# Festivo PHP SDK

Official PHP SDK for the Festivo Public Holidays API.

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
```
