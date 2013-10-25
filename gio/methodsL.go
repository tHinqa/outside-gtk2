// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type LoadableIcon struct{}

var (
	LoadableIconGetType func() O.Type

	loadableIconLoad       func(l *LoadableIcon, size int, t **T.Char, cancellable *Cancellable, err **T.GError) *InputStream
	loadableIconLoadAsync  func(l *LoadableIcon, size int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	loadableIconLoadFinish func(l *LoadableIcon, res *AsyncResult, typ **T.Char, err **T.GError) *InputStream
)

func (l *LoadableIcon) Load(size int, t **T.Char, cancellable *Cancellable, err **T.GError) *InputStream {
	return loadableIconLoad(l, size, t, cancellable, err)
}
func (l *LoadableIcon) LoadAsync(size int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	loadableIconLoadAsync(l, size, cancellable, callback, userData)
}
func (l *LoadableIcon) LoadFinish(res *AsyncResult, typ **T.Char, err **T.GError) *InputStream {
	return loadableIconLoadFinish(l, res, typ, err)
}
