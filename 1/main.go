package main

import (
	"bufio"
	"log"
	"os"
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

func (e elve) lineUp(other *elve) (bigger *elve, smaller *elve) {
	if e.calories > other.calories {
		return &e, other
	} else {
		return other, &e
	}
}

func main() {

	f, err := os.Open("./1/input.txt")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panic(err)
		}
	}(f)
	if err != nil {
		log.Panic(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	richElve := newElve()
	currentElve := newElve()
	for s.Scan() {
		packageComplete, err := currentElve.takePackage(s.Text())
		if err != nil {
			panic(err)
		}
		if packageComplete {
			richElve, _ = richElve.lineUp(currentElve)
			currentElve = newElve()
		}
	}
	log.Printf("%v", richElve.calories)

}
