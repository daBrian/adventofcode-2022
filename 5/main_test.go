package main

import (
	"github.com/daBrian/adventofcode-2022/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildAndRearrangeStacks(t *testing.T) {
	ls, _ := internal.LineScannerFromString(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)
	stacks, err := buildAndRearrangeStacks(ls, false)
	got := peeksStacks(stacks)
	assert.NoError(t, err)
	assert.Equal(t, "CMZ", got)
}
func Test_buildAndRearrangeStacks9001(t *testing.T) {
	ls, _ := internal.LineScannerFromString(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)
	stacks, err := buildAndRearrangeStacks(ls, true)
	got := peeksStacks(stacks)
	assert.NoError(t, err)
	assert.Equal(t, "MCD", got)
}
