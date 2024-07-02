package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	texture  rl.Texture2D
	position rl.Vector2
}

func NewButton(fileName string, position rl.Vector2) *Button {
	return &Button{
		texture:  rl.LoadTexture(fileName),
		position: position,
	}
}

func (b *Button) Unload() {
	rl.UnloadTexture(b.texture)
}

func (b *Button) Draw() {
	rl.DrawTextureV(b.texture, b.position, rl.White)
}

func (b *Button) Pressed(mousePos rl.Vector2, mousePressed bool) bool {
	rect := rl.NewRectangle(b.position.X, b.position.Y, float32(b.texture.Width), float32(b.texture.Height))

	if rl.CheckCollisionPointRec(mousePos, rect) && mousePressed {
		return true
	}
	return false
}
