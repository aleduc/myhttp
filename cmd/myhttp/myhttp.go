package main

import (
	"flag"
	"myhttp/internal"
	"net/http"
	"time"
)

func main() {
	p := internal.NewPool(
		internal.NewURLTask(
			&internal.Hash{},
			// Almost default settings.
			internal.NewHTTPWrap(&http.Client{Timeout: 5 * time.Second}),
			internal.NewPrintForMock()),
		)
	p.Start(parseInputParams())
}

func parseInputParams() (int,[]string) {
	maxConcurrent := flag.Int("parallel",10,"Maximum concurrent requests.")
	flag.Parse()
	return *maxConcurrent, flag.Args()
}