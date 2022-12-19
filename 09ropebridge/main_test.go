package main

import (
	"github.com/daBrian/adventofcode-2022/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleScanner = internal.LineScannerFromString(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2

`)

func Test(t *testing.T) {
	positions := make(map[tailKnot]bool)
	head := headKnot{}
	tail := tailKnot{0, 0}
	err := scanAndMoveAll(exampleScanner, head, tail, positions)
	assert.NoError(t, err)
	assert.Equal(t, 13, len(positions))
}
