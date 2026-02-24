# Festivo JavaScript/TypeScript SDK

[![npm version](https://img.shields.io/npm/v/@festivo-io/festivo-sdk.svg)](https://www.npmjs.com/package/@festivo-io/festivo-sdk)
[![License](https://img.shields.io/npm/l/@festivo-io/festivo-sdk.svg)](https://www.npmjs.com/package/@festivo-io/festivo-sdk)
[![TypeScript](https://img.shields.io/badge/TypeScript-Ready-blue.svg)](https://www.typescriptlang.org/)

Official JavaScript/TypeScript SDK for the [Festivo Public Holidays API](https://getfestivo.com). Access holiday data for 250+ countries with accurate UTC dates, regional variations, and city-level holidays.

## Features

- 🌍 **250+ Countries** - Comprehensive holiday coverage worldwide
- 📅 **UTC Dates** - Accurate timezone handling for global holidays
- 🏙️ **City & Regional** - Support for local holidays (Pro/Builder plans)
- 💪 **TypeScript Native** - Full type definitions included
- ⚡ **Modern ESM** - ES Modules with tree-shaking support
- 🧪 **Well Tested** - Comprehensive test suite with Vitest

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
npm test
```

## Links

- [Official Website](https://getfestivo.com)
- [API Documentation](https://docs.getfestivo.com)
- [GitHub Repository](https://github.com/festivo-io/festivo-sdk)
- [npm Package](https://www.npmjs.com/package/@festivo-io/festivo-sdk)

## Support

- Email: support@getfestivo.com
- [Report Issues](https://github.com/festivo-io/festivo-sdk/issues)

## License

MIT License - see [LICENSE](../LICENSE) for details.
