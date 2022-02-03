package play

import (
	"fmt"
	"os"
	"strings"

	"github.com/guesslin/wordle/internal/check"
	"github.com/guesslin/wordle/internal/dictionary"
	"github.com/guesslin/wordle/internal/pick"
)

const (
	head = "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z"
)

type Round struct {
	// Pre-Configuration
	maxRounds int

	// Pre-Generated question
	question *check.Wordle

	// Records while playing
	history  []string
	alphabet []check.Status
}

func NewRound(max int) *Round {
	result := &Round{
		maxRounds: max,
		question:  check.NewWordle(pick.Pick(dictionary.Pool)),
		history:   make([]string, 0),
		alphabet:  make([]check.Status, 26),
	}

	return result
}

func (r *Round) Play() {
	for len(r.history) < r.maxRounds {
		// ask for input
		answer := ensureAsk(len(r.history) + 1)

		if r.check(answer) {
			return
		}

		r.history = append(r.history, answer)
	}
	fmt.Printf("The Question is %s\n", r.question)
	fmt.Printf("Your guess history: %+v\n", r.history)
	r.printStat()
}

func (r *Round) check(answer string) bool {
	status := r.question.Check(answer)
	if check.Passed(status) {
		fmt.Println("Found")
		return true
	}

	// Update owned alphabet with status
	for i, c := range []byte(answer) {
		r.alphabet[num(c)] = status[i]
	}
	r.printStat()

	return false
}

func (r *Round) printStat() {
	// Print out alphabet status
	fmt.Println(head)
	fmt.Println(buildStat(r.alphabet))
}

func buildStat(input []check.Status) string {
	status := make([]string, 0, len(input))
	for _, s := range input {
		status = append(status, s.String())
	}
	return strings.Join(status, " ")
}

func ensureAsk(i int) string {
	// check if answer is valid
	for {
		var input string
		fmt.Printf("Your %s Guess: ", cardinal(i))
		fmt.Fscanln(os.Stdin, &input)
		answer := strings.ToLower(input)
		if valid(answer) {
			return answer
		}
		fmt.Printf("%s is not a valid word!\n", answer)
	}
}

func valid(input string) bool {
	return dictionary.Have(input)
}

func cardinal(i int) string {
	if i < 0 {
		return "negative"
	}
	switch i {
	case 0:
		return "zero"
	case 1:
		return "1st"
	case 2:
		return "2nd"
	case 3:
		return "3rd"
	}
	return fmt.Sprintf("%dth", i)
}

func num(b byte) int {
	return int(b - 'a')
}
