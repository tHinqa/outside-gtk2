package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type (
	Calendar struct {
		Widget            T.GtkWidget
		Header_style      *T.GtkStyle
		Label_style       *T.GtkStyle
		Month             int
		Year              int
		Selected_day      int
		Day_month         [6][7]int
		Day               [6][7]int
		Num_marked_dates  int
		Marked_date       [31]int
		Display_flags     CalendarDisplayOptions
		Marked_date_color [31]T.GdkColor
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

type CalendarDisplayOptions T.Enum

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
	CalendarNew                   func() *T.GtkWidget
)

var (
	CalendarClearMarks          func(c *Calendar)
	CalendarFreeze              func(c *Calendar)
	CalendarGetDate             func(c *Calendar, year, month, day *uint)
	CalendarGetDetailHeightRows func(c *Calendar) int
	CalendarGetDetailWidthChars func(c *Calendar) int
	CalendarGetDisplayOptions   func(c *Calendar) CalendarDisplayOptions
	CalendarMarkDay             func(c *Calendar, day uint) T.Gboolean
	CalendarSelectDay           func(c *Calendar, day uint)
	CalendarSelectMonth         func(c *Calendar, month, year uint) T.Gboolean
	CalendarSetDetailFunc       func(c *Calendar, f CalendarDetailFunc, dataGpointer, destroy T.GDestroyNotify)
	CalendarSetDetailHeightRows func(c *Calendar, rows int)
	CalendarSetDetailWidthChars func(c *Calendar, chars int)
	CalendarSetDisplayOptions   func(c *Calendar, flags CalendarDisplayOptions)
	CalendarThaw                func(c *Calendar)
	CalendarUnmarkDay           func(c *Calendar, day uint) T.Gboolean

//NOTE(t):Deprecated; We remove because name clash w/enum
//CalendarDisplayOptions      func(c *Calendar, flags CalendarDisplayOptionsEnum)
)

func (c *Calendar) SelectMonth(month, year uint) T.Gboolean {
	return CalendarSelectMonth(c, month, year)
}

func (c *Calendar) SelectDay(day uint) { CalendarSelectDay(c, day) }

func (c *Calendar) MarkDay(day uint) T.Gboolean {
	return CalendarMarkDay(c, day)
}

func (c *Calendar) UnmarkDay(day uint) T.Gboolean {
	return CalendarUnmarkDay(c, day)
}

func (c *Calendar) ClearMarks() { CalendarClearMarks(c) }

func (c *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	CalendarSetDisplayOptions(c, flags)
}

func (c *Calendar) GetDisplayOptions() CalendarDisplayOptions {
	return CalendarGetDisplayOptions(c)
}

//NOTE(t):Deprecated; We remove because name clash w/enum
// func (c *Calendar) DisplayOptions(flags CalendarDisplayOptionsEnum) {
// 	CalendarDisplayOptions(c, flags)
// }

func (c *Calendar) GetDate(year, month, day *uint) {
	CalendarGetDate(c, year, month, day)
}

func (c *Calendar) SetDetailFunc(f CalendarDetailFunc, dataGpointer, destroy T.GDestroyNotify) {
	CalendarSetDetailFunc(c, f, dataGpointer, destroy)
}

func (c *Calendar) SetDetailWidthChars(chars int) {
	CalendarSetDetailWidthChars(c, chars)
}

func (c *Calendar) SetDetailHeightRows(rows int) {
	CalendarSetDetailHeightRows(c, rows)
}

func (c *Calendar) GetDetailWidthChars() int {
	return CalendarGetDetailWidthChars(c)
}

func (c *Calendar) GetDetailHeightRows() int {
	return CalendarGetDetailHeightRows(c)
}

func (c *Calendar) Freeze() { CalendarFreeze(c) }

func (c *Calendar) Thaw() { CalendarThaw(c) }

type CellEditable struct{}

var CellEditableGetType func() T.GType

var (
	CellEditableEditingDone  func(c *CellEditable)
	CellEditableRemoveWidget func(c *CellEditable)
	CellEditableStartEditing func(c *CellEditable, event *T.GdkEvent)
)

func (c *CellEditable) StartEditing(event *T.GdkEvent) {
	CellEditableStartEditing(c, event)
}

func (c *CellEditable) EditingDone() { CellEditableEditingDone(c) }

func (c *CellEditable) RemoveWidget() { CellEditableRemoveWidget(c) }

type CellLayout struct{}

var (
	CellLayoutGetType func() T.GType

	CellLayoutAddAttribute    func(c *CellLayout, cell *CellRenderer, attribute string, column int)
	CellLayoutClear           func(c *CellLayout)
	CellLayoutClearAttributes func(c *CellLayout, cell *CellRenderer)
	CellLayoutGetCells        func(c *CellLayout) *T.GList
	CellLayoutPackEnd         func(c *CellLayout, cell *CellRenderer, expand T.Gboolean)
	CellLayoutPackStart       func(c *CellLayout, cell *CellRenderer, expand T.Gboolean)
	CellLayoutReorder         func(c *CellLayout, cell *CellRenderer, position int)
	CellLayoutSetAttributes   func(c *CellLayout, cell *CellRenderer, v ...VArg)
	CellLayoutSetCellDataFunc func(c *CellLayout, cell *CellRenderer, f CellLayoutDataFunc, funcDataGpointer, destroy T.GDestroyNotify)
)

type CellLayoutDataFunc func(
	cellLayout *CellLayout,
	cell *CellRenderer,
	treeModel *T.GtkTreeModel,
	iter *T.GtkTreeIter,
	data T.Gpointer)

func (c *CellLayout) PackStart(
	cell *CellRenderer, expand T.Gboolean) {
	CellLayoutPackStart(c, cell, expand)
}

func (c *CellLayout) PackEnd(
	cell *CellRenderer, expand T.Gboolean) {
	CellLayoutPackEnd(c, cell, expand)
}

func (c *CellLayout) GetCells() *T.GList {
	return CellLayoutGetCells(c)
}

func (c *CellLayout) Clear() { CellLayoutClear(c) }

func (c *CellLayout) SetAttributes(
	cell *CellRenderer, v ...VArg) {
	CellLayoutSetAttributes(c, cell, v)
}

func (c *CellLayout) AddAttribute(
	cell *CellRenderer, attribute string, column int) {
	CellLayoutAddAttribute(c, cell, attribute, column)
}

func (c *CellLayout) SetCellDataFunc(
	cell *CellRenderer, f CellLayoutDataFunc,
	funcDataGpointer, destroy T.GDestroyNotify) {
	CellLayoutSetCellDataFunc(c, cell, f, funcDataGpointer, destroy)
}

func (c *CellLayout) ClearAttributes(cell *CellRenderer) {
	CellLayoutClearAttributes(c, cell)
}

func (c *CellLayout) Reorder(cell *CellRenderer, position int) {
	CellLayoutReorder(c, cell, position)
}

type CellRenderer struct {
	Parent T.GtkObject
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

type CellRendererState T.Enum

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
	CellRendererActivate        func(c *CellRenderer, event *T.GdkEvent, widget *T.GtkWidget, path string, backgroundArea, cellArea *T.GdkRectangle, flags CellRendererState) T.Gboolean
	CellRendererEditingCanceled func(c *CellRenderer)
	CellRendererGetAlignment    func(c *CellRenderer, xalign, yalign *float32)
	CellRendererGetFixedSize    func(c *CellRenderer, width, height *int)
	CellRendererGetPadding      func(c *CellRenderer, xpad, ypad *int)
	CellRendererGetSensitive    func(c *CellRenderer) T.Gboolean
	CellRendererGetSize         func(c *CellRenderer, widget *T.GtkWidget, cellArea *T.GdkRectangle, xOffset, yOffset, width, height *int)
	CellRendererGetVisible      func(c *CellRenderer) T.Gboolean
	CellRendererRender          func(c *CellRenderer, window *T.GdkWindow, widget *T.GtkWidget, backgroundArea, cellArea, exposeArea *T.GdkRectangle, flags CellRendererState)
	CellRendererSetAlignment    func(c *CellRenderer, xalign, yalign float32)
	CellRendererSetFixedSize    func(c *CellRenderer, width, height int)
	CellRendererSetPadding      func(c *CellRenderer, xpad, ypad int)
	CellRendererSetSensitive    func(c *CellRenderer, sensitive T.Gboolean)
	CellRendererSetVisible      func(c *CellRenderer, visible T.Gboolean)
	CellRendererStartEditing    func(c *CellRenderer, event *T.GdkEvent, widget *T.GtkWidget, path string, backgroundArea, cellArea *T.GdkRectangle, flags CellRendererState) *CellEditable
	CellRendererStopEditing     func(c *CellRenderer, canceled T.Gboolean)
)

func (c *CellRenderer) GetSize(widget *T.GtkWidget, cellArea *T.GdkRectangle, xOffset, yOffset, width, height *int) {
	CellRendererGetSize(c, widget, cellArea, xOffset, yOffset, width, height)
}

func (c *CellRenderer) Render(window *T.GdkWindow, widget *T.GtkWidget, backgroundArea, cellArea, exposeArea *T.GdkRectangle, flags CellRendererState) {
	CellRendererRender(c, window, widget, backgroundArea, cellArea, exposeArea, flags)
}

func (c *CellRenderer) Activate(event *T.GdkEvent, widget *T.GtkWidget, path string, backgroundArea, cellArea *T.GdkRectangle, flags CellRendererState) T.Gboolean {
	return CellRendererActivate(c, event, widget, path, backgroundArea, cellArea, flags)
}

func (c *CellRenderer) StartEditing(event *T.GdkEvent, widget *T.GtkWidget, path string, backgroundArea, cellArea *T.GdkRectangle, flags CellRendererState) *CellEditable {
	return CellRendererStartEditing(c, event, widget, path, backgroundArea, cellArea, flags)
}

func (c *CellRenderer) SetFixedSize(width, height int) {
	CellRendererSetFixedSize(c, width, height)
}

func (c *CellRenderer) GetFixedSize(width, height *int) {
	CellRendererGetFixedSize(c, width, height)
}

func (c *CellRenderer) SetAlignment(xalign, yalign float32) {
	CellRendererSetAlignment(c, xalign, yalign)
}

func (c *CellRenderer) GetAlignment(xalign, yalign *float32) {
	CellRendererGetAlignment(c, xalign, yalign)
}

func (c *CellRenderer) SetPadding(xpad, ypad int) {
	CellRendererSetPadding(c, xpad, ypad)
}

func (c *CellRenderer) GetPadding(xpad, ypad *int) {
	CellRendererGetPadding(c, xpad, ypad)
}

func (c *CellRenderer) SetVisible(visible T.Gboolean) {
	CellRendererSetVisible(c, visible)
}

func (c *CellRenderer) GetVisible() T.Gboolean {
	return CellRendererGetVisible(c)
}

func (c *CellRenderer) SetSensitive(sensitive T.Gboolean) {
	CellRendererSetSensitive(c, sensitive)
}

func (c *CellRenderer) GetSensitive() T.Gboolean {
	return CellRendererGetSensitive(c)
}

func (c *CellRenderer) EditingCanceled() {
	CellRendererEditingCanceled(c)
}

func (c *CellRenderer) StopEditing(canceled T.Gboolean) {
	CellRendererStopEditing(c, canceled)
}

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
	CellRendererToggleGetActivatable func(c *CellRendererToggle) T.Gboolean
	CellRendererToggleGetActive      func(c *CellRendererToggle) T.Gboolean
	CellRendererToggleGetRadio       func(c *CellRendererToggle) T.Gboolean
	CellRendererToggleSetActivatable func(c *CellRendererToggle, setting T.Gboolean)
	CellRendererToggleSetActive      func(c *CellRendererToggle, setting T.Gboolean)
	CellRendererToggleSetRadio       func(c *CellRendererToggle, radio T.Gboolean)
)

func (c *CellRendererToggle) GetRadio() T.Gboolean {
	return CellRendererToggleGetRadio(c)
}

func (c *CellRendererToggle) SetRadio(radio T.Gboolean) {
	CellRendererToggleSetRadio(c, radio)
}

func (c *CellRendererToggle) GetActive() T.Gboolean {
	return CellRendererToggleGetActive(c)
}

func (c *CellRendererToggle) SetActive(setting T.Gboolean) {
	CellRendererToggleSetActive(c, setting)
}

func (c *CellRendererToggle) GetActivatable() T.Gboolean {
	return CellRendererToggleGetActivatable(c)
}

func (c *CellRendererToggle) SetActivatable(setting T.Gboolean) {
	CellRendererToggleSetActivatable(c, setting)
}

type CellView struct {
	Parent T.GtkWidget
	_      struct{}
}

var CellTypeGetType func() T.GType

var (
	CellViewGetType       func() T.GType
	CellViewNew           func() *T.GtkWidget
	CellViewNewWithText   func(text string) *T.GtkWidget
	CellViewNewWithMarkup func(markup string) *T.GtkWidget
	CellViewNewWithPixbuf func(pixbuf *T.GdkPixbuf) *T.GtkWidget
)

var (
	CellViewGetCellRenderers   func(c *CellView) *T.GList
	CellViewGetDisplayedRow    func(c *CellView) *T.GtkTreePath
	CellViewGetModel           func(c *CellView) *T.GtkTreeModel
	CellViewGetSizeOfRow       func(c *CellView, path *T.GtkTreePath, requisition *T.GtkRequisition) T.Gboolean
	CellViewSetBackgroundColor func(c *CellView, color *T.GdkColor)
	CellViewSetDisplayedRow    func(c *CellView, path *T.GtkTreePath)
	CellViewSetModel           func(c *CellView, model *T.GtkTreeModel)
)

func (c *CellView) SetModel(model *T.GtkTreeModel) {
	CellViewSetModel(c, model)
}

func (c *CellView) iewGetModel() *T.GtkTreeModel {
	return CellViewGetModel(c)
}

func (c *CellView) iewSetDisplayedRow(path *T.GtkTreePath) {
	CellViewSetDisplayedRow(c, path)
}

func (c *CellView) iewGetDisplayedRow() *T.GtkTreePath {
	return CellViewGetDisplayedRow(c)
}

func (c *CellView) iewGetSizeOfRow(path *T.GtkTreePath,
	requisition *T.GtkRequisition) T.Gboolean {
	return CellViewGetSizeOfRow(c, path, requisition)
}

func (c *CellView) iewSetBackgroundColor(color *T.GdkColor) {
	CellViewSetBackgroundColor(c, color)
}

func (c *CellView) iewGetCellRenderers() *T.GList {
	return CellViewGetCellRenderers(c)
}

var (
	CheckButtonGetType         func() T.GType
	CheckButtonNew             func() *T.GtkWidget
	CheckButtonNewWithLabel    func(label string) *T.GtkWidget
	CheckButtonNewWithMnemonic func(label string) *T.GtkWidget
)

type CheckMenuItem struct {
	Menu_item T.GtkMenuItem
	Bits      uint
	// Active : 1
	// AlwaysShowToggle : 1
	// Inconsistent : 1
	// DrawAsRadio : 1
}

var (
	CheckMenuItemGetType         func() T.GType
	CheckMenuItemNew             func() *T.GtkWidget
	CheckMenuItemNewWithLabel    func(label string) *T.GtkWidget
	CheckMenuItemNewWithMnemonic func(label string) *T.GtkWidget
)

var (
	CheckMenuItemGetActive       func(c *CheckMenuItem) T.Gboolean
	CheckMenuItemGetDrawAsRadio  func(c *CheckMenuItem) T.Gboolean
	CheckMenuItemGetInconsistent func(c *CheckMenuItem) T.Gboolean
	CheckMenuItemSetActive       func(c *CheckMenuItem, isActive T.Gboolean)
	CheckMenuItemSetDrawAsRadio  func(c *CheckMenuItem, drawAsRadio T.Gboolean)
	CheckMenuItemSetInconsistent func(c *CheckMenuItem, setting T.Gboolean)
	CheckMenuItemSetShowToggle   func(c *CheckMenuItem, always T.Gboolean)
	CheckMenuItemToggled         func(c *CheckMenuItem)
)

func (c *CheckMenuItem) SetActive(isActive T.Gboolean) {
	CheckMenuItemSetActive(c, isActive)
}

func (c *CheckMenuItem) GetActive() T.Gboolean {
	return CheckMenuItemGetActive(c)
}

func (c *CheckMenuItem) Toggled() {
	CheckMenuItemToggled(c)
}

func (c *CheckMenuItem) SetInconsistent(setting T.Gboolean) {
	CheckMenuItemSetInconsistent(c, setting)
}

func (c *CheckMenuItem) GetInconsistent() T.Gboolean {
	return CheckMenuItemGetInconsistent(c)
}

func (c *CheckMenuItem) SetDrawAsRadio(drawAsRadio T.Gboolean) {
	CheckMenuItemSetDrawAsRadio(c, drawAsRadio)
}

func (c *CheckMenuItem) GetDrawAsRadio() T.Gboolean {
	return CheckMenuItemGetDrawAsRadio(c)
}

func (c *CheckMenuItem) SetShowToggle(always T.Gboolean) {
	CheckMenuItemSetShowToggle(c, always)
}

type Clipboard struct{}

var (
	ClipboardGetType       func() T.GType
	ClipboardGet           func(selection T.GdkAtom) *Clipboard
	ClipboardGetForDisplay func(display *T.GdkDisplay, selection T.GdkAtom) *Clipboard
	WidgetGetClipboard     func(widget *T.GtkWidget, selection T.GdkAtom) *Clipboard
)

var (
	ClipboardClear                   func(c *Clipboard)
	ClipboardGetDisplay              func(c *Clipboard) *T.GdkDisplay
	ClipboardGetOwner                func(c *Clipboard) *T.GObject
	ClipboardRequestContents         func(c *Clipboard, target T.GdkAtom, callback ClipboardReceivedFunc, userData T.Gpointer)
	ClipboardRequestImage            func(c *Clipboard, callback ClipboardImageReceivedFunc, userData T.Gpointer)
	ClipboardRequestRichText         func(c *Clipboard, buffer *T.GtkTextBuffer, callback ClipboardRichTextReceivedFunc, userData T.Gpointer)
	ClipboardRequestTargets          func(c *Clipboard, callback ClipboardTargetsReceivedFunc, userData T.Gpointer)
	ClipboardRequestText             func(c *Clipboard, callback ClipboardTextReceivedFunc, userData T.Gpointer)
	ClipboardRequestUris             func(c *Clipboard, callback ClipboardURIReceivedFunc, userData T.Gpointer)
	ClipboardSetCanStore             func(c *Clipboard, targets *T.GtkTargetEntry, nTargets int)
	ClipboardSetImage                func(c *Clipboard, pixbuf *T.GdkPixbuf)
	ClipboardSetText                 func(c *Clipboard, text string, len int)
	ClipboardSetWithData             func(c *Clipboard, targets *T.GtkTargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, userData T.Gpointer) T.Gboolean
	ClipboardSetWithOwner            func(c *Clipboard, targets *T.GtkTargetEntry, nTargets uint, getFunc ClipboardGetFunc, clearFunc ClipboardClearFunc, owner *T.GObject) T.Gboolean
	ClipboardStore                   func(c *Clipboard)
	ClipboardWaitForContents         func(c *Clipboard, target T.GdkAtom) *T.GtkSelectionData
	ClipboardWaitForImage            func(c *Clipboard) *T.GdkPixbuf
	ClipboardWaitForRichText         func(c *Clipboard, buffer *T.GtkTextBuffer, format *T.GdkAtom, length *T.Gsize) *uint8
	ClipboardWaitForTargets          func(c *Clipboard, targets **T.GdkAtom, nTargets *int) T.Gboolean
	ClipboardWaitForText             func(c *Clipboard) string
	ClipboardWaitForUris             func(c *Clipboard) **T.Gchar
	ClipboardWaitIsImageAvailable    func(c *Clipboard) T.Gboolean
	ClipboardWaitIsRichTextAvailable func(c *Clipboard, buffer *T.GtkTextBuffer) T.Gboolean
	ClipboardWaitIsTargetAvailable   func(c *Clipboard, target T.GdkAtom) T.Gboolean
	ClipboardWaitIsTextAvailable     func(c *Clipboard) T.Gboolean
	ClipboardWaitIsUrisAvailable     func(c *Clipboard) T.Gboolean
)

type (
	ClipboardClearFunc func(
		clipboard *Clipboard,
		userDataOrOwner T.Gpointer)

	ClipboardGetFunc func(
		clipboard *Clipboard,
		selectionData *T.GtkSelectionData,
		info uint,
		userDataOrOwner T.Gpointer)

	ClipboardImageReceivedFunc func(
		clipboard *Clipboard,
		pixbuf *T.GdkPixbuf,
		data T.Gpointer)

	ClipboardReceivedFunc func(
		clipboard *Clipboard,
		selectionData *T.GtkSelectionData,
		data T.Gpointer)

	ClipboardRichTextReceivedFunc func(
		clipboard *Clipboard,
		format T.GdkAtom,
		text *uint8,
		length T.Gsize,
		data T.Gpointer)

	ClipboardTargetsReceivedFunc func(
		clipboard *Clipboard,
		atoms *T.GdkAtom,
		nAtoms int,
		data T.Gpointer)

	ClipboardTextReceivedFunc func(
		clipboard *Clipboard,
		text string,
		data T.Gpointer)

	ClipboardURIReceivedFunc func(
		clipboard *Clipboard,
		uris *string,
		data T.Gpointer)
)

func (c *Clipboard) GetDisplay() *T.GdkDisplay {
	return ClipboardGetDisplay(c)
}

func (c *Clipboard) SetWithData(targets *T.GtkTargetEntry,
	nTargets uint, getFunc ClipboardGetFunc,
	clearFunc ClipboardClearFunc, userData T.Gpointer) T.Gboolean {
	return ClipboardSetWithData(
		c, targets, nTargets, getFunc, clearFunc, userData)
}

func (c *Clipboard) SetWithOwner(targets *T.GtkTargetEntry,
	nTargets uint, getFunc ClipboardGetFunc,
	clearFunc ClipboardClearFunc, owner *T.GObject) T.Gboolean {
	return ClipboardSetWithOwner(
		c, targets, nTargets, getFunc, clearFunc, owner)
}

func (c *Clipboard) GetOwner() *T.GObject {
	return ClipboardGetOwner(c)
}

func (c *Clipboard) Clear() {
	ClipboardClear(c)
}

func (c *Clipboard) SetText(text string, leng int) {
	ClipboardSetText(c, text, leng)
}

func (c *Clipboard) SetImage(pixbuf *T.GdkPixbuf) {
	ClipboardSetImage(c, pixbuf)
}

func (c *Clipboard) RequestContents(target T.GdkAtom,
	callback ClipboardReceivedFunc, userData T.Gpointer) {
	ClipboardRequestContents(c, target, callback, userData)
}

func (c *Clipboard) RequestText(
	callback ClipboardTextReceivedFunc, userData T.Gpointer) {
	ClipboardRequestText(c, callback, userData)
}

func (c *Clipboard) RequestRichText(buffer *T.GtkTextBuffer,
	callback ClipboardRichTextReceivedFunc, userData T.Gpointer) {
	ClipboardRequestRichText(c, buffer, callback, userData)
}

func (c *Clipboard) RequestImage(
	callback ClipboardImageReceivedFunc, userData T.Gpointer) {
	ClipboardRequestImage(c, callback, userData)
}

func (c *Clipboard) RequestUris(
	callback ClipboardURIReceivedFunc, userData T.Gpointer) {
	ClipboardRequestUris(c, callback, userData)
}

func (c *Clipboard) RequestTargets(
	callback ClipboardTargetsReceivedFunc, userData T.Gpointer) {
	ClipboardRequestTargets(c, callback, userData)
}

func (c *Clipboard) WaitForContents(
	target T.GdkAtom) *T.GtkSelectionData {
	return ClipboardWaitForContents(c, target)
}

func (c *Clipboard) WaitForText() string {
	return ClipboardWaitForText(c)
}

func (c *Clipboard) WaitForRichText(buffer *T.GtkTextBuffer,
	format *T.GdkAtom, length *T.Gsize) *uint8 {
	return ClipboardWaitForRichText(c, buffer, format, length)
}

func (c *Clipboard) WaitForImage() *T.GdkPixbuf {
	return ClipboardWaitForImage(c)
}

func (c *Clipboard) WaitForUris() **T.Gchar {
	return ClipboardWaitForUris(c)
}

func (c *Clipboard) WaitForTargets(
	targets **T.GdkAtom, nTargets *int) T.Gboolean {
	return ClipboardWaitForTargets(c, targets, nTargets)
}

func (c *Clipboard) WaitIsTextAvailable() T.Gboolean {
	return ClipboardWaitIsTextAvailable(c)
}

func (c *Clipboard) WaitIsRichTextAvailable(
	buffer *T.GtkTextBuffer) T.Gboolean {
	return ClipboardWaitIsRichTextAvailable(c, buffer)
}

func (c *Clipboard) WaitIsImageAvailable() T.Gboolean {
	return ClipboardWaitIsImageAvailable(c)
}

func (c *Clipboard) WaitIsUrisAvailable() T.Gboolean {
	return ClipboardWaitIsUrisAvailable(c)
}

func (c *Clipboard) WaitIsTargetAvailable(
	target T.GdkAtom) T.Gboolean {
	return ClipboardWaitIsTargetAvailable(c, target)
}

func (c *Clipboard) SetCanStore(
	targets *T.GtkTargetEntry, nTargets int) {
	ClipboardSetCanStore(c, targets, nTargets)
}

func (c *Clipboard) Store() { ClipboardStore(c) }

type (
	CList struct {
		Container          T.GtkContainer
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
		TitleWindow        *T.GdkWindow
		Column             *CListColumn
		ClistWindow        *T.GdkWindow
		ClistWindowWidth   int
		ClistWindowHeight  int
		Hoffset            int
		Voffset            int
		ShadowType         T.GtkShadowType
		SelectionMode      T.GtkSelectionMode
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
		CursorDrag         *T.GdkCursor
		XDrag              int
		FocusRow           int
		FocusHeaderColumn  int
		Anchor             int
		AnchorState        T.GtkStateType
		DragPos            int
		Htimer             int
		Vtimer             int
		SortType           T.GtkSortType
		Compare            CListCompareFunc
		SortColumn         int
		DragHighlightRow   int
		DragHighlightPos   CListDragPos
	}

	CListColumn struct {
		Title         *T.Gchar
		Area          T.GdkRectangle
		Button        *T.GtkWidget
		Window        *T.GdkWindow
		Width         int
		Min_width     int
		Max_width     int
		Justification T.GtkJustification
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
		Cell       *T.GtkCell
		State      T.GtkStateType
		Foreground T.GdkColor
		Background T.GdkColor
		Style      *T.GtkStyle
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

type CListDragPos T.Enum

const (
	CLIST_DRAG_NONE CListDragPos = iota
	CLIST_DRAG_BEFORE
	CLIST_DRAG_INTO
	CLIST_DRAG_AFTER
)

var (
	ClistGetType        func() T.GType
	ClistNew            func(columns int) *T.GtkWidget
	ClistNewWithTitles  func(columns int, titles **T.Gchar) *T.GtkWidget
	ClistDragPosGetType func() T.GType
)

var (
	ClistAppend                 func(c *CList, text **T.Gchar) int
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
	ClistGetCellStyle           func(c *CList, row, column int) *T.GtkStyle
	ClistGetCellType            func(c *CList, row, column int) T.GtkCellType
	ClistGetColumnTitle         func(c *CList, column int) string
	ClistGetColumnWidget        func(c *CList, column int) *T.GtkWidget
	ClistGetHadjustment         func(c *CList) *Adjustment
	ClistGetPixmap              func(c *CList, row, column int, pixmap, mask **T.GdkBitmap) int
	ClistGetPixtext             func(c *CList, row, column int, text **T.Gchar, spacing *uint8, pixmap, mask **T.GdkBitmap) int
	ClistGetRowData             func(c *CList, row int) T.Gpointer
	ClistGetRowStyle            func(c *CList, row int) *T.GtkStyle
	ClistGetSelectable          func(c *CList, row int) T.Gboolean
	ClistGetSelectionInfo       func(c *CList, x, y int, row, column *int) int
	ClistGetText                func(c *CList, row, column int, text **T.Gchar) int
	ClistGetVadjustment         func(c *CList) *Adjustment
	ClistInsert                 func(c *CList, row int, text **T.Gchar) int
	ClistMoveto                 func(c *CList, row, column int, rowAlign, colAlign float32)
	ClistOptimalColumnWidth     func(c *CList, column int) int
	ClistPrepend                func(c *CList, text **T.Gchar) int
	ClistRemove                 func(c *CList, row int)
	ClistRowIsVisible           func(c *CList, row int) T.GtkVisibility
	ClistRowMove                func(c *CList, sourceRow, destRow int)
	ClistSelectAll              func(c *CList)
	ClistSelectRow              func(c *CList, row, column int)
	ClistSetAutoSort            func(c *CList, autoSort T.Gboolean)
	ClistSetBackground          func(c *CList, row int, color *T.GdkColor)
	ClistSetButtonActions       func(c *CList, button uint, buttonActions uint8)
	ClistSetCellStyle           func(c *CList, row, column int, style *T.GtkStyle)
	ClistSetColumnAutoResize    func(c *CList, column int, autoResize T.Gboolean)
	ClistSetColumnJustification func(c *CList, column int, justification T.GtkJustification)
	ClistSetColumnMaxWidth      func(c *CList, column, maxWidth int)
	ClistSetColumnMinWidth      func(c *CList, column, minWidth int)
	ClistSetColumnResizeable    func(c *CList, column int, resizeable T.Gboolean)
	ClistSetColumnTitle         func(c *CList, column int, title string)
	ClistSetColumnVisibility    func(c *CList, column int, visible T.Gboolean)
	ClistSetColumnWidget        func(c *CList, column int, widget *T.GtkWidget)
	ClistSetColumnWidth         func(c *CList, column, width int)
	ClistSetCompareFunc         func(c *CList, cmpFunc CListCompareFunc)
	ClistSetForeground          func(c *CList, row int, color *T.GdkColor)
	ClistSetHadjustment         func(c *CList, adjustment *Adjustment)
	ClistSetPixmap              func(c *CList, row, column int, pixmap *T.GdkPixmap, mask *T.GdkBitmap)
	ClistSetPixtext             func(c *CList, row, column int, text string, spacing uint8, pixmap, mask *T.GdkBitmap)
	ClistSetReorderable         func(c *CList, reorderable T.Gboolean)
	ClistSetRowData             func(c *CList, row int, data T.Gpointer)
	ClistSetRowDataFull         func(c *CList, row int, dataGpointer, destroy T.GDestroyNotify)
	ClistSetRowHeight           func(c *CList, height uint)
	ClistSetRowStyle            func(c *CList, row, style *T.GtkStyle)
	ClistSetSelectable          func(c *CList, row int, selectable T.Gboolean)
	ClistSetSelectionMode       func(c *CList, mode T.GtkSelectionMode)
	ClistSetShadowType          func(c *CList, t T.GtkShadowType)
	ClistSetShift               func(c *CList, row, column, vertical, horizontal int)
	ClistSetSortColumn          func(c *CList, column int)
	ClistSetSortType            func(c *CList, sortType T.GtkSortType)
	ClistSetText                func(c *CList, row, column int, text string)
	ClistSetUseDragIcons        func(c *CList, useIcons T.Gboolean)
	ClistSetVadjustment         func(c *CList, adjustment *Adjustment)
	ClistSort                   func(c *CList)
	ClistSwapRows               func(c *CList, row1, row2 int)
	ClistThaw                   func(c *CList)
	ClistUndoSelection          func(c *CList)
	ClistUnselectAll            func(c *CList)
	ClistUnselectRow            func(c *CList, row, column int)
)

func (c *CList) SetHadjustment(adjustment *Adjustment) {
	ClistSetHadjustment(c, adjustment)
}

func (c *CList) SetVadjustment(adjustment *Adjustment) {
	ClistSetVadjustment(c, adjustment)
}

func (c *CList) GetHadjustment() *Adjustment {
	return ClistGetHadjustment(c)
}

func (c *CList) GetVadjustment() *Adjustment {
	return ClistGetVadjustment(c)
}

func (c *CList) SetShadowType(t T.GtkShadowType) {
	ClistSetShadowType(c, t)
}

func (c *CList) SetSelectionMode(mode T.GtkSelectionMode) {
	ClistSetSelectionMode(c, mode)
}

func (c *CList) SetReorderable(reorderable T.Gboolean) {
	ClistSetReorderable(c, reorderable)
}

func (c *CList) SetUseDragIcons(useIcons T.Gboolean) {
	ClistSetUseDragIcons(c, useIcons)
}

func (c *CList) SetButtonActions(button uint, buttonActions uint8) {
	ClistSetButtonActions(c, button, buttonActions)
}

func (c *CList) Freeze() { ClistFreeze(c) }

func (c *CList) Thaw() { ClistThaw(c) }

func (c *CList) ColumnTitlesShow() { ClistColumnTitlesShow(c) }

func (c *CList) ColumnTitlesHide() { ClistColumnTitlesHide(c) }

func (c *CList) ColumnTitleActive(column int) {
	ClistColumnTitleActive(c, column)
}

func (c *CList) ColumnTitlePassive(column int) {
	ClistColumnTitlePassive(c, column)
}

func (c *CList) ColumnTitlesActive() {
	ClistColumnTitlesActive(c)
}

func (c *CList) ColumnTitlesPassive() {
	ClistColumnTitlesPassive(c)
}

func (c *CList) SetColumnTitle(column int, title string) {
	ClistSetColumnTitle(c, column, title)
}

func (c *CList) GetColumnTitle(column int) string {
	return ClistGetColumnTitle(c, column)
}

func (c *CList) SetColumnWidget(column int, widget *T.GtkWidget) {
	ClistSetColumnWidget(c, column, widget)
}

func (c *CList) GetColumnWidget(column int) *T.GtkWidget {
	return ClistGetColumnWidget(c, column)
}

func (c *CList) SetColumnJustification(
	column int, justification T.GtkJustification) {
	ClistSetColumnJustification(c, column, justification)
}

func (c *CList) SetColumnVisibility(
	column int, visible T.Gboolean) {
	ClistSetColumnVisibility(c, column, visible)
}

func (c *CList) SetColumnResizeable(
	column int, resizeable T.Gboolean) {
	ClistSetColumnResizeable(c, column, resizeable)
}

func (c *CList) SetColumnAutoResize(
	column int, autoResize T.Gboolean) {
	ClistSetColumnAutoResize(c, column, autoResize)
}

func (c *CList) ColumnsAutosize() int {
	return ClistColumnsAutosize(c)
}

func (c *CList) OptimalColumnWidth(column int) int {
	return ClistOptimalColumnWidth(c, column)
}

func (c *CList) SetColumnWidth(column, width int) {
	ClistSetColumnWidth(c, column, width)
}

func (c *CList) SetColumnMinWidth(column, minWidth int) {
	ClistSetColumnMinWidth(c, column, minWidth)
}

func (c *CList) SetColumnMaxWidth(column, maxWidth int) {
	ClistSetColumnMaxWidth(c, column, maxWidth)
}

func (c *CList) SetRowHeight(height uint) {
	ClistSetRowHeight(c, height)
}

func (c *CList) Moveto(
	row, column int, rowAlign, colAlign float32) {
	ClistMoveto(c, row, column, rowAlign, colAlign)
}

func (c *CList) RowIsVisible(row int) T.GtkVisibility {
	return ClistRowIsVisible(c, row)
}

func (c *CList) GetCellType(row, column int) T.GtkCellType {
	return ClistGetCellType(c, row, column)
}

func (c *CList) SetText(row, column int, text string) {
	ClistSetText(c, row, column, text)
}

func (c *CList) GetText(row, column int, text **T.Gchar) int {
	return ClistGetText(c, row, column, text)
}

func (c *CList) SetPixmap(
	row, column int, pixmap *T.GdkPixmap, mask *T.GdkBitmap) {
	ClistSetPixmap(c, row, column, pixmap, mask)
}

func (c *CList) GetPixmap(
	row, column int, pixmap, mask **T.GdkBitmap) int {
	return ClistGetPixmap(c, row, column, pixmap, mask)
}

func (c *CList) SetPixtext(row, column int, text string,
	spacing uint8, pixmap, mask *T.GdkBitmap) {
	ClistSetPixtext(
		c, row, column, text, spacing, pixmap, mask)
}

func (c *CList) GetPixtext(row, column int, text **T.Gchar,
	spacing *uint8, pixmap, mask **T.GdkBitmap) int {
	return ClistGetPixtext(
		c, row, column, text, spacing, pixmap, mask)
}

func (c *CList) SetForeground(row int, color *T.GdkColor) {
	ClistSetForeground(c, row, color)
}

func (c *CList) SetBackground(row int, color *T.GdkColor) {
	ClistSetBackground(c, row, color)
}

func (c *CList) SetCellStyle(row, column int, style *T.GtkStyle) {
	ClistSetCellStyle(c, row, column, style)
}

func (c *CList) GetCellStyle(row, column int) *T.GtkStyle {
	return ClistGetCellStyle(c, row, column)
}

func (c *CList) SetRowStyle(row, style *T.GtkStyle) {
	ClistSetRowStyle(c, row, style)
}

func (c *CList) GetRowStyle(row int) *T.GtkStyle {
	return ClistGetRowStyle(c, row)
}

func (c *CList) SetShift(row, column, vertical, horizontal int) {
	ClistSetShift(c, row, column, vertical, horizontal)
}

func (c *CList) SetSelectable(row int, selectable T.Gboolean) {
	ClistSetSelectable(c, row, selectable)
}

func (c *CList) GetSelectable(row int) T.Gboolean {
	return ClistGetSelectable(c, row)
}

func (c *CList) Prepend(text **T.Gchar) int {
	return ClistPrepend(c, text)
}

func (c *CList) Append(text **T.Gchar) int {
	return ClistAppend(c, text)
}

func (c *CList) Insert(row int, text **T.Gchar) int {
	return ClistInsert(c, row, text)
}

func (c *CList) Remove(row int) {
	ClistRemove(c, row)
}

func (c *CList) SetRowData(row int, data T.Gpointer) {
	ClistSetRowData(c, row, data)
}

func (c *CList) SetRowDataFull(
	row int, dataGpointer, destroy T.GDestroyNotify) {
	ClistSetRowDataFull(c, row, dataGpointer, destroy)
}

func (c *CList) GetRowData(row int) T.Gpointer {
	return ClistGetRowData(c, row)
}

func (c *CList) FindRowFromData(data T.Gpointer) int {
	return ClistFindRowFromData(c, data)
}

func (c *CList) SelectRow(row, column int) {
	ClistSelectRow(c, row, column)
}

func (c *CList) UnselectRow(row, column int) {
	ClistUnselectRow(c, row, column)
}

func (c *CList) UndoSelection() { ClistUndoSelection(c) }

func (c *CList) Clear() { ClistClear(c) }

func (c *CList) GetSelectionInfo(x, y int, row, column *int) int {
	return ClistGetSelectionInfo(c, x, y, row, column)
}

func (c *CList) SelectAll() { ClistSelectAll(c) }

func (c *CList) UnselectAll() { ClistUnselectAll(c) }

func (c *CList) SwapRows(row1, row2 int) {
	ClistSwapRows(c, row1, row2)
}

func (c *CList) RowMove(sourceRow, destRow int) {
	ClistRowMove(c, sourceRow, destRow)
}

func (c *CList) SetCompareFunc(cmpFunc CListCompareFunc) {
	ClistSetCompareFunc(c, cmpFunc)
}

func (c *CList) SetSortColumn(column int) {
	ClistSetSortColumn(c, column)
}

func (c *CList) SetSortType(sortType T.GtkSortType) {
	ClistSetSortType(c, sortType)
}

func (c *CList) Sort() { ClistSort(c) }

func (c *CList) SetAutoSort(autoSort T.Gboolean) {
	ClistSetAutoSort(c, autoSort)
}

type ColorButton struct {
	Button Button
	_      *struct{}
}

var (
	ColorButtonGetType      func() T.GType
	ColorButtonNew          func() *T.GtkWidget
	ColorButtonNewWithColor func(color *T.GdkColor) *T.GtkWidget
)

var (
	ColorButtonGetAlpha    func(c *ColorButton) uint16
	ColorButtonGetColor    func(c *ColorButton, color *T.GdkColor)
	ColorButtonGetTitle    func(c *ColorButton) string
	ColorButtonGetUseAlpha func(c *ColorButton) T.Gboolean
	ColorButtonSetAlpha    func(c *ColorButton, alpha uint16)
	ColorButtonSetColor    func(c *ColorButton, color *T.GdkColor)
	ColorButtonSetTitle    func(c *ColorButton, title string)
	ColorButtonSetUseAlpha func(c *ColorButton, useAlpha T.Gboolean)
)

func (c *ColorButton) SetColor(color *T.GdkColor) {
	ColorButtonSetColor(c, color)
}

func (c *ColorButton) SetAlpha(alpha uint16) {
	ColorButtonSetAlpha(c, alpha)
}

func (c *ColorButton) GetColor(color *T.GdkColor) {
	ColorButtonGetColor(c, color)
}

func (c *ColorButton) GetAlpha() uint16 {
	return ColorButtonGetAlpha(c)
}

func (c *ColorButton) SetUseAlpha(useAlpha T.Gboolean) {
	ColorButtonSetUseAlpha(c, useAlpha)
}

func (c *ColorButton) GetUseAlpha() T.Gboolean {
	return ColorButtonGetUseAlpha(c)
}

func (c *ColorButton) SetTitle(title string) {
	ColorButtonSetTitle(c, title)
}

func (c *ColorButton) GetTitle() string {
	return ColorButtonGetTitle(c)
}

type (
	ColorSelection struct {
		Parent T.GtkVBox
		_      T.Gpointer
	}

	ColorSelectionChangePaletteFunc func(
		colors *T.GdkColor, nColors int)

	ColorSelectionChangePaletteWithScreenFunc func(
		screen *T.GdkScreen, colors *T.GdkColor, nColors int)
)

var (
	ColorSelectionGetType func() T.GType
	ColorSelectionNew     func() *T.GtkWidget

	ColorSelectionPaletteFromString func(str string, colors **T.GdkColor, nColors *int) T.Gboolean
	ColorSelectionPaletteToString   func(colors *T.GdkColor, nColors int) string

	ColorSelectionSetChangePaletteHook           func(f ColorSelectionChangePaletteFunc) ColorSelectionChangePaletteFunc
	ColorSelectionSetChangePaletteWithScreenHook func(f ColorSelectionChangePaletteWithScreenFunc) ColorSelectionChangePaletteWithScreenFunc
)

var (
	ColorSelectionGetColor             func(c *ColorSelection, color *float64)
	ColorSelectionGetCurrentAlpha      func(c *ColorSelection) uint16
	ColorSelectionGetCurrentColor      func(c *ColorSelection, color *T.GdkColor)
	ColorSelectionGetHasOpacityControl func(c *ColorSelection) T.Gboolean
	ColorSelectionGetHasPalette        func(c *ColorSelection) T.Gboolean
	ColorSelectionGetPreviousAlpha     func(c *ColorSelection) uint16
	ColorSelectionGetPreviousColor     func(c *ColorSelection, color *T.GdkColor)
	ColorSelectionIsAdjusting          func(c *ColorSelection) T.Gboolean
	ColorSelectionSetColor             func(c *ColorSelection, color *float64)
	ColorSelectionSetCurrentAlpha      func(c *ColorSelection, alpha uint16)
	ColorSelectionSetCurrentColor      func(c *ColorSelection, color *T.GdkColor)
	ColorSelectionSetHasOpacityControl func(c *ColorSelection, hasOpacity T.Gboolean)
	ColorSelectionSetHasPalette        func(c *ColorSelection, hasPalette T.Gboolean)
	ColorSelectionSetPreviousAlpha     func(c *ColorSelection, alpha uint16)
	ColorSelectionSetPreviousColor     func(c *ColorSelection, color *T.GdkColor)
	ColorSelectionSetUpdatePolicy      func(c *ColorSelection, policy T.GtkUpdateType)
)

func (c *ColorSelection) GetHasOpacityControl() T.Gboolean {
	return ColorSelectionGetHasOpacityControl(c)
}

func (c *ColorSelection) SetHasOpacityControl(
	hasOpacity T.Gboolean) {
	ColorSelectionSetHasOpacityControl(c, hasOpacity)
}

func (c *ColorSelection) GetHasPalette() T.Gboolean {
	return ColorSelectionGetHasPalette(c)
}

func (c *ColorSelection) SetHasPalette(hasPalette T.Gboolean) {
	ColorSelectionSetHasPalette(c, hasPalette)
}

func (c *ColorSelection) SetCurrentColor(color *T.GdkColor) {
	ColorSelectionSetCurrentColor(c, color)
}

func (c *ColorSelection) SetCurrentAlpha(alpha uint16) {
	ColorSelectionSetCurrentAlpha(c, alpha)
}

func (c *ColorSelection) GetCurrentColor(color *T.GdkColor) {
	ColorSelectionGetCurrentColor(c, color)
}

func (c *ColorSelection) GetCurrentAlpha() uint16 {
	return ColorSelectionGetCurrentAlpha(c)
}

func (c *ColorSelection) SetPreviousColor(color *T.GdkColor) {
	ColorSelectionSetPreviousColor(c, color)
}

func (c *ColorSelection) SetPreviousAlpha(alpha uint16) {
	ColorSelectionSetPreviousAlpha(c, alpha)
}

func (c *ColorSelection) GetPreviousColor(color *T.GdkColor) {
	ColorSelectionGetPreviousColor(c, color)
}

func (c *ColorSelection) GetPreviousAlpha() uint16 {
	return ColorSelectionGetPreviousAlpha(c)
}

func (c *ColorSelection) IsAdjusting() T.Gboolean {
	return ColorSelectionIsAdjusting(c)
}

func (c *ColorSelection) SetColor(color *float64) {
	ColorSelectionSetColor(c, color)
}

func (c *ColorSelection) GetColor(color *float64) {
	ColorSelectionGetColor(c, color)
}

func (c *ColorSelection) SetUpdatePolicy(policy T.GtkUpdateType) {
	ColorSelectionSetUpdatePolicy(c, policy)
}

type ColorSelectionDialog struct {
	Parent       Dialog
	Colorsel     *T.GtkWidget
	OkButton     *T.GtkWidget
	CancelButton *T.GtkWidget
	HelpButton   *T.GtkWidget
}

var (
	ColorSelectionDialogGetType func() T.GType
	ColorSelectionDialogNew     func(title string) *T.GtkWidget

	ColorSelectionDialogGetColorSelection func(c *ColorSelectionDialog) *T.GtkWidget
)

type Combo struct {
	Hbox          T.GtkHBox
	Entry         *T.GtkWidget
	Button        *T.GtkWidget
	Popup         *T.GtkWidget
	Popwin        *T.GtkWidget
	List          *T.GtkWidget
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
	ComboNew     func() *T.GtkWidget
)

var (
	ComboSetValueInList     func(c *Combo, val T.Gboolean, okIfEmpty T.Gboolean)
	ComboSetUseArrows       func(c *Combo, val T.Gboolean)
	ComboSetUseArrowsAlways func(c *Combo, val T.Gboolean)
	ComboSetCaseSensitive   func(c *Combo, val T.Gboolean)
	ComboSetItemString      func(c *Combo, item *Item, itemValue string)
	ComboSetPopdownStrings  func(c *Combo, strings *T.GList)
	ComboDisableActivate    func(c *Combo)
)

func (c *Combo) SetValueInList(
	val T.Gboolean, okIfEmpty T.Gboolean) {
	ComboSetValueInList(c, val, okIfEmpty)
}

func (c *Combo) SetUseArrows(val T.Gboolean) {
	ComboSetUseArrows(c, val)
}

func (c *Combo) SetUseArrowsAlways(val T.Gboolean) {
	ComboSetUseArrowsAlways(c, val)
}

func (c *Combo) SetCaseSensitive(val T.Gboolean) {
	ComboSetCaseSensitive(c, val)
}

func (c *Combo) SetItemString(item *Item, itemValue string) {
	ComboSetItemString(c, item, itemValue)
}

func (c *Combo) SetPopdownStrings(strings *T.GList) {
	ComboSetPopdownStrings(c, strings)
}

func (c *Combo) DisableActivate() { ComboDisableActivate(c) }

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
	ComboBoxNew                  func() *T.GtkWidget
	ComboBoxNewWithEntry         func() *T.GtkWidget
	ComboBoxNewWithModel         func(model *T.GtkTreeModel) *T.GtkWidget
	ComboBoxNewWithModelAndEntry func(model *T.GtkTreeModel) *T.GtkWidget

	ComboBoxTextGetType      func() T.GType
	ComboBoxTextNew          func() *T.GtkWidget
	ComboBoxTextNewWithEntry func() *T.GtkWidget

	ComboBoxNewText func() *T.GtkWidget

	ComboBoxEntryGetType      func() T.GType
	ComboBoxEntryNew          func() *T.GtkWidget
	ComboBoxEntryNewWithModel func(model *T.GtkTreeModel, textColumn int) *T.GtkWidget
	ComboBoxEntryNewText      func() *T.GtkWidget
)

var (
	ComboBoxAppendText           func(c *ComboBox, text string)
	ComboBoxGetActive            func(c *ComboBox) int
	ComboBoxGetActiveIter        func(c *ComboBox, iter *T.GtkTreeIter) T.Gboolean
	ComboBoxGetActiveText        func(c *ComboBox) string
	ComboBoxGetAddTearoffs       func(c *ComboBox) T.Gboolean
	ComboBoxGetButtonSensitivity func(c *ComboBox) T.GtkSensitivityType
	ComboBoxGetColumnSpanColumn  func(c *ComboBox) int
	ComboBoxGetEntryTextColumn   func(c *ComboBox) int
	ComboBoxGetFocusOnClick      func(c *ComboBox) T.Gboolean
	ComboBoxGetHasEntry          func(c *ComboBox) T.Gboolean
	ComboBoxGetModel             func(c *ComboBox) *T.GtkTreeModel
	ComboBoxGetPopupAccessible   func(c *ComboBox) *T.AtkObject
	ComboBoxGetRowSeparatorFunc  func(c *ComboBox) T.GtkTreeViewRowSeparatorFunc
	ComboBoxGetRowSpanColumn     func(c *ComboBox) int
	ComboBoxGetTitle             func(c *ComboBox) string
	ComboBoxGetWrapWidth         func(c *ComboBox) int
	ComboBoxInsertText           func(c *ComboBox, position int, text string)
	ComboBoxPopdown              func(c *ComboBox)
	ComboBoxPopup                func(c *ComboBox)
	ComboBoxPrependText          func(c *ComboBox, text string)
	ComboBoxRemoveText           func(c *ComboBox, position int)
	ComboBoxSetActive            func(c *ComboBox, index int)
	ComboBoxSetActiveIter        func(c *ComboBox, iter *T.GtkTreeIter)
	ComboBoxSetAddTearoffs       func(c *ComboBox, addTearoffs T.Gboolean)
	ComboBoxSetButtonSensitivity func(c *ComboBox, sensitivity T.GtkSensitivityType)
	ComboBoxSetColumnSpanColumn  func(c *ComboBox, columnSpan int)
	ComboBoxSetEntryTextColumn   func(c *ComboBox, textColumn int)
	ComboBoxSetFocusOnClick      func(c *ComboBox, focusOnClick T.Gboolean)
	ComboBoxSetModel             func(c *ComboBox, model *T.GtkTreeModel)
	ComboBoxSetRowSeparatorFunc  func(c *ComboBox, f T.GtkTreeViewRowSeparatorFunc, dataGpointer, destroy T.GDestroyNotify)
	ComboBoxSetRowSpanColumn     func(c *ComboBox, rowSpan int)
	ComboBoxSetTitle             func(c *ComboBox, title string)
	ComboBoxSetWrapWidth         func(c *ComboBox, width int)
)

func (c *ComboBox) GetWrapWidth() int {
	return ComboBoxGetWrapWidth(c)
}

func (c *ComboBox) SetWrapWidth(width int) {
	ComboBoxSetWrapWidth(c, width)
}

func (c *ComboBox) GetRowSpanColumn() int {
	return ComboBoxGetRowSpanColumn(c)
}

func (c *ComboBox) SetRowSpanColumn(rowSpan int) {
	ComboBoxSetRowSpanColumn(c, rowSpan)
}

func (c *ComboBox) GetColumnSpanColumn() int {
	return ComboBoxGetColumnSpanColumn(c)
}

func (c *ComboBox) SetColumnSpanColumn(columnSpan int) {
	ComboBoxSetColumnSpanColumn(c, columnSpan)
}

func (c *ComboBox) GetAddTearoffs() T.Gboolean {
	return ComboBoxGetAddTearoffs(c)
}

func (c *ComboBox) SetAddTearoffs(addTearoffs T.Gboolean) {
	ComboBoxSetAddTearoffs(c, addTearoffs)
}

func (c *ComboBox) GetTitle() string {
	return ComboBoxGetTitle(c)
}

func (c *ComboBox) SetTitle(title string) {
	ComboBoxSetTitle(c, title)
}

func (c *ComboBox) GetFocusOnClick() T.Gboolean {
	return ComboBoxGetFocusOnClick(c)
}

func (c *ComboBox) SetFocusOnClick(focusOnClick T.Gboolean) {
	ComboBoxSetFocusOnClick(c, focusOnClick)
}

func (c *ComboBox) GetActive() int { return ComboBoxGetActive(c) }

func (c *ComboBox) SetActive(index int) {
	ComboBoxSetActive(c, index)
}

func (c *ComboBox) GetActiveIter(iter *T.GtkTreeIter) T.Gboolean {
	return ComboBoxGetActiveIter(c, iter)
}

func (c *ComboBox) SetActiveIter(iter *T.GtkTreeIter) {
	ComboBoxSetActiveIter(c, iter)
}

func (c *ComboBox) SetModel(model *T.GtkTreeModel) {
	ComboBoxSetModel(c, model)
}

func (c *ComboBox) GetModel() *T.GtkTreeModel {
	return ComboBoxGetModel(c)
}

func (c *ComboBox) GetRowSeparatorFunc() T.GtkTreeViewRowSeparatorFunc {
	return ComboBoxGetRowSeparatorFunc(c)
}

func (c *ComboBox) SetRowSeparatorFunc(f T.GtkTreeViewRowSeparatorFunc, dataGpointer, destroy T.GDestroyNotify) {
	ComboBoxSetRowSeparatorFunc(c, f, dataGpointer, destroy)
}

func (c *ComboBox) SetButtonSensitivity(sensitivity T.GtkSensitivityType) {
	ComboBoxSetButtonSensitivity(c, sensitivity)
}

func (c *ComboBox) GetButtonSensitivity() T.GtkSensitivityType {
	return ComboBoxGetButtonSensitivity(c)
}

func (c *ComboBox) GetHasEntry() T.Gboolean {
	return ComboBoxGetHasEntry(c)
}

func (c *ComboBox) SetEntryTextColumn(textColumn int) {
	ComboBoxSetEntryTextColumn(c, textColumn)
}

func (c *ComboBox) GetEntryTextColumn() int {
	return ComboBoxGetEntryTextColumn(c)
}

func (c *ComboBox) AppendText(text string) {
	ComboBoxAppendText(c, text)
}

func (c *ComboBox) InsertText(position int, text string) {
	ComboBoxInsertText(c, position, text)
}

func (c *ComboBox) PrependText(text string) {
	ComboBoxPrependText(c, text)
}

func (c *ComboBox) RemoveText(position int) {
	ComboBoxRemoveText(c, position)
}

func (c *ComboBox) GetActiveText() string {
	return ComboBoxGetActiveText(c)
}

func (c *ComboBox) Popup() { ComboBoxPopup(c) }

func (c *ComboBox) Popdown() { ComboBoxPopdown(c) }

func (c *ComboBox) GetPopupAccessible() *T.AtkObject {
	return ComboBoxGetPopupAccessible(c)
}

var (
	ComboBoxTextAppendText    func(c *ComboBoxText, text string)
	ComboBoxTextGetActiveText func(c *ComboBoxText) string
	ComboBoxTextInsertText    func(c *ComboBoxText, position int, text string)
	ComboBoxTextPrependText   func(c *ComboBoxText, text string)
	ComboBoxTextRemove        func(c *ComboBoxText, position int)
)

func (c *ComboBoxText) AppendText(text string) {
	ComboBoxTextAppendText(c, text)
}

func (c *ComboBoxText) InsertText(position int, text string) {
	ComboBoxTextInsertText(c, position, text)
}

func (c *ComboBoxText) PrependText(text string) {
	ComboBoxTextPrependText(c, text)
}

func (c *ComboBoxText) Remove(position int) {
	ComboBoxTextRemove(c, position)
}

func (c *ComboBoxText) GetActiveText() string {
	return ComboBoxTextGetActiveText(c)
}

var (
	ComboBoxEntrySetTextColumn func(e *ComboBoxEntry, textColumn int)
	ComboBoxEntryGetTextColumn func(e *ComboBoxEntry) int
)

func (e *ComboBoxEntry) SetTextColumn(textColumn int) {
	ComboBoxEntrySetTextColumn(e, textColumn)
}

func (e *ComboBoxEntry) GetTextColumn() int {
	return ComboBoxEntryGetTextColumn(e)
}

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
		Pixmap_closed *T.GdkPixmap
		Mask_closed   *T.GdkBitmap
		Pixmap_opened *T.GdkPixmap
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

type CTreeExpanderStyle T.Enum

const (
	CTREE_EXPANDER_NONE CTreeExpanderStyle = iota
	CTREE_EXPANDER_SQUARE
	CTREE_EXPANDER_TRIANGLE
	CTREE_EXPANDER_CIRCULAR
)

type CTreeLineStyle T.Enum

const (
	CTREE_LINES_NONE CTreeLineStyle = iota
	CTREE_LINES_SOLID
	CTREE_LINES_DOTTED
	CTREE_LINES_TABBED
)

var (
	CtreeGetType       func() T.GType
	CtreeNew           func(columns int, treeColumn int) *T.GtkWidget
	CtreeNewWithTitles func(columns int, treeColumn int, titles **T.Gchar) *T.GtkWidget

	CtreeNodeGetType func() T.GType
	CtreePosGetType  func() T.GType

	CtreeExpanderStyleGetType func() T.GType
	CtreeExpansionTypeGetType func() T.GType
	CtreeLineStyleGetType     func() T.GType
)

var (
	CtreeCollapse                 func(c *CTree, node *CTreeNode)
	CtreeCollapseRecursive        func(c *CTree, node *CTreeNode)
	CtreeCollapseToDepth          func(c *CTree, node *CTreeNode, depth int)
	CtreeExpand                   func(c *CTree, node *CTreeNode)
	CtreeExpandRecursive          func(c *CTree, node *CTreeNode)
	CtreeExpandToDepth            func(c *CTree, node *CTreeNode, depth int)
	CtreeExportToGnode            func(c *CTree, parent, sibling *T.GNode, node *CTreeNode, f CTreeGNodeFunc, data T.Gpointer) *T.GNode
	CtreeFind                     func(c *CTree, node *CTreeNode, child *CTreeNode) T.Gboolean
	CtreeFindAllByRowData         func(c *CTree, node *CTreeNode, data T.Gpointer) *T.GList
	CtreeFindAllByRowDataCustom   func(c *CTree, node *CTreeNode, dataGpointer, f T.GCompareFunc) *T.GList
	CtreeFindByRowData            func(c *CTree, node *CTreeNode, data T.Gpointer) *CTreeNode
	CtreeFindByRowDataCustom      func(c *CTree, node *CTreeNode, dataGpointer, f T.GCompareFunc) *CTreeNode
	CtreeFindNodePtr              func(c *CTree, ctreeRow *CTreeRow) *CTreeNode
	CtreeGetNodeInfo              func(c *CTree, node *CTreeNode, text **T.Gchar, spacing *uint8, pixmapClosed **T.GdkPixmap, maskClosed **T.GdkBitmap, pixmapOpened **T.GdkPixmap, maskOpened **T.GdkBitmap, isLeaf, expanded *T.Gboolean) T.Gboolean
	CtreeInsertGnode              func(c *CTree, parent *CTreeNode, sibling *CTreeNode, gnode *T.GNode, f CTreeGNodeFunc, data T.Gpointer) *CTreeNode
	CtreeInsertNode               func(c *CTree, parent *CTreeNode, sibling *CTreeNode, text **T.Gchar, spacing uint8, pixmapClosed *T.GdkPixmap, maskClosed *T.GdkBitmap, pixmapOpened *T.GdkPixmap, maskOpened *T.GdkBitmap, isLeaf T.Gboolean, expanded T.Gboolean) *CTreeNode
	CtreeIsAncestor               func(c *CTree, node *CTreeNode, child *CTreeNode) T.Gboolean
	CtreeIsHotSpot                func(c *CTree, x int, y int) T.Gboolean
	CtreeIsViewable               func(c *CTree, node *CTreeNode) T.Gboolean
	CtreeLast                     func(c *CTree, node *CTreeNode) *CTreeNode
	CtreeMove                     func(c *CTree, node *CTreeNode, newParent *CTreeNode, newSibling *CTreeNode)
	CtreeNodeGetCellStyle         func(c *CTree, node *CTreeNode, column int) *T.GtkStyle
	CtreeNodeGetCellType          func(c *CTree, node *CTreeNode, column int) T.GtkCellType
	CtreeNodeGetPixmap            func(c *CTree, node *CTreeNode, column int, pixmap **T.GdkPixmap, mask **T.GdkBitmap) T.Gboolean
	CtreeNodeGetPixtext           func(c *CTree, node *CTreeNode, column int, text **T.Gchar, spacing *uint8, pixmap **T.GdkPixmap, mask **T.GdkBitmap) T.Gboolean
	CtreeNodeGetRowData           func(c *CTree, node *CTreeNode) T.Gpointer
	CtreeNodeGetRowStyle          func(c *CTree, node *CTreeNode) *T.GtkStyle
	CtreeNodeGetSelectable        func(c *CTree, node *CTreeNode) T.Gboolean
	CtreeNodeGetText              func(c *CTree, node *CTreeNode, column int, text **T.Gchar) T.Gboolean
	CtreeNodeIsVisible            func(c *CTree, node *CTreeNode) T.GtkVisibility
	CtreeNodeMoveto               func(c *CTree, node *CTreeNode, column int, rowAlign, colAlign float32)
	CtreeNodeNth                  func(c *CTree, row uint) *CTreeNode
	CtreeNodeSetBackground        func(c *CTree, node *CTreeNode, color *T.GdkColor)
	CtreeNodeSetCellStyle         func(c *CTree, node *CTreeNode, column int, style *T.GtkStyle)
	CtreeNodeSetForeground        func(c *CTree, node *CTreeNode, color *T.GdkColor)
	CtreeNodeSetPixmap            func(c *CTree, node *CTreeNode, column int, pixmap *T.GdkPixmap, mask *T.GdkBitmap)
	CtreeNodeSetPixtext           func(c *CTree, node *CTreeNode, column int, text string, spacing uint8, pixmap *T.GdkPixmap, mask *T.GdkBitmap)
	CtreeNodeSetRowData           func(c *CTree, node *CTreeNode, data T.Gpointer)
	CtreeNodeSetRowDataFull       func(c *CTree, node *CTreeNode, dataGpointer, destroy T.GDestroyNotify)
	CtreeNodeSetRowStyle          func(c *CTree, node *CTreeNode, style *T.GtkStyle)
	CtreeNodeSetSelectable        func(c *CTree, node *CTreeNode, selectable T.Gboolean)
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
	CtreeSetNodeInfo              func(c *CTree, node *CTreeNode, text string, spacing uint8, pixmapClosed *T.GdkPixmap, maskClosed *T.GdkBitmap, pixmapOpened *T.GdkPixmap, maskOpened *T.GdkBitmap, isLeaf, expanded T.Gboolean)
	CtreeSetShowStub              func(c *CTree, showStub T.Gboolean)
	CtreeSetSpacing               func(c *CTree, spacing int)
	CtreeSortNode                 func(c *CTree, node *CTreeNode)
	CtreeSortRecursive            func(c *CTree, node *CTreeNode)
	CtreeToggleExpansion          func(c *CTree, node *CTreeNode)
	CtreeToggleExpansionRecursive func(c *CTree, node *CTreeNode)
	CtreeUnselect                 func(c *CTree, node *CTreeNode)
	CtreeUnselectRecursive        func(c *CTree, node *CTreeNode)
)

func (c *CTree) InsertNode(parent, sibling *CTreeNode,
	text **T.Gchar, spacing uint8, pixmapClosed *T.GdkPixmap,
	maskClosed *T.GdkBitmap, pixmapOpened *T.GdkPixmap,
	maskOpened *T.GdkBitmap, isLeaf,
	expanded T.Gboolean) *CTreeNode {
	return CtreeInsertNode(c, parent, sibling, text,
		spacing, pixmapClosed, maskClosed, pixmapOpened,
		maskOpened, isLeaf, expanded)
}

func (c *CTree) RemoveNode(node *CTreeNode) {
	CtreeRemoveNode(c, node)
}

func (c *CTree) InsertGnode(parent, sibling *CTreeNode,
	gnode *T.GNode, f CTreeGNodeFunc,
	data T.Gpointer) *CTreeNode {
	return CtreeInsertGnode(c, parent, sibling, gnode, f, data)
}

func (c *CTree) ExportToGnode(parent *T.GNode, sibling *T.GNode,
	node *CTreeNode, f CTreeGNodeFunc,
	data T.Gpointer) *T.GNode {
	return CtreeExportToGnode(c, parent, sibling, node, f, data)
}

func (c *CTree) PostRecursive(node *CTreeNode,
	f CTreeFunc, data T.Gpointer) {
	CtreePostRecursive(c, node, f, data)
}

func (c *CTree) PostRecursiveToDepth(node *CTreeNode,
	depth int, f CTreeFunc, data T.Gpointer) {
	CtreePostRecursiveToDepth(c, node, depth, f, data)
}

func (c *CTree) PreRecursive(node *CTreeNode,
	f CTreeFunc, data T.Gpointer) {
	CtreePreRecursive(c, node, f, data)
}

func (c *CTree) PreRecursiveToDepth(node *CTreeNode,
	depth int, f CTreeFunc, data T.Gpointer) {
	CtreePreRecursiveToDepth(c, node, depth, f, data)
}

func (c *CTree) IsViewable(node *CTreeNode) T.Gboolean {
	return CtreeIsViewable(c, node)
}

func (c *CTree) Last(node *CTreeNode) *CTreeNode {
	return CtreeLast(c, node)
}

func (c *CTree) FindNodePtr(
	ctreeRow *CTreeRow) *CTreeNode {
	return CtreeFindNodePtr(c, ctreeRow)
}

func (c *CTree) NodeNth(row uint) *CTreeNode {
	return CtreeNodeNth(c, row)
}

func (c *CTree) Find(node, child *CTreeNode) T.Gboolean {
	return CtreeFind(c, node, child)
}

func (c *CTree) IsAncestor(
	node, child *CTreeNode) T.Gboolean {
	return CtreeIsAncestor(c, node, child)
}

func (c *CTree) FindByRowData(
	node *CTreeNode, data T.Gpointer) *CTreeNode {
	return CtreeFindByRowData(c, node, data)
}

func (c *CTree) FindAllByRowData(
	node *CTreeNode, data T.Gpointer) *T.GList {
	return CtreeFindAllByRowData(c, node, data)
}

func (c *CTree) FindByRowDataCustom(node *CTreeNode,
	dataGpointer, f T.GCompareFunc) *CTreeNode {
	return CtreeFindByRowDataCustom(c, node, dataGpointer, f)
}

func (c *CTree) FindAllByRowDataCustom(node *CTreeNode,
	dataGpointer, f T.GCompareFunc) *T.GList {
	return CtreeFindAllByRowDataCustom(c, node, dataGpointer, f)
}

func (c *CTree) IsHotSpot(x int, y int) T.Gboolean {
	return CtreeIsHotSpot(c, x, y)
}

func (c *CTree) Move(node, newParent, newSibling *CTreeNode) {
	CtreeMove(c, node, newParent, newSibling)
}

func (c *CTree) Expand(node *CTreeNode) {
	CtreeExpand(c, node)
}

func (c *CTree) ExpandRecursive(node *CTreeNode) {
	CtreeExpandRecursive(c, node)
}

func (c *CTree) ExpandToDepth(node *CTreeNode, depth int) {
	CtreeExpandToDepth(c, node, depth)
}

func (c *CTree) Collapse(node *CTreeNode) {
	CtreeCollapse(c, node)
}

func (c *CTree) CollapseRecursive(node *CTreeNode) {
	CtreeCollapseRecursive(c, node)
}

func (c *CTree) CollapseToDepth(node *CTreeNode, depth int) {
	CtreeCollapseToDepth(c, node, depth)
}

func (c *CTree) ToggleExpansion(node *CTreeNode) {
	CtreeToggleExpansion(c, node)
}

func (c *CTree) ToggleExpansionRecursive(node *CTreeNode) {
	CtreeToggleExpansionRecursive(c, node)
}

func (c *CTree) Select(node *CTreeNode) {
	CtreeSelect(c, node)
}

func (c *CTree) SelectRecursive(node *CTreeNode) {
	CtreeSelectRecursive(c, node)
}

func (c *CTree) Unselect(node *CTreeNode) {
	CtreeUnselect(c, node)
}

func (c *CTree) UnselectRecursive(node *CTreeNode) {
	CtreeUnselectRecursive(c, node)
}

func (c *CTree) RealSelectRecursive(
	node *CTreeNode, state int) {
	CtreeRealSelectRecursive(c, node, state)
}

func (c *CTree) NodeSetText(
	node *CTreeNode, column int, text string) {
	CtreeNodeSetText(c, node, column, text)
}

func (c *CTree) NodeSetPixmap(node *CTreeNode, column int, pixmap *T.GdkPixmap, mask *T.GdkBitmap) {
	CtreeNodeSetPixmap(c, node, column, pixmap, mask)
}

func (c *CTree) NodeSetPixtext(node *CTreeNode, column int,
	text string, spacing uint8, pixmap *T.GdkPixmap,
	mask *T.GdkBitmap) {
	CtreeNodeSetPixtext(
		c, node, column, text, spacing, pixmap, mask)
}

func (c *CTree) SetNodeInfo(node *CTreeNode, text string,
	spacing uint8, pixmapClosed *T.GdkPixmap,
	maskClosed *T.GdkBitmap, pixmapOpened *T.GdkPixmap,
	maskOpened *T.GdkBitmap, isLeaf, expanded T.Gboolean) {
	CtreeSetNodeInfo(c, node, text, spacing, pixmapClosed, maskClosed, pixmapOpened, maskOpened, isLeaf, expanded)
}

func (c *CTree) NodeSetShift(
	node *CTreeNode, column, vertical, horizontal int) {
	CtreeNodeSetShift(c, node, column, vertical, horizontal)
}

func (c *CTree) NodeSetSelectable(
	node *CTreeNode, selectable T.Gboolean) {
	CtreeNodeSetSelectable(c, node, selectable)
}

func (c *CTree) NodeGetSelectable(
	node *CTreeNode) T.Gboolean {
	return CtreeNodeGetSelectable(c, node)
}

func (c *CTree) NodeGetCellType(
	node *CTreeNode, column int) T.GtkCellType {
	return CtreeNodeGetCellType(c, node, column)
}

func (c *CTree) NodeGetText(node *CTreeNode,
	column int, text **T.Gchar) T.Gboolean {
	return CtreeNodeGetText(c, node, column, text)
}

func (c *CTree) NodeGetPixmap(node *CTreeNode, column int,
	pixmap **T.GdkPixmap, mask **T.GdkBitmap) T.Gboolean {
	return CtreeNodeGetPixmap(c, node, column, pixmap, mask)
}

func (c *CTree) NodeGetPixtext(node *CTreeNode, column int,
	text **T.Gchar, spacing *uint8, pixmap **T.GdkPixmap,
	mask **T.GdkBitmap) T.Gboolean {
	return CtreeNodeGetPixtext(
		c, node, column, text, spacing, pixmap, mask)
}

func (c *CTree) GetNodeInfo(node *CTreeNode, text **T.Gchar,
	spacing *uint8, pixmapClosed **T.GdkPixmap,
	maskClosed **T.GdkBitmap, pixmapOpened **T.GdkPixmap,
	maskOpened **T.GdkBitmap, isLeaf,
	expanded *T.Gboolean) T.Gboolean {
	return CtreeGetNodeInfo(c, node, text, spacing, pixmapClosed,
		maskClosed, pixmapOpened, maskOpened, isLeaf, expanded)
}

func (c *CTree) NodeSetRowStyle(
	node *CTreeNode, style *T.GtkStyle) {
	CtreeNodeSetRowStyle(c, node, style)
}

func (c *CTree) NodeGetRowStyle(
	node *CTreeNode) *T.GtkStyle {
	return CtreeNodeGetRowStyle(c, node)
}

func (c *CTree) NodeSetCellStyle(
	node *CTreeNode, column int, style *T.GtkStyle) {
	CtreeNodeSetCellStyle(c, node, column, style)
}

func (c *CTree) NodeGetCellStyle(
	node *CTreeNode, column int) *T.GtkStyle {
	return CtreeNodeGetCellStyle(c, node, column)
}

func (c *CTree) NodeSetForeground(
	node *CTreeNode, color *T.GdkColor) {
	CtreeNodeSetForeground(c, node, color)
}

func (c *CTree) NodeSetBackground(
	node *CTreeNode, color *T.GdkColor) {
	CtreeNodeSetBackground(c, node, color)
}

func (c *CTree) NodeSetRowData(
	node *CTreeNode, data T.Gpointer) {
	CtreeNodeSetRowData(c, node, data)
}

func (c *CTree) NodeSetRowDataFull(node *CTreeNode,
	dataGpointer, destroy T.GDestroyNotify) {
	CtreeNodeSetRowDataFull(c, node, dataGpointer, destroy)
}

func (c *CTree) NodeGetRowData(node *CTreeNode) T.Gpointer {
	return CtreeNodeGetRowData(c, node)
}

func (c *CTree) NodeMoveto(node *CTreeNode,
	column int, rowAlign, colAlign float32) {
	CtreeNodeMoveto(c, node, column, rowAlign, colAlign)
}

func (c *CTree) NodeIsVisible(
	node *CTreeNode) T.GtkVisibility {
	return CtreeNodeIsVisible(c, node)
}

func (c *CTree) SetIndent(indent int) {
	CtreeSetIndent(c, indent)
}

func (c *CTree) SetSpacing(spacing int) {
	CtreeSetSpacing(c, spacing)
}

func (c *CTree) SetShowStub(showStub T.Gboolean) {
	CtreeSetShowStub(c, showStub)
}

func (c *CTree) SetLineStyle(lineStyle CTreeLineStyle) {
	CtreeSetLineStyle(c, lineStyle)
}

func (c *CTree) SetExpanderStyle(
	expanderStyle CTreeExpanderStyle) {
	CtreeSetExpanderStyle(c, expanderStyle)
}

func (c *CTree) SetDragCompareFunc(
	cmpFunc CTreeCompareDragFunc) {
	CtreeSetDragCompareFunc(c, cmpFunc)
}

func (c *CTree) SortNode(node *CTreeNode) {
	CtreeSortNode(c, node)
}

func (c *CTree) SortRecursive(node *CTreeNode) {
	CtreeSortRecursive(c, node)
}

type Curve struct {
	Graph        DrawingArea
	CursorType   int
	MinX         float32
	MaxX         float32
	MinY         float32
	MaxY         float32
	Pixmap       *T.GdkPixmap
	CurveType    T.GtkCurveType
	Height       int
	GrabPoint    int
	Last         int
	NumPoints    int
	Point        *T.GdkPoint
	NumCtlpoints int
	Ctlpoint     [2]*float32 //TODO(t): float32 (*ctlpoint)[2]; ???
}

var (
	CurveGetType     func() T.GType
	CurveNew         func() *T.GtkWidget
	CurveTypeGetType func() T.GType

	CurveGetVector    func(c *Curve, veclen int, vector *float32)
	CurveReset        func(c *Curve)
	CurveSetCurveType func(c *Curve, t T.GtkCurveType)
	CurveSetGamma     func(c *Curve, gamma float32)
	CurveSetRange     func(c *Curve, minX, maxX, minY, maxY float32)
	CurveSetVector    func(c *Curve, veclen int, vector *float32)
)

func (c *Curve) Reset() { CurveReset(c) }

func (c *Curve) SetGamma(gamma float32) {
	CurveSetGamma(c, gamma)
}

func (c *Curve) SetRange(minX, maxX, minY, maxY float32) {
	CurveSetRange(c, minX, maxX, minY, maxY)
}

func (c *Curve) GetVector(veclen int, vector *float32) {
	CurveGetVector(c, veclen, vector)
}

func (c *Curve) SetVector(veclen int, vector *float32) {
	CurveSetVector(c, veclen, vector)
}

func (c *Curve) SetCurveType(t T.GtkCurveType) {
	CurveSetCurveType(c, t)
}
