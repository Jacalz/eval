package eval

import (
	"fmt"
	"testing"
)

func TestInfixToRPN(t *testing.T) {
	input := "3 + 4 * 2 / ( 1 − 5 ) ^ 2 ^ 3"
	//expected := []string{"3", "4", "2", "*", "1", "5", "−", "2", "3", "^", "^", "/", "+"}

	actual, err := infixToRPN(input)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(actual)
}
