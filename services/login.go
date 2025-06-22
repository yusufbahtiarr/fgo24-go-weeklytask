package services

import (
	"fmt"
	"go-booking-menu/utils"
	"time"
)

func LoginMenu() {
	for {
		utils.Clear()
		fmt.Println("------- LOGIN -------")
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
	}
}
