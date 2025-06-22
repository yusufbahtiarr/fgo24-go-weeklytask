package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Cart() {
	var wg sync.WaitGroup
	checkoutChan := make(chan bool)
	deleteChan := make(chan string)

	for {
		utils.Clear()
		menus.CafeName()
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetCart()
		}()
		wg.Wait()

		fmt.Printf("\n[d] Delete Item | [c] Checkout | [b] Back Menu\n")
		fmt.Print("\nChoose option : ")
		input := utils.Input()
		switch strings.ToLower(input) {
		case "c":
			wg.Add(1)
			go func() {
				defer wg.Done()
				Checkout()
				checkoutChan <- true
			}()
			<-checkoutChan
			return
		case "b":
			return
		case "d":
			fmt.Print("Choose No Item Product : ")
			deleteInput := utils.Input()
			wg.Add(1)
			go func(input string) {
				DeleteProductCart(input, &wg)
				deleteChan <- input
			}(deleteInput)
			<-deleteChan
		default:
			utils.InvalidInput()
		}
	}
}

func GetCart() {
	if len(menus.Orders) == 0 {
		fmt.Println("\nYour Cart is Empty.")
		time.Sleep(time.Second)
		MainMenu()
	}
	fmt.Println("\nYour Cart:")
	var total int
	for i, item := range menus.Orders {
		total += item.Price * item.Quantity
		fmt.Printf("%d. %s\n", i+1, item.Display())
	}
	fmt.Printf("\nTotal: %s\n", menus.FormatRupiah(total))
}

func DeleteProductCart(input string, wg *sync.WaitGroup) {
	defer wg.Done()
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || index < 1 || index > len(menus.Orders) {
		fmt.Println("\nInvalid item number.")
		time.Sleep(time.Second)
		return
	}

	index--
	menus.Orders = append(menus.Orders[:index], menus.Orders[index+1:]...)

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
