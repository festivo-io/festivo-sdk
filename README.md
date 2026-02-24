# Festivo SDKs

[![License](https://img.shields.io/github/license/festivo-io/festivo-sdk)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/festivo-io/festivo-sdk)](https://github.com/festivo-io/festivo-sdk/releases)

Official SDKs for the [Festivo Public Holidays API](https://getfestivo.com) - Access holiday data for 250+ countries with accurate UTC dates, regional variations, and city-level holidays.

## Available SDKs

| Language | Package | Version | Documentation |
|----------|---------|---------|---------------|
| **JavaScript/TypeScript** | [@festivo-io/festivo-sdk](https://www.npmjs.com/package/@festivo-io/festivo-sdk) | [![npm](https://img.shields.io/npm/v/@festivo-io/festivo-sdk.svg)](https://www.npmjs.com/package/@festivo-io/festivo-sdk) | [README](js/README.md) |
| **Python** | [festivo-python](https://pypi.org/project/festivo-python/) | [![PyPI](https://img.shields.io/pypi/v/festivo-python.svg)](https://pypi.org/project/festivo-python/) | [README](python/README.md) |
| **PHP** | [festivo-io/festivo-php](https://packagist.org/packages/festivo-io/festivo-php) | [![Packagist](https://img.shields.io/packagist/v/festivo-io/festivo-php.svg)](https://packagist.org/packages/festivo-io/festivo-php) | [README](php/README.md) |
| **Go** | [festivo-sdk-go](https://pkg.go.dev/github.com/festivo-io/festivo-sdk-go) | [![Go](https://img.shields.io/github/v/tag/festivo-io/festivo-sdk)](https://github.com/festivo-io/festivo-sdk/tags) | [README](go/README.md) |
| **Ruby** | festivo | Coming Soon | [README](ruby/README.md) |
| **Java** | com.festivo:festivo-sdk | Coming Soon | [README](java/README.md) |

## Features

- 🌍 **250+ Countries** - Comprehensive holiday coverage worldwide
- 📅 **UTC Dates** - Accurate timezone handling for global holidays
- 🏙️ **City & Regional** - Support for local holidays (Pro/Builder plans)
- ✅ **Type Safe** - Full type definitions in all supported languages
- ⚡ **Modern** - Built with current best practices for each language
- 🧪 **Well Tested** - Comprehensive test suites included

## Quick Start

### JavaScript/TypeScript

```bash
npm install @festivo-io/festivo-sdk
```

```typescript
import { FestivoClient } from '@festivo-io/festivo-sdk';
const client = new FestivoClient({ apiKey: 'YOUR_API_KEY' });
const { holidays } = await client.getHolidays('US', 2026);
```

### Python

```bash
pip install festivo-python
```

```python
from festivo import FestivoClient
client = FestivoClient(api_key='YOUR_API_KEY')
result = client.get_holidays('US', 2026)
```

### PHP

```bash
composer require festivo-io/festivo-php
```

```php
use Festivo\FestivoClient;
$client = new FestivoClient('YOUR_API_KEY');
$result = $client->getHolidays('US', 2026);
```

### Go

```bash
go get github.com/festivo-io/festivo-sdk-go
```

```go
import "github.com/festivo-io/festivo-sdk-go/festivo"
client := festivo.NewClient("YOUR_API_KEY")
result, _ := client.GetHolidays("US", 2026, nil)
```

## Documentation

- [Official Website](https://getfestivo.com)
- [API Documentation](https://docs.getfestivo.com)
- [Get Your API Key](https://getfestivo.com/dashboard)

## Support

- Email: support@getfestivo.com
- [Report Issues](https://github.com/festivo-io/festivo-sdk/issues)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) for details.
