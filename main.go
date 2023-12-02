package main

import (
	"go-worker-pool/workerpool"
	"log"
	"time"
)

const (
	maxWorker = 5
	maxJob    = 50
)

type result struct {
	id  int
	val int
}

func main() {
	wpool := workerpool.NewWorkerPool(maxWorker)
	wpool.Init()

	resultChan := make(chan result, maxJob)

	for i := 0; i < maxJob; i++ {
		id := i + 1
		wpool.AddJob(func() {
			log.Printf("Running job %d", id)
			time.Sleep(2 * time.Second)
			resultChan <- result{id, id * 10}
		})
	}

	for i := 0; i < maxJob; i++ {
		r := <-resultChan
		log.Printf("Job %d finished with response %d", r.id, r.val)
	}
}
