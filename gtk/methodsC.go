// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
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
		Gc                *T.GdkGC
		Xor_gc            *T.GdkGC
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
	CalendarDisplayOptionsGetType func() T.GType
	CalendarGetType               func() T.GType
	CalendarNew                   func() *Widget
)

var (
	calendarClearMarks          func(c *Calendar)
	calendarDisplayOptions      func(c *Calendar, flags CalendarDisplayOptions)
	calendarFreeze              func(c *Calendar)
	calendarGetDate             func(c *Calendar, year, month, day *uint)
	calendarGetDetailHeightRows func(c *Calendar) int
	calendarGetDetailWidthChars func(c *Calendar) int
	calendarGetDisplayOptions   func(c *Calendar) CalendarDisplayOptions
	calendarMarkDay             func(c *Calendar, day uint) T.Gboolean
	calendarSelectDay           func(c *Calendar, day uint)
	calendarSelectMonth         func(c *Calendar, month, year uint) T.Gboolean
	calendarSetDetailFunc       func(c *Calendar, f CalendarDetailFunc, data T.Gpointer, destroy T.GDestroyNotify)
	calendarSetDetailHeightRows func(c *Calendar, rows int)
	calendarSetDetailWidthChars func(c *Calendar, chars int)
	calendarSetDisplayOptions   func(c *Calendar, flags CalendarDisplayOptions)
	calendarThaw                func(c *Calendar)
	calendarUnmarkDay           func(c *Calendar, day uint) T.Gboolean
)

func (c *Calendar) ClearMarks()                                 { calendarClearMarks(c) }
func (c *Calendar) DisplayOptions(flags CalendarDisplayOptions) { calendarDisplayOptions(c, flags) }
func (c *Calendar) Freeze()                                     { calendarFreeze(c) }
func (c *Calendar) GetDate(year, month, day *uint)              { calendarGetDate(c, year, month, day) }
func (c *Calendar) GetDetailHeightRows() int                    { return calendarGetDetailHeightRows(c) }
func (c *Calendar) GetDetailWidthChars() int                    { return calendarGetDetailWidthChars(c) }
func (c *Calendar) GetDisplayOptions() CalendarDisplayOptions   { return calendarGetDisplayOptions(c) }
func (c *Calendar) MarkDay(day uint) T.Gboolean                 { return calendarMarkDay(c, day) }
func (c *Calendar) SelectDay(day uint)                          { calendarSelectDay(c, day) }
func (c *Calendar) SelectMonth(month, year uint) T.Gboolean {
	return calendarSelectMonth(c, month, year)
}
func (c *Calendar) SetDetailFunc(f CalendarDetailFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	calendarSetDetailFunc(c, f, data, destroy)
}
func (c *Calendar) SetDetailHeightRows(rows int)  { calendarSetDetailHeightRows(c, rows) }
func (c *Calendar) SetDetailWidthChars(chars int) { calendarSetDetailWidthChars(c, chars) }
func (c *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	calendarSetDisplayOptions(c, flags)
}
func (c *Calendar) Thaw()                         { calendarThaw(c) }
func (c *Calendar) UnmarkDay(day uint) T.Gboolean { return calendarUnmarkDay(c, day) }

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

var CellEditableGetType func() T.GType

var (
	cellEditableEditingDone  func(c *CellEditable)
	cellEditableRemoveWidget func(c *CellEditable)
	cellEditableStartEditing func(c *CellEditable, event *D.Event)
)

func (c *CellEditable) EditingDone()                { cellEditableEditingDone(c) }
func (c *CellEditable) RemoveWidget()               { cellEditableRemoveWidget(c) }
func (c *CellEditable) StartEditing(event *D.Event) { cellEditableStartEditing(c, event) }

type CellLayout struct{}

type CellLayoutDataFunc func(
	cellLayout *CellLayout,
	cell *CellRenderer,
	treeModel *TreeModel,
	iter *TreeIter,
	data T.Gpointer)

var (
	CellLayoutGetType func() T.GType

	cellLayoutAddAttribute    func(c *CellLayout, cell *CellRenderer, attribute string, column int)
	cellLayoutClear           func(c *CellLayout)
	cellLayoutClearAttributes func(c *CellLayout, cell *CellRenderer)
	cellLayoutGetCells        func(c *CellLayout) *T.GList
	cellLayoutPackEnd         func(c *CellLayout, cell *CellRenderer, expand T.Gboolean)
	cellLayoutPackStart       func(c *CellLayout, cell *CellRenderer, expand T.Gboolean)
	cellLayoutReorder         func(c *CellLayout, cell *CellRenderer, position int)
	cellLayoutSetAttributes   func(c *CellLayout, cell *CellRenderer, v ...VArg)
	cellLayoutSetCellDataFunc func(c *CellLayout, cell *CellRenderer, f CellLayoutDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
)

func (c *CellLayout) AddAttribute(cell *CellRenderer, attribute string, column int) {
	cellLayoutAddAttribute(c, cell, attribute, column)
}
func (c *CellLayout) Clear()                             { cellLayoutClear(c) }
func (c *CellLayout) ClearAttributes(cell *CellRenderer) { cellLayoutClearAttributes(c, cell) }
func (c *CellLayout) GetCells() *T.GList                 { return cellLayoutGetCells(c) }
func (c *CellLayout) PackEnd(cell *CellRenderer, expand T.Gboolean) {
	cellLayoutPackEnd(c, cell, expand)
}
func (c *CellLayout) PackStart(cell *CellRenderer, expand T.Gboolean) {
	cellLayoutPackStart(c, cell, expand)
}
func (c *CellLayout) Reorder(cell *CellRenderer, position int)    { cellLayoutReorder(c, cell, position) }
func (c *CellLayout) SetAttributes(cell *CellRenderer, v ...VArg) { cellLayoutSetAttributes(c, cell, v) }
func (c *CellLayout) SetCellDataFunc(cell *CellRenderer, f CellLayoutDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	cellLayoutSetCellDataFunc(c, cell, f, funcData, destroy)
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
	CellRendererGetType          func() T.GType
	CellRendererAccelGetType     func() T.GType
	CellRendererAccelNew         func() *CellRenderer
	CellRendererAccelModeGetType func() T.GType
	CellRendererComboGetType     func() T.GType
	CellRendererComboNew         func() *CellRenderer
	CellRendererModeGetType      func() T.GType
	CellRendererPixbufGetType    func() T.GType
	CellRendererPixbufNew        func() *CellRenderer
	CellRendererProgressGetType  func() T.GType
	CellRendererProgressNew      func() *CellRenderer
	CellRendererSpinGetType      func() T.GType
	CellRendererSpinNew          func() *CellRenderer
	CellRendererSpinnerGetType   func() T.GType
	CellRendererSpinnerNew       func() *CellRenderer
	CellRendererStateGetType     func() T.GType
	CellRendererToggleGetType    func() T.GType
	CellRendererToggleNew        func() *CellRenderer
)

var (
	cellRendererActivate        func(c *CellRenderer, event *D.Event, widget *Widget, path string, backgroundArea, cellArea *T.GdkRectangle, flags CellRendererState) T.Gboolean
	cellRendererEditingCanceled func(c *CellRenderer)
	cellRendererGetAlignment    func(c *CellRenderer, xalign, yalign *float32)
	cellRendererGetFixedSize    func(c *CellRenderer, width, height *int)
	cellRendererGetPadding      func(c *CellRenderer, xpad, ypad *int)
	cellRendererGetSensitive    func(c *CellRenderer) T.Gboolean
	cellRendererGetSize         func(c *CellRenderer, widget *Widget, cellArea *T.GdkRectangle, xOffset, yOffset, width, height *int)
	cellRendererGetVisible      func(c *CellRenderer) T.Gboolean
	cellRendererRender          func(c *CellRenderer, window *D.Window, widget *Widget, backgroundArea, cellArea, exposeArea *T.GdkRectangle, flags CellRendererState)
	cellRendererSetAlignment    func(c *CellRenderer, xalign, yalign float32)
	cellRendererSetFixedSize    func(c *CellRenderer, width, height int)
	cellRendererSetPadding      func(c *CellRenderer, xpad, ypad int)
	cellRendererSetSensitive    func(c *CellRenderer, sensitive T.Gboolean)
	cellRendererSetVisible      func(c *CellRenderer, visible T.Gboolean)
	cellRendererStartEditing    func(c *CellRenderer, event *D.Event, widget *Widget, path string, backgroundArea, cellArea *T.GdkRectangle, flags CellRendererState) *CellEditable
	cellRendererStopEditing     func(c *CellRenderer, canceled T.Gboolean)
)

func (c *CellRenderer) Activate(event *D.Event, widget *Widget, path string, backgroundArea, cellArea *T.GdkRectangle, flags CellRendererState) T.Gboolean {
	return cellRendererActivate(c, event, widget, path, backgroundArea, cellArea, flags)
}
func (c *CellRenderer) EditingCanceled() { cellRendererEditingCanceled(c) }
func (c *CellRenderer) GetAlignment(xalign, yalign *float32) {
	cellRendererGetAlignment(c, xalign, yalign)
}
func (c *CellRenderer) GetFixedSize(width, height *int) { cellRendererGetFixedSize(c, width, height) }
func (c *CellRenderer) GetPadding(xpad, ypad *int)      { cellRendererGetPadding(c, xpad, ypad) }
func (c *CellRenderer) GetSensitive() T.Gboolean        { return cellRendererGetSensitive(c) }
func (c *CellRenderer) GetSize(widget *Widget, cellArea *T.GdkRectangle, xOffset, yOffset, width, height *int) {
	cellRendererGetSize(c, widget, cellArea, xOffset, yOffset, width, height)
}
func (c *CellRenderer) GetVisible() T.Gboolean { return cellRendererGetVisible(c) }
func (c *CellRenderer) Render(window *D.Window, widget *Widget, backgroundArea, cellArea, exposeArea *T.GdkRectangle, flags CellRendererState) {
	cellRendererRender(c, window, widget, backgroundArea, cellArea, exposeArea, flags)
}
func (c *CellRenderer) SetAlignment(xalign, yalign float32) {
	cellRendererSetAlignment(c, xalign, yalign)
}
func (c *CellRenderer) SetFixedSize(width, height int)    { cellRendererSetFixedSize(c, width, height) }
func (c *CellRenderer) SetPadding(xpad, ypad int)         { cellRendererSetPadding(c, xpad, ypad) }
func (c *CellRenderer) SetSensitive(sensitive T.Gboolean) { cellRendererSetSensitive(c, sensitive) }
func (c *CellRenderer) SetVisible(visible T.Gboolean)     { cellRendererSetVisible(c, visible) }
func (c *CellRenderer) StartEditing(event *D.Event, widget *Widget, path string, backgroundArea, cellArea *T.GdkRectangle, flags CellRendererState) *CellEditable {
	return cellRendererStartEditing(c, event, widget, path, backgroundArea, cellArea, flags)
}
func (c *CellRenderer) StopEditing(canceled T.Gboolean) { cellRendererStopEditing(c, canceled) }

type CellRendererText struct {
	Parent          CellRenderer
	Text            *T.Gchar
	Font            *T.PangoFontDescription
	FontScale       float64
	Foreground      T.PangoColor
	Background      T.PangoColor
	ExtraAttrs      *T.PangoAttrList
	UnderlineStyle  T.PangoUnderline
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
	CellRendererTextGetType func() T.GType
	CellRendererTextNew     func() *CellRenderer

	cellRendererTextSetFixedHeightFromFont func(renderer *CellRendererText, numberOfRows int)
)

func (c *CellRendererText) SetFixedHeightFromFont(numberOfRows int) {
	cellRendererTextSetFixedHeightFromFont(c, numberOfRows)
}

type CellRendererToggle struct {
	Parent CellRenderer
	Bits   uint
	// Active : 1
	// Activatable : 1
	// Radio : 1
}

var (
	cellRendererToggleGetActivatable func(c *CellRendererToggle) T.Gboolean
	cellRendererToggleGetActive      func(c *CellRendererToggle) T.Gboolean
	cellRendererToggleGetRadio       func(c *CellRendererToggle) T.Gboolean
	cellRendererToggleSetActivatable func(c *CellRendererToggle, setting T.Gboolean)
	cellRendererToggleSetActive      func(c *CellRendererToggle, setting T.Gboolean)
	cellRendererToggleSetRadio       func(c *CellRendererToggle, radio T.Gboolean)
)

func (c *CellRendererToggle) GetActivatable() T.Gboolean { return cellRendererToggleGetActivatable(c) }
func (c *CellRendererToggle) GetActive() T.Gboolean      { return cellRendererToggleGetActive(c) }
func (c *CellRendererToggle) GetRadio() T.Gboolean       { return cellRendererToggleGetRadio(c) }
func (c *CellRendererToggle) SetActivatable(setting T.Gboolean) {
	cellRendererToggleSetActivatable(c, setting)
}
func (c *CellRendererToggle) SetActive(setting T.Gboolean) { cellRendererToggleSetActive(c, setting) }
func (c *CellRendererToggle) SetRadio(radio T.Gboolean)    { cellRendererToggleSetRadio(c, radio) }

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

var CellTypeGetType func() T.GType

var (
	CellViewGetType       func() T.GType
	CellViewNew           func() *Widget
	CellViewNewWithText   func(text string) *Widget
	CellViewNewWithMarkup func(markup string) *Widget
	CellViewNewWithPixbuf func(pixbuf *D.Pixbuf) *Widget
)

var (
	cellViewGetCellRenderers   func(c *CellView) *T.GList
	cellViewGetDisplayedRow    func(c *CellView) *TreePath
	cellViewGetModel           func(c *CellView) *TreeModel
	cellViewGetSizeOfRow       func(c *CellView, path *TreePath, requisition *Requisition) T.Gboolean
	cellViewSetBackgroundColor func(c *CellView, color *D.Color)
	cellViewSetDisplayedRow    func(c *CellView, path *TreePath)
	cellViewSetModel           func(c *CellView, model *TreeModel)
)

func (c *CellView) GetCellRenderers() *T.GList { return cellViewGetCellRenderers(c) }
func (c *CellView) GetDisplayedRow() *TreePath { return cellViewGetDisplayedRow(c) }
func (c *CellView) GetModel() *TreeModel       { return cellViewGetModel(c) }
func (c *CellView) GetSizeOfRow(path *TreePath, requisition *Requisition) T.Gboolean {
	return cellViewGetSizeOfRow(c, path, requisition)
}
func (c *CellView) SetBackgroundColor(color *D.Color) { cellViewSetBackgroundColor(c, color) }
func (c *CellView) SetDisplayedRow(path *TreePath)    { cellViewSetDisplayedRow(c, path) }
func (c *CellView) SetModel(model *TreeModel)         { cellViewSetModel(c, model) }

var (
	CheckButtonNew             func() *Widget
	CheckButtonNewWithLabel    func(label string) *Widget
	CheckButtonNewWithMnemonic func(label string) *Widget
)

type CheckButton struct {
	ToggleButton ToggleButton
}

var CheckButtonGetType func() T.GType

type CheckMenuItem struct {
	MenuItem MenuItem
	Bits     uint
	// Active : 1
	// AlwaysShowToggle : 1
	// Inconsistent : 1
	// DrawAsRadio : 1
}

var (
	CheckMenuItemGetType         func() T.GType
	CheckMenuItemNew             func() *Widget
	CheckMenuItemNewWithLabel    func(label string) *Widget
	CheckMenuItemNewWithMnemonic func(label string) *Widget
)

var (
	checkMenuItemGetActive       func(c *CheckMenuItem) T.Gboolean
	checkMenuItemGetDrawAsRadio  func(c *CheckMenuItem) T.Gboolean
	checkMenuItemGetInconsistent func(c *CheckMenuItem) T.Gboolean
	checkMenuItemSetActive       func(c *CheckMenuItem, isActive T.Gboolean)
	checkMenuItemSetDrawAsRadio  func(c *CheckMenuItem, drawAsRadio T.Gboolean)
	checkMenuItemSetInconsistent func(c *CheckMenuItem, setting T.Gboolean)
	checkMenuItemSetShowToggle   func(c *CheckMenuItem, always T.Gboolean)
	checkMenuItemToggled         func(c *CheckMenuItem)
)

func (c *CheckMenuItem) GetActive() T.Gboolean         { return checkMenuItemGetActive(c) }
func (c *CheckMenuItem) GetDrawAsRadio() T.Gboolean    { return checkMenuItemGetDrawAsRadio(c) }
func (c *CheckMenuItem) GetInconsistent() T.Gboolean   { return checkMenuItemGetInconsistent(c) }
func (c *CheckMenuItem) SetActive(isActive T.Gboolean) { checkMenuItemSetActive(c, isActive) }
func (c *CheckMenuItem) SetDrawAsRadio(drawAsRadio T.Gboolean) {
	checkMenuItemSetDrawAsRadio(c, drawAsRadio)
}
func (c *CheckMenuItem) SetInconsistent(setting T.Gboolean) { checkMenuItemSetInconsistent(c, setting) }
func (c *CheckMenuItem) SetShowToggle(always T.Gboolean)    { checkMenuItemSetShowToggle(c, always) }
func (c *CheckMenuItem) Toggled()                           { checkMenuItemToggled(c) }

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
	ClipboardGetType       func() T.GType
	ClipboardGet           func(selection D.Atom) *Clipboard
	ClipboardGetForDisplay func(display *D.Display, selection D.Atom) *Clipboard
	WidgetGetClipboard     func(widget *Widget, selection D.Atom) *Clipboard
)

var (
	clipboardClear                   func(c *Clipboard)
	clipboardGetDisplay              func(c *Clipboard) *D.Display
	clipboardGetOwner                func(c *Clipboard) *T.GObject
	clipboardRequestContents         func(c *Clipboard, target D.Atom, callback ClipboardReceivedFunc, userData T.Gpointer)
	clipboardRequestImage            func(c *Clipboard, callback ClipboardImageReceivedFunc, userData T.Gpointer)
	clipboardRequestRichText         func(c *Clipboard, buffer *TextBuffer, callback ClipboardRichTextReceivedFunc, userData T.Gpointer)
	clipboardRequestTargets          func(c *Clipboard, callback ClipboardTargetsReceivedFunc, userData T.Gpointer)
	clipboardRequestText             func(c *Clipboard, callback ClipboardTextReceivedFunc, userData T.Gpointer)
	clipboardRequestUris             func(c *Clipboard, callback ClipboardURIReceivedFunc, userData T.Gpointer)
	clipboardSetCanStore             func(c *Clipboard, targets *TargetEntry, nTargets int)
	clipboardSetImage                func(c *Clipboard, pixbuf *D.Pixbuf)
	clipboardSetText                 func(c *Clipboard, text string, len int)
	clipboardSetWithData             func(c *Clipboard, targets *TargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, userData T.Gpointer) T.Gboolean
	clipboardSetWithOwner            func(c *Clipboard, targets *TargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, owner *T.GObject) T.Gboolean
	clipboardStore                   func(c *Clipboard)
	clipboardWaitForContents         func(c *Clipboard, target D.Atom) *SelectionData
	clipboardWaitForImage            func(c *Clipboard) *D.Pixbuf
	clipboardWaitForRichText         func(c *Clipboard, buffer *TextBuffer, format *D.Atom, length *T.Gsize) *uint8
	clipboardWaitForTargets          func(c *Clipboard, targets **D.Atom, nTargets *int) T.Gboolean
	clipboardWaitForText             func(c *Clipboard) string
	clipboardWaitForUris             func(c *Clipboard) **T.Gchar
	clipboardWaitIsImageAvailable    func(c *Clipboard) T.Gboolean
	clipboardWaitIsRichTextAvailable func(c *Clipboard, buffer *TextBuffer) T.Gboolean
	clipboardWaitIsTargetAvailable   func(c *Clipboard, target D.Atom) T.Gboolean
	clipboardWaitIsTextAvailable     func(c *Clipboard) T.Gboolean
	clipboardWaitIsUrisAvailable     func(c *Clipboard) T.Gboolean
)

func (c *Clipboard) Clear()                 { clipboardClear(c) }
func (c *Clipboard) GetDisplay() *D.Display { return clipboardGetDisplay(c) }
func (c *Clipboard) GetOwner() *T.GObject   { return clipboardGetOwner(c) }
func (c *Clipboard) RequestContents(target D.Atom, callback ClipboardReceivedFunc, userData T.Gpointer) {
	clipboardRequestContents(c, target, callback, userData)
}
func (c *Clipboard) RequestImage(callback ClipboardImageReceivedFunc, userData T.Gpointer) {
	clipboardRequestImage(c, callback, userData)
}
func (c *Clipboard) RequestRichText(buffer *TextBuffer, callback ClipboardRichTextReceivedFunc, userData T.Gpointer) {
	clipboardRequestRichText(c, buffer, callback, userData)
}
func (c *Clipboard) RequestTargets(callback ClipboardTargetsReceivedFunc, userData T.Gpointer) {
	clipboardRequestTargets(c, callback, userData)
}
func (c *Clipboard) RequestText(callback ClipboardTextReceivedFunc, userData T.Gpointer) {
	clipboardRequestText(c, callback, userData)
}
func (c *Clipboard) RequestUris(callback ClipboardURIReceivedFunc, userData T.Gpointer) {
	clipboardRequestUris(c, callback, userData)
}
func (c *Clipboard) SetCanStore(targets *TargetEntry, nTargets int) {
	clipboardSetCanStore(c, targets, nTargets)
}
func (c *Clipboard) SetImage(pixbuf *D.Pixbuf)     { clipboardSetImage(c, pixbuf) }
func (c *Clipboard) SetText(text string, leng int) { clipboardSetText(c, text, leng) }
func (c *Clipboard) SetWithData(targets *TargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, userData T.Gpointer) T.Gboolean {
	return clipboardSetWithData(c, targets, nTargets, getFunc, clearFunc, userData)
}
func (c *Clipboard) SetWithOwner(targets *TargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, owner *T.GObject) T.Gboolean {
	return clipboardSetWithOwner(c, targets, nTargets, getFunc, clearFunc, owner)
}
func (c *Clipboard) Store() { clipboardStore(c) }
func (c *Clipboard) WaitForContents(target D.Atom) *SelectionData {
	return clipboardWaitForContents(c, target)
}
func (c *Clipboard) WaitForImage() *D.Pixbuf { return clipboardWaitForImage(c) }
func (c *Clipboard) WaitForRichText(buffer *TextBuffer, format *D.Atom, length *T.Gsize) *uint8 {
	return clipboardWaitForRichText(c, buffer, format, length)
}
func (c *Clipboard) WaitForTargets(targets **D.Atom, nTargets *int) T.Gboolean {
	return clipboardWaitForTargets(c, targets, nTargets)
}
func (c *Clipboard) WaitForText() string              { return clipboardWaitForText(c) }
func (c *Clipboard) WaitForUris() **T.Gchar           { return clipboardWaitForUris(c) }
func (c *Clipboard) WaitIsImageAvailable() T.Gboolean { return clipboardWaitIsImageAvailable(c) }
func (c *Clipboard) WaitIsRichTextAvailable(buffer *TextBuffer) T.Gboolean {
	return clipboardWaitIsRichTextAvailable(c, buffer)
}
func (c *Clipboard) WaitIsTargetAvailable(target D.Atom) T.Gboolean {
	return clipboardWaitIsTargetAvailable(c, target)
}
func (c *Clipboard) WaitIsTextAvailable() T.Gboolean { return clipboardWaitIsTextAvailable(c) }
func (c *Clipboard) WaitIsUrisAvailable() T.Gboolean { return clipboardWaitIsUrisAvailable(c) }

type (
	CList struct {
		Container          Container
		Flags              uint16
		_                  T.Gpointer
		_                  T.Gpointer
		Freeze_count       uint
		InternalAllocation T.GdkRectangle
		Rows               int
		RowHeight          int
		RowList            *T.GList
		RowListEnd         *T.GList
		Columns            int
		ColumnTitleArea    T.GdkRectangle
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
		XorGc              *T.GdkGC
		FgGc               *T.GdkGC
		BgGc               *T.GdkGC
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
		Area          T.GdkRectangle
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
	ClistGetType        func() T.GType
	ClistNew            func(columns int) *Widget
	ClistNewWithTitles  func(columns int, titles **T.Gchar) *Widget
	ClistDragPosGetType func() T.GType
)

var (
	clistAppend                 func(c *CList, text **T.Gchar) int
	clistClear                  func(c *CList)
	clistColumnsAutosize        func(c *CList) int
	clistColumnTitleActive      func(c *CList, column int)
	clistColumnTitlePassive     func(c *CList, column int)
	clistColumnTitlesActive     func(c *CList)
	clistColumnTitlesHide       func(c *CList)
	clistColumnTitlesPassive    func(c *CList)
	clistColumnTitlesShow       func(c *CList)
	clistFindRowFromData        func(c *CList, data T.Gpointer) int
	clistFreeze                 func(c *CList)
	clistGetCellStyle           func(c *CList, row, column int) *Style
	clistGetCellType            func(c *CList, row, column int) CellType
	clistGetColumnTitle         func(c *CList, column int) string
	clistGetColumnWidget        func(c *CList, column int) *Widget
	clistGetHadjustment         func(c *CList) *Adjustment
	clistGetPixmap              func(c *CList, row, column int, pixmap, mask **T.GdkBitmap) int
	clistGetPixtext             func(c *CList, row, column int, text **T.Gchar, spacing *uint8, pixmap, mask **T.GdkBitmap) int
	clistGetRowData             func(c *CList, row int) T.Gpointer
	clistGetRowStyle            func(c *CList, row int) *Style
	clistGetSelectable          func(c *CList, row int) T.Gboolean
	clistGetSelectionInfo       func(c *CList, x, y int, row, column *int) int
	clistGetText                func(c *CList, row, column int, text **T.Gchar) int
	clistGetVadjustment         func(c *CList) *Adjustment
	clistInsert                 func(c *CList, row int, text **T.Gchar) int
	clistMoveto                 func(c *CList, row, column int, rowAlign, colAlign float32)
	clistOptimalColumnWidth     func(c *CList, column int) int
	clistPrepend                func(c *CList, text **T.Gchar) int
	clistRemove                 func(c *CList, row int)
	clistRowIsVisible           func(c *CList, row int) Visibility
	clistRowMove                func(c *CList, sourceRow, destRow int)
	clistSelectAll              func(c *CList)
	clistSelectRow              func(c *CList, row, column int)
	clistSetAutoSort            func(c *CList, autoSort T.Gboolean)
	clistSetBackground          func(c *CList, row int, color *D.Color)
	clistSetButtonActions       func(c *CList, button uint, buttonActions uint8)
	clistSetCellStyle           func(c *CList, row, column int, style *Style)
	clistSetColumnAutoResize    func(c *CList, column int, autoResize T.Gboolean)
	clistSetColumnJustification func(c *CList, column int, justification Justification)
	clistSetColumnMaxWidth      func(c *CList, column, maxWidth int)
	clistSetColumnMinWidth      func(c *CList, column, minWidth int)
	clistSetColumnResizeable    func(c *CList, column int, resizeable T.Gboolean)
	clistSetColumnTitle         func(c *CList, column int, title string)
	clistSetColumnVisibility    func(c *CList, column int, visible T.Gboolean)
	clistSetColumnWidget        func(c *CList, column int, widget *Widget)
	clistSetColumnWidth         func(c *CList, column, width int)
	clistSetCompareFunc         func(c *CList, cmpFunc CListCompareFunc)
	clistSetForeground          func(c *CList, row int, color *D.Color)
	clistSetHadjustment         func(c *CList, adjustment *Adjustment)
	clistSetPixmap              func(c *CList, row, column int, pixmap *D.Pixmap, mask *T.GdkBitmap)
	clistSetPixtext             func(c *CList, row, column int, text string, spacing uint8, pixmap, mask *T.GdkBitmap)
	clistSetReorderable         func(c *CList, reorderable T.Gboolean)
	clistSetRowData             func(c *CList, row int, data T.Gpointer)
	clistSetRowDataFull         func(c *CList, row int, data T.Gpointer, destroy T.GDestroyNotify)
	clistSetRowHeight           func(c *CList, height uint)
	clistSetRowStyle            func(c *CList, row, style *Style)
	clistSetSelectable          func(c *CList, row int, selectable T.Gboolean)
	clistSetSelectionMode       func(c *CList, mode SelectionMode)
	clistSetShadowType          func(c *CList, t ShadowType)
	clistSetShift               func(c *CList, row, column, vertical, horizontal int)
	clistSetSortColumn          func(c *CList, column int)
	clistSetSortType            func(c *CList, sortType SortType)
	clistSetText                func(c *CList, row, column int, text string)
	clistSetUseDragIcons        func(c *CList, useIcons T.Gboolean)
	clistSetVadjustment         func(c *CList, adjustment *Adjustment)
	clistSort                   func(c *CList)
	clistSwapRows               func(c *CList, row1, row2 int)
	clistThaw                   func(c *CList)
	clistUndoSelection          func(c *CList)
	clistUnselectAll            func(c *CList)
	clistUnselectRow            func(c *CList, row, column int)
)

func (c *CList) Append(text **T.Gchar) int            { return clistAppend(c, text) }
func (c *CList) Clear()                               { clistClear(c) }
func (c *CList) ColumnsAutosize() int                 { return clistColumnsAutosize(c) }
func (c *CList) ColumnTitleActive(column int)         { clistColumnTitleActive(c, column) }
func (c *CList) ColumnTitlePassive(column int)        { clistColumnTitlePassive(c, column) }
func (c *CList) ColumnTitlesActive()                  { clistColumnTitlesActive(c) }
func (c *CList) ColumnTitlesHide()                    { clistColumnTitlesHide(c) }
func (c *CList) ColumnTitlesPassive()                 { clistColumnTitlesPassive(c) }
func (c *CList) ColumnTitlesShow()                    { clistColumnTitlesShow(c) }
func (c *CList) FindRowFromData(data T.Gpointer) int  { return clistFindRowFromData(c, data) }
func (c *CList) Freeze()                              { clistFreeze(c) }
func (c *CList) GetCellStyle(row, column int) *Style  { return clistGetCellStyle(c, row, column) }
func (c *CList) GetCellType(row, column int) CellType { return clistGetCellType(c, row, column) }
func (c *CList) GetColumnTitle(column int) string     { return clistGetColumnTitle(c, column) }
func (c *CList) GetColumnWidget(column int) *Widget   { return clistGetColumnWidget(c, column) }
func (c *CList) GetHadjustment() *Adjustment          { return clistGetHadjustment(c) }
func (c *CList) GetPixmap(row, column int, pixmap, mask **T.GdkBitmap) int {
	return clistGetPixmap(c, row, column, pixmap, mask)
}
func (c *CList) GetPixtext(row, column int, text **T.Gchar, spacing *uint8, pixmap, mask **T.GdkBitmap) int {
	return clistGetPixtext(c, row, column, text, spacing, pixmap, mask)
}
func (c *CList) GetRowData(row int) T.Gpointer    { return clistGetRowData(c, row) }
func (c *CList) GetRowStyle(row int) *Style       { return clistGetRowStyle(c, row) }
func (c *CList) GetSelectable(row int) T.Gboolean { return clistGetSelectable(c, row) }
func (c *CList) GetSelectionInfo(x, y int, row, column *int) int {
	return clistGetSelectionInfo(c, x, y, row, column)
}
func (c *CList) GetText(row, column int, text **T.Gchar) int {
	return clistGetText(c, row, column, text)
}
func (c *CList) GetVadjustment() *Adjustment        { return clistGetVadjustment(c) }
func (c *CList) Insert(row int, text **T.Gchar) int { return clistInsert(c, row, text) }
func (c *CList) Moveto(row, column int, rowAlign, colAlign float32) {
	clistMoveto(c, row, column, rowAlign, colAlign)
}
func (c *CList) OptimalColumnWidth(column int) int     { return clistOptimalColumnWidth(c, column) }
func (c *CList) Prepend(text **T.Gchar) int            { return clistPrepend(c, text) }
func (c *CList) Remove(row int)                        { clistRemove(c, row) }
func (c *CList) RowIsVisible(row int) Visibility       { return clistRowIsVisible(c, row) }
func (c *CList) RowMove(sourceRow, destRow int)        { clistRowMove(c, sourceRow, destRow) }
func (c *CList) SelectAll()                            { clistSelectAll(c) }
func (c *CList) SelectRow(row, column int)             { clistSelectRow(c, row, column) }
func (c *CList) SetAutoSort(autoSort T.Gboolean)       { clistSetAutoSort(c, autoSort) }
func (c *CList) SetBackground(row int, color *D.Color) { clistSetBackground(c, row, color) }
func (c *CList) SetButtonActions(button uint, buttonActions uint8) {
	clistSetButtonActions(c, button, buttonActions)
}
func (c *CList) SetCellStyle(row, column int, style *Style) {
	clistSetCellStyle(c, row, column, style)
}
func (c *CList) SetColumnAutoResize(column int, autoResize T.Gboolean) {
	clistSetColumnAutoResize(c, column, autoResize)
}
func (c *CList) SetColumnJustification(column int, justification Justification) {
	clistSetColumnJustification(c, column, justification)
}
func (c *CList) SetColumnMaxWidth(column, maxWidth int) { clistSetColumnMaxWidth(c, column, maxWidth) }
func (c *CList) SetColumnMinWidth(column, minWidth int) { clistSetColumnMinWidth(c, column, minWidth) }
func (c *CList) SetColumnResizeable(column int, resizeable T.Gboolean) {
	clistSetColumnResizeable(c, column, resizeable)
}
func (c *CList) SetColumnTitle(column int, title string) { clistSetColumnTitle(c, column, title) }
func (c *CList) SetColumnVisibility(column int, visible T.Gboolean) {
	clistSetColumnVisibility(c, column, visible)
}
func (c *CList) SetColumnWidget(column int, widget *Widget) { clistSetColumnWidget(c, column, widget) }
func (c *CList) SetColumnWidth(column, width int)           { clistSetColumnWidth(c, column, width) }
func (c *CList) SetCompareFunc(cmpFunc CListCompareFunc)    { clistSetCompareFunc(c, cmpFunc) }
func (c *CList) SetForeground(row int, color *D.Color)      { clistSetForeground(c, row, color) }
func (c *CList) SetHadjustment(adjustment *Adjustment)      { clistSetHadjustment(c, adjustment) }
func (c *CList) SetPixmap(row, column int, pixmap *D.Pixmap, mask *T.GdkBitmap) {
	clistSetPixmap(c, row, column, pixmap, mask)
}
func (c *CList) SetPixtext(row, column int, text string, spacing uint8, pixmap, mask *T.GdkBitmap) {
	clistSetPixtext(c, row, column, text, spacing, pixmap, mask)
}
func (c *CList) SetReorderable(reorderable T.Gboolean) { clistSetReorderable(c, reorderable) }
func (c *CList) SetRowData(row int, data T.Gpointer)   { clistSetRowData(c, row, data) }
func (c *CList) SetRowDataFull(row int, data T.Gpointer, destroy T.GDestroyNotify) {
	clistSetRowDataFull(c, row, data, destroy)
}
func (c *CList) SetRowHeight(height uint)                     { clistSetRowHeight(c, height) }
func (c *CList) SetRowStyle(row, style *Style)                { clistSetRowStyle(c, row, style) }
func (c *CList) SetSelectable(row int, selectable T.Gboolean) { clistSetSelectable(c, row, selectable) }
func (c *CList) SetSelectionMode(mode SelectionMode)          { clistSetSelectionMode(c, mode) }
func (c *CList) SetShadowType(t ShadowType)                   { clistSetShadowType(c, t) }
func (c *CList) SetShift(row, column, vertical, horizontal int) {
	clistSetShift(c, row, column, vertical, horizontal)
}
func (c *CList) SetSortColumn(column int)              { clistSetSortColumn(c, column) }
func (c *CList) SetSortType(sortType SortType)         { clistSetSortType(c, sortType) }
func (c *CList) SetText(row, column int, text string)  { clistSetText(c, row, column, text) }
func (c *CList) SetUseDragIcons(useIcons T.Gboolean)   { clistSetUseDragIcons(c, useIcons) }
func (c *CList) SetVadjustment(adjustment *Adjustment) { clistSetVadjustment(c, adjustment) }
func (c *CList) Sort()                                 { clistSort(c) }
func (c *CList) SwapRows(row1, row2 int)               { clistSwapRows(c, row1, row2) }
func (c *CList) Thaw()                                 { clistThaw(c) }
func (c *CList) UndoSelection()                        { clistUndoSelection(c) }
func (c *CList) UnselectAll()                          { clistUnselectAll(c) }
func (c *CList) UnselectRow(row, column int)           { clistUnselectRow(c, row, column) }

type ColorButton struct {
	Button Button
	_      *struct{}
}

var (
	ColorButtonGetType      func() T.GType
	ColorButtonNew          func() *Widget
	ColorButtonNewWithColor func(color *D.Color) *Widget
)

var (
	colorButtonGetAlpha    func(c *ColorButton) uint16
	colorButtonGetColor    func(c *ColorButton, color *D.Color)
	colorButtonGetTitle    func(c *ColorButton) string
	colorButtonGetUseAlpha func(c *ColorButton) T.Gboolean
	colorButtonSetAlpha    func(c *ColorButton, alpha uint16)
	colorButtonSetColor    func(c *ColorButton, color *D.Color)
	colorButtonSetTitle    func(c *ColorButton, title string)
	colorButtonSetUseAlpha func(c *ColorButton, useAlpha T.Gboolean)
)

func (c *ColorButton) GetAlpha() uint16                { return colorButtonGetAlpha(c) }
func (c *ColorButton) GetColor(color *D.Color)         { colorButtonGetColor(c, color) }
func (c *ColorButton) GetTitle() string                { return colorButtonGetTitle(c) }
func (c *ColorButton) GetUseAlpha() T.Gboolean         { return colorButtonGetUseAlpha(c) }
func (c *ColorButton) SetAlpha(alpha uint16)           { colorButtonSetAlpha(c, alpha) }
func (c *ColorButton) SetColor(color *D.Color)         { colorButtonSetColor(c, color) }
func (c *ColorButton) SetTitle(title string)           { colorButtonSetTitle(c, title) }
func (c *ColorButton) SetUseAlpha(useAlpha T.Gboolean) { colorButtonSetUseAlpha(c, useAlpha) }

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
	ColorSelectionGetType func() T.GType
	ColorSelectionNew     func() *Widget

	ColorSelectionPaletteFromString func(str string, colors **D.Color, nColors *int) T.Gboolean
	ColorSelectionPaletteToString   func(colors *D.Color, nColors int) string

	ColorSelectionSetChangePaletteHook           func(f ColorSelectionChangePaletteFunc) ColorSelectionChangePaletteFunc
	ColorSelectionSetChangePaletteWithScreenHook func(f ColorSelectionChangePaletteWithScreenFunc) ColorSelectionChangePaletteWithScreenFunc
)

var (
	colorSelectionGetColor             func(c *ColorSelection, color *float64)
	colorSelectionGetCurrentAlpha      func(c *ColorSelection) uint16
	colorSelectionGetCurrentColor      func(c *ColorSelection, color *D.Color)
	colorSelectionGetHasOpacityControl func(c *ColorSelection) T.Gboolean
	colorSelectionGetHasPalette        func(c *ColorSelection) T.Gboolean
	colorSelectionGetPreviousAlpha     func(c *ColorSelection) uint16
	colorSelectionGetPreviousColor     func(c *ColorSelection, color *D.Color)
	colorSelectionIsAdjusting          func(c *ColorSelection) T.Gboolean
	colorSelectionSetColor             func(c *ColorSelection, color *float64)
	colorSelectionSetCurrentAlpha      func(c *ColorSelection, alpha uint16)
	colorSelectionSetCurrentColor      func(c *ColorSelection, color *D.Color)
	colorSelectionSetHasOpacityControl func(c *ColorSelection, hasOpacity T.Gboolean)
	colorSelectionSetHasPalette        func(c *ColorSelection, hasPalette T.Gboolean)
	colorSelectionSetPreviousAlpha     func(c *ColorSelection, alpha uint16)
	colorSelectionSetPreviousColor     func(c *ColorSelection, color *D.Color)
	colorSelectionSetUpdatePolicy      func(c *ColorSelection, policy UpdateType)
)

func (c *ColorSelection) GetColor(color *float64)        { colorSelectionGetColor(c, color) }
func (c *ColorSelection) GetCurrentAlpha() uint16        { return colorSelectionGetCurrentAlpha(c) }
func (c *ColorSelection) GetCurrentColor(color *D.Color) { colorSelectionGetCurrentColor(c, color) }
func (c *ColorSelection) GetHasOpacityControl() T.Gboolean {
	return colorSelectionGetHasOpacityControl(c)
}
func (c *ColorSelection) GetHasPalette() T.Gboolean       { return colorSelectionGetHasPalette(c) }
func (c *ColorSelection) GetPreviousAlpha() uint16        { return colorSelectionGetPreviousAlpha(c) }
func (c *ColorSelection) GetPreviousColor(color *D.Color) { colorSelectionGetPreviousColor(c, color) }
func (c *ColorSelection) IsAdjusting() T.Gboolean         { return colorSelectionIsAdjusting(c) }
func (c *ColorSelection) SetColor(color *float64)         { colorSelectionSetColor(c, color) }
func (c *ColorSelection) SetCurrentAlpha(alpha uint16)    { colorSelectionSetCurrentAlpha(c, alpha) }
func (c *ColorSelection) SetCurrentColor(color *D.Color)  { colorSelectionSetCurrentColor(c, color) }
func (c *ColorSelection) SetHasOpacityControl(hasOpacity T.Gboolean) {
	colorSelectionSetHasOpacityControl(c, hasOpacity)
}
func (c *ColorSelection) SetHasPalette(hasPalette T.Gboolean) {
	colorSelectionSetHasPalette(c, hasPalette)
}
func (c *ColorSelection) SetPreviousAlpha(alpha uint16)   { colorSelectionSetPreviousAlpha(c, alpha) }
func (c *ColorSelection) SetPreviousColor(color *D.Color) { colorSelectionSetPreviousColor(c, color) }
func (c *ColorSelection) SetUpdatePolicy(policy UpdateType) {
	colorSelectionSetUpdatePolicy(c, policy)
}

type ColorSelectionDialog struct {
	Parent       Dialog
	Colorsel     *Widget
	OkButton     *Widget
	CancelButton *Widget
	HelpButton   *Widget
}

var (
	ColorSelectionDialogGetType func() T.GType
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
	ComboGetType func() T.GType
	ComboNew     func() *Widget
)

var (
	comboSetValueInList     func(c *Combo, val T.Gboolean, okIfEmpty T.Gboolean)
	comboSetUseArrows       func(c *Combo, val T.Gboolean)
	comboSetUseArrowsAlways func(c *Combo, val T.Gboolean)
	comboSetCaseSensitive   func(c *Combo, val T.Gboolean)
	comboSetItemString      func(c *Combo, item *Item, itemValue string)
	comboSetPopdownStrings  func(c *Combo, strings *T.GList)
	comboDisableActivate    func(c *Combo)
)

func (c *Combo) DisableActivate()                           { comboDisableActivate(c) }
func (c *Combo) SetCaseSensitive(val T.Gboolean)            { comboSetCaseSensitive(c, val) }
func (c *Combo) SetItemString(item *Item, itemValue string) { comboSetItemString(c, item, itemValue) }
func (c *Combo) SetPopdownStrings(strings *T.GList)         { comboSetPopdownStrings(c, strings) }
func (c *Combo) SetUseArrows(val T.Gboolean)                { comboSetUseArrows(c, val) }
func (c *Combo) SetUseArrowsAlways(val T.Gboolean)          { comboSetUseArrowsAlways(c, val) }
func (c *Combo) SetValueInList(val T.Gboolean, okIfEmpty T.Gboolean) {
	comboSetValueInList(c, val, okIfEmpty)
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
	ComboBoxGetType              func() T.GType
	ComboBoxNew                  func() *Widget
	ComboBoxNewWithEntry         func() *Widget
	ComboBoxNewWithModel         func(model *TreeModel) *Widget
	ComboBoxNewWithModelAndEntry func(model *TreeModel) *Widget

	ComboBoxTextGetType      func() T.GType
	ComboBoxTextNew          func() *Widget
	ComboBoxTextNewWithEntry func() *Widget

	ComboBoxNewText func() *Widget

	ComboBoxEntryGetType      func() T.GType
	ComboBoxEntryNew          func() *Widget
	ComboBoxEntryNewWithModel func(model *TreeModel, textColumn int) *Widget
	ComboBoxEntryNewText      func() *Widget
)

var (
	comboBoxAppendText           func(c *ComboBox, text string)
	comboBoxGetActive            func(c *ComboBox) int
	comboBoxGetActiveIter        func(c *ComboBox, iter *TreeIter) T.Gboolean
	comboBoxGetActiveText        func(c *ComboBox) string
	comboBoxGetAddTearoffs       func(c *ComboBox) T.Gboolean
	comboBoxGetButtonSensitivity func(c *ComboBox) SensitivityType
	comboBoxGetColumnSpanColumn  func(c *ComboBox) int
	comboBoxGetEntryTextColumn   func(c *ComboBox) int
	comboBoxGetFocusOnClick      func(c *ComboBox) T.Gboolean
	comboBoxGetHasEntry          func(c *ComboBox) T.Gboolean
	comboBoxGetModel             func(c *ComboBox) *TreeModel
	comboBoxGetPopupAccessible   func(c *ComboBox) *T.AtkObject
	comboBoxGetRowSeparatorFunc  func(c *ComboBox) TreeViewRowSeparatorFunc
	comboBoxGetRowSpanColumn     func(c *ComboBox) int
	comboBoxGetTitle             func(c *ComboBox) string
	comboBoxGetWrapWidth         func(c *ComboBox) int
	comboBoxInsertText           func(c *ComboBox, position int, text string)
	comboBoxPopdown              func(c *ComboBox)
	comboBoxPopup                func(c *ComboBox)
	comboBoxPrependText          func(c *ComboBox, text string)
	comboBoxRemoveText           func(c *ComboBox, position int)
	comboBoxSetActive            func(c *ComboBox, index int)
	comboBoxSetActiveIter        func(c *ComboBox, iter *TreeIter)
	comboBoxSetAddTearoffs       func(c *ComboBox, addTearoffs T.Gboolean)
	comboBoxSetButtonSensitivity func(c *ComboBox, sensitivity SensitivityType)
	comboBoxSetColumnSpanColumn  func(c *ComboBox, columnSpan int)
	comboBoxSetEntryTextColumn   func(c *ComboBox, textColumn int)
	comboBoxSetFocusOnClick      func(c *ComboBox, focusOnClick T.Gboolean)
	comboBoxSetModel             func(c *ComboBox, model *TreeModel)
	comboBoxSetRowSeparatorFunc  func(c *ComboBox, f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify)
	comboBoxSetRowSpanColumn     func(c *ComboBox, rowSpan int)
	comboBoxSetTitle             func(c *ComboBox, title string)
	comboBoxSetWrapWidth         func(c *ComboBox, width int)
)

func (c *ComboBox) AppendText(text string) { comboBoxAppendText(c, text) }
func (c *ComboBox) GetActive() int         { return comboBoxGetActive(c) }
func (c *ComboBox) GetActiveIter(iter *TreeIter) T.Gboolean {
	return comboBoxGetActiveIter(c, iter)
}
func (c *ComboBox) GetActiveText() string                 { return comboBoxGetActiveText(c) }
func (c *ComboBox) GetAddTearoffs() T.Gboolean            { return comboBoxGetAddTearoffs(c) }
func (c *ComboBox) GetButtonSensitivity() SensitivityType { return comboBoxGetButtonSensitivity(c) }
func (c *ComboBox) GetColumnSpanColumn() int              { return comboBoxGetColumnSpanColumn(c) }
func (c *ComboBox) GetEntryTextColumn() int               { return comboBoxGetEntryTextColumn(c) }
func (c *ComboBox) GetFocusOnClick() T.Gboolean           { return comboBoxGetFocusOnClick(c) }
func (c *ComboBox) GetHasEntry() T.Gboolean               { return comboBoxGetHasEntry(c) }
func (c *ComboBox) GetModel() *TreeModel                  { return comboBoxGetModel(c) }
func (c *ComboBox) GetPopupAccessible() *T.AtkObject      { return comboBoxGetPopupAccessible(c) }
func (c *ComboBox) GetRowSeparatorFunc() TreeViewRowSeparatorFunc {
	return comboBoxGetRowSeparatorFunc(c)
}
func (c *ComboBox) GetRowSpanColumn() int                 { return comboBoxGetRowSpanColumn(c) }
func (c *ComboBox) GetTitle() string                      { return comboBoxGetTitle(c) }
func (c *ComboBox) GetWrapWidth() int                     { return comboBoxGetWrapWidth(c) }
func (c *ComboBox) InsertText(position int, text string)  { comboBoxInsertText(c, position, text) }
func (c *ComboBox) Popdown()                              { comboBoxPopdown(c) }
func (c *ComboBox) Popup()                                { comboBoxPopup(c) }
func (c *ComboBox) PrependText(text string)               { comboBoxPrependText(c, text) }
func (c *ComboBox) RemoveText(position int)               { comboBoxRemoveText(c, position) }
func (c *ComboBox) SetActive(index int)                   { comboBoxSetActive(c, index) }
func (c *ComboBox) SetActiveIter(iter *TreeIter)          { comboBoxSetActiveIter(c, iter) }
func (c *ComboBox) SetAddTearoffs(addTearoffs T.Gboolean) { comboBoxSetAddTearoffs(c, addTearoffs) }
func (c *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	comboBoxSetButtonSensitivity(c, sensitivity)
}
func (c *ComboBox) SetColumnSpanColumn(columnSpan int)      { comboBoxSetColumnSpanColumn(c, columnSpan) }
func (c *ComboBox) SetEntryTextColumn(textColumn int)       { comboBoxSetEntryTextColumn(c, textColumn) }
func (c *ComboBox) SetFocusOnClick(focusOnClick T.Gboolean) { comboBoxSetFocusOnClick(c, focusOnClick) }
func (c *ComboBox) SetModel(model *TreeModel)               { comboBoxSetModel(c, model) }
func (c *ComboBox) SetRowSeparatorFunc(f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	comboBoxSetRowSeparatorFunc(c, f, data, destroy)
}
func (c *ComboBox) SetRowSpanColumn(rowSpan int) { comboBoxSetRowSpanColumn(c, rowSpan) }
func (c *ComboBox) SetTitle(title string)        { comboBoxSetTitle(c, title) }
func (c *ComboBox) SetWrapWidth(width int)       { comboBoxSetWrapWidth(c, width) }

var (
	comboBoxTextAppendText    func(c *ComboBoxText, text string)
	comboBoxTextGetActiveText func(c *ComboBoxText) string
	comboBoxTextInsertText    func(c *ComboBoxText, position int, text string)
	comboBoxTextPrependText   func(c *ComboBoxText, text string)
	comboBoxTextRemove        func(c *ComboBoxText, position int)
)

func (c *ComboBoxText) AppendText(text string) { comboBoxTextAppendText(c, text) }
func (c *ComboBoxText) GetActiveText() string  { return comboBoxTextGetActiveText(c) }
func (c *ComboBoxText) InsertText(position int, text string) {
	comboBoxTextInsertText(c, position, text)
}
func (c *ComboBoxText) PrependText(text string) { comboBoxTextPrependText(c, text) }
func (c *ComboBoxText) Remove(position int)     { comboBoxTextRemove(c, position) }

var (
	comboBoxEntrySetTextColumn func(e *ComboBoxEntry, textColumn int)
	comboBoxEntryGetTextColumn func(e *ComboBoxEntry) int
)

func (e *ComboBoxEntry) GetTextColumn() int           { return comboBoxEntryGetTextColumn(e) }
func (e *ComboBoxEntry) SetTextColumn(textColumn int) { comboBoxEntrySetTextColumn(e, textColumn) }

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
	ContainerGetType func() T.GType

	containerAdd                  func(c *Container, widget *Widget)
	containerAddWithProperties    func(c *Container, widget *Widget, firstPropName string, v ...VArg)
	containerCheckResize          func(c *Container)
	containerChildGet             func(c *Container, child *Widget, firstPropName string, v ...VArg)
	containerChildGetProperty     func(c *Container, child *Widget, propertyName string, value *T.GValue)
	containerChildGetValist       func(c *Container, child *Widget, firstPropertyName string, varArgs T.VaList)
	containerChildSet             func(c *Container, child *Widget, firstPropName string, v ...VArg)
	containerChildSetProperty     func(c *Container, child *Widget, propertyName string, value *T.GValue)
	containerChildSetValist       func(c *Container, child *Widget, firstPropertyName string, varArgs T.VaList)
	containerChildType            func(c *Container) T.GType
	containerForall               func(c *Container, callback Callback, callbackData T.Gpointer)
	containerForeach              func(c *Container, callback Callback, callbackData T.Gpointer)
	containerForeachFull          func(c *Container, callback Callback, marshal CallbackMarshal, callbackData T.Gpointer, notify T.GDestroyNotify)
	containerGetBorderWidth       func(c *Container) uint
	containerGetChildren          func(c *Container) *T.GList
	containerGetFocusChain        func(c *Container, focusableWidgets **T.GList) T.Gboolean
	containerGetFocusChild        func(c *Container) *Widget
	containerGetFocusHadjustment  func(c *Container) *Adjustment
	containerGetFocusVadjustment  func(c *Container) *Adjustment
	containerGetResizeMode        func(c *Container) ResizeMode
	containerPropagateExpose      func(c *Container, child *Widget, event *D.EventExpose)
	containerRemove               func(c *Container, widget *Widget)
	containerResizeChildren       func(c *Container)
	containerSetBorderWidth       func(c *Container, borderWidth uint)
	containerSetFocusChain        func(c *Container, focusableWidgets *T.GList)
	containerSetFocusChild        func(c *Container, child *Widget)
	containerSetFocusHadjustment  func(c *Container, adjustment *Adjustment)
	containerSetFocusVadjustment  func(c *Container, adjustment *Adjustment)
	containerSetReallocateRedraws func(c *Container, needsRedraws T.Gboolean)
	containerSetResizeMode        func(c *Container, resizeMode ResizeMode)
	containerUnsetFocusChain      func(c *Container)
)

func (c *Container) Add(widget *Widget) { containerAdd(c, widget) }
func (c *Container) AddWithProperties(widget *Widget, firstPropName string, v ...VArg) {
	containerAddWithProperties(c, widget, firstPropName, v)
}
func (c *Container) CheckResize() { containerCheckResize(c) }
func (c *Container) ChildGet(child *Widget, firstPropName string, v ...VArg) {
	containerChildGet(c, child, firstPropName, v)
}
func (c *Container) ChildGetProperty(child *Widget, propertyName string, value *T.GValue) {
	containerChildGetProperty(c, child, propertyName, value)
}
func (c *Container) ChildGetValist(child *Widget, firstPropertyName string, varArgs T.VaList) {
	containerChildGetValist(c, child, firstPropertyName, varArgs)
}
func (c *Container) ChildSet(child *Widget, firstPropName string, v ...VArg) {
	containerChildSet(c, child, firstPropName, v)
}
func (c *Container) ChildSetProperty(child *Widget, propertyName string, value *T.GValue) {
	containerChildSetProperty(c, child, propertyName, value)
}
func (c *Container) ChildSetValist(child *Widget, firstPropertyName string, varArgs T.VaList) {
	containerChildSetValist(c, child, firstPropertyName, varArgs)
}
func (c *Container) ChildType() T.GType { return containerChildType(c) }
func (c *Container) Forall(callback Callback, callbackData T.Gpointer) {
	containerForall(c, callback, callbackData)
}
func (c *Container) Foreach(callback Callback, callbackData T.Gpointer) {
	containerForeach(c, callback, callbackData)
}
func (c *Container) ForeachFull(callback Callback, marshal CallbackMarshal, callbackData T.Gpointer, notify T.GDestroyNotify) {
	containerForeachFull(c, callback, marshal, callbackData, notify)
}
func (c *Container) GetBorderWidth() uint  { return containerGetBorderWidth(c) }
func (c *Container) GetChildren() *T.GList { return containerGetChildren(c) }
func (c *Container) GetFocusChain(focusableWidgets **T.GList) T.Gboolean {
	return containerGetFocusChain(c, focusableWidgets)
}
func (c *Container) GetFocusChild() *Widget           { return containerGetFocusChild(c) }
func (c *Container) GetFocusHadjustment() *Adjustment { return containerGetFocusHadjustment(c) }
func (c *Container) GetFocusVadjustment() *Adjustment { return containerGetFocusVadjustment(c) }
func (c *Container) GetResizeMode() ResizeMode        { return containerGetResizeMode(c) }
func (c *Container) PropagateExpose(child *Widget, event *D.EventExpose) {
	containerPropagateExpose(c, child, event)
}
func (c *Container) Remove(widget *Widget)           { containerRemove(c, widget) }
func (c *Container) ResizeChildren()                 { containerResizeChildren(c) }
func (c *Container) SetBorderWidth(borderWidth uint) { containerSetBorderWidth(c, borderWidth) }
func (c *Container) SetFocusChain(focusableWidgets *T.GList) {
	containerSetFocusChain(c, focusableWidgets)
}
func (c *Container) SetFocusChild(child *Widget) { containerSetFocusChild(c, child) }
func (c *Container) SetFocusHadjustment(adjustment *Adjustment) {
	containerSetFocusHadjustment(c, adjustment)
}
func (c *Container) SetFocusVadjustment(adjustment *Adjustment) {
	containerSetFocusVadjustment(c, adjustment)
}
func (c *Container) SetReallocateRedraws(needsRedraws T.Gboolean) {
	containerSetReallocateRedraws(c, needsRedraws)
}
func (c *Container) SetResizeMode(resizeMode ResizeMode) { containerSetResizeMode(c, resizeMode) }
func (c *Container) UnsetFocusChain()                    { containerUnsetFocusChain(c) }

type ContainerClass struct {
	ParentClass      WidgetClass
	Add              func(c *Container, w *Widget)
	Remove           func(c *Container, w *Widget)
	CheckResize      func(c *Container)
	Forall           func(c *Container, includeInternals T.Gboolean, callback Callback, callbackData T.Gpointer)
	SetFocusChild    func(c *Container, w *Widget)
	ChildType        func(c *Container) T.GType
	CompositeName    func(c *Container, child *Widget) *T.Gchar
	SetChildProperty func(c *Container, child *Widget, propertyId uint, value *T.GValue, pspec *T.GParamSpec)
	GetChildProperty func(c *Container, child *Widget, propertyId uint, value *T.GValue, pspec *T.GParamSpec)
	_, _, _, _       func()
}

var (
	ContainerClassFindChildProperty   func(cclass *T.GObjectClass, propertyName string) *T.GParamSpec
	ContainerClassListChildProperties func(cclass *T.GObjectClass, nProperties *uint) **T.GParamSpec

	containerClassInstallChildProperty func(cclass *ContainerClass, propertyId uint, pspec *T.GParamSpec)
)

func (c *ContainerClass) InstallChildProperty(propertyId uint, pspec *T.GParamSpec) {
	containerClassInstallChildProperty(c, propertyId, pspec)
}

type CornerType Enum

const (
	CORNER_TOP_LEFT CornerType = iota
	CORNER_BOTTOM_LEFT
	CORNER_TOP_RIGHT
	CORNER_BOTTOM_RIGHT
)

var CornerTypeGetType func() T.GType

type (
	CTree struct {
		Clist       CList
		LinesGc     *T.GdkGC
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
		sourceNode, newParent, newSibling *CTreeNode) T.Gboolean

	CTreeGNodeFunc func(ctree *CTree, depth uint,
		gnode *T.GNode, cnode *CTreeNode,
		data T.Gpointer) T.Gboolean
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
	CtreeGetType       func() T.GType
	CtreeNew           func(columns int, treeColumn int) *Widget
	CtreeNewWithTitles func(columns int, treeColumn int, titles **T.Gchar) *Widget

	CtreeNodeGetType func() T.GType
	CtreePosGetType  func() T.GType

	CtreeExpanderStyleGetType func() T.GType
	CtreeExpansionTypeGetType func() T.GType
	CtreeLineStyleGetType     func() T.GType
)

var (
	ctreeCollapse                 func(c *CTree, node *CTreeNode)
	ctreeCollapseRecursive        func(c *CTree, node *CTreeNode)
	ctreeCollapseToDepth          func(c *CTree, node *CTreeNode, depth int)
	ctreeExpand                   func(c *CTree, node *CTreeNode)
	ctreeExpandRecursive          func(c *CTree, node *CTreeNode)
	ctreeExpandToDepth            func(c *CTree, node *CTreeNode, depth int)
	ctreeExportToGnode            func(c *CTree, parent, sibling *T.GNode, node *CTreeNode, f CTreeGNodeFunc, data T.Gpointer) *T.GNode
	ctreeFind                     func(c *CTree, node *CTreeNode, child *CTreeNode) T.Gboolean
	ctreeFindAllByRowData         func(c *CTree, node *CTreeNode, data T.Gpointer) *T.GList
	ctreeFindAllByRowDataCustom   func(c *CTree, node *CTreeNode, data T.Gpointer, f T.GCompareFunc) *T.GList
	ctreeFindByRowData            func(c *CTree, node *CTreeNode, data T.Gpointer) *CTreeNode
	ctreeFindByRowDataCustom      func(c *CTree, node *CTreeNode, data T.Gpointer, f T.GCompareFunc) *CTreeNode
	ctreeFindNodePtr              func(c *CTree, ctreeRow *CTreeRow) *CTreeNode
	ctreeGetNodeInfo              func(c *CTree, node *CTreeNode, text **T.Gchar, spacing *uint8, pixmapClosed **D.Pixmap, maskClosed **T.GdkBitmap, pixmapOpened **D.Pixmap, maskOpened **T.GdkBitmap, isLeaf, expanded *T.Gboolean) T.Gboolean
	ctreeInsertGnode              func(c *CTree, parent *CTreeNode, sibling *CTreeNode, gnode *T.GNode, f CTreeGNodeFunc, data T.Gpointer) *CTreeNode
	ctreeInsertNode               func(c *CTree, parent *CTreeNode, sibling *CTreeNode, text **T.Gchar, spacing uint8, pixmapClosed *D.Pixmap, maskClosed *T.GdkBitmap, pixmapOpened *D.Pixmap, maskOpened *T.GdkBitmap, isLeaf T.Gboolean, expanded T.Gboolean) *CTreeNode
	ctreeIsAncestor               func(c *CTree, node *CTreeNode, child *CTreeNode) T.Gboolean
	ctreeIsHotSpot                func(c *CTree, x int, y int) T.Gboolean
	ctreeIsViewable               func(c *CTree, node *CTreeNode) T.Gboolean
	ctreeLast                     func(c *CTree, node *CTreeNode) *CTreeNode
	ctreeMove                     func(c *CTree, node *CTreeNode, newParent *CTreeNode, newSibling *CTreeNode)
	ctreeNodeGetCellStyle         func(c *CTree, node *CTreeNode, column int) *Style
	ctreeNodeGetCellType          func(c *CTree, node *CTreeNode, column int) CellType
	ctreeNodeGetPixmap            func(c *CTree, node *CTreeNode, column int, pixmap **D.Pixmap, mask **T.GdkBitmap) T.Gboolean
	ctreeNodeGetPixtext           func(c *CTree, node *CTreeNode, column int, text **T.Gchar, spacing *uint8, pixmap **D.Pixmap, mask **T.GdkBitmap) T.Gboolean
	ctreeNodeGetRowData           func(c *CTree, node *CTreeNode) T.Gpointer
	ctreeNodeGetRowStyle          func(c *CTree, node *CTreeNode) *Style
	ctreeNodeGetSelectable        func(c *CTree, node *CTreeNode) T.Gboolean
	ctreeNodeGetText              func(c *CTree, node *CTreeNode, column int, text **T.Gchar) T.Gboolean
	ctreeNodeIsVisible            func(c *CTree, node *CTreeNode) Visibility
	ctreeNodeMoveto               func(c *CTree, node *CTreeNode, column int, rowAlign, colAlign float32)
	ctreeNodeNth                  func(c *CTree, row uint) *CTreeNode
	ctreeNodeSetBackground        func(c *CTree, node *CTreeNode, color *D.Color)
	ctreeNodeSetCellStyle         func(c *CTree, node *CTreeNode, column int, style *Style)
	ctreeNodeSetForeground        func(c *CTree, node *CTreeNode, color *D.Color)
	ctreeNodeSetPixmap            func(c *CTree, node *CTreeNode, column int, pixmap *D.Pixmap, mask *T.GdkBitmap)
	ctreeNodeSetPixtext           func(c *CTree, node *CTreeNode, column int, text string, spacing uint8, pixmap *D.Pixmap, mask *T.GdkBitmap)
	ctreeNodeSetRowData           func(c *CTree, node *CTreeNode, data T.Gpointer)
	ctreeNodeSetRowDataFull       func(c *CTree, node *CTreeNode, data T.Gpointer, destroy T.GDestroyNotify)
	ctreeNodeSetRowStyle          func(c *CTree, node *CTreeNode, style *Style)
	ctreeNodeSetSelectable        func(c *CTree, node *CTreeNode, selectable T.Gboolean)
	ctreeNodeSetShift             func(c *CTree, node *CTreeNode, column, vertical, horizontal int)
	ctreeNodeSetText              func(c *CTree, node *CTreeNode, column int, text string)
	ctreePostRecursive            func(c *CTree, node *CTreeNode, f CTreeFunc, data T.Gpointer)
	ctreePostRecursiveToDepth     func(c *CTree, node *CTreeNode, depth int, f CTreeFunc, data T.Gpointer)
	ctreePreRecursive             func(c *CTree, node *CTreeNode, f CTreeFunc, data T.Gpointer)
	ctreePreRecursiveToDepth      func(c *CTree, node *CTreeNode, depth int, f CTreeFunc, data T.Gpointer)
	ctreeRealSelectRecursive      func(c *CTree, node *CTreeNode, state int)
	ctreeRemoveNode               func(c *CTree, node *CTreeNode)
	ctreeSelect                   func(c *CTree, node *CTreeNode)
	ctreeSelectRecursive          func(c *CTree, node *CTreeNode)
	ctreeSetDragCompareFunc       func(c *CTree, cmpFunc CTreeCompareDragFunc)
	ctreeSetExpanderStyle         func(c *CTree, expanderStyle CTreeExpanderStyle)
	ctreeSetIndent                func(c *CTree, indent int)
	ctreeSetLineStyle             func(c *CTree, lineStyle CTreeLineStyle)
	ctreeSetNodeInfo              func(c *CTree, node *CTreeNode, text string, spacing uint8, pixmapClosed *D.Pixmap, maskClosed *T.GdkBitmap, pixmapOpened *D.Pixmap, maskOpened *T.GdkBitmap, isLeaf, expanded T.Gboolean)
	ctreeSetShowStub              func(c *CTree, showStub T.Gboolean)
	ctreeSetSpacing               func(c *CTree, spacing int)
	ctreeSortNode                 func(c *CTree, node *CTreeNode)
	ctreeSortRecursive            func(c *CTree, node *CTreeNode)
	ctreeToggleExpansion          func(c *CTree, node *CTreeNode)
	ctreeToggleExpansionRecursive func(c *CTree, node *CTreeNode)
	ctreeUnselect                 func(c *CTree, node *CTreeNode)
	ctreeUnselectRecursive        func(c *CTree, node *CTreeNode)
)

func (c *CTree) Collapse(node *CTreeNode)                   { ctreeCollapse(c, node) }
func (c *CTree) CollapseRecursive(node *CTreeNode)          { ctreeCollapseRecursive(c, node) }
func (c *CTree) CollapseToDepth(node *CTreeNode, depth int) { ctreeCollapseToDepth(c, node, depth) }
func (c *CTree) Expand(node *CTreeNode)                     { ctreeExpand(c, node) }
func (c *CTree) ExpandRecursive(node *CTreeNode)            { ctreeExpandRecursive(c, node) }
func (c *CTree) ExpandToDepth(node *CTreeNode, depth int)   { ctreeExpandToDepth(c, node, depth) }
func (c *CTree) ExportToGnode(parent *T.GNode, sibling *T.GNode, node *CTreeNode, f CTreeGNodeFunc, data T.Gpointer) *T.GNode {
	return ctreeExportToGnode(c, parent, sibling, node, f, data)
}
func (c *CTree) Find(node, child *CTreeNode) T.Gboolean { return ctreeFind(c, node, child) }
func (c *CTree) FindAllByRowData(node *CTreeNode, data T.Gpointer) *T.GList {
	return ctreeFindAllByRowData(c, node, data)
}
func (c *CTree) FindAllByRowDataCustom(node *CTreeNode, data T.Gpointer, f T.GCompareFunc) *T.GList {
	return ctreeFindAllByRowDataCustom(c, node, data, f)
}
func (c *CTree) FindByRowData(node *CTreeNode, data T.Gpointer) *CTreeNode {
	return ctreeFindByRowData(c, node, data)
}
func (c *CTree) FindByRowDataCustom(node *CTreeNode, data T.Gpointer, f T.GCompareFunc) *CTreeNode {
	return ctreeFindByRowDataCustom(c, node, data, f)
}
func (c *CTree) FindNodePtr(ctreeRow *CTreeRow) *CTreeNode { return ctreeFindNodePtr(c, ctreeRow) }
func (c *CTree) GetNodeInfo(node *CTreeNode, text **T.Gchar, spacing *uint8, pixmapClosed **D.Pixmap, maskClosed **T.GdkBitmap, pixmapOpened **D.Pixmap, maskOpened **T.GdkBitmap, isLeaf, expanded *T.Gboolean) T.Gboolean {
	return ctreeGetNodeInfo(c, node, text, spacing, pixmapClosed, maskClosed, pixmapOpened, maskOpened, isLeaf, expanded)
}
func (c *CTree) InsertGnode(parent, sibling *CTreeNode, gnode *T.GNode, f CTreeGNodeFunc, data T.Gpointer) *CTreeNode {
	return ctreeInsertGnode(c, parent, sibling, gnode, f, data)
}
func (c *CTree) InsertNode(parent, sibling *CTreeNode, text **T.Gchar, spacing uint8, pixmapClosed *D.Pixmap, maskClosed *T.GdkBitmap, pixmapOpened *D.Pixmap, maskOpened *T.GdkBitmap, isLeaf, expanded T.Gboolean) *CTreeNode {
	return ctreeInsertNode(c, parent, sibling, text, spacing, pixmapClosed, maskClosed, pixmapOpened, maskOpened, isLeaf, expanded)
}
func (c *CTree) IsAncestor(node, child *CTreeNode) T.Gboolean { return ctreeIsAncestor(c, node, child) }
func (c *CTree) IsHotSpot(x int, y int) T.Gboolean            { return ctreeIsHotSpot(c, x, y) }
func (c *CTree) IsViewable(node *CTreeNode) T.Gboolean        { return ctreeIsViewable(c, node) }
func (c *CTree) Last(node *CTreeNode) *CTreeNode              { return ctreeLast(c, node) }
func (c *CTree) Move(node, newParent, newSibling *CTreeNode) {
	ctreeMove(c, node, newParent, newSibling)
}
func (c *CTree) NodeGetCellStyle(node *CTreeNode, column int) *Style {
	return ctreeNodeGetCellStyle(c, node, column)
}
func (c *CTree) NodeGetCellType(node *CTreeNode, column int) CellType {
	return ctreeNodeGetCellType(c, node, column)
}
func (c *CTree) NodeGetPixmap(node *CTreeNode, column int, pixmap **D.Pixmap, mask **T.GdkBitmap) T.Gboolean {
	return ctreeNodeGetPixmap(c, node, column, pixmap, mask)
}
func (c *CTree) NodeGetPixtext(node *CTreeNode, column int, text **T.Gchar, spacing *uint8, pixmap **D.Pixmap, mask **T.GdkBitmap) T.Gboolean {
	return ctreeNodeGetPixtext(c, node, column, text, spacing, pixmap, mask)
}
func (c *CTree) NodeGetRowData(node *CTreeNode) T.Gpointer    { return ctreeNodeGetRowData(c, node) }
func (c *CTree) NodeGetRowStyle(node *CTreeNode) *Style       { return ctreeNodeGetRowStyle(c, node) }
func (c *CTree) NodeGetSelectable(node *CTreeNode) T.Gboolean { return ctreeNodeGetSelectable(c, node) }
func (c *CTree) NodeGetText(node *CTreeNode, column int, text **T.Gchar) T.Gboolean {
	return ctreeNodeGetText(c, node, column, text)
}
func (c *CTree) NodeIsVisible(node *CTreeNode) Visibility { return ctreeNodeIsVisible(c, node) }
func (c *CTree) NodeMoveto(node *CTreeNode, column int, rowAlign, colAlign float32) {
	ctreeNodeMoveto(c, node, column, rowAlign, colAlign)
}
func (c *CTree) NodeNth(row uint) *CTreeNode { return ctreeNodeNth(c, row) }
func (c *CTree) NodeSetBackground(node *CTreeNode, color *D.Color) {
	ctreeNodeSetBackground(c, node, color)
}
func (c *CTree) NodeSetCellStyle(node *CTreeNode, column int, style *Style) {
	ctreeNodeSetCellStyle(c, node, column, style)
}
func (c *CTree) NodeSetForeground(node *CTreeNode, color *D.Color) {
	ctreeNodeSetForeground(c, node, color)
}
func (c *CTree) NodeSetPixmap(node *CTreeNode, column int, pixmap *D.Pixmap, mask *T.GdkBitmap) {
	ctreeNodeSetPixmap(c, node, column, pixmap, mask)
}
func (c *CTree) NodeSetPixtext(node *CTreeNode, column int, text string, spacing uint8, pixmap *D.Pixmap, mask *T.GdkBitmap) {
	ctreeNodeSetPixtext(c, node, column, text, spacing, pixmap, mask)
}
func (c *CTree) NodeSetRowData(node *CTreeNode, data T.Gpointer) { ctreeNodeSetRowData(c, node, data) }
func (c *CTree) NodeSetRowDataFull(node *CTreeNode, data T.Gpointer, destroy T.GDestroyNotify) {
	ctreeNodeSetRowDataFull(c, node, data, destroy)
}
func (c *CTree) NodeSetRowStyle(node *CTreeNode, style *Style) {
	ctreeNodeSetRowStyle(c, node, style)
}
func (c *CTree) NodeSetSelectable(node *CTreeNode, selectable T.Gboolean) {
	ctreeNodeSetSelectable(c, node, selectable)
}
func (c *CTree) NodeSetShift(node *CTreeNode, column, vertical, horizontal int) {
	ctreeNodeSetShift(c, node, column, vertical, horizontal)
}
func (c *CTree) NodeSetText(node *CTreeNode, column int, text string) {
	ctreeNodeSetText(c, node, column, text)
}
func (c *CTree) PostRecursive(node *CTreeNode, f CTreeFunc, data T.Gpointer) {
	ctreePostRecursive(c, node, f, data)
}
func (c *CTree) PostRecursiveToDepth(node *CTreeNode, depth int, f CTreeFunc, data T.Gpointer) {
	ctreePostRecursiveToDepth(c, node, depth, f, data)
}
func (c *CTree) PreRecursive(node *CTreeNode, f CTreeFunc, data T.Gpointer) {
	ctreePreRecursive(c, node, f, data)
}
func (c *CTree) PreRecursiveToDepth(node *CTreeNode, depth int, f CTreeFunc, data T.Gpointer) {
	ctreePreRecursiveToDepth(c, node, depth, f, data)
}
func (c *CTree) RealSelectRecursive(node *CTreeNode, state int) {
	ctreeRealSelectRecursive(c, node, state)
}
func (c *CTree) RemoveNode(node *CTreeNode)                      { ctreeRemoveNode(c, node) }
func (c *CTree) Select(node *CTreeNode)                          { ctreeSelect(c, node) }
func (c *CTree) SelectRecursive(node *CTreeNode)                 { ctreeSelectRecursive(c, node) }
func (c *CTree) SetDragCompareFunc(cmpFunc CTreeCompareDragFunc) { ctreeSetDragCompareFunc(c, cmpFunc) }
func (c *CTree) SetExpanderStyle(expanderStyle CTreeExpanderStyle) {
	ctreeSetExpanderStyle(c, expanderStyle)
}
func (c *CTree) SetIndent(indent int)                  { ctreeSetIndent(c, indent) }
func (c *CTree) SetLineStyle(lineStyle CTreeLineStyle) { ctreeSetLineStyle(c, lineStyle) }
func (c *CTree) SetNodeInfo(node *CTreeNode, text string, spacing uint8, pixmapClosed *D.Pixmap, maskClosed *T.GdkBitmap, pixmapOpened *D.Pixmap, maskOpened *T.GdkBitmap, isLeaf, expanded T.Gboolean) {
	ctreeSetNodeInfo(c, node, text, spacing, pixmapClosed, maskClosed, pixmapOpened, maskOpened, isLeaf, expanded)
}
func (c *CTree) SetShowStub(showStub T.Gboolean)          { ctreeSetShowStub(c, showStub) }
func (c *CTree) SetSpacing(spacing int)                   { ctreeSetSpacing(c, spacing) }
func (c *CTree) SortNode(node *CTreeNode)                 { ctreeSortNode(c, node) }
func (c *CTree) SortRecursive(node *CTreeNode)            { ctreeSortRecursive(c, node) }
func (c *CTree) ToggleExpansion(node *CTreeNode)          { ctreeToggleExpansion(c, node) }
func (c *CTree) ToggleExpansionRecursive(node *CTreeNode) { ctreeToggleExpansionRecursive(c, node) }
func (c *CTree) Unselect(node *CTreeNode)                 { ctreeUnselect(c, node) }
func (c *CTree) UnselectRecursive(node *CTreeNode)        { ctreeUnselectRecursive(c, node) }

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
	Point        *T.GdkPoint
	NumCtlpoints int
	Ctlpoint     [2]*float32 //TODO(t): float32 (*ctlpoint)[2]; ???
}

var (
	CurveGetType func() T.GType
	CurveNew     func() *Widget

	curveGetVector    func(c *Curve, veclen int, vector *float32)
	curveReset        func(c *Curve)
	curveSetCurveType func(c *Curve, t CurveType)
	curveSetGamma     func(c *Curve, gamma float32)
	curveSetRange     func(c *Curve, minX, maxX, minY, maxY float32)
	curveSetVector    func(c *Curve, veclen int, vector *float32)
)

func (c *Curve) GetVector(veclen int, vector *float32)   { curveGetVector(c, veclen, vector) }
func (c *Curve) Reset()                                  { curveReset(c) }
func (c *Curve) SetCurveType(t CurveType)                { curveSetCurveType(c, t) }
func (c *Curve) SetGamma(gamma float32)                  { curveSetGamma(c, gamma) }
func (c *Curve) SetRange(minX, maxX, minY, maxY float32) { curveSetRange(c, minX, maxX, minY, maxY) }
func (c *Curve) SetVector(veclen int, vector *float32)   { curveSetVector(c, veclen, vector) }

type CurveType Enum

const (
	CURVE_TYPE_LINEAR CurveType = iota
	CURVE_TYPE_SPLINE
	CURVE_TYPE_FREE
)

var CurveTypeGetType func() T.GType
