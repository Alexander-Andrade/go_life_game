package surfaces

import (
	"fmt"
	"testing"
)

func TestCorrectedIndex(t *testing.T) {
	axisSize := 3
	var testCases = []struct {
		ind    int
		result int
	}{
		{2, 2},
		{3, 0},
		{4, 1},
		{6, 0},
		{-1, 2},
		{-2, 1},
		{-4, 2},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("When index within boundaries: %d", testCase.ind)

		t.Run(testname, func(t *testing.T) {
			index := correctedIndex(testCase.ind, axisSize)
			if index != testCase.result {
				t.Errorf("CorrectedIndex(%d, %d) = %d; want %d", testCase.ind, axisSize,
					index, testCase.result)
			}
		})
	}
}

func TestGet(t *testing.T) {
	torSurface := TorSurface{Data: make([]bool, 12), Width: 3, Height: 4}
	torSurface.Data[torSurface.Width*3+2] = true
	var testCases = []struct {
		x, y   int
		result bool
	}{
		{2, 3, true},
		{5, -1, true},
		{3, 3, false},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("Get the cell with coords: %d, %d", testCase.x, testCase.y)

		t.Run(testname, func(t *testing.T) {
			result := torSurface.Get(testCase.x, testCase.y)
			if result != testCase.result {
				t.Errorf("Get(%d, %d) = %v; want %v", testCase.x, testCase.y, result, testCase.result)
			}
		})
	}
}

func TestSet(t *testing.T) {
	width := 3
	height := 4
	x := 2
	y := 3

	torSurface := TorSurface{Data: make([]bool, width*height), Width: width, Height: height}
	torSurface.Set(x, y, true)
	if torSurface.Data[torSurface.Width*y+x] != true {
		t.Errorf("Get(%d, %d) = %v; want %v", x, y, false, true)
	}
}
