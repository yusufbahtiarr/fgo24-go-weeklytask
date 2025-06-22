package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"time"
)

func LoginMenu() {
	maxAttempts := 3
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		utils.Clear()
		menus.CafeName()
		fmt.Print("\nLOGIN \n")
		fmt.Print("Input Username : ")
		username := utils.Input()
		fmt.Print("Input Password : ")
		password := utils.Input()
		auth := LoginUser(username, password)
		if auth {
			fmt.Printf("\nLogin Success!")
			time.Sleep(time.Second)
			MainMenu()
		} else {
			fmt.Printf("\nInvalid username or password.")
			time.Sleep(time.Second)
		}

		if attempts == maxAttempts {
			{
				fmt.Printf("\n\nYou have exceeded the maximum number of login attempts.\n")
				utils.Exit()
			}
		}
	}
}
