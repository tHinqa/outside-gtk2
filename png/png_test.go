package png

import "testing"

func TestInit(t *testing.T) {
	t.Log("LibPNG version", Png_access_version_number())
}
