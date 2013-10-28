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

var (
	NewFace func(
		library T.FT_Library,
		filepathname string,
		faceIndex T.FT_Long,
		aface *T.FTFace) T.FT_Error

	NewMemoryFace func(
		library T.FT_Library,
		fileBase *T.FT_Byte,
		fileSize T.FT_Long,
		faceIndex T.FT_Long,
		aface *T.FTFace) T.FT_Error

	OpenFace func(
		library T.FT_Library,
		args *T.FT_Open_Args,
		faceIndex T.FT_Long,
		aface *T.FTFace) T.FT_Error

	AttachFile func(
		face T.FTFace,
		filepathname string) T.FT_Error

	AttachStream func(
		face T.FTFace,
		parameters *T.FT_Open_Args) T.FT_Error

	ReferenceFace func(
		face T.FTFace) T.FT_Error

	DoneFace func(
		face T.FTFace) T.FT_Error

	SelectSize func(
		face T.FTFace,
		strikeIndex int) T.FT_Error

	RequestSize func(
		face T.FTFace,
		req T.FT_Size_Request) T.FT_Error

	SetCharSize func(
		face T.FTFace,
		charWidth T.FT_F26Dot6,
		charHeight T.FT_F26Dot6,
		horzResolution uint,
		vertResolution uint) T.FT_Error

	SetPixelSizes func(
		face T.FTFace,
		pixelWidth uint,
		pixelHeight uint) T.FT_Error

	LoadGlyph func(
		face T.FTFace,
		glyphIndex uint,
		loadFlags T.FT_Int32) T.FT_Error

	LoadChar func(
		face T.FTFace,
		charCode T.FT_ULong,
		loadFlags T.FT_Int32) T.FT_Error

	SetTransform func(
		face T.FTFace,
		matrix *T.FT_Matrix,
		delta *T.FT_Vector)

	GetKerning func(
		face T.FTFace,
		leftGlyph uint,
		rightGlyph uint,
		kernMode uint,
		akerning *T.FT_Vector) T.FT_Error

	GetTrackKerning func(
		face T.FTFace,
		pointSize T.FT_Fixed,
		degree int,
		akerning *T.FT_Fixed) T.FT_Error

	GetGlyphName func(
		face T.FTFace,
		glyphIndex uint,
		buffer T.FT_Pointer,
		bufferMax uint) T.FT_Error

	GetPostscriptName func(
		face T.FTFace) string

	SelectCharmap func(
		face T.FTFace,
		encoding T.FT_Encoding) T.FT_Error

	SetCharmap func(
		face T.FTFace,
		charmap T.FT_CharMap) T.FT_Error

	GetCharmapIndex func(
		charmap T.FT_CharMap) int

	GetCharIndex func(
		face T.FTFace,
		charcode T.FT_ULong) uint

	GetFirstChar func(
		face T.FTFace,
		agindex *uint) T.FT_ULong

	GetNextChar func(
		face T.FTFace,
		charCode T.FT_ULong,
		agindex *uint) T.FT_ULong

	GetNameIndex func(
		face T.FTFace,
		glyphName *T.FT_String) uint

	GetSubGlyphInfo func(
		glyph T.FT_GlyphSlot,
		subIndex uint,
		index *int,
		flags *uint,
		arg1 *int,
		arg2 *int,
		transform *T.FT_Matrix) T.FT_Error

	GetFSTypeFlags func(face T.FTFace) uint16

	FaceGetCharVariantIndex func(
		face T.FTFace,
		charcode T.FT_ULong,
		variantSelector T.FT_ULong) uint

	FaceGetCharVariantIsDefault func(
		face T.FTFace,
		charcode T.FT_ULong,
		variantSelector T.FT_ULong) int

	FaceGetVariantSelectors func(
		face T.FTFace) *T.FT_UInt32

	FaceGetVariantsOfChar func(
		face T.FTFace,
		charcode T.FT_ULong) *T.FT_UInt32

	FaceGetCharsOfVariant func(
		face T.FTFace,
		variantSelector T.FT_ULong) *T.FT_UInt32

	MulDiv func(
		a T.FT_Long,
		b T.FT_Long,
		c T.FT_Long) T.FT_Long

	DivFix func(
		a T.FT_Long,
		b T.FT_Long) T.FT_Long

	RoundFix func(
		a T.FT_Fixed) T.FT_Fixed

	CeilFix func(
		a T.FT_Fixed) T.FT_Fixed

	FloorFix func(
		a T.FT_Fixed) T.FT_Fixed

	VectorTransform func(
		vec *T.FT_Vector,
		matrix *T.FT_Matrix)

	LibraryVersion func(
		library T.FT_Library,
		amajor *int,
		aminor *int,
		apatch *int)

	FaceCheckTrueTypePatents func(
		face T.FTFace) T.FT_Bool

	FaceSetUnpatentedHinting func(
		face T.FTFace,
		value T.FT_Bool) T.FT_Bool

	HasPSGlyphNames func(
		face T.FTFace) int

	GetPSFontInfo func(
		face T.FTFace,
		afontInfo T.PS_FontInfo) T.FT_Error

	GetPSFontPrivate func(
		face T.FTFace,
		afontPrivate T.PS_Private) T.FT_Error

	GetMultiMaster func(
		face T.FTFace,
		amaster *T.FT_Multi_Master) T.FT_Error

	GetMMVar func(
		face T.FTFace,
		amaster **T.FT_MM_Var) T.FT_Error

	SetMMDesignCoordinates func(
		face T.FTFace,
		numCoords uint,
		coords *T.FT_Long) T.FT_Error

	SetVarDesignCoordinates func(
		face T.FTFace,
		numCoords uint,
		coords *T.FT_Fixed) T.FT_Error

	SetMMBlendCoordinates func(
		face T.FTFace,
		numCoords uint,
		coords *T.FT_Fixed) T.FT_Error

	SetVarBlendCoordinates func(
		face T.FTFace,
		numCoords uint,
		coords *T.FT_Fixed) T.FT_Error

	GetGlyph func(
		slot T.FT_GlyphSlot,
		aglyph *T.FT_Glyph) T.FT_Error

	GlyphCopy func(
		source T.FT_Glyph,
		target *T.FT_Glyph) T.FT_Error

	GlyphTransform func(
		glyph T.FT_Glyph,
		matrix *T.FT_Matrix,
		delta *T.FT_Vector) T.FT_Error

	GlyphGetCBox func(
		glyph T.FT_Glyph,
		bboxMode uint,
		acbox *T.FT_BBox)

	GlyphToBitmap func(
		theGlyph *T.FT_Glyph,
		renderMode T.FT_Render_Mode,
		origin *T.FT_Vector,
		destroy T.FT_Bool) T.FT_Error

	DoneGlyph func(
		glyph T.FT_Glyph)

	MatrixMultiply func(
		a *T.FT_Matrix,
		b *T.FT_Matrix)

	MatrixInvert func(
		matrix *T.FT_Matrix) T.FT_Error

	ManagerNew func(
		library T.FT_Library,
		maxFaces uint,
		maxSizes uint,
		maxBytes T.FT_ULong,
		requester T.FTC_Face_Requester,
		reqData T.FT_Pointer,
		amanager *T.FTC_Manager) T.FT_Error

	ManagerReset func(
		manager T.FTC_Manager)

	ManagerDone func(
		manager T.FTC_Manager)

	ManagerLookupFace func(
		manager T.FTC_Manager,
		faceId T.FTC_FaceID,
		aface *T.FTFace) T.FT_Error

	ManagerLookupSize func(
		manager T.FTC_Manager,
		scaler T.FTC_Scaler,
		asize *T.FT_Size) T.FT_Error

	NodeUnref func(
		node T.FTC_Node,
		manager T.FTC_Manager)

	ManagerRemoveFaceID func(
		manager T.FTC_Manager,
		faceId T.FTC_FaceID)

	CMapCacheNew func(
		manager T.FTC_Manager,
		acache *T.FTC_CMapCache) T.FT_Error

	CMapCacheLookup func(
		cache T.FTC_CMapCache,
		face_id T.FTC_FaceID,
		cmap_index int,
		char_code T.FT_UInt32) uint

	ImageCacheNew func(
		manager T.FTC_Manager,
		acache *T.FTC_ImageCache) T.FT_Error

	ImageCacheLookup func(
		cache T.FTC_ImageCache,
		t T.FTC_ImageType,
		gindex uint,
		aglyph *T.FT_Glyph,
		anode *T.FTC_Node) T.FT_Error

	ImageCacheLookupScaler func(
		cache T.FTC_ImageCache,
		scaler T.FTC_Scaler,
		loadFlags T.FT_ULong,
		gindex uint,
		aglyph *T.FT_Glyph,
		anode *T.FTC_Node) T.FT_Error

	SBitCacheNew func(
		manager T.FTC_Manager,
		acache *T.FTC_SBitCache) T.FT_Error

	SBitCacheLookup func(
		cache T.FTC_SBitCache,
		t T.FTC_ImageType,
		gindex uint,
		sbit *T.FTC_SBit,
		anode *T.FTC_Node) T.FT_Error

	SBitCacheLookupScaler func(
		cache T.FTC_SBitCache,
		scaler T.FTC_Scaler,
		loadFlags T.FT_ULong,
		gindex uint,
		sbit *T.FTC_SBit,
		anode *T.FTC_Node) T.FT_Error

	NewSize func(
		face T.FTFace,
		size *T.FT_Size) T.FT_Error

	DoneSize func(
		size T.FT_Size) T.FT_Error

	ActivateSize func(
		size T.FT_Size) T.FT_Error

	ListFind func(
		list T.FT_List,
		data *T.Void) T.FT_ListNode

	ListAdd func(
		list T.FT_List,
		node T.FT_ListNode)

	ListInsert func(
		list T.FT_List,
		node T.FT_ListNode)

	ListRemove func(
		list T.FT_List,
		node T.FT_ListNode)

	ListUp func(
		list T.FT_List,
		node T.FT_ListNode)

	ListIterate func(
		list T.FT_List,
		iterator T.FT_List_Iterator,
		user *T.Void) T.FT_Error

	ListFinalize func(
		list T.FT_List,
		destroy T.FT_List_Destructor,
		memory T.FT_Memory,
		user *T.Void)

	OutlineDecompose func(
		outline *T.FT_Outline,
		funcInterface *T.FT_Outline_Funcs,
		user *T.Void) T.FT_Error

	OutlineNew func(
		library T.FT_Library,
		numPoints uint,
		numContours int,
		anoutline *T.FT_Outline) T.FT_Error

	OutlineNewInternal func(
		memory T.FT_Memory,
		numPoints uint,
		numContours int,
		anoutline *T.FT_Outline) T.FT_Error

	OutlineDone func(
		library T.FT_Library,
		outline *T.FT_Outline) T.FT_Error

	OutlineDoneInternal func(
		memory T.FT_Memory,
		outline *T.FT_Outline) T.FT_Error

	OutlineCheck func(
		outline *T.FT_Outline) T.FT_Error

	OutlineGetCBox func(
		outline *T.FT_Outline,
		acbox *T.FT_BBox)

	OutlineTranslate func(
		outline *T.FT_Outline,
		xOffset T.FT_Pos,
		yOffset T.FT_Pos)

	OutlineCopy func(
		source *T.FT_Outline,
		target *T.FT_Outline) T.FT_Error

	OutlineTransform func(
		outline *T.FT_Outline,
		matrix *T.FT_Matrix)

	OutlineEmbolden func(
		outline *T.FT_Outline,
		strength T.FT_Pos) T.FT_Error

	OutlineReverse func(
		outline *T.FT_Outline)

	OutlineGetBitmap func(
		library T.FT_Library,
		outline *T.FT_Outline,
		abitmap *T.FT_Bitmap) T.FT_Error

	OutlineRender func(
		library T.FT_Library,
		outline *T.FT_Outline,
		params *T.FT_Raster_Params) T.FT_Error

	OutlineGetOrientation func(
		outline *T.FT_Outline) T.FT_Orientation

	AddModule func(
		library T.FT_Library,
		clazz *T.FT_Module_Class) T.FT_Error

	GetModule func(
		library T.FT_Library,
		moduleName string) T.FT_Module

	RemoveModule func(
		library T.FT_Library,
		module T.FT_Module) T.FT_Error

	ReferenceLibrary func(library T.FT_Library) T.FT_Error

	NewLibrary func(
		memory T.FT_Memory,
		alibrary *T.FT_Library) T.FT_Error

	DoneLibrary func(library T.FT_Library) T.FT_Error

	SetDebugHook func(
		library T.FT_Library,
		hookIndex uint,
		debugHook T.FT_DebugHook_Func)

	AddDefaultModules func(library T.FT_Library)

	GetTrueTypeEngineType func(
		library T.FT_Library) T.FT_TrueTypeEngineType

	GetRenderer func(library T.FT_Library,
		format T.FT_Glyph_Format) T.FT_Renderer

	SetRenderer func(
		library T.FT_Library,
		renderer T.FT_Renderer,
		numParams uint,
		parameters *T.FT_Parameter) T.FT_Error

	GetSfntTable func(
		face T.FTFace, tag T.FT_Sfnt_Tag) *T.Void

	LoadSfntTable func(
		face T.FTFace,
		tag T.FT_ULong,
		offset T.FT_Long,
		buffer *T.FT_Byte,
		length *T.FT_ULong) T.FT_Error

	SfntTableInfo func(
		face T.FTFace,
		tableIndex uint,
		tag, length *T.FT_ULong) T.FT_Error

	GetCMapLanguageID func(charmap T.FT_CharMap) T.FT_ULong

	GetCMapFormat func(charmap T.FT_CharMap) T.FT_Long

	GetBDFCharsetID func(face T.FTFace,
		acharsetEncoding, acharset_registry **T.Char) T.FT_Error

	GetBDFProperty func(
		face T.FTFace,
		propName string,
		aproperty *T.BDF_PropertyRec) T.FT_Error

	GetCIDRegistryOrderingSupplement func(
		face T.FTFace,
		registry, ordering **T.Char,
		supplement *int) T.FT_Error

	GetCIDIsInternallyCIDKeyed func(
		face T.FTFace,
		isCid *T.FT_Bool) T.FT_Error

	GetCIDFromGlyphIndex func(
		face T.FTFace,
		glyphIndex uint,
		cid *uint) T.FT_Error

	StreamOpenGzip func(
		stream, source T.FT_Stream) T.FT_Error

	StreamOpenLZW func(
		stream, source T.FT_Stream) T.FT_Error

	GetWinFNTHeader func(
		face T.FTFace,
		aheader *T.FT_WinFNT_HeaderRec) T.FT_Error

	BitmapNew func(abitmap *T.FT_Bitmap)

	BitmapCopy func(
		library T.FT_Library,
		source, target *T.FT_Bitmap) T.FT_Error

	BitmapEmbolden func(
		library T.FT_Library,
		bitmap *T.FT_Bitmap,
		xStrength, yStrength T.FT_Pos) T.FT_Error

	BitmapConvert func(
		library T.FT_Library,
		source, target *T.FT_Bitmap,
		alignment int) T.FT_Error

	GlyphSlotOwnBitmap func(slot T.FT_GlyphSlot) T.FT_Error

	BitmapDone func(
		library T.FT_Library, bitmap *T.FT_Bitmap) T.FT_Error

	OutlineGetBBox func(
		outline *T.FT_Outline, abbox *T.FT_BBox) T.FT_Error

	GetSfntNameCount func(face T.FTFace) uint

	GetSfntName func(
		face T.FTFace,
		idx uint,
		aname *T.FT_SfntName) T.FT_Error

	OpenTypeValidate func(
		face T.FTFace, validation_flags uint,
		BASETable, GDEF_table, GPOS_table,
		GSUBTable, JSTF_table *T.FT_Bytes) T.FT_Error

	OpenTypeFree func(face T.FTFace, table T.FT_Bytes)

	TrueTypeGXValidate func(
		face T.FTFace,
		validationFlags uint,
		tables [VALIDATE_GX_LENGTH]T.FT_Bytes,
		tableLength uint) T.FT_Error

	TrueTypeGXFree func(face T.FTFace, table T.FT_Bytes)

	ClassicKernValidate func(
		face T.FTFace,
		validation_flags uint,
		ckernTable *T.FT_Bytes) T.FT_Error

	ClassicKernFree func(
		face T.FTFace, table T.FT_Bytes)

	GetPFRMetrics func(face T.FTFace,
		aoutline_resolution, ametrics_resolution *uint,
		ametrics_x_scale, ametrics_y_scale *T.FT_Fixed) T.FT_Error

	GetPFRKerning func(
		face T.FTFace,
		left, right uint,
		avector *T.FT_Vector) T.FT_Error

	GetPFRAdvance func(
		face T.FTFace,
		gindex uint,
		aadvance *T.FT_Pos) T.FT_Error

	OutlineGetInsideBorder func(
		outline *T.FT_Outline) T.FT_StrokerBorder

	OutlineGetOutsideBorder func(
		outline *T.FT_Outline) T.FT_StrokerBorder

	StrokerNew func(
		library T.FT_Library,
		astroker *T.FT_Stroker) T.FT_Error

	StrokerSet func(
		stroker T.FT_Stroker,
		radius T.FT_Fixed,
		lineCap T.FT_Stroker_LineCap,
		lineJoin T.FT_Stroker_LineJoin,
		miterLimit T.FT_Fixed)

	StrokerRewind func(stroker T.FT_Stroker)

	StrokerParseOutline func(stroker T.FT_Stroker,
		outline *T.FT_Outline, opened T.FT_Bool) T.FT_Error

	StrokerBeginSubPath func(stroker T.FT_Stroker,
		to *T.FT_Vector, open T.FT_Bool) T.FT_Error

	StrokerEndSubPath func(stroker T.FT_Stroker) T.FT_Error

	StrokerLineTo func(
		stroker T.FT_Stroker, to *T.FT_Vector) T.FT_Error

	StrokerConicTo func(stroker T.FT_Stroker,
		control, to *T.FT_Vector) T.FT_Error

	StrokerCubicTo func(stroker T.FT_Stroker,
		control1, control2, to *T.FT_Vector) T.FT_Error

	StrokerGetBorderCounts func(
		stroker T.FT_Stroker, border T.FT_StrokerBorder,
		anumPoints, anum_contours *uint) T.FT_Error

	StrokerExportBorder func(stroker T.FT_Stroker,
		border T.FT_StrokerBorder, outline *T.FT_Outline)

	StrokerGetCounts func(stroker T.FT_Stroker,
		anumPoints, anum_contours *uint) T.FT_Error

	StrokerExport func(
		stroker T.FT_Stroker, outline *T.FT_Outline)

	StrokerDone func(stroker T.FT_Stroker)

	GlyphStroke func(pglyph *T.FT_Glyph,
		stroker T.FT_Stroker, destroy T.FT_Bool) T.FT_Error

	GlyphStrokeBorder func(
		pglyph *T.FT_Glyph, stroker T.FT_Stroker,
		inside, destroy T.FT_Bool) T.FT_Error

	GlyphSlotEmbolden func(slot T.FT_GlyphSlot)

	GlyphSlotOblique func(slot T.FT_GlyphSlot)

	GetX11FontFormat func(face T.FTFace) string

	Sin func(angle T.FT_Angle) T.FT_Fixed

	Cos func(angle T.FT_Angle) T.FT_Fixed

	Tan func(angle T.FT_Angle) T.FT_Fixed

	Atan2 func(x, y T.FT_Fixed) T.FT_Angle

	AngleDiff func(angle1, angle2 T.FT_Angle) T.FT_Angle

	VectorUnit func(vec *T.FT_Vector, angle T.FT_Angle)

	VectorRotate func(vec *T.FT_Vector, angle T.FT_Angle)

	VectorLength func(vec *T.FT_Vector) T.FT_Fixed

	VectorPolarize func(
		vec *T.FT_Vector, length *T.FT_Fixed, angle *T.FT_Angle)

	VectorFromPolar func(
		vec *T.FT_Vector, length T.FT_Fixed, angle T.FT_Angle)

	LibrarySetLcdFilter func(
		library T.FT_Library,
		filter T.FT_LcdFilter) T.FT_Error

	LibrarySetLcdFilterWeights func(
		library T.FT_Library,
		weights *T.UnsignedChar) T.FT_Error

	GetGasp func(
		face T.FTFace, ppem uint) int

	GetAdvance func(
		face T.FTFace,
		gindex uint,
		loadFlags T.FT_Int32,
		padvance *T.FT_Fixed) T.FT_Error

	GetAdvances func(
		face T.FTFace,
		start, count uint,
		loadFlags T.FT_Int32,
		padvances *T.FT_Fixed) T.FT_Error

	DoneFreeType func(library T.FT_Library) T.FT_Error

	InitFreeType func(alibrary *T.FT_Library) T.FT_Error

	MulFix func(a, b T.FT_Long) T.FT_Long

	RenderGlyph func(
		slot T.FT_GlyphSlot,
		renderMode T.FT_Render_Mode) T.FT_Error
)

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
