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

type (
	Table struct {
		Container     Container
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
	TableGetType func() O.Type
	TableNew     func(rows, columns uint, homogeneous bool) *Widget

	TableAttach               func(t *Table, child *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint, xoptions, yoptions AttachOptions, xpadding, ypadding uint)
	TableAttachDefaults       func(t *Table, widget *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint)
	TableGetColSpacing        func(t *Table, column uint) uint
	TableGetDefaultColSpacing func(t *Table) uint
	TableGetDefaultRowSpacing func(t *Table) uint
	TableGetHomogeneous       func(t *Table) bool
	TableGetRowSpacing        func(t *Table, row uint) uint
	TableGetSize              func(t *Table, rows, columns *uint)
	TableResize               func(t *Table, rows, columns uint)
	TableSetColSpacing        func(t *Table, column, spacing uint)
	TableSetColSpacings       func(t *Table, spacing uint)
	TableSetHomogeneous       func(t *Table, homogeneous bool)
	TableSetRowSpacing        func(t *Table, row, spacing uint)
	TableSetRowSpacings       func(t *Table, spacing uint)
)

func (t *Table) Attach(child *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint, xoptions, yoptions AttachOptions, xpadding, ypadding uint) {
	TableAttach(t, child, leftAttach, rightAttach, topAttach, bottomAttach, xoptions, yoptions, xpadding, ypadding)
}
func (t *Table) AttachDefaults(widget *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint) {
	TableAttachDefaults(t, widget, leftAttach, rightAttach, topAttach, bottomAttach)
}
func (t *Table) GetColSpacing(column uint) uint     { return TableGetColSpacing(t, column) }
func (t *Table) GetDefaultColSpacing() uint         { return TableGetDefaultColSpacing(t) }
func (t *Table) GetDefaultRowSpacing() uint         { return TableGetDefaultRowSpacing(t) }
func (t *Table) GetHomogeneous() bool               { return TableGetHomogeneous(t) }
func (t *Table) GetRowSpacing(row uint) uint        { return TableGetRowSpacing(t, row) }
func (t *Table) GetSize(rows, columns *uint)        { TableGetSize(t, rows, columns) }
func (t *Table) Resize(rows, columns uint)          { TableResize(t, rows, columns) }
func (t *Table) SetColSpacing(column, spacing uint) { TableSetColSpacing(t, column, spacing) }
func (t *Table) SetColSpacings(spacing uint)        { TableSetColSpacings(t, spacing) }
func (t *Table) SetHomogeneous(homogeneous bool)    { TableSetHomogeneous(t, homogeneous) }
func (t *Table) SetRowSpacing(row, spacing uint)    { TableSetRowSpacing(t, row, spacing) }
func (t *Table) SetRowSpacings(spacing uint)        { TableSetRowSpacings(t, spacing) }

type TargetEntry struct {
	Target *T.Gchar
	Flags  uint
	Info   uint
}

type TargetFlags Enum

const (
	TARGET_SAME_APP TargetFlags = 1 << iota
	TARGET_SAME_WIDGET
	TARGET_OTHER_APP
	TARGET_OTHER_WIDGET
)

var TargetFlagsGetType func() O.Type

type TargetList struct {
	List     *T.GList
	RefCount uint
}

var (
	TargetListGetType func() O.Type
	TargetListNew     func(targets *TargetEntry, ntargets uint) *TargetList

	TargetTableFree        func(targets *TargetEntry, nTargets int)
	TargetTableNewFromList func(tl *TargetList, nTargets *int) *TargetEntry

	TargetListRef                func(t *TargetList) *TargetList
	TargetListUnref              func(t *TargetList)
	TargetListAdd                func(t *TargetList, target D.Atom, flags, info uint)
	TargetListAddTextTargets     func(t *TargetList, info uint)
	TargetListAddRichTextTargets func(t *TargetList, info uint, deserializable bool, buffer *TextBuffer)
	TargetListAddImageTargets    func(t *TargetList, info uint, writable bool)
	TargetListAddUriTargets      func(t *TargetList, info uint)
	TargetListAddTable           func(t *TargetList, targets *TargetEntry, ntargets uint)
	TargetListRemove             func(t *TargetList, target D.Atom)
	TargetListFind               func(t *TargetList, target D.Atom, info *uint) bool
)

func (t *TargetList) Ref() *TargetList                    { return TargetListRef(t) }
func (t *TargetList) Unref()                              { TargetListUnref(t) }
func (t *TargetList) Add(target D.Atom, flags, info uint) { TargetListAdd(t, target, flags, info) }
func (t *TargetList) AddTextTargets(info uint)            { TargetListAddTextTargets(t, info) }
func (t *TargetList) AddRichTextTargets(info uint, deserializable bool, buffer *TextBuffer) {
	TargetListAddRichTextTargets(t, info, deserializable, buffer)
}
func (t *TargetList) AddImageTargets(info uint, writable bool) {
	TargetListAddImageTargets(t, info, writable)
}
func (t *TargetList) AddUriTargets(info uint) { TargetListAddUriTargets(t, info) }
func (t *TargetList) AddTable(targets *TargetEntry, ntargets uint) {
	TargetListAddTable(t, targets, ntargets)
}
func (t *TargetList) Remove(target D.Atom) { TargetListRemove(t, target) }
func (t *TargetList) Find(target D.Atom, info *uint) bool {
	return TargetListFind(t, target, info)
}

var (
	TargetsIncludeText     func(targets *D.Atom, nTargets int) bool
	TargetsIncludeRichText func(targets *D.Atom, nTargets int, buffer *TextBuffer) bool
	TargetsIncludeImage    func(targets *D.Atom, nTargets int, writable bool) bool
	TargetsIncludeUri      func(targets *D.Atom, nTargets int) bool
)

var (
	TearoffMenuItemGetType func() O.Type  //TODO(t):Use?
	TearoffMenuItemNew     func() *Widget //TODO(t):Use?
)

type TextAppearance struct {
	BgColor   D.Color
	FgColor   D.Color
	BgStipple *T.GdkBitmap
	FgStipple *T.GdkBitmap
	Rise      int
	Padding1  T.Gpointer
	Bits      uint
	// Underline : 4
	// Strikethrough : 1
	// DrawBg : 1
	// InsideSelection : 1
	// IsText : 1
	// Pad1 : 1
	// Pad2 : 1
	// Pad3 : 1
	// Pad4 : 1
}

type TextAttributes struct {
	Refcount         uint
	Appearance       TextAppearance
	Justification    Justification
	Direction        TextDirection
	Font             *P.FontDescription
	FontScale        float64
	LeftMargin       int
	Indent           int
	RightMargin      int
	PixelsAboveLines int
	PixelsBelowLines int
	PixelsInsideWrap int
	Tabs             *P.TabArray
	WrapMode         WrapMode
	Language         *P.Language
	PgBgColor        *D.Color
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
	TextAttributesGetType func() O.Type
	TextAttributesNew     func() *TextAttributes

	TextAttributesCopy       func(t *TextAttributes) *TextAttributes
	TextAttributesCopyValues func(t *TextAttributes, dest *TextAttributes)
	TextAttributesRef        func(t *TextAttributes) *TextAttributes
	TextAttributesUnref      func(t *TextAttributes)
)

func (t *TextAttributes) Copy() *TextAttributes           { return TextAttributesCopy(t) }
func (t *TextAttributes) CopyValues(dest *TextAttributes) { TextAttributesCopyValues(t, dest) }
func (t *TextAttributes) Ref() *TextAttributes            { return TextAttributesRef(t) }
func (t *TextAttributes) Unref()                          { TextAttributesUnref(t) }

type TextBTree struct{}

type (
	TextBuffer struct {
		Parent                   O.Object
		TagTable                 *TextTagTable
		Btree                    *TextBTree
		ClipboardContentsBuffers *L.SList
		SelectionClipboards      *L.SList
		LogAttrCache             *TextLogAttrCache
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
		createTags bool, userData T.Gpointer,
		e **T.GError) bool
)

var (
	TextBufferGetType           func() O.Type
	TextBufferNew               func(table *TextTagTable) *TextBuffer
	TextBufferTargetInfoGetType func() O.Type

	TextBufferAddMark                     func(t *TextBuffer, mark *TextMark, where *TextIter)
	TextBufferAddSelectionClipboard       func(t *TextBuffer, clipboard *Clipboard)
	TextBufferApplyTag                    func(t *TextBuffer, tag *TextTag, start, end *TextIter)
	TextBufferApplyTagByName              func(t *TextBuffer, name string, start, end *TextIter)
	TextBufferBackspace                   func(t *TextBuffer, iter *TextIter, interactive, defaultEditable bool) bool
	TextBufferBeginUserAction             func(t *TextBuffer)
	TextBufferCopyClipboard               func(t *TextBuffer, clipboard *Clipboard)
	TextBufferCreateChildAnchor           func(t *TextBuffer, iter *TextIter) *TextChildAnchor
	TextBufferCreateMark                  func(t *TextBuffer, markName string, where *TextIter, leftGravity bool) *TextMark
	TextBufferCreateTag                   func(t *TextBuffer, tagName, firstPropertyName string, v ...VArg) *TextTag
	TextBufferCutClipboard                func(t *TextBuffer, clipboard *Clipboard, defaultEditable bool)
	TextBufferDelete                      func(t *TextBuffer, start, end *TextIter)
	TextBufferDeleteInteractive           func(t *TextBuffer, startIter, endIter *TextIter, defaultEditable bool) bool
	TextBufferDeleteMark                  func(t *TextBuffer, mark *TextMark)
	TextBufferDeleteMarkByName            func(t *TextBuffer, name string)
	TextBufferDeleteSelection             func(t *TextBuffer, interactive, defaultEditable bool) bool
	TextBufferDeserialize                 func(t *TextBuffer, contentBuffer *TextBuffer, format D.Atom, iter *TextIter, data *uint8, length T.Gsize, error **T.GError) bool
	TextBufferDeserializeGetCanCreateTags func(t *TextBuffer, format D.Atom) bool
	TextBufferDeserializeSetCanCreateTags func(t *TextBuffer, format D.Atom, canCreateTags bool)
	TextBufferEndUserAction               func(t *TextBuffer)
	TextBufferGetBounds                   func(t *TextBuffer, start, end *TextIter)
	TextBufferGetCharCount                func(t *TextBuffer) int
	TextBufferGetCopyTargetList           func(t *TextBuffer) *TargetList
	TextBufferGetDeserializeFormats       func(t *TextBuffer, nFormats *int) *D.Atom
	TextBufferGetEndIter                  func(t *TextBuffer, iter *TextIter)
	TextBufferGetHasSelection             func(t *TextBuffer) bool
	TextBufferGetInsert                   func(t *TextBuffer) *TextMark
	TextBufferGetIterAtChildAnchor        func(t *TextBuffer, iter *TextIter, anchor *TextChildAnchor)
	TextBufferGetIterAtLine               func(t *TextBuffer, iter *TextIter, lineNumber int)
	TextBufferGetIterAtLineIndex          func(t *TextBuffer, iter *TextIter, lineNumber, byteIndex int)
	TextBufferGetIterAtLineOffset         func(t *TextBuffer, iter *TextIter, lineNumber, charOffset int)
	TextBufferGetIterAtMark               func(t *TextBuffer, iter *TextIter, mark *TextMark)
	TextBufferGetIterAtOffset             func(t *TextBuffer, iter *TextIter, charOffset int)
	TextBufferGetLineCount                func(t *TextBuffer) int
	TextBufferGetMark                     func(t *TextBuffer, name string) *TextMark
	TextBufferGetModified                 func(t *TextBuffer) bool
	TextBufferGetPasteTargetList          func(t *TextBuffer) *TargetList
	TextBufferGetSelectionBound           func(t *TextBuffer) *TextMark
	TextBufferGetSelectionBounds          func(t *TextBuffer, start, end *TextIter) bool
	TextBufferGetSerializeFormats         func(t *TextBuffer, nFormats *int) *D.Atom
	TextBufferGetSlice                    func(t *TextBuffer, start, end *TextIter, includeHiddenChars bool) string
	TextBufferGetStartIter                func(t *TextBuffer, iter *TextIter)
	TextBufferGetTagTable                 func(t *TextBuffer) *TextTagTable
	TextBufferGetText                     func(t *TextBuffer, start, end *TextIter, includeHiddenChars bool) string
	TextBufferInsert                      func(t *TextBuffer, iter *TextIter, text string, leng int)
	TextBufferInsertAtCursor              func(t *TextBuffer, text string, leng int)
	TextBufferInsertChildAnchor           func(t *TextBuffer, iter *TextIter, anchor *TextChildAnchor)
	TextBufferInsertInteractive           func(t *TextBuffer, iter *TextIter, text string, leng int, defaultEditable bool) bool
	TextBufferInsertInteractiveAtCursor   func(t *TextBuffer, text string, leng int, defaultEditable bool) bool
	TextBufferInsertPixbuf                func(t *TextBuffer, iter *TextIter, pixbuf *D.Pixbuf)
	TextBufferInsertRange                 func(t *TextBuffer, iter, start, end *TextIter)
	TextBufferInsertRangeInteractive      func(t *TextBuffer, iter, start, end *TextIter, defaultEditable bool) bool
	TextBufferInsertWithTags              func(t *TextBuffer, iter *TextIter, text string, leng int, firstTag *TextTag, v ...VArg)
	TextBufferInsertWithTagsByName        func(t *TextBuffer, iter *TextIter, text string, leng int, firstTagName string, v ...VArg)
	TextBufferMoveMark                    func(t *TextBuffer, mark *TextMark, where *TextIter)
	TextBufferMoveMarkByName              func(t *TextBuffer, name string, where *TextIter)
	TextBufferPasteClipboard              func(t *TextBuffer, clipboard *Clipboard, overrideLocation *TextIter, defaultEditable bool)
	TextBufferPlaceCursor                 func(t *TextBuffer, where *TextIter)
	TextBufferRegisterDeserializeFormat   func(t *TextBuffer, mimeType string, function TextBufferDeserializeFunc, userData T.Gpointer, userDataDestroy T.GDestroyNotify) D.Atom
	TextBufferRegisterDeserializeTagset   func(t *TextBuffer, tagsetName string) D.Atom
	TextBufferRegisterSerializeFormat     func(t *TextBuffer, mimeType string, function TextBufferSerializeFunc, userData T.Gpointer, userDataDestroy T.GDestroyNotify) D.Atom
	TextBufferRegisterSerializeTagset     func(t *TextBuffer, tagsetName string) D.Atom
	TextBufferRemoveAllTags               func(t *TextBuffer, start, end *TextIter)
	TextBufferRemoveSelectionClipboard    func(t *TextBuffer, clipboard *Clipboard)
	TextBufferRemoveTag                   func(t *TextBuffer, tag *TextTag, start, end *TextIter)
	TextBufferRemoveTagByName             func(t *TextBuffer, name string, start, end *TextIter)
	TextBufferSelectRange                 func(t *TextBuffer, ins, bound *TextIter)
	TextBufferSerialize                   func(t *TextBuffer, contentBuffer *TextBuffer, format D.Atom, start, end *TextIter, length *T.Gsize) *uint8
	TextBufferSetModified                 func(t *TextBuffer, setting bool)
	TextBufferSetText                     func(t *TextBuffer, text string, leng int)
	TextBufferUnregisterDeserializeFormat func(t *TextBuffer, format D.Atom)
	TextBufferUnregisterSerializeFormat   func(t *TextBuffer, format D.Atom)
)

func (t *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	TextBufferAddMark(t, mark, where)
}
func (t *TextBuffer) AddSelectionClipboard(clipboard *Clipboard) {
	TextBufferAddSelectionClipboard(t, clipboard)
}
func (t *TextBuffer) ApplyTag(tag *TextTag, start, end *TextIter) {
	TextBufferApplyTag(t, tag, start, end)
}
func (t *TextBuffer) ApplyTagByName(name string, start, end *TextIter) {
	TextBufferApplyTagByName(t, name, start, end)
}
func (t *TextBuffer) Backspace(iter *TextIter, interactive, defaultEditable bool) bool {
	return TextBufferBackspace(t, iter, interactive, defaultEditable)
}
func (t *TextBuffer) BeginUserAction()                   { TextBufferBeginUserAction(t) }
func (t *TextBuffer) CopyClipboard(clipboard *Clipboard) { TextBufferCopyClipboard(t, clipboard) }
func (t *TextBuffer) CreateChildAnchor(iter *TextIter) *TextChildAnchor {
	return TextBufferCreateChildAnchor(t, iter)
}
func (t *TextBuffer) CreateMark(markName string, where *TextIter, leftGravity bool) *TextMark {
	return TextBufferCreateMark(t, markName, where, leftGravity)
}
func (t *TextBuffer) CreateTag(tagName string, firstPropertyName string, v ...VArg) *TextTag {
	return TextBufferCreateTag(t, tagName, firstPropertyName, v)
}
func (t *TextBuffer) CutClipboard(clipboard *Clipboard, defaultEditable bool) {
	TextBufferCutClipboard(t, clipboard, defaultEditable)
}
func (t *TextBuffer) Delete(start, end *TextIter) { TextBufferDelete(t, start, end) }
func (t *TextBuffer) DeleteInteractive(startIter, endIter *TextIter, defaultEditable bool) bool {
	return TextBufferDeleteInteractive(t, startIter, endIter, defaultEditable)
}
func (t *TextBuffer) DeleteMark(mark *TextMark)    { TextBufferDeleteMark(t, mark) }
func (t *TextBuffer) DeleteMarkByName(name string) { TextBufferDeleteMarkByName(t, name) }
func (t *TextBuffer) DeleteSelection(interactive, defaultEditable bool) bool {
	return TextBufferDeleteSelection(t, interactive, defaultEditable)
}
func (t *TextBuffer) Deserialize(contentBuffer *TextBuffer, format D.Atom, iter *TextIter, data *uint8, length T.Gsize, err **T.GError) bool {
	return TextBufferDeserialize(t, contentBuffer, format, iter, data, length, err)
}
func (t *TextBuffer) DeserializeGetCanCreateTags(format D.Atom) bool {
	return TextBufferDeserializeGetCanCreateTags(t, format)
}
func (t *TextBuffer) DeserializeSetCanCreateTags(format D.Atom, canCreateTags bool) {
	TextBufferDeserializeSetCanCreateTags(t, format, canCreateTags)
}
func (t *TextBuffer) EndUserAction()                 { TextBufferEndUserAction(t) }
func (t *TextBuffer) GetBounds(start, end *TextIter) { TextBufferGetBounds(t, start, end) }
func (t *TextBuffer) GetCharCount() int              { return TextBufferGetCharCount(t) }
func (t *TextBuffer) GetCopyTargetList() *TargetList { return TextBufferGetCopyTargetList(t) }
func (t *TextBuffer) GetDeserializeFormats(nFormats *int) *D.Atom {
	return TextBufferGetDeserializeFormats(t, nFormats)
}
func (t *TextBuffer) GetEndIter(iter *TextIter) { TextBufferGetEndIter(t, iter) }
func (t *TextBuffer) GetHasSelection() bool     { return TextBufferGetHasSelection(t) }
func (t *TextBuffer) GetInsert() *TextMark      { return TextBufferGetInsert(t) }
func (t *TextBuffer) GetIterAtChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	TextBufferGetIterAtChildAnchor(t, iter, anchor)
}
func (t *TextBuffer) GetIterAtLine(iter *TextIter, lineNumber int) {
	TextBufferGetIterAtLine(t, iter, lineNumber)
}
func (t *TextBuffer) GetIterAtLineIndex(iter *TextIter, lineNumber, byteIndex int) {
	TextBufferGetIterAtLineIndex(t, iter, lineNumber, byteIndex)
}
func (t *TextBuffer) GetIterAtLineOffset(iter *TextIter, lineNumber, charOffset int) {
	TextBufferGetIterAtLineOffset(t, iter, lineNumber, charOffset)
}
func (t *TextBuffer) GetIterAtMark(iter *TextIter, mark *TextMark) {
	TextBufferGetIterAtMark(t, iter, mark)
}
func (t *TextBuffer) GetIterAtOffset(iter *TextIter, charOffset int) {
	TextBufferGetIterAtOffset(t, iter, charOffset)
}
func (t *TextBuffer) GetLineCount() int               { return TextBufferGetLineCount(t) }
func (t *TextBuffer) GetMark(name string) *TextMark   { return TextBufferGetMark(t, name) }
func (t *TextBuffer) GetModified() bool               { return TextBufferGetModified(t) }
func (t *TextBuffer) GetPasteTargetList() *TargetList { return TextBufferGetPasteTargetList(t) }
func (t *TextBuffer) GetSelectionBound() *TextMark    { return TextBufferGetSelectionBound(t) }
func (t *TextBuffer) GetSelectionBounds(start, end *TextIter) bool {
	return TextBufferGetSelectionBounds(t, start, end)
}
func (t *TextBuffer) GetSerializeFormats(nFormats *int) *D.Atom {
	return TextBufferGetSerializeFormats(t, nFormats)
}
func (t *TextBuffer) GetSlice(start, end *TextIter, includeHiddenChars bool) string {
	return TextBufferGetSlice(t, start, end, includeHiddenChars)
}
func (t *TextBuffer) GetStartIter(iter *TextIter) { TextBufferGetStartIter(t, iter) }
func (t *TextBuffer) GetTagTable() *TextTagTable  { return TextBufferGetTagTable(t) }
func (t *TextBuffer) GetText(start, end *TextIter, includeHiddenChars bool) string {
	return TextBufferGetText(t, start, end, includeHiddenChars)
}
func (t *TextBuffer) Insert(iter *TextIter, text string, leng int) {
	TextBufferInsert(t, iter, text, leng)
}
func (t *TextBuffer) InsertAtCursor(text string, leng int) { TextBufferInsertAtCursor(t, text, leng) }
func (t *TextBuffer) InsertChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	TextBufferInsertChildAnchor(t, iter, anchor)
}
func (t *TextBuffer) InsertInteractive(iter *TextIter, text string, leng int, defaultEditable bool) bool {
	return TextBufferInsertInteractive(t, iter, text, leng, defaultEditable)
}
func (t *TextBuffer) InsertInteractiveAtCursor(text string, leng int, defaultEditable bool) bool {
	return TextBufferInsertInteractiveAtCursor(t, text, leng, defaultEditable)
}
func (t *TextBuffer) InsertPixbuf(iter *TextIter, pixbuf *D.Pixbuf) {
	TextBufferInsertPixbuf(t, iter, pixbuf)
}
func (t *TextBuffer) InsertRange(iter, start, end *TextIter) {
	TextBufferInsertRange(t, iter, start, end)
}
func (t *TextBuffer) InsertRangeInteractive(iter, start, end *TextIter, defaultEditable bool) bool {
	return TextBufferInsertRangeInteractive(t, iter, start, end, defaultEditable)
}
func (t *TextBuffer) InsertWithTags(iter *TextIter, text string, leng int, firstTag *TextTag, v ...VArg) {
	TextBufferInsertWithTags(t, iter, text, leng, firstTag, v)
}
func (t *TextBuffer) InsertWithTagsByName(iter *TextIter, text string, leng int, firstTagName string, v ...VArg) {
	TextBufferInsertWithTagsByName(t, iter, text, leng, firstTagName, v)
}
func (t *TextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	TextBufferMoveMark(t, mark, where)
}
func (t *TextBuffer) MoveMarkByName(name string, where *TextIter) {
	TextBufferMoveMarkByName(t, name, where)
}
func (t *TextBuffer) PasteClipboard(clipboard *Clipboard, overrideLocation *TextIter, defaultEditable bool) {
	TextBufferPasteClipboard(t, clipboard, overrideLocation, defaultEditable)
}
func (t *TextBuffer) PlaceCursor(where *TextIter) { TextBufferPlaceCursor(t, where) }
func (t *TextBuffer) RegisterDeserializeFormat(mimeType string, function TextBufferDeserializeFunc, userData T.Gpointer, userDataDestroy T.GDestroyNotify) D.Atom {
	return TextBufferRegisterDeserializeFormat(t, mimeType, function, userData, userDataDestroy)
}
func (t *TextBuffer) RegisterDeserializeTagset(tagsetName string) D.Atom {
	return TextBufferRegisterDeserializeTagset(t, tagsetName)
}
func (t *TextBuffer) RegisterSerializeFormat(mimeType string, function TextBufferSerializeFunc, userData T.Gpointer, userDataDestroy T.GDestroyNotify) D.Atom {
	return TextBufferRegisterSerializeFormat(t, mimeType, function, userData, userDataDestroy)
}
func (t *TextBuffer) RegisterSerializeTagset(tagsetName string) D.Atom {
	return TextBufferRegisterSerializeTagset(t, tagsetName)
}
func (t *TextBuffer) RemoveAllTags(start, end *TextIter) { TextBufferRemoveAllTags(t, start, end) }
func (t *TextBuffer) RemoveSelectionClipboard(clipboard *Clipboard) {
	TextBufferRemoveSelectionClipboard(t, clipboard)
}
func (t *TextBuffer) RemoveTag(tag *TextTag, start, end *TextIter) {
	TextBufferRemoveTag(t, tag, start, end)
}
func (t *TextBuffer) RemoveTagByName(name string, start, end *TextIter) {
	TextBufferRemoveTagByName(t, name, start, end)
}
func (t *TextBuffer) SelectRange(ins, bound *TextIter) { TextBufferSelectRange(t, ins, bound) }
func (t *TextBuffer) Serialize(contentBuffer *TextBuffer, format D.Atom, start, end *TextIter, length *T.Gsize) *uint8 {
	return TextBufferSerialize(t, contentBuffer, format, start, end, length)
}
func (t *TextBuffer) SetModified(setting bool)      { TextBufferSetModified(t, setting) }
func (t *TextBuffer) SetText(text string, leng int) { TextBufferSetText(t, text, leng) }
func (t *TextBuffer) UnregisterDeserializeFormat(format D.Atom) {
	TextBufferUnregisterDeserializeFormat(t, format)
}
func (t *TextBuffer) UnregisterSerializeFormat(format D.Atom) {
	TextBufferUnregisterSerializeFormat(t, format)
}

type TextCharPredicate func(
	ch T.Gunichar, userData T.Gpointer) bool

type TextChildAnchor struct {
	Parent  O.Object
	Segment T.Gpointer
}

var (
	TextChildAnchorGetType func() O.Type
	TextChildAnchorNew     func() *TextChildAnchor

	TextChildAnchorGetDeleted func(t *TextChildAnchor) bool
	TextChildAnchorGetWidgets func(t *TextChildAnchor) *T.GList
)

func (t *TextChildAnchor) GetDeleted() bool     { return TextChildAnchorGetDeleted(t) }
func (t *TextChildAnchor) GetWidgets() *T.GList { return TextChildAnchorGetWidgets(t) }

var TextDirectionGetType func() O.Type

type TextDirection Enum

const (
	TEXT_DIR_NONE TextDirection = iota
	TEXT_DIR_LTR
	TEXT_DIR_RTL
)

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
	TextIterGetType func() O.Type

	TextIterBackwardChar                   func(t *TextIter) bool
	TextIterBackwardChars                  func(t *TextIter, count int) bool
	TextIterBackwardCursorPosition         func(t *TextIter) bool
	TextIterBackwardCursorPositions        func(t *TextIter, count int) bool
	TextIterBackwardFindChar               func(t *TextIter, pred TextCharPredicate, userData T.Gpointer, limit *TextIter) bool
	TextIterBackwardLine                   func(t *TextIter) bool
	TextIterBackwardLines                  func(t *TextIter, count int) bool
	TextIterBackwardSearch                 func(t *TextIter, str string, flags TextSearchFlags, matchStart, matchEnd, limit *TextIter) bool
	TextIterBackwardSentenceStart          func(t *TextIter) bool
	TextIterBackwardSentenceStarts         func(t *TextIter, count int) bool
	TextIterBackwardToTagToggle            func(t *TextIter, tag *TextTag) bool
	TextIterBackwardVisibleCursorPosition  func(t *TextIter) bool
	TextIterBackwardVisibleCursorPositions func(t *TextIter, count int) bool
	TextIterBackwardVisibleLine            func(t *TextIter) bool
	TextIterBackwardVisibleLines           func(t *TextIter, count int) bool
	TextIterBackwardVisibleWordStart       func(t *TextIter) bool
	TextIterBackwardVisibleWordStarts      func(t *TextIter, count int) bool
	TextIterBackwardWordStart              func(t *TextIter) bool
	TextIterBackwardWordStarts             func(t *TextIter, count int) bool
	TextIterBeginsTag                      func(t *TextIter, tag *TextTag) bool
	TextIterCanInsert                      func(t *TextIter, defaultEditability bool) bool
	TextIterCompare                        func(t *TextIter, rhs *TextIter) int
	TextIterCopy                           func(t *TextIter) *TextIter
	TextIterEditable                       func(t *TextIter, defaultSetting bool) bool
	TextIterEndsLine                       func(t *TextIter) bool
	TextIterEndsSentence                   func(t *TextIter) bool
	TextIterEndsTag                        func(t *TextIter, tag *TextTag) bool
	TextIterEndsWord                       func(t *TextIter) bool
	TextIterEqual                          func(t *TextIter, rhs *TextIter) bool
	TextIterForwardChar                    func(t *TextIter) bool
	TextIterForwardChars                   func(t *TextIter, count int) bool
	TextIterForwardCursorPosition          func(t *TextIter) bool
	TextIterForwardCursorPositions         func(t *TextIter, count int) bool
	TextIterForwardFindChar                func(t *TextIter, pred TextCharPredicate, userData T.Gpointer, limit *TextIter) bool
	TextIterForwardLine                    func(t *TextIter) bool
	TextIterForwardLines                   func(t *TextIter, count int) bool
	TextIterForwardSearch                  func(t *TextIter, str string, flags TextSearchFlags, matchStart, matchEnd, limit *TextIter) bool
	TextIterForwardSentenceEnd             func(t *TextIter) bool
	TextIterForwardSentenceEnds            func(t *TextIter, count int) bool
	TextIterForwardToEnd                   func(t *TextIter)
	TextIterForwardToLineEnd               func(t *TextIter) bool
	TextIterForwardToTagToggle             func(t *TextIter, tag *TextTag) bool
	TextIterForwardVisibleCursorPosition   func(t *TextIter) bool
	TextIterForwardVisibleCursorPositions  func(t *TextIter, count int) bool
	TextIterForwardVisibleLine             func(t *TextIter) bool
	TextIterForwardVisibleLines            func(t *TextIter, count int) bool
	TextIterForwardVisibleWordEnd          func(t *TextIter) bool
	TextIterForwardVisibleWordEnds         func(t *TextIter, count int) bool
	TextIterForwardWordEnd                 func(t *TextIter) bool
	TextIterForwardWordEnds                func(t *TextIter, count int) bool
	TextIterFree                           func(t *TextIter)
	TextIterGetAttributes                  func(t *TextIter, values *TextAttributes) bool
	TextIterGetBuffer                      func(t *TextIter) *TextBuffer
	TextIterGetBytesInLine                 func(t *TextIter) int
	TextIterGetChar                        func(t *TextIter) T.Gunichar
	TextIterGetCharsInLine                 func(t *TextIter) int
	TextIterGetChildAnchor                 func(t *TextIter) *TextChildAnchor
	TextIterGetLanguage                    func(t *TextIter) *P.Language
	TextIterGetLine                        func(t *TextIter) int
	TextIterGetLineIndex                   func(t *TextIter) int
	TextIterGetLineOffset                  func(t *TextIter) int
	TextIterGetMarks                       func(t *TextIter) *L.SList
	TextIterGetOffset                      func(t *TextIter) int
	TextIterGetPixbuf                      func(t *TextIter) *D.Pixbuf
	TextIterGetSlice                       func(t *TextIter, end *TextIter) string
	TextIterGetTags                        func(t *TextIter) *L.SList
	TextIterGetText                        func(t *TextIter, end *TextIter) string
	TextIterGetToggledTags                 func(t *TextIter, toggledOn bool) *L.SList
	TextIterGetVisibleLineIndex            func(t *TextIter) int
	TextIterGetVisibleLineOffset           func(t *TextIter) int
	TextIterGetVisibleSlice                func(t *TextIter, end *TextIter) string
	TextIterGetVisibleText                 func(t *TextIter, end *TextIter) string
	TextIterHasTag                         func(t *TextIter, tag *TextTag) bool
	TextIterInRange                        func(t *TextIter, start, end *TextIter) bool
	TextIterInsideSentence                 func(t *TextIter) bool
	TextIterInsideWord                     func(t *TextIter) bool
	TextIterIsCursorPosition               func(t *TextIter) bool
	TextIterIsEnd                          func(t *TextIter) bool
	TextIterIsStart                        func(t *TextIter) bool
	TextIterOrder                          func(t *TextIter, second *TextIter)
	TextIterSetLine                        func(t *TextIter, lineNumber int)
	TextIterSetLineIndex                   func(t *TextIter, byteOnLine int)
	TextIterSetLineOffset                  func(t *TextIter, charOnLine int)
	TextIterSetOffset                      func(t *TextIter, charOffset int)
	TextIterSetVisibleLineIndex            func(t *TextIter, byteOnLine int)
	TextIterSetVisibleLineOffset           func(t *TextIter, charOnLine int)
	TextIterStartsLine                     func(t *TextIter) bool
	TextIterStartsSentence                 func(t *TextIter) bool
	TextIterStartsWord                     func(t *TextIter) bool
	TextIterTogglesTag                     func(t *TextIter, tag *TextTag) bool
)

func (t *TextIter) BackwardChar() bool           { return TextIterBackwardChar(t) }
func (t *TextIter) BackwardChars(count int) bool { return TextIterBackwardChars(t, count) }
func (t *TextIter) BackwardCursorPosition() bool { return TextIterBackwardCursorPosition(t) }
func (t *TextIter) BackwardCursorPositions(count int) bool {
	return TextIterBackwardCursorPositions(t, count)
}
func (t *TextIter) BackwardFindChar(pred TextCharPredicate, userData T.Gpointer, limit *TextIter) bool {
	return TextIterBackwardFindChar(t, pred, userData, limit)
}
func (t *TextIter) BackwardLine() bool           { return TextIterBackwardLine(t) }
func (t *TextIter) BackwardLines(count int) bool { return TextIterBackwardLines(t, count) }
func (t *TextIter) BackwardSearch(str string, flags TextSearchFlags, matchStart, matchEnd, limit *TextIter) bool {
	return TextIterBackwardSearch(t, str, flags, matchStart, matchEnd, limit)
}
func (t *TextIter) BackwardSentenceStart() bool { return TextIterBackwardSentenceStart(t) }
func (t *TextIter) BackwardSentenceStarts(count int) bool {
	return TextIterBackwardSentenceStarts(t, count)
}
func (t *TextIter) BackwardToTagToggle(tag *TextTag) bool {
	return TextIterBackwardToTagToggle(t, tag)
}
func (t *TextIter) BackwardVisibleCursorPosition() bool {
	return TextIterBackwardVisibleCursorPosition(t)
}
func (t *TextIter) BackwardVisibleCursorPositions(count int) bool {
	return TextIterBackwardVisibleCursorPositions(t, count)
}
func (t *TextIter) BackwardVisibleLine() bool { return TextIterBackwardVisibleLine(t) }
func (t *TextIter) BackwardVisibleLines(count int) bool {
	return TextIterBackwardVisibleLines(t, count)
}
func (t *TextIter) BackwardVisibleWordStart() bool { return TextIterBackwardVisibleWordStart(t) }
func (t *TextIter) BackwardVisibleWordStarts(count int) bool {
	return TextIterBackwardVisibleWordStarts(t, count)
}
func (t *TextIter) BackwardWordStart() bool { return TextIterBackwardWordStart(t) }
func (t *TextIter) BackwardWordStarts(count int) bool {
	return TextIterBackwardWordStarts(t, count)
}
func (t *TextIter) BeginsTag(tag *TextTag) bool { return TextIterBeginsTag(t, tag) }
func (t *TextIter) CanInsert(defaultEditability bool) bool {
	return TextIterCanInsert(t, defaultEditability)
}
func (t *TextIter) Compare(rhs *TextIter) int { return TextIterCompare(t, rhs) }
func (t *TextIter) Copy() *TextIter           { return TextIterCopy(t) }
func (t *TextIter) Editable(defaultSetting bool) bool {
	return TextIterEditable(t, defaultSetting)
}
func (t *TextIter) EndsLine() bool              { return TextIterEndsLine(t) }
func (t *TextIter) EndsSentence() bool          { return TextIterEndsSentence(t) }
func (t *TextIter) EndsTag(tag *TextTag) bool   { return TextIterEndsTag(t, tag) }
func (t *TextIter) EndsWord() bool              { return TextIterEndsWord(t) }
func (t *TextIter) Equal(rhs *TextIter) bool    { return TextIterEqual(t, rhs) }
func (t *TextIter) ForwardChar() bool           { return TextIterForwardChar(t) }
func (t *TextIter) ForwardChars(count int) bool { return TextIterForwardChars(t, count) }
func (t *TextIter) ForwardCursorPosition() bool { return TextIterForwardCursorPosition(t) }
func (t *TextIter) ForwardCursorPositions(count int) bool {
	return TextIterForwardCursorPositions(t, count)
}
func (t *TextIter) ForwardFindChar(pred TextCharPredicate, userData T.Gpointer, limit *TextIter) bool {
	return TextIterForwardFindChar(t, pred, userData, limit)
}
func (t *TextIter) ForwardLine() bool           { return TextIterForwardLine(t) }
func (t *TextIter) ForwardLines(count int) bool { return TextIterForwardLines(t, count) }
func (t *TextIter) ForwardSearch(str string, flags TextSearchFlags, matchStart, matchEnd, limit *TextIter) bool {
	return TextIterForwardSearch(t, str, flags, matchStart, matchEnd, limit)
}
func (t *TextIter) ForwardSentenceEnd() bool { return TextIterForwardSentenceEnd(t) }
func (t *TextIter) ForwardSentenceEnds(count int) bool {
	return TextIterForwardSentenceEnds(t, count)
}
func (t *TextIter) ForwardToEnd()          { TextIterForwardToEnd(t) }
func (t *TextIter) ForwardToLineEnd() bool { return TextIterForwardToLineEnd(t) }
func (t *TextIter) ForwardToTagToggle(tag *TextTag) bool {
	return TextIterForwardToTagToggle(t, tag)
}
func (t *TextIter) ForwardVisibleCursorPosition() bool {
	return TextIterForwardVisibleCursorPosition(t)
}
func (t *TextIter) ForwardVisibleCursorPositions(count int) bool {
	return TextIterForwardVisibleCursorPositions(t, count)
}
func (t *TextIter) ForwardVisibleLine() bool { return TextIterForwardVisibleLine(t) }
func (t *TextIter) ForwardVisibleLines(count int) bool {
	return TextIterForwardVisibleLines(t, count)
}
func (t *TextIter) ForwardVisibleWordEnd() bool { return TextIterForwardVisibleWordEnd(t) }
func (t *TextIter) ForwardVisibleWordEnds(count int) bool {
	return TextIterForwardVisibleWordEnds(t, count)
}
func (t *TextIter) ForwardWordEnd() bool           { return TextIterForwardWordEnd(t) }
func (t *TextIter) ForwardWordEnds(count int) bool { return TextIterForwardWordEnds(t, count) }
func (t *TextIter) Free()                          { TextIterFree(t) }
func (t *TextIter) GetAttributes(values *TextAttributes) bool {
	return TextIterGetAttributes(t, values)
}
func (t *TextIter) GetBuffer() *TextBuffer           { return TextIterGetBuffer(t) }
func (t *TextIter) GetBytesInLine() int              { return TextIterGetBytesInLine(t) }
func (t *TextIter) GetChar() T.Gunichar              { return TextIterGetChar(t) }
func (t *TextIter) GetCharsInLine() int              { return TextIterGetCharsInLine(t) }
func (t *TextIter) GetChildAnchor() *TextChildAnchor { return TextIterGetChildAnchor(t) }
func (t *TextIter) GetLanguage() *P.Language         { return TextIterGetLanguage(t) }
func (t *TextIter) GetLine() int                     { return TextIterGetLine(t) }
func (t *TextIter) GetLineIndex() int                { return TextIterGetLineIndex(t) }
func (t *TextIter) GetLineOffset() int               { return TextIterGetLineOffset(t) }
func (t *TextIter) GetMarks() *L.SList               { return TextIterGetMarks(t) }
func (t *TextIter) GetOffset() int                   { return TextIterGetOffset(t) }
func (t *TextIter) GetPixbuf() *D.Pixbuf             { return TextIterGetPixbuf(t) }
func (t *TextIter) GetSlice(end *TextIter) string    { return TextIterGetSlice(t, end) }
func (t *TextIter) GetTags() *L.SList                { return TextIterGetTags(t) }
func (t *TextIter) GetText(end *TextIter) string     { return TextIterGetText(t, end) }
func (t *TextIter) GetToggledTags(toggledOn bool) *L.SList {
	return TextIterGetToggledTags(t, toggledOn)
}
func (t *TextIter) GetVisibleLineIndex() int             { return TextIterGetVisibleLineIndex(t) }
func (t *TextIter) GetVisibleLineOffset() int            { return TextIterGetVisibleLineOffset(t) }
func (t *TextIter) GetVisibleSlice(end *TextIter) string { return TextIterGetVisibleSlice(t, end) }
func (t *TextIter) GetVisibleText(end *TextIter) string  { return TextIterGetVisibleText(t, end) }
func (t *TextIter) HasTag(tag *TextTag) bool             { return TextIterHasTag(t, tag) }
func (t *TextIter) InRange(start, end *TextIter) bool    { return TextIterInRange(t, start, end) }
func (t *TextIter) InsideSentence() bool                 { return TextIterInsideSentence(t) }
func (t *TextIter) InsideWord() bool                     { return TextIterInsideWord(t) }
func (t *TextIter) IsCursorPosition() bool               { return TextIterIsCursorPosition(t) }
func (t *TextIter) IsEnd() bool                          { return TextIterIsEnd(t) }
func (t *TextIter) IsStart() bool                        { return TextIterIsStart(t) }
func (t *TextIter) Order(second *TextIter)               { TextIterOrder(t, second) }
func (t *TextIter) SetLine(lineNumber int)               { TextIterSetLine(t, lineNumber) }
func (t *TextIter) SetLineIndex(byteOnLine int)          { TextIterSetLineIndex(t, byteOnLine) }
func (t *TextIter) SetLineOffset(charOnLine int)         { TextIterSetLineOffset(t, charOnLine) }
func (t *TextIter) SetOffset(charOffset int)             { TextIterSetOffset(t, charOffset) }
func (t *TextIter) SetVisibleLineIndex(byteOnLine int)   { TextIterSetVisibleLineIndex(t, byteOnLine) }
func (t *TextIter) SetVisibleLineOffset(charOnLine int)  { TextIterSetVisibleLineOffset(t, charOnLine) }
func (t *TextIter) StartsLine() bool                     { return TextIterStartsLine(t) }
func (t *TextIter) StartsSentence() bool                 { return TextIterStartsSentence(t) }
func (t *TextIter) StartsWord() bool                     { return TextIterStartsWord(t) }
func (t *TextIter) TogglesTag(tag *TextTag) bool         { return TextIterTogglesTag(t, tag) }

type (
	TextLayout       struct{}
	TextLogAttrCache struct{}
)

type TextMark struct {
	Parent  O.Object
	Segment T.Gpointer
}

var (
	TextMarkGetType func() O.Type
	TextMarkNew     func(name string, leftGravity bool) *TextMark

	TextMarkSetVisible     func(t *TextMark, setting bool)
	TextMarkGetVisible     func(t *TextMark) bool
	TextMarkGetName        func(t *TextMark) string
	TextMarkGetDeleted     func(t *TextMark) bool
	TextMarkGetBuffer      func(t *TextMark) *TextBuffer
	TextMarkGetLeftGravity func(t *TextMark) bool
)

type TextPendingScroll struct{}

var TextSearchFlagsGetType func() O.Type

type TextSearchFlags Enum

const (
	TEXT_SEARCH_VISIBLE_ONLY TextSearchFlags = 1 << iota
	TEXT_SEARCH_TEXT_ONLY
)

type TextTag struct {
	Parent   O.Object
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
	TextTagGetType func() O.Type
	TextTagNew     func(name string) *TextTag

	TextTagEvent       func(t *TextTag, eventObject *O.Object, event *D.Event, iter *TextIter) bool
	TextTagGetPriority func(t *TextTag) int
	TextTagSetPriority func(t *TextTag, priority int)
)

func (t *TextTag) Event(eventObject *O.Object, event *D.Event, iter *TextIter) bool {
	return TextTagEvent(t, eventObject, event, iter)
}
func (t *TextTag) GetPriority() int         { return TextTagGetPriority(t) }
func (t *TextTag) SetPriority(priority int) { TextTagSetPriority(t, priority) }

type (
	TextTagTable struct {
		Parent    O.Object
		Hash      *L.HashTable
		Anonymous *L.SList
		AnonCount int
		Buffers   *L.SList
	}

	TextTagTableForeachFunc func(tag *TextTag, data T.Gpointer)
)

var (
	TextTagTableGetType func() O.Type
	TextTagTableNew     func() *TextTagTable

	TextTagTableAdd     func(t *TextTagTable, tag *TextTag)
	TextTagTableForeach func(t *TextTagTable, f TextTagTableForeachFunc, data T.Gpointer)
	TextTagTableGetSize func(t *TextTagTable) int
	TextTagTableLookup  func(t *TextTagTable, name string) *TextTag
	TextTagTableRemove  func(t *TextTagTable, tag *TextTag)
)

func (t *TextTagTable) Add(tag *TextTag) { TextTagTableAdd(t, tag) }
func (t *TextTagTable) Foreach(f TextTagTableForeachFunc, data T.Gpointer) {
	TextTagTableForeach(t, f, data)
}
func (t *TextTagTable) GetSize() int                { return TextTagTableGetSize(t) }
func (t *TextTagTable) Lookup(name string) *TextTag { return TextTagTableLookup(t, name) }
func (t *TextTagTable) Remove(tag *TextTag)         { TextTagTableRemove(t, tag) }

type TextView struct {
	Parent               Container
	Layout               *TextLayout
	Buffer               *TextBuffer
	SelectionDragHandler uint
	ScrollTimeout        uint
	PixelsAboveLines     int
	PixelsBelowLines     int
	PixelsInsideWrap     int
	WrapMode             WrapMode
	Justify              Justification
	LeftMargin           int
	RightMargin          int
	Indent               int
	Tabs                 *P.TabArray
	Bits                 uint
	// Editable : 1
	// OverwriteMode : 1
	// CursorVisible : 1
	// NeedImReset : 1
	// AcceptsTab : 1
	// WidthChanged : 1
	// OnscreenValidated : 1
	// MouseCursorObscured : 1
	TextWindow               *TextWindow
	LeftWindow               *TextWindow
	RightWindow              *TextWindow
	TopWindow                *TextWindow
	BottomWindow             *TextWindow
	Hadjustment              *Adjustment
	Vadjustment              *Adjustment
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
	Children                 *L.SList
	PendingScroll            *TextPendingScroll
	PendingPlaceCursorButton int
}

var (
	TextViewGetType       func() O.Type
	TextViewNew           func() *Widget
	TextViewNewWithBuffer func(buffer *TextBuffer) *Widget

	TextViewSetBuffer                func(t *TextView, buffer *TextBuffer)
	TextViewGetBuffer                func(t *TextView) *TextBuffer
	TextViewScrollToIter             func(t *TextView, iter *TextIter, withinMargin float64, useAlign bool, xalign, yalign float64) bool
	TextViewScrollToMark             func(t *TextView, mark *TextMark, withinMargin float64, useAlign bool, xalign, yalign float64)
	TextViewScrollMarkOnscreen       func(t *TextView, mark *TextMark)
	TextViewMoveMarkOnscreen         func(t *TextView, mark *TextMark) bool
	TextViewPlaceCursorOnscreen      func(t *TextView) bool
	TextViewGetVisibleRect           func(t *TextView, visibleRect *D.Rectangle)
	TextViewSetCursorVisible         func(t *TextView, setting bool)
	TextViewGetCursorVisible         func(t *TextView) bool
	TextViewGetIterLocation          func(t *TextView, iter *TextIter, location *D.Rectangle)
	TextViewGetIterAtLocation        func(t *TextView, iter *TextIter, x, y int)
	TextViewGetIterAtPosition        func(t *TextView, iter *TextIter, trailing *int, x, y int)
	TextViewGetLineYrange            func(t *TextView, iter *TextIter, y, height *int)
	TextViewGetLineAtY               func(t *TextView, targetIter *TextIter, y int, lineTop *int)
	TextViewBufferToWindowCoords     func(t *TextView, win TextWindowType, bufferX, bufferY int, windowX, windowY *int)
	TextViewWindowToBufferCoords     func(t *TextView, win TextWindowType, windowX, windowY int, bufferX, bufferY *int)
	TextViewGetHadjustment           func(t *TextView) *Adjustment
	TextViewGetVadjustment           func(t *TextView) *Adjustment
	TextViewGetWindow                func(t *TextView, win TextWindowType) *D.Window
	TextViewGetWindowType            func(t *TextView, window *D.Window) TextWindowType
	TextViewSetBorderWindowSize      func(t *TextView, tw TextWindowType, size int)
	TextViewGetBorderWindowSize      func(t *TextView, tw TextWindowType) int
	TextViewForwardDisplayLine       func(t *TextView, iter *TextIter) bool
	TextViewBackwardDisplayLine      func(t *TextView, iter *TextIter) bool
	TextViewForwardDisplayLineEnd    func(t *TextView, iter *TextIter) bool
	TextViewBackwardDisplayLineStart func(t *TextView, iter *TextIter) bool
	TextViewStartsDisplayLine        func(t *TextView, iter *TextIter) bool
	TextViewMoveVisually             func(t *TextView, iter *TextIter, count int) bool
	TextViewImContextFilterKeypress  func(t *TextView, event *D.EventKey) bool
	TextViewResetImContext           func(t *TextView)
	TextViewAddChildAtAnchor         func(t *TextView, child *Widget, anchor *TextChildAnchor)
	TextViewAddChildInWindow         func(t *TextView, child *Widget, whichWindow TextWindowType, xpos, ypos int)
	TextViewMoveChild                func(t *TextView, child *Widget, xpos, ypos int)
	TextViewSetWrapMode              func(t *TextView, wrapMode WrapMode)
	TextViewGetWrapMode              func(t *TextView) WrapMode
	TextViewSetEditable              func(t *TextView, setting bool)
	TextViewGetEditable              func(t *TextView) bool
	TextViewSetOverwrite             func(t *TextView, overwrite bool)
	TextViewGetOverwrite             func(t *TextView) bool
	TextViewSetAcceptsTab            func(t *TextView, acceptsTab bool)
	TextViewGetAcceptsTab            func(t *TextView) bool
	TextViewSetPixelsAboveLines      func(t *TextView, pixelsAboveLines int)
	TextViewGetPixelsAboveLines      func(t *TextView) int
	TextViewSetPixelsBelowLines      func(t *TextView, pixelsBelowLines int)
	TextViewGetPixelsBelowLines      func(t *TextView) int
	TextViewSetPixelsInsideWrap      func(t *TextView, pixelsInsideWrap int)
	TextViewGetPixelsInsideWrap      func(t *TextView) int
	TextViewSetJustification         func(t *TextView, justification Justification)
	TextViewGetJustification         func(t *TextView) Justification
	TextViewSetLeftMargin            func(t *TextView, leftMargin int)
	TextViewGetLeftMargin            func(t *TextView) int
	TextViewSetRightMargin           func(t *TextView, rightMargin int)
	TextViewGetRightMargin           func(t *TextView) int
	TextViewSetIndent                func(t *TextView, indent int)
	TextViewGetIndent                func(t *TextView) int
	TextViewSetTabs                  func(t *TextView, tabs *P.TabArray)
	TextViewGetTabs                  func(t *TextView) *P.TabArray
	TextViewGetDefaultAttributes     func(t *TextView) *TextAttributes
)

type TextWindow struct{}

var TextWindowTypeGetType func() O.Type

type TextWindowType Enum

const (
	TEXT_WINDOW_PRIVATE TextWindowType = iota
	TEXT_WINDOW_WIDGET
	TEXT_WINDOW_TEXT
	TEXT_WINDOW_LEFT
	TEXT_WINDOW_RIGHT
	TEXT_WINDOW_TOP
	TEXT_WINDOW_BOTTOM
)

var (
	TimeoutAdd     func(interval T.GUint32, function Function, data T.Gpointer) uint
	TimeoutAddFull func(interval T.GUint32, function Function, marshal CallbackMarshal, data T.Gpointer, destroy T.GDestroyNotify) uint
	TimeoutRemove  func(timeoutHandlerId uint)
)

type TipsQuery struct {
	Label Label
	Bits  uint
	// EmitAlways : 1
	// InQuery : 1
	LabelInactive *T.Gchar
	LabelNoTip    *T.Gchar
	Caller        *Widget
	LastCrossed   *Widget
	QueryCursor   *D.Cursor
}

var (
	TipsQueryGetType func() O.Type
	TipsQueryNew     func() *Widget

	TipsQuerySetCaller  func(t *TipsQuery, caller *Widget)
	TipsQuerySetLabels  func(t *TipsQuery, labelInactive, labelNoTip string)
	TipsQueryStartQuery func(t *TipsQuery)
	TipsQueryStopQuery  func(t *TipsQuery)
)

func (t *TipsQuery) SetCaller(caller *Widget) { TipsQuerySetCaller(t, caller) }
func (t *TipsQuery) SetLabels(labelInactive, labelNoTip string) {
	TipsQuerySetLabels(t, labelInactive, labelNoTip)
}
func (t *TipsQuery) StartQuery() { TipsQueryStartQuery(t) }
func (t *TipsQuery) StopQuery()  { TipsQueryStopQuery(t) }

type ToggleAction struct {
	Parent Action
	_      *struct{}
}

var (
	ToggleActionGetType func() O.Type
	ToggleActionNew     func(name, label, tooltip, stockId string) *ToggleAction

	ToggleActionGetActive      func(t *ToggleAction) bool
	ToggleActionGetDrawAsRadio func(t *ToggleAction) bool
	ToggleActionSetActive      func(t *ToggleAction, isActive bool)
	ToggleActionSetDrawAsRadio func(t *ToggleAction, drawAsRadio bool)
	ToggleActionToggled        func(t *ToggleAction)
)

func (t *ToggleAction) GetActive() bool         { return ToggleActionGetActive(t) }
func (t *ToggleAction) GetDrawAsRadio() bool    { return ToggleActionGetDrawAsRadio(t) }
func (t *ToggleAction) SetActive(isActive bool) { ToggleActionSetActive(t, isActive) }
func (t *ToggleAction) SetDrawAsRadio(drawAsRadio bool) {
	ToggleActionSetDrawAsRadio(t, drawAsRadio)
}
func (t *ToggleAction) Toggled() { ToggleActionToggled(t) }

type ToggleActionEntry struct {
	Name        *T.Gchar
	StockId     *T.Gchar
	Label       *T.Gchar
	Accelerator *T.Gchar
	Tooltip     *T.Gchar
	Callback    T.GCallback
	IsActive    T.Gboolean
}

type ToggleButton struct {
	Button Button
	Bits   uint
	// Active : 1
	// DrawIndicator : 1
	// Inconsistent : 1
}

var (
	ToggleButtonGetType         func() O.Type
	ToggleButtonNew             func() *Widget
	ToggleButtonNewWithLabel    func(label string) *Widget
	ToggleButtonNewWithMnemonic func(label string) *Widget

	ToggleButtonGetActive       func(t *ToggleButton) bool
	ToggleButtonGetInconsistent func(t *ToggleButton) bool
	ToggleButtonGetMode         func(t *ToggleButton) bool
	ToggleButtonSetActive       func(t *ToggleButton, isActive bool)
	ToggleButtonSetInconsistent func(t *ToggleButton, setting bool)
	ToggleButtonSetMode         func(t *ToggleButton, drawIndicator bool)
	ToggleButtonToggled         func(t *ToggleButton)
)

func (t *ToggleButton) GetActive() bool              { return ToggleButtonGetActive(t) }
func (t *ToggleButton) GetInconsistent() bool        { return ToggleButtonGetInconsistent(t) }
func (t *ToggleButton) GetMode() bool                { return ToggleButtonGetMode(t) }
func (t *ToggleButton) SetActive(isActive bool)      { ToggleButtonSetActive(t, isActive) }
func (t *ToggleButton) SetInconsistent(setting bool) { ToggleButtonSetInconsistent(t, setting) }
func (t *ToggleButton) SetMode(drawIndicator bool)   { ToggleButtonSetMode(t, drawIndicator) }
func (t *ToggleButton) Toggled()                     { ToggleButtonToggled(t) }

type ToggleToolButton struct {
	Parent ToolButton
	_      *struct{}
}

var (
	ToggleToolButtonGetType      func() O.Type
	ToggleToolButtonNew          func() *ToolItem
	ToggleToolButtonNewFromStock func(stockId string) *ToolItem

	ToggleToolButtonGetActive func(t *ToggleToolButton) bool
	ToggleToolButtonSetActive func(t *ToggleToolButton, isActive bool)
)

func (t *ToggleToolButton) GetActive() bool         { return ToggleToolButtonGetActive(t) }
func (t *ToggleToolButton) SetActive(isActive bool) { ToggleToolButtonSetActive(t, isActive) }

type (
	Toolbar struct {
		Container   Container
		NumChildren int
		Children    *T.GList
		Orientation Orientation
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

type ToolbarStyle Enum

const (
	TOOLBAR_ICONS ToolbarStyle = iota
	TOOLBAR_TEXT
	TOOLBAR_BOTH
	TOOLBAR_BOTH_HORIZ
)

type ToolbarChildType Enum

const (
	TOOLBAR_CHILD_SPACE ToolbarChildType = iota
	TOOLBAR_CHILD_BUTTON
	TOOLBAR_CHILD_TOGGLEBUTTON
	TOOLBAR_CHILD_RADIOBUTTON
	TOOLBAR_CHILD_WIDGET
)

var (
	ToolbarGetType           func() O.Type
	ToolbarNew               func() *Widget
	ToolbarChildTypeGetType  func() O.Type
	ToolbarSpaceStyleGetType func() O.Type
	ToolbarStyleGetType      func() O.Type

	ToolbarAppendElement        func(t *Toolbar, ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget
	ToolbarAppendItem           func(t *Toolbar, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget
	ToolbarAppendSpace          func(t *Toolbar)
	ToolbarAppendWidget         func(t *Toolbar, widget *Widget, tooltipText, tooltipPrivateText string)
	ToolbarGetDropIndex         func(t *Toolbar, x, y int) int
	ToolbarGetIconSize          func(t *Toolbar) IconSize
	ToolbarGetItemIndex         func(t *Toolbar, item *ToolItem) int
	ToolbarGetNItems            func(t *Toolbar) int
	ToolbarGetNthItem           func(t *Toolbar, n int) *ToolItem
	ToolbarGetOrientation       func(t *Toolbar) Orientation
	ToolbarGetReliefStyle       func(t *Toolbar) ReliefStyle
	ToolbarGetShowArrow         func(t *Toolbar) bool
	ToolbarGetStyle             func(t *Toolbar) ToolbarStyle
	ToolbarGetTooltips          func(t *Toolbar) bool
	ToolbarInsert               func(t *Toolbar, item *ToolItem, pos int)
	ToolbarInsertElement        func(t *Toolbar, ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer, position int) *Widget
	ToolbarInsertItem           func(t *Toolbar, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer, position int) *Widget
	ToolbarInsertSpace          func(t *Toolbar, position int)
	ToolbarInsertStock          func(t *Toolbar, stockId, tooltipText, tooltipPrivateText string, callback T.GCallback, userData T.Gpointer, position int) *Widget
	ToolbarInsertWidget         func(t *Toolbar, widget *Widget, tooltipText, tooltipPrivateText string, position int)
	ToolbarPrependElement       func(t *Toolbar, ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget
	ToolbarPrependItem          func(t *Toolbar, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget
	ToolbarPrependSpace         func(t *Toolbar)
	ToolbarPrependWidget        func(t *Toolbar, widget *Widget, tooltipText, tooltipPrivateText string)
	ToolbarRemoveSpace          func(t *Toolbar, position int)
	ToolbarSetDropHighlightItem func(t *Toolbar, toolItem *ToolItem, index int)
	ToolbarSetIconSize          func(t *Toolbar, iconSize IconSize)
	ToolbarSetOrientation       func(t *Toolbar, orientation Orientation)
	ToolbarSetShowArrow         func(t *Toolbar, showArrow bool)
	ToolbarSetStyle             func(t *Toolbar, style ToolbarStyle)
	ToolbarSetTooltips          func(t *Toolbar, enable bool)
	ToolbarUnsetIconSize        func(t *Toolbar)
	ToolbarUnsetStyle           func(t *Toolbar)
)

func (t *Toolbar) AppendElement(ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget {
	return ToolbarAppendElement(t, ct, widget, text, tooltipText, tooltipPrivateText, icon, callback, userData)
}
func (t *Toolbar) AppendItem(text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget {
	return ToolbarAppendItem(t, text, tooltipText, tooltipPrivateText, icon, callback, userData)
}
func (t *Toolbar) AppendSpace() { ToolbarAppendSpace(t) }
func (t *Toolbar) AppendWidget(widget *Widget, tooltipText, tooltipPrivateText string) {
	ToolbarAppendWidget(t, widget, tooltipText, tooltipPrivateText)
}
func (t *Toolbar) GetDropIndex(x, y int) int       { return ToolbarGetDropIndex(t, x, y) }
func (t *Toolbar) GetIconSize() IconSize           { return ToolbarGetIconSize(t) }
func (t *Toolbar) GetItemIndex(item *ToolItem) int { return ToolbarGetItemIndex(t, item) }
func (t *Toolbar) GetNItems() int                  { return ToolbarGetNItems(t) }
func (t *Toolbar) GetNthItem(n int) *ToolItem      { return ToolbarGetNthItem(t, n) }
func (t *Toolbar) GetOrientation() Orientation     { return ToolbarGetOrientation(t) }
func (t *Toolbar) GetReliefStyle() ReliefStyle     { return ToolbarGetReliefStyle(t) }
func (t *Toolbar) GetShowArrow() bool              { return ToolbarGetShowArrow(t) }
func (t *Toolbar) GetStyle() ToolbarStyle          { return ToolbarGetStyle(t) }
func (t *Toolbar) GetTooltips() bool               { return ToolbarGetTooltips(t) }
func (t *Toolbar) Insert(item *ToolItem, pos int)  { ToolbarInsert(t, item, pos) }
func (t *Toolbar) InsertElement(ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer, position int) *Widget {
	return ToolbarInsertElement(t, ct, widget, text, tooltipText, tooltipPrivateText, icon, callback, userData, position)
}
func (t *Toolbar) InsertItem(text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer, position int) *Widget {
	return ToolbarInsertItem(t, text, tooltipText, tooltipPrivateText, icon, callback, userData, position)
}
func (t *Toolbar) InsertSpace(position int) { ToolbarInsertSpace(t, position) }
func (t *Toolbar) InsertStock(stockId, tooltipText, tooltipPrivateText string, callback T.GCallback, userData T.Gpointer, position int) *Widget {
	return ToolbarInsertStock(t, stockId, tooltipText, tooltipPrivateText, callback, userData, position)
}
func (t *Toolbar) InsertWidget(widget *Widget, tooltipText, tooltipPrivateText string, position int) {
	ToolbarInsertWidget(t, widget, tooltipText, tooltipPrivateText, position)
}
func (t *Toolbar) PrependElement(ct ToolbarChildType, widget *Widget, text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget {
	return ToolbarPrependElement(t, ct, widget, text, tooltipText, tooltipPrivateText, icon, callback, userData)
}
func (t *Toolbar) PrependItem(text, tooltipText, tooltipPrivateText string, icon *Widget, callback T.GCallback, userData T.Gpointer) *Widget {
	return ToolbarPrependItem(t, text, tooltipText, tooltipPrivateText, icon, callback, userData)
}
func (t *Toolbar) PrependSpace() { ToolbarPrependSpace(t) }
func (t *Toolbar) PrependWidget(widget *Widget, tooltipText, tooltipPrivateText string) {
	ToolbarPrependWidget(t, widget, tooltipText, tooltipPrivateText)
}
func (t *Toolbar) RemoveSpace(position int) { ToolbarRemoveSpace(t, position) }
func (t *Toolbar) SetDropHighlightItem(toolItem *ToolItem, index int) {
	ToolbarSetDropHighlightItem(t, toolItem, index)
}
func (t *Toolbar) SetIconSize(iconSize IconSize)          { ToolbarSetIconSize(t, iconSize) }
func (t *Toolbar) SetOrientation(orientation Orientation) { ToolbarSetOrientation(t, orientation) }
func (t *Toolbar) SetShowArrow(showArrow bool)            { ToolbarSetShowArrow(t, showArrow) }
func (t *Toolbar) SetStyle(style ToolbarStyle)            { ToolbarSetStyle(t, style) }
func (t *Toolbar) SetTooltips(enable bool)                { ToolbarSetTooltips(t, enable) }
func (t *Toolbar) UnsetIconSize()                         { ToolbarUnsetIconSize(t) }
func (t *Toolbar) UnsetStyle()                            { ToolbarUnsetStyle(t) }

type ToolButton struct {
	Parent ToolItem
	_      *struct{}
}

var (
	ToolButtonGetType      func() O.Type
	ToolButtonNew          func(iconWidget *Widget, label string) *ToolItem
	ToolButtonNewFromStock func(stockId string) *ToolItem

	ToolButtonGetIconName     func(t *ToolButton) string
	ToolButtonGetIconWidget   func(t *ToolButton) *Widget
	ToolButtonGetLabel        func(t *ToolButton) string
	ToolButtonGetLabelWidget  func(t *ToolButton) *Widget
	ToolButtonGetStockId      func(t *ToolButton) string
	ToolButtonGetUseUnderline func(t *ToolButton) bool
	ToolButtonSetIconName     func(t *ToolButton, iconName string)
	ToolButtonSetIconWidget   func(t *ToolButton, iconWidget *Widget)
	ToolButtonSetLabel        func(t *ToolButton, label string)
	ToolButtonSetLabelWidget  func(t *ToolButton, labelWidget *Widget)
	ToolButtonSetStockId      func(t *ToolButton, stockId string)
	ToolButtonSetUseUnderline func(t *ToolButton, useUnderline bool)
)

func (t *ToolButton) GetIconName() string                { return ToolButtonGetIconName(t) }
func (t *ToolButton) GetIconWidget() *Widget             { return ToolButtonGetIconWidget(t) }
func (t *ToolButton) GetLabel() string                   { return ToolButtonGetLabel(t) }
func (t *ToolButton) GetLabelWidget() *Widget            { return ToolButtonGetLabelWidget(t) }
func (t *ToolButton) GetStockId() string                 { return ToolButtonGetStockId(t) }
func (t *ToolButton) GetUseUnderline() bool              { return ToolButtonGetUseUnderline(t) }
func (t *ToolButton) SetIconName(iconName string)        { ToolButtonSetIconName(t, iconName) }
func (t *ToolButton) SetIconWidget(iconWidget *Widget)   { ToolButtonSetIconWidget(t, iconWidget) }
func (t *ToolButton) SetLabel(label string)              { ToolButtonSetLabel(t, label) }
func (t *ToolButton) SetLabelWidget(labelWidget *Widget) { ToolButtonSetLabelWidget(t, labelWidget) }
func (t *ToolButton) SetStockId(stockId string)          { ToolButtonSetStockId(t, stockId) }
func (t *ToolButton) SetUseUnderline(useUnderline bool) {
	ToolButtonSetUseUnderline(t, useUnderline)
}

type ToolItem struct {
	Parent Bin
	_      *struct{}
}

var (
	ToolItemGetType func() O.Type
	ToolItemNew     func() *ToolItem

	ToolItemGetEllipsizeMode      func(t *ToolItem) P.EllipsizeMode
	ToolItemGetExpand             func(t *ToolItem) bool
	ToolItemGetHomogeneous        func(t *ToolItem) bool
	ToolItemGetIconSize           func(t *ToolItem) IconSize
	ToolItemGetIsImportant        func(t *ToolItem) bool
	ToolItemGetOrientation        func(t *ToolItem) Orientation
	ToolItemGetProxyMenuItem      func(t *ToolItem, menuItemId string) *Widget
	ToolItemGetReliefStyle        func(t *ToolItem) ReliefStyle
	ToolItemGetTextAlignment      func(t *ToolItem) float32
	ToolItemGetTextOrientation    func(t *ToolItem) Orientation
	ToolItemGetTextSizeGroup      func(t *ToolItem) *SizeGroup
	ToolItemGetToolbarStyle       func(t *ToolItem) ToolbarStyle
	ToolItemGetUseDragWindow      func(t *ToolItem) bool
	ToolItemGetVisibleHorizontal  func(t *ToolItem) bool
	ToolItemGetVisibleVertical    func(t *ToolItem) bool
	ToolItemRebuildMenu           func(t *ToolItem)
	ToolItemRetrieveProxyMenuItem func(t *ToolItem) *Widget
	ToolItemSetExpand             func(t *ToolItem, expand bool)
	ToolItemSetHomogeneous        func(t *ToolItem, homogeneous bool)
	ToolItemSetIsImportant        func(t *ToolItem, isImportant bool)
	ToolItemSetProxyMenuItem      func(t *ToolItem, menuItemId string, menuItem *Widget)
	ToolItemSetTooltip            func(t *ToolItem, tooltips *Tooltips, tipText, tipPrivate string)
	ToolItemSetTooltipMarkup      func(t *ToolItem, markup string)
	ToolItemSetTooltipText        func(t *ToolItem, text string)
	ToolItemSetUseDragWindow      func(t *ToolItem, useDragWindow bool)
	ToolItemSetVisibleHorizontal  func(t *ToolItem, visibleHorizontal bool)
	ToolItemSetVisibleVertical    func(t *ToolItem, visibleVertical bool)
	ToolItemToolbarReconfigured   func(t *ToolItem)
)

func (t *ToolItem) GetEllipsizeMode() P.EllipsizeMode { return ToolItemGetEllipsizeMode(t) }
func (t *ToolItem) GetExpand() bool                   { return ToolItemGetExpand(t) }
func (t *ToolItem) GetHomogeneous() bool              { return ToolItemGetHomogeneous(t) }
func (t *ToolItem) GetIconSize() IconSize             { return ToolItemGetIconSize(t) }
func (t *ToolItem) GetIsImportant() bool              { return ToolItemGetIsImportant(t) }
func (t *ToolItem) GetOrientation() Orientation       { return ToolItemGetOrientation(t) }
func (t *ToolItem) GetProxyMenuItem(menuItemId string) *Widget {
	return ToolItemGetProxyMenuItem(t, menuItemId)
}
func (t *ToolItem) GetReliefStyle() ReliefStyle     { return ToolItemGetReliefStyle(t) }
func (t *ToolItem) GetTextAlignment() float32       { return ToolItemGetTextAlignment(t) }
func (t *ToolItem) GetTextOrientation() Orientation { return ToolItemGetTextOrientation(t) }
func (t *ToolItem) GetTextSizeGroup() *SizeGroup    { return ToolItemGetTextSizeGroup(t) }
func (t *ToolItem) GetToolbarStyle() ToolbarStyle   { return ToolItemGetToolbarStyle(t) }
func (t *ToolItem) GetUseDragWindow() bool          { return ToolItemGetUseDragWindow(t) }
func (t *ToolItem) GetVisibleHorizontal() bool      { return ToolItemGetVisibleHorizontal(t) }
func (t *ToolItem) GetVisibleVertical() bool        { return ToolItemGetVisibleVertical(t) }
func (t *ToolItem) RebuildMenu()                    { ToolItemRebuildMenu(t) }
func (t *ToolItem) RetrieveProxyMenuItem() *Widget  { return ToolItemRetrieveProxyMenuItem(t) }
func (t *ToolItem) SetExpand(expand bool)           { ToolItemSetExpand(t, expand) }
func (t *ToolItem) SetHomogeneous(homogeneous bool) { ToolItemSetHomogeneous(t, homogeneous) }
func (t *ToolItem) SetIsImportant(isImportant bool) { ToolItemSetIsImportant(t, isImportant) }
func (t *ToolItem) SetProxyMenuItem(menuItemId string, menuItem *Widget) {
	ToolItemSetProxyMenuItem(t, menuItemId, menuItem)
}
func (t *ToolItem) SetTooltip(tooltips *Tooltips, tipText, tipPrivate string) {
	ToolItemSetTooltip(t, tooltips, tipText, tipPrivate)
}
func (t *ToolItem) SetTooltipMarkup(markup string) { ToolItemSetTooltipMarkup(t, markup) }
func (t *ToolItem) SetTooltipText(text string)     { ToolItemSetTooltipText(t, text) }
func (t *ToolItem) SetUseDragWindow(useDragWindow bool) {
	ToolItemSetUseDragWindow(t, useDragWindow)
}
func (t *ToolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	ToolItemSetVisibleHorizontal(t, visibleHorizontal)
}
func (t *ToolItem) SetVisibleVertical(visibleVertical bool) {
	ToolItemSetVisibleVertical(t, visibleVertical)
}
func (t *ToolItem) ToolbarReconfigured() { ToolItemToolbarReconfigured(t) }

type ToolItemGroup struct {
	Parent Container
	_      *struct{}
}

var (
	ToolItemGroupGetType func() O.Type
	ToolItemGroupNew     func(label string) *Widget

	ToolItemGroupGetCollapsed    func(t *ToolItemGroup) bool
	ToolItemGroupGetDropItem     func(t *ToolItemGroup, x, y int) *ToolItem
	ToolItemGroupGetEllipsize    func(t *ToolItemGroup) P.EllipsizeMode
	ToolItemGroupGetHeaderRelief func(t *ToolItemGroup) ReliefStyle
	ToolItemGroupGetItemPosition func(t *ToolItemGroup, item *ToolItem) int
	ToolItemGroupGetLabel        func(t *ToolItemGroup) string
	ToolItemGroupGetLabelWidget  func(t *ToolItemGroup) *Widget
	ToolItemGroupGetNItems       func(t *ToolItemGroup) uint
	ToolItemGroupGetNthItem      func(t *ToolItemGroup, index uint) *ToolItem
	ToolItemGroupInsert          func(t *ToolItemGroup, item *ToolItem, position int)
	ToolItemGroupSetCollapsed    func(t *ToolItemGroup, collapsed bool)
	ToolItemGroupSetEllipsize    func(t *ToolItemGroup, ellipsize P.EllipsizeMode)
	ToolItemGroupSetHeaderRelief func(t *ToolItemGroup, style ReliefStyle)
	ToolItemGroupSetItemPosition func(t *ToolItemGroup, item *ToolItem, position int)
	ToolItemGroupSetLabel        func(t *ToolItemGroup, label string)
	ToolItemGroupSetLabelWidget  func(t *ToolItemGroup, labelWidget *Widget)
)

func (t *ToolItemGroup) GetCollapsed() bool             { return ToolItemGroupGetCollapsed(t) }
func (t *ToolItemGroup) GetDropItem(x, y int) *ToolItem { return ToolItemGroupGetDropItem(t, x, y) }
func (t *ToolItemGroup) GetEllipsize() P.EllipsizeMode  { return ToolItemGroupGetEllipsize(t) }
func (t *ToolItemGroup) GetHeaderRelief() ReliefStyle   { return ToolItemGroupGetHeaderRelief(t) }
func (t *ToolItemGroup) GetItemPosition(item *ToolItem) int {
	return ToolItemGroupGetItemPosition(t, item)
}
func (t *ToolItemGroup) GetLabel() string                    { return ToolItemGroupGetLabel(t) }
func (t *ToolItemGroup) GetLabelWidget() *Widget             { return ToolItemGroupGetLabelWidget(t) }
func (t *ToolItemGroup) GetNItems() uint                     { return ToolItemGroupGetNItems(t) }
func (t *ToolItemGroup) GetNthItem(index uint) *ToolItem     { return ToolItemGroupGetNthItem(t, index) }
func (t *ToolItemGroup) Insert(item *ToolItem, position int) { ToolItemGroupInsert(t, item, position) }
func (t *ToolItemGroup) SetCollapsed(collapsed bool)         { ToolItemGroupSetCollapsed(t, collapsed) }
func (t *ToolItemGroup) SetEllipsize(ellipsize P.EllipsizeMode) {
	ToolItemGroupSetEllipsize(t, ellipsize)
}
func (t *ToolItemGroup) SetHeaderRelief(style ReliefStyle) {
	ToolItemGroupSetHeaderRelief(t, style)
}
func (t *ToolItemGroup) SetItemPosition(item *ToolItem, position int) {
	ToolItemGroupSetItemPosition(t, item, position)
}
func (t *ToolItemGroup) SetLabel(label string) { ToolItemGroupSetLabel(t, label) }
func (t *ToolItemGroup) SetLabelWidget(labelWidget *Widget) {
	ToolItemGroupSetLabelWidget(t, labelWidget)
}

type ToolPalette struct {
	Parent Container
	_      *struct{}
}

type ToolPaletteDragTargets Enum

const (
	TOOL_PALETTE_DRAG_ITEMS ToolPaletteDragTargets = 1 << iota
	TOOL_PALETTE_DRAG_GROUPS
)

var (
	ToolPaletteGetType func() O.Type
	ToolPaletteNew     func() *Widget

	ToolPaletteDragTargetsGetType func() O.Type
	ToolPaletteGetDragTargetGroup func() *TargetEntry
	ToolPaletteGetDragTargetItem  func() *TargetEntry

	ToolPaletteAddDragDest      func(t *ToolPalette, widget *Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions D.DragAction)
	ToolPaletteGetDragItem      func(t *ToolPalette, selection *SelectionData) *Widget
	ToolPaletteGetDropGroup     func(t *ToolPalette, x, y int) *ToolItemGroup
	ToolPaletteGetDropItem      func(t *ToolPalette, x, y int) *ToolItem
	ToolPaletteGetExclusive     func(t *ToolPalette, group *ToolItemGroup) bool
	ToolPaletteGetExpand        func(t *ToolPalette, group *ToolItemGroup) bool
	ToolPaletteGetGroupPosition func(t *ToolPalette, group *ToolItemGroup) int
	ToolPaletteGetHadjustment   func(t *ToolPalette) *Adjustment
	ToolPaletteGetIconSize      func(t *ToolPalette) IconSize
	ToolPaletteGetStyle         func(t *ToolPalette) ToolbarStyle
	ToolPaletteGetVadjustment   func(t *ToolPalette) *Adjustment
	ToolPaletteSetDragSource    func(t *ToolPalette, targets ToolPaletteDragTargets)
	ToolPaletteSetExclusive     func(t *ToolPalette, group *ToolItemGroup, exclusive bool)
	ToolPaletteSetExpand        func(t *ToolPalette, group *ToolItemGroup, expand bool)
	ToolPaletteSetGroupPosition func(t *ToolPalette, group *ToolItemGroup, position int)
	ToolPaletteSetIconSize      func(t *ToolPalette, iconSize IconSize)
	ToolPaletteSetStyle         func(t *ToolPalette, style ToolbarStyle)
	ToolPaletteUnsetIconSize    func(t *ToolPalette)
	ToolPaletteUnsetStyle       func(t *ToolPalette)
)

func (t *ToolPalette) AddDragDest(widget *Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions D.DragAction) {
	ToolPaletteAddDragDest(t, widget, flags, targets, actions)
}
func (t *ToolPalette) GetDragItem(selection *SelectionData) *Widget {
	return ToolPaletteGetDragItem(t, selection)
}
func (t *ToolPalette) GetDropGroup(x, y int) *ToolItemGroup { return ToolPaletteGetDropGroup(t, x, y) }
func (t *ToolPalette) GetDropItem(x, y int) *ToolItem       { return ToolPaletteGetDropItem(t, x, y) }
func (t *ToolPalette) GetExclusive(group *ToolItemGroup) bool {
	return ToolPaletteGetExclusive(t, group)
}
func (t *ToolPalette) GetExpand(group *ToolItemGroup) bool {
	return ToolPaletteGetExpand(t, group)
}
func (t *ToolPalette) GetGroupPosition(group *ToolItemGroup) int {
	return ToolPaletteGetGroupPosition(t, group)
}
func (t *ToolPalette) GetHadjustment() *Adjustment { return ToolPaletteGetHadjustment(t) }
func (t *ToolPalette) GetIconSize() IconSize       { return ToolPaletteGetIconSize(t) }
func (t *ToolPalette) GetStyle() ToolbarStyle      { return ToolPaletteGetStyle(t) }
func (t *ToolPalette) GetVadjustment() *Adjustment { return ToolPaletteGetVadjustment(t) }
func (t *ToolPalette) SetDragSource(targets ToolPaletteDragTargets) {
	ToolPaletteSetDragSource(t, targets)
}
func (t *ToolPalette) SetExclusive(group *ToolItemGroup, exclusive bool) {
	ToolPaletteSetExclusive(t, group, exclusive)
}
func (t *ToolPalette) SetExpand(group *ToolItemGroup, expand bool) {
	ToolPaletteSetExpand(t, group, expand)
}
func (t *ToolPalette) SetGroupPosition(group *ToolItemGroup, position int) {
	ToolPaletteSetGroupPosition(t, group, position)
}
func (t *ToolPalette) SetIconSize(iconSize IconSize) { ToolPaletteSetIconSize(t, iconSize) }
func (t *ToolPalette) SetStyle(style ToolbarStyle)   { ToolPaletteSetStyle(t, style) }
func (t *ToolPalette) UnsetIconSize()                { ToolPaletteUnsetIconSize(t) }
func (t *ToolPalette) UnsetStyle()                   { ToolPaletteUnsetStyle(t) }

type ToolShell struct{}

var (
	ToolShellGetType func() O.Type

	ToolShellGetEllipsizeMode   func(t *ToolShell) P.EllipsizeMode
	ToolShellGetIconSize        func(t *ToolShell) IconSize
	ToolShellGetOrientation     func(t *ToolShell) Orientation
	ToolShellGetReliefStyle     func(t *ToolShell) ReliefStyle
	ToolShellGetStyle           func(t *ToolShell) ToolbarStyle
	ToolShellGetTextAlignment   func(t *ToolShell) float32
	ToolShellGetTextOrientation func(t *ToolShell) Orientation
	ToolShellGetTextSizeGroup   func(t *ToolShell) *SizeGroup
	ToolShellRebuildMenu        func(t *ToolShell)
)

func (t *ToolShell) GetEllipsizeMode() P.EllipsizeMode { return ToolShellGetEllipsizeMode(t) }
func (t *ToolShell) GetIconSize() IconSize             { return ToolShellGetIconSize(t) }
func (t *ToolShell) GetOrientation() Orientation       { return ToolShellGetOrientation(t) }
func (t *ToolShell) GetReliefStyle() ReliefStyle       { return ToolShellGetReliefStyle(t) }
func (t *ToolShell) GetStyle() ToolbarStyle            { return ToolShellGetStyle(t) }
func (t *ToolShell) GetTextAlignment() float32         { return ToolShellGetTextAlignment(t) }
func (t *ToolShell) GetTextOrientation() Orientation   { return ToolShellGetTextOrientation(t) }
func (t *ToolShell) GetTextSizeGroup() *SizeGroup      { return ToolShellGetTextSizeGroup(t) }
func (t *ToolShell) RebuildMenu()                      { ToolShellRebuildMenu(t) }

type Tooltip struct{}

var (
	TooltipGetType func() O.Type

	TooltipTriggerTooltipQuery func(display *D.Display)

	TooltipSetCustom           func(t *Tooltip, customWidget *Widget)
	TooltipSetIcon             func(t *Tooltip, pixbuf *D.Pixbuf)
	TooltipSetIconFromGicon    func(t *Tooltip, gicon *I.Icon, size IconSize)
	TooltipSetIconFromIconName func(t *Tooltip, iconName string, size IconSize)
	TooltipSetIconFromStock    func(t *Tooltip, stockId string, size IconSize)
	TooltipSetMarkup           func(t *Tooltip, markup string)
	TooltipSetText             func(t *Tooltip, text string)
	TooltipSetTipArea          func(t *Tooltip, rect *D.Rectangle)
)

func (t *Tooltip) SetCustom(customWidget *Widget) { TooltipSetCustom(t, customWidget) }
func (t *Tooltip) SetIcon(pixbuf *D.Pixbuf)       { TooltipSetIcon(t, pixbuf) }
func (t *Tooltip) SetIconFromGicon(gicon *I.Icon, size IconSize) {
	TooltipSetIconFromGicon(t, gicon, size)
}
func (t *Tooltip) SetIconFromIconName(iconName string, size IconSize) {
	TooltipSetIconFromIconName(t, iconName, size)
}
func (t *Tooltip) SetIconFromStock(stockId string, size IconSize) {
	TooltipSetIconFromStock(t, stockId, size)
}
func (t *Tooltip) SetMarkup(markup string)      { TooltipSetMarkup(t, markup) }
func (t *Tooltip) SetText(text string)          { TooltipSetText(t, text) }
func (t *Tooltip) SetTipArea(rect *D.Rectangle) { TooltipSetTipArea(t, rect) }

type (
	Tooltips struct {
		Parent         Object
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
		LastPopdown L.TimeVal
	}

	TooltipsData struct {
		Tooltips   *Tooltips
		Widget     *Widget
		TipText    *T.Gchar
		TipPrivate *T.Gchar
	}
)

var (
	TooltipsGetType func() O.Type
	TooltipsNew     func() *Tooltips

	TooltipsDataGet              func(widget *Widget) *TooltipsData
	TooltipsGetInfoFromTipWindow func(tipWindow *Window, tooltips **Tooltips, currentWidget **Widget) bool

	TooltipsDisable     func(t *Tooltips)
	TooltipsEnable      func(t *Tooltips)
	TooltipsForceWindow func(t *Tooltips)
	TooltipsSetDelay    func(t *Tooltips, delay uint)
	TooltipsSetTip      func(t *Tooltips, widget *Widget, tipText, tipPrivate string)
)

func (t *Tooltips) Disable()            { TooltipsDisable(t) }
func (t *Tooltips) Enable()             { TooltipsEnable(t) }
func (t *Tooltips) ForceWindow()        { TooltipsForceWindow(t) }
func (t *Tooltips) SetDelay(delay uint) { TooltipsSetDelay(t, delay) }
func (t *Tooltips) SetTip(widget *Widget, tipText, tipPrivate string) {
	TooltipsSetTip(t, widget, tipText, tipPrivate)
}

type TranslateFunc func(path string, funcData T.Gpointer) string

type (
	Tree struct {
		Container     Container
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
	TreeGetType func() O.Type

	TreeAppend           func(t *Tree, treeItem *Widget)
	TreeChildPosition    func(t *Tree, child *Widget) int
	TreeClearItems       func(t *Tree, start int, end int)
	TreeInsert           func(t *Tree, treeItem *Widget, position int)
	TreePrepend          func(t *Tree, treeItem *Widget)
	TreeRemoveItem       func(t *Tree, child *Widget)
	TreeRemoveItems      func(t *Tree, items *T.GList)
	TreeSelectChild      func(t *Tree, treeItem *Widget)
	TreeSelectItem       func(t *Tree, item int)
	TreeSetSelectionMode func(t *Tree, mode SelectionMode)
	TreeSetViewLines     func(t *Tree, flag bool)
	TreeSetViewMode      func(t *Tree, mode TreeViewMode)
	TreeUnselectChild    func(t *Tree, treeItem *Widget)
	TreeUnselectItem     func(t *Tree, item int)
)

func (t *Tree) Append(treeItem *Widget)               { TreeAppend(t, treeItem) }
func (t *Tree) ChildPosition(child *Widget) int       { return TreeChildPosition(t, child) }
func (t *Tree) ClearItems(start, end int)             { TreeClearItems(t, start, end) }
func (t *Tree) Insert(treeItem *Widget, position int) { TreeInsert(t, treeItem, position) }
func (t *Tree) Prepend(treeItem *Widget)              { TreePrepend(t, treeItem) }
func (t *Tree) RemoveItem(child *Widget)              { TreeRemoveItem(t, child) }
func (t *Tree) RemoveItems(items *T.GList)            { TreeRemoveItems(t, items) }
func (t *Tree) SelectChild(treeItem *Widget)          { TreeSelectChild(t, treeItem) }
func (t *Tree) SelectItem(item int)                   { TreeSelectItem(t, item) }
func (t *Tree) SetSelectionMode(mode SelectionMode)   { TreeSetSelectionMode(t, mode) }
func (t *Tree) SetViewLines(flag bool)                { TreeSetViewLines(t, flag) }
func (t *Tree) SetViewMode(mode TreeViewMode)         { TreeSetViewMode(t, mode) }
func (t *Tree) UnselectChild(treeItem *Widget)        { TreeUnselectChild(t, treeItem) }
func (t *Tree) UnselectItem(item int)                 { TreeUnselectItem(t, item) }

type TreeDragSource struct{}

var (
	TreeDragSourceGetType func() O.Type

	TreeDragSourceDragDataDelete func(t *TreeDragSource, path *TreePath) bool
	TreeDragSourceDragDataGet    func(t *TreeDragSource, path *TreePath, selectionData *SelectionData) bool
	TreeDragSourceRowDraggable   func(t *TreeDragSource, path *TreePath) bool
)

func (t *TreeDragSource) DragDataDelete(path *TreePath) bool {
	return TreeDragSourceDragDataDelete(t, path)
}
func (t *TreeDragSource) DragDataGet(path *TreePath, selectionData *SelectionData) bool {
	return TreeDragSourceDragDataGet(t, path, selectionData)
}
func (t *TreeDragSource) RowDraggable(path *TreePath) bool {
	return TreeDragSourceRowDraggable(t, path)
}

type TreeDragDest struct{}

var (
	TreeDragDestGetType func() O.Type

	TreeDragDestDragDataReceived func(t *TreeDragDest, dest *TreePath, selectionData *SelectionData) bool
	TreeDragDestRowDropPossible  func(t *TreeDragDest, destPath *TreePath, selectionData *SelectionData) bool
)

func (t *TreeDragDest) DragDataReceived(dest *TreePath, selectionData *SelectionData) bool {
	return TreeDragDestDragDataReceived(t, dest, selectionData)
}
func (t *TreeDragDest) RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool {
	return TreeDragDestRowDropPossible(t, destPath, selectionData)
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
	TreeItemGetType      func() O.Type
	TreeItemNew          func() *Widget
	TreeItemNewWithLabel func(label string) *Widget

	TreeItemCollapse      func(t *TreeItem)
	TreeItemDeselect      func(t *TreeItem)
	TreeItemExpand        func(t *TreeItem)
	TreeItemRemoveSubtree func(t *TreeItem)
	TreeItemSelect        func(t *TreeItem)
	TreeItemSetSubtree    func(t *TreeItem, subtree *Widget)
)

func (t *TreeItem) Collapse()                  { TreeItemCollapse(t) }
func (t *TreeItem) Deselect()                  { TreeItemDeselect(t) }
func (t *TreeItem) Expand()                    { TreeItemExpand(t) }
func (t *TreeItem) RemoveSubtree()             { TreeItemRemoveSubtree(t) }
func (t *TreeItem) Select()                    { TreeItemSelect(t) }
func (t *TreeItem) SetSubtree(subtree *Widget) { TreeItemSetSubtree(t, subtree) }

type TreeIter struct {
	Stamp     int
	UserData  T.Gpointer
	UserData2 T.Gpointer
	UserData3 T.Gpointer
}

var (
	TreeIterGetType func() O.Type

	TreeIterCopy func(t *TreeIter) *TreeIter
	TreeIterFree func(t *TreeIter)
)

func (t *TreeIter) Copy() *TreeIter { return TreeIterCopy(t) }
func (t *TreeIter) Free()           { TreeIterFree(t) }

type (
	TreeModel struct{}

	TreeModelForeachFunc func(
		model *TreeModel,
		path *TreePath,
		iter *TreeIter,
		data T.Gpointer) bool
)

var (
	TreeModelGetType func() O.Type

	TreeGetRowDragData func(s *SelectionData, treeModel **TreeModel, path **TreePath) bool
	TreeSetRowDragData func(s *SelectionData, treeModel *TreeModel, path *TreePath) bool

	TreeModelForeach            func(t *TreeModel, f TreeModelForeachFunc, userData T.Gpointer)
	TreeModelGet                func(t *TreeModel, iter *TreeIter, v ...VArg)
	TreeModelGetColumnType      func(t *TreeModel, index int) O.Type
	TreeModelGetFlags           func(t *TreeModel) TreeModelFlags
	TreeModelGetIter            func(t *TreeModel, iter *TreeIter, path *TreePath) bool
	TreeModelGetIterFirst       func(t *TreeModel, iter *TreeIter) bool
	TreeModelGetIterFromString  func(t *TreeModel, iter *TreeIter, pathString string) bool
	TreeModelGetNColumns        func(t *TreeModel) int
	TreeModelGetPath            func(t *TreeModel, iter *TreeIter) *TreePath
	TreeModelGetStringFromIter  func(t *TreeModel, iter *TreeIter) string
	TreeModelGetValist          func(t *TreeModel, iter *TreeIter, varArgs T.VaList)
	TreeModelGetValue           func(t *TreeModel, iter *TreeIter, column int, value *T.GValue)
	TreeModelIterChildren       func(t *TreeModel, iter *TreeIter, parent *TreeIter) bool
	TreeModelIterHasChild       func(t *TreeModel, iter *TreeIter) bool
	TreeModelIterNChildren      func(t *TreeModel, iter *TreeIter) int
	TreeModelIterNext           func(t *TreeModel, iter *TreeIter) bool
	TreeModelIterNthChild       func(t *TreeModel, iter *TreeIter, parent *TreeIter, n int) bool
	TreeModelIterParent         func(t *TreeModel, iter *TreeIter, child *TreeIter) bool
	TreeModelRefNode            func(t *TreeModel, iter *TreeIter)
	TreeModelRowChanged         func(t *TreeModel, path *TreePath, iter *TreeIter)
	TreeModelRowDeleted         func(t *TreeModel, path *TreePath)
	TreeModelRowHasChildToggled func(t *TreeModel, path *TreePath, iter *TreeIter)
	TreeModelRowInserted        func(t *TreeModel, path *TreePath, iter *TreeIter)
	TreeModelRowsReordered      func(t *TreeModel, path *TreePath, iter *TreeIter, newOrder *int)
	TreeModelUnrefNode          func(t *TreeModel, iter *TreeIter)
)

func (t *TreeModel) Foreach(f TreeModelForeachFunc, userData T.Gpointer) {
	TreeModelForeach(t, f, userData)
}
func (t *TreeModel) Get(iter *TreeIter, v ...VArg)  { TreeModelGet(t, iter, v) }
func (t *TreeModel) GetColumnType(index int) O.Type { return TreeModelGetColumnType(t, index) }
func (t *TreeModel) GetFlags() TreeModelFlags       { return TreeModelGetFlags(t) }
func (t *TreeModel) GetIter(iter *TreeIter, path *TreePath) bool {
	return TreeModelGetIter(t, iter, path)
}
func (t *TreeModel) GetIterFirst(iter *TreeIter) bool { return TreeModelGetIterFirst(t, iter) }
func (t *TreeModel) GetIterFromString(iter *TreeIter, pathString string) bool {
	return TreeModelGetIterFromString(t, iter, pathString)
}
func (t *TreeModel) GetNColumns() int                 { return TreeModelGetNColumns(t) }
func (t *TreeModel) GetPath(iter *TreeIter) *TreePath { return TreeModelGetPath(t, iter) }
func (t *TreeModel) GetStringFromIter(iter *TreeIter) string {
	return TreeModelGetStringFromIter(t, iter)
}
func (t *TreeModel) GetValist(iter *TreeIter, varArgs T.VaList) { TreeModelGetValist(t, iter, varArgs) }
func (t *TreeModel) GetValue(iter *TreeIter, column int, value *T.GValue) {
	TreeModelGetValue(t, iter, column, value)
}
func (t *TreeModel) IterChildren(iter *TreeIter, parent *TreeIter) bool {
	return TreeModelIterChildren(t, iter, parent)
}
func (t *TreeModel) IterHasChild(iter *TreeIter) bool { return TreeModelIterHasChild(t, iter) }
func (t *TreeModel) IterNChildren(iter *TreeIter) int { return TreeModelIterNChildren(t, iter) }
func (t *TreeModel) IterNext(iter *TreeIter) bool     { return TreeModelIterNext(t, iter) }
func (t *TreeModel) IterNthChild(iter, parent *TreeIter, n int) bool {
	return TreeModelIterNthChild(t, iter, parent, n)
}
func (t *TreeModel) IterParent(iter, child *TreeIter) bool {
	return TreeModelIterParent(t, iter, child)
}
func (t *TreeModel) RefNode(iter *TreeIter) { TreeModelRefNode(t, iter) }
func (t *TreeModel) RowChanged(path *TreePath, iter *TreeIter) {
	TreeModelRowChanged(t, path, iter)
}
func (t *TreeModel) RowDeleted(path *TreePath) { TreeModelRowDeleted(t, path) }
func (t *TreeModel) RowHasChildToggled(path *TreePath, iter *TreeIter) {
	TreeModelRowHasChildToggled(t, path, iter)
}
func (t *TreeModel) RowInserted(path *TreePath, iter *TreeIter) {
	TreeModelRowInserted(t, path, iter)
}
func (t *TreeModel) RowsReordered(path *TreePath, iter *TreeIter, newOrder *int) {
	TreeModelRowsReordered(t, path, iter, newOrder)
}
func (t *TreeModel) UnrefNode(iter *TreeIter) { TreeModelUnrefNode(t, iter) }

type (
	TreeModelFilter struct {
		Parent O.Object
		_      *struct{}
	}

	TreeModelFilterVisibleFunc func(
		model *TreeModel,
		iter *TreeIter,
		data T.Gpointer) bool

	TreeModelFilterModifyFunc func(
		model *TreeModel,
		iter *TreeIter,
		value *T.GValue,
		column int,
		data T.Gpointer)
)

var (
	TreeModelFilterGetType func() O.Type
	TreeModelFilterNew     func(childModel *TreeModel, root *TreePath) *TreeModel

	TreeModelFilterClearCache             func(t *TreeModelFilter)
	TreeModelFilterConvertChildIterToIter func(t *TreeModelFilter, filterIter *TreeIter, childIter *TreeIter) bool
	TreeModelFilterConvertChildPathToPath func(t *TreeModelFilter, childPath *TreePath) *TreePath
	TreeModelFilterConvertIterToChildIter func(t *TreeModelFilter, childIter *TreeIter, filterIter *TreeIter)
	TreeModelFilterConvertPathToChildPath func(t *TreeModelFilter, filterPath *TreePath) *TreePath
	TreeModelFilterGetModel               func(t *TreeModelFilter) *TreeModel
	TreeModelFilterRefilter               func(t *TreeModelFilter)
	TreeModelFilterSetModifyFunc          func(t *TreeModelFilter, nColumns int, types *O.Type, f TreeModelFilterModifyFunc, data T.Gpointer, destroy T.GDestroyNotify)
	TreeModelFilterSetVisibleColumn       func(t *TreeModelFilter, column int)
	TreeModelFilterSetVisibleFunc         func(t *TreeModelFilter, f TreeModelFilterVisibleFunc, data T.Gpointer, destroy T.GDestroyNotify)
)

func (t *TreeModelFilter) ClearCache() { TreeModelFilterClearCache(t) }
func (t *TreeModelFilter) ConvertChildIterToIter(filterIter *TreeIter, childIter *TreeIter) bool {
	return TreeModelFilterConvertChildIterToIter(t, filterIter, childIter)
}
func (t *TreeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	return TreeModelFilterConvertChildPathToPath(t, childPath)
}
func (t *TreeModelFilter) ConvertIterToChildIter(childIter *TreeIter, filterIter *TreeIter) {
	TreeModelFilterConvertIterToChildIter(t, childIter, filterIter)
}
func (t *TreeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	return TreeModelFilterConvertPathToChildPath(t, filterPath)
}
func (t *TreeModelFilter) GetModel() *TreeModel { return TreeModelFilterGetModel(t) }
func (t *TreeModelFilter) Refilter()            { TreeModelFilterRefilter(t) }
func (t *TreeModelFilter) SetModifyFunc(nColumns int, types *O.Type, f TreeModelFilterModifyFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	TreeModelFilterSetModifyFunc(t, nColumns, types, f, data, destroy)
}
func (t *TreeModelFilter) SetVisibleColumn(column int) { TreeModelFilterSetVisibleColumn(t, column) }
func (t *TreeModelFilter) SetVisibleFunc(f TreeModelFilterVisibleFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	TreeModelFilterSetVisibleFunc(t, f, data, destroy)
}

type TreeModelFlags Enum

const (
	TREE_MODEL_ITERS_PERSIST TreeModelFlags = 1 << iota
	TREE_MODEL_LIST_ONLY
)

var TreeModelFlagsGetType func() O.Type

type TreeModelSort struct {
	Parent             O.Object
	Root               T.Gpointer
	Stamp              int
	ChildFlags         uint
	ChildModel         *TreeModel
	ZeroRefCount       int
	SortList           *T.GList
	SortColumnId       int
	Order              SortType
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
	TreeModelSortGetType      func() O.Type
	TreeModelSortNewWithModel func(childModel *TreeModel) *TreeModel

	TreeModelSortClearCache             func(t *TreeModelSort)
	TreeModelSortConvertChildIterToIter func(t *TreeModelSort, sortIter *TreeIter, childIter *TreeIter) bool
	TreeModelSortConvertChildPathToPath func(t *TreeModelSort, childPath *TreePath) *TreePath
	TreeModelSortConvertIterToChildIter func(t *TreeModelSort, childIter *TreeIter, sortedIter *TreeIter)
	TreeModelSortConvertPathToChildPath func(t *TreeModelSort, sortedPath *TreePath) *TreePath
	TreeModelSortGetModel               func(t *TreeModelSort) *TreeModel
	TreeModelSortIterIsValid            func(t *TreeModelSort, iter *TreeIter) bool
	TreeModelSortResetDefaultSortFunc   func(t *TreeModelSort)
)

func (t *TreeModelSort) ClearCache() { TreeModelSortClearCache(t) }
func (t *TreeModelSort) ConvertChildIterToIter(sortIter *TreeIter, childIter *TreeIter) bool {
	return TreeModelSortConvertChildIterToIter(t, sortIter, childIter)
}
func (t *TreeModelSort) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	return TreeModelSortConvertChildPathToPath(t, childPath)
}
func (t *TreeModelSort) ConvertIterToChildIter(childIter *TreeIter, sortedIter *TreeIter) {
	TreeModelSortConvertIterToChildIter(t, childIter, sortedIter)
}
func (t *TreeModelSort) ConvertPathToChildPath(sortedPath *TreePath) *TreePath {
	return TreeModelSortConvertPathToChildPath(t, sortedPath)
}
func (t *TreeModelSort) GetModel() *TreeModel { return TreeModelSortGetModel(t) }
func (t *TreeModelSort) IterIsValid(iter *TreeIter) bool {
	return TreeModelSortIterIsValid(t, iter)
}
func (t *TreeModelSort) ResetDefaultSortFunc() { TreeModelSortResetDefaultSortFunc(t) }

type TreePath struct{}

var (
	TreePathGetType        func() O.Type
	TreePathNew            func() *TreePath
	TreePathNewFirst       func() *TreePath
	TreePathNewFromString  func(path string) *TreePath
	TreePathNewFromIndices func(firstIndex int, v ...VArg) *TreePath

	TreePathAppendIndex         func(t *TreePath, index int)
	TreePathCompare             func(t *TreePath, p2 *TreePath) int
	TreePathCopy                func(t *TreePath) *TreePath
	TreePathDown                func(t *TreePath)
	TreePathFree                func(t *TreePath)
	TreePathGetDepth            func(t *TreePath) int
	TreePathGetIndices          func(t *TreePath) *int
	TreePathGetIndicesWithDepth func(t *TreePath, depth *int) *int
	TreePathIsAncestor          func(t *TreePath, descendant *TreePath) bool
	TreePathIsDescendant        func(t *TreePath, ancestor *TreePath) bool
	TreePathNext                func(t *TreePath)
	TreePathPrependIndex        func(t *TreePath, index int)
	TreePathPrev                func(t *TreePath) bool
	TreePathToString            func(t *TreePath) string
	TreePathUp                  func(t *TreePath) bool
)

func (t *TreePath) AppendIndex(index int)               { TreePathAppendIndex(t, index) }
func (t *TreePath) Compare(p2 *TreePath) int            { return TreePathCompare(t, p2) }
func (t *TreePath) Copy() *TreePath                     { return TreePathCopy(t) }
func (t *TreePath) Down()                               { TreePathDown(t) }
func (t *TreePath) Free()                               { TreePathFree(t) }
func (t *TreePath) GetDepth() int                       { return TreePathGetDepth(t) }
func (t *TreePath) GetIndices() *int                    { return TreePathGetIndices(t) }
func (t *TreePath) GetIndicesWithDepth(depth *int) *int { return TreePathGetIndicesWithDepth(t, depth) }
func (t *TreePath) IsAncestor(descendant *TreePath) bool {
	return TreePathIsAncestor(t, descendant)
}
func (t *TreePath) IsDescendant(ancestor *TreePath) bool {
	return TreePathIsDescendant(t, ancestor)
}
func (t *TreePath) Next()                  { TreePathNext(t) }
func (t *TreePath) PrependIndex(index int) { TreePathPrependIndex(t, index) }
func (t *TreePath) Prev() bool             { return TreePathPrev(t) }
func (t *TreePath) ToString() string       { return TreePathToString(t) }
func (t *TreePath) Up() bool               { return TreePathUp(t) }

type TreeRowReference struct{}

var (
	TreeRowReferenceGetType func() O.Type
	TreeRowReferenceNew     func(model *TreeModel, path *TreePath) *TreeRowReference

	TreeRowReferenceDeleted   func(proxy *O.Object, path *TreePath)
	TreeRowReferenceInserted  func(proxy *O.Object, path *TreePath)
	TreeRowReferenceNewProxy  func(proxy *O.Object, model *TreeModel, path *TreePath) *TreeRowReference
	TreeRowReferenceReordered func(proxy *O.Object, path *TreePath, iter *TreeIter, newOrder *int)

	TreeRowReferenceCopy     func(t *TreeRowReference) *TreeRowReference
	TreeRowReferenceFree     func(t *TreeRowReference)
	TreeRowReferenceGetModel func(t *TreeRowReference) *TreeModel
	TreeRowReferenceGetPath  func(t *TreeRowReference) *TreePath
	TreeRowReferenceValid    func(t *TreeRowReference) bool
)

func (t *TreeRowReference) Copy() *TreeRowReference { return TreeRowReferenceCopy(t) }
func (t *TreeRowReference) Free()                   { TreeRowReferenceFree(t) }
func (t *TreeRowReference) GetModel() *TreeModel    { return TreeRowReferenceGetModel(t) }
func (t *TreeRowReference) GetPath() *TreePath      { return TreeRowReferenceGetPath(t) }
func (t *TreeRowReference) Valid() bool             { return TreeRowReferenceValid(t) }

type (
	TreeSelection struct {
		Parent   O.Object
		TreeView *TreeView
		Type     SelectionMode
		UserFunc TreeSelectionFunc
		UserData T.Gpointer
		Destroy  T.GDestroyNotify
	}

	TreeSelectionForeachFunc func(model *TreeModel,
		path *TreePath, iter *TreeIter, data T.Gpointer)
)

var (
	TreeSelectionGetType func() O.Type

	TreeSelectionCountSelectedRows func(t *TreeSelection) int
	TreeSelectionGetMode           func(t *TreeSelection) SelectionMode
	TreeSelectionGetSelected       func(t *TreeSelection, model **TreeModel, iter *TreeIter) bool
	TreeSelectionGetSelectedRows   func(t *TreeSelection, model **TreeModel) *T.GList
	TreeSelectionGetSelectFunction func(t *TreeSelection) TreeSelectionFunc
	TreeSelectionGetTreeView       func(t *TreeSelection) *TreeView
	TreeSelectionGetUserData       func(t *TreeSelection) T.Gpointer
	TreeSelectionIterIsSelected    func(t *TreeSelection, iter *TreeIter) bool
	TreeSelectionPathIsSelected    func(t *TreeSelection, path *TreePath) bool
	TreeSelectionSelectAll         func(t *TreeSelection)
	TreeSelectionSelectedForeach   func(t *TreeSelection, f TreeSelectionForeachFunc, data T.Gpointer)
	TreeSelectionSelectIter        func(t *TreeSelection, iter *TreeIter)
	TreeSelectionSelectPath        func(t *TreeSelection, path *TreePath)
	TreeSelectionSelectRange       func(t *TreeSelection, startPath *TreePath, endPath *TreePath)
	TreeSelectionSetMode           func(t *TreeSelection, typ SelectionMode)
	TreeSelectionSetSelectFunction func(t *TreeSelection, f TreeSelectionFunc, data T.Gpointer, destroy T.GDestroyNotify)
	TreeSelectionUnselectAll       func(t *TreeSelection)
	TreeSelectionUnselectIter      func(t *TreeSelection, iter *TreeIter)
	TreeSelectionUnselectPath      func(t *TreeSelection, path *TreePath)
	TreeSelectionUnselectRange     func(t *TreeSelection, startPath *TreePath, endPath *TreePath)
)

func (t *TreeSelection) CountSelectedRows() int { return TreeSelectionCountSelectedRows(t) }
func (t *TreeSelection) GetMode() SelectionMode { return TreeSelectionGetMode(t) }
func (t *TreeSelection) GetSelected(model **TreeModel, iter *TreeIter) bool {
	return TreeSelectionGetSelected(t, model, iter)
}
func (t *TreeSelection) GetSelectedRows(model **TreeModel) *T.GList {
	return TreeSelectionGetSelectedRows(t, model)
}
func (t *TreeSelection) GetSelectFunction() TreeSelectionFunc {
	return TreeSelectionGetSelectFunction(t)
}
func (t *TreeSelection) GetTreeView() *TreeView  { return TreeSelectionGetTreeView(t) }
func (t *TreeSelection) GetUserData() T.Gpointer { return TreeSelectionGetUserData(t) }
func (t *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	return TreeSelectionIterIsSelected(t, iter)
}
func (t *TreeSelection) PathIsSelected(path *TreePath) bool {
	return TreeSelectionPathIsSelected(t, path)
}
func (t *TreeSelection) SelectAll() { TreeSelectionSelectAll(t) }
func (t *TreeSelection) SelectedForeach(f TreeSelectionForeachFunc, data T.Gpointer) {
	TreeSelectionSelectedForeach(t, f, data)
}
func (t *TreeSelection) SelectIter(iter *TreeIter) { TreeSelectionSelectIter(t, iter) }
func (t *TreeSelection) SelectPath(path *TreePath) { TreeSelectionSelectPath(t, path) }
func (t *TreeSelection) SelectRange(startPath, endPath *TreePath) {
	TreeSelectionSelectRange(t, startPath, endPath)
}
func (t *TreeSelection) SetMode(typ SelectionMode) { TreeSelectionSetMode(t, typ) }
func (t *TreeSelection) SetSelectFunction(f TreeSelectionFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	TreeSelectionSetSelectFunction(t, f, data, destroy)
}
func (t *TreeSelection) UnselectAll()                { TreeSelectionUnselectAll(t) }
func (t *TreeSelection) UnselectIter(iter *TreeIter) { TreeSelectionUnselectIter(t, iter) }
func (t *TreeSelection) UnselectPath(path *TreePath) { TreeSelectionUnselectPath(t, path) }
func (t *TreeSelection) UnselectRange(startPath, endPath *TreePath) {
	TreeSelectionUnselectRange(t, startPath, endPath)
}

type TreeSortable struct{}

var (
	TreeSortableGetType func() O.Type

	TreeSortableGetSortColumnId    func(t *TreeSortable, sortColumnId *int, order *SortType) bool
	TreeSortableHasDefaultSortFunc func(t *TreeSortable) bool
	TreeSortableSetDefaultSortFunc func(t *TreeSortable, sortFunc TreeIterCompareFunc, userData T.Gpointer, destroy T.GDestroyNotify)
	TreeSortableSetSortColumnId    func(t *TreeSortable, sortColumnId int, order SortType)
	TreeSortableSetSortFunc        func(t *TreeSortable, sortColumnId int, sortFunc TreeIterCompareFunc, userData T.Gpointer, destroy T.GDestroyNotify)
	TreeSortableSortColumnChanged  func(t *TreeSortable)
)

func (t *TreeSortable) GetSortColumnId(sortColumnId *int, order *SortType) bool {
	return TreeSortableGetSortColumnId(t, sortColumnId, order)
}
func (t *TreeSortable) HasDefaultSortFunc() bool { return TreeSortableHasDefaultSortFunc(t) }
func (t *TreeSortable) SetDefaultSortFunc(sortFunc TreeIterCompareFunc, userData T.Gpointer, destroy T.GDestroyNotify) {
	TreeSortableSetDefaultSortFunc(t, sortFunc, userData, destroy)
}
func (t *TreeSortable) SetSortColumnId(sortColumnId int, order SortType) {
	TreeSortableSetSortColumnId(t, sortColumnId, order)
}
func (t *TreeSortable) SetSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc, userData T.Gpointer, destroy T.GDestroyNotify) {
	TreeSortableSetSortFunc(t, sortColumnId, sortFunc, userData, destroy)
}
func (t *TreeSortable) SortColumnChanged() { TreeSortableSortColumnChanged(t) }

type TreeStore struct {
	Parent             O.Object
	Stamp              int
	Root               T.Gpointer
	Last               T.Gpointer
	NColumns           int
	SortColumnId       int
	SortList           *T.GList
	Order              SortType
	ColumnHeaders      *O.Type
	DefaultSortFunc    TreeIterCompareFunc
	DefaultSortData    T.Gpointer
	DefaultSortDestroy T.GDestroyNotify
	ColumnsDirty       uint // : 1
}

var (
	TreeStoreGetType func() O.Type
	TreeStoreNew     func(nColumns int, v ...VArg) *TreeStore
	TreeStoreNewv    func(nColumns int, types *O.Type) *TreeStore

	TreeStoreAppend            func(t *TreeStore, iter *TreeIter, parent *TreeIter)
	TreeStoreClear             func(t *TreeStore)
	TreeStoreInsert            func(t *TreeStore, iter *TreeIter, parent *TreeIter, position int)
	TreeStoreInsertAfter       func(t *TreeStore, iter *TreeIter, parent *TreeIter, sibling *TreeIter)
	TreeStoreInsertBefore      func(t *TreeStore, iter *TreeIter, parent *TreeIter, sibling *TreeIter)
	TreeStoreInsertWithValues  func(t *TreeStore, iter *TreeIter, parent *TreeIter, position int, v ...VArg)
	TreeStoreInsertWithValuesv func(t *TreeStore, iter *TreeIter, parent *TreeIter, position int, columns *int, values *T.GValue, nValues int)
	TreeStoreIsAncestor        func(t *TreeStore, iter *TreeIter, descendant *TreeIter) bool
	TreeStoreIterDepth         func(t *TreeStore, iter *TreeIter) int
	TreeStoreIterIsValid       func(t *TreeStore, iter *TreeIter) bool
	TreeStoreMoveAfter         func(t *TreeStore, iter *TreeIter, position *TreeIter)
	TreeStoreMoveBefore        func(t *TreeStore, iter *TreeIter, position *TreeIter)
	TreeStorePrepend           func(t *TreeStore, iter *TreeIter, parent *TreeIter)
	TreeStoreRemove            func(t *TreeStore, iter *TreeIter) bool
	TreeStoreReorder           func(t *TreeStore, parent *TreeIter, newOrder *int)
	TreeStoreSet               func(t *TreeStore, iter *TreeIter, v ...VArg)
	TreeStoreSetColumnTypes    func(t *TreeStore, nColumns int, types *O.Type)
	TreeStoreSetValist         func(t *TreeStore, iter *TreeIter, varArgs T.VaList)
	TreeStoreSetValue          func(t *TreeStore, iter *TreeIter, column int, value *T.GValue)
	TreeStoreSetValuesv        func(t *TreeStore, iter *TreeIter, columns *int, values *T.GValue, nValues int)
	TreeStoreSwap              func(t *TreeStore, a *TreeIter, b *TreeIter)
)

func (t *TreeStore) Append(iter, parent *TreeIter) { TreeStoreAppend(t, iter, parent) }
func (t *TreeStore) Clear()                        { TreeStoreClear(t) }
func (t *TreeStore) Insert(iter, parent *TreeIter, position int) {
	TreeStoreInsert(t, iter, parent, position)
}
func (t *TreeStore) InsertAfter(iter, parent, sibling *TreeIter) {
	TreeStoreInsertAfter(t, iter, parent, sibling)
}
func (t *TreeStore) InsertBefore(iter, parent, sibling *TreeIter) {
	TreeStoreInsertBefore(t, iter, parent, sibling)
}
func (t *TreeStore) InsertWithValues(iter, parent *TreeIter, position int, v ...VArg) {
	TreeStoreInsertWithValues(t, iter, parent, position, v)
}
func (t *TreeStore) InsertWithValuesv(iter, parent *TreeIter, position int, columns *int, values *T.GValue, nValues int) {
	TreeStoreInsertWithValuesv(t, iter, parent, position, columns, values, nValues)
}
func (t *TreeStore) IsAncestor(iter, descendant *TreeIter) bool {
	return TreeStoreIsAncestor(t, iter, descendant)
}
func (t *TreeStore) IterDepth(iter *TreeIter) int            { return TreeStoreIterDepth(t, iter) }
func (t *TreeStore) IterIsValid(iter *TreeIter) bool         { return TreeStoreIterIsValid(t, iter) }
func (t *TreeStore) MoveAfter(iter, position *TreeIter)      { TreeStoreMoveAfter(t, iter, position) }
func (t *TreeStore) MoveBefore(iter, position *TreeIter)     { TreeStoreMoveBefore(t, iter, position) }
func (t *TreeStore) Prepend(iter, parent *TreeIter)          { TreeStorePrepend(t, iter, parent) }
func (t *TreeStore) Remove(iter *TreeIter) bool              { return TreeStoreRemove(t, iter) }
func (t *TreeStore) Reorder(parent *TreeIter, newOrder *int) { TreeStoreReorder(t, parent, newOrder) }
func (t *TreeStore) Set(iter *TreeIter, v ...VArg)           { TreeStoreSet(t, iter, v) }
func (t *TreeStore) SetColumnTypes(nColumns int, types *O.Type) {
	TreeStoreSetColumnTypes(t, nColumns, types)
}
func (t *TreeStore) SetValist(iter *TreeIter, varArgs T.VaList) { TreeStoreSetValist(t, iter, varArgs) }
func (t *TreeStore) SetValue(iter *TreeIter, column int, value *T.GValue) {
	TreeStoreSetValue(t, iter, column, value)
}
func (t *TreeStore) SetValuesv(iter *TreeIter, columns *int, values *T.GValue, nValues int) {
	TreeStoreSetValuesv(t, iter, columns, values, nValues)
}
func (t *TreeStore) Swap(a, b *TreeIter) { TreeStoreSwap(t, a, b) }

type (
	TreeView struct {
		Parent Container
		_      *struct{}
	}

	TreeViewRowSeparatorFunc func(
		model *TreeModel,
		iter *TreeIter,
		data T.Gpointer) bool

	TreeViewSearchEqualFunc func(
		model *TreeModel,
		column int,
		key string,
		iter *TreeIter,
		searchData T.Gpointer) bool

	TreeViewSearchPositionFunc func(
		tree_view *TreeView,
		search_dialog *Widget,
		user_data T.Gpointer)

	TreeSelectionFunc func(selection *TreeSelection,
		model *TreeModel, path *TreePath,
		pathCurrentlySelected bool, data T.Gpointer) bool

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
		data T.Gpointer) bool

	TreeDestroyCountFunc func(
		treeView *TreeView,
		path *TreePath,
		children int,
		userData T.Gpointer)

	TreeIterCompareFunc func(model *TreeModel,
		a, b *TreeIter, userData T.Gpointer) int
)

type TreeViewDropPosition Enum

const (
	TREE_VIEW_DROP_BEFORE TreeViewDropPosition = iota
	TREE_VIEW_DROP_AFTER
	TREE_VIEW_DROP_INTO_OR_BEFORE
	TREE_VIEW_DROP_INTO_OR_AFTER
)

type TreeViewGridLines Enum

const (
	TREE_VIEW_GRID_LINES_NONE TreeViewGridLines = iota
	TREE_VIEW_GRID_LINES_HORIZONTAL
	TREE_VIEW_GRID_LINES_VERTICAL
	TREE_VIEW_GRID_LINES_BOTH
)

type TreeViewColumnSizing Enum

const (
	TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = iota
	TREE_VIEW_COLUMN_AUTOSIZE
	TREE_VIEW_COLUMN_FIXED
)

type TreeViewMode Enum

const (
	TREE_VIEW_LINE TreeViewMode = iota
	TREE_VIEW_ITEM
)

var (
	TreeViewGetType      func() O.Type
	TreeViewNew          func() *Widget
	TreeViewNewWithModel func(model *TreeModel) *Widget

	TreeViewDropPositionGetType func() O.Type
	TreeViewGridLinesGetType    func() O.Type
	TreeViewModeGetType         func() O.Type

	TreeViewAppendColumn                   func(t *TreeView, column *TreeViewColumn) int
	TreeViewCollapseAll                    func(t *TreeView)
	TreeViewCollapseRow                    func(t *TreeView, path *TreePath) bool
	TreeViewColumnsAutosize                func(t *TreeView)
	TreeViewConvertBinWindowToTreeCoords   func(t *TreeView, bx, by int, tx, ty *int)
	TreeViewConvertBinWindowToWidgetCoords func(t *TreeView, bx, by int, wx, wy *int)
	TreeViewConvertTreeToBinWindowCoords   func(t *TreeView, tx, ty int, bx, by *int)
	TreeViewConvertTreeToWidgetCoords      func(t *TreeView, tx, ty int, wx, wy *int)
	TreeViewConvertWidgetToBinWindowCoords func(t *TreeView, wx, wy int, bx, by *int)
	TreeViewConvertWidgetToTreeCoords      func(t *TreeView, wx, wy int, tx, ty *int)
	TreeViewCreateRowDragIcon              func(t *TreeView, path *TreePath) *D.Pixmap
	TreeViewEnableModelDragDest            func(t *TreeView, targets *TargetEntry, nTargets int, actions D.DragAction)
	TreeViewEnableModelDragSource          func(t *TreeView, startButtonMask T.GdkModifierType, targets *TargetEntry, nTargets int, actions D.DragAction)
	TreeViewExpandAll                      func(t *TreeView)
	TreeViewExpandRow                      func(t *TreeView, path *TreePath, openAll bool) bool
	TreeViewExpandToPath                   func(t *TreeView, path *TreePath)
	TreeViewGetBackgroundArea              func(t *TreeView, path *TreePath, column *TreeViewColumn, rect *D.Rectangle)
	TreeViewGetBinWindow                   func(t *TreeView) *D.Window
	TreeViewGetCellArea                    func(t *TreeView, path *TreePath, column *TreeViewColumn, rect *D.Rectangle)
	TreeViewGetColumn                      func(t *TreeView, n int) *TreeViewColumn
	TreeViewGetColumns                     func(t *TreeView) *T.GList
	TreeViewGetCursor                      func(t *TreeView, path **TreePath, focusColumn **TreeViewColumn)
	TreeViewGetDestRowAtPos                func(t *TreeView, dragX, dragY int, path **TreePath, pos *TreeViewDropPosition) bool
	TreeViewGetDragDestRow                 func(t *TreeView, path **TreePath, pos *TreeViewDropPosition)
	TreeViewGetEnableSearch                func(t *TreeView) bool
	TreeViewGetEnableTreeLines             func(t *TreeView) bool
	TreeViewGetExpanderColumn              func(t *TreeView) *TreeViewColumn
	TreeViewGetFixedHeightMode             func(t *TreeView) bool
	TreeViewGetGridLines                   func(t *TreeView) TreeViewGridLines
	TreeViewGetHadjustment                 func(t *TreeView) *Adjustment
	TreeViewGetHeadersClickable            func(t *TreeView) bool
	TreeViewGetHeadersVisible              func(t *TreeView) bool
	TreeViewGetHoverExpand                 func(t *TreeView) bool
	TreeViewGetHoverSelection              func(t *TreeView) bool
	TreeViewGetLevelIndentation            func(t *TreeView) int
	TreeViewGetModel                       func(t *TreeView) *TreeModel
	TreeViewGetPathAtPos                   func(t *TreeView, x, y int, path **TreePath, column **TreeViewColumn, cellX, cellY *int) bool
	TreeViewGetReorderable                 func(t *TreeView) bool
	TreeViewGetRowSeparatorFunc            func(t *TreeView) TreeViewRowSeparatorFunc
	TreeViewGetRubberBanding               func(t *TreeView) bool
	TreeViewGetRulesHint                   func(t *TreeView) bool
	TreeViewGetSearchColumn                func(t *TreeView) int
	TreeViewGetSearchEntry                 func(t *TreeView) *Entry
	TreeViewGetSearchEqualFunc             func(t *TreeView) TreeViewSearchEqualFunc
	TreeViewGetSearchPositionFunc          func(t *TreeView) TreeViewSearchPositionFunc
	TreeViewGetSelection                   func(t *TreeView) *TreeSelection
	TreeViewGetShowExpanders               func(t *TreeView) bool
	TreeViewGetTooltipColumn               func(t *TreeView) int
	TreeViewGetTooltipContext              func(t *TreeView, x, y *int, keyboardTip bool, model **TreeModel, path **TreePath, iter *TreeIter) bool
	TreeViewGetVadjustment                 func(t *TreeView) *Adjustment
	TreeViewGetVisibleRange                func(t *TreeView, startPath **TreePath, endPath **TreePath) bool
	TreeViewGetVisibleRect                 func(t *TreeView, visibleRect *D.Rectangle)
	TreeViewInsertColumn                   func(t *TreeView, column *TreeViewColumn, position int) int
	TreeViewInsertColumnWithAttributes     func(t *TreeView, position int, title string, cell *CellRenderer, v ...VArg) int
	TreeViewInsertColumnWithDataFunc       func(t *TreeView, position int, title string, cell *CellRenderer, f TreeCellDataFunc, data T.Gpointer, dnotify T.GDestroyNotify) int
	TreeViewIsRubberBandingActive          func(t *TreeView) bool
	TreeViewMapExpandedRows                func(t *TreeView, f TreeViewMappingFunc, data T.Gpointer)
	TreeViewMoveColumnAfter                func(t *TreeView, column *TreeViewColumn, baseColumn *TreeViewColumn)
	TreeViewRemoveColumn                   func(t *TreeView, column *TreeViewColumn) int
	TreeViewRowActivated                   func(t *TreeView, path *TreePath, column *TreeViewColumn)
	TreeViewRowExpanded                    func(t *TreeView, path *TreePath) bool
	TreeViewScrollToCell                   func(t *TreeView, path *TreePath, column *TreeViewColumn, useAlign bool, rowAlign, colAlign float32)
	TreeViewScrollToPoint                  func(t *TreeView, treeX, treeY int)
	TreeViewSetColumnDragFunction          func(t *TreeView, f TreeViewColumnDropFunc, userData T.Gpointer, destroy T.GDestroyNotify)
	TreeViewSetCursor                      func(t *TreeView, path *TreePath, focusColumn *TreeViewColumn, startEditing bool)
	TreeViewSetCursorOnCell                func(t *TreeView, path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing bool)
	TreeViewSetDestroyCountFunc            func(t *TreeView, f TreeDestroyCountFunc, data T.Gpointer, destroy T.GDestroyNotify)
	TreeViewSetDragDestRow                 func(t *TreeView, path *TreePath, pos TreeViewDropPosition)
	TreeViewSetEnableSearch                func(t *TreeView, enableSearch bool)
	TreeViewSetEnableTreeLines             func(t *TreeView, enabled bool)
	TreeViewSetExpanderColumn              func(t *TreeView, column *TreeViewColumn)
	TreeViewSetFixedHeightMode             func(t *TreeView, enable bool)
	TreeViewSetGridLines                   func(t *TreeView, gridLines TreeViewGridLines)
	TreeViewSetHadjustment                 func(t *TreeView, adjustment *Adjustment)
	TreeViewSetHeadersClickable            func(t *TreeView, setting bool)
	TreeViewSetHeadersVisible              func(t *TreeView, headersVisible bool)
	TreeViewSetHoverExpand                 func(t *TreeView, expand bool)
	TreeViewSetHoverSelection              func(t *TreeView, hover bool)
	TreeViewSetLevelIndentation            func(t *TreeView, indentation int)
	TreeViewSetModel                       func(t *TreeView, model *TreeModel)
	TreeViewSetReorderable                 func(t *TreeView, reorderable bool)
	TreeViewSetRowSeparatorFunc            func(t *TreeView, f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify)
	TreeViewSetRubberBanding               func(t *TreeView, enable bool)
	TreeViewSetRulesHint                   func(t *TreeView, setting bool)
	TreeViewSetSearchColumn                func(t *TreeView, column int)
	TreeViewSetSearchEntry                 func(t *TreeView, entry *Entry)
	TreeViewSetSearchEqualFunc             func(t *TreeView, searchEqualFunc TreeViewSearchEqualFunc, searchUserData T.Gpointer, searchDestroy T.GDestroyNotify)
	TreeViewSetSearchPositionFunc          func(t *TreeView, f TreeViewSearchPositionFunc, data T.Gpointer, destroy T.GDestroyNotify)
	TreeViewSetShowExpanders               func(t *TreeView, enabled bool)
	TreeViewSetTooltipCell                 func(t *TreeView, tooltip *Tooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer)
	TreeViewSetTooltipColumn               func(t *TreeView, column int)
	TreeViewSetTooltipRow                  func(t *TreeView, tooltip *Tooltip, path *TreePath)
	TreeViewSetVadjustment                 func(t *TreeView, adjustment *Adjustment)
	TreeViewTreeToWidgetCoords             func(t *TreeView, tx, ty int, wx, wy *int)
	TreeViewUnsetRowsDragDest              func(t *TreeView)
	TreeViewUnsetRowsDragSource            func(t *TreeView)
	TreeViewWidgetToTreeCoords             func(t *TreeView, wx, wy int, tx, ty *int)
)

func (t *TreeView) AppendColumn(column *TreeViewColumn) int { return TreeViewAppendColumn(t, column) }
func (t *TreeView) CollapseAll()                            { TreeViewCollapseAll(t) }
func (t *TreeView) CollapseRow(path *TreePath) bool         { return TreeViewCollapseRow(t, path) }
func (t *TreeView) ColumnsAutosize()                        { TreeViewColumnsAutosize(t) }
func (t *TreeView) ConvertBinWindowToTreeCoords(bx, by int, tx, ty *int) {
	TreeViewConvertBinWindowToTreeCoords(t, bx, by, tx, ty)
}
func (t *TreeView) ConvertBinWindowToWidgetCoords(bx, by int, wx, wy *int) {
	TreeViewConvertBinWindowToWidgetCoords(t, bx, by, wx, wy)
}
func (t *TreeView) ConvertTreeToBinWindowCoords(tx, ty int, bx, by *int) {
	TreeViewConvertTreeToBinWindowCoords(t, tx, ty, bx, by)
}
func (t *TreeView) ConvertTreeToWidgetCoords(tx, ty int, wx, wy *int) {
	TreeViewConvertTreeToWidgetCoords(t, tx, ty, wx, wy)
}
func (t *TreeView) ConvertWidgetToBinWindowCoords(wx, wy int, bx, by *int) {
	TreeViewConvertWidgetToBinWindowCoords(t, wx, wy, bx, by)
}
func (t *TreeView) ConvertWidgetToTreeCoords(wx, wy int, tx, ty *int) {
	TreeViewConvertWidgetToTreeCoords(t, wx, wy, tx, ty)
}
func (t *TreeView) CreateRowDragIcon(path *TreePath) *D.Pixmap {
	return TreeViewCreateRowDragIcon(t, path)
}
func (t *TreeView) EnableModelDragDest(targets *TargetEntry, nTargets int, actions D.DragAction) {
	TreeViewEnableModelDragDest(t, targets, nTargets, actions)
}
func (t *TreeView) EnableModelDragSource(startButtonMask T.GdkModifierType, targets *TargetEntry, nTargets int, actions D.DragAction) {
	TreeViewEnableModelDragSource(t, startButtonMask, targets, nTargets, actions)
}
func (t *TreeView) ExpandAll() { TreeViewExpandAll(t) }
func (t *TreeView) ExpandRow(path *TreePath, openAll bool) bool {
	return TreeViewExpandRow(t, path, openAll)
}
func (t *TreeView) ExpandToPath(path *TreePath) { TreeViewExpandToPath(t, path) }
func (t *TreeView) GetBackgroundArea(path *TreePath, column *TreeViewColumn, rect *D.Rectangle) {
	TreeViewGetBackgroundArea(t, path, column, rect)
}
func (t *TreeView) GetBinWindow() *D.Window { return TreeViewGetBinWindow(t) }
func (t *TreeView) GetCellArea(path *TreePath, column *TreeViewColumn, rect *D.Rectangle) {
	TreeViewGetCellArea(t, path, column, rect)
}
func (t *TreeView) GetColumn(n int) *TreeViewColumn { return TreeViewGetColumn(t, n) }
func (t *TreeView) GetColumns() *T.GList            { return TreeViewGetColumns(t) }
func (t *TreeView) GetCursor(path **TreePath, focusColumn **TreeViewColumn) {
	TreeViewGetCursor(t, path, focusColumn)
}
func (t *TreeView) GetDestRowAtPos(dragX, dragY int, path **TreePath, pos *TreeViewDropPosition) bool {
	return TreeViewGetDestRowAtPos(t, dragX, dragY, path, pos)
}
func (t *TreeView) GetDragDestRow(path **TreePath, pos *TreeViewDropPosition) {
	TreeViewGetDragDestRow(t, path, pos)
}
func (t *TreeView) GetEnableSearch() bool              { return TreeViewGetEnableSearch(t) }
func (t *TreeView) GetEnableTreeLines() bool           { return TreeViewGetEnableTreeLines(t) }
func (t *TreeView) GetExpanderColumn() *TreeViewColumn { return TreeViewGetExpanderColumn(t) }
func (t *TreeView) GetFixedHeightMode() bool           { return TreeViewGetFixedHeightMode(t) }
func (t *TreeView) GetGridLines() TreeViewGridLines    { return TreeViewGetGridLines(t) }
func (t *TreeView) GetHadjustment() *Adjustment        { return TreeViewGetHadjustment(t) }
func (t *TreeView) GetHeadersClickable() bool          { return TreeViewGetHeadersClickable(t) }
func (t *TreeView) GetHeadersVisible() bool            { return TreeViewGetHeadersVisible(t) }
func (t *TreeView) GetHoverExpand() bool               { return TreeViewGetHoverExpand(t) }
func (t *TreeView) GetHoverSelection() bool            { return TreeViewGetHoverSelection(t) }
func (t *TreeView) GetLevelIndentation() int           { return TreeViewGetLevelIndentation(t) }
func (t *TreeView) GetModel() *TreeModel               { return TreeViewGetModel(t) }
func (t *TreeView) GetPathAtPos(x, y int, path **TreePath, column **TreeViewColumn, cellX, cellY *int) bool {
	return TreeViewGetPathAtPos(t, x, y, path, column, cellX, cellY)
}
func (t *TreeView) GetReorderable() bool { return TreeViewGetReorderable(t) }
func (t *TreeView) GetRowSeparatorFunc() TreeViewRowSeparatorFunc {
	return TreeViewGetRowSeparatorFunc(t)
}
func (t *TreeView) GetRubberBanding() bool { return TreeViewGetRubberBanding(t) }
func (t *TreeView) GetRulesHint() bool     { return TreeViewGetRulesHint(t) }
func (t *TreeView) GetSearchColumn() int   { return TreeViewGetSearchColumn(t) }
func (t *TreeView) GetSearchEntry() *Entry { return TreeViewGetSearchEntry(t) }
func (t *TreeView) GetSearchEqualFunc() TreeViewSearchEqualFunc {
	return TreeViewGetSearchEqualFunc(t)
}
func (t *TreeView) GetSearchPositionFunc() TreeViewSearchPositionFunc {
	return TreeViewGetSearchPositionFunc(t)
}
func (t *TreeView) GetSelection() *TreeSelection { return TreeViewGetSelection(t) }
func (t *TreeView) GetShowExpanders() bool       { return TreeViewGetShowExpanders(t) }
func (t *TreeView) GetTooltipColumn() int        { return TreeViewGetTooltipColumn(t) }
func (t *TreeView) GetTooltipContext(x, y *int, keyboardTip bool, model **TreeModel, path **TreePath, iter *TreeIter) bool {
	return TreeViewGetTooltipContext(t, x, y, keyboardTip, model, path, iter)
}
func (t *TreeView) GetVadjustment() *Adjustment { return TreeViewGetVadjustment(t) }
func (t *TreeView) GetVisibleRange(startPath **TreePath, endPath **TreePath) bool {
	return TreeViewGetVisibleRange(t, startPath, endPath)
}
func (t *TreeView) GetVisibleRect(visibleRect *D.Rectangle) {
	TreeViewGetVisibleRect(t, visibleRect)
}
func (t *TreeView) InsertColumn(column *TreeViewColumn, position int) int {
	return TreeViewInsertColumn(t, column, position)
}
func (t *TreeView) InsertColumnWithAttributes(position int, title string, cell *CellRenderer, v ...VArg) int {
	return TreeViewInsertColumnWithAttributes(t, position, title, cell, v)
}
func (t *TreeView) InsertColumnWithDataFunc(position int, title string, cell *CellRenderer, f TreeCellDataFunc, data T.Gpointer, dnotify T.GDestroyNotify) int {
	return TreeViewInsertColumnWithDataFunc(t, position, title, cell, f, data, dnotify)
}
func (t *TreeView) IsRubberBandingActive() bool { return TreeViewIsRubberBandingActive(t) }
func (t *TreeView) MapExpandedRows(f TreeViewMappingFunc, data T.Gpointer) {
	TreeViewMapExpandedRows(t, f, data)
}
func (t *TreeView) MoveColumnAfter(column *TreeViewColumn, baseColumn *TreeViewColumn) {
	TreeViewMoveColumnAfter(t, column, baseColumn)
}
func (t *TreeView) RemoveColumn(column *TreeViewColumn) int { return TreeViewRemoveColumn(t, column) }
func (t *TreeView) RowActivated(path *TreePath, column *TreeViewColumn) {
	TreeViewRowActivated(t, path, column)
}
func (t *TreeView) RowExpanded(path *TreePath) bool { return TreeViewRowExpanded(t, path) }
func (t *TreeView) ScrollToCell(path *TreePath, column *TreeViewColumn, useAlign bool, rowAlign, colAlign float32) {
	TreeViewScrollToCell(t, path, column, useAlign, rowAlign, colAlign)
}
func (t *TreeView) ScrollToPoint(treeX, treeY int) { TreeViewScrollToPoint(t, treeX, treeY) }
func (t *TreeView) SetColumnDragFunction(f TreeViewColumnDropFunc, userData T.Gpointer, destroy T.GDestroyNotify) {
	TreeViewSetColumnDragFunction(t, f, userData, destroy)
}
func (t *TreeView) SetCursor(path *TreePath, focusColumn *TreeViewColumn, startEditing bool) {
	TreeViewSetCursor(t, path, focusColumn, startEditing)
}
func (t *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing bool) {
	TreeViewSetCursorOnCell(t, path, focusColumn, focusCell, startEditing)
}
func (t *TreeView) SetDestroyCountFunc(f TreeDestroyCountFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	TreeViewSetDestroyCountFunc(t, f, data, destroy)
}
func (t *TreeView) SetDragDestRow(path *TreePath, pos TreeViewDropPosition) {
	TreeViewSetDragDestRow(t, path, pos)
}
func (t *TreeView) SetEnableSearch(enableSearch bool)        { TreeViewSetEnableSearch(t, enableSearch) }
func (t *TreeView) SetEnableTreeLines(enabled bool)          { TreeViewSetEnableTreeLines(t, enabled) }
func (t *TreeView) SetExpanderColumn(column *TreeViewColumn) { TreeViewSetExpanderColumn(t, column) }
func (t *TreeView) SetFixedHeightMode(enable bool)           { TreeViewSetFixedHeightMode(t, enable) }
func (t *TreeView) SetGridLines(gridLines TreeViewGridLines) { TreeViewSetGridLines(t, gridLines) }
func (t *TreeView) SetHadjustment(adjustment *Adjustment)    { TreeViewSetHadjustment(t, adjustment) }
func (t *TreeView) SetHeadersClickable(setting bool)         { TreeViewSetHeadersClickable(t, setting) }
func (t *TreeView) SetHeadersVisible(headersVisible bool) {
	TreeViewSetHeadersVisible(t, headersVisible)
}
func (t *TreeView) SetHoverExpand(expand bool)          { TreeViewSetHoverExpand(t, expand) }
func (t *TreeView) SetHoverSelection(hover bool)        { TreeViewSetHoverSelection(t, hover) }
func (t *TreeView) SetLevelIndentation(indentation int) { TreeViewSetLevelIndentation(t, indentation) }
func (t *TreeView) SetModel(model *TreeModel)           { TreeViewSetModel(t, model) }
func (t *TreeView) SetReorderable(reorderable bool)     { TreeViewSetReorderable(t, reorderable) }
func (t *TreeView) SetRowSeparatorFunc(f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	TreeViewSetRowSeparatorFunc(t, f, data, destroy)
}
func (t *TreeView) SetRubberBanding(enable bool) { TreeViewSetRubberBanding(t, enable) }
func (t *TreeView) SetRulesHint(setting bool)    { TreeViewSetRulesHint(t, setting) }
func (t *TreeView) SetSearchColumn(column int)   { TreeViewSetSearchColumn(t, column) }
func (t *TreeView) SetSearchEntry(entry *Entry)  { TreeViewSetSearchEntry(t, entry) }
func (t *TreeView) SetSearchEqualFunc(searchEqualFunc TreeViewSearchEqualFunc, searchUserData T.Gpointer, searchDestroy T.GDestroyNotify) {
	TreeViewSetSearchEqualFunc(t, searchEqualFunc, searchUserData, searchDestroy)
}
func (t *TreeView) SetSearchPositionFunc(f TreeViewSearchPositionFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	TreeViewSetSearchPositionFunc(t, f, data, destroy)
}
func (t *TreeView) SetShowExpanders(enabled bool) { TreeViewSetShowExpanders(t, enabled) }
func (t *TreeView) SetTooltipCell(tooltip *Tooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer) {
	TreeViewSetTooltipCell(t, tooltip, path, column, cell)
}
func (t *TreeView) SetTooltipColumn(column int) { TreeViewSetTooltipColumn(t, column) }
func (t *TreeView) SetTooltipRow(tooltip *Tooltip, path *TreePath) {
	TreeViewSetTooltipRow(t, tooltip, path)
}
func (t *TreeView) SetVadjustment(adjustment *Adjustment) { TreeViewSetVadjustment(t, adjustment) }
func (t *TreeView) TreeToWidgetCoords(tx, ty int, wx, wy *int) {
	TreeViewTreeToWidgetCoords(t, tx, ty, wx, wy)
}
func (t *TreeView) UnsetRowsDragDest()   { TreeViewUnsetRowsDragDest(t) }
func (t *TreeView) UnsetRowsDragSource() { TreeViewUnsetRowsDragSource(t) }
func (t *TreeView) WidgetToTreeCoords(wx, wy int, tx, ty *int) {
	TreeViewWidgetToTreeCoords(t, wx, wy, tx, ty)
}

type TreeViewColumn struct {
	Parent                  Object
	TreeView                *Widget
	Button                  *Widget
	Child                   *Widget
	Arrow                   *Widget
	Alignment               *Widget
	Window                  *D.Window
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
	SortOrder               SortType
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
	TreeViewColumnGetType           func() O.Type
	TreeViewColumnNew               func() *TreeViewColumn
	TreeViewColumnNewWithAttributes func(title string, cell *CellRenderer, v ...VArg) *TreeViewColumn

	TreeViewColumnSizingGetType func() O.Type

	TreeViewColumnAddAttribute     func(t *TreeViewColumn, cellRenderer *CellRenderer, attribute string, column int)
	TreeViewColumnCellGetPosition  func(t *TreeViewColumn, cellRenderer *CellRenderer, startPos, width *int) bool
	TreeViewColumnCellGetSize      func(t *TreeViewColumn, cellArea *D.Rectangle, xOffset, yOffset, width, height *int)
	TreeViewColumnCellIsVisible    func(t *TreeViewColumn) bool
	TreeViewColumnCellSetCellData  func(t *TreeViewColumn, treeModel *TreeModel, iter *TreeIter, isExpander, isExpanded bool)
	TreeViewColumnClear            func(t *TreeViewColumn)
	TreeViewColumnClearAttributes  func(t *TreeViewColumn, cellRenderer *CellRenderer)
	TreeViewColumnClicked          func(t *TreeViewColumn)
	TreeViewColumnFocusCell        func(t *TreeViewColumn, cell *CellRenderer)
	TreeViewColumnGetAlignment     func(t *TreeViewColumn) float32
	TreeViewColumnGetCellRenderers func(t *TreeViewColumn) *T.GList
	TreeViewColumnGetClickable     func(t *TreeViewColumn) bool
	TreeViewColumnGetExpand        func(t *TreeViewColumn) bool
	TreeViewColumnGetFixedWidth    func(t *TreeViewColumn) int
	TreeViewColumnGetMaxWidth      func(t *TreeViewColumn) int
	TreeViewColumnGetMinWidth      func(t *TreeViewColumn) int
	TreeViewColumnGetReorderable   func(t *TreeViewColumn) bool
	TreeViewColumnGetResizable     func(t *TreeViewColumn) bool
	TreeViewColumnGetSizing        func(t *TreeViewColumn) TreeViewColumnSizing
	TreeViewColumnGetSortColumnId  func(t *TreeViewColumn) int
	TreeViewColumnGetSortIndicator func(t *TreeViewColumn) bool
	TreeViewColumnGetSortOrder     func(t *TreeViewColumn) SortType
	TreeViewColumnGetSpacing       func(t *TreeViewColumn) int
	TreeViewColumnGetTitle         func(t *TreeViewColumn) string
	TreeViewColumnGetTreeView      func(t *TreeViewColumn) *Widget
	TreeViewColumnGetVisible       func(t *TreeViewColumn) bool
	TreeViewColumnGetWidget        func(t *TreeViewColumn) *Widget
	TreeViewColumnGetWidth         func(t *TreeViewColumn) int
	TreeViewColumnPackEnd          func(t *TreeViewColumn, cell *CellRenderer, expand bool)
	TreeViewColumnPackStart        func(t *TreeViewColumn, cell *CellRenderer, expand bool)
	TreeViewColumnQueueResize      func(t *TreeViewColumn)
	TreeViewColumnSetAlignment     func(t *TreeViewColumn, xalign float32)
	TreeViewColumnSetAttributes    func(t *TreeViewColumn, cellRenderer *CellRenderer, v ...VArg)
	TreeViewColumnSetCellDataFunc  func(t *TreeViewColumn, cellRenderer *CellRenderer, f TreeCellDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
	TreeViewColumnSetClickable     func(t *TreeViewColumn, clickable bool)
	TreeViewColumnSetExpand        func(t *TreeViewColumn, expand bool)
	TreeViewColumnSetFixedWidth    func(t *TreeViewColumn, fixedWidth int)
	TreeViewColumnSetMaxWidth      func(t *TreeViewColumn, maxWidth int)
	TreeViewColumnSetMinWidth      func(t *TreeViewColumn, minWidth int)
	TreeViewColumnSetReorderable   func(t *TreeViewColumn, reorderable bool)
	TreeViewColumnSetResizable     func(t *TreeViewColumn, resizable bool)
	TreeViewColumnSetSizing        func(t *TreeViewColumn, typ TreeViewColumnSizing)
	TreeViewColumnSetSortColumnId  func(t *TreeViewColumn, sortColumnId int)
	TreeViewColumnSetSortIndicator func(t *TreeViewColumn, setting bool)
	TreeViewColumnSetSortOrder     func(t *TreeViewColumn, order SortType)
	TreeViewColumnSetSpacing       func(t *TreeViewColumn, spacing int)
	TreeViewColumnSetTitle         func(t *TreeViewColumn, title string)
	TreeViewColumnSetVisible       func(t *TreeViewColumn, visible bool)
	TreeViewColumnSetWidget        func(t *TreeViewColumn, widget *Widget)
)

func (t *TreeViewColumn) AddAttribute(cellRenderer *CellRenderer, attribute string, column int) {
	TreeViewColumnAddAttribute(t, cellRenderer, attribute, column)
}
func (t *TreeViewColumn) CellGetPosition(cellRenderer *CellRenderer, startPos *int, width *int) bool {
	return TreeViewColumnCellGetPosition(t, cellRenderer, startPos, width)
}
func (t *TreeViewColumn) CellGetSize(cellArea *D.Rectangle, xOffset, yOffset, width, height *int) {
	TreeViewColumnCellGetSize(t, cellArea, xOffset, yOffset, width, height)
}
func (t *TreeViewColumn) CellIsVisible() bool { return TreeViewColumnCellIsVisible(t) }
func (t *TreeViewColumn) CellSetCellData(treeModel *TreeModel, iter *TreeIter, isExpander, isExpanded bool) {
	TreeViewColumnCellSetCellData(t, treeModel, iter, isExpander, isExpanded)
}
func (t *TreeViewColumn) Clear() { TreeViewColumnClear(t) }
func (t *TreeViewColumn) ClearAttributes(cellRenderer *CellRenderer) {
	TreeViewColumnClearAttributes(t, cellRenderer)
}
func (t *TreeViewColumn) Clicked()                        { TreeViewColumnClicked(t) }
func (t *TreeViewColumn) FocusCell(cell *CellRenderer)    { TreeViewColumnFocusCell(t, cell) }
func (t *TreeViewColumn) GetAlignment() float32           { return TreeViewColumnGetAlignment(t) }
func (t *TreeViewColumn) GetCellRenderers() *T.GList      { return TreeViewColumnGetCellRenderers(t) }
func (t *TreeViewColumn) GetClickable() bool              { return TreeViewColumnGetClickable(t) }
func (t *TreeViewColumn) GetExpand() bool                 { return TreeViewColumnGetExpand(t) }
func (t *TreeViewColumn) GetFixedWidth() int              { return TreeViewColumnGetFixedWidth(t) }
func (t *TreeViewColumn) GetMaxWidth() int                { return TreeViewColumnGetMaxWidth(t) }
func (t *TreeViewColumn) GetMinWidth() int                { return TreeViewColumnGetMinWidth(t) }
func (t *TreeViewColumn) GetReorderable() bool            { return TreeViewColumnGetReorderable(t) }
func (t *TreeViewColumn) GetResizable() bool              { return TreeViewColumnGetResizable(t) }
func (t *TreeViewColumn) GetSizing() TreeViewColumnSizing { return TreeViewColumnGetSizing(t) }
func (t *TreeViewColumn) GetSortColumnId() int            { return TreeViewColumnGetSortColumnId(t) }
func (t *TreeViewColumn) GetSortIndicator() bool          { return TreeViewColumnGetSortIndicator(t) }
func (t *TreeViewColumn) GetSortOrder() SortType          { return TreeViewColumnGetSortOrder(t) }
func (t *TreeViewColumn) GetSpacing() int                 { return TreeViewColumnGetSpacing(t) }
func (t *TreeViewColumn) GetTitle() string                { return TreeViewColumnGetTitle(t) }
func (t *TreeViewColumn) GetTreeView() *Widget            { return TreeViewColumnGetTreeView(t) }
func (t *TreeViewColumn) GetVisible() bool                { return TreeViewColumnGetVisible(t) }
func (t *TreeViewColumn) GetWidget() *Widget              { return TreeViewColumnGetWidget(t) }
func (t *TreeViewColumn) GetWidth() int                   { return TreeViewColumnGetWidth(t) }
func (t *TreeViewColumn) PackStart(cell *CellRenderer, expand bool) {
	TreeViewColumnPackEnd(t, cell, expand)
}
func (t *TreeViewColumn) QueueResize()                { TreeViewColumnQueueResize(t) }
func (t *TreeViewColumn) SetAlignment(xalign float32) { TreeViewColumnSetAlignment(t, xalign) }
func (t *TreeViewColumn) SetAttributes(cellRenderer *CellRenderer, v ...VArg) {
	TreeViewColumnSetAttributes(t, cellRenderer, v)
}
func (t *TreeViewColumn) SetCellDataFunc(cellRenderer *CellRenderer, f TreeCellDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	TreeViewColumnSetCellDataFunc(t, cellRenderer, f, funcData, destroy)
}
func (t *TreeViewColumn) SetClickable(clickable bool)  { TreeViewColumnSetClickable(t, clickable) }
func (t *TreeViewColumn) SetExpand(expand bool)        { TreeViewColumnSetExpand(t, expand) }
func (t *TreeViewColumn) SetFixedWidth(fixedWidth int) { TreeViewColumnSetFixedWidth(t, fixedWidth) }
func (t *TreeViewColumn) SetMaxWidth(maxWidth int)     { TreeViewColumnSetMaxWidth(t, maxWidth) }
func (t *TreeViewColumn) SetMinWidth(minWidth int)     { TreeViewColumnSetMinWidth(t, minWidth) }
func (t *TreeViewColumn) SetReorderable(reorderable bool) {
	TreeViewColumnSetReorderable(t, reorderable)
}
func (t *TreeViewColumn) SetResizable(resizable bool)        { TreeViewColumnSetResizable(t, resizable) }
func (t *TreeViewColumn) SetSizing(typ TreeViewColumnSizing) { TreeViewColumnSetSizing(t, typ) }
func (t *TreeViewColumn) SetSortColumnId(sortColumnId int) {
	TreeViewColumnSetSortColumnId(t, sortColumnId)
}
func (t *TreeViewColumn) SetSortIndicator(setting bool) {
	TreeViewColumnSetSortIndicator(t, setting)
}
func (t *TreeViewColumn) SetSortOrder(order SortType) { TreeViewColumnSetSortOrder(t, order) }
func (t *TreeViewColumn) SetSpacing(spacing int)      { TreeViewColumnSetSpacing(t, spacing) }
func (t *TreeViewColumn) SetTitle(title string)       { TreeViewColumnSetTitle(t, title) }
func (t *TreeViewColumn) SetVisible(visible bool)     { TreeViewColumnSetVisible(t, visible) }
func (t *TreeViewColumn) SetWidget(widget *Widget)    { TreeViewColumnSetWidget(t, widget) }

type (
	Type O.Type

	TypeInfo struct {
		TypeName          *T.Gchar
		ObjectSize        uint
		ClassSize         uint
		ClassInitFunc     ClassInitFunc
		ObjectInitFunc    ObjectInitFunc
		_, _              T.Gpointer
		BaseClassInitFunc ClassInitFunc
	}
)

var (
	TypeInit func(debugFlags O.TypeDebugFlags)

	TypeClass  func(t Type) T.Gpointer
	TypeNew    func(t Type) T.Gpointer
	TypeUnique func(t Type, gtkinfo *TypeInfo) Type

	TypeEnumFindValue  func(enumType Type, valueName string) *EnumValue
	TypeEnumGetValues  func(enumType Type) *EnumValue
	TypeFlagsFindValue func(flagsType Type, valueName string) *FlagValue
	TypeFlagsGetValues func(flagsType Type) *FlagValue
)

func (t Type) Class() T.Gpointer             { return TypeClass(t) }
func (t Type) New() T.Gpointer               { return TypeNew(t) }
func (t Type) Unique(gtkinfo *TypeInfo) Type { return TypeUnique(t, gtkinfo) }

func (t Type) EnumFindValue(valueName string) *EnumValue  { return TypeEnumFindValue(t, valueName) }
func (t Type) EnumGetValues() *EnumValue                  { return TypeEnumGetValues(t) }
func (t Type) FlagsFindValue(valueName string) *FlagValue { return TypeFlagsFindValue(t, valueName) }
func (t Type) FlagsGetValues() *FlagValue                 { return TypeFlagsGetValues(t) }
