package workerRunner

import (
	"finder/internal/parser"
	"sync"
)

type Answer struct {
	Count      int64
	SourceName string
	Err        error
}

func Run(parsers chan parser.Parser, goroutinesCount int64) chan Answer {
	answers := make(chan Answer)
	concurrentGoroutines := make(chan struct{}, goroutinesCount)

	var i int64 = 0
	for ; i < goroutinesCount; i++ {
		concurrentGoroutines <- struct{}{}
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for p := range parsers {
			wg.Add(1)
			<-concurrentGoroutines

			go func(p parser.Parser) {
				defer wg.Done()

				answer := doWork(p)

				answers <- answer
				concurrentGoroutines <- struct{}{}
			}(p)
		}
	}()

	go func() {
		wg.Wait()
		close(concurrentGoroutines)
		close(answers)
	}()

	return answers
}

func doWork(p parser.Parser) Answer {
	count, err := p.GetTargetCount()
	return Answer{
		Count:      count,
		SourceName: p.GetSourceName(),
		Err:        err,
	}
}
