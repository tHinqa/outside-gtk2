// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	C "github.com/tHinqa/outside-gtk2/cairo"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Window Drawable

var (
	WindowNew func(parent *Window, attributes *WindowAttr, attributesMask int) *Window

	WindowAtPointer         func(winX *int, winY *int) *Window
	WindowGetToplevels      func() *T.GList
	WindowProcessAllUpdates func()

	windowAddFilter                       func(w *Window, function T.GdkFilterFunc, data T.Gpointer)
	windowBeep                            func(w *Window)
	windowBeginMoveDrag                   func(w *Window, button, rootX, rootY int, timestamp T.GUint32)
	windowBeginPaintRect                  func(w *Window, rectangle *Rectangle)
	windowBeginPaintRegion                func(w *Window, region *Region)
	windowBeginResizeDrag                 func(w *Window, edge WindowEdge, button, rootX, rootY int, timestamp T.GUint32)
	windowClear                           func(w *Window)
	windowClearArea                       func(w *Window, x, y, width, height int)
	windowClearAreaE                      func(w *Window, x, y, width, height int)
	windowConfigureFinished               func(w *Window)
	windowCoordsFromParent                func(w *Window, parentX, parentY float64, x, y *float64)
	windowCoordsToParent                  func(w *Window, x, y float64, parentX, parentY *float64)
	windowCreateSimilarSurface            func(w *Window, content C.Content, width, height int) *C.Surface
	windowDeiconify                       func(w *Window)
	windowDestroy                         func(w *Window)
	windowDestroyNotify                   func(w *Window)
	windowEnableSynchronizedConfigure     func(w *Window)
	windowEndPaint                        func(w *Window)
	windowEnsureNative                    func(w *Window) T.Gboolean
	windowFlush                           func(w *Window)
	windowFocus                           func(w *Window, timestamp T.GUint32)
	windowFreezeToplevelUpdatesLibgtkOnly func(w *Window)
	windowFreezeUpdates                   func(w *Window)
	windowFullscreen                      func(w *Window)
	windowGeometryChanged                 func(w *Window)
	windowGetAcceptFocus                  func(w *Window) T.Gboolean
	windowGetBackgroundPattern            func(w *Window) *C.Pattern
	windowGetChildren                     func(w *Window) *T.GList
	windowGetComposited                   func(w *Window) T.Gboolean
	windowGetCursor                       func(w *Window) *Cursor
	windowGetDecorations                  func(w *Window, decorations *T.GdkWMDecoration) T.Gboolean
	windowGetDeskrelativeOrigin           func(w *Window, x, y *int) T.Gboolean
	windowGetDisplay                      func(w *Window) *Display
	windowGetEffectiveParent              func(w *Window) *Window
	windowGetEffectiveToplevel            func(w *Window) *Window
	windowGetEvents                       func(w *Window) EventMask
	windowGetFocusOnMap                   func(w *Window) T.Gboolean
	windowGetFrameExtents                 func(w *Window, rect *Rectangle)
	windowGetGeometry                     func(w *Window, x, y, width, height, depth *int)
	windowGetGroup                        func(w *Window) *Window
	windowGetHeight                       func(w *Window) int
	windowGetInternalPaintInfo            func(w *Window, realDrawable **Drawable, xOffset, yOffset *int)
	windowGetModalHint                    func(w *Window) T.Gboolean
	windowGetOrigin                       func(w *Window, x, y *int) int
	windowGetParent                       func(w *Window) *Window
	windowGetPointer                      func(w *Window, x, y *int, mask *T.GdkModifierType) *Window
	windowGetPosition                     func(w *Window, x, y *int)
	windowGetRootCoords                   func(w *Window, x, y int, rootX, rootY *int)
	windowGetRootOrigin                   func(w *Window, x, y *int)
	windowGetScreen                       func(w *Window) *Screen
	windowGetState                        func(w *Window) WindowState
	windowGetToplevel                     func(w *Window) *Window
	windowGetTypeHint                     func(w *Window) WindowTypeHint
	windowGetUpdateArea                   func(w *Window) *Region
	windowGetUserData                     func(w *Window, data *T.Gpointer)
	windowGetVisual                       func(w *Window) *Visual
	windowGetWidth                        func(w *Window) int
	windowGetWindowType                   func(w *Window) WindowType
	windowHasNative                       func(w *Window) T.Gboolean
	windowHide                            func(w *Window)
	windowIconify                         func(w *Window)
	windowInputShapeCombineMask           func(w *Window, mask *T.GdkBitmap, x, y int)
	windowInputShapeCombineRegion         func(w *Window, shapeRegion *Region, offsetX, offsetY int)
	windowInvalidateMaybeRecurse          func(w *Window, region *Region, childFunc func(*Window, T.Gpointer) T.Gboolean, userData T.Gpointer)
	windowInvalidateRect                  func(w *Window, rect *Rectangle, invalidateChildren T.Gboolean)
	windowInvalidateRegion                func(w *Window, region *Region, invalidateChildren T.Gboolean)
	windowIsDestroyed                     func(w *Window) T.Gboolean
	windowIsInputOnly                     func(w *Window) T.Gboolean
	windowIsShaped                        func(w *Window) T.Gboolean
	windowIsViewable                      func(w *Window) T.Gboolean
	windowIsVisible                       func(w *Window) T.Gboolean
	windowLower                           func(w *Window)
	windowMaximize                        func(w *Window)
	windowMergeChildInputShapes           func(w *Window)
	windowMergeChildShapes                func(w *Window)
	windowMove                            func(w *Window, x, y int)
	windowMoveRegion                      func(w *Window, region *Region, dx, dy int)
	windowMoveResize                      func(w *Window, x, y, width, height int)
	windowPeekChildren                    func(w *Window) *T.GList
	windowProcessUpdates                  func(w *Window, updateChildren T.Gboolean)
	windowRaise                           func(w *Window)
	windowRedirectToDrawable              func(w *Window, drawable *Drawable, srcX, srcY, destX, destY, width, height int)
	windowRegisterDnd                     func(w *Window)
	windowRemoveFilter                    func(w *Window, function T.GdkFilterFunc, data T.Gpointer)
	windowRemoveRedirection               func(w *Window)
	windowReparent                        func(w *Window, newParent *Window, x, y int)
	windowResize                          func(w *Window, width, height int)
	windowRestack                         func(w, sibling *Window, above T.Gboolean)
	windowScroll                          func(w *Window, dx, dy int)
	windowSetAcceptFocus                  func(w *Window, acceptFocus T.Gboolean)
	windowSetBackground                   func(w *Window, color *Color)
	windowSetBackPixmap                   func(w *Window, pixmap *Pixmap, parentRelative T.Gboolean)
	windowSetChildInputShapes             func(w *Window)
	windowSetChildShapes                  func(w *Window)
	windowSetComposited                   func(w *Window, composited T.Gboolean)
	windowSetCursor                       func(w *Window, cursor *Cursor)
	windowSetDecorations                  func(w *Window, decorations T.GdkWMDecoration)
	windowSetEvents                       func(w *Window, eventMask EventMask)
	windowSetFocusOnMap                   func(w *Window, focusOnMap T.Gboolean)
	windowSetFunctions                    func(w *Window, functions T.GdkWMFunction)
	windowSetGeometryHints                func(w *Window, geometry *Geometry, geomMask WindowHints)
	windowSetGroup                        func(w *Window, leader *Window)
	windowSetHints                        func(w *Window, x, y, minWidth, minHeight, maxWidth, maxHeight, flags int)
	windowSetIcon                         func(w, iconWindow *Window, pixmap *Pixmap, mask *T.GdkBitmap)
	windowSetIconList                     func(w *Window, pixbufs *T.GList)
	windowSetIconName                     func(w *Window, name string)
	windowSetKeepAbove                    func(w *Window, setting T.Gboolean)
	windowSetKeepBelow                    func(w *Window, setting T.Gboolean)
	windowSetModalHint                    func(w *Window, modal T.Gboolean)
	windowSetOpacity                      func(w *Window, opacity float64)
	windowSetOverrideRedirect             func(w *Window, overrideRedirect T.Gboolean)
	windowSetRole                         func(w *Window, role string)
	windowSetSkipPagerHint                func(w *Window, skipsPager T.Gboolean)
	windowSetSkipTaskbarHint              func(w *Window, skipsTaskbar T.Gboolean)
	windowSetStartupId                    func(w *Window, startupId string)
	windowSetStaticGravities              func(w *Window, useStatic T.Gboolean) T.Gboolean
	windowSetTitle                        func(w *Window, title string)
	windowSetTransientFor                 func(w *Window, parent *Window)
	windowSetTypeHint                     func(w *Window, hint WindowTypeHint)
	windowSetUrgencyHint                  func(w *Window, urgent T.Gboolean)
	windowSetUserData                     func(w *Window, userData T.Gpointer)
	windowShapeCombineMask                func(w *Window, mask *T.GdkBitmap, x, y int)
	windowShapeCombineRegion              func(w *Window, shapeRegion *Region, offsetX, offsetY int)
	windowShow                            func(w *Window)
	windowShowUnraised                    func(w *Window)
	windowStick                           func(w *Window)
	windowThawToplevelUpdatesLibgtkOnly   func(w *Window)
	windowThawUpdates                     func(w *Window)
	windowUnfullscreen                    func(w *Window)
	windowUnmaximize                      func(w *Window)
	windowUnstick                         func(w *Window)
	windowWithdraw                        func(w *Window)
)

func (w *Window) AddFilter(function T.GdkFilterFunc, data T.Gpointer) {
	windowAddFilter(w, function, data)
}
func (w *Window) Beep() { windowBeep(w) }
func (w *Window) BeginMoveDrag(button, rootX, rootY int, timestamp T.GUint32) {
	windowBeginMoveDrag(w, button, rootX, rootY, timestamp)
}
func (w *Window) BeginPaintRect(rectangle *Rectangle) { windowBeginPaintRect(w, rectangle) }
func (w *Window) BeginPaintRegion(region *Region)     { windowBeginPaintRegion(w, region) }
func (w *Window) BeginResizeDrag(edge WindowEdge, button, rootX, rootY int, timestamp T.GUint32) {
	windowBeginResizeDrag(w, edge, button, rootX, rootY, timestamp)
}
func (w *Window) Clear()                             { windowClear(w) }
func (w *Window) ClearArea(x, y, width, height int)  { windowClearArea(w, x, y, width, height) }
func (w *Window) ClearAreaE(x, y, width, height int) { windowClearAreaE(w, x, y, width, height) }
func (w *Window) ConfigureFinished()                 { windowConfigureFinished(w) }
func (w *Window) CoordsFromParent(parentX, parentY float64, x, y *float64) {
	windowCoordsFromParent(w, parentX, parentY, x, y)
}
func (w *Window) CoordsToParent(x, y float64, parentX, parentY *float64) {
	windowCoordsToParent(w, x, y, parentX, parentY)
}
func (w *Window) CreateSimilarSurface(content C.Content, width, height int) *C.Surface {
	return windowCreateSimilarSurface(w, content, width, height)
}
func (w *Window) Deiconify()                       { windowDeiconify(w) }
func (w *Window) Destroy()                         { windowDestroy(w) }
func (w *Window) DestroyNotify()                   { windowDestroyNotify(w) }
func (w *Window) EnableSynchronizedConfigure()     { windowEnableSynchronizedConfigure(w) }
func (w *Window) EndPaint()                        { windowEndPaint(w) }
func (w *Window) EnsureNative() T.Gboolean         { return windowEnsureNative(w) }
func (w *Window) Flush()                           { windowFlush(w) }
func (w *Window) Focus(timestamp T.GUint32)        { windowFocus(w, timestamp) }
func (w *Window) FreezeToplevelUpdatesLibgtkOnly() { windowFreezeToplevelUpdatesLibgtkOnly(w) }
func (w *Window) FreezeUpdates()                   { windowFreezeUpdates(w) }
func (w *Window) Fullscreen()                      { windowFullscreen(w) }
func (w *Window) GeometryChanged()                 { windowGeometryChanged(w) }
func (w *Window) GetAcceptFocus() T.Gboolean       { return windowGetAcceptFocus(w) }
func (w *Window) GetBackgroundPattern() *C.Pattern { return windowGetBackgroundPattern(w) }
func (w *Window) GetChildren() *T.GList            { return windowGetChildren(w) }
func (w *Window) GetComposited() T.Gboolean        { return windowGetComposited(w) }
func (w *Window) GetCursor() *Cursor               { return windowGetCursor(w) }
func (w *Window) GetDecorations(decorations *T.GdkWMDecoration) T.Gboolean {
	return windowGetDecorations(w, decorations)
}
func (w *Window) GetDeskrelativeOrigin(x, y *int) T.Gboolean {
	return windowGetDeskrelativeOrigin(w, x, y)
}
func (w *Window) GetDisplay() *Display            { return windowGetDisplay(w) }
func (w *Window) GetEffectiveParent() *Window     { return windowGetEffectiveParent(w) }
func (w *Window) GetEffectiveToplevel() *Window   { return windowGetEffectiveToplevel(w) }
func (w *Window) GetEvents() EventMask            { return windowGetEvents(w) }
func (w *Window) GetFocusOnMap() T.Gboolean       { return windowGetFocusOnMap(w) }
func (w *Window) GetFrameExtents(rect *Rectangle) { windowGetFrameExtents(w, rect) }
func (w *Window) GetGeometry(x, y, width, height, depth *int) {
	windowGetGeometry(w, x, y, width, height, depth)
}
func (w *Window) GetGroup() *Window { return windowGetGroup(w) }
func (w *Window) GetHeight() int    { return windowGetHeight(w) }
func (w *Window) GetInternalPaintInfo(realDrawable **Drawable, xOffset, yOffset *int) {
	windowGetInternalPaintInfo(w, realDrawable, xOffset, yOffset)
}
func (w *Window) GetModalHint() T.Gboolean { return windowGetModalHint(w) }
func (w *Window) GetOrigin(x, y *int) int  { return windowGetOrigin(w, x, y) }
func (w *Window) GetParent() *Window       { return windowGetParent(w) }
func (w *Window) GetPointer(x, y *int, mask *T.GdkModifierType) *Window {
	return windowGetPointer(w, x, y, mask)
}
func (w *Window) GetPosition(x, y *int) { windowGetPosition(w, x, y) }
func (w *Window) GetRootCoords(x, y int, rootX, rootY *int) {
	windowGetRootCoords(w, x, y, rootX, rootY)
}
func (w *Window) GetRootOrigin(x, y *int)      { windowGetRootOrigin(w, x, y) }
func (w *Window) GetScreen() *Screen           { return windowGetScreen(w) }
func (w *Window) GetState() WindowState        { return windowGetState(w) }
func (w *Window) GetToplevel() *Window         { return windowGetToplevel(w) }
func (w *Window) GetTypeHint() WindowTypeHint  { return windowGetTypeHint(w) }
func (w *Window) GetUpdateArea() *Region       { return windowGetUpdateArea(w) }
func (w *Window) GetUserData(data *T.Gpointer) { windowGetUserData(w, data) }
func (w *Window) GetVisual() *Visual           { return windowGetVisual(w) }
func (w *Window) GetWidth() int                { return windowGetWidth(w) }
func (w *Window) GetWindowType() WindowType    { return windowGetWindowType(w) }
func (w *Window) HasNative() T.Gboolean        { return windowHasNative(w) }
func (w *Window) Hide()                        { windowHide(w) }
func (w *Window) Iconify()                     { windowIconify(w) }
func (w *Window) InputShapeCombineMask(mask *T.GdkBitmap, x, y int) {
	windowInputShapeCombineMask(w, mask, x, y)
}
func (w *Window) InputShapeCombineRegion(shapeRegion *Region, offsetX, offsetY int) {
	windowInputShapeCombineRegion(w, shapeRegion, offsetX, offsetY)
}
func (w *Window) InvalidateMaybeRecurse(region *Region, childFunc func(*Window, T.Gpointer) T.Gboolean, userData T.Gpointer) {
	windowInvalidateMaybeRecurse(w, region, childFunc, userData)
}
func (w *Window) InvalidateRect(rect *Rectangle, invalidateChildren T.Gboolean) {
	windowInvalidateRect(w, rect, invalidateChildren)
}
func (w *Window) InvalidateRegion(region *Region, invalidateChildren T.Gboolean) {
	windowInvalidateRegion(w, region, invalidateChildren)
}
func (w *Window) IsDestroyed() T.Gboolean                  { return windowIsDestroyed(w) }
func (w *Window) IsInputOnly() T.Gboolean                  { return windowIsInputOnly(w) }
func (w *Window) IsShaped() T.Gboolean                     { return windowIsShaped(w) }
func (w *Window) IsViewable() T.Gboolean                   { return windowIsViewable(w) }
func (w *Window) IsVisible() T.Gboolean                    { return windowIsVisible(w) }
func (w *Window) Lower()                                   { windowLower(w) }
func (w *Window) Maximize()                                { windowMaximize(w) }
func (w *Window) MergeChildInputShapes()                   { windowMergeChildInputShapes(w) }
func (w *Window) MergeChildShapes()                        { windowMergeChildShapes(w) }
func (w *Window) Move(x, y int)                            { windowMove(w, x, y) }
func (w *Window) MoveRegion(region *Region, dx, dy int)    { windowMoveRegion(w, region, dx, dy) }
func (w *Window) MoveResize(x, y, width, height int)       { windowMoveResize(w, x, y, width, height) }
func (w *Window) PeekChildren() *T.GList                   { return windowPeekChildren(w) }
func (w *Window) ProcessUpdates(updateChildren T.Gboolean) { windowProcessUpdates(w, updateChildren) }
func (w *Window) Raise()                                   { windowRaise(w) }
func (w *Window) RedirectToDrawable(drawable *Drawable, srcX, srcY, destX, destY, width, height int) {
	windowRedirectToDrawable(w, drawable, srcX, srcY, destX, destY, width, height)
}
func (w *Window) RegisterDnd() { windowRegisterDnd(w) }
func (w *Window) RemoveFilter(function T.GdkFilterFunc, data T.Gpointer) {
	windowRemoveFilter(w, function, data)
}
func (w *Window) RemoveRedirection()                        { windowRemoveRedirection(w) }
func (w *Window) Reparent(newParent *Window, x, y int)      { windowReparent(w, newParent, x, y) }
func (w *Window) Resize(width, height int)                  { windowResize(w, width, height) }
func (w *Window) Restack(sibling *Window, above T.Gboolean) { windowRestack(w, sibling, above) }
func (w *Window) Scroll(dx, dy int)                         { windowScroll(w, dx, dy) }
func (w *Window) SetAcceptFocus(acceptFocus T.Gboolean)     { windowSetAcceptFocus(w, acceptFocus) }
func (w *Window) SetBackground(color *Color)                { windowSetBackground(w, color) }
func (w *Window) SetBackPixmap(pixmap *Pixmap, parentRelative T.Gboolean) {
	windowSetBackPixmap(w, pixmap, parentRelative)
}
func (w *Window) SetChildInputShapes()                         { windowSetChildInputShapes(w) }
func (w *Window) SetChildShapes()                              { windowSetChildShapes(w) }
func (w *Window) SetComposited(composited T.Gboolean)          { windowSetComposited(w, composited) }
func (w *Window) SetCursor(cursor *Cursor)                     { windowSetCursor(w, cursor) }
func (w *Window) SetDecorations(decorations T.GdkWMDecoration) { windowSetDecorations(w, decorations) }
func (w *Window) SetEvents(eventMask EventMask)                { windowSetEvents(w, eventMask) }
func (w *Window) SetFocusOnMap(focusOnMap T.Gboolean)          { windowSetFocusOnMap(w, focusOnMap) }
func (w *Window) SetFunctions(functions T.GdkWMFunction)       { windowSetFunctions(w, functions) }
func (w *Window) SetGeometryHints(geometry *Geometry, geomMask WindowHints) {
	windowSetGeometryHints(w, geometry, geomMask)
}
func (w *Window) SetGroup(leader *Window) { windowSetGroup(w, leader) }
func (w *Window) SetHints(x, y, minWidth, minHeight, maxWidth, maxHeight, flags int) {
	windowSetHints(w, x, y, minWidth, minHeight, maxWidth, maxHeight, flags)
}
func (w *Window) SetIcon(iconWindow *Window, pixmap *Pixmap, mask *T.GdkBitmap) {
	windowSetIcon(w, iconWindow, pixmap, mask)
}
func (w *Window) SetIconList(pixbufs *T.GList)    { windowSetIconList(w, pixbufs) }
func (w *Window) SetIconName(name string)         { windowSetIconName(w, name) }
func (w *Window) SetKeepAbove(setting T.Gboolean) { windowSetKeepAbove(w, setting) }
func (w *Window) SetKeepBelow(setting T.Gboolean) { windowSetKeepBelow(w, setting) }
func (w *Window) SetModalHint(modal T.Gboolean)   { windowSetModalHint(w, modal) }
func (w *Window) SetOpacity(opacity float64)      { windowSetOpacity(w, opacity) }
func (w *Window) SetOverrideRedirect(overrideRedirect T.Gboolean) {
	windowSetOverrideRedirect(w, overrideRedirect)
}
func (w *Window) SetRole(role string)                    { windowSetRole(w, role) }
func (w *Window) SetSkipPagerHint(skipsPager T.Gboolean) { windowSetSkipPagerHint(w, skipsPager) }
func (w *Window) SetSkipTaskbarHint(skipsTaskbar T.Gboolean) {
	windowSetSkipTaskbarHint(w, skipsTaskbar)
}
func (w *Window) SetStartupId(startupId string) { windowSetStartupId(w, startupId) }
func (w *Window) SetStaticGravities(useStatic T.Gboolean) T.Gboolean {
	return windowSetStaticGravities(w, useStatic)
}
func (w *Window) SetTitle(title string)                        { windowSetTitle(w, title) }
func (w *Window) SetTransientFor(parent *Window)               { windowSetTransientFor(w, parent) }
func (w *Window) SetTypeHint(hint WindowTypeHint)              { windowSetTypeHint(w, hint) }
func (w *Window) SetUrgencyHint(urgent T.Gboolean)             { windowSetUrgencyHint(w, urgent) }
func (w *Window) SetUserData(userData T.Gpointer)              { windowSetUserData(w, userData) }
func (w *Window) ShapeCombineMask(mask *T.GdkBitmap, x, y int) { windowShapeCombineMask(w, mask, x, y) }
func (w *Window) ShapeCombineRegion(shapeRegion *Region, offsetX, offsetY int) {
	windowShapeCombineRegion(w, shapeRegion, offsetX, offsetY)
}
func (w *Window) Show()                          { windowShow(w) }
func (w *Window) ShowUnraised()                  { windowShowUnraised(w) }
func (w *Window) Stick()                         { windowStick(w) }
func (w *Window) ThawToplevelUpdatesLibgtkOnly() { windowThawToplevelUpdatesLibgtkOnly(w) }
func (w *Window) ThawUpdates()                   { windowThawUpdates(w) }
func (w *Window) Unfullscreen()                  { windowUnfullscreen(w) }
func (w *Window) Unmaximize()                    { windowUnmaximize(w) }
func (w *Window) Unstick()                       { windowUnstick(w) }
func (w *Window) Withdraw()                      { windowWithdraw(w) }

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
