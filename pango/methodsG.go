// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

var (
	GetLibSubdirectory     func() string
	GetLogAttrs            func(text string, length, level int, language *Language, logAttrs *LogAttr, attrsLen int)
	GetMirrorChar          func(ch T.Gunichar, mirroredCh *T.Gunichar) T.Gboolean
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

	glyphItemApplyAttrs       func(glyphItem *GlyphItem, text string, list *AttrList) *T.GSList
	glyphItemCopy             func(orig *GlyphItem) *GlyphItem
	glyphItemFree             func(glyphItem *GlyphItem)
	glyphItemGetLogicalWidths func(glyphItem *GlyphItem, text string, logicalWidths *int)
	glyphItemLetterSpace      func(glyphItem *GlyphItem, text string, logAttrs *LogAttr, letterSpacing int)
	glyphItemSplit            func(orig *GlyphItem, text string, splitIndex int) *GlyphItem
)

func (g *GlyphItem) ApplyAttrs(text string, list *AttrList) *T.GSList {
	return glyphItemApplyAttrs(g, text, list)
}
func (g *GlyphItem) Copy() *GlyphItem { return glyphItemCopy(g) }
func (g *GlyphItem) Free()            { glyphItemFree(g) }
func (g *GlyphItem) GetLogicalWidths(text string, logicalWidths *int) {
	glyphItemGetLogicalWidths(g, text, logicalWidths)
}
func (g *GlyphItem) LetterSpace(text string, logAttrs *LogAttr, letterSpacing int) {
	glyphItemLetterSpace(g, text, logAttrs, letterSpacing)
}
func (g *GlyphItem) Split(text string, splitIndex int) *GlyphItem {
	return glyphItemSplit(g, text, splitIndex)
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

	glyphItemIterCopy        func(orig *GlyphItemIter) *GlyphItemIter
	glyphItemIterFree        func(iter *GlyphItemIter)
	glyphItemIterInitEnd     func(iter *GlyphItemIter, glyphItem *GlyphItem, text string) T.Gboolean
	glyphItemIterInitStart   func(iter *GlyphItemIter, glyphItem *GlyphItem, text string) T.Gboolean
	glyphItemIterNextCluster func(iter *GlyphItemIter) T.Gboolean
	glyphItemIterPrevCluster func(iter *GlyphItemIter) T.Gboolean
)

func (g *GlyphItemIter) Copy() *GlyphItemIter { return glyphItemIterCopy(g) }
func (g *GlyphItemIter) Free()                { glyphItemIterFree(g) }
func (g *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) T.Gboolean {
	return glyphItemIterInitEnd(g, glyphItem, text)
}
func (g *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) T.Gboolean {
	return glyphItemIterInitStart(g, glyphItem, text)
}
func (g *GlyphItemIter) NextCluster() T.Gboolean { return glyphItemIterNextCluster(g) }
func (g *GlyphItemIter) PrevCluster() T.Gboolean { return glyphItemIterPrevCluster(g) }

type GlyphString struct {
	NumGlyphs   int
	Glyphs      *GlyphInfo
	LogClusters *int
	Space       int
}

var (
	GlyphStringGetType func() O.Type
	GlyphStringNew     func() *GlyphString

	glyphStringCopy             func(g *GlyphString) *GlyphString
	glyphStringExtents          func(g *GlyphString, font *Font, inkRect *Rectangle, logicalRect *Rectangle)
	glyphStringExtentsRange     func(g *GlyphString, start, end int, font *Font, inkRect *Rectangle, logicalRect *Rectangle)
	glyphStringFree             func(g *GlyphString)
	glyphStringGetLogicalWidths func(g *GlyphString, text string, length, embeddingLevel int, logicalWidths *int)
	glyphStringGetWidth         func(g *GlyphString) int
	glyphStringIndexToX         func(g *GlyphString, text string, length int, Analysis *Analysis, index int, trailing T.Gboolean, xPos *int)
	glyphStringSetSize          func(g *GlyphString, newLen int)
	glyphStringXToIndex         func(g *GlyphString, text string, length int, Analysis *Analysis, xPos int, index, trailing *int)
)

func (g *GlyphString) Copy() *GlyphString { return glyphStringCopy(g) }
func (g *GlyphString) Extents(font *Font, inkRect *Rectangle, logicalRect *Rectangle) {
	glyphStringExtents(g, font, inkRect, logicalRect)
}
func (g *GlyphString) ExtentsRange(start, end int, font *Font, inkRect *Rectangle, logicalRect *Rectangle) {
	glyphStringExtentsRange(g, start, end, font, inkRect, logicalRect)
}
func (g *GlyphString) Free() { glyphStringFree(g) }
func (g *GlyphString) GetLogicalWidths(text string, length, embeddingLevel int, logicalWidths *int) {
	glyphStringGetLogicalWidths(g, text, length, embeddingLevel, logicalWidths)
}
func (g *GlyphString) GetWidth() int { return glyphStringGetWidth(g) }
func (g *GlyphString) IndexToX(text string, length int, Analysis *Analysis, index int, trailing T.Gboolean, xPos *int) {
	glyphStringIndexToX(g, text, length, Analysis, index, trailing, xPos)
}
func (g *GlyphString) SetSize(newLen int) { glyphStringSetSize(g, newLen) }
func (g *GlyphString) XToIndex(text string, length int, Analysis *Analysis, xPos int, index, trailing *int) {
	glyphStringXToIndex(g, text, length, Analysis, xPos, index, trailing)
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
	GravityGetForScriptAndWidth func(script Script, wide T.Gboolean, baseGravity Gravity, hint GravityHint) Gravity

	gravityToRotation func(g Gravity) float64
)

func (g Gravity) ToRotation() float64 { return gravityToRotation(g) }

type GravityHint Enum

const (
	GRAVITY_HINT_NATURAL GravityHint = iota
	GRAVITY_HINT_STRONG
	GRAVITY_HINT_LINE
)

var GravityHintGetType func() O.Type
