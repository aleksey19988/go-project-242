package tests

import (
	"code"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	res, _ := code.GetPathSize("../testdata", false, false, false)
	require.Equal(t, "32248B", res)

	res, _ = code.GetPathSize("../testdata", false, true, false)
	require.Equal(t, "39170B", res)

	res, _ = code.GetPathSize("../testdata", true, false, false)
	require.Equal(t, "31.5KB", res)

	res, _ = code.GetPathSize("../testdata", true, true, false)
	require.Equal(t, "38.3KB", res)

	res, _ = code.GetPathSize("otherpath", true, true, false)
	require.Equal(t, "Error: open otherpath: no such file or directory", res)

	res, _ = code.GetPathSize("../testdata/emptydir", false, false, false)
	require.Equal(t, "0B", res)
}

func TestFile(t *testing.T) {
	res, _ := code.GetPathSize("../testdata/text.txt", false, false, false)
	require.Equal(t, "25326B", res)

	res, _ = code.GetPathSize("../testdata/text.txt", true, false, false)
	require.Equal(t, "24.7KB", res)
}

func TestRecursive(t *testing.T) {
	res, _ := code.GetPathSize("../testdata", false, false, false)
	require.Equal(t, "32248B", res)

	res, _ = code.GetPathSize("../testdata", false, false, true)
	require.Equal(t, "79610B", res)

	res, _ = code.GetPathSize("../testdata", false, true, true)
	require.Equal(t, "86532B", res)
}
