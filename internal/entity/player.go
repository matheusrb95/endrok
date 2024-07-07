package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	numFrames  = 6
	spriteSize = 192
	frameDelay = 6

	initLocationX = 1024
	initLocationY = 1024
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
	Hit       bool
	Speed     int

	texture  *rl.Texture2D
	frameRec rl.Rectangle

	frameCounter int
	frameIndex   int

	color rl.Color
}

func NewPlayer(tx *rl.Texture2D) *Player {
	return &Player{
		texture:  tx,
		Position: rl.NewVector2(initLocationX, initLocationY),
		frameRec: rl.NewRectangle(0, 0, float32(spriteSize), float32(spriteSize)),
		Speed:    200,
		color:    rl.White,
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

func (p *Player) Hitted() {
	if !p.Hit {
		p.frameIndex = 0
		p.Hit = true
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

		if p.Hit {
			if p.frameIndex%2 == 0 {
				p.color = rl.NewColor(255, 255, 255, 150)
			} else {
				p.color = rl.White
			}

			if p.frameIndex == 5 {
				p.Hit = false
			}
		}

		p.frameRec.X = float32(spriteSize * p.frameIndex)
	}
}

func (p *Player) Draw() {
	rec := rl.NewRectangle(p.Position.X+spriteSize/2-35, p.Position.Y+30, 70, 10)
	rl.DrawRectangleRec(rec, rl.Green)
	rl.DrawRectangleLinesEx(rec, 3, rl.Black)
	rl.DrawTextureRec(*p.texture, p.frameRec, p.Position, p.color)
}

func (p *Player) MoveHitbox() rl.Rectangle {
	return rl.NewRectangle(p.Position.X+75, p.Position.Y+120, 45, 18)
}

func (p *Player) AttackHitbox() rl.Rectangle {
	var result rl.Rectangle

	topSprite := rl.NewVector2(p.Position.X+spriteSize/2, p.Position.Y)
	offsetX := float32(25)
	offsetY := float32(25)
	attackWidth := float32(60)
	attackHeight := float32(110)

	pos := rl.NewVector2(topSprite.X+offsetX, topSprite.Y+offsetY)
	if p.frameRec.Width < 0 {
		pos = rl.NewVector2(topSprite.X-offsetX-attackWidth, topSprite.Y+offsetY)
	}

	if p.Attacking {
		result = rl.NewRectangle(pos.X, pos.Y, attackWidth, attackHeight)
	}

	return result
}

func (p *Player) DamageHitbox() rl.Rectangle {
	return rl.NewRectangle(p.Position.X+75, p.Position.Y+60, 45, 60)
}

func (p *Player) DrawAttackHitbox() {
	if p.Attacking {
		rl.DrawRectangleLinesEx(p.AttackHitbox(), 2, rl.Red)
	}
}

func (p *Player) DrawMoveHitbox() {
	rl.DrawRectangleLinesEx(p.MoveHitbox(), 2, rl.Green)
}

func (p *Player) DrawDamageHitbox() {
	rl.DrawRectangleLinesEx(p.DamageHitbox(), 2, rl.Blue)
}
