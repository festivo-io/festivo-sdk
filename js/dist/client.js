export class FestivoClient {
    constructor(config = {}) {
        this.baseUrl = config.baseUrl ?? "https://api.getfestivo.com";
        this.apiKey = config.apiKey;
    }
    async request(path, method = "GET", body) {
        const headers = { "Accept": "application/json" };
        if (this.apiKey)
            headers["Authorization"] = `Bearer ${this.apiKey}`;
        if (body)
            headers["Content-Type"] = "application/json";
        const res = await fetch(`${this.baseUrl}${path}`, {
            method,
            headers,
            body: body ? JSON.stringify(body) : undefined,
        });
        return res.json();
    }
    async getInvoice(id) {
        return this.request(`/invoices/${id}`);
    }
}
