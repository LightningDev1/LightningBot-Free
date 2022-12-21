package utils

import (
	"os"
	"strings"

	"golang.org/x/term"
)

type StringUtils struct{}

func (StringUtils) CenterSlices(rows []string) string {
	// Get the width of the terminal
	fd := int(os.Stdout.Fd())
	w, _, err := term.GetSize(fd)
	if err != nil {
		// If we can't get the width, just return the rows
		return strings.Join(rows, "\n")
	}

	// Get the length of the longest string
	max := 0
	for _, row := range rows {
		if len(row) > max {
			max = len(row)
		}
	}

	// Center the rows
	spaceCount := (w - max) / 2
	for i, row := range rows {
		rows[i] = strings.Repeat(" ", spaceCount) + row
	}

	return strings.Join(rows, "\n")
}

func (StringUtils) Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func (StringUtils) RemoveWords(text string, amount int) string {
	words := strings.Split(text, " ")
	if amount > len(words) {
		return ""
	}
	return strings.Join(words[amount:], " ")
}

func (StringUtils) Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

var String = &StringUtils{}
