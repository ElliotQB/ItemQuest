package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1280, 720, "ItemQuest")
	defer rl.CloseWindow()

	game := NewGame()
	game.PopulateGame()
	game.LoadLevel("level.txt")
	game.Camera.Camera.Target = rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2)

	camera := &game.Camera

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		// game logic
		if rl.IsKeyDown(rl.KeyRight) {
			game.Camera.MoveCamera(rl.NewVector2(camera.CamX()+1, camera.CamY()+1))
		}

		// drawing
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(game.Camera.Camera)

		game.DrawTiles()
		camera.DrawCameraMarker()

		rl.DrawCircle(int32(game.Camera.screenLeft), int32(game.Camera.screenTop), 20, rl.Blue)
		rl.DrawCircle(int32(game.Camera.screenRight), int32(game.Camera.screenBottom), 20, rl.Red)

		rl.EndMode2D()

		rl.EndDrawing()
	}
}
