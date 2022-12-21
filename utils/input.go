package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type InputUtils struct{}

var stdInReader = bufio.NewReader(os.Stdin)

func (InputUtils) GetInputString(prompt string) string {
	Input.PrintPrompt(prompt)

	input, err := stdInReader.ReadString('\n')
	if err != nil {
		Logging.Error("Error reading input:", err)
		return ""
	}

	return strings.TrimSpace(input)
}

func (InputUtils) GetInputInt(prompt string) int {
	return 0
}

func (InputUtils) PrintPrompt(prompt string) {
	t := time.Now().Local().Format("02/01/2006 15:04:05")
	fmt.Printf("[\x1b[34;1m%s\x1b[0m] [\x1b[35;1mINPUT\x1b[0m] %s: ", t, prompt)
}

var Input = &InputUtils{}
