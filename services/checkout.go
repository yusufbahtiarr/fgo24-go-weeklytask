package services

import (
	"fmt"
	"go-booking-menu/menus"
	"go-booking-menu/utils"
	"os"
	"sync"
	"time"
)

func Checkout() {
	utils.Clear()
	menus.CafeName()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		GetCart()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("\nProcessing payment...")
		time.Sleep(2 * time.Second)
		fmt.Println("\nPayment completed!")
	}()

	wg.Wait()

	fmt.Printf("\nThanks for choosing Temu Deka!")
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
