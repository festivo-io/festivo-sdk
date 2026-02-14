# Festivo PHP SDK (starter)

Install dependencies with Composer:

```bash
cd festivo-sdk/php
composer install
```

Usage:

```php
require 'vendor/autoload.php';
use Festivo\FestivoClient;

$c = new FestivoClient(getenv('FESTIVO_KEY'));
$invoice = $c->getInvoice('inv_123');
print_r($invoice);
```
