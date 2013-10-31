// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
	F "github.com/tHinqa/outside-gtk2/fontconfig"
	FT "github.com/tHinqa/outside-gtk2/freetype2"
	T "github.com/tHinqa/outside-gtk2/types"
)

type FillRule Enum

const (
	FILL_RULE_WINDING FillRule = iota
	FILL_RULE_EVEN_ODD
)

type Filter Enum

const (
	FILTER_FAST Filter = iota
	FILTER_GOOD
	FILTER_BEST
	FILTER_NEAREST
	FILTER_BILINEAR
	FILTER_GAUSSIAN
)

type FontExtentsS struct { //NOTE(t): 'S' added; func name clash
	Ascent      float64
	Descent     float64
	Height      float64
	MaxXAdvance float64
	MaxYAdvance float64
}

type (
	FontFace struct{}

	UserScaledFontInitFunc func(
		scaled_font *ScaledFont,
		cr *Cairo,
		extents *FontExtentsS) Status

	UserScaledFontRenderGlyphFunc func(
		scaled_font *ScaledFont,
		glyph T.UnsignedLong,
		cr *Cairo,
		extents *TextExtents) Status

	UserScaledFontTextToGlyphsFunc func(
		scaled_font *ScaledFont,
		utf8 string,
		utf8Len int,
		glyphs **Glyph,
		numGlyphs *int,
		clusters **TextCluster,
		numClusters *int,
		clusterFlags *TextClusterFlags) Status

	UserScaledFontUnicodeToGlyphFunc func(
		scaledFont *ScaledFont,
		unicode T.UnsignedLong,
		glyphIndex *T.UnsignedLong) Status
)

var (
	ToyFontFaceCreate  func(family string, slant FontSlant, weight FontWeight) *FontFace
	UserFontFaceCreate func() *FontFace

	FontFaceDestroy           func(f *FontFace)
	FontFaceGetReferenceCount func(f *FontFace) uint
	FontFaceGetType           func(f *FontFace) FontType
	FontFaceGetUserData       func(f *FontFace, key *UserDataKey) *T.Void
	FontFaceReference         func(f *FontFace) *FontFace
	FontFaceSetUserData       func(f *FontFace, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	FontFaceStatus            func(f *FontFace) Status
	ToyFontFaceGetFamily      func(f *FontFace) string
	ToyFontFaceGetSlant       func(f *FontFace) FontSlant
	ToyFontFaceGetWeight      func(f *FontFace) FontWeight
)

func (f *FontFace) Destroy()                             { FontFaceDestroy(f) }
func (f *FontFace) GetFamily() string                    { return ToyFontFaceGetFamily(f) }
func (f *FontFace) GetReferenceCount() uint              { return FontFaceGetReferenceCount(f) }
func (f *FontFace) GetSlant() FontSlant                  { return ToyFontFaceGetSlant(f) }
func (f *FontFace) GetType() FontType                    { return FontFaceGetType(f) }
func (f *FontFace) GetUserData(key *UserDataKey) *T.Void { return FontFaceGetUserData(f, key) }
func (f *FontFace) GetWeight() FontWeight                { return ToyFontFaceGetWeight(f) }
func (f *FontFace) Reference() *FontFace                 { return FontFaceReference(f) }
func (f *FontFace) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return FontFaceSetUserData(f, key, userData, destroy)
}
func (f *FontFace) Status() Status { return FontFaceStatus(f) }

var (
	UserFontFaceGetInitFunc           func(f *FontFace) UserScaledFontInitFunc
	UserFontFaceGetRenderGlyphFunc    func(f *FontFace) UserScaledFontRenderGlyphFunc
	UserFontFaceGetTextToGlyphsFunc   func(f *FontFace) UserScaledFontTextToGlyphsFunc
	UserFontFaceGetUnicodeToGlyphFunc func(f *FontFace) UserScaledFontUnicodeToGlyphFunc
	UserFontFaceSetInitFunc           func(f *FontFace, initFunc UserScaledFontInitFunc)
	UserFontFaceSetRenderGlyphFunc    func(f *FontFace, renderGlyphFunc UserScaledFontRenderGlyphFunc)
	UserFontFaceSetTextToGlyphsFunc   func(f *FontFace, textToGlyphsFunc UserScaledFontTextToGlyphsFunc)
	UserFontFaceSetUnicodeToGlyphFunc func(f *FontFace, unicodeToGlyphFunc UserScaledFontUnicodeToGlyphFunc)
)

func (f *FontFace) GetInitFunc() UserScaledFontInitFunc { return UserFontFaceGetInitFunc(f) }
func (f *FontFace) GetRenderGlyphFunc() UserScaledFontRenderGlyphFunc {
	return UserFontFaceGetRenderGlyphFunc(f)
}
func (f *FontFace) GetTextToGlyphsFunc() UserScaledFontTextToGlyphsFunc {
	return UserFontFaceGetTextToGlyphsFunc(f)
}
func (f *FontFace) GetUnicodeToGlyphFunc() UserScaledFontUnicodeToGlyphFunc {
	return UserFontFaceGetUnicodeToGlyphFunc(f)
}
func (f *FontFace) SetInitFunc(initFunc UserScaledFontInitFunc) {
	UserFontFaceSetInitFunc(f, initFunc)
}
func (f *FontFace) SetRenderGlyphFunc(renderGlyphFunc UserScaledFontRenderGlyphFunc) {
	UserFontFaceSetRenderGlyphFunc(f, renderGlyphFunc)
}
func (f *FontFace) SetTextToGlyphsFunc(textToGlyphsFunc UserScaledFontTextToGlyphsFunc) {
	UserFontFaceSetTextToGlyphsFunc(f, textToGlyphsFunc)
}
func (f *FontFace) SetUnicodeToGlyphFunc(unicodeToGlyphFunc UserScaledFontUnicodeToGlyphFunc) {
	UserFontFaceSetUnicodeToGlyphFunc(f, unicodeToGlyphFunc)
}

type FontOptions struct{}

var (
	FontOptionsCreate func() *FontOptions

	FontOptionsCopy             func(f *FontOptions) *FontOptions
	FontOptionsDestroy          func(f *FontOptions)
	FontOptionsEqual            func(f *FontOptions, other *FontOptions) Bool
	FontOptionsGetAntialias     func(f *FontOptions) Antialias
	FontOptionsGetHintMetrics   func(f *FontOptions) HintMetrics
	FontOptionsGetHintStyle     func(f *FontOptions) HintStyle
	FontOptionsGetSubpixelOrder func(f *FontOptions) SubpixelOrder
	FontOptionsHash             func(f *FontOptions) T.UnsignedLong
	FontOptionsMerge            func(f *FontOptions, other *FontOptions)
	FontOptionsSetAntialias     func(f *FontOptions, antialias Antialias)
	FontOptionsSetHintMetrics   func(f *FontOptions, hintMetrics HintMetrics)
	FontOptionsSetHintStyle     func(f *FontOptions, hintStyle HintStyle)
	FontOptionsSetSubpixelOrder func(f *FontOptions, subpixelOrder SubpixelOrder)
	FontOptionsStatus           func(f *FontOptions) Status
)

func (f *FontOptions) Copy() *FontOptions               { return FontOptionsCopy(f) }
func (f *FontOptions) Destroy()                         { FontOptionsDestroy(f) }
func (f *FontOptions) Equal(other *FontOptions) Bool    { return FontOptionsEqual(f, other) }
func (f *FontOptions) GetAntialias() Antialias          { return FontOptionsGetAntialias(f) }
func (f *FontOptions) GetHintMetrics() HintMetrics      { return FontOptionsGetHintMetrics(f) }
func (f *FontOptions) GetHintStyle() HintStyle          { return FontOptionsGetHintStyle(f) }
func (f *FontOptions) GetSubpixelOrder() SubpixelOrder  { return FontOptionsGetSubpixelOrder(f) }
func (f *FontOptions) Hash() T.UnsignedLong             { return FontOptionsHash(f) }
func (f *FontOptions) Merge(other *FontOptions)         { FontOptionsMerge(f, other) }
func (f *FontOptions) SetAntialias(antialias Antialias) { FontOptionsSetAntialias(f, antialias) }
func (f *FontOptions) SetHintMetrics(hintMetrics HintMetrics) {
	FontOptionsSetHintMetrics(f, hintMetrics)
}
func (f *FontOptions) SetHintStyle(hintStyle HintStyle) { FontOptionsSetHintStyle(f, hintStyle) }
func (f *FontOptions) SetSubpixelOrder(subpixelOrder SubpixelOrder) {
	FontOptionsSetSubpixelOrder(f, subpixelOrder)
}
func (f *FontOptions) Status() Status { return FontOptionsStatus(f) }

type FontSlant Enum

const (
	FONT_SLANT_NORMAL FontSlant = iota
	FONT_SLANT_ITALIC
	FONT_SLANT_OBLIQUE
)

type FontType Enum

const (
	FONT_TYPE_TOY FontType = iota
	FONT_TYPE_FT
	FONT_TYPE_WIN32
	FONT_TYPE_QUARTZ
	FONT_TYPE_USER
)

type FontWeight Enum

const (
	FONT_WEIGHT_NORMAL FontWeight = iota
	FONT_WEIGHT_BOLD
)

type Format Enum

const (
	FORMAT_INVALID Format = iota - 1
	FORMAT_ARGB32
	FORMAT_RGB24
	FORMAT_A8
	FORMAT_A1
	FORMAT_RGB16_565
)

var FormatStrideForWidth func(format Format, width int) int

var (
	FtFontFaceCreateForFtFace  func(face FT.Face, loadFlags int) *FontFace
	FtFontFaceCreateForPattern func(pattern *F.Pattern) *FontFace
	FtFontOptionsSubstitute    func(options *FontOptions, pattern *F.Pattern)
)

type ScaledFont struct{}

var (
	FtScaledFontLockFace   func(s *ScaledFont) FT.Face
	FtScaledFontUnlockFace func(s *ScaledFont)
)

func (s *ScaledFont) LockFace() FT.Face { return FtScaledFontLockFace(s) }
func (s *ScaledFont) UnlockFace()       { FtScaledFontUnlockFace(s) }
