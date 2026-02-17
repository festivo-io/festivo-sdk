# Festivo PHP SDK

[![Latest Version](https://img.shields.io/packagist/v/festivo-io/festivo-php.svg)](https://packagist.org/packages/festivo-io/festivo-php)
[![License](https://img.shields.io/packagist/l/festivo-io/festivo-php.svg)](https://packagist.org/packages/festivo-io/festivo-php)
[![PHP Version](https://img.shields.io/packagist/php-v/festivo-io/festivo-php.svg)](https://packagist.org/packages/festivo-io/festivo-php)

Official PHP SDK for the [Festivo Public Holidays API](https://getfestivo.com). Access holiday data for 100+ countries with accurate UTC dates, regional variations, and city-level holidays.

## Features

- 🌍 **100+ Countries** - Comprehensive holiday coverage worldwide
- 📅 **UTC Dates** - Accurate timezone handling for global holidays
- 🏙️ **City & Regional** - Support for local holidays (Pro/Builder plans)
- ✅ **Type Safe** - Full type hints for better IDE support
- ⚡ **Modern PHP** - Built for PHP 8.0+
- 🧪 **Well Tested** - Comprehensive test suite included

## Installation

```bash
composer require festivo-io/festivo-php
```

## Usage

```php
<?php
require 'vendor/autoload.php';
use Festivo\FestivoClient;

// Initialize client
$client = new FestivoClient(getenv('FESTIVO_API_KEY'));

// Get all holidays for a country
$result = $client->getHolidays('US', 2026);
echo "Found " . count($result['holidays']) . " holidays\n";

// Get city-level holidays (Pro plan)
$milan = $client->getCityHolidays('IT', 'IT-MILAN', 2026);
print_r($milan['holidays']);

// Get regional holidays (Builder plan)
$scotland = $client->getRegionalHolidays('GB', 'GB-SCT', 2026);

// Check if a specific date is a holiday
$check = $client->checkHoliday('US', '2026-12-25');
if ($check['is_holiday']) {
    echo $check['holiday']['name'] . " on " . $check['holiday']['observed'] . "\n";
}
```

## API Methods

- `getHolidays($country, $year, $options = [])` - Get all holidays
- `getCityHolidays($country, $cityCode, $year, $options = [])` - Get city-level holidays
- `getRegionalHolidays($country, $regionCode, $year, $options = [])` - Get regional holidays
- `checkHoliday($country, $date, $regions = null)` - Check if date is a holiday

## Development

```bash
cd festivo-sdk/php
composer install
composer test
```

## Links

- [Official Website](https://getfestivo.com)
- [API Documentation](https://docs.getfestivo.com)
- [GitHub Repository](https://github.com/festivo-io/festivo-sdk)
- [Packagist Package](https://packagist.org/packages/festivo-io/festivo-php)

## Support

- Email: support@getfestivo.com
- [Report Issues](https://github.com/festivo-io/festivo-sdk/issues)

## License

MIT License - see [LICENSE](../LICENSE) for details.
