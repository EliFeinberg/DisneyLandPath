package main

import (
	"fmt"

	"github.com/EliFeinberg/DisneyLandPath/utils"
)

func main() {
	// test parsing csv
	rows := utils.ParseCSV("./CSV/rides.csv")
	fmt.Println(rows)
}
