package main

import (
	"finder/internal/parserMaker"
	"finder/internal/reader"
	"finder/internal/workerRunner"
	"fmt"
	"os"
	"sync"
	"time"
)

const target = "Go"
const goroutinesCount = 5
const timeout = time.Second * 5

func main() {
	lines := reader.Read(os.Stdin)
	parsers, errCh := parserMaker.Make(lines, target, timeout)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range errCh {
			fmt.Println(err.Error())
		}
	}()
	go wg.Wait()

	answers := workerRunner.Run(parsers, goroutinesCount)

	var total int64 = 0
	for answer := range answers {
		total += answer.Count
		if answer.Err != nil {
			fmt.Printf("error happened with source %s: %+v\n", answer.SourceName, answer.Err)
		} else {
			fmt.Printf("Count for %s: %d\n", answer.SourceName, answer.Count)
		}
	}
	fmt.Printf("Total: %d\n", total)
}
