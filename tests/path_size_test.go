package tests

import (
	"code/cmd/functions"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	res := functions.GetPathSize("../testdata")
	require.Equal(t, "6936B", res)
}

func TestEmptyPath(t *testing.T) {
	res := functions.GetPathSize("")
	require.Equal(t, "Error: open : no such file or directory\n", res)
}

func TestUnexpectedPath(t *testing.T) {
	res := functions.GetPathSize("otherpath")
	require.Equal(t, "Error: open otherpath: no such file or directory\n", res)
}
