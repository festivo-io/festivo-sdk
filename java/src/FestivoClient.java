package com.festivo;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Map;
import java.util.HashMap;
import org.json.JSONArray;
import org.json.JSONObject;

public class FestivoClient {
    private final String apiKey;
    private final String baseUrl;

    public FestivoClient(String apiKey) {
        this(apiKey, "https://api.getfestivo.com");
    }

    public FestivoClient(String apiKey, String baseUrl) {
        this.apiKey = apiKey;
        this.baseUrl = baseUrl;
    }

    public HolidaysResponse getHolidays(String country, int year, Map<String, String> options) throws Exception {
        Map<String, String> params = new HashMap<>();
        params.put("country", country);
        params.put("year", String.valueOf(year));
        if (options != null) params.putAll(options);
        JSONObject json = request("/v3/public-holidays/list", params);
        return HolidaysResponse.fromJson(json);
    }

    public HolidaysResponse getCityHolidays(String country, String cityCode, int year, Map<String, String> options) throws Exception {
        Map<String, String> opts = options == null ? new HashMap<>() : new HashMap<>(options);
        opts.put("regions", cityCode);
        return getHolidays(country, year, opts);
    }

    public HolidaysResponse getRegionalHolidays(String country, String regionCode, int year, Map<String, String> options) throws Exception {
        Map<String, String> opts = options == null ? new HashMap<>() : new HashMap<>(options);
        opts.put("regions", regionCode);
        return getHolidays(country, year, opts);
    }

    public JSONObject checkHoliday(String country, String date, String regions) throws Exception {
        Map<String, String> params = new HashMap<>();
        params.put("country", country);
        params.put("date", date);
        if (regions != null) params.put("regions", regions);
        return request("/v3/public-holidays/check", params);
    }

    private JSONObject request(String path, Map<String, String> params) throws Exception {
        StringBuilder urlBuilder = new StringBuilder(baseUrl + path + "?");
        for (Map.Entry<String, String> entry : params.entrySet()) {
            urlBuilder.append(entry.getKey()).append("=").append(entry.getValue()).append("&");
        }
        URL url = new URL(urlBuilder.toString());
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        conn.setRequestMethod("GET");
        conn.setRequestProperty("Accept", "application/json");
        conn.setRequestProperty("Authorization", "Bearer " + apiKey);
        int status = conn.getResponseCode();
        BufferedReader reader = new BufferedReader(new InputStreamReader(conn.getInputStream()));
        StringBuilder response = new StringBuilder();
        String line;
        while ((line = reader.readLine()) != null) {
            response.append(line);
        }
        reader.close();
        if (status != 200) throw new Exception("Festivo API error: " + status);
        return new JSONObject(response.toString());
    }

    // Minimal model classes
    public static class HolidaysResponse {
        public final JSONArray holidays;
        public HolidaysResponse(JSONArray holidays) {
            this.holidays = holidays;
        }
        public static HolidaysResponse fromJson(JSONObject json) {
            return new HolidaysResponse(json.getJSONArray("holidays"));
        }
    }
}
