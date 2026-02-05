package oats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/persona-mp3/client/common"
)

// Begins the WSS Protocol by following the redirect
// address provided by addr.
//
// It contacts the server using the DialContext method.
func BeginOatsProtocol(addr string) error {
	if len(addr) == 0 {
		return fmt.Errorf(" invalid address provided from server: %s", addr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, addr, nil)
	if err != nil {
		return fmt.Errorf(" could not dial wsserver: \n%w", err)
	}

	if err := mainCollesium(conn); err != nil {
		return err
	}
	return nil
}

// Response for initiating stdin channel and down propagates
// and maps all values to the event handler
//
// Responsible for reading messages from the wss
func mainCollesium(conn *websocket.Conn) error {
	stdinCh := readStdin()
	go readFromServer(conn)
	for {
		val, ok := <-stdinCh
		if !ok {
			fmt.Println("stdinch has been closed?")
			break
		}

		if err := HandleEvent(conn, val, stdinCh); err != nil {
			log.Println(err)
		}

	}

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)

	go func() {
		defer close(interruptCh)

		sig := <-interruptCh
		fmt.Printf(" recvd: %v\n", sig)
		closeConnection(conn)
	}()
	return nil
}

func readFromServer(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(" could not read: ", err)
		}

		fmt.Printf(" [#]: %s\n", msg)
	}
}

type Event struct {
	Name      int
	conn      *websocket.Conn
	stdinChan <-chan string
}

func createEvent(name int, conn *websocket.Conn, stdinChan <-chan string) *Event {
	return &Event{
		Name:      name,
		conn:      conn,
		stdinChan: stdinChan,
	}
}

// Maps all events gotten from stdin channel in similarity of vim motions
//
// For example, 'i' passed from stdin would call the ChatEvent Handler
func HandleEvent(conn *websocket.Conn, val string, stdin <-chan string) error {
	switch val {
	case "i":
		evt := createEvent(common.ChatEvent, conn, stdin)
		if err := evt.ChatEvent(); err != nil {
			return err
		}
	case "q":
		closeConnection(conn)

	default:
		fmt.Println(" unsupported cmd:", val)
		fmt.Print(" [*] ")
	}
	return nil
}

type Message struct {
	From    string `json:"from"`
	Dest    string `json:"dest"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func (evt *Event) ChatEvent() error {
	var msg string
	fmt.Print(" [send-chat]  ")

	val, ok := <-evt.stdinChan
	if !ok {
		return fmt.Errorf(" stdinch closed unexpectedly")
	}
	msg = val

	fmt.Print(" [*] ")

	// content, err := parseToJson(msg)
	// if err != nil {
	// 	return err
	// }
	//
	// err := evt.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	// evt.conn.WriteJSON()
	err := evt.conn.WriteJSON(createMessage(msg))
	if err != nil {
		return fmt.Errorf("could not send %w", err)
	}
	return nil
}

func parseToJson(data any) ([]byte, error) {
	content, err := json.Marshal(data)
	if err != nil {
		return []byte{}, fmt.Errorf(" error in marsahlling: %w\n", err)
	}
	return content, nil
}

func createMessage(msg string) *Message {
	return &Message{
		From:    "master_user",
		Dest:    "ladiesman217",
		Message: msg,
		Time:    time.Now().String(),
	}
}
