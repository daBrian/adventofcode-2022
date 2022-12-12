package main

import (
	"fmt"
	"github.com/daBrian/adventofcode-2022/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
var examplePrint = `- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)`
var _ = examplePrint

func Test_parseCommands(t *testing.T) {
	tests := []struct {
		name  string
		input string
		print string
	}{
		{"only root", "$ cd /", "- / (dir)"},
		{"root and file", "$ cd /\n$ ls\n14848514 b.txt", "- / (dir)\n  - b.txt (file, size=14848514)"},
		{"root and file", "$ cd /\n$ ls\n14848514 b.txt", "- / (dir)\n  - b.txt (file, size=14848514)"},
		//{"codeexample", exampleInput, examplePrint},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ls := internal.LineScannerFromString(test.input)
			got, err := parseCommands(ls)
			assert.NoError(t, err)
			assert.Equal(t, test.print, got.DeepString())
		})
	}
}
func TestSizes(t *testing.T) {
	ls := internal.LineScannerFromString(exampleInput)
	root, err := parseCommands(ls)
	assert.NoError(t, err)
	dirs, _ := root.collectDirsWithSizes()
	for _, dir := range dirs {
		fmt.Printf("%v: %v\n", dir.dir.Path(), dir.totalSize)
	}
	smallerDirs := dirsWithAtMost(dirs, 100000)
	fmt.Printf("7a - summarized size of small dirs is %v\n", sumUpSizes(smallerDirs))

}
