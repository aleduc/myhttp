package internal

import (
	"fmt"
	"sync"
)

type hasher interface {
	GetMD5(input []byte) string
}

type requester interface {
	MakeRequest(url string) []byte
}

type printer interface {
	Println(a ...interface{}) (n int, err error)
}


// URLTask is responsible for processing url.
type URLTask struct {
	hasher hasher
	requester requester
	printer printer
}


func NewURLTask(h hasher, r requester, p printer) *URLTask {
	return &URLTask{hasher: h, requester: r, printer: p}
}


func (u *URLTask) Process(jobs <-chan string, group *sync.WaitGroup) {
	defer group.Done()
	for j := range jobs {
		res := u.requester.MakeRequest(j)
		u.printer.Println(fmt.Sprintf("%v%v %v", defaultSchema, j, u.hasher.GetMD5(res)))
	}
}

