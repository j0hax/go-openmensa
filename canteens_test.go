package openmensa

import (
	"testing"
)

func TestAllCanteens(t *testing.T) {
	canteens, err := AllCanteens()
	if err != nil {
		t.Error("Could not retrieve information for canteens")
	}

	amount := len(canteens)

	t.Logf("Retrieved data for %d canteens", amount)
}

func TestSearchCanteens(t *testing.T) {
	canteens, err := SearchCanteens("Hannover")
	if err != nil {
		t.Error(err)
	}

	for i, c := range canteens {
		t.Logf("%d: %s", i+1, c)
	}
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
