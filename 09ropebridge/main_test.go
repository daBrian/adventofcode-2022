package main

import (
	. "github.com/daBrian/adventofcode-2022/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleScanner9a = LineScannerFromString(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2

`)
var exampleScanner9b = LineScannerFromString(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`)

func Test9a(t *testing.T) {
	positions := make(map[tailKnot]bool)
	head := headKnot{}
	tail := []tailKnot{tailKnot{knot{0, 0}}}
	err := scanAndMoveAll(exampleScanner9a, head, tail, positions, false)
	assert.NoError(t, err)
	assert.Equal(t, 13, len(positions))
}

func Test9b(t *testing.T) {
	tests := []struct {
		name              string
		inputScanner      *LineScanner
		lengthOfTails     int
		expectedPositions int
	}{
		{name: "9a", inputScanner: exampleScanner9a, lengthOfTails: 1, expectedPositions: 13},
		{name: "9b", inputScanner: exampleScanner9b, lengthOfTails: 9, expectedPositions: 36},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			positions := make(map[tailKnot]bool)
			head := headKnot{}
			tails := make([]tailKnot, test.lengthOfTails)
			for i := range tails {
				tails[i] = tailKnot{knot{x: 0, y: 0}}
			}
			err := scanAndMoveAll(test.inputScanner, head, tails, positions, false)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedPositions, len(positions))
		})
	}
}
