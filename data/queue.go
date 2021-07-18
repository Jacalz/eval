package data

// Queue is a simple FIFO queue implementation.
// NOTE: The queue is not concurrency safe.
type Queue struct {
	Items []string
}

// Enqueue adds the element to the back of the queue.
func (q *Queue) Enqueue(element string) {
	q.Items = append(q.Items, element)
}

// Dequeue removes the first item from the queue.
func (q *Queue) Dequeue() string {
	first := q.Items[0]
	q.Items = q.Items[1:]
	return first
}

// NewQueue returns a new queue with a set default size.
// The queue can continue to grow after the size is reached.
func NewQueue(size int) *Queue {
	return &Queue{
		Items: make([]string, 0, size),
	}
}
