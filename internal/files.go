package internal

import (
	"bufio"
	"os"
)

type LineScanner struct {
	s *bufio.Scanner
	f *os.File
}

func (ls LineScanner) Close() error {
	return ls.f.Close()
}
func (ls LineScanner) Scan() bool {
	return ls.s.Scan()
}

func (ls LineScanner) Text() string {
	return ls.s.Text()
}

func NewLineScanner(fileLocation string) (*LineScanner, error) {
	f, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	return &LineScanner{s: s, f: f}, nil
}
