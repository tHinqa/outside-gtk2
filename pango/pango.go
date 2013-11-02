// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package pango provides API definitions for accessing
//libpango-1.0-0.dll and libpangocairo-1.0-0.dll.
package pango

import (
	"github.com/tHinqa/outside"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllApis(dllCairo, false, apiListCairo)
}

type Enum int

type BidiType Enum

const (
	BIDI_TYPE_L BidiType = iota
	BIDI_TYPE_LRE
	BIDI_TYPE_LRO
	BIDI_TYPE_R
	BIDI_TYPE_AL
	BIDI_TYPE_RLE
	BIDI_TYPE_RLO
	BIDI_TYPE_PDF
	BIDI_TYPE_EN
	BIDI_TYPE_ES
	BIDI_TYPE_ET
	BIDI_TYPE_AN
	BIDI_TYPE_CS
	BIDI_TYPE_NSM
	BIDI_TYPE_BN
	BIDI_TYPE_B
	BIDI_TYPE_S
	BIDI_TYPE_WS
	BIDI_TYPE_ON
)

var BidiTypeGetType func() O.Type

var (
	Break func(text string, length int, Analysis *Analysis,
		Attrs *LogAttr, Attrs_len int)

	BidiTypeForUnichar func(
		ch L.Unichar) BidiType
)

var DefaultBreak func(text string, length int,
	Analysis *Analysis, Attrs *LogAttr, AttrsLen int)

type Direction Enum

const (
	DIRECTION_LTR Direction = iota
	DIRECTION_RTL
	DIRECTION_TTB_LTR
	DIRECTION_TTB_RTL
	DIRECTION_WEAK_LTR
	DIRECTION_WEAK_RTL
	DIRECTION_NEUTRAL
)

var DirectionGetType func() O.Type

type EllipsizeMode Enum

const (
	ELLIPSIZE_NONE EllipsizeMode = iota
	ELLIPSIZE_START
	ELLIPSIZE_MIDDLE
	ELLIPSIZE_END
)

var EllipsizeModeGetType func() O.Type

type (
	EngineShape struct{}
	EngineLang  struct{}
)

type Engine struct{ parent O.Object }

var EngineGetType func() O.Type

type EngineInfo struct {
	Id         *T.Gchar
	EngineType *T.Gchar
	RenderType *T.Gchar
	Scripts    *EngineScriptInfo
	NScripts   int
}

type EngineScriptInfo struct {
	Script Script
	Langs  *T.Gchar
}

var (
	EngineLangGetType func() O.Type

	EngineShapeGetType func() O.Type

	ExtentsToPixels func(
		inclusive,
		nearest *Rectangle)
)

type IncludedModule struct {
	List   func(engines **EngineInfo, nEngines *int)
	Init   func(module *O.TypeModule)
	Exit   func()
	Create func(id *T.Char) *Engine
}

type (
	OTBuffer struct{}

	OTFeatureMap struct {
		FeatureName [5]T.Char
		PropertyBit T.Gulong
	}

	OTGlyph struct {
		Glyph      T.GUint32
		Properties uint
		Cluster    uint
		Component  uint16
		LigID      uint16
		Internal   uint
	}

	OTInfo struct{}

	OTRuleset struct{}

	OTRulesetDescription struct {
		Script              Script
		Language            *Language
		StaticGsubFeatures  *OTFeatureMap
		NStaticGsubFeatures uint
		StaticGposFeatures  *OTFeatureMap
		NStaticGposFeatures uint
		OtherFeatures       *OTFeatureMap
		NOtherFeatures      uint
	}

	OTTag T.GUint32
)

type OTTableType Enum

const (
	OT_TABLE_GSUB OTTableType = iota
	OT_TABLE_GPOS
)

var (
	ParseEnum func(
		t O.Type,
		str string,
		value *int,
		warn bool,
		possible_values **T.Char) bool

	ParseMarkup func(
		markup_text string,
		length int,
		AccelMarker L.Unichar,
		Attr_list **AttrList,
		text **T.Char,
		AccelChar *L.Unichar,
		error **L.Error) bool

	ParseStretch func(str string,
		stretch *Stretch, warn bool) bool

	ParseStyle func(str string,
		style *Style, warn bool) bool

	ParseVariant func(str string,
		variant *Variant, warn bool) bool

	ParseWeight func(str string,
		weight *Weight, warn bool) bool

	QuantizeLineGeometry func(thickness, position *int)
)

type Underline Enum

const (
	UNDERLINE_NONE Underline = iota
	UNDERLINE_SINGLE
	UNDERLINE_DOUBLE
	UNDERLINE_LOW
	UNDERLINE_ERROR
)

var UnderlineGetType func() O.Type

var (
	UnicharDirection func(
		ch L.Unichar) Direction

	UnitsFromDouble func(
		d float64) int

	UnitsToDouble func(
		i int) float64
)

type Variant Enum

const (
	VARIANT_NORMAL Variant = iota
	VARIANT_SMALL_CAPS
)

var VariantGetType func() O.Type

var (
	Version func() int

	VersionCheck func(
		requiredMajor, requiredMinor, requiredMicro int) string

	VersionString func() string
)

type Weight Enum

const (
	WEIGHT_THIN       Weight = 100
	WEIGHT_ULTRALIGHT Weight = 200
	WEIGHT_LIGHT      Weight = 300
	WEIGHT_BOOK       Weight = 380
	WEIGHT_NORMAL     Weight = 400
	WEIGHT_MEDIUM     Weight = 500
	WEIGHT_SEMIBOLD   Weight = 600
	WEIGHT_BOLD       Weight = 700
	WEIGHT_ULTRABOLD  Weight = 800
	WEIGHT_HEAVY      Weight = 900
	WEIGHT_ULTRAHEAVY Weight = 1000
)

var WeightGetType func() O.Type

type WrapMode Enum

const (
	WRAP_WORD WrapMode = iota
	WRAP_CHAR
	WRAP_WORD_CHAR
)

var WrapModeGetType func() O.Type

var dll = "libpango-1.0-0.dll"

var apiList = outside.Apis{
	{"pango_alignment_get_type", &AlignmentGetType},
	{"pango_attr_background_new", &AttrBackgroundNew},
	{"pango_attr_fallback_new", &AttrFallbackNew},
	{"pango_attr_family_new", &AttrFamilyNew},
	{"pango_attr_font_desc_new", &AttrFontDescNew},
	{"pango_attr_foreground_new", &AttrForegroundNew},
	{"pango_attr_gravity_hint_new", &AttrGravityHintNew},
	{"pango_attr_gravity_new", &AttrGravityNew},
	{"pango_attr_iterator_copy", &AttrIteratorCopy},
	{"pango_attr_iterator_destroy", &AttrIteratorDestroy},
	{"pango_attr_iterator_get", &AttrIteratorGet},
	{"pango_attr_iterator_get_attrs", &AttrIteratorGetAttrs},
	{"pango_attr_iterator_get_font", &AttrIteratorGetFont},
	{"pango_attr_iterator_next", &AttrIteratorNext},
	{"pango_attr_iterator_range", &AttrIteratorRange},
	{"pango_attr_language_new", &AttrLanguageNew},
	{"pango_attr_letter_spacing_new", &AttrLetterSpacingNew},
	{"pango_attr_list_change", &AttrListChange},
	{"pango_attr_list_copy", &AttrListCopy},
	{"pango_attr_list_filter", &AttrListFilter},
	{"pango_attr_list_get_iterator", &AttrListGetIterator},
	{"pango_attr_list_get_type", &AttrListGetType},
	{"pango_attr_list_insert", &AttrListInsert},
	{"pango_attr_list_insert_before", &AttrListInsertBefore},
	{"pango_attr_list_new", &AttrListNew},
	{"pango_attr_list_ref", &AttrListRef},
	{"pango_attr_list_splice", &AttrListSplice},
	{"pango_attr_list_unref", &AttrListUnref},
	{"pango_attr_rise_new", &AttrRiseNew},
	{"pango_attr_scale_new", &AttrScaleNew},
	{"pango_attr_shape_new", &AttrShapeNew},
	{"pango_attr_shape_new_with_data", &AttrShapeNewWithData},
	{"pango_attr_size_new", &AttrSizeNew},
	{"pango_attr_size_new_absolute", &AttrSizeNewAbsolute},
	{"pango_attr_stretch_new", &AttrStretchNew},
	{"pango_attr_strikethrough_color_new", &AttrStrikethroughColorNew},
	{"pango_attr_strikethrough_new", &AttrStrikethroughNew},
	{"pango_attr_style_new", &AttrStyleNew},
	{"pango_attr_type_get_name", &AttrTypeGetName},
	{"pango_attr_type_get_type", &AttrTypeGetType},
	{"pango_attr_type_register", &AttrTypeRegister},
	{"pango_attr_underline_color_new", &AttrUnderlineColorNew},
	{"pango_attr_underline_new", &AttrUnderlineNew},
	{"pango_attr_variant_new", &AttrVariantNew},
	{"pango_attr_weight_new", &AttrWeightNew},
	{"pango_attribute_copy", &AttributeCopy},
	{"pango_attribute_destroy", &AttributeDestroy},
	{"pango_attribute_equal", &AttributeEqual},
	{"pango_attribute_init", &AttributeInit},
	{"pango_bidi_type_for_unichar", &BidiTypeForUnichar},
	{"pango_bidi_type_get_type", &BidiTypeGetType},
	{"pango_break", &Break},
	{"pango_color_copy", &ColorCopy},
	{"pango_color_free", &ColorFree},
	{"pango_color_get_type", &ColorGetType},
	{"pango_color_parse", &ColorParse},
	{"pango_color_to_string", &ColorToString},
	{"pango_config_key_get", &ConfigKeyGet},
	{"pango_context_get_base_dir", &ContextGetBaseDir},
	{"pango_context_get_base_gravity", &ContextGetBaseGravity},
	{"pango_context_get_font_description", &ContextGetFontDescription},
	{"pango_context_get_font_map", &ContextGetFontMap},
	{"pango_context_get_gravity", &ContextGetGravity},
	{"pango_context_get_gravity_hint", &ContextGetGravityHint},
	{"pango_context_get_language", &ContextGetLanguage},
	{"pango_context_get_matrix", &ContextGetMatrix},
	{"pango_context_get_metrics", &ContextGetMetrics},
	{"pango_context_get_type", &ContextGetType},
	{"pango_context_list_families", &ContextListFamilies},
	{"pango_context_load_font", &ContextLoadFont},
	{"pango_context_load_fontset", &ContextLoadFontset},
	{"pango_context_new", &ContextNew},
	{"pango_context_set_base_dir", &ContextSetBaseDir},
	{"pango_context_set_base_gravity", &ContextSetBaseGravity},
	{"pango_context_set_font_description", &ContextSetFontDescription},
	{"pango_context_set_font_map", &ContextSetFontMap},
	{"pango_context_set_gravity_hint", &ContextSetGravityHint},
	{"pango_context_set_language", &ContextSetLanguage},
	{"pango_context_set_matrix", &ContextSetMatrix},
	{"pango_coverage_copy", &CoverageCopy},
	{"pango_coverage_from_bytes", &CoverageFromBytes},
	{"pango_coverage_get", &CoverageGet},
	{"pango_coverage_level_get_type", &CoverageLevelGetType},
	{"pango_coverage_max", &CoverageMax},
	{"pango_coverage_new", &CoverageNew},
	{"pango_coverage_ref", &CoverageRef},
	{"pango_coverage_set", &CoverageSet},
	{"pango_coverage_to_bytes", &CoverageToBytes},
	{"pango_coverage_unref", &CoverageUnref},
	{"pango_default_break", &DefaultBreak},
	{"pango_direction_get_type", &DirectionGetType},
	{"pango_ellipsize_mode_get_type", &EllipsizeModeGetType},
	{"pango_engine_get_type", &EngineGetType},
	{"pango_engine_lang_get_type", &EngineLangGetType},
	{"pango_engine_shape_get_type", &EngineShapeGetType},
	{"pango_extents_to_pixels", &ExtentsToPixels},
	{"pango_find_base_dir", &FindBaseDir},
	{"pango_find_map", &FindMap},
	{"pango_find_paragraph_boundary", &FindParagraphBoundary},
	{"pango_font_describe", &FontDescribe},
	{"pango_font_describe_with_absolute_size", &FontDescribeWithAbsoluteSize},
	{"pango_font_description_better_match", &FontDescriptionBetterMatch},
	{"pango_font_description_copy", &FontDescriptionCopy},
	{"pango_font_description_copy_static", &FontDescriptionCopyStatic},
	{"pango_font_description_equal", &FontDescriptionEqual},
	{"pango_font_description_free", &FontDescriptionFree},
	{"pango_font_description_from_string", &FontDescriptionFromString},
	{"pango_font_description_get_family", &FontDescriptionGetFamily},
	{"pango_font_description_get_gravity", &FontDescriptionGetGravity},
	{"pango_font_description_get_set_fields", &FontDescriptionGetSetFields},
	{"pango_font_description_get_size", &FontDescriptionGetSize},
	{"pango_font_description_get_size_is_absolute", &FontDescriptionGetSizeIsAbsolute},
	{"pango_font_description_get_stretch", &FontDescriptionGetStretch},
	{"pango_font_description_get_style", &FontDescriptionGetStyle},
	{"pango_font_description_get_type", &FontDescriptionGetType},
	{"pango_font_description_get_variant", &FontDescriptionGetVariant},
	{"pango_font_description_get_weight", &FontDescriptionGetWeight},
	{"pango_font_description_hash", &FontDescriptionHash},
	{"pango_font_description_merge", &FontDescriptionMerge},
	{"pango_font_description_merge_static", &FontDescriptionMergeStatic},
	{"pango_font_description_new", &FontDescriptionNew},
	{"pango_font_description_set_absolute_size", &FontDescriptionSetAbsoluteSize},
	{"pango_font_description_set_family", &FontDescriptionSetFamily},
	{"pango_font_description_set_family_static", &FontDescriptionSetFamilyStatic},
	{"pango_font_description_set_gravity", &FontDescriptionSetGravity},
	{"pango_font_description_set_size", &FontDescriptionSetSize},
	{"pango_font_description_set_stretch", &FontDescriptionSetStretch},
	{"pango_font_description_set_style", &FontDescriptionSetStyle},
	{"pango_font_description_set_variant", &FontDescriptionSetVariant},
	{"pango_font_description_set_weight", &FontDescriptionSetWeight},
	{"pango_font_description_to_filename", &FontDescriptionToFilename},
	{"pango_font_description_to_string", &FontDescriptionToString},
	{"pango_font_description_unset_fields", &FontDescriptionUnsetFields},
	{"pango_font_descriptions_free", &FontDescriptionsFree},
	{"pango_font_face_describe", &FontFaceDescribe},
	{"pango_font_face_get_face_name", &FontFaceGetFaceName},
	{"pango_font_face_get_type", &FontFaceGetType},
	{"pango_font_face_is_synthesized", &FontFaceIsSynthesized},
	{"pango_font_face_list_sizes", &FontFaceListSizes},
	{"pango_font_family_get_name", &FontFamilyGetName},
	{"pango_font_family_get_type", &FontFamilyGetType},
	{"pango_font_family_is_monospace", &FontFamilyIsMonospace},
	{"pango_font_family_list_faces", &FontFamilyListFaces},
	{"pango_font_find_shaper", &FontFindShaper},
	{"pango_font_get_coverage", &FontGetCoverage},
	{"pango_font_get_font_map", &FontGetFontMap},
	{"pango_font_get_glyph_extents", &FontGetGlyphExtents},
	{"pango_font_get_metrics", &FontGetMetrics},
	{"pango_font_get_type", &FontGetType},
	{"pango_font_map_create_context", &FontMapCreateContext},
	{"pango_font_map_get_shape_engine_type", &FontMapGetShapeEngineType},
	{"pango_font_map_get_type", &FontMapGetType},
	{"pango_font_map_list_families", &FontMapListFamilies},
	{"pango_font_map_load_font", &FontMapLoadFont},
	{"pango_font_map_load_fontset", &FontMapLoadFontset},
	{"pango_font_mask_get_type", &FontMaskGetType},
	{"pango_font_metrics_get_approximate_char_width", &FontMetricsGetApproximateCharWidth},
	{"pango_font_metrics_get_approximate_digit_width", &FontMetricsGetApproximateDigitWidth},
	{"pango_font_metrics_get_ascent", &FontMetricsGetAscent},
	{"pango_font_metrics_get_descent", &FontMetricsGetDescent},
	{"pango_font_metrics_get_strikethrough_position", &FontMetricsGetStrikethroughPosition},
	{"pango_font_metrics_get_strikethrough_thickness", &FontMetricsGetStrikethroughThickness},
	{"pango_font_metrics_get_type", &FontMetricsGetType},
	{"pango_font_metrics_get_underline_position", &FontMetricsGetUnderlinePosition},
	{"pango_font_metrics_get_underline_thickness", &FontMetricsGetUnderlineThickness},
	{"pango_font_metrics_new", &FontMetricsNew},
	{"pango_font_metrics_ref", &FontMetricsRef},
	{"pango_font_metrics_unref", &FontMetricsUnref},
	{"pango_fontset_foreach", &FontsetForeach},
	{"pango_fontset_get_font", &FontsetGetFont},
	{"pango_fontset_get_metrics", &FontsetGetMetrics},
	{"pango_fontset_get_type", &FontsetGetType},
	{"pango_fontset_simple_append", &FontsetSimpleAppend},
	{"pango_fontset_simple_get_type", &FontsetSimpleGetType},
	{"pango_fontset_simple_new", &FontsetSimpleNew},
	{"pango_fontset_simple_size", &FontsetSimpleSize},
	{"pango_get_lib_subdirectory", &GetLibSubdirectory},
	{"pango_get_log_attrs", &GetLogAttrs},
	{"pango_get_mirror_char", &GetMirrorChar},
	{"pango_get_sysconf_subdirectory", &GetSysconfSubdirectory},
	{"pango_glyph_item_apply_attrs", &GlyphItemApplyAttrs},
	{"pango_glyph_item_copy", &GlyphItemCopy},
	{"pango_glyph_item_free", &GlyphItemFree},
	{"pango_glyph_item_get_logical_widths", &GlyphItemGetLogicalWidths},
	{"pango_glyph_item_get_type", &GlyphItemGetType},
	{"pango_glyph_item_iter_copy", &GlyphItemIterCopy},
	{"pango_glyph_item_iter_free", &GlyphItemIterFree},
	{"pango_glyph_item_iter_get_type", &GlyphItemIterGetType},
	{"pango_glyph_item_iter_init_end", &GlyphItemIterInitEnd},
	{"pango_glyph_item_iter_init_start", &GlyphItemIterInitStart},
	{"pango_glyph_item_iter_next_cluster", &GlyphItemIterNextCluster},
	{"pango_glyph_item_iter_prev_cluster", &GlyphItemIterPrevCluster},
	{"pango_glyph_item_letter_space", &GlyphItemLetterSpace},
	{"pango_glyph_item_split", &GlyphItemSplit},
	{"pango_glyph_string_copy", &GlyphStringCopy},
	{"pango_glyph_string_extents", &GlyphStringExtents},
	{"pango_glyph_string_extents_range", &GlyphStringExtentsRange},
	{"pango_glyph_string_free", &GlyphStringFree},
	{"pango_glyph_string_get_logical_widths", &GlyphStringGetLogicalWidths},
	{"pango_glyph_string_get_type", &GlyphStringGetType},
	{"pango_glyph_string_get_width", &GlyphStringGetWidth},
	{"pango_glyph_string_index_to_x", &GlyphStringIndexToX},
	{"pango_glyph_string_new", &GlyphStringNew},
	{"pango_glyph_string_set_size", &GlyphStringSetSize},
	{"pango_glyph_string_x_to_index", &GlyphStringXToIndex},
	{"pango_gravity_get_for_matrix", &GravityGetForMatrix},
	{"pango_gravity_get_for_script", &GravityGetForScript},
	{"pango_gravity_get_for_script_and_width", &GravityGetForScriptAndWidth},
	{"pango_gravity_get_type", &GravityGetType},
	{"pango_gravity_hint_get_type", &GravityHintGetType},
	{"pango_gravity_to_rotation", &GravityToRotation},
	{"pango_is_zero_width", &IsZeroWidth},
	{"pango_item_copy", &ItemCopy},
	{"pango_item_free", &ItemFree},
	{"pango_item_get_type", &ItemGetType},
	{"pango_item_new", &ItemNew},
	{"pango_item_split", &ItemSplit},
	{"pango_itemize", &Itemize},
	{"pango_itemize_with_base_dir", &ItemizeWithBaseDir},
	{"pango_language_from_string", &LanguageFromString},
	{"pango_language_get_default", &LanguageGetDefault},
	{"pango_language_get_sample_string", &LanguageGetSampleString},
	{"pango_language_get_scripts", &LanguageGetScripts},
	{"pango_language_get_type", &LanguageGetType},
	{"pango_language_includes_script", &LanguageIncludesScript},
	{"pango_language_matches", &LanguageMatches},
	{"pango_language_to_string", &LanguageToString},
	{"pango_layout_context_changed", &LayoutContextChanged},
	{"pango_layout_copy", &LayoutCopy},
	{"pango_layout_get_alignment", &LayoutGetAlignment},
	{"pango_layout_get_attributes", &LayoutGetAttributes},
	{"pango_layout_get_auto_dir", &LayoutGetAutoDir},
	{"pango_layout_get_baseline", &LayoutGetBaseline},
	{"pango_layout_get_character_count", &LayoutGetCharacterCount},
	{"pango_layout_get_context", &LayoutGetContext},
	{"pango_layout_get_cursor_pos", &LayoutGetCursorPos},
	{"pango_layout_get_ellipsize", &LayoutGetEllipsize},
	{"pango_layout_get_extents", &LayoutGetExtents},
	{"pango_layout_get_font_description", &LayoutGetFontDescription},
	{"pango_layout_get_height", &LayoutGetHeight},
	{"pango_layout_get_indent", &LayoutGetIndent},
	{"pango_layout_get_iter", &LayoutGetIter},
	{"pango_layout_get_justify", &LayoutGetJustify},
	{"pango_layout_get_line", &LayoutGetLine},
	{"pango_layout_get_line_count", &LayoutGetLineCount},
	{"pango_layout_get_line_readonly", &LayoutGetLineReadonly},
	{"pango_layout_get_lines", &LayoutGetLines},
	{"pango_layout_get_lines_readonly", &LayoutGetLinesReadonly},
	{"pango_layout_get_log_attrs", &LayoutGetLogAttrs},
	{"pango_layout_get_log_attrs_readonly", &LayoutGetLogAttrsReadonly},
	{"pango_layout_get_pixel_extents", &LayoutGetPixelExtents},
	{"pango_layout_get_pixel_size", &LayoutGetPixelSize},
	{"pango_layout_get_single_paragraph_mode", &LayoutGetSingleParagraphMode},
	{"pango_layout_get_size", &LayoutGetSize},
	{"pango_layout_get_spacing", &LayoutGetSpacing},
	{"pango_layout_get_tabs", &LayoutGetTabs},
	{"pango_layout_get_text", &LayoutGetText},
	{"pango_layout_get_type", &LayoutGetType},
	{"pango_layout_get_unknown_glyphs_count", &LayoutGetUnknownGlyphsCount},
	{"pango_layout_get_width", &LayoutGetWidth},
	{"pango_layout_get_wrap", &LayoutGetWrap},
	{"pango_layout_index_to_line_x", &LayoutIndexToLineX},
	{"pango_layout_index_to_pos", &LayoutIndexToPos},
	{"pango_layout_is_ellipsized", &LayoutIsEllipsized},
	{"pango_layout_is_wrapped", &LayoutIsWrapped},
	{"pango_layout_iter_at_last_line", &LayoutIterAtLastLine},
	{"pango_layout_iter_copy", &LayoutIterCopy},
	{"pango_layout_iter_free", &LayoutIterFree},
	{"pango_layout_iter_get_baseline", &LayoutIterGetBaseline},
	{"pango_layout_iter_get_char_extents", &LayoutIterGetCharExtents},
	{"pango_layout_iter_get_cluster_extents", &LayoutIterGetClusterExtents},
	{"pango_layout_iter_get_index", &LayoutIterGetIndex},
	{"pango_layout_iter_get_layout", &LayoutIterGetLayout},
	{"pango_layout_iter_get_layout_extents", &LayoutIterGetLayoutExtents},
	{"pango_layout_iter_get_line", &LayoutIterGetLine},
	{"pango_layout_iter_get_line_extents", &LayoutIterGetLineExtents},
	{"pango_layout_iter_get_line_readonly", &LayoutIterGetLineReadonly},
	{"pango_layout_iter_get_line_yrange", &LayoutIterGetLineYrange},
	{"pango_layout_iter_get_run", &LayoutIterGetRun},
	{"pango_layout_iter_get_run_extents", &LayoutIterGetRunExtents},
	{"pango_layout_iter_get_run_readonly", &LayoutIterGetRunReadonly},
	{"pango_layout_iter_get_type", &LayoutIterGetType},
	{"pango_layout_iter_next_char", &LayoutIterNextChar},
	{"pango_layout_iter_next_cluster", &LayoutIterNextCluster},
	{"pango_layout_iter_next_line", &LayoutIterNextLine},
	{"pango_layout_iter_next_run", &LayoutIterNextRun},
	{"pango_layout_line_get_extents", &LayoutLineGetExtents},
	{"pango_layout_line_get_pixel_extents", &LayoutLineGetPixelExtents},
	{"pango_layout_line_get_type", &LayoutLineGetType},
	{"pango_layout_line_get_x_ranges", &LayoutLineGetXRanges},
	{"pango_layout_line_index_to_x", &LayoutLineIndexToX},
	{"pango_layout_line_ref", &LayoutLineRef},
	{"pango_layout_line_unref", &LayoutLineUnref},
	{"pango_layout_line_x_to_index", &LayoutLineXToIndex},
	{"pango_layout_move_cursor_visually", &LayoutMoveCursorVisually},
	{"pango_layout_new", &LayoutNew},
	{"pango_layout_set_alignment", &LayoutSetAlignment},
	{"pango_layout_set_attributes", &LayoutSetAttributes},
	{"pango_layout_set_auto_dir", &LayoutSetAutoDir},
	{"pango_layout_set_ellipsize", &LayoutSetEllipsize},
	{"pango_layout_set_font_description", &LayoutSetFontDescription},
	{"pango_layout_set_height", &LayoutSetHeight},
	{"pango_layout_set_indent", &LayoutSetIndent},
	{"pango_layout_set_justify", &LayoutSetJustify},
	{"pango_layout_set_markup", &LayoutSetMarkup},
	{"pango_layout_set_markup_with_accel", &LayoutSetMarkupWithAccel},
	{"pango_layout_set_single_paragraph_mode", &LayoutSetSingleParagraphMode},
	{"pango_layout_set_spacing", &LayoutSetSpacing},
	{"pango_layout_set_tabs", &LayoutSetTabs},
	{"pango_layout_set_text", &LayoutSetText},
	{"pango_layout_set_width", &LayoutSetWidth},
	{"pango_layout_set_wrap", &LayoutSetWrap},
	{"pango_layout_xy_to_index", &LayoutXyToIndex},
	{"pango_log2vis_get_embedding_levels", &Log2visGetEmbeddingLevels},
	{"pango_lookup_aliases", &LookupAliases},
	{"pango_map_get_engine", &MapGetEngine},
	{"pango_map_get_engines", &MapGetEngines},
	{"pango_matrix_concat", &MatrixConcat},
	{"pango_matrix_copy", &MatrixCopy},
	{"pango_matrix_free", &MatrixFree},
	{"pango_matrix_get_font_scale_factor", &MatrixGetFontScaleFactor},
	{"pango_matrix_get_type", &MatrixGetType},
	{"pango_matrix_rotate", &MatrixRotate},
	{"pango_matrix_scale", &MatrixScale},
	{"pango_matrix_transform_distance", &MatrixTransformDistance},
	{"pango_matrix_transform_pixel_rectangle", &MatrixTransformPixelRectangle},
	{"pango_matrix_transform_point", &MatrixTransformPoint},
	{"pango_matrix_transform_rectangle", &MatrixTransformRectangle},
	{"pango_matrix_translate", &MatrixTranslate},
	{"pango_module_register", &ModuleRegister},
	{"pango_parse_enum", &ParseEnum},
	{"pango_parse_markup", &ParseMarkup},
	{"pango_parse_stretch", &ParseStretch},
	{"pango_parse_style", &ParseStyle},
	{"pango_parse_variant", &ParseVariant},
	{"pango_parse_weight", &ParseWeight},
	{"pango_quantize_line_geometry", &QuantizeLineGeometry},
	{"pango_read_line", &ReadLine},
	{"pango_render_part_get_type", &RenderPartGetType},
	{"pango_renderer_activate", &RendererActivate},
	{"pango_renderer_deactivate", &RendererDeactivate},
	{"pango_renderer_draw_error_underline", &RendererDrawErrorUnderline},
	{"pango_renderer_draw_glyph", &RendererDrawGlyph},
	{"pango_renderer_draw_glyph_item", &RendererDrawGlyphItem},
	{"pango_renderer_draw_glyphs", &RendererDrawGlyphs},
	{"pango_renderer_draw_layout", &RendererDrawLayout},
	{"pango_renderer_draw_layout_line", &RendererDrawLayoutLine},
	{"pango_renderer_draw_rectangle", &RendererDrawRectangle},
	{"pango_renderer_draw_trapezoid", &RendererDrawTrapezoid},
	{"pango_renderer_get_color", &RendererGetColor},
	{"pango_renderer_get_layout", &RendererGetLayout},
	{"pango_renderer_get_layout_line", &RendererGetLayoutLine},
	{"pango_renderer_get_matrix", &RendererGetMatrix},
	{"pango_renderer_get_type", &RendererGetType},
	{"pango_renderer_part_changed", &RendererPartChanged},
	{"pango_renderer_set_color", &RendererSetColor},
	{"pango_renderer_set_matrix", &RendererSetMatrix},
	{"pango_reorder_items", &ReorderItems},
	{"pango_scan_int", &ScanInt},
	{"pango_scan_string", &ScanString},
	{"pango_scan_word", &ScanWord},
	{"pango_script_for_unichar", &ScriptForUnichar},
	{"pango_script_get_sample_language", &ScriptGetSampleLanguage},
	{"pango_script_get_type", &ScriptGetType},
	{"pango_script_iter_free", &ScriptIterFree},
	{"pango_script_iter_get_range", &ScriptIterGetRange},
	{"pango_script_iter_new", &ScriptIterNew},
	{"pango_script_iter_next", &ScriptIterNext},
	{"pango_shape", &Shape},
	{"pango_skip_space", &SkipSpace},
	{"pango_split_file_list", &SplitFileList},
	{"pango_stretch_get_type", &StretchGetType},
	{"pango_style_get_type", &StyleGetType},
	{"pango_tab_align_get_type", &TabAlignGetType},
	{"pango_tab_array_copy", &TabArrayCopy},
	{"pango_tab_array_free", &TabArrayFree},
	{"pango_tab_array_get_positions_in_pixels", &TabArrayGetPositionsInPixels},
	{"pango_tab_array_get_size", &TabArrayGetSize},
	{"pango_tab_array_get_tab", &TabArrayGetTab},
	{"pango_tab_array_get_tabs", &TabArrayGetTabs},
	{"pango_tab_array_get_type", &TabArrayGetType},
	{"pango_tab_array_new", &TabArrayNew},
	{"pango_tab_array_new_with_positions", &TabArrayNewWithPositions},
	{"pango_tab_array_resize", &TabArrayResize},
	{"pango_tab_array_set_tab", &TabArraySetTab},
	{"pango_trim_string", &TrimString},
	{"pango_underline_get_type", &UnderlineGetType},
	{"pango_unichar_direction", &UnicharDirection},
	{"pango_units_from_double", &UnitsFromDouble},
	{"pango_units_to_double", &UnitsToDouble},
	{"pango_variant_get_type", &VariantGetType},
	{"pango_version", &Version},
	{"pango_version_check", &VersionCheck},
	{"pango_version_string", &VersionString},
	{"pango_weight_get_type", &WeightGetType},
	{"pango_wrap_mode_get_type", &WrapModeGetType},
}

var dllCairo = "libpangocairo-1.0-0.dll"

var apiListCairo = outside.Apis{
	{"pango_cairo_context_get_font_options", &CairoContextGetFontOptions},
	{"pango_cairo_context_get_resolution", &CairoContextGetResolution},
	{"pango_cairo_context_get_shape_renderer", &CairoContextGetShapeRenderer},
	{"pango_cairo_context_set_font_options", &CairoContextSetFontOptions},
	{"pango_cairo_context_set_resolution", &CairoContextSetResolution},
	{"pango_cairo_context_set_shape_renderer", &CairoContextSetShapeRenderer},
	{"pango_cairo_create_context", &CairoCreateContext},
	{"pango_cairo_create_layout", &CairoCreateLayout},
	{"pango_cairo_error_underline_path", &CairoErrorUnderlinePath},
	{"pango_cairo_font_get_scaled_font", &CairoFontGetScaledFont},
	{"pango_cairo_font_get_type", &CairoFontGetType},
	{"pango_cairo_font_map_create_context", &CairoFontMapCreateContext},
	{"pango_cairo_font_map_get_default", &CairoFontMapGetDefault},
	{"pango_cairo_font_map_get_font_type", &CairoFontMapGetFontType},
	{"pango_cairo_font_map_get_resolution", &CairoFontMapGetResolution},
	{"pango_cairo_font_map_get_type", &CairoFontMapGetType},
	{"pango_cairo_font_map_new", &CairoFontMapNew},
	{"pango_cairo_font_map_new_for_font_type", &CairoFontMapNewForFontType},
	{"pango_cairo_font_map_set_default", &CairoFontMapSetDefault},
	{"pango_cairo_font_map_set_resolution", &CairoFontMapSetResolution},
	{"pango_cairo_glyph_string_path", &CairoGlyphStringPath},
	{"pango_cairo_layout_line_path", &CairoLayoutLinePath},
	{"pango_cairo_layout_path", &CairoLayoutPath},
	// Undocumented {"pango_cairo_renderer_get_type", &CairoRendererGetType},
	{"pango_cairo_show_error_underline", &CairoShowErrorUnderline},
	{"pango_cairo_show_glyph_item", &CairoShowGlyphItem},
	{"pango_cairo_show_glyph_string", &CairoShowGlyphString},
	{"pango_cairo_show_layout", &CairoShowLayout},
	{"pango_cairo_show_layout_line", &CairoShowLayoutLine},
	{"pango_cairo_update_context", &CairoUpdateContext},
	{"pango_cairo_update_layout", &CairoUpdateLayout},
}
