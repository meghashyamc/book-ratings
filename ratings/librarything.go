package ratings

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/meghashyamc/book-ratings/models"
	log "github.com/sirupsen/logrus"
)

func GetLibraryThingRating(book string) (float32, error) {

	requestMap := map[string]string{reqType: http.MethodGet, url: os.Getenv("BING_URL"), paramKey: "q", paramVal: book + " book librarything", headerKey: os.Getenv("BING_TOKEN_FIELD"), headerVal: os.Getenv("BING_TOKEN")}
	// searching for the book on Bing
	response, err := makeHttpReq(requestMap)
	if err != nil {
		return 0.0, err
	}

	content, err := readResponseAndCheckStatusCode(response, requestMap)
	if err != nil {
		return 0.0, err
	}
	bingObj := models.BingAnswer{}

	err = json.Unmarshal(content, &bingObj)
	if err != nil {
		log.WithFields(log.Fields{
			"err":           err.Error(),
			"response_body": content,
		}).Error("Could not unmarshal response received from Bing(for Library Thing)")

		return 0.0, err
	}

	// first URL result in Bing
	libraryThingUrlBing := bingObj.WebPages.Value[0].URL

	// getting Library Thing number for the book

	slashIndex := strings.LastIndex(libraryThingUrlBing, "/")

	if slashIndex == -1 {
		err := errors.New("Could not get Library Thing book number from Bing search")
		log.WithFields(log.Fields{
			"err":           err.Error(),
			"response_body": content,
		}).Error("Could not fetch Library Thing book number")

		return 0.0, err
	}

	bookNum := libraryThingUrlBing[slashIndex+1:]

	// contact library thing with the book number
	libraryThingUrl := os.Getenv("LIBRARY_THING_URL") + os.Getenv("LIBRARY_THING_KEY") + "&id=" + bookNum
	response, err = http.Get(libraryThingUrl)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"url": libraryThingUrl,
		}).Error("Got an error when calling Library Thing URL")

		return 0.0, err
	}

	content, err = readResponseAndCheckStatusCode(response, nil)
	libObj := models.LibraryThingResponse{}

	if err = xml.Unmarshal(content, &libObj); err != nil {
		log.WithFields(log.Fields{
			"err":           err.Error(),
			"response_body": content,
		}).Error("Could not unmarshal XML response received from Library Thing")

		return 0.0, err
	}

	return libObj.Ltml.Item.Rating / 2, nil
}
