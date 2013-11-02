// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	A "github.com/tHinqa/outside-gtk2/atk"
	D "github.com/tHinqa/outside-gtk2/gdk"
	L "github.com/tHinqa/outside-gtk2/glib"
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

	WidgetActivate              func(w *Widget) bool
	WidgetAddAccelerator        func(w *Widget, accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags)
	WidgetAddEvents             func(w *Widget, events int)
	WidgetAddMnemonicLabel      func(w *Widget, label *Widget)
	WidgetCanActivateAccel      func(w *Widget, signalId uint) bool
	WidgetChildFocus            func(w *Widget, direction DirectionType) bool
	WidgetChildNotify           func(w *Widget, childProperty string)
	WidgetClassPath             func(w *Widget, pathLength *uint, path, pathReversed **T.Gchar)
	WidgetCreatePangoContext    func(w *Widget) *P.Context
	WidgetCreatePangoLayout     func(w *Widget, text string) *P.Layout
	WidgetDestroy               func(w *Widget)
	WidgetDestroyed             func(w *Widget, widgetPointer **Widget)
	WidgetDraw                  func(w *Widget, area *D.Rectangle)
	WidgetEnsureStyle           func(w *Widget)
	WidgetErrorBell             func(w *Widget)
	WidgetEvent                 func(w *Widget, event *D.Event) bool
	WidgetFreezeChildNotify     func(w *Widget)
	WidgetGetAccessible         func(w *Widget) *A.Object
	WidgetGetAllocation         func(w *Widget, allocation *Allocation)
	WidgetGetAncestor           func(w *Widget, widgetType O.Type) *Widget
	WidgetGetAppPaintable       func(w *Widget) bool
	WidgetGetCanDefault         func(w *Widget) bool
	WidgetGetCanFocus           func(w *Widget) bool
	WidgetGetChildRequisition   func(w *Widget, requisition *Requisition)
	WidgetGetChildVisible       func(w *Widget) bool
	WidgetGetColormap           func(w *Widget) *D.Colormap
	WidgetGetCompositeName      func(w *Widget) string
	WidgetGetDirection          func(w *Widget) TextDirection
	WidgetGetDisplay            func(w *Widget) *D.Display
	WidgetGetDoubleBuffered     func(w *Widget) bool
	WidgetGetEvents             func(w *Widget) int
	WidgetGetExtensionEvents    func(w *Widget) D.ExtensionMode
	WidgetGetHasTooltip         func(w *Widget) bool
	WidgetGetHasWindow          func(w *Widget) bool
	WidgetGetMapped             func(w *Widget) bool
	WidgetGetModifierStyle      func(w *Widget) *RcStyle
	WidgetGetName               func(w *Widget) string
	WidgetGetNoShowAll          func(w *Widget) bool
	WidgetGetPangoContext       func(w *Widget) *P.Context
	WidgetGetParent             func(w *Widget) *Widget
	WidgetGetParentWindow       func(w *Widget) *D.Window
	WidgetGetPointer            func(w *Widget, x *int, y *int)
	WidgetGetRealized           func(w *Widget) bool
	WidgetGetReceivesDefault    func(w *Widget) bool
	WidgetGetRequisition        func(w *Widget, requisition *Requisition)
	WidgetGetRootWindow         func(w *Widget) *D.Window
	WidgetGetScreen             func(w *Widget) *D.Screen
	WidgetGetSensitive          func(w *Widget) bool
	WidgetGetSettings           func(w *Widget) *Settings
	WidgetGetSizeRequest        func(w *Widget, width *int, height *int)
	WidgetGetSnapshot           func(w *Widget, clipRect *D.Rectangle) *D.Pixmap
	WidgetGetState              func(w *Widget) StateType
	WidgetGetStyle              func(w *Widget) *Style
	WidgetGetTooltipMarkup      func(w *Widget) string
	WidgetGetTooltipText        func(w *Widget) string
	WidgetGetTooltipWindow      func(w *Widget) *Window
	WidgetGetToplevel           func(w *Widget) *Widget
	WidgetGetVisible            func(w *Widget) bool
	WidgetGetVisual             func(w *Widget) *D.Visual
	WidgetGetWindow             func(w *Widget) *D.Window
	WidgetGrabDefault           func(w *Widget)
	WidgetGrabFocus             func(w *Widget)
	WidgetHasDefault            func(w *Widget) bool
	WidgetHasFocus              func(w *Widget) bool
	WidgetHasGrab               func(w *Widget) bool
	WidgetHasRcStyle            func(w *Widget) bool
	WidgetHasScreen             func(w *Widget) bool
	WidgetHide                  func(w *Widget)
	WidgetHideAll               func(w *Widget)
	WidgetHideOnDelete          func(w *Widget) bool
	WidgetInputShapeCombineMask func(w *Widget, shapeMask *T.GdkBitmap, offsetX, offsetY int)
	WidgetIntersect             func(w *Widget, area, intersection *D.Rectangle) bool
	WidgetIsAncestor            func(w *Widget, ancestor *Widget) bool
	WidgetIsComposited          func(w *Widget) bool
	WidgetIsDrawable            func(w *Widget) bool
	WidgetIsFocus               func(w *Widget) bool
	WidgetIsSensitive           func(w *Widget) bool
	WidgetIsToplevel            func(w *Widget) bool
	WidgetKeynavFailed          func(w *Widget, direction DirectionType) bool
	WidgetListAccelClosures     func(w *Widget) *L.List
	WidgetListMnemonicLabels    func(w *Widget) *L.List
	WidgetMap                   func(w *Widget)
	WidgetMnemonicActivate      func(w *Widget, groupCycling bool) bool
	WidgetModifyBase            func(w *Widget, state StateType, color *D.Color)
	WidgetModifyBg              func(w *Widget, state StateType, color *D.Color)
	WidgetModifyCursor          func(w *Widget, primary *D.Color, secondary *D.Color)
	WidgetModifyFg              func(w *Widget, state StateType, color *D.Color)
	WidgetModifyFont            func(w *Widget, fontDesc *P.FontDescription)
	WidgetModifyStyle           func(w *Widget, style *RcStyle)
	WidgetModifyText            func(w *Widget, state StateType, color *D.Color)
	WidgetPath                  func(w *Widget, pathLength *uint, path, pathReversed **T.Gchar)
	WidgetQueueClear            func(w *Widget)
	WidgetQueueClearArea        func(w *Widget, x, y, width, height int)
	WidgetQueueDraw             func(w *Widget)
	WidgetQueueDrawArea         func(w *Widget, x, y, width, height int)
	WidgetQueueResize           func(w *Widget)
	WidgetQueueResizeNoRedraw   func(w *Widget)
	WidgetRealize               func(w *Widget)
	WidgetRef                   func(w *Widget) *Widget
	WidgetRegionIntersect       func(w *Widget, region *D.Region) *D.Region
	WidgetRemoveAccelerator     func(w *Widget, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType) bool
	WidgetRemoveMnemonicLabel   func(w *Widget, label *Widget)
	WidgetRenderIcon            func(w *Widget, stockId string, size IconSize, detail string) *D.Pixbuf
	WidgetReparent              func(w, newParent *Widget)
	WidgetResetRcStyles         func(w *Widget)
	WidgetResetShapes           func(w *Widget)
	WidgetSendExpose            func(w *Widget, event *D.Event) int
	WidgetSendFocusChange       func(w *Widget, event *D.Event) bool
	WidgetSet                   func(w *Widget, firstPropertyName string, v ...VArg)
	WidgetSetAccelPath          func(w *Widget, accelPath string, accelGroup *AccelGroup)
	WidgetSetAllocation         func(w *Widget, allocation *Allocation)
	WidgetSetAppPaintable       func(w *Widget, appPaintable bool)
	WidgetSetCanDefault         func(w *Widget, canDefault bool)
	WidgetSetCanFocus           func(w *Widget, canFocus bool)
	WidgetSetChildVisible       func(w *Widget, isVisible bool)
	WidgetSetColormap           func(w *Widget, colormap *D.Colormap)
	WidgetSetCompositeName      func(w *Widget, name string)
	WidgetSetDirection          func(w *Widget, dir TextDirection)
	WidgetSetDoubleBuffered     func(w *Widget, doubleBuffered bool)
	WidgetSetEvents             func(w *Widget, events int)
	WidgetSetExtensionEvents    func(w *Widget, mode D.ExtensionMode)
	WidgetSetHasTooltip         func(w *Widget, hasTooltip bool)
	WidgetSetHasWindow          func(w *Widget, hasWindow bool)
	WidgetSetMapped             func(w *Widget, mapped bool)
	WidgetSetName               func(w *Widget, name string)
	WidgetSetNoShowAll          func(w *Widget, noShowAll bool)
	WidgetSetParent             func(w *Widget, parent *Widget)
	WidgetSetParentWindow       func(w *Widget, parentWindow *D.Window)
	WidgetSetRealized           func(w *Widget, realized bool)
	WidgetSetReceivesDefault    func(w *Widget, receivesDefault bool)
	WidgetSetRedrawOnAllocate   func(w *Widget, redrawOnAllocate bool)
	WidgetSetScrollAdjustments  func(w *Widget, hadjustment, vadjustment *Adjustment) bool
	WidgetSetSensitive          func(w *Widget, sensitive bool)
	WidgetSetSizeRequest        func(w *Widget, width int, height int)
	WidgetSetState              func(w *Widget, state StateType)
	WidgetSetStyle              func(w *Widget, style *Style)
	WidgetSetTooltipMarkup      func(w *Widget, markup string)
	WidgetSetTooltipText        func(w *Widget, text string)
	WidgetSetTooltipWindow      func(w *Widget, customWindow *Window)
	WidgetSetUposition          func(w *Widget, x int, y int)
	WidgetSetUsize              func(w *Widget, width int, height int)
	WidgetSetVisible            func(w *Widget, visible bool)
	WidgetSetWindow             func(w *Widget, window *D.Window)
	WidgetShapeCombineMask      func(w *Widget, shapeMask *T.GdkBitmap, offsetX, offsetY int)
	WidgetShow                  func(w *Widget)
	WidgetShowAll               func(w *Widget)
	WidgetShowNow               func(w *Widget)
	WidgetSizeAllocate          func(w *Widget, allocation *Allocation)
	WidgetSizeRequest           func(w *Widget, requisition *Requisition)
	WidgetStyleGet              func(w *Widget, firstPropertyName string, v ...VArg)
	WidgetStyleGetProperty      func(w *Widget, propertyName string, value *O.Value)
	WidgetStyleGetValist        func(w *Widget, firstPropertyName string, varArgs T.VaList)
	WidgetThawChildNotify       func(w *Widget)
	WidgetTriggerTooltipQuery   func(w *Widget)
	WidgetUnmap                 func(w *Widget)
	WidgetUnparent              func(w *Widget)
	WidgetUnrealize             func(w *Widget)
	WidgetUnref                 func(w *Widget)
)

func (w *Widget) AsAboutDialog() *AboutDialog { return (*AboutDialog)(unsafe.Pointer(w)) }
func (w *Widget) AsContainer() *Container     { return (*Container)(unsafe.Pointer(w)) }
func (w *Widget) AsPointer() T.Gpointer       { return (T.Gpointer)(unsafe.Pointer(w)) }
func (w *Widget) AsWindow() *Window           { return (*Window)(unsafe.Pointer(w)) }

func (w *Widget) Activate() bool { return WidgetActivate(w) }
func (w *Widget) AddAccelerator(accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags) {
	WidgetAddAccelerator(w, accelSignal, accelGroup, accelKey, accelMods, accelFlags)
}
func (w *Widget) AddEvents(events int)           { WidgetAddEvents(w, events) }
func (w *Widget) AddMnemonicLabel(label *Widget) { WidgetAddMnemonicLabel(w, label) }
func (w *Widget) CanActivateAccel(signalId uint) bool {
	return WidgetCanActivateAccel(w, signalId)
}
func (w *Widget) ChildFocus(direction DirectionType) bool {
	return WidgetChildFocus(w, direction)
}
func (w *Widget) ChildNotify(childProperty string) { WidgetChildNotify(w, childProperty) }
func (w *Widget) ClassPath(pathLength *uint, path, pathReversed **T.Gchar) {
	WidgetClassPath(w, pathLength, path, pathReversed)
}
func (w *Widget) CreatePangoContext() *P.Context { return WidgetCreatePangoContext(w) }
func (w *Widget) CreatePangoLayout(text string) *P.Layout {
	return WidgetCreatePangoLayout(w, text)
}
func (w *Widget) Destroy()                              { WidgetDestroy(w) }
func (w *Widget) Destroyed(widgetPointer **Widget)      { WidgetDestroyed(w, widgetPointer) }
func (w *Widget) Draw(area *D.Rectangle)                { WidgetDraw(w, area) }
func (w *Widget) EnsureStyle()                          { WidgetEnsureStyle(w) }
func (w *Widget) ErrorBell()                            { WidgetErrorBell(w) }
func (w *Widget) Event(event *D.Event) bool             { return WidgetEvent(w, event) }
func (w *Widget) FreezeChildNotify()                    { WidgetFreezeChildNotify(w) }
func (w *Widget) GetAccessible() *A.Object              { return WidgetGetAccessible(w) }
func (w *Widget) GetAllocation(allocation *Allocation)  { WidgetGetAllocation(w, allocation) }
func (w *Widget) GetAncestor(widgetType O.Type) *Widget { return WidgetGetAncestor(w, widgetType) }
func (w *Widget) GetAppPaintable() bool                 { return WidgetGetAppPaintable(w) }
func (w *Widget) GetCanDefault() bool                   { return WidgetGetCanDefault(w) }
func (w *Widget) GetCanFocus() bool                     { return WidgetGetCanFocus(w) }
func (w *Widget) GetChildRequisition(requisition *Requisition) {
	WidgetGetChildRequisition(w, requisition)
}
func (w *Widget) GetChildVisible() bool                   { return WidgetGetChildVisible(w) }
func (w *Widget) GetColormap() *D.Colormap                { return WidgetGetColormap(w) }
func (w *Widget) GetCompositeName() string                { return WidgetGetCompositeName(w) }
func (w *Widget) GetDirection() TextDirection             { return WidgetGetDirection(w) }
func (w *Widget) GetDisplay() *D.Display                  { return WidgetGetDisplay(w) }
func (w *Widget) GetDoubleBuffered() bool                 { return WidgetGetDoubleBuffered(w) }
func (w *Widget) GetEvents() int                          { return WidgetGetEvents(w) }
func (w *Widget) GetExtensionEvents() D.ExtensionMode     { return WidgetGetExtensionEvents(w) }
func (w *Widget) GetHasTooltip() bool                     { return WidgetGetHasTooltip(w) }
func (w *Widget) GetHasWindow() bool                      { return WidgetGetHasWindow(w) }
func (w *Widget) GetMapped() bool                         { return WidgetGetMapped(w) }
func (w *Widget) GetModifierStyle() *RcStyle              { return WidgetGetModifierStyle(w) }
func (w *Widget) GetName() string                         { return WidgetGetName(w) }
func (w *Widget) GetNoShowAll() bool                      { return WidgetGetNoShowAll(w) }
func (w *Widget) GetPangoContext() *P.Context             { return WidgetGetPangoContext(w) }
func (w *Widget) GetParent() *Widget                      { return WidgetGetParent(w) }
func (w *Widget) GetParentWindow() *D.Window              { return WidgetGetParentWindow(w) }
func (w *Widget) GetPointer(x *int, y *int)               { WidgetGetPointer(w, x, y) }
func (w *Widget) GetRealized() bool                       { return WidgetGetRealized(w) }
func (w *Widget) GetReceivesDefault() bool                { return WidgetGetReceivesDefault(w) }
func (w *Widget) GetRequisition(requisition *Requisition) { WidgetGetRequisition(w, requisition) }
func (w *Widget) GetRootWindow() *D.Window                { return WidgetGetRootWindow(w) }
func (w *Widget) GetScreen() *D.Screen                    { return WidgetGetScreen(w) }
func (w *Widget) GetSensitive() bool                      { return WidgetGetSensitive(w) }
func (w *Widget) GetSettings() *Settings                  { return WidgetGetSettings(w) }
func (w *Widget) GetSizeRequest(width, height *int)       { WidgetGetSizeRequest(w, width, height) }
func (w *Widget) GetSnapshot(clipRect *D.Rectangle) *D.Pixmap {
	return WidgetGetSnapshot(w, clipRect)
}
func (w *Widget) GetState() StateType       { return WidgetGetState(w) }
func (w *Widget) GetStyle() *Style          { return WidgetGetStyle(w) }
func (w *Widget) GetTooltipMarkup() string  { return WidgetGetTooltipMarkup(w) }
func (w *Widget) GetTooltipText() string    { return WidgetGetTooltipText(w) }
func (w *Widget) GetTooltipWindow() *Window { return WidgetGetTooltipWindow(w) }
func (w *Widget) GetToplevel() *Widget      { return WidgetGetToplevel(w) }
func (w *Widget) GetVisible() bool          { return WidgetGetVisible(w) }
func (w *Widget) GetVisual() *D.Visual      { return WidgetGetVisual(w) }
func (w *Widget) GetWindow() *D.Window      { return WidgetGetWindow(w) }
func (w *Widget) GrabDefault()              { WidgetGrabDefault(w) }
func (w *Widget) GrabFocus()                { WidgetGrabFocus(w) }
func (w *Widget) HasDefault() bool          { return WidgetHasDefault(w) }
func (w *Widget) HasFocus() bool            { return WidgetHasFocus(w) }
func (w *Widget) HasGrab() bool             { return WidgetHasGrab(w) }
func (w *Widget) HasRcStyle() bool          { return WidgetHasRcStyle(w) }
func (w *Widget) HasScreen() bool           { return WidgetHasScreen(w) }
func (w *Widget) Hide()                     { WidgetHide(w) }
func (w *Widget) HideAll()                  { WidgetHideAll(w) }
func (w *Widget) HideOnDelete() bool        { return WidgetHideOnDelete(w) }
func (w *Widget) InputShapeCombineMask(shapeMask *T.GdkBitmap, offsetX, offsetY int) {
	WidgetInputShapeCombineMask(w, shapeMask, offsetX, offsetY)
}
func (w *Widget) Intersect(area, intersection *D.Rectangle) bool {
	return WidgetIntersect(w, area, intersection)
}
func (w *Widget) IsAncestor(ancestor *Widget) bool { return WidgetIsAncestor(w, ancestor) }
func (w *Widget) IsComposited() bool               { return WidgetIsComposited(w) }
func (w *Widget) IsDrawable() bool                 { return WidgetIsDrawable(w) }
func (w *Widget) IsFocus() bool                    { return WidgetIsFocus(w) }
func (w *Widget) IsSensitive() bool                { return WidgetIsSensitive(w) }
func (w *Widget) IsToplevel() bool                 { return WidgetIsToplevel(w) }
func (w *Widget) KeynavFailed(direction DirectionType) bool {
	return WidgetKeynavFailed(w, direction)
}
func (w *Widget) ListAccelClosures() *L.List  { return WidgetListAccelClosures(w) }
func (w *Widget) ListMnemonicLabels() *L.List { return WidgetListMnemonicLabels(w) }
func (w *Widget) Map()                        { WidgetMap(w) }
func (w *Widget) MnemonicActivate(groupCycling bool) bool {
	return WidgetMnemonicActivate(w, groupCycling)
}
func (w *Widget) ModifyBase(state StateType, color *D.Color) {
	WidgetModifyBase(w, state, color)
}
func (w *Widget) ModifyBg(state StateType, color *D.Color) { WidgetModifyBg(w, state, color) }
func (w *Widget) ModifyCursor(primary *D.Color, secondary *D.Color) {
	WidgetModifyCursor(w, primary, secondary)
}
func (w *Widget) ModifyFg(state StateType, color *D.Color) { WidgetModifyFg(w, state, color) }
func (w *Widget) ModifyFont(fontDesc *P.FontDescription)   { WidgetModifyFont(w, fontDesc) }
func (w *Widget) ModifyStyle(style *RcStyle)               { WidgetModifyStyle(w, style) }
func (w *Widget) ModifyText(state StateType, color *D.Color) {
	WidgetModifyText(w, state, color)
}
func (w *Widget) Path(pathLength *uint, path, pathReversed **T.Gchar) {
	WidgetPath(w, pathLength, path, pathReversed)
}
func (w *Widget) QueueClear()                            { WidgetQueueClear(w) }
func (w *Widget) QueueClearArea(x, y, width, height int) { WidgetQueueClearArea(w, x, y, width, height) }
func (w *Widget) QueueDraw()                             { WidgetQueueDraw(w) }
func (w *Widget) QueueDrawArea(x, y, width, height int)  { WidgetQueueDrawArea(w, x, y, width, height) }
func (w *Widget) QueueResize()                           { WidgetQueueResize(w) }
func (w *Widget) QueueResizeNoRedraw()                   { WidgetQueueResizeNoRedraw(w) }
func (w *Widget) Realize()                               { WidgetRealize(w) }
func (w *Widget) Ref() *Widget                           { return WidgetRef(w) }
func (w *Widget) RegionIntersect(region *D.Region) *D.Region {
	return WidgetRegionIntersect(w, region)
}
func (w *Widget) RemoveAccelerator(accelGroup *AccelGroup, accelKey uint, accelMods T.GdkModifierType) bool {
	return WidgetRemoveAccelerator(w, accelGroup, accelKey, accelMods)
}
func (w *Widget) RemoveMnemonicLabel(label *Widget) { WidgetRemoveMnemonicLabel(w, label) }
func (w *Widget) RenderIcon(stockId string, size IconSize, detail string) *D.Pixbuf {
	return WidgetRenderIcon(w, stockId, size, detail)
}
func (w *Widget) Reparent(newParent *Widget)              { WidgetReparent(w, newParent) }
func (w *Widget) ResetRcStyles()                          { WidgetResetRcStyles(w) }
func (w *Widget) ResetShapes()                            { WidgetResetShapes(w) }
func (w *Widget) SendExpose(event *D.Event) int           { return WidgetSendExpose(w, event) }
func (w *Widget) SendFocusChange(event *D.Event) bool     { return WidgetSendFocusChange(w, event) }
func (w *Widget) Set(firstPropertyName string, v ...VArg) { WidgetSet(w, firstPropertyName, v) }
func (w *Widget) SetAccelPath(accelPath string, accelGroup *AccelGroup) {
	WidgetSetAccelPath(w, accelPath, accelGroup)
}
func (w *Widget) SetAllocation(allocation *Allocation) { WidgetSetAllocation(w, allocation) }
func (w *Widget) SetAppPaintable(appPaintable bool)    { WidgetSetAppPaintable(w, appPaintable) }
func (w *Widget) SetCanDefault(canDefault bool)        { WidgetSetCanDefault(w, canDefault) }
func (w *Widget) SetCanFocus(canFocus bool)            { WidgetSetCanFocus(w, canFocus) }
func (w *Widget) SetChildVisible(isVisible bool)       { WidgetSetChildVisible(w, isVisible) }
func (w *Widget) SetColormap(colormap *D.Colormap)     { WidgetSetColormap(w, colormap) }
func (w *Widget) SetCompositeName(name string)         { WidgetSetCompositeName(w, name) }
func (w *Widget) SetDirection(dir TextDirection)       { WidgetSetDirection(w, dir) }
func (w *Widget) SetDoubleBuffered(doubleBuffered bool) {
	WidgetSetDoubleBuffered(w, doubleBuffered)
}
func (w *Widget) SetEvents(events int)                    { WidgetSetEvents(w, events) }
func (w *Widget) SetExtensionEvents(mode D.ExtensionMode) { WidgetSetExtensionEvents(w, mode) }
func (w *Widget) SetHasTooltip(hasTooltip bool)           { WidgetSetHasTooltip(w, hasTooltip) }
func (w *Widget) SetHasWindow(hasWindow bool)             { WidgetSetHasWindow(w, hasWindow) }
func (w *Widget) SetMapped(mapped bool)                   { WidgetSetMapped(w, mapped) }
func (w *Widget) SetName(name string)                     { WidgetSetName(w, name) }
func (w *Widget) SetNoShowAll(noShowAll bool)             { WidgetSetNoShowAll(w, noShowAll) }
func (w *Widget) SetParent(parent *Widget)                { WidgetSetParent(w, parent) }
func (w *Widget) SetParentWindow(parentWindow *D.Window)  { WidgetSetParentWindow(w, parentWindow) }
func (w *Widget) SetRealized(realized bool)               { WidgetSetRealized(w, realized) }
func (w *Widget) SetReceivesDefault(receivesDefault bool) {
	WidgetSetReceivesDefault(w, receivesDefault)
}
func (w *Widget) SetRedrawOnAllocate(redrawOnAllocate bool) {
	WidgetSetRedrawOnAllocate(w, redrawOnAllocate)
}
func (w *Widget) SetScrollAdjustments(hadjustment, vadjustment *Adjustment) bool {
	return WidgetSetScrollAdjustments(w, hadjustment, vadjustment)
}
func (w *Widget) SetSensitive(sensitive bool)           { WidgetSetSensitive(w, sensitive) }
func (w *Widget) SetSizeRequest(width, height int)      { WidgetSetSizeRequest(w, width, height) }
func (w *Widget) SetState(state StateType)              { WidgetSetState(w, state) }
func (w *Widget) SetStyle(style *Style)                 { WidgetSetStyle(w, style) }
func (w *Widget) SetTooltipMarkup(markup string)        { WidgetSetTooltipMarkup(w, markup) }
func (w *Widget) SetTooltipText(text string)            { WidgetSetTooltipText(w, text) }
func (w *Widget) SetTooltipWindow(customWindow *Window) { WidgetSetTooltipWindow(w, customWindow) }
func (w *Widget) SetUposition(x, y int)                 { WidgetSetUposition(w, x, y) }
func (w *Widget) SetUsize(width, height int)            { WidgetSetUsize(w, width, height) }
func (w *Widget) SetVisible(visible bool)               { WidgetSetVisible(w, visible) }
func (w *Widget) SetWindow(window *D.Window)            { WidgetSetWindow(w, window) }
func (w *Widget) ShapeCombineMask(shapeMask *T.GdkBitmap, offsetX, offsetY int) {
	WidgetShapeCombineMask(w, shapeMask, offsetX, offsetY)
}
func (w *Widget) Show()                                { WidgetShow(w) }
func (w *Widget) ShowAll()                             { WidgetShowAll(w) }
func (w *Widget) ShowNow()                             { WidgetShowNow(w) }
func (w *Widget) SizeAllocate(allocation *Allocation)  { WidgetSizeAllocate(w, allocation) }
func (w *Widget) SizeRequest(requisition *Requisition) { WidgetSizeRequest(w, requisition) }
func (w *Widget) StyleGet(firstPropertyName string, v ...VArg) {
	WidgetStyleGet(w, firstPropertyName, v)
}
func (w *Widget) StyleGetProperty(propertyName string, value *O.Value) {
	WidgetStyleGetProperty(w, propertyName, value)
}
func (w *Widget) StyleGetValist(firstPropertyName string, varArgs T.VaList) {
	WidgetStyleGetValist(w, firstPropertyName, varArgs)
}
func (w *Widget) ThawChildNotify()     { WidgetThawChildNotify(w) }
func (w *Widget) TriggerTooltipQuery() { WidgetTriggerTooltipQuery(w) }
func (w *Widget) Unmap()               { WidgetUnmap(w) }
func (w *Widget) Unparent()            { WidgetUnparent(w) }
func (w *Widget) Unrealize()           { WidgetUnrealize(w) }
func (w *Widget) Unref()               { WidgetUnref(w) }

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
	WidgetClassFindStyleProperty          func(w *WidgetClass, propertyName string) *T.GParamSpec
	WidgetClassInstallStyleProperty       func(w *WidgetClass, pspec *T.GParamSpec)
	WidgetClassInstallStylePropertyParser func(w *WidgetClass, pspec *T.GParamSpec, parser RcPropertyParser)
	WidgetClassListStyleProperties        func(w *WidgetClass, nProperties *uint) **T.GParamSpec
)

func (w *WidgetClass) FindStyleProperty(propertyName string) *T.GParamSpec {
	return WidgetClassFindStyleProperty(w, propertyName)
}
func (w *WidgetClass) InstallStyleProperty(pspec *T.GParamSpec) {
	WidgetClassInstallStyleProperty(w, pspec)
}
func (w *WidgetClass) InstallStylePropertyParser(pspec *T.GParamSpec, parser RcPropertyParser) {
	WidgetClassInstallStylePropertyParser(w, pspec, parser)
}
func (w *WidgetClass) ListStyleProperties(nProperties *uint) **T.GParamSpec {
	return WidgetClassListStyleProperties(w, nProperties)
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

	WindowGetDefaultIconList         func() *L.List
	WindowGetDefaultIconName         func() string // 2.16
	WindowListToplevels              func() *L.List
	WindowSetAutoStartupNotification func(setting bool)                         // 2.2
	WindowSetDefaultIcon             func(icon *D.Pixbuf)                       // 2.4
	WindowSetDefaultIconFromFile     func(filename string, err **L.Error) bool // 2.2
	WindowSetDefaultIconList         func(list *L.List)
	WindowSetDefaultIconName         func(name string) // 2.6

	WindowActivateDefault       func(w *Window) bool
	WindowActivateFocus         func(w *Window) bool
	WindowActivateKey           func(w *Window, event *D.EventKey) bool
	WindowAddAccelGroup         func(w *Window, accelGroup *AccelGroup)
	WindowAddEmbeddedXid        func(w *Window, xid T.GdkNativeWindow)
	WindowAddMnemonic           func(w *Window, keyval uint, target *Widget)
	WindowBeginMoveDrag         func(w *Window, button, rootX, rootY int, timestamp T.GUint32)
	WindowBeginResizeDrag       func(w *Window, edge D.WindowEdge, button, rootX, rootY int, timestamp T.GUint32)
	WindowDeiconify             func(w *Window)
	WindowFullscreen            func(w *Window)      // 2.2
	WindowGetAcceptFocus        func(w *Window) bool // 2.4
	WindowGetDecorated          func(w *Window) bool
	WindowGetDefaultSize        func(w *Window, width, height *int)
	WindowGetDefaultWidget      func(w *Window) *Widget // 2.14
	WindowGetDeletable          func(w *Window) bool    // 2.10
	WindowGetDestroyWithParent  func(w *Window) bool
	WindowGetFocus              func(w *Window) *Widget
	WindowGetFocusOnMap         func(w *Window) bool // 2.6
	WindowGetFrameDimensions    func(w *Window, left, top, right, bottom *int)
	WindowGetGravity            func(w *Window) D.Gravity
	WindowGetGroup              func(w *Window) *WindowGroup // 2.10
	WindowGetHasFrame           func(w *Window) bool
	WindowGetIcon               func(w *Window) *D.Pixbuf
	WindowGetIconList           func(w *Window) *L.List
	WindowGetIconName           func(w *Window) string // 2.6
	WindowGetMnemonicModifier   func(w *Window) T.GdkModifierType
	WindowGetMnemonicsVisible   func(w *Window) bool
	WindowGetModal              func(w *Window) bool
	WindowGetOpacity            func(w *Window) float64 // 2.12
	WindowGetPosition           func(w *Window, rootX, rootY *int)
	WindowGetResizable          func(w *Window) bool
	WindowGetRole               func(w *Window) string
	WindowGetScreen             func(w *Window) *D.Screen // 2.2
	WindowGetSize               func(w *Window, width, height *int)
	WindowGetSkipPagerHint      func(w *Window) bool // 2.2
	WindowGetSkipTaskbarHint    func(w *Window) bool // 2.2
	WindowGetTitle              func(w *Window) string
	WindowGetTransientFor       func(w *Window) *Window
	WindowGetTypeHint           func(w *Window) D.WindowTypeHint
	WindowGetUrgencyHint        func(w *Window) bool // 2.8
	WindowGetWindowType         func(w *Window) WindowType
	WindowHasGroup              func(w *Window) bool
	WindowHasToplevelFocus      func(w *Window) bool // 2.4
	WindowIconify               func(w *Window)
	WindowIsActive              func(w *Window) bool // 2.4
	WindowMaximize              func(w *Window)
	WindowMnemonicActivate      func(w *Window, keyval uint, modifier T.GdkModifierType) bool
	WindowMove                  func(w *Window, x, y int)
	WindowParseGeometry         func(w *Window, geometry string) bool
	WindowPresent               func(w *Window)
	WindowPresentWithTime       func(w *Window, timestamp T.GUint32) // 2.8
	WindowPropagateKeyEvent     func(w *Window, event *D.EventKey) bool
	WindowRemoveAccelGroup      func(w *Window, accelGroup *AccelGroup)
	WindowRemoveEmbeddedXid     func(w *Window, xid T.GdkNativeWindow)
	WindowRemoveMnemonic        func(w *Window, keyval uint, target *Widget)
	WindowReshowWithInitialSize func(w *Window)
	WindowResize                func(w *Window, width, height int)
	WindowSetAcceptFocus        func(w *Window, setting bool) // 2.4
	WindowSetDecorated          func(w *Window, setting bool)
	WindowSetDefault            func(w *Window, defaultWidget *Widget)
	WindowSetDefaultSize        func(w *Window, width, height int)
	WindowSetDeletable          func(w *Window, setting bool) // 2.10
	WindowSetDestroyWithParent  func(w *Window, setting bool)
	WindowSetFocus              func(w *Window, focus *Widget)
	WindowSetFocusOnMap         func(w *Window, setting bool) // 2.6
	WindowSetFrameDimensions    func(w *Window, left, top, right, bottom int)
	WindowSetGeometryHints      func(w *Window, geometryWidget *Widget, geometry *D.Geometry, geomMask D.WindowHints)
	WindowSetGravity            func(w *Window, gravity D.Gravity)
	WindowSetHasFrame           func(w *Window, setting bool)
	WindowSetIcon               func(w *Window, icon *D.Pixbuf)
	WindowSetIconFromFile       func(w *Window, filename string, err **L.Error) bool // 2.2
	WindowSetIconList           func(w *Window, list *L.List)
	WindowSetIconName           func(w *Window, name string)  // 2.6
	WindowSetKeepAbove          func(w *Window, setting bool) // 2.4
	WindowSetKeepBelow          func(w *Window, setting bool) // 2.4
	WindowSetMnemonicModifier   func(w *Window, modifier T.GdkModifierType)
	WindowSetMnemonicsVisible   func(w *Window, setting bool)
	WindowSetModal              func(w *Window, modal bool)
	WindowSetOpacity            func(w *Window, opacity float64) // 2.12
	// Deprecated windowSetPolicy func(w *Window, allowShrink, allowGrow, autoShrink int)
	WindowSetPosition        func(w *Window, position WindowPosition)
	WindowSetResizable       func(w *Window, resizable bool)
	WindowSetRole            func(w *Window, role string)
	WindowSetScreen          func(w *Window, screen *D.Screen) // 2.2
	WindowSetSkipPagerHint   func(w *Window, setting bool)     // 2.2
	WindowSetSkipTaskbarHint func(w *Window, setting bool)     // 2.2
	WindowSetStartupId       func(w *Window, startupId string) // 2.12
	WindowSetTitle           func(w *Window, title string)
	WindowSetTransientFor    func(w *Window, parent *Window)
	WindowSetTypeHint        func(w *Window, hint D.WindowTypeHint)
	WindowSetUrgencyHint     func(w *Window, setting bool) // 2.8
	WindowSetWmclass         func(w *Window, wmclassName, wmclassClass string)
	WindowStick              func(w *Window)
	WindowUnfullscreen       func(w *Window)
	WindowUnmaximize         func(w *Window)
	WindowUnstick            func(w *Window)
)

func NewWindow(t WindowType) (wt *Widget, w *Window) {
	wt = WindowNew(t)
	w = (*Window)(unsafe.Pointer(wt))
	return
}

func (w *Window) AsContainer() *Container { return (*Container)(unsafe.Pointer(w)) }

func (w *Window) ActivateDefault() bool { return WindowActivateDefault(w) }
func (w *Window) ActivateFocus() bool   { return WindowActivateFocus(w) }
func (w *Window) ActivateKey(event *D.EventKey) bool {
	return WindowActivateKey(w, event)
}
func (w *Window) AddAccelGroup(accelGroup *AccelGroup)    { WindowAddAccelGroup(w, accelGroup) }
func (w *Window) AddEmbeddedXid(xid T.GdkNativeWindow)    { WindowAddEmbeddedXid(w, xid) }
func (w *Window) AddMnemonic(keyval uint, target *Widget) { WindowAddMnemonic(w, keyval, target) }
func (w *Window) BeginMoveDrag(button, rootX, rootY int, timestamp T.GUint32) {
	WindowBeginMoveDrag(w, button, rootX, rootY, timestamp)
}
func (w *Window) BeginResizeDrag(edge D.WindowEdge, button, rootX, rootY int, timestamp T.GUint32) {
	WindowBeginResizeDrag(w, edge, button, rootX, rootY, timestamp)
}
func (w *Window) Deiconify()        { WindowDeiconify(w) }
func (w *Window) Fullscreen()       { WindowFullscreen(w) }
func (w *Window) AcceptFocus() bool { return WindowGetAcceptFocus(w) }
func (w *Window) Decorated() bool   { return WindowGetDecorated(w) }
func (w *Window) DefaultSize(width, height *int) {
	WindowGetDefaultSize(w, width, height)
}
func (w *Window) GetDefaultWidget() *Widget { return WindowGetDefaultWidget(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) Deletable() bool           { return WindowGetDeletable(w) }
func (w *Window) DestroyWithParent() bool   { return WindowGetDestroyWithParent(w) }
func (w *Window) Focus() *Widget            { return WindowGetFocus(w) }
func (w *Window) FocusOnMap() bool          { return WindowGetFocusOnMap(w) }
func (w *Window) FrameDimensions(left, top, right, bottom *int) {
	WindowGetFrameDimensions(w, left, top, right, bottom)
}
func (w *Window) Gravity() D.Gravity                     { return WindowGetGravity(w) }
func (w *Window) GetGroup() *WindowGroup                 { return WindowGetGroup(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) HasFrame() bool                         { return WindowGetHasFrame(w) }
func (w *Window) Icon() *D.Pixbuf                        { return WindowGetIcon(w) }
func (w *Window) IconList() *L.List                      { return WindowGetIconList(w) }
func (w *Window) IconName() string                       { return WindowGetIconName(w) }
func (w *Window) GetMnemonicModifier() T.GdkModifierType { return WindowGetMnemonicModifier(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) MnemonicsVisible() bool                 { return WindowGetMnemonicsVisible(w) }
func (w *Window) Modal() bool                            { return WindowGetModal(w) }
func (w *Window) Opacity() float64                       { return WindowGetOpacity(w) }
func (w *Window) Position(rootX, rootY *int)             { WindowGetPosition(w, rootX, rootY) }
func (w *Window) Resizable() bool                        { return WindowGetResizable(w) }
func (w *Window) Role() string                           { return WindowGetRole(w) }
func (w *Window) GetScreen() *D.Screen                   { return WindowGetScreen(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) Size(width, height *int)                { WindowGetSize(w, width, height) }
func (w *Window) SkipPagerHint() bool                    { return WindowGetSkipPagerHint(w) }
func (w *Window) SkipTaskbarHint() bool                  { return WindowGetSkipTaskbarHint(w) }
func (w *Window) GetTitle() string                       { return WindowGetTitle(w) } //Note(t):removing Get introduces ambiguity
func (w *Window) TransientFor() *Window                  { return WindowGetTransientFor(w) }
func (w *Window) TypeHint() D.WindowTypeHint             { return WindowGetTypeHint(w) }
func (w *Window) UrgencyHint() bool                      { return WindowGetUrgencyHint(w) }
func (w *Window) WindowType() WindowType                 { return WindowGetWindowType(w) }
func (w *Window) HasGroup() bool                         { return WindowHasGroup(w) }
func (w *Window) HasToplevelFocus() bool                 { return WindowHasToplevelFocus(w) }
func (w *Window) Iconify()                               { WindowIconify(w) }
func (w *Window) IsActive() bool                         { return WindowIsActive(w) }
func (w *Window) Maximize()                              { WindowMaximize(w) }
func (w *Window) MnemonicActivate(keyval uint, modifier T.GdkModifierType) bool {
	return WindowMnemonicActivate(w, keyval, modifier)
}
func (w *Window) Move(x int, y int)                   { WindowMove(w, x, y) }
func (w *Window) ParseGeometry(geometry string) bool  { return WindowParseGeometry(w, geometry) }
func (w *Window) Present()                            { WindowPresent(w) }
func (w *Window) PresentWithTime(timestamp T.GUint32) { WindowPresentWithTime(w, timestamp) }
func (w *Window) PropagateKeyEvent(event *D.EventKey) bool {
	return WindowPropagateKeyEvent(w, event)
}
func (w *Window) RemoveAccelGroup(accelGroup *AccelGroup)    { WindowRemoveAccelGroup(w, accelGroup) }
func (w *Window) RemoveEmbeddedXid(xid T.GdkNativeWindow)    { WindowRemoveEmbeddedXid(w, xid) }
func (w *Window) RemoveMnemonic(keyval uint, target *Widget) { WindowRemoveMnemonic(w, keyval, target) }
func (w *Window) ReshowWithInitialSize()                     { WindowReshowWithInitialSize(w) }
func (w *Window) Resize(width, height int)                   { WindowResize(w, width, height) }
func (w *Window) SetAcceptFocus(setting bool)                { WindowSetAcceptFocus(w, setting) }
func (w *Window) SetDecorated(setting bool)                  { WindowSetDecorated(w, setting) }
func (w *Window) SetDefault(defaultWidget *Widget)           { WindowSetDefault(w, defaultWidget) }
func (w *Window) SetDefaultSize(width, height int)           { WindowSetDefaultSize(w, width, height) }
func (w *Window) SetDeletable(setting bool)                  { WindowSetDeletable(w, setting) }
func (w *Window) SetDestroyWithParent(setting bool)          { WindowSetDestroyWithParent(w, setting) }
func (w *Window) SetFocus(focus *Widget)                     { WindowSetFocus(w, focus) }
func (w *Window) SetFocusOnMap(setting bool)                 { WindowSetFocusOnMap(w, setting) }
func (w *Window) SetFrameDimensions(left, top, right, bottom int) {
	WindowSetFrameDimensions(w, left, top, right, bottom)
}
func (w *Window) SetGeometryHints(geometryWidget *Widget, geometry *D.Geometry, geomMask D.WindowHints) {
	WindowSetGeometryHints(w, geometryWidget, geometry, geomMask)
}
func (w *Window) SetGravity(gravity D.Gravity) { WindowSetGravity(w, gravity) }
func (w *Window) SetHasFrame(setting bool)     { WindowSetHasFrame(w, setting) }
func (w *Window) SetIcon(icon *D.Pixbuf)       { WindowSetIcon(w, icon) }
func (w *Window) SetIconFromFile(filename string, err **L.Error) bool {
	return WindowSetIconFromFile(w, filename, err)
}
func (w *Window) SetIconList(list *L.List)  { WindowSetIconList(w, list) }
func (w *Window) SetIconName(name string)   { WindowSetIconName(w, name) }
func (w *Window) SetKeepAbove(setting bool) { WindowSetKeepAbove(w, setting) }
func (w *Window) SetKeepBelow(setting bool) { WindowSetKeepBelow(w, setting) }
func (w *Window) SetMnemonicModifier(modifier T.GdkModifierType) {
	WindowSetMnemonicModifier(w, modifier)
}
func (w *Window) SetMnemonicsVisible(setting bool) { WindowSetMnemonicsVisible(w, setting) }
func (w *Window) SetModal(modal bool)              { WindowSetModal(w, modal) }
func (w *Window) SetOpacity(opacity float64)       { WindowSetOpacity(w, opacity) }

// Deprecated
// func (w *Window) SetPolicy(allowShrink, allowGrow, autoShrink int) {
// 	windowSetPolicy(w, allowShrink, allowGrow, autoShrink)
// }
func (w *Window) SetPosition(position WindowPosition) { WindowSetPosition(w, position) }
func (w *Window) SetResizable(resizable bool)         { WindowSetResizable(w, resizable) }
func (w *Window) SetRole(role string)                 { WindowSetRole(w, role) }
func (w *Window) SetScreen(screen *D.Screen)          { WindowSetScreen(w, screen) }
func (w *Window) SetSkipPagerHint(setting bool)       { WindowSetSkipPagerHint(w, setting) }
func (w *Window) SetSkipTaskbarHint(setting bool)     { WindowSetSkipTaskbarHint(w, setting) }
func (w *Window) SetStartupId(startupId string)       { WindowSetStartupId(w, startupId) }
func (w *Window) SetTitle(title string)               { WindowSetTitle(w, title) }
func (w *Window) SetTransientFor(parent *Window)      { WindowSetTransientFor(w, parent) }
func (w *Window) SetTypeHint(hint D.WindowTypeHint)   { WindowSetTypeHint(w, hint) }
func (w *Window) SetUrgencyHint(setting bool)         { WindowSetUrgencyHint(w, setting) }
func (w *Window) SetWmclass(wmclassName, wmclassClass string) {
	WindowSetWmclass(w, wmclassName, wmclassClass)
}
func (w *Window) Stick()        { WindowStick(w) }
func (w *Window) Unfullscreen() { WindowUnfullscreen(w) }
func (w *Window) Unmaximize()   { WindowUnmaximize(w) }
func (w *Window) Unstick()      { WindowUnstick(w) }

type WindowGeometryInfo struct{}

type WindowGroup struct {
	Parent O.Object
	Grabs  *L.SList
}

var (
	WindowGroupGetType func() O.Type
	WindowGroupNew     func() *WindowGroup

	WindowGroupAddWindow      func(w *WindowGroup, window *Window)
	WindowGroupGetCurrentGrab func(w *WindowGroup) *Widget
	WindowGroupListWindows    func(w *WindowGroup) *L.List
	WindowGroupRemoveWindow   func(w *WindowGroup, window *Window)
)

func (w *WindowGroup) AddWindow(window *Window)    { WindowGroupAddWindow(w, window) }
func (w *WindowGroup) GetCurrentGrab() *Widget     { return WindowGroupGetCurrentGrab(w) }
func (w *WindowGroup) ListWindows() *L.List        { return WindowGroupListWindows(w) }
func (w *WindowGroup) RemoveWindow(window *Window) { WindowGroupRemoveWindow(w, window) }

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
