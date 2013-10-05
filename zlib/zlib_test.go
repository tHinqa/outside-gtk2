package zlib

import "testing"

func TestInit(t *testing.T) {
	t.Log("Zlib version", ZlibVersion())
}
