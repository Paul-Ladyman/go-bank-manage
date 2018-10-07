package csvutil

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCsvReader(t *testing.T) {
	expected := [][]string{{"first column", "second column"}, {"data 1", "data 2"}}
	result := CsvReader("./test/example.csv")
	fmt.Print(result)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("%q was not equal to %q", result, expected)
	}
}
