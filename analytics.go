package main

import (
	"flag"
	"logreader/helper"
)

func main() {
	timeAge := flag.String("t", "10m", "for specify last n minutes of the data to get")
	directory := flag.String("d", ".", "to specify the directory containing the log")

	flag.Parse()

	logReader := helper.NewLogReader(directory,timeAge)
	logReader.PrintLogFileContents()
}
