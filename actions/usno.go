package actions

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

func USNOGet(c buffalo.Context) error {
	url := "http://api.usno.navy.mil/rstt/oneday?" +
		"date=today&" +
		"coords=" +
		c.Param("lat") + "," +
		c.Param("long") + "&"

	// Since the US Military has a bad habit of signing their own SSL
	// certificates, we have to bypass checking the CA.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
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
		r.Func("application/json", func(w io.Writer, d render.Data) error {
			_, err := w.Write(body)
			return err
		}),
	)
}
