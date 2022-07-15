package openmensa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Opening time.Time

func (o *Opening) UnmarshalJSON(text []byte) error {
	s := strings.Trim(string(text), "\"")
	t, err := time.Parse("2006-01-02", string(s))
	if err != nil {
		return err
	}

	*o = Opening(t)
	return nil
}

func (o Opening) String() string {
	t := time.Time(o)
	return t.Format("2006-01-02")
}

// Represents opening and closing days of a canteen
type Day struct {
	Date   Opening `json:"date"`
	Closed bool    `json:"closed"`
}

func (d *Day) String() string {
	var desc string
	if d.Closed {
		desc = "Closed"
	} else {
		desc = "Open"
	}

	return fmt.Sprintf("%s on %s", desc, d.Date)
}

// Returns upcoming opening and closing days for a canteen
func GetDays(canteenId int) (*[]Day, error) {
	response, err := http.Get(fmt.Sprintf("%s/canteens/%d/days", endpoint, canteenId))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []Day
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}

// Returns opening and closing day for a specific canteen and day
func GetDay(canteenId int, date string) (*Day, error) {
	response, err := http.Get(fmt.Sprintf("%s/canteens/%d/days/%s", endpoint, canteenId, date))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject Day
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}
