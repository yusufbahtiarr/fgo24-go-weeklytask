package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strings"
)

type Pagination struct {
	Items        []menus.Menu
	CurrentPage  int
	ItemsPerPage int
}

func NewPagination(items []menus.Menu, perPage int) *Pagination {
	return &Pagination{
		Items:        items,
		CurrentPage:  0,
		ItemsPerPage: perPage,
	}
}

// Get total number of pages
func (p *Pagination) TotalPages() int {
	if len(p.Items) == 0 || p.ItemsPerPage <= 0 {
		return 1
	}
	return (len(p.Items) + p.ItemsPerPage - 1) / p.ItemsPerPage
}

func (p *Pagination) CurrentItems() []menus.Menu {
	start := p.CurrentPage * p.ItemsPerPage
	end := start + p.ItemsPerPage

	if start > len(p.Items) {
		return []menus.Menu{}
	}
	if end > len(p.Items) {
		end = len(p.Items)
	}
	return p.Items[start:end]
}

func (p *Pagination) Next() {
	if p.CurrentPage < p.TotalPages()-1 {
		p.CurrentPage++
	}
}

func (p *Pagination) Previous() {
	if p.CurrentPage > 0 {
		p.CurrentPage--
	}
}

func DefaultPagination(items []menus.Menu) *Pagination {
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
		CafeName()
		fmt.Printf("\n%s (Page %d of %d)\n\n", title, p.CurrentPage+1, totalPages)

		start := p.CurrentPage * p.ItemsPerPage
		end := start + p.ItemsPerPage
		if end > totalItems {
			end = totalItems
		}

		for _, item := range p.Items[start:end] {
			fmt.Printf("- %-32s | %-10s | %s\n", item.Name, item.Category, item.FormatPrice())
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
		}
	}
}
