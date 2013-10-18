package pango

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dllFt, false, apiListFt)
}

var (
	Pango_fc_decoder_get_type func() T.GType

	Pango_fc_decoder_get_charset func(
		decoder *T.PangoFcDecoder,
		fcfont *T.PangoFcFont) *T.FcCharSet

	Pango_fc_decoder_get_glyph func(
		decoder *T.PangoFcDecoder,
		fcfont *T.PangoFcFont,
		wc T.GUint32) T.PangoGlyph

	Pango_fc_font_has_char func(
		font *T.PangoFcFont,
		wc T.Gunichar) T.Gboolean

	Pango_fc_font_get_glyph func(
		font *T.PangoFcFont,
		wc T.Gunichar) uint

	Pango_fc_font_get_unknown_glyph func(
		font *T.PangoFcFont,
		wc T.Gunichar) T.PangoGlyph

	Pango_fc_font_kern_glyphs func(
		font *T.PangoFcFont,
		glyphs *T.PangoGlyphString)

	Pango_fc_font_get_type func() T.GType

	Pango_fc_font_lock_face func(
		font *T.PangoFcFont) T.FT_Face

	Pango_fc_font_unlock_face func(
		font *T.PangoFcFont)

	Pango_fc_fontset_key_get_language func(
		key *T.PangoFcFontsetKey) *T.PangoLanguage

	Pango_fc_fontset_key_get_description func(
		key *T.PangoFcFontsetKey) *T.PangoFontDescription

	Pango_fc_fontset_key_get_matrix func(
		key *T.PangoFcFontsetKey) *T.PangoMatrix

	Pango_fc_fontset_key_get_absolute_size func(
		key *T.PangoFcFontsetKey) float64

	Pango_fc_fontset_key_get_resolution func(
		key *T.PangoFcFontsetKey) float64

	Pango_fc_fontset_key_get_context_key func(
		key *T.PangoFcFontsetKey) T.Gpointer

	Pango_fc_font_key_get_pattern func(
		key *T.PangoFcFontKey) *T.FcPattern

	Pango_fc_font_key_get_matrix func(
		key *T.PangoFcFontKey) *T.PangoMatrix

	Pango_fc_font_key_get_context_key func(
		key *T.PangoFcFontKey) T.Gpointer

	Pango_fc_font_map_create_context func(
		fcfontmap *T.PangoFcFontMap) *T.PangoContext

	Pango_fc_font_map_shutdown func(
		fcfontmap *T.PangoFcFontMap)

	Pango_fc_font_map_get_type func() T.GType

	Pango_fc_font_map_cache_clear func(
		fcfontmap *T.PangoFcFontMap)

	Pango_fc_font_map_add_decoder_find_func func(
		fcfontmap *T.PangoFcFontMap,
		findfunc T.PangoFcDecoderFindFunc,
		user_data T.Gpointer,
		dnotify T.GDestroyNotify)

	Pango_fc_font_map_find_decoder func(
		fcfontmap *T.PangoFcFontMap,
		pattern *T.FcPattern) *T.PangoFcDecoder

	Pango_fc_font_description_from_pattern func(
		pattern *T.FcPattern,
		include_size T.Gboolean) *T.PangoFontDescription

	Pango_ft2_render func(
		bitmap *T.FT_Bitmap,
		font *T.PangoFont,
		glyphs *T.PangoGlyphString,
		x, y int)

	Pango_ft2_render_transformed func(
		bitmap *T.FT_Bitmap,
		matrix *T.PangoMatrix,
		font *T.PangoFont,
		glyphs *T.PangoGlyphString,
		x, y int)

	Pango_ft2_render_layout_line func(
		bitmap *T.FT_Bitmap,
		line *T.PangoLayoutLine,
		x int,
		y int)

	Pango_ft2_render_layout_line_subpixel func(
		bitmap *T.FT_Bitmap,
		line *T.PangoLayoutLine,
		x int,
		y int)

	Pango_ft2_render_layout func(
		bitmap *T.FT_Bitmap,
		layout *T.PangoLayout,
		x int,
		y int)

	Pango_ft2_render_layout_subpixel func(
		bitmap *T.FT_Bitmap,
		layout *T.PangoLayout,
		x int,
		y int)

	Pango_ft2_font_map_get_type func() T.GType

	Pango_ft2_font_map_new func() *T.PangoFontMap

	Pango_ft2_font_map_set_resolution func(
		fontmap *T.PangoFT2FontMap,
		dpi_x, dpi_y float64)

	Pango_ft2_font_map_set_default_substitute func(
		fontmap *T.PangoFT2FontMap,
		f T.PangoFT2SubstituteFunc,
		data T.Gpointer,
		notify T.GDestroyNotify)

	Pango_ft2_font_map_substitute_changed func(
		fontmap *T.PangoFT2FontMap)

	Pango_ft2_font_map_create_context func(
		fontmap *T.PangoFT2FontMap) *T.PangoContext

	Pango_ft2_get_context func(
		dpi_x, dpi_y float64) *T.PangoContext

	Pango_ft2_font_map_for_display func() *T.PangoFontMap

	Pango_ft2_shutdown_display func()

	Pango_ft2_get_unknown_glyph func(
		font *T.PangoFont) T.PangoGlyph

	Pango_ft2_font_get_kerning func(
		font *T.PangoFont,
		left T.PangoGlyph,
		right T.PangoGlyph) int

	Pango_ft2_font_get_face func(
		font *T.PangoFont) T.FT_Face

	Pango_ft2_font_get_coverage func(
		font *T.PangoFont,
		language *T.PangoLanguage) *T.PangoCoverage

	Pango_ot_info_get func(
		face T.FT_Face) *T.PangoOTInfo

	Pango_ot_info_find_script func(
		info *T.PangoOTInfo,
		table_type T.PangoOTTableType,
		script_tag T.PangoOTTag,
		script_index *uint) T.Gboolean

	Pango_ot_info_find_language func(
		info *T.PangoOTInfo,
		table_type T.PangoOTTableType,
		script_index uint,
		language_tag T.PangoOTTag,
		language_index *uint,
		required_feature_index *uint) T.Gboolean

	Pango_ot_info_find_feature func(
		info *T.PangoOTInfo,
		table_type T.PangoOTTableType,
		feature_tag T.PangoOTTag,
		script_index uint,
		language_index uint,
		feature_index *uint) T.Gboolean

	Pango_ot_info_list_scripts func(
		info *T.PangoOTInfo,
		table_type T.PangoOTTableType) *T.PangoOTTag

	Pango_ot_info_list_languages func(
		info *T.PangoOTInfo,
		table_type T.PangoOTTableType,
		script_index uint,
		language_tag T.PangoOTTag) *T.PangoOTTag

	Pango_ot_info_list_features func(
		info *T.PangoOTInfo,
		table_type T.PangoOTTableType,
		tag T.PangoOTTag,
		script_index uint,
		language_index uint) *T.PangoOTTag

	Pango_ot_buffer_new func(
		font *T.PangoFcFont) *T.PangoOTBuffer

	Pango_ot_buffer_destroy func(
		buffer *T.PangoOTBuffer)

	Pango_ot_buffer_clear func(
		buffer *T.PangoOTBuffer)

	Pango_ot_buffer_set_rtl func(
		buffer *T.PangoOTBuffer,
		rtl T.Gboolean)

	Pango_ot_buffer_add_glyph func(
		buffer *T.PangoOTBuffer,
		glyph uint,
		properties uint,
		cluster uint)

	Pango_ot_buffer_get_glyphs func(
		buffer *T.PangoOTBuffer,
		glyphs **T.PangoOTGlyph,
		n_glyphs *int)

	Pango_ot_buffer_output func(
		buffer *T.PangoOTBuffer,
		glyphs *T.PangoGlyphString)

	Pango_ot_buffer_set_zero_width_marks func(
		buffer *T.PangoOTBuffer,
		zero_width_marks T.Gboolean)

	Pango_ot_ruleset_get_for_description func(
		info *T.PangoOTInfo,
		desc *T.PangoOTRulesetDescription) *T.PangoOTRuleset

	Pango_ot_ruleset_new func(
		info *T.PangoOTInfo) *T.PangoOTRuleset

	Pango_ot_ruleset_new_for func(
		info *T.PangoOTInfo,
		script T.PangoScript,
		language *T.PangoLanguage) *T.PangoOTRuleset

	Pango_ot_ruleset_new_from_description func(
		info *T.PangoOTInfo,
		desc *T.PangoOTRulesetDescription) *T.PangoOTRuleset

	Pango_ot_ruleset_add_feature func(
		ruleset *T.PangoOTRuleset,
		table_type T.PangoOTTableType,
		feature_index uint,
		property_bit T.Gulong)

	Pango_ot_ruleset_maybe_add_feature func(
		ruleset *T.PangoOTRuleset,
		table_type T.PangoOTTableType,
		feature_tag T.PangoOTTag,
		property_bit T.Gulong) T.Gboolean

	Pango_ot_ruleset_maybe_add_features func(
		ruleset *T.PangoOTRuleset,
		table_type T.PangoOTTableType,
		features *T.PangoOTFeatureMap,
		n_features uint) uint

	Pango_ot_ruleset_get_feature_count func(
		ruleset *T.PangoOTRuleset,
		n_gsub_features *uint,
		n_gpos_features *uint) uint

	Pango_ot_ruleset_substitute func(
		ruleset *T.PangoOTRuleset,
		buffer *T.PangoOTBuffer)

	Pango_ot_ruleset_position func(
		ruleset *T.PangoOTRuleset,
		buffer *T.PangoOTBuffer)

	Pango_ot_tag_to_script func(
		script_tag T.PangoOTTag) T.PangoScript

	Pango_ot_tag_from_script func(
		script T.PangoScript) T.PangoOTTag

	Pango_ot_tag_to_language func(
		language_tag T.PangoOTTag) *T.PangoLanguage

	Pango_ot_tag_from_language func(
		language *T.PangoLanguage) T.PangoOTTag

	Pango_ot_ruleset_description_hash func(
		desc *T.PangoOTRulesetDescription) uint

	Pango_ot_ruleset_description_equal func(
		desc1 *T.PangoOTRulesetDescription,
		desc2 *T.PangoOTRulesetDescription) T.Gboolean

	Pango_ot_ruleset_description_copy func(
		desc *T.PangoOTRulesetDescription) *T.PangoOTRulesetDescription

	Pango_ot_ruleset_description_free func(
		desc *T.PangoOTRulesetDescription)

	Pango_ot_info_get_type func() T.GType

	Pango_ot_ruleset_get_type func() T.GType
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
