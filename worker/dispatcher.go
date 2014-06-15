package main

import "fmt"

var WorkerQueue chan chan WorkRequest

func StartDispatcher(nworkers int) {
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker ", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work request")
				go func() {
					worker := <-WorkerQueue

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
