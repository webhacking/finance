import json

from pandas import DataFrame, to_datetime
from pandas_datareader.base import _DailyBaseReader
from pandas_datareader.yahoo.daily import _adjust_prices, _calc_return_index


class LocalDailyReader(_DailyBaseReader):
    """The basic idea is to read data locally to construct a DataFrame. Most of
    code was copied from pandas_datareader.yahoo.daily.YahooDailyReader. A lot
    of unnecessary code remains unchanged for compatibility reasons.
    """

    def __init__(self, symbols=None, start=None, end=None, retry_count=3,
                 pause=0.1, session=None, adjust_price=False,
                 ret_index=False, chunksize=1, interval='d',
                 get_actions=False, adjust_dividends=False):
        super(LocalDailyReader, self).__init__(
            symbols=symbols, start=start, end=end, retry_count=retry_count,
            pause=pause, session=session, chunksize=chunksize)

        self.adjust_price = adjust_price
        self.ret_index = ret_index
        self.interval = interval
        self._get_actions = get_actions
        self.interval = '1' + self.interval
        self.adjust_dividends = adjust_dividends

    @property
    def get_actions(self):
        return self._get_actions

    @property
    def url(self):
        return None

    def _get_params(self, symbol):
        return {
            'symbol': symbol,
        }

    def _read_one_data(self, url, params):

        symbol = params['symbol']

        with open(f'{symbol}.json') as fin:
            raw = fin.read()
            data = json.loads(raw)

        # price data
        prices = DataFrame(data['prices'])
        prices.columns = [col.capitalize() for col in prices.columns]
        prices['Date'] = to_datetime(
            to_datetime(prices['Date'], unit='s').dt.date)

        if 'Data' in prices.columns:
            prices = prices[prices['Data'].isnull()]
        prices = prices[['Date', 'High', 'Low', 'Open', 'Close', 'Volume',
                         'Adjclose']]
        prices = prices.rename(columns={'Adjclose': 'Adj Close'})

        prices = prices.set_index('Date')
        prices = prices.sort_index().dropna(how='all')

        if self.ret_index:
            prices['Ret_Index'] = \
                _calc_return_index(prices['Adj Close'])
        if self.adjust_price:
            prices = _adjust_prices(prices)

        # dividends & splits data
        if self.get_actions and data['eventsData']:

            actions = DataFrame(data['eventsData'])
            actions.columns = [col.capitalize() for col in actions.columns]
            actions['Date'] = to_datetime(
                to_datetime(actions['Date'], unit='s').dt.date)

            types = actions['Type'].unique()
            if 'DIVIDEND' in types:
                divs = actions[actions.Type == 'DIVIDEND'].copy()
                divs = divs[['Date', 'Amount']].reset_index(drop=True)
                divs = divs.set_index('Date')
                divs = divs.rename(columns={'Amount': 'Dividends'})
                prices = prices.join(divs, how='outer')

            if 'SPLIT' in types:

                def split_ratio(row):
                    if float(row['Numerator']) > 0:
                        return eval(row['Splitratio'])
                    else:
                        return 1

                splits = actions[actions.Type == 'SPLIT'].copy()
                splits['SplitRatio'] = splits.apply(split_ratio, axis=1)
                splits = splits.reset_index(drop=True)
                splits = splits.set_index('Date')
                splits['Splits'] = splits['SplitRatio']
                prices = prices.join(splits['Splits'], how='outer')

                if 'DIVIDEND' in types and self.adjust_dividends:
                    # Adjust dividends to deal with splits
                    adj = prices['Splits'].sort_index(ascending=False).fillna(
                        1).cumprod()
                    prices['Dividends'] = prices['Dividends'] * adj

        return prices
