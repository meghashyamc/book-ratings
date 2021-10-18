package ratings

import (
	"errors"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func makeHttpReq(requestMap map[string]string) (*http.Response, error) {

	request, err := http.NewRequest(requestMap[reqType], requestMap[url], nil)
	if err != nil {
		log.WithFields(log.Fields{
			"err":             err.Error(),
			"request_details": requestMap,
		}).Error("Error in creating new request")

		return nil, err
	}

	q := request.URL.Query()
	q.Add(requestMap[paramKey], requestMap[paramVal])

	request.URL.RawQuery = q.Encode()

	if requestMap[headerKey] != "" {
		request.Header.Add(requestMap[headerKey], requestMap[headerVal])
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.WithFields(log.Fields{
			"err":             err.Error(),
			"request_details": requestMap,
		}).Error("Error received when calling URL")
		return nil, err
	}

	return response, err

}

func readResponseAndCheckStatusCode(response *http.Response, requestMap map[string]string) ([]byte, error) {
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err":             err.Error(),
			"request_details": requestMap,
		}).Error("Could not read response body after making HTTP request")

		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"err":             content,
			"status_code":     response.StatusCode,
			"request_details": requestMap,
		}).Error("Non-OK response received after making HTTP request")
		return nil, errors.New("Non-OK response received after making HTTP request")
	}

	return content, nil
}
