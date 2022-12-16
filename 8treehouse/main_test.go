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

func TestFindHiddenTrees(t *testing.T) {
	woods, _ := loadWood(exampleScanner)
	found := findVisibleTrees(woods)
	//for t2, _ := range found {
	//	fmt.Println(t2)
	//}
	assert.NotNil(t, found)
	assert.Equal(t, 21, len(found))
}

func TestFindBestScenicScore(t *testing.T) {
	woods, _ := loadWood(exampleScanner)
	score := findBestScenicScore(woods)
	assert.Equal(t, 8, score)
}

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		name          string
		posX          int
		posY          int
		expectedScore int
	}{
		{"no points at the top", 1, 0, 0},
		{"no points at the right", 0, 1, 0},
		{"no points at the bottom", 1, 4, 0},
		{"no points at the bottom", 4, 1, 0},
		{"example 1 - four pointer", 2, 1, 4},
		{"example 2 - eight pointer", 2, 3, 8},
	}
	woods, _ := loadWood(exampleScanner)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := woods[test.posY][test.posX]
			assert.Equal(t, test.expectedScore, tree.calcuclateScore(&woods))
		})
	}
}

func TestCountUp(t *testing.T) {
	tests := []struct {
		name          string
		posX          int
		posY          int
		expectedCount int
	}{
		{"no points at the top", 1, 0, 0},
		{"one point in the second row", 2, 1, 1},
		{"two points in the third row", 0, 2, 2},
		{"one point in the first example", 2, 1, 1},
		{"two points in the second example", 2, 3, 2},
	}
	woods, _ := loadWood(exampleScanner)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := woods[test.posY][test.posX]
			assert.Equal(t, test.expectedCount, tree.countUp(woods, tree.height))
		})
	}
}
func TestCountLeft(t *testing.T) {
	tests := []struct {
		name          string
		posX          int
		posY          int
		expectedCount int
	}{
		{"no points on left", 0, 2, 0},
		{"one point in the second col", 1, 4, 1},
		{"two points in the third col", 2, 3, 2},
		{"one point in the first example", 2, 1, 1},
		{"two points in the second example", 2, 3, 2},
	}
	woods, _ := loadWood(exampleScanner)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := woods[test.posY][test.posX]
			assert.Equal(t, test.expectedCount, tree.countLeft(woods, tree.height))
		})
	}
}

func TestCountRight(t *testing.T) {
	tests := []struct {
		name          string
		posX          int
		posY          int
		expectedCount int
	}{
		{"no points on the right", 4, 2, 0},
		{"one point in the 2nd to right", 3, 0, 1},
		{"two points in the 3rd to right", 2, 1, 2},
		{"two point in the first example", 2, 1, 2},
		{"two points in the second example", 2, 3, 2},
	}
	woods, _ := loadWood(exampleScanner)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := woods[test.posY][test.posX]
			assert.Equal(t, test.expectedCount, tree.countRight(woods, tree.height))
		})
	}
}

func TestCountDown(t *testing.T) {
	tests := []struct {
		name          string
		posX          int
		posY          int
		expectedCount int
	}{
		{"no points on the Bottom", 2, 4, 0},
		{"one point in the 2nd to Bottom", 0, 3, 1},
		{"two points in the 3rd to Bottom", 2, 1, 2},
		{"two point in the first example", 2, 1, 2},
		{"two points in the second example", 2, 3, 1},
	}
	woods, _ := loadWood(exampleScanner)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := woods[test.posY][test.posX]
			assert.Equal(t, test.expectedCount, tree.countDown(woods, tree.height))
		})
	}
}
