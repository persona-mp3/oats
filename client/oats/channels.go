package oats

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

func fromStdinCh(ctx context.Context) <-chan string {
	stdin := make(chan string, 1)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		defer close(stdin)

		fmt.Printf("  [*] ")
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

func fromServerCh(ctx context.Context, conn *websocket.Conn) <-chan ServerResponse {
	serverMsgCh := make(chan ServerResponse, 1)
	var msg ServerResponse
	go func() {
		defer close(serverMsgCh)

		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return

			default:
				err := conn.ReadJSON(&msg)
				if err != nil {
					fmt.Println(" readjson error:", err)
					return
				}

				serverMsgCh <- msg
			}

		}
	}()
	return serverMsgCh
}
