package internal

import (
	"sync"
)

// Pool struct for control our worker pool.
type Pool struct {
	worker Worker
}

// Worker process jobs from  channel.
type Worker interface {
	Process(jobs <-chan string, group *sync.WaitGroup)
}

// NewPool init new Pool.
func NewPool(w Worker) *Pool {
	return &Pool{worker: w}
}

// Process run all workers/consumer, produce jobs, wait results.
func (p *Pool) Start(maxConcurrent int, urls []string) {
	var wg sync.WaitGroup

	actualMaxConcurrent := maxConcurrent
	if len(urls) < maxConcurrent {
		actualMaxConcurrent = len(urls)
	}

	jobs := make(chan string, len(urls))
	for i := 1; i <= actualMaxConcurrent; i++ {
		wg.Add(1)
		go p.worker.Process(jobs, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg.Wait()
}