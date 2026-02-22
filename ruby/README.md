# Festivo Ruby SDK

Official Ruby SDK for the Festivo Public Holidays API.

## Installation

```bash
gem install festivo # (when published)
```

## Usage

```ruby
require 'festivo'
client = Festivo::Client.new('YOUR_API_KEY')

# Get all holidays for a country and year
holidays = client.get_holidays('US', 2026)

# Get city-level holidays
city_holidays = client.get_city_holidays('IT', 'IT-MILAN', 2026)

# Get regional holidays
regional_holidays = client.get_regional_holidays('GB', 'GB-SCT', 2026)

# Check if a date is a holiday
check = client.check_holiday('US', '2026-12-25')
```

## API Methods

- `get_holidays(country, year, regions: nil, type: nil, language: nil, timezone: nil)`
- `get_city_holidays(country, city_code, year, type: nil, language: nil, timezone: nil)`
- `get_regional_holidays(country, region_code, year, type: nil, language: nil, timezone: nil)`
- `check_holiday(country, date, regions: nil)`

## Development

```bash
cd festivo-sdk/ruby
ruby example.rb
```

## Links

- [Official Website](https://getfestivo.com)
- [API Documentation](https://docs.getfestivo.com)
- [GitHub Repository](https://github.com/festivo-io/festivo-sdk)

## Support

- Email: support@getfestivo.com
- [Report Issues](https://github.com/festivo-io/festivo-sdk/issues)

## License

MIT License - see [LICENSE](../LICENSE) for details.
