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
	G_signal_newv func(
		signal_name string,
		itype T.GType,
		signal_flags T.GSignalFlags,
		class_closure *T.GClosure,
		accumulator T.GSignalAccumulator,
		accu_data T.Gpointer,
		c_marshaller T.GSignalCMarshaller,
		return_type T.GType,
		n_params uint,
		param_types *T.GType) uint

	G_signal_new_valist func(
		signal_name string,
		itype T.GType,
		signal_flags T.GSignalFlags,
		class_closure *T.GClosure,
		accumulator T.GSignalAccumulator,
		accu_data T.Gpointer,
		c_marshaller T.GSignalCMarshaller,
		return_type T.GType,
		n_params uint,
		args T.Va_list) uint

	G_signal_new func(signal_name string, itype T.GType,
		signal_flags T.GSignalFlags, class_offset uint,
		accumulator T.GSignalAccumulator, accu_data T.Gpointer,
		c_marshaller T.GSignalCMarshaller, return_type T.GType,
		n_params uint, v ...VArg) uint

	G_signal_new_class_handler func(signal_name string,
		itype T.GType, signal_flags T.GSignalFlags,
		class_handler T.GCallback, accumulator T.GSignalAccumulator,
		accu_data T.Gpointer, c_marshaller T.GSignalCMarshaller,
		return_type T.GType, n_params uint, v ...VArg) uint

	G_signal_emitv func(
		instance_and_params *T.GValue,
		signal_id uint,
		detail T.GQuark,
		return_value *T.GValue)

	G_signal_emit_valist func(
		instance T.Gpointer,
		signal_id uint,
		detail T.GQuark,
		var_args T.Va_list)

	G_signal_emit func(instance T.Gpointer, signal_id uint,
		detail T.GQuark, v ...VArg)

	G_signal_emit_by_name func(instance T.Gpointer,
		detailed_signal string, v ...VArg)

	G_signal_lookup func(name string, itype T.GType) uint

	G_signal_name func(signal_id uint) string

	G_signal_query func(signal_id uint, query *T.GSignalQuery)

	G_signal_list_ids func(itype T.GType, n_ids *uint) *uint

	G_signal_parse_name func(
		detailed_signal string,
		itype T.GType,
		signal_id_p *uint,
		detail_p *T.GQuark,
		force_detail_quark T.Gboolean) T.Gboolean

	G_signal_get_invocation_hint func(
		instance T.Gpointer) *T.GSignalInvocationHint

	G_signal_stop_emission func(
		instance T.Gpointer,
		signal_id uint,
		detail T.GQuark)

	G_signal_stop_emission_by_name func(
		instance T.Gpointer,
		detailed_signal string)

	G_signal_add_emission_hook func(
		signal_id uint,
		detail T.GQuark,
		hook_func T.GSignalEmissionHook,
		hook_data T.Gpointer,
		data_destroy T.GDestroyNotify) T.Gulong

	G_signal_remove_emission_hook func(
		signal_id uint,
		hook_id T.Gulong)

	G_signal_has_handler_pending func(
		instance T.Gpointer,
		signal_id uint,
		detail T.GQuark,
		may_be_blocked T.Gboolean) T.Gboolean

	G_signal_connect_closure_by_id func(
		instance T.Gpointer,
		signal_id uint,
		detail T.GQuark,
		closure *T.GClosure,
		after T.Gboolean) T.Gulong

	G_signal_connect_closure func(
		instance T.Gpointer,
		detailed_signal string,
		closure *T.GClosure,
		after T.Gboolean) T.Gulong

	G_signal_connect_data func(
		instance T.Gpointer,
		detailed_signal string,
		c_handler T.GCallback,
		data T.Gpointer,
		destroy_data T.GClosureNotify,
		connect_flags T.GConnectFlags) T.Gulong

	G_signal_handler_block func(
		instance T.Gpointer,
		handler_id T.Gulong)

	G_signal_handler_unblock func(
		instance T.Gpointer,
		handler_id T.Gulong)

	G_signal_handler_disconnect func(
		instance T.Gpointer,
		handler_id T.Gulong)

	G_signal_handler_is_connected func(
		instance T.Gpointer,
		handler_id T.Gulong) T.Gboolean

	G_signal_handler_find func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signal_id uint,
		detail T.GQuark,
		closure *T.GClosure,
		fnc T.Gpointer,
		data T.Gpointer) T.Gulong

	G_signal_handlers_block_matched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signal_id uint,
		detail T.GQuark,
		closure *T.GClosure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	G_signal_handlers_unblock_matched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signal_id uint,
		detail T.GQuark,
		closure *T.GClosure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	G_signal_handlers_disconnect_matched func(
		instance T.Gpointer,
		mask T.GSignalMatchType,
		signal_id uint,
		detail T.GQuark,
		closure *T.GClosure,
		fnc T.Gpointer,
		data T.Gpointer) uint

	G_signal_override_class_closure func(
		signal_id uint,
		instance_type T.GType,
		class_closure *T.GClosure)

	G_signal_override_class_handler func(
		signal_name string,
		instance_type T.GType,
		class_handler T.GCallback)

	G_signal_chain_from_overridden func(
		instance_and_params *T.GValue,
		return_value *T.GValue)

	G_signal_chain_from_overridden_handler func(
		instance T.Gpointer, v ...VArg)

	G_signal_accumulator_true_handled func(
		ihint *T.GSignalInvocationHint,
		return_accu *T.GValue,
		handler_return *T.GValue,
		dummy T.Gpointer) T.Gboolean

	G_signal_accumulator_first_wins func(
		ihint *T.GSignalInvocationHint,
		return_accu *T.GValue,
		handler_return *T.GValue,
		dummy T.Gpointer) T.Gboolean

	G_signal_handlers_destroy func(
		instance T.Gpointer)

	G_initially_unowned_get_type func() T.GType

	G_object_class_install_property func(
		oclass *T.GObjectClass,
		property_id uint,
		pspec *T.GParamSpec)

	G_object_class_find_property func(
		oclass *T.GObjectClass,
		property_name string) *T.GParamSpec

	G_object_class_list_properties func(
		oclass *T.GObjectClass,
		n_properties *uint) **T.GParamSpec

	G_object_class_override_property func(
		oclass *T.GObjectClass,
		property_id uint,
		name string)

	G_object_class_install_properties func(
		oclass *T.GObjectClass,
		n_pspecs uint,
		pspecs **T.GParamSpec)

	G_object_interface_install_property func(
		g_iface T.Gpointer,
		pspec *T.GParamSpec)

	G_object_interface_find_property func(
		g_iface T.Gpointer,
		property_name string) *T.GParamSpec

	G_object_interface_list_properties func(
		g_iface T.Gpointer,
		n_properties_p *uint) **T.GParamSpec

	G_object_get_type func() T.GType

	G_object_new func(object_type T.GType,
		first_property_name string, v ...VArg) T.Gpointer

	G_object_newv func(
		object_type T.GType,
		n_parameters uint,
		parameters *T.GParameter) T.Gpointer

	G_object_new_valist func(
		object_type T.GType,
		first_property_name string,
		var_args T.Va_list) *T.GObject

	G_object_set func(object T.Gpointer,
		first_property_name string, v ...VArg)

	G_object_get func(object T.Gpointer,
		first_property_name string, v ...VArg)

	G_object_connect func(object T.Gpointer,
		signal_spec string, v ...VArg) T.Gpointer

	G_object_disconnect func(object T.Gpointer,
		signal_spec string, v ...VArg)

	G_object_set_valist func(
		object *T.GObject,
		first_property_name string,
		var_args T.Va_list)

	G_object_get_valist func(
		object *T.GObject,
		first_property_name string,
		var_args T.Va_list)

	G_object_set_property func(
		object *T.GObject,
		property_name string,
		value *T.GValue)

	G_object_get_property func(
		object *T.GObject,
		property_name string,
		value *T.GValue)

	G_object_freeze_notify func(
		object *T.GObject)

	G_object_notify func(
		object *T.GObject,
		property_name string)

	G_object_notify_by_pspec func(
		object *T.GObject,
		pspec *T.GParamSpec)

	G_object_thaw_notify func(
		object *T.GObject)

	G_object_is_floating func(
		object T.Gpointer) T.Gboolean

	G_object_ref_sink func(
		object T.Gpointer) T.Gpointer

	G_object_ref func(
		object T.Gpointer) T.Gpointer

	G_object_unref func(
		object T.Gpointer)

	G_object_weak_ref func(
		object *T.GObject,
		notify T.GWeakNotify,
		data T.Gpointer)

	G_object_weak_unref func(
		object *T.GObject,
		notify T.GWeakNotify,
		data T.Gpointer)

	G_object_add_weak_pointer func(
		object *T.GObject,
		weak_pointer_location *T.Gpointer)

	G_object_remove_weak_pointer func(
		object *T.GObject,
		weak_pointer_location *T.Gpointer)

	G_object_add_toggle_ref func(
		object *T.GObject,
		notify T.GToggleNotify,
		data T.Gpointer)

	G_object_remove_toggle_ref func(
		object *T.GObject,
		notify T.GToggleNotify,
		data T.Gpointer)

	G_object_get_qdata func(
		object *T.GObject,
		quark T.GQuark) T.Gpointer

	G_object_set_qdata func(
		object *T.GObject,
		quark T.GQuark,
		data T.Gpointer)

	G_object_set_qdata_full func(
		object *T.GObject,
		quark T.GQuark,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	G_object_steal_qdata func(
		object *T.GObject,
		quark T.GQuark) T.Gpointer

	G_object_get_data func(
		object *T.GObject,
		key string) T.Gpointer

	G_object_set_data func(
		object *T.GObject,
		key string,
		data T.Gpointer)

	G_object_set_data_full func(
		object *T.GObject,
		key string,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	G_object_steal_data func(
		object *T.GObject,
		key string) T.Gpointer

	G_object_watch_closure func(
		object *T.GObject,
		closure *T.GClosure)

	G_cclosure_new_object func(
		callback_func T.GCallback,
		object *T.GObject) *T.GClosure

	G_cclosure_new_object_swap func(
		callback_func T.GCallback,
		object *T.GObject) *T.GClosure

	G_closure_new_object func(
		sizeof_closure uint,
		object *T.GObject) *T.GClosure

	G_value_set_object func(
		value *T.GValue,
		v_object T.Gpointer)

	G_value_get_object func(
		value *T.GValue) T.Gpointer

	G_value_dup_object func(
		value *T.GValue) T.Gpointer

	G_signal_connect_object func(
		instance T.Gpointer,
		detailed_signal string,
		c_handler T.GCallback,
		gobject T.Gpointer,
		connect_flags T.GConnectFlags) T.Gulong

	G_object_force_floating func(
		object *T.GObject)

	G_object_run_dispose func(
		object *T.GObject)

	G_value_take_object func(
		value *T.GValue,
		v_object T.Gpointer)

	G_value_set_object_take_ownership func(
		value *T.GValue,
		v_object T.Gpointer)

	G_object_compat_control func(
		what T.Gsize,
		data T.Gpointer) T.Gsize

	G_clear_object func(object_ptr **T.GObject)

	G_binding_flags_get_type func() T.GType

	G_binding_get_type func() T.GType

	G_binding_get_flags func(binding *T.GBinding) T.GBindingFlags

	G_binding_get_source func(binding *T.GBinding) *T.GObject

	G_binding_get_target func(binding *T.GBinding) *T.GObject

	G_binding_get_source_property func(binding *T.GBinding) string

	G_binding_get_target_property func(binding *T.GBinding) string

	G_object_bind_property func(
		source T.Gpointer,
		source_property string,
		target T.Gpointer,
		target_property string,
		flags T.GBindingFlags) *T.GBinding

	G_object_bind_property_full func(
		source T.Gpointer,
		source_property string,
		target T.Gpointer,
		target_property string,
		flags T.GBindingFlags,
		transform_to T.GBindingTransformFunc,
		transform_from T.GBindingTransformFunc,
		user_data T.Gpointer,
		notify T.GDestroyNotify) *T.GBinding

	G_object_bind_property_with_closures func(
		source T.Gpointer,
		source_property string,
		target T.Gpointer,
		target_property string,
		flags T.GBindingFlags,
		transform_to *T.GClosure,
		transform_from *T.GClosure) *T.GBinding

	G_boxed_copy func(
		boxed_type T.GType,
		src_boxed T.Gconstpointer) T.Gpointer

	G_boxed_free func(
		boxed_type T.GType,
		boxed T.Gpointer)

	G_value_set_boxed func(
		value *T.GValue,
		v_boxed T.Gconstpointer)

	G_value_set_static_boxed func(
		value *T.GValue,
		v_boxed T.Gconstpointer)

	G_value_get_boxed func(
		value *T.GValue) T.Gpointer

	G_value_dup_boxed func(
		value *T.GValue) T.Gpointer

	G_boxed_type_register_static func(
		name string,
		boxed_copy T.GBoxedCopyFunc,
		boxed_free T.GBoxedFreeFunc) T.GType

	G_value_take_boxed func(
		value *T.GValue,
		v_boxed T.Gconstpointer)

	G_value_set_boxed_take_ownership func(
		value *T.GValue,
		v_boxed T.Gconstpointer)

	G_closure_get_type func() T.GType

	G_value_get_type func() T.GType

	G_value_array_get_type func() T.GType

	G_date_get_type func() T.GType

	G_strv_get_type func() T.GType

	G_gstring_get_type func() T.GType

	G_hash_table_get_type func() T.GType

	G_array_get_type func() T.GType

	G_byte_array_get_type func() T.GType

	G_ptr_array_get_type func() T.GType

	G_variant_type_get_gtype func() T.GType

	G_regex_get_type func() T.GType

	G_error_get_type func() T.GType

	G_date_time_get_type func() T.GType

	G_variant_get_gtype func() T.GType

	G_type_init func()

	G_type_init_with_debug_flags func(
		debug_flags T.GTypeDebugFlags)

	G_type_name func(t T.GType) string

	G_type_qname func(t T.GType) T.GQuark

	G_type_from_name func(name string) T.GType

	G_type_parent func(t T.GType) T.GType

	G_type_depth func(t T.GType) uint

	G_type_next_base func(leaf_type, root_type T.GType) T.GType

	G_type_is_a func(t, is_a_type T.GType) T.Gboolean

	G_type_class_ref func(
		t T.GType) T.Gpointer

	G_type_class_peek func(
		t T.GType) T.Gpointer

	G_type_class_peek_static func(
		t T.GType) T.Gpointer

	G_type_class_unref func(
		g_class T.Gpointer)

	G_type_class_peek_parent func(
		g_class T.Gpointer) T.Gpointer

	G_type_interface_peek func(
		instance_class T.Gpointer,
		iface_type T.GType) T.Gpointer

	G_type_interface_peek_parent func(
		g_iface T.Gpointer) T.Gpointer

	G_type_default_interface_ref func(
		g_type T.GType) T.Gpointer

	G_type_default_interface_peek func(
		g_type T.GType) T.Gpointer

	G_type_default_interface_unref func(
		g_iface T.Gpointer)

	G_type_children func(
		t T.GType,
		n_children *uint) *T.GType

	G_type_interfaces func(
		t T.GType,
		n_interfaces *uint) *T.GType

	G_type_set_qdata func(
		t T.GType,
		quark T.GQuark,
		data T.Gpointer)

	G_type_get_qdata func(
		t T.GType,
		quark T.GQuark) T.Gpointer

	G_type_query func(
		t T.GType,
		query *T.GTypeQuery)

	G_cclosure_new func(
		callback_func T.GCallback,
		user_data T.Gpointer,
		destroy_data T.GClosureNotify) *T.GClosure

	G_cclosure_new_swap func(
		callback_func T.GCallback,
		user_data T.Gpointer,
		destroy_data T.GClosureNotify) *T.GClosure

	G_signal_type_cclosure_new func(
		itype T.GType,
		struct_offset uint) *T.GClosure

	G_closure_ref func(
		closure *T.GClosure) *T.GClosure

	G_closure_sink func(
		closure *T.GClosure)

	G_closure_unref func(
		closure *T.GClosure)

	G_closure_new_simple func(
		sizeof_closure uint,
		data T.Gpointer) *T.GClosure

	G_closure_add_finalize_notifier func(
		closure *T.GClosure,
		notify_data T.Gpointer,
		notify_func T.GClosureNotify)

	G_closure_remove_finalize_notifier func(
		closure *T.GClosure,
		notify_data T.Gpointer,
		notify_func T.GClosureNotify)

	G_closure_add_invalidate_notifier func(
		closure *T.GClosure,
		notify_data T.Gpointer,
		notify_func T.GClosureNotify)

	G_closure_remove_invalidate_notifier func(
		closure *T.GClosure,
		notify_data T.Gpointer,
		notify_func T.GClosureNotify)

	G_closure_add_marshal_guards func(
		closure *T.GClosure,
		pre_marshal_data T.Gpointer,
		pre_marshal_notify T.GClosureNotify,
		post_marshal_data T.Gpointer,
		post_marshal_notify T.GClosureNotify)

	G_closure_set_marshal func(
		closure *T.GClosure,
		marshal T.GClosureMarshal)

	G_closure_set_meta_marshal func(
		closure *T.GClosure,
		marshal_data T.Gpointer,
		meta_marshal T.GClosureMarshal)

	G_closure_invalidate func(
		closure *T.GClosure)

	G_closure_invoke func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer)

	G_cclosure_marshal_VOID__VOID func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__BOOLEAN func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__CHAR func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__UCHAR func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__INT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__UINT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__LONG func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__ULONG func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__ENUM func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__FLAGS func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__FLOAT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__DOUBLE func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__STRING func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__PARAM func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__BOXED func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__OBJECT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__VARIANT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_VOID__UINT_POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_BOOLEAN__FLAGS func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_STRING__OBJECT_POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_cclosure_marshal_BOOLEAN__BOXED_BOXED func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values uint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	G_enum_get_value func(
		enum_class *T.GEnumClass,
		value int) *T.GEnumValue

	G_enum_get_value_by_name func(
		enum_class *T.GEnumClass,
		name string) *T.GEnumValue

	G_enum_get_value_by_nick func(
		enum_class *T.GEnumClass,
		nick string) *T.GEnumValue

	G_flags_get_first_value func(
		flags_class *T.GFlagsClass,
		value uint) *T.GFlagsValue

	G_flags_get_value_by_name func(
		flags_class *T.GFlagsClass,
		name string) *T.GFlagsValue

	G_flags_get_value_by_nick func(
		flags_class *T.GFlagsClass,
		nick string) *T.GFlagsValue

	G_value_set_enum func(
		value *T.GValue,
		v_enum int)

	G_value_get_enum func(
		value *T.GValue) int

	G_value_set_flags func(
		value *T.GValue,
		v_flags uint)

	G_value_get_flags func(
		value *T.GValue) uint

	G_enum_register_static func(
		name string,
		const_static_values *T.GEnumValue) T.GType

	G_flags_register_static func(
		name string,
		const_static_values *T.GFlagsValue) T.GType

	G_enum_complete_type_info func(
		g_enum_type T.GType,
		info *T.GTypeInfo,
		const_values *T.GEnumValue)

	G_flags_complete_type_info func(
		g_flags_type T.GType,
		info *T.GTypeInfo,
		const_values *T.GFlagsValue)

	G_param_spec_ref func(
		pspec *T.GParamSpec) *T.GParamSpec

	G_param_spec_unref func(
		pspec *T.GParamSpec)

	G_param_spec_sink func(
		pspec *T.GParamSpec)

	G_param_spec_ref_sink func(
		pspec *T.GParamSpec) *T.GParamSpec

	G_param_spec_get_qdata func(
		pspec *T.GParamSpec,
		quark T.GQuark) T.Gpointer

	G_param_spec_set_qdata func(
		pspec *T.GParamSpec,
		quark T.GQuark,
		data T.Gpointer)

	G_param_spec_set_qdata_full func(
		pspec *T.GParamSpec,
		quark T.GQuark,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	G_param_spec_steal_qdata func(
		pspec *T.GParamSpec,
		quark T.GQuark) T.Gpointer

	G_param_spec_get_redirect_target func(
		pspec *T.GParamSpec) *T.GParamSpec

	G_param_value_set_default func(
		pspec *T.GParamSpec,
		value *T.GValue)

	G_param_value_defaults func(
		pspec *T.GParamSpec,
		value *T.GValue) T.Gboolean

	G_param_value_validate func(
		pspec *T.GParamSpec,
		value *T.GValue) T.Gboolean

	G_param_value_convert func(
		pspec *T.GParamSpec,
		src_value *T.GValue,
		dest_value *T.GValue,
		strict_validation T.Gboolean) T.Gboolean

	G_param_values_cmp func(
		pspec *T.GParamSpec,
		value1 *T.GValue,
		value2 *T.GValue) int

	G_param_spec_get_name func(
		pspec *T.GParamSpec) string

	G_param_spec_get_nick func(
		pspec *T.GParamSpec) string

	G_param_spec_get_blurb func(
		pspec *T.GParamSpec) string

	G_value_set_param func(
		value *T.GValue,
		param *T.GParamSpec)

	G_value_get_param func(
		value *T.GValue) *T.GParamSpec

	G_value_dup_param func(
		value *T.GValue) *T.GParamSpec

	G_value_take_param func(
		value *T.GValue,
		param *T.GParamSpec)

	G_value_set_param_take_ownership func(
		value *T.GValue,
		param *T.GParamSpec)

	G_value_array_get_nth func(
		value_array *T.GValueArray,
		index_ uint) *T.GValue

	G_value_array_new func(n_prealloced uint) *T.GValueArray

	G_value_array_free func(value_array *T.GValueArray)

	G_value_array_copy func(
		value_array *T.GValueArray) *T.GValueArray

	G_value_array_prepend func(
		value_array *T.GValueArray,
		value *T.GValue) *T.GValueArray

	G_value_array_append func(
		value_array *T.GValueArray,
		value *T.GValue) *T.GValueArray

	G_value_array_insert func(
		value_array *T.GValueArray,
		index_ uint,
		value *T.GValue) *T.GValueArray

	G_value_array_remove func(
		value_array *T.GValueArray,
		index_ uint) *T.GValueArray

	G_value_array_sort func(
		value_array *T.GValueArray,
		compare_func T.GCompareFunc) *T.GValueArray

	G_value_array_sort_with_data func(
		value_array *T.GValueArray,
		compare_func T.GCompareDataFunc,
		user_data T.Gpointer) *T.GValueArray

	G_value_set_char func(value *T.GValue, v_char T.Gchar)

	G_value_get_char func(value *T.GValue) T.Gchar

	G_value_set_uchar func(
		value *T.GValue, v_uchar T.Guchar)

	G_value_get_uchar func(value *T.GValue) T.Guchar

	G_value_set_boolean func(
		value *T.GValue, v_boolean T.Gboolean)

	G_value_get_boolean func(value *T.GValue) T.Gboolean

	G_value_set_int func(value *T.GValue, v_int int)

	G_value_get_int func(value *T.GValue) int

	G_value_set_uint func(value *T.GValue, v_uint uint)

	G_value_get_uint func(value *T.GValue) uint

	G_value_set_long func(value *T.GValue, v_long T.Glong)

	G_value_get_long func(value *T.GValue) T.Glong

	G_value_set_ulong func(value *T.GValue, v_ulong T.Gulong)

	G_value_get_ulong func(value *T.GValue) T.Gulong

	G_value_set_int64 func(value *T.GValue, v_int64 int64)

	G_value_get_int64 func(value *T.GValue) int64

	G_value_set_uint64 func(value *T.GValue, v_uint64 uint64)

	G_value_get_uint64 func(value *T.GValue) uint64

	G_value_set_float func(value *T.GValue, v_float float32)

	G_value_get_float func(value *T.GValue) float32

	G_value_set_double func(value *T.GValue, v_double float64)

	G_value_get_double func(value *T.GValue) float64

	G_value_set_string func(value *T.GValue, v_string string)

	G_value_set_static_string func(
		value *T.GValue, v_string string)

	G_value_get_string func(value *T.GValue) string

	G_value_dup_string func(value *T.GValue) string

	G_value_set_pointer func(
		value *T.GValue, v_pointer T.Gpointer)

	G_value_get_pointer func(value *T.GValue) T.Gpointer

	G_gtype_get_type func() T.GType

	G_value_set_gtype func(value *T.GValue, v_gtype T.GType)

	G_value_get_gtype func(value *T.GValue) T.GType

	G_value_set_variant func(
		value *T.GValue, variant *T.GVariant)

	G_value_take_variant func(
		value *T.GValue, variant *T.GVariant)

	G_value_get_variant func(value *T.GValue) *T.GVariant

	G_value_dup_variant func(value *T.GValue) *T.GVariant

	G_pointer_type_register_static func(name string) T.GType

	G_strdup_value_contents func(value *T.GValue) string

	G_value_take_string func(
		value *T.GValue, v_string string)

	G_value_set_string_take_ownership func(
		value *T.GValue, v_string string)

	G_type_register_static func(
		parent_type T.GType,
		type_name string,
		info *T.GTypeInfo,
		flags T.GTypeFlags) T.GType

	G_type_register_static_simple func(
		parent_type T.GType,
		type_name string,
		class_size uint,
		class_init T.GClassInitFunc,
		instance_size uint,
		instance_init T.GInstanceInitFunc,
		flags T.GTypeFlags) T.GType

	G_type_register_dynamic func(
		parent_type T.GType,
		type_name string,
		plugin *T.GTypePlugin,
		flags T.GTypeFlags) T.GType

	G_type_register_fundamental func(
		type_id T.GType,
		type_name string,
		info *T.GTypeInfo,
		finfo *T.GTypeFundamentalInfo,
		flags T.GTypeFlags) T.GType

	G_type_add_interface_static func(
		instance_type T.GType,
		interface_type T.GType,
		info *T.GInterfaceInfo)

	G_type_add_interface_dynamic func(
		instance_type T.GType,
		interface_type T.GType,
		plugin *T.GTypePlugin)

	G_type_interface_add_prerequisite func(
		interface_type T.GType,
		prerequisite_type T.GType)

	G_type_interface_prerequisites func(
		interface_type T.GType,
		n_prerequisites *uint) *T.GType

	G_type_class_add_private func(
		g_class T.Gpointer,
		private_size T.Gsize)

	G_type_instance_get_private func(
		instance *T.GTypeInstance,
		private_type T.GType) T.Gpointer

	G_type_add_class_private func(
		class_type T.GType,
		private_size T.Gsize)

	G_type_class_get_private func(
		klass *T.GTypeClass,
		private_type T.GType) T.Gpointer

	G_type_get_plugin func(
		t T.GType) *T.GTypePlugin

	G_type_interface_get_plugin func(
		instance_type T.GType,
		interface_type T.GType) *T.GTypePlugin

	G_type_fundamental_next func() T.GType

	G_type_fundamental func(
		type_id T.GType) T.GType

	G_type_create_instance func(
		t T.GType) *T.GTypeInstance

	G_type_free_instance func(
		instance *T.GTypeInstance)

	G_type_add_class_cache_func func(
		cache_data T.Gpointer,
		cache_func T.GTypeClassCacheFunc)

	G_type_remove_class_cache_func func(
		cache_data T.Gpointer,
		cache_func T.GTypeClassCacheFunc)

	G_type_class_unref_uncached func(
		g_class T.Gpointer)

	G_type_add_interface_check func(
		check_data T.Gpointer,
		check_func T.GTypeInterfaceCheckFunc)

	G_type_remove_interface_check func(
		check_data T.Gpointer,
		check_func T.GTypeInterfaceCheckFunc)

	G_type_value_table_peek func(
		t T.GType) *T.GTypeValueTable

	G_type_check_instance func(
		instance *T.GTypeInstance) T.Gboolean

	G_type_check_instance_cast func(
		instance *T.GTypeInstance,
		iface_type T.GType) *T.GTypeInstance

	G_type_check_instance_is_a func(
		instance *T.GTypeInstance,
		iface_type T.GType) T.Gboolean

	G_type_check_class_cast func(
		g_class *T.GTypeClass,
		is_a_type T.GType) *T.GTypeClass

	G_type_check_class_is_a func(
		g_class *T.GTypeClass,
		is_a_type T.GType) T.Gboolean

	G_type_check_is_value_type func(
		t T.GType) T.Gboolean

	G_type_check_value func(
		value *T.GValue) T.Gboolean

	G_type_check_value_holds func(
		value *T.GValue,
		t T.GType) T.Gboolean

	G_type_test_flags func(
		t T.GType,
		flags uint) T.Gboolean

	G_type_name_from_instance func(
		instance *T.GTypeInstance) string

	G_type_name_from_class func(
		g_class *T.GTypeClass) string

	G_value_c_init func()

	G_value_types_init func()

	G_enum_types_init func()

	G_param_type_init func()

	G_boxed_type_init func()

	G_object_type_init func()

	G_param_spec_types_init func()

	G_value_transforms_init func()

	G_signal_init func()

	G_value_init func(
		value *T.GValue,
		g_type T.GType) *T.GValue

	G_value_copy func(
		src_value *T.GValue,
		dest_value *T.GValue)

	G_value_reset func(
		value *T.GValue) *T.GValue

	G_value_unset func(
		value *T.GValue)

	G_value_set_instance func(
		value *T.GValue,
		instance T.Gpointer)

	G_value_fits_pointer func(
		value *T.GValue) T.Gboolean

	G_value_peek_pointer func(
		value *T.GValue) T.Gpointer

	G_value_type_compatible func(
		src_type T.GType,
		dest_type T.GType) T.Gboolean

	G_value_type_transformable func(
		src_type T.GType,
		dest_type T.GType) T.Gboolean

	G_value_transform func(
		src_value *T.GValue,
		dest_value *T.GValue) T.Gboolean

	G_value_register_transform_func func(
		src_type T.GType,
		dest_type T.GType,
		transform_func T.GValueTransform)

	G_type_plugin_get_type func() T.GType

	G_type_plugin_use func(
		plugin *T.GTypePlugin)

	G_type_plugin_unuse func(
		plugin *T.GTypePlugin)

	G_type_plugin_complete_type_info func(
		plugin *T.GTypePlugin,
		g_type T.GType,
		info *T.GTypeInfo,
		value_table *T.GTypeValueTable)

	G_type_plugin_complete_interface_info func(
		plugin *T.GTypePlugin,
		instance_type T.GType,
		interface_type T.GType,
		info *T.GInterfaceInfo)

	G_type_module_get_type func() T.GType

	G_type_module_use func(
		module *T.GTypeModule) T.Gboolean

	G_type_module_unuse func(
		module *T.GTypeModule)

	G_type_module_set_name func(
		module *T.GTypeModule,
		name string)

	G_type_module_register_type func(
		module *T.GTypeModule,
		parent_type T.GType,
		type_name string,
		type_info *T.GTypeInfo,
		flags T.GTypeFlags) T.GType

	G_type_module_add_interface func(
		module *T.GTypeModule,
		instance_type T.GType,
		interface_type T.GType,
		interface_info *T.GInterfaceInfo)

	G_type_module_register_enum func(
		module *T.GTypeModule,
		name string,
		const_static_values *T.GEnumValue) T.GType

	G_type_module_register_flags func(
		module *T.GTypeModule,
		name string,
		const_static_values *T.GFlagsValue) T.GType

	G_param_spec_char func(
		name string,
		nick string,
		blurb string,
		minimum int8,
		maximum int8,
		default_value int8,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_uchar func(
		name string,
		nick string,
		blurb string,
		minimum uint8,
		maximum uint8,
		default_value uint8,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_boolean func(
		name string,
		nick string,
		blurb string,
		default_value T.Gboolean,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_int func(
		name string,
		nick string,
		blurb string,
		minimum int,
		maximum int,
		default_value int,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_uint func(
		name string,
		nick string,
		blurb string,
		minimum uint,
		maximum uint,
		default_value uint,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_long func(
		name string,
		nick string,
		blurb string,
		minimum T.Glong,
		maximum T.Glong,
		default_value T.Glong,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_ulong func(
		name string,
		nick string,
		blurb string,
		minimum T.Gulong,
		maximum T.Gulong,
		default_value T.Gulong,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_int64 func(
		name string,
		nick string,
		blurb string,
		minimum int64,
		maximum int64,
		default_value int64,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_uint64 func(
		name string,
		nick string,
		blurb string,
		minimum uint64,
		maximum uint64,
		default_value uint64,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_unichar func(
		name string,
		nick string,
		blurb string,
		default_value T.Gunichar,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_enum func(
		name string,
		nick string,
		blurb string,
		enum_type T.GType,
		default_value int,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_flags func(
		name string,
		nick string,
		blurb string,
		flags_type T.GType,
		default_value uint,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_float func(
		name string,
		nick string,
		blurb string,
		minimum float32,
		maximum float32,
		default_value float32,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_double func(
		name string,
		nick string,
		blurb string,
		minimum float64,
		maximum float64,
		default_value float64,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_string func(
		name string,
		nick string,
		blurb string,
		default_value string,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_param func(
		name string,
		nick string,
		blurb string,
		param_type T.GType,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_boxed func(
		name string,
		nick string,
		blurb string,
		boxed_type T.GType,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_pointer func(
		name string,
		nick string,
		blurb string,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_value_array func(
		name string,
		nick string,
		blurb string,
		element_spec *T.GParamSpec,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_object func(
		name string,
		nick string,
		blurb string,
		object_type T.GType,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_override func(
		name string,
		overridden *T.GParamSpec) *T.GParamSpec

	G_param_spec_gtype func(
		name string,
		nick string,
		blurb string,
		is_a_type T.GType,
		flags T.GParamFlags) *T.GParamSpec

	G_param_spec_variant func(
		name string,
		nick string,
		blurb string,
		t *T.GVariantType,
		default_value *T.GVariant,
		flags T.GParamFlags) *T.GParamSpec

	G_source_set_closure func(
		source *T.GSource,
		closure *T.GClosure)

	G_source_set_dummy_callback func(
		source *T.GSource)

	G_io_channel_get_type func() T.GType

	G_io_condition_get_type func() T.GType

	G_param_type_register_static func(
		name string,
		pspec_info *T.GParamSpecTypeInfo) T.GType

	G_param_spec_internal func(
		param_type T.GType,
		name string,
		nick string,
		blurb string,
		flags T.GParamFlags) T.Gpointer

	G_param_spec_pool_new func(
		type_prefixing T.Gboolean) *T.GParamSpecPool

	G_param_spec_pool_insert func(
		pool *T.GParamSpecPool,
		pspec *T.GParamSpec,
		owner_type T.GType)

	G_param_spec_pool_remove func(
		pool *T.GParamSpecPool,
		pspec *T.GParamSpec)

	G_param_spec_pool_lookup func(
		pool *T.GParamSpecPool,
		param_name string,
		owner_type T.GType,
		walk_ancestors T.Gboolean) *T.GParamSpec

	G_param_spec_pool_list_owned func(
		pool *T.GParamSpecPool,
		owner_type T.GType) *T.GList

	G_param_spec_pool_list func(
		pool *T.GParamSpecPool,
		owner_type T.GType,
		n_pspecs_p *uint) **T.GParamSpec

	G_unichar_validate func(
		ch T.Gunichar) T.Gboolean

	G_slist_remove_all func(
		list *T.GSList,
		data T.Gconstpointer) *T.GSList
)

var dll = "libgobject-2.0-0.dll"

var apiList = outside.Apis{
	{"g_array_get_type", &G_array_get_type},
	{"g_binding_flags_get_type", &G_binding_flags_get_type},
	{"g_binding_get_flags", &G_binding_get_flags},
	{"g_binding_get_source", &G_binding_get_source},
	{"g_binding_get_source_property", &G_binding_get_source_property},
	{"g_binding_get_target", &G_binding_get_target},
	{"g_binding_get_target_property", &G_binding_get_target_property},
	{"g_binding_get_type", &G_binding_get_type},
	{"g_boxed_copy", &G_boxed_copy},
	{"g_boxed_free", &G_boxed_free},
	{"g_boxed_type_register_static", &G_boxed_type_register_static},
	{"g_byte_array_get_type", &G_byte_array_get_type},
	{"g_cclosure_marshal_BOOLEAN__BOXED_BOXED", &G_cclosure_marshal_BOOLEAN__BOXED_BOXED},
	{"g_cclosure_marshal_BOOLEAN__FLAGS", &G_cclosure_marshal_BOOLEAN__FLAGS},
	{"g_cclosure_marshal_STRING__OBJECT_POINTER", &G_cclosure_marshal_STRING__OBJECT_POINTER},
	{"g_cclosure_marshal_VOID__BOOLEAN", &G_cclosure_marshal_VOID__BOOLEAN},
	{"g_cclosure_marshal_VOID__BOXED", &G_cclosure_marshal_VOID__BOXED},
	{"g_cclosure_marshal_VOID__CHAR", &G_cclosure_marshal_VOID__CHAR},
	{"g_cclosure_marshal_VOID__DOUBLE", &G_cclosure_marshal_VOID__DOUBLE},
	{"g_cclosure_marshal_VOID__ENUM", &G_cclosure_marshal_VOID__ENUM},
	{"g_cclosure_marshal_VOID__FLAGS", &G_cclosure_marshal_VOID__FLAGS},
	{"g_cclosure_marshal_VOID__FLOAT", &G_cclosure_marshal_VOID__FLOAT},
	{"g_cclosure_marshal_VOID__INT", &G_cclosure_marshal_VOID__INT},
	{"g_cclosure_marshal_VOID__LONG", &G_cclosure_marshal_VOID__LONG},
	{"g_cclosure_marshal_VOID__OBJECT", &G_cclosure_marshal_VOID__OBJECT},
	{"g_cclosure_marshal_VOID__PARAM", &G_cclosure_marshal_VOID__PARAM},
	{"g_cclosure_marshal_VOID__POINTER", &G_cclosure_marshal_VOID__POINTER},
	{"g_cclosure_marshal_VOID__STRING", &G_cclosure_marshal_VOID__STRING},
	{"g_cclosure_marshal_VOID__UCHAR", &G_cclosure_marshal_VOID__UCHAR},
	{"g_cclosure_marshal_VOID__UINT", &G_cclosure_marshal_VOID__UINT},
	{"g_cclosure_marshal_VOID__UINT_POINTER", &G_cclosure_marshal_VOID__UINT_POINTER},
	{"g_cclosure_marshal_VOID__ULONG", &G_cclosure_marshal_VOID__ULONG},
	{"g_cclosure_marshal_VOID__VARIANT", &G_cclosure_marshal_VOID__VARIANT},
	{"g_cclosure_marshal_VOID__VOID", &G_cclosure_marshal_VOID__VOID},
	{"g_cclosure_new", &G_cclosure_new},
	{"g_cclosure_new_object", &G_cclosure_new_object},
	{"g_cclosure_new_object_swap", &G_cclosure_new_object_swap},
	{"g_cclosure_new_swap", &G_cclosure_new_swap},
	{"g_clear_object", &G_clear_object},
	{"g_closure_add_finalize_notifier", &G_closure_add_finalize_notifier},
	{"g_closure_add_invalidate_notifier", &G_closure_add_invalidate_notifier},
	{"g_closure_add_marshal_guards", &G_closure_add_marshal_guards},
	{"g_closure_get_type", &G_closure_get_type},
	{"g_closure_invalidate", &G_closure_invalidate},
	{"g_closure_invoke", &G_closure_invoke},
	{"g_closure_new_object", &G_closure_new_object},
	{"g_closure_new_simple", &G_closure_new_simple},
	{"g_closure_ref", &G_closure_ref},
	{"g_closure_remove_finalize_notifier", &G_closure_remove_finalize_notifier},
	{"g_closure_remove_invalidate_notifier", &G_closure_remove_invalidate_notifier},
	{"g_closure_set_marshal", &G_closure_set_marshal},
	{"g_closure_set_meta_marshal", &G_closure_set_meta_marshal},
	{"g_closure_sink", &G_closure_sink},
	{"g_closure_unref", &G_closure_unref},
	{"g_date_get_type", &G_date_get_type},
	{"g_date_time_get_type", &G_date_time_get_type},
	{"g_enum_complete_type_info", &G_enum_complete_type_info},
	{"g_enum_get_value", &G_enum_get_value},
	{"g_enum_get_value_by_name", &G_enum_get_value_by_name},
	{"g_enum_get_value_by_nick", &G_enum_get_value_by_nick},
	{"g_enum_register_static", &G_enum_register_static},
	{"g_error_get_type", &G_error_get_type},
	{"g_flags_complete_type_info", &G_flags_complete_type_info},
	{"g_flags_get_first_value", &G_flags_get_first_value},
	{"g_flags_get_value_by_name", &G_flags_get_value_by_name},
	{"g_flags_get_value_by_nick", &G_flags_get_value_by_nick},
	{"g_flags_register_static", &G_flags_register_static},
	{"g_gstring_get_type", &G_gstring_get_type},
	{"g_gtype_get_type", &G_gtype_get_type},
	{"g_hash_table_get_type", &G_hash_table_get_type},
	{"g_initially_unowned_get_type", &G_initially_unowned_get_type},
	{"g_io_channel_get_type", &G_io_channel_get_type},
	{"g_io_condition_get_type", &G_io_condition_get_type},
	{"g_object_add_toggle_ref", &G_object_add_toggle_ref},
	{"g_object_add_weak_pointer", &G_object_add_weak_pointer},
	{"g_object_bind_property", &G_object_bind_property},
	{"g_object_bind_property_full", &G_object_bind_property_full},
	{"g_object_bind_property_with_closures", &G_object_bind_property_with_closures},
	{"g_object_class_find_property", &G_object_class_find_property},
	{"g_object_class_install_properties", &G_object_class_install_properties},
	{"g_object_class_install_property", &G_object_class_install_property},
	{"g_object_class_list_properties", &G_object_class_list_properties},
	{"g_object_class_override_property", &G_object_class_override_property},
	{"g_object_compat_control", &G_object_compat_control},
	{"g_object_connect", &G_object_connect},
	{"g_object_disconnect", &G_object_disconnect},
	{"g_object_force_floating", &G_object_force_floating},
	{"g_object_freeze_notify", &G_object_freeze_notify},
	{"g_object_get", &G_object_get},
	{"g_object_get_data", &G_object_get_data},
	{"g_object_get_property", &G_object_get_property},
	{"g_object_get_qdata", &G_object_get_qdata},
	{"g_object_get_type", &G_object_get_type},
	{"g_object_get_valist", &G_object_get_valist},
	{"g_object_interface_find_property", &G_object_interface_find_property},
	{"g_object_interface_install_property", &G_object_interface_install_property},
	{"g_object_interface_list_properties", &G_object_interface_list_properties},
	{"g_object_is_floating", &G_object_is_floating},
	{"g_object_new", &G_object_new},
	{"g_object_new_valist", &G_object_new_valist},
	{"g_object_newv", &G_object_newv},
	{"g_object_notify", &G_object_notify},
	{"g_object_notify_by_pspec", &G_object_notify_by_pspec},
	{"g_object_ref", &G_object_ref},
	{"g_object_ref_sink", &G_object_ref_sink},
	{"g_object_remove_toggle_ref", &G_object_remove_toggle_ref},
	{"g_object_remove_weak_pointer", &G_object_remove_weak_pointer},
	{"g_object_run_dispose", &G_object_run_dispose},
	{"g_object_set", &G_object_set},
	{"g_object_set_data", &G_object_set_data},
	{"g_object_set_data_full", &G_object_set_data_full},
	{"g_object_set_property", &G_object_set_property},
	{"g_object_set_qdata", &G_object_set_qdata},
	{"g_object_set_qdata_full", &G_object_set_qdata_full},
	{"g_object_set_valist", &G_object_set_valist},
	{"g_object_steal_data", &G_object_steal_data},
	{"g_object_steal_qdata", &G_object_steal_qdata},
	{"g_object_thaw_notify", &G_object_thaw_notify},
	{"g_object_unref", &G_object_unref},
	{"g_object_watch_closure", &G_object_watch_closure},
	{"g_object_weak_ref", &G_object_weak_ref},
	{"g_object_weak_unref", &G_object_weak_unref},
	{"g_param_spec_boolean", &G_param_spec_boolean},
	{"g_param_spec_boxed", &G_param_spec_boxed},
	{"g_param_spec_char", &G_param_spec_char},
	{"g_param_spec_double", &G_param_spec_double},
	{"g_param_spec_enum", &G_param_spec_enum},
	{"g_param_spec_flags", &G_param_spec_flags},
	{"g_param_spec_float", &G_param_spec_float},
	{"g_param_spec_get_blurb", &G_param_spec_get_blurb},
	{"g_param_spec_get_name", &G_param_spec_get_name},
	{"g_param_spec_get_nick", &G_param_spec_get_nick},
	{"g_param_spec_get_qdata", &G_param_spec_get_qdata},
	{"g_param_spec_get_redirect_target", &G_param_spec_get_redirect_target},
	{"g_param_spec_gtype", &G_param_spec_gtype},
	{"g_param_spec_int", &G_param_spec_int},
	{"g_param_spec_int64", &G_param_spec_int64},
	{"g_param_spec_internal", &G_param_spec_internal},
	{"g_param_spec_long", &G_param_spec_long},
	{"g_param_spec_object", &G_param_spec_object},
	{"g_param_spec_override", &G_param_spec_override},
	{"g_param_spec_param", &G_param_spec_param},
	{"g_param_spec_pointer", &G_param_spec_pointer},
	{"g_param_spec_pool_insert", &G_param_spec_pool_insert},
	{"g_param_spec_pool_list", &G_param_spec_pool_list},
	{"g_param_spec_pool_list_owned", &G_param_spec_pool_list_owned},
	{"g_param_spec_pool_lookup", &G_param_spec_pool_lookup},
	{"g_param_spec_pool_new", &G_param_spec_pool_new},
	{"g_param_spec_pool_remove", &G_param_spec_pool_remove},
	{"g_param_spec_ref", &G_param_spec_ref},
	{"g_param_spec_ref_sink", &G_param_spec_ref_sink},
	{"g_param_spec_set_qdata", &G_param_spec_set_qdata},
	{"g_param_spec_set_qdata_full", &G_param_spec_set_qdata_full},
	{"g_param_spec_sink", &G_param_spec_sink},
	{"g_param_spec_steal_qdata", &G_param_spec_steal_qdata},
	{"g_param_spec_string", &G_param_spec_string},
	{"g_param_spec_uchar", &G_param_spec_uchar},
	{"g_param_spec_uint", &G_param_spec_uint},
	{"g_param_spec_uint64", &G_param_spec_uint64},
	{"g_param_spec_ulong", &G_param_spec_ulong},
	{"g_param_spec_unichar", &G_param_spec_unichar},
	{"g_param_spec_unref", &G_param_spec_unref},
	{"g_param_spec_value_array", &G_param_spec_value_array},
	{"g_param_spec_variant", &G_param_spec_variant},
	{"g_param_type_register_static", &G_param_type_register_static},
	{"g_param_value_convert", &G_param_value_convert},
	{"g_param_value_defaults", &G_param_value_defaults},
	{"g_param_value_set_default", &G_param_value_set_default},
	{"g_param_value_validate", &G_param_value_validate},
	{"g_param_values_cmp", &G_param_values_cmp},
	{"g_pointer_type_register_static", &G_pointer_type_register_static},
	{"g_ptr_array_get_type", &G_ptr_array_get_type},
	{"g_regex_get_type", &G_regex_get_type},
	{"g_signal_accumulator_first_wins", &G_signal_accumulator_first_wins},
	{"g_signal_accumulator_true_handled", &G_signal_accumulator_true_handled},
	{"g_signal_add_emission_hook", &G_signal_add_emission_hook},
	{"g_signal_chain_from_overridden", &G_signal_chain_from_overridden},
	{"g_signal_chain_from_overridden_handler", &G_signal_chain_from_overridden_handler},
	{"g_signal_connect_closure", &G_signal_connect_closure},
	{"g_signal_connect_closure_by_id", &G_signal_connect_closure_by_id},
	{"g_signal_connect_data", &G_signal_connect_data},
	{"g_signal_connect_object", &G_signal_connect_object},
	{"g_signal_emit", &G_signal_emit},
	{"g_signal_emit_by_name", &G_signal_emit_by_name},
	{"g_signal_emit_valist", &G_signal_emit_valist},
	{"g_signal_emitv", &G_signal_emitv},
	{"g_signal_get_invocation_hint", &G_signal_get_invocation_hint},
	{"g_signal_handler_block", &G_signal_handler_block},
	{"g_signal_handler_disconnect", &G_signal_handler_disconnect},
	{"g_signal_handler_find", &G_signal_handler_find},
	{"g_signal_handler_is_connected", &G_signal_handler_is_connected},
	{"g_signal_handler_unblock", &G_signal_handler_unblock},
	{"g_signal_handlers_block_matched", &G_signal_handlers_block_matched},
	{"g_signal_handlers_destroy", &G_signal_handlers_destroy},
	{"g_signal_handlers_disconnect_matched", &G_signal_handlers_disconnect_matched},
	{"g_signal_handlers_unblock_matched", &G_signal_handlers_unblock_matched},
	{"g_signal_has_handler_pending", &G_signal_has_handler_pending},
	{"g_signal_list_ids", &G_signal_list_ids},
	{"g_signal_lookup", &G_signal_lookup},
	{"g_signal_name", &G_signal_name},
	{"g_signal_new", &G_signal_new},
	{"g_signal_new_class_handler", &G_signal_new_class_handler},
	{"g_signal_new_valist", &G_signal_new_valist},
	{"g_signal_newv", &G_signal_newv},
	{"g_signal_override_class_closure", &G_signal_override_class_closure},
	{"g_signal_override_class_handler", &G_signal_override_class_handler},
	{"g_signal_parse_name", &G_signal_parse_name},
	{"g_signal_query", &G_signal_query},
	{"g_signal_remove_emission_hook", &G_signal_remove_emission_hook},
	{"g_signal_stop_emission", &G_signal_stop_emission},
	{"g_signal_stop_emission_by_name", &G_signal_stop_emission_by_name},
	{"g_signal_type_cclosure_new", &G_signal_type_cclosure_new},
	{"g_slist_remove_all", &G_slist_remove_all},
	{"g_source_set_closure", &G_source_set_closure},
	{"g_source_set_dummy_callback", &G_source_set_dummy_callback},
	{"g_strdup_value_contents", &G_strdup_value_contents},
	{"g_strv_get_type", &G_strv_get_type},
	{"g_type_add_class_cache_func", &G_type_add_class_cache_func},
	{"g_type_add_class_private", &G_type_add_class_private},
	{"g_type_add_interface_check", &G_type_add_interface_check},
	{"g_type_add_interface_dynamic", &G_type_add_interface_dynamic},
	{"g_type_add_interface_static", &G_type_add_interface_static},
	{"g_type_check_class_cast", &G_type_check_class_cast},
	{"g_type_check_class_is_a", &G_type_check_class_is_a},
	{"g_type_check_instance", &G_type_check_instance},
	{"g_type_check_instance_cast", &G_type_check_instance_cast},
	{"g_type_check_instance_is_a", &G_type_check_instance_is_a},
	{"g_type_check_is_value_type", &G_type_check_is_value_type},
	{"g_type_check_value", &G_type_check_value},
	{"g_type_check_value_holds", &G_type_check_value_holds},
	{"g_type_children", &G_type_children},
	{"g_type_class_add_private", &G_type_class_add_private},
	{"g_type_class_get_private", &G_type_class_get_private},
	{"g_type_class_peek", &G_type_class_peek},
	{"g_type_class_peek_parent", &G_type_class_peek_parent},
	{"g_type_class_peek_static", &G_type_class_peek_static},
	{"g_type_class_ref", &G_type_class_ref},
	{"g_type_class_unref", &G_type_class_unref},
	{"g_type_class_unref_uncached", &G_type_class_unref_uncached},
	{"g_type_create_instance", &G_type_create_instance},
	{"g_type_default_interface_peek", &G_type_default_interface_peek},
	{"g_type_default_interface_ref", &G_type_default_interface_ref},
	{"g_type_default_interface_unref", &G_type_default_interface_unref},
	{"g_type_depth", &G_type_depth},
	{"g_type_free_instance", &G_type_free_instance},
	{"g_type_from_name", &G_type_from_name},
	{"g_type_fundamental", &G_type_fundamental},
	{"g_type_fundamental_next", &G_type_fundamental_next},
	{"g_type_get_plugin", &G_type_get_plugin},
	{"g_type_get_qdata", &G_type_get_qdata},
	{"g_type_init", &G_type_init},
	{"g_type_init_with_debug_flags", &G_type_init_with_debug_flags},
	{"g_type_instance_get_private", &G_type_instance_get_private},
	{"g_type_interface_add_prerequisite", &G_type_interface_add_prerequisite},
	{"g_type_interface_get_plugin", &G_type_interface_get_plugin},
	{"g_type_interface_peek", &G_type_interface_peek},
	{"g_type_interface_peek_parent", &G_type_interface_peek_parent},
	{"g_type_interface_prerequisites", &G_type_interface_prerequisites},
	{"g_type_interfaces", &G_type_interfaces},
	{"g_type_is_a", &G_type_is_a},
	{"g_type_module_add_interface", &G_type_module_add_interface},
	{"g_type_module_get_type", &G_type_module_get_type},
	{"g_type_module_register_enum", &G_type_module_register_enum},
	{"g_type_module_register_flags", &G_type_module_register_flags},
	{"g_type_module_register_type", &G_type_module_register_type},
	{"g_type_module_set_name", &G_type_module_set_name},
	{"g_type_module_unuse", &G_type_module_unuse},
	{"g_type_module_use", &G_type_module_use},
	{"g_type_name", &G_type_name},
	{"g_type_name_from_class", &G_type_name_from_class},
	{"g_type_name_from_instance", &G_type_name_from_instance},
	{"g_type_next_base", &G_type_next_base},
	{"g_type_parent", &G_type_parent},
	{"g_type_plugin_complete_interface_info", &G_type_plugin_complete_interface_info},
	{"g_type_plugin_complete_type_info", &G_type_plugin_complete_type_info},
	{"g_type_plugin_get_type", &G_type_plugin_get_type},
	{"g_type_plugin_unuse", &G_type_plugin_unuse},
	{"g_type_plugin_use", &G_type_plugin_use},
	{"g_type_qname", &G_type_qname},
	{"g_type_query", &G_type_query},
	{"g_type_register_dynamic", &G_type_register_dynamic},
	{"g_type_register_fundamental", &G_type_register_fundamental},
	{"g_type_register_static", &G_type_register_static},
	{"g_type_register_static_simple", &G_type_register_static_simple},
	{"g_type_remove_class_cache_func", &G_type_remove_class_cache_func},
	{"g_type_remove_interface_check", &G_type_remove_interface_check},
	{"g_type_set_qdata", &G_type_set_qdata},
	{"g_type_test_flags", &G_type_test_flags},
	{"g_type_value_table_peek", &G_type_value_table_peek},
	{"g_unichar_validate", &G_unichar_validate},
	{"g_value_array_append", &G_value_array_append},
	{"g_value_array_copy", &G_value_array_copy},
	{"g_value_array_free", &G_value_array_free},
	{"g_value_array_get_nth", &G_value_array_get_nth},
	{"g_value_array_get_type", &G_value_array_get_type},
	{"g_value_array_insert", &G_value_array_insert},
	{"g_value_array_new", &G_value_array_new},
	{"g_value_array_prepend", &G_value_array_prepend},
	{"g_value_array_remove", &G_value_array_remove},
	{"g_value_array_sort", &G_value_array_sort},
	{"g_value_array_sort_with_data", &G_value_array_sort_with_data},
	{"g_value_copy", &G_value_copy},
	{"g_value_dup_boxed", &G_value_dup_boxed},
	{"g_value_dup_object", &G_value_dup_object},
	{"g_value_dup_param", &G_value_dup_param},
	{"g_value_dup_string", &G_value_dup_string},
	{"g_value_dup_variant", &G_value_dup_variant},
	{"g_value_fits_pointer", &G_value_fits_pointer},
	{"g_value_get_boolean", &G_value_get_boolean},
	{"g_value_get_boxed", &G_value_get_boxed},
	{"g_value_get_char", &G_value_get_char},
	{"g_value_get_double", &G_value_get_double},
	{"g_value_get_enum", &G_value_get_enum},
	{"g_value_get_flags", &G_value_get_flags},
	{"g_value_get_float", &G_value_get_float},
	{"g_value_get_gtype", &G_value_get_gtype},
	{"g_value_get_int", &G_value_get_int},
	{"g_value_get_int64", &G_value_get_int64},
	{"g_value_get_long", &G_value_get_long},
	{"g_value_get_object", &G_value_get_object},
	{"g_value_get_param", &G_value_get_param},
	{"g_value_get_pointer", &G_value_get_pointer},
	{"g_value_get_string", &G_value_get_string},
	{"g_value_get_type", &G_value_get_type},
	{"g_value_get_uchar", &G_value_get_uchar},
	{"g_value_get_uint", &G_value_get_uint},
	{"g_value_get_uint64", &G_value_get_uint64},
	{"g_value_get_ulong", &G_value_get_ulong},
	{"g_value_get_variant", &G_value_get_variant},
	{"g_value_init", &G_value_init},
	{"g_value_peek_pointer", &G_value_peek_pointer},
	{"g_value_register_transform_func", &G_value_register_transform_func},
	{"g_value_reset", &G_value_reset},
	{"g_value_set_boolean", &G_value_set_boolean},
	{"g_value_set_boxed", &G_value_set_boxed},
	{"g_value_set_boxed_take_ownership", &G_value_set_boxed_take_ownership},
	{"g_value_set_char", &G_value_set_char},
	{"g_value_set_double", &G_value_set_double},
	{"g_value_set_enum", &G_value_set_enum},
	{"g_value_set_flags", &G_value_set_flags},
	{"g_value_set_float", &G_value_set_float},
	{"g_value_set_gtype", &G_value_set_gtype},
	{"g_value_set_instance", &G_value_set_instance},
	{"g_value_set_int", &G_value_set_int},
	{"g_value_set_int64", &G_value_set_int64},
	{"g_value_set_long", &G_value_set_long},
	{"g_value_set_object", &G_value_set_object},
	{"g_value_set_object_take_ownership", &G_value_set_object_take_ownership},
	{"g_value_set_param", &G_value_set_param},
	{"g_value_set_param_take_ownership", &G_value_set_param_take_ownership},
	{"g_value_set_pointer", &G_value_set_pointer},
	{"g_value_set_static_boxed", &G_value_set_static_boxed},
	{"g_value_set_static_string", &G_value_set_static_string},
	{"g_value_set_string", &G_value_set_string},
	{"g_value_set_string_take_ownership", &G_value_set_string_take_ownership},
	{"g_value_set_uchar", &G_value_set_uchar},
	{"g_value_set_uint", &G_value_set_uint},
	{"g_value_set_uint64", &G_value_set_uint64},
	{"g_value_set_ulong", &G_value_set_ulong},
	{"g_value_set_variant", &G_value_set_variant},
	{"g_value_take_boxed", &G_value_take_boxed},
	{"g_value_take_object", &G_value_take_object},
	{"g_value_take_param", &G_value_take_param},
	{"g_value_take_string", &G_value_take_string},
	{"g_value_take_variant", &G_value_take_variant},
	{"g_value_transform", &G_value_transform},
	{"g_value_type_compatible", &G_value_type_compatible},
	{"g_value_type_transformable", &G_value_type_transformable},
	{"g_value_unset", &G_value_unset},
	{"g_variant_get_gtype", &G_variant_get_gtype},
	{"g_variant_type_get_gtype", &G_variant_type_get_gtype},
}

var dataList = outside.Data{
	{"g_param_spec_types", (*T.GType)(nil)},
}
