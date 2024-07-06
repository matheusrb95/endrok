package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Background struct {
	texture  rl.Texture2D
	position rl.Vector2
}

func NewBackground(fileName string, position rl.Vector2) *Background {
	return &Background{
		texture:  rl.LoadTexture(fileName),
		position: position,
	}
}

func (b *Background) Unload() {
	rl.UnloadTexture(b.texture)
}

func (b *Background) Draw() {
	sourceRec := rl.NewRectangle(float32(b.texture.Width/2)-800/2, 0, float32(b.texture.Width), float32(b.texture.Height))

	rl.DrawTextureRec(b.texture, sourceRec, b.position, rl.White)
}
