package utils

import (
	"container/list"
	"fmt"
	"time"

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

type ConsoleMessage struct {
	Title string
	Color string
	Info  *list.List
}

func (c *ConsoleMessage) SetTitle(title string) *ConsoleMessage {
	c.Title = title
	return c
}

func (c *ConsoleMessage) SetColor(color string) *ConsoleMessage {
	c.Color = color
	return c
}

func (c *ConsoleMessage) AddInfo(key, value string) *ConsoleMessage {
	c.Info.PushBack([]string{key, value})
	return c
}

func (c *ConsoleMessage) AddInfoConditional(key, value string, condition bool) *ConsoleMessage {
	if condition {
		c.AddInfo(key, value)
	}
	return c
}

func (c *ConsoleMessage) Show() {
	r, g, b := Misc.HexToRGB(c.Color)
	timeFormat := Console.Colored(time.Now().Local().Format("15:04:05"), 193, 239, 255, false)
	message := Console.Colored("[", 238, 238, 238, false) + timeFormat + Console.Colored("] ", 238, 238, 238, false) + Console.Colored(c.Title, r, g, b, false) + "\n"

	for e := c.Info.Front(); e != nil; e = e.Next() {
		info := e.Value.([]string)
		if info[0] != "" {
			message += Console.Colored("  "+info[0]+": ", 238, 238, 238, false) + Console.Colored(info[1], 193, 239, 255, false) + "\n"
		} else {
			message += Console.Colored("  "+info[1], 193, 239, 255, false) + "\n"
		}
	}

	fmt.Println(message)
}

func (ConsoleUtils) NewConsoleMessage() *ConsoleMessage {
	return &ConsoleMessage{
		Info: list.New(),
	}
}

var Console = &ConsoleUtils{}
