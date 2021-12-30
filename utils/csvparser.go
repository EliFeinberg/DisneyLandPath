package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

// parse csv files and return columns as a list of lists of strings
func ParseCSV(filePath string) [][]string {
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create csv reader
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	// read all rows
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return rows
}

// Calculates the index of a ride time given its string time
func CalculateIdx(time string) int {
	// convert time to int
	colon := strings.Index(time, ":")
	hour, err := strconv.Atoi(time[:colon])
	if err != nil {
		panic(err)
	}
	minute, err := strconv.Atoi(time[colon+1:])
	if err != nil {
		panic(err)
	}
	return (hour*2 + minute/30)
}
