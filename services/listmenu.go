package services

import (
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"sort"
	"strings"
)

func ListMenu() {
	for {
		utils.Clear()
		menus.DisplayMenu(menus.ListMenus)
		input := utils.Input()

		sortedProducts := make([]menus.Product, len(menus.ListProduct))
		copy(sortedProducts, menus.ListProduct)

		switch input {
		case "1":
			ShowAllMenu(sortedProducts, CartServices.GetCarts())
		case "2":
			SortMostPopular(sortedProducts, CartServices.GetCarts())
		case "3":
			SortCheapestProduct(sortedProducts, CartServices.GetCarts())
		case "4":
			SortMostExpensiveProduct(sortedProducts, CartServices.GetCarts())
		case "5":
			SortMenuAscending(sortedProducts, CartServices.GetCarts())
		case "6":
			SortMenuDescending(sortedProducts, CartServices.GetCarts())
		case "0":
			utils.Clear()
			return
		default:
			utils.InvalidInput()
			return
		}
	}
}

func ShowAllMenu(sortedProducts []menus.Product, order *[]Cart) {
	DisplayPagination("Show All Menu ", DefaultPagination(sortedProducts, order))
}

func SortMostPopular(sortedProducts []menus.Product, order *[]Cart) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Rating > sortedProducts[j].Rating
	})
	DisplayPagination("Sort Menu by Most Popular", DefaultPagination(sortedProducts, order))
}

func SortCheapestProduct(sortedProducts []menus.Product, order *[]Cart) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Price < sortedProducts[j].Price
	})
	DisplayPagination("Sort Menu by Cheapest Price", DefaultPagination(sortedProducts, order))
}

func SortMostExpensiveProduct(sortedProducts []menus.Product, order *[]Cart) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Price > sortedProducts[j].Price
	})
	DisplayPagination("Sort Menu by Cheapest Price", DefaultPagination(sortedProducts, order))
}
func SortMenuAscending(sortedProducts []menus.Product, order *[]Cart) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return strings.ToLower(sortedProducts[i].Name) < strings.ToLower(sortedProducts[j].Name)
	})
	DisplayPagination("Sort Menu by Name (A-Z)", DefaultPagination(sortedProducts, order))
}
func SortMenuDescending(sortedProducts []menus.Product, order *[]Cart) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return strings.ToLower(sortedProducts[i].Name) > strings.ToLower(sortedProducts[j].Name)
	})
	DisplayPagination("Sort Menu by Name (Z-A)", DefaultPagination(sortedProducts, order))
}
