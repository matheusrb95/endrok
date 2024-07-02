package main

import (
	"github.com/matheusrb95/endrok/internal/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type InitScreen struct {
	Background  rl.Texture2D
	Title       *entity.Button
	StartButton *entity.Button
	ExitButton  *entity.Button
}

func NewInitScreen() *InitScreen {
	return &InitScreen{
		Background:  rl.LoadTexture("assets/background.png"),
		Title:       entity.NewButton("assets/title.png", rl.NewVector2(screenWidth/2-160, 50)),
		StartButton: entity.NewButton("assets/button.png", rl.NewVector2(screenWidth/2-96, 200)),
		ExitButton:  entity.NewButton("assets/button.png", rl.NewVector2(screenWidth/2-96, 300)),
	}
}

func (s *InitScreen) Unload() {
	s.Title.Unload()
	s.StartButton.Unload()
	s.ExitButton.Unload()
	rl.UnloadTexture(s.Background)
}

func (s *InitScreen) Update() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	sourceRec := rl.NewRectangle(float32(s.Background.Width/2)-screenWidth/2, 0, float32(s.Background.Width), float32(s.Background.Height))
	position := rl.NewVector2(0, 0)
	rl.DrawTextureRec(s.Background, sourceRec, position, rl.White)
	s.Title.Draw()
	rl.DrawText("Quebra Castelo", screenWidth/2-120, 60, 30, rl.RayWhite)
	s.StartButton.Draw()
	rl.DrawText("Start", screenWidth/2-40, 210, 30, rl.RayWhite)
	s.ExitButton.Draw()
	rl.DrawText("Exit", screenWidth/2-30, 310, 30, rl.RayWhite)

	rl.EndMode2D()
	rl.EndDrawing()
}

func (s *InitScreen) Start() bool {
	return s.StartButton.Pressed(rl.GetMousePosition(), rl.IsMouseButtonPressed(rl.MouseButtonLeft))
}

func (s *InitScreen) Exit() bool {
	return s.ExitButton.Pressed(rl.GetMousePosition(), rl.IsMouseButtonPressed(rl.MouseButtonLeft))
}
