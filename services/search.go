package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
	"strings"
)

func Search() {
	reader := bufio.NewReader(os.Stdin)
	for {
		menus.CafeName()
		fmt.Printf("\nSearch Menu by Name \n")

		fmt.Printf("\nInput keyword : ")
		input, _ := reader.ReadString('\n')
		keyword := strings.ToLower(strings.TrimSpace(input))

		if keyword == "" {
			fmt.Printf("\nKeyword cannot be empty.\n")
			fmt.Printf("\nPress Enter to try again...")
			reader.ReadString('\n')
			utils.Clear()
			continue
		}

		var results []menus.Menu
		for _, item := range menus.ListMenu {
			if strings.Contains(strings.ToLower(item.Name), keyword) {
				results = append(results, item)
			}
		}

		if len(results) == 0 {
			fmt.Printf("\nNo menu items found with the given keyword.\n")
			utils.GoBack(reader)
			utils.Clear()
			return
		}
		p := DefaultPagination(results)
		DisplayPagination(fmt.Sprintf("Search Results for '%s'", keyword), p, reader)
		return
	}
}
