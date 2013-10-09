// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package gobject provides API definitions for accessing
//libgobject-2.0-0.dll.
package gobject

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
	AddDllData(dll, false, dataList)
}

var (
	G_signal_newv func(
		signal_name string,
		itype GType,
		signal_flags GSignalFlags,
		class_closure *GClosure,
		accumulator GSignalAccumulator,
		accu_data Gpointer,
		c_marshaller GSignalCMarshaller,
		return_type GType,
		n_params Guint,
		param_types *GType) Guint

	G_signal_new_valist func(
		signal_name string,
		itype GType,
		signal_flags GSignalFlags,
		class_closure *GClosure,
		accumulator GSignalAccumulator,
		accu_data Gpointer,
		c_marshaller GSignalCMarshaller,
		return_type GType,
		n_params Guint,
		args Va_list) Guint

	G_signal_new func(signal_name string, itype GType,
		signal_flags GSignalFlags, class_offset Guint,
		accumulator GSignalAccumulator, accu_data Gpointer,
		c_marshaller GSignalCMarshaller, return_type GType,
		n_params Guint, v ...VArg) Guint

	G_signal_new_class_handler func(signal_name string,
		itype GType, signal_flags GSignalFlags,
		class_handler GCallback, accumulator GSignalAccumulator,
		accu_data Gpointer, c_marshaller GSignalCMarshaller,
		return_type GType, n_params Guint, v ...VArg) Guint

	G_signal_emitv func(
		instance_and_params *GValue,
		signal_id Guint,
		detail GQuark,
		return_value *GValue)

	G_signal_emit_valist func(
		instance Gpointer,
		signal_id Guint,
		detail GQuark,
		var_args Va_list)

	G_signal_emit func(instance Gpointer, signal_id Guint,
		detail GQuark, v ...VArg)

	G_signal_emit_by_name func(instance Gpointer,
		detailed_signal string, v ...VArg)

	G_signal_lookup func(name string, itype GType) Guint

	G_signal_name func(signal_id Guint) string

	G_signal_query func(signal_id Guint, query *GSignalQuery)

	G_signal_list_ids func(itype GType, n_ids *Guint) *Guint

	G_signal_parse_name func(
		detailed_signal string,
		itype GType,
		signal_id_p *Guint,
		detail_p *GQuark,
		force_detail_quark Gboolean) Gboolean

	G_signal_get_invocation_hint func(
		instance Gpointer) *GSignalInvocationHint

	G_signal_stop_emission func(
		instance Gpointer,
		signal_id Guint,
		detail GQuark)

	G_signal_stop_emission_by_name func(
		instance Gpointer,
		detailed_signal string)

	G_signal_add_emission_hook func(
		signal_id Guint,
		detail GQuark,
		hook_func GSignalEmissionHook,
		hook_data Gpointer,
		data_destroy GDestroyNotify) Gulong

	G_signal_remove_emission_hook func(
		signal_id Guint,
		hook_id Gulong)

	G_signal_has_handler_pending func(
		instance Gpointer,
		signal_id Guint,
		detail GQuark,
		may_be_blocked Gboolean) Gboolean

	G_signal_connect_closure_by_id func(
		instance Gpointer,
		signal_id Guint,
		detail GQuark,
		closure *GClosure,
		after Gboolean) Gulong

	G_signal_connect_closure func(
		instance Gpointer,
		detailed_signal string,
		closure *GClosure,
		after Gboolean) Gulong

	G_signal_connect_data func(
		instance Gpointer,
		detailed_signal string,
		c_handler GCallback,
		data Gpointer,
		destroy_data GClosureNotify,
		connect_flags GConnectFlags) Gulong

	G_signal_handler_block func(
		instance Gpointer,
		handler_id Gulong)

	G_signal_handler_unblock func(
		instance Gpointer,
		handler_id Gulong)

	G_signal_handler_disconnect func(
		instance Gpointer,
		handler_id Gulong)

	G_signal_handler_is_connected func(
		instance Gpointer,
		handler_id Gulong) Gboolean

	G_signal_handler_find func(
		instance Gpointer,
		mask GSignalMatchType,
		signal_id Guint,
		detail GQuark,
		closure *GClosure,
		fnc Gpointer,
		data Gpointer) Gulong

	G_signal_handlers_block_matched func(
		instance Gpointer,
		mask GSignalMatchType,
		signal_id Guint,
		detail GQuark,
		closure *GClosure,
		fnc Gpointer,
		data Gpointer) Guint

	G_signal_handlers_unblock_matched func(
		instance Gpointer,
		mask GSignalMatchType,
		signal_id Guint,
		detail GQuark,
		closure *GClosure,
		fnc Gpointer,
		data Gpointer) Guint

	G_signal_handlers_disconnect_matched func(
		instance Gpointer,
		mask GSignalMatchType,
		signal_id Guint,
		detail GQuark,
		closure *GClosure,
		fnc Gpointer,
		data Gpointer) Guint

	G_signal_override_class_closure func(
		signal_id Guint,
		instance_type GType,
		class_closure *GClosure)

	G_signal_override_class_handler func(
		signal_name string,
		instance_type GType,
		class_handler GCallback)

	G_signal_chain_from_overridden func(
		instance_and_params *GValue,
		return_value *GValue)

	G_signal_chain_from_overridden_handler func(
		instance Gpointer, v ...VArg)

	G_signal_accumulator_true_handled func(
		ihint *GSignalInvocationHint,
		return_accu *GValue,
		handler_return *GValue,
		dummy Gpointer) Gboolean

	G_signal_accumulator_first_wins func(
		ihint *GSignalInvocationHint,
		return_accu *GValue,
		handler_return *GValue,
		dummy Gpointer) Gboolean

	G_signal_handlers_destroy func(
		instance Gpointer)

	G_initially_unowned_get_type func() GType

	G_object_class_install_property func(
		oclass *GObjectClass,
		property_id Guint,
		pspec *GParamSpec)

	G_object_class_find_property func(
		oclass *GObjectClass,
		property_name string) *GParamSpec

	G_object_class_list_properties func(
		oclass *GObjectClass,
		n_properties *Guint) **GParamSpec

	G_object_class_override_property func(
		oclass *GObjectClass,
		property_id Guint,
		name string)

	G_object_class_install_properties func(
		oclass *GObjectClass,
		n_pspecs Guint,
		pspecs **GParamSpec)

	G_object_interface_install_property func(
		g_iface Gpointer,
		pspec *GParamSpec)

	G_object_interface_find_property func(
		g_iface Gpointer,
		property_name string) *GParamSpec

	G_object_interface_list_properties func(
		g_iface Gpointer,
		n_properties_p *Guint) **GParamSpec

	G_object_get_type func() GType

	G_object_new func(object_type GType,
		first_property_name string, v ...VArg) Gpointer

	G_object_newv func(
		object_type GType,
		n_parameters Guint,
		parameters *GParameter) Gpointer

	G_object_new_valist func(
		object_type GType,
		first_property_name string,
		var_args Va_list) *GObject

	G_object_set func(object Gpointer,
		first_property_name string, v ...VArg)

	G_object_get func(object Gpointer,
		first_property_name string, v ...VArg)

	G_object_connect func(object Gpointer,
		signal_spec string, v ...VArg) Gpointer

	G_object_disconnect func(object Gpointer,
		signal_spec string, v ...VArg)

	G_object_set_valist func(
		object *GObject,
		first_property_name string,
		var_args Va_list)

	G_object_get_valist func(
		object *GObject,
		first_property_name string,
		var_args Va_list)

	G_object_set_property func(
		object *GObject,
		property_name string,
		value *GValue)

	G_object_get_property func(
		object *GObject,
		property_name string,
		value *GValue)

	G_object_freeze_notify func(
		object *GObject)

	G_object_notify func(
		object *GObject,
		property_name string)

	G_object_notify_by_pspec func(
		object *GObject,
		pspec *GParamSpec)

	G_object_thaw_notify func(
		object *GObject)

	G_object_is_floating func(
		object Gpointer) Gboolean

	G_object_ref_sink func(
		object Gpointer) Gpointer

	G_object_ref func(
		object Gpointer) Gpointer

	G_object_unref func(
		object Gpointer)

	G_object_weak_ref func(
		object *GObject,
		notify GWeakNotify,
		data Gpointer)

	G_object_weak_unref func(
		object *GObject,
		notify GWeakNotify,
		data Gpointer)

	G_object_add_weak_pointer func(
		object *GObject,
		weak_pointer_location *Gpointer)

	G_object_remove_weak_pointer func(
		object *GObject,
		weak_pointer_location *Gpointer)

	G_object_add_toggle_ref func(
		object *GObject,
		notify GToggleNotify,
		data Gpointer)

	G_object_remove_toggle_ref func(
		object *GObject,
		notify GToggleNotify,
		data Gpointer)

	G_object_get_qdata func(
		object *GObject,
		quark GQuark) Gpointer

	G_object_set_qdata func(
		object *GObject,
		quark GQuark,
		data Gpointer)

	G_object_set_qdata_full func(
		object *GObject,
		quark GQuark,
		data Gpointer,
		destroy GDestroyNotify)

	G_object_steal_qdata func(
		object *GObject,
		quark GQuark) Gpointer

	G_object_get_data func(
		object *GObject,
		key string) Gpointer

	G_object_set_data func(
		object *GObject,
		key string,
		data Gpointer)

	G_object_set_data_full func(
		object *GObject,
		key string,
		data Gpointer,
		destroy GDestroyNotify)

	G_object_steal_data func(
		object *GObject,
		key string) Gpointer

	G_object_watch_closure func(
		object *GObject,
		closure *GClosure)

	G_cclosure_new_object func(
		callback_func GCallback,
		object *GObject) *GClosure

	G_cclosure_new_object_swap func(
		callback_func GCallback,
		object *GObject) *GClosure

	G_closure_new_object func(
		sizeof_closure Guint,
		object *GObject) *GClosure

	G_value_set_object func(
		value *GValue,
		v_object Gpointer)

	G_value_get_object func(
		value *GValue) Gpointer

	G_value_dup_object func(
		value *GValue) Gpointer

	G_signal_connect_object func(
		instance Gpointer,
		detailed_signal string,
		c_handler GCallback,
		gobject Gpointer,
		connect_flags GConnectFlags) Gulong

	G_object_force_floating func(
		object *GObject)

	G_object_run_dispose func(
		object *GObject)

	G_value_take_object func(
		value *GValue,
		v_object Gpointer)

	G_value_set_object_take_ownership func(
		value *GValue,
		v_object Gpointer)

	G_object_compat_control func(
		what Gsize,
		data Gpointer) Gsize

	G_clear_object func(object_ptr **GObject)

	G_binding_flags_get_type func() GType

	G_binding_get_type func() GType

	G_binding_get_flags func(binding *GBinding) GBindingFlags

	G_binding_get_source func(binding *GBinding) *GObject

	G_binding_get_target func(binding *GBinding) *GObject

	G_binding_get_source_property func(binding *GBinding) string

	G_binding_get_target_property func(binding *GBinding) string

	G_object_bind_property func(
		source Gpointer,
		source_property string,
		target Gpointer,
		target_property string,
		flags GBindingFlags) *GBinding

	G_object_bind_property_full func(
		source Gpointer,
		source_property string,
		target Gpointer,
		target_property string,
		flags GBindingFlags,
		transform_to GBindingTransformFunc,
		transform_from GBindingTransformFunc,
		user_data Gpointer,
		notify GDestroyNotify) *GBinding

	G_object_bind_property_with_closures func(
		source Gpointer,
		source_property string,
		target Gpointer,
		target_property string,
		flags GBindingFlags,
		transform_to *GClosure,
		transform_from *GClosure) *GBinding

	G_boxed_copy func(
		boxed_type GType,
		src_boxed Gconstpointer) Gpointer

	G_boxed_free func(
		boxed_type GType,
		boxed Gpointer)

	G_value_set_boxed func(
		value *GValue,
		v_boxed Gconstpointer)

	G_value_set_static_boxed func(
		value *GValue,
		v_boxed Gconstpointer)

	G_value_get_boxed func(
		value *GValue) Gpointer

	G_value_dup_boxed func(
		value *GValue) Gpointer

	G_boxed_type_register_static func(
		name string,
		boxed_copy GBoxedCopyFunc,
		boxed_free GBoxedFreeFunc) GType

	G_value_take_boxed func(
		value *GValue,
		v_boxed Gconstpointer)

	G_value_set_boxed_take_ownership func(
		value *GValue,
		v_boxed Gconstpointer)

	G_closure_get_type func() GType

	G_value_get_type func() GType

	G_value_array_get_type func() GType

	G_date_get_type func() GType

	G_strv_get_type func() GType

	G_gstring_get_type func() GType

	G_hash_table_get_type func() GType

	G_array_get_type func() GType

	G_byte_array_get_type func() GType

	G_ptr_array_get_type func() GType

	G_variant_type_get_gtype func() GType

	G_regex_get_type func() GType

	G_error_get_type func() GType

	G_date_time_get_type func() GType

	G_variant_get_gtype func() GType

	G_type_init func()

	G_type_init_with_debug_flags func(
		debug_flags GTypeDebugFlags)

	G_type_name func(t GType) string

	G_type_qname func(t GType) GQuark

	G_type_from_name func(name string) GType

	G_type_parent func(t GType) GType

	G_type_depth func(t GType) Guint

	G_type_next_base func(leaf_type, root_type GType) GType

	G_type_is_a func(t, is_a_type GType) Gboolean

	G_type_class_ref func(
		t GType) Gpointer

	G_type_class_peek func(
		t GType) Gpointer

	G_type_class_peek_static func(
		t GType) Gpointer

	G_type_class_unref func(
		g_class Gpointer)

	G_type_class_peek_parent func(
		g_class Gpointer) Gpointer

	G_type_interface_peek func(
		instance_class Gpointer,
		iface_type GType) Gpointer

	G_type_interface_peek_parent func(
		g_iface Gpointer) Gpointer

	G_type_default_interface_ref func(
		g_type GType) Gpointer

	G_type_default_interface_peek func(
		g_type GType) Gpointer

	G_type_default_interface_unref func(
		g_iface Gpointer)

	G_type_children func(
		t GType,
		n_children *Guint) *GType

	G_type_interfaces func(
		t GType,
		n_interfaces *Guint) *GType

	G_type_set_qdata func(
		t GType,
		quark GQuark,
		data Gpointer)

	G_type_get_qdata func(
		t GType,
		quark GQuark) Gpointer

	G_type_query func(
		t GType,
		query *GTypeQuery)

	G_cclosure_new func(
		callback_func GCallback,
		user_data Gpointer,
		destroy_data GClosureNotify) *GClosure

	G_cclosure_new_swap func(
		callback_func GCallback,
		user_data Gpointer,
		destroy_data GClosureNotify) *GClosure

	G_signal_type_cclosure_new func(
		itype GType,
		struct_offset Guint) *GClosure

	G_closure_ref func(
		closure *GClosure) *GClosure

	G_closure_sink func(
		closure *GClosure)

	G_closure_unref func(
		closure *GClosure)

	G_closure_new_simple func(
		sizeof_closure Guint,
		data Gpointer) *GClosure

	G_closure_add_finalize_notifier func(
		closure *GClosure,
		notify_data Gpointer,
		notify_func GClosureNotify)

	G_closure_remove_finalize_notifier func(
		closure *GClosure,
		notify_data Gpointer,
		notify_func GClosureNotify)

	G_closure_add_invalidate_notifier func(
		closure *GClosure,
		notify_data Gpointer,
		notify_func GClosureNotify)

	G_closure_remove_invalidate_notifier func(
		closure *GClosure,
		notify_data Gpointer,
		notify_func GClosureNotify)

	G_closure_add_marshal_guards func(
		closure *GClosure,
		pre_marshal_data Gpointer,
		pre_marshal_notify GClosureNotify,
		post_marshal_data Gpointer,
		post_marshal_notify GClosureNotify)

	G_closure_set_marshal func(
		closure *GClosure,
		marshal GClosureMarshal)

	G_closure_set_meta_marshal func(
		closure *GClosure,
		marshal_data Gpointer,
		meta_marshal GClosureMarshal)

	G_closure_invalidate func(
		closure *GClosure)

	G_closure_invoke func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer)

	G_cclosure_marshal_VOID__VOID func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__BOOLEAN func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__CHAR func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__UCHAR func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__INT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__UINT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__LONG func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__ULONG func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__ENUM func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__FLAGS func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__FLOAT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__DOUBLE func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__STRING func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__PARAM func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__BOXED func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__OBJECT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__VARIANT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_VOID__UINT_POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_BOOLEAN__FLAGS func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_STRING__OBJECT_POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_cclosure_marshal_BOOLEAN__BOXED_BOXED func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	G_enum_get_value func(
		enum_class *GEnumClass,
		value Gint) *GEnumValue

	G_enum_get_value_by_name func(
		enum_class *GEnumClass,
		name string) *GEnumValue

	G_enum_get_value_by_nick func(
		enum_class *GEnumClass,
		nick string) *GEnumValue

	G_flags_get_first_value func(
		flags_class *GFlagsClass,
		value Guint) *GFlagsValue

	G_flags_get_value_by_name func(
		flags_class *GFlagsClass,
		name string) *GFlagsValue

	G_flags_get_value_by_nick func(
		flags_class *GFlagsClass,
		nick string) *GFlagsValue

	G_value_set_enum func(
		value *GValue,
		v_enum Gint)

	G_value_get_enum func(
		value *GValue) Gint

	G_value_set_flags func(
		value *GValue,
		v_flags Guint)

	G_value_get_flags func(
		value *GValue) Guint

	G_enum_register_static func(
		name string,
		const_static_values *GEnumValue) GType

	G_flags_register_static func(
		name string,
		const_static_values *GFlagsValue) GType

	G_enum_complete_type_info func(
		g_enum_type GType,
		info *GTypeInfo,
		const_values *GEnumValue)

	G_flags_complete_type_info func(
		g_flags_type GType,
		info *GTypeInfo,
		const_values *GFlagsValue)

	G_param_spec_ref func(
		pspec *GParamSpec) *GParamSpec

	G_param_spec_unref func(
		pspec *GParamSpec)

	G_param_spec_sink func(
		pspec *GParamSpec)

	G_param_spec_ref_sink func(
		pspec *GParamSpec) *GParamSpec

	G_param_spec_get_qdata func(
		pspec *GParamSpec,
		quark GQuark) Gpointer

	G_param_spec_set_qdata func(
		pspec *GParamSpec,
		quark GQuark,
		data Gpointer)

	G_param_spec_set_qdata_full func(
		pspec *GParamSpec,
		quark GQuark,
		data Gpointer,
		destroy GDestroyNotify)

	G_param_spec_steal_qdata func(
		pspec *GParamSpec,
		quark GQuark) Gpointer

	G_param_spec_get_redirect_target func(
		pspec *GParamSpec) *GParamSpec

	G_param_value_set_default func(
		pspec *GParamSpec,
		value *GValue)

	G_param_value_defaults func(
		pspec *GParamSpec,
		value *GValue) Gboolean

	G_param_value_validate func(
		pspec *GParamSpec,
		value *GValue) Gboolean

	G_param_value_convert func(
		pspec *GParamSpec,
		src_value *GValue,
		dest_value *GValue,
		strict_validation Gboolean) Gboolean

	G_param_values_cmp func(
		pspec *GParamSpec,
		value1 *GValue,
		value2 *GValue) Gint

	G_param_spec_get_name func(
		pspec *GParamSpec) string

	G_param_spec_get_nick func(
		pspec *GParamSpec) string

	G_param_spec_get_blurb func(
		pspec *GParamSpec) string

	G_value_set_param func(
		value *GValue,
		param *GParamSpec)

	G_value_get_param func(
		value *GValue) *GParamSpec

	G_value_dup_param func(
		value *GValue) *GParamSpec

	G_value_take_param func(
		value *GValue,
		param *GParamSpec)

	G_value_set_param_take_ownership func(
		value *GValue,
		param *GParamSpec)

	G_value_array_get_nth func(
		value_array *GValueArray,
		index_ Guint) *GValue

	G_value_array_new func(n_prealloced Guint) *GValueArray

	G_value_array_free func(value_array *GValueArray)

	G_value_array_copy func(
		value_array *GValueArray) *GValueArray

	G_value_array_prepend func(
		value_array *GValueArray,
		value *GValue) *GValueArray

	G_value_array_append func(
		value_array *GValueArray,
		value *GValue) *GValueArray

	G_value_array_insert func(
		value_array *GValueArray,
		index_ Guint,
		value *GValue) *GValueArray

	G_value_array_remove func(
		value_array *GValueArray,
		index_ Guint) *GValueArray

	G_value_array_sort func(
		value_array *GValueArray,
		compare_func GCompareFunc) *GValueArray

	G_value_array_sort_with_data func(
		value_array *GValueArray,
		compare_func GCompareDataFunc,
		user_data Gpointer) *GValueArray

	G_value_set_char func(value *GValue, v_char Gchar)

	G_value_get_char func(value *GValue) Gchar

	G_value_set_uchar func(
		value *GValue, v_uchar Guchar)

	G_value_get_uchar func(value *GValue) Guchar

	G_value_set_boolean func(
		value *GValue, v_boolean Gboolean)

	G_value_get_boolean func(value *GValue) Gboolean

	G_value_set_int func(value *GValue, v_int Gint)

	G_value_get_int func(value *GValue) Gint

	G_value_set_uint func(value *GValue, v_uint Guint)

	G_value_get_uint func(value *GValue) Guint

	G_value_set_long func(value *GValue, v_long Glong)

	G_value_get_long func(value *GValue) Glong

	G_value_set_ulong func(value *GValue, v_ulong Gulong)

	G_value_get_ulong func(value *GValue) Gulong

	G_value_set_int64 func(value *GValue, v_int64 Gint64)

	G_value_get_int64 func(value *GValue) Gint64

	G_value_set_uint64 func(value *GValue, v_uint64 Guint64)

	G_value_get_uint64 func(value *GValue) Guint64

	G_value_set_float func(value *GValue, v_float Gfloat)

	G_value_get_float func(value *GValue) Gfloat

	G_value_set_double func(value *GValue, v_double Gdouble)

	G_value_get_double func(value *GValue) Gdouble

	G_value_set_string func(value *GValue, v_string string)

	G_value_set_static_string func(
		value *GValue, v_string string)

	G_value_get_string func(value *GValue) string

	G_value_dup_string func(value *GValue) string

	G_value_set_pointer func(
		value *GValue, v_pointer Gpointer)

	G_value_get_pointer func(value *GValue) Gpointer

	G_gtype_get_type func() GType

	G_value_set_gtype func(value *GValue, v_gtype GType)

	G_value_get_gtype func(value *GValue) GType

	G_value_set_variant func(
		value *GValue, variant *GVariant)

	G_value_take_variant func(
		value *GValue, variant *GVariant)

	G_value_get_variant func(value *GValue) *GVariant

	G_value_dup_variant func(value *GValue) *GVariant

	G_pointer_type_register_static func(name string) GType

	G_strdup_value_contents func(value *GValue) string

	G_value_take_string func(
		value *GValue, v_string string)

	G_value_set_string_take_ownership func(
		value *GValue, v_string string)

	G_type_register_static func(
		parent_type GType,
		type_name string,
		info *GTypeInfo,
		flags GTypeFlags) GType

	G_type_register_static_simple func(
		parent_type GType,
		type_name string,
		class_size Guint,
		class_init GClassInitFunc,
		instance_size Guint,
		instance_init GInstanceInitFunc,
		flags GTypeFlags) GType

	G_type_register_dynamic func(
		parent_type GType,
		type_name string,
		plugin *GTypePlugin,
		flags GTypeFlags) GType

	G_type_register_fundamental func(
		type_id GType,
		type_name string,
		info *GTypeInfo,
		finfo *GTypeFundamentalInfo,
		flags GTypeFlags) GType

	G_type_add_interface_static func(
		instance_type GType,
		interface_type GType,
		info *GInterfaceInfo)

	G_type_add_interface_dynamic func(
		instance_type GType,
		interface_type GType,
		plugin *GTypePlugin)

	G_type_interface_add_prerequisite func(
		interface_type GType,
		prerequisite_type GType)

	G_type_interface_prerequisites func(
		interface_type GType,
		n_prerequisites *Guint) *GType

	G_type_class_add_private func(
		g_class Gpointer,
		private_size Gsize)

	G_type_instance_get_private func(
		instance *GTypeInstance,
		private_type GType) Gpointer

	G_type_add_class_private func(
		class_type GType,
		private_size Gsize)

	G_type_class_get_private func(
		klass *GTypeClass,
		private_type GType) Gpointer

	G_type_get_plugin func(
		t GType) *GTypePlugin

	G_type_interface_get_plugin func(
		instance_type GType,
		interface_type GType) *GTypePlugin

	G_type_fundamental_next func() GType

	G_type_fundamental func(
		type_id GType) GType

	G_type_create_instance func(
		t GType) *GTypeInstance

	G_type_free_instance func(
		instance *GTypeInstance)

	G_type_add_class_cache_func func(
		cache_data Gpointer,
		cache_func GTypeClassCacheFunc)

	G_type_remove_class_cache_func func(
		cache_data Gpointer,
		cache_func GTypeClassCacheFunc)

	G_type_class_unref_uncached func(
		g_class Gpointer)

	G_type_add_interface_check func(
		check_data Gpointer,
		check_func GTypeInterfaceCheckFunc)

	G_type_remove_interface_check func(
		check_data Gpointer,
		check_func GTypeInterfaceCheckFunc)

	G_type_value_table_peek func(
		t GType) *GTypeValueTable

	G_type_check_instance func(
		instance *GTypeInstance) Gboolean

	G_type_check_instance_cast func(
		instance *GTypeInstance,
		iface_type GType) *GTypeInstance

	G_type_check_instance_is_a func(
		instance *GTypeInstance,
		iface_type GType) Gboolean

	G_type_check_class_cast func(
		g_class *GTypeClass,
		is_a_type GType) *GTypeClass

	G_type_check_class_is_a func(
		g_class *GTypeClass,
		is_a_type GType) Gboolean

	G_type_check_is_value_type func(
		t GType) Gboolean

	G_type_check_value func(
		value *GValue) Gboolean

	G_type_check_value_holds func(
		value *GValue,
		t GType) Gboolean

	G_type_test_flags func(
		t GType,
		flags Guint) Gboolean

	G_type_name_from_instance func(
		instance *GTypeInstance) string

	G_type_name_from_class func(
		g_class *GTypeClass) string

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
		value *GValue,
		g_type GType) *GValue

	G_value_copy func(
		src_value *GValue,
		dest_value *GValue)

	G_value_reset func(
		value *GValue) *GValue

	G_value_unset func(
		value *GValue)

	G_value_set_instance func(
		value *GValue,
		instance Gpointer)

	G_value_fits_pointer func(
		value *GValue) Gboolean

	G_value_peek_pointer func(
		value *GValue) Gpointer

	G_value_type_compatible func(
		src_type GType,
		dest_type GType) Gboolean

	G_value_type_transformable func(
		src_type GType,
		dest_type GType) Gboolean

	G_value_transform func(
		src_value *GValue,
		dest_value *GValue) Gboolean

	G_value_register_transform_func func(
		src_type GType,
		dest_type GType,
		transform_func GValueTransform)

	G_type_plugin_get_type func() GType

	G_type_plugin_use func(
		plugin *GTypePlugin)

	G_type_plugin_unuse func(
		plugin *GTypePlugin)

	G_type_plugin_complete_type_info func(
		plugin *GTypePlugin,
		g_type GType,
		info *GTypeInfo,
		value_table *GTypeValueTable)

	G_type_plugin_complete_interface_info func(
		plugin *GTypePlugin,
		instance_type GType,
		interface_type GType,
		info *GInterfaceInfo)

	G_type_module_get_type func() GType

	G_type_module_use func(
		module *GTypeModule) Gboolean

	G_type_module_unuse func(
		module *GTypeModule)

	G_type_module_set_name func(
		module *GTypeModule,
		name string)

	G_type_module_register_type func(
		module *GTypeModule,
		parent_type GType,
		type_name string,
		type_info *GTypeInfo,
		flags GTypeFlags) GType

	G_type_module_add_interface func(
		module *GTypeModule,
		instance_type GType,
		interface_type GType,
		interface_info *GInterfaceInfo)

	G_type_module_register_enum func(
		module *GTypeModule,
		name string,
		const_static_values *GEnumValue) GType

	G_type_module_register_flags func(
		module *GTypeModule,
		name string,
		const_static_values *GFlagsValue) GType

	G_param_spec_char func(
		name string,
		nick string,
		blurb string,
		minimum Gint8,
		maximum Gint8,
		default_value Gint8,
		flags GParamFlags) *GParamSpec

	G_param_spec_uchar func(
		name string,
		nick string,
		blurb string,
		minimum Guint8,
		maximum Guint8,
		default_value Guint8,
		flags GParamFlags) *GParamSpec

	G_param_spec_boolean func(
		name string,
		nick string,
		blurb string,
		default_value Gboolean,
		flags GParamFlags) *GParamSpec

	G_param_spec_int func(
		name string,
		nick string,
		blurb string,
		minimum Gint,
		maximum Gint,
		default_value Gint,
		flags GParamFlags) *GParamSpec

	G_param_spec_uint func(
		name string,
		nick string,
		blurb string,
		minimum Guint,
		maximum Guint,
		default_value Guint,
		flags GParamFlags) *GParamSpec

	G_param_spec_long func(
		name string,
		nick string,
		blurb string,
		minimum Glong,
		maximum Glong,
		default_value Glong,
		flags GParamFlags) *GParamSpec

	G_param_spec_ulong func(
		name string,
		nick string,
		blurb string,
		minimum Gulong,
		maximum Gulong,
		default_value Gulong,
		flags GParamFlags) *GParamSpec

	G_param_spec_int64 func(
		name string,
		nick string,
		blurb string,
		minimum Gint64,
		maximum Gint64,
		default_value Gint64,
		flags GParamFlags) *GParamSpec

	G_param_spec_uint64 func(
		name string,
		nick string,
		blurb string,
		minimum Guint64,
		maximum Guint64,
		default_value Guint64,
		flags GParamFlags) *GParamSpec

	G_param_spec_unichar func(
		name string,
		nick string,
		blurb string,
		default_value Gunichar,
		flags GParamFlags) *GParamSpec

	G_param_spec_enum func(
		name string,
		nick string,
		blurb string,
		enum_type GType,
		default_value Gint,
		flags GParamFlags) *GParamSpec

	G_param_spec_flags func(
		name string,
		nick string,
		blurb string,
		flags_type GType,
		default_value Guint,
		flags GParamFlags) *GParamSpec

	G_param_spec_float func(
		name string,
		nick string,
		blurb string,
		minimum Gfloat,
		maximum Gfloat,
		default_value Gfloat,
		flags GParamFlags) *GParamSpec

	G_param_spec_double func(
		name string,
		nick string,
		blurb string,
		minimum Gdouble,
		maximum Gdouble,
		default_value Gdouble,
		flags GParamFlags) *GParamSpec

	G_param_spec_string func(
		name string,
		nick string,
		blurb string,
		default_value string,
		flags GParamFlags) *GParamSpec

	G_param_spec_param func(
		name string,
		nick string,
		blurb string,
		param_type GType,
		flags GParamFlags) *GParamSpec

	G_param_spec_boxed func(
		name string,
		nick string,
		blurb string,
		boxed_type GType,
		flags GParamFlags) *GParamSpec

	G_param_spec_pointer func(
		name string,
		nick string,
		blurb string,
		flags GParamFlags) *GParamSpec

	G_param_spec_value_array func(
		name string,
		nick string,
		blurb string,
		element_spec *GParamSpec,
		flags GParamFlags) *GParamSpec

	G_param_spec_object func(
		name string,
		nick string,
		blurb string,
		object_type GType,
		flags GParamFlags) *GParamSpec

	G_param_spec_override func(
		name string,
		overridden *GParamSpec) *GParamSpec

	G_param_spec_gtype func(
		name string,
		nick string,
		blurb string,
		is_a_type GType,
		flags GParamFlags) *GParamSpec

	G_param_spec_variant func(
		name string,
		nick string,
		blurb string,
		t *GVariantType,
		default_value *GVariant,
		flags GParamFlags) *GParamSpec

	G_source_set_closure func(
		source *GSource,
		closure *GClosure)

	G_source_set_dummy_callback func(
		source *GSource)

	G_io_channel_get_type func() GType

	G_io_condition_get_type func() GType

	G_param_type_register_static func(
		name string,
		pspec_info *GParamSpecTypeInfo) GType

	G_param_spec_internal func(
		param_type GType,
		name string,
		nick string,
		blurb string,
		flags GParamFlags) Gpointer

	G_param_spec_pool_new func(
		type_prefixing Gboolean) *GParamSpecPool

	G_param_spec_pool_insert func(
		pool *GParamSpecPool,
		pspec *GParamSpec,
		owner_type GType)

	G_param_spec_pool_remove func(
		pool *GParamSpecPool,
		pspec *GParamSpec)

	G_param_spec_pool_lookup func(
		pool *GParamSpecPool,
		param_name string,
		owner_type GType,
		walk_ancestors Gboolean) *GParamSpec

	G_param_spec_pool_list_owned func(
		pool *GParamSpecPool,
		owner_type GType) *GList

	G_param_spec_pool_list func(
		pool *GParamSpecPool,
		owner_type GType,
		n_pspecs_p *Guint) **GParamSpec

	G_unichar_validate func(
		ch Gunichar) Gboolean

	G_slist_remove_all func(
		list *GSList,
		data Gconstpointer) *GSList
)

var dll = "libgobject-2.0-0.dll"

var apiList = Apis{
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

var dataList = Data{
// {"g_param_spec_types", new(G_param_spec_types)},
}
