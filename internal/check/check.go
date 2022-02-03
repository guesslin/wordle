package check

import "strings"

type Status int

const (
	NotAppear Status = iota // X: Not appear in the target string
	Appear                  // B: Appear in the target string, but wrong position
	Same                    // A: Appear in the target string, and correct position
)

type Wordle struct {
	question []byte
	lookup   map[byte]bool
}

func NewWordle(raw string) *Wordle {
	question := []byte(strings.ToLower(raw))
	lookup := make(map[byte]bool)
	for _, c := range question {
		lookup[c] = true
	}

	return &Wordle{
		question: question,
		lookup:   lookup,
	}
}

func (w *Wordle) Ensure(answer string) bool {
	return len(answer) == len(w.question)
}

func (w *Wordle) Check(input string) []Status {
	answer := []byte(strings.ToLower(input))

	res := make([]Status, len(w.question))

	for i, c := range answer {
		if !w.has(c) {
			continue
		}
		res[i]++
		if w.at(i, c) {
			res[i]++
		}
	}

	return res
}

func (w *Wordle) has(c byte) bool {
	return w.lookup[c]
}

func (w *Wordle) at(p int, c byte) bool {
	return c == w.question[p]
}
