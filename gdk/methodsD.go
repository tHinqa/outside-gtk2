// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Device struct {
	Parent    T.GObject
	Name      *T.Gchar
	Source    T.GdkInputSource
	Mode      T.GdkInputMode
	HasCursor T.Gboolean
	NumAxes   int
	Axes      *DeviceAxis
	NumKeys   int
	Keys      *DeviceKey
}

var (
	DeviceGetType func() T.GType

	DeviceFreeHistory    func(events **T.GdkTimeCoord, nEvents int)
	DeviceGetCorePointer func() *Device
	DevicesList          func() *T.GList

	deviceGetAxis      func(d *Device, axes *float64, use T.GdkAxisUse, value *float64) T.Gboolean
	deviceGetAxisUse   func(d *Device, index uint) T.GdkAxisUse
	deviceGetHasCursor func(d *Device) T.Gboolean
	deviceGetHistory   func(d *Device, window *Window, start, stop T.GUint32, events ***T.GdkTimeCoord, nEvents *int) T.Gboolean
	deviceGetKey       func(d *Device, index uint, keyval *uint, modifiers *T.GdkModifierType)
	deviceGetMode      func(d *Device) T.GdkInputMode
	deviceGetName      func(d *Device) string
	deviceGetNAxes     func(d *Device) int
	deviceGetNKeys     func(d *Device) int
	deviceGetSource    func(d *Device) T.GdkInputSource
	deviceGetState     func(d *Device, window *Window, axes *float64, mask *T.GdkModifierType)
	deviceSetAxisUse   func(d *Device, index uint, use T.GdkAxisUse)
	deviceSetKey       func(d *Device, index, keyval uint, modifiers T.GdkModifierType)
	deviceSetMode      func(d *Device, mode T.GdkInputMode) T.Gboolean
	deviceSetSource    func(d *Device, source T.GdkInputSource)
)

func (d *Device) GetAxis(axes *float64, use T.GdkAxisUse, value *float64) T.Gboolean {
	return deviceGetAxis(d, axes, use, value)
}
func (d *Device) GetAxisUse(index uint) T.GdkAxisUse { return deviceGetAxisUse(d, index) }
func (d *Device) GetHasCursor() T.Gboolean           { return deviceGetHasCursor(d) }
func (d *Device) GetHistory(window *Window, start, stop T.GUint32, events ***T.GdkTimeCoord, nEvents *int) T.Gboolean {
	return deviceGetHistory(d, window, start, stop, events, nEvents)
}
func (d *Device) GetKey(index uint, keyval *uint, modifiers *T.GdkModifierType) {
	deviceGetKey(d, index, keyval, modifiers)
}
func (d *Device) GetMode() T.GdkInputMode     { return deviceGetMode(d) }
func (d *Device) GetName() string             { return deviceGetName(d) }
func (d *Device) GetNAxes() int               { return deviceGetNAxes(d) }
func (d *Device) GetNKeys() int               { return deviceGetNKeys(d) }
func (d *Device) GetSource() T.GdkInputSource { return deviceGetSource(d) }
func (d *Device) GetState(window *Window, axes *float64, mask *T.GdkModifierType) {
	deviceGetState(d, window, axes, mask)
}
func (d *Device) SetAxisUse(index uint, use T.GdkAxisUse) { deviceSetAxisUse(d, index, use) }
func (d *Device) SetKey(index, keyval uint, modifiers T.GdkModifierType) {
	deviceSetKey(d, index, keyval, modifiers)
}
func (d *Device) SetMode(mode T.GdkInputMode) T.Gboolean { return deviceSetMode(d, mode) }
func (d *Device) SetSource(source T.GdkInputSource)      { deviceSetSource(d, source) }

type DeviceKey struct {
	Keyval    uint
	Modifiers T.GdkModifierType
}

type DeviceAxis struct {
	Use T.GdkAxisUse
	Min float64
	Max float64
}

type Display struct {
	Parent          T.GObject
	QueuedEvents    *T.GList
	QueuedTail      *T.GList
	ButtonClickTime [2]T.GUint32
	ButtonWindow    [2]*Window
	ButtonNumber    [2]int
	DoubleClickTime uint
	CorePointer     *Device
	PointerHooks    *DisplayPointerHooks
	Bits            uint
	// Closed : 1
	// IgnoreCoreEvents : 1
	DoubleClickDistance uint
	ButtonX             [2]int
	ButtonY             [2]int
	PointerGrabs        *T.GList
	KeyboardGrab        T.GdkKeyboardGrabInfo
	PointerInfo         T.GdkPointerWindowInfo
	LastEventTime       T.GUint32
}

var (
	DisplayGetType func() T.GType

	DisplayGetDefault            func() *Display
	DisplayOpen                  func(displayName string) *Display
	DisplayOpenDefaultLibgtkOnly func() *Display

	displayAddClientMessageFilter        func(d *Display, messageType Atom, f T.GdkFilterFunc, data T.Gpointer)
	displayBeep                          func(d *Display)
	displayClose                         func(d *Display)
	displayFlush                         func(d *Display)
	displayGetCorePointer                func(d *Display) *Device
	displayGetDefaultCursorSize          func(d *Display) uint
	displayGetDefaultGroup               func(d *Display) *Window
	displayGetDefaultScreen              func(d *Display) *Screen
	displayGetEvent                      func(d *Display) *Event
	displayGetMaximalCursorSize          func(d *Display, width, height *uint)
	displayGetName                       func(d *Display) string
	displayGetNScreens                   func(d *Display) int
	displayGetPointer                    func(d *Display, screen **Screen, x, y *int, mask *T.GdkModifierType)
	displayGetScreen                     func(d *Display, screenNum int) *Screen
	displayGetWindowAtPointer            func(d *Display, winX, winY *int) *Window
	displayIsClosed                      func(d *Display) T.Gboolean
	displayKeyboardUngrab                func(d *Display, time T.GUint32)
	displayListDevices                   func(d *Display) *T.GList
	displayPeekEvent                     func(d *Display) *Event
	displayPointerIsGrabbed              func(d *Display) T.Gboolean
	displayPointerUngrab                 func(d *Display, time T.GUint32)
	displayPutEvent                      func(d *Display, event *Event)
	displayRequestSelectionNotification  func(d *Display, selection Atom) T.Gboolean
	displaySetDoubleClickDistance        func(d *Display, distance uint)
	displaySetDoubleClickTime            func(d *Display, msec uint)
	displaySetPointerHooks               func(d *Display, newHooks *DisplayPointerHooks) *DisplayPointerHooks
	displayStoreClipboard                func(d *Display, clipboardWindow *Window, time T.GUint32, targets *Atom, nTargets int)
	displaySupportsClipboardPersistence  func(d *Display) T.Gboolean
	displaySupportsComposite             func(d *Display) T.Gboolean
	displaySupportsCursorAlpha           func(d *Display) T.Gboolean
	displaySupportsCursorColor           func(d *Display) T.Gboolean
	displaySupportsInputShapes           func(d *Display) T.Gboolean
	displaySupportsSelectionNotification func(d *Display) T.Gboolean
	displaySupportsShapes                func(d *Display) T.Gboolean
	displaySync                          func(d *Display)
	displayWarpPointer                   func(d *Display, screen *Screen, x, y int)
)

func (d *Display) AddClientMessageFilter(messageType Atom, f T.GdkFilterFunc, data T.Gpointer) {
	displayAddClientMessageFilter(d, messageType, f, data)
}
func (d *Display) Beep()                      { displayBeep(d) }
func (d *Display) Close()                     { displayClose(d) }
func (d *Display) Flush()                     { displayFlush(d) }
func (d *Display) GetCorePointer() *Device    { return displayGetCorePointer(d) }
func (d *Display) GetDefaultCursorSize() uint { return displayGetDefaultCursorSize(d) }
func (d *Display) GetDefaultGroup() *Window   { return displayGetDefaultGroup(d) }
func (d *Display) GetDefaultScreen() *Screen  { return displayGetDefaultScreen(d) }
func (d *Display) GetEvent() *Event           { return displayGetEvent(d) }
func (d *Display) GetMaximalCursorSize(width, height *uint) {
	displayGetMaximalCursorSize(d, width, height)
}
func (d *Display) GetName() string  { return displayGetName(d) }
func (d *Display) GetNScreens() int { return displayGetNScreens(d) }
func (d *Display) GetPointer(screen **Screen, x, y *int, mask *T.GdkModifierType) {
	displayGetPointer(d, screen, x, y, mask)
}
func (d *Display) GetScreen(screenNum int) *Screen { return displayGetScreen(d, screenNum) }
func (d *Display) GetWindowAtPointer(winX, winY *int) *Window {
	return displayGetWindowAtPointer(d, winX, winY)
}
func (d *Display) IsClosed() T.Gboolean          { return displayIsClosed(d) }
func (d *Display) KeyboardUngrab(time T.GUint32) { displayKeyboardUngrab(d, time) }
func (d *Display) ListDevices() *T.GList         { return displayListDevices(d) }
func (d *Display) PeekEvent() *Event             { return displayPeekEvent(d) }
func (d *Display) PointerIsGrabbed() T.Gboolean  { return displayPointerIsGrabbed(d) }
func (d *Display) PointerUngrab(time T.GUint32)  { displayPointerUngrab(d, time) }
func (d *Display) PutEvent(event *Event)         { displayPutEvent(d, event) }
func (d *Display) RequestSelectionNotification(selection Atom) T.Gboolean {
	return displayRequestSelectionNotification(d, selection)
}
func (d *Display) SetDoubleClickDistance(distance uint) { displaySetDoubleClickDistance(d, distance) }
func (d *Display) SetDoubleClickTime(msec uint)         { displaySetDoubleClickTime(d, msec) }
func (d *Display) SetPointerHooks(newHooks *DisplayPointerHooks) *DisplayPointerHooks {
	return displaySetPointerHooks(d, newHooks)
}
func (d *Display) StoreClipboard(clipboardWindow *Window, time T.GUint32, targets *Atom, nTargets int) {
	displayStoreClipboard(d, clipboardWindow, time, targets, nTargets)
}
func (d *Display) SupportsClipboardPersistence() T.Gboolean {
	return displaySupportsClipboardPersistence(d)
}
func (d *Display) SupportsComposite() T.Gboolean   { return displaySupportsComposite(d) }
func (d *Display) SupportsCursorAlpha() T.Gboolean { return displaySupportsCursorAlpha(d) }
func (d *Display) SupportsCursorColor() T.Gboolean { return displaySupportsCursorColor(d) }
func (d *Display) SupportsInputShapes() T.Gboolean { return displaySupportsInputShapes(d) }
func (d *Display) SupportsSelectionNotification() T.Gboolean {
	return displaySupportsSelectionNotification(d)
}
func (d *Display) SupportsShapes() T.Gboolean           { return displaySupportsShapes(d) }
func (d *Display) Sync()                                { displaySync(d) }
func (d *Display) WarpPointer(screen *Screen, x, y int) { displayWarpPointer(d, screen, x, y) }

type DisplayPointerHooks struct {
	GetPointer func(
		display *Display, screen **Screen,
		x, y *int, mask *T.GdkModifierType)
	WindowGetPointer func(
		display *Display, window *Window,
		x, y *int, mask *T.GdkModifierType) *Window
	WindowAtPointer func(
		display *Display, winX, winY *int) *Window
}

type DisplayManager struct{}

var (
	DisplayManagerGetType func() T.GType

	DisplayManagerGet func() *DisplayManager

	displayManagerGetDefaultDisplay func(d *DisplayManager) *Display
	displayManagerListDisplays      func(d *DisplayManager) *T.GSList
	displayManagerSetDefaultDisplay func(d *DisplayManager, display *Display)
)

func (d *DisplayManager) GetDefaultDisplay() *Display { return displayManagerGetDefaultDisplay(d) }
func (d *DisplayManager) ListDisplays() *T.GSList     { return displayManagerListDisplays(d) }
func (d *DisplayManager) SetDefaultDisplay(display *Display) {
	displayManagerSetDefaultDisplay(d, display)
}

type DragAction Enum

const (
	ACTION_DEFAULT DragAction = 1 << iota
	ACTION_COPY
	ACTION_MOVE
	ACTION_LINK
	ACTION_PRIVATE
	ACTION_ASK
)

var DragActionGetType func() T.GType

type DragContext struct {
	Parent          T.GObject
	Protocol        DragProtocol
	Is_source       T.Gboolean
	SourceWindow    *Window
	DestWindow      *Window
	Targets         *T.GList
	Actions         DragAction
	SuggestedAction DragAction
	Action          DragAction
	StartTime       T.GUint32
	WindowingData   T.Gpointer
}

var (
	DragContextGetType func() T.GType
	DragContextNew     func() *DragContext

	DragAbort                 func(context *DragContext, time T.GUint32)
	DragBegin                 func(window *Window, targets *T.GList) *DragContext
	DragGetProtocol           func(xid T.GdkNativeWindow, protocol *DragProtocol) T.GdkNativeWindow
	DragGetProtocolForDisplay func(display *Display, xid T.GdkNativeWindow, protocol *DragProtocol) T.GdkNativeWindow

	DragDrop                func(context *DragContext, time T.GUint32)
	DragDropSucceeded       func(context *DragContext) T.Gboolean
	DragFindWindow          func(context *DragContext, dragWindow *Window, xRoot int, yRoot int, destWindow **Window, protocol *DragProtocol)
	DragFindWindowForScreen func(context *DragContext, dragWindow *Window, screen *Screen, xRoot int, yRoot int, destWindow **Window, protocol *DragProtocol)
	DragGetSelection        func(context *DragContext) Atom
	DragMotion              func(context *DragContext, destWindow *Window, protocol DragProtocol, xRoot int, yRoot int, suggestedAction DragAction, possibleActions DragAction, time T.GUint32) T.Gboolean
	DragStatus              func(context *DragContext, action DragAction, time T.GUint32)
	DropFinish              func(context *DragContext, success T.Gboolean, time T.GUint32)
	DropReply               func(context *DragContext, ok T.Gboolean, time T.GUint32)

	dragContextGetActions         func(context *DragContext) DragAction
	dragContextGetDestWindow      func(context *DragContext) *Window
	dragContextGetProtocol        func(context *DragContext) DragProtocol
	dragContextGetSelectedAction  func(context *DragContext) DragAction
	dragContextGetSourceWindow    func(context *DragContext) *Window
	dragContextGetSuggestedAction func(context *DragContext) DragAction
	dragContextListTargets        func(context *DragContext) *T.GList
	dragContextRef                func(context *DragContext)
	dragContextUnref              func(context *DragContext)
)

func (d *DragContext) GetActions() DragAction         { return dragContextGetActions(d) }
func (d *DragContext) GetDestWindow() *Window         { return dragContextGetDestWindow(d) }
func (d *DragContext) GetProtocol() DragProtocol      { return dragContextGetProtocol(d) }
func (d *DragContext) GetSelectedAction() DragAction  { return dragContextGetSelectedAction(d) }
func (d *DragContext) GetSourceWindow() *Window       { return dragContextGetSourceWindow(d) }
func (d *DragContext) GetSuggestedAction() DragAction { return dragContextGetSuggestedAction(d) }
func (d *DragContext) ListTargets() *T.GList          { return dragContextListTargets(d) }
func (d *DragContext) Ref()                           { dragContextRef(d) }
func (d *DragContext) Unref()                         { dragContextUnref(d) }

type DragProtocol Enum

const (
	DRAG_PROTO_MOTIF DragProtocol = iota
	DRAG_PROTO_XDND
	DRAG_PROTO_ROOTWIN
	DRAG_PROTO_NONE
	DRAG_PROTO_WIN32_DROPFILES
	DRAG_PROTO_OLE2
	DRAG_PROTO_LOCAL
)

var DragProtocolGetType func() T.GType

type Drawable simpleObject

var (
	DrawableGetType func() T.GType

	DrawArc                  func(drawable *Drawable, gc *T.GdkGC, filled T.Gboolean, x, y, width, height, angle1, angle2 int)
	DrawDrawable             func(drawable *Drawable, gc *T.GdkGC, src *Drawable, xsrc, ysrc, xdest, ydest, width, height int)
	DrawGlyphs               func(drawable *Drawable, gc *T.GdkGC, font *T.PangoFont, x, y int, glyphs *T.PangoGlyphString)
	DrawGlyphsTransformed    func(drawable *Drawable, gc *T.GdkGC, matrix *T.PangoMatrix, font *T.PangoFont, x, y int, glyphs *T.PangoGlyphString)
	DrawGrayImage            func(drawable *Drawable, gc *T.GdkGC, x, y, width, height int, dith T.GdkRgbDither, buf *T.Guchar, rowstride int)
	DrawImage                func(drawable *Drawable, gc *T.GdkGC, image *Image, xsrc, ysrc, xdest, ydest, width, height int)
	DrawIndexedImage         func(drawable *Drawable, gc *T.GdkGC, x, y, width, height int, dith T.GdkRgbDither, buf *T.Guchar, rowstride int, cmap *T.GdkRgbCmap)
	DrawLayout               func(drawable *Drawable, gc *T.GdkGC, x, y int, layout *T.PangoLayout)
	DrawLayoutLine           func(drawable *Drawable, gc *T.GdkGC, x, y int, line *T.PangoLayoutLine)
	DrawLayoutLineWithColors func(drawable *Drawable, gc *T.GdkGC, x, y int, line *T.PangoLayoutLine, foreground, background *Color)
	DrawLayoutWithColors     func(drawable *Drawable, gc *T.GdkGC, x, y int, layout *T.PangoLayout, foreground, background *Color)
	DrawLine                 func(drawable *Drawable, gc *T.GdkGC, x1, y1, x2, y2 int)
	DrawLines                func(drawable *Drawable, gc *T.GdkGC, points *T.GdkPoint, nPoints int)
	DrawPixbuf               func(drawable *Drawable, gc *T.GdkGC, pixbuf *Pixbuf, srcX, srcY, destX, destY, width, height int, dither T.GdkRgbDither, xDither, yDither int)
	DrawPoint                func(drawable *Drawable, gc *T.GdkGC, x, y int)
	DrawPoints               func(drawable *Drawable, gc *T.GdkGC, points *T.GdkPoint, nPoints int)
	DrawPolygon              func(drawable *Drawable, gc *T.GdkGC, filled T.Gboolean, points *T.GdkPoint, nPoints int)
	DrawRectangle            func(drawable *Drawable, gc *T.GdkGC, filled T.Gboolean, x, y, width, height int)
	DrawRgb32Image           func(drawable *Drawable, gc *T.GdkGC, x, y, width, height int, dith T.GdkRgbDither, buf *T.Guchar, rowstride int)
	DrawRgb32ImageDithalign  func(drawable *Drawable, gc *T.GdkGC, x, y, width, height int, dith T.GdkRgbDither, buf *T.Guchar, rowstride, xdith, ydith int)
	DrawRgbImage             func(drawable *Drawable, gc *T.GdkGC, x, y, width, height int, dith T.GdkRgbDither, rgbBuf *T.Guchar, rowstride int)
	DrawRgbImageDithalign    func(drawable *Drawable, gc *T.GdkGC, x, y, width, height int, dith T.GdkRgbDither, rgbBuf *T.Guchar, rowstride, xdith, ydith int)
	DrawSegments             func(drawable *Drawable, gc *T.GdkGC, segs *T.GdkSegment, nSegs int)
	DrawString               func(drawable *Drawable, font *T.GdkFont, gc *T.GdkGC, x, y int, s string)
	DrawText                 func(drawable *Drawable, font *T.GdkFont, gc *T.GdkGC, x, y int, text string, textLength int)
	DrawTextWc               func(drawable *Drawable, font *T.GdkFont, gc *T.GdkGC, x, y int, text *T.GdkWChar, textLength int)
	DrawTrapezoids           func(drawable *Drawable, gc *T.GdkGC, trapezoids *T.GdkTrapezoid, nTrapezoids int)

	drawableCopyToImage      func(d *Drawable, image *Image, srcX, srcY, destX, destY, width, height int) *Image
	drawableGetClipRegion    func(d *Drawable) *T.GdkRegion
	drawableGetColormap      func(d *Drawable) *Colormap
	drawableGetData          func(d *Drawable, key string) T.Gpointer
	drawableGetDepth         func(d *Drawable) int
	drawableGetDisplay       func(d *Drawable) *Display
	drawableGetImage         func(d *Drawable, x, y, width, height int) *Image
	drawableGetScreen        func(d *Drawable) *Screen
	drawableGetSize          func(d *Drawable, width, height *int)
	drawableGetVisibleRegion func(d *Drawable) *T.GdkRegion
	drawableGetVisual        func(d *Drawable) *Visual
	drawableRef              func(d *Drawable) *Drawable
	drawableSetColormap      func(d *Drawable, colormap *Colormap)
	drawableSetData          func(d *Drawable, key string, data T.Gpointer, destroyFunc T.GDestroyNotify)
	drawableUnref            func(d *Drawable)
)

func (d *Drawable) CopyToImage(image *Image, srcX, srcY, destX, destY, width, height int) *Image {
	return drawableCopyToImage(d, image, srcX, srcY, destX, destY, width, height)
}
func (d *Drawable) GetClipRegion() *T.GdkRegion   { return drawableGetClipRegion(d) }
func (d *Drawable) GetColormap() *Colormap        { return drawableGetColormap(d) }
func (d *Drawable) GetData(key string) T.Gpointer { return drawableGetData(d, key) }
func (d *Drawable) GetDepth() int                 { return drawableGetDepth(d) }
func (d *Drawable) GetDisplay() *Display          { return drawableGetDisplay(d) }
func (d *Drawable) GetImage(x, y, width, height int) *Image {
	return drawableGetImage(d, x, y, width, height)
}
func (d *Drawable) GetScreen() *Screen             { return drawableGetScreen(d) }
func (d *Drawable) GetSize(width, height *int)     { drawableGetSize(d, width, height) }
func (d *Drawable) GetVisibleRegion() *T.GdkRegion { return drawableGetVisibleRegion(d) }
func (d *Drawable) GetVisual() *Visual             { return drawableGetVisual(d) }
func (d *Drawable) Ref() *Drawable                 { return drawableRef(d) }
func (d *Drawable) SetColormap(colormap *Colormap) { drawableSetColormap(d, colormap) }
func (d *Drawable) SetData(key string, data T.Gpointer, destroyFunc T.GDestroyNotify) {
	drawableSetData(d, key, data, destroyFunc)
}
func (d *Drawable) Unref() { drawableUnref(d) }
