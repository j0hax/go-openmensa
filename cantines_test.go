package openmensa

import (
	"testing"
)

func TestAllCantines(t *testing.T) {
	canteens, err := AllCanteens()
	if err != nil {
		t.Error("Could not retrieve information for canteens")
	}

	amount := len(canteens)

	t.Logf("Retrieved data for %d canteens", amount)
}

func matchesID(c *Canteen, expected int, t *testing.T) {
	if c == nil {
		t.Fatal("Cantine is nil pointer")
	}

	if c.Id != expected {
		t.Errorf("Got Canteen with ID %d, expected %d", c.Id, expected)
	}
}

func TestFindCanteen(t *testing.T) {
	const (
		contineID = 7
		garbsenID = 8
	)

	c, err := FindCanteen(`Contine`)
	if err != nil {
		t.Fatal(err)
	}

	matchesID(c, contineID, t)

	c, err = FindCanteen(`Mensa PZH`)
	if err != nil {
		t.Fatal(err)
	}

	matchesID(c, garbsenID, t)
}
