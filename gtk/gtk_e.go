// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	I "github.com/tHinqa/outside-gtk2/gio"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Editable struct{}

var (
	EditableGetType func() O.Type

	editableCopyClipboard      func(e *Editable)
	editableCutClipboard       func(e *Editable)
	editableDeleteSelection    func(e *Editable)
	editableDeleteText         func(e *Editable, startPos, endPos int)
	editableGetChars           func(e *Editable, startPos, endPos int) string
	editableGetEditable        func(e *Editable) bool
	editableGetPosition        func(e *Editable) int
	editableGetSelectionBounds func(e *Editable, startPos, endPos *int) bool
	editableInsertText         func(e *Editable, newText string, newTextLength int, position *int)
	editablePasteClipboard     func(e *Editable)
	editableSelectRegion       func(e *Editable, startPos, endPos int)
	editableSetEditable        func(e *Editable, isEditable bool)
	editableSetPosition        func(e *Editable, position int)
)

func (e *Editable) CopyClipboard()                       { editableCopyClipboard(e) }
func (e *Editable) CutClipboard()                        { editableCutClipboard(e) }
func (e *Editable) DeleteSelection()                     { editableDeleteSelection(e) }
func (e *Editable) DeleteText(startPos, endPos int)      { editableDeleteText(e, startPos, endPos) }
func (e *Editable) GetChars(startPos, endPos int) string { return editableGetChars(e, startPos, endPos) }
func (e *Editable) GetEditable() bool                    { return editableGetEditable(e) }
func (e *Editable) GetPosition() int                     { return editableGetPosition(e) }
func (e *Editable) GetSelectionBounds(startPos, endPos *int) bool {
	return editableGetSelectionBounds(e, startPos, endPos)
}
func (e *Editable) InsertText(newText string, newTextLength int, position *int) {
	editableInsertText(e, newText, newTextLength, position)
}
func (e *Editable) PasteClipboard()                   { editablePasteClipboard(e) }
func (e *Editable) SelectRegion(startPos, endPos int) { editableSelectRegion(e, startPos, endPos) }
func (e *Editable) SetEditable(isEditable bool)       { editableSetEditable(e, isEditable) }
func (e *Editable) SetPosition(position int)          { editableSetPosition(e, position) }

type Entry struct {
	Widget Widget
	Text   *T.Gchar
	Bits   uint
	// Editable : 1
	// Visible : 1
	// OverwriteMode : 1
	// InDrag : 1
	TextLength     uint16
	TextMaxLength  uint16
	TextArea       *D.Window
	ImContext      *IMContext
	PopupMenu      *Widget
	CurrentPos     int
	SelectionBound int
	CachedLayout   *P.Layout
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

type EntryIconPosition Enum

const (
	ENTRY_ICON_PRIMARY EntryIconPosition = iota
	ENTRY_ICON_SECONDARY
)

var (
	EntryGetType          func() O.Type
	EntryNew              func() *Widget
	EntryNewWithMaxLength func(max int) *Widget
	EntryNewWithBuffer    func(buffer *EntryBuffer) *Widget

	EntryIconPositionGetType func() O.Type

	entryAppendText               func(e *Entry, text string)
	entryGetActivatesDefault      func(e *Entry) bool
	entryGetAlignment             func(e *Entry) float32
	entryGetBuffer                func(e *Entry) *EntryBuffer
	entryGetCompletion            func(e *Entry) *EntryCompletion
	entryGetCurrentIconDragSource func(e *Entry) int
	entryGetCursorHadjustment     func(e *Entry) *Adjustment
	entryGetHasFrame              func(e *Entry) bool
	entryGetIconActivatable       func(e *Entry, iconPos EntryIconPosition) bool
	entryGetIconAtPos             func(e *Entry, x, y int) int
	entryGetIconGicon             func(e *Entry, iconPos EntryIconPosition) *I.Icon
	entryGetIconName              func(e *Entry, iconPos EntryIconPosition) string
	entryGetIconPixbuf            func(e *Entry, iconPos EntryIconPosition) *D.Pixbuf
	entryGetIconSensitive         func(e *Entry, iconPos EntryIconPosition) bool
	entryGetIconStock             func(e *Entry, iconPos EntryIconPosition) string
	entryGetIconStorageType       func(e *Entry, iconPos EntryIconPosition) ImageType
	entryGetIconTooltipMarkup     func(e *Entry, iconPos EntryIconPosition) string
	entryGetIconTooltipText       func(e *Entry, iconPos EntryIconPosition) string
	entryGetIconWindow            func(e *Entry, iconPos EntryIconPosition) *D.Window
	entryGetInnerBorder           func(e *Entry) *Border
	entryGetInvisibleChar         func(e *Entry) T.Gunichar
	entryGetLayout                func(e *Entry) *P.Layout
	entryGetLayoutOffsets         func(e *Entry, x, y *int)
	entryGetMaxLength             func(e *Entry) int
	entryGetOverwriteMode         func(e *Entry) bool
	entryGetProgressFraction      func(e *Entry) float64
	entryGetProgressPulseStep     func(e *Entry) float64
	entryGetText                  func(e *Entry) string
	entryGetTextLength            func(e *Entry) uint16
	entryGetTextWindow            func(e *Entry) *D.Window
	entryGetVisibility            func(e *Entry) bool
	entryGetWidthChars            func(e *Entry) int
	entryImContextFilterKeypress  func(e *Entry, event *D.EventKey) bool
	entryLayoutIndexToTextIndex   func(e *Entry, layoutIndex int) int
	entryPrependText              func(e *Entry, text string)
	entryProgressPulse            func(e *Entry)
	entryResetImContext           func(e *Entry)
	entrySelectRegion             func(e *Entry, start, end int)
	entrySetActivatesDefault      func(e *Entry, setting bool)
	entrySetAlignment             func(e *Entry, xalign float32)
	entrySetBuffer                func(e *Entry, buffer *EntryBuffer)
	entrySetCompletion            func(e *Entry, completion *EntryCompletion)
	entrySetCursorHadjustment     func(e *Entry, adjustment *Adjustment)
	entrySetEditable              func(e *Entry, editable bool)
	entrySetHasFrame              func(e *Entry, setting bool)
	entrySetIconActivatable       func(e *Entry, iconPos EntryIconPosition, activatable bool)
	entrySetIconDragSource        func(e *Entry, iconPos EntryIconPosition, targetList *TargetList, actions D.DragAction)
	entrySetIconFromGicon         func(e *Entry, iconPos EntryIconPosition, icon *I.Icon)
	entrySetIconFromIconName      func(e *Entry, iconPos EntryIconPosition, iconName string)
	entrySetIconFromPixbuf        func(e *Entry, iconPos EntryIconPosition, pixbuf *D.Pixbuf)
	entrySetIconFromStock         func(e *Entry, iconPos EntryIconPosition, stockId string)
	entrySetIconSensitive         func(e *Entry, iconPos EntryIconPosition, sensitive bool)
	entrySetIconTooltipMarkup     func(e *Entry, iconPos EntryIconPosition, tooltip string)
	entrySetIconTooltipText       func(e *Entry, iconPos EntryIconPosition, tooltip string)
	entrySetInnerBorder           func(e *Entry, border *Border)
	entrySetInvisibleChar         func(e *Entry, ch T.Gunichar)
	entrySetMaxLength             func(e *Entry, max int)
	entrySetOverwriteMode         func(e *Entry, overwrite bool)
	entrySetPosition              func(e *Entry, position int)
	entrySetProgressFraction      func(e *Entry, fraction float64)
	entrySetProgressPulseStep     func(e *Entry, fraction float64)
	entrySetText                  func(e *Entry, text string)
	entrySetVisibility            func(e *Entry, visible bool)
	entrySetWidthChars            func(e *Entry, nChars int)
	entryTextIndexToLayoutIndex   func(e *Entry, textIndex int) int
	entryUnsetInvisibleChar       func(e *Entry)
)

func (e *Entry) AppendText(text string)            { entryAppendText(e, text) }
func (e *Entry) GetActivatesDefault() bool         { return entryGetActivatesDefault(e) }
func (e *Entry) GetAlignment() float32             { return entryGetAlignment(e) }
func (e *Entry) GetBuffer() *EntryBuffer           { return entryGetBuffer(e) }
func (e *Entry) GetCompletion() *EntryCompletion   { return entryGetCompletion(e) }
func (e *Entry) GetCurrentIconDragSource() int     { return entryGetCurrentIconDragSource(e) }
func (e *Entry) GetCursorHadjustment() *Adjustment { return entryGetCursorHadjustment(e) }
func (e *Entry) GetHasFrame() bool                 { return entryGetHasFrame(e) }
func (e *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	return entryGetIconActivatable(e, iconPos)
}
func (e *Entry) GetIconAtPos(x, y int) int                      { return entryGetIconAtPos(e, x, y) }
func (e *Entry) GetIconGicon(iconPos EntryIconPosition) *I.Icon { return entryGetIconGicon(e, iconPos) }
func (e *Entry) GetIconName(iconPos EntryIconPosition) string   { return entryGetIconName(e, iconPos) }
func (e *Entry) GetIconPixbuf(iconPos EntryIconPosition) *D.Pixbuf {
	return entryGetIconPixbuf(e, iconPos)
}
func (e *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	return entryGetIconSensitive(e, iconPos)
}
func (e *Entry) GetIconStock(iconPos EntryIconPosition) string { return entryGetIconStock(e, iconPos) }
func (e *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	return entryGetIconStorageType(e, iconPos)
}
func (e *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) string {
	return entryGetIconTooltipMarkup(e, iconPos)
}
func (e *Entry) GetIconTooltipText(iconPos EntryIconPosition) string {
	return entryGetIconTooltipText(e, iconPos)
}
func (e *Entry) GetIconWindow(iconPos EntryIconPosition) *D.Window {
	return entryGetIconWindow(e, iconPos)
}
func (e *Entry) GetInnerBorder() *Border       { return entryGetInnerBorder(e) }
func (e *Entry) GetInvisibleChar() T.Gunichar  { return entryGetInvisibleChar(e) }
func (e *Entry) GetLayout() *P.Layout          { return entryGetLayout(e) }
func (e *Entry) GetLayoutOffsets(x, y *int)    { entryGetLayoutOffsets(e, x, y) }
func (e *Entry) GetMaxLength() int             { return entryGetMaxLength(e) }
func (e *Entry) GetOverwriteMode() bool        { return entryGetOverwriteMode(e) }
func (e *Entry) GetProgressFraction() float64  { return entryGetProgressFraction(e) }
func (e *Entry) GetProgressPulseStep() float64 { return entryGetProgressPulseStep(e) }
func (e *Entry) GetText() string               { return entryGetText(e) }
func (e *Entry) GetTextLength() uint16         { return entryGetTextLength(e) }
func (e *Entry) GetTextWindow() *D.Window      { return entryGetTextWindow(e) }
func (e *Entry) GetVisibility() bool           { return entryGetVisibility(e) }
func (e *Entry) GetWidthChars() int            { return entryGetWidthChars(e) }
func (e *Entry) ImContextFilterKeypress(event *D.EventKey) bool {
	return entryImContextFilterKeypress(e, event)
}
func (e *Entry) LayoutIndexToTextIndex(layoutIndex int) int {
	return entryLayoutIndexToTextIndex(e, layoutIndex)
}
func (e *Entry) PrependText(text string)                     { entryPrependText(e, text) }
func (e *Entry) ProgressPulse()                              { entryProgressPulse(e) }
func (e *Entry) ResetImContext()                             { entryResetImContext(e) }
func (e *Entry) SelectRegion(start, end int)                 { entrySelectRegion(e, start, end) }
func (e *Entry) SetActivatesDefault(setting bool)            { entrySetActivatesDefault(e, setting) }
func (e *Entry) SetAlignment(xalign float32)                 { entrySetAlignment(e, xalign) }
func (e *Entry) SetBuffer(buffer *EntryBuffer)               { entrySetBuffer(e, buffer) }
func (e *Entry) SetCompletion(completion *EntryCompletion)   { entrySetCompletion(e, completion) }
func (e *Entry) SetCursorHadjustment(adjustment *Adjustment) { entrySetCursorHadjustment(e, adjustment) }
func (e *Entry) SetEditable(editable bool)                   { entrySetEditable(e, editable) }
func (e *Entry) SetHasFrame(setting bool)                    { entrySetHasFrame(e, setting) }
func (e *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	entrySetIconActivatable(e, iconPos, activatable)
}
func (e *Entry) SetIconDragSource(iconPos EntryIconPosition, targetList *TargetList, actions D.DragAction) {
	entrySetIconDragSource(e, iconPos, targetList, actions)
}
func (e *Entry) SetIconFromGicon(iconPos EntryIconPosition, icon *I.Icon) {
	entrySetIconFromGicon(e, iconPos, icon)
}
func (e *Entry) SetIconFromIconName(iconPos EntryIconPosition, iconName string) {
	entrySetIconFromIconName(e, iconPos, iconName)
}
func (e *Entry) SetIconFromPixbuf(iconPos EntryIconPosition, pixbuf *D.Pixbuf) {
	entrySetIconFromPixbuf(e, iconPos, pixbuf)
}
func (e *Entry) SetIconFromStock(iconPos EntryIconPosition, stockId string) {
	entrySetIconFromStock(e, iconPos, stockId)
}
func (e *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	entrySetIconSensitive(e, iconPos, sensitive)
}
func (e *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	entrySetIconTooltipMarkup(e, iconPos, tooltip)
}
func (e *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	entrySetIconTooltipText(e, iconPos, tooltip)
}
func (e *Entry) SetInnerBorder(border *Border)         { entrySetInnerBorder(e, border) }
func (e *Entry) SetInvisibleChar(ch T.Gunichar)        { entrySetInvisibleChar(e, ch) }
func (e *Entry) SetMaxLength(max int)                  { entrySetMaxLength(e, max) }
func (e *Entry) SetOverwriteMode(overwrite bool)       { entrySetOverwriteMode(e, overwrite) }
func (e *Entry) SetPosition(position int)              { entrySetPosition(e, position) }
func (e *Entry) SetProgressFraction(fraction float64)  { entrySetProgressFraction(e, fraction) }
func (e *Entry) SetProgressPulseStep(fraction float64) { entrySetProgressPulseStep(e, fraction) }
func (e *Entry) SetText(text string)                   { entrySetText(e, text) }
func (e *Entry) SetVisibility(visible bool)            { entrySetVisibility(e, visible) }
func (e *Entry) SetWidthChars(nChars int)              { entrySetWidthChars(e, nChars) }
func (e *Entry) TextIndexToLayoutIndex(textIndex int) int {
	return entryTextIndexToLayoutIndex(e, textIndex)
}
func (e *Entry) UnsetInvisibleChar() { entryUnsetInvisibleChar(e) }

var (
	EntryBufferGetType func() O.Type
	EntryBufferNew     func(
		initialChars string, nInitialChars int) *EntryBuffer

	entryBufferGetBytes         func(e *EntryBuffer) T.Gsize
	entryBufferGetLength        func(e *EntryBuffer) uint
	entryBufferGetText          func(e *EntryBuffer) string
	entryBufferSetText          func(e *EntryBuffer, chars string, nChars int)
	entryBufferSetMaxLength     func(e *EntryBuffer, maxLength int)
	entryBufferGetMaxLength     func(e *EntryBuffer) int
	entryBufferInsertText       func(e *EntryBuffer, position uint, chars string, nChars int) uint
	entryBufferDeleteText       func(e *EntryBuffer, position uint, nChars int) uint
	entryBufferEmitInsertedText func(e *EntryBuffer, position uint, chars string, nChars uint)
	entryBufferEmitDeletedText  func(e *EntryBuffer, position uint, nChars uint)
)

func (e *EntryBuffer) DeleteText(position uint, nChars int) uint {
	return entryBufferDeleteText(e, position, nChars)
}
func (e *EntryBuffer) EmitDeletedText(position uint, nChars uint) {
	entryBufferEmitDeletedText(e, position, nChars)
}
func (e *EntryBuffer) EmitInsertedText(position uint, chars string, nChars uint) {
	entryBufferEmitInsertedText(e, position, chars, nChars)
}
func (e *EntryBuffer) GetBytes() T.Gsize { return entryBufferGetBytes(e) }
func (e *EntryBuffer) GetLength() uint   { return entryBufferGetLength(e) }
func (e *EntryBuffer) GetMaxLength() int { return entryBufferGetMaxLength(e) }
func (e *EntryBuffer) GetText() string   { return entryBufferGetText(e) }
func (e *EntryBuffer) InsertText(position uint, chars string, nChars int) uint {
	return entryBufferInsertText(e, position, chars, nChars)
}
func (e *EntryBuffer) SetMaxLength(maxLength int)       { entryBufferSetMaxLength(e, maxLength) }
func (e *EntryBuffer) SetText(chars string, nChars int) { entryBufferSetText(e, chars, nChars) }

type (
	GtkEntryCompletion struct {
		Parent T.GObject
		_      *struct{}
	}

	EntryCompletionMatchFunc func(
		completion *EntryCompletion,
		key string,
		iter *TreeIter,
		userData T.Gpointer) bool
)

var (
	EntryCompletionGetType func() O.Type
	EntryCompletionNew     func() *EntryCompletion

	entryCompletionComplete            func(completion *EntryCompletion)
	entryCompletionDeleteAction        func(completion *EntryCompletion, index int)
	entryCompletionGetCompletionPrefix func(completion *EntryCompletion) string
	entryCompletionGetEntry            func(completion *EntryCompletion) *Widget
	entryCompletionGetInlineCompletion func(completion *EntryCompletion) bool
	entryCompletionGetInlineSelection  func(completion *EntryCompletion) bool
	entryCompletionGetMinimumKeyLength func(completion *EntryCompletion) int
	entryCompletionGetModel            func(completion *EntryCompletion) *TreeModel
	entryCompletionGetPopupCompletion  func(completion *EntryCompletion) bool
	entryCompletionGetPopupSetWidth    func(completion *EntryCompletion) bool
	entryCompletionGetPopupSingleMatch func(completion *EntryCompletion) bool
	entryCompletionGetTextColumn       func(completion *EntryCompletion) int
	entryCompletionInsertActionMarkup  func(completion *EntryCompletion, index int, markup string)
	entryCompletionInsertActionText    func(completion *EntryCompletion, index int, text string)
	entryCompletionInsertPrefix        func(completion *EntryCompletion)
	entryCompletionSetInlineCompletion func(completion *EntryCompletion, inlineCompletion bool)
	entryCompletionSetInlineSelection  func(completion *EntryCompletion, inlineSelection bool)
	entryCompletionSetMatchFunc        func(completion *EntryCompletion, f EntryCompletionMatchFunc, funcData T.Gpointer, funcNotify T.GDestroyNotify)
	entryCompletionSetMinimumKeyLength func(completion *EntryCompletion, length int)
	entryCompletionSetModel            func(completion *EntryCompletion, model *TreeModel)
	entryCompletionSetPopupCompletion  func(completion *EntryCompletion, popupCompletion bool)
	entryCompletionSetPopupSetWidth    func(completion *EntryCompletion, popupSetWidth bool)
	entryCompletionSetPopupSingleMatch func(completion *EntryCompletion, popupSingleMatch bool)
	entryCompletionSetTextColumn       func(completion *EntryCompletion, column int)
)

func (e *EntryCompletion) Complete()                   { entryCompletionComplete(e) }
func (e *EntryCompletion) DeleteAction(index int)      { entryCompletionDeleteAction(e, index) }
func (e *EntryCompletion) GetCompletionPrefix() string { return entryCompletionGetCompletionPrefix(e) }
func (e *EntryCompletion) GetEntry() *Widget           { return entryCompletionGetEntry(e) }
func (e *EntryCompletion) GetInlineCompletion() bool {
	return entryCompletionGetInlineCompletion(e)
}
func (e *EntryCompletion) GetInlineSelection() bool { return entryCompletionGetInlineSelection(e) }
func (e *EntryCompletion) GetMinimumKeyLength() int { return entryCompletionGetMinimumKeyLength(e) }
func (e *EntryCompletion) GetModel() *TreeModel     { return entryCompletionGetModel(e) }
func (e *EntryCompletion) GetPopupCompletion() bool { return entryCompletionGetPopupCompletion(e) }
func (e *EntryCompletion) GetPopupSetWidth() bool   { return entryCompletionGetPopupSetWidth(e) }
func (e *EntryCompletion) GetPopupSingleMatch() bool {
	return entryCompletionGetPopupSingleMatch(e)
}
func (e *EntryCompletion) GetTextColumn() int { return entryCompletionGetTextColumn(e) }
func (e *EntryCompletion) InsertActionMarkup(index int, markup string) {
	entryCompletionInsertActionMarkup(e, index, markup)
}
func (e *EntryCompletion) InsertActionText(index int, text string) {
	entryCompletionInsertActionText(e, index, text)
}
func (e *EntryCompletion) InsertPrefix() { entryCompletionInsertPrefix(e) }
func (e *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	entryCompletionSetInlineCompletion(e, inlineCompletion)
}
func (e *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	entryCompletionSetInlineSelection(e, inlineSelection)
}
func (e *EntryCompletion) SetMatchFunc(f EntryCompletionMatchFunc, funcData T.Gpointer, funcNotify T.GDestroyNotify) {
	entryCompletionSetMatchFunc(e, f, funcData, funcNotify)
}
func (e *EntryCompletion) SetMinimumKeyLength(length int) {
	entryCompletionSetMinimumKeyLength(e, length)
}
func (e *EntryCompletion) SetModel(model *TreeModel) { entryCompletionSetModel(e, model) }
func (e *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	entryCompletionSetPopupCompletion(e, popupCompletion)
}
func (e *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	entryCompletionSetPopupSetWidth(e, popupSetWidth)
}
func (e *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	entryCompletionSetPopupSingleMatch(e, popupSingleMatch)
}
func (e *EntryCompletion) SetTextColumn(column int) { entryCompletionSetTextColumn(e, column) }

type EnumValue O.EnumValue

type EventBox struct{ Bin Bin }

var (
	EventBoxGetType func() O.Type
	EventBoxNew     func() *Widget

	eventBoxGetAboveChild    func(e *EventBox) bool
	eventBoxGetVisibleWindow func(e *EventBox) bool
	eventBoxSetAboveChild    func(e *EventBox, aboveChild bool)
	eventBoxSetVisibleWindow func(e *EventBox, visibleWindow bool)
)

func (e *EventBox) GetAboveChild() bool           { return eventBoxGetAboveChild(e) }
func (e *EventBox) GetVisibleWindow() bool        { return eventBoxGetVisibleWindow(e) }
func (e *EventBox) SetAboveChild(aboveChild bool) { eventBoxSetAboveChild(e, aboveChild) }
func (e *EventBox) SetVisibleWindow(visibleWindow bool) {
	eventBoxSetVisibleWindow(e, visibleWindow)
}

type Expander struct {
	Bin Bin
	_   *struct{}
}

type ExpanderStyle Enum

const (
	EXPANDER_COLLAPSED ExpanderStyle = iota
	EXPANDER_SEMI_COLLAPSED
	EXPANDER_SEMI_EXPANDED
	EXPANDER_EXPANDED
)

var (
	ExpanderGetType         func() O.Type
	ExpanderNew             func(label string) *Widget
	ExpanderNewWithMnemonic func(label string) *Widget

	ExpanderStyleGetType func() O.Type

	expanderSetExpanded     func(e *Expander, expanded bool)
	expanderGetExpanded     func(e *Expander) bool
	expanderSetSpacing      func(e *Expander, spacing int)
	expanderGetSpacing      func(e *Expander) int
	expanderSetLabel        func(e *Expander, label string)
	expanderGetLabel        func(e *Expander) string
	expanderSetUseUnderline func(e *Expander, useUnderline bool)
	expanderGetUseUnderline func(e *Expander) bool
	expanderSetUseMarkup    func(e *Expander, useMarkup bool)
	expanderGetUseMarkup    func(e *Expander) bool
	expanderSetLabelWidget  func(e *Expander, labelWidget *Widget)
	expanderGetLabelWidget  func(e *Expander) *Widget
	expanderSetLabelFill    func(e *Expander, labelFill bool)
	expanderGetLabelFill    func(e *Expander) bool
)

func (e *Expander) GetExpanded() bool                  { return expanderGetExpanded(e) }
func (e *Expander) GetLabel() string                   { return expanderGetLabel(e) }
func (e *Expander) GetLabelFill() bool                 { return expanderGetLabelFill(e) }
func (e *Expander) GetLabelWidget() *Widget            { return expanderGetLabelWidget(e) }
func (e *Expander) GetSpacing() int                    { return expanderGetSpacing(e) }
func (e *Expander) GetUseMarkup() bool                 { return expanderGetUseMarkup(e) }
func (e *Expander) GetUseUnderline() bool              { return expanderGetUseUnderline(e) }
func (e *Expander) SetExpanded(expanded bool)          { expanderSetExpanded(e, expanded) }
func (e *Expander) SetLabel(label string)              { expanderSetLabel(e, label) }
func (e *Expander) SetLabelFill(labelFill bool)        { expanderSetLabelFill(e, labelFill) }
func (e *Expander) SetLabelWidget(labelWidget *Widget) { expanderSetLabelWidget(e, labelWidget) }
func (e *Expander) SetSpacing(spacing int)             { expanderSetSpacing(e, spacing) }
func (e *Expander) SetUseMarkup(useMarkup bool)        { expanderSetUseMarkup(e, useMarkup) }
func (e *Expander) SetUseUnderline(useUnderline bool)  { expanderSetUseUnderline(e, useUnderline) }
