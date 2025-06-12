package services

import (
	"bufio"
	"fmt"
	"go-booking-menu/utils"
	"os"
)

func ListMenu(){
	reader := bufio.NewReader(os.Stdin)
	CafeName()
	fmt.Println(`
1. All Menu
2. Search Menu
3. Sorting Menu

0. Back to Main Menu`)
	fmt.Print("\nPress Enter to go back.. ")
	reader.ReadString('\n')
	utils.Clear()

}