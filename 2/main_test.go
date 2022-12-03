package main

import (
	"reflect"
	"testing"
)

func TestNewBet(t *testing.T) {
	tests := []struct {
		name           string
		args           string
		wantCurrentBet *round
		wantErr        bool
	}{
		{name: "'A ?' has A",
			args:           "A X",
			wantCurrentBet: &round{other: 'A', my: 'X'},
			wantErr:        false,
		}, {name: "'B ?' has B",
			args:           "B X",
			wantCurrentBet: &round{other: 'B', my: 'X'},
			wantErr:        false,
		}, {name: "'C ?' has C",
			args:           "C X",
			wantCurrentBet: &round{other: 'C', my: 'X'},
			wantErr:        false,
		}, {name: "Only accepts [A-C] as first character",
			args:           "D X",
			wantCurrentBet: nil,
			wantErr:        true,
		}, {name: "'? X' has X",
			args:           "A X",
			wantCurrentBet: &round{other: 'A', my: 'X'},
			wantErr:        false,
		}, {name: "'? Y' has Y",
			args:           "A Y",
			wantCurrentBet: &round{other: 'A', my: 'Y'},
			wantErr:        false,
		}, {name: "'? Z' has Z",
			args:           "A Z",
			wantCurrentBet: &round{other: 'A', my: 'Z'},
			wantErr:        false,
		}, {name: "Only accepts [X-Z] as first character",
			args:           "A U",
			wantCurrentBet: nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCurrentBet, err := NewRound(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRound() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCurrentBet, tt.wantCurrentBet) {
				t.Errorf("NewRound() gotCurrentBet = %v, want %v", gotCurrentBet, tt.wantCurrentBet)
			}
		})
	}
}

func Test_round_valueMyChoice(t *testing.T) {
	type fields struct {
		other rune
		my    rune
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Rock (X) is 1",
			fields: fields(round{my: 'X', other: 0}),
			want:   1},
		{name: "Paper (Y) is 2",
			fields: fields(round{my: 'Y', other: 0}),
			want:   2},
		{name: "Scissors (Z) is 3",
			fields: fields(round{my: 'Z', other: 0}),
			want:   3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := round{
				other: tt.fields.other,
				my:    tt.fields.my,
			}
			if got := b.valueMyChoice(); got != tt.want {
				t.Errorf("valueMyChoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round_valueRoundResult(t *testing.T) {
	type fields struct {
		other rune
		my    rune
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Win with Rock (C X)",
			fields: fields(round{other: 'C', my: 'X'}),
			want:   6},
		{name: "Win with Paper (A Y)",
			fields: fields(round{other: 'A', my: 'Y'}),
			want:   6},
		{name: "Win with Scissors (B Z)",
			fields: fields(round{other: 'B', my: 'Z'}),
			want:   6},
		{name: "Draw with Rock (A X)",
			fields: fields(round{other: 'A', my: 'X'}),
			want:   3},
		{name: "Draw with Paper (B Y)",
			fields: fields(round{other: 'B', my: 'Y'}),
			want:   3},
		{name: "Draw with Scissors (C Z)",
			fields: fields(round{other: 'C', my: 'Z'}),
			want:   3},
		{name: "Lose with Rock (B X)",
			fields: fields(round{other: 'B', my: 'X'}),
			want:   0},
		{name: "Lose with Paper (C Y)",
			fields: fields(round{other: 'C', my: 'Y'}),
			want:   0},
		{name: "Lose with Scissors (B Z)",
			fields: fields(round{other: 'A', my: 'Z'}),
			want:   0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := round{
				other: tt.fields.other,
				my:    tt.fields.my,
			}
			if got := b.valueRoundResult(); got != tt.want {
				t.Errorf("valueRoundResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateRoundPoints(t *testing.T) {
	tests := []struct {
		name       string
		args       string
		wantPoints int
	}{
		{"Example 1", "A Y", 8},
		{"Example 2", "B X", 1},
		{"Example 3", "C Z", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := NewRound(tt.args)
			if gotPoints := calculateRoundPoints(*r); gotPoints != tt.wantPoints {
				t.Errorf("calculateRoundPoints() = %v, want %v", gotPoints, tt.wantPoints)
			}
		})
	}
}
