package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Calendar struct {
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
	Display_flags     T.GtkCalendarDisplayOptions
	Marked_date_color [31]T.GdkColor
	Gc                *T.GdkGC
	Xor_gc            *T.GdkGC
	Focus_row         int
	Focus_col         int
	Highlight_row     int
	Highlight_col     int
	Priv              *T.GtkCalendarPrivate
	Grow_space        [32]T.Gchar
	_, _, _, _        func()
}

var (
	CalendarDisplayOptionsGetType func() T.GType
	CalendarGetType               func() T.GType
	CalendarNew                   func() *T.GtkWidget
)

var (
	CalendarClearMarks          func(c *Calendar)
	CalendarDisplayOptions      func(c *Calendar, flags T.GtkCalendarDisplayOptions)
	CalendarFreeze              func(c *Calendar)
	CalendarGetDate             func(c *Calendar, year, month, day *uint)
	CalendarGetDetailHeightRows func(c *Calendar) int
	CalendarGetDetailWidthChars func(c *Calendar) int
	CalendarGetDisplayOptions   func(c *Calendar) T.GtkCalendarDisplayOptions
	CalendarMarkDay             func(c *Calendar, day uint) T.Gboolean
	CalendarSelectDay           func(c *Calendar, day uint)
	CalendarSelectMonth         func(c *Calendar, month, year uint) T.Gboolean
	CalendarSetDetailFunc       func(c *Calendar, f T.GtkCalendarDetailFunc, dataGpointer, destroy T.GDestroyNotify)
	CalendarSetDetailHeightRows func(c *Calendar, rows int)
	CalendarSetDetailWidthChars func(c *Calendar, chars int)
	CalendarSetDisplayOptions   func(c *Calendar, flags T.GtkCalendarDisplayOptions)
	CalendarThaw                func(c *Calendar)
	CalendarUnmarkDay           func(c *Calendar, day uint) T.Gboolean
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

func (c *Calendar) SetDisplayOptions(flags T.GtkCalendarDisplayOptions) {
	CalendarSetDisplayOptions(c, flags)
}

func (c *Calendar) GetDisplayOptions() T.GtkCalendarDisplayOptions {
	return CalendarGetDisplayOptions(c)
}

func (c *Calendar) DisplayOptions(flags T.GtkCalendarDisplayOptions) {
	CalendarDisplayOptions(c, flags)
}

func (c *Calendar) GetDate(year, month, day *uint) {
	CalendarGetDate(c, year, month, day)
}

func (c *Calendar) SetDetailFunc(f T.GtkCalendarDetailFunc, dataGpointer, destroy T.GDestroyNotify) {
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

//============================================================

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
	CellRendererAccelNew         func() *T.GtkCellRenderer
	CellRendererAccelModeGetType func() T.GType
	CellRendererComboGetType     func() T.GType
	CellRendererComboNew         func() *T.GtkCellRenderer
	CellRendererModeGetType      func() T.GType
	CellRendererPixbufGetType    func() T.GType
	CellRendererPixbufNew        func() *T.GtkCellRenderer
	CellRendererProgressGetType  func() T.GType
	CellRendererProgressNew      func() *T.GtkCellRenderer
	CellRendererSpinGetType      func() T.GType
	CellRendererSpinNew          func() *T.GtkCellRenderer
	CellRendererSpinnerGetType   func() T.GType
	CellRendererSpinnerNew       func() *T.GtkCellRenderer
	CellRendererStateGetType     func() T.GType
	CellRendererToggleGetType    func() T.GType
	CellRendererToggleNew        func() *T.GtkCellRenderer
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
	CellRendererTextNew     func() *T.GtkCellRenderer

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

type CList struct {
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
	Column             *T.GtkCListColumn
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
	ClickCell          T.GtkCListCellInfo
	Hadjustment        *T.GtkAdjustment
	Vadjustment        *T.GtkAdjustment
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
	Compare            T.GtkCListCompareFunc
	SortColumn         int
	DragHighlightRow   int
	DragHighlightPos   T.GtkCListDragPos
}

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
	ClistGetHadjustment         func(c *CList) *T.GtkAdjustment
	ClistGetPixmap              func(c *CList, row, column int, pixmap, mask **T.GdkBitmap) int
	ClistGetPixtext             func(c *CList, row, column int, text **T.Gchar, spacing *uint8, pixmap, mask **T.GdkBitmap) int
	ClistGetRowData             func(c *CList, row int) T.Gpointer
	ClistGetRowStyle            func(c *CList, row int) *T.GtkStyle
	ClistGetSelectable          func(c *CList, row int) T.Gboolean
	ClistGetSelectionInfo       func(c *CList, x, y int, row, column *int) int
	ClistGetText                func(c *CList, row, column int, text **T.Gchar) int
	ClistGetVadjustment         func(c *CList) *T.GtkAdjustment
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
	ClistSetCompareFunc         func(c *CList, cmpFunc T.GtkCListCompareFunc)
	ClistSetForeground          func(c *CList, row int, color *T.GdkColor)
	ClistSetHadjustment         func(c *CList, adjustment *T.GtkAdjustment)
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
	ClistSetVadjustment         func(c *CList, adjustment *T.GtkAdjustment)
	ClistSort                   func(c *CList)
	ClistSwapRows               func(c *CList, row1, row2 int)
	ClistThaw                   func(c *CList)
	ClistUndoSelection          func(c *CList)
	ClistUnselectAll            func(c *CList)
	ClistUnselectRow            func(c *CList, row, column int)
)

func (c *CList) SetHadjustment(adjustment *T.GtkAdjustment) {
	ClistSetHadjustment(c, adjustment)
}

func (c *CList) SetVadjustment(adjustment *T.GtkAdjustment) {
	ClistSetVadjustment(c, adjustment)
}

func (c *CList) GetHadjustment() *T.GtkAdjustment {
	return ClistGetHadjustment(c)
}

func (c *CList) GetVadjustment() *T.GtkAdjustment {
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

func (c *CList) SetCompareFunc(cmpFunc T.GtkCListCompareFunc) {
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

//============================================================

type ComboBox struct {
	ParentInstance T.GtkBin
	_              *T.GtkComboBoxPrivate
}

type ComboBoxText struct {
	ParentInstance *ComboBox
	_              *T.GtkComboBoxTextPrivate
}

type ComboBoxEntry struct {
	ParentInstance ComboBox
	_              *T.GtkComboBoxEntryPrivate
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
