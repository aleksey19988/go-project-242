package code

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, humanReadable, withHidden, recursive bool) (string, error) {
	entries, err := os.ReadDir(path)
	size := 0
	if err != nil {
		if strings.Contains(err.Error(), "not a directory") {
			f, fileError := os.Lstat(path)
			if fileError != nil {
				return "", err
			}
			size += getFileSize(f, strings.HasPrefix(f.Name(), "."), withHidden)
			return formatSize(size, humanReadable), nil
		}
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if recursive {
				s, err := getIncludeDirSize(filepath.Join(path, entry.Name()), withHidden, recursive)
				if err != nil {
					return "", err
				}
				size += s
			} else {
				continue
			}
		}
		fileInfo, err := entry.Info()
		if err != nil {
			return "", err
		}

		size += getFileSize(fileInfo, strings.HasPrefix(fileInfo.Name(), "."), withHidden)
	}

	return formatSize(size, humanReadable), nil
}

func getIncludeDirSize(path string, withHidden, recursive bool) (int, error) {
	entries, err := os.ReadDir(path)
	size := 0
	if err != nil {
		if strings.Contains(err.Error(), "not a directory") {
			f, fileError := os.Lstat(path)
			if fileError != nil {
				return 0, fileError
			}
			return getFileSize(f, strings.HasPrefix(f.Name(), "."), withHidden), nil
		}
		return 0, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if recursive {
				s, err := getIncludeDirSize(filepath.Join(path, entry.Name()), withHidden, recursive)
				if err != nil {
					return 0, err
				}
				size += s
			} else {
				continue
			}
		}
		fileInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}

		size += int(fileInfo.Size())
	}

	return size, nil
}

var getFileSize = func(fileInfo fs.FileInfo, hasPrefix, withHidden bool) int {
	if strings.HasPrefix(fileInfo.Name(), ".") {
		if withHidden {
			return int(fileInfo.Size())
		}
	} else {
		return int(fileInfo.Size())
	}
	return 0
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
