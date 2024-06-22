package main

import (
	"encoding/json"
	"log"
	"net"

	"github.com/matheusrb95/endrok/internal/game"
	"github.com/matheusrb95/endrok/types"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gameTitle    = "Paz de Bol"
)

var sessionID int

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var game game.Game
	go func() {
		buf := make([]byte, 4*1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Println("Read error: ", err)
			}

			err = json.Unmarshal(buf[:n], &game)
			if err != nil {
				log.Println("Error unmarshaling data: ", err)
				continue
			}

			log.Println(game)
		}
	}()

	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
			v := types.WSMessage{
				SessionID: 1,
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
				SessionID: 1,
				Type:      "mov",
				Data:      []byte("RIGHT"),
			}
			js, err := json.Marshal(v)
			if err != nil {
				log.Println("Error marshalling json: ", err)
			}
			conn.Write(js)
		}

		if !game.Ball.Active {
			if rl.IsKeyDown(rl.KeySpace) {
				v := types.WSMessage{
					SessionID: 1,
					Type:      "ball",
					Data:      []byte("SPACE"),
				}
				js, err := json.Marshal(v)
				if err != nil {
					log.Println("Error marshalling json: ", err)
				}
				conn.Write(js)
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		rl.DrawRectangle(int32(game.Player1.Position.X-game.Player1.Size.X/2), int32(game.Player1.Position.Y-game.Player1.Size.Y/2), int32(game.Player1.Size.X), int32(game.Player1.Size.Y), rl.Black)
		for i := 0; i < game.Player1.Life; i++ {
			rl.DrawRectangle(int32(20+40*i), screenHeight-30, 35, 10, rl.LightGray)
		}

		rl.DrawRectangle(int32(game.Player2.Position.X-game.Player2.Size.X/2), int32(game.Player2.Position.Y-game.Player2.Size.Y/2), int32(game.Player2.Size.X), int32(game.Player2.Size.Y), rl.Black)
		for i := 0; i < game.Player2.Life; i++ {
			rl.DrawRectangle(int32(20+40*i), 30, 35, 10, rl.LightGray)
		}

		rl.DrawCircleV(game.Ball.Position, game.Ball.Radius, rl.Maroon)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
