package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const CELL_SIZE = 64

type Game struct {
	loaded bool
	Tiles  [200][200]bool
	Camera Camera
	Player Player
	Input  Input
}

func NewGame() Game {
	return Game{
		loaded: false,
	}
}

func (g *Game) PopulateGame() {
	g.Camera = NewCamera(g)
	g.Player = NewPlayer(g)
	g.Input = NewInput(g)
}

func (g *Game) LoadLevel(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	level := string(buf[:n-1])

	split := strings.Split(level, ",")

	g.Tiles = [200][200]bool{}

	for i := 0; i < len(split)-1; i += 3 {
		x, _ := strconv.Atoi(split[i])
		y, _ := strconv.Atoi(split[i+1])
		tile, _ := strconv.Atoi(split[i+2])

		if tile == 1 {
			g.Tiles[x][y] = true
		} else if tile == 2 {
			g.Player.Pos = rl.NewVector2(float32(x*(CELL_SIZE)), float32(y*CELL_SIZE)-5)
			g.Camera.MoveCamera(g.Player.Pos)
			g.Camera.tweenPos = g.Player.Pos
		}
	}
	g.loaded = true
}

func (g *Game) DrawTiles() {
	for i := max(0, StepDown(g.Camera.screenLeft, CELL_SIZE)); i <= min(200*CELL_SIZE, StepUp(g.Camera.screenRight, CELL_SIZE)); i += CELL_SIZE {
		for j := max(0, StepDown(g.Camera.screenTop, CELL_SIZE)); j <= min(200*CELL_SIZE, StepUp(g.Camera.screenBottom, CELL_SIZE)); j += CELL_SIZE {
			if i >= 0 && j >= 0 && i < 200*CELL_SIZE && j < 200*CELL_SIZE {
				if g.Tiles[int(i/CELL_SIZE)][int(j/CELL_SIZE)] {
					rl.DrawRectangle(int32(i), int32(j), CELL_SIZE, CELL_SIZE, rl.Gray)
				}
			}
		}
	}
}

// general functions

// bring value down to the nearest multiple of step size
func StepDown(val float32, stepSize float32) float32 {
	return Floor32(val/stepSize) * stepSize
}

// bring value up to the nearest multiple of step size
func StepUp(val float32, stepSize float32) float32 {
	return Ceil32(val/stepSize) * stepSize
}

// like floor but with float32 values instead of float64
func Floor32(val float32) float32 {
	return float32(math.Floor(float64(val)))
}

// like floor but with float32 values instead of float64
func Ceil32(val float32) float32 {
	return float32(math.Ceil(float64(val)))
}

// converts a boolean to an integer value for boolean arithmetic
func BoolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}

// move a value to a destination by a given step size
func MoveValue(val float32, dest float32, step float32) float32 {
	orig := dest-val > 0
	if dest-val > 0 {
		val += step
	} else {
		val -= step
	}
	if (dest-val > 0) != orig {
		return dest
	} else {
		return val
	}
}

// returns a normalized version of the value
func Sign(val float32) float32 {
	if val > 0 {
		return 1
	} else if val < 0 {
		return -1
	} else {
		return 0
	}
}

// returns a normalized version of the value
func BoolSign(val bool) float32 {
	if val {
		return 1
	} else {
		return -1
	}
}
