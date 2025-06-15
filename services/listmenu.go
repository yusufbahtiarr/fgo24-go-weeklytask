package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
	"sort"
	"strings"
)

func ListMenu() {
	reader := bufio.NewReader(os.Stdin)
	utils.Clear()
	for {
		CafeName()
		fmt.Println(`
Menu Show All Menu & Sorting
1. Show All Menus
2. Sort by Most Popular
3. Sort by Cheapest Price
4. Sort by Name (A-Z)
5. Sort by Name (Z-A)
0. Back to Main Menu`)

		fmt.Print("\nChoose an option (0-5): ")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		sortedMenus := make([]menus.Menu, len(menus.ListMenu))
		copy(sortedMenus, menus.ListMenu)

		switch choice {
		case "1":
			utils.Clear()
			CafeName()
			displayedSortedMenus("Show All Menu ", sortedMenus, reader)
		case "2":
			sort.Slice(sortedMenus, func(i, j int) bool {
				return sortedMenus[i].Rating > sortedMenus[j].Rating
			})
			displayedSortedMenus("Sort Menu by Most Popular ", sortedMenus, reader)
		case "3":
			sort.Slice(sortedMenus, func(i, j int) bool {
				return sortedMenus[i].Price < sortedMenus[j].Price
			})
			displayedSortedMenus("Sort Menu by Cheapest Price ", sortedMenus, reader)
		case "4":
			sort.Slice(sortedMenus, func(i, j int) bool {
				return strings.ToLower(sortedMenus[i].Name) < strings.ToLower(sortedMenus[j].Name)
			})
			displayedSortedMenus("Sort Menu by Name (A-Z) ", sortedMenus, reader)

		case "5":
			sort.Slice(sortedMenus, func(i, j int) bool {
				return strings.ToLower(sortedMenus[i].Name) > strings.ToLower(sortedMenus[j].Name)
			})
			displayedSortedMenus("Sort Menu by Name (Z-A) ", sortedMenus, reader)
		case "0":
			utils.Clear()
			return
		}

	}
}

func displayedSortedMenus(title string, menus []menus.Menu, reader *bufio.Reader) {
	const itemPerPage = 5
	totalItems := len(menus)
	totalPages := (totalItems + itemPerPage - 1) / itemPerPage
	currentPage := 0

	for {
		if currentPage < 0 {
			currentPage = 0
		}
		if currentPage >= totalPages {
			currentPage = totalPages - 1
		}
		utils.Clear()
		CafeName()
		fmt.Printf("\n%s (Page %d of %d)\n\n", title, currentPage+1, totalPages)

		start := currentPage * itemPerPage
		end := start + itemPerPage
		if end > totalItems {
			end = totalItems
		}
		for _, item := range menus[start:end] {
			fmt.Printf("- %-32s | %-8s | %s\n", item.Name, item.Category, item.FormatPrice())
		}

		fmt.Println("\n[N] Next page | [P] Previous Page | [B] Back ")
		fmt.Print("Choose an option : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "n", "N":
			if currentPage < totalPages-1 {
				currentPage++
			}
		case "p", "P":
			if currentPage > 0 {
				currentPage--
			}
		case "b", "B":
			utils.Clear()
			return
		default:
			// fmt.Println("Invalid input.")
		}
	}
}
