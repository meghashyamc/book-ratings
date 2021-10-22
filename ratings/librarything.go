package ratings

import (
	"encoding/xml"
	"net/http"
	"os"
	"strings"

	"github.com/meghashyamc/book-ratings/models"
	log "github.com/sirupsen/logrus"
)

func GetLibraryThingRating(book string) (float32, error) {

	res, err := getBookDetailsFromSearch(book, libraryThing)
	if err != nil {
		return 0.0, err
	}

	resURLParts := strings.Split(res.FormattedUrl, "/")
	libraryThingIndex := resURLParts[len(resURLParts)-1]

	// contact library thing with the book number
	libraryThingURL := os.Getenv("LIBRARY_THING_URL") + os.Getenv("LIBRARY_THING_KEY") + "&id=" + libraryThingIndex
	response, err := http.Get(libraryThingURL)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"url": libraryThingURL,
		}).Error("Got an error when calling Library Thing URL")

		return 0.0, err
	}

	content, err := readResponseAndCheckStatusCode(response, nil)
	libraryThingResponse := models.LibraryThingResponse{}

	if err = xml.Unmarshal(content, &libraryThingResponse); err != nil {
		log.WithFields(log.Fields{
			"err":           err.Error(),
			"response_body": content,
		}).Error("Could not unmarshal XML response received from Library Thing")

		return 0.0, err
	}

	return libraryThingResponse.Ltml.Item.Rating / 2, nil
}
