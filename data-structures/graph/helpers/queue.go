package helpers

type Queue[T any] struct {
	elements []T
}


func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value) 
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, true
}


func (q *Queue[T]) Peek() (T, bool) {
	if len(q.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return q.elements[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}


