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
)

type (
	Calendar struct {
		Widget            Widget
		Header_style      *Style
		Label_style       *Style
		Month             int
		Year              int
		Selected_day      int
		Day_month         [6][7]int
		Day               [6][7]int
		Num_marked_dates  int
		Marked_date       [31]int
		Display_flags     CalendarDisplayOptions
		Marked_date_color [31]D.Color
		Gc                *D.GC
		Xor_gc            *D.GC
		Focus_row         int
		Focus_col         int
		Highlight_row     int
		Highlight_col     int
		_                 *struct{}
		Grow_space        [32]T.Gchar
		_, _, _, _        func()
	}

	CalendarDetailFunc func(calendar *Calendar,
		year, month, day uint, userData T.Gpointer) string
)

type CalendarDisplayOptions Enum

const (
	CALENDAR_SHOW_HEADING CalendarDisplayOptions = 1 << iota
	CALENDAR_SHOW_DAY_NAMES
	CALENDAR_NO_MONTH_CHANGE
	CALENDAR_SHOW_WEEK_NUMBERS
	CALENDAR_WEEK_START_MONDAY
	CALENDAR_SHOW_DETAILS
)

var (
	CalendarDisplayOptionsGetType func() O.Type
	CalendarGetType               func() O.Type
	CalendarNew                   func() *Widget
)

var (
	CalendarClearMarks          func(c *Calendar)
	CalendarDisplayOptions_     func(c *Calendar, flags CalendarDisplayOptions)
	CalendarFreeze              func(c *Calendar)
	CalendarGetDate             func(c *Calendar, year, month, day *uint)
	CalendarGetDetailHeightRows func(c *Calendar) int
	CalendarGetDetailWidthChars func(c *Calendar) int
	CalendarGetDisplayOptions   func(c *Calendar) CalendarDisplayOptions
	CalendarMarkDay             func(c *Calendar, day uint) bool
	CalendarSelectDay           func(c *Calendar, day uint)
	CalendarSelectMonth         func(c *Calendar, month, year uint) bool
	CalendarSetDetailFunc       func(c *Calendar, f CalendarDetailFunc, data T.Gpointer, destroy T.GDestroyNotify)
	CalendarSetDetailHeightRows func(c *Calendar, rows int)
	CalendarSetDetailWidthChars func(c *Calendar, chars int)
	CalendarSetDisplayOptions   func(c *Calendar, flags CalendarDisplayOptions)
	CalendarThaw                func(c *Calendar)
	CalendarUnmarkDay           func(c *Calendar, day uint) bool
)

func (c *Calendar) ClearMarks()                                 { CalendarClearMarks(c) }
func (c *Calendar) DisplayOptions(flags CalendarDisplayOptions) { CalendarDisplayOptions_(c, flags) }
func (c *Calendar) Freeze()                                     { CalendarFreeze(c) }
func (c *Calendar) GetDate(year, month, day *uint)              { CalendarGetDate(c, year, month, day) }
func (c *Calendar) GetDetailHeightRows() int                    { return CalendarGetDetailHeightRows(c) }
func (c *Calendar) GetDetailWidthChars() int                    { return CalendarGetDetailWidthChars(c) }
func (c *Calendar) GetDisplayOptions() CalendarDisplayOptions   { return CalendarGetDisplayOptions(c) }
func (c *Calendar) MarkDay(day uint) bool                       { return CalendarMarkDay(c, day) }
func (c *Calendar) SelectDay(day uint)                          { CalendarSelectDay(c, day) }
func (c *Calendar) SelectMonth(month, year uint) bool {
	return CalendarSelectMonth(c, month, year)
}
func (c *Calendar) SetDetailFunc(f CalendarDetailFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	CalendarSetDetailFunc(c, f, data, destroy)
}
func (c *Calendar) SetDetailHeightRows(rows int)  { CalendarSetDetailHeightRows(c, rows) }
func (c *Calendar) SetDetailWidthChars(chars int) { CalendarSetDetailWidthChars(c, chars) }
func (c *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	CalendarSetDisplayOptions(c, flags)
}
func (c *Calendar) Thaw()                   { CalendarThaw(c) }
func (c *Calendar) UnmarkDay(day uint) bool { return CalendarUnmarkDay(c, day) }

type Callback func(widget *Widget, data T.Gpointer)

type CallbackMarshal func(object *Object,
	data T.Gpointer, nArgs uint, args *Arg)

type Cell struct { // TODO(t):fix union
	Type       CellType
	Vertical   int16
	Horizontal int16
	Style      *Style
	/* union
	   text  *T.Gchar
	   struct {
			pixmap  *D.Pixmap
			mask  *T.GdkBitmap
	   }
	   struct {
			text  *T.Gchar
			spacing  uint8
			pixmap  *D.Pixmap
			mask  *GdkBitmap
	   }
	   widget  *Widget
	*/
}

type CellEditable struct{}

var CellEditableGetType func() O.Type

var (
	CellEditableEditingDone  func(c *CellEditable)
	CellEditableRemoveWidget func(c *CellEditable)
	CellEditableStartEditing func(c *CellEditable, event *D.Event)
)

func (c *CellEditable) EditingDone()                { CellEditableEditingDone(c) }
func (c *CellEditable) RemoveWidget()               { CellEditableRemoveWidget(c) }
func (c *CellEditable) StartEditing(event *D.Event) { CellEditableStartEditing(c, event) }

type CellLayout struct{}

type CellLayoutDataFunc func(
	CellLayout *CellLayout,
	Cell *CellRenderer,
	treeModel *TreeModel,
	iter *TreeIter,
	data T.Gpointer)

var (
	CellLayoutGetType func() O.Type

	CellLayoutAddAttribute    func(c *CellLayout, cell *CellRenderer, attribute string, column int)
	CellLayoutClear           func(c *CellLayout)
	CellLayoutClearAttributes func(c *CellLayout, cell *CellRenderer)
	CellLayoutGetCells        func(c *CellLayout) *T.GList
	CellLayoutPackEnd         func(c *CellLayout, cell *CellRenderer, expand bool)
	CellLayoutPackStart       func(c *CellLayout, cell *CellRenderer, expand bool)
	CellLayoutReorder         func(c *CellLayout, cell *CellRenderer, position int)
	CellLayoutSetAttributes   func(c *CellLayout, cell *CellRenderer, v ...VArg)
	CellLayoutSetCellDataFunc func(c *CellLayout, cell *CellRenderer, f CellLayoutDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
)

func (c *CellLayout) AddAttribute(cell *CellRenderer, attribute string, column int) {
	CellLayoutAddAttribute(c, cell, attribute, column)
}
func (c *CellLayout) Clear()                             { CellLayoutClear(c) }
func (c *CellLayout) ClearAttributes(cell *CellRenderer) { CellLayoutClearAttributes(c, cell) }
func (c *CellLayout) GetCells() *T.GList                 { return CellLayoutGetCells(c) }
func (c *CellLayout) PackEnd(cell *CellRenderer, expand bool) {
	CellLayoutPackEnd(c, cell, expand)
}
func (c *CellLayout) PackStart(cell *CellRenderer, expand bool) {
	CellLayoutPackStart(c, cell, expand)
}
func (c *CellLayout) Reorder(cell *CellRenderer, position int)    { CellLayoutReorder(c, cell, position) }
func (c *CellLayout) SetAttributes(cell *CellRenderer, v ...VArg) { CellLayoutSetAttributes(c, cell, v) }
func (c *CellLayout) SetCellDataFunc(cell *CellRenderer, f CellLayoutDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	CellLayoutSetCellDataFunc(c, cell, f, funcData, destroy)
}

type CellRenderer struct {
	Parent Object
	Xalign float32
	Yalign float32
	Width  int
	Height int
	Xpad   uint16
	Ypad   uint16
	Bits   uint
	// Mode : 2;
	// Visible : 1
	// IsExpander : 1
	// IsExpanded : 1
	// CellBackgroundSet : 1
	// Sensitive : 1
	// Editing : 1
}

type CellRendererState Enum

const (
	CELL_RENDERER_SELECTED CellRendererState = 1 << iota
	CELL_RENDERER_PRELIT
	CELL_RENDERER_INSENSITIVE
	CELL_RENDERER_SORTED
	CELL_RENDERER_FOCUSED
)

var (
	CellRendererGetType          func() O.Type
	CellRendererAccelGetType     func() O.Type
	CellRendererAccelNew         func() *CellRenderer
	CellRendererAccelModeGetType func() O.Type
	CellRendererComboGetType     func() O.Type
	CellRendererComboNew         func() *CellRenderer
	CellRendererModeGetType      func() O.Type
	CellRendererPixbufGetType    func() O.Type
	CellRendererPixbufNew        func() *CellRenderer
	CellRendererProgressGetType  func() O.Type
	CellRendererProgressNew      func() *CellRenderer
	CellRendererSpinGetType      func() O.Type
	CellRendererSpinNew          func() *CellRenderer
	CellRendererSpinnerGetType   func() O.Type
	CellRendererSpinnerNew       func() *CellRenderer
	CellRendererStateGetType     func() O.Type
	CellRendererToggleGetType    func() O.Type
	CellRendererToggleNew        func() *CellRenderer
)

var (
	CellRendererActivate        func(c *CellRenderer, event *D.Event, widget *Widget, path string, backgroundArea, cellArea *D.Rectangle, flags CellRendererState) bool
	CellRendererEditingCanceled func(c *CellRenderer)
	CellRendererGetAlignment    func(c *CellRenderer, xalign, yalign *float32)
	CellRendererGetFixedSize    func(c *CellRenderer, width, height *int)
	CellRendererGetPadding      func(c *CellRenderer, xpad, ypad *int)
	CellRendererGetSensitive    func(c *CellRenderer) bool
	CellRendererGetSize         func(c *CellRenderer, widget *Widget, cellArea *D.Rectangle, xOffset, yOffset, width, height *int)
	CellRendererGetVisible      func(c *CellRenderer) bool
	CellRendererRender          func(c *CellRenderer, window *D.Window, widget *Widget, backgroundArea, cellArea, exposeArea *D.Rectangle, flags CellRendererState)
	CellRendererSetAlignment    func(c *CellRenderer, xalign, yalign float32)
	CellRendererSetFixedSize    func(c *CellRenderer, width, height int)
	CellRendererSetPadding      func(c *CellRenderer, xpad, ypad int)
	CellRendererSetSensitive    func(c *CellRenderer, sensitive bool)
	CellRendererSetVisible      func(c *CellRenderer, visible bool)
	CellRendererStartEditing    func(c *CellRenderer, event *D.Event, widget *Widget, path string, backgroundArea, cellArea *D.Rectangle, flags CellRendererState) *CellEditable
	CellRendererStopEditing     func(c *CellRenderer, canceled bool)
)

func (c *CellRenderer) Activate(event *D.Event, widget *Widget, path string, backgroundArea, cellArea *D.Rectangle, flags CellRendererState) bool {
	return CellRendererActivate(c, event, widget, path, backgroundArea, cellArea, flags)
}
func (c *CellRenderer) EditingCanceled() { CellRendererEditingCanceled(c) }
func (c *CellRenderer) GetAlignment(xalign, yalign *float32) {
	CellRendererGetAlignment(c, xalign, yalign)
}
func (c *CellRenderer) GetFixedSize(width, height *int) { CellRendererGetFixedSize(c, width, height) }
func (c *CellRenderer) GetPadding(xpad, ypad *int)      { CellRendererGetPadding(c, xpad, ypad) }
func (c *CellRenderer) GetSensitive() bool              { return CellRendererGetSensitive(c) }
func (c *CellRenderer) GetSize(widget *Widget, cellArea *D.Rectangle, xOffset, yOffset, width, height *int) {
	CellRendererGetSize(c, widget, cellArea, xOffset, yOffset, width, height)
}
func (c *CellRenderer) GetVisible() bool { return CellRendererGetVisible(c) }
func (c *CellRenderer) Render(window *D.Window, widget *Widget, backgroundArea, cellArea, exposeArea *D.Rectangle, flags CellRendererState) {
	CellRendererRender(c, window, widget, backgroundArea, cellArea, exposeArea, flags)
}
func (c *CellRenderer) SetAlignment(xalign, yalign float32) {
	CellRendererSetAlignment(c, xalign, yalign)
}
func (c *CellRenderer) SetFixedSize(width, height int) { CellRendererSetFixedSize(c, width, height) }
func (c *CellRenderer) SetPadding(xpad, ypad int)      { CellRendererSetPadding(c, xpad, ypad) }
func (c *CellRenderer) SetSensitive(sensitive bool)    { CellRendererSetSensitive(c, sensitive) }
func (c *CellRenderer) SetVisible(visible bool)        { CellRendererSetVisible(c, visible) }
func (c *CellRenderer) StartEditing(event *D.Event, widget *Widget, path string, backgroundArea, cellArea *D.Rectangle, flags CellRendererState) *CellEditable {
	return CellRendererStartEditing(c, event, widget, path, backgroundArea, cellArea, flags)
}
func (c *CellRenderer) StopEditing(canceled bool) { CellRendererStopEditing(c, canceled) }

type CellRendererText struct {
	Parent          CellRenderer
	Text            *T.Gchar
	Font            *P.FontDescription
	FontScale       float64
	Foreground      P.Color
	Background      P.Color
	ExtraAttrs      *P.AttrList
	UnderlineStyle  P.Underline
	Rise            int
	FixedHeightRows int
	Bits            uint
	// Strikethrough : 1
	// Editable : 1
	// ScaleSet : 1
	// ForegroundSet : 1
	// BackgroundSet : 1
	// UnderlineSet : 1
	// RiseSet : 1
	// StrikethroughSet : 1
	// EditableSet : 1
	// CalcFixedHeight : 1
}

var (
	CellRendererTextGetType func() O.Type
	CellRendererTextNew     func() *CellRenderer

	CellRendererTextSetFixedHeightFromFont func(renderer *CellRendererText, numberOfRows int)
)

func (c *CellRendererText) SetFixedHeightFromFont(numberOfRows int) {
	CellRendererTextSetFixedHeightFromFont(c, numberOfRows)
}

type CellRendererToggle struct {
	Parent CellRenderer
	Bits   uint
	// Active : 1
	// Activatable : 1
	// Radio : 1
}

var (
	CellRendererToggleGetActivatable func(c *CellRendererToggle) bool
	CellRendererToggleGetActive      func(c *CellRendererToggle) bool
	CellRendererToggleGetRadio       func(c *CellRendererToggle) bool
	CellRendererToggleSetActivatable func(c *CellRendererToggle, setting bool)
	CellRendererToggleSetActive      func(c *CellRendererToggle, setting bool)
	CellRendererToggleSetRadio       func(c *CellRendererToggle, radio bool)
)

func (c *CellRendererToggle) GetActivatable() bool { return CellRendererToggleGetActivatable(c) }
func (c *CellRendererToggle) GetActive() bool      { return CellRendererToggleGetActive(c) }
func (c *CellRendererToggle) GetRadio() bool       { return CellRendererToggleGetRadio(c) }
func (c *CellRendererToggle) SetActivatable(setting bool) {
	CellRendererToggleSetActivatable(c, setting)
}
func (c *CellRendererToggle) SetActive(setting bool) { CellRendererToggleSetActive(c, setting) }
func (c *CellRendererToggle) SetRadio(radio bool)    { CellRendererToggleSetRadio(c, radio) }

type CellView struct {
	Parent Widget
	_      struct{}
}

type CellType Enum

const (
	CELL_EMPTY CellType = iota
	CELL_TEXT
	CELL_PIXMAP
	CELL_PIXTEXT
	CELL_WIDGET
)

var CellTypeGetType func() O.Type

var (
	CellViewGetType       func() O.Type
	CellViewNew           func() *Widget
	CellViewNewWithText   func(text string) *Widget
	CellViewNewWithMarkup func(markup string) *Widget
	CellViewNewWithPixbuf func(pixbuf *D.Pixbuf) *Widget
)

var (
	CellViewGetCellRenderers   func(c *CellView) *T.GList
	CellViewGetDisplayedRow    func(c *CellView) *TreePath
	CellViewGetModel           func(c *CellView) *TreeModel
	CellViewGetSizeOfRow       func(c *CellView, path *TreePath, requisition *Requisition) bool
	CellViewSetBackgroundColor func(c *CellView, color *D.Color)
	CellViewSetDisplayedRow    func(c *CellView, path *TreePath)
	CellViewSetModel           func(c *CellView, model *TreeModel)
)

func (c *CellView) GetCellRenderers() *T.GList { return CellViewGetCellRenderers(c) }
func (c *CellView) GetDisplayedRow() *TreePath { return CellViewGetDisplayedRow(c) }
func (c *CellView) GetModel() *TreeModel       { return CellViewGetModel(c) }
func (c *CellView) GetSizeOfRow(path *TreePath, requisition *Requisition) bool {
	return CellViewGetSizeOfRow(c, path, requisition)
}
func (c *CellView) SetBackgroundColor(color *D.Color) { CellViewSetBackgroundColor(c, color) }
func (c *CellView) SetDisplayedRow(path *TreePath)    { CellViewSetDisplayedRow(c, path) }
func (c *CellView) SetModel(model *TreeModel)         { CellViewSetModel(c, model) }

var (
	CheckButtonNew             func() *Widget
	CheckButtonNewWithLabel    func(label string) *Widget
	CheckButtonNewWithMnemonic func(label string) *Widget
)

type CheckButton struct {
	ToggleButton ToggleButton
}

var CheckButtonGetType func() O.Type

type CheckMenuItem struct {
	MenuItem MenuItem
	Bits     uint
	// Active : 1
	// AlwaysShowToggle : 1
	// Inconsistent : 1
	// DrawAsRadio : 1
}

var (
	CheckMenuItemGetType         func() O.Type
	CheckMenuItemNew             func() *Widget
	CheckMenuItemNewWithLabel    func(label string) *Widget
	CheckMenuItemNewWithMnemonic func(label string) *Widget
)

var (
	CheckMenuItemGetActive       func(c *CheckMenuItem) bool
	CheckMenuItemGetDrawAsRadio  func(c *CheckMenuItem) bool
	CheckMenuItemGetInconsistent func(c *CheckMenuItem) bool
	CheckMenuItemSetActive       func(c *CheckMenuItem, isActive bool)
	CheckMenuItemSetDrawAsRadio  func(c *CheckMenuItem, drawAsRadio bool)
	CheckMenuItemSetInconsistent func(c *CheckMenuItem, setting bool)
	CheckMenuItemSetShowToggle   func(c *CheckMenuItem, always bool)
	CheckMenuItemToggled         func(c *CheckMenuItem)
)

func (c *CheckMenuItem) GetActive() bool         { return CheckMenuItemGetActive(c) }
func (c *CheckMenuItem) GetDrawAsRadio() bool    { return CheckMenuItemGetDrawAsRadio(c) }
func (c *CheckMenuItem) GetInconsistent() bool   { return CheckMenuItemGetInconsistent(c) }
func (c *CheckMenuItem) SetActive(isActive bool) { CheckMenuItemSetActive(c, isActive) }
func (c *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	CheckMenuItemSetDrawAsRadio(c, drawAsRadio)
}
func (c *CheckMenuItem) SetInconsistent(setting bool) { CheckMenuItemSetInconsistent(c, setting) }
func (c *CheckMenuItem) SetShowToggle(always bool)    { CheckMenuItemSetShowToggle(c, always) }
func (c *CheckMenuItem) Toggled()                     { CheckMenuItemToggled(c) }

var CheckVersion func(
	requiredMajor, requiredMinor, requiredMicro uint) string

type ClassInitFunc T.GBaseInitFunc

type Clipboard struct{}

type (
	ClipboardClearFunc            func(clipboard *Clipboard, userDataOrOwner T.Gpointer)
	ClipboardGetFunc              func(clipboard *Clipboard, selectionData *SelectionData, info uint, userDataOrOwner T.Gpointer)
	ClipboardImageReceivedFunc    func(clipboard *Clipboard, pixbuf *D.Pixbuf, data T.Gpointer)
	ClipboardReceivedFunc         func(clipboard *Clipboard, selectionData *SelectionData, data T.Gpointer)
	ClipboardRichTextReceivedFunc func(clipboard *Clipboard, format D.Atom, text *uint8, length T.Gsize, data T.Gpointer)
	ClipboardTargetsReceivedFunc  func(clipboard *Clipboard, atoms *D.Atom, nAtoms int, data T.Gpointer)
	ClipboardTextReceivedFunc     func(clipboard *Clipboard, text string, data T.Gpointer)
	ClipboardURIReceivedFunc      func(clipboard *Clipboard, uris *string, data T.Gpointer)
)

var (
	ClipboardGetType       func() O.Type
	ClipboardGet           func(selection D.Atom) *Clipboard
	ClipboardGetForDisplay func(display *D.Display, selection D.Atom) *Clipboard
	WidgetGetClipboard     func(widget *Widget, selection D.Atom) *Clipboard
)

var (
	ClipboardClear                   func(c *Clipboard)
	ClipboardGetDisplay              func(c *Clipboard) *D.Display
	ClipboardGetOwner                func(c *Clipboard) *T.GObject
	ClipboardRequestContents         func(c *Clipboard, target D.Atom, callback ClipboardReceivedFunc, userData T.Gpointer)
	ClipboardRequestImage            func(c *Clipboard, callback ClipboardImageReceivedFunc, userData T.Gpointer)
	ClipboardRequestRichText         func(c *Clipboard, buffer *TextBuffer, callback ClipboardRichTextReceivedFunc, userData T.Gpointer)
	ClipboardRequestTargets          func(c *Clipboard, callback ClipboardTargetsReceivedFunc, userData T.Gpointer)
	ClipboardRequestText             func(c *Clipboard, callback ClipboardTextReceivedFunc, userData T.Gpointer)
	ClipboardRequestUris             func(c *Clipboard, callback ClipboardURIReceivedFunc, userData T.Gpointer)
	ClipboardSetCanStore             func(c *Clipboard, targets *TargetEntry, nTargets int)
	ClipboardSetImage                func(c *Clipboard, pixbuf *D.Pixbuf)
	ClipboardSetText                 func(c *Clipboard, text string, len int)
	ClipboardSetWithData             func(c *Clipboard, targets *TargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, userData T.Gpointer) bool
	ClipboardSetWithOwner            func(c *Clipboard, targets *TargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, owner *T.GObject) bool
	ClipboardStore                   func(c *Clipboard)
	ClipboardWaitForContents         func(c *Clipboard, target D.Atom) *SelectionData
	ClipboardWaitForImage            func(c *Clipboard) *D.Pixbuf
	ClipboardWaitForRichText         func(c *Clipboard, buffer *TextBuffer, format *D.Atom, length *T.Gsize) *uint8
	ClipboardWaitForTargets          func(c *Clipboard, targets **D.Atom, nTargets *int) bool
	ClipboardWaitForText             func(c *Clipboard) string
	ClipboardWaitForUris             func(c *Clipboard) []string
	ClipboardWaitIsImageAvailable    func(c *Clipboard) bool
	ClipboardWaitIsRichTextAvailable func(c *Clipboard, buffer *TextBuffer) bool
	ClipboardWaitIsTargetAvailable   func(c *Clipboard, target D.Atom) bool
	ClipboardWaitIsTextAvailable     func(c *Clipboard) bool
	ClipboardWaitIsUrisAvailable     func(c *Clipboard) bool
)

func (c *Clipboard) Clear()                 { ClipboardClear(c) }
func (c *Clipboard) GetDisplay() *D.Display { return ClipboardGetDisplay(c) }
func (c *Clipboard) GetOwner() *T.GObject   { return ClipboardGetOwner(c) }
func (c *Clipboard) RequestContents(target D.Atom, callback ClipboardReceivedFunc, userData T.Gpointer) {
	ClipboardRequestContents(c, target, callback, userData)
}
func (c *Clipboard) RequestImage(callback ClipboardImageReceivedFunc, userData T.Gpointer) {
	ClipboardRequestImage(c, callback, userData)
}
func (c *Clipboard) RequestRichText(buffer *TextBuffer, callback ClipboardRichTextReceivedFunc, userData T.Gpointer) {
	ClipboardRequestRichText(c, buffer, callback, userData)
}
func (c *Clipboard) RequestTargets(callback ClipboardTargetsReceivedFunc, userData T.Gpointer) {
	ClipboardRequestTargets(c, callback, userData)
}
func (c *Clipboard) RequestText(callback ClipboardTextReceivedFunc, userData T.Gpointer) {
	ClipboardRequestText(c, callback, userData)
}
func (c *Clipboard) RequestUris(callback ClipboardURIReceivedFunc, userData T.Gpointer) {
	ClipboardRequestUris(c, callback, userData)
}
func (c *Clipboard) SetCanStore(targets *TargetEntry, nTargets int) {
	ClipboardSetCanStore(c, targets, nTargets)
}
func (c *Clipboard) SetImage(pixbuf *D.Pixbuf)     { ClipboardSetImage(c, pixbuf) }
func (c *Clipboard) SetText(text string, leng int) { ClipboardSetText(c, text, leng) }
func (c *Clipboard) SetWithData(targets *TargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, userData T.Gpointer) bool {
	return ClipboardSetWithData(c, targets, nTargets, getFunc, clearFunc, userData)
}
func (c *Clipboard) SetWithOwner(targets *TargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, owner *T.GObject) bool {
	return ClipboardSetWithOwner(c, targets, nTargets, getFunc, clearFunc, owner)
}
func (c *Clipboard) Store() { ClipboardStore(c) }
func (c *Clipboard) WaitForContents(target D.Atom) *SelectionData {
	return ClipboardWaitForContents(c, target)
}
func (c *Clipboard) WaitForImage() *D.Pixbuf { return ClipboardWaitForImage(c) }
func (c *Clipboard) WaitForRichText(buffer *TextBuffer, format *D.Atom, length *T.Gsize) *uint8 {
	return ClipboardWaitForRichText(c, buffer, format, length)
}
func (c *Clipboard) WaitForTargets(targets **D.Atom, nTargets *int) bool {
	return ClipboardWaitForTargets(c, targets, nTargets)
}
func (c *Clipboard) WaitForText() string        { return ClipboardWaitForText(c) }
func (c *Clipboard) WaitForUris() []string      { return ClipboardWaitForUris(c) }
func (c *Clipboard) WaitIsImageAvailable() bool { return ClipboardWaitIsImageAvailable(c) }
func (c *Clipboard) WaitIsRichTextAvailable(buffer *TextBuffer) bool {
	return ClipboardWaitIsRichTextAvailable(c, buffer)
}
func (c *Clipboard) WaitIsTargetAvailable(target D.Atom) bool {
	return ClipboardWaitIsTargetAvailable(c, target)
}
func (c *Clipboard) WaitIsTextAvailable() bool { return ClipboardWaitIsTextAvailable(c) }
func (c *Clipboard) WaitIsUrisAvailable() bool { return ClipboardWaitIsUrisAvailable(c) }

type (
	CList struct {
		Container          Container
		Flags              uint16
		_                  T.Gpointer
		_                  T.Gpointer
		Freeze_count       uint
		InternalAllocation D.Rectangle
		Rows               int
		RowHeight          int
		RowList            *T.GList
		RowListEnd         *T.GList
		Columns            int
		ColumnTitleArea    D.Rectangle
		TitleWindow        *D.Window
		Column             *CListColumn
		ClistWindow        *D.Window
		ClistWindowWidth   int
		ClistWindowHeight  int
		Hoffset            int
		Voffset            int
		ShadowType         ShadowType
		SelectionMode      SelectionMode
		Selection          *T.GList
		SelectionWnd       *T.GList
		UndoSelectionList  *T.GList
		UndoUnselection    *T.GList
		UndoAnchor         int
		ButtonActions      [5]uint8
		DragButton         uint8
		ClickCell          CListCellInfo
		Hadjustment        *Adjustment
		Vadjustment        *Adjustment
		XorGc              *D.GC
		FgGc               *D.GC
		BgGc               *D.GC
		CursorDrag         *D.Cursor
		XDrag              int
		FocusRow           int
		FocusHeaderColumn  int
		Anchor             int
		AnchorState        StateType
		DragPos            int
		Htimer             int
		Vtimer             int
		SortType           SortType
		Compare            CListCompareFunc
		SortColumn         int
		DragHighlightRow   int
		DragHighlightPos   CListDragPos
	}

	CListColumn struct {
		Title         *T.Gchar
		Area          D.Rectangle
		Button        *Widget
		Window        *D.Window
		Width         int
		Min_width     int
		Max_width     int
		Justification Justification
		Bits          uint
		// Visible : 1
		// WidthSet : 1
		// Resizeable : 1
		// AutoResize : 1
		// ButtonPassive : 1
	}

	CListCellInfo struct {
		Row    int
		Column int
	}

	CListRow struct {
		Cell       *Cell
		State      StateType
		Foreground D.Color
		Background D.Color
		Style      *Style
		Data       T.Gpointer
		Destroy    T.GDestroyNotify
		Bits       uint
		// FgSet : 1
		// BgSet : 1
		// Selectable : 1
	}

	CListCompareFunc func(clist *CList,
		ptr1, ptr2 T.Gconstpointer) int
)

type CListDragPos Enum

const (
	CLIST_DRAG_NONE CListDragPos = iota
	CLIST_DRAG_BEFORE
	CLIST_DRAG_INTO
	CLIST_DRAG_AFTER
)

var (
	ClistGetType        func() O.Type
	ClistNew            func(columns int) *Widget
	ClistNewWithTitles  func(columns int, titles []string) *Widget
	ClistDragPosGetType func() O.Type
)

var (
	ClistAppend                 func(c *CList, text []string) int
	ClistClear                  func(c *CList)
	ClistColumnsAutosize        func(c *CList) int
	ClistColumnTitleActive      func(c *CList, column int)
	ClistColumnTitlePassive     func(c *CList, column int)
	ClistColumnTitlesActive     func(c *CList)
	ClistColumnTitlesHide       func(c *CList)
	ClistColumnTitlesPassive    func(c *CList)
	ClistColumnTitlesShow       func(c *CList)
	ClistFindRowFromData        func(c *CList, data T.Gpointer) int
	ClistFreeze                 func(c *CList)
	ClistGetCellStyle           func(c *CList, row, column int) *Style
	ClistGetCellType            func(c *CList, row, column int) CellType
	ClistGetColumnTitle         func(c *CList, column int) string
	ClistGetColumnWidget        func(c *CList, column int) *Widget
	ClistGetHadjustment         func(c *CList) *Adjustment
	ClistGetPixmap              func(c *CList, row, column int, pixmap, mask **T.GdkBitmap) int
	ClistGetPixtext             func(c *CList, row, column int, text []string, spacing *uint8, pixmap, mask **T.GdkBitmap) int
	ClistGetRowData             func(c *CList, row int) T.Gpointer
	ClistGetRowStyle            func(c *CList, row int) *Style
	ClistGetSelectable          func(c *CList, row int) bool
	ClistGetSelectionInfo       func(c *CList, x, y int, row, column *int) int
	ClistGetText                func(c *CList, row, column int, text []string) int
	ClistGetVadjustment         func(c *CList) *Adjustment
	ClistInsert                 func(c *CList, row int, text []string) int
	ClistMoveto                 func(c *CList, row, column int, rowAlign, colAlign float32)
	ClistOptimalColumnWidth     func(c *CList, column int) int
	ClistPrepend                func(c *CList, text []string) int
	ClistRemove                 func(c *CList, row int)
	ClistRowIsVisible           func(c *CList, row int) Visibility
	ClistRowMove                func(c *CList, sourceRow, destRow int)
	ClistSelectAll              func(c *CList)
	ClistSelectRow              func(c *CList, row, column int)
	ClistSetAutoSort            func(c *CList, autoSort bool)
	ClistSetBackground          func(c *CList, row int, color *D.Color)
	ClistSetButtonActions       func(c *CList, button uint, buttonActions uint8)
	ClistSetCellStyle           func(c *CList, row, column int, style *Style)
	ClistSetColumnAutoResize    func(c *CList, column int, autoResize bool)
	ClistSetColumnJustification func(c *CList, column int, justification Justification)
	ClistSetColumnMaxWidth      func(c *CList, column, maxWidth int)
	ClistSetColumnMinWidth      func(c *CList, column, minWidth int)
	ClistSetColumnResizeable    func(c *CList, column int, resizeable bool)
	ClistSetColumnTitle         func(c *CList, column int, title string)
	ClistSetColumnVisibility    func(c *CList, column int, visible bool)
	ClistSetColumnWidget        func(c *CList, column int, widget *Widget)
	ClistSetColumnWidth         func(c *CList, column, width int)
	ClistSetCompareFunc         func(c *CList, cmpFunc CListCompareFunc)
	ClistSetForeground          func(c *CList, row int, color *D.Color)
	ClistSetHadjustment         func(c *CList, adjustment *Adjustment)
	ClistSetPixmap              func(c *CList, row, column int, pixmap *D.Pixmap, mask *T.GdkBitmap)
	ClistSetPixtext             func(c *CList, row, column int, text string, spacing uint8, pixmap, mask *T.GdkBitmap)
	ClistSetReorderable         func(c *CList, reorderable bool)
	ClistSetRowData             func(c *CList, row int, data T.Gpointer)
	ClistSetRowDataFull         func(c *CList, row int, data T.Gpointer, destroy T.GDestroyNotify)
	ClistSetRowHeight           func(c *CList, height uint)
	ClistSetRowStyle            func(c *CList, row, style *Style)
	ClistSetSelectable          func(c *CList, row int, selectable bool)
	ClistSetSelectionMode       func(c *CList, mode SelectionMode)
	ClistSetShadowType          func(c *CList, t ShadowType)
	ClistSetShift               func(c *CList, row, column, vertical, horizontal int)
	ClistSetSortColumn          func(c *CList, column int)
	ClistSetSortType            func(c *CList, sortType SortType)
	ClistSetText                func(c *CList, row, column int, text string)
	ClistSetUseDragIcons        func(c *CList, useIcons bool)
	ClistSetVadjustment         func(c *CList, adjustment *Adjustment)
	ClistSort                   func(c *CList)
	ClistSwapRows               func(c *CList, row1, row2 int)
	ClistThaw                   func(c *CList)
	ClistUndoSelection          func(c *CList)
	ClistUnselectAll            func(c *CList)
	ClistUnselectRow            func(c *CList, row, column int)
)

func (c *CList) Append(text []string) int             { return ClistAppend(c, text) }
func (c *CList) Clear()                               { ClistClear(c) }
func (c *CList) ColumnsAutosize() int                 { return ClistColumnsAutosize(c) }
func (c *CList) ColumnTitleActive(column int)         { ClistColumnTitleActive(c, column) }
func (c *CList) ColumnTitlePassive(column int)        { ClistColumnTitlePassive(c, column) }
func (c *CList) ColumnTitlesActive()                  { ClistColumnTitlesActive(c) }
func (c *CList) ColumnTitlesHide()                    { ClistColumnTitlesHide(c) }
func (c *CList) ColumnTitlesPassive()                 { ClistColumnTitlesPassive(c) }
func (c *CList) ColumnTitlesShow()                    { ClistColumnTitlesShow(c) }
func (c *CList) FindRowFromData(data T.Gpointer) int  { return ClistFindRowFromData(c, data) }
func (c *CList) Freeze()                              { ClistFreeze(c) }
func (c *CList) GetCellStyle(row, column int) *Style  { return ClistGetCellStyle(c, row, column) }
func (c *CList) GetCellType(row, column int) CellType { return ClistGetCellType(c, row, column) }
func (c *CList) GetColumnTitle(column int) string     { return ClistGetColumnTitle(c, column) }
func (c *CList) GetColumnWidget(column int) *Widget   { return ClistGetColumnWidget(c, column) }
func (c *CList) GetHadjustment() *Adjustment          { return ClistGetHadjustment(c) }
func (c *CList) GetPixmap(row, column int, pixmap, mask **T.GdkBitmap) int {
	return ClistGetPixmap(c, row, column, pixmap, mask)
}
func (c *CList) GetPixtext(row, column int, text []string, spacing *uint8, pixmap, mask **T.GdkBitmap) int {
	return ClistGetPixtext(c, row, column, text, spacing, pixmap, mask)
}
func (c *CList) GetRowData(row int) T.Gpointer { return ClistGetRowData(c, row) }
func (c *CList) GetRowStyle(row int) *Style    { return ClistGetRowStyle(c, row) }
func (c *CList) GetSelectable(row int) bool    { return ClistGetSelectable(c, row) }
func (c *CList) GetSelectionInfo(x, y int, row, column *int) int {
	return ClistGetSelectionInfo(c, x, y, row, column)
}
func (c *CList) GetText(row, column int, text []string) int {
	return ClistGetText(c, row, column, text)
}
func (c *CList) GetVadjustment() *Adjustment       { return ClistGetVadjustment(c) }
func (c *CList) Insert(row int, text []string) int { return ClistInsert(c, row, text) }
func (c *CList) Moveto(row, column int, rowAlign, colAlign float32) {
	ClistMoveto(c, row, column, rowAlign, colAlign)
}
func (c *CList) OptimalColumnWidth(column int) int     { return ClistOptimalColumnWidth(c, column) }
func (c *CList) Prepend(text []string) int             { return ClistPrepend(c, text) }
func (c *CList) Remove(row int)                        { ClistRemove(c, row) }
func (c *CList) RowIsVisible(row int) Visibility       { return ClistRowIsVisible(c, row) }
func (c *CList) RowMove(sourceRow, destRow int)        { ClistRowMove(c, sourceRow, destRow) }
func (c *CList) SelectAll()                            { ClistSelectAll(c) }
func (c *CList) SelectRow(row, column int)             { ClistSelectRow(c, row, column) }
func (c *CList) SetAutoSort(autoSort bool)             { ClistSetAutoSort(c, autoSort) }
func (c *CList) SetBackground(row int, color *D.Color) { ClistSetBackground(c, row, color) }
func (c *CList) SetButtonActions(button uint, buttonActions uint8) {
	ClistSetButtonActions(c, button, buttonActions)
}
func (c *CList) SetCellStyle(row, column int, style *Style) {
	ClistSetCellStyle(c, row, column, style)
}
func (c *CList) SetColumnAutoResize(column int, autoResize bool) {
	ClistSetColumnAutoResize(c, column, autoResize)
}
func (c *CList) SetColumnJustification(column int, justification Justification) {
	ClistSetColumnJustification(c, column, justification)
}
func (c *CList) SetColumnMaxWidth(column, maxWidth int) { ClistSetColumnMaxWidth(c, column, maxWidth) }
func (c *CList) SetColumnMinWidth(column, minWidth int) { ClistSetColumnMinWidth(c, column, minWidth) }
func (c *CList) SetColumnResizeable(column int, resizeable bool) {
	ClistSetColumnResizeable(c, column, resizeable)
}
func (c *CList) SetColumnTitle(column int, title string) { ClistSetColumnTitle(c, column, title) }
func (c *CList) SetColumnVisibility(column int, visible bool) {
	ClistSetColumnVisibility(c, column, visible)
}
func (c *CList) SetColumnWidget(column int, widget *Widget) { ClistSetColumnWidget(c, column, widget) }
func (c *CList) SetColumnWidth(column, width int)           { ClistSetColumnWidth(c, column, width) }
func (c *CList) SetCompareFunc(cmpFunc CListCompareFunc)    { ClistSetCompareFunc(c, cmpFunc) }
func (c *CList) SetForeground(row int, color *D.Color)      { ClistSetForeground(c, row, color) }
func (c *CList) SetHadjustment(adjustment *Adjustment)      { ClistSetHadjustment(c, adjustment) }
func (c *CList) SetPixmap(row, column int, pixmap *D.Pixmap, mask *T.GdkBitmap) {
	ClistSetPixmap(c, row, column, pixmap, mask)
}
func (c *CList) SetPixtext(row, column int, text string, spacing uint8, pixmap, mask *T.GdkBitmap) {
	ClistSetPixtext(c, row, column, text, spacing, pixmap, mask)
}
func (c *CList) SetReorderable(reorderable bool)     { ClistSetReorderable(c, reorderable) }
func (c *CList) SetRowData(row int, data T.Gpointer) { ClistSetRowData(c, row, data) }
func (c *CList) SetRowDataFull(row int, data T.Gpointer, destroy T.GDestroyNotify) {
	ClistSetRowDataFull(c, row, data, destroy)
}
func (c *CList) SetRowHeight(height uint)               { ClistSetRowHeight(c, height) }
func (c *CList) SetRowStyle(row, style *Style)          { ClistSetRowStyle(c, row, style) }
func (c *CList) SetSelectable(row int, selectable bool) { ClistSetSelectable(c, row, selectable) }
func (c *CList) SetSelectionMode(mode SelectionMode)    { ClistSetSelectionMode(c, mode) }
func (c *CList) SetShadowType(t ShadowType)             { ClistSetShadowType(c, t) }
func (c *CList) SetShift(row, column, vertical, horizontal int) {
	ClistSetShift(c, row, column, vertical, horizontal)
}
func (c *CList) SetSortColumn(column int)              { ClistSetSortColumn(c, column) }
func (c *CList) SetSortType(sortType SortType)         { ClistSetSortType(c, sortType) }
func (c *CList) SetText(row, column int, text string)  { ClistSetText(c, row, column, text) }
func (c *CList) SetUseDragIcons(useIcons bool)         { ClistSetUseDragIcons(c, useIcons) }
func (c *CList) SetVadjustment(adjustment *Adjustment) { ClistSetVadjustment(c, adjustment) }
func (c *CList) Sort()                                 { ClistSort(c) }
func (c *CList) SwapRows(row1, row2 int)               { ClistSwapRows(c, row1, row2) }
func (c *CList) Thaw()                                 { ClistThaw(c) }
func (c *CList) UndoSelection()                        { ClistUndoSelection(c) }
func (c *CList) UnselectAll()                          { ClistUnselectAll(c) }
func (c *CList) UnselectRow(row, column int)           { ClistUnselectRow(c, row, column) }

type ColorButton struct {
	Button Button
	_      *struct{}
}

var (
	ColorButtonGetType      func() O.Type
	ColorButtonNew          func() *Widget
	ColorButtonNewWithColor func(color *D.Color) *Widget
)

var (
	ColorButtonGetAlpha    func(c *ColorButton) uint16
	ColorButtonGetColor    func(c *ColorButton, color *D.Color)
	ColorButtonGetTitle    func(c *ColorButton) string
	ColorButtonGetUseAlpha func(c *ColorButton) bool
	ColorButtonSetAlpha    func(c *ColorButton, alpha uint16)
	ColorButtonSetColor    func(c *ColorButton, color *D.Color)
	ColorButtonSetTitle    func(c *ColorButton, title string)
	ColorButtonSetUseAlpha func(c *ColorButton, useAlpha bool)
)

func (c *ColorButton) GetAlpha() uint16          { return ColorButtonGetAlpha(c) }
func (c *ColorButton) GetColor(color *D.Color)   { ColorButtonGetColor(c, color) }
func (c *ColorButton) GetTitle() string          { return ColorButtonGetTitle(c) }
func (c *ColorButton) GetUseAlpha() bool         { return ColorButtonGetUseAlpha(c) }
func (c *ColorButton) SetAlpha(alpha uint16)     { ColorButtonSetAlpha(c, alpha) }
func (c *ColorButton) SetColor(color *D.Color)   { ColorButtonSetColor(c, color) }
func (c *ColorButton) SetTitle(title string)     { ColorButtonSetTitle(c, title) }
func (c *ColorButton) SetUseAlpha(useAlpha bool) { ColorButtonSetUseAlpha(c, useAlpha) }

type (
	ColorSelection struct {
		Parent VBox
		_      T.Gpointer
	}

	ColorSelectionChangePaletteFunc func(
		colors *D.Color, nColors int)

	ColorSelectionChangePaletteWithScreenFunc func(
		screen *D.Screen, colors *D.Color, nColors int)
)

var (
	ColorSelectionGetType func() O.Type
	ColorSelectionNew     func() *Widget

	ColorSelectionPaletteFromString func(str string, colors **D.Color, nColors *int) bool
	ColorSelectionPaletteToString   func(colors *D.Color, nColors int) string

	ColorSelectionSetChangePaletteHook           func(f ColorSelectionChangePaletteFunc) ColorSelectionChangePaletteFunc
	ColorSelectionSetChangePaletteWithScreenHook func(f ColorSelectionChangePaletteWithScreenFunc) ColorSelectionChangePaletteWithScreenFunc
)

var (
	ColorSelectionGetColor             func(c *ColorSelection, color *float64)
	ColorSelectionGetCurrentAlpha      func(c *ColorSelection) uint16
	ColorSelectionGetCurrentColor      func(c *ColorSelection, color *D.Color)
	ColorSelectionGetHasOpacityControl func(c *ColorSelection) bool
	ColorSelectionGetHasPalette        func(c *ColorSelection) bool
	ColorSelectionGetPreviousAlpha     func(c *ColorSelection) uint16
	ColorSelectionGetPreviousColor     func(c *ColorSelection, color *D.Color)
	ColorSelectionIsAdjusting          func(c *ColorSelection) bool
	ColorSelectionSetColor             func(c *ColorSelection, color *float64)
	ColorSelectionSetCurrentAlpha      func(c *ColorSelection, alpha uint16)
	ColorSelectionSetCurrentColor      func(c *ColorSelection, color *D.Color)
	ColorSelectionSetHasOpacityControl func(c *ColorSelection, hasOpacity bool)
	ColorSelectionSetHasPalette        func(c *ColorSelection, hasPalette bool)
	ColorSelectionSetPreviousAlpha     func(c *ColorSelection, alpha uint16)
	ColorSelectionSetPreviousColor     func(c *ColorSelection, color *D.Color)
	ColorSelectionSetUpdatePolicy      func(c *ColorSelection, policy UpdateType)
)

func (c *ColorSelection) GetColor(color *float64)        { ColorSelectionGetColor(c, color) }
func (c *ColorSelection) GetCurrentAlpha() uint16        { return ColorSelectionGetCurrentAlpha(c) }
func (c *ColorSelection) GetCurrentColor(color *D.Color) { ColorSelectionGetCurrentColor(c, color) }
func (c *ColorSelection) GetHasOpacityControl() bool {
	return ColorSelectionGetHasOpacityControl(c)
}
func (c *ColorSelection) GetHasPalette() bool             { return ColorSelectionGetHasPalette(c) }
func (c *ColorSelection) GetPreviousAlpha() uint16        { return ColorSelectionGetPreviousAlpha(c) }
func (c *ColorSelection) GetPreviousColor(color *D.Color) { ColorSelectionGetPreviousColor(c, color) }
func (c *ColorSelection) IsAdjusting() bool               { return ColorSelectionIsAdjusting(c) }
func (c *ColorSelection) SetColor(color *float64)         { ColorSelectionSetColor(c, color) }
func (c *ColorSelection) SetCurrentAlpha(alpha uint16)    { ColorSelectionSetCurrentAlpha(c, alpha) }
func (c *ColorSelection) SetCurrentColor(color *D.Color)  { ColorSelectionSetCurrentColor(c, color) }
func (c *ColorSelection) SetHasOpacityControl(hasOpacity bool) {
	ColorSelectionSetHasOpacityControl(c, hasOpacity)
}
func (c *ColorSelection) SetHasPalette(hasPalette bool) {
	ColorSelectionSetHasPalette(c, hasPalette)
}
func (c *ColorSelection) SetPreviousAlpha(alpha uint16)   { ColorSelectionSetPreviousAlpha(c, alpha) }
func (c *ColorSelection) SetPreviousColor(color *D.Color) { ColorSelectionSetPreviousColor(c, color) }
func (c *ColorSelection) SetUpdatePolicy(policy UpdateType) {
	ColorSelectionSetUpdatePolicy(c, policy)
}

type ColorSelectionDialog struct {
	Parent       Dialog
	Colorsel     *Widget
	OkButton     *Widget
	CancelButton *Widget
	HelpButton   *Widget
}

var (
	ColorSelectionDialogGetType func() O.Type
	ColorSelectionDialogNew     func(title string) *Widget

	ColorSelectionDialogGetColorSelection func(c *ColorSelectionDialog) *Widget
)

type Combo struct {
	Hbox          HBox
	Entry         *Widget
	Button        *Widget
	Popup         *Widget
	Popwin        *Widget
	List          *Widget
	EntryChangeId uint
	ListChangeId  uint
	Bits          uint
	// ValueInList:1
	// OkIfEmpty:1
	// CaseSensitive:1
	// UseArrows:1
	// UseArrowsAlways:1
	CurrentButton uint16
	ActivateId    uint
}

var (
	ComboGetType func() O.Type
	ComboNew     func() *Widget
)

var (
	ComboSetValueInList     func(c *Combo, val bool, okIfEmpty bool)
	ComboSetUseArrows       func(c *Combo, val bool)
	ComboSetUseArrowsAlways func(c *Combo, val bool)
	ComboSetCaseSensitive   func(c *Combo, val bool)
	ComboSetItemString      func(c *Combo, item *Item, itemValue string)
	ComboSetPopdownStrings  func(c *Combo, strings *T.GList)
	ComboDisableActivate    func(c *Combo)
)

func (c *Combo) DisableActivate()                           { ComboDisableActivate(c) }
func (c *Combo) SetCaseSensitive(val bool)                  { ComboSetCaseSensitive(c, val) }
func (c *Combo) SetItemString(item *Item, itemValue string) { ComboSetItemString(c, item, itemValue) }
func (c *Combo) SetPopdownStrings(strings *T.GList)         { ComboSetPopdownStrings(c, strings) }
func (c *Combo) SetUseArrows(val bool)                      { ComboSetUseArrows(c, val) }
func (c *Combo) SetUseArrowsAlways(val bool)                { ComboSetUseArrowsAlways(c, val) }
func (c *Combo) SetValueInList(val bool, okIfEmpty bool) {
	ComboSetValueInList(c, val, okIfEmpty)
}

type ComboBox struct {
	ParentInstance Bin
	_              *struct{}
}

type ComboBoxText struct {
	ParentInstance *ComboBox
	_              *struct{}
}

type ComboBoxEntry struct {
	ParentInstance ComboBox
	_              *struct{}
}

var (
	ComboBoxGetType              func() O.Type
	ComboBoxNew                  func() *Widget
	ComboBoxNewWithEntry         func() *Widget
	ComboBoxNewWithModel         func(model *TreeModel) *Widget
	ComboBoxNewWithModelAndEntry func(model *TreeModel) *Widget

	ComboBoxTextGetType      func() O.Type
	ComboBoxTextNew          func() *Widget
	ComboBoxTextNewWithEntry func() *Widget

	ComboBoxNewText func() *Widget

	ComboBoxEntryGetType      func() O.Type
	ComboBoxEntryNew          func() *Widget
	ComboBoxEntryNewWithModel func(model *TreeModel, textColumn int) *Widget
	ComboBoxEntryNewText      func() *Widget
)

var (
	ComboBoxAppendText           func(c *ComboBox, text string)
	ComboBoxGetActive            func(c *ComboBox) int
	ComboBoxGetActiveIter        func(c *ComboBox, iter *TreeIter) bool
	ComboBoxGetActiveText        func(c *ComboBox) string
	ComboBoxGetAddTearoffs       func(c *ComboBox) bool
	ComboBoxGetButtonSensitivity func(c *ComboBox) SensitivityType
	ComboBoxGetColumnSpanColumn  func(c *ComboBox) int
	ComboBoxGetEntryTextColumn   func(c *ComboBox) int
	ComboBoxGetFocusOnClick      func(c *ComboBox) bool
	ComboBoxGetHasEntry          func(c *ComboBox) bool
	ComboBoxGetModel             func(c *ComboBox) *TreeModel
	ComboBoxGetPopupAccessible   func(c *ComboBox) *A.Object
	ComboBoxGetRowSeparatorFunc  func(c *ComboBox) TreeViewRowSeparatorFunc
	ComboBoxGetRowSpanColumn     func(c *ComboBox) int
	ComboBoxGetTitle             func(c *ComboBox) string
	ComboBoxGetWrapWidth         func(c *ComboBox) int
	ComboBoxInsertText           func(c *ComboBox, position int, text string)
	ComboBoxPopdown              func(c *ComboBox)
	ComboBoxPopup                func(c *ComboBox)
	ComboBoxPrependText          func(c *ComboBox, text string)
	ComboBoxRemoveText           func(c *ComboBox, position int)
	ComboBoxSetActive            func(c *ComboBox, index int)
	ComboBoxSetActiveIter        func(c *ComboBox, iter *TreeIter)
	ComboBoxSetAddTearoffs       func(c *ComboBox, addTearoffs bool)
	ComboBoxSetButtonSensitivity func(c *ComboBox, sensitivity SensitivityType)
	ComboBoxSetColumnSpanColumn  func(c *ComboBox, columnSpan int)
	ComboBoxSetEntryTextColumn   func(c *ComboBox, textColumn int)
	ComboBoxSetFocusOnClick      func(c *ComboBox, focusOnClick bool)
	ComboBoxSetModel             func(c *ComboBox, model *TreeModel)
	ComboBoxSetRowSeparatorFunc  func(c *ComboBox, f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify)
	ComboBoxSetRowSpanColumn     func(c *ComboBox, rowSpan int)
	ComboBoxSetTitle             func(c *ComboBox, title string)
	ComboBoxSetWrapWidth         func(c *ComboBox, width int)
)

func (c *ComboBox) AppendText(text string) { ComboBoxAppendText(c, text) }
func (c *ComboBox) GetActive() int         { return ComboBoxGetActive(c) }
func (c *ComboBox) GetActiveIter(iter *TreeIter) bool {
	return ComboBoxGetActiveIter(c, iter)
}
func (c *ComboBox) GetActiveText() string                 { return ComboBoxGetActiveText(c) }
func (c *ComboBox) GetAddTearoffs() bool                  { return ComboBoxGetAddTearoffs(c) }
func (c *ComboBox) GetButtonSensitivity() SensitivityType { return ComboBoxGetButtonSensitivity(c) }
func (c *ComboBox) GetColumnSpanColumn() int              { return ComboBoxGetColumnSpanColumn(c) }
func (c *ComboBox) GetEntryTextColumn() int               { return ComboBoxGetEntryTextColumn(c) }
func (c *ComboBox) GetFocusOnClick() bool                 { return ComboBoxGetFocusOnClick(c) }
func (c *ComboBox) GetHasEntry() bool                     { return ComboBoxGetHasEntry(c) }
func (c *ComboBox) GetModel() *TreeModel                  { return ComboBoxGetModel(c) }
func (c *ComboBox) GetPopupAccessible() *A.Object         { return ComboBoxGetPopupAccessible(c) }
func (c *ComboBox) GetRowSeparatorFunc() TreeViewRowSeparatorFunc {
	return ComboBoxGetRowSeparatorFunc(c)
}
func (c *ComboBox) GetRowSpanColumn() int                { return ComboBoxGetRowSpanColumn(c) }
func (c *ComboBox) GetTitle() string                     { return ComboBoxGetTitle(c) }
func (c *ComboBox) GetWrapWidth() int                    { return ComboBoxGetWrapWidth(c) }
func (c *ComboBox) InsertText(position int, text string) { ComboBoxInsertText(c, position, text) }
func (c *ComboBox) Popdown()                             { ComboBoxPopdown(c) }
func (c *ComboBox) Popup()                               { ComboBoxPopup(c) }
func (c *ComboBox) PrependText(text string)              { ComboBoxPrependText(c, text) }
func (c *ComboBox) RemoveText(position int)              { ComboBoxRemoveText(c, position) }
func (c *ComboBox) SetActive(index int)                  { ComboBoxSetActive(c, index) }
func (c *ComboBox) SetActiveIter(iter *TreeIter)         { ComboBoxSetActiveIter(c, iter) }
func (c *ComboBox) SetAddTearoffs(addTearoffs bool)      { ComboBoxSetAddTearoffs(c, addTearoffs) }
func (c *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	ComboBoxSetButtonSensitivity(c, sensitivity)
}
func (c *ComboBox) SetColumnSpanColumn(columnSpan int) { ComboBoxSetColumnSpanColumn(c, columnSpan) }
func (c *ComboBox) SetEntryTextColumn(textColumn int)  { ComboBoxSetEntryTextColumn(c, textColumn) }
func (c *ComboBox) SetFocusOnClick(focusOnClick bool)  { ComboBoxSetFocusOnClick(c, focusOnClick) }
func (c *ComboBox) SetModel(model *TreeModel)          { ComboBoxSetModel(c, model) }
func (c *ComboBox) SetRowSeparatorFunc(f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	ComboBoxSetRowSeparatorFunc(c, f, data, destroy)
}
func (c *ComboBox) SetRowSpanColumn(rowSpan int) { ComboBoxSetRowSpanColumn(c, rowSpan) }
func (c *ComboBox) SetTitle(title string)        { ComboBoxSetTitle(c, title) }
func (c *ComboBox) SetWrapWidth(width int)       { ComboBoxSetWrapWidth(c, width) }

var (
	ComboBoxTextAppendText    func(c *ComboBoxText, text string)
	ComboBoxTextGetActiveText func(c *ComboBoxText) string
	ComboBoxTextInsertText    func(c *ComboBoxText, position int, text string)
	ComboBoxTextPrependText   func(c *ComboBoxText, text string)
	ComboBoxTextRemove        func(c *ComboBoxText, position int)
)

func (c *ComboBoxText) AppendText(text string) { ComboBoxTextAppendText(c, text) }
func (c *ComboBoxText) GetActiveText() string  { return ComboBoxTextGetActiveText(c) }
func (c *ComboBoxText) InsertText(position int, text string) {
	ComboBoxTextInsertText(c, position, text)
}
func (c *ComboBoxText) PrependText(text string) { ComboBoxTextPrependText(c, text) }
func (c *ComboBoxText) Remove(position int)     { ComboBoxTextRemove(c, position) }

var (
	ComboBoxEntrySetTextColumn func(e *ComboBoxEntry, textColumn int)
	ComboBoxEntryGetTextColumn func(e *ComboBoxEntry) int
)

func (e *ComboBoxEntry) GetTextColumn() int           { return ComboBoxEntryGetTextColumn(e) }
func (e *ComboBoxEntry) SetTextColumn(textColumn int) { ComboBoxEntrySetTextColumn(e, textColumn) }

type Container struct {
	Widget     Widget
	FocusChild *Widget
	Bits       uint
	// BorderWidth : 16
	// NeedResize : 1
	// ResizeMode : 2
	// ReallocateRedraws : 1
	// HasFocusChain : 1
}

var (
	ContainerGetType func() O.Type

	ContainerAdd                  func(c *Container, widget *Widget)
	ContainerAddWithProperties    func(c *Container, widget *Widget, firstPropName string, v ...VArg)
	ContainerCheckResize          func(c *Container)
	ContainerChildGet             func(c *Container, child *Widget, firstPropName string, v ...VArg)
	ContainerChildGetProperty     func(c *Container, child *Widget, propertyName string, value *T.GValue)
	ContainerChildGetValist       func(c *Container, child *Widget, firstPropertyName string, varArgs T.VaList)
	ContainerChildSet             func(c *Container, child *Widget, firstPropName string, v ...VArg)
	ContainerChildSetProperty     func(c *Container, child *Widget, propertyName string, value *T.GValue)
	ContainerChildSetValist       func(c *Container, child *Widget, firstPropertyName string, varArgs T.VaList)
	ContainerChildType            func(c *Container) O.Type
	ContainerForall               func(c *Container, callback Callback, callbackData T.Gpointer)
	ContainerForeach              func(c *Container, callback Callback, callbackData T.Gpointer)
	ContainerForeachFull          func(c *Container, callback Callback, marshal CallbackMarshal, callbackData T.Gpointer, notify T.GDestroyNotify)
	ContainerGetBorderWidth       func(c *Container) uint
	ContainerGetChildren          func(c *Container) *T.GList
	ContainerGetFocusChain        func(c *Container, focusableWidgets **T.GList) bool
	ContainerGetFocusChild        func(c *Container) *Widget
	ContainerGetFocusHadjustment  func(c *Container) *Adjustment
	ContainerGetFocusVadjustment  func(c *Container) *Adjustment
	ContainerGetResizeMode        func(c *Container) ResizeMode
	ContainerPropagateExpose      func(c *Container, child *Widget, event *D.EventExpose)
	ContainerRemove               func(c *Container, widget *Widget)
	ContainerResizeChildren       func(c *Container)
	ContainerSetBorderWidth       func(c *Container, borderWidth uint)
	ContainerSetFocusChain        func(c *Container, focusableWidgets *T.GList)
	ContainerSetFocusChild        func(c *Container, child *Widget)
	ContainerSetFocusHadjustment  func(c *Container, adjustment *Adjustment)
	ContainerSetFocusVadjustment  func(c *Container, adjustment *Adjustment)
	ContainerSetReallocateRedraws func(c *Container, needsRedraws bool)
	ContainerSetResizeMode        func(c *Container, resizeMode ResizeMode)
	ContainerUnsetFocusChain      func(c *Container)
)

func (c *Container) Add(widget *Widget) { ContainerAdd(c, widget) }
func (c *Container) AddWithProperties(widget *Widget, firstPropName string, v ...VArg) {
	ContainerAddWithProperties(c, widget, firstPropName, v)
}
func (c *Container) CheckResize() { ContainerCheckResize(c) }
func (c *Container) ChildGet(child *Widget, firstPropName string, v ...VArg) {
	ContainerChildGet(c, child, firstPropName, v)
}
func (c *Container) ChildGetProperty(child *Widget, propertyName string, value *T.GValue) {
	ContainerChildGetProperty(c, child, propertyName, value)
}
func (c *Container) ChildGetValist(child *Widget, firstPropertyName string, varArgs T.VaList) {
	ContainerChildGetValist(c, child, firstPropertyName, varArgs)
}
func (c *Container) ChildSet(child *Widget, firstPropName string, v ...VArg) {
	ContainerChildSet(c, child, firstPropName, v)
}
func (c *Container) ChildSetProperty(child *Widget, propertyName string, value *T.GValue) {
	ContainerChildSetProperty(c, child, propertyName, value)
}
func (c *Container) ChildSetValist(child *Widget, firstPropertyName string, varArgs T.VaList) {
	ContainerChildSetValist(c, child, firstPropertyName, varArgs)
}
func (c *Container) ChildType() O.Type { return ContainerChildType(c) }
func (c *Container) Forall(callback Callback, callbackData T.Gpointer) {
	ContainerForall(c, callback, callbackData)
}
func (c *Container) Foreach(callback Callback, callbackData T.Gpointer) {
	ContainerForeach(c, callback, callbackData)
}
func (c *Container) ForeachFull(callback Callback, marshal CallbackMarshal, callbackData T.Gpointer, notify T.GDestroyNotify) {
	ContainerForeachFull(c, callback, marshal, callbackData, notify)
}
func (c *Container) GetBorderWidth() uint  { return ContainerGetBorderWidth(c) }
func (c *Container) GetChildren() *T.GList { return ContainerGetChildren(c) }
func (c *Container) GetFocusChain(focusableWidgets **T.GList) bool {
	return ContainerGetFocusChain(c, focusableWidgets)
}
func (c *Container) GetFocusChild() *Widget           { return ContainerGetFocusChild(c) }
func (c *Container) GetFocusHadjustment() *Adjustment { return ContainerGetFocusHadjustment(c) }
func (c *Container) GetFocusVadjustment() *Adjustment { return ContainerGetFocusVadjustment(c) }
func (c *Container) GetResizeMode() ResizeMode        { return ContainerGetResizeMode(c) }
func (c *Container) PropagateExpose(child *Widget, event *D.EventExpose) {
	ContainerPropagateExpose(c, child, event)
}
func (c *Container) Remove(widget *Widget)           { ContainerRemove(c, widget) }
func (c *Container) ResizeChildren()                 { ContainerResizeChildren(c) }
func (c *Container) SetBorderWidth(borderWidth uint) { ContainerSetBorderWidth(c, borderWidth) }
func (c *Container) SetFocusChain(focusableWidgets *T.GList) {
	ContainerSetFocusChain(c, focusableWidgets)
}
func (c *Container) SetFocusChild(child *Widget) { ContainerSetFocusChild(c, child) }
func (c *Container) SetFocusHadjustment(adjustment *Adjustment) {
	ContainerSetFocusHadjustment(c, adjustment)
}
func (c *Container) SetFocusVadjustment(adjustment *Adjustment) {
	ContainerSetFocusVadjustment(c, adjustment)
}
func (c *Container) SetReallocateRedraws(needsRedraws bool) {
	ContainerSetReallocateRedraws(c, needsRedraws)
}
func (c *Container) SetResizeMode(resizeMode ResizeMode) { ContainerSetResizeMode(c, resizeMode) }
func (c *Container) UnsetFocusChain()                    { ContainerUnsetFocusChain(c) }

type ContainerClass struct {
	ParentClass      WidgetClass
	Add              func(c *Container, w *Widget)
	Remove           func(c *Container, w *Widget)
	CheckResize      func(c *Container)
	Forall           func(c *Container, includeInternals bool, callback Callback, callbackData T.Gpointer)
	SetFocusChild    func(c *Container, w *Widget)
	ChildType        func(c *Container) O.Type
	CompositeName    func(c *Container, child *Widget) *T.Gchar
	SetChildProperty func(c *Container, child *Widget, propertyId uint, value *T.GValue, pspec *T.GParamSpec)
	GetChildProperty func(c *Container, child *Widget, propertyId uint, value *T.GValue, pspec *T.GParamSpec)
	_, _, _, _       func()
}

var (
	ContainerClassFindChildProperty   func(cclass *T.GObjectClass, propertyName string) *T.GParamSpec
	ContainerClassListChildProperties func(cclass *T.GObjectClass, nProperties *uint) **T.GParamSpec

	ContainerClassInstallChildProperty func(cclass *ContainerClass, propertyId uint, pspec *T.GParamSpec)
)

func (c *ContainerClass) InstallChildProperty(propertyId uint, pspec *T.GParamSpec) {
	ContainerClassInstallChildProperty(c, propertyId, pspec)
}

type CornerType Enum

const (
	CORNER_TOP_LEFT CornerType = iota
	CORNER_BOTTOM_LEFT
	CORNER_TOP_RIGHT
	CORNER_BOTTOM_RIGHT
)

var CornerTypeGetType func() O.Type

type (
	CTree struct {
		Clist       CList
		LinesGc     *D.GC
		TreeIndent  int
		TreeSpacing int
		TreeColumn  int
		Bits        uint
		// LineStyle : 2
		// ExpanderStyle : 2
		// ShowStub : 1
		DragCompare CTreeCompareDragFunc
	}

	CTreeNode struct {
		List T.GList
	}

	CTreeRow struct {
		Row           CListRow
		Parent        *CTreeNode
		Sibling       *CTreeNode
		Children      *CTreeNode
		Pixmap_closed *D.Pixmap
		Mask_closed   *T.GdkBitmap
		Pixmap_opened *D.Pixmap
		Mask_opened   *T.GdkBitmap
		Level         uint16
		Bits          uint
		// IsLeaf : 1
		// Expanded : 1
	}

	CTreeFunc func(ctree *CTree, node *CTreeNode, data T.Gpointer)

	CTreeCompareDragFunc func(c *CTree,
		sourceNode, newParent, newSibling *CTreeNode) bool

	CTreeGNodeFunc func(ctree *CTree, depth uint,
		gnode *T.GNode, cnode *CTreeNode,
		data T.Gpointer) bool
)

type CTreeExpanderStyle Enum

const (
	CTREE_EXPANDER_NONE CTreeExpanderStyle = iota
	CTREE_EXPANDER_SQUARE
	CTREE_EXPANDER_TRIANGLE
	CTREE_EXPANDER_CIRCULAR
)

type CTreeLineStyle Enum

const (
	CTREE_LINES_NONE CTreeLineStyle = iota
	CTREE_LINES_SOLID
	CTREE_LINES_DOTTED
	CTREE_LINES_TABBED
)

var (
	CtreeGetType       func() O.Type
	CtreeNew           func(columns int, treeColumn int) *Widget
	CtreeNewWithTitles func(columns int, treeColumn int, titles []string) *Widget

	CtreeNodeGetType func() O.Type
	CtreePosGetType  func() O.Type

	CtreeExpanderStyleGetType func() O.Type
	CtreeExpansionTypeGetType func() O.Type
	CtreeLineStyleGetType     func() O.Type
)

var (
	CtreeCollapse                 func(c *CTree, node *CTreeNode)
	CtreeCollapseRecursive        func(c *CTree, node *CTreeNode)
	CtreeCollapseToDepth          func(c *CTree, node *CTreeNode, depth int)
	CtreeExpand                   func(c *CTree, node *CTreeNode)
	CtreeExpandRecursive          func(c *CTree, node *CTreeNode)
	CtreeExpandToDepth            func(c *CTree, node *CTreeNode, depth int)
	CtreeExportToGnode            func(c *CTree, parent, sibling *T.GNode, node *CTreeNode, f CTreeGNodeFunc, data T.Gpointer) *T.GNode
	CtreeFind                     func(c *CTree, node *CTreeNode, child *CTreeNode) bool
	CtreeFindAllByRowData         func(c *CTree, node *CTreeNode, data T.Gpointer) *T.GList
	CtreeFindAllByRowDataCustom   func(c *CTree, node *CTreeNode, data T.Gpointer, f T.GCompareFunc) *T.GList
	CtreeFindByRowData            func(c *CTree, node *CTreeNode, data T.Gpointer) *CTreeNode
	CtreeFindByRowDataCustom      func(c *CTree, node *CTreeNode, data T.Gpointer, f T.GCompareFunc) *CTreeNode
	CtreeFindNodePtr              func(c *CTree, ctreeRow *CTreeRow) *CTreeNode
	CtreeGetNodeInfo              func(c *CTree, node *CTreeNode, text []string, spacing *uint8, pixmapClosed **D.Pixmap, maskClosed **T.GdkBitmap, pixmapOpened **D.Pixmap, maskOpened **T.GdkBitmap, isLeaf, expanded *bool) bool
	CtreeInsertGnode              func(c *CTree, parent *CTreeNode, sibling *CTreeNode, gnode *T.GNode, f CTreeGNodeFunc, data T.Gpointer) *CTreeNode
	CtreeInsertNode               func(c *CTree, parent *CTreeNode, sibling *CTreeNode, text []string, spacing uint8, pixmapClosed *D.Pixmap, maskClosed *T.GdkBitmap, pixmapOpened *D.Pixmap, maskOpened *T.GdkBitmap, isLeaf bool, expanded bool) *CTreeNode
	CtreeIsAncestor               func(c *CTree, node *CTreeNode, child *CTreeNode) bool
	CtreeIsHotSpot                func(c *CTree, x int, y int) bool
	CtreeIsViewable               func(c *CTree, node *CTreeNode) bool
	CtreeLast                     func(c *CTree, node *CTreeNode) *CTreeNode
	CtreeMove                     func(c *CTree, node *CTreeNode, newParent *CTreeNode, newSibling *CTreeNode)
	CtreeNodeGetCellStyle         func(c *CTree, node *CTreeNode, column int) *Style
	CtreeNodeGetCellType          func(c *CTree, node *CTreeNode, column int) CellType
	CtreeNodeGetPixmap            func(c *CTree, node *CTreeNode, column int, pixmap **D.Pixmap, mask **T.GdkBitmap) bool
	CtreeNodeGetPixtext           func(c *CTree, node *CTreeNode, column int, text []string, spacing *uint8, pixmap **D.Pixmap, mask **T.GdkBitmap) bool
	CtreeNodeGetRowData           func(c *CTree, node *CTreeNode) T.Gpointer
	CtreeNodeGetRowStyle          func(c *CTree, node *CTreeNode) *Style
	CtreeNodeGetSelectable        func(c *CTree, node *CTreeNode) bool
	CtreeNodeGetText              func(c *CTree, node *CTreeNode, column int, text []string) bool
	CtreeNodeIsVisible            func(c *CTree, node *CTreeNode) Visibility
	CtreeNodeMoveto               func(c *CTree, node *CTreeNode, column int, rowAlign, colAlign float32)
	CtreeNodeNth                  func(c *CTree, row uint) *CTreeNode
	CtreeNodeSetBackground        func(c *CTree, node *CTreeNode, color *D.Color)
	CtreeNodeSetCellStyle         func(c *CTree, node *CTreeNode, column int, style *Style)
	CtreeNodeSetForeground        func(c *CTree, node *CTreeNode, color *D.Color)
	CtreeNodeSetPixmap            func(c *CTree, node *CTreeNode, column int, pixmap *D.Pixmap, mask *T.GdkBitmap)
	CtreeNodeSetPixtext           func(c *CTree, node *CTreeNode, column int, text string, spacing uint8, pixmap *D.Pixmap, mask *T.GdkBitmap)
	CtreeNodeSetRowData           func(c *CTree, node *CTreeNode, data T.Gpointer)
	CtreeNodeSetRowDataFull       func(c *CTree, node *CTreeNode, data T.Gpointer, destroy T.GDestroyNotify)
	CtreeNodeSetRowStyle          func(c *CTree, node *CTreeNode, style *Style)
	CtreeNodeSetSelectable        func(c *CTree, node *CTreeNode, selectable bool)
	CtreeNodeSetShift             func(c *CTree, node *CTreeNode, column, vertical, horizontal int)
	CtreeNodeSetText              func(c *CTree, node *CTreeNode, column int, text string)
	CtreePostRecursive            func(c *CTree, node *CTreeNode, f CTreeFunc, data T.Gpointer)
	CtreePostRecursiveToDepth     func(c *CTree, node *CTreeNode, depth int, f CTreeFunc, data T.Gpointer)
	CtreePreRecursive             func(c *CTree, node *CTreeNode, f CTreeFunc, data T.Gpointer)
	CtreePreRecursiveToDepth      func(c *CTree, node *CTreeNode, depth int, f CTreeFunc, data T.Gpointer)
	CtreeRealSelectRecursive      func(c *CTree, node *CTreeNode, state int)
	CtreeRemoveNode               func(c *CTree, node *CTreeNode)
	CtreeSelect                   func(c *CTree, node *CTreeNode)
	CtreeSelectRecursive          func(c *CTree, node *CTreeNode)
	CtreeSetDragCompareFunc       func(c *CTree, cmpFunc CTreeCompareDragFunc)
	CtreeSetExpanderStyle         func(c *CTree, expanderStyle CTreeExpanderStyle)
	CtreeSetIndent                func(c *CTree, indent int)
	CtreeSetLineStyle             func(c *CTree, lineStyle CTreeLineStyle)
	CtreeSetNodeInfo              func(c *CTree, node *CTreeNode, text string, spacing uint8, pixmapClosed *D.Pixmap, maskClosed *T.GdkBitmap, pixmapOpened *D.Pixmap, maskOpened *T.GdkBitmap, isLeaf, expanded bool)
	CtreeSetShowStub              func(c *CTree, showStub bool)
	CtreeSetSpacing               func(c *CTree, spacing int)
	CtreeSortNode                 func(c *CTree, node *CTreeNode)
	CtreeSortRecursive            func(c *CTree, node *CTreeNode)
	CtreeToggleExpansion          func(c *CTree, node *CTreeNode)
	CtreeToggleExpansionRecursive func(c *CTree, node *CTreeNode)
	CtreeUnselect                 func(c *CTree, node *CTreeNode)
	CtreeUnselectRecursive        func(c *CTree, node *CTreeNode)
)

func (c *CTree) Collapse(node *CTreeNode)                   { CtreeCollapse(c, node) }
func (c *CTree) CollapseRecursive(node *CTreeNode)          { CtreeCollapseRecursive(c, node) }
func (c *CTree) CollapseToDepth(node *CTreeNode, depth int) { CtreeCollapseToDepth(c, node, depth) }
func (c *CTree) Expand(node *CTreeNode)                     { CtreeExpand(c, node) }
func (c *CTree) ExpandRecursive(node *CTreeNode)            { CtreeExpandRecursive(c, node) }
func (c *CTree) ExpandToDepth(node *CTreeNode, depth int)   { CtreeExpandToDepth(c, node, depth) }
func (c *CTree) ExportToGnode(parent *T.GNode, sibling *T.GNode, node *CTreeNode, f CTreeGNodeFunc, data T.Gpointer) *T.GNode {
	return CtreeExportToGnode(c, parent, sibling, node, f, data)
}
func (c *CTree) Find(node, child *CTreeNode) bool { return CtreeFind(c, node, child) }
func (c *CTree) FindAllByRowData(node *CTreeNode, data T.Gpointer) *T.GList {
	return CtreeFindAllByRowData(c, node, data)
}
func (c *CTree) FindAllByRowDataCustom(node *CTreeNode, data T.Gpointer, f T.GCompareFunc) *T.GList {
	return CtreeFindAllByRowDataCustom(c, node, data, f)
}
func (c *CTree) FindByRowData(node *CTreeNode, data T.Gpointer) *CTreeNode {
	return CtreeFindByRowData(c, node, data)
}
func (c *CTree) FindByRowDataCustom(node *CTreeNode, data T.Gpointer, f T.GCompareFunc) *CTreeNode {
	return CtreeFindByRowDataCustom(c, node, data, f)
}
func (c *CTree) FindNodePtr(ctreeRow *CTreeRow) *CTreeNode { return CtreeFindNodePtr(c, ctreeRow) }
func (c *CTree) GetNodeInfo(node *CTreeNode, text []string, spacing *uint8, pixmapClosed **D.Pixmap, maskClosed **T.GdkBitmap, pixmapOpened **D.Pixmap, maskOpened **T.GdkBitmap, isLeaf, expanded *bool) bool {
	return CtreeGetNodeInfo(c, node, text, spacing, pixmapClosed, maskClosed, pixmapOpened, maskOpened, isLeaf, expanded)
}
func (c *CTree) InsertGnode(parent, sibling *CTreeNode, gnode *T.GNode, f CTreeGNodeFunc, data T.Gpointer) *CTreeNode {
	return CtreeInsertGnode(c, parent, sibling, gnode, f, data)
}
func (c *CTree) InsertNode(parent, sibling *CTreeNode, text []string, spacing uint8, pixmapClosed *D.Pixmap, maskClosed *T.GdkBitmap, pixmapOpened *D.Pixmap, maskOpened *T.GdkBitmap, isLeaf, expanded bool) *CTreeNode {
	return CtreeInsertNode(c, parent, sibling, text, spacing, pixmapClosed, maskClosed, pixmapOpened, maskOpened, isLeaf, expanded)
}
func (c *CTree) IsAncestor(node, child *CTreeNode) bool { return CtreeIsAncestor(c, node, child) }
func (c *CTree) IsHotSpot(x int, y int) bool            { return CtreeIsHotSpot(c, x, y) }
func (c *CTree) IsViewable(node *CTreeNode) bool        { return CtreeIsViewable(c, node) }
func (c *CTree) Last(node *CTreeNode) *CTreeNode        { return CtreeLast(c, node) }
func (c *CTree) Move(node, newParent, newSibling *CTreeNode) {
	CtreeMove(c, node, newParent, newSibling)
}
func (c *CTree) NodeGetCellStyle(node *CTreeNode, column int) *Style {
	return CtreeNodeGetCellStyle(c, node, column)
}
func (c *CTree) NodeGetCellType(node *CTreeNode, column int) CellType {
	return CtreeNodeGetCellType(c, node, column)
}
func (c *CTree) NodeGetPixmap(node *CTreeNode, column int, pixmap **D.Pixmap, mask **T.GdkBitmap) bool {
	return CtreeNodeGetPixmap(c, node, column, pixmap, mask)
}
func (c *CTree) NodeGetPixtext(node *CTreeNode, column int, text []string, spacing *uint8, pixmap **D.Pixmap, mask **T.GdkBitmap) bool {
	return CtreeNodeGetPixtext(c, node, column, text, spacing, pixmap, mask)
}
func (c *CTree) NodeGetRowData(node *CTreeNode) T.Gpointer { return CtreeNodeGetRowData(c, node) }
func (c *CTree) NodeGetRowStyle(node *CTreeNode) *Style    { return CtreeNodeGetRowStyle(c, node) }
func (c *CTree) NodeGetSelectable(node *CTreeNode) bool    { return CtreeNodeGetSelectable(c, node) }
func (c *CTree) NodeGetText(node *CTreeNode, column int, text []string) bool {
	return CtreeNodeGetText(c, node, column, text)
}
func (c *CTree) NodeIsVisible(node *CTreeNode) Visibility { return CtreeNodeIsVisible(c, node) }
func (c *CTree) NodeMoveto(node *CTreeNode, column int, rowAlign, colAlign float32) {
	CtreeNodeMoveto(c, node, column, rowAlign, colAlign)
}
func (c *CTree) NodeNth(row uint) *CTreeNode { return CtreeNodeNth(c, row) }
func (c *CTree) NodeSetBackground(node *CTreeNode, color *D.Color) {
	CtreeNodeSetBackground(c, node, color)
}
func (c *CTree) NodeSetCellStyle(node *CTreeNode, column int, style *Style) {
	CtreeNodeSetCellStyle(c, node, column, style)
}
func (c *CTree) NodeSetForeground(node *CTreeNode, color *D.Color) {
	CtreeNodeSetForeground(c, node, color)
}
func (c *CTree) NodeSetPixmap(node *CTreeNode, column int, pixmap *D.Pixmap, mask *T.GdkBitmap) {
	CtreeNodeSetPixmap(c, node, column, pixmap, mask)
}
func (c *CTree) NodeSetPixtext(node *CTreeNode, column int, text string, spacing uint8, pixmap *D.Pixmap, mask *T.GdkBitmap) {
	CtreeNodeSetPixtext(c, node, column, text, spacing, pixmap, mask)
}
func (c *CTree) NodeSetRowData(node *CTreeNode, data T.Gpointer) { CtreeNodeSetRowData(c, node, data) }
func (c *CTree) NodeSetRowDataFull(node *CTreeNode, data T.Gpointer, destroy T.GDestroyNotify) {
	CtreeNodeSetRowDataFull(c, node, data, destroy)
}
func (c *CTree) NodeSetRowStyle(node *CTreeNode, style *Style) {
	CtreeNodeSetRowStyle(c, node, style)
}
func (c *CTree) NodeSetSelectable(node *CTreeNode, selectable bool) {
	CtreeNodeSetSelectable(c, node, selectable)
}
func (c *CTree) NodeSetShift(node *CTreeNode, column, vertical, horizontal int) {
	CtreeNodeSetShift(c, node, column, vertical, horizontal)
}
func (c *CTree) NodeSetText(node *CTreeNode, column int, text string) {
	CtreeNodeSetText(c, node, column, text)
}
func (c *CTree) PostRecursive(node *CTreeNode, f CTreeFunc, data T.Gpointer) {
	CtreePostRecursive(c, node, f, data)
}
func (c *CTree) PostRecursiveToDepth(node *CTreeNode, depth int, f CTreeFunc, data T.Gpointer) {
	CtreePostRecursiveToDepth(c, node, depth, f, data)
}
func (c *CTree) PreRecursive(node *CTreeNode, f CTreeFunc, data T.Gpointer) {
	CtreePreRecursive(c, node, f, data)
}
func (c *CTree) PreRecursiveToDepth(node *CTreeNode, depth int, f CTreeFunc, data T.Gpointer) {
	CtreePreRecursiveToDepth(c, node, depth, f, data)
}
func (c *CTree) RealSelectRecursive(node *CTreeNode, state int) {
	CtreeRealSelectRecursive(c, node, state)
}
func (c *CTree) RemoveNode(node *CTreeNode)                      { CtreeRemoveNode(c, node) }
func (c *CTree) Select(node *CTreeNode)                          { CtreeSelect(c, node) }
func (c *CTree) SelectRecursive(node *CTreeNode)                 { CtreeSelectRecursive(c, node) }
func (c *CTree) SetDragCompareFunc(cmpFunc CTreeCompareDragFunc) { CtreeSetDragCompareFunc(c, cmpFunc) }
func (c *CTree) SetExpanderStyle(expanderStyle CTreeExpanderStyle) {
	CtreeSetExpanderStyle(c, expanderStyle)
}
func (c *CTree) SetIndent(indent int)                  { CtreeSetIndent(c, indent) }
func (c *CTree) SetLineStyle(lineStyle CTreeLineStyle) { CtreeSetLineStyle(c, lineStyle) }
func (c *CTree) SetNodeInfo(node *CTreeNode, text string, spacing uint8, pixmapClosed *D.Pixmap, maskClosed *T.GdkBitmap, pixmapOpened *D.Pixmap, maskOpened *T.GdkBitmap, isLeaf, expanded bool) {
	CtreeSetNodeInfo(c, node, text, spacing, pixmapClosed, maskClosed, pixmapOpened, maskOpened, isLeaf, expanded)
}
func (c *CTree) SetShowStub(showStub bool)                { CtreeSetShowStub(c, showStub) }
func (c *CTree) SetSpacing(spacing int)                   { CtreeSetSpacing(c, spacing) }
func (c *CTree) SortNode(node *CTreeNode)                 { CtreeSortNode(c, node) }
func (c *CTree) SortRecursive(node *CTreeNode)            { CtreeSortRecursive(c, node) }
func (c *CTree) ToggleExpansion(node *CTreeNode)          { CtreeToggleExpansion(c, node) }
func (c *CTree) ToggleExpansionRecursive(node *CTreeNode) { CtreeToggleExpansionRecursive(c, node) }
func (c *CTree) Unselect(node *CTreeNode)                 { CtreeUnselect(c, node) }
func (c *CTree) UnselectRecursive(node *CTreeNode)        { CtreeUnselectRecursive(c, node) }

type Curve struct {
	Graph        DrawingArea
	CursorType   int
	MinX         float32
	MaxX         float32
	MinY         float32
	MaxY         float32
	Pixmap       *D.Pixmap
	CurveType    CurveType
	Height       int
	GrabPoint    int
	Last         int
	NumPoints    int
	Point        *D.Point
	NumCtlpoints int
	Ctlpoint     [2]*float32 //TODO(t): float32 (*ctlpoint)[2]; ???
}

var (
	CurveGetType func() O.Type
	CurveNew     func() *Widget

	CurveGetVector    func(c *Curve, veclen int, vector *float32)
	CurveReset        func(c *Curve)
	CurveSetCurveType func(c *Curve, t CurveType)
	CurveSetGamma     func(c *Curve, gamma float32)
	CurveSetRange     func(c *Curve, minX, maxX, minY, maxY float32)
	CurveSetVector    func(c *Curve, veclen int, vector *float32)
)

func (c *Curve) GetVector(veclen int, vector *float32)   { CurveGetVector(c, veclen, vector) }
func (c *Curve) Reset()                                  { CurveReset(c) }
func (c *Curve) SetCurveType(t CurveType)                { CurveSetCurveType(c, t) }
func (c *Curve) SetGamma(gamma float32)                  { CurveSetGamma(c, gamma) }
func (c *Curve) SetRange(minX, maxX, minY, maxY float32) { CurveSetRange(c, minX, maxX, minY, maxY) }
func (c *Curve) SetVector(veclen int, vector *float32)   { CurveSetVector(c, veclen, vector) }

type CurveType Enum

const (
	CURVE_TYPE_LINEAR CurveType = iota
	CURVE_TYPE_SPLINE
	CURVE_TYPE_FREE
)

var CurveTypeGetType func() O.Type
