package csvutil

import (
	"testing"

	"github.com/Paul-Ladyman/go-bank-manage/domain"
)

func TestCsvParser(t *testing.T) {
	csvData := [][]string{{"Transaction Date"}, {"1/1/1970"}}
	expected := domain.Transaction{Date: "1/1/1970"}
	result := CsvParser(csvData)

	if result != expected {
		t.Errorf("%q was not equal to %q", result, expected)
	}
}
