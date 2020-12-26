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
	reader.FieldsPerRecord = 2
	lines, err := reader.ReadAll()
	
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	for _, line := range lines {
		key[line[0]] = line[1]
	}

	return
}
