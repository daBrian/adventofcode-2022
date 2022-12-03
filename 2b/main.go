package main

import (
	"errors"
	"github.com/daBrian/adventofcode-2022/internal"
	"log"
	"unicode/utf8"
)

func main() {
	s, err := internal.LineScannerFromFile("./2/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	result, err := calculateTotalScore(s)
	if err != nil {
		panic(err)
	}
	log.Printf("Total scores: %v", result)
}

type round struct {
	other  rune
	result rune
}

func (b round) check() (err error) {
	switch b.other {
	case 'A', 'B', 'C':
		break
	default:
		return errors.New("Other's round must be one of A B C")
	}
	switch b.result {
	case 'X', 'Y', 'Z':
		break
	default:
		return errors.New("My round must be one of X Y Z")
	}
	return nil

}

func (b round) valueRoundOutcome() int {
	switch b.result {
	case 'X': //loose
		return 0
	case 'Y': //draw
		return 3
	case 'Z': //win
		return 6
	}
	return -1
}

var ruleset = [3][3]int{
	// Rock - 1
	// Paper - 2
	// Scissors - 3
	//  X, Y, Z
	{3, 1, 2}, //A - Rock
	{1, 2, 3}, //B - Paper
	{2, 3, 1}} //C - Scissors

func (b round) valueRoundChoice() int {
	return ruleset[b.other-65][b.result-88]
}

func newRound(currentResult string) (currentRound *round, err error) {
	other, _ := utf8.DecodeRuneInString(currentResult[0:1])
	my, _ := utf8.DecodeRuneInString(currentResult[len(currentResult)-1:])
	b := round{other, my}
	err = b.check()
	if err != nil {
		return nil, err
	}
	return &b, err
}

func calculateTotalScore(s *internal.LineScanner) (int, error) {
	points := 0
	for s.Scan() {
		currentResult := s.Text()
		currentRound, err := newRound(currentResult)
		if err != nil {
			return 0, err
		}
		log.Printf("%v: %v", currentResult, calculateRoundPoints(*currentRound))
		points += calculateRoundPoints(*currentRound)
	}
	return points, nil
}

func calculateRoundPoints(r round) (points int) {
	return r.valueRoundOutcome() + r.valueRoundChoice()
}
