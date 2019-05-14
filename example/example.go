package main

import (
	"fmt"
	"log"

	"github.com/kennykarnama/theysaidsogo/client"
)

func main() {

	getQuoteOfTheDay()
	getQuoteByCategory()

}

func getQuoteOfTheDay() {
	quoteClient := client.NewClient()
	quotes, err := quoteClient.GetQuoteOfTheDay()
	if err != nil {
		log.Fatal(err)
	}
	for _, quote := range quotes.Contents.Quotes {
		fmt.Printf("Quote of the day: %s\n", quote.Quote)
	}
}

func getQuoteByCategory() {
	quoteClient := client.NewClient()
	quotes, err := quoteClient.GetQuotesByCategory("test")
	if err != nil {
		log.Fatal(err)
	}
	for _, quote := range quotes.Contents.Quotes {
		fmt.Printf("Quote of the day: %s\n", quote.Quote)
	}
}
