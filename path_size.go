package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, humanReadable, withHidden, recursive bool) (string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	size, err := getSize(path, withHidden, recursive)
	if err != nil {
		return "", err
	}
	return formatSize(size, humanReadable), nil
}

func getSize(
	path string,
	withHidden,
	recursive bool,
) (int, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	size := 0

	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				if strings.HasPrefix(entry.Name(), ".") && !withHidden {
					continue
				}
				if recursive {
					s, err := getSize(filepath.Join(path, entry.Name()), withHidden, recursive)
					if err != nil {
						return 0, err
					}
					size += s
				} else {
					continue
				}
			} else {
				if strings.HasPrefix(entry.Name(), ".") && !withHidden {
					continue
				}
				s, err := getFileSize(filepath.Join(path, entry.Name()))
				if err != nil {
					return 0, err
				}
				size += s
			}
		}
	} else {
		if strings.HasPrefix(fileInfo.Name(), ".") && !withHidden {
			return size, nil
		}
		s, err := getFileSize(path)
		if err != nil {
			return 0, err
		}
		return s, nil
	}

	return size, nil
}

func getFileSize(path string) (int, error) {
	fileInfo, fileError := os.Stat(path)
	if fileError != nil {
		return 0, fileError
	}
	return int(fileInfo.Size()), nil
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
