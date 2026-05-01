package tests

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func buildCLIBinary(t *testing.T) string {
	t.Helper()

	binDir := t.TempDir()
	//binDir := "bin"
	binName := "hexlet-path-size"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	binPath := filepath.Join(binDir, binName)

	cmd := exec.Command("go", "build", "-o", binPath, "./cmd/hexlet-path-size")
	//cmd.Dir = "../code"
	cmd.Dir = "../"
	cmd.Env = append(os.Environ(), "GOCACHE="+filepath.Join(binDir, "go-build-cache"))

	output, err := cmd.CombinedOutput()
	require.NoErrorf(t, err, "failed to build CLI binary: %s", string(output))

	return binPath
}

func runCLI(t *testing.T, binPath string, args ...string) (string, string, error) {
	t.Helper()

	cmd := exec.Command(binPath, args...)
	cmd.Dir = ".."

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func TestCLIOutput(t *testing.T) {
	binPath := buildCLIBinary(t)

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "basic file size",
			args:     []string{"testdata/fixture/file.txt"},
			expected: "5B\ttestdata/fixture/file.txt\n",
		},
		{
			name:     "human-readable file size",
			args:     []string{"--human", "testdata/fixture/large_file.txt"},
			expected: "2.0KB\ttestdata/fixture/large_file.txt\n",
		},
		{
			name:     "recursive directory size",
			args:     []string{"--recursive", "testdata/fixture/dir1"},
			expected: "7B\ttestdata/fixture/dir1\n",
		},
		{
			name:     "recursive directory size with hidden files",
			args:     []string{"-r", "-a", "testdata/fixture/dir1"},
			expected: "14B\ttestdata/fixture/dir1\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, err := runCLI(t, binPath, tt.args...)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, stdout)
			assert.Equal(t, "", stderr)
		})
	}
}
