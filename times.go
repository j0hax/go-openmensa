package openmensa

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Opening is a wrapper type for dates.
type Opening time.Time

// UnmarshalJSON parses a YYYY-MM-DD date to an Opening type.
func (o *Opening) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	t, err := time.Parse(dateLayout, string(s))
	if err != nil {
		return err
	}

	*o = Opening(t)
	return nil
}

// String returns a human-readable representation of a canteen's opening status.
func (o Opening) String() string {
	t := time.Time(o)
	return t.Format(dateLayout)
}

// Day represents a canteen's opening status.
type Day struct {
	// Date is the given date of operation.
	Date Opening `json:"date"`
	// Closed indicates if the canteen is closed on the given date.
	Closed bool `json:"closed"`
}

// String returns a human-readable representation of a canteen's opening data.
func (d *Day) String() string {
	var desc string
	if d.Closed {
		desc = "Closed"
	} else {
		desc = "Open"
	}

	return fmt.Sprintf("%s on %s", desc, d.Date)
}

// Days returns upcoming open/closed dates of a canteen.
func (c Canteen) Days() ([]Day, error) {
	var responseObject []Day
	cid := strconv.Itoa(c.Id)
	err := getUnmarshal(&responseObject, "canteens", cid, "days")
	return responseObject, err
}

// GetDay returns specific opening information of a given canteen on a given date.
func (c Canteen) Day(date time.Time) (*Day, error) {
	var responseObject Day
	strDate := date.Format(dateLayout)
	cid := strconv.Itoa(c.Id)
	err := getUnmarshal(&responseObject, "canteens", cid, "days", strDate)
	return &responseObject, err
}
