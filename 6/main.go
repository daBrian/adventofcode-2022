package main

import (
	"errors"
	. "github.com/daBrian/adventofcode-2022/internal"
	"log"
	"strings"
)

func main() {
	s, err := LineScannerFromFile("./6/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	if !s.Scan() {
		log.Panic(errors.New("File seems to be empty"))
	}

	pos, err := findPackagePosition(s.Text())
	if err != nil {
		panic(err)
	}

	log.Printf("6 - first marker after character %v", pos)

}

func findPackagePosition(datastream string) (int, error) {
	for i := range datastream {
		if isBeforeMarker(datastream[i : i+4]) {
			return i + 4, nil
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
