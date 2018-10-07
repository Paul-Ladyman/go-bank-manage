package main

import (
	"fmt"

	"github.com/Paul-Ladyman/go-bank-manage/csvutil"
)

func main() {
	records := csvutil.CsvReader("./statement.csv")
	statement := csvutil.CsvParser(records)
	fmt.Println("Statement:")
	fmt.Println(statement)
}
