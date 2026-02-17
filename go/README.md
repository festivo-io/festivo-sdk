# Festivo Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/festivo-io/festivo-sdk-go.svg)](https://pkg.go.dev/github.com/festivo-io/festivo-sdk-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/festivo-io/festivo-sdk)](https://github.com/festivo-io/festivo-sdk)
[![License](https://img.shields.io/github/license/festivo-io/festivo-sdk)](https://github.com/festivo-io/festivo-sdk)

Official Go SDK for the [Festivo Public Holidays API](https://getfestivo.com). Access holiday data for 100+ countries with accurate UTC dates, regional variations, and city-level holidays.

## Features

- 🌍 **100+ Countries** - Comprehensive holiday coverage worldwide
- 📅 **UTC Dates** - Accurate timezone handling for global holidays
- 🏙️ **City & Regional** - Support for local holidays (Pro/Builder plans)
- ⚡ **Standard Library** - Uses only Go standard lib (net/http)
- 🔒 **Type Safe** - Full struct definitions for responses
- 🧪 **Well Tested** - Comprehensive test suite included

## Installation

```bash
go get github.com/festivo-io/festivo-sdk-go
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "os"
    "github.com/festivo-io/festivo-sdk-go/festivo"
)

func main() {
    // Initialize client
    client := festivo.NewClient(os.Getenv("FESTIVO_API_KEY"))

    // Get all holidays for a country
    result, err := client.GetHolidays("US", 2026, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d holidays\n", len(result.Holidays))

    // Get city-level holidays (Pro plan)
    milan, _ := client.GetCityHolidays("IT", "IT-MILAN", 2026, nil)
    fmt.Println(milan.Holidays)

    // Get regional holidays (Builder plan)
    scotland, _ := client.GetRegionalHolidays("GB", "GB-SCT", 2026, nil)

    // Check if a specific date is a holiday
    check, _ := client.CheckHoliday("US", "2026-12-25", "")
    if check.IsHoliday {
        fmt.Printf("%s on %s\n", check.Holiday.Name, check.Holiday.Observed)
    }
}
```

## API Methods

- `GetHolidays(country, year, opts)` - Get all holidays
- `GetCityHolidays(country, cityCode, year, opts)` - Get city-level holidays
- `GetRegionalHolidays(country, regionCode, year, opts)` - Get regional holidays
- `CheckHoliday(country, date, regions)` - Check if date is a holiday

## Development

```bash
cd festivo-sdk/go
go test ./...
go build ./...
```

## Links

- [Official Website](https://getfestivo.com)
- [API Documentation](https://docs.getfestivo.com)
- [GitHub Repository](https://github.com/festivo-io/festivo-sdk)
- [pkg.go.dev](https://pkg.go.dev/github.com/festivo-io/festivo-sdk-go)

## Support

- Email: support@getfestivo.com
- [Report Issues](https://github.com/festivo-io/festivo-sdk/issues)

## License

MIT License - see [LICENSE](../LICENSE) for details.
