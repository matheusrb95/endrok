package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gameTitle    = "Quebra Castelo"

	spriteSize = 192

	frameDelay  = 6
	numFrames   = 6
	playerSpeed = 3
)

type Knight struct {
	Position  rl.Vector2
	Velocity  rl.Vector2
	Moving    bool
	Attacking bool
}

type Game struct {
	TxKnightBlue   rl.Texture2D
	TxKnightPurple rl.Texture2D
	TxMap          rl.Texture2D
	SourceMapRec   rl.Rectangle
	SourceRecP1    rl.Rectangle
	SourceRecP2    rl.Rectangle
	KnightBlue     Knight
	KnightPurple   Knight
	FrameCounter   int32
	FrameIndex     int32
	Camera2D       rl.Camera2D
}

func NewGame() (g Game) {
	g.Init()
	return
}

func main() {
	game := NewGame()

	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(60)

	game.Load()

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	game.Unload()
	rl.CloseWindow()
}

func (g *Game) Init() {
	g.KnightBlue = Knight{rl.NewVector2(0, 0), rl.NewVector2(0, 0), false, false}
	g.KnightPurple = Knight{rl.NewVector2(0, 0), rl.NewVector2(0, 0), false, false}
	g.SourceRecP1 = rl.NewRectangle(0, 0, spriteSize, spriteSize)
	g.SourceRecP2 = rl.NewRectangle(0, 0, spriteSize, spriteSize)

	g.Camera2D = rl.Camera2D{
		Target:   rl.NewVector2(0, 0),
		Offset:   rl.NewVector2(screenWidth/2, screenHeight/2),
		Rotation: 0,
		Zoom:     1,
	}
}

func (g *Game) Load() {
	g.TxKnightBlue = rl.LoadTexture("assets/knight_blue.png")
	g.TxKnightPurple = rl.LoadTexture("assets/knight_purple.png")
	g.TxMap = rl.LoadTexture("assets/map.png")
	g.SourceMapRec = rl.NewRectangle(0, 0, float32(g.TxMap.Width), float32(g.TxMap.Height))
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.TxKnightBlue)
	rl.UnloadTexture(g.TxKnightPurple)
	rl.UnloadTexture(g.TxMap)
}

func (g *Game) Update() {
	g.KnightBlue.Velocity = rl.NewVector2(0, 0)

	if !g.KnightBlue.Attacking {
		if rl.IsKeyDown(rl.KeyD) {
			if g.SourceRecP1.Width < 0 {
				g.SourceRecP1.Width *= -1
			}

			g.KnightBlue.Velocity.X = 1
		} else if rl.IsKeyDown(rl.KeyA) {
			if g.SourceRecP1.Width > 0 {
				g.SourceRecP1.Width *= -1
			}

			g.KnightBlue.Velocity.X = -1
		}

		if rl.IsKeyDown(rl.KeyW) {
			g.KnightBlue.Velocity.Y = -1
		} else if rl.IsKeyDown(rl.KeyS) {
			g.KnightBlue.Velocity.Y = 1
		}

		g.KnightBlue.Position = rl.Vector2Add(g.KnightBlue.Position, rl.Vector2Scale(rl.Vector2Normalize(g.KnightBlue.Velocity), float32(playerSpeed)))
	}

	if rl.IsKeyDown(rl.KeyJ) {
		g.KnightBlue.Attacking = true
	}

	g.KnightBlue.Moving = g.KnightBlue.Velocity.X != 0 || g.KnightBlue.Velocity.Y != 0

	g.Camera2D.Target = rl.Vector2AddValue(g.KnightBlue.Position, spriteSize/2)

	g.Animate()
}

func (g *Game) Animate() {
	g.FrameCounter++
	if g.FrameCounter >= frameDelay {
		g.FrameCounter = 0

		if g.KnightBlue.Moving {
			g.FrameIndex++
			g.FrameIndex %= numFrames
			g.SourceRecP1.Y = spriteSize * 1
		} else if g.KnightBlue.Attacking {
			g.FrameIndex++
			g.FrameIndex %= numFrames
			g.SourceRecP1.Y = spriteSize * 2

			if g.FrameIndex == 5 {
				g.KnightBlue.Attacking = false
			}
		} else {
			g.FrameIndex++
			g.FrameIndex %= numFrames
			g.SourceRecP1.Y = spriteSize * 0
		}
		g.SourceRecP1.X = float32(spriteSize * g.FrameIndex)
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(g.Camera2D)

	rl.DrawTextureRec(g.TxMap, g.SourceMapRec, rl.NewVector2(0, 0), rl.White)
	rl.DrawTextureRec(g.TxKnightBlue, g.SourceRecP1, g.KnightBlue.Position, rl.White)
	rl.DrawTextureRec(g.TxKnightPurple, g.SourceRecP2, g.KnightPurple.Position, rl.White)

	rl.EndMode2D()
	rl.EndDrawing()
}
