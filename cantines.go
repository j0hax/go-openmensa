package openmensa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Represents an invidual canteen
type Canteen struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	Coordinates []float64 `json:"coordinates"`
}

func (m Canteen) String() string {
	return m.Name
}

func GetCanteens() (*[]Canteen, error) {
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

	return &responseObject, nil
}

// Get attributes of a specific canteen from its ID number
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
