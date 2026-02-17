# Festivo Python SDK

Official Python SDK for the Festivo Public Holidays API.

## Installation

```bash
pip install festivo-python
```

## Usage

```python
from festivo import FestivoClient
import os

# Initialize client
client = FestivoClient(api_key=os.getenv("FESTIVO_API_KEY"))

# Get all holidays for a country
result = client.get_holidays('US', 2026)
print(f"Found {len(result['holidays'])} holidays")

# Get city-level holidays (Pro plan)
milan = client.get_city_holidays('IT', 'IT-MILAN', 2026)
print(milan['holidays'])

# Get regional holidays (Builder plan)
scotland = client.get_regional_holidays('GB', 'GB-SCT', 2026)

# Check if a specific date is a holiday
check = client.check_holiday('US', '2026-12-25')
if check['is_holiday']:
    print(f"{check['holiday']['name']} on {check['holiday']['observed']}")
```

## API Methods

- `get_holidays(country, year, **options)` - Get all holidays
- `get_city_holidays(country, city_code, year, **options)` - Get city-level holidays
- `get_regional_holidays(country, region_code, year, **options)` - Get regional holidays
- `check_holiday(country, date, regions=None)` - Check if date is a holiday

## Development

For local development with Poetry:

```bash
cd festivo-sdk/python
poetry install
poetry build
```
