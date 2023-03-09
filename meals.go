package openmensa

import (
	"encoding/json"
	"strconv"
	"time"
)

// Date Format used by OpenMensa (simplified ISO 8601)
const dateLayout = "2006-01-02"

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

// Menu represents all meals served by a canteen on a given day.
//
// In German, this is the semantic equivalent to a "Speiseplan"
type Menu struct {
	Day   Day    `json:"date"`
	Meals []Meal `json:"meals"`
}

// Custom unmarshaler to get around the inconsistency between
// /canteens/{id}/days/{date}, which returns an object with two attributes, and
// /canteens/{id}/meals, which returns two attributes and meals directly
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
func (c Canteen) MenuOn(date time.Time) (*Menu, error) {
	strDate := date.Format(dateLayout)
	cid := strconv.Itoa(c.Id)

	var dateResponse Day
	err := getUnmarshal(&dateResponse, "canteens", cid, "days", strDate)
	if err != nil {
		return nil, err
	}

	var mList []Meal
	err = getUnmarshal(&mList, "canteens", cid, "days", strDate, "meals")
	if err != nil {
		return nil, err
	}

	menu := Menu{
		Day:   dateResponse,
		Meals: mList,
	}

	return &menu, err
}

// CurrentMenu returns returns all meals served by a canteen on today's date.
func (c Canteen) CurrentMenu() (*Menu, error) {
	date := time.Now()
	return c.MenuOn(date)
}

// AllMenus returns all meals for all upcoming dates
func (c Canteen) AllMenus() ([]Menu, error) {
	var responseData []Menu
	cid := strconv.Itoa(c.Id)
	err := getUnmarshal(&responseData, "canteens", cid, "meals")

	return responseData, err
}

// Meal returns a specific meal.
//
// A single meal is identified by the day it is served on and its ID.
func (c Canteen) Meal(date time.Time, mealId int) (*Meal, error) {
	strDate := date.Format(dateLayout)
	var responseObject Meal
	cid := strconv.Itoa(c.Id)
	mid := strconv.Itoa(mealId)
	err := getUnmarshal(&responseObject, "canteens", cid, "days", strDate, "meals", mid)
	return &responseObject, err
}

// String returns a human-readable representation of a meal.
//
// Currently, this is simply the meal's name.
func (m Meal) String() string {
	return m.Name
}
