package utils

import (
	"os"
)

func OpenMock(path string) []byte {

	readFile, readErr := os.ReadFile(path)
	if readErr != nil {
		return nil
	}
	return readFile

}
