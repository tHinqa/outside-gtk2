// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gdk provides API definitions for accessing
//libgdk-win32-2.0-0.dll and libgdk_pixbuf-2.0-0.dll.
package gdk

import (
	"github.com/tHinqa/outside"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
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

	Enum int
)

var (
	InterpTypeGetType func() O.Type

	InputSetExtensionEvents func(
		window *Window, mask int, mode ExtensionMode)

	SetShowEvents func(showEvents bool)

	GetShowEvents func() bool

	AddClientMessageFilter func(
		messageType Atom, f T.GdkFilterFunc, data T.Gpointer)

	SettingGet func(name string, value *T.GValue) bool

	RgbInit func()

	RgbXpixelFromRgb func(rgb T.GUint32) T.Gulong

	RgbGcSetForeground func(gc *GC, rgb T.GUint32)

	RgbGcSetBackground func(gc *GC, rgb T.GUint32)

	RgbFindColor func(colormap *Colormap, color *Color)

	RgbCmapNew func(colors *T.GUint32, nColors int) *T.GdkRgbCmap

	RgbCmapFree func(cmap *T.GdkRgbCmap)

	RgbSetVerbose func(verbose bool)

	RgbSetInstall func(install bool)

	RgbSetMinColors func(minColors int)

	RgbGetColormap func() *Colormap

	RgbGetVisual func() *Visual

	RgbDitherable func() bool

	RgbColormapDitherable func(cmap *Colormap) bool

	FilterReturnGetType func() O.Type

	ScrollDirectionGetType func() O.Type

	NotifyTypeGetType func() O.Type

	SettingActionGetType func() O.Type

	OwnerChangeGetType func() O.Type

	FontTypeGetType func() O.Type

	FillGetType func() O.Type

	FunctionGetType func() O.Type

	JoinStyleGetType func() O.Type

	LineStyleGetType func() O.Type

	SubwindowModeGetType func() O.Type

	InputSourceGetType func() O.Type

	InputModeGetType func() O.Type

	AxisUseGetType func() O.Type

	FillRuleGetType func() O.Type

	OverlapTypeGetType func() O.Type

	RgbDitherGetType func() O.Type

	ByteOrderGetType func() O.Type

	ModifierTypeGetType func() O.Type

	InputConditionGetType func() O.Type

	StatusGetType func() O.Type

	WindowAttributesTypeGetType func() O.Type

	WmDecorationGetType func() O.Type

	WmFunctionGetType func() O.Type

	StringWidth func(font *Font, s string) int

	TextWidth func(font *Font, text string, textLength int) int

	TextWidthWc func(
		font *Font, text *T.GdkWChar, textLength int) int

	CharWidth func(font *Font, character T.Gchar) int

	CharWidthWc func(font *Font, character T.GdkWChar) int

	StringMeasure func(font *Font, s string) int

	TextMeasure func(font *Font, text string, textLength int) int

	CharMeasure func(font *Font, character T.Gchar) int

	StringHeight func(font *Font, s string) int

	TextHeight func(font *Font, text string, textLength int) int

	CharHeight func(font *Font, character T.Gchar) int

	TextExtents func(font *Font, text string, textLength int,
		lbearing, rbearing, width, ascent, descent *int)

	TextExtentsWc func(font *Font, text *T.GdkWChar,
		txtLeng, lbearing, rbearing, width, ascent, descent *int)

	StringExtents func(font *Font, s string,
		lbearing, rbearing, width, ascent, descent *int)

	FontGetDisplay func(font *Font) *Display

	KeyvalName func(keyval uint) string

	KeyvalFromName func(keyvalName string) uint

	KeyvalConvertCase func(symbol uint, lower, upper *uint)

	KeyvalToUpper func(keyval uint) uint

	KeyvalToLower func(keyval uint) uint

	KeyvalIsUpper func(keyval uint) bool

	KeyvalIsLower func(keyval uint) bool

	KeyvalToUnicode func(keyval uint) T.GUint32

	UnicodeToKeyval func(wc T.GUint32) uint

	AtomIntern func(
		atomName string, onlyIfExists bool) Atom

	AtomInternStaticString func(atomName string) Atom

	AtomName func(atom Atom) string

	PropertyGet func(window *Window, property, typ Atom,
		offset, length T.Gulong, pdelete int,
		actualPropertyType *Atom, actualFormat, actualLength *int,
		data **T.Guchar) bool

	PropertyChange func(window *Window, property, typ Atom,
		format int, mode PropMode, data *T.Guchar,
		nelements int)

	PropertyDelete func(window *Window, property Atom)

	TextPropertyToTextList func(encoding Atom, format int,
		text *T.Guchar, length int, list ***T.Gchar) int

	Utf8ToCompoundText func(str string, encoding *Atom,
		format *int, ctext **T.Guchar, length *int) bool

	StringToCompoundText func(str string, encoding *Atom,
		format *int, ctext **T.Guchar, length *int) int

	TextPropertyToUtf8List func(encoding Atom, format int,
		text *T.Guchar, length int, list ***T.Gchar) int

	TextPropertyToUtf8ListForDisplay func(
		display *Display, encoding Atom, format int,
		text *T.Guchar, length int, list ***T.Gchar) int

	Utf8ToStringTarget func(str string) string

	TextPropertyToTextListForDisplay func(
		display *Display, encoding Atom, format int,
		text *T.Guchar, length int, list ***T.Gchar) int

	StringToCompoundTextForDisplay func(
		display *Display, str string, encoding *Atom,
		format *int, ctext **T.Guchar, length *int) int

	Utf8ToCompoundTextForDisplay func(
		display *Display, str string, encoding *Atom,
		format *int, ctext **T.Guchar, length *int) bool

	FreeTextList func(list **T.Gchar)

	FreeCompoundText func(ctext *T.Guchar)

	SelectionOwnerSet func(owner *Window, selection Atom,
		time T.GUint32, sendEvent bool) bool

	SelectionOwnerGet func(selection Atom) *Window

	SelectionOwnerSetForDisplay func(display *Display,
		owner *Window, selection Atom, time T.GUint32,
		sendEvent bool) bool

	SelectionOwnerGetForDisplay func(
		display *Display, selection Atom) *Window

	SelectionConvert func(requestor *Window,
		selection Atom, target Atom, time T.GUint32)

	SelectionPropertyGet func(requestor *Window,
		data **T.Guchar, propType *Atom, propFormat *int) int

	SelectionSendNotify func(requestor T.GdkNativeWindow,
		selection, target, property Atom, time T.GUint32)

	SelectionSendNotifyForDisplay func(
		display *Display, requestor T.GdkNativeWindow,
		selection, target, property Atom, time T.GUint32)

	SpawnOnScreen func(screen *Screen, workingDirectory string,
		argv, envp []string, flags L.SpawnFlags,
		childSetup L.SpawnChildSetupFunc, userData T.Gpointer,
		childPid *int, e **T.GError) bool

	SpawnOnScreenWithPipes func(screen *Screen,
		workingDirectory string, argv, envp []string,
		flags L.SpawnFlags, childSetup L.SpawnChildSetupFunc,
		userData T.Gpointer, childPid, standardInput,
		standardOutput, standardError *int,
		e **T.GError) bool

	SpawnCommandLineOnScreen func(
		screen *Screen, commandLine string, e **T.GError) bool

	WindowForeignNew func(anid T.GdkNativeWindow) *Window

	WindowLookup func(anid T.GdkNativeWindow) *Window

	WindowForeignNewForDisplay func(
		display *Display, anid T.GdkNativeWindow) *Window

	WindowLookupForDisplay func(
		display *Display, anid T.GdkNativeWindow) *Window

	WindowObjectGetType func() O.Type

	SetSmClientId func(smClientId string)

	WindowSetDebugUpdates func(setting bool)

	WindowConstrainSize func(geometry *Geometry, flags uint,
		width, height int, newWidth, newHeight *int)

	GetDefaultRootWindow func() *Window

	OffscreenWindowGetPixmap func(window *Window) *Pixmap

	OffscreenWindowSetEmbedder func(window, embedder *Window)

	OffscreenWindowGetEmbedder func(window *Window) *Window

	SetPointerHooks func(
		newHooks *T.GdkPointerHooks) *T.GdkPointerHooks

	TestRenderSync func(window *Window)

	TestSimulateKey func(window *Window, x, y int,
		keyval uint, modifiers T.GdkModifierType,
		keyPressrelease EventType) bool

	TestSimulateButton func(window *Window, x, y int,
		button uint, modifiers T.GdkModifierType,
		buttonPressrelease EventType) bool

	QueryDepths func(depths **int, count *int)

	QueryVisualTypes func(visualTypes **VisualType, count *int)

	ListVisuals func() *T.GList

	ParseArgs func(argc *int, argv ***T.Gchar)

	Init func(argc *int, argv ***T.Gchar)

	InitCheck func(argc *int, argv ***T.Gchar) bool

	AddOptionEntriesLibgtkOnly func(group *T.GOptionGroup)

	PreParseLibgtkOnly func()

	Exit func(errorCode int)

	SetLocale func() string

	GetProgramClass func() string

	SetProgramClass func(programClass string)

	ErrorTrapPush func()

	ErrorTrapPop func() int

	SetUseXshm func(useXshm bool)

	GetUseXshm func() bool

	GetDisplay func() string

	GetDisplayArgName func() string

	InputAddFull func(source int, condition T.GdkInputCondition,
		function T.GdkInputFunction, data T.Gpointer,
		destroy T.GDestroyNotify) int

	InputAdd func(source int, condition T.GdkInputCondition,
		function T.GdkInputFunction, data T.Gpointer) int

	InputRemove func(tag int)

	PointerGrab func(window *Window, ownerEvents bool,
		eventMask EventMask, confineTo *Window,
		cursor *Cursor, time T.GUint32) GrabStatus

	KeyboardGrab func(window *Window, ownerEvents bool,
		time T.GUint32) GrabStatus

	PointerGrabInfoLibgtkOnly func(display *Display,
		grabWindow **Window, ownerEvents *bool) bool

	KeyboardGrabInfoLibgtkOnly func(display *Display,
		grabWindow **Window, ownerEvents *bool) bool

	PointerUngrab func(time T.GUint32)

	KeyboardUngrab func(time T.GUint32)

	PointerIsGrabbed func() bool

	ScreenWidth func() int

	ScreenHeight func() int

	ScreenWidthMm func() int

	ScreenHeightMm func() int

	Beep func()

	Flush func()

	SetDoubleClickTime func(msec uint)

	Wcstombs func(src *T.GdkWChar) string

	Mbstowcs func(dest *T.GdkWChar, src string, destMax int) int

	NotifyStartupComplete func()

	NotifyStartupCompleteWithId func(startupId string)

	ThreadsEnter func()

	ThreadsLeave func()

	ThreadsInit func()

	ThreadsSetLockFunctions func(enterFn, leaveFn T.GCallback)

	ThreadsAddIdleFull func(priority int, function O.SourceFunc,
		data T.Gpointer, notify T.GDestroyNotify) uint

	ThreadsAddIdle func(
		function O.SourceFunc, data T.Gpointer) uint

	ThreadsAddTimeoutFull func(
		priority int, interval uint, function O.SourceFunc,
		data T.Gpointer, notify T.GDestroyNotify) uint

	ThreadsAddTimeout func(interval uint,
		function O.SourceFunc, data T.Gpointer) uint

	ThreadsAddTimeoutSecondsFull func(
		priority int, interval uint, function O.SourceFunc,
		data T.Gpointer, notify T.GDestroyNotify) uint

	ThreadsAddTimeoutSeconds func(interval uint,
		function O.SourceFunc, data T.Gpointer) uint

	SynthesizeWindowState func(window *Window,
		unsetFlags WindowState, setFlags WindowState)

	Win32WindowIsWin32 func(window *Window) bool

	Win32WindowGetImplHwnd func(window *Window) HWND

	Win32HandleTableLookup func(handle T.GdkNativeWindow) T.Gpointer

	Win32DrawableGetHandle func(drawable *Drawable) HGDIOBJ

	Win32HdcGet func(drawable *Drawable, gc *GC, usage GCValuesMask) HDC

	Win32HdcRelease func(drawable *Drawable, gc *GC, usage GCValuesMask)

	Win32SelectionAddTargets func(owner *Window, selection Atom, nTargets int, targets *Atom)

	Win32IconToPixbufLibgtkOnly func(hicon HICON) *Pixbuf

	Win32PixbufToHiconLibgtkOnly func(pixbuf *Pixbuf) HICON

	Win32SetModalDialogLibgtkOnly func(window HWND)

	Win32BeginDirectDrawLibgtkOnly func(drawable *Drawable, gc *GC, privData *T.Gpointer, xOffsetOut *int, yOffsetOut *int) *Drawable

	Win32EndDirectDrawLibgtkOnly func(privData T.Gpointer)

	Win32WindowForeignNewForDisplay func(display *Display, anid T.GdkNativeWindow) *Window

	Win32WindowLookupForDisplay func(display *Display, anid T.GdkNativeWindow) *Window

	BitmapCreateFromData func(drawable *Drawable, data string, width, height int) *T.GdkBitmap
)

var dll = "libgdk-win32-2.0-0.dll"

var dllPixbuf = "libgdk_pixbuf-2.0-0.dll"

var apiList = outside.Apis{
	{"gdk_add_client_message_filter", &AddClientMessageFilter},
	{"gdk_add_option_entries_libgtk_only", &AddOptionEntriesLibgtkOnly},
	{"gdk_app_launch_context_get_type", &AppLaunchContextGetType},
	{"gdk_app_launch_context_new", &AppLaunchContextNew},
	{"gdk_app_launch_context_set_desktop", &AppLaunchContextSetDesktop},
	{"gdk_app_launch_context_set_display", &AppLaunchContextSetDisplay},
	{"gdk_app_launch_context_set_icon", &AppLaunchContextSetIcon},
	{"gdk_app_launch_context_set_icon_name", &AppLaunchContextSetIconName},
	{"gdk_app_launch_context_set_screen", &AppLaunchContextSetScreen},
	{"gdk_app_launch_context_set_timestamp", &AppLaunchContextSetTimestamp},
	{"gdk_atom_intern", &AtomIntern},
	{"gdk_atom_intern_static_string", &AtomInternStaticString},
	{"gdk_atom_name", &AtomName},
	{"gdk_axis_use_get_type", &AxisUseGetType},
	{"gdk_beep", &Beep},
	{"gdk_bitmap_create_from_data", &BitmapCreateFromData},
	{"gdk_byte_order_get_type", &ByteOrderGetType},
	{"gdk_cairo_create", &CairoCreate},
	{"gdk_cairo_rectangle", &CairoRectangle},
	{"gdk_cairo_region", &CairoRegion},
	{"gdk_cairo_reset_clip", &CairoResetClip},
	{"gdk_cairo_set_source_color", &CairoSetSourceColor},
	{"gdk_cairo_set_source_pixbuf", &CairoSetSourcePixbuf},
	{"gdk_cairo_set_source_pixmap", &CairoSetSourcePixmap},
	{"gdk_cairo_set_source_window", &CairoSetSourceWindow},
	{"gdk_cap_style_get_type", &CapStyleGetType},
	{"gdk_char_height", &CharHeight},
	{"gdk_char_measure", &CharMeasure},
	{"gdk_char_width", &CharWidth},
	{"gdk_char_width_wc", &CharWidthWc},
	{"gdk_color_alloc", &ColorAlloc},
	{"gdk_color_black", &ColorBlack},
	{"gdk_color_change", &ColorChange},
	{"gdk_color_copy", &ColorCopy},
	{"gdk_color_equal", &ColorEqual},
	{"gdk_color_free", &ColorFree},
	{"gdk_color_get_type", &ColorGetType},
	{"gdk_color_hash", &ColorHash},
	{"gdk_color_parse", &ColorParse},
	{"gdk_color_to_string", &ColorToString},
	{"gdk_color_white", &ColorWhite},
	{"gdk_colormap_alloc_color", &ColormapAllocColor},
	{"gdk_colormap_alloc_colors", &ColormapAllocColors},
	{"gdk_colormap_change", &ColormapChange},
	{"gdk_colormap_free_colors", &ColormapFreeColors},
	{"gdk_colormap_get_screen", &ColormapGetScreen},
	{"gdk_colormap_get_system", &ColormapGetSystem},
	{"gdk_colormap_get_system_size", &ColormapGetSystemSize},
	{"gdk_colormap_get_type", &ColormapGetType},
	{"gdk_colormap_get_visual", &ColormapGetVisual},
	{"gdk_colormap_new", &ColormapNew},
	{"gdk_colormap_query_color", &ColormapQueryColor},
	{"gdk_colormap_ref", &ColormapRef},
	{"gdk_colormap_unref", &ColormapUnref},
	{"gdk_colors_alloc", &ColorsAlloc},
	{"gdk_colors_free", &ColorsFree},
	{"gdk_colors_store", &ColorsStore},
	{"gdk_crossing_mode_get_type", &CrossingModeGetType},
	{"gdk_cursor_get_cursor_type", &CursorGetCursorType},
	{"gdk_cursor_get_display", &CursorGetDisplay},
	{"gdk_cursor_get_image", &CursorGetImage},
	{"gdk_cursor_get_type", &CursorGetType},
	{"gdk_cursor_new", &CursorNew},
	{"gdk_cursor_new_for_display", &CursorNewForDisplay},
	{"gdk_cursor_new_from_name", &CursorNewFromName},
	{"gdk_cursor_new_from_pixbuf", &CursorNewFromPixbuf},
	{"gdk_cursor_new_from_pixmap", &CursorNewFromPixmap},
	{"gdk_cursor_ref", &CursorRef},
	{"gdk_cursor_type_get_type", &CursorTypeGetType},
	{"gdk_cursor_unref", &CursorUnref},
	{"gdk_device_free_history", &DeviceFreeHistory},
	{"gdk_device_get_axis", &DeviceGetAxis},
	{"gdk_device_get_axis_use", &DeviceGetAxisUse},
	{"gdk_device_get_core_pointer", &DeviceGetCorePointer},
	{"gdk_device_get_has_cursor", &DeviceGetHasCursor},
	{"gdk_device_get_history", &DeviceGetHistory},
	{"gdk_device_get_key", &DeviceGetKey},
	{"gdk_device_get_mode", &DeviceGetMode},
	{"gdk_device_get_n_axes", &DeviceGetNAxes},
	{"gdk_device_get_n_keys", &DeviceGetNKeys},
	{"gdk_device_get_name", &DeviceGetName},
	{"gdk_device_get_source", &DeviceGetSource},
	{"gdk_device_get_state", &DeviceGetState},
	{"gdk_device_get_type", &DeviceGetType},
	{"gdk_device_set_axis_use", &DeviceSetAxisUse},
	{"gdk_device_set_key", &DeviceSetKey},
	{"gdk_device_set_mode", &DeviceSetMode},
	{"gdk_device_set_source", &DeviceSetSource},
	{"gdk_devices_list", &DevicesList},
	{"gdk_display_add_client_message_filter", &DisplayAddClientMessageFilter},
	{"gdk_display_beep", &DisplayBeep},
	{"gdk_display_close", &DisplayClose},
	{"gdk_display_flush", &DisplayFlush},
	{"gdk_display_get_core_pointer", &DisplayGetCorePointer},
	{"gdk_display_get_default", &DisplayGetDefault},
	{"gdk_display_get_default_cursor_size", &DisplayGetDefaultCursorSize},
	{"gdk_display_get_default_group", &DisplayGetDefaultGroup},
	{"gdk_display_get_default_screen", &DisplayGetDefaultScreen},
	{"gdk_display_get_event", &DisplayGetEvent},
	{"gdk_display_get_maximal_cursor_size", &DisplayGetMaximalCursorSize},
	{"gdk_display_get_n_screens", &DisplayGetNScreens},
	{"gdk_display_get_name", &DisplayGetName},
	{"gdk_display_get_pointer", &DisplayGetPointer},
	{"gdk_display_get_screen", &DisplayGetScreen},
	{"gdk_display_get_type", &DisplayGetType},
	{"gdk_display_get_window_at_pointer", &DisplayGetWindowAtPointer},
	{"gdk_display_is_closed", &DisplayIsClosed},
	{"gdk_display_keyboard_ungrab", &DisplayKeyboardUngrab},
	{"gdk_display_list_devices", &DisplayListDevices},
	{"gdk_display_manager_get", &DisplayManagerGet},
	{"gdk_display_manager_get_default_display", &DisplayManagerGetDefaultDisplay},
	{"gdk_display_manager_get_type", &DisplayManagerGetType},
	{"gdk_display_manager_list_displays", &DisplayManagerListDisplays},
	{"gdk_display_manager_set_default_display", &DisplayManagerSetDefaultDisplay},
	{"gdk_display_open", &DisplayOpen},
	{"gdk_display_open_default_libgtk_only", &DisplayOpenDefaultLibgtkOnly},
	{"gdk_display_peek_event", &DisplayPeekEvent},
	{"gdk_display_pointer_is_grabbed", &DisplayPointerIsGrabbed},
	{"gdk_display_pointer_ungrab", &DisplayPointerUngrab},
	{"gdk_display_put_event", &DisplayPutEvent},
	{"gdk_display_request_selection_notification", &DisplayRequestSelectionNotification},
	{"gdk_display_set_double_click_distance", &DisplaySetDoubleClickDistance},
	{"gdk_display_set_double_click_time", &DisplaySetDoubleClickTime},
	{"gdk_display_set_pointer_hooks", &DisplaySetPointerHooks},
	{"gdk_display_store_clipboard", &DisplayStoreClipboard},
	{"gdk_display_supports_clipboard_persistence", &DisplaySupportsClipboardPersistence},
	{"gdk_display_supports_composite", &DisplaySupportsComposite},
	{"gdk_display_supports_cursor_alpha", &DisplaySupportsCursorAlpha},
	{"gdk_display_supports_cursor_color", &DisplaySupportsCursorColor},
	{"gdk_display_supports_input_shapes", &DisplaySupportsInputShapes},
	{"gdk_display_supports_selection_notification", &DisplaySupportsSelectionNotification},
	{"gdk_display_supports_shapes", &DisplaySupportsShapes},
	{"gdk_display_sync", &DisplaySync},
	{"gdk_display_warp_pointer", &DisplayWarpPointer},
	{"gdk_drag_abort", &DragAbort},
	{"gdk_drag_action_get_type", &DragActionGetType},
	{"gdk_drag_begin", &DragBegin},
	{"gdk_drag_context_get_actions", &DragContextGetActions},
	{"gdk_drag_context_get_dest_window", &DragContextGetDestWindow},
	{"gdk_drag_context_get_protocol", &DragContextGetProtocol},
	{"gdk_drag_context_get_selected_action", &DragContextGetSelectedAction},
	{"gdk_drag_context_get_source_window", &DragContextGetSourceWindow},
	{"gdk_drag_context_get_suggested_action", &DragContextGetSuggestedAction},
	{"gdk_drag_context_get_type", &DragContextGetType},
	{"gdk_drag_context_list_targets", &DragContextListTargets},
	{"gdk_drag_context_new", &DragContextNew},
	{"gdk_drag_context_ref", &DragContextRef},
	{"gdk_drag_context_unref", &DragContextUnref},
	{"gdk_drag_drop", &DragDrop},
	{"gdk_drag_drop_succeeded", &DragDropSucceeded},
	{"gdk_drag_find_window", &DragFindWindow},
	{"gdk_drag_find_window_for_screen", &DragFindWindowForScreen},
	{"gdk_drag_get_protocol", &DragGetProtocol},
	{"gdk_drag_get_protocol_for_display", &DragGetProtocolForDisplay},
	{"gdk_drag_get_selection", &DragGetSelection},
	{"gdk_drag_motion", &DragMotion},
	{"gdk_drag_protocol_get_type", &DragProtocolGetType},
	{"gdk_drag_status", &DragStatus},
	{"gdk_draw_arc", &DrawArc},
	{"gdk_draw_drawable", &DrawDrawable},
	{"gdk_draw_glyphs", &DrawGlyphs},
	{"gdk_draw_glyphs_transformed", &DrawGlyphsTransformed},
	{"gdk_draw_gray_image", &DrawGrayImage},
	{"gdk_draw_image", &DrawImage},
	{"gdk_draw_indexed_image", &DrawIndexedImage},
	{"gdk_draw_layout", &DrawLayout},
	{"gdk_draw_layout_line", &DrawLayoutLine},
	{"gdk_draw_layout_line_with_colors", &DrawLayoutLineWithColors},
	{"gdk_draw_layout_with_colors", &DrawLayoutWithColors},
	{"gdk_draw_line", &DrawLine},
	{"gdk_draw_lines", &DrawLines},
	{"gdk_draw_pixbuf", &DrawPixbuf},
	{"gdk_draw_point", &DrawPoint},
	{"gdk_draw_points", &DrawPoints},
	{"gdk_draw_polygon", &DrawPolygon},
	{"gdk_draw_rectangle", &DrawRectangle},
	{"gdk_draw_rgb_32_image", &DrawRgb32Image},
	{"gdk_draw_rgb_32_image_dithalign", &DrawRgb32ImageDithalign},
	{"gdk_draw_rgb_image", &DrawRgbImage},
	{"gdk_draw_rgb_image_dithalign", &DrawRgbImageDithalign},
	{"gdk_draw_segments", &DrawSegments},
	{"gdk_draw_string", &DrawString},
	{"gdk_draw_text", &DrawText},
	{"gdk_draw_text_wc", &DrawTextWc},
	{"gdk_draw_trapezoids", &DrawTrapezoids},
	{"gdk_drawable_copy_to_image", &DrawableCopyToImage},
	{"gdk_drawable_get_clip_region", &DrawableGetClipRegion},
	{"gdk_drawable_get_colormap", &DrawableGetColormap},
	{"gdk_drawable_get_data", &DrawableGetData},
	{"gdk_drawable_get_depth", &DrawableGetDepth},
	{"gdk_drawable_get_display", &DrawableGetDisplay},
	{"gdk_drawable_get_image", &DrawableGetImage},
	{"gdk_drawable_get_screen", &DrawableGetScreen},
	{"gdk_drawable_get_size", &DrawableGetSize},
	{"gdk_drawable_get_type", &DrawableGetType},
	{"gdk_drawable_get_visible_region", &DrawableGetVisibleRegion},
	{"gdk_drawable_get_visual", &DrawableGetVisual},
	{"gdk_drawable_ref", &DrawableRef},
	{"gdk_drawable_set_colormap", &DrawableSetColormap},
	{"gdk_drawable_set_data", &DrawableSetData},
	{"gdk_drawable_unref", &DrawableUnref},
	{"gdk_drop_finish", &DropFinish},
	{"gdk_drop_reply", &DropReply},
	{"gdk_error_trap_pop", &ErrorTrapPop},
	{"gdk_error_trap_push", &ErrorTrapPush},
	{"gdk_event_copy", &EventCopy},
	{"gdk_event_free", &EventFree},
	{"gdk_event_get", &EventGet},
	{"gdk_event_get_axis", &EventGetAxis},
	{"gdk_event_get_coords", &EventGetCoords},
	{"gdk_event_get_graphics_expose", &EventGetGraphicsExpose},
	{"gdk_event_get_root_coords", &EventGetRootCoords},
	{"gdk_event_get_screen", &EventGetScreen},
	{"gdk_event_get_state", &EventGetState},
	{"gdk_event_get_time", &EventGetTime},
	{"gdk_event_get_type", &EventGetType},
	{"gdk_event_handler_set", &EventHandlerSet},
	{"gdk_event_mask_get_type", &EventMaskGetType},
	{"gdk_event_new", &EventNew},
	{"gdk_event_peek", &EventPeek},
	{"gdk_event_put", &EventPut},
	{"gdk_event_request_motions", &EventRequestMotions},
	{"gdk_event_send_client_message", &EventSendClientMessage},
	{"gdk_event_send_client_message_for_display", &EventSendClientMessageForDisplay},
	{"gdk_event_send_clientmessage_toall", &EventSendClientmessageToall},
	{"gdk_event_set_screen", &EventSetScreen},
	{"gdk_event_type_get_type", &EventTypeGetType},
	{"gdk_events_pending", &EventsPending},
	{"gdk_exit", &Exit},
	{"gdk_extension_mode_get_type", &ExtensionModeGetType},
	{"gdk_fill_get_type", &FillGetType},
	{"gdk_fill_rule_get_type", &FillRuleGetType},
	{"gdk_filter_return_get_type", &FilterReturnGetType},
	{"gdk_flush", &Flush},
	{"gdk_font_equal", &FontEqual},
	{"gdk_font_from_description", &FontFromDescription},
	{"gdk_font_from_description_for_display", &FontFromDescriptionForDisplay},
	{"gdk_font_get_display", &FontGetDisplay},
	{"gdk_font_get_type", &FontGetType},
	{"gdk_font_id", &FontId},
	{"gdk_font_load", &FontLoad},
	{"gdk_font_load_for_display", &FontLoadForDisplay},
	{"gdk_font_ref", &FontRef},
	{"gdk_font_type_get_type", &FontTypeGetType},
	{"gdk_font_unref", &FontUnref},
	{"gdk_fontset_load", &FontsetLoad},
	{"gdk_fontset_load_for_display", &FontsetLoadForDisplay},
	{"gdk_free_compound_text", &FreeCompoundText},
	{"gdk_free_text_list", &FreeTextList},
	{"gdk_function_get_type", &FunctionGetType},
	{"gdk_gc_copy", &GcCopy},
	{"gdk_gc_get_colormap", &GcGetColormap},
	{"gdk_gc_get_screen", &GcGetScreen},
	{"gdk_gc_get_type", &GcGetType},
	{"gdk_gc_get_values", &GcGetValues},
	{"gdk_gc_new", &GcNew},
	{"gdk_gc_new_with_values", &GcNewWithValues},
	{"gdk_gc_offset", &GcOffset},
	{"gdk_gc_ref", &GcRef},
	{"gdk_gc_set_background", &GcSetBackground},
	{"gdk_gc_set_clip_mask", &GcSetClipMask},
	{"gdk_gc_set_clip_origin", &GcSetClipOrigin},
	{"gdk_gc_set_clip_rectangle", &GcSetClipRectangle},
	{"gdk_gc_set_clip_region", &GcSetClipRegion},
	{"gdk_gc_set_colormap", &GcSetColormap},
	{"gdk_gc_set_dashes", &GcSetDashes},
	{"gdk_gc_set_exposures", &GcSetExposures},
	{"gdk_gc_set_fill", &GcSetFill},
	{"gdk_gc_set_font", &GcSetFont},
	{"gdk_gc_set_foreground", &GcSetForeground},
	{"gdk_gc_set_function", &GcSetFunction},
	{"gdk_gc_set_line_attributes", &GcSetLineAttributes},
	{"gdk_gc_set_rgb_bg_color", &GcSetRgbBgColor},
	{"gdk_gc_set_rgb_fg_color", &GcSetRgbFgColor},
	{"gdk_gc_set_stipple", &GcSetStipple},
	{"gdk_gc_set_subwindow", &GcSetSubwindow},
	{"gdk_gc_set_tile", &GcSetTile},
	{"gdk_gc_set_ts_origin", &GcSetTsOrigin},
	{"gdk_gc_set_values", &GcSetValues},
	{"gdk_gc_unref", &GcUnref},
	{"gdk_gc_values_mask_get_type", &GcValuesMaskGetType},
	{"gdk_get_default_root_window", &GetDefaultRootWindow},
	{"gdk_get_display", &GetDisplay},
	{"gdk_get_display_arg_name", &GetDisplayArgName},
	{"gdk_get_program_class", &GetProgramClass},
	{"gdk_get_show_events", &GetShowEvents},
	{"gdk_get_use_xshm", &GetUseXshm},
	{"gdk_grab_status_get_type", &GrabStatusGetType},
	{"gdk_gravity_get_type", &GravityGetType},
	{"gdk_image_get", &ImageGet},
	{"gdk_image_get_bits_per_pixel", &ImageGetBitsPerPixel},
	{"gdk_image_get_byte_order", &ImageGetByteOrder},
	{"gdk_image_get_bytes_per_line", &ImageGetBytesPerLine},
	{"gdk_image_get_bytes_per_pixel", &ImageGetBytesPerPixel},
	{"gdk_image_get_colormap", &ImageGetColormap},
	{"gdk_image_get_depth", &ImageGetDepth},
	{"gdk_image_get_height", &ImageGetHeight},
	{"gdk_image_get_image_type", &ImageGetImageType},
	{"gdk_image_get_pixel", &ImageGetPixel},
	{"gdk_image_get_pixels", &ImageGetPixels},
	{"gdk_image_get_type", &ImageGetType},
	{"gdk_image_get_visual", &ImageGetVisual},
	{"gdk_image_get_width", &ImageGetWidth},
	{"gdk_image_new", &ImageNew},
	{"gdk_image_put_pixel", &ImagePutPixel},
	{"gdk_image_ref", &ImageRef},
	{"gdk_image_set_colormap", &ImageSetColormap},
	{"gdk_image_type_get_type", &ImageTypeGetType},
	{"gdk_image_unref", &ImageUnref},
	{"gdk_init", &Init},
	{"gdk_init_check", &InitCheck},
	{"gdk_input_add", &InputAdd},
	{"gdk_input_add_full", &InputAddFull},
	{"gdk_input_condition_get_type", &InputConditionGetType},
	{"gdk_input_mode_get_type", &InputModeGetType},
	{"gdk_input_remove", &InputRemove},
	{"gdk_input_set_extension_events", &InputSetExtensionEvents},
	{"gdk_input_source_get_type", &InputSourceGetType},
	{"gdk_join_style_get_type", &JoinStyleGetType},
	{"gdk_keyboard_grab", &KeyboardGrab},
	{"gdk_keyboard_grab_info_libgtk_only", &KeyboardGrabInfoLibgtkOnly},
	{"gdk_keyboard_ungrab", &KeyboardUngrab},
	{"gdk_keymap_add_virtual_modifiers", &KeymapAddVirtualModifiers},
	{"gdk_keymap_get_caps_lock_state", &KeymapGetCapsLockState},
	{"gdk_keymap_get_default", &KeymapGetDefault},
	{"gdk_keymap_get_direction", &KeymapGetDirection},
	{"gdk_keymap_get_entries_for_keycode", &KeymapGetEntriesForKeycode},
	{"gdk_keymap_get_entries_for_keyval", &KeymapGetEntriesForKeyval},
	{"gdk_keymap_get_for_display", &KeymapGetForDisplay},
	{"gdk_keymap_get_type", &KeymapGetType},
	{"gdk_keymap_have_bidi_layouts", &KeymapHaveBidiLayouts},
	{"gdk_keymap_lookup_key", &KeymapLookupKey},
	{"gdk_keymap_map_virtual_modifiers", &KeymapMapVirtualModifiers},
	{"gdk_keymap_translate_keyboard_state", &KeymapTranslateKeyboardState},
	{"gdk_keyval_convert_case", &KeyvalConvertCase},
	{"gdk_keyval_from_name", &KeyvalFromName},
	{"gdk_keyval_is_lower", &KeyvalIsLower},
	{"gdk_keyval_is_upper", &KeyvalIsUpper},
	{"gdk_keyval_name", &KeyvalName},
	{"gdk_keyval_to_lower", &KeyvalToLower},
	{"gdk_keyval_to_unicode", &KeyvalToUnicode},
	{"gdk_keyval_to_upper", &KeyvalToUpper},
	{"gdk_line_style_get_type", &LineStyleGetType},
	{"gdk_list_visuals", &ListVisuals},
	{"gdk_mbstowcs", &Mbstowcs},
	{"gdk_modifier_type_get_type", &ModifierTypeGetType},
	// Undocumented {"gdk_net_wm_supports", &NetWmSupports},
	{"gdk_notify_startup_complete", &NotifyStartupComplete},
	{"gdk_notify_startup_complete_with_id", &NotifyStartupCompleteWithId},
	{"gdk_notify_type_get_type", &NotifyTypeGetType},
	{"gdk_offscreen_window_get_embedder", &OffscreenWindowGetEmbedder},
	{"gdk_offscreen_window_get_pixmap", &OffscreenWindowGetPixmap},
	// Undocumented {"gdk_offscreen_window_get_type", &OffscreenWindowGetType},
	{"gdk_offscreen_window_set_embedder", &OffscreenWindowSetEmbedder},
	{"gdk_overlap_type_get_type", &OverlapTypeGetType},
	{"gdk_owner_change_get_type", &OwnerChangeGetType},
	{"gdk_pango_attr_emboss_color_new", &PangoAttrEmbossColorNew},
	{"gdk_pango_attr_embossed_new", &PangoAttrEmbossedNew},
	{"gdk_pango_attr_stipple_new", &PangoAttrStippleNew},
	{"gdk_pango_context_get", &PangoContextGet},
	{"gdk_pango_context_get_for_screen", &PangoContextGetForScreen},
	{"gdk_pango_context_set_colormap", &PangoContextSetColormap},
	{"gdk_pango_layout_get_clip_region", &PangoLayoutGetClipRegion},
	{"gdk_pango_layout_line_get_clip_region", &PangoLayoutLineGetClipRegion},
	{"gdk_pango_renderer_get_default", &PangoRendererGetDefault},
	{"gdk_pango_renderer_get_type", &PangoRendererGetType},
	{"gdk_pango_renderer_new", &PangoRendererNew},
	{"gdk_pango_renderer_set_drawable", &PangoRendererSetDrawable},
	{"gdk_pango_renderer_set_gc", &PangoRendererSetGc},
	{"gdk_pango_renderer_set_override_color", &PangoRendererSetOverrideColor},
	{"gdk_pango_renderer_set_stipple", &PangoRendererSetStipple},
	{"gdk_parse_args", &ParseArgs},
	{"gdk_pixbuf_get_from_drawable", &PixbufGetFromDrawable},
	{"gdk_pixbuf_get_from_image", &PixbufGetFromImage},
	{"gdk_pixbuf_render_pixmap_and_mask", &PixbufRenderPixmapAndMask},
	{"gdk_pixbuf_render_pixmap_and_mask_for_colormap", &PixbufRenderPixmapAndMaskForColormap},
	{"gdk_pixbuf_render_threshold_alpha", &PixbufRenderThresholdAlpha},
	{"gdk_pixbuf_render_to_drawable", &PixbufRenderToDrawable},
	{"gdk_pixbuf_render_to_drawable_alpha", &PixbufRenderToDrawableAlpha},
	{"gdk_pixmap_colormap_create_from_xpm", &PixmapColormapCreateFromXpm},
	{"gdk_pixmap_colormap_create_from_xpm_d", &PixmapColormapCreateFromXpmD},
	{"gdk_pixmap_create_from_data", &PixmapCreateFromData},
	{"gdk_pixmap_create_from_xpm", &PixmapCreateFromXpm},
	{"gdk_pixmap_create_from_xpm_d", &PixmapCreateFromXpmD},
	{"gdk_pixmap_foreign_new", &PixmapForeignNew},
	{"gdk_pixmap_foreign_new_for_display", &PixmapForeignNewForDisplay},
	{"gdk_pixmap_foreign_new_for_screen", &PixmapForeignNewForScreen},
	{"gdk_pixmap_get_size", &PixmapGetSize},
	{"gdk_pixmap_get_type", &PixmapGetType},
	{"gdk_pixmap_lookup", &PixmapLookup},
	{"gdk_pixmap_lookup_for_display", &PixmapLookupForDisplay},
	{"gdk_pixmap_new", &PixmapNew},
	{"gdk_pointer_grab", &PointerGrab},
	{"gdk_pointer_grab_info_libgtk_only", &PointerGrabInfoLibgtkOnly},
	{"gdk_pointer_is_grabbed", &PointerIsGrabbed},
	{"gdk_pointer_ungrab", &PointerUngrab},
	{"gdk_pre_parse_libgtk_only", &PreParseLibgtkOnly},
	{"gdk_prop_mode_get_type", &PropModeGetType},
	{"gdk_property_change", &PropertyChange},
	{"gdk_property_delete", &PropertyDelete},
	{"gdk_property_get", &PropertyGet},
	{"gdk_property_state_get_type", &PropertyStateGetType},
	{"gdk_query_depths", &QueryDepths},
	{"gdk_query_visual_types", &QueryVisualTypes},
	{"gdk_rectangle_get_type", &RectangleGetType},
	{"gdk_rectangle_intersect", &RectangleIntersect},
	{"gdk_rectangle_union", &RectangleUnion},
	{"gdk_region_copy", &RegionCopy},
	{"gdk_region_destroy", &RegionDestroy},
	{"gdk_region_empty", &RegionEmpty},
	{"gdk_region_equal", &RegionEqual},
	{"gdk_region_get_clipbox", &RegionGetClipbox},
	{"gdk_region_get_rectangles", &RegionGetRectangles},
	{"gdk_region_intersect", &RegionIntersect},
	{"gdk_region_new", &RegionNew},
	{"gdk_region_offset", &RegionOffset},
	{"gdk_region_point_in", &RegionPointIn},
	{"gdk_region_polygon", &RegionPolygon},
	{"gdk_region_rect_equal", &RegionRectEqual},
	{"gdk_region_rect_in", &RegionRectIn},
	{"gdk_region_rectangle", &RegionRectangle},
	{"gdk_region_shrink", &RegionShrink},
	{"gdk_region_spans_intersect_foreach", &RegionSpansIntersectForeach},
	{"gdk_region_subtract", &RegionSubtract},
	{"gdk_region_union", &RegionUnion},
	{"gdk_region_union_with_rect", &RegionUnionWithRect},
	{"gdk_region_xor", &RegionXor},
	{"gdk_rgb_cmap_free", &RgbCmapFree},
	{"gdk_rgb_cmap_new", &RgbCmapNew},
	{"gdk_rgb_colormap_ditherable", &RgbColormapDitherable},
	{"gdk_rgb_dither_get_type", &RgbDitherGetType},
	{"gdk_rgb_ditherable", &RgbDitherable},
	{"gdk_rgb_find_color", &RgbFindColor},
	{"gdk_rgb_gc_set_background", &RgbGcSetBackground},
	{"gdk_rgb_gc_set_foreground", &RgbGcSetForeground},
	{"gdk_rgb_get_colormap", &RgbGetColormap},
	{"gdk_rgb_get_visual", &RgbGetVisual},
	{"gdk_rgb_init", &RgbInit},
	{"gdk_rgb_set_install", &RgbSetInstall},
	{"gdk_rgb_set_min_colors", &RgbSetMinColors},
	{"gdk_rgb_set_verbose", &RgbSetVerbose},
	{"gdk_rgb_xpixel_from_rgb", &RgbXpixelFromRgb},
	{"gdk_screen_broadcast_client_message", &ScreenBroadcastClientMessage},
	{"gdk_screen_get_active_window", &ScreenGetActiveWindow},
	{"gdk_screen_get_default", &ScreenGetDefault},
	{"gdk_screen_get_default_colormap", &ScreenGetDefaultColormap},
	{"gdk_screen_get_display", &ScreenGetDisplay},
	{"gdk_screen_get_font_options", &ScreenGetFontOptions},
	{"gdk_screen_get_height", &ScreenGetHeight},
	{"gdk_screen_get_height_mm", &ScreenGetHeightMm},
	{"gdk_screen_get_monitor_at_point", &ScreenGetMonitorAtPoint},
	{"gdk_screen_get_monitor_at_window", &ScreenGetMonitorAtWindow},
	{"gdk_screen_get_monitor_geometry", &ScreenGetMonitorGeometry},
	{"gdk_screen_get_monitor_height_mm", &ScreenGetMonitorHeightMm},
	{"gdk_screen_get_monitor_plug_name", &ScreenGetMonitorPlugName},
	{"gdk_screen_get_monitor_width_mm", &ScreenGetMonitorWidthMm},
	{"gdk_screen_get_n_monitors", &ScreenGetNMonitors},
	{"gdk_screen_get_number", &ScreenGetNumber},
	{"gdk_screen_get_primary_monitor", &ScreenGetPrimaryMonitor},
	{"gdk_screen_get_resolution", &ScreenGetResolution},
	{"gdk_screen_get_rgb_colormap", &ScreenGetRgbColormap},
	{"gdk_screen_get_rgb_visual", &ScreenGetRgbVisual},
	{"gdk_screen_get_rgba_colormap", &ScreenGetRgbaColormap},
	{"gdk_screen_get_rgba_visual", &ScreenGetRgbaVisual},
	{"gdk_screen_get_root_window", &ScreenGetRootWindow},
	{"gdk_screen_get_setting", &ScreenGetSetting},
	{"gdk_screen_get_system_colormap", &ScreenGetSystemColormap},
	{"gdk_screen_get_system_visual", &ScreenGetSystemVisual},
	{"gdk_screen_get_toplevel_windows", &ScreenGetToplevelWindows},
	{"gdk_screen_get_type", &ScreenGetType},
	{"gdk_screen_get_width", &ScreenGetWidth},
	{"gdk_screen_get_width_mm", &ScreenGetWidthMm},
	{"gdk_screen_get_window_stack", &ScreenGetWindowStack},
	{"gdk_screen_height", &ScreenHeight},
	{"gdk_screen_height_mm", &ScreenHeightMm},
	{"gdk_screen_is_composited", &ScreenIsComposited},
	{"gdk_screen_list_visuals", &ScreenListVisuals},
	{"gdk_screen_make_display_name", &ScreenMakeDisplayName},
	{"gdk_screen_set_default_colormap", &ScreenSetDefaultColormap},
	{"gdk_screen_set_font_options", &ScreenSetFontOptions},
	{"gdk_screen_set_resolution", &ScreenSetResolution},
	{"gdk_screen_width", &ScreenWidth},
	{"gdk_screen_width_mm", &ScreenWidthMm},
	{"gdk_scroll_direction_get_type", &ScrollDirectionGetType},
	{"gdk_selection_convert", &SelectionConvert},
	{"gdk_selection_owner_get", &SelectionOwnerGet},
	{"gdk_selection_owner_get_for_display", &SelectionOwnerGetForDisplay},
	{"gdk_selection_owner_set", &SelectionOwnerSet},
	{"gdk_selection_owner_set_for_display", &SelectionOwnerSetForDisplay},
	{"gdk_selection_property_get", &SelectionPropertyGet},
	{"gdk_selection_send_notify", &SelectionSendNotify},
	{"gdk_selection_send_notify_for_display", &SelectionSendNotifyForDisplay},
	{"gdk_set_double_click_time", &SetDoubleClickTime},
	{"gdk_set_locale", &SetLocale},
	{"gdk_set_pointer_hooks", &SetPointerHooks},
	{"gdk_set_program_class", &SetProgramClass},
	{"gdk_set_show_events", &SetShowEvents},
	{"gdk_set_sm_client_id", &SetSmClientId},
	{"gdk_set_use_xshm", &SetUseXshm},
	{"gdk_setting_action_get_type", &SettingActionGetType},
	{"gdk_setting_get", &SettingGet},
	{"gdk_spawn_command_line_on_screen", &SpawnCommandLineOnScreen},
	{"gdk_spawn_on_screen", &SpawnOnScreen},
	{"gdk_spawn_on_screen_with_pipes", &SpawnOnScreenWithPipes},
	{"gdk_status_get_type", &StatusGetType},
	{"gdk_string_extents", &StringExtents},
	{"gdk_string_height", &StringHeight},
	{"gdk_string_measure", &StringMeasure},
	{"gdk_string_to_compound_text", &StringToCompoundText},
	{"gdk_string_to_compound_text_for_display", &StringToCompoundTextForDisplay},
	{"gdk_string_width", &StringWidth},
	{"gdk_subwindow_mode_get_type", &SubwindowModeGetType},
	{"gdk_synthesize_window_state", &SynthesizeWindowState},
	{"gdk_test_render_sync", &TestRenderSync},
	{"gdk_test_simulate_button", &TestSimulateButton},
	{"gdk_test_simulate_key", &TestSimulateKey},
	{"gdk_text_extents", &TextExtents},
	{"gdk_text_extents_wc", &TextExtentsWc},
	{"gdk_text_height", &TextHeight},
	{"gdk_text_measure", &TextMeasure},
	{"gdk_text_property_to_text_list", &TextPropertyToTextList},
	{"gdk_text_property_to_text_list_for_display", &TextPropertyToTextListForDisplay},
	{"gdk_text_property_to_utf8_list", &TextPropertyToUtf8List},
	{"gdk_text_property_to_utf8_list_for_display", &TextPropertyToUtf8ListForDisplay},
	{"gdk_text_width", &TextWidth},
	{"gdk_text_width_wc", &TextWidthWc},
	{"gdk_threads_add_idle", &ThreadsAddIdle},
	{"gdk_threads_add_idle_full", &ThreadsAddIdleFull},
	{"gdk_threads_add_timeout", &ThreadsAddTimeout},
	{"gdk_threads_add_timeout_full", &ThreadsAddTimeoutFull},
	{"gdk_threads_add_timeout_seconds", &ThreadsAddTimeoutSeconds},
	{"gdk_threads_add_timeout_seconds_full", &ThreadsAddTimeoutSecondsFull},
	{"gdk_threads_enter", &ThreadsEnter},
	{"gdk_threads_init", &ThreadsInit},
	{"gdk_threads_leave", &ThreadsLeave},
	{"gdk_threads_set_lock_functions", &ThreadsSetLockFunctions},
	{"gdk_unicode_to_keyval", &UnicodeToKeyval},
	{"gdk_utf8_to_compound_text", &Utf8ToCompoundText},
	{"gdk_utf8_to_compound_text_for_display", &Utf8ToCompoundTextForDisplay},
	{"gdk_utf8_to_string_target", &Utf8ToStringTarget},
	{"gdk_visibility_state_get_type", &VisibilityStateGetType},
	{"gdk_visual_get_best", &VisualGetBest},
	{"gdk_visual_get_best_depth", &VisualGetBestDepth},
	{"gdk_visual_get_best_type", &VisualGetBestType},
	{"gdk_visual_get_best_with_both", &VisualGetBestWithBoth},
	{"gdk_visual_get_best_with_depth", &VisualGetBestWithDepth},
	{"gdk_visual_get_best_with_type", &VisualGetBestWithType},
	{"gdk_visual_get_bits_per_rgb", &VisualGetBitsPerRgb},
	{"gdk_visual_get_blue_pixel_details", &VisualGetBluePixelDetails},
	{"gdk_visual_get_byte_order", &VisualGetByteOrder},
	{"gdk_visual_get_colormap_size", &VisualGetColormapSize},
	{"gdk_visual_get_depth", &VisualGetDepth},
	{"gdk_visual_get_green_pixel_details", &VisualGetGreenPixelDetails},
	{"gdk_visual_get_red_pixel_details", &VisualGetRedPixelDetails},
	{"gdk_visual_get_screen", &VisualGetScreen},
	{"gdk_visual_get_system", &VisualGetSystem},
	{"gdk_visual_get_type", &VisualGetType},
	{"gdk_visual_get_visual_type", &VisualGetVisualType},
	{"gdk_visual_type_get_type", &VisualTypeGetType},
	{"gdk_wcstombs", &Wcstombs},
	{"gdk_win32_begin_direct_draw_libgtk_only", &Win32BeginDirectDrawLibgtkOnly},
	{"gdk_win32_drawable_get_handle", &Win32DrawableGetHandle},
	{"gdk_win32_end_direct_draw_libgtk_only", &Win32EndDirectDrawLibgtkOnly},
	{"gdk_win32_handle_table_lookup", &Win32HandleTableLookup},
	{"gdk_win32_hdc_get", &Win32HdcGet},
	{"gdk_win32_hdc_release", &Win32HdcRelease},
	{"gdk_win32_icon_to_pixbuf_libgtk_only", &Win32IconToPixbufLibgtkOnly},
	{"gdk_win32_pixbuf_to_hicon_libgtk_only", &Win32PixbufToHiconLibgtkOnly},
	{"gdk_win32_selection_add_targets", &Win32SelectionAddTargets},
	{"gdk_win32_set_modal_dialog_libgtk_only", &Win32SetModalDialogLibgtkOnly},
	{"gdk_win32_window_foreign_new_for_display", &Win32WindowForeignNewForDisplay},
	{"gdk_win32_window_get_impl_hwnd", &Win32WindowGetImplHwnd},
	{"gdk_win32_window_is_win32", &Win32WindowIsWin32},
	{"gdk_win32_window_lookup_for_display", &Win32WindowLookupForDisplay},
	{"gdk_window_add_filter", &WindowAddFilter},
	{"gdk_window_at_pointer", &WindowAtPointer},
	{"gdk_window_attributes_type_get_type", &WindowAttributesTypeGetType},
	{"gdk_window_beep", &WindowBeep},
	{"gdk_window_begin_move_drag", &WindowBeginMoveDrag},
	{"gdk_window_begin_paint_rect", &WindowBeginPaintRect},
	{"gdk_window_begin_paint_region", &WindowBeginPaintRegion},
	{"gdk_window_begin_resize_drag", &WindowBeginResizeDrag},
	{"gdk_window_class_get_type", &WindowClassGetType},
	{"gdk_window_clear", &WindowClear},
	{"gdk_window_clear_area", &WindowClearArea},
	{"gdk_window_clear_area_e", &WindowClearAreaE},
	{"gdk_window_configure_finished", &WindowConfigureFinished},
	{"gdk_window_constrain_size", &WindowConstrainSize},
	{"gdk_window_coords_from_parent", &WindowCoordsFromParent},
	{"gdk_window_coords_to_parent", &WindowCoordsToParent},
	{"gdk_window_create_similar_surface", &WindowCreateSimilarSurface},
	{"gdk_window_deiconify", &WindowDeiconify},
	{"gdk_window_destroy", &WindowDestroy},
	{"gdk_window_destroy_notify", &WindowDestroyNotify},
	{"gdk_window_edge_get_type", &WindowEdgeGetType},
	{"gdk_window_enable_synchronized_configure", &WindowEnableSynchronizedConfigure},
	{"gdk_window_end_paint", &WindowEndPaint},
	{"gdk_window_ensure_native", &WindowEnsureNative},
	{"gdk_window_flush", &WindowFlush},
	{"gdk_window_focus", &WindowFocus},
	{"gdk_window_foreign_new", &WindowForeignNew},
	{"gdk_window_foreign_new_for_display", &WindowForeignNewForDisplay},
	{"gdk_window_freeze_toplevel_updates_libgtk_only", &WindowFreezeToplevelUpdatesLibgtkOnly},
	{"gdk_window_freeze_updates", &WindowFreezeUpdates},
	{"gdk_window_fullscreen", &WindowFullscreen},
	{"gdk_window_geometry_changed", &WindowGeometryChanged},
	{"gdk_window_get_accept_focus", &WindowGetAcceptFocus},
	{"gdk_window_get_background_pattern", &WindowGetBackgroundPattern},
	{"gdk_window_get_children", &WindowGetChildren},
	{"gdk_window_get_composited", &WindowGetComposited},
	{"gdk_window_get_cursor", &WindowGetCursor},
	{"gdk_window_get_decorations", &WindowGetDecorations},
	{"gdk_window_get_deskrelative_origin", &WindowGetDeskrelativeOrigin},
	{"gdk_window_get_display", &WindowGetDisplay},
	{"gdk_window_get_effective_parent", &WindowGetEffectiveParent},
	{"gdk_window_get_effective_toplevel", &WindowGetEffectiveToplevel},
	{"gdk_window_get_events", &WindowGetEvents},
	{"gdk_window_get_focus_on_map", &WindowGetFocusOnMap},
	{"gdk_window_get_frame_extents", &WindowGetFrameExtents},
	{"gdk_window_get_geometry", &WindowGetGeometry},
	{"gdk_window_get_group", &WindowGetGroup},
	{"gdk_window_get_height", &WindowGetHeight},
	{"gdk_window_get_internal_paint_info", &WindowGetInternalPaintInfo},
	{"gdk_window_get_modal_hint", &WindowGetModalHint},
	{"gdk_window_get_origin", &WindowGetOrigin},
	{"gdk_window_get_parent", &WindowGetParent},
	{"gdk_window_get_pointer", &WindowGetPointer},
	{"gdk_window_get_position", &WindowGetPosition},
	{"gdk_window_get_root_coords", &WindowGetRootCoords},
	{"gdk_window_get_root_origin", &WindowGetRootOrigin},
	{"gdk_window_get_screen", &WindowGetScreen},
	{"gdk_window_get_state", &WindowGetState},
	{"gdk_window_get_toplevel", &WindowGetToplevel},
	{"gdk_window_get_toplevels", &WindowGetToplevels},
	{"gdk_window_get_type_hint", &WindowGetTypeHint},
	{"gdk_window_get_update_area", &WindowGetUpdateArea},
	{"gdk_window_get_user_data", &WindowGetUserData},
	{"gdk_window_get_visual", &WindowGetVisual},
	{"gdk_window_get_width", &WindowGetWidth},
	{"gdk_window_get_window_type", &WindowGetWindowType},
	{"gdk_window_has_native", &WindowHasNative},
	{"gdk_window_hide", &WindowHide},
	{"gdk_window_hints_get_type", &WindowHintsGetType},
	{"gdk_window_iconify", &WindowIconify},
	// Undocumented {"gdk_window_impl_get_type", &WindowImplGetType},
	{"gdk_window_input_shape_combine_mask", &WindowInputShapeCombineMask},
	{"gdk_window_input_shape_combine_region", &WindowInputShapeCombineRegion},
	{"gdk_window_invalidate_maybe_recurse", &WindowInvalidateMaybeRecurse},
	{"gdk_window_invalidate_rect", &WindowInvalidateRect},
	{"gdk_window_invalidate_region", &WindowInvalidateRegion},
	{"gdk_window_is_destroyed", &WindowIsDestroyed},
	{"gdk_window_is_input_only", &WindowIsInputOnly},
	{"gdk_window_is_shaped", &WindowIsShaped},
	{"gdk_window_is_viewable", &WindowIsViewable},
	{"gdk_window_is_visible", &WindowIsVisible},
	{"gdk_window_lookup", &WindowLookup},
	{"gdk_window_lookup_for_display", &WindowLookupForDisplay},
	{"gdk_window_lower", &WindowLower},
	{"gdk_window_maximize", &WindowMaximize},
	{"gdk_window_merge_child_input_shapes", &WindowMergeChildInputShapes},
	{"gdk_window_merge_child_shapes", &WindowMergeChildShapes},
	{"gdk_window_move", &WindowMove},
	{"gdk_window_move_region", &WindowMoveRegion},
	{"gdk_window_move_resize", &WindowMoveResize},
	{"gdk_window_new", &WindowNew},
	{"gdk_window_object_get_type", &WindowObjectGetType},
	{"gdk_window_peek_children", &WindowPeekChildren},
	{"gdk_window_process_all_updates", &WindowProcessAllUpdates},
	{"gdk_window_process_updates", &WindowProcessUpdates},
	{"gdk_window_raise", &WindowRaise},
	{"gdk_window_redirect_to_drawable", &WindowRedirectToDrawable},
	{"gdk_window_register_dnd", &WindowRegisterDnd},
	{"gdk_window_remove_filter", &WindowRemoveFilter},
	{"gdk_window_remove_redirection", &WindowRemoveRedirection},
	{"gdk_window_reparent", &WindowReparent},
	{"gdk_window_resize", &WindowResize},
	{"gdk_window_restack", &WindowRestack},
	{"gdk_window_scroll", &WindowScroll},
	{"gdk_window_set_accept_focus", &WindowSetAcceptFocus},
	{"gdk_window_set_back_pixmap", &WindowSetBackPixmap},
	{"gdk_window_set_background", &WindowSetBackground},
	{"gdk_window_set_child_input_shapes", &WindowSetChildInputShapes},
	{"gdk_window_set_child_shapes", &WindowSetChildShapes},
	{"gdk_window_set_composited", &WindowSetComposited},
	{"gdk_window_set_cursor", &WindowSetCursor},
	{"gdk_window_set_debug_updates", &WindowSetDebugUpdates},
	{"gdk_window_set_decorations", &WindowSetDecorations},
	{"gdk_window_set_events", &WindowSetEvents},
	{"gdk_window_set_focus_on_map", &WindowSetFocusOnMap},
	{"gdk_window_set_functions", &WindowSetFunctions},
	{"gdk_window_set_geometry_hints", &WindowSetGeometryHints},
	{"gdk_window_set_group", &WindowSetGroup},
	{"gdk_window_set_hints", &WindowSetHints},
	{"gdk_window_set_icon", &WindowSetIcon},
	{"gdk_window_set_icon_list", &WindowSetIconList},
	{"gdk_window_set_icon_name", &WindowSetIconName},
	{"gdk_window_set_keep_above", &WindowSetKeepAbove},
	{"gdk_window_set_keep_below", &WindowSetKeepBelow},
	{"gdk_window_set_modal_hint", &WindowSetModalHint},
	{"gdk_window_set_opacity", &WindowSetOpacity},
	{"gdk_window_set_override_redirect", &WindowSetOverrideRedirect},
	{"gdk_window_set_role", &WindowSetRole},
	{"gdk_window_set_skip_pager_hint", &WindowSetSkipPagerHint},
	{"gdk_window_set_skip_taskbar_hint", &WindowSetSkipTaskbarHint},
	{"gdk_window_set_startup_id", &WindowSetStartupId},
	{"gdk_window_set_static_gravities", &WindowSetStaticGravities},
	{"gdk_window_set_title", &WindowSetTitle},
	{"gdk_window_set_transient_for", &WindowSetTransientFor},
	{"gdk_window_set_type_hint", &WindowSetTypeHint},
	{"gdk_window_set_urgency_hint", &WindowSetUrgencyHint},
	{"gdk_window_set_user_data", &WindowSetUserData},
	{"gdk_window_shape_combine_mask", &WindowShapeCombineMask},
	{"gdk_window_shape_combine_region", &WindowShapeCombineRegion},
	{"gdk_window_show", &WindowShow},
	{"gdk_window_show_unraised", &WindowShowUnraised},
	{"gdk_window_state_get_type", &WindowStateGetType},
	{"gdk_window_stick", &WindowStick},
	{"gdk_window_thaw_toplevel_updates_libgtk_only", &WindowThawToplevelUpdatesLibgtkOnly},
	{"gdk_window_thaw_updates", &WindowThawUpdates},
	{"gdk_window_type_get_type", &WindowTypeGetType},
	{"gdk_window_type_hint_get_type", &WindowTypeHintGetType},
	{"gdk_window_unfullscreen", &WindowUnfullscreen},
	{"gdk_window_unmaximize", &WindowUnmaximize},
	{"gdk_window_unstick", &WindowUnstick},
	{"gdk_window_withdraw", &WindowWithdraw},
	{"gdk_wm_decoration_get_type", &WmDecorationGetType},
	{"gdk_wm_function_get_type", &WmFunctionGetType},
}

var apiListPixbuf = outside.Apis{
	{"gdk_colorspace_get_type", &ColorspaceGetType},
	{"gdk_interp_type_get_type", &InterpTypeGetType},
	{"gdk_pixbuf_add_alpha", &PixbufAddAlpha},
	{"gdk_pixbuf_alpha_mode_get_type", &PixbufAlphaModeGetType},
	{"gdk_pixbuf_animation_get_height", &PixbufAnimationGetHeight},
	{"gdk_pixbuf_animation_get_iter", &PixbufAnimationGetIter},
	{"gdk_pixbuf_animation_get_static_image", &PixbufAnimationGetStaticImage},
	{"gdk_pixbuf_animation_get_type", &PixbufAnimationGetType},
	{"gdk_pixbuf_animation_get_width", &PixbufAnimationGetWidth},
	{"gdk_pixbuf_animation_is_static_image", &PixbufAnimationIsStaticImage},
	{"gdk_pixbuf_animation_iter_advance", &PixbufAnimationIterAdvance},
	{"gdk_pixbuf_animation_iter_get_delay_time", &PixbufAnimationIterGetDelayTime},
	{"gdk_pixbuf_animation_iter_get_pixbuf", &PixbufAnimationIterGetPixbuf},
	{"gdk_pixbuf_animation_iter_get_type", &PixbufAnimationIterGetType},
	{"gdk_pixbuf_animation_iter_on_currently_loading_frame", &PixbufAnimationIterOnCurrentlyLoadingFrame},
	{"gdk_pixbuf_animation_new_from_file", &PixbufAnimationNewFromFile},
	{"gdk_pixbuf_animation_new_from_file_utf8", &PixbufAnimationNewFromFileUtf8},
	{"gdk_pixbuf_animation_ref", &PixbufAnimationRef},
	{"gdk_pixbuf_animation_unref", &PixbufAnimationUnref},
	{"gdk_pixbuf_apply_embedded_orientation", &PixbufApplyEmbeddedOrientation},
	{"gdk_pixbuf_composite", &PixbufComposite},
	{"gdk_pixbuf_composite_color", &PixbufCompositeColor},
	{"gdk_pixbuf_composite_color_simple", &PixbufCompositeColorSimple},
	{"gdk_pixbuf_copy", &PixbufCopy},
	{"gdk_pixbuf_copy_area", &PixbufCopyArea},
	{"gdk_pixbuf_error_get_type", &PixbufErrorGetType},
	{"gdk_pixbuf_error_quark", &PixbufErrorQuark},
	{"gdk_pixbuf_fill", &PixbufFill},
	{"gdk_pixbuf_flip", &PixbufFlip},
	{"gdk_pixbuf_format_copy", &PixbufFormatCopy},
	{"gdk_pixbuf_format_free", &PixbufFormatFree},
	{"gdk_pixbuf_format_get_description", &PixbufFormatGetDescription},
	{"gdk_pixbuf_format_get_extensions", &PixbufFormatGetExtensions},
	{"gdk_pixbuf_format_get_license", &PixbufFormatGetLicense},
	{"gdk_pixbuf_format_get_mime_types", &PixbufFormatGetMimeTypes},
	{"gdk_pixbuf_format_get_name", &PixbufFormatGetName},
	{"gdk_pixbuf_format_get_type", &PixbufFormatGetType},
	{"gdk_pixbuf_format_is_disabled", &PixbufFormatIsDisabled},
	{"gdk_pixbuf_format_is_scalable", &PixbufFormatIsScalable},
	{"gdk_pixbuf_format_is_writable", &PixbufFormatIsWritable},
	{"gdk_pixbuf_format_set_disabled", &PixbufFormatSetDisabled},
	{"gdk_pixbuf_from_pixdata", &PixbufFromPixdata},
	{"gdk_pixbuf_get_bits_per_sample", &PixbufGetBitsPerSample},
	{"gdk_pixbuf_get_colorspace", &PixbufGetColorspace},
	{"gdk_pixbuf_get_file_info", &PixbufGetFileInfo},
	{"gdk_pixbuf_get_formats", &PixbufGetFormats},
	{"gdk_pixbuf_get_has_alpha", &PixbufGetHasAlpha},
	{"gdk_pixbuf_get_height", &PixbufGetHeight},
	{"gdk_pixbuf_get_n_channels", &PixbufGetNChannels},
	{"gdk_pixbuf_get_option", &PixbufGetOption},
	{"gdk_pixbuf_get_pixels", &PixbufGetPixels},
	{"gdk_pixbuf_get_rowstride", &PixbufGetRowstride},
	{"gdk_pixbuf_get_type", &PixbufGetType},
	{"gdk_pixbuf_get_width", &PixbufGetWidth},
	// Undocumented {"gdk_pixbuf_gettext", &PixbufGettext},
	{"gdk_pixbuf_loader_close", &PixbufLoaderClose},
	{"gdk_pixbuf_loader_get_animation", &PixbufLoaderGetAnimation},
	{"gdk_pixbuf_loader_get_format", &PixbufLoaderGetFormat},
	{"gdk_pixbuf_loader_get_pixbuf", &PixbufLoaderGetPixbuf},
	{"gdk_pixbuf_loader_get_type", &PixbufLoaderGetType},
	{"gdk_pixbuf_loader_new", &PixbufLoaderNew},
	{"gdk_pixbuf_loader_new_with_mime_type", &PixbufLoaderNewWithMimeType},
	{"gdk_pixbuf_loader_new_with_type", &PixbufLoaderNewWithType},
	{"gdk_pixbuf_loader_set_size", &PixbufLoaderSetSize},
	{"gdk_pixbuf_loader_write", &PixbufLoaderWrite},
	{"gdk_pixbuf_new", &PixbufNew},
	{"gdk_pixbuf_new_from_data", &PixbufNewFromData},
	// Win32 uses _utf8 {"gdk_pixbuf_new_from_file", &PixbufNewFromFile},
	// Win32 uese _utf8 {"gdk_pixbuf_new_from_file_at_scale", &PixbufNewFromFileAtScale},
	{"gdk_pixbuf_new_from_file_at_scale_utf8", &PixbufNewFromFileAtScale},
	// Win32 uses _utf8 {"gdk_pixbuf_new_from_file_at_size", &PixbufNewFromFileAtSize},
	{"gdk_pixbuf_new_from_file_at_size_utf8", &PixbufNewFromFileAtSize},
	{"gdk_pixbuf_new_from_file_utf8", &PixbufNewFromFile},
	{"gdk_pixbuf_new_from_inline", &PixbufNewFromInline},
	{"gdk_pixbuf_new_from_stream", &PixbufNewFromStream},
	{"gdk_pixbuf_new_from_stream_async", &PixbufNewFromStreamAsync},
	{"gdk_pixbuf_new_from_stream_at_scale", &PixbufNewFromStreamAtScale},
	{"gdk_pixbuf_new_from_stream_at_scale_async", &PixbufNewFromStreamAtScaleAsync},
	{"gdk_pixbuf_new_from_stream_finish", &PixbufNewFromStreamFinish},
	{"gdk_pixbuf_new_from_xpm_data", &PixbufNewFromXpmData},
	{"gdk_pixbuf_new_subpixbuf", &PixbufNewSubpixbuf},
	{"gdk_pixbuf_non_anim_get_type", &PixbufNonAnimGetType},
	{"gdk_pixbuf_non_anim_new", &PixbufNonAnimNew},
	{"gdk_pixbuf_ref", &PixbufRef},
	{"gdk_pixbuf_rotate_simple", &PixbufRotateSimple},
	{"gdk_pixbuf_rotation_get_type", &PixbufRotationGetType},
	{"gdk_pixbuf_saturate_and_pixelate", &PixbufSaturateAndPixelate},
	// Win32 uses _utf8 {"gdk_pixbuf_save", &PixbufSave},
	{"gdk_pixbuf_save_to_buffer", &PixbufSaveToBuffer},
	{"gdk_pixbuf_save_to_bufferv", &PixbufSaveToBufferv},
	{"gdk_pixbuf_save_to_callback", &PixbufSaveToCallback},
	{"gdk_pixbuf_save_to_callbackv", &PixbufSaveToCallbackv},
	{"gdk_pixbuf_save_to_stream", &PixbufSaveToStream},
	{"gdk_pixbuf_save_to_stream_async", &PixbufSaveToStreamAsync},
	{"gdk_pixbuf_save_to_stream_finish", &PixbufSaveToStreamFinish},
	{"gdk_pixbuf_save_utf8", &PixbufSaveUtf8},
	// Win32 uses _utf8 {"gdk_pixbuf_savev", &PixbufSavev},
	{"gdk_pixbuf_savev_utf8", &PixbufSavevUtf8},
	{"gdk_pixbuf_scale", &PixbufScale},
	{"gdk_pixbuf_scale_simple", &PixbufScaleSimple},
	// Undocumented {"gdk_pixbuf_scaled_anim_get_type", &PixbufScaledAnimGetType},
	// Undocumented {"gdk_pixbuf_scaled_anim_iter_get_type", &PixbufScaledAnimIterGetType},
	{"gdk_pixbuf_set_option", &PixbufSetOption},
	{"gdk_pixbuf_simple_anim_add_frame", &PixbufSimpleAnimAddFrame},
	{"gdk_pixbuf_simple_anim_get_loop", &PixbufSimpleAnimGetLoop},
	{"gdk_pixbuf_simple_anim_get_type", &PixbufSimpleAnimGetType},
	{"gdk_pixbuf_simple_anim_iter_get_type", &PixbufSimpleAnimIterGetType},
	{"gdk_pixbuf_simple_anim_new", &PixbufSimpleAnimNew},
	{"gdk_pixbuf_simple_anim_set_loop", &PixbufSimpleAnimSetLoop},
	{"gdk_pixbuf_unref", &PixbufUnref},
	{"gdk_pixdata_deserialize", &PixdataDeserialize},
	{"gdk_pixdata_from_pixbuf", &PixdataFromPixbuf},
	{"gdk_pixdata_serialize", &PixdataSerialize},
	{"gdk_pixdata_to_csource", &PixdataToCsource},
}

var dataList = outside.Data{
	{"gdk_threads_lock", (*T.GCallback)(nil)},
	{"gdk_threads_mutex", (*T.GMutex)(nil)},
	{"gdk_threads_unlock", (*T.GCallback)(nil)},
}

var dataListPixbuf = outside.Data{
	{"gdk_pixbuf_major_version", (*uint)(nil)},
	{"gdk_pixbuf_micro_version", (*uint)(nil)},
	{"gdk_pixbuf_minor_version", (*uint)(nil)},
	{"gdk_pixbuf_version", (*uint8)(nil)},
}
