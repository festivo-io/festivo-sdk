# Festivo Python SDK (starter)

Install locally using Poetry:

```bash
cd festivo-sdk/python
poetry install
```

Usage:

```py
from festivo import FestivoClient

c = FestivoClient()
invoice = c.get_invoice('inv_123')
print(invoice)
```
