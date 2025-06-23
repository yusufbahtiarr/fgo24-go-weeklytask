package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strconv"
	"time"
)

type Pagination struct {
	Items        []menus.Product
	Cart         *[]Cart
	CurrentPage  int
	ItemsPerPage int
}

func NewPagination(items []menus.Product, perPage int, cart *[]Cart) *Pagination {
	return &Pagination{
		Items:        items,
		Cart:         cart,
		CurrentPage:  0,
		ItemsPerPage: perPage,
	}
}

func DefaultPagination(items []menus.Product, order *[]Cart) *Pagination {
	return NewPagination(items, 5, order)
}

func DisplayPagination(title string, p *Pagination) {
	totalItems := len(p.Items)
	totalPages := (totalItems + p.ItemsPerPage - 1) / p.ItemsPerPage

	for {
		if p.CurrentPage < 0 {
			p.CurrentPage = 0
		}
		if p.CurrentPage >= totalPages {
			p.CurrentPage = totalPages - 1
		}

		utils.Clear()
		menus.CafeName()
		fmt.Printf("\n%s (Page %d of %d)\n\n", title, p.CurrentPage+1, totalPages)

		start := p.CurrentPage * p.ItemsPerPage
		end := start + p.ItemsPerPage
		if end > totalItems {
			end = totalItems
		}

		for i, item := range p.Items[start:end] {
			fmt.Printf("%d. %s \n", start+i+1, item.DisplayWithCategory())
		}

		fmt.Println("\n[N] Next page | [P] Previous Page | [No_Item] Add To Cart | [B] Back ")
		fmt.Print("Choose an option : ")
		input := utils.Input()

		switch input {
		case "n", "N":
			if p.CurrentPage < totalPages-1 {
				p.CurrentPage++
			}
		case "p", "P":
			if p.CurrentPage > 0 {
				p.CurrentPage--
			}
		case "b", "B":
			utils.Clear()
			return
		default:
			index, err := strconv.Atoi(input)
			if err != nil || index < 1 || index > len(p.Items) {
				utils.InvalidInput()
				continue
			}

			selectedItem := p.Items[index-1]

			fmt.Printf("Enter quantity for %s: ", selectedItem.Name)
			qtyInput := utils.Input()
			qty, err := strconv.Atoi(qtyInput)
			if err != nil || qty < 1 {
				fmt.Println("Invalid quantity.")
				time.Sleep(time.Second)
				continue
			}

			AddToCart(p.Cart, selectedItem, qty)
			time.Sleep(1 * time.Second)
		}
	}
}
