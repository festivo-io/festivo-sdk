import os
import requests
from typing import Optional, Dict, List, Any


class FestivoClient:
    """
    Festivo API client for accessing public holiday data.
    """

    def __init__(self, api_key: Optional[str] = None, base_url: str = "https://api.getfestivo.com"):
        """
        Initialize the Festivo API client.
        
        Args:
            api_key: Your Festivo API key (or set FESTIVO_API_KEY env var)
            base_url: API base URL (default: https://api.getfestivo.com)
        """
        self.api_key = api_key or os.getenv("FESTIVO_API_KEY")
        self.base_url = base_url

    def _headers(self) -> Dict[str, str]:
        h = {"Accept": "application/json"}
        if self.api_key:
            h["Authorization"] = f"Bearer {self.api_key}"
        return h

    def _request(self, path: str, params: Optional[Dict[str, Any]] = None) -> Dict[str, Any]:
        """Internal request method."""
        url = f"{self.base_url}{path}"
        resp = requests.get(url, headers=self._headers(), params=params or {})
        resp.raise_for_status()
        return resp.json()

    def get_holidays(
        self,
        country: str,
        year: int,
        regions: Optional[str] = None,
        type: Optional[str] = None,
        language: Optional[str] = None,
        timezone: Optional[str] = None,
    ) -> Dict[str, Any]:
        """
        Get all holidays for a country and year.
        
        Args:
            country: ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "IT")
            year: Year (e.g., 2026)
            regions: Optional comma-separated region/city codes
            type: Optional holiday type filter (e.g., "public", "bank")
            language: Optional language code for holiday names
            timezone: Optional IANA timezone (e.g., "America/New_York")
            
        Returns:
            Dict with 'holidays' list and 'total' count
        """
        params = {"country": country, "year": year}
        if regions:
            params["regions"] = regions
        if type:
            params["type"] = type
        if language:
            params["language"] = language
        if timezone:
            params["timezone"] = timezone
        return self._request("/v3/public-holidays/list", params)

    def get_city_holidays(
        self,
        country: str,
        city_code: str,
        year: int,
        type: Optional[str] = None,
        language: Optional[str] = None,
        timezone: Optional[str] = None,
    ) -> Dict[str, Any]:
        """
        Get holidays for a specific city (Pro plan).
        
        Args:
            country: ISO 3166-1 alpha-2 country code (e.g., "IT")
            city_code: City code in format {COUNTRY}-{CITY} (e.g., "IT-MILAN")
            year: Year (e.g., 2026)
            type: Optional holiday type filter
            language: Optional language code for holiday names
            timezone: Optional IANA timezone
            
        Returns:
            Dict with 'holidays' list and 'total' count
        """
        return self.get_holidays(country, year, regions=city_code, type=type, language=language, timezone=timezone)

    def get_regional_holidays(
        self,
        country: str,
        region_code: str,
        year: int,
        type: Optional[str] = None,
        language: Optional[str] = None,
        timezone: Optional[str] = None,
    ) -> Dict[str, Any]:
        """
        Get holidays for a specific region using ISO 3166-2 codes (Builder plan).
        
        Args:
            country: ISO 3166-1 alpha-2 country code (e.g., "GB")
            region_code: ISO 3166-2 subdivision code (e.g., "GB-SCT" for Scotland)
            year: Year (e.g., 2026)
            type: Optional holiday type filter
            language: Optional language code for holiday names
            timezone: Optional IANA timezone
            
        Returns:
            Dict with 'holidays' list and 'total' count
        """
        return self.get_holidays(country, year, regions=region_code, type=type, language=language, timezone=timezone)

    def check_holiday(
        self,
        country: str,
        date: str,
        regions: Optional[str] = None,
    ) -> Dict[str, Any]:
        """
        Check if a specific date is a holiday.
        
        Args:
            country: ISO 3166-1 alpha-2 country code (e.g., "US")
            date: Date in format YYYY-MM-DD (e.g., "2026-12-25")
            regions: Optional comma-separated region/city codes
            
        Returns:
            Dict with 'is_holiday' bool and optional 'holiday' dict
        """
        params = {"country": country, "date": date}
        if regions:
            params["regions"] = regions
        return self._request("/v3/public-holidays/check", params)

