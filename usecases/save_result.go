package usecases

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveResult(res interface{}, destination string) error {
	buf, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return fmt.Errorf("Error encoditg result")
	}

	file, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Error opening file")
	}
	defer file.Close()

	if _, err := file.Write(buf); err != nil {
		return fmt.Errorf("Error writing to file")
	}

	return nil
}
