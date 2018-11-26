package actions

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

func ADDSGet(c buffalo.Context) error {
	url := "https://www.aviationweather.gov/adds/dataserver_current/" +
		"httpparam?dataSource=metars&" +
		"requestType=retrieve&" +
		"format=xml&" +
		"mostRecentForEachStation=constraint&" +
		"radialDistance=" + c.Param("dist") + ";" +
		c.Param("long") + "," +
		c.Param("lat") + "&" +
		"hoursBeforeNow=" + c.Param("hoursBeforeNow")

	resp, err := http.Get(url)
	// TODO: Handle request errors intelligently
	if err != nil {
		c.Logger().Errorf("Error on GET: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// TODO: Handle read errors intelligently
	if err != nil {
		c.Logger().Errorf("Error on read of body: %s", err)
	}

	return c.Render(http.StatusOK,
		r.Func("application/xml", func(w io.Writer, d render.Data) error {
			_, err := w.Write(body)
			return err
		}),
	)
}
