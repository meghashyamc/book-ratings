package models

type GoodreadsPageMap struct {
	Review []ReviewDetails `json:"review"`
}

type ReviewDetails struct {
	RatingStars string `json:"ratingstars"`
}
