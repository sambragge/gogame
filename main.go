package main

import . "github.com/gen2brain/raylib-go/raylib"


func main() {
	game := NewGame()

	for !WindowShouldClose() {
		game.draw()
	}

	CloseWindow()
}