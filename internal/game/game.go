package game

import (
	"encoding/json"
	"math"

	"github.com/matheusrb95/endrok/types"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth     = 800
	screenHeight    = 600
	PLAYER_MAX_LIFE = 5
)

type Player struct {
	Position rl.Vector2
	Size     rl.Vector2
	Life     int
}

type Ball struct {
	Position rl.Vector2
	Speed    rl.Vector2
	Radius   float32
	Active   bool
}

type Game struct {
	Player1 Player
	Player2 Player
	Ball    Ball
}

func New() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.Player1.Position = rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight * 7 / 8)}
	g.Player1.Size = rl.Vector2{X: float32(screenWidth / 6), Y: 20}
	g.Player1.Life = PLAYER_MAX_LIFE

	g.Player2.Position = rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight * 1 / 8)}
	g.Player2.Size = rl.Vector2{X: float32(screenWidth / 6), Y: 20}
	g.Player2.Life = PLAYER_MAX_LIFE

	g.Ball.Position = rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight*7/8 - 30)}
	g.Ball.Speed = rl.Vector2{X: 0, Y: 0}
	g.Ball.Radius = 10
	g.Ball.Active = false
}

func (g *Game) UpdateMoves(msg types.Message) {
	if msg.SessionID == 1 && msg.Type == "mov" && string(msg.Data) == "LEFT" {
		g.Player1.Position.X -= 8
	}
	if (g.Player1.Position.X - g.Player1.Size.X/2) <= 0 {
		g.Player1.Position.X = g.Player1.Size.X / 2
	}

	if msg.SessionID == 1 && msg.Type == "mov" && string(msg.Data) == "RIGHT" {
		g.Player1.Position.X += 8
	}
	if (g.Player1.Position.X + g.Player1.Size.X/2) >= screenWidth {
		g.Player1.Position.X = screenWidth - g.Player1.Size.X/2
	}

	if msg.SessionID == 2 && msg.Type == "mov" && string(msg.Data) == "LEFT" {
		g.Player2.Position.X -= 8
	}
	if (g.Player2.Position.X - g.Player2.Size.X/2) <= 0 {
		g.Player2.Position.X = g.Player2.Size.X / 2
	}

	if msg.SessionID == 2 && msg.Type == "mov" && string(msg.Data) == "RIGHT" {
		g.Player2.Position.X += 8
	}
	if (g.Player2.Position.X + g.Player2.Size.X/2) >= screenWidth {
		g.Player2.Position.X = screenWidth - g.Player2.Size.X/2
	}

	if !g.Ball.Active {
		if msg.SessionID == 1 && msg.Type == "ball" && string(msg.Data) == "SPACE" {
			g.Ball.Active = true
			g.Ball.Speed = rl.Vector2{X: 0, Y: -10}
		}
	}
}

func (g *Game) Update() {
	if g.Ball.Active {
		g.Ball.Position.X += g.Ball.Speed.X
		g.Ball.Position.Y += g.Ball.Speed.Y
	} else {
		g.Ball.Position = rl.Vector2{X: g.Player1.Position.X, Y: screenHeight*7/8 - 30}
	}

	if ((g.Ball.Position.X + g.Ball.Radius) >= screenWidth) || ((g.Ball.Position.X - g.Ball.Radius) <= 0) {
		g.Ball.Speed.X *= -1
	}

	if (g.Ball.Position.Y + g.Ball.Radius) >= screenHeight {
		g.Ball.Speed = rl.Vector2{X: 0, Y: 0}
		g.Ball.Active = false

		g.Player1.Life--
	}
	if (g.Ball.Position.Y + g.Ball.Radius) <= 0 {
		g.Ball.Speed = rl.Vector2{X: 0, Y: 0}
		g.Ball.Active = false

		g.Player2.Life--
	}

	if (rl.CheckCollisionCircleRec(g.Ball.Position, g.Ball.Radius,
		rl.Rectangle{g.Player1.Position.X - g.Player1.Size.X/2, g.Player1.Position.Y - g.Player1.Size.Y/2, g.Player1.Size.X, g.Player1.Size.Y})) {
		if g.Ball.Speed.Y > 0 {
			g.Ball.Speed.Y *= -1
			g.Ball.Speed.X = (g.Ball.Position.X - g.Player1.Position.X) / (g.Player1.Size.X / 2) * 5
		}
	}

	if ((g.Ball.Position.Y - g.Ball.Radius) <= (g.Player2.Position.Y + g.Player2.Size.Y/2)) &&
		((g.Ball.Position.Y - g.Ball.Radius) > (g.Player2.Position.Y + g.Player2.Size.Y/2 + g.Ball.Speed.Y)) &&
		((float32(math.Abs(float64(g.Ball.Position.X - g.Player2.Position.X)))) < (g.Player2.Size.X/2 + g.Ball.Radius*2/3)) &&
		(g.Ball.Speed.Y < 0) {

		g.Ball.Speed.Y *= -1
	}
}

func (g *Game) Serialize() ([]byte, error) {
	return json.Marshal(g)
}

func Deserialize(data []byte) (*Game, error) {
	var g Game
	err := json.Unmarshal(data, &g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}
