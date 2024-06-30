package main

import (
	"github.com/matheusrb95/endrok/internal/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gameTitle    = "Quebra Castelo"

	spriteSize = 192
)

type Textures struct {
	Map    rl.Texture2D
	Player rl.Texture2D
}

type Game struct {
	Camera   rl.Camera2D
	Map      *entity.Map
	Player   *entity.Player
	Textures Textures
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(60)

	game := NewGame()

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	game.Unload()
	rl.CloseWindow()
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.Load()

	g.Map = entity.NewMap(&g.Textures.Map)
	g.Player = entity.NewPlayer(&g.Textures.Player)

	g.Camera = rl.Camera2D{
		Target:   rl.NewVector2(0, 0),
		Offset:   rl.NewVector2(screenWidth/2, screenHeight/2),
		Rotation: 0,
		Zoom:     1,
	}
}

func (g *Game) Load() {
	g.Textures.Map = rl.LoadTexture("assets/map.png")
	g.Textures.Player = rl.LoadTexture("assets/knight_blue.png")
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.Textures.Map)
	rl.UnloadTexture(g.Textures.Player)
}

func (g *Game) Update() {
	g.Player.Velocity = rl.NewVector2(0, 0)

	if rl.IsKeyDown(rl.KeyD) {
		g.Player.GoRight()
	} else if rl.IsKeyDown(rl.KeyA) {
		g.Player.GoLeft()
	}

	if rl.IsKeyDown(rl.KeyW) {
		g.Player.GoUp()
	} else if rl.IsKeyDown(rl.KeyS) {
		g.Player.GoDown()
	}

	if rl.IsKeyDown(rl.KeyJ) {
		g.Player.Attack()
	}

	g.Player.Update()

	g.Camera.Target = rl.Vector2AddValue(g.Player.Position, spriteSize/2)
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(g.Camera)

	g.Map.Draw()
	g.Player.Draw()

	rl.EndMode2D()
	rl.EndDrawing()
}
