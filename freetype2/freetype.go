package freetype

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	FT_New_Face func(
		library FT_Library,
		filepathname *Char,
		face_index FT_Long,
		aface *FT_Face) FT_Error

	FT_New_Memory_Face func(
		library FT_Library,
		file_base *FT_Byte,
		file_size FT_Long,
		face_index FT_Long,
		aface *FT_Face) FT_Error

	FT_Open_Face func(
		library FT_Library,
		args *FT_Open_Args,
		face_index FT_Long,
		aface *FT_Face) FT_Error

	FT_Attach_File func(
		face FT_Face,
		filepathname *Char) FT_Error

	FT_Attach_Stream func(
		face FT_Face,
		parameters *FT_Open_Args) FT_Error

	FT_Reference_Face func(
		face FT_Face) FT_Error

	FT_Done_Face func(
		face FT_Face) FT_Error

	FT_Select_Size func(
		face FT_Face,
		strike_index FT_Int) FT_Error

	FT_Request_Size func(
		face FT_Face,
		req FT_Size_Request) FT_Error

	FT_Set_Char_Size func(
		face FT_Face,
		char_width FT_F26Dot6,
		char_height FT_F26Dot6,
		horz_resolution FT_UInt,
		vert_resolution FT_UInt) FT_Error

	FT_Set_Pixel_Sizes func(
		face FT_Face,
		pixel_width FT_UInt,
		pixel_height FT_UInt) FT_Error

	FT_Load_Glyph func(
		face FT_Face,
		glyph_index FT_UInt,
		load_flags FT_Int32) FT_Error

	FT_Load_Char func(
		face FT_Face,
		char_code FT_ULong,
		load_flags FT_Int32) FT_Error

	FT_Set_Transform func(
		face FT_Face,
		matrix *FT_Matrix,
		delta *FT_Vector)

	FT_Get_Kerning func(
		face FT_Face,
		left_glyph FT_UInt,
		right_glyph FT_UInt,
		kern_mode FT_UInt,
		akerning *FT_Vector) FT_Error

	FT_Get_Track_Kerning func(
		face FT_Face,
		point_size FT_Fixed,
		degree FT_Int,
		akerning *FT_Fixed) FT_Error

	FT_Get_Glyph_Name func(
		face FT_Face,
		glyph_index FT_UInt,
		buffer FT_Pointer,
		buffer_max FT_UInt) FT_Error

	FT_Get_Postscript_Name func(
		face FT_Face) *Char

	FT_Select_Charmap func(
		face FT_Face,
		encoding FT_Encoding) FT_Error

	FT_Set_Charmap func(
		face FT_Face,
		charmap FT_CharMap) FT_Error

	FT_Get_Charmap_Index func(
		charmap FT_CharMap) FT_Int

	FT_Get_Char_Index func(
		face FT_Face,
		charcode FT_ULong) FT_UInt

	FT_Get_First_Char func(
		face FT_Face,
		agindex *FT_UInt) FT_ULong

	FT_Get_Next_Char func(
		face FT_Face,
		char_code FT_ULong,
		agindex *FT_UInt) FT_ULong

	FT_Get_Name_Index func(
		face FT_Face,
		glyph_name *FT_String) FT_UInt

	FT_Get_SubGlyph_Info func(
		glyph FT_GlyphSlot,
		sub_index FT_UInt,
		p_index *FT_Int,
		p_flags *FT_UInt,
		p_arg1 *FT_Int,
		p_arg2 *FT_Int,
		p_transform *FT_Matrix) FT_Error

	FT_Get_FSType_Flags func(
		face FT_Face) FT_UShort

	FT_Face_GetCharVariantIndex func(
		face FT_Face,
		charcode FT_ULong,
		variantSelector FT_ULong) FT_UInt

	FT_Face_GetCharVariantIsDefault func(
		face FT_Face,
		charcode FT_ULong,
		variantSelector FT_ULong) FT_Int

	FT_Face_GetVariantSelectors func(
		face FT_Face) *FT_UInt32

	FT_Face_GetVariantsOfChar func(
		face FT_Face,
		charcode FT_ULong) *FT_UInt32

	FT_Face_GetCharsOfVariant func(
		face FT_Face,
		variantSelector FT_ULong) *FT_UInt32

	FT_MulDiv func(
		a FT_Long,
		b FT_Long,
		c FT_Long) FT_Long

	FT_DivFix func(
		a FT_Long,
		b FT_Long) FT_Long

	FT_RoundFix func(
		a FT_Fixed) FT_Fixed

	FT_CeilFix func(
		a FT_Fixed) FT_Fixed

	FT_FloorFix func(
		a FT_Fixed) FT_Fixed

	FT_Vector_Transform func(
		vec *FT_Vector,
		matrix *FT_Matrix)

	FT_Library_Version func(
		library FT_Library,
		amajor *FT_Int,
		aminor *FT_Int,
		apatch *FT_Int)

	FT_Face_CheckTrueTypePatents func(
		face FT_Face) FT_Bool

	FT_Face_SetUnpatentedHinting func(
		face FT_Face,
		value FT_Bool) FT_Bool

	FT_Has_PS_Glyph_Names func(
		face FT_Face) FT_Int

	FT_Get_PS_Font_Info func(
		face FT_Face,
		afont_info PS_FontInfo) FT_Error

	FT_Get_PS_Font_Private func(
		face FT_Face,
		afont_private PS_Private) FT_Error

	FT_Get_Multi_Master func(
		face FT_Face,
		amaster *FT_Multi_Master) FT_Error

	FT_Get_MM_Var func(
		face FT_Face,
		amaster **FT_MM_Var) FT_Error

	FT_Set_MM_Design_Coordinates func(
		face FT_Face,
		num_coords FT_UInt,
		coords *FT_Long) FT_Error

	FT_Set_Var_Design_Coordinates func(
		face FT_Face,
		num_coords FT_UInt,
		coords *FT_Fixed) FT_Error

	FT_Set_MM_Blend_Coordinates func(
		face FT_Face,
		num_coords FT_UInt,
		coords *FT_Fixed) FT_Error

	FT_Set_Var_Blend_Coordinates func(
		face FT_Face,
		num_coords FT_UInt,
		coords *FT_Fixed) FT_Error

	FT_Get_Glyph func(
		slot FT_GlyphSlot,
		aglyph *FT_Glyph) FT_Error

	FT_Glyph_Copy func(
		source FT_Glyph,
		target *FT_Glyph) FT_Error

	FT_Glyph_Transform func(
		glyph FT_Glyph,
		matrix *FT_Matrix,
		delta *FT_Vector) FT_Error

	FT_Glyph_Get_CBox func(
		glyph FT_Glyph,
		bbox_mode FT_UInt,
		acbox *FT_BBox)

	FT_Glyph_To_Bitmap func(
		the_glyph *FT_Glyph,
		render_mode FT_Render_Mode,
		origin *FT_Vector,
		destroy FT_Bool) FT_Error

	FT_Done_Glyph func(
		glyph FT_Glyph)

	FT_Matrix_Multiply func(
		a *FT_Matrix,
		b *FT_Matrix)

	FT_Matrix_Invert func(
		matrix *FT_Matrix) FT_Error

	FTC_Manager_New func(
		library FT_Library,
		max_faces FT_UInt,
		max_sizes FT_UInt,
		max_bytes FT_ULong,
		requester FTC_Face_Requester,
		req_data FT_Pointer,
		amanager *FTC_Manager) FT_Error

	FTC_Manager_Reset func(
		manager FTC_Manager)

	FTC_Manager_Done func(
		manager FTC_Manager)

	FTC_Manager_LookupFace func(
		manager FTC_Manager,
		face_id FTC_FaceID,
		aface *FT_Face) FT_Error

	FTC_Manager_LookupSize func(
		manager FTC_Manager,
		scaler FTC_Scaler,
		asize *FT_Size) FT_Error

	FTC_Node_Unref func(
		node FTC_Node,
		manager FTC_Manager)

	FTC_Manager_RemoveFaceID func(
		manager FTC_Manager,
		face_id FTC_FaceID)

	FTC_CMapCache_New func(
		manager FTC_Manager,
		acache *FTC_CMapCache) FT_Error

	FTC_CMapCache_Lookup func(
		cache FTC_CMapCache,
		face_id FTC_FaceID,
		cmap_index FT_Int,
		char_code FT_UInt32) FT_UInt

	FTC_ImageCache_New func(
		manager FTC_Manager,
		acache *FTC_ImageCache) FT_Error

	FTC_ImageCache_Lookup func(
		cache FTC_ImageCache,
		t FTC_ImageType,
		gindex FT_UInt,
		aglyph *FT_Glyph,
		anode *FTC_Node) FT_Error

	FTC_ImageCache_LookupScaler func(
		cache FTC_ImageCache,
		scaler FTC_Scaler,
		load_flags FT_ULong,
		gindex FT_UInt,
		aglyph *FT_Glyph,
		anode *FTC_Node) FT_Error

	FTC_SBitCache_New func(
		manager FTC_Manager,
		acache *FTC_SBitCache) FT_Error

	FTC_SBitCache_Lookup func(
		cache FTC_SBitCache,
		t FTC_ImageType,
		gindex FT_UInt,
		sbit *FTC_SBit,
		anode *FTC_Node) FT_Error

	FTC_SBitCache_LookupScaler func(
		cache FTC_SBitCache,
		scaler FTC_Scaler,
		load_flags FT_ULong,
		gindex FT_UInt,
		sbit *FTC_SBit,
		anode *FTC_Node) FT_Error

	FTC_Manager_Lookup_Face func(
		manager FTC_Manager,
		face_id FTC_FaceID,
		aface *FT_Face) FT_Error

	FTC_Manager_Lookup_Size func(
		manager FTC_Manager,
		font FTC_Font,
		aface *FT_Face,
		asize *FT_Size) FT_Error

	FT_New_Size func(
		face FT_Face,
		size *FT_Size) FT_Error

	FT_Done_Size func(
		size FT_Size) FT_Error

	FT_Activate_Size func(
		size FT_Size) FT_Error

	FT_List_Find func(
		list FT_List,
		data *Void) FT_ListNode

	FT_List_Add func(
		list FT_List,
		node FT_ListNode)

	FT_List_Insert func(
		list FT_List,
		node FT_ListNode)

	FT_List_Remove func(
		list FT_List,
		node FT_ListNode)

	FT_List_Up func(
		list FT_List,
		node FT_ListNode)

	FT_List_Iterate func(
		list FT_List,
		iterator FT_List_Iterator,
		user *Void) FT_Error

	FT_List_Finalize func(
		list FT_List,
		destroy FT_List_Destructor,
		memory FT_Memory,
		user *Void)

	FT_Outline_Decompose func(
		outline *FT_Outline,
		func_interface *FT_Outline_Funcs,
		user *Void) FT_Error

	FT_Outline_New func(
		library FT_Library,
		numPoints FT_UInt,
		numContours FT_Int,
		anoutline *FT_Outline) FT_Error

	FT_Outline_New_Internal func(
		memory FT_Memory,
		numPoints FT_UInt,
		numContours FT_Int,
		anoutline *FT_Outline) FT_Error

	FT_Outline_Done func(
		library FT_Library,
		outline *FT_Outline) FT_Error

	FT_Outline_Done_Internal func(
		memory FT_Memory,
		outline *FT_Outline) FT_Error

	FT_Outline_Check func(
		outline *FT_Outline) FT_Error

	FT_Outline_Get_CBox func(
		outline *FT_Outline,
		acbox *FT_BBox)

	FT_Outline_Translate func(
		outline *FT_Outline,
		xOffset FT_Pos,
		yOffset FT_Pos)

	FT_Outline_Copy func(
		source *FT_Outline,
		target *FT_Outline) FT_Error

	FT_Outline_Transform func(
		outline *FT_Outline,
		matrix *FT_Matrix)

	FT_Outline_Embolden func(
		outline *FT_Outline,
		strength FT_Pos) FT_Error

	FT_Outline_Reverse func(
		outline *FT_Outline)

	FT_Outline_Get_Bitmap func(
		library FT_Library,
		outline *FT_Outline,
		abitmap *FT_Bitmap) FT_Error

	FT_Outline_Render func(
		library FT_Library,
		outline *FT_Outline,
		params *FT_Raster_Params) FT_Error

	FT_Outline_Get_Orientation func(
		outline *FT_Outline) FT_Orientation

	FT_Add_Module func(
		library FT_Library,
		clazz *FT_Module_Class) FT_Error

	FT_Get_Module func(
		library FT_Library,
		module_name *Char) FT_Module

	FT_Remove_Module func(
		library FT_Library,
		module FT_Module) FT_Error

	FT_Reference_Library func(
		library FT_Library) FT_Error

	FT_New_Library func(
		memory FT_Memory,
		alibrary *FT_Library) FT_Error

	FT_Done_Library func(
		library FT_Library) FT_Error

	FT_Set_Debug_Hook func(
		library FT_Library,
		hook_index FT_UInt,
		debug_hook FT_DebugHook_Func)

	FT_Add_Default_Modules func(
		library FT_Library)

	FT_Get_TrueType_Engine_Type func(
		library FT_Library) FT_TrueTypeEngineType

	FT_Get_Renderer func(
		library FT_Library,
		format FT_Glyph_Format) FT_Renderer

	FT_Set_Renderer func(
		library FT_Library,
		renderer FT_Renderer,
		num_params FT_UInt,
		parameters *FT_Parameter) FT_Error

	FT_Get_Sfnt_Table func(
		face FT_Face,
		tag FT_Sfnt_Tag) *Void

	FT_Load_Sfnt_Table func(
		face FT_Face,
		tag FT_ULong,
		offset FT_Long,
		buffer *FT_Byte,
		length *FT_ULong) FT_Error

	FT_Sfnt_Table_Info func(
		face FT_Face,
		table_index FT_UInt,
		tag *FT_ULong,
		length *FT_ULong) FT_Error

	FT_Get_CMap_Language_ID func(
		charmap FT_CharMap) FT_ULong

	FT_Get_CMap_Format func(
		charmap FT_CharMap) FT_Long

	FT_Get_BDF_Charset_ID func(
		face FT_Face,
		acharset_encoding **Char,
		acharset_registry **Char) FT_Error

	FT_Get_BDF_Property func(
		face FT_Face,
		prop_name *Char,
		aproperty *BDF_PropertyRec) FT_Error

	FT_Get_CID_Registry_Ordering_Supplement func(
		face FT_Face,
		registry **Char,
		ordering **Char,
		supplement *FT_Int) FT_Error

	FT_Get_CID_Is_Internally_CID_Keyed func(
		face FT_Face,
		is_cid *FT_Bool) FT_Error

	FT_Get_CID_From_Glyph_Index func(
		face FT_Face,
		glyph_index FT_UInt,
		cid *FT_UInt) FT_Error

	FT_Stream_OpenGzip func(
		stream FT_Stream,
		source FT_Stream) FT_Error

	FT_Stream_OpenLZW func(
		stream FT_Stream,
		source FT_Stream) FT_Error

	FT_Get_WinFNT_Header func(
		face FT_Face,
		aheader *FT_WinFNT_HeaderRec) FT_Error

	FT_Bitmap_New func(
		abitmap *FT_Bitmap)

	FT_Bitmap_Copy func(
		library FT_Library,
		source *FT_Bitmap,
		target *FT_Bitmap) FT_Error

	FT_Bitmap_Embolden func(
		library FT_Library,
		bitmap *FT_Bitmap,
		xStrength FT_Pos,
		yStrength FT_Pos) FT_Error

	FT_Bitmap_Convert func(
		library FT_Library,
		source *FT_Bitmap,
		target *FT_Bitmap,
		alignment FT_Int) FT_Error

	FT_GlyphSlot_Own_Bitmap func(
		slot FT_GlyphSlot) FT_Error

	FT_Bitmap_Done func(
		library FT_Library,
		bitmap *FT_Bitmap) FT_Error

	FT_Outline_Get_BBox func(
		outline *FT_Outline,
		abbox *FT_BBox) FT_Error

	FT_Get_Sfnt_Name_Count func(
		face FT_Face) FT_UInt

	FT_Get_Sfnt_Name func(
		face FT_Face,
		idx FT_UInt,
		aname *FT_SfntName) FT_Error

	FT_OpenType_Validate func(
		face FT_Face,
		validation_flags FT_UInt,
		BASE_table *FT_Bytes,
		GDEF_table *FT_Bytes,
		GPOS_table *FT_Bytes,
		GSUB_table *FT_Bytes,
		JSTF_table *FT_Bytes) FT_Error

	FT_OpenType_Free func(
		face FT_Face,
		table FT_Bytes)

	FT_TrueTypeGX_Validate func(
		face FT_Face,
		validation_flags FT_UInt,
		tables [9 + 1]FT_Bytes,
		table_length FT_UInt) FT_Error

	FT_TrueTypeGX_Free func(
		face FT_Face,
		table FT_Bytes)

	FT_ClassicKern_Validate func(
		face FT_Face,
		validation_flags FT_UInt,
		ckern_table *FT_Bytes) FT_Error

	FT_ClassicKern_Free func(
		face FT_Face,
		table FT_Bytes)

	FT_Get_PFR_Metrics func(
		face FT_Face,
		aoutline_resolution *FT_UInt,
		ametrics_resolution *FT_UInt,
		ametrics_x_scale *FT_Fixed,
		ametrics_y_scale *FT_Fixed) FT_Error

	FT_Get_PFR_Kerning func(
		face FT_Face,
		left FT_UInt,
		right FT_UInt,
		avector *FT_Vector) FT_Error

	FT_Get_PFR_Advance func(
		face FT_Face,
		gindex FT_UInt,
		aadvance *FT_Pos) FT_Error

	FT_Outline_GetInsideBorder func(
		outline *FT_Outline) FT_StrokerBorder

	FT_Outline_GetOutsideBorder func(
		outline *FT_Outline) FT_StrokerBorder

	FT_Stroker_New func(
		library FT_Library,
		astroker *FT_Stroker) FT_Error

	FT_Stroker_Set func(
		stroker FT_Stroker,
		radius FT_Fixed,
		line_cap FT_Stroker_LineCap,
		line_join FT_Stroker_LineJoin,
		miter_limit FT_Fixed)

	FT_Stroker_Rewind func(
		stroker FT_Stroker)

	FT_Stroker_ParseOutline func(
		stroker FT_Stroker,
		outline *FT_Outline,
		opened FT_Bool) FT_Error

	FT_Stroker_BeginSubPath func(
		stroker FT_Stroker,
		to *FT_Vector,
		open FT_Bool) FT_Error

	FT_Stroker_EndSubPath func(
		stroker FT_Stroker) FT_Error

	FT_Stroker_LineTo func(
		stroker FT_Stroker,
		to *FT_Vector) FT_Error

	FT_Stroker_ConicTo func(
		stroker FT_Stroker,
		control *FT_Vector,
		to *FT_Vector) FT_Error

	FT_Stroker_CubicTo func(
		stroker FT_Stroker,
		control1 *FT_Vector,
		control2 *FT_Vector,
		to *FT_Vector) FT_Error

	FT_Stroker_GetBorderCounts func(
		stroker FT_Stroker,
		border FT_StrokerBorder,
		anum_points *FT_UInt,
		anum_contours *FT_UInt) FT_Error

	FT_Stroker_ExportBorder func(
		stroker FT_Stroker,
		border FT_StrokerBorder,
		outline *FT_Outline)

	FT_Stroker_GetCounts func(
		stroker FT_Stroker,
		anum_points *FT_UInt,
		anum_contours *FT_UInt) FT_Error

	FT_Stroker_Export func(
		stroker FT_Stroker,
		outline *FT_Outline)

	FT_Stroker_Done func(
		stroker FT_Stroker)

	FT_Glyph_Stroke func(
		pglyph *FT_Glyph,
		stroker FT_Stroker,
		destroy FT_Bool) FT_Error

	FT_Glyph_StrokeBorder func(
		pglyph *FT_Glyph,
		stroker FT_Stroker,
		inside FT_Bool,
		destroy FT_Bool) FT_Error

	FT_GlyphSlot_Embolden func(
		slot FT_GlyphSlot)

	FT_GlyphSlot_Oblique func(
		slot FT_GlyphSlot)

	FT_Get_X11_Font_Format func(
		face FT_Face) *Char

	FT_Sin func(
		angle FT_Angle) FT_Fixed

	FT_Cos func(
		angle FT_Angle) FT_Fixed

	FT_Tan func(
		angle FT_Angle) FT_Fixed

	FT_Atan2 func(
		x FT_Fixed,
		y FT_Fixed) FT_Angle

	FT_Angle_Diff func(
		angle1 FT_Angle,
		angle2 FT_Angle) FT_Angle

	FT_Vector_Unit func(
		vec *FT_Vector,
		angle FT_Angle)

	FT_Vector_Rotate func(
		vec *FT_Vector,
		angle FT_Angle)

	FT_Vector_Length func(
		vec *FT_Vector) FT_Fixed

	FT_Vector_Polarize func(
		vec *FT_Vector,
		length *FT_Fixed,
		angle *FT_Angle)

	FT_Vector_From_Polar func(
		vec *FT_Vector,
		length FT_Fixed,
		angle FT_Angle)

	FT_Library_SetLcdFilter func(
		library FT_Library,
		filter FT_LcdFilter) FT_Error

	FT_Library_SetLcdFilterWeights func(
		library FT_Library,
		weights *Unsigned_char) FT_Error

	FT_Get_Gasp func(
		face FT_Face,
		ppem FT_UInt) FT_Int

	FT_Get_Advance func(
		face FT_Face,
		gindex FT_UInt,
		load_flags FT_Int32,
		padvance *FT_Fixed) FT_Error

	FT_Get_Advances func(
		face FT_Face,
		start FT_UInt,
		count FT_UInt,
		load_flags FT_Int32,
		padvances *FT_Fixed) FT_Error
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
	// {"FT_Done_FreeType", &FT_Done_FreeType},
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
	// {"FT_Init_FreeType", &FT_Init_FreeType},
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
	// {"FT_MulFix", &FT_MulFix},
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
	// {"FT_Render_Glyph", &FT_Render_Glyph},
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
	// {"TT_New_Context", &TT_New_Context},
	// {"TT_RunIns", &TT_RunIns},
}
