package main

import (
	"fmt"

	"github.com/asifk1997/mini_dth_software/input_output"
	"github.com/asifk1997/mini_dth_software/models"
)

// greeting function
func greetingAndInstructions() {
	fmt.Println("Welcome to DishTV")
	fmt.Println("1. View current balance and user details")
	fmt.Println("2. Recharge Account")
	fmt.Println("3. View available packs, channels and services")
	fmt.Println("4. Subscribe to base packs")
	fmt.Println("5. Add channels to an existing subscription")
	fmt.Println("6. Subscribe to special services")
	fmt.Println("7. View current subscription details")
	fmt.Println("8. Update email and phone number for notifications")
	fmt.Println("9. Exit")
}

// entrypoint

func main() {
	dth := models.CreateDTHClient()
	user := models.CreateUser()

	greetingAndInstructions()

forloop:
	for true {
		fmt.Println()
		fmt.Println("--------------------Enter your choice-------------------")
		text := input_output.TakeInput()
		switch text {
		case "1":
			user.ViewCurrentBalanceAndUserDetails()
		case "2":
			user.RechargeAccount()
		case "3":
			dth.ViewAvailablePlan()
		case "4":
			user.SubscribeAPlan(&dth)
		case "5":
			user.GetAddOnChannels(&dth)
		case "6":
			user.SubscribeSpecialServices(&dth)
		case "7":
			user.ViewCurrentSubscriptionDetails()
		case "8":
			user.UpdateMailAndPhone()
		case "9":
			fmt.Printf("Exiting")
			break forloop
		default:
			fmt.Printf("%s.\n", text)
		}
	}

}
