package client

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	okResponse = `{
		"success": {
			"total": 1
		},
		"contents": {
			"quotes": [
				{
					"quote": "Sometimes you climb out of bed in the morning and you think, I'm not going to make it, but you laugh inside — remembering all the times you've felt that way.",
					"author": "Charles Bukowski",
					"length": "164",
					"tags": [
						"depression",
						"inspire",
						"tso-life"
					],
					"category": "inspire",
					"title": "Inspiring Quote of the day",
					"date": "2019-05-14",
					"id": null
				}
			],
			"copyright": "2017-19 theysaidso.com"
		}
	}`
	notFoundResponse = `{
		"failure": 1,
		"total": 0,
		"reason": "QOD Category not supported"
	}`
)

func TestGetQuoteOfTheDay(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(okResponse))
	})
	httpClient, teardown := testingHttpClient(h)
	defer teardown()

	cli := NewClient()
	cli.HttpClient = httpClient

	quotes, err := cli.GetQuoteOfTheDay()
	assert.Nil(t, err)
	assert.Equal(t, 1, quotes.Status.Total)

}

func TestGetQuoteOfTheDayWhenCategoryNotExist(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(notFoundResponse))
	})
	httpClient, teardown := testingHttpClient(h)
	defer teardown()

	cli := NewClient()
	cli.HttpClient = httpClient

	_, err := cli.GetQuotesByCategory("test")
	assert.NotNil(t, err)
}

func testingHttpClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewServer(handler)
	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}
	return cli, s.Close
}
