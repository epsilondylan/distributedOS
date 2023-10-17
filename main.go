package main

import (
	"fmt"
)

type Queue struct {
	capacity  int
	queue     chan string
	enqueueCh chan string
	dequeueCh chan string
	sizeCh    chan int
	capCh     chan int
}

func NewQueue(capacity int) *Queue {
	q := &Queue{
		capacity:  capacity,
		queue:     make(chan string, capacity),
		enqueueCh: make(chan string),
		dequeueCh: make(chan string),
		sizeCh:    make(chan int),
		capCh:     make(chan int),
	}

	go q.processQueue()
	return q
}

func (q *Queue) enqueue(item string) int {
	q.enqueueCh <- item
	return <-q.sizeCh
}

func (q *Queue) dequeue() (string, int) {
	q.dequeueCh <- ""
	return <-q.dequeueCh, <-q.sizeCh
}

func (q *Queue) size() int {
	q.sizeCh <- 0
	return <-q.sizeCh
}

func (q *Queue) capacity() int {
	q.capCh <- 0
	return <-q.capCh
}

func (q *Queue) processQueue() {
	var queue []string

	for {
		select {
		case item := <-q.enqueueCh:
			if len(queue) < q.capacity {
				queue = append(queue, item)
				q.sizeCh <- len(queue)
			} else {
				// Block if the queue is full
				q.sizeCh <- -1
			}
		case q.dequeueCh <- "":
			if len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]
				q.sizeCh <- len(queue)
				q.dequeueCh <- item
			} else {
				// Block if the queue is empty
				q.dequeueCh <- ""
			}
		case q.sizeCh <- 0:
			q.sizeCh <- len(queue)
		case q.capCh <- 0:
			q.capCh <- q.capacity
		}
	}
}

func main() {
	q := NewQueue(3) // Initialize a queue with a capacity of 3

	q.enqueue("Item 1")
	q.enqueue("Item 2")
	q.enqueue("Item 3")

	fmt.Println("Queue size:", q.size())     // Output: 3
	fmt.Println("Queue capacity:", q.capacity()) // Output: 3

	item, _ := q.dequeue()
	fmt.Println("Dequeued:", item) // Output: Dequeued: Item 1
	fmt.Println("Queue size:", q.size()) // Output: 2

	q.enqueue("Item 4")
	fmt.Println("Queue size:", q.size()) // Output: 3
}
