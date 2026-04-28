package tests

import (
	"code/cmd/functions"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	res := functions.GetPathSize("../testdata", false, false)
	require.Equal(t, "32248B", res)

	res = functions.GetPathSize("../testdata", false, true)
	require.Equal(t, "39170B", res)

	res = functions.GetPathSize("../testdata", true, false)
	require.Equal(t, "31.5KB", res)

	res = functions.GetPathSize("../testdata", true, true)
	require.Equal(t, "38.3KB", res)

	res = functions.GetPathSize("otherpath", true, true)
	require.Equal(t, "Error: open otherpath: no such file or directory", res)

	res = functions.GetPathSize("../testdata/emptydir", false, false)
	require.Equal(t, "0B", res)
}

func TestFile(t *testing.T) {
	res := functions.GetPathSize("../testdata/text.txt", false, false)
	require.Equal(t, "25326B", res)

	res = functions.GetPathSize("../testdata/text.txt", false, true)
	require.Equal(t, "0B", res)

	res = functions.GetPathSize("../testdata/text.txt", true, false)
	require.Equal(t, "24.7KB", res)
}
