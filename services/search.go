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
		utils.Clear()
		menus.CafeName()
		fmt.Printf("\nSearch Product by Name \n")

		fmt.Printf("\nEnter the product you're looking for: ")
		input := utils.Input()

		if input == "" {
			fmt.Printf("\nKeyword cannot be empty.\n")
			time.Sleep(time.Second)
			continue
		}

		resultChan := make(chan []menus.Product)

		go func(keyword string, ch chan<- []menus.Product) {
			var found []menus.Product
			lowered := strings.ToLower(keyword)

			for _, item := range menus.ListProduct {
				if strings.Contains(strings.ToLower(item.Name), lowered) {
					found = append(found, item)
				}
			}
			ch <- found
		}(input, resultChan)

		fmt.Printf("\nSearching...")
		time.Sleep(2 * time.Second)

		results := <-resultChan
		close(resultChan)

		if len(results) == 0 {
			fmt.Printf("\nNo product found.\n")
			time.Sleep(time.Second)
			return
		}
		DisplayPagination(fmt.Sprintf("Search Results for '%s'", input), DefaultPagination(results))
		return
	}
}
