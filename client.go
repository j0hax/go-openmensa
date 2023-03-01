// Package openmensa provides an API to interface with OpenMensa.org.
package openmensa

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

// API endpoint URL
var Endpoint = "https://openmensa.org/api/v2"

// API User Endpoint
const defaultUserAgent = "go-openmensa/0.3"

// The client to use for HTTP requests
var c = http.Client{Timeout: time.Second * 10}

// get is a wrapper for http.Get(), using the predifined endpoint and custom headers.
func get(query url.Values, elem ...string) ([]byte, error) {

	path, err := url.JoinPath(Endpoint, elem...)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	if query != nil {
		url.RawQuery = query.Encode()
	}

	req, err := http.NewRequest("GET", url.String(), nil)
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

// getUnmarshal GETs JSON data at the endpoint and unmarshals it into v
func getUnmarshal(v any, elem ...string) error {
	// Grab the data
	data, err := get(nil, elem...)
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
