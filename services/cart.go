package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
)

func Cart() {
	reader := bufio.NewReader(os.Stdin)
	menus.CafeName()
	fmt.Println("\nYour Order:")
	var total int
	for i, item := range menus.Orders {
		total += item.Price * item.Quantity
		fmt.Printf("%d. %s\n", i+1, item.Display())
	}
	fmt.Printf("\nTotal: %s\n", menus.FormatRupiah(total))
	fmt.Print("\nPress Enter to go back.. ")
	reader.ReadString('\n')
	utils.Clear()
	MainMenu()
}
