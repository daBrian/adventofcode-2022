package main

import (
	"errors"
	"fmt"
	. "github.com/daBrian/adventofcode-2022/internal"
	"log"
	"strconv"
	"strings"
)

func main() {

	run5a()
}

func run5a() {
	s, err := LineScannerFromFile("./5/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	stacks, err := buildAndRearrangeStacks(s)
	if err != nil {
		panic(err)
	}

	log.Printf("5a Top line: %v", peeksStacks(stacks))

}

func peeksStacks(stacks []Stack[crane]) string {
	var b strings.Builder
	for _, c := range stacks {
		top, hasValue := c.Peek()
		if !hasValue {
			panic("unexpected top")
		}
		b.WriteString(string(top))
	}
	return b.String()
}

func buildAndRearrangeStacks(s *LineScanner) (stacks []Stack[crane], err error) {
	stacks, err = parseStacks(s)
	if err != nil {
		return nil, err
	}
	err = reArrangeStacks(stacks, s)
	return
}

func reArrangeStacks(stacks []Stack[crane], s *LineScanner) error {
	for s.Scan() {
		err := move(s.Text(), stacks, s.LineNumber())
		if err != nil {
			return err
		}
	}
	return nil
}

func move(order string, stacks []Stack[crane], lineNumber int) error {
	actions := strings.Fields(order)
	n, err := strconv.Atoi(actions[1])
	if err != nil {
		return fmt.Errorf("Could not parse iterations amount from line '%v'", order)
	}
	from, err := strconv.Atoi(actions[3])
	if err != nil {
		return fmt.Errorf("Could not parse 'from' number from line '%v'", order)
	}
	to, err := strconv.Atoi(actions[5])
	if err != nil {
		return fmt.Errorf("Could not parse 'to' number from line '%v'", order)
	}
	for i := 0; i < n; i++ {
		c, remaining := stacks[from-1].Pop()
		if !remaining {
			return fmt.Errorf("Stack %v is empty in line %v", from, lineNumber)
		}
		stacks[to-1].Push(c)
	}
	return nil
}

type crane string

func parseStacks(s *LineScanner) (craneStacks []Stack[crane], err error) {
	lines := Stack[string]{}
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			break
		}
		lines.Push(line)
	}
	craneStacks, err = initializeCraneStacksFromLines(lines)
	if err != nil {
		return craneStacks, err
	}
	return craneStacks, nil
}

func initializeCraneStacksFromLines(lines Stack[string]) ([]Stack[crane], error) {
	nextLine, hasLine := lines.Pop()
	if !hasLine {
		return nil, errors.New("unexpected end of file")
	}
	nOfStacks := len(strings.Fields(nextLine))
	stacks := make([]Stack[crane], nOfStacks)
	for {
		nextLine, hasLine = lines.Pop()

		if !hasLine || len(nextLine) == 0 {
			break
		}
		chunks := ChunkString(nextLine, 4)
		for i, chunk := range chunks {
			crane, exists := newCrane(chunk)
			if exists {
				stacks[i].Push(*crane)
			}
		}
	}

	return stacks, nil
}

func newCrane(chunk string) (*crane, bool) {
	chunk = strings.TrimSpace(chunk)
	if len(chunk) == 0 {
		return nil, false
	}
	c := crane(chunk[1])
	return &c, true
}
