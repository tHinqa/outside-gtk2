// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Scale struct {
	Range  Range
	Digits int
	Bits   uint
	// DrawValue : 1
	// ValuePos : 2
}

var (
	ScaleGetType   func() T.GType
	ScaleSetDigits func(scale *Scale, digits int)

	scaleAddMark          func(s *Scale, value float64, position PositionType, markup string)
	scaleClearMarks       func(s *Scale)
	scaleGetDigits        func(s *Scale) int
	scaleGetDrawValue     func(s *Scale) T.Gboolean
	scaleGetLayout        func(s *Scale) *T.PangoLayout
	scaleGetLayoutOffsets func(s *Scale, x, y *int)
	scaleGetValuePos      func(s *Scale) PositionType
	scaleSetDrawValue     func(s *Scale, drawValue T.Gboolean)
	scaleSetValuePos      func(s *Scale, pos PositionType)
)

func (s *Scale) AddMark(value float64, position PositionType, markup string) {
	scaleAddMark(s, value, position, markup)
}
func (s *Scale) ClearMarks()                       { scaleClearMarks(s) }
func (s *Scale) GetDigits() int                    { return scaleGetDigits(s) }
func (s *Scale) GetDrawValue() T.Gboolean          { return scaleGetDrawValue(s) }
func (s *Scale) GetLayout() *T.PangoLayout         { return scaleGetLayout(s) }
func (s *Scale) GetLayoutOffsets(x, y *int)        { scaleGetLayoutOffsets(s, x, y) }
func (s *Scale) GetValuePos() PositionType         { return scaleGetValuePos(s) }
func (s *Scale) SetDrawValue(drawValue T.Gboolean) { scaleSetDrawValue(s, drawValue) }
func (s *Scale) SetValuePos(pos PositionType)      { scaleSetValuePos(s, pos) }

type ScaleButton struct {
	Parent      Button
	PlusButton  *Widget
	MinusButton *Widget
	_           *struct{}
}

var (
	ScaleButtonGetType func() T.GType
	ScaleButtonNew     func(size IconSize, min, max, step float64, icons **T.Gchar) *Widget

	scaleButtonGetAdjustment  func(s *ScaleButton) *Adjustment
	scaleButtonGetMinusButton func(s *ScaleButton) *Widget
	scaleButtonGetOrientation func(s *ScaleButton) Orientation
	scaleButtonGetPlusButton  func(s *ScaleButton) *Widget
	scaleButtonGetPopup       func(s *ScaleButton) *Widget
	scaleButtonGetValue       func(s *ScaleButton) float64
	scaleButtonSetAdjustment  func(s *ScaleButton, adjustment *Adjustment)
	scaleButtonSetIcons       func(s *ScaleButton, icons **T.Gchar)
	scaleButtonSetOrientation func(s *ScaleButton, orientation Orientation)
	scaleButtonSetValue       func(s *ScaleButton, value float64)
)

func (s *ScaleButton) GetAdjustment() *Adjustment           { return scaleButtonGetAdjustment(s) }
func (s *ScaleButton) GetMinusButton() *Widget              { return scaleButtonGetMinusButton(s) }
func (s *ScaleButton) GetOrientation() Orientation          { return scaleButtonGetOrientation(s) }
func (s *ScaleButton) GetPlusButton() *Widget               { return scaleButtonGetPlusButton(s) }
func (s *ScaleButton) GetPopup() *Widget                    { return scaleButtonGetPopup(s) }
func (s *ScaleButton) GetValue() float64                    { return scaleButtonGetValue(s) }
func (s *ScaleButton) SetAdjustment(adjustment *Adjustment) { scaleButtonSetAdjustment(s, adjustment) }
func (s *ScaleButton) SetIcons(icons **T.Gchar)             { scaleButtonSetIcons(s, icons) }
func (s *ScaleButton) SetOrientation(orientation Orientation) {
	scaleButtonSetOrientation(s, orientation)
}
func (s *ScaleButton) SetValue(value float64) { scaleButtonSetValue(s, value) }

var ScrollbarGetType func() T.GType //TODO(t):Use?

type ScrolledWindow struct {
	Container  Bin
	Hscrollbar *Widget
	Vscrollbar *Widget
	Bits       uint
	// HscrollbarPolicy : 2
	// VscrollbarPolicy : 2
	// HscrollbarVisible : 1
	// VscrollbarVisible : 1
	// WindowPlacement : 2
	// FocusOut : 1
	ShadowType uint16
}

var (
	ScrolledWindowGetType func() T.GType
	ScrolledWindowNew     func(hadjustment, vadjustment *Adjustment) *Widget

	scrolledWindowAddWithViewport func(s *ScrolledWindow, child *Widget)
	scrolledWindowGetHadjustment  func(s *ScrolledWindow) *Adjustment
	scrolledWindowGetHscrollbar   func(s *ScrolledWindow) *Widget
	scrolledWindowGetPlacement    func(s *ScrolledWindow) CornerType
	scrolledWindowGetPolicy       func(s *ScrolledWindow, hscrollbarPolicy, vscrollbarPolicy *PolicyType)
	scrolledWindowGetShadowType   func(s *ScrolledWindow) ShadowType
	scrolledWindowGetVadjustment  func(s *ScrolledWindow) *Adjustment
	scrolledWindowGetVscrollbar   func(s *ScrolledWindow) *Widget
	scrolledWindowSetHadjustment  func(s *ScrolledWindow, hadjustment *Adjustment)
	scrolledWindowSetPlacement    func(s *ScrolledWindow, windowPlacement CornerType)
	scrolledWindowSetPolicy       func(s *ScrolledWindow, hscrollbarPolicy, vscrollbarPolicy PolicyType)
	scrolledWindowSetShadowType   func(s *ScrolledWindow, t ShadowType)
	scrolledWindowSetVadjustment  func(s *ScrolledWindow, vadjustment *Adjustment)
	scrolledWindowUnsetPlacement  func(s *ScrolledWindow)
)

func (s *ScrolledWindow) AddWithViewport(child *Widget) {
	scrolledWindowAddWithViewport(s, child)
}
func (s *ScrolledWindow) GetHadjustment() *Adjustment { return scrolledWindowGetHadjustment(s) }
func (s *ScrolledWindow) GetHscrollbar() *Widget      { return scrolledWindowGetHscrollbar(s) }
func (s *ScrolledWindow) GetPlacement() CornerType    { return scrolledWindowGetPlacement(s) }
func (s *ScrolledWindow) GetPolicy(hscrollbarPolicy, vscrollbarPolicy *PolicyType) {
	scrolledWindowGetPolicy(s, hscrollbarPolicy, vscrollbarPolicy)
}
func (s *ScrolledWindow) GetShadowType() ShadowType   { return scrolledWindowGetShadowType(s) }
func (s *ScrolledWindow) GetVadjustment() *Adjustment { return scrolledWindowGetVadjustment(s) }
func (s *ScrolledWindow) GetVscrollbar() *Widget      { return scrolledWindowGetVscrollbar(s) }
func (s *ScrolledWindow) SetHadjustment(hadjustment *Adjustment) {
	scrolledWindowSetHadjustment(s, hadjustment)
}
func (s *ScrolledWindow) SetPlacement(windowPlacement CornerType) {
	scrolledWindowSetPlacement(s, windowPlacement)
}
func (s *ScrolledWindow) SetPolicy(hscrollbarPolicy, vscrollbarPolicy PolicyType) {
	scrolledWindowSetPolicy(s, hscrollbarPolicy, vscrollbarPolicy)
}
func (s *ScrolledWindow) SetShadowType(t ShadowType) { scrolledWindowSetShadowType(s, t) }
func (s *ScrolledWindow) SetVadjustment(vadjustment *Adjustment) {
	scrolledWindowSetVadjustment(s, vadjustment)
}
func (s *ScrolledWindow) UnsetPlacement() { scrolledWindowUnsetPlacement(s) }

var ScrollStepGetType func() T.GType //TODO(t):Use?

type ScrollType Enum

const (
	SCROLL_NONE ScrollType = iota
	SCROLL_JUMP
	SCROLL_STEP_BACKWARD
	SCROLL_STEP_FORWARD
	SCROLL_PAGE_BACKWARD
	SCROLL_PAGE_FORWARD
	SCROLL_STEP_UP
	SCROLL_STEP_DOWN
	SCROLL_PAGE_UP
	SCROLL_PAGE_DOWN
	SCROLL_STEP_LEFT
	SCROLL_STEP_RIGHT
	SCROLL_PAGE_LEFT
	SCROLL_PAGE_RIGHT
	SCROLL_START
	SCROLL_END
)

var ScrollTypeGetType func() T.GType

type SelectionData struct {
	Selection D.Atom
	Target    D.Atom
	Type      D.Atom
	Format    int
	Data      *T.Guchar
	Length    int
	Display   *D.Display
}

var (
	SelectionDataGetType func() T.GType

	SelectionAddTarget          func(widget *Widget, selection, target D.Atom, info uint)
	SelectionAddTargets         func(widget *Widget, selection D.Atom, targets *TargetEntry, ntargets uint)
	SelectionClear              func(widget *Widget, event *D.EventSelection) T.Gboolean
	SelectionClearTargets       func(widget *Widget, selection D.Atom)
	SelectionConvert            func(widget *Widget, selection, target D.Atom, time T.GUint32) T.Gboolean
	selectionDataCopy           func(s *SelectionData) *SelectionData
	SelectionOwnerSet           func(widget *Widget, selection D.Atom, time T.GUint32) T.Gboolean
	SelectionOwnerSetForDisplay func(display *D.Display, widget *Widget, selection D.Atom, time T.GUint32) T.Gboolean
	SelectionRemoveAll          func(widget *Widget)

	selectionDataFree                   func(s *SelectionData)
	selectionDataGetData                func(s *SelectionData) *T.Guchar
	selectionDataGetDataType            func(s *SelectionData) D.Atom
	selectionDataGetDisplay             func(s *SelectionData) *D.Display
	selectionDataGetFormat              func(s *SelectionData) int
	selectionDataGetLength              func(s *SelectionData) int
	selectionDataGetPixbuf              func(s *SelectionData) *D.Pixbuf
	selectionDataGetSelection           func(s *SelectionData) D.Atom
	selectionDataGetTarget              func(s *SelectionData) D.Atom
	selectionDataGetTargets             func(s *SelectionData, targets **D.Atom, nAtoms *int) T.Gboolean
	selectionDataGetText                func(s *SelectionData) *T.Guchar
	selectionDataGetUris                func(s *SelectionData) **T.Gchar
	selectionDataSet                    func(s *SelectionData, typ D.Atom, format int, data *T.Guchar, length int)
	selectionDataSetPixbuf              func(s *SelectionData, pixbuf *D.Pixbuf) T.Gboolean
	selectionDataSetText                func(s *SelectionData, str string, len int) T.Gboolean
	selectionDataSetUris                func(s *SelectionData, uris **T.Gchar) T.Gboolean
	selectionDataTargetsIncludeImage    func(s *SelectionData, writable T.Gboolean) T.Gboolean
	selectionDataTargetsIncludeRichText func(s *SelectionData, buffer *TextBuffer) T.Gboolean
	selectionDataTargetsIncludeText     func(s *SelectionData) T.Gboolean
	selectionDataTargetsIncludeUri      func(s *SelectionData) T.Gboolean
)

func (s *SelectionData) Copy() *SelectionData   { return selectionDataCopy(s) }
func (s *SelectionData) Free()                  { selectionDataFree(s) }
func (s *SelectionData) GetData() *T.Guchar     { return selectionDataGetData(s) }
func (s *SelectionData) GetDataType() D.Atom    { return selectionDataGetDataType(s) }
func (s *SelectionData) GetDisplay() *D.Display { return selectionDataGetDisplay(s) }
func (s *SelectionData) GetFormat() int         { return selectionDataGetFormat(s) }
func (s *SelectionData) GetLength() int         { return selectionDataGetLength(s) }
func (s *SelectionData) GetPixbuf() *D.Pixbuf   { return selectionDataGetPixbuf(s) }
func (s *SelectionData) GetSelection() D.Atom   { return selectionDataGetSelection(s) }
func (s *SelectionData) GetTarget() D.Atom      { return selectionDataGetTarget(s) }
func (s *SelectionData) GetTargets(targets **D.Atom, nAtoms *int) T.Gboolean {
	return selectionDataGetTargets(s, targets, nAtoms)
}
func (s *SelectionData) GetText() *T.Guchar { return selectionDataGetText(s) }
func (s *SelectionData) GetUris() **T.Gchar { return selectionDataGetUris(s) }
func (s *SelectionData) Set(typ D.Atom, format int, data *T.Guchar, length int) {
	selectionDataSet(s, typ, format, data, length)
}
func (s *SelectionData) SetPixbuf(pixbuf *D.Pixbuf) T.Gboolean {
	return selectionDataSetPixbuf(s, pixbuf)
}
func (s *SelectionData) SetText(str string, leng int) T.Gboolean {
	return selectionDataSetText(s, str, leng)
}
func (s *SelectionData) SetUris(uris **T.Gchar) T.Gboolean { return selectionDataSetUris(s, uris) }
func (s *SelectionData) TargetsIncludeImage(writable T.Gboolean) T.Gboolean {
	return selectionDataTargetsIncludeImage(s, writable)
}
func (s *SelectionData) TargetsIncludeRichText(buffer *TextBuffer) T.Gboolean {
	return selectionDataTargetsIncludeRichText(s, buffer)
}
func (s *SelectionData) TargetsIncludeText() T.Gboolean { return selectionDataTargetsIncludeText(s) }
func (s *SelectionData) TargetsIncludeUri() T.Gboolean  { return selectionDataTargetsIncludeUri(s) }

type SelectionMode Enum

const (
	SELECTION_NONE SelectionMode = iota
	SELECTION_SINGLE
	SELECTION_BROWSE
	SELECTION_MULTIPLE
	SELECTION_EXTENDED = SELECTION_MULTIPLE
)

var SelectionModeGetType func() T.GType

type SensitivityType Enum

const (
	SENSITIVITY_AUTO SensitivityType = iota
	SENSITIVITY_ON
	SENSITIVITY_OFF
)

var SensitivityTypeGetType func() T.GType

var (
	SeparatorGetType func() T.GType

	SeparatorMenuItemGetType func() T.GType
	SeparatorMenuItemNew     func() *Widget
)

type SeparatorToolItem struct {
	Parent ToolItem
	_      *struct{}
}

var (
	SeparatorToolItemGetType func() T.GType
	SeparatorToolItemNew     func() *ToolItem

	separatorToolItemGetDraw func(s *SeparatorToolItem) T.Gboolean
	separatorToolItemSetDraw func(s *SeparatorToolItem, draw T.Gboolean)
)

func (s *SeparatorToolItem) GetDraw() T.Gboolean     { return separatorToolItemGetDraw(s) }
func (s *SeparatorToolItem) SetDraw(draw T.Gboolean) { separatorToolItemSetDraw(s, draw) }

type Settings struct {
	Parent          T.GObject
	QueuedSettings  *T.GData
	Property_values *SettingsPropertyValue
	RcContext       *RcContext
	Screen          *D.Screen
}

var (
	SettingsGetType    func() T.GType
	SettingsGetDefault func() *Settings

	SettingsGetForScreen          func(screen *D.Screen) *Settings
	SettingsInstallProperty       func(pspec *T.GParamSpec)
	SettingsInstallPropertyParser func(pspec *T.GParamSpec, parser RcPropertyParser)

	settingsSetDoubleProperty func(s *Settings, name string, vDouble float64, origin string)
	settingsSetLongProperty   func(s *Settings, name string, vLong T.Glong, origin string)
	settingsSetPropertyValue  func(s *Settings, name string, svalue *SettingsValue)
	settingsSetStringProperty func(s *Settings, name, vString, origin string)
)

func (s *Settings) SetDoubleProperty(name string, vDouble float64, origin string) {
	settingsSetDoubleProperty(s, name, vDouble, origin)
}
func (s *Settings) SetLongProperty(name string, vLong T.Glong, origin string) {
	settingsSetLongProperty(s, name, vLong, origin)
}
func (s *Settings) SetPropertyValue(name string, svalue *SettingsValue) {
	settingsSetPropertyValue(s, name, svalue)
}
func (s *Settings) SetStringProperty(name, vString, origin string) {
	settingsSetStringProperty(s, name, vString, origin)
}

type SettingsPropertyValue struct{}

type SettingsValue struct {
	Origin *T.Gchar
	Value  T.GValue
}

type ShadowType Enum

const (
	SHADOW_NONE ShadowType = iota
	SHADOW_IN
	SHADOW_OUT
	SHADOW_ETCHED_IN
	SHADOW_ETCHED_OUT
)

var ShadowTypeGetType func() T.GType

type SideType Enum

const (
	SIDE_TOP SideType = iota
	SIDE_BOTTOM
	SIDE_LEFT
	SIDE_RIGHT
)

var SideTypeGetType func() T.GType

type SignalRunType Enum

const (
	RUN_FIRST      SignalRunType = SignalRunType(T.G_SIGNAL_RUN_FIRST)
	RUN_LAST                     = SignalRunType(T.G_SIGNAL_RUN_LAST)
	RUN_BOTH                     = SignalRunType(RUN_FIRST | RUN_LAST)
	RUN_NO_RECURSE               = SignalRunType(T.G_SIGNAL_NO_RECURSE)
	RUN_ACTION                   = SignalRunType(T.G_SIGNAL_ACTION)
	RUN_NO_HOOKS                 = SignalRunType(T.G_SIGNAL_NO_HOOKS)
)

var SignalRunTypeGetType func() T.GType

type SizeGroup struct {
	Parent  T.GObject
	Widgets *T.GSList
	Mode    uint8
	Bits    uint
	// HaveWidth : 1
	// HaveHeight : 1
	// IgnoreHidden : 1
	Requisition Requisition
}

type SizeGroupMode Enum

const (
	SIZE_GROUP_NONE SizeGroupMode = iota
	SIZE_GROUP_HORIZONTAL
	SIZE_GROUP_VERTICAL
	SIZE_GROUP_BOTH
)

var (
	SizeGroupGetType func() T.GType
	SizeGroupNew     func(mode SizeGroupMode) *SizeGroup

	SizeGroupModeGetType func() T.GType

	sizeGroupAddWidget       func(s *SizeGroup, widget *Widget)
	sizeGroupGetIgnoreHidden func(s *SizeGroup) T.Gboolean
	sizeGroupGetMode         func(s *SizeGroup) SizeGroupMode
	sizeGroupGetWidgets      func(s *SizeGroup) *T.GSList
	sizeGroupRemoveWidget    func(s *SizeGroup, widget *Widget)
	sizeGroupSetIgnoreHidden func(s *SizeGroup, ignoreHidden T.Gboolean)
	sizeGroupSetMode         func(s *SizeGroup, mode SizeGroupMode)
)

func (s *SizeGroup) AddWidget(widget *Widget)    { sizeGroupAddWidget(s, widget) }
func (s *SizeGroup) GetIgnoreHidden() T.Gboolean { return sizeGroupGetIgnoreHidden(s) }
func (s *SizeGroup) GetMode() SizeGroupMode      { return sizeGroupGetMode(s) }
func (s *SizeGroup) GetWidgets() *T.GSList       { return sizeGroupGetWidgets(s) }
func (s *SizeGroup) RemoveWidget(widget *Widget) { sizeGroupRemoveWidget(s, widget) }
func (s *SizeGroup) SetIgnoreHidden(ignoreHidden T.Gboolean) {
	sizeGroupSetIgnoreHidden(s, ignoreHidden)
}
func (s *SizeGroup) SetMode(mode SizeGroupMode) { sizeGroupSetMode(s, mode) }

type Socket struct {
	Container     Container
	RequestWidth  uint16
	RequestHeight uint16
	CurrentWidth  uint16
	CurrentHeight uint16
	PlugWindow    *D.Window
	PlugWidget    *Widget
	XembedVersion int16
	Bits          uint
	// SameApp : 1
	// FocusIn : 1
	// HaveSize : 1
	// NeedMap : 1
	// IsMapped : 1
	// Active : 1
	AccelGroup *AccelGroup
	Toplevel   *Widget
}

var (
	SocketGetType func() T.GType
	SocketNew     func() *Widget

	socketAddId         func(s *Socket, windowId T.GdkNativeWindow)
	socketGetId         func(s *Socket) T.GdkNativeWindow
	socketGetPlugWindow func(s *Socket) *D.Window
	socketSteal         func(s *Socket, wid T.GdkNativeWindow)
)

func (s *Socket) AddId(windowId T.GdkNativeWindow) { socketAddId(s, windowId) }
func (s *Socket) GetId() T.GdkNativeWindow         { return socketGetId(s) }
func (s *Socket) GetPlugWindow() *D.Window         { return socketGetPlugWindow(s) }
func (s *Socket) Steal(wid T.GdkNativeWindow)      { socketSteal(s, wid) }

type SortType Enum

const (
	SORT_ASCENDING SortType = iota
	SORT_DESCENDING
)

var SortTypeGetType func() T.GType

type SpinButton struct {
	Entry        Entry
	Adjustment   *Adjustment
	Panel        *D.Window
	Timer        T.GUint32
	ClimbRate    float64
	TimerStep    float64
	UpdatePolicy SpinButtonUpdatePolicy
	Bits         uint
	// InChild : 2
	// ClickChild : 2
	// Button : 2
	// NeedTimer : 1
	// TimerCalls : 3
	// Digits : 10
	// Numeric : 1
	// Wrap : 1
	// SnapToTicks : 1
}

type SpinButtonUpdatePolicy Enum

const (
	UPDATE_ALWAYS SpinButtonUpdatePolicy = iota
	UPDATE_IF_VALID
)

var (
	SpinButtonGetType      func() T.GType
	SpinButtonNew          func(adjustment *Adjustment, climbRate float64, digits uint) *Widget
	SpinButtonNewWithRange func(min, max, step float64) *Widget

	SpinButtonUpdatePolicyGetType func() T.GType

	spinButtonConfigure       func(s *SpinButton, adjustment *Adjustment, climbRate float64, digits uint)
	spinButtonGetAdjustment   func(s *SpinButton) *Adjustment
	spinButtonGetDigits       func(s *SpinButton) uint
	spinButtonGetIncrements   func(s *SpinButton, step, page *float64)
	spinButtonGetNumeric      func(s *SpinButton) T.Gboolean
	spinButtonGetRange        func(s *SpinButton, min, max *float64)
	spinButtonGetSnapToTicks  func(s *SpinButton) T.Gboolean
	spinButtonGetUpdatePolicy func(s *SpinButton) SpinButtonUpdatePolicy
	spinButtonGetValue        func(s *SpinButton) float64
	spinButtonGetValueAsInt   func(s *SpinButton) int
	spinButtonGetWrap         func(s *SpinButton) T.Gboolean
	spinButtonSetAdjustment   func(s *SpinButton, adjustment *Adjustment)
	spinButtonSetDigits       func(s *SpinButton, digits uint)
	spinButtonSetIncrements   func(s *SpinButton, step, page float64)
	spinButtonSetNumeric      func(s *SpinButton, numeric T.Gboolean)
	spinButtonSetRange        func(s *SpinButton, min, max float64)
	spinButtonSetSnapToTicks  func(s *SpinButton, snapToTicks T.Gboolean)
	spinButtonSetUpdatePolicy func(s *SpinButton, policy SpinButtonUpdatePolicy)
	spinButtonSetValue        func(s *SpinButton, value float64)
	spinButtonSetWrap         func(s *SpinButton, wrap T.Gboolean)
	spinButtonSpin            func(s *SpinButton, direction SpinType, increment float64)
	spinButtonUpdate          func(s *SpinButton)
)

func (s *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint) {
	spinButtonConfigure(s, adjustment, climbRate, digits)
}
func (s *SpinButton) GetAdjustment() *Adjustment              { return spinButtonGetAdjustment(s) }
func (s *SpinButton) GetDigits() uint                         { return spinButtonGetDigits(s) }
func (s *SpinButton) GetIncrements(step, page *float64)       { spinButtonGetIncrements(s, step, page) }
func (s *SpinButton) GetNumeric() T.Gboolean                  { return spinButtonGetNumeric(s) }
func (s *SpinButton) GetRange(min, max *float64)              { spinButtonGetRange(s, min, max) }
func (s *SpinButton) GetSnapToTicks() T.Gboolean              { return spinButtonGetSnapToTicks(s) }
func (s *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy { return spinButtonGetUpdatePolicy(s) }
func (s *SpinButton) GetValue() float64                       { return spinButtonGetValue(s) }
func (s *SpinButton) GetValueAsInt() int                      { return spinButtonGetValueAsInt(s) }
func (s *SpinButton) GetWrap() T.Gboolean                     { return spinButtonGetWrap(s) }
func (s *SpinButton) SetAdjustment(adjustment *Adjustment)    { spinButtonSetAdjustment(s, adjustment) }
func (s *SpinButton) SetDigits(digits uint)                   { spinButtonSetDigits(s, digits) }
func (s *SpinButton) SetIncrements(step, page float64)        { spinButtonSetIncrements(s, step, page) }
func (s *SpinButton) SetNumeric(numeric T.Gboolean)           { spinButtonSetNumeric(s, numeric) }
func (s *SpinButton) SetRange(min, max float64)               { spinButtonSetRange(s, min, max) }
func (s *SpinButton) SetSnapToTicks(snapToTicks T.Gboolean)   { spinButtonSetSnapToTicks(s, snapToTicks) }
func (s *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	spinButtonSetUpdatePolicy(s, policy)
}
func (s *SpinButton) SetValue(value float64)  { spinButtonSetValue(s, value) }
func (s *SpinButton) SetWrap(wrap T.Gboolean) { spinButtonSetWrap(s, wrap) }
func (s *SpinButton) Spin(direction SpinType, increment float64) {
	spinButtonSpin(s, direction, increment)
}
func (s *SpinButton) Update() { spinButtonUpdate(s) }

type Spinner struct {
	Parent DrawingArea
	_      *struct{}
}

var (
	SpinnerGetType func() T.GType
	SpinnerNew     func() *Widget

	spinnerStart func(s *Spinner)
	spinnerStop  func(s *Spinner)
)

func (s *Spinner) Start() { spinnerStart(s) }
func (s *Spinner) Stop()  { spinnerStop(s) }

type SpinType Enum

const (
	SPIN_STEP_FORWARD SpinType = iota
	SPIN_STEP_BACKWARD
	SPIN_PAGE_FORWARD
	SPIN_PAGE_BACKWARD
	SPIN_HOME
	SPIN_END
	SPIN_USER_DEFINED
)

var SpinTypeGetType func() T.GType

type StateType Enum

const (
	STATE_NORMAL StateType = iota
	STATE_ACTIVE
	STATE_PRELIGHT
	STATE_SELECTED
	STATE_INSENSITIVE
)

var StateTypeGetType func() T.GType

type Statusbar struct {
	ParentWidget  HBox
	Frame         *Widget
	Label         *Widget
	Messages      *T.GSList
	Keys          *T.GSList
	SeqContextId  uint
	SeqMessageId  uint
	Grip_window   *D.Window
	HasResizeGrip uint // : 1
}

var (
	StatusbarGetType func() T.GType
	StatusbarNew     func() *Widget

	statusbarGetContextId     func(s *Statusbar, contextDescription string) uint
	statusbarGetHasResizeGrip func(s *Statusbar) T.Gboolean
	statusbarGetMessageArea   func(s *Statusbar) *Widget
	statusbarPop              func(s *Statusbar, contextId uint)
	statusbarPush             func(s *Statusbar, contextId uint, text string) uint
	statusbarRemove           func(s *Statusbar, contextId uint, messageId uint)
	statusbarRemoveAll        func(s *Statusbar, contextId uint)
	statusbarSetHasResizeGrip func(s *Statusbar, setting T.Gboolean)
)

func (s *Statusbar) GetContextId(contextDescription string) uint {
	return statusbarGetContextId(s, contextDescription)
}
func (s *Statusbar) GetHasResizeGrip() T.Gboolean          { return statusbarGetHasResizeGrip(s) }
func (s *Statusbar) GetMessageArea() *Widget               { return statusbarGetMessageArea(s) }
func (s *Statusbar) Pop(contextId uint)                    { statusbarPop(s, contextId) }
func (s *Statusbar) Push(contextId uint, text string) uint { return statusbarPush(s, contextId, text) }
func (s *Statusbar) Remove(contextId uint, messageId uint) { statusbarRemove(s, contextId, messageId) }
func (s *Statusbar) RemoveAll(contextId uint)              { statusbarRemoveAll(s, contextId) }
func (s *Statusbar) SetHasResizeGrip(setting T.Gboolean)   { statusbarSetHasResizeGrip(s, setting) }

type StatusIcon struct {
	Parent T.GObject
	_      *struct{}
}

var (
	StatusIconGetType         func() T.GType
	StatusIconNew             func() *StatusIcon
	StatusIconNewFromFile     func(filename string) *StatusIcon
	StatusIconNewFromGicon    func(icon *T.GIcon) *StatusIcon
	StatusIconNewFromIconName func(iconName string) *StatusIcon
	StatusIconNewFromPixbuf   func(pixbuf *D.Pixbuf) *StatusIcon
	StatusIconNewFromStock    func(stockId string) *StatusIcon

	StatusIconPositionMenu func(menu *Menu, x, y *int, pushIn *T.Gboolean, userData T.Gpointer)

	statusIconGetBlinking      func(s *StatusIcon) T.Gboolean
	statusIconGetGeometry      func(s *StatusIcon, screen **D.Screen, area *D.Rectangle, orientation *Orientation) T.Gboolean
	statusIconGetGicon         func(s *StatusIcon) *T.GIcon
	statusIconGetHasTooltip    func(s *StatusIcon) T.Gboolean
	statusIconGetIconName      func(s *StatusIcon) string
	statusIconGetPixbuf        func(s *StatusIcon) *D.Pixbuf
	statusIconGetScreen        func(s *StatusIcon) *D.Screen
	statusIconGetSize          func(s *StatusIcon) int
	statusIconGetStock         func(s *StatusIcon) string
	statusIconGetStorageType   func(s *StatusIcon) ImageType
	statusIconGetTitle         func(s *StatusIcon) string
	statusIconGetTooltipMarkup func(s *StatusIcon) string
	statusIconGetTooltipText   func(s *StatusIcon) string
	statusIconGetVisible       func(s *StatusIcon) T.Gboolean
	statusIconGetX11WindowId   func(s *StatusIcon) T.GUint32
	statusIconIsEmbedded       func(s *StatusIcon) T.Gboolean
	statusIconSetBlinking      func(s *StatusIcon, blinking T.Gboolean)
	statusIconSetFromFile      func(s *StatusIcon, filename string)
	statusIconSetFromGicon     func(s *StatusIcon, icon *T.GIcon)
	statusIconSetFromIconName  func(s *StatusIcon, iconName string)
	statusIconSetFromPixbuf    func(s *StatusIcon, pixbuf *D.Pixbuf)
	statusIconSetFromStock     func(s *StatusIcon, stockId string)
	statusIconSetHasTooltip    func(s *StatusIcon, hasTooltip T.Gboolean)
	statusIconSetName          func(s *StatusIcon, name string)
	statusIconSetScreen        func(s *StatusIcon, screen *D.Screen)
	statusIconSetTitle         func(s *StatusIcon, title string)
	statusIconSetTooltip       func(s *StatusIcon, tooltipText string)
	statusIconSetTooltipMarkup func(s *StatusIcon, markup string)
	statusIconSetTooltipText   func(s *StatusIcon, text string)
	statusIconSetVisible       func(s *StatusIcon, visible T.Gboolean)
)

func (s *StatusIcon) GetBlinking() T.Gboolean { return statusIconGetBlinking(s) }
func (s *StatusIcon) GetGeometry(screen **D.Screen, area *D.Rectangle, orientation *Orientation) T.Gboolean {
	return statusIconGetGeometry(s, screen, area, orientation)
}
func (s *StatusIcon) GetGicon() *T.GIcon                  { return statusIconGetGicon(s) }
func (s *StatusIcon) GetHasTooltip() T.Gboolean           { return statusIconGetHasTooltip(s) }
func (s *StatusIcon) GetIconName() string                 { return statusIconGetIconName(s) }
func (s *StatusIcon) GetPixbuf() *D.Pixbuf                { return statusIconGetPixbuf(s) }
func (s *StatusIcon) GetScreen() *D.Screen                { return statusIconGetScreen(s) }
func (s *StatusIcon) GetSize() int                        { return statusIconGetSize(s) }
func (s *StatusIcon) GetStock() string                    { return statusIconGetStock(s) }
func (s *StatusIcon) GetStorageType() ImageType           { return statusIconGetStorageType(s) }
func (s *StatusIcon) GetTitle() string                    { return statusIconGetTitle(s) }
func (s *StatusIcon) GetTooltipMarkup() string            { return statusIconGetTooltipMarkup(s) }
func (s *StatusIcon) GetTooltipText() string              { return statusIconGetTooltipText(s) }
func (s *StatusIcon) GetVisible() T.Gboolean              { return statusIconGetVisible(s) }
func (s *StatusIcon) GetX11WindowId() T.GUint32           { return statusIconGetX11WindowId(s) }
func (s *StatusIcon) IsEmbedded() T.Gboolean              { return statusIconIsEmbedded(s) }
func (s *StatusIcon) SetBlinking(blinking T.Gboolean)     { statusIconSetBlinking(s, blinking) }
func (s *StatusIcon) SetFromFile(filename string)         { statusIconSetFromFile(s, filename) }
func (s *StatusIcon) SetFromGicon(icon *T.GIcon)          { statusIconSetFromGicon(s, icon) }
func (s *StatusIcon) SetFromIconName(iconName string)     { statusIconSetFromIconName(s, iconName) }
func (s *StatusIcon) SetFromPixbuf(pixbuf *D.Pixbuf)      { statusIconSetFromPixbuf(s, pixbuf) }
func (s *StatusIcon) SetFromStock(stockId string)         { statusIconSetFromStock(s, stockId) }
func (s *StatusIcon) SetHasTooltip(hasTooltip T.Gboolean) { statusIconSetHasTooltip(s, hasTooltip) }
func (s *StatusIcon) SetName(name string)                 { statusIconSetName(s, name) }
func (s *StatusIcon) SetScreen(screen *D.Screen)          { statusIconSetScreen(s, screen) }
func (s *StatusIcon) SetTitle(title string)               { statusIconSetTitle(s, title) }
func (s *StatusIcon) SetTooltip(tooltipText string)       { statusIconSetTooltip(s, tooltipText) }
func (s *StatusIcon) SetTooltipMarkup(markup string)      { statusIconSetTooltipMarkup(s, markup) }
func (s *StatusIcon) SetTooltipText(text string)          { statusIconSetTooltipText(s, text) }
func (s *StatusIcon) SetVisible(visible T.Gboolean)       { statusIconSetVisible(s, visible) }

type StockItem struct {
	StockId           *T.Gchar
	Label             *T.Gchar
	Modifier          T.GdkModifierType
	Keyval            uint
	TranslationDomain *T.Gchar
}

var (
	StockListIds          func() *T.GSList
	StockLookup           func(stockId string, item *StockItem) T.Gboolean
	StockSetTranslateFunc func(domain string, f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify)

	StockAdd       func(s *StockItem, nItems uint)
	StockAddStatic func(s *StockItem, nItems uint)
	StockItemCopy  func(s *StockItem) *StockItem
	StockItemFree  func(s *StockItem)
)

type Style struct {
	Parent          T.GObject
	Fg              [5]D.Color
	Bg              [5]D.Color
	Light           [5]D.Color
	Dark            [5]D.Color
	Mid             [5]D.Color
	Text            [5]D.Color
	Base            [5]D.Color
	TextAa          [5]D.Color
	Black           D.Color
	White           D.Color
	FontDesc        *T.PangoFontDescription
	Xthickness      int
	Ythickness      int
	FgGc            *[5]D.GC //TODO(t): CHECK
	BgGc            *[5]D.GC //TODO(t): CHECK
	LightGc         *[5]D.GC //TODO(t): CHECK
	DarkGc          *[5]D.GC //TODO(t): CHECK
	MidGc           *[5]D.GC //TODO(t): CHECK
	TextGc          *[5]D.GC //TODO(t): CHECK
	BaseGc          *[5]D.GC //TODO(t): CHECK
	TextAaGc        *[5]D.GC //TODO(t): CHECK
	BlackGc         *D.GC
	WhiteGc         *D.GC
	BgPixmap        *[5]D.Pixmap //TODO(t): CHECK
	AttachCount     int
	Depth           int
	Colormap        *D.Colormap
	PrivateFont     *D.Font
	PrivateFontDesc *T.PangoFontDescription
	RcStyle         *RcStyle
	Styles          *T.GSList
	PropertyCache   *T.GArray
	IconFactories   *T.GSList
}

var (
	StyleGetType func() T.GType
	StyleNew     func() *Style

	styleApplyDefaultBackground func(s *Style, window *D.Window, setBg T.Gboolean, stateType StateType, area *D.Rectangle, x, y, width, height int)
	styleAttach                 func(s *Style, window *D.Window) *Style
	styleCopy                   func(s *Style) *Style
	styleDetach                 func(s *Style)
	styleGet                    func(s *Style, widgetType T.GType, firstPropertyName string, v ...VArg)
	styleGetFont                func(s *Style) *D.Font
	styleGetStyleProperty       func(s *Style, widgetType T.GType, propertyName string, value *T.GValue)
	styleGetValist              func(s *Style, widgetType T.GType, firstPropertyName string, varArgs T.VaList)
	styleLookupColor            func(s *Style, colorName string, color *D.Color) T.Gboolean
	styleLookupIconSet          func(s *Style, stockId string) *IconSet
	styleRef                    func(s *Style) *Style
	styleRenderIcon             func(s *Style, source *IconSource, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *D.Pixbuf
	styleSetBackground          func(s *Style, window *D.Window, stateType StateType)
	styleSetFont                func(s *Style, font *D.Font)
	styleUnref                  func(s *Style)
)

func (s *Style) ApplyDefaultBackground(window *D.Window, setBg T.Gboolean, stateType StateType, area *D.Rectangle, x, y, width, height int) {
	styleApplyDefaultBackground(s, window, setBg, stateType, area, x, y, width, height)
}
func (s *Style) Attach(window *D.Window) *Style { return styleAttach(s, window) }
func (s *Style) Copy() *Style                   { return styleCopy(s) }
func (s *Style) Detach()                        { styleDetach(s) }
func (s *Style) Get(widgetType T.GType, firstPropertyName string, v ...VArg) {
	styleGet(s, widgetType, firstPropertyName, v)
}
func (s *Style) GetFont() *D.Font { return styleGetFont(s) }
func (s *Style) GetStyleProperty(widgetType T.GType, propertyName string, value *T.GValue) {
	styleGetStyleProperty(s, widgetType, propertyName, value)
}
func (s *Style) GetValist(widgetType T.GType, firstPropertyName string, varArgs T.VaList) {
	styleGetValist(s, widgetType, firstPropertyName, varArgs)
}
func (s *Style) LookupColor(colorName string, color *D.Color) T.Gboolean {
	return styleLookupColor(s, colorName, color)
}
func (s *Style) LookupIconSet(stockId string) *IconSet { return styleLookupIconSet(s, stockId) }
func (s *Style) Ref() *Style                           { return styleRef(s) }
func (s *Style) RenderIcon(source *IconSource, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *D.Pixbuf {
	return styleRenderIcon(s, source, direction, state, size, widget, detail)
}
func (s *Style) SetBackground(window *D.Window, stateType StateType) {
	styleSetBackground(s, window, stateType)
}
func (s *Style) SetFont(font *D.Font) { styleSetFont(s, font) }
func (s *Style) Unref()               { styleUnref(s) }

type SubmenuDirection Enum

const (
	DIRECTION_LEFT SubmenuDirection = iota
	DIRECTION_RIGHT
)

var SubmenuDirectionGetType func() T.GType

type SubmenuPlacement Enum

const (
	TOP_BOTTOM SubmenuPlacement = iota
	LEFT_RIGHT
)

var SubmenuPlacementGetType func() T.GType
