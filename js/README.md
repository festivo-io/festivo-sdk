# JavaScript SDK (starter)

Install / build locally:

```bash
cd festivo-sdk/js
npm install
npm run build
```

Basic usage (Node 18+ or browser):

```js
import { FestivoClient } from './dist/index.js';

const c = new FestivoClient({ apiKey: process.env.FESTIVO_KEY });
const invoice = await c.getInvoice('inv_123');
console.log(invoice);
```
