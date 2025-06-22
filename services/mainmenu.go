package services

import (
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
)

func MainMenu() {
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
			utils.Exit()
		default:
			utils.Input()
			continue
		}
	}
}
