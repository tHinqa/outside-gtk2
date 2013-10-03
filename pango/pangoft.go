package pango

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dllFt, false, apiListFt)
}

var (
	Pango_fc_decoder_get_type func() GType

	Pango_fc_decoder_get_charset func(
		decoder *PangoFcDecoder,
		fcfont *PangoFcFont) *FcCharSet

	Pango_fc_decoder_get_glyph func(
		decoder *PangoFcDecoder,
		fcfont *PangoFcFont,
		wc Guint32) PangoGlyph

	Pango_fc_font_has_char func(
		font *PangoFcFont,
		wc Gunichar) Gboolean

	Pango_fc_font_get_glyph func(
		font *PangoFcFont,
		wc Gunichar) Guint

	Pango_fc_font_get_unknown_glyph func(
		font *PangoFcFont,
		wc Gunichar) PangoGlyph

	Pango_fc_font_kern_glyphs func(
		font *PangoFcFont,
		glyphs *PangoGlyphString)

	Pango_fc_font_get_type func() GType

	Pango_fc_font_lock_face func(
		font *PangoFcFont) FT_Face

	Pango_fc_font_unlock_face func(
		font *PangoFcFont)

	Pango_fc_fontset_key_get_language func(
		key *PangoFcFontsetKey) *PangoLanguage

	Pango_fc_fontset_key_get_description func(
		key *PangoFcFontsetKey) *PangoFontDescription

	Pango_fc_fontset_key_get_matrix func(
		key *PangoFcFontsetKey) *PangoMatrix

	Pango_fc_fontset_key_get_absolute_size func(
		key *PangoFcFontsetKey) Double

	Pango_fc_fontset_key_get_resolution func(
		key *PangoFcFontsetKey) Double

	Pango_fc_fontset_key_get_context_key func(
		key *PangoFcFontsetKey) Gpointer

	Pango_fc_font_key_get_pattern func(
		key *PangoFcFontKey) *FcPattern

	Pango_fc_font_key_get_matrix func(
		key *PangoFcFontKey) *PangoMatrix

	Pango_fc_font_key_get_context_key func(
		key *PangoFcFontKey) Gpointer

	Pango_fc_font_map_create_context func(
		fcfontmap *PangoFcFontMap) *PangoContext

	Pango_fc_font_map_shutdown func(
		fcfontmap *PangoFcFontMap)

	Pango_fc_font_map_get_type func() GType

	Pango_fc_font_map_cache_clear func(
		fcfontmap *PangoFcFontMap)

	Pango_fc_font_map_add_decoder_find_func func(
		fcfontmap *PangoFcFontMap,
		findfunc PangoFcDecoderFindFunc,
		user_data Gpointer,
		dnotify GDestroyNotify)

	Pango_fc_font_map_find_decoder func(
		fcfontmap *PangoFcFontMap,
		pattern *FcPattern) *PangoFcDecoder

	Pango_fc_font_description_from_pattern func(
		pattern *FcPattern,
		include_size Gboolean) *PangoFontDescription

	Pango_ft2_render func(
		bitmap *FT_Bitmap,
		font *PangoFont,
		glyphs *PangoGlyphString,
		x, y Gint)

	Pango_ft2_render_transformed func(
		bitmap *FT_Bitmap,
		matrix *PangoMatrix,
		font *PangoFont,
		glyphs *PangoGlyphString,
		x, y int)

	Pango_ft2_render_layout_line func(
		bitmap *FT_Bitmap,
		line *PangoLayoutLine,
		x int,
		y int)

	Pango_ft2_render_layout_line_subpixel func(
		bitmap *FT_Bitmap,
		line *PangoLayoutLine,
		x int,
		y int)

	Pango_ft2_render_layout func(
		bitmap *FT_Bitmap,
		layout *PangoLayout,
		x int,
		y int)

	Pango_ft2_render_layout_subpixel func(
		bitmap *FT_Bitmap,
		layout *PangoLayout,
		x int,
		y int)

	Pango_ft2_font_map_get_type func() GType

	Pango_ft2_font_map_new func() *PangoFontMap

	Pango_ft2_font_map_set_resolution func(
		fontmap *PangoFT2FontMap,
		dpi_x, dpi_y Double)

	Pango_ft2_font_map_set_default_substitute func(
		fontmap *PangoFT2FontMap,
		f PangoFT2SubstituteFunc,
		data Gpointer,
		notify GDestroyNotify)

	Pango_ft2_font_map_substitute_changed func(
		fontmap *PangoFT2FontMap)

	Pango_ft2_font_map_create_context func(
		fontmap *PangoFT2FontMap) *PangoContext

	Pango_ft2_get_context func(
		dpi_x, dpi_y Double) *PangoContext

	Pango_ft2_font_map_for_display func() *PangoFontMap

	Pango_ft2_shutdown_display func()

	Pango_ft2_get_unknown_glyph func(
		font *PangoFont) PangoGlyph

	Pango_ft2_font_get_kerning func(
		font *PangoFont,
		left PangoGlyph,
		right PangoGlyph) int

	Pango_ft2_font_get_face func(
		font *PangoFont) FT_Face

	Pango_ft2_font_get_coverage func(
		font *PangoFont,
		language *PangoLanguage) *PangoCoverage

	Pango_ot_info_get func(
		face FT_Face) *PangoOTInfo

	Pango_ot_info_find_script func(
		info *PangoOTInfo,
		table_type PangoOTTableType,
		script_tag PangoOTTag,
		script_index *Guint) Gboolean

	Pango_ot_info_find_language func(
		info *PangoOTInfo,
		table_type PangoOTTableType,
		script_index Guint,
		language_tag PangoOTTag,
		language_index *Guint,
		required_feature_index *Guint) Gboolean

	Pango_ot_info_find_feature func(
		info *PangoOTInfo,
		table_type PangoOTTableType,
		feature_tag PangoOTTag,
		script_index Guint,
		language_index Guint,
		feature_index *Guint) Gboolean

	Pango_ot_info_list_scripts func(
		info *PangoOTInfo,
		table_type PangoOTTableType) *PangoOTTag

	Pango_ot_info_list_languages func(
		info *PangoOTInfo,
		table_type PangoOTTableType,
		script_index Guint,
		language_tag PangoOTTag) *PangoOTTag

	Pango_ot_info_list_features func(
		info *PangoOTInfo,
		table_type PangoOTTableType,
		tag PangoOTTag,
		script_index Guint,
		language_index Guint) *PangoOTTag

	Pango_ot_buffer_new func(
		font *PangoFcFont) *PangoOTBuffer

	Pango_ot_buffer_destroy func(
		buffer *PangoOTBuffer)

	Pango_ot_buffer_clear func(
		buffer *PangoOTBuffer)

	Pango_ot_buffer_set_rtl func(
		buffer *PangoOTBuffer,
		rtl Gboolean)

	Pango_ot_buffer_add_glyph func(
		buffer *PangoOTBuffer,
		glyph Guint,
		properties Guint,
		cluster Guint)

	Pango_ot_buffer_get_glyphs func(
		buffer *PangoOTBuffer,
		glyphs **PangoOTGlyph,
		n_glyphs *int)

	Pango_ot_buffer_output func(
		buffer *PangoOTBuffer,
		glyphs *PangoGlyphString)

	Pango_ot_buffer_set_zero_width_marks func(
		buffer *PangoOTBuffer,
		zero_width_marks Gboolean)

	Pango_ot_ruleset_get_for_description func(
		info *PangoOTInfo,
		desc *PangoOTRulesetDescription) *PangoOTRuleset

	Pango_ot_ruleset_new func(
		info *PangoOTInfo) *PangoOTRuleset

	Pango_ot_ruleset_new_for func(
		info *PangoOTInfo,
		script PangoScript,
		language *PangoLanguage) *PangoOTRuleset

	Pango_ot_ruleset_new_from_description func(
		info *PangoOTInfo,
		desc *PangoOTRulesetDescription) *PangoOTRuleset

	Pango_ot_ruleset_add_feature func(
		ruleset *PangoOTRuleset,
		table_type PangoOTTableType,
		feature_index Guint,
		property_bit Gulong)

	Pango_ot_ruleset_maybe_add_feature func(
		ruleset *PangoOTRuleset,
		table_type PangoOTTableType,
		feature_tag PangoOTTag,
		property_bit Gulong) Gboolean

	Pango_ot_ruleset_maybe_add_features func(
		ruleset *PangoOTRuleset,
		table_type PangoOTTableType,
		features *PangoOTFeatureMap,
		n_features Guint) Guint

	Pango_ot_ruleset_get_feature_count func(
		ruleset *PangoOTRuleset,
		n_gsub_features *Guint,
		n_gpos_features *Guint) Guint

	Pango_ot_ruleset_substitute func(
		ruleset *PangoOTRuleset,
		buffer *PangoOTBuffer)

	Pango_ot_ruleset_position func(
		ruleset *PangoOTRuleset,
		buffer *PangoOTBuffer)

	Pango_ot_tag_to_script func(
		script_tag PangoOTTag) PangoScript

	Pango_ot_tag_from_script func(
		script PangoScript) PangoOTTag

	Pango_ot_tag_to_language func(
		language_tag PangoOTTag) *PangoLanguage

	Pango_ot_tag_from_language func(
		language *PangoLanguage) PangoOTTag

	Pango_ot_ruleset_description_hash func(
		desc *PangoOTRulesetDescription) Guint

	Pango_ot_ruleset_description_equal func(
		desc1 *PangoOTRulesetDescription,
		desc2 *PangoOTRulesetDescription) Gboolean

	Pango_ot_ruleset_description_copy func(
		desc *PangoOTRulesetDescription) *PangoOTRulesetDescription

	Pango_ot_ruleset_description_free func(
		desc *PangoOTRulesetDescription)

	Pango_ot_info_get_type func() GType

	Pango_ot_ruleset_get_type func() GType
)

var dllFt = "libpangoft2-1.0-0.dll"

var apiListFt = outside.Apis{
	{"pango_fc_decoder_get_charset", &Pango_fc_decoder_get_charset},
	{"pango_fc_decoder_get_glyph", &Pango_fc_decoder_get_glyph},
	{"pango_fc_decoder_get_type", &Pango_fc_decoder_get_type},
	// Undocumented {"pango_fc_font_create_base_metrics_for_context", &Pango_fc_font_create_base_metrics_for_context},
	{"pango_fc_font_description_from_pattern", &Pango_fc_font_description_from_pattern},
	{"pango_fc_font_get_glyph", &Pango_fc_font_get_glyph},
	// Undocumented {"pango_fc_font_get_raw_extents", &Pango_fc_font_get_raw_extents},
	{"pango_fc_font_get_type", &Pango_fc_font_get_type},
	{"pango_fc_font_get_unknown_glyph", &Pango_fc_font_get_unknown_glyph},
	{"pango_fc_font_has_char", &Pango_fc_font_has_char},
	{"pango_fc_font_kern_glyphs", &Pango_fc_font_kern_glyphs},
	{"pango_fc_font_key_get_context_key", &Pango_fc_font_key_get_context_key},
	{"pango_fc_font_key_get_matrix", &Pango_fc_font_key_get_matrix},
	{"pango_fc_font_key_get_pattern", &Pango_fc_font_key_get_pattern},
	{"pango_fc_font_lock_face", &Pango_fc_font_lock_face},
	{"pango_fc_font_map_add_decoder_find_func", &Pango_fc_font_map_add_decoder_find_func},
	{"pango_fc_font_map_cache_clear", &Pango_fc_font_map_cache_clear},
	{"pango_fc_font_map_create_context", &Pango_fc_font_map_create_context},
	{"pango_fc_font_map_find_decoder", &Pango_fc_font_map_find_decoder},
	{"pango_fc_font_map_get_type", &Pango_fc_font_map_get_type},
	{"pango_fc_font_map_shutdown", &Pango_fc_font_map_shutdown},
	{"pango_fc_font_unlock_face", &Pango_fc_font_unlock_face},
	{"pango_fc_fontset_key_get_absolute_size", &Pango_fc_fontset_key_get_absolute_size},
	{"pango_fc_fontset_key_get_context_key", &Pango_fc_fontset_key_get_context_key},
	{"pango_fc_fontset_key_get_description", &Pango_fc_fontset_key_get_description},
	{"pango_fc_fontset_key_get_language", &Pango_fc_fontset_key_get_language},
	{"pango_fc_fontset_key_get_matrix", &Pango_fc_fontset_key_get_matrix},
	{"pango_fc_fontset_key_get_resolution", &Pango_fc_fontset_key_get_resolution},
	{"pango_ft2_font_get_coverage", &Pango_ft2_font_get_coverage},
	{"pango_ft2_font_get_face", &Pango_ft2_font_get_face},
	{"pango_ft2_font_get_kerning", &Pango_ft2_font_get_kerning},
	// Undocumented {"pango_ft2_font_get_type", &Pango_ft2_font_get_type},
	{"pango_ft2_font_map_create_context", &Pango_ft2_font_map_create_context},
	{"pango_ft2_font_map_for_display", &Pango_ft2_font_map_for_display},
	{"pango_ft2_font_map_get_type", &Pango_ft2_font_map_get_type},
	{"pango_ft2_font_map_new", &Pango_ft2_font_map_new},
	{"pango_ft2_font_map_set_default_substitute", &Pango_ft2_font_map_set_default_substitute},
	{"pango_ft2_font_map_set_resolution", &Pango_ft2_font_map_set_resolution},
	{"pango_ft2_font_map_substitute_changed", &Pango_ft2_font_map_substitute_changed},
	{"pango_ft2_get_context", &Pango_ft2_get_context},
	{"pango_ft2_get_unknown_glyph", &Pango_ft2_get_unknown_glyph},
	{"pango_ft2_render", &Pango_ft2_render},
	{"pango_ft2_render_layout", &Pango_ft2_render_layout},
	{"pango_ft2_render_layout_line", &Pango_ft2_render_layout_line},
	{"pango_ft2_render_layout_line_subpixel", &Pango_ft2_render_layout_line_subpixel},
	{"pango_ft2_render_layout_subpixel", &Pango_ft2_render_layout_subpixel},
	{"pango_ft2_render_transformed", &Pango_ft2_render_transformed},
	// Undocumented {"pango_ft2_renderer_get_type", &Pango_ft2_renderer_get_type},
	{"pango_ft2_shutdown_display", &Pango_ft2_shutdown_display},
	{"pango_ot_buffer_add_glyph", &Pango_ot_buffer_add_glyph},
	{"pango_ot_buffer_clear", &Pango_ot_buffer_clear},
	{"pango_ot_buffer_destroy", &Pango_ot_buffer_destroy},
	{"pango_ot_buffer_get_glyphs", &Pango_ot_buffer_get_glyphs},
	{"pango_ot_buffer_new", &Pango_ot_buffer_new},
	{"pango_ot_buffer_output", &Pango_ot_buffer_output},
	{"pango_ot_buffer_set_rtl", &Pango_ot_buffer_set_rtl},
	{"pango_ot_buffer_set_zero_width_marks", &Pango_ot_buffer_set_zero_width_marks},
	{"pango_ot_info_find_feature", &Pango_ot_info_find_feature},
	{"pango_ot_info_find_language", &Pango_ot_info_find_language},
	{"pango_ot_info_find_script", &Pango_ot_info_find_script},
	{"pango_ot_info_get", &Pango_ot_info_get},
	{"pango_ot_info_get_type", &Pango_ot_info_get_type},
	{"pango_ot_info_list_features", &Pango_ot_info_list_features},
	{"pango_ot_info_list_languages", &Pango_ot_info_list_languages},
	{"pango_ot_info_list_scripts", &Pango_ot_info_list_scripts},
	{"pango_ot_ruleset_add_feature", &Pango_ot_ruleset_add_feature},
	{"pango_ot_ruleset_description_copy", &Pango_ot_ruleset_description_copy},
	{"pango_ot_ruleset_description_equal", &Pango_ot_ruleset_description_equal},
	{"pango_ot_ruleset_description_free", &Pango_ot_ruleset_description_free},
	{"pango_ot_ruleset_description_hash", &Pango_ot_ruleset_description_hash},
	{"pango_ot_ruleset_get_feature_count", &Pango_ot_ruleset_get_feature_count},
	{"pango_ot_ruleset_get_for_description", &Pango_ot_ruleset_get_for_description},
	{"pango_ot_ruleset_get_type", &Pango_ot_ruleset_get_type},
	{"pango_ot_ruleset_maybe_add_feature", &Pango_ot_ruleset_maybe_add_feature},
	{"pango_ot_ruleset_maybe_add_features", &Pango_ot_ruleset_maybe_add_features},
	{"pango_ot_ruleset_new", &Pango_ot_ruleset_new},
	{"pango_ot_ruleset_new_for", &Pango_ot_ruleset_new_for},
	{"pango_ot_ruleset_new_from_description", &Pango_ot_ruleset_new_from_description},
	{"pango_ot_ruleset_position", &Pango_ot_ruleset_position},
	{"pango_ot_ruleset_substitute", &Pango_ot_ruleset_substitute},
	{"pango_ot_tag_from_language", &Pango_ot_tag_from_language},
	{"pango_ot_tag_from_script", &Pango_ot_tag_from_script},
	{"pango_ot_tag_to_language", &Pango_ot_tag_to_language},
	{"pango_ot_tag_to_script", &Pango_ot_tag_to_script},
}
