package entity

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	tileSize = 64
)

type Map struct {
	texture    *rl.Texture2D
	textureTop *rl.Texture2D
	frameRec   rl.Rectangle
}

func NewMap(tx *rl.Texture2D, txTop *rl.Texture2D) *Map {

	return &Map{
		texture:    tx,
		textureTop: txTop,
		frameRec:   rl.NewRectangle(0, 0, float32(tx.Width), float32(tx.Height)),
	}
}

func (m *Map) Draw() {
	rl.DrawTextureRec(*m.texture, m.frameRec, rl.NewVector2(0, 0), rl.White)
}

func (m *Map) DrawTop() {
	rl.DrawTextureRec(*m.textureTop, m.frameRec, rl.NewVector2(0, 0), rl.White)
}

func (m *Map) Obstacles() []rl.Rectangle {
	var result []rl.Rectangle

	result = append(result, m.borderCollision()...)
	result = append(result, m.objectCollision()...)

	return result
}

func (m *Map) DrawObstacles() {
	for _, rec := range m.Obstacles() {
		rl.DrawRectangleLinesEx(rec, 5, rl.Red)
	}
}

func (m *Map) borderCollision() []rl.Rectangle {
	var result []rl.Rectangle

	leftBorder := rl.NewRectangle(0, 0, tileSize*8, float32(m.texture.Height))
	rightBorder := rl.NewRectangle(float32(m.texture.Width)-tileSize*8, 0, tileSize*8, float32(m.texture.Height))
	topBorder := rl.NewRectangle(0, 0, float32(m.texture.Width), tileSize*8)
	bottonBorder := rl.NewRectangle(tileSize*8, float32(m.texture.Height)-tileSize*8, float32(m.texture.Width), tileSize*8)

	result = append(result, leftBorder)
	result = append(result, rightBorder)
	result = append(result, topBorder)
	result = append(result, bottonBorder)

	return result
}

func (m *Map) objectCollision() []rl.Rectangle {
	var result []rl.Rectangle

	mountain := rl.NewRectangle(576, 576, 192, 256)

	mountain2a := rl.NewRectangle(1088, 768, 64, 64)
	mountain2b := rl.NewRectangle(1280, 768, 64, 64)
	mountain2c := rl.NewRectangle(1088, 576, 2, 256)
	mountain2d := rl.NewRectangle(1344, 576, 2, 256)
	mountain2e := rl.NewRectangle(1088, 576, 256, 2)

	castle := rl.NewRectangle(1000, 1020, 240, 110)

	tree := rl.NewRectangle(720, 1100, 35, 50)

	result = append(result, mountain)
	result = append(result, mountain2a)
	result = append(result, mountain2b)
	result = append(result, mountain2c)
	result = append(result, mountain2d)
	result = append(result, mountain2e)
	result = append(result, castle)
	result = append(result, tree)

	return result
}
