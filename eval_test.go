package eval

import "testing"

func TestEval(t *testing.T) {
	input := "( 6 - 2 * ( 6 / 3 ) ) ^ 3"

	expected := 8.0
	actual, err := Eval(input)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Got %f, expected %f", actual, expected)
	}
}
