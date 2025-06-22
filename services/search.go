package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strings"
	"time"
)

func Search() {
	for {
		menus.CafeName()
		fmt.Printf("\nSearch Product by Name \n")

		fmt.Printf("\nEnter the product you're looking for:  : ")
		input := utils.Input()

		if input == "" {
			fmt.Printf("\nKeyword cannot be empty.\n")
			time.Sleep(time.Second)
			continue
		}

		var results []menus.Product
		for _, item := range menus.ListProduct {
			if strings.Contains(strings.ToLower(item.Name), input) {
				results = append(results, item)
			}
		}

		if len(results) == 0 {
			fmt.Printf("\nNo product found.\n")
			time.Sleep(time.Second)
			utils.Clear()
			return
		}
		p := DefaultPagination(results)
		DisplayPagination(fmt.Sprintf("Search Results for '%s'", input), p)
		return
	}
}
