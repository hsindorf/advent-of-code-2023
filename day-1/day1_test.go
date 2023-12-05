package main

import (
	"testing"
)

func Test_getCalibrationTotal(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Example case",
			input: `1abc2
			pqr3stu8vwx
			a1b2c3d4e5f
			treb7uchet`,
			want: 142,
		},
		{
			name: "Part 2",
			input: `two1nine
			eightwothree
			abcone2threexyz
			xtwone3four
			4nineeightseven2
			zoneight234
			7pqrstsixteen`,
			want: 281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCalibrationTotal(tt.input); got != tt.want {
				t.Errorf("getCalibrationTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCalibrationValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example row 1",
			input: "1abc2",
			want:  12,
		},
		{
			name:  "example row 2",
			input: "pqr3stu8vwx",
			want:  38,
		},
		{
			name:  "example row 3",
			input: "a1b2c3d4e5f",
			want:  15,
		},
		{
			name:  "example row 4",
			input: "treb7uchet",
			want:  77,
		},
		{
			name:  "part 2 example",
			input: "zoneight234",
			want:  14,
		},
		{
			name:  "part 2 example 2",
			input: "eightwothree",
			want:  83,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCalibrationValueVersion2(tt.input); got != tt.want {
				t.Errorf("getCalibrationValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
