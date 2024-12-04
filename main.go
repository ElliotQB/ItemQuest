package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1920, 1080, "ItemQuest")
	defer rl.CloseWindow()

	game := NewGame()
	game.PopulateGame()
	game.LoadLevel("level.txt")

	camera := &game.Camera
	player := &game.Player

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		// game logic
		game.Input.InputStep()
		game.Player.PlayerStep()
		camera.MoveCamera(rl.NewVector2(game.Player.Pos.X+(PLAYER_WIDTH/2), game.Player.Pos.Y+(PLAYER_HEIGHT/2)))
		camera.CameraStep()

		// drawing
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(game.Camera.Camera)

		game.DrawTiles()
		//camera.DrawCameraMarker()
		player.DrawPlayer()

		rl.DrawCircle(int32(game.Camera.screenLeft), int32(game.Camera.screenTop), 20, rl.Blue)
		rl.DrawCircle(int32(game.Camera.screenRight), int32(game.Camera.screenBottom), 20, rl.Red)

		rl.EndMode2D()

		rl.EndDrawing()
	}
}
