package menus

import "fmt"

type MenuType string

const (
	MainMenus  MenuType = "main"
	ListMenus  MenuType = "list"
	OrderMenus MenuType = "order"
	AuthMenus  MenuType = "auth"
)

type MenuOption struct {
	Number int
	Label  string
}

type MenuConfig struct {
	Title   string
	Options []MenuOption
}

func CafeName() {
	fmt.Println("==================================")
	fmt.Println("||           TEMU DEKA          ||")
	fmt.Println("==================================")
}

func DisplayMenu(menuType MenuType) {
	configs := map[MenuType]MenuConfig{
		MainMenus: {
			Title: "Choose Menu",
			Options: []MenuOption{
				{1, "All Menu"},
				{2, "Search Menu"},
				{3, "Filter"},
				{4, "Cart"},
				{5, "Checkout"},
				{6, "Back to Auth Menu"},
				{0, "Exit"},
			},
		},
		ListMenus: {
			Title: "Menu Show All Menu & Sorting",
			Options: []MenuOption{
				{1, "Show All Menus"},
				{2, "Sort by Most Popular"},
				{3, "Sort by Cheapest Price"},
				{4, "Sort by Name (A-Z)"},
				{5, "Sort by Name (Z-A)"},
				{0, "Back to Main Menu"},
			},
		},
		OrderMenus: {
			Title: "Choose Category",
			Options: []MenuOption{
				{1, "Foods"},
				{2, "Drinks"},
				{3, "Snacks"},
				{4, "Appetizers"},
				{0, "Back to Main Menu"},
			},
		},
		AuthMenus: {
			Title: "Authenticaton",
			Options: []MenuOption{
				{1, "Login"},
				{2, "Register"},
				{0, "Exit"},
			},
		},
	}

	choice, exists := configs[menuType]
	if !exists {
		fmt.Println("Invalid menu")
		return
	}

	CafeName()
	fmt.Printf("\n%s:\n", choice.Title)
	for _, item := range choice.Options {
		fmt.Printf("%d. %s\n", item.Number, item.Label)
	}
	fmt.Printf("\nChoose an option (0-%d): ", len(choice.Options)-1)
}
