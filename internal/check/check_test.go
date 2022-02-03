package check_test

import (
	"testing"

	"github.com/guesslin/wordle/internal/check"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWordle_Check(t *testing.T) {
	testCases := []struct {
		title    string
		question string
		answer   string
		expected []check.Status
	}{
		{
			title:    "All Matched",
			question: "audio",
			answer:   "audio",
			expected: []check.Status{check.Same, check.Same, check.Same, check.Same, check.Same},
		},
		{
			title:    "All Matched, Answer in UpperCases",
			question: "audio",
			answer:   "AUDIO",
			expected: []check.Status{check.Same, check.Same, check.Same, check.Same, check.Same},
		},
		{
			title:    "All Matched, Question in UpperCases",
			question: "AUDIO",
			answer:   "audio",
			expected: []check.Status{check.Same, check.Same, check.Same, check.Same, check.Same},
		},
		{
			title:    "All not Matched",
			question: "audio",
			answer:   "perky",
			expected: []check.Status{check.NotAppear, check.NotAppear, check.NotAppear, check.NotAppear, check.NotAppear},
		},
		{
			title:    "AAAXX",
			question: "phone",
			answer:   "phoca",
			expected: []check.Status{check.Same, check.Same, check.Same, check.NotAppear, check.NotAppear},
		},
		{
			title:    "BBBAB",
			question: "world",
			answer:   "droll",
			expected: []check.Status{check.Appear, check.Appear, check.Appear, check.Same, check.Appear},
		},
	}
	for _, c := range testCases {
		target := check.NewWordle(c.question)
		require.True(t, target.Ensure(c.answer))
		res := target.Check(c.answer)
		assert.Equal(t, c.expected, res, c.title)
	}

}
