package main

import "fmt"

// Go doesn't support generics... Queue<T>, but in Case of Queues it would be nice to have generics
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

//this function is just for formatting purposes
func (q *Queue) String() string {
	return fmt.Sprint(q.slice)
}

func main() {
	var q *Queue = new(Queue)
	q.Enqueue(123)
	q.Enqueue(43)
	q.Enqueue(99)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	q.Enqueue(101)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
}
