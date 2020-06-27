package models

import (
	"testing"
)

func TestDoesExistInChannelList(t *testing.T) {
	var dth DTH
	//fmt.Println(dth)
	// var user User
	sab := Channel{"Sab", 50}
	cn := Channel{"CN", 50}
	zee := Channel{"Zee", 10}
	tables := []struct {
		x Channel
		y bool
	}{
		{sab, false},
		{cn, false},
		{zee, true},
	}

	for _, table := range tables {
		total := dth.availableChannels.DoesExistInChannelList(table.x.Name)
		if total == true {
			t.Errorf("Error since channel does not exist now")
		}
	}

	dth = CreateDTHClient() // when client is created it adds all the channels in available channels

	for _, table := range tables {
		total := dth.availableChannels.DoesExistInChannelList(table.x.Name)
		if total == table.y {

		} else {
			t.Errorf(" not matching with the list provided")
		}
	}
}
