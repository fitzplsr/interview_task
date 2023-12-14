package fileParser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type FileParser struct {
	fileName string
	target   string
}

func NewFileParser(fileName, target string) *FileParser {
	return &FileParser{
		fileName: fileName,
		target:   target,
	}
}

func (fp *FileParser) GetTargetCount() (int64, error) {
	var count int64
	var file *os.File

	file, err := os.Open(fp.fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(file)

	if err != nil {
		return 0, err
	}

	in := bufio.NewScanner(file)
	for in.Scan() {
		count += int64(strings.Count(in.Text(), fp.target))
	}
	return count, nil
}

func (fp *FileParser) GetSourceName() string {
	return fp.fileName
}
