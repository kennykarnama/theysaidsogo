package main

import (
	"fmt"
	"log"

	"github.com/kennykarnama/theysaidsogo/theysaidsogo"
)

func main() {
	quotes, err := theysaidsogo.GetQuoteOfTheDay()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There are %d quotes\n", quotes.Status.Total)
	for _, quote := range quotes.Contents.Quotes {
		fmt.Printf("Quote: %s\n", quote.Quote)
	}
}
