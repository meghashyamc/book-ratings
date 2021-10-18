package models

import "encoding/xml"

type LibraryThingResponse struct {
	XMLName xml.Name `xml:"response"`

	Ltml LtmlItem `xml:"ltml"`
}

type LtmlItem struct {
	Item ItemWithRating `xml:"item"`
}

type ItemWithRating struct {
	Rating float32 `xml:"rating"`
}
