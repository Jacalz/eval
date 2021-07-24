package eval

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jacalz/eval/data"
)

var errMismatchedParenthesis = errors.New("mismatched parenthesis")

var priorities = map[string]int{
	"^": 2,
	"*": 1,
	"/": 1,
	"+": 0,
	"-": 0,
}

var rightAssociated = map[string]bool{
	"^": true,
	"*": false,
	"/": false,
	"+": false,
	"-": false,
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// infixToRPN converts an infix notation string to reverse polish notation using a shunting yard algorithm.
func infixToRPN(input string) ([]string, error) {
	output := data.NewQueue(16)
	operators := data.NewStack(16)

	for _, t := range strings.Fields(input) {
		if t == "(" {
			operators.Push(t)
		} else if t == ")" {
			foundLeftMatch := false

			// Pop items from the stack to the queue until the matching parenthesis is found.
			for len(operators.Items) > 0 {
				oper, err := operators.Pop()
				if err != nil {
					return nil, err
				}

				if oper == "(" {
					foundLeftMatch = true
					break
				}

				output.Enqueue(oper)
			}

			if !foundLeftMatch {
				return nil, errMismatchedParenthesis
			}

			// If the top in the stack is not an operator, pop over the corresponding function.
			top, err := operators.Peek()
			if err != nil {
				return nil, err
			}

			if _, ok := priorities[top]; !ok {
				operators.Pop()
				output.Enqueue(top)
			}

		} else if priority, ok := priorities[t]; ok {
			for len(operators.Items) > 0 {
				top, err := operators.Peek()
				if err != nil {
					return nil, err
				}

				if top == "(" {
					break
				}

				if (priorities[top] > priority && rightAssociated[t]) ||
					(priorities[top] <= priority && !rightAssociated[t]) {
					operators.Pop()
					output.Enqueue(top)
				}
			}

			operators.Push(t)
		} else if isNumeric(t) {
			output.Enqueue(t)
		} else { // Token is a function.
			operators.Push(t)
		}
	}

	// Pop remaining items from the stack to the queue.
	for len(operators.Items) > 0 {
		oper, err := operators.Pop()
		if err != nil {
			return nil, err
		}

		if oper == "(" {
			return nil, errors.New("mismatched parenthesis")
		}

		output.Enqueue(oper)
	}

	return output.Items, nil
}

// Eval evaluates the mathematical expression and returns the result.
func Eval(input string) float64 {
	return 0
}
