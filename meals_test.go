package openmensa

import (
	"testing"
	"time"
)

func TestCurrentMenu(t *testing.T) {
	contine, err := GetCanteen(7)
	if err != nil {
		t.Error(err)
	}

	menu, err := contine.CurrentMenu()
	if err != nil {
		t.Error(err)
	}

	t.Log(menu.Day.String())

	// Ensure our meal isn't empty
	for _, meal := range menu.Meals {
		t.Logf("\t%s\n", meal)
	}
}

func TestAllMenus(t *testing.T) {
	contine, err := GetCanteen(7)
	if err != nil {
		t.Error(err)
	}

	menu, err := contine.AllMenus()
	if err != nil {
		t.Error(err)
	}

	for _, entry := range menu {
		t.Log(entry.Day.String())

		// Ensure our meal isn't empty
		for _, meal := range entry.Meals {
			t.Logf("\t%s\n", meal)
		}
	}
}

func TestNonExistantMenu(t *testing.T) {
	contine, err := GetCanteen(7)
	if err != nil {
		t.Error(err)
	}

	_, err = contine.MenuOn(time.Date(1, time.January, 1, 1, 1, 1, 1, time.Local))

	t.Log(err)

	if err == nil {
		t.Fail()
	}
}
