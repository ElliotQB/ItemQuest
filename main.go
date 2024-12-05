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

		playerRec := rl.Rectangle{game.Player.Pos.X, game.Player.Pos.Y, PLAYER_WIDTH, PLAYER_HEIGHT}
		for i := 0; i < len(game.Collectables); i++ {
			if rl.CheckCollisionCircleRec(game.Collectables[i].Pos, 40, playerRec) {
				game.Collectables[i].Collect()

				switch game.Collectables[i].Type {
				case DOUBLEJUMPGEM:
					game.Player.canDoubleJump = true
				case TRIPLEJUMPGEM:
					game.Player.canTripleJump = true
				case WALLJUMPGEM:
					game.Player.canWallJump = true
				case CAT:
					game.Player.canWallJump = true
				}
			}
		}

		// drawing
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(game.Camera.Camera)

		game.DrawTiles()
		//camera.DrawCameraMarker()
		player.DrawPlayer()

		for i := 0; i < len(game.Collectables); i++ {
			if !game.Collectables[i].Collected {
				game.Collectables[i].DrawCollectable()
			}
		}

		rl.DrawCircle(int32(game.Camera.screenLeft), int32(game.Camera.screenTop), 20, rl.Blue)
		rl.DrawCircle(int32(game.Camera.screenRight), int32(game.Camera.screenBottom), 20, rl.Red)

		rl.EndMode2D()

		rl.EndDrawing()
	}
}
