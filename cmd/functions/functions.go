package functions

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(path string, humanReadable, withHidden, recursive bool) string {
	entries, err := os.ReadDir(path)
	size := 0
	if err != nil {
		if strings.Contains(err.Error(), "not a directory") {
			f, fileError := os.Lstat(path)
			if fileError != nil {
				return fmt.Sprintf("Error: %s", fileError.Error())
			}
			if strings.HasPrefix(f.Name(), ".") && withHidden {
				size = int(f.Size())
			} else {
				size = int(f.Size())
			}
			return formatSize(size, humanReadable)
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

		if strings.HasPrefix(fileInfo.Name(), ".") {
			if withHidden {
				size += int(fileInfo.Size())
			}
		} else {
			size += int(fileInfo.Size())
		}
	}

	return formatSize(size, humanReadable)
}

func formatSize(sizeInBytes int, humanReadable bool) string {
	if !humanReadable {
		return fmt.Sprintf("%dB", sizeInBytes)
	}

	const (
		_  = iota //ignore first value by assigning to blank identifier
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	res := 0.0
	unit := "B"

	switch {
	case sizeInBytes >= EB:
		res = float64(sizeInBytes) / EB
		unit = "EB"
	case sizeInBytes >= PB:
		res = float64(sizeInBytes) / PB
		unit = "PB"
	case sizeInBytes >= TB:
		res = float64(sizeInBytes) / TB
		unit = "TB"
	case sizeInBytes >= GB:
		res = float64(sizeInBytes) / GB
		unit = "GB"
	case sizeInBytes >= MB:
		res = float64(sizeInBytes) / MB
		unit = "MB"
	case sizeInBytes >= KB:
		res = float64(sizeInBytes) / KB
		unit = "KB"
	default:
		return fmt.Sprintf("%dB", sizeInBytes)
	}

	return fmt.Sprintf("%.1f%s", res, unit)
}
