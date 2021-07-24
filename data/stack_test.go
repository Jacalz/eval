package data

import "testing"

var (
	exampleStack         = []string{"0", "1", "2"}
	exampleStackReversed = []string{"2", "1", "0"}
)

func TestStack_Push(t *testing.T) {
	s := NewStack(3)
	s.Push("0")
	s.Push("1")
	s.Push("2")

	for i, actual := range s.Items {
		expected := exampleStack[i]
		if actual != expected {
			t.Errorf("Got %s, expected %s", actual, expected)
		}
	}

	if len(s.Items) != 3 {
		t.Error("Unexpected items in the stack")
	}
}

func TestStack_Pop(t *testing.T) {
	s := Stack{Items: exampleStack}

	for i := range s.Items {
		actual, err := s.Pop()
		if err != nil {
			t.Error(err)
		}

		expected := exampleStackReversed[i]
		if actual != expected {
			t.Errorf("Got %s, expected %s", actual, expected)
		}
	}

	got, err := s.Pop()
	if got != "" || err != errEmptyStack {
		t.Error("The pop should have failed")
	}

	if len(s.Items) != 0 {
		t.Error("Unexpected items in the queue")
	}
}

func TestStack_Peek(t *testing.T) {
	s := Stack{Items: exampleStack}

	actual, err := s.Peek()
	if err != nil {
		t.Error(err)
	}

	expected := exampleStackReversed[0]
	if actual != expected {
		t.Errorf("Got %s, expected %s", actual, expected)
	}

	if len(s.Items) != 3 {
		t.Error("Unexpected items in the queue")
	}

	s.Items = []string{}
	got, err := s.Peek()
	if got != "" || err != errEmptyStack {
		t.Error("The pop should have failed")
	}
}

func TestStack_PushAndPop(t *testing.T) {
	s := Stack{}

	s.Push("0")
	s.Push("1")
	s.Push("2")

	s.Push("3")
	_, err := s.Pop()
	if err != nil {
		t.Error(err)
	}

	s.Push("4")
	_, err = s.Pop()
	if err != nil {
		t.Error(err)
	}

	for i, actual := range s.Items {
		expected := exampleStack[i]
		if actual != expected {
			t.Errorf("Got %s, expected %s", actual, expected)
		}
	}
}
