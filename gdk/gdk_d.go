// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Device struct {
	Parent    O.Object
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
	DeviceGetType func() O.Type

	DeviceFreeHistory    func(events **T.GdkTimeCoord, nEvents int)
	DeviceGetCorePointer func() *Device
	DevicesList          func() *L.List

	DeviceGetAxis      func(d *Device, axes *float64, use T.GdkAxisUse, value *float64) bool
	DeviceGetAxisUse   func(d *Device, index uint) T.GdkAxisUse
	DeviceGetHasCursor func(d *Device) bool
	DeviceGetHistory   func(d *Device, window *Window, start, stop T.GUint32, events ***T.GdkTimeCoord, nEvents *int) bool
	DeviceGetKey       func(d *Device, index uint, keyval *uint, modifiers *T.GdkModifierType)
	DeviceGetMode      func(d *Device) T.GdkInputMode
	DeviceGetName      func(d *Device) string
	DeviceGetNAxes     func(d *Device) int
	DeviceGetNKeys     func(d *Device) int
	DeviceGetSource    func(d *Device) T.GdkInputSource
	DeviceGetState     func(d *Device, window *Window, axes *float64, mask *T.GdkModifierType)
	DeviceSetAxisUse   func(d *Device, index uint, use T.GdkAxisUse)
	DeviceSetKey       func(d *Device, index, keyval uint, modifiers T.GdkModifierType)
	DeviceSetMode      func(d *Device, mode T.GdkInputMode) bool
	DeviceSetSource    func(d *Device, source T.GdkInputSource)
)

func (d *Device) GetAxis(axes *float64, use T.GdkAxisUse, value *float64) bool {
	return DeviceGetAxis(d, axes, use, value)
}
func (d *Device) GetAxisUse(index uint) T.GdkAxisUse { return DeviceGetAxisUse(d, index) }
func (d *Device) GetHasCursor() bool                 { return DeviceGetHasCursor(d) }
func (d *Device) GetHistory(window *Window, start, stop T.GUint32, events ***T.GdkTimeCoord, nEvents *int) bool {
	return DeviceGetHistory(d, window, start, stop, events, nEvents)
}
func (d *Device) GetKey(index uint, keyval *uint, modifiers *T.GdkModifierType) {
	DeviceGetKey(d, index, keyval, modifiers)
}
func (d *Device) GetMode() T.GdkInputMode     { return DeviceGetMode(d) }
func (d *Device) GetName() string             { return DeviceGetName(d) }
func (d *Device) GetNAxes() int               { return DeviceGetNAxes(d) }
func (d *Device) GetNKeys() int               { return DeviceGetNKeys(d) }
func (d *Device) GetSource() T.GdkInputSource { return DeviceGetSource(d) }
func (d *Device) GetState(window *Window, axes *float64, mask *T.GdkModifierType) {
	DeviceGetState(d, window, axes, mask)
}
func (d *Device) SetAxisUse(index uint, use T.GdkAxisUse) { DeviceSetAxisUse(d, index, use) }
func (d *Device) SetKey(index, keyval uint, modifiers T.GdkModifierType) {
	DeviceSetKey(d, index, keyval, modifiers)
}
func (d *Device) SetMode(mode T.GdkInputMode) bool  { return DeviceSetMode(d, mode) }
func (d *Device) SetSource(source T.GdkInputSource) { DeviceSetSource(d, source) }

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
	Parent          O.Object
	QueuedEvents    *L.List
	QueuedTail      *L.List
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
	PointerGrabs        *L.List
	KeyboardGrab        T.GdkKeyboardGrabInfo
	PointerInfo         T.GdkPointerWindowInfo
	LastEventTime       T.GUint32
}

var (
	DisplayGetType func() O.Type

	DisplayGetDefault            func() *Display
	DisplayOpen                  func(displayName string) *Display
	DisplayOpenDefaultLibgtkOnly func() *Display

	DisplayAddClientMessageFilter        func(d *Display, messageType Atom, f T.GdkFilterFunc, data T.Gpointer)
	DisplayBeep                          func(d *Display)
	DisplayClose                         func(d *Display)
	DisplayFlush                         func(d *Display)
	DisplayGetCorePointer                func(d *Display) *Device
	DisplayGetDefaultCursorSize          func(d *Display) uint
	DisplayGetDefaultGroup               func(d *Display) *Window
	DisplayGetDefaultScreen              func(d *Display) *Screen
	DisplayGetEvent                      func(d *Display) *Event
	DisplayGetMaximalCursorSize          func(d *Display, width, height *uint)
	DisplayGetName                       func(d *Display) string
	DisplayGetNScreens                   func(d *Display) int
	DisplayGetPointer                    func(d *Display, screen **Screen, x, y *int, mask *T.GdkModifierType)
	DisplayGetScreen                     func(d *Display, screenNum int) *Screen
	DisplayGetWindowAtPointer            func(d *Display, winX, winY *int) *Window
	DisplayIsClosed                      func(d *Display) bool
	DisplayKeyboardUngrab                func(d *Display, time T.GUint32)
	DisplayListDevices                   func(d *Display) *L.List
	DisplayPeekEvent                     func(d *Display) *Event
	DisplayPointerIsGrabbed              func(d *Display) bool
	DisplayPointerUngrab                 func(d *Display, time T.GUint32)
	DisplayPutEvent                      func(d *Display, event *Event)
	DisplayRequestSelectionNotification  func(d *Display, selection Atom) bool
	DisplaySetDoubleClickDistance        func(d *Display, distance uint)
	DisplaySetDoubleClickTime            func(d *Display, msec uint)
	DisplaySetPointerHooks               func(d *Display, newHooks *DisplayPointerHooks) *DisplayPointerHooks
	DisplayStoreClipboard                func(d *Display, clipboardWindow *Window, time T.GUint32, targets *Atom, nTargets int)
	DisplaySupportsClipboardPersistence  func(d *Display) bool
	DisplaySupportsComposite             func(d *Display) bool
	DisplaySupportsCursorAlpha           func(d *Display) bool
	DisplaySupportsCursorColor           func(d *Display) bool
	DisplaySupportsInputShapes           func(d *Display) bool
	DisplaySupportsSelectionNotification func(d *Display) bool
	DisplaySupportsShapes                func(d *Display) bool
	DisplaySync                          func(d *Display)
	DisplayWarpPointer                   func(d *Display, screen *Screen, x, y int)
)

func (d *Display) AddClientMessageFilter(messageType Atom, f T.GdkFilterFunc, data T.Gpointer) {
	DisplayAddClientMessageFilter(d, messageType, f, data)
}
func (d *Display) Beep()                      { DisplayBeep(d) }
func (d *Display) Close()                     { DisplayClose(d) }
func (d *Display) Flush()                     { DisplayFlush(d) }
func (d *Display) GetCorePointer() *Device    { return DisplayGetCorePointer(d) }
func (d *Display) GetDefaultCursorSize() uint { return DisplayGetDefaultCursorSize(d) }
func (d *Display) GetDefaultGroup() *Window   { return DisplayGetDefaultGroup(d) }
func (d *Display) GetDefaultScreen() *Screen  { return DisplayGetDefaultScreen(d) }
func (d *Display) GetEvent() *Event           { return DisplayGetEvent(d) }
func (d *Display) GetMaximalCursorSize(width, height *uint) {
	DisplayGetMaximalCursorSize(d, width, height)
}
func (d *Display) GetName() string  { return DisplayGetName(d) }
func (d *Display) GetNScreens() int { return DisplayGetNScreens(d) }
func (d *Display) GetPointer(screen **Screen, x, y *int, mask *T.GdkModifierType) {
	DisplayGetPointer(d, screen, x, y, mask)
}
func (d *Display) GetScreen(screenNum int) *Screen { return DisplayGetScreen(d, screenNum) }
func (d *Display) GetWindowAtPointer(winX, winY *int) *Window {
	return DisplayGetWindowAtPointer(d, winX, winY)
}
func (d *Display) IsClosed() bool                { return DisplayIsClosed(d) }
func (d *Display) KeyboardUngrab(time T.GUint32) { DisplayKeyboardUngrab(d, time) }
func (d *Display) ListDevices() *L.List          { return DisplayListDevices(d) }
func (d *Display) PeekEvent() *Event             { return DisplayPeekEvent(d) }
func (d *Display) PointerIsGrabbed() bool        { return DisplayPointerIsGrabbed(d) }
func (d *Display) PointerUngrab(time T.GUint32)  { DisplayPointerUngrab(d, time) }
func (d *Display) PutEvent(event *Event)         { DisplayPutEvent(d, event) }
func (d *Display) RequestSelectionNotification(selection Atom) bool {
	return DisplayRequestSelectionNotification(d, selection)
}
func (d *Display) SetDoubleClickDistance(distance uint) { DisplaySetDoubleClickDistance(d, distance) }
func (d *Display) SetDoubleClickTime(msec uint)         { DisplaySetDoubleClickTime(d, msec) }
func (d *Display) SetPointerHooks(newHooks *DisplayPointerHooks) *DisplayPointerHooks {
	return DisplaySetPointerHooks(d, newHooks)
}
func (d *Display) StoreClipboard(clipboardWindow *Window, time T.GUint32, targets *Atom, nTargets int) {
	DisplayStoreClipboard(d, clipboardWindow, time, targets, nTargets)
}
func (d *Display) SupportsClipboardPersistence() bool {
	return DisplaySupportsClipboardPersistence(d)
}
func (d *Display) SupportsComposite() bool   { return DisplaySupportsComposite(d) }
func (d *Display) SupportsCursorAlpha() bool { return DisplaySupportsCursorAlpha(d) }
func (d *Display) SupportsCursorColor() bool { return DisplaySupportsCursorColor(d) }
func (d *Display) SupportsInputShapes() bool { return DisplaySupportsInputShapes(d) }
func (d *Display) SupportsSelectionNotification() bool {
	return DisplaySupportsSelectionNotification(d)
}
func (d *Display) SupportsShapes() bool                 { return DisplaySupportsShapes(d) }
func (d *Display) Sync()                                { DisplaySync(d) }
func (d *Display) WarpPointer(screen *Screen, x, y int) { DisplayWarpPointer(d, screen, x, y) }

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
	DisplayManagerGetType func() O.Type

	DisplayManagerGet func() *DisplayManager

	DisplayManagerGetDefaultDisplay func(d *DisplayManager) *Display
	DisplayManagerListDisplays      func(d *DisplayManager) *L.SList
	DisplayManagerSetDefaultDisplay func(d *DisplayManager, display *Display)
)

func (d *DisplayManager) GetDefaultDisplay() *Display { return DisplayManagerGetDefaultDisplay(d) }
func (d *DisplayManager) ListDisplays() *L.SList      { return DisplayManagerListDisplays(d) }
func (d *DisplayManager) SetDefaultDisplay(display *Display) {
	DisplayManagerSetDefaultDisplay(d, display)
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

var DragActionGetType func() O.Type

type DragContext struct {
	Parent          O.Object
	Protocol        DragProtocol
	Is_source       bool
	SourceWindow    *Window
	DestWindow      *Window
	Targets         *L.List
	Actions         DragAction
	SuggestedAction DragAction
	Action          DragAction
	StartTime       T.GUint32
	WindowingData   T.Gpointer
}

var (
	DragContextGetType func() O.Type
	DragContextNew     func() *DragContext

	DragAbort                 func(context *DragContext, time T.GUint32)
	DragBegin                 func(window *Window, targets *L.List) *DragContext
	DragGetProtocol           func(xid T.GdkNativeWindow, protocol *DragProtocol) T.GdkNativeWindow
	DragGetProtocolForDisplay func(display *Display, xid T.GdkNativeWindow, protocol *DragProtocol) T.GdkNativeWindow

	DragDrop                func(context *DragContext, time T.GUint32)
	DragDropSucceeded       func(context *DragContext) bool
	DragFindWindow          func(context *DragContext, dragWindow *Window, xRoot int, yRoot int, destWindow **Window, protocol *DragProtocol)
	DragFindWindowForScreen func(context *DragContext, dragWindow *Window, screen *Screen, xRoot int, yRoot int, destWindow **Window, protocol *DragProtocol)
	DragGetSelection        func(context *DragContext) Atom
	DragMotion              func(context *DragContext, destWindow *Window, protocol DragProtocol, xRoot int, yRoot int, suggestedAction DragAction, possibleActions DragAction, time T.GUint32) bool
	DragStatus              func(context *DragContext, action DragAction, time T.GUint32)
	DropFinish              func(context *DragContext, success bool, time T.GUint32)
	DropReply               func(context *DragContext, ok bool, time T.GUint32)

	DragContextGetActions         func(context *DragContext) DragAction
	DragContextGetDestWindow      func(context *DragContext) *Window
	DragContextGetProtocol        func(context *DragContext) DragProtocol
	DragContextGetSelectedAction  func(context *DragContext) DragAction
	DragContextGetSourceWindow    func(context *DragContext) *Window
	DragContextGetSuggestedAction func(context *DragContext) DragAction
	DragContextListTargets        func(context *DragContext) *L.List
	DragContextRef                func(context *DragContext)
	DragContextUnref              func(context *DragContext)
)

func (d *DragContext) GetActions() DragAction         { return DragContextGetActions(d) }
func (d *DragContext) GetDestWindow() *Window         { return DragContextGetDestWindow(d) }
func (d *DragContext) GetProtocol() DragProtocol      { return DragContextGetProtocol(d) }
func (d *DragContext) GetSelectedAction() DragAction  { return DragContextGetSelectedAction(d) }
func (d *DragContext) GetSourceWindow() *Window       { return DragContextGetSourceWindow(d) }
func (d *DragContext) GetSuggestedAction() DragAction { return DragContextGetSuggestedAction(d) }
func (d *DragContext) ListTargets() *L.List           { return DragContextListTargets(d) }
func (d *DragContext) Ref()                           { DragContextRef(d) }
func (d *DragContext) Unref()                         { DragContextUnref(d) }

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

var DragProtocolGetType func() O.Type

type Drawable struct{ parent O.Object }

var (
	DrawableGetType func() O.Type

	DrawArc                  func(drawable *Drawable, gc *GC, filled bool, x, y, width, height, angle1, angle2 int)
	DrawDrawable             func(drawable *Drawable, gc *GC, src *Drawable, xsrc, ysrc, xdest, ydest, width, height int)
	DrawGlyphs               func(drawable *Drawable, gc *GC, font *P.Font, x, y int, glyphs *P.GlyphString)
	DrawGlyphsTransformed    func(drawable *Drawable, gc *GC, matrix *P.Matrix, font *P.Font, x, y int, glyphs *P.GlyphString)
	DrawGrayImage            func(drawable *Drawable, gc *GC, x, y, width, height int, dith T.GdkRgbDither, buf *T.Guchar, rowstride int)
	DrawImage                func(drawable *Drawable, gc *GC, image *Image, xsrc, ysrc, xdest, ydest, width, height int)
	DrawIndexedImage         func(drawable *Drawable, gc *GC, x, y, width, height int, dith T.GdkRgbDither, buf *T.Guchar, rowstride int, cmap *T.GdkRgbCmap)
	DrawLayout               func(drawable *Drawable, gc *GC, x, y int, layout *P.Layout)
	DrawLayoutLine           func(drawable *Drawable, gc *GC, x, y int, line *P.LayoutLine)
	DrawLayoutLineWithColors func(drawable *Drawable, gc *GC, x, y int, line *P.LayoutLine, foreground, background *Color)
	DrawLayoutWithColors     func(drawable *Drawable, gc *GC, x, y int, layout *P.Layout, foreground, background *Color)
	DrawLine                 func(drawable *Drawable, gc *GC, x1, y1, x2, y2 int)
	DrawLines                func(drawable *Drawable, gc *GC, points *Point, nPoints int)
	DrawPixbuf               func(drawable *Drawable, gc *GC, pixbuf *Pixbuf, srcX, srcY, destX, destY, width, height int, dither T.GdkRgbDither, xDither, yDither int)
	DrawPoint                func(drawable *Drawable, gc *GC, x, y int)
	DrawPoints               func(drawable *Drawable, gc *GC, points *Point, nPoints int)
	DrawPolygon              func(drawable *Drawable, gc *GC, filled bool, points *Point, nPoints int)
	DrawRectangle            func(drawable *Drawable, gc *GC, filled bool, x, y, width, height int)
	DrawRgb32Image           func(drawable *Drawable, gc *GC, x, y, width, height int, dith T.GdkRgbDither, buf *T.Guchar, rowstride int)
	DrawRgb32ImageDithalign  func(drawable *Drawable, gc *GC, x, y, width, height int, dith T.GdkRgbDither, buf *T.Guchar, rowstride, xdith, ydith int)
	DrawRgbImage             func(drawable *Drawable, gc *GC, x, y, width, height int, dith T.GdkRgbDither, rgbBuf *T.Guchar, rowstride int)
	DrawRgbImageDithalign    func(drawable *Drawable, gc *GC, x, y, width, height int, dith T.GdkRgbDither, rgbBuf *T.Guchar, rowstride, xdith, ydith int)
	DrawSegments             func(drawable *Drawable, gc *GC, segs *T.GdkSegment, nSegs int)
	DrawString               func(drawable *Drawable, font *Font, gc *GC, x, y int, s string)
	DrawText                 func(drawable *Drawable, font *Font, gc *GC, x, y int, text string, textLength int)
	DrawTextWc               func(drawable *Drawable, font *Font, gc *GC, x, y int, text *T.GdkWChar, textLength int)
	DrawTrapezoids           func(drawable *Drawable, gc *GC, trapezoids *T.GdkTrapezoid, nTrapezoids int)

	DrawableCopyToImage      func(d *Drawable, image *Image, srcX, srcY, destX, destY, width, height int) *Image
	DrawableGetClipRegion    func(d *Drawable) *Region
	DrawableGetColormap      func(d *Drawable) *Colormap
	DrawableGetData          func(d *Drawable, key string) T.Gpointer
	DrawableGetDepth         func(d *Drawable) int
	DrawableGetDisplay       func(d *Drawable) *Display
	DrawableGetImage         func(d *Drawable, x, y, width, height int) *Image
	DrawableGetScreen        func(d *Drawable) *Screen
	DrawableGetSize          func(d *Drawable, width, height *int)
	DrawableGetVisibleRegion func(d *Drawable) *Region
	DrawableGetVisual        func(d *Drawable) *Visual
	DrawableRef              func(d *Drawable) *Drawable
	DrawableSetColormap      func(d *Drawable, colormap *Colormap)
	DrawableSetData          func(d *Drawable, key string, data T.Gpointer, destroyFunc O.DestroyNotify)
	DrawableUnref            func(d *Drawable)
)

func (d *Drawable) CopyToImage(image *Image, srcX, srcY, destX, destY, width, height int) *Image {
	return DrawableCopyToImage(d, image, srcX, srcY, destX, destY, width, height)
}
func (d *Drawable) GetClipRegion() *Region        { return DrawableGetClipRegion(d) }
func (d *Drawable) GetColormap() *Colormap        { return DrawableGetColormap(d) }
func (d *Drawable) GetData(key string) T.Gpointer { return DrawableGetData(d, key) }
func (d *Drawable) GetDepth() int                 { return DrawableGetDepth(d) }
func (d *Drawable) GetDisplay() *Display          { return DrawableGetDisplay(d) }
func (d *Drawable) GetImage(x, y, width, height int) *Image {
	return DrawableGetImage(d, x, y, width, height)
}
func (d *Drawable) GetScreen() *Screen             { return DrawableGetScreen(d) }
func (d *Drawable) GetSize(width, height *int)     { DrawableGetSize(d, width, height) }
func (d *Drawable) GetVisibleRegion() *Region      { return DrawableGetVisibleRegion(d) }
func (d *Drawable) GetVisual() *Visual             { return DrawableGetVisual(d) }
func (d *Drawable) Ref() *Drawable                 { return DrawableRef(d) }
func (d *Drawable) SetColormap(colormap *Colormap) { DrawableSetColormap(d, colormap) }
func (d *Drawable) SetData(key string, data T.Gpointer, destroyFunc O.DestroyNotify) {
	DrawableSetData(d, key, data, destroyFunc)
}
func (d *Drawable) Unref() { DrawableUnref(d) }
