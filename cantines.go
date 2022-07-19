package openmensa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/exp/slices"
)

// Canteen contains information associated with a specific canteen, cafe, cafeteria, etc.
type Canteen struct {
	// Id is a unique identifier of the canteen.
	Id int `json:"id"`
	// Name of the canteen.
	Name string `json:"name"`
	// City the canteen is located in.
	City string `json:"city"`
	// Address of the canteen.
	Address string `json:"address"`
	// Geographic coordinates of the canteen.
	Coordinates []float64 `json:"coordinates"`
}

// String returns a human-readable representation of the canteen.
//
// Currently, this is simply the canteen's name.
func (m Canteen) String() string {
	return m.Name
}

// GetCanteens returns a list of all canteens and their metadata.
func GetCanteens() ([]Canteen, error) {
	response, err := http.Get(fmt.Sprintf("%s/canteens", endpoint))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []Canteen
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return responseObject, nil
}

// GetCanteen returns data about a specific canteen.
func GetCanteen(canteenId int) (*Canteen, error) {
	response, err := http.Get(fmt.Sprintf("%s/canteens/%d", endpoint, canteenId))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject Canteen
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}

// FindCanteen searches the list of canteens and return the first canteen
// whose name matches the specified pattern
func FindCanteen(pattern string) (*Canteen, error) {
	canteens, err := GetCanteens()

	if err != nil {
		return nil, err
	}

	i := slices.IndexFunc(canteens, func(c Canteen) bool {
		matched, err := regexp.MatchString(pattern, c.Name)
		if err != nil {
			log.Panic(err)
			return false
		}
		return matched
	})

	if i < 0 {
		return nil, errors.New("No matching canteen found")
	}

	return &(canteens[i]), nil
}
