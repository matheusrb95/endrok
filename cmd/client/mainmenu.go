package main

import (
	"github.com/matheusrb95/endrok/internal/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	Background  *entity.Background
	Title       *entity.Button
	StartButton *entity.Button
	ExitButton  *entity.Button
	Selector    *entity.Button
}

func NewMainMenu() *MainMenu {
	return &MainMenu{
		Background:  entity.NewBackground("assets/background.png", rl.Vector2Zero()),
		Title:       entity.NewButton("assets/title.png", rl.NewVector2(screenWidth/2-160, 50)),
		StartButton: entity.NewButton("assets/button.png", rl.NewVector2(screenWidth/2-96, 200)),
		ExitButton:  entity.NewButton("assets/button.png", rl.NewVector2(screenWidth/2-96, 300)),
		Selector:    entity.NewButton("assets/selector.png", rl.NewVector2(0, 0)),
	}
}

func (m *MainMenu) Unload() {
	m.Background.Unload()
	m.Title.Unload()
	m.StartButton.Unload()
	m.ExitButton.Unload()
	m.Selector.Unload()
}

func (m *MainMenu) Update() {
	if m.Start() {
		ApplicationState = Run
		return
	}
	if m.Exit() {
		ApplicationState = Quit
		return
	}
}

func (m *MainMenu) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	m.Background.Draw()
	m.Title.Draw()
	rl.DrawText("Quebra Castelo", screenWidth/2-120, 60, 30, rl.RayWhite)
	m.StartButton.Draw()
	rl.DrawText("Start", screenWidth/2-40, 210, 30, rl.RayWhite)
	m.ExitButton.Draw()
	rl.DrawText("Exit", screenWidth/2-30, 310, 30, rl.RayWhite)

	rl.EndMode2D()
	rl.EndDrawing()
}

func (m *MainMenu) Start() bool {
	return m.StartButton.Pressed(rl.GetMousePosition(), rl.IsMouseButtonPressed(rl.MouseButtonLeft))
}

func (m *MainMenu) Exit() bool {
	return m.ExitButton.Pressed(rl.GetMousePosition(), rl.IsMouseButtonPressed(rl.MouseButtonLeft))
}
