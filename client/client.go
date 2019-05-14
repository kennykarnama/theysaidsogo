package client

import (
	"encoding/json"
	"errors"
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

// Get quotes by category
// Parameter : category
// Return *model.QuoteSearchResult
func (cli *Client) GetQuotesByCategory(category string) (*model.QuoteSearchResult, error) {

	url := config.ApiEndpoint + "/" + config.QuoteOfTheDay + config.JsonResult

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	param := req.URL.Query()
	param.Add(config.CategoryQuery, category)

	req.URL.RawQuery = param.Encode()

	resp, err := cli.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Quotes search failed %s\n", resp.Status)
	}

	var result model.QuoteSearchResult

	tmpResult := make(map[string]interface{})

	if err := json.NewDecoder(resp.Body).Decode(&tmpResult); err != nil {
		return nil, err
	}

	if _, ok := tmpResult["failure"]; ok {
		return nil, errors.New(tmpResult["reason"].(string))
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
