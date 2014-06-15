package main

import (
	"fmt"
	"time"
)

type Worker struct {
	ID        int
	Work      chan WorkRequest
	WorkQueue chan chan WorkRequest
	QuitChan  chan bool
}

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	worker := Worker{
		ID:        id,
		Work:      make(chan WorkRequest),
		WorkQueue: workerQueue,
		QuitChan:  make(chan bool),
	}

	return worker
}

func (w Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue
			w.WorkQueue <- w.Work

			select {
			case work := <-w.Work:
				// Receive a work request
				fmt.Printf("worker%d: Received work request, delaying for %f seconds\n", w.ID, work.Delay.Seconds())

				time.Sleep(work.Delay)
				fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)

			case <-w.QuitChan:
				fmt.Printf("worker%d sttoping\n", w.ID)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
