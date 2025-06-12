package services

import (
	"fmt"
	"go-booking-menu/data"
	"go-booking-menu/menus"
	"os"
)

func Checkout(){
	CafeName()
	fmt.Println("\nYour Order:")
	var total int
	for i, item := range data.Orders {
		total += item.Price * item.Quantity
		 fmt.Printf("%d. %s\n", i+1, item.DisplayOrder()) 
	}
	fmt.Printf("\nTotal: %s\n", menus.FormatRupiah(total))
	fmt.Println("\nThanks for choosing Temu Degayu!")
	os.Exit(0)
}