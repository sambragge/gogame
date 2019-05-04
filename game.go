package main

import . "github.com/gen2brain/raylib-go/raylib"


type options struct {
	width int32
	height int32
	title string
	fps int32
}

type game struct {
	options
	*World
}

func NewGame() *game {

	_options := options{
		800,
		500,
		"Spektor",
		60,
	}

	x := &game{
		_options,
		newWorld(),

	}

	InitWindow(x.width, x.height, x.title)
	SetTargetFPS(x.fps)

	return x
}

