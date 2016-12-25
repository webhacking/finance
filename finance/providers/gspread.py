"""Import data from Google spreadsheets."""

import os

import gspread
from oauth2client.service_account import ServiceAccountCredentials

from finance.providers.provider import Provider
from finance.utils import extract_numbers, parse_date


DATE_FORMAT = '%m/%d/%Y'


class GSpread(Provider):
    scopes = ['https://spreadsheets.google.com/feeds']
    gc = None

    def auth(self):
        key_file_path = os.environ['GSPREAD_KEY_FILE_PATH']
        credentials = ServiceAccountCredentials.from_json_keyfile_name(
            key_file_path, self.scopes)
        self.gc = gspread.authorize(credentials)

    def parse_int(self, raw):
        try:
            return int(extract_numbers(raw))
        except ValueError:
            return 0

    def fetch_data(self):
        if self.gc is None:
            self.auth()
        doc_key = os.environ['GSPREAD_DOC_KEY']
        doc = self.gc.open_by_key(doc_key)
        sheet = doc.worksheet('P2P Bonds')

        # FIMXE: Do not use a hard-coded value
        currency = 'KRW'

        # NOTE: sheet.get_all_records() seems to return a list of dict
        # containing all cell values. This may have significant performance
        # implications.
        for record in sheet.get_all_records():
            bond_name = record['Name']
            if not bond_name:
                continue

            date = parse_date(record['Date'], DATE_FORMAT)
            invested_amount = self.parse_int(record['Invested'])
            returned_amount = self.parse_int(record['Returned'])

            if invested_amount:
                yield 'invested', date, bond_name, invested_amount, 0, 0, 0, \
                    currency
            elif returned_amount:
                principle = self.parse_int(record['Returned Principle'])
                interest = self.parse_int(record['Returned Interest'])
                tax = self.parse_int(record['Tax'])
                fees = self.parse_int(record['Fees'])

                assert principle + interest - tax - fees == returned_amount, \
                    'Incorrect returned amount: {} {}'.format(
                        date, bond_name)

                yield 'returned', date, bond_name, principle, interest, tax, \
                    fees, currency
