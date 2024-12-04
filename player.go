package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	PLAYER_WIDTH        = 32
	PLAYER_HEIGHT       = 64
	PLAYER_MAXXSPEED    = 5
	PLAYER_GROUND_ACCEL = 1
	PLAYER_GROUND_DECEL = 1
	PLAYER_AIR_ACCEL    = 0.3
	PLAYER_AIR_DECEL    = 0.3
	JUMP_HEIGHT         = 10
	RISING_GRV          = 0.4
	FALLING_GRV         = 0.6
)

type Player struct {
	Pos    rl.Vector2
	Vel    rl.Vector2
	Width  float32
	Height float32
	game   *Game
}

func NewPlayer(g *Game) Player {
	return Player{Width: PLAYER_WIDTH, Height: PLAYER_HEIGHT, game: g}
}

func (p *Player) DrawPlayer() {
	rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), int32(p.Width), int32(p.Height), rl.Green)
}

func (p *Player) PlayerTick() {

}

func (p *Player) TileMeeting(x float32, y float32) bool {
	pLeft := StepDown(x, CELL_SIZE) / CELL_SIZE
	pRight := StepDown(x, CELL_SIZE) / CELL_SIZE

	pTop := StepDown(y-PLAYER_HEIGHT, CELL_SIZE) / CELL_SIZE
	pBottom := StepDown(y+PLAYER_HEIGHT, CELL_SIZE) / CELL_SIZE

	for i := pLeft; i <= pRight; i++ {
		for j := pTop; j <= pBottom; j++ {
			if p.game.Tiles[int(i)][int(j)] {
				return true
			}
		}
	}
	return false

}
