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
	player := &game.Player

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		// game logic
		game.Input.InputTick()
		game.Player.PlayerTick()
		camera.MoveCamera(rl.NewVector2(game.Player.Pos.X+(PLAYER_WIDTH/2), game.Player.Pos.Y+(PLAYER_HEIGHT/2)))

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
