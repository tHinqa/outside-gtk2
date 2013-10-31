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

type Enum int

var (
	GTypeInit func()

	ArrayGetType            func() Type
	ByteArrayGetType        func() Type
	DateGetType             func() Type
	DateTimeGetType         func() Type
	ErrorGetType            func() Type
	GstringGetType          func() Type
	GtypeGetType            func() Type
	HashTableGetType        func() Type
	InitiallyUnownedGetType func() Type
	IoChannelGetType        func() Type
	IoConditionGetType      func() Type
	PtrArrayGetType         func() Type
	RegexGetType            func() Type
	StrvGetType             func() Type
	VariantGetGtype         func() Type
	VariantTypeGetGtype     func() Type

	SignalNewv func(
		signalName string,
		itype Type,
		signalFlags T.GSignalFlags,
		classClosure *Closure,
		accumulator T.GSignalAccumulator,
		accuData T.Gpointer,
		cMarshaller T.GSignalCMarshaller,
		returnType Type,
		nParams uint,
		paramTypes *Type) uint

	SignalNewValist func(
		signalName string,
		itype Type,
		signalFlags T.GSignalFlags,
		classClosure *Closure,
		accumulator T.GSignalAccumulator,
		accuData T.Gpointer,
		cMarshaller T.GSignalCMarshaller,
		returnType Type,
		nParams uint,
		args T.VaList) uint

	SignalNew func(signalName string, itype Type,
		signalFlags T.GSignalFlags, classOffset uint,
		accumulator T.GSignalAccumulator, accuData T.Gpointer,
		cMarshaller T.GSignalCMarshaller, returnType Type,
		nParams uint, v ...VArg) uint

	SignalNewClassHandler func(signalName string,
		itype Type, signalFlags T.GSignalFlags,
		classHandler T.GCallback, accumulator T.GSignalAccumulator,
		accuData T.Gpointer, cMarshaller T.GSignalCMarshaller,
		returnType Type, nParams uint, v ...VArg) uint

	SignalEmitv func(
		instanceAndParams *Value,
		signalId uint,
		detail T.GQuark,
		returnValue *Value)

	SignalEmitValist func(
		instance T.Gpointer,
		signalId uint,
		detail T.GQuark,
		varArgs T.VaList)

	SignalEmit func(instance T.Gpointer, signalId uint,
		detail T.GQuark, v ...VArg)

	SignalEmitByName func(instance T.Gpointer,
		detailedSignal string, v ...VArg)

	SignalLookup func(name string, itype Type) uint

	SignalName func(signalId uint) string

	SignalQuery func(signalId uint, query *T.GSignalQuery)

	SignalListIds func(itype Type, nIds *uint) *uint

	SignalParseName func(
		detailedSignal string,
		itype Type,
		signalIdP *uint,
		detailP *T.GQuark,
		forceDetailQuark bool) bool

	SignalGetInvocationHint func(
		instance T.Gpointer) *T.GSignalInvocationHint

	SignalStopEmission func(
		instance T.Gpointer, signalId uint, detail T.GQuark)

	SignalStopEmissionByName func(
		instance T.Gpointer, detailedSignal string)

	SignalAddEmissionHook func(
		signalId uint,
		detail T.GQuark,
		hookFunc T.GSignalEmissionHook,
		hookData T.Gpointer,
		dataDestroy T.GDestroyNotify) T.Gulong

	SignalRemoveEmissionHook func(signalId uint, hookId T.Gulong)

	SignalHasHandlerPending func(
		instance T.Gpointer,
		signalId uint,
		detail T.GQuark,
		mayBeBlocked bool) bool

	SignalConnectClosureById func(
		instance T.Gpointer,
		signalId uint,
		detail T.GQuark,
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
		cHandler T.GCallback,
		data T.Gpointer,
		destroyData ClosureNotify,
		connectFlags T.GConnectFlags) T.Gulong

	SignalHandlerBlock func(instance T.Gpointer, handlerId T.Gulong)

	SignalHandlerUnblock func(instance T.Gpointer, handlerId T.Gulong)

	SignalHandlerDisconnect func(
		instance T.Gpointer, handlerId T.Gulong)

	SignalHandlerIsConnected func(
		instance T.Gpointer, handlerId T.Gulong) bool

	SignalHandlerFind func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signalId uint,
		detail T.GQuark,
		closure *Closure,
		fnc T.Gpointer,
		data T.Gpointer) T.Gulong

	SignalHandlersBlockMatched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signalId uint,
		detail T.GQuark,
		closure *Closure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalHandlersUnblockMatched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signalId uint,
		detail T.GQuark,
		closure *Closure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalHandlersDisconnectMatched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signalId uint,
		detail T.GQuark,
		closure *Closure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	SignalOverrideClassClosure func(
		signalId uint, instanceType Type, classClosure *Closure)

	SignalOverrideClassHandler func(
		signalName string, instanceType Type, classHandler T.GCallback)

	SignalChainFromOverridden func(
		instanceAndParams *Value, returnValue *Value)

	SignalChainFromOverriddenHandler func(
		instance T.Gpointer, v ...VArg)

	SignalAccumulatorTrueHandled func(ihint *T.GSignalInvocationHint,
		returnAccu, handlerReturn *Value, dummy T.Gpointer) bool

	SignalAccumulatorFirstWins func(ihint *T.GSignalInvocationHint,
		returnAccu, handlerReturn *Value, dummy T.Gpointer) bool

	SignalHandlersDestroy func(instance T.Gpointer)

	ObjectInterfaceInstallProperty func(
		gIface T.Gpointer, pspec *T.GParamSpec)

	ObjectInterfaceFindProperty func(
		gIface T.Gpointer, propertyName string) *T.GParamSpec

	ObjectInterfaceListProperties func(
		gIface T.Gpointer, nPropertiesP *uint) **T.GParamSpec

	CclosureNewObject func(
		callbackFunc T.GCallback, object *Object) *Closure

	CclosureNewObjectSwap func(
		callbackFunc T.GCallback, object *Object) *Closure

	SignalConnectObject func(instance T.Gpointer,
		detailedSignal string, cHandler T.GCallback,
		gobject T.Gpointer, connectFlags T.GConnectFlags) T.Gulong

	ObjectCompatControl func(what T.Gsize, data T.Gpointer) T.Gsize

	ClearObject func(objectPtr **Object)

	BoxedCopy func(boxedType Type, srcBoxed T.Gconstpointer) T.Gpointer

	BoxedFree func(boxedType Type, boxed T.Gpointer)

	BoxedTypeRegisterStatic func(name string,
		boxedCopy T.GBoxedCopyFunc, boxedFree T.GBoxedFreeFunc) Type

	CclosureNew func(callbackFunc T.GCallback,
		userData T.Gpointer, destroyData ClosureNotify) *Closure

	CclosureNewSwap func(callbackFunc T.GCallback,
		userData T.Gpointer, destroyData ClosureNotify) *Closure

	SignalTypeCclosureNew func(
		itype Type, structOffset uint) *Closure

	CclosureMarshal_VOID__VOID func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__BOOLEAN func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__CHAR func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__UCHAR func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__INT func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__UINT func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__LONG func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__ULONG func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__ENUM func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__FLAGS func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__FLOAT func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__DOUBLE func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__STRING func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__PARAM func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__BOXED func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__POINTER func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__OBJECT func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__VARIANT func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_VOID__UINT_POINTER func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_BOOLEAN__FLAGS func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_STRING__OBJECT_POINTER func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	CclosureMarshal_BOOLEAN__BOXED_BOXED func(
		closure *Closure,
		returnValue *Value,
		nParamValues uint,
		paramValues *Value,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	PointerTypeRegisterStatic func(name string) Type

	StrdupValueContents func(value *Value) string

	UnicharValidate func(ch T.Gunichar) bool

	SlistRemoveAll func(list *T.GSList, data T.Gconstpointer) *T.GSList

	//NOTE(t): No entry in dll
	// EnumTypesInit func()
	// ValueCInit func()
	// ValueTypesInit func()
	// ParamTypeInit func()
	// BoxedTypeInit func()
	// ParamSpecTypesInit func()
	// ValueTransformsInit func()
	// SignalInit func()
)

func SignalConnect(i T.Gpointer, s string,
	cb T.GCallback, d T.Gpointer) T.Gulong {
	return SignalConnectData(i, s, cb, d, nil, 0)
}

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
	{"g_closure_add_finalize_notifier", &closureAddFinalizeNotifier},
	{"g_closure_add_invalidate_notifier", &closureAddInvalidateNotifier},
	{"g_closure_add_marshal_guards", &closureAddMarshalGuards},
	{"g_closure_get_type", &ClosureGetType},
	{"g_closure_invalidate", &closureInvalidate},
	{"g_closure_invoke", &closureInvoke},
	{"g_closure_new_object", &ClosureNewObject},
	{"g_closure_new_simple", &ClosureNewSimple},
	{"g_closure_ref", &closureRef},
	{"g_closure_remove_finalize_notifier", &closureRemoveFinalizeNotifier},
	{"g_closure_remove_invalidate_notifier", &closureRemoveInvalidateNotifier},
	{"g_closure_set_marshal", &closureSetMarshal},
	{"g_closure_set_meta_marshal", &closureSetMetaMarshal},
	{"g_closure_sink", &closureSink},
	{"g_closure_unref", &closureUnref},
	{"g_date_get_type", &DateGetType},
	{"g_date_time_get_type", &DateTimeGetType},
	{"g_enum_complete_type_info", &EnumCompleteTypeInfo},
	{"g_enum_get_value", &EnumGetValue},
	{"g_enum_get_value_by_name", &EnumGetValueByName},
	{"g_enum_get_value_by_nick", &EnumGetValueByNick},
	{"g_enum_register_static", &EnumRegisterStatic},
	{"g_error_get_type", &ErrorGetType},
	{"g_flags_complete_type_info", &FlagsCompleteTypeInfo},
	{"g_flags_get_first_value", &flagsGetFirstValue},
	{"g_flags_get_value_by_name", &flagsGetValueByName},
	{"g_flags_get_value_by_nick", &flagsGetValueByNick},
	{"g_flags_register_static", &FlagsRegisterStatic},
	{"g_gstring_get_type", &GstringGetType},
	{"g_gtype_get_type", &GtypeGetType},
	{"g_hash_table_get_type", &HashTableGetType},
	{"g_initially_unowned_get_type", &InitiallyUnownedGetType},
	{"g_io_channel_get_type", &IoChannelGetType},
	{"g_io_condition_get_type", &IoConditionGetType},
	{"g_object_add_toggle_ref", &objectAddToggleRef},
	{"g_object_add_weak_pointer", &objectAddWeakPointer},
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
	{"g_object_force_floating", &objectForceFloating},
	{"g_object_freeze_notify", &objectFreezeNotify},
	{"g_object_get", &ObjectGet},
	{"g_object_get_data", &objectGetData},
	{"g_object_get_property", &objectGetProperty},
	{"g_object_get_qdata", &objectGetQdata},
	{"g_object_get_type", &ObjectGetType},
	{"g_object_get_valist", &objectGetValist},
	{"g_object_interface_find_property", &ObjectInterfaceFindProperty},
	{"g_object_interface_install_property", &ObjectInterfaceInstallProperty},
	{"g_object_interface_list_properties", &ObjectInterfaceListProperties},
	{"g_object_is_floating", &ObjectIsFloating},
	{"g_object_new", &ObjectNew},
	{"g_object_new_valist", &ObjectNewValist},
	{"g_object_newv", &ObjectNewv},
	{"g_object_notify", &objectNotify},
	{"g_object_notify_by_pspec", &objectNotifyByPspec},
	{"g_object_ref", &ObjectRef},
	{"g_object_ref_sink", &ObjectRefSink},
	{"g_object_remove_toggle_ref", &objectRemoveToggleRef},
	{"g_object_remove_weak_pointer", &objectRemoveWeakPointer},
	{"g_object_run_dispose", &objectRunDispose},
	{"g_object_set", &ObjectSet},
	{"g_object_set_data", &objectSetData},
	{"g_object_set_data_full", &objectSetDataFull},
	{"g_object_set_property", &objectSetProperty},
	{"g_object_set_qdata", &objectSetQdata},
	{"g_object_set_qdata_full", &objectSetQdataFull},
	{"g_object_set_valist", &objectSetValist},
	{"g_object_steal_data", &objectStealData},
	{"g_object_steal_qdata", &objectStealQdata},
	{"g_object_thaw_notify", &objectThawNotify},
	{"g_object_unref", &ObjectUnref},
	{"g_object_watch_closure", &objectWatchClosure},
	{"g_object_weak_ref", &objectWeakRef},
	{"g_object_weak_unref", &objectWeakUnref},
	{"g_param_spec_boolean", &ParamSpecBoolean},
	{"g_param_spec_boxed", &ParamSpecBoxed},
	{"g_param_spec_char", &ParamSpecChar},
	{"g_param_spec_double", &ParamSpecDouble},
	{"g_param_spec_enum", &ParamSpecEnum},
	{"g_param_spec_flags", &ParamSpecFlags},
	{"g_param_spec_float", &ParamSpecFloat},
	{"g_param_spec_get_blurb", &paramSpecGetBlurb},
	{"g_param_spec_get_name", &paramSpecGetName},
	{"g_param_spec_get_nick", &paramSpecGetNick},
	{"g_param_spec_get_qdata", &paramSpecGetQdata},
	{"g_param_spec_get_redirect_target", &paramSpecGetRedirectTarget},
	{"g_param_spec_gtype", &ParamSpecGtype},
	{"g_param_spec_int", &ParamSpecInt},
	{"g_param_spec_int64", &ParamSpecInt64},
	{"g_param_spec_internal", &ParamSpecInternal},
	{"g_param_spec_long", &ParamSpecLong},
	{"g_param_spec_object", &ParamSpecObject},
	{"g_param_spec_override", &ParamSpecOverride},
	{"g_param_spec_param", &ParamSpecParam},
	{"g_param_spec_pointer", &ParamSpecPointer},
	{"g_param_spec_pool_insert", &paramSpecPoolInsert},
	{"g_param_spec_pool_list", &paramSpecPoolList},
	{"g_param_spec_pool_list_owned", &paramSpecPoolListOwned},
	{"g_param_spec_pool_lookup", &paramSpecPoolLookup},
	{"g_param_spec_pool_new", &ParamSpecPoolNew},
	{"g_param_spec_pool_remove", &paramSpecPoolRemove},
	{"g_param_spec_ref", &paramSpecRef},
	{"g_param_spec_ref_sink", &paramSpecRefSink},
	{"g_param_spec_set_qdata", &paramSpecSetQdata},
	{"g_param_spec_set_qdata_full", &paramSpecSetQdataFull},
	{"g_param_spec_sink", &paramSpecSink},
	{"g_param_spec_steal_qdata", &paramSpecStealQdata},
	{"g_param_spec_string", &ParamSpecString},
	{"g_param_spec_uchar", &ParamSpecUchar},
	{"g_param_spec_uint", &ParamSpecUint},
	{"g_param_spec_uint64", &ParamSpecUint64},
	{"g_param_spec_ulong", &ParamSpecUlong},
	{"g_param_spec_unichar", &ParamSpecUnichar},
	{"g_param_spec_unref", &paramSpecUnref},
	{"g_param_spec_value_array", &ParamSpecValueArray},
	{"g_param_spec_variant", &ParamSpecVariant},
	{"g_param_type_register_static", &ParamTypeRegisterStatic},
	{"g_param_value_convert", &paramValueConvert},
	{"g_param_value_defaults", &paramValueDefaults},
	{"g_param_value_set_default", &paramValueSetDefault},
	{"g_param_value_validate", &paramValueValidate},
	{"g_param_values_cmp", &paramValuesCmp},
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
	{"g_type_add_class_private", &typeAddClassPrivate},
	{"g_type_add_interface_check", &TypeAddInterfaceCheck},
	{"g_type_add_interface_dynamic", &typeAddInterfaceDynamic},
	{"g_type_add_interface_static", &typeAddInterfaceStatic},
	{"g_type_check_class_cast", &typeCheckClassCast},
	{"g_type_check_class_is_a", &typeCheckClassIsA},
	{"g_type_check_instance", &typeCheckInstance},
	{"g_type_check_instance_cast", &typeCheckInstanceCast},
	{"g_type_check_instance_is_a", &typeCheckInstanceIsA},
	{"g_type_check_is_value_type", &typeCheckIsValueType},
	{"g_type_check_value", &TypeCheckValue},
	{"g_type_check_value_holds", &TypeCheckValueHolds},
	{"g_type_children", &typeChildren},
	{"g_type_class_add_private", &TypeClassAddPrivate},
	{"g_type_class_get_private", &typeClassGetPrivate},
	{"g_type_class_peek", &typeClassPeek},
	{"g_type_class_peek_parent", &TypeClassPeekParent},
	{"g_type_class_peek_static", &typeClassPeekStatic},
	{"g_type_class_ref", &typeClassRef},
	{"g_type_class_unref", &TypeClassUnref},
	{"g_type_class_unref_uncached", &TypeClassUnrefUncached},
	{"g_type_create_instance", &TypeCreateInstance},
	{"g_type_default_interface_peek", &typeDefaultInterfacePeek},
	{"g_type_default_interface_ref", &typeDefaultInterfaceRef},
	{"g_type_default_interface_unref", &TypeDefaultInterfaceUnref},
	{"g_type_depth", &typeDepth},
	{"g_type_free_instance", &typeFreeInstance},
	{"g_type_from_name", &TypeFromName},
	{"g_type_fundamental", &typeFundamental},
	{"g_type_fundamental_next", &TypeFundamentalNext},
	{"g_type_get_plugin", &typeGetPlugin},
	{"g_type_get_qdata", &typeGetQdata},
	{"g_type_init", &GTypeInit},
	{"g_type_init_with_debug_flags", &TypeInitWithDebugFlags},
	{"g_type_instance_get_private", &typeInstanceGetPrivate},
	{"g_type_interface_add_prerequisite", &typeInterfaceAddPrerequisite},
	{"g_type_interface_get_plugin", &typeInterfaceGetPlugin},
	{"g_type_interface_peek", &TypeInterfacePeek},
	{"g_type_interface_peek_parent", &TypeInterfacePeekParent},
	{"g_type_interface_prerequisites", &typeInterfacePrerequisites},
	{"g_type_interfaces", &typeInterfaces},
	{"g_type_is_a", &typeIsA},
	{"g_type_module_add_interface", &typeModuleAddInterface},
	{"g_type_module_get_type", &TypeModuleGetType},
	{"g_type_module_register_enum", &typeModuleRegisterEnum},
	{"g_type_module_register_flags", &typeModuleRegisterFlags},
	{"g_type_module_register_type", &typeModuleRegisterType},
	{"g_type_module_set_name", &typeModuleSetName},
	{"g_type_module_unuse", &typeModuleUnuse},
	{"g_type_module_use", &typeModuleUse},
	{"g_type_name", &typeName},
	{"g_type_name_from_class", &typeNameFromClass},
	{"g_type_name_from_instance", &typeNameFromInstance},
	{"g_type_next_base", &typeNextBase},
	{"g_type_parent", &typeParent},
	{"g_type_plugin_complete_interface_info", &typePluginCompleteInterfaceInfo},
	{"g_type_plugin_complete_type_info", &typePluginCompleteTypeInfo},
	{"g_type_plugin_get_type", &TypePluginGetType},
	{"g_type_plugin_unuse", &typePluginUnuse},
	{"g_type_plugin_use", &typePluginUse},
	{"g_type_qname", &typeQname},
	{"g_type_query", &typeQuery},
	{"g_type_register_dynamic", &typeRegisterDynamic},
	{"g_type_register_fundamental", &typeRegisterFundamental},
	{"g_type_register_static", &typeRegisterStatic},
	{"g_type_register_static_simple", &typeRegisterStaticSimple},
	{"g_type_remove_class_cache_func", &TypeRemoveClassCacheFunc},
	{"g_type_remove_interface_check", &TypeRemoveInterfaceCheck},
	{"g_type_set_qdata", &typeSetQdata},
	{"g_type_test_flags", &typeTestFlags},
	{"g_type_value_table_peek", &typeValueTablePeek},
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
	{"g_value_copy", &valueCopy},
	{"g_value_dup_boxed", &valueDupBoxed},
	{"g_value_dup_object", &valueDupObject},
	{"g_value_dup_param", &valueDupParam},
	{"g_value_dup_string", &valueDupString},
	{"g_value_dup_variant", &valueDupVariant},
	{"g_value_fits_pointer", &valueFitsPointer},
	{"g_value_get_boolean", &valueGetBoolean},
	{"g_value_get_boxed", &valueGetBoxed},
	{"g_value_get_char", &valueGetChar},
	{"g_value_get_double", &valueGetDouble},
	{"g_value_get_enum", &valueGetEnum},
	{"g_value_get_flags", &valueGetFlags},
	{"g_value_get_float", &valueGetFloat},
	{"g_value_get_gtype", &valueGetGtype},
	{"g_value_get_int", &valueGetInt},
	{"g_value_get_int64", &valueGetInt64},
	{"g_value_get_long", &valueGetLong},
	{"g_value_get_object", &valueGetObject},
	{"g_value_get_param", &valueGetParam},
	{"g_value_get_pointer", &valueGetPointer},
	{"g_value_get_string", &valueGetString},
	{"g_value_get_type", &ValueGetType},
	{"g_value_get_uchar", &valueGetUchar},
	{"g_value_get_uint", &valueGetUint},
	{"g_value_get_uint64", &valueGetUint64},
	{"g_value_get_ulong", &valueGetUlong},
	{"g_value_get_variant", &valueGetVariant},
	{"g_value_init", &valueInit},
	{"g_value_peek_pointer", &valuePeekPointer},
	{"g_value_register_transform_func", &ValueRegisterTransformFunc},
	{"g_value_reset", &valueReset},
	{"g_value_set_boolean", &valueSetBoolean},
	{"g_value_set_boxed", &valueSetBoxed},
	{"g_value_set_boxed_take_ownership", &valueSetBoxedTakeOwnership},
	{"g_value_set_char", &valueSetChar},
	{"g_value_set_double", &valueSetDouble},
	{"g_value_set_enum", &valueSetEnum},
	{"g_value_set_flags", &valueSetFlags},
	{"g_value_set_float", &valueSetFloat},
	{"g_value_set_gtype", &valueSetGtype},
	{"g_value_set_instance", &valueSetInstance},
	{"g_value_set_int", &valueSetInt},
	{"g_value_set_int64", &valueSetInt64},
	{"g_value_set_long", &valueSetLong},
	{"g_value_set_object", &valueSetObject},
	{"g_value_set_object_take_ownership", &valueSetObjectTakeOwnership},
	{"g_value_set_param", &valueSetParam},
	{"g_value_set_param_take_ownership", &valueSetParamTakeOwnership},
	{"g_value_set_pointer", &valueSetPointer},
	{"g_value_set_static_boxed", &valueSetStaticBoxed},
	{"g_value_set_static_string", &valueSetStaticString},
	{"g_value_set_string", &valueSetString},
	{"g_value_set_string_take_ownership", &valueSetStringTakeOwnership},
	{"g_value_set_uchar", &valueSetUchar},
	{"g_value_set_uint", &valueSetUint},
	{"g_value_set_uint64", &valueSetUint64},
	{"g_value_set_ulong", &valueSetUlong},
	{"g_value_set_variant", &valueSetVariant},
	{"g_value_take_boxed", &valueTakeBoxed},
	{"g_value_take_object", &valueTakeObject},
	{"g_value_take_param", &valueTakeParam},
	{"g_value_take_string", &valueTakeString},
	{"g_value_take_variant", &valueTakeVariant},
	{"g_value_transform", &valueTransform},
	{"g_value_type_compatible", &ValueTypeCompatible},
	{"g_value_type_transformable", &ValueTypeTransformable},
	{"g_value_unset", &valueUnset},
	{"g_variant_get_gtype", &VariantGetGtype},
	{"g_variant_type_get_gtype", &VariantTypeGetGtype},
}

var dataList = outside.Data{
	{"g_param_spec_types", (*Type)(nil)},
}
