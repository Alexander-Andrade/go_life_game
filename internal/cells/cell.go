package cells

import (
	"github.com/Alexander-Andrade/go_life_game/internal/surfaces"
)

func CalculateNextCellState(coords surfaces.Coords, surface surfaces.Surfacable) bool {
	neighboursCount := aliveNeigboursCount(coords, surface)

	if isAlive(coords, surface) && isLonely(neighboursCount) {
		return false
	} else if isAlive(coords, surface) && isOverpopulated(neighboursCount) {
		return false
	} else if isAlive(coords, surface) && hasOptimalNeighboursCount(neighboursCount) {
		return true
	} else if isDead(coords, surface) && canReproduceItself(neighboursCount) {
		return true
	}

	return false
}

func isAlive(coords surfaces.Coords, surface surfaces.Surfacable) bool {
	return surface.Get(coords.X, coords.Y)
}

func isDead(coords surfaces.Coords, surface surfaces.Surfacable) bool {
	return !surface.Get(coords.X, coords.Y)
}

func isLonely(neighboursCount int) bool {
	return neighboursCount < 2
}

func isOverpopulated(neighboursCount int) bool {
	return neighboursCount > 3
}

func hasOptimalNeighboursCount(neighboursCount int) bool {
	return neighboursCount == 2 || neighboursCount == 3
}

func canReproduceItself(neighboursCount int) bool {
	return neighboursCount == 3
}

func aliveNeigboursCount(coords surfaces.Coords, surface surfaces.Surfacable) int {
	neighborsCount := 0

	for i := coords.X - 1; i <= coords.X+1; i++ {
		for j := coords.Y - 1; j <= coords.Y+1; j++ {
			if i == coords.X && j == coords.Y {
				continue
			}

			if surface.Get(i, j) {
				neighborsCount += 1
			}
		}
	}

	return neighborsCount
}
