package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

type forecastZone struct {
	Features []struct {
		Properties struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			State string `json:"state"`
		} `json:"properties"`
	} `json:"features"`
}

func ForecastGet(c buffalo.Context) error {
	baseURL := "https://api.weather.gov/zones"

	// Look up the forecast zone
	lookupURL := baseURL + "?type=forecast&" +
		"point=" + c.Param("lat") + "," + c.Param("long")
	resp, err := http.Get(lookupURL)
	// TODO: Handle request errors intelligently
	if err != nil {
		c.Logger().Errorf("Error on GET Lookup: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// TODO: Handle read errors intelligently
	if err != nil {
		c.Logger().Errorf("Error on read of body: %s", err)
	}
	var lr forecastZone
	err = json.Unmarshal(body, &lr)
	if err != nil {
		c.Logger().Errorf("Error unmarshalling body: %s", err)
	}

	// Get the forecast for this zone
	forecastURL := baseURL + "/forecast/" +
		lr.Features[0].Properties.ID + "/forecast"
	resp, err = http.Get(forecastURL)
	// TODO: Handle request errors intelligently
	if err != nil {
		c.Logger().Errorf("Error on GET Lookup: %s", err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	// TODO: Handle read errors intelligently
	if err != nil {
		c.Logger().Errorf("Error on read of body: %s", err)
	}
	var forecast map[string]interface{}
	err = json.Unmarshal(body, &forecast)

	// Tack the zone information onto the end of the response object
	forecast["zone"] = lr.Features[0].Properties

	return c.Render(http.StatusOK, r.JSON(forecast))
}
