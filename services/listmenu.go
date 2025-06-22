package services

import (
	"bufio"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
	"sort"
	"strings"
)

func ListMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		menus.DisplayMenu(menus.ListMenus)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		sortedProducts := make([]menus.Product, len(menus.ListProduct))
		copy(sortedProducts, menus.ListProduct)

		switch choice {
		case "1":
			p := DefaultPagination(sortedProducts)
			DisplayPagination("Show All Menu ", p, reader)

		case "2":
			sort.Slice(sortedProducts, func(i, j int) bool {
				return sortedProducts[i].Rating > sortedProducts[j].Rating
			})
			DisplayPagination("Sort Menu by Most Popular", DefaultPagination(sortedProducts), reader)

		case "3":
			sort.Slice(sortedProducts, func(i, j int) bool {
				return sortedProducts[i].Price < sortedProducts[j].Price
			})
			DisplayPagination("Sort Menu by Cheapest Price", DefaultPagination(sortedProducts), reader)

		case "4":
			sort.Slice(sortedProducts, func(i, j int) bool {
				return strings.ToLower(sortedProducts[i].Name) < strings.ToLower(sortedProducts[j].Name)
			})
			DisplayPagination("Sort Menu by Name (A-Z)", DefaultPagination(sortedProducts), reader)

		case "5":
			sort.Slice(sortedProducts, func(i, j int) bool {
				return strings.ToLower(sortedProducts[i].Name) > strings.ToLower(sortedProducts[j].Name)
			})
			DisplayPagination("Sort Menu by Name (Z-A)", DefaultPagination(sortedProducts), reader)

		case "0":
			utils.Clear()
			return

		default:
			utils.Clear()
			return
		}

	}
}
