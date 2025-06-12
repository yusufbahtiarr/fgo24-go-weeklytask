package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/utils"
	"os"
	"strings"
)


func MainMenu(){
	reader := bufio.NewReader(os.Stdin)
	for {
		CafeName()
		fmt.Print(`
Choose Menu:
1. List Menu
2. Order
3. Cart
4. Checkout

0. Exit

Choose an option (0-4)
`)
	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)

	switch choice {
	case "0":
		os.Exit(0)
	case "1":
		utils.Clear()
		Order()
	case "2":
		utils.Clear()
		Order()
	case "3":
		utils.Clear()
		Cart()
	case "4":
		utils.Clear()
		Checkout()
	default:
		utils.Clear()
		return
	}
	}
}