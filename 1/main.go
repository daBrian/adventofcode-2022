package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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
	richElve := 0
	elvePackage := 0
	for s.Scan() {
		caloriesCall := s.Text()
		switch caloriesCall {
		case "":
			if elvePackage > richElve {
				richElve = elvePackage
			}
			elvePackage = 0
		default:
			calories, err := strconv.Atoi(caloriesCall)
			if err != nil {
				panic(err)
			}
			elvePackage += calories
		}
	}
	log.Printf("%v", richElve)

}
