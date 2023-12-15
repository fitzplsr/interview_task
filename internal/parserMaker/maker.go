package parserMaker

import (
	"finder/internal/parser"
	"finder/internal/parser/fileParser"
	"finder/internal/parser/urlParser"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

func Make(in chan string, target string, timeout time.Duration) (chan parser.Parser, chan error) {
	parsers := make(chan parser.Parser)
	errorCh := make(chan error)

	go func() {
		defer func() {
			close(parsers)
			close(errorCh)
		}()

		for s := range in {
			var p parser.Parser
			source := strings.Trim(s, " ")
			_, err := os.Stat(source)
			if err == nil {
				p = fileParser.NewFileParser(source, target)
				parsers <- p
				continue
			}

			parsedUrl, err := url.ParseRequestURI(source)
			if err == nil && parsedUrl.Scheme != "" {
				p = urlParser.NewUrlParser(source, target, timeout)
				parsers <- p
				continue
			}

			errorCh <- fmt.Errorf("invalid argument: %s", source)
		}
	}()

	return parsers, errorCh
}
