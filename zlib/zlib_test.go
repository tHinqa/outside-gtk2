package zlib

import "testing"

func TestInit(t *testing.T) {
	t.Log("ZlibVersion()", ZlibVersion())
	t.Logf("ZlibCompileFlags() %b", ZlibCompileFlags())
}
