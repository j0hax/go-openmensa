package openmensa

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"regexp"
	"strconv"

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
	q := url.Values{}
	page := 1

	var allCanteens []Canteen

	// Repeatedly query the next page until none are returned
	for {
		q.Set("page", strconv.Itoa(page))
		var canteens []Canteen

		// Grab data with custom page query and unmarshal it
		data, err := get(q, "canteens")
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &canteens)
		if err != nil {
			return nil, err
		}

		if len(canteens) == 0 {
			break
		}

		// Save and continue to next page
		page = page + 1
		allCanteens = append(allCanteens, canteens...)
	}

	return allCanteens, nil
}

// GetCanteen returns data about a specific canteen.
func GetCanteen(canteenId int) (*Canteen, error) {
	var responseObject Canteen
	err := getUnmarshal(&responseObject, "canteens", strconv.Itoa(canteenId))
	return &responseObject, err
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
		return nil, errors.New("no matching canteen found")
	}

	return &(canteens[i]), nil
}
