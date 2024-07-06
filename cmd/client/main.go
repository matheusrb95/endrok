package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gameTitle    = "Quebra Castelo"
)

const (
	Menu = iota
	Run
	Quit
)

var ApplicationState = Run

func main() {
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(60)

	mainMenu := NewMainMenu()
	game := NewGame()

	for !rl.WindowShouldClose() && ApplicationState != Quit {

		switch ApplicationState {
		case Menu:
			mainMenu.Update()
			mainMenu.Draw()
		case Run:
			game.Update()
			game.Draw()
		}
	}

	mainMenu.Unload()
	game.Unload()
	rl.CloseWindow()
}
