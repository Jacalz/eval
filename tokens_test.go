package eval

import "testing"

func equal(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func TestSplitIntoTokens_Simple(t *testing.T) {
	input := "56 + 27 * 2"
	expected := []string{"56", "+", "27", "*", "2"}

	actual := splitIntoTokens(input)
	if !equal(actual, expected) {
		t.Errorf("Got %#v, expected %#v", actual, expected)
	}

	input = "56+27*2"

	actual = splitIntoTokens(input)
	if !equal(actual, expected) {
		t.Errorf("Got %#v, expected %#v", actual, expected)
	}
}

func TestSplitIntoTokens_Complex(t *testing.T) {
	input := "2 * ( 5 + 5 ) / 2 - 16 + 2 ^ 3"
	expected := []string{"2", "*", "(", "5", "+", "5", ")", "/", "2", "-", "16", "+", "2", "^", "3"}

	actual := splitIntoTokens(input)
	if !equal(actual, expected) {
		t.Errorf("Got %#v,\nexpected %#v", actual, expected)
	}

	input = "2*(5+5)/2-16+2^3"

	actual = splitIntoTokens(input)
	if !equal(actual, expected) {
		t.Errorf("Got %#v,\nexpected %#v", actual, expected)
	}
}

func TestSplitIntoTokens_Invalid(t *testing.T) {
	input := ")(+2-)("
	expected := []string{")", "(", "+", "2", "-", ")", "("}

	actual := splitIntoTokens(input)
	if !equal(actual, expected) {
		t.Errorf("Got %#v, expected %#v", actual, expected)
	}
}

func TestSplitIntoTokens_Negative(t *testing.T) {
	input := "2 + -10"
	expected := []string{"2", "+", "-", "10"}

	actual := splitIntoTokens(input)
	if !equal(actual, expected) {
		t.Errorf("Got %#v, expected %#v", actual, expected)
	}
}
