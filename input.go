package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Input struct {
	left  bool
	right bool
	up    bool
	down  bool
	jump  bool
	jumpI bool

	game *Game
}

func NewInput(game *Game) Input {
	return Input{game: game}
}

func (i *Input) InputStep() {
	i.left = rl.IsKeyDown(rl.KeyD)
	i.right = rl.IsKeyDown(rl.KeyA)
	i.up = rl.IsKeyDown(rl.KeyW)
	i.down = rl.IsKeyDown(rl.KeyS)
	i.jump = rl.IsKeyDown(rl.KeySpace)
	i.jumpI = rl.IsKeyPressed(rl.KeySpace)
}
