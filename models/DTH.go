package models

import "fmt"

type DTH struct {
	availableServices AvailableServices
	availablePacks    AvailablePacks
	availableChannels AvailableChannels
}

/*
	function for populating dth object data like services, packs and channels
*/
func fillDetails(dth *DTH) {
	dth.availableServices = getListOfServices()
	dth.availablePacks = getListOfPacks()
	dth.availableChannels = getListOfChannels()
}

/*
	Make and return dth object
*/
func CreateDTHClient() DTH {
	dth := DTH{}
	fillDetails(&dth)
	return dth
}

/*
	function to query all the available plans,addons channels, and services.
*/
func (d *DTH) ViewAvailablePlan() {
	plans := d.availablePacks
	packs := plans.Packs
	fmt.Println("-List of Packs ")
	for i := 0; i < len(packs); i++ {
		fmt.Print("---Pack", i, " ")
		fmt.Println(packs[i].Type, packs[i].Price)
		ch_list := packs[i].Channels
		for j := 0; j < len(ch_list); j++ {
			ch := ch_list[j]
			fmt.Println("---------Channel", ch.Name, ch.Price)
		}
	}
	services := d.availableServices.Services
	fmt.Println("-List of Services")
	for i := 0; i < len(services); i++ {
		svc := services[i]
		fmt.Print("---Service", i, " ")
		fmt.Println(svc.Name, svc.Price)
	}
}
