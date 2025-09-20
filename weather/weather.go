package weather

import (
	"encoding/json"
	"fmt"
	"net/url"
	"net/http"
)

// CurrentWeather struct, the part of the Open-metro response we want
type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	WindSpeed float64 `json:"windspeed"`
	WindDirection float64 `json:"winddirection"`
	WeatherCode int `json:"weathercode"`
	Time string `json"time"`
}


//internal struct to decode the api response
type openMetroResponse struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CurrentWeather CurrentWeather `json:"curent_weather"`
}

// WeatherClient holds the config for calling the weather API
// The pointer field for http.Client: *http.Client for practise
type WeatherClient struct {
	BaseURL string
	HTTPClient *http.Client
}


// Newclient returns a pointer to WeatherClient
// Returning *Weatherclient is idiomatic when the client holds the state or large fields
func NewClient() *WeatherClient {
	return &WeatherClient{
		BaseURL: "https://api.open-metro.com/v1/forecast",
		HTTPClient: &http.Client{},
	}
}

// GetCurrentWeather uses a pointer reciever (c *WeatherClient)
// It fetches current weather for given latitude/longitdude and returns to CurrentWeather
func (c *WeatherClient) GetCurrentWeather(lat, lon float64) (*CurrentWeather, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("latidude", fmt.Sprintf("%f", lat))
	q.Set("longitude", fmt.Sprintf("%f", lon))
	q.Set("current_weather", "true")
	u.RawQuery = q.Encode()

	res, err := c.HTTPClient.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Weather API returned status: %s", res.StatusCode)
	}

	var om openMetroRes
	if err := json.NewDecoder(res.Body)Decode(&om); err != nil {
		return nil, err
	}

	// return pointer to CurrentWeather (safe - go will allocate on the heap if needed)
	return &om.CurrentWeather, nil
}