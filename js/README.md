# Festivo JavaScript/TypeScript SDK

 Official SDK for the Festivo Public Holidays API.

## Installation

```bash
npm install @festivo-io/festivo-sdk
```

## Usage

```typescript
import { FestivoClient } from '@festivo-io/festivo-sdk';

// Initialize client
const client = new FestivoClient({ apiKey: process.env.FESTIVO_API_KEY });

// Get all holidays for a country
const { holidays } = await client.getHolidays('US', 2026);
console.log(`Found ${holidays.length} holidays`);

// Get city-level holidays (Pro plan)
const milan = await client.getCityHolidays('IT', 'IT-MILAN', 2026);
console.log(milan.holidays);

// Get regional holidays (Builder plan)
const scotland = await client.getRegionalHolidays('GB', 'GB-SCT', 2026);

// Check if a specific date is a holiday
const check = await client.checkHoliday('US', '2026-12-25');
if (check.is_holiday) {
  console.log(`${check.holiday.name} on ${check.holiday.observed}`);
}
```

## API Methods

- `getHolidays(country, year, options?)` - Get all holidays
- `getCityHolidays(country, cityCode, year, options?)` - Get city-level holidays
- `getRegionalHolidays(country, regionCode, year, options?)` - Get regional holidays
- `checkHoliday(country, date, options?)` - Check if date is a holiday

## Development

```bash
cd festivo-sdk/js
npm install
npm run build
```
