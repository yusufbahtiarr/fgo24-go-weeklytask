package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strconv"
	"strings"
	"time"
)

func Cart() {
	for {
		utils.Clear()
		menus.CafeName()
		if len(menus.Orders) == 0 {
			fmt.Println("\nYour Cart is Empty:")
			time.Sleep(time.Second)
			return
		}
		fmt.Println("\nYour Cart:")
		var total int
		for i, item := range menus.Orders {
			total += item.Price * item.Quantity
			fmt.Printf("%d. %s\n", i+1, item.Display())
		}
		fmt.Printf("\nTotal: %s\n", menus.FormatRupiah(total))
		fmt.Printf("\n[d] Delete Item | [c] Checkout | [b] Back Menu\n")
		fmt.Print("\nChoose option : ")
		input := utils.Input()
		switch strings.ToLower(input) {
		case "c":
			Checkout()
		case "b":
			return
		case "d":
			fmt.Print("Choose No Item Product : ")
			input := utils.Input()
			DeleteProductCart(input)
		default:
			utils.InvalidInput()
		}
	}
}

func DeleteProductCart(input string) {
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || index < 1 || index > len(menus.Orders) {
		fmt.Println("\nInvalid item number.")
		time.Sleep(time.Second)
		return
	}

	done := make(chan bool)

	go func() {
		index--
		menus.Orders = append(menus.Orders[:index], menus.Orders[index+1:]...)
		done <- true
	}()

	<-done

	fmt.Println("\nItem has been removed from your cart.")
	time.Sleep(time.Second)
}

func AddToCart(order *[]menus.OrderProduct, selectedItem menus.Product, qty int) {
	for i, item := range *order {
		if item.Name == selectedItem.Name {
			(*order)[i].Quantity += qty
			fmt.Printf("\nAdded %d more of %s to your cart. Total: %d\n", qty, selectedItem.Name, (*order)[i].Quantity)
			return
		}
	}

	*order = append(*order, menus.OrderProduct{
		Name:     selectedItem.Name,
		Price:    selectedItem.Price,
		Quantity: qty,
	})

	fmt.Printf("\nSuccessfully added %d x %s to your order.\n", qty, selectedItem.Name)
}
