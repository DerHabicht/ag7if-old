package actions

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

type forecastZone struct {
	Features []struct {
		Properties struct {
			ID string `json:"id"`
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

	return c.Render(http.StatusOK,
		r.Func("application/json", func(w io.Writer, d render.Data) error {
			_, err := w.Write(body)
			return err
		}),
	)
}
