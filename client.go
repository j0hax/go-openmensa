// Package openmensa provides an API to interface with OpenMensa.org.
package openmensa

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

// API endpoint
const endpoint = "https://openmensa.org/api/v2"

// API User Endpoint
const defaultUserAgent = "go-openmensa/0.3"

// The client to use for HTTP requests
var c = http.Client{Timeout: time.Second * 10}

// Function Get is a wrapper for http.Get(),
// using the predifined endpoint and custom headers.
func Get(elem ...string) ([]byte, error) {
	url, err := url.JoinPath(endpoint, elem...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", defaultUserAgent)

	response, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Function GetUnmarshal GETs JSON data at the endpoint and unmarshals it into v
func GetUnmarshal(v any, elem ...string) error {
	// Grab the data
	data, err := Get(elem...)
	if err != nil {
		return err
	}

	// Unmarshal it
	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	return nil
}
