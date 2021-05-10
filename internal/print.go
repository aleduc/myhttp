package internal

import "fmt"

type PrintForMock struct {
	printFunc func(a ...interface{}) (n int, err error)
}

func NewPrintForMock() *PrintForMock {
	return &PrintForMock{printFunc: fmt.Println}
}
func (p *PrintForMock) Println(a ...interface{}) (n int, err error) {
	return p.printFunc(a)
}