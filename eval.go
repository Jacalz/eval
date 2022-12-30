package eval

import (
	"math"
	"strconv"
)

func evaluateRPN(tokens []string) (float64, error) {
	stack := make([]float64, 0, 8)
	for _, t := range tokens {
		switch t {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "^":
			stack[len(stack)-2] = math.Pow(stack[len(stack)-2], stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		default:
			f, err := strconv.ParseFloat(t, 64)
			if err != nil {
				return 0, err
			}

			stack = append(stack, f)
		}
	}

	return stack[0], nil
}

// Numerical evaluates a numerical mathematical expression and returns the result.
func Numerical(input string) (float64, error) {
	rpn, err := infixToRPN(input)
	if err != nil {
		return 0, err
	}

	result, err := evaluateRPN(rpn)
	if err != nil {
		return 0, err
	}

	return result, nil
}
