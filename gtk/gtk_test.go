package gtk

import (
	o "github.com/tHinqa/outside"
	G "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// r "reflect"
	"testing"
	// "unsafe"
)

func init() {
	Init(nil, nil) // Gtk Init prereq
}

func TestIni(t *testing.T) {
	minor := uint(12)
	v := CheckVersion(2, minor, 0)
	if v != "" {
		t.Errorf("Expected version 2.%d.0+; got: %s\nNote: For 'micro mismatch' read 'minor mismatch'", minor, v)
	}
	a := o.GetData("gtk_major_version")
	b := o.GetData("gtk_minor_version")
	c := o.GetData("gtk_micro_version")
	t.Logf("outside.GetData(\"gtk_..._version\") says major: %v minor: %v micro: %v", *(a).(*uint), *(b).(*uint), *(c).(*uint))
}

func AsPointer(a func()) T.Gpointer { return nil }

func wDestroy() (_ T.Dummy) {
	MainQuit()
	return
}

func TestAbout(t *testing.T) {
	// aw := AboutDialogNew()
	// a := aw.AsAboutDialog() or
	aw, a := NewAboutDialog()
	aw.Show()
	a.SetProgramName("outside-gtk2")
	p := []string{"Tony Wilson http://twitter.com/tHinqa",
		"Foo, Bar, and Associates",
		"Rhubarb, Roux Bob and Roo Baab"}
	l := `Copyright (c) 2013 Tony Wilson. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of the copyright owner nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.`
	c := "An easy way to implement Gtk+2 in Go. Boring rhubarb rhubarb rhubarb."
	a.SetWebsite("https://github.com/tHinqa/outside-gtk2")

	a.SetWebsiteLabel("Source repository")
	a.SetAuthors(p[:1])
	a.SetArtists(p[1:])
	a.SetComments(c)
	a.SetVersion("v0.0.0alpha")
	a.SetDocumenters(p)
	a.SetLicense(l)
	a.SetLogoIconName("")
	G.SignalConnect(aw.AsPointer(), "close", wDestroy, nil)    // [X]
	G.SignalConnect(aw.AsPointer(), "response", wDestroy, nil) // [Close]

	// ShowAboutDialog(nil,
	// 	"program-name", "outside-gtk2",
	// 	"title", "Hello World",
	// 	"authors", p[:1],
	// 	// "artists", p[1:], // too many args for syscall (17)
	// 	"documenters", p,
	// 	"comments", c,
	// 	"license", l, 0)

	// Main()
}

func TestWindow(t *testing.T) {
	// wt := WindowNew(WINDOW_TOPLEVEL)
	// w := wt.AsWindow() or
	wt, w := NewWindow(WINDOW_TOPLEVEL)
	wt.Show()
	w.Move(960, 0)
	w.Resize(800, 600)
	w.SetOpacity(.5)
	w.SetTitle("Go does Gtk+2")
	G.SignalConnect(wt.AsPointer(), "destroy", wDestroy, nil)
	Main()
}
