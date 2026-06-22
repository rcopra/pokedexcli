package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "trims surrounding spaces",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "lowercases words",
			input:    "  HELLO  WorLD  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "single word",
			input:    "charmander",
			expected: []string{"charmander"},
		},
		{
			name:     "empty input",
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("got %v, want %v", actual, c.expected)
			}
		})
	}
}
