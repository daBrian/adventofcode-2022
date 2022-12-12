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
	b Stack[string]
	i int
}

func (ls *LineScanner) Close() {
	err := ls.f.Close()
	if err != nil {
		panic(err)
	}
}
func (ls *LineScanner) Scan() bool {
	ls.i++
	if !ls.b.IsEmpty() {
		return true
	} else {
		return ls.s.Scan()
	}
}

func (ls *LineScanner) Text() string {
	if !ls.b.IsEmpty() {
		t, _ := ls.b.Pop()
		return t
	}
	return ls.s.Text()
}
func (ls *LineScanner) Push(line string) {
	ls.b.Push(line)
	ls.i--
}

func (ls *LineScanner) PushMore(l []string) {
	ls.b.PushMore(l)
}

func (ls *LineScanner) LineNumber() int {
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

func LineScannerFromString(input string) *LineScanner {
	reader := io.NopCloser(strings.NewReader(input))
	scanner, err := createScanner(reader)
	if err != nil {
		panic(err)
	}
	return scanner
}
