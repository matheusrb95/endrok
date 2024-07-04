package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	texture  rl.Texture2D
	position rl.Vector2
	selector rl.Texture2D
}

func NewButton(fileName string, position rl.Vector2) *Button {
	return &Button{
		texture:  rl.LoadTexture(fileName),
		position: position,
		selector: rl.LoadTexture("assets/selector.png"),
	}
}

func (b *Button) Unload() {
	rl.UnloadTexture(b.texture)
	rl.UnloadTexture(b.selector)
}

func (b *Button) Draw() {
	rl.DrawTextureV(b.texture, b.position, rl.White)
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(b.position.X, b.position.Y, float32(b.texture.Width), float32(b.texture.Height))) {
		pos := rl.NewVector2(b.position.X-35, b.position.Y-35)
		rl.DrawTextureV(b.selector, pos, rl.White)
	}
}

func (b *Button) Pressed(mousePos rl.Vector2, mousePressed bool) bool {
	rect := rl.NewRectangle(b.position.X, b.position.Y, float32(b.texture.Width), float32(b.texture.Height))

	if rl.CheckCollisionPointRec(mousePos, rect) && mousePressed {
		return true
	}
	return false
}
