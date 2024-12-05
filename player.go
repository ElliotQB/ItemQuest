package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PLAYER_WIDTH              = 50
	PLAYER_HEIGHT             = 85
	PLAYER_MAXXSPEED          = 9
	PLAYER_MAX_SLIDING_YSPEED = 3
	PLAYER_WALLJUMP_HLAUNCH   = 11
	PLAYER_GROUND_ACCEL       = 1
	PLAYER_GROUND_DECEL       = 1
	PLAYER_AIR_ACCEL          = 0.3
	PLAYER_AIR_DECEL          = 0.3
	PLAYER_JUMP_HEIGHT        = 12
	PLAYER_RISING_GRV         = 0.4
	PLAYER_FALLING_GRV        = 0.5
	PLAYER_SLIDING_GRV        = 0.2
	PLAYER_COYOTE_TIME        = 9
	PLAYER_JUMP_BUFFER        = 15
	PLAYER_RESPAWN_TIME       = 60
)

const (
	PLAYER_STATE_NEUTRAL    = 0
	PLAYER_STATE_RESPAWNING = 1
)

const (
	ANIM_IDLE        = 0
	ANIM_IDLE_PATH   = "textures/idle-Sheet.png"
	ANIM_IDLE_FRAMES = 4

	ANIM_RUN        = 1
	ANIM_RUN_PATH   = "textures/run-Sheet.png"
	ANIM_RUN_FRAMES = 8

	ANIM_JUMP        = 2
	ANIM_JUMP_PATH   = "textures/jump-Sheet.png"
	ANIM_JUMP_FRAMES = 1

	ANIM_FALL        = 3
	ANIM_FALL_PATH   = "textures/fall-Sheet.png"
	ANIM_FALL_FRAMES = 1

	ANIM_LAND        = 4
	ANIM_LAND_PATH   = "textures/land-Sheet.png"
	ANIM_LAND_FRAMES = 4

	ANIM_WALLSLIDE        = 5
	ANIM_WALLSLIDE_PATH   = "textures/wallslide-Sheet.png"
	ANIM_WALLSLIDE_FRAMES = 1

	ANIM_AIR = 6

	ANIM_DB = 7
)

func LoadAnims() []rl.Texture2D {
	anims := []rl.Texture2D{}
	anims = append(anims, rl.LoadTexture(ANIM_IDLE_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_RUN_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_JUMP_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_FALL_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_LAND_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_WALLSLIDE_PATH))

	return anims
}

type Player struct {
	Pos          rl.Vector2
	Vel          rl.Vector2
	Width        float32
	Height       float32
	Dir          bool
	State        int
	StateCounter float32

	coyoteTime  float32
	jumpBuffer  float32
	numJumps    int
	wallsliding bool
	lastSavePos rl.Vector2

	canDoubleJump bool
	canTripleJump bool
	canWallJump   bool

	game           *Game
	renderer       *SpriteRenderer
	AnimationState uint
	ASTimer        float32
	Anims          []rl.Texture2D
}

func NewPlayer(g *Game) Player {
	anims := LoadAnims()
	return Player{
		Width:    PLAYER_WIDTH,
		Height:   PLAYER_HEIGHT,
		game:     g,
		renderer: NewSpriteRenderer(anims[ANIM_IDLE], ANIM_IDLE_FRAMES, 4, 3, rl.NewVector2(0, -16)),
		Anims:    anims,
		Dir:      true,

		canDoubleJump: false,
		canTripleJump: false,
		canWallJump:   false,
	}
}

func (p *Player) DrawPlayer() {
	//rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), int32(p.Width), int32(p.Height), rl.Green)
	if p.State == PLAYER_STATE_NEUTRAL {
		p.renderer.Render(float32(int(p.Pos.X+(PLAYER_WIDTH/2))), float32(int(p.Pos.Y+(PLAYER_HEIGHT/2))))
	}
}

func (p *Player) PlayerStep() {
	inp := &p.game.Input

	p.StateCounter += 1

	if p.State == PLAYER_STATE_NEUTRAL {
		// capture input and ground state
		move := BoolToInt(inp.left) - BoolToInt(inp.right)
		onground := p.TileMeeting(p.Pos.X, p.Pos.Y+2, WALL)

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
			p.Vel.X = p.Vel.X + accel
			if p.Vel.X > PLAYER_MAXXSPEED {
				p.Vel.X = max(PLAYER_MAXXSPEED, p.Vel.X-0.6)
			}
			if onground {
				p.Dir = true
				if p.Vel.X > PLAYER_MAXXSPEED {
					p.Vel.X = PLAYER_MAXXSPEED
				}
			}
		} else if move == -1 {
			p.Vel.X = p.Vel.X - accel
			if p.Vel.X < -PLAYER_MAXXSPEED {
				p.Vel.X = min(-PLAYER_MAXXSPEED, p.Vel.X+0.6)
			}
			if onground {
				p.Dir = false
				if p.Vel.X < -PLAYER_MAXXSPEED {
					p.Vel.X = -PLAYER_MAXXSPEED
				}
			}
		} else {
			p.Vel.X = MoveValue(p.Vel.X, 0, decel)
		}

		// wall sliding
		if p.canWallJump {
			p.wallsliding = false
			if !onground {
				if p.TileMeeting(p.Pos.X-1, p.Pos.Y, WALL) {
					p.Dir = true
					p.wallsliding = true
				} else if p.TileMeeting(p.Pos.X+1, p.Pos.Y, WALL) {
					p.Dir = false
					p.wallsliding = true
				}
			}
		}

		// record last safe position
		if onground && !p.TileMeeting(p.Pos.X, p.Pos.Y, HAZARD) {
			p.lastSavePos.X = p.Pos.X
			p.lastSavePos.Y = p.Pos.Y
		}

		// collision with hazard
		if p.TileMeeting(p.Pos.X, p.Pos.Y, HAZARD) || p.Pos.Y > 200*CELL_SIZE {
			p.Pos = p.lastSavePos
			p.State = PLAYER_STATE_RESPAWNING
			p.StateCounter = 0
			p.Vel.X = 0
			p.Vel.Y = 0
			p.jumpBuffer = 0
			p.coyoteTime = 0
		}

		// coyote time
		if !onground && p.coyoteTime == -1 {
			p.coyoteTime = PLAYER_COYOTE_TIME
			p.SetNumJumps()
		}
		if onground {
			p.coyoteTime = -1
		}
		if p.coyoteTime != -1 {
			p.coyoteTime = max(0, p.coyoteTime-1)
		}

		// set animation state to air when off ground
		if p.wallsliding {
			p.AnimationState = ANIM_WALLSLIDE
			p.renderer.animationFrame = 0
		} else {
			if !onground {
				if p.AnimationState != ANIM_DB {
					p.AnimationState = ANIM_AIR
				}
			} else {
				if p.AnimationState == ANIM_AIR {
					p.AnimationState = ANIM_LAND
					p.ASTimer = 0
				} else {
					if Abs32(p.Vel.X) < 0.1 || move == 0 {
						if p.AnimationState != ANIM_IDLE && p.AnimationState != ANIM_LAND {
							p.AnimationState = ANIM_IDLE
							p.ASTimer = 0
						}
					} else {
						if p.AnimationState != ANIM_RUN {
							p.AnimationState = ANIM_RUN
							p.ASTimer = 0
						}
					}
				}
			}
		}

		// clean grv speed
		var grv float32
		grv = PLAYER_RISING_GRV
		if p.Vel.Y > 0 {
			grv = PLAYER_FALLING_GRV
		}
		if p.wallsliding && p.Vel.Y > 0 {
			grv = PLAYER_SLIDING_GRV
		}

		// apply grv
		p.Vel.Y += grv

		if p.wallsliding && p.Vel.Y > PLAYER_MAX_SLIDING_YSPEED {
			p.Vel.Y = max(PLAYER_MAX_SLIDING_YSPEED, p.Vel.Y-0.7)
		}

		// wall jump
		if p.wallsliding && inp.jumpI {
			p.Vel.Y = -PLAYER_JUMP_HEIGHT
			p.Vel.X = PLAYER_WALLJUMP_HLAUNCH * BoolSign(p.Dir)
		}

		// set jump buffer when player presses jump
		if inp.jumpI {
			p.jumpBuffer = PLAYER_JUMP_BUFFER
		}

		// launch player up when jump buffer is valid and coyote time is valid or player is on the ground
		if p.jumpBuffer > 0 && (p.coyoteTime > 0 || onground) {
			p.Vel.Y = -PLAYER_JUMP_HEIGHT
			p.coyoteTime = 0
			p.SetNumJumps()
		}

		// aerial jump
		if inp.jumpI && p.coyoteTime == 0 && !onground && p.numJumps >= 1 && !p.wallsliding {
			p.Vel.Y = -PLAYER_JUMP_HEIGHT
			p.numJumps = max(0, p.numJumps-1)

			if Sign(p.Vel.X) != Sign(float32(move)) && move != 0 {
				p.Vel.X = -p.Vel.X
				p.Dir = FloatToBool(float32(move))
			}

			p.AnimationState = ANIM_DB
			p.ASTimer = 0
		}

		// increment animation state
		p.PlayerAnimationStateStep()

		// decrease jump buffer
		p.jumpBuffer = max(0, p.jumpBuffer-1)

		// horizontal collision
		if p.TileMeeting(p.Pos.X+p.Vel.X, p.Pos.Y, WALL) {
			for !p.TileMeeting(p.Pos.X+Sign(p.Vel.X), p.Pos.Y, WALL) {
				p.Pos.X += Sign(p.Vel.X)
			}
			p.Vel.X = 0
		}

		// vertical collision
		if p.TileMeeting(p.Pos.X, p.Pos.Y+p.Vel.Y, WALL) {
			for !p.TileMeeting(p.Pos.X, p.Pos.Y+Sign(p.Vel.Y), WALL) {
				p.Pos.Y += Sign(p.Vel.Y)
			}
			p.Vel.Y = 0
		}

		// push player out in case they're trapped in a block forever
		if p.TileMeeting(p.Pos.X, p.Pos.Y, WALL) && p.Vel.X == 0 && p.Vel.Y == 0 {
			p.PushOut()
		}

		// apply x and y speeds
		p.Pos.X += p.Vel.X
		p.Pos.Y += p.Vel.Y

	} else if p.State == PLAYER_STATE_RESPAWNING {
		if p.StateCounter >= PLAYER_RESPAWN_TIME {
			p.State = PLAYER_STATE_NEUTRAL
		}
	}

}

func (p *Player) PlayerAnimationStateStep() {

	switch p.AnimationState {
	case ANIM_IDLE:
		p.renderer.flip = !p.Dir
		if p.ASTimer == 0 {
			p.renderer.sprite = p.Anims[ANIM_IDLE]
			p.renderer.numberFrames = ANIM_IDLE_FRAMES
			p.renderer.animationFrame = 0
			p.renderer.animationSpeed = 0.07
		}

	case ANIM_RUN:
		p.renderer.flip = !p.Dir
		if p.ASTimer == 0 {
			p.renderer.sprite = p.Anims[ANIM_RUN]
			p.renderer.numberFrames = ANIM_RUN_FRAMES
			p.renderer.animationFrame = 0
			p.renderer.animationSpeed = 0.2
		}

	case ANIM_LAND:
		p.renderer.flip = !p.Dir
		if p.ASTimer == 0 {
			p.renderer.sprite = p.Anims[ANIM_LAND]
			p.renderer.numberFrames = ANIM_LAND_FRAMES
			p.renderer.animationFrame = 0
			p.renderer.animationSpeed = 0.12
		}

		if p.renderer.animationFrame >= ANIM_LAND_FRAMES-1 {
			p.AnimationState = ANIM_IDLE
			p.renderer.sprite = p.Anims[ANIM_IDLE]
			p.renderer.numberFrames = ANIM_IDLE_FRAMES
		}

	case ANIM_WALLSLIDE:
		p.renderer.flip = !p.Dir
		p.renderer.sprite = p.Anims[ANIM_WALLSLIDE]
		p.renderer.numberFrames = ANIM_WALLSLIDE_FRAMES

	case ANIM_AIR:
		p.renderer.flip = !p.Dir
		if p.Vel.Y < 0 {
			p.renderer.sprite = p.Anims[ANIM_JUMP]
			p.renderer.numberFrames = ANIM_JUMP_FRAMES
		} else {
			p.renderer.sprite = p.Anims[ANIM_FALL]
			p.renderer.numberFrames = ANIM_FALL_FRAMES
		}
		p.renderer.animationFrame = 0

	case ANIM_DB:
		p.renderer.flip = !p.Dir

		if p.ASTimer == 0 {
			p.renderer.sprite = p.Anims[ANIM_FALL]
			p.renderer.numberFrames = ANIM_FALL_FRAMES
			p.renderer.animationFrame = 0
		}

		if p.ASTimer >= 10 {
			p.AnimationState = ANIM_AIR
		}
	}

	p.ASTimer++
	p.renderer.AnimationStep()
}

func (p *Player) TileMeeting(x float32, y float32, tileType int) bool {
	pLeft := StepDown(x, CELL_SIZE) / CELL_SIZE
	pRight := StepDown(x+PLAYER_WIDTH, CELL_SIZE) / CELL_SIZE

	pTop := StepDown(y, CELL_SIZE) / CELL_SIZE
	pBottom := StepDown(y+PLAYER_HEIGHT, CELL_SIZE) / CELL_SIZE

	for i := pLeft; i <= pRight; i++ {
		for j := pTop; j <= pBottom; j++ {
			if i >= 0 && j >= 0 && i < 200 && j < 200 {
				//rl.DrawRectangle(int32(i*CELL_SIZE), int32(j*CELL_SIZE), CELL_SIZE, CELL_SIZE, rl.Blue)
				if p.game.Tiles[int(i)][int(j)] == tileType {
					return true
				}
			}
		}
	}
	return false

}

func (p *Player) SetNumJumps() {
	if !p.canDoubleJump {
		p.numJumps = 0
	} else if !p.canTripleJump {
		p.numJumps = 1
	} else {
		p.numJumps = 2
	}
}

func (p *Player) PushOut() {
	origX := p.Pos.X
	origY := p.Pos.Y

	// rotate around every possible direction to push the player out of a wall until you find the shortest path
	for i := 1; i < 500; i++ {
		f_i := float32(i)
		if !p.TileMeeting(origX, origY+f_i, WALL) {
			p.Pos.X = origX
			p.Pos.Y = origY + f_i
			return
		}
		if !p.TileMeeting(origX, origY-f_i, WALL) {
			p.Pos.X = origX
			p.Pos.Y = origY - f_i
			return
		}
		if !p.TileMeeting(origX+f_i, origY, WALL) {
			p.Pos.X = origX + f_i
			p.Pos.Y = origY
			return
		}
		if !p.TileMeeting(origX-f_i, origY, WALL) {
			p.Pos.X = origX - f_i
			p.Pos.Y = origY
			return
		}
	}

	// failsafe; push player up
	for !p.TileMeeting(p.Pos.X, p.Pos.Y, WALL) {
		p.Pos.Y--
	}
}
