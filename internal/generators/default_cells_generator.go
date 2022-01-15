package generators

import (
	"math/rand"

	"github.com/Alexander-Andrade/go_life_game/internal/surfaces"
)

func GenerateCells(surface surfaces.Surfacable) {
	for y := 0; y < surface.GetHeight(); y++ {
		for x := 0; x < surface.GetWidth(); x++ {
			surface.Set(x, y, rand.Intn(2) == 0)
		}
	}
}
