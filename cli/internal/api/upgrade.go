package api

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// not supposed to be the api package but for now temp
func ClearScreen() {
	screen := exec.Command("clear")
	screen.Stdout = os.Stdout
	screen.Run()
}

// from gorilla's documentation
func (req *Req) requestUpgrade(info *RedirectInfo) error {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := info.url

	wssUrl := u.String()

	c, _, err := websocket.DefaultDialer.Dial(wssUrl, nil)
	if err != nil {
		return fmt.Errorf("couldn't dial wss: %w", err)
	}

	ClearScreen()
	fmt.Printf("\n\n connection opened\n\n")

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("\n [recv]:  %s\n", msg)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return nil
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()+"  welcome to the oats domain"))
			if err != nil {
				log.Println("write:", err)
				return err
			}
		case <-interrupt:
			log.Println("disconnecting from server")
			err := c.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)
			if err != nil {
				return fmt.Errorf("[err] in disconnecting cleanly: %w", err)
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}

			return nil
		}
	}

}
