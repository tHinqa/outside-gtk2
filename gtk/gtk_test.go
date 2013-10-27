package gtk

import (
	o "github.com/tHinqa/outside"
	"testing"
)

func TestIni(t *testing.T) { // TestInit name conflict
	minor := uint(10)
	v := CheckVersion(2, minor, 0)
	if v != "" {
		t.Errorf("Expected version 2.%d.0+; got: %s\nNote: For 'micro mismatch' read 'minor mismatch'", minor, v)
	}
	a := o.GetData("gtk_major_version")
	b := o.GetData("gtk_minor_version")
	c := o.GetData("gtk_micro_version")
	t.Logf("outside.GetData(\"gtk_..._version\") says major: %v minor: %v micro: %v", *(a).(*uint), *(b).(*uint), *(c).(*uint))
}
