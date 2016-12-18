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
            invested_amount = record['Invested']
            if invested_amount:
                # NOTE: int() or float()?
                invested_amount = int(extract_numbers(invested_amount))
                date = parse_date(record['Date'], DATE_FORMAT)
                bond_name = record['Name']

                yield date, bond_name, invested_amount, currency
