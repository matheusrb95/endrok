package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Map struct {
	texture  *rl.Texture2D
	frameRec rl.Rectangle
	borders  []rl.Rectangle
}

func NewMap(tx *rl.Texture2D) *Map {
	var borders []rl.Rectangle

	leftBorder := rl.NewRectangle(0, 0, spriteSize, float32(tx.Height))
	rightBorder := rl.NewRectangle(float32(tx.Width)-spriteSize, 0, spriteSize, float32(tx.Height))
	topBorder := rl.NewRectangle(0, 0, float32(tx.Width), spriteSize)
	bottonBorder := rl.NewRectangle(float32(tx.Width)-spriteSize, float32(tx.Height)-spriteSize, float32(tx.Width), spriteSize)

	borders = append(borders, leftBorder)
	borders = append(borders, rightBorder)
	borders = append(borders, topBorder)
	borders = append(borders, bottonBorder)

	return &Map{
		texture:  tx,
		frameRec: rl.NewRectangle(0, 0, float32(tx.Width), float32(tx.Height)),
		borders:  borders,
	}
}

func (m *Map) Draw() {
	rl.DrawTextureRec(*m.texture, m.frameRec, rl.NewVector2(0, 0), rl.White)
}

func (m *Map) Obstacles() []rl.Rectangle {
	return m.borders
}
