package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Paul-Ladyman/go-bank-manage/csvutil"
)

func main() {
	filename := "./statement.csv"
	csvFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	csvString := string(csvFileData)

	r := csv.NewReader(strings.NewReader(csvString))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	statement := csvutil.CsvParser(records)
	fmt.Println("Statement:")
	fmt.Println(statement)
}
