package play

import (
	"fmt"
	"os"

	"github.com/guesslin/mywordle/internal/dictionary"
	"github.com/guesslin/mywordle/internal/pick"
)

type Round struct {
	// Pre-Configuration
	maxRounds int

	// Pre-Generated question
	question string

	// Records while playing
	history []string
}

func NewRound(max int) *Round {
	result := &Round{
		maxRounds: max,
		question:  pick.Pick(dictionary.Pool),
		history:   make([]string, 0),
	}

	fmt.Println(result.question)

	return result
}

func (r *Round) Play() {
	for len(r.history) < r.maxRounds {
		// ask for input
		answer := ensureAsk(len(r.history) + 1)

		if answer == r.question {
			fmt.Println("Found")
			return
		}
		r.history = append(r.history, answer)
	}
	fmt.Printf("The Question is %s\n", r.question)
}

func ensureAsk(i int) string {
	// check if answer is valid
	for {
		var answer string
		fmt.Printf("Your %s Guess: ", cardinal(i))
		fmt.Fscanln(os.Stdin, &answer)
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
