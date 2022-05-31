package printcolor

import (
	"github.com/fatih/color"
)

func PrintlnR(str string) {
	if len(str) == 0 {
		return
	}

	color.Red(str)
}

func PrintlnB(str string) {
	if len(str) == 0 {
		return
	}

	color.Blue(str)
}

func PrintlnY(str string) {
	if len(str) == 0 {
		return
	}
	color.Yellow(str)
}
