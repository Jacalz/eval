package data

// Stack is a simple LIFO stack implementation.
// NOTE: The stack is not concurrency safe.
type Stack struct {
	Items []string
}

// Push adds a new element to the stack.
func (s *Stack) Push(element string) {
	s.Items = append(s.Items, element)
}

// Pop removes the top most element from the stack.
func (s *Stack) Pop() string {
	top := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return top
}

// Peek returns the top most element from the stack.
func (s *Stack) Peek() string {
	if len(s.Items) < 1 {
		return ""
	}

	return s.Items[len(s.Items)-1]
}

// NewStack returns a new stack with a set inital size.
// The stack can continue to grow after the size is reached.
func NewStack(size int) *Stack {
	return &Stack{
		Items: make([]string, 0, size),
	}
}
