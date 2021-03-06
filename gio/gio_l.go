// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type LoadableIcon struct{}

var (
	LoadableIconGetType func() O.Type

	LoadableIconLoad       func(l *LoadableIcon, size int, t **T.Char, cancellable *Cancellable, err **L.Error) *InputStream
	LoadableIconLoadAsync  func(l *LoadableIcon, size int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	LoadableIconLoadFinish func(l *LoadableIcon, res *AsyncResult, typ **T.Char, err **L.Error) *InputStream
)

func (l *LoadableIcon) Load(size int, t **T.Char, cancellable *Cancellable, err **L.Error) *InputStream {
	return LoadableIconLoad(l, size, t, cancellable, err)
}
func (l *LoadableIcon) LoadAsync(size int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	LoadableIconLoadAsync(l, size, cancellable, callback, userData)
}
func (l *LoadableIcon) LoadFinish(res *AsyncResult, typ **T.Char, err **L.Error) *InputStream {
	return LoadableIconLoadFinish(l, res, typ, err)
}
