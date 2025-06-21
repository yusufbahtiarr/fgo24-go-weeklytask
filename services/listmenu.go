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

		sortedMenus := make([]menus.Menu, len(menus.ListMenu))
		copy(sortedMenus, menus.ListMenu)

		switch choice {
		case "1":
			p := DefaultPagination(sortedMenus)
			DisplayPagination("Show All Menu ", p, reader)

		case "2":
			sort.Slice(sortedMenus, func(i, j int) bool {
				return sortedMenus[i].Rating > sortedMenus[j].Rating
			})
			DisplayPagination("Sort Menu by Most Popular", DefaultPagination(sortedMenus), reader)

		case "3":
			sort.Slice(sortedMenus, func(i, j int) bool {
				return sortedMenus[i].Price < sortedMenus[j].Price
			})
			DisplayPagination("Sort Menu by Cheapest Price", DefaultPagination(sortedMenus), reader)

		case "4":
			sort.Slice(sortedMenus, func(i, j int) bool {
				return strings.ToLower(sortedMenus[i].Name) < strings.ToLower(sortedMenus[j].Name)
			})
			DisplayPagination("Sort Menu by Name (A-Z)", DefaultPagination(sortedMenus), reader)

		case "5":
			sort.Slice(sortedMenus, func(i, j int) bool {
				return strings.ToLower(sortedMenus[i].Name) > strings.ToLower(sortedMenus[j].Name)
			})
			DisplayPagination("Sort Menu by Name (Z-A)", DefaultPagination(sortedMenus), reader)

		case "0":
			utils.Clear()
			return

		default:
			utils.Clear()
			return
		}

	}
}
