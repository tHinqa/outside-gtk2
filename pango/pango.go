// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package pango provides API definitions for accessing
//libpango-1.0-0.dll and libpangocairo-1.0-0.dll.
package pango

import (
	"github.com/tHinqa/outside"
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
		ch T.Gunichar) BidiType
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
	Init   func(module *T.GTypeModule)
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
		warn T.Gboolean,
		possible_values **T.Char) T.Gboolean

	ParseMarkup func(
		markup_text string,
		length int,
		AccelMarker T.Gunichar,
		Attr_list **AttrList,
		text **T.Char,
		AccelChar *T.Gunichar,
		error **T.GError) T.Gboolean

	ParseStretch func(str string,
		stretch *Stretch, warn T.Gboolean) T.Gboolean

	ParseStyle func(str string,
		style *Style, warn T.Gboolean) T.Gboolean

	ParseVariant func(str string,
		variant *Variant, warn T.Gboolean) T.Gboolean

	ParseWeight func(str string,
		weight *Weight, warn T.Gboolean) T.Gboolean

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
		ch T.Gunichar) Direction

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
	{"pango_attr_iterator_copy", &attrIteratorCopy},
	{"pango_attr_iterator_destroy", &attrIteratorDestroy},
	{"pango_attr_iterator_get", &attrIteratorGet},
	{"pango_attr_iterator_get_attrs", &attrIteratorGetAttrs},
	{"pango_attr_iterator_get_font", &attrIteratorGetFont},
	{"pango_attr_iterator_next", &attrIteratorNext},
	{"pango_attr_iterator_range", &attrIteratorRange},
	{"pango_attr_language_new", &AttrLanguageNew},
	{"pango_attr_letter_spacing_new", &AttrLetterSpacingNew},
	{"pango_attr_list_change", &attrListChange},
	{"pango_attr_list_copy", &attrListCopy},
	{"pango_attr_list_filter", &attrListFilter},
	{"pango_attr_list_get_iterator", &attrListGetIterator},
	{"pango_attr_list_get_type", &AttrListGetType},
	{"pango_attr_list_insert", &attrListInsert},
	{"pango_attr_list_insert_before", &attrListInsertBefore},
	{"pango_attr_list_new", &AttrListNew},
	{"pango_attr_list_ref", &attrListRef},
	{"pango_attr_list_splice", &attrListSplice},
	{"pango_attr_list_unref", &attrListUnref},
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
	{"pango_attribute_copy", &attributeCopy},
	{"pango_attribute_destroy", &attributeDestroy},
	{"pango_attribute_equal", &attributeEqual},
	{"pango_attribute_init", &attributeInit},
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
	{"pango_font_description_better_match", &fontDescriptionBetterMatch},
	{"pango_font_description_copy", &fontDescriptionCopy},
	{"pango_font_description_copy_static", &fontDescriptionCopyStatic},
	{"pango_font_description_equal", &fontDescriptionEqual},
	{"pango_font_description_free", &fontDescriptionFree},
	{"pango_font_description_from_string", &FontDescriptionFromString},
	{"pango_font_description_get_family", &fontDescriptionGetFamily},
	{"pango_font_description_get_gravity", &fontDescriptionGetGravity},
	{"pango_font_description_get_set_fields", &fontDescriptionGetSetFields},
	{"pango_font_description_get_size", &fontDescriptionGetSize},
	{"pango_font_description_get_size_is_absolute", &fontDescriptionGetSizeIsAbsolute},
	{"pango_font_description_get_stretch", &fontDescriptionGetStretch},
	{"pango_font_description_get_style", &fontDescriptionGetStyle},
	{"pango_font_description_get_type", &FontDescriptionGetType},
	{"pango_font_description_get_variant", &fontDescriptionGetVariant},
	{"pango_font_description_get_weight", &fontDescriptionGetWeight},
	{"pango_font_description_hash", &fontDescriptionHash},
	{"pango_font_description_merge", &fontDescriptionMerge},
	{"pango_font_description_merge_static", &fontDescriptionMergeStatic},
	{"pango_font_description_new", &FontDescriptionNew},
	{"pango_font_description_set_absolute_size", &fontDescriptionSetAbsoluteSize},
	{"pango_font_description_set_family", &fontDescriptionSetFamily},
	{"pango_font_description_set_family_static", &fontDescriptionSetFamilyStatic},
	{"pango_font_description_set_gravity", &fontDescriptionSetGravity},
	{"pango_font_description_set_size", &fontDescriptionSetSize},
	{"pango_font_description_set_stretch", &fontDescriptionSetStretch},
	{"pango_font_description_set_style", &fontDescriptionSetStyle},
	{"pango_font_description_set_variant", &fontDescriptionSetVariant},
	{"pango_font_description_set_weight", &fontDescriptionSetWeight},
	{"pango_font_description_to_filename", &fontDescriptionToFilename},
	{"pango_font_description_to_string", &fontDescriptionToString},
	{"pango_font_description_unset_fields", &fontDescriptionUnsetFields},
	{"pango_font_descriptions_free", &FontDescriptionsFree},
	{"pango_font_face_describe", &fontFaceDescribe},
	{"pango_font_face_get_face_name", &fontFaceGetFaceName},
	{"pango_font_face_get_type", &FontFaceGetType},
	{"pango_font_face_is_synthesized", &fontFaceIsSynthesized},
	{"pango_font_face_list_sizes", &fontFaceListSizes},
	{"pango_font_family_get_name", &fontFamilyGetName},
	{"pango_font_family_get_type", &FontFamilyGetType},
	{"pango_font_family_is_monospace", &fontFamilyIsMonospace},
	{"pango_font_family_list_faces", &fontFamilyListFaces},
	{"pango_font_find_shaper", &fontFindShaper},
	{"pango_font_get_coverage", &fontGetCoverage},
	{"pango_font_get_font_map", &fontGetFontMap},
	{"pango_font_get_glyph_extents", &FontGetGlyphExtents},
	{"pango_font_get_metrics", &fontGetMetrics},
	{"pango_font_get_type", &FontGetType},
	{"pango_font_map_create_context", &fontMapCreateContext},
	{"pango_font_map_get_shape_engine_type", &fontMapGetShapeEngineType},
	{"pango_font_map_get_type", &FontMapGetType},
	{"pango_font_map_list_families", &fontMapListFamilies},
	{"pango_font_map_load_font", &fontMapLoadFont},
	{"pango_font_map_load_fontset", &fontMapLoadFontset},
	{"pango_font_mask_get_type", &FontMaskGetType},
	{"pango_font_metrics_get_approximate_char_width", &fontMetricsGetApproximateCharWidth},
	{"pango_font_metrics_get_approximate_digit_width", &fontMetricsGetApproximateDigitWidth},
	{"pango_font_metrics_get_ascent", &fontMetricsGetAscent},
	{"pango_font_metrics_get_descent", &fontMetricsGetDescent},
	{"pango_font_metrics_get_strikethrough_position", &fontMetricsGetStrikethroughPosition},
	{"pango_font_metrics_get_strikethrough_thickness", &fontMetricsGetStrikethroughThickness},
	{"pango_font_metrics_get_type", &FontMetricsGetType},
	{"pango_font_metrics_get_underline_position", &fontMetricsGetUnderlinePosition},
	{"pango_font_metrics_get_underline_thickness", &fontMetricsGetUnderlineThickness},
	{"pango_font_metrics_new", &FontMetricsNew},
	{"pango_font_metrics_ref", &fontMetricsRef},
	{"pango_font_metrics_unref", &fontMetricsUnref},
	{"pango_fontset_foreach", &fontsetForeach},
	{"pango_fontset_get_font", &fontsetGetFont},
	{"pango_fontset_get_metrics", &fontsetGetMetrics},
	{"pango_fontset_get_type", &FontsetGetType},
	{"pango_fontset_simple_append", &fontsetSimpleAppend},
	{"pango_fontset_simple_get_type", &FontsetSimpleGetType},
	{"pango_fontset_simple_new", &FontsetSimpleNew},
	{"pango_fontset_simple_size", &fontsetSimpleSize},
	{"pango_get_lib_subdirectory", &GetLibSubdirectory},
	{"pango_get_log_attrs", &GetLogAttrs},
	{"pango_get_mirror_char", &GetMirrorChar},
	{"pango_get_sysconf_subdirectory", &GetSysconfSubdirectory},
	{"pango_glyph_item_apply_attrs", &glyphItemApplyAttrs},
	{"pango_glyph_item_copy", &glyphItemCopy},
	{"pango_glyph_item_free", &glyphItemFree},
	{"pango_glyph_item_get_logical_widths", &glyphItemGetLogicalWidths},
	{"pango_glyph_item_get_type", &GlyphItemGetType},
	{"pango_glyph_item_iter_copy", &glyphItemIterCopy},
	{"pango_glyph_item_iter_free", &glyphItemIterFree},
	{"pango_glyph_item_iter_get_type", &GlyphItemIterGetType},
	{"pango_glyph_item_iter_init_end", &glyphItemIterInitEnd},
	{"pango_glyph_item_iter_init_start", &glyphItemIterInitStart},
	{"pango_glyph_item_iter_next_cluster", &glyphItemIterNextCluster},
	{"pango_glyph_item_iter_prev_cluster", &glyphItemIterPrevCluster},
	{"pango_glyph_item_letter_space", &glyphItemLetterSpace},
	{"pango_glyph_item_split", &glyphItemSplit},
	{"pango_glyph_string_copy", &glyphStringCopy},
	{"pango_glyph_string_extents", &glyphStringExtents},
	{"pango_glyph_string_extents_range", &glyphStringExtentsRange},
	{"pango_glyph_string_free", &glyphStringFree},
	{"pango_glyph_string_get_logical_widths", &glyphStringGetLogicalWidths},
	{"pango_glyph_string_get_type", &GlyphStringGetType},
	{"pango_glyph_string_get_width", &glyphStringGetWidth},
	{"pango_glyph_string_index_to_x", &glyphStringIndexToX},
	{"pango_glyph_string_new", &GlyphStringNew},
	{"pango_glyph_string_set_size", &glyphStringSetSize},
	{"pango_glyph_string_x_to_index", &glyphStringXToIndex},
	{"pango_gravity_get_for_matrix", &GravityGetForMatrix},
	{"pango_gravity_get_for_script", &GravityGetForScript},
	{"pango_gravity_get_for_script_and_width", &GravityGetForScriptAndWidth},
	{"pango_gravity_get_type", &GravityGetType},
	{"pango_gravity_hint_get_type", &GravityHintGetType},
	{"pango_gravity_to_rotation", &gravityToRotation},
	{"pango_is_zero_width", &IsZeroWidth},
	{"pango_item_copy", &itemCopy},
	{"pango_item_free", &itemFree},
	{"pango_item_get_type", &ItemGetType},
	{"pango_item_new", &ItemNew},
	{"pango_item_split", &itemSplit},
	{"pango_itemize", &Itemize},
	{"pango_itemize_with_base_dir", &ItemizeWithBaseDir},
	{"pango_language_from_string", &LanguageFromString},
	{"pango_language_get_default", &LanguageGetDefault},
	{"pango_language_get_sample_string", &languageGetSampleString},
	{"pango_language_get_scripts", &languageGetScripts},
	{"pango_language_get_type", &LanguageGetType},
	{"pango_language_includes_script", &languageIncludesScript},
	{"pango_language_matches", &languageMatches},
	{"pango_language_to_string", &languageToString},
	{"pango_layout_context_changed", &layoutContextChanged},
	{"pango_layout_copy", &layoutCopy},
	{"pango_layout_get_alignment", &layoutGetAlignment},
	{"pango_layout_get_attributes", &layoutGetAttributes},
	{"pango_layout_get_auto_dir", &layoutGetAutoDir},
	{"pango_layout_get_baseline", &layoutGetBaseline},
	{"pango_layout_get_character_count", &layoutGetCharacterCount},
	{"pango_layout_get_context", &layoutGetContext},
	{"pango_layout_get_cursor_pos", &layoutGetCursorPos},
	{"pango_layout_get_ellipsize", &layoutGetEllipsize},
	{"pango_layout_get_extents", &layoutGetExtents},
	{"pango_layout_get_font_description", &layoutGetFontDescription},
	{"pango_layout_get_height", &layoutGetHeight},
	{"pango_layout_get_indent", &layoutGetIndent},
	{"pango_layout_get_iter", &layoutGetIter},
	{"pango_layout_get_justify", &layoutGetJustify},
	{"pango_layout_get_line", &layoutGetLine},
	{"pango_layout_get_line_count", &layoutGetLineCount},
	{"pango_layout_get_line_readonly", &layoutGetLineReadonly},
	{"pango_layout_get_lines", &layoutGetLines},
	{"pango_layout_get_lines_readonly", &layoutGetLinesReadonly},
	{"pango_layout_get_log_attrs", &layoutGetLogAttrs},
	{"pango_layout_get_log_attrs_readonly", &layoutGetLogAttrsReadonly},
	{"pango_layout_get_pixel_extents", &layoutGetPixelExtents},
	{"pango_layout_get_pixel_size", &layoutGetPixelSize},
	{"pango_layout_get_single_paragraph_mode", &layoutGetSingleParagraphMode},
	{"pango_layout_get_size", &layoutGetSize},
	{"pango_layout_get_spacing", &layoutGetSpacing},
	{"pango_layout_get_tabs", &layoutGetTabs},
	{"pango_layout_get_text", &layoutGetText},
	{"pango_layout_get_type", &layoutGetType},
	{"pango_layout_get_unknown_glyphs_count", &layoutGetUnknownGlyphsCount},
	{"pango_layout_get_width", &layoutGetWidth},
	{"pango_layout_get_wrap", &layoutGetWrap},
	{"pango_layout_index_to_line_x", &layoutIndexToLineX},
	{"pango_layout_index_to_pos", &layoutIndexToPos},
	{"pango_layout_is_ellipsized", &layoutIsEllipsized},
	{"pango_layout_is_wrapped", &layoutIsWrapped},
	{"pango_layout_iter_at_last_line", &layoutIterAtLastLine},
	{"pango_layout_iter_copy", &layoutIterCopy},
	{"pango_layout_iter_free", &layoutIterFree},
	{"pango_layout_iter_get_baseline", &layoutIterGetBaseline},
	{"pango_layout_iter_get_char_extents", &layoutIterGetCharExtents},
	{"pango_layout_iter_get_cluster_extents", &layoutIterGetClusterExtents},
	{"pango_layout_iter_get_index", &layoutIterGetIndex},
	{"pango_layout_iter_get_layout", &layoutIterGetLayout},
	{"pango_layout_iter_get_layout_extents", &layoutIterGetLayoutExtents},
	{"pango_layout_iter_get_line", &layoutIterGetLine},
	{"pango_layout_iter_get_line_extents", &layoutIterGetLineExtents},
	{"pango_layout_iter_get_line_readonly", &layoutIterGetLineReadonly},
	{"pango_layout_iter_get_line_yrange", &layoutIterGetLineYrange},
	{"pango_layout_iter_get_run", &layoutIterGetRun},
	{"pango_layout_iter_get_run_extents", &layoutIterGetRunExtents},
	{"pango_layout_iter_get_run_readonly", &layoutIterGetRunReadonly},
	{"pango_layout_iter_get_type", &LayoutIterGetType},
	{"pango_layout_iter_next_char", &layoutIterNextChar},
	{"pango_layout_iter_next_cluster", &layoutIterNextCluster},
	{"pango_layout_iter_next_line", &layoutIterNextLine},
	{"pango_layout_iter_next_run", &layoutIterNextRun},
	{"pango_layout_line_get_extents", &layoutLineGetExtents},
	{"pango_layout_line_get_pixel_extents", &layoutLineGetPixelExtents},
	{"pango_layout_line_get_type", &LayoutLineGetType},
	{"pango_layout_line_get_x_ranges", &layoutLineGetXRanges},
	{"pango_layout_line_index_to_x", &layoutLineIndexToX},
	{"pango_layout_line_ref", &layoutLineRef},
	{"pango_layout_line_unref", &layoutLineUnref},
	{"pango_layout_line_x_to_index", &layoutLineXToIndex},
	{"pango_layout_move_cursor_visually", &layoutMoveCursorVisually},
	{"pango_layout_new", &layoutNew},
	{"pango_layout_set_alignment", &layoutSetAlignment},
	{"pango_layout_set_attributes", &layoutSetAttributes},
	{"pango_layout_set_auto_dir", &layoutSetAutoDir},
	{"pango_layout_set_ellipsize", &layoutSetEllipsize},
	{"pango_layout_set_font_description", &layoutSetFontDescription},
	{"pango_layout_set_height", &layoutSetHeight},
	{"pango_layout_set_indent", &layoutSetIndent},
	{"pango_layout_set_justify", &layoutSetJustify},
	{"pango_layout_set_markup", &layoutSetMarkup},
	{"pango_layout_set_markup_with_accel", &layoutSetMarkupWithAccel},
	{"pango_layout_set_single_paragraph_mode", &layoutSetSingleParagraphMode},
	{"pango_layout_set_spacing", &layoutSetSpacing},
	{"pango_layout_set_tabs", &layoutSetTabs},
	{"pango_layout_set_text", &layoutSetText},
	{"pango_layout_set_width", &layoutSetWidth},
	{"pango_layout_set_wrap", &layoutSetWrap},
	{"pango_layout_xy_to_index", &layoutXyToIndex},
	{"pango_log2vis_get_embedding_levels", &Log2visGetEmbeddingLevels},
	{"pango_lookup_aliases", &LookupAliases},
	{"pango_map_get_engine", &mapGetEngine},
	{"pango_map_get_engines", &mapGetEngines},
	{"pango_matrix_concat", &matrixConcat},
	{"pango_matrix_copy", &matrixCopy},
	{"pango_matrix_free", &matrixFree},
	{"pango_matrix_get_font_scale_factor", &matrixGetFontScaleFactor},
	{"pango_matrix_get_type", &MatrixGetType},
	{"pango_matrix_rotate", &matrixRotate},
	{"pango_matrix_scale", &matrixScale},
	{"pango_matrix_transform_distance", &matrixTransformDistance},
	{"pango_matrix_transform_pixel_rectangle", &matrixTransformPixelRectangle},
	{"pango_matrix_transform_point", &matrixTransformPoint},
	{"pango_matrix_transform_rectangle", &matrixTransformRectangle},
	{"pango_matrix_translate", &matrixTranslate},
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
	{"pango_renderer_activate", &rendererActivate},
	{"pango_renderer_deactivate", &rendererDeactivate},
	{"pango_renderer_draw_error_underline", &rendererDrawErrorUnderline},
	{"pango_renderer_draw_glyph", &rendererDrawGlyph},
	{"pango_renderer_draw_glyph_item", &rendererDrawGlyphItem},
	{"pango_renderer_draw_glyphs", &rendererDrawGlyphs},
	{"pango_renderer_draw_layout", &rendererDrawLayout},
	{"pango_renderer_draw_layout_line", &rendererDrawLayoutLine},
	{"pango_renderer_draw_rectangle", &rendererDrawRectangle},
	{"pango_renderer_draw_trapezoid", &rendererDrawTrapezoid},
	{"pango_renderer_get_color", &rendererGetColor},
	{"pango_renderer_get_layout", &rendererGetLayout},
	{"pango_renderer_get_layout_line", &rendererGetLayoutLine},
	{"pango_renderer_get_matrix", &rendererGetMatrix},
	{"pango_renderer_get_type", &RendererGetType},
	{"pango_renderer_part_changed", &rendererPartChanged},
	{"pango_renderer_set_color", &rendererSetColor},
	{"pango_renderer_set_matrix", &rendererSetMatrix},
	{"pango_reorder_items", &ReorderItems},
	{"pango_scan_int", &ScanInt},
	{"pango_scan_string", &ScanString},
	{"pango_scan_word", &ScanWord},
	{"pango_script_for_unichar", &ScriptForUnichar},
	{"pango_script_get_sample_language", &scriptGetSampleLanguage},
	{"pango_script_get_type", &ScriptGetType},
	{"pango_script_iter_free", &scriptIterFree},
	{"pango_script_iter_get_range", &scriptIterGetRange},
	{"pango_script_iter_new", &ScriptIterNew},
	{"pango_script_iter_next", &scriptIterNext},
	{"pango_shape", &Shape},
	{"pango_skip_space", &SkipSpace},
	{"pango_split_file_list", &SplitFileList},
	{"pango_stretch_get_type", &StretchGetType},
	{"pango_style_get_type", &StyleGetType},
	{"pango_tab_align_get_type", &TabAlignGetType},
	{"pango_tab_array_copy", &tabArrayCopy},
	{"pango_tab_array_free", &tabArrayFree},
	{"pango_tab_array_get_positions_in_pixels", &tabArrayGetPositionsInPixels},
	{"pango_tab_array_get_size", &tabArrayGetSize},
	{"pango_tab_array_get_tab", &tabArrayGetTab},
	{"pango_tab_array_get_tabs", &tabArrayGetTabs},
	{"pango_tab_array_get_type", &TabArrayGetType},
	{"pango_tab_array_new", &TabArrayNew},
	{"pango_tab_array_new_with_positions", &TabArrayNewWithPositions},
	{"pango_tab_array_resize", &tabArrayResize},
	{"pango_tab_array_set_tab", &tabArraySetTab},
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
