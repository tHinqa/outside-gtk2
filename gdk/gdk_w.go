// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	C "github.com/tHinqa/outside-gtk2/cairo"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Window Drawable

var (
	WindowNew func(parent *Window, attributes *WindowAttr, attributesMask int) *Window

	WindowAtPointer         func(winX *int, winY *int) *Window
	WindowGetToplevels      func() *L.List
	WindowProcessAllUpdates func()

	WindowAddFilter                       func(w *Window, function T.GdkFilterFunc, data T.Gpointer)
	WindowBeep                            func(w *Window)
	WindowBeginMoveDrag                   func(w *Window, button, rootX, rootY int, timestamp T.GUint32)
	WindowBeginPaintRect                  func(w *Window, rectangle *Rectangle)
	WindowBeginPaintRegion                func(w *Window, region *Region)
	WindowBeginResizeDrag                 func(w *Window, edge WindowEdge, button, rootX, rootY int, timestamp T.GUint32)
	WindowClear                           func(w *Window)
	WindowClearArea                       func(w *Window, x, y, width, height int)
	WindowClearAreaE                      func(w *Window, x, y, width, height int)
	WindowConfigureFinished               func(w *Window)
	WindowCoordsFromParent                func(w *Window, parentX, parentY float64, x, y *float64)
	WindowCoordsToParent                  func(w *Window, x, y float64, parentX, parentY *float64)
	WindowCreateSimilarSurface            func(w *Window, content C.Content, width, height int) *C.Surface
	WindowDeiconify                       func(w *Window)
	WindowDestroy                         func(w *Window)
	WindowDestroyNotify                   func(w *Window)
	WindowEnableSynchronizedConfigure     func(w *Window)
	WindowEndPaint                        func(w *Window)
	WindowEnsureNative                    func(w *Window) bool
	WindowFlush                           func(w *Window)
	WindowFocus                           func(w *Window, timestamp T.GUint32)
	WindowFreezeToplevelUpdatesLibgtkOnly func(w *Window)
	WindowFreezeUpdates                   func(w *Window)
	WindowFullscreen                      func(w *Window)
	WindowGeometryChanged                 func(w *Window)
	WindowGetAcceptFocus                  func(w *Window) bool
	WindowGetBackgroundPattern            func(w *Window) *C.Pattern
	WindowGetChildren                     func(w *Window) *L.List
	WindowGetComposited                   func(w *Window) bool
	WindowGetCursor                       func(w *Window) *Cursor
	WindowGetDecorations                  func(w *Window, decorations *T.GdkWMDecoration) bool
	WindowGetDeskrelativeOrigin           func(w *Window, x, y *int) bool
	WindowGetDisplay                      func(w *Window) *Display
	WindowGetEffectiveParent              func(w *Window) *Window
	WindowGetEffectiveToplevel            func(w *Window) *Window
	WindowGetEvents                       func(w *Window) EventMask
	WindowGetFocusOnMap                   func(w *Window) bool
	WindowGetFrameExtents                 func(w *Window, rect *Rectangle)
	WindowGetGeometry                     func(w *Window, x, y, width, height, depth *int)
	WindowGetGroup                        func(w *Window) *Window
	WindowGetHeight                       func(w *Window) int
	WindowGetInternalPaintInfo            func(w *Window, realDrawable **Drawable, xOffset, yOffset *int)
	WindowGetModalHint                    func(w *Window) bool
	WindowGetOrigin                       func(w *Window, x, y *int) int
	WindowGetParent                       func(w *Window) *Window
	WindowGetPointer                      func(w *Window, x, y *int, mask *T.GdkModifierType) *Window
	WindowGetPosition                     func(w *Window, x, y *int)
	WindowGetRootCoords                   func(w *Window, x, y int, rootX, rootY *int)
	WindowGetRootOrigin                   func(w *Window, x, y *int)
	WindowGetScreen                       func(w *Window) *Screen
	WindowGetState                        func(w *Window) WindowState
	WindowGetToplevel                     func(w *Window) *Window
	WindowGetTypeHint                     func(w *Window) WindowTypeHint
	WindowGetUpdateArea                   func(w *Window) *Region
	WindowGetUserData                     func(w *Window, data *T.Gpointer)
	WindowGetVisual                       func(w *Window) *Visual
	WindowGetWidth                        func(w *Window) int
	WindowGetWindowType                   func(w *Window) WindowType
	WindowHasNative                       func(w *Window) bool
	WindowHide                            func(w *Window)
	WindowIconify                         func(w *Window)
	WindowInputShapeCombineMask           func(w *Window, mask *T.GdkBitmap, x, y int)
	WindowInputShapeCombineRegion         func(w *Window, shapeRegion *Region, offsetX, offsetY int)
	WindowInvalidateMaybeRecurse          func(w *Window, region *Region, childFunc func(*Window, T.Gpointer) bool, userData T.Gpointer)
	WindowInvalidateRect                  func(w *Window, rect *Rectangle, invalidateChildren bool)
	WindowInvalidateRegion                func(w *Window, region *Region, invalidateChildren bool)
	WindowIsDestroyed                     func(w *Window) bool
	WindowIsInputOnly                     func(w *Window) bool
	WindowIsShaped                        func(w *Window) bool
	WindowIsViewable                      func(w *Window) bool
	WindowIsVisible                       func(w *Window) bool
	WindowLower                           func(w *Window)
	WindowMaximize                        func(w *Window)
	WindowMergeChildInputShapes           func(w *Window)
	WindowMergeChildShapes                func(w *Window)
	WindowMove                            func(w *Window, x, y int)
	WindowMoveRegion                      func(w *Window, region *Region, dx, dy int)
	WindowMoveResize                      func(w *Window, x, y, width, height int)
	WindowPeekChildren                    func(w *Window) *L.List
	WindowProcessUpdates                  func(w *Window, updateChildren bool)
	WindowRaise                           func(w *Window)
	WindowRedirectToDrawable              func(w *Window, drawable *Drawable, srcX, srcY, destX, destY, width, height int)
	WindowRegisterDnd                     func(w *Window)
	WindowRemoveFilter                    func(w *Window, function T.GdkFilterFunc, data T.Gpointer)
	WindowRemoveRedirection               func(w *Window)
	WindowReparent                        func(w *Window, newParent *Window, x, y int)
	WindowResize                          func(w *Window, width, height int)
	WindowRestack                         func(w, sibling *Window, above bool)
	WindowScroll                          func(w *Window, dx, dy int)
	WindowSetAcceptFocus                  func(w *Window, acceptFocus bool)
	WindowSetBackground                   func(w *Window, color *Color)
	WindowSetBackPixmap                   func(w *Window, pixmap *Pixmap, parentRelative bool)
	WindowSetChildInputShapes             func(w *Window)
	WindowSetChildShapes                  func(w *Window)
	WindowSetComposited                   func(w *Window, composited bool)
	WindowSetCursor                       func(w *Window, cursor *Cursor)
	WindowSetDecorations                  func(w *Window, decorations T.GdkWMDecoration)
	WindowSetEvents                       func(w *Window, eventMask EventMask)
	WindowSetFocusOnMap                   func(w *Window, focusOnMap bool)
	WindowSetFunctions                    func(w *Window, functions T.GdkWMFunction)
	WindowSetGeometryHints                func(w *Window, geometry *Geometry, geomMask WindowHints)
	WindowSetGroup                        func(w *Window, leader *Window)
	WindowSetHints                        func(w *Window, x, y, minWidth, minHeight, maxWidth, maxHeight, flags int)
	WindowSetIcon                         func(w, iconWindow *Window, pixmap *Pixmap, mask *T.GdkBitmap)
	WindowSetIconList                     func(w *Window, pixbufs *L.List)
	WindowSetIconName                     func(w *Window, name string)
	WindowSetKeepAbove                    func(w *Window, setting bool)
	WindowSetKeepBelow                    func(w *Window, setting bool)
	WindowSetModalHint                    func(w *Window, modal bool)
	WindowSetOpacity                      func(w *Window, opacity float64)
	WindowSetOverrideRedirect             func(w *Window, overrideRedirect bool)
	WindowSetRole                         func(w *Window, role string)
	WindowSetSkipPagerHint                func(w *Window, skipsPager bool)
	WindowSetSkipTaskbarHint              func(w *Window, skipsTaskbar bool)
	WindowSetStartupId                    func(w *Window, startupId string)
	WindowSetStaticGravities              func(w *Window, useStatic bool) bool
	WindowSetTitle                        func(w *Window, title string)
	WindowSetTransientFor                 func(w *Window, parent *Window)
	WindowSetTypeHint                     func(w *Window, hint WindowTypeHint)
	WindowSetUrgencyHint                  func(w *Window, urgent bool)
	WindowSetUserData                     func(w *Window, userData T.Gpointer)
	WindowShapeCombineMask                func(w *Window, mask *T.GdkBitmap, x, y int)
	WindowShapeCombineRegion              func(w *Window, shapeRegion *Region, offsetX, offsetY int)
	WindowShow                            func(w *Window)
	WindowShowUnraised                    func(w *Window)
	WindowStick                           func(w *Window)
	WindowThawToplevelUpdatesLibgtkOnly   func(w *Window)
	WindowThawUpdates                     func(w *Window)
	WindowUnfullscreen                    func(w *Window)
	WindowUnmaximize                      func(w *Window)
	WindowUnstick                         func(w *Window)
	WindowWithdraw                        func(w *Window)
)

func (w *Window) AddFilter(function T.GdkFilterFunc, data T.Gpointer) {
	WindowAddFilter(w, function, data)
}
func (w *Window) Beep() { WindowBeep(w) }
func (w *Window) BeginMoveDrag(button, rootX, rootY int, timestamp T.GUint32) {
	WindowBeginMoveDrag(w, button, rootX, rootY, timestamp)
}
func (w *Window) BeginPaintRect(rectangle *Rectangle) { WindowBeginPaintRect(w, rectangle) }
func (w *Window) BeginPaintRegion(region *Region)     { WindowBeginPaintRegion(w, region) }
func (w *Window) BeginResizeDrag(edge WindowEdge, button, rootX, rootY int, timestamp T.GUint32) {
	WindowBeginResizeDrag(w, edge, button, rootX, rootY, timestamp)
}
func (w *Window) Clear()                             { WindowClear(w) }
func (w *Window) ClearArea(x, y, width, height int)  { WindowClearArea(w, x, y, width, height) }
func (w *Window) ClearAreaE(x, y, width, height int) { WindowClearAreaE(w, x, y, width, height) }
func (w *Window) ConfigureFinished()                 { WindowConfigureFinished(w) }
func (w *Window) CoordsFromParent(parentX, parentY float64, x, y *float64) {
	WindowCoordsFromParent(w, parentX, parentY, x, y)
}
func (w *Window) CoordsToParent(x, y float64, parentX, parentY *float64) {
	WindowCoordsToParent(w, x, y, parentX, parentY)
}
func (w *Window) CreateSimilarSurface(content C.Content, width, height int) *C.Surface {
	return WindowCreateSimilarSurface(w, content, width, height)
}
func (w *Window) Deiconify()                       { WindowDeiconify(w) }
func (w *Window) Destroy()                         { WindowDestroy(w) }
func (w *Window) DestroyNotify()                   { WindowDestroyNotify(w) }
func (w *Window) EnableSynchronizedConfigure()     { WindowEnableSynchronizedConfigure(w) }
func (w *Window) EndPaint()                        { WindowEndPaint(w) }
func (w *Window) EnsureNative() bool               { return WindowEnsureNative(w) }
func (w *Window) Flush()                           { WindowFlush(w) }
func (w *Window) Focus(timestamp T.GUint32)        { WindowFocus(w, timestamp) }
func (w *Window) FreezeToplevelUpdatesLibgtkOnly() { WindowFreezeToplevelUpdatesLibgtkOnly(w) }
func (w *Window) FreezeUpdates()                   { WindowFreezeUpdates(w) }
func (w *Window) Fullscreen()                      { WindowFullscreen(w) }
func (w *Window) GeometryChanged()                 { WindowGeometryChanged(w) }
func (w *Window) GetAcceptFocus() bool             { return WindowGetAcceptFocus(w) }
func (w *Window) GetBackgroundPattern() *C.Pattern { return WindowGetBackgroundPattern(w) }
func (w *Window) GetChildren() *L.List             { return WindowGetChildren(w) }
func (w *Window) GetComposited() bool              { return WindowGetComposited(w) }
func (w *Window) GetCursor() *Cursor               { return WindowGetCursor(w) }
func (w *Window) GetDecorations(decorations *T.GdkWMDecoration) bool {
	return WindowGetDecorations(w, decorations)
}
func (w *Window) GetDeskrelativeOrigin(x, y *int) bool {
	return WindowGetDeskrelativeOrigin(w, x, y)
}
func (w *Window) GetDisplay() *Display            { return WindowGetDisplay(w) }
func (w *Window) GetEffectiveParent() *Window     { return WindowGetEffectiveParent(w) }
func (w *Window) GetEffectiveToplevel() *Window   { return WindowGetEffectiveToplevel(w) }
func (w *Window) GetEvents() EventMask            { return WindowGetEvents(w) }
func (w *Window) GetFocusOnMap() bool             { return WindowGetFocusOnMap(w) }
func (w *Window) GetFrameExtents(rect *Rectangle) { WindowGetFrameExtents(w, rect) }
func (w *Window) GetGeometry(x, y, width, height, depth *int) {
	WindowGetGeometry(w, x, y, width, height, depth)
}
func (w *Window) GetGroup() *Window { return WindowGetGroup(w) }
func (w *Window) GetHeight() int    { return WindowGetHeight(w) }
func (w *Window) GetInternalPaintInfo(realDrawable **Drawable, xOffset, yOffset *int) {
	WindowGetInternalPaintInfo(w, realDrawable, xOffset, yOffset)
}
func (w *Window) GetModalHint() bool      { return WindowGetModalHint(w) }
func (w *Window) GetOrigin(x, y *int) int { return WindowGetOrigin(w, x, y) }
func (w *Window) GetParent() *Window      { return WindowGetParent(w) }
func (w *Window) GetPointer(x, y *int, mask *T.GdkModifierType) *Window {
	return WindowGetPointer(w, x, y, mask)
}
func (w *Window) GetPosition(x, y *int) { WindowGetPosition(w, x, y) }
func (w *Window) GetRootCoords(x, y int, rootX, rootY *int) {
	WindowGetRootCoords(w, x, y, rootX, rootY)
}
func (w *Window) GetRootOrigin(x, y *int)      { WindowGetRootOrigin(w, x, y) }
func (w *Window) GetScreen() *Screen           { return WindowGetScreen(w) }
func (w *Window) GetState() WindowState        { return WindowGetState(w) }
func (w *Window) GetToplevel() *Window         { return WindowGetToplevel(w) }
func (w *Window) GetTypeHint() WindowTypeHint  { return WindowGetTypeHint(w) }
func (w *Window) GetUpdateArea() *Region       { return WindowGetUpdateArea(w) }
func (w *Window) GetUserData(data *T.Gpointer) { WindowGetUserData(w, data) }
func (w *Window) GetVisual() *Visual           { return WindowGetVisual(w) }
func (w *Window) GetWidth() int                { return WindowGetWidth(w) }
func (w *Window) GetWindowType() WindowType    { return WindowGetWindowType(w) }
func (w *Window) HasNative() bool              { return WindowHasNative(w) }
func (w *Window) Hide()                        { WindowHide(w) }
func (w *Window) Iconify()                     { WindowIconify(w) }
func (w *Window) InputShapeCombineMask(mask *T.GdkBitmap, x, y int) {
	WindowInputShapeCombineMask(w, mask, x, y)
}
func (w *Window) InputShapeCombineRegion(shapeRegion *Region, offsetX, offsetY int) {
	WindowInputShapeCombineRegion(w, shapeRegion, offsetX, offsetY)
}
func (w *Window) InvalidateMaybeRecurse(region *Region, childFunc func(*Window, T.Gpointer) bool, userData T.Gpointer) {
	WindowInvalidateMaybeRecurse(w, region, childFunc, userData)
}
func (w *Window) InvalidateRect(rect *Rectangle, invalidateChildren bool) {
	WindowInvalidateRect(w, rect, invalidateChildren)
}
func (w *Window) InvalidateRegion(region *Region, invalidateChildren bool) {
	WindowInvalidateRegion(w, region, invalidateChildren)
}
func (w *Window) IsDestroyed() bool                     { return WindowIsDestroyed(w) }
func (w *Window) IsInputOnly() bool                     { return WindowIsInputOnly(w) }
func (w *Window) IsShaped() bool                        { return WindowIsShaped(w) }
func (w *Window) IsViewable() bool                      { return WindowIsViewable(w) }
func (w *Window) IsVisible() bool                       { return WindowIsVisible(w) }
func (w *Window) Lower()                                { WindowLower(w) }
func (w *Window) Maximize()                             { WindowMaximize(w) }
func (w *Window) MergeChildInputShapes()                { WindowMergeChildInputShapes(w) }
func (w *Window) MergeChildShapes()                     { WindowMergeChildShapes(w) }
func (w *Window) Move(x, y int)                         { WindowMove(w, x, y) }
func (w *Window) MoveRegion(region *Region, dx, dy int) { WindowMoveRegion(w, region, dx, dy) }
func (w *Window) MoveResize(x, y, width, height int)    { WindowMoveResize(w, x, y, width, height) }
func (w *Window) PeekChildren() *L.List                 { return WindowPeekChildren(w) }
func (w *Window) ProcessUpdates(updateChildren bool)    { WindowProcessUpdates(w, updateChildren) }
func (w *Window) Raise()                                { WindowRaise(w) }
func (w *Window) RedirectToDrawable(drawable *Drawable, srcX, srcY, destX, destY, width, height int) {
	WindowRedirectToDrawable(w, drawable, srcX, srcY, destX, destY, width, height)
}
func (w *Window) RegisterDnd() { WindowRegisterDnd(w) }
func (w *Window) RemoveFilter(function T.GdkFilterFunc, data T.Gpointer) {
	WindowRemoveFilter(w, function, data)
}
func (w *Window) RemoveRedirection()                   { WindowRemoveRedirection(w) }
func (w *Window) Reparent(newParent *Window, x, y int) { WindowReparent(w, newParent, x, y) }
func (w *Window) Resize(width, height int)             { WindowResize(w, width, height) }
func (w *Window) Restack(sibling *Window, above bool)  { WindowRestack(w, sibling, above) }
func (w *Window) Scroll(dx, dy int)                    { WindowScroll(w, dx, dy) }
func (w *Window) SetAcceptFocus(acceptFocus bool)      { WindowSetAcceptFocus(w, acceptFocus) }
func (w *Window) SetBackground(color *Color)           { WindowSetBackground(w, color) }
func (w *Window) SetBackPixmap(pixmap *Pixmap, parentRelative bool) {
	WindowSetBackPixmap(w, pixmap, parentRelative)
}
func (w *Window) SetChildInputShapes()                         { WindowSetChildInputShapes(w) }
func (w *Window) SetChildShapes()                              { WindowSetChildShapes(w) }
func (w *Window) SetComposited(composited bool)                { WindowSetComposited(w, composited) }
func (w *Window) SetCursor(cursor *Cursor)                     { WindowSetCursor(w, cursor) }
func (w *Window) SetDecorations(decorations T.GdkWMDecoration) { WindowSetDecorations(w, decorations) }
func (w *Window) SetEvents(eventMask EventMask)                { WindowSetEvents(w, eventMask) }
func (w *Window) SetFocusOnMap(focusOnMap bool)                { WindowSetFocusOnMap(w, focusOnMap) }
func (w *Window) SetFunctions(functions T.GdkWMFunction)       { WindowSetFunctions(w, functions) }
func (w *Window) SetGeometryHints(geometry *Geometry, geomMask WindowHints) {
	WindowSetGeometryHints(w, geometry, geomMask)
}
func (w *Window) SetGroup(leader *Window) { WindowSetGroup(w, leader) }
func (w *Window) SetHints(x, y, minWidth, minHeight, maxWidth, maxHeight, flags int) {
	WindowSetHints(w, x, y, minWidth, minHeight, maxWidth, maxHeight, flags)
}
func (w *Window) SetIcon(iconWindow *Window, pixmap *Pixmap, mask *T.GdkBitmap) {
	WindowSetIcon(w, iconWindow, pixmap, mask)
}
func (w *Window) SetIconList(pixbufs *L.List) { WindowSetIconList(w, pixbufs) }
func (w *Window) SetIconName(name string)     { WindowSetIconName(w, name) }
func (w *Window) SetKeepAbove(setting bool)   { WindowSetKeepAbove(w, setting) }
func (w *Window) SetKeepBelow(setting bool)   { WindowSetKeepBelow(w, setting) }
func (w *Window) SetModalHint(modal bool)     { WindowSetModalHint(w, modal) }
func (w *Window) SetOpacity(opacity float64)  { WindowSetOpacity(w, opacity) }
func (w *Window) SetOverrideRedirect(overrideRedirect bool) {
	WindowSetOverrideRedirect(w, overrideRedirect)
}
func (w *Window) SetRole(role string)              { WindowSetRole(w, role) }
func (w *Window) SetSkipPagerHint(skipsPager bool) { WindowSetSkipPagerHint(w, skipsPager) }
func (w *Window) SetSkipTaskbarHint(skipsTaskbar bool) {
	WindowSetSkipTaskbarHint(w, skipsTaskbar)
}
func (w *Window) SetStartupId(startupId string) { WindowSetStartupId(w, startupId) }
func (w *Window) SetStaticGravities(useStatic bool) bool {
	return WindowSetStaticGravities(w, useStatic)
}
func (w *Window) SetTitle(title string)                        { WindowSetTitle(w, title) }
func (w *Window) SetTransientFor(parent *Window)               { WindowSetTransientFor(w, parent) }
func (w *Window) SetTypeHint(hint WindowTypeHint)              { WindowSetTypeHint(w, hint) }
func (w *Window) SetUrgencyHint(urgent bool)                   { WindowSetUrgencyHint(w, urgent) }
func (w *Window) SetUserData(userData T.Gpointer)              { WindowSetUserData(w, userData) }
func (w *Window) ShapeCombineMask(mask *T.GdkBitmap, x, y int) { WindowShapeCombineMask(w, mask, x, y) }
func (w *Window) ShapeCombineRegion(shapeRegion *Region, offsetX, offsetY int) {
	WindowShapeCombineRegion(w, shapeRegion, offsetX, offsetY)
}
func (w *Window) Show()                          { WindowShow(w) }
func (w *Window) ShowUnraised()                  { WindowShowUnraised(w) }
func (w *Window) Stick()                         { WindowStick(w) }
func (w *Window) ThawToplevelUpdatesLibgtkOnly() { WindowThawToplevelUpdatesLibgtkOnly(w) }
func (w *Window) ThawUpdates()                   { WindowThawUpdates(w) }
func (w *Window) Unfullscreen()                  { WindowUnfullscreen(w) }
func (w *Window) Unmaximize()                    { WindowUnmaximize(w) }
func (w *Window) Unstick()                       { WindowUnstick(w) }
func (w *Window) Withdraw()                      { WindowWithdraw(w) }

type WindowAttr struct {
	Title            *T.Gchar
	EventMask        int
	X, Y             int
	Width            int
	Height           int
	Wclass           WindowClass
	Visual           *Visual
	Colormap         *Colormap
	WindowType       WindowType
	Cursor           *Cursor
	WmclassName      *T.Gchar
	WmclassClass     *T.Gchar
	OverrideRedirect T.Gboolean
	TypeHint         WindowTypeHint
}

type WindowClass Enum

const (
	INPUT_OUTPUT WindowClass = iota
	INPUT_ONLY
)

var WindowClassGetType func() O.Type

type WindowEdge Enum

const (
	WINDOW_EDGE_NORTH_WEST WindowEdge = iota
	WINDOW_EDGE_NORTH
	WINDOW_EDGE_NORTH_EAST
	WINDOW_EDGE_WEST
	WINDOW_EDGE_EAST
	WINDOW_EDGE_SOUTH_WEST
	WINDOW_EDGE_SOUTH
	WINDOW_EDGE_SOUTH_EAST
)

var WindowEdgeGetType func() O.Type

type WindowHints Enum

const (
	HINT_POS WindowHints = 1 << iota
	HINT_MIN_SIZE
	HINT_MAX_SIZE
	HINT_BASE_SIZE
	HINT_ASPECT
	HINT_RESIZE_INC
	HINT_WIN_GRAVITY
	HINT_USER_POS
	HINT_USER_SIZE
)

var WindowHintsGetType func() O.Type

type WindowState Enum

const (
	WINDOW_STATE_WITHDRAWN WindowState = 1 << iota
	WINDOW_STATE_ICONIFIED
	WINDOW_STATE_MAXIMIZED
	WINDOW_STATE_STICKY
	WINDOW_STATE_FULLSCREEN
	WINDOW_STATE_ABOVE
	WINDOW_STATE_BELOW
)

var WindowStateGetType func() O.Type

type WindowType Enum

const (
	WINDOW_ROOT WindowType = iota
	WINDOW_TOPLEVEL
	WINDOW_CHILD
	WINDOW_DIALOG
	WINDOW_TEMP
	WINDOW_FOREIGN
	WINDOW_OFFSCREEN
)

var WindowTypeGetType func() O.Type

type WindowTypeHint Enum

const (
	WINDOW_TYPE_HINT_NORMAL WindowTypeHint = iota
	WINDOW_TYPE_HINT_DIALOG
	WINDOW_TYPE_HINT_MENU
	WINDOW_TYPE_HINT_TOOLBAR
	WINDOW_TYPE_HINT_SPLASHSCREEN
	WINDOW_TYPE_HINT_UTILITY
	WINDOW_TYPE_HINT_DOCK
	WINDOW_TYPE_HINT_DESKTOP
	WINDOW_TYPE_HINT_DROPDOWN_MENU
	WINDOW_TYPE_HINT_POPUP_MENU
	WINDOW_TYPE_HINT_TOOLTIP
	WINDOW_TYPE_HINT_NOTIFICATION
	WINDOW_TYPE_HINT_COMBO
	WINDOW_TYPE_HINT_DND
)

var WindowTypeHintGetType func() O.Type
