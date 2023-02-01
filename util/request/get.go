package request

import (
	"github.com/jbillote/YAB/config"
	"github.com/jbillote/YAB/util/logger"

	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var log = logger.GetLogger("HTTP Request")

func GetRequest(uri string) (io.ReadCloser, error) {
	c := config.GetConfig()
	log.Info(fmt.Sprintf("Making GET request to %s", uri))

	r, err := http.Get(c.APIHost + uri)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to make GET request to %s", uri))
		log.Error(err)
		return nil, err
	}

	if r.StatusCode != 200 {
		log.Error(fmt.Sprintf("Request completed with non-200 status code %d", r.StatusCode))

		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Unable to read response body")
			log.Error(err)
			return nil, err
		}

		return nil, errors.New(string(response))
	}

	log.Info(fmt.Sprintf("Successfully made GET request to %s", uri))

	return r.Body, nil
}
