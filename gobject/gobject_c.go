// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Callback func() T.Dummy

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
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)
	Data      T.Gpointer
	Notifiers *ClosureNotifyData
}

var (
	ClosureGetType   func() Type
	ClosureNewObject func(sizeofClosure uint, object *Object) *Closure
	ClosureNewSimple func(sizeofClosure uint, data T.Gpointer) *Closure

	ClosureAddFinalizeNotifier      func(c Closure, notifyData T.Gpointer, notifyFunc ClosureNotify)
	ClosureAddInvalidateNotifier    func(c Closure, notifyData T.Gpointer, notifyFunc ClosureNotify)
	ClosureAddMarshalGuards         func(c Closure, preMarshalData T.Gpointer, preMarshalNotify ClosureNotify, postMarshalData T.Gpointer, postMarshalNotify ClosureNotify)
	ClosureInvalidate               func(c Closure)
	ClosureInvoke                   func(c Closure, returnValue *Value, nParamValues uint, paramValues *Value, invocationHint T.Gpointer)
	ClosureRef                      func(c Closure) *Closure
	ClosureRemoveFinalizeNotifier   func(c Closure, notifyData T.Gpointer, notifyFunc ClosureNotify)
	ClosureRemoveInvalidateNotifier func(c Closure, notifyData T.Gpointer, notifyFunc ClosureNotify)
	ClosureSetMarshal               func(c Closure, marshal ClosureMarshal)
	ClosureSetMetaMarshal           func(c Closure, marshalData T.Gpointer, metaMarshal ClosureMarshal)
	ClosureSink                     func(c Closure)
	ClosureUnref                    func(c Closure)
)

func (c Closure) AddFinalizeNotifier(notifyData T.Gpointer, notifyFunc ClosureNotify) {
	ClosureAddFinalizeNotifier(c, notifyData, notifyFunc)
}
func (c Closure) AddInvalidateNotifier(notifyData T.Gpointer, notifyFunc ClosureNotify) {
	ClosureAddInvalidateNotifier(c, notifyData, notifyFunc)
}
func (c Closure) AddMarshalGuards(preMarshalData T.Gpointer, preMarshalNotify ClosureNotify, postMarshalData T.Gpointer, postMarshalNotify ClosureNotify) {
	ClosureAddMarshalGuards(c, preMarshalData, preMarshalNotify, postMarshalData, postMarshalNotify)
}
func (c Closure) Invalidate() { ClosureInvalidate(c) }
func (c Closure) Invoke(returnValue *Value, nParamValues uint, paramValues *Value, invocationHint T.Gpointer) {
	ClosureInvoke(c, returnValue, nParamValues, paramValues, invocationHint)
}
func (c Closure) Ref() *Closure { return ClosureRef(c) }
func (c Closure) RemoveFinalizeNotifier(notifyData T.Gpointer, notifyFunc ClosureNotify) {
	ClosureRemoveFinalizeNotifier(c, notifyData, notifyFunc)
}
func (c Closure) RemoveInvalidateNotifier(notifyData T.Gpointer, notifyFunc ClosureNotify) {
	ClosureRemoveInvalidateNotifier(c, notifyData, notifyFunc)
}
func (c Closure) SetMarshal(marshal ClosureMarshal) { ClosureSetMarshal(c, marshal) }
func (c Closure) SetMetaMarshal(marshalData T.Gpointer, metaMarshal ClosureMarshal) {
	ClosureSetMetaMarshal(c, marshalData, metaMarshal)
}
func (c Closure) Sink()  { ClosureSink(c) }
func (c Closure) Unref() { ClosureUnref(c) }

type ClosureNotify func(data T.Gpointer, closure *Closure)

type ClosureNotifyData struct {
	Data   T.Gpointer
	Notify ClosureNotify
}

type ClosureMarshal func(
	Closure *Closure,
	returnValue *Value,
	nParamValues uint,
	paramValues *Value,
	invocationHint, marshalData T.Gpointer)
