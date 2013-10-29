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

	bitmapConvert  func(library *Library, source, target *Bitmap, alignment int) Error
	bitmapCopy     func(library *Library, source, target *Bitmap) Error
	bitmapDone     func(library *Library, bitmap *Bitmap) Error
	bitmapEmbolden func(library *Library, bitmap *Bitmap, xStrength, yStrength Pos) Error
)

func (b *Bitmap) Convert(target *Bitmap, library *Library, alignment int) Error {
	return bitmapConvert(library, b, target, alignment)
}
func (b *Bitmap) Copy(target *Bitmap, library *Library) Error { return bitmapCopy(library, b, target) }
func (b *Bitmap) Done(bitmap *Bitmap, library *Library) Error { return bitmapDone(library, b) }
func (b *Bitmap) Embolden(bitmap *Bitmap, library *Library, xStrength, yStrength Pos) Error {
	return bitmapEmbolden(library, b, xStrength, yStrength)
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
	getCMapLanguageID func(c *CharMap) ULong
	getCMapFormat     func(c *CharMap) Long
	getCharmapIndex   func(c *CharMap) int

	cMapCacheLookup func(cache *CMapCache, faceId FaceID, cmapIndex int, charCode UInt32) uint
)

func (c *CharMap) LanguageID() ULong { return getCMapLanguageID(c) }
func (c *CharMap) Format() Long      { return getCMapFormat(c) }
func (c *CharMap) Index() int        { return getCharmapIndex(c) }

func (c *CMapCache) Lookup(faceId FaceID, cmapIndex int, charCode UInt32) uint {
	return cMapCacheLookup(c, faceId, cmapIndex, charCode)
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
	faceCheckTrueTypePatents    func(f *Face) Bool
	faceGetCharsOfVariant       func(f *Face, variantSelector ULong) *UInt32
	faceGetCharVariantIndex     func(f *Face, charcode, variantSelector ULong) uint
	faceGetCharVariantIsDefault func(f *Face, charcode, variantSelector ULong) int
	faceGetVariantSelectors     func(f *Face) *UInt32
	faceGetVariantsOfChar       func(f *Face, charcode ULong) *UInt32
	faceSetUnpatentedHinting    func(f *Face, value Bool) Bool

	attachFile                       func(f *Face, filepathname string) Error
	attachStream                     func(f *Face, parameters *OpenArgs) Error
	classicKernFree                  func(f *Face, table *byte)
	classicKernValidate              func(f *Face, validation_flags uint, ckernTable **byte) Error
	doneFace                         func(f *Face) Error
	getAdvance                       func(f *Face, gindex uint, loadFlags Int32, padvance *Fixed) Error
	getAdvances                      func(f *Face, start, count uint, loadFlags Int32, padvances *Fixed) Error
	getBDFCharsetID                  func(f *Face, acharsetEncoding, acharset_registry **T.Char) Error
	getBDFProperty                   func(f *Face, propName string, aproperty *BDFProperty) Error
	getCharIndex                     func(f *Face, charcode ULong) uint
	getCIDFromGlyphIndex             func(f *Face, glyphIndex uint, cid *uint) Error
	getCIDIsInternallyCIDKeyed       func(f *Face, isCid *Bool) Error
	getCIDRegistryOrderingSupplement func(f *Face, registry, ordering **T.Char, supplement *int) Error
	getFirstChar                     func(f *Face, agindex *uint) ULong
	getFSTypeFlags                   func(f *Face) uint16
	getGasp                          func(f *Face, ppem uint) int
	getGlyphName                     func(f *Face, glyphIndex uint, buffer *Void, bufferMax uint) Error
	getKerning                       func(f *Face, leftGlyph, rightGlyph, kernMode uint, akerning *Vector) Error
	getMMVar                         func(f *Face, amaster **MMVar) Error
	getMultiMaster                   func(f *Face, amaster *MultiMaster) Error
	getNameIndex                     func(f *Face, glyphName *String) uint
	getNextChar                      func(f *Face, charCode ULong, agindex *uint) ULong
	getPFRAdvance                    func(f *Face, gindex uint, aadvance *Pos) Error
	getPFRKerning                    func(f *Face, left, right uint, avector *Vector) Error
	getPFRMetrics                    func(f *Face, aoutline_resolution, ametrics_resolution *uint, ametrics_x_scale, ametrics_y_scale *Fixed) Error
	getPostscriptName                func(f *Face) string
	getPSFontInfo                    func(f *Face, afontInfo *PSFontInfo) Error
	getPSFontPrivate                 func(f *Face, afontPrivate *PSPrivate) Error
	getSfntName                      func(f *Face, idx uint, aname *SfntName) Error
	getSfntNameCount                 func(f *Face) uint
	getSfntTable                     func(f *Face, tag SfntTag) *T.Void
	getTrackKerning                  func(f *Face, pointSize Fixed, degree int, akerning *Fixed) Error
	getWinFNTHeader                  func(f *Face, aheader *WinFNTHeaderRec) Error
	getX11FontFormat                 func(f *Face) string
	hasPSGlyphNames                  func(f *Face) int
	loadChar                         func(f *Face, charCode ULong, loadFlags Int32) Error
	loadGlyph                        func(f *Face, glyphIndex, loadFlags Int32) Error
	loadSfntTable                    func(f *Face, tag ULong, offset Long, buffer *byte, length *ULong) Error
	newSize                          func(f *Face, size **Size) Error
	openTypeFree                     func(f *Face, table *byte)
	openTypeValidate                 func(f *Face, validation_flags uint, BASETable, GDEF_table, GPOS_table, GSUBTable, JSTF_table **byte) Error
	referenceFace                    func(f *Face) Error
	requestSize                      func(f *Face, req *SizeRequest) Error
	selectCharmap                    func(f *Face, encoding Encoding) Error
	selectSize                       func(f *Face, strikeIndex int) Error
	setCharmap                       func(f *Face, charmap *CharMap) Error
	setCharSize                      func(f *Face, charWidth, charHeight F26Dot6, horzResolution, vertResolution uint) Error
	setMMBlendCoordinates            func(f *Face, numCoords uint, coords *Fixed) Error
	setMMDesignCoordinates           func(f *Face, numCoords uint, coords *Long) Error
	setPixelSizes                    func(f *Face, pixelWidth, pixelHeight uint) Error
	setTransform                     func(f *Face, matrix *Matrix, delta *Vector)
	setVarBlendCoordinates           func(f *Face, numCoords uint, coords *Fixed) Error
	setVarDesignCoordinates          func(f *Face, numCoords uint, coords *Fixed) Error
	sfntTableInfo                    func(f *Face, tableIndex uint, tag, length *ULong) Error
	trueTypeGXFree                   func(f *Face, table *byte)
	trueTypeGXValidate               func(f *Face, validationFlags uint, tables [VALIDATE_GX_LENGTH]*byte, tableLength uint) Error
)

func (f *Face) CheckTrueTypePatents() Bool { return faceCheckTrueTypePatents(f) }
func (f *Face) GetCharsOfVariant(variantSelector ULong) *UInt32 {
	return faceGetCharsOfVariant(f, variantSelector)
}
func (f *Face) GetCharVariantIndex(charcode, variantSelector ULong) uint {
	return faceGetCharVariantIndex(f, charcode, variantSelector)
}
func (f *Face) GetCharVariantIsDefault(charcode, variantSelector ULong) int {
	return faceGetCharVariantIsDefault(f, charcode, variantSelector)
}
func (f *Face) GetVariantSelectors() *UInt32             { return faceGetVariantSelectors(f) }
func (f *Face) GetVariantsOfChar(charcode ULong) *UInt32 { return faceGetVariantsOfChar(f, charcode) }
func (f *Face) SetUnpatentedHinting(value Bool) Bool     { return faceSetUnpatentedHinting(f, value) }

func (f *Face) AttachFile(filepathname string) Error    { return attachFile(f, filepathname) }
func (f *Face) AttachStream(parameters *OpenArgs) Error { return attachStream(f, parameters) }
func (f *Face) ClassicKernFree(table *byte)             { classicKernFree(f, table) }
func (f *Face) ClassicKernValidate(validation_flags uint, ckernTable **byte) Error {
	return classicKernValidate(f, validation_flags, ckernTable)
}
func (f *Face) DoneFace() Error { return doneFace(f) }
func (f *Face) GetAdvance(gindex uint, loadFlags Int32, padvance *Fixed) Error {
	return getAdvance(f, gindex, loadFlags, padvance)
}
func (f *Face) GetAdvances(start, count uint, loadFlags Int32, padvances *Fixed) Error {
	return getAdvances(f, start, count, loadFlags, padvances)
}
func (f *Face) GetBDFCharsetID(acharsetEncoding, acharset_registry **T.Char) Error {
	return getBDFCharsetID(f, acharsetEncoding, acharset_registry)
}
func (f *Face) GetBDFProperty(propName string, aproperty *BDFProperty) Error {
	return getBDFProperty(f, propName, aproperty)
}
func (f *Face) GetCharIndex(charcode ULong) uint { return getCharIndex(f, charcode) }
func (f *Face) GetCIDFromGlyphIndex(glyphIndex uint, cid *uint) Error {
	return getCIDFromGlyphIndex(f, glyphIndex, cid)
}
func (f *Face) GetCIDIsInternallyCIDKeyed(isCid *Bool) Error {
	return getCIDIsInternallyCIDKeyed(f, isCid)
}
func (f *Face) GetCIDRegistryOrderingSupplement(registry, ordering **T.Char, supplement *int) Error {
	return getCIDRegistryOrderingSupplement(f, registry, ordering, supplement)
}
func (f *Face) GetFirstChar(agindex *uint) ULong { return getFirstChar(f, agindex) }
func (f *Face) GetFSTypeFlags() uint16           { return getFSTypeFlags(f) }
func (f *Face) GetGasp(ppem uint) int            { return getGasp(f, ppem) }
func (f *Face) GetGlyphName(glyphIndex uint, buffer *Void, bufferMax uint) Error {
	return getGlyphName(f, glyphIndex, buffer, bufferMax)
}
func (f *Face) GetKerning(leftGlyph, rightGlyph, kernMode uint, akerning *Vector) Error {
	return getKerning(f, leftGlyph, rightGlyph, kernMode, akerning)
}
func (f *Face) GetMMVar(amaster **MMVar) Error            { return getMMVar(f, amaster) }
func (f *Face) GetMultiMaster(amaster *MultiMaster) Error { return getMultiMaster(f, amaster) }
func (f *Face) GetNameIndex(glyphName *String) uint       { return getNameIndex(f, glyphName) }
func (f *Face) GetNextChar(charCode ULong, agindex *uint) ULong {
	return getNextChar(f, charCode, agindex)
}
func (f *Face) GetPFRAdvance(gindex uint, aadvance *Pos) Error {
	return getPFRAdvance(f, gindex, aadvance)
}
func (f *Face) GetPFRKerning(left, right uint, avector *Vector) Error {
	return getPFRKerning(f, left, right, avector)
}
func (f *Face) GetPFRMetrics(aoutline_resolution, ametrics_resolution *uint, ametrics_x_scale, ametrics_y_scale *Fixed) Error {
	return getPFRMetrics(f, aoutline_resolution, ametrics_resolution, ametrics_x_scale, ametrics_y_scale)
}
func (f *Face) GetPostscriptName() string                 { return getPostscriptName(f) }
func (f *Face) GetPSFontInfo(afontInfo *PSFontInfo) Error { return getPSFontInfo(f, afontInfo) }
func (f *Face) GetPSFontPrivate(afontPrivate *PSPrivate) Error {
	return getPSFontPrivate(f, afontPrivate)
}
func (f *Face) GetSfntName(idx uint, aname *SfntName) Error { return getSfntName(f, idx, aname) }
func (f *Face) GetSfntNameCount() uint                      { return getSfntNameCount(f) }
func (f *Face) GetSfntTable(tag SfntTag) *T.Void            { return getSfntTable(f, tag) }
func (f *Face) GetTrackKerning(pointSize Fixed, degree int, akerning *Fixed) Error {
	return getTrackKerning(f, pointSize, degree, akerning)
}
func (f *Face) GetWinFNTHeader(aheader *WinFNTHeaderRec) Error { return getWinFNTHeader(f, aheader) }
func (f *Face) GetX11FontFormat() string                       { return getX11FontFormat(f) }
func (f *Face) HasPSGlyphNames() int                           { return hasPSGlyphNames(f) }
func (f *Face) LoadChar(charCode ULong, loadFlags Int32) Error {
	return loadChar(f, charCode, loadFlags)
}
func (f *Face) LoadGlyph(glyphIndex, loadFlags Int32) Error {
	return loadGlyph(f, glyphIndex, loadFlags)
}
func (f *Face) LoadSfntTable(tag ULong, offset Long, buffer *byte, length *ULong) Error {
	return loadSfntTable(f, tag, offset, buffer, length)
}
func (f *Face) NewSize(size **Size) Error { return newSize(f, size) }
func (f *Face) OpenTypeFree(table *byte)  { openTypeFree(f, table) }
func (f *Face) OpenTypeValidate(validation_flags uint, BASETable, GDEF_table, GPOS_table, GSUBTable, JSTF_table **byte) Error {
	return openTypeValidate(f, validation_flags, BASETable, GDEF_table, GPOS_table, GSUBTable, JSTF_table)
}
func (f *Face) ReferenceFace() Error                  { return referenceFace(f) }
func (f *Face) RequestSize(req *SizeRequest) Error    { return requestSize(f, req) }
func (f *Face) SelectCharmap(encoding Encoding) Error { return selectCharmap(f, encoding) }
func (f *Face) SelectSize(strikeIndex int) Error      { return selectSize(f, strikeIndex) }
func (f *Face) SetCharmap(charmap *CharMap) Error     { return setCharmap(f, charmap) }
func (f *Face) SetCharSize(charWidth, charHeight F26Dot6, horzResolution, vertResolution uint) Error {
	return setCharSize(f, charWidth, charHeight, horzResolution, vertResolution)
}
func (f *Face) SetMMBlendCoordinates(numCoords uint, coords *Fixed) Error {
	return setMMBlendCoordinates(f, numCoords, coords)
}
func (f *Face) SetMMDesignCoordinates(numCoords uint, coords *Long) Error {
	return setMMDesignCoordinates(f, numCoords, coords)
}
func (f *Face) SetPixelSizes(pixelWidth, pixelHeight uint) Error {
	return setPixelSizes(f, pixelWidth, pixelHeight)
}
func (f *Face) SetTransform(matrix *Matrix, delta *Vector) { setTransform(f, matrix, delta) }
func (f *Face) SetVarBlendCoordinates(numCoords uint, coords *Fixed) Error {
	return setVarBlendCoordinates(f, numCoords, coords)
}
func (f *Face) SetVarDesignCoordinates(numCoords uint, coords *Fixed) Error {
	return setVarDesignCoordinates(f, numCoords, coords)
}
func (f *Face) SfntTableInfo(tableIndex uint, tag, length *ULong) Error {
	return sfntTableInfo(f, tableIndex, tag, length)
}
func (f *Face) TrueTypeGXFree(table *byte) { trueTypeGXFree(f, table) }
func (f *Face) TrueTypeGXValidate(validationFlags uint, tables [VALIDATE_GX_LENGTH]*byte, tableLength uint) Error {
	return trueTypeGXValidate(f, validationFlags, tables, tableLength)
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
	glyphCopy      func(g *Glyph, target **Glyph) Error
	glyphTransform func(g *Glyph, matrix *Matrix, delta *Vector) Error
	glyphGetCBox   func(g *Glyph, bboxMode uint, acbox *BBox)
	doneGlyph      func(g *Glyph)
)

func (g *Glyph) Copy(target **Glyph) Error { return glyphCopy(g, target) }
func (g *Glyph) Transform(matrix *Matrix, delta *Vector) Error {
	return glyphTransform(g, matrix, delta)
}
func (g *Glyph) GetCBox(bboxMode uint, acbox *BBox) { glyphGetCBox(g, bboxMode, acbox) }
func (g *Glyph) Done()                              { doneGlyph(g) }

type (
	GlyphCopy func(source, target *Glyph) Error

	GlyphDone func(glyph *Glyph)

	GlyphGetBBox func(glyph *Glyph, abbox *BBox)

	GlyphInit func(glyph *Glyph, slot *GlyphSlot) Error

	GlyphPrepare func(glyph *Glyph, slot *GlyphSlot) Error

	GlyphTransform func(
		glyph *Glyph, matrix *Matrix, delta *Vector)
)

type GlyphClass struct {
	GlyphSize      Long
	GlyphFormat    GlyphFormat
	GlyphInit      GlyphInit
	GlyphDone      GlyphDone
	GlyphCopy      GlyphCopy
	GlyphTransform GlyphTransform
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
	getGlyph           func(g *GlyphSlot, aglyph **Glyph) Error
	getSubGlyphInfo    func(g *GlyphSlot, subIndex uint, index *int, flags *uint, arg1, arg2 *int, transform *Matrix) Error
	glyphSlotEmbolden  func(g *GlyphSlot)
	glyphSlotOblique   func(g *GlyphSlot)
	glyphSlotOwnBitmap func(slot *GlyphSlot) Error

	renderGlyph func(slot *GlyphSlot, renderMode RenderMode) Error
)

func (g *GlyphSlot) Get(aglyph **Glyph) Error { return getGlyph(g, aglyph) }
func (g *GlyphSlot) Embolden()                { glyphSlotEmbolden(g) }
func (g *GlyphSlot) Oblique()                 { glyphSlotOblique(g) }
func (g *GlyphSlot) Info(subIndex uint, index *int, flags *uint, arg1, arg2 *int, transform *Matrix) Error {
	return getSubGlyphInfo(g, subIndex, index, flags, arg1, arg2, transform)
}
func (g *GlyphSlot) OwnBitmap() Error                        { return glyphSlotOwnBitmap(g) }
func (s *GlyphSlot) RenderGlyph(renderMode RenderMode) Error { return renderGlyph(s, renderMode) }

var (
	imageCacheLookup       func(i *ImageCache, t *ImageType, gindex uint, aglyph **Glyph, anode **Node) Error
	imageCacheLookupScaler func(i *ImageCache, scaler *Scaler, loadFlags ULong, gindex uint, aglyph **Glyph, anode **Node) Error
)

func (i *ImageCache) Lookup(t *ImageType, gindex uint, aglyph **Glyph, anode **Node) Error {
	return imageCacheLookup(i, t, gindex, aglyph, anode)
}
func (i *ImageCache) LookupScaler(scaler *Scaler, loadFlags ULong, gindex uint, aglyph **Glyph, anode **Node) Error {
	return imageCacheLookupScaler(i, scaler, loadFlags, gindex, aglyph, anode)
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

	librarySetLcdFilter        func(l *Library, filter LcdFilter) Error
	librarySetLcdFilterWeights func(l *Library, weights *T.UnsignedChar) Error
	libraryVersion             func(l *Library, amajor, aminor, apatch *int)

	addDefaultModules     func(l *Library)
	addModule             func(l *Library, clazz *ModuleClass) Error
	doneFreeType          func(l *Library) Error
	doneLibrary           func(l *Library) Error
	getModule             func(l *Library, moduleName string) *Module
	getRenderer           func(l *Library, format GlyphFormat) *Renderer
	getTrueTypeEngineType func(l *Library) TrueTypeEngineType
	newFace               func(l *Library, filepathname string, faceIndex Long, aface **Face) Error
	newMemoryFace         func(l *Library, fileBase *byte, fileSize, faceIndex Long, aface **Face) Error
	openFace              func(l *Library, args *OpenArgs, faceIndex Long, aface **Face) Error
	referenceLibrary      func(l *Library) Error
	removeModule          func(l *Library, module *Module) Error
	setDebugHook          func(l *Library, hookIndex uint, debugHook DebugHookFunc)
	setRenderer           func(l *Library, renderer *Renderer, numParams uint, parameters *Parameter) Error

	strokerNew func(l *Library, astroker **Stroker) Error
	managerNew func(l *Library, maxFaces, maxSizes uint, maxBytes ULong, requester FaceRequester, reqData *Void, amanager **Manager) Error
	outlineNew func(l *Library, numPoints uint, numContours int, anoutline **Outline) Error
)

func (l *Library) SetLcdFilter(filter LcdFilter) Error { return librarySetLcdFilter(l, filter) }
func (l *Library) SetLcdFilterWeights(weights *T.UnsignedChar) Error {
	return librarySetLcdFilterWeights(l, weights)
}
func (l *Library) Version(amajor, aminor, apatch *int) { libraryVersion(l, amajor, aminor, apatch) }

func (l *Library) AddDefaultModules()                 { addDefaultModules(l) }
func (l *Library) AddModule(clazz *ModuleClass) Error { return addModule(l, clazz) }
func (l *Library) DoneFreeType() Error                { return doneFreeType(l) }
func (l *Library) DoneLibrary() Error                 { return doneLibrary(l) }
func (l *Library) ManagerNew(maxFaces, maxSizes uint, maxBytes ULong, requester FaceRequester, reqData *Void, amanager **Manager) Error {
	return managerNew(l, maxFaces, maxSizes, maxBytes, requester, reqData, amanager)
}
func (l *Library) Module(moduleName string) *Module { return getModule(l, moduleName) }
func (l *Library) NewFace(filepathname string, faceIndex Long, aface **Face) Error {
	return newFace(l, filepathname, faceIndex, aface)
}
func (l *Library) NewMemoryFace(fileBase *byte, fileSize, faceIndex Long, aface **Face) Error {
	return newMemoryFace(l, fileBase, fileSize, faceIndex, aface)
}
func (l *Library) OpenFace(args *OpenArgs, faceIndex Long, aface **Face) Error {
	return openFace(l, args, faceIndex, aface)
}
func (l *Library) Renderer(format GlyphFormat) *Renderer { return getRenderer(l, format) }
func (l *Library) ReferenceLibrary() Error               { return referenceLibrary(l) }
func (l *Library) RemoveModule(module *Module) Error     { return removeModule(l, module) }
func (l *Library) SetDebugHook(hookIndex uint, debugHook DebugHookFunc) {
	setDebugHook(l, hookIndex, debugHook)
}
func (l *Library) SetRenderer(renderer *Renderer, numParams uint, parameters *Parameter) Error {
	return setRenderer(l, renderer, numParams, parameters)
}
func (l *Library) StrokerNew(astroker **Stroker) Error    { return strokerNew(l, astroker) }
func (l *Library) TrueTypeEngineType() TrueTypeEngineType { return getTrueTypeEngineType(l) }
func (l *Library) OutlineNew(numPoints uint, numContours int, anoutline **Outline) Error {
	return outlineNew(l, numPoints, numContours, anoutline)
}

type (
	List struct{ Head, Tail *ListNode }

	ListDestructor func(memory *Memory, data, user *Void)

	ListIterator func(node *ListNode, user *Void) Error
)

var (
	listAdd      func(l *List, node *ListNode)
	listFinalize func(l *List, destroy ListDestructor, memory *Memory, user *T.Void)
	listFind     func(l *List, data *T.Void) *ListNode
	listInsert   func(l *List, node *ListNode)
	listIterate  func(l *List, iterator ListIterator, user *T.Void) Error
	listRemove   func(l *List, node *ListNode)
	listUp       func(l *List, node *ListNode)
)

func (l *List) Add(node *ListNode) { listAdd(l, node) }
func (l *List) Finalize(destroy ListDestructor, memory *Memory, user *T.Void) {
	listFinalize(l, destroy, memory, user)
}
func (l *List) Find(data *T.Void) *ListNode { return listFind(l, data) }
func (l *List) Insert(node *ListNode)       { listInsert(l, node) }
func (l *List) Iterate(iterator ListIterator, user *T.Void) Error {
	return listIterate(l, iterator, user)
}
func (l *List) Remove(node *ListNode) { listRemove(l, node) }
func (l *List) Up(node *ListNode)     { listUp(l, node) }

type ListNode struct {
	Prev, Next *ListNode
	Data       *Void
}

type Manager struct{}

var (
	managerDone         func(m *Manager)
	managerLookupFace   func(m *Manager, faceId FaceID, aface **Face) Error
	managerLookupSize   func(m *Manager, scaler *Scaler, asize **Size) Error
	managerRemoveFaceID func(m *Manager, faceId FaceID)
	managerReset        func(m *Manager)

	cMapCacheNew  func(m *Manager, acache **CMapCache) Error
	imageCacheNew func(m *Manager, acache **ImageCache) Error
	sBitCacheNew  func(m *Manager, acache **SBitCache) Error
)

func (m *Manager) Done() { managerDone(m) }
func (m *Manager) LookupFace(faceId FaceID, aface **Face) Error {
	return managerLookupFace(m, faceId, aface)
}
func (m *Manager) LookupSize(scaler *Scaler, asize **Size) Error {
	return managerLookupSize(m, scaler, asize)
}
func (m *Manager) RemoveFaceID(faceId FaceID) { managerRemoveFaceID(m, faceId) }
func (m *Manager) Reset()                     { managerReset(m) }

func (m *Manager) CMapCacheNew(acache **CMapCache) Error   { return cMapCacheNew(m, acache) }
func (m *Manager) ImageCacheNew(acache **ImageCache) Error { return imageCacheNew(m, acache) }
func (m *Manager) SBitCacheNew(acache **SBitCache) Error   { return sBitCacheNew(m, acache) }

type Matrix struct{ XX, XY, YX, YY Fixed }

var (
	matrixMultiply func(a, b *Matrix)
	matrixInvert   func(matrix *Matrix) Error
)

func (m *Matrix) Multiply(m2 *Matrix) { matrixMultiply(m, m2) }
func (m *Matrix) Invert() Error       { return matrixInvert(m) }

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
	newLibrary         func(m *Memory, alibrary **Library) Error
	outlineNewInternal func(m *Memory, numPoints uint, numContours int, anoutline **Outline) Error
)

func (m *Memory) NewLibrary(alibrary **Library) Error { return newLibrary(m, alibrary) }
func (m *Memory) OutlineNewInternal(numPoints uint, numContours int, anoutline **Outline) Error {
	return outlineNewInternal(m, numPoints, numContours, anoutline)
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

type MultiMaster struct {
	NumAxis    uint
	NumDesigns uint
	Axis       [4]MMAxis
}

type Node struct{}

var nodeUnref func(node *Node, manager *Manager)

func (n *Node) Unref(manager *Manager) { nodeUnref(n, manager) }

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
	outlineCheck            func(o *Outline) Error
	outlineCopy             func(o *Outline, target *Outline) Error
	outlineDecompose        func(o *Outline, funcInterface *OutlineFuncs, user *T.Void) Error
	outlineEmbolden         func(o *Outline, strength Pos) Error
	outlineGetBBox          func(o *Outline, abbox *BBox) Error
	outlineGetCBox          func(o *Outline, acbox *BBox)
	outlineGetInsideBorder  func(o *Outline) StrokerBorder
	outlineGetOrientation   func(o *Outline) Orientation
	outlineGetOutsideBorder func(o *Outline) StrokerBorder
	outlineReverse          func(o *Outline)
	outlineTransform        func(o *Outline, matrix *Matrix)
	outlineTranslate        func(o *Outline, xOffset, yOffset Pos)

	outlineDoneInternal func(memory *Memory, outline *Outline) Error
	outlineGetBitmap    func(library *Library, outline *Outline, abitmap *Bitmap) Error
	outlineRender       func(library *Library, outline *Outline, params *RasterParams) Error
	outlineDone         func(library *Library, outline *Outline) Error
)

func (o *Outline) Check() Error               { return outlineCheck(o) }
func (o *Outline) Copy(target *Outline) Error { return outlineCopy(o, target) }
func (o *Outline) Decompose(funcInterface *OutlineFuncs, user *T.Void) Error {
	return outlineDecompose(o, funcInterface, user)
}
func (o *Outline) Embolden(strength Pos) Error     { return outlineEmbolden(o, strength) }
func (o *Outline) GetBBox(abbox *BBox) Error       { return outlineGetBBox(o, abbox) }
func (o *Outline) GetCBox(acbox *BBox)             { outlineGetCBox(o, acbox) }
func (o *Outline) GetInsideBorder() StrokerBorder  { return outlineGetInsideBorder(o) }
func (o *Outline) GetOrientation() Orientation     { return outlineGetOrientation(o) }
func (o *Outline) GetOutsideBorder() StrokerBorder { return outlineGetOutsideBorder(o) }
func (o *Outline) Reverse()                        { outlineReverse(o) }
func (o *Outline) Transform(matrix *Matrix)        { outlineTransform(o, matrix) }
func (o *Outline) Translate(xOffset, yOffset Pos)  { outlineTranslate(o, xOffset, yOffset) }

func (o *Outline) DoneInternal(memory *Memory) Error { return outlineDoneInternal(memory, o) }
func (o *Outline) GetBitmap(library *Library, abitmap *Bitmap) Error {
	return outlineGetBitmap(library, o, abitmap)
}
func (o *Outline) Render(library *Library, params *RasterParams) Error {
	return outlineRender(library, o, params)
}
func (o *Outline) Done(library *Library) Error { return outlineDone(library, o) }

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
	sBitCacheLookup       func(s *SBitCache, t *ImageType, gindex uint, sbit **SBit, anode **Node) Error
	sBitCacheLookupScaler func(s *SBitCache, scaler *Scaler, loadFlags ULong, gindex uint, sbit **SBit, anode **Node) Error
)

func (s *SBitCache) Lookup(t *ImageType, gindex uint, sbit **SBit, anode **Node) Error {
	return sBitCacheLookup(s, t, gindex, sbit, anode)
}
func (s *SBitCache) LookupScaler(scaler *Scaler, loadFlags ULong, gindex uint, sbit **SBit, anode **Node) Error {
	return sBitCacheLookupScaler(s, scaler, loadFlags, gindex, sbit, anode)
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
	doneSize     func(size *Size) Error
	activateSize func(size *Size) Error
)

func (s *Size) Done() Error     { return doneSize(s) }
func (s *Size) Activate() Error { return activateSize(s) }

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
	streamOpenGzip func(stream, source *Stream) Error
	streamOpenLZW  func(stream, source *Stream) Error
)

func (s *Stream) OpenGzip(source *Stream) Error { return streamOpenGzip(s, source) }
func (s *Stream) OpenLZW(source *Stream) Error  { return streamOpenLZW(s, source) }

type StrokerBorder Enum

const (
	STROKER_BORDER_LEFT StrokerBorder = iota
	STROKER_BORDER_RIGHT
)

type Stroker struct{}

var (
	strokerBeginSubPath    func(s *Stroker, to *Vector, open Bool) Error
	strokerConicTo         func(s *Stroker, control, to *Vector) Error
	strokerCubicTo         func(s *Stroker, control1, control2, to *Vector) Error
	strokerDone            func(s *Stroker)
	strokerEndSubPath      func(s *Stroker) Error
	strokerExport          func(s *Stroker, outline *Outline)
	strokerExportBorder    func(s *Stroker, border StrokerBorder, outline *Outline)
	strokerGetBorderCounts func(s *Stroker, border StrokerBorder, anumPoints, anum_contours *uint) Error
	strokerGetCounts       func(s *Stroker, anumPoints, anum_contours *uint) Error
	strokerLineTo          func(s *Stroker, to *Vector) Error
	strokerParseOutline    func(s *Stroker, outline *Outline, opened Bool) Error
	strokerRewind          func(s *Stroker)
	strokerSet             func(s *Stroker, radius Fixed, lineCap StrokerLineCap, lineJoin StrokerLineJoin, miterLimit Fixed)
)

func (s *Stroker) BeginSubPath(to *Vector, open Bool) Error { return strokerBeginSubPath(s, to, open) }
func (s *Stroker) ConicTo(control, to *Vector) Error        { return strokerConicTo(s, control, to) }
func (s *Stroker) CubicTo(control1, control2, to *Vector) Error {
	return strokerCubicTo(s, control1, control2, to)
}
func (s *Stroker) Done()                   { strokerDone(s) }
func (s *Stroker) EndSubPath() Error       { return strokerEndSubPath(s) }
func (s *Stroker) Export(outline *Outline) { strokerExport(s, outline) }
func (s *Stroker) ExportBorder(border StrokerBorder, outline *Outline) {
	strokerExportBorder(s, border, outline)
}
func (s *Stroker) GetBorderCounts(border StrokerBorder, anumPoints, anum_contours *uint) Error {
	return strokerGetBorderCounts(s, border, anumPoints, anum_contours)
}
func (s *Stroker) GetCounts(anumPoints, anum_contours *uint) Error {
	return strokerGetCounts(s, anumPoints, anum_contours)
}
func (s *Stroker) LineTo(to *Vector) Error { return strokerLineTo(s, to) }
func (s *Stroker) ParseOutline(outline *Outline, opened Bool) Error {
	return strokerParseOutline(s, outline, opened)
}
func (s *Stroker) Rewind() { strokerRewind(s) }
func (s *Stroker) Set(radius Fixed, lineCap StrokerLineCap, lineJoin StrokerLineJoin, miterLimit Fixed) {
	strokerSet(s, radius, lineCap, lineJoin, miterLimit)
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
	vectorFromPolar func(v *Vector, length Fixed, angle Angle)
	vectorLength    func(v *Vector) Fixed
	vectorPolarize  func(v *Vector, length *Fixed, angle *Angle)
	vectorRotate    func(v *Vector, angle Angle)
	vectorTransform func(v *Vector, matrix *Matrix)
	vectorUnit      func(v *Vector, angle Angle)
)

func (v *Vector) FromPolar(length Fixed, angle Angle)  { vectorFromPolar(v, length, angle) }
func (v *Vector) Length() Fixed                        { return vectorLength(v) }
func (v *Vector) Polarize(length *Fixed, angle *Angle) { vectorPolarize(v, length, angle) }
func (v *Vector) Rotate(angle Angle)                   { vectorRotate(v, angle) }
func (v *Vector) Transform(matrix *Matrix)             { vectorTransform(v, matrix) }
func (v *Vector) Unit(angle Angle)                     { vectorUnit(v, angle) }

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
	{"FTC_CMapCache_Lookup", &cMapCacheLookup},
	{"FTC_CMapCache_New", &cMapCacheNew},
	{"FTC_ImageCache_Lookup", &imageCacheLookup},
	{"FTC_ImageCache_LookupScaler", &imageCacheLookupScaler},
	{"FTC_ImageCache_New", &imageCacheNew},
	{"FTC_Manager_Done", &managerDone},
	{"FTC_Manager_LookupFace", &managerLookupFace},
	{"FTC_Manager_LookupSize", &managerLookupSize},
	{"FTC_Manager_Lookup_Face", &managerLookupFace},
	{"FTC_Manager_Lookup_Size", &managerLookupSize},
	{"FTC_Manager_New", &managerNew},
	{"FTC_Manager_RemoveFaceID", &managerRemoveFaceID},
	{"FTC_Manager_Reset", &managerReset},
	{"FTC_Node_Unref", &nodeUnref},
	{"FTC_SBitCache_Lookup", &sBitCacheLookup},
	{"FTC_SBitCache_LookupScaler", &sBitCacheLookupScaler},
	{"FTC_SBitCache_New", &sBitCacheNew},
	{"FT_Activate_Size", &activateSize},
	{"FT_Add_Default_Modules", &addDefaultModules},
	{"FT_Add_Module", &addModule},
	{"FT_Angle_Diff", &AngleDiff},
	{"FT_Atan2", &Atan2},
	{"FT_Attach_File", &attachFile},
	{"FT_Attach_Stream", &attachStream},
	{"FT_Bitmap_Convert", &bitmapConvert},
	{"FT_Bitmap_Copy", &bitmapCopy},
	{"FT_Bitmap_Done", &bitmapDone},
	{"FT_Bitmap_Embolden", &bitmapEmbolden},
	{"FT_Bitmap_New", &BitmapNew},
	{"FT_CeilFix", &CeilFix},
	{"FT_ClassicKern_Free", &classicKernFree},
	{"FT_ClassicKern_Validate", &classicKernValidate},
	{"FT_Cos", &Cos},
	{"FT_DivFix", &DivFix},
	{"FT_Done_Face", &doneFace},
	{"FT_Done_FreeType", &doneFreeType},
	{"FT_Done_Glyph", &doneGlyph},
	{"FT_Done_Library", &doneLibrary},
	{"FT_Done_Size", &doneSize},
	{"FT_Face_CheckTrueTypePatents", &faceCheckTrueTypePatents},
	{"FT_Face_GetCharVariantIndex", &faceGetCharVariantIndex},
	{"FT_Face_GetCharVariantIsDefault", &faceGetCharVariantIsDefault},
	{"FT_Face_GetCharsOfVariant", &faceGetCharsOfVariant},
	{"FT_Face_GetVariantSelectors", &faceGetVariantSelectors},
	{"FT_Face_GetVariantsOfChar", &faceGetVariantsOfChar},
	{"FT_Face_SetUnpatentedHinting", &faceSetUnpatentedHinting},
	{"FT_FloorFix", &FloorFix},
	{"FT_Get_Advance", &getAdvance},
	{"FT_Get_Advances", &getAdvances},
	{"FT_Get_BDF_Charset_ID", &getBDFCharsetID},
	{"FT_Get_BDF_Property", &getBDFProperty},
	{"FT_Get_CID_From_Glyph_Index", &getCIDFromGlyphIndex},
	{"FT_Get_CID_Is_Internally_CID_Keyed", &getCIDIsInternallyCIDKeyed},
	{"FT_Get_CID_Registry_Ordering_Supplement", &getCIDRegistryOrderingSupplement},
	{"FT_Get_CMap_Format", &getCMapFormat},
	{"FT_Get_CMap_Language_ID", &getCMapLanguageID},
	{"FT_Get_Char_Index", &getCharIndex},
	{"FT_Get_Charmap_Index", &getCharmapIndex},
	{"FT_Get_FSType_Flags", &getFSTypeFlags},
	{"FT_Get_First_Char", &getFirstChar},
	{"FT_Get_Gasp", &getGasp},
	{"FT_Get_Glyph", &getGlyph},
	{"FT_Get_Glyph_Name", &getGlyphName},
	{"FT_Get_Kerning", &getKerning},
	{"FT_Get_MM_Var", &getMMVar},
	{"FT_Get_Module", &getModule},
	{"FT_Get_Multi_Master", &getMultiMaster},
	{"FT_Get_Name_Index", &getNameIndex},
	{"FT_Get_Next_Char", &getNextChar},
	{"FT_Get_PFR_Advance", &getPFRAdvance},
	{"FT_Get_PFR_Kerning", &getPFRKerning},
	{"FT_Get_PFR_Metrics", &getPFRMetrics},
	{"FT_Get_PS_Font_Info", &getPSFontInfo},
	{"FT_Get_PS_Font_Private", &getPSFontPrivate},
	{"FT_Get_Postscript_Name", &getPostscriptName},
	{"FT_Get_Renderer", &getRenderer},
	{"FT_Get_Sfnt_Name", &getSfntName},
	{"FT_Get_Sfnt_Name_Count", &getSfntNameCount},
	{"FT_Get_Sfnt_Table", &getSfntTable},
	{"FT_Get_SubGlyph_Info", &getSubGlyphInfo},
	{"FT_Get_Track_Kerning", &getTrackKerning},
	{"FT_Get_TrueType_Engine_Type", &getTrueTypeEngineType},
	{"FT_Get_WinFNT_Header", &getWinFNTHeader},
	{"FT_Get_X11_Font_Format", &getX11FontFormat},
	{"FT_GlyphSlot_Embolden", &glyphSlotEmbolden},
	{"FT_GlyphSlot_Oblique", &glyphSlotOblique},
	{"FT_GlyphSlot_Own_Bitmap", &glyphSlotOwnBitmap},
	{"FT_Glyph_Copy", &glyphCopy},
	{"FT_Glyph_Get_CBox", &glyphGetCBox},
	{"FT_Glyph_Stroke", &GlyphStroke},
	{"FT_Glyph_StrokeBorder", &GlyphStrokeBorder},
	{"FT_Glyph_To_Bitmap", &GlyphToBitmap},
	{"FT_Glyph_Transform", &glyphTransform},
	{"FT_Has_PS_Glyph_Names", &hasPSGlyphNames},
	{"FT_Init_FreeType", &InitFreeType},
	{"FT_Library_SetLcdFilter", &librarySetLcdFilter},
	{"FT_Library_SetLcdFilterWeights", &librarySetLcdFilterWeights},
	{"FT_Library_Version", &libraryVersion},
	{"FT_List_Add", &listAdd},
	{"FT_List_Finalize", &listFinalize},
	{"FT_List_Find", &listFind},
	{"FT_List_Insert", &listInsert},
	{"FT_List_Iterate", &listIterate},
	{"FT_List_Remove", &listRemove},
	{"FT_List_Up", &listUp},
	{"FT_Load_Char", &loadChar},
	{"FT_Load_Glyph", &loadGlyph},
	{"FT_Load_Sfnt_Table", &loadSfntTable},
	{"FT_Matrix_Invert", &matrixInvert},
	{"FT_Matrix_Multiply", &matrixMultiply},
	{"FT_MulDiv", &MulDiv},
	{"FT_MulFix", &MulFix},
	{"FT_New_Face", &newFace},
	{"FT_New_Library", &newLibrary},
	{"FT_New_Memory_Face", &newMemoryFace},
	{"FT_New_Size", &newSize},
	{"FT_OpenType_Free", &openTypeFree},
	{"FT_OpenType_Validate", &openTypeValidate},
	{"FT_Open_Face", &openFace},
	{"FT_Outline_Check", &outlineCheck},
	{"FT_Outline_Copy", &outlineCopy},
	{"FT_Outline_Decompose", &outlineDecompose},
	{"FT_Outline_Done", &outlineDone},
	{"FT_Outline_Done_Internal", &outlineDoneInternal},
	{"FT_Outline_Embolden", &outlineEmbolden},
	{"FT_Outline_GetInsideBorder", &outlineGetInsideBorder},
	{"FT_Outline_GetOutsideBorder", &outlineGetOutsideBorder},
	{"FT_Outline_Get_BBox", &outlineGetBBox},
	{"FT_Outline_Get_Bitmap", &outlineGetBitmap},
	{"FT_Outline_Get_CBox", &outlineGetCBox},
	{"FT_Outline_Get_Orientation", &outlineGetOrientation},
	{"FT_Outline_New", &outlineNew},
	{"FT_Outline_New_Internal", &outlineNewInternal},
	{"FT_Outline_Render", &outlineRender},
	{"FT_Outline_Reverse", &outlineReverse},
	{"FT_Outline_Transform", &outlineTransform},
	{"FT_Outline_Translate", &outlineTranslate},
	{"FT_Reference_Face", &referenceFace},
	{"FT_Reference_Library", &referenceLibrary},
	{"FT_Remove_Module", &removeModule},
	{"FT_Render_Glyph", &renderGlyph},
	{"FT_Request_Size", &requestSize},
	{"FT_RoundFix", &RoundFix},
	{"FT_Select_Charmap", &selectCharmap},
	{"FT_Select_Size", &selectSize},
	{"FT_Set_Char_Size", &setCharSize},
	{"FT_Set_Charmap", &setCharmap},
	{"FT_Set_Debug_Hook", &setDebugHook},
	{"FT_Set_MM_Blend_Coordinates", &setMMBlendCoordinates},
	{"FT_Set_MM_Design_Coordinates", &setMMDesignCoordinates},
	{"FT_Set_Pixel_Sizes", &setPixelSizes},
	{"FT_Set_Renderer", &setRenderer},
	{"FT_Set_Transform", &setTransform},
	{"FT_Set_Var_Blend_Coordinates", &setVarBlendCoordinates},
	{"FT_Set_Var_Design_Coordinates", &setVarDesignCoordinates},
	{"FT_Sfnt_Table_Info", &sfntTableInfo},
	{"FT_Sin", &Sin},
	{"FT_Stream_OpenGzip", &streamOpenGzip},
	{"FT_Stream_OpenLZW", &streamOpenLZW},
	{"FT_Stroker_BeginSubPath", &strokerBeginSubPath},
	{"FT_Stroker_ConicTo", &strokerConicTo},
	{"FT_Stroker_CubicTo", &strokerCubicTo},
	{"FT_Stroker_Done", &strokerDone},
	{"FT_Stroker_EndSubPath", &strokerEndSubPath},
	{"FT_Stroker_Export", &strokerExport},
	{"FT_Stroker_ExportBorder", &strokerExportBorder},
	{"FT_Stroker_GetBorderCounts", &strokerGetBorderCounts},
	{"FT_Stroker_GetCounts", &strokerGetCounts},
	{"FT_Stroker_LineTo", &strokerLineTo},
	{"FT_Stroker_New", &strokerNew},
	{"FT_Stroker_ParseOutline", &strokerParseOutline},
	{"FT_Stroker_Rewind", &strokerRewind},
	{"FT_Stroker_Set", &strokerSet},
	{"FT_Tan", &Tan},
	{"FT_TrueTypeGX_Free", &trueTypeGXFree},
	{"FT_TrueTypeGX_Validate", &trueTypeGXValidate},
	{"FT_Vector_From_Polar", &vectorFromPolar},
	{"FT_Vector_Length", &vectorLength},
	{"FT_Vector_Polarize", &vectorPolarize},
	{"FT_Vector_Rotate", &vectorRotate},
	{"FT_Vector_Transform", &vectorTransform},
	{"FT_Vector_Unit", &vectorUnit},
	// Undocumented {"TT_New_Context", &TTNewContext},
	// Undocumented {"TT_RunIns", &TTRunIns},
}
