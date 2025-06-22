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
			ShowAllMenu(sortedProducts, &menus.Orders)
		case "2":
			SortMostPopular(sortedProducts, &menus.Orders)
		case "3":
			SortCheapestProduct(sortedProducts, &menus.Orders)
		case "4":
			SortMostExpensiveProduct(sortedProducts, &menus.Orders)
		case "5":
			SortMenuAscending(sortedProducts, &menus.Orders)
		case "6":
			SortMenuDescending(sortedProducts, &menus.Orders)
		case "0":
			utils.Clear()
			return
		default:
			utils.InvalidInput()
			return
		}
	}
}

func ShowAllMenu(sortedProducts []menus.Product, order *[]menus.OrderProduct) {
	DisplayPagination("Show All Menu ", DefaultPagination(sortedProducts, order))
}

func SortMostPopular(sortedProducts []menus.Product, order *[]menus.OrderProduct) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Rating > sortedProducts[j].Rating
	})
	DisplayPagination("Sort Menu by Most Popular", DefaultPagination(sortedProducts, order))
}

func SortCheapestProduct(sortedProducts []menus.Product, order *[]menus.OrderProduct) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Price < sortedProducts[j].Price
	})
	DisplayPagination("Sort Menu by Cheapest Price", DefaultPagination(sortedProducts, order))
}

func SortMostExpensiveProduct(sortedProducts []menus.Product, order *[]menus.OrderProduct) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Price > sortedProducts[j].Price
	})
	DisplayPagination("Sort Menu by Cheapest Price", DefaultPagination(sortedProducts, order))
}
func SortMenuAscending(sortedProducts []menus.Product, order *[]menus.OrderProduct) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return strings.ToLower(sortedProducts[i].Name) < strings.ToLower(sortedProducts[j].Name)
	})
	DisplayPagination("Sort Menu by Name (A-Z)", DefaultPagination(sortedProducts, order))
}
func SortMenuDescending(sortedProducts []menus.Product, order *[]menus.OrderProduct) {
	sort.Slice(sortedProducts, func(i, j int) bool {
		return strings.ToLower(sortedProducts[i].Name) > strings.ToLower(sortedProducts[j].Name)
	})
	DisplayPagination("Sort Menu by Name (Z-A)", DefaultPagination(sortedProducts, order))
}
