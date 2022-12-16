package main

import (
	"fmt"
	. "github.com/daBrian/adventofcode-2022/internal"
	"log"
)

func main() {
	r7a()
}

type tree struct {
	height int
	posx   int
	posy   int
}

func (t tree) String() string {
	return fmt.Sprintf("%v/%v: %v", t.posx, t.posy, t.height)
}

func r7a() {
	s, err := LineScannerFromFile("./8treehouse/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	wood, err := loadWood(s)
	visibleTrees := findVisibleTrees(wood)
	fmt.Printf("Found %v trees", len(visibleTrees))
}

func findVisibleTrees(wood [][]tree) map[tree]any {
	var visibleTrees = make(map[tree]any)

	for i := 0; i < 4; i++ {
		//show(wood)
		findVisibleTreesFromLeft(wood, &visibleTrees)
		wood = transpose(wood)
	}
	return visibleTrees
}

func show(wood [][]tree) {
	for _, trees := range wood {
		for _, t := range trees {
			fmt.Printf("%v | ", t)
		}
		println()
	}
	println()
}
func findVisibleTreesFromLeft(wood [][]tree, trees *map[tree]any) {
	for y, line := range wood {
		maxInCol := -1
		for x, t := range line {
			if t.height > maxInCol {
				(*trees)[t] = nil
				maxInCol = t.height
				_, _ = x, y
			}
		}

	}
	return
}

func transpose(matrix [][]tree) [][]tree {
	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}
func loadWood(s *LineScanner) (wood [][]tree, err error) {
	wood = [][]tree{}
	for s.Scan() {
		line, err := loadTreeLine(s.Text(), s.LineNumber()-1)
		if err != nil {
			return nil, fmt.Errorf("error while parsing line %v", s.LineNumber())
		}
		wood = append(wood, line)
	}
	return wood, nil
}

func loadTreeLine(text string, y int) ([]tree, error) {
	var line = make([]tree, len(text))
	for x, c := range text {
		t, err := newTree(c, x, y)
		if err != nil {
			return nil, err
		}
		line[x] = t
	}
	return line, nil
}

func newTree(r rune, x int, y int) (tree, error) {
	return tree{int(r - '0'), x, y}, nil
}
