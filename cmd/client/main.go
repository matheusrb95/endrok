package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/matheusrb95/endrok/types"

	rl "github.com/gen2brain/raylib-go/raylib"
	"nhooyr.io/websocket"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gameTitle    = "Paz de Bol"
)

var sessionID int

type Player struct {
	Position rl.Vector2
	Size     rl.Vector2
}

type Game struct {
	Who     string
	Player1 Player
	Player2 Player
}

func main() {
	url := "ws://localhost:8000"
	ctx := context.Background()
	c, resp, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection response: ", resp)

	conn := websocket.NetConn(ctx, c, websocket.MessageBinary)

	msg := make([]byte, 4*1024)
	var wsMsg types.WSMessage
	n, err := conn.Read(msg)
	if err != nil {
		log.Println("Read error: ", err)
	}

	err = json.Unmarshal(msg[:n], &wsMsg)
	if err != nil {
		log.Println("Error unmarshaling data: ", err)
	}

	log.Println(wsMsg)
	sessionID = wsMsg.SessionID

	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(60)

	var game Game
	game.Player1 = Player{
		Position: rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight * 7 / 8)},
		Size:     rl.Vector2{X: float32(screenWidth / 6), Y: 20},
	}

	game.Player2 = Player{
		Position: rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight * 1 / 8)},
		Size:     rl.Vector2{X: float32(screenWidth / 6), Y: 20},
	}

	go func() {
		msg := make([]byte, 4*1024)
		for {
			n, err := conn.Read(msg)
			if err != nil {
				log.Println("Read error: ", err)
			}

			err = json.Unmarshal(msg[:n], &game)
			if err != nil {
				log.Println("Error unmarshaling data: ", err)
			}
			log.Println(game)
		}
	}()

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
			v := types.WSMessage{
				SessionID: sessionID,
				Type:      "mov",
				Data:      []byte("LEFT"),
			}
			js, err := json.Marshal(v)
			if err != nil {
				log.Println("Error marshalling json: ", err)
			}
			conn.Write(js)
		}

		if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
			v := types.WSMessage{
				SessionID: sessionID,
				Type:      "mov",
				Data:      []byte("RIGHT"),
			}
			js, err := json.Marshal(v)
			if err != nil {
				log.Println("Error marshalling json: ", err)
			}
			conn.Write(js)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		rl.DrawRectangle(int32(game.Player1.Position.X-game.Player1.Size.X/2), int32(game.Player1.Position.Y-game.Player1.Size.Y/2), int32(game.Player1.Size.X), int32(game.Player1.Size.Y), rl.Black)
		rl.DrawRectangle(int32(game.Player2.Position.X-game.Player2.Size.X/2), int32(game.Player2.Position.Y-game.Player2.Size.Y/2), int32(game.Player2.Size.X), int32(game.Player2.Size.Y), rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
