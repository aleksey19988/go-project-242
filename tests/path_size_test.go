package tests

import (
	"code/cmd/functions"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	res := functions.GetPathSize("../testdata", false, false, false)
	require.Equal(t, "32248B", res)

	res = functions.GetPathSize("../testdata", false, true, false)
	require.Equal(t, "39170B", res)

	res = functions.GetPathSize("../testdata", true, false, false)
	require.Equal(t, "31.5KB", res)

	res = functions.GetPathSize("../testdata", true, true, false)
	require.Equal(t, "38.3KB", res)

	res = functions.GetPathSize("otherpath", true, true, false)
	require.Equal(t, "Error: open otherpath: no such file or directory", res)

	res = functions.GetPathSize("../testdata/emptydir", false, false, false)
	require.Equal(t, "0B", res)
}

func TestFile(t *testing.T) {
	res := functions.GetPathSize("../testdata/text.txt", false, false, false)
	require.Equal(t, "25326B", res)

	res = functions.GetPathSize("../testdata/text.txt", false, true, false)
	require.Equal(t, "0B", res)

	res = functions.GetPathSize("../testdata/text.txt", true, false, false)
	require.Equal(t, "24.7KB", res)
}
