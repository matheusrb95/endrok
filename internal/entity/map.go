package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Map struct {
	texture  *rl.Texture2D
	frameRec rl.Rectangle
}

func NewMap(tx *rl.Texture2D) *Map {
	return &Map{
		texture:  tx,
		frameRec: rl.NewRectangle(0, 0, float32(tx.Width), float32(tx.Height)),
	}
}

func (m *Map) Draw() {
	rl.DrawTextureRec(*m.texture, m.frameRec, rl.NewVector2(0, 0), rl.White)
}
