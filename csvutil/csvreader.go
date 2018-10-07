package csvutil

import (
	"encoding/csv"
	"io/ioutil"
	"strings"
)

// CsvReader reads a csv file and returns all the data
func CsvReader(filepath string) [][]string {
	csvFileData, _ := ioutil.ReadFile(filepath)
	csvString := string(csvFileData)

	r := csv.NewReader(strings.NewReader(csvString))

	records, _ := r.ReadAll()
	return records
}
