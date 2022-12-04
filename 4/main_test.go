package main

import (
	"github.com/daBrian/adventofcode-2022/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countEnclosingPairs(t *testing.T) {
	ls, _ := internal.LineScannerFromString(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)
	got, err := countEnclosingPairs(ls)
	assert.NoError(t, err)
	assert.Equal(t, 2, got)
}
