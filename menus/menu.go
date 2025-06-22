package menus

import (
	"fmt"
	"strings"
)

type ProductItem interface {
	FormatPrice() string
	Display() string
	DisplayWithCategory() string
}

type OrderProduct struct {
	Name     string
	Price    int
	Quantity int
}

type Product struct {
	Name, Category string
	Price          int
	Rating         float32
}

type ProductCategory struct {
	Name string
}

type User struct {
	Username string
	Password string
}

func (pr Product) FormatPrice() string {
	return FormatRupiah(pr.Price)
}
func (om OrderProduct) FormatPrice() string {
	return FormatRupiah(om.Price * om.Quantity)
}

func (pr Product) Display() string {
	return fmt.Sprintf("%s - %s", pr.Name, pr.FormatPrice())
}
func (pr Product) DisplayWithCategory() string {
	return fmt.Sprintf("%s (%s) - %s", pr.Name, pr.Category, pr.FormatPrice())
}

func (om OrderProduct) Display() string {
	return fmt.Sprintf("%s x %d - %s", om.Name, om.Quantity, om.FormatPrice())
}

var Orders []OrderProduct

var Users []User

var ListProduct = []Product{
	{Category: "Food", Name: "Nasi Ayam Rica", Price: 30000, Rating: 4.7},
	{Category: "Food", Name: "Nasi Ayam Asam Manis", Price: 32000, Rating: 4.5},
	{Category: "Food", Name: "Nasi Ayam Lada Hitam", Price: 31000, Rating: 4.6},
	{Category: "Food", Name: "Nasi Sapi Mentega", Price: 35000, Rating: 4.8},
	{Category: "Food", Name: "Nasi Dori Saus Singapura", Price: 40000, Rating: 4.4},
	{Category: "Food", Name: "Nasi Dori Saus Thailand", Price: 40000, Rating: 4.3},
	{Category: "Food", Name: "Mie Seafood", Price: 28000, Rating: 4.2},
	{Category: "Food", Name: "Ifumi Seafood", Price: 30000, Rating: 4.1},
	{Category: "Food", Name: "Kwetiaw Siram Sapi", Price: 34000, Rating: 4.6},
	{Category: "Food", Name: "Cumi Rica", Price: 34000, Rating: 4.5},
	{Category: "Drink", Name: "Iced Coffee Degayu", Price: 30000, Rating: 4.6},
	{Category: "Drink", Name: "Iced Coffee Joglo", Price: 30000, Rating: 4.5},
	{Category: "Drink", Name: "Berry Cream Soda", Price: 26000, Rating: 4.2},
	{Category: "Drink", Name: "Sunrises", Price: 25000, Rating: 4.3},
	{Category: "Drink", Name: "Matcha Latte", Price: 23000, Rating: 4.4},
	{Category: "Drink", Name: "Peanut Butter Cream", Price: 22000, Rating: 4.1},
	{Category: "Drink", Name: "Mango Bloom Tea", Price: 20000, Rating: 4.3},
	{Category: "Drink", Name: "Lychee Gum Pop", Price: 18000, Rating: 4.0},
	{Category: "Drink", Name: "Mango Juice", Price: 20000, Rating: 4.2},
	{Category: "Drink", Name: "Mineral Water", Price: 8000, Rating: 4.0},
	{Category: "Snack", Name: "French Fries", Price: 18000, Rating: 4.3},
	{Category: "Snack", Name: "Crispy Tofu", Price: 16000, Rating: 4.2},
	{Category: "Snack", Name: "Fried Banana with Chocolate", Price: 20000, Rating: 4.4},
	{Category: "Snack", Name: "Chicken Dim Sum", Price: 20000, Rating: 4.5},
	{Category: "Snack", Name: "Fried Tapioca with Rujak Sauce", Price: 18000, Rating: 4.1},
	{Category: "Snack", Name: "Grilled Bread with Chocolate", Price: 18000, Rating: 4.3},
	{Category: "Snack", Name: "Steamed Chicken Dumpling", Price: 17000, Rating: 4.4},
	{Category: "Snack", Name: "Vegetable Fritters", Price: 18000, Rating: 4.2},
	{Category: "Snack", Name: "Grilled Fish Cake", Price: 20000, Rating: 4.3},
	{Category: "Snack", Name: "Cassava with Cheese", Price: 18000, Rating: 4.1},
	{Category: "Appetizer", Name: "Caesar Salad", Price: 22000, Rating: 4.5},
	{Category: "Appetizer", Name: "Bruschetta", Price: 20000, Rating: 4.2},
	{Category: "Appetizer", Name: "Cold Cuts Platter", Price: 24000, Rating: 4.3},
	{Category: "Appetizer", Name: "Chicken Wings", Price: 22000, Rating: 4.5},
	{Category: "Appetizer", Name: "Caprese Skewers", Price: 18000, Rating: 4.0},
	{Category: "Appetizer", Name: "Fruit Salad with Mint", Price: 20000, Rating: 4.3},
	{Category: "Appetizer", Name: "Mini Sushi Roll", Price: 22000, Rating: 4.4},
	{Category: "Appetizer", Name: "Spring Rolls", Price: 16000, Rating: 4.2},
	{Category: "Appetizer", Name: "Mozzarella Sticks", Price: 19000, Rating: 4.3},
}

var ListCategory = []ProductCategory{
	{Name: "Food"},
	{Name: "Drink"},
	{Name: "Snack"},
	{Name: "Appetizer"},
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
