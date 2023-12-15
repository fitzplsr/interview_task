package reader

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	reader := strings.NewReader("https://golang.org\n/etc/passwd\nhttps://golang.org\nhttps://golang.org")
	linesArr := [4]string{
		"https://golang.org",
		"/etc/passwd",
		"https://golang.org",
		"https://golang.org",
	}
	lines := Read(reader)

	var index int
	for line := range lines {
		assert.Equal(t, linesArr[index], line, fmt.Sprintf("got %s, want %s", line, linesArr[index]))
		index++
	}
	assert.Equal(t, len(linesArr), index, fmt.Sprintf("got %d, want %d", index, len(linesArr)))
}
