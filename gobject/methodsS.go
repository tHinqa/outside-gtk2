// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Source struct {
	CallbackData  T.Gpointer
	CallbackFuncs *SourceCallbackFuncs
	Source_funcs  *SourceFuncs
	RefCount      uint
	Context       *T.GMainContext
	Priority      int
	Flags         uint
	SourceId      uint
	PollFds       *T.GSList
	Prev          *Source
	Next          *Source
	Name          *T.Char
	_             *struct{}
}

var (
	SourceSetClosure       func(s *Source, closure *Closure)
	SourceSetDummyCallback func(s *Source)
)

type SourceCallbackFuncs struct {
	Ref   func(cbData T.Gpointer)
	Unref func(cbData T.Gpointer)
	Get   func(cbData T.Gpointer,
		source *Source, fnc *SourceFunc, data *T.Gpointer)
}

type SourceDummyMarshal func()

type SourceFunc func(data T.Gpointer) T.Gboolean

type SourceFuncs struct {
	Prepare  func(source *Source, timeout *int) T.Gboolean
	Check    func(source *Source) T.Gboolean
	Dispatch func(
		source *Source, callback SourceFunc, userData T.Gpointer) T.Gboolean
	Finalize func(source *Source)

	ClosureCallback SourceFunc
	ClosureMarshal  SourceDummyMarshal
}