package utils

import (
	"encoding/csv"
	"os"
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
