package models

type AvailableChannels struct {
	Channels []Channel
}

/*
	Used for appending channel to Channels array inside AvailableChannel datastructure
*/
func (MyAvailableChannel *AvailableChannels) AddChannels(ch Channel) {
	MyAvailableChannel.Channels = append(MyAvailableChannel.Channels, ch)
}

/*
	Function for popualating initial data of all the channels
*/
func getListOfChannels() AvailableChannels {
	zee := Channel{"Zee", 10}
	sony := Channel{"Sony", 15}
	starplus := Channel{"StarPlus", 20}
	discovery := Channel{"Discovery", 10}
	natgeo := Channel{"NatGeo", 20}
	max := Channel{"Max", 30}
	var availableChannels AvailableChannels
	availableChannels.AddChannels(zee)
	availableChannels.AddChannels(sony)
	availableChannels.AddChannels(starplus)
	availableChannels.AddChannels(discovery)
	availableChannels.AddChannels(natgeo)
	availableChannels.AddChannels(max)
	return availableChannels
}

/*
	This Function checks whether a selected channel exists in array of channels in available channels
*/
func (MyAvailableChannel *AvailableChannels) DoesExistInChannelList(ch string) bool {
	var Exists bool
	Exists = false

	for i := 0; i < len(MyAvailableChannel.Channels); i++ {
		if ch == MyAvailableChannel.Channels[i].Name {
			Exists = true
			break
		}
	}
	return Exists
}

/*
	This Function returns the channel object mathcing with channel name.
*/
func (MyAvailableChannel *AvailableChannels) GetChannelFromChannelList(ch string) Channel {
	var Exists Channel
	for i := 0; i < len(MyAvailableChannel.Channels); i++ {
		if ch == MyAvailableChannel.Channels[i].Name {
			Exists = MyAvailableChannel.Channels[i]
			break
		}
	}
	return Exists
}
