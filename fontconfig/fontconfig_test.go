package fontconfig

import "testing"

func TestInit(t *testing.T) {
	t.Log("Fontconfig version", FcGetVersion())
}
