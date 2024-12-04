package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	screenLeft   float32
	screenRight  float32
	screenTop    float32
	screenBottom float32
	Camera       rl.Camera2D
	tweenPos     rl.Vector2
	game         *Game
}

func NewCamera(game *Game) Camera {
	return Camera{
		screenLeft:   0,
		screenRight:  float32(rl.GetScreenWidth()),
		screenTop:    0,
		screenBottom: float32(rl.GetScreenHeight()),
		Camera:       rl.NewCamera2D(rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2), rl.NewVector2(0, 0), 0, 1),
		tweenPos:     rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
		game:         game,
	}
}

func (c *Camera) MoveCamera(pos rl.Vector2) {
	c.tweenPos = pos
}

func (c *Camera) CameraTick() {
	c.Camera.Target.X += (c.tweenPos.X - c.Camera.Target.X) * 0.05
	c.Camera.Target.Y += (c.tweenPos.Y - c.Camera.Target.Y) * 0.05

	c.screenLeft = float32(c.Camera.Target.X) - float32(rl.GetScreenWidth())/2
	c.screenRight = float32(c.Camera.Target.X) + float32(rl.GetScreenWidth())/2
	c.screenTop = float32(c.Camera.Target.Y) - float32(rl.GetScreenHeight())/2
	c.screenBottom = float32(c.Camera.Target.Y) + float32(rl.GetScreenHeight())/2
}

func (c Camera) DrawCameraMarker() {
	rl.DrawCircle(int32(c.Camera.Target.X), int32(c.Camera.Target.Y), 5, rl.Blue)
}

func (c *Camera) CamX() float32 {
	return c.tweenPos.X
}

func (c *Camera) CamY() float32 {
	return c.tweenPos.Y
}
