package gio

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type (
	GWin32OutputStreamPrivate      struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	_                              struct{}
	GWin32InputStreamPrivate       struct{}
	GAction                        struct{}
	GActionGroup                   struct{}
	GAppInfo                       struct{}
	GApplicationCommandLinePrivate struct{}
	GApplicationPrivate            struct{}
	GAsyncInitable                 struct{}
	GBufferedInputStreamPrivate    struct{}
	GBufferedOutputStreamPrivate   struct{}
	GcharsetConverter              struct{}
	GConverter                     struct{}
	GConverterInputStreamPrivate   struct{}
	GConverterOutputStreamPrivate  struct{}
	GCredentials                   struct{}
	GDataInputStreamPrivate        struct{}
	GDataOutputStreamPrivate       struct{}
	GDBusAuthObserver              struct{}
	GDBusConnection                struct{}
	GDBusMessage                   struct{}
	GDBusMethodInvocation          struct{}
	GDBusProxyPrivate              struct{}
	GDBusServer                    struct{}
	GDrive                         struct{}
	GEmblem                        struct{}
	GEmblemedIconPrivate           struct{}
	GFileAttributeMatcher          struct{}
	GFileEnumeratorPrivate         struct{}
	GFileIcon                      struct{}
	GFileInfo                      struct{}
	GFileInputStreamPrivate        struct{}
	GFileIOStreamPrivate           struct{}
	GFileMonitorPrivate            struct{}
	GFilenameCompleter             struct{}
	GFileOutputStreamPrivate       struct{}
	GInetAddressPrivate            struct{}
	GInetSocketAddressPrivate      struct{}
	GInitable                      struct{}
	GIOExtension                   struct{}
	GIOExtensionPoint              struct{}
	GIOModule                      struct{}
	GIOSchedulerJob                struct{}
	GIOStreamPrivate               struct{}
	GLoadableIcon                  struct{}
	GMemoryInputStreamPrivate      struct{}
	GMemoryOutputStreamPrivate     struct{}
	GMount                         struct{}
	GNetworkAddressPrivate         struct{}
	GNetworkServicePrivate         struct{}
	GOutputStreamPrivate           struct{}
	GPermissionPrivate             struct{}
	GPollableInputStream           struct{}
	GPollableOutputStream          struct{}
	GProxy                         struct{}
	GProxyAddressPrivate           struct{}
	GProxyResolver                 struct{}
	GResolverPrivate               struct{}
	GSeekable                      struct{}
	GSettingsBackend               struct{}
	GSettingsPrivate               struct{}
	GSimpleActionGroupPrivate      struct{}
	GSimpleActionPrivate           struct{}
	GSimpleAsyncResult             struct{}
	GSocketClientPrivate           struct{}
	GSocketConnectable             struct{}
	GSocketConnectionPrivate       struct{}
	GSocketControlMessagePrivate   struct{}
	GSocketListenerPrivate         struct{}
	GSocketPrivate                 struct{}
	GSocketServicePrivate          struct{}
	GSrvTarget                     struct{}
	GTcpConnectionPrivate          struct{}
	GTcpWrapperConnectionPrivate   struct{}
	GThemedIcon                    struct{}
	GTlsBackend                    struct{}
	GTlsCertificatePrivate         struct{}
	GTlsClientConnection           struct{}
	GTlsConnectionPrivate          struct{}
	GUnixFDList                    struct{}
	GVolume                        struct{}
	GZlibCompressor                struct{}
	GZlibDecompressor              struct{}

	GBusAcquiredCallback func(
		connection *GDBusConnection,
		name *Gchar,
		user_data Gpointer)

	GBusNameAcquiredCallback func(
		connection *GDBusConnection,
		name *Gchar,
		user_data Gpointer)

	GBusNameLostCallback func(
		connection *GDBusConnection,
		name *Gchar,
		user_data Gpointer)

	GBusNameAppearedCallback func(
		connection *GDBusConnection,
		name *Gchar,
		name_owner *Gchar,
		user_data Gpointer)

	GBusNameVanishedCallback func(
		connection *GDBusConnection,
		name *Gchar,
		user_data Gpointer)

	GDBusInterfaceMethodCallFunc func(
		connection *GDBusConnection,
		sender *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		method_name *Gchar,
		parameters *GVariant,
		invocation *GDBusMethodInvocation,
		user_data Gpointer)

	GDBusInterfaceGetPropertyFunc func(
		connection *GDBusConnection,
		sender *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		property_name *Gchar,
		err **GError,
		user_data Gpointer)

	GDBusInterfaceSetPropertyFunc func(
		connection *GDBusConnection,
		sender *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		property_name *Gchar,
		value *GVariant,
		err **GError,
		user_data Gpointer)

	GDBusSubtreeEnumerateFunc func(
		connection *GDBusConnection,
		sender *Gchar,
		object_path *Gchar,
		user_data Gpointer) **Gchar

	GDBusSubtreeIntrospectFunc func(
		connection *GDBusConnection,
		sender *Gchar,
		object_path *Gchar,
		node *Gchar,
		user_data Gpointer) **GDBusInterfaceInfo

	GDBusSubtreeDispatchFunc func(
		connection *GDBusConnection,
		sender *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		node *Gchar,
		out_user_data *Gpointer,
		user_data Gpointer) *GDBusInterfaceVTable

	GDBusSignalCallback func(
		connection *GDBusConnection,
		sender_name *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		signal_name *Gchar,
		parameters *GVariant,
		user_data Gpointer)

	GDBusMessageFilterFunction func(
		connection *GDBusConnection,
		message *GDBusMessage,
		incoming Gboolean,
		user_data Gpointer) *GDBusMessage

	GFileProgressCallback func(
		current_num_bytes Goffset,
		total_num_bytes Goffset,
		user_data Gpointer)

	GIOSchedulerJobFunc func(
		job *GIOSchedulerJob,
		cancellable *GCancellable,
		user_data Gpointer) Gboolean

	GFileReadMoreCallback func(
		file_contents *Char,
		file_size Goffset,
		callback_data Gpointer) Gboolean

	GReallocFunc func(
		data Gpointer,
		size Gsize) Gpointer

	GSettingsBindSetMapping func(
		value *GValue,
		expected_type *GVariantType,
		user_data Gpointer) *GVariant

	GSettingsBindGetMapping func(
		value *GValue,
		variant *GVariant,
		user_data Gpointer) Gboolean

	GSettingsGetMapping func(
		value *GVariant,
		result *Gpointer,
		user_data Gpointer) Gboolean

	GSimpleAsyncThreadFunc func(
		res *GSimpleAsyncResult,
		object *GObject,
		cancellable *GCancellable)

	GApplication struct {
		parent_instance GObject
		priv            *GApplicationPrivate
	}

	GApplicationCommandLine struct {
		parent_instance GObject
		priv            *GApplicationCommandLinePrivate
	}

	GParameter struct {
		name  *Gchar
		value GValue
	}

	GBufferedInputStream struct {
		parent_instance GFilterInputStream
		priv            *GBufferedInputStreamPrivate
	}

	GFilterInputStream struct {
		parent_instance GInputStream
		base_stream     *GInputStream
	}

	GOutputStream struct {
		parent_instance GObject
		priv            *GOutputStreamPrivate
	}

	GBufferedOutputStream struct {
		parent_instance GFilterOutputStream
		priv            *GBufferedOutputStreamPrivate
	}

	GFilterOutputStream struct {
		parent_instance GOutputStream
		base_stream     *GOutputStream
	}

	GConverterInputStream struct {
		parent_instance GFilterInputStream
		priv            *GConverterInputStreamPrivate
	}

	GConverterOutputStream struct {
		parent_instance GFilterOutputStream
		priv            *GConverterOutputStreamPrivate
	}

	GDataInputStream struct {
		parent_instance GBufferedInputStream
		priv            *GDataInputStreamPrivate
	}

	GDataOutputStream struct {
		parent_instance GFilterOutputStream
		priv            *GDataOutputStreamPrivate
	}

	GIOStream struct {
		parent_instance GObject
		priv            *GIOStreamPrivate
	}

	GDBusInterfaceInfo struct {
		ref_count   Gint
		name        *Gchar
		methods     **GDBusMethodInfo
		signals     **GDBusSignalInfo
		properties  **GDBusPropertyInfo
		annotations **GDBusAnnotationInfo
	}

	GDBusMethodInfo struct {
		ref_count   Gint
		name        *Gchar
		in_args     **GDBusArgInfo
		out_args    **GDBusArgInfo
		annotations **GDBusAnnotationInfo
	}

	GDBusSignalInfo struct {
		ref_count   Gint
		name        *Gchar
		args        **GDBusArgInfo
		annotations **GDBusAnnotationInfo
	}

	GDBusPropertyInfo struct {
		ref_count   Gint
		name        *Gchar
		signature   *Gchar
		flags       GDBusPropertyInfoFlags
		annotations **GDBusAnnotationInfo
	}

	GDBusArgInfo struct {
		ref_count   Gint
		name        *Gchar
		signature   *Gchar
		annotations **GDBusAnnotationInfo
	}

	GDBusAnnotationInfo struct {
		ref_count   Gint
		key         *Gchar
		value       *Gchar
		annotations **GDBusAnnotationInfo
	}

	GDBusInterfaceVTable struct {
		method_call  GDBusInterfaceMethodCallFunc
		get_property GDBusInterfaceGetPropertyFunc
		set_property GDBusInterfaceSetPropertyFunc
		padding      [8]Gpointer
	}

	GDBusNodeInfo struct {
		ref_count   Gint
		path        *Gchar
		interfaces  **GDBusInterfaceInfo
		nodes       **GDBusNodeInfo
		annotations **GDBusAnnotationInfo
	}

	GDBusProxy struct {
		parent_instance GObject
		priv            *GDBusProxyPrivate
	}

	GDBusSubtreeVTable struct {
		enumerate  GDBusSubtreeEnumerateFunc
		introspect GDBusSubtreeIntrospectFunc
		dispatch   GDBusSubtreeDispatchFunc
		padding    [8]Gpointer
	}

	GDBusErrorEntry struct {
		error_code      Gint
		dbus_error_name *Gchar
	}

	GEmblemedIcon struct {
		parent_instance GObject
		priv            *GEmblemedIconPrivate
	}

	GFileAttributeInfoList struct {
		infos   *GFileAttributeInfo
		n_infos int
	}

	GFileAttributeInfo struct {
		name  *Char
		Type  GFileAttributeType
		flags GFileAttributeInfoFlags
	}

	GFileOutputStream struct {
		parent_instance GOutputStream
		priv            *GFileOutputStreamPrivate
	}

	GFileIOStream struct {
		parent_instance GIOStream
		priv            *GFileIOStreamPrivate
	}

	GFileEnumerator struct {
		parent_instance GObject
		priv            *GFileEnumeratorPrivate
	}

	GFileInputStream struct {
		parent_instance GInputStream
		priv            *GFileInputStreamPrivate
	}

	GFileMonitor struct {
		parent_instance GObject
		priv            *GFileMonitorPrivate
	}

	GInetAddress struct {
		parent_instance GObject
		priv            *GInetAddressPrivate
	}

	GInetSocketAddress struct {
		parent_instance GSocketAddress
		priv            *GInetSocketAddressPrivate
	}

	GSocketAddress struct {
		parent_instance GObject
	}

	GMemoryOutputStream struct {
		parent_instance GOutputStream
		priv            *GMemoryOutputStreamPrivate
	}

	GNetworkAddress struct {
		parent_instance GObject
		priv            *GNetworkAddressPrivate
	}

	GNetworkService struct {
		parent_instance GObject
		priv            *GNetworkServicePrivate
	}

	GMemoryInputStream struct {
		parent_instance GInputStream
		priv            *GMemoryInputStreamPrivate
	}

	GPermission struct {
		parent_instance GObject
		priv            *GPermissionPrivate
	}

	GProxyAddress struct {
		parent_instance GInetSocketAddress
		priv            *GProxyAddressPrivate
	}

	GResolver struct {
		parent_instance GObject
		priv            *GResolverPrivate
	}

	GSettings struct {
		parent_instance GObject
		priv            *GSettingsPrivate
	}

	GSimpleActionGroup struct {
		parent_instance GObject
		priv            *GSimpleActionGroupPrivate
	}

	GSimpleAction struct {
		parent_instance GObject
		priv            *GSimpleActionPrivate
	}

	GSocketAddressEnumerator struct {
		parent_instance GObject
	}

	GSocket struct {
		parent_instance GObject
		priv            *GSocketPrivate
	}

	GSocketClient struct {
		parent_instance GObject
		priv            *GSocketClientPrivate
	}

	GSocketConnection struct {
		parent_instance GIOStream
		priv            *GSocketConnectionPrivate
	}

	GSocketControlMessage struct {
		parent_instance GObject
		priv            *GSocketControlMessagePrivate
	}

	GSocketListener struct {
		parent_instance GObject
		priv            *GSocketListenerPrivate
	}

	GSocketService struct {
		parent_instance GSocketListener
		priv            *GSocketServicePrivate
	}

	GTlsCertificate struct {
		parent_instance GObject
		priv            *GTlsCertificatePrivate
	}

	GTcpConnection struct {
		parent_instance GSocketConnection
		priv            *GTcpConnectionPrivate
	}

	GTlsConnection struct {
		parent_instance GIOStream
		priv            *GTlsConnectionPrivate
	}

	GVfs struct {
		parent_instance GObject
	}

	GVolumeMonitor struct {
		parent_instance GObject
		priv            Gpointer
	}

	GInputVector struct {
		buffer Gpointer
		size   Gsize
	}

	GOutputVector struct {
		buffer Gconstpointer
		size   Gsize
	}

	GTcpWrapperConnection struct {
		parent_instance GTcpConnection
		priv            *GTcpWrapperConnectionPrivate
	}

	GWin32InputStream struct {
		parent_instance GInputStream
		priv            *GWin32InputStreamPrivate
	}

	GWin32OutputStream struct {
		parent_instance GOutputStream
		priv            *GWin32OutputStreamPrivate
	}
)

var (
	G_app_info_get_type func() GType

	G_app_info_create_from_commandline func(
		commandline *Char,
		application_name *Char,
		flags GAppInfoCreateFlags,
		err **GError) *GAppInfo

	G_app_info_dup func(appinfo *GAppInfo) *GAppInfo

	G_app_info_equal func(
		appinfo1 *GAppInfo, appinfo2 *GAppInfo) Gboolean

	G_app_info_get_id func(appinfo *GAppInfo) *Char

	G_app_info_get_name func(appinfo *GAppInfo) *Char

	G_app_info_get_display_name func(appinfo *GAppInfo) *Char

	G_app_info_get_description func(appinfo *GAppInfo) *Char

	G_app_info_get_executable func(appinfo *GAppInfo) *Char

	G_app_info_get_commandline func(appinfo *GAppInfo) *Char

	G_app_info_get_icon func(appinfo *GAppInfo) *GIcon

	G_app_info_launch func(
		appinfo *GAppInfo,
		files *GList,
		launch_context *GAppLaunchContext,
		err **GError) Gboolean

	G_app_info_supports_uris func(appinfo *GAppInfo) Gboolean

	G_app_info_supports_files func(appinfo *GAppInfo) Gboolean

	G_app_info_launch_uris func(
		appinfo *GAppInfo,
		uris *GList,
		launch_context *GAppLaunchContext,
		err **GError) Gboolean

	G_app_info_should_show func(appinfo *GAppInfo) Gboolean

	G_app_info_set_as_default_for_type func(
		appinfo *GAppInfo,
		content_type *Char,
		err **GError) Gboolean

	G_app_info_set_as_default_for_extension func(
		appinfo *GAppInfo,
		extension *Char,
		err **GError) Gboolean

	G_app_info_add_supports_type func(
		appinfo *GAppInfo,
		content_type *Char,
		err **GError) Gboolean

	G_app_info_can_remove_supports_type func(
		appinfo *GAppInfo) Gboolean

	G_app_info_remove_supports_type func(
		appinfo *GAppInfo,
		content_type *Char,
		err **GError) Gboolean

	G_app_info_can_delete func(appinfo *GAppInfo) Gboolean

	G_app_info_delete func(appinfo *GAppInfo) Gboolean

	G_app_info_set_as_last_used_for_type func(
		appinfo *GAppInfo,
		content_type *Char,
		err **GError) Gboolean

	G_app_info_get_all func() *GList

	G_app_info_get_all_for_type func(content_type *Char) *GList

	G_app_info_get_recommended_for_type func(
		content_type *Gchar) *GList

	G_app_info_get_fallback_for_type func(
		content_type *Gchar) *GList

	G_app_info_reset_type_associations func(content_type *Char)

	G_app_info_get_default_for_type func(
		content_type *Char,
		must_support_uris Gboolean) *GAppInfo

	G_app_info_get_default_for_uri_scheme func(
		uri_scheme *Char) *GAppInfo

	G_app_info_launch_default_for_uri func(
		uri *Char,
		launch_context *GAppLaunchContext,
		err **GError) Gboolean

	G_app_launch_context_get_type func() GType

	G_app_launch_context_new func() *GAppLaunchContext

	G_app_launch_context_get_display func(
		context *GAppLaunchContext,
		info *GAppInfo,
		files *GList) *Char

	G_app_launch_context_get_startup_notify_id func(
		context *GAppLaunchContext,
		info *GAppInfo,
		files *GList) *Char

	G_app_launch_context_launch_failed func(
		context *GAppLaunchContext,
		startup_notify_id *Char)

	G_action_get_type func() GType

	G_action_get_name func(action *GAction) *Gchar

	G_action_get_parameter_type func(action *GAction) *GVariantType

	G_action_get_state_type func(action *GAction) *GVariantType

	G_action_get_state_hint func(action *GAction) *GVariant

	G_action_get_enabled func(action *GAction) Gboolean

	G_action_get_state func(action *GAction) *GVariant

	G_action_set_state func(action *GAction, value *GVariant)

	G_action_activate func(action *GAction, parameter *GVariant)

	G_simple_action_get_type func() GType

	G_simple_action_new func(
		name *Gchar, parameter_type *GVariantType) *GSimpleAction

	G_simple_action_new_stateful func(
		name *Gchar,
		parameter_type *GVariantType,
		state *GVariant) *GSimpleAction

	G_simple_action_set_enabled func(
		simple *GSimpleAction,
		enabled Gboolean)

	G_action_group_get_type func() GType

	G_action_group_has_action func(
		action_group *GActionGroup,
		action_name *Gchar) Gboolean

	G_action_group_list_actions func(
		action_group *GActionGroup) **Gchar

	G_action_group_get_action_parameter_type func(
		action_group *GActionGroup,
		action_name *Gchar) *GVariantType

	G_action_group_get_action_state_type func(
		action_group *GActionGroup,
		action_name *Gchar) *GVariantType

	G_action_group_get_action_state_hint func(
		action_group *GActionGroup,
		action_name *Gchar) *GVariant

	G_action_group_get_action_enabled func(
		action_group *GActionGroup,
		action_name *Gchar) Gboolean

	G_action_group_get_action_state func(
		action_group *GActionGroup,
		action_name *Gchar) *GVariant

	G_action_group_change_action_state func(
		action_group *GActionGroup,
		action_name *Gchar,
		value *GVariant)

	G_action_group_activate_action func(
		action_group *GActionGroup,
		action_name *Gchar,
		parameter *GVariant)

	G_action_group_action_added func(
		action_group *GActionGroup,
		action_name *Gchar)

	G_action_group_action_removed func(
		action_group *GActionGroup,
		action_name *Gchar)

	G_action_group_action_enabled_changed func(
		action_group *GActionGroup,
		action_name *Gchar,
		enabled Gboolean)

	G_action_group_action_state_changed func(
		action_group *GActionGroup,
		action_name *Gchar,
		state *GVariant)

	G_simple_action_group_get_type func() GType

	G_simple_action_group_new func() *GSimpleActionGroup

	G_simple_action_group_lookup func(
		simple *GSimpleActionGroup,
		action_name *Gchar) *GAction

	G_simple_action_group_insert func(
		simple *GSimpleActionGroup,
		action *GAction)

	G_simple_action_group_remove func(
		simple *GSimpleActionGroup,
		action_name *Gchar)

	G_application_get_type func() GType

	G_application_id_is_valid func(
		application_id *Gchar) Gboolean

	G_application_new func(
		application_id *Gchar,
		flags GApplicationFlags) *GApplication

	G_application_get_application_id func(
		application *GApplication) *Gchar

	G_application_set_application_id func(
		application *GApplication,
		application_id *Gchar)

	G_application_get_inactivity_timeout func(
		application *GApplication) Guint

	G_application_set_inactivity_timeout func(
		application *GApplication,
		inactivity_timeout Guint)

	G_application_get_flags func(
		application *GApplication) GApplicationFlags

	G_application_set_flags func(
		application *GApplication,
		flags GApplicationFlags)

	G_application_set_action_group func(
		application *GApplication,
		action_group *GActionGroup)

	G_application_get_is_registered func(
		application *GApplication) Gboolean

	G_application_get_is_remote func(
		application *GApplication) Gboolean

	G_application_register func(
		application *GApplication,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_application_hold func(
		application *GApplication)

	G_application_release func(
		application *GApplication)

	G_application_activate func(
		application *GApplication)

	G_application_open func(
		application *GApplication,
		files **GFile,
		n_files Gint,
		hint *Gchar)

	G_application_run func(
		application *GApplication,
		argc int,
		argv **Char) int

	G_application_command_line_get_type func() GType

	G_application_command_line_get_arguments func(
		cmdline *GApplicationCommandLine,
		argc *int) **Gchar

	G_application_command_line_get_environ func(
		cmdline *GApplicationCommandLine) **Gchar

	G_application_command_line_getenv func(
		cmdline *GApplicationCommandLine,
		name *Gchar) *Gchar

	G_application_command_line_get_cwd func(
		cmdline *GApplicationCommandLine) *Gchar

	G_application_command_line_get_is_remote func(
		cmdline *GApplicationCommandLine) Gboolean

	//TODO(t): Variant
	//g_application_command_line_print func(cmdline  *GApplicationCommandLine, format  *Gchar, ...)
	//g_application_command_line_printerr func(cmdline  *GApplicationCommandLine, format  *Gchar, ...)

	G_application_command_line_get_exit_status func(
		cmdline *GApplicationCommandLine) int

	G_application_command_line_set_exit_status func(
		cmdline *GApplicationCommandLine,
		exit_status int)

	G_application_command_line_get_platform_data func(
		cmdline *GApplicationCommandLine) *GVariant

	G_initable_get_type func() GType

	G_initable_init func(
		initable *GInitable,
		cancellable *GCancellable,
		err **GError) Gboolean

	//TODO(t):Variant
	//g_initable_new func(object_type  GType, cancellable  *GCancellable, error  **GError, first_property_name  *Gchar, ...) Gpointer

	G_initable_newv func(
		object_type GType,
		n_parameters Guint,
		parameters *GParameter,
		cancellable *GCancellable,
		err **GError) Gpointer

	G_initable_new_valist func(
		object_type GType,
		first_property_name *Gchar,
		var_args Va_list,
		cancellable *GCancellable,
		err **GError) *GObject

	G_async_initable_get_type func() GType

	G_async_initable_init_async func(
		initable *GAsyncInitable,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_async_initable_init_finish func(
		initable *GAsyncInitable,
		res *GAsyncResult,
		err **GError) Gboolean

	//TODO(t):Variant
	//g_async_initable_new_async func(object_type  GType, io_priority  int, cancellable  *GCancellable, callback  GAsyncReadyCallback, user_data  Gpointer, first_property_name  *Gchar, ...)

	G_async_initable_newv_async func(
		object_type GType,
		n_parameters Guint,
		parameters *GParameter,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_async_initable_new_valist_async func(
		object_type GType,
		first_property_name *Gchar,
		var_args Va_list,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_async_initable_new_finish func(
		initable *GAsyncInitable,
		res *GAsyncResult,
		err **GError) *GObject

	G_async_result_get_type func() GType

	G_async_result_get_user_data func(
		res *GAsyncResult) Gpointer

	G_async_result_get_source_object func(
		res *GAsyncResult) *GObject

	G_input_stream_get_type func() GType

	G_input_stream_read func(
		stream *GInputStream,
		buffer *Void,
		count Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_input_stream_read_all func(
		stream *GInputStream,
		buffer *Void,
		count Gsize,
		bytes_read *Gsize,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_input_stream_skip func(
		stream *GInputStream,
		count Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_input_stream_close func(
		stream *GInputStream,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_input_stream_read_async func(
		stream *GInputStream,
		buffer *Void,
		count Gsize,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_input_stream_read_finish func(
		stream *GInputStream,
		result *GAsyncResult,
		err **GError) Gssize

	G_input_stream_skip_async func(
		stream *GInputStream,
		count Gsize,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_input_stream_skip_finish func(
		stream *GInputStream,
		result *GAsyncResult,
		err **GError) Gssize

	G_input_stream_close_async func(
		stream *GInputStream,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_input_stream_close_finish func(
		stream *GInputStream,
		result *GAsyncResult,
		err **GError) Gboolean

	G_input_stream_is_closed func(
		stream *GInputStream) Gboolean

	G_input_stream_has_pending func(
		stream *GInputStream) Gboolean

	G_input_stream_set_pending func(
		stream *GInputStream,
		err **GError) Gboolean

	G_input_stream_clear_pending func(
		stream *GInputStream)

	G_filter_input_stream_get_type func() GType

	G_filter_input_stream_get_base_stream func(
		stream *GFilterInputStream) *GInputStream

	G_filter_input_stream_get_close_base_stream func(
		stream *GFilterInputStream) Gboolean

	G_filter_input_stream_set_close_base_stream func(
		stream *GFilterInputStream,
		close_base Gboolean)

	G_buffered_input_stream_get_type func() GType

	G_buffered_input_stream_new func(
		base_stream *GInputStream) *GInputStream

	G_buffered_input_stream_new_sized func(
		base_stream *GInputStream,
		size Gsize) *GInputStream

	G_buffered_input_stream_get_buffer_size func(
		stream *GBufferedInputStream) Gsize

	G_buffered_input_stream_set_buffer_size func(
		stream *GBufferedInputStream,
		size Gsize)

	G_buffered_input_stream_get_available func(
		stream *GBufferedInputStream) Gsize

	G_buffered_input_stream_peek func(
		stream *GBufferedInputStream,
		buffer *Void,
		offset Gsize,
		count Gsize) Gsize

	G_buffered_input_stream_peek_buffer func(
		stream *GBufferedInputStream,
		count *Gsize) *Void

	G_buffered_input_stream_fill func(
		stream *GBufferedInputStream,
		count Gssize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_buffered_input_stream_fill_async func(
		stream *GBufferedInputStream,
		count Gssize,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_buffered_input_stream_fill_finish func(
		stream *GBufferedInputStream,
		result *GAsyncResult,
		err **GError) Gssize

	G_buffered_input_stream_read_byte func(
		stream *GBufferedInputStream,
		cancellable *GCancellable,
		err **GError) int

	G_output_stream_get_type func() GType

	G_output_stream_write func(
		stream *GOutputStream,
		buffer *Void,
		count Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_output_stream_write_all func(
		stream *GOutputStream,
		buffer *Void,
		count Gsize,
		bytes_written *Gsize,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_output_stream_splice func(
		stream *GOutputStream,
		source *GInputStream,
		flags GOutputStreamSpliceFlags,
		cancellable *GCancellable,
		err **GError) Gssize

	G_output_stream_flush func(
		stream *GOutputStream,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_output_stream_close func(
		stream *GOutputStream,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_output_stream_write_async func(
		stream *GOutputStream,
		buffer *Void,
		count Gsize,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_output_stream_write_finish func(
		stream *GOutputStream,
		result *GAsyncResult,
		err **GError) Gssize

	G_output_stream_splice_async func(
		stream *GOutputStream,
		source *GInputStream,
		flags GOutputStreamSpliceFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_output_stream_splice_finish func(
		stream *GOutputStream,
		result *GAsyncResult,
		err **GError) Gssize

	G_output_stream_flush_async func(
		stream *GOutputStream,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_output_stream_flush_finish func(
		stream *GOutputStream,
		result *GAsyncResult,
		err **GError) Gboolean

	G_output_stream_close_async func(
		stream *GOutputStream,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_output_stream_close_finish func(
		stream *GOutputStream,
		result *GAsyncResult,
		err **GError) Gboolean

	G_output_stream_is_closed func(
		stream *GOutputStream) Gboolean

	G_output_stream_is_closing func(
		stream *GOutputStream) Gboolean

	G_output_stream_has_pending func(
		stream *GOutputStream) Gboolean

	G_output_stream_set_pending func(
		stream *GOutputStream,
		err **GError) Gboolean

	G_output_stream_clear_pending func(
		stream *GOutputStream)

	G_filter_output_stream_get_type func() GType

	G_filter_output_stream_get_base_stream func(
		stream *GFilterOutputStream) *GOutputStream

	G_filter_output_stream_get_close_base_stream func(
		stream *GFilterOutputStream) Gboolean

	G_filter_output_stream_set_close_base_stream func(
		stream *GFilterOutputStream,
		close_base Gboolean)

	G_buffered_output_stream_get_type func() GType

	G_buffered_output_stream_new func(
		base_stream *GOutputStream) *GOutputStream

	G_buffered_output_stream_new_sized func(
		base_stream *GOutputStream,
		size Gsize) *GOutputStream

	G_buffered_output_stream_get_buffer_size func(
		stream *GBufferedOutputStream) Gsize

	G_buffered_output_stream_set_buffer_size func(
		stream *GBufferedOutputStream,
		size Gsize)

	G_buffered_output_stream_get_auto_grow func(
		stream *GBufferedOutputStream) Gboolean

	G_buffered_output_stream_set_auto_grow func(
		stream *GBufferedOutputStream,
		auto_grow Gboolean)

	G_cancellable_get_type func() GType

	G_cancellable_new func() *GCancellable

	G_cancellable_is_cancelled func(
		cancellable *GCancellable) Gboolean

	G_cancellable_set_error_if_cancelled func(
		cancellable *GCancellable,
		err **GError) Gboolean

	G_cancellable_get_fd func(
		cancellable *GCancellable) int

	G_cancellable_make_pollfd func(
		cancellable *GCancellable,
		pollfd *GPollFD) Gboolean

	G_cancellable_release_fd func(
		cancellable *GCancellable)

	G_cancellable_source_new func(
		cancellable *GCancellable) *GSource

	G_cancellable_get_current func() *GCancellable

	G_cancellable_push_current func(
		cancellable *GCancellable)

	G_cancellable_pop_current func(
		cancellable *GCancellable)

	G_cancellable_reset func(
		cancellable *GCancellable)

	G_cancellable_connect func(
		cancellable *GCancellable,
		callback GCallback,
		data Gpointer,
		data_destroy_func GDestroyNotify) Gulong

	G_cancellable_disconnect func(
		cancellable *GCancellable,
		handler_id Gulong)

	G_cancellable_cancel func(
		cancellable *GCancellable)

	G_converter_get_type func() GType

	G_converter_convert func(
		converter *GConverter,
		inbuf *Void,
		inbuf_size Gsize,
		outbuf *Void,
		outbuf_size Gsize,
		flags GConverterFlags,
		bytes_read *Gsize,
		bytes_written *Gsize,
		err **GError) GConverterResult

	G_converter_reset func(
		converter *GConverter)

	G_charset_converter_get_type func() GType

	G_charset_converter_new func(
		to_charset *Gchar,
		from_charset *Gchar,
		err **GError) *GcharsetConverter

	G_charset_converter_set_use_fallback func(
		converter *GcharsetConverter,
		use_fallback Gboolean)

	G_charset_converter_get_use_fallback func(
		converter *GcharsetConverter) Gboolean

	G_charset_converter_get_num_fallbacks func(
		converter *GcharsetConverter) Guint

	G_content_type_equals func(
		type1 *Gchar,
		type2 *Gchar) Gboolean

	G_content_type_is_a func(
		typ *Gchar,
		supertype *Gchar) Gboolean

	G_content_type_is_unknown func(
		typ *Gchar) Gboolean

	G_content_type_get_description func(
		typ *Gchar) *Gchar

	G_content_type_get_mime_type func(
		typ *Gchar) *Gchar

	G_content_type_get_icon func(
		typ *Gchar) *GIcon

	G_content_type_can_be_executable func(
		typ *Gchar) Gboolean

	G_content_type_from_mime_type func(
		mime_type *Gchar) *Gchar

	G_content_type_guess func(
		filename *Gchar,
		data *Guchar,
		data_size Gsize,
		result_uncertain *Gboolean) *Gchar

	G_content_type_guess_for_tree func(
		root *GFile) **Gchar

	G_content_types_get_registered func() *GList

	G_converter_input_stream_get_type func() GType

	G_converter_input_stream_new func(
		base_stream *GInputStream,
		converter *GConverter) *GInputStream

	G_converter_input_stream_get_converter func(
		converter_stream *GConverterInputStream) *GConverter

	G_converter_output_stream_get_type func() GType

	G_converter_output_stream_new func(
		base_stream *GOutputStream,
		converter *GConverter) *GOutputStream

	G_converter_output_stream_get_converter func(
		converter_stream *GConverterOutputStream) *GConverter

	G_credentials_get_type func() GType

	G_credentials_new func() *GCredentials

	G_credentials_to_string func(
		credentials *GCredentials) *Gchar

	G_credentials_get_native func(
		credentials *GCredentials,
		native_type GCredentialsType) Gpointer

	G_credentials_set_native func(
		credentials *GCredentials,
		native_type GCredentialsType,
		native Gpointer)

	G_credentials_is_same_user func(
		credentials *GCredentials,
		other_credentials *GCredentials,
		err **GError) Gboolean

	G_data_input_stream_get_type func() GType

	G_data_input_stream_new func(
		base_stream *GInputStream) *GDataInputStream

	G_data_input_stream_set_byte_order func(
		stream *GDataInputStream,
		order GDataStreamByteOrder)

	G_data_input_stream_get_byte_order func(
		stream *GDataInputStream) GDataStreamByteOrder

	G_data_input_stream_set_newline_type func(
		stream *GDataInputStream,
		typ GDataStreamNewlineType)

	G_data_input_stream_get_newline_type func(
		stream *GDataInputStream) GDataStreamNewlineType

	G_data_input_stream_read_byte func(
		stream *GDataInputStream,
		cancellable *GCancellable,
		err **GError) Guchar

	G_data_input_stream_read_int16 func(
		stream *GDataInputStream,
		cancellable *GCancellable,
		err **GError) Gint16

	G_data_input_stream_read_uint16 func(
		stream *GDataInputStream,
		cancellable *GCancellable,
		err **GError) Guint16

	G_data_input_stream_read_int32 func(
		stream *GDataInputStream,
		cancellable *GCancellable,
		err **GError) Gint32

	G_data_input_stream_read_uint32 func(
		stream *GDataInputStream,
		cancellable *GCancellable,
		err **GError) Guint32

	G_data_input_stream_read_int64 func(
		stream *GDataInputStream,
		cancellable *GCancellable,
		err **GError) Gint64

	G_data_input_stream_read_uint64 func(
		stream *GDataInputStream,
		cancellable *GCancellable,
		err **GError) Guint64

	G_data_input_stream_read_line func(
		stream *GDataInputStream,
		length *Gsize,
		cancellable *GCancellable,
		err **GError) *Char

	G_data_input_stream_read_line_async func(
		stream *GDataInputStream,
		io_priority Gint,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_data_input_stream_read_line_finish func(
		stream *GDataInputStream,
		result *GAsyncResult,
		length *Gsize,
		err **GError) *Char

	G_data_input_stream_read_until func(
		stream *GDataInputStream,
		stop_chars *Gchar,
		length *Gsize,
		cancellable *GCancellable,
		err **GError) *Char

	G_data_input_stream_read_until_async func(
		stream *GDataInputStream,
		stop_chars *Gchar,
		io_priority Gint,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_data_input_stream_read_until_finish func(
		stream *GDataInputStream,
		result *GAsyncResult,
		length *Gsize,
		err **GError) *Char

	G_data_input_stream_read_upto func(
		stream *GDataInputStream,
		stop_chars *Gchar,
		stop_chars_len Gssize,
		length *Gsize,
		cancellable *GCancellable,
		err **GError) *Char

	G_data_input_stream_read_upto_async func(
		stream *GDataInputStream,
		stop_chars *Gchar,
		stop_chars_len Gssize,
		io_priority Gint,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_data_input_stream_read_upto_finish func(
		stream *GDataInputStream,
		result *GAsyncResult,
		length *Gsize,
		err **GError) *Char

	G_data_output_stream_get_type func() GType

	G_data_output_stream_new func(
		base_stream *GOutputStream) *GDataOutputStream

	G_data_output_stream_set_byte_order func(
		stream *GDataOutputStream,
		order GDataStreamByteOrder)

	G_data_output_stream_get_byte_order func(
		stream *GDataOutputStream) GDataStreamByteOrder

	G_data_output_stream_put_byte func(
		stream *GDataOutputStream,
		data Guchar,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_data_output_stream_put_int16 func(
		stream *GDataOutputStream,
		data Gint16,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_data_output_stream_put_uint16 func(
		stream *GDataOutputStream,
		data Guint16,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_data_output_stream_put_int32 func(
		stream *GDataOutputStream,
		data Gint32,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_data_output_stream_put_uint32 func(
		stream *GDataOutputStream,
		data Guint32,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_data_output_stream_put_int64 func(
		stream *GDataOutputStream,
		data Gint64,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_data_output_stream_put_uint64 func(
		stream *GDataOutputStream,
		data Guint64,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_data_output_stream_put_string func(
		stream *GDataOutputStream,
		str *Char,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_dbus_is_address func(
		string *Gchar) Gboolean

	G_dbus_is_supported_address func(
		string *Gchar,
		err **GError) Gboolean

	G_dbus_address_get_stream func(
		address *Gchar,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_address_get_stream_finish func(
		res *GAsyncResult,
		out_guid **Gchar,
		err **GError) *GIOStream

	G_dbus_address_get_stream_sync func(
		address *Gchar,
		out_guid **Gchar,
		cancellable *GCancellable,
		err **GError) *GIOStream

	G_dbus_address_get_for_bus_sync func(
		bus_type GBusType,
		cancellable *GCancellable,
		err **GError) *Gchar

	G_dbus_auth_observer_get_type func() GType

	G_dbus_auth_observer_new func() *GDBusAuthObserver

	G_dbus_auth_observer_authorize_authenticated_peer func(
		observer *GDBusAuthObserver,
		stream *GIOStream,
		credentials *GCredentials) Gboolean

	G_dbus_connection_get_type func() GType

	G_bus_get func(
		bus_type GBusType,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_bus_get_finish func(
		res *GAsyncResult,
		err **GError) *GDBusConnection

	G_bus_get_sync func(
		bus_type GBusType,
		cancellable *GCancellable,
		err **GError) *GDBusConnection

	G_dbus_connection_new func(
		stream *GIOStream,
		guid *Gchar,
		flags GDBusConnectionFlags,
		observer *GDBusAuthObserver,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_connection_new_finish func(
		res *GAsyncResult,
		err **GError) *GDBusConnection

	G_dbus_connection_new_sync func(
		stream *GIOStream,
		guid *Gchar,
		flags GDBusConnectionFlags,
		observer *GDBusAuthObserver,
		cancellable *GCancellable,
		err **GError) *GDBusConnection

	G_dbus_connection_new_for_address func(
		address *Gchar,
		flags GDBusConnectionFlags,
		observer *GDBusAuthObserver,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_connection_new_for_address_finish func(
		res *GAsyncResult,
		err **GError) *GDBusConnection

	G_dbus_connection_new_for_address_sync func(
		address *Gchar,
		flags GDBusConnectionFlags,
		observer *GDBusAuthObserver,
		cancellable *GCancellable,
		err **GError) *GDBusConnection

	G_dbus_connection_start_message_processing func(
		connection *GDBusConnection)

	G_dbus_connection_is_closed func(
		connection *GDBusConnection) Gboolean

	G_dbus_connection_get_stream func(
		connection *GDBusConnection) *GIOStream

	G_dbus_connection_get_guid func(
		connection *GDBusConnection) *Gchar

	G_dbus_connection_get_unique_name func(
		connection *GDBusConnection) *Gchar

	G_dbus_connection_get_peer_credentials func(
		connection *GDBusConnection) *GCredentials

	G_dbus_connection_get_exit_on_close func(
		connection *GDBusConnection) Gboolean

	G_dbus_connection_set_exit_on_close func(
		connection *GDBusConnection,
		exit_on_close Gboolean)

	G_dbus_connection_get_capabilities func(
		connection *GDBusConnection) GDBusCapabilityFlags

	G_dbus_connection_close func(
		connection *GDBusConnection,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_connection_close_finish func(
		connection *GDBusConnection,
		res *GAsyncResult,
		err **GError) Gboolean

	G_dbus_connection_close_sync func(
		connection *GDBusConnection,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_dbus_connection_flush func(
		connection *GDBusConnection,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_connection_flush_finish func(
		connection *GDBusConnection,
		res *GAsyncResult,
		err **GError) Gboolean

	G_dbus_connection_flush_sync func(
		connection *GDBusConnection,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_dbus_connection_send_message func(
		connection *GDBusConnection,
		message *GDBusMessage,
		flags GDBusSendMessageFlags,
		out_serial *Guint32,
		err **GError) Gboolean

	G_dbus_connection_send_message_with_reply func(
		connection *GDBusConnection,
		message *GDBusMessage,
		flags GDBusSendMessageFlags,
		timeout_msec Gint,
		out_serial *Guint32,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_connection_send_message_with_reply_finish func(
		connection *GDBusConnection,
		res *GAsyncResult,
		err **GError) *GDBusMessage

	G_dbus_connection_send_message_with_reply_sync func(
		connection *GDBusConnection,
		message *GDBusMessage,
		flags GDBusSendMessageFlags,
		timeout_msec Gint,
		out_serial *Guint32,
		cancellable *GCancellable,
		err **GError) *GDBusMessage

	G_dbus_connection_emit_signal func(
		connection *GDBusConnection,
		destination_bus_name *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		signal_name *Gchar,
		parameters *GVariant,
		err **GError) Gboolean

	G_dbus_connection_call func(
		connection *GDBusConnection,
		bus_name *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		method_name *Gchar,
		parameters *GVariant,
		reply_type *GVariantType,
		flags GDBusCallFlags,
		timeout_msec Gint,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_connection_call_finish func(
		connection *GDBusConnection,
		res *GAsyncResult,
		err **GError) *GVariant

	G_dbus_connection_call_sync func(
		connection *GDBusConnection,
		bus_name *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		method_name *Gchar,
		parameters *GVariant,
		reply_type *GVariantType,
		flags GDBusCallFlags,
		timeout_msec Gint,
		cancellable *GCancellable,
		err **GError) *GVariant

	G_dbus_connection_register_object func(
		connection *GDBusConnection,
		object_path *Gchar,
		interface_info *GDBusInterfaceInfo,
		vtable *GDBusInterfaceVTable,
		user_data Gpointer,
		user_data_free_func GDestroyNotify,
		err **GError) Guint

	G_dbus_connection_unregister_object func(
		connection *GDBusConnection,
		registration_id Guint) Gboolean

	G_dbus_connection_register_subtree func(
		connection *GDBusConnection,
		object_path *Gchar,
		vtable *GDBusSubtreeVTable,
		flags GDBusSubtreeFlags,
		user_data Gpointer,
		user_data_free_func GDestroyNotify,
		err **GError) Guint

	G_dbus_connection_unregister_subtree func(
		connection *GDBusConnection,
		registration_id Guint) Gboolean

	G_dbus_connection_signal_subscribe func(
		connection *GDBusConnection,
		sender *Gchar,
		interface_name *Gchar,
		member *Gchar,
		object_path *Gchar,
		arg0 *Gchar,
		flags GDBusSignalFlags,
		callback GDBusSignalCallback,
		user_data Gpointer,
		user_data_free_func GDestroyNotify) Guint

	G_dbus_connection_signal_unsubscribe func(
		connection *GDBusConnection,
		subscription_id Guint)

	G_dbus_connection_add_filter func(
		connection *GDBusConnection,
		filter_function GDBusMessageFilterFunction,
		user_data Gpointer,
		user_data_free_func GDestroyNotify) Guint

	G_dbus_connection_remove_filter func(
		connection *GDBusConnection,
		filter_id Guint)

	G_dbus_error_quark func() GQuark

	G_dbus_error_is_remote_error func(
		err *GError) Gboolean

	G_dbus_error_get_remote_error func(
		err *GError) *Gchar

	G_dbus_error_strip_remote_error func(
		err *GError) Gboolean

	G_dbus_error_register_error func(
		error_domain GQuark,
		error_code Gint,
		dbus_error_name *Gchar) Gboolean

	G_dbus_error_unregister_error func(
		error_domain GQuark,
		error_code Gint,
		dbus_error_name *Gchar) Gboolean

	G_dbus_error_register_error_domain func(
		error_domain_quark_name *Gchar,
		quark_volatile *Gsize,
		entries *GDBusErrorEntry,
		num_entries Guint)

	G_dbus_error_new_for_dbus_error func(
		dbus_error_name *Gchar,
		dbus_error_message *Gchar) *GError

	//g_dbus_error_set_dbus_error func(error  **GError, dbus_error_name  *Gchar, dbus_error_message  *Gchar, format  *Gchar, ...)

	G_dbus_error_set_dbus_error_valist func(
		err **GError,
		dbus_error_name *Gchar,
		dbus_error_message *Gchar,
		format *Gchar,
		var_args Va_list)

	G_dbus_error_encode_gerror func(
		err *GError) *Gchar

	G_dbus_annotation_info_lookup func(
		annotations **GDBusAnnotationInfo,
		name *Gchar) *Gchar

	G_dbus_interface_info_lookup_method func(
		info *GDBusInterfaceInfo,
		name *Gchar) *GDBusMethodInfo

	G_dbus_interface_info_lookup_signal func(
		info *GDBusInterfaceInfo,
		name *Gchar) *GDBusSignalInfo

	G_dbus_interface_info_lookup_property func(
		info *GDBusInterfaceInfo,
		name *Gchar) *GDBusPropertyInfo

	G_dbus_interface_info_generate_xml func(
		info *GDBusInterfaceInfo,
		indent Guint,
		string_builder *GString)

	G_dbus_node_info_new_for_xml func(
		xml_data *Gchar,
		err **GError) *GDBusNodeInfo

	G_dbus_node_info_lookup_interface func(
		info *GDBusNodeInfo,
		name *Gchar) *GDBusInterfaceInfo

	G_dbus_node_info_generate_xml func(
		info *GDBusNodeInfo,
		indent Guint,
		string_builder *GString)

	G_dbus_node_info_ref func(
		info *GDBusNodeInfo) *GDBusNodeInfo

	G_dbus_interface_info_ref func(
		info *GDBusInterfaceInfo) *GDBusInterfaceInfo

	G_dbus_method_info_ref func(
		info *GDBusMethodInfo) *GDBusMethodInfo

	G_dbus_signal_info_ref func(
		info *GDBusSignalInfo) *GDBusSignalInfo

	G_dbus_property_info_ref func(
		info *GDBusPropertyInfo) *GDBusPropertyInfo

	G_dbus_arg_info_ref func(
		info *GDBusArgInfo) *GDBusArgInfo

	G_dbus_annotation_info_ref func(
		info *GDBusAnnotationInfo) *GDBusAnnotationInfo

	G_dbus_node_info_unref func(
		info *GDBusNodeInfo)

	G_dbus_interface_info_unref func(
		info *GDBusInterfaceInfo)

	G_dbus_method_info_unref func(
		info *GDBusMethodInfo)

	G_dbus_signal_info_unref func(
		info *GDBusSignalInfo)

	G_dbus_property_info_unref func(
		info *GDBusPropertyInfo)

	G_dbus_arg_info_unref func(
		info *GDBusArgInfo)

	G_dbus_annotation_info_unref func(
		info *GDBusAnnotationInfo)

	G_dbus_node_info_get_type func() GType

	G_dbus_interface_info_get_type func() GType

	G_dbus_method_info_get_type func() GType

	G_dbus_signal_info_get_type func() GType

	G_dbus_property_info_get_type func() GType

	G_dbus_arg_info_get_type func() GType

	G_dbus_annotation_info_get_type func() GType

	G_dbus_message_get_type func() GType

	G_dbus_message_new func() *GDBusMessage

	G_dbus_message_new_signal func(
		path *Gchar,
		interface_ *Gchar,
		signal *Gchar) *GDBusMessage

	G_dbus_message_new_method_call func(
		name *Gchar,
		path *Gchar,
		interface_ *Gchar,
		method *Gchar) *GDBusMessage

	G_dbus_message_new_method_reply func(
		method_call_message *GDBusMessage) *GDBusMessage

	//g_dbus_message_new_method_error func(method_call_message  *GDBusMessage, error_name  *Gchar, error_message_format  *Gchar, ...) *GDBusMessage

	G_dbus_message_new_method_error_valist func(
		method_call_message *GDBusMessage,
		error_name *Gchar,
		error_message_format *Gchar,
		var_args Va_list) *GDBusMessage

	G_dbus_message_new_method_error_literal func(
		method_call_message *GDBusMessage,
		error_name *Gchar,
		error_message *Gchar) *GDBusMessage

	G_dbus_message_print func(
		message *GDBusMessage,
		indent Guint) *Gchar

	G_dbus_message_get_locked func(
		message *GDBusMessage) Gboolean

	G_dbus_message_lock func(
		message *GDBusMessage)

	G_dbus_message_copy func(
		message *GDBusMessage,
		err **GError) *GDBusMessage

	G_dbus_message_get_byte_order func(
		message *GDBusMessage) GDBusMessageByteOrder

	G_dbus_message_set_byte_order func(
		message *GDBusMessage,
		byte_order GDBusMessageByteOrder)

	G_dbus_message_get_message_type func(
		message *GDBusMessage) GDBusMessageType

	G_dbus_message_set_message_type func(
		message *GDBusMessage,
		typ GDBusMessageType)

	G_dbus_message_get_flags func(
		message *GDBusMessage) GDBusMessageFlags

	G_dbus_message_set_flags func(
		message *GDBusMessage,
		flags GDBusMessageFlags)

	G_dbus_message_get_serial func(
		message *GDBusMessage) Guint32

	G_dbus_message_set_serial func(
		message *GDBusMessage,
		serial Guint32)

	G_dbus_message_get_header func(
		message *GDBusMessage,
		header_field GDBusMessageHeaderField) *GVariant

	G_dbus_message_set_header func(
		message *GDBusMessage,
		header_field GDBusMessageHeaderField,
		value *GVariant)

	G_dbus_message_get_header_fields func(
		message *GDBusMessage) *Guchar

	G_dbus_message_get_body func(
		message *GDBusMessage) *GVariant

	G_dbus_message_set_body func(
		message *GDBusMessage,
		body *GVariant)

	G_dbus_message_get_unix_fd_list func(
		message *GDBusMessage) *GUnixFDList

	G_dbus_message_set_unix_fd_list func(
		message *GDBusMessage,
		fd_list *GUnixFDList)

	G_dbus_message_get_reply_serial func(
		message *GDBusMessage) Guint32

	G_dbus_message_set_reply_serial func(
		message *GDBusMessage,
		value Guint32)

	G_dbus_message_get_interface func(
		message *GDBusMessage) *Gchar

	G_dbus_message_set_interface func(
		message *GDBusMessage,
		value *Gchar)

	G_dbus_message_get_member func(
		message *GDBusMessage) *Gchar

	G_dbus_message_set_member func(
		message *GDBusMessage,
		value *Gchar)

	G_dbus_message_get_path func(
		message *GDBusMessage) *Gchar

	G_dbus_message_set_path func(
		message *GDBusMessage,
		value *Gchar)

	G_dbus_message_get_sender func(
		message *GDBusMessage) *Gchar

	G_dbus_message_set_sender func(
		message *GDBusMessage,
		value *Gchar)

	G_dbus_message_get_destination func(
		message *GDBusMessage) *Gchar

	G_dbus_message_set_destination func(
		message *GDBusMessage,
		value *Gchar)

	G_dbus_message_get_error_name func(
		message *GDBusMessage) *Gchar

	G_dbus_message_set_error_name func(
		message *GDBusMessage,
		value *Gchar)

	G_dbus_message_get_signature func(
		message *GDBusMessage) *Gchar

	G_dbus_message_set_signature func(
		message *GDBusMessage,
		value *Gchar)

	G_dbus_message_get_num_unix_fds func(
		message *GDBusMessage) Guint32

	G_dbus_message_set_num_unix_fds func(
		message *GDBusMessage,
		value Guint32)

	G_dbus_message_get_arg0 func(
		message *GDBusMessage) *Gchar

	G_dbus_message_new_from_blob func(
		blob *Guchar,
		blob_len Gsize,
		capabilities GDBusCapabilityFlags,
		err **GError) *GDBusMessage

	G_dbus_message_bytes_needed func(
		blob *Guchar,
		blob_len Gsize,
		err **GError) Gssize

	G_dbus_message_to_blob func(
		message *GDBusMessage,
		out_size *Gsize,
		capabilities GDBusCapabilityFlags,
		err **GError) *Guchar

	G_dbus_message_to_gerror func(
		message *GDBusMessage,
		err **GError) Gboolean

	G_dbus_method_invocation_get_type func() GType

	G_dbus_method_invocation_get_sender func(
		invocation *GDBusMethodInvocation) *Gchar

	G_dbus_method_invocation_get_object_path func(
		invocation *GDBusMethodInvocation) *Gchar

	G_dbus_method_invocation_get_interface_name func(
		invocation *GDBusMethodInvocation) *Gchar

	G_dbus_method_invocation_get_method_name func(
		invocation *GDBusMethodInvocation) *Gchar

	G_dbus_method_invocation_get_method_info func(
		invocation *GDBusMethodInvocation) *GDBusMethodInfo

	G_dbus_method_invocation_get_connection func(
		invocation *GDBusMethodInvocation) *GDBusConnection

	G_dbus_method_invocation_get_message func(
		invocation *GDBusMethodInvocation) *GDBusMessage

	G_dbus_method_invocation_get_parameters func(
		invocation *GDBusMethodInvocation) *GVariant

	G_dbus_method_invocation_get_user_data func(
		invocation *GDBusMethodInvocation) Gpointer

	G_dbus_method_invocation_return_value func(
		invocation *GDBusMethodInvocation,
		parameters *GVariant)

	//g_dbus_method_invocation_return_error func(invocation  *GDBusMethodInvocation, domain  GQuark, code  Gint, format  *Gchar, ...)

	G_dbus_method_invocation_return_error_valist func(
		invocation *GDBusMethodInvocation,
		domain GQuark,
		code Gint,
		format *Gchar,
		var_args Va_list)

	G_dbus_method_invocation_return_error_literal func(
		invocation *GDBusMethodInvocation,
		domain GQuark,
		code Gint,
		message *Gchar)

	G_dbus_method_invocation_return_gerror func(
		invocation *GDBusMethodInvocation,
		err *GError)

	G_dbus_method_invocation_return_dbus_error func(
		invocation *GDBusMethodInvocation,
		error_name *Gchar,
		error_message *Gchar)

	G_bus_own_name func(
		bus_type GBusType,
		name *Gchar,
		flags GBusNameOwnerFlags,
		bus_acquired_handler GBusAcquiredCallback,
		name_acquired_handler GBusNameAcquiredCallback,
		name_lost_handler GBusNameLostCallback,
		user_data Gpointer,
		user_data_free_func GDestroyNotify) Guint

	G_bus_own_name_on_connection func(
		connection *GDBusConnection,
		name *Gchar,
		flags GBusNameOwnerFlags,
		name_acquired_handler GBusNameAcquiredCallback,
		name_lost_handler GBusNameLostCallback,
		user_data Gpointer,
		user_data_free_func GDestroyNotify) Guint

	G_bus_own_name_with_closures func(
		bus_type GBusType,
		name *Gchar,
		flags GBusNameOwnerFlags,
		bus_acquired_closure *GClosure,
		name_acquired_closure *GClosure,
		name_lost_closure *GClosure) Guint

	G_bus_own_name_on_connection_with_closures func(

		connection *GDBusConnection,
		name *Gchar,
		flags GBusNameOwnerFlags,
		name_acquired_closure *GClosure,
		name_lost_closure *GClosure) Guint

	G_bus_unown_name func(
		owner_id Guint)

	G_bus_watch_name func(
		bus_type GBusType,
		name *Gchar,
		flags GBusNameWatcherFlags,
		name_appeared_handler GBusNameAppearedCallback,
		name_vanished_handler GBusNameVanishedCallback,
		user_data Gpointer,
		user_data_free_func GDestroyNotify) Guint

	G_bus_watch_name_on_connection func(
		connection *GDBusConnection,
		name *Gchar,
		flags GBusNameWatcherFlags,
		name_appeared_handler GBusNameAppearedCallback,
		name_vanished_handler GBusNameVanishedCallback,
		user_data Gpointer,
		user_data_free_func GDestroyNotify) Guint

	G_bus_watch_name_with_closures func(
		bus_type GBusType,
		name *Gchar,
		flags GBusNameWatcherFlags,
		name_appeared_closure *GClosure,
		name_vanished_closure *GClosure) Guint

	G_bus_watch_name_on_connection_with_closures func(

		connection *GDBusConnection,
		name *Gchar,
		flags GBusNameWatcherFlags,
		name_appeared_closure *GClosure,
		name_vanished_closure *GClosure) Guint

	G_bus_unwatch_name func(
		watcher_id Guint)

	G_dbus_proxy_get_type func() GType

	G_dbus_proxy_new func(
		connection *GDBusConnection,
		flags GDBusProxyFlags,
		info *GDBusInterfaceInfo,
		name *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_proxy_new_finish func(
		res *GAsyncResult,
		err **GError) *GDBusProxy

	G_dbus_proxy_new_sync func(
		connection *GDBusConnection,
		flags GDBusProxyFlags,
		info *GDBusInterfaceInfo,
		name *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		cancellable *GCancellable,
		err **GError) *GDBusProxy

	G_dbus_proxy_new_for_bus func(
		bus_type GBusType,
		flags GDBusProxyFlags,
		info *GDBusInterfaceInfo,
		name *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_proxy_new_for_bus_finish func(
		res *GAsyncResult,
		err **GError) *GDBusProxy

	G_dbus_proxy_new_for_bus_sync func(
		bus_type GBusType,
		flags GDBusProxyFlags,
		info *GDBusInterfaceInfo,
		name *Gchar,
		object_path *Gchar,
		interface_name *Gchar,
		cancellable *GCancellable,
		err **GError) *GDBusProxy

	G_dbus_proxy_get_connection func(
		proxy *GDBusProxy) *GDBusConnection

	G_dbus_proxy_get_flags func(
		proxy *GDBusProxy) GDBusProxyFlags

	G_dbus_proxy_get_name func(
		proxy *GDBusProxy) *Gchar

	G_dbus_proxy_get_name_owner func(
		proxy *GDBusProxy) *Gchar

	G_dbus_proxy_get_object_path func(
		proxy *GDBusProxy) *Gchar

	G_dbus_proxy_get_interface_name func(
		proxy *GDBusProxy) *Gchar

	G_dbus_proxy_get_default_timeout func(
		proxy *GDBusProxy) Gint

	G_dbus_proxy_set_default_timeout func(
		proxy *GDBusProxy,
		timeout_msec Gint)

	G_dbus_proxy_get_interface_info func(
		proxy *GDBusProxy) *GDBusInterfaceInfo

	G_dbus_proxy_set_interface_info func(
		proxy *GDBusProxy,
		info *GDBusInterfaceInfo)

	G_dbus_proxy_get_cached_property func(
		proxy *GDBusProxy,
		property_name *Gchar) *GVariant

	G_dbus_proxy_set_cached_property func(
		proxy *GDBusProxy,
		property_name *Gchar,
		value *GVariant)

	G_dbus_proxy_get_cached_property_names func(
		proxy *GDBusProxy) **Gchar

	G_dbus_proxy_call func(
		proxy *GDBusProxy,
		method_name *Gchar,
		parameters *GVariant,
		flags GDBusCallFlags,
		timeout_msec Gint,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_dbus_proxy_call_finish func(
		proxy *GDBusProxy,
		res *GAsyncResult,
		err **GError) *GVariant

	G_dbus_proxy_call_sync func(
		proxy *GDBusProxy,
		method_name *Gchar,
		parameters *GVariant,
		flags GDBusCallFlags,
		timeout_msec Gint,
		cancellable *GCancellable,
		err **GError) *GVariant

	G_dbus_server_get_type func() GType

	G_dbus_server_new_sync func(
		address *Gchar,
		flags GDBusServerFlags,
		guid *Gchar,
		observer *GDBusAuthObserver,
		cancellable *GCancellable,
		err **GError) *GDBusServer

	G_dbus_server_get_client_address func(
		server *GDBusServer) *Gchar

	G_dbus_server_get_guid func(
		server *GDBusServer) *Gchar

	G_dbus_server_get_flags func(
		server *GDBusServer) GDBusServerFlags

	G_dbus_server_start func(
		server *GDBusServer)

	G_dbus_server_stop func(
		server *GDBusServer)

	G_dbus_server_is_active func(
		server *GDBusServer) Gboolean

	G_dbus_is_guid func(
		string *Gchar) Gboolean

	G_dbus_generate_guid func() *Gchar

	G_dbus_is_name func(
		string *Gchar) Gboolean

	G_dbus_is_unique_name func(
		string *Gchar) Gboolean

	G_dbus_is_member_name func(
		string *Gchar) Gboolean

	G_dbus_is_interface_name func(
		string *Gchar) Gboolean

	G_drive_get_type func() GType

	G_drive_get_name func(drive *GDrive) *Char

	G_drive_get_icon func(drive *GDrive) *GIcon

	G_drive_has_volumes func(drive *GDrive) Gboolean

	G_drive_get_volumes func(drive *GDrive) *GList

	G_drive_is_media_removable func(drive *GDrive) Gboolean

	G_drive_has_media func(drive *GDrive) Gboolean

	G_drive_is_media_check_automatic func(drive *GDrive) Gboolean

	G_drive_can_poll_for_media func(drive *GDrive) Gboolean

	G_drive_can_eject func(drive *GDrive) Gboolean

	G_drive_eject func(
		drive *GDrive,
		flags GMountUnmountFlags,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_drive_eject_finish func(
		drive *GDrive,
		result *GAsyncResult,
		err **GError) Gboolean

	G_drive_poll_for_media func(
		drive *GDrive,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_drive_poll_for_media_finish func(
		drive *GDrive,
		result *GAsyncResult,
		err **GError) Gboolean

	G_drive_get_identifier func(drive *GDrive, kind *Char) *Char

	G_drive_enumerate_identifiers func(
		drive *GDrive) **Char

	G_drive_get_start_stop_type func(
		drive *GDrive) GDriveStartStopType

	G_drive_can_start func(drive *GDrive) Gboolean

	G_drive_can_start_degraded func(drive *GDrive) Gboolean

	G_drive_start func(
		drive *GDrive,
		flags GDriveStartFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_drive_start_finish func(
		drive *GDrive,
		result *GAsyncResult,
		err **GError) Gboolean

	G_drive_can_stop func(drive *GDrive) Gboolean

	G_drive_stop func(
		drive *GDrive,
		flags GMountUnmountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_drive_stop_finish func(
		drive *GDrive,
		result *GAsyncResult,
		err **GError) Gboolean

	G_drive_eject_with_operation func(
		drive *GDrive,
		flags GMountUnmountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_drive_eject_with_operation_finish func(
		drive *GDrive,
		result *GAsyncResult,
		err **GError) Gboolean

	G_icon_get_type func() GType

	G_icon_hash func(icon Gconstpointer) Guint

	G_icon_equal func(icon1 *GIcon, icon2 *GIcon) Gboolean

	G_icon_to_string func(icon *GIcon) *Gchar

	G_icon_new_for_string func(str *Gchar, err **GError) *GIcon

	G_emblem_get_type func() GType

	G_emblem_new func(icon *GIcon) *GEmblem

	G_emblem_new_with_origin func(
		icon *GIcon, origin GEmblemOrigin) *GEmblem

	G_emblem_get_icon func(emblem *GEmblem) *GIcon

	G_emblem_get_origin func(emblem *GEmblem) GEmblemOrigin

	G_emblemed_icon_get_type func() GType

	G_emblemed_icon_new func(icon *GIcon, emblem *GEmblem) *GIcon

	G_emblemed_icon_get_icon func(emblemed *GEmblemedIcon) *GIcon

	G_emblemed_icon_get_emblems func(
		emblemed *GEmblemedIcon) *GList

	G_emblemed_icon_add_emblem func(
		emblemed *GEmblemedIcon, emblem *GEmblem)

	G_emblemed_icon_clear_emblems func(emblemed *GEmblemedIcon)

	G_file_attribute_info_list_get_type func() GType

	G_file_attribute_info_list_new func() *GFileAttributeInfoList

	G_file_attribute_info_list_ref func(
		list *GFileAttributeInfoList) *GFileAttributeInfoList

	G_file_attribute_info_list_unref func(
		list *GFileAttributeInfoList)

	G_file_attribute_info_list_dup func(
		list *GFileAttributeInfoList) *GFileAttributeInfoList

	G_file_attribute_info_list_lookup func(
		list *GFileAttributeInfoList,
		name *Char) *GFileAttributeInfo

	G_file_attribute_info_list_add func(
		list *GFileAttributeInfoList,
		name *Char,
		typ GFileAttributeType,
		flags GFileAttributeInfoFlags)

	G_file_enumerator_get_type func() GType

	G_file_enumerator_next_file func(
		enumerator *GFileEnumerator,
		cancellable *GCancellable,
		err **GError) *GFileInfo

	G_file_enumerator_close func(
		enumerator *GFileEnumerator,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_enumerator_next_files_async func(
		enumerator *GFileEnumerator,
		num_files int,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_enumerator_next_files_finish func(
		enumerator *GFileEnumerator,
		result *GAsyncResult,
		err **GError) *GList

	G_file_enumerator_close_async func(
		enumerator *GFileEnumerator,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_enumerator_close_finish func(
		enumerator *GFileEnumerator,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_enumerator_is_closed func(
		enumerator *GFileEnumerator) Gboolean

	G_file_enumerator_has_pending func(
		enumerator *GFileEnumerator) Gboolean

	G_file_enumerator_set_pending func(
		enumerator *GFileEnumerator,
		pending Gboolean)

	G_file_enumerator_get_container func(
		enumerator *GFileEnumerator) *GFile

	G_file_get_type func() GType

	G_file_new_for_path func(path *Char) *GFile

	G_file_new_for_uri func(uri *Char) *GFile

	G_file_new_for_commandline_arg func(arg *Char) *GFile

	G_file_parse_name func(parse_name *Char) *GFile

	G_file_dup func(file *GFile) *GFile

	G_file_hash func(file Gconstpointer) Guint

	G_file_equal func(file1 *GFile, file2 *GFile) Gboolean

	G_file_get_basename func(file *GFile) *Char

	G_file_get_path func(file *GFile) *Char

	G_file_get_uri func(file *GFile) *Char

	G_file_get_parse_name func(file *GFile) *Char

	G_file_get_parent func(file *GFile) *GFile

	G_file_has_parent func(file *GFile, parent *GFile) Gboolean

	G_file_get_child func(file *GFile, name *Char) *GFile

	G_file_get_child_for_display_name func(
		file *GFile, display_name *Char, err **GError) *GFile

	G_file_has_prefix func(file *GFile, prefix *GFile) Gboolean

	G_file_get_relative_path func(
		parent *GFile, descendant *GFile) *Char

	G_file_resolve_relative_path func(
		file *GFile, relative_path *Char) *GFile

	G_file_is_native func(file *GFile) Gboolean

	G_file_has_uri_scheme func(
		file *GFile, uri_scheme *Char) Gboolean

	G_file_get_uri_scheme func(file *GFile) *Char

	G_file_read func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) *GFileInputStream

	G_file_read_async func(
		file *GFile,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_read_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileInputStream

	G_file_append_to func(
		file *GFile,
		flags GFileCreateFlags,
		cancellable *GCancellable,
		err **GError) *GFileOutputStream

	G_file_create func(
		file *GFile,
		flags GFileCreateFlags,
		cancellable *GCancellable,
		err **GError) *GFileOutputStream

	G_file_replace func(
		file *GFile,
		etag *Char,
		make_backup Gboolean,
		flags GFileCreateFlags,
		cancellable *GCancellable,
		err **GError) *GFileOutputStream

	G_file_append_to_async func(
		file *GFile,
		flags GFileCreateFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_append_to_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileOutputStream

	G_file_create_async func(
		file *GFile,
		flags GFileCreateFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_create_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileOutputStream

	G_file_replace_async func(
		file *GFile,
		etag *Char,
		make_backup Gboolean,
		flags GFileCreateFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_replace_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileOutputStream

	G_file_open_readwrite func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) *GFileIOStream

	G_file_open_readwrite_async func(
		file *GFile,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_open_readwrite_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileIOStream

	G_file_create_readwrite func(
		file *GFile,
		flags GFileCreateFlags,
		cancellable *GCancellable,
		err **GError) *GFileIOStream

	G_file_create_readwrite_async func(
		file *GFile,
		flags GFileCreateFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_create_readwrite_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileIOStream

	G_file_replace_readwrite func(
		file *GFile,
		etag *Char,
		make_backup Gboolean,
		flags GFileCreateFlags,
		cancellable *GCancellable,
		err **GError) *GFileIOStream

	G_file_replace_readwrite_async func(
		file *GFile,
		etag *Char,
		make_backup Gboolean,
		flags GFileCreateFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_replace_readwrite_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileIOStream

	G_file_query_exists func(
		file *GFile,
		cancellable *GCancellable) Gboolean

	G_file_query_file_type func(
		file *GFile,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable) GFileType

	G_file_query_info func(
		file *GFile,
		attributes *Char,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) *GFileInfo

	G_file_query_info_async func(
		file *GFile,
		attributes *Char,
		flags GFileQueryInfoFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_query_info_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileInfo

	G_file_query_filesystem_info func(
		file *GFile,
		attributes *Char,
		cancellable *GCancellable,
		err **GError) *GFileInfo

	G_file_query_filesystem_info_async func(
		file *GFile,
		attributes *Char,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_query_filesystem_info_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileInfo

	G_file_find_enclosing_mount func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) *GMount

	G_file_find_enclosing_mount_async func(
		file *GFile,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_find_enclosing_mount_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GMount

	G_file_enumerate_children func(
		file *GFile,
		attributes *Char,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) *GFileEnumerator

	G_file_enumerate_children_async func(
		file *GFile,
		attributes *Char,
		flags GFileQueryInfoFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_enumerate_children_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFileEnumerator

	G_file_set_display_name func(
		file *GFile,
		display_name *Char,
		cancellable *GCancellable,
		err **GError) *GFile

	G_file_set_display_name_async func(
		file *GFile,
		display_name *Char,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_set_display_name_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) *GFile

	G_file_delete func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_trash func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_copy func(
		source *GFile,
		destination *GFile,
		flags GFileCopyFlags,
		cancellable *GCancellable,
		progress_callback GFileProgressCallback,
		progress_callback_data Gpointer,
		err **GError) Gboolean

	G_file_copy_async func(
		source *GFile,
		destination *GFile,
		flags GFileCopyFlags,
		io_priority int,
		cancellable *GCancellable,
		progress_callback GFileProgressCallback,
		progress_callback_data Gpointer,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_copy_finish func(
		file *GFile,
		res *GAsyncResult,
		err **GError) Gboolean

	G_file_move func(
		source *GFile,
		destination *GFile,
		flags GFileCopyFlags,
		cancellable *GCancellable,
		progress_callback GFileProgressCallback,
		progress_callback_data Gpointer,
		err **GError) Gboolean

	G_file_make_directory func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_make_directory_with_parents func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_make_symbolic_link func(
		file *GFile,
		symlink_value *Char,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_query_settable_attributes func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) *GFileAttributeInfoList

	G_file_query_writable_namespaces func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) *GFileAttributeInfoList

	G_file_set_attribute func(
		file *GFile,
		attribute *Char,
		typ GFileAttributeType,
		value_p Gpointer,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_set_attributes_from_info func(
		file *GFile,
		info *GFileInfo,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_set_attributes_async func(
		file *GFile,
		info *GFileInfo,
		flags GFileQueryInfoFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_set_attributes_finish func(
		file *GFile,
		result *GAsyncResult,
		info **GFileInfo,
		err **GError) Gboolean

	G_file_set_attribute_string func(
		file *GFile,
		attribute *Char,
		value *Char,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_set_attribute_byte_string func(
		file *GFile,
		attribute *Char,
		value *Char,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_set_attribute_uint32 func(
		file *GFile,
		attribute *Char,
		value Guint32,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_set_attribute_int32 func(
		file *GFile,
		attribute *Char,
		value Gint32,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_set_attribute_uint64 func(
		file *GFile,
		attribute *Char,
		value Guint64,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_set_attribute_int64 func(
		file *GFile,
		attribute *Char,
		value Gint64,
		flags GFileQueryInfoFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_mount_enclosing_volume func(
		location *GFile,
		flags GMountMountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_mount_enclosing_volume_finish func(
		location *GFile,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_mount_mountable func(
		file *GFile,
		flags GMountMountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_mount_mountable_finish func(
		file *GFile,
		result *GAsyncResult,
		err **GError) *GFile

	G_file_unmount_mountable func(
		file *GFile,
		flags GMountUnmountFlags,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_unmount_mountable_finish func(
		file *GFile,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_unmount_mountable_with_operation func(
		file *GFile,
		flags GMountUnmountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_unmount_mountable_with_operation_finish func(
		file *GFile,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_eject_mountable func(
		file *GFile,
		flags GMountUnmountFlags,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_eject_mountable_finish func(
		file *GFile,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_eject_mountable_with_operation func(
		file *GFile,
		flags GMountUnmountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_eject_mountable_with_operation_finish func(
		file *GFile,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_copy_attributes func(
		source *GFile,
		destination *GFile,
		flags GFileCopyFlags,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_monitor_directory func(
		file *GFile,
		flags GFileMonitorFlags,
		cancellable *GCancellable,
		err **GError) *GFileMonitor

	G_file_monitor_file func(
		file *GFile,
		flags GFileMonitorFlags,
		cancellable *GCancellable,
		err **GError) *GFileMonitor

	G_file_monitor func(
		file *GFile,
		flags GFileMonitorFlags,
		cancellable *GCancellable,
		err **GError) *GFileMonitor

	G_file_start_mountable func(
		file *GFile,
		flags GDriveStartFlags,
		start_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_start_mountable_finish func(
		file *GFile,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_stop_mountable func(
		file *GFile,
		flags GMountUnmountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_stop_mountable_finish func(
		file *GFile,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_poll_mountable func(
		file *GFile,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_poll_mountable_finish func(
		file *GFile,
		result *GAsyncResult,
		err **GError) Gboolean

	G_file_query_default_handler func(
		file *GFile,
		cancellable *GCancellable,
		err **GError) *GAppInfo

	G_file_load_contents func(
		file *GFile,
		cancellable *GCancellable,
		contents **Char,
		length *Gsize,
		etag_out **Char,
		err **GError) Gboolean

	G_file_load_contents_async func(
		file *GFile,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_load_contents_finish func(
		file *GFile,
		res *GAsyncResult,
		contents **Char,
		length *Gsize,
		etag_out **Char,
		err **GError) Gboolean

	G_file_load_partial_contents_async func(
		file *GFile,
		cancellable *GCancellable,
		read_more_callback GFileReadMoreCallback,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_load_partial_contents_finish func(
		file *GFile,
		res *GAsyncResult,
		contents **Char,
		length *Gsize,
		etag_out **Char,
		err **GError) Gboolean

	G_file_replace_contents func(
		file *GFile,
		contents *Char,
		length Gsize,
		etag *Char,
		make_backup Gboolean,
		flags GFileCreateFlags,
		new_etag **Char,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_file_replace_contents_async func(
		file *GFile,
		contents *Char,
		length Gsize,
		etag *Char,
		make_backup Gboolean,
		flags GFileCreateFlags,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_replace_contents_finish func(
		file *GFile,
		res *GAsyncResult,
		new_etag **Char,
		err **GError) Gboolean

	G_file_supports_thread_contexts func(file *GFile) Gboolean

	G_file_icon_get_type func() GType

	G_file_icon_new func(file *GFile) *GIcon

	G_file_icon_get_file func(icon *GFileIcon) *GFile

	G_file_info_get_type func() GType

	G_file_info_new func() *GFileInfo

	G_file_info_dup func(other *GFileInfo) *GFileInfo

	G_file_info_copy_into func(src_info, dest_info *GFileInfo)

	G_file_info_has_attribute func(
		info *GFileInfo, attribute *Char) Gboolean

	G_file_info_has_namespace func(
		info *GFileInfo, name_space *Char) Gboolean

	G_file_info_list_attributes func(
		info *GFileInfo, name_space *Char) **Char

	G_file_info_get_attribute_data func(
		info *GFileInfo,
		attribute *Char,
		typ *GFileAttributeType,
		value_pp *Gpointer,
		status *GFileAttributeStatus) Gboolean

	G_file_info_get_attribute_type func(
		info *GFileInfo,
		attribute *Char) GFileAttributeType

	G_file_info_remove_attribute func(
		info *GFileInfo,
		attribute *Char)

	G_file_info_get_attribute_status func(
		info *GFileInfo,
		attribute *Char) GFileAttributeStatus

	G_file_info_set_attribute_status func(
		info *GFileInfo,
		attribute *Char,
		status GFileAttributeStatus) Gboolean

	G_file_info_get_attribute_as_string func(
		info *GFileInfo,
		attribute *Char) *Char

	G_file_info_get_attribute_string func(
		info *GFileInfo,
		attribute *Char) *Char

	G_file_info_get_attribute_byte_string func(
		info *GFileInfo,
		attribute *Char) *Char

	G_file_info_get_attribute_boolean func(
		info *GFileInfo,
		attribute *Char) Gboolean

	G_file_info_get_attribute_uint32 func(
		info *GFileInfo,
		attribute *Char) Guint32

	G_file_info_get_attribute_int32 func(
		info *GFileInfo,
		attribute *Char) Gint32

	G_file_info_get_attribute_uint64 func(
		info *GFileInfo,
		attribute *Char) Guint64

	G_file_info_get_attribute_int64 func(
		info *GFileInfo,
		attribute *Char) Gint64

	G_file_info_get_attribute_object func(
		info *GFileInfo,
		attribute *Char) *GObject

	G_file_info_get_attribute_stringv func(
		info *GFileInfo,
		attribute *Char) **Char

	G_file_info_set_attribute func(
		info *GFileInfo,
		attribute *Char,
		typ GFileAttributeType,
		value_p Gpointer)

	G_file_info_set_attribute_string func(
		info *GFileInfo,
		attribute *Char,
		attr_value *Char)

	G_file_info_set_attribute_byte_string func(
		info *GFileInfo,
		attribute *Char,
		attr_value *Char)

	G_file_info_set_attribute_boolean func(
		info *GFileInfo,
		attribute *Char,
		attr_value Gboolean)

	G_file_info_set_attribute_uint32 func(
		info *GFileInfo,
		attribute *Char,
		attr_value Guint32)

	G_file_info_set_attribute_int32 func(
		info *GFileInfo,
		attribute *Char,
		attr_value Gint32)

	G_file_info_set_attribute_uint64 func(
		info *GFileInfo,
		attribute *Char,
		attr_value Guint64)

	G_file_info_set_attribute_int64 func(
		info *GFileInfo,
		attribute *Char,
		attr_value Gint64)

	G_file_info_set_attribute_object func(
		info *GFileInfo,
		attribute *Char,
		attr_value *GObject)

	G_file_info_set_attribute_stringv func(
		info *GFileInfo,
		attribute *Char,
		attr_value **Char)

	G_file_info_clear_status func(
		info *GFileInfo)

	G_file_info_get_file_type func(
		info *GFileInfo) GFileType

	G_file_info_get_is_hidden func(
		info *GFileInfo) Gboolean

	G_file_info_get_is_backup func(
		info *GFileInfo) Gboolean

	G_file_info_get_is_symlink func(
		info *GFileInfo) Gboolean

	G_file_info_get_name func(
		info *GFileInfo) *Char

	G_file_info_get_display_name func(
		info *GFileInfo) *Char

	G_file_info_get_edit_name func(
		info *GFileInfo) *Char

	G_file_info_get_icon func(
		info *GFileInfo) *GIcon

	G_file_info_get_content_type func(
		info *GFileInfo) *Char

	G_file_info_get_size func(
		info *GFileInfo) Goffset

	G_file_info_get_modification_time func(
		info *GFileInfo,
		result *GTimeVal)

	G_file_info_get_symlink_target func(
		info *GFileInfo) *Char

	G_file_info_get_etag func(
		info *GFileInfo) *Char

	G_file_info_get_sort_order func(
		info *GFileInfo) Gint32

	G_file_info_set_attribute_mask func(
		info *GFileInfo,
		mask *GFileAttributeMatcher)

	G_file_info_unset_attribute_mask func(
		info *GFileInfo)

	G_file_info_set_file_type func(
		info *GFileInfo,
		typ GFileType)

	G_file_info_set_is_hidden func(
		info *GFileInfo,
		is_hidden Gboolean)

	G_file_info_set_is_symlink func(
		info *GFileInfo,
		is_symlink Gboolean)

	G_file_info_set_name func(
		info *GFileInfo,
		name *Char)

	G_file_info_set_display_name func(
		info *GFileInfo,
		display_name *Char)

	G_file_info_set_edit_name func(
		info *GFileInfo,
		edit_name *Char)

	G_file_info_set_icon func(
		info *GFileInfo,
		icon *GIcon)

	G_file_info_set_content_type func(
		info *GFileInfo,
		content_type *Char)

	G_file_info_set_size func(
		info *GFileInfo,
		size Goffset)

	G_file_info_set_modification_time func(
		info *GFileInfo,
		mtime *GTimeVal)

	G_file_info_set_symlink_target func(
		info *GFileInfo,
		symlink_target *Char)

	G_file_info_set_sort_order func(
		info *GFileInfo,
		sort_order Gint32)

	G_file_attribute_matcher_get_type func() GType

	G_file_attribute_matcher_new func(
		attributes *Char) *GFileAttributeMatcher

	G_file_attribute_matcher_ref func(
		matcher *GFileAttributeMatcher) *GFileAttributeMatcher

	G_file_attribute_matcher_unref func(
		matcher *GFileAttributeMatcher)

	G_file_attribute_matcher_matches func(
		matcher *GFileAttributeMatcher,
		attribute *Char) Gboolean

	G_file_attribute_matcher_matches_only func(
		matcher *GFileAttributeMatcher,
		attribute *Char) Gboolean

	G_file_attribute_matcher_enumerate_namespace func(
		matcher *GFileAttributeMatcher,
		ns *Char) Gboolean

	G_file_attribute_matcher_enumerate_next func(
		matcher *GFileAttributeMatcher) *Char

	G_file_input_stream_get_type func() GType

	G_file_input_stream_query_info func(
		stream *GFileInputStream,
		attributes *Char,
		cancellable *GCancellable,
		err **GError) *GFileInfo

	G_file_input_stream_query_info_async func(
		stream *GFileInputStream,
		attributes *Char,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_input_stream_query_info_finish func(
		stream *GFileInputStream,
		result *GAsyncResult,
		err **GError) *GFileInfo

	G_io_error_quark func() GQuark

	G_io_error_from_errno func(
		err_no Gint) GIOErrorEnum

	G_io_error_from_win32_error func(
		error_code Gint) GIOErrorEnum

	G_io_stream_get_type func() GType

	G_io_stream_get_input_stream func(
		stream *GIOStream) *GInputStream

	G_io_stream_get_output_stream func(
		stream *GIOStream) *GOutputStream

	G_io_stream_splice_async func(
		stream1 *GIOStream,
		stream2 *GIOStream,
		flags GIOStreamSpliceFlags,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_io_stream_splice_finish func(
		result *GAsyncResult,
		err **GError) Gboolean

	G_io_stream_close func(
		stream *GIOStream,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_io_stream_close_async func(
		stream *GIOStream,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_io_stream_close_finish func(
		stream *GIOStream,
		result *GAsyncResult,
		err **GError) Gboolean

	G_io_stream_is_closed func(
		stream *GIOStream) Gboolean

	G_io_stream_has_pending func(
		stream *GIOStream) Gboolean

	G_io_stream_set_pending func(
		stream *GIOStream,
		err **GError) Gboolean

	G_io_stream_clear_pending func(
		stream *GIOStream)

	G_file_io_stream_get_type func() GType

	G_file_io_stream_query_info func(
		stream *GFileIOStream,
		attributes *Char,
		cancellable *GCancellable,
		err **GError) *GFileInfo

	G_file_io_stream_query_info_async func(
		stream *GFileIOStream,
		attributes *Char,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_io_stream_query_info_finish func(
		stream *GFileIOStream,
		result *GAsyncResult,
		err **GError) *GFileInfo

	G_file_io_stream_get_etag func(
		stream *GFileIOStream) *Char

	G_file_monitor_get_type func() GType

	G_file_monitor_cancel func(
		monitor *GFileMonitor) Gboolean

	G_file_monitor_is_cancelled func(
		monitor *GFileMonitor) Gboolean

	G_file_monitor_set_rate_limit func(
		monitor *GFileMonitor,
		limit_msecs Gint)

	G_file_monitor_emit_event func(
		monitor *GFileMonitor,
		child *GFile,
		other_file *GFile,
		event_type GFileMonitorEvent)

	G_filename_completer_get_type func() GType

	G_filename_completer_new func() *GFilenameCompleter

	G_filename_completer_get_completion_suffix func(
		completer *GFilenameCompleter,
		initial_text *Char) *Char

	G_filename_completer_get_completions func(
		completer *GFilenameCompleter,
		initial_text *Char) **Char

	G_filename_completer_set_dirs_only func(
		completer *GFilenameCompleter,
		dirs_only Gboolean)

	G_file_output_stream_get_type func() GType

	G_file_output_stream_query_info func(
		stream *GFileOutputStream,
		attributes *Char,
		cancellable *GCancellable,
		err **GError) *GFileInfo

	G_file_output_stream_query_info_async func(
		stream *GFileOutputStream,
		attributes *Char,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_file_output_stream_query_info_finish func(
		stream *GFileOutputStream,
		result *GAsyncResult,
		err **GError) *GFileInfo

	G_file_output_stream_get_etag func(
		stream *GFileOutputStream) *Char

	G_inet_address_get_type func() GType

	G_inet_address_new_from_string func(
		string *Gchar) *GInetAddress

	G_inet_address_new_from_bytes func(
		bytes *Guint8,
		family GSocketFamily) *GInetAddress

	G_inet_address_new_loopback func(
		family GSocketFamily) *GInetAddress

	G_inet_address_new_any func(
		family GSocketFamily) *GInetAddress

	G_inet_address_to_string func(
		address *GInetAddress) *Gchar

	G_inet_address_to_bytes func(
		address *GInetAddress) *Guint8

	G_inet_address_get_native_size func(
		address *GInetAddress) Gsize

	G_inet_address_get_family func(
		address *GInetAddress) GSocketFamily

	G_inet_address_get_is_any func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_loopback func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_link_local func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_site_local func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_multicast func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_mc_global func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_mc_link_local func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_mc_node_local func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_mc_org_local func(
		address *GInetAddress) Gboolean

	G_inet_address_get_is_mc_site_local func(
		address *GInetAddress) Gboolean

	G_socket_address_get_type func() GType

	G_socket_address_get_family func(
		address *GSocketAddress) GSocketFamily

	G_socket_address_new_from_native func(
		native Gpointer,
		len Gsize) *GSocketAddress

	G_socket_address_to_native func(
		address *GSocketAddress,
		dest Gpointer,
		destlen Gsize,
		err **GError) Gboolean

	G_socket_address_get_native_size func(
		address *GSocketAddress) Gssize

	G_inet_socket_address_get_type func() GType

	G_inet_socket_address_new func(
		address *GInetAddress,
		port Guint16) *GSocketAddress

	G_inet_socket_address_get_address func(
		address *GInetSocketAddress) *GInetAddress

	G_inet_socket_address_get_port func(
		address *GInetSocketAddress) Guint16

	G_app_info_create_flags_get_type func() GType

	G_converter_flags_get_type func() GType

	G_converter_result_get_type func() GType

	G_data_stream_byte_order_get_type func() GType

	G_data_stream_newline_type_get_type func() GType

	G_file_attribute_type_get_type func() GType

	G_file_attribute_info_flags_get_type func() GType

	G_file_attribute_status_get_type func() GType

	G_file_query_info_flags_get_type func() GType

	G_file_create_flags_get_type func() GType

	G_mount_mount_flags_get_type func() GType

	G_mount_unmount_flags_get_type func() GType

	G_drive_start_flags_get_type func() GType

	G_drive_start_stop_type_get_type func() GType

	G_file_copy_flags_get_type func() GType

	G_file_monitor_flags_get_type func() GType

	G_file_type_get_type func() GType

	G_filesystem_preview_type_get_type func() GType

	G_file_monitor_event_get_type func() GType

	G_io_error_enum_get_type func() GType

	G_ask_password_flags_get_type func() GType

	G_password_save_get_type func() GType

	G_mount_operation_result_get_type func() GType

	G_output_stream_splice_flags_get_type func() GType

	G_io_stream_splice_flags_get_type func() GType

	G_emblem_origin_get_type func() GType

	G_resolver_error_get_type func() GType

	G_socket_family_get_type func() GType

	G_socket_type_get_type func() GType

	G_socket_msg_flags_get_type func() GType

	G_socket_protocol_get_type func() GType

	G_zlib_compressor_format_get_type func() GType

	G_unix_socket_address_type_get_type func() GType

	G_bus_type_get_type func() GType

	G_bus_name_owner_flags_get_type func() GType

	G_bus_name_watcher_flags_get_type func() GType

	G_dbus_proxy_flags_get_type func() GType

	G_dbus_error_get_type func() GType

	G_dbus_connection_flags_get_type func() GType

	G_dbus_capability_flags_get_type func() GType

	G_dbus_call_flags_get_type func() GType

	G_dbus_message_type_get_type func() GType

	G_dbus_message_flags_get_type func() GType

	G_dbus_message_header_field_get_type func() GType

	G_dbus_property_info_flags_get_type func() GType

	G_dbus_subtree_flags_get_type func() GType

	G_dbus_server_flags_get_type func() GType

	G_dbus_signal_flags_get_type func() GType

	G_dbus_send_message_flags_get_type func() GType

	G_credentials_type_get_type func() GType

	G_dbus_message_byte_order_get_type func() GType

	G_application_flags_get_type func() GType

	G_tls_error_get_type func() GType

	G_tls_certificate_flags_get_type func() GType

	G_tls_authentication_mode_get_type func() GType

	G_tls_rehandshake_mode_get_type func() GType

	G_settings_bind_flags_get_type func() GType

	G_io_module_get_type func() GType

	G_io_module_new func(
		filename *Gchar) *GIOModule

	G_io_modules_scan_all_in_directory func(
		dirname *Char)

	G_io_modules_load_all_in_directory func(
		dirname *Gchar) *GList

	G_io_extension_point_register func(
		name *Char) *GIOExtensionPoint

	G_io_extension_point_lookup func(
		name *Char) *GIOExtensionPoint

	G_io_extension_point_set_required_type func(
		extension_point *GIOExtensionPoint,
		typ GType)

	G_io_extension_point_get_required_type func(
		extension_point *GIOExtensionPoint) GType

	G_io_extension_point_get_extensions func(
		extension_point *GIOExtensionPoint) *GList

	G_io_extension_point_get_extension_by_name func(
		extension_point *GIOExtensionPoint,
		name *Char) *GIOExtension

	G_io_extension_point_implement func(
		extension_point_name *Char,
		typ GType,
		extension_name *Char,
		priority Gint) *GIOExtension

	G_io_extension_get_type func(
		extension *GIOExtension) GType

	G_io_extension_get_name func(
		extension *GIOExtension) *Char

	G_io_extension_get_priority func(
		extension *GIOExtension) Gint

	G_io_extension_ref_class func(
		extension *GIOExtension) *GTypeClass

	G_io_module_load func(
		module *GIOModule)

	G_io_module_unload func(
		module *GIOModule)

	G_io_module_query func() **Char

	G_io_scheduler_push_job func(
		job_func GIOSchedulerJobFunc,
		user_data Gpointer,
		notify GDestroyNotify,
		io_priority Gint,
		cancellable *GCancellable)

	G_io_scheduler_cancel_all_jobs func()

	G_io_scheduler_job_send_to_mainloop func(
		job *GIOSchedulerJob,
		f GSourceFunc,
		user_data Gpointer,
		notify GDestroyNotify) Gboolean

	G_io_scheduler_job_send_to_mainloop_async func(
		job *GIOSchedulerJob,
		f GSourceFunc,
		user_data Gpointer,
		notify GDestroyNotify)

	G_loadable_icon_get_type func() GType

	G_loadable_icon_load func(
		icon *GLoadableIcon,
		size int,
		t **Char,
		cancellable *GCancellable,
		err **GError) *GInputStream

	G_loadable_icon_load_async func(
		icon *GLoadableIcon,
		size int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_loadable_icon_load_finish func(
		icon *GLoadableIcon,
		res *GAsyncResult,
		typ **Char,
		err **GError) *GInputStream

	G_memory_input_stream_get_type func() GType

	G_memory_input_stream_new func() *GInputStream

	G_memory_input_stream_new_from_data func(
		data *Void,
		len Gssize,
		destroy GDestroyNotify) *GInputStream

	G_memory_input_stream_add_data func(
		stream *GMemoryInputStream,
		data *Void,
		len Gssize,
		destroy GDestroyNotify)

	G_memory_output_stream_get_type func() GType

	G_memory_output_stream_new func(
		data Gpointer,
		size Gsize,
		realloc_function GReallocFunc,
		destroy_function GDestroyNotify) *GOutputStream

	G_memory_output_stream_get_data func(
		ostream *GMemoryOutputStream) Gpointer

	G_memory_output_stream_get_size func(
		ostream *GMemoryOutputStream) Gsize

	G_memory_output_stream_get_data_size func(
		ostream *GMemoryOutputStream) Gsize

	G_memory_output_stream_steal_data func(
		ostream *GMemoryOutputStream) Gpointer

	G_mount_get_type func() GType

	G_mount_get_root func(
		mount *GMount) *GFile

	G_mount_get_default_location func(
		mount *GMount) *GFile

	G_mount_get_name func(
		mount *GMount) *Char

	G_mount_get_icon func(
		mount *GMount) *GIcon

	G_mount_get_uuid func(
		mount *GMount) *Char

	G_mount_get_volume func(
		mount *GMount) *GVolume

	G_mount_get_drive func(
		mount *GMount) *GDrive

	G_mount_can_unmount func(
		mount *GMount) Gboolean

	G_mount_can_eject func(
		mount *GMount) Gboolean

	G_mount_unmount func(
		mount *GMount,
		flags GMountUnmountFlags,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_mount_unmount_finish func(
		mount *GMount,
		result *GAsyncResult,
		err **GError) Gboolean

	G_mount_eject func(
		mount *GMount,
		flags GMountUnmountFlags,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_mount_eject_finish func(
		mount *GMount,
		result *GAsyncResult,
		err **GError) Gboolean

	G_mount_remount func(
		mount *GMount,
		flags GMountMountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_mount_remount_finish func(
		mount *GMount,
		result *GAsyncResult,
		err **GError) Gboolean

	G_mount_guess_content_type func(
		mount *GMount,
		force_rescan Gboolean,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_mount_guess_content_type_finish func(
		mount *GMount,
		result *GAsyncResult,
		err **GError) **Gchar

	G_mount_guess_content_type_sync func(
		mount *GMount,
		force_rescan Gboolean,
		cancellable *GCancellable,
		err **GError) **Gchar

	G_mount_is_shadowed func(
		mount *GMount) Gboolean

	G_mount_shadow func(
		mount *GMount)

	G_mount_unshadow func(
		mount *GMount)

	G_mount_unmount_with_operation func(
		mount *GMount,
		flags GMountUnmountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_mount_unmount_with_operation_finish func(
		mount *GMount,
		result *GAsyncResult,
		err **GError) Gboolean

	G_mount_eject_with_operation func(
		mount *GMount,
		flags GMountUnmountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_mount_eject_with_operation_finish func(
		mount *GMount,
		result *GAsyncResult,
		err **GError) Gboolean

	G_mount_operation_get_type func() GType

	G_mount_operation_new func() *GMountOperation

	G_mount_operation_get_username func(
		op *GMountOperation) *Char

	G_mount_operation_set_username func(
		op *GMountOperation,
		username *Char)

	G_mount_operation_get_password func(
		op *GMountOperation) *Char

	G_mount_operation_set_password func(
		op *GMountOperation,
		password *Char)

	G_mount_operation_get_anonymous func(
		op *GMountOperation) Gboolean

	G_mount_operation_set_anonymous func(
		op *GMountOperation,
		anonymous Gboolean)

	G_mount_operation_get_domain func(
		op *GMountOperation) *Char

	G_mount_operation_set_domain func(
		op *GMountOperation,
		domain *Char)

	G_mount_operation_get_password_save func(
		op *GMountOperation) GPasswordSave

	G_mount_operation_set_password_save func(
		op *GMountOperation,
		save GPasswordSave)

	G_mount_operation_get_choice func(
		op *GMountOperation) int

	G_mount_operation_set_choice func(
		op *GMountOperation,
		choice int)

	G_mount_operation_reply func(
		op *GMountOperation,
		result GMountOperationResult)

	G_volume_monitor_get_type func() GType

	G_volume_monitor_get func() *GVolumeMonitor

	G_volume_monitor_get_connected_drives func(
		volume_monitor *GVolumeMonitor) *GList

	G_volume_monitor_get_volumes func(
		volume_monitor *GVolumeMonitor) *GList

	G_volume_monitor_get_mounts func(
		volume_monitor *GVolumeMonitor) *GList

	G_volume_monitor_get_volume_for_uuid func(
		volume_monitor *GVolumeMonitor,
		uuid *Char) *GVolume

	G_volume_monitor_get_mount_for_uuid func(
		volume_monitor *GVolumeMonitor,
		uuid *Char) *GMount

	G_volume_monitor_adopt_orphan_mount func(
		mount *GMount) *GVolume

	G_native_volume_monitor_get_type func() GType

	G_network_address_get_type func() GType

	G_network_address_new func(
		hostname *Gchar,
		port Guint16) *GSocketConnectable

	G_network_address_parse func(
		host_and_port *Gchar,
		default_port Guint16,
		err **GError) *GSocketConnectable

	G_network_address_parse_uri func(
		uri *Gchar,
		default_port Guint16,
		err **GError) *GSocketConnectable

	G_network_address_get_hostname func(
		addr *GNetworkAddress) *Gchar

	G_network_address_get_port func(
		addr *GNetworkAddress) Guint16

	G_network_address_get_scheme func(
		addr *GNetworkAddress) *Gchar

	G_network_service_get_type func() GType

	G_network_service_new func(
		service *Gchar,
		protocol *Gchar,
		domain *Gchar) *GSocketConnectable

	G_network_service_get_service func(
		srv *GNetworkService) *Gchar

	G_network_service_get_protocol func(
		srv *GNetworkService) *Gchar

	G_network_service_get_domain func(
		srv *GNetworkService) *Gchar

	G_network_service_get_scheme func(
		srv *GNetworkService) *Gchar

	G_network_service_set_scheme func(
		srv *GNetworkService,
		scheme *Gchar)

	G_permission_get_type func() GType

	G_permission_acquire func(
		permission *GPermission,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_permission_acquire_async func(
		permission *GPermission,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_permission_acquire_finish func(
		permission *GPermission,
		result *GAsyncResult,
		err **GError) Gboolean

	G_permission_release func(
		permission *GPermission,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_permission_release_async func(
		permission *GPermission,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_permission_release_finish func(
		permission *GPermission,
		result *GAsyncResult,
		err **GError) Gboolean

	G_permission_get_allowed func(
		permission *GPermission) Gboolean

	G_permission_get_can_acquire func(
		permission *GPermission) Gboolean

	G_permission_get_can_release func(
		permission *GPermission) Gboolean

	G_permission_impl_update func(
		permission *GPermission,
		allowed Gboolean,
		can_acquire Gboolean,
		can_release Gboolean)

	G_pollable_input_stream_get_type func() GType

	G_pollable_input_stream_can_poll func(
		stream *GPollableInputStream) Gboolean

	G_pollable_input_stream_is_readable func(
		stream *GPollableInputStream) Gboolean

	G_pollable_input_stream_create_source func(
		stream *GPollableInputStream,
		cancellable *GCancellable) *GSource

	G_pollable_input_stream_read_nonblocking func(
		stream *GPollableInputStream,
		buffer *Void,
		size Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_pollable_source_new func(
		pollable_stream *GObject) *GSource

	G_pollable_output_stream_get_type func() GType

	G_pollable_output_stream_can_poll func(
		stream *GPollableOutputStream) Gboolean

	G_pollable_output_stream_is_writable func(
		stream *GPollableOutputStream) Gboolean

	G_pollable_output_stream_create_source func(
		stream *GPollableOutputStream,
		cancellable *GCancellable) *GSource

	G_pollable_output_stream_write_nonblocking func(
		stream *GPollableOutputStream,
		buffer *Void,
		size Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_proxy_get_type func() GType

	G_proxy_get_default_for_protocol func(
		protocol *Gchar) *GProxy

	G_proxy_connect func(
		proxy *GProxy,
		connection *GIOStream,
		proxy_address *GProxyAddress,
		cancellable *GCancellable,
		err **GError) *GIOStream

	G_proxy_connect_async func(
		proxy *GProxy,
		connection *GIOStream,
		proxy_address *GProxyAddress,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_proxy_connect_finish func(
		proxy *GProxy,
		result *GAsyncResult,
		err **GError) *GIOStream

	G_proxy_supports_hostname func(
		proxy *GProxy) Gboolean

	G_proxy_address_get_type func() GType

	G_proxy_address_new func(
		inetaddr *GInetAddress,
		port Guint16,
		protocol *Gchar,
		dest_hostname *Gchar,
		dest_port Guint16,
		username *Gchar,
		password *Gchar) *GSocketAddress

	G_proxy_address_get_protocol func(
		proxy *GProxyAddress) *Gchar

	G_proxy_address_get_destination_hostname func(
		proxy *GProxyAddress) *Gchar

	G_proxy_address_get_destination_port func(
		proxy *GProxyAddress) Guint16

	G_proxy_address_get_username func(
		proxy *GProxyAddress) *Gchar

	G_proxy_address_get_password func(
		proxy *GProxyAddress) *Gchar

	G_socket_address_enumerator_get_type func() GType

	G_socket_address_enumerator_next func(
		enumerator *GSocketAddressEnumerator,
		cancellable *GCancellable,
		err **GError) *GSocketAddress

	G_socket_address_enumerator_next_async func(
		enumerator *GSocketAddressEnumerator,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_socket_address_enumerator_next_finish func(
		enumerator *GSocketAddressEnumerator,
		result *GAsyncResult,
		err **GError) *GSocketAddress

	G_proxy_address_enumerator_get_type func() GType

	G_proxy_resolver_get_type func() GType

	G_proxy_resolver_get_default func() *GProxyResolver

	G_proxy_resolver_is_supported func(
		resolver *GProxyResolver) Gboolean

	G_proxy_resolver_lookup func(
		resolver *GProxyResolver,
		uri *Gchar,
		cancellable *GCancellable,
		err **GError) **Gchar

	G_proxy_resolver_lookup_async func(
		resolver *GProxyResolver,
		uri *Gchar,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_proxy_resolver_lookup_finish func(
		resolver *GProxyResolver,
		result *GAsyncResult,
		err **GError) **Gchar

	G_resolver_get_type func() GType

	G_resolver_get_default func() *GResolver

	G_resolver_set_default func(
		resolver *GResolver)

	G_resolver_lookup_by_name func(
		resolver *GResolver,
		hostname *Gchar,
		cancellable *GCancellable,
		err **GError) *GList

	G_resolver_lookup_by_name_async func(
		resolver *GResolver,
		hostname *Gchar,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_resolver_lookup_by_name_finish func(
		resolver *GResolver,
		result *GAsyncResult,
		err **GError) *GList

	G_resolver_free_addresses func(
		addresses *GList)

	G_resolver_lookup_by_address func(
		resolver *GResolver,
		address *GInetAddress,
		cancellable *GCancellable,
		err **GError) *Gchar

	G_resolver_lookup_by_address_async func(
		resolver *GResolver,
		address *GInetAddress,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_resolver_lookup_by_address_finish func(
		resolver *GResolver,
		result *GAsyncResult,
		err **GError) *Gchar

	G_resolver_lookup_service func(
		resolver *GResolver,
		service *Gchar,
		protocol *Gchar,
		domain *Gchar,
		cancellable *GCancellable,
		err **GError) *GList

	G_resolver_lookup_service_async func(
		resolver *GResolver,
		service *Gchar,
		protocol *Gchar,
		domain *Gchar,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_resolver_lookup_service_finish func(
		resolver *GResolver,
		result *GAsyncResult,
		err **GError) *GList

	G_resolver_free_targets func(
		targets *GList)

	G_resolver_error_quark func() GQuark

	G_seekable_get_type func() GType

	G_seekable_tell func(
		seekable *GSeekable) Goffset

	G_seekable_can_seek func(
		seekable *GSeekable) Gboolean

	G_seekable_seek func(
		seekable *GSeekable,
		offset Goffset,
		typ GSeekType,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_seekable_can_truncate func(
		seekable *GSeekable) Gboolean

	G_seekable_truncate func(
		seekable *GSeekable,
		offset Goffset,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_settings_get_type func() GType

	G_settings_list_schemas func() **Gchar

	G_settings_list_relocatable_schemas func() **Gchar

	G_settings_new func(
		schema *Gchar) *GSettings

	G_settings_new_with_path func(
		schema *Gchar,
		path *Gchar) *GSettings

	G_settings_new_with_backend func(
		schema *Gchar,
		backend *GSettingsBackend) *GSettings

	G_settings_new_with_backend_and_path func(
		schema *Gchar,
		backend *GSettingsBackend,
		path *Gchar) *GSettings

	G_settings_list_children func(
		settings *GSettings) **Gchar

	G_settings_list_keys func(
		settings *GSettings) **Gchar

	G_settings_get_range func(
		settings *GSettings,
		key *Gchar) *GVariant

	G_settings_range_check func(
		settings *GSettings,
		key *Gchar,
		value *GVariant) Gboolean

	G_settings_set_value func(
		settings *GSettings,
		key *Gchar,
		value *GVariant) Gboolean

	G_settings_get_value func(
		settings *GSettings,
		key *Gchar) *GVariant

	//TODO(t):Variant
	//g_settings_set func(settings  *GSettings, key  *Gchar, format  *Gchar, ...) Gboolean
	//g_settings_get func(settings  *GSettings, key  *Gchar, format  *Gchar, ...)

	G_settings_reset func(
		settings *GSettings,
		key *Gchar)

	G_settings_get_int func(
		settings *GSettings,
		key *Gchar) Gint

	G_settings_set_int func(
		settings *GSettings,
		key *Gchar,
		value Gint) Gboolean

	G_settings_get_string func(
		settings *GSettings,
		key *Gchar) *Gchar

	G_settings_set_string func(
		settings *GSettings,
		key *Gchar,
		value *Gchar) Gboolean

	G_settings_get_boolean func(
		settings *GSettings,
		key *Gchar) Gboolean

	G_settings_set_boolean func(
		settings *GSettings,
		key *Gchar,
		value Gboolean) Gboolean

	G_settings_get_double func(
		settings *GSettings,
		key *Gchar) Gdouble

	G_settings_set_double func(
		settings *GSettings,
		key *Gchar,
		value Gdouble) Gboolean

	G_settings_get_strv func(
		settings *GSettings,
		key *Gchar) **Gchar

	G_settings_set_strv func(
		settings *GSettings,
		key *Gchar,
		value **Gchar) Gboolean

	G_settings_get_enum func(
		settings *GSettings,
		key *Gchar) Gint

	G_settings_set_enum func(
		settings *GSettings,
		key *Gchar,
		value Gint) Gboolean

	G_settings_get_flags func(
		settings *GSettings,
		key *Gchar) Guint

	G_settings_set_flags func(
		settings *GSettings,
		key *Gchar,
		value Guint) Gboolean

	G_settings_get_child func(
		settings *GSettings,
		name *Gchar) *GSettings

	G_settings_is_writable func(
		settings *GSettings,
		name *Gchar) Gboolean

	G_settings_delay func(
		settings *GSettings)

	G_settings_apply func(
		settings *GSettings)

	G_settings_revert func(
		settings *GSettings)

	G_settings_get_has_unapplied func(
		settings *GSettings) Gboolean

	G_settings_sync func()

	G_settings_bind func(
		settings *GSettings,
		key *Gchar,
		object Gpointer,
		property *Gchar,
		flags GSettingsBindFlags)

	G_settings_bind_with_mapping func(
		settings *GSettings,
		key *Gchar,
		object Gpointer,
		property *Gchar,
		flags GSettingsBindFlags,
		get_mapping GSettingsBindGetMapping,
		set_mapping GSettingsBindSetMapping,
		user_data Gpointer,
		destroy GDestroyNotify)

	G_settings_bind_writable func(
		settings *GSettings,
		key *Gchar,
		object Gpointer,
		property *Gchar,
		inverted Gboolean)

	G_settings_unbind func(
		object Gpointer,
		property *Gchar)

	G_settings_get_mapped func(
		settings *GSettings,
		key *Gchar,
		mapping GSettingsGetMapping,
		user_data Gpointer) Gpointer

	G_simple_async_result_get_type func() GType

	G_simple_async_result_new func(
		source_object *GObject,
		callback GAsyncReadyCallback,
		user_data Gpointer,
		source_tag Gpointer) *GSimpleAsyncResult

	//TODO(t):Variant
	//g_simple_async_result_new_error func(source_object  *GObject, callback  GAsyncReadyCallback, user_data  Gpointer, domain  GQuark, code  Gint, format  *Char, ...) *GSimpleAsyncResult

	G_simple_async_result_new_from_error func(
		source_object *GObject,
		callback GAsyncReadyCallback,
		user_data Gpointer,
		err *GError) *GSimpleAsyncResult

	G_simple_async_result_new_take_error func(
		source_object *GObject,
		callback GAsyncReadyCallback,
		user_data Gpointer,
		err *GError) *GSimpleAsyncResult

	G_simple_async_result_set_op_res_gpointer func(
		simple *GSimpleAsyncResult,
		op_res Gpointer,
		destroy_op_res GDestroyNotify)

	G_simple_async_result_get_op_res_gpointer func(
		simple *GSimpleAsyncResult) Gpointer

	G_simple_async_result_set_op_res_gssize func(
		simple *GSimpleAsyncResult,
		op_res Gssize)

	G_simple_async_result_get_op_res_gssize func(
		simple *GSimpleAsyncResult) Gssize

	G_simple_async_result_set_op_res_gboolean func(
		simple *GSimpleAsyncResult,
		op_res Gboolean)

	G_simple_async_result_get_op_res_gboolean func(
		simple *GSimpleAsyncResult) Gboolean

	G_simple_async_result_get_source_tag func(
		simple *GSimpleAsyncResult) Gpointer

	G_simple_async_result_set_handle_cancellation func(
		simple *GSimpleAsyncResult,
		handle_cancellation Gboolean)

	G_simple_async_result_complete func(
		simple *GSimpleAsyncResult)

	G_simple_async_result_complete_in_idle func(
		simple *GSimpleAsyncResult)

	G_simple_async_result_run_in_thread func(
		simple *GSimpleAsyncResult,
		f GSimpleAsyncThreadFunc,
		io_priority int,
		cancellable *GCancellable)

	G_simple_async_result_set_from_error func(
		simple *GSimpleAsyncResult,
		err *GError)

	G_simple_async_result_take_error func(
		simple *GSimpleAsyncResult,
		err *GError)

	G_simple_async_result_propagate_error func(
		simple *GSimpleAsyncResult,
		dest **GError) Gboolean

	//TODO(t):Variant
	//g_simple_async_result_set_error func(simple  *GSimpleAsyncResult, domain  GQuark, code  Gint, format  *Char, ...)

	G_simple_async_result_set_error_va func(
		simple *GSimpleAsyncResult,
		domain GQuark,
		code Gint,
		format *Char,
		args Va_list)

	G_simple_async_result_is_valid func(
		result *GAsyncResult,
		source *GObject,
		source_tag Gpointer) Gboolean

	//TODO(t):Variant
	//g_simple_async_report_error_in_idle func(object  *GObject, callback  GAsyncReadyCallback, user_data  Gpointer, domain  GQuark, code  Gint, format  *Char, ...)

	G_simple_async_report_gerror_in_idle func(
		object *GObject,
		callback GAsyncReadyCallback,
		user_data Gpointer,
		err *GError)

	G_simple_async_report_take_gerror_in_idle func(
		object *GObject,
		callback GAsyncReadyCallback,
		user_data Gpointer,
		err *GError)

	G_simple_permission_get_type func() GType

	G_simple_permission_new func(
		allowed Gboolean) *GPermission

	G_socket_client_get_type func() GType

	G_socket_client_new func() *GSocketClient

	G_socket_client_get_family func(
		client *GSocketClient) GSocketFamily

	G_socket_client_set_family func(
		client *GSocketClient,
		family GSocketFamily)

	G_socket_client_get_socket_type func(
		client *GSocketClient) GSocketType

	G_socket_client_set_socket_type func(
		client *GSocketClient,
		t GSocketType)

	G_socket_client_get_protocol func(
		client *GSocketClient) GSocketProtocol

	G_socket_client_set_protocol func(
		client *GSocketClient,
		protocol GSocketProtocol)

	G_socket_client_get_local_address func(
		client *GSocketClient) *GSocketAddress

	G_socket_client_set_local_address func(
		client *GSocketClient,
		address *GSocketAddress)

	G_socket_client_get_timeout func(
		client *GSocketClient) Guint

	G_socket_client_set_timeout func(
		client *GSocketClient,
		timeout Guint)

	G_socket_client_get_enable_proxy func(
		client *GSocketClient) Gboolean

	G_socket_client_set_enable_proxy func(
		client *GSocketClient,
		enable Gboolean)

	G_socket_client_get_tls func(
		client *GSocketClient) Gboolean

	G_socket_client_set_tls func(
		client *GSocketClient,
		tls Gboolean)

	G_socket_client_get_tls_validation_flags func(
		client *GSocketClient) GTlsCertificateFlags

	G_socket_client_set_tls_validation_flags func(
		client *GSocketClient,
		flags GTlsCertificateFlags)

	G_socket_client_connect func(
		client *GSocketClient,
		connectable *GSocketConnectable,
		cancellable *GCancellable,
		err **GError) *GSocketConnection

	G_socket_client_connect_to_host func(
		client *GSocketClient,
		host_and_port *Gchar,
		default_port Guint16,
		cancellable *GCancellable,
		err **GError) *GSocketConnection

	G_socket_client_connect_to_service func(
		client *GSocketClient,
		domain *Gchar,
		service *Gchar,
		cancellable *GCancellable,
		err **GError) *GSocketConnection

	G_socket_client_connect_to_uri func(
		client *GSocketClient,
		uri *Gchar,
		default_port Guint16,
		cancellable *GCancellable,
		err **GError) *GSocketConnection

	G_socket_client_connect_async func(
		client *GSocketClient,
		connectable *GSocketConnectable,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_socket_client_connect_finish func(
		client *GSocketClient,
		result *GAsyncResult,
		err **GError) *GSocketConnection

	G_socket_client_connect_to_host_async func(
		client *GSocketClient,
		host_and_port *Gchar,
		default_port Guint16,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_socket_client_connect_to_host_finish func(
		client *GSocketClient,
		result *GAsyncResult,
		err **GError) *GSocketConnection

	G_socket_client_connect_to_service_async func(
		client *GSocketClient,
		domain *Gchar,
		service *Gchar,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_socket_client_connect_to_service_finish func(
		client *GSocketClient,
		result *GAsyncResult,
		err **GError) *GSocketConnection

	G_socket_client_connect_to_uri_async func(
		client *GSocketClient,
		uri *Gchar,
		default_port Guint16,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_socket_client_connect_to_uri_finish func(
		client *GSocketClient,
		result *GAsyncResult,
		err **GError) *GSocketConnection

	G_socket_client_add_application_proxy func(
		client *GSocketClient,
		protocol *Gchar)

	G_socket_connectable_get_type func() GType

	G_socket_connectable_enumerate func(
		connectable *GSocketConnectable) *GSocketAddressEnumerator

	G_socket_connectable_proxy_enumerate func(
		connectable *GSocketConnectable) *GSocketAddressEnumerator

	G_socket_get_type func() GType

	G_socket_new func(
		family GSocketFamily,
		t GSocketType,
		protocol GSocketProtocol,
		err **GError) *GSocket

	G_socket_new_from_fd func(
		fd Gint,
		err **GError) *GSocket

	G_socket_get_fd func(
		socket *GSocket) int

	G_socket_get_family func(
		socket *GSocket) GSocketFamily

	G_socket_get_socket_type func(
		socket *GSocket) GSocketType

	G_socket_get_protocol func(
		socket *GSocket) GSocketProtocol

	G_socket_get_local_address func(
		socket *GSocket,
		err **GError) *GSocketAddress

	G_socket_get_remote_address func(
		socket *GSocket,
		err **GError) *GSocketAddress

	G_socket_set_blocking func(
		socket *GSocket,
		blocking Gboolean)

	G_socket_get_blocking func(
		socket *GSocket) Gboolean

	G_socket_set_keepalive func(
		socket *GSocket,
		keepalive Gboolean)

	G_socket_get_keepalive func(
		socket *GSocket) Gboolean

	G_socket_get_listen_backlog func(
		socket *GSocket) Gint

	G_socket_set_listen_backlog func(
		socket *GSocket,
		backlog Gint)

	G_socket_get_timeout func(
		socket *GSocket) Guint

	G_socket_set_timeout func(
		socket *GSocket,
		timeout Guint)

	G_socket_is_connected func(
		socket *GSocket) Gboolean

	G_socket_bind func(
		socket *GSocket,
		address *GSocketAddress,
		allow_reuse Gboolean,
		err **GError) Gboolean

	G_socket_connect func(
		socket *GSocket,
		address *GSocketAddress,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_socket_check_connect_result func(
		socket *GSocket,
		err **GError) Gboolean

	G_socket_condition_check func(
		socket *GSocket,
		condition GIOCondition) GIOCondition

	G_socket_condition_wait func(
		socket *GSocket,
		condition GIOCondition,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_socket_accept func(
		socket *GSocket,
		cancellable *GCancellable,
		err **GError) *GSocket

	G_socket_listen func(
		socket *GSocket,
		err **GError) Gboolean

	G_socket_receive func(
		socket *GSocket,
		buffer *Gchar,
		size Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_socket_receive_from func(
		socket *GSocket,
		address **GSocketAddress,
		buffer *Gchar,
		size Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_socket_send func(
		socket *GSocket,
		buffer *Gchar,
		size Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_socket_send_to func(
		socket *GSocket,
		address *GSocketAddress,
		buffer *Gchar,
		size Gsize,
		cancellable *GCancellable,
		err **GError) Gssize

	G_socket_receive_message func(
		socket *GSocket,
		address **GSocketAddress,
		vectors *GInputVector,
		num_vectors Gint,
		messages ***GSocketControlMessage,
		num_messages *Gint,
		flags *Gint,
		cancellable *GCancellable,
		err **GError) Gssize

	G_socket_send_message func(
		socket *GSocket,
		address *GSocketAddress,
		vectors *GOutputVector,
		num_vectors Gint,
		messages **GSocketControlMessage,
		num_messages Gint,
		flags Gint,
		cancellable *GCancellable,
		err **GError) Gssize

	G_socket_close func(
		socket *GSocket,
		err **GError) Gboolean

	G_socket_shutdown func(
		socket *GSocket,
		shutdown_read Gboolean,
		shutdown_write Gboolean,
		err **GError) Gboolean

	G_socket_is_closed func(
		socket *GSocket) Gboolean

	G_socket_create_source func(
		socket *GSocket,
		condition GIOCondition,
		cancellable *GCancellable) *GSource

	G_socket_speaks_ipv4 func(
		socket *GSocket) Gboolean

	G_socket_get_credentials func(
		socket *GSocket,
		err **GError) *GCredentials

	G_socket_receive_with_blocking func(
		socket *GSocket,
		buffer *Gchar,
		size Gsize,
		blocking Gboolean,
		cancellable *GCancellable,
		err **GError) Gssize

	G_socket_send_with_blocking func(
		socket *GSocket,
		buffer *Gchar,
		size Gsize,
		blocking Gboolean,
		cancellable *GCancellable,
		err **GError) Gssize

	G_socket_connection_get_type func() GType

	G_socket_connection_get_socket func(
		connection *GSocketConnection) *GSocket

	G_socket_connection_get_local_address func(
		connection *GSocketConnection,
		err **GError) *GSocketAddress

	G_socket_connection_get_remote_address func(
		connection *GSocketConnection,
		err **GError) *GSocketAddress

	G_socket_connection_factory_register_type func(

		G_type GType,
		family GSocketFamily,
		t GSocketType,
		protocol Gint)

	G_socket_connection_factory_lookup_type func(
		family GSocketFamily,
		t GSocketType,
		protocol_id Gint) GType

	G_socket_connection_factory_create_connection func(
		socket *GSocket) *GSocketConnection

	G_socket_control_message_get_type func() GType

	G_socket_control_message_get_size func(
		message *GSocketControlMessage) Gsize

	G_socket_control_message_get_level func(
		message *GSocketControlMessage) int

	G_socket_control_message_get_msg_type func(
		message *GSocketControlMessage) int

	G_socket_control_message_serialize func(
		message *GSocketControlMessage,
		data Gpointer)

	G_socket_control_message_deserialize func(
		level int,
		typ int,
		size Gsize,
		data Gpointer) *GSocketControlMessage

	G_socket_listener_get_type func() GType

	G_socket_listener_new func() *GSocketListener

	G_socket_listener_set_backlog func(
		listener *GSocketListener,
		listen_backlog int)

	G_socket_listener_add_socket func(
		listener *GSocketListener,
		socket *GSocket,
		source_object *GObject,
		err **GError) Gboolean

	G_socket_listener_add_address func(
		listener *GSocketListener,
		address *GSocketAddress,
		t GSocketType,
		protocol GSocketProtocol,
		source_object *GObject,
		effective_address **GSocketAddress,
		err **GError) Gboolean

	G_socket_listener_add_inet_port func(
		listener *GSocketListener,
		port Guint16,
		source_object *GObject,
		err **GError) Gboolean

	G_socket_listener_add_any_inet_port func(
		listener *GSocketListener,
		source_object *GObject,
		err **GError) Guint16

	G_socket_listener_accept_socket func(
		listener *GSocketListener,
		source_object **GObject,
		cancellable *GCancellable,
		err **GError) *GSocket

	G_socket_listener_accept_socket_async func(
		listener *GSocketListener,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_socket_listener_accept_socket_finish func(
		listener *GSocketListener,
		result *GAsyncResult,
		source_object **GObject,
		err **GError) *GSocket

	G_socket_listener_accept func(
		listener *GSocketListener,
		source_object **GObject,
		cancellable *GCancellable,
		err **GError) *GSocketConnection

	G_socket_listener_accept_async func(
		listener *GSocketListener,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_socket_listener_accept_finish func(
		listener *GSocketListener,
		result *GAsyncResult,
		source_object **GObject,
		err **GError) *GSocketConnection

	G_socket_listener_close func(
		listener *GSocketListener)

	G_socket_service_get_type func() GType

	G_socket_service_new func() *GSocketService

	G_socket_service_start func(
		service *GSocketService)

	G_socket_service_stop func(
		service *GSocketService)

	G_socket_service_is_active func(
		service *GSocketService) Gboolean

	G_srv_target_get_type func() GType

	G_srv_target_new func(
		hostname *Gchar,
		port Guint16,
		priority Guint16,
		weight Guint16) *GSrvTarget

	G_srv_target_copy func(
		target *GSrvTarget) *GSrvTarget

	G_srv_target_free func(
		target *GSrvTarget)

	G_srv_target_get_hostname func(
		target *GSrvTarget) *Gchar

	G_srv_target_get_port func(
		target *GSrvTarget) Guint16

	G_srv_target_get_priority func(
		target *GSrvTarget) Guint16

	G_srv_target_get_weight func(
		target *GSrvTarget) Guint16

	G_srv_target_list_sort func(
		targets *GList) *GList

	G_tcp_connection_get_type func() GType

	G_tcp_connection_set_graceful_disconnect func(
		connection *GTcpConnection,
		graceful_disconnect Gboolean)

	G_tcp_connection_get_graceful_disconnect func(
		connection *GTcpConnection) Gboolean

	G_tcp_wrapper_connection_get_type func() GType

	G_tcp_wrapper_connection_new func(
		base_io_stream *GIOStream,
		socket *GSocket) *GSocketConnection

	G_tcp_wrapper_connection_get_base_io_stream func(
		conn *GTcpWrapperConnection) *GIOStream

	G_themed_icon_get_type func() GType

	G_themed_icon_new func(
		iconname *Char) *GIcon

	G_themed_icon_new_with_default_fallbacks func(
		iconname *Char) *GIcon

	G_themed_icon_new_from_names func(
		iconnames **Char,
		len int) *GIcon

	G_themed_icon_prepend_name func(
		icon *GThemedIcon,
		iconname *Char)

	G_themed_icon_append_name func(
		icon *GThemedIcon,
		iconname *Char)

	G_themed_icon_get_names func(
		icon *GThemedIcon) **Gchar

	G_threaded_socket_service_get_type func() GType

	G_threaded_socket_service_new func(
		max_threads int) *GSocketService

	G_tls_backend_get_type func() GType

	G_tls_backend_get_default func() *GTlsBackend

	G_tls_backend_supports_tls func(
		backend *GTlsBackend) Gboolean

	G_tls_backend_get_certificate_type func(
		backend *GTlsBackend) GType

	G_tls_backend_get_client_connection_type func(
		backend *GTlsBackend) GType

	G_tls_backend_get_server_connection_type func(
		backend *GTlsBackend) GType

	G_tls_certificate_get_type func() GType

	G_tls_certificate_new_from_pem func(
		data *Gchar,
		length Gssize,
		err **GError) *GTlsCertificate

	G_tls_certificate_new_from_file func(
		file *Gchar,
		err **GError) *GTlsCertificate

	G_tls_certificate_new_from_files func(
		cert_file *Gchar,
		key_file *Gchar,
		err **GError) *GTlsCertificate

	G_tls_certificate_list_new_from_file func(
		file *Gchar,
		err **GError) *GList

	G_tls_certificate_get_issuer func(
		cert *GTlsCertificate) *GTlsCertificate

	G_tls_certificate_verify func(
		cert *GTlsCertificate,
		identity *GSocketConnectable,
		trusted_ca *GTlsCertificate) GTlsCertificateFlags

	G_tls_connection_get_type func() GType

	G_tls_connection_set_use_system_certdb func(
		conn *GTlsConnection,
		use_system_certdb Gboolean)

	G_tls_connection_get_use_system_certdb func(
		conn *GTlsConnection) Gboolean

	G_tls_connection_set_certificate func(
		conn *GTlsConnection,
		certificate *GTlsCertificate)

	G_tls_connection_get_certificate func(
		conn *GTlsConnection) *GTlsCertificate

	G_tls_connection_get_peer_certificate func(
		conn *GTlsConnection) *GTlsCertificate

	G_tls_connection_get_peer_certificate_errors func(
		conn *GTlsConnection) GTlsCertificateFlags

	G_tls_connection_set_require_close_notify func(
		conn *GTlsConnection,
		require_close_notify Gboolean)

	G_tls_connection_get_require_close_notify func(
		conn *GTlsConnection) Gboolean

	G_tls_connection_set_rehandshake_mode func(
		conn *GTlsConnection,
		mode GTlsRehandshakeMode)

	G_tls_connection_get_rehandshake_mode func(
		conn *GTlsConnection) GTlsRehandshakeMode

	G_tls_connection_handshake func(
		conn *GTlsConnection,
		cancellable *GCancellable,
		err **GError) Gboolean

	G_tls_connection_handshake_async func(
		conn *GTlsConnection,
		io_priority int,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_tls_connection_handshake_finish func(
		conn *GTlsConnection,
		result *GAsyncResult,
		err **GError) Gboolean

	G_tls_error_quark func() GQuark

	G_tls_connection_emit_accept_certificate func(
		conn *GTlsConnection,
		peer_cert *GTlsCertificate,
		errors GTlsCertificateFlags) Gboolean

	G_tls_client_connection_get_type func() GType

	G_tls_client_connection_new func(
		base_io_stream *GIOStream,
		server_identity *GSocketConnectable,
		err **GError) *GIOStream

	G_tls_client_connection_get_validation_flags func(
		conn *GTlsClientConnection) GTlsCertificateFlags

	G_tls_client_connection_set_validation_flags func(
		conn *GTlsClientConnection,
		flags GTlsCertificateFlags)

	G_tls_client_connection_get_server_identity func(
		conn *GTlsClientConnection) *GSocketConnectable

	G_tls_client_connection_set_server_identity func(
		conn *GTlsClientConnection,
		identity *GSocketConnectable)

	G_tls_client_connection_get_use_ssl3 func(
		conn *GTlsClientConnection) Gboolean

	G_tls_client_connection_set_use_ssl3 func(
		conn *GTlsClientConnection,
		use_ssl3 Gboolean)

	G_tls_client_connection_get_accepted_cas func(
		conn *GTlsClientConnection) *GList

	G_tls_server_connection_get_type func() GType

	G_tls_server_connection_new func(
		base_io_stream *GIOStream,
		certificate *GTlsCertificate,
		err **GError) *GIOStream

	G_vfs_get_type func() GType

	G_vfs_is_active func(
		vfs *GVfs) Gboolean

	G_vfs_get_file_for_path func(
		vfs *GVfs,
		path *Char) *GFile

	G_vfs_get_file_for_uri func(
		vfs *GVfs,
		uri *Char) *GFile

	G_vfs_get_supported_uri_schemes func(
		vfs *GVfs) **Gchar

	G_vfs_parse_name func(
		vfs *GVfs,
		parse_name *Char) *GFile

	G_vfs_get_default func() *GVfs

	G_vfs_get_local func() *GVfs

	G_volume_get_type func() GType

	G_volume_get_name func(
		volume *GVolume) *Char

	G_volume_get_icon func(
		volume *GVolume) *GIcon

	G_volume_get_uuid func(
		volume *GVolume) *Char

	G_volume_get_drive func(
		volume *GVolume) *GDrive

	G_volume_get_mount func(
		volume *GVolume) *GMount

	G_volume_can_mount func(
		volume *GVolume) Gboolean

	G_volume_can_eject func(
		volume *GVolume) Gboolean

	G_volume_should_automount func(
		volume *GVolume) Gboolean

	G_volume_mount func(
		volume *GVolume,
		flags GMountMountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_volume_mount_finish func(
		volume *GVolume,
		result *GAsyncResult,
		err **GError) Gboolean

	G_volume_eject func(
		volume *GVolume,
		flags GMountUnmountFlags,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_volume_eject_finish func(
		volume *GVolume,
		result *GAsyncResult,
		err **GError) Gboolean

	G_volume_get_identifier func(
		volume *GVolume,
		kind *Char) *Char

	G_volume_enumerate_identifiers func(
		volume *GVolume) **Char

	G_volume_get_activation_root func(
		volume *GVolume) *GFile

	G_volume_eject_with_operation func(
		volume *GVolume,
		flags GMountUnmountFlags,
		mount_operation *GMountOperation,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	G_volume_eject_with_operation_finish func(
		volume *GVolume,
		result *GAsyncResult,
		err **GError) Gboolean

	G_zlib_compressor_get_type func() GType

	G_zlib_compressor_new func(
		format GZlibCompressorFormat,
		level int) *GZlibCompressor

	G_zlib_compressor_get_file_info func(
		compressor *GZlibCompressor) *GFileInfo

	G_zlib_compressor_set_file_info func(
		compressor *GZlibCompressor,
		file_info *GFileInfo)

	G_zlib_decompressor_get_type func() GType

	G_zlib_decompressor_new func(
		format GZlibCompressorFormat) *GZlibDecompressor

	G_zlib_decompressor_get_file_info func(
		decompressor *GZlibDecompressor) *GFileInfo

	G_win32_input_stream_get_type func() GType

	G_win32_input_stream_new func(
		handle *Void,
		close_handle Gboolean) *GInputStream

	G_win32_input_stream_set_close_handle func(
		stream *GWin32InputStream,
		close_handle Gboolean)

	G_win32_input_stream_get_close_handle func(
		stream *GWin32InputStream) Gboolean

	G_win32_input_stream_get_handle func(
		stream *GWin32InputStream) *Void

	G_win32_output_stream_get_type func() GType

	G_win32_output_stream_new func(
		handle *Void,
		close_handle Gboolean) *GOutputStream

	G_win32_output_stream_set_close_handle func(
		stream *GWin32OutputStream,
		close_handle Gboolean)

	G_win32_output_stream_get_close_handle func(
		stream *GWin32OutputStream) Gboolean

	G_win32_output_stream_get_handle func(
		stream *GWin32OutputStream) *Void

	G_settings_backend_get_type func() GType

	G_settings_backend_changed func(
		backend *GSettingsBackend,
		key *Gchar,
		origin_tag Gpointer)

	G_settings_backend_path_changed func(
		backend *GSettingsBackend,
		path *Gchar,
		origin_tag Gpointer)

	G_settings_backend_flatten_tree func(
		tree *GTree,
		path **Gchar,
		keys ***Gchar,
		values ***GVariant)

	G_settings_backend_keys_changed func(
		backend *GSettingsBackend,
		path *Gchar,
		items **Gchar,
		origin_tag Gpointer)

	G_settings_backend_path_writable_changed func(
		backend *GSettingsBackend,
		path *Gchar)

	G_settings_backend_writable_changed func(
		backend *GSettingsBackend,
		key *Gchar)

	G_settings_backend_changed_tree func(
		backend *GSettingsBackend,
		tree *GTree,
		origin_tag Gpointer)

	G_settings_backend_get_default func() *GSettingsBackend

	G_keyfile_settings_backend_new func(
		filename *Gchar,
		root_path *Gchar,
		root_group *Gchar) *GSettingsBackend

	G_null_settings_backend_new func() *GSettingsBackend

	G_memory_settings_backend_new func() *GSettingsBackend
)
var dll = "libgio-2.0-0.dll"

var apiList = outside.Apis{
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
	// TODO():Variant {"g_application_command_line_print", &G_application_command_line_print},
	// TODO():Variant {"g_application_command_line_printerr", &G_application_command_line_printerr},
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
	// TODO():Variant {"g_async_initable_new_async", &G_async_initable_new_async},
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
	// TODO():Variant {"g_dbus_error_set_dbus_error", &G_dbus_error_set_dbus_error},
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
	// TODO():Variant {"g_dbus_message_new_method_error", &G_dbus_message_new_method_error},
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
	// TODO():Variant {"g_dbus_method_invocation_return_error", &G_dbus_method_invocation_return_error},
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
	//TODO(t):Variant {"g_initable_new", &G_initable_new},
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
	//TODO(t):Variant {"g_settings_get", &G_settings_get},
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
	//TODO(t):Variant {"g_settings_set", &G_settings_set},
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
	//TODO(t):Variant {"g_simple_async_report_error_in_idle", &G_simple_async_report_error_in_idle},
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
	//TODO(t):Variant {"g_simple_async_result_new_error", &G_simple_async_result_new_error},
	{"g_simple_async_result_new_from_error", &G_simple_async_result_new_from_error},
	{"g_simple_async_result_new_take_error", &G_simple_async_result_new_take_error},
	{"g_simple_async_result_propagate_error", &G_simple_async_result_propagate_error},
	{"g_simple_async_result_run_in_thread", &G_simple_async_result_run_in_thread},
	//TODO(t):Variant {"g_simple_async_result_set_error", &G_simple_async_result_set_error},
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
