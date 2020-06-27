package models

type Pack struct {
	Type     string // basic plan - silver or gold
	Price    int64
	Channels []Channel
}

/*
	function to add a channel to array of channels in a pack
*/
func (MyPack *Pack) AddChannel(MyChannel Channel) {
	MyPack.Channels = append(MyPack.Channels, MyChannel)
}

/*
	function to check whether a channel exists in a pack
*/
func (pack *Pack) DoesChannelExistInCurrentPack(ch string) bool {
	var Exists bool
	Exists = false

	for i := 0; i < len(pack.Channels); i++ {
		if ch == pack.Channels[i].Name {
			Exists = true
			break
		}
	}
	return Exists
}
