package ratings

import (
	"encoding/xml"
	"net/http"
	"os"

	"github.com/meghashyamc/book-ratings/models"
	log "github.com/sirupsen/logrus"
)

func GetGoodreadsRating(book string) (float32, error) {

	requestMap := map[string]string{reqType: http.MethodGet, url: os.Getenv("GOODREADS_URL") + os.Getenv("GOODREADS_KEY"), paramKey: "q", paramVal: book}
	response, err := makeHttpReq(requestMap)
	if err != nil {
		return 0.0, err
	}
	content, err := readResponseAndCheckStatusCode(response, requestMap)
	if err != nil {
		return 0.0, err
	}
	goodreadsObj := models.GoodreadsResponse{}

	if err = xml.Unmarshal(content, &goodreadsObj); err != nil {
		log.WithFields(log.Fields{
			"err":           err.Error(),
			"response_body": content,
		}).Error("Could not unmarshal XML response received from Goodreads")

		return 0.0, err
	}

	return goodreadsObj.Search.Results.Works[0].AverageRating, nil

}
