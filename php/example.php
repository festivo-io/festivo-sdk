<?php
require __DIR__ . '/vendor/autoload.php';
use Festivo\FestivoClient;

$c = new FestivoClient(getenv('FESTIVO_KEY'));
print_r($c->getInvoice('inv_123'));
