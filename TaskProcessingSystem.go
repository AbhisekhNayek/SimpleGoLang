package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a simple task with an ID.
type Task struct {
	ID int
}

// Worker represents a worker that processes tasks.
type Worker struct {
	ID        int
	taskQueue chan Task
	done      chan bool
}

// NewWorker creates a new worker.
func NewWorker(id int, taskQueue chan Task, wg *sync.WaitGroup) *Worker {
	worker := &Worker{
		ID:        id,
		taskQueue: taskQueue,
		done:      make(chan bool),
	}

	go worker.start(wg)
	return worker
}

// start starts the worker's processing loop.
func (w *Worker) start(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task := <-w.taskQueue:
			fmt.Printf("Worker %d processing task %d\n", w.ID, task.ID)
			// Simulate some work
			time.Sleep(time.Second)
		case <-w.done:
			return
		}
	}
}

// Stop stops the worker.
func (w *Worker) Stop() {
	close(w.done)
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	// Create a task queue and a wait group
	taskQueue := make(chan Task, numTasks)
	var wg sync.WaitGroup

	// Create workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		worker := NewWorker(i, taskQueue, &wg)
		defer worker.Stop() // Ensure workers are stopped when the main function exits
	}

	// Enqueue tasks
	for i := 1; i <= numTasks; i++ {
		taskQueue <- Task{ID: i}
	}

	// Close the task queue to signal workers to stop
	close(taskQueue)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All tasks processed.")
}
