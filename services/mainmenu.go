package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
)

func MainMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		utils.Clear()
		menus.DisplayMenu(menus.MainMenus)

		input := utils.Input()
		switch input {
		case "0":
			os.Exit(0)
		case "1":
			ListMenu()
		case "2":
			Search()
		case "3":
			Order()
		case "4":
			Cart()
		case "5":
			MenuAuth()
		case "6":
			return
		default:
			fmt.Print("Invalid options. ")
			reader.ReadString('\n')
			continue
		}
	}
}
