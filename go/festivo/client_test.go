package festivo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test-key")

	if client.BaseURL != "https://api.getfestivo.com" {
		t.Errorf("Expected base URL to be https://api.getfestivo.com, got %s", client.BaseURL)
	}

	if client.APIKey != "test-key" {
		t.Errorf("Expected API key to be test-key, got %s", client.APIKey)
	}

	if client.HTTP == nil {
		t.Error("Expected HTTP client to be initialized")
	}
}

func TestGetHolidays(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("Expected Authorization header with bearer token")
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header to be application/json")
		}

		// Check query parameters
		query := r.URL.Query()
		if query.Get("country") != "US" {
			t.Errorf("Expected country to be US, got %s", query.Get("country"))
		}
		if query.Get("year") != "2026" {
			t.Errorf("Expected year to be 2026, got %s", query.Get("year"))
		}

		// Return mock response
		resp := HolidaysResponse{
			Holidays: []Holiday{
				{
					Date:     "2026-01-01",
					Name:     "New Year's Day",
					Type:     "public",
					Observed: "2026-01-01",
					Public:   true,
					Country:  "US",
				},
			},
			Total: 1,
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client with mock server
	client := NewClient("test-key")
	client.BaseURL = server.URL

	// Test
	result, err := client.GetHolidays("US", 2026, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Total != 1 {
		t.Errorf("Expected total to be 1, got %d", result.Total)
	}

	if len(result.Holidays) != 1 {
		t.Fatalf("Expected 1 holiday, got %d", len(result.Holidays))
	}

	if result.Holidays[0].Name != "New Year's Day" {
		t.Errorf("Expected holiday name to be New Year's Day, got %s", result.Holidays[0].Name)
	}
}

func TestGetHolidaysWithOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		
		if query.Get("regions") != "GB-SCT" {
			t.Errorf("Expected regions to be GB-SCT, got %s", query.Get("regions"))
		}
		if query.Get("type") != "public" {
			t.Errorf("Expected type to be public, got %s", query.Get("type"))
		}
		if query.Get("language") != "en" {
			t.Errorf("Expected language to be en, got %s", query.Get("language"))
		}

		resp := HolidaysResponse{Holidays: []Holiday{}, Total: 0}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("test-key")
	client.BaseURL = server.URL

	opts := &HolidayOptions{
		Regions:  "GB-SCT",
		Type:     "public",
		Language: "en",
	}

	_, err := client.GetHolidays("GB", 2026, opts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetCityHolidays(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		
		if query.Get("country") != "IT" {
			t.Errorf("Expected country to be IT, got %s", query.Get("country"))
		}
		if query.Get("regions") != "IT-MILAN" {
			t.Errorf("Expected regions to be IT-MILAN, got %s", query.Get("regions"))
		}
		if query.Get("year") != "2026" {
			t.Errorf("Expected year to be 2026, got %s", query.Get("year"))
		}

		resp := HolidaysResponse{Holidays: []Holiday{}, Total: 0}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("test-key")
	client.BaseURL = server.URL

	_, err := client.GetCityHolidays("IT", "IT-MILAN", 2026, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetRegionalHolidays(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		
		if query.Get("country") != "GB" {
			t.Errorf("Expected country to be GB, got %s", query.Get("country"))
		}
		if query.Get("regions") != "GB-SCT" {
			t.Errorf("Expected regions to be GB-SCT, got %s", query.Get("regions"))
		}

		resp := HolidaysResponse{Holidays: []Holiday{}, Total: 0}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("test-key")
	client.BaseURL = server.URL

	_, err := client.GetRegionalHolidays("GB", "GB-SCT", 2026, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCheckHoliday(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		
		if query.Get("country") != "US" {
			t.Errorf("Expected country to be US, got %s", query.Get("country"))
		}
		if query.Get("date") != "2026-12-25" {
			t.Errorf("Expected date to be 2026-12-25, got %s", query.Get("date"))
		}

		if r.URL.Path != "/v3/public-holidays/check" {
			t.Errorf("Expected path to be /v3/public-holidays/check, got %s", r.URL.Path)
		}

		resp := CheckHolidayResponse{
			IsHoliday: true,
			Holiday: &Holiday{
				Date:     "2026-12-25",
				Name:     "Christmas Day",
				Type:     "public",
				Observed: "2026-12-25",
				Public:   true,
				Country:  "US",
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("test-key")
	client.BaseURL = server.URL

	result, err := client.CheckHoliday("US", "2026-12-25", "")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !result.IsHoliday {
		t.Error("Expected is_holiday to be true")
	}

	if result.Holiday.Name != "Christmas Day" {
		t.Errorf("Expected holiday name to be Christmas Day, got %s", result.Holiday.Name)
	}
}

func TestCheckHolidayWithRegions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		
		if query.Get("regions") != "IT-MILAN" {
			t.Errorf("Expected regions to be IT-MILAN, got %s", query.Get("regions"))
		}

		resp := CheckHolidayResponse{IsHoliday: true}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("test-key")
	client.BaseURL = server.URL

	_, err := client.CheckHoliday("IT", "2026-12-07", "IT-MILAN")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Unauthorized"}`))
	}))
	defer server.Close()

	client := NewClient("test-key")
	client.BaseURL = server.URL

	_, err := client.GetHolidays("US", 2026, nil)
	if err == nil {
		t.Error("Expected error for 401 response, got nil")
	}
}

func TestClientWithoutAPIKey(t *testing.T) {
	os.Unsetenv("FESTIVO_API_KEY")
	client := NewClient("")

	if client.APIKey != "" {
		t.Errorf("Expected empty API key, got %s", client.APIKey)
	}
}
