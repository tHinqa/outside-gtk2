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
	FT_VALIDATE_lcar_INDEX    = 9
	FT_VALIDATE_GX_LAST_INDEX = FT_VALIDATE_lcar_INDEX
	FT_VALIDATE_GX_LENGTH     = FT_VALIDATE_GX_LAST_INDEX + 1
)

var (
	FT_New_Face func(
		library T.FT_Library,
		filepathname string,
		face_index T.FT_Long,
		aface *T.FT_Face) T.FT_Error

	FT_New_Memory_Face func(
		library T.FT_Library,
		file_base *T.FT_Byte,
		file_size T.FT_Long,
		face_index T.FT_Long,
		aface *T.FT_Face) T.FT_Error

	FT_Open_Face func(
		library T.FT_Library,
		args *T.FT_Open_Args,
		face_index T.FT_Long,
		aface *T.FT_Face) T.FT_Error

	FT_Attach_File func(
		face T.FT_Face,
		filepathname string) T.FT_Error

	FT_Attach_Stream func(
		face T.FT_Face,
		parameters *T.FT_Open_Args) T.FT_Error

	FT_Reference_Face func(
		face T.FT_Face) T.FT_Error

	FT_Done_Face func(
		face T.FT_Face) T.FT_Error

	FT_Select_Size func(
		face T.FT_Face,
		strike_index int) T.FT_Error

	FT_Request_Size func(
		face T.FT_Face,
		req T.FT_Size_Request) T.FT_Error

	FT_Set_Char_Size func(
		face T.FT_Face,
		char_width T.FT_F26Dot6,
		char_height T.FT_F26Dot6,
		horz_resolution uint,
		vert_resolution uint) T.FT_Error

	FT_Set_Pixel_Sizes func(
		face T.FT_Face,
		pixel_width uint,
		pixel_height uint) T.FT_Error

	FT_Load_Glyph func(
		face T.FT_Face,
		glyph_index uint,
		load_flags T.FT_Int32) T.FT_Error

	FT_Load_Char func(
		face T.FT_Face,
		char_code T.FT_ULong,
		load_flags T.FT_Int32) T.FT_Error

	FT_Set_Transform func(
		face T.FT_Face,
		matrix *T.FT_Matrix,
		delta *T.FT_Vector)

	FT_Get_Kerning func(
		face T.FT_Face,
		left_glyph uint,
		right_glyph uint,
		kern_mode uint,
		akerning *T.FT_Vector) T.FT_Error

	FT_Get_Track_Kerning func(
		face T.FT_Face,
		point_size T.FT_Fixed,
		degree int,
		akerning *T.FT_Fixed) T.FT_Error

	FT_Get_Glyph_Name func(
		face T.FT_Face,
		glyph_index uint,
		buffer T.FT_Pointer,
		buffer_max uint) T.FT_Error

	FT_Get_Postscript_Name func(
		face T.FT_Face) string

	FT_Select_Charmap func(
		face T.FT_Face,
		encoding T.FT_Encoding) T.FT_Error

	FT_Set_Charmap func(
		face T.FT_Face,
		charmap T.FT_CharMap) T.FT_Error

	FT_Get_Charmap_Index func(
		charmap T.FT_CharMap) int

	FT_Get_Char_Index func(
		face T.FT_Face,
		charcode T.FT_ULong) uint

	FT_Get_First_Char func(
		face T.FT_Face,
		agindex *uint) T.FT_ULong

	FT_Get_Next_Char func(
		face T.FT_Face,
		char_code T.FT_ULong,
		agindex *uint) T.FT_ULong

	FT_Get_Name_Index func(
		face T.FT_Face,
		glyph_name *T.FT_String) uint

	FT_Get_SubGlyph_Info func(
		glyph T.FT_GlyphSlot,
		sub_index uint,
		p_index *int,
		p_flags *uint,
		p_arg1 *int,
		p_arg2 *int,
		p_transform *T.FT_Matrix) T.FT_Error

	FT_Get_FSType_Flags func(face T.FT_Face) uint16

	FT_Face_GetCharVariantIndex func(
		face T.FT_Face,
		charcode T.FT_ULong,
		variantSelector T.FT_ULong) uint

	FT_Face_GetCharVariantIsDefault func(
		face T.FT_Face,
		charcode T.FT_ULong,
		variantSelector T.FT_ULong) int

	FT_Face_GetVariantSelectors func(
		face T.FT_Face) *T.FT_UInt32

	FT_Face_GetVariantsOfChar func(
		face T.FT_Face,
		charcode T.FT_ULong) *T.FT_UInt32

	FT_Face_GetCharsOfVariant func(
		face T.FT_Face,
		variantSelector T.FT_ULong) *T.FT_UInt32

	FT_MulDiv func(
		a T.FT_Long,
		b T.FT_Long,
		c T.FT_Long) T.FT_Long

	FT_DivFix func(
		a T.FT_Long,
		b T.FT_Long) T.FT_Long

	FT_RoundFix func(
		a T.FT_Fixed) T.FT_Fixed

	FT_CeilFix func(
		a T.FT_Fixed) T.FT_Fixed

	FT_FloorFix func(
		a T.FT_Fixed) T.FT_Fixed

	FT_Vector_Transform func(
		vec *T.FT_Vector,
		matrix *T.FT_Matrix)

	FT_Library_Version func(
		library T.FT_Library,
		amajor *int,
		aminor *int,
		apatch *int)

	FT_Face_CheckTrueTypePatents func(
		face T.FT_Face) T.FT_Bool

	FT_Face_SetUnpatentedHinting func(
		face T.FT_Face,
		value T.FT_Bool) T.FT_Bool

	FT_Has_PS_Glyph_Names func(
		face T.FT_Face) int

	FT_Get_PS_Font_Info func(
		face T.FT_Face,
		afont_info T.PS_FontInfo) T.FT_Error

	FT_Get_PS_Font_Private func(
		face T.FT_Face,
		afont_private T.PS_Private) T.FT_Error

	FT_Get_Multi_Master func(
		face T.FT_Face,
		amaster *T.FT_Multi_Master) T.FT_Error

	FT_Get_MM_Var func(
		face T.FT_Face,
		amaster **T.FT_MM_Var) T.FT_Error

	FT_Set_MM_Design_Coordinates func(
		face T.FT_Face,
		num_coords uint,
		coords *T.FT_Long) T.FT_Error

	FT_Set_Var_Design_Coordinates func(
		face T.FT_Face,
		num_coords uint,
		coords *T.FT_Fixed) T.FT_Error

	FT_Set_MM_Blend_Coordinates func(
		face T.FT_Face,
		num_coords uint,
		coords *T.FT_Fixed) T.FT_Error

	FT_Set_Var_Blend_Coordinates func(
		face T.FT_Face,
		num_coords uint,
		coords *T.FT_Fixed) T.FT_Error

	FT_Get_Glyph func(
		slot T.FT_GlyphSlot,
		aglyph *T.FT_Glyph) T.FT_Error

	FT_Glyph_Copy func(
		source T.FT_Glyph,
		target *T.FT_Glyph) T.FT_Error

	FT_Glyph_Transform func(
		glyph T.FT_Glyph,
		matrix *T.FT_Matrix,
		delta *T.FT_Vector) T.FT_Error

	FT_Glyph_Get_CBox func(
		glyph T.FT_Glyph,
		bbox_mode uint,
		acbox *T.FT_BBox)

	FT_Glyph_To_Bitmap func(
		the_glyph *T.FT_Glyph,
		render_mode T.FT_Render_Mode,
		origin *T.FT_Vector,
		destroy T.FT_Bool) T.FT_Error

	FT_Done_Glyph func(
		glyph T.FT_Glyph)

	FT_Matrix_Multiply func(
		a *T.FT_Matrix,
		b *T.FT_Matrix)

	FT_Matrix_Invert func(
		matrix *T.FT_Matrix) T.FT_Error

	FTC_Manager_New func(
		library T.FT_Library,
		max_faces uint,
		max_sizes uint,
		max_bytes T.FT_ULong,
		requester T.FTC_Face_Requester,
		req_data T.FT_Pointer,
		amanager *T.FTC_Manager) T.FT_Error

	FTC_Manager_Reset func(
		manager T.FTC_Manager)

	FTC_Manager_Done func(
		manager T.FTC_Manager)

	FTC_Manager_LookupFace func(
		manager T.FTC_Manager,
		face_id T.FTC_FaceID,
		aface *T.FT_Face) T.FT_Error

	FTC_Manager_LookupSize func(
		manager T.FTC_Manager,
		scaler T.FTC_Scaler,
		asize *T.FT_Size) T.FT_Error

	FTC_Node_Unref func(
		node T.FTC_Node,
		manager T.FTC_Manager)

	FTC_Manager_RemoveFaceID func(
		manager T.FTC_Manager,
		face_id T.FTC_FaceID)

	FTC_CMapCache_New func(
		manager T.FTC_Manager,
		acache *T.FTC_CMapCache) T.FT_Error

	FTC_CMapCache_Lookup func(
		cache T.FTC_CMapCache,
		face_id T.FTC_FaceID,
		cmap_index int,
		char_code T.FT_UInt32) uint

	FTC_ImageCache_New func(
		manager T.FTC_Manager,
		acache *T.FTC_ImageCache) T.FT_Error

	FTC_ImageCache_Lookup func(
		cache T.FTC_ImageCache,
		t T.FTC_ImageType,
		gindex uint,
		aglyph *T.FT_Glyph,
		anode *T.FTC_Node) T.FT_Error

	FTC_ImageCache_LookupScaler func(
		cache T.FTC_ImageCache,
		scaler T.FTC_Scaler,
		load_flags T.FT_ULong,
		gindex uint,
		aglyph *T.FT_Glyph,
		anode *T.FTC_Node) T.FT_Error

	FTC_SBitCache_New func(
		manager T.FTC_Manager,
		acache *T.FTC_SBitCache) T.FT_Error

	FTC_SBitCache_Lookup func(
		cache T.FTC_SBitCache,
		t T.FTC_ImageType,
		gindex uint,
		sbit *T.FTC_SBit,
		anode *T.FTC_Node) T.FT_Error

	FTC_SBitCache_LookupScaler func(
		cache T.FTC_SBitCache,
		scaler T.FTC_Scaler,
		load_flags T.FT_ULong,
		gindex uint,
		sbit *T.FTC_SBit,
		anode *T.FTC_Node) T.FT_Error

	FTC_Manager_Lookup_Face func(
		manager T.FTC_Manager,
		face_id T.FTC_FaceID,
		aface *T.FT_Face) T.FT_Error

	FTC_Manager_Lookup_Size func(
		manager T.FTC_Manager,
		font T.FTC_Font,
		aface *T.FT_Face,
		asize *T.FT_Size) T.FT_Error

	FT_New_Size func(
		face T.FT_Face,
		size *T.FT_Size) T.FT_Error

	FT_Done_Size func(
		size T.FT_Size) T.FT_Error

	FT_Activate_Size func(
		size T.FT_Size) T.FT_Error

	FT_List_Find func(
		list T.FT_List,
		data *T.Void) T.FT_ListNode

	FT_List_Add func(
		list T.FT_List,
		node T.FT_ListNode)

	FT_List_Insert func(
		list T.FT_List,
		node T.FT_ListNode)

	FT_List_Remove func(
		list T.FT_List,
		node T.FT_ListNode)

	FT_List_Up func(
		list T.FT_List,
		node T.FT_ListNode)

	FT_List_Iterate func(
		list T.FT_List,
		iterator T.FT_List_Iterator,
		user *T.Void) T.FT_Error

	FT_List_Finalize func(
		list T.FT_List,
		destroy T.FT_List_Destructor,
		memory T.FT_Memory,
		user *T.Void)

	FT_Outline_Decompose func(
		outline *T.FT_Outline,
		func_interface *T.FT_Outline_Funcs,
		user *T.Void) T.FT_Error

	FT_Outline_New func(
		library T.FT_Library,
		numPoints uint,
		numContours int,
		anoutline *T.FT_Outline) T.FT_Error

	FT_Outline_New_Internal func(
		memory T.FT_Memory,
		numPoints uint,
		numContours int,
		anoutline *T.FT_Outline) T.FT_Error

	FT_Outline_Done func(
		library T.FT_Library,
		outline *T.FT_Outline) T.FT_Error

	FT_Outline_Done_Internal func(
		memory T.FT_Memory,
		outline *T.FT_Outline) T.FT_Error

	FT_Outline_Check func(
		outline *T.FT_Outline) T.FT_Error

	FT_Outline_Get_CBox func(
		outline *T.FT_Outline,
		acbox *T.FT_BBox)

	FT_Outline_Translate func(
		outline *T.FT_Outline,
		xOffset T.FT_Pos,
		yOffset T.FT_Pos)

	FT_Outline_Copy func(
		source *T.FT_Outline,
		target *T.FT_Outline) T.FT_Error

	FT_Outline_Transform func(
		outline *T.FT_Outline,
		matrix *T.FT_Matrix)

	FT_Outline_Embolden func(
		outline *T.FT_Outline,
		strength T.FT_Pos) T.FT_Error

	FT_Outline_Reverse func(
		outline *T.FT_Outline)

	FT_Outline_Get_Bitmap func(
		library T.FT_Library,
		outline *T.FT_Outline,
		abitmap *T.FT_Bitmap) T.FT_Error

	FT_Outline_Render func(
		library T.FT_Library,
		outline *T.FT_Outline,
		params *T.FT_Raster_Params) T.FT_Error

	FT_Outline_Get_Orientation func(
		outline *T.FT_Outline) T.FT_Orientation

	FT_Add_Module func(
		library T.FT_Library,
		clazz *T.FT_Module_Class) T.FT_Error

	FT_Get_Module func(
		library T.FT_Library,
		module_name string) T.FT_Module

	FT_Remove_Module func(
		library T.FT_Library,
		module T.FT_Module) T.FT_Error

	FT_Reference_Library func(library T.FT_Library) T.FT_Error

	FT_New_Library func(
		memory T.FT_Memory,
		alibrary *T.FT_Library) T.FT_Error

	FT_Done_Library func(library T.FT_Library) T.FT_Error

	FT_Set_Debug_Hook func(
		library T.FT_Library,
		hook_index uint,
		debug_hook T.FT_DebugHook_Func)

	FT_Add_Default_Modules func(library T.FT_Library)

	FT_Get_TrueType_Engine_Type func(
		library T.FT_Library) T.FT_TrueTypeEngineType

	FT_Get_Renderer func(library T.FT_Library,
		format T.FT_Glyph_Format) T.FT_Renderer

	FT_Set_Renderer func(
		library T.FT_Library,
		renderer T.FT_Renderer,
		num_params uint,
		parameters *T.FT_Parameter) T.FT_Error

	FT_Get_Sfnt_Table func(
		face T.FT_Face, tag T.FT_Sfnt_Tag) *T.Void

	FT_Load_Sfnt_Table func(
		face T.FT_Face,
		tag T.FT_ULong,
		offset T.FT_Long,
		buffer *T.FT_Byte,
		length *T.FT_ULong) T.FT_Error

	FT_Sfnt_Table_Info func(
		face T.FT_Face,
		table_index uint,
		tag, length *T.FT_ULong) T.FT_Error

	FT_Get_CMap_Language_ID func(charmap T.FT_CharMap) T.FT_ULong

	FT_Get_CMap_Format func(charmap T.FT_CharMap) T.FT_Long

	FT_Get_BDF_Charset_ID func(face T.FT_Face,
		acharset_encoding, acharset_registry **T.Char) T.FT_Error

	FT_Get_BDF_Property func(
		face T.FT_Face,
		prop_name string,
		aproperty *T.BDF_PropertyRec) T.FT_Error

	FT_Get_CID_Registry_Ordering_Supplement func(
		face T.FT_Face,
		registry, ordering **T.Char,
		supplement *int) T.FT_Error

	FT_Get_CID_Is_Internally_CID_Keyed func(
		face T.FT_Face,
		is_cid *T.FT_Bool) T.FT_Error

	FT_Get_CID_From_Glyph_Index func(
		face T.FT_Face,
		glyph_index uint,
		cid *uint) T.FT_Error

	FT_Stream_OpenGzip func(
		stream, source T.FT_Stream) T.FT_Error

	FT_Stream_OpenLZW func(
		stream, source T.FT_Stream) T.FT_Error

	FT_Get_WinFNT_Header func(
		face T.FT_Face,
		aheader *T.FT_WinFNT_HeaderRec) T.FT_Error

	FT_Bitmap_New func(abitmap *T.FT_Bitmap)

	FT_Bitmap_Copy func(
		library T.FT_Library,
		source, target *T.FT_Bitmap) T.FT_Error

	FT_Bitmap_Embolden func(
		library T.FT_Library,
		bitmap *T.FT_Bitmap,
		xStrength, yStrength T.FT_Pos) T.FT_Error

	FT_Bitmap_Convert func(
		library T.FT_Library,
		source, target *T.FT_Bitmap,
		alignment int) T.FT_Error

	FT_GlyphSlot_Own_Bitmap func(slot T.FT_GlyphSlot) T.FT_Error

	FT_Bitmap_Done func(
		library T.FT_Library, bitmap *T.FT_Bitmap) T.FT_Error

	FT_Outline_Get_BBox func(
		outline *T.FT_Outline, abbox *T.FT_BBox) T.FT_Error

	FT_Get_Sfnt_Name_Count func(face T.FT_Face) uint

	FT_Get_Sfnt_Name func(
		face T.FT_Face,
		idx uint,
		aname *T.FT_SfntName) T.FT_Error

	FT_OpenType_Validate func(
		face T.FT_Face, validation_flags uint,
		BASE_table, GDEF_table, GPOS_table,
		GSUB_table, JSTF_table *T.FT_Bytes) T.FT_Error

	FT_OpenType_Free func(face T.FT_Face, table T.FT_Bytes)

	FT_TrueTypeGX_Validate func(
		face T.FT_Face,
		validation_flags uint,
		tables [FT_VALIDATE_GX_LENGTH]T.FT_Bytes,
		table_length uint) T.FT_Error

	FT_TrueTypeGX_Free func(face T.FT_Face, table T.FT_Bytes)

	FT_ClassicKern_Validate func(
		face T.FT_Face,
		validation_flags uint,
		ckern_table *T.FT_Bytes) T.FT_Error

	FT_ClassicKern_Free func(
		face T.FT_Face, table T.FT_Bytes)

	FT_Get_PFR_Metrics func(face T.FT_Face,
		aoutline_resolution, ametrics_resolution *uint,
		ametrics_x_scale, ametrics_y_scale *T.FT_Fixed) T.FT_Error

	FT_Get_PFR_Kerning func(
		face T.FT_Face,
		left, right uint,
		avector *T.FT_Vector) T.FT_Error

	FT_Get_PFR_Advance func(
		face T.FT_Face,
		gindex uint,
		aadvance *T.FT_Pos) T.FT_Error

	FT_Outline_GetInsideBorder func(
		outline *T.FT_Outline) T.FT_StrokerBorder

	FT_Outline_GetOutsideBorder func(
		outline *T.FT_Outline) T.FT_StrokerBorder

	FT_Stroker_New func(
		library T.FT_Library,
		astroker *T.FT_Stroker) T.FT_Error

	FT_Stroker_Set func(
		stroker T.FT_Stroker,
		radius T.FT_Fixed,
		line_cap T.FT_Stroker_LineCap,
		line_join T.FT_Stroker_LineJoin,
		miter_limit T.FT_Fixed)

	FT_Stroker_Rewind func(stroker T.FT_Stroker)

	FT_Stroker_ParseOutline func(stroker T.FT_Stroker,
		outline *T.FT_Outline, opened T.FT_Bool) T.FT_Error

	FT_Stroker_BeginSubPath func(stroker T.FT_Stroker,
		to *T.FT_Vector, open T.FT_Bool) T.FT_Error

	FT_Stroker_EndSubPath func(stroker T.FT_Stroker) T.FT_Error

	FT_Stroker_LineTo func(
		stroker T.FT_Stroker, to *T.FT_Vector) T.FT_Error

	FT_Stroker_ConicTo func(stroker T.FT_Stroker,
		control, to *T.FT_Vector) T.FT_Error

	FT_Stroker_CubicTo func(stroker T.FT_Stroker,
		control1, control2, to *T.FT_Vector) T.FT_Error

	FT_Stroker_GetBorderCounts func(
		stroker T.FT_Stroker, border T.FT_StrokerBorder,
		anum_points, anum_contours *uint) T.FT_Error

	FT_Stroker_ExportBorder func(stroker T.FT_Stroker,
		border T.FT_StrokerBorder, outline *T.FT_Outline)

	FT_Stroker_GetCounts func(stroker T.FT_Stroker,
		anum_points, anum_contours *uint) T.FT_Error

	FT_Stroker_Export func(
		stroker T.FT_Stroker, outline *T.FT_Outline)

	FT_Stroker_Done func(stroker T.FT_Stroker)

	FT_Glyph_Stroke func(pglyph *T.FT_Glyph,
		stroker T.FT_Stroker, destroy T.FT_Bool) T.FT_Error

	FT_Glyph_StrokeBorder func(
		pglyph *T.FT_Glyph, stroker T.FT_Stroker,
		inside, destroy T.FT_Bool) T.FT_Error

	FT_GlyphSlot_Embolden func(slot T.FT_GlyphSlot)

	FT_GlyphSlot_Oblique func(slot T.FT_GlyphSlot)

	FT_Get_X11_Font_Format func(face T.FT_Face) string

	FT_Sin func(angle T.FT_Angle) T.FT_Fixed

	FT_Cos func(angle T.FT_Angle) T.FT_Fixed

	FT_Tan func(angle T.FT_Angle) T.FT_Fixed

	FT_Atan2 func(x, y T.FT_Fixed) T.FT_Angle

	FT_Angle_Diff func(angle1, angle2 T.FT_Angle) T.FT_Angle

	FT_Vector_Unit func(vec *T.FT_Vector, angle T.FT_Angle)

	FT_Vector_Rotate func(vec *T.FT_Vector, angle T.FT_Angle)

	FT_Vector_Length func(vec *T.FT_Vector) T.FT_Fixed

	FT_Vector_Polarize func(
		vec *T.FT_Vector, length *T.FT_Fixed, angle *T.FT_Angle)

	FT_Vector_From_Polar func(
		vec *T.FT_Vector, length T.FT_Fixed, angle T.FT_Angle)

	FT_Library_SetLcdFilter func(
		library T.FT_Library,
		filter T.FT_LcdFilter) T.FT_Error

	FT_Library_SetLcdFilterWeights func(
		library T.FT_Library,
		weights *T.Unsigned_char) T.FT_Error

	FT_Get_Gasp func(
		face T.FT_Face, ppem uint) int

	FT_Get_Advance func(
		face T.FT_Face,
		gindex uint,
		load_flags T.FT_Int32,
		padvance *T.FT_Fixed) T.FT_Error

	FT_Get_Advances func(
		face T.FT_Face,
		start, count uint,
		load_flags T.FT_Int32,
		padvances *T.FT_Fixed) T.FT_Error

	FT_Done_FreeType func(library T.FT_Library) T.FT_Error

	FT_Init_FreeType func(alibrary *T.FT_Library) T.FT_Error

	FT_MulFix func(a, b T.FT_Long) T.FT_Long

	FT_Render_Glyph func(
		slot T.FT_GlyphSlot,
		render_mode T.FT_Render_Mode) T.FT_Error
)

var dll = "freetype6.dll"

var apiList = outside.Apis{
	{"FTC_CMapCache_Lookup", &FTC_CMapCache_Lookup},
	{"FTC_CMapCache_New", &FTC_CMapCache_New},
	{"FTC_ImageCache_Lookup", &FTC_ImageCache_Lookup},
	{"FTC_ImageCache_LookupScaler", &FTC_ImageCache_LookupScaler},
	{"FTC_ImageCache_New", &FTC_ImageCache_New},
	{"FTC_Manager_Done", &FTC_Manager_Done},
	{"FTC_Manager_LookupFace", &FTC_Manager_LookupFace},
	{"FTC_Manager_LookupSize", &FTC_Manager_LookupSize},
	{"FTC_Manager_Lookup_Face", &FTC_Manager_Lookup_Face},
	{"FTC_Manager_Lookup_Size", &FTC_Manager_Lookup_Size},
	{"FTC_Manager_New", &FTC_Manager_New},
	{"FTC_Manager_RemoveFaceID", &FTC_Manager_RemoveFaceID},
	{"FTC_Manager_Reset", &FTC_Manager_Reset},
	{"FTC_Node_Unref", &FTC_Node_Unref},
	{"FTC_SBitCache_Lookup", &FTC_SBitCache_Lookup},
	{"FTC_SBitCache_LookupScaler", &FTC_SBitCache_LookupScaler},
	{"FTC_SBitCache_New", &FTC_SBitCache_New},
	{"FT_Activate_Size", &FT_Activate_Size},
	{"FT_Add_Default_Modules", &FT_Add_Default_Modules},
	{"FT_Add_Module", &FT_Add_Module},
	{"FT_Angle_Diff", &FT_Angle_Diff},
	{"FT_Atan2", &FT_Atan2},
	{"FT_Attach_File", &FT_Attach_File},
	{"FT_Attach_Stream", &FT_Attach_Stream},
	{"FT_Bitmap_Convert", &FT_Bitmap_Convert},
	{"FT_Bitmap_Copy", &FT_Bitmap_Copy},
	{"FT_Bitmap_Done", &FT_Bitmap_Done},
	{"FT_Bitmap_Embolden", &FT_Bitmap_Embolden},
	{"FT_Bitmap_New", &FT_Bitmap_New},
	{"FT_CeilFix", &FT_CeilFix},
	{"FT_ClassicKern_Free", &FT_ClassicKern_Free},
	{"FT_ClassicKern_Validate", &FT_ClassicKern_Validate},
	{"FT_Cos", &FT_Cos},
	{"FT_DivFix", &FT_DivFix},
	{"FT_Done_Face", &FT_Done_Face},
	{"FT_Done_FreeType", &FT_Done_FreeType},
	{"FT_Done_Glyph", &FT_Done_Glyph},
	{"FT_Done_Library", &FT_Done_Library},
	{"FT_Done_Size", &FT_Done_Size},
	{"FT_Face_CheckTrueTypePatents", &FT_Face_CheckTrueTypePatents},
	{"FT_Face_GetCharVariantIndex", &FT_Face_GetCharVariantIndex},
	{"FT_Face_GetCharVariantIsDefault", &FT_Face_GetCharVariantIsDefault},
	{"FT_Face_GetCharsOfVariant", &FT_Face_GetCharsOfVariant},
	{"FT_Face_GetVariantSelectors", &FT_Face_GetVariantSelectors},
	{"FT_Face_GetVariantsOfChar", &FT_Face_GetVariantsOfChar},
	{"FT_Face_SetUnpatentedHinting", &FT_Face_SetUnpatentedHinting},
	{"FT_FloorFix", &FT_FloorFix},
	{"FT_Get_Advance", &FT_Get_Advance},
	{"FT_Get_Advances", &FT_Get_Advances},
	{"FT_Get_BDF_Charset_ID", &FT_Get_BDF_Charset_ID},
	{"FT_Get_BDF_Property", &FT_Get_BDF_Property},
	{"FT_Get_CID_From_Glyph_Index", &FT_Get_CID_From_Glyph_Index},
	{"FT_Get_CID_Is_Internally_CID_Keyed", &FT_Get_CID_Is_Internally_CID_Keyed},
	{"FT_Get_CID_Registry_Ordering_Supplement", &FT_Get_CID_Registry_Ordering_Supplement},
	{"FT_Get_CMap_Format", &FT_Get_CMap_Format},
	{"FT_Get_CMap_Language_ID", &FT_Get_CMap_Language_ID},
	{"FT_Get_Char_Index", &FT_Get_Char_Index},
	{"FT_Get_Charmap_Index", &FT_Get_Charmap_Index},
	{"FT_Get_FSType_Flags", &FT_Get_FSType_Flags},
	{"FT_Get_First_Char", &FT_Get_First_Char},
	{"FT_Get_Gasp", &FT_Get_Gasp},
	{"FT_Get_Glyph", &FT_Get_Glyph},
	{"FT_Get_Glyph_Name", &FT_Get_Glyph_Name},
	{"FT_Get_Kerning", &FT_Get_Kerning},
	{"FT_Get_MM_Var", &FT_Get_MM_Var},
	{"FT_Get_Module", &FT_Get_Module},
	{"FT_Get_Multi_Master", &FT_Get_Multi_Master},
	{"FT_Get_Name_Index", &FT_Get_Name_Index},
	{"FT_Get_Next_Char", &FT_Get_Next_Char},
	{"FT_Get_PFR_Advance", &FT_Get_PFR_Advance},
	{"FT_Get_PFR_Kerning", &FT_Get_PFR_Kerning},
	{"FT_Get_PFR_Metrics", &FT_Get_PFR_Metrics},
	{"FT_Get_PS_Font_Info", &FT_Get_PS_Font_Info},
	{"FT_Get_PS_Font_Private", &FT_Get_PS_Font_Private},
	{"FT_Get_Postscript_Name", &FT_Get_Postscript_Name},
	{"FT_Get_Renderer", &FT_Get_Renderer},
	{"FT_Get_Sfnt_Name", &FT_Get_Sfnt_Name},
	{"FT_Get_Sfnt_Name_Count", &FT_Get_Sfnt_Name_Count},
	{"FT_Get_Sfnt_Table", &FT_Get_Sfnt_Table},
	{"FT_Get_SubGlyph_Info", &FT_Get_SubGlyph_Info},
	{"FT_Get_Track_Kerning", &FT_Get_Track_Kerning},
	{"FT_Get_TrueType_Engine_Type", &FT_Get_TrueType_Engine_Type},
	{"FT_Get_WinFNT_Header", &FT_Get_WinFNT_Header},
	{"FT_Get_X11_Font_Format", &FT_Get_X11_Font_Format},
	{"FT_GlyphSlot_Embolden", &FT_GlyphSlot_Embolden},
	{"FT_GlyphSlot_Oblique", &FT_GlyphSlot_Oblique},
	{"FT_GlyphSlot_Own_Bitmap", &FT_GlyphSlot_Own_Bitmap},
	{"FT_Glyph_Copy", &FT_Glyph_Copy},
	{"FT_Glyph_Get_CBox", &FT_Glyph_Get_CBox},
	{"FT_Glyph_Stroke", &FT_Glyph_Stroke},
	{"FT_Glyph_StrokeBorder", &FT_Glyph_StrokeBorder},
	{"FT_Glyph_To_Bitmap", &FT_Glyph_To_Bitmap},
	{"FT_Glyph_Transform", &FT_Glyph_Transform},
	{"FT_Has_PS_Glyph_Names", &FT_Has_PS_Glyph_Names},
	{"FT_Init_FreeType", &FT_Init_FreeType},
	{"FT_Library_SetLcdFilter", &FT_Library_SetLcdFilter},
	{"FT_Library_SetLcdFilterWeights", &FT_Library_SetLcdFilterWeights},
	{"FT_Library_Version", &FT_Library_Version},
	{"FT_List_Add", &FT_List_Add},
	{"FT_List_Finalize", &FT_List_Finalize},
	{"FT_List_Find", &FT_List_Find},
	{"FT_List_Insert", &FT_List_Insert},
	{"FT_List_Iterate", &FT_List_Iterate},
	{"FT_List_Remove", &FT_List_Remove},
	{"FT_List_Up", &FT_List_Up},
	{"FT_Load_Char", &FT_Load_Char},
	{"FT_Load_Glyph", &FT_Load_Glyph},
	{"FT_Load_Sfnt_Table", &FT_Load_Sfnt_Table},
	{"FT_Matrix_Invert", &FT_Matrix_Invert},
	{"FT_Matrix_Multiply", &FT_Matrix_Multiply},
	{"FT_MulDiv", &FT_MulDiv},
	{"FT_MulFix", &FT_MulFix},
	{"FT_New_Face", &FT_New_Face},
	{"FT_New_Library", &FT_New_Library},
	{"FT_New_Memory_Face", &FT_New_Memory_Face},
	{"FT_New_Size", &FT_New_Size},
	{"FT_OpenType_Free", &FT_OpenType_Free},
	{"FT_OpenType_Validate", &FT_OpenType_Validate},
	{"FT_Open_Face", &FT_Open_Face},
	{"FT_Outline_Check", &FT_Outline_Check},
	{"FT_Outline_Copy", &FT_Outline_Copy},
	{"FT_Outline_Decompose", &FT_Outline_Decompose},
	{"FT_Outline_Done", &FT_Outline_Done},
	{"FT_Outline_Done_Internal", &FT_Outline_Done_Internal},
	{"FT_Outline_Embolden", &FT_Outline_Embolden},
	{"FT_Outline_GetInsideBorder", &FT_Outline_GetInsideBorder},
	{"FT_Outline_GetOutsideBorder", &FT_Outline_GetOutsideBorder},
	{"FT_Outline_Get_BBox", &FT_Outline_Get_BBox},
	{"FT_Outline_Get_Bitmap", &FT_Outline_Get_Bitmap},
	{"FT_Outline_Get_CBox", &FT_Outline_Get_CBox},
	{"FT_Outline_Get_Orientation", &FT_Outline_Get_Orientation},
	{"FT_Outline_New", &FT_Outline_New},
	{"FT_Outline_New_Internal", &FT_Outline_New_Internal},
	{"FT_Outline_Render", &FT_Outline_Render},
	{"FT_Outline_Reverse", &FT_Outline_Reverse},
	{"FT_Outline_Transform", &FT_Outline_Transform},
	{"FT_Outline_Translate", &FT_Outline_Translate},
	{"FT_Reference_Face", &FT_Reference_Face},
	{"FT_Reference_Library", &FT_Reference_Library},
	{"FT_Remove_Module", &FT_Remove_Module},
	{"FT_Render_Glyph", &FT_Render_Glyph},
	{"FT_Request_Size", &FT_Request_Size},
	{"FT_RoundFix", &FT_RoundFix},
	{"FT_Select_Charmap", &FT_Select_Charmap},
	{"FT_Select_Size", &FT_Select_Size},
	{"FT_Set_Char_Size", &FT_Set_Char_Size},
	{"FT_Set_Charmap", &FT_Set_Charmap},
	{"FT_Set_Debug_Hook", &FT_Set_Debug_Hook},
	{"FT_Set_MM_Blend_Coordinates", &FT_Set_MM_Blend_Coordinates},
	{"FT_Set_MM_Design_Coordinates", &FT_Set_MM_Design_Coordinates},
	{"FT_Set_Pixel_Sizes", &FT_Set_Pixel_Sizes},
	{"FT_Set_Renderer", &FT_Set_Renderer},
	{"FT_Set_Transform", &FT_Set_Transform},
	{"FT_Set_Var_Blend_Coordinates", &FT_Set_Var_Blend_Coordinates},
	{"FT_Set_Var_Design_Coordinates", &FT_Set_Var_Design_Coordinates},
	{"FT_Sfnt_Table_Info", &FT_Sfnt_Table_Info},
	{"FT_Sin", &FT_Sin},
	{"FT_Stream_OpenGzip", &FT_Stream_OpenGzip},
	{"FT_Stream_OpenLZW", &FT_Stream_OpenLZW},
	{"FT_Stroker_BeginSubPath", &FT_Stroker_BeginSubPath},
	{"FT_Stroker_ConicTo", &FT_Stroker_ConicTo},
	{"FT_Stroker_CubicTo", &FT_Stroker_CubicTo},
	{"FT_Stroker_Done", &FT_Stroker_Done},
	{"FT_Stroker_EndSubPath", &FT_Stroker_EndSubPath},
	{"FT_Stroker_Export", &FT_Stroker_Export},
	{"FT_Stroker_ExportBorder", &FT_Stroker_ExportBorder},
	{"FT_Stroker_GetBorderCounts", &FT_Stroker_GetBorderCounts},
	{"FT_Stroker_GetCounts", &FT_Stroker_GetCounts},
	{"FT_Stroker_LineTo", &FT_Stroker_LineTo},
	{"FT_Stroker_New", &FT_Stroker_New},
	{"FT_Stroker_ParseOutline", &FT_Stroker_ParseOutline},
	{"FT_Stroker_Rewind", &FT_Stroker_Rewind},
	{"FT_Stroker_Set", &FT_Stroker_Set},
	{"FT_Tan", &FT_Tan},
	{"FT_TrueTypeGX_Free", &FT_TrueTypeGX_Free},
	{"FT_TrueTypeGX_Validate", &FT_TrueTypeGX_Validate},
	{"FT_Vector_From_Polar", &FT_Vector_From_Polar},
	{"FT_Vector_Length", &FT_Vector_Length},
	{"FT_Vector_Polarize", &FT_Vector_Polarize},
	{"FT_Vector_Rotate", &FT_Vector_Rotate},
	{"FT_Vector_Transform", &FT_Vector_Transform},
	{"FT_Vector_Unit", &FT_Vector_Unit},
	// Undocumented {"TT_New_Context", &TT_New_Context},
	// Undocumented {"TT_RunIns", &TT_RunIns},
}
