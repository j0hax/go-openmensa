package openmensa

import (
	"errors"
	"log"
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
	var responseObject []Canteen
	err := GetUnmarshal(&responseObject, "canteens")
	return responseObject, err
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
