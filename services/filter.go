package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strings"
)

func Filter() {
	for {
		utils.Clear()
		menus.DisplayMenu(menus.OrderMenus)
		input := utils.Input()

		switch input {
		case "0":
			utils.Clear()
			return
		case "1":
			filterCategory("Food", CartServices.GetCarts())
		case "2":
			filterCategory("Drink", CartServices.GetCarts())
		case "3":
			filterCategory("Snack", CartServices.GetCarts())
		case "4":
			filterCategory("Appetizer", CartServices.GetCarts())
		default:
			utils.InvalidInput()
			continue
		}
	}
}

func filterCategory(category string, order *[]Cart) {
	utils.Clear()
	menus.CafeName()

	var selectedItems []menus.Product
	GetProductByCategory(&selectedItems, category)
	DisplayPagination("Menu Category ", DefaultPagination(selectedItems, order))

}

func GetProductByCategory(selectedItems *[]menus.Product, category string) {
	for _, item := range menus.ListProduct {
		if strings.EqualFold(item.Category, category) {
			*selectedItems = append(*selectedItems, item)
		}
	}

	fmt.Printf("\n%s Menu:\n", category)
	for idx, item := range *selectedItems {
		fmt.Printf("%d. %s\n", idx+1, item.Display())
	}
}
