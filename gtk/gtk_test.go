package gtk

import (
	o "github.com/tHinqa/outside"
	"testing"
)

func init() {
	Init(nil, nil) // Gtk Init prereq for About
}

func TestIni(t *testing.T) {
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

func TestAbout(t *testing.T) {
	a, _ := AboutDialogNew()
	t.Logf("%+v", a)
	a.SetProgramName("Hello.World.exe")
	t.Log(a.ProgramName())
	a.SetWrapLicense(true)
	t.Log(a.WrapLicense())
	a.SetDocumenters([]string{
		"Tony Wilson",
		"Foo, Bar, and Associates",
		"Rhubarb, Roux Bob and Roo Baab"})
	t.Log(a.Documenters())
}
