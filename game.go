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
}

func NewGame() Game {
	return Game{
		loaded: false,
	}
}

func (g *Game) PopulateGame() {
	g.Camera = NewCamera(g)
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
		}
	}
	g.loaded = true
}

func (g *Game) DrawTiles() {
	for i := StepDown(g.Camera.screenLeft, CELL_SIZE); i <= StepUp(g.Camera.screenRight, CELL_SIZE); i += CELL_SIZE {
		for j := StepDown(g.Camera.screenTop, CELL_SIZE); j <= StepUp(g.Camera.screenBottom, CELL_SIZE); j += CELL_SIZE {
			if g.Tiles[int(i/CELL_SIZE)][int(j/CELL_SIZE)] {
				rl.DrawRectangle(int32(i), int32(j), CELL_SIZE, CELL_SIZE, rl.Gray)
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
