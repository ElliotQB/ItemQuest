package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SpriteRenderer struct {
	sprite         rl.Texture2D
	animationFrame float32
	animationSpeed float32
	numberFrames   int
	scale          float32
	flip           bool
	offset         rl.Vector2
}

func NewSpriteRenderer(sprite rl.Texture2D, numberFrames int, animationSpeed float32, scale float32, offset rl.Vector2) *SpriteRenderer {
	return &SpriteRenderer{sprite: sprite, animationFrame: 0, animationSpeed: animationSpeed, numberFrames: numberFrames, scale: scale, flip: false, offset: offset}
}

func (r SpriteRenderer) Render(x float32, y float32) {
	spriteWidth := r.sprite.Width / int32(r.numberFrames)

	animFrame := (int32(Floor32(r.animationFrame))) % int32(r.numberFrames)

	source := rl.NewRectangle(float32(spriteWidth*animFrame), 0, float32(spriteWidth)*BoolSign(!r.flip), float32(r.sprite.Height))
	dest := rl.NewRectangle(x-(float32(spriteWidth*int32(r.scale))/2)+r.offset.X, y-(float32(r.sprite.Height*int32(r.scale))/2)+r.offset.Y, float32(spriteWidth)*r.scale, float32(r.sprite.Height)*r.scale)

	rl.DrawTexturePro(r.sprite, source, dest, rl.NewVector2(0, 0), 0, rl.White)
}

func (r *SpriteRenderer) AnimationStep() {
	r.animationFrame += r.animationSpeed
}
