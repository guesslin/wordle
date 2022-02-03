package dictionary_test

import (
	"testing"

	"github.com/guesslin/mywordle/internal/dictionary"

	"github.com/stretchr/testify/assert"
)

func TestDictionary_Have(t *testing.T) {
	testCases := []struct {
		message  string
		input    string
		expected bool
	}{
		{
			message:  "Empty input should not be found",
			input:    "",
			expected: false,
		},
		{
			message:  "Normal word",
			input:    "audio",
			expected: true,
		},
		{
			message:  "Abnormal word",
			input:    "phpie",
			expected: false,
		},
		{
			message:  "Word has more than 5 char",
			input:    "normal",
			expected: false,
		},
		{
			message:  "Word has less than 5 char",
			input:    "less",
			expected: false,
		},
	}
	for _, c := range testCases {
		res := dictionary.Have(c.input)
		assert.Equal(t, c.expected, res, c.message)
	}

}
