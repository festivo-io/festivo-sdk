export type Config = {
    apiKey?: string;
    baseUrl?: string;
};
export type HolidayOptions = {
    regions?: string;
    type?: string;
    language?: string;
    timezone?: string;
};
export type Holiday = {
    date: string;
    name: string;
    name_local?: string;
    type: string;
    observed: string;
    public: boolean;
    country: string;
    subdivisions: string[];
    regions?: Array<{
        code: string;
        type: string;
    }>;
};
export type HolidaysResponse = {
    holidays: Holiday[];
    total: number;
};
export type CheckHolidayResponse = {
    is_holiday: boolean;
    holiday?: Holiday;
};
export declare class FestivoClient {
    baseUrl: string;
    apiKey?: string;
    constructor(config?: Config);
    private request;
    /**
     * Get all holidays for a country and year
     * @param country ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "IT")
     * @param year Year (e.g., 2026)
     * @param options Optional filters (regions, type, language, timezone)
     */
    getHolidays(country: string, year: number, options?: HolidayOptions): Promise<HolidaysResponse>;
    /**
     * Get holidays for a specific city (Pro plan)
     * @param country ISO 3166-1 alpha-2 country code (e.g., "IT")
     * @param cityCode City code in format {COUNTRY}-{CITY} (e.g., "IT-MILAN")
     * @param year Year (e.g., 2026)
     * @param options Optional filters (type, language, timezone)
     */
    getCityHolidays(country: string, cityCode: string, year: number, options?: Omit<HolidayOptions, "regions">): Promise<HolidaysResponse>;
    /**
     * Get holidays for a specific region using ISO 3166-2 subdivision codes (Builder plan)
     * @param country ISO 3166-1 alpha-2 country code (e.g., "GB")
     * @param regionCode ISO 3166-2 subdivision code (e.g., "GB-SCT" for Scotland)
     * @param year Year (e.g., 2026)
     * @param options Optional filters (type, language, timezone)
     */
    getRegionalHolidays(country: string, regionCode: string, year: number, options?: Omit<HolidayOptions, "regions">): Promise<HolidaysResponse>;
    /**
     * Check if a specific date is a holiday
     * @param country ISO 3166-1 alpha-2 country code (e.g., "US")
     * @param date Date in format YYYY-MM-DD (e.g., "2026-12-25")
     * @param options Optional filters (regions)
     */
    checkHoliday(country: string, date: string, options?: Pick<HolidayOptions, "regions">): Promise<CheckHolidayResponse>;
}
