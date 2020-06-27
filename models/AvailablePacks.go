package models

type AvailablePacks struct {
	Packs []Pack
}

/*
	Used for appending channel to Channels array inside AvailablePacks datastructure
*/
func (MyAvailablePack *AvailablePacks) AddPacks(MyPack Pack) {
	MyAvailablePack.Packs = append(MyAvailablePack.Packs, MyPack)
}

/*
	Function for getting intial list of all the packs
*/
func getListOfPacks() AvailablePacks {
	zee := Channel{"Zee", 10}
	sony := Channel{"Sony", 15}
	starplus := Channel{"StarPlus", 20}
	discovery := Channel{"Discovery", 10}
	natgeo := Channel{"NatGeo", 20}

	silverpack := Pack{
		Type:  "Silver",
		Price: 50,
	}
	silverpack.AddChannel(zee)
	silverpack.AddChannel(sony)
	silverpack.AddChannel(starplus)

	goldpack := Pack{
		Type:  "Gold",
		Price: 100,
	}
	goldpack.AddChannel(zee)
	goldpack.AddChannel(sony)
	goldpack.AddChannel(starplus)
	goldpack.AddChannel(discovery)
	goldpack.AddChannel(natgeo)

	var packs AvailablePacks
	packs.AddPacks(silverpack)
	packs.AddPacks(goldpack)
	return packs
}
