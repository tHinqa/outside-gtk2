// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	C "github.com/tHinqa/outside-gtk2/cairo"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type CairoFont struct{}

type CairoFontMap struct{}

var (
	CairoFontMapGetType func() O.Type
	CairoFontMapNew     func() *FontMap

	CairoFontMapNewForFontType func(fonttype C.FontType) *FontMap
	CairoFontMapGetDefault     func() *FontMap

	CairoFontGetType       func() O.Type
	CairoFontGetScaledFont func(font *CairoFont) *C.ScaledFont

	CairoFontMapSetDefault    func(fontmap *CairoFontMap)
	CairoFontMapGetFontType   func(fontmap *CairoFontMap) C.FontType
	CairoFontMapSetResolution func(fontmap *CairoFontMap, dpi float64)
	CairoFontMapGetResolution func(fontmap *CairoFontMap) float64
	CairoFontMapCreateContext func(fontmap *CairoFontMap) *Context
)

var (
	CairoContextSetFontOptions   func(c *Context, options *C.FontOptions)
	CairoContextGetFontOptions   func(c *Context) *C.FontOptions
	CairoContextSetResolution    func(c *Context, dpi float64)
	CairoContextGetResolution    func(c *Context) float64
	CairoContextSetShapeRenderer func(c *Context, f CairoShapeRendererFunc, data T.Gpointer, dnotify O.DestroyNotify)
	CairoContextGetShapeRenderer func(c *Context, data *T.Gpointer) CairoShapeRendererFunc
)

var (
	CairoCreateContext      func(c *C.Cairo) *Context
	CairoCreateLayout       func(c *C.Cairo) *Layout
	CairoUpdateLayout       func(c *C.Cairo, layout *Layout)
	CairoShowGlyphString    func(c *C.Cairo, font *Font, glyphs *GlyphString)
	CairoShowGlyphItem      func(c *C.Cairo, text string, glyphItem *GlyphItem)
	CairoShowLayoutLine     func(c *C.Cairo, line *LayoutLine)
	CairoShowLayout         func(c *C.Cairo, layout *Layout)
	CairoShowErrorUnderline func(c *C.Cairo, x, y, width, height float64)
	CairoGlyphStringPath    func(c *C.Cairo, font *Font, glyphs *GlyphString)
	CairoLayoutLinePath     func(c *C.Cairo, line *LayoutLine)
	CairoLayoutPath         func(c *C.Cairo, layout *Layout)
	CairoErrorUnderlinePath func(c *C.Cairo, x, y, width, height float64)
	CairoUpdateContext      func(c *C.Cairo, context *Context)
)

type CairoShapeRendererFunc func(cr *C.Cairo,
	attr *AttrShape, doPath bool, data T.Gpointer)

type Color struct {
	Red   uint16
	Green uint16
	Blue  uint16
}

var (
	ColorGetType func() O.Type

	ColorCopy     func(src *Color) *Color
	ColorFree     func(c *Color)
	ColorParse    func(c *Color, spec string) bool
	ColorToString func(c *Color) string
)

type Context struct{}

var (
	ContextGetType func() O.Type
	ContextNew     func() *Context

	ContextGetBaseDir         func(c *Context) Direction
	ContextGetBaseGravity     func(c *Context) Gravity
	ContextGetFontDescription func(c *Context) *FontDescription
	ContextGetFontMap         func(c *Context) *FontMap
	ContextGetGravity         func(c *Context) Gravity
	ContextGetGravityHint     func(c *Context) GravityHint
	ContextGetLanguage        func(c *Context) *Language
	ContextGetMatrix          func(c *Context) *Matrix
	ContextGetMetrics         func(c *Context, desc *FontDescription, language *Language) *FontMetrics
	ContextListFamilies       func(c *Context, families ***FontFamily, nFamilies *int)
	ContextLoadFont           func(c *Context, desc *FontDescription) *Font
	ContextLoadFontset        func(c *Context, desc *FontDescription, language *Language) *Fontset
	ContextSetBaseDir         func(c *Context, direction Direction)
	ContextSetBaseGravity     func(c *Context, gravity Gravity)
	ContextSetFontDescription func(c *Context, desc *FontDescription)
	ContextSetFontMap         func(c *Context, fontMap *FontMap)
	ContextSetGravityHint     func(c *Context, hint GravityHint)
	ContextSetLanguage        func(c *Context, language *Language)
	ContextSetMatrix          func(c *Context, matrix *Matrix)
)

var ConfigKeyGet func(key string) string

type Coverage struct{}

var (
	CoverageNew       func() *Coverage
	CoverageFromBytes func(bytes *T.Guchar, nBytes int) *Coverage

	CoverageCopy    func(c *Coverage) *Coverage
	CoverageGet     func(c *Coverage, index int) CoverageLevel
	CoverageMax     func(c *Coverage, other *Coverage)
	CoverageRef     func(c *Coverage) *Coverage
	CoverageSet     func(c *Coverage, index int, level CoverageLevel)
	CoverageToBytes func(c *Coverage, bytes **T.Guchar, nBytes *int)
	CoverageUnref   func(c *Coverage)
)

type CoverageLevel Enum

const (
	COVERAGE_NONE CoverageLevel = iota
	COVERAGE_FALLBACK
	COVERAGE_APPROXIMATE
	COVERAGE_EXACT
)

var CoverageLevelGetType func() O.Type
