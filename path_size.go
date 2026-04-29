package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, humanReadable, withHidden, recursive bool) (string, error) {
	// Если переданный путь - конечный файл
	if _, err := os.ReadDir(path); err != nil {
		if isFile(err) {
			size, err := getFileSize(path, withHidden)
			if err != nil {
				return "", err
			}
			return formatSize(size, humanReadable), nil
		} else {
			return "", err
		}
	} else {
		size, err := getSize(path, withHidden, recursive)
		if err != nil {
			return "", err
		}

		return formatSize(size, humanReadable), nil
	}
}

func isFile(err error) bool {
	return strings.Contains(err.Error(), "not a directory")
}

func getSize(dirPath string, withHidden, recursive bool) (int, error) {
	entries, err := os.ReadDir(dirPath)
	size := 0
	if err != nil {
		if isFile(err) {
			size, err = getFileSize(dirPath, withHidden)
			if err != nil {
				return 0, err
			}
			return size, nil
		} else {
			return 0, err
		}
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if strings.HasPrefix(entry.Name(), ".") && !withHidden {
				continue
			}
			if recursive {
				s, err := getSize(filepath.Join(dirPath, entry.Name()), withHidden, recursive)
				if err != nil {
					return 0, err
				}
				size += s
			} else {
				continue
			}
		} else {
			s, err := getFileSize(filepath.Join(dirPath, entry.Name()), withHidden)
			if err != nil {
				return 0, err
			}
			size += s
		}
	}

	return size, nil
}

func getFileSize(path string, withHidden bool) (int, error) {
	fileInfo, fileError := os.Lstat(path)
	if fileError != nil {
		return 0, fileError
	}

	if strings.HasPrefix(fileInfo.Name(), ".") {
		if withHidden {
			return int(fileInfo.Size()), nil
		}
	} else {
		return int(fileInfo.Size()), nil
	}
	return 0, nil
}

func formatSize(sizeInBytes int, humanReadable bool) string {
	if !humanReadable {
		return fmt.Sprintf("%dB", sizeInBytes)
	}

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	switch {
	case sizeInBytes >= EB:
		return fmt.Sprintf("%.1fEB", float64(sizeInBytes)/EB)
	case sizeInBytes >= PB:
		return fmt.Sprintf("%.1fPB", float64(sizeInBytes)/PB)
	case sizeInBytes >= TB:
		return fmt.Sprintf("%.1fTB", float64(sizeInBytes)/TB)
	case sizeInBytes >= GB:
		return fmt.Sprintf("%.1fGB", float64(sizeInBytes)/GB)
	case sizeInBytes >= MB:
		return fmt.Sprintf("%.1fMB", float64(sizeInBytes)/MB)
	case sizeInBytes >= KB:
		return fmt.Sprintf("%.1fKB", float64(sizeInBytes)/KB)
	default:
		return fmt.Sprintf("%dB", sizeInBytes)
	}
}
