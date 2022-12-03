package main

import (
	"errors"
	"fmt"
	"github.com/daBrian/adventofcode-2022/internal"
	"log"
	"strings"
)

func main() {
	s, err := internal.LineScannerFromFile("./3/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	result, err := calculateScore(s, false)
	if err != nil {
		panic(err)
	}
	log.Printf("Total scores: %v", result)
}

func calculateScore(s *internal.LineScanner, debug bool) (sum int, err error) {
	for s.Scan() {
		line := s.Text()
		char, err := locateItem(line)
		if err != nil {
			panic(err)
		}
		score, err := calculateRuneScore(char)
		if err != nil {
			panic(err)
		}
		sum += score
		if debug {
			log.Printf("%c %v\t%v", char, score, line)
		}
	}
	return
}

func locateItem(line string) (rune, error) {
	for _, c := range line[0 : len(line)/2] {
		if strings.Count(line[len(line)/2:], string(c)) > 0 {
			return c, nil
		}
	}
	return 0, fmt.Errorf("No double item found in %v", line)
}

func calculateRuneScore(c rune) (priority int, err error) {
	priority = int(c)
	if priority > 96 && priority < 123 {
		return priority - 96, nil
	} else if priority > 64 && priority < 91 {
		return priority - 64 + 26, nil
	}
	return 0, errors.New("")
}
