package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/matheusrb95/endrok/types"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/net/netutil"
	"nhooyr.io/websocket"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type client chan<- *Game

var count int

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan *Game)
)

type Player struct {
	Position rl.Vector2
	Size     rl.Vector2
}

type Game struct {
	Who     string
	Player1 Player
	Player2 Player
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	listener = netutil.LimitListener(listener, 2)

	var game Game

	s := &http.Server{
		Handler: websocketServer{Game: &game},
	}

	log.Println("Starting Server ", listener.Addr().String())

	errc := make(chan error, 1)
	go func() {
		errc <- s.Serve(listener)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	select {
	case err := <-errc:
		log.Println("Failed to serve: ", err)
	case sig := <-sigs:
		log.Println("Terminating: ", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = s.Shutdown(ctx)
	if err != nil {
		log.Println("Error shutting down server: ", err)
	}
}

type websocketServer struct {
	Game *Game
}

func (s websocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println("Error accepting websocket: ", err)
		return
	}

	ctx := context.Background()
	conn := websocket.NetConn(ctx, c, websocket.MessageBinary)
	go s.ServeNetConn(conn)
}

func (s websocketServer) ServeNetConn(conn net.Conn) {
	count++
	welcomeMsg := types.WSMessage{
		SessionID: count,
		Type:      "welcome",
		Data:      nil,
	}
	js, err := json.Marshal(welcomeMsg)
	if err != nil {
		log.Println("Error marshal: ", err)
		return
	}
	conn.Write(js)

	ch := make(chan *Game)
	go s.clientWriter(conn, ch)

	s.Game.Who = conn.RemoteAddr().String()
	ch <- s.Game
	messages <- s.Game
	entering <- ch

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing conn: ", err)
		}
	}()

	timeoutTime := 60 * time.Second
	timeout := make(chan uint8, 1)
	const StopTimeout uint8 = 0
	const ContTimeout uint8 = 1
	const MaxMsgSize int = 4 * 1024

	if s.Game.Player1 == (Player{}) {
		s.Game.Player1 = Player{
			Position: rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight * 7 / 8)},
			Size:     rl.Vector2{X: float32(screenWidth / 6), Y: 20},
		}
	} else if s.Game.Player2 == (Player{}) {
		s.Game.Player2 = Player{
			Position: rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight * 1 / 8)},
			Size:     rl.Vector2{X: float32(screenWidth / 6), Y: 20},
		}
	}

	go func() {
		buf := make([]byte, MaxMsgSize)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Println("Read error: ", err)
				timeout <- StopTimeout
				return
			}

			msg := types.WSMessage{}
			err = json.Unmarshal(buf[:n], &msg)
			if err != nil {
				log.Println("Error unmarshal: ", err)
			}

			if msg.SessionID == 1 && msg.Type == "mov" && string(msg.Data) == "LEFT" {
				s.Game.Player1.Position.X -= 8
				messages <- s.Game
			}
			if (s.Game.Player1.Position.X - s.Game.Player1.Size.X/2) <= 0 {
				s.Game.Player1.Position.X = s.Game.Player1.Size.X / 2
				messages <- s.Game
			}

			if msg.SessionID == 1 && msg.Type == "mov" && string(msg.Data) == "RIGHT" {
				s.Game.Player1.Position.X += 8
				messages <- s.Game
			}
			if (s.Game.Player1.Position.X + s.Game.Player1.Size.X/2) >= screenWidth {
				s.Game.Player1.Position.X = screenWidth - s.Game.Player1.Size.X/2
				messages <- s.Game
			}

			if msg.SessionID == 2 && msg.Type == "mov" && string(msg.Data) == "LEFT" {
				s.Game.Player2.Position.X -= 8
				messages <- s.Game
			}
			if (s.Game.Player2.Position.X - s.Game.Player2.Size.X/2) <= 0 {
				s.Game.Player2.Position.X = s.Game.Player2.Size.X / 2
				messages <- s.Game
			}

			if msg.SessionID == 2 && msg.Type == "mov" && string(msg.Data) == "RIGHT" {
				s.Game.Player2.Position.X += 8
				messages <- s.Game
			}
			if (s.Game.Player2.Position.X + s.Game.Player2.Size.X/2) >= screenWidth {
				s.Game.Player2.Position.X = screenWidth - s.Game.Player2.Size.X/2
				messages <- s.Game
			}

			timeout <- ContTimeout
		}
	}()

ExitTimeout:
	for {
		select {
		case res := <-timeout:
			if res == StopTimeout {
				log.Println("Manually stopping timeout manager")
				break ExitTimeout
			}
		case <-time.After(timeoutTime):
			log.Println("User timed out")
			break ExitTimeout
		}
	}
}

func (s websocketServer) clientWriter(conn net.Conn, ch <-chan *Game) {
	for game := range ch {
		js, err := json.Marshal(game)
		if err != nil {
			log.Println("Error marshaling json: ", err)
			return
		}

		_, err = conn.Write(js)
		if err != nil {
			log.Println("Write error: ", err)
			return
		}
	}
}
