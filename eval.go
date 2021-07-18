package eval

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jacalz/eval/data"
)

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
	outputQueue := data.NewQueue(16)
	operatorStack := data.NewStack(16)

	for _, t := range strings.Fields(input) {
		fmt.Println("Queue:", outputQueue, "Stack:", operatorStack)

		if isNumeric(t) {
			outputQueue.Enqueue(t)
		} else if t == "(" {
			operatorStack.Push(t)
		} else if t == ")" {
			foundLeftMatch := false

			// Pop items from the stack to the queue until the matching parenthesis is found.
			for len(operatorStack.Items) > 0 {
				oper := operatorStack.Pop()
				if oper == "(" {
					foundLeftMatch = true
					break
				}

				outputQueue.Enqueue(oper)
			}

			if !foundLeftMatch {
				return nil, errors.New("mismatched parenthesis")
			}

			// If the top in the queue is not an operator, pop over the corresponding function.
			if _, ok := priorities[operatorStack.Peek()]; !ok {
				outputQueue.Enqueue(operatorStack.Pop())
			}

		} else if priority, ok := priorities[t]; ok {
			for len(operatorStack.Items) > 0 {
				top := operatorStack.Peek()
				if top == "(" {
					break
				}

				if (priorities[top] > priority && rightAssociated[t]) ||
					(priorities[top] <= priority && !rightAssociated[t]) {
					outputQueue.Enqueue(operatorStack.Pop())
				}
			}

			operatorStack.Push(t)
		} else { // Token is a function.
			operatorStack.Push(t)
		}
	}

	// Pop remaining items from the stack to the queue.
	for len(operatorStack.Items) > 0 {
		oper := operatorStack.Pop()
		if oper == "(" {
			return nil, errors.New("mismatched parenthesis")
		}

		outputQueue.Enqueue(oper)
	}

	return outputQueue.Items, nil
}

// Eval evaluates the mathematical expression and returns the result.
func Eval(input string) float64 {
	return 0
}
