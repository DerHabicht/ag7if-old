package actions

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

func AlertsGet(c buffalo.Context) error {
	url := "https://api.weather.gov/alerts/active?" +
		"point=" + c.Param("lat") + "," + c.Param("long")

	resp, err := http.Get(url)
	// TODO: Handle request errors intelligently
	if err != nil {
		c.Logger().Errorf("Error on GET: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return c.Render(http.StatusOK,
		r.Func("application/json", func(w io.Writer, d render.Data) error {
			_, err := w.Write(body)
			return err
		}),
	)
}
