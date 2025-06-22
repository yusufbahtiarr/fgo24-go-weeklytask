package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func Order() {
	reader := bufio.NewReader(os.Stdin)

	for {
		menus.DisplayMenu(menus.OrderMenus)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "0":
			utils.Clear()
			return
		case "1":
			utils.Clear()
			handleCategory("Food", reader, &menus.Orders)
		case "2":
			utils.Clear()
			handleCategory("Drink", reader, &menus.Orders)
		case "3":
			utils.Clear()
			handleCategory("Snack", reader, &menus.Orders)
		case "4":
			utils.Clear()
			handleCategory("Appetizer", reader, &menus.Orders)
		default:
			utils.Clear()
			fmt.Print("Invalid option.")
			reader.ReadString('\n')
			utils.Clear()
			continue
		}

	}
}

func handleCategory(category string, reader *bufio.Reader, order *[]menus.OrderProduct) {
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
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		utils.Clear()
		return
	}

	num, err := strconv.Atoi(input)
	if err != nil || num < 1 || num > len(selectedItems) {
		fmt.Print("Invalid choice. ")
		reader.ReadString('\n')
		handleCategory(category, reader, order)
		return
	}

	selectedItem := selectedItems[num-1]

	var qty int
	for {
		fmt.Print("Enter quantity: ")
		qtyInput, _ := reader.ReadString('\n')
		qtyInput = strings.TrimSpace(qtyInput)

		var err error
		qty, err = strconv.Atoi(qtyInput)
		if err != nil || qty < 1 {
			fmt.Println("Invalid Input.")
			continue
		}
		break
	}

	*order = append(*order, menus.OrderProduct{
		Name:     selectedItem.Name,
		Price:    selectedItem.Price,
		Quantity: qty,
	})

	fmt.Printf("\nSuccessfully added %d x %s to your order.\n", qty, selectedItem.Name)

	for {
		fmt.Print("\nWould you like to place another order? (y/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "y":
			handleCategory(category, reader, order)
			return
		case "n":
			utils.Clear()
			return
		default:
			fmt.Println("Invalid option.")
			time.Sleep(time.Second)
		}
	}
}
