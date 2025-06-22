package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strings"
	"time"
)

type Pagination struct {
	Items        []menus.Product
	CurrentPage  int
	ItemsPerPage int
}

func NewPagination(items []menus.Product, perPage int) *Pagination {
	return &Pagination{
		Items:        items,
		CurrentPage:  0,
		ItemsPerPage: perPage,
	}
}

func DefaultPagination(items []menus.Product) *Pagination {
	return NewPagination(items, 5)
}

func DisplayPagination(title string, p *Pagination, reader *bufio.Reader) {
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

		fmt.Println("\n[N] Next page | [P] Previous Page | [B] Back ")
		fmt.Print("Choose an option : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

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
			fmt.Println("Invalid input.")
			time.Sleep(time.Second)
		}
	}
}
