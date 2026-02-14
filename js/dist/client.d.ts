export type Config = {
    apiKey?: string;
    baseUrl?: string;
};
export declare class FestivoClient {
    baseUrl: string;
    apiKey?: string;
    constructor(config?: Config);
    request(path: string, method?: string, body?: any): Promise<any>;
    getInvoice(id: string): Promise<any>;
}
