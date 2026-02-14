import { FestivoClient } from './src/client.js';

const c = new FestivoClient({ apiKey: process.env.FESTIVO_KEY });
(async () => {
  try {
    const invoice = await c.getInvoice('inv_123');
    console.log(invoice);
  } catch (e) {
    console.error('error', e);
  }
})();
