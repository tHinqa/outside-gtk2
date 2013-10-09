// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package pango provides API definitions for accessing
//libpango-1.0-0.dll and libpangocairo-1.0-0.dll.
package pango

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
	AddDllApis(dllCairo, false, apiListCairo)
}

var (
	Pango_coverage_new func() *PangoCoverage

	Pango_coverage_ref func(
		coverage *PangoCoverage) *PangoCoverage

	Pango_coverage_unref func(
		coverage *PangoCoverage)

	Pango_coverage_copy func(
		coverage *PangoCoverage) *PangoCoverage

	Pango_coverage_get func(
		coverage *PangoCoverage,
		index int) PangoCoverageLevel

	Pango_coverage_set func(
		coverage *PangoCoverage,
		index int,
		level PangoCoverageLevel)

	Pango_coverage_max func(
		coverage *PangoCoverage,
		other *PangoCoverage)

	Pango_coverage_to_bytes func(
		coverage *PangoCoverage,
		bytes **Guchar,
		n_bytes *int)

	Pango_coverage_from_bytes func(
		bytes *Guchar,
		n_bytes int) *PangoCoverage

	Pango_units_from_double func(
		d Double) int

	Pango_units_to_double func(
		i int) Double

	Pango_extents_to_pixels func(
		inclusive *PangoRectangle,
		nearest *PangoRectangle)

	Pango_matrix_get_type func() GType

	Pango_matrix_copy func(
		matrix *PangoMatrix) *PangoMatrix

	Pango_matrix_free func(
		matrix *PangoMatrix)

	Pango_matrix_translate func(
		matrix *PangoMatrix,
		tx Double,
		ty Double)

	Pango_matrix_scale func(
		matrix *PangoMatrix,
		scale_x Double,
		scale_y Double)

	Pango_matrix_rotate func(
		matrix *PangoMatrix,
		degrees Double)

	Pango_matrix_concat func(
		matrix *PangoMatrix,
		new_matrix *PangoMatrix)

	Pango_matrix_transform_point func(
		matrix *PangoMatrix,
		x *Double,
		y *Double)

	Pango_matrix_transform_distance func(
		matrix *PangoMatrix,
		dx *Double,
		dy *Double)

	Pango_matrix_transform_rectangle func(
		matrix *PangoMatrix,
		rect *PangoRectangle)

	Pango_matrix_transform_pixel_rectangle func(
		matrix *PangoMatrix,
		rect *PangoRectangle)

	Pango_matrix_get_font_scale_factor func(
		matrix *PangoMatrix) Double

	Pango_script_for_unichar func(
		ch Gunichar) PangoScript

	Pango_script_iter_new func(
		text string,
		length int) *PangoScriptIter

	Pango_script_iter_get_range func(
		iter *PangoScriptIter,
		start **Char,
		end **Char,
		script *PangoScript)

	Pango_script_iter_next func(
		iter *PangoScriptIter) Gboolean

	Pango_script_iter_free func(
		iter *PangoScriptIter)

	Pango_language_get_type func() GType

	Pango_language_from_string func(
		language string) *PangoLanguage

	Pango_language_to_string func(
		language *PangoLanguage) string

	Pango_language_get_sample_string func(
		language *PangoLanguage) string

	Pango_language_get_default func() *PangoLanguage

	Pango_language_matches func(
		language *PangoLanguage,
		range_list string) Gboolean

	Pango_language_includes_script func(
		language *PangoLanguage,
		script PangoScript) Gboolean

	Pango_language_get_scripts func(
		language *PangoLanguage,
		num_scripts *int) *PangoScript

	Pango_script_get_sample_language func(
		script PangoScript) *PangoLanguage

	Pango_gravity_to_rotation func(
		gravity PangoGravity) Double

	Pango_gravity_get_for_matrix func(
		matrix *PangoMatrix) PangoGravity

	Pango_gravity_get_for_script func(
		script PangoScript,
		base_gravity PangoGravity,
		hint PangoGravityHint) PangoGravity

	Pango_gravity_get_for_script_and_width func(
		script PangoScript,
		wide Gboolean,
		base_gravity PangoGravity,
		hint PangoGravityHint) PangoGravity

	Pango_bidi_type_for_unichar func(
		ch Gunichar) PangoBidiType

	Pango_unichar_direction func(
		ch Gunichar) PangoDirection

	Pango_find_base_dir func(
		text string,
		length Gint) PangoDirection

	Pango_get_mirror_char func(
		ch Gunichar,
		mirrored_ch *Gunichar) Gboolean

	Pango_font_description_get_type func() GType

	Pango_font_description_new func() *PangoFontDescription

	Pango_font_description_copy func(
		desc *PangoFontDescription) *PangoFontDescription

	Pango_font_description_copy_static func(
		desc *PangoFontDescription) *PangoFontDescription

	Pango_font_description_hash func(
		desc *PangoFontDescription) Guint

	Pango_font_description_equal func(
		desc1 *PangoFontDescription,
		desc2 *PangoFontDescription) Gboolean

	Pango_font_description_free func(
		desc *PangoFontDescription)

	Pango_font_descriptions_free func(
		descs **PangoFontDescription,
		n_descs int)

	Pango_font_description_set_family func(
		desc *PangoFontDescription,
		family string)

	Pango_font_description_set_family_static func(
		desc *PangoFontDescription,
		family string)

	Pango_font_description_get_family func(
		desc *PangoFontDescription) string

	Pango_font_description_set_style func(
		desc *PangoFontDescription,
		style PangoStyle)

	Pango_font_description_get_style func(
		desc *PangoFontDescription) PangoStyle

	Pango_font_description_set_variant func(
		desc *PangoFontDescription,
		variant PangoVariant)

	Pango_font_description_get_variant func(
		desc *PangoFontDescription) PangoVariant

	Pango_font_description_set_weight func(
		desc *PangoFontDescription,
		weight PangoWeight)

	Pango_font_description_get_weight func(
		desc *PangoFontDescription) PangoWeight

	Pango_font_description_set_stretch func(
		desc *PangoFontDescription,
		stretch PangoStretch)

	Pango_font_description_get_stretch func(
		desc *PangoFontDescription) PangoStretch

	Pango_font_description_set_size func(
		desc *PangoFontDescription,
		size Gint)

	Pango_font_description_get_size func(
		desc *PangoFontDescription) Gint

	Pango_font_description_set_absolute_size func(
		desc *PangoFontDescription,
		size Double)

	Pango_font_description_get_size_is_absolute func(
		desc *PangoFontDescription) Gboolean

	Pango_font_description_set_gravity func(
		desc *PangoFontDescription,
		gravity PangoGravity)

	Pango_font_description_get_gravity func(
		desc *PangoFontDescription) PangoGravity

	Pango_font_description_get_set_fields func(
		desc *PangoFontDescription) PangoFontMask

	Pango_font_description_unset_fields func(
		desc *PangoFontDescription,
		to_unset PangoFontMask)

	Pango_font_description_merge func(
		desc *PangoFontDescription,
		desc_to_merge *PangoFontDescription,
		replace_existing Gboolean)

	Pango_font_description_merge_static func(
		desc *PangoFontDescription,
		desc_to_merge *PangoFontDescription,
		replace_existing Gboolean)

	Pango_font_description_better_match func(
		desc *PangoFontDescription,
		old_match *PangoFontDescription,
		new_match *PangoFontDescription) Gboolean

	Pango_font_description_from_string func(
		str string) *PangoFontDescription

	Pango_font_description_to_string func(
		desc *PangoFontDescription) string

	Pango_font_description_to_filename func(
		desc *PangoFontDescription) string

	Pango_font_metrics_get_type func() GType

	Pango_font_metrics_ref func(
		metrics *PangoFontMetrics) *PangoFontMetrics

	Pango_font_metrics_unref func(
		metrics *PangoFontMetrics)

	Pango_font_metrics_get_ascent func(
		metrics *PangoFontMetrics) int

	Pango_font_metrics_get_descent func(
		metrics *PangoFontMetrics) int

	Pango_font_metrics_get_approximate_char_width func(
		metrics *PangoFontMetrics) int

	Pango_font_metrics_get_approximate_digit_width func(
		metrics *PangoFontMetrics) int

	Pango_font_metrics_get_underline_position func(
		metrics *PangoFontMetrics) int

	Pango_font_metrics_get_underline_thickness func(
		metrics *PangoFontMetrics) int

	Pango_font_metrics_get_strikethrough_position func(
		metrics *PangoFontMetrics) int

	Pango_font_metrics_get_strikethrough_thickness func(
		metrics *PangoFontMetrics) int

	Pango_font_family_get_type func() GType

	Pango_font_family_list_faces func(
		family *PangoFontFamily,
		faces ***PangoFontFace,
		n_faces *int)

	Pango_font_family_get_name func(
		family *PangoFontFamily) string

	Pango_font_family_is_monospace func(
		family *PangoFontFamily) Gboolean

	Pango_font_face_get_type func() GType

	Pango_font_face_describe func(
		face *PangoFontFace) *PangoFontDescription

	Pango_font_face_get_face_name func(
		face *PangoFontFace) string

	Pango_font_face_list_sizes func(
		face *PangoFontFace,
		sizes **int,
		n_sizes *int)

	Pango_font_face_is_synthesized func(
		face *PangoFontFace) Gboolean

	Pango_font_get_type func() GType

	Pango_font_describe func(
		font *PangoFont) *PangoFontDescription

	Pango_font_describe_with_absolute_size func(
		font *PangoFont) *PangoFontDescription

	Pango_font_get_coverage func(
		font *PangoFont,
		language *PangoLanguage) *PangoCoverage

	Pango_font_find_shaper func(
		font *PangoFont,
		language *PangoLanguage,
		ch Guint32) *PangoEngineShape

	Pango_font_get_metrics func(
		font *PangoFont,
		language *PangoLanguage) *PangoFontMetrics

	Pango_font_get_glyph_extents func(
		font *PangoFont,
		glyph PangoGlyph,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_font_get_font_map func(
		font *PangoFont) *PangoFontMap

	Pango_color_get_type func() GType

	Pango_color_copy func(
		src *PangoColor) *PangoColor

	Pango_color_free func(
		color *PangoColor)

	Pango_color_parse func(
		color *PangoColor,
		spec string) Gboolean

	Pango_color_to_string func(
		color *PangoColor) string

	Pango_attr_type_register func(
		name string) PangoAttrType

	Pango_attr_type_get_name func(
		t PangoAttrType) string

	Pango_attribute_init func(
		attr *PangoAttribute,
		klass *PangoAttrClass)

	Pango_attribute_copy func(
		attr *PangoAttribute) *PangoAttribute

	Pango_attribute_destroy func(
		attr *PangoAttribute)

	Pango_attribute_equal func(
		attr1 *PangoAttribute,
		attr2 *PangoAttribute) Gboolean

	Pango_attr_language_new func(
		language *PangoLanguage) *PangoAttribute

	Pango_attr_family_new func(
		family string) *PangoAttribute

	Pango_attr_foreground_new func(
		red Guint16,
		green Guint16,
		blue Guint16) *PangoAttribute

	Pango_attr_background_new func(
		red Guint16,
		green Guint16,
		blue Guint16) *PangoAttribute

	Pango_attr_size_new func(
		size int) *PangoAttribute

	Pango_attr_size_new_absolute func(
		size int) *PangoAttribute

	Pango_attr_style_new func(
		style PangoStyle) *PangoAttribute

	Pango_attr_weight_new func(
		weight PangoWeight) *PangoAttribute

	Pango_attr_variant_new func(
		variant PangoVariant) *PangoAttribute

	Pango_attr_stretch_new func(
		stretch PangoStretch) *PangoAttribute

	Pango_attr_font_desc_new func(
		desc *PangoFontDescription) *PangoAttribute

	Pango_attr_underline_new func(
		underline PangoUnderline) *PangoAttribute

	Pango_attr_underline_color_new func(
		red Guint16,
		green Guint16,
		blue Guint16) *PangoAttribute

	Pango_attr_strikethrough_new func(
		strikethrough Gboolean) *PangoAttribute

	Pango_attr_strikethrough_color_new func(
		red Guint16,
		green Guint16,
		blue Guint16) *PangoAttribute

	Pango_attr_rise_new func(
		rise int) *PangoAttribute

	Pango_attr_scale_new func(
		scale_factor Double) *PangoAttribute

	Pango_attr_fallback_new func(
		enable_fallback Gboolean) *PangoAttribute

	Pango_attr_letter_spacing_new func(
		letter_spacing int) *PangoAttribute

	Pango_attr_shape_new func(
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle) *PangoAttribute

	Pango_attr_shape_new_with_data func(
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle,
		data Gpointer,
		copy_func PangoAttrDataCopyFunc,
		destroy_func GDestroyNotify) *PangoAttribute

	Pango_attr_gravity_new func(
		gravity PangoGravity) *PangoAttribute

	Pango_attr_gravity_hint_new func(
		hint PangoGravityHint) *PangoAttribute

	Pango_attr_list_get_type func() GType

	Pango_attr_list_new func() *PangoAttrList

	Pango_attr_list_ref func(
		list *PangoAttrList) *PangoAttrList

	Pango_attr_list_unref func(
		list *PangoAttrList)

	Pango_attr_list_copy func(
		list *PangoAttrList) *PangoAttrList

	Pango_attr_list_insert func(
		list *PangoAttrList,
		attr *PangoAttribute)

	Pango_attr_list_insert_before func(
		list *PangoAttrList,
		attr *PangoAttribute)

	Pango_attr_list_change func(
		list *PangoAttrList,
		attr *PangoAttribute)

	Pango_attr_list_splice func(
		list *PangoAttrList,
		other *PangoAttrList,
		pos Gint,
		len Gint)

	Pango_attr_list_filter func(
		list *PangoAttrList,
		f PangoAttrFilterFunc,
		data Gpointer) *PangoAttrList

	Pango_attr_list_get_iterator func(
		list *PangoAttrList) *PangoAttrIterator

	Pango_attr_iterator_range func(
		iterator *PangoAttrIterator,
		start *Gint,
		end *Gint)

	Pango_attr_iterator_next func(
		iterator *PangoAttrIterator) Gboolean

	Pango_attr_iterator_copy func(
		iterator *PangoAttrIterator) *PangoAttrIterator

	Pango_attr_iterator_destroy func(
		iterator *PangoAttrIterator)

	Pango_attr_iterator_get func(
		iterator *PangoAttrIterator,
		t PangoAttrType) *PangoAttribute

	Pango_attr_iterator_get_font func(
		iterator *PangoAttrIterator,
		desc *PangoFontDescription,
		language **PangoLanguage,
		extra_attrs **GSList)

	Pango_attr_iterator_get_attrs func(
		iterator *PangoAttrIterator) *GSList

	Pango_parse_markup func(
		markup_text string,
		length int,
		accel_marker Gunichar,
		attr_list **PangoAttrList,
		text **Char,
		accel_char *Gunichar,
		error **GError) Gboolean

	Pango_item_get_type func() GType

	Pango_item_new func() *PangoItem

	Pango_item_copy func(
		item *PangoItem) *PangoItem

	Pango_item_free func(
		item *PangoItem)

	Pango_item_split func(
		orig *PangoItem,
		split_index int,
		split_offset int) *PangoItem

	Pango_break func(
		text string,
		length int,
		analysis *PangoAnalysis,
		attrs *PangoLogAttr,
		attrs_len int)

	Pango_find_paragraph_boundary func(
		text string,
		length Gint,
		paragraph_delimiter_index *Gint,
		next_paragraph_start *Gint)

	Pango_get_log_attrs func(
		text string,
		length int,
		level int,
		language *PangoLanguage,
		log_attrs *PangoLogAttr,
		attrs_len int)

	Pango_fontset_get_type func() GType

	Pango_fontset_get_font func(
		fontset *PangoFontset,
		wc Guint) *PangoFont

	Pango_fontset_get_metrics func(
		fontset *PangoFontset) *PangoFontMetrics

	Pango_fontset_foreach func(
		fontset *PangoFontset,
		f PangoFontsetForeachFunc,
		data Gpointer)

	Pango_font_map_get_type func() GType

	Pango_font_map_create_context func(
		fontmap *PangoFontMap) *PangoContext

	Pango_font_map_load_font func(
		fontmap *PangoFontMap,
		context *PangoContext,
		desc *PangoFontDescription) *PangoFont

	Pango_font_map_load_fontset func(
		fontmap *PangoFontMap,
		context *PangoContext,
		desc *PangoFontDescription,
		language *PangoLanguage) *PangoFontset

	Pango_font_map_list_families func(
		fontmap *PangoFontMap,
		families ***PangoFontFamily,
		n_families *int)

	Pango_context_get_type func() GType

	Pango_context_new func() *PangoContext

	Pango_context_set_font_map func(
		context *PangoContext,
		font_map *PangoFontMap)

	Pango_context_get_font_map func(
		context *PangoContext) *PangoFontMap

	Pango_context_list_families func(
		context *PangoContext,
		families ***PangoFontFamily,
		n_families *int)

	Pango_context_load_font func(
		context *PangoContext,
		desc *PangoFontDescription) *PangoFont

	Pango_context_load_fontset func(
		context *PangoContext,
		desc *PangoFontDescription,
		language *PangoLanguage) *PangoFontset

	Pango_context_get_metrics func(
		context *PangoContext,
		desc *PangoFontDescription,
		language *PangoLanguage) *PangoFontMetrics

	Pango_context_set_font_description func(
		context *PangoContext,
		desc *PangoFontDescription)

	Pango_context_get_font_description func(
		context *PangoContext) *PangoFontDescription

	Pango_context_get_language func(
		context *PangoContext) *PangoLanguage

	Pango_context_set_language func(
		context *PangoContext,
		language *PangoLanguage)

	Pango_context_set_base_dir func(
		context *PangoContext,
		direction PangoDirection)

	Pango_context_get_base_dir func(
		context *PangoContext) PangoDirection

	Pango_context_set_base_gravity func(
		context *PangoContext,
		gravity PangoGravity)

	Pango_context_get_base_gravity func(
		context *PangoContext) PangoGravity

	Pango_context_get_gravity func(
		context *PangoContext) PangoGravity

	Pango_context_set_gravity_hint func(
		context *PangoContext,
		hint PangoGravityHint)

	Pango_context_get_gravity_hint func(
		context *PangoContext) PangoGravityHint

	Pango_context_set_matrix func(
		context *PangoContext,
		matrix *PangoMatrix)

	Pango_context_get_matrix func(
		context *PangoContext) *PangoMatrix

	Pango_itemize func(
		context *PangoContext,
		text string,
		start_index int,
		length int,
		attrs *PangoAttrList,
		cached_iter *PangoAttrIterator) *GList

	Pango_itemize_with_base_dir func(
		context *PangoContext,
		base_dir PangoDirection,
		text string,
		start_index int,
		length int,
		attrs *PangoAttrList,
		cached_iter *PangoAttrIterator) *GList

	Pango_glyph_string_new func() *PangoGlyphString

	Pango_glyph_string_set_size func(
		string *PangoGlyphString,
		new_len Gint)

	Pango_glyph_string_get_type func() GType

	Pango_glyph_string_copy func(
		string *PangoGlyphString) *PangoGlyphString

	Pango_glyph_string_free func(
		string *PangoGlyphString)

	Pango_glyph_string_extents func(
		glyphs *PangoGlyphString,
		font *PangoFont,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_glyph_string_get_width func(
		glyphs *PangoGlyphString) int

	Pango_glyph_string_extents_range func(
		glyphs *PangoGlyphString,
		start int,
		end int,
		font *PangoFont,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_glyph_string_get_logical_widths func(
		glyphs *PangoGlyphString,
		text string,
		length int,
		embedding_level int,
		logical_widths *int)

	Pango_glyph_string_index_to_x func(
		glyphs *PangoGlyphString,
		text string,
		length int,
		analysis *PangoAnalysis,
		index int,
		trailing Gboolean,
		x_pos *int)

	Pango_glyph_string_x_to_index func(
		glyphs *PangoGlyphString,
		text string,
		length int,
		analysis *PangoAnalysis,
		x_pos int,
		index *int,
		trailing *int)

	Pango_shape func(
		text string,
		length Gint,
		analysis *PangoAnalysis,
		glyphs *PangoGlyphString)

	Pango_reorder_items func(
		logical_items *GList) *GList

	Pango_attr_type_get_type func() GType

	Pango_underline_get_type func() GType

	Pango_bidi_type_get_type func() GType

	Pango_direction_get_type func() GType

	Pango_coverage_level_get_type func() GType

	Pango_style_get_type func() GType

	Pango_variant_get_type func() GType

	Pango_weight_get_type func() GType

	Pango_stretch_get_type func() GType

	Pango_font_mask_get_type func() GType

	Pango_gravity_get_type func() GType

	Pango_gravity_hint_get_type func() GType

	Pango_alignment_get_type func() GType

	Pango_wrap_mode_get_type func() GType

	Pango_ellipsize_mode_get_type func() GType

	Pango_render_part_get_type func() GType

	Pango_script_get_type func() GType

	Pango_tab_align_get_type func() GType

	Pango_glyph_item_get_type func() GType

	Pango_glyph_item_split func(
		orig *PangoGlyphItem,
		text string,
		split_index int) *PangoGlyphItem

	Pango_glyph_item_copy func(
		orig *PangoGlyphItem) *PangoGlyphItem

	Pango_glyph_item_free func(
		glyph_item *PangoGlyphItem)

	Pango_glyph_item_apply_attrs func(
		glyph_item *PangoGlyphItem,
		text string,
		list *PangoAttrList) *GSList

	Pango_glyph_item_letter_space func(
		glyph_item *PangoGlyphItem,
		text string,
		log_attrs *PangoLogAttr,
		letter_spacing int)

	Pango_glyph_item_get_logical_widths func(
		glyph_item *PangoGlyphItem,
		text string,
		logical_widths *int)

	Pango_glyph_item_iter_get_type func() GType

	Pango_glyph_item_iter_copy func(
		orig *PangoGlyphItemIter) *PangoGlyphItemIter

	Pango_glyph_item_iter_free func(
		iter *PangoGlyphItemIter)

	Pango_glyph_item_iter_init_start func(
		iter *PangoGlyphItemIter,
		glyph_item *PangoGlyphItem,
		text string) Gboolean

	Pango_glyph_item_iter_init_end func(
		iter *PangoGlyphItemIter,
		glyph_item *PangoGlyphItem,
		text string) Gboolean

	Pango_glyph_item_iter_next_cluster func(
		iter *PangoGlyphItemIter) Gboolean

	Pango_glyph_item_iter_prev_cluster func(
		iter *PangoGlyphItemIter) Gboolean

	Pango_tab_array_new func(
		initial_size Gint,
		positions_in_pixels Gboolean) *PangoTabArray

	Pango_tab_array_new_with_positions func(size Gint,
		positions_in_pixels Gboolean, first_alignment PangoTabAlign,
		first_position Gint, v ...VArg) *PangoTabArray

	Pango_tab_array_get_type func() GType

	Pango_tab_array_copy func(
		src *PangoTabArray) *PangoTabArray

	Pango_tab_array_free func(
		tab_array *PangoTabArray)

	Pango_tab_array_get_size func(
		tab_array *PangoTabArray) Gint

	Pango_tab_array_resize func(
		tab_array *PangoTabArray,
		new_size Gint)

	Pango_tab_array_set_tab func(
		tab_array *PangoTabArray,
		tab_index Gint,
		alignment PangoTabAlign,
		location Gint)

	Pango_tab_array_get_tab func(
		tab_array *PangoTabArray,
		tab_index Gint,
		alignment *PangoTabAlign,
		location *Gint)

	Pango_tab_array_get_tabs func(
		tab_array *PangoTabArray,
		alignments **PangoTabAlign,
		locations **Gint)

	Pango_tab_array_get_positions_in_pixels func(
		tab_array *PangoTabArray) Gboolean

	Pango_layout_get_type func() GType

	Pango_layout_new func(
		context *PangoContext) *PangoLayout

	Pango_layout_copy func(
		src *PangoLayout) *PangoLayout

	Pango_layout_get_context func(
		layout *PangoLayout) *PangoContext

	Pango_layout_set_attributes func(
		layout *PangoLayout,
		attrs *PangoAttrList)

	Pango_layout_get_attributes func(
		layout *PangoLayout) *PangoAttrList

	Pango_layout_set_text func(
		layout *PangoLayout,
		text string,
		length int)

	Pango_layout_get_text func(
		layout *PangoLayout) string

	Pango_layout_get_character_count func(
		layout *PangoLayout) Gint

	Pango_layout_set_markup func(
		layout *PangoLayout,
		markup string,
		length int)

	Pango_layout_set_markup_with_accel func(
		layout *PangoLayout,
		markup string,
		length int,
		accel_marker Gunichar,
		accel_char *Gunichar)

	Pango_layout_set_font_description func(
		layout *PangoLayout,
		desc *PangoFontDescription)

	Pango_layout_get_font_description func(
		layout *PangoLayout) *PangoFontDescription

	Pango_layout_set_width func(
		layout *PangoLayout,
		width int)

	Pango_layout_get_width func(
		layout *PangoLayout) int

	Pango_layout_set_height func(
		layout *PangoLayout,
		height int)

	Pango_layout_get_height func(
		layout *PangoLayout) int

	Pango_layout_set_wrap func(
		layout *PangoLayout,
		wrap PangoWrapMode)

	Pango_layout_get_wrap func(
		layout *PangoLayout) PangoWrapMode

	Pango_layout_is_wrapped func(
		layout *PangoLayout) Gboolean

	Pango_layout_set_indent func(
		layout *PangoLayout,
		indent int)

	Pango_layout_get_indent func(
		layout *PangoLayout) int

	Pango_layout_set_spacing func(
		layout *PangoLayout,
		spacing int)

	Pango_layout_get_spacing func(
		layout *PangoLayout) int

	Pango_layout_set_justify func(
		layout *PangoLayout,
		justify Gboolean)

	Pango_layout_get_justify func(
		layout *PangoLayout) Gboolean

	Pango_layout_set_auto_dir func(
		layout *PangoLayout,
		auto_dir Gboolean)

	Pango_layout_get_auto_dir func(
		layout *PangoLayout) Gboolean

	Pango_layout_set_alignment func(
		layout *PangoLayout,
		alignment PangoAlignment)

	Pango_layout_get_alignment func(
		layout *PangoLayout) PangoAlignment

	Pango_layout_set_tabs func(
		layout *PangoLayout,
		tabs *PangoTabArray)

	Pango_layout_get_tabs func(
		layout *PangoLayout) *PangoTabArray

	Pango_layout_set_single_paragraph_mode func(
		layout *PangoLayout,
		setting Gboolean)

	Pango_layout_get_single_paragraph_mode func(
		layout *PangoLayout) Gboolean

	Pango_layout_set_ellipsize func(
		layout *PangoLayout,
		ellipsize PangoEllipsizeMode)

	Pango_layout_get_ellipsize func(
		layout *PangoLayout) PangoEllipsizeMode

	Pango_layout_is_ellipsized func(
		layout *PangoLayout) Gboolean

	Pango_layout_get_unknown_glyphs_count func(
		layout *PangoLayout) int

	Pango_layout_context_changed func(
		layout *PangoLayout)

	Pango_layout_get_log_attrs func(
		layout *PangoLayout,
		attrs **PangoLogAttr,
		n_attrs *Gint)

	Pango_layout_get_log_attrs_readonly func(
		layout *PangoLayout,
		n_attrs *Gint) *PangoLogAttr

	Pango_layout_index_to_pos func(
		layout *PangoLayout,
		index int,
		pos *PangoRectangle)

	Pango_layout_index_to_line_x func(
		layout *PangoLayout,
		index int,
		trailing Gboolean,
		line *int,
		x_pos *int)

	Pango_layout_get_cursor_pos func(
		layout *PangoLayout,
		index int,
		strong_pos *PangoRectangle,
		weak_pos *PangoRectangle)

	Pango_layout_move_cursor_visually func(
		layout *PangoLayout,
		strong Gboolean,
		old_index int,
		old_trailing int,
		direction int,
		new_index *int,
		new_trailing *int)

	Pango_layout_xy_to_index func(
		layout *PangoLayout,
		x int,
		y int,
		index *int,
		trailing *int) Gboolean

	Pango_layout_get_extents func(
		layout *PangoLayout,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_layout_get_pixel_extents func(
		layout *PangoLayout,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_layout_get_size func(
		layout *PangoLayout,
		width *int,
		height *int)

	Pango_layout_get_pixel_size func(
		layout *PangoLayout,
		width *int,
		height *int)

	Pango_layout_get_baseline func(
		layout *PangoLayout) int

	Pango_layout_get_line_count func(
		layout *PangoLayout) int

	Pango_layout_get_line func(
		layout *PangoLayout,
		line int) *PangoLayoutLine

	Pango_layout_get_line_readonly func(
		layout *PangoLayout,
		line int) *PangoLayoutLine

	Pango_layout_get_lines func(
		layout *PangoLayout) *GSList

	Pango_layout_get_lines_readonly func(
		layout *PangoLayout) *GSList

	Pango_layout_line_get_type func() GType

	Pango_layout_line_ref func(
		line *PangoLayoutLine) *PangoLayoutLine

	Pango_layout_line_unref func(
		line *PangoLayoutLine)

	Pango_layout_line_x_to_index func(
		line *PangoLayoutLine,
		x_pos int,
		index *int,
		trailing *int) Gboolean

	Pango_layout_line_index_to_x func(
		line *PangoLayoutLine,
		index int,
		trailing Gboolean,
		x_pos *int)

	Pango_layout_line_get_x_ranges func(
		line *PangoLayoutLine,
		start_index int,
		end_index int,
		ranges **int,
		n_ranges *int)

	Pango_layout_line_get_extents func(
		line *PangoLayoutLine,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_layout_line_get_pixel_extents func(
		layout_line *PangoLayoutLine,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_layout_iter_get_type func() GType

	Pango_layout_get_iter func(
		layout *PangoLayout) *PangoLayoutIter

	Pango_layout_iter_copy func(
		iter *PangoLayoutIter) *PangoLayoutIter

	Pango_layout_iter_free func(
		iter *PangoLayoutIter)

	Pango_layout_iter_get_index func(
		iter *PangoLayoutIter) int

	Pango_layout_iter_get_run func(
		iter *PangoLayoutIter) *PangoLayoutRun

	Pango_layout_iter_get_run_readonly func(
		iter *PangoLayoutIter) *PangoLayoutRun

	Pango_layout_iter_get_line func(
		iter *PangoLayoutIter) *PangoLayoutLine

	Pango_layout_iter_get_line_readonly func(
		iter *PangoLayoutIter) *PangoLayoutLine

	Pango_layout_iter_at_last_line func(
		iter *PangoLayoutIter) Gboolean

	Pango_layout_iter_get_layout func(
		iter *PangoLayoutIter) *PangoLayout

	Pango_layout_iter_next_char func(
		iter *PangoLayoutIter) Gboolean

	Pango_layout_iter_next_cluster func(
		iter *PangoLayoutIter) Gboolean

	Pango_layout_iter_next_run func(
		iter *PangoLayoutIter) Gboolean

	Pango_layout_iter_next_line func(
		iter *PangoLayoutIter) Gboolean

	Pango_layout_iter_get_char_extents func(
		iter *PangoLayoutIter,
		logical_rect *PangoRectangle)

	Pango_layout_iter_get_cluster_extents func(
		iter *PangoLayoutIter,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_layout_iter_get_run_extents func(
		iter *PangoLayoutIter,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_layout_iter_get_line_extents func(
		iter *PangoLayoutIter,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_layout_iter_get_line_yrange func(
		iter *PangoLayoutIter,
		y0 *int,
		y1 *int)

	Pango_layout_iter_get_layout_extents func(
		iter *PangoLayoutIter,
		ink_rect *PangoRectangle,
		logical_rect *PangoRectangle)

	Pango_layout_iter_get_baseline func(
		iter *PangoLayoutIter) int

	Pango_renderer_get_type func() GType

	Pango_renderer_draw_layout func(
		renderer *PangoRenderer,
		layout *PangoLayout,
		x int,
		y int)

	Pango_renderer_draw_layout_line func(
		renderer *PangoRenderer,
		line *PangoLayoutLine,
		x int,
		y int)

	Pango_renderer_draw_glyphs func(
		renderer *PangoRenderer,
		font *PangoFont,
		glyphs *PangoGlyphString,
		x int,
		y int)

	Pango_renderer_draw_glyph_item func(
		renderer *PangoRenderer,
		text string,
		glyph_item *PangoGlyphItem,
		x int,
		y int)

	Pango_renderer_draw_rectangle func(
		renderer *PangoRenderer,
		part PangoRenderPart,
		x int,
		y int,
		width int,
		height int)

	Pango_renderer_draw_error_underline func(
		renderer *PangoRenderer,
		x int,
		y int,
		width int,
		height int)

	Pango_renderer_draw_trapezoid func(
		renderer *PangoRenderer,
		part PangoRenderPart,
		y1 Double,
		x11 Double,
		x21 Double,
		y2 Double,
		x12 Double,
		x22 Double)

	Pango_renderer_draw_glyph func(
		renderer *PangoRenderer,
		font *PangoFont,
		glyph PangoGlyph,
		x Double,
		y Double)

	Pango_renderer_activate func(
		renderer *PangoRenderer)

	Pango_renderer_deactivate func(
		renderer *PangoRenderer)

	Pango_renderer_part_changed func(
		renderer *PangoRenderer,
		part PangoRenderPart)

	Pango_renderer_set_color func(
		renderer *PangoRenderer,
		part PangoRenderPart,
		color *PangoColor)

	Pango_renderer_get_color func(
		renderer *PangoRenderer,
		part PangoRenderPart) *PangoColor

	Pango_renderer_set_matrix func(
		renderer *PangoRenderer,
		matrix *PangoMatrix)

	Pango_renderer_get_matrix func(
		renderer *PangoRenderer) *PangoMatrix

	Pango_renderer_get_layout func(
		renderer *PangoRenderer) *PangoLayout

	Pango_renderer_get_layout_line func(
		renderer *PangoRenderer) *PangoLayoutLine

	Pango_split_file_list func(
		str string) **Char

	Pango_trim_string func(
		str string) string

	Pango_read_line func(
		stream *FILE,
		str *GString) Gint

	Pango_skip_space func(
		pos **Char) Gboolean

	Pango_scan_word func(
		pos **Char,
		out *GString) Gboolean

	Pango_scan_string func(
		pos **Char,
		out *GString) Gboolean

	Pango_scan_int func(
		pos **Char,
		out *int) Gboolean

	Pango_parse_enum func(
		t GType,
		str string,
		value *int,
		warn Gboolean,
		possible_values **Char) Gboolean

	Pango_parse_style func(
		str string,
		style *PangoStyle,
		warn Gboolean) Gboolean

	Pango_parse_variant func(
		str string,
		variant *PangoVariant,
		warn Gboolean) Gboolean

	Pango_parse_weight func(
		str string,
		weight *PangoWeight,
		warn Gboolean) Gboolean

	Pango_parse_stretch func(
		str string,
		stretch *PangoStretch,
		warn Gboolean) Gboolean

	Pango_quantize_line_geometry func(
		thickness *int,
		position *int)

	Pango_log2vis_get_embedding_levels func(
		text string,
		length int,
		pbase_dir *PangoDirection) *Guint8

	Pango_is_zero_width func(
		ch Gunichar) Gboolean

	Pango_version func() int

	Pango_version_string func() string

	Pango_version_check func(
		required_major int,
		required_minor int,
		required_micro int) string

	Pango_cairo_font_map_get_type func() GType

	Pango_cairo_font_map_new func() *PangoFontMap

	Pango_cairo_font_map_new_for_font_type func(
		fonttype Cairo_font_type_t) *PangoFontMap

	Pango_cairo_font_map_get_default func() *PangoFontMap

	Pango_cairo_font_map_set_default func(
		fontmap *PangoCairoFontMap)

	Pango_cairo_font_map_get_font_type func(
		fontmap *PangoCairoFontMap) Cairo_font_type_t

	Pango_cairo_font_map_set_resolution func(
		fontmap *PangoCairoFontMap,
		dpi Double)

	Pango_cairo_font_map_get_resolution func(
		fontmap *PangoCairoFontMap) Double

	Pango_cairo_font_map_create_context func(
		fontmap *PangoCairoFontMap) *PangoContext

	Pango_cairo_font_get_type func() GType

	Pango_cairo_font_get_scaled_font func(
		font *PangoCairoFont) *Cairo_scaled_font_t

	Pango_cairo_update_context func(
		cr *Cairo_t,
		context *PangoContext)

	Pango_cairo_context_set_font_options func(
		context *PangoContext,
		options *Cairo_font_options_t)

	Pango_cairo_context_get_font_options func(
		context *PangoContext) *Cairo_font_options_t

	Pango_cairo_context_set_resolution func(
		context *PangoContext,
		dpi Double)

	Pango_cairo_context_get_resolution func(
		context *PangoContext) Double

	Pango_cairo_context_set_shape_renderer func(
		context *PangoContext,
		f PangoCairoShapeRendererFunc,
		data Gpointer,
		dnotify GDestroyNotify)

	Pango_cairo_context_get_shape_renderer func(
		context *PangoContext,
		data *Gpointer) PangoCairoShapeRendererFunc

	Pango_cairo_create_context func(
		cr *Cairo_t) *PangoContext

	Pango_cairo_create_layout func(
		cr *Cairo_t) *PangoLayout

	Pango_cairo_update_layout func(
		cr *Cairo_t,
		layout *PangoLayout)

	Pango_cairo_show_glyph_string func(
		cr *Cairo_t,
		font *PangoFont,
		glyphs *PangoGlyphString)

	Pango_cairo_show_glyph_item func(
		cr *Cairo_t,
		text string,
		glyph_item *PangoGlyphItem)

	Pango_cairo_show_layout_line func(
		cr *Cairo_t,
		line *PangoLayoutLine)

	Pango_cairo_show_layout func(
		cr *Cairo_t,
		layout *PangoLayout)

	Pango_cairo_show_error_underline func(
		cr *Cairo_t,
		x Double,
		y Double,
		width Double,
		height Double)

	Pango_cairo_glyph_string_path func(
		cr *Cairo_t,
		font *PangoFont,
		glyphs *PangoGlyphString)

	Pango_cairo_layout_line_path func(
		cr *Cairo_t,
		line *PangoLayoutLine)

	Pango_cairo_layout_path func(
		cr *Cairo_t,
		layout *PangoLayout)

	Pango_cairo_error_underline_path func(
		cr *Cairo_t,
		x Double,
		y Double,
		width Double,
		height Double)

	Pango_config_key_get func(
		key string) string

	Pango_default_break func(
		text string,
		length int,
		analysis *PangoAnalysis,
		attrs *PangoLogAttr,
		attrs_len int)

	Pango_engine_get_type func() GType

	Pango_engine_lang_get_type func() GType

	Pango_engine_shape_get_type func() GType

	Pango_find_map func(
		language *PangoLanguage,
		engine_type_id Guint,
		render_type_id Guint) *PangoMap

	Pango_font_map_get_shape_engine_type func(
		fontmap *PangoFontMap) string

	Pango_font_metrics_new func() *PangoFontMetrics

	Pango_fontset_simple_append func(
		fontset *PangoFontsetSimple,
		font *PangoFont)

	Pango_fontset_simple_size func(
		fontset *PangoFontsetSimple) int

	Pango_fontset_simple_new func(
		language *PangoLanguage) *PangoFontsetSimple

	Pango_fontset_simple_get_type func() GType

	Pango_get_lib_subdirectory func() string

	Pango_get_sysconf_subdirectory func() string

	Pango_lookup_aliases func(
		fontname string,
		families ***Char,
		n_families *int)

	Pango_map_get_engine func(
		map_ *PangoMap,
		script PangoScript) *PangoEngine

	Pango_map_get_engines func(
		map_ *PangoMap,
		script PangoScript,
		exact_engines **GSList,
		fallback_engines **GSList)

	Pango_module_register func(
		module *PangoIncludedModule)
)

var dll = "libpango-1.0-0.dll"

var apiList = Apis{
	{"pango_alignment_get_type", &Pango_alignment_get_type},
	{"pango_attr_background_new", &Pango_attr_background_new},
	{"pango_attr_fallback_new", &Pango_attr_fallback_new},
	{"pango_attr_family_new", &Pango_attr_family_new},
	{"pango_attr_font_desc_new", &Pango_attr_font_desc_new},
	{"pango_attr_foreground_new", &Pango_attr_foreground_new},
	{"pango_attr_gravity_hint_new", &Pango_attr_gravity_hint_new},
	{"pango_attr_gravity_new", &Pango_attr_gravity_new},
	{"pango_attr_iterator_copy", &Pango_attr_iterator_copy},
	{"pango_attr_iterator_destroy", &Pango_attr_iterator_destroy},
	{"pango_attr_iterator_get", &Pango_attr_iterator_get},
	{"pango_attr_iterator_get_attrs", &Pango_attr_iterator_get_attrs},
	{"pango_attr_iterator_get_font", &Pango_attr_iterator_get_font},
	{"pango_attr_iterator_next", &Pango_attr_iterator_next},
	{"pango_attr_iterator_range", &Pango_attr_iterator_range},
	{"pango_attr_language_new", &Pango_attr_language_new},
	{"pango_attr_letter_spacing_new", &Pango_attr_letter_spacing_new},
	{"pango_attr_list_change", &Pango_attr_list_change},
	{"pango_attr_list_copy", &Pango_attr_list_copy},
	{"pango_attr_list_filter", &Pango_attr_list_filter},
	{"pango_attr_list_get_iterator", &Pango_attr_list_get_iterator},
	{"pango_attr_list_get_type", &Pango_attr_list_get_type},
	{"pango_attr_list_insert", &Pango_attr_list_insert},
	{"pango_attr_list_insert_before", &Pango_attr_list_insert_before},
	{"pango_attr_list_new", &Pango_attr_list_new},
	{"pango_attr_list_ref", &Pango_attr_list_ref},
	{"pango_attr_list_splice", &Pango_attr_list_splice},
	{"pango_attr_list_unref", &Pango_attr_list_unref},
	{"pango_attr_rise_new", &Pango_attr_rise_new},
	{"pango_attr_scale_new", &Pango_attr_scale_new},
	{"pango_attr_shape_new", &Pango_attr_shape_new},
	{"pango_attr_shape_new_with_data", &Pango_attr_shape_new_with_data},
	{"pango_attr_size_new", &Pango_attr_size_new},
	{"pango_attr_size_new_absolute", &Pango_attr_size_new_absolute},
	{"pango_attr_stretch_new", &Pango_attr_stretch_new},
	{"pango_attr_strikethrough_color_new", &Pango_attr_strikethrough_color_new},
	{"pango_attr_strikethrough_new", &Pango_attr_strikethrough_new},
	{"pango_attr_style_new", &Pango_attr_style_new},
	{"pango_attr_type_get_name", &Pango_attr_type_get_name},
	{"pango_attr_type_get_type", &Pango_attr_type_get_type},
	{"pango_attr_type_register", &Pango_attr_type_register},
	{"pango_attr_underline_color_new", &Pango_attr_underline_color_new},
	{"pango_attr_underline_new", &Pango_attr_underline_new},
	{"pango_attr_variant_new", &Pango_attr_variant_new},
	{"pango_attr_weight_new", &Pango_attr_weight_new},
	{"pango_attribute_copy", &Pango_attribute_copy},
	{"pango_attribute_destroy", &Pango_attribute_destroy},
	{"pango_attribute_equal", &Pango_attribute_equal},
	{"pango_attribute_init", &Pango_attribute_init},
	{"pango_bidi_type_for_unichar", &Pango_bidi_type_for_unichar},
	{"pango_bidi_type_get_type", &Pango_bidi_type_get_type},
	{"pango_break", &Pango_break},
	{"pango_color_copy", &Pango_color_copy},
	{"pango_color_free", &Pango_color_free},
	{"pango_color_get_type", &Pango_color_get_type},
	{"pango_color_parse", &Pango_color_parse},
	{"pango_color_to_string", &Pango_color_to_string},
	{"pango_config_key_get", &Pango_config_key_get},
	{"pango_context_get_base_dir", &Pango_context_get_base_dir},
	{"pango_context_get_base_gravity", &Pango_context_get_base_gravity},
	{"pango_context_get_font_description", &Pango_context_get_font_description},
	{"pango_context_get_font_map", &Pango_context_get_font_map},
	{"pango_context_get_gravity", &Pango_context_get_gravity},
	{"pango_context_get_gravity_hint", &Pango_context_get_gravity_hint},
	{"pango_context_get_language", &Pango_context_get_language},
	{"pango_context_get_matrix", &Pango_context_get_matrix},
	{"pango_context_get_metrics", &Pango_context_get_metrics},
	{"pango_context_get_type", &Pango_context_get_type},
	{"pango_context_list_families", &Pango_context_list_families},
	{"pango_context_load_font", &Pango_context_load_font},
	{"pango_context_load_fontset", &Pango_context_load_fontset},
	{"pango_context_new", &Pango_context_new},
	{"pango_context_set_base_dir", &Pango_context_set_base_dir},
	{"pango_context_set_base_gravity", &Pango_context_set_base_gravity},
	{"pango_context_set_font_description", &Pango_context_set_font_description},
	{"pango_context_set_font_map", &Pango_context_set_font_map},
	{"pango_context_set_gravity_hint", &Pango_context_set_gravity_hint},
	{"pango_context_set_language", &Pango_context_set_language},
	{"pango_context_set_matrix", &Pango_context_set_matrix},
	{"pango_coverage_copy", &Pango_coverage_copy},
	{"pango_coverage_from_bytes", &Pango_coverage_from_bytes},
	{"pango_coverage_get", &Pango_coverage_get},
	{"pango_coverage_level_get_type", &Pango_coverage_level_get_type},
	{"pango_coverage_max", &Pango_coverage_max},
	{"pango_coverage_new", &Pango_coverage_new},
	{"pango_coverage_ref", &Pango_coverage_ref},
	{"pango_coverage_set", &Pango_coverage_set},
	{"pango_coverage_to_bytes", &Pango_coverage_to_bytes},
	{"pango_coverage_unref", &Pango_coverage_unref},
	{"pango_default_break", &Pango_default_break},
	{"pango_direction_get_type", &Pango_direction_get_type},
	{"pango_ellipsize_mode_get_type", &Pango_ellipsize_mode_get_type},
	{"pango_engine_get_type", &Pango_engine_get_type},
	{"pango_engine_lang_get_type", &Pango_engine_lang_get_type},
	{"pango_engine_shape_get_type", &Pango_engine_shape_get_type},
	{"pango_extents_to_pixels", &Pango_extents_to_pixels},
	{"pango_find_base_dir", &Pango_find_base_dir},
	{"pango_find_map", &Pango_find_map},
	{"pango_find_paragraph_boundary", &Pango_find_paragraph_boundary},
	{"pango_font_describe", &Pango_font_describe},
	{"pango_font_describe_with_absolute_size", &Pango_font_describe_with_absolute_size},
	{"pango_font_description_better_match", &Pango_font_description_better_match},
	{"pango_font_description_copy", &Pango_font_description_copy},
	{"pango_font_description_copy_static", &Pango_font_description_copy_static},
	{"pango_font_description_equal", &Pango_font_description_equal},
	{"pango_font_description_free", &Pango_font_description_free},
	{"pango_font_description_from_string", &Pango_font_description_from_string},
	{"pango_font_description_get_family", &Pango_font_description_get_family},
	{"pango_font_description_get_gravity", &Pango_font_description_get_gravity},
	{"pango_font_description_get_set_fields", &Pango_font_description_get_set_fields},
	{"pango_font_description_get_size", &Pango_font_description_get_size},
	{"pango_font_description_get_size_is_absolute", &Pango_font_description_get_size_is_absolute},
	{"pango_font_description_get_stretch", &Pango_font_description_get_stretch},
	{"pango_font_description_get_style", &Pango_font_description_get_style},
	{"pango_font_description_get_type", &Pango_font_description_get_type},
	{"pango_font_description_get_variant", &Pango_font_description_get_variant},
	{"pango_font_description_get_weight", &Pango_font_description_get_weight},
	{"pango_font_description_hash", &Pango_font_description_hash},
	{"pango_font_description_merge", &Pango_font_description_merge},
	{"pango_font_description_merge_static", &Pango_font_description_merge_static},
	{"pango_font_description_new", &Pango_font_description_new},
	{"pango_font_description_set_absolute_size", &Pango_font_description_set_absolute_size},
	{"pango_font_description_set_family", &Pango_font_description_set_family},
	{"pango_font_description_set_family_static", &Pango_font_description_set_family_static},
	{"pango_font_description_set_gravity", &Pango_font_description_set_gravity},
	{"pango_font_description_set_size", &Pango_font_description_set_size},
	{"pango_font_description_set_stretch", &Pango_font_description_set_stretch},
	{"pango_font_description_set_style", &Pango_font_description_set_style},
	{"pango_font_description_set_variant", &Pango_font_description_set_variant},
	{"pango_font_description_set_weight", &Pango_font_description_set_weight},
	{"pango_font_description_to_filename", &Pango_font_description_to_filename},
	{"pango_font_description_to_string", &Pango_font_description_to_string},
	{"pango_font_description_unset_fields", &Pango_font_description_unset_fields},
	{"pango_font_descriptions_free", &Pango_font_descriptions_free},
	{"pango_font_face_describe", &Pango_font_face_describe},
	{"pango_font_face_get_face_name", &Pango_font_face_get_face_name},
	{"pango_font_face_get_type", &Pango_font_face_get_type},
	{"pango_font_face_is_synthesized", &Pango_font_face_is_synthesized},
	{"pango_font_face_list_sizes", &Pango_font_face_list_sizes},
	{"pango_font_family_get_name", &Pango_font_family_get_name},
	{"pango_font_family_get_type", &Pango_font_family_get_type},
	{"pango_font_family_is_monospace", &Pango_font_family_is_monospace},
	{"pango_font_family_list_faces", &Pango_font_family_list_faces},
	{"pango_font_find_shaper", &Pango_font_find_shaper},
	{"pango_font_get_coverage", &Pango_font_get_coverage},
	{"pango_font_get_font_map", &Pango_font_get_font_map},
	{"pango_font_get_glyph_extents", &Pango_font_get_glyph_extents},
	{"pango_font_get_metrics", &Pango_font_get_metrics},
	{"pango_font_get_type", &Pango_font_get_type},
	{"pango_font_map_create_context", &Pango_font_map_create_context},
	{"pango_font_map_get_shape_engine_type", &Pango_font_map_get_shape_engine_type},
	{"pango_font_map_get_type", &Pango_font_map_get_type},
	{"pango_font_map_list_families", &Pango_font_map_list_families},
	{"pango_font_map_load_font", &Pango_font_map_load_font},
	{"pango_font_map_load_fontset", &Pango_font_map_load_fontset},
	{"pango_font_mask_get_type", &Pango_font_mask_get_type},
	{"pango_font_metrics_get_approximate_char_width", &Pango_font_metrics_get_approximate_char_width},
	{"pango_font_metrics_get_approximate_digit_width", &Pango_font_metrics_get_approximate_digit_width},
	{"pango_font_metrics_get_ascent", &Pango_font_metrics_get_ascent},
	{"pango_font_metrics_get_descent", &Pango_font_metrics_get_descent},
	{"pango_font_metrics_get_strikethrough_position", &Pango_font_metrics_get_strikethrough_position},
	{"pango_font_metrics_get_strikethrough_thickness", &Pango_font_metrics_get_strikethrough_thickness},
	{"pango_font_metrics_get_type", &Pango_font_metrics_get_type},
	{"pango_font_metrics_get_underline_position", &Pango_font_metrics_get_underline_position},
	{"pango_font_metrics_get_underline_thickness", &Pango_font_metrics_get_underline_thickness},
	{"pango_font_metrics_new", &Pango_font_metrics_new},
	{"pango_font_metrics_ref", &Pango_font_metrics_ref},
	{"pango_font_metrics_unref", &Pango_font_metrics_unref},
	{"pango_fontset_foreach", &Pango_fontset_foreach},
	{"pango_fontset_get_font", &Pango_fontset_get_font},
	{"pango_fontset_get_metrics", &Pango_fontset_get_metrics},
	{"pango_fontset_get_type", &Pango_fontset_get_type},
	{"pango_fontset_simple_append", &Pango_fontset_simple_append},
	{"pango_fontset_simple_get_type", &Pango_fontset_simple_get_type},
	{"pango_fontset_simple_new", &Pango_fontset_simple_new},
	{"pango_fontset_simple_size", &Pango_fontset_simple_size},
	{"pango_get_lib_subdirectory", &Pango_get_lib_subdirectory},
	{"pango_get_log_attrs", &Pango_get_log_attrs},
	{"pango_get_mirror_char", &Pango_get_mirror_char},
	{"pango_get_sysconf_subdirectory", &Pango_get_sysconf_subdirectory},
	{"pango_glyph_item_apply_attrs", &Pango_glyph_item_apply_attrs},
	{"pango_glyph_item_copy", &Pango_glyph_item_copy},
	{"pango_glyph_item_free", &Pango_glyph_item_free},
	{"pango_glyph_item_get_logical_widths", &Pango_glyph_item_get_logical_widths},
	{"pango_glyph_item_get_type", &Pango_glyph_item_get_type},
	{"pango_glyph_item_iter_copy", &Pango_glyph_item_iter_copy},
	{"pango_glyph_item_iter_free", &Pango_glyph_item_iter_free},
	{"pango_glyph_item_iter_get_type", &Pango_glyph_item_iter_get_type},
	{"pango_glyph_item_iter_init_end", &Pango_glyph_item_iter_init_end},
	{"pango_glyph_item_iter_init_start", &Pango_glyph_item_iter_init_start},
	{"pango_glyph_item_iter_next_cluster", &Pango_glyph_item_iter_next_cluster},
	{"pango_glyph_item_iter_prev_cluster", &Pango_glyph_item_iter_prev_cluster},
	{"pango_glyph_item_letter_space", &Pango_glyph_item_letter_space},
	{"pango_glyph_item_split", &Pango_glyph_item_split},
	{"pango_glyph_string_copy", &Pango_glyph_string_copy},
	{"pango_glyph_string_extents", &Pango_glyph_string_extents},
	{"pango_glyph_string_extents_range", &Pango_glyph_string_extents_range},
	{"pango_glyph_string_free", &Pango_glyph_string_free},
	{"pango_glyph_string_get_logical_widths", &Pango_glyph_string_get_logical_widths},
	{"pango_glyph_string_get_type", &Pango_glyph_string_get_type},
	{"pango_glyph_string_get_width", &Pango_glyph_string_get_width},
	{"pango_glyph_string_index_to_x", &Pango_glyph_string_index_to_x},
	{"pango_glyph_string_new", &Pango_glyph_string_new},
	{"pango_glyph_string_set_size", &Pango_glyph_string_set_size},
	{"pango_glyph_string_x_to_index", &Pango_glyph_string_x_to_index},
	{"pango_gravity_get_for_matrix", &Pango_gravity_get_for_matrix},
	{"pango_gravity_get_for_script", &Pango_gravity_get_for_script},
	{"pango_gravity_get_for_script_and_width", &Pango_gravity_get_for_script_and_width},
	{"pango_gravity_get_type", &Pango_gravity_get_type},
	{"pango_gravity_hint_get_type", &Pango_gravity_hint_get_type},
	{"pango_gravity_to_rotation", &Pango_gravity_to_rotation},
	{"pango_is_zero_width", &Pango_is_zero_width},
	{"pango_item_copy", &Pango_item_copy},
	{"pango_item_free", &Pango_item_free},
	{"pango_item_get_type", &Pango_item_get_type},
	{"pango_item_new", &Pango_item_new},
	{"pango_item_split", &Pango_item_split},
	{"pango_itemize", &Pango_itemize},
	{"pango_itemize_with_base_dir", &Pango_itemize_with_base_dir},
	{"pango_language_from_string", &Pango_language_from_string},
	{"pango_language_get_default", &Pango_language_get_default},
	{"pango_language_get_sample_string", &Pango_language_get_sample_string},
	{"pango_language_get_scripts", &Pango_language_get_scripts},
	{"pango_language_get_type", &Pango_language_get_type},
	{"pango_language_includes_script", &Pango_language_includes_script},
	{"pango_language_matches", &Pango_language_matches},
	{"pango_language_to_string", &Pango_language_to_string},
	{"pango_layout_context_changed", &Pango_layout_context_changed},
	{"pango_layout_copy", &Pango_layout_copy},
	{"pango_layout_get_alignment", &Pango_layout_get_alignment},
	{"pango_layout_get_attributes", &Pango_layout_get_attributes},
	{"pango_layout_get_auto_dir", &Pango_layout_get_auto_dir},
	{"pango_layout_get_baseline", &Pango_layout_get_baseline},
	{"pango_layout_get_character_count", &Pango_layout_get_character_count},
	{"pango_layout_get_context", &Pango_layout_get_context},
	{"pango_layout_get_cursor_pos", &Pango_layout_get_cursor_pos},
	{"pango_layout_get_ellipsize", &Pango_layout_get_ellipsize},
	{"pango_layout_get_extents", &Pango_layout_get_extents},
	{"pango_layout_get_font_description", &Pango_layout_get_font_description},
	{"pango_layout_get_height", &Pango_layout_get_height},
	{"pango_layout_get_indent", &Pango_layout_get_indent},
	{"pango_layout_get_iter", &Pango_layout_get_iter},
	{"pango_layout_get_justify", &Pango_layout_get_justify},
	{"pango_layout_get_line", &Pango_layout_get_line},
	{"pango_layout_get_line_count", &Pango_layout_get_line_count},
	{"pango_layout_get_line_readonly", &Pango_layout_get_line_readonly},
	{"pango_layout_get_lines", &Pango_layout_get_lines},
	{"pango_layout_get_lines_readonly", &Pango_layout_get_lines_readonly},
	{"pango_layout_get_log_attrs", &Pango_layout_get_log_attrs},
	{"pango_layout_get_log_attrs_readonly", &Pango_layout_get_log_attrs_readonly},
	{"pango_layout_get_pixel_extents", &Pango_layout_get_pixel_extents},
	{"pango_layout_get_pixel_size", &Pango_layout_get_pixel_size},
	{"pango_layout_get_single_paragraph_mode", &Pango_layout_get_single_paragraph_mode},
	{"pango_layout_get_size", &Pango_layout_get_size},
	{"pango_layout_get_spacing", &Pango_layout_get_spacing},
	{"pango_layout_get_tabs", &Pango_layout_get_tabs},
	{"pango_layout_get_text", &Pango_layout_get_text},
	{"pango_layout_get_type", &Pango_layout_get_type},
	{"pango_layout_get_unknown_glyphs_count", &Pango_layout_get_unknown_glyphs_count},
	{"pango_layout_get_width", &Pango_layout_get_width},
	{"pango_layout_get_wrap", &Pango_layout_get_wrap},
	{"pango_layout_index_to_line_x", &Pango_layout_index_to_line_x},
	{"pango_layout_index_to_pos", &Pango_layout_index_to_pos},
	{"pango_layout_is_ellipsized", &Pango_layout_is_ellipsized},
	{"pango_layout_is_wrapped", &Pango_layout_is_wrapped},
	{"pango_layout_iter_at_last_line", &Pango_layout_iter_at_last_line},
	{"pango_layout_iter_copy", &Pango_layout_iter_copy},
	{"pango_layout_iter_free", &Pango_layout_iter_free},
	{"pango_layout_iter_get_baseline", &Pango_layout_iter_get_baseline},
	{"pango_layout_iter_get_char_extents", &Pango_layout_iter_get_char_extents},
	{"pango_layout_iter_get_cluster_extents", &Pango_layout_iter_get_cluster_extents},
	{"pango_layout_iter_get_index", &Pango_layout_iter_get_index},
	{"pango_layout_iter_get_layout", &Pango_layout_iter_get_layout},
	{"pango_layout_iter_get_layout_extents", &Pango_layout_iter_get_layout_extents},
	{"pango_layout_iter_get_line", &Pango_layout_iter_get_line},
	{"pango_layout_iter_get_line_extents", &Pango_layout_iter_get_line_extents},
	{"pango_layout_iter_get_line_readonly", &Pango_layout_iter_get_line_readonly},
	{"pango_layout_iter_get_line_yrange", &Pango_layout_iter_get_line_yrange},
	{"pango_layout_iter_get_run", &Pango_layout_iter_get_run},
	{"pango_layout_iter_get_run_extents", &Pango_layout_iter_get_run_extents},
	{"pango_layout_iter_get_run_readonly", &Pango_layout_iter_get_run_readonly},
	{"pango_layout_iter_get_type", &Pango_layout_iter_get_type},
	{"pango_layout_iter_next_char", &Pango_layout_iter_next_char},
	{"pango_layout_iter_next_cluster", &Pango_layout_iter_next_cluster},
	{"pango_layout_iter_next_line", &Pango_layout_iter_next_line},
	{"pango_layout_iter_next_run", &Pango_layout_iter_next_run},
	{"pango_layout_line_get_extents", &Pango_layout_line_get_extents},
	{"pango_layout_line_get_pixel_extents", &Pango_layout_line_get_pixel_extents},
	{"pango_layout_line_get_type", &Pango_layout_line_get_type},
	{"pango_layout_line_get_x_ranges", &Pango_layout_line_get_x_ranges},
	{"pango_layout_line_index_to_x", &Pango_layout_line_index_to_x},
	{"pango_layout_line_ref", &Pango_layout_line_ref},
	{"pango_layout_line_unref", &Pango_layout_line_unref},
	{"pango_layout_line_x_to_index", &Pango_layout_line_x_to_index},
	{"pango_layout_move_cursor_visually", &Pango_layout_move_cursor_visually},
	{"pango_layout_new", &Pango_layout_new},
	{"pango_layout_set_alignment", &Pango_layout_set_alignment},
	{"pango_layout_set_attributes", &Pango_layout_set_attributes},
	{"pango_layout_set_auto_dir", &Pango_layout_set_auto_dir},
	{"pango_layout_set_ellipsize", &Pango_layout_set_ellipsize},
	{"pango_layout_set_font_description", &Pango_layout_set_font_description},
	{"pango_layout_set_height", &Pango_layout_set_height},
	{"pango_layout_set_indent", &Pango_layout_set_indent},
	{"pango_layout_set_justify", &Pango_layout_set_justify},
	{"pango_layout_set_markup", &Pango_layout_set_markup},
	{"pango_layout_set_markup_with_accel", &Pango_layout_set_markup_with_accel},
	{"pango_layout_set_single_paragraph_mode", &Pango_layout_set_single_paragraph_mode},
	{"pango_layout_set_spacing", &Pango_layout_set_spacing},
	{"pango_layout_set_tabs", &Pango_layout_set_tabs},
	{"pango_layout_set_text", &Pango_layout_set_text},
	{"pango_layout_set_width", &Pango_layout_set_width},
	{"pango_layout_set_wrap", &Pango_layout_set_wrap},
	{"pango_layout_xy_to_index", &Pango_layout_xy_to_index},
	{"pango_log2vis_get_embedding_levels", &Pango_log2vis_get_embedding_levels},
	{"pango_lookup_aliases", &Pango_lookup_aliases},
	{"pango_map_get_engine", &Pango_map_get_engine},
	{"pango_map_get_engines", &Pango_map_get_engines},
	{"pango_matrix_concat", &Pango_matrix_concat},
	{"pango_matrix_copy", &Pango_matrix_copy},
	{"pango_matrix_free", &Pango_matrix_free},
	{"pango_matrix_get_font_scale_factor", &Pango_matrix_get_font_scale_factor},
	{"pango_matrix_get_type", &Pango_matrix_get_type},
	{"pango_matrix_rotate", &Pango_matrix_rotate},
	{"pango_matrix_scale", &Pango_matrix_scale},
	{"pango_matrix_transform_distance", &Pango_matrix_transform_distance},
	{"pango_matrix_transform_pixel_rectangle", &Pango_matrix_transform_pixel_rectangle},
	{"pango_matrix_transform_point", &Pango_matrix_transform_point},
	{"pango_matrix_transform_rectangle", &Pango_matrix_transform_rectangle},
	{"pango_matrix_translate", &Pango_matrix_translate},
	{"pango_module_register", &Pango_module_register},
	{"pango_parse_enum", &Pango_parse_enum},
	{"pango_parse_markup", &Pango_parse_markup},
	{"pango_parse_stretch", &Pango_parse_stretch},
	{"pango_parse_style", &Pango_parse_style},
	{"pango_parse_variant", &Pango_parse_variant},
	{"pango_parse_weight", &Pango_parse_weight},
	{"pango_quantize_line_geometry", &Pango_quantize_line_geometry},
	{"pango_read_line", &Pango_read_line},
	{"pango_render_part_get_type", &Pango_render_part_get_type},
	{"pango_renderer_activate", &Pango_renderer_activate},
	{"pango_renderer_deactivate", &Pango_renderer_deactivate},
	{"pango_renderer_draw_error_underline", &Pango_renderer_draw_error_underline},
	{"pango_renderer_draw_glyph", &Pango_renderer_draw_glyph},
	{"pango_renderer_draw_glyph_item", &Pango_renderer_draw_glyph_item},
	{"pango_renderer_draw_glyphs", &Pango_renderer_draw_glyphs},
	{"pango_renderer_draw_layout", &Pango_renderer_draw_layout},
	{"pango_renderer_draw_layout_line", &Pango_renderer_draw_layout_line},
	{"pango_renderer_draw_rectangle", &Pango_renderer_draw_rectangle},
	{"pango_renderer_draw_trapezoid", &Pango_renderer_draw_trapezoid},
	{"pango_renderer_get_color", &Pango_renderer_get_color},
	{"pango_renderer_get_layout", &Pango_renderer_get_layout},
	{"pango_renderer_get_layout_line", &Pango_renderer_get_layout_line},
	{"pango_renderer_get_matrix", &Pango_renderer_get_matrix},
	{"pango_renderer_get_type", &Pango_renderer_get_type},
	{"pango_renderer_part_changed", &Pango_renderer_part_changed},
	{"pango_renderer_set_color", &Pango_renderer_set_color},
	{"pango_renderer_set_matrix", &Pango_renderer_set_matrix},
	{"pango_reorder_items", &Pango_reorder_items},
	{"pango_scan_int", &Pango_scan_int},
	{"pango_scan_string", &Pango_scan_string},
	{"pango_scan_word", &Pango_scan_word},
	{"pango_script_for_unichar", &Pango_script_for_unichar},
	{"pango_script_get_sample_language", &Pango_script_get_sample_language},
	{"pango_script_get_type", &Pango_script_get_type},
	{"pango_script_iter_free", &Pango_script_iter_free},
	{"pango_script_iter_get_range", &Pango_script_iter_get_range},
	{"pango_script_iter_new", &Pango_script_iter_new},
	{"pango_script_iter_next", &Pango_script_iter_next},
	{"pango_shape", &Pango_shape},
	{"pango_skip_space", &Pango_skip_space},
	{"pango_split_file_list", &Pango_split_file_list},
	{"pango_stretch_get_type", &Pango_stretch_get_type},
	{"pango_style_get_type", &Pango_style_get_type},
	{"pango_tab_align_get_type", &Pango_tab_align_get_type},
	{"pango_tab_array_copy", &Pango_tab_array_copy},
	{"pango_tab_array_free", &Pango_tab_array_free},
	{"pango_tab_array_get_positions_in_pixels", &Pango_tab_array_get_positions_in_pixels},
	{"pango_tab_array_get_size", &Pango_tab_array_get_size},
	{"pango_tab_array_get_tab", &Pango_tab_array_get_tab},
	{"pango_tab_array_get_tabs", &Pango_tab_array_get_tabs},
	{"pango_tab_array_get_type", &Pango_tab_array_get_type},
	{"pango_tab_array_new", &Pango_tab_array_new},
	{"pango_tab_array_new_with_positions", &Pango_tab_array_new_with_positions},
	{"pango_tab_array_resize", &Pango_tab_array_resize},
	{"pango_tab_array_set_tab", &Pango_tab_array_set_tab},
	{"pango_trim_string", &Pango_trim_string},
	{"pango_underline_get_type", &Pango_underline_get_type},
	{"pango_unichar_direction", &Pango_unichar_direction},
	{"pango_units_from_double", &Pango_units_from_double},
	{"pango_units_to_double", &Pango_units_to_double},
	{"pango_variant_get_type", &Pango_variant_get_type},
	{"pango_version", &Pango_version},
	{"pango_version_check", &Pango_version_check},
	{"pango_version_string", &Pango_version_string},
	{"pango_weight_get_type", &Pango_weight_get_type},
	{"pango_wrap_mode_get_type", &Pango_wrap_mode_get_type},
}

var dllCairo = "libpangocairo-1.0-0.dll"

var apiListCairo = Apis{
	{"pango_cairo_context_get_font_options", &Pango_cairo_context_get_font_options},
	{"pango_cairo_context_get_resolution", &Pango_cairo_context_get_resolution},
	{"pango_cairo_context_get_shape_renderer", &Pango_cairo_context_get_shape_renderer},
	{"pango_cairo_context_set_font_options", &Pango_cairo_context_set_font_options},
	{"pango_cairo_context_set_resolution", &Pango_cairo_context_set_resolution},
	{"pango_cairo_context_set_shape_renderer", &Pango_cairo_context_set_shape_renderer},
	{"pango_cairo_create_context", &Pango_cairo_create_context},
	{"pango_cairo_create_layout", &Pango_cairo_create_layout},
	{"pango_cairo_error_underline_path", &Pango_cairo_error_underline_path},
	{"pango_cairo_font_get_scaled_font", &Pango_cairo_font_get_scaled_font},
	{"pango_cairo_font_get_type", &Pango_cairo_font_get_type},
	{"pango_cairo_font_map_create_context", &Pango_cairo_font_map_create_context},
	{"pango_cairo_font_map_get_default", &Pango_cairo_font_map_get_default},
	{"pango_cairo_font_map_get_font_type", &Pango_cairo_font_map_get_font_type},
	{"pango_cairo_font_map_get_resolution", &Pango_cairo_font_map_get_resolution},
	{"pango_cairo_font_map_get_type", &Pango_cairo_font_map_get_type},
	{"pango_cairo_font_map_new", &Pango_cairo_font_map_new},
	{"pango_cairo_font_map_new_for_font_type", &Pango_cairo_font_map_new_for_font_type},
	{"pango_cairo_font_map_set_default", &Pango_cairo_font_map_set_default},
	{"pango_cairo_font_map_set_resolution", &Pango_cairo_font_map_set_resolution},
	{"pango_cairo_glyph_string_path", &Pango_cairo_glyph_string_path},
	{"pango_cairo_layout_line_path", &Pango_cairo_layout_line_path},
	{"pango_cairo_layout_path", &Pango_cairo_layout_path},
	// Undocumented {"pango_cairo_renderer_get_type", &Pango_cairo_renderer_get_type},
	{"pango_cairo_show_error_underline", &Pango_cairo_show_error_underline},
	{"pango_cairo_show_glyph_item", &Pango_cairo_show_glyph_item},
	{"pango_cairo_show_glyph_string", &Pango_cairo_show_glyph_string},
	{"pango_cairo_show_layout", &Pango_cairo_show_layout},
	{"pango_cairo_show_layout_line", &Pango_cairo_show_layout_line},
	{"pango_cairo_update_context", &Pango_cairo_update_context},
	{"pango_cairo_update_layout", &Pango_cairo_update_layout},
}
