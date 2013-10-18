// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gdk provides API definitions for accessing
//libgdk-win32-2.0-0.dll and libgdk_pixbuf-2.0-0.dll.
package gdk

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllData(dll, false, dataList)
	outside.AddDllApis(dllPixbuf, false, apiListPixbuf)
	outside.AddDllData(dllPixbuf, false, dataListPixbuf)
}

type (
	HICON   uint32
	HGDIOBJ uint32
	HWND    uint32
	HDC     uint32
)

var (
	Gdk_colormap_get_type func() T.GType

	Gdk_colormap_new func(
		visual *T.GdkVisual,
		allocate T.Gboolean) *T.GdkColormap

	Gdk_colormap_ref func(
		cmap *T.GdkColormap) *T.GdkColormap

	Gdk_colormap_unref func(
		cmap *T.GdkColormap)

	Gdk_colormap_get_system func() *T.GdkColormap

	Gdk_colormap_get_screen func(
		cmap *T.GdkColormap) *T.GdkScreen

	Gdk_colormap_get_system_size func() T.Gint

	Gdk_colormap_change func(
		colormap *T.GdkColormap,
		ncolors T.Gint)

	Gdk_colormap_alloc_colors func(
		colormap *T.GdkColormap,
		colors *T.GdkColor,
		n_colors T.Gint,
		writeable T.Gboolean,
		best_match T.Gboolean,
		success *T.Gboolean) T.Gint

	Gdk_colormap_alloc_color func(
		colormap *T.GdkColormap,
		color *T.GdkColor,
		writeable T.Gboolean,
		best_match T.Gboolean) T.Gboolean

	Gdk_colormap_free_colors func(
		colormap *T.GdkColormap,
		colors *T.GdkColor,
		n_colors T.Gint)

	Gdk_colormap_query_color func(
		colormap *T.GdkColormap,
		pixel T.Gulong,
		result *T.GdkColor)

	Gdk_colormap_get_visual func(
		colormap *T.GdkColormap) *T.GdkVisual

	Gdk_color_copy func(
		color *T.GdkColor) *T.GdkColor

	Gdk_color_free func(
		color *T.GdkColor)

	Gdk_color_parse func(
		spec string,
		color *T.GdkColor) T.Gboolean

	Gdk_color_hash func(
		colora *T.GdkColor) T.Guint

	Gdk_color_equal func(
		colora *T.GdkColor,
		colorb *T.GdkColor) T.Gboolean

	Gdk_color_to_string func(
		color *T.GdkColor) string

	Gdk_color_get_type func() T.GType

	Gdk_colors_store func(
		colormap *T.GdkColormap,
		colors *T.GdkColor,
		ncolors T.Gint)

	Gdk_color_white func(
		colormap *T.GdkColormap,
		color *T.GdkColor) T.Gint

	Gdk_color_black func(
		colormap *T.GdkColormap,
		color *T.GdkColor) T.Gint

	Gdk_color_alloc func(
		colormap *T.GdkColormap,
		color *T.GdkColor) T.Gint

	Gdk_color_change func(
		colormap *T.GdkColormap,
		color *T.GdkColor) T.Gint

	Gdk_colors_alloc func(
		colormap *T.GdkColormap,
		contiguous T.Gboolean,
		planes *T.Gulong,
		nplanes T.Gint,
		pixels *T.Gulong,
		npixels T.Gint) T.Gint

	Gdk_colors_free func(
		colormap *T.GdkColormap,
		pixels *T.Gulong,
		npixels T.Gint,
		planes T.Gulong)

	Gdk_drag_context_get_type func() T.GType

	Gdk_drag_context_new func() *T.GdkDragContext

	Gdk_drag_context_list_targets func(
		context *T.GdkDragContext) *T.GList

	Gdk_drag_context_get_actions func(
		context *T.GdkDragContext) T.GdkDragAction

	Gdk_drag_context_get_suggested_action func(
		context *T.GdkDragContext) T.GdkDragAction

	Gdk_drag_context_get_selected_action func(
		context *T.GdkDragContext) T.GdkDragAction

	Gdk_drag_context_get_source_window func(
		context *T.GdkDragContext) *T.GdkWindow

	Gdk_drag_context_get_dest_window func(
		context *T.GdkDragContext) *T.GdkWindow

	Gdk_drag_context_get_protocol func(
		context *T.GdkDragContext) T.GdkDragProtocol

	Gdk_drag_context_ref func(
		context *T.GdkDragContext)

	Gdk_drag_context_unref func(
		context *T.GdkDragContext)

	Gdk_drag_status func(
		context *T.GdkDragContext,
		action T.GdkDragAction,
		time_ T.Guint32)

	Gdk_drop_reply func(
		context *T.GdkDragContext,
		ok T.Gboolean,
		time_ T.Guint32)

	Gdk_drop_finish func(
		context *T.GdkDragContext,
		success T.Gboolean,
		time_ T.Guint32)

	Gdk_drag_get_selection func(
		context *T.GdkDragContext) T.GdkAtom

	Gdk_drag_begin func(
		window *T.GdkWindow,
		targets *T.GList) *T.GdkDragContext

	Gdk_drag_get_protocol_for_display func(
		display *T.GdkDisplay,
		xid T.GdkNativeWindow,
		protocol *T.GdkDragProtocol) T.GdkNativeWindow

	Gdk_drag_find_window_for_screen func(
		context *T.GdkDragContext,
		drag_window *T.GdkWindow,
		screen *T.GdkScreen,
		x_root T.Gint,
		y_root T.Gint,
		dest_window **T.GdkWindow,
		protocol *T.GdkDragProtocol)

	Gdk_drag_get_protocol func(
		xid T.GdkNativeWindow,
		protocol *T.GdkDragProtocol) T.GdkNativeWindow

	Gdk_drag_find_window func(
		context *T.GdkDragContext,
		drag_window *T.GdkWindow,
		x_root T.Gint,
		y_root T.Gint,
		dest_window **T.GdkWindow,
		protocol *T.GdkDragProtocol)

	Gdk_drag_motion func(
		context *T.GdkDragContext,
		dest_window *T.GdkWindow,
		protocol T.GdkDragProtocol,
		x_root T.Gint,
		y_root T.Gint,
		suggested_action T.GdkDragAction,
		possible_actions T.GdkDragAction,
		time_ T.Guint32) T.Gboolean

	Gdk_drag_drop func(
		context *T.GdkDragContext,
		time_ T.Guint32)

	Gdk_drag_abort func(
		context *T.GdkDragContext,
		time_ T.Guint32)

	Gdk_drag_drop_succeeded func(
		context *T.GdkDragContext) T.Gboolean

	Gdk_device_get_type func() T.GType

	Gdk_devices_list func() *T.GList

	Gdk_device_get_name func(
		device *T.GdkDevice) string

	Gdk_device_get_source func(
		device *T.GdkDevice) T.GdkInputSource

	Gdk_device_get_mode func(
		device *T.GdkDevice) T.GdkInputMode

	Gdk_device_get_has_cursor func(
		device *T.GdkDevice) T.Gboolean

	Gdk_device_get_key func(
		device *T.GdkDevice,
		index T.Guint,
		keyval *T.Guint,
		modifiers *T.GdkModifierType)

	Gdk_device_get_axis_use func(
		device *T.GdkDevice,
		index T.Guint) T.GdkAxisUse

	Gdk_device_get_n_keys func(
		device *T.GdkDevice) T.Gint

	Gdk_device_get_n_axes func(
		device *T.GdkDevice) T.Gint

	Gdk_device_set_source func(
		device *T.GdkDevice,
		source T.GdkInputSource)

	Gdk_device_set_mode func(
		device *T.GdkDevice,
		mode T.GdkInputMode) T.Gboolean

	Gdk_device_set_key func(
		device *T.GdkDevice,
		index_ T.Guint,
		keyval T.Guint,
		modifiers T.GdkModifierType)

	Gdk_device_set_axis_use func(
		device *T.GdkDevice,
		index_ T.Guint,
		use T.GdkAxisUse)

	Gdk_device_get_state func(
		device *T.GdkDevice,
		window *T.GdkWindow,
		axes *T.Gdouble,
		mask *T.GdkModifierType)

	Gdk_device_get_history func(
		device *T.GdkDevice,
		window *T.GdkWindow,
		start T.Guint32,
		stop T.Guint32,
		events ***T.GdkTimeCoord,
		n_events *T.Gint) T.Gboolean

	Gdk_device_free_history func(
		events **T.GdkTimeCoord,
		n_events T.Gint)

	Gdk_device_get_axis func(
		device *T.GdkDevice,
		axes *T.Gdouble,
		use T.GdkAxisUse,
		value *T.Gdouble) T.Gboolean

	Gdk_input_set_extension_events func(
		window *T.GdkWindow,
		mask T.Gint,
		mode T.GdkExtensionMode)

	Gdk_device_get_core_pointer func() *T.GdkDevice

	Gdk_event_get_type func() T.GType

	Gdk_events_pending func() T.Gboolean

	Gdk_event_get func() *T.GdkEvent

	Gdk_event_peek func() *T.GdkEvent

	Gdk_event_get_graphics_expose func(
		window *T.GdkWindow) *T.GdkEvent

	Gdk_event_put func(
		event *T.GdkEvent)

	Gdk_event_new func(
		t T.GdkEventType) *T.GdkEvent

	Gdk_event_copy func(
		event *T.GdkEvent) *T.GdkEvent

	Gdk_event_free func(
		event *T.GdkEvent)

	Gdk_event_get_time func(
		event *T.GdkEvent) T.Guint32

	Gdk_event_get_state func(
		event *T.GdkEvent,
		state *T.GdkModifierType) T.Gboolean

	Gdk_event_get_coords func(
		event *T.GdkEvent,
		x_win *T.Gdouble,
		y_win *T.Gdouble) T.Gboolean

	Gdk_event_get_root_coords func(
		event *T.GdkEvent,
		x_root *T.Gdouble,
		y_root *T.Gdouble) T.Gboolean

	Gdk_event_get_axis func(
		event *T.GdkEvent,
		axis_use T.GdkAxisUse,
		value *T.Gdouble) T.Gboolean

	Gdk_event_request_motions func(
		event *T.GdkEventMotion)

	Gdk_event_handler_set func(
		f T.GdkEventFunc,
		data T.Gpointer,
		notify T.GDestroyNotify)

	Gdk_event_set_screen func(
		event *T.GdkEvent,
		screen *T.GdkScreen)

	Gdk_event_get_screen func(
		event *T.GdkEvent) *T.GdkScreen

	Gdk_set_show_events func(
		show_events T.Gboolean)

	Gdk_get_show_events func() T.Gboolean

	Gdk_add_client_message_filter func(
		message_type T.GdkAtom,
		f T.GdkFilterFunc,
		data T.Gpointer)

	Gdk_setting_get func(
		name string,
		value *T.GValue) T.Gboolean

	Gdk_display_get_type func() T.GType

	Gdk_display_open func(
		display_name string) *T.GdkDisplay

	Gdk_display_get_name func(
		display *T.GdkDisplay) string

	Gdk_display_get_n_screens func(
		display *T.GdkDisplay) T.Gint

	Gdk_display_get_screen func(
		display *T.GdkDisplay,
		screen_num T.Gint) *T.GdkScreen

	Gdk_display_get_default_screen func(
		display *T.GdkDisplay) *T.GdkScreen

	Gdk_display_pointer_ungrab func(
		display *T.GdkDisplay,
		time_ T.Guint32)

	Gdk_display_keyboard_ungrab func(
		display *T.GdkDisplay,
		time_ T.Guint32)

	Gdk_display_pointer_is_grabbed func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_display_beep func(
		display *T.GdkDisplay)

	Gdk_display_sync func(
		display *T.GdkDisplay)

	Gdk_display_flush func(
		display *T.GdkDisplay)

	Gdk_display_close func(
		display *T.GdkDisplay)

	Gdk_display_is_closed func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_display_list_devices func(
		display *T.GdkDisplay) *T.GList

	Gdk_display_get_event func(
		display *T.GdkDisplay) *T.GdkEvent

	Gdk_display_peek_event func(
		display *T.GdkDisplay) *T.GdkEvent

	Gdk_display_put_event func(
		display *T.GdkDisplay,
		event *T.GdkEvent)

	Gdk_display_add_client_message_filter func(
		display *T.GdkDisplay,
		message_type T.GdkAtom,
		f T.GdkFilterFunc,
		data T.Gpointer)

	Gdk_display_set_double_click_time func(
		display *T.GdkDisplay,
		msec T.Guint)

	Gdk_display_set_double_click_distance func(
		display *T.GdkDisplay,
		distance T.Guint)

	Gdk_display_get_default func() *T.GdkDisplay

	Gdk_display_get_core_pointer func(
		display *T.GdkDisplay) *T.GdkDevice

	Gdk_display_get_pointer func(
		display *T.GdkDisplay,
		screen **T.GdkScreen,
		x *T.Gint,
		y *T.Gint,
		mask *T.GdkModifierType)

	Gdk_display_get_window_at_pointer func(
		display *T.GdkDisplay,
		win_x *T.Gint,
		win_y *T.Gint) *T.GdkWindow

	Gdk_display_warp_pointer func(
		display *T.GdkDisplay,
		screen *T.GdkScreen,
		x T.Gint,
		y T.Gint)

	Gdk_display_set_pointer_hooks func(
		display *T.GdkDisplay,
		new_hooks *T.GdkDisplayPointerHooks) *T.GdkDisplayPointerHooks

	Gdk_display_open_default_libgtk_only func() *T.GdkDisplay

	Gdk_display_supports_cursor_alpha func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_display_supports_cursor_color func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_display_get_default_cursor_size func(
		display *T.GdkDisplay) T.Guint

	Gdk_display_get_maximal_cursor_size func(
		display *T.GdkDisplay,
		width *T.Guint,
		height *T.Guint)

	Gdk_display_get_default_group func(
		display *T.GdkDisplay) *T.GdkWindow

	Gdk_display_supports_selection_notification func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_display_request_selection_notification func(
		display *T.GdkDisplay,
		selection T.GdkAtom) T.Gboolean

	Gdk_display_supports_clipboard_persistence func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_display_store_clipboard func(
		display *T.GdkDisplay,
		clipboard_window *T.GdkWindow,
		time_ T.Guint32,
		targets *T.GdkAtom,
		n_targets T.Gint)

	Gdk_display_supports_shapes func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_display_supports_input_shapes func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_display_supports_composite func(
		display *T.GdkDisplay) T.Gboolean

	Gdk_screen_get_type func() T.GType

	Gdk_screen_get_default_colormap func(
		screen *T.GdkScreen) *T.GdkColormap

	Gdk_screen_set_default_colormap func(
		screen *T.GdkScreen,
		colormap *T.GdkColormap)

	Gdk_screen_get_system_colormap func(
		screen *T.GdkScreen) *T.GdkColormap

	Gdk_screen_get_system_visual func(
		screen *T.GdkScreen) *T.GdkVisual

	Gdk_screen_get_rgb_colormap func(
		screen *T.GdkScreen) *T.GdkColormap

	Gdk_screen_get_rgb_visual func(
		screen *T.GdkScreen) *T.GdkVisual

	Gdk_screen_get_rgba_colormap func(
		screen *T.GdkScreen) *T.GdkColormap

	Gdk_screen_get_rgba_visual func(
		screen *T.GdkScreen) *T.GdkVisual

	Gdk_screen_is_composited func(
		screen *T.GdkScreen) T.Gboolean

	Gdk_screen_get_root_window func(
		screen *T.GdkScreen) *T.GdkWindow

	Gdk_screen_get_display func(
		screen *T.GdkScreen) *T.GdkDisplay

	Gdk_screen_get_number func(
		screen *T.GdkScreen) T.Gint

	Gdk_screen_get_width func(
		screen *T.GdkScreen) T.Gint

	Gdk_screen_get_height func(
		screen *T.GdkScreen) T.Gint

	Gdk_screen_get_width_mm func(
		screen *T.GdkScreen) T.Gint

	Gdk_screen_get_height_mm func(
		screen *T.GdkScreen) T.Gint

	Gdk_screen_list_visuals func(
		screen *T.GdkScreen) *T.GList

	Gdk_screen_get_toplevel_windows func(
		screen *T.GdkScreen) *T.GList

	Gdk_screen_make_display_name func(
		screen *T.GdkScreen) string

	Gdk_screen_get_n_monitors func(
		screen *T.GdkScreen) T.Gint

	Gdk_screen_get_primary_monitor func(
		screen *T.GdkScreen) T.Gint

	Gdk_screen_get_monitor_geometry func(
		screen *T.GdkScreen,
		monitor_num T.Gint,
		dest *T.GdkRectangle)

	Gdk_screen_get_monitor_at_point func(
		screen *T.GdkScreen,
		x T.Gint,
		y T.Gint) T.Gint

	Gdk_screen_get_monitor_at_window func(
		screen *T.GdkScreen,
		window *T.GdkWindow) T.Gint

	Gdk_screen_get_monitor_width_mm func(
		screen *T.GdkScreen,
		monitor_num T.Gint) T.Gint

	Gdk_screen_get_monitor_height_mm func(
		screen *T.GdkScreen,
		monitor_num T.Gint) T.Gint

	Gdk_screen_get_monitor_plug_name func(
		screen *T.GdkScreen,
		monitor_num T.Gint) string

	Gdk_screen_broadcast_client_message func(
		screen *T.GdkScreen,
		event *T.GdkEvent)

	Gdk_screen_get_default func() *T.GdkScreen

	Gdk_screen_get_setting func(
		screen *T.GdkScreen,
		name string,
		value *T.GValue) T.Gboolean

	Gdk_screen_set_font_options func(
		screen *T.GdkScreen,
		options *T.Cairo_font_options_t)

	Gdk_screen_get_font_options func(
		screen *T.GdkScreen) *T.Cairo_font_options_t

	Gdk_screen_set_resolution func(
		screen *T.GdkScreen,
		dpi T.Gdouble)

	Gdk_screen_get_resolution func(
		screen *T.GdkScreen) T.Gdouble

	Gdk_screen_get_active_window func(
		screen *T.GdkScreen) *T.GdkWindow

	Gdk_screen_get_window_stack func(
		screen *T.GdkScreen) *T.GList

	Gdk_app_launch_context_get_type func() T.GType

	Gdk_app_launch_context_new func() *T.GdkAppLaunchContext

	Gdk_app_launch_context_set_display func(
		context *T.GdkAppLaunchContext,
		display *T.GdkDisplay)

	Gdk_app_launch_context_set_screen func(
		context *T.GdkAppLaunchContext,
		screen *T.GdkScreen)

	Gdk_app_launch_context_set_desktop func(
		context *T.GdkAppLaunchContext,
		desktop T.Gint)

	Gdk_app_launch_context_set_timestamp func(
		context *T.GdkAppLaunchContext,
		timestamp T.Guint32)

	Gdk_app_launch_context_set_icon func(
		context *T.GdkAppLaunchContext,
		icon *T.GIcon)

	Gdk_app_launch_context_set_icon_name func(
		context *T.GdkAppLaunchContext,
		icon_name string)

	Gdk_rgb_init func()

	Gdk_rgb_xpixel_from_rgb func(
		rgb T.Guint32) T.Gulong

	Gdk_rgb_gc_set_foreground func(
		gc *T.GdkGC,
		rgb T.Guint32)

	Gdk_rgb_gc_set_background func(
		gc *T.GdkGC,
		rgb T.Guint32)

	Gdk_rgb_find_color func(
		colormap *T.GdkColormap,
		color *T.GdkColor)

	Gdk_draw_rgb_image func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		dith T.GdkRgbDither,
		rgb_buf *T.Guchar,
		rowstride T.Gint)

	Gdk_draw_rgb_image_dithalign func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		dith T.GdkRgbDither,
		rgb_buf *T.Guchar,
		rowstride T.Gint,
		xdith T.Gint,
		ydith T.Gint)

	Gdk_draw_rgb_32_image func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		dith T.GdkRgbDither,
		buf *T.Guchar,
		rowstride T.Gint)

	Gdk_draw_rgb_32_image_dithalign func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		dith T.GdkRgbDither,
		buf *T.Guchar,
		rowstride T.Gint,
		xdith T.Gint,
		ydith T.Gint)

	Gdk_draw_gray_image func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		dith T.GdkRgbDither,
		buf *T.Guchar,
		rowstride T.Gint)

	Gdk_draw_indexed_image func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		dith T.GdkRgbDither,
		buf *T.Guchar,
		rowstride T.Gint,
		cmap *T.GdkRgbCmap)

	Gdk_rgb_cmap_new func(
		colors *T.Guint32,
		n_colors T.Gint) *T.GdkRgbCmap

	Gdk_rgb_cmap_free func(
		cmap *T.GdkRgbCmap)

	Gdk_rgb_set_verbose func(
		verbose T.Gboolean)

	Gdk_rgb_set_install func(
		install T.Gboolean)

	Gdk_rgb_set_min_colors func(
		min_colors T.Gint)

	Gdk_rgb_get_colormap func() *T.GdkColormap

	Gdk_rgb_get_visual func() *T.GdkVisual

	Gdk_rgb_ditherable func() T.Gboolean

	Gdk_rgb_colormap_ditherable func(
		cmap *T.GdkColormap) T.Gboolean

	Gdk_pixbuf_error_quark func() T.GQuark

	Gdk_pixbuf_get_type func() T.GType

	Gdk_pixbuf_ref func(
		pixbuf *T.GdkPixbuf) *T.GdkPixbuf

	Gdk_pixbuf_unref func(
		pixbuf *T.GdkPixbuf)

	Gdk_pixbuf_get_colorspace func(
		pixbuf *T.GdkPixbuf) T.GdkColorspace

	Gdk_pixbuf_get_n_channels func(
		pixbuf *T.GdkPixbuf) int

	Gdk_pixbuf_get_has_alpha func(
		pixbuf *T.GdkPixbuf) T.Gboolean

	Gdk_pixbuf_get_bits_per_sample func(
		pixbuf *T.GdkPixbuf) int

	Gdk_pixbuf_get_pixels func(
		pixbuf *T.GdkPixbuf) *T.Guchar

	Gdk_pixbuf_get_width func(
		pixbuf *T.GdkPixbuf) int

	Gdk_pixbuf_get_height func(
		pixbuf *T.GdkPixbuf) int

	Gdk_pixbuf_get_rowstride func(
		pixbuf *T.GdkPixbuf) int

	Gdk_pixbuf_new func(
		colorspace T.GdkColorspace,
		has_alpha T.Gboolean,
		bits_per_sample int,
		width int,
		height int) *T.GdkPixbuf

	Gdk_pixbuf_copy func(
		pixbuf *T.GdkPixbuf) *T.GdkPixbuf

	Gdk_pixbuf_new_subpixbuf func(
		src_pixbuf *T.GdkPixbuf,
		src_x int,
		src_y int,
		width int,
		height int) *T.GdkPixbuf

	Gdk_pixbuf_new_from_file func(
		filename string,
		e **T.GError) *T.GdkPixbuf

	Gdk_pixbuf_new_from_file_at_size func(
		filename string,
		width int,
		height int,
		e **T.GError) *T.GdkPixbuf

	Gdk_pixbuf_new_from_file_at_scale func(
		filename string,
		width int,
		height int,
		preserve_aspect_ratio T.Gboolean,
		e **T.GError) *T.GdkPixbuf

	Gdk_pixbuf_new_from_data func(
		data *T.Guchar,
		colorspace T.GdkColorspace,
		has_alpha T.Gboolean,
		bits_per_sample int,
		width int,
		height int,
		rowstride int,
		destroy_fn T.GdkPixbufDestroyNotify,
		destroy_fn_data T.Gpointer) *T.GdkPixbuf

	Gdk_pixbuf_new_from_xpm_data func(
		data **T.Char) *T.GdkPixbuf

	Gdk_pixbuf_new_from_inline func(
		data_length T.Gint,
		data *T.Guint8,
		copy_pixels T.Gboolean,
		e **T.GError) *T.GdkPixbuf

	Gdk_pixbuf_fill func(
		pixbuf *T.GdkPixbuf,
		pixel T.Guint32)

	Gdk_pixbuf_save_utf8 func(pixbuf *T.GdkPixbuf,
		filename, typ string, e **T.GError, v ...VArg) T.Gboolean

	Gdk_pixbuf_savev_utf8 func(
		pixbuf *T.GdkPixbuf,
		filename string,
		typ string,
		option_keys **T.Char,
		option_values **T.Char,
		e **T.GError) T.Gboolean

	Gdk_pixbuf_save_to_callback func(pixbuf *T.GdkPixbuf,
		save_func T.GdkPixbufSaveFunc,
		user_data T.Gpointer,
		typ string, err **T.GError, v ...VArg) T.Gboolean

	Gdk_pixbuf_save_to_callbackv func(
		pixbuf *T.GdkPixbuf,
		save_func T.GdkPixbufSaveFunc,
		user_data T.Gpointer,
		typ string,
		option_keys **T.Char,
		option_values **T.Char,
		e **T.GError) T.Gboolean

	Gdk_pixbuf_save_to_buffer func(pixbuf *T.GdkPixbuf,
		buffer **T.Gchar, buffer_size *T.Gsize,
		typ string, e **T.GError, v ...VArg) T.Gboolean

	Gdk_pixbuf_save_to_bufferv func(
		pixbuf *T.GdkPixbuf,
		buffer **T.Gchar,
		buffer_size *T.Gsize,
		typ string,
		option_keys **T.Char,
		option_values **T.Char,
		e **T.GError) T.Gboolean

	Gdk_pixbuf_new_from_stream func(
		stream *T.GInputStream,
		cancellable *T.GCancellable,
		e **T.GError) *T.GdkPixbuf

	Gdk_pixbuf_new_from_stream_async func(
		stream *T.GInputStream,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	Gdk_pixbuf_new_from_stream_finish func(
		async_result *T.GAsyncResult,
		e **T.GError) *T.GdkPixbuf

	Gdk_pixbuf_new_from_stream_at_scale func(
		stream *T.GInputStream,
		width T.Gint,
		height T.Gint,
		preserve_aspect_ratio T.Gboolean,
		cancellable *T.GCancellable,
		e **T.GError) *T.GdkPixbuf

	Gdk_pixbuf_new_from_stream_at_scale_async func(
		stream *T.GInputStream,
		width T.Gint,
		height T.Gint,
		preserve_aspect_ratio T.Gboolean,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		user_data T.Gpointer)

	Gdk_pixbuf_save_to_stream func(pixbuf *T.GdkPixbuf,
		stream *T.GOutputStream, typ string,
		cancellable *T.GCancellable, e **T.GError, v ...VArg) T.Gboolean

	Gdk_pixbuf_save_to_stream_async func(pixbuf *T.GdkPixbuf,
		stream *T.GOutputStream, typ string,
		cancellable *T.GCancellable, callback T.GAsyncReadyCallback,
		user_data T.Gpointer, v ...VArg)

	Gdk_pixbuf_save_to_stream_finish func(
		async_result *T.GAsyncResult,
		e **T.GError) T.Gboolean

	Gdk_pixbuf_add_alpha func(
		pixbuf *T.GdkPixbuf,
		substitute_color T.Gboolean,
		r, g, b T.Guchar) *T.GdkPixbuf

	Gdk_pixbuf_copy_area func(
		src_pixbuf *T.GdkPixbuf,
		src_x, src_y, width, height int,
		dest_pixbuf *T.GdkPixbuf,
		dest_x, dest_y int)

	Gdk_pixbuf_saturate_and_pixelate func(
		src, dest *T.GdkPixbuf,
		saturation T.Gfloat,
		pixelate T.Gboolean)

	Gdk_pixbuf_apply_embedded_orientation func(
		src *T.GdkPixbuf) *T.GdkPixbuf

	Gdk_pixbuf_get_option func(
		pixbuf *T.GdkPixbuf,
		key string) string

	Gdk_pixbuf_scale func(
		src *T.GdkPixbuf,
		dest *T.GdkPixbuf,
		dest_x int,
		dest_y int,
		dest_width int,
		dest_height int,
		offset_x T.Double,
		offset_y T.Double,
		scale_x T.Double,
		scale_y T.Double,
		interp_type T.GdkInterpType)

	Gdk_pixbuf_composite func(
		src *T.GdkPixbuf,
		dest *T.GdkPixbuf,
		dest_x int,
		dest_y int,
		dest_width int,
		dest_height int,
		offset_x T.Double,
		offset_y T.Double,
		scale_x T.Double,
		scale_y T.Double,
		interp_type T.GdkInterpType,
		overall_alpha int)

	Gdk_pixbuf_composite_color func(
		src *T.GdkPixbuf,
		dest *T.GdkPixbuf,
		dest_x int,
		dest_y int,
		dest_width int,
		dest_height int,
		offset_x T.Double,
		offset_y T.Double,
		scale_x T.Double,
		scale_y T.Double,
		interp_type T.GdkInterpType,
		overall_alpha int,
		check_x int,
		check_y int,
		check_size int,
		color1 T.Guint32,
		color2 T.Guint32)

	Gdk_pixbuf_scale_simple func(
		src *T.GdkPixbuf,
		dest_width int,
		dest_height int,
		interp_type T.GdkInterpType) *T.GdkPixbuf

	Gdk_pixbuf_composite_color_simple func(
		src *T.GdkPixbuf,
		dest_width int,
		dest_height int,
		interp_type T.GdkInterpType,
		overall_alpha int,
		check_size int,
		color1 T.Guint32,
		color2 T.Guint32) *T.GdkPixbuf

	Gdk_pixbuf_rotate_simple func(
		src *T.GdkPixbuf,
		angle T.GdkPixbufRotation) *T.GdkPixbuf

	Gdk_pixbuf_flip func(
		src *T.GdkPixbuf,
		horizontal T.Gboolean) *T.GdkPixbuf

	Gdk_pixbuf_animation_get_type func() T.GType

	Gdk_pixbuf_animation_new_from_file_utf8 func(
		filename string,
		e **T.GError) *T.GdkPixbufAnimation

	Gdk_pixbuf_animation_ref func(
		animation *T.GdkPixbufAnimation) *T.GdkPixbufAnimation

	Gdk_pixbuf_animation_unref func(
		animation *T.GdkPixbufAnimation)

	Gdk_pixbuf_animation_get_width func(
		animation *T.GdkPixbufAnimation) int

	Gdk_pixbuf_animation_get_height func(
		animation *T.GdkPixbufAnimation) int

	Gdk_pixbuf_animation_is_static_image func(
		animation *T.GdkPixbufAnimation) T.Gboolean

	Gdk_pixbuf_animation_get_static_image func(
		animation *T.GdkPixbufAnimation) *T.GdkPixbuf

	Gdk_pixbuf_animation_get_iter func(
		animation *T.GdkPixbufAnimation,
		start_time *T.GTimeVal) *T.GdkPixbufAnimationIter

	Gdk_pixbuf_animation_iter_get_type func() T.GType

	Gdk_pixbuf_animation_iter_get_delay_time func(
		iter *T.GdkPixbufAnimationIter) int

	Gdk_pixbuf_animation_iter_get_pixbuf func(
		iter *T.GdkPixbufAnimationIter) *T.GdkPixbuf

	Gdk_pixbuf_animation_iter_on_currently_loading_frame func(
		iter *T.GdkPixbufAnimationIter) T.Gboolean

	Gdk_pixbuf_animation_iter_advance func(
		iter *T.GdkPixbufAnimationIter,
		current_time *T.GTimeVal) T.Gboolean

	Gdk_pixbuf_simple_anim_get_type func() T.GType

	Gdk_pixbuf_simple_anim_iter_get_type func() T.GType

	Gdk_pixbuf_simple_anim_new func(
		width T.Gint,
		height T.Gint,
		rate T.Gfloat) *T.GdkPixbufSimpleAnim

	Gdk_pixbuf_simple_anim_add_frame func(
		animation *T.GdkPixbufSimpleAnim,
		pixbuf *T.GdkPixbuf)

	Gdk_pixbuf_simple_anim_set_loop func(
		animation *T.GdkPixbufSimpleAnim,
		loop T.Gboolean)

	Gdk_pixbuf_simple_anim_get_loop func(
		animation *T.GdkPixbufSimpleAnim) T.Gboolean

	Gdk_pixbuf_format_get_type func() T.GType

	Gdk_pixbuf_get_formats func() *T.GSList

	Gdk_pixbuf_format_get_name func(
		format *T.GdkPixbufFormat) string

	Gdk_pixbuf_format_get_description func(
		format *T.GdkPixbufFormat) string

	Gdk_pixbuf_format_get_mime_types func(
		format *T.GdkPixbufFormat) **T.Gchar

	Gdk_pixbuf_format_get_extensions func(
		format *T.GdkPixbufFormat) **T.Gchar

	Gdk_pixbuf_format_is_writable func(
		format *T.GdkPixbufFormat) T.Gboolean

	Gdk_pixbuf_format_is_scalable func(
		format *T.GdkPixbufFormat) T.Gboolean

	Gdk_pixbuf_format_is_disabled func(
		format *T.GdkPixbufFormat) T.Gboolean

	Gdk_pixbuf_format_set_disabled func(
		format *T.GdkPixbufFormat,
		disabled T.Gboolean)

	Gdk_pixbuf_format_get_license func(
		format *T.GdkPixbufFormat) string

	Gdk_pixbuf_get_file_info func(
		filename string,
		width *T.Gint,
		height *T.Gint) *T.GdkPixbufFormat

	Gdk_pixbuf_format_copy func(
		format *T.GdkPixbufFormat) *T.GdkPixbufFormat

	Gdk_pixbuf_format_free func(
		format *T.GdkPixbufFormat)

	Gdk_pixbuf_loader_get_type func() T.GType

	Gdk_pixbuf_loader_new func() *T.GdkPixbufLoader

	Gdk_pixbuf_loader_new_with_type func(
		image_type string,
		e **T.GError) *T.GdkPixbufLoader

	Gdk_pixbuf_loader_new_with_mime_type func(
		mime_type string,
		e **T.GError) *T.GdkPixbufLoader

	Gdk_pixbuf_loader_set_size func(
		loader *T.GdkPixbufLoader,
		width int,
		height int)

	Gdk_pixbuf_loader_write func(
		loader *T.GdkPixbufLoader,
		buf *T.Guchar,
		count T.Gsize,
		e **T.GError) T.Gboolean

	Gdk_pixbuf_loader_get_pixbuf func(
		loader *T.GdkPixbufLoader) *T.GdkPixbuf

	Gdk_pixbuf_loader_get_animation func(
		loader *T.GdkPixbufLoader) *T.GdkPixbufAnimation

	Gdk_pixbuf_loader_close func(
		loader *T.GdkPixbufLoader,
		e **T.GError) T.Gboolean

	Gdk_pixbuf_loader_get_format func(
		loader *T.GdkPixbufLoader) *T.GdkPixbufFormat

	Gdk_pixbuf_alpha_mode_get_type func() T.GType

	Gdk_colorspace_get_type func() T.GType

	Gdk_pixbuf_error_get_type func() T.GType

	Gdk_interp_type_get_type func() T.GType

	Gdk_pixbuf_rotation_get_type func() T.GType

	Gdk_pixbuf_render_threshold_alpha func(
		pixbuf *T.GdkPixbuf,
		bitmap *T.GdkBitmap,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int,
		alpha_threshold int)

	Gdk_pixbuf_render_to_drawable func(
		pixbuf *T.GdkPixbuf,
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int,
		dither T.GdkRgbDither,
		x_dither int,
		y_dither int)

	Gdk_pixbuf_render_to_drawable_alpha func(
		pixbuf *T.GdkPixbuf,
		drawable *T.GdkDrawable,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int,
		alpha_mode T.GdkPixbufAlphaMode,
		alpha_threshold int,
		dither T.GdkRgbDither,
		x_dither int,
		y_dither int)

	Gdk_pixbuf_render_pixmap_and_mask_for_colormap func(
		pixbuf *T.GdkPixbuf,
		colormap *T.GdkColormap,
		pixmap_return **T.GdkPixmap,
		mask_return **T.GdkBitmap,
		alpha_threshold int)

	Gdk_pixbuf_render_pixmap_and_mask func(
		pixbuf *T.GdkPixbuf,
		pixmap_return **T.GdkPixmap,
		mask_return **T.GdkBitmap,
		alpha_threshold int)

	Gdk_pixbuf_get_from_drawable func(
		dest *T.GdkPixbuf,
		src *T.GdkDrawable,
		cmap *T.GdkColormap,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int) *T.GdkPixbuf

	Gdk_pixbuf_get_from_image func(
		dest *T.GdkPixbuf,
		src *T.GdkImage,
		cmap *T.GdkColormap,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int) *T.GdkPixbuf

	Gdk_cairo_create func(
		drawable *T.GdkDrawable) *T.Cairo_t

	Gdk_cairo_reset_clip func(
		cr *T.Cairo_t,
		drawable *T.GdkDrawable)

	Gdk_cairo_set_source_color func(
		cr *T.Cairo_t,
		color *T.GdkColor)

	Gdk_cairo_set_source_pixbuf func(
		cr *T.Cairo_t,
		pixbuf *T.GdkPixbuf,
		pixbuf_x T.Double,
		pixbuf_y T.Double)

	Gdk_cairo_set_source_pixmap func(
		cr *T.Cairo_t,
		pixmap *T.GdkPixmap,
		pixmap_x T.Double,
		pixmap_y T.Double)

	Gdk_cairo_set_source_window func(
		cr *T.Cairo_t,
		window *T.GdkWindow,
		x T.Double,
		y T.Double)

	Gdk_cairo_rectangle func(
		cr *T.Cairo_t,
		rectangle *T.GdkRectangle)

	Gdk_cairo_region func(
		cr *T.Cairo_t,
		region *T.GdkRegion)

	Gdk_cursor_get_type func() T.GType

	Gdk_cursor_new_for_display func(
		display *T.GdkDisplay,
		cursor_type T.GdkCursorType) *T.GdkCursor

	Gdk_cursor_new func(
		cursor_type T.GdkCursorType) *T.GdkCursor

	Gdk_cursor_new_from_pixmap func(
		source *T.GdkPixmap,
		mask *T.GdkPixmap,
		fg *T.GdkColor,
		bg *T.GdkColor,
		x T.Gint,
		y T.Gint) *T.GdkCursor

	Gdk_cursor_new_from_pixbuf func(
		display *T.GdkDisplay,
		pixbuf *T.GdkPixbuf,
		x T.Gint,
		y T.Gint) *T.GdkCursor

	Gdk_cursor_get_display func(
		cursor *T.GdkCursor) *T.GdkDisplay

	Gdk_cursor_ref func(
		cursor *T.GdkCursor) *T.GdkCursor

	Gdk_cursor_unref func(
		cursor *T.GdkCursor)

	Gdk_cursor_new_from_name func(
		display *T.GdkDisplay,
		name string) *T.GdkCursor

	Gdk_cursor_get_image func(
		cursor *T.GdkCursor) *T.GdkPixbuf

	Gdk_cursor_get_cursor_type func(
		cursor *T.GdkCursor) T.GdkCursorType

	Gdk_display_manager_get_type func() T.GType

	Gdk_display_manager_get func() *T.GdkDisplayManager

	Gdk_display_manager_get_default_display func(
		display_manager *T.GdkDisplayManager) *T.GdkDisplay

	Gdk_display_manager_set_default_display func(
		display_manager *T.GdkDisplayManager,
		display *T.GdkDisplay)

	Gdk_display_manager_list_displays func(
		display_manager *T.GdkDisplayManager) *T.GSList

	Gdk_gc_get_type func() T.GType

	Gdk_gc_new func(
		drawable *T.GdkDrawable) *T.GdkGC

	Gdk_gc_new_with_values func(
		drawable *T.GdkDrawable,
		values *T.GdkGCValues,
		values_mask T.GdkGCValuesMask) *T.GdkGC

	Gdk_gc_ref func(
		gc *T.GdkGC) *T.GdkGC

	Gdk_gc_unref func(
		gc *T.GdkGC)

	Gdk_gc_get_values func(
		gc *T.GdkGC,
		values *T.GdkGCValues)

	Gdk_gc_set_values func(
		gc *T.GdkGC,
		values *T.GdkGCValues,
		values_mask T.GdkGCValuesMask)

	Gdk_gc_set_foreground func(
		gc *T.GdkGC,
		color *T.GdkColor)

	Gdk_gc_set_background func(
		gc *T.GdkGC,
		color *T.GdkColor)

	Gdk_gc_set_font func(
		gc *T.GdkGC,
		font *T.GdkFont)

	Gdk_gc_set_function func(
		gc *T.GdkGC,
		function T.GdkFunction)

	Gdk_gc_set_fill func(
		gc *T.GdkGC,
		fill T.GdkFill)

	Gdk_gc_set_tile func(
		gc *T.GdkGC,
		tile *T.GdkPixmap)

	Gdk_gc_set_stipple func(
		gc *T.GdkGC,
		stipple *T.GdkPixmap)

	Gdk_gc_set_ts_origin func(
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint)

	Gdk_gc_set_clip_origin func(
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint)

	Gdk_gc_set_clip_mask func(
		gc *T.GdkGC,
		mask *T.GdkBitmap)

	Gdk_gc_set_clip_rectangle func(
		gc *T.GdkGC,
		rectangle *T.GdkRectangle)

	Gdk_gc_set_clip_region func(
		gc *T.GdkGC,
		region *T.GdkRegion)

	Gdk_gc_set_subwindow func(
		gc *T.GdkGC,
		mode T.GdkSubwindowMode)

	Gdk_gc_set_exposures func(
		gc *T.GdkGC,
		exposures T.Gboolean)

	Gdk_gc_set_line_attributes func(
		gc *T.GdkGC,
		line_width T.Gint,
		line_style T.GdkLineStyle,
		cap_style T.GdkCapStyle,
		join_style T.GdkJoinStyle)

	Gdk_gc_set_dashes func(
		gc *T.GdkGC,
		dash_offset T.Gint,
		dash_list *T.Gint8,
		n T.Gint)

	Gdk_gc_offset func(
		gc *T.GdkGC,
		x_offset T.Gint,
		y_offset T.Gint)

	Gdk_gc_copy func(
		dst_gc *T.GdkGC,
		src_gc *T.GdkGC)

	Gdk_gc_set_colormap func(
		gc *T.GdkGC,
		colormap *T.GdkColormap)

	Gdk_gc_get_colormap func(
		gc *T.GdkGC) *T.GdkColormap

	Gdk_gc_set_rgb_fg_color func(
		gc *T.GdkGC,
		color *T.GdkColor)

	Gdk_gc_set_rgb_bg_color func(
		gc *T.GdkGC,
		color *T.GdkColor)

	Gdk_gc_get_screen func(
		gc *T.GdkGC) *T.GdkScreen

	Gdk_drawable_get_type func() T.GType

	Gdk_drawable_set_data func(
		drawable *T.GdkDrawable,
		key string,
		data T.Gpointer,
		destroy_func T.GDestroyNotify)

	Gdk_drawable_get_data func(
		drawable *T.GdkDrawable,
		key string) T.Gpointer

	Gdk_drawable_set_colormap func(
		drawable *T.GdkDrawable,
		colormap *T.GdkColormap)

	Gdk_drawable_get_colormap func(
		drawable *T.GdkDrawable) *T.GdkColormap

	Gdk_drawable_get_depth func(
		drawable *T.GdkDrawable) T.Gint

	Gdk_drawable_get_size func(
		drawable *T.GdkDrawable,
		width *T.Gint,
		height *T.Gint)

	Gdk_drawable_get_visual func(
		drawable *T.GdkDrawable) *T.GdkVisual

	Gdk_drawable_get_screen func(
		drawable *T.GdkDrawable) *T.GdkScreen

	Gdk_drawable_get_display func(
		drawable *T.GdkDrawable) *T.GdkDisplay

	Gdk_drawable_ref func(
		drawable *T.GdkDrawable) *T.GdkDrawable

	Gdk_drawable_unref func(
		drawable *T.GdkDrawable)

	Gdk_draw_point func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint)

	Gdk_draw_line func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x1_ T.Gint,
		y1_ T.Gint,
		x2_ T.Gint,
		y2_ T.Gint)

	Gdk_draw_rectangle func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		filled T.Gboolean,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gdk_draw_arc func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		filled T.Gboolean,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		angle1 T.Gint,
		angle2 T.Gint)

	Gdk_draw_polygon func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		filled T.Gboolean,
		points *T.GdkPoint,
		n_points T.Gint)

	Gdk_draw_string func(
		drawable *T.GdkDrawable,
		font *T.GdkFont,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		s string)

	Gdk_draw_text func(
		drawable *T.GdkDrawable,
		font *T.GdkFont,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		text string,
		text_length T.Gint)

	Gdk_draw_text_wc func(
		drawable *T.GdkDrawable,
		font *T.GdkFont,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		text *T.GdkWChar,
		text_length T.Gint)

	Gdk_draw_drawable func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		src *T.GdkDrawable,
		xsrc T.Gint,
		ysrc T.Gint,
		xdest T.Gint,
		ydest T.Gint,
		width T.Gint,
		height T.Gint)

	Gdk_draw_image func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		image *T.GdkImage,
		xsrc T.Gint,
		ysrc T.Gint,
		xdest T.Gint,
		ydest T.Gint,
		width T.Gint,
		height T.Gint)

	Gdk_draw_points func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		points *T.GdkPoint,
		n_points T.Gint)

	Gdk_draw_segments func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		segs *T.GdkSegment,
		n_segs T.Gint)

	Gdk_draw_lines func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		points *T.GdkPoint,
		n_points T.Gint)

	Gdk_draw_pixbuf func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		pixbuf *T.GdkPixbuf,
		src_x T.Gint,
		src_y T.Gint,
		dest_x T.Gint,
		dest_y T.Gint,
		width T.Gint,
		height T.Gint,
		dither T.GdkRgbDither,
		x_dither T.Gint,
		y_dither T.Gint)

	Gdk_draw_glyphs func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		font *T.PangoFont,
		x T.Gint,
		y T.Gint,
		glyphs *T.PangoGlyphString)

	Gdk_draw_layout_line func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		line *T.PangoLayoutLine)

	Gdk_draw_layout func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		layout *T.PangoLayout)

	Gdk_draw_layout_line_with_colors func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		line *T.PangoLayoutLine,
		foreground *T.GdkColor,
		background *T.GdkColor)

	Gdk_draw_layout_with_colors func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		x T.Gint,
		y T.Gint,
		layout *T.PangoLayout,
		foreground *T.GdkColor,
		background *T.GdkColor)

	Gdk_draw_glyphs_transformed func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		matrix *T.PangoMatrix,
		font *T.PangoFont,
		x T.Gint,
		y T.Gint,
		glyphs *T.PangoGlyphString)

	Gdk_draw_trapezoids func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		trapezoids *T.GdkTrapezoid,
		n_trapezoids T.Gint)

	Gdk_drawable_get_image func(
		drawable *T.GdkDrawable,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint) *T.GdkImage

	Gdk_drawable_copy_to_image func(
		drawable *T.GdkDrawable,
		image *T.GdkImage,
		src_x T.Gint,
		src_y T.Gint,
		dest_x T.Gint,
		dest_y T.Gint,
		width T.Gint,
		height T.Gint) *T.GdkImage

	Gdk_drawable_get_clip_region func(
		drawable *T.GdkDrawable) *T.GdkRegion

	Gdk_drawable_get_visible_region func(
		drawable *T.GdkDrawable) *T.GdkRegion

	Gdk_cursor_type_get_type func() T.GType

	Gdk_drag_action_get_type func() T.GType

	Gdk_drag_protocol_get_type func() T.GType

	Gdk_filter_return_get_type func() T.GType

	Gdk_event_type_get_type func() T.GType

	Gdk_event_mask_get_type func() T.GType

	Gdk_visibility_state_get_type func() T.GType

	Gdk_scroll_direction_get_type func() T.GType

	Gdk_notify_type_get_type func() T.GType

	Gdk_crossing_mode_get_type func() T.GType

	Gdk_property_state_get_type func() T.GType

	Gdk_window_state_get_type func() T.GType

	Gdk_setting_action_get_type func() T.GType

	Gdk_owner_change_get_type func() T.GType

	Gdk_font_type_get_type func() T.GType

	Gdk_cap_style_get_type func() T.GType

	Gdk_fill_get_type func() T.GType

	Gdk_function_get_type func() T.GType

	Gdk_join_style_get_type func() T.GType

	Gdk_line_style_get_type func() T.GType

	Gdk_subwindow_mode_get_type func() T.GType

	Gdk_gc_values_mask_get_type func() T.GType

	Gdk_image_type_get_type func() T.GType

	Gdk_extension_mode_get_type func() T.GType

	Gdk_input_source_get_type func() T.GType

	Gdk_input_mode_get_type func() T.GType

	Gdk_axis_use_get_type func() T.GType

	Gdk_prop_mode_get_type func() T.GType

	Gdk_fill_rule_get_type func() T.GType

	Gdk_overlap_type_get_type func() T.GType

	Gdk_rgb_dither_get_type func() T.GType

	Gdk_byte_order_get_type func() T.GType

	Gdk_modifier_type_get_type func() T.GType

	Gdk_input_condition_get_type func() T.GType

	Gdk_status_get_type func() T.GType

	Gdk_grab_status_get_type func() T.GType

	Gdk_visual_type_get_type func() T.GType

	Gdk_window_class_get_type func() T.GType

	Gdk_window_type_get_type func() T.GType

	Gdk_window_attributes_type_get_type func() T.GType

	Gdk_window_hints_get_type func() T.GType

	Gdk_window_type_hint_get_type func() T.GType

	Gdk_wm_decoration_get_type func() T.GType

	Gdk_wm_function_get_type func() T.GType

	Gdk_gravity_get_type func() T.GType

	Gdk_window_edge_get_type func() T.GType

	Gdk_font_get_type func() T.GType

	Gdk_font_ref func(
		font *T.GdkFont) *T.GdkFont

	Gdk_font_unref func(
		font *T.GdkFont)

	Gdk_font_id func(
		font *T.GdkFont) T.Gint

	Gdk_font_equal func(
		fonta *T.GdkFont,
		fontb *T.GdkFont) T.Gboolean

	Gdk_font_load_for_display func(
		display *T.GdkDisplay,
		font_name string) *T.GdkFont

	Gdk_fontset_load_for_display func(
		display *T.GdkDisplay,
		fontset_name string) *T.GdkFont

	Gdk_font_from_description_for_display func(
		display *T.GdkDisplay,
		font_desc *T.PangoFontDescription) *T.GdkFont

	Gdk_font_load func(
		font_name string) *T.GdkFont

	Gdk_fontset_load func(
		fontset_name string) *T.GdkFont

	Gdk_font_from_description func(
		font_desc *T.PangoFontDescription) *T.GdkFont

	Gdk_string_width func(
		font *T.GdkFont,
		s string) T.Gint

	Gdk_text_width func(
		font *T.GdkFont,
		text string,
		text_length T.Gint) T.Gint

	Gdk_text_width_wc func(
		font *T.GdkFont,
		text *T.GdkWChar,
		text_length T.Gint) T.Gint

	Gdk_char_width func(
		font *T.GdkFont,
		character T.Gchar) T.Gint

	Gdk_char_width_wc func(
		font *T.GdkFont,
		character T.GdkWChar) T.Gint

	Gdk_string_measure func(
		font *T.GdkFont,
		s string) T.Gint

	Gdk_text_measure func(
		font *T.GdkFont,
		text string,
		text_length T.Gint) T.Gint

	Gdk_char_measure func(
		font *T.GdkFont,
		character T.Gchar) T.Gint

	Gdk_string_height func(
		font *T.GdkFont,
		s string) T.Gint

	Gdk_text_height func(
		font *T.GdkFont,
		text string,
		text_length T.Gint) T.Gint

	Gdk_char_height func(
		font *T.GdkFont,
		character T.Gchar) T.Gint

	Gdk_text_extents func(
		font *T.GdkFont,
		text string,
		text_length T.Gint,
		lbearing *T.Gint,
		rbearing *T.Gint,
		width *T.Gint,
		ascent *T.Gint,
		descent *T.Gint)

	Gdk_text_extents_wc func(
		font *T.GdkFont,
		text *T.GdkWChar,
		text_length T.Gint,
		lbearing *T.Gint,
		rbearing *T.Gint,
		width *T.Gint,
		ascent *T.Gint,
		descent *T.Gint)

	Gdk_string_extents func(
		font *T.GdkFont,
		s string,
		lbearing *T.Gint,
		rbearing *T.Gint,
		width *T.Gint,
		ascent *T.Gint,
		descent *T.Gint)

	Gdk_font_get_display func(
		font *T.GdkFont) *T.GdkDisplay

	Gdk_image_get_type func() T.GType

	Gdk_image_new func(
		typ T.GdkImageType,
		visual *T.GdkVisual,
		width T.Gint,
		height T.Gint) *T.GdkImage

	Gdk_image_get func(
		drawable *T.GdkDrawable,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint) *T.GdkImage

	Gdk_image_ref func(
		image *T.GdkImage) *T.GdkImage

	Gdk_image_unref func(
		image *T.GdkImage)

	Gdk_image_put_pixel func(
		image *T.GdkImage,
		x T.Gint,
		y T.Gint,
		pixel T.Guint32)

	Gdk_image_get_pixel func(
		image *T.GdkImage,
		x T.Gint,
		y T.Gint) T.Guint32

	Gdk_image_set_colormap func(
		image *T.GdkImage,
		colormap *T.GdkColormap)

	Gdk_image_get_colormap func(
		image *T.GdkImage) *T.GdkColormap

	Gdk_image_get_image_type func(
		image *T.GdkImage) T.GdkImageType

	Gdk_image_get_visual func(
		image *T.GdkImage) *T.GdkVisual

	Gdk_image_get_byte_order func(
		image *T.GdkImage) T.GdkByteOrder

	Gdk_image_get_width func(
		image *T.GdkImage) T.Gint

	Gdk_image_get_height func(
		image *T.GdkImage) T.Gint

	Gdk_image_get_depth func(
		image *T.GdkImage) T.Guint16

	Gdk_image_get_bytes_per_pixel func(
		image *T.GdkImage) T.Guint16

	Gdk_image_get_bytes_per_line func(
		image *T.GdkImage) T.Guint16

	Gdk_image_get_bits_per_pixel func(
		image *T.GdkImage) T.Guint16

	Gdk_image_get_pixels func(
		image *T.GdkImage) T.Gpointer

	Gdk_keymap_get_type func() T.GType

	Gdk_keymap_get_default func() *T.GdkKeymap

	Gdk_keymap_get_for_display func(
		display *T.GdkDisplay) *T.GdkKeymap

	Gdk_keymap_lookup_key func(
		keymap *T.GdkKeymap,
		key *T.GdkKeymapKey) T.Guint

	Gdk_keymap_translate_keyboard_state func(
		keymap *T.GdkKeymap,
		hardware_keycode T.Guint,
		state T.GdkModifierType,
		group T.Gint,
		keyval *T.Guint,
		effective_group *T.Gint,
		level *T.Gint,
		consumed_modifiers *T.GdkModifierType) T.Gboolean

	Gdk_keymap_get_entries_for_keyval func(
		keymap *T.GdkKeymap,
		keyval T.Guint,
		keys **T.GdkKeymapKey,
		n_keys *T.Gint) T.Gboolean

	Gdk_keymap_get_entries_for_keycode func(
		keymap *T.GdkKeymap,
		hardware_keycode T.Guint,
		keys **T.GdkKeymapKey,
		keyvals **T.Guint,
		n_entries *T.Gint) T.Gboolean

	Gdk_keymap_get_direction func(
		keymap *T.GdkKeymap) T.PangoDirection

	Gdk_keymap_have_bidi_layouts func(
		keymap *T.GdkKeymap) T.Gboolean

	Gdk_keymap_get_caps_lock_state func(
		keymap *T.GdkKeymap) T.Gboolean

	Gdk_keymap_add_virtual_modifiers func(
		keymap *T.GdkKeymap,
		state *T.GdkModifierType)

	Gdk_keymap_map_virtual_modifiers func(
		keymap *T.GdkKeymap,
		state *T.GdkModifierType) T.Gboolean

	Gdk_keyval_name func(
		keyval T.Guint) string

	Gdk_keyval_from_name func(
		keyval_name string) T.Guint

	Gdk_keyval_convert_case func(
		symbol T.Guint,
		lower *T.Guint,
		upper *T.Guint)

	Gdk_keyval_to_upper func(
		keyval T.Guint) T.Guint

	Gdk_keyval_to_lower func(
		keyval T.Guint) T.Guint

	Gdk_keyval_is_upper func(
		keyval T.Guint) T.Gboolean

	Gdk_keyval_is_lower func(
		keyval T.Guint) T.Gboolean

	Gdk_keyval_to_unicode func(
		keyval T.Guint) T.Guint32

	Gdk_unicode_to_keyval func(
		wc T.Guint32) T.Guint

	Gdk_pango_renderer_get_type func() T.GType

	Gdk_pango_renderer_new func(
		screen *T.GdkScreen) *T.PangoRenderer

	Gdk_pango_renderer_get_default func(
		screen *T.GdkScreen) *T.PangoRenderer

	Gdk_pango_renderer_set_drawable func(

		Gdk_renderer *T.GdkPangoRenderer,
		drawable *T.GdkDrawable)

	Gdk_pango_renderer_set_gc func(

		Gdk_renderer *T.GdkPangoRenderer,
		gc *T.GdkGC)

	Gdk_pango_renderer_set_stipple func(

		Gdk_renderer *T.GdkPangoRenderer,
		part T.PangoRenderPart,
		stipple *T.GdkBitmap)

	Gdk_pango_renderer_set_override_color func(

		Gdk_renderer *T.GdkPangoRenderer,
		part T.PangoRenderPart,
		color *T.GdkColor)

	Gdk_pango_context_get_for_screen func(
		screen *T.GdkScreen) *T.PangoContext

	Gdk_pango_context_get func() *T.PangoContext

	Gdk_pango_context_set_colormap func(
		context *T.PangoContext,
		colormap *T.GdkColormap)

	Gdk_pango_layout_line_get_clip_region func(
		line *T.PangoLayoutLine,
		x_origin T.Gint,
		y_origin T.Gint,
		index_ranges *T.Gint,
		n_ranges T.Gint) *T.GdkRegion

	Gdk_pango_layout_get_clip_region func(
		layout *T.PangoLayout,
		x_origin T.Gint,
		y_origin T.Gint,
		index_ranges *T.Gint,
		n_ranges T.Gint) *T.GdkRegion

	Gdk_pango_attr_stipple_new func(
		stipple *T.GdkBitmap) *T.PangoAttribute

	Gdk_pango_attr_embossed_new func(
		embossed T.Gboolean) *T.PangoAttribute

	Gdk_pango_attr_emboss_color_new func(
		color *T.GdkColor) *T.PangoAttribute

	Gdk_pixmap_get_type func() T.GType

	Gdk_pixmap_new func(
		drawable *T.GdkDrawable,
		width T.Gint,
		height T.Gint,
		depth T.Gint) *T.GdkPixmap

	Gdk_bitmap_create_from_data func(
		drawable *T.GdkDrawable,
		data string,
		width T.Gint,
		height T.Gint) *T.GdkBitmap

	Gdk_pixmap_create_from_data func(
		drawable *T.GdkDrawable,
		data string,
		width T.Gint,
		height T.Gint,
		depth T.Gint,
		fg *T.GdkColor,
		bg *T.GdkColor) *T.GdkPixmap

	Gdk_pixmap_create_from_xpm func(
		drawable *T.GdkDrawable,
		mask **T.GdkBitmap,
		transparent_color *T.GdkColor,
		filename string) *T.GdkPixmap

	Gdk_pixmap_colormap_create_from_xpm func(
		drawable *T.GdkDrawable,
		colormap *T.GdkColormap,
		mask **T.GdkBitmap,
		transparent_color *T.GdkColor,
		filename string) *T.GdkPixmap

	Gdk_pixmap_create_from_xpm_d func(
		drawable *T.GdkDrawable,
		mask **T.GdkBitmap,
		transparent_color *T.GdkColor,
		data **T.Gchar) *T.GdkPixmap

	Gdk_pixmap_colormap_create_from_xpm_d func(
		drawable *T.GdkDrawable,
		colormap *T.GdkColormap,
		mask **T.GdkBitmap,
		transparent_color *T.GdkColor,
		data **T.Gchar) *T.GdkPixmap

	Gdk_pixmap_get_size func(
		pixmap *T.GdkPixmap,
		width *T.Gint,
		height *T.Gint)

	Gdk_pixmap_foreign_new func(
		anid T.GdkNativeWindow) *T.GdkPixmap

	Gdk_pixmap_lookup func(
		anid T.GdkNativeWindow) *T.GdkPixmap

	Gdk_pixmap_foreign_new_for_display func(
		display *T.GdkDisplay,
		anid T.GdkNativeWindow) *T.GdkPixmap

	Gdk_pixmap_lookup_for_display func(
		display *T.GdkDisplay,
		anid T.GdkNativeWindow) *T.GdkPixmap

	Gdk_pixmap_foreign_new_for_screen func(
		screen *T.GdkScreen,
		anid T.GdkNativeWindow,
		width T.Gint,
		height T.Gint,
		depth T.Gint) *T.GdkPixmap

	Gdk_atom_intern func(
		atom_name string,
		only_if_exists T.Gboolean) T.GdkAtom

	Gdk_atom_intern_static_string func(
		atom_name string) T.GdkAtom

	Gdk_atom_name func(
		atom T.GdkAtom) string

	Gdk_property_get func(
		window *T.GdkWindow,
		property T.GdkAtom,
		typ T.GdkAtom,
		offset T.Gulong,
		length T.Gulong,
		pdelete T.Gint,
		actual_property_type *T.GdkAtom,
		actual_format *T.Gint,
		actual_length *T.Gint,
		data **T.Guchar) T.Gboolean

	Gdk_property_change func(
		window *T.GdkWindow,
		property T.GdkAtom,
		typ T.GdkAtom,
		format T.Gint,
		mode T.GdkPropMode,
		data *T.Guchar,
		nelements T.Gint)

	Gdk_property_delete func(
		window *T.GdkWindow,
		property T.GdkAtom)

	Gdk_text_property_to_text_list func(
		encoding T.GdkAtom,
		format T.Gint,
		text *T.Guchar,
		length T.Gint,
		list ***T.Gchar) T.Gint

	Gdk_utf8_to_compound_text func(
		str string,
		encoding *T.GdkAtom,
		format *T.Gint,
		ctext **T.Guchar,
		length *T.Gint) T.Gboolean

	Gdk_string_to_compound_text func(
		str string,
		encoding *T.GdkAtom,
		format *T.Gint,
		ctext **T.Guchar,
		length *T.Gint) T.Gint

	Gdk_text_property_to_utf8_list func(
		encoding T.GdkAtom,
		format T.Gint,
		text *T.Guchar,
		length T.Gint,
		list ***T.Gchar) T.Gint

	Gdk_text_property_to_utf8_list_for_display func(
		display *T.GdkDisplay,
		encoding T.GdkAtom,
		format T.Gint,
		text *T.Guchar,
		length T.Gint,
		list ***T.Gchar) T.Gint

	Gdk_utf8_to_string_target func(
		str string) string

	Gdk_text_property_to_text_list_for_display func(
		display *T.GdkDisplay,
		encoding T.GdkAtom,
		format T.Gint,
		text *T.Guchar,
		length T.Gint,
		list ***T.Gchar) T.Gint

	Gdk_string_to_compound_text_for_display func(
		display *T.GdkDisplay,
		str string,
		encoding *T.GdkAtom,
		format *T.Gint,
		ctext **T.Guchar,
		length *T.Gint) T.Gint

	Gdk_utf8_to_compound_text_for_display func(
		display *T.GdkDisplay,
		str string,
		encoding *T.GdkAtom,
		format *T.Gint,
		ctext **T.Guchar,
		length *T.Gint) T.Gboolean

	Gdk_free_text_list func(
		list **T.Gchar)

	Gdk_free_compound_text func(
		ctext *T.Guchar)

	Gdk_region_new func() *T.GdkRegion

	Gdk_region_polygon func(
		points *T.GdkPoint,
		n_points T.Gint,
		fill_rule T.GdkFillRule) *T.GdkRegion

	Gdk_region_copy func(
		region *T.GdkRegion) *T.GdkRegion

	Gdk_region_rectangle func(
		rectangle *T.GdkRectangle) *T.GdkRegion

	Gdk_region_destroy func(
		region *T.GdkRegion)

	Gdk_region_get_clipbox func(
		region *T.GdkRegion,
		rectangle *T.GdkRectangle)

	Gdk_region_get_rectangles func(
		region *T.GdkRegion,
		rectangles **T.GdkRectangle,
		n_rectangles *T.Gint)

	Gdk_region_empty func(
		region *T.GdkRegion) T.Gboolean

	Gdk_region_equal func(
		region1 *T.GdkRegion,
		region2 *T.GdkRegion) T.Gboolean

	Gdk_region_rect_equal func(
		region *T.GdkRegion,
		rectangle *T.GdkRectangle) T.Gboolean

	Gdk_region_point_in func(
		region *T.GdkRegion,
		x int,
		y int) T.Gboolean

	Gdk_region_rect_in func(
		region *T.GdkRegion,
		rectangle *T.GdkRectangle) T.GdkOverlapType

	Gdk_region_offset func(
		region *T.GdkRegion,
		dx T.Gint,
		dy T.Gint)

	Gdk_region_shrink func(
		region *T.GdkRegion,
		dx T.Gint,
		dy T.Gint)

	Gdk_region_union_with_rect func(
		region *T.GdkRegion,
		rect *T.GdkRectangle)

	Gdk_region_intersect func(
		source1 *T.GdkRegion,
		source2 *T.GdkRegion)

	Gdk_region_union func(
		source1 *T.GdkRegion,
		source2 *T.GdkRegion)

	Gdk_region_subtract func(
		source1 *T.GdkRegion,
		source2 *T.GdkRegion)

	Gdk_region_xor func(
		source1 *T.GdkRegion,
		source2 *T.GdkRegion)

	Gdk_region_spans_intersect_foreach func(
		region *T.GdkRegion,
		spans *T.GdkSpan,
		n_spans int,
		sorted T.Gboolean,
		function T.GdkSpanFunc,
		data T.Gpointer)

	Gdk_selection_owner_set func(
		owner *T.GdkWindow,
		selection T.GdkAtom,
		time_ T.Guint32,
		send_event T.Gboolean) T.Gboolean

	Gdk_selection_owner_get func(
		selection T.GdkAtom) *T.GdkWindow

	Gdk_selection_owner_set_for_display func(
		display *T.GdkDisplay,
		owner *T.GdkWindow,
		selection T.GdkAtom,
		time_ T.Guint32,
		send_event T.Gboolean) T.Gboolean

	Gdk_selection_owner_get_for_display func(
		display *T.GdkDisplay,
		selection T.GdkAtom) *T.GdkWindow

	Gdk_selection_convert func(
		requestor *T.GdkWindow,
		selection T.GdkAtom,
		target T.GdkAtom,
		time_ T.Guint32)

	Gdk_selection_property_get func(
		requestor *T.GdkWindow,
		data **T.Guchar,
		prop_type *T.GdkAtom,
		prop_format *T.Gint) T.Gint

	Gdk_selection_send_notify func(
		requestor T.GdkNativeWindow,
		selection T.GdkAtom,
		target T.GdkAtom,
		property T.GdkAtom,
		time_ T.Guint32)

	Gdk_selection_send_notify_for_display func(
		display *T.GdkDisplay,
		requestor T.GdkNativeWindow,
		selection T.GdkAtom,
		target T.GdkAtom,
		property T.GdkAtom,
		time_ T.Guint32)

	Gdk_spawn_on_screen func(
		screen *T.GdkScreen,
		working_directory string,
		argv **T.Gchar,
		envp **T.Gchar,
		flags T.GSpawnFlags,
		child_setup T.GSpawnChildSetupFunc,
		user_data T.Gpointer,
		child_pid *T.Gint,
		e **T.GError) T.Gboolean

	Gdk_spawn_on_screen_with_pipes func(
		screen *T.GdkScreen,
		working_directory string,
		argv **T.Gchar,
		envp **T.Gchar,
		flags T.GSpawnFlags,
		child_setup T.GSpawnChildSetupFunc,
		user_data T.Gpointer,
		child_pid *T.Gint,
		standard_input *T.Gint,
		standard_output *T.Gint,
		standard_error *T.Gint,
		e **T.GError) T.Gboolean

	Gdk_spawn_command_line_on_screen func(
		screen *T.GdkScreen,
		command_line string,
		e **T.GError) T.Gboolean

	Gdk_window_object_get_type func() T.GType

	Gdk_window_new func(
		parent *T.GdkWindow,
		attributes *T.GdkWindowAttr,
		attributes_mask T.Gint) *T.GdkWindow

	Gdk_window_destroy func(
		window *T.GdkWindow)

	Gdk_window_get_window_type func(
		window *T.GdkWindow) T.GdkWindowType

	Gdk_window_is_destroyed func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_get_screen func(
		window *T.GdkWindow) *T.GdkScreen

	Gdk_window_get_display func(
		window *T.GdkWindow) *T.GdkDisplay

	Gdk_window_get_visual func(
		window *T.GdkWindow) *T.GdkVisual

	Gdk_window_get_width func(
		window *T.GdkWindow) int

	Gdk_window_get_height func(
		window *T.GdkWindow) int

	Gdk_window_at_pointer func(
		win_x *T.Gint,
		win_y *T.Gint) *T.GdkWindow

	Gdk_window_show func(
		window *T.GdkWindow)

	Gdk_window_hide func(
		window *T.GdkWindow)

	Gdk_window_withdraw func(
		window *T.GdkWindow)

	Gdk_window_show_unraised func(
		window *T.GdkWindow)

	Gdk_window_move func(
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint)

	Gdk_window_resize func(
		window *T.GdkWindow,
		width T.Gint,
		height T.Gint)

	Gdk_window_move_resize func(
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gdk_window_reparent func(
		window *T.GdkWindow,
		new_parent *T.GdkWindow,
		x T.Gint,
		y T.Gint)

	Gdk_window_clear func(
		window *T.GdkWindow)

	Gdk_window_clear_area func(
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gdk_window_clear_area_e func(
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gdk_window_raise func(
		window *T.GdkWindow)

	Gdk_window_lower func(
		window *T.GdkWindow)

	Gdk_window_restack func(
		window *T.GdkWindow,
		sibling *T.GdkWindow,
		above T.Gboolean)

	Gdk_window_focus func(
		window *T.GdkWindow,
		timestamp T.Guint32)

	Gdk_window_set_user_data func(
		window *T.GdkWindow,
		user_data T.Gpointer)

	Gdk_window_set_override_redirect func(
		window *T.GdkWindow,
		override_redirect T.Gboolean)

	Gdk_window_get_accept_focus func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_set_accept_focus func(
		window *T.GdkWindow,
		accept_focus T.Gboolean)

	Gdk_window_get_focus_on_map func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_set_focus_on_map func(
		window *T.GdkWindow,
		focus_on_map T.Gboolean)

	Gdk_window_add_filter func(
		window *T.GdkWindow,
		function T.GdkFilterFunc,
		data T.Gpointer)

	Gdk_window_remove_filter func(
		window *T.GdkWindow,
		function T.GdkFilterFunc,
		data T.Gpointer)

	Gdk_window_scroll func(
		window *T.GdkWindow,
		dx T.Gint,
		dy T.Gint)

	Gdk_window_move_region func(
		window *T.GdkWindow,
		region *T.GdkRegion,
		dx T.Gint,
		dy T.Gint)

	Gdk_window_ensure_native func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_shape_combine_mask func(
		window *T.GdkWindow,
		mask *T.GdkBitmap,
		x T.Gint,
		y T.Gint)

	Gdk_window_shape_combine_region func(
		window *T.GdkWindow,
		shape_region *T.GdkRegion,
		offset_x T.Gint,
		offset_y T.Gint)

	Gdk_window_set_child_shapes func(
		window *T.GdkWindow)

	Gdk_window_get_composited func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_set_composited func(
		window *T.GdkWindow,
		composited T.Gboolean)

	Gdk_window_merge_child_shapes func(
		window *T.GdkWindow)

	Gdk_window_input_shape_combine_mask func(
		window *T.GdkWindow,
		mask *T.GdkBitmap,
		x T.Gint,
		y T.Gint)

	Gdk_window_input_shape_combine_region func(
		window *T.GdkWindow,
		shape_region *T.GdkRegion,
		offset_x T.Gint,
		offset_y T.Gint)

	Gdk_window_set_child_input_shapes func(
		window *T.GdkWindow)

	Gdk_window_merge_child_input_shapes func(
		window *T.GdkWindow)

	Gdk_window_is_visible func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_is_viewable func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_is_input_only func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_is_shaped func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_get_state func(
		window *T.GdkWindow) T.GdkWindowState

	Gdk_window_set_static_gravities func(
		window *T.GdkWindow,
		use_static T.Gboolean) T.Gboolean

	Gdk_window_foreign_new func(
		anid T.GdkNativeWindow) *T.GdkWindow

	Gdk_window_lookup func(
		anid T.GdkNativeWindow) *T.GdkWindow

	Gdk_window_foreign_new_for_display func(
		display *T.GdkDisplay,
		anid T.GdkNativeWindow) *T.GdkWindow

	Gdk_window_lookup_for_display func(
		display *T.GdkDisplay,
		anid T.GdkNativeWindow) *T.GdkWindow

	Gdk_window_has_native func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_set_hints func(
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint,
		min_width T.Gint,
		min_height T.Gint,
		max_width T.Gint,
		max_height T.Gint,
		flags T.Gint)

	Gdk_window_set_type_hint func(
		window *T.GdkWindow,
		hint T.GdkWindowTypeHint)

	Gdk_window_get_type_hint func(
		window *T.GdkWindow) T.GdkWindowTypeHint

	Gdk_window_get_modal_hint func(
		window *T.GdkWindow) T.Gboolean

	Gdk_window_set_modal_hint func(
		window *T.GdkWindow,
		modal T.Gboolean)

	Gdk_window_set_skip_taskbar_hint func(
		window *T.GdkWindow,
		skips_taskbar T.Gboolean)

	Gdk_window_set_skip_pager_hint func(
		window *T.GdkWindow,
		skips_pager T.Gboolean)

	Gdk_window_set_urgency_hint func(
		window *T.GdkWindow,
		urgent T.Gboolean)

	Gdk_window_set_geometry_hints func(
		window *T.GdkWindow,
		geometry *T.GdkGeometry,
		geom_mask T.GdkWindowHints)

	Gdk_set_sm_client_id func(
		sm_client_id string)

	Gdk_window_begin_paint_rect func(
		window *T.GdkWindow,
		rectangle *T.GdkRectangle)

	Gdk_window_begin_paint_region func(
		window *T.GdkWindow,
		region *T.GdkRegion)

	Gdk_window_end_paint func(
		window *T.GdkWindow)

	Gdk_window_flush func(
		window *T.GdkWindow)

	Gdk_window_set_title func(
		window *T.GdkWindow,
		title string)

	Gdk_window_set_role func(
		window *T.GdkWindow,
		role string)

	Gdk_window_set_startup_id func(
		window *T.GdkWindow,
		startup_id string)

	Gdk_window_set_transient_for func(
		window *T.GdkWindow,
		parent *T.GdkWindow)

	Gdk_window_set_background func(
		window *T.GdkWindow,
		color *T.GdkColor)

	Gdk_window_set_back_pixmap func(
		window *T.GdkWindow,
		pixmap *T.GdkPixmap,
		parent_relative T.Gboolean)

	Gdk_window_get_background_pattern func(
		window *T.GdkWindow) *T.Cairo_pattern_t

	Gdk_window_set_cursor func(
		window *T.GdkWindow,
		cursor *T.GdkCursor)

	Gdk_window_get_cursor func(
		window *T.GdkWindow) *T.GdkCursor

	Gdk_window_get_user_data func(
		window *T.GdkWindow,
		data *T.Gpointer)

	Gdk_window_get_geometry func(
		window *T.GdkWindow,
		x *T.Gint,
		y *T.Gint,
		width *T.Gint,
		height *T.Gint,
		depth *T.Gint)

	Gdk_window_get_position func(
		window *T.GdkWindow,
		x *T.Gint,
		y *T.Gint)

	Gdk_window_get_origin func(
		window *T.GdkWindow,
		x *T.Gint,
		y *T.Gint) T.Gint

	Gdk_window_get_root_coords func(
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint,
		root_x *T.Gint,
		root_y *T.Gint)

	Gdk_window_coords_to_parent func(
		window *T.GdkWindow,
		x T.Gdouble,
		y T.Gdouble,
		parent_x *T.Gdouble,
		parent_y *T.Gdouble)

	Gdk_window_coords_from_parent func(
		window *T.GdkWindow,
		parent_x T.Gdouble,
		parent_y T.Gdouble,
		x *T.Gdouble,
		y *T.Gdouble)

	Gdk_window_get_deskrelative_origin func(
		window *T.GdkWindow,
		x *T.Gint,
		y *T.Gint) T.Gboolean

	Gdk_window_get_root_origin func(
		window *T.GdkWindow,
		x *T.Gint,
		y *T.Gint)

	Gdk_window_get_frame_extents func(
		window *T.GdkWindow,
		rect *T.GdkRectangle)

	Gdk_window_get_pointer func(
		window *T.GdkWindow,
		x *T.Gint,
		y *T.Gint,
		mask *T.GdkModifierType) *T.GdkWindow

	Gdk_window_get_parent func(
		window *T.GdkWindow) *T.GdkWindow

	Gdk_window_get_toplevel func(
		window *T.GdkWindow) *T.GdkWindow

	Gdk_window_get_effective_parent func(
		window *T.GdkWindow) *T.GdkWindow

	Gdk_window_get_effective_toplevel func(
		window *T.GdkWindow) *T.GdkWindow

	Gdk_window_get_children func(
		window *T.GdkWindow) *T.GList

	Gdk_window_peek_children func(
		window *T.GdkWindow) *T.GList

	Gdk_window_get_events func(
		window *T.GdkWindow) T.GdkEventMask

	Gdk_window_set_events func(
		window *T.GdkWindow,
		event_mask T.GdkEventMask)

	Gdk_window_set_icon_list func(
		window *T.GdkWindow,
		pixbufs *T.GList)

	Gdk_window_set_icon func(
		window *T.GdkWindow,
		icon_window *T.GdkWindow,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	Gdk_window_set_icon_name func(
		window *T.GdkWindow,
		name string)

	Gdk_window_set_group func(
		window *T.GdkWindow,
		leader *T.GdkWindow)

	Gdk_window_get_group func(
		window *T.GdkWindow) *T.GdkWindow

	Gdk_window_set_decorations func(
		window *T.GdkWindow,
		decorations T.GdkWMDecoration)

	Gdk_window_get_decorations func(
		window *T.GdkWindow,
		decorations *T.GdkWMDecoration) T.Gboolean

	Gdk_window_set_functions func(
		window *T.GdkWindow,
		functions T.GdkWMFunction)

	Gdk_window_get_toplevels func() *T.GList

	Gdk_window_create_similar_surface func(
		window *T.GdkWindow,
		content T.Cairo_content_t,
		width int,
		height int) *T.Cairo_surface_t

	Gdk_window_beep func(
		window *T.GdkWindow)

	Gdk_window_iconify func(
		window *T.GdkWindow)

	Gdk_window_deiconify func(
		window *T.GdkWindow)

	Gdk_window_stick func(
		window *T.GdkWindow)

	Gdk_window_unstick func(
		window *T.GdkWindow)

	Gdk_window_maximize func(
		window *T.GdkWindow)

	Gdk_window_unmaximize func(
		window *T.GdkWindow)

	Gdk_window_fullscreen func(
		window *T.GdkWindow)

	Gdk_window_unfullscreen func(
		window *T.GdkWindow)

	Gdk_window_set_keep_above func(
		window *T.GdkWindow,
		setting T.Gboolean)

	Gdk_window_set_keep_below func(
		window *T.GdkWindow,
		setting T.Gboolean)

	Gdk_window_set_opacity func(
		window *T.GdkWindow,
		opacity T.Gdouble)

	Gdk_window_register_dnd func(
		window *T.GdkWindow)

	Gdk_window_begin_resize_drag func(
		window *T.GdkWindow,
		edge T.GdkWindowEdge,
		button T.Gint,
		root_x T.Gint,
		root_y T.Gint,
		timestamp T.Guint32)

	Gdk_window_begin_move_drag func(
		window *T.GdkWindow,
		button T.Gint,
		root_x T.Gint,
		root_y T.Gint,
		timestamp T.Guint32)

	Gdk_window_invalidate_rect func(
		window *T.GdkWindow,
		rect *T.GdkRectangle,
		invalidate_children T.Gboolean)

	Gdk_window_invalidate_region func(
		window *T.GdkWindow,
		region *T.GdkRegion,
		invalidate_children T.Gboolean)

	Gdk_window_invalidate_maybe_recurse func(
		window *T.GdkWindow,
		region *T.GdkRegion,
		child_func func(*T.GdkWindow, T.Gpointer) T.Gboolean,
		user_data T.Gpointer)

	Gdk_window_get_update_area func(
		window *T.GdkWindow) *T.GdkRegion

	Gdk_window_freeze_updates func(
		window *T.GdkWindow)

	Gdk_window_thaw_updates func(
		window *T.GdkWindow)

	Gdk_window_freeze_toplevel_updates_libgtk_only func(
		window *T.GdkWindow)

	Gdk_window_thaw_toplevel_updates_libgtk_only func(
		window *T.GdkWindow)

	Gdk_window_process_all_updates func()

	Gdk_window_process_updates func(
		window *T.GdkWindow,
		update_children T.Gboolean)

	Gdk_window_set_debug_updates func(
		setting T.Gboolean)

	Gdk_window_constrain_size func(
		geometry *T.GdkGeometry,
		flags T.Guint,
		width T.Gint,
		height T.Gint,
		new_width *T.Gint,
		new_height *T.Gint)

	Gdk_window_get_internal_paint_info func(
		window *T.GdkWindow,
		real_drawable **T.GdkDrawable,
		x_offset *T.Gint,
		y_offset *T.Gint)

	Gdk_window_enable_synchronized_configure func(
		window *T.GdkWindow)

	Gdk_window_configure_finished func(
		window *T.GdkWindow)

	Gdk_get_default_root_window func() *T.GdkWindow

	Gdk_offscreen_window_get_pixmap func(
		window *T.GdkWindow) *T.GdkPixmap

	Gdk_offscreen_window_set_embedder func(
		window *T.GdkWindow,
		embedder *T.GdkWindow)

	Gdk_offscreen_window_get_embedder func(
		window *T.GdkWindow) *T.GdkWindow

	Gdk_window_geometry_changed func(
		window *T.GdkWindow)

	Gdk_window_redirect_to_drawable func(
		window *T.GdkWindow,
		drawable *T.GdkDrawable,
		src_x T.Gint,
		src_y T.Gint,
		dest_x T.Gint,
		dest_y T.Gint,
		width T.Gint,
		height T.Gint)

	Gdk_window_remove_redirection func(
		window *T.GdkWindow)

	Gdk_set_pointer_hooks func(
		new_hooks *T.GdkPointerHooks) *T.GdkPointerHooks

	Gdk_test_render_sync func(
		window *T.GdkWindow)

	Gdk_test_simulate_key func(
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint,
		keyval T.Guint,
		modifiers T.GdkModifierType,
		key_pressrelease T.GdkEventType) T.Gboolean

	Gdk_test_simulate_button func(
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint,
		button T.Guint,
		modifiers T.GdkModifierType,
		button_pressrelease T.GdkEventType) T.Gboolean

	Gdk_visual_get_type func() T.GType

	Gdk_visual_get_best_depth func() T.Gint

	Gdk_visual_get_best_type func() T.GdkVisualType

	Gdk_visual_get_system func() *T.GdkVisual

	Gdk_visual_get_best func() *T.GdkVisual

	Gdk_visual_get_best_with_depth func(
		depth T.Gint) *T.GdkVisual

	Gdk_visual_get_best_with_type func(
		visual_type T.GdkVisualType) *T.GdkVisual

	Gdk_visual_get_best_with_both func(
		depth T.Gint,
		visual_type T.GdkVisualType) *T.GdkVisual

	Gdk_query_depths func(
		depths **T.Gint,
		count *T.Gint)

	Gdk_query_visual_types func(
		visual_types **T.GdkVisualType,
		count *T.Gint)

	Gdk_list_visuals func() *T.GList

	Gdk_visual_get_screen func(
		visual *T.GdkVisual) *T.GdkScreen

	Gdk_visual_get_visual_type func(
		visual *T.GdkVisual) T.GdkVisualType

	Gdk_visual_get_depth func(
		visual *T.GdkVisual) T.Gint

	Gdk_visual_get_byte_order func(
		visual *T.GdkVisual) T.GdkByteOrder

	Gdk_visual_get_colormap_size func(
		visual *T.GdkVisual) T.Gint

	Gdk_visual_get_bits_per_rgb func(
		visual *T.GdkVisual) T.Gint

	Gdk_visual_get_red_pixel_details func(
		visual *T.GdkVisual,
		mask *T.Guint32,
		shift *T.Gint,
		precision *T.Gint)

	Gdk_visual_get_green_pixel_details func(
		visual *T.GdkVisual,
		mask *T.Guint32,
		shift *T.Gint,
		precision *T.Gint)

	Gdk_visual_get_blue_pixel_details func(
		visual *T.GdkVisual,
		mask *T.Guint32,
		shift *T.Gint,
		precision *T.Gint)

	Gdk_parse_args func(
		argc *T.Gint,
		argv ***T.Gchar)

	Gdk_init func(
		argc *T.Gint,
		argv ***T.Gchar)

	Gdk_init_check func(
		argc *T.Gint,
		argv ***T.Gchar) T.Gboolean

	Gdk_add_option_entries_libgtk_only func(
		group *T.GOptionGroup)

	Gdk_pre_parse_libgtk_only func()

	Gdk_exit func(
		error_code T.Gint)

	Gdk_set_locale func() string

	Gdk_get_program_class func() string

	Gdk_set_program_class func(
		program_class string)

	Gdk_error_trap_push func()

	Gdk_error_trap_pop func() T.Gint

	Gdk_set_use_xshm func(
		use_xshm T.Gboolean)

	Gdk_get_use_xshm func() T.Gboolean

	Gdk_get_display func() string

	Gdk_get_display_arg_name func() string

	Gdk_input_add_full func(
		source T.Gint,
		condition T.GdkInputCondition,
		function T.GdkInputFunction,
		data T.Gpointer,
		destroy T.GDestroyNotify) T.Gint

	Gdk_input_add func(
		source T.Gint,
		condition T.GdkInputCondition,
		function T.GdkInputFunction,
		data T.Gpointer) T.Gint

	Gdk_input_remove func(
		tag T.Gint)

	Gdk_pointer_grab func(
		window *T.GdkWindow,
		owner_events T.Gboolean,
		event_mask T.GdkEventMask,
		confine_to *T.GdkWindow,
		cursor *T.GdkCursor,
		time_ T.Guint32) T.GdkGrabStatus

	Gdk_keyboard_grab func(
		window *T.GdkWindow,
		owner_events T.Gboolean,
		time_ T.Guint32) T.GdkGrabStatus

	Gdk_pointer_grab_info_libgtk_only func(
		display *T.GdkDisplay,
		grab_window **T.GdkWindow,
		owner_events *T.Gboolean) T.Gboolean

	Gdk_keyboard_grab_info_libgtk_only func(
		display *T.GdkDisplay,
		grab_window **T.GdkWindow,
		owner_events *T.Gboolean) T.Gboolean

	Gdk_pointer_ungrab func(
		time_ T.Guint32)

	Gdk_keyboard_ungrab func(
		time_ T.Guint32)

	Gdk_pointer_is_grabbed func() T.Gboolean

	Gdk_screen_width func() T.Gint

	Gdk_screen_height func() T.Gint

	Gdk_screen_width_mm func() T.Gint

	Gdk_screen_height_mm func() T.Gint

	Gdk_beep func()

	Gdk_flush func()

	Gdk_set_double_click_time func(
		msec T.Guint)

	Gdk_rectangle_intersect func(
		src1 *T.GdkRectangle,
		src2 *T.GdkRectangle,
		dest *T.GdkRectangle) T.Gboolean

	Gdk_rectangle_union func(
		src1 *T.GdkRectangle,
		src2 *T.GdkRectangle,
		dest *T.GdkRectangle)

	Gdk_rectangle_get_type func() T.GType

	Gdk_wcstombs func(
		src *T.GdkWChar) string

	Gdk_mbstowcs func(
		dest *T.GdkWChar,
		src string,
		dest_max T.Gint) T.Gint

	Gdk_event_send_client_message func(
		event *T.GdkEvent,
		winid T.GdkNativeWindow) T.Gboolean

	Gdk_event_send_clientmessage_toall func(
		event *T.GdkEvent)

	Gdk_event_send_client_message_for_display func(
		display *T.GdkDisplay,
		event *T.GdkEvent,
		winid T.GdkNativeWindow) T.Gboolean

	Gdk_notify_startup_complete func()

	Gdk_notify_startup_complete_with_id func(
		startup_id string)

	Gdk_threads_enter func()

	Gdk_threads_leave func()

	Gdk_threads_init func()

	Gdk_threads_set_lock_functions func(
		enter_fn T.GCallback,
		leave_fn T.GCallback)

	Gdk_threads_add_idle_full func(
		priority T.Gint,
		function T.GSourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify) T.Guint

	Gdk_threads_add_idle func(
		function T.GSourceFunc,
		data T.Gpointer) T.Guint

	Gdk_threads_add_timeout_full func(
		priority T.Gint,
		interval T.Guint,
		function T.GSourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify) T.Guint

	Gdk_threads_add_timeout func(
		interval T.Guint,
		function T.GSourceFunc,
		data T.Gpointer) T.Guint

	Gdk_threads_add_timeout_seconds_full func(
		priority T.Gint,
		interval T.Guint,
		function T.GSourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify) T.Guint

	Gdk_threads_add_timeout_seconds func(
		interval T.Guint,
		function T.GSourceFunc,
		data T.Gpointer) T.Guint

	Gdk_pixbuf_non_anim_get_type func() T.GType

	Gdk_pixbuf_non_anim_new func(
		pixbuf *T.GdkPixbuf) *T.GdkPixbufAnimation

	Gdk_pixbuf_animation_new_from_file func(
		filename string,
		e **T.GError) *T.GdkPixbufAnimation

	Gdk_pixdata_serialize func(
		pixdata *T.GdkPixdata,
		stream_length_p *T.Guint) *T.Guint8

	Gdk_pixdata_deserialize func(
		pixdata *T.GdkPixdata,
		stream_length T.Guint,
		stream *T.Guint8,
		e **T.GError) T.Gboolean

	Gdk_pixdata_from_pixbuf func(
		pixdata *T.GdkPixdata,
		pixbuf *T.GdkPixbuf,
		use_rle T.Gboolean) T.Gpointer

	Gdk_pixbuf_from_pixdata func(
		pixdata *T.GdkPixdata,
		copy_pixels T.Gboolean,
		e **T.GError) *T.GdkPixbuf

	Gdk_pixdata_to_csource func(
		pixdata *T.GdkPixdata,
		name string,
		dump_type T.GdkPixdataDumpType) *T.GString

	Gdk_pixbuf_set_option func(
		pixbuf *T.GdkPixbuf,
		key string,
		value string) T.Gboolean

	Gdk_window_destroy_notify func(
		window *T.GdkWindow)

	Gdk_synthesize_window_state func(
		window *T.GdkWindow,
		unset_flags T.GdkWindowState,
		set_flags T.GdkWindowState)

	Gdk_win32_window_is_win32 func(
		window *T.GdkWindow) T.Gboolean

	Gdk_win32_window_get_impl_hwnd func(
		window *T.GdkWindow) HWND

	Gdk_win32_handle_table_lookup func(
		handle T.GdkNativeWindow) T.Gpointer

	Gdk_win32_drawable_get_handle func(
		drawable *T.GdkDrawable) HGDIOBJ

	Gdk_win32_hdc_get func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		usage T.GdkGCValuesMask) HDC

	Gdk_win32_hdc_release func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		usage T.GdkGCValuesMask)

	Gdk_win32_selection_add_targets func(
		owner *T.GdkWindow,
		selection T.GdkAtom,
		n_targets T.Gint,
		targets *T.GdkAtom)

	Gdk_win32_icon_to_pixbuf_libgtk_only func(
		hicon HICON) *T.GdkPixbuf

	Gdk_win32_pixbuf_to_hicon_libgtk_only func(
		pixbuf *T.GdkPixbuf) HICON

	Gdk_win32_set_modal_dialog_libgtk_only func(
		window HWND)

	Gdk_win32_begin_direct_draw_libgtk_only func(
		drawable *T.GdkDrawable,
		gc *T.GdkGC,
		priv_data *T.Gpointer,
		x_offset_out *T.Gint,
		y_offset_out *T.Gint) *T.GdkDrawable

	Gdk_win32_end_direct_draw_libgtk_only func(
		priv_data T.Gpointer)

	Gdk_win32_window_foreign_new_for_display func(
		display *T.GdkDisplay,
		anid T.GdkNativeWindow) *T.GdkWindow

	Gdk_win32_window_lookup_for_display func(
		display *T.GdkDisplay,
		anid T.GdkNativeWindow) *T.GdkWindow
)

var dll = "libgdk-win32-2.0-0.dll"

var dllPixbuf = "libgdk_pixbuf-2.0-0.dll"

var apiList = outside.Apis{
	{"gdk_add_client_message_filter", &Gdk_add_client_message_filter},
	{"gdk_add_option_entries_libgtk_only", &Gdk_add_option_entries_libgtk_only},
	{"gdk_app_launch_context_get_type", &Gdk_app_launch_context_get_type},
	{"gdk_app_launch_context_new", &Gdk_app_launch_context_new},
	{"gdk_app_launch_context_set_desktop", &Gdk_app_launch_context_set_desktop},
	{"gdk_app_launch_context_set_display", &Gdk_app_launch_context_set_display},
	{"gdk_app_launch_context_set_icon", &Gdk_app_launch_context_set_icon},
	{"gdk_app_launch_context_set_icon_name", &Gdk_app_launch_context_set_icon_name},
	{"gdk_app_launch_context_set_screen", &Gdk_app_launch_context_set_screen},
	{"gdk_app_launch_context_set_timestamp", &Gdk_app_launch_context_set_timestamp},
	{"gdk_atom_intern", &Gdk_atom_intern},
	{"gdk_atom_intern_static_string", &Gdk_atom_intern_static_string},
	{"gdk_atom_name", &Gdk_atom_name},
	{"gdk_axis_use_get_type", &Gdk_axis_use_get_type},
	{"gdk_beep", &Gdk_beep},
	{"gdk_bitmap_create_from_data", &Gdk_bitmap_create_from_data},
	{"gdk_byte_order_get_type", &Gdk_byte_order_get_type},
	{"gdk_cairo_create", &Gdk_cairo_create},
	{"gdk_cairo_rectangle", &Gdk_cairo_rectangle},
	{"gdk_cairo_region", &Gdk_cairo_region},
	{"gdk_cairo_reset_clip", &Gdk_cairo_reset_clip},
	{"gdk_cairo_set_source_color", &Gdk_cairo_set_source_color},
	{"gdk_cairo_set_source_pixbuf", &Gdk_cairo_set_source_pixbuf},
	{"gdk_cairo_set_source_pixmap", &Gdk_cairo_set_source_pixmap},
	{"gdk_cairo_set_source_window", &Gdk_cairo_set_source_window},
	{"gdk_cap_style_get_type", &Gdk_cap_style_get_type},
	{"gdk_char_height", &Gdk_char_height},
	{"gdk_char_measure", &Gdk_char_measure},
	{"gdk_char_width", &Gdk_char_width},
	{"gdk_char_width_wc", &Gdk_char_width_wc},
	{"gdk_color_alloc", &Gdk_color_alloc},
	{"gdk_color_black", &Gdk_color_black},
	{"gdk_color_change", &Gdk_color_change},
	{"gdk_color_copy", &Gdk_color_copy},
	{"gdk_color_equal", &Gdk_color_equal},
	{"gdk_color_free", &Gdk_color_free},
	{"gdk_color_get_type", &Gdk_color_get_type},
	{"gdk_color_hash", &Gdk_color_hash},
	{"gdk_color_parse", &Gdk_color_parse},
	{"gdk_color_to_string", &Gdk_color_to_string},
	{"gdk_color_white", &Gdk_color_white},
	{"gdk_colormap_alloc_color", &Gdk_colormap_alloc_color},
	{"gdk_colormap_alloc_colors", &Gdk_colormap_alloc_colors},
	{"gdk_colormap_change", &Gdk_colormap_change},
	{"gdk_colormap_free_colors", &Gdk_colormap_free_colors},
	{"gdk_colormap_get_screen", &Gdk_colormap_get_screen},
	{"gdk_colormap_get_system", &Gdk_colormap_get_system},
	{"gdk_colormap_get_system_size", &Gdk_colormap_get_system_size},
	{"gdk_colormap_get_type", &Gdk_colormap_get_type},
	{"gdk_colormap_get_visual", &Gdk_colormap_get_visual},
	{"gdk_colormap_new", &Gdk_colormap_new},
	{"gdk_colormap_query_color", &Gdk_colormap_query_color},
	{"gdk_colormap_ref", &Gdk_colormap_ref},
	{"gdk_colormap_unref", &Gdk_colormap_unref},
	{"gdk_colors_alloc", &Gdk_colors_alloc},
	{"gdk_colors_free", &Gdk_colors_free},
	{"gdk_colors_store", &Gdk_colors_store},
	{"gdk_crossing_mode_get_type", &Gdk_crossing_mode_get_type},
	{"gdk_cursor_get_cursor_type", &Gdk_cursor_get_cursor_type},
	{"gdk_cursor_get_display", &Gdk_cursor_get_display},
	{"gdk_cursor_get_image", &Gdk_cursor_get_image},
	{"gdk_cursor_get_type", &Gdk_cursor_get_type},
	{"gdk_cursor_new", &Gdk_cursor_new},
	{"gdk_cursor_new_for_display", &Gdk_cursor_new_for_display},
	{"gdk_cursor_new_from_name", &Gdk_cursor_new_from_name},
	{"gdk_cursor_new_from_pixbuf", &Gdk_cursor_new_from_pixbuf},
	{"gdk_cursor_new_from_pixmap", &Gdk_cursor_new_from_pixmap},
	{"gdk_cursor_ref", &Gdk_cursor_ref},
	{"gdk_cursor_type_get_type", &Gdk_cursor_type_get_type},
	{"gdk_cursor_unref", &Gdk_cursor_unref},
	{"gdk_device_free_history", &Gdk_device_free_history},
	{"gdk_device_get_axis", &Gdk_device_get_axis},
	{"gdk_device_get_axis_use", &Gdk_device_get_axis_use},
	{"gdk_device_get_core_pointer", &Gdk_device_get_core_pointer},
	{"gdk_device_get_has_cursor", &Gdk_device_get_has_cursor},
	{"gdk_device_get_history", &Gdk_device_get_history},
	{"gdk_device_get_key", &Gdk_device_get_key},
	{"gdk_device_get_mode", &Gdk_device_get_mode},
	{"gdk_device_get_n_axes", &Gdk_device_get_n_axes},
	{"gdk_device_get_n_keys", &Gdk_device_get_n_keys},
	{"gdk_device_get_name", &Gdk_device_get_name},
	{"gdk_device_get_source", &Gdk_device_get_source},
	{"gdk_device_get_state", &Gdk_device_get_state},
	{"gdk_device_get_type", &Gdk_device_get_type},
	{"gdk_device_set_axis_use", &Gdk_device_set_axis_use},
	{"gdk_device_set_key", &Gdk_device_set_key},
	{"gdk_device_set_mode", &Gdk_device_set_mode},
	{"gdk_device_set_source", &Gdk_device_set_source},
	{"gdk_devices_list", &Gdk_devices_list},
	{"gdk_display_add_client_message_filter", &Gdk_display_add_client_message_filter},
	{"gdk_display_beep", &Gdk_display_beep},
	{"gdk_display_close", &Gdk_display_close},
	{"gdk_display_flush", &Gdk_display_flush},
	{"gdk_display_get_core_pointer", &Gdk_display_get_core_pointer},
	{"gdk_display_get_default", &Gdk_display_get_default},
	{"gdk_display_get_default_cursor_size", &Gdk_display_get_default_cursor_size},
	{"gdk_display_get_default_group", &Gdk_display_get_default_group},
	{"gdk_display_get_default_screen", &Gdk_display_get_default_screen},
	{"gdk_display_get_event", &Gdk_display_get_event},
	{"gdk_display_get_maximal_cursor_size", &Gdk_display_get_maximal_cursor_size},
	{"gdk_display_get_n_screens", &Gdk_display_get_n_screens},
	{"gdk_display_get_name", &Gdk_display_get_name},
	{"gdk_display_get_pointer", &Gdk_display_get_pointer},
	{"gdk_display_get_screen", &Gdk_display_get_screen},
	{"gdk_display_get_type", &Gdk_display_get_type},
	{"gdk_display_get_window_at_pointer", &Gdk_display_get_window_at_pointer},
	{"gdk_display_is_closed", &Gdk_display_is_closed},
	{"gdk_display_keyboard_ungrab", &Gdk_display_keyboard_ungrab},
	{"gdk_display_list_devices", &Gdk_display_list_devices},
	{"gdk_display_manager_get", &Gdk_display_manager_get},
	{"gdk_display_manager_get_default_display", &Gdk_display_manager_get_default_display},
	{"gdk_display_manager_get_type", &Gdk_display_manager_get_type},
	{"gdk_display_manager_list_displays", &Gdk_display_manager_list_displays},
	{"gdk_display_manager_set_default_display", &Gdk_display_manager_set_default_display},
	{"gdk_display_open", &Gdk_display_open},
	{"gdk_display_open_default_libgtk_only", &Gdk_display_open_default_libgtk_only},
	{"gdk_display_peek_event", &Gdk_display_peek_event},
	{"gdk_display_pointer_is_grabbed", &Gdk_display_pointer_is_grabbed},
	{"gdk_display_pointer_ungrab", &Gdk_display_pointer_ungrab},
	{"gdk_display_put_event", &Gdk_display_put_event},
	{"gdk_display_request_selection_notification", &Gdk_display_request_selection_notification},
	{"gdk_display_set_double_click_distance", &Gdk_display_set_double_click_distance},
	{"gdk_display_set_double_click_time", &Gdk_display_set_double_click_time},
	{"gdk_display_set_pointer_hooks", &Gdk_display_set_pointer_hooks},
	{"gdk_display_store_clipboard", &Gdk_display_store_clipboard},
	{"gdk_display_supports_clipboard_persistence", &Gdk_display_supports_clipboard_persistence},
	{"gdk_display_supports_composite", &Gdk_display_supports_composite},
	{"gdk_display_supports_cursor_alpha", &Gdk_display_supports_cursor_alpha},
	{"gdk_display_supports_cursor_color", &Gdk_display_supports_cursor_color},
	{"gdk_display_supports_input_shapes", &Gdk_display_supports_input_shapes},
	{"gdk_display_supports_selection_notification", &Gdk_display_supports_selection_notification},
	{"gdk_display_supports_shapes", &Gdk_display_supports_shapes},
	{"gdk_display_sync", &Gdk_display_sync},
	{"gdk_display_warp_pointer", &Gdk_display_warp_pointer},
	{"gdk_drag_abort", &Gdk_drag_abort},
	{"gdk_drag_action_get_type", &Gdk_drag_action_get_type},
	{"gdk_drag_begin", &Gdk_drag_begin},
	{"gdk_drag_context_get_actions", &Gdk_drag_context_get_actions},
	{"gdk_drag_context_get_dest_window", &Gdk_drag_context_get_dest_window},
	{"gdk_drag_context_get_protocol", &Gdk_drag_context_get_protocol},
	{"gdk_drag_context_get_selected_action", &Gdk_drag_context_get_selected_action},
	{"gdk_drag_context_get_source_window", &Gdk_drag_context_get_source_window},
	{"gdk_drag_context_get_suggested_action", &Gdk_drag_context_get_suggested_action},
	{"gdk_drag_context_get_type", &Gdk_drag_context_get_type},
	{"gdk_drag_context_list_targets", &Gdk_drag_context_list_targets},
	{"gdk_drag_context_new", &Gdk_drag_context_new},
	{"gdk_drag_context_ref", &Gdk_drag_context_ref},
	{"gdk_drag_context_unref", &Gdk_drag_context_unref},
	{"gdk_drag_drop", &Gdk_drag_drop},
	{"gdk_drag_drop_succeeded", &Gdk_drag_drop_succeeded},
	{"gdk_drag_find_window", &Gdk_drag_find_window},
	{"gdk_drag_find_window_for_screen", &Gdk_drag_find_window_for_screen},
	{"gdk_drag_get_protocol", &Gdk_drag_get_protocol},
	{"gdk_drag_get_protocol_for_display", &Gdk_drag_get_protocol_for_display},
	{"gdk_drag_get_selection", &Gdk_drag_get_selection},
	{"gdk_drag_motion", &Gdk_drag_motion},
	{"gdk_drag_protocol_get_type", &Gdk_drag_protocol_get_type},
	{"gdk_drag_status", &Gdk_drag_status},
	{"gdk_draw_arc", &Gdk_draw_arc},
	{"gdk_draw_drawable", &Gdk_draw_drawable},
	{"gdk_draw_glyphs", &Gdk_draw_glyphs},
	{"gdk_draw_glyphs_transformed", &Gdk_draw_glyphs_transformed},
	{"gdk_draw_gray_image", &Gdk_draw_gray_image},
	{"gdk_draw_image", &Gdk_draw_image},
	{"gdk_draw_indexed_image", &Gdk_draw_indexed_image},
	{"gdk_draw_layout", &Gdk_draw_layout},
	{"gdk_draw_layout_line", &Gdk_draw_layout_line},
	{"gdk_draw_layout_line_with_colors", &Gdk_draw_layout_line_with_colors},
	{"gdk_draw_layout_with_colors", &Gdk_draw_layout_with_colors},
	{"gdk_draw_line", &Gdk_draw_line},
	{"gdk_draw_lines", &Gdk_draw_lines},
	{"gdk_draw_pixbuf", &Gdk_draw_pixbuf},
	{"gdk_draw_point", &Gdk_draw_point},
	{"gdk_draw_points", &Gdk_draw_points},
	{"gdk_draw_polygon", &Gdk_draw_polygon},
	{"gdk_draw_rectangle", &Gdk_draw_rectangle},
	{"gdk_draw_rgb_32_image", &Gdk_draw_rgb_32_image},
	{"gdk_draw_rgb_32_image_dithalign", &Gdk_draw_rgb_32_image_dithalign},
	{"gdk_draw_rgb_image", &Gdk_draw_rgb_image},
	{"gdk_draw_rgb_image_dithalign", &Gdk_draw_rgb_image_dithalign},
	{"gdk_draw_segments", &Gdk_draw_segments},
	{"gdk_draw_string", &Gdk_draw_string},
	{"gdk_draw_text", &Gdk_draw_text},
	{"gdk_draw_text_wc", &Gdk_draw_text_wc},
	{"gdk_draw_trapezoids", &Gdk_draw_trapezoids},
	{"gdk_drawable_copy_to_image", &Gdk_drawable_copy_to_image},
	{"gdk_drawable_get_clip_region", &Gdk_drawable_get_clip_region},
	{"gdk_drawable_get_colormap", &Gdk_drawable_get_colormap},
	{"gdk_drawable_get_data", &Gdk_drawable_get_data},
	{"gdk_drawable_get_depth", &Gdk_drawable_get_depth},
	{"gdk_drawable_get_display", &Gdk_drawable_get_display},
	{"gdk_drawable_get_image", &Gdk_drawable_get_image},
	{"gdk_drawable_get_screen", &Gdk_drawable_get_screen},
	{"gdk_drawable_get_size", &Gdk_drawable_get_size},
	{"gdk_drawable_get_type", &Gdk_drawable_get_type},
	{"gdk_drawable_get_visible_region", &Gdk_drawable_get_visible_region},
	{"gdk_drawable_get_visual", &Gdk_drawable_get_visual},
	{"gdk_drawable_ref", &Gdk_drawable_ref},
	{"gdk_drawable_set_colormap", &Gdk_drawable_set_colormap},
	{"gdk_drawable_set_data", &Gdk_drawable_set_data},
	{"gdk_drawable_unref", &Gdk_drawable_unref},
	{"gdk_drop_finish", &Gdk_drop_finish},
	{"gdk_drop_reply", &Gdk_drop_reply},
	{"gdk_error_trap_pop", &Gdk_error_trap_pop},
	{"gdk_error_trap_push", &Gdk_error_trap_push},
	{"gdk_event_copy", &Gdk_event_copy},
	{"gdk_event_free", &Gdk_event_free},
	{"gdk_event_get", &Gdk_event_get},
	{"gdk_event_get_axis", &Gdk_event_get_axis},
	{"gdk_event_get_coords", &Gdk_event_get_coords},
	{"gdk_event_get_graphics_expose", &Gdk_event_get_graphics_expose},
	{"gdk_event_get_root_coords", &Gdk_event_get_root_coords},
	{"gdk_event_get_screen", &Gdk_event_get_screen},
	{"gdk_event_get_state", &Gdk_event_get_state},
	{"gdk_event_get_time", &Gdk_event_get_time},
	{"gdk_event_get_type", &Gdk_event_get_type},
	{"gdk_event_handler_set", &Gdk_event_handler_set},
	{"gdk_event_mask_get_type", &Gdk_event_mask_get_type},
	{"gdk_event_new", &Gdk_event_new},
	{"gdk_event_peek", &Gdk_event_peek},
	{"gdk_event_put", &Gdk_event_put},
	{"gdk_event_request_motions", &Gdk_event_request_motions},
	{"gdk_event_send_client_message", &Gdk_event_send_client_message},
	{"gdk_event_send_client_message_for_display", &Gdk_event_send_client_message_for_display},
	{"gdk_event_send_clientmessage_toall", &Gdk_event_send_clientmessage_toall},
	{"gdk_event_set_screen", &Gdk_event_set_screen},
	{"gdk_event_type_get_type", &Gdk_event_type_get_type},
	{"gdk_events_pending", &Gdk_events_pending},
	{"gdk_exit", &Gdk_exit},
	{"gdk_extension_mode_get_type", &Gdk_extension_mode_get_type},
	{"gdk_fill_get_type", &Gdk_fill_get_type},
	{"gdk_fill_rule_get_type", &Gdk_fill_rule_get_type},
	{"gdk_filter_return_get_type", &Gdk_filter_return_get_type},
	{"gdk_flush", &Gdk_flush},
	{"gdk_font_equal", &Gdk_font_equal},
	{"gdk_font_from_description", &Gdk_font_from_description},
	{"gdk_font_from_description_for_display", &Gdk_font_from_description_for_display},
	{"gdk_font_get_display", &Gdk_font_get_display},
	{"gdk_font_get_type", &Gdk_font_get_type},
	{"gdk_font_id", &Gdk_font_id},
	{"gdk_font_load", &Gdk_font_load},
	{"gdk_font_load_for_display", &Gdk_font_load_for_display},
	{"gdk_font_ref", &Gdk_font_ref},
	{"gdk_font_type_get_type", &Gdk_font_type_get_type},
	{"gdk_font_unref", &Gdk_font_unref},
	{"gdk_fontset_load", &Gdk_fontset_load},
	{"gdk_fontset_load_for_display", &Gdk_fontset_load_for_display},
	{"gdk_free_compound_text", &Gdk_free_compound_text},
	{"gdk_free_text_list", &Gdk_free_text_list},
	{"gdk_function_get_type", &Gdk_function_get_type},
	{"gdk_gc_copy", &Gdk_gc_copy},
	{"gdk_gc_get_colormap", &Gdk_gc_get_colormap},
	{"gdk_gc_get_screen", &Gdk_gc_get_screen},
	{"gdk_gc_get_type", &Gdk_gc_get_type},
	{"gdk_gc_get_values", &Gdk_gc_get_values},
	{"gdk_gc_new", &Gdk_gc_new},
	{"gdk_gc_new_with_values", &Gdk_gc_new_with_values},
	{"gdk_gc_offset", &Gdk_gc_offset},
	{"gdk_gc_ref", &Gdk_gc_ref},
	{"gdk_gc_set_background", &Gdk_gc_set_background},
	{"gdk_gc_set_clip_mask", &Gdk_gc_set_clip_mask},
	{"gdk_gc_set_clip_origin", &Gdk_gc_set_clip_origin},
	{"gdk_gc_set_clip_rectangle", &Gdk_gc_set_clip_rectangle},
	{"gdk_gc_set_clip_region", &Gdk_gc_set_clip_region},
	{"gdk_gc_set_colormap", &Gdk_gc_set_colormap},
	{"gdk_gc_set_dashes", &Gdk_gc_set_dashes},
	{"gdk_gc_set_exposures", &Gdk_gc_set_exposures},
	{"gdk_gc_set_fill", &Gdk_gc_set_fill},
	{"gdk_gc_set_font", &Gdk_gc_set_font},
	{"gdk_gc_set_foreground", &Gdk_gc_set_foreground},
	{"gdk_gc_set_function", &Gdk_gc_set_function},
	{"gdk_gc_set_line_attributes", &Gdk_gc_set_line_attributes},
	{"gdk_gc_set_rgb_bg_color", &Gdk_gc_set_rgb_bg_color},
	{"gdk_gc_set_rgb_fg_color", &Gdk_gc_set_rgb_fg_color},
	{"gdk_gc_set_stipple", &Gdk_gc_set_stipple},
	{"gdk_gc_set_subwindow", &Gdk_gc_set_subwindow},
	{"gdk_gc_set_tile", &Gdk_gc_set_tile},
	{"gdk_gc_set_ts_origin", &Gdk_gc_set_ts_origin},
	{"gdk_gc_set_values", &Gdk_gc_set_values},
	{"gdk_gc_unref", &Gdk_gc_unref},
	{"gdk_gc_values_mask_get_type", &Gdk_gc_values_mask_get_type},
	{"gdk_get_default_root_window", &Gdk_get_default_root_window},
	{"gdk_get_display", &Gdk_get_display},
	{"gdk_get_display_arg_name", &Gdk_get_display_arg_name},
	{"gdk_get_program_class", &Gdk_get_program_class},
	{"gdk_get_show_events", &Gdk_get_show_events},
	{"gdk_get_use_xshm", &Gdk_get_use_xshm},
	{"gdk_grab_status_get_type", &Gdk_grab_status_get_type},
	{"gdk_gravity_get_type", &Gdk_gravity_get_type},
	{"gdk_image_get", &Gdk_image_get},
	{"gdk_image_get_bits_per_pixel", &Gdk_image_get_bits_per_pixel},
	{"gdk_image_get_byte_order", &Gdk_image_get_byte_order},
	{"gdk_image_get_bytes_per_line", &Gdk_image_get_bytes_per_line},
	{"gdk_image_get_bytes_per_pixel", &Gdk_image_get_bytes_per_pixel},
	{"gdk_image_get_colormap", &Gdk_image_get_colormap},
	{"gdk_image_get_depth", &Gdk_image_get_depth},
	{"gdk_image_get_height", &Gdk_image_get_height},
	{"gdk_image_get_image_type", &Gdk_image_get_image_type},
	{"gdk_image_get_pixel", &Gdk_image_get_pixel},
	{"gdk_image_get_pixels", &Gdk_image_get_pixels},
	{"gdk_image_get_type", &Gdk_image_get_type},
	{"gdk_image_get_visual", &Gdk_image_get_visual},
	{"gdk_image_get_width", &Gdk_image_get_width},
	{"gdk_image_new", &Gdk_image_new},
	{"gdk_image_put_pixel", &Gdk_image_put_pixel},
	{"gdk_image_ref", &Gdk_image_ref},
	{"gdk_image_set_colormap", &Gdk_image_set_colormap},
	{"gdk_image_type_get_type", &Gdk_image_type_get_type},
	{"gdk_image_unref", &Gdk_image_unref},
	{"gdk_init", &Gdk_init},
	{"gdk_init_check", &Gdk_init_check},
	{"gdk_input_add", &Gdk_input_add},
	{"gdk_input_add_full", &Gdk_input_add_full},
	{"gdk_input_condition_get_type", &Gdk_input_condition_get_type},
	{"gdk_input_mode_get_type", &Gdk_input_mode_get_type},
	{"gdk_input_remove", &Gdk_input_remove},
	{"gdk_input_set_extension_events", &Gdk_input_set_extension_events},
	{"gdk_input_source_get_type", &Gdk_input_source_get_type},
	{"gdk_join_style_get_type", &Gdk_join_style_get_type},
	{"gdk_keyboard_grab", &Gdk_keyboard_grab},
	{"gdk_keyboard_grab_info_libgtk_only", &Gdk_keyboard_grab_info_libgtk_only},
	{"gdk_keyboard_ungrab", &Gdk_keyboard_ungrab},
	{"gdk_keymap_add_virtual_modifiers", &Gdk_keymap_add_virtual_modifiers},
	{"gdk_keymap_get_caps_lock_state", &Gdk_keymap_get_caps_lock_state},
	{"gdk_keymap_get_default", &Gdk_keymap_get_default},
	{"gdk_keymap_get_direction", &Gdk_keymap_get_direction},
	{"gdk_keymap_get_entries_for_keycode", &Gdk_keymap_get_entries_for_keycode},
	{"gdk_keymap_get_entries_for_keyval", &Gdk_keymap_get_entries_for_keyval},
	{"gdk_keymap_get_for_display", &Gdk_keymap_get_for_display},
	{"gdk_keymap_get_type", &Gdk_keymap_get_type},
	{"gdk_keymap_have_bidi_layouts", &Gdk_keymap_have_bidi_layouts},
	{"gdk_keymap_lookup_key", &Gdk_keymap_lookup_key},
	{"gdk_keymap_map_virtual_modifiers", &Gdk_keymap_map_virtual_modifiers},
	{"gdk_keymap_translate_keyboard_state", &Gdk_keymap_translate_keyboard_state},
	{"gdk_keyval_convert_case", &Gdk_keyval_convert_case},
	{"gdk_keyval_from_name", &Gdk_keyval_from_name},
	{"gdk_keyval_is_lower", &Gdk_keyval_is_lower},
	{"gdk_keyval_is_upper", &Gdk_keyval_is_upper},
	{"gdk_keyval_name", &Gdk_keyval_name},
	{"gdk_keyval_to_lower", &Gdk_keyval_to_lower},
	{"gdk_keyval_to_unicode", &Gdk_keyval_to_unicode},
	{"gdk_keyval_to_upper", &Gdk_keyval_to_upper},
	{"gdk_line_style_get_type", &Gdk_line_style_get_type},
	{"gdk_list_visuals", &Gdk_list_visuals},
	{"gdk_mbstowcs", &Gdk_mbstowcs},
	{"gdk_modifier_type_get_type", &Gdk_modifier_type_get_type},
	// Undocumented {"gdk_net_wm_supports", &Gdk_net_wm_supports},
	{"gdk_notify_startup_complete", &Gdk_notify_startup_complete},
	{"gdk_notify_startup_complete_with_id", &Gdk_notify_startup_complete_with_id},
	{"gdk_notify_type_get_type", &Gdk_notify_type_get_type},
	{"gdk_offscreen_window_get_embedder", &Gdk_offscreen_window_get_embedder},
	{"gdk_offscreen_window_get_pixmap", &Gdk_offscreen_window_get_pixmap},
	// Undocumented {"gdk_offscreen_window_get_type", &Gdk_offscreen_window_get_type},
	{"gdk_offscreen_window_set_embedder", &Gdk_offscreen_window_set_embedder},
	{"gdk_overlap_type_get_type", &Gdk_overlap_type_get_type},
	{"gdk_owner_change_get_type", &Gdk_owner_change_get_type},
	{"gdk_pango_attr_emboss_color_new", &Gdk_pango_attr_emboss_color_new},
	{"gdk_pango_attr_embossed_new", &Gdk_pango_attr_embossed_new},
	{"gdk_pango_attr_stipple_new", &Gdk_pango_attr_stipple_new},
	{"gdk_pango_context_get", &Gdk_pango_context_get},
	{"gdk_pango_context_get_for_screen", &Gdk_pango_context_get_for_screen},
	{"gdk_pango_context_set_colormap", &Gdk_pango_context_set_colormap},
	{"gdk_pango_layout_get_clip_region", &Gdk_pango_layout_get_clip_region},
	{"gdk_pango_layout_line_get_clip_region", &Gdk_pango_layout_line_get_clip_region},
	{"gdk_pango_renderer_get_default", &Gdk_pango_renderer_get_default},
	{"gdk_pango_renderer_get_type", &Gdk_pango_renderer_get_type},
	{"gdk_pango_renderer_new", &Gdk_pango_renderer_new},
	{"gdk_pango_renderer_set_drawable", &Gdk_pango_renderer_set_drawable},
	{"gdk_pango_renderer_set_gc", &Gdk_pango_renderer_set_gc},
	{"gdk_pango_renderer_set_override_color", &Gdk_pango_renderer_set_override_color},
	{"gdk_pango_renderer_set_stipple", &Gdk_pango_renderer_set_stipple},
	{"gdk_parse_args", &Gdk_parse_args},
	{"gdk_pixbuf_get_from_drawable", &Gdk_pixbuf_get_from_drawable},
	{"gdk_pixbuf_get_from_image", &Gdk_pixbuf_get_from_image},
	{"gdk_pixbuf_render_pixmap_and_mask", &Gdk_pixbuf_render_pixmap_and_mask},
	{"gdk_pixbuf_render_pixmap_and_mask_for_colormap", &Gdk_pixbuf_render_pixmap_and_mask_for_colormap},
	{"gdk_pixbuf_render_threshold_alpha", &Gdk_pixbuf_render_threshold_alpha},
	{"gdk_pixbuf_render_to_drawable", &Gdk_pixbuf_render_to_drawable},
	{"gdk_pixbuf_render_to_drawable_alpha", &Gdk_pixbuf_render_to_drawable_alpha},
	{"gdk_pixmap_colormap_create_from_xpm", &Gdk_pixmap_colormap_create_from_xpm},
	{"gdk_pixmap_colormap_create_from_xpm_d", &Gdk_pixmap_colormap_create_from_xpm_d},
	{"gdk_pixmap_create_from_data", &Gdk_pixmap_create_from_data},
	{"gdk_pixmap_create_from_xpm", &Gdk_pixmap_create_from_xpm},
	{"gdk_pixmap_create_from_xpm_d", &Gdk_pixmap_create_from_xpm_d},
	{"gdk_pixmap_foreign_new", &Gdk_pixmap_foreign_new},
	{"gdk_pixmap_foreign_new_for_display", &Gdk_pixmap_foreign_new_for_display},
	{"gdk_pixmap_foreign_new_for_screen", &Gdk_pixmap_foreign_new_for_screen},
	{"gdk_pixmap_get_size", &Gdk_pixmap_get_size},
	{"gdk_pixmap_get_type", &Gdk_pixmap_get_type},
	{"gdk_pixmap_lookup", &Gdk_pixmap_lookup},
	{"gdk_pixmap_lookup_for_display", &Gdk_pixmap_lookup_for_display},
	{"gdk_pixmap_new", &Gdk_pixmap_new},
	{"gdk_pointer_grab", &Gdk_pointer_grab},
	{"gdk_pointer_grab_info_libgtk_only", &Gdk_pointer_grab_info_libgtk_only},
	{"gdk_pointer_is_grabbed", &Gdk_pointer_is_grabbed},
	{"gdk_pointer_ungrab", &Gdk_pointer_ungrab},
	{"gdk_pre_parse_libgtk_only", &Gdk_pre_parse_libgtk_only},
	{"gdk_prop_mode_get_type", &Gdk_prop_mode_get_type},
	{"gdk_property_change", &Gdk_property_change},
	{"gdk_property_delete", &Gdk_property_delete},
	{"gdk_property_get", &Gdk_property_get},
	{"gdk_property_state_get_type", &Gdk_property_state_get_type},
	{"gdk_query_depths", &Gdk_query_depths},
	{"gdk_query_visual_types", &Gdk_query_visual_types},
	{"gdk_rectangle_get_type", &Gdk_rectangle_get_type},
	{"gdk_rectangle_intersect", &Gdk_rectangle_intersect},
	{"gdk_rectangle_union", &Gdk_rectangle_union},
	{"gdk_region_copy", &Gdk_region_copy},
	{"gdk_region_destroy", &Gdk_region_destroy},
	{"gdk_region_empty", &Gdk_region_empty},
	{"gdk_region_equal", &Gdk_region_equal},
	{"gdk_region_get_clipbox", &Gdk_region_get_clipbox},
	{"gdk_region_get_rectangles", &Gdk_region_get_rectangles},
	{"gdk_region_intersect", &Gdk_region_intersect},
	{"gdk_region_new", &Gdk_region_new},
	{"gdk_region_offset", &Gdk_region_offset},
	{"gdk_region_point_in", &Gdk_region_point_in},
	{"gdk_region_polygon", &Gdk_region_polygon},
	{"gdk_region_rect_equal", &Gdk_region_rect_equal},
	{"gdk_region_rect_in", &Gdk_region_rect_in},
	{"gdk_region_rectangle", &Gdk_region_rectangle},
	{"gdk_region_shrink", &Gdk_region_shrink},
	{"gdk_region_spans_intersect_foreach", &Gdk_region_spans_intersect_foreach},
	{"gdk_region_subtract", &Gdk_region_subtract},
	{"gdk_region_union", &Gdk_region_union},
	{"gdk_region_union_with_rect", &Gdk_region_union_with_rect},
	{"gdk_region_xor", &Gdk_region_xor},
	{"gdk_rgb_cmap_free", &Gdk_rgb_cmap_free},
	{"gdk_rgb_cmap_new", &Gdk_rgb_cmap_new},
	{"gdk_rgb_colormap_ditherable", &Gdk_rgb_colormap_ditherable},
	{"gdk_rgb_dither_get_type", &Gdk_rgb_dither_get_type},
	{"gdk_rgb_ditherable", &Gdk_rgb_ditherable},
	{"gdk_rgb_find_color", &Gdk_rgb_find_color},
	{"gdk_rgb_gc_set_background", &Gdk_rgb_gc_set_background},
	{"gdk_rgb_gc_set_foreground", &Gdk_rgb_gc_set_foreground},
	{"gdk_rgb_get_colormap", &Gdk_rgb_get_colormap},
	{"gdk_rgb_get_visual", &Gdk_rgb_get_visual},
	{"gdk_rgb_init", &Gdk_rgb_init},
	{"gdk_rgb_set_install", &Gdk_rgb_set_install},
	{"gdk_rgb_set_min_colors", &Gdk_rgb_set_min_colors},
	{"gdk_rgb_set_verbose", &Gdk_rgb_set_verbose},
	{"gdk_rgb_xpixel_from_rgb", &Gdk_rgb_xpixel_from_rgb},
	{"gdk_screen_broadcast_client_message", &Gdk_screen_broadcast_client_message},
	{"gdk_screen_get_active_window", &Gdk_screen_get_active_window},
	{"gdk_screen_get_default", &Gdk_screen_get_default},
	{"gdk_screen_get_default_colormap", &Gdk_screen_get_default_colormap},
	{"gdk_screen_get_display", &Gdk_screen_get_display},
	{"gdk_screen_get_font_options", &Gdk_screen_get_font_options},
	{"gdk_screen_get_height", &Gdk_screen_get_height},
	{"gdk_screen_get_height_mm", &Gdk_screen_get_height_mm},
	{"gdk_screen_get_monitor_at_point", &Gdk_screen_get_monitor_at_point},
	{"gdk_screen_get_monitor_at_window", &Gdk_screen_get_monitor_at_window},
	{"gdk_screen_get_monitor_geometry", &Gdk_screen_get_monitor_geometry},
	{"gdk_screen_get_monitor_height_mm", &Gdk_screen_get_monitor_height_mm},
	{"gdk_screen_get_monitor_plug_name", &Gdk_screen_get_monitor_plug_name},
	{"gdk_screen_get_monitor_width_mm", &Gdk_screen_get_monitor_width_mm},
	{"gdk_screen_get_n_monitors", &Gdk_screen_get_n_monitors},
	{"gdk_screen_get_number", &Gdk_screen_get_number},
	{"gdk_screen_get_primary_monitor", &Gdk_screen_get_primary_monitor},
	{"gdk_screen_get_resolution", &Gdk_screen_get_resolution},
	{"gdk_screen_get_rgb_colormap", &Gdk_screen_get_rgb_colormap},
	{"gdk_screen_get_rgb_visual", &Gdk_screen_get_rgb_visual},
	{"gdk_screen_get_rgba_colormap", &Gdk_screen_get_rgba_colormap},
	{"gdk_screen_get_rgba_visual", &Gdk_screen_get_rgba_visual},
	{"gdk_screen_get_root_window", &Gdk_screen_get_root_window},
	{"gdk_screen_get_setting", &Gdk_screen_get_setting},
	{"gdk_screen_get_system_colormap", &Gdk_screen_get_system_colormap},
	{"gdk_screen_get_system_visual", &Gdk_screen_get_system_visual},
	{"gdk_screen_get_toplevel_windows", &Gdk_screen_get_toplevel_windows},
	{"gdk_screen_get_type", &Gdk_screen_get_type},
	{"gdk_screen_get_width", &Gdk_screen_get_width},
	{"gdk_screen_get_width_mm", &Gdk_screen_get_width_mm},
	{"gdk_screen_get_window_stack", &Gdk_screen_get_window_stack},
	{"gdk_screen_height", &Gdk_screen_height},
	{"gdk_screen_height_mm", &Gdk_screen_height_mm},
	{"gdk_screen_is_composited", &Gdk_screen_is_composited},
	{"gdk_screen_list_visuals", &Gdk_screen_list_visuals},
	{"gdk_screen_make_display_name", &Gdk_screen_make_display_name},
	{"gdk_screen_set_default_colormap", &Gdk_screen_set_default_colormap},
	{"gdk_screen_set_font_options", &Gdk_screen_set_font_options},
	{"gdk_screen_set_resolution", &Gdk_screen_set_resolution},
	{"gdk_screen_width", &Gdk_screen_width},
	{"gdk_screen_width_mm", &Gdk_screen_width_mm},
	{"gdk_scroll_direction_get_type", &Gdk_scroll_direction_get_type},
	{"gdk_selection_convert", &Gdk_selection_convert},
	{"gdk_selection_owner_get", &Gdk_selection_owner_get},
	{"gdk_selection_owner_get_for_display", &Gdk_selection_owner_get_for_display},
	{"gdk_selection_owner_set", &Gdk_selection_owner_set},
	{"gdk_selection_owner_set_for_display", &Gdk_selection_owner_set_for_display},
	{"gdk_selection_property_get", &Gdk_selection_property_get},
	{"gdk_selection_send_notify", &Gdk_selection_send_notify},
	{"gdk_selection_send_notify_for_display", &Gdk_selection_send_notify_for_display},
	{"gdk_set_double_click_time", &Gdk_set_double_click_time},
	{"gdk_set_locale", &Gdk_set_locale},
	{"gdk_set_pointer_hooks", &Gdk_set_pointer_hooks},
	{"gdk_set_program_class", &Gdk_set_program_class},
	{"gdk_set_show_events", &Gdk_set_show_events},
	{"gdk_set_sm_client_id", &Gdk_set_sm_client_id},
	{"gdk_set_use_xshm", &Gdk_set_use_xshm},
	{"gdk_setting_action_get_type", &Gdk_setting_action_get_type},
	{"gdk_setting_get", &Gdk_setting_get},
	{"gdk_spawn_command_line_on_screen", &Gdk_spawn_command_line_on_screen},
	{"gdk_spawn_on_screen", &Gdk_spawn_on_screen},
	{"gdk_spawn_on_screen_with_pipes", &Gdk_spawn_on_screen_with_pipes},
	{"gdk_status_get_type", &Gdk_status_get_type},
	{"gdk_string_extents", &Gdk_string_extents},
	{"gdk_string_height", &Gdk_string_height},
	{"gdk_string_measure", &Gdk_string_measure},
	{"gdk_string_to_compound_text", &Gdk_string_to_compound_text},
	{"gdk_string_to_compound_text_for_display", &Gdk_string_to_compound_text_for_display},
	{"gdk_string_width", &Gdk_string_width},
	{"gdk_subwindow_mode_get_type", &Gdk_subwindow_mode_get_type},
	{"gdk_synthesize_window_state", &Gdk_synthesize_window_state},
	{"gdk_test_render_sync", &Gdk_test_render_sync},
	{"gdk_test_simulate_button", &Gdk_test_simulate_button},
	{"gdk_test_simulate_key", &Gdk_test_simulate_key},
	{"gdk_text_extents", &Gdk_text_extents},
	{"gdk_text_extents_wc", &Gdk_text_extents_wc},
	{"gdk_text_height", &Gdk_text_height},
	{"gdk_text_measure", &Gdk_text_measure},
	{"gdk_text_property_to_text_list", &Gdk_text_property_to_text_list},
	{"gdk_text_property_to_text_list_for_display", &Gdk_text_property_to_text_list_for_display},
	{"gdk_text_property_to_utf8_list", &Gdk_text_property_to_utf8_list},
	{"gdk_text_property_to_utf8_list_for_display", &Gdk_text_property_to_utf8_list_for_display},
	{"gdk_text_width", &Gdk_text_width},
	{"gdk_text_width_wc", &Gdk_text_width_wc},
	{"gdk_threads_add_idle", &Gdk_threads_add_idle},
	{"gdk_threads_add_idle_full", &Gdk_threads_add_idle_full},
	{"gdk_threads_add_timeout", &Gdk_threads_add_timeout},
	{"gdk_threads_add_timeout_full", &Gdk_threads_add_timeout_full},
	{"gdk_threads_add_timeout_seconds", &Gdk_threads_add_timeout_seconds},
	{"gdk_threads_add_timeout_seconds_full", &Gdk_threads_add_timeout_seconds_full},
	{"gdk_threads_enter", &Gdk_threads_enter},
	{"gdk_threads_init", &Gdk_threads_init},
	{"gdk_threads_leave", &Gdk_threads_leave},
	{"gdk_threads_set_lock_functions", &Gdk_threads_set_lock_functions},
	{"gdk_unicode_to_keyval", &Gdk_unicode_to_keyval},
	{"gdk_utf8_to_compound_text", &Gdk_utf8_to_compound_text},
	{"gdk_utf8_to_compound_text_for_display", &Gdk_utf8_to_compound_text_for_display},
	{"gdk_utf8_to_string_target", &Gdk_utf8_to_string_target},
	{"gdk_visibility_state_get_type", &Gdk_visibility_state_get_type},
	{"gdk_visual_get_best", &Gdk_visual_get_best},
	{"gdk_visual_get_best_depth", &Gdk_visual_get_best_depth},
	{"gdk_visual_get_best_type", &Gdk_visual_get_best_type},
	{"gdk_visual_get_best_with_both", &Gdk_visual_get_best_with_both},
	{"gdk_visual_get_best_with_depth", &Gdk_visual_get_best_with_depth},
	{"gdk_visual_get_best_with_type", &Gdk_visual_get_best_with_type},
	{"gdk_visual_get_bits_per_rgb", &Gdk_visual_get_bits_per_rgb},
	{"gdk_visual_get_blue_pixel_details", &Gdk_visual_get_blue_pixel_details},
	{"gdk_visual_get_byte_order", &Gdk_visual_get_byte_order},
	{"gdk_visual_get_colormap_size", &Gdk_visual_get_colormap_size},
	{"gdk_visual_get_depth", &Gdk_visual_get_depth},
	{"gdk_visual_get_green_pixel_details", &Gdk_visual_get_green_pixel_details},
	{"gdk_visual_get_red_pixel_details", &Gdk_visual_get_red_pixel_details},
	{"gdk_visual_get_screen", &Gdk_visual_get_screen},
	{"gdk_visual_get_system", &Gdk_visual_get_system},
	{"gdk_visual_get_type", &Gdk_visual_get_type},
	{"gdk_visual_get_visual_type", &Gdk_visual_get_visual_type},
	{"gdk_visual_type_get_type", &Gdk_visual_type_get_type},
	{"gdk_wcstombs", &Gdk_wcstombs},
	{"gdk_win32_begin_direct_draw_libgtk_only", &Gdk_win32_begin_direct_draw_libgtk_only},
	{"gdk_win32_drawable_get_handle", &Gdk_win32_drawable_get_handle},
	{"gdk_win32_end_direct_draw_libgtk_only", &Gdk_win32_end_direct_draw_libgtk_only},
	{"gdk_win32_handle_table_lookup", &Gdk_win32_handle_table_lookup},
	{"gdk_win32_hdc_get", &Gdk_win32_hdc_get},
	{"gdk_win32_hdc_release", &Gdk_win32_hdc_release},
	{"gdk_win32_icon_to_pixbuf_libgtk_only", &Gdk_win32_icon_to_pixbuf_libgtk_only},
	{"gdk_win32_pixbuf_to_hicon_libgtk_only", &Gdk_win32_pixbuf_to_hicon_libgtk_only},
	{"gdk_win32_selection_add_targets", &Gdk_win32_selection_add_targets},
	{"gdk_win32_set_modal_dialog_libgtk_only", &Gdk_win32_set_modal_dialog_libgtk_only},
	{"gdk_win32_window_foreign_new_for_display", &Gdk_win32_window_foreign_new_for_display},
	{"gdk_win32_window_get_impl_hwnd", &Gdk_win32_window_get_impl_hwnd},
	{"gdk_win32_window_is_win32", &Gdk_win32_window_is_win32},
	{"gdk_win32_window_lookup_for_display", &Gdk_win32_window_lookup_for_display},
	{"gdk_window_add_filter", &Gdk_window_add_filter},
	{"gdk_window_at_pointer", &Gdk_window_at_pointer},
	{"gdk_window_attributes_type_get_type", &Gdk_window_attributes_type_get_type},
	{"gdk_window_beep", &Gdk_window_beep},
	{"gdk_window_begin_move_drag", &Gdk_window_begin_move_drag},
	{"gdk_window_begin_paint_rect", &Gdk_window_begin_paint_rect},
	{"gdk_window_begin_paint_region", &Gdk_window_begin_paint_region},
	{"gdk_window_begin_resize_drag", &Gdk_window_begin_resize_drag},
	{"gdk_window_class_get_type", &Gdk_window_class_get_type},
	{"gdk_window_clear", &Gdk_window_clear},
	{"gdk_window_clear_area", &Gdk_window_clear_area},
	{"gdk_window_clear_area_e", &Gdk_window_clear_area_e},
	{"gdk_window_configure_finished", &Gdk_window_configure_finished},
	{"gdk_window_constrain_size", &Gdk_window_constrain_size},
	{"gdk_window_coords_from_parent", &Gdk_window_coords_from_parent},
	{"gdk_window_coords_to_parent", &Gdk_window_coords_to_parent},
	{"gdk_window_create_similar_surface", &Gdk_window_create_similar_surface},
	{"gdk_window_deiconify", &Gdk_window_deiconify},
	{"gdk_window_destroy", &Gdk_window_destroy},
	{"gdk_window_destroy_notify", &Gdk_window_destroy_notify},
	{"gdk_window_edge_get_type", &Gdk_window_edge_get_type},
	{"gdk_window_enable_synchronized_configure", &Gdk_window_enable_synchronized_configure},
	{"gdk_window_end_paint", &Gdk_window_end_paint},
	{"gdk_window_ensure_native", &Gdk_window_ensure_native},
	{"gdk_window_flush", &Gdk_window_flush},
	{"gdk_window_focus", &Gdk_window_focus},
	{"gdk_window_foreign_new", &Gdk_window_foreign_new},
	{"gdk_window_foreign_new_for_display", &Gdk_window_foreign_new_for_display},
	{"gdk_window_freeze_toplevel_updates_libgtk_only", &Gdk_window_freeze_toplevel_updates_libgtk_only},
	{"gdk_window_freeze_updates", &Gdk_window_freeze_updates},
	{"gdk_window_fullscreen", &Gdk_window_fullscreen},
	{"gdk_window_geometry_changed", &Gdk_window_geometry_changed},
	{"gdk_window_get_accept_focus", &Gdk_window_get_accept_focus},
	{"gdk_window_get_background_pattern", &Gdk_window_get_background_pattern},
	{"gdk_window_get_children", &Gdk_window_get_children},
	{"gdk_window_get_composited", &Gdk_window_get_composited},
	{"gdk_window_get_cursor", &Gdk_window_get_cursor},
	{"gdk_window_get_decorations", &Gdk_window_get_decorations},
	{"gdk_window_get_deskrelative_origin", &Gdk_window_get_deskrelative_origin},
	{"gdk_window_get_display", &Gdk_window_get_display},
	{"gdk_window_get_effective_parent", &Gdk_window_get_effective_parent},
	{"gdk_window_get_effective_toplevel", &Gdk_window_get_effective_toplevel},
	{"gdk_window_get_events", &Gdk_window_get_events},
	{"gdk_window_get_focus_on_map", &Gdk_window_get_focus_on_map},
	{"gdk_window_get_frame_extents", &Gdk_window_get_frame_extents},
	{"gdk_window_get_geometry", &Gdk_window_get_geometry},
	{"gdk_window_get_group", &Gdk_window_get_group},
	{"gdk_window_get_height", &Gdk_window_get_height},
	{"gdk_window_get_internal_paint_info", &Gdk_window_get_internal_paint_info},
	{"gdk_window_get_modal_hint", &Gdk_window_get_modal_hint},
	{"gdk_window_get_origin", &Gdk_window_get_origin},
	{"gdk_window_get_parent", &Gdk_window_get_parent},
	{"gdk_window_get_pointer", &Gdk_window_get_pointer},
	{"gdk_window_get_position", &Gdk_window_get_position},
	{"gdk_window_get_root_coords", &Gdk_window_get_root_coords},
	{"gdk_window_get_root_origin", &Gdk_window_get_root_origin},
	{"gdk_window_get_screen", &Gdk_window_get_screen},
	{"gdk_window_get_state", &Gdk_window_get_state},
	{"gdk_window_get_toplevel", &Gdk_window_get_toplevel},
	{"gdk_window_get_toplevels", &Gdk_window_get_toplevels},
	{"gdk_window_get_type_hint", &Gdk_window_get_type_hint},
	{"gdk_window_get_update_area", &Gdk_window_get_update_area},
	{"gdk_window_get_user_data", &Gdk_window_get_user_data},
	{"gdk_window_get_visual", &Gdk_window_get_visual},
	{"gdk_window_get_width", &Gdk_window_get_width},
	{"gdk_window_get_window_type", &Gdk_window_get_window_type},
	{"gdk_window_has_native", &Gdk_window_has_native},
	{"gdk_window_hide", &Gdk_window_hide},
	{"gdk_window_hints_get_type", &Gdk_window_hints_get_type},
	{"gdk_window_iconify", &Gdk_window_iconify},
	// Undocumented {"gdk_window_impl_get_type", &Gdk_window_impl_get_type},
	{"gdk_window_input_shape_combine_mask", &Gdk_window_input_shape_combine_mask},
	{"gdk_window_input_shape_combine_region", &Gdk_window_input_shape_combine_region},
	{"gdk_window_invalidate_maybe_recurse", &Gdk_window_invalidate_maybe_recurse},
	{"gdk_window_invalidate_rect", &Gdk_window_invalidate_rect},
	{"gdk_window_invalidate_region", &Gdk_window_invalidate_region},
	{"gdk_window_is_destroyed", &Gdk_window_is_destroyed},
	{"gdk_window_is_input_only", &Gdk_window_is_input_only},
	{"gdk_window_is_shaped", &Gdk_window_is_shaped},
	{"gdk_window_is_viewable", &Gdk_window_is_viewable},
	{"gdk_window_is_visible", &Gdk_window_is_visible},
	{"gdk_window_lookup", &Gdk_window_lookup},
	{"gdk_window_lookup_for_display", &Gdk_window_lookup_for_display},
	{"gdk_window_lower", &Gdk_window_lower},
	{"gdk_window_maximize", &Gdk_window_maximize},
	{"gdk_window_merge_child_input_shapes", &Gdk_window_merge_child_input_shapes},
	{"gdk_window_merge_child_shapes", &Gdk_window_merge_child_shapes},
	{"gdk_window_move", &Gdk_window_move},
	{"gdk_window_move_region", &Gdk_window_move_region},
	{"gdk_window_move_resize", &Gdk_window_move_resize},
	{"gdk_window_new", &Gdk_window_new},
	{"gdk_window_object_get_type", &Gdk_window_object_get_type},
	{"gdk_window_peek_children", &Gdk_window_peek_children},
	{"gdk_window_process_all_updates", &Gdk_window_process_all_updates},
	{"gdk_window_process_updates", &Gdk_window_process_updates},
	{"gdk_window_raise", &Gdk_window_raise},
	{"gdk_window_redirect_to_drawable", &Gdk_window_redirect_to_drawable},
	{"gdk_window_register_dnd", &Gdk_window_register_dnd},
	{"gdk_window_remove_filter", &Gdk_window_remove_filter},
	{"gdk_window_remove_redirection", &Gdk_window_remove_redirection},
	{"gdk_window_reparent", &Gdk_window_reparent},
	{"gdk_window_resize", &Gdk_window_resize},
	{"gdk_window_restack", &Gdk_window_restack},
	{"gdk_window_scroll", &Gdk_window_scroll},
	{"gdk_window_set_accept_focus", &Gdk_window_set_accept_focus},
	{"gdk_window_set_back_pixmap", &Gdk_window_set_back_pixmap},
	{"gdk_window_set_background", &Gdk_window_set_background},
	{"gdk_window_set_child_input_shapes", &Gdk_window_set_child_input_shapes},
	{"gdk_window_set_child_shapes", &Gdk_window_set_child_shapes},
	{"gdk_window_set_composited", &Gdk_window_set_composited},
	{"gdk_window_set_cursor", &Gdk_window_set_cursor},
	{"gdk_window_set_debug_updates", &Gdk_window_set_debug_updates},
	{"gdk_window_set_decorations", &Gdk_window_set_decorations},
	{"gdk_window_set_events", &Gdk_window_set_events},
	{"gdk_window_set_focus_on_map", &Gdk_window_set_focus_on_map},
	{"gdk_window_set_functions", &Gdk_window_set_functions},
	{"gdk_window_set_geometry_hints", &Gdk_window_set_geometry_hints},
	{"gdk_window_set_group", &Gdk_window_set_group},
	{"gdk_window_set_hints", &Gdk_window_set_hints},
	{"gdk_window_set_icon", &Gdk_window_set_icon},
	{"gdk_window_set_icon_list", &Gdk_window_set_icon_list},
	{"gdk_window_set_icon_name", &Gdk_window_set_icon_name},
	{"gdk_window_set_keep_above", &Gdk_window_set_keep_above},
	{"gdk_window_set_keep_below", &Gdk_window_set_keep_below},
	{"gdk_window_set_modal_hint", &Gdk_window_set_modal_hint},
	{"gdk_window_set_opacity", &Gdk_window_set_opacity},
	{"gdk_window_set_override_redirect", &Gdk_window_set_override_redirect},
	{"gdk_window_set_role", &Gdk_window_set_role},
	{"gdk_window_set_skip_pager_hint", &Gdk_window_set_skip_pager_hint},
	{"gdk_window_set_skip_taskbar_hint", &Gdk_window_set_skip_taskbar_hint},
	{"gdk_window_set_startup_id", &Gdk_window_set_startup_id},
	{"gdk_window_set_static_gravities", &Gdk_window_set_static_gravities},
	{"gdk_window_set_title", &Gdk_window_set_title},
	{"gdk_window_set_transient_for", &Gdk_window_set_transient_for},
	{"gdk_window_set_type_hint", &Gdk_window_set_type_hint},
	{"gdk_window_set_urgency_hint", &Gdk_window_set_urgency_hint},
	{"gdk_window_set_user_data", &Gdk_window_set_user_data},
	{"gdk_window_shape_combine_mask", &Gdk_window_shape_combine_mask},
	{"gdk_window_shape_combine_region", &Gdk_window_shape_combine_region},
	{"gdk_window_show", &Gdk_window_show},
	{"gdk_window_show_unraised", &Gdk_window_show_unraised},
	{"gdk_window_state_get_type", &Gdk_window_state_get_type},
	{"gdk_window_stick", &Gdk_window_stick},
	{"gdk_window_thaw_toplevel_updates_libgtk_only", &Gdk_window_thaw_toplevel_updates_libgtk_only},
	{"gdk_window_thaw_updates", &Gdk_window_thaw_updates},
	{"gdk_window_type_get_type", &Gdk_window_type_get_type},
	{"gdk_window_type_hint_get_type", &Gdk_window_type_hint_get_type},
	{"gdk_window_unfullscreen", &Gdk_window_unfullscreen},
	{"gdk_window_unmaximize", &Gdk_window_unmaximize},
	{"gdk_window_unstick", &Gdk_window_unstick},
	{"gdk_window_withdraw", &Gdk_window_withdraw},
	{"gdk_wm_decoration_get_type", &Gdk_wm_decoration_get_type},
	{"gdk_wm_function_get_type", &Gdk_wm_function_get_type},
}

var apiListPixbuf = outside.Apis{
	{"gdk_colorspace_get_type", &Gdk_colorspace_get_type},
	{"gdk_interp_type_get_type", &Gdk_interp_type_get_type},
	{"gdk_pixbuf_add_alpha", &Gdk_pixbuf_add_alpha},
	{"gdk_pixbuf_alpha_mode_get_type", &Gdk_pixbuf_alpha_mode_get_type},
	{"gdk_pixbuf_animation_get_height", &Gdk_pixbuf_animation_get_height},
	{"gdk_pixbuf_animation_get_iter", &Gdk_pixbuf_animation_get_iter},
	{"gdk_pixbuf_animation_get_static_image", &Gdk_pixbuf_animation_get_static_image},
	{"gdk_pixbuf_animation_get_type", &Gdk_pixbuf_animation_get_type},
	{"gdk_pixbuf_animation_get_width", &Gdk_pixbuf_animation_get_width},
	{"gdk_pixbuf_animation_is_static_image", &Gdk_pixbuf_animation_is_static_image},
	{"gdk_pixbuf_animation_iter_advance", &Gdk_pixbuf_animation_iter_advance},
	{"gdk_pixbuf_animation_iter_get_delay_time", &Gdk_pixbuf_animation_iter_get_delay_time},
	{"gdk_pixbuf_animation_iter_get_pixbuf", &Gdk_pixbuf_animation_iter_get_pixbuf},
	{"gdk_pixbuf_animation_iter_get_type", &Gdk_pixbuf_animation_iter_get_type},
	{"gdk_pixbuf_animation_iter_on_currently_loading_frame", &Gdk_pixbuf_animation_iter_on_currently_loading_frame},
	{"gdk_pixbuf_animation_new_from_file", &Gdk_pixbuf_animation_new_from_file},
	{"gdk_pixbuf_animation_new_from_file_utf8", &Gdk_pixbuf_animation_new_from_file_utf8},
	{"gdk_pixbuf_animation_ref", &Gdk_pixbuf_animation_ref},
	{"gdk_pixbuf_animation_unref", &Gdk_pixbuf_animation_unref},
	{"gdk_pixbuf_apply_embedded_orientation", &Gdk_pixbuf_apply_embedded_orientation},
	{"gdk_pixbuf_composite", &Gdk_pixbuf_composite},
	{"gdk_pixbuf_composite_color", &Gdk_pixbuf_composite_color},
	{"gdk_pixbuf_composite_color_simple", &Gdk_pixbuf_composite_color_simple},
	{"gdk_pixbuf_copy", &Gdk_pixbuf_copy},
	{"gdk_pixbuf_copy_area", &Gdk_pixbuf_copy_area},
	{"gdk_pixbuf_error_get_type", &Gdk_pixbuf_error_get_type},
	{"gdk_pixbuf_error_quark", &Gdk_pixbuf_error_quark},
	{"gdk_pixbuf_fill", &Gdk_pixbuf_fill},
	{"gdk_pixbuf_flip", &Gdk_pixbuf_flip},
	{"gdk_pixbuf_format_copy", &Gdk_pixbuf_format_copy},
	{"gdk_pixbuf_format_free", &Gdk_pixbuf_format_free},
	{"gdk_pixbuf_format_get_description", &Gdk_pixbuf_format_get_description},
	{"gdk_pixbuf_format_get_extensions", &Gdk_pixbuf_format_get_extensions},
	{"gdk_pixbuf_format_get_license", &Gdk_pixbuf_format_get_license},
	{"gdk_pixbuf_format_get_mime_types", &Gdk_pixbuf_format_get_mime_types},
	{"gdk_pixbuf_format_get_name", &Gdk_pixbuf_format_get_name},
	{"gdk_pixbuf_format_get_type", &Gdk_pixbuf_format_get_type},
	{"gdk_pixbuf_format_is_disabled", &Gdk_pixbuf_format_is_disabled},
	{"gdk_pixbuf_format_is_scalable", &Gdk_pixbuf_format_is_scalable},
	{"gdk_pixbuf_format_is_writable", &Gdk_pixbuf_format_is_writable},
	{"gdk_pixbuf_format_set_disabled", &Gdk_pixbuf_format_set_disabled},
	{"gdk_pixbuf_from_pixdata", &Gdk_pixbuf_from_pixdata},
	{"gdk_pixbuf_get_bits_per_sample", &Gdk_pixbuf_get_bits_per_sample},
	{"gdk_pixbuf_get_colorspace", &Gdk_pixbuf_get_colorspace},
	{"gdk_pixbuf_get_file_info", &Gdk_pixbuf_get_file_info},
	{"gdk_pixbuf_get_formats", &Gdk_pixbuf_get_formats},
	{"gdk_pixbuf_get_has_alpha", &Gdk_pixbuf_get_has_alpha},
	{"gdk_pixbuf_get_height", &Gdk_pixbuf_get_height},
	{"gdk_pixbuf_get_n_channels", &Gdk_pixbuf_get_n_channels},
	{"gdk_pixbuf_get_option", &Gdk_pixbuf_get_option},
	{"gdk_pixbuf_get_pixels", &Gdk_pixbuf_get_pixels},
	{"gdk_pixbuf_get_rowstride", &Gdk_pixbuf_get_rowstride},
	{"gdk_pixbuf_get_type", &Gdk_pixbuf_get_type},
	{"gdk_pixbuf_get_width", &Gdk_pixbuf_get_width},
	// Undocumented {"gdk_pixbuf_gettext", &Gdk_pixbuf_gettext},
	{"gdk_pixbuf_loader_close", &Gdk_pixbuf_loader_close},
	{"gdk_pixbuf_loader_get_animation", &Gdk_pixbuf_loader_get_animation},
	{"gdk_pixbuf_loader_get_format", &Gdk_pixbuf_loader_get_format},
	{"gdk_pixbuf_loader_get_pixbuf", &Gdk_pixbuf_loader_get_pixbuf},
	{"gdk_pixbuf_loader_get_type", &Gdk_pixbuf_loader_get_type},
	{"gdk_pixbuf_loader_new", &Gdk_pixbuf_loader_new},
	{"gdk_pixbuf_loader_new_with_mime_type", &Gdk_pixbuf_loader_new_with_mime_type},
	{"gdk_pixbuf_loader_new_with_type", &Gdk_pixbuf_loader_new_with_type},
	{"gdk_pixbuf_loader_set_size", &Gdk_pixbuf_loader_set_size},
	{"gdk_pixbuf_loader_write", &Gdk_pixbuf_loader_write},
	{"gdk_pixbuf_new", &Gdk_pixbuf_new},
	{"gdk_pixbuf_new_from_data", &Gdk_pixbuf_new_from_data},
	// Win32 uses _utf8 {"gdk_pixbuf_new_from_file", &Gdk_pixbuf_new_from_file},
	// Win32 uese _utf8 {"gdk_pixbuf_new_from_file_at_scale", &Gdk_pixbuf_new_from_file_at_scale},
	{"gdk_pixbuf_new_from_file_at_scale_utf8", &Gdk_pixbuf_new_from_file_at_scale},
	// Win32 uses _utf8 {"gdk_pixbuf_new_from_file_at_size", &Gdk_pixbuf_new_from_file_at_size},
	{"gdk_pixbuf_new_from_file_at_size_utf8", &Gdk_pixbuf_new_from_file_at_size},
	{"gdk_pixbuf_new_from_file_utf8", &Gdk_pixbuf_new_from_file},
	{"gdk_pixbuf_new_from_inline", &Gdk_pixbuf_new_from_inline},
	{"gdk_pixbuf_new_from_stream", &Gdk_pixbuf_new_from_stream},
	{"gdk_pixbuf_new_from_stream_async", &Gdk_pixbuf_new_from_stream_async},
	{"gdk_pixbuf_new_from_stream_at_scale", &Gdk_pixbuf_new_from_stream_at_scale},
	{"gdk_pixbuf_new_from_stream_at_scale_async", &Gdk_pixbuf_new_from_stream_at_scale_async},
	{"gdk_pixbuf_new_from_stream_finish", &Gdk_pixbuf_new_from_stream_finish},
	{"gdk_pixbuf_new_from_xpm_data", &Gdk_pixbuf_new_from_xpm_data},
	{"gdk_pixbuf_new_subpixbuf", &Gdk_pixbuf_new_subpixbuf},
	{"gdk_pixbuf_non_anim_get_type", &Gdk_pixbuf_non_anim_get_type},
	{"gdk_pixbuf_non_anim_new", &Gdk_pixbuf_non_anim_new},
	{"gdk_pixbuf_ref", &Gdk_pixbuf_ref},
	{"gdk_pixbuf_rotate_simple", &Gdk_pixbuf_rotate_simple},
	{"gdk_pixbuf_rotation_get_type", &Gdk_pixbuf_rotation_get_type},
	{"gdk_pixbuf_saturate_and_pixelate", &Gdk_pixbuf_saturate_and_pixelate},
	// Win32 uses _utf8 {"gdk_pixbuf_save", &Gdk_pixbuf_save},
	{"gdk_pixbuf_save_to_buffer", &Gdk_pixbuf_save_to_buffer},
	{"gdk_pixbuf_save_to_bufferv", &Gdk_pixbuf_save_to_bufferv},
	{"gdk_pixbuf_save_to_callback", &Gdk_pixbuf_save_to_callback},
	{"gdk_pixbuf_save_to_callbackv", &Gdk_pixbuf_save_to_callbackv},
	{"gdk_pixbuf_save_to_stream", &Gdk_pixbuf_save_to_stream},
	{"gdk_pixbuf_save_to_stream_async", &Gdk_pixbuf_save_to_stream_async},
	{"gdk_pixbuf_save_to_stream_finish", &Gdk_pixbuf_save_to_stream_finish},
	{"gdk_pixbuf_save_utf8", &Gdk_pixbuf_save_utf8},
	// Win32 uses _utf8 {"gdk_pixbuf_savev", &Gdk_pixbuf_savev},
	{"gdk_pixbuf_savev_utf8", &Gdk_pixbuf_savev_utf8},
	{"gdk_pixbuf_scale", &Gdk_pixbuf_scale},
	{"gdk_pixbuf_scale_simple", &Gdk_pixbuf_scale_simple},
	// Undocumented {"gdk_pixbuf_scaled_anim_get_type", &Gdk_pixbuf_scaled_anim_get_type},
	// Undocumented {"gdk_pixbuf_scaled_anim_iter_get_type", &Gdk_pixbuf_scaled_anim_iter_get_type},
	{"gdk_pixbuf_set_option", &Gdk_pixbuf_set_option},
	{"gdk_pixbuf_simple_anim_add_frame", &Gdk_pixbuf_simple_anim_add_frame},
	{"gdk_pixbuf_simple_anim_get_loop", &Gdk_pixbuf_simple_anim_get_loop},
	{"gdk_pixbuf_simple_anim_get_type", &Gdk_pixbuf_simple_anim_get_type},
	{"gdk_pixbuf_simple_anim_iter_get_type", &Gdk_pixbuf_simple_anim_iter_get_type},
	{"gdk_pixbuf_simple_anim_new", &Gdk_pixbuf_simple_anim_new},
	{"gdk_pixbuf_simple_anim_set_loop", &Gdk_pixbuf_simple_anim_set_loop},
	{"gdk_pixbuf_unref", &Gdk_pixbuf_unref},
	{"gdk_pixdata_deserialize", &Gdk_pixdata_deserialize},
	{"gdk_pixdata_from_pixbuf", &Gdk_pixdata_from_pixbuf},
	{"gdk_pixdata_serialize", &Gdk_pixdata_serialize},
	{"gdk_pixdata_to_csource", &Gdk_pixdata_to_csource},
}

var dataList = outside.Data{
	{"gdk_threads_lock", (*T.GCallback)(nil)},
	{"gdk_threads_mutex", (*T.GMutex)(nil)},
	{"gdk_threads_unlock", (*T.GCallback)(nil)},
}

var dataListPixbuf = outside.Data{
	{"gdk_pixbuf_major_version", (*T.Guint)(nil)},
	{"gdk_pixbuf_micro_version", (*T.Guint)(nil)},
	{"gdk_pixbuf_minor_version", (*T.Guint)(nil)},
	{"gdk_pixbuf_version", (*uint8)(nil)},
}
