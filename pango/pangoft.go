package pango

import (
	"github.com/tHinqa/outside"
	FT "github.com/tHinqa/outside-gtk2/freetype2"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dllFt, false, apiListFt)
}

var (
	OtInfoGet func(
		face FT.Face) *OTInfo

	OtInfoFindScript func(
		info *OTInfo,
		tableType OTTableType,
		scriptTag OTTag,
		scriptIndex *uint) bool

	OtInfoFindLanguage func(
		info *OTInfo,
		tableType OTTableType,
		scriptIndex uint,
		languageTag OTTag,
		languageIndex *uint,
		requiredFeatureIndex *uint) bool

	OtInfoFindFeature func(
		info *OTInfo,
		tableType OTTableType,
		featureTag OTTag,
		scriptIndex uint,
		languageIndex uint,
		featureIndex *uint) bool

	OtInfoListScripts func(
		info *OTInfo,
		tableType OTTableType) *OTTag

	OtInfoListLanguages func(
		info *OTInfo,
		tableType OTTableType,
		scriptIndex uint,
		languageTag OTTag) *OTTag

	OtInfoListFeatures func(
		info *OTInfo,
		tableType OTTableType,
		tag OTTag,
		scriptIndex uint,
		languageIndex uint) *OTTag

	OtBufferNew func(
		font *FcFont) *OTBuffer

	OtBufferDestroy func(
		buffer *OTBuffer)

	OtBufferClear func(
		buffer *OTBuffer)

	OtBufferSetRtl func(
		buffer *OTBuffer,
		rtl bool)

	OtBufferAddGlyph func(
		buffer *OTBuffer,
		glyph uint,
		properties uint,
		cluster uint)

	OtBufferGetGlyphs func(
		buffer *OTBuffer,
		glyphs **OTGlyph,
		nGlyphs *int)

	OtBufferOutput func(
		buffer *OTBuffer,
		glyphs *GlyphString)

	OtBufferSetZeroWidthMarks func(
		buffer *OTBuffer,
		zeroWidthMarks bool)

	OtRulesetGetForDescription func(
		info *OTInfo,
		desc *OTRulesetDescription) *OTRuleset

	OtRulesetNew func(
		info *OTInfo) *OTRuleset

	OtRulesetNewFor func(
		info *OTInfo,
		script Script,
		language *Language) *OTRuleset

	OtRulesetNewFromDescription func(
		info *OTInfo,
		desc *OTRulesetDescription) *OTRuleset

	OtRulesetAddFeature func(
		ruleset *OTRuleset,
		tableType OTTableType,
		featureIndex uint,
		propertyBit T.Gulong)

	OtRulesetMaybeAddFeature func(
		ruleset *OTRuleset,
		tableType OTTableType,
		featureTag OTTag,
		propertyBit T.Gulong) bool

	OtRulesetMaybeAddFeatures func(
		ruleset *OTRuleset,
		tableType OTTableType,
		features *OTFeatureMap,
		nFeatures uint) uint

	OtRulesetGetFeatureCount func(
		ruleset *OTRuleset,
		nGsubFeatures *uint,
		nGposFeatures *uint) uint

	OtRulesetSubstitute func(
		ruleset *OTRuleset,
		buffer *OTBuffer)

	OtRulesetPosition func(
		ruleset *OTRuleset,
		buffer *OTBuffer)

	OtTagToScript func(
		scriptTag OTTag) Script

	OtTagFromScript func(
		script Script) OTTag

	OtTagToLanguage func(
		languageTag OTTag) *Language

	OtTagFromLanguage func(
		language *Language) OTTag

	OtRulesetDescriptionHash func(
		desc *OTRulesetDescription) uint

	OtRulesetDescriptionEqual func(
		desc1 *OTRulesetDescription,
		desc2 *OTRulesetDescription) bool

	OtRulesetDescriptionCopy func(
		desc *OTRulesetDescription) *OTRulesetDescription

	OtRulesetDescriptionFree func(
		desc *OTRulesetDescription)

	OtInfoGetType func() O.Type

	OtRulesetGetType func() O.Type
)

var dllFt = "libpangoft2-1.0-0.dll"

var apiListFt = outside.Apis{
	{"pango_fc_decoder_get_charset", &fcDecoderGetCharset},
	{"pango_fc_decoder_get_glyph", &fcDecoderGetGlyph},
	{"pango_fc_decoder_get_type", &FcDecoderGetType},
	// Undocumented {"pango_fc_font_create_base_metrics_for_context", &FcFontCreateBaseMetricsForContext},
	{"pango_fc_font_description_from_pattern", &FcFontDescriptionFromPattern},
	{"pango_fc_font_get_glyph", &fcFontGetGlyph},
	// Undocumented {"pango_fc_font_get_raw_extents", &FcFontGetRawExtents},
	{"pango_fc_font_get_type", &FcFontGetType},
	{"pango_fc_font_get_unknown_glyph", &fcFontGetUnknownGlyph},
	{"pango_fc_font_has_char", &fcFontHasChar},
	{"pango_fc_font_kern_glyphs", &fcFontKernGlyphs},
	{"pango_fc_font_key_get_context_key", &fcFontKeyGetContextKey},
	{"pango_fc_font_key_get_matrix", &fcFontKeyGetMatrix},
	{"pango_fc_font_key_get_pattern", &fcFontKeyGetPattern},
	{"pango_fc_font_lock_face", &fcFontLockFace},
	{"pango_fc_font_map_add_decoder_find_func", &fcFontMapAddDecoderFindFunc},
	{"pango_fc_font_map_cache_clear", &fcFontMapCacheClear},
	{"pango_fc_font_map_create_context", &fcFontMapCreateContext},
	{"pango_fc_font_map_find_decoder", &fcFontMapFindDecoder},
	{"pango_fc_font_map_get_type", &FcFontMapGetType},
	{"pango_fc_font_map_shutdown", &fcFontMapShutdown},
	{"pango_fc_font_unlock_face", &fcFontUnlockFace},
	{"pango_fc_fontset_key_get_absolute_size", &fcFontsetKeyGetAbsoluteSize},
	{"pango_fc_fontset_key_get_context_key", &fcFontsetKeyGetContextKey},
	{"pango_fc_fontset_key_get_description", &fcFontsetKeyGetDescription},
	{"pango_fc_fontset_key_get_language", &fcFontsetKeyGetLanguage},
	{"pango_fc_fontset_key_get_matrix", &fcFontsetKeyGetMatrix},
	{"pango_fc_fontset_key_get_resolution", &fcFontsetKeyGetResolution},
	{"pango_ft2_font_get_coverage", &Ft2FontGetCoverage},
	{"pango_ft2_font_get_face", &Ft2FontGetFace},
	{"pango_ft2_font_get_kerning", &Ft2FontGetKerning},
	// Undocumented {"pango_ft2_font_get_type", &Ft2FontGetType},
	{"pango_ft2_font_map_create_context", &ft2FontMapCreateContext},
	{"pango_ft2_font_map_for_display", &Ft2FontMapForDisplay},
	{"pango_ft2_font_map_get_type", &Ft2FontMapGetType},
	{"pango_ft2_font_map_new", &Ft2FontMapNew},
	{"pango_ft2_font_map_set_default_substitute", &ft2FontMapSetDefaultSubstitute},
	{"pango_ft2_font_map_set_resolution", &ft2FontMapSetResolution},
	{"pango_ft2_font_map_substitute_changed", &ft2FontMapSubstituteChanged},
	{"pango_ft2_get_context", &Ft2GetContext},
	{"pango_ft2_get_unknown_glyph", &Ft2GetUnknownGlyph},
	{"pango_ft2_render", &Ft2Render},
	{"pango_ft2_render_layout", &Ft2RenderLayout},
	{"pango_ft2_render_layout_line", &Ft2RenderLayoutLine},
	{"pango_ft2_render_layout_line_subpixel", &Ft2RenderLayoutLineSubpixel},
	{"pango_ft2_render_layout_subpixel", &Ft2RenderLayoutSubpixel},
	{"pango_ft2_render_transformed", &Ft2RenderTransformed},
	// Undocumented {"pango_ft2_renderer_get_type", &Ft2RendererGetType},
	{"pango_ft2_shutdown_display", &Ft2ShutdownDisplay},
	{"pango_ot_buffer_add_glyph", &OtBufferAddGlyph},
	{"pango_ot_buffer_clear", &OtBufferClear},
	{"pango_ot_buffer_destroy", &OtBufferDestroy},
	{"pango_ot_buffer_get_glyphs", &OtBufferGetGlyphs},
	{"pango_ot_buffer_new", &OtBufferNew},
	{"pango_ot_buffer_output", &OtBufferOutput},
	{"pango_ot_buffer_set_rtl", &OtBufferSetRtl},
	{"pango_ot_buffer_set_zero_width_marks", &OtBufferSetZeroWidthMarks},
	{"pango_ot_info_find_feature", &OtInfoFindFeature},
	{"pango_ot_info_find_language", &OtInfoFindLanguage},
	{"pango_ot_info_find_script", &OtInfoFindScript},
	{"pango_ot_info_get", &OtInfoGet},
	{"pango_ot_info_get_type", &OtInfoGetType},
	{"pango_ot_info_list_features", &OtInfoListFeatures},
	{"pango_ot_info_list_languages", &OtInfoListLanguages},
	{"pango_ot_info_list_scripts", &OtInfoListScripts},
	{"pango_ot_ruleset_add_feature", &OtRulesetAddFeature},
	{"pango_ot_ruleset_description_copy", &OtRulesetDescriptionCopy},
	{"pango_ot_ruleset_description_equal", &OtRulesetDescriptionEqual},
	{"pango_ot_ruleset_description_free", &OtRulesetDescriptionFree},
	{"pango_ot_ruleset_description_hash", &OtRulesetDescriptionHash},
	{"pango_ot_ruleset_get_feature_count", &OtRulesetGetFeatureCount},
	{"pango_ot_ruleset_get_for_description", &OtRulesetGetForDescription},
	{"pango_ot_ruleset_get_type", &OtRulesetGetType},
	{"pango_ot_ruleset_maybe_add_feature", &OtRulesetMaybeAddFeature},
	{"pango_ot_ruleset_maybe_add_features", &OtRulesetMaybeAddFeatures},
	{"pango_ot_ruleset_new", &OtRulesetNew},
	{"pango_ot_ruleset_new_for", &OtRulesetNewFor},
	{"pango_ot_ruleset_new_from_description", &OtRulesetNewFromDescription},
	{"pango_ot_ruleset_position", &OtRulesetPosition},
	{"pango_ot_ruleset_substitute", &OtRulesetSubstitute},
	{"pango_ot_tag_from_language", &OtTagFromLanguage},
	{"pango_ot_tag_from_script", &OtTagFromScript},
	{"pango_ot_tag_to_language", &OtTagToLanguage},
	{"pango_ot_tag_to_script", &OtTagToScript},
}
