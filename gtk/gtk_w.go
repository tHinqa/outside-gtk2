// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	A "github.com/tHinqa/outside-gtk2/atk"
	D "github.com/tHinqa/outside-gtk2/gdk"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
	"unsafe"
)

type WrapMode Enum

const (
	WRAP_NONE WrapMode = iota
	WRAP_CHAR
	WRAP_WORD
	WRAP_WORD_CHAR
)

var WrapModeGetType func() O.Type

type Widget struct {
	Object       Object
	PrivateFlags uint16
	State        uint8
	SavedState   uint8
	Name         *T.Gchar
	Style        *Style
	Requisition  Requisition
	Allocation   Allocation
	Window       *D.Window
	Parent       *Widget
}

var (
	WidgetGetType func() O.Type
	WidgetNew     func(t O.Type, firstPropertyName string, v ...VArg) *Widget

	WidgetFlagsGetType func() O.Type

	WidgetGetDefaultColormap  func() *D.Colormap
	WidgetGetDefaultDirection func() TextDirection
	WidgetGetDefaultStyle     func() *Style
	WidgetGetDefaultVisual    func() *D.Visual
	WidgetPopColormap         func()
	WidgetPopCompositeChild   func()
	WidgetPushColormap        func(cmap *D.Colormap)
	WidgetPushCompositeChild  func()
	WidgetSetDefaultColormap  func(colormap *D.Colormap)
	WidgetSetDefaultDirection func(dir TextDirection)

	//TODO(t): Methods?
	WidgetTranslateCoordinates func(srcWidget *Widget, destWidget *Widget, srcX int, srcY int, destX *int, destY *int) bool
	WidgetStyleAttach          func(style *Widget)

	widgetActivate              func(w *Widget) bool
	widgetAddAccelerator        func(w *Widget, accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags)
	widgetAddEvents             func(w *Widget, events int)
	widgetAddMnemonicLabel      func(w *Widget, label *Widget)
	widgetCanActivateAccel      func(w *Widget, signalId uint) bool
	widgetChildFocus            func(w *Widget, direction DirectionType) bool
	widgetChildNotify           func(w *Widget, childProperty string)
	widgetClassPath             func(w *Widget, pathLength *uint, path, pathReversed **T.Gchar)
	widgetCreatePangoContext    func(w *Widget) *P.Context
	widgetCreatePangoLayout     func(w *Widget, text string) *P.Layout
	widgetDestroy               func(w *Widget)
	widgetDestroyed             func(w *Widget, widgetPointer **Widget)
	widgetDraw                  func(w *Widget, area *D.Rectangle)
	widgetEnsureStyle           func(w *Widget)
	widgetErrorBell             func(w *Widget)
	widgetEvent                 func(w *Widget, event *D.Event) bool
	widgetFreezeChildNotify     func(w *Widget)
	widgetGetAccessible         func(w *Widget) *A.Object
	widgetGetAllocation         func(w *Widget, allocation *Allocation)
	widgetGetAncestor           func(w *Widget, widgetType O.Type) *Widget
	widgetGetAppPaintable       func(w *Widget) bool
	widgetGetCanDefault         func(w *Widget) bool
	widgetGetCanFocus           func(w *Widget) bool
	widgetGetChildRequisition   func(w *Widget, requisition *Requisition)
	widgetGetChildVisible       func(w *Widget) bool
	widgetGetColormap           func(w *Widget) *D.Colormap
	widgetGetCompositeName      func(w *Widget) string
	widgetGetDirection          func(w *Widget) TextDirection
	widgetGetDisplay            func(w *Widget) *D.Display
	widgetGetDoubleBuffered     func(w *Widget) bool
	widgetGetEvents             func(w *Widget) int
	widgetGetExtensionEvents    func(w *Widget) D.ExtensionMode
	widgetGetHasTooltip         func(w *Widget) bool
	widgetGetHasWindow          func(w *Widget) bool
	widgetGetMapped             func(w *Widget) bool
	widgetGetModifierStyle      func(w *Widget) *RcStyle
	widgetGetName               func(w *Widget) string
	widgetGetNoShowAll          func(w *Widget) bool
	widgetGetPangoContext       func(w *Widget) *P.Context
	widgetGetParent             func(w *Widget) *Widget
	widgetGetParentWindow       func(w *Widget) *D.Window
	widgetGetPointer            func(w *Widget, x *int, y *int)
	widgetGetRealized           func(w *Widget) bool
	widgetGetReceivesDefault    func(w *Widget) bool
	widgetGetRequisition        func(w *Widget, requisition *Requisition)
	widgetGetRootWindow         func(w *Widget) *D.Window
	widgetGetScreen             func(w *Widget) *D.Screen
	widgetGetSensitive          func(w *Widget) bool
	widgetGetSettings           func(w *Widget) *Settings
	widgetGetSizeRequest        func(w *Widget, width *int, height *int)
	widgetGetSnapshot           func(w *Widget, clipRect *D.Rectangle) *D.Pixmap
	widgetGetState              func(w *Widget) StateType
	widgetGetStyle              func(w *Widget) *Style
	widgetGetTooltipMarkup      func(w *Widget) string
	widgetGetTooltipText        func(w *Widget) string
	widgetGetTooltipWindow      func(w *Widget) *Window
	widgetGetToplevel           func(w *Widget) *Widget
	widgetGetVisible            func(w *Widget) bool
	widgetGetVisual             func(w *Widget) *D.Visual
	widgetGetWindow             func(w *Widget) *D.Window
	widgetGrabDefault           func(w *Widget)
	widgetGrabFocus             func(w *Widget)
	widgetHasDefault            func(w *Widget) bool
	widgetHasFocus              func(w *Widget) bool
	widgetHasGrab               func(w *Widget) bool
	widgetHasRcStyle            func(w *Widget) bool
	widgetHasScreen             func(w *Widget) bool
	widgetHide                  func(w *Widget)
	widgetHideAll               func(w *Widget)
	widgetHideOnDelete          func(w *Widget) bool
	widgetInputShapeCombineMask func(w *Widget, shapeMask *T.GdkBitmap, offsetX, offsetY int)
	widgetIntersect             func(w *Widget, area, intersection *D.Rectangle) bool
	widgetIsAncestor            func(w *Widget, ancestor *Widget) bool
	widgetIsComposited          func(w *Widget) bool
	widgetIsDrawable            func(w *Widget) bool
	widgetIsFocus               func(w *Widget) bool
	widgetIsSensitive           func(w *Widget) bool
	widgetIsToplevel            func(w *Widget) bool
	widgetKeynavFailed          func(w *Widget, direction DirectionType) bool
	widgetListAccelClosures     func(w *Widget) *T.GList
	widgetListMnemonicLabels    func(w *Widget) *T.GList
	widgetMap                   func(w *Widget)
	widgetMnemonicActivate      func(w *Widget, groupCycling bool) bool
	widgetModifyBase            func(w *Widget, state StateType, color *D.Color)
	widgetModifyBg              func(w *Widget, state StateType, color *D.Color)
	widgetModifyCursor          func(w *Widget, primary *D.Color, secondary *D.Color)
	widgetModifyFg              func(w *Widget, state StateType, color *D.Color)
	widgetModifyFont            func(w *Widget, fontDesc *P.FontDescription)
	widgetModifyStyle           func(w *Widget, style *RcStyle)
	widgetModifyText            func(w *Widget, state StateType, color *D.Color)
	widgetPath                  func(w *Widget, pathLength *uint, path, pathReversed **T.Gchar)
	widgetQueueClear            func(w *Widget)
	widgetQueueClearArea        func(w *Widget, x, y, width, height int)
	widgetQueueDraw             func(w *Widget)
	widgetQueueDrawArea         func(w *Widget, x, y, width, height int)
	widgetQueueResize           func(w *Widget)
	widgetQueueResizeNoRedraw   func(w *Widget)
	widgetRealize               func(w *Widget)
	widgetRef                   func(w *Widget) *Widget
	widgetRegionIntersect       func(w *Widget, region *D.Region) *D.Region
	widgetRemoveAccelerator     func(w *Widget, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType) bool
	widgetRemoveMnemonicLabel   func(w *Widget, label *Widget)
	widgetRenderIcon            func(w *Widget, stockId string, size IconSize, detail string) *D.Pixbuf
	widgetReparent              func(w, newParent *Widget)
	widgetResetRcStyles         func(w *Widget)
	widgetResetShapes           func(w *Widget)
	widgetSendExpose            func(w *Widget, event *D.Event) int
	widgetSendFocusChange       func(w *Widget, event *D.Event) bool
	widgetSet                   func(w *Widget, firstPropertyName string, v ...VArg)
	widgetSetAccelPath          func(w *Widget, accelPath string, accelGroup *AccelGroup)
	widgetSetAllocation         func(w *Widget, allocation *Allocation)
	widgetSetAppPaintable       func(w *Widget, appPaintable bool)
	widgetSetCanDefault         func(w *Widget, canDefault bool)
	widgetSetCanFocus           func(w *Widget, canFocus bool)
	widgetSetChildVisible       func(w *Widget, isVisible bool)
	widgetSetColormap           func(w *Widget, colormap *D.Colormap)
	widgetSetCompositeName      func(w *Widget, name string)
	widgetSetDirection          func(w *Widget, dir TextDirection)
	widgetSetDoubleBuffered     func(w *Widget, doubleBuffered bool)
	widgetSetEvents             func(w *Widget, events int)
	widgetSetExtensionEvents    func(w *Widget, mode D.ExtensionMode)
	widgetSetHasTooltip         func(w *Widget, hasTooltip bool)
	widgetSetHasWindow          func(w *Widget, hasWindow bool)
	widgetSetMapped             func(w *Widget, mapped bool)
	widgetSetName               func(w *Widget, name string)
	widgetSetNoShowAll          func(w *Widget, noShowAll bool)
	widgetSetParent             func(w *Widget, parent *Widget)
	widgetSetParentWindow       func(w *Widget, parentWindow *D.Window)
	widgetSetRealized           func(w *Widget, realized bool)
	widgetSetReceivesDefault    func(w *Widget, receivesDefault bool)
	widgetSetRedrawOnAllocate   func(w *Widget, redrawOnAllocate bool)
	widgetSetScrollAdjustments  func(w *Widget, hadjustment, vadjustment *Adjustment) bool
	widgetSetSensitive          func(w *Widget, sensitive bool)
	widgetSetSizeRequest        func(w *Widget, width int, height int)
	widgetSetState              func(w *Widget, state StateType)
	widgetSetStyle              func(w *Widget, style *Style)
	widgetSetTooltipMarkup      func(w *Widget, markup string)
	widgetSetTooltipText        func(w *Widget, text string)
	widgetSetTooltipWindow      func(w *Widget, customWindow *Window)
	widgetSetUposition          func(w *Widget, x int, y int)
	widgetSetUsize              func(w *Widget, width int, height int)
	widgetSetVisible            func(w *Widget, visible bool)
	widgetSetWindow             func(w *Widget, window *D.Window)
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

func (w *Widget) AsAboutDialog() *AboutDialog { return (*AboutDialog)(unsafe.Pointer(w)) }
func (w *Widget) AsContainer() *Container     { return (*Container)(unsafe.Pointer(w)) }
func (w *Widget) AsPointer() T.Gpointer       { return (T.Gpointer)(unsafe.Pointer(w)) }
func (w *Widget) AsWindow() *Window           { return (*Window)(unsafe.Pointer(w)) }

func (w *Widget) Activate() bool { return widgetActivate(w) }
func (w *Widget) AddAccelerator(accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags) {
	widgetAddAccelerator(w, accelSignal, accelGroup, accelKey, accelMods, accelFlags)
}
func (w *Widget) AddEvents(events int)           { widgetAddEvents(w, events) }
func (w *Widget) AddMnemonicLabel(label *Widget) { widgetAddMnemonicLabel(w, label) }
func (w *Widget) CanActivateAccel(signalId uint) bool {
	return widgetCanActivateAccel(w, signalId)
}
func (w *Widget) ChildFocus(direction DirectionType) bool {
	return widgetChildFocus(w, direction)
}
func (w *Widget) ChildNotify(childProperty string) { widgetChildNotify(w, childProperty) }
func (w *Widget) ClassPath(pathLength *uint, path, pathReversed **T.Gchar) {
	widgetClassPath(w, pathLength, path, pathReversed)
}
func (w *Widget) CreatePangoContext() *P.Context { return widgetCreatePangoContext(w) }
func (w *Widget) CreatePangoLayout(text string) *P.Layout {
	return widgetCreatePangoLayout(w, text)
}
func (w *Widget) Destroy()                              { widgetDestroy(w) }
func (w *Widget) Destroyed(widgetPointer **Widget)      { widgetDestroyed(w, widgetPointer) }
func (w *Widget) Draw(area *D.Rectangle)                { widgetDraw(w, area) }
func (w *Widget) EnsureStyle()                          { widgetEnsureStyle(w) }
func (w *Widget) ErrorBell()                            { widgetErrorBell(w) }
func (w *Widget) Event(event *D.Event) bool             { return widgetEvent(w, event) }
func (w *Widget) FreezeChildNotify()                    { widgetFreezeChildNotify(w) }
func (w *Widget) GetAccessible() *A.Object              { return widgetGetAccessible(w) }
func (w *Widget) GetAllocation(allocation *Allocation)  { widgetGetAllocation(w, allocation) }
func (w *Widget) GetAncestor(widgetType O.Type) *Widget { return widgetGetAncestor(w, widgetType) }
func (w *Widget) GetAppPaintable() bool                 { return widgetGetAppPaintable(w) }
func (w *Widget) GetCanDefault() bool                   { return widgetGetCanDefault(w) }
func (w *Widget) GetCanFocus() bool                     { return widgetGetCanFocus(w) }
func (w *Widget) GetChildRequisition(requisition *Requisition) {
	widgetGetChildRequisition(w, requisition)
}
func (w *Widget) GetChildVisible() bool                   { return widgetGetChildVisible(w) }
func (w *Widget) GetColormap() *D.Colormap                { return widgetGetColormap(w) }
func (w *Widget) GetCompositeName() string                { return widgetGetCompositeName(w) }
func (w *Widget) GetDirection() TextDirection             { return widgetGetDirection(w) }
func (w *Widget) GetDisplay() *D.Display                  { return widgetGetDisplay(w) }
func (w *Widget) GetDoubleBuffered() bool                 { return widgetGetDoubleBuffered(w) }
func (w *Widget) GetEvents() int                          { return widgetGetEvents(w) }
func (w *Widget) GetExtensionEvents() D.ExtensionMode     { return widgetGetExtensionEvents(w) }
func (w *Widget) GetHasTooltip() bool                     { return widgetGetHasTooltip(w) }
func (w *Widget) GetHasWindow() bool                      { return widgetGetHasWindow(w) }
func (w *Widget) GetMapped() bool                         { return widgetGetMapped(w) }
func (w *Widget) GetModifierStyle() *RcStyle              { return widgetGetModifierStyle(w) }
func (w *Widget) GetName() string                         { return widgetGetName(w) }
func (w *Widget) GetNoShowAll() bool                      { return widgetGetNoShowAll(w) }
func (w *Widget) GetPangoContext() *P.Context             { return widgetGetPangoContext(w) }
func (w *Widget) GetParent() *Widget                      { return widgetGetParent(w) }
func (w *Widget) GetParentWindow() *D.Window              { return widgetGetParentWindow(w) }
func (w *Widget) GetPointer(x *int, y *int)               { widgetGetPointer(w, x, y) }
func (w *Widget) GetRealized() bool                       { return widgetGetRealized(w) }
func (w *Widget) GetReceivesDefault() bool                { return widgetGetReceivesDefault(w) }
func (w *Widget) GetRequisition(requisition *Requisition) { widgetGetRequisition(w, requisition) }
func (w *Widget) GetRootWindow() *D.Window                { return widgetGetRootWindow(w) }
func (w *Widget) GetScreen() *D.Screen                    { return widgetGetScreen(w) }
func (w *Widget) GetSensitive() bool                      { return widgetGetSensitive(w) }
func (w *Widget) GetSettings() *Settings                  { return widgetGetSettings(w) }
func (w *Widget) GetSizeRequest(width, height *int)       { widgetGetSizeRequest(w, width, height) }
func (w *Widget) GetSnapshot(clipRect *D.Rectangle) *D.Pixmap {
	return widgetGetSnapshot(w, clipRect)
}
func (w *Widget) GetState() StateType       { return widgetGetState(w) }
func (w *Widget) GetStyle() *Style          { return widgetGetStyle(w) }
func (w *Widget) GetTooltipMarkup() string  { return widgetGetTooltipMarkup(w) }
func (w *Widget) GetTooltipText() string    { return widgetGetTooltipText(w) }
func (w *Widget) GetTooltipWindow() *Window { return widgetGetTooltipWindow(w) }
func (w *Widget) GetToplevel() *Widget      { return widgetGetToplevel(w) }
func (w *Widget) GetVisible() bool          { return widgetGetVisible(w) }
func (w *Widget) GetVisual() *D.Visual      { return widgetGetVisual(w) }
func (w *Widget) GetWindow() *D.Window      { return widgetGetWindow(w) }
func (w *Widget) GrabDefault()              { widgetGrabDefault(w) }
func (w *Widget) GrabFocus()                { widgetGrabFocus(w) }
func (w *Widget) HasDefault() bool          { return widgetHasDefault(w) }
func (w *Widget) HasFocus() bool            { return widgetHasFocus(w) }
func (w *Widget) HasGrab() bool             { return widgetHasGrab(w) }
func (w *Widget) HasRcStyle() bool          { return widgetHasRcStyle(w) }
func (w *Widget) HasScreen() bool           { return widgetHasScreen(w) }
func (w *Widget) Hide()                     { widgetHide(w) }
func (w *Widget) HideAll()                  { widgetHideAll(w) }
func (w *Widget) HideOnDelete() bool        { return widgetHideOnDelete(w) }
func (w *Widget) InputShapeCombineMask(shapeMask *T.GdkBitmap, offsetX, offsetY int) {
	widgetInputShapeCombineMask(w, shapeMask, offsetX, offsetY)
}
func (w *Widget) Intersect(area, intersection *D.Rectangle) bool {
	return widgetIntersect(w, area, intersection)
}
func (w *Widget) IsAncestor(ancestor *Widget) bool { return widgetIsAncestor(w, ancestor) }
func (w *Widget) IsComposited() bool               { return widgetIsComposited(w) }
func (w *Widget) IsDrawable() bool                 { return widgetIsDrawable(w) }
func (w *Widget) IsFocus() bool                    { return widgetIsFocus(w) }
func (w *Widget) IsSensitive() bool                { return widgetIsSensitive(w) }
func (w *Widget) IsToplevel() bool                 { return widgetIsToplevel(w) }
func (w *Widget) KeynavFailed(direction DirectionType) bool {
	return widgetKeynavFailed(w, direction)
}
func (w *Widget) ListAccelClosures() *T.GList  { return widgetListAccelClosures(w) }
func (w *Widget) ListMnemonicLabels() *T.GList { return widgetListMnemonicLabels(w) }
func (w *Widget) Map()                         { widgetMap(w) }
func (w *Widget) MnemonicActivate(groupCycling bool) bool {
	return widgetMnemonicActivate(w, groupCycling)
}
func (w *Widget) ModifyBase(state StateType, color *D.Color) {
	widgetModifyBase(w, state, color)
}
func (w *Widget) ModifyBg(state StateType, color *D.Color) { widgetModifyBg(w, state, color) }
func (w *Widget) ModifyCursor(primary *D.Color, secondary *D.Color) {
	widgetModifyCursor(w, primary, secondary)
}
func (w *Widget) ModifyFg(state StateType, color *D.Color) { widgetModifyFg(w, state, color) }
func (w *Widget) ModifyFont(fontDesc *P.FontDescription)   { widgetModifyFont(w, fontDesc) }
func (w *Widget) ModifyStyle(style *RcStyle)               { widgetModifyStyle(w, style) }
func (w *Widget) ModifyText(state StateType, color *D.Color) {
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
func (w *Widget) RegionIntersect(region *D.Region) *D.Region {
	return widgetRegionIntersect(w, region)
}
func (w *Widget) RemoveAccelerator(accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType) bool {
	return widgetRemoveAccelerator(w, accelGroup, accelKey, accelMods)
}
func (w *Widget) RemoveMnemonicLabel(label *Widget) { widgetRemoveMnemonicLabel(w, label) }
func (w *Widget) RenderIcon(stockId string, size IconSize, detail string) *D.Pixbuf {
	return widgetRenderIcon(w, stockId, size, detail)
}
func (w *Widget) Reparent(newParent *Widget)              { widgetReparent(w, newParent) }
func (w *Widget) ResetRcStyles()                          { widgetResetRcStyles(w) }
func (w *Widget) ResetShapes()                            { widgetResetShapes(w) }
func (w *Widget) SendExpose(event *D.Event) int           { return widgetSendExpose(w, event) }
func (w *Widget) SendFocusChange(event *D.Event) bool     { return widgetSendFocusChange(w, event) }
func (w *Widget) Set(firstPropertyName string, v ...VArg) { widgetSet(w, firstPropertyName, v) }
func (w *Widget) SetAccelPath(accelPath string, accelGroup *AccelGroup) {
	widgetSetAccelPath(w, accelPath, accelGroup)
}
func (w *Widget) SetAllocation(allocation *Allocation) { widgetSetAllocation(w, allocation) }
func (w *Widget) SetAppPaintable(appPaintable bool)    { widgetSetAppPaintable(w, appPaintable) }
func (w *Widget) SetCanDefault(canDefault bool)        { widgetSetCanDefault(w, canDefault) }
func (w *Widget) SetCanFocus(canFocus bool)            { widgetSetCanFocus(w, canFocus) }
func (w *Widget) SetChildVisible(isVisible bool)       { widgetSetChildVisible(w, isVisible) }
func (w *Widget) SetColormap(colormap *D.Colormap)     { widgetSetColormap(w, colormap) }
func (w *Widget) SetCompositeName(name string)         { widgetSetCompositeName(w, name) }
func (w *Widget) SetDirection(dir TextDirection)       { widgetSetDirection(w, dir) }
func (w *Widget) SetDoubleBuffered(doubleBuffered bool) {
	widgetSetDoubleBuffered(w, doubleBuffered)
}
func (w *Widget) SetEvents(events int)                    { widgetSetEvents(w, events) }
func (w *Widget) SetExtensionEvents(mode D.ExtensionMode) { widgetSetExtensionEvents(w, mode) }
func (w *Widget) SetHasTooltip(hasTooltip bool)           { widgetSetHasTooltip(w, hasTooltip) }
func (w *Widget) SetHasWindow(hasWindow bool)             { widgetSetHasWindow(w, hasWindow) }
func (w *Widget) SetMapped(mapped bool)                   { widgetSetMapped(w, mapped) }
func (w *Widget) SetName(name string)                     { widgetSetName(w, name) }
func (w *Widget) SetNoShowAll(noShowAll bool)             { widgetSetNoShowAll(w, noShowAll) }
func (w *Widget) SetParent(parent *Widget)                { widgetSetParent(w, parent) }
func (w *Widget) SetParentWindow(parentWindow *D.Window)  { widgetSetParentWindow(w, parentWindow) }
func (w *Widget) SetRealized(realized bool)               { widgetSetRealized(w, realized) }
func (w *Widget) SetReceivesDefault(receivesDefault bool) {
	widgetSetReceivesDefault(w, receivesDefault)
}
func (w *Widget) SetRedrawOnAllocate(redrawOnAllocate bool) {
	widgetSetRedrawOnAllocate(w, redrawOnAllocate)
}
func (w *Widget) SetScrollAdjustments(hadjustment, vadjustment *Adjustment) bool {
	return widgetSetScrollAdjustments(w, hadjustment, vadjustment)
}
func (w *Widget) SetSensitive(sensitive bool)           { widgetSetSensitive(w, sensitive) }
func (w *Widget) SetSizeRequest(width, height int)      { widgetSetSizeRequest(w, width, height) }
func (w *Widget) SetState(state StateType)              { widgetSetState(w, state) }
func (w *Widget) SetStyle(style *Style)                 { widgetSetStyle(w, style) }
func (w *Widget) SetTooltipMarkup(markup string)        { widgetSetTooltipMarkup(w, markup) }
func (w *Widget) SetTooltipText(text string)            { widgetSetTooltipText(w, text) }
func (w *Widget) SetTooltipWindow(customWindow *Window) { widgetSetTooltipWindow(w, customWindow) }
func (w *Widget) SetUposition(x, y int)                 { widgetSetUposition(w, x, y) }
func (w *Widget) SetUsize(width, height int)            { widgetSetUsize(w, width, height) }
func (w *Widget) SetVisible(visible bool)               { widgetSetVisible(w, visible) }
func (w *Widget) SetWindow(window *D.Window)            { widgetSetWindow(w, window) }
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
		wasGrabbed bool)
	ChildNotify func(widget *Widget,
		pspec *T.GParamSpec)
	MnemonicActivate func(widget *Widget,
		groupCycling bool) bool
	GrabFocus func(widget *Widget)
	Focus     func(widget *Widget,
		direction DirectionType) bool
	Event func(widget *Widget,
		event *D.Event) bool
	ButtonPressEvent func(widget *Widget,
		event *D.EventButton) bool
	ButtonReleaseEvent func(widget *Widget,
		event *D.EventButton) bool
	ScrollEvent func(widget *Widget,
		event *D.EventScroll) bool
	MotionNotifyEvent func(widget *Widget,
		event *D.EventMotion) bool
	DeleteEvent func(widget *Widget,
		event *D.EventAny) bool
	DestroyEvent func(widget *Widget,
		event *D.EventAny) bool
	ExposeEvent func(widget *Widget,
		event *D.EventExpose) bool
	KeyPressEvent func(widget *Widget,
		event *D.EventKey) bool
	KeyReleaseEvent func(widget *Widget,
		event *D.EventKey) bool
	EnterNotifyEvent func(widget *Widget,
		event *D.EventCrossing) bool
	LeaveNotifyEvent func(widget *Widget,
		event *D.EventCrossing) bool
	ConfigureEvent func(widget *Widget,
		event *D.EventConfigure) bool
	FocusInEvent func(widget *Widget,
		event *D.EventFocus) bool
	FocusOutEvent func(widget *Widget,
		event *D.EventFocus) bool
	MapEvent func(widget *Widget,
		event *D.EventAny) bool
	UnmapEvent func(widget *Widget,
		event *D.EventAny) bool
	PropertyNotifyEvent func(widget *Widget,
		event *D.EventProperty) bool
	SelectionClearEvent func(widget *Widget,
		event *D.EventSelection) bool
	SelectionRequestEvent func(widget *Widget,
		event *D.EventSelection) bool
	SelectionNotifyEvent func(widget *Widget,
		event *D.EventSelection) bool
	ProximityInEvent func(widget *Widget,
		event *D.EventProximity) bool
	ProximityOutEvent func(widget *Widget,
		event *D.EventProximity) bool
	VisibilityNotifyEvent func(widget *Widget,
		event *D.EventVisibility) bool
	ClientEvent func(widget *Widget,
		event *D.EventClient) bool
	NoExposeEvent func(widget *Widget,
		event *D.EventAny) bool
	WindowStateEvent func(widget *Widget,
		event *D.EventWindowState) bool
	SelectionGet func(widget *Widget,
		selectionData *SelectionData,
		info, time uint)
	SelectionReceived func(widget *Widget,
		selectionData *SelectionData,
		time uint)
	DragBegin func(widget *Widget,
		context *D.DragContext)
	DragEnd func(widget *Widget,
		context *D.DragContext)
	DragDataGet func(widget *Widget,
		context *D.DragContext,
		selectionData *SelectionData,
		info, time uint)
	DragDataDelete func(widget *Widget,
		context *D.DragContext)
	DragLeave func(widget *Widget,
		context *D.DragContext,
		time uint)
	DragMotion func(widget *Widget,
		context *D.DragContext,
		x, y int, time uint) bool
	DragDrop func(widget *Widget,
		context *D.DragContext,
		x, y int, time uint) bool
	DragDataReceived func(widget *Widget,
		context *D.DragContext,
		x, y int,
		selectionData *SelectionData,
		info, time uint)
	PopupMenu func(widget *Widget) bool
	ShowHelp  func(widget *Widget,
		helpType WidgetHelpType) bool
	GetAccessible func(widget *Widget) *A.Object
	ScreenChanged func(widget *Widget,
		previousScreen *D.Screen) bool
	CanActivateAccel func(widget *Widget,
		signalId uint) bool
	GrabBrokenEvent func(widget *Widget,
		event *D.EventGrabBroken) bool
	CompositedChanged func(widget *Widget)
	QueryTooltip      func(widget *Widget,
		x, y int, keyboardTooltip bool,
		tooltip *Tooltip) bool
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

type WidgetHelpType Enum

const (
	WIDGET_HELP_TOOLTIP WidgetHelpType = iota
	WIDGET_HELP_WHATS_THIS
)

var WidgetHelpTypeGetType func() O.Type

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
	Screen             *D.Screen
}

var (
	WindowGetType func() O.Type
	WindowNew     func(t WindowType) *Widget

	WindowGetDefaultIconList         func() *T.GList
	WindowGetDefaultIconName         func() string // 2.16
	WindowListToplevels              func() *T.GList
	WindowSetAutoStartupNotification func(setting bool)                         // 2.2
	WindowSetDefaultIcon             func(icon *D.Pixbuf)                       // 2.4
	WindowSetDefaultIconFromFile     func(filename string, err **T.GError) bool // 2.2
	WindowSetDefaultIconList         func(list *T.GList)
	WindowSetDefaultIconName         func(name string) // 2.6

	windowActivateDefault       func(w *Window) bool
	windowActivateFocus         func(w *Window) bool
	windowActivateKey           func(w *Window, event *D.EventKey) bool
	windowAddAccelGroup         func(w *Window, accelGroup *AccelGroup)
	windowAddEmbeddedXid        func(w *Window, xid T.GdkNativeWindow)
	windowAddMnemonic           func(w *Window, keyval uint, target *Widget)
	windowBeginMoveDrag         func(w *Window, button, rootX, rootY int, timestamp T.GUint32)
	windowBeginResizeDrag       func(w *Window, edge D.WindowEdge, button, rootX, rootY int, timestamp T.GUint32)
	windowDeiconify             func(w *Window)
	windowFullscreen            func(w *Window)      // 2.2
	windowGetAcceptFocus        func(w *Window) bool // 2.4
	windowGetDecorated          func(w *Window) bool
	windowGetDefaultSize        func(w *Window, width, height *int)
	windowGetDefaultWidget      func(w *Window) *Widget // 2.14
	windowGetDeletable          func(w *Window) bool    // 2.10
	windowGetDestroyWithParent  func(w *Window) bool
	windowGetFocus              func(w *Window) *Widget
	windowGetFocusOnMap         func(w *Window) bool // 2.6
	windowGetFrameDimensions    func(w *Window, left, top, right, bottom *int)
	windowGetGravity            func(w *Window) D.Gravity
	windowGetGroup              func(w *Window) *WindowGroup // 2.10
	windowGetHasFrame           func(w *Window) bool
	windowGetIcon               func(w *Window) *D.Pixbuf
	windowGetIconList           func(w *Window) *T.GList
	windowGetIconName           func(w *Window) string // 2.6
	windowGetMnemonicModifier   func(w *Window) T.GdkModifierType
	windowGetMnemonicsVisible   func(w *Window) bool
	windowGetModal              func(w *Window) bool
	windowGetOpacity            func(w *Window) float64 // 2.12
	windowGetPosition           func(w *Window, rootX, rootY *int)
	windowGetResizable          func(w *Window) bool
	windowGetRole               func(w *Window) string
	windowGetScreen             func(w *Window) *D.Screen // 2.2
	windowGetSize               func(w *Window, width, height *int)
	windowGetSkipPagerHint      func(w *Window) bool // 2.2
	windowGetSkipTaskbarHint    func(w *Window) bool // 2.2
	windowGetTitle              func(w *Window) string
	windowGetTransientFor       func(w *Window) *Window
	windowGetTypeHint           func(w *Window) D.WindowTypeHint
	windowGetUrgencyHint        func(w *Window) bool // 2.8
	windowGetWindowType         func(w *Window) WindowType
	windowHasGroup              func(w *Window) bool
	windowHasToplevelFocus      func(w *Window) bool // 2.4
	windowIconify               func(w *Window)
	windowIsActive              func(w *Window) bool // 2.4
	windowMaximize              func(w *Window)
	windowMnemonicActivate      func(w *Window, keyval uint, modifier T.GdkModifierType) bool
	windowMove                  func(w *Window, x, y int)
	windowParseGeometry         func(w *Window, geometry string) bool
	windowPresent               func(w *Window)
	windowPresentWithTime       func(w *Window, timestamp T.GUint32) // 2.8
	windowPropagateKeyEvent     func(w *Window, event *D.EventKey) bool
	windowRemoveAccelGroup      func(w *Window, accelGroup *AccelGroup)
	windowRemoveEmbeddedXid     func(w *Window, xid T.GdkNativeWindow)
	windowRemoveMnemonic        func(w *Window, keyval uint, target *Widget)
	windowReshowWithInitialSize func(w *Window)
	windowResize                func(w *Window, width, height int)
	windowSetAcceptFocus        func(w *Window, setting bool) // 2.4
	windowSetDecorated          func(w *Window, setting bool)
	windowSetDefault            func(w *Window, defaultWidget *Widget)
	windowSetDefaultSize        func(w *Window, width, height int)
	windowSetDeletable          func(w *Window, setting bool) // 2.10
	windowSetDestroyWithParent  func(w *Window, setting bool)
	windowSetFocus              func(w *Window, focus *Widget)
	windowSetFocusOnMap         func(w *Window, setting bool) // 2.6
	windowSetFrameDimensions    func(w *Window, left, top, right, bottom int)
	windowSetGeometryHints      func(w *Window, geometryWidget *Widget, geometry *D.Geometry, geomMask D.WindowHints)
	windowSetGravity            func(w *Window, gravity D.Gravity)
	windowSetHasFrame           func(w *Window, setting bool)
	windowSetIcon               func(w *Window, icon *D.Pixbuf)
	windowSetIconFromFile       func(w *Window, filename string, err **T.GError) bool // 2.2
	windowSetIconList           func(w *Window, list *T.GList)
	windowSetIconName           func(w *Window, name string)  // 2.6
	windowSetKeepAbove          func(w *Window, setting bool) // 2.4
	windowSetKeepBelow          func(w *Window, setting bool) // 2.4
	windowSetMnemonicModifier   func(w *Window, modifier T.GdkModifierType)
	windowSetMnemonicsVisible   func(w *Window, setting bool)
	windowSetModal              func(w *Window, modal bool)
	windowSetOpacity            func(w *Window, opacity float64) // 2.12
	// Deprecated windowSetPolicy func(w *Window, allowShrink, allowGrow, autoShrink int)
	windowSetPosition        func(w *Window, position WindowPosition)
	windowSetResizable       func(w *Window, resizable bool)
	windowSetRole            func(w *Window, role string)
	windowSetScreen          func(w *Window, screen *D.Screen) // 2.2
	windowSetSkipPagerHint   func(w *Window, setting bool)     // 2.2
	windowSetSkipTaskbarHint func(w *Window, setting bool)     // 2.2
	windowSetStartupId       func(w *Window, startupId string) // 2.12
	windowSetTitle           func(w *Window, title string)
	windowSetTransientFor    func(w *Window, parent *Window)
	windowSetTypeHint        func(w *Window, hint D.WindowTypeHint)
	windowSetUrgencyHint     func(w *Window, setting bool) // 2.8
	windowSetWmclass         func(w *Window, wmclassName, wmclassClass string)
	windowStick              func(w *Window)
	windowUnfullscreen       func(w *Window)
	windowUnmaximize         func(w *Window)
	windowUnstick            func(w *Window)
)

func NewWindow(t WindowType) (wt *Widget, w *Window) {
	wt = WindowNew(t)
	w = (*Window)(unsafe.Pointer(wt))
	return
}

func (w *Window) AsContainer() *Container { return (*Container)(unsafe.Pointer(w)) }

func (w *Window) ActivateDefault() bool { return windowActivateDefault(w) }
func (w *Window) ActivateFocus() bool   { return windowActivateFocus(w) }
func (w *Window) ActivateKey(event *D.EventKey) bool {
	return windowActivateKey(w, event)
}
func (w *Window) AddAccelGroup(accelGroup *AccelGroup)    { windowAddAccelGroup(w, accelGroup) }
func (w *Window) AddEmbeddedXid(xid T.GdkNativeWindow)    { windowAddEmbeddedXid(w, xid) }
func (w *Window) AddMnemonic(keyval uint, target *Widget) { windowAddMnemonic(w, keyval, target) }
func (w *Window) BeginMoveDrag(button, rootX, rootY int, timestamp T.GUint32) {
	windowBeginMoveDrag(w, button, rootX, rootY, timestamp)
}
func (w *Window) BeginResizeDrag(edge D.WindowEdge, button, rootX, rootY int, timestamp T.GUint32) {
	windowBeginResizeDrag(w, edge, button, rootX, rootY, timestamp)
}
func (w *Window) Deiconify()        { windowDeiconify(w) }
func (w *Window) Fullscreen()       { windowFullscreen(w) }
func (w *Window) AcceptFocus() bool { return windowGetAcceptFocus(w) }
func (w *Window) Decorated() bool   { return windowGetDecorated(w) }
func (w *Window) DefaultSize(width, height *int) {
	windowGetDefaultSize(w, width, height)
}
func (w *Window) GetDefaultWidget() *Widget { return windowGetDefaultWidget(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) Deletable() bool           { return windowGetDeletable(w) }
func (w *Window) DestroyWithParent() bool   { return windowGetDestroyWithParent(w) }
func (w *Window) Focus() *Widget            { return windowGetFocus(w) }
func (w *Window) FocusOnMap() bool          { return windowGetFocusOnMap(w) }
func (w *Window) FrameDimensions(left, top, right, bottom *int) {
	windowGetFrameDimensions(w, left, top, right, bottom)
}
func (w *Window) Gravity() D.Gravity                     { return windowGetGravity(w) }
func (w *Window) GetGroup() *WindowGroup                 { return windowGetGroup(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) HasFrame() bool                         { return windowGetHasFrame(w) }
func (w *Window) Icon() *D.Pixbuf                        { return windowGetIcon(w) }
func (w *Window) IconList() *T.GList                     { return windowGetIconList(w) }
func (w *Window) IconName() string                       { return windowGetIconName(w) }
func (w *Window) GetMnemonicModifier() T.GdkModifierType { return windowGetMnemonicModifier(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) MnemonicsVisible() bool                 { return windowGetMnemonicsVisible(w) }
func (w *Window) Modal() bool                            { return windowGetModal(w) }
func (w *Window) Opacity() float64                       { return windowGetOpacity(w) }
func (w *Window) Position(rootX, rootY *int)             { windowGetPosition(w, rootX, rootY) }
func (w *Window) Resizable() bool                        { return windowGetResizable(w) }
func (w *Window) Role() string                           { return windowGetRole(w) }
func (w *Window) GetScreen() *D.Screen                   { return windowGetScreen(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) Size(width, height *int)                { windowGetSize(w, width, height) }
func (w *Window) SkipPagerHint() bool                    { return windowGetSkipPagerHint(w) }
func (w *Window) SkipTaskbarHint() bool                  { return windowGetSkipTaskbarHint(w) }
func (w *Window) GetTitle() string                       { return windowGetTitle(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) TransientFor() *Window                  { return windowGetTransientFor(w) }
func (w *Window) TypeHint() D.WindowTypeHint             { return windowGetTypeHint(w) }
func (w *Window) UrgencyHint() bool                      { return windowGetUrgencyHint(w) }
func (w *Window) WindowType() WindowType                 { return windowGetWindowType(w) }
func (w *Window) HasGroup() bool                         { return windowHasGroup(w) }
func (w *Window) HasToplevelFocus() bool                 { return windowHasToplevelFocus(w) }
func (w *Window) Iconify()                               { windowIconify(w) }
func (w *Window) IsActive() bool                         { return windowIsActive(w) }
func (w *Window) Maximize()                              { windowMaximize(w) }
func (w *Window) MnemonicActivate(keyval uint, modifier T.GdkModifierType) bool {
	return windowMnemonicActivate(w, keyval, modifier)
}
func (w *Window) Move(x int, y int)                   { windowMove(w, x, y) }
func (w *Window) ParseGeometry(geometry string) bool  { return windowParseGeometry(w, geometry) }
func (w *Window) Present()                            { windowPresent(w) }
func (w *Window) PresentWithTime(timestamp T.GUint32) { windowPresentWithTime(w, timestamp) }
func (w *Window) PropagateKeyEvent(event *D.EventKey) bool {
	return windowPropagateKeyEvent(w, event)
}
func (w *Window) RemoveAccelGroup(accelGroup *AccelGroup)    { windowRemoveAccelGroup(w, accelGroup) }
func (w *Window) RemoveEmbeddedXid(xid T.GdkNativeWindow)    { windowRemoveEmbeddedXid(w, xid) }
func (w *Window) RemoveMnemonic(keyval uint, target *Widget) { windowRemoveMnemonic(w, keyval, target) }
func (w *Window) ReshowWithInitialSize()                     { windowReshowWithInitialSize(w) }
func (w *Window) Resize(width, height int)                   { windowResize(w, width, height) }
func (w *Window) SetAcceptFocus(setting bool)                { windowSetAcceptFocus(w, setting) }
func (w *Window) SetDecorated(setting bool)                  { windowSetDecorated(w, setting) }
func (w *Window) SetDefault(defaultWidget *Widget)           { windowSetDefault(w, defaultWidget) }
func (w *Window) SetDefaultSize(width, height int)           { windowSetDefaultSize(w, width, height) }
func (w *Window) SetDeletable(setting bool)                  { windowSetDeletable(w, setting) }
func (w *Window) SetDestroyWithParent(setting bool)          { windowSetDestroyWithParent(w, setting) }
func (w *Window) SetFocus(focus *Widget)                     { windowSetFocus(w, focus) }
func (w *Window) SetFocusOnMap(setting bool)                 { windowSetFocusOnMap(w, setting) }
func (w *Window) SetFrameDimensions(left, top, right, bottom int) {
	windowSetFrameDimensions(w, left, top, right, bottom)
}
func (w *Window) SetGeometryHints(geometryWidget *Widget, geometry *D.Geometry, geomMask D.WindowHints) {
	windowSetGeometryHints(w, geometryWidget, geometry, geomMask)
}
func (w *Window) SetGravity(gravity D.Gravity) { windowSetGravity(w, gravity) }
func (w *Window) SetHasFrame(setting bool)     { windowSetHasFrame(w, setting) }
func (w *Window) SetIcon(icon *D.Pixbuf)       { windowSetIcon(w, icon) }
func (w *Window) SetIconFromFile(filename string, err **T.GError) bool {
	return windowSetIconFromFile(w, filename, err)
}
func (w *Window) SetIconList(list *T.GList) { windowSetIconList(w, list) }
func (w *Window) SetIconName(name string)   { windowSetIconName(w, name) }
func (w *Window) SetKeepAbove(setting bool) { windowSetKeepAbove(w, setting) }
func (w *Window) SetKeepBelow(setting bool) { windowSetKeepBelow(w, setting) }
func (w *Window) SetMnemonicModifier(modifier T.GdkModifierType) {
	windowSetMnemonicModifier(w, modifier)
}
func (w *Window) SetMnemonicsVisible(setting bool) { windowSetMnemonicsVisible(w, setting) }
func (w *Window) SetModal(modal bool)              { windowSetModal(w, modal) }
func (w *Window) SetOpacity(opacity float64)       { windowSetOpacity(w, opacity) }

// Deprecated
// func (w *Window) SetPolicy(allowShrink, allowGrow, autoShrink int) {
// 	windowSetPolicy(w, allowShrink, allowGrow, autoShrink)
// }
func (w *Window) SetPosition(position WindowPosition) { windowSetPosition(w, position) }
func (w *Window) SetResizable(resizable bool)         { windowSetResizable(w, resizable) }
func (w *Window) SetRole(role string)                 { windowSetRole(w, role) }
func (w *Window) SetScreen(screen *D.Screen)          { windowSetScreen(w, screen) }
func (w *Window) SetSkipPagerHint(setting bool)       { windowSetSkipPagerHint(w, setting) }
func (w *Window) SetSkipTaskbarHint(setting bool)     { windowSetSkipTaskbarHint(w, setting) }
func (w *Window) SetStartupId(startupId string)       { windowSetStartupId(w, startupId) }
func (w *Window) SetTitle(title string)               { windowSetTitle(w, title) }
func (w *Window) SetTransientFor(parent *Window)      { windowSetTransientFor(w, parent) }
func (w *Window) SetTypeHint(hint D.WindowTypeHint)   { windowSetTypeHint(w, hint) }
func (w *Window) SetUrgencyHint(setting bool)         { windowSetUrgencyHint(w, setting) }
func (w *Window) SetWmclass(wmclassName, wmclassClass string) {
	windowSetWmclass(w, wmclassName, wmclassClass)
}
func (w *Window) Stick()        { windowStick(w) }
func (w *Window) Unfullscreen() { windowUnfullscreen(w) }
func (w *Window) Unmaximize()   { windowUnmaximize(w) }
func (w *Window) Unstick()      { windowUnstick(w) }

type WindowGeometryInfo struct{}

type WindowGroup struct {
	Parent O.Object
	Grabs  *T.GSList
}

var (
	WindowGroupGetType func() O.Type
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

type WindowPosition Enum

const (
	WIN_POS_NONE WindowPosition = iota
	WIN_POS_CENTER
	WIN_POS_MOUSE
	WIN_POS_CENTER_ALWAYS
	WIN_POS_CENTER_ON_PARENT
)

var WindowPositionGetType func() O.Type

type WindowType Enum

const (
	WINDOW_TOPLEVEL WindowType = iota
	WINDOW_POPUP
)

var WindowTypeGetType func() O.Type

/*
properties            ver  default                     actions
"has-toplevel-focus"       false                       read
"is-active"                false                       read
"accept-focus"             true                        read/write
"allow-grow"               true                        read/write
"allow-shrink"             false                       read/write
"decorated"           2.4  true                        read/write
"default-height"           -1                          read/write
"default-width"            -1                          read/write
"deletable"           2.10 true                        read/write
"destroy-with-parent"      false                       read/write
"focus-on-map"             true                        read/write
"gravity"             2.4  gdk.GRAVITY_NORTH_WEST      read/write
"icon"                     ???                         read/write
"icon-name"                nil                         read/write
"modal"                    false                       read/write
"opacity"                  1.0                         read/write
"resizable"                true                        read/write
"role"                     nil                         read/write
"screen"                   ???                         read/write
"skip-pager-hint"          false                       read/write
"skip-taskbar-hint"        false                       read/write
"title"                    nil                         read/write
"transient-for"       2.10 ???                         read/write/construct
"type"                     gtk.WINDOW_TOPLEVEL         read/write/construct only
"type-hint"                gdk.WINDOW_TYPE_HINT_NORMAL read/write
"urgency-hint"             false                       read/write
"window-position"          gtk.WIN_POS_NONE            read/write
"startup-id"          2.12 nil                         write

signals
"activate-default"  run last/action
"activate-focus"    run last/action
"frame-event"       run last
"keys-changed"      run first
"set-focus"         run last
*/
