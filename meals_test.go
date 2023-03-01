package openmensa

import (
	"testing"
)

func TestMeals(t *testing.T) {
	contine, err := GetCanteen(7)
	if err != nil {
		t.Error(err)
	}

	meals, day, err := contine.GetNextMeals()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Next opening day: %s\n", day.Date)

	for _, m := range meals {
		// Ensure our meal isn't empty
		if len(m.Name) == 0 {
			t.Errorf("Meal ID %d has empty name", m.Id)
		}

		t.Log(m)
	}
}
