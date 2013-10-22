// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type (
	Table struct {
		Container     T.GtkContainer
		Children      *T.GList
		Rows          *TableRowCol
		Cols          *TableRowCol
		Nrows         uint16
		Ncols         uint16
		ColumnSpacing uint16
		RowSpacing    uint16
		Bits          uint
		// Homogeneous : 1
	}

	TableRowCol struct {
		Requisition uint16
		Allocation  uint16
		Spacing     uint16
		Bits        uint
		// NeedExpand : 1
		// NeedShrink : 1
		// Expand : 1
		// Shrink : 1
		// Empty : 1
	}
)

var (
	TableGetType func() T.GType
	TableNew     func(rows, columns uint, homogeneous T.Gboolean) *Widget

	tableAttach               func(t *Table, child *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint, xoptions, yoptions T.GtkAttachOptions, xpadding, ypadding uint)
	tableAttachDefaults       func(t *Table, widget *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint)
	tableGetColSpacing        func(t *Table, column uint) uint
	tableGetDefaultColSpacing func(t *Table) uint
	tableGetDefaultRowSpacing func(t *Table) uint
	tableGetHomogeneous       func(t *Table) T.Gboolean
	tableGetRowSpacing        func(t *Table, row uint) uint
	tableGetSize              func(t *Table, rows, columns *uint)
	tableResize               func(t *Table, rows, columns uint)
	tableSetColSpacing        func(t *Table, column, spacing uint)
	tableSetColSpacings       func(t *Table, spacing uint)
	tableSetHomogeneous       func(t *Table, homogeneous T.Gboolean)
	tableSetRowSpacing        func(t *Table, row, spacing uint)
	tableSetRowSpacings       func(t *Table, spacing uint)
)

func (t *Table) Attach(child *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint, xoptions, yoptions T.GtkAttachOptions, xpadding, ypadding uint) {
	tableAttach(t, child, leftAttach, rightAttach, topAttach, bottomAttach, xoptions, yoptions, xpadding, ypadding)
}
func (t *Table) AttachDefaults(widget *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint) {
	tableAttachDefaults(t, widget, leftAttach, rightAttach, topAttach, bottomAttach)
}
func (t *Table) GetColSpacing(column uint) uint        { return tableGetColSpacing(t, column) }
func (t *Table) GetDefaultColSpacing() uint            { return tableGetDefaultColSpacing(t) }
func (t *Table) GetDefaultRowSpacing() uint            { return tableGetDefaultRowSpacing(t) }
func (t *Table) GetHomogeneous() T.Gboolean            { return tableGetHomogeneous(t) }
func (t *Table) GetRowSpacing(row uint) uint           { return tableGetRowSpacing(t, row) }
func (t *Table) GetSize(rows, columns *uint)           { tableGetSize(t, rows, columns) }
func (t *Table) Resize(rows, columns uint)             { tableResize(t, rows, columns) }
func (t *Table) SetColSpacing(column, spacing uint)    { tableSetColSpacing(t, column, spacing) }
func (t *Table) SetColSpacings(spacing uint)           { tableSetColSpacings(t, spacing) }
func (t *Table) SetHomogeneous(homogeneous T.Gboolean) { tableSetHomogeneous(t, homogeneous) }
func (t *Table) SetRowSpacing(row, spacing uint)       { tableSetRowSpacing(t, row, spacing) }
func (t *Table) SetRowSpacings(spacing uint)           { tableSetRowSpacings(t, spacing) }

type TargetList struct {
	List     *T.GList
	RefCount uint
}

var (
	TargetListGetType func() T.GType
	TargetListNew     func(targets *T.GtkTargetEntry, ntargets uint) *TargetList

	TargetTableFree        func(targets *T.GtkTargetEntry, nTargets int)
	TargetTableNewFromList func(tl *TargetList, nTargets *int) *T.GtkTargetEntry

	targetListRef                func(t *TargetList) *TargetList
	targetListUnref              func(t *TargetList)
	targetListAdd                func(t *TargetList, target T.GdkAtom, flags, info uint)
	targetListAddTextTargets     func(t *TargetList, info uint)
	targetListAddRichTextTargets func(t *TargetList, info uint, deserializable T.Gboolean, buffer *TextBuffer)
	targetListAddImageTargets    func(t *TargetList, info uint, writable T.Gboolean)
	targetListAddUriTargets      func(t *TargetList, info uint)
	targetListAddTable           func(t *TargetList, targets *T.GtkTargetEntry, ntargets uint)
	targetListRemove             func(t *TargetList, target T.GdkAtom)
	targetListFind               func(t *TargetList, target T.GdkAtom, info *uint) T.Gboolean
)

func (t *TargetList) Ref() *TargetList                       { return targetListRef(t) }
func (t *TargetList) Unref()                                 { targetListUnref(t) }
func (t *TargetList) Add(target T.GdkAtom, flags, info uint) { targetListAdd(t, target, flags, info) }
func (t *TargetList) AddTextTargets(info uint)               { targetListAddTextTargets(t, info) }
func (t *TargetList) AddRichTextTargets(info uint, deserializable T.Gboolean, buffer *TextBuffer) {
	targetListAddRichTextTargets(t, info, deserializable, buffer)
}
func (t *TargetList) AddImageTargets(info uint, writable T.Gboolean) {
	targetListAddImageTargets(t, info, writable)
}
func (t *TargetList) AddUriTargets(info uint) { targetListAddUriTargets(t, info) }
func (t *TargetList) AddTable(targets *T.GtkTargetEntry, ntargets uint) {
	targetListAddTable(t, targets, ntargets)
}
func (t *TargetList) Remove(target T.GdkAtom) { targetListRemove(t, target) }
func (t *TargetList) Find(target T.GdkAtom, info *uint) T.Gboolean {
	return targetListFind(t, target, info)
}

type TipsQuery struct {
	Label Label
	Bits  uint
	// EmitAlways : 1
	// InQuery : 1
	LabelInactive *T.Gchar
	LabelNoTip    *T.Gchar
	Caller        *Widget
	LastCrossed   *Widget
	QueryCursor   *T.GdkCursor
}

var (
	TipsQueryGetType func() T.GType
	TipsQueryNew     func() *Widget

	tipsQuerySetCaller  func(t *TipsQuery, caller *Widget)
	tipsQuerySetLabels  func(t *TipsQuery, labelInactive, labelNoTip string)
	tipsQueryStartQuery func(t *TipsQuery)
	tipsQueryStopQuery  func(t *TipsQuery)
)

func (t *TipsQuery) SetCaller(caller *Widget) { tipsQuerySetCaller(t, caller) }
func (t *TipsQuery) SetLabels(labelInactive, labelNoTip string) {
	tipsQuerySetLabels(t, labelInactive, labelNoTip)
}
func (t *TipsQuery) StartQuery() { tipsQueryStartQuery(t) }
func (t *TipsQuery) StopQuery()  { tipsQueryStopQuery(t) }

type ToggleAction struct {
	Parent       Action
	Private_data *T.GtkToggleActionPrivate
}

var (
	ToggleActionGetType func() T.GType
	ToggleActionNew     func(name, label, tooltip, stockId string) *T.GtkToggleAction

	toggleActionGetActive      func(t *ToggleAction) T.Gboolean
	toggleActionGetDrawAsRadio func(t *ToggleAction) T.Gboolean
	toggleActionSetActive      func(t *ToggleAction, isActive T.Gboolean)
	toggleActionSetDrawAsRadio func(t *ToggleAction, drawAsRadio T.Gboolean)
	toggleActionToggled        func(t *ToggleAction)
)

func (t *ToggleAction) GetActive() T.Gboolean         { return toggleActionGetActive(t) }
func (t *ToggleAction) GetDrawAsRadio() T.Gboolean    { return toggleActionGetDrawAsRadio(t) }
func (t *ToggleAction) SetActive(isActive T.Gboolean) { toggleActionSetActive(t, isActive) }
func (t *ToggleAction) SetDrawAsRadio(drawAsRadio T.Gboolean) {
	toggleActionSetDrawAsRadio(t, drawAsRadio)
}
func (t *ToggleAction) Toggled() { toggleActionToggled(t) }

type ToggleButton struct {
	Button Button
	Bits   uint
	// Active : 1
	// DrawIndicator : 1
	// Inconsistent : 1
}

var (
	ToggleButtonGetType         func() T.GType
	ToggleButtonNew             func() *Widget
	ToggleButtonNewWithLabel    func(label string) *Widget
	ToggleButtonNewWithMnemonic func(label string) *Widget

	toggleButtonGetActive       func(t *ToggleButton) T.Gboolean
	toggleButtonGetInconsistent func(t *ToggleButton) T.Gboolean
	toggleButtonGetMode         func(t *ToggleButton) T.Gboolean
	toggleButtonSetActive       func(t *ToggleButton, isActive T.Gboolean)
	toggleButtonSetInconsistent func(t *ToggleButton, setting T.Gboolean)
	toggleButtonSetMode         func(t *ToggleButton, drawIndicator T.Gboolean)
	toggleButtonToggled         func(t *ToggleButton)
)

func (t *ToggleButton) GetActive() T.Gboolean              { return toggleButtonGetActive(t) }
func (t *ToggleButton) GetInconsistent() T.Gboolean        { return toggleButtonGetInconsistent(t) }
func (t *ToggleButton) GetMode() T.Gboolean                { return toggleButtonGetMode(t) }
func (t *ToggleButton) SetActive(isActive T.Gboolean)      { toggleButtonSetActive(t, isActive) }
func (t *ToggleButton) SetInconsistent(setting T.Gboolean) { toggleButtonSetInconsistent(t, setting) }
func (t *ToggleButton) SetMode(drawIndicator T.Gboolean)   { toggleButtonSetMode(t, drawIndicator) }
func (t *ToggleButton) Toggled()                           { toggleButtonToggled(t) }

type ToggleToolButton struct {
	Parent ToolButton
	_      *struct{}
}

var (
	ToggleToolButtonGetType      func() T.GType
	ToggleToolButtonNew          func() *ToolItem
	ToggleToolButtonNewFromStock func(stockId string) *ToolItem

	toggleToolButtonGetActive func(t *ToggleToolButton) T.Gboolean
	toggleToolButtonSetActive func(t *ToggleToolButton, isActive T.Gboolean)
)

func (t *ToggleToolButton) GetActive() T.Gboolean         { return toggleToolButtonGetActive(t) }
func (t *ToggleToolButton) SetActive(isActive T.Gboolean) { toggleToolButtonSetActive(t, isActive) }

type (
	Toolbar struct {
		Container   T.GtkContainer
		NumChildren int
		Children    *T.GList
		Orientation T.GtkOrientation
		Style       ToolbarStyle
		IconSize    IconSize
		Tooltips    *Tooltips
		ButtonMaxw  int
		ButtonMaxh  int
		_, _        uint
		Bits        uint
		// StyleSet : 1
		// IconSizeSet : 1
	}
)

type TextAttributes struct {
	Refcount         uint
	Appearance       T.GtkTextAppearance
	Justification    T.GtkJustification
	Direction        T.GtkTextDirection
	Font             *T.PangoFontDescription
	FontScale        float64
	LeftMargin       int
	Indent           int
	RightMargin      int
	PixelsAboveLines int
	PixelsBelowLines int
	PixelsInsideWrap int
	Tabs             *T.PangoTabArray
	WrapMode         T.GtkWrapMode
	Language         *T.PangoLanguage
	PgBgColor        *T.GdkColor
	Bits             uint
	// Invisible : 1
	// BgFullHeight : 1
	// Editable : 1
	// Realized : 1
	// Pad1 : 1
	// Pad2 : 1
	// Pad3 : 1
	// Pad4 : 1
}

var (
	TextAttributesGetType func() T.GType
	TextAttributesNew     func() *TextAttributes

	textAttributesCopy       func(t *TextAttributes) *TextAttributes
	textAttributesCopyValues func(t *TextAttributes, dest *TextAttributes)
	textAttributesRef        func(t *TextAttributes) *TextAttributes
	textAttributesUnref      func(t *TextAttributes)
)

func (t *TextAttributes) Copy() *TextAttributes           { return textAttributesCopy(t) }
func (t *TextAttributes) CopyValues(dest *TextAttributes) { textAttributesCopyValues(t, dest) }
func (t *TextAttributes) Ref() *TextAttributes            { return textAttributesRef(t) }
func (t *TextAttributes) Unref()                          { textAttributesUnref(t) }

type (
	TextBuffer struct {
		Parent                   T.GObject
		TagTable                 *TextTagTable
		Btree                    *T.GtkTextBTree
		ClipboardContentsBuffers *T.GSList
		SelectionClipboards      *T.GSList
		LogAttrCache             *T.GtkTextLogAttrCache
		UserActionCount          uint
		Bits                     uint
		// Modified : 1
		// HasSelection : 1
	}

	TextBufferSerializeFunc func(
		registerBuffer, contentBuffer *TextBuffer,
		start, end *TextIter,
		length *T.Gsize, userData T.Gpointer) *uint8

	TextBufferDeserializeFunc func(
		registerBuffer, contentBuffer *TextBuffer,
		iter *TextIter, data *uint8, length T.Gsize,
		createTags T.Gboolean, userData T.Gpointer,
		e **T.GError) T.Gboolean
)

var (
	TextBufferGetType           func() T.GType
	TextBufferNew               func(table *TextTagTable) *TextBuffer
	TextBufferTargetInfoGetType func() T.GType

	textBufferAddMark                     func(t *TextBuffer, mark *TextMark, where *TextIter)
	textBufferAddSelectionClipboard       func(t *TextBuffer, clipboard *Clipboard)
	textBufferApplyTag                    func(t *TextBuffer, tag *TextTag, start, end *TextIter)
	textBufferApplyTagByName              func(t *TextBuffer, name string, start, end *TextIter)
	textBufferBackspace                   func(t *TextBuffer, iter *TextIter, interactive, defaultEditable T.Gboolean) T.Gboolean
	textBufferBeginUserAction             func(t *TextBuffer)
	textBufferCopyClipboard               func(t *TextBuffer, clipboard *Clipboard)
	textBufferCreateChildAnchor           func(t *TextBuffer, iter *TextIter) *TextChildAnchor
	textBufferCreateMark                  func(t *TextBuffer, markName string, where *TextIter, leftGravity T.Gboolean) *TextMark
	textBufferCreateTag                   func(t *TextBuffer, tagName, firstPropertyName string, v ...VArg) *TextTag
	textBufferCutClipboard                func(t *TextBuffer, clipboard *Clipboard, defaultEditable T.Gboolean)
	textBufferDelete                      func(t *TextBuffer, start, end *TextIter)
	textBufferDeleteInteractive           func(t *TextBuffer, startIter, endIter *TextIter, defaultEditable T.Gboolean) T.Gboolean
	textBufferDeleteMark                  func(t *TextBuffer, mark *TextMark)
	textBufferDeleteMarkByName            func(t *TextBuffer, name string)
	textBufferDeleteSelection             func(t *TextBuffer, interactive, defaultEditable T.Gboolean) T.Gboolean
	textBufferDeserialize                 func(t *TextBuffer, contentBuffer *TextBuffer, format T.GdkAtom, iter *TextIter, data *uint8, length T.Gsize, error **T.GError) T.Gboolean
	textBufferDeserializeGetCanCreateTags func(t *TextBuffer, format T.GdkAtom) T.Gboolean
	textBufferDeserializeSetCanCreateTags func(t *TextBuffer, format T.GdkAtom, canCreateTags T.Gboolean)
	textBufferEndUserAction               func(t *TextBuffer)
	textBufferGetBounds                   func(t *TextBuffer, start, end *TextIter)
	textBufferGetCharCount                func(t *TextBuffer) int
	textBufferGetCopyTargetList           func(t *TextBuffer) *TargetList
	textBufferGetDeserializeFormats       func(t *TextBuffer, nFormats *int) *T.GdkAtom
	textBufferGetEndIter                  func(t *TextBuffer, iter *TextIter)
	textBufferGetHasSelection             func(t *TextBuffer) T.Gboolean
	textBufferGetInsert                   func(t *TextBuffer) *TextMark
	textBufferGetIterAtChildAnchor        func(t *TextBuffer, iter *TextIter, anchor *TextChildAnchor)
	textBufferGetIterAtLine               func(t *TextBuffer, iter *TextIter, lineNumber int)
	textBufferGetIterAtLineIndex          func(t *TextBuffer, iter *TextIter, lineNumber, byteIndex int)
	textBufferGetIterAtLineOffset         func(t *TextBuffer, iter *TextIter, lineNumber, charOffset int)
	textBufferGetIterAtMark               func(t *TextBuffer, iter *TextIter, mark *TextMark)
	textBufferGetIterAtOffset             func(t *TextBuffer, iter *TextIter, charOffset int)
	textBufferGetLineCount                func(t *TextBuffer) int
	textBufferGetMark                     func(t *TextBuffer, name string) *TextMark
	textBufferGetModified                 func(t *TextBuffer) T.Gboolean
	textBufferGetPasteTargetList          func(t *TextBuffer) *TargetList
	textBufferGetSelectionBound           func(t *TextBuffer) *TextMark
	textBufferGetSelectionBounds          func(t *TextBuffer, start, end *TextIter) T.Gboolean
	textBufferGetSerializeFormats         func(t *TextBuffer, nFormats *int) *T.GdkAtom
	textBufferGetSlice                    func(t *TextBuffer, start, end *TextIter, includeHiddenChars T.Gboolean) string
	textBufferGetStartIter                func(t *TextBuffer, iter *TextIter)
	textBufferGetTagTable                 func(t *TextBuffer) *TextTagTable
	textBufferGetText                     func(t *TextBuffer, start, end *TextIter, includeHiddenChars T.Gboolean) string
	textBufferInsert                      func(t *TextBuffer, iter *TextIter, text string, leng int)
	textBufferInsertAtCursor              func(t *TextBuffer, text string, leng int)
	textBufferInsertChildAnchor           func(t *TextBuffer, iter *TextIter, anchor *TextChildAnchor)
	textBufferInsertInteractive           func(t *TextBuffer, iter *TextIter, text string, leng int, defaultEditable T.Gboolean) T.Gboolean
	textBufferInsertInteractiveAtCursor   func(t *TextBuffer, text string, leng int, defaultEditable T.Gboolean) T.Gboolean
	textBufferInsertPixbuf                func(t *TextBuffer, iter *TextIter, pixbuf *T.GdkPixbuf)
	textBufferInsertRange                 func(t *TextBuffer, iter, start, end *TextIter)
	textBufferInsertRangeInteractive      func(t *TextBuffer, iter, start, end *TextIter, defaultEditable T.Gboolean) T.Gboolean
	textBufferInsertWithTags              func(t *TextBuffer, iter *TextIter, text string, leng int, firstTag *TextTag, v ...VArg)
	textBufferInsertWithTagsByName        func(t *TextBuffer, iter *TextIter, text string, leng int, firstTagName string, v ...VArg)
	textBufferMoveMark                    func(t *TextBuffer, mark *TextMark, where *TextIter)
	textBufferMoveMarkByName              func(t *TextBuffer, name string, where *TextIter)
	textBufferPasteClipboard              func(t *TextBuffer, clipboard *Clipboard, overrideLocation *TextIter, defaultEditable T.Gboolean)
	textBufferPlaceCursor                 func(t *TextBuffer, where *TextIter)
	textBufferRegisterDeserializeFormat   func(t *TextBuffer, mimeType string, function TextBufferDeserializeFunc, userData T.Gpointer, userDataDestroy T.GDestroyNotify) T.GdkAtom
	textBufferRegisterDeserializeTagset   func(t *TextBuffer, tagsetName string) T.GdkAtom
	textBufferRegisterSerializeFormat     func(t *TextBuffer, mimeType string, function TextBufferSerializeFunc, userData T.Gpointer, userDataDestroy T.GDestroyNotify) T.GdkAtom
	textBufferRegisterSerializeTagset     func(t *TextBuffer, tagsetName string) T.GdkAtom
	textBufferRemoveAllTags               func(t *TextBuffer, start, end *TextIter)
	textBufferRemoveSelectionClipboard    func(t *TextBuffer, clipboard *Clipboard)
	textBufferRemoveTag                   func(t *TextBuffer, tag *TextTag, start, end *TextIter)
	textBufferRemoveTagByName             func(t *TextBuffer, name string, start, end *TextIter)
	textBufferSelectRange                 func(t *TextBuffer, ins, bound *TextIter)
	textBufferSerialize                   func(t *TextBuffer, contentBuffer *TextBuffer, format T.GdkAtom, start, end *TextIter, length *T.Gsize) *uint8
	textBufferSetModified                 func(t *TextBuffer, setting T.Gboolean)
	textBufferSetText                     func(t *TextBuffer, text string, leng int)
	textBufferUnregisterDeserializeFormat func(t *TextBuffer, format T.GdkAtom)
	textBufferUnregisterSerializeFormat   func(t *TextBuffer, format T.GdkAtom)
)

func (t *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	textBufferAddMark(t, mark, where)
}
func (t *TextBuffer) AddSelectionClipboard(clipboard *Clipboard) {
	textBufferAddSelectionClipboard(t, clipboard)
}
func (t *TextBuffer) ApplyTag(tag *TextTag, start, end *TextIter) {
	textBufferApplyTag(t, tag, start, end)
}
func (t *TextBuffer) ApplyTagByName(name string, start, end *TextIter) {
	textBufferApplyTagByName(t, name, start, end)
}
func (t *TextBuffer) Backspace(iter *TextIter, interactive, defaultEditable T.Gboolean) T.Gboolean {
	return textBufferBackspace(t, iter, interactive, defaultEditable)
}
func (t *TextBuffer) BeginUserAction()                   { textBufferBeginUserAction(t) }
func (t *TextBuffer) CopyClipboard(clipboard *Clipboard) { textBufferCopyClipboard(t, clipboard) }
func (t *TextBuffer) CreateChildAnchor(iter *TextIter) *TextChildAnchor {
	return textBufferCreateChildAnchor(t, iter)
}
func (t *TextBuffer) CreateMark(markName string, where *TextIter, leftGravity T.Gboolean) *TextMark {
	return textBufferCreateMark(t, markName, where, leftGravity)
}
func (t *TextBuffer) CreateTag(tagName string, firstPropertyName string, v ...VArg) *TextTag {
	return textBufferCreateTag(t, tagName, firstPropertyName, v)
}
func (t *TextBuffer) CutClipboard(clipboard *Clipboard, defaultEditable T.Gboolean) {
	textBufferCutClipboard(t, clipboard, defaultEditable)
}
func (t *TextBuffer) Delete(start, end *TextIter) { textBufferDelete(t, start, end) }
func (t *TextBuffer) DeleteInteractive(startIter, endIter *TextIter, defaultEditable T.Gboolean) T.Gboolean {
	return textBufferDeleteInteractive(t, startIter, endIter, defaultEditable)
}
func (t *TextBuffer) DeleteMark(mark *TextMark)    { textBufferDeleteMark(t, mark) }
func (t *TextBuffer) DeleteMarkByName(name string) { textBufferDeleteMarkByName(t, name) }
func (t *TextBuffer) DeleteSelection(interactive, defaultEditable T.Gboolean) T.Gboolean {
	return textBufferDeleteSelection(t, interactive, defaultEditable)
}
func (t *TextBuffer) Deserialize(contentBuffer *TextBuffer, format T.GdkAtom, iter *TextIter, data *uint8, length T.Gsize, err **T.GError) T.Gboolean {
	return textBufferDeserialize(t, contentBuffer, format, iter, data, length, err)
}
func (t *TextBuffer) DeserializeGetCanCreateTags(format T.GdkAtom) T.Gboolean {
	return textBufferDeserializeGetCanCreateTags(t, format)
}
func (t *TextBuffer) DeserializeSetCanCreateTags(format T.GdkAtom, canCreateTags T.Gboolean) {
	textBufferDeserializeSetCanCreateTags(t, format, canCreateTags)
}
func (t *TextBuffer) EndUserAction()                 { textBufferEndUserAction(t) }
func (t *TextBuffer) GetBounds(start, end *TextIter) { textBufferGetBounds(t, start, end) }
func (t *TextBuffer) GetCharCount() int              { return textBufferGetCharCount(t) }
func (t *TextBuffer) GetCopyTargetList() *TargetList { return textBufferGetCopyTargetList(t) }
func (t *TextBuffer) GetDeserializeFormats(nFormats *int) *T.GdkAtom {
	return textBufferGetDeserializeFormats(t, nFormats)
}
func (t *TextBuffer) GetEndIter(iter *TextIter)   { textBufferGetEndIter(t, iter) }
func (t *TextBuffer) GetHasSelection() T.Gboolean { return textBufferGetHasSelection(t) }
func (t *TextBuffer) GetInsert() *TextMark        { return textBufferGetInsert(t) }
func (t *TextBuffer) GetIterAtChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	textBufferGetIterAtChildAnchor(t, iter, anchor)
}
func (t *TextBuffer) GetIterAtLine(iter *TextIter, lineNumber int) {
	textBufferGetIterAtLine(t, iter, lineNumber)
}
func (t *TextBuffer) GetIterAtLineIndex(iter *TextIter, lineNumber, byteIndex int) {
	textBufferGetIterAtLineIndex(t, iter, lineNumber, byteIndex)
}
func (t *TextBuffer) GetIterAtLineOffset(iter *TextIter, lineNumber, charOffset int) {
	textBufferGetIterAtLineOffset(t, iter, lineNumber, charOffset)
}
func (t *TextBuffer) GetIterAtMark(iter *TextIter, mark *TextMark) {
	textBufferGetIterAtMark(t, iter, mark)
}
func (t *TextBuffer) GetIterAtOffset(iter *TextIter, charOffset int) {
	textBufferGetIterAtOffset(t, iter, charOffset)
}
func (t *TextBuffer) GetLineCount() int               { return textBufferGetLineCount(t) }
func (t *TextBuffer) GetMark(name string) *TextMark   { return textBufferGetMark(t, name) }
func (t *TextBuffer) GetModified() T.Gboolean         { return textBufferGetModified(t) }
func (t *TextBuffer) GetPasteTargetList() *TargetList { return textBufferGetPasteTargetList(t) }
func (t *TextBuffer) GetSelectionBound() *TextMark    { return textBufferGetSelectionBound(t) }
func (t *TextBuffer) GetSelectionBounds(start, end *TextIter) T.Gboolean {
	return textBufferGetSelectionBounds(t, start, end)
}
func (t *TextBuffer) GetSerializeFormats(nFormats *int) *T.GdkAtom {
	return textBufferGetSerializeFormats(t, nFormats)
}
func (t *TextBuffer) GetSlice(start, end *TextIter, includeHiddenChars T.Gboolean) string {
	return textBufferGetSlice(t, start, end, includeHiddenChars)
}
func (t *TextBuffer) GetStartIter(iter *TextIter) { textBufferGetStartIter(t, iter) }
func (t *TextBuffer) GetTagTable() *TextTagTable  { return textBufferGetTagTable(t) }
func (t *TextBuffer) GetText(start, end *TextIter, includeHiddenChars T.Gboolean) string {
	return textBufferGetText(t, start, end, includeHiddenChars)
}
func (t *TextBuffer) Insert(iter *TextIter, text string, leng int) {
	textBufferInsert(t, iter, text, leng)
}
func (t *TextBuffer) InsertAtCursor(text string, leng int) { textBufferInsertAtCursor(t, text, leng) }
func (t *TextBuffer) InsertChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	textBufferInsertChildAnchor(t, iter, anchor)
}
func (t *TextBuffer) InsertInteractive(iter *TextIter, text string, leng int, defaultEditable T.Gboolean) T.Gboolean {
	return textBufferInsertInteractive(t, iter, text, leng, defaultEditable)
}
func (t *TextBuffer) InsertInteractiveAtCursor(text string, leng int, defaultEditable T.Gboolean) T.Gboolean {
	return textBufferInsertInteractiveAtCursor(t, text, leng, defaultEditable)
}
func (t *TextBuffer) InsertPixbuf(iter *TextIter, pixbuf *T.GdkPixbuf) {
	textBufferInsertPixbuf(t, iter, pixbuf)
}
func (t *TextBuffer) InsertRange(iter, start, end *TextIter) {
	textBufferInsertRange(t, iter, start, end)
}
func (t *TextBuffer) InsertRangeInteractive(iter, start, end *TextIter, defaultEditable T.Gboolean) T.Gboolean {
	return textBufferInsertRangeInteractive(t, iter, start, end, defaultEditable)
}
func (t *TextBuffer) InsertWithTags(iter *TextIter, text string, leng int, firstTag *TextTag, v ...VArg) {
	textBufferInsertWithTags(t, iter, text, leng, firstTag, v)
}
func (t *TextBuffer) InsertWithTagsByName(iter *TextIter, text string, leng int, firstTagName string, v ...VArg) {
	textBufferInsertWithTagsByName(t, iter, text, leng, firstTagName, v)
}
func (t *TextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	textBufferMoveMark(t, mark, where)
}
func (t *TextBuffer) MoveMarkByName(name string, where *TextIter) {
	textBufferMoveMarkByName(t, name, where)
}
func (t *TextBuffer) PasteClipboard(clipboard *Clipboard, overrideLocation *TextIter, defaultEditable T.Gboolean) {
	textBufferPasteClipboard(t, clipboard, overrideLocation, defaultEditable)
}
func (t *TextBuffer) PlaceCursor(where *TextIter) { textBufferPlaceCursor(t, where) }
func (t *TextBuffer) RegisterDeserializeFormat(mimeType string, function TextBufferDeserializeFunc, userData T.Gpointer, userDataDestroy T.GDestroyNotify) T.GdkAtom {
	return textBufferRegisterDeserializeFormat(t, mimeType, function, userData, userDataDestroy)
}
func (t *TextBuffer) RegisterDeserializeTagset(tagsetName string) T.GdkAtom {
	return textBufferRegisterDeserializeTagset(t, tagsetName)
}
func (t *TextBuffer) RegisterSerializeFormat(mimeType string, function TextBufferSerializeFunc, userData T.Gpointer, userDataDestroy T.GDestroyNotify) T.GdkAtom {
	return textBufferRegisterSerializeFormat(t, mimeType, function, userData, userDataDestroy)
}
func (t *TextBuffer) RegisterSerializeTagset(tagsetName string) T.GdkAtom {
	return textBufferRegisterSerializeTagset(t, tagsetName)
}
func (t *TextBuffer) RemoveAllTags(start, end *TextIter) { textBufferRemoveAllTags(t, start, end) }
func (t *TextBuffer) RemoveSelectionClipboard(clipboard *Clipboard) {
	textBufferRemoveSelectionClipboard(t, clipboard)
}
func (t *TextBuffer) RemoveTag(tag *TextTag, start, end *TextIter) {
	textBufferRemoveTag(t, tag, start, end)
}
func (t *TextBuffer) RemoveTagByName(name string, start, end *TextIter) {
	textBufferRemoveTagByName(t, name, start, end)
}
func (t *TextBuffer) SelectRange(ins, bound *TextIter) { textBufferSelectRange(t, ins, bound) }
func (t *TextBuffer) Serialize(contentBuffer *TextBuffer, format T.GdkAtom, start, end *TextIter, length *T.Gsize) *uint8 {
	return textBufferSerialize(t, contentBuffer, format, start, end, length)
}
func (t *TextBuffer) SetModified(setting T.Gboolean) { textBufferSetModified(t, setting) }
func (t *TextBuffer) SetText(text string, leng int)  { textBufferSetText(t, text, leng) }
func (t *TextBuffer) UnregisterDeserializeFormat(format T.GdkAtom) {
	textBufferUnregisterDeserializeFormat(t, format)
}
func (t *TextBuffer) UnregisterSerializeFormat(format T.GdkAtom) {
	textBufferUnregisterSerializeFormat(t, format)
}

type TextChildAnchor struct {
	Parent  T.GObject
	Segment T.Gpointer
}

var (
	TextChildAnchorGetType func() T.GType
	TextChildAnchorNew     func() *TextChildAnchor

	textChildAnchorGetDeleted func(t *TextChildAnchor) T.Gboolean
	textChildAnchorGetWidgets func(t *TextChildAnchor) *T.GList
)

func (t *TextChildAnchor) GetDeleted() T.Gboolean { return textChildAnchorGetDeleted(t) }
func (t *TextChildAnchor) GetWidgets() *T.GList   { return textChildAnchorGetWidgets(t) }

type TextIter struct { //TODO(t): ALL _?
	Dummy1  T.Gpointer
	Dummy2  T.Gpointer
	Dummy3  int
	Dummy4  int
	Dummy5  int
	Dummy6  int
	Dummy7  int
	Dummy8  int
	Dummy9  T.Gpointer
	Dummy10 T.Gpointer
	Dummy11 int
	Dummy12 int
	Dummy13 int
	Dummy14 T.Gpointer
}

var (
	TextIterGetType func() T.GType

	textIterBackwardChar                   func(t *TextIter) T.Gboolean
	textIterBackwardChars                  func(t *TextIter, count int) T.Gboolean
	textIterBackwardCursorPosition         func(t *TextIter) T.Gboolean
	textIterBackwardCursorPositions        func(t *TextIter, count int) T.Gboolean
	textIterBackwardFindChar               func(t *TextIter, pred T.GtkTextCharPredicate, userData T.Gpointer, limit *TextIter) T.Gboolean
	textIterBackwardLine                   func(t *TextIter) T.Gboolean
	textIterBackwardLines                  func(t *TextIter, count int) T.Gboolean
	textIterBackwardSearch                 func(t *TextIter, str string, flags T.GtkTextSearchFlags, matchStart, matchEnd, limit *TextIter) T.Gboolean
	textIterBackwardSentenceStart          func(t *TextIter) T.Gboolean
	textIterBackwardSentenceStarts         func(t *TextIter, count int) T.Gboolean
	textIterBackwardToTagToggle            func(t *TextIter, tag *TextTag) T.Gboolean
	textIterBackwardVisibleCursorPosition  func(t *TextIter) T.Gboolean
	textIterBackwardVisibleCursorPositions func(t *TextIter, count int) T.Gboolean
	textIterBackwardVisibleLine            func(t *TextIter) T.Gboolean
	textIterBackwardVisibleLines           func(t *TextIter, count int) T.Gboolean
	textIterBackwardVisibleWordStart       func(t *TextIter) T.Gboolean
	textIterBackwardVisibleWordStarts      func(t *TextIter, count int) T.Gboolean
	textIterBackwardWordStart              func(t *TextIter) T.Gboolean
	textIterBackwardWordStarts             func(t *TextIter, count int) T.Gboolean
	textIterBeginsTag                      func(t *TextIter, tag *TextTag) T.Gboolean
	textIterCanInsert                      func(t *TextIter, defaultEditability T.Gboolean) T.Gboolean
	textIterCompare                        func(t *TextIter, rhs *TextIter) int
	textIterCopy                           func(t *TextIter) *TextIter
	textIterEditable                       func(t *TextIter, defaultSetting T.Gboolean) T.Gboolean
	textIterEndsLine                       func(t *TextIter) T.Gboolean
	textIterEndsSentence                   func(t *TextIter) T.Gboolean
	textIterEndsTag                        func(t *TextIter, tag *TextTag) T.Gboolean
	textIterEndsWord                       func(t *TextIter) T.Gboolean
	textIterEqual                          func(t *TextIter, rhs *TextIter) T.Gboolean
	textIterForwardChar                    func(t *TextIter) T.Gboolean
	textIterForwardChars                   func(t *TextIter, count int) T.Gboolean
	textIterForwardCursorPosition          func(t *TextIter) T.Gboolean
	textIterForwardCursorPositions         func(t *TextIter, count int) T.Gboolean
	textIterForwardFindChar                func(t *TextIter, pred T.GtkTextCharPredicate, userData T.Gpointer, limit *TextIter) T.Gboolean
	textIterForwardLine                    func(t *TextIter) T.Gboolean
	textIterForwardLines                   func(t *TextIter, count int) T.Gboolean
	textIterForwardSearch                  func(t *TextIter, str string, flags T.GtkTextSearchFlags, matchStart, matchEnd, limit *TextIter) T.Gboolean
	textIterForwardSentenceEnd             func(t *TextIter) T.Gboolean
	textIterForwardSentenceEnds            func(t *TextIter, count int) T.Gboolean
	textIterForwardToEnd                   func(t *TextIter)
	textIterForwardToLineEnd               func(t *TextIter) T.Gboolean
	textIterForwardToTagToggle             func(t *TextIter, tag *TextTag) T.Gboolean
	textIterForwardVisibleCursorPosition   func(t *TextIter) T.Gboolean
	textIterForwardVisibleCursorPositions  func(t *TextIter, count int) T.Gboolean
	textIterForwardVisibleLine             func(t *TextIter) T.Gboolean
	textIterForwardVisibleLines            func(t *TextIter, count int) T.Gboolean
	textIterForwardVisibleWordEnd          func(t *TextIter) T.Gboolean
	textIterForwardVisibleWordEnds         func(t *TextIter, count int) T.Gboolean
	textIterForwardWordEnd                 func(t *TextIter) T.Gboolean
	textIterForwardWordEnds                func(t *TextIter, count int) T.Gboolean
	textIterFree                           func(t *TextIter)
	textIterGetAttributes                  func(t *TextIter, values *TextAttributes) T.Gboolean
	textIterGetBuffer                      func(t *TextIter) *TextBuffer
	textIterGetBytesInLine                 func(t *TextIter) int
	textIterGetChar                        func(t *TextIter) T.Gunichar
	textIterGetCharsInLine                 func(t *TextIter) int
	textIterGetChildAnchor                 func(t *TextIter) *TextChildAnchor
	textIterGetLanguage                    func(t *TextIter) *T.PangoLanguage
	textIterGetLine                        func(t *TextIter) int
	textIterGetLineIndex                   func(t *TextIter) int
	textIterGetLineOffset                  func(t *TextIter) int
	textIterGetMarks                       func(t *TextIter) *T.GSList
	textIterGetOffset                      func(t *TextIter) int
	textIterGetPixbuf                      func(t *TextIter) *T.GdkPixbuf
	textIterGetSlice                       func(t *TextIter, end *TextIter) string
	textIterGetTags                        func(t *TextIter) *T.GSList
	textIterGetText                        func(t *TextIter, end *TextIter) string
	textIterGetToggledTags                 func(t *TextIter, toggledOn T.Gboolean) *T.GSList
	textIterGetVisibleLineIndex            func(t *TextIter) int
	textIterGetVisibleLineOffset           func(t *TextIter) int
	textIterGetVisibleSlice                func(t *TextIter, end *TextIter) string
	textIterGetVisibleText                 func(t *TextIter, end *TextIter) string
	textIterHasTag                         func(t *TextIter, tag *TextTag) T.Gboolean
	textIterInRange                        func(t *TextIter, start, end *TextIter) T.Gboolean
	textIterInsideSentence                 func(t *TextIter) T.Gboolean
	textIterInsideWord                     func(t *TextIter) T.Gboolean
	textIterIsCursorPosition               func(t *TextIter) T.Gboolean
	textIterIsEnd                          func(t *TextIter) T.Gboolean
	textIterIsStart                        func(t *TextIter) T.Gboolean
	textIterOrder                          func(t *TextIter, second *TextIter)
	textIterSetLine                        func(t *TextIter, lineNumber int)
	textIterSetLineIndex                   func(t *TextIter, byteOnLine int)
	textIterSetLineOffset                  func(t *TextIter, charOnLine int)
	textIterSetOffset                      func(t *TextIter, charOffset int)
	textIterSetVisibleLineIndex            func(t *TextIter, byteOnLine int)
	textIterSetVisibleLineOffset           func(t *TextIter, charOnLine int)
	textIterStartsLine                     func(t *TextIter) T.Gboolean
	textIterStartsSentence                 func(t *TextIter) T.Gboolean
	textIterStartsWord                     func(t *TextIter) T.Gboolean
	textIterTogglesTag                     func(t *TextIter, tag *TextTag) T.Gboolean
)

func (t *TextIter) BackwardChar() T.Gboolean           { return textIterBackwardChar(t) }
func (t *TextIter) BackwardChars(count int) T.Gboolean { return textIterBackwardChars(t, count) }
func (t *TextIter) BackwardCursorPosition() T.Gboolean { return textIterBackwardCursorPosition(t) }
func (t *TextIter) BackwardCursorPositions(count int) T.Gboolean {
	return textIterBackwardCursorPositions(t, count)
}
func (t *TextIter) BackwardFindChar(pred T.GtkTextCharPredicate, userData T.Gpointer, limit *TextIter) T.Gboolean {
	return textIterBackwardFindChar(t, pred, userData, limit)
}
func (t *TextIter) BackwardLine() T.Gboolean           { return textIterBackwardLine(t) }
func (t *TextIter) BackwardLines(count int) T.Gboolean { return textIterBackwardLines(t, count) }
func (t *TextIter) BackwardSearch(str string, flags T.GtkTextSearchFlags, matchStart, matchEnd, limit *TextIter) T.Gboolean {
	return textIterBackwardSearch(t, str, flags, matchStart, matchEnd, limit)
}
func (t *TextIter) BackwardSentenceStart() T.Gboolean { return textIterBackwardSentenceStart(t) }
func (t *TextIter) BackwardSentenceStarts(count int) T.Gboolean {
	return textIterBackwardSentenceStarts(t, count)
}
func (t *TextIter) BackwardToTagToggle(tag *TextTag) T.Gboolean {
	return textIterBackwardToTagToggle(t, tag)
}
func (t *TextIter) BackwardVisibleCursorPosition() T.Gboolean {
	return textIterBackwardVisibleCursorPosition(t)
}
func (t *TextIter) BackwardVisibleCursorPositions(count int) T.Gboolean {
	return textIterBackwardVisibleCursorPositions(t, count)
}
func (t *TextIter) BackwardVisibleLine() T.Gboolean { return textIterBackwardVisibleLine(t) }
func (t *TextIter) BackwardVisibleLines(count int) T.Gboolean {
	return textIterBackwardVisibleLines(t, count)
}
func (t *TextIter) BackwardVisibleWordStart() T.Gboolean { return textIterBackwardVisibleWordStart(t) }
func (t *TextIter) BackwardVisibleWordStarts(count int) T.Gboolean {
	return textIterBackwardVisibleWordStarts(t, count)
}
func (t *TextIter) BackwardWordStart() T.Gboolean { return textIterBackwardWordStart(t) }
func (t *TextIter) BackwardWordStarts(count int) T.Gboolean {
	return textIterBackwardWordStarts(t, count)
}
func (t *TextIter) BeginsTag(tag *TextTag) T.Gboolean { return textIterBeginsTag(t, tag) }
func (t *TextIter) CanInsert(defaultEditability T.Gboolean) T.Gboolean {
	return textIterCanInsert(t, defaultEditability)
}
func (t *TextIter) Compare(rhs *TextIter) int { return textIterCompare(t, rhs) }
func (t *TextIter) Copy() *TextIter           { return textIterCopy(t) }
func (t *TextIter) Editable(defaultSetting T.Gboolean) T.Gboolean {
	return textIterEditable(t, defaultSetting)
}
func (t *TextIter) EndsLine() T.Gboolean              { return textIterEndsLine(t) }
func (t *TextIter) EndsSentence() T.Gboolean          { return textIterEndsSentence(t) }
func (t *TextIter) EndsTag(tag *TextTag) T.Gboolean   { return textIterEndsTag(t, tag) }
func (t *TextIter) EndsWord() T.Gboolean              { return textIterEndsWord(t) }
func (t *TextIter) Equal(rhs *TextIter) T.Gboolean    { return textIterEqual(t, rhs) }
func (t *TextIter) ForwardChar() T.Gboolean           { return textIterForwardChar(t) }
func (t *TextIter) ForwardChars(count int) T.Gboolean { return textIterForwardChars(t, count) }
func (t *TextIter) ForwardCursorPosition() T.Gboolean { return textIterForwardCursorPosition(t) }
func (t *TextIter) ForwardCursorPositions(count int) T.Gboolean {
	return textIterForwardCursorPositions(t, count)
}
func (t *TextIter) ForwardFindChar(pred T.GtkTextCharPredicate, userData T.Gpointer, limit *TextIter) T.Gboolean {
	return textIterForwardFindChar(t, pred, userData, limit)
}
func (t *TextIter) ForwardLine() T.Gboolean           { return textIterForwardLine(t) }
func (t *TextIter) ForwardLines(count int) T.Gboolean { return textIterForwardLines(t, count) }
func (t *TextIter) ForwardSearch(str string, flags T.GtkTextSearchFlags, matchStart, matchEnd, limit *TextIter) T.Gboolean {
	return textIterForwardSearch(t, str, flags, matchStart, matchEnd, limit)
}
func (t *TextIter) ForwardSentenceEnd() T.Gboolean { return textIterForwardSentenceEnd(t) }
func (t *TextIter) ForwardSentenceEnds(count int) T.Gboolean {
	return textIterForwardSentenceEnds(t, count)
}
func (t *TextIter) ForwardToEnd()                { textIterForwardToEnd(t) }
func (t *TextIter) ForwardToLineEnd() T.Gboolean { return textIterForwardToLineEnd(t) }
func (t *TextIter) ForwardToTagToggle(tag *TextTag) T.Gboolean {
	return textIterForwardToTagToggle(t, tag)
}
func (t *TextIter) ForwardVisibleCursorPosition() T.Gboolean {
	return textIterForwardVisibleCursorPosition(t)
}
func (t *TextIter) ForwardVisibleCursorPositions(count int) T.Gboolean {
	return textIterForwardVisibleCursorPositions(t, count)
}
func (t *TextIter) ForwardVisibleLine() T.Gboolean { return textIterForwardVisibleLine(t) }
func (t *TextIter) ForwardVisibleLines(count int) T.Gboolean {
	return textIterForwardVisibleLines(t, count)
}
func (t *TextIter) ForwardVisibleWordEnd() T.Gboolean { return textIterForwardVisibleWordEnd(t) }
func (t *TextIter) ForwardVisibleWordEnds(count int) T.Gboolean {
	return textIterForwardVisibleWordEnds(t, count)
}
func (t *TextIter) ForwardWordEnd() T.Gboolean           { return textIterForwardWordEnd(t) }
func (t *TextIter) ForwardWordEnds(count int) T.Gboolean { return textIterForwardWordEnds(t, count) }
func (t *TextIter) Free()                                { textIterFree(t) }
func (t *TextIter) GetAttributes(values *TextAttributes) T.Gboolean {
	return textIterGetAttributes(t, values)
}
func (t *TextIter) GetBuffer() *TextBuffer           { return textIterGetBuffer(t) }
func (t *TextIter) GetBytesInLine() int              { return textIterGetBytesInLine(t) }
func (t *TextIter) GetChar() T.Gunichar              { return textIterGetChar(t) }
func (t *TextIter) GetCharsInLine() int              { return textIterGetCharsInLine(t) }
func (t *TextIter) GetChildAnchor() *TextChildAnchor { return textIterGetChildAnchor(t) }
func (t *TextIter) GetLanguage() *T.PangoLanguage    { return textIterGetLanguage(t) }
func (t *TextIter) GetLine() int                     { return textIterGetLine(t) }
func (t *TextIter) GetLineIndex() int                { return textIterGetLineIndex(t) }
func (t *TextIter) GetLineOffset() int               { return textIterGetLineOffset(t) }
func (t *TextIter) GetMarks() *T.GSList              { return textIterGetMarks(t) }
func (t *TextIter) GetOffset() int                   { return textIterGetOffset(t) }
func (t *TextIter) GetPixbuf() *T.GdkPixbuf          { return textIterGetPixbuf(t) }
func (t *TextIter) GetSlice(end *TextIter) string    { return textIterGetSlice(t, end) }
func (t *TextIter) GetTags() *T.GSList               { return textIterGetTags(t) }
func (t *TextIter) GetText(end *TextIter) string     { return textIterGetText(t, end) }
func (t *TextIter) GetToggledTags(toggledOn T.Gboolean) *T.GSList {
	return textIterGetToggledTags(t, toggledOn)
}
func (t *TextIter) GetVisibleLineIndex() int                { return textIterGetVisibleLineIndex(t) }
func (t *TextIter) GetVisibleLineOffset() int               { return textIterGetVisibleLineOffset(t) }
func (t *TextIter) GetVisibleSlice(end *TextIter) string    { return textIterGetVisibleSlice(t, end) }
func (t *TextIter) GetVisibleText(end *TextIter) string     { return textIterGetVisibleText(t, end) }
func (t *TextIter) HasTag(tag *TextTag) T.Gboolean          { return textIterHasTag(t, tag) }
func (t *TextIter) InRange(start, end *TextIter) T.Gboolean { return textIterInRange(t, start, end) }
func (t *TextIter) InsideSentence() T.Gboolean              { return textIterInsideSentence(t) }
func (t *TextIter) InsideWord() T.Gboolean                  { return textIterInsideWord(t) }
func (t *TextIter) IsCursorPosition() T.Gboolean            { return textIterIsCursorPosition(t) }
func (t *TextIter) IsEnd() T.Gboolean                       { return textIterIsEnd(t) }
func (t *TextIter) IsStart() T.Gboolean                     { return textIterIsStart(t) }
func (t *TextIter) Order(second *TextIter)                  { textIterOrder(t, second) }
func (t *TextIter) SetLine(lineNumber int)                  { textIterSetLine(t, lineNumber) }
func (t *TextIter) SetLineIndex(byteOnLine int)             { textIterSetLineIndex(t, byteOnLine) }
func (t *TextIter) SetLineOffset(charOnLine int)            { textIterSetLineOffset(t, charOnLine) }
func (t *TextIter) SetOffset(charOffset int)                { textIterSetOffset(t, charOffset) }
func (t *TextIter) SetVisibleLineIndex(byteOnLine int)      { textIterSetVisibleLineIndex(t, byteOnLine) }
func (t *TextIter) SetVisibleLineOffset(charOnLine int)     { textIterSetVisibleLineOffset(t, charOnLine) }
func (t *TextIter) StartsLine() T.Gboolean                  { return textIterStartsLine(t) }
func (t *TextIter) StartsSentence() T.Gboolean              { return textIterStartsSentence(t) }
func (t *TextIter) StartsWord() T.Gboolean                  { return textIterStartsWord(t) }
func (t *TextIter) TogglesTag(tag *TextTag) T.Gboolean      { return textIterTogglesTag(t, tag) }

type TextMark struct {
	Parent  T.GObject
	Segment T.Gpointer
}

var (
	TextMarkGetType func() T.GType
	TextMarkNew     func(name string, leftGravity T.Gboolean) *TextMark

	TextMarkSetVisible     func(t *TextMark, setting T.Gboolean)
	TextMarkGetVisible     func(t *TextMark) T.Gboolean
	TextMarkGetName        func(t *TextMark) string
	TextMarkGetDeleted     func(t *TextMark) T.Gboolean
	TextMarkGetBuffer      func(t *TextMark) *TextBuffer
	TextMarkGetLeftGravity func(t *TextMark) T.Gboolean
)

type TextView struct {
	Parent               T.GtkContainer
	Layout               *T.GtkTextLayout
	Buffer               *TextBuffer
	SelectionDragHandler uint
	ScrollTimeout        uint
	PixelsAboveLines     int
	PixelsBelowLines     int
	PixelsInsideWrap     int
	WrapMode             T.GtkWrapMode
	Justify              T.GtkJustification
	LeftMargin           int
	RightMargin          int
	Indent               int
	Tabs                 *T.PangoTabArray
	Bits                 uint
	// Editable : 1
	// OverwriteMode : 1
	// CursorVisible : 1
	// NeedImReset : 1
	// AcceptsTab : 1
	// WidthChanged : 1
	// OnscreenValidated : 1
	// MouseCursorObscured : 1
	TextWindow               *T.GtkTextWindow
	LeftWindow               *T.GtkTextWindow
	RightWindow              *T.GtkTextWindow
	TopWindow                *T.GtkTextWindow
	BottomWindow             *T.GtkTextWindow
	Hadjustment              *T.GtkAdjustment
	Vadjustment              *T.GtkAdjustment
	Xoffset                  int
	Yoffset                  int
	Width                    int
	Height                   int
	VirtualCursorX           int
	VirtualCursorY           int
	FirstParaMark            *TextMark
	FirstParaPixels          int
	DndMark                  *TextMark
	BlinkTimeout             uint
	FirstValidateIdle        uint
	IncrementalValidateIdle  uint
	ImContext                *IMContext
	PopupMenu                *Widget
	DragStartX               int
	DragStartY               int
	Children                 *T.GSList
	PendingScroll            *T.GtkTextPendingScroll
	PendingPlaceCursorButton int
}

var (
	TextViewGetType       func() T.GType
	TextViewNew           func() *Widget
	TextViewNewWithBuffer func(buffer *TextBuffer) *Widget

	TextViewSetBuffer                func(t *TextView, buffer *TextBuffer)
	TextViewGetBuffer                func(t *TextView) *TextBuffer
	TextViewScrollToIter             func(t *TextView, iter *TextIter, withinMargin float64, useAlign T.Gboolean, xalign, yalign float64) T.Gboolean
	TextViewScrollToMark             func(t *TextView, mark *TextMark, withinMargin float64, useAlign T.Gboolean, xalign, yalign float64)
	TextViewScrollMarkOnscreen       func(t *TextView, mark *TextMark)
	TextViewMoveMarkOnscreen         func(t *TextView, mark *TextMark) T.Gboolean
	TextViewPlaceCursorOnscreen      func(t *TextView) T.Gboolean
	TextViewGetVisibleRect           func(t *TextView, visibleRect *T.GdkRectangle)
	TextViewSetCursorVisible         func(t *TextView, setting T.Gboolean)
	TextViewGetCursorVisible         func(t *TextView) T.Gboolean
	TextViewGetIterLocation          func(t *TextView, iter *TextIter, location *T.GdkRectangle)
	TextViewGetIterAtLocation        func(t *TextView, iter *TextIter, x, y int)
	TextViewGetIterAtPosition        func(t *TextView, iter *TextIter, trailing *int, x, y int)
	TextViewGetLineYrange            func(t *TextView, iter *TextIter, y, height *int)
	TextViewGetLineAtY               func(t *TextView, targetIter *TextIter, y int, lineTop *int)
	TextViewBufferToWindowCoords     func(t *TextView, win T.GtkTextWindowType, bufferX, bufferY int, windowX, windowY *int)
	TextViewWindowToBufferCoords     func(t *TextView, win T.GtkTextWindowType, windowX, windowY int, bufferX, bufferY *int)
	TextViewGetHadjustment           func(t *TextView) *Adjustment
	TextViewGetVadjustment           func(t *TextView) *Adjustment
	TextViewGetWindow                func(t *TextView, win T.GtkTextWindowType) *T.GdkWindow
	TextViewGetWindowType            func(t *TextView, window *T.GdkWindow) T.GtkTextWindowType
	TextViewSetBorderWindowSize      func(t *TextView, tw T.GtkTextWindowType, size int)
	TextViewGetBorderWindowSize      func(t *TextView, tw T.GtkTextWindowType) int
	TextViewForwardDisplayLine       func(t *TextView, iter *TextIter) T.Gboolean
	TextViewBackwardDisplayLine      func(t *TextView, iter *TextIter) T.Gboolean
	TextViewForwardDisplayLineEnd    func(t *TextView, iter *TextIter) T.Gboolean
	TextViewBackwardDisplayLineStart func(t *TextView, iter *TextIter) T.Gboolean
	TextViewStartsDisplayLine        func(t *TextView, iter *TextIter) T.Gboolean
	TextViewMoveVisually             func(t *TextView, iter *TextIter, count int) T.Gboolean
	TextViewImContextFilterKeypress  func(t *TextView, event *T.GdkEventKey) T.Gboolean
	TextViewResetImContext           func(t *TextView)
	TextViewAddChildAtAnchor         func(t *TextView, child *Widget, anchor *TextChildAnchor)
	TextViewAddChildInWindow         func(t *TextView, child *Widget, whichWindow T.GtkTextWindowType, xpos, ypos int)
	TextViewMoveChild                func(t *TextView, child *Widget, xpos, ypos int)
	TextViewSetWrapMode              func(t *TextView, wrapMode T.GtkWrapMode)
	TextViewGetWrapMode              func(t *TextView) T.GtkWrapMode
	TextViewSetEditable              func(t *TextView, setting T.Gboolean)
	TextViewGetEditable              func(t *TextView) T.Gboolean
	TextViewSetOverwrite             func(t *TextView, overwrite T.Gboolean)
	TextViewGetOverwrite             func(t *TextView) T.Gboolean
	TextViewSetAcceptsTab            func(t *TextView, acceptsTab T.Gboolean)
	TextViewGetAcceptsTab            func(t *TextView) T.Gboolean
	TextViewSetPixelsAboveLines      func(t *TextView, pixelsAboveLines int)
	TextViewGetPixelsAboveLines      func(t *TextView) int
	TextViewSetPixelsBelowLines      func(t *TextView, pixelsBelowLines int)
	TextViewGetPixelsBelowLines      func(t *TextView) int
	TextViewSetPixelsInsideWrap      func(t *TextView, pixelsInsideWrap int)
	TextViewGetPixelsInsideWrap      func(t *TextView) int
	TextViewSetJustification         func(t *TextView, justification T.GtkJustification)
	TextViewGetJustification         func(t *TextView) T.GtkJustification
	TextViewSetLeftMargin            func(t *TextView, leftMargin int)
	TextViewGetLeftMargin            func(t *TextView) int
	TextViewSetRightMargin           func(t *TextView, rightMargin int)
	TextViewGetRightMargin           func(t *TextView) int
	TextViewSetIndent                func(t *TextView, indent int)
	TextViewGetIndent                func(t *TextView) int
	TextViewSetTabs                  func(t *TextView, tabs *T.PangoTabArray)
	TextViewGetTabs                  func(t *TextView) *T.PangoTabArray
	TextViewGetDefaultAttributes     func(t *TextView) *TextAttributes
)

type TextTag struct {
	Parent   T.GObject
	Table    *TextTagTable
	Name     *T.Char
	Priority int
	Values   *TextAttributes
	Bits     uint
	// BgColorSet : 1
	// BgStippleSet : 1
	// FgColorSet : 1
	// ScaleSet : 1
	// FgStippleSet : 1
	// JustificationSet : 1
	// LeftMarginSet : 1
	// IndentSet : 1
	// RiseSet : 1
	// StrikethroughSet : 1
	// RightMarginSet : 1
	// PixelsAboveLinesSet : 1
	// PixelsBelowLinesSet : 1
	// PixelsInsideWrapSet : 1
	// TabsSet : 1
	// UnderlineSet : 1
	// WrapModeSet : 1
	// BgFullHeightSet : 1
	// InvisibleSet : 1
	// EditableSet : 1
	// LanguageSet : 1
	// PgBgColorSet : 1
	// AccumulativeMargin : 1
	// Pad1 : 1
}

var (
	TextTagGetType func() T.GType
	TextTagNew     func(name string) *TextTag

	textTagEvent       func(t *TextTag, eventObject *T.GObject, event *T.GdkEvent, iter *TextIter) T.Gboolean
	textTagGetPriority func(t *TextTag) int
	textTagSetPriority func(t *TextTag, priority int)
)

func (t *TextTag) Event(eventObject *T.GObject, event *T.GdkEvent, iter *TextIter) T.Gboolean {
	return textTagEvent(t, eventObject, event, iter)
}
func (t *TextTag) GetPriority() int         { return textTagGetPriority(t) }
func (t *TextTag) SetPriority(priority int) { textTagSetPriority(t, priority) }

type (
	TextTagTable struct {
		Parent    T.GObject
		Hash      *T.GHashTable
		Anonymous *T.GSList
		AnonCount int
		Buffers   *T.GSList
	}

	TextTagTableForeach func(tag *TextTag, data T.Gpointer)
)

var (
	TextTagTableGetType func() T.GType
	TextTagTableNew     func() *TextTagTable

	textTagTableAdd     func(t *TextTagTable, tag *TextTag)
	textTagTableForeach func(t *TextTagTable, f TextTagTableForeach, data T.Gpointer)
	textTagTableGetSize func(t *TextTagTable) int
	textTagTableLookup  func(t *TextTagTable, name string) *TextTag
	textTagTableRemove  func(t *TextTagTable, tag *TextTag)
)

func (t *TextTagTable) Add(tag *TextTag) { textTagTableAdd(t, tag) }
func (t *TextTagTable) Foreach(f TextTagTableForeach, data T.Gpointer) {
	textTagTableForeach(t, f, data)
}
func (t *TextTagTable) GetSize() int                { return textTagTableGetSize(t) }
func (t *TextTagTable) Lookup(name string) *TextTag { return textTagTableLookup(t, name) }
func (t *TextTagTable) Remove(tag *TextTag)         { textTagTableRemove(t, tag) }

type ToolbarStyle T.Enum

const (
	TOOLBAR_ICONS ToolbarStyle = iota
	TOOLBAR_TEXT
	TOOLBAR_BOTH
	TOOLBAR_BOTH_HORIZ
)

type ToolbarChildType T.Enum

const (
	TOOLBAR_CHILD_SPACE ToolbarChildType = iota
	TOOLBAR_CHILD_BUTTON
	TOOLBAR_CHILD_TOGGLEBUTTON
	TOOLBAR_CHILD_RADIOBUTTON
	TOOLBAR_CHILD_WIDGET
)

var (
	ToolbarGetType           func() T.GType
	ToolbarNew               func() *Widget
	ToolbarChildTypeGetType  func() T.GType
	ToolbarSpaceStyleGetType func() T.GType
	ToolbarStyleGetType      func() T.GType

	toolbarAppendElement        func(t *Toolbar, ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget
	toolbarAppendItem           func(t *Toolbar, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget
	toolbarAppendSpace          func(t *Toolbar)
	toolbarAppendWidget         func(t *Toolbar, widget *Widget, tooltipText, tooltipPrivateText string)
	toolbarGetDropIndex         func(t *Toolbar, x, y int) int
	toolbarGetIconSize          func(t *Toolbar) IconSize
	toolbarGetItemIndex         func(t *Toolbar, item *ToolItem) int
	toolbarGetNItems            func(t *Toolbar) int
	toolbarGetNthItem           func(t *Toolbar, n int) *ToolItem
	toolbarGetOrientation       func(t *Toolbar) T.GtkOrientation
	toolbarGetReliefStyle       func(t *Toolbar) T.GtkReliefStyle
	toolbarGetShowArrow         func(t *Toolbar) T.Gboolean
	toolbarGetStyle             func(t *Toolbar) ToolbarStyle
	toolbarGetTooltips          func(t *Toolbar) T.Gboolean
	toolbarInsert               func(t *Toolbar, item *ToolItem, pos int)
	toolbarInsertElement        func(t *Toolbar, ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer, position int) *Widget
	toolbarInsertItem           func(t *Toolbar, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer, position int) *Widget
	toolbarInsertSpace          func(t *Toolbar, position int)
	toolbarInsertStock          func(t *Toolbar, stockId, tooltipText, tooltipPrivateText string, callback T.GCallback, userData T.Gpointer, position int) *Widget
	toolbarInsertWidget         func(t *Toolbar, widget *Widget, tooltipText, tooltipPrivateText string, position int)
	toolbarPrependElement       func(t *Toolbar, ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget
	toolbarPrependItem          func(t *Toolbar, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget
	toolbarPrependSpace         func(t *Toolbar)
	toolbarPrependWidget        func(t *Toolbar, widget *Widget, tooltipText, tooltipPrivateText string)
	toolbarRemoveSpace          func(t *Toolbar, position int)
	toolbarSetDropHighlightItem func(t *Toolbar, toolItem *ToolItem, index int)
	toolbarSetIconSize          func(t *Toolbar, iconSize IconSize)
	toolbarSetOrientation       func(t *Toolbar, orientation T.GtkOrientation)
	toolbarSetShowArrow         func(t *Toolbar, showArrow T.Gboolean)
	toolbarSetStyle             func(t *Toolbar, style ToolbarStyle)
	toolbarSetTooltips          func(t *Toolbar, enable T.Gboolean)
	toolbarUnsetIconSize        func(t *Toolbar)
	toolbarUnsetStyle           func(t *Toolbar)
)

func (t *Toolbar) AppendElement(ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget {
	return toolbarAppendElement(t, ct, widget, text, tooltipText, tooltipPrivateText, icon, callback, userData)
}
func (t *Toolbar) AppendItem(text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget {
	return toolbarAppendItem(t, text, tooltipText, tooltipPrivateText, icon, callback, userData)
}
func (t *Toolbar) AppendSpace() { toolbarAppendSpace(t) }
func (t *Toolbar) AppendWidget(widget *Widget, tooltipText, tooltipPrivateText string) {
	toolbarAppendWidget(t, widget, tooltipText, tooltipPrivateText)
}
func (t *Toolbar) GetDropIndex(x, y int) int        { return toolbarGetDropIndex(t, x, y) }
func (t *Toolbar) GetIconSize() IconSize            { return toolbarGetIconSize(t) }
func (t *Toolbar) GetItemIndex(item *ToolItem) int  { return toolbarGetItemIndex(t, item) }
func (t *Toolbar) GetNItems() int                   { return toolbarGetNItems(t) }
func (t *Toolbar) GetNthItem(n int) *ToolItem       { return toolbarGetNthItem(t, n) }
func (t *Toolbar) GetOrientation() T.GtkOrientation { return toolbarGetOrientation(t) }
func (t *Toolbar) GetReliefStyle() T.GtkReliefStyle { return toolbarGetReliefStyle(t) }
func (t *Toolbar) GetShowArrow() T.Gboolean         { return toolbarGetShowArrow(t) }
func (t *Toolbar) GetStyle() ToolbarStyle           { return toolbarGetStyle(t) }
func (t *Toolbar) GetTooltips() T.Gboolean          { return toolbarGetTooltips(t) }
func (t *Toolbar) Insert(item *ToolItem, pos int)   { toolbarInsert(t, item, pos) }
func (t *Toolbar) InsertElement(ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer, position int) *Widget {
	return toolbarInsertElement(t, ct, widget, text, tooltipText, tooltipPrivateText, icon, callback, userData, position)
}
func (t *Toolbar) InsertItem(text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer, position int) *Widget {
	return toolbarInsertItem(t, text, tooltipText, tooltipPrivateText, icon, callback, userData, position)
}
func (t *Toolbar) InsertSpace(position int) { toolbarInsertSpace(t, position) }
func (t *Toolbar) InsertStock(stockId, tooltipText, tooltipPrivateText string, callback T.GCallback, userData T.Gpointer, position int) *Widget {
	return toolbarInsertStock(t, stockId, tooltipText, tooltipPrivateText, callback, userData, position)
}
func (t *Toolbar) InsertWidget(widget *Widget, tooltipText, tooltipPrivateText string, position int) {
	toolbarInsertWidget(t, widget, tooltipText, tooltipPrivateText, position)
}
func (t *Toolbar) PrependElement(ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget {
	return toolbarPrependElement(t, ct, widget, text, tooltipText, tooltipPrivateText, icon, callback, userData)
}
func (t *Toolbar) PrependItem(text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget {
	return toolbarPrependItem(t, text, tooltipText, tooltipPrivateText, icon, callback, userData)
}
func (t *Toolbar) PrependSpace() { toolbarPrependSpace(t) }
func (t *Toolbar) PrependWidget(widget *Widget, tooltipText, tooltipPrivateText string) {
	toolbarPrependWidget(t, widget, tooltipText, tooltipPrivateText)
}
func (t *Toolbar) RemoveSpace(position int) { toolbarRemoveSpace(t, position) }
func (t *Toolbar) SetDropHighlightItem(toolItem *ToolItem, index int) {
	toolbarSetDropHighlightItem(t, toolItem, index)
}
func (t *Toolbar) SetIconSize(iconSize IconSize)               { toolbarSetIconSize(t, iconSize) }
func (t *Toolbar) SetOrientation(orientation T.GtkOrientation) { toolbarSetOrientation(t, orientation) }
func (t *Toolbar) SetShowArrow(showArrow T.Gboolean)           { toolbarSetShowArrow(t, showArrow) }
func (t *Toolbar) SetStyle(style ToolbarStyle)                 { toolbarSetStyle(t, style) }
func (t *Toolbar) SetTooltips(enable T.Gboolean)               { toolbarSetTooltips(t, enable) }
func (t *Toolbar) UnsetIconSize()                              { toolbarUnsetIconSize(t) }
func (t *Toolbar) UnsetStyle()                                 { toolbarUnsetStyle(t) }

type ToolButton struct {
	Parent ToolItem
	_      *struct{}
}

var (
	ToolButtonGetType      func() T.GType
	ToolButtonNew          func(iconWidget *Widget, label string) *ToolItem
	ToolButtonNewFromStock func(stockId string) *ToolItem

	toolButtonGetIconName     func(t *ToolButton) string
	toolButtonGetIconWidget   func(t *ToolButton) *Widget
	toolButtonGetLabel        func(t *ToolButton) string
	toolButtonGetLabelWidget  func(t *ToolButton) *Widget
	toolButtonGetStockId      func(t *ToolButton) string
	toolButtonGetUseUnderline func(t *ToolButton) T.Gboolean
	toolButtonSetIconName     func(t *ToolButton, iconName string)
	toolButtonSetIconWidget   func(t *ToolButton, iconWidget *Widget)
	toolButtonSetLabel        func(t *ToolButton, label string)
	toolButtonSetLabelWidget  func(t *ToolButton, labelWidget *Widget)
	toolButtonSetStockId      func(t *ToolButton, stockId string)
	toolButtonSetUseUnderline func(t *ToolButton, useUnderline T.Gboolean)
)

func (t *ToolButton) GetIconName() string                { return toolButtonGetIconName(t) }
func (t *ToolButton) GetIconWidget() *Widget             { return toolButtonGetIconWidget(t) }
func (t *ToolButton) GetLabel() string                   { return toolButtonGetLabel(t) }
func (t *ToolButton) GetLabelWidget() *Widget            { return toolButtonGetLabelWidget(t) }
func (t *ToolButton) GetStockId() string                 { return toolButtonGetStockId(t) }
func (t *ToolButton) GetUseUnderline() T.Gboolean        { return toolButtonGetUseUnderline(t) }
func (t *ToolButton) SetIconName(iconName string)        { toolButtonSetIconName(t, iconName) }
func (t *ToolButton) SetIconWidget(iconWidget *Widget)   { toolButtonSetIconWidget(t, iconWidget) }
func (t *ToolButton) SetLabel(label string)              { toolButtonSetLabel(t, label) }
func (t *ToolButton) SetLabelWidget(labelWidget *Widget) { toolButtonSetLabelWidget(t, labelWidget) }
func (t *ToolButton) SetStockId(stockId string)          { toolButtonSetStockId(t, stockId) }
func (t *ToolButton) SetUseUnderline(useUnderline T.Gboolean) {
	toolButtonSetUseUnderline(t, useUnderline)
}

type ToolItem struct {
	Parent Bin
	_      *struct{}
}

var (
	ToolItemGetType func() T.GType
	ToolItemNew     func() *ToolItem

	toolItemGetEllipsizeMode      func(t *ToolItem) T.PangoEllipsizeMode
	toolItemGetExpand             func(t *ToolItem) T.Gboolean
	toolItemGetHomogeneous        func(t *ToolItem) T.Gboolean
	toolItemGetIconSize           func(t *ToolItem) IconSize
	toolItemGetIsImportant        func(t *ToolItem) T.Gboolean
	toolItemGetOrientation        func(t *ToolItem) T.GtkOrientation
	toolItemGetProxyMenuItem      func(t *ToolItem, menuItemId string) *Widget
	toolItemGetReliefStyle        func(t *ToolItem) T.GtkReliefStyle
	toolItemGetTextAlignment      func(t *ToolItem) float32
	toolItemGetTextOrientation    func(t *ToolItem) T.GtkOrientation
	toolItemGetTextSizeGroup      func(t *ToolItem) *SizeGroup
	toolItemGetToolbarStyle       func(t *ToolItem) ToolbarStyle
	toolItemGetUseDragWindow      func(t *ToolItem) T.Gboolean
	toolItemGetVisibleHorizontal  func(t *ToolItem) T.Gboolean
	toolItemGetVisibleVertical    func(t *ToolItem) T.Gboolean
	toolItemRebuildMenu           func(t *ToolItem)
	toolItemRetrieveProxyMenuItem func(t *ToolItem) *Widget
	toolItemSetExpand             func(t *ToolItem, expand T.Gboolean)
	toolItemSetHomogeneous        func(t *ToolItem, homogeneous T.Gboolean)
	toolItemSetIsImportant        func(t *ToolItem, isImportant T.Gboolean)
	toolItemSetProxyMenuItem      func(t *ToolItem, menuItemId string, menuItem *Widget)
	toolItemSetTooltip            func(t *ToolItem, tooltips *Tooltips, tipText, tipPrivate string)
	toolItemSetTooltipMarkup      func(t *ToolItem, markup string)
	toolItemSetTooltipText        func(t *ToolItem, text string)
	toolItemSetUseDragWindow      func(t *ToolItem, useDragWindow T.Gboolean)
	toolItemSetVisibleHorizontal  func(t *ToolItem, visibleHorizontal T.Gboolean)
	toolItemSetVisibleVertical    func(t *ToolItem, visibleVertical T.Gboolean)
	toolItemToolbarReconfigured   func(t *ToolItem)
)

func (t *ToolItem) GetEllipsizeMode() T.PangoEllipsizeMode { return toolItemGetEllipsizeMode(t) }
func (t *ToolItem) GetExpand() T.Gboolean                  { return toolItemGetExpand(t) }
func (t *ToolItem) GetHomogeneous() T.Gboolean             { return toolItemGetHomogeneous(t) }
func (t *ToolItem) GetIconSize() IconSize                  { return toolItemGetIconSize(t) }
func (t *ToolItem) GetIsImportant() T.Gboolean             { return toolItemGetIsImportant(t) }
func (t *ToolItem) GetOrientation() T.GtkOrientation       { return toolItemGetOrientation(t) }
func (t *ToolItem) GetProxyMenuItem(menuItemId string) *Widget {
	return toolItemGetProxyMenuItem(t, menuItemId)
}
func (t *ToolItem) GetReliefStyle() T.GtkReliefStyle      { return toolItemGetReliefStyle(t) }
func (t *ToolItem) GetTextAlignment() float32             { return toolItemGetTextAlignment(t) }
func (t *ToolItem) GetTextOrientation() T.GtkOrientation  { return toolItemGetTextOrientation(t) }
func (t *ToolItem) GetTextSizeGroup() *SizeGroup          { return toolItemGetTextSizeGroup(t) }
func (t *ToolItem) GetToolbarStyle() ToolbarStyle         { return toolItemGetToolbarStyle(t) }
func (t *ToolItem) GetUseDragWindow() T.Gboolean          { return toolItemGetUseDragWindow(t) }
func (t *ToolItem) GetVisibleHorizontal() T.Gboolean      { return toolItemGetVisibleHorizontal(t) }
func (t *ToolItem) GetVisibleVertical() T.Gboolean        { return toolItemGetVisibleVertical(t) }
func (t *ToolItem) RebuildMenu()                          { toolItemRebuildMenu(t) }
func (t *ToolItem) RetrieveProxyMenuItem() *Widget        { return toolItemRetrieveProxyMenuItem(t) }
func (t *ToolItem) SetExpand(expand T.Gboolean)           { toolItemSetExpand(t, expand) }
func (t *ToolItem) SetHomogeneous(homogeneous T.Gboolean) { toolItemSetHomogeneous(t, homogeneous) }
func (t *ToolItem) SetIsImportant(isImportant T.Gboolean) { toolItemSetIsImportant(t, isImportant) }
func (t *ToolItem) SetProxyMenuItem(menuItemId string, menuItem *Widget) {
	toolItemSetProxyMenuItem(t, menuItemId, menuItem)
}
func (t *ToolItem) SetTooltip(tooltips *Tooltips, tipText, tipPrivate string) {
	toolItemSetTooltip(t, tooltips, tipText, tipPrivate)
}
func (t *ToolItem) SetTooltipMarkup(markup string) { toolItemSetTooltipMarkup(t, markup) }
func (t *ToolItem) SetTooltipText(text string)     { toolItemSetTooltipText(t, text) }
func (t *ToolItem) SetUseDragWindow(useDragWindow T.Gboolean) {
	toolItemSetUseDragWindow(t, useDragWindow)
}
func (t *ToolItem) SetVisibleHorizontal(visibleHorizontal T.Gboolean) {
	toolItemSetVisibleHorizontal(t, visibleHorizontal)
}
func (t *ToolItem) SetVisibleVertical(visibleVertical T.Gboolean) {
	toolItemSetVisibleVertical(t, visibleVertical)
}
func (t *ToolItem) ToolbarReconfigured() { toolItemToolbarReconfigured(t) }

type ToolItemGroup struct {
	Parent T.GtkContainer
	_      *struct{}
}

var (
	ToolItemGroupGetType func() T.GType
	ToolItemGroupNew     func(label string) *Widget

	toolItemGroupGetCollapsed    func(t *ToolItemGroup) T.Gboolean
	toolItemGroupGetDropItem     func(t *ToolItemGroup, x, y int) *ToolItem
	toolItemGroupGetEllipsize    func(t *ToolItemGroup) T.PangoEllipsizeMode
	toolItemGroupGetHeaderRelief func(t *ToolItemGroup) T.GtkReliefStyle
	toolItemGroupGetItemPosition func(t *ToolItemGroup, item *ToolItem) int
	toolItemGroupGetLabel        func(t *ToolItemGroup) string
	toolItemGroupGetLabelWidget  func(t *ToolItemGroup) *Widget
	toolItemGroupGetNItems       func(t *ToolItemGroup) uint
	toolItemGroupGetNthItem      func(t *ToolItemGroup, index uint) *ToolItem
	toolItemGroupInsert          func(t *ToolItemGroup, item *ToolItem, position int)
	toolItemGroupSetCollapsed    func(t *ToolItemGroup, collapsed T.Gboolean)
	toolItemGroupSetEllipsize    func(t *ToolItemGroup, ellipsize T.PangoEllipsizeMode)
	toolItemGroupSetHeaderRelief func(t *ToolItemGroup, style T.GtkReliefStyle)
	toolItemGroupSetItemPosition func(t *ToolItemGroup, item *ToolItem, position int)
	toolItemGroupSetLabel        func(t *ToolItemGroup, label string)
	toolItemGroupSetLabelWidget  func(t *ToolItemGroup, labelWidget *Widget)
)

func (t *ToolItemGroup) GetCollapsed() T.Gboolean           { return toolItemGroupGetCollapsed(t) }
func (t *ToolItemGroup) GetDropItem(x, y int) *ToolItem     { return toolItemGroupGetDropItem(t, x, y) }
func (t *ToolItemGroup) GetEllipsize() T.PangoEllipsizeMode { return toolItemGroupGetEllipsize(t) }
func (t *ToolItemGroup) GetHeaderRelief() T.GtkReliefStyle  { return toolItemGroupGetHeaderRelief(t) }
func (t *ToolItemGroup) GetItemPosition(item *ToolItem) int {
	return toolItemGroupGetItemPosition(t, item)
}
func (t *ToolItemGroup) GetLabel() string                    { return toolItemGroupGetLabel(t) }
func (t *ToolItemGroup) GetLabelWidget() *Widget             { return toolItemGroupGetLabelWidget(t) }
func (t *ToolItemGroup) GetNItems() uint                     { return toolItemGroupGetNItems(t) }
func (t *ToolItemGroup) GetNthItem(index uint) *ToolItem     { return toolItemGroupGetNthItem(t, index) }
func (t *ToolItemGroup) Insert(item *ToolItem, position int) { toolItemGroupInsert(t, item, position) }
func (t *ToolItemGroup) SetCollapsed(collapsed T.Gboolean)   { toolItemGroupSetCollapsed(t, collapsed) }
func (t *ToolItemGroup) SetEllipsize(ellipsize T.PangoEllipsizeMode) {
	toolItemGroupSetEllipsize(t, ellipsize)
}
func (t *ToolItemGroup) SetHeaderRelief(style T.GtkReliefStyle) {
	toolItemGroupSetHeaderRelief(t, style)
}
func (t *ToolItemGroup) SetItemPosition(item *ToolItem, position int) {
	toolItemGroupSetItemPosition(t, item, position)
}
func (t *ToolItemGroup) SetLabel(label string) { toolItemGroupSetLabel(t, label) }
func (t *ToolItemGroup) SetLabelWidget(labelWidget *Widget) {
	toolItemGroupSetLabelWidget(t, labelWidget)
}

type ToolPalette struct {
	Parent T.GtkContainer
	_      *struct{}
}

type ToolPaletteDragTargets T.Enum

const (
	TOOL_PALETTE_DRAG_ITEMS ToolPaletteDragTargets = 1 << iota
	TOOL_PALETTE_DRAG_GROUPS
)

var (
	ToolPaletteGetType func() T.GType
	ToolPaletteNew     func() *Widget

	ToolPaletteDragTargetsGetType func() T.GType
	ToolPaletteGetDragTargetGroup func() *T.GtkTargetEntry
	ToolPaletteGetDragTargetItem  func() *T.GtkTargetEntry

	toolPaletteAddDragDest      func(t *ToolPalette, widget *Widget, flags T.GtkDestDefaults, targets ToolPaletteDragTargets, actions T.GdkDragAction)
	toolPaletteGetDragItem      func(t *ToolPalette, selection *T.GtkSelectionData) *Widget
	toolPaletteGetDropGroup     func(t *ToolPalette, x, y int) *ToolItemGroup
	toolPaletteGetDropItem      func(t *ToolPalette, x, y int) *ToolItem
	toolPaletteGetExclusive     func(t *ToolPalette, group *ToolItemGroup) T.Gboolean
	toolPaletteGetExpand        func(t *ToolPalette, group *ToolItemGroup) T.Gboolean
	toolPaletteGetGroupPosition func(t *ToolPalette, group *ToolItemGroup) int
	toolPaletteGetHadjustment   func(t *ToolPalette) *Adjustment
	toolPaletteGetIconSize      func(t *ToolPalette) IconSize
	toolPaletteGetStyle         func(t *ToolPalette) ToolbarStyle
	toolPaletteGetVadjustment   func(t *ToolPalette) *Adjustment
	toolPaletteSetDragSource    func(t *ToolPalette, targets ToolPaletteDragTargets)
	toolPaletteSetExclusive     func(t *ToolPalette, group *ToolItemGroup, exclusive T.Gboolean)
	toolPaletteSetExpand        func(t *ToolPalette, group *ToolItemGroup, expand T.Gboolean)
	toolPaletteSetGroupPosition func(t *ToolPalette, group *ToolItemGroup, position int)
	toolPaletteSetIconSize      func(t *ToolPalette, iconSize IconSize)
	toolPaletteSetStyle         func(t *ToolPalette, style ToolbarStyle)
	toolPaletteUnsetIconSize    func(t *ToolPalette)
	toolPaletteUnsetStyle       func(t *ToolPalette)
)

func (t *ToolPalette) AddDragDest(widget *Widget, flags T.GtkDestDefaults, targets ToolPaletteDragTargets, actions T.GdkDragAction) {
	toolPaletteAddDragDest(t, widget, flags, targets, actions)
}
func (t *ToolPalette) GetDragItem(selection *T.GtkSelectionData) *Widget {
	return toolPaletteGetDragItem(t, selection)
}
func (t *ToolPalette) GetDropGroup(x, y int) *ToolItemGroup { return toolPaletteGetDropGroup(t, x, y) }
func (t *ToolPalette) GetDropItem(x, y int) *ToolItem       { return toolPaletteGetDropItem(t, x, y) }
func (t *ToolPalette) GetExclusive(group *ToolItemGroup) T.Gboolean {
	return toolPaletteGetExclusive(t, group)
}
func (t *ToolPalette) GetExpand(group *ToolItemGroup) T.Gboolean {
	return toolPaletteGetExpand(t, group)
}
func (t *ToolPalette) GetGroupPosition(group *ToolItemGroup) int {
	return toolPaletteGetGroupPosition(t, group)
}
func (t *ToolPalette) GetHadjustment() *Adjustment { return toolPaletteGetHadjustment(t) }
func (t *ToolPalette) GetIconSize() IconSize       { return toolPaletteGetIconSize(t) }
func (t *ToolPalette) GetStyle() ToolbarStyle      { return toolPaletteGetStyle(t) }
func (t *ToolPalette) GetVadjustment() *Adjustment { return toolPaletteGetVadjustment(t) }
func (t *ToolPalette) SetDragSource(targets ToolPaletteDragTargets) {
	toolPaletteSetDragSource(t, targets)
}
func (t *ToolPalette) SetExclusive(group *ToolItemGroup, exclusive T.Gboolean) {
	toolPaletteSetExclusive(t, group, exclusive)
}
func (t *ToolPalette) SetExpand(group *ToolItemGroup, expand T.Gboolean) {
	toolPaletteSetExpand(t, group, expand)
}
func (t *ToolPalette) SetGroupPosition(group *ToolItemGroup, position int) {
	toolPaletteSetGroupPosition(t, group, position)
}
func (t *ToolPalette) SetIconSize(iconSize IconSize) { toolPaletteSetIconSize(t, iconSize) }
func (t *ToolPalette) SetStyle(style ToolbarStyle)   { toolPaletteSetStyle(t, style) }
func (t *ToolPalette) UnsetIconSize()                { toolPaletteUnsetIconSize(t) }
func (t *ToolPalette) UnsetStyle()                   { toolPaletteUnsetStyle(t) }

type ToolShell struct{}

var (
	ToolShellGetType func() T.GType

	toolShellGetEllipsizeMode   func(t *ToolShell) T.PangoEllipsizeMode
	toolShellGetIconSize        func(t *ToolShell) IconSize
	toolShellGetOrientation     func(t *ToolShell) T.GtkOrientation
	toolShellGetReliefStyle     func(t *ToolShell) T.GtkReliefStyle
	toolShellGetStyle           func(t *ToolShell) ToolbarStyle
	toolShellGetTextAlignment   func(t *ToolShell) float32
	toolShellGetTextOrientation func(t *ToolShell) T.GtkOrientation
	toolShellGetTextSizeGroup   func(t *ToolShell) *SizeGroup
	toolShellRebuildMenu        func(t *ToolShell)
)

func (t *ToolShell) GetEllipsizeMode() T.PangoEllipsizeMode { return toolShellGetEllipsizeMode(t) }
func (t *ToolShell) GetIconSize() IconSize                  { return toolShellGetIconSize(t) }
func (t *ToolShell) GetOrientation() T.GtkOrientation       { return toolShellGetOrientation(t) }
func (t *ToolShell) GetReliefStyle() T.GtkReliefStyle       { return toolShellGetReliefStyle(t) }
func (t *ToolShell) GetStyle() ToolbarStyle                 { return toolShellGetStyle(t) }
func (t *ToolShell) GetTextAlignment() float32              { return toolShellGetTextAlignment(t) }
func (t *ToolShell) GetTextOrientation() T.GtkOrientation   { return toolShellGetTextOrientation(t) }
func (t *ToolShell) GetTextSizeGroup() *SizeGroup           { return toolShellGetTextSizeGroup(t) }
func (t *ToolShell) RebuildMenu()                           { toolShellRebuildMenu(t) }

var (
	TooltipGetType func() T.GType

	tooltipSetCustom           func(t *Tooltip, customWidget *Widget)
	tooltipSetIcon             func(t *Tooltip, pixbuf *T.GdkPixbuf)
	tooltipSetIconFromGicon    func(t *Tooltip, gicon *T.GIcon, size IconSize)
	tooltipSetIconFromIconName func(t *Tooltip, iconName string, size IconSize)
	tooltipSetIconFromStock    func(t *Tooltip, stockId string, size IconSize)
	tooltipSetMarkup           func(t *Tooltip, markup string)
	tooltipSetText             func(t *Tooltip, text string)
	tooltipSetTipArea          func(t *Tooltip, rect *T.GdkRectangle)
)

type Tooltip struct{}

func (t *Tooltip) SetCustom(customWidget *Widget) { tooltipSetCustom(t, customWidget) }
func (t *Tooltip) SetIcon(pixbuf *T.GdkPixbuf)    { tooltipSetIcon(t, pixbuf) }
func (t *Tooltip) SetIconFromGicon(gicon *T.GIcon, size IconSize) {
	tooltipSetIconFromGicon(t, gicon, size)
}
func (t *Tooltip) SetIconFromIconName(iconName string, size IconSize) {
	tooltipSetIconFromIconName(t, iconName, size)
}
func (t *Tooltip) SetIconFromStock(stockId string, size IconSize) {
	tooltipSetIconFromStock(t, stockId, size)
}
func (t *Tooltip) SetMarkup(markup string)         { tooltipSetMarkup(t, markup) }
func (t *Tooltip) SetText(text string)             { tooltipSetText(t, text) }
func (t *Tooltip) SetTipArea(rect *T.GdkRectangle) { tooltipSetTipArea(t, rect) }

type (
	Tooltips struct {
		Parent         T.GtkObject
		TipWindow      *Widget
		TipLabel       *Widget
		ActiveTipsData *TooltipsData
		TipsDataList   *T.GList
		Bits, Bits2    uint //TODO(t): 33 bits Alignment/size?
		// delay : 30
		// enabled : 1
		// have_grab : 1
		// use_sticky_delay : 1
		TimerTag    int
		LastPopdown T.GTimeVal
	}

	TooltipsData struct {
		Tooltips    *Tooltips
		Widget      *Widget
		Tip_text    *T.Gchar
		Tip_private *T.Gchar
	}
)

var (
	TooltipsGetType func() T.GType
	TooltipsNew     func() *Tooltips

	tooltipsDisable     func(t *Tooltips)
	tooltipsEnable      func(t *Tooltips)
	tooltipsForceWindow func(t *Tooltips)
	tooltipsSetDelay    func(t *Tooltips, delay uint)
	tooltipsSetTip      func(t *Tooltips, widget *Widget, tipText, tipPrivate string)
)

func (t *Tooltips) Disable()            { tooltipsDisable(t) }
func (t *Tooltips) Enable()             { tooltipsEnable(t) }
func (t *Tooltips) ForceWindow()        { tooltipsForceWindow(t) }
func (t *Tooltips) SetDelay(delay uint) { tooltipsSetDelay(t, delay) }
func (t *Tooltips) SetTip(widget *Widget, tipText, tipPrivate string) {
	tooltipsSetTip(t, widget, tipText, tipPrivate)
}

type (
	Tree struct {
		Container     T.GtkContainer
		Children      *T.GList
		RootTree      *Tree
		TreeOwner     *Widget
		Selection     *T.GList
		Level         uint
		IndentValue   uint
		CurrentIndent uint
		Bits          uint
		// SelectionMode : 2
		// ViewMode : 1
		// ViewLine : 1
	}
)

var (
	TreeNew     func() *Widget
	TreeGetType func() T.GType

	treeAppend           func(t *Tree, treeItem *Widget)
	treeChildPosition    func(t *Tree, child *Widget) int
	treeClearItems       func(t *Tree, start int, end int)
	treeInsert           func(t *Tree, treeItem *Widget, position int)
	treePrepend          func(t *Tree, treeItem *Widget)
	treeRemoveItem       func(t *Tree, child *Widget)
	treeRemoveItems      func(t *Tree, items *T.GList)
	treeSelectChild      func(t *Tree, treeItem *Widget)
	treeSelectItem       func(t *Tree, item int)
	treeSetSelectionMode func(t *Tree, mode T.GtkSelectionMode)
	treeSetViewLines     func(t *Tree, flag T.Gboolean)
	treeSetViewMode      func(t *Tree, mode TreeViewMode)
	treeUnselectChild    func(t *Tree, treeItem *Widget)
	treeUnselectItem     func(t *Tree, item int)
)

func (t *Tree) Append(treeItem *Widget)                  { treeAppend(t, treeItem) }
func (t *Tree) ChildPosition(child *Widget) int          { return treeChildPosition(t, child) }
func (t *Tree) ClearItems(start, end int)                { treeClearItems(t, start, end) }
func (t *Tree) Insert(treeItem *Widget, position int)    { treeInsert(t, treeItem, position) }
func (t *Tree) Prepend(treeItem *Widget)                 { treePrepend(t, treeItem) }
func (t *Tree) RemoveItem(child *Widget)                 { treeRemoveItem(t, child) }
func (t *Tree) RemoveItems(items *T.GList)               { treeRemoveItems(t, items) }
func (t *Tree) SelectChild(treeItem *Widget)             { treeSelectChild(t, treeItem) }
func (t *Tree) SelectItem(item int)                      { treeSelectItem(t, item) }
func (t *Tree) SetSelectionMode(mode T.GtkSelectionMode) { treeSetSelectionMode(t, mode) }
func (t *Tree) SetViewLines(flag T.Gboolean)             { treeSetViewLines(t, flag) }
func (t *Tree) SetViewMode(mode TreeViewMode)            { treeSetViewMode(t, mode) }
func (t *Tree) UnselectChild(treeItem *Widget)           { treeUnselectChild(t, treeItem) }
func (t *Tree) UnselectItem(item int)                    { treeUnselectItem(t, item) }

type TreeDragSource struct{}

var (
	TreeDragSourceGetType func() T.GType

	treeDragSourceDragDataDelete func(t *TreeDragSource, path *TreePath) T.Gboolean
	treeDragSourceDragDataGet    func(t *TreeDragSource, path *TreePath, selectionData *T.GtkSelectionData) T.Gboolean
	treeDragSourceRowDraggable   func(t *TreeDragSource, path *TreePath) T.Gboolean
)

func (t *TreeDragSource) DragDataDelete(path *TreePath) T.Gboolean {
	return treeDragSourceDragDataDelete(t, path)
}
func (t *TreeDragSource) DragDataGet(path *TreePath, selectionData *T.GtkSelectionData) T.Gboolean {
	return treeDragSourceDragDataGet(t, path, selectionData)
}
func (t *TreeDragSource) RowDraggable(path *TreePath) T.Gboolean {
	return treeDragSourceRowDraggable(t, path)
}

type TreeDragDest struct{}

var (
	TreeDragDestGetType func() T.GType

	treeDragDestDragDataReceived func(t *TreeDragDest, dest *TreePath, selectionData *T.GtkSelectionData) T.Gboolean
	treeDragDestRowDropPossible  func(t *TreeDragDest, destPath *TreePath, selectionData *T.GtkSelectionData) T.Gboolean
)

func (t *TreeDragDest) DragDataReceived(dest *TreePath, selectionData *T.GtkSelectionData) T.Gboolean {
	return treeDragDestDragDataReceived(t, dest, selectionData)
}
func (t *TreeDragDest) RowDropPossible(destPath *TreePath, selectionData *T.GtkSelectionData) T.Gboolean {
	return treeDragDestRowDropPossible(t, destPath, selectionData)
}

type TreeItem struct {
	Item             Item
	Subtree          *Widget
	Pixmaps_box      *Widget
	Plus_pix_widget  *Widget
	Minus_pix_widget *Widget
	Pixmaps          *T.GList
	Expanded         uint // : 1
}

var (
	TreeItemGetType      func() T.GType
	TreeItemNew          func() *Widget
	TreeItemNewWithLabel func(label string) *Widget

	treeItemCollapse      func(t *TreeItem)
	treeItemDeselect      func(t *TreeItem)
	treeItemExpand        func(t *TreeItem)
	treeItemRemoveSubtree func(t *TreeItem)
	treeItemSelect        func(t *TreeItem)
	treeItemSetSubtree    func(t *TreeItem, subtree *Widget)
)

func (t *TreeItem) Collapse()                  { treeItemCollapse(t) }
func (t *TreeItem) Deselect()                  { treeItemDeselect(t) }
func (t *TreeItem) Expand()                    { treeItemExpand(t) }
func (t *TreeItem) RemoveSubtree()             { treeItemRemoveSubtree(t) }
func (t *TreeItem) Select()                    { treeItemSelect(t) }
func (t *TreeItem) SetSubtree(subtree *Widget) { treeItemSetSubtree(t, subtree) }

type TreeIter struct {
	Stamp     int
	UserData  T.Gpointer
	UserData2 T.Gpointer
	UserData3 T.Gpointer
}

var (
	TreeIterGetType func() T.GType

	treeIterCopy func(t *TreeIter) *TreeIter
	treeIterFree func(t *TreeIter)
)

func (t *TreeIter) Copy() *TreeIter { return treeIterCopy(t) }
func (t *TreeIter) Free()           { treeIterFree(t) }

type (
	TreeModel struct{}

	TreeModelForeachFunc func(
		model *TreeModel,
		path *TreePath,
		iter *TreeIter,
		data T.Gpointer) T.Gboolean
)

type TreeModelFlags T.Enum

const (
	TREE_MODEL_ITERS_PERSIST TreeModelFlags = 1 << iota
	TREE_MODEL_LIST_ONLY
)

var (
	TreeModelGetType func() T.GType

	TreeModelFlagsGetType func() T.GType

	treeModelForeach            func(t *TreeModel, f TreeModelForeachFunc, userData T.Gpointer)
	treeModelGet                func(t *TreeModel, iter *TreeIter, v ...VArg)
	treeModelGetColumnType      func(t *TreeModel, index int) T.GType
	treeModelGetFlags           func(t *TreeModel) TreeModelFlags
	treeModelGetIter            func(t *TreeModel, iter *TreeIter, path *TreePath) T.Gboolean
	treeModelGetIterFirst       func(t *TreeModel, iter *TreeIter) T.Gboolean
	treeModelGetIterFromString  func(t *TreeModel, iter *TreeIter, pathString string) T.Gboolean
	treeModelGetNColumns        func(t *TreeModel) int
	treeModelGetPath            func(t *TreeModel, iter *TreeIter) *TreePath
	treeModelGetStringFromIter  func(t *TreeModel, iter *TreeIter) string
	treeModelGetValist          func(t *TreeModel, iter *TreeIter, varArgs T.VaList)
	treeModelGetValue           func(t *TreeModel, iter *TreeIter, column int, value *T.GValue)
	treeModelIterChildren       func(t *TreeModel, iter *TreeIter, parent *TreeIter) T.Gboolean
	treeModelIterHasChild       func(t *TreeModel, iter *TreeIter) T.Gboolean
	treeModelIterNChildren      func(t *TreeModel, iter *TreeIter) int
	treeModelIterNext           func(t *TreeModel, iter *TreeIter) T.Gboolean
	treeModelIterNthChild       func(t *TreeModel, iter *TreeIter, parent *TreeIter, n int) T.Gboolean
	treeModelIterParent         func(t *TreeModel, iter *TreeIter, child *TreeIter) T.Gboolean
	treeModelRefNode            func(t *TreeModel, iter *TreeIter)
	treeModelRowChanged         func(t *TreeModel, path *TreePath, iter *TreeIter)
	treeModelRowDeleted         func(t *TreeModel, path *TreePath)
	treeModelRowHasChildToggled func(t *TreeModel, path *TreePath, iter *TreeIter)
	treeModelRowInserted        func(t *TreeModel, path *TreePath, iter *TreeIter)
	treeModelRowsReordered      func(t *TreeModel, path *TreePath, iter *TreeIter, newOrder *int)
	treeModelUnrefNode          func(t *TreeModel, iter *TreeIter)
)

func (t *TreeModel) Foreach(f TreeModelForeachFunc, userData T.Gpointer) {
	treeModelForeach(t, f, userData)
}
func (t *TreeModel) Get(iter *TreeIter, v ...VArg)   { treeModelGet(t, iter, v) }
func (t *TreeModel) GetColumnType(index int) T.GType { return treeModelGetColumnType(t, index) }
func (t *TreeModel) GetFlags() TreeModelFlags        { return treeModelGetFlags(t) }
func (t *TreeModel) GetIter(iter *TreeIter, path *TreePath) T.Gboolean {
	return treeModelGetIter(t, iter, path)
}
func (t *TreeModel) GetIterFirst(iter *TreeIter) T.Gboolean { return treeModelGetIterFirst(t, iter) }
func (t *TreeModel) GetIterFromString(iter *TreeIter, pathString string) T.Gboolean {
	return treeModelGetIterFromString(t, iter, pathString)
}
func (t *TreeModel) GetNColumns() int                 { return treeModelGetNColumns(t) }
func (t *TreeModel) GetPath(iter *TreeIter) *TreePath { return treeModelGetPath(t, iter) }
func (t *TreeModel) GetStringFromIter(iter *TreeIter) string {
	return treeModelGetStringFromIter(t, iter)
}
func (t *TreeModel) GetValist(iter *TreeIter, varArgs T.VaList) { treeModelGetValist(t, iter, varArgs) }
func (t *TreeModel) GetValue(iter *TreeIter, column int, value *T.GValue) {
	treeModelGetValue(t, iter, column, value)
}
func (t *TreeModel) IterChildren(iter *TreeIter, parent *TreeIter) T.Gboolean {
	return treeModelIterChildren(t, iter, parent)
}
func (t *TreeModel) IterHasChild(iter *TreeIter) T.Gboolean { return treeModelIterHasChild(t, iter) }
func (t *TreeModel) IterNChildren(iter *TreeIter) int       { return treeModelIterNChildren(t, iter) }
func (t *TreeModel) IterNext(iter *TreeIter) T.Gboolean     { return treeModelIterNext(t, iter) }
func (t *TreeModel) IterNthChild(iter, parent *TreeIter, n int) T.Gboolean {
	return treeModelIterNthChild(t, iter, parent, n)
}
func (t *TreeModel) IterParent(iter, child *TreeIter) T.Gboolean {
	return treeModelIterParent(t, iter, child)
}
func (t *TreeModel) RefNode(iter *TreeIter) { treeModelRefNode(t, iter) }
func (t *TreeModel) RowChanged(path *TreePath, iter *TreeIter) {
	treeModelRowChanged(t, path, iter)
}
func (t *TreeModel) RowDeleted(path *TreePath) { treeModelRowDeleted(t, path) }
func (t *TreeModel) RowHasChildToggled(path *TreePath, iter *TreeIter) {
	treeModelRowHasChildToggled(t, path, iter)
}
func (t *TreeModel) RowInserted(path *TreePath, iter *TreeIter) {
	treeModelRowInserted(t, path, iter)
}
func (t *TreeModel) RowsReordered(path *TreePath, iter *TreeIter, newOrder *int) {
	treeModelRowsReordered(t, path, iter, newOrder)
}
func (t *TreeModel) UnrefNode(iter *TreeIter) { treeModelUnrefNode(t, iter) }

type (
	TreeModelFilter struct {
		Parent T.GObject
		_      *struct{}
	}

	TreeModelFilterVisibleFunc func(
		model *TreeModel,
		iter *TreeIter,
		data T.Gpointer) T.Gboolean

	TreeModelFilterModifyFunc func(
		model *TreeModel,
		iter *TreeIter,
		value *T.GValue,
		column int,
		data T.Gpointer)
)

var (
	TreeModelFilterGetType func() T.GType
	TreeModelFilterNew     func(childModel *TreeModel, root *TreePath) *TreeModel

	treeModelFilterClearCache             func(t *TreeModelFilter)
	treeModelFilterConvertChildIterToIter func(t *TreeModelFilter, filterIter *TreeIter, childIter *TreeIter) T.Gboolean
	treeModelFilterConvertChildPathToPath func(t *TreeModelFilter, childPath *TreePath) *TreePath
	treeModelFilterConvertIterToChildIter func(t *TreeModelFilter, childIter *TreeIter, filterIter *TreeIter)
	treeModelFilterConvertPathToChildPath func(t *TreeModelFilter, filterPath *TreePath) *TreePath
	treeModelFilterGetModel               func(t *TreeModelFilter) *TreeModel
	treeModelFilterRefilter               func(t *TreeModelFilter)
	treeModelFilterSetModifyFunc          func(t *TreeModelFilter, nColumns int, types *T.GType, f TreeModelFilterModifyFunc, data T.Gpointer, destroy T.GDestroyNotify)
	treeModelFilterSetVisibleColumn       func(t *TreeModelFilter, column int)
	treeModelFilterSetVisibleFunc         func(t *TreeModelFilter, f TreeModelFilterVisibleFunc, data T.Gpointer, destroy T.GDestroyNotify)
)

func (t *TreeModelFilter) ClearCache() { treeModelFilterClearCache(t) }
func (t *TreeModelFilter) ConvertChildIterToIter(filterIter *TreeIter, childIter *TreeIter) T.Gboolean {
	return treeModelFilterConvertChildIterToIter(t, filterIter, childIter)
}
func (t *TreeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	return treeModelFilterConvertChildPathToPath(t, childPath)
}
func (t *TreeModelFilter) ConvertIterToChildIter(childIter *TreeIter, filterIter *TreeIter) {
	treeModelFilterConvertIterToChildIter(t, childIter, filterIter)
}
func (t *TreeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	return treeModelFilterConvertPathToChildPath(t, filterPath)
}
func (t *TreeModelFilter) GetModel() *TreeModel { return treeModelFilterGetModel(t) }
func (t *TreeModelFilter) Refilter()            { treeModelFilterRefilter(t) }
func (t *TreeModelFilter) SetModifyFunc(nColumns int, types *T.GType, f TreeModelFilterModifyFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeModelFilterSetModifyFunc(t, nColumns, types, f, data, destroy)
}
func (t *TreeModelFilter) SetVisibleColumn(column int) { treeModelFilterSetVisibleColumn(t, column) }
func (t *TreeModelFilter) SetVisibleFunc(f TreeModelFilterVisibleFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeModelFilterSetVisibleFunc(t, f, data, destroy)
}

type TreeModelSort struct {
	Parent             T.GObject
	Root               T.Gpointer
	Stamp              int
	ChildFlags         uint
	ChildModel         *TreeModel
	ZeroRefCount       int
	SortList           *T.GList
	SortColumnId       int
	Order              T.GtkSortType
	DefaultSortFunc    TreeIterCompareFunc
	DefaultSortData    T.Gpointer
	DefaultSortDestroy T.GDestroyNotify
	ChangedId          uint
	InsertedId         uint
	HasChildToggledId  uint
	DeletedId          uint
	ReorderedId        uint
}

var (
	TreeModelSortGetType      func() T.GType
	TreeModelSortNewWithModel func(childModel *TreeModel) *TreeModel

	treeModelSortClearCache             func(t *TreeModelSort)
	treeModelSortConvertChildIterToIter func(t *TreeModelSort, sortIter *TreeIter, childIter *TreeIter) T.Gboolean
	treeModelSortConvertChildPathToPath func(t *TreeModelSort, childPath *TreePath) *TreePath
	treeModelSortConvertIterToChildIter func(t *TreeModelSort, childIter *TreeIter, sortedIter *TreeIter)
	treeModelSortConvertPathToChildPath func(t *TreeModelSort, sortedPath *TreePath) *TreePath
	treeModelSortGetModel               func(t *TreeModelSort) *TreeModel
	treeModelSortIterIsValid            func(t *TreeModelSort, iter *TreeIter) T.Gboolean
	treeModelSortResetDefaultSortFunc   func(t *TreeModelSort)
)

func (t *TreeModelSort) ClearCache() { treeModelSortClearCache(t) }
func (t *TreeModelSort) ConvertChildIterToIter(sortIter *TreeIter, childIter *TreeIter) T.Gboolean {
	return treeModelSortConvertChildIterToIter(t, sortIter, childIter)
}
func (t *TreeModelSort) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	return treeModelSortConvertChildPathToPath(t, childPath)
}
func (t *TreeModelSort) ConvertIterToChildIter(childIter *TreeIter, sortedIter *TreeIter) {
	treeModelSortConvertIterToChildIter(t, childIter, sortedIter)
}
func (t *TreeModelSort) ConvertPathToChildPath(sortedPath *TreePath) *TreePath {
	return treeModelSortConvertPathToChildPath(t, sortedPath)
}
func (t *TreeModelSort) GetModel() *TreeModel { return treeModelSortGetModel(t) }
func (t *TreeModelSort) IterIsValid(iter *TreeIter) T.Gboolean {
	return treeModelSortIterIsValid(t, iter)
}
func (t *TreeModelSort) ResetDefaultSortFunc() { treeModelSortResetDefaultSortFunc(t) }

type TreePath struct{}

var (
	TreePathGetType        func() T.GType
	TreePathNew            func() *TreePath
	TreePathNewFirst       func() *TreePath
	TreePathNewFromString  func(path string) *TreePath
	TreePathNewFromIndices func(firstIndex int, v ...VArg) *TreePath

	treePathAppendIndex         func(t *TreePath, index int)
	treePathCompare             func(t *TreePath, p2 *TreePath) int
	treePathCopy                func(t *TreePath) *TreePath
	treePathDown                func(t *TreePath)
	treePathFree                func(t *TreePath)
	treePathGetDepth            func(t *TreePath) int
	treePathGetIndices          func(t *TreePath) *int
	treePathGetIndicesWithDepth func(t *TreePath, depth *int) *int
	treePathIsAncestor          func(t *TreePath, descendant *TreePath) T.Gboolean
	treePathIsDescendant        func(t *TreePath, ancestor *TreePath) T.Gboolean
	treePathNext                func(t *TreePath)
	treePathPrependIndex        func(t *TreePath, index int)
	treePathPrev                func(t *TreePath) T.Gboolean
	treePathToString            func(t *TreePath) string
	treePathUp                  func(t *TreePath) T.Gboolean
)

func (t *TreePath) AppendIndex(index int)               { treePathAppendIndex(t, index) }
func (t *TreePath) Compare(p2 *TreePath) int            { return treePathCompare(t, p2) }
func (t *TreePath) Copy() *TreePath                     { return treePathCopy(t) }
func (t *TreePath) Down()                               { treePathDown(t) }
func (t *TreePath) Free()                               { treePathFree(t) }
func (t *TreePath) GetDepth() int                       { return treePathGetDepth(t) }
func (t *TreePath) GetIndices() *int                    { return treePathGetIndices(t) }
func (t *TreePath) GetIndicesWithDepth(depth *int) *int { return treePathGetIndicesWithDepth(t, depth) }
func (t *TreePath) IsAncestor(descendant *TreePath) T.Gboolean {
	return treePathIsAncestor(t, descendant)
}
func (t *TreePath) IsDescendant(ancestor *TreePath) T.Gboolean {
	return treePathIsDescendant(t, ancestor)
}
func (t *TreePath) Next()                  { treePathNext(t) }
func (t *TreePath) PrependIndex(index int) { treePathPrependIndex(t, index) }
func (t *TreePath) Prev() T.Gboolean       { return treePathPrev(t) }
func (t *TreePath) ToString() string       { return treePathToString(t) }
func (t *TreePath) Up() T.Gboolean         { return treePathUp(t) }

type TreeRowReference struct{}

var (
	TreeRowReferenceGetType func() T.GType
	TreeRowReferenceNew     func(model *TreeModel, path *TreePath) *TreeRowReference

	TreeRowReferenceDeleted   func(proxy *T.GObject, path *TreePath)
	TreeRowReferenceInserted  func(proxy *T.GObject, path *TreePath)
	TreeRowReferenceNewProxy  func(proxy *T.GObject, model *TreeModel, path *TreePath) *TreeRowReference
	TreeRowReferenceReordered func(proxy *T.GObject, path *TreePath, iter *TreeIter, newOrder *int)

	treeRowReferenceCopy     func(t *TreeRowReference) *TreeRowReference
	treeRowReferenceFree     func(t *TreeRowReference)
	treeRowReferenceGetModel func(t *TreeRowReference) *TreeModel
	treeRowReferenceGetPath  func(t *TreeRowReference) *TreePath
	treeRowReferenceValid    func(t *TreeRowReference) T.Gboolean
)

func (t *TreeRowReference) Copy() *TreeRowReference { return treeRowReferenceCopy(t) }
func (t *TreeRowReference) Free()                   { treeRowReferenceFree(t) }
func (t *TreeRowReference) GetModel() *TreeModel    { return treeRowReferenceGetModel(t) }
func (t *TreeRowReference) GetPath() *TreePath      { return treeRowReferenceGetPath(t) }
func (t *TreeRowReference) Valid() T.Gboolean       { return treeRowReferenceValid(t) }

type (
	TreeSelection struct {
		Parent   T.GObject
		TreeView *TreeView
		Type     T.GtkSelectionMode
		UserFunc TreeSelectionFunc
		UserData T.Gpointer
		Destroy  T.GDestroyNotify
	}

	TreeSelectionForeachFunc func(model *TreeModel,
		path *TreePath, iter *TreeIter, data T.Gpointer)
)

var (
	TreeSelectionGetType func() T.GType

	treeSelectionCountSelectedRows func(t *TreeSelection) int
	treeSelectionGetMode           func(t *TreeSelection) T.GtkSelectionMode
	treeSelectionGetSelected       func(t *TreeSelection, model **TreeModel, iter *TreeIter) T.Gboolean
	treeSelectionGetSelectedRows   func(t *TreeSelection, model **TreeModel) *T.GList
	treeSelectionGetSelectFunction func(t *TreeSelection) TreeSelectionFunc
	treeSelectionGetTreeView       func(t *TreeSelection) *TreeView
	treeSelectionGetUserData       func(t *TreeSelection) T.Gpointer
	treeSelectionIterIsSelected    func(t *TreeSelection, iter *TreeIter) T.Gboolean
	treeSelectionPathIsSelected    func(t *TreeSelection, path *TreePath) T.Gboolean
	treeSelectionSelectAll         func(t *TreeSelection)
	treeSelectionSelectedForeach   func(t *TreeSelection, f TreeSelectionForeachFunc, data T.Gpointer)
	treeSelectionSelectIter        func(t *TreeSelection, iter *TreeIter)
	treeSelectionSelectPath        func(t *TreeSelection, path *TreePath)
	treeSelectionSelectRange       func(t *TreeSelection, startPath *TreePath, endPath *TreePath)
	treeSelectionSetMode           func(t *TreeSelection, typ T.GtkSelectionMode)
	treeSelectionSetSelectFunction func(t *TreeSelection, f TreeSelectionFunc, data T.Gpointer, destroy T.GDestroyNotify)
	treeSelectionUnselectAll       func(t *TreeSelection)
	treeSelectionUnselectIter      func(t *TreeSelection, iter *TreeIter)
	treeSelectionUnselectPath      func(t *TreeSelection, path *TreePath)
	treeSelectionUnselectRange     func(t *TreeSelection, startPath *TreePath, endPath *TreePath)
)

func (t *TreeSelection) CountSelectedRows() int      { return treeSelectionCountSelectedRows(t) }
func (t *TreeSelection) GetMode() T.GtkSelectionMode { return treeSelectionGetMode(t) }
func (t *TreeSelection) GetSelected(model **TreeModel, iter *TreeIter) T.Gboolean {
	return treeSelectionGetSelected(t, model, iter)
}
func (t *TreeSelection) GetSelectedRows(model **TreeModel) *T.GList {
	return treeSelectionGetSelectedRows(t, model)
}
func (t *TreeSelection) GetSelectFunction() TreeSelectionFunc {
	return treeSelectionGetSelectFunction(t)
}
func (t *TreeSelection) GetTreeView() *TreeView  { return treeSelectionGetTreeView(t) }
func (t *TreeSelection) GetUserData() T.Gpointer { return treeSelectionGetUserData(t) }
func (t *TreeSelection) IterIsSelected(iter *TreeIter) T.Gboolean {
	return treeSelectionIterIsSelected(t, iter)
}
func (t *TreeSelection) PathIsSelected(path *TreePath) T.Gboolean {
	return treeSelectionPathIsSelected(t, path)
}
func (t *TreeSelection) SelectAll() { treeSelectionSelectAll(t) }
func (t *TreeSelection) SelectedForeach(f TreeSelectionForeachFunc, data T.Gpointer) {
	treeSelectionSelectedForeach(t, f, data)
}
func (t *TreeSelection) SelectIter(iter *TreeIter) { treeSelectionSelectIter(t, iter) }
func (t *TreeSelection) SelectPath(path *TreePath) { treeSelectionSelectPath(t, path) }
func (t *TreeSelection) SelectRange(startPath, endPath *TreePath) {
	treeSelectionSelectRange(t, startPath, endPath)
}
func (t *TreeSelection) SetMode(typ T.GtkSelectionMode) { treeSelectionSetMode(t, typ) }
func (t *TreeSelection) SetSelectFunction(f TreeSelectionFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeSelectionSetSelectFunction(t, f, data, destroy)
}
func (t *TreeSelection) UnselectAll()                { treeSelectionUnselectAll(t) }
func (t *TreeSelection) UnselectIter(iter *TreeIter) { treeSelectionUnselectIter(t, iter) }
func (t *TreeSelection) UnselectPath(path *TreePath) { treeSelectionUnselectPath(t, path) }
func (t *TreeSelection) UnselectRange(startPath, endPath *TreePath) {
	treeSelectionUnselectRange(t, startPath, endPath)
}

type TreeSortable struct{}

var (
	TreeSortableGetType func() T.GType

	treeSortableGetSortColumnId    func(t *TreeSortable, sortColumnId *int, order *T.GtkSortType) T.Gboolean
	treeSortableHasDefaultSortFunc func(t *TreeSortable) T.Gboolean
	treeSortableSetDefaultSortFunc func(t *TreeSortable, sortFunc TreeIterCompareFunc, userData T.Gpointer, destroy T.GDestroyNotify)
	treeSortableSetSortColumnId    func(t *TreeSortable, sortColumnId int, order T.GtkSortType)
	treeSortableSetSortFunc        func(t *TreeSortable, sortColumnId int, sortFunc TreeIterCompareFunc, userData T.Gpointer, destroy T.GDestroyNotify)
	treeSortableSortColumnChanged  func(t *TreeSortable)
)

func (t *TreeSortable) GetSortColumnId(sortColumnId *int, order *T.GtkSortType) T.Gboolean {
	return treeSortableGetSortColumnId(t, sortColumnId, order)
}
func (t *TreeSortable) HasDefaultSortFunc() T.Gboolean { return treeSortableHasDefaultSortFunc(t) }
func (t *TreeSortable) SetDefaultSortFunc(sortFunc TreeIterCompareFunc, userData T.Gpointer, destroy T.GDestroyNotify) {
	treeSortableSetDefaultSortFunc(t, sortFunc, userData, destroy)
}
func (t *TreeSortable) SetSortColumnId(sortColumnId int, order T.GtkSortType) {
	treeSortableSetSortColumnId(t, sortColumnId, order)
}
func (t *TreeSortable) SetSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc, userData T.Gpointer, destroy T.GDestroyNotify) {
	treeSortableSetSortFunc(t, sortColumnId, sortFunc, userData, destroy)
}
func (t *TreeSortable) SortColumnChanged() { treeSortableSortColumnChanged(t) }

type TreeStore struct {
	Parent             T.GObject
	Stamp              int
	Root               T.Gpointer
	Last               T.Gpointer
	NColumns           int
	SortColumnId       int
	SortList           *T.GList
	Order              T.GtkSortType
	ColumnHeaders      *T.GType
	DefaultSortFunc    TreeIterCompareFunc
	DefaultSortData    T.Gpointer
	DefaultSortDestroy T.GDestroyNotify
	ColumnsDirty       uint // : 1
}

var (
	TreeStoreGetType func() T.GType
	TreeStoreNew     func(nColumns int, v ...VArg) *TreeStore
	TreeStoreNewv    func(nColumns int, types *T.GType) *TreeStore

	treeStoreAppend            func(t *TreeStore, iter *TreeIter, parent *TreeIter)
	treeStoreClear             func(t *TreeStore)
	treeStoreInsert            func(t *TreeStore, iter *TreeIter, parent *TreeIter, position int)
	treeStoreInsertAfter       func(t *TreeStore, iter *TreeIter, parent *TreeIter, sibling *TreeIter)
	treeStoreInsertBefore      func(t *TreeStore, iter *TreeIter, parent *TreeIter, sibling *TreeIter)
	treeStoreInsertWithValues  func(t *TreeStore, iter *TreeIter, parent *TreeIter, position int, v ...VArg)
	treeStoreInsertWithValuesv func(t *TreeStore, iter *TreeIter, parent *TreeIter, position int, columns *int, values *T.GValue, nValues int)
	treeStoreIsAncestor        func(t *TreeStore, iter *TreeIter, descendant *TreeIter) T.Gboolean
	treeStoreIterDepth         func(t *TreeStore, iter *TreeIter) int
	treeStoreIterIsValid       func(t *TreeStore, iter *TreeIter) T.Gboolean
	treeStoreMoveAfter         func(t *TreeStore, iter *TreeIter, position *TreeIter)
	treeStoreMoveBefore        func(t *TreeStore, iter *TreeIter, position *TreeIter)
	treeStorePrepend           func(t *TreeStore, iter *TreeIter, parent *TreeIter)
	treeStoreRemove            func(t *TreeStore, iter *TreeIter) T.Gboolean
	treeStoreReorder           func(t *TreeStore, parent *TreeIter, newOrder *int)
	treeStoreSet               func(t *TreeStore, iter *TreeIter, v ...VArg)
	treeStoreSetColumnTypes    func(t *TreeStore, nColumns int, types *T.GType)
	treeStoreSetValist         func(t *TreeStore, iter *TreeIter, varArgs T.VaList)
	treeStoreSetValue          func(t *TreeStore, iter *TreeIter, column int, value *T.GValue)
	treeStoreSetValuesv        func(t *TreeStore, iter *TreeIter, columns *int, values *T.GValue, nValues int)
	treeStoreSwap              func(t *TreeStore, a *TreeIter, b *TreeIter)
)

func (t *TreeStore) Append(iter, parent *TreeIter) { treeStoreAppend(t, iter, parent) }
func (t *TreeStore) Clear()                        { treeStoreClear(t) }
func (t *TreeStore) Insert(iter, parent *TreeIter, position int) {
	treeStoreInsert(t, iter, parent, position)
}
func (t *TreeStore) InsertAfter(iter, parent, sibling *TreeIter) {
	treeStoreInsertAfter(t, iter, parent, sibling)
}
func (t *TreeStore) InsertBefore(iter, parent, sibling *TreeIter) {
	treeStoreInsertBefore(t, iter, parent, sibling)
}
func (t *TreeStore) InsertWithValues(iter, parent *TreeIter, position int, v ...VArg) {
	treeStoreInsertWithValues(t, iter, parent, position, v)
}
func (t *TreeStore) InsertWithValuesv(iter, parent *TreeIter, position int, columns *int, values *T.GValue, nValues int) {
	treeStoreInsertWithValuesv(t, iter, parent, position, columns, values, nValues)
}
func (t *TreeStore) IsAncestor(iter, descendant *TreeIter) T.Gboolean {
	return treeStoreIsAncestor(t, iter, descendant)
}
func (t *TreeStore) IterDepth(iter *TreeIter) int            { return treeStoreIterDepth(t, iter) }
func (t *TreeStore) IterIsValid(iter *TreeIter) T.Gboolean   { return treeStoreIterIsValid(t, iter) }
func (t *TreeStore) MoveAfter(iter, position *TreeIter)      { treeStoreMoveAfter(t, iter, position) }
func (t *TreeStore) MoveBefore(iter, position *TreeIter)     { treeStoreMoveBefore(t, iter, position) }
func (t *TreeStore) Prepend(iter, parent *TreeIter)          { treeStorePrepend(t, iter, parent) }
func (t *TreeStore) Remove(iter *TreeIter) T.Gboolean        { return treeStoreRemove(t, iter) }
func (t *TreeStore) Reorder(parent *TreeIter, newOrder *int) { treeStoreReorder(t, parent, newOrder) }
func (t *TreeStore) Set(iter *TreeIter, v ...VArg)           { treeStoreSet(t, iter, v) }
func (t *TreeStore) SetColumnTypes(nColumns int, types *T.GType) {
	treeStoreSetColumnTypes(t, nColumns, types)
}
func (t *TreeStore) SetValist(iter *TreeIter, varArgs T.VaList) { treeStoreSetValist(t, iter, varArgs) }
func (t *TreeStore) SetValue(iter *TreeIter, column int, value *T.GValue) {
	treeStoreSetValue(t, iter, column, value)
}
func (t *TreeStore) SetValuesv(iter *TreeIter, columns *int, values *T.GValue, nValues int) {
	treeStoreSetValuesv(t, iter, columns, values, nValues)
}
func (t *TreeStore) Swap(a, b *TreeIter) { treeStoreSwap(t, a, b) }

type (
	TreeView struct {
		Parent T.GtkContainer
		_      *struct{}
	}

	TreeViewRowSeparatorFunc func(
		model *TreeModel,
		iter *TreeIter,
		data T.Gpointer) T.Gboolean

	TreeViewSearchEqualFunc func(
		model *TreeModel,
		column int,
		key string,
		iter *TreeIter,
		searchData T.Gpointer) T.Gboolean

	TreeViewSearchPositionFunc func(
		tree_view *TreeView,
		search_dialog *Widget,
		user_data T.Gpointer)

	TreeSelectionFunc func(selection *TreeSelection,
		model *TreeModel, path *TreePath,
		pathCurrentlySelected T.Gboolean, data T.Gpointer) T.Gboolean

	TreeCellDataFunc func(
		tree_column *TreeViewColumn,
		cell *CellRenderer,
		treeModel *TreeModel,
		iter *TreeIter,
		data T.Gpointer)

	TreeViewMappingFunc func(
		treeView *TreeView,
		path *TreePath,
		userData T.Gpointer)

	TreeViewColumnDropFunc func(
		treeView *TreeView,
		column *TreeViewColumn,
		prevColumn *TreeViewColumn,
		nextColumn *TreeViewColumn,
		data T.Gpointer) T.Gboolean

	TreeDestroyCountFunc func(
		treeView *TreeView,
		path *TreePath,
		children int,
		userData T.Gpointer)

	TreeIterCompareFunc func(model *TreeModel,
		a, b *TreeIter, userData T.Gpointer) int
)

type TreeViewDropPosition T.Enum

const (
	TREE_VIEW_DROP_BEFORE TreeViewDropPosition = iota
	TREE_VIEW_DROP_AFTER
	TREE_VIEW_DROP_INTO_OR_BEFORE
	TREE_VIEW_DROP_INTO_OR_AFTER
)

type TreeViewGridLines T.Enum

const (
	TREE_VIEW_GRID_LINES_NONE TreeViewGridLines = iota
	TREE_VIEW_GRID_LINES_HORIZONTAL
	TREE_VIEW_GRID_LINES_VERTICAL
	TREE_VIEW_GRID_LINES_BOTH
)

type TreeViewColumnSizing T.Enum

const (
	TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = iota
	TREE_VIEW_COLUMN_AUTOSIZE
	TREE_VIEW_COLUMN_FIXED
)

type TreeViewMode T.Enum

const (
	TREE_VIEW_LINE TreeViewMode = iota
	TREE_VIEW_ITEM
)

var (
	TreeViewGetType      func() T.GType
	TreeViewNew          func() *Widget
	TreeViewNewWithModel func(model *TreeModel) *Widget

	TreeViewDropPositionGetType func() T.GType
	TreeViewGridLinesGetType    func() T.GType
	TreeViewModeGetType         func() T.GType

	treeViewAppendColumn                   func(t *TreeView, column *TreeViewColumn) int
	treeViewCollapseAll                    func(t *TreeView)
	treeViewCollapseRow                    func(t *TreeView, path *TreePath) T.Gboolean
	treeViewColumnsAutosize                func(t *TreeView)
	treeViewConvertBinWindowToTreeCoords   func(t *TreeView, bx, by int, tx, ty *int)
	treeViewConvertBinWindowToWidgetCoords func(t *TreeView, bx, by int, wx, wy *int)
	treeViewConvertTreeToBinWindowCoords   func(t *TreeView, tx, ty int, bx, by *int)
	treeViewConvertTreeToWidgetCoords      func(t *TreeView, tx, ty int, wx, wy *int)
	treeViewConvertWidgetToBinWindowCoords func(t *TreeView, wx, wy int, bx, by *int)
	treeViewConvertWidgetToTreeCoords      func(t *TreeView, wx, wy int, tx, ty *int)
	treeViewCreateRowDragIcon              func(t *TreeView, path *TreePath) *T.GdkPixmap
	treeViewEnableModelDragDest            func(t *TreeView, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction)
	treeViewEnableModelDragSource          func(t *TreeView, startButtonMask T.GdkModifierType, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction)
	treeViewExpandAll                      func(t *TreeView)
	treeViewExpandRow                      func(t *TreeView, path *TreePath, openAll T.Gboolean) T.Gboolean
	treeViewExpandToPath                   func(t *TreeView, path *TreePath)
	treeViewGetBackgroundArea              func(t *TreeView, path *TreePath, column *TreeViewColumn, rect *T.GdkRectangle)
	treeViewGetBinWindow                   func(t *TreeView) *T.GdkWindow
	treeViewGetCellArea                    func(t *TreeView, path *TreePath, column *TreeViewColumn, rect *T.GdkRectangle)
	treeViewGetColumn                      func(t *TreeView, n int) *TreeViewColumn
	treeViewGetColumns                     func(t *TreeView) *T.GList
	treeViewGetCursor                      func(t *TreeView, path **TreePath, focusColumn **TreeViewColumn)
	treeViewGetDestRowAtPos                func(t *TreeView, dragX, dragY int, path **TreePath, pos *TreeViewDropPosition) T.Gboolean
	treeViewGetDragDestRow                 func(t *TreeView, path **TreePath, pos *TreeViewDropPosition)
	treeViewGetEnableSearch                func(t *TreeView) T.Gboolean
	treeViewGetEnableTreeLines             func(t *TreeView) T.Gboolean
	treeViewGetExpanderColumn              func(t *TreeView) *TreeViewColumn
	treeViewGetFixedHeightMode             func(t *TreeView) T.Gboolean
	treeViewGetGridLines                   func(t *TreeView) TreeViewGridLines
	treeViewGetHadjustment                 func(t *TreeView) *Adjustment
	treeViewGetHeadersClickable            func(t *TreeView) T.Gboolean
	treeViewGetHeadersVisible              func(t *TreeView) T.Gboolean
	treeViewGetHoverExpand                 func(t *TreeView) T.Gboolean
	treeViewGetHoverSelection              func(t *TreeView) T.Gboolean
	treeViewGetLevelIndentation            func(t *TreeView) int
	treeViewGetModel                       func(t *TreeView) *TreeModel
	treeViewGetPathAtPos                   func(t *TreeView, x, y int, path **TreePath, column **TreeViewColumn, cellX, cellY *int) T.Gboolean
	treeViewGetReorderable                 func(t *TreeView) T.Gboolean
	treeViewGetRowSeparatorFunc            func(t *TreeView) TreeViewRowSeparatorFunc
	treeViewGetRubberBanding               func(t *TreeView) T.Gboolean
	treeViewGetRulesHint                   func(t *TreeView) T.Gboolean
	treeViewGetSearchColumn                func(t *TreeView) int
	treeViewGetSearchEntry                 func(t *TreeView) *Entry
	treeViewGetSearchEqualFunc             func(t *TreeView) TreeViewSearchEqualFunc
	treeViewGetSearchPositionFunc          func(t *TreeView) TreeViewSearchPositionFunc
	treeViewGetSelection                   func(t *TreeView) *TreeSelection
	treeViewGetShowExpanders               func(t *TreeView) T.Gboolean
	treeViewGetTooltipColumn               func(t *TreeView) int
	treeViewGetTooltipContext              func(t *TreeView, x, y *int, keyboardTip T.Gboolean, model **TreeModel, path **TreePath, iter *TreeIter) T.Gboolean
	treeViewGetVadjustment                 func(t *TreeView) *Adjustment
	treeViewGetVisibleRange                func(t *TreeView, startPath **TreePath, endPath **TreePath) T.Gboolean
	treeViewGetVisibleRect                 func(t *TreeView, visibleRect *T.GdkRectangle)
	treeViewInsertColumn                   func(t *TreeView, column *TreeViewColumn, position int) int
	treeViewInsertColumnWithAttributes     func(t *TreeView, position int, title string, cell *CellRenderer, v ...VArg) int
	treeViewInsertColumnWithDataFunc       func(t *TreeView, position int, title string, cell *CellRenderer, f TreeCellDataFunc, data T.Gpointer, dnotify T.GDestroyNotify) int
	treeViewIsRubberBandingActive          func(t *TreeView) T.Gboolean
	treeViewMapExpandedRows                func(t *TreeView, f TreeViewMappingFunc, data T.Gpointer)
	treeViewMoveColumnAfter                func(t *TreeView, column *TreeViewColumn, baseColumn *TreeViewColumn)
	treeViewRemoveColumn                   func(t *TreeView, column *TreeViewColumn) int
	treeViewRowActivated                   func(t *TreeView, path *TreePath, column *TreeViewColumn)
	treeViewRowExpanded                    func(t *TreeView, path *TreePath) T.Gboolean
	treeViewScrollToCell                   func(t *TreeView, path *TreePath, column *TreeViewColumn, useAlign T.Gboolean, rowAlign, colAlign float32)
	treeViewScrollToPoint                  func(t *TreeView, treeX, treeY int)
	treeViewSetColumnDragFunction          func(t *TreeView, f TreeViewColumnDropFunc, userData T.Gpointer, destroy T.GDestroyNotify)
	treeViewSetCursor                      func(t *TreeView, path *TreePath, focusColumn *TreeViewColumn, startEditing T.Gboolean)
	treeViewSetCursorOnCell                func(t *TreeView, path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing T.Gboolean)
	treeViewSetDestroyCountFunc            func(t *TreeView, f TreeDestroyCountFunc, data T.Gpointer, destroy T.GDestroyNotify)
	treeViewSetDragDestRow                 func(t *TreeView, path *TreePath, pos TreeViewDropPosition)
	treeViewSetEnableSearch                func(t *TreeView, enableSearch T.Gboolean)
	treeViewSetEnableTreeLines             func(t *TreeView, enabled T.Gboolean)
	treeViewSetExpanderColumn              func(t *TreeView, column *TreeViewColumn)
	treeViewSetFixedHeightMode             func(t *TreeView, enable T.Gboolean)
	treeViewSetGridLines                   func(t *TreeView, gridLines TreeViewGridLines)
	treeViewSetHadjustment                 func(t *TreeView, adjustment *Adjustment)
	treeViewSetHeadersClickable            func(t *TreeView, setting T.Gboolean)
	treeViewSetHeadersVisible              func(t *TreeView, headersVisible T.Gboolean)
	treeViewSetHoverExpand                 func(t *TreeView, expand T.Gboolean)
	treeViewSetHoverSelection              func(t *TreeView, hover T.Gboolean)
	treeViewSetLevelIndentation            func(t *TreeView, indentation int)
	treeViewSetModel                       func(t *TreeView, model *TreeModel)
	treeViewSetReorderable                 func(t *TreeView, reorderable T.Gboolean)
	treeViewSetRowSeparatorFunc            func(t *TreeView, f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify)
	treeViewSetRubberBanding               func(t *TreeView, enable T.Gboolean)
	treeViewSetRulesHint                   func(t *TreeView, setting T.Gboolean)
	treeViewSetSearchColumn                func(t *TreeView, column int)
	treeViewSetSearchEntry                 func(t *TreeView, entry *Entry)
	treeViewSetSearchEqualFunc             func(t *TreeView, searchEqualFunc TreeViewSearchEqualFunc, searchUserData T.Gpointer, searchDestroy T.GDestroyNotify)
	treeViewSetSearchPositionFunc          func(t *TreeView, f TreeViewSearchPositionFunc, data T.Gpointer, destroy T.GDestroyNotify)
	treeViewSetShowExpanders               func(t *TreeView, enabled T.Gboolean)
	treeViewSetTooltipCell                 func(t *TreeView, tooltip *Tooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer)
	treeViewSetTooltipColumn               func(t *TreeView, column int)
	treeViewSetTooltipRow                  func(t *TreeView, tooltip *Tooltip, path *TreePath)
	treeViewSetVadjustment                 func(t *TreeView, adjustment *Adjustment)
	treeViewTreeToWidgetCoords             func(t *TreeView, tx, ty int, wx, wy *int)
	treeViewUnsetRowsDragDest              func(t *TreeView)
	treeViewUnsetRowsDragSource            func(t *TreeView)
	treeViewWidgetToTreeCoords             func(t *TreeView, wx, wy int, tx, ty *int)
)

func (t *TreeView) AppendColumn(column *TreeViewColumn) int { return treeViewAppendColumn(t, column) }
func (t *TreeView) CollapseAll()                            { treeViewCollapseAll(t) }
func (t *TreeView) CollapseRow(path *TreePath) T.Gboolean   { return treeViewCollapseRow(t, path) }
func (t *TreeView) ColumnsAutosize()                        { treeViewColumnsAutosize(t) }
func (t *TreeView) ConvertBinWindowToTreeCoords(bx, by int, tx, ty *int) {
	treeViewConvertBinWindowToTreeCoords(t, bx, by, tx, ty)
}
func (t *TreeView) ConvertBinWindowToWidgetCoords(bx, by int, wx, wy *int) {
	treeViewConvertBinWindowToWidgetCoords(t, bx, by, wx, wy)
}
func (t *TreeView) ConvertTreeToBinWindowCoords(tx, ty int, bx, by *int) {
	treeViewConvertTreeToBinWindowCoords(t, tx, ty, bx, by)
}
func (t *TreeView) ConvertTreeToWidgetCoords(tx, ty int, wx, wy *int) {
	treeViewConvertTreeToWidgetCoords(t, tx, ty, wx, wy)
}
func (t *TreeView) ConvertWidgetToBinWindowCoords(wx, wy int, bx, by *int) {
	treeViewConvertWidgetToBinWindowCoords(t, wx, wy, bx, by)
}
func (t *TreeView) ConvertWidgetToTreeCoords(wx, wy int, tx, ty *int) {
	treeViewConvertWidgetToTreeCoords(t, wx, wy, tx, ty)
}
func (t *TreeView) CreateRowDragIcon(path *TreePath) *T.GdkPixmap {
	return treeViewCreateRowDragIcon(t, path)
}
func (t *TreeView) EnableModelDragDest(targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction) {
	treeViewEnableModelDragDest(t, targets, nTargets, actions)
}
func (t *TreeView) EnableModelDragSource(startButtonMask T.GdkModifierType, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction) {
	treeViewEnableModelDragSource(t, startButtonMask, targets, nTargets, actions)
}
func (t *TreeView) ExpandAll() { treeViewExpandAll(t) }
func (t *TreeView) ExpandRow(path *TreePath, openAll T.Gboolean) T.Gboolean {
	return treeViewExpandRow(t, path, openAll)
}
func (t *TreeView) ExpandToPath(path *TreePath) { treeViewExpandToPath(t, path) }
func (t *TreeView) GetBackgroundArea(path *TreePath, column *TreeViewColumn, rect *T.GdkRectangle) {
	treeViewGetBackgroundArea(t, path, column, rect)
}
func (t *TreeView) GetBinWindow() *T.GdkWindow { return treeViewGetBinWindow(t) }
func (t *TreeView) GetCellArea(path *TreePath, column *TreeViewColumn, rect *T.GdkRectangle) {
	treeViewGetCellArea(t, path, column, rect)
}
func (t *TreeView) GetColumn(n int) *TreeViewColumn { return treeViewGetColumn(t, n) }
func (t *TreeView) GetColumns() *T.GList            { return treeViewGetColumns(t) }
func (t *TreeView) GetCursor(path **TreePath, focusColumn **TreeViewColumn) {
	treeViewGetCursor(t, path, focusColumn)
}
func (t *TreeView) GetDestRowAtPos(dragX, dragY int, path **TreePath, pos *TreeViewDropPosition) T.Gboolean {
	return treeViewGetDestRowAtPos(t, dragX, dragY, path, pos)
}
func (t *TreeView) GetDragDestRow(path **TreePath, pos *TreeViewDropPosition) {
	treeViewGetDragDestRow(t, path, pos)
}
func (t *TreeView) GetEnableSearch() T.Gboolean        { return treeViewGetEnableSearch(t) }
func (t *TreeView) GetEnableTreeLines() T.Gboolean     { return treeViewGetEnableTreeLines(t) }
func (t *TreeView) GetExpanderColumn() *TreeViewColumn { return treeViewGetExpanderColumn(t) }
func (t *TreeView) GetFixedHeightMode() T.Gboolean     { return treeViewGetFixedHeightMode(t) }
func (t *TreeView) GetGridLines() TreeViewGridLines    { return treeViewGetGridLines(t) }
func (t *TreeView) GetHadjustment() *Adjustment        { return treeViewGetHadjustment(t) }
func (t *TreeView) GetHeadersClickable() T.Gboolean    { return treeViewGetHeadersClickable(t) }
func (t *TreeView) GetHeadersVisible() T.Gboolean      { return treeViewGetHeadersVisible(t) }
func (t *TreeView) GetHoverExpand() T.Gboolean         { return treeViewGetHoverExpand(t) }
func (t *TreeView) GetHoverSelection() T.Gboolean      { return treeViewGetHoverSelection(t) }
func (t *TreeView) GetLevelIndentation() int           { return treeViewGetLevelIndentation(t) }
func (t *TreeView) GetModel() *TreeModel               { return treeViewGetModel(t) }
func (t *TreeView) GetPathAtPos(x, y int, path **TreePath, column **TreeViewColumn, cellX, cellY *int) T.Gboolean {
	return treeViewGetPathAtPos(t, x, y, path, column, cellX, cellY)
}
func (t *TreeView) GetReorderable() T.Gboolean { return treeViewGetReorderable(t) }
func (t *TreeView) GetRowSeparatorFunc() TreeViewRowSeparatorFunc {
	return treeViewGetRowSeparatorFunc(t)
}
func (t *TreeView) GetRubberBanding() T.Gboolean { return treeViewGetRubberBanding(t) }
func (t *TreeView) GetRulesHint() T.Gboolean     { return treeViewGetRulesHint(t) }
func (t *TreeView) GetSearchColumn() int         { return treeViewGetSearchColumn(t) }
func (t *TreeView) GetSearchEntry() *Entry       { return treeViewGetSearchEntry(t) }
func (t *TreeView) GetSearchEqualFunc() TreeViewSearchEqualFunc {
	return treeViewGetSearchEqualFunc(t)
}
func (t *TreeView) GetSearchPositionFunc() TreeViewSearchPositionFunc {
	return treeViewGetSearchPositionFunc(t)
}
func (t *TreeView) GetSelection() *TreeSelection { return treeViewGetSelection(t) }
func (t *TreeView) GetShowExpanders() T.Gboolean { return treeViewGetShowExpanders(t) }
func (t *TreeView) GetTooltipColumn() int        { return treeViewGetTooltipColumn(t) }
func (t *TreeView) GetTooltipContext(x, y *int, keyboardTip T.Gboolean, model **TreeModel, path **TreePath, iter *TreeIter) T.Gboolean {
	return treeViewGetTooltipContext(t, x, y, keyboardTip, model, path, iter)
}
func (t *TreeView) GetVadjustment() *Adjustment { return treeViewGetVadjustment(t) }
func (t *TreeView) GetVisibleRange(startPath **TreePath, endPath **TreePath) T.Gboolean {
	return treeViewGetVisibleRange(t, startPath, endPath)
}
func (t *TreeView) GetVisibleRect(visibleRect *T.GdkRectangle) {
	treeViewGetVisibleRect(t, visibleRect)
}
func (t *TreeView) InsertColumn(column *TreeViewColumn, position int) int {
	return treeViewInsertColumn(t, column, position)
}
func (t *TreeView) InsertColumnWithAttributes(position int, title string, cell *CellRenderer, v ...VArg) int {
	return treeViewInsertColumnWithAttributes(t, position, title, cell, v)
}
func (t *TreeView) InsertColumnWithDataFunc(position int, title string, cell *CellRenderer, f TreeCellDataFunc, data T.Gpointer, dnotify T.GDestroyNotify) int {
	return treeViewInsertColumnWithDataFunc(t, position, title, cell, f, data, dnotify)
}
func (t *TreeView) IsRubberBandingActive() T.Gboolean { return treeViewIsRubberBandingActive(t) }
func (t *TreeView) MapExpandedRows(f TreeViewMappingFunc, data T.Gpointer) {
	treeViewMapExpandedRows(t, f, data)
}
func (t *TreeView) MoveColumnAfter(column *TreeViewColumn, baseColumn *TreeViewColumn) {
	treeViewMoveColumnAfter(t, column, baseColumn)
}
func (t *TreeView) RemoveColumn(column *TreeViewColumn) int { return treeViewRemoveColumn(t, column) }
func (t *TreeView) RowActivated(path *TreePath, column *TreeViewColumn) {
	treeViewRowActivated(t, path, column)
}
func (t *TreeView) RowExpanded(path *TreePath) T.Gboolean { return treeViewRowExpanded(t, path) }
func (t *TreeView) ScrollToCell(path *TreePath, column *TreeViewColumn, useAlign T.Gboolean, rowAlign, colAlign float32) {
	treeViewScrollToCell(t, path, column, useAlign, rowAlign, colAlign)
}
func (t *TreeView) ScrollToPoint(treeX, treeY int) { treeViewScrollToPoint(t, treeX, treeY) }
func (t *TreeView) SetColumnDragFunction(f TreeViewColumnDropFunc, userData T.Gpointer, destroy T.GDestroyNotify) {
	treeViewSetColumnDragFunction(t, f, userData, destroy)
}
func (t *TreeView) SetCursor(path *TreePath, focusColumn *TreeViewColumn, startEditing T.Gboolean) {
	treeViewSetCursor(t, path, focusColumn, startEditing)
}
func (t *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing T.Gboolean) {
	treeViewSetCursorOnCell(t, path, focusColumn, focusCell, startEditing)
}
func (t *TreeView) SetDestroyCountFunc(f TreeDestroyCountFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeViewSetDestroyCountFunc(t, f, data, destroy)
}
func (t *TreeView) SetDragDestRow(path *TreePath, pos TreeViewDropPosition) {
	treeViewSetDragDestRow(t, path, pos)
}
func (t *TreeView) SetEnableSearch(enableSearch T.Gboolean)  { treeViewSetEnableSearch(t, enableSearch) }
func (t *TreeView) SetEnableTreeLines(enabled T.Gboolean)    { treeViewSetEnableTreeLines(t, enabled) }
func (t *TreeView) SetExpanderColumn(column *TreeViewColumn) { treeViewSetExpanderColumn(t, column) }
func (t *TreeView) SetFixedHeightMode(enable T.Gboolean)     { treeViewSetFixedHeightMode(t, enable) }
func (t *TreeView) SetGridLines(gridLines TreeViewGridLines) { treeViewSetGridLines(t, gridLines) }
func (t *TreeView) SetHadjustment(adjustment *Adjustment)    { treeViewSetHadjustment(t, adjustment) }
func (t *TreeView) SetHeadersClickable(setting T.Gboolean)   { treeViewSetHeadersClickable(t, setting) }
func (t *TreeView) SetHeadersVisible(headersVisible T.Gboolean) {
	treeViewSetHeadersVisible(t, headersVisible)
}
func (t *TreeView) SetHoverExpand(expand T.Gboolean)      { treeViewSetHoverExpand(t, expand) }
func (t *TreeView) SetHoverSelection(hover T.Gboolean)    { treeViewSetHoverSelection(t, hover) }
func (t *TreeView) SetLevelIndentation(indentation int)   { treeViewSetLevelIndentation(t, indentation) }
func (t *TreeView) SetModel(model *TreeModel)             { treeViewSetModel(t, model) }
func (t *TreeView) SetReorderable(reorderable T.Gboolean) { treeViewSetReorderable(t, reorderable) }
func (t *TreeView) SetRowSeparatorFunc(f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeViewSetRowSeparatorFunc(t, f, data, destroy)
}
func (t *TreeView) SetRubberBanding(enable T.Gboolean) { treeViewSetRubberBanding(t, enable) }
func (t *TreeView) SetRulesHint(setting T.Gboolean)    { treeViewSetRulesHint(t, setting) }
func (t *TreeView) SetSearchColumn(column int)         { treeViewSetSearchColumn(t, column) }
func (t *TreeView) SetSearchEntry(entry *Entry)        { treeViewSetSearchEntry(t, entry) }
func (t *TreeView) SetSearchEqualFunc(searchEqualFunc TreeViewSearchEqualFunc, searchUserData T.Gpointer, searchDestroy T.GDestroyNotify) {
	treeViewSetSearchEqualFunc(t, searchEqualFunc, searchUserData, searchDestroy)
}
func (t *TreeView) SetSearchPositionFunc(f TreeViewSearchPositionFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeViewSetSearchPositionFunc(t, f, data, destroy)
}
func (t *TreeView) SetShowExpanders(enabled T.Gboolean) { treeViewSetShowExpanders(t, enabled) }
func (t *TreeView) SetTooltipCell(tooltip *Tooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer) {
	treeViewSetTooltipCell(t, tooltip, path, column, cell)
}
func (t *TreeView) SetTooltipColumn(column int) { treeViewSetTooltipColumn(t, column) }
func (t *TreeView) SetTooltipRow(tooltip *Tooltip, path *TreePath) {
	treeViewSetTooltipRow(t, tooltip, path)
}
func (t *TreeView) SetVadjustment(adjustment *Adjustment) { treeViewSetVadjustment(t, adjustment) }
func (t *TreeView) TreeToWidgetCoords(tx, ty int, wx, wy *int) {
	treeViewTreeToWidgetCoords(t, tx, ty, wx, wy)
}
func (t *TreeView) UnsetRowsDragDest()   { treeViewUnsetRowsDragDest(t) }
func (t *TreeView) UnsetRowsDragSource() { treeViewUnsetRowsDragSource(t) }
func (t *TreeView) WidgetToTreeCoords(wx, wy int, tx, ty *int) {
	treeViewWidgetToTreeCoords(t, wx, wy, tx, ty)
}

type TreeViewColumn struct {
	Parent                  T.GtkObject
	TreeView                *T.GtkWidget
	Button                  *T.GtkWidget
	Child                   *T.GtkWidget
	Arrow                   *T.GtkWidget
	Alignment               *T.GtkWidget
	Window                  *T.GdkWindow
	EditableWidget          *CellEditable
	Xalign                  float32
	PropertyChangedSignal   uint
	Spacing                 int
	ColumnType              TreeViewColumnSizing
	RequestedWidth          int
	ButtonRequest           int
	ResizedWidth            int
	Width                   int
	FixedWidth              int
	MinWidth                int
	MaxWidth                int
	DragX                   int
	DragY                   int
	Title                   *T.Gchar
	CellList                *T.GList
	SortClickedSignal       uint
	SortColumnChangedSignal uint
	SortColumnId            int
	SortOrder               T.GtkSortType
	Bits                    uint
	// Visible : 1
	// Resizable : 1
	// Clickable : 1
	// Dirty : 1
	// ShowSortIndicator : 1
	// MaybeReordered : 1
	// Reorderable : 1
	// UseResizedWidth : 1
	// Expand : 1
}

var (
	TreeViewColumnGetType           func() T.GType
	TreeViewColumnNew               func() *TreeViewColumn
	TreeViewColumnNewWithAttributes func(title string, cell *CellRenderer, v ...VArg) *TreeViewColumn

	TreeViewColumnSizingGetType func() T.GType

	treeViewColumnAddAttribute     func(t *TreeViewColumn, cellRenderer *CellRenderer, attribute string, column int)
	treeViewColumnCellGetPosition  func(t *TreeViewColumn, cellRenderer *CellRenderer, startPos, width *int) T.Gboolean
	treeViewColumnCellGetSize      func(t *TreeViewColumn, cellArea *T.GdkRectangle, xOffset, yOffset, width, height *int)
	treeViewColumnCellIsVisible    func(t *TreeViewColumn) T.Gboolean
	treeViewColumnCellSetCellData  func(t *TreeViewColumn, treeModel *TreeModel, iter *TreeIter, isExpander, isExpanded T.Gboolean)
	treeViewColumnClear            func(t *TreeViewColumn)
	treeViewColumnClearAttributes  func(t *TreeViewColumn, cellRenderer *CellRenderer)
	treeViewColumnClicked          func(t *TreeViewColumn)
	treeViewColumnFocusCell        func(t *TreeViewColumn, cell *CellRenderer)
	treeViewColumnGetAlignment     func(t *TreeViewColumn) float32
	treeViewColumnGetCellRenderers func(t *TreeViewColumn) *T.GList
	treeViewColumnGetClickable     func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetExpand        func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetFixedWidth    func(t *TreeViewColumn) int
	treeViewColumnGetMaxWidth      func(t *TreeViewColumn) int
	treeViewColumnGetMinWidth      func(t *TreeViewColumn) int
	treeViewColumnGetReorderable   func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetResizable     func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetSizing        func(t *TreeViewColumn) TreeViewColumnSizing
	treeViewColumnGetSortColumnId  func(t *TreeViewColumn) int
	treeViewColumnGetSortIndicator func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetSortOrder     func(t *TreeViewColumn) T.GtkSortType
	treeViewColumnGetSpacing       func(t *TreeViewColumn) int
	treeViewColumnGetTitle         func(t *TreeViewColumn) string
	treeViewColumnGetTreeView      func(t *TreeViewColumn) *Widget
	treeViewColumnGetVisible       func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetWidget        func(t *TreeViewColumn) *Widget
	treeViewColumnGetWidth         func(t *TreeViewColumn) int
	treeViewColumnPackEnd          func(t *TreeViewColumn, cell *CellRenderer, expand T.Gboolean)
	treeViewColumnPackStart        func(t *TreeViewColumn, cell *CellRenderer, expand T.Gboolean)
	treeViewColumnQueueResize      func(t *TreeViewColumn)
	treeViewColumnSetAlignment     func(t *TreeViewColumn, xalign float32)
	treeViewColumnSetAttributes    func(t *TreeViewColumn, cellRenderer *CellRenderer, v ...VArg)
	treeViewColumnSetCellDataFunc  func(t *TreeViewColumn, cellRenderer *CellRenderer, f TreeCellDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
	treeViewColumnSetClickable     func(t *TreeViewColumn, clickable T.Gboolean)
	treeViewColumnSetExpand        func(t *TreeViewColumn, expand T.Gboolean)
	treeViewColumnSetFixedWidth    func(t *TreeViewColumn, fixedWidth int)
	treeViewColumnSetMaxWidth      func(t *TreeViewColumn, maxWidth int)
	treeViewColumnSetMinWidth      func(t *TreeViewColumn, minWidth int)
	treeViewColumnSetReorderable   func(t *TreeViewColumn, reorderable T.Gboolean)
	treeViewColumnSetResizable     func(t *TreeViewColumn, resizable T.Gboolean)
	treeViewColumnSetSizing        func(t *TreeViewColumn, typ TreeViewColumnSizing)
	treeViewColumnSetSortColumnId  func(t *TreeViewColumn, sortColumnId int)
	treeViewColumnSetSortIndicator func(t *TreeViewColumn, setting T.Gboolean)
	treeViewColumnSetSortOrder     func(t *TreeViewColumn, order T.GtkSortType)
	treeViewColumnSetSpacing       func(t *TreeViewColumn, spacing int)
	treeViewColumnSetTitle         func(t *TreeViewColumn, title string)
	treeViewColumnSetVisible       func(t *TreeViewColumn, visible T.Gboolean)
	treeViewColumnSetWidget        func(t *TreeViewColumn, widget *Widget)
)

func (t *TreeViewColumn) AddAttribute(cellRenderer *CellRenderer, attribute string, column int) {
	treeViewColumnAddAttribute(t, cellRenderer, attribute, column)
}
func (t *TreeViewColumn) CellGetPosition(cellRenderer *CellRenderer, startPos *int, width *int) T.Gboolean {
	return treeViewColumnCellGetPosition(t, cellRenderer, startPos, width)
}
func (t *TreeViewColumn) CellGetSize(cellArea *T.GdkRectangle, xOffset, yOffset, width, height *int) {
	treeViewColumnCellGetSize(t, cellArea, xOffset, yOffset, width, height)
}
func (t *TreeViewColumn) CellIsVisible() T.Gboolean { return treeViewColumnCellIsVisible(t) }
func (t *TreeViewColumn) CellSetCellData(treeModel *TreeModel, iter *TreeIter, isExpander, isExpanded T.Gboolean) {
	treeViewColumnCellSetCellData(t, treeModel, iter, isExpander, isExpanded)
}
func (t *TreeViewColumn) Clear() { treeViewColumnClear(t) }
func (t *TreeViewColumn) ClearAttributes(cellRenderer *CellRenderer) {
	treeViewColumnClearAttributes(t, cellRenderer)
}
func (t *TreeViewColumn) Clicked()                        { treeViewColumnClicked(t) }
func (t *TreeViewColumn) FocusCell(cell *CellRenderer)    { treeViewColumnFocusCell(t, cell) }
func (t *TreeViewColumn) GetAlignment() float32           { return treeViewColumnGetAlignment(t) }
func (t *TreeViewColumn) GetCellRenderers() *T.GList      { return treeViewColumnGetCellRenderers(t) }
func (t *TreeViewColumn) GetClickable() T.Gboolean        { return treeViewColumnGetClickable(t) }
func (t *TreeViewColumn) GetExpand() T.Gboolean           { return treeViewColumnGetExpand(t) }
func (t *TreeViewColumn) GetFixedWidth() int              { return treeViewColumnGetFixedWidth(t) }
func (t *TreeViewColumn) GetMaxWidth() int                { return treeViewColumnGetMaxWidth(t) }
func (t *TreeViewColumn) GetMinWidth() int                { return treeViewColumnGetMinWidth(t) }
func (t *TreeViewColumn) GetReorderable() T.Gboolean      { return treeViewColumnGetReorderable(t) }
func (t *TreeViewColumn) GetResizable() T.Gboolean        { return treeViewColumnGetResizable(t) }
func (t *TreeViewColumn) GetSizing() TreeViewColumnSizing { return treeViewColumnGetSizing(t) }
func (t *TreeViewColumn) GetSortColumnId() int            { return treeViewColumnGetSortColumnId(t) }
func (t *TreeViewColumn) GetSortIndicator() T.Gboolean    { return treeViewColumnGetSortIndicator(t) }
func (t *TreeViewColumn) GetSortOrder() T.GtkSortType     { return treeViewColumnGetSortOrder(t) }
func (t *TreeViewColumn) GetSpacing() int                 { return treeViewColumnGetSpacing(t) }
func (t *TreeViewColumn) GetTitle() string                { return treeViewColumnGetTitle(t) }
func (t *TreeViewColumn) GetTreeView() *Widget            { return treeViewColumnGetTreeView(t) }
func (t *TreeViewColumn) GetVisible() T.Gboolean          { return treeViewColumnGetVisible(t) }
func (t *TreeViewColumn) GetWidget() *Widget              { return treeViewColumnGetWidget(t) }
func (t *TreeViewColumn) GetWidth() int                   { return treeViewColumnGetWidth(t) }
func (t *TreeViewColumn) PackStart(cell *CellRenderer, expand T.Gboolean) {
	treeViewColumnPackEnd(t, cell, expand)
}
func (t *TreeViewColumn) QueueResize()                { treeViewColumnQueueResize(t) }
func (t *TreeViewColumn) SetAlignment(xalign float32) { treeViewColumnSetAlignment(t, xalign) }
func (t *TreeViewColumn) SetAttributes(cellRenderer *CellRenderer, v ...VArg) {
	treeViewColumnSetAttributes(t, cellRenderer, v)
}
func (t *TreeViewColumn) SetCellDataFunc(cellRenderer *CellRenderer, f TreeCellDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	treeViewColumnSetCellDataFunc(t, cellRenderer, f, funcData, destroy)
}
func (t *TreeViewColumn) SetClickable(clickable T.Gboolean) { treeViewColumnSetClickable(t, clickable) }
func (t *TreeViewColumn) SetExpand(expand T.Gboolean)       { treeViewColumnSetExpand(t, expand) }
func (t *TreeViewColumn) SetFixedWidth(fixedWidth int)      { treeViewColumnSetFixedWidth(t, fixedWidth) }
func (t *TreeViewColumn) SetMaxWidth(maxWidth int)          { treeViewColumnSetMaxWidth(t, maxWidth) }
func (t *TreeViewColumn) SetMinWidth(minWidth int)          { treeViewColumnSetMinWidth(t, minWidth) }
func (t *TreeViewColumn) SetReorderable(reorderable T.Gboolean) {
	treeViewColumnSetReorderable(t, reorderable)
}
func (t *TreeViewColumn) SetResizable(resizable T.Gboolean)  { treeViewColumnSetResizable(t, resizable) }
func (t *TreeViewColumn) SetSizing(typ TreeViewColumnSizing) { treeViewColumnSetSizing(t, typ) }
func (t *TreeViewColumn) SetSortColumnId(sortColumnId int) {
	treeViewColumnSetSortColumnId(t, sortColumnId)
}
func (t *TreeViewColumn) SetSortIndicator(setting T.Gboolean) {
	treeViewColumnSetSortIndicator(t, setting)
}
func (t *TreeViewColumn) SetSortOrder(order T.GtkSortType) { treeViewColumnSetSortOrder(t, order) }
func (t *TreeViewColumn) SetSpacing(spacing int)           { treeViewColumnSetSpacing(t, spacing) }
func (t *TreeViewColumn) SetTitle(title string)            { treeViewColumnSetTitle(t, title) }
func (t *TreeViewColumn) SetVisible(visible T.Gboolean)    { treeViewColumnSetVisible(t, visible) }
func (t *TreeViewColumn) SetWidget(widget *Widget)         { treeViewColumnSetWidget(t, widget) }
