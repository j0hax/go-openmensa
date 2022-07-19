package openmensa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Meal is the representation of a canteen's menu item.
type Meal struct {
	// Id is a unique identifier for the meal.
	Id int `json:"id"`
	// Name is the title of the meal.
	Name string `json:"name"`
	// Notes include extra information, such as allergens.
	Notes []string `json:"notes"`
	// Prices vary for different groups of patrons.
	//
	// Note that the groups vary by canteen operator.
	// Typically these include "students", "employees", "others", and "pupils".
	Prices map[string]float64 `json:"prices"`
}

// GetMealsOn returns returns all meals served by a canteen on a given date.
func GetMealsOn(canteenId int, date string) ([]Meal, error) {
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

	return responseObject, nil
}

// GetMeals returns returns all current meals served by a canteen on today's date.
func GetMeals(canteenId int) ([]Meal, error) {
	response, err := http.Get(fmt.Sprintf("%s/canteens/%d/meals", endpoint, canteenId))

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

	return responseObject, nil
}

// GetMeal returns a specific meal.
//
// A single meal is identified by its serving canteen, the day it is served on and its ID.
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

// String returns a human-readable representation of a meal.
//
// Currently, this is simply the meal's name.
func (m Meal) String() string {
	return m.Name
}
