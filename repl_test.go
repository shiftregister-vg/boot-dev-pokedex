package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i, word := range actual {
			expected := c.expected[i]
			if expected != word {
				t.Errorf("word did not match expectation. Expected: '%s', Actual: '%s'", expected, word)
				t.Fail()
			}
		}
	}
}
