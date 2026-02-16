# Festivo Go SDK (starter)

Usage example:

```go
package main

import (
    "fmt"
    "github.com/festivo-io/festivo-sdk-go/festivo"
)

func main() {
    client := festivo.NewClient("YOUR_API_KEY")
    invoice, err := client.GetInvoice("inv_123")
    if err != nil {
        panic(err)
    }
    fmt.Println(invoice)
}
```
