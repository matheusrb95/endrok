package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gameTitle    = "Quebra Castelo"

	spriteSize = 192
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(60)

	init := NewInitScreen()
	game := NewGame()
	startGame := false

	for !rl.WindowShouldClose() {

		if !startGame {
			if init.Exit() {
				break
			} else if init.Start() {
				init.Unload()
				startGame = true
			} else {
				init.Update()
			}
		} else {
			game.Update()
			game.Draw()
		}
	}

	init.Unload()
	game.Unload()
	rl.CloseWindow()
}
