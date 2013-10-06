package intl

import "testing"

func TestInit(t *testing.T) {
	t.Log(Libintl_textdomain(""))
}
