package openmensa

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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
func (c Canteen) String() string {
	return c.Name
}

// queryCanteens returns a list of all canteens and their metadata.
//
// The parameter q allows for specific tuning of pagination and other parameters:
// For more information see https://docs.openmensa.org/api/v2/canteens/#list-params
func queryCanteens(q url.Values) ([]Canteen, error) {
	currentPage := 1

	var allCanteens []Canteen

	// Repeatedly query the next page until none are returned
	for {
		q.Set("page", strconv.Itoa(currentPage))

		// Grab data with custom page query
		data, hdr, err := get(q, "canteens")
		if err != nil {
			return nil, fmt.Errorf("retrieve all canteens: %w", err)
		}

		// Preallocate main mensa list if needed
		if len(allCanteens) == 0 {
			size, err := strconv.Atoi(hdr.Get("x-total-count"))
			if err != nil {
				return nil, fmt.Errorf("could not determine total mensa count: %w", err)
			}
			allCanteens = make([]Canteen, 0, size)
		}

		// Preallocate current page of mensas
		size, err := strconv.Atoi(hdr.Get("x-per-page"))
		if err != nil {
			return nil, fmt.Errorf("could not determine page size: %w", err)
		}
		canteens := make([]Canteen, 0, size)

		// Marshal the current page
		err = json.Unmarshal(data, &canteens)
		if err != nil {
			return nil, err
		}

		// Check if we have reached the last page
		totalPages, err := strconv.Atoi(hdr.Get("x-total-pages"))
		if err != nil {
			return nil, fmt.Errorf("could not determine total page count: %w", err)
		}

		// Save and continue to next page
		allCanteens = append(allCanteens, canteens...)
		currentPage++

		if currentPage > totalPages {
			break
		}
	}

	return allCanteens, nil
}

// AllCanteens returns all canteens listed in OpenMensa
func AllCanteens() ([]Canteen, error) {
	q := url.Values{}
	return queryCanteens(q)
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

	return queryCanteens(q)
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
	q := url.Values{}

	// convert to string arraw
	var stringIDs = make([]string, len(canteenIds))
	for i := range canteenIds {
		stringIDs[i] = strconv.Itoa(canteenIds[i])
	}

	idString := strings.Join(stringIDs, ",")
	q.Set("ids", idString)

	return queryCanteens(q)
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
