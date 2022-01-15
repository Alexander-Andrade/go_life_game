package surfaces

import (
	"math"
)

type TorSurface struct {
	Data   []bool
	Width  int
	Height int
}

func CreateTorSurface(width, height int) *TorSurface {
	return &TorSurface{Data: make([]bool, width*height), Width: width, Height: height}
}

func (surface *TorSurface) Get(x, y int) bool {
	xCorrected := correctedIndex(x, surface.Width)
	yCorrected := correctedIndex(y, surface.Height)
	index := surface.Width*yCorrected + xCorrected
	return surface.Data[index]
}

func (surface *TorSurface) Set(x, y int, value bool) bool {
	xCorrected := correctedIndex(x, surface.Width)
	yCorrected := correctedIndex(y, surface.Height)
	index := surface.Width*yCorrected + xCorrected

	surface.Data[index] = value
	return surface.Data[index]
}

func (surface *TorSurface) GetWidth() int {
	return surface.Width
}

func (surface *TorSurface) GetHeight() int {
	return surface.Height
}

func correctedIndex(ind int, axisSize int) int {
	if ind >= 0 && ind < axisSize {
		return ind
	}

	if axisSize <= ind {
		return ind % axisSize
	}

	return axisSize - (int(math.Abs(float64(ind))) % axisSize)
}
