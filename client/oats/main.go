package oats

import (
	"bufio"
	"context"
	"fmt"

	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/persona-mp3/client/common"
	"github.com/persona-mp3/client/internal/renderer"
	"github.com/persona-mp3/client/utils"
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

	utils.ClearScreen()

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
			// fmt.Printf(" [#] %+v\n", serverRes)
			if err := processServerResponse(&serverRes); err != nil {
				return err
			}

		case sig := <-interruptSigCh:
			fmt.Printf(" [*colosseum] recvd deadly signal %v\n", sig)
			return nil
		}
	}

}

var ContentPaintType = 0
var ChatMessageType = 1

type ServerResponse struct {
	// Represents how the message should be processed
	// 0 means FirstContentPaint and the Paint field can be accessed
	// 1 means it's a normal Message ie came from a sender or the server
	MessageType int `json:"messageType"`

	Body common.Message `json:"body"`

	// This field is optional as it only happens on the inital connection
	Paint []common.Friend `json:"paint"`
}

type Friend struct {
	Name     string `json:"name"`
	LastSeen string `json:"lastSeen"`
}

//	type Message struct {
//		Dest    string `json:"dest"`
//		From    string `json:"from"`
//		Time    string `json:"time"`
//		Message string `json:"message"`
//	}
func processServerResponse(res *ServerResponse) error {
	responseType := res.MessageType

	switch responseType {
	case ContentPaintType:
		return renderer.RenderContentPaint(&res.Paint)

	case ChatMessageType:
		renderer.RenderChatMessage(&res.Body)
		return nil
	}

	return nil
}

func handleStdInEvts(ctx context.Context, val string, conn *websocket.Conn, stdinCh <-chan string) error {
	switch val {
	case "i":
		return startChat(ctx, conn, stdinCh)
	default:
		fmt.Println("  unrecognised command")
	}
	return nil

}

func startChat(ctx context.Context, conn *websocket.Conn, stdin <-chan string) error {
	testMsg := &common.Message{
		Dest: "node_client",
		From: "go_client",
		Time: time.Now().String(),
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
				return fmt.Errorf(" write json error: %w", err)
			}
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
