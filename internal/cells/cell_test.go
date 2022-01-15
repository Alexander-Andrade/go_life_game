package cells

import (
	"fmt"
	"testing"

	"github.com/Alexander-Andrade/go_life_game/internal/surfaces"
)

func TestAliveNeigboursCount(t *testing.T) {
	width := 3
	height := 4
	torSurface := surfaces.TorSurface{
		Data: []bool{
			true, false, true,
			false, true, false,
			false, true, false,
			false, false, false,
		},
		Width:  width,
		Height: height,
	}
	var testCases = []struct {
		x, y           int
		neighborsCount int
	}{
		{2, 3, 3},
		{1, 1, 3},
		{1, 2, 1},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("neighbours count for cell: %d, %d", testCase.x, testCase.y)

		t.Run(testname, func(t *testing.T) {
			neighborsCount := aliveNeigboursCount(surfaces.Coords{X: testCase.x, Y: testCase.y}, &torSurface)
			if neighborsCount != testCase.neighborsCount {
				t.Errorf("aliveNeigboursCount(%d, %d) = %d; want %d", testCase.x, testCase.y, neighborsCount, testCase.neighborsCount)
			}
		})
	}
}

func TestIsAlive(t *testing.T) {
	torSurface := surfaces.TorSurface{
		Data: []bool{
			true, false, true,
			false, true, false,
			false, true, false,
			false, false, false,
		},
		Width:  3,
		Height: 4,
	}

	result := isAlive(surfaces.Coords{X: 1, Y: 1}, &torSurface)
	if result != true {
		t.Errorf("isAlive(coords, surface) = %t; want %t", result, true)
	}

	result = isAlive(surfaces.Coords{X: 0, Y: 1}, &torSurface)
	if result != false {
		t.Errorf("isAlive(coords, surface) = %t; want %t", result, false)
	}
}

func TestIsDead(t *testing.T) {
	torSurface := surfaces.TorSurface{
		Data: []bool{
			true, false, true,
			false, true, false,
			false, true, false,
			false, false, false,
		},
		Width:  3,
		Height: 4,
	}

	result := isDead(surfaces.Coords{X: 1, Y: 1}, &torSurface)
	if result != false {
		t.Errorf("isDead(coords, surface) = %t; want %t", result, false)
	}

	result = isDead(surfaces.Coords{X: 0, Y: 1}, &torSurface)
	if result != true {
		t.Errorf("isDead(coords, surface) = %t; want %t", result, true)
	}
}

func TestIsLonely(t *testing.T) {
	result := isLonely(1)
	if result != true {
		t.Errorf("isLonely(1) = %t; want %t", result, true)
	}

	result = isLonely(3)
	if result != false {
		t.Errorf("isLonely(3) = %t; want %t", result, false)
	}
}

func TestIsOverpopulated(t *testing.T) {
	result := isOverpopulated(1)
	if result != false {
		t.Errorf("isOverpopulated(1) = %t; want %t", result, false)
	}

	result = isOverpopulated(4)
	if result != true {
		t.Errorf("isOverpopulated(4) = %t; want %t", result, true)
	}
}

func TestHasOptimalNeighboursCount(t *testing.T) {
	result := hasOptimalNeighboursCount(2)
	if result != true {
		t.Errorf("hasOptimalNeighboursCount(2) = %t; want %t", result, true)
	}

	result = hasOptimalNeighboursCount(3)
	if result != true {
		t.Errorf("hasOptimalNeighboursCount(3) = %t; want %t", result, true)
	}

	result = hasOptimalNeighboursCount(1)
	if result != false {
		t.Errorf("hasOptimalNeighboursCount(3) = %t; want %t", result, false)
	}
}

func TestCanReproduceItself(t *testing.T) {
	result := canReproduceItself(3)
	if result != true {
		t.Errorf("canReproduceItself(3) = %t; want %t", result, true)
	}

	result = canReproduceItself(1)
	if result != false {
		t.Errorf("canReproduceItself(1) = %t; want %t", result, false)
	}

	result = canReproduceItself(4)
	if result != false {
		t.Errorf("canReproduceItself(4) = %t; want %t", result, false)
	}
}

func TestCalculateNextCellState(t *testing.T) {
	torSurface := surfaces.TorSurface{
		Data: []bool{
			true, false, true,
			false, true, false,
			false, true, false,
			false, false, false,
		},
		Width:  3,
		Height: 4,
	}

	var testCases = []struct {
		x, y          int
		nextCellState bool
	}{
		// alive and 3 neighbours
		{1, 1, true},
		// alive and 1 neighbour
		{1, 2, false},
		// dead and 3 neighbours
		{1, 3, true},
		// dead and 4 neighbours
		{0, 1, false},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("cell coords: %d, %d", testCase.x, testCase.y)

		t.Run(testname, func(t *testing.T) {
			resultCellState := CalculateNextCellState(surfaces.Coords{X: testCase.x, Y: testCase.y}, &torSurface)
			if resultCellState != testCase.nextCellState {
				t.Errorf("CalculateNextCellState(%d, %d) = %t; want %t", testCase.x, testCase.y, resultCellState, testCase.nextCellState)
			}
		})
	}

	t.Run("overpopulated cell", func(t *testing.T) {
		torSurface := surfaces.TorSurface{
			Data: []bool{
				true, true, true,
				false, true, false,
				false, true, false,
				false, false, false,
			},
			Width:  3,
			Height: 4,
		}

		resultCellState := CalculateNextCellState(surfaces.Coords{X: 1, Y: 1}, &torSurface)
		if resultCellState != false {
			t.Errorf("CalculateNextCellState(1, 1) = %t; want %t", resultCellState, false)
		}
	})
}
