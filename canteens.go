package openmensa

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
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

// AllCanteens returns a list of all canteens and their metadata.
func AllCanteens() ([]Canteen, error) {
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
			return nil, fmt.Errorf("retrieve all canteens: %w", err)
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

// CanteensNear returns canteens in the radius of the given latitude and longitude
func CanteensNear(latitude, longitude, distance float64) ([]Canteen, error) {
	q := url.Values{}

	lat := fmt.Sprintf("%f", latitude)
	lng := fmt.Sprintf("%f", longitude)
	dist := fmt.Sprintf("%f", distance)

	q.Set("near[lat]", lat)
	q.Set("near[lng]", lng)
	q.Set("near[dist]", dist)

	page := 1

	var nearby []Canteen

	// Repeatedly query the next page until none are returned
	for {
		q.Set("page", strconv.Itoa(page))
		var canteens []Canteen

		// Grab data with custom page query and unmarshal it
		data, err := get(q, "canteens")
		if err != nil {
			return nil, fmt.Errorf("retrieve canteens near [%f, %f]: %w", latitude, longitude, err)
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
		nearby = append(nearby, canteens...)
	}

	return nearby, nil
}

// GetCanteen returns data about a specific canteen.
func GetCanteen(canteenId int) (*Canteen, error) {
	var responseObject Canteen
	err := getUnmarshal(&responseObject, "canteens", strconv.Itoa(canteenId))
	if err != nil {
		return nil, fmt.Errorf("retrieve canteen with ID %d: %w", canteenId, err)
	}
	return &responseObject, nil
}

// GetCanteens retrieves multiple canteens specified by their IDs.
func GetCanteens(canteenIds ...int) ([]Canteen, error) {
	var canteens []Canteen
	for _, id := range canteenIds {
		canteen, err := GetCanteen(id)
		if err != nil {
			return nil, err
		}

		canteens = append(canteens, *canteen)
	}

	return canteens, nil
}

// SearchCanteens returns a slice of canteens whose names match the given pattern
func SearchCanteens(pattern string) ([]Canteen, error) {
	all, err := AllCanteens()
	if err != nil {
		return nil, err
	}

	var matches []Canteen

	for _, c := range all {
		matched, err := regexp.MatchString(pattern, c.Name)
		if err != nil {
			return nil, err
		}

		if matched {
			matches = append(matches, c)
		}
	}

	return matches, nil
}
