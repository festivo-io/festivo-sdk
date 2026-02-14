import os
import requests
from typing import Optional


class FestivoClient:
    def __init__(self, api_key: Optional[str] = None, base_url: str = "https://api.getfestivo.com"):
        self.api_key = api_key or os.getenv("FESTIVO_KEY")
        self.base_url = base_url

    def _headers(self):
        h = {"Accept": "application/json"}
        if self.api_key:
            h["Authorization"] = f"Bearer {self.api_key}"
        return h

    def get_invoice(self, invoice_id: str):
        resp = requests.get(f"{self.base_url}/invoices/{invoice_id}", headers=self._headers())
        resp.raise_for_status()
        return resp.json()
