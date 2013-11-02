// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package freetype provides API definitions for accessing
//freetype6.dll.
package freetype

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

const (
	VALIDATE_lcar_INDEX    = 9
	VALIDATE_GX_LAST_INDEX = VALIDATE_lcar_INDEX
	VALIDATE_GX_LENGTH     = VALIDATE_GX_LAST_INDEX + 1
)

type (
	Enum int

	Angle   Fixed
	Bool    uint8
	Char    int8
	Error   int
	F26Dot6 Long
	Fixed   Long
	Int32   int   // ANOMALLY int != int32
	Long    int32 //TODO(t): signed long SIZE ?
	Pos     Long
	String  int8
	UInt32  uint   // ANOMALLY unsigned int != uint32
	ULong   uint32 //unsigned long SIZE ?

	Renderer   struct{}
	Void       struct{}
	ImageCache struct{}

	SpanFunc func(y, count int, spans *Span, user *Void)

	RasterBitTest func(y, x int, user *Void) int

	RasterBitSet func(y, x int, user *Void)

	DebugHookFunc func(arg *Void)

	FaceRequester func(faceId FaceID, library **Library,
		requestData *Void, aface **Face) Error
)

var (
	MulDiv func(a, b, c Long) Long

	DivFix func(a, b Long) Long

	RoundFix func(a Fixed) Fixed

	CeilFix func(a Fixed) Fixed

	FloorFix func(a Fixed) Fixed

	GlyphToBitmap func(
		theGlyph **Glyph,
		renderMode RenderMode,
		origin *Vector,
		destroy Bool) Error

	GlyphStroke func(
		pglyph **Glyph,
		stroker *Stroker,
		destroy Bool) Error

	GlyphStrokeBorder func(
		pglyph **Glyph,
		stroker *Stroker,
		inside, destroy Bool) Error

	Sin func(angle Angle) Fixed

	Cos func(angle Angle) Fixed

	Tan func(angle Angle) Fixed

	Atan2 func(x, y Fixed) Angle

	AngleDiff func(angle1, angle2 Angle) Angle

	MulFix func(a, b Long) Long
)

type BBox struct {
	XMin, YMin Pos
	XMax, YMax Pos
}

type BDFProperty struct {
	Type BDFPropertyType
	// Union
	Atom *Char
	// Integer  FT_Int32
	// Cardinal FT_UInt32
}

type BDFPropertyType Enum

const (
	BDF_PROPERTY_TYPE_NONE BDFPropertyType = iota
	BDF_PROPERTY_TYPE_ATOM
	BDF_PROPERTY_TYPE_INTEGER
	BDF_PROPERTY_TYPE_CARDINAL
)

type Bitmap struct {
	Rows        int
	Width       int
	Pitch       int
	Buffer      *uint8
	NumGrays    int16
	PixelMode   int8
	PaletteMode int8
	Palette     *Void
}

var (
	BitmapNew func(abitmap **Bitmap)

	BitmapConvert  func(library *Library, source, target *Bitmap, alignment int) Error
	BitmapCopy     func(library *Library, source, target *Bitmap) Error
	BitmapDone     func(library *Library, bitmap *Bitmap) Error
	BitmapEmbolden func(library *Library, bitmap *Bitmap, xStrength, yStrength Pos) Error
)

func (b *Bitmap) Convert(target *Bitmap, library *Library, alignment int) Error {
	return BitmapConvert(library, b, target, alignment)
}
func (b *Bitmap) Copy(target *Bitmap, library *Library) Error { return BitmapCopy(library, b, target) }
func (b *Bitmap) Done(bitmap *Bitmap, library *Library) Error { return BitmapDone(library, b) }
func (b *Bitmap) Embolden(bitmap *Bitmap, library *Library, xStrength, yStrength Pos) Error {
	return BitmapEmbolden(library, b, xStrength, yStrength)
}

type BitmapSize struct {
	Height int16
	Width  int16
	Size   Pos
	XPpem  Pos
	YPpem  Pos
}

type (
	CharMap struct {
		Face       *Face
		Encoding   Encoding
		PlatformId uint16
		EncodingId uint16
	}
	CMapCache struct{}
)

var (
	GetCMapLanguageID func(c *CharMap) ULong
	GetCMapFormat     func(c *CharMap) Long
	GetCharmapIndex   func(c *CharMap) int

	CMapCacheLookup func(cache *CMapCache, faceId FaceID, cmapIndex int, charCode UInt32) uint
)

func (c *CharMap) LanguageID() ULong { return GetCMapLanguageID(c) }
func (c *CharMap) Format() Long      { return GetCMapFormat(c) }
func (c *CharMap) Index() int        { return GetCharmapIndex(c) }

func (c *CMapCache) Lookup(faceId FaceID, cmapIndex int, charCode UInt32) uint {
	return CMapCacheLookup(c, faceId, cmapIndex, charCode)
}

type Driver struct{}

type Encoding uint32 // TODO(t): enum/Size?

const (
	ENCODING_NONE           Encoding = 0
	ENCODING_MS_SYMBOL      Encoding = (('s' << 24) | ('y' << 16) | ('m' << 8) | 'b')
	ENCODING_UNICODE        Encoding = (('u' << 24) | ('n' << 16) | ('i' << 8) | 'c')
	ENCODING_SJIS           Encoding = (('s' << 24) | ('j' << 16) | ('i' << 8) | 's')
	ENCODING_GB2312         Encoding = (('g' << 24) | ('b' << 16) | (' ' << 8) | ' ')
	ENCODING_BIG5           Encoding = (('b' << 24) | ('i' << 16) | ('g' << 8) | '5')
	ENCODING_WANSUNG        Encoding = (('w' << 24) | ('a' << 16) | ('n' << 8) | 's')
	ENCODING_JOHAB          Encoding = (('j' << 24) | ('o' << 16) | ('h' << 8) | 'a')
	ENCODING_MS_SJIS                 = ENCODING_SJIS
	ENCODING_MS_GB2312               = ENCODING_GB2312
	ENCODING_MS_BIG5                 = ENCODING_BIG5
	ENCODING_MS_WANSUNG              = ENCODING_WANSUNG
	ENCODING_MS_JOHAB                = ENCODING_JOHAB
	ENCODING_ADOBE_STANDARD Encoding = (('A' << 24) | ('D' << 16) | ('O' << 8) | 'B')
	ENCODING_ADOBE_EXPERT   Encoding = (('A' << 24) | ('D' << 16) | ('B' << 8) | 'E')
	ENCODING_ADOBE_CUSTOM   Encoding = (('A' << 24) | ('D' << 16) | ('B' << 8) | 'C')
	ENCODING_ADOBE_LATIN_1  Encoding = (('l' << 24) | ('a' << 16) | ('t' << 8) | '1')
	ENCODING_OLD_LATIN_2    Encoding = (('l' << 24) | ('a' << 16) | ('t' << 8) | '2')
	ENCODING_APPLE_ROMAN    Encoding = (('a' << 24) | ('r' << 16) | ('m' << 8) | 'n')
)

type Face struct {
	NumFaces           Long
	FaceIndex          Long
	FaceFlags          Long
	StyleFlags         Long
	NumGlyphs          Long
	FamilyName         *String
	StyleName          *String
	NumFixedSizes      int
	AvailableSizes     *BitmapSize
	NumCharmaps        int
	Charmaps           **CharMap
	Generic            Generic
	Bbox               BBox
	UnitsPer_EM        uint16
	Ascender           int16
	Descender          int16
	Height             int16
	MaxAdvanceWidth    int16
	MaxAdvanceHeight   int16
	UnderlinePosition  int16
	UnderlineThickness int16
	Glyph              *GlyphSlot
	Size               *Size
	Charmap            *CharMap
	Driver             *Driver
	Memory             *Memory
	Stream             *Stream
	SizesList          List
	Autohint           Generic
	Extensions         *Void
	Internal           *FaceInternal
}

var (
	FaceCheckTrueTypePatents    func(f *Face) Bool
	FaceGetCharsOfVariant       func(f *Face, variantSelector ULong) *UInt32
	FaceGetCharVariantIndex     func(f *Face, charcode, variantSelector ULong) uint
	FaceGetCharVariantIsDefault func(f *Face, charcode, variantSelector ULong) int
	FaceGetVariantSelectors     func(f *Face) *UInt32
	FaceGetVariantsOfChar       func(f *Face, charcode ULong) *UInt32
	FaceSetUnpatentedHinting    func(f *Face, value Bool) Bool

	AttachFile                       func(f *Face, filepathname string) Error
	AttachStream                     func(f *Face, parameters *OpenArgs) Error
	ClassicKernFree                  func(f *Face, table *byte)
	ClassicKernValidate              func(f *Face, validation_flags uint, ckernTable **byte) Error
	DoneFace                         func(f *Face) Error
	GetAdvance                       func(f *Face, gindex uint, loadFlags Int32, padvance *Fixed) Error
	GetAdvances                      func(f *Face, start, count uint, loadFlags Int32, padvances *Fixed) Error
	GetBDFCharsetID                  func(f *Face, acharsetEncoding, acharset_registry **T.Char) Error
	GetBDFProperty                   func(f *Face, propName string, aproperty *BDFProperty) Error
	GetCharIndex                     func(f *Face, charcode ULong) uint
	GetCIDFromGlyphIndex             func(f *Face, glyphIndex uint, cid *uint) Error
	GetCIDIsInternallyCIDKeyed       func(f *Face, isCid *Bool) Error
	GetCIDRegistryOrderingSupplement func(f *Face, registry, ordering **T.Char, supplement *int) Error
	GetFirstChar                     func(f *Face, agindex *uint) ULong
	GetFSTypeFlags                   func(f *Face) uint16
	GetGasp                          func(f *Face, ppem uint) int
	GetGlyphName                     func(f *Face, glyphIndex uint, buffer *Void, bufferMax uint) Error
	GetKerning                       func(f *Face, leftGlyph, rightGlyph, kernMode uint, akerning *Vector) Error
	GetMMVar                         func(f *Face, amaster **MMVar) Error
	GetMultiMaster                   func(f *Face, amaster *MultiMaster) Error
	GetNameIndex                     func(f *Face, glyphName *String) uint
	GetNextChar                      func(f *Face, charCode ULong, agindex *uint) ULong
	GetPFRAdvance                    func(f *Face, gindex uint, aadvance *Pos) Error
	GetPFRKerning                    func(f *Face, left, right uint, avector *Vector) Error
	GetPFRMetrics                    func(f *Face, aoutline_resolution, ametrics_resolution *uint, ametrics_x_scale, ametrics_y_scale *Fixed) Error
	GetPostscriptName                func(f *Face) string
	GetPSFontInfo                    func(f *Face, afontInfo *PSFontInfo) Error
	GetPSFontPrivate                 func(f *Face, afontPrivate *PSPrivate) Error
	GetSfntName                      func(f *Face, idx uint, aname *SfntName) Error
	GetSfntNameCount                 func(f *Face) uint
	GetSfntTable                     func(f *Face, tag SfntTag) *T.Void
	GetTrackKerning                  func(f *Face, pointSize Fixed, degree int, akerning *Fixed) Error
	GetWinFNTHeader                  func(f *Face, aheader *WinFNTHeaderRec) Error
	GetX11FontFormat                 func(f *Face) string
	HasPSGlyphNames                  func(f *Face) int
	LoadChar                         func(f *Face, charCode ULong, loadFlags Int32) Error
	LoadGlyph                        func(f *Face, glyphIndex, loadFlags Int32) Error
	LoadSfntTable                    func(f *Face, tag ULong, offset Long, buffer *byte, length *ULong) Error
	NewSize                          func(f *Face, size **Size) Error
	OpenTypeFree                     func(f *Face, table *byte)
	OpenTypeValidate                 func(f *Face, validation_flags uint, BASETable, GDEF_table, GPOS_table, GSUBTable, JSTF_table **byte) Error
	ReferenceFace                    func(f *Face) Error
	RequestSize                      func(f *Face, req *SizeRequest) Error
	SelectCharmap                    func(f *Face, encoding Encoding) Error
	SelectSize                       func(f *Face, strikeIndex int) Error
	SetCharmap                       func(f *Face, charmap *CharMap) Error
	SetCharSize                      func(f *Face, charWidth, charHeight F26Dot6, horzResolution, vertResolution uint) Error
	SetMMBlendCoordinates            func(f *Face, numCoords uint, coords *Fixed) Error
	SetMMDesignCoordinates           func(f *Face, numCoords uint, coords *Long) Error
	SetPixelSizes                    func(f *Face, pixelWidth, pixelHeight uint) Error
	SetTransform                     func(f *Face, matrix *Matrix, delta *Vector)
	SetVarBlendCoordinates           func(f *Face, numCoords uint, coords *Fixed) Error
	SetVarDesignCoordinates          func(f *Face, numCoords uint, coords *Fixed) Error
	SfntTableInfo                    func(f *Face, tableIndex uint, tag, length *ULong) Error
	TrueTypeGXFree                   func(f *Face, table *byte)
	TrueTypeGXValidate               func(f *Face, validationFlags uint, tables [VALIDATE_GX_LENGTH]*byte, tableLength uint) Error
)

func (f *Face) CheckTrueTypePatents() Bool { return FaceCheckTrueTypePatents(f) }
func (f *Face) GetCharsOfVariant(variantSelector ULong) *UInt32 {
	return FaceGetCharsOfVariant(f, variantSelector)
}
func (f *Face) GetCharVariantIndex(charcode, variantSelector ULong) uint {
	return FaceGetCharVariantIndex(f, charcode, variantSelector)
}
func (f *Face) GetCharVariantIsDefault(charcode, variantSelector ULong) int {
	return FaceGetCharVariantIsDefault(f, charcode, variantSelector)
}
func (f *Face) GetVariantSelectors() *UInt32             { return FaceGetVariantSelectors(f) }
func (f *Face) GetVariantsOfChar(charcode ULong) *UInt32 { return FaceGetVariantsOfChar(f, charcode) }
func (f *Face) SetUnpatentedHinting(value Bool) Bool     { return FaceSetUnpatentedHinting(f, value) }

func (f *Face) AttachFile(filepathname string) Error    { return AttachFile(f, filepathname) }
func (f *Face) AttachStream(parameters *OpenArgs) Error { return AttachStream(f, parameters) }
func (f *Face) ClassicKernFree(table *byte)             { ClassicKernFree(f, table) }
func (f *Face) ClassicKernValidate(validation_flags uint, ckernTable **byte) Error {
	return ClassicKernValidate(f, validation_flags, ckernTable)
}
func (f *Face) DoneFace() Error { return DoneFace(f) }
func (f *Face) GetAdvance(gindex uint, loadFlags Int32, padvance *Fixed) Error {
	return GetAdvance(f, gindex, loadFlags, padvance)
}
func (f *Face) GetAdvances(start, count uint, loadFlags Int32, padvances *Fixed) Error {
	return GetAdvances(f, start, count, loadFlags, padvances)
}
func (f *Face) GetBDFCharsetID(acharsetEncoding, acharset_registry **T.Char) Error {
	return GetBDFCharsetID(f, acharsetEncoding, acharset_registry)
}
func (f *Face) GetBDFProperty(propName string, aproperty *BDFProperty) Error {
	return GetBDFProperty(f, propName, aproperty)
}
func (f *Face) GetCharIndex(charcode ULong) uint { return GetCharIndex(f, charcode) }
func (f *Face) GetCIDFromGlyphIndex(glyphIndex uint, cid *uint) Error {
	return GetCIDFromGlyphIndex(f, glyphIndex, cid)
}
func (f *Face) GetCIDIsInternallyCIDKeyed(isCid *Bool) Error {
	return GetCIDIsInternallyCIDKeyed(f, isCid)
}
func (f *Face) GetCIDRegistryOrderingSupplement(registry, ordering **T.Char, supplement *int) Error {
	return GetCIDRegistryOrderingSupplement(f, registry, ordering, supplement)
}
func (f *Face) GetFirstChar(agindex *uint) ULong { return GetFirstChar(f, agindex) }
func (f *Face) GetFSTypeFlags() uint16           { return GetFSTypeFlags(f) }
func (f *Face) GetGasp(ppem uint) int            { return GetGasp(f, ppem) }
func (f *Face) GetGlyphName(glyphIndex uint, buffer *Void, bufferMax uint) Error {
	return GetGlyphName(f, glyphIndex, buffer, bufferMax)
}
func (f *Face) GetKerning(leftGlyph, rightGlyph, kernMode uint, akerning *Vector) Error {
	return GetKerning(f, leftGlyph, rightGlyph, kernMode, akerning)
}
func (f *Face) GetMMVar(amaster **MMVar) Error            { return GetMMVar(f, amaster) }
func (f *Face) GetMultiMaster(amaster *MultiMaster) Error { return GetMultiMaster(f, amaster) }
func (f *Face) GetNameIndex(glyphName *String) uint       { return GetNameIndex(f, glyphName) }
func (f *Face) GetNextChar(charCode ULong, agindex *uint) ULong {
	return GetNextChar(f, charCode, agindex)
}
func (f *Face) GetPFRAdvance(gindex uint, aadvance *Pos) Error {
	return GetPFRAdvance(f, gindex, aadvance)
}
func (f *Face) GetPFRKerning(left, right uint, avector *Vector) Error {
	return GetPFRKerning(f, left, right, avector)
}
func (f *Face) GetPFRMetrics(aoutline_resolution, ametrics_resolution *uint, ametrics_x_scale, ametrics_y_scale *Fixed) Error {
	return GetPFRMetrics(f, aoutline_resolution, ametrics_resolution, ametrics_x_scale, ametrics_y_scale)
}
func (f *Face) GetPostscriptName() string                 { return GetPostscriptName(f) }
func (f *Face) GetPSFontInfo(afontInfo *PSFontInfo) Error { return GetPSFontInfo(f, afontInfo) }
func (f *Face) GetPSFontPrivate(afontPrivate *PSPrivate) Error {
	return GetPSFontPrivate(f, afontPrivate)
}
func (f *Face) GetSfntName(idx uint, aname *SfntName) Error { return GetSfntName(f, idx, aname) }
func (f *Face) GetSfntNameCount() uint                      { return GetSfntNameCount(f) }
func (f *Face) GetSfntTable(tag SfntTag) *T.Void            { return GetSfntTable(f, tag) }
func (f *Face) GetTrackKerning(pointSize Fixed, degree int, akerning *Fixed) Error {
	return GetTrackKerning(f, pointSize, degree, akerning)
}
func (f *Face) GetWinFNTHeader(aheader *WinFNTHeaderRec) Error { return GetWinFNTHeader(f, aheader) }
func (f *Face) GetX11FontFormat() string                       { return GetX11FontFormat(f) }
func (f *Face) HasPSGlyphNames() int                           { return HasPSGlyphNames(f) }
func (f *Face) LoadChar(charCode ULong, loadFlags Int32) Error {
	return LoadChar(f, charCode, loadFlags)
}
func (f *Face) LoadGlyph(glyphIndex, loadFlags Int32) Error {
	return LoadGlyph(f, glyphIndex, loadFlags)
}
func (f *Face) LoadSfntTable(tag ULong, offset Long, buffer *byte, length *ULong) Error {
	return LoadSfntTable(f, tag, offset, buffer, length)
}
func (f *Face) NewSize(size **Size) Error { return NewSize(f, size) }
func (f *Face) OpenTypeFree(table *byte)  { OpenTypeFree(f, table) }
func (f *Face) OpenTypeValidate(validation_flags uint, BASETable, GDEF_table, GPOS_table, GSUBTable, JSTF_table **byte) Error {
	return OpenTypeValidate(f, validation_flags, BASETable, GDEF_table, GPOS_table, GSUBTable, JSTF_table)
}
func (f *Face) ReferenceFace() Error                  { return ReferenceFace(f) }
func (f *Face) RequestSize(req *SizeRequest) Error    { return RequestSize(f, req) }
func (f *Face) SelectCharmap(encoding Encoding) Error { return SelectCharmap(f, encoding) }
func (f *Face) SelectSize(strikeIndex int) Error      { return SelectSize(f, strikeIndex) }
func (f *Face) SetCharmap(charmap *CharMap) Error     { return SetCharmap(f, charmap) }
func (f *Face) SetCharSize(charWidth, charHeight F26Dot6, horzResolution, vertResolution uint) Error {
	return SetCharSize(f, charWidth, charHeight, horzResolution, vertResolution)
}
func (f *Face) SetMMBlendCoordinates(numCoords uint, coords *Fixed) Error {
	return SetMMBlendCoordinates(f, numCoords, coords)
}
func (f *Face) SetMMDesignCoordinates(numCoords uint, coords *Long) Error {
	return SetMMDesignCoordinates(f, numCoords, coords)
}
func (f *Face) SetPixelSizes(pixelWidth, pixelHeight uint) Error {
	return SetPixelSizes(f, pixelWidth, pixelHeight)
}
func (f *Face) SetTransform(matrix *Matrix, delta *Vector) { SetTransform(f, matrix, delta) }
func (f *Face) SetVarBlendCoordinates(numCoords uint, coords *Fixed) Error {
	return SetVarBlendCoordinates(f, numCoords, coords)
}
func (f *Face) SetVarDesignCoordinates(numCoords uint, coords *Fixed) Error {
	return SetVarDesignCoordinates(f, numCoords, coords)
}
func (f *Face) SfntTableInfo(tableIndex uint, tag, length *ULong) Error {
	return SfntTableInfo(f, tableIndex, tag, length)
}
func (f *Face) TrueTypeGXFree(table *byte) { TrueTypeGXFree(f, table) }
func (f *Face) TrueTypeGXValidate(validationFlags uint, tables [VALIDATE_GX_LENGTH]*byte, tableLength uint) Error {
	return TrueTypeGXValidate(f, validationFlags, tables, tableLength)
}

type FaceID *struct{}

type FaceInternal struct{}

//NOTE(t): Unreferenced
// type Font *FontRec
// type FontRec struct {
// 	FaceId    FaceID
// 	PixWidth  uint16
// 	PixHeight uint16
// }

type Generic struct {
	Data      *Void
	Finalizer GenericFinalizer
}

type GenericFinalizer func(Object *Void)

type Glyph struct {
	Library *Library
	Clazz   *GlyphClass
	Format  GlyphFormat
	Advance Vector
}

var (
	GlyphCopy      func(g *Glyph, target **Glyph) Error
	GlyphTransform func(g *Glyph, matrix *Matrix, delta *Vector) Error
	GlyphGetCBox   func(g *Glyph, bboxMode uint, acbox *BBox)
	DoneGlyph      func(g *Glyph)
)

func (g *Glyph) Copy(target **Glyph) Error { return GlyphCopy(g, target) }
func (g *Glyph) Transform(matrix *Matrix, delta *Vector) Error {
	return GlyphTransform(g, matrix, delta)
}
func (g *Glyph) GetCBox(bboxMode uint, acbox *BBox) { GlyphGetCBox(g, bboxMode, acbox) }
func (g *Glyph) Done()                              { DoneGlyph(g) }

type (
	GlyphCopyFunc func(source, target *Glyph) Error

	GlyphDone func(glyph *Glyph)

	GlyphGetBBox func(glyph *Glyph, abbox *BBox)

	GlyphInit func(glyph *Glyph, slot *GlyphSlot) Error

	GlyphPrepare func(glyph *Glyph, slot *GlyphSlot) Error

	GlyphTransformFunc func(
		glyph *Glyph, matrix *Matrix, delta *Vector)
)

type GlyphClass struct {
	GlyphSize      Long
	GlyphFormat    GlyphFormat
	GlyphInit      GlyphInit
	GlyphDone      GlyphDone
	GlyphCopy      GlyphCopyFunc
	GlyphTransform GlyphTransformFunc
	GlyphBbox      GlyphGetBBox
	GlyphPrepare   GlyphPrepare
}

type GlyphFormat uint32 // TODO(t): enum/Size?

const (
	GLYPH_FORMAT_NONE      GlyphFormat = ((0 << 24) | (0 << 16) | (0 << 8) | 0)
	GLYPH_FORMAT_COMPOSITE GlyphFormat = (('c' << 24) | ('o' << 16) | ('m' << 8) | 'p')
	GLYPH_FORMAT_BITMAP    GlyphFormat = (('b' << 24) | ('i' << 16) | ('t' << 8) | 's')
	GLYPH_FORMAT_OUTLINE   GlyphFormat = (('o' << 24) | ('u' << 16) | ('t' << 8) | 'l')
	GLYPH_FORMAT_PLOTTER   GlyphFormat = (('p' << 24) | ('l' << 16) | ('o' << 8) | 't')
)

type GlyphMetrics struct {
	Width        Pos
	Height       Pos
	HoriBearingX Pos
	HoriBearingY Pos
	HoriAdvance  Pos
	VertBearingX Pos
	VertBearingY Pos
	VertAdvance  Pos
}

type GlyphSlot struct {
	Library           *Library
	Face              *Face
	Next              *GlyphSlot
	Reserved          uint
	Generic           Generic
	Metrics           GlyphMetrics
	LinearHoriAdvance Fixed
	LinearVertAdvance Fixed
	Advance           Vector
	Format            GlyphFormat
	Bitmap            Bitmap
	BitmapLeft        int
	BitmapTop         int
	Outline           Outline
	NumSubglyphs      uint
	Subglyphs         *SubGlyph
	ControlData       *Void
	ControlLen        Long
	LsbDelta          Pos
	RsbDelta          Pos
	Other             *Void
	Internal          *SlotInternal
}

var (
	GetGlyph           func(g *GlyphSlot, aglyph **Glyph) Error
	GetSubGlyphInfo    func(g *GlyphSlot, subIndex uint, index *int, flags *uint, arg1, arg2 *int, transform *Matrix) Error
	GlyphSlotEmbolden  func(g *GlyphSlot)
	GlyphSlotOblique   func(g *GlyphSlot)
	GlyphSlotOwnBitmap func(slot *GlyphSlot) Error

	RenderGlyph func(slot *GlyphSlot, renderMode RenderMode) Error
)

func (g *GlyphSlot) Get(aglyph **Glyph) Error { return GetGlyph(g, aglyph) }
func (g *GlyphSlot) Embolden()                { GlyphSlotEmbolden(g) }
func (g *GlyphSlot) Oblique()                 { GlyphSlotOblique(g) }
func (g *GlyphSlot) Info(subIndex uint, index *int, flags *uint, arg1, arg2 *int, transform *Matrix) Error {
	return GetSubGlyphInfo(g, subIndex, index, flags, arg1, arg2, transform)
}
func (g *GlyphSlot) OwnBitmap() Error                        { return GlyphSlotOwnBitmap(g) }
func (s *GlyphSlot) RenderGlyph(renderMode RenderMode) Error { return RenderGlyph(s, renderMode) }

var (
	ImageCacheLookup       func(i *ImageCache, t *ImageType, gindex uint, aglyph **Glyph, anode **Node) Error
	ImageCacheLookupScaler func(i *ImageCache, scaler *Scaler, loadFlags ULong, gindex uint, aglyph **Glyph, anode **Node) Error
)

func (i *ImageCache) Lookup(t *ImageType, gindex uint, aglyph **Glyph, anode **Node) Error {
	return ImageCacheLookup(i, t, gindex, aglyph, anode)
}
func (i *ImageCache) LookupScaler(scaler *Scaler, loadFlags ULong, gindex uint, aglyph **Glyph, anode **Node) Error {
	return ImageCacheLookupScaler(i, scaler, loadFlags, gindex, aglyph, anode)
}

type ImageType struct {
	FaceId FaceID
	Width  int
	Height int
	Flags  Int32
}

type LcdFilter Enum

const (
	LCD_FILTER_DEFAULT LcdFilter = 1 << iota
	LCD_FILTER_LIGHT
	_
	_
	LCD_FILTER_LEGACY
	LCD_FILTER_MAX            = LCD_FILTER_LEGACY + 1
	LCD_FILTER_NONE LcdFilter = 0
)

type Library struct{}

var (
	InitFreeType func(alibrary **Library) Error

	LibrarySetLcdFilter        func(l *Library, filter LcdFilter) Error
	LibrarySetLcdFilterWeights func(l *Library, weights *T.UnsignedChar) Error
	LibraryVersion             func(l *Library, amajor, aminor, apatch *int)

	AddDefaultModules     func(l *Library)
	AddModule             func(l *Library, clazz *ModuleClass) Error
	DoneFreeType          func(l *Library) Error
	DoneLibrary           func(l *Library) Error
	GetModule             func(l *Library, moduleName string) *Module
	GetRenderer           func(l *Library, format GlyphFormat) *Renderer
	GetTrueTypeEngineType func(l *Library) TrueTypeEngineType
	NewFace               func(l *Library, filepathname string, faceIndex Long, aface **Face) Error
	NewMemoryFace         func(l *Library, fileBase *byte, fileSize, faceIndex Long, aface **Face) Error
	OpenFace              func(l *Library, args *OpenArgs, faceIndex Long, aface **Face) Error
	ReferenceLibrary      func(l *Library) Error
	RemoveModule          func(l *Library, module *Module) Error
	SetDebugHook          func(l *Library, hookIndex uint, debugHook DebugHookFunc)
	SetRenderer           func(l *Library, renderer *Renderer, numParams uint, parameters *Parameter) Error

	StrokerNew func(l *Library, astroker **Stroker) Error
	ManagerNew func(l *Library, maxFaces, maxSizes uint, maxBytes ULong, requester FaceRequester, reqData *Void, amanager **Manager) Error
	OutlineNew func(l *Library, numPoints uint, numContours int, anoutline **Outline) Error
)

func (l *Library) SetLcdFilter(filter LcdFilter) Error { return LibrarySetLcdFilter(l, filter) }
func (l *Library) SetLcdFilterWeights(weights *T.UnsignedChar) Error {
	return LibrarySetLcdFilterWeights(l, weights)
}
func (l *Library) Version(amajor, aminor, apatch *int) { LibraryVersion(l, amajor, aminor, apatch) }

func (l *Library) AddDefaultModules()                 { AddDefaultModules(l) }
func (l *Library) AddModule(clazz *ModuleClass) Error { return AddModule(l, clazz) }
func (l *Library) DoneFreeType() Error                { return DoneFreeType(l) }
func (l *Library) DoneLibrary() Error                 { return DoneLibrary(l) }
func (l *Library) ManagerNew(maxFaces, maxSizes uint, maxBytes ULong, requester FaceRequester, reqData *Void, amanager **Manager) Error {
	return ManagerNew(l, maxFaces, maxSizes, maxBytes, requester, reqData, amanager)
}
func (l *Library) Module(moduleName string) *Module { return GetModule(l, moduleName) }
func (l *Library) NewFace(filepathname string, faceIndex Long, aface **Face) Error {
	return NewFace(l, filepathname, faceIndex, aface)
}
func (l *Library) NewMemoryFace(fileBase *byte, fileSize, faceIndex Long, aface **Face) Error {
	return NewMemoryFace(l, fileBase, fileSize, faceIndex, aface)
}
func (l *Library) OpenFace(args *OpenArgs, faceIndex Long, aface **Face) Error {
	return OpenFace(l, args, faceIndex, aface)
}
func (l *Library) Renderer(format GlyphFormat) *Renderer { return GetRenderer(l, format) }
func (l *Library) ReferenceLibrary() Error               { return ReferenceLibrary(l) }
func (l *Library) RemoveModule(module *Module) Error     { return RemoveModule(l, module) }
func (l *Library) SetDebugHook(hookIndex uint, debugHook DebugHookFunc) {
	SetDebugHook(l, hookIndex, debugHook)
}
func (l *Library) SetRenderer(renderer *Renderer, numParams uint, parameters *Parameter) Error {
	return SetRenderer(l, renderer, numParams, parameters)
}
func (l *Library) StrokerNew(astroker **Stroker) Error    { return StrokerNew(l, astroker) }
func (l *Library) TrueTypeEngineType() TrueTypeEngineType { return GetTrueTypeEngineType(l) }
func (l *Library) OutlineNew(numPoints uint, numContours int, anoutline **Outline) Error {
	return OutlineNew(l, numPoints, numContours, anoutline)
}

type (
	List struct{ Head, Tail *ListNode }

	ListDestructor func(memory *Memory, data, user *Void)

	ListIterator func(node *ListNode, user *Void) Error
)

var (
	ListAdd      func(l *List, node *ListNode)
	ListFinalize func(l *List, destroy ListDestructor, memory *Memory, user *T.Void)
	ListFind     func(l *List, data *T.Void) *ListNode
	ListInsert   func(l *List, node *ListNode)
	ListIterate  func(l *List, iterator ListIterator, user *T.Void) Error
	ListRemove   func(l *List, node *ListNode)
	ListUp       func(l *List, node *ListNode)
)

func (l *List) Add(node *ListNode) { ListAdd(l, node) }
func (l *List) Finalize(destroy ListDestructor, memory *Memory, user *T.Void) {
	ListFinalize(l, destroy, memory, user)
}
func (l *List) Find(data *T.Void) *ListNode { return ListFind(l, data) }
func (l *List) Insert(node *ListNode)       { ListInsert(l, node) }
func (l *List) Iterate(iterator ListIterator, user *T.Void) Error {
	return ListIterate(l, iterator, user)
}
func (l *List) Remove(node *ListNode) { ListRemove(l, node) }
func (l *List) Up(node *ListNode)     { ListUp(l, node) }

type ListNode struct {
	Prev, Next *ListNode
	Data       *Void
}

type Manager struct{}

var (
	ManagerDone         func(m *Manager)
	ManagerLookupFace   func(m *Manager, faceId FaceID, aface **Face) Error
	ManagerLookupSize   func(m *Manager, scaler *Scaler, asize **Size) Error
	ManagerRemoveFaceID func(m *Manager, faceId FaceID)
	ManagerReset        func(m *Manager)

	CMapCacheNew  func(m *Manager, acache **CMapCache) Error
	ImageCacheNew func(m *Manager, acache **ImageCache) Error
	SBitCacheNew  func(m *Manager, acache **SBitCache) Error
)

func (m *Manager) Done() { ManagerDone(m) }
func (m *Manager) LookupFace(faceId FaceID, aface **Face) Error {
	return ManagerLookupFace(m, faceId, aface)
}
func (m *Manager) LookupSize(scaler *Scaler, asize **Size) Error {
	return ManagerLookupSize(m, scaler, asize)
}
func (m *Manager) RemoveFaceID(faceId FaceID) { ManagerRemoveFaceID(m, faceId) }
func (m *Manager) Reset()                     { ManagerReset(m) }

func (m *Manager) CMapCacheNew(acache **CMapCache) Error   { return CMapCacheNew(m, acache) }
func (m *Manager) ImageCacheNew(acache **ImageCache) Error { return ImageCacheNew(m, acache) }
func (m *Manager) SBitCacheNew(acache **SBitCache) Error   { return SBitCacheNew(m, acache) }

type Matrix struct{ XX, XY, YX, YY Fixed }

var (
	MatrixMultiply func(a, b *Matrix)
	MatrixInvert   func(matrix *Matrix) Error
)

func (m *Matrix) Multiply(m2 *Matrix) { MatrixMultiply(m, m2) }
func (m *Matrix) Invert() Error       { return MatrixInvert(m) }

type (
	Memory struct {
		User    *Void
		Alloc   AllocFunc
		Free    FreeFunc
		Realloc ReallocFunc
	}

	AllocFunc func(memory *Memory, size Long) *Void

	FreeFunc func(memory *Memory, block *Void)

	ReallocFunc func(
		memory *Memory, curSize, newSize Long, block *Void) *Void
)

var (
	NewLibrary         func(m *Memory, alibrary **Library) Error
	OutlineNewInternal func(m *Memory, numPoints uint, numContours int, anoutline **Outline) Error
)

func (m *Memory) NewLibrary(alibrary **Library) Error { return NewLibrary(m, alibrary) }
func (m *Memory) OutlineNewInternal(numPoints uint, numContours int, anoutline **Outline) Error {
	return OutlineNewInternal(m, numPoints, numContours, anoutline)
}

type MMAxis struct {
	Name    *String
	Minimum Long
	Maximum Long
}

type MMVar struct {
	NumAxis        uint
	NumDesigns     uint
	NumNamedstyles uint
	Axis           *VarAxis
	Namedstyle     *VarNamedStyle
}

type (
	Module struct{}

	ModuleClass struct {
		ModuleFlags     ULong
		ModuleSize      Long
		ModuleName      *String
		ModuleVersion   Fixed
		ModuleRequires  Fixed
		ModuleInterface *Void
		ModuleInit      ModuleConstructor
		ModuleDone      ModuleDestructor
		GetInterface    ModuleRequester
	}

	ModuleConstructor func(module *Module) Error

	ModuleDestructor func(module *Module)

	ModuleInterface struct{}

	ModuleRequester func(
		module *Module, name string) *ModuleInterface
)

const T1_MAX_MM_AXIS = 4

type MultiMaster struct {
	NumAxis    uint
	NumDesigns uint
	Axis       [T1_MAX_MM_AXIS]MMAxis
}

type Node struct{}

var NodeUnref func(node *Node, manager *Manager)

func (n *Node) Unref(manager *Manager) { NodeUnref(n, manager) }

type OpenArgs struct {
	Flags      uint
	MemoryBase *byte
	MemorySize Long
	Pathname   *String
	Stream     Stream
	Driver     *Module
	NumParams  int
	Params     *Parameter
}

type Orientation Enum

const (
	ORIENTATION_TRUETYPE Orientation = iota
	ORIENTATION_POSTSCRIPT
	ORIENTATION_NONE
	ORIENTATION_FILL_RIGHT = ORIENTATION_TRUETYPE
	ORIENTATION_FILL_LEFT  = ORIENTATION_POSTSCRIPT
)

type Outline struct {
	NContours int16
	NPoints   int16
	Points    *Vector
	Tags      *Char
	Contours  *int16
	Flags     int
}

var (
	OutlineCheck            func(o *Outline) Error
	OutlineCopy             func(o *Outline, target *Outline) Error
	OutlineDecompose        func(o *Outline, funcInterface *OutlineFuncs, user *T.Void) Error
	OutlineEmbolden         func(o *Outline, strength Pos) Error
	OutlineGetBBox          func(o *Outline, abbox *BBox) Error
	OutlineGetCBox          func(o *Outline, acbox *BBox)
	OutlineGetInsideBorder  func(o *Outline) StrokerBorder
	OutlineGetOrientation   func(o *Outline) Orientation
	OutlineGetOutsideBorder func(o *Outline) StrokerBorder
	OutlineReverse          func(o *Outline)
	OutlineTransform        func(o *Outline, matrix *Matrix)
	OutlineTranslate        func(o *Outline, xOffset, yOffset Pos)

	OutlineDoneInternal func(memory *Memory, outline *Outline) Error
	OutlineGetBitmap    func(library *Library, outline *Outline, abitmap *Bitmap) Error
	OutlineRender       func(library *Library, outline *Outline, params *RasterParams) Error
	OutlineDone         func(library *Library, outline *Outline) Error
)

func (o *Outline) Check() Error               { return OutlineCheck(o) }
func (o *Outline) Copy(target *Outline) Error { return OutlineCopy(o, target) }
func (o *Outline) Decompose(funcInterface *OutlineFuncs, user *T.Void) Error {
	return OutlineDecompose(o, funcInterface, user)
}
func (o *Outline) Embolden(strength Pos) Error     { return OutlineEmbolden(o, strength) }
func (o *Outline) GetBBox(abbox *BBox) Error       { return OutlineGetBBox(o, abbox) }
func (o *Outline) GetCBox(acbox *BBox)             { OutlineGetCBox(o, acbox) }
func (o *Outline) GetInsideBorder() StrokerBorder  { return OutlineGetInsideBorder(o) }
func (o *Outline) GetOrientation() Orientation     { return OutlineGetOrientation(o) }
func (o *Outline) GetOutsideBorder() StrokerBorder { return OutlineGetOutsideBorder(o) }
func (o *Outline) Reverse()                        { OutlineReverse(o) }
func (o *Outline) Transform(matrix *Matrix)        { OutlineTransform(o, matrix) }
func (o *Outline) Translate(xOffset, yOffset Pos)  { OutlineTranslate(o, xOffset, yOffset) }

func (o *Outline) DoneInternal(memory *Memory) Error { return OutlineDoneInternal(memory, o) }
func (o *Outline) GetBitmap(library *Library, abitmap *Bitmap) Error {
	return OutlineGetBitmap(library, o, abitmap)
}
func (o *Outline) Render(library *Library, params *RasterParams) Error {
	return OutlineRender(library, o, params)
}
func (o *Outline) Done(library *Library) Error { return OutlineDone(library, o) }

type (
	OutlineFuncs struct {
		MoveTo  OutlineMoveTo
		LineTo  OutlineLineTo
		ConicTo OutlineConicTo
		CubicTo OutlineCubicTo
		Shift   int
		Delta   Pos
	}

	OutlineMoveTo func(to *Vector, user *Void) int

	OutlineLineTo func(to *Vector, user *Void) int

	OutlineConicTo func(
		control, to *Vector, user *Void) int

	OutlineCubicTo func(
		control1, control2 *Vector,
		to *Vector, user *Void) int
)

type Parameter struct {
	Tag  ULong
	Data *Void
}

type PSFontInfo struct {
	Version            *String
	Notice             *String
	FullName           *String
	FamilyName         *String
	Weight             *String
	ItalicAngle        Long
	IsFixedPitch       Bool
	UnderlinePosition  int16
	UnderlineThickness uint16
}

type PSPrivate struct {
	UniqueId            int
	LenIV               int
	NumBlueValues       byte
	NumOtherBlues       byte
	NumFamilyBlues      byte
	NumFamilyOtherBlues byte
	BlueValues          [14]int16
	OtherBlues          [10]int16
	FamilyBlues         [14]int16
	FamilyOtherBlues    [10]int16
	BlueScale           Fixed
	BlueShift           int
	BlueFuzz            int
	StandardWidth       [1]uint16
	StandardHeight      [1]uint16
	NumSnapWidths       byte
	NumSnapHeights      byte
	ForceBold           Bool
	RoundStemUp         Bool
	SnapWidths          [13]int16
	SnapHeights         [13]int16
	ExpansionFactor     Fixed
	LanguageGroup       Long
	Password            Long
	MinFeature          [2]int16
}

type RasterParams struct {
	Target     *Bitmap
	Source     *Void
	Flags      int
	GraySpans  SpanFunc
	BlackSpans SpanFunc
	BitTest    RasterBitTest
	BitSet     RasterBitSet
	User       *Void
	ClipBox    BBox
}

type RenderMode Enum

const (
	RENDER_MODE_NORMAL RenderMode = iota
	RENDER_MODE_LIGHT
	RENDER_MODE_MONO
	RENDER_MODE_LCD
	RENDER_MODE_LCD_V
	RENDER_MODE_MAX
)

type SBit struct {
	Width    byte
	Height   byte
	Left     Char
	Top      Char
	Format   byte
	MaxGrays byte
	Pitch    int16
	Xadvance Char
	Yadvance Char
	Buffer   *byte
}

var (
	SBitCacheLookup       func(s *SBitCache, t *ImageType, gindex uint, sbit **SBit, anode **Node) Error
	SBitCacheLookupScaler func(s *SBitCache, scaler *Scaler, loadFlags ULong, gindex uint, sbit **SBit, anode **Node) Error
)

func (s *SBitCache) Lookup(t *ImageType, gindex uint, sbit **SBit, anode **Node) Error {
	return SBitCacheLookup(s, t, gindex, sbit, anode)
}
func (s *SBitCache) LookupScaler(scaler *Scaler, loadFlags ULong, gindex uint, sbit **SBit, anode **Node) Error {
	return SBitCacheLookupScaler(s, scaler, loadFlags, gindex, sbit, anode)
}

type SBitCache struct{}

type Scaler struct {
	FaceId FaceID
	Width  uint
	Height uint
	Pixel  int
	XRes   uint
	YRes   uint
}

type SfntName struct {
	PlatformId uint16
	EncodingId uint16
	LanguageId uint16
	NameId     uint16
	String     *byte //STRING?
	StringLen  uint
}

type SfntTag Enum

const (
	Ft_sfnt_head SfntTag = iota
	Ft_sfnt_maxp
	Ft_sfnt_os2
	Ft_sfnt_hhea
	Ft_sfnt_vhea
	Ft_sfnt_post
	Ft_sfnt_pclt
	Sfnt_max
)

type Size struct {
	Face     *Face
	Generic  Generic
	Metrics  SizeMetrics
	Internal *SizeInternal
}

var (
	DoneSize     func(size *Size) Error
	ActivateSize func(size *Size) Error
)

func (s *Size) Done() Error     { return DoneSize(s) }
func (s *Size) Activate() Error { return ActivateSize(s) }

type SizeInternal struct{}

type SizeMetrics struct {
	XPpem      uint16
	YPpem      uint16
	XScale     Fixed
	YScale     Fixed
	Ascender   Pos
	Descender  Pos
	Height     Pos
	MaxAdvance Pos
}

type SizeRequest struct {
	Type           SizeRequestType
	Width          Long
	Height         Long
	HoriResolution uint
	VertResolution uint
}

type SizeRequestType Enum

const (
	SIZE_REQUEST_TYPE_NOMINAL SizeRequestType = iota
	SIZE_REQUEST_TYPE_REAL_DIM
	SIZE_REQUEST_TYPE_BBOX
	SIZE_REQUEST_TYPE_CELL
	SIZE_REQUEST_TYPE_SCALES
	SIZE_REQUEST_TYPE_MAX
)

type SlotInternal struct{}

type Span struct {
	X        int16
	Len      uint16
	Coverage uint8
}

type Stream struct{}

// type StreamDesc struct {
// 	//Union
// 	Value   Long
// 	Pointer *Void
// }

var (
	StreamOpenGzip func(stream, source *Stream) Error
	StreamOpenLZW  func(stream, source *Stream) Error
)

func (s *Stream) OpenGzip(source *Stream) Error { return StreamOpenGzip(s, source) }
func (s *Stream) OpenLZW(source *Stream) Error  { return StreamOpenLZW(s, source) }

type StrokerBorder Enum

const (
	STROKER_BORDER_LEFT StrokerBorder = iota
	STROKER_BORDER_RIGHT
)

type Stroker struct{}

var (
	StrokerBeginSubPath    func(s *Stroker, to *Vector, open Bool) Error
	StrokerConicTo         func(s *Stroker, control, to *Vector) Error
	StrokerCubicTo         func(s *Stroker, control1, control2, to *Vector) Error
	StrokerDone            func(s *Stroker)
	StrokerEndSubPath      func(s *Stroker) Error
	StrokerExport          func(s *Stroker, outline *Outline)
	StrokerExportBorder    func(s *Stroker, border StrokerBorder, outline *Outline)
	StrokerGetBorderCounts func(s *Stroker, border StrokerBorder, anumPoints, anum_contours *uint) Error
	StrokerGetCounts       func(s *Stroker, anumPoints, anum_contours *uint) Error
	StrokerLineTo          func(s *Stroker, to *Vector) Error
	StrokerParseOutline    func(s *Stroker, outline *Outline, opened Bool) Error
	StrokerRewind          func(s *Stroker)
	StrokerSet             func(s *Stroker, radius Fixed, lineCap StrokerLineCap, lineJoin StrokerLineJoin, miterLimit Fixed)
)

func (s *Stroker) BeginSubPath(to *Vector, open Bool) Error { return StrokerBeginSubPath(s, to, open) }
func (s *Stroker) ConicTo(control, to *Vector) Error        { return StrokerConicTo(s, control, to) }
func (s *Stroker) CubicTo(control1, control2, to *Vector) Error {
	return StrokerCubicTo(s, control1, control2, to)
}
func (s *Stroker) Done()                   { StrokerDone(s) }
func (s *Stroker) EndSubPath() Error       { return StrokerEndSubPath(s) }
func (s *Stroker) Export(outline *Outline) { StrokerExport(s, outline) }
func (s *Stroker) ExportBorder(border StrokerBorder, outline *Outline) {
	StrokerExportBorder(s, border, outline)
}
func (s *Stroker) GetBorderCounts(border StrokerBorder, anumPoints, anum_contours *uint) Error {
	return StrokerGetBorderCounts(s, border, anumPoints, anum_contours)
}
func (s *Stroker) GetCounts(anumPoints, anum_contours *uint) Error {
	return StrokerGetCounts(s, anumPoints, anum_contours)
}
func (s *Stroker) LineTo(to *Vector) Error { return StrokerLineTo(s, to) }
func (s *Stroker) ParseOutline(outline *Outline, opened Bool) Error {
	return StrokerParseOutline(s, outline, opened)
}
func (s *Stroker) Rewind() { StrokerRewind(s) }
func (s *Stroker) Set(radius Fixed, lineCap StrokerLineCap, lineJoin StrokerLineJoin, miterLimit Fixed) {
	StrokerSet(s, radius, lineCap, lineJoin, miterLimit)
}

type StrokerLineCap Enum

const (
	STROKER_LINECAP_BUTT StrokerLineCap = iota
	STROKER_LINECAP_ROUND
	STROKER_LINECAP_SQUARE
)

type StrokerLineJoin Enum

const (
	STROKER_LINEJOIN_ROUND StrokerLineJoin = iota
	STROKER_LINEJOIN_BEVEL
	STROKER_LINEJOIN_MITER
)

type SubGlyph struct{}

type TrueTypeEngineType Enum

const (
	TRUETYPE_ENGINE_TYPE_NONE TrueTypeEngineType = iota
	TRUETYPE_ENGINE_TYPE_UNPATENTED
	TRUETYPE_ENGINE_TYPE_PATENTED
)

type VarAxis struct {
	Name    *String
	Minimum Fixed
	Def     Fixed
	Maximum Fixed
	Tag     ULong
	Strid   uint
}

type VarNamedStyle struct {
	Coords *Fixed
	Strid  uint
}

type Vector struct {
	X Pos
	Y Pos
}

var (
	VectorFromPolar func(v *Vector, length Fixed, angle Angle)
	VectorLength    func(v *Vector) Fixed
	VectorPolarize  func(v *Vector, length *Fixed, angle *Angle)
	VectorRotate    func(v *Vector, angle Angle)
	VectorTransform func(v *Vector, matrix *Matrix)
	VectorUnit      func(v *Vector, angle Angle)
)

func (v *Vector) FromPolar(length Fixed, angle Angle)  { VectorFromPolar(v, length, angle) }
func (v *Vector) Length() Fixed                        { return VectorLength(v) }
func (v *Vector) Polarize(length *Fixed, angle *Angle) { VectorPolarize(v, length, angle) }
func (v *Vector) Rotate(angle Angle)                   { VectorRotate(v, angle) }
func (v *Vector) Transform(matrix *Matrix)             { VectorTransform(v, matrix) }
func (v *Vector) Unit(angle Angle)                     { VectorUnit(v, angle) }

type WinFNTHeaderRec struct {
	Version              uint16
	FileSize             ULong
	Copyright            [60]byte
	FileType             uint16
	NominalPointSize     uint16
	VerticalResolution   uint16
	HorizontalResolution uint16
	Ascent               uint16
	InternalLeading      uint16
	ExternalLeading      uint16
	Italic               byte
	Underline            byte
	StrikeOut            byte
	Weight               uint16
	Charset              byte
	PixelWidth           uint16
	PixelHeight          uint16
	PitchAndFamily       byte
	AvgWidth             uint16
	MaxWidth             uint16
	FirstChar            byte
	LastChar             byte
	DefaultChar          byte
	BreakChar            byte
	BytesPerRow          uint16
	DeviceOffset         ULong
	FaceNameOffset       ULong
	BitsPointer          ULong
	BitsOffset           ULong
	Reserved             byte
	Flags                ULong
	ASpace               uint16
	BSpace               uint16
	CSpace               uint16
	ColorTableOffset     uint16
	Reserved1            [4]ULong
}

var dll = "freetype6.dll"

var apiList = outside.Apis{
	{"FTC_CMapCache_Lookup", &CMapCacheLookup},
	{"FTC_CMapCache_New", &CMapCacheNew},
	{"FTC_ImageCache_Lookup", &ImageCacheLookup},
	{"FTC_ImageCache_LookupScaler", &ImageCacheLookupScaler},
	{"FTC_ImageCache_New", &ImageCacheNew},
	{"FTC_Manager_Done", &ManagerDone},
	{"FTC_Manager_LookupFace", &ManagerLookupFace},
	{"FTC_Manager_LookupSize", &ManagerLookupSize},
	{"FTC_Manager_Lookup_Face", &ManagerLookupFace},
	{"FTC_Manager_Lookup_Size", &ManagerLookupSize},
	{"FTC_Manager_New", &ManagerNew},
	{"FTC_Manager_RemoveFaceID", &ManagerRemoveFaceID},
	{"FTC_Manager_Reset", &ManagerReset},
	{"FTC_Node_Unref", &NodeUnref},
	{"FTC_SBitCache_Lookup", &SBitCacheLookup},
	{"FTC_SBitCache_LookupScaler", &SBitCacheLookupScaler},
	{"FTC_SBitCache_New", &SBitCacheNew},
	{"FT_Activate_Size", &ActivateSize},
	{"FT_Add_Default_Modules", &AddDefaultModules},
	{"FT_Add_Module", &AddModule},
	{"FT_Angle_Diff", &AngleDiff},
	{"FT_Atan2", &Atan2},
	{"FT_Attach_File", &AttachFile},
	{"FT_Attach_Stream", &AttachStream},
	{"FT_Bitmap_Convert", &BitmapConvert},
	{"FT_Bitmap_Copy", &BitmapCopy},
	{"FT_Bitmap_Done", &BitmapDone},
	{"FT_Bitmap_Embolden", &BitmapEmbolden},
	{"FT_Bitmap_New", &BitmapNew},
	{"FT_CeilFix", &CeilFix},
	{"FT_ClassicKern_Free", &ClassicKernFree},
	{"FT_ClassicKern_Validate", &ClassicKernValidate},
	{"FT_Cos", &Cos},
	{"FT_DivFix", &DivFix},
	{"FT_Done_Face", &DoneFace},
	{"FT_Done_FreeType", &DoneFreeType},
	{"FT_Done_Glyph", &DoneGlyph},
	{"FT_Done_Library", &DoneLibrary},
	{"FT_Done_Size", &DoneSize},
	{"FT_Face_CheckTrueTypePatents", &FaceCheckTrueTypePatents},
	{"FT_Face_GetCharVariantIndex", &FaceGetCharVariantIndex},
	{"FT_Face_GetCharVariantIsDefault", &FaceGetCharVariantIsDefault},
	{"FT_Face_GetCharsOfVariant", &FaceGetCharsOfVariant},
	{"FT_Face_GetVariantSelectors", &FaceGetVariantSelectors},
	{"FT_Face_GetVariantsOfChar", &FaceGetVariantsOfChar},
	{"FT_Face_SetUnpatentedHinting", &FaceSetUnpatentedHinting},
	{"FT_FloorFix", &FloorFix},
	{"FT_Get_Advance", &GetAdvance},
	{"FT_Get_Advances", &GetAdvances},
	{"FT_Get_BDF_Charset_ID", &GetBDFCharsetID},
	{"FT_Get_BDF_Property", &GetBDFProperty},
	{"FT_Get_CID_From_Glyph_Index", &GetCIDFromGlyphIndex},
	{"FT_Get_CID_Is_Internally_CID_Keyed", &GetCIDIsInternallyCIDKeyed},
	{"FT_Get_CID_Registry_Ordering_Supplement", &GetCIDRegistryOrderingSupplement},
	{"FT_Get_CMap_Format", &GetCMapFormat},
	{"FT_Get_CMap_Language_ID", &GetCMapLanguageID},
	{"FT_Get_Char_Index", &GetCharIndex},
	{"FT_Get_Charmap_Index", &GetCharmapIndex},
	{"FT_Get_FSType_Flags", &GetFSTypeFlags},
	{"FT_Get_First_Char", &GetFirstChar},
	{"FT_Get_Gasp", &GetGasp},
	{"FT_Get_Glyph", &GetGlyph},
	{"FT_Get_Glyph_Name", &GetGlyphName},
	{"FT_Get_Kerning", &GetKerning},
	{"FT_Get_MM_Var", &GetMMVar},
	{"FT_Get_Module", &GetModule},
	{"FT_Get_Multi_Master", &GetMultiMaster},
	{"FT_Get_Name_Index", &GetNameIndex},
	{"FT_Get_Next_Char", &GetNextChar},
	{"FT_Get_PFR_Advance", &GetPFRAdvance},
	{"FT_Get_PFR_Kerning", &GetPFRKerning},
	{"FT_Get_PFR_Metrics", &GetPFRMetrics},
	{"FT_Get_PS_Font_Info", &GetPSFontInfo},
	{"FT_Get_PS_Font_Private", &GetPSFontPrivate},
	{"FT_Get_Postscript_Name", &GetPostscriptName},
	{"FT_Get_Renderer", &GetRenderer},
	{"FT_Get_Sfnt_Name", &GetSfntName},
	{"FT_Get_Sfnt_Name_Count", &GetSfntNameCount},
	{"FT_Get_Sfnt_Table", &GetSfntTable},
	{"FT_Get_SubGlyph_Info", &GetSubGlyphInfo},
	{"FT_Get_Track_Kerning", &GetTrackKerning},
	{"FT_Get_TrueType_Engine_Type", &GetTrueTypeEngineType},
	{"FT_Get_WinFNT_Header", &GetWinFNTHeader},
	{"FT_Get_X11_Font_Format", &GetX11FontFormat},
	{"FT_GlyphSlot_Embolden", &GlyphSlotEmbolden},
	{"FT_GlyphSlot_Oblique", &GlyphSlotOblique},
	{"FT_GlyphSlot_Own_Bitmap", &GlyphSlotOwnBitmap},
	{"FT_Glyph_Copy", &GlyphCopy},
	{"FT_Glyph_Get_CBox", &GlyphGetCBox},
	{"FT_Glyph_Stroke", &GlyphStroke},
	{"FT_Glyph_StrokeBorder", &GlyphStrokeBorder},
	{"FT_Glyph_To_Bitmap", &GlyphToBitmap},
	{"FT_Glyph_Transform", &GlyphTransform},
	{"FT_Has_PS_Glyph_Names", &HasPSGlyphNames},
	{"FT_Init_FreeType", &InitFreeType},
	{"FT_Library_SetLcdFilter", &LibrarySetLcdFilter},
	{"FT_Library_SetLcdFilterWeights", &LibrarySetLcdFilterWeights},
	{"FT_Library_Version", &LibraryVersion},
	{"FT_List_Add", &ListAdd},
	{"FT_List_Finalize", &ListFinalize},
	{"FT_List_Find", &ListFind},
	{"FT_List_Insert", &ListInsert},
	{"FT_List_Iterate", &ListIterate},
	{"FT_List_Remove", &ListRemove},
	{"FT_List_Up", &ListUp},
	{"FT_Load_Char", &LoadChar},
	{"FT_Load_Glyph", &LoadGlyph},
	{"FT_Load_Sfnt_Table", &LoadSfntTable},
	{"FT_Matrix_Invert", &MatrixInvert},
	{"FT_Matrix_Multiply", &MatrixMultiply},
	{"FT_MulDiv", &MulDiv},
	{"FT_MulFix", &MulFix},
	{"FT_New_Face", &NewFace},
	{"FT_New_Library", &NewLibrary},
	{"FT_New_Memory_Face", &NewMemoryFace},
	{"FT_New_Size", &NewSize},
	{"FT_OpenType_Free", &OpenTypeFree},
	{"FT_OpenType_Validate", &OpenTypeValidate},
	{"FT_Open_Face", &OpenFace},
	{"FT_Outline_Check", &OutlineCheck},
	{"FT_Outline_Copy", &OutlineCopy},
	{"FT_Outline_Decompose", &OutlineDecompose},
	{"FT_Outline_Done", &OutlineDone},
	{"FT_Outline_Done_Internal", &OutlineDoneInternal},
	{"FT_Outline_Embolden", &OutlineEmbolden},
	{"FT_Outline_GetInsideBorder", &OutlineGetInsideBorder},
	{"FT_Outline_GetOutsideBorder", &OutlineGetOutsideBorder},
	{"FT_Outline_Get_BBox", &OutlineGetBBox},
	{"FT_Outline_Get_Bitmap", &OutlineGetBitmap},
	{"FT_Outline_Get_CBox", &OutlineGetCBox},
	{"FT_Outline_Get_Orientation", &OutlineGetOrientation},
	{"FT_Outline_New", &OutlineNew},
	{"FT_Outline_New_Internal", &OutlineNewInternal},
	{"FT_Outline_Render", &OutlineRender},
	{"FT_Outline_Reverse", &OutlineReverse},
	{"FT_Outline_Transform", &OutlineTransform},
	{"FT_Outline_Translate", &OutlineTranslate},
	{"FT_Reference_Face", &ReferenceFace},
	{"FT_Reference_Library", &ReferenceLibrary},
	{"FT_Remove_Module", &RemoveModule},
	{"FT_Render_Glyph", &RenderGlyph},
	{"FT_Request_Size", &RequestSize},
	{"FT_RoundFix", &RoundFix},
	{"FT_Select_Charmap", &SelectCharmap},
	{"FT_Select_Size", &SelectSize},
	{"FT_Set_Char_Size", &SetCharSize},
	{"FT_Set_Charmap", &SetCharmap},
	{"FT_Set_Debug_Hook", &SetDebugHook},
	{"FT_Set_MM_Blend_Coordinates", &SetMMBlendCoordinates},
	{"FT_Set_MM_Design_Coordinates", &SetMMDesignCoordinates},
	{"FT_Set_Pixel_Sizes", &SetPixelSizes},
	{"FT_Set_Renderer", &SetRenderer},
	{"FT_Set_Transform", &SetTransform},
	{"FT_Set_Var_Blend_Coordinates", &SetVarBlendCoordinates},
	{"FT_Set_Var_Design_Coordinates", &SetVarDesignCoordinates},
	{"FT_Sfnt_Table_Info", &SfntTableInfo},
	{"FT_Sin", &Sin},
	{"FT_Stream_OpenGzip", &StreamOpenGzip},
	{"FT_Stream_OpenLZW", &StreamOpenLZW},
	{"FT_Stroker_BeginSubPath", &StrokerBeginSubPath},
	{"FT_Stroker_ConicTo", &StrokerConicTo},
	{"FT_Stroker_CubicTo", &StrokerCubicTo},
	{"FT_Stroker_Done", &StrokerDone},
	{"FT_Stroker_EndSubPath", &StrokerEndSubPath},
	{"FT_Stroker_Export", &StrokerExport},
	{"FT_Stroker_ExportBorder", &StrokerExportBorder},
	{"FT_Stroker_GetBorderCounts", &StrokerGetBorderCounts},
	{"FT_Stroker_GetCounts", &StrokerGetCounts},
	{"FT_Stroker_LineTo", &StrokerLineTo},
	{"FT_Stroker_New", &StrokerNew},
	{"FT_Stroker_ParseOutline", &StrokerParseOutline},
	{"FT_Stroker_Rewind", &StrokerRewind},
	{"FT_Stroker_Set", &StrokerSet},
	{"FT_Tan", &Tan},
	{"FT_TrueTypeGX_Free", &TrueTypeGXFree},
	{"FT_TrueTypeGX_Validate", &TrueTypeGXValidate},
	{"FT_Vector_From_Polar", &VectorFromPolar},
	{"FT_Vector_Length", &VectorLength},
	{"FT_Vector_Polarize", &VectorPolarize},
	{"FT_Vector_Rotate", &VectorRotate},
	{"FT_Vector_Transform", &VectorTransform},
	{"FT_Vector_Unit", &VectorUnit},
	// Undocumented {"TT_New_Context", &TTNewContext},
	// Undocumented {"TT_RunIns", &TTRunIns},
}
