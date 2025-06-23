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

type CartService interface {
	AddToCart(selected menus.Product, qty int)
	GetCart()
	DeleteProductCart(index int)
	GetCarts() *[]Cart
}

type Cart struct {
	Name     string
	Price    int
	Quantity int
}

type Carts struct {
	carts []Cart
}

var CartServices CartService = &Carts{
	carts: []Cart{},
}

func (c *Carts) GetCarts() *[]Cart {
	return &c.carts
}

func (om Cart) FormatPrice() string {
	return menus.FormatRupiah(om.Price * om.Quantity)
}
func (om Cart) Display() string {
	return fmt.Sprintf("%s x %d - %s", om.Name, om.Quantity, om.FormatPrice())
}

func CartView() {
	var wg sync.WaitGroup
	checkoutChan := make(chan bool)

	for {
		utils.Clear()
		menus.CafeName()

		wg.Add(1)
		go func() {
			defer wg.Done()
			CartServices.GetCart()
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
			inputProduct, err := strconv.Atoi(strings.TrimSpace(deleteInput))
			if err != nil {
				fmt.Println("\nInvalid input. Must be a number.")
				time.Sleep(time.Second)
				break
			}
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				CartServices.DeleteProductCart(idx)
			}(inputProduct)

			wg.Wait()
		default:
			utils.InvalidInput()
		}
	}
}

func (c *Carts) GetCart() {
	if len(c.carts) == 0 {
		fmt.Println("\nYour Cart is Empty.")
		time.Sleep(time.Second)
		MainMenu()
	}
	fmt.Println("\nYour Cart:")
	var total int
	for i, item := range c.carts {
		total += item.Price * item.Quantity
		fmt.Printf("%d. %s\n", i+1, item.Display())
	}
	fmt.Printf("\nTotal: %s\n", menus.FormatRupiah(total))
}

func (c *Carts) DeleteProductCart(index int) {
	if index < 1 || index > len(c.carts) {
		fmt.Println("\nInvalid item number.")
		time.Sleep(time.Second)
		return
	}
	index--
	c.carts = append(c.carts[:index], c.carts[index+1:]...)
	fmt.Println("\nItem has been removed from your cart.")
	time.Sleep(time.Second)
}

func (c *Carts) AddToCart(selected menus.Product, qty int) {
	for i, item := range c.carts {
		if item.Name == selected.Name {
			c.carts[i].Quantity += qty
			fmt.Printf("\nAdded %d more of %s to your cart. Total: %d\n", qty, selected.Name, c.carts[i].Quantity)
			return
		}
	}
	c.carts = append(c.carts, Cart{
		Name:     selected.Name,
		Price:    selected.Price,
		Quantity: qty,
	})
	fmt.Printf("\nSuccessfully added %d x %s to your cart.\n", qty, selected.Name)
}
