package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strconv"
	"strings"
)

func Order() {
	for {
		utils.Clear()
		menus.DisplayMenu(menus.OrderMenus)
		input := utils.Input()

		switch input {
		case "0":
			utils.Clear()
			return
		case "1":
			handleCategory("Food", &menus.Orders)
		case "2":
			handleCategory("Drink", &menus.Orders)
		case "3":
			handleCategory("Snack", &menus.Orders)
		case "4":
			handleCategory("Appetizer", &menus.Orders)
		default:
			utils.InvalidInput()
			continue
		}

	}
}

func handleCategory(category string, order *[]menus.OrderProduct) {
	utils.Clear()
	menus.CafeName()

	var selectedItems []menus.Product
	for _, item := range menus.ListProduct {
		if strings.EqualFold(item.Category, category) {
			selectedItems = append(selectedItems, item)
		}
	}

	fmt.Printf("\n%s Menu:\n", category)
	for idx, item := range selectedItems {
		fmt.Printf("%d. %s\n", idx+1, item.Display())
	}

	fmt.Print("\nChoose an option, or press Enter to go back: ")
	input := utils.Input()

	if input == "" {
		utils.Clear()
		return
	}

	num, err := strconv.Atoi(input)
	if err != nil || num < 1 || num > len(selectedItems) {
		utils.InvalidInput()
		handleCategory(category, order)
		return
	}

	selectedItem := selectedItems[num-1]

	var qty int
	for {
		fmt.Print("Enter quantity: ")
		qtyInput := utils.Input()

		var err error
		qty, err = strconv.Atoi(qtyInput)
		if err != nil || qty < 1 {
			fmt.Println("Invalid Input.")
			continue
		}
		break
	}

	AddToCart(order, selectedItem, qty)

	for {
		fmt.Print("\nWould you like to place another order? (y/n): ")
		input := utils.Input()

		switch input {
		case "y":
			handleCategory(category, order)
			return
		case "n":
			utils.Clear()
			return
		default:
			utils.InvalidInput()
		}
	}
}
