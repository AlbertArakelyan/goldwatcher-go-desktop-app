package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.HTTPClient = client
	os.Exit(m.Run())
}

var jsonToReturn = `
	{
		"ts": 1727032022278,
		"tsj": 1727032015131,
		"date": "Sep 22nd 2024, 03:06:55 pm NY",
		"items": [
			{
				"curr": "USD",
				"xauPrice": 2622.3,
				"xagPrice": 31.1725,
				"chgXau": 34.95,
				"chgXag": 0.3925,
				"pcXau": 1.3508,
				"pcXag": 1.2752,
				"xauClose": 2587.35,
				"xagClose": 30.78
			}
		]
	}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
