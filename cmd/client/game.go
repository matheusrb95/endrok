package main

import (
	"github.com/matheusrb95/endrok/internal/entity"
	"github.com/matheusrb95/endrok/internal/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	spriteSize = 192
)

var (
	Obstacle  = rl.NewRectangle(600, 600, 250, 250)
	Obstacles = []rl.Rectangle{Obstacle}
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

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.Load()

	g.Map = entity.NewMap(&g.Textures.Map)
	g.Player = entity.NewPlayer(&g.Textures.Player)

	Obstacles = append(Obstacles, g.Map.Obstacles()...)

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
	g.Player.Direction = rl.NewVector2(0, 0)

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

	physics.Move(g.Player, Obstacles)

	g.Player.Update()

	g.Camera.Target = rl.Vector2AddValue(g.Player.Position, spriteSize/2)
}

func (g *Game) Draw() {
	coliding := rl.CheckCollisionRecs(g.Player.MoveHitbox(), Obstacle)
	attackColiding := rl.CheckCollisionCircleRec(g.Player.AttackHitbox(), spriteSize/2, Obstacle)

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(g.Camera)

	g.Map.Draw()
	g.Player.Draw()
	g.Player.DrawMoveHitbox(coliding)
	g.Player.DrawAttackHitbox(attackColiding)
	rl.DrawRectangleLinesEx(Obstacle, 5, rl.Black)

	rl.EndMode2D()
	rl.EndDrawing()
}
