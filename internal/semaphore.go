package internal

import (
	"fmt"
	"sync"
)

type response struct {
	url string
	body []byte
}

type semaphoreRequester interface {
	MakeRequest(url string, result chan<- response)
}


// Semaphore is responsible for sending requests.
type Semaphore struct {
	hasher hasher
	requester semaphoreRequester
	printer printer
}

func NewSemaphore(h hasher, r semaphoreRequester, p printer) *Semaphore {
	return &Semaphore{hasher: h, requester: r, printer: p}
}

// Process process HTTP responses asynchronously using the buffered channel as a semaphore.
// Simple solution, only for fun.
func (r *Semaphore) Process(maxConcurrent int, urls []string) {
	boundChan := make(chan struct{}, maxConcurrent)
	resultChan := make(chan response)

	go func() {
		defer close(boundChan)
		defer close(resultChan)
		wg := sync.WaitGroup{}
		for _, u := range urls {
			boundChan <- struct{}{}
			wg.Add(1)
			go func(url string) {
				r.requester.MakeRequest(url, resultChan)
				<-boundChan
				wg.Done()
			}(u)
		}
		wg.Wait()
	}()

	for res := range resultChan {
		r.printer.Println(fmt.Sprintf("%v %v", res.url, r.hasher.GetMD5(res.body)))
	}
}


