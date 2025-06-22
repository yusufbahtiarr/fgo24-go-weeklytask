package services

import (
	"go-booking-menu/menus"
	"go-booking-menu/utils"
)

func MenuAuth() {
	for {
		utils.Clear()
		menus.DisplayMenu(menus.AuthMenus)
		input := utils.Input()
		switch input {
		case "1":
			LoginMenu()
		case "2":
			RegisterMenu()
		case "0":
			utils.Exit()
		default:
			utils.InvalidInput()
		}
	}

}
