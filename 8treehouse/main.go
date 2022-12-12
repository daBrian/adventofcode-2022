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
	return fmt.Sprintf("%v/%v: %v", t.height, t.posx, t.posy)
}

func r7a() {
	s, err := LineScannerFromFile("./7tree/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	wood, err := loadWood(s)
	_ = wood
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
