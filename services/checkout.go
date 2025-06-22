package services

import (
	"fmt"
	"go-booking-menu/menus"
	"os"
)

func Checkout() {
	menus.CafeName()
	fmt.Println("\nYour Order:")
	var total int
	for i, item := range menus.Orders {
		total += item.Price * item.Quantity
		fmt.Printf("%d. %s\n", i+1, item.Display())
	}
	fmt.Printf("\nTotal: %s\n", menus.FormatRupiah(total))
	fmt.Println("\nThanks for choosing Temu Deka!")
	os.Exit(0)
}
