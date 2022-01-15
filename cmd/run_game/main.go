package main

import (
	"math/rand"
	"time"

	"github.com/Alexander-Andrade/go_life_game/internal/cells"
	"github.com/Alexander-Andrade/go_life_game/internal/generators"
	"github.com/Alexander-Andrade/go_life_game/internal/surfaces"
)

const width = 60
const height = 30

func main() {
	rand.Seed(time.Now().UnixNano())
	surface := surfaces.CreateTorSurface(width, height)
	nextSurface := surfaces.CreateTorSurface(width, height)
	generators.GenerateCells(surface)

	surfaces.Draw(surface)

	for {
		surfaces.Draw(surface)

		for y := 0; y < surface.GetHeight(); y++ {
			for x := 0; x < surface.GetWidth(); x++ {
				nextSurface.Set(x, y, cells.CalculateNextCellState(surfaces.Coords{X: x, Y: y}, surface))
			}
		}

		surface, nextSurface = nextSurface, surface
		time.Sleep(1 * time.Second)
	}
}
