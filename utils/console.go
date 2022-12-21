package utils

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

type ConsoleUtils struct{}

func (ConsoleUtils) Clear() {
	fmt.Print("\033[H\033[2J")
}

func (ConsoleUtils) PrintBanner() {
	Console.Clear()

	banner := String.CenterSlices(figure.NewFigure("LightningBot", "stop", true).Slicify())

	fmt.Println(Console.Colored(banner, 0, 203, 255, true) + "\n\n")
}

func (ConsoleUtils) Colored(text string, r, g, b int, bold bool) string {
	boldArg := ";1"
	if !bold {
		boldArg = ""
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%d%sm%s\033[0m", r, g, b, boldArg, text)
}

var Console = &ConsoleUtils{}
