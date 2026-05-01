package tests

import (
	"code"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getFixturePath(name string) string {
	path := filepath.Join("../testdata/fixture", name)

	return path
}

func TestGetPathSize_Basic(t *testing.T) {
	path := getFixturePath("dir1")

	got, err := code.GetPathSize(path, false, false, false)
	require.NoError(t, err)

	// Включён только a.txt → "a" = 1 байт + перенос строки
	expected := int64(len("a\n"))
	assert.Equal(t, fmt.Sprintf("%dB", expected), got)
}

func TestGetPathSize_Recursive(t *testing.T) {
	path := getFixturePath("dir1")

	got, err := code.GetPathSize(path, true, false, false)
	require.NoError(t, err)

	// Включены: a.txt + nested/deep.txt → "a" + "deep"
	expected := int64(len("a\n") + len("deep\n"))
	assert.Equal(t, fmt.Sprintf("%dB", expected), got)
}

func TestGetPathSize_All(t *testing.T) {
	path := getFixturePath("dir1")

	got, err := code.GetPathSize(path, true, false, true)
	require.NoError(t, err)

	// Включены: a.txt, .hidden.txt, nested/deep.txt → "a" + "hidden" + "deep" + переносы строк
	expected := int64(len("a\n") + len("hidden\n") + len("deep\n"))
	assert.Equal(t, fmt.Sprintf("%dB", expected), got)
}

func TestGetPathSize_EmptyDir(t *testing.T) {
	path := getFixturePath("empty_dir")

	got, err := code.GetPathSize(path, false, false, false)
	require.NoError(t, err)

	// Пустая директория → 0 байт
	expected := int64(0)
	assert.Equal(t, fmt.Sprintf("%dB", expected), got)
}

func TestGetPathSize_SingleFile(t *testing.T) {
	path := getFixturePath("file.txt")

	got, err := code.GetPathSize(path, false, false, false)
	require.NoError(t, err)

	// Один файл → "file" = 4 байта
	expected := int64(len("file\n"))
	assert.Equal(t, fmt.Sprintf("%dB", expected), got)
}

func TestGetPathSize_HiddenFilesIgnored(t *testing.T) {
	path := getFixturePath("dir1")

	got, err := code.GetPathSize(path, true, false, false)
	require.NoError(t, err)

	// Включены: a.txt, nested/deep.txt → "a" + "deep" + переносы строк
	expected := int64(len("a\n") + len("deep\n"))
	assert.Equal(t, fmt.Sprintf("%dB", expected), got)
}

func TestGetPathSize_HumanReadable(t *testing.T) {
	path := getFixturePath("file.txt")

	got, err := code.GetPathSize(path, false, true, false)
	require.NoError(t, err)

	// Один файл → "file\n" = 5 байта → 5B в человекочитаемом формате
	expected := "5B"
	assert.Equal(t, expected, got)
}

func TestGetPathSize_HumanReadableKilobytes(t *testing.T) {
	path := getFixturePath("large_file.txt")

	got, err := code.GetPathSize(path, false, true, false)
	require.NoError(t, err)

	// Файл размером 2048 байт → 2.0KB в человекочитаемом формате
	expected := "2.0KB"
	assert.Equal(t, expected, got)
}
