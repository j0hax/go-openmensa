package openmensa

import (
	"testing"
)

func TestMeals(t *testing.T) {
	meals, err := GetMeals(7)
	if err != nil {
		t.Error(err)
	}

	for _, m := range meals {
		// Ensure our meal isn't empty
		if len(m.Name) == 0 {
			t.Errorf("Meal ID %d has empty name", m.Id)
		}

		t.Log(m)
	}
}
