// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

var ReadLine func(stream *T.FILE, str *L.String) int

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Renderer struct {
	Parent        O.Object
	Underline     Underline
	Strikethrough T.Gboolean
	ActiveCount   int
	Matrix        *Matrix
	_             *struct{}
}

var (
	RendererGetType func() O.Type

	RendererActivate           func(renderer *Renderer)
	RendererDeactivate         func(renderer *Renderer)
	RendererDrawErrorUnderline func(renderer *Renderer, x, y, width, height int)
	RendererDrawGlyph          func(renderer *Renderer, font *Font, glyph Glyph, x, y float64)
	RendererDrawGlyphItem      func(renderer *Renderer, text string, glyphItem *GlyphItem, x, y int)
	RendererDrawGlyphs         func(renderer *Renderer, font *Font, glyphs *GlyphString, x, y int)
	RendererDrawLayout         func(renderer *Renderer, layout *Layout, x, y int)
	RendererDrawLayoutLine     func(renderer *Renderer, line *LayoutLine, x, y int)
	RendererDrawRectangle      func(renderer *Renderer, part RenderPart, x, y, width, height int)
	RendererDrawTrapezoid      func(renderer *Renderer, part RenderPart, y1, x11, x21, y2, x12, x22 float64)
	RendererGetColor           func(renderer *Renderer, part RenderPart) *Color
	RendererGetLayout          func(renderer *Renderer) *Layout
	RendererGetLayoutLine      func(renderer *Renderer) *LayoutLine
	RendererGetMatrix          func(renderer *Renderer) *Matrix
	RendererPartChanged        func(renderer *Renderer, part RenderPart)
	RendererSetColor           func(renderer *Renderer, part RenderPart, color *Color)
	RendererSetMatrix          func(renderer *Renderer, matrix *Matrix)
)

func (r *Renderer) Activate()   { RendererActivate(r) }
func (r *Renderer) Deactivate() { RendererDeactivate(r) }
func (r *Renderer) DrawErrorUnderline(x, y, width, height int) {
	RendererDrawErrorUnderline(r, x, y, width, height)
}
func (r *Renderer) DrawGlyph(font *Font, glyph Glyph, x, y float64) {
	RendererDrawGlyph(r, font, glyph, x, y)
}
func (r *Renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x, y int) {
	RendererDrawGlyphItem(r, text, glyphItem, x, y)
}
func (r *Renderer) DrawGlyphs(font *Font, glyphs *GlyphString, x, y int) {
	RendererDrawGlyphs(r, font, glyphs, x, y)
}
func (r *Renderer) DrawLayout(layout *Layout, x, y int)       { RendererDrawLayout(r, layout, x, y) }
func (r *Renderer) DrawLayoutLine(line *LayoutLine, x, y int) { RendererDrawLayoutLine(r, line, x, y) }
func (r *Renderer) DrawRectangle(part RenderPart, x, y, width, height int) {
	RendererDrawRectangle(r, part, x, y, width, height)
}
func (r *Renderer) DrawTrapezoid(part RenderPart, y1, x11, x21, y2, x12, x22 float64) {
	RendererDrawTrapezoid(r, part, y1, x11, x21, y2, x12, x22)
}
func (r *Renderer) GetColor(part RenderPart) *Color        { return RendererGetColor(r, part) }
func (r *Renderer) GetLayout() *Layout                     { return RendererGetLayout(r) }
func (r *Renderer) GetLayoutLine() *LayoutLine             { return RendererGetLayoutLine(r) }
func (r *Renderer) GetMatrix() *Matrix                     { return RendererGetMatrix(r) }
func (r *Renderer) PartChanged(part RenderPart)            { RendererPartChanged(r, part) }
func (r *Renderer) SetColor(part RenderPart, color *Color) { RendererSetColor(r, part, color) }
func (r *Renderer) SetMatrix(matrix *Matrix)               { RendererSetMatrix(r, matrix) }

type RenderPart Enum

const (
	RENDER_PART_FOREGROUND RenderPart = iota
	RENDER_PART_BACKGROUND
	RENDER_PART_UNDERLINE
	RENDER_PART_STRIKETHROUGH
)

var RenderPartGetType func() O.Type

var ReorderItems func(logicalItems *L.List) *L.List
