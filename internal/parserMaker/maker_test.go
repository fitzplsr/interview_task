package parserMaker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMake(t *testing.T) {
	in := make(chan string)
	target := "Go"
	timeout := time.Second * 5
	lines := [2]string{"https://golang.org:", "https://golang.org:"}
	go func() {
		defer close(in)
		for _, line := range lines {
			in <- line
		}
	}()
	parsers, errCh := Make(in, target, timeout)

	var parsersCount int64
	done := make(chan struct{})
	go func() {
		for _ = range parsers {
			parsersCount++
		}
		done <- struct{}{}
	}()

	var errCount int64
	for _ = range errCh {
		errCount++
	}
	assert.Equal(t, int64(0), errCount, fmt.Sprintf("got %d, want %d", errCount, 0))

	<-done
	assert.Equal(t, int64(2), parsersCount, fmt.Sprintf("got %d, want %d", parsersCount, 2))

}
