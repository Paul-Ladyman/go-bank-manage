package csvutil

import (
	"github.com/Paul-Ladyman/go-bank-manage/domain"
)

// CsvParser takes a multi-dimensional array of CSV data and return structs
func CsvParser(csv [][]string) domain.Transaction {
	date := csv[1][0]
	return domain.Transaction{Date: date}
}
