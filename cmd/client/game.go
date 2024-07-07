package main

import (
	"github.com/matheusrb95/endrok/internal/entity"
	"github.com/matheusrb95/endrok/internal/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	spriteSize = 192
)

type Textures struct {
	Map     rl.Texture2D
	MapTop  rl.Texture2D
	Player  rl.Texture2D
	Player2 rl.Texture2D
}

type Game struct {
	Camera    rl.Camera2D
	Map       *entity.Map
	Player    *entity.Player
	Player2   *entity.Player
	Textures  Textures
	Obstacles []rl.Rectangle
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.Load()

	g.Map = entity.NewMap(&g.Textures.Map, &g.Textures.MapTop)
	g.Player = entity.NewPlayer(&g.Textures.Player)
	g.Player2 = entity.NewPlayer(&g.Textures.Player2)

	g.Obstacles = g.Map.Obstacles()

	g.Camera = rl.Camera2D{
		Target:   rl.NewVector2(0, 0),
		Offset:   rl.NewVector2(screenWidth/2, screenHeight/2),
		Rotation: 0,
		Zoom:     1,
	}
}

func (g *Game) Load() {
	g.Textures.Map = rl.LoadTexture("assets/map.png")
	g.Textures.MapTop = rl.LoadTexture("assets/map_top.png")
	g.Textures.Player = rl.LoadTexture("assets/knight_blue.png")
	g.Textures.Player2 = rl.LoadTexture("assets/knight_purple.png")
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.Textures.Map)
	rl.UnloadTexture(g.Textures.MapTop)
	rl.UnloadTexture(g.Textures.Player)
	rl.UnloadTexture(g.Textures.Player2)
}

func (g *Game) Update() {
	g.Player.Direction = rl.NewVector2(0, 0)
	g.Player2.Direction = rl.NewVector2(0, 0)

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

	if rl.IsKeyDown(rl.KeyQ) {
		g.Player.Attack()
	}

	if rl.IsKeyDown(rl.KeyL) {
		g.Player2.GoRight()
	} else if rl.IsKeyDown(rl.KeyJ) {
		g.Player2.GoLeft()
	}

	if rl.IsKeyDown(rl.KeyI) {
		g.Player2.GoUp()
	} else if rl.IsKeyDown(rl.KeyK) {
		g.Player2.GoDown()
	}

	if rl.IsKeyDown(rl.KeyU) {
		g.Player2.Attack()
	}

	physics.Move(g.Player, g.Obstacles)
	physics.Move(g.Player2, g.Obstacles)

	if rl.CheckCollisionRecs(g.Player.DamageHitbox(), g.Player2.AttackHitbox()) {
		g.Player.Hit = true
	}
	if rl.CheckCollisionRecs(g.Player2.DamageHitbox(), g.Player.AttackHitbox()) {
		g.Player2.Hit = true
	}

	g.Player.Update()
	g.Player2.Update()

	g.Camera.Target = rl.Vector2AddValue(g.Player.Position, spriteSize/2)
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(g.Camera)

	g.Map.Draw()
	g.Player.Draw()
	g.Player2.Draw()
	g.Map.DrawTop()

	// g.Player.DrawAttackHitbox()
	// g.Player.DrawMoveHitbox()
	// g.Player.DrawDamageHitbox()
	// g.Map.DrawObstacles()

	rl.EndMode2D()
	rl.EndDrawing()
}
