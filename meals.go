package openmensa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Represents a menu item
type Meal struct {
	Id     int                `json:"id"`
	Name   string             `json:"name"`
	Notes  []string           `json:"notes"`
	Prices map[string]float64 `json:"prices"`
}

// Get all meal information served by a cateen on a given date
func GetMeals(canteenId int, date string) (*[]Meal, error) {
	response, err := http.Get(fmt.Sprintf("%s/canteens/%d/days/%s/meals", endpoint, canteenId, date))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []Meal
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}

// Get specific meal information
func GetMeal(canteenId int, date string, mealId int) (*Meal, error) {
	response, err := http.Get(fmt.Sprintf("%s/canteens/%d/days/%s/meals/%d", endpoint, canteenId, date, mealId))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject Meal
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}

func (m Meal) String() string {
	return m.Name
}
