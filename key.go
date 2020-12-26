package transliterate

import (
	"encoding/csv"
	"fmt"
	"os"
)

// LoadKey loads
func LoadKey(path string) (key map[string]string, err error) {
	csvfile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = 2
	
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	key = map[string]string{}
	for _, line := range lines[1:] {
		key[line[0]] = line[1]
	}

	return
}
