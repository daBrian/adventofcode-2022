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
	log.Printf("3a Total scores: %v", result)

	s2, err := internal.LineScannerFromFile("./3/input.txt")
	defer s2.Close()
	if err != nil {
		log.Panic(err)
	}
	result, err = calculateGroupScore(s2, false)
	if err != nil {
		panic(err)
	}
	log.Printf("3b Group scores: %v", result)
}

func calculateScore(s *internal.LineScanner, debug bool) (sum int, err error) {
	for s.Scan() {
		line := s.Text()
		char, err := locateDoubleItem(line)
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
func calculateGroupScore(s *internal.LineScanner, debug bool) (sum int, err error) {
	for s.Scan() {
		var lines [3]string
		lines[0] = s.Text()
		if !s.Scan() {
			return 0, errors.New("unexpected end of list")
		}
		lines[1] = s.Text()
		if !s.Scan() {
			return 0, errors.New("unexpected end of list")
		}
		lines[2] = s.Text()
		char, err := locateGroupItem(lines)
		if err != nil {
			panic(err)
		}
		score, err := calculateRuneScore(char)
		if err != nil {
			panic(err)
		}
		sum += score
		if debug {
			log.Printf("%c %v\t%v", char, score, lines)
		}
	}
	return
}

func locateDoubleItem(line string) (rune, error) {
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
func locateGroupItem(lines [3]string) (r rune, err error) {
	var sb strings.Builder
	for _, line := range lines {
		cleanLine, err := eliminateDuplicate(line)
		if err != nil {
			panic(err)
		}
		sb.WriteString(cleanLine)
	}
	runes := make(map[rune]int)
	for _, r := range sb.String() {
		occurence := runes[r]
		if occurence == 2 {
			return r, nil
		}
		runes[r] = occurence + 1
	}
	return 0, errors.New("Missing common item")
}

func eliminateDuplicate(line string) (string, error) {
	m := map[rune]bool{}
	for _, r := range line {
		m[r] = true
	}
	var sb strings.Builder
	sb.Grow(len(m))
	for r := range m {
		sb.WriteRune(r)
	}
	return sb.String(), nil
}
