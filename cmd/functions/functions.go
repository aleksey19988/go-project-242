package functions

import (
	"fmt"
	"os"
)

func GetPathSize(path string) string {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Sprintf("Error: %s\n", err.Error())
	}
	size := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fileInfo, err := entry.Info()
		if err != nil {
			return fmt.Sprintf("Error: %s\n", err.Error())
		}
		size += int(fileInfo.Size())
	}

	return fmt.Sprintf("%dB", size)
}
