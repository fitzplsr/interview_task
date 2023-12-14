package urlParser

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewUrlParser(url, target string, timeout time.Duration) *UrlParser {
	return &UrlParser{
		url:     url,
		target:  target,
		timeout: timeout,
	}
}

type UrlParser struct {
	url     string
	target  string
	timeout time.Duration
}

func (up *UrlParser) GetTargetCount() (int64, error) {
	var count int64
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, up.timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", up.url, nil)
	if err != nil {
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	in := bufio.NewScanner(res.Body)
	for in.Scan() {
		count += int64(strings.Count(in.Text(), up.target))
	}
	return count, nil
}

func (up *UrlParser) GetSourceName() string {
	return up.url
}
