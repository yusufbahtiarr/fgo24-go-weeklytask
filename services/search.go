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
		CafeName()
		fmt.Println("Search Menu by Name ")

		fmt.Print("Input keyword : ")
		input, _ := reader.ReadString('\n')
		keyword := strings.ToLower(strings.TrimSpace(input))

		if keyword == "" {
			fmt.Println("\nKeyword cannot be empty.")
			fmt.Print("Press Enter to try again...")
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
		fmt.Println("\nSearch Results:")
		if len(results) == 0 {
			fmt.Println("No menu items found with the given keyword.")
		} else {
			for _, item := range results {
				fmt.Printf("- %s (%s) - %s\n", item.Name, item.Category, item.FormatPrice())
			}
		}

		fmt.Print("\nPress Enter to search again or type 'back' to return: ")
		inputopt, _ := reader.ReadString('\n')
		inputopt = strings.TrimSpace(strings.ToLower(inputopt))
		utils.Clear()

		if inputopt == "back" {
			utils.Clear()
			break
		}
	}
}
