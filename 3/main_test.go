package main

import (
	"fmt"
	"github.com/daBrian/adventofcode-2022/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculateScore(t *testing.T) {
	ls, _ := internal.LineScannerFromString(
		`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)
	sum, _ := calculateScore(ls, true)
	assert.Equal(t, 157, sum)
}
func Test_LocateItem(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp"
	doubleItem, err := locateItem(input)
	assert.NoError(t, err)
	assert.Equal(t, 'p', doubleItem)
}

func Test_calculateRuneScore(t *testing.T) {
	tests := []struct {
		input  rune
		expect int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
		{'p', 16},
		{'L', 38},
		{'P', 42},
		{'v', 22},
		{'t', 20},
		{'s', 19},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%c", test.input), func(t *testing.T) {
			score, err := calculateRuneScore(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.expect, score)
		})
	}
}
