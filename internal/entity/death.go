package entity

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Start = iota
	End
)

const (
	deathSpriteSize = 128
	deathNumFrames  = 14
)

type Death struct {
	position     rl.Vector2
	texture      *rl.Texture2D
	frameRec     rl.Rectangle
	frameCounter int
	frameIndex   int
	dead         bool
}

func NewDeath(tx *rl.Texture2D, pos rl.Vector2) *Death {
	return &Death{
		texture:  tx,
		position: pos,
		frameRec: rl.NewRectangle(0, 0, float32(deathSpriteSize), float32(deathSpriteSize)),
		dead:     true,
	}
}

func (d *Death) Update() {
	if d.dead {
		d.frameCounter++
		if d.frameCounter >= frameDelay {
			d.frameCounter = 0

			d.frameIndex++
			if d.frameIndex == deathNumFrames {
				d.frameIndex = 0
				d.dead = false
			}

			d.frameRec.X = float32(d.frameIndex % 7 * deathSpriteSize)
			d.frameRec.Y = float32(d.frameIndex / 7 * deathSpriteSize)
		}
	}
}

func (d *Death) Draw() {
	if d.dead {
		rl.DrawTextureRec(*d.texture, d.frameRec, d.position, rl.White)
	}
}
