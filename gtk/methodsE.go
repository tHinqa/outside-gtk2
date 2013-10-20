package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Editable struct{}

var (
	EditableGetType func() T.GType

	EditableCopyClipboard      func(e *Editable)
	EditableCutClipboard       func(e *Editable)
	EditableDeleteSelection    func(e *Editable)
	EditableDeleteText         func(e *Editable, startPos, endPos int)
	EditableGetChars           func(e *Editable, startPos, endPos int) string
	EditableGetEditable        func(e *Editable) T.Gboolean
	EditableGetPosition        func(e *Editable) int
	EditableGetSelectionBounds func(e *Editable, startPos, endPos *int) T.Gboolean
	EditableInsertText         func(e *Editable, newText string, newTextLength int, position *int)
	EditablePasteClipboard     func(e *Editable)
	EditableSelectRegion       func(e *Editable, startPos, endPos int)
	EditableSetEditable        func(e *Editable, isEditable T.Gboolean)
	EditableSetPosition        func(e *Editable, position int)
)

func (e *Editable) SelectRegion(startPos, endPos int) {
	EditableSelectRegion(e, startPos, endPos)
}

func (e *Editable) GetSelectionBounds(startPos, endPos *int) T.Gboolean {
	return EditableGetSelectionBounds(e, startPos, endPos)
}

func (e *Editable) InsertText(newText string, newTextLength int, position *int) {
	EditableInsertText(e, newText, newTextLength, position)
}

func (e *Editable) DeleteText(startPos, endPos int) {
	EditableDeleteText(e, startPos, endPos)
}

func (e *Editable) GetChars(startPos, endPos int) string {
	return EditableGetChars(e, startPos, endPos)
}

func (e *Editable) CutClipboard() { EditableCutClipboard(e) }

func (e *Editable) CopyClipboard() { EditableCopyClipboard(e) }

func (e *Editable) PasteClipboard() { EditablePasteClipboard(e) }

func (e *Editable) DeleteSelection() { EditableDeleteSelection(e) }

func (e *Editable) SetPosition(position int) {
	EditableSetPosition(e, position)
}

func (e *Editable) GetPosition() int {
	return EditableGetPosition(e)
}

func (e *Editable) SetEditable(isEditable T.Gboolean) {
	EditableSetEditable(e, isEditable)
}

func (e *Editable) GetEditable() T.Gboolean {
	return EditableGetEditable(e)
}

type Entry struct {
	Widget T.GtkWidget
	Text   *T.Gchar
	Bits   uint
	// Editable : 1
	// Visible : 1
	// OverwriteMode : 1
	// InDrag : 1
	TextLength     uint16
	TextMaxLength  uint16
	TextArea       *T.GdkWindow
	ImContext      *T.GtkIMContext
	PopupMenu      *T.GtkWidget
	CurrentPos     int
	SelectionBound int
	CachedLayout   *T.PangoLayout
	Bits2          uint
	// CacheIncludesPreedit : 1
	// NeedImReset : 1
	// HasFrame : 1
	// ActivatesDefault : 1
	// CursorVisible : 1
	// InClick : 1
	// IsCellRenderer : 1
	// EditingCanceled : 1
	// MouseCursorObscured : 1
	// SelectWords : 1
	// SelectLines : 1
	// ResolvedDir : 4
	// TruncateMultiline : 1
	Button         uint
	Blink_timeout  uint
	Recompute_idle uint
	Scroll_offset  int
	Ascent         int
	Descent        int
	XTextSize      uint16
	XNBytes        uint16
	PreeditLength  uint16
	PreeditCursor  uint16
	DndPosition    int
	DragStartC     int
	DragStartY     int
	InvisibleChar  T.Gunichar
	WidthChars     int
}

type EntryBuffer struct {
	Parent T.GObject
	_      *struct{}
}

type EntryCompletion struct {
	Parent T.GObject
	_      *struct{}
}

type EntryIconPosition T.Enum

const (
	ENTRY_ICON_PRIMARY EntryIconPosition = iota
	ENTRY_ICON_SECONDARY
)

var (
	EntryGetType          func() T.GType
	EntryNew              func() *T.GtkWidget
	EntryNewWithMaxLength func(max int) *T.GtkWidget
	EntryNewWithBuffer    func(buffer *EntryBuffer) *T.GtkWidget

	EntryIconPositionGetType func() T.GType

	EntryAppendText               func(e *Entry, text string)
	EntryGetActivatesDefault      func(e *Entry) T.Gboolean
	EntryGetAlignment             func(e *Entry) float32
	EntryGetBuffer                func(e *Entry) *EntryBuffer
	EntryGetCompletion            func(e *Entry) *EntryCompletion
	EntryGetCurrentIconDragSource func(e *Entry) int
	EntryGetCursorHadjustment     func(e *Entry) *T.GtkAdjustment
	EntryGetHasFrame              func(e *Entry) T.Gboolean
	EntryGetIconActivatable       func(e *Entry, iconPos EntryIconPosition) T.Gboolean
	EntryGetIconAtPos             func(e *Entry, x, y int) int
	EntryGetIconGicon             func(e *Entry, iconPos EntryIconPosition) *T.GIcon
	EntryGetIconName              func(e *Entry, iconPos EntryIconPosition) string
	EntryGetIconPixbuf            func(e *Entry, iconPos EntryIconPosition) *T.GdkPixbuf
	EntryGetIconSensitive         func(e *Entry, iconPos EntryIconPosition) T.Gboolean
	EntryGetIconStock             func(e *Entry, iconPos EntryIconPosition) string
	EntryGetIconStorageType       func(e *Entry, iconPos EntryIconPosition) ImageType
	EntryGetIconTooltipMarkup     func(e *Entry, iconPos EntryIconPosition) string
	EntryGetIconTooltipText       func(e *Entry, iconPos EntryIconPosition) string
	EntryGetIconWindow            func(e *Entry, iconPos EntryIconPosition) *T.GdkWindow
	EntryGetInnerBorder           func(e *Entry) *Border
	EntryGetInvisibleChar         func(e *Entry) T.Gunichar
	EntryGetLayout                func(e *Entry) *T.PangoLayout
	EntryGetLayoutOffsets         func(e *Entry, x, y *int)
	EntryGetMaxLength             func(e *Entry) int
	EntryGetOverwriteMode         func(e *Entry) T.Gboolean
	EntryGetProgressFraction      func(e *Entry) float64
	EntryGetProgressPulseStep     func(e *Entry) float64
	EntryGetText                  func(e *Entry) string
	EntryGetTextLength            func(e *Entry) uint16
	EntryGetTextWindow            func(e *Entry) *T.GdkWindow
	EntryGetVisibility            func(e *Entry) T.Gboolean
	EntryGetWidthChars            func(e *Entry) int
	EntryImContextFilterKeypress  func(e *Entry, event *T.GdkEventKey) T.Gboolean
	EntryLayoutIndexToTextIndex   func(e *Entry, layoutIndex int) int
	EntryPrependText              func(e *Entry, text string)
	EntryProgressPulse            func(e *Entry)
	EntryResetImContext           func(e *Entry)
	EntrySelectRegion             func(e *Entry, start, end int)
	EntrySetActivatesDefault      func(e *Entry, setting T.Gboolean)
	EntrySetAlignment             func(e *Entry, xalign float32)
	EntrySetBuffer                func(e *Entry, buffer *EntryBuffer)
	EntrySetCompletion            func(e *Entry, completion *EntryCompletion)
	EntrySetCursorHadjustment     func(e *Entry, adjustment *T.GtkAdjustment)
	EntrySetEditable              func(e *Entry, editable T.Gboolean)
	EntrySetHasFrame              func(e *Entry, setting T.Gboolean)
	EntrySetIconActivatable       func(e *Entry, iconPos EntryIconPosition, activatable T.Gboolean)
	EntrySetIconDragSource        func(e *Entry, iconPos EntryIconPosition, targetList *T.GtkTargetList, actions T.GdkDragAction)
	EntrySetIconFromGicon         func(e *Entry, iconPos EntryIconPosition, icon *T.GIcon)
	EntrySetIconFromIconName      func(e *Entry, iconPos EntryIconPosition, iconName string)
	EntrySetIconFromPixbuf        func(e *Entry, iconPos EntryIconPosition, pixbuf *T.GdkPixbuf)
	EntrySetIconFromStock         func(e *Entry, iconPos EntryIconPosition, stockId string)
	EntrySetIconSensitive         func(e *Entry, iconPos EntryIconPosition, sensitive T.Gboolean)
	EntrySetIconTooltipMarkup     func(e *Entry, iconPos EntryIconPosition, tooltip string)
	EntrySetIconTooltipText       func(e *Entry, iconPos EntryIconPosition, tooltip string)
	EntrySetInnerBorder           func(e *Entry, border *Border)
	EntrySetInvisibleChar         func(e *Entry, ch T.Gunichar)
	EntrySetMaxLength             func(e *Entry, max int)
	EntrySetOverwriteMode         func(e *Entry, overwrite T.Gboolean)
	EntrySetPosition              func(e *Entry, position int)
	EntrySetProgressFraction      func(e *Entry, fraction float64)
	EntrySetProgressPulseStep     func(e *Entry, fraction float64)
	EntrySetText                  func(e *Entry, text string)
	EntrySetVisibility            func(e *Entry, visible T.Gboolean)
	EntrySetWidthChars            func(e *Entry, nChars int)
	EntryTextIndexToLayoutIndex   func(e *Entry, textIndex int) int
	EntryUnsetInvisibleChar       func(e *Entry)
)

func (e *Entry) GetBuffer() *EntryBuffer {
	return EntryGetBuffer(e)
}

func (e *Entry) SetBuffer(buffer *EntryBuffer) {
	EntrySetBuffer(e, buffer)
}

func (e *Entry) GetTextWindow() *T.GdkWindow {
	return EntryGetTextWindow(e)
}

func (e *Entry) SetVisibility(visible T.Gboolean) {
	EntrySetVisibility(e, visible)
}

func (e *Entry) GetVisibility() T.Gboolean {
	return EntryGetVisibility(e)
}

func (e *Entry) SetInvisibleChar(ch T.Gunichar) {
	EntrySetInvisibleChar(e, ch)
}

func (e *Entry) GetInvisibleChar() T.Gunichar {
	return EntryGetInvisibleChar(e)
}

func (e *Entry) UnsetInvisibleChar() { EntryUnsetInvisibleChar(e) }

func (e *Entry) SetHasFrame(setting T.Gboolean) {
	EntrySetHasFrame(e, setting)
}

func (e *Entry) GetHasFrame() T.Gboolean {
	return EntryGetHasFrame(e)
}

func (e *Entry) SetInnerBorder(border *Border) {
	EntrySetInnerBorder(e, border)
}

func (e *Entry) GetInnerBorder() *Border {
	return EntryGetInnerBorder(e)
}

func (e *Entry) SetOverwriteMode(overwrite T.Gboolean) {
	EntrySetOverwriteMode(e, overwrite)
}

func (e *Entry) GetOverwriteMode() T.Gboolean {
	return EntryGetOverwriteMode(e)
}

func (e *Entry) SetMaxLength(max int) { EntrySetMaxLength(e, max) }

func (e *Entry) GetMaxLength() int { return EntryGetMaxLength(e) }

func (e *Entry) GetTextLength() uint16 {
	return EntryGetTextLength(e)
}

func (e *Entry) SetActivatesDefault(setting T.Gboolean) {
	EntrySetActivatesDefault(e, setting)
}

func (e *Entry) GetActivatesDefault() T.Gboolean {
	return EntryGetActivatesDefault(e)
}

func (e *Entry) SetWidthChars(nChars int) {
	EntrySetWidthChars(e, nChars)
}

func (e *Entry) GetWidthChars() int { return EntryGetWidthChars(e) }

func (e *Entry) SetText(text string) { EntrySetText(e, text) }

func (e *Entry) GetText() string { return EntryGetText(e) }

func (e *Entry) GetLayout() *T.PangoLayout {
	return EntryGetLayout(e)
}

func (e *Entry) GetLayoutOffsets(x, y *int) {
	EntryGetLayoutOffsets(e, x, y)
}

func (e *Entry) SetAlignment(xalign float32) {
	EntrySetAlignment(e, xalign)
}

func (e *Entry) GetAlignment() float32 {
	return EntryGetAlignment(e)
}

func (e *Entry) SetCompletion(completion *EntryCompletion) {
	EntrySetCompletion(e, completion)
}

func (e *Entry) GetCompletion() *EntryCompletion {
	return EntryGetCompletion(e)
}

func (e *Entry) LayoutIndexToTextIndex(layoutIndex int) int {
	return EntryLayoutIndexToTextIndex(e, layoutIndex)
}

func (e *Entry) TextIndexToLayoutIndex(textIndex int) int {
	return EntryTextIndexToLayoutIndex(e, textIndex)
}

func (e *Entry) SetCursorHadjustment(adjustment *T.GtkAdjustment) {
	EntrySetCursorHadjustment(e, adjustment)
}

func (e *Entry) GetCursorHadjustment() *T.GtkAdjustment {
	return EntryGetCursorHadjustment(e)
}

func (e *Entry) SetProgressFraction(fraction float64) {
	EntrySetProgressFraction(e, fraction)
}

func (e *Entry) GetProgressFraction() float64 {
	return EntryGetProgressFraction(e)
}

func (e *Entry) SetProgressPulseStep(fraction float64) {
	EntrySetProgressPulseStep(e, fraction)
}

func (e *Entry) GetProgressPulseStep() float64 {
	return EntryGetProgressPulseStep(e)
}

func (e *Entry) ProgressPulse() { EntryProgressPulse(e) }

func (e *Entry) SetIconFromPixbuf(
	iconPos EntryIconPosition, pixbuf *T.GdkPixbuf) {
	EntrySetIconFromPixbuf(e, iconPos, pixbuf)
}

func (e *Entry) SetIconFromStock(
	iconPos EntryIconPosition, stockId string) {
	EntrySetIconFromStock(e, iconPos, stockId)
}

func (e *Entry) SetIconFromIconName(
	iconPos EntryIconPosition, iconName string) {
	EntrySetIconFromIconName(e, iconPos, iconName)
}

func (e *Entry) SetIconFromGicon(
	iconPos EntryIconPosition, icon *T.GIcon) {
	EntrySetIconFromGicon(e, iconPos, icon)
}

func (e *Entry) GetIconStorageType(
	iconPos EntryIconPosition) ImageType {
	return EntryGetIconStorageType(e, iconPos)
}

func (e *Entry) GetIconPixbuf(
	iconPos EntryIconPosition) *T.GdkPixbuf {
	return EntryGetIconPixbuf(e, iconPos)
}

func (e *Entry) GetIconStock(
	iconPos EntryIconPosition) string {
	return EntryGetIconStock(e, iconPos)
}

func (e *Entry) GetIconName(
	iconPos EntryIconPosition) string {
	return EntryGetIconName(e, iconPos)
}

func (e *Entry) GetIconGicon(
	iconPos EntryIconPosition) *T.GIcon {
	return EntryGetIconGicon(e, iconPos)
}

func (e *Entry) SetIconActivatable(
	iconPos EntryIconPosition, activatable T.Gboolean) {
	EntrySetIconActivatable(e, iconPos, activatable)
}

func (e *Entry) GetIconActivatable(
	iconPos EntryIconPosition) T.Gboolean {
	return EntryGetIconActivatable(e, iconPos)
}

func (e *Entry) SetIconSensitive(
	iconPos EntryIconPosition, sensitive T.Gboolean) {
	EntrySetIconSensitive(e, iconPos, sensitive)
}

func (e *Entry) GetIconSensitive(
	iconPos EntryIconPosition) T.Gboolean {
	return EntryGetIconSensitive(e, iconPos)
}

func (e *Entry) GetIconAtPos(x, y int) int {
	return EntryGetIconAtPos(e, x, y)
}

func (e *Entry) SetIconTooltipText(
	iconPos EntryIconPosition, tooltip string) {
	EntrySetIconTooltipText(e, iconPos, tooltip)
}

func (e *Entry) GetIconTooltipText(
	iconPos EntryIconPosition) string {
	return EntryGetIconTooltipText(e, iconPos)
}

func (e *Entry) SetIconTooltipMarkup(
	iconPos EntryIconPosition, tooltip string) {
	EntrySetIconTooltipMarkup(e, iconPos, tooltip)
}

func (e *Entry) GetIconTooltipMarkup(
	iconPos EntryIconPosition) string {
	return EntryGetIconTooltipMarkup(e, iconPos)
}

func (e *Entry) SetIconDragSource(iconPos EntryIconPosition,
	targetList *T.GtkTargetList, actions T.GdkDragAction) {
	EntrySetIconDragSource(e, iconPos, targetList, actions)
}

func (e *Entry) GetCurrentIconDragSource() int {
	return EntryGetCurrentIconDragSource(e)
}

func (e *Entry) GetIconWindow(
	iconPos EntryIconPosition) *T.GdkWindow {
	return EntryGetIconWindow(e, iconPos)
}

func (e *Entry) ImContextFilterKeypress(
	event *T.GdkEventKey) T.Gboolean {
	return EntryImContextFilterKeypress(e, event)
}

func (e *Entry) ResetImContext() { EntryResetImContext(e) }

func (e *Entry) AppendText(text string) {
	EntryAppendText(e, text)
}

func (e *Entry) PrependText(text string) {
	EntryPrependText(e, text)
}

func (e *Entry) SetPosition(position int) {
	EntrySetPosition(e, position)
}

func (e *Entry) SelectRegion(start, end int) {
	EntrySelectRegion(e, start, end)
}

func (e *Entry) SetEditable(editable T.Gboolean) {
	EntrySetEditable(e, editable)
}

var (
	EntryBufferGetType func() T.GType
	EntryBufferNew     func(
		initialChars string, nInitialChars int) *EntryBuffer

	EntryBufferGetBytes         func(e *EntryBuffer) T.Gsize
	EntryBufferGetLength        func(e *EntryBuffer) uint
	EntryBufferGetText          func(e *EntryBuffer) string
	EntryBufferSetText          func(e *EntryBuffer, chars string, nChars int)
	EntryBufferSetMaxLength     func(e *EntryBuffer, maxLength int)
	EntryBufferGetMaxLength     func(e *EntryBuffer) int
	EntryBufferInsertText       func(e *EntryBuffer, position uint, chars string, nChars int) uint
	EntryBufferDeleteText       func(e *EntryBuffer, position uint, nChars int) uint
	EntryBufferEmitInsertedText func(e *EntryBuffer, position uint, chars string, nChars uint)
	EntryBufferEmitDeletedText  func(e *EntryBuffer, position uint, nChars uint)
)

func (e *EntryBuffer) GetBytes() T.Gsize {
	return EntryBufferGetBytes(e)
}

func (e *EntryBuffer) GetLength() uint {
	return EntryBufferGetLength(e)
}

func (e *EntryBuffer) GetText() string {
	return EntryBufferGetText(e)
}

func (e *EntryBuffer) SetText(chars string, nChars int) {
	EntryBufferSetText(e, chars, nChars)
}

func (e *EntryBuffer) SetMaxLength(maxLength int) {
	EntryBufferSetMaxLength(e, maxLength)
}

func (e *EntryBuffer) GetMaxLength() int {
	return EntryBufferGetMaxLength(e)
}

func (e *EntryBuffer) InsertText(position uint, chars string, nChars int) uint {
	return EntryBufferInsertText(e, position, chars, nChars)
}

func (e *EntryBuffer) DeleteText(position uint, nChars int) uint {
	return EntryBufferDeleteText(e, position, nChars)
}

func (e *EntryBuffer) EmitInsertedText(position uint, chars string, nChars uint) {
	EntryBufferEmitInsertedText(e, position, chars, nChars)
}

func (e *EntryBuffer) EmitDeletedText(position uint, nChars uint) {
	EntryBufferEmitDeletedText(e, position, nChars)
}

type (
	GtkEntryCompletion struct {
		Parent T.GObject
		_      *struct{}
	}

	EntryCompletionMatchFunc func(
		completion *EntryCompletion,
		key string,
		iter *T.GtkTreeIter,
		userData T.Gpointer) T.Gboolean
)

var (
	EntryCompletionGetType func() T.GType
	EntryCompletionNew     func() *EntryCompletion

	EntryCompletionComplete            func(completion *EntryCompletion)
	EntryCompletionDeleteAction        func(completion *EntryCompletion, index int)
	EntryCompletionGetCompletionPrefix func(completion *EntryCompletion) string
	EntryCompletionGetEntry            func(completion *EntryCompletion) *T.GtkWidget
	EntryCompletionGetInlineCompletion func(completion *EntryCompletion) T.Gboolean
	EntryCompletionGetInlineSelection  func(completion *EntryCompletion) T.Gboolean
	EntryCompletionGetMinimumKeyLength func(completion *EntryCompletion) int
	EntryCompletionGetModel            func(completion *EntryCompletion) *T.GtkTreeModel
	EntryCompletionGetPopupCompletion  func(completion *EntryCompletion) T.Gboolean
	EntryCompletionGetPopupSetWidth    func(completion *EntryCompletion) T.Gboolean
	EntryCompletionGetPopupSingleMatch func(completion *EntryCompletion) T.Gboolean
	EntryCompletionGetTextColumn       func(completion *EntryCompletion) int
	EntryCompletionInsertActionMarkup  func(completion *EntryCompletion, index int, markup string)
	EntryCompletionInsertActionText    func(completion *EntryCompletion, index int, text string)
	EntryCompletionInsertPrefix        func(completion *EntryCompletion)
	EntryCompletionSetInlineCompletion func(completion *EntryCompletion, inlineCompletion T.Gboolean)
	EntryCompletionSetInlineSelection  func(completion *EntryCompletion, inlineSelection T.Gboolean)
	EntryCompletionSetMatchFunc        func(completion *EntryCompletion, f EntryCompletionMatchFunc, funcDataGpointer, funcNotify T.GDestroyNotify)
	EntryCompletionSetMinimumKeyLength func(completion *EntryCompletion, length int)
	EntryCompletionSetModel            func(completion *EntryCompletion, model *T.GtkTreeModel)
	EntryCompletionSetPopupCompletion  func(completion *EntryCompletion, popupCompletion T.Gboolean)
	EntryCompletionSetPopupSetWidth    func(completion *EntryCompletion, popupSetWidth T.Gboolean)
	EntryCompletionSetPopupSingleMatch func(completion *EntryCompletion, popupSingleMatch T.Gboolean)
	EntryCompletionSetTextColumn       func(completion *EntryCompletion, column int)
)

func (e *EntryCompletion) GetEntry() *T.GtkWidget {
	return EntryCompletionGetEntry(e)
}

func (e *EntryCompletion) SetModel(model *T.GtkTreeModel) {
	EntryCompletionSetModel(e, model)
}

func (e *EntryCompletion) GetModel() *T.GtkTreeModel {
	return EntryCompletionGetModel(e)
}

func (e *EntryCompletion) SetMatchFunc(f EntryCompletionMatchFunc,
	funcDataGpointer, funcNotify T.GDestroyNotify) {
	EntryCompletionSetMatchFunc(
		e, f, funcDataGpointer, funcNotify)
}

func (e *EntryCompletion) SetMinimumKeyLength(length int) {
	EntryCompletionSetMinimumKeyLength(e, length)
}

func (e *EntryCompletion) GetMinimumKeyLength() int {
	return EntryCompletionGetMinimumKeyLength(e)
}

func (e *EntryCompletion) Complete() {
	EntryCompletionComplete(e)
}

func (e *EntryCompletion) InsertPrefix() {
	EntryCompletionInsertPrefix(e)
}

func (e *EntryCompletion) InsertActionText(index int, text string) {
	EntryCompletionInsertActionText(e, index, text)
}

func (e *EntryCompletion) InsertActionMarkup(
	index int, markup string) {
	EntryCompletionInsertActionMarkup(e, index, markup)
}

func (e *EntryCompletion) DeleteAction(index int) {
	EntryCompletionDeleteAction(e, index)
}

func (e *EntryCompletion) SetInlineCompletion(
	inlineCompletion T.Gboolean) {
	EntryCompletionSetInlineCompletion(e, inlineCompletion)
}

func (e *EntryCompletion) GetInlineCompletion() T.Gboolean {
	return EntryCompletionGetInlineCompletion(e)
}

func (e *EntryCompletion) SetInlineSelection(
	inlineSelection T.Gboolean) {
	EntryCompletionSetInlineSelection(e, inlineSelection)
}

func (e *EntryCompletion) GetInlineSelection() T.Gboolean {
	return EntryCompletionGetInlineSelection(e)
}

func (e *EntryCompletion) SetPopupCompletion(
	popupCompletion T.Gboolean) {
	EntryCompletionSetPopupCompletion(e, popupCompletion)
}

func (e *EntryCompletion) GetPopupCompletion() T.Gboolean {
	return EntryCompletionGetPopupCompletion(e)
}

func (e *EntryCompletion) SetPopupSetWidth(
	popupSetWidth T.Gboolean) {
	EntryCompletionSetPopupSetWidth(e, popupSetWidth)
}

func (e *EntryCompletion) GetPopupSetWidth() T.Gboolean {
	return EntryCompletionGetPopupSetWidth(e)
}

func (e *EntryCompletion) SetPopupSingleMatch(
	popupSingleMatch T.Gboolean) {
	EntryCompletionSetPopupSingleMatch(e, popupSingleMatch)
}

func (e *EntryCompletion) GetPopupSingleMatch() T.Gboolean {
	return EntryCompletionGetPopupSingleMatch(e)
}

func (e *EntryCompletion) GetCompletionPrefix() string {
	return EntryCompletionGetCompletionPrefix(e)
}

func (e *EntryCompletion) SetTextColumn(column int) {
	EntryCompletionSetTextColumn(e, column)
}

func (e *EntryCompletion) GetTextColumn() int {
	return EntryCompletionGetTextColumn(e)
}

type EventBox struct {
	Bin T.GtkBin
}

var (
	EventBoxGetType func() T.GType
	EventBoxNew     func() *T.GtkWidget

	EventBoxGetAboveChild    func(e *EventBox) T.Gboolean
	EventBoxGetVisibleWindow func(e *EventBox) T.Gboolean
	EventBoxSetAboveChild    func(e *EventBox, aboveChild T.Gboolean)
	EventBoxSetVisibleWindow func(e *EventBox, visibleWindow T.Gboolean)
)

func (e *EventBox) GetVisibleWindow() T.Gboolean {
	return EventBoxGetVisibleWindow(e)
}

func (e *EventBox) SetVisibleWindow(visibleWindow T.Gboolean) {
	EventBoxSetVisibleWindow(e, visibleWindow)
}

func (e *EventBox) GetAboveChild() T.Gboolean {
	return EventBoxGetAboveChild(e)
}

func (e *EventBox) SetAboveChild(aboveChild T.Gboolean) {
	EventBoxSetAboveChild(e, aboveChild)
}

type Expander struct {
	Bin T.GtkBin
	_   *struct{}
}

var (
	ExpanderGetType         func() T.GType
	ExpanderNew             func(label string) *T.GtkWidget
	ExpanderNewWithMnemonic func(label string) *T.GtkWidget

	ExpanderStyleGetType func() T.GType

	ExpanderSetExpanded     func(e *Expander, expanded T.Gboolean)
	ExpanderGetExpanded     func(e *Expander) T.Gboolean
	ExpanderSetSpacing      func(e *Expander, spacing int)
	ExpanderGetSpacing      func(e *Expander) int
	ExpanderSetLabel        func(e *Expander, label string)
	ExpanderGetLabel        func(e *Expander) string
	ExpanderSetUseUnderline func(e *Expander, useUnderline T.Gboolean)
	ExpanderGetUseUnderline func(e *Expander) T.Gboolean
	ExpanderSetUseMarkup    func(e *Expander, useMarkup T.Gboolean)
	ExpanderGetUseMarkup    func(e *Expander) T.Gboolean
	ExpanderSetLabelWidget  func(e *Expander, labelWidget *T.GtkWidget)
	ExpanderGetLabelWidget  func(e *Expander) *T.GtkWidget
	ExpanderSetLabelFill    func(e *Expander, labelFill T.Gboolean)
	ExpanderGetLabelFill    func(e *Expander) T.Gboolean
)

func (e *Expander) SetExpanded(expanded T.Gboolean) {
	ExpanderSetExpanded(e, expanded)
}

func (e *Expander) GetExpanded() T.Gboolean {
	return ExpanderGetExpanded(e)
}

func (e *Expander) SetSpacing(spacing int) {
	ExpanderSetSpacing(e, spacing)
}

func (e *Expander) GetSpacing() int {
	return ExpanderGetSpacing(e)
}

func (e *Expander) SetLabel(label string) {
	ExpanderSetLabel(e, label)
}

func (e *Expander) GetLabel() string {
	return ExpanderGetLabel(e)
}

func (e *Expander) SetUseUnderline(useUnderline T.Gboolean) {
	ExpanderSetUseUnderline(e, useUnderline)
}

func (e *Expander) GetUseUnderline() T.Gboolean {
	return ExpanderGetUseUnderline(e)
}

func (e *Expander) SetUseMarkup(useMarkup T.Gboolean) {
	ExpanderSetUseMarkup(e, useMarkup)
}

func (e *Expander) GetUseMarkup() T.Gboolean {
	return ExpanderGetUseMarkup(e)
}

func (e *Expander) SetLabelWidget(labelWidget *T.GtkWidget) {
	ExpanderSetLabelWidget(e, labelWidget)
}

func (e *Expander) GetLabelWidget() *T.GtkWidget {
	return ExpanderGetLabelWidget(e)
}

func (e *Expander) SetLabelFill(labelFill T.Gboolean) {
	ExpanderSetLabelFill(e, labelFill)
}

func (e *Expander) GetLabelFill() T.Gboolean {
	return ExpanderGetLabelFill(e)
}
