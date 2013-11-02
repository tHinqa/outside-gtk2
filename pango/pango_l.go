// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Language struct{}

var (
	LanguageGetType func() O.Type

	LanguageFromString func(language string) *Language
	LanguageGetDefault func() *Language

	LanguageGetSampleString func(l *Language) string
	LanguageGetScripts      func(l *Language, numScripts *int) *Script
	LanguageIncludesScript  func(l *Language, script Script) bool
	LanguageMatches         func(l *Language, rangeList string) bool
	LanguageToString        func(l *Language) string
)

func (l *Language) GetSampleString() string { return LanguageGetSampleString(l) }
func (l *Language) GetScripts(numScripts *int) *Script {
	return LanguageGetScripts(l, numScripts)
}
func (l *Language) IncludesScript(script Script) bool {
	return LanguageIncludesScript(l, script)
}
func (l *Language) Matches(rangeList string) bool { return LanguageMatches(l, rangeList) }
func (l *Language) ToString() string              { return LanguageToString(l) }

type Layout struct{}

var (
	LayoutGetType func() O.Type
	LayoutNew     func(context *Context) *Layout

	LayoutContextChanged         func(l *Layout)
	LayoutCopy                   func(l *Layout) *Layout
	LayoutGetAlignment           func(l *Layout) Alignment
	LayoutGetAttributes          func(l *Layout) *AttrList
	LayoutGetAutoDir             func(l *Layout) bool
	LayoutGetBaseline            func(l *Layout) int
	LayoutGetCharacterCount      func(l *Layout) int
	LayoutGetContext             func(l *Layout) *Context
	LayoutGetCursorPos           func(l *Layout, index int, strongPos *Rectangle, weakPos *Rectangle)
	LayoutGetEllipsize           func(l *Layout) EllipsizeMode
	LayoutGetExtents             func(l *Layout, inkRect *Rectangle, logicalRect *Rectangle)
	LayoutGetFontDescription     func(l *Layout) *FontDescription
	LayoutGetHeight              func(l *Layout) int
	LayoutGetIndent              func(l *Layout) int
	LayoutGetIter                func(l *Layout) *LayoutIter
	LayoutGetJustify             func(l *Layout) bool
	LayoutGetLine                func(l *Layout, line int) *LayoutLine
	LayoutGetLineCount           func(l *Layout) int
	LayoutGetLineReadonly        func(l *Layout, line int) *LayoutLine
	LayoutGetLines               func(l *Layout) *L.SList
	LayoutGetLinesReadonly       func(l *Layout) *L.SList
	LayoutGetLogAttrs            func(l *Layout, Attrs **LogAttr, nAttrs *int)
	LayoutGetLogAttrsReadonly    func(l *Layout, nAttrs *int) *LogAttr
	LayoutGetPixelExtents        func(l *Layout, inkRect *Rectangle, logicalRect *Rectangle)
	LayoutGetPixelSize           func(l *Layout, width *int, height *int)
	LayoutGetSingleParagraphMode func(l *Layout) bool
	LayoutGetSize                func(l *Layout, width *int, height *int)
	LayoutGetSpacing             func(l *Layout) int
	LayoutGetTabs                func(l *Layout) *TabArray
	LayoutGetText                func(l *Layout) string
	LayoutGetUnknownGlyphsCount  func(l *Layout) int
	LayoutGetWidth               func(l *Layout) int
	LayoutGetWrap                func(l *Layout) WrapMode
	LayoutIndexToLineX           func(l *Layout, index int, trailing bool, line *int, xPos *int)
	LayoutIndexToPos             func(l *Layout, index int, pos *Rectangle)
	LayoutIsEllipsized           func(l *Layout) bool
	LayoutIsWrapped              func(l *Layout) bool
	LayoutMoveCursorVisually     func(l *Layout, strong bool, oldIndex int, oldTrailing int, direction int, newIndex *int, newTrailing *int)
	LayoutSetAlignment           func(l *Layout, Alignment Alignment)
	LayoutSetAttributes          func(l *Layout, Attrs *AttrList)
	LayoutSetAutoDir             func(l *Layout, AutoDir bool)
	LayoutSetEllipsize           func(l *Layout, ellipsize EllipsizeMode)
	LayoutSetFontDescription     func(l *Layout, desc *FontDescription)
	LayoutSetHeight              func(l *Layout, height int)
	LayoutSetIndent              func(l *Layout, indent int)
	LayoutSetJustify             func(l *Layout, justify bool)
	LayoutSetMarkup              func(l *Layout, markup string, length int)
	LayoutSetMarkupWithAccel     func(l *Layout, markup string, length int, AccelMarker L.Unichar, AccelChar *L.Unichar)
	LayoutSetSingleParagraphMode func(l *Layout, setting bool)
	LayoutSetSpacing             func(l *Layout, spacing int)
	LayoutSetTabs                func(l *Layout, tabs *TabArray)
	LayoutSetText                func(l *Layout, text string, length int)
	LayoutSetWidth               func(l *Layout, width int)
	LayoutSetWrap                func(l *Layout, wrap WrapMode)
	LayoutXyToIndex              func(l *Layout, x int, y int, index *int, trailing *int) bool
)

func (l *Layout) ContextChanged()          { LayoutContextChanged(l) }
func (l *Layout) Copy() *Layout            { return LayoutCopy(l) }
func (l *Layout) GetAlignment() Alignment  { return LayoutGetAlignment(l) }
func (l *Layout) GetAttributes() *AttrList { return LayoutGetAttributes(l) }
func (l *Layout) GetAutoDir() bool         { return LayoutGetAutoDir(l) }
func (l *Layout) GetBaseline() int         { return LayoutGetBaseline(l) }
func (l *Layout) GetCharacterCount() int   { return LayoutGetCharacterCount(l) }
func (l *Layout) GetContext() *Context     { return LayoutGetContext(l) }
func (l *Layout) GetCursorPos(index int, strongPos *Rectangle, weakPos *Rectangle) {
	LayoutGetCursorPos(l, index, strongPos, weakPos)
}
func (l *Layout) GetEllipsize() EllipsizeMode { return LayoutGetEllipsize(l) }
func (l *Layout) GetExtents(inkRect *Rectangle, logicalRect *Rectangle) {
	LayoutGetExtents(l, inkRect, logicalRect)
}
func (l *Layout) GetFontDescription() *FontDescription     { return LayoutGetFontDescription(l) }
func (l *Layout) GetHeight() int                           { return LayoutGetHeight(l) }
func (l *Layout) GetIndent() int                           { return LayoutGetIndent(l) }
func (l *Layout) GetIter() *LayoutIter                     { return LayoutGetIter(l) }
func (l *Layout) GetJustify() bool                         { return LayoutGetJustify(l) }
func (l *Layout) GetLine(line int) *LayoutLine             { return LayoutGetLine(l, line) }
func (l *Layout) GetLineCount() int                        { return LayoutGetLineCount(l) }
func (l *Layout) GetLineReadonly(line int) *LayoutLine     { return LayoutGetLineReadonly(l, line) }
func (l *Layout) GetLines() *L.SList                       { return LayoutGetLines(l) }
func (l *Layout) GetLinesReadonly() *L.SList               { return LayoutGetLinesReadonly(l) }
func (l *Layout) GetLogAttrs(Attrs **LogAttr, nAttrs *int) { LayoutGetLogAttrs(l, Attrs, nAttrs) }
func (l *Layout) GetLogAttrsReadonly(nAttrs *int) *LogAttr {
	return LayoutGetLogAttrsReadonly(l, nAttrs)
}
func (l *Layout) GetPixelExtents(inkRect *Rectangle, logicalRect *Rectangle) {
	LayoutGetPixelExtents(l, inkRect, logicalRect)
}
func (l *Layout) GetPixelSize(width, height *int) { LayoutGetPixelSize(l, width, height) }
func (l *Layout) GetSingleParagraphMode() bool    { return LayoutGetSingleParagraphMode(l) }
func (l *Layout) GetSize(width, height *int)      { LayoutGetSize(l, width, height) }
func (l *Layout) GetSpacing() int                 { return LayoutGetSpacing(l) }
func (l *Layout) GetTabs() *TabArray              { return LayoutGetTabs(l) }
func (l *Layout) GetText() string                 { return LayoutGetText(l) }
func (l *Layout) GetUnknownGlyphsCount() int      { return LayoutGetUnknownGlyphsCount(l) }
func (l *Layout) GetWidth() int                   { return LayoutGetWidth(l) }
func (l *Layout) GetWrap() WrapMode               { return LayoutGetWrap(l) }
func (l *Layout) IndexToLineX(index int, trailing bool, line, xPos *int) {
	LayoutIndexToLineX(l, index, trailing, line, xPos)
}
func (l *Layout) IndexToPos(index int, pos *Rectangle) { LayoutIndexToPos(l, index, pos) }
func (l *Layout) IsEllipsized() bool                   { return LayoutIsEllipsized(l) }
func (l *Layout) IsWrapped() bool                      { return LayoutIsWrapped(l) }
func (l *Layout) MoveCursorVisually(strong bool, oldIndex, oldTrailing, direction int, newIndex, newTrailing *int) {
	LayoutMoveCursorVisually(l, strong, oldIndex, oldTrailing, direction, newIndex, newTrailing)
}
func (l *Layout) SetAlignment(Alignment Alignment)         { LayoutSetAlignment(l, Alignment) }
func (l *Layout) SetAttributes(Attrs *AttrList)            { LayoutSetAttributes(l, Attrs) }
func (l *Layout) SetAutoDir(AutoDir bool)                  { LayoutSetAutoDir(l, AutoDir) }
func (l *Layout) SetEllipsize(ellipsize EllipsizeMode)     { LayoutSetEllipsize(l, ellipsize) }
func (l *Layout) SetFontDescription(desc *FontDescription) { LayoutSetFontDescription(l, desc) }
func (l *Layout) SetHeight(height int)                     { LayoutSetHeight(l, height) }
func (l *Layout) SetIndent(indent int)                     { LayoutSetIndent(l, indent) }
func (l *Layout) SetJustify(justify bool)                  { LayoutSetJustify(l, justify) }
func (l *Layout) SetMarkup(markup string, length int)      { LayoutSetMarkup(l, markup, length) }
func (l *Layout) SetMarkupWithAccel(markup string, length int, AccelMarker L.Unichar, AccelChar *L.Unichar) {
	LayoutSetMarkupWithAccel(l, markup, length, AccelMarker, AccelChar)
}
func (l *Layout) SetSingleParagraphMode(setting bool) { LayoutSetSingleParagraphMode(l, setting) }
func (l *Layout) SetSpacing(spacing int)              { LayoutSetSpacing(l, spacing) }
func (l *Layout) SetTabs(tabs *TabArray)              { LayoutSetTabs(l, tabs) }
func (l *Layout) SetText(text string, length int)     { LayoutSetText(l, text, length) }
func (l *Layout) SetWidth(width int)                  { LayoutSetWidth(l, width) }
func (l *Layout) SetWrap(wrap WrapMode)               { LayoutSetWrap(l, wrap) }
func (l *Layout) XyToIndex(x, y int, index, trailing *int) bool {
	return LayoutXyToIndex(l, x, y, index, trailing)
}

type LayoutLine struct {
	Layout     *Layout
	StartIndex int
	Length     int
	Runs       *L.SList
	Bits       uint
	// IsParagraphStart : 1
	// ResolvedDir : 3
}

var (
	LayoutLineGetType func() O.Type

	LayoutLineGetExtents      func(l *LayoutLine, inkRect, logicalRect *Rectangle)
	LayoutLineGetPixelExtents func(l *LayoutLine, inkRect, logicalRect *Rectangle)
	LayoutLineGetXRanges      func(l *LayoutLine, startIndex, endIndex int, ranges **int, nRanges *int)
	LayoutLineIndexToX        func(l *LayoutLine, index int, trailing bool, xPos *int)
	LayoutLineRef             func(l *LayoutLine) *LayoutLine
	LayoutLineUnref           func(l *LayoutLine)
	LayoutLineXToIndex        func(l *LayoutLine, xPos int, index, trailing *int) bool
)

func (l *LayoutLine) GetExtents(inkRect, logicalRect *Rectangle) {
	LayoutLineGetExtents(l, inkRect, logicalRect)
}
func (l *LayoutLine) GetPixelExtents(inkRect, logicalRect *Rectangle) {
	LayoutLineGetPixelExtents(l, inkRect, logicalRect)
}
func (l *LayoutLine) GetXRanges(startIndex, endIndex int, ranges **int, nRanges *int) {
	LayoutLineGetXRanges(l, startIndex, endIndex, ranges, nRanges)
}
func (l *LayoutLine) IndexToX(index int, trailing bool, xPos *int) {
	LayoutLineIndexToX(l, index, trailing, xPos)
}
func (l *LayoutLine) Ref() *LayoutLine { return LayoutLineRef(l) }
func (l *LayoutLine) Unref()           { LayoutLineUnref(l) }
func (l *LayoutLine) XToIndex(xPos int, index, trailing *int) bool {
	return LayoutLineXToIndex(l, xPos, index, trailing)
}

type LayoutIter struct{}

var (
	LayoutIterGetType func() O.Type

	LayoutIterAtLastLine        func(l *LayoutIter) bool
	LayoutIterCopy              func(l *LayoutIter) *LayoutIter
	LayoutIterFree              func(l *LayoutIter)
	LayoutIterGetBaseline       func(l *LayoutIter) int
	LayoutIterGetCharExtents    func(l *LayoutIter, logicalRect *Rectangle)
	LayoutIterGetClusterExtents func(l *LayoutIter, inkRect, logicalRect *Rectangle)
	LayoutIterGetIndex          func(l *LayoutIter) int
	LayoutIterGetLayout         func(l *LayoutIter) *Layout
	LayoutIterGetLayoutExtents  func(l *LayoutIter, inkRect, logicalRect *Rectangle)
	LayoutIterGetLine           func(l *LayoutIter) *LayoutLine
	LayoutIterGetLineExtents    func(l *LayoutIter, inkRect, logicalRect *Rectangle)
	LayoutIterGetLineReadonly   func(l *LayoutIter) *LayoutLine
	LayoutIterGetLineYrange     func(l *LayoutIter, y0, y1 *int)
	LayoutIterGetRun            func(l *LayoutIter) *LayoutRun
	LayoutIterGetRunExtents     func(l *LayoutIter, inkRect, logicalRect *Rectangle)
	LayoutIterGetRunReadonly    func(l *LayoutIter) *LayoutRun
	LayoutIterNextChar          func(l *LayoutIter) bool
	LayoutIterNextCluster       func(l *LayoutIter) bool
	LayoutIterNextLine          func(l *LayoutIter) bool
	LayoutIterNextRun           func(l *LayoutIter) bool
)

func (l *LayoutIter) AtLastLine() bool  { return LayoutIterAtLastLine(l) }
func (l *LayoutIter) Copy() *LayoutIter { return LayoutIterCopy(l) }
func (l *LayoutIter) Free()             { LayoutIterFree(l) }
func (l *LayoutIter) GetBaseline() int  { return LayoutIterGetBaseline(l) }
func (l *LayoutIter) GetCharExtents(logicalRect *Rectangle) {
	LayoutIterGetCharExtents(l, logicalRect)
}
func (l *LayoutIter) GetClusterExtents(inkRect, logicalRect *Rectangle) {
	LayoutIterGetClusterExtents(l, inkRect, logicalRect)
}
func (l *LayoutIter) GetIndex() int      { return LayoutIterGetIndex(l) }
func (l *LayoutIter) GetLayout() *Layout { return LayoutIterGetLayout(l) }
func (l *LayoutIter) GetLayoutExtents(inkRect, logicalRect *Rectangle) {
	LayoutIterGetLayoutExtents(l, inkRect, logicalRect)
}
func (l *LayoutIter) GetLine() *LayoutLine { return LayoutIterGetLine(l) }
func (l *LayoutIter) GetLineExtents(inkRect, logicalRect *Rectangle) {
	LayoutIterGetLineExtents(l, inkRect, logicalRect)
}
func (l *LayoutIter) GetLineReadonly() *LayoutLine { return LayoutIterGetLineReadonly(l) }
func (l *LayoutIter) GetLineYrange(y0, y1 *int)    { LayoutIterGetLineYrange(l, y0, y1) }
func (l *LayoutIter) GetRun() *LayoutRun           { return LayoutIterGetRun(l) }
func (l *LayoutIter) GetRunExtents(inkRect, logicalRect *Rectangle) {
	LayoutIterGetRunExtents(l, inkRect, logicalRect)
}
func (l *LayoutIter) GetRunReadonly() *LayoutRun { return LayoutIterGetRunReadonly(l) }
func (l *LayoutIter) NextChar() bool             { return LayoutIterNextChar(l) }
func (l *LayoutIter) NextCluster() bool          { return LayoutIterNextCluster(l) }
func (l *LayoutIter) NextLine() bool             { return LayoutIterNextLine(l) }
func (l *LayoutIter) NextRun() bool              { return LayoutIterNextRun(l) }

type LayoutRun GlyphItem

var Log2visGetEmbeddingLevels func(text string,
	Length int, pbaseDir *Direction) *uint8

type LogAttr struct {
	Bits uint
	// IsLineBreak : 1
	// IsMandatoryBreak : 1
	// IsCharBreak : 1
	// IsWhite : 1
	// IsCursorPosition : 1
	// IsWordStart : 1
	// IsWordEnd : 1
	// IsSentenceBoundary : 1
	// IsSentenceStart : 1
	// IsSentenceEnd : 1
	// BackspaceDeletesCharacter : 1
	// IsExpandableSpace : 1
	// IsWordBoundary : 1
}

var LookupAliases func(
	fontname string, families ***T.Char, nFamilies *int)
