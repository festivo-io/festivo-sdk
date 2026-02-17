import pytest
from unittest.mock import Mock, patch, MagicMock
from festivo.client import FestivoClient


class TestFestivoClient:
    @pytest.fixture
    def client(self):
        """Create a test client instance."""
        return FestivoClient(api_key="test-key")

    @pytest.fixture
    def mock_response(self):
        """Create a mock successful response."""
        mock_resp = Mock()
        mock_resp.raise_for_status = Mock()
        mock_resp.json = Mock(return_value={"holidays": [], "total": 0})
        return mock_resp

    def test_initialization(self):
        """Test client initialization with default values."""
        client = FestivoClient(api_key="test-key")
        assert client.api_key == "test-key"
        assert client.base_url == "https://api.getfestivo.com"

    def test_initialization_custom_base_url(self):
        """Test client initialization with custom base URL."""
        client = FestivoClient(api_key="test-key", base_url="https://custom.api.com")
        assert client.base_url == "https://custom.api.com"

    def test_initialization_env_var(self, monkeypatch):
        """Test client reads API key from environment variable."""
        monkeypatch.setenv("FESTIVO_API_KEY", "env-key")
        client = FestivoClient()
        assert client.api_key == "env-key"

    def test_headers_with_api_key(self, client):
        """Test headers include authorization when API key is set."""
        headers = client._headers()
        assert headers["Accept"] == "application/json"
        assert headers["Authorization"] == "Bearer test-key"

    def test_headers_without_api_key(self):
        """Test headers without authorization when no API key."""
        client = FestivoClient()
        headers = client._headers()
        assert headers["Accept"] == "application/json"
        assert "Authorization" not in headers

    @patch("festivo.client.requests.get")
    def test_get_holidays_basic(self, mock_get, client, mock_response):
        """Test basic getHolidays call."""
        mock_get.return_value = mock_response

        result = client.get_holidays("US", 2026)

        mock_get.assert_called_once()
        call_url, call_kwargs = mock_get.call_args
        assert call_url[0] == "https://api.getfestivo.com/v3/public-holidays/list"
        assert call_kwargs["params"]["country"] == "US"
        assert call_kwargs["params"]["year"] == 2026
        assert result == {"holidays": [], "total": 0}

    @patch("festivo.client.requests.get")
    def test_get_holidays_with_options(self, mock_get, client, mock_response):
        """Test getHolidays with optional parameters."""
        mock_get.return_value = mock_response

        client.get_holidays(
            "GB", 2026, regions="GB-SCT", type="public", language="en", timezone="Europe/London"
        )

        call_kwargs = mock_get.call_args[1]
        params = call_kwargs["params"]
        assert params["country"] == "GB"
        assert params["year"] == 2026
        assert params["regions"] == "GB-SCT"
        assert params["type"] == "public"
        assert params["language"] == "en"
        assert params["timezone"] == "Europe/London"

    @patch("festivo.client.requests.get")
    def test_get_city_holidays(self, mock_get, client, mock_response):
        """Test getCityHolidays method."""
        mock_get.return_value = mock_response

        client.get_city_holidays("IT", "IT-MILAN", 2026)

        call_kwargs = mock_get.call_args[1]
        params = call_kwargs["params"]
        assert params["country"] == "IT"
        assert params["year"] == 2026
        assert params["regions"] == "IT-MILAN"

    @patch("festivo.client.requests.get")
    def test_get_city_holidays_with_options(self, mock_get, client, mock_response):
        """Test getCityHolidays with additional options."""
        mock_get.return_value = mock_response

        client.get_city_holidays("IT", "IT-ROME", 2026, type="public", language="it")

        call_kwargs = mock_get.call_args[1]
        params = call_kwargs["params"]
        assert params["regions"] == "IT-ROME"
        assert params["type"] == "public"
        assert params["language"] == "it"

    @patch("festivo.client.requests.get")
    def test_get_regional_holidays(self, mock_get, client, mock_response):
        """Test getRegionalHolidays method."""
        mock_get.return_value = mock_response

        client.get_regional_holidays("GB", "GB-SCT", 2026)

        call_kwargs = mock_get.call_args[1]
        params = call_kwargs["params"]
        assert params["country"] == "GB"
        assert params["year"] == 2026
        assert params["regions"] == "GB-SCT"

    @patch("festivo.client.requests.get")
    def test_check_holiday(self, mock_get, client):
        """Test checkHoliday method."""
        mock_resp = Mock()
        mock_resp.raise_for_status = Mock()
        mock_resp.json = Mock(return_value={"is_holiday": True})
        mock_get.return_value = mock_resp

        result = client.check_holiday("US", "2026-12-25")

        call_url = mock_get.call_args[0][0]
        call_kwargs = mock_get.call_args[1]
        assert "/v3/public-holidays/list/check" in call_url
        assert call_kwargs["params"]["country"] == "US"
        assert call_kwargs["params"]["date"] == "2026-12-25"
        assert result["is_holiday"] is True

    @patch("festivo.client.requests.get")
    def test_check_holiday_with_regions(self, mock_get, client):
        """Test checkHoliday with regions parameter."""
        mock_resp = Mock()
        mock_resp.raise_for_status = Mock()
        mock_resp.json = Mock(return_value={"is_holiday": True})
        mock_get.return_value = mock_resp

        client.check_holiday("IT", "2026-12-07", regions="IT-MILAN")

        call_kwargs = mock_get.call_args[1]
        assert call_kwargs["params"]["regions"] == "IT-MILAN"

    @patch("festivo.client.requests.get")
    def test_api_error_handling(self, mock_get, client):
        """Test that API errors are properly raised."""
        mock_resp = Mock()
        mock_resp.raise_for_status.side_effect = Exception("401 Unauthorized")
        mock_get.return_value = mock_resp

        with pytest.raises(Exception, match="401 Unauthorized"):
            client.get_holidays("US", 2026)

    @patch("festivo.client.requests.get")
    def test_response_structure(self, mock_get, client):
        """Test that response structure matches expected format."""
        mock_resp = Mock()
        mock_resp.raise_for_status = Mock()
        mock_resp.json = Mock(
            return_value={
                "holidays": [
                    {
                        "date": "2026-01-01",
                        "name": "New Year's Day",
                        "type": "public",
                        "observed": "2026-01-01",
                        "public": True,
                        "country": "US",
                        "subdivisions": [],
                    }
                ],
                "total": 1,
            }
        )
        mock_get.return_value = mock_resp

        result = client.get_holidays("US", 2026)

        assert "holidays" in result
        assert "total" in result
        assert isinstance(result["holidays"], list)
        assert isinstance(result["total"], int)
        assert len(result["holidays"]) == 1
        assert result["holidays"][0]["date"] == "2026-01-01"
