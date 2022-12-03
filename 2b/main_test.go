package main

import (
	"github.com/daBrian/adventofcode-2022/internal"
	"log"
	"testing"
)

func Test_calculateTotalScore(t *testing.T) {
	ls, _ := internal.LineScannerFromString("A Y\nB X\nC Z")
	got, _ := calculateTotalScore(ls)
	log.Printf("Sum is %v", got)
}
