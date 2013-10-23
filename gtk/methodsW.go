// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type WrapMode T.Enum

const (
	WRAP_NONE WrapMode = iota
	WRAP_CHAR
	WRAP_WORD
	WRAP_WORD_CHAR
)

var WrapModeGetType func() T.GType

type Widget struct {
	Object       Object
	PrivateFlags uint16
	State        uint8
	SavedState   uint8
	Name         *T.Gchar
	Style        *Style
	Requisition  Requisition
	Allocation   Allocation
	Window       *T.GdkWindow
	Parent       *Widget
}

var (
	WidgetGetType func() T.GType
	WidgetNew     func(t T.GType, firstPropertyName string, v ...VArg) *Widget

	WidgetFlagsGetType func() T.GType

	WidgetGetDefaultColormap  func() *T.GdkColormap
	WidgetGetDefaultDirection func() TextDirection
	WidgetGetDefaultStyle     func() *Style
	WidgetGetDefaultVisual    func() *T.GdkVisual
	WidgetPopColormap         func()
	WidgetPopCompositeChild   func()
	WidgetPushColormap        func(cmap *T.GdkColormap)
	WidgetPushCompositeChild  func()
	WidgetSetDefaultColormap  func(colormap *T.GdkColormap)
	WidgetSetDefaultDirection func(dir TextDirection)

	//TODO(t): Methods?
	WidgetTranslateCoordinates func(srcWidget *Widget, destWidget *Widget, srcX int, srcY int, destX *int, destY *int) T.Gboolean
	WidgetStyleAttach          func(style *Widget)

	widgetActivate              func(w *Widget) T.Gboolean
	widgetAddAccelerator        func(w *Widget, accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags)
	widgetAddEvents             func(w *Widget, events int)
	widgetAddMnemonicLabel      func(w *Widget, label *Widget)
	widgetCanActivateAccel      func(w *Widget, signalId uint) T.Gboolean
	widgetChildFocus            func(w *Widget, direction DirectionType) T.Gboolean
	widgetChildNotify           func(w *Widget, childProperty string)
	widgetClassPath             func(w *Widget, pathLength *uint, path, pathReversed **T.Gchar)
	widgetCreatePangoContext    func(w *Widget) *T.PangoContext
	widgetCreatePangoLayout     func(w *Widget, text string) *T.PangoLayout
	widgetDestroy               func(w *Widget)
	widgetDestroyed             func(w *Widget, widgetPointer **Widget)
	widgetDraw                  func(w *Widget, area *T.GdkRectangle)
	widgetEnsureStyle           func(w *Widget)
	widgetErrorBell             func(w *Widget)
	widgetEvent                 func(w *Widget, event *T.GdkEvent) T.Gboolean
	widgetFreezeChildNotify     func(w *Widget)
	widgetGetAccessible         func(w *Widget) *T.AtkObject
	widgetGetAllocation         func(w *Widget, allocation *Allocation)
	widgetGetAncestor           func(w *Widget, widgetType T.GType) *Widget
	widgetGetAppPaintable       func(w *Widget) T.Gboolean
	widgetGetCanDefault         func(w *Widget) T.Gboolean
	widgetGetCanFocus           func(w *Widget) T.Gboolean
	widgetGetChildRequisition   func(w *Widget, requisition *Requisition)
	widgetGetChildVisible       func(w *Widget) T.Gboolean
	widgetGetColormap           func(w *Widget) *T.GdkColormap
	widgetGetCompositeName      func(w *Widget) string
	widgetGetDirection          func(w *Widget) TextDirection
	widgetGetDisplay            func(w *Widget) *T.GdkDisplay
	widgetGetDoubleBuffered     func(w *Widget) T.Gboolean
	widgetGetEvents             func(w *Widget) int
	widgetGetExtensionEvents    func(w *Widget) T.GdkExtensionMode
	widgetGetHasTooltip         func(w *Widget) T.Gboolean
	widgetGetHasWindow          func(w *Widget) T.Gboolean
	widgetGetMapped             func(w *Widget) T.Gboolean
	widgetGetModifierStyle      func(w *Widget) *RcStyle
	widgetGetName               func(w *Widget) string
	widgetGetNoShowAll          func(w *Widget) T.Gboolean
	widgetGetPangoContext       func(w *Widget) *T.PangoContext
	widgetGetParent             func(w *Widget) *Widget
	widgetGetParentWindow       func(w *Widget) *T.GdkWindow
	widgetGetPointer            func(w *Widget, x *int, y *int)
	widgetGetRealized           func(w *Widget) T.Gboolean
	widgetGetReceivesDefault    func(w *Widget) T.Gboolean
	widgetGetRequisition        func(w *Widget, requisition *Requisition)
	widgetGetRootWindow         func(w *Widget) *T.GdkWindow
	widgetGetScreen             func(w *Widget) *T.GdkScreen
	widgetGetSensitive          func(w *Widget) T.Gboolean
	widgetGetSettings           func(w *Widget) *Settings
	widgetGetSizeRequest        func(w *Widget, width *int, height *int)
	widgetGetSnapshot           func(w *Widget, clipRect *T.GdkRectangle) *T.GdkPixmap
	widgetGetState              func(w *Widget) StateType
	widgetGetStyle              func(w *Widget) *Style
	widgetGetTooltipMarkup      func(w *Widget) string
	widgetGetTooltipText        func(w *Widget) string
	widgetGetTooltipWindow      func(w *Widget) *Window
	widgetGetToplevel           func(w *Widget) *Widget
	widgetGetVisible            func(w *Widget) T.Gboolean
	widgetGetVisual             func(w *Widget) *T.GdkVisual
	widgetGetWindow             func(w *Widget) *T.GdkWindow
	widgetGrabDefault           func(w *Widget)
	widgetGrabFocus             func(w *Widget)
	widgetHasDefault            func(w *Widget) T.Gboolean
	widgetHasFocus              func(w *Widget) T.Gboolean
	widgetHasGrab               func(w *Widget) T.Gboolean
	widgetHasRcStyle            func(w *Widget) T.Gboolean
	widgetHasScreen             func(w *Widget) T.Gboolean
	widgetHide                  func(w *Widget)
	widgetHideAll               func(w *Widget)
	widgetHideOnDelete          func(w *Widget) T.Gboolean
	widgetInputShapeCombineMask func(w *Widget, shapeMask *T.GdkBitmap, offsetX, offsetY int)
	widgetIntersect             func(w *Widget, area, intersection *T.GdkRectangle) T.Gboolean
	widgetIsAncestor            func(w *Widget, ancestor *Widget) T.Gboolean
	widgetIsComposited          func(w *Widget) T.Gboolean
	widgetIsDrawable            func(w *Widget) T.Gboolean
	widgetIsFocus               func(w *Widget) T.Gboolean
	widgetIsSensitive           func(w *Widget) T.Gboolean
	widgetIsToplevel            func(w *Widget) T.Gboolean
	widgetKeynavFailed          func(w *Widget, direction DirectionType) T.Gboolean
	widgetListAccelClosures     func(w *Widget) *T.GList
	widgetListMnemonicLabels    func(w *Widget) *T.GList
	widgetMap                   func(w *Widget)
	widgetMnemonicActivate      func(w *Widget, groupCycling T.Gboolean) T.Gboolean
	widgetModifyBase            func(w *Widget, state StateType, color *T.GdkColor)
	widgetModifyBg              func(w *Widget, state StateType, color *T.GdkColor)
	widgetModifyCursor          func(w *Widget, primary *T.GdkColor, secondary *T.GdkColor)
	widgetModifyFg              func(w *Widget, state StateType, color *T.GdkColor)
	widgetModifyFont            func(w *Widget, fontDesc *T.PangoFontDescription)
	widgetModifyStyle           func(w *Widget, style *RcStyle)
	widgetModifyText            func(w *Widget, state StateType, color *T.GdkColor)
	widgetPath                  func(w *Widget, pathLength *uint, path, pathReversed **T.Gchar)
	widgetQueueClear            func(w *Widget)
	widgetQueueClearArea        func(w *Widget, x, y, width, height int)
	widgetQueueDraw             func(w *Widget)
	widgetQueueDrawArea         func(w *Widget, x, y, width, height int)
	widgetQueueResize           func(w *Widget)
	widgetQueueResizeNoRedraw   func(w *Widget)
	widgetRealize               func(w *Widget)
	widgetRef                   func(w *Widget) *Widget
	widgetRegionIntersect       func(w *Widget, region *T.GdkRegion) *T.GdkRegion
	widgetRemoveAccelerator     func(w *Widget, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType) T.Gboolean
	widgetRemoveMnemonicLabel   func(w *Widget, label *Widget)
	widgetRenderIcon            func(w *Widget, stockId string, size IconSize, detail string) *T.GdkPixbuf
	widgetReparent              func(w, newParent *Widget)
	widgetResetRcStyles         func(w *Widget)
	widgetResetShapes           func(w *Widget)
	widgetSendExpose            func(w *Widget, event *T.GdkEvent) int
	widgetSendFocusChange       func(w *Widget, event *T.GdkEvent) T.Gboolean
	widgetSet                   func(w *Widget, firstPropertyName string, v ...VArg)
	widgetSetAccelPath          func(w *Widget, accelPath string, accelGroup *AccelGroup)
	widgetSetAllocation         func(w *Widget, allocation *Allocation)
	widgetSetAppPaintable       func(w *Widget, appPaintable T.Gboolean)
	widgetSetCanDefault         func(w *Widget, canDefault T.Gboolean)
	widgetSetCanFocus           func(w *Widget, canFocus T.Gboolean)
	widgetSetChildVisible       func(w *Widget, isVisible T.Gboolean)
	widgetSetColormap           func(w *Widget, colormap *T.GdkColormap)
	widgetSetCompositeName      func(w *Widget, name string)
	widgetSetDirection          func(w *Widget, dir TextDirection)
	widgetSetDoubleBuffered     func(w *Widget, doubleBuffered T.Gboolean)
	widgetSetEvents             func(w *Widget, events int)
	widgetSetExtensionEvents    func(w *Widget, mode T.GdkExtensionMode)
	widgetSetHasTooltip         func(w *Widget, hasTooltip T.Gboolean)
	widgetSetHasWindow          func(w *Widget, hasWindow T.Gboolean)
	widgetSetMapped             func(w *Widget, mapped T.Gboolean)
	widgetSetName               func(w *Widget, name string)
	widgetSetNoShowAll          func(w *Widget, noShowAll T.Gboolean)
	widgetSetParent             func(w *Widget, parent *Widget)
	widgetSetParentWindow       func(w *Widget, parentWindow *T.GdkWindow)
	widgetSetRealized           func(w *Widget, realized T.Gboolean)
	widgetSetReceivesDefault    func(w *Widget, receivesDefault T.Gboolean)
	widgetSetRedrawOnAllocate   func(w *Widget, redrawOnAllocate T.Gboolean)
	widgetSetScrollAdjustments  func(w *Widget, hadjustment, vadjustment *Adjustment) T.Gboolean
	widgetSetSensitive          func(w *Widget, sensitive T.Gboolean)
	widgetSetSizeRequest        func(w *Widget, width int, height int)
	widgetSetState              func(w *Widget, state StateType)
	widgetSetStyle              func(w *Widget, style *Style)
	widgetSetTooltipMarkup      func(w *Widget, markup string)
	widgetSetTooltipText        func(w *Widget, text string)
	widgetSetTooltipWindow      func(w *Widget, customWindow *Window)
	widgetSetUposition          func(w *Widget, x int, y int)
	widgetSetUsize              func(w *Widget, width int, height int)
	widgetSetVisible            func(w *Widget, visible T.Gboolean)
	widgetSetWindow             func(w *Widget, window *T.GdkWindow)
	widgetShapeCombineMask      func(w *Widget, shapeMask *T.GdkBitmap, offsetX, offsetY int)
	widgetShow                  func(w *Widget)
	widgetShowAll               func(w *Widget)
	widgetShowNow               func(w *Widget)
	widgetSizeAllocate          func(w *Widget, allocation *Allocation)
	widgetSizeRequest           func(w *Widget, requisition *Requisition)
	widgetStyleGet              func(w *Widget, firstPropertyName string, v ...VArg)
	widgetStyleGetProperty      func(w *Widget, propertyName string, value *T.GValue)
	widgetStyleGetValist        func(w *Widget, firstPropertyName string, varArgs T.VaList)
	widgetThawChildNotify       func(w *Widget)
	widgetTriggerTooltipQuery   func(w *Widget)
	widgetUnmap                 func(w *Widget)
	widgetUnparent              func(w *Widget)
	widgetUnrealize             func(w *Widget)
	widgetUnref                 func(w *Widget)
)

func (w *Widget) Activate() T.Gboolean { return widgetActivate(w) }
func (w *Widget) AddAccelerator(accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags) {
	widgetAddAccelerator(w, accelSignal, accelGroup, accelKey, accelMods, accelFlags)
}
func (w *Widget) AddEvents(events int)           { widgetAddEvents(w, events) }
func (w *Widget) AddMnemonicLabel(label *Widget) { widgetAddMnemonicLabel(w, label) }
func (w *Widget) CanActivateAccel(signalId uint) T.Gboolean {
	return widgetCanActivateAccel(w, signalId)
}
func (w *Widget) ChildFocus(direction DirectionType) T.Gboolean {
	return widgetChildFocus(w, direction)
}
func (w *Widget) ChildNotify(childProperty string) { widgetChildNotify(w, childProperty) }
func (w *Widget) ClassPath(pathLength *uint, path, pathReversed **T.Gchar) {
	widgetClassPath(w, pathLength, path, pathReversed)
}
func (w *Widget) CreatePangoContext() *T.PangoContext { return widgetCreatePangoContext(w) }
func (w *Widget) CreatePangoLayout(text string) *T.PangoLayout {
	return widgetCreatePangoLayout(w, text)
}
func (w *Widget) Destroy()                               { widgetDestroy(w) }
func (w *Widget) Destroyed(widgetPointer **Widget)       { widgetDestroyed(w, widgetPointer) }
func (w *Widget) Draw(area *T.GdkRectangle)              { widgetDraw(w, area) }
func (w *Widget) EnsureStyle()                           { widgetEnsureStyle(w) }
func (w *Widget) ErrorBell()                             { widgetErrorBell(w) }
func (w *Widget) Event(event *T.GdkEvent) T.Gboolean     { return widgetEvent(w, event) }
func (w *Widget) FreezeChildNotify()                     { widgetFreezeChildNotify(w) }
func (w *Widget) GetAccessible() *T.AtkObject            { return widgetGetAccessible(w) }
func (w *Widget) GetAllocation(allocation *Allocation)   { widgetGetAllocation(w, allocation) }
func (w *Widget) GetAncestor(widgetType T.GType) *Widget { return widgetGetAncestor(w, widgetType) }
func (w *Widget) GetAppPaintable() T.Gboolean            { return widgetGetAppPaintable(w) }
func (w *Widget) GetCanDefault() T.Gboolean              { return widgetGetCanDefault(w) }
func (w *Widget) GetCanFocus() T.Gboolean                { return widgetGetCanFocus(w) }
func (w *Widget) GetChildRequisition(requisition *Requisition) {
	widgetGetChildRequisition(w, requisition)
}
func (w *Widget) GetChildVisible() T.Gboolean             { return widgetGetChildVisible(w) }
func (w *Widget) GetColormap() *T.GdkColormap             { return widgetGetColormap(w) }
func (w *Widget) GetCompositeName() string                { return widgetGetCompositeName(w) }
func (w *Widget) GetDirection() TextDirection             { return widgetGetDirection(w) }
func (w *Widget) GetDisplay() *T.GdkDisplay               { return widgetGetDisplay(w) }
func (w *Widget) GetDoubleBuffered() T.Gboolean           { return widgetGetDoubleBuffered(w) }
func (w *Widget) GetEvents() int                          { return widgetGetEvents(w) }
func (w *Widget) GetExtensionEvents() T.GdkExtensionMode  { return widgetGetExtensionEvents(w) }
func (w *Widget) GetHasTooltip() T.Gboolean               { return widgetGetHasTooltip(w) }
func (w *Widget) GetHasWindow() T.Gboolean                { return widgetGetHasWindow(w) }
func (w *Widget) GetMapped() T.Gboolean                   { return widgetGetMapped(w) }
func (w *Widget) GetModifierStyle() *RcStyle              { return widgetGetModifierStyle(w) }
func (w *Widget) GetName() string                         { return widgetGetName(w) }
func (w *Widget) GetNoShowAll() T.Gboolean                { return widgetGetNoShowAll(w) }
func (w *Widget) GetPangoContext() *T.PangoContext        { return widgetGetPangoContext(w) }
func (w *Widget) GetParent() *Widget                      { return widgetGetParent(w) }
func (w *Widget) GetParentWindow() *T.GdkWindow           { return widgetGetParentWindow(w) }
func (w *Widget) GetPointer(x *int, y *int)               { widgetGetPointer(w, x, y) }
func (w *Widget) GetRealized() T.Gboolean                 { return widgetGetRealized(w) }
func (w *Widget) GetReceivesDefault() T.Gboolean          { return widgetGetReceivesDefault(w) }
func (w *Widget) GetRequisition(requisition *Requisition) { widgetGetRequisition(w, requisition) }
func (w *Widget) GetRootWindow() *T.GdkWindow             { return widgetGetRootWindow(w) }
func (w *Widget) GetScreen() *T.GdkScreen                 { return widgetGetScreen(w) }
func (w *Widget) GetSensitive() T.Gboolean                { return widgetGetSensitive(w) }
func (w *Widget) GetSettings() *Settings                  { return widgetGetSettings(w) }
func (w *Widget) GetSizeRequest(width, height *int)       { widgetGetSizeRequest(w, width, height) }
func (w *Widget) GetSnapshot(clipRect *T.GdkRectangle) *T.GdkPixmap {
	return widgetGetSnapshot(w, clipRect)
}
func (w *Widget) GetState() StateType       { return widgetGetState(w) }
func (w *Widget) GetStyle() *Style          { return widgetGetStyle(w) }
func (w *Widget) GetTooltipMarkup() string  { return widgetGetTooltipMarkup(w) }
func (w *Widget) GetTooltipText() string    { return widgetGetTooltipText(w) }
func (w *Widget) GetTooltipWindow() *Window { return widgetGetTooltipWindow(w) }
func (w *Widget) GetToplevel() *Widget      { return widgetGetToplevel(w) }
func (w *Widget) GetVisible() T.Gboolean    { return widgetGetVisible(w) }
func (w *Widget) GetVisual() *T.GdkVisual   { return widgetGetVisual(w) }
func (w *Widget) GetWindow() *T.GdkWindow   { return widgetGetWindow(w) }
func (w *Widget) GrabDefault()              { widgetGrabDefault(w) }
func (w *Widget) GrabFocus()                { widgetGrabFocus(w) }
func (w *Widget) HasDefault() T.Gboolean    { return widgetHasDefault(w) }
func (w *Widget) HasFocus() T.Gboolean      { return widgetHasFocus(w) }
func (w *Widget) HasGrab() T.Gboolean       { return widgetHasGrab(w) }
func (w *Widget) HasRcStyle() T.Gboolean    { return widgetHasRcStyle(w) }
func (w *Widget) HasScreen() T.Gboolean     { return widgetHasScreen(w) }
func (w *Widget) Hide()                     { widgetHide(w) }
func (w *Widget) HideAll()                  { widgetHideAll(w) }
func (w *Widget) HideOnDelete() T.Gboolean  { return widgetHideOnDelete(w) }
func (w *Widget) InputShapeCombineMask(shapeMask *T.GdkBitmap, offsetX, offsetY int) {
	widgetInputShapeCombineMask(w, shapeMask, offsetX, offsetY)
}
func (w *Widget) Intersect(area, intersection *T.GdkRectangle) T.Gboolean {
	return widgetIntersect(w, area, intersection)
}
func (w *Widget) IsAncestor(ancestor *Widget) T.Gboolean { return widgetIsAncestor(w, ancestor) }
func (w *Widget) IsComposited() T.Gboolean               { return widgetIsComposited(w) }
func (w *Widget) IsDrawable() T.Gboolean                 { return widgetIsDrawable(w) }
func (w *Widget) IsFocus() T.Gboolean                    { return widgetIsFocus(w) }
func (w *Widget) IsSensitive() T.Gboolean                { return widgetIsSensitive(w) }
func (w *Widget) IsToplevel() T.Gboolean                 { return widgetIsToplevel(w) }
func (w *Widget) KeynavFailed(direction DirectionType) T.Gboolean {
	return widgetKeynavFailed(w, direction)
}
func (w *Widget) ListAccelClosures() *T.GList  { return widgetListAccelClosures(w) }
func (w *Widget) ListMnemonicLabels() *T.GList { return widgetListMnemonicLabels(w) }
func (w *Widget) Map()                         { widgetMap(w) }
func (w *Widget) MnemonicActivate(groupCycling T.Gboolean) T.Gboolean {
	return widgetMnemonicActivate(w, groupCycling)
}
func (w *Widget) ModifyBase(state StateType, color *T.GdkColor) {
	widgetModifyBase(w, state, color)
}
func (w *Widget) ModifyBg(state StateType, color *T.GdkColor) { widgetModifyBg(w, state, color) }
func (w *Widget) ModifyCursor(primary *T.GdkColor, secondary *T.GdkColor) {
	widgetModifyCursor(w, primary, secondary)
}
func (w *Widget) ModifyFg(state StateType, color *T.GdkColor) { widgetModifyFg(w, state, color) }
func (w *Widget) ModifyFont(fontDesc *T.PangoFontDescription) { widgetModifyFont(w, fontDesc) }
func (w *Widget) ModifyStyle(style *RcStyle)                  { widgetModifyStyle(w, style) }
func (w *Widget) ModifyText(state StateType, color *T.GdkColor) {
	widgetModifyText(w, state, color)
}
func (w *Widget) Path(pathLength *uint, path, pathReversed **T.Gchar) {
	widgetPath(w, pathLength, path, pathReversed)
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
func (w *Widget) RemoveMnemonicLabel(label *Widget) { widgetRemoveMnemonicLabel(w, label) }
func (w *Widget) RenderIcon(stockId string, size IconSize, detail string) *T.GdkPixbuf {
	return widgetRenderIcon(w, stockId, size, detail)
}
func (w *Widget) Reparent(newParent *Widget)                   { widgetReparent(w, newParent) }
func (w *Widget) ResetRcStyles()                               { widgetResetRcStyles(w) }
func (w *Widget) ResetShapes()                                 { widgetResetShapes(w) }
func (w *Widget) SendExpose(event *T.GdkEvent) int             { return widgetSendExpose(w, event) }
func (w *Widget) SendFocusChange(event *T.GdkEvent) T.Gboolean { return widgetSendFocusChange(w, event) }
func (w *Widget) Set(firstPropertyName string, v ...VArg)      { widgetSet(w, firstPropertyName, v) }
func (w *Widget) SetAccelPath(accelPath string, accelGroup *AccelGroup) {
	widgetSetAccelPath(w, accelPath, accelGroup)
}
func (w *Widget) SetAllocation(allocation *Allocation)    { widgetSetAllocation(w, allocation) }
func (w *Widget) SetAppPaintable(appPaintable T.Gboolean) { widgetSetAppPaintable(w, appPaintable) }
func (w *Widget) SetCanDefault(canDefault T.Gboolean)     { widgetSetCanDefault(w, canDefault) }
func (w *Widget) SetCanFocus(canFocus T.Gboolean)         { widgetSetCanFocus(w, canFocus) }
func (w *Widget) SetChildVisible(isVisible T.Gboolean)    { widgetSetChildVisible(w, isVisible) }
func (w *Widget) SetColormap(colormap *T.GdkColormap)     { widgetSetColormap(w, colormap) }
func (w *Widget) SetCompositeName(name string)            { widgetSetCompositeName(w, name) }
func (w *Widget) SetDirection(dir TextDirection)          { widgetSetDirection(w, dir) }
func (w *Widget) SetDoubleBuffered(doubleBuffered T.Gboolean) {
	widgetSetDoubleBuffered(w, doubleBuffered)
}
func (w *Widget) SetEvents(events int)                       { widgetSetEvents(w, events) }
func (w *Widget) SetExtensionEvents(mode T.GdkExtensionMode) { widgetSetExtensionEvents(w, mode) }
func (w *Widget) SetHasTooltip(hasTooltip T.Gboolean)        { widgetSetHasTooltip(w, hasTooltip) }
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
func (w *Widget) SetSensitive(sensitive T.Gboolean)     { widgetSetSensitive(w, sensitive) }
func (w *Widget) SetSizeRequest(width, height int)      { widgetSetSizeRequest(w, width, height) }
func (w *Widget) SetState(state StateType)              { widgetSetState(w, state) }
func (w *Widget) SetStyle(style *Style)                 { widgetSetStyle(w, style) }
func (w *Widget) SetTooltipMarkup(markup string)        { widgetSetTooltipMarkup(w, markup) }
func (w *Widget) SetTooltipText(text string)            { widgetSetTooltipText(w, text) }
func (w *Widget) SetTooltipWindow(customWindow *Window) { widgetSetTooltipWindow(w, customWindow) }
func (w *Widget) SetUposition(x, y int)                 { widgetSetUposition(w, x, y) }
func (w *Widget) SetUsize(width, height int)            { widgetSetUsize(w, width, height) }
func (w *Widget) SetVisible(visible T.Gboolean)         { widgetSetVisible(w, visible) }
func (w *Widget) SetWindow(window *T.GdkWindow)         { widgetSetWindow(w, window) }
func (w *Widget) ShapeCombineMask(shapeMask *T.GdkBitmap, offsetX, offsetY int) {
	widgetShapeCombineMask(w, shapeMask, offsetX, offsetY)
}
func (w *Widget) Show()                                { widgetShow(w) }
func (w *Widget) ShowAll()                             { widgetShowAll(w) }
func (w *Widget) ShowNow()                             { widgetShowNow(w) }
func (w *Widget) SizeAllocate(allocation *Allocation)  { widgetSizeAllocate(w, allocation) }
func (w *Widget) SizeRequest(requisition *Requisition) { widgetSizeRequest(w, requisition) }
func (w *Widget) StyleGet(firstPropertyName string, v ...VArg) {
	widgetStyleGet(w, firstPropertyName, v)
}
func (w *Widget) StyleGetProperty(propertyName string, value *T.GValue) {
	widgetStyleGetProperty(w, propertyName, value)
}
func (w *Widget) StyleGetValist(firstPropertyName string, varArgs T.VaList) {
	widgetStyleGetValist(w, firstPropertyName, varArgs)
}
func (w *Widget) ThawChildNotify()     { widgetThawChildNotify(w) }
func (w *Widget) TriggerTooltipQuery() { widgetTriggerTooltipQuery(w) }
func (w *Widget) Unmap()               { widgetUnmap(w) }
func (w *Widget) Unparent()            { widgetUnparent(w) }
func (w *Widget) Unrealize()           { widgetUnrealize(w) }
func (w *Widget) Unref()               { widgetUnref(w) }

type WidgetClass struct {
	ParentClass                    ObjectClass
	ActivateSignal                 uint
	SetScrollAdjustmentsSignal     uint
	DispatchChildPropertiesChanged func(
		widget *Widget,
		nPspecs uint,
		pspecs **T.GParamSpec)
	Show        func(widget *Widget)
	ShowAll     func(widget *Widget)
	Hide        func(widget *Widget)
	HideAll     func(widget *Widget)
	Map         func(widget *Widget)
	Unmap       func(widget *Widget)
	Realize     func(widget *Widget)
	Unrealize   func(widget *Widget)
	SizeRequest func(widget *Widget,
		requisition *Requisition)
	SizeAllocate func(widget *Widget,
		allocation *Allocation)
	StateChanged func(widget *Widget,
		previousState StateType)
	ParentSet func(widget *Widget,
		previousParent *Widget)
	HierarchyChanged func(widget *Widget,
		previousToplevel *Widget)
	StyleSet func(widget *Widget,
		previousStyle *Style)
	DirectionChanged func(widget *Widget,
		previousDirection TextDirection)
	GrabNotify func(widget *Widget,
		wasGrabbed T.Gboolean)
	ChildNotify func(widget *Widget,
		pspec *T.GParamSpec)
	MnemonicActivate func(widget *Widget,
		groupCycling T.Gboolean) T.Gboolean
	GrabFocus func(widget *Widget)
	Focus     func(widget *Widget,
		direction DirectionType) T.Gboolean
	Event func(widget *Widget,
		event *T.GdkEvent) T.Gboolean
	ButtonPressEvent func(widget *Widget,
		event *T.GdkEventButton) T.Gboolean
	ButtonReleaseEvent func(widget *Widget,
		event *T.GdkEventButton) T.Gboolean
	ScrollEvent func(widget *Widget,
		event *T.GdkEventScroll) T.Gboolean
	MotionNotifyEvent func(widget *Widget,
		event *T.GdkEventMotion) T.Gboolean
	DeleteEvent func(widget *Widget,
		event *T.GdkEventAny) T.Gboolean
	DestroyEvent func(widget *Widget,
		event *T.GdkEventAny) T.Gboolean
	ExposeEvent func(widget *Widget,
		event *T.GdkEventExpose) T.Gboolean
	KeyPressEvent func(widget *Widget,
		event *T.GdkEventKey) T.Gboolean
	KeyReleaseEvent func(widget *Widget,
		event *T.GdkEventKey) T.Gboolean
	EnterNotifyEvent func(widget *Widget,
		event *T.GdkEventCrossing) T.Gboolean
	LeaveNotifyEvent func(widget *Widget,
		event *T.GdkEventCrossing) T.Gboolean
	ConfigureEvent func(widget *Widget,
		event *T.GdkEventConfigure) T.Gboolean
	FocusInEvent func(widget *Widget,
		event *T.GdkEventFocus) T.Gboolean
	FocusOutEvent func(widget *Widget,
		event *T.GdkEventFocus) T.Gboolean
	MapEvent func(widget *Widget,
		event *T.GdkEventAny) T.Gboolean
	UnmapEvent func(widget *Widget,
		event *T.GdkEventAny) T.Gboolean
	PropertyNotifyEvent func(widget *Widget,
		event *T.GdkEventProperty) T.Gboolean
	SelectionClearEvent func(widget *Widget,
		event *T.GdkEventSelection) T.Gboolean
	SelectionRequestEvent func(widget *Widget,
		event *T.GdkEventSelection) T.Gboolean
	SelectionNotifyEvent func(widget *Widget,
		event *T.GdkEventSelection) T.Gboolean
	ProximityInEvent func(widget *Widget,
		event *T.GdkEventProximity) T.Gboolean
	ProximityOutEvent func(widget *Widget,
		event *T.GdkEventProximity) T.Gboolean
	VisibilityNotifyEvent func(widget *Widget,
		event *T.GdkEventVisibility) T.Gboolean
	ClientEvent func(widget *Widget,
		event *T.GdkEventClient) T.Gboolean
	NoExposeEvent func(widget *Widget,
		event *T.GdkEventAny) T.Gboolean
	WindowStateEvent func(widget *Widget,
		event *T.GdkEventWindowState) T.Gboolean
	SelectionGet func(widget *Widget,
		selectionData *SelectionData,
		info, time uint)
	SelectionReceived func(widget *Widget,
		selectionData *SelectionData,
		time uint)
	DragBegin func(widget *Widget,
		context *T.GdkDragContext)
	DragEnd func(widget *Widget,
		context *T.GdkDragContext)
	DragDataGet func(widget *Widget,
		context *T.GdkDragContext,
		selectionData *SelectionData,
		info, time uint)
	DragDataDelete func(widget *Widget,
		context *T.GdkDragContext)
	DragLeave func(widget *Widget,
		context *T.GdkDragContext,
		time uint)
	DragMotion func(widget *Widget,
		context *T.GdkDragContext,
		x, y int, time uint) T.Gboolean
	DragDrop func(widget *Widget,
		context *T.GdkDragContext,
		x, y int, time uint) T.Gboolean
	DragDataReceived func(widget *Widget,
		context *T.GdkDragContext,
		x, y int,
		selectionData *SelectionData,
		info, time uint)
	PopupMenu func(widget *Widget) T.Gboolean
	ShowHelp  func(widget *Widget,
		helpType WidgetHelpType) T.Gboolean
	GetAccessible func(widget *Widget) *T.AtkObject
	ScreenChanged func(widget *Widget,
		previousScreen *T.GdkScreen) T.Gboolean
	CanActivateAccel func(widget *Widget,
		signalId uint) T.Gboolean
	GrabBrokenEvent func(widget *Widget,
		event *T.GdkEventGrabBroken) T.Gboolean
	CompositedChanged func(widget *Widget)
	QueryTooltip      func(widget *Widget,
		x, y int, keyboardTooltip T.Gboolean,
		tooltip *Tooltip) T.Gboolean
	_, _, _ func()
}

var (
	widgetClassFindStyleProperty          func(w *WidgetClass, propertyName string) *T.GParamSpec
	widgetClassInstallStyleProperty       func(w *WidgetClass, pspec *T.GParamSpec)
	widgetClassInstallStylePropertyParser func(w *WidgetClass, pspec *T.GParamSpec, parser RcPropertyParser)
	widgetClassListStyleProperties        func(w *WidgetClass, nProperties *uint) **T.GParamSpec
)

func (w *WidgetClass) FindStyleProperty(propertyName string) *T.GParamSpec {
	return widgetClassFindStyleProperty(w, propertyName)
}
func (w *WidgetClass) InstallStyleProperty(pspec *T.GParamSpec) {
	widgetClassInstallStyleProperty(w, pspec)
}
func (w *WidgetClass) InstallStylePropertyParser(pspec *T.GParamSpec, parser RcPropertyParser) {
	widgetClassInstallStylePropertyParser(w, pspec, parser)
}
func (w *WidgetClass) ListStyleProperties(nProperties *uint) **T.GParamSpec {
	return widgetClassListStyleProperties(w, nProperties)
}

type WidgetHelpType T.Enum

const (
	WIDGET_HELP_TOOLTIP WidgetHelpType = iota
	WIDGET_HELP_WHATS_THIS
)

var WidgetHelpTypeGetType func() T.GType

type Window struct {
	Bin                   Bin
	Title                 *T.Gchar
	WmclassName           *T.Gchar
	WmclassClass          *T.Gchar
	WmRole                *T.Gchar
	FocusWidget           *Widget
	DefaultWidget         *Widget
	TransientParent       *Window
	GeometryInfo          *WindowGeometryInfo
	Frame                 *Window
	Group                 *WindowGroup
	ConfigureRequestCount uint16
	Bits                  uint
	// AllowShrink : 1
	// AllowGrow : 1
	// ConfigureNotifyReceived : 1
	// NeedDefaultPosition : 1
	// NeedDefaultSize : 1
	// Position : 3
	// Type : 4
	// HasUserRefCount : 1
	// HasFocus : 1
	// Modal : 1
	// DestroyWithParent : 1
	// HasFrame : 1
	// IconifyInitially : 1
	// StickInitially : 1
	// MaximizeInitially : 1
	// Decorated : 1
	// TypeHint : 3
	// Gravity : 5
	// IsActive : 1
	// HasToplevelFocus : 1
	FrameLeft          uint
	FrameTop           uint
	FrameRight         uint
	FrameBottom        uint
	KeysChangedHandler uint
	MnemonicModifier   T.GdkModifierType
	Screen             *T.GdkScreen
}

var (
	WindowGetType func() T.GType
	WindowNew     func(t WindowType) *Widget

	WindowSetDefaultIconList         func(list *T.GList)
	WindowGetDefaultIconList         func() *T.GList
	WindowSetDefaultIcon             func(icon *T.GdkPixbuf)
	WindowSetDefaultIconName         func(name string)
	WindowGetDefaultIconName         func() string
	WindowSetDefaultIconFromFile     func(filename string, err **T.GError) T.Gboolean
	WindowListToplevels              func() *T.GList
	WindowSetAutoStartupNotification func(setting T.Gboolean)

	windowActivateDefault       func(w *Window) T.Gboolean
	windowActivateFocus         func(w *Window) T.Gboolean
	windowActivateKey           func(w *Window, event *T.GdkEventKey) T.Gboolean
	windowAddAccelGroup         func(w *Window, accelGroup *AccelGroup)
	windowAddEmbeddedXid        func(w *Window, xid T.GdkNativeWindow)
	windowAddMnemonic           func(w *Window, keyval uint, target *Widget)
	windowBeginMoveDrag         func(w *Window, button, rootX, rootY int, timestamp T.GUint32)
	windowBeginResizeDrag       func(w *Window, edge T.GdkWindowEdge, button, rootX, rootY int, timestamp T.GUint32)
	windowDeiconify             func(w *Window)
	windowFullscreen            func(w *Window)
	windowGetAcceptFocus        func(w *Window) T.Gboolean
	windowGetDecorated          func(w *Window) T.Gboolean
	windowGetDefaultSize        func(w *Window, width, height *int)
	windowGetDefaultWidget      func(w *Window) *Widget
	windowGetDeletable          func(w *Window) T.Gboolean
	windowGetDestroyWithParent  func(w *Window) T.Gboolean
	windowGetFocus              func(w *Window) *Widget
	windowGetFocusOnMap         func(w *Window) T.Gboolean
	windowGetFrameDimensions    func(w *Window, left, top, right, bottom *int)
	windowGetGravity            func(w *Window) T.GdkGravity
	windowGetGroup              func(w *Window) *WindowGroup
	windowGetHasFrame           func(w *Window) T.Gboolean
	windowGetIcon               func(w *Window) *T.GdkPixbuf
	windowGetIconList           func(w *Window) *T.GList
	windowGetIconName           func(w *Window) string
	windowGetMnemonicModifier   func(w *Window) T.GdkModifierType
	windowGetMnemonicsVisible   func(w *Window) T.Gboolean
	windowGetModal              func(w *Window) T.Gboolean
	windowGetOpacity            func(w *Window) float64
	windowGetPosition           func(w *Window, rootX, rootY *int)
	windowGetResizable          func(w *Window) T.Gboolean
	windowGetRole               func(w *Window) string
	windowGetScreen             func(w *Window) *T.GdkScreen
	windowGetSize               func(w *Window, width, height *int)
	windowGetSkipPagerHint      func(w *Window) T.Gboolean
	windowGetSkipTaskbarHint    func(w *Window) T.Gboolean
	windowGetTitle              func(w *Window) string
	windowGetTransientFor       func(w *Window) *Window
	windowGetTypeHint           func(w *Window) T.GdkWindowTypeHint
	windowGetUrgencyHint        func(w *Window) T.Gboolean
	windowGetWindowType         func(w *Window) WindowType
	windowHasGroup              func(w *Window) T.Gboolean
	windowHasToplevelFocus      func(w *Window) T.Gboolean
	windowIconify               func(w *Window)
	windowIsActive              func(w *Window) T.Gboolean
	windowMaximize              func(w *Window)
	windowMnemonicActivate      func(w *Window, keyval uint, modifier T.GdkModifierType) T.Gboolean
	windowMove                  func(w *Window, x, y int)
	windowParseGeometry         func(w *Window, geometry string) T.Gboolean
	windowPresent               func(w *Window)
	windowPresentWithTime       func(w *Window, timestamp T.GUint32)
	windowPropagateKeyEvent     func(w *Window, event *T.GdkEventKey) T.Gboolean
	windowRemoveAccelGroup      func(w *Window, accelGroup *AccelGroup)
	windowRemoveEmbeddedXid     func(w *Window, xid T.GdkNativeWindow)
	windowRemoveMnemonic        func(w *Window, keyval uint, target *Widget)
	windowReshowWithInitialSize func(w *Window)
	windowResize                func(w *Window, width, height int)
	windowSetAcceptFocus        func(w *Window, setting T.Gboolean)
	windowSetDecorated          func(w *Window, setting T.Gboolean)
	windowSetDefault            func(w *Window, defaultWidget *Widget)
	windowSetDefaultSize        func(w *Window, width, height int)
	windowSetDeletable          func(w *Window, setting T.Gboolean)
	windowSetDestroyWithParent  func(w *Window, setting T.Gboolean)
	windowSetFocus              func(w *Window, focus *Widget)
	windowSetFocusOnMap         func(w *Window, setting T.Gboolean)
	windowSetFrameDimensions    func(w *Window, left, top, right, bottom int)
	windowSetGeometryHints      func(w *Window, geometryWidget *Widget, geometry *T.GdkGeometry, geomMask T.GdkWindowHints)
	windowSetGravity            func(w *Window, gravity T.GdkGravity)
	windowSetHasFrame           func(w *Window, setting T.Gboolean)
	windowSetIcon               func(w *Window, icon *T.GdkPixbuf)
	windowSetIconFromFile       func(w *Window, filename string, err **T.GError) T.Gboolean
	windowSetIconList           func(w *Window, list *T.GList)
	windowSetIconName           func(w *Window, name string)
	windowSetKeepAbove          func(w *Window, setting T.Gboolean)
	windowSetKeepBelow          func(w *Window, setting T.Gboolean)
	windowSetMnemonicModifier   func(w *Window, modifier T.GdkModifierType)
	windowSetMnemonicsVisible   func(w *Window, setting T.Gboolean)
	windowSetModal              func(w *Window, modal T.Gboolean)
	windowSetOpacity            func(w *Window, opacity float64)
	windowSetPolicy             func(w *Window, allowShrink, allowGrow, autoShrink int)
	windowSetPosition           func(w *Window, position WindowPosition)
	windowSetResizable          func(w *Window, resizable T.Gboolean)
	windowSetRole               func(w *Window, role string)
	windowSetScreen             func(w *Window, screen *T.GdkScreen)
	windowSetSkipPagerHint      func(w *Window, setting T.Gboolean)
	windowSetSkipTaskbarHint    func(w *Window, setting T.Gboolean)
	windowSetStartupId          func(w *Window, startupId string)
	windowSetTitle              func(w *Window, title string)
	windowSetTransientFor       func(w *Window, parent *Window)
	windowSetTypeHint           func(w *Window, hint T.GdkWindowTypeHint)
	windowSetUrgencyHint        func(w *Window, setting T.Gboolean)
	windowSetWmclass            func(w *Window, wmclassName, wmclassClass string)
	windowStick                 func(w *Window)
	windowUnfullscreen          func(w *Window)
	windowUnmaximize            func(w *Window)
	windowUnstick               func(w *Window)
)

func (w *Window) ActivateDefault() T.Gboolean { return windowActivateDefault(w) }
func (w *Window) ActivateFocus() T.Gboolean   { return windowActivateFocus(w) }
func (w *Window) ActivateKey(event *T.GdkEventKey) T.Gboolean {
	return windowActivateKey(w, event)
}
func (w *Window) AddAccelGroup(accelGroup *AccelGroup)    { windowAddAccelGroup(w, accelGroup) }
func (w *Window) AddEmbeddedXid(xid T.GdkNativeWindow)    { windowAddEmbeddedXid(w, xid) }
func (w *Window) AddMnemonic(keyval uint, target *Widget) { windowAddMnemonic(w, keyval, target) }
func (w *Window) BeginMoveDrag(button, rootX, rootY int, timestamp T.GUint32) {
	windowBeginMoveDrag(w, button, rootX, rootY, timestamp)
}
func (w *Window) BeginResizeDrag(edge T.GdkWindowEdge, button, rootX, rootY int, timestamp T.GUint32) {
	windowBeginResizeDrag(w, edge, button, rootX, rootY, timestamp)
}
func (w *Window) Deiconify()                 { windowDeiconify(w) }
func (w *Window) Fullscreen()                { windowFullscreen(w) }
func (w *Window) GetAcceptFocus() T.Gboolean { return windowGetAcceptFocus(w) }
func (w *Window) GetDecorated() T.Gboolean   { return windowGetDecorated(w) }
func (w *Window) GetDefaultSize(width, height *int) {
	windowGetDefaultSize(w, width, height)
}
func (w *Window) GetDefaultWidget() *Widget        { return windowGetDefaultWidget(w) }
func (w *Window) GetDeletable() T.Gboolean         { return windowGetDeletable(w) }
func (w *Window) GetDestroyWithParent() T.Gboolean { return windowGetDestroyWithParent(w) }
func (w *Window) GetFocus() *Widget                { return windowGetFocus(w) }
func (w *Window) GetFocusOnMap() T.Gboolean        { return windowGetFocusOnMap(w) }
func (w *Window) GetFrameDimensions(left, top, right, bottom *int) {
	windowGetFrameDimensions(w, left, top, right, bottom)
}
func (w *Window) GetGravity() T.GdkGravity               { return windowGetGravity(w) }
func (w *Window) GetGroup() *WindowGroup                 { return windowGetGroup(w) }
func (w *Window) GetHasFrame() T.Gboolean                { return windowGetHasFrame(w) }
func (w *Window) GetIcon() *T.GdkPixbuf                  { return windowGetIcon(w) }
func (w *Window) GetIconList() *T.GList                  { return windowGetIconList(w) }
func (w *Window) GetIconName() string                    { return windowGetIconName(w) }
func (w *Window) GetMnemonicModifier() T.GdkModifierType { return windowGetMnemonicModifier(w) }
func (w *Window) GetMnemonicsVisible() T.Gboolean        { return windowGetMnemonicsVisible(w) }
func (w *Window) GetModal() T.Gboolean                   { return windowGetModal(w) }
func (w *Window) GetOpacity() float64                    { return windowGetOpacity(w) }
func (w *Window) GetPosition(rootX, rootY *int)          { windowGetPosition(w, rootX, rootY) }
func (w *Window) GetResizable() T.Gboolean               { return windowGetResizable(w) }
func (w *Window) GetRole() string                        { return windowGetRole(w) }
func (w *Window) GetScreen() *T.GdkScreen                { return windowGetScreen(w) }
func (w *Window) GetSize(width, height *int)             { windowGetSize(w, width, height) }
func (w *Window) GetSkipPagerHint() T.Gboolean           { return windowGetSkipPagerHint(w) }
func (w *Window) GetSkipTaskbarHint() T.Gboolean         { return windowGetSkipTaskbarHint(w) }
func (w *Window) GetTitle() string                       { return windowGetTitle(w) }
func (w *Window) GetTransientFor() *Window               { return windowGetTransientFor(w) }
func (w *Window) GetTypeHint() T.GdkWindowTypeHint       { return windowGetTypeHint(w) }
func (w *Window) GetUrgencyHint() T.Gboolean             { return windowGetUrgencyHint(w) }
func (w *Window) GetWindowType() WindowType              { return windowGetWindowType(w) }
func (w *Window) HasGroup() T.Gboolean                   { return windowHasGroup(w) }
func (w *Window) HasToplevelFocus() T.Gboolean           { return windowHasToplevelFocus(w) }
func (w *Window) Iconify()                               { windowIconify(w) }
func (w *Window) IsActive() T.Gboolean                   { return windowIsActive(w) }
func (w *Window) Maximize()                              { windowMaximize(w) }
func (w *Window) MnemonicActivate(keyval uint, modifier T.GdkModifierType) T.Gboolean {
	return windowMnemonicActivate(w, keyval, modifier)
}
func (w *Window) Move(x int, y int)                        { windowMove(w, x, y) }
func (w *Window) ParseGeometry(geometry string) T.Gboolean { return windowParseGeometry(w, geometry) }
func (w *Window) Present()                                 { windowPresent(w) }
func (w *Window) PresentWithTime(timestamp T.GUint32)      { windowPresentWithTime(w, timestamp) }
func (w *Window) PropagateKeyEvent(event *T.GdkEventKey) T.Gboolean {
	return windowPropagateKeyEvent(w, event)
}
func (w *Window) RemoveAccelGroup(accelGroup *AccelGroup)    { windowRemoveAccelGroup(w, accelGroup) }
func (w *Window) RemoveEmbeddedXid(xid T.GdkNativeWindow)    { windowRemoveEmbeddedXid(w, xid) }
func (w *Window) RemoveMnemonic(keyval uint, target *Widget) { windowRemoveMnemonic(w, keyval, target) }
func (w *Window) ReshowWithInitialSize()                     { windowReshowWithInitialSize(w) }
func (w *Window) Resize(width, height int)                   { windowResize(w, width, height) }
func (w *Window) SetAcceptFocus(setting T.Gboolean)          { windowSetAcceptFocus(w, setting) }
func (w *Window) SetDecorated(setting T.Gboolean)            { windowSetDecorated(w, setting) }
func (w *Window) SetDefault(defaultWidget *Widget)           { windowSetDefault(w, defaultWidget) }
func (w *Window) SetDefaultSize(width, height int)           { windowSetDefaultSize(w, width, height) }
func (w *Window) SetDeletable(setting T.Gboolean)            { windowSetDeletable(w, setting) }
func (w *Window) SetDestroyWithParent(setting T.Gboolean)    { windowSetDestroyWithParent(w, setting) }
func (w *Window) SetFocus(focus *Widget)                     { windowSetFocus(w, focus) }
func (w *Window) SetFocusOnMap(setting T.Gboolean)           { windowSetFocusOnMap(w, setting) }
func (w *Window) SetFrameDimensions(left, top, right, bottom int) {
	windowSetFrameDimensions(w, left, top, right, bottom)
}
func (w *Window) SetGeometryHints(geometryWidget *Widget, geometry *T.GdkGeometry, geomMask T.GdkWindowHints) {
	windowSetGeometryHints(w, geometryWidget, geometry, geomMask)
}
func (w *Window) SetGravity(gravity T.GdkGravity) { windowSetGravity(w, gravity) }
func (w *Window) SetHasFrame(setting T.Gboolean)  { windowSetHasFrame(w, setting) }
func (w *Window) SetIcon(icon *T.GdkPixbuf)       { windowSetIcon(w, icon) }
func (w *Window) SetIconFromFile(filename string, err **T.GError) T.Gboolean {
	return windowSetIconFromFile(w, filename, err)
}
func (w *Window) SetIconList(list *T.GList)       { windowSetIconList(w, list) }
func (w *Window) SetIconName(name string)         { windowSetIconName(w, name) }
func (w *Window) SetKeepAbove(setting T.Gboolean) { windowSetKeepAbove(w, setting) }
func (w *Window) SetKeepBelow(setting T.Gboolean) { windowSetKeepBelow(w, setting) }
func (w *Window) SetMnemonicModifier(modifier T.GdkModifierType) {
	windowSetMnemonicModifier(w, modifier)
}
func (w *Window) SetMnemonicsVisible(setting T.Gboolean) { windowSetMnemonicsVisible(w, setting) }
func (w *Window) SetModal(modal T.Gboolean)              { windowSetModal(w, modal) }
func (w *Window) SetOpacity(opacity float64)             { windowSetOpacity(w, opacity) }
func (w *Window) SetPolicy(allowShrink, allowGrow, autoShrink int) {
	windowSetPolicy(w, allowShrink, allowGrow, autoShrink)
}
func (w *Window) SetPosition(position WindowPosition)   { windowSetPosition(w, position) }
func (w *Window) SetResizable(resizable T.Gboolean)     { windowSetResizable(w, resizable) }
func (w *Window) SetRole(role string)                   { windowSetRole(w, role) }
func (w *Window) SetScreen(screen *T.GdkScreen)         { windowSetScreen(w, screen) }
func (w *Window) SetSkipPagerHint(setting T.Gboolean)   { windowSetSkipPagerHint(w, setting) }
func (w *Window) SetSkipTaskbarHint(setting T.Gboolean) { windowSetSkipTaskbarHint(w, setting) }
func (w *Window) SetStartupId(startupId string)         { windowSetStartupId(w, startupId) }
func (w *Window) SetTitle(title string)                 { windowSetTitle(w, title) }
func (w *Window) SetTransientFor(parent *Window)        { windowSetTransientFor(w, parent) }
func (w *Window) SetTypeHint(hint T.GdkWindowTypeHint)  { windowSetTypeHint(w, hint) }
func (w *Window) SetUrgencyHint(setting T.Gboolean)     { windowSetUrgencyHint(w, setting) }
func (w *Window) SetWmclass(wmclassName, wmclassClass string) {
	windowSetWmclass(w, wmclassName, wmclassClass)
}
func (w *Window) Stick()        { windowStick(w) }
func (w *Window) Unfullscreen() { windowUnfullscreen(w) }
func (w *Window) Unmaximize()   { windowUnmaximize(w) }
func (w *Window) Unstick()      { windowUnstick(w) }

type WindowGeometryInfo struct{}

type WindowGroup struct {
	Parent T.GObject
	Grabs  *T.GSList
}

var (
	WindowGroupGetType func() T.GType
	WindowGroupNew     func() *WindowGroup

	windowGroupAddWindow      func(w *WindowGroup, window *Window)
	windowGroupGetCurrentGrab func(w *WindowGroup) *Widget
	windowGroupListWindows    func(w *WindowGroup) *T.GList
	windowGroupRemoveWindow   func(w *WindowGroup, window *Window)
)

func (w *WindowGroup) AddWindow(window *Window)    { windowGroupAddWindow(w, window) }
func (w *WindowGroup) GetCurrentGrab() *Widget     { return windowGroupGetCurrentGrab(w) }
func (w *WindowGroup) ListWindows() *T.GList       { return windowGroupListWindows(w) }
func (w *WindowGroup) RemoveWindow(window *Window) { windowGroupRemoveWindow(w, window) }

type WindowPosition T.Enum

const (
	WIN_POS_NONE WindowPosition = iota
	WIN_POS_CENTER
	WIN_POS_MOUSE
	WIN_POS_CENTER_ALWAYS
	WIN_POS_CENTER_ON_PARENT
)

var WindowPositionGetType func() T.GType

type WindowType T.Enum

const (
	WINDOW_TOPLEVEL WindowType = iota
	WINDOW_POPUP
)

var WindowTypeGetType func() T.GType
