package fontconfig

import "testing"

func TestInit(t *testing.T) {
	t.Log("FcGetVersion():", FcGetVersion())
}
