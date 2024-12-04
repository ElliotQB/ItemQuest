package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	PLAYER_WIDTH        = 50
	PLAYER_HEIGHT       = 85
	PLAYER_MAXXSPEED    = 6
	PLAYER_GROUND_ACCEL = 1
	PLAYER_GROUND_DECEL = 1
	PLAYER_AIR_ACCEL    = 0.3
	PLAYER_AIR_DECEL    = 0.3
	PLAYER_JUMP_HEIGHT  = 12
	PLAYER_RISING_GRV   = 0.4
	PLAYER_FALLING_GRV  = 0.5
	PLAYER_COYOTE_TIME  = 9
	PLAYER_JUMP_BUFFER  = 15
)

const (
	ANIM_IDLE        = 0
	ANIM_IDLE_PATH   = "textures/idle-Sheet.png"
	ANIM_IDLE_FRAMES = 4

	ANIM_RUN        = 1
	ANIM_RUN_PATH   = "textures/run-Sheet.png"
	ANIM_RUN_FRAMES = 9

	ANIM_JUMP        = 2
	ANIM_JUMP_PATH   = "textures/jump-Sheet.png"
	ANIM_JUMP_FRAMES = 1

	ANIM_FALL        = 3
	ANIM_FALL_PATH   = "textures/fall-Sheet.png"
	ANIM_FALL_FRAMES = 1

	ANIM_LAND        = 4
	ANIM_LAND_PATH   = "textures/land-Sheet.png"
	ANIM_LAND_FRAMES = 4
)

func LoadAnims() []rl.Texture2D {
	anims := []rl.Texture2D{}
	anims = append(anims, rl.LoadTexture(ANIM_IDLE_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_RUN_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_JUMP_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_FALL_PATH))
	anims = append(anims, rl.LoadTexture(ANIM_LAND_PATH))

	return anims
}

type Player struct {
	Pos    rl.Vector2
	Vel    rl.Vector2
	Width  float32
	Height float32

	coyoteTime float32
	jumpBuffer float32

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
	}
}

func (p *Player) DrawPlayer() {
	//rl.DrawRectangle(int32(p.Pos.X), int32(p.Pos.Y), int32(p.Width), int32(p.Height), rl.Green)
	p.renderer.Render(float32(int(p.Pos.X+(PLAYER_WIDTH/2))), float32(int(p.Pos.Y+(PLAYER_HEIGHT/2))))
}

func (p *Player) PlayerStep() {
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

	if !onground && p.coyoteTime == -1 {
		p.coyoteTime = PLAYER_COYOTE_TIME
	}
	if onground {
		p.coyoteTime = -1
	}
	if p.coyoteTime != -1 {
		p.coyoteTime = max(0, p.coyoteTime-1)
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

	// launch player up when jump buffer is valid and coyote time is valid or player is on the ground
	if p.jumpBuffer > 0 && (p.coyoteTime > 0 || onground) {
		p.Vel.Y = -PLAYER_JUMP_HEIGHT
		p.coyoteTime = 0
	}

	// increment animation state
	p.PlayerAnimationStateStep()

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

	// push player out in case they're trapped in a block forever
	if p.TileMeeting(p.Pos.X, p.Pos.Y) && p.Vel.X == 0 && p.Vel.Y == 0 {
		p.PushOut()
	}

	// apply x and y speeds
	p.Pos.X += p.Vel.X
	p.Pos.Y += p.Vel.Y

}

func (p *Player) PlayerAnimationStateStep() {

	switch p.AnimationState {
	case ANIM_IDLE:
		if p.ASTimer == 0 {
			p.renderer.sprite = p.Anims[ANIM_IDLE]
			p.renderer.animationFrame = 0
			p.renderer.animationSpeed = 4
		}

	}

	p.ASTimer++
	p.renderer.AnimationStep()
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

func (p *Player) PushOut() {
	origX := p.Pos.X
	origY := p.Pos.Y

	// rotate around every possible direction to push the player out of a wall until you find the shortest path
	for i := 1; i < 500; i++ {
		f_i := float32(i)
		if !p.TileMeeting(origX, origY+f_i) {
			p.Pos.X = origX
			p.Pos.Y = origY + f_i
			return
		}
		if !p.TileMeeting(origX, origY-f_i) {
			p.Pos.X = origX
			p.Pos.Y = origY - f_i
			return
		}
		if !p.TileMeeting(origX+f_i, origY) {
			p.Pos.X = origX + f_i
			p.Pos.Y = origY
			return
		}
		if !p.TileMeeting(origX-f_i, origY) {
			p.Pos.X = origX - f_i
			p.Pos.Y = origY
			return
		}
	}

	// failsafe; push player up
	for !p.TileMeeting(p.Pos.X, p.Pos.Y) {
		p.Pos.Y--
	}
}
