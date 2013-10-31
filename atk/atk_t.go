// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Table struct{}

var (
	TableGetType func() O.Type

	TableAddColumnSelection    func(t *Table, column int) bool
	TableAddRowSelection       func(t *Table, row int) bool
	TableGetCaption            func(t *Table) *Object
	TableGetColumnAtIndex      func(t *Table, index int) int
	TableGetColumnDescription  func(t *Table, column int) string
	TableGetColumnExtentAt     func(t *Table, row, column int) int
	TableGetColumnHeader       func(t *Table, column int) *Object
	TableGetIndexAt            func(t *Table, row, column int) int
	TableGetNColumns           func(t *Table) int
	TableGetNRows              func(t *Table) int
	TableGetRowAtIndex         func(t *Table, index int) int
	TableGetRowDescription     func(t *Table, row int) string
	TableGetRowExtentAt        func(t *Table, row, column int) int
	TableGetRowHeader          func(t *Table, row int) *Object
	TableGetSelectedColumns    func(t *Table, selected **int) int
	TableGetSelectedRows       func(t *Table, selected **int) int
	TableGetSummary            func(t *Table) *Object
	TableIsColumnSelected      func(t *Table, column int) bool
	TableIsRowSelected         func(t *Table, row int) bool
	TableIsSelected            func(t *Table, row, column int) bool
	TableRefAt                 func(t *Table, row, column int) *Object
	TableRemoveColumnSelection func(t *Table, column int) bool
	TableRemoveRowSelection    func(t *Table, row int) bool
	TableSetCaption            func(t *Table, caption *Object)
	TableSetColumnDescription  func(t *Table, column int, description string)
	TableSetColumnHeader       func(t *Table, column int, header *Object)
	TableSetRowDescription     func(t *Table, row int, description string)
	TableSetRowHeader          func(t *Table, row int, header *Object)
	TableSetSummary            func(t *Table, accessible *Object)
)

func (t *Table) AddColumnSelection(column int) bool     { return TableAddColumnSelection(t, column) }
func (t *Table) AddRowSelection(row int) bool           { return TableAddRowSelection(t, row) }
func (t *Table) GetCaption() *Object                    { return TableGetCaption(t) }
func (t *Table) GetColumnAtIndex(index int) int         { return TableGetColumnAtIndex(t, index) }
func (t *Table) GetColumnDescription(column int) string { return TableGetColumnDescription(t, column) }
func (t *Table) GetColumnExtentAt(row, column int) int  { return TableGetColumnExtentAt(t, row, column) }
func (t *Table) GetColumnHeader(column int) *Object     { return TableGetColumnHeader(t, column) }
func (t *Table) GetIndexAt(row, column int) int         { return TableGetIndexAt(t, row, column) }
func (t *Table) GetNColumns() int                       { return TableGetNColumns(t) }
func (t *Table) GetNRows() int                          { return TableGetNRows(t) }
func (t *Table) GetRowAtIndex(index int) int            { return TableGetRowAtIndex(t, index) }
func (t *Table) GetRowDescription(row int) string       { return TableGetRowDescription(t, row) }
func (t *Table) GetRowExtentAt(row, column int) int     { return TableGetRowExtentAt(t, row, column) }
func (t *Table) GetRowHeader(row int) *Object           { return TableGetRowHeader(t, row) }
func (t *Table) GetSelectedColumns(selected **int) int  { return TableGetSelectedColumns(t, selected) }
func (t *Table) GetSelectedRows(selected **int) int     { return TableGetSelectedRows(t, selected) }
func (t *Table) GetSummary() *Object                    { return TableGetSummary(t) }
func (t *Table) IsColumnSelected(column int) bool       { return TableIsColumnSelected(t, column) }
func (t *Table) IsRowSelected(row int) bool             { return TableIsRowSelected(t, row) }
func (t *Table) IsSelected(row, column int) bool        { return TableIsSelected(t, row, column) }
func (t *Table) RefAt(row, column int) *Object          { return TableRefAt(t, row, column) }
func (t *Table) RemoveColumnSelection(column int) bool {
	return TableRemoveColumnSelection(t, column)
}
func (t *Table) RemoveRowSelection(row int) bool { return TableRemoveRowSelection(t, row) }
func (t *Table) SetCaption(caption *Object)      { TableSetCaption(t, caption) }
func (t *Table) SetColumnDescription(column int, description string) {
	TableSetColumnDescription(t, column, description)
}
func (t *Table) SetColumnHeader(column int, header *Object) { TableSetColumnHeader(t, column, header) }
func (t *Table) SetRowDescription(row int, description string) {
	TableSetRowDescription(t, row, description)
}
func (t *Table) SetRowHeader(row int, header *Object) { TableSetRowHeader(t, row, header) }
func (t *Table) SetSummary(accessible *Object)        { TableSetSummary(t, accessible) }

type Text struct{}

var (
	TextGetType func() O.Type

	TextAttributeRegister func(name string) TextAttribute
	TextFreeRanges        func(ranges **TextRange)

	TextAddSelection         func(t *Text, startOffset, endOffset int) bool
	TextGetBoundedRanges     func(t *Text, rect *TextRectangle, coordType CoordType, xClipType, yClipType TextClipType) **TextRange
	TextGetCaretOffset       func(t *Text) int
	TextGetCharacterAtOffset func(t *Text, offset int) T.Gunichar
	TextGetCharacterCount    func(t *Text) int
	TextGetCharacterExtents  func(t *Text, offset int, x, y, width, height *int, coords CoordType)
	TextGetDefaultAttributes func(t *Text) *AttributeSet
	TextGetNSelections       func(t *Text) int
	TextGetOffsetAtPoint     func(t *Text, x, y int, coords CoordType) int
	TextGetRangeExtents      func(t *Text, startOffset, endOffset int, coordType CoordType, rect *TextRectangle)
	TextGetRunAttributes     func(t *Text, offset int, startOffset, endOffset *int) *AttributeSet
	TextGetSelection         func(t *Text, selectionNum int, startOffset, endOffset *int) string
	TextGetText              func(t *Text, startOffset, endOffset int) string
	TextGetTextAfterOffset   func(t *Text, offset int, boundaryType TextBoundary, startOffset, endOffset *int) string
	TextGetTextAtOffset      func(t *Text, offset int, boundaryType TextBoundary, startOffset, endOffset *int) string
	TextGetTextBeforeOffset  func(t *Text, offset int, boundaryType TextBoundary, startOffset, endOffset *int) string
	TextRemoveSelection      func(t *Text, selectionNum int) bool
	TextSetCaretOffset       func(t *Text, offset int) bool
	TextSetSelection         func(t *Text, selectionNum, startOffset, endOffset int) bool
)

func (t *Text) AddSelection(startOffset, endOffset int) bool {
	return TextAddSelection(t, startOffset, endOffset)
}
func (t *Text) GetBoundedRanges(rect *TextRectangle, coordType CoordType, xClipType, yClipType TextClipType) **TextRange {
	return TextGetBoundedRanges(t, rect, coordType, xClipType, yClipType)
}
func (t *Text) GetCaretOffset() int                        { return TextGetCaretOffset(t) }
func (t *Text) GetCharacterAtOffset(offset int) T.Gunichar { return TextGetCharacterAtOffset(t, offset) }
func (t *Text) GetCharacterCount() int                     { return TextGetCharacterCount(t) }
func (t *Text) GetCharacterExtents(offset int, x, y, width, height *int, coords CoordType) {
	TextGetCharacterExtents(t, offset, x, y, width, height, coords)
}
func (t *Text) GetDefaultAttributes() *AttributeSet { return TextGetDefaultAttributes(t) }
func (t *Text) GetNSelections() int                 { return TextGetNSelections(t) }
func (t *Text) GetOffsetAtPoint(x, y int, coords CoordType) int {
	return TextGetOffsetAtPoint(t, x, y, coords)
}
func (t *Text) GetRangeExtents(startOffset, endOffset int, coordType CoordType, rect *TextRectangle) {
	TextGetRangeExtents(t, startOffset, endOffset, coordType, rect)
}
func (t *Text) GetRunAttributes(offset int, startOffset, endOffset *int) *AttributeSet {
	return TextGetRunAttributes(t, offset, startOffset, endOffset)
}
func (t *Text) GetSelection(selectionNum int, startOffset, endOffset *int) string {
	return TextGetSelection(t, selectionNum, startOffset, endOffset)
}
func (t *Text) GetText(startOffset, endOffset int) string {
	return TextGetText(t, startOffset, endOffset)
}
func (t *Text) GetTextAfterOffset(offset int, boundaryType TextBoundary, startOffset, endOffset *int) string {
	return TextGetTextAfterOffset(t, offset, boundaryType, startOffset, endOffset)
}
func (t *Text) GetTextAtOffset(offset int, boundaryType TextBoundary, startOffset, endOffset *int) string {
	return TextGetTextAtOffset(t, offset, boundaryType, startOffset, endOffset)
}
func (t *Text) GetTextBeforeOffset(offset int, boundaryType TextBoundary, startOffset, endOffset *int) string {
	return TextGetTextBeforeOffset(t, offset, boundaryType, startOffset, endOffset)
}
func (t *Text) RemoveSelection(selectionNum int) bool {
	return TextRemoveSelection(t, selectionNum)
}
func (t *Text) SetCaretOffset(offset int) bool { return TextSetCaretOffset(t, offset) }
func (t *Text) SetSelection(selectionNum, startOffset, endOffset int) bool {
	return TextSetSelection(t, selectionNum, startOffset, endOffset)
}

type TextAttribute Enum

const (
	TEXT_ATTR_INVALID TextAttribute = iota
	TEXT_ATTR_LEFT_MARGIN
	TEXT_ATTR_RIGHT_MARGIN
	TEXT_ATTR_INDENT
	TEXT_ATTR_INVISIBLE
	TEXT_ATTR_EDITABLE
	TEXT_ATTR_PIXELS_ABOVE_LINES
	TEXT_ATTR_PIXELS_BELOW_LINES
	TEXT_ATTR_PIXELS_INSIDE_WRAP
	TEXT_ATTR_BG_FULL_HEIGHT
	TEXT_ATTR_RISE
	TEXT_ATTR_UNDERLINE
	TEXT_ATTR_STRIKETHROUGH
	TEXT_ATTR_SIZE
	TEXT_ATTR_SCALE
	TEXT_ATTR_WEIGHT
	TEXT_ATTR_LANGUAGE
	TEXT_ATTR_FAMILY_NAME
	TEXT_ATTR_BG_COLOR
	TEXT_ATTR_FG_COLOR
	TEXT_ATTR_BG_STIPPLE
	TEXT_ATTR_FG_STIPPLE
	TEXT_ATTR_WRAP_MODE
	TEXT_ATTR_DIRECTION
	TEXT_ATTR_JUSTIFICATION
	TEXT_ATTR_STRETCH
	TEXT_ATTR_VARIANT
	TEXT_ATTR_STYLE
	TEXT_ATTR_LAST_DEFINED
)

var (
	TextAttributeGetType func() O.Type
	TextAttributeForName func(name string) TextAttribute

	TextAttributeGetName  func(t TextAttribute) string
	TextAttributeGetValue func(t TextAttribute, index int) string
)

func (t TextAttribute) GetName() string           { return TextAttributeGetName(t) }
func (t TextAttribute) GetValue(index int) string { return TextAttributeGetValue(t, index) }

type TextBoundary Enum

const (
	TEXT_BOUNDARY_CHAR TextBoundary = iota
	TEXT_BOUNDARY_WORD_START
	TEXT_BOUNDARY_WORD_END
	TEXT_BOUNDARY_SENTENCE_START
	TEXT_BOUNDARY_SENTENCE_END
	TEXT_BOUNDARY_LINE_START
	TEXT_BOUNDARY_LINE_END
)

var TextBoundaryGetType func() O.Type

type TextClipType Enum

const (
	TEXT_CLIP_NONE TextClipType = iota
	TEXT_CLIP_MIN
	TEXT_CLIP_MAX
	TEXT_CLIP_BOTH
)

var TextClipTypeGetType func() O.Type

type TextRange struct {
	Bounds      TextRectangle
	StartOffset int
	EndOffset   int
	Content     *T.Gchar
}

type TextRectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}
