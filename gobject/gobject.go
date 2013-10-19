// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gobject provides API definitions for accessing
//libgobject-2.0-0.dll.
package gobject

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllData(dll, false, dataList)
}

var (
	SignalNewv func(
		signalName string,
		itype T.GType,
		signalFlags T.GSignalFlags,
		classClosure *T.GClosure,
		accumulator T.GSignalAccumulator,
		accuData T.Gpointer,
		cMarshaller T.GSignalCMarshaller,
		returnType T.GType,
		nParams uint,
		paramTypes *T.GType) uint

	SignalNewValist func(
		signalName string,
		itype T.GType,
		signalFlags T.GSignalFlags,
		classClosure *T.GClosure,
		accumulator T.GSignalAccumulator,
		accuData T.Gpointer,
		cMarshaller T.GSignalCMarshaller,
		returnType T.GType,
		nParams uint,
		args T.VaList) uint

	SignalNew func(signalName string, itype T.GType,
		signalFlags T.GSignalFlags, classOffset uint,
		accumulator T.GSignalAccumulator, accuData T.Gpointer,
		cMarshaller T.GSignalCMarshaller, returnType T.GType,
		nParams uint, v ...VArg) uint

	SignalNewClassHandler func(signalName string,
		itype T.GType, signalFlags T.GSignalFlags,
		classHandler T.GCallback, accumulator T.GSignalAccumulator,
		accuData T.Gpointer, cMarshaller T.GSignalCMarshaller,
		returnType T.GType, nParams uint, v ...VArg) uint

	SignalEmitv func(
		instanceAndParams *T.GValue,
		signalId uint,
		detail T.GQuark,
		returnValue *T.GValue)

	SignalEmitValist func(
		instance T.Gpointer,
		signalId uint,
		detail T.GQuark,
		varArgs T.VaList)

	SignalEmit func(instance T.Gpointer, signalId uint,
		detail T.GQuark, v ...VArg)

	SignalEmitByName func(instance T.Gpointer,
		detailedSignal string, v ...VArg)

	SignalLookup func(name string, itype T.GType) uint

	SignalName func(signalId uint) string

	SignalQuery func(signalId uint, query *T.GSignalQuery)

	SignalListIds func(itype T.GType, nIds *uint) *uint

	SignalParseName func(
		detailedSignal string,
		itype T.GType,
		signalIdP *uint,
		detailP *T.GQuark,
		forceDetailQuark T.Gboolean) T.Gboolean

	SignalGetInvocationHint func(
		instance T.Gpointer) *T.GSignalInvocationHint

	SignalStopEmission func(
		instance T.Gpointer,
		signalId uint,
		detail T.GQuark)

	SignalStopEmissionByName func(
		instance T.Gpointer,
		detailedSignal string)

	SignalAddEmissionHook func(
		signalId uint,
		detail T.GQuark,
		hookFunc T.GSignalEmissionHook,
		hookData T.Gpointer,
		dataDestroy T.GDestroyNotify) T.Gulong

	SignalRemoveEmissionHook func(
		signalId uint,
		hookId T.Gulong)

	SignalHasHandlerPending func(
		instance T.Gpointer,
		signalId uint,
		detail T.GQuark,
		mayBeBlocked T.Gboolean) T.Gboolean

	SignalConnectClosureById func(
		instance T.Gpointer,
		signalId uint,
		detail T.GQuark,
		closure *T.GClosure,
		after T.Gboolean) T.Gulong

	SignalConnectClosure func(
		instance T.Gpointer,
		detailedSignal string,
		closure *T.GClosure,
		after T.Gboolean) T.Gulong

	SignalConnectData func(
		instance T.Gpointer,
		detailedSignal string,
		cHandler T.GCallback,
		data T.Gpointer,
		destroyData T.GClosureNotify,
		connectFlags T.GConnectFlags) T.Gulong

	SignalHandlerBlock func(
		instance T.Gpointer,
		handlerId T.Gulong)

	SignalHandlerUnblock func(
		instance T.Gpointer,
		handlerId T.Gulong)

	SignalHandlerDisconnect func(
		instance T.Gpointer,
		handlerId T.Gulong)

	SignalHandlerIsConnected func(
		instance T.Gpointer,
		handlerId T.Gulong) T.Gboolean

	SignalHandlerFind func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signalId uint,
		detail T.GQuark,
		closure *T.GClosure,
		fnc T.Gpointer,
		data T.Gpointer) T.Gulong

	SignalHandlersBlockMatched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signalId uint,
		detail T.GQuark,
		closure *T.GClosure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalHandlersUnblockMatched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signalId uint,
		detail T.GQuark,
		closure *T.GClosure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalHandlersDisconnectMatched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signalId uint,
		detail T.GQuark,
		closure *T.GClosure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalOverrideClassClosure func(
		signalId uint,
		instanceType T.GType,
		classClosure *T.GClosure)

	SignalOverrideClassHandler func(
		signalName string,
		instanceType T.GType,
		classHandler T.GCallback)

	SignalChainFromOverridden func(
		instanceAndParams *T.GValue,
		returnValue *T.GValue)

	SignalChainFromOverriddenHandler func(
		instance T.Gpointer, v ...VArg)

	SignalAccumulatorTrueHandled func(
		ihint *T.GSignalInvocationHint,
		returnAccu *T.GValue,
		handlerReturn *T.GValue,
		dummy T.Gpointer) T.Gboolean

	SignalAccumulatorFirstWins func(
		ihint *T.GSignalInvocationHint,
		returnAccu *T.GValue,
		handlerReturn *T.GValue,
		dummy T.Gpointer) T.Gboolean

	SignalHandlersDestroy func(
		instance T.Gpointer)

	InitiallyUnownedGetType func() T.GType

	ObjectClassInstallProperty func(
		oclass *T.GObjectClass,
		propertyId uint,
		pspec *T.GParamSpec)

	ObjectClassFindProperty func(
		oclass *T.GObjectClass,
		propertyName string) *T.GParamSpec

	ObjectClassListProperties func(
		oclass *T.GObjectClass,
		nProperties *uint) **T.GParamSpec

	ObjectClassOverrideProperty func(
		oclass *T.GObjectClass,
		propertyId uint,
		name string)

	ObjectClassInstallProperties func(
		oclass *T.GObjectClass,
		nPspecs uint,
		pspecs **T.GParamSpec)

	ObjectInterfaceInstallProperty func(
		gIface T.Gpointer,
		pspec *T.GParamSpec)

	ObjectInterfaceFindProperty func(
		gIface T.Gpointer,
		propertyName string) *T.GParamSpec

	ObjectInterfaceListProperties func(
		gIface T.Gpointer,
		nPropertiesP *uint) **T.GParamSpec

	ObjectGetType func() T.GType

	ObjectNew func(objectType T.GType,
		firstPropertyName string, v ...VArg) T.Gpointer

	ObjectNewv func(
		objectType T.GType,
		nParameters uint,
		parameters *T.GParameter) T.Gpointer

	ObjectNewValist func(
		objectType T.GType,
		firstPropertyName string,
		varArgs T.VaList) *T.GObject

	ObjectSet func(object T.Gpointer,
		firstPropertyName string, v ...VArg)

	ObjectGet func(object T.Gpointer,
		firstPropertyName string, v ...VArg)

	ObjectConnect func(object T.Gpointer,
		signalSpec string, v ...VArg) T.Gpointer

	ObjectDisconnect func(object T.Gpointer,
		signalSpec string, v ...VArg)

	ObjectSetValist func(
		object *T.GObject,
		firstPropertyName string,
		varArgs T.VaList)

	ObjectGetValist func(
		object *T.GObject,
		firstPropertyName string,
		varArgs T.VaList)

	ObjectSetProperty func(
		object *T.GObject,
		propertyName string,
		value *T.GValue)

	ObjectGetProperty func(
		object *T.GObject,
		propertyName string,
		value *T.GValue)

	ObjectFreezeNotify func(
		object *T.GObject)

	ObjectNotify func(
		object *T.GObject,
		propertyName string)

	ObjectNotifyByPspec func(
		object *T.GObject,
		pspec *T.GParamSpec)

	ObjectThawNotify func(
		object *T.GObject)

	ObjectIsFloating func(
		object T.Gpointer) T.Gboolean

	ObjectRefSink func(
		object T.Gpointer) T.Gpointer

	ObjectRef func(
		object T.Gpointer) T.Gpointer

	ObjectUnref func(
		object T.Gpointer)

	ObjectWeakRef func(
		object *T.GObject,
		notify T.GWeakNotify,
		data T.Gpointer)

	ObjectWeakUnref func(
		object *T.GObject,
		notify T.GWeakNotify,
		data T.Gpointer)

	ObjectAddWeakPointer func(
		object *T.GObject,
		weakPointerLocation *T.Gpointer)

	ObjectRemoveWeakPointer func(
		object *T.GObject,
		weakPointerLocation *T.Gpointer)

	ObjectAddToggleRef func(
		object *T.GObject,
		notify T.GToggleNotify,
		data T.Gpointer)

	ObjectRemoveToggleRef func(
		object *T.GObject,
		notify T.GToggleNotify,
		data T.Gpointer)

	ObjectGetQdata func(
		object *T.GObject,
		quark T.GQuark) T.Gpointer

	ObjectSetQdata func(
		object *T.GObject,
		quark T.GQuark,
		data T.Gpointer)

	ObjectSetQdataFull func(
		object *T.GObject,
		quark T.GQuark,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	ObjectStealQdata func(
		object *T.GObject,
		quark T.GQuark) T.Gpointer

	ObjectGetData func(
		object *T.GObject,
		key string) T.Gpointer

	ObjectSetData func(
		object *T.GObject,
		key string,
		data T.Gpointer)

	ObjectSetDataFull func(
		object *T.GObject,
		key string,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	ObjectStealData func(
		object *T.GObject,
		key string) T.Gpointer

	ObjectWatchClosure func(
		object *T.GObject,
		closure *T.GClosure)

	CclosureNewObject func(
		callbackFunc T.GCallback,
		object *T.GObject) *T.GClosure

	CclosureNewObjectSwap func(
		callbackFunc T.GCallback,
		object *T.GObject) *T.GClosure

	ClosureNewObject func(
		sizeofClosure uint,
		object *T.GObject) *T.GClosure

	ValueSetObject func(
		value *T.GValue,
		vObject T.Gpointer)

	ValueGetObject func(
		value *T.GValue) T.Gpointer

	ValueDupObject func(
		value *T.GValue) T.Gpointer

	SignalConnectObject func(
		instance T.Gpointer,
		detailedSignal string,
		cHandler T.GCallback,
		gobject T.Gpointer,
		connectFlags T.GConnectFlags) T.Gulong

	ObjectForceFloating func(
		object *T.GObject)

	ObjectRunDispose func(
		object *T.GObject)

	ValueTakeObject func(
		value *T.GValue,
		vObject T.Gpointer)

	ValueSetObjectTakeOwnership func(
		value *T.GValue,
		vObject T.Gpointer)

	ObjectCompatControl func(
		what T.Gsize,
		data T.Gpointer) T.Gsize

	ClearObject func(objectPtr **T.GObject)

	BindingFlagsGetType func() T.GType

	BindingGetType func() T.GType

	BindingGetFlags func(binding *T.GBinding) T.GBindingFlags

	BindingGetSource func(binding *T.GBinding) *T.GObject

	BindingGetTarget func(binding *T.GBinding) *T.GObject

	BindingGetSourceProperty func(binding *T.GBinding) string

	BindingGetTargetProperty func(binding *T.GBinding) string

	ObjectBindProperty func(
		source T.Gpointer,
		sourceProperty string,
		target T.Gpointer,
		targetProperty string,
		flags T.GBindingFlags) *T.GBinding

	ObjectBindPropertyFull func(
		source T.Gpointer,
		sourceProperty string,
		target T.Gpointer,
		targetProperty string,
		flags T.GBindingFlags,
		transformTo T.GBindingTransformFunc,
		transformFrom T.GBindingTransformFunc,
		userData T.Gpointer,
		notify T.GDestroyNotify) *T.GBinding

	ObjectBindPropertyWithClosures func(
		source T.Gpointer,
		sourceProperty string,
		target T.Gpointer,
		targetProperty string,
		flags T.GBindingFlags,
		transformTo *T.GClosure,
		transformFrom *T.GClosure) *T.GBinding

	BoxedCopy func(
		boxedType T.GType,
		srcBoxed T.Gconstpointer) T.Gpointer

	BoxedFree func(
		boxedType T.GType,
		boxed T.Gpointer)

	ValueSetBoxed func(
		value *T.GValue,
		vBoxed T.Gconstpointer)

	ValueSetStaticBoxed func(
		value *T.GValue,
		vBoxed T.Gconstpointer)

	ValueGetBoxed func(
		value *T.GValue) T.Gpointer

	ValueDupBoxed func(
		value *T.GValue) T.Gpointer

	BoxedTypeRegisterStatic func(
		name string,
		boxedCopy T.GBoxedCopyFunc,
		boxedFree T.GBoxedFreeFunc) T.GType

	ValueTakeBoxed func(
		value *T.GValue,
		vBoxed T.Gconstpointer)

	ValueSetBoxedTakeOwnership func(
		value *T.GValue,
		vBoxed T.Gconstpointer)

	ClosureGetType func() T.GType

	ValueGetType func() T.GType

	ValueArrayGetType func() T.GType

	DateGetType func() T.GType

	StrvGetType func() T.GType

	GstringGetType func() T.GType

	HashTableGetType func() T.GType

	ArrayGetType func() T.GType

	ByteArrayGetType func() T.GType

	PtrArrayGetType func() T.GType

	VariantTypeGetGtype func() T.GType

	RegexGetType func() T.GType

	ErrorGetType func() T.GType

	DateTimeGetType func() T.GType

	VariantGetGtype func() T.GType

	GTypeInit func()

	TypeInitWithDebugFlags func(
		debugFlags T.GTypeDebugFlags)

	TypeName func(t T.GType) string

	TypeQname func(t T.GType) T.GQuark

	TypeFromName func(name string) T.GType

	TypeParent func(t T.GType) T.GType

	TypeDepth func(t T.GType) uint

	TypeNextBase func(leafType, rootType T.GType) T.GType

	TypeIsA func(t, isAType T.GType) T.Gboolean

	TypeClassRef func(
		t T.GType) T.Gpointer

	TypeClassPeek func(
		t T.GType) T.Gpointer

	TypeClassPeekStatic func(
		t T.GType) T.Gpointer

	TypeClassUnref func(
		gClass T.Gpointer)

	TypeClassPeekParent func(
		gClass T.Gpointer) T.Gpointer

	TypeInterfacePeek func(
		instanceClass T.Gpointer,
		ifaceType T.GType) T.Gpointer

	TypeInterfacePeekParent func(
		gIface T.Gpointer) T.Gpointer

	TypeDefaultInterfaceRef func(
		gType T.GType) T.Gpointer

	TypeDefaultInterfacePeek func(
		gType T.GType) T.Gpointer

	TypeDefaultInterfaceUnref func(
		gIface T.Gpointer)

	TypeChildren func(
		t T.GType,
		nChildren *uint) *T.GType

	TypeInterfaces func(
		t T.GType,
		nInterfaces *uint) *T.GType

	TypeSetQdata func(
		t T.GType,
		quark T.GQuark,
		data T.Gpointer)

	TypeGetQdata func(
		t T.GType,
		quark T.GQuark) T.Gpointer

	TypeQuery func(
		t T.GType,
		query *T.GTypeQuery)

	CclosureNew func(
		callbackFunc T.GCallback,
		userData T.Gpointer,
		destroyData T.GClosureNotify) *T.GClosure

	CclosureNewSwap func(
		callbackFunc T.GCallback,
		userData T.Gpointer,
		destroyData T.GClosureNotify) *T.GClosure

	SignalTypeCclosureNew func(
		itype T.GType,
		structOffset uint) *T.GClosure

	ClosureRef func(
		closure *T.GClosure) *T.GClosure

	ClosureSink func(
		closure *T.GClosure)

	ClosureUnref func(
		closure *T.GClosure)

	ClosureNewSimple func(
		sizeofClosure uint,
		data T.Gpointer) *T.GClosure

	ClosureAddFinalizeNotifier func(
		closure *T.GClosure,
		notifyData T.Gpointer,
		notifyFunc T.GClosureNotify)

	ClosureRemoveFinalizeNotifier func(
		closure *T.GClosure,
		notifyData T.Gpointer,
		notifyFunc T.GClosureNotify)

	ClosureAddInvalidateNotifier func(
		closure *T.GClosure,
		notifyData T.Gpointer,
		notifyFunc T.GClosureNotify)

	ClosureRemoveInvalidateNotifier func(
		closure *T.GClosure,
		notifyData T.Gpointer,
		notifyFunc T.GClosureNotify)

	ClosureAddMarshalGuards func(
		closure *T.GClosure,
		preMarshalData T.Gpointer,
		preMarshalNotify T.GClosureNotify,
		postMarshalData T.Gpointer,
		postMarshalNotify T.GClosureNotify)

	ClosureSetMarshal func(
		closure *T.GClosure,
		marshal T.GClosureMarshal)

	ClosureSetMetaMarshal func(
		closure *T.GClosure,
		marshalData T.Gpointer,
		metaMarshal T.GClosureMarshal)

	ClosureInvalidate func(
		closure *T.GClosure)

	ClosureInvoke func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer)

	CclosureMarshal_VOID__VOID func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__BOOLEAN func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__CHAR func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__UCHAR func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__LONG func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__ULONG func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__ENUM func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__FLAGS func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__FLOAT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__DOUBLE func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__STRING func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__PARAM func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__BOXED func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__OBJECT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__VARIANT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__UINT_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_BOOLEAN__FLAGS func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_STRING__OBJECT_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_BOOLEAN__BOXED_BOXED func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	EnumGetValue func(
		enumClass *T.GEnumClass,
		value int) *T.GEnumValue

	EnumGetValueByName func(
		enumClass *T.GEnumClass,
		name string) *T.GEnumValue

	EnumGetValueByNick func(
		enumClass *T.GEnumClass,
		nick string) *T.GEnumValue

	FlagsGetFirstValue func(
		flagsClass *T.GFlagsClass,
		value uint) *T.GFlagsValue

	FlagsGetValueByName func(
		flagsClass *T.GFlagsClass,
		name string) *T.GFlagsValue

	FlagsGetValueByNick func(
		flagsClass *T.GFlagsClass,
		nick string) *T.GFlagsValue

	ValueSetEnum func(
		value *T.GValue,
		vEnum int)

	ValueGetEnum func(
		value *T.GValue) int

	ValueSetFlags func(
		value *T.GValue,
		vFlags uint)

	ValueGetFlags func(
		value *T.GValue) uint

	EnumRegisterStatic func(
		name string,
		constStaticValues *T.GEnumValue) T.GType

	FlagsRegisterStatic func(
		name string,
		constStaticValues *T.GFlagsValue) T.GType

	EnumCompleteTypeInfo func(
		gEnumType T.GType,
		info *T.GTypeInfo,
		constValues *T.GEnumValue)

	FlagsCompleteTypeInfo func(
		gFlagsType T.GType,
		info *T.GTypeInfo,
		constValues *T.GFlagsValue)

	ParamSpecRef func(
		pspec *T.GParamSpec) *T.GParamSpec

	ParamSpecUnref func(
		pspec *T.GParamSpec)

	ParamSpecSink func(
		pspec *T.GParamSpec)

	ParamSpecRefSink func(
		pspec *T.GParamSpec) *T.GParamSpec

	ParamSpecGetQdata func(
		pspec *T.GParamSpec,
		quark T.GQuark) T.Gpointer

	ParamSpecSetQdata func(
		pspec *T.GParamSpec,
		quark T.GQuark,
		data T.Gpointer)

	ParamSpecSetQdataFull func(
		pspec *T.GParamSpec,
		quark T.GQuark,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	ParamSpecStealQdata func(
		pspec *T.GParamSpec,
		quark T.GQuark) T.Gpointer

	ParamSpecGetRedirectTarget func(
		pspec *T.GParamSpec) *T.GParamSpec

	ParamValueSetDefault func(
		pspec *T.GParamSpec,
		value *T.GValue)

	ParamValueDefaults func(
		pspec *T.GParamSpec,
		value *T.GValue) T.Gboolean

	ParamValueValidate func(
		pspec *T.GParamSpec,
		value *T.GValue) T.Gboolean

	ParamValueConvert func(
		pspec *T.GParamSpec,
		srcValue *T.GValue,
		destValue *T.GValue,
		strictValidation T.Gboolean) T.Gboolean

	ParamValuesCmp func(
		pspec *T.GParamSpec,
		value1 *T.GValue,
		value2 *T.GValue) int

	ParamSpecGetName func(
		pspec *T.GParamSpec) string

	ParamSpecGetNick func(
		pspec *T.GParamSpec) string

	ParamSpecGetBlurb func(
		pspec *T.GParamSpec) string

	ValueSetParam func(
		value *T.GValue,
		param *T.GParamSpec)

	ValueGetParam func(
		value *T.GValue) *T.GParamSpec

	ValueDupParam func(
		value *T.GValue) *T.GParamSpec

	ValueTakeParam func(
		value *T.GValue,
		param *T.GParamSpec)

	ValueSetParamTakeOwnership func(
		value *T.GValue,
		param *T.GParamSpec)

	ValueArrayGetNth func(
		valueArray *T.GValueArray,
		index uint) *T.GValue

	ValueArrayNew func(nPrealloced uint) *T.GValueArray

	ValueArrayFree func(valueArray *T.GValueArray)

	ValueArrayCopy func(
		valueArray *T.GValueArray) *T.GValueArray

	ValueArrayPrepend func(
		valueArray *T.GValueArray,
		value *T.GValue) *T.GValueArray

	ValueArrayAppend func(
		valueArray *T.GValueArray,
		value *T.GValue) *T.GValueArray

	ValueArrayInsert func(
		valueArray *T.GValueArray,
		index uint,
		value *T.GValue) *T.GValueArray

	ValueArrayRemove func(
		valueArray *T.GValueArray,
		index uint) *T.GValueArray

	ValueArraySort func(
		valueArray *T.GValueArray,
		compareFunc T.GCompareFunc) *T.GValueArray

	ValueArraySortWithData func(
		valueArray *T.GValueArray,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer) *T.GValueArray

	ValueSetChar func(value *T.GValue, vChar T.Gchar)

	ValueGetChar func(value *T.GValue) T.Gchar

	ValueSetUchar func(
		value *T.GValue, vUchar T.Guchar)

	ValueGetUchar func(value *T.GValue) T.Guchar

	ValueSetBoolean func(
		value *T.GValue, vBoolean T.Gboolean)

	ValueGetBoolean func(value *T.GValue) T.Gboolean

	ValueSetInt func(value *T.GValue, vInt int)

	ValueGetInt func(value *T.GValue) int

	ValueSetUint func(value *T.GValue, vUint uint)

	ValueGetUint func(value *T.GValue) uint

	ValueSetLong func(value *T.GValue, vLong T.Glong)

	ValueGetLong func(value *T.GValue) T.Glong

	ValueSetUlong func(value *T.GValue, vUlong T.Gulong)

	ValueGetUlong func(value *T.GValue) T.Gulong

	ValueSetInt64 func(value *T.GValue, vInt64 int64)

	ValueGetInt64 func(value *T.GValue) int64

	ValueSetUint64 func(value *T.GValue, vUint64 uint64)

	ValueGetUint64 func(value *T.GValue) uint64

	ValueSetFloat func(value *T.GValue, vFloat float32)

	ValueGetFloat func(value *T.GValue) float32

	ValueSetDouble func(value *T.GValue, vDouble float64)

	ValueGetDouble func(value *T.GValue) float64

	ValueSetString func(value *T.GValue, vString string)

	ValueSetStaticString func(
		value *T.GValue, vString string)

	ValueGetString func(value *T.GValue) string

	ValueDupString func(value *T.GValue) string

	ValueSetPointer func(
		value *T.GValue, vPointer T.Gpointer)

	ValueGetPointer func(value *T.GValue) T.Gpointer

	GtypeGetType func() T.GType

	ValueSetGtype func(value *T.GValue, vGtype T.GType)

	ValueGetGtype func(value *T.GValue) T.GType

	ValueSetVariant func(
		value *T.GValue, variant *T.GVariant)

	ValueTakeVariant func(
		value *T.GValue, variant *T.GVariant)

	ValueGetVariant func(value *T.GValue) *T.GVariant

	ValueDupVariant func(value *T.GValue) *T.GVariant

	PointerTypeRegisterStatic func(name string) T.GType

	StrdupValueContents func(value *T.GValue) string

	ValueTakeString func(
		value *T.GValue, vString string)

	ValueSetStringTakeOwnership func(
		value *T.GValue, vString string)

	TypeRegisterStatic func(
		parentType T.GType,
		typeName string,
		info *T.GTypeInfo,
		flags T.GTypeFlags) T.GType

	TypeRegisterStaticSimple func(
		parentType T.GType,
		typeName string,
		classSize uint,
		classInit T.GClassInitFunc,
		instanceSize uint,
		instanceInit T.GInstanceInitFunc,
		flags T.GTypeFlags) T.GType

	TypeRegisterDynamic func(
		parentType T.GType,
		typeName string,
		plugin *T.GTypePlugin,
		flags T.GTypeFlags) T.GType

	TypeRegisterFundamental func(
		typeId T.GType,
		typeName string,
		info *T.GTypeInfo,
		finfo *T.GTypeFundamentalInfo,
		flags T.GTypeFlags) T.GType

	TypeAddInterfaceStatic func(
		instanceType T.GType,
		interfaceType T.GType,
		info *T.GInterfaceInfo)

	TypeAddInterfaceDynamic func(
		instanceType T.GType,
		interfaceType T.GType,
		plugin *T.GTypePlugin)

	TypeInterfaceAddPrerequisite func(
		interfaceType T.GType,
		prerequisiteType T.GType)

	TypeInterfacePrerequisites func(
		interfaceType T.GType,
		nPrerequisites *uint) *T.GType

	TypeClassAddPrivate func(
		gClass T.Gpointer,
		privateSize T.Gsize)

	TypeInstanceGetPrivate func(
		instance *T.GTypeInstance,
		privateType T.GType) T.Gpointer

	TypeAddClassPrivate func(
		classType T.GType,
		privateSize T.Gsize)

	TypeClassGetPrivate func(
		klass *T.GTypeClass,
		privateType T.GType) T.Gpointer

	TypeGetPlugin func(
		t T.GType) *T.GTypePlugin

	TypeInterfaceGetPlugin func(
		instanceType T.GType,
		interfaceType T.GType) *T.GTypePlugin

	TypeFundamentalNext func() T.GType

	TypeFundamental func(
		typeId T.GType) T.GType

	TypeCreateInstance func(
		t T.GType) *T.GTypeInstance

	TypeFreeInstance func(
		instance *T.GTypeInstance)

	TypeAddClassCacheFunc func(
		cacheData T.Gpointer,
		cacheFunc T.GTypeClassCacheFunc)

	TypeRemoveClassCacheFunc func(
		cacheData T.Gpointer,
		cacheFunc T.GTypeClassCacheFunc)

	TypeClassUnrefUncached func(
		gClass T.Gpointer)

	TypeAddInterfaceCheck func(
		checkData T.Gpointer,
		checkFunc T.GTypeInterfaceCheckFunc)

	TypeRemoveInterfaceCheck func(
		checkData T.Gpointer,
		checkFunc T.GTypeInterfaceCheckFunc)

	TypeValueTablePeek func(
		t T.GType) *T.GTypeValueTable

	TypeCheckInstance func(
		instance *T.GTypeInstance) T.Gboolean

	TypeCheckInstanceCast func(
		instance *T.GTypeInstance,
		ifaceType T.GType) *T.GTypeInstance

	TypeCheckInstanceIsA func(
		instance *T.GTypeInstance,
		ifaceType T.GType) T.Gboolean

	TypeCheckClassCast func(
		gClass *T.GTypeClass,
		isAType T.GType) *T.GTypeClass

	TypeCheckClassIsA func(
		gClass *T.GTypeClass,
		isAType T.GType) T.Gboolean

	TypeCheckIsValueType func(
		t T.GType) T.Gboolean

	TypeCheckValue func(
		value *T.GValue) T.Gboolean

	TypeCheckValueHolds func(
		value *T.GValue,
		t T.GType) T.Gboolean

	TypeTestFlags func(
		t T.GType,
		flags uint) T.Gboolean

	TypeNameFromInstance func(
		instance *T.GTypeInstance) string

	TypeNameFromClass func(
		gClass *T.GTypeClass) string

	ValueCInit func()

	ValueTypesInit func()

	EnumTypesInit func()

	ParamTypeInit func()

	BoxedTypeInit func()

	ObjectTypeInit func()

	ParamSpecTypesInit func()

	ValueTransformsInit func()

	SignalInit func()

	ValueInit func(
		value *T.GValue,
		gType T.GType) *T.GValue

	ValueCopy func(
		srcValue *T.GValue,
		destValue *T.GValue)

	ValueReset func(
		value *T.GValue) *T.GValue

	ValueUnset func(
		value *T.GValue)

	ValueSetInstance func(
		value *T.GValue,
		instance T.Gpointer)

	ValueFitsPointer func(
		value *T.GValue) T.Gboolean

	ValuePeekPointer func(
		value *T.GValue) T.Gpointer

	ValueTypeCompatible func(
		srcType T.GType,
		destType T.GType) T.Gboolean

	ValueTypeTransformable func(
		srcType T.GType,
		destType T.GType) T.Gboolean

	ValueTransform func(
		srcValue *T.GValue,
		destValue *T.GValue) T.Gboolean

	ValueRegisterTransformFunc func(
		srcType T.GType,
		destType T.GType,
		transformFunc T.GValueTransform)

	TypePluginGetType func() T.GType

	TypePluginUse func(
		plugin *T.GTypePlugin)

	TypePluginUnuse func(
		plugin *T.GTypePlugin)

	TypePluginCompleteTypeInfo func(
		plugin *T.GTypePlugin,
		gType T.GType,
		info *T.GTypeInfo,
		valueTable *T.GTypeValueTable)

	TypePluginCompleteInterfaceInfo func(
		plugin *T.GTypePlugin,
		instanceType T.GType,
		interfaceType T.GType,
		info *T.GInterfaceInfo)

	TypeModuleGetType func() T.GType

	TypeModuleUse func(
		module *T.GTypeModule) T.Gboolean

	TypeModuleUnuse func(
		module *T.GTypeModule)

	TypeModuleSetName func(
		module *T.GTypeModule,
		name string)

	TypeModuleRegisterType func(
		module *T.GTypeModule,
		parentType T.GType,
		typeName string,
		typeInfo *T.GTypeInfo,
		flags T.GTypeFlags) T.GType

	TypeModuleAddInterface func(
		module *T.GTypeModule,
		instanceType T.GType,
		interfaceType T.GType,
		interfaceInfo *T.GInterfaceInfo)

	TypeModuleRegisterEnum func(
		module *T.GTypeModule,
		name string,
		constStaticValues *T.GEnumValue) T.GType

	TypeModuleRegisterFlags func(
		module *T.GTypeModule,
		name string,
		constStaticValues *T.GFlagsValue) T.GType

	ParamSpecChar func(
		name string,
		nick string,
		blurb string,
		minimum int8,
		maximum int8,
		defaultValue int8,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecUchar func(
		name string,
		nick string,
		blurb string,
		minimum uint8,
		maximum uint8,
		defaultValue uint8,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecBoolean func(
		name string,
		nick string,
		blurb string,
		defaultValue T.Gboolean,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecInt func(
		name string,
		nick string,
		blurb string,
		minimum int,
		maximum int,
		defaultValue int,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecUint func(
		name string,
		nick string,
		blurb string,
		minimum uint,
		maximum uint,
		defaultValue uint,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecLong func(
		name string,
		nick string,
		blurb string,
		minimum T.Glong,
		maximum T.Glong,
		defaultValue T.Glong,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecUlong func(
		name string,
		nick string,
		blurb string,
		minimum T.Gulong,
		maximum T.Gulong,
		defaultValue T.Gulong,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecInt64 func(
		name string,
		nick string,
		blurb string,
		minimum int64,
		maximum int64,
		defaultValue int64,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecUint64 func(
		name string,
		nick string,
		blurb string,
		minimum uint64,
		maximum uint64,
		defaultValue uint64,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecUnichar func(
		name string,
		nick string,
		blurb string,
		defaultValue T.Gunichar,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecEnum func(
		name string,
		nick string,
		blurb string,
		enumType T.GType,
		defaultValue int,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecFlags func(
		name string,
		nick string,
		blurb string,
		flagsType T.GType,
		defaultValue uint,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecFloat func(
		name string,
		nick string,
		blurb string,
		minimum float32,
		maximum float32,
		defaultValue float32,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecDouble func(
		name string,
		nick string,
		blurb string,
		minimum float64,
		maximum float64,
		defaultValue float64,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecString func(
		name string,
		nick string,
		blurb string,
		defaultValue string,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecParam func(
		name string,
		nick string,
		blurb string,
		paramType T.GType,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecBoxed func(
		name string,
		nick string,
		blurb string,
		boxedType T.GType,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecPointer func(
		name string,
		nick string,
		blurb string,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecValueArray func(
		name string,
		nick string,
		blurb string,
		elementSpec *T.GParamSpec,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecObject func(
		name string,
		nick string,
		blurb string,
		objectType T.GType,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecOverride func(
		name string,
		overridden *T.GParamSpec) *T.GParamSpec

	ParamSpecGtype func(
		name string,
		nick string,
		blurb string,
		isAType T.GType,
		flags T.GParamFlags) *T.GParamSpec

	ParamSpecVariant func(
		name string,
		nick string,
		blurb string,
		t *T.GVariantType,
		defaultValue *T.GVariant,
		flags T.GParamFlags) *T.GParamSpec

	SourceSetClosure func(
		source *T.GSource,
		closure *T.GClosure)

	SourceSetDummyCallback func(
		source *T.GSource)

	IoChannelGetType func() T.GType

	IoConditionGetType func() T.GType

	ParamTypeRegisterStatic func(
		name string,
		pspecInfo *T.GParamSpecTypeInfo) T.GType

	ParamSpecInternal func(
		paramType T.GType,
		name string,
		nick string,
		blurb string,
		flags T.GParamFlags) T.Gpointer

	ParamSpecPoolNew func(
		typePrefixing T.Gboolean) *T.GParamSpecPool

	ParamSpecPoolInsert func(
		pool *T.GParamSpecPool,
		pspec *T.GParamSpec,
		ownerType T.GType)

	ParamSpecPoolRemove func(
		pool *T.GParamSpecPool,
		pspec *T.GParamSpec)

	ParamSpecPoolLookup func(
		pool *T.GParamSpecPool,
		paramName string,
		ownerType T.GType,
		walkAncestors T.Gboolean) *T.GParamSpec

	ParamSpecPoolListOwned func(
		pool *T.GParamSpecPool,
		ownerType T.GType) *T.GList

	ParamSpecPoolList func(
		pool *T.GParamSpecPool,
		ownerType T.GType,
		nPspecsP *uint) **T.GParamSpec

	UnicharValidate func(
		ch T.Gunichar) T.Gboolean

	SlistRemoveAll func(
		list *T.GSList,
		data T.Gconstpointer) *T.GSList
)

var dll = "libgobject-2.0-0.dll"

var apiList = outside.Apis{
	{"g_array_get_type", &ArrayGetType},
	{"g_binding_flags_get_type", &BindingFlagsGetType},
	{"g_binding_get_flags", &BindingGetFlags},
	{"g_binding_get_source", &BindingGetSource},
	{"g_binding_get_source_property", &BindingGetSourceProperty},
	{"g_binding_get_target", &BindingGetTarget},
	{"g_binding_get_target_property", &BindingGetTargetProperty},
	{"g_binding_get_type", &BindingGetType},
	{"g_boxed_copy", &BoxedCopy},
	{"g_boxed_free", &BoxedFree},
	{"g_boxed_type_register_static", &BoxedTypeRegisterStatic},
	{"g_byte_array_get_type", &ByteArrayGetType},
	{"g_cclosure_marshal_BOOLEAN__BOXED_BOXED", &CclosureMarshal_BOOLEAN__BOXED_BOXED},
	{"g_cclosure_marshal_BOOLEAN__FLAGS", &CclosureMarshal_BOOLEAN__FLAGS},
	{"g_cclosure_marshal_STRING__OBJECT_POINTER", &CclosureMarshal_STRING__OBJECT_POINTER},
	{"g_cclosure_marshal_VOID__BOOLEAN", &CclosureMarshal_VOID__BOOLEAN},
	{"g_cclosure_marshal_VOID__BOXED", &CclosureMarshal_VOID__BOXED},
	{"g_cclosure_marshal_VOID__CHAR", &CclosureMarshal_VOID__CHAR},
	{"g_cclosure_marshal_VOID__DOUBLE", &CclosureMarshal_VOID__DOUBLE},
	{"g_cclosure_marshal_VOID__ENUM", &CclosureMarshal_VOID__ENUM},
	{"g_cclosure_marshal_VOID__FLAGS", &CclosureMarshal_VOID__FLAGS},
	{"g_cclosure_marshal_VOID__FLOAT", &CclosureMarshal_VOID__FLOAT},
	{"g_cclosure_marshal_VOID__INT", &CclosureMarshal_VOID__INT},
	{"g_cclosure_marshal_VOID__LONG", &CclosureMarshal_VOID__LONG},
	{"g_cclosure_marshal_VOID__OBJECT", &CclosureMarshal_VOID__OBJECT},
	{"g_cclosure_marshal_VOID__PARAM", &CclosureMarshal_VOID__PARAM},
	{"g_cclosure_marshal_VOID__POINTER", &CclosureMarshal_VOID__POINTER},
	{"g_cclosure_marshal_VOID__STRING", &CclosureMarshal_VOID__STRING},
	{"g_cclosure_marshal_VOID__UCHAR", &CclosureMarshal_VOID__UCHAR},
	{"g_cclosure_marshal_VOID__UINT", &CclosureMarshal_VOID__UINT},
	{"g_cclosure_marshal_VOID__UINT_POINTER", &CclosureMarshal_VOID__UINT_POINTER},
	{"g_cclosure_marshal_VOID__ULONG", &CclosureMarshal_VOID__ULONG},
	{"g_cclosure_marshal_VOID__VARIANT", &CclosureMarshal_VOID__VARIANT},
	{"g_cclosure_marshal_VOID__VOID", &CclosureMarshal_VOID__VOID},
	{"g_cclosure_new", &CclosureNew},
	{"g_cclosure_new_object", &CclosureNewObject},
	{"g_cclosure_new_object_swap", &CclosureNewObjectSwap},
	{"g_cclosure_new_swap", &CclosureNewSwap},
	{"g_clear_object", &ClearObject},
	{"g_closure_add_finalize_notifier", &ClosureAddFinalizeNotifier},
	{"g_closure_add_invalidate_notifier", &ClosureAddInvalidateNotifier},
	{"g_closure_add_marshal_guards", &ClosureAddMarshalGuards},
	{"g_closure_get_type", &ClosureGetType},
	{"g_closure_invalidate", &ClosureInvalidate},
	{"g_closure_invoke", &ClosureInvoke},
	{"g_closure_new_object", &ClosureNewObject},
	{"g_closure_new_simple", &ClosureNewSimple},
	{"g_closure_ref", &ClosureRef},
	{"g_closure_remove_finalize_notifier", &ClosureRemoveFinalizeNotifier},
	{"g_closure_remove_invalidate_notifier", &ClosureRemoveInvalidateNotifier},
	{"g_closure_set_marshal", &ClosureSetMarshal},
	{"g_closure_set_meta_marshal", &ClosureSetMetaMarshal},
	{"g_closure_sink", &ClosureSink},
	{"g_closure_unref", &ClosureUnref},
	{"g_date_get_type", &DateGetType},
	{"g_date_time_get_type", &DateTimeGetType},
	{"g_enum_complete_type_info", &EnumCompleteTypeInfo},
	{"g_enum_get_value", &EnumGetValue},
	{"g_enum_get_value_by_name", &EnumGetValueByName},
	{"g_enum_get_value_by_nick", &EnumGetValueByNick},
	{"g_enum_register_static", &EnumRegisterStatic},
	{"g_error_get_type", &ErrorGetType},
	{"g_flags_complete_type_info", &FlagsCompleteTypeInfo},
	{"g_flags_get_first_value", &FlagsGetFirstValue},
	{"g_flags_get_value_by_name", &FlagsGetValueByName},
	{"g_flags_get_value_by_nick", &FlagsGetValueByNick},
	{"g_flags_register_static", &FlagsRegisterStatic},
	{"g_gstring_get_type", &GstringGetType},
	{"g_gtype_get_type", &GtypeGetType},
	{"g_hash_table_get_type", &HashTableGetType},
	{"g_initially_unowned_get_type", &InitiallyUnownedGetType},
	{"g_io_channel_get_type", &IoChannelGetType},
	{"g_io_condition_get_type", &IoConditionGetType},
	{"g_object_add_toggle_ref", &ObjectAddToggleRef},
	{"g_object_add_weak_pointer", &ObjectAddWeakPointer},
	{"g_object_bind_property", &ObjectBindProperty},
	{"g_object_bind_property_full", &ObjectBindPropertyFull},
	{"g_object_bind_property_with_closures", &ObjectBindPropertyWithClosures},
	{"g_object_class_find_property", &ObjectClassFindProperty},
	{"g_object_class_install_properties", &ObjectClassInstallProperties},
	{"g_object_class_install_property", &ObjectClassInstallProperty},
	{"g_object_class_list_properties", &ObjectClassListProperties},
	{"g_object_class_override_property", &ObjectClassOverrideProperty},
	{"g_object_compat_control", &ObjectCompatControl},
	{"g_object_connect", &ObjectConnect},
	{"g_object_disconnect", &ObjectDisconnect},
	{"g_object_force_floating", &ObjectForceFloating},
	{"g_object_freeze_notify", &ObjectFreezeNotify},
	{"g_object_get", &ObjectGet},
	{"g_object_get_data", &ObjectGetData},
	{"g_object_get_property", &ObjectGetProperty},
	{"g_object_get_qdata", &ObjectGetQdata},
	{"g_object_get_type", &ObjectGetType},
	{"g_object_get_valist", &ObjectGetValist},
	{"g_object_interface_find_property", &ObjectInterfaceFindProperty},
	{"g_object_interface_install_property", &ObjectInterfaceInstallProperty},
	{"g_object_interface_list_properties", &ObjectInterfaceListProperties},
	{"g_object_is_floating", &ObjectIsFloating},
	{"g_object_new", &ObjectNew},
	{"g_object_new_valist", &ObjectNewValist},
	{"g_object_newv", &ObjectNewv},
	{"g_object_notify", &ObjectNotify},
	{"g_object_notify_by_pspec", &ObjectNotifyByPspec},
	{"g_object_ref", &ObjectRef},
	{"g_object_ref_sink", &ObjectRefSink},
	{"g_object_remove_toggle_ref", &ObjectRemoveToggleRef},
	{"g_object_remove_weak_pointer", &ObjectRemoveWeakPointer},
	{"g_object_run_dispose", &ObjectRunDispose},
	{"g_object_set", &ObjectSet},
	{"g_object_set_data", &ObjectSetData},
	{"g_object_set_data_full", &ObjectSetDataFull},
	{"g_object_set_property", &ObjectSetProperty},
	{"g_object_set_qdata", &ObjectSetQdata},
	{"g_object_set_qdata_full", &ObjectSetQdataFull},
	{"g_object_set_valist", &ObjectSetValist},
	{"g_object_steal_data", &ObjectStealData},
	{"g_object_steal_qdata", &ObjectStealQdata},
	{"g_object_thaw_notify", &ObjectThawNotify},
	{"g_object_unref", &ObjectUnref},
	{"g_object_watch_closure", &ObjectWatchClosure},
	{"g_object_weak_ref", &ObjectWeakRef},
	{"g_object_weak_unref", &ObjectWeakUnref},
	{"g_param_spec_boolean", &ParamSpecBoolean},
	{"g_param_spec_boxed", &ParamSpecBoxed},
	{"g_param_spec_char", &ParamSpecChar},
	{"g_param_spec_double", &ParamSpecDouble},
	{"g_param_spec_enum", &ParamSpecEnum},
	{"g_param_spec_flags", &ParamSpecFlags},
	{"g_param_spec_float", &ParamSpecFloat},
	{"g_param_spec_get_blurb", &ParamSpecGetBlurb},
	{"g_param_spec_get_name", &ParamSpecGetName},
	{"g_param_spec_get_nick", &ParamSpecGetNick},
	{"g_param_spec_get_qdata", &ParamSpecGetQdata},
	{"g_param_spec_get_redirect_target", &ParamSpecGetRedirectTarget},
	{"g_param_spec_gtype", &ParamSpecGtype},
	{"g_param_spec_int", &ParamSpecInt},
	{"g_param_spec_int64", &ParamSpecInt64},
	{"g_param_spec_internal", &ParamSpecInternal},
	{"g_param_spec_long", &ParamSpecLong},
	{"g_param_spec_object", &ParamSpecObject},
	{"g_param_spec_override", &ParamSpecOverride},
	{"g_param_spec_param", &ParamSpecParam},
	{"g_param_spec_pointer", &ParamSpecPointer},
	{"g_param_spec_pool_insert", &ParamSpecPoolInsert},
	{"g_param_spec_pool_list", &ParamSpecPoolList},
	{"g_param_spec_pool_list_owned", &ParamSpecPoolListOwned},
	{"g_param_spec_pool_lookup", &ParamSpecPoolLookup},
	{"g_param_spec_pool_new", &ParamSpecPoolNew},
	{"g_param_spec_pool_remove", &ParamSpecPoolRemove},
	{"g_param_spec_ref", &ParamSpecRef},
	{"g_param_spec_ref_sink", &ParamSpecRefSink},
	{"g_param_spec_set_qdata", &ParamSpecSetQdata},
	{"g_param_spec_set_qdata_full", &ParamSpecSetQdataFull},
	{"g_param_spec_sink", &ParamSpecSink},
	{"g_param_spec_steal_qdata", &ParamSpecStealQdata},
	{"g_param_spec_string", &ParamSpecString},
	{"g_param_spec_uchar", &ParamSpecUchar},
	{"g_param_spec_uint", &ParamSpecUint},
	{"g_param_spec_uint64", &ParamSpecUint64},
	{"g_param_spec_ulong", &ParamSpecUlong},
	{"g_param_spec_unichar", &ParamSpecUnichar},
	{"g_param_spec_unref", &ParamSpecUnref},
	{"g_param_spec_value_array", &ParamSpecValueArray},
	{"g_param_spec_variant", &ParamSpecVariant},
	{"g_param_type_register_static", &ParamTypeRegisterStatic},
	{"g_param_value_convert", &ParamValueConvert},
	{"g_param_value_defaults", &ParamValueDefaults},
	{"g_param_value_set_default", &ParamValueSetDefault},
	{"g_param_value_validate", &ParamValueValidate},
	{"g_param_values_cmp", &ParamValuesCmp},
	{"g_pointer_type_register_static", &PointerTypeRegisterStatic},
	{"g_ptr_array_get_type", &PtrArrayGetType},
	{"g_regex_get_type", &RegexGetType},
	{"g_signal_accumulator_first_wins", &SignalAccumulatorFirstWins},
	{"g_signal_accumulator_true_handled", &SignalAccumulatorTrueHandled},
	{"g_signal_add_emission_hook", &SignalAddEmissionHook},
	{"g_signal_chain_from_overridden", &SignalChainFromOverridden},
	{"g_signal_chain_from_overridden_handler", &SignalChainFromOverriddenHandler},
	{"g_signal_connect_closure", &SignalConnectClosure},
	{"g_signal_connect_closure_by_id", &SignalConnectClosureById},
	{"g_signal_connect_data", &SignalConnectData},
	{"g_signal_connect_object", &SignalConnectObject},
	{"g_signal_emit", &SignalEmit},
	{"g_signal_emit_by_name", &SignalEmitByName},
	{"g_signal_emit_valist", &SignalEmitValist},
	{"g_signal_emitv", &SignalEmitv},
	{"g_signal_get_invocation_hint", &SignalGetInvocationHint},
	{"g_signal_handler_block", &SignalHandlerBlock},
	{"g_signal_handler_disconnect", &SignalHandlerDisconnect},
	{"g_signal_handler_find", &SignalHandlerFind},
	{"g_signal_handler_is_connected", &SignalHandlerIsConnected},
	{"g_signal_handler_unblock", &SignalHandlerUnblock},
	{"g_signal_handlers_block_matched", &SignalHandlersBlockMatched},
	{"g_signal_handlers_destroy", &SignalHandlersDestroy},
	{"g_signal_handlers_disconnect_matched", &SignalHandlersDisconnectMatched},
	{"g_signal_handlers_unblock_matched", &SignalHandlersUnblockMatched},
	{"g_signal_has_handler_pending", &SignalHasHandlerPending},
	{"g_signal_list_ids", &SignalListIds},
	{"g_signal_lookup", &SignalLookup},
	{"g_signal_name", &SignalName},
	{"g_signal_new", &SignalNew},
	{"g_signal_new_class_handler", &SignalNewClassHandler},
	{"g_signal_new_valist", &SignalNewValist},
	{"g_signal_newv", &SignalNewv},
	{"g_signal_override_class_closure", &SignalOverrideClassClosure},
	{"g_signal_override_class_handler", &SignalOverrideClassHandler},
	{"g_signal_parse_name", &SignalParseName},
	{"g_signal_query", &SignalQuery},
	{"g_signal_remove_emission_hook", &SignalRemoveEmissionHook},
	{"g_signal_stop_emission", &SignalStopEmission},
	{"g_signal_stop_emission_by_name", &SignalStopEmissionByName},
	{"g_signal_type_cclosure_new", &SignalTypeCclosureNew},
	{"g_slist_remove_all", &SlistRemoveAll},
	{"g_source_set_closure", &SourceSetClosure},
	{"g_source_set_dummy_callback", &SourceSetDummyCallback},
	{"g_strdup_value_contents", &StrdupValueContents},
	{"g_strv_get_type", &StrvGetType},
	{"g_type_add_class_cache_func", &TypeAddClassCacheFunc},
	{"g_type_add_class_private", &TypeAddClassPrivate},
	{"g_type_add_interface_check", &TypeAddInterfaceCheck},
	{"g_type_add_interface_dynamic", &TypeAddInterfaceDynamic},
	{"g_type_add_interface_static", &TypeAddInterfaceStatic},
	{"g_type_check_class_cast", &TypeCheckClassCast},
	{"g_type_check_class_is_a", &TypeCheckClassIsA},
	{"g_type_check_instance", &TypeCheckInstance},
	{"g_type_check_instance_cast", &TypeCheckInstanceCast},
	{"g_type_check_instance_is_a", &TypeCheckInstanceIsA},
	{"g_type_check_is_value_type", &TypeCheckIsValueType},
	{"g_type_check_value", &TypeCheckValue},
	{"g_type_check_value_holds", &TypeCheckValueHolds},
	{"g_type_children", &TypeChildren},
	{"g_type_class_add_private", &TypeClassAddPrivate},
	{"g_type_class_get_private", &TypeClassGetPrivate},
	{"g_type_class_peek", &TypeClassPeek},
	{"g_type_class_peek_parent", &TypeClassPeekParent},
	{"g_type_class_peek_static", &TypeClassPeekStatic},
	{"g_type_class_ref", &TypeClassRef},
	{"g_type_class_unref", &TypeClassUnref},
	{"g_type_class_unref_uncached", &TypeClassUnrefUncached},
	{"g_type_create_instance", &TypeCreateInstance},
	{"g_type_default_interface_peek", &TypeDefaultInterfacePeek},
	{"g_type_default_interface_ref", &TypeDefaultInterfaceRef},
	{"g_type_default_interface_unref", &TypeDefaultInterfaceUnref},
	{"g_type_depth", &TypeDepth},
	{"g_type_free_instance", &TypeFreeInstance},
	{"g_type_from_name", &TypeFromName},
	{"g_type_fundamental", &TypeFundamental},
	{"g_type_fundamental_next", &TypeFundamentalNext},
	{"g_type_get_plugin", &TypeGetPlugin},
	{"g_type_get_qdata", &TypeGetQdata},
	{"g_type_init", &GTypeInit},
	{"g_type_init_with_debug_flags", &TypeInitWithDebugFlags},
	{"g_type_instance_get_private", &TypeInstanceGetPrivate},
	{"g_type_interface_add_prerequisite", &TypeInterfaceAddPrerequisite},
	{"g_type_interface_get_plugin", &TypeInterfaceGetPlugin},
	{"g_type_interface_peek", &TypeInterfacePeek},
	{"g_type_interface_peek_parent", &TypeInterfacePeekParent},
	{"g_type_interface_prerequisites", &TypeInterfacePrerequisites},
	{"g_type_interfaces", &TypeInterfaces},
	{"g_type_is_a", &TypeIsA},
	{"g_type_module_add_interface", &TypeModuleAddInterface},
	{"g_type_module_get_type", &TypeModuleGetType},
	{"g_type_module_register_enum", &TypeModuleRegisterEnum},
	{"g_type_module_register_flags", &TypeModuleRegisterFlags},
	{"g_type_module_register_type", &TypeModuleRegisterType},
	{"g_type_module_set_name", &TypeModuleSetName},
	{"g_type_module_unuse", &TypeModuleUnuse},
	{"g_type_module_use", &TypeModuleUse},
	{"g_type_name", &TypeName},
	{"g_type_name_from_class", &TypeNameFromClass},
	{"g_type_name_from_instance", &TypeNameFromInstance},
	{"g_type_next_base", &TypeNextBase},
	{"g_type_parent", &TypeParent},
	{"g_type_plugin_complete_interface_info", &TypePluginCompleteInterfaceInfo},
	{"g_type_plugin_complete_type_info", &TypePluginCompleteTypeInfo},
	{"g_type_plugin_get_type", &TypePluginGetType},
	{"g_type_plugin_unuse", &TypePluginUnuse},
	{"g_type_plugin_use", &TypePluginUse},
	{"g_type_qname", &TypeQname},
	{"g_type_query", &TypeQuery},
	{"g_type_register_dynamic", &TypeRegisterDynamic},
	{"g_type_register_fundamental", &TypeRegisterFundamental},
	{"g_type_register_static", &TypeRegisterStatic},
	{"g_type_register_static_simple", &TypeRegisterStaticSimple},
	{"g_type_remove_class_cache_func", &TypeRemoveClassCacheFunc},
	{"g_type_remove_interface_check", &TypeRemoveInterfaceCheck},
	{"g_type_set_qdata", &TypeSetQdata},
	{"g_type_test_flags", &TypeTestFlags},
	{"g_type_value_table_peek", &TypeValueTablePeek},
	{"g_unichar_validate", &UnicharValidate},
	{"g_value_array_append", &ValueArrayAppend},
	{"g_value_array_copy", &ValueArrayCopy},
	{"g_value_array_free", &ValueArrayFree},
	{"g_value_array_get_nth", &ValueArrayGetNth},
	{"g_value_array_get_type", &ValueArrayGetType},
	{"g_value_array_insert", &ValueArrayInsert},
	{"g_value_array_new", &ValueArrayNew},
	{"g_value_array_prepend", &ValueArrayPrepend},
	{"g_value_array_remove", &ValueArrayRemove},
	{"g_value_array_sort", &ValueArraySort},
	{"g_value_array_sort_with_data", &ValueArraySortWithData},
	{"g_value_copy", &ValueCopy},
	{"g_value_dup_boxed", &ValueDupBoxed},
	{"g_value_dup_object", &ValueDupObject},
	{"g_value_dup_param", &ValueDupParam},
	{"g_value_dup_string", &ValueDupString},
	{"g_value_dup_variant", &ValueDupVariant},
	{"g_value_fits_pointer", &ValueFitsPointer},
	{"g_value_get_boolean", &ValueGetBoolean},
	{"g_value_get_boxed", &ValueGetBoxed},
	{"g_value_get_char", &ValueGetChar},
	{"g_value_get_double", &ValueGetDouble},
	{"g_value_get_enum", &ValueGetEnum},
	{"g_value_get_flags", &ValueGetFlags},
	{"g_value_get_float", &ValueGetFloat},
	{"g_value_get_gtype", &ValueGetGtype},
	{"g_value_get_int", &ValueGetInt},
	{"g_value_get_int64", &ValueGetInt64},
	{"g_value_get_long", &ValueGetLong},
	{"g_value_get_object", &ValueGetObject},
	{"g_value_get_param", &ValueGetParam},
	{"g_value_get_pointer", &ValueGetPointer},
	{"g_value_get_string", &ValueGetString},
	{"g_value_get_type", &ValueGetType},
	{"g_value_get_uchar", &ValueGetUchar},
	{"g_value_get_uint", &ValueGetUint},
	{"g_value_get_uint64", &ValueGetUint64},
	{"g_value_get_ulong", &ValueGetUlong},
	{"g_value_get_variant", &ValueGetVariant},
	{"g_value_init", &ValueInit},
	{"g_value_peek_pointer", &ValuePeekPointer},
	{"g_value_register_transform_func", &ValueRegisterTransformFunc},
	{"g_value_reset", &ValueReset},
	{"g_value_set_boolean", &ValueSetBoolean},
	{"g_value_set_boxed", &ValueSetBoxed},
	{"g_value_set_boxed_take_ownership", &ValueSetBoxedTakeOwnership},
	{"g_value_set_char", &ValueSetChar},
	{"g_value_set_double", &ValueSetDouble},
	{"g_value_set_enum", &ValueSetEnum},
	{"g_value_set_flags", &ValueSetFlags},
	{"g_value_set_float", &ValueSetFloat},
	{"g_value_set_gtype", &ValueSetGtype},
	{"g_value_set_instance", &ValueSetInstance},
	{"g_value_set_int", &ValueSetInt},
	{"g_value_set_int64", &ValueSetInt64},
	{"g_value_set_long", &ValueSetLong},
	{"g_value_set_object", &ValueSetObject},
	{"g_value_set_object_take_ownership", &ValueSetObjectTakeOwnership},
	{"g_value_set_param", &ValueSetParam},
	{"g_value_set_param_take_ownership", &ValueSetParamTakeOwnership},
	{"g_value_set_pointer", &ValueSetPointer},
	{"g_value_set_static_boxed", &ValueSetStaticBoxed},
	{"g_value_set_static_string", &ValueSetStaticString},
	{"g_value_set_string", &ValueSetString},
	{"g_value_set_string_take_ownership", &ValueSetStringTakeOwnership},
	{"g_value_set_uchar", &ValueSetUchar},
	{"g_value_set_uint", &ValueSetUint},
	{"g_value_set_uint64", &ValueSetUint64},
	{"g_value_set_ulong", &ValueSetUlong},
	{"g_value_set_variant", &ValueSetVariant},
	{"g_value_take_boxed", &ValueTakeBoxed},
	{"g_value_take_object", &ValueTakeObject},
	{"g_value_take_param", &ValueTakeParam},
	{"g_value_take_string", &ValueTakeString},
	{"g_value_take_variant", &ValueTakeVariant},
	{"g_value_transform", &ValueTransform},
	{"g_value_type_compatible", &ValueTypeCompatible},
	{"g_value_type_transformable", &ValueTypeTransformable},
	{"g_value_unset", &ValueUnset},
	{"g_variant_get_gtype", &VariantGetGtype},
	{"g_variant_type_get_gtype", &VariantTypeGetGtype},
}

var dataList = outside.Data{
	{"g_param_spec_types", (*T.GType)(nil)},
}
