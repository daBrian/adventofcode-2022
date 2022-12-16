package main

import (
	. "github.com/daBrian/adventofcode-2022/internal"
	"log"
)

func main() {
	r9a()
	r9b()
}

func r9a() {
	s, err := LineScannerFromFile("./09ropebridge/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}

}

func r9b() {
	s, err := LineScannerFromFile("./09ropebridge/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}

}
