package functions

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(path string) string {
	entries, err := os.ReadDir(path)
	size := 0
	if err != nil {
		if strings.Contains(err.Error(), "not a directory") {
			f, fileError := os.Lstat(path)
			if fileError != nil {
				return fmt.Sprintf("Error: %s", fileError.Error())
			}
			size = int(f.Size())
			return fmt.Sprintf("%dB", size)
		}
		return fmt.Sprintf("Error: %s", err.Error())
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fileInfo, err := entry.Info()
		if err != nil {
			return fmt.Sprintf("Error: %s", err.Error())
		}
		size += int(fileInfo.Size())
	}

	return fmt.Sprintf("%dB", size)
}
