package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	DOUBLEJUMPGEM = 0
	TRIPLEJUMPGEM = 1
	WALLJUMPGEM   = 2
	CAT           = 3
)

type Collectable struct {
	Pos       rl.Vector2
	Type      int
	Rad       float32
	Collected bool
}

func NewCollectable(x float32, y float32, collectable int) Collectable {
	return Collectable{Pos: rl.NewVector2(x, y), Type: collectable, Rad: 40, Collected: false}
}

func (c *Collectable) DrawCollectable() {
	color := rl.White
	switch c.Type {
	case DOUBLEJUMPGEM:
		color = rl.Blue
	case TRIPLEJUMPGEM:
		color = rl.SkyBlue
	case WALLJUMPGEM:
		color = rl.Green
	}
	rl.DrawCircle(int32(c.Pos.X), int32(c.Pos.Y), c.Rad, color)
}

func (c *Collectable) Collect() {
	c.Collected = true
}
