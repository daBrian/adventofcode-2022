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
	doubleItem, err := locateDoubleItem(input)
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
func Test_calculateGroupScore(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{"first group", "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg", 18},
		{"second group", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 52},
		{"both groups", "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 70},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ls, _ := internal.LineScannerFromString(test.input)
			score, err := calculateGroupScore(ls, true)
			assert.NoError(t, err)
			assert.Equal(t, test.expect, score)
		})
	}
}

func Test_eliminateDuplicate(t *testing.T) {
	got, err := eliminateDuplicate("abcbd")
	assert.NoError(t, err)
	assert.Equalf(t, "acbd", got, "eliminateDuplicate(abcbd)")
}
