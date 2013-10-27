// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Language struct{}

var (
	LanguageGetType func() O.Type

	LanguageFromString func(language string) *Language
	LanguageGetDefault func() *Language

	languageGetSampleString func(l *Language) string
	languageGetScripts      func(l *Language, numScripts *int) *Script
	languageIncludesScript  func(l *Language, script Script) T.Gboolean
	languageMatches         func(l *Language, rangeList string) T.Gboolean
	languageToString        func(l *Language) string
)

func (l *Language) GetSampleString() string { return languageGetSampleString(l) }
func (l *Language) GetScripts(numScripts *int) *Script {
	return languageGetScripts(l, numScripts)
}
func (l *Language) IncludesScript(script Script) T.Gboolean {
	return languageIncludesScript(l, script)
}
func (l *Language) Matches(rangeList string) T.Gboolean { return languageMatches(l, rangeList) }
func (l *Language) ToString() string                    { return languageToString(l) }

type Layout struct{}

var (
	layoutGetType func() O.Type
	layoutNew     func(context *Context) *Layout

	layoutContextChanged         func(l *Layout)
	layoutCopy                   func(l *Layout) *Layout
	layoutGetAlignment           func(l *Layout) Alignment
	layoutGetAttributes          func(l *Layout) *AttrList
	layoutGetAutoDir             func(l *Layout) T.Gboolean
	layoutGetBaseline            func(l *Layout) int
	layoutGetCharacterCount      func(l *Layout) int
	layoutGetContext             func(l *Layout) *Context
	layoutGetCursorPos           func(l *Layout, index int, strongPos *Rectangle, weakPos *Rectangle)
	layoutGetEllipsize           func(l *Layout) EllipsizeMode
	layoutGetExtents             func(l *Layout, inkRect *Rectangle, logicalRect *Rectangle)
	layoutGetFontDescription     func(l *Layout) *FontDescription
	layoutGetHeight              func(l *Layout) int
	layoutGetIndent              func(l *Layout) int
	layoutGetIter                func(l *Layout) *LayoutIter
	layoutGetJustify             func(l *Layout) T.Gboolean
	layoutGetLine                func(l *Layout, line int) *LayoutLine
	layoutGetLineCount           func(l *Layout) int
	layoutGetLineReadonly        func(l *Layout, line int) *LayoutLine
	layoutGetLines               func(l *Layout) *T.GSList
	layoutGetLinesReadonly       func(l *Layout) *T.GSList
	layoutGetLogAttrs            func(l *Layout, Attrs **LogAttr, nAttrs *int)
	layoutGetLogAttrsReadonly    func(l *Layout, nAttrs *int) *LogAttr
	layoutGetPixelExtents        func(l *Layout, inkRect *Rectangle, logicalRect *Rectangle)
	layoutGetPixelSize           func(l *Layout, width *int, height *int)
	layoutGetSingleParagraphMode func(l *Layout) T.Gboolean
	layoutGetSize                func(l *Layout, width *int, height *int)
	layoutGetSpacing             func(l *Layout) int
	layoutGetTabs                func(l *Layout) *TabArray
	layoutGetText                func(l *Layout) string
	layoutGetUnknownGlyphsCount  func(l *Layout) int
	layoutGetWidth               func(l *Layout) int
	layoutGetWrap                func(l *Layout) WrapMode
	layoutIndexToLineX           func(l *Layout, index int, trailing T.Gboolean, line *int, xPos *int)
	layoutIndexToPos             func(l *Layout, index int, pos *Rectangle)
	layoutIsEllipsized           func(l *Layout) T.Gboolean
	layoutIsWrapped              func(l *Layout) T.Gboolean
	layoutMoveCursorVisually     func(l *Layout, strong T.Gboolean, oldIndex int, oldTrailing int, direction int, newIndex *int, newTrailing *int)
	layoutSetAlignment           func(l *Layout, Alignment Alignment)
	layoutSetAttributes          func(l *Layout, Attrs *AttrList)
	layoutSetAutoDir             func(l *Layout, AutoDir T.Gboolean)
	layoutSetEllipsize           func(l *Layout, ellipsize EllipsizeMode)
	layoutSetFontDescription     func(l *Layout, desc *FontDescription)
	layoutSetHeight              func(l *Layout, height int)
	layoutSetIndent              func(l *Layout, indent int)
	layoutSetJustify             func(l *Layout, justify T.Gboolean)
	layoutSetMarkup              func(l *Layout, markup string, length int)
	layoutSetMarkupWithAccel     func(l *Layout, markup string, length int, AccelMarker T.Gunichar, AccelChar *T.Gunichar)
	layoutSetSingleParagraphMode func(l *Layout, setting T.Gboolean)
	layoutSetSpacing             func(l *Layout, spacing int)
	layoutSetTabs                func(l *Layout, tabs *TabArray)
	layoutSetText                func(l *Layout, text string, length int)
	layoutSetWidth               func(l *Layout, width int)
	layoutSetWrap                func(l *Layout, wrap WrapMode)
	layoutXyToIndex              func(l *Layout, x int, y int, index *int, trailing *int) T.Gboolean
)

func (l *Layout) ContextChanged()          { layoutContextChanged(l) }
func (l *Layout) Copy() *Layout            { return layoutCopy(l) }
func (l *Layout) GetAlignment() Alignment  { return layoutGetAlignment(l) }
func (l *Layout) GetAttributes() *AttrList { return layoutGetAttributes(l) }
func (l *Layout) GetAutoDir() T.Gboolean   { return layoutGetAutoDir(l) }
func (l *Layout) GetBaseline() int         { return layoutGetBaseline(l) }
func (l *Layout) GetCharacterCount() int   { return layoutGetCharacterCount(l) }
func (l *Layout) GetContext() *Context     { return layoutGetContext(l) }
func (l *Layout) GetCursorPos(index int, strongPos *Rectangle, weakPos *Rectangle) {
	layoutGetCursorPos(l, index, strongPos, weakPos)
}
func (l *Layout) GetEllipsize() EllipsizeMode { return layoutGetEllipsize(l) }
func (l *Layout) GetExtents(inkRect *Rectangle, logicalRect *Rectangle) {
	layoutGetExtents(l, inkRect, logicalRect)
}
func (l *Layout) GetFontDescription() *FontDescription     { return layoutGetFontDescription(l) }
func (l *Layout) GetHeight() int                           { return layoutGetHeight(l) }
func (l *Layout) GetIndent() int                           { return layoutGetIndent(l) }
func (l *Layout) GetIter() *LayoutIter                     { return layoutGetIter(l) }
func (l *Layout) GetJustify() T.Gboolean                   { return layoutGetJustify(l) }
func (l *Layout) GetLine(line int) *LayoutLine             { return layoutGetLine(l, line) }
func (l *Layout) GetLineCount() int                        { return layoutGetLineCount(l) }
func (l *Layout) GetLineReadonly(line int) *LayoutLine     { return layoutGetLineReadonly(l, line) }
func (l *Layout) GetLines() *T.GSList                      { return layoutGetLines(l) }
func (l *Layout) GetLinesReadonly() *T.GSList              { return layoutGetLinesReadonly(l) }
func (l *Layout) GetLogAttrs(Attrs **LogAttr, nAttrs *int) { layoutGetLogAttrs(l, Attrs, nAttrs) }
func (l *Layout) GetLogAttrsReadonly(nAttrs *int) *LogAttr {
	return layoutGetLogAttrsReadonly(l, nAttrs)
}
func (l *Layout) GetPixelExtents(inkRect *Rectangle, logicalRect *Rectangle) {
	layoutGetPixelExtents(l, inkRect, logicalRect)
}
func (l *Layout) GetPixelSize(width, height *int)    { layoutGetPixelSize(l, width, height) }
func (l *Layout) GetSingleParagraphMode() T.Gboolean { return layoutGetSingleParagraphMode(l) }
func (l *Layout) GetSize(width, height *int)         { layoutGetSize(l, width, height) }
func (l *Layout) GetSpacing() int                    { return layoutGetSpacing(l) }
func (l *Layout) GetTabs() *TabArray                 { return layoutGetTabs(l) }
func (l *Layout) GetText() string                    { return layoutGetText(l) }
func (l *Layout) GetUnknownGlyphsCount() int         { return layoutGetUnknownGlyphsCount(l) }
func (l *Layout) GetWidth() int                      { return layoutGetWidth(l) }
func (l *Layout) GetWrap() WrapMode                  { return layoutGetWrap(l) }
func (l *Layout) IndexToLineX(index int, trailing T.Gboolean, line, xPos *int) {
	layoutIndexToLineX(l, index, trailing, line, xPos)
}
func (l *Layout) IndexToPos(index int, pos *Rectangle) { layoutIndexToPos(l, index, pos) }
func (l *Layout) IsEllipsized() T.Gboolean             { return layoutIsEllipsized(l) }
func (l *Layout) IsWrapped() T.Gboolean                { return layoutIsWrapped(l) }
func (l *Layout) MoveCursorVisually(strong T.Gboolean, oldIndex, oldTrailing, direction int, newIndex, newTrailing *int) {
	layoutMoveCursorVisually(l, strong, oldIndex, oldTrailing, direction, newIndex, newTrailing)
}
func (l *Layout) SetAlignment(Alignment Alignment)         { layoutSetAlignment(l, Alignment) }
func (l *Layout) SetAttributes(Attrs *AttrList)            { layoutSetAttributes(l, Attrs) }
func (l *Layout) SetAutoDir(AutoDir T.Gboolean)            { layoutSetAutoDir(l, AutoDir) }
func (l *Layout) SetEllipsize(ellipsize EllipsizeMode)     { layoutSetEllipsize(l, ellipsize) }
func (l *Layout) SetFontDescription(desc *FontDescription) { layoutSetFontDescription(l, desc) }
func (l *Layout) SetHeight(height int)                     { layoutSetHeight(l, height) }
func (l *Layout) SetIndent(indent int)                     { layoutSetIndent(l, indent) }
func (l *Layout) SetJustify(justify T.Gboolean)            { layoutSetJustify(l, justify) }
func (l *Layout) SetMarkup(markup string, length int)      { layoutSetMarkup(l, markup, length) }
func (l *Layout) SetMarkupWithAccel(markup string, length int, AccelMarker T.Gunichar, AccelChar *T.Gunichar) {
	layoutSetMarkupWithAccel(l, markup, length, AccelMarker, AccelChar)
}
func (l *Layout) SetSingleParagraphMode(setting T.Gboolean) { layoutSetSingleParagraphMode(l, setting) }
func (l *Layout) SetSpacing(spacing int)                    { layoutSetSpacing(l, spacing) }
func (l *Layout) SetTabs(tabs *TabArray)                    { layoutSetTabs(l, tabs) }
func (l *Layout) SetText(text string, length int)           { layoutSetText(l, text, length) }
func (l *Layout) SetWidth(width int)                        { layoutSetWidth(l, width) }
func (l *Layout) SetWrap(wrap WrapMode)                     { layoutSetWrap(l, wrap) }
func (l *Layout) XyToIndex(x, y int, index, trailing *int) T.Gboolean {
	return layoutXyToIndex(l, x, y, index, trailing)
}

type LayoutLine struct {
	Layout     *Layout
	StartIndex int
	Length     int
	Runs       *T.GSList
	Bits       uint
	// IsParagraphStart : 1
	// ResolvedDir : 3
}

var (
	LayoutLineGetType func() O.Type

	layoutLineGetExtents      func(l *LayoutLine, inkRect, logicalRect *Rectangle)
	layoutLineGetPixelExtents func(l *LayoutLine, inkRect, logicalRect *Rectangle)
	layoutLineGetXRanges      func(l *LayoutLine, startIndex, endIndex int, ranges **int, nRanges *int)
	layoutLineIndexToX        func(l *LayoutLine, index int, trailing T.Gboolean, xPos *int)
	layoutLineRef             func(l *LayoutLine) *LayoutLine
	layoutLineUnref           func(l *LayoutLine)
	layoutLineXToIndex        func(l *LayoutLine, xPos int, index, trailing *int) T.Gboolean
)

func (l *LayoutLine) GetExtents(inkRect, logicalRect *Rectangle) {
	layoutLineGetExtents(l, inkRect, logicalRect)
}
func (l *LayoutLine) GetPixelExtents(inkRect, logicalRect *Rectangle) {
	layoutLineGetPixelExtents(l, inkRect, logicalRect)
}
func (l *LayoutLine) GetXRanges(startIndex, endIndex int, ranges **int, nRanges *int) {
	layoutLineGetXRanges(l, startIndex, endIndex, ranges, nRanges)
}
func (l *LayoutLine) IndexToX(index int, trailing T.Gboolean, xPos *int) {
	layoutLineIndexToX(l, index, trailing, xPos)
}
func (l *LayoutLine) Ref() *LayoutLine { return layoutLineRef(l) }
func (l *LayoutLine) Unref()           { layoutLineUnref(l) }
func (l *LayoutLine) XToIndex(xPos int, index, trailing *int) T.Gboolean {
	return layoutLineXToIndex(l, xPos, index, trailing)
}

type LayoutIter struct{}

var (
	LayoutIterGetType func() O.Type

	layoutIterAtLastLine        func(l *LayoutIter) T.Gboolean
	layoutIterCopy              func(l *LayoutIter) *LayoutIter
	layoutIterFree              func(l *LayoutIter)
	layoutIterGetBaseline       func(l *LayoutIter) int
	layoutIterGetCharExtents    func(l *LayoutIter, logicalRect *Rectangle)
	layoutIterGetClusterExtents func(l *LayoutIter, inkRect, logicalRect *Rectangle)
	layoutIterGetIndex          func(l *LayoutIter) int
	layoutIterGetLayout         func(l *LayoutIter) *Layout
	layoutIterGetLayoutExtents  func(l *LayoutIter, inkRect, logicalRect *Rectangle)
	layoutIterGetLine           func(l *LayoutIter) *LayoutLine
	layoutIterGetLineExtents    func(l *LayoutIter, inkRect, logicalRect *Rectangle)
	layoutIterGetLineReadonly   func(l *LayoutIter) *LayoutLine
	layoutIterGetLineYrange     func(l *LayoutIter, y0, y1 *int)
	layoutIterGetRun            func(l *LayoutIter) *LayoutRun
	layoutIterGetRunExtents     func(l *LayoutIter, inkRect, logicalRect *Rectangle)
	layoutIterGetRunReadonly    func(l *LayoutIter) *LayoutRun
	layoutIterNextChar          func(l *LayoutIter) T.Gboolean
	layoutIterNextCluster       func(l *LayoutIter) T.Gboolean
	layoutIterNextLine          func(l *LayoutIter) T.Gboolean
	layoutIterNextRun           func(l *LayoutIter) T.Gboolean
)

func (l *LayoutIter) AtLastLine() T.Gboolean { return layoutIterAtLastLine(l) }
func (l *LayoutIter) Copy() *LayoutIter      { return layoutIterCopy(l) }
func (l *LayoutIter) Free()                  { layoutIterFree(l) }
func (l *LayoutIter) GetBaseline() int       { return layoutIterGetBaseline(l) }
func (l *LayoutIter) GetCharExtents(logicalRect *Rectangle) {
	layoutIterGetCharExtents(l, logicalRect)
}
func (l *LayoutIter) GetClusterExtents(inkRect, logicalRect *Rectangle) {
	layoutIterGetClusterExtents(l, inkRect, logicalRect)
}
func (l *LayoutIter) GetIndex() int      { return layoutIterGetIndex(l) }
func (l *LayoutIter) GetLayout() *Layout { return layoutIterGetLayout(l) }
func (l *LayoutIter) GetLayoutExtents(inkRect, logicalRect *Rectangle) {
	layoutIterGetLayoutExtents(l, inkRect, logicalRect)
}
func (l *LayoutIter) GetLine() *LayoutLine { return layoutIterGetLine(l) }
func (l *LayoutIter) GetLineExtents(inkRect, logicalRect *Rectangle) {
	layoutIterGetLineExtents(l, inkRect, logicalRect)
}
func (l *LayoutIter) GetLineReadonly() *LayoutLine { return layoutIterGetLineReadonly(l) }
func (l *LayoutIter) GetLineYrange(y0, y1 *int)    { layoutIterGetLineYrange(l, y0, y1) }
func (l *LayoutIter) GetRun() *LayoutRun           { return layoutIterGetRun(l) }
func (l *LayoutIter) GetRunExtents(inkRect, logicalRect *Rectangle) {
	layoutIterGetRunExtents(l, inkRect, logicalRect)
}
func (l *LayoutIter) GetRunReadonly() *LayoutRun { return layoutIterGetRunReadonly(l) }
func (l *LayoutIter) NextChar() T.Gboolean       { return layoutIterNextChar(l) }
func (l *LayoutIter) NextCluster() T.Gboolean    { return layoutIterNextCluster(l) }
func (l *LayoutIter) NextLine() T.Gboolean       { return layoutIterNextLine(l) }
func (l *LayoutIter) NextRun() T.Gboolean        { return layoutIterNextRun(l) }

type LayoutRun GlyphItem

var Log2visGetEmbeddingLevels func(text string,
	length int, pbaseDir *Direction) *uint8

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
