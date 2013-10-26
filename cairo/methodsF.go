// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
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

	fontFaceDestroy           func(f *FontFace)
	fontFaceGetReferenceCount func(f *FontFace) uint
	fontFaceGetType           func(f *FontFace) FontType
	fontFaceGetUserData       func(f *FontFace, key *UserDataKey) *T.Void
	fontFaceReference         func(f *FontFace) *FontFace
	fontFaceSetUserData       func(f *FontFace, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	fontFaceStatus            func(f *FontFace) Status
	toyFontFaceGetFamily      func(f *FontFace) string
	toyFontFaceGetSlant       func(f *FontFace) FontSlant
	toyFontFaceGetWeight      func(f *FontFace) FontWeight
)

func (f *FontFace) Destroy()                             { fontFaceDestroy(f) }
func (f *FontFace) GetFamily() string                    { return toyFontFaceGetFamily(f) }
func (f *FontFace) GetReferenceCount() uint              { return fontFaceGetReferenceCount(f) }
func (f *FontFace) GetSlant() FontSlant                  { return toyFontFaceGetSlant(f) }
func (f *FontFace) GetType() FontType                    { return fontFaceGetType(f) }
func (f *FontFace) GetUserData(key *UserDataKey) *T.Void { return fontFaceGetUserData(f, key) }
func (f *FontFace) GetWeight() FontWeight                { return toyFontFaceGetWeight(f) }
func (f *FontFace) Reference() *FontFace                 { return fontFaceReference(f) }
func (f *FontFace) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return fontFaceSetUserData(f, key, userData, destroy)
}
func (f *FontFace) Status() Status { return fontFaceStatus(f) }

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

	fontOptionsCopy             func(f *FontOptions) *FontOptions
	fontOptionsDestroy          func(f *FontOptions)
	fontOptionsEqual            func(f *FontOptions, other *FontOptions) Bool
	fontOptionsGetAntialias     func(f *FontOptions) Antialias
	fontOptionsGetHintMetrics   func(f *FontOptions) HintMetrics
	fontOptionsGetHintStyle     func(f *FontOptions) HintStyle
	fontOptionsGetSubpixelOrder func(f *FontOptions) SubpixelOrder
	fontOptionsHash             func(f *FontOptions) T.UnsignedLong
	fontOptionsMerge            func(f *FontOptions, other *FontOptions)
	fontOptionsSetAntialias     func(f *FontOptions, antialias Antialias)
	fontOptionsSetHintMetrics   func(f *FontOptions, hintMetrics HintMetrics)
	fontOptionsSetHintStyle     func(f *FontOptions, hintStyle HintStyle)
	fontOptionsSetSubpixelOrder func(f *FontOptions, subpixelOrder SubpixelOrder)
	fontOptionsStatus           func(f *FontOptions) Status
)

func (f *FontOptions) Copy() *FontOptions               { return fontOptionsCopy(f) }
func (f *FontOptions) Destroy()                         { fontOptionsDestroy(f) }
func (f *FontOptions) Equal(other *FontOptions) Bool    { return fontOptionsEqual(f, other) }
func (f *FontOptions) GetAntialias() Antialias          { return fontOptionsGetAntialias(f) }
func (f *FontOptions) GetHintMetrics() HintMetrics      { return fontOptionsGetHintMetrics(f) }
func (f *FontOptions) GetHintStyle() HintStyle          { return fontOptionsGetHintStyle(f) }
func (f *FontOptions) GetSubpixelOrder() SubpixelOrder  { return fontOptionsGetSubpixelOrder(f) }
func (f *FontOptions) Hash() T.UnsignedLong             { return fontOptionsHash(f) }
func (f *FontOptions) Merge(other *FontOptions)         { fontOptionsMerge(f, other) }
func (f *FontOptions) SetAntialias(antialias Antialias) { fontOptionsSetAntialias(f, antialias) }
func (f *FontOptions) SetHintMetrics(hintMetrics HintMetrics) {
	fontOptionsSetHintMetrics(f, hintMetrics)
}
func (f *FontOptions) SetHintStyle(hintStyle HintStyle) { fontOptionsSetHintStyle(f, hintStyle) }
func (f *FontOptions) SetSubpixelOrder(subpixelOrder SubpixelOrder) {
	fontOptionsSetSubpixelOrder(f, subpixelOrder)
}
func (f *FontOptions) Status() Status { return fontOptionsStatus(f) }

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
	FtFontFaceCreateForFtFace  func(face T.FTFace, loadFlags int) *FontFace
	FtFontFaceCreateForPattern func(pattern *T.FcPattern) *FontFace
	FtFontOptionsSubstitute    func(options *FontOptions, pattern *T.FcPattern)
)

type ScaledFont struct{}

var (
	ftScaledFontLockFace   func(s *ScaledFont) T.FTFace
	ftScaledFontUnlockFace func(s *ScaledFont)
)

func (s *ScaledFont) LockFace() T.FTFace { return ftScaledFontLockFace(s) }
func (s *ScaledFont) UnlockFace()        { ftScaledFontUnlockFace(s) }
