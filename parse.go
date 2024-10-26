package eval

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jacalz/eval/data"
)

var errMismatchedParenthesis = errors.New("mismatched parenthesis")

func priority(token string) (int, bool) {
	switch token {
	case "^":
		return 2, true
	case "*", "/", "%":
		return 1, true
	case "+", "-":
		return 0, true
	}

	return 0, false
}

func rightAssociated(token string) bool {
	switch token {
	case "^":
		return true
	case "*", "/", "%", "+", "-":
		return false
	}

	return false
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
				continue // Don't fail if stack is empty.
			}

			if _, ok := priority(top); !ok {
				operators.Pop()
				output.Enqueue(top)
			}

		} else if target, ok := priority(t); ok {
			for len(operators.Items) > 0 {
				top, err := operators.Peek()
				if err != nil {
					return nil, err
				}

				if top == "(" {
					break
				}

				prio, _ := priority(top)
				if (prio > target && rightAssociated(t)) ||
					(target <= prio && !rightAssociated(t)) {
					operators.Pop()
					output.Enqueue(top)
				}

				break
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
			return nil, errMismatchedParenthesis
		}

		output.Enqueue(oper)
	}

	return output.Items, nil
}
