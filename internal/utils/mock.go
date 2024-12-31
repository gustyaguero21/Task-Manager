package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func OpenMock[T any](path string) (T, error) {
	var result T

	bytes, err := os.ReadFile(path)
	if err != nil {
		return result, fmt.Errorf("error reading json file: %w", err)
	}

	if err := json.Unmarshal(bytes, &result); err != nil {
		return result, fmt.Errorf("error parsing data: %w", err)
	}

	return result, nil
}
