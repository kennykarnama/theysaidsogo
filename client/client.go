package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kennykarnama/theysaidsogo/config"
	"github.com/kennykarnama/theysaidsogo/model"
)

type Client struct {
	HttpClient *http.Client
}

// Construct the new Client
// Return pointer to Client
func NewClient() *Client {
	cli := Client{
		HttpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	return &cli
}

// Get quote of the day
// return *QUoteSearchResult, error
func (cli *Client) GetQuoteOfTheDay() (*model.QuoteSearchResult, error) {
	url := config.ApiEndpoint + "/" + config.QuoteOfTheDay + config.JsonResult
	resp, err := cli.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Quotes search failed %s\n", resp.Status)
	}

	var result model.QuoteSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
