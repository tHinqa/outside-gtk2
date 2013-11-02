// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	I "github.com/tHinqa/outside-gtk2/gio"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
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
	ScaleGetType   func() O.Type
	ScaleSetDigits func(scale *Scale, digits int)

	ScaleAddMark          func(s *Scale, value float64, position PositionType, markup string)
	ScaleClearMarks       func(s *Scale)
	ScaleGetDigits        func(s *Scale) int
	ScaleGetDrawValue     func(s *Scale) bool
	ScaleGetLayout        func(s *Scale) *P.Layout
	ScaleGetLayoutOffsets func(s *Scale, x, y *int)
	ScaleGetValuePos      func(s *Scale) PositionType
	ScaleSetDrawValue     func(s *Scale, drawValue bool)
	ScaleSetValuePos      func(s *Scale, pos PositionType)
)

func (s *Scale) AddMark(value float64, position PositionType, markup string) {
	ScaleAddMark(s, value, position, markup)
}
func (s *Scale) ClearMarks()                  { ScaleClearMarks(s) }
func (s *Scale) GetDigits() int               { return ScaleGetDigits(s) }
func (s *Scale) GetDrawValue() bool           { return ScaleGetDrawValue(s) }
func (s *Scale) GetLayout() *P.Layout         { return ScaleGetLayout(s) }
func (s *Scale) GetLayoutOffsets(x, y *int)   { ScaleGetLayoutOffsets(s, x, y) }
func (s *Scale) GetValuePos() PositionType    { return ScaleGetValuePos(s) }
func (s *Scale) SetDrawValue(drawValue bool)  { ScaleSetDrawValue(s, drawValue) }
func (s *Scale) SetValuePos(pos PositionType) { ScaleSetValuePos(s, pos) }

type ScaleButton struct {
	Parent      Button
	PlusButton  *Widget
	MinusButton *Widget
	_           *struct{}
}

var (
	ScaleButtonGetType func() O.Type
	ScaleButtonNew     func(size IconSize, min, max, step float64, icons []string) *Widget

	ScaleButtonGetAdjustment  func(s *ScaleButton) *Adjustment
	ScaleButtonGetMinusButton func(s *ScaleButton) *Widget
	ScaleButtonGetOrientation func(s *ScaleButton) Orientation
	ScaleButtonGetPlusButton  func(s *ScaleButton) *Widget
	ScaleButtonGetPopup       func(s *ScaleButton) *Widget
	ScaleButtonGetValue       func(s *ScaleButton) float64
	ScaleButtonSetAdjustment  func(s *ScaleButton, adjustment *Adjustment)
	ScaleButtonSetIcons       func(s *ScaleButton, icons []string)
	ScaleButtonSetOrientation func(s *ScaleButton, orientation Orientation)
	ScaleButtonSetValue       func(s *ScaleButton, value float64)
)

func (s *ScaleButton) GetAdjustment() *Adjustment           { return ScaleButtonGetAdjustment(s) }
func (s *ScaleButton) GetMinusButton() *Widget              { return ScaleButtonGetMinusButton(s) }
func (s *ScaleButton) GetOrientation() Orientation          { return ScaleButtonGetOrientation(s) }
func (s *ScaleButton) GetPlusButton() *Widget               { return ScaleButtonGetPlusButton(s) }
func (s *ScaleButton) GetPopup() *Widget                    { return ScaleButtonGetPopup(s) }
func (s *ScaleButton) GetValue() float64                    { return ScaleButtonGetValue(s) }
func (s *ScaleButton) SetAdjustment(adjustment *Adjustment) { ScaleButtonSetAdjustment(s, adjustment) }
func (s *ScaleButton) SetIcons(icons []string)              { ScaleButtonSetIcons(s, icons) }
func (s *ScaleButton) SetOrientation(orientation Orientation) {
	ScaleButtonSetOrientation(s, orientation)
}
func (s *ScaleButton) SetValue(value float64) { ScaleButtonSetValue(s, value) }

var ScrollbarGetType func() O.Type //TODO(t):Use?

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
	ScrolledWindowGetType func() O.Type
	ScrolledWindowNew     func(hadjustment, vadjustment *Adjustment) *Widget

	ScrolledWindowAddWithViewport func(s *ScrolledWindow, child *Widget)
	ScrolledWindowGetHadjustment  func(s *ScrolledWindow) *Adjustment
	ScrolledWindowGetHscrollbar   func(s *ScrolledWindow) *Widget
	ScrolledWindowGetPlacement    func(s *ScrolledWindow) CornerType
	ScrolledWindowGetPolicy       func(s *ScrolledWindow, hscrollbarPolicy, vscrollbarPolicy *PolicyType)
	ScrolledWindowGetShadowType   func(s *ScrolledWindow) ShadowType
	ScrolledWindowGetVadjustment  func(s *ScrolledWindow) *Adjustment
	ScrolledWindowGetVscrollbar   func(s *ScrolledWindow) *Widget
	ScrolledWindowSetHadjustment  func(s *ScrolledWindow, hadjustment *Adjustment)
	ScrolledWindowSetPlacement    func(s *ScrolledWindow, windowPlacement CornerType)
	ScrolledWindowSetPolicy       func(s *ScrolledWindow, hscrollbarPolicy, vscrollbarPolicy PolicyType)
	ScrolledWindowSetShadowType   func(s *ScrolledWindow, t ShadowType)
	ScrolledWindowSetVadjustment  func(s *ScrolledWindow, vadjustment *Adjustment)
	ScrolledWindowUnsetPlacement  func(s *ScrolledWindow)
)

func (s *ScrolledWindow) AddWithViewport(child *Widget) {
	ScrolledWindowAddWithViewport(s, child)
}
func (s *ScrolledWindow) GetHadjustment() *Adjustment { return ScrolledWindowGetHadjustment(s) }
func (s *ScrolledWindow) GetHscrollbar() *Widget      { return ScrolledWindowGetHscrollbar(s) }
func (s *ScrolledWindow) GetPlacement() CornerType    { return ScrolledWindowGetPlacement(s) }
func (s *ScrolledWindow) GetPolicy(hscrollbarPolicy, vscrollbarPolicy *PolicyType) {
	ScrolledWindowGetPolicy(s, hscrollbarPolicy, vscrollbarPolicy)
}
func (s *ScrolledWindow) GetShadowType() ShadowType   { return ScrolledWindowGetShadowType(s) }
func (s *ScrolledWindow) GetVadjustment() *Adjustment { return ScrolledWindowGetVadjustment(s) }
func (s *ScrolledWindow) GetVscrollbar() *Widget      { return ScrolledWindowGetVscrollbar(s) }
func (s *ScrolledWindow) SetHadjustment(hadjustment *Adjustment) {
	ScrolledWindowSetHadjustment(s, hadjustment)
}
func (s *ScrolledWindow) SetPlacement(windowPlacement CornerType) {
	ScrolledWindowSetPlacement(s, windowPlacement)
}
func (s *ScrolledWindow) SetPolicy(hscrollbarPolicy, vscrollbarPolicy PolicyType) {
	ScrolledWindowSetPolicy(s, hscrollbarPolicy, vscrollbarPolicy)
}
func (s *ScrolledWindow) SetShadowType(t ShadowType) { ScrolledWindowSetShadowType(s, t) }
func (s *ScrolledWindow) SetVadjustment(vadjustment *Adjustment) {
	ScrolledWindowSetVadjustment(s, vadjustment)
}
func (s *ScrolledWindow) UnsetPlacement() { ScrolledWindowUnsetPlacement(s) }

var ScrollStepGetType func() O.Type //TODO(t):Use?

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

var ScrollTypeGetType func() O.Type

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
	SelectionDataGetType func() O.Type

	SelectionAddTarget          func(widget *Widget, selection, target D.Atom, info uint)
	SelectionAddTargets         func(widget *Widget, selection D.Atom, targets *TargetEntry, ntargets uint)
	SelectionClear              func(widget *Widget, event *D.EventSelection) bool
	SelectionClearTargets       func(widget *Widget, selection D.Atom)
	SelectionConvert            func(widget *Widget, selection, target D.Atom, time T.GUint32) bool
	SelectionDataCopy           func(s *SelectionData) *SelectionData
	SelectionOwnerSet           func(widget *Widget, selection D.Atom, time T.GUint32) bool
	SelectionOwnerSetForDisplay func(display *D.Display, widget *Widget, selection D.Atom, time T.GUint32) bool
	SelectionRemoveAll          func(widget *Widget)

	SelectionDataFree                   func(s *SelectionData)
	SelectionDataGetData                func(s *SelectionData) *T.Guchar
	SelectionDataGetDataType            func(s *SelectionData) D.Atom
	SelectionDataGetDisplay             func(s *SelectionData) *D.Display
	SelectionDataGetFormat              func(s *SelectionData) int
	SelectionDataGetLength              func(s *SelectionData) int
	SelectionDataGetPixbuf              func(s *SelectionData) *D.Pixbuf
	SelectionDataGetSelection           func(s *SelectionData) D.Atom
	SelectionDataGetTarget              func(s *SelectionData) D.Atom
	SelectionDataGetTargets             func(s *SelectionData, targets **D.Atom, nAtoms *int) bool
	SelectionDataGetText                func(s *SelectionData) []string //TODO(t): *T.Guchar?
	SelectionDataGetUris                func(s *SelectionData) []string
	SelectionDataSet                    func(s *SelectionData, typ D.Atom, format int, data *T.Guchar, length int)
	SelectionDataSetPixbuf              func(s *SelectionData, pixbuf *D.Pixbuf) bool
	SelectionDataSetText                func(s *SelectionData, str string, len int) bool
	SelectionDataSetUris                func(s *SelectionData, uris []string) bool
	SelectionDataTargetsIncludeImage    func(s *SelectionData, writable bool) bool
	SelectionDataTargetsIncludeRichText func(s *SelectionData, buffer *TextBuffer) bool
	SelectionDataTargetsIncludeText     func(s *SelectionData) bool
	SelectionDataTargetsIncludeUri      func(s *SelectionData) bool
)

func (s *SelectionData) Copy() *SelectionData   { return SelectionDataCopy(s) }
func (s *SelectionData) Free()                  { SelectionDataFree(s) }
func (s *SelectionData) GetData() *T.Guchar     { return SelectionDataGetData(s) }
func (s *SelectionData) GetDataType() D.Atom    { return SelectionDataGetDataType(s) }
func (s *SelectionData) GetDisplay() *D.Display { return SelectionDataGetDisplay(s) }
func (s *SelectionData) GetFormat() int         { return SelectionDataGetFormat(s) }
func (s *SelectionData) GetLength() int         { return SelectionDataGetLength(s) }
func (s *SelectionData) GetPixbuf() *D.Pixbuf   { return SelectionDataGetPixbuf(s) }
func (s *SelectionData) GetSelection() D.Atom   { return SelectionDataGetSelection(s) }
func (s *SelectionData) GetTarget() D.Atom      { return SelectionDataGetTarget(s) }
func (s *SelectionData) GetTargets(targets **D.Atom, nAtoms *int) bool {
	return SelectionDataGetTargets(s, targets, nAtoms)
}
func (s *SelectionData) GetText() []string { return SelectionDataGetText(s) } //TODO(t): was*T.Guchar?
func (s *SelectionData) GetUris() []string { return SelectionDataGetUris(s) }
func (s *SelectionData) Set(typ D.Atom, format int, data *T.Guchar, length int) {
	SelectionDataSet(s, typ, format, data, length)
}
func (s *SelectionData) SetPixbuf(pixbuf *D.Pixbuf) bool {
	return SelectionDataSetPixbuf(s, pixbuf)
}
func (s *SelectionData) SetText(str string, leng int) bool {
	return SelectionDataSetText(s, str, leng)
}
func (s *SelectionData) SetUris(uris []string) bool { return SelectionDataSetUris(s, uris) }
func (s *SelectionData) TargetsIncludeImage(writable bool) bool {
	return SelectionDataTargetsIncludeImage(s, writable)
}
func (s *SelectionData) TargetsIncludeRichText(buffer *TextBuffer) bool {
	return SelectionDataTargetsIncludeRichText(s, buffer)
}
func (s *SelectionData) TargetsIncludeText() bool { return SelectionDataTargetsIncludeText(s) }
func (s *SelectionData) TargetsIncludeUri() bool  { return SelectionDataTargetsIncludeUri(s) }

type SelectionMode Enum

const (
	SELECTION_NONE SelectionMode = iota
	SELECTION_SINGLE
	SELECTION_BROWSE
	SELECTION_MULTIPLE
	SELECTION_EXTENDED = SELECTION_MULTIPLE
)

var SelectionModeGetType func() O.Type

type SensitivityType Enum

const (
	SENSITIVITY_AUTO SensitivityType = iota
	SENSITIVITY_ON
	SENSITIVITY_OFF
)

var SensitivityTypeGetType func() O.Type

var (
	SeparatorGetType func() O.Type

	SeparatorMenuItemGetType func() O.Type
	SeparatorMenuItemNew     func() *Widget
)

type SeparatorToolItem struct {
	Parent ToolItem
	_      *struct{}
}

var (
	SeparatorToolItemGetType func() O.Type
	SeparatorToolItemNew     func() *ToolItem

	SeparatorToolItemGetDraw func(s *SeparatorToolItem) bool
	SeparatorToolItemSetDraw func(s *SeparatorToolItem, draw bool)
)

func (s *SeparatorToolItem) GetDraw() bool     { return SeparatorToolItemGetDraw(s) }
func (s *SeparatorToolItem) SetDraw(draw bool) { SeparatorToolItemSetDraw(s, draw) }

type Settings struct {
	Parent          O.Object
	QueuedSettings  *T.GData
	Property_values *SettingsPropertyValue
	RcContext       *RcContext
	Screen          *D.Screen
}

var (
	SettingsGetType    func() O.Type
	SettingsGetDefault func() *Settings

	SettingsGetForScreen          func(screen *D.Screen) *Settings
	SettingsInstallProperty       func(pspec *T.GParamSpec)
	SettingsInstallPropertyParser func(pspec *T.GParamSpec, parser RcPropertyParser)

	SettingsSetDoubleProperty func(s *Settings, name string, vDouble float64, origin string)
	SettingsSetLongProperty   func(s *Settings, name string, vLong T.Glong, origin string)
	SettingsSetPropertyValue  func(s *Settings, name string, svalue *SettingsValue)
	SettingsSetStringProperty func(s *Settings, name, vString, origin string)
)

func (s *Settings) SetDoubleProperty(name string, vDouble float64, origin string) {
	SettingsSetDoubleProperty(s, name, vDouble, origin)
}
func (s *Settings) SetLongProperty(name string, vLong T.Glong, origin string) {
	SettingsSetLongProperty(s, name, vLong, origin)
}
func (s *Settings) SetPropertyValue(name string, svalue *SettingsValue) {
	SettingsSetPropertyValue(s, name, svalue)
}
func (s *Settings) SetStringProperty(name, vString, origin string) {
	SettingsSetStringProperty(s, name, vString, origin)
}

type SettingsPropertyValue struct{}

type SettingsValue struct {
	Origin *T.Gchar
	Value  O.Value
}

type ShadowType Enum

const (
	SHADOW_NONE ShadowType = iota
	SHADOW_IN
	SHADOW_OUT
	SHADOW_ETCHED_IN
	SHADOW_ETCHED_OUT
)

var ShadowTypeGetType func() O.Type

type SideType Enum

const (
	SIDE_TOP SideType = iota
	SIDE_BOTTOM
	SIDE_LEFT
	SIDE_RIGHT
)

var SideTypeGetType func() O.Type

type SignalRunType Enum

const (
	RUN_FIRST      SignalRunType = SignalRunType(O.SIGNAL_RUN_FIRST)
	RUN_LAST                     = SignalRunType(O.SIGNAL_RUN_LAST)
	RUN_BOTH                     = SignalRunType(RUN_FIRST | RUN_LAST)
	RUN_NO_RECURSE               = SignalRunType(O.SIGNAL_NO_RECURSE)
	RUN_ACTION                   = SignalRunType(O.SIGNAL_ACTION)
	RUN_NO_HOOKS                 = SignalRunType(O.SIGNAL_NO_HOOKS)
)

var SignalRunTypeGetType func() O.Type

type SizeGroup struct {
	Parent  O.Object
	Widgets *L.SList
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
	SizeGroupGetType func() O.Type
	SizeGroupNew     func(mode SizeGroupMode) *SizeGroup

	SizeGroupModeGetType func() O.Type

	SizeGroupAddWidget       func(s *SizeGroup, widget *Widget)
	SizeGroupGetIgnoreHidden func(s *SizeGroup) bool
	SizeGroupGetMode         func(s *SizeGroup) SizeGroupMode
	SizeGroupGetWidgets      func(s *SizeGroup) *L.SList
	SizeGroupRemoveWidget    func(s *SizeGroup, widget *Widget)
	SizeGroupSetIgnoreHidden func(s *SizeGroup, ignoreHidden bool)
	SizeGroupSetMode         func(s *SizeGroup, mode SizeGroupMode)
)

func (s *SizeGroup) AddWidget(widget *Widget)    { SizeGroupAddWidget(s, widget) }
func (s *SizeGroup) GetIgnoreHidden() bool       { return SizeGroupGetIgnoreHidden(s) }
func (s *SizeGroup) GetMode() SizeGroupMode      { return SizeGroupGetMode(s) }
func (s *SizeGroup) GetWidgets() *L.SList        { return SizeGroupGetWidgets(s) }
func (s *SizeGroup) RemoveWidget(widget *Widget) { SizeGroupRemoveWidget(s, widget) }
func (s *SizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	SizeGroupSetIgnoreHidden(s, ignoreHidden)
}
func (s *SizeGroup) SetMode(mode SizeGroupMode) { SizeGroupSetMode(s, mode) }

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
	SocketGetType func() O.Type
	SocketNew     func() *Widget

	SocketAddId         func(s *Socket, windowId T.GdkNativeWindow)
	SocketGetId         func(s *Socket) T.GdkNativeWindow
	SocketGetPlugWindow func(s *Socket) *D.Window
	SocketSteal         func(s *Socket, wid T.GdkNativeWindow)
)

func (s *Socket) AddId(windowId T.GdkNativeWindow) { SocketAddId(s, windowId) }
func (s *Socket) GetId() T.GdkNativeWindow         { return SocketGetId(s) }
func (s *Socket) GetPlugWindow() *D.Window         { return SocketGetPlugWindow(s) }
func (s *Socket) Steal(wid T.GdkNativeWindow)      { SocketSteal(s, wid) }

type SortType Enum

const (
	SORT_ASCENDING SortType = iota
	SORT_DESCENDING
)

var SortTypeGetType func() O.Type

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
	SpinButtonGetType      func() O.Type
	SpinButtonNew          func(adjustment *Adjustment, climbRate float64, digits uint) *Widget
	SpinButtonNewWithRange func(min, max, step float64) *Widget

	SpinButtonUpdatePolicyGetType func() O.Type

	SpinButtonConfigure       func(s *SpinButton, adjustment *Adjustment, climbRate float64, digits uint)
	SpinButtonGetAdjustment   func(s *SpinButton) *Adjustment
	SpinButtonGetDigits       func(s *SpinButton) uint
	SpinButtonGetIncrements   func(s *SpinButton, step, page *float64)
	SpinButtonGetNumeric      func(s *SpinButton) bool
	SpinButtonGetRange        func(s *SpinButton, min, max *float64)
	SpinButtonGetSnapToTicks  func(s *SpinButton) bool
	SpinButtonGetUpdatePolicy func(s *SpinButton) SpinButtonUpdatePolicy
	SpinButtonGetValue        func(s *SpinButton) float64
	SpinButtonGetValueAsInt   func(s *SpinButton) int
	SpinButtonGetWrap         func(s *SpinButton) bool
	SpinButtonSetAdjustment   func(s *SpinButton, adjustment *Adjustment)
	SpinButtonSetDigits       func(s *SpinButton, digits uint)
	SpinButtonSetIncrements   func(s *SpinButton, step, page float64)
	SpinButtonSetNumeric      func(s *SpinButton, numeric bool)
	SpinButtonSetRange        func(s *SpinButton, min, max float64)
	SpinButtonSetSnapToTicks  func(s *SpinButton, snapToTicks bool)
	SpinButtonSetUpdatePolicy func(s *SpinButton, policy SpinButtonUpdatePolicy)
	SpinButtonSetValue        func(s *SpinButton, value float64)
	SpinButtonSetWrap         func(s *SpinButton, wrap bool)
	SpinButtonSpin            func(s *SpinButton, direction SpinType, increment float64)
	SpinButtonUpdate          func(s *SpinButton)
)

func (s *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint) {
	SpinButtonConfigure(s, adjustment, climbRate, digits)
}
func (s *SpinButton) GetAdjustment() *Adjustment              { return SpinButtonGetAdjustment(s) }
func (s *SpinButton) GetDigits() uint                         { return SpinButtonGetDigits(s) }
func (s *SpinButton) GetIncrements(step, page *float64)       { SpinButtonGetIncrements(s, step, page) }
func (s *SpinButton) GetNumeric() bool                        { return SpinButtonGetNumeric(s) }
func (s *SpinButton) GetRange(min, max *float64)              { SpinButtonGetRange(s, min, max) }
func (s *SpinButton) GetSnapToTicks() bool                    { return SpinButtonGetSnapToTicks(s) }
func (s *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy { return SpinButtonGetUpdatePolicy(s) }
func (s *SpinButton) GetValue() float64                       { return SpinButtonGetValue(s) }
func (s *SpinButton) GetValueAsInt() int                      { return SpinButtonGetValueAsInt(s) }
func (s *SpinButton) GetWrap() bool                           { return SpinButtonGetWrap(s) }
func (s *SpinButton) SetAdjustment(adjustment *Adjustment)    { SpinButtonSetAdjustment(s, adjustment) }
func (s *SpinButton) SetDigits(digits uint)                   { SpinButtonSetDigits(s, digits) }
func (s *SpinButton) SetIncrements(step, page float64)        { SpinButtonSetIncrements(s, step, page) }
func (s *SpinButton) SetNumeric(numeric bool)                 { SpinButtonSetNumeric(s, numeric) }
func (s *SpinButton) SetRange(min, max float64)               { SpinButtonSetRange(s, min, max) }
func (s *SpinButton) SetSnapToTicks(snapToTicks bool)         { SpinButtonSetSnapToTicks(s, snapToTicks) }
func (s *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	SpinButtonSetUpdatePolicy(s, policy)
}
func (s *SpinButton) SetValue(value float64) { SpinButtonSetValue(s, value) }
func (s *SpinButton) SetWrap(wrap bool)      { SpinButtonSetWrap(s, wrap) }
func (s *SpinButton) Spin(direction SpinType, increment float64) {
	SpinButtonSpin(s, direction, increment)
}
func (s *SpinButton) Update() { SpinButtonUpdate(s) }

type Spinner struct {
	Parent DrawingArea
	_      *struct{}
}

var (
	SpinnerGetType func() O.Type
	SpinnerNew     func() *Widget

	SpinnerStart func(s *Spinner)
	SpinnerStop  func(s *Spinner)
)

func (s *Spinner) Start() { SpinnerStart(s) }
func (s *Spinner) Stop()  { SpinnerStop(s) }

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

var SpinTypeGetType func() O.Type

type StateType Enum

const (
	STATE_NORMAL StateType = iota
	STATE_ACTIVE
	STATE_PRELIGHT
	STATE_SELECTED
	STATE_INSENSITIVE
)

var StateTypeGetType func() O.Type

type Statusbar struct {
	ParentWidget  HBox
	Frame         *Widget
	Label         *Widget
	Messages      *L.SList
	Keys          *L.SList
	SeqContextId  uint
	SeqMessageId  uint
	Grip_window   *D.Window
	HasResizeGrip uint // : 1
}

var (
	StatusbarGetType func() O.Type
	StatusbarNew     func() *Widget

	StatusbarGetContextId     func(s *Statusbar, contextDescription string) uint
	StatusbarGetHasResizeGrip func(s *Statusbar) bool
	StatusbarGetMessageArea   func(s *Statusbar) *Widget
	StatusbarPop              func(s *Statusbar, contextId uint)
	StatusbarPush             func(s *Statusbar, contextId uint, text string) uint
	StatusbarRemove           func(s *Statusbar, contextId uint, messageId uint)
	StatusbarRemoveAll        func(s *Statusbar, contextId uint)
	StatusbarSetHasResizeGrip func(s *Statusbar, setting bool)
)

func (s *Statusbar) GetContextId(contextDescription string) uint {
	return StatusbarGetContextId(s, contextDescription)
}
func (s *Statusbar) GetHasResizeGrip() bool                { return StatusbarGetHasResizeGrip(s) }
func (s *Statusbar) GetMessageArea() *Widget               { return StatusbarGetMessageArea(s) }
func (s *Statusbar) Pop(contextId uint)                    { StatusbarPop(s, contextId) }
func (s *Statusbar) Push(contextId uint, text string) uint { return StatusbarPush(s, contextId, text) }
func (s *Statusbar) Remove(contextId uint, messageId uint) { StatusbarRemove(s, contextId, messageId) }
func (s *Statusbar) RemoveAll(contextId uint)              { StatusbarRemoveAll(s, contextId) }
func (s *Statusbar) SetHasResizeGrip(setting bool)         { StatusbarSetHasResizeGrip(s, setting) }

type StatusIcon struct {
	Parent O.Object
	_      *struct{}
}

var (
	StatusIconGetType         func() O.Type
	StatusIconNew             func() *StatusIcon
	StatusIconNewFromFile     func(filename string) *StatusIcon
	StatusIconNewFromGicon    func(icon *I.Icon) *StatusIcon
	StatusIconNewFromIconName func(iconName string) *StatusIcon
	StatusIconNewFromPixbuf   func(pixbuf *D.Pixbuf) *StatusIcon
	StatusIconNewFromStock    func(stockId string) *StatusIcon

	StatusIconPositionMenu func(menu *Menu, x, y *int, pushIn *bool, userData T.Gpointer)

	StatusIconGetBlinking      func(s *StatusIcon) bool
	StatusIconGetGeometry      func(s *StatusIcon, screen **D.Screen, area *D.Rectangle, orientation *Orientation) bool
	StatusIconGetGicon         func(s *StatusIcon) *I.Icon
	StatusIconGetHasTooltip    func(s *StatusIcon) bool
	StatusIconGetIconName      func(s *StatusIcon) string
	StatusIconGetPixbuf        func(s *StatusIcon) *D.Pixbuf
	StatusIconGetScreen        func(s *StatusIcon) *D.Screen
	StatusIconGetSize          func(s *StatusIcon) int
	StatusIconGetStock         func(s *StatusIcon) string
	StatusIconGetStorageType   func(s *StatusIcon) ImageType
	StatusIconGetTitle         func(s *StatusIcon) string
	StatusIconGetTooltipMarkup func(s *StatusIcon) string
	StatusIconGetTooltipText   func(s *StatusIcon) string
	StatusIconGetVisible       func(s *StatusIcon) bool
	StatusIconGetX11WindowId   func(s *StatusIcon) T.GUint32
	StatusIconIsEmbedded       func(s *StatusIcon) bool
	StatusIconSetBlinking      func(s *StatusIcon, blinking bool)
	StatusIconSetFromFile      func(s *StatusIcon, filename string)
	StatusIconSetFromGicon     func(s *StatusIcon, icon *I.Icon)
	StatusIconSetFromIconName  func(s *StatusIcon, iconName string)
	StatusIconSetFromPixbuf    func(s *StatusIcon, pixbuf *D.Pixbuf)
	StatusIconSetFromStock     func(s *StatusIcon, stockId string)
	StatusIconSetHasTooltip    func(s *StatusIcon, hasTooltip bool)
	StatusIconSetName          func(s *StatusIcon, name string)
	StatusIconSetScreen        func(s *StatusIcon, screen *D.Screen)
	StatusIconSetTitle         func(s *StatusIcon, title string)
	StatusIconSetTooltip       func(s *StatusIcon, tooltipText string)
	StatusIconSetTooltipMarkup func(s *StatusIcon, markup string)
	StatusIconSetTooltipText   func(s *StatusIcon, text string)
	StatusIconSetVisible       func(s *StatusIcon, visible bool)
)

func (s *StatusIcon) GetBlinking() bool { return StatusIconGetBlinking(s) }
func (s *StatusIcon) GetGeometry(screen **D.Screen, area *D.Rectangle, orientation *Orientation) bool {
	return StatusIconGetGeometry(s, screen, area, orientation)
}
func (s *StatusIcon) GetGicon() *I.Icon               { return StatusIconGetGicon(s) }
func (s *StatusIcon) GetHasTooltip() bool             { return StatusIconGetHasTooltip(s) }
func (s *StatusIcon) GetIconName() string             { return StatusIconGetIconName(s) }
func (s *StatusIcon) GetPixbuf() *D.Pixbuf            { return StatusIconGetPixbuf(s) }
func (s *StatusIcon) GetScreen() *D.Screen            { return StatusIconGetScreen(s) }
func (s *StatusIcon) GetSize() int                    { return StatusIconGetSize(s) }
func (s *StatusIcon) GetStock() string                { return StatusIconGetStock(s) }
func (s *StatusIcon) GetStorageType() ImageType       { return StatusIconGetStorageType(s) }
func (s *StatusIcon) GetTitle() string                { return StatusIconGetTitle(s) }
func (s *StatusIcon) GetTooltipMarkup() string        { return StatusIconGetTooltipMarkup(s) }
func (s *StatusIcon) GetTooltipText() string          { return StatusIconGetTooltipText(s) }
func (s *StatusIcon) GetVisible() bool                { return StatusIconGetVisible(s) }
func (s *StatusIcon) GetX11WindowId() T.GUint32       { return StatusIconGetX11WindowId(s) }
func (s *StatusIcon) IsEmbedded() bool                { return StatusIconIsEmbedded(s) }
func (s *StatusIcon) SetBlinking(blinking bool)       { StatusIconSetBlinking(s, blinking) }
func (s *StatusIcon) SetFromFile(filename string)     { StatusIconSetFromFile(s, filename) }
func (s *StatusIcon) SetFromGicon(icon *I.Icon)       { StatusIconSetFromGicon(s, icon) }
func (s *StatusIcon) SetFromIconName(iconName string) { StatusIconSetFromIconName(s, iconName) }
func (s *StatusIcon) SetFromPixbuf(pixbuf *D.Pixbuf)  { StatusIconSetFromPixbuf(s, pixbuf) }
func (s *StatusIcon) SetFromStock(stockId string)     { StatusIconSetFromStock(s, stockId) }
func (s *StatusIcon) SetHasTooltip(hasTooltip bool)   { StatusIconSetHasTooltip(s, hasTooltip) }
func (s *StatusIcon) SetName(name string)             { StatusIconSetName(s, name) }
func (s *StatusIcon) SetScreen(screen *D.Screen)      { StatusIconSetScreen(s, screen) }
func (s *StatusIcon) SetTitle(title string)           { StatusIconSetTitle(s, title) }
func (s *StatusIcon) SetTooltip(tooltipText string)   { StatusIconSetTooltip(s, tooltipText) }
func (s *StatusIcon) SetTooltipMarkup(markup string)  { StatusIconSetTooltipMarkup(s, markup) }
func (s *StatusIcon) SetTooltipText(text string)      { StatusIconSetTooltipText(s, text) }
func (s *StatusIcon) SetVisible(visible bool)         { StatusIconSetVisible(s, visible) }

type StockItem struct {
	StockId           *T.Gchar
	Label             *T.Gchar
	Modifier          T.GdkModifierType
	Keyval            uint
	TranslationDomain *T.Gchar
}

var (
	StockListIds          func() *L.SList
	StockLookup           func(stockId string, item *StockItem) bool
	StockSetTranslateFunc func(domain string, f TranslateFunc, data T.Gpointer, notify O.DestroyNotify)

	StockAdd       func(s *StockItem, nItems uint)
	StockAddStatic func(s *StockItem, nItems uint)
	StockItemCopy  func(s *StockItem) *StockItem
	StockItemFree  func(s *StockItem)
)

type Style struct {
	Parent          O.Object
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
	FontDesc        *P.FontDescription
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
	PrivateFontDesc *P.FontDescription
	RcStyle         *RcStyle
	Styles          *L.SList
	PropertyCache   *L.Array
	IconFactories   *L.SList
}

var (
	StyleGetType func() O.Type
	StyleNew     func() *Style

	StyleApplyDefaultBackground func(s *Style, window *D.Window, setBg bool, stateType StateType, area *D.Rectangle, x, y, width, height int)
	StyleAttach                 func(s *Style, window *D.Window) *Style
	StyleCopy                   func(s *Style) *Style
	StyleDetach                 func(s *Style)
	StyleGet                    func(s *Style, widgetType O.Type, firstPropertyName string, v ...VArg)
	StyleGetFont                func(s *Style) *D.Font
	StyleGetStyleProperty       func(s *Style, widgetType O.Type, propertyName string, value *O.Value)
	StyleGetValist              func(s *Style, widgetType O.Type, firstPropertyName string, varArgs T.VaList)
	StyleLookupColor            func(s *Style, colorName string, color *D.Color) bool
	StyleLookupIconSet          func(s *Style, stockId string) *IconSet
	StyleRef                    func(s *Style) *Style
	StyleRenderIcon             func(s *Style, source *IconSource, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *D.Pixbuf
	StyleSetBackground          func(s *Style, window *D.Window, stateType StateType)
	StyleSetFont                func(s *Style, font *D.Font)
	StyleUnref                  func(s *Style)
)

func (s *Style) ApplyDefaultBackground(window *D.Window, setBg bool, stateType StateType, area *D.Rectangle, x, y, width, height int) {
	StyleApplyDefaultBackground(s, window, setBg, stateType, area, x, y, width, height)
}
func (s *Style) Attach(window *D.Window) *Style { return StyleAttach(s, window) }
func (s *Style) Copy() *Style                   { return StyleCopy(s) }
func (s *Style) Detach()                        { StyleDetach(s) }
func (s *Style) Get(widgetType O.Type, firstPropertyName string, v ...VArg) {
	StyleGet(s, widgetType, firstPropertyName, v)
}
func (s *Style) GetFont() *D.Font { return StyleGetFont(s) }
func (s *Style) GetStyleProperty(widgetType O.Type, propertyName string, value *O.Value) {
	StyleGetStyleProperty(s, widgetType, propertyName, value)
}
func (s *Style) GetValist(widgetType O.Type, firstPropertyName string, varArgs T.VaList) {
	StyleGetValist(s, widgetType, firstPropertyName, varArgs)
}
func (s *Style) LookupColor(colorName string, color *D.Color) bool {
	return StyleLookupColor(s, colorName, color)
}
func (s *Style) LookupIconSet(stockId string) *IconSet { return StyleLookupIconSet(s, stockId) }
func (s *Style) Ref() *Style                           { return StyleRef(s) }
func (s *Style) RenderIcon(source *IconSource, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *D.Pixbuf {
	return StyleRenderIcon(s, source, direction, state, size, widget, detail)
}
func (s *Style) SetBackground(window *D.Window, stateType StateType) {
	StyleSetBackground(s, window, stateType)
}
func (s *Style) SetFont(font *D.Font) { StyleSetFont(s, font) }
func (s *Style) Unref()               { StyleUnref(s) }

type SubmenuDirection Enum

const (
	DIRECTION_LEFT SubmenuDirection = iota
	DIRECTION_RIGHT
)

var SubmenuDirectionGetType func() O.Type

type SubmenuPlacement Enum

const (
	TOP_BOTTOM SubmenuPlacement = iota
	LEFT_RIGHT
)

var SubmenuPlacementGetType func() O.Type
