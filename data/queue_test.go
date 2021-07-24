package data

import "testing"

var exampleQueue = []string{"0", "1", "2"}

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue(3)
	q.Enqueue("0")
	q.Enqueue("1")
	q.Enqueue("2")

	for i, actual := range q.Items {
		expected := exampleQueue[i]
		if actual != expected {
			t.Errorf("Got %s, expected %s", actual, expected)
		}
	}

	if len(q.Items) != 3 {
		t.Error("Unexpected items in the queue")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := Queue{Items: exampleQueue}

	for i := range q.Items {
		actual, err := q.Dequeue()
		if err != nil {
			t.Error(err)
		}

		expected := exampleQueue[i]
		if actual != expected {
			t.Errorf("Got %s, expected %s", actual, expected)
		}
	}

	got, err := q.Dequeue()
	if got != "" || err != errEmptyQueue {
		t.Error("The dequeue should have failed")
	}

	if len(q.Items) != 0 {
		t.Error("Unexpected items in the queue")
	}
}

func TestQueue_EnqueueAndDequeue(t *testing.T) {
	q := Queue{}

	q.Enqueue("-2")
	_, err := q.Dequeue()
	if err != nil {
		t.Error(err)
	}

	q.Enqueue("-1")
	_, err = q.Dequeue()
	if err != nil {
		t.Error(err)
	}

	q.Enqueue("0")
	q.Enqueue("1")
	q.Enqueue("2")

	for i, actual := range q.Items {
		expected := exampleQueue[i]
		if actual != expected {
			t.Errorf("Got %s, expected %s", actual, expected)
		}
	}
}
