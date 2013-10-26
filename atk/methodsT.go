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

	tableAddColumnSelection    func(t *Table, column int) T.Gboolean
	tableAddRowSelection       func(t *Table, row int) T.Gboolean
	tableGetCaption            func(t *Table) *Object
	tableGetColumnAtIndex      func(t *Table, index int) int
	tableGetColumnDescription  func(t *Table, column int) string
	tableGetColumnExtentAt     func(t *Table, row, column int) int
	tableGetColumnHeader       func(t *Table, column int) *Object
	tableGetIndexAt            func(t *Table, row, column int) int
	tableGetNColumns           func(t *Table) int
	tableGetNRows              func(t *Table) int
	tableGetRowAtIndex         func(t *Table, index int) int
	tableGetRowDescription     func(t *Table, row int) string
	tableGetRowExtentAt        func(t *Table, row, column int) int
	tableGetRowHeader          func(t *Table, row int) *Object
	tableGetSelectedColumns    func(t *Table, selected **int) int
	tableGetSelectedRows       func(t *Table, selected **int) int
	tableGetSummary            func(t *Table) *Object
	tableIsColumnSelected      func(t *Table, column int) T.Gboolean
	tableIsRowSelected         func(t *Table, row int) T.Gboolean
	tableIsSelected            func(t *Table, row, column int) T.Gboolean
	tableRefAt                 func(t *Table, row, column int) *Object
	tableRemoveColumnSelection func(t *Table, column int) T.Gboolean
	tableRemoveRowSelection    func(t *Table, row int) T.Gboolean
	tableSetCaption            func(t *Table, caption *Object)
	tableSetColumnDescription  func(t *Table, column int, description string)
	tableSetColumnHeader       func(t *Table, column int, header *Object)
	tableSetRowDescription     func(t *Table, row int, description string)
	tableSetRowHeader          func(t *Table, row int, header *Object)
	tableSetSummary            func(t *Table, accessible *Object)
)

func (t *Table) AddColumnSelection(column int) T.Gboolean { return tableAddColumnSelection(t, column) }
func (t *Table) AddRowSelection(row int) T.Gboolean       { return tableAddRowSelection(t, row) }
func (t *Table) GetCaption() *Object                      { return tableGetCaption(t) }
func (t *Table) GetColumnAtIndex(index int) int           { return tableGetColumnAtIndex(t, index) }
func (t *Table) GetColumnDescription(column int) string   { return tableGetColumnDescription(t, column) }
func (t *Table) GetColumnExtentAt(row, column int) int    { return tableGetColumnExtentAt(t, row, column) }
func (t *Table) GetColumnHeader(column int) *Object       { return tableGetColumnHeader(t, column) }
func (t *Table) GetIndexAt(row, column int) int           { return tableGetIndexAt(t, row, column) }
func (t *Table) GetNColumns() int                         { return tableGetNColumns(t) }
func (t *Table) GetNRows() int                            { return tableGetNRows(t) }
func (t *Table) GetRowAtIndex(index int) int              { return tableGetRowAtIndex(t, index) }
func (t *Table) GetRowDescription(row int) string         { return tableGetRowDescription(t, row) }
func (t *Table) GetRowExtentAt(row, column int) int       { return tableGetRowExtentAt(t, row, column) }
func (t *Table) GetRowHeader(row int) *Object             { return tableGetRowHeader(t, row) }
func (t *Table) GetSelectedColumns(selected **int) int    { return tableGetSelectedColumns(t, selected) }
func (t *Table) GetSelectedRows(selected **int) int       { return tableGetSelectedRows(t, selected) }
func (t *Table) GetSummary() *Object                      { return tableGetSummary(t) }
func (t *Table) IsColumnSelected(column int) T.Gboolean   { return tableIsColumnSelected(t, column) }
func (t *Table) IsRowSelected(row int) T.Gboolean         { return tableIsRowSelected(t, row) }
func (t *Table) IsSelected(row, column int) T.Gboolean    { return tableIsSelected(t, row, column) }
func (t *Table) RefAt(row, column int) *Object            { return tableRefAt(t, row, column) }
func (t *Table) RemoveColumnSelection(column int) T.Gboolean {
	return tableRemoveColumnSelection(t, column)
}
func (t *Table) RemoveRowSelection(row int) T.Gboolean { return tableRemoveRowSelection(t, row) }
func (t *Table) SetCaption(caption *Object)            { tableSetCaption(t, caption) }
func (t *Table) SetColumnDescription(column int, description string) {
	tableSetColumnDescription(t, column, description)
}
func (t *Table) SetColumnHeader(column int, header *Object) { tableSetColumnHeader(t, column, header) }
func (t *Table) SetRowDescription(row int, description string) {
	tableSetRowDescription(t, row, description)
}
func (t *Table) SetRowHeader(row int, header *Object) { tableSetRowHeader(t, row, header) }
func (t *Table) SetSummary(accessible *Object)        { tableSetSummary(t, accessible) }

type Text struct{}

var (
	TextGetType func() O.Type

	TextAttributeRegister func(name string) TextAttribute
	TextFreeRanges        func(ranges **TextRange)

	textAddSelection         func(t *Text, startOffset, endOffset int) T.Gboolean
	textGetBoundedRanges     func(t *Text, rect *TextRectangle, coordType CoordType, xClipType, yClipType TextClipType) **TextRange
	textGetCaretOffset       func(t *Text) int
	textGetCharacterAtOffset func(t *Text, offset int) T.Gunichar
	textGetCharacterCount    func(t *Text) int
	textGetCharacterExtents  func(t *Text, offset int, x, y, width, height *int, coords CoordType)
	textGetDefaultAttributes func(t *Text) *AttributeSet
	textGetNSelections       func(t *Text) int
	textGetOffsetAtPoint     func(t *Text, x, y int, coords CoordType) int
	textGetRangeExtents      func(t *Text, startOffset, endOffset int, coordType CoordType, rect *TextRectangle)
	textGetRunAttributes     func(t *Text, offset int, startOffset, endOffset *int) *AttributeSet
	textGetSelection         func(t *Text, selectionNum int, startOffset, endOffset *int) string
	textGetText              func(t *Text, startOffset, endOffset int) string
	textGetTextAfterOffset   func(t *Text, offset int, boundaryType TextBoundary, startOffset, endOffset *int) string
	textGetTextAtOffset      func(t *Text, offset int, boundaryType TextBoundary, startOffset, endOffset *int) string
	textGetTextBeforeOffset  func(t *Text, offset int, boundaryType TextBoundary, startOffset, endOffset *int) string
	textRemoveSelection      func(t *Text, selectionNum int) T.Gboolean
	textSetCaretOffset       func(t *Text, offset int) T.Gboolean
	textSetSelection         func(t *Text, selectionNum, startOffset, endOffset int) T.Gboolean
)

func (t *Text) AddSelection(startOffset, endOffset int) T.Gboolean {
	return textAddSelection(t, startOffset, endOffset)
}
func (t *Text) GetBoundedRanges(rect *TextRectangle, coordType CoordType, xClipType, yClipType TextClipType) **TextRange {
	return textGetBoundedRanges(t, rect, coordType, xClipType, yClipType)
}
func (t *Text) GetCaretOffset() int                        { return textGetCaretOffset(t) }
func (t *Text) GetCharacterAtOffset(offset int) T.Gunichar { return textGetCharacterAtOffset(t, offset) }
func (t *Text) GetCharacterCount() int                     { return textGetCharacterCount(t) }
func (t *Text) GetCharacterExtents(offset int, x, y, width, height *int, coords CoordType) {
	textGetCharacterExtents(t, offset, x, y, width, height, coords)
}
func (t *Text) GetDefaultAttributes() *AttributeSet { return textGetDefaultAttributes(t) }
func (t *Text) GetNSelections() int                 { return textGetNSelections(t) }
func (t *Text) GetOffsetAtPoint(x, y int, coords CoordType) int {
	return textGetOffsetAtPoint(t, x, y, coords)
}
func (t *Text) GetRangeExtents(startOffset, endOffset int, coordType CoordType, rect *TextRectangle) {
	textGetRangeExtents(t, startOffset, endOffset, coordType, rect)
}
func (t *Text) GetRunAttributes(offset int, startOffset, endOffset *int) *AttributeSet {
	return textGetRunAttributes(t, offset, startOffset, endOffset)
}
func (t *Text) GetSelection(selectionNum int, startOffset, endOffset *int) string {
	return textGetSelection(t, selectionNum, startOffset, endOffset)
}
func (t *Text) GetText(startOffset, endOffset int) string {
	return textGetText(t, startOffset, endOffset)
}
func (t *Text) GetTextAfterOffset(offset int, boundaryType TextBoundary, startOffset, endOffset *int) string {
	return textGetTextAfterOffset(t, offset, boundaryType, startOffset, endOffset)
}
func (t *Text) GetTextAtOffset(offset int, boundaryType TextBoundary, startOffset, endOffset *int) string {
	return textGetTextAtOffset(t, offset, boundaryType, startOffset, endOffset)
}
func (t *Text) GetTextBeforeOffset(offset int, boundaryType TextBoundary, startOffset, endOffset *int) string {
	return textGetTextBeforeOffset(t, offset, boundaryType, startOffset, endOffset)
}
func (t *Text) RemoveSelection(selectionNum int) T.Gboolean {
	return textRemoveSelection(t, selectionNum)
}
func (t *Text) SetCaretOffset(offset int) T.Gboolean { return textSetCaretOffset(t, offset) }
func (t *Text) SetSelection(selectionNum, startOffset, endOffset int) T.Gboolean {
	return textSetSelection(t, selectionNum, startOffset, endOffset)
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

	textAttributeGetName  func(t TextAttribute) string
	textAttributeGetValue func(t TextAttribute, index int) string
)

func (t TextAttribute) GetName() string           { return textAttributeGetName(t) }
func (t TextAttribute) GetValue(index int) string { return textAttributeGetValue(t, index) }

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
