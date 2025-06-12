package menus

import (
	"fmt"
	"strings"
)


type OrderMenu struct {
	Name string
	Price int
	Quantity int
}

type MenuItem struct {
	Name string
	Price int
	Rating float32
}

type MenuCategory struct {
	Name string
	Items []MenuItem
}

type MenuCollection struct {
	Categories []MenuCategory
}

func NewMenuCollection() *MenuCollection{
	return &MenuCollection{
		Categories: []MenuCategory{
			{
				Name: "Foods",
				Items: []MenuItem{
					{"Nasi Ayam Rica", 30000, 4.7},
					{"Nasi Ayam Asam Manis", 32000, 4.5},
					{"Nasi Ayam Lada Hitam", 31000, 4.6},
					{"Nasi Sapi Mentega", 35000, 4.8},
					{"Nasi Dori Saus Singapura", 40000, 4.4},
					{"Nasi Dori Saus Thailand", 40000, 4.3},
					{"Mie Seafood", 28000, 4.2},
					{"Ifumi Seafood", 30000, 4.1},
					{"Kwetiaw Siram Sapi", 34000, 4.6},
					{"Cumi Rica", 34000, 4.5},
				},
			},
			{
				Name: "Drinks",
				Items: []MenuItem{
					{"Iced Coffee Degayu", 30000, 4.6},
					{"Iced Coffee Joglo", 30000, 4.5},
					{"Berry Cream Soda", 26000, 4.2},
					{"Sunrises", 25000, 4.3},
					{"Matcha Latte", 23000, 4.4},
					{"Peanut Butter Cream", 22000, 4.1},
					{"Mango Bloom Tea", 20000, 4.3},
					{"Lychee Gum Pop", 18000, 4.0},
					{"Mango Juice", 20000, 4.2},
					{"Mango Juice", 20000, 4.2},
					{"Mineral Water", 8000, 4.0},
				},
			},
			{
				Name: "Snacks",
				Items: []MenuItem{
					{"French Fries", 18000, 4.3},
					{"Crispy Tofu", 16000, 4.2},
					{"Fried Banana with Chocolate", 20000, 4.4},
					{"Chicken Dim Sum", 20000, 4.5},
					{"Fried Tapioca with Spicy Rujak Sauce", 18000, 4.1},
					{"Grilled Bread with Chocolate & Cheese", 18000, 4.3},
					{"Steamed Chicken Dumpling", 17000, 4.4},
					{"Vegetable Fritters", 18000, 4.2},
					{"Grilled Fish Cake", 20000, 4.3},
					{"Cassava with Cheese", 18000, 4.1},
				},
			},
			{
				Name: "Appetizers",
				Items: []MenuItem{
					{"Caesar Salad", 22000, 4.5},
					{"Bruschetta", 20000, 4.2},
					{"Cold Cuts Platter", 24000, 4.3},
					{"Chicken Wings", 22000, 4.5},
					{"Caprese Skewers", 18000, 4.0},
					{"Fruit Salad with Mint", 20000, 4.3},
					{"Mini Sushi Roll", 22000, 4.4},
					{"Chicken Wings", 22000, 4.5},
					{"Spring Rolls", 16000, 4.2},
					{"Mozzarella Sticks", 19000, 4.3},
				},
			},
		},
	}
}

func (mi MenuItem) FormatPrice() string {
	return FormatRupiah(mi.Price)
}

func (mi MenuItem) DisplayName() string {
	return fmt.Sprintf("%s - %s", mi.Name, mi.FormatPrice())
}

func (om OrderMenu) DisplayOrder() string {
    return fmt.Sprintf("%s x %d - %s", om.Name, om.Quantity, FormatRupiah(om.Price*om.Quantity))
}

func (mc *MenuCollection) ShowMenu(){
	for _, category := range mc.Categories {
		fmt.Printf("\nMenu %s:\n", category.Name)
		for i, item:= range category.Items {
			fmt.Printf("%d. %s\n", i+1, item.DisplayName())
		}
	}
}

func FormatRupiah(value int) string {
	str := fmt.Sprintf("%d", value)
	n := len(str)

	if n <= 3 {
		return "Rp " + str
	}

	var result []string
	for i, c := range str {
		if (n-i)%3 == 0 && i != 0 {
			result = append(result, ".")
		}
		result = append(result, string(c))
	}

	return "Rp " + strings.Join(result, "")
}
