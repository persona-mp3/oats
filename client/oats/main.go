package oats

import (
	"bufio"
	"context"
	"fmt"
	// "log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/persona-mp3/client/common"
)

func StartProtocol(info *common.RedirectInfo) error {
	serverAddr := info.Location.String()
	if len(serverAddr) < 5 {
		return fmt.Errorf(" malformed server address: %s", serverAddr)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, serverAddr, nil)
	if err != nil {
		return fmt.Errorf(" could not dial wsserver: %w", err)
	}

	return colosseum(ctx, conn)
}

func colosseum(ctx context.Context, conn *websocket.Conn) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	defer func() {
		fmt.Println(" [!] closing application")
		conn.Close()
	}()

	interruptSigCh := make(chan os.Signal, 1)
	signal.Notify(interruptSigCh, os.Interrupt)
	defer signal.Stop(interruptSigCh)

	fromStdin := fromStdinCh(ctx)
	fromServer := fromServerCh(ctx, conn)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(" [*colosseum] some sibling got done")
			return ctx.Err()

		case msg, ok := <-fromStdin:
			if !ok {
				return fmt.Errorf(" stdinchannel closed")
			}

			if msg == common.QuitChat {
				return nil
			}
			// fmt.Println("  [*stdin] -> ", msg)
			if err := handleStdInEvts(ctx, msg, conn, fromStdin); err != nil {
				return err
			}

		case serverRes, ok := <-fromServer:
			if !ok {
				return fmt.Errorf(" connection with server has been closed")
			}
			fmt.Printf(" [#] %+v\n", serverRes)

		case sig := <-interruptSigCh:
			fmt.Printf(" [*colosseum] recvd deadly signal %v\n", sig)
			return nil
			// return handleCleanUp(ctx, conn)
		}
	}

}

//	func handleCleanUp(ctx context.Context, conn *websocket.Conn) error {
//		select {
//		case <-ctx.Done():
//			return ctx.Err()
//		default:
//			return conn.Close()
//		}
//		// return nil
//	}
func handleStdInEvts(ctx context.Context, val string, conn *websocket.Conn, stdinCh <-chan string) error {
	switch val {
	case "i":
		fmt.Println("   * chat initiated  *")
		return startChat(ctx, conn, stdinCh)
	default:
		fmt.Println("  unrecognised command")
	}
	return nil

}

func startChat(ctx context.Context, conn *websocket.Conn, stdin <-chan string) error {
	testMsg := &MessageJson{
		Dest:    "node_server",
		From:    "go_client",
		Time:    time.Now().String(),
		Message: "testing writing channel",
	}

	fmt.Printf("  [*i] ")
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case chat, ok := <-stdin:
			if !ok {
				return nil
			}

			testMsg.Message = chat
			if err := conn.WriteJSON(&testMsg); err != nil {
				return fmt.Errorf(" write_jsone_rror: %w", err)
			}
			fmt.Println(" chat sent...")
			return nil
		}

	}

}

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

type MessageJson struct {
	Dest    string `json:"dest"`
	From    string `json:"from"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

func fromServerCh(ctx context.Context, conn *websocket.Conn) <-chan MessageJson {
	serverMsgCh := make(chan MessageJson, 1)
	var msg MessageJson
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
