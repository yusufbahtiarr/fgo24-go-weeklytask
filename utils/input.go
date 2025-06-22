package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func InvalidInput() {
	fmt.Printf("\nInvalid input. Please try again.\n")
	time.Sleep(time.Second)
}

func Exit() {
	fmt.Printf("\nExit Program | TEKA DEGA...\n")
	time.Sleep(time.Second)
	os.Exit(0)
}
