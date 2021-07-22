package helper

import (
	"bufio"
	"github.com/araddon/dateparse"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const (
	duration  = 5
	directory = "./dummylogs"

	nowStr = "2021-07-22 11:00:00"
)

type dummyLogFileContent struct {
	text  string
	aging float64
}

func TestReadLog(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, nowStr)
	var dummyLogFileContents []dummyLogFileContent
	var logFilePaths []string

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), suffixFile) {
			if file.ModTime().Minute() <= int(duration) {
				logFilePaths = append(logFilePaths, filepath.Join(directory, file.Name()))
			}
		}
	}

	for _, logFilePath := range logFilePaths {
		f, err := os.Open(logFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			textArr := strings.Split(scanner.Text(), " ")
			textRune := []rune(textArr[3])

			length := len(textRune)
			contentTimeStr := string(textRune[1:length])
			contentTime, _ := dateparse.ParseAny(contentTimeStr)

			var contentAging float64
			contentAging = math.Round(now.Sub(contentTime).Minutes())
			if contentAging <= duration {
				dummyLogFileContents = append(dummyLogFileContents, dummyLogFileContent{
					text:  scanner.Text(),
					aging: contentAging,
				})
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		f.Close()
	}

	assert.NotEmpty(t, dummyLogFileContents)
	assert.NotEmpty(t, logFilePaths)
}
