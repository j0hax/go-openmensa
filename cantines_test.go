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

func TestLocalCanteens(t *testing.T) {
	canteens, err := CanteensNear(52.3830, 9.7197, 10)
	if err != nil {
		t.Error(err)
	}

	for i, c := range canteens {
		t.Logf("%d: %s [%f, %f]", i+1, c, c.Coordinates[0], c.Coordinates[1])
	}

	total := len(canteens)
	expected := 12
	if total != expected {
		t.Errorf("Expected %d canteens, got %d", expected, total)
	}
}
