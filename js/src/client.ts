export type Config = {
  apiKey?: string;
  baseUrl?: string;
};

export class FestivoClient {
  baseUrl: string;
  apiKey?: string;

  constructor(config: Config = {}) {
    this.baseUrl = config.baseUrl ?? "https://api.getfestivo.com";
    this.apiKey = config.apiKey;
  }

  async request(path: string, method = "GET", body?: any) {
    const headers: Record<string, string> = { "Accept": "application/json" };
    if (this.apiKey) headers["Authorization"] = `Bearer ${this.apiKey}`;
    if (body) headers["Content-Type"] = "application/json";

    const res = await fetch(`${this.baseUrl}${path}`, {
      method,
      headers,
      body: body ? JSON.stringify(body) : undefined,
    });
    return res.json();
  }

  async getInvoice(id: string) {
    return this.request(`/invoices/${id}`);
  }
}
