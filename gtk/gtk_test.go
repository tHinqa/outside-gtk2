package gtk

import "testing"

func TestInit(t *testing.T) {
	v := Gtk_check_version(2, 0, 0)
	if v != "" {
		t.Error(v)
	}
}
