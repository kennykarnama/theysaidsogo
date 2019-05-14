package main

import (
	"fmt"
	"log"

	"github.com/kennykarnama/theysaidsogo/theysaidsogo/client"
)

func main() {
	saidClient := client.NewClient()
	quotes, err := saidClient.GetQuoteOfTheDay()
	if err != nil {
		log.Fatal(err)
	}
	for _, quote := range quotes.Contents.Quotes {
		fmt.Printf("Quote of the day: %s\n", quote.Quote)
	}
}
