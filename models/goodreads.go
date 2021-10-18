package models

import (
	"encoding/xml"
)

type GoodreadsResponse struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`

	Search SearchResults `xml:"search"`
}

type SearchResults struct {
	Results ResultsWithWorks `xml:"results"`
}

type ResultsWithWorks struct {
	Works []Work `xml:"work"`
}
type Work struct {
	AverageRating float32 `xml:"average_rating"`
}
