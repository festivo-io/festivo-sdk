<?php
namespace Festivo;

use GuzzleHttp\Client as Guzzle;

/**
 * Festivo API client for accessing public holiday data.
 */
class FestivoClient {
    private $baseUrl;
    private $apiKey;
    private $http;

    /**
     * Initialize the Festivo API client.
     *
     * @param string|null $apiKey Your Festivo API key (or set FESTIVO_API_KEY env var)
     * @param string $baseUrl API base URL (default: https://api.getfestivo.com)
     */
    public function __construct($apiKey = null, $baseUrl = 'https://api.getfestivo.com') {
        $this->apiKey = $apiKey ?: getenv('FESTIVO_API_KEY');
        $this->baseUrl = $baseUrl;
        $this->http = new Guzzle(['base_uri' => $this->baseUrl]);
    }

    private function headers() {
        $h = ['Accept' => 'application/json'];
        if ($this->apiKey) {
            $h['Authorization'] = 'Bearer ' . $this->apiKey;
        }
        return $h;
    }

    private function request(string $path, array $params = []) {
        $res = $this->http->request('GET', $path, [
            'headers' => $this->headers(),
            'query' => array_filter($params, function($value) {
                return $value !== null;
            })
        ]);
        return json_decode($res->getBody()->getContents(), true);
    }

    /**
     * Get all holidays for a country and year.
     *
     * @param string $country ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "IT")
     * @param int $year Year (e.g., 2026)
     * @param array $options Optional filters: regions, type, language, timezone
     * @return array Array with 'holidays' list and 'total' count
     */
    public function getHolidays(string $country, int $year, array $options = []) {
        $params = array_merge(['country' => $country, 'year' => $year], $options);
        return $this->request('/v3/public-holidays/list', $params);
    }

    /**
     * Get holidays for a specific city (Pro plan).
     *
     * @param string $country ISO 3166-1 alpha-2 country code (e.g., "IT")
     * @param string $cityCode City code in format {COUNTRY}-{CITY} (e.g., "IT-MILAN")
     * @param int $year Year (e.g., 2026)
     * @param array $options Optional filters: type, language, timezone
     * @return array Array with 'holidays' list and 'total' count
     */
    public function getCityHolidays(string $country, string $cityCode, int $year, array $options = []) {
        return $this->getHolidays($country, $year, array_merge(['regions' => $cityCode], $options));
    }

    /**
     * Get holidays for a specific region using ISO 3166-2 codes (Builder plan).
     *
     * @param string $country ISO 3166-1 alpha-2 country code (e.g., "GB")
     * @param string $regionCode ISO 3166-2 subdivision code (e.g., "GB-SCT" for Scotland)
     * @param int $year Year (e.g., 2026)
     * @param array $options Optional filters: type, language, timezone
     * @return array Array with 'holidays' list and 'total' count
     */
    public function getRegionalHolidays(string $country, string $regionCode, int $year, array $options = []) {
        return $this->getHolidays($country, $year, array_merge(['regions' => $regionCode], $options));
    }

    /**
     * Check if a specific date is a holiday.
     *
     * @param string $country ISO 3166-1 alpha-2 country code (e.g., "US")
     * @param string $date Date in format YYYY-MM-DD (e.g., "2026-12-25")
     * @param string|null $regions Optional comma-separated region/city codes
     * @return array Array with 'is_holiday' bool and optional 'holiday' array
     */
    public function checkHoliday(string $country, string $date, $regions = null) {
        $params = ['country' => $country, 'date' => $date];
        if ($regions !== null) {
            $params['regions'] = $regions;
        }
        return $this->request('/v3/public-holidays/check', $params);
    }
}

