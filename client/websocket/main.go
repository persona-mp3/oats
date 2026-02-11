package websocket

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/persona-mp3/client/shared"
)

var (
	MessageTypePaint = 0
	MessageTypeChat  = 1
)

func StartWebSocketProtocol(info *shared.RedirectInfo) error {
	addr := info.Url.String()
	if len(strings.ReplaceAll(addr, " ", "")) <= 5 {
		return fmt.Errorf("the server sent an invalid address %s", addr)
	}

	if info.StatusCode != http.StatusFound {
		return fmt.Errorf("status code for redirect info is not 'found' %d", info.StatusCode)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, addr, nil)
	if err != nil {
		return fmt.Errorf("could not dial ws-server: %w", err)
	}

	return Colosseum(ctx, conn)
}

// At the moment, this the main eventloop, that is responsible for the
// following:
// 1. Reading from stdin
// 2. Reading from server
// 3. Listening for os signals ie Ctrl+C
// These are all handled with channels, from their goroutines,
func Colosseum(ctx context.Context, conn *websocket.Conn) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stdin := fromStdin(ctx)
	res := fromServer(ctx, conn)

	sigInterrupt := make(chan os.Signal, 1)
	signal.Notify(sigInterrupt, os.Interrupt)
	defer signal.Stop(sigInterrupt)

	defer func() {
		fmt.Println(" [!] closing connection")
		conn.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
			// Commands are provided by the UI Framework. They tell the server
			// what to do dependent on the user input, if required. As
			// a standin, we'll use stdin for now
		case command, ok := <-stdin:
			if !ok {
				return fmt.Errorf("stdin channel has been closed! %w", ctx.Err())
			}
			processStdinCommands(command)

		case event, ok := <-res:
			if !ok {
				return fmt.Errorf("server channel has been closed! %w", ctx.Err())
			}
			handleServerResponses(&event)

		case quit := <-sigInterrupt:
			fmt.Printf(" [*] closing application, recvd: %+v\n", quit)
			return nil
		}
	}
}

type Message struct {
	From    string `json:"from"`
	Dest    string `json:"dest"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

type Friend struct {
	Name     string `json:"name"`
	LastSeen string `json:"lastSeen"`
}

type ServerResponse struct {
	MessageType int      `json:"messageType"`
	Paint       []Friend `json:"paint"`
	Body        Message  `json:"body"`
}

// Commands are provided by the UI, interpreting user commands
// to specific actions. A ChatCommand will prompt the websocket
// connection to send the ChatCommand.Message to the websocket server
func processStdinCommands(cmd any) {
	fmt.Println("new command from ui", cmd)
}

// All new server responses are treated as Events and are
// sent to the UI to handle it's content
func handleServerResponses(res *ServerResponse) {
	fmt.Println("handling server response", res)
	fmt.Printf("%s sent ", res.Body.From)
	fmt.Printf("%s\n", res.Body)
}
