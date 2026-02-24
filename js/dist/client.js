export class FestivoClient {
    constructor(config = {}) {
        this.baseUrl = config.baseUrl ?? "https://api.getfestivo.com";
        this.apiKey = config.apiKey;
    }
    async request(path, params) {
        const headers = { "Accept": "application/json" };
        if (this.apiKey)
            headers["X-API-Key"] = this.apiKey;
        const url = new URL(`${this.baseUrl}${path}`);
        if (params) {
            Object.entries(params).forEach(([key, value]) => {
                if (value !== undefined && value !== null) {
                    url.searchParams.append(key, String(value));
                }
            });
        }
        const res = await fetch(url.toString(), { headers });
        if (!res.ok) {
            throw new Error(`API error: ${res.status} ${res.statusText}`);
        }
        return res.json();
    }
    /**
     * Get all holidays for a country and year
     * @param country ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "IT")
     * @param year Year (e.g., 2026)
     * @param options Optional filters (regions, type, language, timezone)
     */
    async getHolidays(country, year, options) {
        return this.request("/v3/public-holidays/list", {
            country,
            year,
            ...options,
        });
    }
    /**
     * Get holidays for a specific city (Pro plan)
     * @param country ISO 3166-1 alpha-2 country code (e.g., "IT")
     * @param cityCode City code in format {COUNTRY}-{CITY} (e.g., "IT-MILAN")
     * @param year Year (e.g., 2026)
     * @param options Optional filters (type, language, timezone)
     */
    async getCityHolidays(country, cityCode, year, options) {
        return this.request("/v3/public-holidays/list", {
            country,
            year,
            regions: cityCode,
            ...options,
        });
    }
    /**
     * Get holidays for a specific region using ISO 3166-2 subdivision codes (Builder plan)
     * @param country ISO 3166-1 alpha-2 country code (e.g., "GB")
     * @param regionCode ISO 3166-2 subdivision code (e.g., "GB-SCT" for Scotland)
     * @param year Year (e.g., 2026)
     * @param options Optional filters (type, language, timezone)
     */
    async getRegionalHolidays(country, regionCode, year, options) {
        return this.request("/v3/public-holidays/list", {
            country,
            year,
            regions: regionCode,
            ...options,
        });
    }
    /**
     * Check if a specific date is a holiday
     * @param country ISO 3166-1 alpha-2 country code (e.g., "US")
     * @param date Date in format YYYY-MM-DD (e.g., "2026-12-25")
     * @param options Optional filters (regions)
     */
    async checkHoliday(country, date, options) {
        return this.request("/v3/public-holidays/check", {
            country,
            date,
            ...options,
        });
    }
}
