package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	conn, err := net.Dial("tcp", "192.168.15.4:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	sessionID, err := getSessionID(conn)
	if err != nil {
		fmt.Println("error getting session id:", err)
		return
	}

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
			v := types.Message{
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
			v := types.Message{
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

		if !game.Ball.Active {
			if rl.IsKeyDown(rl.KeySpace) {
				v := types.Message{
					SessionID: sessionID,
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

func getSessionID(conn net.Conn) (int, error) {
	buf := make([]byte, 4*1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Read error: ", err)
	}

	var msg types.Message
	err = json.Unmarshal(buf[:n], &msg)
	if err != nil {
		log.Println("Error unmarshaling data: ", err)
		return 0, err
	}

	return msg.SessionID, nil
}
