package main

import (
	"github.com/daBrian/adventofcode-2022/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleScanner = internal.LineScannerFromString(`30373
25512
65332
33549
35390
`)

func TestWoodLoading(t *testing.T) {
	woods, err := loadWood(exampleScanner)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(woods))
	assert.Equal(t, 5, len(woods[0]))
}