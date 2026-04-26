package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " tesT    ",
			expected: []string{"test"}},
		{
			input:    "this Is a Test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "  hello  world    ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLen := len(actual)
		expectedLen := len(c.expected)
		if actualLen != expectedLen {
			t.Errorf("Incorrect slice length; expected: %d, got %d", expectedLen, actualLen)
		}
		for i, word := range actual {
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Incorrect word; expected: %s, got %s", expectedWord, word)
			}
		}
	}
}
