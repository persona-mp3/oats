package oats

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/persona-mp3/client/common"
	"github.com/persona-mp3/client/utils"
)

func BeginWSProtocol(addr string) error {
	utils.ClearScreen()

	if len(addr) == 0 {
		return fmt.Errorf(" invalid address provided for ws_server %s", addr)
	}

	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	log.Println(" connection opened successfully")
	if err != nil {
		return nil
	}

	if err := Collesium(conn); err != nil {
		return err
	}

	return nil
}

// func startChat(conn *websocket.Conn) error {
// 	interrupt := make(chan os.Signal, 1)
// 	signal.Notify(interrupt, os.Interrupt)
//
// 	sendMsg := make(chan os.File)
//
// 	done := make(chan struct{})
//
// 	for {
// 		go func() {
// 			defer close(done)
// 			err := readMessage(conn)
// 			log.Println(err)
// 		}()
//
// 		go writeMessage(conn)
// 		break
// 	}
//
// 	return nil
// }

func Collesium(conn *websocket.Conn) error {
	userEvent := GetUserEvents()
	switch userEvent.name {
	case common.ChatEvent:
		userEvent.Chat(conn)
	default:
		fmt.Println(" unknown event")
	}

	return nil
}

type ChatMsg struct{}

func (event *Event) Chat(conn *websocket.Conn) error {
	fmt.Println(" [:chat] sending message")
	err := conn.WriteMessage(websocket.TextMessage, []byte(event.value))
	if err != nil {
		log.Println(err)
	}
	return nil
}

type Message struct {
	Msg     string `json:"msg"`
	From    string `json:"from"`
	Media   string `json:"media"`
	Isgroup bool   `json:"isGroup"`
}

func readMessage(conn *websocket.Conn) error {
	msgType, msg, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf(" could not read msg: %w", err)
	}

	log.Println(" [msg_type]: ", msgType)
	log.Println(" *[new_msg]")

	fmt.Printf(`
			%s

	`, msg)

	return nil
}

// Reads from stdin and parses /in to provide
// and input for reading msg to send
// Not sure yet, but we might set a channel to listen for it
func waitForFlagInput() <-chan string {
	userFlags := make(chan string)

	go func() {
		defer close(userFlags)

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			userFlags <- scanner.Text()
		}
	}()

	return userFlags
}

type Event struct {
	value string
	name  int
}

func GetUserEvents() Event {
	readCh := waitForFlagInput()

	var evt Event
	for cmd := range readCh {
		switch cmd {
		case ":in":
			evt.value = ReadInputFromCli()
			evt.name = common.InEvent
			evt.WriteMessage()
			return evt

		case ":search":
			fmt.Println("searching event")
			evt.Search()

		default:
			fmt.Println(" unsupported:", cmd)
		}

	}

	return evt
}

func (e *Event) WriteMessage() {

}

func (e *Event) Search() {}
