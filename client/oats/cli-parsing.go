package oats

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInputFromCli() string {
	var msg string
	fmt.Printf(" > ")
	scanner := bufio.NewScanner(os.Stdin)
	msg = scanner.Text()
	return msg
}
