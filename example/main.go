package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PLAYER_MAX_LIFE = 5
)

type Player struct {
	position rl.Vector2
	size     rl.Vector2
	life     int
}

type Ball struct {
	position rl.Vector2
	speed    rl.Vector2
	radius   float32
	active   bool
}

const (
	screenWidth  = 800
	screenHeight = 600
	gameTitle    = "Paz de Bol"
)

type Game struct {
	gameOver bool
	player1  Player
	player2  Player
	ball     Ball
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, gameTitle)

	game := NewGame()
	game.gameOver = true

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	rl.CloseWindow()
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.player1.position = rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight * 7 / 8)}
	g.player1.size = rl.Vector2{X: float32(screenWidth / 6), Y: 20}
	g.player1.life = PLAYER_MAX_LIFE

	g.player2.position = rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight * 1 / 8)}
	g.player2.size = rl.Vector2{X: float32(screenWidth / 6), Y: 20}
	g.player2.life = PLAYER_MAX_LIFE

	g.ball.position = rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight*7/8 - 30)}
	g.ball.speed = rl.Vector2{X: 0, Y: 0}
	g.ball.radius = 10
	g.ball.active = false
}

func (g *Game) Update() {
	if !g.gameOver {
		if rl.IsKeyDown(rl.KeyLeft) {
			g.player1.position.X -= 8
		}
		if (g.player1.position.X - g.player1.size.X/2) <= 0 {
			g.player1.position.X = g.player1.size.X / 2
		}

		if rl.IsKeyDown(rl.KeyA) {
			g.player2.position.X -= 8
		}
		if (g.player2.position.X - g.player2.size.X/2) <= 0 {
			g.player2.position.X = g.player2.size.X / 2
		}

		if rl.IsKeyDown(rl.KeyRight) {
			g.player1.position.X += 8
		}
		if (g.player1.position.X + g.player1.size.X/2) >= screenWidth {
			g.player1.position.X = screenWidth - g.player1.size.X/2
		}

		if rl.IsKeyDown(rl.KeyD) {
			g.player2.position.X += 8
		}
		if (g.player2.position.X + g.player2.size.X/2) >= screenWidth {
			g.player2.position.X = screenWidth - g.player2.size.X/2
		}

		if !g.ball.active {
			if rl.IsKeyPressed(rl.KeySpace) {
				g.ball.active = true
				g.ball.speed = rl.Vector2{X: 0, Y: -5}
			}
		}

		if g.ball.active {
			g.ball.position.X += g.ball.speed.X
			g.ball.position.Y += g.ball.speed.Y
		} else {
			g.ball.position = rl.Vector2{X: g.player1.position.X, Y: screenHeight*7/8 - 30}
		}

		if ((g.ball.position.X + g.ball.radius) >= screenWidth) || ((g.ball.position.X - g.ball.radius) <= 0) {
			g.ball.speed.X *= -1
		}

		if (g.ball.position.Y + g.ball.radius) >= screenHeight {
			g.ball.speed = rl.Vector2{X: 0, Y: 0}
			g.ball.active = false

			g.player1.life--
		}
		if (g.ball.position.Y + g.ball.radius) <= 0 {
			g.ball.speed = rl.Vector2{X: 0, Y: 0}
			g.ball.active = false

			g.player2.life--
		}

		if (rl.CheckCollisionCircleRec(g.ball.position, g.ball.radius,
			rl.Rectangle{g.player1.position.X - g.player1.size.X/2, g.player1.position.Y - g.player1.size.Y/2, g.player1.size.X, g.player1.size.Y})) {
			if g.ball.speed.Y > 0 {
				g.ball.speed.Y *= -1
				g.ball.speed.X = (g.ball.position.X - g.player1.position.X) / (g.player1.size.X / 2) * 5
			}
		}

		if ((g.ball.position.Y - g.ball.radius) <= (g.player2.position.Y + g.player2.size.Y/2)) &&
			((g.ball.position.Y - g.ball.radius) > (g.player2.position.Y + g.player2.size.Y/2 + g.ball.speed.Y)) &&
			((float32(math.Abs(float64(g.ball.position.X - g.player2.position.X)))) < (g.player2.size.X/2 + g.ball.radius*2/3)) &&
			(g.ball.speed.Y < 0) {

			g.ball.speed.Y *= -1
		}

		if g.player1.life <= 0 || g.player2.life <= 0 {
			g.gameOver = true
		} else {
			g.gameOver = false
		}

	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.Init()
			g.gameOver = false
		}
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	if !g.gameOver {
		rl.DrawRectangle(int32(g.player1.position.X-g.player1.size.X/2), int32(g.player1.position.Y-g.player1.size.Y/2), int32(g.player1.size.X), int32(g.player1.size.Y), rl.Black)
		for i := 0; i < g.player1.life; i++ {
			rl.DrawRectangle(int32(20+40*i), screenHeight-30, 35, 10, rl.LightGray)
		}

		rl.DrawRectangle(int32(g.player2.position.X-g.player2.size.X/2), int32(g.player2.position.Y-g.player2.size.Y/2), int32(g.player2.size.X), int32(g.player2.size.Y), rl.Black)
		for i := 0; i < g.player2.life; i++ {
			rl.DrawRectangle(int32(20+40*i), 30, 35, 10, rl.LightGray)
		}

		rl.DrawCircleV(g.ball.position, g.ball.radius, rl.Maroon)
	} else {
		str := "PRESS [ENTER] TO PLAY AGAIN"
		x := int(rl.GetScreenWidth()/2) - int(rl.MeasureText(str, 20)/2)
		y := rl.GetScreenHeight()/2 - 50
		rl.DrawText(str, int32(x), int32(y), 20, rl.Gray)
	}

	rl.EndDrawing()
}
