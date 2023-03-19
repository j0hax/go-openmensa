package openmensa

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Simplified ISO 8601 date layout used by OpenMensa.
//
// Can be used for Time.Format() et al.
const DateLayout = "2006-01-02"

// Meal is the representation of a canteen's menu item.
type Meal struct {
	// Id is a unique identifier for the meal.
	Id int `json:"id"`
	// Name is the title of the meal.
	Name string `json:"name"`
	// Optional category metadata
	Category string `json:"category"`
	// Notes include extra information, such as allergens.
	Notes []string `json:"notes"`
	// Prices vary for different groups of patrons.
	//
	// Note that the groups vary by canteen operator.
	// Typically these include "students", "employees", "others", and "pupils".
	Prices map[string]float64 `json:"prices"`
}

// UnmarshalJSON is a custom unmarshaller for Meals.
// This function acts as a regular json.Unmarshal, but removes duplicate note entries.
//
// This function may be removed should the official OpenMensa API perform server-side duplicate handling one day.
func (m *Meal) UnmarshalJSON(data []byte) error {
	// To avoid infinite unmarshal recursions, a type alias is used.
	// Big thanks to https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers/
	type TmpType Meal
	var tmpMeal TmpType

	err := json.Unmarshal(data, &tmpMeal)
	if err != nil {
		return err
	}

	// Convert from the temporary type to our regular struct
	*m = Meal(tmpMeal)

	// Remove duplicate notes
	exists := make(map[string]bool, len(m.Notes))
	new := make([]string, 0, len(m.Notes))
	for _, item := range m.Notes {
		if !exists[item] {
			exists[item] = true
			new = append(new, item)
		}
	}
	m.Notes = new

	return nil
}

// Menu represents all meals served by a canteen on a given day.
//
// In German, this is the semantic equivalent to a "Speiseplan"
type Menu struct {
	Day   Day    `json:"date"`
	Meals []Meal `json:"meals"`
}

// Custom unmarshaller to get around the inconsistency between
// /canteens/{id}/days/{date}, which returns an object with two attributes, and
// /canteens/{id}/meals, which returns two attributes and meals directly
//
// This function may be removed should the official OpenMensa API return opening information as a single JSON object one day.
func (m *Menu) UnmarshalJSON(data []byte) error {
	var interim struct {
		Date   Opening `json:"date"`
		Closed bool    `json:"closed"`
		Meals  []Meal  `json:"meals"`
	}

	err := json.Unmarshal(data, &interim)
	if err != nil {
		return err
	}

	*m = Menu{
		Day: Day{
			Date:   interim.Date,
			Closed: interim.Closed,
		},
		Meals: interim.Meals,
	}

	return nil
}

// MenuOn returns returns all meals served by a canteen on a given date.
func (c *Canteen) MenuOn(date time.Time) (*Menu, error) {
	dateResponse, err := c.Day(date)
	if err != nil {
		return nil, err
	}

	strDate := date.Format(DateLayout)
	cid := strconv.Itoa(c.Id)

	mList := make([]Meal, 0, 10)
	err = getUnmarshal(&mList, "canteens", cid, "days", strDate, "meals")
	if err != nil {
		return nil, fmt.Errorf("retrieve menu for canteen ID %d on %s: %w", c.Id, strDate, err)
	}

	menu := Menu{
		Day:   *dateResponse,
		Meals: mList,
	}

	return &menu, nil
}

// CurrentMenu returns returns all meals served by a canteen on today's date.
func (c *Canteen) CurrentMenu() (*Menu, error) {
	date := time.Now()
	return c.MenuOn(date)
}

// AllMenus returns all meals for all upcoming dates
func (c *Canteen) AllMenus() ([]Menu, error) {
	// Preallocate enough for a week
	responseData := make([]Menu, 0, 7)

	cid := strconv.Itoa(c.Id)
	err := getUnmarshal(&responseData, "canteens", cid, "meals")
	if err != nil {
		return nil, fmt.Errorf("retrieve menus for canteen ID %d: %w", c.Id, err)
	}

	return responseData, err
}

// Meal returns a specific meal.
//
// A single meal is identified by the day it is served on and its ID.
func (c *Canteen) Meal(date time.Time, mealId int) (*Meal, error) {
	strDate := date.Format(DateLayout)
	var responseObject Meal
	cid := strconv.Itoa(c.Id)
	mid := strconv.Itoa(mealId)
	err := getUnmarshal(&responseObject, "canteens", cid, "days", strDate, "meals", mid)
	if err != nil {
		return nil, fmt.Errorf("retrieve meal ID %d served by canteen ID %d on %s: %w", mealId, c.Id, strDate, err)
	}

	return &responseObject, nil
}

// String returns a human-readable representation of a meal.
//
// Currently, this is simply the meal's name.
func (m Meal) String() string {
	return m.Name
}
