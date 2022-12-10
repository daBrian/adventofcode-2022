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
var dirPattern = regexp.MustCompile(`^dir [a-z]+$`)
var lsPattern = regexp.MustCompile(`^\$ ls$`)
var cdPattern = regexp.MustCompile(`^\$ cd [a-z]+$`)
var cdUpPattern = regexp.MustCompile(`^\$ cd \.\.$`)

type Element interface {
	GetName() string
	DeepString() string
	Path() string
	AsDirNode() (*DirNode, error)
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

func (n FileNode) AsDirNode() (*DirNode, error) {
	return nil, errors.New("node is file, not dir")
}

func (n FileNode) DeepString() string {
	return n.String()
}

func (n TreeNode) Path() string {
	if n.depth == 0 {
		return "/"
	} else {
		return fmt.Sprintf("%v/%v", n.Parent.Path(), n.Name)
	}
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

func (d DirNode) AsDirNode() (*DirNode, error) {
	return &d, nil
}

func NewDirNode(parent *DirNode, line string) (Element, error) {
	fields := strings.Fields(line)
	return &DirNode{
		TreeNode: TreeNode{
			Parent: parent,
			Name:   fields[1],
			depth:  parent.depth + 1,
		},
		Children: map[string]Element{},
	}, nil
}

func (d DirNode) String() string {
	return fmt.Sprintf("- %v (dir)", d.Name)
}

func (d DirNode) DeepString() string {
	var list = []string{}
	list = append(list, fmt.Sprintf("- %v (dir)", d.Name))
	for _, c := range d.Children {
		list = append(list, c.DeepString())
	}
	return strings.Join(list, "\n")
}

func (d *DirNode) parseNext(s *LineScanner) (err error) {
	for s.Scan() {
		var child Element
		line := s.Text()
		switch {
		case filePattern.MatchString(line):
			child, err = NewFileNode(d, line)
			_, ok := d.Children[child.GetName()]
			if ok {
				return fmt.Errorf("Line %v: element with name %v already exists in %v",
					s.LineNumber(), child.GetName(), d.GetName())
			}
			d.Children[child.GetName()] = child
			fmt.Println(child.Path())
		case dirPattern.MatchString(line):
			child, err = NewDirNode(d, line)
			d.Children[child.GetName()] = child
			fmt.Println(child.Path())
		case lsPattern.MatchString(line):
			continue
		case cdPattern.MatchString(line):
			dir := strings.Fields(line)[2]
			dn, err := d.Cd(dir)
			if err != nil {
				return err
			}
			d = dn
		case cdUpPattern.MatchString(line):
			d = d.Parent
		default:
			return fmt.Errorf("Unknown line %v: %v", s.LineNumber(), line)
		}
		if err != nil {
			return fmt.Errorf("Line %v: Could not create FileNode from '%v'", s.LineNumber(), line)
		}
	}
	return nil
}

func (d DirNode) Cd(dir string) (*DirNode, error) {
	element := d.Children[dir]
	dirNode, err := element.AsDirNode()
	if err != nil {
		return nil, err
	}
	return dirNode, nil
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
		return nil, err
	}
	return &root, nil
}
