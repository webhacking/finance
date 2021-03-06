from datetime import datetime
import os

from logbook import Logger
from sqlalchemy.exc import IntegrityError, InvalidRequestError

from finance import create_app
from finance.exceptions import AssetNotFoundException
from finance.fetchers import fetch_stock_values
from finance.models import Asset, AssetType, AssetValue, db, Granularity
from finance.providers import Yahoo
from finance.utils import (
    date_to_datetime,
    parse_date,
)


log = Logger("finance")


# TODO: Write logs to CloudWatch


def fetch_asset_values(code, start_time, end_time):
    try:
        asset = Asset.get_by_symbol(code)
    except AssetNotFoundException:
        log.info("Asset {0} does not exist. Creating an Asset record...", code)
        asset = Asset.create(name=code, code=code, type=AssetType.stock)

    rows = fetch_stock_values(code, start_time, end_time, Granularity.min)

    for date, open_, high, low, close_, volume, source in rows:
        insert_asset_value(
            asset, date, Granularity.min, open_, high, low, close_, volume, source
        )

    try:
        db.session.commit()
    except (IntegrityError, InvalidRequestError):
        log.exception("Something went wrong")
        db.session.rollback()
    else:
        log.info("Asset values for {0} have been imported", code)


def insert_asset_value(
    asset, date, granularity, open_, high, low, close_, volume, source
):
    # FIXME: This kind of approach may not be safe in multithreading
    # environments
    if AssetValue.exists(asset_id=asset.id, evaluated_at=date, granularity=granularity):
        log.warn("AssetValue for {0} on {1} already exist", asset.code, date)
    else:
        asset_value = AssetValue.create(
            evaluated_at=date,
            granularity=Granularity.min,
            asset=asset,
            open=open_,
            high=high,
            low=low,
            close=close_,
            volume=int(volume),
            source=source,
            commit=False,
        )
        log.info("Record has been create: {0}", asset_value)


# TODO: Have a list of stock symbols to be fetched


if __name__ == "__main__":
    request_import_stock_values_handler({}, None)
