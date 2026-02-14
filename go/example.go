package main

import (
    "fmt"
    "github.com/getfestivo/festivo-sdk-go/festivo"
)

func main() {
    c := festivo.NewClient("YOUR_API_KEY")
    invoice, err := c.GetInvoice("inv_123")
    if err != nil {
        panic(err)
    }
    fmt.Println(invoice)
}
