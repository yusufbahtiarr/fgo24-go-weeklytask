package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("\n\n\n")
	}
}

func Pause(reader *bufio.Reader) {
	fmt.Print("\nPress Enter to back...")
	reader.ReadString('\n')
}
