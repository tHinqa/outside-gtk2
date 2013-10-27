// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

var ReadLine func(stream *T.FILE, str *T.GString) int

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

	rendererActivate           func(renderer *Renderer)
	rendererDeactivate         func(renderer *Renderer)
	rendererDrawErrorUnderline func(renderer *Renderer, x, y, width, height int)
	rendererDrawGlyph          func(renderer *Renderer, font *Font, glyph Glyph, x, y float64)
	rendererDrawGlyphItem      func(renderer *Renderer, text string, glyphItem *GlyphItem, x, y int)
	rendererDrawGlyphs         func(renderer *Renderer, font *Font, glyphs *GlyphString, x, y int)
	rendererDrawLayout         func(renderer *Renderer, layout *Layout, x, y int)
	rendererDrawLayoutLine     func(renderer *Renderer, line *LayoutLine, x, y int)
	rendererDrawRectangle      func(renderer *Renderer, part RenderPart, x, y, width, height int)
	rendererDrawTrapezoid      func(renderer *Renderer, part RenderPart, y1, x11, x21, y2, x12, x22 float64)
	rendererGetColor           func(renderer *Renderer, part RenderPart) *Color
	rendererGetLayout          func(renderer *Renderer) *Layout
	rendererGetLayoutLine      func(renderer *Renderer) *LayoutLine
	rendererGetMatrix          func(renderer *Renderer) *Matrix
	rendererPartChanged        func(renderer *Renderer, part RenderPart)
	rendererSetColor           func(renderer *Renderer, part RenderPart, color *Color)
	rendererSetMatrix          func(renderer *Renderer, matrix *Matrix)
)

func (r *Renderer) Activate()   { rendererActivate(r) }
func (r *Renderer) Deactivate() { rendererDeactivate(r) }
func (r *Renderer) DrawErrorUnderline(x, y, width, height int) {
	rendererDrawErrorUnderline(r, x, y, width, height)
}
func (r *Renderer) DrawGlyph(font *Font, glyph Glyph, x, y float64) {
	rendererDrawGlyph(r, font, glyph, x, y)
}
func (r *Renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x, y int) {
	rendererDrawGlyphItem(r, text, glyphItem, x, y)
}
func (r *Renderer) DrawGlyphs(font *Font, glyphs *GlyphString, x, y int) {
	rendererDrawGlyphs(r, font, glyphs, x, y)
}
func (r *Renderer) DrawLayout(layout *Layout, x, y int)       { rendererDrawLayout(r, layout, x, y) }
func (r *Renderer) DrawLayoutLine(line *LayoutLine, x, y int) { rendererDrawLayoutLine(r, line, x, y) }
func (r *Renderer) DrawRectangle(part RenderPart, x, y, width, height int) {
	rendererDrawRectangle(r, part, x, y, width, height)
}
func (r *Renderer) DrawTrapezoid(part RenderPart, y1, x11, x21, y2, x12, x22 float64) {
	rendererDrawTrapezoid(r, part, y1, x11, x21, y2, x12, x22)
}
func (r *Renderer) GetColor(part RenderPart) *Color        { return rendererGetColor(r, part) }
func (r *Renderer) GetLayout() *Layout                     { return rendererGetLayout(r) }
func (r *Renderer) GetLayoutLine() *LayoutLine             { return rendererGetLayoutLine(r) }
func (r *Renderer) GetMatrix() *Matrix                     { return rendererGetMatrix(r) }
func (r *Renderer) PartChanged(part RenderPart)            { rendererPartChanged(r, part) }
func (r *Renderer) SetColor(part RenderPart, color *Color) { rendererSetColor(r, part, color) }
func (r *Renderer) SetMatrix(matrix *Matrix)               { rendererSetMatrix(r, matrix) }

type RenderPart Enum

const (
	RENDER_PART_FOREGROUND RenderPart = iota
	RENDER_PART_BACKGROUND
	RENDER_PART_UNDERLINE
	RENDER_PART_STRIKETHROUGH
)

var RenderPartGetType func() O.Type

var ReorderItems func(logicalItems *T.GList) *T.GList
