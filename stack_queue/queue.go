package main

import "fmt"

type Queue struct {
	slice []int
}

// Enqueue adds the integer provided to the back of the queue
func (q *Queue) Enqueue(i int) {
	//Add i to the queue
	q.slice = append(q.slice, i)
}

// Dequeue returns the first item in the Queue and removes that item form the Queue or panics if there isn't one
func (q *Queue) Dequeue() int {
	//Return the first item in the queue
	var ret int = q.slice[0]
	//Removes the first item from the queue
	q.slice = q.slice[1:len(q.slice)]
	return ret
}

func (q *Queue) String() string {
	return fmt.Sprint(q.slice)
}

func main() {
	var q *Queue = new(Queue)
	q.Enqueue(123)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
}
