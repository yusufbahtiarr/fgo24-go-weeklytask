package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/data"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
	"strconv"
	"strings"
)

func Order() {
	reader := bufio.NewReader(os.Stdin)

	for {
		CafeName()
		fmt.Print(`Choose Category:
1. Foods
2. Drinks
3. Snacks
4. Appetizers

0. Back to Main Menu

Choose an option (0-4): `)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "0":
			utils.Clear()
			return
		case "1":
			utils.Clear()
			handleCategory("Food", reader, &data.Orders)
		case "2":
			utils.Clear()
			handleCategory("Drink", reader, &data.Orders)
		case "3":
			utils.Clear()
			handleCategory("Snack", reader, &data.Orders)
		case "4":
			utils.Clear()
			handleCategory("Appetizer", reader, &data.Orders)
		default:
			utils.Clear()
			fmt.Print("Invalid option.")
			reader.ReadString('\n')
			utils.Clear()
			continue
		}

	}
}

func handleCategory(category string, reader *bufio.Reader, order *[]menus.OrderMenu) {
	utils.Clear()
	CafeName()

	var selectedItems []menus.Menu
	for _, item := range menus.ListMenu {
		if strings.EqualFold(item.Category, category) {
			selectedItems = append(selectedItems, item)
		}
	}

	fmt.Printf("\n%s Menu:\n", category)
	for idx, item := range selectedItems {
		fmt.Printf("%d. %s\n", idx+1, item.DisplayName())
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
			fmt.Println("Invalid quantity. Please enter a positive number.")
			continue
		}
		break
	}

	*order = append(*order, menus.OrderMenu{
		Name:     selectedItem.Name,
		Price:    selectedItem.Price,
		Quantity: qty,
	})

	fmt.Printf("\nSuccessfully added %d x %s to your order.\n", qty, selectedItem.Name)

	for {
		fmt.Print("\nWould you like to place another order? (y/n): ")
		again, _ := reader.ReadString('\n')
		again = strings.TrimSpace(strings.ToLower(again))

		if again == "y" {
			handleCategory(category, reader, order)
			return
		} else if again == "n" {
			utils.Clear()
			return
		} else {
			fmt.Println("Please choose 'y' or 'n'.")
			continue
		}
	}
}

func CafeName() {
	fmt.Println("==================================")
	fmt.Println("||           TEMU DEKA          ||")
	fmt.Println("==================================")
}
