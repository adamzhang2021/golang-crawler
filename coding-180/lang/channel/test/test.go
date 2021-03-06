package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		done <- true
	}
}

func createWorker(id int) worker { //箭头代表发数据
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {

	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	chanDemo()
}
