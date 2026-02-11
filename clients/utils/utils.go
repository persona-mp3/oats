package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	screen := exec.Command("clear")
	screen.Stdout = os.Stdout
	screen.Run()
}

func ReadInputFromCli() string {
	var msg string
	fmt.Printf(" > ")
	scanner := bufio.NewScanner(os.Stdin)
	msg = scanner.Text()
	return msg
}
