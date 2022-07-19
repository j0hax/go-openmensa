package openmensa

import (
	"testing"
)

func TestAllCantines(t *testing.T) {
	canteens, err := GetCanteens()
	if err != nil {
		t.Error("Could not retrieve information for canteens")
	}

	amount := len(canteens)

	t.Logf("Retrieved data for %d canteens", amount)
}
