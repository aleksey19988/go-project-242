package tests

import (
	"code/cmd/functions"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	res := functions.GetPathSize("../testdata", false)
	require.Equal(t, "6936B", res)
}

func TestEmptyPath(t *testing.T) {
	res := functions.GetPathSize("", false)
	require.Equal(t, "Error: open : no such file or directory", res)
}

func TestUnexpectedPath(t *testing.T) {
	res := functions.GetPathSize("otherpath", false)
	require.Equal(t, "Error: open otherpath: no such file or directory", res)
}

func TestFile(t *testing.T) {
	res := functions.GetPathSize("../testdata/test.txt", false)
	require.Equal(t, "14B", res)
}

func TestHumanReadable(t *testing.T) {
	res := functions.GetPathSize("../testdata", true)
	require.Equal(t, "6.8KB", res)
}
