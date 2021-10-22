package ratings

import (
	"encoding/json"
	"strconv"

	"github.com/meghashyamc/book-ratings/models"
	log "github.com/sirupsen/logrus"
)

func GetGoodreadsRating(book string) (float32, error) {

	res, err := getBookDetailsFromSearch(book, goodreads)
	if err != nil {
		return 0.0, err
	}
	goodreadsPageMap := models.GoodreadsPageMap{}
	if err := json.Unmarshal(res.Pagemap, &goodreadsPageMap); err != nil {
		log.WithFields(log.Fields{
			"err":  err.Error(),
			"book": book,
			"site": goodreads,
		}).Error("Got an error when unmarshaling search result to read Goodreads rating")
		return 0.0, err
	}

	goodreadsRating, err := strconv.ParseFloat(goodreadsPageMap.Review[0].RatingStars, 32)
	if err != nil {
		log.WithFields(log.Fields{
			"err":  err.Error(),
			"book": book,
			"site": goodreads,
		}).Error("Could not convert Goodreads rating to a number")
		return 0.0, err

	}
	return float32(goodreadsRating), nil

}
