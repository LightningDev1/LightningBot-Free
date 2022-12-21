package utils

import (
	"os/exec"
	"runtime"
)

type MiscUtils struct{}

// Returns a if a is not empty, otherwise returns b
func (MiscUtils) Or(a string, b string) string {
	if a != "" {
		return a
	}
	return b
}

// Returns lowest of three integers
func (MiscUtils) Minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

// GetDistance returns the Levenshtein distance between two strings.
func (MiscUtils) GetDistance(str1, str2 string) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastKey := x - 1
		for y := 1; y <= s1len; y++ {
			oldKey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = Misc.Minimum(column[y]+1, column[y-1]+1, lastKey+incr)
			lastKey = oldKey
		}
	}
	return column[s1len]
}

func (MiscUtils) OpenURL(url string) error {
	if runtime.GOOS == "windows" {
		return exec.Command("explorer", url).Start()
	} else if runtime.GOOS == "darwin" {
		return exec.Command("open", url).Start()
	} else {
		return exec.Command("xdg-open", url).Start()
	}
}

var Misc = &MiscUtils{}
