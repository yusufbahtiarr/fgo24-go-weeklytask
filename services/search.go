package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"strings"
	"sync"
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

		var wg sync.WaitGroup
		resultChan := make(chan []menus.Product)

		wg.Add(1)

		go ResultSearch(input, menus.ListProduct, resultChan, &wg)

		fmt.Printf("\nSearching...")
		time.Sleep(2 * time.Second)

		results := <-resultChan
		close(resultChan)

		if len(results) == 0 {
			fmt.Printf("\nNo product found.\n")
			time.Sleep(time.Second)
			return
		}
		DisplayPagination(fmt.Sprintf("Search Results for '%s'", input), DefaultPagination(results, CartServices.GetCarts()))
		return
	}
}

func ResultSearch(keyword string, items []menus.Product, resultChan chan<- []menus.Product, wg *sync.WaitGroup) {
	defer wg.Done()

	var results []menus.Product
	keylower := strings.ToLower(keyword)

	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), keylower) {
			results = append(results, item)
		}
	}

	resultChan <- results
}
