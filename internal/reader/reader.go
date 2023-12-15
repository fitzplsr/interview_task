package reader

import (
	"bufio"
	"io"
)

func Read(reader io.Reader) chan string {
	in := bufio.NewScanner(reader)
	out := make(chan string)

	go func() {
		defer close(out)
		for in.Scan() {
			out <- in.Text()
		}
	}()

	return out
}
