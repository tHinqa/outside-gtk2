// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

var (
	SignalTypeCclosureNew func(
		itype Type, structOffset uint) *Closure

	SignalNewv func(
		signalName string,
		itype Type,
		signalFlags SignalFlags,
		classClosure *Closure,
		accumulator SignalAccumulator,
		accuData T.Gpointer,
		cMarshaller SignalCMarshaller,
		returnType Type,
		nParams uint,
		paramTypes *Type) uint

	SignalNewValist func(
		signalName string,
		itype Type,
		signalFlags SignalFlags,
		classClosure *Closure,
		accumulator SignalAccumulator,
		accuData T.Gpointer,
		cMarshaller SignalCMarshaller,
		returnType Type,
		nParams uint,
		args VAList) uint

	SignalNew func(signalName string, itype Type,
		signalFlags SignalFlags, classOffset uint,
		accumulator SignalAccumulator, accuData T.Gpointer,
		cMarshaller SignalCMarshaller, returnType Type,
		nParams uint, v ...VArg) uint

	SignalNewClassHandler func(signalName string,
		itype Type, signalFlags SignalFlags,
		classHandler Callback, accumulator SignalAccumulator,
		accuData T.Gpointer, cMarshaller SignalCMarshaller,
		returnType Type, nParams uint, v ...VArg) uint

	SignalEmitv func(
		instanceAndParams *Value,
		signalId uint,
		detail T.Quark,
		returnValue *Value)

	SignalEmitValist func(
		instance T.Gpointer,
		signalId uint,
		detail T.Quark,
		varArgs VAList)

	SignalEmit func(instance T.Gpointer, signalId uint,
		detail T.Quark, v ...VArg)

	SignalEmitByName func(instance T.Gpointer,
		detailedSignal string, v ...VArg)

	SignalLookup func(name string, itype Type) uint

	SignalName func(signalId uint) string

	SignalQuery func(signalId uint, query SignalQueryFunc)

	SignalListIds func(itype Type, nIds *uint) *uint

	SignalParseName func(
		detailedSignal string,
		itype Type,
		signalIdP *uint,
		detailP *T.Quark,
		forceDetailQuark bool) bool

	SignalGetInvocationHint func(
		instance T.Gpointer) *SignalInvocationHint

	SignalStopEmission func(
		instance T.Gpointer, signalId uint, detail T.Quark)

	SignalStopEmissionByName func(
		instance T.Gpointer, detailedSignal string)

	SignalAddEmissionHook func(
		signalId uint,
		detail T.Quark,
		hookFunc SignalEmissionHook,
		hookData T.Gpointer,
		dataDestroy DestroyNotify) T.Gulong

	SignalRemoveEmissionHook func(signalId uint, hookId T.Gulong)

	SignalHasHandlerPending func(
		instance T.Gpointer,
		signalId uint,
		detail T.Quark,
		mayBeBlocked bool) bool

	SignalConnectClosureById func(
		instance T.Gpointer,
		signalId uint,
		detail T.Quark,
		closure *Closure,
		after bool) T.Gulong

	SignalConnectClosure func(
		instance T.Gpointer,
		detailedSignal string,
		closure *Closure,
		after bool) T.Gulong

	SignalConnectData func(
		instance T.Gpointer,
		detailedSignal string,
		cHandler Callback,
		data T.Gpointer,
		destroyData ClosureNotify,
		connectFlags ConnectFlags) T.Gulong

	SignalHandlerBlock func(instance T.Gpointer, handlerId T.Gulong)

	SignalHandlerUnblock func(instance T.Gpointer, handlerId T.Gulong)

	SignalHandlerDisconnect func(
		instance T.Gpointer, handlerId T.Gulong)

	SignalHandlerIsConnected func(
		instance T.Gpointer, handlerId T.Gulong) bool

	SignalHandlerFind func(
		instance T.Gpointer,
		mask SignalMatchType,
		signalId uint,
		detail T.Quark,
		closure *Closure,
		fnc T.Gpointer,
		data T.Gpointer) T.Gulong

	SignalHandlersBlockMatched func(
		instance T.Gpointer,
		mask SignalMatchType,
		signalId uint,
		detail T.Quark,
		closure *Closure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalHandlersUnblockMatched func(
		instance T.Gpointer,
		mask SignalMatchType,
		signalId uint,
		detail T.Quark,
		closure *Closure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalHandlersDisconnectMatched func(
		instance T.Gpointer,
		mask SignalMatchType,
		signalId uint,
		detail T.Quark,
		closure *Closure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalOverrideClassClosure func(
		signalId uint, instanceType Type, classClosure *Closure)

	SignalOverrideClassHandler func(
		signalName string, instanceType Type, classHandler Callback)

	SignalChainFromOverridden func(
		instanceAndParams *Value, returnValue *Value)

	SignalChainFromOverriddenHandler func(
		instance T.Gpointer, v ...VArg)

	SignalAccumulatorTrueHandled func(ihint *SignalInvocationHint,
		returnAccu, handlerReturn *Value, dummy T.Gpointer) bool

	SignalAccumulatorFirstWins func(ihint *SignalInvocationHint,
		returnAccu, handlerReturn *Value, dummy T.Gpointer) bool

	SignalHandlersDestroy func(instance T.Gpointer)
)

type SignalAccumulator func(
	ihint *SignalInvocationHint,
	returnAccu *Value,
	handlerReturn *Value,
	data T.Gpointer) T.Gboolean

type SignalCMarshaller ClosureMarshal

type SignalEmissionHook func(
	ihint *SignalInvocationHint,
	nParamValues uint,
	paramValues *Value,
	data T.Gpointer) T.Gboolean

type SignalFlags Enum

const (
	SIGNAL_RUN_FIRST SignalFlags = 1 << iota
	SIGNAL_RUN_LAST
	SIGNAL_RUN_CLEANUP
	SIGNAL_NO_RECURSE
	SIGNAL_DETAILED
	SIGNAL_ACTION
	SIGNAL_NO_HOOKS
)

type SignalInvocationHint struct {
	SignalId uint
	Detail   T.Quark
	RunType  SignalFlags
}

type SignalMatchType Enum

const (
	SIGNAL_MATCH_ID SignalMatchType = 1 << iota
	SIGNAL_MATCH_DETAIL
	SIGNAL_MATCH_CLOSURE
	SIGNAL_MATCH_FUNC
	SIGNAL_MATCH_DATA
	SIGNAL_MATCH_UNBLOCKED
)

type SignalQueryFunc struct {
	SignalId    uint
	SignalName  string
	Itype       T.GType
	SignalFlags SignalFlags
	ReturnType  T.GType
	NParams     uint
	ParamTypes  *T.GType
}

type Source struct {
	CallbackData  T.Gpointer
	CallbackFuncs *SourceCallbackFuncs
	Source_funcs  *SourceFuncs
	RefCount      uint
	Context       *T.MainContext
	Priority      int
	Flags         uint
	SourceId      uint
	PollFds       *SList
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

type SourceFunc func(data T.Gpointer) bool

type SourceFuncs struct {
	Prepare  func(source *Source, timeout *int) bool
	Check    func(source *Source) bool
	Dispatch func(
		source *Source, callback SourceFunc, userData T.Gpointer) bool
	Finalize func(source *Source)

	ClosureCallback SourceFunc
	ClosureMarshal  SourceDummyMarshal
}
