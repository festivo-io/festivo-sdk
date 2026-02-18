package festivo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client is the Festivo API client for accessing public holiday data.
type Client struct {
	BaseURL string
	APIKey  string
	HTTP    *http.Client
}

// HolidayOptions contains optional filters for holiday queries.
type HolidayOptions struct {
	Regions  string
	Type     string
	Language string
	Timezone string
}

// Holiday represents a public holiday.
type Holiday struct {
	Date        string                   `json:"date"`
	Name        string                   `json:"name"`
	NameLocal   string                   `json:"name_local,omitempty"`
	Type        string                   `json:"type"`
	Observed    string                   `json:"observed"`
	Public      bool                     `json:"public"`
	Country     string                   `json:"country"`
	Subdivisions []string                `json:"subdivisions"`
	Regions     []map[string]interface{} `json:"regions,omitempty"`
}

// HolidaysResponse represents the response from the holidays list endpoint.
type HolidaysResponse struct {
	Holidays []Holiday `json:"holidays"`
	Total    int       `json:"total"`
}

// CheckHolidayResponse represents the response from the check holiday endpoint.
type CheckHolidayResponse struct {
	IsHoliday bool     `json:"is_holiday"`
	Holiday   *Holiday `json:"holiday,omitempty"`
}

// NewClient creates a new Festivo API client.
func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: "https://api.getfestivo.com",
		APIKey:  apiKey,
		HTTP:    http.DefaultClient,
	}
}

func (c *Client) request(path string, params url.Values) ([]byte, error) {
	urlStr := fmt.Sprintf("%s%s?%s", c.BaseURL, path, params.Encode())
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d %s", resp.StatusCode, resp.Status)
	}

	var result []byte
	_, err = resp.Body.Read(result)
	if err != nil {
		return nil, err
	}

	// Read full body
	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(tmp)
		if n > 0 {
			buf = append(buf, tmp[:n]...)
		}
		if err != nil {
			break
		}
	}

	return buf, nil
}

// GetHolidays gets all holidays for a country and year.
//
// country: ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "IT")
// year: Year (e.g., 2026)
// opts: Optional filters (regions, type, language, timezone)
func (c *Client) GetHolidays(country string, year int, opts *HolidayOptions) (*HolidaysResponse, error) {
	params := url.Values{}
	params.Set("country", country)
	params.Set("year", fmt.Sprintf("%d", year))

	if opts != nil {
		if opts.Regions != "" {
			params.Set("regions", opts.Regions)
		}
		if opts.Type != "" {
			params.Set("type", opts.Type)
		}
		if opts.Language != "" {
			params.Set("language", opts.Language)
		}
		if opts.Timezone != "" {
			params.Set("timezone", opts.Timezone)
		}
	}

	data, err := c.request("/v3/public-holidays/list", params)
	if err != nil {
		return nil, err
	}

	var resp HolidaysResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetCityHolidays gets holidays for a specific city (Pro plan).
//
// country: ISO 3166-1 alpha-2 country code (e.g., "IT")
// cityCode: City code in format {COUNTRY}-{CITY} (e.g., "IT-MILAN")
// year: Year (e.g., 2026)
// opts: Optional filters (type, language, timezone)
func (c *Client) GetCityHolidays(country, cityCode string, year int, opts *HolidayOptions) (*HolidaysResponse, error) {
	if opts == nil {
		opts = &HolidayOptions{}
	}
	opts.Regions = cityCode
	return c.GetHolidays(country, year, opts)
}

// GetRegionalHolidays gets holidays for a specific region using ISO 3166-2 codes (Builder plan).
//
// country: ISO 3166-1 alpha-2 country code (e.g., "GB")
// regionCode: ISO 3166-2 subdivision code (e.g., "GB-SCT" for Scotland)
// year: Year (e.g., 2026)
// opts: Optional filters (type, language, timezone)
func (c *Client) GetRegionalHolidays(country, regionCode string, year int, opts *HolidayOptions) (*HolidaysResponse, error) {
	if opts == nil {
		opts = &HolidayOptions{}
	}
	opts.Regions = regionCode
	return c.GetHolidays(country, year, opts)
}

// CheckHoliday checks if a specific date is a holiday.
//
// country: ISO 3166-1 alpha-2 country code (e.g., "US")
// date: Date in format YYYY-MM-DD (e.g., "2026-12-25")
// regions: Optional comma-separated region/city codes
func (c *Client) CheckHoliday(country, date string, regions string) (*CheckHolidayResponse, error) {
	params := url.Values{}
	params.Set("country", country)
	params.Set("date", date)
	if regions != "" {
		params.Set("regions", regions)
	}

	data, err := c.request("/v3/public-holidays/check", params)
	if err != nil {
		return nil, err
	}

	var resp CheckHolidayResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

