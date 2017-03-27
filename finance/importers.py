"""A collection of data import functions."""

from logbook import Logger

from finance.models import Asset, AssetValue, get_asset_by_stock_code, \
    Granularity
from finance.providers import Yahoo
from finance.utils import DictReader


log = Logger(__name__)


def import_8percent_data(parsed_data, account_checking, account_8p, asset_krw):
    """Import 8percent `AssetValue`s and `Record`s altogether."""
    from finance.models import Asset, AssetType, AssetValue, Record, \
        Transaction

    assert account_checking
    assert account_8p
    assert asset_krw

    parsed_data = DictReader(parsed_data)
    asset_data = {
        'started_at': parsed_data.started_at.isoformat()
    }
    keys = ['annual_percentage_yield', 'amount', 'grade', 'duration',
            'originator']
    for key in keys:
        asset_data[key] = parsed_data[key]

    asset_8p = Asset.create(name=parsed_data.name, type=AssetType.p2p_bond,
                            data=asset_data)
    remaining_value = parsed_data.amount
    started_at = parsed_data.started_at

    with Transaction.create() as t:
        Record.create(
            created_at=started_at, transaction=t, account=account_checking,
            asset=asset_krw, quantity=-remaining_value)
        Record.create(
            created_at=started_at, transaction=t, account=account_8p,
            asset=asset_8p, quantity=1)
    AssetValue.create(
        evaluated_at=started_at, asset=asset_8p,
        base_asset=asset_krw, granularity='1day', close=remaining_value)

    for record in parsed_data.records:
        date, principle, interest, tax, fees = record
        returned = principle + interest - (tax + fees)
        remaining_value -= principle
        with Transaction.create() as t:
            Record.create(
                created_at=date, transaction=t,
                account=account_checking, asset=asset_krw, quantity=returned)
        AssetValue.create(
            evaluated_at=date, asset=asset_8p,
            base_asset=asset_krw, granularity='1day', close=remaining_value)


def import_stock_values(code, from_date, to_date):
    provider = Yahoo()
    asset = get_asset_by_stock_code(code)
    data = provider.fetch_data(code, from_date, to_date)

    # NOTE: Any better way to handle this?
    asset_krw = Asset.query.filter_by(name='KRW').first()

    for date, open_, high, low, close_, volume, adj_close in data:
        log.info('AssetValue(date={}, asset={}, open={}, close={}',
                 date, asset, open_, close_)
        AssetValue.create(
            evaluated_at=date, granularity=Granularity.day, asset=asset,
            base_asset=asset_krw, open=open_, high=high, low=low, close=close_,
            volume=volume)
