package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPackagePosition(t *testing.T) {
	tests := []struct {
		dataStream    string
		nOfCharacters int
		want          int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v(...) first %v-marker after character %v", tt.dataStream[0:5], tt.nOfCharacters, tt.want), func(t *testing.T) {
			got, err := findPackagePosition(tt.dataStream, tt.nOfCharacters)
			assert.NoError(t, err)
			if got != tt.want {
				t.Errorf("findPackagePosition() got = %v, want %v", got, tt.want)
			}
		})
	}
}
