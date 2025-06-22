package services

import (
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"sort"
	"strings"
)

func ListMenu() {
	for {
		menus.DisplayMenu(menus.ListMenus)
		input := utils.Input()

		sortedProducts := make([]menus.Product, len(menus.ListProduct))
		copy(sortedProducts, menus.ListProduct)

		switch input {
		case "1":
			ShowAllMenu(sortedProducts)
		case "2":
			SortMostPopular(sortedProducts)
		case "3":
			SortCheapestProduct(sortedProducts)
		case "4":
			SortMenuAscending(sortedProducts)
		case "5":
			SortMenuDescending(sortedProducts)
		case "0":
			utils.Clear()
			return
		default:
			utils.InvalidInput()
			return
		}
	}
}

func ShowAllMenu(sortedProducts []menus.Product) {
	DisplayPagination("Show All Menu ", DefaultPagination(sortedProducts))
}

func SortMostPopular(sortedProducts []menus.Product) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Rating > sortedProducts[j].Rating
	})
	DisplayPagination("Sort Menu by Most Popular", DefaultPagination(sortedProducts))
}

func SortCheapestProduct(sortedProducts []menus.Product) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Price < sortedProducts[j].Price
	})
	DisplayPagination("Sort Menu by Cheapest Price", DefaultPagination(sortedProducts))
}
func SortMenuAscending(sortedProducts []menus.Product) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return strings.ToLower(sortedProducts[i].Name) < strings.ToLower(sortedProducts[j].Name)
	})
	DisplayPagination("Sort Menu by Name (A-Z)", DefaultPagination(sortedProducts))
}
func SortMenuDescending(sortedProducts []menus.Product) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return strings.ToLower(sortedProducts[i].Name) > strings.ToLower(sortedProducts[j].Name)
	})
	DisplayPagination("Sort Menu by Name (Z-A)", DefaultPagination(sortedProducts))
}
