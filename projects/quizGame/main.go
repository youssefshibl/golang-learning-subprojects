package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	ReadCSVFile()
}

const DefaultFileName = "problems.csv"
const TimeLimit = 20

var score int
var correctScore int

func ReadCSVFile() {

	CsvFileNamePointer := flag.String("csv", DefaultFileName, "a csv file name of problems")
	flag.Parse()
	CsvFileName := *CsvFileNamePointer

	f, err := os.Open(CsvFileName)
	if err != nil {
		// panic(err)
		log.Fatal("Unable to open this file ", CsvFileName)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+CsvFileName, err)
	}
	score = len(records)

	askChanel := AskUser(records)

	timer := time.NewTimer(time.Duration(TimeLimit) * time.Second)

	select {
	case <-timer.C:
	case <-askChanel:
	}

	fmt.Printf("\nYour Score is %d/%d \n", correctScore, score)
}

func AskUser(questions [][]string) chan bool {
	channel := make(chan bool)
	go func() {
		defer func() { channel <- true }()

		for _, question := range questions {

			answer := ""
			fmt.Print(question[0], " :")
			fmt.Scan(&answer)
			if strings.TrimSpace(answer) == question[1] {
				correctScore++
			}

		}

	}()
	return channel

}
