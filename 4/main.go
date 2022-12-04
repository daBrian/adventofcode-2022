package main

import (
	"fmt"
	"github.com/daBrian/adventofcode-2022/internal"
	"log"
	"strconv"
	"strings"
)

func main() {
	run4a()
}

func run4a() {
	s, err := internal.LineScannerFromFile("./4/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	result, err := countEnclosingPairs(s)
	if err != nil {
		panic(err)
	}
	log.Printf("3a Total scores: %v", result)

}

type elve struct {
	start int
	end   int
}

func (e elve) covers(other *elve) bool {
	return e.start <= other.start && e.end >= other.end
}

func newElve(s string) (*elve, error) {
	scope := strings.Split(s, "-")
	if len(scope) != 2 {
		return nil, fmt.Errorf("no valid range: %v", s)
	}
	start, err := strconv.Atoi(scope[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(scope[1])
	if err != nil {
		return nil, err
	}
	return &elve{start, end}, nil
}

func countEnclosingPairs(s *internal.LineScanner) (int, error) {
	var count int
	for s.Scan() {
		first, second, err := getPairs(s.Text())
		if err != nil {
			return 0, err
		}
		if first.covers(second) || second.covers(first) {
			count++
		}
	}
	return count, nil
}

func getPairs(line string) (first *elve, second *elve, err error) {
	sep := strings.Split(line, ",")
	if len(sep) != 2 {
		return nil, nil, fmt.Errorf("%v contains %v elves instead of 2", line, len(sep))
	}
	first, err = newElve(sep[0])
	if err != nil {
		return nil, nil, err
	}
	second, err = newElve(sep[1])
	if err != nil {
		return nil, nil, err
	}
	return first, second, nil
}
