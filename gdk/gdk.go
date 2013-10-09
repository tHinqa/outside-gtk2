// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package gdk provides API definitions for accessing
//libgdk-win32-2.0-0.dll and libgdk_pixbuf-2.0-0.dll.
package gdk

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
	AddDllApis(dllPixbuf, false, apiListPixbuf)
}

type (
	HICON   uint32
	HGDIOBJ uint32
	HWND    uint32
	HDC     uint32
)

var (
	Gdk_colormap_get_type func() GType

	Gdk_colormap_new func(
		visual *GdkVisual,
		allocate Gboolean) *GdkColormap

	Gdk_colormap_ref func(
		cmap *GdkColormap) *GdkColormap

	Gdk_colormap_unref func(
		cmap *GdkColormap)

	Gdk_colormap_get_system func() *GdkColormap

	Gdk_colormap_get_screen func(
		cmap *GdkColormap) *GdkScreen

	Gdk_colormap_get_system_size func() Gint

	Gdk_colormap_change func(
		colormap *GdkColormap,
		ncolors Gint)

	Gdk_colormap_alloc_colors func(
		colormap *GdkColormap,
		colors *GdkColor,
		n_colors Gint,
		writeable Gboolean,
		best_match Gboolean,
		success *Gboolean) Gint

	Gdk_colormap_alloc_color func(
		colormap *GdkColormap,
		color *GdkColor,
		writeable Gboolean,
		best_match Gboolean) Gboolean

	Gdk_colormap_free_colors func(
		colormap *GdkColormap,
		colors *GdkColor,
		n_colors Gint)

	Gdk_colormap_query_color func(
		colormap *GdkColormap,
		pixel Gulong,
		result *GdkColor)

	Gdk_colormap_get_visual func(
		colormap *GdkColormap) *GdkVisual

	Gdk_color_copy func(
		color *GdkColor) *GdkColor

	Gdk_color_free func(
		color *GdkColor)

	Gdk_color_parse func(
		spec string,
		color *GdkColor) Gboolean

	Gdk_color_hash func(
		colora *GdkColor) Guint

	Gdk_color_equal func(
		colora *GdkColor,
		colorb *GdkColor) Gboolean

	Gdk_color_to_string func(
		color *GdkColor) string

	Gdk_color_get_type func() GType

	Gdk_colors_store func(
		colormap *GdkColormap,
		colors *GdkColor,
		ncolors Gint)

	Gdk_color_white func(
		colormap *GdkColormap,
		color *GdkColor) Gint

	Gdk_color_black func(
		colormap *GdkColormap,
		color *GdkColor) Gint

	Gdk_color_alloc func(
		colormap *GdkColormap,
		color *GdkColor) Gint

	Gdk_color_change func(
		colormap *GdkColormap,
		color *GdkColor) Gint

	Gdk_colors_alloc func(
		colormap *GdkColormap,
		contiguous Gboolean,
		planes *Gulong,
		nplanes Gint,
		pixels *Gulong,
		npixels Gint) Gint

	Gdk_colors_free func(
		colormap *GdkColormap,
		pixels *Gulong,
		npixels Gint,
		planes Gulong)

	Gdk_drag_context_get_type func() GType

	Gdk_drag_context_new func() *GdkDragContext

	Gdk_drag_context_list_targets func(
		context *GdkDragContext) *GList

	Gdk_drag_context_get_actions func(
		context *GdkDragContext) GdkDragAction

	Gdk_drag_context_get_suggested_action func(
		context *GdkDragContext) GdkDragAction

	Gdk_drag_context_get_selected_action func(
		context *GdkDragContext) GdkDragAction

	Gdk_drag_context_get_source_window func(
		context *GdkDragContext) *GdkWindow

	Gdk_drag_context_get_dest_window func(
		context *GdkDragContext) *GdkWindow

	Gdk_drag_context_get_protocol func(
		context *GdkDragContext) GdkDragProtocol

	Gdk_drag_context_ref func(
		context *GdkDragContext)

	Gdk_drag_context_unref func(
		context *GdkDragContext)

	Gdk_drag_status func(
		context *GdkDragContext,
		action GdkDragAction,
		time_ Guint32)

	Gdk_drop_reply func(
		context *GdkDragContext,
		ok Gboolean,
		time_ Guint32)

	Gdk_drop_finish func(
		context *GdkDragContext,
		success Gboolean,
		time_ Guint32)

	Gdk_drag_get_selection func(
		context *GdkDragContext) GdkAtom

	Gdk_drag_begin func(
		window *GdkWindow,
		targets *GList) *GdkDragContext

	Gdk_drag_get_protocol_for_display func(
		display *GdkDisplay,
		xid GdkNativeWindow,
		protocol *GdkDragProtocol) GdkNativeWindow

	Gdk_drag_find_window_for_screen func(
		context *GdkDragContext,
		drag_window *GdkWindow,
		screen *GdkScreen,
		x_root Gint,
		y_root Gint,
		dest_window **GdkWindow,
		protocol *GdkDragProtocol)

	Gdk_drag_get_protocol func(
		xid GdkNativeWindow,
		protocol *GdkDragProtocol) GdkNativeWindow

	Gdk_drag_find_window func(
		context *GdkDragContext,
		drag_window *GdkWindow,
		x_root Gint,
		y_root Gint,
		dest_window **GdkWindow,
		protocol *GdkDragProtocol)

	Gdk_drag_motion func(
		context *GdkDragContext,
		dest_window *GdkWindow,
		protocol GdkDragProtocol,
		x_root Gint,
		y_root Gint,
		suggested_action GdkDragAction,
		possible_actions GdkDragAction,
		time_ Guint32) Gboolean

	Gdk_drag_drop func(
		context *GdkDragContext,
		time_ Guint32)

	Gdk_drag_abort func(
		context *GdkDragContext,
		time_ Guint32)

	Gdk_drag_drop_succeeded func(
		context *GdkDragContext) Gboolean

	Gdk_device_get_type func() GType

	Gdk_devices_list func() *GList

	Gdk_device_get_name func(
		device *GdkDevice) string

	Gdk_device_get_source func(
		device *GdkDevice) GdkInputSource

	Gdk_device_get_mode func(
		device *GdkDevice) GdkInputMode

	Gdk_device_get_has_cursor func(
		device *GdkDevice) Gboolean

	Gdk_device_get_key func(
		device *GdkDevice,
		index Guint,
		keyval *Guint,
		modifiers *GdkModifierType)

	Gdk_device_get_axis_use func(
		device *GdkDevice,
		index Guint) GdkAxisUse

	Gdk_device_get_n_keys func(
		device *GdkDevice) Gint

	Gdk_device_get_n_axes func(
		device *GdkDevice) Gint

	Gdk_device_set_source func(
		device *GdkDevice,
		source GdkInputSource)

	Gdk_device_set_mode func(
		device *GdkDevice,
		mode GdkInputMode) Gboolean

	Gdk_device_set_key func(
		device *GdkDevice,
		index_ Guint,
		keyval Guint,
		modifiers GdkModifierType)

	Gdk_device_set_axis_use func(
		device *GdkDevice,
		index_ Guint,
		use GdkAxisUse)

	Gdk_device_get_state func(
		device *GdkDevice,
		window *GdkWindow,
		axes *Gdouble,
		mask *GdkModifierType)

	Gdk_device_get_history func(
		device *GdkDevice,
		window *GdkWindow,
		start Guint32,
		stop Guint32,
		events ***GdkTimeCoord,
		n_events *Gint) Gboolean

	Gdk_device_free_history func(
		events **GdkTimeCoord,
		n_events Gint)

	Gdk_device_get_axis func(
		device *GdkDevice,
		axes *Gdouble,
		use GdkAxisUse,
		value *Gdouble) Gboolean

	Gdk_input_set_extension_events func(
		window *GdkWindow,
		mask Gint,
		mode GdkExtensionMode)

	Gdk_device_get_core_pointer func() *GdkDevice

	Gdk_event_get_type func() GType

	Gdk_events_pending func() Gboolean

	Gdk_event_get func() *GdkEvent

	Gdk_event_peek func() *GdkEvent

	Gdk_event_get_graphics_expose func(
		window *GdkWindow) *GdkEvent

	Gdk_event_put func(
		event *GdkEvent)

	Gdk_event_new func(
		t GdkEventType) *GdkEvent

	Gdk_event_copy func(
		event *GdkEvent) *GdkEvent

	Gdk_event_free func(
		event *GdkEvent)

	Gdk_event_get_time func(
		event *GdkEvent) Guint32

	Gdk_event_get_state func(
		event *GdkEvent,
		state *GdkModifierType) Gboolean

	Gdk_event_get_coords func(
		event *GdkEvent,
		x_win *Gdouble,
		y_win *Gdouble) Gboolean

	Gdk_event_get_root_coords func(
		event *GdkEvent,
		x_root *Gdouble,
		y_root *Gdouble) Gboolean

	Gdk_event_get_axis func(
		event *GdkEvent,
		axis_use GdkAxisUse,
		value *Gdouble) Gboolean

	Gdk_event_request_motions func(
		event *GdkEventMotion)

	Gdk_event_handler_set func(
		f GdkEventFunc,
		data Gpointer,
		notify GDestroyNotify)

	Gdk_event_set_screen func(
		event *GdkEvent,
		screen *GdkScreen)

	Gdk_event_get_screen func(
		event *GdkEvent) *GdkScreen

	Gdk_set_show_events func(
		show_events Gboolean)

	Gdk_get_show_events func() Gboolean

	Gdk_add_client_message_filter func(
		message_type GdkAtom,
		f GdkFilterFunc,
		data Gpointer)

	Gdk_setting_get func(
		name string,
		value *GValue) Gboolean

	Gdk_display_get_type func() GType

	Gdk_display_open func(
		display_name string) *GdkDisplay

	Gdk_display_get_name func(
		display *GdkDisplay) string

	Gdk_display_get_n_screens func(
		display *GdkDisplay) Gint

	Gdk_display_get_screen func(
		display *GdkDisplay,
		screen_num Gint) *GdkScreen

	Gdk_display_get_default_screen func(
		display *GdkDisplay) *GdkScreen

	Gdk_display_pointer_ungrab func(
		display *GdkDisplay,
		time_ Guint32)

	Gdk_display_keyboard_ungrab func(
		display *GdkDisplay,
		time_ Guint32)

	Gdk_display_pointer_is_grabbed func(
		display *GdkDisplay) Gboolean

	Gdk_display_beep func(
		display *GdkDisplay)

	Gdk_display_sync func(
		display *GdkDisplay)

	Gdk_display_flush func(
		display *GdkDisplay)

	Gdk_display_close func(
		display *GdkDisplay)

	Gdk_display_is_closed func(
		display *GdkDisplay) Gboolean

	Gdk_display_list_devices func(
		display *GdkDisplay) *GList

	Gdk_display_get_event func(
		display *GdkDisplay) *GdkEvent

	Gdk_display_peek_event func(
		display *GdkDisplay) *GdkEvent

	Gdk_display_put_event func(
		display *GdkDisplay,
		event *GdkEvent)

	Gdk_display_add_client_message_filter func(
		display *GdkDisplay,
		message_type GdkAtom,
		f GdkFilterFunc,
		data Gpointer)

	Gdk_display_set_double_click_time func(
		display *GdkDisplay,
		msec Guint)

	Gdk_display_set_double_click_distance func(
		display *GdkDisplay,
		distance Guint)

	Gdk_display_get_default func() *GdkDisplay

	Gdk_display_get_core_pointer func(
		display *GdkDisplay) *GdkDevice

	Gdk_display_get_pointer func(
		display *GdkDisplay,
		screen **GdkScreen,
		x *Gint,
		y *Gint,
		mask *GdkModifierType)

	Gdk_display_get_window_at_pointer func(
		display *GdkDisplay,
		win_x *Gint,
		win_y *Gint) *GdkWindow

	Gdk_display_warp_pointer func(
		display *GdkDisplay,
		screen *GdkScreen,
		x Gint,
		y Gint)

	Gdk_display_set_pointer_hooks func(
		display *GdkDisplay,
		new_hooks *GdkDisplayPointerHooks) *GdkDisplayPointerHooks

	Gdk_display_open_default_libgtk_only func() *GdkDisplay

	Gdk_display_supports_cursor_alpha func(
		display *GdkDisplay) Gboolean

	Gdk_display_supports_cursor_color func(
		display *GdkDisplay) Gboolean

	Gdk_display_get_default_cursor_size func(
		display *GdkDisplay) Guint

	Gdk_display_get_maximal_cursor_size func(
		display *GdkDisplay,
		width *Guint,
		height *Guint)

	Gdk_display_get_default_group func(
		display *GdkDisplay) *GdkWindow

	Gdk_display_supports_selection_notification func(
		display *GdkDisplay) Gboolean

	Gdk_display_request_selection_notification func(
		display *GdkDisplay,
		selection GdkAtom) Gboolean

	Gdk_display_supports_clipboard_persistence func(
		display *GdkDisplay) Gboolean

	Gdk_display_store_clipboard func(
		display *GdkDisplay,
		clipboard_window *GdkWindow,
		time_ Guint32,
		targets *GdkAtom,
		n_targets Gint)

	Gdk_display_supports_shapes func(
		display *GdkDisplay) Gboolean

	Gdk_display_supports_input_shapes func(
		display *GdkDisplay) Gboolean

	Gdk_display_supports_composite func(
		display *GdkDisplay) Gboolean

	Gdk_screen_get_type func() GType

	Gdk_screen_get_default_colormap func(
		screen *GdkScreen) *GdkColormap

	Gdk_screen_set_default_colormap func(
		screen *GdkScreen,
		colormap *GdkColormap)

	Gdk_screen_get_system_colormap func(
		screen *GdkScreen) *GdkColormap

	Gdk_screen_get_system_visual func(
		screen *GdkScreen) *GdkVisual

	Gdk_screen_get_rgb_colormap func(
		screen *GdkScreen) *GdkColormap

	Gdk_screen_get_rgb_visual func(
		screen *GdkScreen) *GdkVisual

	Gdk_screen_get_rgba_colormap func(
		screen *GdkScreen) *GdkColormap

	Gdk_screen_get_rgba_visual func(
		screen *GdkScreen) *GdkVisual

	Gdk_screen_is_composited func(
		screen *GdkScreen) Gboolean

	Gdk_screen_get_root_window func(
		screen *GdkScreen) *GdkWindow

	Gdk_screen_get_display func(
		screen *GdkScreen) *GdkDisplay

	Gdk_screen_get_number func(
		screen *GdkScreen) Gint

	Gdk_screen_get_width func(
		screen *GdkScreen) Gint

	Gdk_screen_get_height func(
		screen *GdkScreen) Gint

	Gdk_screen_get_width_mm func(
		screen *GdkScreen) Gint

	Gdk_screen_get_height_mm func(
		screen *GdkScreen) Gint

	Gdk_screen_list_visuals func(
		screen *GdkScreen) *GList

	Gdk_screen_get_toplevel_windows func(
		screen *GdkScreen) *GList

	Gdk_screen_make_display_name func(
		screen *GdkScreen) string

	Gdk_screen_get_n_monitors func(
		screen *GdkScreen) Gint

	Gdk_screen_get_primary_monitor func(
		screen *GdkScreen) Gint

	Gdk_screen_get_monitor_geometry func(
		screen *GdkScreen,
		monitor_num Gint,
		dest *GdkRectangle)

	Gdk_screen_get_monitor_at_point func(
		screen *GdkScreen,
		x Gint,
		y Gint) Gint

	Gdk_screen_get_monitor_at_window func(
		screen *GdkScreen,
		window *GdkWindow) Gint

	Gdk_screen_get_monitor_width_mm func(
		screen *GdkScreen,
		monitor_num Gint) Gint

	Gdk_screen_get_monitor_height_mm func(
		screen *GdkScreen,
		monitor_num Gint) Gint

	Gdk_screen_get_monitor_plug_name func(
		screen *GdkScreen,
		monitor_num Gint) string

	Gdk_screen_broadcast_client_message func(
		screen *GdkScreen,
		event *GdkEvent)

	Gdk_screen_get_default func() *GdkScreen

	Gdk_screen_get_setting func(
		screen *GdkScreen,
		name string,
		value *GValue) Gboolean

	Gdk_screen_set_font_options func(
		screen *GdkScreen,
		options *Cairo_font_options_t)

	Gdk_screen_get_font_options func(
		screen *GdkScreen) *Cairo_font_options_t

	Gdk_screen_set_resolution func(
		screen *GdkScreen,
		dpi Gdouble)

	Gdk_screen_get_resolution func(
		screen *GdkScreen) Gdouble

	Gdk_screen_get_active_window func(
		screen *GdkScreen) *GdkWindow

	Gdk_screen_get_window_stack func(
		screen *GdkScreen) *GList

	Gdk_app_launch_context_get_type func() GType

	Gdk_app_launch_context_new func() *GdkAppLaunchContext

	Gdk_app_launch_context_set_display func(
		context *GdkAppLaunchContext,
		display *GdkDisplay)

	Gdk_app_launch_context_set_screen func(
		context *GdkAppLaunchContext,
		screen *GdkScreen)

	Gdk_app_launch_context_set_desktop func(
		context *GdkAppLaunchContext,
		desktop Gint)

	Gdk_app_launch_context_set_timestamp func(
		context *GdkAppLaunchContext,
		timestamp Guint32)

	Gdk_app_launch_context_set_icon func(
		context *GdkAppLaunchContext,
		icon *GIcon)

	Gdk_app_launch_context_set_icon_name func(
		context *GdkAppLaunchContext,
		icon_name string)

	Gdk_rgb_init func()

	Gdk_rgb_xpixel_from_rgb func(
		rgb Guint32) Gulong

	Gdk_rgb_gc_set_foreground func(
		gc *GdkGC,
		rgb Guint32)

	Gdk_rgb_gc_set_background func(
		gc *GdkGC,
		rgb Guint32)

	Gdk_rgb_find_color func(
		colormap *GdkColormap,
		color *GdkColor)

	Gdk_draw_rgb_image func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		dith GdkRgbDither,
		rgb_buf *Guchar,
		rowstride Gint)

	Gdk_draw_rgb_image_dithalign func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		dith GdkRgbDither,
		rgb_buf *Guchar,
		rowstride Gint,
		xdith Gint,
		ydith Gint)

	Gdk_draw_rgb_32_image func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		dith GdkRgbDither,
		buf *Guchar,
		rowstride Gint)

	Gdk_draw_rgb_32_image_dithalign func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		dith GdkRgbDither,
		buf *Guchar,
		rowstride Gint,
		xdith Gint,
		ydith Gint)

	Gdk_draw_gray_image func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		dith GdkRgbDither,
		buf *Guchar,
		rowstride Gint)

	Gdk_draw_indexed_image func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		dith GdkRgbDither,
		buf *Guchar,
		rowstride Gint,
		cmap *GdkRgbCmap)

	Gdk_rgb_cmap_new func(
		colors *Guint32,
		n_colors Gint) *GdkRgbCmap

	Gdk_rgb_cmap_free func(
		cmap *GdkRgbCmap)

	Gdk_rgb_set_verbose func(
		verbose Gboolean)

	Gdk_rgb_set_install func(
		install Gboolean)

	Gdk_rgb_set_min_colors func(
		min_colors Gint)

	Gdk_rgb_get_colormap func() *GdkColormap

	Gdk_rgb_get_visual func() *GdkVisual

	Gdk_rgb_ditherable func() Gboolean

	Gdk_rgb_colormap_ditherable func(
		cmap *GdkColormap) Gboolean

	Gdk_pixbuf_error_quark func() GQuark

	Gdk_pixbuf_get_type func() GType

	Gdk_pixbuf_ref func(
		pixbuf *GdkPixbuf) *GdkPixbuf

	Gdk_pixbuf_unref func(
		pixbuf *GdkPixbuf)

	Gdk_pixbuf_get_colorspace func(
		pixbuf *GdkPixbuf) GdkColorspace

	Gdk_pixbuf_get_n_channels func(
		pixbuf *GdkPixbuf) int

	Gdk_pixbuf_get_has_alpha func(
		pixbuf *GdkPixbuf) Gboolean

	Gdk_pixbuf_get_bits_per_sample func(
		pixbuf *GdkPixbuf) int

	Gdk_pixbuf_get_pixels func(
		pixbuf *GdkPixbuf) *Guchar

	Gdk_pixbuf_get_width func(
		pixbuf *GdkPixbuf) int

	Gdk_pixbuf_get_height func(
		pixbuf *GdkPixbuf) int

	Gdk_pixbuf_get_rowstride func(
		pixbuf *GdkPixbuf) int

	Gdk_pixbuf_new func(
		colorspace GdkColorspace,
		has_alpha Gboolean,
		bits_per_sample int,
		width int,
		height int) *GdkPixbuf

	Gdk_pixbuf_copy func(
		pixbuf *GdkPixbuf) *GdkPixbuf

	Gdk_pixbuf_new_subpixbuf func(
		src_pixbuf *GdkPixbuf,
		src_x int,
		src_y int,
		width int,
		height int) *GdkPixbuf

	Gdk_pixbuf_new_from_file func(
		filename string,
		e **GError) *GdkPixbuf

	Gdk_pixbuf_new_from_file_at_size func(
		filename string,
		width int,
		height int,
		e **GError) *GdkPixbuf

	Gdk_pixbuf_new_from_file_at_scale func(
		filename string,
		width int,
		height int,
		preserve_aspect_ratio Gboolean,
		e **GError) *GdkPixbuf

	Gdk_pixbuf_new_from_data func(
		data *Guchar,
		colorspace GdkColorspace,
		has_alpha Gboolean,
		bits_per_sample int,
		width int,
		height int,
		rowstride int,
		destroy_fn GdkPixbufDestroyNotify,
		destroy_fn_data Gpointer) *GdkPixbuf

	Gdk_pixbuf_new_from_xpm_data func(
		data **Char) *GdkPixbuf

	Gdk_pixbuf_new_from_inline func(
		data_length Gint,
		data *Guint8,
		copy_pixels Gboolean,
		e **GError) *GdkPixbuf

	Gdk_pixbuf_fill func(
		pixbuf *GdkPixbuf,
		pixel Guint32)

	Gdk_pixbuf_save_utf8 func(pixbuf *GdkPixbuf,
		filename, typ string, e **GError, v ...VArg) Gboolean

	Gdk_pixbuf_savev_utf8 func(
		pixbuf *GdkPixbuf,
		filename string,
		typ string,
		option_keys **Char,
		option_values **Char,
		e **GError) Gboolean

	Gdk_pixbuf_save_to_callback func(pixbuf *GdkPixbuf,
		save_func GdkPixbufSaveFunc,
		user_data Gpointer,
		typ string, err **GError, v ...VArg) Gboolean

	Gdk_pixbuf_save_to_callbackv func(
		pixbuf *GdkPixbuf,
		save_func GdkPixbufSaveFunc,
		user_data Gpointer,
		typ string,
		option_keys **Char,
		option_values **Char,
		e **GError) Gboolean

	Gdk_pixbuf_save_to_buffer func(pixbuf *GdkPixbuf,
		buffer **Gchar, buffer_size *Gsize,
		typ string, e **GError, v ...VArg) Gboolean

	Gdk_pixbuf_save_to_bufferv func(
		pixbuf *GdkPixbuf,
		buffer **Gchar,
		buffer_size *Gsize,
		typ string,
		option_keys **Char,
		option_values **Char,
		e **GError) Gboolean

	Gdk_pixbuf_new_from_stream func(
		stream *GInputStream,
		cancellable *GCancellable,
		e **GError) *GdkPixbuf

	Gdk_pixbuf_new_from_stream_async func(
		stream *GInputStream,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	Gdk_pixbuf_new_from_stream_finish func(
		async_result *GAsyncResult,
		e **GError) *GdkPixbuf

	Gdk_pixbuf_new_from_stream_at_scale func(
		stream *GInputStream,
		width Gint,
		height Gint,
		preserve_aspect_ratio Gboolean,
		cancellable *GCancellable,
		e **GError) *GdkPixbuf

	Gdk_pixbuf_new_from_stream_at_scale_async func(
		stream *GInputStream,
		width Gint,
		height Gint,
		preserve_aspect_ratio Gboolean,
		cancellable *GCancellable,
		callback GAsyncReadyCallback,
		user_data Gpointer)

	Gdk_pixbuf_save_to_stream func(pixbuf *GdkPixbuf,
		stream *GOutputStream, typ string,
		cancellable *GCancellable, e **GError, v ...VArg) Gboolean

	Gdk_pixbuf_save_to_stream_async func(pixbuf *GdkPixbuf,
		stream *GOutputStream, typ string,
		cancellable *GCancellable, callback GAsyncReadyCallback,
		user_data Gpointer, v ...VArg)

	Gdk_pixbuf_save_to_stream_finish func(
		async_result *GAsyncResult,
		e **GError) Gboolean

	Gdk_pixbuf_add_alpha func(
		pixbuf *GdkPixbuf,
		substitute_color Gboolean,
		r, g, b Guchar) *GdkPixbuf

	Gdk_pixbuf_copy_area func(
		src_pixbuf *GdkPixbuf,
		src_x, src_y, width, height int,
		dest_pixbuf *GdkPixbuf,
		dest_x, dest_y int)

	Gdk_pixbuf_saturate_and_pixelate func(
		src, dest *GdkPixbuf,
		saturation Gfloat,
		pixelate Gboolean)

	Gdk_pixbuf_apply_embedded_orientation func(
		src *GdkPixbuf) *GdkPixbuf

	Gdk_pixbuf_get_option func(
		pixbuf *GdkPixbuf,
		key string) string

	Gdk_pixbuf_scale func(
		src *GdkPixbuf,
		dest *GdkPixbuf,
		dest_x int,
		dest_y int,
		dest_width int,
		dest_height int,
		offset_x Double,
		offset_y Double,
		scale_x Double,
		scale_y Double,
		interp_type GdkInterpType)

	Gdk_pixbuf_composite func(
		src *GdkPixbuf,
		dest *GdkPixbuf,
		dest_x int,
		dest_y int,
		dest_width int,
		dest_height int,
		offset_x Double,
		offset_y Double,
		scale_x Double,
		scale_y Double,
		interp_type GdkInterpType,
		overall_alpha int)

	Gdk_pixbuf_composite_color func(
		src *GdkPixbuf,
		dest *GdkPixbuf,
		dest_x int,
		dest_y int,
		dest_width int,
		dest_height int,
		offset_x Double,
		offset_y Double,
		scale_x Double,
		scale_y Double,
		interp_type GdkInterpType,
		overall_alpha int,
		check_x int,
		check_y int,
		check_size int,
		color1 Guint32,
		color2 Guint32)

	Gdk_pixbuf_scale_simple func(
		src *GdkPixbuf,
		dest_width int,
		dest_height int,
		interp_type GdkInterpType) *GdkPixbuf

	Gdk_pixbuf_composite_color_simple func(
		src *GdkPixbuf,
		dest_width int,
		dest_height int,
		interp_type GdkInterpType,
		overall_alpha int,
		check_size int,
		color1 Guint32,
		color2 Guint32) *GdkPixbuf

	Gdk_pixbuf_rotate_simple func(
		src *GdkPixbuf,
		angle GdkPixbufRotation) *GdkPixbuf

	Gdk_pixbuf_flip func(
		src *GdkPixbuf,
		horizontal Gboolean) *GdkPixbuf

	Gdk_pixbuf_animation_get_type func() GType

	Gdk_pixbuf_animation_new_from_file_utf8 func(
		filename string,
		e **GError) *GdkPixbufAnimation

	Gdk_pixbuf_animation_ref func(
		animation *GdkPixbufAnimation) *GdkPixbufAnimation

	Gdk_pixbuf_animation_unref func(
		animation *GdkPixbufAnimation)

	Gdk_pixbuf_animation_get_width func(
		animation *GdkPixbufAnimation) int

	Gdk_pixbuf_animation_get_height func(
		animation *GdkPixbufAnimation) int

	Gdk_pixbuf_animation_is_static_image func(
		animation *GdkPixbufAnimation) Gboolean

	Gdk_pixbuf_animation_get_static_image func(
		animation *GdkPixbufAnimation) *GdkPixbuf

	Gdk_pixbuf_animation_get_iter func(
		animation *GdkPixbufAnimation,
		start_time *GTimeVal) *GdkPixbufAnimationIter

	Gdk_pixbuf_animation_iter_get_type func() GType

	Gdk_pixbuf_animation_iter_get_delay_time func(
		iter *GdkPixbufAnimationIter) int

	Gdk_pixbuf_animation_iter_get_pixbuf func(
		iter *GdkPixbufAnimationIter) *GdkPixbuf

	Gdk_pixbuf_animation_iter_on_currently_loading_frame func(
		iter *GdkPixbufAnimationIter) Gboolean

	Gdk_pixbuf_animation_iter_advance func(
		iter *GdkPixbufAnimationIter,
		current_time *GTimeVal) Gboolean

	Gdk_pixbuf_simple_anim_get_type func() GType

	Gdk_pixbuf_simple_anim_iter_get_type func() GType

	Gdk_pixbuf_simple_anim_new func(
		width Gint,
		height Gint,
		rate Gfloat) *GdkPixbufSimpleAnim

	Gdk_pixbuf_simple_anim_add_frame func(
		animation *GdkPixbufSimpleAnim,
		pixbuf *GdkPixbuf)

	Gdk_pixbuf_simple_anim_set_loop func(
		animation *GdkPixbufSimpleAnim,
		loop Gboolean)

	Gdk_pixbuf_simple_anim_get_loop func(
		animation *GdkPixbufSimpleAnim) Gboolean

	Gdk_pixbuf_format_get_type func() GType

	Gdk_pixbuf_get_formats func() *GSList

	Gdk_pixbuf_format_get_name func(
		format *GdkPixbufFormat) string

	Gdk_pixbuf_format_get_description func(
		format *GdkPixbufFormat) string

	Gdk_pixbuf_format_get_mime_types func(
		format *GdkPixbufFormat) **Gchar

	Gdk_pixbuf_format_get_extensions func(
		format *GdkPixbufFormat) **Gchar

	Gdk_pixbuf_format_is_writable func(
		format *GdkPixbufFormat) Gboolean

	Gdk_pixbuf_format_is_scalable func(
		format *GdkPixbufFormat) Gboolean

	Gdk_pixbuf_format_is_disabled func(
		format *GdkPixbufFormat) Gboolean

	Gdk_pixbuf_format_set_disabled func(
		format *GdkPixbufFormat,
		disabled Gboolean)

	Gdk_pixbuf_format_get_license func(
		format *GdkPixbufFormat) string

	Gdk_pixbuf_get_file_info func(
		filename string,
		width *Gint,
		height *Gint) *GdkPixbufFormat

	Gdk_pixbuf_format_copy func(
		format *GdkPixbufFormat) *GdkPixbufFormat

	Gdk_pixbuf_format_free func(
		format *GdkPixbufFormat)

	Gdk_pixbuf_loader_get_type func() GType

	Gdk_pixbuf_loader_new func() *GdkPixbufLoader

	Gdk_pixbuf_loader_new_with_type func(
		image_type string,
		e **GError) *GdkPixbufLoader

	Gdk_pixbuf_loader_new_with_mime_type func(
		mime_type string,
		e **GError) *GdkPixbufLoader

	Gdk_pixbuf_loader_set_size func(
		loader *GdkPixbufLoader,
		width int,
		height int)

	Gdk_pixbuf_loader_write func(
		loader *GdkPixbufLoader,
		buf *Guchar,
		count Gsize,
		e **GError) Gboolean

	Gdk_pixbuf_loader_get_pixbuf func(
		loader *GdkPixbufLoader) *GdkPixbuf

	Gdk_pixbuf_loader_get_animation func(
		loader *GdkPixbufLoader) *GdkPixbufAnimation

	Gdk_pixbuf_loader_close func(
		loader *GdkPixbufLoader,
		e **GError) Gboolean

	Gdk_pixbuf_loader_get_format func(
		loader *GdkPixbufLoader) *GdkPixbufFormat

	Gdk_pixbuf_alpha_mode_get_type func() GType

	Gdk_colorspace_get_type func() GType

	Gdk_pixbuf_error_get_type func() GType

	Gdk_interp_type_get_type func() GType

	Gdk_pixbuf_rotation_get_type func() GType

	Gdk_pixbuf_render_threshold_alpha func(
		pixbuf *GdkPixbuf,
		bitmap *GdkBitmap,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int,
		alpha_threshold int)

	Gdk_pixbuf_render_to_drawable func(
		pixbuf *GdkPixbuf,
		drawable *GdkDrawable,
		gc *GdkGC,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int,
		dither GdkRgbDither,
		x_dither int,
		y_dither int)

	Gdk_pixbuf_render_to_drawable_alpha func(
		pixbuf *GdkPixbuf,
		drawable *GdkDrawable,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int,
		alpha_mode GdkPixbufAlphaMode,
		alpha_threshold int,
		dither GdkRgbDither,
		x_dither int,
		y_dither int)

	Gdk_pixbuf_render_pixmap_and_mask_for_colormap func(
		pixbuf *GdkPixbuf,
		colormap *GdkColormap,
		pixmap_return **GdkPixmap,
		mask_return **GdkBitmap,
		alpha_threshold int)

	Gdk_pixbuf_render_pixmap_and_mask func(
		pixbuf *GdkPixbuf,
		pixmap_return **GdkPixmap,
		mask_return **GdkBitmap,
		alpha_threshold int)

	Gdk_pixbuf_get_from_drawable func(
		dest *GdkPixbuf,
		src *GdkDrawable,
		cmap *GdkColormap,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int) *GdkPixbuf

	Gdk_pixbuf_get_from_image func(
		dest *GdkPixbuf,
		src *GdkImage,
		cmap *GdkColormap,
		src_x int,
		src_y int,
		dest_x int,
		dest_y int,
		width int,
		height int) *GdkPixbuf

	Gdk_cairo_create func(
		drawable *GdkDrawable) *Cairo_t

	Gdk_cairo_reset_clip func(
		cr *Cairo_t,
		drawable *GdkDrawable)

	Gdk_cairo_set_source_color func(
		cr *Cairo_t,
		color *GdkColor)

	Gdk_cairo_set_source_pixbuf func(
		cr *Cairo_t,
		pixbuf *GdkPixbuf,
		pixbuf_x Double,
		pixbuf_y Double)

	Gdk_cairo_set_source_pixmap func(
		cr *Cairo_t,
		pixmap *GdkPixmap,
		pixmap_x Double,
		pixmap_y Double)

	Gdk_cairo_set_source_window func(
		cr *Cairo_t,
		window *GdkWindow,
		x Double,
		y Double)

	Gdk_cairo_rectangle func(
		cr *Cairo_t,
		rectangle *GdkRectangle)

	Gdk_cairo_region func(
		cr *Cairo_t,
		region *GdkRegion)

	Gdk_cursor_get_type func() GType

	Gdk_cursor_new_for_display func(
		display *GdkDisplay,
		cursor_type GdkCursorType) *GdkCursor

	Gdk_cursor_new func(
		cursor_type GdkCursorType) *GdkCursor

	Gdk_cursor_new_from_pixmap func(
		source *GdkPixmap,
		mask *GdkPixmap,
		fg *GdkColor,
		bg *GdkColor,
		x Gint,
		y Gint) *GdkCursor

	Gdk_cursor_new_from_pixbuf func(
		display *GdkDisplay,
		pixbuf *GdkPixbuf,
		x Gint,
		y Gint) *GdkCursor

	Gdk_cursor_get_display func(
		cursor *GdkCursor) *GdkDisplay

	Gdk_cursor_ref func(
		cursor *GdkCursor) *GdkCursor

	Gdk_cursor_unref func(
		cursor *GdkCursor)

	Gdk_cursor_new_from_name func(
		display *GdkDisplay,
		name string) *GdkCursor

	Gdk_cursor_get_image func(
		cursor *GdkCursor) *GdkPixbuf

	Gdk_cursor_get_cursor_type func(
		cursor *GdkCursor) GdkCursorType

	Gdk_display_manager_get_type func() GType

	Gdk_display_manager_get func() *GdkDisplayManager

	Gdk_display_manager_get_default_display func(
		display_manager *GdkDisplayManager) *GdkDisplay

	Gdk_display_manager_set_default_display func(
		display_manager *GdkDisplayManager,
		display *GdkDisplay)

	Gdk_display_manager_list_displays func(
		display_manager *GdkDisplayManager) *GSList

	Gdk_gc_get_type func() GType

	Gdk_gc_new func(
		drawable *GdkDrawable) *GdkGC

	Gdk_gc_new_with_values func(
		drawable *GdkDrawable,
		values *GdkGCValues,
		values_mask GdkGCValuesMask) *GdkGC

	Gdk_gc_ref func(
		gc *GdkGC) *GdkGC

	Gdk_gc_unref func(
		gc *GdkGC)

	Gdk_gc_get_values func(
		gc *GdkGC,
		values *GdkGCValues)

	Gdk_gc_set_values func(
		gc *GdkGC,
		values *GdkGCValues,
		values_mask GdkGCValuesMask)

	Gdk_gc_set_foreground func(
		gc *GdkGC,
		color *GdkColor)

	Gdk_gc_set_background func(
		gc *GdkGC,
		color *GdkColor)

	Gdk_gc_set_font func(
		gc *GdkGC,
		font *GdkFont)

	Gdk_gc_set_function func(
		gc *GdkGC,
		function GdkFunction)

	Gdk_gc_set_fill func(
		gc *GdkGC,
		fill GdkFill)

	Gdk_gc_set_tile func(
		gc *GdkGC,
		tile *GdkPixmap)

	Gdk_gc_set_stipple func(
		gc *GdkGC,
		stipple *GdkPixmap)

	Gdk_gc_set_ts_origin func(
		gc *GdkGC,
		x Gint,
		y Gint)

	Gdk_gc_set_clip_origin func(
		gc *GdkGC,
		x Gint,
		y Gint)

	Gdk_gc_set_clip_mask func(
		gc *GdkGC,
		mask *GdkBitmap)

	Gdk_gc_set_clip_rectangle func(
		gc *GdkGC,
		rectangle *GdkRectangle)

	Gdk_gc_set_clip_region func(
		gc *GdkGC,
		region *GdkRegion)

	Gdk_gc_set_subwindow func(
		gc *GdkGC,
		mode GdkSubwindowMode)

	Gdk_gc_set_exposures func(
		gc *GdkGC,
		exposures Gboolean)

	Gdk_gc_set_line_attributes func(
		gc *GdkGC,
		line_width Gint,
		line_style GdkLineStyle,
		cap_style GdkCapStyle,
		join_style GdkJoinStyle)

	Gdk_gc_set_dashes func(
		gc *GdkGC,
		dash_offset Gint,
		dash_list []Gint8,
		n Gint)

	Gdk_gc_offset func(
		gc *GdkGC,
		x_offset Gint,
		y_offset Gint)

	Gdk_gc_copy func(
		dst_gc *GdkGC,
		src_gc *GdkGC)

	Gdk_gc_set_colormap func(
		gc *GdkGC,
		colormap *GdkColormap)

	Gdk_gc_get_colormap func(
		gc *GdkGC) *GdkColormap

	Gdk_gc_set_rgb_fg_color func(
		gc *GdkGC,
		color *GdkColor)

	Gdk_gc_set_rgb_bg_color func(
		gc *GdkGC,
		color *GdkColor)

	Gdk_gc_get_screen func(
		gc *GdkGC) *GdkScreen

	Gdk_drawable_get_type func() GType

	Gdk_drawable_set_data func(
		drawable *GdkDrawable,
		key string,
		data Gpointer,
		destroy_func GDestroyNotify)

	Gdk_drawable_get_data func(
		drawable *GdkDrawable,
		key string) Gpointer

	Gdk_drawable_set_colormap func(
		drawable *GdkDrawable,
		colormap *GdkColormap)

	Gdk_drawable_get_colormap func(
		drawable *GdkDrawable) *GdkColormap

	Gdk_drawable_get_depth func(
		drawable *GdkDrawable) Gint

	Gdk_drawable_get_size func(
		drawable *GdkDrawable,
		width *Gint,
		height *Gint)

	Gdk_drawable_get_visual func(
		drawable *GdkDrawable) *GdkVisual

	Gdk_drawable_get_screen func(
		drawable *GdkDrawable) *GdkScreen

	Gdk_drawable_get_display func(
		drawable *GdkDrawable) *GdkDisplay

	Gdk_drawable_ref func(
		drawable *GdkDrawable) *GdkDrawable

	Gdk_drawable_unref func(
		drawable *GdkDrawable)

	Gdk_draw_point func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint)

	Gdk_draw_line func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x1_ Gint,
		y1_ Gint,
		x2_ Gint,
		y2_ Gint)

	Gdk_draw_rectangle func(
		drawable *GdkDrawable,
		gc *GdkGC,
		filled Gboolean,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gdk_draw_arc func(
		drawable *GdkDrawable,
		gc *GdkGC,
		filled Gboolean,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		angle1 Gint,
		angle2 Gint)

	Gdk_draw_polygon func(
		drawable *GdkDrawable,
		gc *GdkGC,
		filled Gboolean,
		points *GdkPoint,
		n_points Gint)

	Gdk_draw_string func(
		drawable *GdkDrawable,
		font *GdkFont,
		gc *GdkGC,
		x Gint,
		y Gint,
		s string)

	Gdk_draw_text func(
		drawable *GdkDrawable,
		font *GdkFont,
		gc *GdkGC,
		x Gint,
		y Gint,
		text string,
		text_length Gint)

	Gdk_draw_text_wc func(
		drawable *GdkDrawable,
		font *GdkFont,
		gc *GdkGC,
		x Gint,
		y Gint,
		text *GdkWChar,
		text_length Gint)

	Gdk_draw_drawable func(
		drawable *GdkDrawable,
		gc *GdkGC,
		src *GdkDrawable,
		xsrc Gint,
		ysrc Gint,
		xdest Gint,
		ydest Gint,
		width Gint,
		height Gint)

	Gdk_draw_image func(
		drawable *GdkDrawable,
		gc *GdkGC,
		image *GdkImage,
		xsrc Gint,
		ysrc Gint,
		xdest Gint,
		ydest Gint,
		width Gint,
		height Gint)

	Gdk_draw_points func(
		drawable *GdkDrawable,
		gc *GdkGC,
		points *GdkPoint,
		n_points Gint)

	Gdk_draw_segments func(
		drawable *GdkDrawable,
		gc *GdkGC,
		segs *GdkSegment,
		n_segs Gint)

	Gdk_draw_lines func(
		drawable *GdkDrawable,
		gc *GdkGC,
		points *GdkPoint,
		n_points Gint)

	Gdk_draw_pixbuf func(
		drawable *GdkDrawable,
		gc *GdkGC,
		pixbuf *GdkPixbuf,
		src_x Gint,
		src_y Gint,
		dest_x Gint,
		dest_y Gint,
		width Gint,
		height Gint,
		dither GdkRgbDither,
		x_dither Gint,
		y_dither Gint)

	Gdk_draw_glyphs func(
		drawable *GdkDrawable,
		gc *GdkGC,
		font *PangoFont,
		x Gint,
		y Gint,
		glyphs *PangoGlyphString)

	Gdk_draw_layout_line func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		line *PangoLayoutLine)

	Gdk_draw_layout func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		layout *PangoLayout)

	Gdk_draw_layout_line_with_colors func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		line *PangoLayoutLine,
		foreground *GdkColor,
		background *GdkColor)

	Gdk_draw_layout_with_colors func(
		drawable *GdkDrawable,
		gc *GdkGC,
		x Gint,
		y Gint,
		layout *PangoLayout,
		foreground *GdkColor,
		background *GdkColor)

	Gdk_draw_glyphs_transformed func(
		drawable *GdkDrawable,
		gc *GdkGC,
		matrix *PangoMatrix,
		font *PangoFont,
		x Gint,
		y Gint,
		glyphs *PangoGlyphString)

	Gdk_draw_trapezoids func(
		drawable *GdkDrawable,
		gc *GdkGC,
		trapezoids *GdkTrapezoid,
		n_trapezoids Gint)

	Gdk_drawable_get_image func(
		drawable *GdkDrawable,
		x Gint,
		y Gint,
		width Gint,
		height Gint) *GdkImage

	Gdk_drawable_copy_to_image func(
		drawable *GdkDrawable,
		image *GdkImage,
		src_x Gint,
		src_y Gint,
		dest_x Gint,
		dest_y Gint,
		width Gint,
		height Gint) *GdkImage

	Gdk_drawable_get_clip_region func(
		drawable *GdkDrawable) *GdkRegion

	Gdk_drawable_get_visible_region func(
		drawable *GdkDrawable) *GdkRegion

	Gdk_cursor_type_get_type func() GType

	Gdk_drag_action_get_type func() GType

	Gdk_drag_protocol_get_type func() GType

	Gdk_filter_return_get_type func() GType

	Gdk_event_type_get_type func() GType

	Gdk_event_mask_get_type func() GType

	Gdk_visibility_state_get_type func() GType

	Gdk_scroll_direction_get_type func() GType

	Gdk_notify_type_get_type func() GType

	Gdk_crossing_mode_get_type func() GType

	Gdk_property_state_get_type func() GType

	Gdk_window_state_get_type func() GType

	Gdk_setting_action_get_type func() GType

	Gdk_owner_change_get_type func() GType

	Gdk_font_type_get_type func() GType

	Gdk_cap_style_get_type func() GType

	Gdk_fill_get_type func() GType

	Gdk_function_get_type func() GType

	Gdk_join_style_get_type func() GType

	Gdk_line_style_get_type func() GType

	Gdk_subwindow_mode_get_type func() GType

	Gdk_gc_values_mask_get_type func() GType

	Gdk_image_type_get_type func() GType

	Gdk_extension_mode_get_type func() GType

	Gdk_input_source_get_type func() GType

	Gdk_input_mode_get_type func() GType

	Gdk_axis_use_get_type func() GType

	Gdk_prop_mode_get_type func() GType

	Gdk_fill_rule_get_type func() GType

	Gdk_overlap_type_get_type func() GType

	Gdk_rgb_dither_get_type func() GType

	Gdk_byte_order_get_type func() GType

	Gdk_modifier_type_get_type func() GType

	Gdk_input_condition_get_type func() GType

	Gdk_status_get_type func() GType

	Gdk_grab_status_get_type func() GType

	Gdk_visual_type_get_type func() GType

	Gdk_window_class_get_type func() GType

	Gdk_window_type_get_type func() GType

	Gdk_window_attributes_type_get_type func() GType

	Gdk_window_hints_get_type func() GType

	Gdk_window_type_hint_get_type func() GType

	Gdk_wm_decoration_get_type func() GType

	Gdk_wm_function_get_type func() GType

	Gdk_gravity_get_type func() GType

	Gdk_window_edge_get_type func() GType

	Gdk_font_get_type func() GType

	Gdk_font_ref func(
		font *GdkFont) *GdkFont

	Gdk_font_unref func(
		font *GdkFont)

	Gdk_font_id func(
		font *GdkFont) Gint

	Gdk_font_equal func(
		fonta *GdkFont,
		fontb *GdkFont) Gboolean

	Gdk_font_load_for_display func(
		display *GdkDisplay,
		font_name string) *GdkFont

	Gdk_fontset_load_for_display func(
		display *GdkDisplay,
		fontset_name string) *GdkFont

	Gdk_font_from_description_for_display func(
		display *GdkDisplay,
		font_desc *PangoFontDescription) *GdkFont

	Gdk_font_load func(
		font_name string) *GdkFont

	Gdk_fontset_load func(
		fontset_name string) *GdkFont

	Gdk_font_from_description func(
		font_desc *PangoFontDescription) *GdkFont

	Gdk_string_width func(
		font *GdkFont,
		s string) Gint

	Gdk_text_width func(
		font *GdkFont,
		text string,
		text_length Gint) Gint

	Gdk_text_width_wc func(
		font *GdkFont,
		text *GdkWChar,
		text_length Gint) Gint

	Gdk_char_width func(
		font *GdkFont,
		character Gchar) Gint

	Gdk_char_width_wc func(
		font *GdkFont,
		character GdkWChar) Gint

	Gdk_string_measure func(
		font *GdkFont,
		s string) Gint

	Gdk_text_measure func(
		font *GdkFont,
		text string,
		text_length Gint) Gint

	Gdk_char_measure func(
		font *GdkFont,
		character Gchar) Gint

	Gdk_string_height func(
		font *GdkFont,
		s string) Gint

	Gdk_text_height func(
		font *GdkFont,
		text string,
		text_length Gint) Gint

	Gdk_char_height func(
		font *GdkFont,
		character Gchar) Gint

	Gdk_text_extents func(
		font *GdkFont,
		text string,
		text_length Gint,
		lbearing *Gint,
		rbearing *Gint,
		width *Gint,
		ascent *Gint,
		descent *Gint)

	Gdk_text_extents_wc func(
		font *GdkFont,
		text *GdkWChar,
		text_length Gint,
		lbearing *Gint,
		rbearing *Gint,
		width *Gint,
		ascent *Gint,
		descent *Gint)

	Gdk_string_extents func(
		font *GdkFont,
		s string,
		lbearing *Gint,
		rbearing *Gint,
		width *Gint,
		ascent *Gint,
		descent *Gint)

	Gdk_font_get_display func(
		font *GdkFont) *GdkDisplay

	Gdk_image_get_type func() GType

	Gdk_image_new func(
		typ GdkImageType,
		visual *GdkVisual,
		width Gint,
		height Gint) *GdkImage

	Gdk_image_get func(
		drawable *GdkDrawable,
		x Gint,
		y Gint,
		width Gint,
		height Gint) *GdkImage

	Gdk_image_ref func(
		image *GdkImage) *GdkImage

	Gdk_image_unref func(
		image *GdkImage)

	Gdk_image_put_pixel func(
		image *GdkImage,
		x Gint,
		y Gint,
		pixel Guint32)

	Gdk_image_get_pixel func(
		image *GdkImage,
		x Gint,
		y Gint) Guint32

	Gdk_image_set_colormap func(
		image *GdkImage,
		colormap *GdkColormap)

	Gdk_image_get_colormap func(
		image *GdkImage) *GdkColormap

	Gdk_image_get_image_type func(
		image *GdkImage) GdkImageType

	Gdk_image_get_visual func(
		image *GdkImage) *GdkVisual

	Gdk_image_get_byte_order func(
		image *GdkImage) GdkByteOrder

	Gdk_image_get_width func(
		image *GdkImage) Gint

	Gdk_image_get_height func(
		image *GdkImage) Gint

	Gdk_image_get_depth func(
		image *GdkImage) Guint16

	Gdk_image_get_bytes_per_pixel func(
		image *GdkImage) Guint16

	Gdk_image_get_bytes_per_line func(
		image *GdkImage) Guint16

	Gdk_image_get_bits_per_pixel func(
		image *GdkImage) Guint16

	Gdk_image_get_pixels func(
		image *GdkImage) Gpointer

	Gdk_keymap_get_type func() GType

	Gdk_keymap_get_default func() *GdkKeymap

	Gdk_keymap_get_for_display func(
		display *GdkDisplay) *GdkKeymap

	Gdk_keymap_lookup_key func(
		keymap *GdkKeymap,
		key *GdkKeymapKey) Guint

	Gdk_keymap_translate_keyboard_state func(
		keymap *GdkKeymap,
		hardware_keycode Guint,
		state GdkModifierType,
		group Gint,
		keyval *Guint,
		effective_group *Gint,
		level *Gint,
		consumed_modifiers *GdkModifierType) Gboolean

	Gdk_keymap_get_entries_for_keyval func(
		keymap *GdkKeymap,
		keyval Guint,
		keys **GdkKeymapKey,
		n_keys *Gint) Gboolean

	Gdk_keymap_get_entries_for_keycode func(
		keymap *GdkKeymap,
		hardware_keycode Guint,
		keys **GdkKeymapKey,
		keyvals **Guint,
		n_entries *Gint) Gboolean

	Gdk_keymap_get_direction func(
		keymap *GdkKeymap) PangoDirection

	Gdk_keymap_have_bidi_layouts func(
		keymap *GdkKeymap) Gboolean

	Gdk_keymap_get_caps_lock_state func(
		keymap *GdkKeymap) Gboolean

	Gdk_keymap_add_virtual_modifiers func(
		keymap *GdkKeymap,
		state *GdkModifierType)

	Gdk_keymap_map_virtual_modifiers func(
		keymap *GdkKeymap,
		state *GdkModifierType) Gboolean

	Gdk_keyval_name func(
		keyval Guint) string

	Gdk_keyval_from_name func(
		keyval_name string) Guint

	Gdk_keyval_convert_case func(
		symbol Guint,
		lower *Guint,
		upper *Guint)

	Gdk_keyval_to_upper func(
		keyval Guint) Guint

	Gdk_keyval_to_lower func(
		keyval Guint) Guint

	Gdk_keyval_is_upper func(
		keyval Guint) Gboolean

	Gdk_keyval_is_lower func(
		keyval Guint) Gboolean

	Gdk_keyval_to_unicode func(
		keyval Guint) Guint32

	Gdk_unicode_to_keyval func(
		wc Guint32) Guint

	Gdk_pango_renderer_get_type func() GType

	Gdk_pango_renderer_new func(
		screen *GdkScreen) *PangoRenderer

	Gdk_pango_renderer_get_default func(
		screen *GdkScreen) *PangoRenderer

	Gdk_pango_renderer_set_drawable func(

		Gdk_renderer *GdkPangoRenderer,
		drawable *GdkDrawable)

	Gdk_pango_renderer_set_gc func(

		Gdk_renderer *GdkPangoRenderer,
		gc *GdkGC)

	Gdk_pango_renderer_set_stipple func(

		Gdk_renderer *GdkPangoRenderer,
		part PangoRenderPart,
		stipple *GdkBitmap)

	Gdk_pango_renderer_set_override_color func(

		Gdk_renderer *GdkPangoRenderer,
		part PangoRenderPart,
		color *GdkColor)

	Gdk_pango_context_get_for_screen func(
		screen *GdkScreen) *PangoContext

	Gdk_pango_context_get func() *PangoContext

	Gdk_pango_context_set_colormap func(
		context *PangoContext,
		colormap *GdkColormap)

	Gdk_pango_layout_line_get_clip_region func(
		line *PangoLayoutLine,
		x_origin Gint,
		y_origin Gint,
		index_ranges *Gint,
		n_ranges Gint) *GdkRegion

	Gdk_pango_layout_get_clip_region func(
		layout *PangoLayout,
		x_origin Gint,
		y_origin Gint,
		index_ranges *Gint,
		n_ranges Gint) *GdkRegion

	Gdk_pango_attr_stipple_new func(
		stipple *GdkBitmap) *PangoAttribute

	Gdk_pango_attr_embossed_new func(
		embossed Gboolean) *PangoAttribute

	Gdk_pango_attr_emboss_color_new func(
		color *GdkColor) *PangoAttribute

	Gdk_pixmap_get_type func() GType

	Gdk_pixmap_new func(
		drawable *GdkDrawable,
		width Gint,
		height Gint,
		depth Gint) *GdkPixmap

	Gdk_bitmap_create_from_data func(
		drawable *GdkDrawable,
		data string,
		width Gint,
		height Gint) *GdkBitmap

	Gdk_pixmap_create_from_data func(
		drawable *GdkDrawable,
		data string,
		width Gint,
		height Gint,
		depth Gint,
		fg *GdkColor,
		bg *GdkColor) *GdkPixmap

	Gdk_pixmap_create_from_xpm func(
		drawable *GdkDrawable,
		mask **GdkBitmap,
		transparent_color *GdkColor,
		filename string) *GdkPixmap

	Gdk_pixmap_colormap_create_from_xpm func(
		drawable *GdkDrawable,
		colormap *GdkColormap,
		mask **GdkBitmap,
		transparent_color *GdkColor,
		filename string) *GdkPixmap

	Gdk_pixmap_create_from_xpm_d func(
		drawable *GdkDrawable,
		mask **GdkBitmap,
		transparent_color *GdkColor,
		data **Gchar) *GdkPixmap

	Gdk_pixmap_colormap_create_from_xpm_d func(
		drawable *GdkDrawable,
		colormap *GdkColormap,
		mask **GdkBitmap,
		transparent_color *GdkColor,
		data **Gchar) *GdkPixmap

	Gdk_pixmap_get_size func(
		pixmap *GdkPixmap,
		width *Gint,
		height *Gint)

	Gdk_pixmap_foreign_new func(
		anid GdkNativeWindow) *GdkPixmap

	Gdk_pixmap_lookup func(
		anid GdkNativeWindow) *GdkPixmap

	Gdk_pixmap_foreign_new_for_display func(
		display *GdkDisplay,
		anid GdkNativeWindow) *GdkPixmap

	Gdk_pixmap_lookup_for_display func(
		display *GdkDisplay,
		anid GdkNativeWindow) *GdkPixmap

	Gdk_pixmap_foreign_new_for_screen func(
		screen *GdkScreen,
		anid GdkNativeWindow,
		width Gint,
		height Gint,
		depth Gint) *GdkPixmap

	Gdk_atom_intern func(
		atom_name string,
		only_if_exists Gboolean) GdkAtom

	Gdk_atom_intern_static_string func(
		atom_name string) GdkAtom

	Gdk_atom_name func(
		atom GdkAtom) string

	Gdk_property_get func(
		window *GdkWindow,
		property GdkAtom,
		typ GdkAtom,
		offset Gulong,
		length Gulong,
		pdelete Gint,
		actual_property_type *GdkAtom,
		actual_format *Gint,
		actual_length *Gint,
		data **Guchar) Gboolean

	Gdk_property_change func(
		window *GdkWindow,
		property GdkAtom,
		typ GdkAtom,
		format Gint,
		mode GdkPropMode,
		data *Guchar,
		nelements Gint)

	Gdk_property_delete func(
		window *GdkWindow,
		property GdkAtom)

	Gdk_text_property_to_text_list func(
		encoding GdkAtom,
		format Gint,
		text *Guchar,
		length Gint,
		list ***Gchar) Gint

	Gdk_utf8_to_compound_text func(
		str string,
		encoding *GdkAtom,
		format *Gint,
		ctext **Guchar,
		length *Gint) Gboolean

	Gdk_string_to_compound_text func(
		str string,
		encoding *GdkAtom,
		format *Gint,
		ctext **Guchar,
		length *Gint) Gint

	Gdk_text_property_to_utf8_list func(
		encoding GdkAtom,
		format Gint,
		text *Guchar,
		length Gint,
		list ***Gchar) Gint

	Gdk_text_property_to_utf8_list_for_display func(
		display *GdkDisplay,
		encoding GdkAtom,
		format Gint,
		text *Guchar,
		length Gint,
		list ***Gchar) Gint

	Gdk_utf8_to_string_target func(
		str string) string

	Gdk_text_property_to_text_list_for_display func(
		display *GdkDisplay,
		encoding GdkAtom,
		format Gint,
		text *Guchar,
		length Gint,
		list ***Gchar) Gint

	Gdk_string_to_compound_text_for_display func(
		display *GdkDisplay,
		str string,
		encoding *GdkAtom,
		format *Gint,
		ctext **Guchar,
		length *Gint) Gint

	Gdk_utf8_to_compound_text_for_display func(
		display *GdkDisplay,
		str string,
		encoding *GdkAtom,
		format *Gint,
		ctext **Guchar,
		length *Gint) Gboolean

	Gdk_free_text_list func(
		list **Gchar)

	Gdk_free_compound_text func(
		ctext *Guchar)

	Gdk_region_new func() *GdkRegion

	Gdk_region_polygon func(
		points *GdkPoint,
		n_points Gint,
		fill_rule GdkFillRule) *GdkRegion

	Gdk_region_copy func(
		region *GdkRegion) *GdkRegion

	Gdk_region_rectangle func(
		rectangle *GdkRectangle) *GdkRegion

	Gdk_region_destroy func(
		region *GdkRegion)

	Gdk_region_get_clipbox func(
		region *GdkRegion,
		rectangle *GdkRectangle)

	Gdk_region_get_rectangles func(
		region *GdkRegion,
		rectangles **GdkRectangle,
		n_rectangles *Gint)

	Gdk_region_empty func(
		region *GdkRegion) Gboolean

	Gdk_region_equal func(
		region1 *GdkRegion,
		region2 *GdkRegion) Gboolean

	Gdk_region_rect_equal func(
		region *GdkRegion,
		rectangle *GdkRectangle) Gboolean

	Gdk_region_point_in func(
		region *GdkRegion,
		x int,
		y int) Gboolean

	Gdk_region_rect_in func(
		region *GdkRegion,
		rectangle *GdkRectangle) GdkOverlapType

	Gdk_region_offset func(
		region *GdkRegion,
		dx Gint,
		dy Gint)

	Gdk_region_shrink func(
		region *GdkRegion,
		dx Gint,
		dy Gint)

	Gdk_region_union_with_rect func(
		region *GdkRegion,
		rect *GdkRectangle)

	Gdk_region_intersect func(
		source1 *GdkRegion,
		source2 *GdkRegion)

	Gdk_region_union func(
		source1 *GdkRegion,
		source2 *GdkRegion)

	Gdk_region_subtract func(
		source1 *GdkRegion,
		source2 *GdkRegion)

	Gdk_region_xor func(
		source1 *GdkRegion,
		source2 *GdkRegion)

	Gdk_region_spans_intersect_foreach func(
		region *GdkRegion,
		spans *GdkSpan,
		n_spans int,
		sorted Gboolean,
		function GdkSpanFunc,
		data Gpointer)

	Gdk_selection_owner_set func(
		owner *GdkWindow,
		selection GdkAtom,
		time_ Guint32,
		send_event Gboolean) Gboolean

	Gdk_selection_owner_get func(
		selection GdkAtom) *GdkWindow

	Gdk_selection_owner_set_for_display func(
		display *GdkDisplay,
		owner *GdkWindow,
		selection GdkAtom,
		time_ Guint32,
		send_event Gboolean) Gboolean

	Gdk_selection_owner_get_for_display func(
		display *GdkDisplay,
		selection GdkAtom) *GdkWindow

	Gdk_selection_convert func(
		requestor *GdkWindow,
		selection GdkAtom,
		target GdkAtom,
		time_ Guint32)

	Gdk_selection_property_get func(
		requestor *GdkWindow,
		data **Guchar,
		prop_type *GdkAtom,
		prop_format *Gint) Gint

	Gdk_selection_send_notify func(
		requestor GdkNativeWindow,
		selection GdkAtom,
		target GdkAtom,
		property GdkAtom,
		time_ Guint32)

	Gdk_selection_send_notify_for_display func(
		display *GdkDisplay,
		requestor GdkNativeWindow,
		selection GdkAtom,
		target GdkAtom,
		property GdkAtom,
		time_ Guint32)

	Gdk_spawn_on_screen func(
		screen *GdkScreen,
		working_directory string,
		argv **Gchar,
		envp **Gchar,
		flags GSpawnFlags,
		child_setup GSpawnChildSetupFunc,
		user_data Gpointer,
		child_pid *Gint,
		e **GError) Gboolean

	Gdk_spawn_on_screen_with_pipes func(
		screen *GdkScreen,
		working_directory string,
		argv **Gchar,
		envp **Gchar,
		flags GSpawnFlags,
		child_setup GSpawnChildSetupFunc,
		user_data Gpointer,
		child_pid *Gint,
		standard_input *Gint,
		standard_output *Gint,
		standard_error *Gint,
		e **GError) Gboolean

	Gdk_spawn_command_line_on_screen func(
		screen *GdkScreen,
		command_line string,
		e **GError) Gboolean

	Gdk_window_object_get_type func() GType

	Gdk_window_new func(
		parent *GdkWindow,
		attributes *GdkWindowAttr,
		attributes_mask Gint) *GdkWindow

	Gdk_window_destroy func(
		window *GdkWindow)

	Gdk_window_get_window_type func(
		window *GdkWindow) GdkWindowType

	Gdk_window_is_destroyed func(
		window *GdkWindow) Gboolean

	Gdk_window_get_screen func(
		window *GdkWindow) *GdkScreen

	Gdk_window_get_display func(
		window *GdkWindow) *GdkDisplay

	Gdk_window_get_visual func(
		window *GdkWindow) *GdkVisual

	Gdk_window_get_width func(
		window *GdkWindow) int

	Gdk_window_get_height func(
		window *GdkWindow) int

	Gdk_window_at_pointer func(
		win_x *Gint,
		win_y *Gint) *GdkWindow

	Gdk_window_show func(
		window *GdkWindow)

	Gdk_window_hide func(
		window *GdkWindow)

	Gdk_window_withdraw func(
		window *GdkWindow)

	Gdk_window_show_unraised func(
		window *GdkWindow)

	Gdk_window_move func(
		window *GdkWindow,
		x Gint,
		y Gint)

	Gdk_window_resize func(
		window *GdkWindow,
		width Gint,
		height Gint)

	Gdk_window_move_resize func(
		window *GdkWindow,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gdk_window_reparent func(
		window *GdkWindow,
		new_parent *GdkWindow,
		x Gint,
		y Gint)

	Gdk_window_clear func(
		window *GdkWindow)

	Gdk_window_clear_area func(
		window *GdkWindow,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gdk_window_clear_area_e func(
		window *GdkWindow,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gdk_window_raise func(
		window *GdkWindow)

	Gdk_window_lower func(
		window *GdkWindow)

	Gdk_window_restack func(
		window *GdkWindow,
		sibling *GdkWindow,
		above Gboolean)

	Gdk_window_focus func(
		window *GdkWindow,
		timestamp Guint32)

	Gdk_window_set_user_data func(
		window *GdkWindow,
		user_data Gpointer)

	Gdk_window_set_override_redirect func(
		window *GdkWindow,
		override_redirect Gboolean)

	Gdk_window_get_accept_focus func(
		window *GdkWindow) Gboolean

	Gdk_window_set_accept_focus func(
		window *GdkWindow,
		accept_focus Gboolean)

	Gdk_window_get_focus_on_map func(
		window *GdkWindow) Gboolean

	Gdk_window_set_focus_on_map func(
		window *GdkWindow,
		focus_on_map Gboolean)

	Gdk_window_add_filter func(
		window *GdkWindow,
		function GdkFilterFunc,
		data Gpointer)

	Gdk_window_remove_filter func(
		window *GdkWindow,
		function GdkFilterFunc,
		data Gpointer)

	Gdk_window_scroll func(
		window *GdkWindow,
		dx Gint,
		dy Gint)

	Gdk_window_move_region func(
		window *GdkWindow,
		region *GdkRegion,
		dx Gint,
		dy Gint)

	Gdk_window_ensure_native func(
		window *GdkWindow) Gboolean

	Gdk_window_shape_combine_mask func(
		window *GdkWindow,
		mask *GdkBitmap,
		x Gint,
		y Gint)

	Gdk_window_shape_combine_region func(
		window *GdkWindow,
		shape_region *GdkRegion,
		offset_x Gint,
		offset_y Gint)

	Gdk_window_set_child_shapes func(
		window *GdkWindow)

	Gdk_window_get_composited func(
		window *GdkWindow) Gboolean

	Gdk_window_set_composited func(
		window *GdkWindow,
		composited Gboolean)

	Gdk_window_merge_child_shapes func(
		window *GdkWindow)

	Gdk_window_input_shape_combine_mask func(
		window *GdkWindow,
		mask *GdkBitmap,
		x Gint,
		y Gint)

	Gdk_window_input_shape_combine_region func(
		window *GdkWindow,
		shape_region *GdkRegion,
		offset_x Gint,
		offset_y Gint)

	Gdk_window_set_child_input_shapes func(
		window *GdkWindow)

	Gdk_window_merge_child_input_shapes func(
		window *GdkWindow)

	Gdk_window_is_visible func(
		window *GdkWindow) Gboolean

	Gdk_window_is_viewable func(
		window *GdkWindow) Gboolean

	Gdk_window_is_input_only func(
		window *GdkWindow) Gboolean

	Gdk_window_is_shaped func(
		window *GdkWindow) Gboolean

	Gdk_window_get_state func(
		window *GdkWindow) GdkWindowState

	Gdk_window_set_static_gravities func(
		window *GdkWindow,
		use_static Gboolean) Gboolean

	Gdk_window_foreign_new func(
		anid GdkNativeWindow) *GdkWindow

	Gdk_window_lookup func(
		anid GdkNativeWindow) *GdkWindow

	Gdk_window_foreign_new_for_display func(
		display *GdkDisplay,
		anid GdkNativeWindow) *GdkWindow

	Gdk_window_lookup_for_display func(
		display *GdkDisplay,
		anid GdkNativeWindow) *GdkWindow

	Gdk_window_has_native func(
		window *GdkWindow) Gboolean

	Gdk_window_set_hints func(
		window *GdkWindow,
		x Gint,
		y Gint,
		min_width Gint,
		min_height Gint,
		max_width Gint,
		max_height Gint,
		flags Gint)

	Gdk_window_set_type_hint func(
		window *GdkWindow,
		hint GdkWindowTypeHint)

	Gdk_window_get_type_hint func(
		window *GdkWindow) GdkWindowTypeHint

	Gdk_window_get_modal_hint func(
		window *GdkWindow) Gboolean

	Gdk_window_set_modal_hint func(
		window *GdkWindow,
		modal Gboolean)

	Gdk_window_set_skip_taskbar_hint func(
		window *GdkWindow,
		skips_taskbar Gboolean)

	Gdk_window_set_skip_pager_hint func(
		window *GdkWindow,
		skips_pager Gboolean)

	Gdk_window_set_urgency_hint func(
		window *GdkWindow,
		urgent Gboolean)

	Gdk_window_set_geometry_hints func(
		window *GdkWindow,
		geometry *GdkGeometry,
		geom_mask GdkWindowHints)

	Gdk_set_sm_client_id func(
		sm_client_id string)

	Gdk_window_begin_paint_rect func(
		window *GdkWindow,
		rectangle *GdkRectangle)

	Gdk_window_begin_paint_region func(
		window *GdkWindow,
		region *GdkRegion)

	Gdk_window_end_paint func(
		window *GdkWindow)

	Gdk_window_flush func(
		window *GdkWindow)

	Gdk_window_set_title func(
		window *GdkWindow,
		title string)

	Gdk_window_set_role func(
		window *GdkWindow,
		role string)

	Gdk_window_set_startup_id func(
		window *GdkWindow,
		startup_id string)

	Gdk_window_set_transient_for func(
		window *GdkWindow,
		parent *GdkWindow)

	Gdk_window_set_background func(
		window *GdkWindow,
		color *GdkColor)

	Gdk_window_set_back_pixmap func(
		window *GdkWindow,
		pixmap *GdkPixmap,
		parent_relative Gboolean)

	Gdk_window_get_background_pattern func(
		window *GdkWindow) *Cairo_pattern_t

	Gdk_window_set_cursor func(
		window *GdkWindow,
		cursor *GdkCursor)

	Gdk_window_get_cursor func(
		window *GdkWindow) *GdkCursor

	Gdk_window_get_user_data func(
		window *GdkWindow,
		data *Gpointer)

	Gdk_window_get_geometry func(
		window *GdkWindow,
		x *Gint,
		y *Gint,
		width *Gint,
		height *Gint,
		depth *Gint)

	Gdk_window_get_position func(
		window *GdkWindow,
		x *Gint,
		y *Gint)

	Gdk_window_get_origin func(
		window *GdkWindow,
		x *Gint,
		y *Gint) Gint

	Gdk_window_get_root_coords func(
		window *GdkWindow,
		x Gint,
		y Gint,
		root_x *Gint,
		root_y *Gint)

	Gdk_window_coords_to_parent func(
		window *GdkWindow,
		x Gdouble,
		y Gdouble,
		parent_x *Gdouble,
		parent_y *Gdouble)

	Gdk_window_coords_from_parent func(
		window *GdkWindow,
		parent_x Gdouble,
		parent_y Gdouble,
		x *Gdouble,
		y *Gdouble)

	Gdk_window_get_deskrelative_origin func(
		window *GdkWindow,
		x *Gint,
		y *Gint) Gboolean

	Gdk_window_get_root_origin func(
		window *GdkWindow,
		x *Gint,
		y *Gint)

	Gdk_window_get_frame_extents func(
		window *GdkWindow,
		rect *GdkRectangle)

	Gdk_window_get_pointer func(
		window *GdkWindow,
		x *Gint,
		y *Gint,
		mask *GdkModifierType) *GdkWindow

	Gdk_window_get_parent func(
		window *GdkWindow) *GdkWindow

	Gdk_window_get_toplevel func(
		window *GdkWindow) *GdkWindow

	Gdk_window_get_effective_parent func(
		window *GdkWindow) *GdkWindow

	Gdk_window_get_effective_toplevel func(
		window *GdkWindow) *GdkWindow

	Gdk_window_get_children func(
		window *GdkWindow) *GList

	Gdk_window_peek_children func(
		window *GdkWindow) *GList

	Gdk_window_get_events func(
		window *GdkWindow) GdkEventMask

	Gdk_window_set_events func(
		window *GdkWindow,
		event_mask GdkEventMask)

	Gdk_window_set_icon_list func(
		window *GdkWindow,
		pixbufs *GList)

	Gdk_window_set_icon func(
		window *GdkWindow,
		icon_window *GdkWindow,
		pixmap *GdkPixmap,
		mask *GdkBitmap)

	Gdk_window_set_icon_name func(
		window *GdkWindow,
		name string)

	Gdk_window_set_group func(
		window *GdkWindow,
		leader *GdkWindow)

	Gdk_window_get_group func(
		window *GdkWindow) *GdkWindow

	Gdk_window_set_decorations func(
		window *GdkWindow,
		decorations GdkWMDecoration)

	Gdk_window_get_decorations func(
		window *GdkWindow,
		decorations *GdkWMDecoration) Gboolean

	Gdk_window_set_functions func(
		window *GdkWindow,
		functions GdkWMFunction)

	Gdk_window_get_toplevels func() *GList

	Gdk_window_create_similar_surface func(
		window *GdkWindow,
		content Cairo_content_t,
		width int,
		height int) *Cairo_surface_t

	Gdk_window_beep func(
		window *GdkWindow)

	Gdk_window_iconify func(
		window *GdkWindow)

	Gdk_window_deiconify func(
		window *GdkWindow)

	Gdk_window_stick func(
		window *GdkWindow)

	Gdk_window_unstick func(
		window *GdkWindow)

	Gdk_window_maximize func(
		window *GdkWindow)

	Gdk_window_unmaximize func(
		window *GdkWindow)

	Gdk_window_fullscreen func(
		window *GdkWindow)

	Gdk_window_unfullscreen func(
		window *GdkWindow)

	Gdk_window_set_keep_above func(
		window *GdkWindow,
		setting Gboolean)

	Gdk_window_set_keep_below func(
		window *GdkWindow,
		setting Gboolean)

	Gdk_window_set_opacity func(
		window *GdkWindow,
		opacity Gdouble)

	Gdk_window_register_dnd func(
		window *GdkWindow)

	Gdk_window_begin_resize_drag func(
		window *GdkWindow,
		edge GdkWindowEdge,
		button Gint,
		root_x Gint,
		root_y Gint,
		timestamp Guint32)

	Gdk_window_begin_move_drag func(
		window *GdkWindow,
		button Gint,
		root_x Gint,
		root_y Gint,
		timestamp Guint32)

	Gdk_window_invalidate_rect func(
		window *GdkWindow,
		rect *GdkRectangle,
		invalidate_children Gboolean)

	Gdk_window_invalidate_region func(
		window *GdkWindow,
		region *GdkRegion,
		invalidate_children Gboolean)

	Gdk_window_invalidate_maybe_recurse func(
		window *GdkWindow,
		region *GdkRegion,
		child_func func(
			*GdkWindow,
			Gpointer) Gboolean,
		user_data Gpointer)

	Gdk_window_get_update_area func(
		window *GdkWindow) *GdkRegion

	Gdk_window_freeze_updates func(
		window *GdkWindow)

	Gdk_window_thaw_updates func(
		window *GdkWindow)

	Gdk_window_freeze_toplevel_updates_libgtk_only func(
		window *GdkWindow)

	Gdk_window_thaw_toplevel_updates_libgtk_only func(
		window *GdkWindow)

	Gdk_window_process_all_updates func()

	Gdk_window_process_updates func(
		window *GdkWindow,
		update_children Gboolean)

	Gdk_window_set_debug_updates func(
		setting Gboolean)

	Gdk_window_constrain_size func(
		geometry *GdkGeometry,
		flags Guint,
		width Gint,
		height Gint,
		new_width *Gint,
		new_height *Gint)

	Gdk_window_get_internal_paint_info func(
		window *GdkWindow,
		real_drawable **GdkDrawable,
		x_offset *Gint,
		y_offset *Gint)

	Gdk_window_enable_synchronized_configure func(
		window *GdkWindow)

	Gdk_window_configure_finished func(
		window *GdkWindow)

	Gdk_get_default_root_window func() *GdkWindow

	Gdk_offscreen_window_get_pixmap func(
		window *GdkWindow) *GdkPixmap

	Gdk_offscreen_window_set_embedder func(
		window *GdkWindow,
		embedder *GdkWindow)

	Gdk_offscreen_window_get_embedder func(
		window *GdkWindow) *GdkWindow

	Gdk_window_geometry_changed func(
		window *GdkWindow)

	Gdk_window_redirect_to_drawable func(
		window *GdkWindow,
		drawable *GdkDrawable,
		src_x Gint,
		src_y Gint,
		dest_x Gint,
		dest_y Gint,
		width Gint,
		height Gint)

	Gdk_window_remove_redirection func(
		window *GdkWindow)

	Gdk_set_pointer_hooks func(
		new_hooks *GdkPointerHooks) *GdkPointerHooks

	Gdk_test_render_sync func(
		window *GdkWindow)

	Gdk_test_simulate_key func(
		window *GdkWindow,
		x Gint,
		y Gint,
		keyval Guint,
		modifiers GdkModifierType,
		key_pressrelease GdkEventType) Gboolean

	Gdk_test_simulate_button func(
		window *GdkWindow,
		x Gint,
		y Gint,
		button Guint,
		modifiers GdkModifierType,
		button_pressrelease GdkEventType) Gboolean

	Gdk_visual_get_type func() GType

	Gdk_visual_get_best_depth func() Gint

	Gdk_visual_get_best_type func() GdkVisualType

	Gdk_visual_get_system func() *GdkVisual

	Gdk_visual_get_best func() *GdkVisual

	Gdk_visual_get_best_with_depth func(
		depth Gint) *GdkVisual

	Gdk_visual_get_best_with_type func(
		visual_type GdkVisualType) *GdkVisual

	Gdk_visual_get_best_with_both func(
		depth Gint,
		visual_type GdkVisualType) *GdkVisual

	Gdk_query_depths func(
		depths **Gint,
		count *Gint)

	Gdk_query_visual_types func(
		visual_types **GdkVisualType,
		count *Gint)

	Gdk_list_visuals func() *GList

	Gdk_visual_get_screen func(
		visual *GdkVisual) *GdkScreen

	Gdk_visual_get_visual_type func(
		visual *GdkVisual) GdkVisualType

	Gdk_visual_get_depth func(
		visual *GdkVisual) Gint

	Gdk_visual_get_byte_order func(
		visual *GdkVisual) GdkByteOrder

	Gdk_visual_get_colormap_size func(
		visual *GdkVisual) Gint

	Gdk_visual_get_bits_per_rgb func(
		visual *GdkVisual) Gint

	Gdk_visual_get_red_pixel_details func(
		visual *GdkVisual,
		mask *Guint32,
		shift *Gint,
		precision *Gint)

	Gdk_visual_get_green_pixel_details func(
		visual *GdkVisual,
		mask *Guint32,
		shift *Gint,
		precision *Gint)

	Gdk_visual_get_blue_pixel_details func(
		visual *GdkVisual,
		mask *Guint32,
		shift *Gint,
		precision *Gint)

	Gdk_parse_args func(
		argc *Gint,
		argv ***Gchar)

	Gdk_init func(
		argc *Gint,
		argv ***Gchar)

	Gdk_init_check func(
		argc *Gint,
		argv ***Gchar) Gboolean

	Gdk_add_option_entries_libgtk_only func(
		group *GOptionGroup)

	Gdk_pre_parse_libgtk_only func()

	Gdk_exit func(
		error_code Gint)

	Gdk_set_locale func() string

	Gdk_get_program_class func() string

	Gdk_set_program_class func(
		program_class string)

	Gdk_error_trap_push func()

	Gdk_error_trap_pop func() Gint

	Gdk_set_use_xshm func(
		use_xshm Gboolean)

	Gdk_get_use_xshm func() Gboolean

	Gdk_get_display func() string

	Gdk_get_display_arg_name func() string

	Gdk_input_add_full func(
		source Gint,
		condition GdkInputCondition,
		function GdkInputFunction,
		data Gpointer,
		destroy GDestroyNotify) Gint

	Gdk_input_add func(
		source Gint,
		condition GdkInputCondition,
		function GdkInputFunction,
		data Gpointer) Gint

	Gdk_input_remove func(
		tag Gint)

	Gdk_pointer_grab func(
		window *GdkWindow,
		owner_events Gboolean,
		event_mask GdkEventMask,
		confine_to *GdkWindow,
		cursor *GdkCursor,
		time_ Guint32) GdkGrabStatus

	Gdk_keyboard_grab func(
		window *GdkWindow,
		owner_events Gboolean,
		time_ Guint32) GdkGrabStatus

	Gdk_pointer_grab_info_libgtk_only func(
		display *GdkDisplay,
		grab_window **GdkWindow,
		owner_events *Gboolean) Gboolean

	Gdk_keyboard_grab_info_libgtk_only func(
		display *GdkDisplay,
		grab_window **GdkWindow,
		owner_events *Gboolean) Gboolean

	Gdk_pointer_ungrab func(
		time_ Guint32)

	Gdk_keyboard_ungrab func(
		time_ Guint32)

	Gdk_pointer_is_grabbed func() Gboolean

	Gdk_screen_width func() Gint

	Gdk_screen_height func() Gint

	Gdk_screen_width_mm func() Gint

	Gdk_screen_height_mm func() Gint

	Gdk_beep func()

	Gdk_flush func()

	Gdk_set_double_click_time func(
		msec Guint)

	Gdk_rectangle_intersect func(
		src1 *GdkRectangle,
		src2 *GdkRectangle,
		dest *GdkRectangle) Gboolean

	Gdk_rectangle_union func(
		src1 *GdkRectangle,
		src2 *GdkRectangle,
		dest *GdkRectangle)

	Gdk_rectangle_get_type func() GType

	Gdk_wcstombs func(
		src *GdkWChar) string

	Gdk_mbstowcs func(
		dest *GdkWChar,
		src string,
		dest_max Gint) Gint

	Gdk_event_send_client_message func(
		event *GdkEvent,
		winid GdkNativeWindow) Gboolean

	Gdk_event_send_clientmessage_toall func(
		event *GdkEvent)

	Gdk_event_send_client_message_for_display func(
		display *GdkDisplay,
		event *GdkEvent,
		winid GdkNativeWindow) Gboolean

	Gdk_notify_startup_complete func()

	Gdk_notify_startup_complete_with_id func(
		startup_id string)

	Gdk_threads_enter func()

	Gdk_threads_leave func()

	Gdk_threads_init func()

	Gdk_threads_set_lock_functions func(
		enter_fn GCallback,
		leave_fn GCallback)

	Gdk_threads_add_idle_full func(
		priority Gint,
		function GSourceFunc,
		data Gpointer,
		notify GDestroyNotify) Guint

	Gdk_threads_add_idle func(
		function GSourceFunc,
		data Gpointer) Guint

	Gdk_threads_add_timeout_full func(
		priority Gint,
		interval Guint,
		function GSourceFunc,
		data Gpointer,
		notify GDestroyNotify) Guint

	Gdk_threads_add_timeout func(
		interval Guint,
		function GSourceFunc,
		data Gpointer) Guint

	Gdk_threads_add_timeout_seconds_full func(
		priority Gint,
		interval Guint,
		function GSourceFunc,
		data Gpointer,
		notify GDestroyNotify) Guint

	Gdk_threads_add_timeout_seconds func(
		interval Guint,
		function GSourceFunc,
		data Gpointer) Guint

	Gdk_pixbuf_non_anim_get_type func() GType

	Gdk_pixbuf_non_anim_new func(
		pixbuf *GdkPixbuf) *GdkPixbufAnimation

	Gdk_pixbuf_animation_new_from_file func(
		filename string,
		e **GError) *GdkPixbufAnimation

	Gdk_pixdata_serialize func(
		pixdata *GdkPixdata,
		stream_length_p *Guint) *Guint8

	Gdk_pixdata_deserialize func(
		pixdata *GdkPixdata,
		stream_length Guint,
		stream *Guint8,
		e **GError) Gboolean

	Gdk_pixdata_from_pixbuf func(
		pixdata *GdkPixdata,
		pixbuf *GdkPixbuf,
		use_rle Gboolean) Gpointer

	Gdk_pixbuf_from_pixdata func(
		pixdata *GdkPixdata,
		copy_pixels Gboolean,
		e **GError) *GdkPixbuf

	Gdk_pixdata_to_csource func(
		pixdata *GdkPixdata,
		name string,
		dump_type GdkPixdataDumpType) *GString

	Gdk_pixbuf_set_option func(
		pixbuf *GdkPixbuf,
		key string,
		value string) Gboolean

	Gdk_window_destroy_notify func(
		window *GdkWindow)

	Gdk_synthesize_window_state func(
		window *GdkWindow,
		unset_flags GdkWindowState,
		set_flags GdkWindowState)

	Gdk_win32_window_is_win32 func(
		window *GdkWindow) Gboolean

	Gdk_win32_window_get_impl_hwnd func(
		window *GdkWindow) HWND

	Gdk_win32_handle_table_lookup func(
		handle GdkNativeWindow) Gpointer

	Gdk_win32_drawable_get_handle func(
		drawable *GdkDrawable) HGDIOBJ

	Gdk_win32_hdc_get func(
		drawable *GdkDrawable,
		gc *GdkGC,
		usage GdkGCValuesMask) HDC

	Gdk_win32_hdc_release func(
		drawable *GdkDrawable,
		gc *GdkGC,
		usage GdkGCValuesMask)

	Gdk_win32_selection_add_targets func(
		owner *GdkWindow,
		selection GdkAtom,
		n_targets Gint,
		targets *GdkAtom)

	Gdk_win32_icon_to_pixbuf_libgtk_only func(
		hicon HICON) *GdkPixbuf

	Gdk_win32_pixbuf_to_hicon_libgtk_only func(
		pixbuf *GdkPixbuf) HICON

	Gdk_win32_set_modal_dialog_libgtk_only func(
		window HWND)

	Gdk_win32_begin_direct_draw_libgtk_only func(
		drawable *GdkDrawable,
		gc *GdkGC,
		priv_data *Gpointer,
		x_offset_out *Gint,
		y_offset_out *Gint) *GdkDrawable

	Gdk_win32_end_direct_draw_libgtk_only func(
		priv_data Gpointer)

	Gdk_win32_window_foreign_new_for_display func(
		display *GdkDisplay,
		anid GdkNativeWindow) *GdkWindow

	Gdk_win32_window_lookup_for_display func(
		display *GdkDisplay,
		anid GdkNativeWindow) *GdkWindow
)

var dll = "libgdk-win32-2.0-0.dll"

var dllPixbuf = "libgdk_pixbuf-2.0-0.dll"

var apiList = Apis{
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

var apiListPixbuf = Apis{
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

type (
	Gdk_pixbuf_major_version Guint
	Gdk_pixbuf_minor_version Guint
	Gdk_pixbuf_micro_version Guint
	Gdk_pixbuf_version       string
	Gdk_threads_mutex        *GMutex
	Gdk_threads_lock         GCallback
	Gdk_threads_unlock       GCallback
)

var DataList = Data{
// {"gdk_threads_lock", new(Gdk_threads_lock)},
// {"gdk_threads_mutex", new(Gdk_threads_mutex)},
// {"gdk_threads_unlock", new(Gdk_threads_unlock)},
// {"gdk_pixbuf_major_version", new(Gdk_pixbuf_major_version)},
// {"gdk_pixbuf_micro_version", new(Gdk_pixbuf_micro_version)},
// {"gdk_pixbuf_minor_version", new(Gdk_pixbuf_minor_version)},
// {"gdk_pixbuf_version", new(Gdk_pixbuf_version)},
}
