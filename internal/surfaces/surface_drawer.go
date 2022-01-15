package surfaces

import "fmt"

func Draw(surface Surfacable) {
	clearConsole()

	for y := 0; y < surface.GetHeight(); y++ {
		for x := 0; x < surface.GetWidth(); x++ {
			printCell(x, y, surface)
			printNewLine(x, surface)
		}
	}
}

func clearConsole() {
	fmt.Print("\033[2J")
}

func printCell(x, y int, surface Surfacable) {
	if surface.Get(x, y) {
		fmt.Print("@ ")
	} else {
		fmt.Print("  ")
	}
}

func printNewLine(i int, surface Surfacable) {
	if ((i + 1) % surface.GetWidth()) != 0 {
		return
	}

	fmt.Println()
}
