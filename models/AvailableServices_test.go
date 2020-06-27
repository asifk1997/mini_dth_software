package models

import (
	"testing"
)

func TestDoesExistIn(t *testing.T) {
	var dth DTH
	service1 := Service{
		Name:  "LearnEnglish",
		Price: 100,
	}
	service2 := Service{
		Name:  "LearnHindi",
		Price: 200,
	}
	service3 := Service{
		Name:  "LearnCooking",
		Price: 300,
	}
	tables := []struct {
		x Service
		y bool
	}{
		{service1, true},
		{service2, true},
		{service3, false},
	}

	for _, table := range tables {
		total := dth.availableServices.DoesServiceExist(table.x.Name)
		if total == true {
			t.Errorf("Error since channel does not exist now")
		}
	}

	dth = CreateDTHClient()

	for _, table := range tables {
		total := dth.availableServices.DoesServiceExist(table.x.Name)
		if total == table.y {

		} else {
			t.Errorf(" not matching with the list provided")
		}
	}
}
