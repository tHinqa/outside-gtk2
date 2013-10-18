package gtk

import "testing"

func TestIni(t *testing.T) { // TestInit name conflict
	v := CheckVersion(2, 0, 0)
	if v != "" {
		t.Error(v)
	}
}
