package helper

import (
	"bufio"
	"fmt"
	"github.com/araddon/dateparse"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ILogReader interface {
	setDuration()

	setFileLogPaths()

	setLogFileContents()

	PrintLogFileContents()
}

type logFileContent struct {
	text  string
	aging float64
}

type LogReader struct {
	timeFlag        string
	directoryFlag   string
	duration        float64
	logFilePaths    []string
	logFileContents []logFileContent
}

func NewLogReader(directoryFlag, timeFlag *string) ILogReader {
	return &LogReader{
		directoryFlag: string(*directoryFlag),
		timeFlag:      string(*timeFlag),
	}
}

const (
	suffixFile = ".log"
)

func (l *LogReader) setDuration() {
	stringLength := len(l.timeFlag)
	timeFlagRune := []rune(l.timeFlag)

	timeAgeInt, _ := strconv.Atoi(string(timeFlagRune[0 : stringLength-1]))
	l.duration = time.Duration(time.Minute * time.Duration(timeAgeInt)).Minutes()
}

func (l *LogReader) setFileLogPaths() {
	files, err := ioutil.ReadDir(l.directoryFlag)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), suffixFile) {
			if file.ModTime().Minute() <= int(l.duration) {
				l.logFilePaths = append(l.logFilePaths, filepath.Join(l.directoryFlag, file.Name()))
			}
		}
	}
}

func (l *LogReader) setLogFileContents() {
	now := time.Now().UTC()

	for _, logFilePath := range l.logFilePaths {
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
			if contentAging <= l.duration {
				l.logFileContents = append(l.logFileContents, logFileContent{
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
}

func (l *LogReader) PrintLogFileContents() {
	l.setDuration()
	l.setFileLogPaths()
	l.setLogFileContents()

	sort.SliceStable(l.logFileContents, func(i, j int) bool {
		return l.logFileContents[i].aging < l.logFileContents[j].aging
	})

	for _, logContents := range l.logFileContents {
		fmt.Println(logContents.text)
	}
}
