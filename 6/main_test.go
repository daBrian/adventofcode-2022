package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPackagePosition(t *testing.T) {
	tests := []struct {
		stream string
		want   int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v(...) first marker after character %v", tt.stream[0:5], tt.want), func(t *testing.T) {
			got, err := findPackagePosition(tt.stream)
			assert.NoError(t, err)
			if got != tt.want {
				t.Errorf("findPackagePosition() got = %v, want %v", got, tt.want)
			}
		})
	}
}
