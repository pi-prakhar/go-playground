package main

//building queues in golang using channels

type Queue struct {
	Data chan int
}

func (Queue *Queue) InitQueue(len int) {
	Queue.Data = make(chan int, len)
}

func (Queue *Queue) Enqueue(val int) bool {
	if len(Queue.Data) == cap(Queue.Data) {
		return false
	} else {
		Queue.Data <- val
		return true
	}
}

func (Queue *Queue) Dequeue() (int, bool) {
	if len(Queue.Data) > 0 {
		return <-Queue.Data, true
	}
	return -1, false
}
