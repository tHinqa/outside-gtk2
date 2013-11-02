// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	F "github.com/tHinqa/outside-gtk2/fontconfig"
	FT "github.com/tHinqa/outside-gtk2/freetype2"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type FcDecoder struct{ parent O.Object }

var (
	FcDecoderGetType func() O.Type

	FcDecoderGetCharset func(f *FcDecoder, fcfont *FcFont) *F.CharSet
	FcDecoderGetGlyph   func(f *FcDecoder, fcfont *FcFont, wc T.GUint32) Glyph
)

func (f *FcDecoder) GetCharset(fcfont *FcFont) *F.CharSet { return FcDecoderGetCharset(f, fcfont) }
func (f *FcDecoder) GetGlyph(fcfont *FcFont, wc T.GUint32) Glyph {
	return FcDecoderGetGlyph(f, fcfont, wc)
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
	MetricsByLang *L.SList
	Bits          uint
	// IsHinted : 1
	// IsTransformed : 1
}

var (
	FcFontGetType                func() O.Type
	FcFontDescriptionFromPattern func(pattern *F.Pattern, includeSize bool) *FontDescription

	FcFontGetGlyph        func(f *FcFont, wc L.Unichar) uint
	FcFontGetUnknownGlyph func(f *FcFont, wc L.Unichar) Glyph
	FcFontHasChar         func(f *FcFont, wc L.Unichar) bool
	FcFontKernGlyphs      func(f *FcFont, glyphs *GlyphString)
	FcFontLockFace        func(f *FcFont) FT.Face
	FcFontUnlockFace      func(f *FcFont)
)

func (f *FcFont) GetGlyph(wc L.Unichar) uint         { return FcFontGetGlyph(f, wc) }
func (f *FcFont) GetUnknownGlyph(wc L.Unichar) Glyph { return FcFontGetUnknownGlyph(f, wc) }
func (f *FcFont) HasChar(wc L.Unichar) bool          { return FcFontHasChar(f, wc) }
func (f *FcFont) KernGlyphs(glyphs *GlyphString)      { FcFontKernGlyphs(f, glyphs) }
func (f *FcFont) LockFace() FT.Face                   { return FcFontLockFace(f) }
func (f *FcFont) UnlockFace()                         { FcFontUnlockFace(f) }

type FcFontKey struct{}

var (
	FcFontKeyGetContextKey func(f *FcFontKey) T.Gpointer
	FcFontKeyGetMatrix     func(f *FcFontKey) *Matrix
	FcFontKeyGetPattern    func(f *FcFontKey) *F.Pattern
)

func (f *FcFontKey) GetContextKey() T.Gpointer { return FcFontKeyGetContextKey(f) }
func (f *FcFontKey) GetMatrix() *Matrix        { return FcFontKeyGetMatrix(f) }
func (f *FcFontKey) GetPattern() *F.Pattern    { return FcFontKeyGetPattern(f) }

type FcFontMap struct {
	Parent FontMap
	_      *struct{}
}

var (
	FcFontMapGetType func() O.Type

	FcFontMapAddDecoderFindFunc func(f *FcFontMap, findfunc FcDecoderFindFunc, userData T.Gpointer, dnotify O.DestroyNotify)
	FcFontMapCacheClear         func(f *FcFontMap)
	FcFontMapCreateContext      func(f *FcFontMap) *Context
	FcFontMapFindDecoder        func(f *FcFontMap, pattern *F.Pattern) *FcDecoder
	FcFontMapShutdown           func(f *FcFontMap)
)

func (f *FcFontMap) AddDecoderFindFunc(findfunc FcDecoderFindFunc, userData T.Gpointer, dnotify O.DestroyNotify) {
	FcFontMapAddDecoderFindFunc(f, findfunc, userData, dnotify)
}
func (f *FcFontMap) CacheClear()             { FcFontMapCacheClear(f) }
func (f *FcFontMap) CreateContext() *Context { return FcFontMapCreateContext(f) }
func (f *FcFontMap) FindDecoder(pattern *F.Pattern) *FcDecoder {
	return FcFontMapFindDecoder(f, pattern)
}
func (f *FcFontMap) Shutdown() { FcFontMapShutdown(f) }

type FcFontsetKey struct{}

var (
	FcFontsetKeyGetAbsoluteSize func(f *FcFontsetKey) float64
	FcFontsetKeyGetContextKey   func(f *FcFontsetKey) T.Gpointer
	FcFontsetKeyGetDescription  func(f *FcFontsetKey) *FontDescription
	FcFontsetKeyGetLanguage     func(f *FcFontsetKey) *Language
	FcFontsetKeyGetMatrix       func(f *FcFontsetKey) *Matrix
	FcFontsetKeyGetResolution   func(f *FcFontsetKey) float64
)

func (f *FcFontsetKey) GetAbsoluteSize() float64         { return FcFontsetKeyGetAbsoluteSize(f) }
func (f *FcFontsetKey) GetContextKey() T.Gpointer        { return FcFontsetKeyGetContextKey(f) }
func (f *FcFontsetKey) GetDescription() *FontDescription { return FcFontsetKeyGetDescription(f) }
func (f *FcFontsetKey) GetLanguage() *Language           { return FcFontsetKeyGetLanguage(f) }
func (f *FcFontsetKey) GetMatrix() *Matrix               { return FcFontsetKeyGetMatrix(f) }
func (f *FcFontsetKey) GetResolution() float64           { return FcFontsetKeyGetResolution(f) }

var (
	FindBaseDir           func(text string, length int) Direction
	FindMap               func(language *Language, engineTypeId uint, renderTypeId uint) *Map
	FindParagraphBoundary func(text string, length int, paragraphDelimiterIndex, nextParagraphStart *int)
)

type Font struct{}

var (
	FontGetType func() O.Type

	FontFindShaper  func(f *Font, language *Language, ch T.GUint32) *EngineShape
	FontGetCoverage func(f *Font, language *Language) *Coverage
	FontGetFontMap  func(f *Font) *FontMap
	FontGetMetrics  func(f *Font, language *Language) *FontMetrics
)

func (f *Font) FindShaper(language *Language, ch T.GUint32) *EngineShape {
	return FontFindShaper(f, language, ch)
}
func (f *Font) GetCoverage(language *Language) *Coverage   { return FontGetCoverage(f, language) }
func (f *Font) GetFontMap() *FontMap                       { return FontGetFontMap(f) }
func (f *Font) GetMetrics(language *Language) *FontMetrics { return FontGetMetrics(f, language) }

type FontDescription struct{}

var (
	FontDescriptionGetType func() O.Type
	FontDescriptionNew     func() *FontDescription

	FontDescribe                 func(font *Font) *FontDescription
	FontDescribeWithAbsoluteSize func(font *Font) *FontDescription
	FontDescriptionFromString    func(str string) *FontDescription
	FontDescriptionsFree         func(descs **FontDescription, nDescs int)
	FontGetGlyphExtents          func(font *Font, glyph Glyph, inkRect *Rectangle, logicalRect *Rectangle)

	FontDescriptionBetterMatch       func(f *FontDescription, oldMatch *FontDescription, newMatch *FontDescription) bool
	FontDescriptionCopy              func(f *FontDescription) *FontDescription
	FontDescriptionCopyStatic        func(f *FontDescription) *FontDescription
	FontDescriptionEqual             func(f *FontDescription, desc2 *FontDescription) bool
	FontDescriptionFree              func(f *FontDescription)
	FontDescriptionGetFamily         func(f *FontDescription) string
	FontDescriptionGetGravity        func(f *FontDescription) Gravity
	FontDescriptionGetSetFields      func(f *FontDescription) FontMask
	FontDescriptionGetSize           func(f *FontDescription) int
	FontDescriptionGetSizeIsAbsolute func(f *FontDescription) bool
	FontDescriptionGetStretch        func(f *FontDescription) Stretch
	FontDescriptionGetStyle          func(f *FontDescription) Style
	FontDescriptionGetVariant        func(f *FontDescription) Variant
	FontDescriptionGetWeight         func(f *FontDescription) Weight
	FontDescriptionHash              func(f *FontDescription) uint
	FontDescriptionMerge             func(f *FontDescription, descToMerge *FontDescription, replaceExisting bool)
	FontDescriptionMergeStatic       func(f *FontDescription, descToMerge *FontDescription, replaceExisting bool)
	FontDescriptionSetAbsoluteSize   func(f *FontDescription, size float64)
	FontDescriptionSetFamily         func(f *FontDescription, family string)
	FontDescriptionSetFamilyStatic   func(f *FontDescription, family string)
	FontDescriptionSetGravity        func(f *FontDescription, gravity Gravity)
	FontDescriptionSetSize           func(f *FontDescription, size int)
	FontDescriptionSetStretch        func(f *FontDescription, stretch Stretch)
	FontDescriptionSetStyle          func(f *FontDescription, style Style)
	FontDescriptionSetVariant        func(f *FontDescription, variant Variant)
	FontDescriptionSetWeight         func(f *FontDescription, weight Weight)
	FontDescriptionToFilename        func(f *FontDescription) string
	FontDescriptionToString          func(f *FontDescription) string
	FontDescriptionUnsetFields       func(f *FontDescription, toUnset FontMask)
)

func (f *FontDescription) BetterMatch(oldMatch *FontDescription, newMatch *FontDescription) bool {
	return FontDescriptionBetterMatch(f, oldMatch, newMatch)
}
func (f *FontDescription) Copy() *FontDescription       { return FontDescriptionCopy(f) }
func (f *FontDescription) CopyStatic() *FontDescription { return FontDescriptionCopyStatic(f) }
func (f *FontDescription) Equal(desc2 *FontDescription) bool {
	return FontDescriptionEqual(f, desc2)
}
func (f *FontDescription) Free()                   { FontDescriptionFree(f) }
func (f *FontDescription) GetFamily() string       { return FontDescriptionGetFamily(f) }
func (f *FontDescription) GetGravity() Gravity     { return FontDescriptionGetGravity(f) }
func (f *FontDescription) GetSetFields() FontMask  { return FontDescriptionGetSetFields(f) }
func (f *FontDescription) GetSize() int            { return FontDescriptionGetSize(f) }
func (f *FontDescription) GetSizeIsAbsolute() bool { return FontDescriptionGetSizeIsAbsolute(f) }
func (f *FontDescription) GetStretch() Stretch     { return FontDescriptionGetStretch(f) }
func (f *FontDescription) GetStyle() Style         { return FontDescriptionGetStyle(f) }
func (f *FontDescription) GetVariant() Variant     { return FontDescriptionGetVariant(f) }
func (f *FontDescription) GetWeight() Weight       { return FontDescriptionGetWeight(f) }
func (f *FontDescription) Hash() uint              { return FontDescriptionHash(f) }
func (f *FontDescription) Merge(descToMerge *FontDescription, replaceExisting bool) {
	FontDescriptionMerge(f, descToMerge, replaceExisting)
}
func (f *FontDescription) MergeStatic(descToMerge *FontDescription, replaceExisting bool) {
	FontDescriptionMergeStatic(f, descToMerge, replaceExisting)
}
func (f *FontDescription) SetAbsoluteSize(size float64)  { FontDescriptionSetAbsoluteSize(f, size) }
func (f *FontDescription) SetFamily(family string)       { FontDescriptionSetFamily(f, family) }
func (f *FontDescription) SetFamilyStatic(family string) { FontDescriptionSetFamilyStatic(f, family) }
func (f *FontDescription) SetGravity(gravity Gravity)    { FontDescriptionSetGravity(f, gravity) }
func (f *FontDescription) SetSize(size int)              { FontDescriptionSetSize(f, size) }
func (f *FontDescription) SetStretch(stretch Stretch)    { FontDescriptionSetStretch(f, stretch) }
func (f *FontDescription) SetStyle(style Style)          { FontDescriptionSetStyle(f, style) }
func (f *FontDescription) SetVariant(variant Variant)    { FontDescriptionSetVariant(f, variant) }
func (f *FontDescription) SetWeight(weight Weight)       { FontDescriptionSetWeight(f, weight) }
func (f *FontDescription) ToFilename() string            { return FontDescriptionToFilename(f) }
func (f *FontDescription) ToString() string              { return FontDescriptionToString(f) }
func (f *FontDescription) UnsetFields(toUnset FontMask)  { FontDescriptionUnsetFields(f, toUnset) }

type FontFace struct{}

var (
	FontFaceGetType func() O.Type

	FontFaceDescribe      func(f *FontFace) *FontDescription
	FontFaceGetFaceName   func(f *FontFace) string
	FontFaceIsSynthesized func(f *FontFace) bool
	FontFaceListSizes     func(f *FontFace, sizes **int, nSizes *int)
)

func (f *FontFace) Describe() *FontDescription         { return FontFaceDescribe(f) }
func (f *FontFace) GetFaceName() string                { return FontFaceGetFaceName(f) }
func (f *FontFace) IsSynthesized() bool                { return FontFaceIsSynthesized(f) }
func (f *FontFace) ListSizes(sizes **int, nSizes *int) { FontFaceListSizes(f, sizes, nSizes) }

type FontFamily struct{}

var (
	FontFamilyGetType func() O.Type

	FontFamilyGetName     func(f *FontFamily) string
	FontFamilyIsMonospace func(f *FontFamily) bool
	FontFamilyListFaces   func(f *FontFamily, faces ***FontFace, nFaces *int)
)

func (f *FontFamily) GetName() string   { return FontFamilyGetName(f) }
func (f *FontFamily) IsMonospace() bool { return FontFamilyIsMonospace(f) }
func (f *FontFamily) ListFaces(faces ***FontFace, nFaces *int) {
	FontFamilyListFaces(f, faces, nFaces)
}

type FontMap struct{}

var (
	FontMapGetType func() O.Type

	FontMapCreateContext      func(f *FontMap) *Context
	FontMapGetShapeEngineType func(f *FontMap) string
	FontMapListFamilies       func(f *FontMap, families ***FontFamily, nFamilies *int)
	FontMapLoadFont           func(f *FontMap, context *Context, desc *FontDescription) *Font
	FontMapLoadFontset        func(f *FontMap, context *Context, desc *FontDescription, language *Language) *Fontset
)

func (f *FontMap) CreateContext() *Context    { return FontMapCreateContext(f) }
func (f *FontMap) GetShapeEngineType() string { return FontMapGetShapeEngineType(f) }
func (f *FontMap) ListFamilies(families ***FontFamily, nFamilies *int) {
	FontMapListFamilies(f, families, nFamilies)
}
func (f *FontMap) LoadFont(context *Context, desc *FontDescription) *Font {
	return FontMapLoadFont(f, context, desc)
}
func (f *FontMap) LoadFontset(context *Context, desc *FontDescription, language *Language) *Fontset {
	return FontMapLoadFontset(f, context, desc, language)
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

	FontMetricsGetApproximateCharWidth   func(f *FontMetrics) int
	FontMetricsGetApproximateDigitWidth  func(f *FontMetrics) int
	FontMetricsGetAscent                 func(f *FontMetrics) int
	FontMetricsGetDescent                func(f *FontMetrics) int
	FontMetricsGetStrikethroughPosition  func(f *FontMetrics) int
	FontMetricsGetStrikethroughThickness func(f *FontMetrics) int
	FontMetricsGetUnderlinePosition      func(f *FontMetrics) int
	FontMetricsGetUnderlineThickness     func(f *FontMetrics) int
	FontMetricsRef                       func(f *FontMetrics) *FontMetrics
	FontMetricsUnref                     func(f *FontMetrics)
)

func (f *FontMetrics) GetApproximateCharWidth() int   { return FontMetricsGetApproximateCharWidth(f) }
func (f *FontMetrics) GetApproximateDigitWidth() int  { return FontMetricsGetApproximateDigitWidth(f) }
func (f *FontMetrics) GetAscent() int                 { return FontMetricsGetAscent(f) }
func (f *FontMetrics) GetDescent() int                { return FontMetricsGetDescent(f) }
func (f *FontMetrics) GetStrikethroughPosition() int  { return FontMetricsGetStrikethroughPosition(f) }
func (f *FontMetrics) GetStrikethroughThickness() int { return FontMetricsGetStrikethroughThickness(f) }
func (f *FontMetrics) GetUnderlinePosition() int      { return FontMetricsGetUnderlinePosition(f) }
func (f *FontMetrics) GetUnderlineThickness() int     { return FontMetricsGetUnderlineThickness(f) }
func (f *FontMetrics) Ref() *FontMetrics              { return FontMetricsRef(f) }
func (f *FontMetrics) Unref()                         { FontMetricsUnref(f) }

type Fontset struct{}

var (
	FontsetGetType func() O.Type

	FontsetForeach    func(f *Fontset, fnc FontsetForeachFunc, data T.Gpointer)
	FontsetGetFont    func(f *Fontset, wc uint) *Font
	FontsetGetMetrics func(f *Fontset) *FontMetrics
)

func (f *Fontset) Foreach(fnc FontsetForeachFunc, data T.Gpointer) { FontsetForeach(f, fnc, data) }
func (f *Fontset) GetFont(wc uint) *Font                           { return FontsetGetFont(f, wc) }
func (f *Fontset) GetMetrics() *FontMetrics                        { return FontsetGetMetrics(f) }

type FontsetForeachFunc func(
	Fontset *Fontset, font *Font, data T.Gpointer) bool

type FontsetSimple struct{}

var (
	FontsetSimpleGetType func() O.Type
	FontsetSimpleNew     func(language *Language) *FontsetSimple

	FontsetSimpleAppend func(f *FontsetSimple, font *Font)
	FontsetSimpleSize   func(f *FontsetSimple) int
)

func (f *FontsetSimple) Append(font *Font) { FontsetSimpleAppend(f, font) }
func (f *FontsetSimple) Size() int         { return FontsetSimpleSize(f) }

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

	Ft2FontMapCreateContext        func(f *FT2FontMap) *Context
	Ft2FontMapSetDefaultSubstitute func(f *FT2FontMap, fnc FT2SubstituteFunc, data T.Gpointer, notify O.DestroyNotify)
	Ft2FontMapSetResolution        func(f *FT2FontMap, dpiX, dpiY float64)
	Ft2FontMapSubstituteChanged    func(f *FT2FontMap)
)

func (f *FT2FontMap) CreateContext() *Context { return Ft2FontMapCreateContext(f) }
func (f *FT2FontMap) SetDefaultSubstitute(fnc FT2SubstituteFunc, data T.Gpointer, notify O.DestroyNotify) {
	Ft2FontMapSetDefaultSubstitute(f, fnc, data, notify)
}
func (f *FT2FontMap) SetResolution(dpiX, dpiY float64) {
	Ft2FontMapSetResolution(f, dpiX, dpiY)
}
func (f *FT2FontMap) SubstituteChanged() { Ft2FontMapSubstituteChanged(f) }

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
