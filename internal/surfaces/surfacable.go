package surfaces

type Surfacable interface {
	Get(x, y int) bool
	Set(x, y int, value bool) bool
	GetWidth() int
	GetHeight() int
}

type Coords struct {
	X, Y int
}
