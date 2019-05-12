package theysaidsogo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kennykarnama/theysaidsogo/theysaidsogo/config"
	"github.com/kennykarnama/theysaidsogo/theysaidsogo/model"
)

const (
	QuoteOfTheDay = "qod"
)

func GetQuoteOfTheDay() (*model.QuoteSearchResult, error) {
	url := config.ApiEndpoint + "/" + QuoteOfTheDay + config.JsonResult
	resp, err := http.Get(url)
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
