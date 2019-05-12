package model

type QuoteSearchResult struct {
	Status   *Total     `json:"success"`
	Contents *QuoteList `json:"contents"`
}

type Total struct {
	Total int `json:"total"`
}

type QuoteList struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Quote    string   `json:"quote"`
	Author   string   `json:"author"`
	Length   string   `json:"length"`
	Tags     []string `json:"tags"`
	Category string   `json:"category"`
	Title    string   `json:"title"`
	Date     string   `json:"date"`
	Id       string   `json:"id"`
}
