package main

import (
	"errors"
	. "github.com/daBrian/adventofcode-2022/internal"
	"log"
	"strings"
)

func main() {
	r6(4)
	r6(14)

}

func r6(nOfCharacters int) {
	s, err := LineScannerFromFile("./6/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	if !s.Scan() {
		log.Panic(errors.New("File seems to be empty"))
	}

	pos, err := findPackagePosition(s.Text(), nOfCharacters)
	if err != nil {
		panic(err)
	}
	log.Printf("6a - first %v-marker after character %v", nOfCharacters, pos)
}

func findPackagePosition(datastream string, nOfCharacters int) (int, error) {
	for i := range datastream {
		if isBeforeMarker(datastream[i : i+nOfCharacters]) {
			return i + nOfCharacters, nil
		}
	}
	return 0, errors.New("no marker found")
}

func isBeforeMarker(s string) bool {
	for _, i2 := range s {
		if strings.Count(s, string(i2)) > 1 {
			return false
		}
	}
	return true
}
