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
		} else {
			const itemPerPage = 5
			totalItems := len(results)
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
				fmt.Printf("\nSearch Results '%s' (Page %d of %d)\n\n", keyword, currentPage+1, totalPages)

				start := currentPage * itemPerPage
				end := start + itemPerPage
				if end > totalItems {
					end = totalItems
				}
				for _, item := range results[start:end] {
					fmt.Printf("- %s (%s) - %s\n", item.Name, item.Category, item.FormatPrice())
				}

				fmt.Println("\n[N] Next page | [P] Previous Page | [B] Back")
				fmt.Printf("\nChoose an option : ")
				input, _ = reader.ReadString('\n')
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
					fmt.Println("Invalid options.")
				}
			}
		}
	}
}
