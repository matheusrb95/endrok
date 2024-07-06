package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	numFrames  = 6
	spriteSize = 192
	frameDelay = 6

	initLocationX = 350
	initLocationY = 250
)

const (
	Idle = iota
	Walking
	AttackSideType1
	AttackSideType2
	AttackDownType1
	AttackDownType2
	AttackUpType1
	AttackUpType2
)

type Player struct {
	Position  rl.Vector2
	Direction rl.Vector2

	Walking   bool
	Attacking bool
	Speed     int

	texture  *rl.Texture2D
	frameRec rl.Rectangle

	frameCounter int
	frameIndex   int
}

func NewPlayer(tx *rl.Texture2D) *Player {
	return &Player{
		texture:  tx,
		Position: rl.NewVector2(initLocationX, initLocationY),
		frameRec: rl.NewRectangle(0, 0, float32(spriteSize), float32(spriteSize)),
		Speed:    200,
	}
}

func (p *Player) GoUp() {
	if !p.Attacking {
		p.Direction.Y = -1
	}
}

func (p *Player) GoDown() {
	if !p.Attacking {
		p.Direction.Y = 1
	}
}

func (p *Player) GoLeft() {
	if !p.Attacking {
		if p.frameRec.Width > 0 {
			p.frameRec.Width *= -1
		}

		p.Direction.X = -1
	}
}

func (p *Player) GoRight() {
	if !p.Attacking {
		if p.frameRec.Width < 0 {
			p.frameRec.Width *= -1
		}

		p.Direction.X = 1
	}
}

func (p *Player) Attack() {
	if !p.Attacking {
		p.frameIndex = 0
		p.Attacking = true
	}
}

func (p *Player) Update() {
	p.Walking = p.Direction != rl.NewVector2(0, 0)

	p.frameCounter++
	if p.frameCounter >= frameDelay {
		p.frameCounter = 0

		if p.Walking {
			p.frameIndex++
			p.frameIndex %= numFrames
			p.frameRec.Y = spriteSize * Walking
		} else if p.Attacking {
			p.frameIndex++
			p.frameIndex %= numFrames
			p.frameRec.Y = spriteSize * AttackSideType1

			if p.frameIndex == 5 {
				p.Attacking = false
			}
		} else {
			p.frameIndex++
			p.frameIndex %= numFrames
			p.frameRec.Y = spriteSize * Idle
		}

		p.frameRec.X = float32(spriteSize * p.frameIndex)
	}
}

func (p *Player) Draw() {
	rec := rl.NewRectangle(p.Position.X+spriteSize/2-35, p.Position.Y+30, 70, 10)
	rl.DrawRectangleRec(rec, rl.Green)
	rl.DrawRectangleLinesEx(rec, 3, rl.Black)
	rl.DrawTextureRec(*p.texture, p.frameRec, p.Position, rl.White)
}

func (p *Player) MoveHitbox() rl.Rectangle {
	return rl.NewRectangle(p.Position.X+75, p.Position.Y+120, 45, 18)
	//eturn rl.NewRectangle(p.Position.X, p.Position.Y, spriteSize, spriteSize)
}

func (p *Player) DrawMoveHitbox(coliding bool) {
	if coliding {
		rl.DrawRectangleLinesEx(p.MoveHitbox(), 3, rl.Red)
	}
}

func (p *Player) AttackHitbox() rl.Vector2 {
	return rl.NewVector2(p.Position.X+spriteSize/2, p.Position.Y+spriteSize/2)
}

func (p *Player) DrawAttackHitbox(coliding bool) {
	if coliding && p.Attacking {
		rl.DrawCircleLines(int32(p.AttackHitbox().X), int32(p.AttackHitbox().Y), 90, rl.Green)
	}
}
