package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	screenLeft   float32
	screenRight  float32
	screenTop    float32
	screenBottom float32
	Camera       rl.Camera2D
	game         *Game
}

func NewCamera(game *Game) Camera {
	return Camera{
		screenLeft:   0,
		screenRight:  float32(rl.GetScreenWidth()),
		screenTop:    0,
		screenBottom: float32(rl.GetScreenHeight()),
		Camera:       rl.NewCamera2D(rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2), rl.NewVector2(0, 0), 0, 1),
		game:         game,
	}
}

func (c *Camera) MoveCamera(pos rl.Vector2) {
	c.Camera.Target = pos
	c.screenLeft = float32(pos.X) - float32(rl.GetScreenWidth())/2
	c.screenRight = float32(pos.X) + float32(rl.GetScreenWidth())/2
	c.screenTop = float32(pos.Y) - float32(rl.GetScreenHeight())/2
	c.screenBottom = float32(pos.Y) + float32(rl.GetScreenHeight())/2
}

func (c Camera) DrawCameraMarker() {
	rl.DrawCircle(int32(c.Camera.Target.X), int32(c.Camera.Target.Y), 5, rl.Blue)
}

func (c *Camera) CamX() float32 {
	return c.Camera.Target.X
}

func (c *Camera) CamY() float32 {
	return c.Camera.Target.Y
}
