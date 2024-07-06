package physics

import (
	"github.com/matheusrb95/endrok/internal/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	spriteSize = 192
)

func CollideWithRects(rect rl.Rectangle, rects []rl.Rectangle) []rl.Rectangle {
	var hitSlice []rl.Rectangle
	for _, rec := range rects {
		if rl.CheckCollisionRecs(rect, rec) {
			hitSlice = append(hitSlice, rec)
		}
	}
	return hitSlice
}

func Move(p *entity.Player, rects []rl.Rectangle) {
	norm := rl.Vector2Scale(rl.Vector2Normalize(p.Direction), rl.GetFrameTime()*float32(p.Speed))

	p.Position.X += norm.X
	hitList := CollideWithRects(p.MoveHitbox(), rects)
	for _, tile := range hitList {
		if norm.X > 0 {
			p.Position.X = tile.X - (p.MoveHitbox().Width + p.MoveHitbox().X - p.Position.X)
		} else if norm.X < 0 {
			p.Position.X = tile.X + tile.Width - (p.MoveHitbox().X - p.Position.X)
		}
	}

	p.Position.Y += norm.Y
	hitList = CollideWithRects(p.MoveHitbox(), rects)
	for _, tile := range hitList {
		if norm.Y > 0 {
			p.Position.Y = tile.Y - (p.MoveHitbox().Height + p.MoveHitbox().Y - p.Position.Y)
		} else if norm.Y < 0 {
			p.Position.Y = tile.Y + tile.Height - (p.MoveHitbox().Y - p.Position.Y)
		}
	}
}
