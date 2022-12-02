package main

import (
	"github.com/daBrian/adventofcode-2022/internal"
	"log"
	"strconv"
)

type elve struct {
	calories int
}

func newElve() *elve {
	return &elve{calories: 0}
}

func (e *elve) takePackage(call string) (packagesComplete bool, err error) {
	if call == "" {
		return true, nil
	} else {
		newPackage, err := strconv.Atoi(call)
		if err != nil {
			return false, err
		}
		e.calories += newPackage
		return false, nil
	}
}

func (e elve) lineUp(other *elve) (bigger elve, smaller *elve) {
	if e.calories > other.calories {
		return e, other
	} else {
		return *other, &e
	}
}

type topElves struct {
	highScores [3]elve
}

func (t *topElves) lineUp(current *elve) {
	for i, next := range t.highScores {
		t.highScores[i], current = next.lineUp(current)
	}
}

func (t *topElves) sumUp() (all int) {

	for _, score := range t.highScores {
		all += score.calories
	}
	return all
}

func main() {

	s, err := internal.NewLineScanner("./1/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}

	topElves := topElves{highScores: [3]elve{*newElve(), *newElve(), *newElve()}}
	currentElve := newElve()
	for s.Scan() {
		packageComplete, err := currentElve.takePackage(s.Text())
		if err != nil {
			panic(err)
		}
		if packageComplete {
			topElves.lineUp(currentElve)
			currentElve = newElve()
		}
	}

	log.Printf("First:\t%v", topElves.highScores[0].calories)
	log.Printf("Second:\t%v", topElves.highScores[1].calories)
	log.Printf("Third:\t%v", topElves.highScores[2].calories)
	log.Printf("Sum:\t%v", topElves.sumUp())

}
