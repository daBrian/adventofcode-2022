package main

import (
	"errors"
	"fmt"
	. "github.com/daBrian/adventofcode-2022/internal"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var filePattern = regexp.MustCompile(`^[0-9]+ .+$`)
var lsPattern = regexp.MustCompile(`^\$ ls$`)

type Element interface {
	GetName() string
	DeepString() string
}
type TreeNode struct {
	Parent *DirNode
	Name   string
	depth  int
}

func (n TreeNode) GetName() string {
	return n.Name
}

type FileNode struct {
	TreeNode
	Size int
}

func (n FileNode) DeepString() string {
	return n.String()
}

func (n FileNode) String() string {
	return fmt.Sprintf("%v - %v (file, size=%v)", strings.Repeat(" ", n.depth), n.Name, n.Size)
}

func NewFileNode(parent *DirNode, line string) (*FileNode, error) {
	fields := strings.Fields(line)
	size, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, errors.New("invalid number")
	}
	return &FileNode{
		TreeNode: TreeNode{
			Parent: parent,
			Name:   fields[1],
			depth:  parent.depth + 1,
		},
		Size: size,
	}, nil
}

type DirNode struct {
	TreeNode
	Children map[string]Element
}

func (d DirNode) String() string {
	var list = []string{}
	list = append(list, fmt.Sprintf("- %v (dir)", d.Name))
	for _, c := range d.Children {
		list = append(list, c.DeepString())
	}
	return strings.Join(list, "\n")
}
func (d DirNode) DeepString() string {
	return d.String()
}

func (d DirNode) parseNext(s *LineScanner) (err error) {
	if s.Scan() {
		line := s.Text()
		switch {
		case lsPattern.MatchString(line):
			err = d.ls(s)
		default:
			return fmt.Errorf("Unknown line %v: %v", s.LineNumber(), line)
		}
		if err != nil {
			return fmt.Errorf("Line %v: Could not create FileNode from '%v'", s.LineNumber(), line)
		}
	}
	return nil
}

func (d DirNode) ls(s *LineScanner) (err error) {
	var child Element
	if s.Scan() {
		line := s.Text()
		switch {
		case filePattern.MatchString(line):
			child, err = NewFileNode(&d, line)
			d.Children[child.GetName()] = child
		default:
			return fmt.Errorf("Unknown line %v: %v", s.LineNumber(), line)
		}
		if err != nil {
			return fmt.Errorf("Line %v: Could not create FileNode from '%v'", s.LineNumber(), line)
		}
	}

	return nil
}
func main() {
	r7a()
}

func r7a() {
	s, err := LineScannerFromFile("./7tree/input.txt")
	defer s.Close()
	if err != nil {
		log.Panic(err)
	}
	_, err = parseCommands(s)
	if err != nil {
		panic(err)
	}
}

func parseCommands(s *LineScanner) (*DirNode, error) {
	if !s.Scan() {
		return nil, errors.New("Empty file.")
	}
	if s.Text() != "$ cd /" {
		return nil, errors.New("Unexpected start of file.")
	}
	root := DirNode{
		TreeNode: TreeNode{
			Parent: &DirNode{},
			Name:   "/",
			depth:  0,
		},
		Children: map[string]Element{},
	}
	err := root.parseNext(s)
	if err != nil {
		panic(err)
	}
	return &root, nil
}
