// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

var (
	GetLibSubdirectory     func() string
	GetLogAttrs            func(text string, length, level int, language *Language, logAttrs *LogAttr, attrsLen int)
	GetMirrorChar          func(ch L.Unichar, mirroredCh *L.Unichar) bool
	GetSysconfSubdirectory func() string
)

type (
	Glyph T.GUint32

	GlyphUnit T.GInt32

	GlyphGeometry struct {
		Width   GlyphUnit
		XOffset GlyphUnit
		YOffset GlyphUnit
	}

	GlyphInfo struct {
		Glyph    Glyph
		Geometry GlyphGeometry
		Attr     GlyphVisAttr
	}

	GlyphItem struct {
		Item   *Item
		Glyphs *GlyphString
	}

	GlyphVisAttr struct {
		IsClusterStart uint // : 1
	}
)

var (
	GlyphItemGetType func() O.Type

	GlyphItemApplyAttrs       func(glyphItem *GlyphItem, text string, list *AttrList) *L.SList
	GlyphItemCopy             func(orig *GlyphItem) *GlyphItem
	GlyphItemFree             func(glyphItem *GlyphItem)
	GlyphItemGetLogicalWidths func(glyphItem *GlyphItem, text string, logicalWidths *int)
	GlyphItemLetterSpace      func(glyphItem *GlyphItem, text string, logAttrs *LogAttr, letterSpacing int)
	GlyphItemSplit            func(orig *GlyphItem, text string, splitIndex int) *GlyphItem
)

func (g *GlyphItem) ApplyAttrs(text string, list *AttrList) *L.SList {
	return GlyphItemApplyAttrs(g, text, list)
}
func (g *GlyphItem) Copy() *GlyphItem { return GlyphItemCopy(g) }
func (g *GlyphItem) Free()            { GlyphItemFree(g) }
func (g *GlyphItem) GetLogicalWidths(text string, logicalWidths *int) {
	GlyphItemGetLogicalWidths(g, text, logicalWidths)
}
func (g *GlyphItem) LetterSpace(text string, logAttrs *LogAttr, letterSpacing int) {
	GlyphItemLetterSpace(g, text, logAttrs, letterSpacing)
}
func (g *GlyphItem) Split(text string, splitIndex int) *GlyphItem {
	return GlyphItemSplit(g, text, splitIndex)
}

type GlyphItemIter struct {
	GlyphItem  *GlyphItem
	Text       *T.Gchar
	StartGlyph int
	StartIndex int
	StartChar  int
	EndGlyph   int
	EndIndex   int
	EndChar    int
}

var (
	GlyphItemIterGetType func() O.Type

	GlyphItemIterCopy        func(orig *GlyphItemIter) *GlyphItemIter
	GlyphItemIterFree        func(iter *GlyphItemIter)
	GlyphItemIterInitEnd     func(iter *GlyphItemIter, glyphItem *GlyphItem, text string) bool
	GlyphItemIterInitStart   func(iter *GlyphItemIter, glyphItem *GlyphItem, text string) bool
	GlyphItemIterNextCluster func(iter *GlyphItemIter) bool
	GlyphItemIterPrevCluster func(iter *GlyphItemIter) bool
)

func (g *GlyphItemIter) Copy() *GlyphItemIter { return GlyphItemIterCopy(g) }
func (g *GlyphItemIter) Free()                { GlyphItemIterFree(g) }
func (g *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	return GlyphItemIterInitEnd(g, glyphItem, text)
}
func (g *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	return GlyphItemIterInitStart(g, glyphItem, text)
}
func (g *GlyphItemIter) NextCluster() bool { return GlyphItemIterNextCluster(g) }
func (g *GlyphItemIter) PrevCluster() bool { return GlyphItemIterPrevCluster(g) }

type GlyphString struct {
	NumGlyphs   int
	Glyphs      *GlyphInfo
	LogClusters *int
	Space       int
}

var (
	GlyphStringGetType func() O.Type
	GlyphStringNew     func() *GlyphString

	GlyphStringCopy             func(g *GlyphString) *GlyphString
	GlyphStringExtents          func(g *GlyphString, font *Font, inkRect *Rectangle, logicalRect *Rectangle)
	GlyphStringExtentsRange     func(g *GlyphString, start, end int, font *Font, inkRect *Rectangle, logicalRect *Rectangle)
	GlyphStringFree             func(g *GlyphString)
	GlyphStringGetLogicalWidths func(g *GlyphString, text string, length, embeddingLevel int, logicalWidths *int)
	GlyphStringGetWidth         func(g *GlyphString) int
	GlyphStringIndexToX         func(g *GlyphString, text string, length int, Analysis *Analysis, index int, trailing bool, xPos *int)
	GlyphStringSetSize          func(g *GlyphString, newLen int)
	GlyphStringXToIndex         func(g *GlyphString, text string, length int, Analysis *Analysis, xPos int, index, trailing *int)
)

func (g *GlyphString) Copy() *GlyphString { return GlyphStringCopy(g) }
func (g *GlyphString) Extents(font *Font, inkRect *Rectangle, logicalRect *Rectangle) {
	GlyphStringExtents(g, font, inkRect, logicalRect)
}
func (g *GlyphString) ExtentsRange(start, end int, font *Font, inkRect *Rectangle, logicalRect *Rectangle) {
	GlyphStringExtentsRange(g, start, end, font, inkRect, logicalRect)
}
func (g *GlyphString) Free() { GlyphStringFree(g) }
func (g *GlyphString) GetLogicalWidths(text string, length, embeddingLevel int, logicalWidths *int) {
	GlyphStringGetLogicalWidths(g, text, length, embeddingLevel, logicalWidths)
}
func (g *GlyphString) GetWidth() int { return GlyphStringGetWidth(g) }
func (g *GlyphString) IndexToX(text string, length int, Analysis *Analysis, index int, trailing bool, xPos *int) {
	GlyphStringIndexToX(g, text, length, Analysis, index, trailing, xPos)
}
func (g *GlyphString) SetSize(newLen int) { GlyphStringSetSize(g, newLen) }
func (g *GlyphString) XToIndex(text string, length int, Analysis *Analysis, xPos int, index, trailing *int) {
	GlyphStringXToIndex(g, text, length, Analysis, xPos, index, trailing)
}

type Gravity Enum

const (
	GRAVITY_SOUTH Gravity = iota
	GRAVITY_EAST
	GRAVITY_NORTH
	GRAVITY_WEST
	GRAVITY_AUTO
)

var (
	GravityGetType func() O.Type

	GravityGetForMatrix         func(matrix *Matrix) Gravity
	GravityGetForScript         func(script Script, baseGravity Gravity, hint GravityHint) Gravity
	GravityGetForScriptAndWidth func(script Script, wide bool, baseGravity Gravity, hint GravityHint) Gravity

	GravityToRotation func(g Gravity) float64
)

func (g Gravity) ToRotation() float64 { return GravityToRotation(g) }

type GravityHint Enum

const (
	GRAVITY_HINT_NATURAL GravityHint = iota
	GRAVITY_HINT_STRONG
	GRAVITY_HINT_LINE
)

var GravityHintGetType func() O.Type
