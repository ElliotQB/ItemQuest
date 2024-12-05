package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PLAYING      = 0
	TITLE_SCREEN = 1
)

func main() {
	rl.InitWindow(1920, 1080, "ItemQuest")
	defer rl.CloseWindow()

	game := NewGame()
	game.PopulateGame()
	game.LoadLevel("level.txt")

	camera := &game.Camera
	player := &game.Player
	gameTimer := 0

	rl.InitAudioDevice()
	bgMusic := rl.LoadMusicStream("audio/wavedash.ppt.mp3")
	rl.PlayMusicStream(bgMusic)

	finishSound := rl.LoadSound("audio/collect.ogg")
	finishSoundPlayed := false

	abilitySound := rl.LoadSound("audio/heal.ogg")

	gameState := TITLE_SCREEN
	numCats := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if gameState == TITLE_SCREEN {
			if rl.IsKeyPressed(rl.KeySpace) {
				gameState = PLAYING
			}
		} else {
			if rl.IsKeyPressed(rl.KeyBackspace) {
				gameState = TITLE_SCREEN
			}

			// game logic
			rl.UpdateMusicStream(bgMusic)
			if gameTimer < 10 {
				camera.Camera.Target = camera.tweenPos
			}
			gameTimer++

			numCats = 0

			game.Input.InputStep()
			game.Player.PlayerStep()
			camera.MoveCamera(rl.NewVector2(game.Player.Pos.X+(PLAYER_WIDTH/2), game.Player.Pos.Y+(PLAYER_HEIGHT/2)))
			camera.CameraStep()

			for i := 0; i < len(game.Collectables); i++ {
				if game.Collectables[i].Type == CAT {
					numCats++
				}
			}

			playerRec := rl.Rectangle{game.Player.Pos.X, game.Player.Pos.Y, PLAYER_WIDTH, PLAYER_HEIGHT}
			for i := 0; i < len(game.Collectables); i++ {
				if rl.CheckCollisionCircleRec(game.Collectables[i].Pos, 40, playerRec) {

					if !game.Collectables[i].Collected {

						switch game.Collectables[i].Type {

						case DOUBLEJUMPGEM:
							game.Player.canDoubleJump = true
							game.Player.SetNumJumps()
							rl.PlaySound(abilitySound)

						case TRIPLEJUMPGEM:
							game.Player.canTripleJump = true
							game.Player.SetNumJumps()
							rl.PlaySound(abilitySound)

						case WALLJUMPGEM:
							game.Player.canWallJump = true
							rl.PlaySound(abilitySound)

						case CAT:
							game.Cats++
							rl.PlaySound(abilitySound)

						}
						game.Collectables[i].Collect()
					}
				}
			}
		}

		// drawing
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		if gameState == TITLE_SCREEN {
			rl.DrawText("Item Quest\nWASD/Space to move\nBackspace to return to title\n\nPress space to begin", 30, 30, 40, rl.White)
		} else {

			rl.BeginMode2D(game.Camera.Camera)

			game.DrawTiles()
			//camera.DrawCameraMarker()
			player.DrawPlayer()

			for i := 0; i < len(game.Collectables); i++ {
				if !game.Collectables[i].Collected {
					game.Collectables[i].DrawCollectable()
				}
			}

			//rl.DrawCircle(int32(game.Camera.screenLeft), int32(game.Camera.screenTop), 20, rl.Blue)
			//rl.DrawCircle(int32(game.Camera.screenRight), int32(game.Camera.screenBottom), 20, rl.Red)

			rl.EndMode2D()

			hud := strconv.Itoa(game.Cats) + "/" + strconv.Itoa(numCats)

			if game.Player.canDoubleJump {
				hud += "\nDouble Jump"
			}
			if game.Player.canWallJump {
				hud += "\nWall Jump"
			}
			if game.Player.canTripleJump {
				hud += "\nTriple Jump"
			}

			if game.Cats == numCats && game.Player.canDoubleJump && game.Player.canTripleJump && game.Player.canWallJump {
				if !finishSoundPlayed {
					rl.PlaySound(finishSound)
				}
				hud += "\n\nCOMPLETE!"
			}

			rl.DrawText(hud, 15, 15, 20, rl.White)
		}

		rl.EndDrawing()
	}
}
