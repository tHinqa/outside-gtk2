// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gio provides API definitions for accessing
//libgio-2.0-0.dll.
package gio

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
}

var (
	G_app_info_get_type func() T.GType

	G_app_info_create_from_commandline func(
		commandline string,
		application_name string,
		flags T.GAppInfoCreateFlags,
		err **T.GError) *T.GAppInfo

	G_app_info_dup func(appinfo *T.GAppInfo) *T.GAppInfo

	G_app_info_equal func(
		appinfo1 *T.GAppInfo, appinfo2 *T.GAppInfo) T.Gboolean

	G_app_info_get_id func(appinfo *T.GAppInfo) string

	G_app_info_get_name func(appinfo *T.GAppInfo) string

	G_app_info_get_display_name func(appinfo *T.GAppInfo) string

	G_app_info_get_description func(appinfo *T.GAppInfo) string

	G_app_info_get_executable func(appinfo *T.GAppInfo) string

	G_app_info_get_commandline func(appinfo *T.GAppInfo) string

	G_app_info_get_icon func(appinfo *T.GAppInfo) *T.GIcon

	G_app_info_launch func(
		appinfo *T.GAppInfo,
		files *T.GList,
		launch_context *T.GAppLaunchContext,
		err **T.GError) T.Gboolean

	G_app_info_supports_uris func(appinfo *T.GAppInfo) T.Gboolean

	G_app_info_supports_files func(appinfo *T.GAppInfo) T.Gboolean

	G_app_info_launch_uris func(
		appinfo *T.GAppInfo,
		uris *T.GList,
		launch_context *T.GAppLaunchContext,
		err **T.GError) T.Gboolean

	G_app_info_should_show func(appinfo *T.GAppInfo) T.Gboolean

	G_app_info_set_as_default_for_type func(
		appinfo *T.GAppInfo,
		content_type string,
		err **T.GError) T.Gboolean

	G_app_info_set_as_default_for_extension func(
		appinfo *T.GAppInfo,
		extension string,
		err **T.GError) T.Gboolean

	G_app_info_add_supports_type func(
		appinfo *T.GAppInfo,
		content_type string,
		err **T.GError) T.Gboolean

	G_app_info_can_remove_supports_type func(
		appinfo *T.GAppInfo) T.Gboolean

	G_app_info_remove_supports_type func(
		appinfo *T.GAppInfo,
		content_type string,
		err **T.GError) T.Gboolean

	G_app_info_can_delete func(appinfo *T.GAppInfo) T.Gboolean

	G_app_info_delete func(appinfo *T.GAppInfo) T.Gboolean

	G_app_info_set_as_last_used_for_type func(
		appinfo *T.GAppInfo,
		content_type string,
		err **T.GError) T.Gboolean

	G_app_info_get_all func() *T.GList

	G_app_info_get_all_for_type func(content_type string) *T.GList

	G_app_info_get_recommended_for_type func(
		content_type string) *T.GList

	G_app_info_get_fallback_for_type func(
		content_type string) *T.GList

	G_app_info_reset_type_associations func(content_type string)

	G_app_info_get_default_for_type func(
		content_type string,
		must_support_uris T.Gboolean) *T.GAppInfo

	G_app_info_get_default_for_uri_scheme func(
		uri_scheme string) *T.GAppInfo

	G_app_info_launch_default_for_uri func(
		uri string,
		launch_context *T.GAppLaunchContext,
		err **T.GError) T.Gboolean

	G_app_launch_context_get_type func() T.GType

	G_app_launch_context_new func() *T.GAppLaunchContext

	G_app_launch_context_get_display func(
		context *T.GAppLaunchContext,
		info *T.GAppInfo,
		files *T.GList) string

	G_app_launch_context_get_startup_notify_id func(
		context *T.GAppLaunchContext,
		info *T.GAppInfo,
		files *T.GList) string

	G_app_launch_context_launch_failed func(
		context *T.GAppLaunchContext,
		startup_notify_id string)

	G_action_get_type func() T.GType

	G_action_get_name func(action *T.GAction) string

	G_action_get_parameter_type func(action *T.GAction) *T.GVariantType

	G_action_get_state_type func(action *T.GAction) *T.GVariantType

	G_action_get_state_hint func(action *T.GAction) *T.GVariant

	G_action_get_enabled func(action *T.GAction) T.Gboolean

	G_action_get_state func(action *T.GAction) *T.GVariant

	G_action_set_state func(action *T.GAction, value *T.GVariant)

	G_action_activate func(action *T.GAction, parameter *T.GVariant)

	G_simple_action_get_type func() T.GType

	G_simple_action_new func(
		name string, parameter_type *T.GVariantType) *T.GSimpleAction

	G_simple_action_new_stateful func(
		name string,
		parameter_type *T.GVariantType,
		state *T.GVariant) *T.GSimpleAction

	G_simple_action_set_enabled func(
		simple *T.GSimpleAction,
		enabled T.Gboolean)

	G_action_group_get_type func() T.GType

	G_action_group_has_action func(
		action_group *T.GActionGroup,
		action_name string) T.Gboolean

	G_action_group_list_actions func(
		action_group *T.GActionGroup) **T.Gchar

	G_action_group_get_action_parameter_type func(
		action_group *T.GActionGroup,
		action_name string) *T.GVariantType

	G_action_group_get_action_state_type func(
		action_group *T.GActionGroup,
		action_name string) *T.GVariantType

	G_action_group_get_action_state_hint func(
		action_group *T.GActionGroup,
		action_name string) *T.GVariant

	G_action_group_get_action_enabled func(
		action_group *T.GActionGroup,
		action_name string) T.Gboolean

	G_action_group_get_action_state func(
		action_group *T.GActionGroup,
		action_name string) *T.GVariant

	G_action_group_change_action_state func(
		action_group *T.GActionGroup,
		action_name string,
		value *T.GVariant)

	G_action_group_activate_action func(
		action_group *T.GActionGroup,
		action_name string,
		parameter *T.GVariant)

	G_action_group_action_added func(
		action_group *T.GActionGroup,
		action_name string)

	G_action_group_action_removed func(
		action_group *T.GActionGroup,
		action_name string)

	G_action_group_action_enabled_changed func(
		action_group *T.GActionGroup,
		action_name string,
		enabled T.Gboolean)

	G_action_group_action_state_changed func(
		action_group *T.GActionGroup,
		action_name string,
		state *T.GVariant)

	G_simple_action_group_get_type func() T.GType

	G_simple_action_group_new func() *T.GSimpleActionGroup

	G_simple_action_group_lookup func(
		simple *T.GSimpleActionGroup,
		action_name string) *T.GAction

	G_simple_action_group_insert func(
		simple *T.GSimpleActionGroup,
		action *T.GAction)

	G_simple_action_group_remove func(
		simple *T.GSimpleActionGroup,
		action_name string)

	G_application_get_type func() T.GType

	G_application_id_is_valid func(
		application_id string) T.Gboolean

	G_application_new func(
		application_id string,
		flags T.GApplicationFlags) *T.GApplication

	G_application_get_application_id func(
		application *T.GApplication) string

	G_application_set_application_id func(
		application *T.GApplication,
		application_id string)

	G_application_get_inactivity_timeout func(
		application *T.GApplication) T.Guint

	G_application_set_inactivity_timeout func(
		application *T.GApplication,
		inactivity_timeout T.Guint)

	G_application_get_flags func(
		application *T.GApplication) T.GApplicationFlags

	G_application_set_flags func(
		application *T.GApplication,
		flags T.GApplicationFlags)

	G_application_set_action_group func(
		application *T.GApplication,
		action_group *T.GActionGroup)

	G_application_get_is_registered func(
		application *T.GApplication) T.Gboolean

	G_application_get_is_remote func(
		application *T.GApplication) T.Gboolean

	G_application_register func(
		application *T.GApplication,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_application_hold func(
		application *T.GApplication)

	G_application_release func(
		application *T.GApplication)

	G_application_activate func(
		application *T.GApplication)

	G_application_open func(
		application *T.GApplication,
		files **T.GFile,
		n_files T.Gint,
		hint string)

	G_application_run func(
		application *T.GApplication,
		argc int,
		argv **T.Char) int

	G_application_command_line_get_type func() T.GType

	G_application_command_line_get_arguments func(
		cmdline *T.GApplicationCommandLine,
		argc *int) **T.Gchar

	G_application_command_line_get_environ func(
		cmdline *T.GApplicationCommandLine) **T.Gchar

	G_application_command_line_getenv func(
		cmdline *T.GApplicationCommandLine,
		name string) string

	G_application_command_line_get_cwd func(
		cmdline *T.GApplicationCommandLine) string

	G_application_command_line_get_is_remote func(
		cmdline *T.GApplicationCommandLine) T.Gboolean

	G_application_command_line_print func(
		cmdline *T.GApplicationCommandLine, format string, v ...VArg)
	G_application_command_line_printerr func(
		cmdline *T.GApplicationCommandLine, format string, v ...VArg)

	G_application_command_line_get_exit_status func(
		cmdline *T.GApplicationCommandLine) int

	G_application_command_line_set_exit_status func(
		cmdline *T.GApplicationCommandLine,
		exit_status int)

	G_application_command_line_get_platform_data func(
		cmdline *T.GApplicationCommandLine) *T.GVariant

	G_initable_get_type func() T.GType

	G_initable_init func(
		initable *T.GInitable,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_initable_new func(object_type T.GType,
		cancellable *T.GCancellable, e **T.GError,
		first_property_name string, v ...VArg) T.Gpointer

	G_initable_newv func(
		object_type T.GType,
		n_parameters T.Guint,
		parameters *T.GParameter,
		cancellable *T.GCancellable,
		err **T.GError) T.Gpointer

	G_initable_new_valist func(
		object_type T.GType,
		first_property_name string,
		var_args T.Va_list,
		cancellable *T.GCancellable,
		err **T.GError) *T.GObject

	G_async_initable_get_type func() T.GType

	G_async_initable_init_async func(
		initable *T.GAsyncInitable,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_async_initable_init_finish func(
		initable *T.GAsyncInitable,
		res *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_async_initable_new_async func(object_type T.GType,
		io_priority int, cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback, user_data T.Gpointer,
		first_property_name string, v ...VArg)

	G_async_initable_newv_async func(
		object_type T.GType,
		n_parameters T.Guint,
		parameters *T.GParameter,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_async_initable_new_valist_async func(
		object_type T.GType,
		first_property_name string,
		var_args T.Va_list,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_async_initable_new_finish func(
		initable *T.GAsyncInitable,
		res *T.GAsyncResult,
		err **T.GError) *T.GObject

	G_async_result_get_type func() T.GType

	G_async_result_get_user_data func(
		res *T.GAsyncResult) T.Gpointer

	G_async_result_get_source_object func(
		res *T.GAsyncResult) *T.GObject

	G_input_stream_get_type func() T.GType

	G_input_stream_read func(
		stream *T.GInputStream,
		buffer *T.Void,
		count T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_input_stream_read_all func(
		stream *T.GInputStream,
		buffer *T.Void,
		count T.Gsize,
		bytes_read *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_input_stream_skip func(
		stream *T.GInputStream,
		count T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_input_stream_close func(
		stream *T.GInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_input_stream_read_async func(
		stream *T.GInputStream,
		buffer *T.Void,
		count T.Gsize,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_input_stream_read_finish func(
		stream *T.GInputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	G_input_stream_skip_async func(
		stream *T.GInputStream,
		count T.Gsize,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_input_stream_skip_finish func(
		stream *T.GInputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	G_input_stream_close_async func(
		stream *T.GInputStream,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_input_stream_close_finish func(
		stream *T.GInputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_input_stream_is_closed func(
		stream *T.GInputStream) T.Gboolean

	G_input_stream_has_pending func(
		stream *T.GInputStream) T.Gboolean

	G_input_stream_set_pending func(
		stream *T.GInputStream,
		err **T.GError) T.Gboolean

	G_input_stream_clear_pending func(
		stream *T.GInputStream)

	G_filter_input_stream_get_type func() T.GType

	G_filter_input_stream_get_base_stream func(
		stream *T.GFilterInputStream) *T.GInputStream

	G_filter_input_stream_get_close_base_stream func(
		stream *T.GFilterInputStream) T.Gboolean

	G_filter_input_stream_set_close_base_stream func(
		stream *T.GFilterInputStream,
		close_base T.Gboolean)

	G_buffered_input_stream_get_type func() T.GType

	G_buffered_input_stream_new func(
		base_stream *T.GInputStream) *T.GInputStream

	G_buffered_input_stream_new_sized func(
		base_stream *T.GInputStream,
		size T.Gsize) *T.GInputStream

	G_buffered_input_stream_get_buffer_size func(
		stream *T.GBufferedInputStream) T.Gsize

	G_buffered_input_stream_set_buffer_size func(
		stream *T.GBufferedInputStream,
		size T.Gsize)

	G_buffered_input_stream_get_available func(
		stream *T.GBufferedInputStream) T.Gsize

	G_buffered_input_stream_peek func(
		stream *T.GBufferedInputStream,
		buffer *T.Void,
		offset T.Gsize,
		count T.Gsize) T.Gsize

	G_buffered_input_stream_peek_buffer func(
		stream *T.GBufferedInputStream,
		count *T.Gsize) *T.Void

	G_buffered_input_stream_fill func(
		stream *T.GBufferedInputStream,
		count T.Gssize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_buffered_input_stream_fill_async func(
		stream *T.GBufferedInputStream,
		count T.Gssize,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_buffered_input_stream_fill_finish func(
		stream *T.GBufferedInputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	G_buffered_input_stream_read_byte func(
		stream *T.GBufferedInputStream,
		cancellable *T.GCancellable,
		err **T.GError) int

	G_output_stream_get_type func() T.GType

	G_output_stream_write func(
		stream *T.GOutputStream,
		buffer *T.Void,
		count T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_output_stream_write_all func(
		stream *T.GOutputStream,
		buffer *T.Void,
		count T.Gsize,
		bytes_written *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_output_stream_splice func(
		stream *T.GOutputStream,
		source *T.GInputStream,
		flags T.GOutputStreamSpliceFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_output_stream_flush func(
		stream *T.GOutputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_output_stream_close func(
		stream *T.GOutputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_output_stream_write_async func(
		stream *T.GOutputStream,
		buffer *T.Void,
		count T.Gsize,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_output_stream_write_finish func(
		stream *T.GOutputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	G_output_stream_splice_async func(
		stream *T.GOutputStream,
		source *T.GInputStream,
		flags T.GOutputStreamSpliceFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_output_stream_splice_finish func(
		stream *T.GOutputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	G_output_stream_flush_async func(
		stream *T.GOutputStream,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_output_stream_flush_finish func(
		stream *T.GOutputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_output_stream_close_async func(
		stream *T.GOutputStream,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_output_stream_close_finish func(
		stream *T.GOutputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_output_stream_is_closed func(
		stream *T.GOutputStream) T.Gboolean

	G_output_stream_is_closing func(
		stream *T.GOutputStream) T.Gboolean

	G_output_stream_has_pending func(
		stream *T.GOutputStream) T.Gboolean

	G_output_stream_set_pending func(
		stream *T.GOutputStream,
		err **T.GError) T.Gboolean

	G_output_stream_clear_pending func(
		stream *T.GOutputStream)

	G_filter_output_stream_get_type func() T.GType

	G_filter_output_stream_get_base_stream func(
		stream *T.GFilterOutputStream) *T.GOutputStream

	G_filter_output_stream_get_close_base_stream func(
		stream *T.GFilterOutputStream) T.Gboolean

	G_filter_output_stream_set_close_base_stream func(
		stream *T.GFilterOutputStream,
		close_base T.Gboolean)

	G_buffered_output_stream_get_type func() T.GType

	G_buffered_output_stream_new func(
		base_stream *T.GOutputStream) *T.GOutputStream

	G_buffered_output_stream_new_sized func(
		base_stream *T.GOutputStream,
		size T.Gsize) *T.GOutputStream

	G_buffered_output_stream_get_buffer_size func(
		stream *T.GBufferedOutputStream) T.Gsize

	G_buffered_output_stream_set_buffer_size func(
		stream *T.GBufferedOutputStream,
		size T.Gsize)

	G_buffered_output_stream_get_auto_grow func(
		stream *T.GBufferedOutputStream) T.Gboolean

	G_buffered_output_stream_set_auto_grow func(
		stream *T.GBufferedOutputStream,
		auto_grow T.Gboolean)

	G_cancellable_get_type func() T.GType

	G_cancellable_new func() *T.GCancellable

	G_cancellable_is_cancelled func(
		cancellable *T.GCancellable) T.Gboolean

	G_cancellable_set_error_if_cancelled func(
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_cancellable_get_fd func(
		cancellable *T.GCancellable) int

	G_cancellable_make_pollfd func(
		cancellable *T.GCancellable,
		pollfd *T.GPollFD) T.Gboolean

	G_cancellable_release_fd func(
		cancellable *T.GCancellable)

	G_cancellable_source_new func(
		cancellable *T.GCancellable) *T.GSource

	G_cancellable_get_current func() *T.GCancellable

	G_cancellable_push_current func(
		cancellable *T.GCancellable)

	G_cancellable_pop_current func(
		cancellable *T.GCancellable)

	G_cancellable_reset func(
		cancellable *T.GCancellable)

	G_cancellable_connect func(
		cancellable *T.GCancellable,
		callback T.GCallback,
		data T.Gpointer,
		data_destroy_func T.GDestroyNotify) T.Gulong

	G_cancellable_disconnect func(
		cancellable *T.GCancellable,
		handler_id T.Gulong)

	G_cancellable_cancel func(
		cancellable *T.GCancellable)

	G_converter_get_type func() T.GType

	G_converter_convert func(
		converter *T.GConverter,
		inbuf *T.Void,
		inbuf_size T.Gsize,
		outbuf *T.Void,
		outbuf_size T.Gsize,
		flags T.GConverterFlags,
		bytes_read *T.Gsize,
		bytes_written *T.Gsize,
		err **T.GError) T.GConverterResult

	G_converter_reset func(
		converter *T.GConverter)

	G_charset_converter_get_type func() T.GType

	G_charset_converter_new func(
		to_charset string,
		from_charset string,
		err **T.GError) *T.GcharsetConverter

	G_charset_converter_set_use_fallback func(
		converter *T.GcharsetConverter,
		use_fallback T.Gboolean)

	G_charset_converter_get_use_fallback func(
		converter *T.GcharsetConverter) T.Gboolean

	G_charset_converter_get_num_fallbacks func(
		converter *T.GcharsetConverter) T.Guint

	G_content_type_equals func(
		type1 string,
		type2 string) T.Gboolean

	G_content_type_is_a func(
		typ string,
		supertype string) T.Gboolean

	G_content_type_is_unknown func(
		typ string) T.Gboolean

	G_content_type_get_description func(
		typ string) string

	G_content_type_get_mime_type func(
		typ string) string

	G_content_type_get_icon func(
		typ string) *T.GIcon

	G_content_type_can_be_executable func(
		typ string) T.Gboolean

	G_content_type_from_mime_type func(
		mime_type string) string

	G_content_type_guess func(
		filename string,
		data *T.Guchar,
		data_size T.Gsize,
		result_uncertain *T.Gboolean) string

	G_content_type_guess_for_tree func(
		root *T.GFile) **T.Gchar

	G_content_types_get_registered func() *T.GList

	G_converter_input_stream_get_type func() T.GType

	G_converter_input_stream_new func(
		base_stream *T.GInputStream,
		converter *T.GConverter) *T.GInputStream

	G_converter_input_stream_get_converter func(
		converter_stream *T.GConverterInputStream) *T.GConverter

	G_converter_output_stream_get_type func() T.GType

	G_converter_output_stream_new func(
		base_stream *T.GOutputStream,
		converter *T.GConverter) *T.GOutputStream

	G_converter_output_stream_get_converter func(
		converter_stream *T.GConverterOutputStream) *T.GConverter

	G_credentials_get_type func() T.GType

	G_credentials_new func() *T.GCredentials

	G_credentials_to_string func(
		credentials *T.GCredentials) string

	G_credentials_get_native func(
		credentials *T.GCredentials,
		native_type T.GCredentialsType) T.Gpointer

	G_credentials_set_native func(
		credentials *T.GCredentials,
		native_type T.GCredentialsType,
		native T.Gpointer)

	G_credentials_is_same_user func(
		credentials *T.GCredentials,
		other_credentials *T.GCredentials,
		err **T.GError) T.Gboolean

	G_data_input_stream_get_type func() T.GType

	G_data_input_stream_new func(
		base_stream *T.GInputStream) *T.GDataInputStream

	G_data_input_stream_set_byte_order func(
		stream *T.GDataInputStream,
		order T.GDataStreamByteOrder)

	G_data_input_stream_get_byte_order func(
		stream *T.GDataInputStream) T.GDataStreamByteOrder

	G_data_input_stream_set_newline_type func(
		stream *T.GDataInputStream,
		typ T.GDataStreamNewlineType)

	G_data_input_stream_get_newline_type func(
		stream *T.GDataInputStream) T.GDataStreamNewlineType

	G_data_input_stream_read_byte func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Guchar

	G_data_input_stream_read_int16 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gint16

	G_data_input_stream_read_uint16 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Guint16

	G_data_input_stream_read_int32 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gint32

	G_data_input_stream_read_uint32 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Guint32

	G_data_input_stream_read_int64 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gint64

	G_data_input_stream_read_uint64 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Guint64

	G_data_input_stream_read_line func(
		stream *T.GDataInputStream,
		length *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) string

	G_data_input_stream_read_line_async func(
		stream *T.GDataInputStream,
		io_priority T.Gint,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_data_input_stream_read_line_finish func(
		stream *T.GDataInputStream,
		result *T.GAsyncResult,
		length *T.Gsize,
		err **T.GError) string

	G_data_input_stream_read_until func(
		stream *T.GDataInputStream,
		stop_chars string,
		length *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) string

	G_data_input_stream_read_until_async func(
		stream *T.GDataInputStream,
		stop_chars string,
		io_priority T.Gint,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_data_input_stream_read_until_finish func(
		stream *T.GDataInputStream,
		result *T.GAsyncResult,
		length *T.Gsize,
		err **T.GError) string

	G_data_input_stream_read_upto func(
		stream *T.GDataInputStream,
		stop_chars string,
		stop_chars_len T.Gssize,
		length *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) string

	G_data_input_stream_read_upto_async func(
		stream *T.GDataInputStream,
		stop_chars string,
		stop_chars_len T.Gssize,
		io_priority T.Gint,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_data_input_stream_read_upto_finish func(
		stream *T.GDataInputStream,
		result *T.GAsyncResult,
		length *T.Gsize,
		err **T.GError) string

	G_data_output_stream_get_type func() T.GType

	G_data_output_stream_new func(
		base_stream *T.GOutputStream) *T.GDataOutputStream

	G_data_output_stream_set_byte_order func(
		stream *T.GDataOutputStream,
		order T.GDataStreamByteOrder)

	G_data_output_stream_get_byte_order func(
		stream *T.GDataOutputStream) T.GDataStreamByteOrder

	G_data_output_stream_put_byte func(
		stream *T.GDataOutputStream,
		data T.Guchar,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_data_output_stream_put_int16 func(
		stream *T.GDataOutputStream,
		data T.Gint16,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_data_output_stream_put_uint16 func(
		stream *T.GDataOutputStream,
		data T.Guint16,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_data_output_stream_put_int32 func(
		stream *T.GDataOutputStream,
		data T.Gint32,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_data_output_stream_put_uint32 func(
		stream *T.GDataOutputStream,
		data T.Guint32,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_data_output_stream_put_int64 func(
		stream *T.GDataOutputStream,
		data T.Gint64,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_data_output_stream_put_uint64 func(
		stream *T.GDataOutputStream,
		data T.Guint64,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_data_output_stream_put_string func(
		stream *T.GDataOutputStream,
		str string,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_dbus_is_address func(
		string string) T.Gboolean

	G_dbus_is_supported_address func(
		string string,
		err **T.GError) T.Gboolean

	G_dbus_address_get_stream func(
		address string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_address_get_stream_finish func(
		res *T.GAsyncResult,
		out_guid **T.Gchar,
		err **T.GError) *T.GIOStream

	G_dbus_address_get_stream_sync func(
		address string,
		out_guid **T.Gchar,
		cancellable *T.GCancellable,
		err **T.GError) *T.GIOStream

	G_dbus_address_get_for_bus_sync func(
		bus_type T.GBusType,
		cancellable *T.GCancellable,
		err **T.GError) string

	G_dbus_auth_observer_get_type func() T.GType

	G_dbus_auth_observer_new func() *T.GDBusAuthObserver

	G_dbus_auth_observer_authorize_authenticated_peer func(
		observer *T.GDBusAuthObserver,
		stream *T.GIOStream,
		credentials *T.GCredentials) T.Gboolean

	G_dbus_connection_get_type func() T.GType

	G_bus_get func(
		bus_type T.GBusType,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_bus_get_finish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusConnection

	G_bus_get_sync func(
		bus_type T.GBusType,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusConnection

	G_dbus_connection_new func(
		stream *T.GIOStream,
		guid string,
		flags T.GDBusConnectionFlags,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_connection_new_finish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusConnection

	G_dbus_connection_new_sync func(
		stream *T.GIOStream,
		guid string,
		flags T.GDBusConnectionFlags,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusConnection

	G_dbus_connection_new_for_address func(
		address string,
		flags T.GDBusConnectionFlags,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_connection_new_for_address_finish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusConnection

	G_dbus_connection_new_for_address_sync func(
		address string,
		flags T.GDBusConnectionFlags,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusConnection

	G_dbus_connection_start_message_processing func(
		connection *T.GDBusConnection)

	G_dbus_connection_is_closed func(
		connection *T.GDBusConnection) T.Gboolean

	G_dbus_connection_get_stream func(
		connection *T.GDBusConnection) *T.GIOStream

	G_dbus_connection_get_guid func(
		connection *T.GDBusConnection) string

	G_dbus_connection_get_unique_name func(
		connection *T.GDBusConnection) string

	G_dbus_connection_get_peer_credentials func(
		connection *T.GDBusConnection) *T.GCredentials

	G_dbus_connection_get_exit_on_close func(
		connection *T.GDBusConnection) T.Gboolean

	G_dbus_connection_set_exit_on_close func(
		connection *T.GDBusConnection,
		exit_on_close T.Gboolean)

	G_dbus_connection_get_capabilities func(
		connection *T.GDBusConnection) T.GDBusCapabilityFlags

	G_dbus_connection_close func(
		connection *T.GDBusConnection,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_connection_close_finish func(
		connection *T.GDBusConnection,
		res *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_dbus_connection_close_sync func(
		connection *T.GDBusConnection,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_dbus_connection_flush func(
		connection *T.GDBusConnection,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_connection_flush_finish func(
		connection *T.GDBusConnection,
		res *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_dbus_connection_flush_sync func(
		connection *T.GDBusConnection,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_dbus_connection_send_message func(
		connection *T.GDBusConnection,
		message *T.GDBusMessage,
		flags T.GDBusSendMessageFlags,
		out_serial *T.Guint32,
		err **T.GError) T.Gboolean

	G_dbus_connection_send_message_with_reply func(
		connection *T.GDBusConnection,
		message *T.GDBusMessage,
		flags T.GDBusSendMessageFlags,
		timeout_msec T.Gint,
		out_serial *T.Guint32,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_connection_send_message_with_reply_finish func(
		connection *T.GDBusConnection,
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusMessage

	G_dbus_connection_send_message_with_reply_sync func(
		connection *T.GDBusConnection,
		message *T.GDBusMessage,
		flags T.GDBusSendMessageFlags,
		timeout_msec T.Gint,
		out_serial *T.Guint32,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusMessage

	G_dbus_connection_emit_signal func(
		connection *T.GDBusConnection,
		destination_bus_name string,
		object_path string,
		interface_name string,
		signal_name string,
		parameters *T.GVariant,
		err **T.GError) T.Gboolean

	G_dbus_connection_call func(
		connection *T.GDBusConnection,
		bus_name string,
		object_path string,
		interface_name string,
		method_name string,
		parameters *T.GVariant,
		reply_type *T.GVariantType,
		flags T.GDBusCallFlags,
		timeout_msec T.Gint,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_connection_call_finish func(
		connection *T.GDBusConnection,
		res *T.GAsyncResult,
		err **T.GError) *T.GVariant

	G_dbus_connection_call_sync func(
		connection *T.GDBusConnection,
		bus_name string,
		object_path string,
		interface_name string,
		method_name string,
		parameters *T.GVariant,
		reply_type *T.GVariantType,
		flags T.GDBusCallFlags,
		timeout_msec T.Gint,
		cancellable *T.GCancellable,
		err **T.GError) *T.GVariant

	G_dbus_connection_register_object func(
		connection *T.GDBusConnection,
		object_path string,
		interface_info *T.GDBusInterfaceInfo,
		vtable *T.GDBusInterfaceVTable,
		user_data T.Gpointer,
		user_data_free_func T.GDestroyNotify,
		err **T.GError) T.Guint

	G_dbus_connection_unregister_object func(
		connection *T.GDBusConnection,
		registration_id T.Guint) T.Gboolean

	G_dbus_connection_register_subtree func(
		connection *T.GDBusConnection,
		object_path string,
		vtable *T.GDBusSubtreeVTable,
		flags T.GDBusSubtreeFlags,
		user_data T.Gpointer,
		user_data_free_func T.GDestroyNotify,
		err **T.GError) T.Guint

	G_dbus_connection_unregister_subtree func(
		connection *T.GDBusConnection,
		registration_id T.Guint) T.Gboolean

	G_dbus_connection_signal_subscribe func(
		connection *T.GDBusConnection,
		sender string,
		interface_name string,
		member string,
		object_path string,
		arg0 string,
		flags T.GDBusSignalFlags,
		callback T.GDBusSignalCallback,
		user_data T.Gpointer,
		user_data_free_func T.GDestroyNotify) T.Guint

	G_dbus_connection_signal_unsubscribe func(
		connection *T.GDBusConnection,
		subscription_id T.Guint)

	G_dbus_connection_add_filter func(
		connection *T.GDBusConnection,
		filter_function T.GDBusMessageFilterFunction,
		user_data T.Gpointer,
		user_data_free_func T.GDestroyNotify) T.Guint

	G_dbus_connection_remove_filter func(
		connection *T.GDBusConnection,
		filter_id T.Guint)

	G_dbus_error_quark func() T.GQuark

	G_dbus_error_is_remote_error func(
		err *T.GError) T.Gboolean

	G_dbus_error_get_remote_error func(
		err *T.GError) string

	G_dbus_error_strip_remote_error func(
		err *T.GError) T.Gboolean

	G_dbus_error_register_error func(
		error_domain T.GQuark,
		error_code T.Gint,
		dbus_error_name string) T.Gboolean

	G_dbus_error_unregister_error func(
		error_domain T.GQuark,
		error_code T.Gint,
		dbus_error_name string) T.Gboolean

	G_dbus_error_register_error_domain func(
		error_domain_quark_name string,
		quark_volatile *T.Gsize,
		entries *T.GDBusErrorEntry,
		num_entries T.Guint)

	G_dbus_error_new_for_dbus_error func(
		dbus_error_name string,
		dbus_error_message string) *T.GError

	G_dbus_error_set_dbus_error func(e **T.GError,
		dbus_error_name, dbus_error_message,
		format string, v ...VArg)

	G_dbus_error_set_dbus_error_valist func(
		err **T.GError,
		dbus_error_name string,
		dbus_error_message string,
		format string,
		var_args T.Va_list)

	G_dbus_error_encode_gerror func(
		err *T.GError) string

	G_dbus_annotation_info_lookup func(
		annotations **T.GDBusAnnotationInfo,
		name string) string

	G_dbus_interface_info_lookup_method func(
		info *T.GDBusInterfaceInfo,
		name string) *T.GDBusMethodInfo

	G_dbus_interface_info_lookup_signal func(
		info *T.GDBusInterfaceInfo,
		name string) *T.GDBusSignalInfo

	G_dbus_interface_info_lookup_property func(
		info *T.GDBusInterfaceInfo,
		name string) *T.GDBusPropertyInfo

	G_dbus_interface_info_generate_xml func(
		info *T.GDBusInterfaceInfo,
		indent T.Guint,
		string_builder *T.GString)

	G_dbus_node_info_new_for_xml func(
		xml_data string,
		err **T.GError) *T.GDBusNodeInfo

	G_dbus_node_info_lookup_interface func(
		info *T.GDBusNodeInfo,
		name string) *T.GDBusInterfaceInfo

	G_dbus_node_info_generate_xml func(
		info *T.GDBusNodeInfo,
		indent T.Guint,
		string_builder *T.GString)

	G_dbus_node_info_ref func(
		info *T.GDBusNodeInfo) *T.GDBusNodeInfo

	G_dbus_interface_info_ref func(
		info *T.GDBusInterfaceInfo) *T.GDBusInterfaceInfo

	G_dbus_method_info_ref func(
		info *T.GDBusMethodInfo) *T.GDBusMethodInfo

	G_dbus_signal_info_ref func(
		info *T.GDBusSignalInfo) *T.GDBusSignalInfo

	G_dbus_property_info_ref func(
		info *T.GDBusPropertyInfo) *T.GDBusPropertyInfo

	G_dbus_arg_info_ref func(
		info *T.GDBusArgInfo) *T.GDBusArgInfo

	G_dbus_annotation_info_ref func(
		info *T.GDBusAnnotationInfo) *T.GDBusAnnotationInfo

	G_dbus_node_info_unref func(
		info *T.GDBusNodeInfo)

	G_dbus_interface_info_unref func(
		info *T.GDBusInterfaceInfo)

	G_dbus_method_info_unref func(
		info *T.GDBusMethodInfo)

	G_dbus_signal_info_unref func(
		info *T.GDBusSignalInfo)

	G_dbus_property_info_unref func(
		info *T.GDBusPropertyInfo)

	G_dbus_arg_info_unref func(
		info *T.GDBusArgInfo)

	G_dbus_annotation_info_unref func(
		info *T.GDBusAnnotationInfo)

	G_dbus_node_info_get_type func() T.GType

	G_dbus_interface_info_get_type func() T.GType

	G_dbus_method_info_get_type func() T.GType

	G_dbus_signal_info_get_type func() T.GType

	G_dbus_property_info_get_type func() T.GType

	G_dbus_arg_info_get_type func() T.GType

	G_dbus_annotation_info_get_type func() T.GType

	G_dbus_message_get_type func() T.GType

	G_dbus_message_new func() *T.GDBusMessage

	G_dbus_message_new_signal func(
		path string,
		interface_ string,
		signal string) *T.GDBusMessage

	G_dbus_message_new_method_call func(
		name string,
		path string,
		interface_ string,
		method string) *T.GDBusMessage

	G_dbus_message_new_method_reply func(
		method_call_message *T.GDBusMessage) *T.GDBusMessage

	G_dbus_message_new_method_error func(
		method_call_message *T.GDBusMessage,
		error_name, error_message_format string,
		v ...VArg) *T.GDBusMessage

	G_dbus_message_new_method_error_valist func(
		method_call_message *T.GDBusMessage,
		error_name string,
		error_message_format string,
		var_args T.Va_list) *T.GDBusMessage

	G_dbus_message_new_method_error_literal func(
		method_call_message *T.GDBusMessage,
		error_name string,
		error_message string) *T.GDBusMessage

	G_dbus_message_print func(
		message *T.GDBusMessage,
		indent T.Guint) string

	G_dbus_message_get_locked func(
		message *T.GDBusMessage) T.Gboolean

	G_dbus_message_lock func(
		message *T.GDBusMessage)

	G_dbus_message_copy func(
		message *T.GDBusMessage,
		err **T.GError) *T.GDBusMessage

	G_dbus_message_get_byte_order func(
		message *T.GDBusMessage) T.GDBusMessageByteOrder

	G_dbus_message_set_byte_order func(
		message *T.GDBusMessage,
		byte_order T.GDBusMessageByteOrder)

	G_dbus_message_get_message_type func(
		message *T.GDBusMessage) T.GDBusMessageType

	G_dbus_message_set_message_type func(
		message *T.GDBusMessage,
		typ T.GDBusMessageType)

	G_dbus_message_get_flags func(
		message *T.GDBusMessage) T.GDBusMessageFlags

	G_dbus_message_set_flags func(
		message *T.GDBusMessage,
		flags T.GDBusMessageFlags)

	G_dbus_message_get_serial func(
		message *T.GDBusMessage) T.Guint32

	G_dbus_message_set_serial func(
		message *T.GDBusMessage,
		serial T.Guint32)

	G_dbus_message_get_header func(
		message *T.GDBusMessage,
		header_field T.GDBusMessageHeaderField) *T.GVariant

	G_dbus_message_set_header func(
		message *T.GDBusMessage,
		header_field T.GDBusMessageHeaderField,
		value *T.GVariant)

	G_dbus_message_get_header_fields func(
		message *T.GDBusMessage) *T.Guchar

	G_dbus_message_get_body func(
		message *T.GDBusMessage) *T.GVariant

	G_dbus_message_set_body func(
		message *T.GDBusMessage,
		body *T.GVariant)

	G_dbus_message_get_unix_fd_list func(
		message *T.GDBusMessage) *T.GUnixFDList

	G_dbus_message_set_unix_fd_list func(
		message *T.GDBusMessage,
		fd_list *T.GUnixFDList)

	G_dbus_message_get_reply_serial func(
		message *T.GDBusMessage) T.Guint32

	G_dbus_message_set_reply_serial func(
		message *T.GDBusMessage,
		value T.Guint32)

	G_dbus_message_get_interface func(
		message *T.GDBusMessage) string

	G_dbus_message_set_interface func(
		message *T.GDBusMessage,
		value string)

	G_dbus_message_get_member func(
		message *T.GDBusMessage) string

	G_dbus_message_set_member func(
		message *T.GDBusMessage,
		value string)

	G_dbus_message_get_path func(
		message *T.GDBusMessage) string

	G_dbus_message_set_path func(
		message *T.GDBusMessage,
		value string)

	G_dbus_message_get_sender func(
		message *T.GDBusMessage) string

	G_dbus_message_set_sender func(
		message *T.GDBusMessage,
		value string)

	G_dbus_message_get_destination func(
		message *T.GDBusMessage) string

	G_dbus_message_set_destination func(
		message *T.GDBusMessage,
		value string)

	G_dbus_message_get_error_name func(
		message *T.GDBusMessage) string

	G_dbus_message_set_error_name func(
		message *T.GDBusMessage,
		value string)

	G_dbus_message_get_signature func(
		message *T.GDBusMessage) string

	G_dbus_message_set_signature func(
		message *T.GDBusMessage,
		value string)

	G_dbus_message_get_num_unix_fds func(
		message *T.GDBusMessage) T.Guint32

	G_dbus_message_set_num_unix_fds func(
		message *T.GDBusMessage,
		value T.Guint32)

	G_dbus_message_get_arg0 func(
		message *T.GDBusMessage) string

	G_dbus_message_new_from_blob func(
		blob *T.Guchar,
		blob_len T.Gsize,
		capabilities T.GDBusCapabilityFlags,
		err **T.GError) *T.GDBusMessage

	G_dbus_message_bytes_needed func(
		blob *T.Guchar,
		blob_len T.Gsize,
		err **T.GError) T.Gssize

	G_dbus_message_to_blob func(
		message *T.GDBusMessage,
		out_size *T.Gsize,
		capabilities T.GDBusCapabilityFlags,
		err **T.GError) *T.Guchar

	G_dbus_message_to_gerror func(
		message *T.GDBusMessage,
		err **T.GError) T.Gboolean

	G_dbus_method_invocation_get_type func() T.GType

	G_dbus_method_invocation_get_sender func(
		invocation *T.GDBusMethodInvocation) string

	G_dbus_method_invocation_get_object_path func(
		invocation *T.GDBusMethodInvocation) string

	G_dbus_method_invocation_get_interface_name func(
		invocation *T.GDBusMethodInvocation) string

	G_dbus_method_invocation_get_method_name func(
		invocation *T.GDBusMethodInvocation) string

	G_dbus_method_invocation_get_method_info func(
		invocation *T.GDBusMethodInvocation) *T.GDBusMethodInfo

	G_dbus_method_invocation_get_connection func(
		invocation *T.GDBusMethodInvocation) *T.GDBusConnection

	G_dbus_method_invocation_get_message func(
		invocation *T.GDBusMethodInvocation) *T.GDBusMessage

	G_dbus_method_invocation_get_parameters func(
		invocation *T.GDBusMethodInvocation) *T.GVariant

	G_dbus_method_invocation_get_user_data func(
		invocation *T.GDBusMethodInvocation) T.Gpointer

	G_dbus_method_invocation_return_value func(
		invocation *T.GDBusMethodInvocation,
		parameters *T.GVariant)

	G_dbus_method_invocation_return_error func(
		invocation *T.GDBusMethodInvocation,
		domain T.GQuark, code T.Gint, format string, v ...VArg)

	G_dbus_method_invocation_return_error_valist func(
		invocation *T.GDBusMethodInvocation,
		domain T.GQuark,
		code T.Gint,
		format string,
		var_args T.Va_list)

	G_dbus_method_invocation_return_error_literal func(
		invocation *T.GDBusMethodInvocation,
		domain T.GQuark,
		code T.Gint,
		message string)

	G_dbus_method_invocation_return_gerror func(
		invocation *T.GDBusMethodInvocation,
		err *T.GError)

	G_dbus_method_invocation_return_dbus_error func(
		invocation *T.GDBusMethodInvocation,
		error_name string,
		error_message string)

	G_bus_own_name func(
		bus_type T.GBusType,
		name string,
		flags T.GBusNameOwnerFlags,
		bus_acquired_handler T.GBusAcquiredCallback,
		name_acquired_handler T.GBusNameAcquiredCallback,
		name_lost_handler T.GBusNameLostCallback,
		user_data T.Gpointer,
		user_data_free_func T.GDestroyNotify) T.Guint

	G_bus_own_name_on_connection func(
		connection *T.GDBusConnection,
		name string,
		flags T.GBusNameOwnerFlags,
		name_acquired_handler T.GBusNameAcquiredCallback,
		name_lost_handler T.GBusNameLostCallback,
		user_data T.Gpointer,
		user_data_free_func T.GDestroyNotify) T.Guint

	G_bus_own_name_with_closures func(
		bus_type T.GBusType,
		name string,
		flags T.GBusNameOwnerFlags,
		bus_acquired_closure *T.GClosure,
		name_acquired_closure *T.GClosure,
		name_lost_closure *T.GClosure) T.Guint

	G_bus_own_name_on_connection_with_closures func(

		connection *T.GDBusConnection,
		name string,
		flags T.GBusNameOwnerFlags,
		name_acquired_closure *T.GClosure,
		name_lost_closure *T.GClosure) T.Guint

	G_bus_unown_name func(
		owner_id T.Guint)

	G_bus_watch_name func(
		bus_type T.GBusType,
		name string,
		flags T.GBusNameWatcherFlags,
		name_appeared_handler T.GBusNameAppearedCallback,
		name_vanished_handler T.GBusNameVanishedCallback,
		user_data T.Gpointer,
		user_data_free_func T.GDestroyNotify) T.Guint

	G_bus_watch_name_on_connection func(
		connection *T.GDBusConnection,
		name string,
		flags T.GBusNameWatcherFlags,
		name_appeared_handler T.GBusNameAppearedCallback,
		name_vanished_handler T.GBusNameVanishedCallback,
		user_data T.Gpointer,
		user_data_free_func T.GDestroyNotify) T.Guint

	G_bus_watch_name_with_closures func(
		bus_type T.GBusType,
		name string,
		flags T.GBusNameWatcherFlags,
		name_appeared_closure *T.GClosure,
		name_vanished_closure *T.GClosure) T.Guint

	G_bus_watch_name_on_connection_with_closures func(

		connection *T.GDBusConnection,
		name string,
		flags T.GBusNameWatcherFlags,
		name_appeared_closure *T.GClosure,
		name_vanished_closure *T.GClosure) T.Guint

	G_bus_unwatch_name func(
		watcher_id T.Guint)

	G_dbus_proxy_get_type func() T.GType

	G_dbus_proxy_new func(
		connection *T.GDBusConnection,
		flags T.GDBusProxyFlags,
		info *T.GDBusInterfaceInfo,
		name string,
		object_path string,
		interface_name string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_proxy_new_finish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusProxy

	G_dbus_proxy_new_sync func(
		connection *T.GDBusConnection,
		flags T.GDBusProxyFlags,
		info *T.GDBusInterfaceInfo,
		name string,
		object_path string,
		interface_name string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusProxy

	G_dbus_proxy_new_for_bus func(
		bus_type T.GBusType,
		flags T.GDBusProxyFlags,
		info *T.GDBusInterfaceInfo,
		name string,
		object_path string,
		interface_name string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_proxy_new_for_bus_finish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusProxy

	G_dbus_proxy_new_for_bus_sync func(
		bus_type T.GBusType,
		flags T.GDBusProxyFlags,
		info *T.GDBusInterfaceInfo,
		name string,
		object_path string,
		interface_name string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusProxy

	G_dbus_proxy_get_connection func(
		proxy *T.GDBusProxy) *T.GDBusConnection

	G_dbus_proxy_get_flags func(
		proxy *T.GDBusProxy) T.GDBusProxyFlags

	G_dbus_proxy_get_name func(
		proxy *T.GDBusProxy) string

	G_dbus_proxy_get_name_owner func(
		proxy *T.GDBusProxy) string

	G_dbus_proxy_get_object_path func(
		proxy *T.GDBusProxy) string

	G_dbus_proxy_get_interface_name func(
		proxy *T.GDBusProxy) string

	G_dbus_proxy_get_default_timeout func(
		proxy *T.GDBusProxy) T.Gint

	G_dbus_proxy_set_default_timeout func(
		proxy *T.GDBusProxy,
		timeout_msec T.Gint)

	G_dbus_proxy_get_interface_info func(
		proxy *T.GDBusProxy) *T.GDBusInterfaceInfo

	G_dbus_proxy_set_interface_info func(
		proxy *T.GDBusProxy,
		info *T.GDBusInterfaceInfo)

	G_dbus_proxy_get_cached_property func(
		proxy *T.GDBusProxy,
		property_name string) *T.GVariant

	G_dbus_proxy_set_cached_property func(
		proxy *T.GDBusProxy,
		property_name string,
		value *T.GVariant)

	G_dbus_proxy_get_cached_property_names func(
		proxy *T.GDBusProxy) **T.Gchar

	G_dbus_proxy_call func(
		proxy *T.GDBusProxy,
		method_name string,
		parameters *T.GVariant,
		flags T.GDBusCallFlags,
		timeout_msec T.Gint,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_dbus_proxy_call_finish func(
		proxy *T.GDBusProxy,
		res *T.GAsyncResult,
		err **T.GError) *T.GVariant

	G_dbus_proxy_call_sync func(
		proxy *T.GDBusProxy,
		method_name string,
		parameters *T.GVariant,
		flags T.GDBusCallFlags,
		timeout_msec T.Gint,
		cancellable *T.GCancellable,
		err **T.GError) *T.GVariant

	G_dbus_server_get_type func() T.GType

	G_dbus_server_new_sync func(
		address string,
		flags T.GDBusServerFlags,
		guid string,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusServer

	G_dbus_server_get_client_address func(
		server *T.GDBusServer) string

	G_dbus_server_get_guid func(
		server *T.GDBusServer) string

	G_dbus_server_get_flags func(
		server *T.GDBusServer) T.GDBusServerFlags

	G_dbus_server_start func(
		server *T.GDBusServer)

	G_dbus_server_stop func(
		server *T.GDBusServer)

	G_dbus_server_is_active func(
		server *T.GDBusServer) T.Gboolean

	G_dbus_is_guid func(
		string string) T.Gboolean

	G_dbus_generate_guid func() string

	G_dbus_is_name func(
		string string) T.Gboolean

	G_dbus_is_unique_name func(
		string string) T.Gboolean

	G_dbus_is_member_name func(
		string string) T.Gboolean

	G_dbus_is_interface_name func(
		string string) T.Gboolean

	G_drive_get_type func() T.GType

	G_drive_get_name func(drive *T.GDrive) string

	G_drive_get_icon func(drive *T.GDrive) *T.GIcon

	G_drive_has_volumes func(drive *T.GDrive) T.Gboolean

	G_drive_get_volumes func(drive *T.GDrive) *T.GList

	G_drive_is_media_removable func(drive *T.GDrive) T.Gboolean

	G_drive_has_media func(drive *T.GDrive) T.Gboolean

	G_drive_is_media_check_automatic func(drive *T.GDrive) T.Gboolean

	G_drive_can_poll_for_media func(drive *T.GDrive) T.Gboolean

	G_drive_can_eject func(drive *T.GDrive) T.Gboolean

	G_drive_eject func(
		drive *T.GDrive,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_drive_eject_finish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_drive_poll_for_media func(
		drive *T.GDrive,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_drive_poll_for_media_finish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_drive_get_identifier func(drive *T.GDrive, kind string) string

	G_drive_enumerate_identifiers func(
		drive *T.GDrive) **T.Char

	G_drive_get_start_stop_type func(
		drive *T.GDrive) T.GDriveStartStopType

	G_drive_can_start func(drive *T.GDrive) T.Gboolean

	G_drive_can_start_degraded func(drive *T.GDrive) T.Gboolean

	G_drive_start func(
		drive *T.GDrive,
		flags T.GDriveStartFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_drive_start_finish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_drive_can_stop func(drive *T.GDrive) T.Gboolean

	G_drive_stop func(
		drive *T.GDrive,
		flags T.GMountUnmountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_drive_stop_finish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_drive_eject_with_operation func(
		drive *T.GDrive,
		flags T.GMountUnmountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_drive_eject_with_operation_finish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_icon_get_type func() T.GType

	G_icon_hash func(icon T.Gconstpointer) T.Guint

	G_icon_equal func(icon1 *T.GIcon, icon2 *T.GIcon) T.Gboolean

	G_icon_to_string func(icon *T.GIcon) string

	G_icon_new_for_string func(str string, err **T.GError) *T.GIcon

	G_emblem_get_type func() T.GType

	G_emblem_new func(icon *T.GIcon) *T.GEmblem

	G_emblem_new_with_origin func(
		icon *T.GIcon, origin T.GEmblemOrigin) *T.GEmblem

	G_emblem_get_icon func(emblem *T.GEmblem) *T.GIcon

	G_emblem_get_origin func(emblem *T.GEmblem) T.GEmblemOrigin

	G_emblemed_icon_get_type func() T.GType

	G_emblemed_icon_new func(icon *T.GIcon, emblem *T.GEmblem) *T.GIcon

	G_emblemed_icon_get_icon func(emblemed *T.GEmblemedIcon) *T.GIcon

	G_emblemed_icon_get_emblems func(
		emblemed *T.GEmblemedIcon) *T.GList

	G_emblemed_icon_add_emblem func(
		emblemed *T.GEmblemedIcon, emblem *T.GEmblem)

	G_emblemed_icon_clear_emblems func(emblemed *T.GEmblemedIcon)

	G_file_attribute_info_list_get_type func() T.GType

	G_file_attribute_info_list_new func() *T.GFileAttributeInfoList

	G_file_attribute_info_list_ref func(
		list *T.GFileAttributeInfoList) *T.GFileAttributeInfoList

	G_file_attribute_info_list_unref func(
		list *T.GFileAttributeInfoList)

	G_file_attribute_info_list_dup func(
		list *T.GFileAttributeInfoList) *T.GFileAttributeInfoList

	G_file_attribute_info_list_lookup func(
		list *T.GFileAttributeInfoList,
		name string) *T.GFileAttributeInfo

	G_file_attribute_info_list_add func(
		list *T.GFileAttributeInfoList,
		name string,
		typ T.GFileAttributeType,
		flags T.GFileAttributeInfoFlags)

	G_file_enumerator_get_type func() T.GType

	G_file_enumerator_next_file func(
		enumerator *T.GFileEnumerator,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	G_file_enumerator_close func(
		enumerator *T.GFileEnumerator,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_enumerator_next_files_async func(
		enumerator *T.GFileEnumerator,
		num_files int,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_enumerator_next_files_finish func(
		enumerator *T.GFileEnumerator,
		result *T.GAsyncResult,
		err **T.GError) *T.GList

	G_file_enumerator_close_async func(
		enumerator *T.GFileEnumerator,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_enumerator_close_finish func(
		enumerator *T.GFileEnumerator,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_enumerator_is_closed func(
		enumerator *T.GFileEnumerator) T.Gboolean

	G_file_enumerator_has_pending func(
		enumerator *T.GFileEnumerator) T.Gboolean

	G_file_enumerator_set_pending func(
		enumerator *T.GFileEnumerator,
		pending T.Gboolean)

	G_file_enumerator_get_container func(
		enumerator *T.GFileEnumerator) *T.GFile

	G_file_get_type func() T.GType

	G_file_new_for_path func(path string) *T.GFile

	G_file_new_for_uri func(uri string) *T.GFile

	G_file_new_for_commandline_arg func(arg string) *T.GFile

	G_file_parse_name func(parse_name string) *T.GFile

	G_file_dup func(file *T.GFile) *T.GFile

	G_file_hash func(file T.Gconstpointer) T.Guint

	G_file_equal func(file1 *T.GFile, file2 *T.GFile) T.Gboolean

	G_file_get_basename func(file *T.GFile) string

	G_file_get_path func(file *T.GFile) string

	G_file_get_uri func(file *T.GFile) string

	G_file_get_parse_name func(file *T.GFile) string

	G_file_get_parent func(file *T.GFile) *T.GFile

	G_file_has_parent func(file *T.GFile, parent *T.GFile) T.Gboolean

	G_file_get_child func(file *T.GFile, name string) *T.GFile

	G_file_get_child_for_display_name func(
		file *T.GFile, display_name string, err **T.GError) *T.GFile

	G_file_has_prefix func(file *T.GFile, prefix *T.GFile) T.Gboolean

	G_file_get_relative_path func(
		parent *T.GFile, descendant *T.GFile) string

	G_file_resolve_relative_path func(
		file *T.GFile, relative_path string) *T.GFile

	G_file_is_native func(file *T.GFile) T.Gboolean

	G_file_has_uri_scheme func(
		file *T.GFile, uri_scheme string) T.Gboolean

	G_file_get_uri_scheme func(file *T.GFile) string

	G_file_read func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInputStream

	G_file_read_async func(
		file *T.GFile,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_read_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileInputStream

	G_file_append_to func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileOutputStream

	G_file_create func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileOutputStream

	G_file_replace func(
		file *T.GFile,
		etag string,
		make_backup T.Gboolean,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileOutputStream

	G_file_append_to_async func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_append_to_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileOutputStream

	G_file_create_async func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_create_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileOutputStream

	G_file_replace_async func(
		file *T.GFile,
		etag string,
		make_backup T.Gboolean,
		flags T.GFileCreateFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_replace_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileOutputStream

	G_file_open_readwrite func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileIOStream

	G_file_open_readwrite_async func(
		file *T.GFile,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_open_readwrite_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileIOStream

	G_file_create_readwrite func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileIOStream

	G_file_create_readwrite_async func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_create_readwrite_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileIOStream

	G_file_replace_readwrite func(
		file *T.GFile,
		etag string,
		make_backup T.Gboolean,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileIOStream

	G_file_replace_readwrite_async func(
		file *T.GFile,
		etag string,
		make_backup T.Gboolean,
		flags T.GFileCreateFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_replace_readwrite_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileIOStream

	G_file_query_exists func(
		file *T.GFile,
		cancellable *T.GCancellable) T.Gboolean

	G_file_query_file_type func(
		file *T.GFile,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable) T.GFileType

	G_file_query_info func(
		file *T.GFile,
		attributes string,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	G_file_query_info_async func(
		file *T.GFile,
		attributes string,
		flags T.GFileQueryInfoFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_query_info_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	G_file_query_filesystem_info func(
		file *T.GFile,
		attributes string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	G_file_query_filesystem_info_async func(
		file *T.GFile,
		attributes string,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_query_filesystem_info_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	G_file_find_enclosing_mount func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GMount

	G_file_find_enclosing_mount_async func(
		file *T.GFile,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_find_enclosing_mount_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GMount

	G_file_enumerate_children func(
		file *T.GFile,
		attributes string,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileEnumerator

	G_file_enumerate_children_async func(
		file *T.GFile,
		attributes string,
		flags T.GFileQueryInfoFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_enumerate_children_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileEnumerator

	G_file_set_display_name func(
		file *T.GFile,
		display_name string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFile

	G_file_set_display_name_async func(
		file *T.GFile,
		display_name string,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_set_display_name_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFile

	G_file_delete func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_trash func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_copy func(
		source *T.GFile,
		destination *T.GFile,
		flags T.GFileCopyFlags,
		cancellable *T.GCancellable,
		progress_callback T.GFileProgressCallback,
		progress_callback_data T.Gpointer,
		err **T.GError) T.Gboolean

	G_file_copy_async func(
		source *T.GFile,
		destination *T.GFile,
		flags T.GFileCopyFlags,
		io_priority int,
		cancellable *T.GCancellable,
		progress_callback T.GFileProgressCallback,
		progress_callback_data T.Gpointer,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_copy_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_move func(
		source *T.GFile,
		destination *T.GFile,
		flags T.GFileCopyFlags,
		cancellable *T.GCancellable,
		progress_callback T.GFileProgressCallback,
		progress_callback_data T.Gpointer,
		err **T.GError) T.Gboolean

	G_file_make_directory func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_make_directory_with_parents func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_make_symbolic_link func(
		file *T.GFile,
		symlink_value string,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_query_settable_attributes func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileAttributeInfoList

	G_file_query_writable_namespaces func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileAttributeInfoList

	G_file_set_attribute func(
		file *T.GFile,
		attribute string,
		typ T.GFileAttributeType,
		value_p T.Gpointer,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_set_attributes_from_info func(
		file *T.GFile,
		info *T.GFileInfo,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_set_attributes_async func(
		file *T.GFile,
		info *T.GFileInfo,
		flags T.GFileQueryInfoFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_set_attributes_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		info **T.GFileInfo,
		err **T.GError) T.Gboolean

	G_file_set_attribute_string func(
		file *T.GFile,
		attribute string,
		value string,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_set_attribute_byte_string func(
		file *T.GFile,
		attribute string,
		value string,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_set_attribute_uint32 func(
		file *T.GFile,
		attribute string,
		value T.Guint32,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_set_attribute_int32 func(
		file *T.GFile,
		attribute string,
		value T.Gint32,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_set_attribute_uint64 func(
		file *T.GFile,
		attribute string,
		value T.Guint64,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_set_attribute_int64 func(
		file *T.GFile,
		attribute string,
		value T.Gint64,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_mount_enclosing_volume func(
		location *T.GFile,
		flags T.GMountMountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_mount_enclosing_volume_finish func(
		location *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_mount_mountable func(
		file *T.GFile,
		flags T.GMountMountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_mount_mountable_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) *T.GFile

	G_file_unmount_mountable func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_unmount_mountable_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_unmount_mountable_with_operation func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_unmount_mountable_with_operation_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_eject_mountable func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_eject_mountable_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_eject_mountable_with_operation func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_eject_mountable_with_operation_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_copy_attributes func(
		source *T.GFile,
		destination *T.GFile,
		flags T.GFileCopyFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_monitor_directory func(
		file *T.GFile,
		flags T.GFileMonitorFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileMonitor

	G_file_monitor_file func(
		file *T.GFile,
		flags T.GFileMonitorFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileMonitor

	G_file_monitor func(
		file *T.GFile,
		flags T.GFileMonitorFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileMonitor

	G_file_start_mountable func(
		file *T.GFile,
		flags T.GDriveStartFlags,
		start_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_start_mountable_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_stop_mountable func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_stop_mountable_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_poll_mountable func(
		file *T.GFile,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_poll_mountable_finish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_file_query_default_handler func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GAppInfo

	G_file_load_contents func(
		file *T.GFile,
		cancellable *T.GCancellable,
		contents **T.Char,
		length *T.Gsize,
		etag_out **T.Char,
		err **T.GError) T.Gboolean

	G_file_load_contents_async func(
		file *T.GFile,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_load_contents_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		contents **T.Char,
		length *T.Gsize,
		etag_out **T.Char,
		err **T.GError) T.Gboolean

	G_file_load_partial_contents_async func(
		file *T.GFile,
		cancellable *T.GCancellable,
		read_more_callback T.GFileReadMoreCallback,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_load_partial_contents_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		contents **T.Char,
		length *T.Gsize,
		etag_out **T.Char,
		err **T.GError) T.Gboolean

	G_file_replace_contents func(
		file *T.GFile,
		contents string,
		length T.Gsize,
		etag string,
		make_backup T.Gboolean,
		flags T.GFileCreateFlags,
		new_etag **T.Char,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_file_replace_contents_async func(
		file *T.GFile,
		contents string,
		length T.Gsize,
		etag string,
		make_backup T.Gboolean,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_replace_contents_finish func(
		file *T.GFile,
		res *T.GAsyncResult,
		new_etag **T.Char,
		err **T.GError) T.Gboolean

	G_file_supports_thread_contexts func(file *T.GFile) T.Gboolean

	G_file_icon_get_type func() T.GType

	G_file_icon_new func(file *T.GFile) *T.GIcon

	G_file_icon_get_file func(icon *T.GFileIcon) *T.GFile

	G_file_info_get_type func() T.GType

	G_file_info_new func() *T.GFileInfo

	G_file_info_dup func(other *T.GFileInfo) *T.GFileInfo

	G_file_info_copy_into func(src_info, dest_info *T.GFileInfo)

	G_file_info_has_attribute func(
		info *T.GFileInfo, attribute string) T.Gboolean

	G_file_info_has_namespace func(
		info *T.GFileInfo, name_space string) T.Gboolean

	G_file_info_list_attributes func(
		info *T.GFileInfo, name_space string) **T.Char

	G_file_info_get_attribute_data func(
		info *T.GFileInfo,
		attribute string,
		typ *T.GFileAttributeType,
		value_pp *T.Gpointer,
		status *T.GFileAttributeStatus) T.Gboolean

	G_file_info_get_attribute_type func(
		info *T.GFileInfo,
		attribute string) T.GFileAttributeType

	G_file_info_remove_attribute func(
		info *T.GFileInfo,
		attribute string)

	G_file_info_get_attribute_status func(
		info *T.GFileInfo,
		attribute string) T.GFileAttributeStatus

	G_file_info_set_attribute_status func(
		info *T.GFileInfo,
		attribute string,
		status T.GFileAttributeStatus) T.Gboolean

	G_file_info_get_attribute_as_string func(
		info *T.GFileInfo,
		attribute string) string

	G_file_info_get_attribute_string func(
		info *T.GFileInfo,
		attribute string) string

	G_file_info_get_attribute_byte_string func(
		info *T.GFileInfo,
		attribute string) string

	G_file_info_get_attribute_boolean func(
		info *T.GFileInfo,
		attribute string) T.Gboolean

	G_file_info_get_attribute_uint32 func(
		info *T.GFileInfo,
		attribute string) T.Guint32

	G_file_info_get_attribute_int32 func(
		info *T.GFileInfo,
		attribute string) T.Gint32

	G_file_info_get_attribute_uint64 func(
		info *T.GFileInfo,
		attribute string) T.Guint64

	G_file_info_get_attribute_int64 func(
		info *T.GFileInfo,
		attribute string) T.Gint64

	G_file_info_get_attribute_object func(
		info *T.GFileInfo,
		attribute string) *T.GObject

	G_file_info_get_attribute_stringv func(
		info *T.GFileInfo,
		attribute string) **T.Char

	G_file_info_set_attribute func(
		info *T.GFileInfo,
		attribute string,
		typ T.GFileAttributeType,
		value_p T.Gpointer)

	G_file_info_set_attribute_string func(
		info *T.GFileInfo,
		attribute string,
		attr_value string)

	G_file_info_set_attribute_byte_string func(
		info *T.GFileInfo,
		attribute string,
		attr_value string)

	G_file_info_set_attribute_boolean func(
		info *T.GFileInfo,
		attribute string,
		attr_value T.Gboolean)

	G_file_info_set_attribute_uint32 func(
		info *T.GFileInfo,
		attribute string,
		attr_value T.Guint32)

	G_file_info_set_attribute_int32 func(
		info *T.GFileInfo,
		attribute string,
		attr_value T.Gint32)

	G_file_info_set_attribute_uint64 func(
		info *T.GFileInfo,
		attribute string,
		attr_value T.Guint64)

	G_file_info_set_attribute_int64 func(
		info *T.GFileInfo,
		attribute string,
		attr_value T.Gint64)

	G_file_info_set_attribute_object func(
		info *T.GFileInfo,
		attribute string,
		attr_value *T.GObject)

	G_file_info_set_attribute_stringv func(
		info *T.GFileInfo,
		attribute string,
		attr_value **T.Char)

	G_file_info_clear_status func(
		info *T.GFileInfo)

	G_file_info_get_file_type func(
		info *T.GFileInfo) T.GFileType

	G_file_info_get_is_hidden func(
		info *T.GFileInfo) T.Gboolean

	G_file_info_get_is_backup func(
		info *T.GFileInfo) T.Gboolean

	G_file_info_get_is_symlink func(
		info *T.GFileInfo) T.Gboolean

	G_file_info_get_name func(
		info *T.GFileInfo) string

	G_file_info_get_display_name func(
		info *T.GFileInfo) string

	G_file_info_get_edit_name func(
		info *T.GFileInfo) string

	G_file_info_get_icon func(
		info *T.GFileInfo) *T.GIcon

	G_file_info_get_content_type func(
		info *T.GFileInfo) string

	G_file_info_get_size func(
		info *T.GFileInfo) T.Goffset

	G_file_info_get_modification_time func(
		info *T.GFileInfo,
		result *T.GTimeVal)

	G_file_info_get_symlink_target func(
		info *T.GFileInfo) string

	G_file_info_get_etag func(
		info *T.GFileInfo) string

	G_file_info_get_sort_order func(
		info *T.GFileInfo) T.Gint32

	G_file_info_set_attribute_mask func(
		info *T.GFileInfo,
		mask *T.GFileAttributeMatcher)

	G_file_info_unset_attribute_mask func(
		info *T.GFileInfo)

	G_file_info_set_file_type func(
		info *T.GFileInfo,
		typ T.GFileType)

	G_file_info_set_is_hidden func(
		info *T.GFileInfo,
		is_hidden T.Gboolean)

	G_file_info_set_is_symlink func(
		info *T.GFileInfo,
		is_symlink T.Gboolean)

	G_file_info_set_name func(
		info *T.GFileInfo,
		name string)

	G_file_info_set_display_name func(
		info *T.GFileInfo,
		display_name string)

	G_file_info_set_edit_name func(
		info *T.GFileInfo,
		edit_name string)

	G_file_info_set_icon func(
		info *T.GFileInfo,
		icon *T.GIcon)

	G_file_info_set_content_type func(
		info *T.GFileInfo,
		content_type string)

	G_file_info_set_size func(
		info *T.GFileInfo,
		size T.Goffset)

	G_file_info_set_modification_time func(
		info *T.GFileInfo,
		mtime *T.GTimeVal)

	G_file_info_set_symlink_target func(
		info *T.GFileInfo,
		symlink_target string)

	G_file_info_set_sort_order func(
		info *T.GFileInfo,
		sort_order T.Gint32)

	G_file_attribute_matcher_get_type func() T.GType

	G_file_attribute_matcher_new func(
		attributes string) *T.GFileAttributeMatcher

	G_file_attribute_matcher_ref func(
		matcher *T.GFileAttributeMatcher) *T.GFileAttributeMatcher

	G_file_attribute_matcher_unref func(
		matcher *T.GFileAttributeMatcher)

	G_file_attribute_matcher_matches func(
		matcher *T.GFileAttributeMatcher,
		attribute string) T.Gboolean

	G_file_attribute_matcher_matches_only func(
		matcher *T.GFileAttributeMatcher,
		attribute string) T.Gboolean

	G_file_attribute_matcher_enumerate_namespace func(
		matcher *T.GFileAttributeMatcher,
		ns string) T.Gboolean

	G_file_attribute_matcher_enumerate_next func(
		matcher *T.GFileAttributeMatcher) string

	G_file_input_stream_get_type func() T.GType

	G_file_input_stream_query_info func(
		stream *T.GFileInputStream,
		attributes string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	G_file_input_stream_query_info_async func(
		stream *T.GFileInputStream,
		attributes string,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_input_stream_query_info_finish func(
		stream *T.GFileInputStream,
		result *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	G_io_error_quark func() T.GQuark

	G_io_error_from_errno func(
		err_no T.Gint) T.GIOErrorEnum

	G_io_error_from_win32_error func(
		error_code T.Gint) T.GIOErrorEnum

	G_io_stream_get_type func() T.GType

	G_io_stream_get_input_stream func(
		stream *T.GIOStream) *T.GInputStream

	G_io_stream_get_output_stream func(
		stream *T.GIOStream) *T.GOutputStream

	G_io_stream_splice_async func(
		stream1 *T.GIOStream,
		stream2 *T.GIOStream,
		flags T.GIOStreamSpliceFlags,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_io_stream_splice_finish func(
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_io_stream_close func(
		stream *T.GIOStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_io_stream_close_async func(
		stream *T.GIOStream,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_io_stream_close_finish func(
		stream *T.GIOStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_io_stream_is_closed func(
		stream *T.GIOStream) T.Gboolean

	G_io_stream_has_pending func(
		stream *T.GIOStream) T.Gboolean

	G_io_stream_set_pending func(
		stream *T.GIOStream,
		err **T.GError) T.Gboolean

	G_io_stream_clear_pending func(
		stream *T.GIOStream)

	G_file_io_stream_get_type func() T.GType

	G_file_io_stream_query_info func(
		stream *T.GFileIOStream,
		attributes string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	G_file_io_stream_query_info_async func(
		stream *T.GFileIOStream,
		attributes string,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_io_stream_query_info_finish func(
		stream *T.GFileIOStream,
		result *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	G_file_io_stream_get_etag func(
		stream *T.GFileIOStream) string

	G_file_monitor_get_type func() T.GType

	G_file_monitor_cancel func(
		monitor *T.GFileMonitor) T.Gboolean

	G_file_monitor_is_cancelled func(
		monitor *T.GFileMonitor) T.Gboolean

	G_file_monitor_set_rate_limit func(
		monitor *T.GFileMonitor,
		limit_msecs T.Gint)

	G_file_monitor_emit_event func(
		monitor *T.GFileMonitor,
		child *T.GFile,
		other_file *T.GFile,
		event_type T.GFileMonitorEvent)

	G_filename_completer_get_type func() T.GType

	G_filename_completer_new func() *T.GFilenameCompleter

	G_filename_completer_get_completion_suffix func(
		completer *T.GFilenameCompleter,
		initial_text string) string

	G_filename_completer_get_completions func(
		completer *T.GFilenameCompleter,
		initial_text string) **T.Char

	G_filename_completer_set_dirs_only func(
		completer *T.GFilenameCompleter,
		dirs_only T.Gboolean)

	G_file_output_stream_get_type func() T.GType

	G_file_output_stream_query_info func(
		stream *T.GFileOutputStream,
		attributes string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	G_file_output_stream_query_info_async func(
		stream *T.GFileOutputStream,
		attributes string,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_file_output_stream_query_info_finish func(
		stream *T.GFileOutputStream,
		result *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	G_file_output_stream_get_etag func(
		stream *T.GFileOutputStream) string

	G_inet_address_get_type func() T.GType

	G_inet_address_new_from_string func(
		string string) *T.GInetAddress

	G_inet_address_new_from_bytes func(
		bytes *T.Guint8,
		family T.GSocketFamily) *T.GInetAddress

	G_inet_address_new_loopback func(
		family T.GSocketFamily) *T.GInetAddress

	G_inet_address_new_any func(
		family T.GSocketFamily) *T.GInetAddress

	G_inet_address_to_string func(
		address *T.GInetAddress) string

	G_inet_address_to_bytes func(
		address *T.GInetAddress) *T.Guint8

	G_inet_address_get_native_size func(
		address *T.GInetAddress) T.Gsize

	G_inet_address_get_family func(
		address *T.GInetAddress) T.GSocketFamily

	G_inet_address_get_is_any func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_loopback func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_link_local func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_site_local func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_multicast func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_mc_global func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_mc_link_local func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_mc_node_local func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_mc_org_local func(
		address *T.GInetAddress) T.Gboolean

	G_inet_address_get_is_mc_site_local func(
		address *T.GInetAddress) T.Gboolean

	G_socket_address_get_type func() T.GType

	G_socket_address_get_family func(
		address *T.GSocketAddress) T.GSocketFamily

	G_socket_address_new_from_native func(
		native T.Gpointer,
		len T.Gsize) *T.GSocketAddress

	G_socket_address_to_native func(
		address *T.GSocketAddress,
		dest T.Gpointer,
		destlen T.Gsize,
		err **T.GError) T.Gboolean

	G_socket_address_get_native_size func(
		address *T.GSocketAddress) T.Gssize

	G_inet_socket_address_get_type func() T.GType

	G_inet_socket_address_new func(
		address *T.GInetAddress,
		port T.Guint16) *T.GSocketAddress

	G_inet_socket_address_get_address func(
		address *T.GInetSocketAddress) *T.GInetAddress

	G_inet_socket_address_get_port func(
		address *T.GInetSocketAddress) T.Guint16

	G_app_info_create_flags_get_type func() T.GType

	G_converter_flags_get_type func() T.GType

	G_converter_result_get_type func() T.GType

	G_data_stream_byte_order_get_type func() T.GType

	G_data_stream_newline_type_get_type func() T.GType

	G_file_attribute_type_get_type func() T.GType

	G_file_attribute_info_flags_get_type func() T.GType

	G_file_attribute_status_get_type func() T.GType

	G_file_query_info_flags_get_type func() T.GType

	G_file_create_flags_get_type func() T.GType

	G_mount_mount_flags_get_type func() T.GType

	G_mount_unmount_flags_get_type func() T.GType

	G_drive_start_flags_get_type func() T.GType

	G_drive_start_stop_type_get_type func() T.GType

	G_file_copy_flags_get_type func() T.GType

	G_file_monitor_flags_get_type func() T.GType

	G_file_type_get_type func() T.GType

	G_filesystem_preview_type_get_type func() T.GType

	G_file_monitor_event_get_type func() T.GType

	G_io_error_enum_get_type func() T.GType

	G_ask_password_flags_get_type func() T.GType

	G_password_save_get_type func() T.GType

	G_mount_operation_result_get_type func() T.GType

	G_output_stream_splice_flags_get_type func() T.GType

	G_io_stream_splice_flags_get_type func() T.GType

	G_emblem_origin_get_type func() T.GType

	G_resolver_error_get_type func() T.GType

	G_socket_family_get_type func() T.GType

	G_socket_type_get_type func() T.GType

	G_socket_msg_flags_get_type func() T.GType

	G_socket_protocol_get_type func() T.GType

	G_zlib_compressor_format_get_type func() T.GType

	G_unix_socket_address_type_get_type func() T.GType

	G_bus_type_get_type func() T.GType

	G_bus_name_owner_flags_get_type func() T.GType

	G_bus_name_watcher_flags_get_type func() T.GType

	G_dbus_proxy_flags_get_type func() T.GType

	G_dbus_error_get_type func() T.GType

	G_dbus_connection_flags_get_type func() T.GType

	G_dbus_capability_flags_get_type func() T.GType

	G_dbus_call_flags_get_type func() T.GType

	G_dbus_message_type_get_type func() T.GType

	G_dbus_message_flags_get_type func() T.GType

	G_dbus_message_header_field_get_type func() T.GType

	G_dbus_property_info_flags_get_type func() T.GType

	G_dbus_subtree_flags_get_type func() T.GType

	G_dbus_server_flags_get_type func() T.GType

	G_dbus_signal_flags_get_type func() T.GType

	G_dbus_send_message_flags_get_type func() T.GType

	G_credentials_type_get_type func() T.GType

	G_dbus_message_byte_order_get_type func() T.GType

	G_application_flags_get_type func() T.GType

	G_tls_error_get_type func() T.GType

	G_tls_certificate_flags_get_type func() T.GType

	G_tls_authentication_mode_get_type func() T.GType

	G_tls_rehandshake_mode_get_type func() T.GType

	G_settings_bind_flags_get_type func() T.GType

	G_io_module_get_type func() T.GType

	G_io_module_new func(
		filename string) *T.GIOModule

	G_io_modules_scan_all_in_directory func(
		dirname string)

	G_io_modules_load_all_in_directory func(
		dirname string) *T.GList

	G_io_extension_point_register func(
		name string) *T.GIOExtensionPoint

	G_io_extension_point_lookup func(
		name string) *T.GIOExtensionPoint

	G_io_extension_point_set_required_type func(
		extension_point *T.GIOExtensionPoint,
		typ T.GType)

	G_io_extension_point_get_required_type func(
		extension_point *T.GIOExtensionPoint) T.GType

	G_io_extension_point_get_extensions func(
		extension_point *T.GIOExtensionPoint) *T.GList

	G_io_extension_point_get_extension_by_name func(
		extension_point *T.GIOExtensionPoint,
		name string) *T.GIOExtension

	G_io_extension_point_implement func(
		extension_point_name string,
		typ T.GType,
		extension_name string,
		priority T.Gint) *T.GIOExtension

	G_io_extension_get_type func(
		extension *T.GIOExtension) T.GType

	G_io_extension_get_name func(
		extension *T.GIOExtension) string

	G_io_extension_get_priority func(
		extension *T.GIOExtension) T.Gint

	G_io_extension_ref_class func(
		extension *T.GIOExtension) *T.GTypeClass

	G_io_module_load func(
		module *T.GIOModule)

	G_io_module_unload func(
		module *T.GIOModule)

	G_io_module_query func() **T.Char

	G_io_scheduler_push_job func(
		job_func T.GIOSchedulerJobFunc,
		user_data T.Gpointer,
		notify T.GDestroyNotify,
		io_priority T.Gint,
		cancellable *T.GCancellable)

	G_io_scheduler_cancel_all_jobs func()

	G_io_scheduler_job_send_to_mainloop func(
		job *T.GIOSchedulerJob,
		f T.GSourceFunc,
		user_data T.Gpointer,
		notify T.GDestroyNotify) T.Gboolean

	G_io_scheduler_job_send_to_mainloop_async func(
		job *T.GIOSchedulerJob,
		f T.GSourceFunc,
		user_data T.Gpointer,
		notify T.GDestroyNotify)

	G_loadable_icon_get_type func() T.GType

	G_loadable_icon_load func(
		icon *T.GLoadableIcon,
		size int,
		t **T.Char,
		cancellable *T.GCancellable,
		err **T.GError) *T.GInputStream

	G_loadable_icon_load_async func(
		icon *T.GLoadableIcon,
		size int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_loadable_icon_load_finish func(
		icon *T.GLoadableIcon,
		res *T.GAsyncResult,
		typ **T.Char,
		err **T.GError) *T.GInputStream

	G_memory_input_stream_get_type func() T.GType

	G_memory_input_stream_new func() *T.GInputStream

	G_memory_input_stream_new_from_data func(
		data *T.Void,
		len T.Gssize,
		destroy T.GDestroyNotify) *T.GInputStream

	G_memory_input_stream_add_data func(
		stream *T.GMemoryInputStream,
		data *T.Void,
		len T.Gssize,
		destroy T.GDestroyNotify)

	G_memory_output_stream_get_type func() T.GType

	G_memory_output_stream_new func(
		data T.Gpointer,
		size T.Gsize,
		realloc_function T.GReallocFunc,
		destroy_function T.GDestroyNotify) *T.GOutputStream

	G_memory_output_stream_get_data func(
		ostream *T.GMemoryOutputStream) T.Gpointer

	G_memory_output_stream_get_size func(
		ostream *T.GMemoryOutputStream) T.Gsize

	G_memory_output_stream_get_data_size func(
		ostream *T.GMemoryOutputStream) T.Gsize

	G_memory_output_stream_steal_data func(
		ostream *T.GMemoryOutputStream) T.Gpointer

	G_mount_get_type func() T.GType

	G_mount_get_root func(
		mount *T.GMount) *T.GFile

	G_mount_get_default_location func(
		mount *T.GMount) *T.GFile

	G_mount_get_name func(
		mount *T.GMount) string

	G_mount_get_icon func(
		mount *T.GMount) *T.GIcon

	G_mount_get_uuid func(
		mount *T.GMount) string

	G_mount_get_volume func(
		mount *T.GMount) *T.GVolume

	G_mount_get_drive func(
		mount *T.GMount) *T.GDrive

	G_mount_can_unmount func(
		mount *T.GMount) T.Gboolean

	G_mount_can_eject func(
		mount *T.GMount) T.Gboolean

	G_mount_unmount func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_mount_unmount_finish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_mount_eject func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_mount_eject_finish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_mount_remount func(
		mount *T.GMount,
		flags T.GMountMountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_mount_remount_finish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_mount_guess_content_type func(
		mount *T.GMount,
		force_rescan T.Gboolean,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_mount_guess_content_type_finish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) **T.Gchar

	G_mount_guess_content_type_sync func(
		mount *T.GMount,
		force_rescan T.Gboolean,
		cancellable *T.GCancellable,
		err **T.GError) **T.Gchar

	G_mount_is_shadowed func(
		mount *T.GMount) T.Gboolean

	G_mount_shadow func(
		mount *T.GMount)

	G_mount_unshadow func(
		mount *T.GMount)

	G_mount_unmount_with_operation func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_mount_unmount_with_operation_finish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_mount_eject_with_operation func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_mount_eject_with_operation_finish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_mount_operation_get_type func() T.GType

	G_mount_operation_new func() *T.GMountOperation

	G_mount_operation_get_username func(
		op *T.GMountOperation) string

	G_mount_operation_set_username func(
		op *T.GMountOperation,
		username string)

	G_mount_operation_get_password func(
		op *T.GMountOperation) string

	G_mount_operation_set_password func(
		op *T.GMountOperation,
		password string)

	G_mount_operation_get_anonymous func(
		op *T.GMountOperation) T.Gboolean

	G_mount_operation_set_anonymous func(
		op *T.GMountOperation,
		anonymous T.Gboolean)

	G_mount_operation_get_domain func(
		op *T.GMountOperation) string

	G_mount_operation_set_domain func(
		op *T.GMountOperation,
		domain string)

	G_mount_operation_get_password_save func(
		op *T.GMountOperation) T.GPasswordSave

	G_mount_operation_set_password_save func(
		op *T.GMountOperation,
		save T.GPasswordSave)

	G_mount_operation_get_choice func(
		op *T.GMountOperation) int

	G_mount_operation_set_choice func(
		op *T.GMountOperation,
		choice int)

	G_mount_operation_reply func(
		op *T.GMountOperation,
		result T.GMountOperationResult)

	G_volume_monitor_get_type func() T.GType

	G_volume_monitor_get func() *T.GVolumeMonitor

	G_volume_monitor_get_connected_drives func(
		volume_monitor *T.GVolumeMonitor) *T.GList

	G_volume_monitor_get_volumes func(
		volume_monitor *T.GVolumeMonitor) *T.GList

	G_volume_monitor_get_mounts func(
		volume_monitor *T.GVolumeMonitor) *T.GList

	G_volume_monitor_get_volume_for_uuid func(
		volume_monitor *T.GVolumeMonitor,
		uuid string) *T.GVolume

	G_volume_monitor_get_mount_for_uuid func(
		volume_monitor *T.GVolumeMonitor,
		uuid string) *T.GMount

	G_volume_monitor_adopt_orphan_mount func(
		mount *T.GMount) *T.GVolume

	G_native_volume_monitor_get_type func() T.GType

	G_network_address_get_type func() T.GType

	G_network_address_new func(
		hostname string,
		port T.Guint16) *T.GSocketConnectable

	G_network_address_parse func(
		host_and_port string,
		default_port T.Guint16,
		err **T.GError) *T.GSocketConnectable

	G_network_address_parse_uri func(
		uri string,
		default_port T.Guint16,
		err **T.GError) *T.GSocketConnectable

	G_network_address_get_hostname func(
		addr *T.GNetworkAddress) string

	G_network_address_get_port func(
		addr *T.GNetworkAddress) T.Guint16

	G_network_address_get_scheme func(
		addr *T.GNetworkAddress) string

	G_network_service_get_type func() T.GType

	G_network_service_new func(
		service string,
		protocol string,
		domain string) *T.GSocketConnectable

	G_network_service_get_service func(
		srv *T.GNetworkService) string

	G_network_service_get_protocol func(
		srv *T.GNetworkService) string

	G_network_service_get_domain func(
		srv *T.GNetworkService) string

	G_network_service_get_scheme func(
		srv *T.GNetworkService) string

	G_network_service_set_scheme func(
		srv *T.GNetworkService,
		scheme string)

	G_permission_get_type func() T.GType

	G_permission_acquire func(
		permission *T.GPermission,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_permission_acquire_async func(
		permission *T.GPermission,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_permission_acquire_finish func(
		permission *T.GPermission,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_permission_release func(
		permission *T.GPermission,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_permission_release_async func(
		permission *T.GPermission,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_permission_release_finish func(
		permission *T.GPermission,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_permission_get_allowed func(
		permission *T.GPermission) T.Gboolean

	G_permission_get_can_acquire func(
		permission *T.GPermission) T.Gboolean

	G_permission_get_can_release func(
		permission *T.GPermission) T.Gboolean

	G_permission_impl_update func(
		permission *T.GPermission,
		allowed T.Gboolean,
		can_acquire T.Gboolean,
		can_release T.Gboolean)

	G_pollable_input_stream_get_type func() T.GType

	G_pollable_input_stream_can_poll func(
		stream *T.GPollableInputStream) T.Gboolean

	G_pollable_input_stream_is_readable func(
		stream *T.GPollableInputStream) T.Gboolean

	G_pollable_input_stream_create_source func(
		stream *T.GPollableInputStream,
		cancellable *T.GCancellable) *T.GSource

	G_pollable_input_stream_read_nonblocking func(
		stream *T.GPollableInputStream,
		buffer *T.Void,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_pollable_source_new func(
		pollable_stream *T.GObject) *T.GSource

	G_pollable_output_stream_get_type func() T.GType

	G_pollable_output_stream_can_poll func(
		stream *T.GPollableOutputStream) T.Gboolean

	G_pollable_output_stream_is_writable func(
		stream *T.GPollableOutputStream) T.Gboolean

	G_pollable_output_stream_create_source func(
		stream *T.GPollableOutputStream,
		cancellable *T.GCancellable) *T.GSource

	G_pollable_output_stream_write_nonblocking func(
		stream *T.GPollableOutputStream,
		buffer *T.Void,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_proxy_get_type func() T.GType

	G_proxy_get_default_for_protocol func(
		protocol string) *T.GProxy

	G_proxy_connect func(
		proxy *T.GProxy,
		connection *T.GIOStream,
		proxy_address *T.GProxyAddress,
		cancellable *T.GCancellable,
		err **T.GError) *T.GIOStream

	G_proxy_connect_async func(
		proxy *T.GProxy,
		connection *T.GIOStream,
		proxy_address *T.GProxyAddress,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_proxy_connect_finish func(
		proxy *T.GProxy,
		result *T.GAsyncResult,
		err **T.GError) *T.GIOStream

	G_proxy_supports_hostname func(
		proxy *T.GProxy) T.Gboolean

	G_proxy_address_get_type func() T.GType

	G_proxy_address_new func(
		inetaddr *T.GInetAddress,
		port T.Guint16,
		protocol string,
		dest_hostname string,
		dest_port T.Guint16,
		username string,
		password string) *T.GSocketAddress

	G_proxy_address_get_protocol func(
		proxy *T.GProxyAddress) string

	G_proxy_address_get_destination_hostname func(
		proxy *T.GProxyAddress) string

	G_proxy_address_get_destination_port func(
		proxy *T.GProxyAddress) T.Guint16

	G_proxy_address_get_username func(
		proxy *T.GProxyAddress) string

	G_proxy_address_get_password func(
		proxy *T.GProxyAddress) string

	G_socket_address_enumerator_get_type func() T.GType

	G_socket_address_enumerator_next func(
		enumerator *T.GSocketAddressEnumerator,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketAddress

	G_socket_address_enumerator_next_async func(
		enumerator *T.GSocketAddressEnumerator,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_socket_address_enumerator_next_finish func(
		enumerator *T.GSocketAddressEnumerator,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketAddress

	G_proxy_address_enumerator_get_type func() T.GType

	G_proxy_resolver_get_type func() T.GType

	G_proxy_resolver_get_default func() *T.GProxyResolver

	G_proxy_resolver_is_supported func(
		resolver *T.GProxyResolver) T.Gboolean

	G_proxy_resolver_lookup func(
		resolver *T.GProxyResolver,
		uri string,
		cancellable *T.GCancellable,
		err **T.GError) **T.Gchar

	G_proxy_resolver_lookup_async func(
		resolver *T.GProxyResolver,
		uri string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_proxy_resolver_lookup_finish func(
		resolver *T.GProxyResolver,
		result *T.GAsyncResult,
		err **T.GError) **T.Gchar

	G_resolver_get_type func() T.GType

	G_resolver_get_default func() *T.GResolver

	G_resolver_set_default func(
		resolver *T.GResolver)

	G_resolver_lookup_by_name func(
		resolver *T.GResolver,
		hostname string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GList

	G_resolver_lookup_by_name_async func(
		resolver *T.GResolver,
		hostname string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_resolver_lookup_by_name_finish func(
		resolver *T.GResolver,
		result *T.GAsyncResult,
		err **T.GError) *T.GList

	G_resolver_free_addresses func(
		addresses *T.GList)

	G_resolver_lookup_by_address func(
		resolver *T.GResolver,
		address *T.GInetAddress,
		cancellable *T.GCancellable,
		err **T.GError) string

	G_resolver_lookup_by_address_async func(
		resolver *T.GResolver,
		address *T.GInetAddress,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_resolver_lookup_by_address_finish func(
		resolver *T.GResolver,
		result *T.GAsyncResult,
		err **T.GError) string

	G_resolver_lookup_service func(
		resolver *T.GResolver,
		service string,
		protocol string,
		domain string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GList

	G_resolver_lookup_service_async func(
		resolver *T.GResolver,
		service string,
		protocol string,
		domain string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_resolver_lookup_service_finish func(
		resolver *T.GResolver,
		result *T.GAsyncResult,
		err **T.GError) *T.GList

	G_resolver_free_targets func(
		targets *T.GList)

	G_resolver_error_quark func() T.GQuark

	G_seekable_get_type func() T.GType

	G_seekable_tell func(
		seekable *T.GSeekable) T.Goffset

	G_seekable_can_seek func(
		seekable *T.GSeekable) T.Gboolean

	G_seekable_seek func(
		seekable *T.GSeekable,
		offset T.Goffset,
		typ T.GSeekType,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_seekable_can_truncate func(
		seekable *T.GSeekable) T.Gboolean

	G_seekable_truncate func(
		seekable *T.GSeekable,
		offset T.Goffset,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_settings_get_type func() T.GType

	G_settings_list_schemas func() **T.Gchar

	G_settings_list_relocatable_schemas func() **T.Gchar

	G_settings_new func(
		schema string) *T.GSettings

	G_settings_new_with_path func(
		schema string,
		path string) *T.GSettings

	G_settings_new_with_backend func(
		schema string,
		backend *T.GSettingsBackend) *T.GSettings

	G_settings_new_with_backend_and_path func(
		schema string,
		backend *T.GSettingsBackend,
		path string) *T.GSettings

	G_settings_list_children func(
		settings *T.GSettings) **T.Gchar

	G_settings_list_keys func(
		settings *T.GSettings) **T.Gchar

	G_settings_get_range func(
		settings *T.GSettings,
		key string) *T.GVariant

	G_settings_range_check func(
		settings *T.GSettings,
		key string,
		value *T.GVariant) T.Gboolean

	G_settings_set_value func(
		settings *T.GSettings,
		key string,
		value *T.GVariant) T.Gboolean

	G_settings_get_value func(
		settings *T.GSettings,
		key string) *T.GVariant

	G_settings_set func(settings *T.GSettings, key, format string,
		v ...VArg) T.Gboolean

	G_settings_get func(settings *T.GSettings, key, format string,
		v ...VArg)

	G_settings_reset func(settings *T.GSettings, key string)

	G_settings_get_int func(settings *T.GSettings, key string) T.Gint

	G_settings_set_int func(
		settings *T.GSettings, key string, value T.Gint) T.Gboolean

	G_settings_get_string func(
		settings *T.GSettings, key string) string

	G_settings_set_string func(
		settings *T.GSettings, key, value string) T.Gboolean

	G_settings_get_boolean func(
		settings *T.GSettings, key string) T.Gboolean

	G_settings_set_boolean func(
		settings *T.GSettings, key string, value T.Gboolean) T.Gboolean

	G_settings_get_double func(
		settings *T.GSettings, key string) T.Gdouble

	G_settings_set_double func(settings *T.GSettings,
		key string, value T.Gdouble) T.Gboolean

	G_settings_get_strv func(
		settings *T.GSettings, key string) **T.Gchar

	G_settings_set_strv func(settings *T.GSettings,
		key string, value **T.Gchar) T.Gboolean

	G_settings_get_enum func(
		settings *T.GSettings, key string) T.Gint

	G_settings_set_enum func(settings *T.GSettings,
		key string, value T.Gint) T.Gboolean

	G_settings_get_flags func(
		settings *T.GSettings, key string) T.Guint

	G_settings_set_flags func(settings *T.GSettings,
		key string, value T.Guint) T.Gboolean

	G_settings_get_child func(
		settings *T.GSettings,
		name string) *T.GSettings

	G_settings_is_writable func(
		settings *T.GSettings, name string) T.Gboolean

	G_settings_delay func(settings *T.GSettings)

	G_settings_apply func(settings *T.GSettings)

	G_settings_revert func(settings *T.GSettings)

	G_settings_get_has_unapplied func(
		settings *T.GSettings) T.Gboolean

	G_settings_sync func()

	G_settings_bind func(
		settings *T.GSettings,
		key string,
		object T.Gpointer,
		property string,
		flags T.GSettingsBindFlags)

	G_settings_bind_with_mapping func(
		settings *T.GSettings,
		key string,
		object T.Gpointer,
		property string,
		flags T.GSettingsBindFlags,
		get_mapping T.GSettingsBindGetMapping,
		set_mapping T.GSettingsBindSetMapping,
		user_data T.Gpointer,
		destroy T.GDestroyNotify)

	G_settings_bind_writable func(
		settings *T.GSettings,
		key string,
		object T.Gpointer,
		property string,
		inverted T.Gboolean)

	G_settings_unbind func(object T.Gpointer, property string)

	G_settings_get_mapped func(
		settings *T.GSettings,
		key string,
		mapping T.GSettingsGetMapping,
		user_data T.Gpointer) T.Gpointer

	G_simple_async_result_get_type func() T.GType

	G_simple_async_result_new func(
		source_object *T.GObject,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer,
		source_tag T.Gpointer) *T.GSimpleAsyncResult

	G_simple_async_result_new_error func(
		source_object *T.GObject, callback T.GAsyncReadyCallback,
		user_data T.Gpointer, domain T.GQuark, code T.Gint,
		format string, v ...VArg) *T.GSimpleAsyncResult

	G_simple_async_result_new_from_error func(
		source_object *T.GObject,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer,
		err *T.GError) *T.GSimpleAsyncResult

	G_simple_async_result_new_take_error func(
		source_object *T.GObject,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer,
		err *T.GError) *T.GSimpleAsyncResult

	G_simple_async_result_set_op_res_gpointer func(
		simple *T.GSimpleAsyncResult,
		op_res T.Gpointer,
		destroy_op_res T.GDestroyNotify)

	G_simple_async_result_get_op_res_gpointer func(
		simple *T.GSimpleAsyncResult) T.Gpointer

	G_simple_async_result_set_op_res_gssize func(
		simple *T.GSimpleAsyncResult,
		op_res T.Gssize)

	G_simple_async_result_get_op_res_gssize func(
		simple *T.GSimpleAsyncResult) T.Gssize

	G_simple_async_result_set_op_res_gboolean func(
		simple *T.GSimpleAsyncResult,
		op_res T.Gboolean)

	G_simple_async_result_get_op_res_gboolean func(
		simple *T.GSimpleAsyncResult) T.Gboolean

	G_simple_async_result_get_source_tag func(
		simple *T.GSimpleAsyncResult) T.Gpointer

	G_simple_async_result_set_handle_cancellation func(
		simple *T.GSimpleAsyncResult,
		handle_cancellation T.Gboolean)

	G_simple_async_result_complete func(
		simple *T.GSimpleAsyncResult)

	G_simple_async_result_complete_in_idle func(
		simple *T.GSimpleAsyncResult)

	G_simple_async_result_run_in_thread func(
		simple *T.GSimpleAsyncResult,
		f T.GSimpleAsyncThreadFunc,
		io_priority int,
		cancellable *T.GCancellable)

	G_simple_async_result_set_from_error func(
		simple *T.GSimpleAsyncResult,
		err *T.GError)

	G_simple_async_result_take_error func(
		simple *T.GSimpleAsyncResult,
		err *T.GError)

	G_simple_async_result_propagate_error func(
		simple *T.GSimpleAsyncResult,
		dest **T.GError) T.Gboolean

	G_simple_async_result_set_error func(
		simple *T.GSimpleAsyncResult, domain T.GQuark,
		code T.Gint, format string, v ...VArg)

	G_simple_async_result_set_error_va func(
		simple *T.GSimpleAsyncResult,
		domain T.GQuark,
		code T.Gint,
		format string,
		args T.Va_list)

	G_simple_async_result_is_valid func(
		result *T.GAsyncResult,
		source *T.GObject,
		source_tag T.Gpointer) T.Gboolean

	G_simple_async_report_error_in_idle func(object *T.GObject,
		callback T.GAsyncReadyCallback, user_data T.Gpointer,
		domain T.GQuark, code T.Gint, format string, v ...VArg)

	G_simple_async_report_gerror_in_idle func(
		object *T.GObject,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer,
		err *T.GError)

	G_simple_async_report_take_gerror_in_idle func(
		object *T.GObject,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer,
		err *T.GError)

	G_simple_permission_get_type func() T.GType

	G_simple_permission_new func(
		allowed T.Gboolean) *T.GPermission

	G_socket_client_get_type func() T.GType

	G_socket_client_new func() *T.GSocketClient

	G_socket_client_get_family func(
		client *T.GSocketClient) T.GSocketFamily

	G_socket_client_set_family func(
		client *T.GSocketClient,
		family T.GSocketFamily)

	G_socket_client_get_socket_type func(
		client *T.GSocketClient) T.GSocketType

	G_socket_client_set_socket_type func(
		client *T.GSocketClient,
		t T.GSocketType)

	G_socket_client_get_protocol func(
		client *T.GSocketClient) T.GSocketProtocol

	G_socket_client_set_protocol func(
		client *T.GSocketClient,
		protocol T.GSocketProtocol)

	G_socket_client_get_local_address func(
		client *T.GSocketClient) *T.GSocketAddress

	G_socket_client_set_local_address func(
		client *T.GSocketClient,
		address *T.GSocketAddress)

	G_socket_client_get_timeout func(
		client *T.GSocketClient) T.Guint

	G_socket_client_set_timeout func(
		client *T.GSocketClient,
		timeout T.Guint)

	G_socket_client_get_enable_proxy func(
		client *T.GSocketClient) T.Gboolean

	G_socket_client_set_enable_proxy func(
		client *T.GSocketClient,
		enable T.Gboolean)

	G_socket_client_get_tls func(
		client *T.GSocketClient) T.Gboolean

	G_socket_client_set_tls func(
		client *T.GSocketClient,
		tls T.Gboolean)

	G_socket_client_get_tls_validation_flags func(
		client *T.GSocketClient) T.GTlsCertificateFlags

	G_socket_client_set_tls_validation_flags func(
		client *T.GSocketClient,
		flags T.GTlsCertificateFlags)

	G_socket_client_connect func(
		client *T.GSocketClient,
		connectable *T.GSocketConnectable,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	G_socket_client_connect_to_host func(
		client *T.GSocketClient,
		host_and_port string,
		default_port T.Guint16,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	G_socket_client_connect_to_service func(
		client *T.GSocketClient,
		domain string,
		service string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	G_socket_client_connect_to_uri func(
		client *T.GSocketClient,
		uri string,
		default_port T.Guint16,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	G_socket_client_connect_async func(
		client *T.GSocketClient,
		connectable *T.GSocketConnectable,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_socket_client_connect_finish func(
		client *T.GSocketClient,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketConnection

	G_socket_client_connect_to_host_async func(
		client *T.GSocketClient,
		host_and_port string,
		default_port T.Guint16,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_socket_client_connect_to_host_finish func(
		client *T.GSocketClient,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketConnection

	G_socket_client_connect_to_service_async func(
		client *T.GSocketClient,
		domain string,
		service string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_socket_client_connect_to_service_finish func(
		client *T.GSocketClient,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketConnection

	G_socket_client_connect_to_uri_async func(
		client *T.GSocketClient,
		uri string,
		default_port T.Guint16,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_socket_client_connect_to_uri_finish func(
		client *T.GSocketClient,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketConnection

	G_socket_client_add_application_proxy func(
		client *T.GSocketClient,
		protocol string)

	G_socket_connectable_get_type func() T.GType

	G_socket_connectable_enumerate func(
		connectable *T.GSocketConnectable) *T.GSocketAddressEnumerator

	G_socket_connectable_proxy_enumerate func(
		connectable *T.GSocketConnectable) *T.GSocketAddressEnumerator

	G_socket_get_type func() T.GType

	G_socket_new func(
		family T.GSocketFamily,
		t T.GSocketType,
		protocol T.GSocketProtocol,
		err **T.GError) *T.GSocket

	G_socket_new_from_fd func(
		fd T.Gint,
		err **T.GError) *T.GSocket

	G_socket_get_fd func(
		socket *T.GSocket) int

	G_socket_get_family func(
		socket *T.GSocket) T.GSocketFamily

	G_socket_get_socket_type func(
		socket *T.GSocket) T.GSocketType

	G_socket_get_protocol func(
		socket *T.GSocket) T.GSocketProtocol

	G_socket_get_local_address func(
		socket *T.GSocket,
		err **T.GError) *T.GSocketAddress

	G_socket_get_remote_address func(
		socket *T.GSocket,
		err **T.GError) *T.GSocketAddress

	G_socket_set_blocking func(
		socket *T.GSocket,
		blocking T.Gboolean)

	G_socket_get_blocking func(
		socket *T.GSocket) T.Gboolean

	G_socket_set_keepalive func(
		socket *T.GSocket,
		keepalive T.Gboolean)

	G_socket_get_keepalive func(
		socket *T.GSocket) T.Gboolean

	G_socket_get_listen_backlog func(
		socket *T.GSocket) T.Gint

	G_socket_set_listen_backlog func(
		socket *T.GSocket,
		backlog T.Gint)

	G_socket_get_timeout func(
		socket *T.GSocket) T.Guint

	G_socket_set_timeout func(
		socket *T.GSocket,
		timeout T.Guint)

	G_socket_is_connected func(
		socket *T.GSocket) T.Gboolean

	G_socket_bind func(
		socket *T.GSocket,
		address *T.GSocketAddress,
		allow_reuse T.Gboolean,
		err **T.GError) T.Gboolean

	G_socket_connect func(
		socket *T.GSocket,
		address *T.GSocketAddress,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_socket_check_connect_result func(
		socket *T.GSocket,
		err **T.GError) T.Gboolean

	G_socket_condition_check func(
		socket *T.GSocket,
		condition T.GIOCondition) T.GIOCondition

	G_socket_condition_wait func(
		socket *T.GSocket,
		condition T.GIOCondition,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_socket_accept func(
		socket *T.GSocket,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocket

	G_socket_listen func(
		socket *T.GSocket,
		err **T.GError) T.Gboolean

	G_socket_receive func(
		socket *T.GSocket,
		buffer string,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_socket_receive_from func(
		socket *T.GSocket,
		address **T.GSocketAddress,
		buffer string,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_socket_send func(
		socket *T.GSocket,
		buffer string,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_socket_send_to func(
		socket *T.GSocket,
		address *T.GSocketAddress,
		buffer string,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_socket_receive_message func(
		socket *T.GSocket,
		address **T.GSocketAddress,
		vectors *T.GInputVector,
		num_vectors T.Gint,
		messages ***T.GSocketControlMessage,
		num_messages *T.Gint,
		flags *T.Gint,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_socket_send_message func(
		socket *T.GSocket,
		address *T.GSocketAddress,
		vectors *T.GOutputVector,
		num_vectors T.Gint,
		messages **T.GSocketControlMessage,
		num_messages T.Gint,
		flags T.Gint,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_socket_close func(
		socket *T.GSocket,
		err **T.GError) T.Gboolean

	G_socket_shutdown func(
		socket *T.GSocket,
		shutdown_read T.Gboolean,
		shutdown_write T.Gboolean,
		err **T.GError) T.Gboolean

	G_socket_is_closed func(
		socket *T.GSocket) T.Gboolean

	G_socket_create_source func(
		socket *T.GSocket,
		condition T.GIOCondition,
		cancellable *T.GCancellable) *T.GSource

	G_socket_speaks_ipv4 func(
		socket *T.GSocket) T.Gboolean

	G_socket_get_credentials func(
		socket *T.GSocket,
		err **T.GError) *T.GCredentials

	G_socket_receive_with_blocking func(
		socket *T.GSocket,
		buffer string,
		size T.Gsize,
		blocking T.Gboolean,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_socket_send_with_blocking func(
		socket *T.GSocket,
		buffer string,
		size T.Gsize,
		blocking T.Gboolean,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	G_socket_connection_get_type func() T.GType

	G_socket_connection_get_socket func(
		connection *T.GSocketConnection) *T.GSocket

	G_socket_connection_get_local_address func(
		connection *T.GSocketConnection,
		err **T.GError) *T.GSocketAddress

	G_socket_connection_get_remote_address func(
		connection *T.GSocketConnection,
		err **T.GError) *T.GSocketAddress

	G_socket_connection_factory_register_type func(

		G_type T.GType,
		family T.GSocketFamily,
		t T.GSocketType,
		protocol T.Gint)

	G_socket_connection_factory_lookup_type func(
		family T.GSocketFamily,
		t T.GSocketType,
		protocol_id T.Gint) T.GType

	G_socket_connection_factory_create_connection func(
		socket *T.GSocket) *T.GSocketConnection

	G_socket_control_message_get_type func() T.GType

	G_socket_control_message_get_size func(
		message *T.GSocketControlMessage) T.Gsize

	G_socket_control_message_get_level func(
		message *T.GSocketControlMessage) int

	G_socket_control_message_get_msg_type func(
		message *T.GSocketControlMessage) int

	G_socket_control_message_serialize func(
		message *T.GSocketControlMessage,
		data T.Gpointer)

	G_socket_control_message_deserialize func(
		level int,
		typ int,
		size T.Gsize,
		data T.Gpointer) *T.GSocketControlMessage

	G_socket_listener_get_type func() T.GType

	G_socket_listener_new func() *T.GSocketListener

	G_socket_listener_set_backlog func(
		listener *T.GSocketListener,
		listen_backlog int)

	G_socket_listener_add_socket func(
		listener *T.GSocketListener,
		socket *T.GSocket,
		source_object *T.GObject,
		err **T.GError) T.Gboolean

	G_socket_listener_add_address func(
		listener *T.GSocketListener,
		address *T.GSocketAddress,
		t T.GSocketType,
		protocol T.GSocketProtocol,
		source_object *T.GObject,
		effective_address **T.GSocketAddress,
		err **T.GError) T.Gboolean

	G_socket_listener_add_inet_port func(
		listener *T.GSocketListener,
		port T.Guint16,
		source_object *T.GObject,
		err **T.GError) T.Gboolean

	G_socket_listener_add_any_inet_port func(
		listener *T.GSocketListener,
		source_object *T.GObject,
		err **T.GError) T.Guint16

	G_socket_listener_accept_socket func(
		listener *T.GSocketListener,
		source_object **T.GObject,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocket

	G_socket_listener_accept_socket_async func(
		listener *T.GSocketListener,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_socket_listener_accept_socket_finish func(
		listener *T.GSocketListener,
		result *T.GAsyncResult,
		source_object **T.GObject,
		err **T.GError) *T.GSocket

	G_socket_listener_accept func(
		listener *T.GSocketListener,
		source_object **T.GObject,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	G_socket_listener_accept_async func(
		listener *T.GSocketListener,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_socket_listener_accept_finish func(
		listener *T.GSocketListener,
		result *T.GAsyncResult,
		source_object **T.GObject,
		err **T.GError) *T.GSocketConnection

	G_socket_listener_close func(
		listener *T.GSocketListener)

	G_socket_service_get_type func() T.GType

	G_socket_service_new func() *T.GSocketService

	G_socket_service_start func(
		service *T.GSocketService)

	G_socket_service_stop func(
		service *T.GSocketService)

	G_socket_service_is_active func(
		service *T.GSocketService) T.Gboolean

	G_srv_target_get_type func() T.GType

	G_srv_target_new func(
		hostname string,
		port T.Guint16,
		priority T.Guint16,
		weight T.Guint16) *T.GSrvTarget

	G_srv_target_copy func(
		target *T.GSrvTarget) *T.GSrvTarget

	G_srv_target_free func(
		target *T.GSrvTarget)

	G_srv_target_get_hostname func(
		target *T.GSrvTarget) string

	G_srv_target_get_port func(
		target *T.GSrvTarget) T.Guint16

	G_srv_target_get_priority func(
		target *T.GSrvTarget) T.Guint16

	G_srv_target_get_weight func(
		target *T.GSrvTarget) T.Guint16

	G_srv_target_list_sort func(
		targets *T.GList) *T.GList

	G_tcp_connection_get_type func() T.GType

	G_tcp_connection_set_graceful_disconnect func(
		connection *T.GTcpConnection,
		graceful_disconnect T.Gboolean)

	G_tcp_connection_get_graceful_disconnect func(
		connection *T.GTcpConnection) T.Gboolean

	G_tcp_wrapper_connection_get_type func() T.GType

	G_tcp_wrapper_connection_new func(
		base_io_stream *T.GIOStream,
		socket *T.GSocket) *T.GSocketConnection

	G_tcp_wrapper_connection_get_base_io_stream func(
		conn *T.GTcpWrapperConnection) *T.GIOStream

	G_themed_icon_get_type func() T.GType

	G_themed_icon_new func(
		iconname string) *T.GIcon

	G_themed_icon_new_with_default_fallbacks func(
		iconname string) *T.GIcon

	G_themed_icon_new_from_names func(
		iconnames **T.Char,
		len int) *T.GIcon

	G_themed_icon_prepend_name func(
		icon *T.GThemedIcon,
		iconname string)

	G_themed_icon_append_name func(
		icon *T.GThemedIcon,
		iconname string)

	G_themed_icon_get_names func(
		icon *T.GThemedIcon) **T.Gchar

	G_threaded_socket_service_get_type func() T.GType

	G_threaded_socket_service_new func(
		max_threads int) *T.GSocketService

	G_tls_backend_get_type func() T.GType

	G_tls_backend_get_default func() *T.GTlsBackend

	G_tls_backend_supports_tls func(
		backend *T.GTlsBackend) T.Gboolean

	G_tls_backend_get_certificate_type func(
		backend *T.GTlsBackend) T.GType

	G_tls_backend_get_client_connection_type func(
		backend *T.GTlsBackend) T.GType

	G_tls_backend_get_server_connection_type func(
		backend *T.GTlsBackend) T.GType

	G_tls_certificate_get_type func() T.GType

	G_tls_certificate_new_from_pem func(
		data string,
		length T.Gssize,
		err **T.GError) *T.GTlsCertificate

	G_tls_certificate_new_from_file func(
		file string,
		err **T.GError) *T.GTlsCertificate

	G_tls_certificate_new_from_files func(
		cert_file string,
		key_file string,
		err **T.GError) *T.GTlsCertificate

	G_tls_certificate_list_new_from_file func(
		file string,
		err **T.GError) *T.GList

	G_tls_certificate_get_issuer func(
		cert *T.GTlsCertificate) *T.GTlsCertificate

	G_tls_certificate_verify func(
		cert *T.GTlsCertificate,
		identity *T.GSocketConnectable,
		trusted_ca *T.GTlsCertificate) T.GTlsCertificateFlags

	G_tls_connection_get_type func() T.GType

	G_tls_connection_set_use_system_certdb func(
		conn *T.GTlsConnection,
		use_system_certdb T.Gboolean)

	G_tls_connection_get_use_system_certdb func(
		conn *T.GTlsConnection) T.Gboolean

	G_tls_connection_set_certificate func(
		conn *T.GTlsConnection,
		certificate *T.GTlsCertificate)

	G_tls_connection_get_certificate func(
		conn *T.GTlsConnection) *T.GTlsCertificate

	G_tls_connection_get_peer_certificate func(
		conn *T.GTlsConnection) *T.GTlsCertificate

	G_tls_connection_get_peer_certificate_errors func(
		conn *T.GTlsConnection) T.GTlsCertificateFlags

	G_tls_connection_set_require_close_notify func(
		conn *T.GTlsConnection,
		require_close_notify T.Gboolean)

	G_tls_connection_get_require_close_notify func(
		conn *T.GTlsConnection) T.Gboolean

	G_tls_connection_set_rehandshake_mode func(
		conn *T.GTlsConnection,
		mode T.GTlsRehandshakeMode)

	G_tls_connection_get_rehandshake_mode func(
		conn *T.GTlsConnection) T.GTlsRehandshakeMode

	G_tls_connection_handshake func(
		conn *T.GTlsConnection,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	G_tls_connection_handshake_async func(
		conn *T.GTlsConnection,
		io_priority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_tls_connection_handshake_finish func(
		conn *T.GTlsConnection,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_tls_error_quark func() T.GQuark

	G_tls_connection_emit_accept_certificate func(
		conn *T.GTlsConnection,
		peer_cert *T.GTlsCertificate,
		errors T.GTlsCertificateFlags) T.Gboolean

	G_tls_client_connection_get_type func() T.GType

	G_tls_client_connection_new func(
		base_io_stream *T.GIOStream,
		server_identity *T.GSocketConnectable,
		err **T.GError) *T.GIOStream

	G_tls_client_connection_get_validation_flags func(
		conn *T.GTlsClientConnection) T.GTlsCertificateFlags

	G_tls_client_connection_set_validation_flags func(
		conn *T.GTlsClientConnection,
		flags T.GTlsCertificateFlags)

	G_tls_client_connection_get_server_identity func(
		conn *T.GTlsClientConnection) *T.GSocketConnectable

	G_tls_client_connection_set_server_identity func(
		conn *T.GTlsClientConnection,
		identity *T.GSocketConnectable)

	G_tls_client_connection_get_use_ssl3 func(
		conn *T.GTlsClientConnection) T.Gboolean

	G_tls_client_connection_set_use_ssl3 func(
		conn *T.GTlsClientConnection,
		use_ssl3 T.Gboolean)

	G_tls_client_connection_get_accepted_cas func(
		conn *T.GTlsClientConnection) *T.GList

	G_tls_server_connection_get_type func() T.GType

	G_tls_server_connection_new func(
		base_io_stream *T.GIOStream,
		certificate *T.GTlsCertificate,
		err **T.GError) *T.GIOStream

	G_vfs_get_type func() T.GType

	G_vfs_is_active func(
		vfs *T.GVfs) T.Gboolean

	G_vfs_get_file_for_path func(
		vfs *T.GVfs,
		path string) *T.GFile

	G_vfs_get_file_for_uri func(
		vfs *T.GVfs,
		uri string) *T.GFile

	G_vfs_get_supported_uri_schemes func(
		vfs *T.GVfs) **T.Gchar

	G_vfs_parse_name func(
		vfs *T.GVfs,
		parse_name string) *T.GFile

	G_vfs_get_default func() *T.GVfs

	G_vfs_get_local func() *T.GVfs

	G_volume_get_type func() T.GType

	G_volume_get_name func(
		volume *T.GVolume) string

	G_volume_get_icon func(
		volume *T.GVolume) *T.GIcon

	G_volume_get_uuid func(
		volume *T.GVolume) string

	G_volume_get_drive func(
		volume *T.GVolume) *T.GDrive

	G_volume_get_mount func(
		volume *T.GVolume) *T.GMount

	G_volume_can_mount func(
		volume *T.GVolume) T.Gboolean

	G_volume_can_eject func(
		volume *T.GVolume) T.Gboolean

	G_volume_should_automount func(
		volume *T.GVolume) T.Gboolean

	G_volume_mount func(
		volume *T.GVolume,
		flags T.GMountMountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_volume_mount_finish func(
		volume *T.GVolume,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_volume_eject func(
		volume *T.GVolume,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_volume_eject_finish func(
		volume *T.GVolume,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_volume_get_identifier func(
		volume *T.GVolume,
		kind string) string

	G_volume_enumerate_identifiers func(
		volume *T.GVolume) **T.Char

	G_volume_get_activation_root func(
		volume *T.GVolume) *T.GFile

	G_volume_eject_with_operation func(
		volume *T.GVolume,
		flags T.GMountUnmountFlags,
		mount_operation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	G_volume_eject_with_operation_finish func(
		volume *T.GVolume,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	G_zlib_compressor_get_type func() T.GType

	G_zlib_compressor_new func(
		format T.GZlibCompressorFormat,
		level int) *T.GZlibCompressor

	G_zlib_compressor_get_file_info func(
		compressor *T.GZlibCompressor) *T.GFileInfo

	G_zlib_compressor_set_file_info func(
		compressor *T.GZlibCompressor,
		file_info *T.GFileInfo)

	G_zlib_decompressor_get_type func() T.GType

	G_zlib_decompressor_new func(
		format T.GZlibCompressorFormat) *T.GZlibDecompressor

	G_zlib_decompressor_get_file_info func(
		decompressor *T.GZlibDecompressor) *T.GFileInfo

	G_win32_input_stream_get_type func() T.GType

	G_win32_input_stream_new func(
		handle *T.Void,
		close_handle T.Gboolean) *T.GInputStream

	G_win32_input_stream_set_close_handle func(
		stream *T.GWin32InputStream,
		close_handle T.Gboolean)

	G_win32_input_stream_get_close_handle func(
		stream *T.GWin32InputStream) T.Gboolean

	G_win32_input_stream_get_handle func(
		stream *T.GWin32InputStream) *T.Void

	G_win32_output_stream_get_type func() T.GType

	G_win32_output_stream_new func(
		handle *T.Void,
		close_handle T.Gboolean) *T.GOutputStream

	G_win32_output_stream_set_close_handle func(
		stream *T.GWin32OutputStream,
		close_handle T.Gboolean)

	G_win32_output_stream_get_close_handle func(
		stream *T.GWin32OutputStream) T.Gboolean

	G_win32_output_stream_get_handle func(
		stream *T.GWin32OutputStream) *T.Void

	G_settings_backend_get_type func() T.GType

	G_settings_backend_changed func(
		backend *T.GSettingsBackend,
		key string,
		origin_tag T.Gpointer)

	G_settings_backend_path_changed func(
		backend *T.GSettingsBackend,
		path string,
		origin_tag T.Gpointer)

	G_settings_backend_flatten_tree func(
		tree *T.GTree,
		path **T.Gchar,
		keys ***T.Gchar,
		values ***T.GVariant)

	G_settings_backend_keys_changed func(
		backend *T.GSettingsBackend,
		path string,
		items **T.Gchar,
		origin_tag T.Gpointer)

	G_settings_backend_path_writable_changed func(
		backend *T.GSettingsBackend,
		path string)

	G_settings_backend_writable_changed func(
		backend *T.GSettingsBackend,
		key string)

	G_settings_backend_changed_tree func(
		backend *T.GSettingsBackend,
		tree *T.GTree,
		origin_tag T.Gpointer)

	G_settings_backend_get_default func() *T.GSettingsBackend

	G_keyfile_settings_backend_new func(
		filename string,
		root_path string,
		root_group string) *T.GSettingsBackend

	G_null_settings_backend_new func() *T.GSettingsBackend

	G_memory_settings_backend_new func() *T.GSettingsBackend
)
var dll = "libgio-2.0-0.dll"

var apiList = Apis{
	{"g_action_activate", &G_action_activate},
	{"g_action_get_enabled", &G_action_get_enabled},
	{"g_action_get_name", &G_action_get_name},
	{"g_action_get_parameter_type", &G_action_get_parameter_type},
	{"g_action_get_state", &G_action_get_state},
	{"g_action_get_state_hint", &G_action_get_state_hint},
	{"g_action_get_state_type", &G_action_get_state_type},
	{"g_action_get_type", &G_action_get_type},
	{"g_action_group_action_added", &G_action_group_action_added},
	{"g_action_group_action_enabled_changed", &G_action_group_action_enabled_changed},
	{"g_action_group_action_removed", &G_action_group_action_removed},
	{"g_action_group_action_state_changed", &G_action_group_action_state_changed},
	{"g_action_group_activate_action", &G_action_group_activate_action},
	{"g_action_group_change_action_state", &G_action_group_change_action_state},
	{"g_action_group_get_action_enabled", &G_action_group_get_action_enabled},
	{"g_action_group_get_action_parameter_type", &G_action_group_get_action_parameter_type},
	{"g_action_group_get_action_state", &G_action_group_get_action_state},
	{"g_action_group_get_action_state_hint", &G_action_group_get_action_state_hint},
	{"g_action_group_get_action_state_type", &G_action_group_get_action_state_type},
	{"g_action_group_get_type", &G_action_group_get_type},
	{"g_action_group_has_action", &G_action_group_has_action},
	{"g_action_group_list_actions", &G_action_group_list_actions},
	{"g_action_set_state", &G_action_set_state},
	{"g_app_info_add_supports_type", &G_app_info_add_supports_type},
	{"g_app_info_can_delete", &G_app_info_can_delete},
	{"g_app_info_can_remove_supports_type", &G_app_info_can_remove_supports_type},
	{"g_app_info_create_flags_get_type", &G_app_info_create_flags_get_type},
	{"g_app_info_create_from_commandline", &G_app_info_create_from_commandline},
	{"g_app_info_delete", &G_app_info_delete},
	{"g_app_info_dup", &G_app_info_dup},
	{"g_app_info_equal", &G_app_info_equal},
	{"g_app_info_get_all", &G_app_info_get_all},
	{"g_app_info_get_all_for_type", &G_app_info_get_all_for_type},
	{"g_app_info_get_commandline", &G_app_info_get_commandline},
	{"g_app_info_get_default_for_type", &G_app_info_get_default_for_type},
	{"g_app_info_get_default_for_uri_scheme", &G_app_info_get_default_for_uri_scheme},
	{"g_app_info_get_description", &G_app_info_get_description},
	{"g_app_info_get_display_name", &G_app_info_get_display_name},
	{"g_app_info_get_executable", &G_app_info_get_executable},
	{"g_app_info_get_fallback_for_type", &G_app_info_get_fallback_for_type},
	{"g_app_info_get_icon", &G_app_info_get_icon},
	{"g_app_info_get_id", &G_app_info_get_id},
	{"g_app_info_get_name", &G_app_info_get_name},
	{"g_app_info_get_recommended_for_type", &G_app_info_get_recommended_for_type},
	{"g_app_info_get_type", &G_app_info_get_type},
	{"g_app_info_launch", &G_app_info_launch},
	{"g_app_info_launch_default_for_uri", &G_app_info_launch_default_for_uri},
	{"g_app_info_launch_uris", &G_app_info_launch_uris},
	{"g_app_info_remove_supports_type", &G_app_info_remove_supports_type},
	{"g_app_info_reset_type_associations", &G_app_info_reset_type_associations},
	{"g_app_info_set_as_default_for_extension", &G_app_info_set_as_default_for_extension},
	{"g_app_info_set_as_default_for_type", &G_app_info_set_as_default_for_type},
	{"g_app_info_set_as_last_used_for_type", &G_app_info_set_as_last_used_for_type},
	{"g_app_info_should_show", &G_app_info_should_show},
	{"g_app_info_supports_files", &G_app_info_supports_files},
	{"g_app_info_supports_uris", &G_app_info_supports_uris},
	{"g_app_launch_context_get_display", &G_app_launch_context_get_display},
	{"g_app_launch_context_get_startup_notify_id", &G_app_launch_context_get_startup_notify_id},
	{"g_app_launch_context_get_type", &G_app_launch_context_get_type},
	{"g_app_launch_context_launch_failed", &G_app_launch_context_launch_failed},
	{"g_app_launch_context_new", &G_app_launch_context_new},
	{"g_application_activate", &G_application_activate},
	{"g_application_command_line_get_arguments", &G_application_command_line_get_arguments},
	{"g_application_command_line_get_cwd", &G_application_command_line_get_cwd},
	{"g_application_command_line_get_environ", &G_application_command_line_get_environ},
	{"g_application_command_line_get_exit_status", &G_application_command_line_get_exit_status},
	{"g_application_command_line_get_is_remote", &G_application_command_line_get_is_remote},
	{"g_application_command_line_get_platform_data", &G_application_command_line_get_platform_data},
	{"g_application_command_line_get_type", &G_application_command_line_get_type},
	{"g_application_command_line_getenv", &G_application_command_line_getenv},
	{"g_application_command_line_print", &G_application_command_line_print},
	{"g_application_command_line_printerr", &G_application_command_line_printerr},
	{"g_application_command_line_set_exit_status", &G_application_command_line_set_exit_status},
	{"g_application_flags_get_type", &G_application_flags_get_type},
	{"g_application_get_application_id", &G_application_get_application_id},
	{"g_application_get_flags", &G_application_get_flags},
	{"g_application_get_inactivity_timeout", &G_application_get_inactivity_timeout},
	{"g_application_get_is_registered", &G_application_get_is_registered},
	{"g_application_get_is_remote", &G_application_get_is_remote},
	{"g_application_get_type", &G_application_get_type},
	{"g_application_hold", &G_application_hold},
	{"g_application_id_is_valid", &G_application_id_is_valid},
	{"g_application_new", &G_application_new},
	{"g_application_open", &G_application_open},
	{"g_application_register", &G_application_register},
	{"g_application_release", &G_application_release},
	{"g_application_run", &G_application_run},
	{"g_application_set_action_group", &G_application_set_action_group},
	{"g_application_set_application_id", &G_application_set_application_id},
	{"g_application_set_flags", &G_application_set_flags},
	{"g_application_set_inactivity_timeout", &G_application_set_inactivity_timeout},
	{"g_ask_password_flags_get_type", &G_ask_password_flags_get_type},
	{"g_async_initable_get_type", &G_async_initable_get_type},
	{"g_async_initable_init_async", &G_async_initable_init_async},
	{"g_async_initable_init_finish", &G_async_initable_init_finish},
	{"g_async_initable_new_async", &G_async_initable_new_async},
	{"g_async_initable_new_finish", &G_async_initable_new_finish},
	{"g_async_initable_new_valist_async", &G_async_initable_new_valist_async},
	{"g_async_initable_newv_async", &G_async_initable_newv_async},
	{"g_async_result_get_source_object", &G_async_result_get_source_object},
	{"g_async_result_get_type", &G_async_result_get_type},
	{"g_async_result_get_user_data", &G_async_result_get_user_data},
	{"g_buffered_input_stream_fill", &G_buffered_input_stream_fill},
	{"g_buffered_input_stream_fill_async", &G_buffered_input_stream_fill_async},
	{"g_buffered_input_stream_fill_finish", &G_buffered_input_stream_fill_finish},
	{"g_buffered_input_stream_get_available", &G_buffered_input_stream_get_available},
	{"g_buffered_input_stream_get_buffer_size", &G_buffered_input_stream_get_buffer_size},
	{"g_buffered_input_stream_get_type", &G_buffered_input_stream_get_type},
	{"g_buffered_input_stream_new", &G_buffered_input_stream_new},
	{"g_buffered_input_stream_new_sized", &G_buffered_input_stream_new_sized},
	{"g_buffered_input_stream_peek", &G_buffered_input_stream_peek},
	{"g_buffered_input_stream_peek_buffer", &G_buffered_input_stream_peek_buffer},
	{"g_buffered_input_stream_read_byte", &G_buffered_input_stream_read_byte},
	{"g_buffered_input_stream_set_buffer_size", &G_buffered_input_stream_set_buffer_size},
	{"g_buffered_output_stream_get_auto_grow", &G_buffered_output_stream_get_auto_grow},
	{"g_buffered_output_stream_get_buffer_size", &G_buffered_output_stream_get_buffer_size},
	{"g_buffered_output_stream_get_type", &G_buffered_output_stream_get_type},
	{"g_buffered_output_stream_new", &G_buffered_output_stream_new},
	{"g_buffered_output_stream_new_sized", &G_buffered_output_stream_new_sized},
	{"g_buffered_output_stream_set_auto_grow", &G_buffered_output_stream_set_auto_grow},
	{"g_buffered_output_stream_set_buffer_size", &G_buffered_output_stream_set_buffer_size},
	{"g_bus_get", &G_bus_get},
	{"g_bus_get_finish", &G_bus_get_finish},
	{"g_bus_get_sync", &G_bus_get_sync},
	{"g_bus_name_owner_flags_get_type", &G_bus_name_owner_flags_get_type},
	{"g_bus_name_watcher_flags_get_type", &G_bus_name_watcher_flags_get_type},
	{"g_bus_own_name", &G_bus_own_name},
	{"g_bus_own_name_on_connection", &G_bus_own_name_on_connection},
	{"g_bus_own_name_on_connection_with_closures", &G_bus_own_name_on_connection_with_closures},
	{"g_bus_own_name_with_closures", &G_bus_own_name_with_closures},
	{"g_bus_type_get_type", &G_bus_type_get_type},
	{"g_bus_unown_name", &G_bus_unown_name},
	{"g_bus_unwatch_name", &G_bus_unwatch_name},
	{"g_bus_watch_name", &G_bus_watch_name},
	{"g_bus_watch_name_on_connection", &G_bus_watch_name_on_connection},
	{"g_bus_watch_name_on_connection_with_closures", &G_bus_watch_name_on_connection_with_closures},
	{"g_bus_watch_name_with_closures", &G_bus_watch_name_with_closures},
	{"g_cancellable_cancel", &G_cancellable_cancel},
	{"g_cancellable_connect", &G_cancellable_connect},
	{"g_cancellable_disconnect", &G_cancellable_disconnect},
	{"g_cancellable_get_current", &G_cancellable_get_current},
	{"g_cancellable_get_fd", &G_cancellable_get_fd},
	{"g_cancellable_get_type", &G_cancellable_get_type},
	{"g_cancellable_is_cancelled", &G_cancellable_is_cancelled},
	{"g_cancellable_make_pollfd", &G_cancellable_make_pollfd},
	{"g_cancellable_new", &G_cancellable_new},
	{"g_cancellable_pop_current", &G_cancellable_pop_current},
	{"g_cancellable_push_current", &G_cancellable_push_current},
	{"g_cancellable_release_fd", &G_cancellable_release_fd},
	{"g_cancellable_reset", &G_cancellable_reset},
	{"g_cancellable_set_error_if_cancelled", &G_cancellable_set_error_if_cancelled},
	{"g_cancellable_source_new", &G_cancellable_source_new},
	{"g_charset_converter_get_num_fallbacks", &G_charset_converter_get_num_fallbacks},
	{"g_charset_converter_get_type", &G_charset_converter_get_type},
	{"g_charset_converter_get_use_fallback", &G_charset_converter_get_use_fallback},
	{"g_charset_converter_new", &G_charset_converter_new},
	{"g_charset_converter_set_use_fallback", &G_charset_converter_set_use_fallback},
	{"g_content_type_can_be_executable", &G_content_type_can_be_executable},
	{"g_content_type_equals", &G_content_type_equals},
	{"g_content_type_from_mime_type", &G_content_type_from_mime_type},
	{"g_content_type_get_description", &G_content_type_get_description},
	{"g_content_type_get_icon", &G_content_type_get_icon},
	{"g_content_type_get_mime_type", &G_content_type_get_mime_type},
	{"g_content_type_guess", &G_content_type_guess},
	{"g_content_type_guess_for_tree", &G_content_type_guess_for_tree},
	{"g_content_type_is_a", &G_content_type_is_a},
	{"g_content_type_is_unknown", &G_content_type_is_unknown},
	{"g_content_types_get_registered", &G_content_types_get_registered},
	{"g_converter_convert", &G_converter_convert},
	{"g_converter_flags_get_type", &G_converter_flags_get_type},
	{"g_converter_get_type", &G_converter_get_type},
	{"g_converter_input_stream_get_converter", &G_converter_input_stream_get_converter},
	{"g_converter_input_stream_get_type", &G_converter_input_stream_get_type},
	{"g_converter_input_stream_new", &G_converter_input_stream_new},
	{"g_converter_output_stream_get_converter", &G_converter_output_stream_get_converter},
	{"g_converter_output_stream_get_type", &G_converter_output_stream_get_type},
	{"g_converter_output_stream_new", &G_converter_output_stream_new},
	{"g_converter_reset", &G_converter_reset},
	{"g_converter_result_get_type", &G_converter_result_get_type},
	{"g_credentials_get_native", &G_credentials_get_native},
	{"g_credentials_get_type", &G_credentials_get_type},
	{"g_credentials_is_same_user", &G_credentials_is_same_user},
	{"g_credentials_new", &G_credentials_new},
	{"g_credentials_set_native", &G_credentials_set_native},
	{"g_credentials_to_string", &G_credentials_to_string},
	{"g_credentials_type_get_type", &G_credentials_type_get_type},
	{"g_data_input_stream_get_byte_order", &G_data_input_stream_get_byte_order},
	{"g_data_input_stream_get_newline_type", &G_data_input_stream_get_newline_type},
	{"g_data_input_stream_get_type", &G_data_input_stream_get_type},
	{"g_data_input_stream_new", &G_data_input_stream_new},
	{"g_data_input_stream_read_byte", &G_data_input_stream_read_byte},
	{"g_data_input_stream_read_int16", &G_data_input_stream_read_int16},
	{"g_data_input_stream_read_int32", &G_data_input_stream_read_int32},
	{"g_data_input_stream_read_int64", &G_data_input_stream_read_int64},
	{"g_data_input_stream_read_line", &G_data_input_stream_read_line},
	{"g_data_input_stream_read_line_async", &G_data_input_stream_read_line_async},
	{"g_data_input_stream_read_line_finish", &G_data_input_stream_read_line_finish},
	{"g_data_input_stream_read_uint16", &G_data_input_stream_read_uint16},
	{"g_data_input_stream_read_uint32", &G_data_input_stream_read_uint32},
	{"g_data_input_stream_read_uint64", &G_data_input_stream_read_uint64},
	{"g_data_input_stream_read_until", &G_data_input_stream_read_until},
	{"g_data_input_stream_read_until_async", &G_data_input_stream_read_until_async},
	{"g_data_input_stream_read_until_finish", &G_data_input_stream_read_until_finish},
	{"g_data_input_stream_read_upto", &G_data_input_stream_read_upto},
	{"g_data_input_stream_read_upto_async", &G_data_input_stream_read_upto_async},
	{"g_data_input_stream_read_upto_finish", &G_data_input_stream_read_upto_finish},
	{"g_data_input_stream_set_byte_order", &G_data_input_stream_set_byte_order},
	{"g_data_input_stream_set_newline_type", &G_data_input_stream_set_newline_type},
	{"g_data_output_stream_get_byte_order", &G_data_output_stream_get_byte_order},
	{"g_data_output_stream_get_type", &G_data_output_stream_get_type},
	{"g_data_output_stream_new", &G_data_output_stream_new},
	{"g_data_output_stream_put_byte", &G_data_output_stream_put_byte},
	{"g_data_output_stream_put_int16", &G_data_output_stream_put_int16},
	{"g_data_output_stream_put_int32", &G_data_output_stream_put_int32},
	{"g_data_output_stream_put_int64", &G_data_output_stream_put_int64},
	{"g_data_output_stream_put_string", &G_data_output_stream_put_string},
	{"g_data_output_stream_put_uint16", &G_data_output_stream_put_uint16},
	{"g_data_output_stream_put_uint32", &G_data_output_stream_put_uint32},
	{"g_data_output_stream_put_uint64", &G_data_output_stream_put_uint64},
	{"g_data_output_stream_set_byte_order", &G_data_output_stream_set_byte_order},
	{"g_data_stream_byte_order_get_type", &G_data_stream_byte_order_get_type},
	{"g_data_stream_newline_type_get_type", &G_data_stream_newline_type_get_type},
	{"g_dbus_address_get_for_bus_sync", &G_dbus_address_get_for_bus_sync},
	{"g_dbus_address_get_stream", &G_dbus_address_get_stream},
	{"g_dbus_address_get_stream_finish", &G_dbus_address_get_stream_finish},
	{"g_dbus_address_get_stream_sync", &G_dbus_address_get_stream_sync},
	{"g_dbus_annotation_info_get_type", &G_dbus_annotation_info_get_type},
	{"g_dbus_annotation_info_lookup", &G_dbus_annotation_info_lookup},
	{"g_dbus_annotation_info_ref", &G_dbus_annotation_info_ref},
	{"g_dbus_annotation_info_unref", &G_dbus_annotation_info_unref},
	{"g_dbus_arg_info_get_type", &G_dbus_arg_info_get_type},
	{"g_dbus_arg_info_ref", &G_dbus_arg_info_ref},
	{"g_dbus_arg_info_unref", &G_dbus_arg_info_unref},
	{"g_dbus_auth_observer_authorize_authenticated_peer", &G_dbus_auth_observer_authorize_authenticated_peer},
	{"g_dbus_auth_observer_get_type", &G_dbus_auth_observer_get_type},
	{"g_dbus_auth_observer_new", &G_dbus_auth_observer_new},
	{"g_dbus_call_flags_get_type", &G_dbus_call_flags_get_type},
	{"g_dbus_capability_flags_get_type", &G_dbus_capability_flags_get_type},
	{"g_dbus_connection_add_filter", &G_dbus_connection_add_filter},
	{"g_dbus_connection_call", &G_dbus_connection_call},
	{"g_dbus_connection_call_finish", &G_dbus_connection_call_finish},
	{"g_dbus_connection_call_sync", &G_dbus_connection_call_sync},
	{"g_dbus_connection_close", &G_dbus_connection_close},
	{"g_dbus_connection_close_finish", &G_dbus_connection_close_finish},
	{"g_dbus_connection_close_sync", &G_dbus_connection_close_sync},
	{"g_dbus_connection_emit_signal", &G_dbus_connection_emit_signal},
	{"g_dbus_connection_flags_get_type", &G_dbus_connection_flags_get_type},
	{"g_dbus_connection_flush", &G_dbus_connection_flush},
	{"g_dbus_connection_flush_finish", &G_dbus_connection_flush_finish},
	{"g_dbus_connection_flush_sync", &G_dbus_connection_flush_sync},
	{"g_dbus_connection_get_capabilities", &G_dbus_connection_get_capabilities},
	{"g_dbus_connection_get_exit_on_close", &G_dbus_connection_get_exit_on_close},
	{"g_dbus_connection_get_guid", &G_dbus_connection_get_guid},
	{"g_dbus_connection_get_peer_credentials", &G_dbus_connection_get_peer_credentials},
	{"g_dbus_connection_get_stream", &G_dbus_connection_get_stream},
	{"g_dbus_connection_get_type", &G_dbus_connection_get_type},
	{"g_dbus_connection_get_unique_name", &G_dbus_connection_get_unique_name},
	{"g_dbus_connection_is_closed", &G_dbus_connection_is_closed},
	{"g_dbus_connection_new", &G_dbus_connection_new},
	{"g_dbus_connection_new_finish", &G_dbus_connection_new_finish},
	{"g_dbus_connection_new_for_address", &G_dbus_connection_new_for_address},
	{"g_dbus_connection_new_for_address_finish", &G_dbus_connection_new_for_address_finish},
	{"g_dbus_connection_new_for_address_sync", &G_dbus_connection_new_for_address_sync},
	{"g_dbus_connection_new_sync", &G_dbus_connection_new_sync},
	{"g_dbus_connection_register_object", &G_dbus_connection_register_object},
	{"g_dbus_connection_register_subtree", &G_dbus_connection_register_subtree},
	{"g_dbus_connection_remove_filter", &G_dbus_connection_remove_filter},
	{"g_dbus_connection_send_message", &G_dbus_connection_send_message},
	{"g_dbus_connection_send_message_with_reply", &G_dbus_connection_send_message_with_reply},
	{"g_dbus_connection_send_message_with_reply_finish", &G_dbus_connection_send_message_with_reply_finish},
	{"g_dbus_connection_send_message_with_reply_sync", &G_dbus_connection_send_message_with_reply_sync},
	{"g_dbus_connection_set_exit_on_close", &G_dbus_connection_set_exit_on_close},
	{"g_dbus_connection_signal_subscribe", &G_dbus_connection_signal_subscribe},
	{"g_dbus_connection_signal_unsubscribe", &G_dbus_connection_signal_unsubscribe},
	{"g_dbus_connection_start_message_processing", &G_dbus_connection_start_message_processing},
	{"g_dbus_connection_unregister_object", &G_dbus_connection_unregister_object},
	{"g_dbus_connection_unregister_subtree", &G_dbus_connection_unregister_subtree},
	{"g_dbus_error_encode_gerror", &G_dbus_error_encode_gerror},
	{"g_dbus_error_get_remote_error", &G_dbus_error_get_remote_error},
	{"g_dbus_error_get_type", &G_dbus_error_get_type},
	{"g_dbus_error_is_remote_error", &G_dbus_error_is_remote_error},
	{"g_dbus_error_new_for_dbus_error", &G_dbus_error_new_for_dbus_error},
	{"g_dbus_error_quark", &G_dbus_error_quark},
	{"g_dbus_error_register_error", &G_dbus_error_register_error},
	{"g_dbus_error_register_error_domain", &G_dbus_error_register_error_domain},
	{"g_dbus_error_set_dbus_error", &G_dbus_error_set_dbus_error},
	{"g_dbus_error_set_dbus_error_valist", &G_dbus_error_set_dbus_error_valist},
	{"g_dbus_error_strip_remote_error", &G_dbus_error_strip_remote_error},
	{"g_dbus_error_unregister_error", &G_dbus_error_unregister_error},
	{"g_dbus_generate_guid", &G_dbus_generate_guid},
	{"g_dbus_interface_info_generate_xml", &G_dbus_interface_info_generate_xml},
	{"g_dbus_interface_info_get_type", &G_dbus_interface_info_get_type},
	{"g_dbus_interface_info_lookup_method", &G_dbus_interface_info_lookup_method},
	{"g_dbus_interface_info_lookup_property", &G_dbus_interface_info_lookup_property},
	{"g_dbus_interface_info_lookup_signal", &G_dbus_interface_info_lookup_signal},
	{"g_dbus_interface_info_ref", &G_dbus_interface_info_ref},
	{"g_dbus_interface_info_unref", &G_dbus_interface_info_unref},
	{"g_dbus_is_address", &G_dbus_is_address},
	{"g_dbus_is_guid", &G_dbus_is_guid},
	{"g_dbus_is_interface_name", &G_dbus_is_interface_name},
	{"g_dbus_is_member_name", &G_dbus_is_member_name},
	{"g_dbus_is_name", &G_dbus_is_name},
	{"g_dbus_is_supported_address", &G_dbus_is_supported_address},
	{"g_dbus_is_unique_name", &G_dbus_is_unique_name},
	{"g_dbus_message_byte_order_get_type", &G_dbus_message_byte_order_get_type},
	{"g_dbus_message_bytes_needed", &G_dbus_message_bytes_needed},
	{"g_dbus_message_copy", &G_dbus_message_copy},
	{"g_dbus_message_flags_get_type", &G_dbus_message_flags_get_type},
	{"g_dbus_message_get_arg0", &G_dbus_message_get_arg0},
	{"g_dbus_message_get_body", &G_dbus_message_get_body},
	{"g_dbus_message_get_byte_order", &G_dbus_message_get_byte_order},
	{"g_dbus_message_get_destination", &G_dbus_message_get_destination},
	{"g_dbus_message_get_error_name", &G_dbus_message_get_error_name},
	{"g_dbus_message_get_flags", &G_dbus_message_get_flags},
	{"g_dbus_message_get_header", &G_dbus_message_get_header},
	{"g_dbus_message_get_header_fields", &G_dbus_message_get_header_fields},
	{"g_dbus_message_get_interface", &G_dbus_message_get_interface},
	{"g_dbus_message_get_locked", &G_dbus_message_get_locked},
	{"g_dbus_message_get_member", &G_dbus_message_get_member},
	{"g_dbus_message_get_message_type", &G_dbus_message_get_message_type},
	{"g_dbus_message_get_num_unix_fds", &G_dbus_message_get_num_unix_fds},
	{"g_dbus_message_get_path", &G_dbus_message_get_path},
	{"g_dbus_message_get_reply_serial", &G_dbus_message_get_reply_serial},
	{"g_dbus_message_get_sender", &G_dbus_message_get_sender},
	{"g_dbus_message_get_serial", &G_dbus_message_get_serial},
	{"g_dbus_message_get_signature", &G_dbus_message_get_signature},
	{"g_dbus_message_get_type", &G_dbus_message_get_type},
	{"g_dbus_message_header_field_get_type", &G_dbus_message_header_field_get_type},
	{"g_dbus_message_lock", &G_dbus_message_lock},
	{"g_dbus_message_new", &G_dbus_message_new},
	{"g_dbus_message_new_from_blob", &G_dbus_message_new_from_blob},
	{"g_dbus_message_new_method_call", &G_dbus_message_new_method_call},
	{"g_dbus_message_new_method_error", &G_dbus_message_new_method_error},
	{"g_dbus_message_new_method_error_literal", &G_dbus_message_new_method_error_literal},
	{"g_dbus_message_new_method_error_valist", &G_dbus_message_new_method_error_valist},
	{"g_dbus_message_new_method_reply", &G_dbus_message_new_method_reply},
	{"g_dbus_message_new_signal", &G_dbus_message_new_signal},
	{"g_dbus_message_print", &G_dbus_message_print},
	{"g_dbus_message_set_body", &G_dbus_message_set_body},
	{"g_dbus_message_set_byte_order", &G_dbus_message_set_byte_order},
	{"g_dbus_message_set_destination", &G_dbus_message_set_destination},
	{"g_dbus_message_set_error_name", &G_dbus_message_set_error_name},
	{"g_dbus_message_set_flags", &G_dbus_message_set_flags},
	{"g_dbus_message_set_header", &G_dbus_message_set_header},
	{"g_dbus_message_set_interface", &G_dbus_message_set_interface},
	{"g_dbus_message_set_member", &G_dbus_message_set_member},
	{"g_dbus_message_set_message_type", &G_dbus_message_set_message_type},
	{"g_dbus_message_set_num_unix_fds", &G_dbus_message_set_num_unix_fds},
	{"g_dbus_message_set_path", &G_dbus_message_set_path},
	{"g_dbus_message_set_reply_serial", &G_dbus_message_set_reply_serial},
	{"g_dbus_message_set_sender", &G_dbus_message_set_sender},
	{"g_dbus_message_set_serial", &G_dbus_message_set_serial},
	{"g_dbus_message_set_signature", &G_dbus_message_set_signature},
	{"g_dbus_message_to_blob", &G_dbus_message_to_blob},
	{"g_dbus_message_to_gerror", &G_dbus_message_to_gerror},
	{"g_dbus_message_type_get_type", &G_dbus_message_type_get_type},
	{"g_dbus_method_info_get_type", &G_dbus_method_info_get_type},
	{"g_dbus_method_info_ref", &G_dbus_method_info_ref},
	{"g_dbus_method_info_unref", &G_dbus_method_info_unref},
	{"g_dbus_method_invocation_get_connection", &G_dbus_method_invocation_get_connection},
	{"g_dbus_method_invocation_get_interface_name", &G_dbus_method_invocation_get_interface_name},
	{"g_dbus_method_invocation_get_message", &G_dbus_method_invocation_get_message},
	{"g_dbus_method_invocation_get_method_info", &G_dbus_method_invocation_get_method_info},
	{"g_dbus_method_invocation_get_method_name", &G_dbus_method_invocation_get_method_name},
	{"g_dbus_method_invocation_get_object_path", &G_dbus_method_invocation_get_object_path},
	{"g_dbus_method_invocation_get_parameters", &G_dbus_method_invocation_get_parameters},
	{"g_dbus_method_invocation_get_sender", &G_dbus_method_invocation_get_sender},
	{"g_dbus_method_invocation_get_type", &G_dbus_method_invocation_get_type},
	{"g_dbus_method_invocation_get_user_data", &G_dbus_method_invocation_get_user_data},
	{"g_dbus_method_invocation_return_dbus_error", &G_dbus_method_invocation_return_dbus_error},
	{"g_dbus_method_invocation_return_error", &G_dbus_method_invocation_return_error},
	{"g_dbus_method_invocation_return_error_literal", &G_dbus_method_invocation_return_error_literal},
	{"g_dbus_method_invocation_return_error_valist", &G_dbus_method_invocation_return_error_valist},
	{"g_dbus_method_invocation_return_gerror", &G_dbus_method_invocation_return_gerror},
	{"g_dbus_method_invocation_return_value", &G_dbus_method_invocation_return_value},
	{"g_dbus_node_info_generate_xml", &G_dbus_node_info_generate_xml},
	{"g_dbus_node_info_get_type", &G_dbus_node_info_get_type},
	{"g_dbus_node_info_lookup_interface", &G_dbus_node_info_lookup_interface},
	{"g_dbus_node_info_new_for_xml", &G_dbus_node_info_new_for_xml},
	{"g_dbus_node_info_ref", &G_dbus_node_info_ref},
	{"g_dbus_node_info_unref", &G_dbus_node_info_unref},
	{"g_dbus_property_info_flags_get_type", &G_dbus_property_info_flags_get_type},
	{"g_dbus_property_info_get_type", &G_dbus_property_info_get_type},
	{"g_dbus_property_info_ref", &G_dbus_property_info_ref},
	{"g_dbus_property_info_unref", &G_dbus_property_info_unref},
	{"g_dbus_proxy_call", &G_dbus_proxy_call},
	{"g_dbus_proxy_call_finish", &G_dbus_proxy_call_finish},
	{"g_dbus_proxy_call_sync", &G_dbus_proxy_call_sync},
	{"g_dbus_proxy_flags_get_type", &G_dbus_proxy_flags_get_type},
	{"g_dbus_proxy_get_cached_property", &G_dbus_proxy_get_cached_property},
	{"g_dbus_proxy_get_cached_property_names", &G_dbus_proxy_get_cached_property_names},
	{"g_dbus_proxy_get_connection", &G_dbus_proxy_get_connection},
	{"g_dbus_proxy_get_default_timeout", &G_dbus_proxy_get_default_timeout},
	{"g_dbus_proxy_get_flags", &G_dbus_proxy_get_flags},
	{"g_dbus_proxy_get_interface_info", &G_dbus_proxy_get_interface_info},
	{"g_dbus_proxy_get_interface_name", &G_dbus_proxy_get_interface_name},
	{"g_dbus_proxy_get_name", &G_dbus_proxy_get_name},
	{"g_dbus_proxy_get_name_owner", &G_dbus_proxy_get_name_owner},
	{"g_dbus_proxy_get_object_path", &G_dbus_proxy_get_object_path},
	{"g_dbus_proxy_get_type", &G_dbus_proxy_get_type},
	{"g_dbus_proxy_new", &G_dbus_proxy_new},
	{"g_dbus_proxy_new_finish", &G_dbus_proxy_new_finish},
	{"g_dbus_proxy_new_for_bus", &G_dbus_proxy_new_for_bus},
	{"g_dbus_proxy_new_for_bus_finish", &G_dbus_proxy_new_for_bus_finish},
	{"g_dbus_proxy_new_for_bus_sync", &G_dbus_proxy_new_for_bus_sync},
	{"g_dbus_proxy_new_sync", &G_dbus_proxy_new_sync},
	{"g_dbus_proxy_set_cached_property", &G_dbus_proxy_set_cached_property},
	{"g_dbus_proxy_set_default_timeout", &G_dbus_proxy_set_default_timeout},
	{"g_dbus_proxy_set_interface_info", &G_dbus_proxy_set_interface_info},
	{"g_dbus_send_message_flags_get_type", &G_dbus_send_message_flags_get_type},
	{"g_dbus_server_flags_get_type", &G_dbus_server_flags_get_type},
	{"g_dbus_server_get_client_address", &G_dbus_server_get_client_address},
	{"g_dbus_server_get_flags", &G_dbus_server_get_flags},
	{"g_dbus_server_get_guid", &G_dbus_server_get_guid},
	{"g_dbus_server_get_type", &G_dbus_server_get_type},
	{"g_dbus_server_is_active", &G_dbus_server_is_active},
	{"g_dbus_server_new_sync", &G_dbus_server_new_sync},
	{"g_dbus_server_start", &G_dbus_server_start},
	{"g_dbus_server_stop", &G_dbus_server_stop},
	{"g_dbus_signal_flags_get_type", &G_dbus_signal_flags_get_type},
	{"g_dbus_signal_info_get_type", &G_dbus_signal_info_get_type},
	{"g_dbus_signal_info_ref", &G_dbus_signal_info_ref},
	{"g_dbus_signal_info_unref", &G_dbus_signal_info_unref},
	{"g_dbus_subtree_flags_get_type", &G_dbus_subtree_flags_get_type},
	{"g_drive_can_eject", &G_drive_can_eject},
	{"g_drive_can_poll_for_media", &G_drive_can_poll_for_media},
	{"g_drive_can_start", &G_drive_can_start},
	{"g_drive_can_start_degraded", &G_drive_can_start_degraded},
	{"g_drive_can_stop", &G_drive_can_stop},
	{"g_drive_eject", &G_drive_eject},
	{"g_drive_eject_finish", &G_drive_eject_finish},
	{"g_drive_eject_with_operation", &G_drive_eject_with_operation},
	{"g_drive_eject_with_operation_finish", &G_drive_eject_with_operation_finish},
	{"g_drive_enumerate_identifiers", &G_drive_enumerate_identifiers},
	{"g_drive_get_icon", &G_drive_get_icon},
	{"g_drive_get_identifier", &G_drive_get_identifier},
	{"g_drive_get_name", &G_drive_get_name},
	{"g_drive_get_start_stop_type", &G_drive_get_start_stop_type},
	{"g_drive_get_type", &G_drive_get_type},
	{"g_drive_get_volumes", &G_drive_get_volumes},
	{"g_drive_has_media", &G_drive_has_media},
	{"g_drive_has_volumes", &G_drive_has_volumes},
	{"g_drive_is_media_check_automatic", &G_drive_is_media_check_automatic},
	{"g_drive_is_media_removable", &G_drive_is_media_removable},
	{"g_drive_poll_for_media", &G_drive_poll_for_media},
	{"g_drive_poll_for_media_finish", &G_drive_poll_for_media_finish},
	{"g_drive_start", &G_drive_start},
	{"g_drive_start_finish", &G_drive_start_finish},
	{"g_drive_start_flags_get_type", &G_drive_start_flags_get_type},
	{"g_drive_start_stop_type_get_type", &G_drive_start_stop_type_get_type},
	{"g_drive_stop", &G_drive_stop},
	{"g_drive_stop_finish", &G_drive_stop_finish},
	{"g_emblem_get_icon", &G_emblem_get_icon},
	{"g_emblem_get_origin", &G_emblem_get_origin},
	{"g_emblem_get_type", &G_emblem_get_type},
	{"g_emblem_new", &G_emblem_new},
	{"g_emblem_new_with_origin", &G_emblem_new_with_origin},
	{"g_emblem_origin_get_type", &G_emblem_origin_get_type},
	{"g_emblemed_icon_add_emblem", &G_emblemed_icon_add_emblem},
	{"g_emblemed_icon_clear_emblems", &G_emblemed_icon_clear_emblems},
	{"g_emblemed_icon_get_emblems", &G_emblemed_icon_get_emblems},
	{"g_emblemed_icon_get_icon", &G_emblemed_icon_get_icon},
	{"g_emblemed_icon_get_type", &G_emblemed_icon_get_type},
	{"g_emblemed_icon_new", &G_emblemed_icon_new},
	{"g_file_append_to", &G_file_append_to},
	{"g_file_append_to_async", &G_file_append_to_async},
	{"g_file_append_to_finish", &G_file_append_to_finish},
	{"g_file_attribute_info_flags_get_type", &G_file_attribute_info_flags_get_type},
	{"g_file_attribute_info_list_add", &G_file_attribute_info_list_add},
	{"g_file_attribute_info_list_dup", &G_file_attribute_info_list_dup},
	{"g_file_attribute_info_list_get_type", &G_file_attribute_info_list_get_type},
	{"g_file_attribute_info_list_lookup", &G_file_attribute_info_list_lookup},
	{"g_file_attribute_info_list_new", &G_file_attribute_info_list_new},
	{"g_file_attribute_info_list_ref", &G_file_attribute_info_list_ref},
	{"g_file_attribute_info_list_unref", &G_file_attribute_info_list_unref},
	{"g_file_attribute_matcher_enumerate_namespace", &G_file_attribute_matcher_enumerate_namespace},
	{"g_file_attribute_matcher_enumerate_next", &G_file_attribute_matcher_enumerate_next},
	{"g_file_attribute_matcher_get_type", &G_file_attribute_matcher_get_type},
	{"g_file_attribute_matcher_matches", &G_file_attribute_matcher_matches},
	{"g_file_attribute_matcher_matches_only", &G_file_attribute_matcher_matches_only},
	{"g_file_attribute_matcher_new", &G_file_attribute_matcher_new},
	{"g_file_attribute_matcher_ref", &G_file_attribute_matcher_ref},
	{"g_file_attribute_matcher_unref", &G_file_attribute_matcher_unref},
	{"g_file_attribute_status_get_type", &G_file_attribute_status_get_type},
	{"g_file_attribute_type_get_type", &G_file_attribute_type_get_type},
	{"g_file_copy", &G_file_copy},
	{"g_file_copy_async", &G_file_copy_async},
	{"g_file_copy_attributes", &G_file_copy_attributes},
	{"g_file_copy_finish", &G_file_copy_finish},
	{"g_file_copy_flags_get_type", &G_file_copy_flags_get_type},
	{"g_file_create", &G_file_create},
	{"g_file_create_async", &G_file_create_async},
	{"g_file_create_finish", &G_file_create_finish},
	{"g_file_create_flags_get_type", &G_file_create_flags_get_type},
	{"g_file_create_readwrite", &G_file_create_readwrite},
	{"g_file_create_readwrite_async", &G_file_create_readwrite_async},
	{"g_file_create_readwrite_finish", &G_file_create_readwrite_finish},
	{"g_file_delete", &G_file_delete},
	{"g_file_dup", &G_file_dup},
	{"g_file_eject_mountable", &G_file_eject_mountable},
	{"g_file_eject_mountable_finish", &G_file_eject_mountable_finish},
	{"g_file_eject_mountable_with_operation", &G_file_eject_mountable_with_operation},
	{"g_file_eject_mountable_with_operation_finish", &G_file_eject_mountable_with_operation_finish},
	{"g_file_enumerate_children", &G_file_enumerate_children},
	{"g_file_enumerate_children_async", &G_file_enumerate_children_async},
	{"g_file_enumerate_children_finish", &G_file_enumerate_children_finish},
	{"g_file_enumerator_close", &G_file_enumerator_close},
	{"g_file_enumerator_close_async", &G_file_enumerator_close_async},
	{"g_file_enumerator_close_finish", &G_file_enumerator_close_finish},
	{"g_file_enumerator_get_container", &G_file_enumerator_get_container},
	{"g_file_enumerator_get_type", &G_file_enumerator_get_type},
	{"g_file_enumerator_has_pending", &G_file_enumerator_has_pending},
	{"g_file_enumerator_is_closed", &G_file_enumerator_is_closed},
	{"g_file_enumerator_next_file", &G_file_enumerator_next_file},
	{"g_file_enumerator_next_files_async", &G_file_enumerator_next_files_async},
	{"g_file_enumerator_next_files_finish", &G_file_enumerator_next_files_finish},
	{"g_file_enumerator_set_pending", &G_file_enumerator_set_pending},
	{"g_file_equal", &G_file_equal},
	{"g_file_find_enclosing_mount", &G_file_find_enclosing_mount},
	{"g_file_find_enclosing_mount_async", &G_file_find_enclosing_mount_async},
	{"g_file_find_enclosing_mount_finish", &G_file_find_enclosing_mount_finish},
	{"g_file_get_basename", &G_file_get_basename},
	{"g_file_get_child", &G_file_get_child},
	{"g_file_get_child_for_display_name", &G_file_get_child_for_display_name},
	{"g_file_get_parent", &G_file_get_parent},
	{"g_file_get_parse_name", &G_file_get_parse_name},
	{"g_file_get_path", &G_file_get_path},
	{"g_file_get_relative_path", &G_file_get_relative_path},
	{"g_file_get_type", &G_file_get_type},
	{"g_file_get_uri", &G_file_get_uri},
	{"g_file_get_uri_scheme", &G_file_get_uri_scheme},
	{"g_file_has_parent", &G_file_has_parent},
	{"g_file_has_prefix", &G_file_has_prefix},
	{"g_file_has_uri_scheme", &G_file_has_uri_scheme},
	{"g_file_hash", &G_file_hash},
	{"g_file_icon_get_file", &G_file_icon_get_file},
	{"g_file_icon_get_type", &G_file_icon_get_type},
	{"g_file_icon_new", &G_file_icon_new},
	{"g_file_info_clear_status", &G_file_info_clear_status},
	{"g_file_info_copy_into", &G_file_info_copy_into},
	{"g_file_info_dup", &G_file_info_dup},
	{"g_file_info_get_attribute_as_string", &G_file_info_get_attribute_as_string},
	{"g_file_info_get_attribute_boolean", &G_file_info_get_attribute_boolean},
	{"g_file_info_get_attribute_byte_string", &G_file_info_get_attribute_byte_string},
	{"g_file_info_get_attribute_data", &G_file_info_get_attribute_data},
	{"g_file_info_get_attribute_int32", &G_file_info_get_attribute_int32},
	{"g_file_info_get_attribute_int64", &G_file_info_get_attribute_int64},
	{"g_file_info_get_attribute_object", &G_file_info_get_attribute_object},
	{"g_file_info_get_attribute_status", &G_file_info_get_attribute_status},
	{"g_file_info_get_attribute_string", &G_file_info_get_attribute_string},
	{"g_file_info_get_attribute_stringv", &G_file_info_get_attribute_stringv},
	{"g_file_info_get_attribute_type", &G_file_info_get_attribute_type},
	{"g_file_info_get_attribute_uint32", &G_file_info_get_attribute_uint32},
	{"g_file_info_get_attribute_uint64", &G_file_info_get_attribute_uint64},
	{"g_file_info_get_content_type", &G_file_info_get_content_type},
	{"g_file_info_get_display_name", &G_file_info_get_display_name},
	{"g_file_info_get_edit_name", &G_file_info_get_edit_name},
	{"g_file_info_get_etag", &G_file_info_get_etag},
	{"g_file_info_get_file_type", &G_file_info_get_file_type},
	{"g_file_info_get_icon", &G_file_info_get_icon},
	{"g_file_info_get_is_backup", &G_file_info_get_is_backup},
	{"g_file_info_get_is_hidden", &G_file_info_get_is_hidden},
	{"g_file_info_get_is_symlink", &G_file_info_get_is_symlink},
	{"g_file_info_get_modification_time", &G_file_info_get_modification_time},
	{"g_file_info_get_name", &G_file_info_get_name},
	{"g_file_info_get_size", &G_file_info_get_size},
	{"g_file_info_get_sort_order", &G_file_info_get_sort_order},
	{"g_file_info_get_symlink_target", &G_file_info_get_symlink_target},
	{"g_file_info_get_type", &G_file_info_get_type},
	{"g_file_info_has_attribute", &G_file_info_has_attribute},
	{"g_file_info_has_namespace", &G_file_info_has_namespace},
	{"g_file_info_list_attributes", &G_file_info_list_attributes},
	{"g_file_info_new", &G_file_info_new},
	{"g_file_info_remove_attribute", &G_file_info_remove_attribute},
	{"g_file_info_set_attribute", &G_file_info_set_attribute},
	{"g_file_info_set_attribute_boolean", &G_file_info_set_attribute_boolean},
	{"g_file_info_set_attribute_byte_string", &G_file_info_set_attribute_byte_string},
	{"g_file_info_set_attribute_int32", &G_file_info_set_attribute_int32},
	{"g_file_info_set_attribute_int64", &G_file_info_set_attribute_int64},
	{"g_file_info_set_attribute_mask", &G_file_info_set_attribute_mask},
	{"g_file_info_set_attribute_object", &G_file_info_set_attribute_object},
	{"g_file_info_set_attribute_status", &G_file_info_set_attribute_status},
	{"g_file_info_set_attribute_string", &G_file_info_set_attribute_string},
	{"g_file_info_set_attribute_stringv", &G_file_info_set_attribute_stringv},
	{"g_file_info_set_attribute_uint32", &G_file_info_set_attribute_uint32},
	{"g_file_info_set_attribute_uint64", &G_file_info_set_attribute_uint64},
	{"g_file_info_set_content_type", &G_file_info_set_content_type},
	{"g_file_info_set_display_name", &G_file_info_set_display_name},
	{"g_file_info_set_edit_name", &G_file_info_set_edit_name},
	{"g_file_info_set_file_type", &G_file_info_set_file_type},
	{"g_file_info_set_icon", &G_file_info_set_icon},
	{"g_file_info_set_is_hidden", &G_file_info_set_is_hidden},
	{"g_file_info_set_is_symlink", &G_file_info_set_is_symlink},
	{"g_file_info_set_modification_time", &G_file_info_set_modification_time},
	{"g_file_info_set_name", &G_file_info_set_name},
	{"g_file_info_set_size", &G_file_info_set_size},
	{"g_file_info_set_sort_order", &G_file_info_set_sort_order},
	{"g_file_info_set_symlink_target", &G_file_info_set_symlink_target},
	{"g_file_info_unset_attribute_mask", &G_file_info_unset_attribute_mask},
	{"g_file_input_stream_get_type", &G_file_input_stream_get_type},
	{"g_file_input_stream_query_info", &G_file_input_stream_query_info},
	{"g_file_input_stream_query_info_async", &G_file_input_stream_query_info_async},
	{"g_file_input_stream_query_info_finish", &G_file_input_stream_query_info_finish},
	{"g_file_io_stream_get_etag", &G_file_io_stream_get_etag},
	{"g_file_io_stream_get_type", &G_file_io_stream_get_type},
	{"g_file_io_stream_query_info", &G_file_io_stream_query_info},
	{"g_file_io_stream_query_info_async", &G_file_io_stream_query_info_async},
	{"g_file_io_stream_query_info_finish", &G_file_io_stream_query_info_finish},
	{"g_file_is_native", &G_file_is_native},
	{"g_file_load_contents", &G_file_load_contents},
	{"g_file_load_contents_async", &G_file_load_contents_async},
	{"g_file_load_contents_finish", &G_file_load_contents_finish},
	{"g_file_load_partial_contents_async", &G_file_load_partial_contents_async},
	{"g_file_load_partial_contents_finish", &G_file_load_partial_contents_finish},
	{"g_file_make_directory", &G_file_make_directory},
	{"g_file_make_directory_with_parents", &G_file_make_directory_with_parents},
	{"g_file_make_symbolic_link", &G_file_make_symbolic_link},
	{"g_file_monitor", &G_file_monitor},
	{"g_file_monitor_cancel", &G_file_monitor_cancel},
	{"g_file_monitor_directory", &G_file_monitor_directory},
	{"g_file_monitor_emit_event", &G_file_monitor_emit_event},
	{"g_file_monitor_event_get_type", &G_file_monitor_event_get_type},
	{"g_file_monitor_file", &G_file_monitor_file},
	{"g_file_monitor_flags_get_type", &G_file_monitor_flags_get_type},
	{"g_file_monitor_get_type", &G_file_monitor_get_type},
	{"g_file_monitor_is_cancelled", &G_file_monitor_is_cancelled},
	{"g_file_monitor_set_rate_limit", &G_file_monitor_set_rate_limit},
	{"g_file_mount_enclosing_volume", &G_file_mount_enclosing_volume},
	{"g_file_mount_enclosing_volume_finish", &G_file_mount_enclosing_volume_finish},
	{"g_file_mount_mountable", &G_file_mount_mountable},
	{"g_file_mount_mountable_finish", &G_file_mount_mountable_finish},
	{"g_file_move", &G_file_move},
	{"g_file_new_for_commandline_arg", &G_file_new_for_commandline_arg},
	{"g_file_new_for_path", &G_file_new_for_path},
	{"g_file_new_for_uri", &G_file_new_for_uri},
	{"g_file_open_readwrite", &G_file_open_readwrite},
	{"g_file_open_readwrite_async", &G_file_open_readwrite_async},
	{"g_file_open_readwrite_finish", &G_file_open_readwrite_finish},
	{"g_file_output_stream_get_etag", &G_file_output_stream_get_etag},
	{"g_file_output_stream_get_type", &G_file_output_stream_get_type},
	{"g_file_output_stream_query_info", &G_file_output_stream_query_info},
	{"g_file_output_stream_query_info_async", &G_file_output_stream_query_info_async},
	{"g_file_output_stream_query_info_finish", &G_file_output_stream_query_info_finish},
	{"g_file_parse_name", &G_file_parse_name},
	{"g_file_poll_mountable", &G_file_poll_mountable},
	{"g_file_poll_mountable_finish", &G_file_poll_mountable_finish},
	{"g_file_query_default_handler", &G_file_query_default_handler},
	{"g_file_query_exists", &G_file_query_exists},
	{"g_file_query_file_type", &G_file_query_file_type},
	{"g_file_query_filesystem_info", &G_file_query_filesystem_info},
	{"g_file_query_filesystem_info_async", &G_file_query_filesystem_info_async},
	{"g_file_query_filesystem_info_finish", &G_file_query_filesystem_info_finish},
	{"g_file_query_info", &G_file_query_info},
	{"g_file_query_info_async", &G_file_query_info_async},
	{"g_file_query_info_finish", &G_file_query_info_finish},
	{"g_file_query_info_flags_get_type", &G_file_query_info_flags_get_type},
	{"g_file_query_settable_attributes", &G_file_query_settable_attributes},
	{"g_file_query_writable_namespaces", &G_file_query_writable_namespaces},
	{"g_file_read", &G_file_read},
	{"g_file_read_async", &G_file_read_async},
	{"g_file_read_finish", &G_file_read_finish},
	{"g_file_replace", &G_file_replace},
	{"g_file_replace_async", &G_file_replace_async},
	{"g_file_replace_contents", &G_file_replace_contents},
	{"g_file_replace_contents_async", &G_file_replace_contents_async},
	{"g_file_replace_contents_finish", &G_file_replace_contents_finish},
	{"g_file_replace_finish", &G_file_replace_finish},
	{"g_file_replace_readwrite", &G_file_replace_readwrite},
	{"g_file_replace_readwrite_async", &G_file_replace_readwrite_async},
	{"g_file_replace_readwrite_finish", &G_file_replace_readwrite_finish},
	{"g_file_resolve_relative_path", &G_file_resolve_relative_path},
	{"g_file_set_attribute", &G_file_set_attribute},
	{"g_file_set_attribute_byte_string", &G_file_set_attribute_byte_string},
	{"g_file_set_attribute_int32", &G_file_set_attribute_int32},
	{"g_file_set_attribute_int64", &G_file_set_attribute_int64},
	{"g_file_set_attribute_string", &G_file_set_attribute_string},
	{"g_file_set_attribute_uint32", &G_file_set_attribute_uint32},
	{"g_file_set_attribute_uint64", &G_file_set_attribute_uint64},
	{"g_file_set_attributes_async", &G_file_set_attributes_async},
	{"g_file_set_attributes_finish", &G_file_set_attributes_finish},
	{"g_file_set_attributes_from_info", &G_file_set_attributes_from_info},
	{"g_file_set_display_name", &G_file_set_display_name},
	{"g_file_set_display_name_async", &G_file_set_display_name_async},
	{"g_file_set_display_name_finish", &G_file_set_display_name_finish},
	{"g_file_start_mountable", &G_file_start_mountable},
	{"g_file_start_mountable_finish", &G_file_start_mountable_finish},
	{"g_file_stop_mountable", &G_file_stop_mountable},
	{"g_file_stop_mountable_finish", &G_file_stop_mountable_finish},
	{"g_file_supports_thread_contexts", &G_file_supports_thread_contexts},
	{"g_file_trash", &G_file_trash},
	{"g_file_type_get_type", &G_file_type_get_type},
	{"g_file_unmount_mountable", &G_file_unmount_mountable},
	{"g_file_unmount_mountable_finish", &G_file_unmount_mountable_finish},
	{"g_file_unmount_mountable_with_operation", &G_file_unmount_mountable_with_operation},
	{"g_file_unmount_mountable_with_operation_finish", &G_file_unmount_mountable_with_operation_finish},
	{"g_filename_completer_get_completion_suffix", &G_filename_completer_get_completion_suffix},
	{"g_filename_completer_get_completions", &G_filename_completer_get_completions},
	{"g_filename_completer_get_type", &G_filename_completer_get_type},
	{"g_filename_completer_new", &G_filename_completer_new},
	{"g_filename_completer_set_dirs_only", &G_filename_completer_set_dirs_only},
	{"g_filesystem_preview_type_get_type", &G_filesystem_preview_type_get_type},
	{"g_filter_input_stream_get_base_stream", &G_filter_input_stream_get_base_stream},
	{"g_filter_input_stream_get_close_base_stream", &G_filter_input_stream_get_close_base_stream},
	{"g_filter_input_stream_get_type", &G_filter_input_stream_get_type},
	{"g_filter_input_stream_set_close_base_stream", &G_filter_input_stream_set_close_base_stream},
	{"g_filter_output_stream_get_base_stream", &G_filter_output_stream_get_base_stream},
	{"g_filter_output_stream_get_close_base_stream", &G_filter_output_stream_get_close_base_stream},
	{"g_filter_output_stream_get_type", &G_filter_output_stream_get_type},
	{"g_filter_output_stream_set_close_base_stream", &G_filter_output_stream_set_close_base_stream},
	{"g_icon_equal", &G_icon_equal},
	{"g_icon_get_type", &G_icon_get_type},
	{"g_icon_hash", &G_icon_hash},
	{"g_icon_new_for_string", &G_icon_new_for_string},
	{"g_icon_to_string", &G_icon_to_string},
	{"g_inet_address_get_family", &G_inet_address_get_family},
	{"g_inet_address_get_is_any", &G_inet_address_get_is_any},
	{"g_inet_address_get_is_link_local", &G_inet_address_get_is_link_local},
	{"g_inet_address_get_is_loopback", &G_inet_address_get_is_loopback},
	{"g_inet_address_get_is_mc_global", &G_inet_address_get_is_mc_global},
	{"g_inet_address_get_is_mc_link_local", &G_inet_address_get_is_mc_link_local},
	{"g_inet_address_get_is_mc_node_local", &G_inet_address_get_is_mc_node_local},
	{"g_inet_address_get_is_mc_org_local", &G_inet_address_get_is_mc_org_local},
	{"g_inet_address_get_is_mc_site_local", &G_inet_address_get_is_mc_site_local},
	{"g_inet_address_get_is_multicast", &G_inet_address_get_is_multicast},
	{"g_inet_address_get_is_site_local", &G_inet_address_get_is_site_local},
	{"g_inet_address_get_native_size", &G_inet_address_get_native_size},
	{"g_inet_address_get_type", &G_inet_address_get_type},
	{"g_inet_address_new_any", &G_inet_address_new_any},
	{"g_inet_address_new_from_bytes", &G_inet_address_new_from_bytes},
	{"g_inet_address_new_from_string", &G_inet_address_new_from_string},
	{"g_inet_address_new_loopback", &G_inet_address_new_loopback},
	{"g_inet_address_to_bytes", &G_inet_address_to_bytes},
	{"g_inet_address_to_string", &G_inet_address_to_string},
	{"g_inet_socket_address_get_address", &G_inet_socket_address_get_address},
	{"g_inet_socket_address_get_port", &G_inet_socket_address_get_port},
	{"g_inet_socket_address_get_type", &G_inet_socket_address_get_type},
	{"g_inet_socket_address_new", &G_inet_socket_address_new},
	{"g_initable_get_type", &G_initable_get_type},
	{"g_initable_init", &G_initable_init},
	{"g_initable_new", &G_initable_new},
	{"g_initable_new_valist", &G_initable_new_valist},
	{"g_initable_newv", &G_initable_newv},
	{"g_input_stream_clear_pending", &G_input_stream_clear_pending},
	{"g_input_stream_close", &G_input_stream_close},
	{"g_input_stream_close_async", &G_input_stream_close_async},
	{"g_input_stream_close_finish", &G_input_stream_close_finish},
	{"g_input_stream_get_type", &G_input_stream_get_type},
	{"g_input_stream_has_pending", &G_input_stream_has_pending},
	{"g_input_stream_is_closed", &G_input_stream_is_closed},
	{"g_input_stream_read", &G_input_stream_read},
	{"g_input_stream_read_all", &G_input_stream_read_all},
	{"g_input_stream_read_async", &G_input_stream_read_async},
	{"g_input_stream_read_finish", &G_input_stream_read_finish},
	{"g_input_stream_set_pending", &G_input_stream_set_pending},
	{"g_input_stream_skip", &G_input_stream_skip},
	{"g_input_stream_skip_async", &G_input_stream_skip_async},
	{"g_input_stream_skip_finish", &G_input_stream_skip_finish},
	{"g_io_error_enum_get_type", &G_io_error_enum_get_type},
	{"g_io_error_from_errno", &G_io_error_from_errno},
	{"g_io_error_from_win32_error", &G_io_error_from_win32_error},
	{"g_io_error_quark", &G_io_error_quark},
	{"g_io_extension_get_name", &G_io_extension_get_name},
	{"g_io_extension_get_priority", &G_io_extension_get_priority},
	{"g_io_extension_get_type", &G_io_extension_get_type},
	{"g_io_extension_point_get_extension_by_name", &G_io_extension_point_get_extension_by_name},
	{"g_io_extension_point_get_extensions", &G_io_extension_point_get_extensions},
	{"g_io_extension_point_get_required_type", &G_io_extension_point_get_required_type},
	{"g_io_extension_point_implement", &G_io_extension_point_implement},
	{"g_io_extension_point_lookup", &G_io_extension_point_lookup},
	{"g_io_extension_point_register", &G_io_extension_point_register},
	{"g_io_extension_point_set_required_type", &G_io_extension_point_set_required_type},
	{"g_io_extension_ref_class", &G_io_extension_ref_class},
	{"g_io_module_get_type", &G_io_module_get_type},
	{"g_io_module_new", &G_io_module_new},
	{"g_io_modules_load_all_in_directory", &G_io_modules_load_all_in_directory},
	{"g_io_modules_scan_all_in_directory", &G_io_modules_scan_all_in_directory},
	{"g_io_scheduler_cancel_all_jobs", &G_io_scheduler_cancel_all_jobs},
	{"g_io_scheduler_job_send_to_mainloop", &G_io_scheduler_job_send_to_mainloop},
	{"g_io_scheduler_job_send_to_mainloop_async", &G_io_scheduler_job_send_to_mainloop_async},
	{"g_io_scheduler_push_job", &G_io_scheduler_push_job},
	{"g_io_stream_clear_pending", &G_io_stream_clear_pending},
	{"g_io_stream_close", &G_io_stream_close},
	{"g_io_stream_close_async", &G_io_stream_close_async},
	{"g_io_stream_close_finish", &G_io_stream_close_finish},
	{"g_io_stream_get_input_stream", &G_io_stream_get_input_stream},
	{"g_io_stream_get_output_stream", &G_io_stream_get_output_stream},
	{"g_io_stream_get_type", &G_io_stream_get_type},
	{"g_io_stream_has_pending", &G_io_stream_has_pending},
	{"g_io_stream_is_closed", &G_io_stream_is_closed},
	{"g_io_stream_set_pending", &G_io_stream_set_pending},
	{"g_io_stream_splice_async", &G_io_stream_splice_async},
	{"g_io_stream_splice_finish", &G_io_stream_splice_finish},
	{"g_io_stream_splice_flags_get_type", &G_io_stream_splice_flags_get_type},
	{"g_keyfile_settings_backend_new", &G_keyfile_settings_backend_new},
	{"g_loadable_icon_get_type", &G_loadable_icon_get_type},
	{"g_loadable_icon_load", &G_loadable_icon_load},
	{"g_loadable_icon_load_async", &G_loadable_icon_load_async},
	{"g_loadable_icon_load_finish", &G_loadable_icon_load_finish},
	{"g_memory_input_stream_add_data", &G_memory_input_stream_add_data},
	{"g_memory_input_stream_get_type", &G_memory_input_stream_get_type},
	{"g_memory_input_stream_new", &G_memory_input_stream_new},
	{"g_memory_input_stream_new_from_data", &G_memory_input_stream_new_from_data},
	{"g_memory_output_stream_get_data", &G_memory_output_stream_get_data},
	{"g_memory_output_stream_get_data_size", &G_memory_output_stream_get_data_size},
	{"g_memory_output_stream_get_size", &G_memory_output_stream_get_size},
	{"g_memory_output_stream_get_type", &G_memory_output_stream_get_type},
	{"g_memory_output_stream_new", &G_memory_output_stream_new},
	{"g_memory_output_stream_steal_data", &G_memory_output_stream_steal_data},
	{"g_memory_settings_backend_new", &G_memory_settings_backend_new},
	{"g_mount_can_eject", &G_mount_can_eject},
	{"g_mount_can_unmount", &G_mount_can_unmount},
	{"g_mount_eject", &G_mount_eject},
	{"g_mount_eject_finish", &G_mount_eject_finish},
	{"g_mount_eject_with_operation", &G_mount_eject_with_operation},
	{"g_mount_eject_with_operation_finish", &G_mount_eject_with_operation_finish},
	{"g_mount_get_default_location", &G_mount_get_default_location},
	{"g_mount_get_drive", &G_mount_get_drive},
	{"g_mount_get_icon", &G_mount_get_icon},
	{"g_mount_get_name", &G_mount_get_name},
	{"g_mount_get_root", &G_mount_get_root},
	{"g_mount_get_type", &G_mount_get_type},
	{"g_mount_get_uuid", &G_mount_get_uuid},
	{"g_mount_get_volume", &G_mount_get_volume},
	{"g_mount_guess_content_type", &G_mount_guess_content_type},
	{"g_mount_guess_content_type_finish", &G_mount_guess_content_type_finish},
	{"g_mount_guess_content_type_sync", &G_mount_guess_content_type_sync},
	{"g_mount_is_shadowed", &G_mount_is_shadowed},
	{"g_mount_mount_flags_get_type", &G_mount_mount_flags_get_type},
	{"g_mount_operation_get_anonymous", &G_mount_operation_get_anonymous},
	{"g_mount_operation_get_choice", &G_mount_operation_get_choice},
	{"g_mount_operation_get_domain", &G_mount_operation_get_domain},
	{"g_mount_operation_get_password", &G_mount_operation_get_password},
	{"g_mount_operation_get_password_save", &G_mount_operation_get_password_save},
	{"g_mount_operation_get_type", &G_mount_operation_get_type},
	{"g_mount_operation_get_username", &G_mount_operation_get_username},
	{"g_mount_operation_new", &G_mount_operation_new},
	{"g_mount_operation_reply", &G_mount_operation_reply},
	{"g_mount_operation_result_get_type", &G_mount_operation_result_get_type},
	{"g_mount_operation_set_anonymous", &G_mount_operation_set_anonymous},
	{"g_mount_operation_set_choice", &G_mount_operation_set_choice},
	{"g_mount_operation_set_domain", &G_mount_operation_set_domain},
	{"g_mount_operation_set_password", &G_mount_operation_set_password},
	{"g_mount_operation_set_password_save", &G_mount_operation_set_password_save},
	{"g_mount_operation_set_username", &G_mount_operation_set_username},
	{"g_mount_remount", &G_mount_remount},
	{"g_mount_remount_finish", &G_mount_remount_finish},
	{"g_mount_shadow", &G_mount_shadow},
	{"g_mount_unmount", &G_mount_unmount},
	{"g_mount_unmount_finish", &G_mount_unmount_finish},
	{"g_mount_unmount_flags_get_type", &G_mount_unmount_flags_get_type},
	{"g_mount_unmount_with_operation", &G_mount_unmount_with_operation},
	{"g_mount_unmount_with_operation_finish", &G_mount_unmount_with_operation_finish},
	{"g_mount_unshadow", &G_mount_unshadow},
	{"g_native_volume_monitor_get_type", &G_native_volume_monitor_get_type},
	{"g_network_address_get_hostname", &G_network_address_get_hostname},
	{"g_network_address_get_port", &G_network_address_get_port},
	{"g_network_address_get_scheme", &G_network_address_get_scheme},
	{"g_network_address_get_type", &G_network_address_get_type},
	{"g_network_address_new", &G_network_address_new},
	{"g_network_address_parse", &G_network_address_parse},
	{"g_network_address_parse_uri", &G_network_address_parse_uri},
	{"g_network_service_get_domain", &G_network_service_get_domain},
	{"g_network_service_get_protocol", &G_network_service_get_protocol},
	{"g_network_service_get_scheme", &G_network_service_get_scheme},
	{"g_network_service_get_service", &G_network_service_get_service},
	{"g_network_service_get_type", &G_network_service_get_type},
	{"g_network_service_new", &G_network_service_new},
	{"g_network_service_set_scheme", &G_network_service_set_scheme},
	{"g_null_settings_backend_new", &G_null_settings_backend_new},
	{"g_output_stream_clear_pending", &G_output_stream_clear_pending},
	{"g_output_stream_close", &G_output_stream_close},
	{"g_output_stream_close_async", &G_output_stream_close_async},
	{"g_output_stream_close_finish", &G_output_stream_close_finish},
	{"g_output_stream_flush", &G_output_stream_flush},
	{"g_output_stream_flush_async", &G_output_stream_flush_async},
	{"g_output_stream_flush_finish", &G_output_stream_flush_finish},
	{"g_output_stream_get_type", &G_output_stream_get_type},
	{"g_output_stream_has_pending", &G_output_stream_has_pending},
	{"g_output_stream_is_closed", &G_output_stream_is_closed},
	{"g_output_stream_is_closing", &G_output_stream_is_closing},
	{"g_output_stream_set_pending", &G_output_stream_set_pending},
	{"g_output_stream_splice", &G_output_stream_splice},
	{"g_output_stream_splice_async", &G_output_stream_splice_async},
	{"g_output_stream_splice_finish", &G_output_stream_splice_finish},
	{"g_output_stream_splice_flags_get_type", &G_output_stream_splice_flags_get_type},
	{"g_output_stream_write", &G_output_stream_write},
	{"g_output_stream_write_all", &G_output_stream_write_all},
	{"g_output_stream_write_async", &G_output_stream_write_async},
	{"g_output_stream_write_finish", &G_output_stream_write_finish},
	{"g_password_save_get_type", &G_password_save_get_type},
	{"g_permission_acquire", &G_permission_acquire},
	{"g_permission_acquire_async", &G_permission_acquire_async},
	{"g_permission_acquire_finish", &G_permission_acquire_finish},
	{"g_permission_get_allowed", &G_permission_get_allowed},
	{"g_permission_get_can_acquire", &G_permission_get_can_acquire},
	{"g_permission_get_can_release", &G_permission_get_can_release},
	{"g_permission_get_type", &G_permission_get_type},
	{"g_permission_impl_update", &G_permission_impl_update},
	{"g_permission_release", &G_permission_release},
	{"g_permission_release_async", &G_permission_release_async},
	{"g_permission_release_finish", &G_permission_release_finish},
	{"g_pollable_input_stream_can_poll", &G_pollable_input_stream_can_poll},
	{"g_pollable_input_stream_create_source", &G_pollable_input_stream_create_source},
	{"g_pollable_input_stream_get_type", &G_pollable_input_stream_get_type},
	{"g_pollable_input_stream_is_readable", &G_pollable_input_stream_is_readable},
	{"g_pollable_input_stream_read_nonblocking", &G_pollable_input_stream_read_nonblocking},
	{"g_pollable_output_stream_can_poll", &G_pollable_output_stream_can_poll},
	{"g_pollable_output_stream_create_source", &G_pollable_output_stream_create_source},
	{"g_pollable_output_stream_get_type", &G_pollable_output_stream_get_type},
	{"g_pollable_output_stream_is_writable", &G_pollable_output_stream_is_writable},
	{"g_pollable_output_stream_write_nonblocking", &G_pollable_output_stream_write_nonblocking},
	{"g_pollable_source_new", &G_pollable_source_new},
	{"g_proxy_address_enumerator_get_type", &G_proxy_address_enumerator_get_type},
	{"g_proxy_address_get_destination_hostname", &G_proxy_address_get_destination_hostname},
	{"g_proxy_address_get_destination_port", &G_proxy_address_get_destination_port},
	{"g_proxy_address_get_password", &G_proxy_address_get_password},
	{"g_proxy_address_get_protocol", &G_proxy_address_get_protocol},
	{"g_proxy_address_get_type", &G_proxy_address_get_type},
	{"g_proxy_address_get_username", &G_proxy_address_get_username},
	{"g_proxy_address_new", &G_proxy_address_new},
	{"g_proxy_connect", &G_proxy_connect},
	{"g_proxy_connect_async", &G_proxy_connect_async},
	{"g_proxy_connect_finish", &G_proxy_connect_finish},
	{"g_proxy_get_default_for_protocol", &G_proxy_get_default_for_protocol},
	{"g_proxy_get_type", &G_proxy_get_type},
	{"g_proxy_resolver_get_default", &G_proxy_resolver_get_default},
	{"g_proxy_resolver_get_type", &G_proxy_resolver_get_type},
	{"g_proxy_resolver_is_supported", &G_proxy_resolver_is_supported},
	{"g_proxy_resolver_lookup", &G_proxy_resolver_lookup},
	{"g_proxy_resolver_lookup_async", &G_proxy_resolver_lookup_async},
	{"g_proxy_resolver_lookup_finish", &G_proxy_resolver_lookup_finish},
	{"g_proxy_supports_hostname", &G_proxy_supports_hostname},
	{"g_resolver_error_get_type", &G_resolver_error_get_type},
	{"g_resolver_error_quark", &G_resolver_error_quark},
	{"g_resolver_free_addresses", &G_resolver_free_addresses},
	{"g_resolver_free_targets", &G_resolver_free_targets},
	{"g_resolver_get_default", &G_resolver_get_default},
	{"g_resolver_get_type", &G_resolver_get_type},
	{"g_resolver_lookup_by_address", &G_resolver_lookup_by_address},
	{"g_resolver_lookup_by_address_async", &G_resolver_lookup_by_address_async},
	{"g_resolver_lookup_by_address_finish", &G_resolver_lookup_by_address_finish},
	{"g_resolver_lookup_by_name", &G_resolver_lookup_by_name},
	{"g_resolver_lookup_by_name_async", &G_resolver_lookup_by_name_async},
	{"g_resolver_lookup_by_name_finish", &G_resolver_lookup_by_name_finish},
	{"g_resolver_lookup_service", &G_resolver_lookup_service},
	{"g_resolver_lookup_service_async", &G_resolver_lookup_service_async},
	{"g_resolver_lookup_service_finish", &G_resolver_lookup_service_finish},
	{"g_resolver_set_default", &G_resolver_set_default},
	{"g_seekable_can_seek", &G_seekable_can_seek},
	{"g_seekable_can_truncate", &G_seekable_can_truncate},
	{"g_seekable_get_type", &G_seekable_get_type},
	{"g_seekable_seek", &G_seekable_seek},
	{"g_seekable_tell", &G_seekable_tell},
	{"g_seekable_truncate", &G_seekable_truncate},
	{"g_settings_apply", &G_settings_apply},
	{"g_settings_backend_changed", &G_settings_backend_changed},
	{"g_settings_backend_changed_tree", &G_settings_backend_changed_tree},
	{"g_settings_backend_flatten_tree", &G_settings_backend_flatten_tree},
	{"g_settings_backend_get_default", &G_settings_backend_get_default},
	{"g_settings_backend_get_type", &G_settings_backend_get_type},
	{"g_settings_backend_keys_changed", &G_settings_backend_keys_changed},
	{"g_settings_backend_path_changed", &G_settings_backend_path_changed},
	{"g_settings_backend_path_writable_changed", &G_settings_backend_path_writable_changed},
	{"g_settings_backend_writable_changed", &G_settings_backend_writable_changed},
	{"g_settings_bind", &G_settings_bind},
	{"g_settings_bind_flags_get_type", &G_settings_bind_flags_get_type},
	{"g_settings_bind_with_mapping", &G_settings_bind_with_mapping},
	{"g_settings_bind_writable", &G_settings_bind_writable},
	{"g_settings_delay", &G_settings_delay},
	{"g_settings_get", &G_settings_get},
	{"g_settings_get_boolean", &G_settings_get_boolean},
	{"g_settings_get_child", &G_settings_get_child},
	{"g_settings_get_double", &G_settings_get_double},
	{"g_settings_get_enum", &G_settings_get_enum},
	{"g_settings_get_flags", &G_settings_get_flags},
	{"g_settings_get_has_unapplied", &G_settings_get_has_unapplied},
	{"g_settings_get_int", &G_settings_get_int},
	{"g_settings_get_mapped", &G_settings_get_mapped},
	{"g_settings_get_range", &G_settings_get_range},
	{"g_settings_get_string", &G_settings_get_string},
	{"g_settings_get_strv", &G_settings_get_strv},
	{"g_settings_get_type", &G_settings_get_type},
	{"g_settings_get_value", &G_settings_get_value},
	{"g_settings_is_writable", &G_settings_is_writable},
	{"g_settings_list_children", &G_settings_list_children},
	{"g_settings_list_keys", &G_settings_list_keys},
	{"g_settings_list_relocatable_schemas", &G_settings_list_relocatable_schemas},
	{"g_settings_list_schemas", &G_settings_list_schemas},
	{"g_settings_new", &G_settings_new},
	{"g_settings_new_with_backend", &G_settings_new_with_backend},
	{"g_settings_new_with_backend_and_path", &G_settings_new_with_backend_and_path},
	{"g_settings_new_with_path", &G_settings_new_with_path},
	{"g_settings_range_check", &G_settings_range_check},
	{"g_settings_reset", &G_settings_reset},
	{"g_settings_revert", &G_settings_revert},
	{"g_settings_set", &G_settings_set},
	{"g_settings_set_boolean", &G_settings_set_boolean},
	{"g_settings_set_double", &G_settings_set_double},
	{"g_settings_set_enum", &G_settings_set_enum},
	{"g_settings_set_flags", &G_settings_set_flags},
	{"g_settings_set_int", &G_settings_set_int},
	{"g_settings_set_string", &G_settings_set_string},
	{"g_settings_set_strv", &G_settings_set_strv},
	{"g_settings_set_value", &G_settings_set_value},
	{"g_settings_sync", &G_settings_sync},
	{"g_settings_unbind", &G_settings_unbind},
	// Undocumented {"g_simple_action_get_parameter_type", &G_simple_action_get_parameter_type},
	{"g_simple_action_get_type", &G_simple_action_get_type},
	{"g_simple_action_group_get_type", &G_simple_action_group_get_type},
	{"g_simple_action_group_insert", &G_simple_action_group_insert},
	{"g_simple_action_group_lookup", &G_simple_action_group_lookup},
	{"g_simple_action_group_new", &G_simple_action_group_new},
	{"g_simple_action_group_remove", &G_simple_action_group_remove},
	{"g_simple_action_new", &G_simple_action_new},
	{"g_simple_action_new_stateful", &G_simple_action_new_stateful},
	{"g_simple_action_set_enabled", &G_simple_action_set_enabled},
	{"g_simple_async_report_error_in_idle", &G_simple_async_report_error_in_idle},
	{"g_simple_async_report_gerror_in_idle", &G_simple_async_report_gerror_in_idle},
	{"g_simple_async_report_take_gerror_in_idle", &G_simple_async_report_take_gerror_in_idle},
	{"g_simple_async_result_complete", &G_simple_async_result_complete},
	{"g_simple_async_result_complete_in_idle", &G_simple_async_result_complete_in_idle},
	{"g_simple_async_result_get_op_res_gboolean", &G_simple_async_result_get_op_res_gboolean},
	{"g_simple_async_result_get_op_res_gpointer", &G_simple_async_result_get_op_res_gpointer},
	{"g_simple_async_result_get_op_res_gssize", &G_simple_async_result_get_op_res_gssize},
	{"g_simple_async_result_get_source_tag", &G_simple_async_result_get_source_tag},
	{"g_simple_async_result_get_type", &G_simple_async_result_get_type},
	{"g_simple_async_result_is_valid", &G_simple_async_result_is_valid},
	{"g_simple_async_result_new", &G_simple_async_result_new},
	{"g_simple_async_result_new_error", &G_simple_async_result_new_error},
	{"g_simple_async_result_new_from_error", &G_simple_async_result_new_from_error},
	{"g_simple_async_result_new_take_error", &G_simple_async_result_new_take_error},
	{"g_simple_async_result_propagate_error", &G_simple_async_result_propagate_error},
	{"g_simple_async_result_run_in_thread", &G_simple_async_result_run_in_thread},
	{"g_simple_async_result_set_error", &G_simple_async_result_set_error},
	{"g_simple_async_result_set_error_va", &G_simple_async_result_set_error_va},
	{"g_simple_async_result_set_from_error", &G_simple_async_result_set_from_error},
	{"g_simple_async_result_set_handle_cancellation", &G_simple_async_result_set_handle_cancellation},
	{"g_simple_async_result_set_op_res_gboolean", &G_simple_async_result_set_op_res_gboolean},
	{"g_simple_async_result_set_op_res_gpointer", &G_simple_async_result_set_op_res_gpointer},
	{"g_simple_async_result_set_op_res_gssize", &G_simple_async_result_set_op_res_gssize},
	{"g_simple_async_result_take_error", &G_simple_async_result_take_error},
	{"g_simple_permission_get_type", &G_simple_permission_get_type},
	{"g_simple_permission_new", &G_simple_permission_new},
	{"g_socket_accept", &G_socket_accept},
	{"g_socket_address_enumerator_get_type", &G_socket_address_enumerator_get_type},
	{"g_socket_address_enumerator_next", &G_socket_address_enumerator_next},
	{"g_socket_address_enumerator_next_async", &G_socket_address_enumerator_next_async},
	{"g_socket_address_enumerator_next_finish", &G_socket_address_enumerator_next_finish},
	{"g_socket_address_get_family", &G_socket_address_get_family},
	{"g_socket_address_get_native_size", &G_socket_address_get_native_size},
	{"g_socket_address_get_type", &G_socket_address_get_type},
	{"g_socket_address_new_from_native", &G_socket_address_new_from_native},
	{"g_socket_address_to_native", &G_socket_address_to_native},
	{"g_socket_bind", &G_socket_bind},
	{"g_socket_check_connect_result", &G_socket_check_connect_result},
	{"g_socket_client_add_application_proxy", &G_socket_client_add_application_proxy},
	{"g_socket_client_connect", &G_socket_client_connect},
	{"g_socket_client_connect_async", &G_socket_client_connect_async},
	{"g_socket_client_connect_finish", &G_socket_client_connect_finish},
	{"g_socket_client_connect_to_host", &G_socket_client_connect_to_host},
	{"g_socket_client_connect_to_host_async", &G_socket_client_connect_to_host_async},
	{"g_socket_client_connect_to_host_finish", &G_socket_client_connect_to_host_finish},
	{"g_socket_client_connect_to_service", &G_socket_client_connect_to_service},
	{"g_socket_client_connect_to_service_async", &G_socket_client_connect_to_service_async},
	{"g_socket_client_connect_to_service_finish", &G_socket_client_connect_to_service_finish},
	{"g_socket_client_connect_to_uri", &G_socket_client_connect_to_uri},
	{"g_socket_client_connect_to_uri_async", &G_socket_client_connect_to_uri_async},
	{"g_socket_client_connect_to_uri_finish", &G_socket_client_connect_to_uri_finish},
	{"g_socket_client_get_enable_proxy", &G_socket_client_get_enable_proxy},
	{"g_socket_client_get_family", &G_socket_client_get_family},
	{"g_socket_client_get_local_address", &G_socket_client_get_local_address},
	{"g_socket_client_get_protocol", &G_socket_client_get_protocol},
	{"g_socket_client_get_socket_type", &G_socket_client_get_socket_type},
	{"g_socket_client_get_timeout", &G_socket_client_get_timeout},
	{"g_socket_client_get_tls", &G_socket_client_get_tls},
	{"g_socket_client_get_tls_validation_flags", &G_socket_client_get_tls_validation_flags},
	{"g_socket_client_get_type", &G_socket_client_get_type},
	{"g_socket_client_new", &G_socket_client_new},
	{"g_socket_client_set_enable_proxy", &G_socket_client_set_enable_proxy},
	{"g_socket_client_set_family", &G_socket_client_set_family},
	{"g_socket_client_set_local_address", &G_socket_client_set_local_address},
	{"g_socket_client_set_protocol", &G_socket_client_set_protocol},
	{"g_socket_client_set_socket_type", &G_socket_client_set_socket_type},
	{"g_socket_client_set_timeout", &G_socket_client_set_timeout},
	{"g_socket_client_set_tls", &G_socket_client_set_tls},
	{"g_socket_client_set_tls_validation_flags", &G_socket_client_set_tls_validation_flags},
	{"g_socket_close", &G_socket_close},
	{"g_socket_condition_check", &G_socket_condition_check},
	{"g_socket_condition_wait", &G_socket_condition_wait},
	{"g_socket_connect", &G_socket_connect},
	{"g_socket_connectable_enumerate", &G_socket_connectable_enumerate},
	{"g_socket_connectable_get_type", &G_socket_connectable_get_type},
	{"g_socket_connectable_proxy_enumerate", &G_socket_connectable_proxy_enumerate},
	{"g_socket_connection_factory_create_connection", &G_socket_connection_factory_create_connection},
	{"g_socket_connection_factory_lookup_type", &G_socket_connection_factory_lookup_type},
	{"g_socket_connection_factory_register_type", &G_socket_connection_factory_register_type},
	{"g_socket_connection_get_local_address", &G_socket_connection_get_local_address},
	{"g_socket_connection_get_remote_address", &G_socket_connection_get_remote_address},
	{"g_socket_connection_get_socket", &G_socket_connection_get_socket},
	{"g_socket_connection_get_type", &G_socket_connection_get_type},
	{"g_socket_control_message_deserialize", &G_socket_control_message_deserialize},
	{"g_socket_control_message_get_level", &G_socket_control_message_get_level},
	{"g_socket_control_message_get_msg_type", &G_socket_control_message_get_msg_type},
	{"g_socket_control_message_get_size", &G_socket_control_message_get_size},
	{"g_socket_control_message_get_type", &G_socket_control_message_get_type},
	{"g_socket_control_message_serialize", &G_socket_control_message_serialize},
	{"g_socket_create_source", &G_socket_create_source},
	{"g_socket_family_get_type", &G_socket_family_get_type},
	{"g_socket_get_blocking", &G_socket_get_blocking},
	{"g_socket_get_credentials", &G_socket_get_credentials},
	{"g_socket_get_family", &G_socket_get_family},
	{"g_socket_get_fd", &G_socket_get_fd},
	{"g_socket_get_keepalive", &G_socket_get_keepalive},
	{"g_socket_get_listen_backlog", &G_socket_get_listen_backlog},
	{"g_socket_get_local_address", &G_socket_get_local_address},
	{"g_socket_get_protocol", &G_socket_get_protocol},
	{"g_socket_get_remote_address", &G_socket_get_remote_address},
	{"g_socket_get_socket_type", &G_socket_get_socket_type},
	{"g_socket_get_timeout", &G_socket_get_timeout},
	{"g_socket_get_type", &G_socket_get_type},
	{"g_socket_is_closed", &G_socket_is_closed},
	{"g_socket_is_connected", &G_socket_is_connected},
	{"g_socket_listen", &G_socket_listen},
	{"g_socket_listener_accept", &G_socket_listener_accept},
	{"g_socket_listener_accept_async", &G_socket_listener_accept_async},
	{"g_socket_listener_accept_finish", &G_socket_listener_accept_finish},
	{"g_socket_listener_accept_socket", &G_socket_listener_accept_socket},
	{"g_socket_listener_accept_socket_async", &G_socket_listener_accept_socket_async},
	{"g_socket_listener_accept_socket_finish", &G_socket_listener_accept_socket_finish},
	{"g_socket_listener_add_address", &G_socket_listener_add_address},
	{"g_socket_listener_add_any_inet_port", &G_socket_listener_add_any_inet_port},
	{"g_socket_listener_add_inet_port", &G_socket_listener_add_inet_port},
	{"g_socket_listener_add_socket", &G_socket_listener_add_socket},
	{"g_socket_listener_close", &G_socket_listener_close},
	{"g_socket_listener_get_type", &G_socket_listener_get_type},
	{"g_socket_listener_new", &G_socket_listener_new},
	{"g_socket_listener_set_backlog", &G_socket_listener_set_backlog},
	{"g_socket_msg_flags_get_type", &G_socket_msg_flags_get_type},
	{"g_socket_new", &G_socket_new},
	{"g_socket_new_from_fd", &G_socket_new_from_fd},
	{"g_socket_protocol_get_type", &G_socket_protocol_get_type},
	{"g_socket_receive", &G_socket_receive},
	{"g_socket_receive_from", &G_socket_receive_from},
	{"g_socket_receive_message", &G_socket_receive_message},
	{"g_socket_receive_with_blocking", &G_socket_receive_with_blocking},
	{"g_socket_send", &G_socket_send},
	{"g_socket_send_message", &G_socket_send_message},
	{"g_socket_send_to", &G_socket_send_to},
	{"g_socket_send_with_blocking", &G_socket_send_with_blocking},
	{"g_socket_service_get_type", &G_socket_service_get_type},
	{"g_socket_service_is_active", &G_socket_service_is_active},
	{"g_socket_service_new", &G_socket_service_new},
	{"g_socket_service_start", &G_socket_service_start},
	{"g_socket_service_stop", &G_socket_service_stop},
	{"g_socket_set_blocking", &G_socket_set_blocking},
	{"g_socket_set_keepalive", &G_socket_set_keepalive},
	{"g_socket_set_listen_backlog", &G_socket_set_listen_backlog},
	{"g_socket_set_timeout", &G_socket_set_timeout},
	{"g_socket_shutdown", &G_socket_shutdown},
	{"g_socket_speaks_ipv4", &G_socket_speaks_ipv4},
	{"g_socket_type_get_type", &G_socket_type_get_type},
	{"g_srv_target_copy", &G_srv_target_copy},
	{"g_srv_target_free", &G_srv_target_free},
	{"g_srv_target_get_hostname", &G_srv_target_get_hostname},
	{"g_srv_target_get_port", &G_srv_target_get_port},
	{"g_srv_target_get_priority", &G_srv_target_get_priority},
	{"g_srv_target_get_type", &G_srv_target_get_type},
	{"g_srv_target_get_weight", &G_srv_target_get_weight},
	{"g_srv_target_list_sort", &G_srv_target_list_sort},
	{"g_srv_target_new", &G_srv_target_new},
	{"g_tcp_connection_get_graceful_disconnect", &G_tcp_connection_get_graceful_disconnect},
	{"g_tcp_connection_get_type", &G_tcp_connection_get_type},
	{"g_tcp_connection_set_graceful_disconnect", &G_tcp_connection_set_graceful_disconnect},
	{"g_tcp_wrapper_connection_get_base_io_stream", &G_tcp_wrapper_connection_get_base_io_stream},
	{"g_tcp_wrapper_connection_get_type", &G_tcp_wrapper_connection_get_type},
	{"g_tcp_wrapper_connection_new", &G_tcp_wrapper_connection_new},
	{"g_themed_icon_append_name", &G_themed_icon_append_name},
	{"g_themed_icon_get_names", &G_themed_icon_get_names},
	{"g_themed_icon_get_type", &G_themed_icon_get_type},
	{"g_themed_icon_new", &G_themed_icon_new},
	{"g_themed_icon_new_from_names", &G_themed_icon_new_from_names},
	{"g_themed_icon_new_with_default_fallbacks", &G_themed_icon_new_with_default_fallbacks},
	{"g_themed_icon_prepend_name", &G_themed_icon_prepend_name},
	// Undocumented {"g_threaded_resolver_get_type", &G_threaded_resolver_get_type},
	{"g_threaded_socket_service_get_type", &G_threaded_socket_service_get_type},
	{"g_threaded_socket_service_new", &G_threaded_socket_service_new},
	{"g_tls_authentication_mode_get_type", &G_tls_authentication_mode_get_type},
	{"g_tls_backend_get_certificate_type", &G_tls_backend_get_certificate_type},
	{"g_tls_backend_get_client_connection_type", &G_tls_backend_get_client_connection_type},
	{"g_tls_backend_get_default", &G_tls_backend_get_default},
	{"g_tls_backend_get_server_connection_type", &G_tls_backend_get_server_connection_type},
	{"g_tls_backend_get_type", &G_tls_backend_get_type},
	{"g_tls_backend_supports_tls", &G_tls_backend_supports_tls},
	{"g_tls_certificate_flags_get_type", &G_tls_certificate_flags_get_type},
	{"g_tls_certificate_get_issuer", &G_tls_certificate_get_issuer},
	{"g_tls_certificate_get_type", &G_tls_certificate_get_type},
	{"g_tls_certificate_list_new_from_file", &G_tls_certificate_list_new_from_file},
	{"g_tls_certificate_new_from_file", &G_tls_certificate_new_from_file},
	{"g_tls_certificate_new_from_files", &G_tls_certificate_new_from_files},
	{"g_tls_certificate_new_from_pem", &G_tls_certificate_new_from_pem},
	{"g_tls_certificate_verify", &G_tls_certificate_verify},
	{"g_tls_client_connection_get_accepted_cas", &G_tls_client_connection_get_accepted_cas},
	{"g_tls_client_connection_get_server_identity", &G_tls_client_connection_get_server_identity},
	{"g_tls_client_connection_get_type", &G_tls_client_connection_get_type},
	{"g_tls_client_connection_get_use_ssl3", &G_tls_client_connection_get_use_ssl3},
	{"g_tls_client_connection_get_validation_flags", &G_tls_client_connection_get_validation_flags},
	{"g_tls_client_connection_new", &G_tls_client_connection_new},
	{"g_tls_client_connection_set_server_identity", &G_tls_client_connection_set_server_identity},
	{"g_tls_client_connection_set_use_ssl3", &G_tls_client_connection_set_use_ssl3},
	{"g_tls_client_connection_set_validation_flags", &G_tls_client_connection_set_validation_flags},
	{"g_tls_connection_emit_accept_certificate", &G_tls_connection_emit_accept_certificate},
	{"g_tls_connection_get_certificate", &G_tls_connection_get_certificate},
	{"g_tls_connection_get_peer_certificate", &G_tls_connection_get_peer_certificate},
	{"g_tls_connection_get_peer_certificate_errors", &G_tls_connection_get_peer_certificate_errors},
	{"g_tls_connection_get_rehandshake_mode", &G_tls_connection_get_rehandshake_mode},
	{"g_tls_connection_get_require_close_notify", &G_tls_connection_get_require_close_notify},
	{"g_tls_connection_get_type", &G_tls_connection_get_type},
	{"g_tls_connection_get_use_system_certdb", &G_tls_connection_get_use_system_certdb},
	{"g_tls_connection_handshake", &G_tls_connection_handshake},
	{"g_tls_connection_handshake_async", &G_tls_connection_handshake_async},
	{"g_tls_connection_handshake_finish", &G_tls_connection_handshake_finish},
	{"g_tls_connection_set_certificate", &G_tls_connection_set_certificate},
	{"g_tls_connection_set_rehandshake_mode", &G_tls_connection_set_rehandshake_mode},
	{"g_tls_connection_set_require_close_notify", &G_tls_connection_set_require_close_notify},
	{"g_tls_connection_set_use_system_certdb", &G_tls_connection_set_use_system_certdb},
	{"g_tls_error_get_type", &G_tls_error_get_type},
	{"g_tls_error_quark", &G_tls_error_quark},
	{"g_tls_rehandshake_mode_get_type", &G_tls_rehandshake_mode_get_type},
	{"g_tls_server_connection_get_type", &G_tls_server_connection_get_type},
	{"g_tls_server_connection_new", &G_tls_server_connection_new},
	{"g_unix_socket_address_type_get_type", &G_unix_socket_address_type_get_type},
	{"g_vfs_get_default", &G_vfs_get_default},
	{"g_vfs_get_file_for_path", &G_vfs_get_file_for_path},
	{"g_vfs_get_file_for_uri", &G_vfs_get_file_for_uri},
	{"g_vfs_get_local", &G_vfs_get_local},
	{"g_vfs_get_supported_uri_schemes", &G_vfs_get_supported_uri_schemes},
	{"g_vfs_get_type", &G_vfs_get_type},
	{"g_vfs_is_active", &G_vfs_is_active},
	{"g_vfs_parse_name", &G_vfs_parse_name},
	{"g_volume_can_eject", &G_volume_can_eject},
	{"g_volume_can_mount", &G_volume_can_mount},
	{"g_volume_eject", &G_volume_eject},
	{"g_volume_eject_finish", &G_volume_eject_finish},
	{"g_volume_eject_with_operation", &G_volume_eject_with_operation},
	{"g_volume_eject_with_operation_finish", &G_volume_eject_with_operation_finish},
	{"g_volume_enumerate_identifiers", &G_volume_enumerate_identifiers},
	{"g_volume_get_activation_root", &G_volume_get_activation_root},
	{"g_volume_get_drive", &G_volume_get_drive},
	{"g_volume_get_icon", &G_volume_get_icon},
	{"g_volume_get_identifier", &G_volume_get_identifier},
	{"g_volume_get_mount", &G_volume_get_mount},
	{"g_volume_get_name", &G_volume_get_name},
	{"g_volume_get_type", &G_volume_get_type},
	{"g_volume_get_uuid", &G_volume_get_uuid},
	{"g_volume_monitor_adopt_orphan_mount", &G_volume_monitor_adopt_orphan_mount},
	{"g_volume_monitor_get", &G_volume_monitor_get},
	{"g_volume_monitor_get_connected_drives", &G_volume_monitor_get_connected_drives},
	{"g_volume_monitor_get_mount_for_uuid", &G_volume_monitor_get_mount_for_uuid},
	{"g_volume_monitor_get_mounts", &G_volume_monitor_get_mounts},
	{"g_volume_monitor_get_type", &G_volume_monitor_get_type},
	{"g_volume_monitor_get_volume_for_uuid", &G_volume_monitor_get_volume_for_uuid},
	{"g_volume_monitor_get_volumes", &G_volume_monitor_get_volumes},
	{"g_volume_mount", &G_volume_mount},
	{"g_volume_mount_finish", &G_volume_mount_finish},
	{"g_volume_should_automount", &G_volume_should_automount},
	{"g_win32_input_stream_get_close_handle", &G_win32_input_stream_get_close_handle},
	{"g_win32_input_stream_get_handle", &G_win32_input_stream_get_handle},
	{"g_win32_input_stream_get_type", &G_win32_input_stream_get_type},
	{"g_win32_input_stream_new", &G_win32_input_stream_new},
	{"g_win32_input_stream_set_close_handle", &G_win32_input_stream_set_close_handle},
	{"g_win32_output_stream_get_close_handle", &G_win32_output_stream_get_close_handle},
	{"g_win32_output_stream_get_handle", &G_win32_output_stream_get_handle},
	{"g_win32_output_stream_get_type", &G_win32_output_stream_get_type},
	{"g_win32_output_stream_new", &G_win32_output_stream_new},
	{"g_win32_output_stream_set_close_handle", &G_win32_output_stream_set_close_handle},
	// Undocumented {"g_win32_resolver_get_type", &G_win32_resolver_get_type},
	{"g_zlib_compressor_format_get_type", &G_zlib_compressor_format_get_type},
	{"g_zlib_compressor_get_file_info", &G_zlib_compressor_get_file_info},
	{"g_zlib_compressor_get_type", &G_zlib_compressor_get_type},
	{"g_zlib_compressor_new", &G_zlib_compressor_new},
	{"g_zlib_compressor_set_file_info", &G_zlib_compressor_set_file_info},
	{"g_zlib_decompressor_get_file_info", &G_zlib_decompressor_get_file_info},
	{"g_zlib_decompressor_get_type", &G_zlib_decompressor_get_type},
	{"g_zlib_decompressor_new", &G_zlib_decompressor_new},
}
