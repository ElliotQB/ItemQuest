package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	PLAYER_WIDTH        = 50
	PLAYER_HEIGHT       = 128
	PLAYER_MAXXSPEED    = 5
	PLAYER_GROUND_ACCEL = 1
	PLAYER_GROUND_DECEL = 1
	PLAYER_AIR_ACCEL    = 0.3
	PLAYER_AIR_DECEL    = 0.3
	PLAYER_JUMP_HEIGHT  = 12
	PLAYER_RISING_GRV   = 0.4
	PLAYER_FALLING_GRV  = 0.5
	PLAYER_COYOTE_TIME  = 15
	PLAYER_JUMP_BUFFER  = 18
)

type Player struct {
	Pos    rl.Vector2
	Vel    rl.Vector2
	Width  float32
	Height float32
	game   *Game

	coyoteTime float32
	jumpBuffer float32
}

func NewPlayer(g *Game) Player {
	return Player{
		Width:  PLAYER_WIDTH,
		Height: PLAYER_HEIGHT,
		game:   g,
	}
}

func (p *Player) DrawPlayer() {
	rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), int32(p.Width), int32(p.Height), rl.Green)
}

func (p *Player) PlayerTick() {
	inp := &p.game.Input

	// capture input and ground state
	move := BoolToInt(inp.left) - BoolToInt(inp.right)
	onground := p.TileMeeting(p.Pos.X, p.Pos.Y+2)

	// clean accel and decel speeds
	var accel float32
	var decel float32
	accel = PLAYER_GROUND_ACCEL
	decel = PLAYER_GROUND_DECEL
	if !onground {
		accel = PLAYER_AIR_ACCEL
		decel = PLAYER_AIR_DECEL
	}

	// accelerate and decelerate player
	if move == 1 {
		p.Vel.X = min(PLAYER_MAXXSPEED, p.Vel.X+accel)
	} else if move == -1 {
		p.Vel.X = max(-PLAYER_MAXXSPEED, p.Vel.X-accel)
	} else {
		p.Vel.X = MoveValue(p.Vel.X, 0, decel)
	}

	// clean grv speed
	var grv float32
	grv = PLAYER_RISING_GRV
	if p.Vel.Y > 0 {
		grv = PLAYER_FALLING_GRV
	}

	// apply grv
	p.Vel.Y += grv

	// set jump buffer when player presses jump
	if inp.jumpI {
		p.jumpBuffer = PLAYER_JUMP_BUFFER
	}

	// launch player up when jump buffer is valid and player is on ground
	if p.jumpBuffer > 0 && onground {
		p.Vel.Y = -PLAYER_JUMP_HEIGHT
	}

	// decrease jump buffer
	p.jumpBuffer = max(0, p.jumpBuffer-1)

	// horizontal collision
	if p.TileMeeting(p.Pos.X+p.Vel.X, p.Pos.Y) {
		for !p.TileMeeting(p.Pos.X+Sign(p.Vel.X), p.Pos.Y) {
			p.Pos.X += Sign(p.Vel.X)
		}
		p.Vel.X = 0
	}

	// vertical collision
	if p.TileMeeting(p.Pos.X, p.Pos.Y+p.Vel.Y) {
		for !p.TileMeeting(p.Pos.X, p.Pos.Y+Sign(p.Vel.Y)) {
			p.Pos.Y += Sign(p.Vel.Y)
		}
		p.Vel.Y = 0
	}

	// push player up in case they're trapped in a block forever
	if p.TileMeeting(p.Pos.X, p.Pos.Y) && p.Vel.X == 0 && p.Vel.Y == 0 {
		for p.TileMeeting(p.Pos.X, p.Pos.Y) {
			p.Pos.Y--
		}
	}

	// apply x and y speeds
	p.Pos.X += p.Vel.X
	p.Pos.Y += p.Vel.Y

}

func (p *Player) TileMeeting(x float32, y float32) bool {
	pLeft := StepDown(x, CELL_SIZE) / CELL_SIZE
	pRight := StepDown(x+PLAYER_WIDTH, CELL_SIZE) / CELL_SIZE

	pTop := StepDown(y, CELL_SIZE) / CELL_SIZE
	pBottom := StepDown(y+PLAYER_HEIGHT, CELL_SIZE) / CELL_SIZE

	for i := pLeft; i <= pRight; i++ {
		for j := pTop; j <= pBottom; j++ {
			if i >= 0 && j >= 0 && i < 200 && j < 200 {
				//rl.DrawRectangle(int32(i*CELL_SIZE), int32(j*CELL_SIZE), CELL_SIZE, CELL_SIZE, rl.Blue)
				if p.game.Tiles[int(i)][int(j)] {
					return true
				}
			}
		}
	}
	return false

}
