package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/data"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
)

func Cart() {
	reader := bufio.NewReader(os.Stdin)
	CafeName()
	fmt.Println("\nYour Order:")
	var total int
	for i, item := range data.Orders {
		total += item.Price * item.Quantity
		fmt.Printf("%d. %s\n", i+1, item.DisplayOrder())
	}
	fmt.Printf("\nTotal: %s\n", menus.FormatRupiah(total))
	fmt.Print("\nPress Enter to go back.. ")
	reader.ReadString('\n')
	utils.Clear()
	MainMenu()
}
