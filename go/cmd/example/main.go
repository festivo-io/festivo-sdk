package main

import (
    "fmt"
    "github.com/festivo-io/festivo-sdk-go/festivo"
)

func main() {
    client := festivo.NewClient("YOUR_API_KEY")
    
    // Get holidays for a country
    holidays, err := client.GetHolidays("US", nil)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Holidays: %+v\n", holidays)
    
    // Check if a specific date is a holiday
    check, err := client.CheckHoliday("US", "2026-12-25", nil)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Is holiday: %v\n", check.IsHoliday)
}
