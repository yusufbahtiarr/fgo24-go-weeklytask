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
			Filter()
		case "4":
			CartView()
		case "5":
			Checkout()
		case "6":
			MenuAuth()
		default:
			utils.Input()
			continue
		}
	}
}
