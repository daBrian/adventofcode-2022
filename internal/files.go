package internal

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type LineScanner struct {
	s *bufio.Scanner
	f io.ReadCloser
	i int
}

func (ls LineScanner) Close() {
	err := ls.f.Close()
	if err != nil {
		panic(err)
	}
}
func (ls LineScanner) Scan() bool {
	ls.i++
	return ls.s.Scan()
}

func (ls LineScanner) Text() string {
	return ls.s.Text()
}

func (ls LineScanner) LineNumber() int {
	return ls.i
}

func LineScannerFromFile(fileLocation string) (*LineScanner, error) {
	f, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	return createScanner(f)
}

func createScanner(f io.ReadCloser) (*LineScanner, error) {
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	return &LineScanner{s: s, f: f}, nil
}

func LineScannerFromString(input string) (*LineScanner, error) {
	reader := io.NopCloser(strings.NewReader(input))
	return createScanner(reader)
}
