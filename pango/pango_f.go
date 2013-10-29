// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	F "github.com/tHinqa/outside-gtk2/fontconfig"
	FT "github.com/tHinqa/outside-gtk2/freetype2"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type FcDecoder struct{ parent O.Object }

var (
	FcDecoderGetType func() O.Type

	fcDecoderGetCharset func(f *FcDecoder, fcfont *FcFont) *F.CharSet
	fcDecoderGetGlyph   func(f *FcDecoder, fcfont *FcFont, wc T.GUint32) Glyph
)

func (f *FcDecoder) GetCharset(fcfont *FcFont) *F.CharSet { return fcDecoderGetCharset(f, fcfont) }
func (f *FcDecoder) GetGlyph(fcfont *FcFont, wc T.GUint32) Glyph {
	return fcDecoderGetGlyph(f, fcfont, wc)
}

type FcDecoderFindFunc func(
	pattern *F.Pattern, userData T.Gpointer) *FcDecoder

type FcFont struct {
	Parent        Font
	FontPattern   *F.Pattern
	Fontmap       *FontMap
	Priv          T.Gpointer
	Matrix        Matrix
	Description   *FontDescription
	MetricsByLang *T.GSList
	Bits          uint
	// IsHinted : 1
	// IsTransformed : 1
}

var (
	FcFontGetType                func() O.Type
	FcFontDescriptionFromPattern func(pattern *F.Pattern, includeSize T.Gboolean) *FontDescription

	fcFontGetGlyph        func(f *FcFont, wc T.Gunichar) uint
	fcFontGetUnknownGlyph func(f *FcFont, wc T.Gunichar) Glyph
	fcFontHasChar         func(f *FcFont, wc T.Gunichar) T.Gboolean
	fcFontKernGlyphs      func(f *FcFont, glyphs *GlyphString)
	fcFontLockFace        func(f *FcFont) FT.Face
	fcFontUnlockFace      func(f *FcFont)
)

func (f *FcFont) GetGlyph(wc T.Gunichar) uint         { return fcFontGetGlyph(f, wc) }
func (f *FcFont) GetUnknownGlyph(wc T.Gunichar) Glyph { return fcFontGetUnknownGlyph(f, wc) }
func (f *FcFont) HasChar(wc T.Gunichar) T.Gboolean    { return fcFontHasChar(f, wc) }
func (f *FcFont) KernGlyphs(glyphs *GlyphString)      { fcFontKernGlyphs(f, glyphs) }
func (f *FcFont) LockFace() FT.Face                   { return fcFontLockFace(f) }
func (f *FcFont) UnlockFace()                         { fcFontUnlockFace(f) }

type FcFontKey struct{}

var (
	fcFontKeyGetContextKey func(f *FcFontKey) T.Gpointer
	fcFontKeyGetMatrix     func(f *FcFontKey) *Matrix
	fcFontKeyGetPattern    func(f *FcFontKey) *F.Pattern
)

func (f *FcFontKey) GetContextKey() T.Gpointer { return fcFontKeyGetContextKey(f) }
func (f *FcFontKey) GetMatrix() *Matrix        { return fcFontKeyGetMatrix(f) }
func (f *FcFontKey) GetPattern() *F.Pattern    { return fcFontKeyGetPattern(f) }

type FcFontMap struct {
	Parent FontMap
	_      *struct{}
}

var (
	FcFontMapGetType func() O.Type

	fcFontMapAddDecoderFindFunc func(f *FcFontMap, findfunc FcDecoderFindFunc, userData T.Gpointer, dnotify T.GDestroyNotify)
	fcFontMapCacheClear         func(f *FcFontMap)
	fcFontMapCreateContext      func(f *FcFontMap) *Context
	fcFontMapFindDecoder        func(f *FcFontMap, pattern *F.Pattern) *FcDecoder
	fcFontMapShutdown           func(f *FcFontMap)
)

func (f *FcFontMap) AddDecoderFindFunc(findfunc FcDecoderFindFunc, userData T.Gpointer, dnotify T.GDestroyNotify) {
	fcFontMapAddDecoderFindFunc(f, findfunc, userData, dnotify)
}
func (f *FcFontMap) CacheClear()             { fcFontMapCacheClear(f) }
func (f *FcFontMap) CreateContext() *Context { return fcFontMapCreateContext(f) }
func (f *FcFontMap) FindDecoder(pattern *F.Pattern) *FcDecoder {
	return fcFontMapFindDecoder(f, pattern)
}
func (f *FcFontMap) Shutdown() { fcFontMapShutdown(f) }

type FcFontsetKey struct{}

var (
	fcFontsetKeyGetAbsoluteSize func(f *FcFontsetKey) float64
	fcFontsetKeyGetContextKey   func(f *FcFontsetKey) T.Gpointer
	fcFontsetKeyGetDescription  func(f *FcFontsetKey) *FontDescription
	fcFontsetKeyGetLanguage     func(f *FcFontsetKey) *Language
	fcFontsetKeyGetMatrix       func(f *FcFontsetKey) *Matrix
	fcFontsetKeyGetResolution   func(f *FcFontsetKey) float64
)

func (f *FcFontsetKey) GetAbsoluteSize() float64         { return fcFontsetKeyGetAbsoluteSize(f) }
func (f *FcFontsetKey) GetContextKey() T.Gpointer        { return fcFontsetKeyGetContextKey(f) }
func (f *FcFontsetKey) GetDescription() *FontDescription { return fcFontsetKeyGetDescription(f) }
func (f *FcFontsetKey) GetLanguage() *Language           { return fcFontsetKeyGetLanguage(f) }
func (f *FcFontsetKey) GetMatrix() *Matrix               { return fcFontsetKeyGetMatrix(f) }
func (f *FcFontsetKey) GetResolution() float64           { return fcFontsetKeyGetResolution(f) }

var (
	FindBaseDir           func(text string, length int) Direction
	FindMap               func(language *Language, engineTypeId uint, renderTypeId uint) *Map
	FindParagraphBoundary func(text string, length int, paragraphDelimiterIndex, nextParagraphStart *int)
)

type Font struct{}

var (
	FontGetType func() O.Type

	fontFindShaper  func(f *Font, language *Language, ch T.GUint32) *EngineShape
	fontGetCoverage func(f *Font, language *Language) *Coverage
	fontGetFontMap  func(f *Font) *FontMap
	fontGetMetrics  func(f *Font, language *Language) *FontMetrics
)

func (f *Font) FindShaper(language *Language, ch T.GUint32) *EngineShape {
	return fontFindShaper(f, language, ch)
}
func (f *Font) GetCoverage(language *Language) *Coverage   { return fontGetCoverage(f, language) }
func (f *Font) GetFontMap() *FontMap                       { return fontGetFontMap(f) }
func (f *Font) GetMetrics(language *Language) *FontMetrics { return fontGetMetrics(f, language) }

type FontDescription struct{}

var (
	FontDescriptionGetType func() O.Type
	FontDescriptionNew     func() *FontDescription

	FontDescribe                 func(font *Font) *FontDescription
	FontDescribeWithAbsoluteSize func(font *Font) *FontDescription
	FontDescriptionFromString    func(str string) *FontDescription
	FontDescriptionsFree         func(descs **FontDescription, nDescs int)
	FontGetGlyphExtents          func(font *Font, glyph Glyph, inkRect *Rectangle, logicalRect *Rectangle)

	fontDescriptionBetterMatch       func(f *FontDescription, oldMatch *FontDescription, newMatch *FontDescription) T.Gboolean
	fontDescriptionCopy              func(f *FontDescription) *FontDescription
	fontDescriptionCopyStatic        func(f *FontDescription) *FontDescription
	fontDescriptionEqual             func(f *FontDescription, desc2 *FontDescription) T.Gboolean
	fontDescriptionFree              func(f *FontDescription)
	fontDescriptionGetFamily         func(f *FontDescription) string
	fontDescriptionGetGravity        func(f *FontDescription) Gravity
	fontDescriptionGetSetFields      func(f *FontDescription) FontMask
	fontDescriptionGetSize           func(f *FontDescription) int
	fontDescriptionGetSizeIsAbsolute func(f *FontDescription) T.Gboolean
	fontDescriptionGetStretch        func(f *FontDescription) Stretch
	fontDescriptionGetStyle          func(f *FontDescription) Style
	fontDescriptionGetVariant        func(f *FontDescription) Variant
	fontDescriptionGetWeight         func(f *FontDescription) Weight
	fontDescriptionHash              func(f *FontDescription) uint
	fontDescriptionMerge             func(f *FontDescription, descToMerge *FontDescription, replaceExisting T.Gboolean)
	fontDescriptionMergeStatic       func(f *FontDescription, descToMerge *FontDescription, replaceExisting T.Gboolean)
	fontDescriptionSetAbsoluteSize   func(f *FontDescription, size float64)
	fontDescriptionSetFamily         func(f *FontDescription, family string)
	fontDescriptionSetFamilyStatic   func(f *FontDescription, family string)
	fontDescriptionSetGravity        func(f *FontDescription, gravity Gravity)
	fontDescriptionSetSize           func(f *FontDescription, size int)
	fontDescriptionSetStretch        func(f *FontDescription, stretch Stretch)
	fontDescriptionSetStyle          func(f *FontDescription, style Style)
	fontDescriptionSetVariant        func(f *FontDescription, variant Variant)
	fontDescriptionSetWeight         func(f *FontDescription, weight Weight)
	fontDescriptionToFilename        func(f *FontDescription) string
	fontDescriptionToString          func(f *FontDescription) string
	fontDescriptionUnsetFields       func(f *FontDescription, toUnset FontMask)
)

func (f *FontDescription) BetterMatch(oldMatch *FontDescription, newMatch *FontDescription) T.Gboolean {
	return fontDescriptionBetterMatch(f, oldMatch, newMatch)
}
func (f *FontDescription) Copy() *FontDescription       { return fontDescriptionCopy(f) }
func (f *FontDescription) CopyStatic() *FontDescription { return fontDescriptionCopyStatic(f) }
func (f *FontDescription) Equal(desc2 *FontDescription) T.Gboolean {
	return fontDescriptionEqual(f, desc2)
}
func (f *FontDescription) Free()                         { fontDescriptionFree(f) }
func (f *FontDescription) GetFamily() string             { return fontDescriptionGetFamily(f) }
func (f *FontDescription) GetGravity() Gravity           { return fontDescriptionGetGravity(f) }
func (f *FontDescription) GetSetFields() FontMask        { return fontDescriptionGetSetFields(f) }
func (f *FontDescription) GetSize() int                  { return fontDescriptionGetSize(f) }
func (f *FontDescription) GetSizeIsAbsolute() T.Gboolean { return fontDescriptionGetSizeIsAbsolute(f) }
func (f *FontDescription) GetStretch() Stretch           { return fontDescriptionGetStretch(f) }
func (f *FontDescription) GetStyle() Style               { return fontDescriptionGetStyle(f) }
func (f *FontDescription) GetVariant() Variant           { return fontDescriptionGetVariant(f) }
func (f *FontDescription) GetWeight() Weight             { return fontDescriptionGetWeight(f) }
func (f *FontDescription) Hash() uint                    { return fontDescriptionHash(f) }
func (f *FontDescription) Merge(descToMerge *FontDescription, replaceExisting T.Gboolean) {
	fontDescriptionMerge(f, descToMerge, replaceExisting)
}
func (f *FontDescription) MergeStatic(descToMerge *FontDescription, replaceExisting T.Gboolean) {
	fontDescriptionMergeStatic(f, descToMerge, replaceExisting)
}
func (f *FontDescription) SetAbsoluteSize(size float64)  { fontDescriptionSetAbsoluteSize(f, size) }
func (f *FontDescription) SetFamily(family string)       { fontDescriptionSetFamily(f, family) }
func (f *FontDescription) SetFamilyStatic(family string) { fontDescriptionSetFamilyStatic(f, family) }
func (f *FontDescription) SetGravity(gravity Gravity)    { fontDescriptionSetGravity(f, gravity) }
func (f *FontDescription) SetSize(size int)              { fontDescriptionSetSize(f, size) }
func (f *FontDescription) SetStretch(stretch Stretch)    { fontDescriptionSetStretch(f, stretch) }
func (f *FontDescription) SetStyle(style Style)          { fontDescriptionSetStyle(f, style) }
func (f *FontDescription) SetVariant(variant Variant)    { fontDescriptionSetVariant(f, variant) }
func (f *FontDescription) SetWeight(weight Weight)       { fontDescriptionSetWeight(f, weight) }
func (f *FontDescription) ToFilename() string            { return fontDescriptionToFilename(f) }
func (f *FontDescription) ToString() string              { return fontDescriptionToString(f) }
func (f *FontDescription) UnsetFields(toUnset FontMask)  { fontDescriptionUnsetFields(f, toUnset) }

type FontFace struct{}

var (
	FontFaceGetType func() O.Type

	fontFaceDescribe      func(f *FontFace) *FontDescription
	fontFaceGetFaceName   func(f *FontFace) string
	fontFaceIsSynthesized func(f *FontFace) T.Gboolean
	fontFaceListSizes     func(f *FontFace, sizes **int, nSizes *int)
)

func (f *FontFace) Describe() *FontDescription         { return fontFaceDescribe(f) }
func (f *FontFace) GetFaceName() string                { return fontFaceGetFaceName(f) }
func (f *FontFace) IsSynthesized() T.Gboolean          { return fontFaceIsSynthesized(f) }
func (f *FontFace) ListSizes(sizes **int, nSizes *int) { fontFaceListSizes(f, sizes, nSizes) }

type FontFamily struct{}

var (
	FontFamilyGetType func() O.Type

	fontFamilyGetName     func(f *FontFamily) string
	fontFamilyIsMonospace func(f *FontFamily) T.Gboolean
	fontFamilyListFaces   func(f *FontFamily, faces ***FontFace, nFaces *int)
)

func (f *FontFamily) GetName() string         { return fontFamilyGetName(f) }
func (f *FontFamily) IsMonospace() T.Gboolean { return fontFamilyIsMonospace(f) }
func (f *FontFamily) ListFaces(faces ***FontFace, nFaces *int) {
	fontFamilyListFaces(f, faces, nFaces)
}

type FontMap struct{}

var (
	FontMapGetType func() O.Type

	fontMapCreateContext      func(f *FontMap) *Context
	fontMapGetShapeEngineType func(f *FontMap) string
	fontMapListFamilies       func(f *FontMap, families ***FontFamily, nFamilies *int)
	fontMapLoadFont           func(f *FontMap, context *Context, desc *FontDescription) *Font
	fontMapLoadFontset        func(f *FontMap, context *Context, desc *FontDescription, language *Language) *Fontset
)

func (f *FontMap) CreateContext() *Context    { return fontMapCreateContext(f) }
func (f *FontMap) GetShapeEngineType() string { return fontMapGetShapeEngineType(f) }
func (f *FontMap) ListFamilies(families ***FontFamily, nFamilies *int) {
	fontMapListFamilies(f, families, nFamilies)
}
func (f *FontMap) LoadFont(context *Context, desc *FontDescription) *Font {
	return fontMapLoadFont(f, context, desc)
}
func (f *FontMap) LoadFontset(context *Context, desc *FontDescription, language *Language) *Fontset {
	return fontMapLoadFontset(f, context, desc, language)
}

type FontMask Enum

const (
	FONT_MASK_FAMILY FontMask = 1 << iota
	FONT_MASK_STYLE
	FONT_MASK_VARIANT
	FONT_MASK_WEIGHT
	FONT_MASK_STRETCH
	FONT_MASK_SIZE
	FONT_MASK_GRAVITY
)

var FontMaskGetType func() O.Type

type FontMetrics struct{}

var (
	FontMetricsGetType func() O.Type
	FontMetricsNew     func() *FontMetrics

	fontMetricsGetApproximateCharWidth   func(f *FontMetrics) int
	fontMetricsGetApproximateDigitWidth  func(f *FontMetrics) int
	fontMetricsGetAscent                 func(f *FontMetrics) int
	fontMetricsGetDescent                func(f *FontMetrics) int
	fontMetricsGetStrikethroughPosition  func(f *FontMetrics) int
	fontMetricsGetStrikethroughThickness func(f *FontMetrics) int
	fontMetricsGetUnderlinePosition      func(f *FontMetrics) int
	fontMetricsGetUnderlineThickness     func(f *FontMetrics) int
	fontMetricsRef                       func(f *FontMetrics) *FontMetrics
	fontMetricsUnref                     func(f *FontMetrics)
)

func (f *FontMetrics) GetApproximateCharWidth() int   { return fontMetricsGetApproximateCharWidth(f) }
func (f *FontMetrics) GetApproximateDigitWidth() int  { return fontMetricsGetApproximateDigitWidth(f) }
func (f *FontMetrics) GetAscent() int                 { return fontMetricsGetAscent(f) }
func (f *FontMetrics) GetDescent() int                { return fontMetricsGetDescent(f) }
func (f *FontMetrics) GetStrikethroughPosition() int  { return fontMetricsGetStrikethroughPosition(f) }
func (f *FontMetrics) GetStrikethroughThickness() int { return fontMetricsGetStrikethroughThickness(f) }
func (f *FontMetrics) GetUnderlinePosition() int      { return fontMetricsGetUnderlinePosition(f) }
func (f *FontMetrics) GetUnderlineThickness() int     { return fontMetricsGetUnderlineThickness(f) }
func (f *FontMetrics) Ref() *FontMetrics              { return fontMetricsRef(f) }
func (f *FontMetrics) Unref()                         { fontMetricsUnref(f) }

type Fontset struct{}

var (
	FontsetGetType func() O.Type

	fontsetForeach    func(f *Fontset, fnc FontsetForeachFunc, data T.Gpointer)
	fontsetGetFont    func(f *Fontset, wc uint) *Font
	fontsetGetMetrics func(f *Fontset) *FontMetrics
)

func (f *Fontset) Foreach(fnc FontsetForeachFunc, data T.Gpointer) { fontsetForeach(f, fnc, data) }
func (f *Fontset) GetFont(wc uint) *Font                           { return fontsetGetFont(f, wc) }
func (f *Fontset) GetMetrics() *FontMetrics                        { return fontsetGetMetrics(f) }

type FontsetForeachFunc func(
	fontset *Fontset, font *Font, data T.Gpointer) T.Gboolean

type FontsetSimple struct{}

var (
	FontsetSimpleGetType func() O.Type
	FontsetSimpleNew     func(language *Language) *FontsetSimple

	fontsetSimpleAppend func(f *FontsetSimple, font *Font)
	fontsetSimpleSize   func(f *FontsetSimple) int
)

func (f *FontsetSimple) Append(font *Font) { fontsetSimpleAppend(f, font) }
func (f *FontsetSimple) Size() int         { return fontsetSimpleSize(f) }

var (
	Ft2FontGetCoverage func(font *Font, language *Language) *Coverage
	Ft2FontGetFace     func(font *Font) FT.Face
	Ft2FontGetKerning  func(font *Font, left Glyph, right Glyph) int
)

type FT2FontMap struct{}

var (
	Ft2FontMapGetType func() O.Type
	Ft2FontMapNew     func() *FontMap

	Ft2FontMapForDisplay func() *FontMap

	ft2FontMapCreateContext        func(f *FT2FontMap) *Context
	ft2FontMapSetDefaultSubstitute func(f *FT2FontMap, fnc FT2SubstituteFunc, data T.Gpointer, notify T.GDestroyNotify)
	ft2FontMapSetResolution        func(f *FT2FontMap, dpiX, dpiY float64)
	ft2FontMapSubstituteChanged    func(f *FT2FontMap)
)

func (f *FT2FontMap) CreateContext() *Context { return ft2FontMapCreateContext(f) }
func (f *FT2FontMap) SetDefaultSubstitute(fnc FT2SubstituteFunc, data T.Gpointer, notify T.GDestroyNotify) {
	ft2FontMapSetDefaultSubstitute(f, fnc, data, notify)
}
func (f *FT2FontMap) SetResolution(dpiX, dpiY float64) {
	ft2FontMapSetResolution(f, dpiX, dpiY)
}
func (f *FT2FontMap) SubstituteChanged() { ft2FontMapSubstituteChanged(f) }

var (
	Ft2GetContext      func(dpiX, dpiY float64) *Context
	Ft2GetUnknownGlyph func(font *Font) Glyph
	Ft2ShutdownDisplay func()
)
var (
	Ft2Render                   func(bitmap *FT.Bitmap, font *Font, glyphs *GlyphString, x, y int)
	Ft2RenderLayout             func(bitmap *FT.Bitmap, layout *Layout, x int, y int)
	Ft2RenderLayoutLine         func(bitmap *FT.Bitmap, line *LayoutLine, x int, y int)
	Ft2RenderLayoutLineSubpixel func(bitmap *FT.Bitmap, line *LayoutLine, x int, y int)
	Ft2RenderLayoutSubpixel     func(bitmap *FT.Bitmap, layout *Layout, x int, y int)
	Ft2RenderTransformed        func(bitmap *FT.Bitmap, matrix *Matrix, font *Font, glyphs *GlyphString, x, y int)
)

type FT2SubstituteFunc func(pattern *F.Pattern, data T.Gpointer)
