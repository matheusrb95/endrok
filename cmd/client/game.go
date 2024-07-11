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
	Map    rl.Texture2D
	MapTop rl.Texture2D
	Player rl.Texture2D
	Enemy  rl.Texture2D
	Death  rl.Texture2D
}

type Game struct {
	Camera    rl.Camera2D
	Map       *entity.Map
	Player    *entity.Player
	Enemy     *entity.Player
	Death     *entity.Death
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
	g.Enemy = entity.NewPlayer(&g.Textures.Enemy)
	g.Death = entity.NewDeath(&g.Textures.Death, rl.NewVector2(1024+32, 1024+32))

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
	g.Textures.Enemy = rl.LoadTexture("assets/torch_red.png")
	g.Textures.Death = rl.LoadTexture("assets/dead.png")
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.Textures.Map)
	rl.UnloadTexture(g.Textures.MapTop)
	rl.UnloadTexture(g.Textures.Player)
	rl.UnloadTexture(g.Textures.Enemy)
	rl.UnloadTexture(g.Textures.Death)
}

func (g *Game) Update() {
	g.Player.Direction = rl.NewVector2(0, 0)
	g.Enemy.Direction = rl.NewVector2(0, 0)

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
		g.Enemy.GoRight()
	} else if rl.IsKeyDown(rl.KeyJ) {
		g.Enemy.GoLeft()
	}

	if rl.IsKeyDown(rl.KeyI) {
		g.Enemy.GoUp()
	} else if rl.IsKeyDown(rl.KeyK) {
		g.Enemy.GoDown()
	}

	if rl.IsKeyDown(rl.KeyU) {
		g.Enemy.Attack()
	}

	physics.Move(g.Player, g.Obstacles)
	physics.Move(g.Enemy, g.Obstacles)

	if rl.CheckCollisionRecs(g.Player.DamageHitbox(), g.Enemy.AttackHitbox()) {
		g.Player.Hitted()
	}
	if rl.CheckCollisionRecs(g.Enemy.DamageHitbox(), g.Player.AttackHitbox()) {
		g.Enemy.Hitted()
	}

	if g.Enemy.Health == 0 {
		g.Death.Update()
	} else {
		g.Enemy.Update()
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
	if g.Enemy.Health == 0 {
		g.Death.Draw()
	} else {
		g.Enemy.Draw()
	}
	g.Map.DrawTop()

	// g.Enemy.DrawAttackHitbox()
	// g.Enemy.DrawMoveHitbox()
	// g.Enemy.DrawDamageHitbox()
	// g.Map.DrawObstacles()

	rl.EndMode2D()
	rl.EndDrawing()
}
