// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Closure struct {
	Bits uint
	// RefCount : 15
	// MetaMarshal : 1
	// NGuards : 1
	// NFnotifiers : 2
	// NInotifiers : 8
	// InInotify : 1
	// Floating : 1
	// DerivativeFlag : 1
	// InMarshal : 1
	// IsInvalid : 1
	Marshal func(
		closure *Closure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)
	Data      T.Gpointer
	Notifiers *ClosureNotifyData
}

var (
	ClosureGetType   func() Type
	ClosureNewObject func(sizeofClosure uint, object *Object) *Closure
	ClosureNewSimple func(sizeofClosure uint, data T.Gpointer) *Closure

	closureAddFinalizeNotifier      func(c Closure, notifyData T.Gpointer, notifyFunc ClosureNotify)
	closureAddInvalidateNotifier    func(c Closure, notifyData T.Gpointer, notifyFunc ClosureNotify)
	closureAddMarshalGuards         func(c Closure, preMarshalData T.Gpointer, preMarshalNotify ClosureNotify, postMarshalData T.Gpointer, postMarshalNotify ClosureNotify)
	closureInvalidate               func(c Closure)
	closureInvoke                   func(c Closure, returnValue *T.GValue, nParamValues uint, paramValues *T.GValue, invocationHint T.Gpointer)
	closureRef                      func(c Closure) *Closure
	closureRemoveFinalizeNotifier   func(c Closure, notifyData T.Gpointer, notifyFunc ClosureNotify)
	closureRemoveInvalidateNotifier func(c Closure, notifyData T.Gpointer, notifyFunc ClosureNotify)
	closureSetMarshal               func(c Closure, marshal ClosureMarshal)
	closureSetMetaMarshal           func(c Closure, marshalData T.Gpointer, metaMarshal ClosureMarshal)
	closureSink                     func(c Closure)
	closureUnref                    func(c Closure)
)

func (c Closure) AddFinalizeNotifier(notifyData T.Gpointer, notifyFunc ClosureNotify) {
	closureAddFinalizeNotifier(c, notifyData, notifyFunc)
}
func (c Closure) AddInvalidateNotifier(notifyData T.Gpointer, notifyFunc ClosureNotify) {
	closureAddInvalidateNotifier(c, notifyData, notifyFunc)
}
func (c Closure) AddMarshalGuards(preMarshalData T.Gpointer, preMarshalNotify ClosureNotify, postMarshalData T.Gpointer, postMarshalNotify ClosureNotify) {
	closureAddMarshalGuards(c, preMarshalData, preMarshalNotify, postMarshalData, postMarshalNotify)
}
func (c Closure) Invalidate() { closureInvalidate(c) }
func (c Closure) Invoke(returnValue *T.GValue, nParamValues uint, paramValues *T.GValue, invocationHint T.Gpointer) {
	closureInvoke(c, returnValue, nParamValues, paramValues, invocationHint)
}
func (c Closure) Ref() *Closure { return closureRef(c) }
func (c Closure) RemoveFinalizeNotifier(notifyData T.Gpointer, notifyFunc ClosureNotify) {
	closureRemoveFinalizeNotifier(c, notifyData, notifyFunc)
}
func (c Closure) RemoveInvalidateNotifier(notifyData T.Gpointer, notifyFunc ClosureNotify) {
	closureRemoveInvalidateNotifier(c, notifyData, notifyFunc)
}
func (c Closure) SetMarshal(marshal ClosureMarshal) { closureSetMarshal(c, marshal) }
func (c Closure) SetMetaMarshal(marshalData T.Gpointer, metaMarshal ClosureMarshal) {
	closureSetMetaMarshal(c, marshalData, metaMarshal)
}
func (c Closure) Sink()  { closureSink(c) }
func (c Closure) Unref() { closureUnref(c) }

type ClosureNotify func(data T.Gpointer, closure *Closure)

type ClosureNotifyData struct {
	Data   T.Gpointer
	Notify ClosureNotify
}

type ClosureMarshal func(
	closure *Closure,
	returnValue *T.GValue,
	nParamValues uint,
	paramValues *T.GValue,
	invocationHint, marshalData T.Gpointer)
