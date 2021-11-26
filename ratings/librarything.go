package ratings

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	log "github.com/sirupsen/logrus"
)

const scheme = "https"

func GetLibraryThingRating(book string) (float32, error) {

	res, err := getBookDetailsFromSearch(book, libraryThing)
	if err != nil {
		return 0.0, err
	}
	libraryThingURL := res.FormattedUrl
	response, err := http.Get(scheme + "://" + libraryThingURL)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"url": libraryThingURL,
		}).Error("Got an error when calling Library Thing URL")

		return 0.0, err
	}

	if response.StatusCode != http.StatusOK {
		err := errors.New("Non-OK response received after making HTTP request to Library Thing")
		log.WithFields(log.Fields{
			"status_code": response.StatusCode,
			"status":      response.Status,
			"header":      response.Header,
			"url":         libraryThingURL,
		}).Error()
		return 0.0, err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	var ratingWithBrackets string
	doc.Find(libraryThingResponseClass).Each(func(i int, element *goquery.Selection) {
		if i > 0 {
			return
		}
		ratingWithBrackets = element.Text()
	})

	libraryThingRating, err := getNumberFromStringWithBrackets(ratingWithBrackets)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Could not convert gotten rating in Library Thing page to a number")
		return 0.0, err
	}

	return libraryThingRating, nil

}

//the function below expects a string like so: ( 4.3 )
func getNumberFromStringWithBrackets(s string) (float32, error) {

	s = strings.ReplaceAll(s, " ", "")
	s = strings.Replace(s, "(", "", 1)
	s = strings.Replace(s, ")", "", 1)

	num, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.WithFields(log.Fields{
			"err":           err.Error(),
			"string_passed": s,
		}).Error("Could not convert string with brackets a number")
		return 0.0, err
	}

	return float32(num), nil

}
