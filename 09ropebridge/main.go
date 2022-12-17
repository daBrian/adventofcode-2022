package main

import (
	"errors"
	. "github.com/daBrian/adventofcode-2022/internal"
	"log"
	"strconv"
	"strings"
)

func main() {
	r9a()
	r9b()
}

type knot struct {
	x, y int
}

type tailKnot knot

func (t tailKnot) follow(h headKnot) {
	if h.x > t.x {
		t.x++
	} else if h.x < t.x {
		t.x--
	}
	if h.y > t.y {
		t.y++
	} else if h.y < t.y {
		t.y--
	}
}

type headKnot struct {
	knot
	traction
}

func (k headKnot) move() {
	if k.xt > 0 {
		k.x++
		k.xt--
	}
	if k.xt < 0 {
		k.x--
		k.xt++
	}
	if k.yt > 0 {
		k.y++
		k.yt--
	}
	if k.yt < 0 {
		k.y--
		k.yt++
	}
}

type traction struct {
	xt, yt int
}

func (t traction) sumUp() int {
	var dx, dy int
	if t.xt < 0 {
		dx = t.xt * -1
	} else {
		dx = t.xt
	}
	if t.yt < 0 {
		dy = t.yt * -1
	} else {
		dy = t.yt
	}
	return dx + dy
}

func loadTractionFromCall(call string) (traction, error) {
	cols := strings.Fields(call)
	steps, err := strconv.Atoi(cols[1])
	if err != nil {
		return traction{}, err
	}
	switch cols[0] {
	case "U":
		return traction{xt: 0, yt: -steps}, nil
	case "D":
		return traction{xt: 0, yt: steps}, nil
	case "L":
		return traction{xt: -steps, yt: 0}, nil
	case "R":
		return traction{xt: steps, yt: 0}, nil
	}
	return traction{}, errors.New(call)
}

func r9a() {
	s, err := LineScannerFromFile("./09ropebridge/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	positions := make(map[tailKnot]bool)
	head := headKnot{}
	tail := tailKnot{0, 0}
	err = scanAndMoveAll(s, head, tail, positions)
	if err != nil {
		panic(err)
	}
}

func scanAndMoveAll(s *LineScanner, head headKnot, tail tailKnot, positions map[tailKnot]bool) (err error) {
	for s.Scan() {
		head.traction, err = loadTractionFromCall(s.Text())
		if err != nil {
			return err
		}
		moveAll(head, tail, positions)
	}
	return nil
}

func moveAll(head headKnot, tail tailKnot, positions map[tailKnot]bool) {
	head.move()
	tail.follow(head)
	positions[tail] = true
	if head.hasTraction() {
		moveAll(head, tail, positions)
	}
}

func (t traction) hasTraction() bool {
	return t.sumUp() > 0
}

func r9b() {
	s, err := LineScannerFromFile("./09ropebridge/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}

}
