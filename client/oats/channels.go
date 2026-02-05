package oats

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

// Reads directly from stdin using bufio.NewScanner.
//
// This should be the only source for reading from the os.Stdin.
func readStdin() <-chan string {
	stdinCh := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" [*] ")

	go func() {
		defer close(stdinCh)
		for scanner.Scan() {
			msg := scanner.Text()
			stdinCh <- msg
		}
	}()

	return stdinCh
}

// Closes the connection to the receiving socket,
// any errors with end in a Fatal
func closeConnection(conn *websocket.Conn) {
	fmt.Println("  closing application")
	if err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	); err != nil {
		log.Fatalf(" could not close normally:\n  %s", err)
	}
	os.Exit(0)
}
