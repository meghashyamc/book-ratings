package ratings

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	search "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func getBookDetailsFromSearch(book, site string) (*search.Result, error) {

	searchService, err := search.NewService(context.Background(), option.WithAPIKey(os.Getenv("SEARCH_API_KEY")))
	if err != nil {
		log.WithFields(log.Fields{
			"err":  err.Error(),
			"book": book,
			"site": site,
		}).Error("Could not form new search service)")
		return nil, err
	}

	searchID := os.Getenv(searchIDEnvVariableNames[site])
	if searchID == "" {
		log.WithFields(log.Fields{
			"book": book,
			"site": site,
		}).Error("Could not get search ID for site)")
		return nil, err
	}

	res, err := getFirstSearchResult(book, searchID, searchService)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func getFirstSearchResult(book, searchID string, searchService *search.Service) (*search.Result, error) {

	searchCall := searchService.Cse.List()

	searchCall.Q(book)
	searchCall.Cx(searchID)
	searchResults, err := searchCall.Do()

	if err != nil {
		log.WithFields(log.Fields{
			"err":      err.Error(),
			"book":     book,
			"searchID": searchID,
		}).Error("Could not get search results after searching for book in site)")
		return nil, err

	}

	if len(searchResults.Items) == 0 {
		log.WithFields(log.Fields{
			"err":      err.Error(),
			"book":     book,
			"searchID": searchID,
		}).Error("Got zero search results after searching for book in site)")
		return nil, err
	}
	return searchResults.Items[0], nil

}
