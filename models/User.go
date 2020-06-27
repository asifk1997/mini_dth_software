package models

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	"github.com/asifk1997/mini_dth_software/input_output"
)

type User struct {
	Name                   string
	Email                  string
	Phone                  string
	Balance                int64
	SubscribedServices     []Service
	SubscribedPacks        Pack
	AddOnChannels          []Channel
	DurationOfSubscription int
}

/*
	function to fill initial info to user
*/
func fillUserInfo(user *User) {
	user.Name = "Asif"
	user.Balance = 100
	user.SubscribedPacks.Type = ""
}

/*
	create an object fill it with some initial data and return it
*/
func CreateUser() User {
	user := User{}
	fillUserInfo(&user)
	return user
}

/*
	function to add a channel to list of addon channels
*/
func (user *User) AddToAddONChannels(MyChannel Channel) {
	user.AddOnChannels = append(user.AddOnChannels, MyChannel)
}

/*
	function to add a service to list of subscribed services
*/
func (user *User) AddToServices(MyService Service) {
	user.SubscribedServices = append(user.SubscribedServices, MyService)
}

/*
	function to view user details
*/
func (v *User) ViewCurrentBalanceAndUserDetails() {
	fmt.Println("Your Name : ", v.Name)
	fmt.Println("Your Email : ", v.Email)
	fmt.Println("Your Phone : ", v.Phone)
	fmt.Println("Your Balance : ", v.Balance)
}

/*
	function which takes amount as input and add it to user balance
*/
func (u *User) RechargeAccount() {
	fmt.Println("your current balance is ", u.Balance)
	fmt.Print("Enter amount for recharge ")
	v := input_output.TakeInput()
	i, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		u.Balance += i
		fmt.Print("your updated balance is ")
		fmt.Println(u.Balance)
	}
}

func PrintAllPlansWithChannels(d *DTH) {
	a := d.availablePacks.Packs
	for i := 0; i < len(a); i++ {
		p := a[i]

		fmt.Println("--Plan Name : ", p.Type, " Price : ", p.Price, "Press ", p.Type, "for subscribing")
		ch_list := p.Channels
		for j := 0; j < len(ch_list); j++ {
			ch := ch_list[j]
			fmt.Println("---------Channel Name : ", ch.Name, " Channel Price : ", ch.Price)
		}
	}
}

/*
	function which subscribes a plan for a user basic or gold
*/
func (u *User) SubscribeAPlan(d *DTH) {

	PrintAllPlansWithChannels(d)

	fmt.Println("enter a plan ")
	text := input_output.TakeInput()
	if text == "Gold" || text == "Silver" { // check whether user input silver or gold or not

		// if u.SubscribedPacks.Type
		if u.SubscribedPacks.Type == "" { // if user is already not subscribed to a plan then only we move forward
			for i := 0; i < len(d.availablePacks.Packs); i++ { // we traverse to all the plans provided by dth
				p := d.availablePacks.Packs[i]
				if text == p.Type { // if user inputed plan matches with plan provided by dth we move further

					fmt.Println("Enter duration of subscription in months") // we ask for months for which user wants basic plan
					duration_text := input_output.TakeInput()
					duration, err := strconv.ParseInt(duration_text, 10, 0)
					if err == nil {
						cost := p.Price * duration
						fmt.Println("cost ", cost, " ")
						if duration >= 3 { // if the duration of subscription is atleast 3 months we apply 10 % discount
							cost = cost * 90 / 100
							fmt.Println(" cost after discount ", cost, " ")
						}
						if u.Balance >= cost { // if user has balance greater than the price of pack
							u.SubscribedPacks = p
							u.Balance -= cost
							fmt.Println("Subscribed ", u.SubscribedPacks.Type)
							fmt.Println("Email notification sent successfully")
							fmt.Println("SMS notification sent successfully")

						} else {
							fmt.Println("You do not have enough balance")
						}

					}
					break
				}
			}
		} else {
			fmt.Print("you have subscribed a plan already")
			fmt.Println(u.SubscribedPacks.Type)
		}

	} else { // if user chooses something other than silver or gold
		fmt.Println("please choose a valid pack")
	}

}

/*
	function to print list of all the addon channels
*/
func PrintAllAddOnChannels(d *DTH) {
	ch_list := d.availableChannels.Channels
	for i := 0; i < len(ch_list); i++ {
		ch := ch_list[i]
		fmt.Println("Channel Name : ", ch.Name, " Price : ", ch.Price)
	}
}

/*
	function to subscribe addon channels
*/
func (u *User) GetAddOnChannels(d *DTH) {

	if u.SubscribedPacks.Type == "" { // if user doesnt have base pack he cannot have addons
		fmt.Println("you should subscibe to a base plan first")
	} else {
		PrintAllAddOnChannels(d) // we show user a list of channels to choose from
		fmt.Println("Enter channel name you wish to add as Addon")
		text := input_output.TakeInput() // we take input from user

		if d.availableChannels.DoesExistInChannelList(text) { // if user input matches with list we showed him or not
			if u.SubscribedPacks.DoesChannelExistInCurrentPack(text) { // if channel is already included in his base plan
				fmt.Println("you do not neet to buy this channel, it is already included in your base plan")
			} else {
				ch := d.availableChannels.GetChannelFromChannelList(text) // get channel from list
				cost := ch.Price
				if u.Balance >= cost { // if current balance is greater than the price of channel
					u.Balance -= cost
					u.AddToAddONChannels(ch)
					fmt.Println("AddOn Channel Added Susscessfully")
					fmt.Println("Your current balance is ", u.Balance)
				} else {
					fmt.Println("Sorry you do not have enough balance")
				}
			}
		} else {
			fmt.Println("channel does not exist in our channel list, Please change your choice")
		}
	}
}

/*
	function to print all the additional services
*/
func PrintAllSpecialServices(d *DTH) {
	fmt.Println("Here are all the special services you can choose")
	for i := 0; i < len(d.availableServices.Services); i++ {
		svc := d.availableServices.Services[i]
		fmt.Println("Service Name : ", svc.Name, ", Service Price", svc.Price)
	}
}

/*
	Subscribe to special service function
*/
func (u *User) SubscribeSpecialServices(d *DTH) {

	if u.SubscribedPacks.Type == "" { // a user should subscibe base plan before subscribing additional services
		fmt.Println("you should subscibe to a base plan first")
	} else {
		PrintAllSpecialServices(d) // show list of services to user
		fmt.Print("Enter name of service : ")
		text := input_output.TakeInput()

		if d.availableServices.DoesServiceExist(text) { // match user input with one of the service in the list
			svc := d.availableServices.GetService(text) // get service
			if u.DoesUserHaveServiceAlready(svc.Name) { // check if user has service already
				fmt.Println("you already have this service, you do not need to buy it")
			} else {
				if u.Balance >= svc.Price { // if user has sufficinet balance
					u.Balance -= svc.Price
					u.AddToServices(svc)
					fmt.Println("Service ", svc.Name, " has been successfully added to your account")
					fmt.Println("Your current balance is ", u.Balance)
				} else {
					fmt.Println("you do not have enough balance")
				}
			}
		} else {
			fmt.Println("This service does not exist, Please enter a valid service")
		}
	}
}

/*
	check if the user has service already
*/
func (u *User) DoesUserHaveServiceAlready(ch string) bool {
	var Exists bool
	Exists = false

	for i := 0; i < len(u.SubscribedServices); i++ {
		if ch == u.SubscribedServices[i].Name {
			Exists = true
			break
		}
	}
	return Exists
}

/*
	print all the information about the subscription details
*/
func (u *User) ViewCurrentSubscriptionDetails() {
	if u.SubscribedPacks.Type == "" { // if user is not subscibed to any plan then subsequently he will not have anything
		fmt.Println("You are subscribed not subscribed to any plan please subscribe")
	} else {
		fmt.Println("You are subscribed to ", u.SubscribedPacks.Type)
		fmt.Println("This is list of channels included in pack", "Number of channels ", len(u.SubscribedPacks.Channels))

		for i := 0; i < len(u.SubscribedPacks.Channels); i++ {
			ch := u.SubscribedPacks.Channels[i]
			fmt.Println("Channel Name ", ch.Name, " Channel Price : ", ch.Price)
		}

		fmt.Println("This is list of addon channels", "Number of addon Channels ", len(u.AddOnChannels))

		for i := 0; i < len(u.AddOnChannels); i++ {
			ch := u.AddOnChannels[i]
			fmt.Println("Channel Name ", ch.Name, " Channel Price : ", ch.Price)
		}

		fmt.Println("This is list of addon services ", "Number of services ", len(u.SubscribedServices))
		for i := 0; i < len(u.SubscribedServices); i++ {
			svc := u.SubscribedServices[i]
			fmt.Println("Service Name ", svc.Name, " Service Price : ", svc.Price)
		}
	}

}

// regex for email validation
var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

/*
Update Mail And Phone function
*/
func (u *User) UpdateMailAndPhone() {
	fmt.Print("Enter your updated Email : ")
	text := input_output.TakeInput()                   //take user input
	if len(text) > 254 || !rxEmail.MatchString(text) { // if is in email format
		fmt.Println("error: ", text, " is not a valid email address")
	} else {
		u.Email = text
		fmt.Println("Email Updated")
		fmt.Print("Enter your updated phonenumber : ")
		phoneNumber := input_output.TakeInput() // take user input for phone
		if isInt(phoneNumber) {                 // if phonenumber consists of all characters as digits
			u.Phone = phoneNumber
			fmt.Println("Phone Updated")
			u.ViewCurrentBalanceAndUserDetails() // call to a function which prints user details on screen
		} else {
			fmt.Println("Phone number should contain only digits")
		}
	}
}

/*
	function to check all the characters are numbers
*/
func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
