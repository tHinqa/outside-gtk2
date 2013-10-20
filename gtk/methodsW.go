// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Widget struct {
	Object        T.GtkObject
	Private_flags uint16
	State         uint8
	Saved_state   uint8
	Name          *T.Gchar
	Style         *T.GtkStyle
	Requisition   T.GtkRequisition
	Allocation    T.GtkAllocation
	Window        *T.GdkWindow
	Parent        *Widget
}

var (
	WidgetGetType func() T.GType
	WidgetNew     func(t T.GType, firstPropertyName string, v ...VArg) *Widget

	//TODO(t): Methods?
	WidgetTranslateCoordinates func(srcWidget *Widget, destWidget *Widget, srcX int, srcY int, destX *int, destY *int) T.Gboolean
	WidgetStyleAttach          func(style *Widget)

	widgetActivate             func(w *Widget) T.Gboolean
	widgetAddAccelerator       func(w *Widget, accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags)
	widgetAddEvents            func(w *Widget, events int)
	widgetCanActivateAccel     func(w *Widget, signalId uint) T.Gboolean
	widgetChildFocus           func(w *Widget, direction T.GtkDirectionType) T.Gboolean
	widgetChildNotify          func(w *Widget, childProperty string)
	widgetCreatePangoContext   func(w *Widget) *T.PangoContext
	widgetCreatePangoLayout    func(w *Widget, text string) *T.PangoLayout
	widgetDestroy              func(w *Widget)
	widgetDestroyed            func(w *Widget, widgetPointer **Widget)
	widgetDraw                 func(w *Widget, area *T.GdkRectangle)
	widgetEnsureStyle          func(w *Widget)
	widgetErrorBell            func(w *Widget)
	widgetEvent                func(w *Widget, event *T.GdkEvent) T.Gboolean
	widgetFreezeChildNotify    func(w *Widget)
	widgetGetAccessible        func(w *Widget) *T.AtkObject
	widgetGetAllocation        func(w *Widget, allocation *T.GtkAllocation)
	widgetGetAncestor          func(w *Widget, widgetType T.GType) *Widget
	widgetGetAppPaintable      func(w *Widget) T.Gboolean
	widgetGetCanDefault        func(w *Widget) T.Gboolean
	widgetGetCanFocus          func(w *Widget) T.Gboolean
	widgetGetChildRequisition  func(w *Widget, requisition *T.GtkRequisition)
	widgetGetChildVisible      func(w *Widget) T.Gboolean
	widgetGetColormap          func(w *Widget) *T.GdkColormap
	widgetGetCompositeName     func(w *Widget) string
	widgetGetDisplay           func(w *Widget) *T.GdkDisplay
	widgetGetDoubleBuffered    func(w *Widget) T.Gboolean
	widgetGetEvents            func(w *Widget) int
	widgetGetExtensionEvents   func(w *Widget) T.GdkExtensionMode
	widgetGetHasWindow         func(w *Widget) T.Gboolean
	widgetGetMapped            func(w *Widget) T.Gboolean
	widgetGetModifierStyle     func(w *Widget) *T.GtkRcStyle
	widgetGetName              func(w *Widget) string
	widgetGetNoShowAll         func(w *Widget) T.Gboolean
	widgetGetPangoContext      func(w *Widget) *T.PangoContext
	widgetGetParent            func(w *Widget) *Widget
	widgetGetParentWindow      func(w *Widget) *T.GdkWindow
	widgetGetPointer           func(w *Widget, x *int, y *int)
	widgetGetRealized          func(w *Widget) T.Gboolean
	widgetGetReceivesDefault   func(w *Widget) T.Gboolean
	widgetGetRequisition       func(w *Widget, requisition *T.GtkRequisition)
	widgetGetRootWindow        func(w *Widget) *T.GdkWindow
	widgetGetScreen            func(w *Widget) *T.GdkScreen
	widgetGetSensitive         func(w *Widget) T.Gboolean
	widgetGetSettings          func(w *Widget) *T.GtkSettings
	widgetGetSizeRequest       func(w *Widget, width *int, height *int)
	widgetGetSnapshot          func(w *Widget, clipRect *T.GdkRectangle) *T.GdkPixmap
	widgetGetState             func(w *Widget) T.GtkStateType
	widgetGetStyle             func(w *Widget) *T.GtkStyle
	widgetGetToplevel          func(w *Widget) *Widget
	widgetGetVisible           func(w *Widget) T.Gboolean
	widgetGetVisual            func(w *Widget) *T.GdkVisual
	widgetGetWindow            func(w *Widget) *T.GdkWindow
	widgetGrabDefault          func(w *Widget)
	widgetGrabFocus            func(w *Widget)
	widgetHasDefault           func(w *Widget) T.Gboolean
	widgetHasFocus             func(w *Widget) T.Gboolean
	widgetHasGrab              func(w *Widget) T.Gboolean
	widgetHasRcStyle           func(w *Widget) T.Gboolean
	widgetHasScreen            func(w *Widget) T.Gboolean
	widgetHide                 func(w *Widget)
	widgetHideAll              func(w *Widget)
	widgetHideOnDelete         func(w *Widget) T.Gboolean
	widgetIntersect            func(w *Widget, area, intersection *T.GdkRectangle) T.Gboolean
	widgetIsAncestor           func(w *Widget, ancestor *Widget) T.Gboolean
	widgetIsDrawable           func(w *Widget) T.Gboolean
	widgetIsFocus              func(w *Widget) T.Gboolean
	widgetIsSensitive          func(w *Widget) T.Gboolean
	widgetIsToplevel           func(w *Widget) T.Gboolean
	widgetKeynavFailed         func(w *Widget, direction T.GtkDirectionType) T.Gboolean
	widgetListAccelClosures    func(w *Widget) *T.GList
	widgetMap                  func(w *Widget)
	widgetMnemonicActivate     func(w *Widget, groupCycling T.Gboolean) T.Gboolean
	widgetModifyBase           func(w *Widget, state T.GtkStateType, color *T.GdkColor)
	widgetModifyBg             func(w *Widget, state T.GtkStateType, color *T.GdkColor)
	widgetModifyCursor         func(w *Widget, primary *T.GdkColor, secondary *T.GdkColor)
	widgetModifyFg             func(w *Widget, state T.GtkStateType, color *T.GdkColor)
	widgetModifyFont           func(w *Widget, fontDesc *T.PangoFontDescription)
	widgetModifyStyle          func(w *Widget, style *T.GtkRcStyle)
	widgetModifyText           func(w *Widget, state T.GtkStateType, color *T.GdkColor)
	widgetQueueClear           func(w *Widget)
	widgetQueueClearArea       func(w *Widget, x, y, width, height int)
	widgetQueueDraw            func(w *Widget)
	widgetQueueDrawArea        func(w *Widget, x, y, width, height int)
	widgetQueueResize          func(w *Widget)
	widgetQueueResizeNoRedraw  func(w *Widget)
	widgetRealize              func(w *Widget)
	widgetRef                  func(w *Widget) *Widget
	widgetRegionIntersect      func(w *Widget, region *T.GdkRegion) *T.GdkRegion
	widgetRemoveAccelerator    func(w *Widget, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType) T.Gboolean
	widgetRenderIcon           func(w *Widget, stockId string, size IconSize, detail string) *T.GdkPixbuf
	widgetReparent             func(widget, newParent *Widget)
	widgetResetRcStyles        func(w *Widget)
	widgetSendExpose           func(w *Widget, event *T.GdkEvent) int
	widgetSendFocusChange      func(w *Widget, event *T.GdkEvent) T.Gboolean
	widgetSet                  func(w *Widget, firstPropertyName string, v ...VArg)
	widgetSetAccelPath         func(w *Widget, accelPath string, accelGroup *AccelGroup)
	widgetSetAllocation        func(w *Widget, allocation *T.GtkAllocation)
	widgetSetAppPaintable      func(w *Widget, appPaintable T.Gboolean)
	widgetSetCanDefault        func(w *Widget, canDefault T.Gboolean)
	widgetSetCanFocus          func(w *Widget, canFocus T.Gboolean)
	widgetSetChildVisible      func(w *Widget, isVisible T.Gboolean)
	widgetSetColormap          func(w *Widget, colormap *T.GdkColormap)
	widgetSetCompositeName     func(w *Widget, name string)
	widgetSetDoubleBuffered    func(w *Widget, doubleBuffered T.Gboolean)
	widgetSetEvents            func(w *Widget, events int)
	widgetSetExtensionEvents   func(w *Widget, mode T.GdkExtensionMode)
	widgetSetHasWindow         func(w *Widget, hasWindow T.Gboolean)
	widgetSetMapped            func(w *Widget, mapped T.Gboolean)
	widgetSetName              func(w *Widget, name string)
	widgetSetNoShowAll         func(w *Widget, noShowAll T.Gboolean)
	widgetSetParent            func(w *Widget, parent *Widget)
	widgetSetParentWindow      func(w *Widget, parentWindow *T.GdkWindow)
	widgetSetRealized          func(w *Widget, realized T.Gboolean)
	widgetSetReceivesDefault   func(w *Widget, receivesDefault T.Gboolean)
	widgetSetRedrawOnAllocate  func(w *Widget, redrawOnAllocate T.Gboolean)
	widgetSetScrollAdjustments func(w *Widget, hadjustment, vadjustment *Adjustment) T.Gboolean
	widgetSetSensitive         func(w *Widget, sensitive T.Gboolean)
	widgetSetSizeRequest       func(w *Widget, width int, height int)
	widgetSetState             func(w *Widget, state T.GtkStateType)
	widgetSetStyle             func(w *Widget, style *T.GtkStyle)
	widgetSetUposition         func(w *Widget, x int, y int)
	widgetSetUsize             func(w *Widget, width int, height int)
	widgetSetVisible           func(w *Widget, visible T.Gboolean)
	widgetSetWindow            func(w *Widget, window *T.GdkWindow)
	widgetShow                 func(w *Widget)
	widgetShowAll              func(w *Widget)
	widgetShowNow              func(w *Widget)
	widgetSizeAllocate         func(w *Widget, allocation *T.GtkAllocation)
	widgetSizeRequest          func(w *Widget, requisition *T.GtkRequisition)
	widgetThawChildNotify      func(w *Widget)
	widgetUnmap                func(w *Widget)
	widgetUnparent             func(w *Widget)
	widgetUnrealize            func(w *Widget)
	widgetUnref                func(w *Widget)
)

func (w *Widget) Activate() T.Gboolean { return widgetActivate(w) }
func (w *Widget) AddAccelerator(accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags) {
	widgetAddAccelerator(w, accelSignal, accelGroup, accelKey, accelMods, accelFlags)
}
func (w *Widget) AddEvents(events int) { widgetAddEvents(w, events) }
func (w *Widget) CanActivateAccel(signalId uint) T.Gboolean {
	return widgetCanActivateAccel(w, signalId)
}
func (w *Widget) ChildFocus(direction T.GtkDirectionType) T.Gboolean {
	return widgetChildFocus(w, direction)
}
func (w *Widget) ChildNotify(childProperty string)    { widgetChildNotify(w, childProperty) }
func (w *Widget) CreatePangoContext() *T.PangoContext { return widgetCreatePangoContext(w) }
func (w *Widget) CreatePangoLayout(text string) *T.PangoLayout {
	return widgetCreatePangoLayout(w, text)
}
func (w *Widget) Destroy()                                  { widgetDestroy(w) }
func (w *Widget) Destroyed(widgetPointer **Widget)          { widgetDestroyed(w, widgetPointer) }
func (w *Widget) Draw(area *T.GdkRectangle)                 { widgetDraw(w, area) }
func (w *Widget) EnsureStyle()                              { widgetEnsureStyle(w) }
func (w *Widget) ErrorBell()                                { widgetErrorBell(w) }
func (w *Widget) Event(event *T.GdkEvent) T.Gboolean        { return widgetEvent(w, event) }
func (w *Widget) FreezeChildNotify()                        { widgetFreezeChildNotify(w) }
func (w *Widget) GetAccessible() *T.AtkObject               { return widgetGetAccessible(w) }
func (w *Widget) GetAllocation(allocation *T.GtkAllocation) { widgetGetAllocation(w, allocation) }
func (w *Widget) GetAncestor(widgetType T.GType) *Widget    { return widgetGetAncestor(w, widgetType) }
func (w *Widget) GetAppPaintable() T.Gboolean               { return widgetGetAppPaintable(w) }
func (w *Widget) GetCanDefault() T.Gboolean                 { return widgetGetCanDefault(w) }
func (w *Widget) GetCanFocus() T.Gboolean                   { return widgetGetCanFocus(w) }
func (w *Widget) GetChildRequisition(requisition *T.GtkRequisition) {
	widgetGetChildRequisition(w, requisition)
}
func (w *Widget) GetChildVisible() T.Gboolean                  { return widgetGetChildVisible(w) }
func (w *Widget) GetColormap() *T.GdkColormap                  { return widgetGetColormap(w) }
func (w *Widget) GetCompositeName() string                     { return widgetGetCompositeName(w) }
func (w *Widget) GetDisplay() *T.GdkDisplay                    { return widgetGetDisplay(w) }
func (w *Widget) GetDoubleBuffered() T.Gboolean                { return widgetGetDoubleBuffered(w) }
func (w *Widget) GetEvents() int                               { return widgetGetEvents(w) }
func (w *Widget) GetExtensionEvents() T.GdkExtensionMode       { return widgetGetExtensionEvents(w) }
func (w *Widget) GetHasWindow() T.Gboolean                     { return widgetGetHasWindow(w) }
func (w *Widget) GetMapped() T.Gboolean                        { return widgetGetMapped(w) }
func (w *Widget) GetModifierStyle() *T.GtkRcStyle              { return widgetGetModifierStyle(w) }
func (w *Widget) GetName() string                              { return widgetGetName(w) }
func (w *Widget) GetNoShowAll() T.Gboolean                     { return widgetGetNoShowAll(w) }
func (w *Widget) GetPangoContext() *T.PangoContext             { return widgetGetPangoContext(w) }
func (w *Widget) GetParent() *Widget                           { return widgetGetParent(w) }
func (w *Widget) GetParentWindow() *T.GdkWindow                { return widgetGetParentWindow(w) }
func (w *Widget) GetPointer(x *int, y *int)                    { widgetGetPointer(w, x, y) }
func (w *Widget) GetRealized() T.Gboolean                      { return widgetGetRealized(w) }
func (w *Widget) GetReceivesDefault() T.Gboolean               { return widgetGetReceivesDefault(w) }
func (w *Widget) GetRequisition(requisition *T.GtkRequisition) { widgetGetRequisition(w, requisition) }
func (w *Widget) GetRootWindow() *T.GdkWindow                  { return widgetGetRootWindow(w) }
func (w *Widget) GetScreen() *T.GdkScreen                      { return widgetGetScreen(w) }
func (w *Widget) GetSensitive() T.Gboolean                     { return widgetGetSensitive(w) }
func (w *Widget) GetSettings() *T.GtkSettings                  { return widgetGetSettings(w) }
func (w *Widget) GetSizeRequest(width, height *int)            { widgetGetSizeRequest(w, width, height) }
func (w *Widget) GetSnapshot(clipRect *T.GdkRectangle) *T.GdkPixmap {
	return widgetGetSnapshot(w, clipRect)
}
func (w *Widget) GetState() T.GtkStateType { return widgetGetState(w) }
func (w *Widget) GetStyle() *T.GtkStyle    { return widgetGetStyle(w) }
func (w *Widget) GetToplevel() *Widget     { return widgetGetToplevel(w) }
func (w *Widget) GetVisible() T.Gboolean   { return widgetGetVisible(w) }
func (w *Widget) GetVisual() *T.GdkVisual  { return widgetGetVisual(w) }
func (w *Widget) GetWindow() *T.GdkWindow  { return widgetGetWindow(w) }
func (w *Widget) GrabDefault()             { widgetGrabDefault(w) }
func (w *Widget) GrabFocus()               { widgetGrabFocus(w) }
func (w *Widget) HasDefault() T.Gboolean   { return widgetHasDefault(w) }
func (w *Widget) HasFocus() T.Gboolean     { return widgetHasFocus(w) }
func (w *Widget) HasGrab() T.Gboolean      { return widgetHasGrab(w) }
func (w *Widget) HasRcStyle() T.Gboolean   { return widgetHasRcStyle(w) }
func (w *Widget) HasScreen() T.Gboolean    { return widgetHasScreen(w) }
func (w *Widget) Hide()                    { widgetHide(w) }
func (w *Widget) HideAll()                 { widgetHideAll(w) }
func (w *Widget) HideOnDelete() T.Gboolean { return widgetHideOnDelete(w) }
func (w *Widget) Intersect(area, intersection *T.GdkRectangle) T.Gboolean {
	return widgetIntersect(w, area, intersection)
}
func (w *Widget) IsAncestor(ancestor *Widget) T.Gboolean { return widgetIsAncestor(w, ancestor) }
func (w *Widget) IsDrawable() T.Gboolean                 { return widgetIsDrawable(w) }
func (w *Widget) IsFocus() T.Gboolean                    { return widgetIsFocus(w) }
func (w *Widget) IsSensitive() T.Gboolean                { return widgetIsSensitive(w) }
func (w *Widget) IsToplevel() T.Gboolean                 { return widgetIsToplevel(w) }
func (w *Widget) KeynavFailed(direction T.GtkDirectionType) T.Gboolean {
	return widgetKeynavFailed(w, direction)
}
func (w *Widget) ListAccelClosures() *T.GList { return widgetListAccelClosures(w) }
func (w *Widget) Map()                        { widgetMap(w) }
func (w *Widget) MnemonicActivate(groupCycling T.Gboolean) T.Gboolean {
	return widgetMnemonicActivate(w, groupCycling)
}
func (w *Widget) ModifyBase(state T.GtkStateType, color *T.GdkColor) {
	widgetModifyBase(w, state, color)
}
func (w *Widget) ModifyBg(state T.GtkStateType, color *T.GdkColor) { widgetModifyBg(w, state, color) }
func (w *Widget) ModifyCursor(primary *T.GdkColor, secondary *T.GdkColor) {
	widgetModifyCursor(w, primary, secondary)
}
func (w *Widget) ModifyFg(state T.GtkStateType, color *T.GdkColor) { widgetModifyFg(w, state, color) }
func (w *Widget) ModifyFont(fontDesc *T.PangoFontDescription)      { widgetModifyFont(w, fontDesc) }
func (w *Widget) ModifyStyle(style *T.GtkRcStyle)                  { widgetModifyStyle(w, style) }
func (w *Widget) ModifyText(state T.GtkStateType, color *T.GdkColor) {
	widgetModifyText(w, state, color)
}
func (w *Widget) QueueClear()                            { widgetQueueClear(w) }
func (w *Widget) QueueClearArea(x, y, width, height int) { widgetQueueClearArea(w, x, y, width, height) }
func (w *Widget) QueueDraw()                             { widgetQueueDraw(w) }
func (w *Widget) QueueDrawArea(x, y, width, height int)  { widgetQueueDrawArea(w, x, y, width, height) }
func (w *Widget) QueueResize()                           { widgetQueueResize(w) }
func (w *Widget) QueueResizeNoRedraw()                   { widgetQueueResizeNoRedraw(w) }
func (w *Widget) Realize()                               { widgetRealize(w) }
func (w *Widget) Ref() *Widget                           { return widgetRef(w) }
func (w *Widget) RegionIntersect(region *T.GdkRegion) *T.GdkRegion {
	return widgetRegionIntersect(w, region)
}
func (w *Widget) RemoveAccelerator(accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType) T.Gboolean {
	return widgetRemoveAccelerator(w, accelGroup, accelKey, accelMods)
}
func (w *Widget) RenderIcon(stockId string, size IconSize, detail string) *T.GdkPixbuf {
	return widgetRenderIcon(w, stockId, size, detail)
}
func (w *Widget) Reparent(newParent *Widget)                   { widgetReparent(w, newParent) }
func (w *Widget) ResetRcStyles()                               { widgetResetRcStyles(w) }
func (w *Widget) SendExpose(event *T.GdkEvent) int             { return widgetSendExpose(w, event) }
func (w *Widget) SendFocusChange(event *T.GdkEvent) T.Gboolean { return widgetSendFocusChange(w, event) }
func (w *Widget) Set(firstPropertyName string, v ...VArg)      { widgetSet(w, firstPropertyName, v) }
func (w *Widget) SetAccelPath(accelPath string, accelGroup *AccelGroup) {
	widgetSetAccelPath(w, accelPath, accelGroup)
}
func (w *Widget) SetAllocation(allocation *T.GtkAllocation) { widgetSetAllocation(w, allocation) }
func (w *Widget) SetAppPaintable(appPaintable T.Gboolean)   { widgetSetAppPaintable(w, appPaintable) }
func (w *Widget) SetCanDefault(canDefault T.Gboolean)       { widgetSetCanDefault(w, canDefault) }
func (w *Widget) SetCanFocus(canFocus T.Gboolean)           { widgetSetCanFocus(w, canFocus) }
func (w *Widget) SetChildVisible(isVisible T.Gboolean)      { widgetSetChildVisible(w, isVisible) }
func (w *Widget) SetColormap(colormap *T.GdkColormap)       { widgetSetColormap(w, colormap) }
func (w *Widget) SetCompositeName(name string)              { widgetSetCompositeName(w, name) }
func (w *Widget) SetDoubleBuffered(doubleBuffered T.Gboolean) {
	widgetSetDoubleBuffered(w, doubleBuffered)
}
func (w *Widget) SetEvents(events int)                       { widgetSetEvents(w, events) }
func (w *Widget) SetExtensionEvents(mode T.GdkExtensionMode) { widgetSetExtensionEvents(w, mode) }
func (w *Widget) SetHasWindow(hasWindow T.Gboolean)          { widgetSetHasWindow(w, hasWindow) }
func (w *Widget) SetMapped(mapped T.Gboolean)                { widgetSetMapped(w, mapped) }
func (w *Widget) SetName(name string)                        { widgetSetName(w, name) }
func (w *Widget) SetNoShowAll(noShowAll T.Gboolean)          { widgetSetNoShowAll(w, noShowAll) }
func (w *Widget) SetParent(parent *Widget)                   { widgetSetParent(w, parent) }
func (w *Widget) SetParentWindow(parentWindow *T.GdkWindow)  { widgetSetParentWindow(w, parentWindow) }
func (w *Widget) SetRealized(realized T.Gboolean)            { widgetSetRealized(w, realized) }
func (w *Widget) SetReceivesDefault(receivesDefault T.Gboolean) {
	widgetSetReceivesDefault(w, receivesDefault)
}
func (w *Widget) SetRedrawOnAllocate(redrawOnAllocate T.Gboolean) {
	widgetSetRedrawOnAllocate(w, redrawOnAllocate)
}
func (w *Widget) SetScrollAdjustments(hadjustment, vadjustment *Adjustment) T.Gboolean {
	return widgetSetScrollAdjustments(w, hadjustment, vadjustment)
}
func (w *Widget) SetSensitive(sensitive T.Gboolean)         { widgetSetSensitive(w, sensitive) }
func (w *Widget) SetSizeRequest(width, height int)          { widgetSetSizeRequest(w, width, height) }
func (w *Widget) SetState(state T.GtkStateType)             { widgetSetState(w, state) }
func (w *Widget) SetStyle(style *T.GtkStyle)                { widgetSetStyle(w, style) }
func (w *Widget) SetUposition(x, y int)                     { widgetSetUposition(w, x, y) }
func (w *Widget) SetUsize(width, height int)                { widgetSetUsize(w, width, height) }
func (w *Widget) SetVisible(visible T.Gboolean)             { widgetSetVisible(w, visible) }
func (w *Widget) SetWindow(window *T.GdkWindow)             { widgetSetWindow(w, window) }
func (w *Widget) Show()                                     { widgetShow(w) }
func (w *Widget) ShowAll()                                  { widgetShowAll(w) }
func (w *Widget) ShowNow()                                  { widgetShowNow(w) }
func (w *Widget) SizeAllocate(allocation *T.GtkAllocation)  { widgetSizeAllocate(w, allocation) }
func (w *Widget) SizeRequest(requisition *T.GtkRequisition) { widgetSizeRequest(w, requisition) }
func (w *Widget) ThawChildNotify()                          { widgetThawChildNotify(w) }
func (w *Widget) Unmap()                                    { widgetUnmap(w) }
func (w *Widget) Unparent()                                 { widgetUnparent(w) }
func (w *Widget) Unrealize()                                { widgetUnrealize(w) }
func (w *Widget) Unref()                                    { widgetUnref(w) }

var (
	WindowGetType func() T.GType
	WindowNew     func(t T.GtkWindowType) *Widget

	WindowSetTitle func(
		window *T.GtkWindow,
		title string)

	WindowGetTitle func(
		window *T.GtkWindow) string

	WindowSetWmclass func(
		window *T.GtkWindow,
		wmclassName string,
		wmclassClass string)

	WindowSetRole func(
		window *T.GtkWindow,
		role string)

	WindowSetStartupId func(
		window *T.GtkWindow,
		startupId string)

	WindowGetRole func(
		window *T.GtkWindow) string

	WindowAddAccelGroup func(
		window *T.GtkWindow,
		accelGroup *AccelGroup)

	WindowRemoveAccelGroup func(
		window *T.GtkWindow,
		accelGroup *AccelGroup)

	WindowSetPosition func(
		window *T.GtkWindow,
		position T.GtkWindowPosition)

	WindowActivateFocus func(
		window *T.GtkWindow) T.Gboolean

	WindowSetFocus func(
		window *T.GtkWindow,
		focus *Widget)

	WindowGetFocus func(
		window *T.GtkWindow) *Widget

	WindowSetDefault func(
		window *T.GtkWindow,
		defaultWidget *Widget)

	WindowGetDefaultWidget func(
		window *T.GtkWindow) *Widget

	WindowActivateDefault func(
		window *T.GtkWindow) T.Gboolean

	WindowSetTransientFor func(
		window *T.GtkWindow,
		parent *T.GtkWindow)

	WindowGetTransientFor func(
		window *T.GtkWindow) *T.GtkWindow

	WindowSetOpacity func(
		window *T.GtkWindow,
		opacity float64)

	WindowGetOpacity func(
		window *T.GtkWindow) float64

	WindowSetTypeHint func(
		window *T.GtkWindow,
		hint T.GdkWindowTypeHint)

	WindowGetTypeHint func(
		window *T.GtkWindow) T.GdkWindowTypeHint

	WindowSetSkipTaskbarHint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetSkipTaskbarHint func(
		window *T.GtkWindow) T.Gboolean

	WindowSetSkipPagerHint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetSkipPagerHint func(
		window *T.GtkWindow) T.Gboolean

	WindowSetUrgencyHint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetUrgencyHint func(
		window *T.GtkWindow) T.Gboolean

	WindowSetAcceptFocus func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetAcceptFocus func(
		window *T.GtkWindow) T.Gboolean

	WindowSetFocusOnMap func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetFocusOnMap func(
		window *T.GtkWindow) T.Gboolean

	WindowSetDestroyWithParent func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetDestroyWithParent func(
		window *T.GtkWindow) T.Gboolean

	WindowSetMnemonicsVisible func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetMnemonicsVisible func(
		window *T.GtkWindow) T.Gboolean

	WindowSetResizable func(
		window *T.GtkWindow,
		resizable T.Gboolean)

	WindowGetResizable func(
		window *T.GtkWindow) T.Gboolean

	WindowSetGravity func(
		window *T.GtkWindow,
		gravity T.GdkGravity)

	WindowGetGravity func(
		window *T.GtkWindow) T.GdkGravity

	WindowSetGeometryHints func(
		window *T.GtkWindow,
		geometryWidget *Widget,
		geometry *T.GdkGeometry,
		geomMask T.GdkWindowHints)

	WindowSetScreen func(
		window *T.GtkWindow,
		screen *T.GdkScreen)

	WindowGetScreen func(
		window *T.GtkWindow) *T.GdkScreen

	WindowIsActive func(
		window *T.GtkWindow) T.Gboolean

	WindowHasToplevelFocus func(
		window *T.GtkWindow) T.Gboolean

	WindowSetHasFrame func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetHasFrame func(
		window *T.GtkWindow) T.Gboolean

	WindowSetFrameDimensions func(
		window *T.GtkWindow,
		left int,
		top int,
		right int,
		bottom int)

	WindowGetFrameDimensions func(
		window *T.GtkWindow,
		left *int,
		top *int,
		right *int,
		bottom *int)

	WindowSetDecorated func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetDecorated func(
		window *T.GtkWindow) T.Gboolean

	WindowSetDeletable func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetDeletable func(
		window *T.GtkWindow) T.Gboolean

	WindowSetIconList func(
		window *T.GtkWindow,
		list *T.GList)

	WindowGetIconList func(
		window *T.GtkWindow) *T.GList

	WindowSetIcon func(
		window *T.GtkWindow,
		icon *T.GdkPixbuf)

	WindowSetIconName func(
		window *T.GtkWindow,
		name string)

	WindowSetIconFromFile func(
		window *T.GtkWindow,
		filename string,
		err **T.GError) T.Gboolean

	WindowGetIcon func(
		window *T.GtkWindow) *T.GdkPixbuf

	WindowGetIconName func(
		window *T.GtkWindow) string

	WindowSetDefaultIconList func(
		list *T.GList)

	WindowGetDefaultIconList func() *T.GList

	WindowSetDefaultIcon func(
		icon *T.GdkPixbuf)

	WindowSetDefaultIconName func(
		name string)

	WindowGetDefaultIconName func() string

	WindowSetDefaultIconFromFile func(
		filename string,
		err **T.GError) T.Gboolean

	WindowSetAutoStartupNotification func(
		setting T.Gboolean)

	WindowSetModal func(
		window *T.GtkWindow,
		modal T.Gboolean)

	WindowGetModal func(
		window *T.GtkWindow) T.Gboolean

	WindowListToplevels func() *T.GList

	WindowAddMnemonic func(
		window *T.GtkWindow,
		keyval uint,
		target *Widget)

	WindowRemoveMnemonic func(
		window *T.GtkWindow,
		keyval uint,
		target *Widget)

	WindowMnemonicActivate func(
		window *T.GtkWindow,
		keyval uint,
		modifier T.GdkModifierType) T.Gboolean

	WindowSetMnemonicModifier func(
		window *T.GtkWindow,
		modifier T.GdkModifierType)

	WindowGetMnemonicModifier func(
		window *T.GtkWindow) T.GdkModifierType

	WindowActivateKey func(
		window *T.GtkWindow,
		event *T.GdkEventKey) T.Gboolean

	WindowPropagateKeyEvent func(
		window *T.GtkWindow,
		event *T.GdkEventKey) T.Gboolean

	WindowPresent func(
		window *T.GtkWindow)

	WindowPresentWithTime func(
		window *T.GtkWindow,
		timestamp T.GUint32)

	WindowIconify func(
		window *T.GtkWindow)

	WindowDeiconify func(
		window *T.GtkWindow)

	WindowStick func(
		window *T.GtkWindow)

	WindowUnstick func(
		window *T.GtkWindow)

	WindowMaximize func(
		window *T.GtkWindow)

	WindowUnmaximize func(
		window *T.GtkWindow)

	WindowFullscreen func(
		window *T.GtkWindow)

	WindowUnfullscreen func(
		window *T.GtkWindow)

	WindowSetKeepAbove func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowSetKeepBelow func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowBeginResizeDrag func(
		window *T.GtkWindow,
		edge T.GdkWindowEdge,
		button int,
		rootX int,
		rootY int,
		timestamp T.GUint32)

	WindowBeginMoveDrag func(
		window *T.GtkWindow,
		button int,
		rootX int,
		rootY int,
		timestamp T.GUint32)

	WindowSetPolicy func(
		window *T.GtkWindow,
		allowShrink int,
		allowGrow int,
		autoShrink int)

	WindowSetDefaultSize func(
		window *T.GtkWindow,
		width int,
		height int)

	WindowGetDefaultSize func(
		window *T.GtkWindow,
		width *int,
		height *int)

	WindowResize func(
		window *T.GtkWindow,
		width int,
		height int)

	WindowGetSize func(
		window *T.GtkWindow,
		width *int,
		height *int)

	WindowMove func(
		window *T.GtkWindow,
		x int,
		y int)

	WindowGetPosition func(
		window *T.GtkWindow,
		rootX *int,
		rootY *int)

	WindowParseGeometry func(
		window *T.GtkWindow,
		geometry string) T.Gboolean

	WindowGetGroup func(
		window *T.GtkWindow) *T.GtkWindowGroup

	WindowHasGroup func(
		window *T.GtkWindow) T.Gboolean

	WindowReshowWithInitialSize func(
		window *T.GtkWindow)

	WindowGetWindowType func(
		window *T.GtkWindow) T.GtkWindowType

	WindowGroupNew func() *T.GtkWindowGroup

	WindowGroupAddWindow func(
		windowGroup *T.GtkWindowGroup,
		window *T.GtkWindow)

	WindowGroupRemoveWindow func(
		windowGroup *T.GtkWindowGroup,
		window *T.GtkWindow)

	WindowGroupListWindows func(
		windowGroup *T.GtkWindowGroup) *T.GList

	WindowRemoveEmbeddedXid func(
		window *T.GtkWindow,
		xid T.GdkNativeWindow)

	WindowAddEmbeddedXid func(
		window *T.GtkWindow,
		xid T.GdkNativeWindow)
)
