package openmensa

import (
	"errors"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
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
	var responseObject []Meal
	cid := strconv.Itoa(canteenId)
	err := GetUnmarshal(&responseObject, "canteens", cid, "days", date, "meals")
	return responseObject, err
}

// GetMeals returns returns all current meals served by a canteen on today's date.
func GetMeals(canteenId int) ([]Meal, error) {
	date := time.Now().Format("2006-01-02")
	return GetMealsOn(canteenId, date)
}

// GetNextMeals gets all meals served by a canteen on the next opening date.
func GetNextMeals(canteenId int) ([]Meal, *Day, error) {
	// Get the opening dates
	days, err := GetDays(canteenId)
	if err != nil {
		return nil, nil, err
	}

	i := slices.IndexFunc(days, func(d Day) bool {
		return !d.Closed
	})

	if i < 0 {
		return nil, nil, errors.New("canteen is closed on all upcoming days")
	}

	firstOpening := days[i]

	meals, err := GetMealsOn(canteenId, firstOpening.Date.String())
	if err != nil {
		return nil, &firstOpening, err
	}

	return meals, &firstOpening, nil
}

// GetMeal returns a specific meal.
//
// A single meal is identified by its serving canteen, the day it is served on and its ID.
func GetMeal(canteenId int, date string, mealId int) (*Meal, error) {
	var responseObject Meal
	cid := strconv.Itoa(canteenId)
	mid := strconv.Itoa(canteenId)
	err := GetUnmarshal(&responseObject, "canteens", cid, "days", date, "meals", mid)
	return &responseObject, err
}

// String returns a human-readable representation of a meal.
//
// Currently, this is simply the meal's name.
func (m Meal) String() string {
	return m.Name
}
