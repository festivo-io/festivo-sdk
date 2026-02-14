package festivo

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Client struct {
    BaseURL string
    APIKey  string
    HTTP    *http.Client
}

func NewClient(apiKey string) *Client {
    return &Client{BaseURL: "https://api.getfestivo.com", APIKey: apiKey, HTTP: http.DefaultClient}
}

func (c *Client) GetInvoice(id string) (map[string]interface{}, error) {
    req, _ := http.NewRequest("GET", fmt.Sprintf("%s/invoices/%s", c.BaseURL, id), nil)
    if c.APIKey != "" {
        req.Header.Set("Authorization", "Bearer "+c.APIKey)
    }
    req.Header.Set("Accept", "application/json")
    resp, err := c.HTTP.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    var out map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
        return nil, err
    }
    return out, nil
}
