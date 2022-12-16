package main

import (
	"fmt"
	. "github.com/daBrian/adventofcode-2022/internal"
	"log"
)

func main() {
	r8a()
	r8b()
}

type tree struct {
	height int
	posx   int
	posy   int
}

func (t tree) String() string {
	return fmt.Sprintf("%v/%v: %v", t.posx, t.posy, t.height)
}

func (t *tree) calcuclateScore(wood *[][]tree) int {
	upTrees := t.countUp(*wood, t.height)
	leftTrees := t.countLeft(*wood, t.height)
	rightTrees := t.countRight(*wood, t.height)
	downTrees := t.countDown(*wood, t.height)
	return upTrees * leftTrees * rightTrees * downTrees
}

func (t *tree) countUp(wood [][]tree, maxHeight int) int {
	if t.posy == 0 {
		return 0
	}
	next := wood[t.posy-1][t.posx]
	if next.height >= maxHeight {
		return 1
	}
	return 1 + next.countUp(wood, maxHeight)
}

func (t *tree) countLeft(wood [][]tree, maxHeight int) int {
	if t.posx == 0 {
		return 0
	}
	next := wood[t.posy][t.posx-1]
	if next.height >= maxHeight {
		return 1
	}
	return 1 + next.countLeft(wood, maxHeight)
}

func (t *tree) countRight(wood [][]tree, maxHeight int) int {
	if t.posx == len(wood[0])-1 {
		return 0
	}
	next := wood[t.posy][t.posx+1]
	if next.height >= maxHeight {
		return 1
	}
	return 1 + next.countRight(wood, maxHeight)
}
func (t *tree) countDown(wood [][]tree, maxHeight int) int {
	if t.posy == len(wood)-1 {
		return 0
	}
	next := wood[t.posy+1][t.posx]
	if next.height >= maxHeight {
		return 1
	}
	return 1 + next.countDown(wood, maxHeight)
}

func r8a() {
	s, err := LineScannerFromFile("./8treehouse/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	wood, err := loadWood(s)
	visibleTrees := findVisibleTrees(wood)
	fmt.Printf("Found %v trees\n", len(visibleTrees))
}

func r8b() {
	s, err := LineScannerFromFile("./8treehouse/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	wood, err := loadWood(s)
	bestTree := findBestScenicScore(wood)
	fmt.Printf("Best scenic score for %v: %v", bestTree, bestTree.calcuclateScore(&wood))
}

func findBestScenicScore(wood [][]tree) tree {
	var bestTree tree
	currentHighscore := 0
	for _, trees := range wood {
		for _, t := range trees {
			newScore := t.calcuclateScore(&wood)
			if newScore > currentHighscore {
				bestTree = t
				currentHighscore = newScore
			}
		}
	}
	return bestTree
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
