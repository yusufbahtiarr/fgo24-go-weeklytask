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
			continue
		case "2":
			utils.Clear()
			sort.Slice(sortedMenus, func(i, j int) bool {
				return sortedMenus[i].Rating > sortedMenus[j].Rating
			})
			CafeName()
			fmt.Printf("\nSort Menu by Most Popular: \n")
			for _, item := range sortedMenus {
				fmt.Printf("- %s (%s) - %s\n", item.Name, item.Category, item.FormatPrice())
			}
			utils.GoBack(reader)
			utils.Clear()
			continue
		case "3":
			utils.Clear()
			fmt.Printf("\nSort Menu by Cheapest Price: \n")
			sort.Slice(sortedMenus, func(i, j int) bool {
				return sortedMenus[i].Price < sortedMenus[j].Price
			})
			for _, item := range sortedMenus {
				fmt.Printf("- %s (%s) - %s\n", item.Name, item.Category, item.FormatPrice())
			}
			utils.GoBack(reader)
			utils.Clear()
			continue
		case "4":
			utils.Clear()
			fmt.Printf("\n Sort Menu by Name (A-Z):\n")
			sort.Slice(sortedMenus, func(i, j int) bool {
				return strings.ToLower(sortedMenus[i].Name) < strings.ToLower(sortedMenus[j].Name)
			})
			for _, item := range sortedMenus {
				fmt.Printf("- %s (%s) - %s\n", item.Name, item.Category, item.FormatPrice())
			}
			utils.GoBack(reader)
			utils.Clear()
			continue
		case "5":
			utils.Clear()
			fmt.Printf("\n Sort Menu by Name (Z-A):\n")
			sort.Slice(sortedMenus, func(i, j int) bool {
				return strings.ToLower(sortedMenus[i].Name) > strings.ToLower(sortedMenus[j].Name)
			})
			for _, item := range sortedMenus {
				fmt.Printf("- %s (%s) - %s\n", item.Name, item.Category, item.FormatPrice())
			}
			utils.GoBack(reader)
			utils.Clear()
			continue
		case "0":
			utils.Clear()
			return
		}

	}
}
