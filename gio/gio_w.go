// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Win32InputStream struct {
	Parent InputStream
	_      *struct{}
}

var (
	Win32InputStreamGetType func() O.Type
	Win32InputStreamNew     func(handle *T.Void, closeHandle bool) *InputStream

	win32InputStreamSetCloseHandle func(stream *Win32InputStream, closeHandle bool)
	win32InputStreamGetCloseHandle func(stream *Win32InputStream) bool
	win32InputStreamGetHandle      func(stream *Win32InputStream) *T.Void
)

func (w *Win32InputStream) SetCloseHandle(closeHandle bool) {
	win32InputStreamSetCloseHandle(w, closeHandle)
}
func (w *Win32InputStream) GetCloseHandle() bool { return win32InputStreamGetCloseHandle(w) }
func (w *Win32InputStream) GetHandle() *T.Void   { return win32InputStreamGetHandle(w) }

type Win32OutputStream struct {
	Parent OutputStream
	_      *struct{}
}

var (
	Win32OutputStreamGetType func() O.Type
	Win32OutputStreamNew     func(handle *T.Void, closeHandle bool) *OutputStream

	win32OutputStreamSetCloseHandle func(stream *Win32OutputStream, closeHandle bool)
	win32OutputStreamGetCloseHandle func(stream *Win32OutputStream) bool
	win32OutputStreamGetHandle      func(stream *Win32OutputStream) *T.Void
)

func (w *Win32OutputStream) SetCloseHandle(closeHandle bool) {
	win32OutputStreamSetCloseHandle(w, closeHandle)
}
func (w *Win32OutputStream) GetCloseHandle() bool { return win32OutputStreamGetCloseHandle(w) }
func (w *Win32OutputStream) GetHandle() *T.Void   { return win32OutputStreamGetHandle(w) }
