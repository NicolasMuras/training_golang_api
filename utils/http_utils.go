package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// check look for an error in the http request that was made.
func Check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] fetch: %v\n", err)
		os.Exit(1)
	}
}

func ApiRequest(url string, host string) []byte {
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}
