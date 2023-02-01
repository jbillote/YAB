package util

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/jbillote/YAB/util/logger"
)

var log = logger.GetLogger("Util")

/*
 * Function: URLValid
 * Test  whether or not a string is a valid URL
 *
 * Params:
 * s: String to test
 *
 * Returns:
 * True if string is a valid URL, else false
 */
func URLValid(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

/*
 * Function: URLAvailable
 * Test whether or not a URL is accessible
 *
 * Params:
 * url: URL to test
 *
 * Returns:
 * True if URL is accessible, else false
 */
func URLAvailable(url string) bool  {
	log.Info(fmt.Sprintf("Checking availability of URL, url=%s", url))

	r, err := http.Get(url)
	available := err == nil && r!= nil && r.StatusCode == 200

	if available {
		log.Info(fmt.Sprintf("URL is available, url=%s", url))
	} else {
		if r != nil {
			log.Error(fmt.Sprintf("URL is not available, url=%s, err=%o, status=%d", url, err, r.StatusCode))
		} else {
			log.Error(fmt.Sprintf("URL is not available, url=%s, err=%o, status=nil", url, err))
		}
	}

	return available
}
