# Festivo Go SDK

Official Go SDK for the Festivo Public Holidays API.

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
