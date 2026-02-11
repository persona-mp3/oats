package websocket

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

// Reads from server in a goroutine and writes to a ServerResponse channel
func fromServer(ctx context.Context, conn *websocket.Conn) <-chan ServerResponse {
	from := make(chan ServerResponse, 100)
	var res ServerResponse
	go func() {
		defer close(from)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := conn.ReadJSON(&res); err != nil {
					fmt.Println("could not read message from server: ", err)
					return
				}

				from <- res
			}
		}
	}()
	return from
}

// Reads from stdin in a goroutine and writes to a string channel
// Later on, when a UI Framework has been integrated, it will write to
// to a Command channel instead
func fromStdin(ctx context.Context) <-chan string {
	stdin := make(chan string, 50) // we should buffer, max of 50 inputs at a time
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		defer close(stdin)

		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			case stdin <- scanner.Text():
			}
		}
	}()
	return stdin
}
