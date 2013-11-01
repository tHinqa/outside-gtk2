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

	EditableCopyClipboard      func(e *Editable)
	EditableCutClipboard       func(e *Editable)
	EditableDeleteSelection    func(e *Editable)
	EditableDeleteText         func(e *Editable, startPos, endPos int)
	EditableGetChars           func(e *Editable, startPos, endPos int) string
	EditableGetEditable        func(e *Editable) bool
	EditableGetPosition        func(e *Editable) int
	EditableGetSelectionBounds func(e *Editable, startPos, endPos *int) bool
	EditableInsertText         func(e *Editable, newText string, newTextLength int, position *int)
	EditablePasteClipboard     func(e *Editable)
	EditableSelectRegion       func(e *Editable, startPos, endPos int)
	EditableSetEditable        func(e *Editable, isEditable bool)
	EditableSetPosition        func(e *Editable, position int)
)

func (e *Editable) CopyClipboard()                       { EditableCopyClipboard(e) }
func (e *Editable) CutClipboard()                        { EditableCutClipboard(e) }
func (e *Editable) DeleteSelection()                     { EditableDeleteSelection(e) }
func (e *Editable) DeleteText(startPos, endPos int)      { EditableDeleteText(e, startPos, endPos) }
func (e *Editable) GetChars(startPos, endPos int) string { return EditableGetChars(e, startPos, endPos) }
func (e *Editable) GetEditable() bool                    { return EditableGetEditable(e) }
func (e *Editable) GetPosition() int                     { return EditableGetPosition(e) }
func (e *Editable) GetSelectionBounds(startPos, endPos *int) bool {
	return EditableGetSelectionBounds(e, startPos, endPos)
}
func (e *Editable) InsertText(newText string, newTextLength int, position *int) {
	EditableInsertText(e, newText, newTextLength, position)
}
func (e *Editable) PasteClipboard()                   { EditablePasteClipboard(e) }
func (e *Editable) SelectRegion(startPos, endPos int) { EditableSelectRegion(e, startPos, endPos) }
func (e *Editable) SetEditable(isEditable bool)       { EditableSetEditable(e, isEditable) }
func (e *Editable) SetPosition(position int)          { EditableSetPosition(e, position) }

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
	Parent O.Object
	_      *struct{}
}

type EntryCompletion struct {
	Parent O.Object
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

	EntryAppendText               func(e *Entry, text string)
	EntryGetActivatesDefault      func(e *Entry) bool
	EntryGetAlignment             func(e *Entry) float32
	EntryGetBuffer                func(e *Entry) *EntryBuffer
	EntryGetCompletion            func(e *Entry) *EntryCompletion
	EntryGetCurrentIconDragSource func(e *Entry) int
	EntryGetCursorHadjustment     func(e *Entry) *Adjustment
	EntryGetHasFrame              func(e *Entry) bool
	EntryGetIconActivatable       func(e *Entry, iconPos EntryIconPosition) bool
	EntryGetIconAtPos             func(e *Entry, x, y int) int
	EntryGetIconGicon             func(e *Entry, iconPos EntryIconPosition) *I.Icon
	EntryGetIconName              func(e *Entry, iconPos EntryIconPosition) string
	EntryGetIconPixbuf            func(e *Entry, iconPos EntryIconPosition) *D.Pixbuf
	EntryGetIconSensitive         func(e *Entry, iconPos EntryIconPosition) bool
	EntryGetIconStock             func(e *Entry, iconPos EntryIconPosition) string
	EntryGetIconStorageType       func(e *Entry, iconPos EntryIconPosition) ImageType
	EntryGetIconTooltipMarkup     func(e *Entry, iconPos EntryIconPosition) string
	EntryGetIconTooltipText       func(e *Entry, iconPos EntryIconPosition) string
	EntryGetIconWindow            func(e *Entry, iconPos EntryIconPosition) *D.Window
	EntryGetInnerBorder           func(e *Entry) *Border
	EntryGetInvisibleChar         func(e *Entry) T.Gunichar
	EntryGetLayout                func(e *Entry) *P.Layout
	EntryGetLayoutOffsets         func(e *Entry, x, y *int)
	EntryGetMaxLength             func(e *Entry) int
	EntryGetOverwriteMode         func(e *Entry) bool
	EntryGetProgressFraction      func(e *Entry) float64
	EntryGetProgressPulseStep     func(e *Entry) float64
	EntryGetText                  func(e *Entry) string
	EntryGetTextLength            func(e *Entry) uint16
	EntryGetTextWindow            func(e *Entry) *D.Window
	EntryGetVisibility            func(e *Entry) bool
	EntryGetWidthChars            func(e *Entry) int
	EntryImContextFilterKeypress  func(e *Entry, event *D.EventKey) bool
	EntryLayoutIndexToTextIndex   func(e *Entry, layoutIndex int) int
	EntryPrependText              func(e *Entry, text string)
	EntryProgressPulse            func(e *Entry)
	EntryResetImContext           func(e *Entry)
	EntrySelectRegion             func(e *Entry, start, end int)
	EntrySetActivatesDefault      func(e *Entry, setting bool)
	EntrySetAlignment             func(e *Entry, xalign float32)
	EntrySetBuffer                func(e *Entry, buffer *EntryBuffer)
	EntrySetCompletion            func(e *Entry, completion *EntryCompletion)
	EntrySetCursorHadjustment     func(e *Entry, adjustment *Adjustment)
	EntrySetEditable              func(e *Entry, editable bool)
	EntrySetHasFrame              func(e *Entry, setting bool)
	EntrySetIconActivatable       func(e *Entry, iconPos EntryIconPosition, activatable bool)
	EntrySetIconDragSource        func(e *Entry, iconPos EntryIconPosition, targetList *TargetList, actions D.DragAction)
	EntrySetIconFromGicon         func(e *Entry, iconPos EntryIconPosition, icon *I.Icon)
	EntrySetIconFromIconName      func(e *Entry, iconPos EntryIconPosition, iconName string)
	EntrySetIconFromPixbuf        func(e *Entry, iconPos EntryIconPosition, pixbuf *D.Pixbuf)
	EntrySetIconFromStock         func(e *Entry, iconPos EntryIconPosition, stockId string)
	EntrySetIconSensitive         func(e *Entry, iconPos EntryIconPosition, sensitive bool)
	EntrySetIconTooltipMarkup     func(e *Entry, iconPos EntryIconPosition, tooltip string)
	EntrySetIconTooltipText       func(e *Entry, iconPos EntryIconPosition, tooltip string)
	EntrySetInnerBorder           func(e *Entry, border *Border)
	EntrySetInvisibleChar         func(e *Entry, ch T.Gunichar)
	EntrySetMaxLength             func(e *Entry, max int)
	EntrySetOverwriteMode         func(e *Entry, overwrite bool)
	EntrySetPosition              func(e *Entry, position int)
	EntrySetProgressFraction      func(e *Entry, fraction float64)
	EntrySetProgressPulseStep     func(e *Entry, fraction float64)
	EntrySetText                  func(e *Entry, text string)
	EntrySetVisibility            func(e *Entry, visible bool)
	EntrySetWidthChars            func(e *Entry, nChars int)
	EntryTextIndexToLayoutIndex   func(e *Entry, textIndex int) int
	EntryUnsetInvisibleChar       func(e *Entry)
)

func (e *Entry) AppendText(text string)            { EntryAppendText(e, text) }
func (e *Entry) GetActivatesDefault() bool         { return EntryGetActivatesDefault(e) }
func (e *Entry) GetAlignment() float32             { return EntryGetAlignment(e) }
func (e *Entry) GetBuffer() *EntryBuffer           { return EntryGetBuffer(e) }
func (e *Entry) GetCompletion() *EntryCompletion   { return EntryGetCompletion(e) }
func (e *Entry) GetCurrentIconDragSource() int     { return EntryGetCurrentIconDragSource(e) }
func (e *Entry) GetCursorHadjustment() *Adjustment { return EntryGetCursorHadjustment(e) }
func (e *Entry) GetHasFrame() bool                 { return EntryGetHasFrame(e) }
func (e *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	return EntryGetIconActivatable(e, iconPos)
}
func (e *Entry) GetIconAtPos(x, y int) int                      { return EntryGetIconAtPos(e, x, y) }
func (e *Entry) GetIconGicon(iconPos EntryIconPosition) *I.Icon { return EntryGetIconGicon(e, iconPos) }
func (e *Entry) GetIconName(iconPos EntryIconPosition) string   { return EntryGetIconName(e, iconPos) }
func (e *Entry) GetIconPixbuf(iconPos EntryIconPosition) *D.Pixbuf {
	return EntryGetIconPixbuf(e, iconPos)
}
func (e *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	return EntryGetIconSensitive(e, iconPos)
}
func (e *Entry) GetIconStock(iconPos EntryIconPosition) string { return EntryGetIconStock(e, iconPos) }
func (e *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	return EntryGetIconStorageType(e, iconPos)
}
func (e *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) string {
	return EntryGetIconTooltipMarkup(e, iconPos)
}
func (e *Entry) GetIconTooltipText(iconPos EntryIconPosition) string {
	return EntryGetIconTooltipText(e, iconPos)
}
func (e *Entry) GetIconWindow(iconPos EntryIconPosition) *D.Window {
	return EntryGetIconWindow(e, iconPos)
}
func (e *Entry) GetInnerBorder() *Border       { return EntryGetInnerBorder(e) }
func (e *Entry) GetInvisibleChar() T.Gunichar  { return EntryGetInvisibleChar(e) }
func (e *Entry) GetLayout() *P.Layout          { return EntryGetLayout(e) }
func (e *Entry) GetLayoutOffsets(x, y *int)    { EntryGetLayoutOffsets(e, x, y) }
func (e *Entry) GetMaxLength() int             { return EntryGetMaxLength(e) }
func (e *Entry) GetOverwriteMode() bool        { return EntryGetOverwriteMode(e) }
func (e *Entry) GetProgressFraction() float64  { return EntryGetProgressFraction(e) }
func (e *Entry) GetProgressPulseStep() float64 { return EntryGetProgressPulseStep(e) }
func (e *Entry) GetText() string               { return EntryGetText(e) }
func (e *Entry) GetTextLength() uint16         { return EntryGetTextLength(e) }
func (e *Entry) GetTextWindow() *D.Window      { return EntryGetTextWindow(e) }
func (e *Entry) GetVisibility() bool           { return EntryGetVisibility(e) }
func (e *Entry) GetWidthChars() int            { return EntryGetWidthChars(e) }
func (e *Entry) ImContextFilterKeypress(event *D.EventKey) bool {
	return EntryImContextFilterKeypress(e, event)
}
func (e *Entry) LayoutIndexToTextIndex(layoutIndex int) int {
	return EntryLayoutIndexToTextIndex(e, layoutIndex)
}
func (e *Entry) PrependText(text string)                     { EntryPrependText(e, text) }
func (e *Entry) ProgressPulse()                              { EntryProgressPulse(e) }
func (e *Entry) ResetImContext()                             { EntryResetImContext(e) }
func (e *Entry) SelectRegion(start, end int)                 { EntrySelectRegion(e, start, end) }
func (e *Entry) SetActivatesDefault(setting bool)            { EntrySetActivatesDefault(e, setting) }
func (e *Entry) SetAlignment(xalign float32)                 { EntrySetAlignment(e, xalign) }
func (e *Entry) SetBuffer(buffer *EntryBuffer)               { EntrySetBuffer(e, buffer) }
func (e *Entry) SetCompletion(completion *EntryCompletion)   { EntrySetCompletion(e, completion) }
func (e *Entry) SetCursorHadjustment(adjustment *Adjustment) { EntrySetCursorHadjustment(e, adjustment) }
func (e *Entry) SetEditable(editable bool)                   { EntrySetEditable(e, editable) }
func (e *Entry) SetHasFrame(setting bool)                    { EntrySetHasFrame(e, setting) }
func (e *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	EntrySetIconActivatable(e, iconPos, activatable)
}
func (e *Entry) SetIconDragSource(iconPos EntryIconPosition, targetList *TargetList, actions D.DragAction) {
	EntrySetIconDragSource(e, iconPos, targetList, actions)
}
func (e *Entry) SetIconFromGicon(iconPos EntryIconPosition, icon *I.Icon) {
	EntrySetIconFromGicon(e, iconPos, icon)
}
func (e *Entry) SetIconFromIconName(iconPos EntryIconPosition, iconName string) {
	EntrySetIconFromIconName(e, iconPos, iconName)
}
func (e *Entry) SetIconFromPixbuf(iconPos EntryIconPosition, pixbuf *D.Pixbuf) {
	EntrySetIconFromPixbuf(e, iconPos, pixbuf)
}
func (e *Entry) SetIconFromStock(iconPos EntryIconPosition, stockId string) {
	EntrySetIconFromStock(e, iconPos, stockId)
}
func (e *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	EntrySetIconSensitive(e, iconPos, sensitive)
}
func (e *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	EntrySetIconTooltipMarkup(e, iconPos, tooltip)
}
func (e *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	EntrySetIconTooltipText(e, iconPos, tooltip)
}
func (e *Entry) SetInnerBorder(border *Border)         { EntrySetInnerBorder(e, border) }
func (e *Entry) SetInvisibleChar(ch T.Gunichar)        { EntrySetInvisibleChar(e, ch) }
func (e *Entry) SetMaxLength(max int)                  { EntrySetMaxLength(e, max) }
func (e *Entry) SetOverwriteMode(overwrite bool)       { EntrySetOverwriteMode(e, overwrite) }
func (e *Entry) SetPosition(position int)              { EntrySetPosition(e, position) }
func (e *Entry) SetProgressFraction(fraction float64)  { EntrySetProgressFraction(e, fraction) }
func (e *Entry) SetProgressPulseStep(fraction float64) { EntrySetProgressPulseStep(e, fraction) }
func (e *Entry) SetText(text string)                   { EntrySetText(e, text) }
func (e *Entry) SetVisibility(visible bool)            { EntrySetVisibility(e, visible) }
func (e *Entry) SetWidthChars(nChars int)              { EntrySetWidthChars(e, nChars) }
func (e *Entry) TextIndexToLayoutIndex(textIndex int) int {
	return EntryTextIndexToLayoutIndex(e, textIndex)
}
func (e *Entry) UnsetInvisibleChar() { EntryUnsetInvisibleChar(e) }

var (
	EntryBufferGetType func() O.Type
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

func (e *EntryBuffer) DeleteText(position uint, nChars int) uint {
	return EntryBufferDeleteText(e, position, nChars)
}
func (e *EntryBuffer) EmitDeletedText(position uint, nChars uint) {
	EntryBufferEmitDeletedText(e, position, nChars)
}
func (e *EntryBuffer) EmitInsertedText(position uint, chars string, nChars uint) {
	EntryBufferEmitInsertedText(e, position, chars, nChars)
}
func (e *EntryBuffer) GetBytes() T.Gsize { return EntryBufferGetBytes(e) }
func (e *EntryBuffer) GetLength() uint   { return EntryBufferGetLength(e) }
func (e *EntryBuffer) GetMaxLength() int { return EntryBufferGetMaxLength(e) }
func (e *EntryBuffer) GetText() string   { return EntryBufferGetText(e) }
func (e *EntryBuffer) InsertText(position uint, chars string, nChars int) uint {
	return EntryBufferInsertText(e, position, chars, nChars)
}
func (e *EntryBuffer) SetMaxLength(maxLength int)       { EntryBufferSetMaxLength(e, maxLength) }
func (e *EntryBuffer) SetText(chars string, nChars int) { EntryBufferSetText(e, chars, nChars) }

type (
	GtkEntryCompletion struct {
		Parent O.Object
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

	EntryCompletionComplete            func(completion *EntryCompletion)
	EntryCompletionDeleteAction        func(completion *EntryCompletion, index int)
	EntryCompletionGetCompletionPrefix func(completion *EntryCompletion) string
	EntryCompletionGetEntry            func(completion *EntryCompletion) *Widget
	EntryCompletionGetInlineCompletion func(completion *EntryCompletion) bool
	EntryCompletionGetInlineSelection  func(completion *EntryCompletion) bool
	EntryCompletionGetMinimumKeyLength func(completion *EntryCompletion) int
	EntryCompletionGetModel            func(completion *EntryCompletion) *TreeModel
	EntryCompletionGetPopupCompletion  func(completion *EntryCompletion) bool
	EntryCompletionGetPopupSetWidth    func(completion *EntryCompletion) bool
	EntryCompletionGetPopupSingleMatch func(completion *EntryCompletion) bool
	EntryCompletionGetTextColumn       func(completion *EntryCompletion) int
	EntryCompletionInsertActionMarkup  func(completion *EntryCompletion, index int, markup string)
	EntryCompletionInsertActionText    func(completion *EntryCompletion, index int, text string)
	EntryCompletionInsertPrefix        func(completion *EntryCompletion)
	EntryCompletionSetInlineCompletion func(completion *EntryCompletion, inlineCompletion bool)
	EntryCompletionSetInlineSelection  func(completion *EntryCompletion, inlineSelection bool)
	EntryCompletionSetMatchFunc        func(completion *EntryCompletion, f EntryCompletionMatchFunc, funcData T.Gpointer, funcNotify T.GDestroyNotify)
	EntryCompletionSetMinimumKeyLength func(completion *EntryCompletion, length int)
	EntryCompletionSetModel            func(completion *EntryCompletion, model *TreeModel)
	EntryCompletionSetPopupCompletion  func(completion *EntryCompletion, popupCompletion bool)
	EntryCompletionSetPopupSetWidth    func(completion *EntryCompletion, popupSetWidth bool)
	EntryCompletionSetPopupSingleMatch func(completion *EntryCompletion, popupSingleMatch bool)
	EntryCompletionSetTextColumn       func(completion *EntryCompletion, column int)
)

func (e *EntryCompletion) Complete()                   { EntryCompletionComplete(e) }
func (e *EntryCompletion) DeleteAction(index int)      { EntryCompletionDeleteAction(e, index) }
func (e *EntryCompletion) GetCompletionPrefix() string { return EntryCompletionGetCompletionPrefix(e) }
func (e *EntryCompletion) GetEntry() *Widget           { return EntryCompletionGetEntry(e) }
func (e *EntryCompletion) GetInlineCompletion() bool {
	return EntryCompletionGetInlineCompletion(e)
}
func (e *EntryCompletion) GetInlineSelection() bool { return EntryCompletionGetInlineSelection(e) }
func (e *EntryCompletion) GetMinimumKeyLength() int { return EntryCompletionGetMinimumKeyLength(e) }
func (e *EntryCompletion) GetModel() *TreeModel     { return EntryCompletionGetModel(e) }
func (e *EntryCompletion) GetPopupCompletion() bool { return EntryCompletionGetPopupCompletion(e) }
func (e *EntryCompletion) GetPopupSetWidth() bool   { return EntryCompletionGetPopupSetWidth(e) }
func (e *EntryCompletion) GetPopupSingleMatch() bool {
	return EntryCompletionGetPopupSingleMatch(e)
}
func (e *EntryCompletion) GetTextColumn() int { return EntryCompletionGetTextColumn(e) }
func (e *EntryCompletion) InsertActionMarkup(index int, markup string) {
	EntryCompletionInsertActionMarkup(e, index, markup)
}
func (e *EntryCompletion) InsertActionText(index int, text string) {
	EntryCompletionInsertActionText(e, index, text)
}
func (e *EntryCompletion) InsertPrefix() { EntryCompletionInsertPrefix(e) }
func (e *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	EntryCompletionSetInlineCompletion(e, inlineCompletion)
}
func (e *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	EntryCompletionSetInlineSelection(e, inlineSelection)
}
func (e *EntryCompletion) SetMatchFunc(f EntryCompletionMatchFunc, funcData T.Gpointer, funcNotify T.GDestroyNotify) {
	EntryCompletionSetMatchFunc(e, f, funcData, funcNotify)
}
func (e *EntryCompletion) SetMinimumKeyLength(length int) {
	EntryCompletionSetMinimumKeyLength(e, length)
}
func (e *EntryCompletion) SetModel(model *TreeModel) { EntryCompletionSetModel(e, model) }
func (e *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	EntryCompletionSetPopupCompletion(e, popupCompletion)
}
func (e *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	EntryCompletionSetPopupSetWidth(e, popupSetWidth)
}
func (e *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	EntryCompletionSetPopupSingleMatch(e, popupSingleMatch)
}
func (e *EntryCompletion) SetTextColumn(column int) { EntryCompletionSetTextColumn(e, column) }

type EnumValue O.EnumValue

type EventBox struct{ Bin Bin }

var (
	EventBoxGetType func() O.Type
	EventBoxNew     func() *Widget

	EventBoxGetAboveChild    func(e *EventBox) bool
	EventBoxGetVisibleWindow func(e *EventBox) bool
	EventBoxSetAboveChild    func(e *EventBox, aboveChild bool)
	EventBoxSetVisibleWindow func(e *EventBox, visibleWindow bool)
)

func (e *EventBox) GetAboveChild() bool           { return EventBoxGetAboveChild(e) }
func (e *EventBox) GetVisibleWindow() bool        { return EventBoxGetVisibleWindow(e) }
func (e *EventBox) SetAboveChild(aboveChild bool) { EventBoxSetAboveChild(e, aboveChild) }
func (e *EventBox) SetVisibleWindow(visibleWindow bool) {
	EventBoxSetVisibleWindow(e, visibleWindow)
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

	ExpanderSetExpanded     func(e *Expander, expanded bool)
	ExpanderGetExpanded     func(e *Expander) bool
	ExpanderSetSpacing      func(e *Expander, spacing int)
	ExpanderGetSpacing      func(e *Expander) int
	ExpanderSetLabel        func(e *Expander, label string)
	ExpanderGetLabel        func(e *Expander) string
	ExpanderSetUseUnderline func(e *Expander, useUnderline bool)
	ExpanderGetUseUnderline func(e *Expander) bool
	ExpanderSetUseMarkup    func(e *Expander, useMarkup bool)
	ExpanderGetUseMarkup    func(e *Expander) bool
	ExpanderSetLabelWidget  func(e *Expander, labelWidget *Widget)
	ExpanderGetLabelWidget  func(e *Expander) *Widget
	ExpanderSetLabelFill    func(e *Expander, labelFill bool)
	ExpanderGetLabelFill    func(e *Expander) bool
)

func (e *Expander) GetExpanded() bool                  { return ExpanderGetExpanded(e) }
func (e *Expander) GetLabel() string                   { return ExpanderGetLabel(e) }
func (e *Expander) GetLabelFill() bool                 { return ExpanderGetLabelFill(e) }
func (e *Expander) GetLabelWidget() *Widget            { return ExpanderGetLabelWidget(e) }
func (e *Expander) GetSpacing() int                    { return ExpanderGetSpacing(e) }
func (e *Expander) GetUseMarkup() bool                 { return ExpanderGetUseMarkup(e) }
func (e *Expander) GetUseUnderline() bool              { return ExpanderGetUseUnderline(e) }
func (e *Expander) SetExpanded(expanded bool)          { ExpanderSetExpanded(e, expanded) }
func (e *Expander) SetLabel(label string)              { ExpanderSetLabel(e, label) }
func (e *Expander) SetLabelFill(labelFill bool)        { ExpanderSetLabelFill(e, labelFill) }
func (e *Expander) SetLabelWidget(labelWidget *Widget) { ExpanderSetLabelWidget(e, labelWidget) }
func (e *Expander) SetSpacing(spacing int)             { ExpanderSetSpacing(e, spacing) }
func (e *Expander) SetUseMarkup(useMarkup bool)        { ExpanderSetUseMarkup(e, useMarkup) }
func (e *Expander) SetUseUnderline(useUnderline bool)  { ExpanderSetUseUnderline(e, useUnderline) }
