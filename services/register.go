package services

import (
	"fmt"
	"go-booking-menu/utils"
	"time"
)

func RegisterMenu() {
	for {
		utils.Clear()
		fmt.Println("------- REGISTER -------")
		fmt.Print("Input Username : ")
		username := utils.Input()
		fmt.Print("Input Password : ")
		password := utils.Input()
		fmt.Print("Confirm Password : ")
		confirm_password := utils.Input()
		if password != confirm_password {
			fmt.Println("Password do not match.")
			time.Sleep(time.Second)
		} else if username == "" || password == "" {
			fmt.Println("Username or Password cannot empty.")
			time.Sleep(time.Second)
		} else {
			AddUser(username, password)
			fmt.Printf("\nRegistration successful!")
			time.Sleep(time.Second)
			return
		}
	}
}
