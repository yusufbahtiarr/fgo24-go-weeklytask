package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/utils"
	"os"
	"strings"
)

func MainMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		CafeName()
		fmt.Print(`
Choose Menu:
1. List Menu
2. Search Menu
3. Order
4. Cart
5. Checkout

0. Exit

Choose an option (0-4): `)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "0":
			os.Exit(0)
		case "1":
			utils.Clear()
			ListMenu()
		case "2":
			utils.Clear()
			Search()
		case "3":
			utils.Clear()
			Order()
		case "4":
			utils.Clear()
			Cart()
		case "5":
			utils.Clear()
			Checkout()
		default:
			utils.Clear()
			fmt.Print("Invalid options. ")
			reader.ReadString('\n')
			utils.Clear()
			continue
		}
	}
}
