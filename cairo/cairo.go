package cairo

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllApis(dllGobject, false, apiListGobject)
	outside.AddDllApis(dllScript, false, apiListScript)
}

type (
	//TODO(t):Fix
	FILE     uintptr
	HDC      uint32
	HFONT    uint32
	LOGFONTW struct{}
)

type Cairo_pdf_version_t Enum

const (
	CAIRO_PDF_VERSION_1_4 Cairo_pdf_version_t = iota
	CAIRO_PDF_VERSION_1_5
)

type Cairo_ps_level_t Enum

const (
	CAIRO_PS_LEVEL_2 Cairo_ps_level_t = iota
	CAIRO_PS_LEVEL_3
)

type Cairo_svg_version_t Enum

const (
	CAIRO_SVG_VERSION_1_1 Cairo_svg_version_t = iota
	CAIRO_SVG_VERSION_1_2
)

var (
	Cairo_version func() int

	Cairo_version_string func() *Char

	Cairo_create func(
		target *Cairo_surface_t) *Cairo_t

	Cairo_reference func(
		cr *Cairo_t) *Cairo_t

	Cairo_destroy func(
		cr *Cairo_t)

	Cairo_get_reference_count func(
		cr *Cairo_t) Unsigned_int

	Cairo_get_user_data func(
		cr *Cairo_t,
		key *Cairo_user_data_key_t) *Void

	Cairo_set_user_data func(
		cr *Cairo_t,
		key *Cairo_user_data_key_t,
		user_data *Void,
		destroy Cairo_destroy_func_t) Cairo_status_t

	Cairo_save func(
		cr *Cairo_t)

	Cairo_restore func(
		cr *Cairo_t)

	Cairo_push_group func(
		cr *Cairo_t)

	Cairo_push_group_with_content func(
		cr *Cairo_t,
		content Cairo_content_t)

	Cairo_pop_group func(
		cr *Cairo_t) *Cairo_pattern_t

	Cairo_pop_group_to_source func(
		cr *Cairo_t)

	Cairo_set_operator func(
		cr *Cairo_t,
		op Cairo_operator_t)

	Cairo_set_source func(
		cr *Cairo_t,
		source *Cairo_pattern_t)

	Cairo_set_source_rgb func(
		cr *Cairo_t,
		red Double,
		green Double,
		blue Double)

	Cairo_set_source_rgba func(
		cr *Cairo_t,
		red Double,
		green Double,
		blue Double,
		alpha Double)

	Cairo_set_source_surface func(
		cr *Cairo_t,
		surface *Cairo_surface_t,
		x Double,
		y Double)

	Cairo_set_tolerance func(
		cr *Cairo_t,
		tolerance Double)

	Cairo_set_antialias func(
		cr *Cairo_t,
		antialias Cairo_antialias_t)

	Cairo_set_fill_rule func(
		cr *Cairo_t,
		fill_rule Cairo_fill_rule_t)

	Cairo_set_line_width func(
		cr *Cairo_t,
		width Double)

	Cairo_set_line_cap func(
		cr *Cairo_t,
		line_cap Cairo_line_cap_t)

	Cairo_set_line_join func(
		cr *Cairo_t,
		line_join Cairo_line_join_t)

	Cairo_set_dash func(
		cr *Cairo_t,
		dashes *Double,
		num_dashes int,
		offset Double)

	Cairo_set_miter_limit func(
		cr *Cairo_t,
		limit Double)

	Cairo_translate func(
		cr *Cairo_t,
		tx Double,
		ty Double)

	Cairo_scale func(
		cr *Cairo_t,
		sx Double,
		sy Double)

	Cairo_rotate func(
		cr *Cairo_t,
		angle Double)

	Cairo_transform func(
		cr *Cairo_t,
		matrix *Cairo_matrix_t)

	Cairo_set_matrix func(
		cr *Cairo_t,
		matrix *Cairo_matrix_t)

	Cairo_identity_matrix func(
		cr *Cairo_t)

	Cairo_user_to_device func(
		cr *Cairo_t,
		x *Double,
		y *Double)

	Cairo_user_to_device_distance func(
		cr *Cairo_t,
		dx *Double,
		dy *Double)

	Cairo_device_to_user func(
		cr *Cairo_t,
		x *Double,
		y *Double)

	Cairo_device_to_user_distance func(
		cr *Cairo_t,
		dx *Double,
		dy *Double)

	Cairo_new_path func(
		cr *Cairo_t)

	Cairo_move_to func(
		cr *Cairo_t,
		x Double,
		y Double)

	Cairo_new_sub_path func(
		cr *Cairo_t)

	Cairo_line_to func(
		cr *Cairo_t,
		x Double,
		y Double)

	Cairo_curve_to func(
		cr *Cairo_t,
		x1 Double,
		y1 Double,
		x2 Double,
		y2 Double,
		x3 Double,
		y3 Double)

	Cairo_arc func(
		cr *Cairo_t,
		xc Double,
		yc Double,
		radius Double,
		angle1 Double,
		angle2 Double)

	Cairo_arc_negative func(
		cr *Cairo_t,
		xc Double,
		yc Double,
		radius Double,
		angle1 Double,
		angle2 Double)

	Cairo_rel_move_to func(
		cr *Cairo_t,
		dx Double,
		dy Double)

	Cairo_rel_line_to func(
		cr *Cairo_t,
		dx Double,
		dy Double)

	Cairo_rel_curve_to func(
		cr *Cairo_t,
		dx1 Double,
		dy1 Double,
		dx2 Double,
		dy2 Double,
		dx3 Double,
		dy3 Double)

	Cairo_rectangle func(
		cr *Cairo_t,
		x Double,
		y Double,
		width Double,
		height Double)

	Cairo_close_path func(
		cr *Cairo_t)

	Cairo_path_extents func(
		cr *Cairo_t,
		x1 *Double,
		y1 *Double,
		x2 *Double,
		y2 *Double)

	Cairo_paint func(
		cr *Cairo_t)

	Cairo_paint_with_alpha func(
		cr *Cairo_t,
		alpha Double)

	Cairo_mask func(
		cr *Cairo_t,
		pattern *Cairo_pattern_t)

	Cairo_mask_surface func(
		cr *Cairo_t,
		surface *Cairo_surface_t,
		surface_x Double,
		surface_y Double)

	Cairo_stroke func(
		cr *Cairo_t)

	Cairo_stroke_preserve func(
		cr *Cairo_t)

	Cairo_fill func(
		cr *Cairo_t)

	Cairo_fill_preserve func(
		cr *Cairo_t)

	Cairo_copy_page func(
		cr *Cairo_t)

	Cairo_show_page func(
		cr *Cairo_t)

	Cairo_in_stroke func(
		cr *Cairo_t,
		x Double,
		y Double) Cairo_bool_t

	Cairo_in_fill func(
		cr *Cairo_t,
		x Double,
		y Double) Cairo_bool_t

	Cairo_in_clip func(
		cr *Cairo_t,
		x Double,
		y Double) Cairo_bool_t

	Cairo_stroke_extents func(
		cr *Cairo_t,
		x1 *Double,
		y1 *Double,
		x2 *Double,
		y2 *Double)

	Cairo_fill_extents func(
		cr *Cairo_t,
		x1 *Double,
		y1 *Double,
		x2 *Double,
		y2 *Double)

	Cairo_reset_clip func(
		cr *Cairo_t)

	Cairo_clip func(
		cr *Cairo_t)

	Cairo_clip_preserve func(
		cr *Cairo_t)

	Cairo_clip_extents func(
		cr *Cairo_t,
		x1 *Double,
		y1 *Double,
		x2 *Double,
		y2 *Double)

	Cairo_copy_clip_rectangle_list func(
		cr *Cairo_t) *Cairo_rectangle_list_t

	Cairo_rectangle_list_destroy func(
		rectangle_list *Cairo_rectangle_list_t)

	Cairo_glyph_allocate func(
		num_glyphs int) *Cairo_glyph_t

	Cairo_glyph_free func(
		glyphs *Cairo_glyph_t)

	Cairo_text_cluster_allocate func(
		num_clusters int) *Cairo_text_cluster_t

	Cairo_text_cluster_free func(
		clusters *Cairo_text_cluster_t)

	Cairo_font_options_create func() *Cairo_font_options_t

	Cairo_font_options_copy func(
		original *Cairo_font_options_t) *Cairo_font_options_t

	Cairo_font_options_destroy func(
		options *Cairo_font_options_t)

	Cairo_font_options_status func(
		options *Cairo_font_options_t) Cairo_status_t

	Cairo_font_options_merge func(
		options *Cairo_font_options_t,
		other *Cairo_font_options_t)

	Cairo_font_options_equal func(
		options *Cairo_font_options_t,
		other *Cairo_font_options_t) Cairo_bool_t

	Cairo_font_options_hash func(
		options *Cairo_font_options_t) Unsigned_long

	Cairo_font_options_set_antialias func(
		options *Cairo_font_options_t,
		antialias Cairo_antialias_t)

	Cairo_font_options_get_antialias func(
		options *Cairo_font_options_t) Cairo_antialias_t

	Cairo_font_options_set_subpixel_order func(
		options *Cairo_font_options_t,
		subpixel_order Cairo_subpixel_order_t)

	Cairo_font_options_get_subpixel_order func(
		options *Cairo_font_options_t) Cairo_subpixel_order_t

	Cairo_font_options_set_hint_style func(
		options *Cairo_font_options_t,
		hint_style Cairo_hint_style_t)

	Cairo_font_options_get_hint_style func(
		options *Cairo_font_options_t) Cairo_hint_style_t

	Cairo_font_options_set_hint_metrics func(
		options *Cairo_font_options_t,
		hint_metrics Cairo_hint_metrics_t)

	Cairo_font_options_get_hint_metrics func(
		options *Cairo_font_options_t) Cairo_hint_metrics_t

	Cairo_select_font_face func(
		cr *Cairo_t,
		family *Char,
		slant Cairo_font_slant_t,
		weight Cairo_font_weight_t)

	Cairo_set_font_size func(
		cr *Cairo_t,
		size Double)

	Cairo_set_font_matrix func(
		cr *Cairo_t,
		matrix *Cairo_matrix_t)

	Cairo_get_font_matrix func(
		cr *Cairo_t,
		matrix *Cairo_matrix_t)

	Cairo_set_font_options func(
		cr *Cairo_t,
		options *Cairo_font_options_t)

	Cairo_get_font_options func(
		cr *Cairo_t,
		options *Cairo_font_options_t)

	Cairo_set_font_face func(
		cr *Cairo_t,
		font_face *Cairo_font_face_t)

	Cairo_get_font_face func(
		cr *Cairo_t) *Cairo_font_face_t

	Cairo_set_scaled_font func(
		cr *Cairo_t,
		scaled_font *Cairo_scaled_font_t)

	Cairo_get_scaled_font func(
		cr *Cairo_t) *Cairo_scaled_font_t

	Cairo_show_text func(
		cr *Cairo_t,
		utf8 *Char)

	Cairo_show_glyphs func(
		cr *Cairo_t,
		glyphs *Cairo_glyph_t,
		num_glyphs int)

	Cairo_show_text_glyphs func(
		cr *Cairo_t,
		utf8 *Char,
		utf8_len int,
		glyphs *Cairo_glyph_t,
		num_glyphs int,
		clusters *Cairo_text_cluster_t,
		num_clusters int,
		cluster_flags Cairo_text_cluster_flags_t)

	Cairo_text_path func(
		cr *Cairo_t,
		utf8 *Char)

	Cairo_glyph_path func(
		cr *Cairo_t,
		glyphs *Cairo_glyph_t,
		num_glyphs int)

	Cairo_text_extents func(
		cr *Cairo_t,
		utf8 *Char,
		extents *Cairo_text_extents_t)

	Cairo_glyph_extents func(
		cr *Cairo_t,
		glyphs *Cairo_glyph_t,
		num_glyphs int,
		extents *Cairo_text_extents_t)

	Cairo_font_extents func(
		cr *Cairo_t,
		extents *Cairo_font_extents_t)

	Cairo_font_face_reference func(
		font_face *Cairo_font_face_t) *Cairo_font_face_t

	Cairo_font_face_destroy func(
		font_face *Cairo_font_face_t)

	Cairo_font_face_get_reference_count func(
		font_face *Cairo_font_face_t) Unsigned_int

	Cairo_font_face_status func(
		font_face *Cairo_font_face_t) Cairo_status_t

	Cairo_font_face_get_type func(
		font_face *Cairo_font_face_t) Cairo_font_type_t

	Cairo_font_face_get_user_data func(
		font_face *Cairo_font_face_t,
		key *Cairo_user_data_key_t) *Void

	Cairo_font_face_set_user_data func(
		font_face *Cairo_font_face_t,
		key *Cairo_user_data_key_t,
		user_data *Void,
		destroy Cairo_destroy_func_t) Cairo_status_t

	Cairo_scaled_font_create func(
		font_face *Cairo_font_face_t,
		font_matrix *Cairo_matrix_t,
		ctm *Cairo_matrix_t,
		options *Cairo_font_options_t) *Cairo_scaled_font_t

	Cairo_scaled_font_reference func(
		scaled_font *Cairo_scaled_font_t) *Cairo_scaled_font_t

	Cairo_scaled_font_destroy func(
		scaled_font *Cairo_scaled_font_t)

	Cairo_scaled_font_get_reference_count func(
		scaled_font *Cairo_scaled_font_t) Unsigned_int

	Cairo_scaled_font_status func(
		scaled_font *Cairo_scaled_font_t) Cairo_status_t

	Cairo_scaled_font_get_type func(
		scaled_font *Cairo_scaled_font_t) Cairo_font_type_t

	Cairo_scaled_font_get_user_data func(
		scaled_font *Cairo_scaled_font_t,
		key *Cairo_user_data_key_t) *Void

	Cairo_scaled_font_set_user_data func(
		scaled_font *Cairo_scaled_font_t,
		key *Cairo_user_data_key_t,
		user_data *Void,
		destroy Cairo_destroy_func_t) Cairo_status_t

	Cairo_scaled_font_extents func(
		scaled_font *Cairo_scaled_font_t,
		extents *Cairo_font_extents_t)

	Cairo_scaled_font_text_extents func(
		scaled_font *Cairo_scaled_font_t,
		utf8 *Char,
		extents *Cairo_text_extents_t)

	Cairo_scaled_font_glyph_extents func(
		scaled_font *Cairo_scaled_font_t,
		glyphs *Cairo_glyph_t,
		num_glyphs int,
		extents *Cairo_text_extents_t)

	Cairo_scaled_font_text_to_glyphs func(
		scaled_font *Cairo_scaled_font_t,
		x Double,
		y Double,
		utf8 *Char,
		utf8_len int,
		glyphs **Cairo_glyph_t,
		num_glyphs *int,
		clusters **Cairo_text_cluster_t,
		num_clusters *int,
		cluster_flags *Cairo_text_cluster_flags_t) Cairo_status_t

	Cairo_scaled_font_get_font_face func(
		scaled_font *Cairo_scaled_font_t) *Cairo_font_face_t

	Cairo_scaled_font_get_font_matrix func(
		scaled_font *Cairo_scaled_font_t,
		font_matrix *Cairo_matrix_t)

	Cairo_scaled_font_get_ctm func(
		scaled_font *Cairo_scaled_font_t,
		ctm *Cairo_matrix_t)

	Cairo_scaled_font_get_scale_matrix func(
		scaled_font *Cairo_scaled_font_t,
		scale_matrix *Cairo_matrix_t)

	Cairo_scaled_font_get_font_options func(
		scaled_font *Cairo_scaled_font_t,
		options *Cairo_font_options_t)

	Cairo_toy_font_face_create func(
		family *Char,
		slant Cairo_font_slant_t,
		weight Cairo_font_weight_t) *Cairo_font_face_t

	Cairo_toy_font_face_get_family func(
		font_face *Cairo_font_face_t) *Char

	Cairo_toy_font_face_get_slant func(
		font_face *Cairo_font_face_t) Cairo_font_slant_t

	Cairo_toy_font_face_get_weight func(
		font_face *Cairo_font_face_t) Cairo_font_weight_t

	Cairo_user_font_face_create func() *Cairo_font_face_t

	Cairo_user_font_face_set_init_func func(
		font_face *Cairo_font_face_t,
		init_func Cairo_user_scaled_font_init_func_t)

	Cairo_user_font_face_set_render_glyph_func func(
		font_face *Cairo_font_face_t,
		render_glyph_func Cairo_user_scaled_font_render_glyph_func_t)

	Cairo_user_font_face_set_text_to_glyphs_func func(
		font_face *Cairo_font_face_t,
		text_to_glyphs_func Cairo_user_scaled_font_text_to_glyphs_func_t)

	Cairo_user_font_face_set_unicode_to_glyph_func func(
		font_face *Cairo_font_face_t,
		unicode_to_glyph_func Cairo_user_scaled_font_unicode_to_glyph_func_t)

	Cairo_user_font_face_get_init_func func(
		font_face *Cairo_font_face_t) Cairo_user_scaled_font_init_func_t

	Cairo_user_font_face_get_render_glyph_func func(
		font_face *Cairo_font_face_t) Cairo_user_scaled_font_render_glyph_func_t

	Cairo_user_font_face_get_text_to_glyphs_func func(
		font_face *Cairo_font_face_t) Cairo_user_scaled_font_text_to_glyphs_func_t

	Cairo_user_font_face_get_unicode_to_glyph_func func(
		font_face *Cairo_font_face_t) Cairo_user_scaled_font_unicode_to_glyph_func_t

	Cairo_get_operator func(
		cr *Cairo_t) Cairo_operator_t

	Cairo_get_source func(
		cr *Cairo_t) *Cairo_pattern_t

	Cairo_get_tolerance func(
		cr *Cairo_t) Double

	Cairo_get_antialias func(
		cr *Cairo_t) Cairo_antialias_t

	Cairo_has_current_point func(
		cr *Cairo_t) Cairo_bool_t

	Cairo_get_current_point func(
		cr *Cairo_t,
		x *Double,
		y *Double)

	Cairo_get_fill_rule func(
		cr *Cairo_t) Cairo_fill_rule_t

	Cairo_get_line_width func(
		cr *Cairo_t) Double

	Cairo_get_line_cap func(
		cr *Cairo_t) Cairo_line_cap_t

	Cairo_get_line_join func(
		cr *Cairo_t) Cairo_line_join_t

	Cairo_get_miter_limit func(
		cr *Cairo_t) Double

	Cairo_get_dash_count func(
		cr *Cairo_t) int

	Cairo_get_dash func(
		cr *Cairo_t,
		dashes *Double,
		offset *Double)

	Cairo_get_matrix func(
		cr *Cairo_t,
		matrix *Cairo_matrix_t)

	Cairo_get_target func(
		cr *Cairo_t) *Cairo_surface_t

	Cairo_get_group_target func(
		cr *Cairo_t) *Cairo_surface_t

	Cairo_copy_path func(
		cr *Cairo_t) *Cairo_path_t

	Cairo_copy_path_flat func(
		cr *Cairo_t) *Cairo_path_t

	Cairo_append_path func(
		cr *Cairo_t,
		path *Cairo_path_t)

	Cairo_path_destroy func(
		path *Cairo_path_t)

	Cairo_status func(
		cr *Cairo_t) Cairo_status_t

	Cairo_status_to_string func(
		status Cairo_status_t) *Char

	Cairo_device_reference func(
		device *Cairo_device_t) *Cairo_device_t

	Cairo_device_get_type func(
		device *Cairo_device_t) Cairo_device_type_t

	Cairo_device_status func(
		device *Cairo_device_t) Cairo_status_t

	Cairo_device_acquire func(
		device *Cairo_device_t) Cairo_status_t

	Cairo_device_release func(
		device *Cairo_device_t)

	Cairo_device_flush func(
		device *Cairo_device_t)

	Cairo_device_finish func(
		device *Cairo_device_t)

	Cairo_device_destroy func(
		device *Cairo_device_t)

	Cairo_device_get_reference_count func(
		device *Cairo_device_t) Unsigned_int

	Cairo_device_get_user_data func(
		device *Cairo_device_t,
		key *Cairo_user_data_key_t) *Void

	Cairo_device_set_user_data func(
		device *Cairo_device_t,
		key *Cairo_user_data_key_t,
		user_data *Void,
		destroy Cairo_destroy_func_t) Cairo_status_t

	Cairo_surface_create_similar func(
		other *Cairo_surface_t,
		content Cairo_content_t,
		width int,
		height int) *Cairo_surface_t

	Cairo_surface_create_for_rectangle func(
		target *Cairo_surface_t,
		x Double,
		y Double,
		width Double,
		height Double) *Cairo_surface_t

	Cairo_surface_reference func(
		surface *Cairo_surface_t) *Cairo_surface_t

	Cairo_surface_finish func(
		surface *Cairo_surface_t)

	Cairo_surface_destroy func(
		surface *Cairo_surface_t)

	Cairo_surface_get_device func(
		surface *Cairo_surface_t) *Cairo_device_t

	Cairo_surface_get_reference_count func(
		surface *Cairo_surface_t) Unsigned_int

	Cairo_surface_status func(
		surface *Cairo_surface_t) Cairo_status_t

	Cairo_surface_get_type func(
		surface *Cairo_surface_t) Cairo_surface_type_t

	Cairo_surface_get_content func(
		surface *Cairo_surface_t) Cairo_content_t

	Cairo_surface_write_to_png func(
		surface *Cairo_surface_t,
		filename *Char) Cairo_status_t

	Cairo_surface_write_to_png_stream func(
		surface *Cairo_surface_t,
		write_func Cairo_write_func_t,
		closure *Void) Cairo_status_t

	Cairo_surface_get_user_data func(
		surface *Cairo_surface_t,
		key *Cairo_user_data_key_t) *Void

	Cairo_surface_set_user_data func(
		surface *Cairo_surface_t,
		key *Cairo_user_data_key_t,
		user_data *Void,
		destroy Cairo_destroy_func_t) Cairo_status_t

	Cairo_surface_get_mime_data func(
		surface *Cairo_surface_t,
		mime_type *Char,
		data **Unsigned_char,
		length *Unsigned_long)

	Cairo_surface_set_mime_data func(
		surface *Cairo_surface_t,
		mime_type *Char,
		data *Unsigned_char,
		length Unsigned_long,
		destroy Cairo_destroy_func_t,
		closure *Void) Cairo_status_t

	Cairo_surface_get_font_options func(
		surface *Cairo_surface_t,
		options *Cairo_font_options_t)

	Cairo_surface_flush func(
		surface *Cairo_surface_t)

	Cairo_surface_mark_dirty func(
		surface *Cairo_surface_t)

	Cairo_surface_mark_dirty_rectangle func(
		surface *Cairo_surface_t,
		x int,
		y int,
		width int,
		height int)

	Cairo_surface_set_device_offset func(
		surface *Cairo_surface_t,
		x_offset Double,
		y_offset Double)

	Cairo_surface_get_device_offset func(
		surface *Cairo_surface_t,
		x_offset *Double,
		y_offset *Double)

	Cairo_surface_set_fallback_resolution func(
		surface *Cairo_surface_t,
		x_pixels_per_inch Double,
		y_pixels_per_inch Double)

	Cairo_surface_get_fallback_resolution func(
		surface *Cairo_surface_t,
		x_pixels_per_inch *Double,
		y_pixels_per_inch *Double)

	Cairo_surface_copy_page func(
		surface *Cairo_surface_t)

	Cairo_surface_show_page func(
		surface *Cairo_surface_t)

	Cairo_surface_has_show_text_glyphs func(
		surface *Cairo_surface_t) Cairo_bool_t

	Cairo_image_surface_create func(
		format Cairo_format_t,
		width int,
		height int) *Cairo_surface_t

	Cairo_format_stride_for_width func(
		format Cairo_format_t,
		width int) int

	Cairo_image_surface_create_for_data func(
		data *Unsigned_char,
		format Cairo_format_t,
		width int,
		height int,
		stride int) *Cairo_surface_t

	Cairo_image_surface_get_data func(
		surface *Cairo_surface_t) *Unsigned_char

	Cairo_image_surface_get_format func(
		surface *Cairo_surface_t) Cairo_format_t

	Cairo_image_surface_get_width func(
		surface *Cairo_surface_t) int

	Cairo_image_surface_get_height func(
		surface *Cairo_surface_t) int

	Cairo_image_surface_get_stride func(
		surface *Cairo_surface_t) int

	Cairo_image_surface_create_from_png func(
		filename *Char) *Cairo_surface_t

	Cairo_image_surface_create_from_png_stream func(
		read_func Cairo_read_func_t,
		closure *Void) *Cairo_surface_t

	Cairo_recording_surface_create func(
		content Cairo_content_t,
		extents *Cairo_rectangle_t) *Cairo_surface_t

	Cairo_recording_surface_ink_extents func(
		surface *Cairo_surface_t,
		x0 *Double,
		y0 *Double,
		width *Double,
		height *Double)

	Cairo_pattern_create_rgb func(
		red Double,
		green Double,
		blue Double) *Cairo_pattern_t

	Cairo_pattern_create_rgba func(
		red Double,
		green Double,
		blue Double,
		alpha Double) *Cairo_pattern_t

	Cairo_pattern_create_for_surface func(
		surface *Cairo_surface_t) *Cairo_pattern_t

	Cairo_pattern_create_linear func(
		x0 Double,
		y0 Double,
		x1 Double,
		y1 Double) *Cairo_pattern_t

	Cairo_pattern_create_radial func(
		cx0 Double,
		cy0 Double,
		radius0 Double,
		cx1 Double,
		cy1 Double,
		radius1 Double) *Cairo_pattern_t

	Cairo_pattern_reference func(
		pattern *Cairo_pattern_t) *Cairo_pattern_t

	Cairo_pattern_destroy func(
		pattern *Cairo_pattern_t)

	Cairo_pattern_get_reference_count func(
		pattern *Cairo_pattern_t) Unsigned_int

	Cairo_pattern_status func(
		pattern *Cairo_pattern_t) Cairo_status_t

	Cairo_pattern_get_user_data func(
		pattern *Cairo_pattern_t,
		key *Cairo_user_data_key_t) *Void

	Cairo_pattern_set_user_data func(
		pattern *Cairo_pattern_t,
		key *Cairo_user_data_key_t,
		user_data *Void,
		destroy Cairo_destroy_func_t) Cairo_status_t

	Cairo_pattern_get_type func(
		pattern *Cairo_pattern_t) Cairo_pattern_type_t

	Cairo_pattern_add_color_stop_rgb func(
		pattern *Cairo_pattern_t,
		offset Double,
		red Double,
		green Double,
		blue Double)

	Cairo_pattern_add_color_stop_rgba func(
		pattern *Cairo_pattern_t,
		offset Double,
		red Double,
		green Double,
		blue Double,
		alpha Double)

	Cairo_pattern_set_matrix func(
		pattern *Cairo_pattern_t,
		matrix *Cairo_matrix_t)

	Cairo_pattern_get_matrix func(
		pattern *Cairo_pattern_t,
		matrix *Cairo_matrix_t)

	Cairo_pattern_set_extend func(
		pattern *Cairo_pattern_t,
		extend Cairo_extend_t)

	Cairo_pattern_get_extend func(
		pattern *Cairo_pattern_t) Cairo_extend_t

	Cairo_pattern_set_filter func(
		pattern *Cairo_pattern_t,
		filter Cairo_filter_t)

	Cairo_pattern_get_filter func(
		pattern *Cairo_pattern_t) Cairo_filter_t

	Cairo_pattern_get_rgba func(
		pattern *Cairo_pattern_t,
		red *Double,
		green *Double,
		blue *Double,
		alpha *Double) Cairo_status_t

	Cairo_pattern_get_surface func(
		pattern *Cairo_pattern_t,
		surface **Cairo_surface_t) Cairo_status_t

	Cairo_pattern_get_color_stop_rgba func(
		pattern *Cairo_pattern_t,
		index int,
		offset *Double,
		red *Double,
		green *Double,
		blue *Double,
		alpha *Double) Cairo_status_t

	Cairo_pattern_get_color_stop_count func(
		pattern *Cairo_pattern_t,
		count *int) Cairo_status_t

	Cairo_pattern_get_linear_points func(
		pattern *Cairo_pattern_t,
		x0 *Double,
		y0 *Double,
		x1 *Double,
		y1 *Double) Cairo_status_t

	Cairo_pattern_get_radial_circles func(
		pattern *Cairo_pattern_t,
		x0 *Double,
		y0 *Double,
		r0 *Double,
		x1 *Double,
		y1 *Double,
		r1 *Double) Cairo_status_t

	Cairo_matrix_init func(
		matrix *Cairo_matrix_t,
		xx Double,
		yx Double,
		xy Double,
		yy Double,
		x0 Double,
		y0 Double)

	Cairo_matrix_init_identity func(
		matrix *Cairo_matrix_t)

	Cairo_matrix_init_translate func(
		matrix *Cairo_matrix_t,
		tx Double,
		ty Double)

	Cairo_matrix_init_scale func(
		matrix *Cairo_matrix_t,
		sx Double,
		sy Double)

	Cairo_matrix_init_rotate func(
		matrix *Cairo_matrix_t,
		radians Double)

	Cairo_matrix_translate func(
		matrix *Cairo_matrix_t,
		tx Double,
		ty Double)

	Cairo_matrix_scale func(
		matrix *Cairo_matrix_t,
		sx Double,
		sy Double)

	Cairo_matrix_rotate func(
		matrix *Cairo_matrix_t,
		radians Double)

	Cairo_matrix_invert func(
		matrix *Cairo_matrix_t) Cairo_status_t

	Cairo_matrix_multiply func(
		result *Cairo_matrix_t,
		a *Cairo_matrix_t,
		b *Cairo_matrix_t)

	Cairo_matrix_transform_distance func(
		matrix *Cairo_matrix_t,
		dx *Double,
		dy *Double)

	Cairo_matrix_transform_point func(
		matrix *Cairo_matrix_t,
		x *Double,
		y *Double)

	Cairo_region_create func() *Cairo_region_t

	Cairo_region_create_rectangle func(
		rectangle *Cairo_rectangle_int_t) *Cairo_region_t

	Cairo_region_create_rectangles func(
		rects *Cairo_rectangle_int_t,
		count int) *Cairo_region_t

	Cairo_region_copy func(
		original *Cairo_region_t) *Cairo_region_t

	Cairo_region_reference func(
		region *Cairo_region_t) *Cairo_region_t

	Cairo_region_destroy func(
		region *Cairo_region_t)

	Cairo_region_equal func(
		a *Cairo_region_t,
		b *Cairo_region_t) Cairo_bool_t

	Cairo_region_status func(
		region *Cairo_region_t) Cairo_status_t

	Cairo_region_get_extents func(
		region *Cairo_region_t,
		extents *Cairo_rectangle_int_t)

	Cairo_region_num_rectangles func(
		region *Cairo_region_t) int

	Cairo_region_get_rectangle func(
		region *Cairo_region_t,
		nth int,
		rectangle *Cairo_rectangle_int_t)

	Cairo_region_is_empty func(
		region *Cairo_region_t) Cairo_bool_t

	Cairo_region_contains_rectangle func(
		region *Cairo_region_t,
		rectangle *Cairo_rectangle_int_t) Cairo_region_overlap_t

	Cairo_region_contains_point func(
		region *Cairo_region_t,
		x int,
		y int) Cairo_bool_t

	Cairo_region_translate func(
		region *Cairo_region_t,
		dx int,
		dy int)

	Cairo_region_subtract func(
		dst *Cairo_region_t,
		other *Cairo_region_t) Cairo_status_t

	Cairo_region_subtract_rectangle func(
		dst *Cairo_region_t,
		rectangle *Cairo_rectangle_int_t) Cairo_status_t

	Cairo_region_intersect func(
		dst *Cairo_region_t,
		other *Cairo_region_t) Cairo_status_t

	Cairo_region_intersect_rectangle func(
		dst *Cairo_region_t,
		rectangle *Cairo_rectangle_int_t) Cairo_status_t

	Cairo_region_union func(
		dst *Cairo_region_t,
		other *Cairo_region_t) Cairo_status_t

	Cairo_region_union_rectangle func(
		dst *Cairo_region_t,
		rectangle *Cairo_rectangle_int_t) Cairo_status_t

	Cairo_region_xor func(
		dst *Cairo_region_t,
		other *Cairo_region_t) Cairo_status_t

	Cairo_region_xor_rectangle func(
		dst *Cairo_region_t,
		rectangle *Cairo_rectangle_int_t) Cairo_status_t

	Cairo_debug_reset_static_data func()

	Cairo_gobject_context_get_type func()

	Cairo_gobject_device_get_type func()

	Cairo_gobject_pattern_get_type func()

	Cairo_gobject_surface_get_type func()

	Cairo_gobject_rectangle_get_type func()

	Cairo_gobject_scaled_font_get_type func()

	Cairo_gobject_font_face_get_type func()

	Cairo_gobject_font_options_get_type func()

	Cairo_gobject_rectangle_int_get_type func()

	Cairo_gobject_region_get_type func()

	Cairo_gobject_status_get_type func()

	Cairo_gobject_content_get_type func()

	Cairo_gobject_operator_get_type func()

	Cairo_gobject_antialias_get_type func()

	Cairo_gobject_fill_rule_get_type func()

	Cairo_gobject_line_cap_get_type func()

	Cairo_gobject_line_join_get_type func()

	Cairo_gobject_text_cluster_flags_get_type func()

	Cairo_gobject_font_slant_get_type func()

	Cairo_gobject_font_weight_get_type func()

	Cairo_gobject_subpixel_order_get_type func()

	Cairo_gobject_hint_style_get_type func()

	Cairo_gobject_hint_metrics_get_type func()

	Cairo_gobject_font_type_get_type func()

	Cairo_gobject_path_data_type_get_type func()

	Cairo_gobject_device_type_get_type func()

	Cairo_gobject_surface_type_get_type func()

	Cairo_gobject_format_get_type func()

	Cairo_gobject_pattern_type_get_type func()

	Cairo_gobject_extend_get_type func()

	Cairo_gobject_filter_get_type func()

	Cairo_gobject_region_overlap_get_type func()

	Cairo_script_interpreter_create func() *Cairo_script_interpreter_t

	Cairo_script_interpreter_install_hooks func(
		ctx *Cairo_script_interpreter_t,
		hooks *Cairo_script_interpreter_hooks_t)

	Cairo_script_interpreter_run func(
		ctx *Cairo_script_interpreter_t,
		filename *Char) Cairo_status_t

	Cairo_script_interpreter_feed_stream func(
		ctx *Cairo_script_interpreter_t,
		stream *FILE) Cairo_status_t

	Cairo_script_interpreter_feed_string func(
		ctx *Cairo_script_interpreter_t,
		line *Char,
		len int) Cairo_status_t

	Cairo_script_interpreter_get_line_number func(
		ctx *Cairo_script_interpreter_t) Unsigned_int

	Cairo_script_interpreter_reference func(
		ctx *Cairo_script_interpreter_t) *Cairo_script_interpreter_t

	Cairo_script_interpreter_finish func(
		ctx *Cairo_script_interpreter_t) Cairo_status_t

	Cairo_script_interpreter_destroy func(
		ctx *Cairo_script_interpreter_t) Cairo_status_t

	Cairo_script_interpreter_translate_stream func(
		stream *FILE,
		write_func Cairo_write_func_t,
		closure *Void) Cairo_status_t

	Cairo_ft_font_face_create_for_ft_face func(
		face FT_Face,
		load_flags int) *Cairo_font_face_t

	Cairo_ft_scaled_font_lock_face func(
		scaled_font *Cairo_scaled_font_t) FT_Face

	Cairo_ft_scaled_font_unlock_face func(
		scaled_font *Cairo_scaled_font_t)

	Cairo_ft_font_face_create_for_pattern func(
		pattern *FcPattern) *Cairo_font_face_t

	Cairo_ft_font_options_substitute func(
		options *Cairo_font_options_t,
		pattern *FcPattern)

	Cairo_pdf_surface_create func(
		filename *Char,
		width_in_points Double,
		height_in_points Double) *Cairo_surface_t

	Cairo_pdf_surface_create_for_stream func(
		write_func Cairo_write_func_t,
		closure *Void,
		width_in_points Double,
		height_in_points Double) *Cairo_surface_t

	Cairo_pdf_surface_restrict_to_version func(
		surface *Cairo_surface_t,
		version Cairo_pdf_version_t)

	Cairo_pdf_get_versions func(
		versions **Cairo_pdf_version_t,
		num_versions *int)

	Cairo_pdf_version_to_string func(
		version Cairo_pdf_version_t) *Char

	Cairo_pdf_surface_set_size func(
		surface *Cairo_surface_t,
		width_in_points Double,
		height_in_points Double)

	Cairo_ps_surface_create func(
		filename *Char,
		width_in_points Double,
		height_in_points Double) *Cairo_surface_t

	Cairo_ps_surface_create_for_stream func(
		write_func Cairo_write_func_t,
		closure *Void,
		width_in_points Double,
		height_in_points Double) *Cairo_surface_t

	Cairo_ps_surface_restrict_to_level func(
		surface *Cairo_surface_t,
		level Cairo_ps_level_t)

	Cairo_ps_get_levels func(
		levels **Cairo_ps_level_t,
		num_levels *int)

	Cairo_ps_level_to_string func(
		level Cairo_ps_level_t)

	Cairo_ps_surface_set_eps func(
		surface *Cairo_surface_t,
		eps Cairo_bool_t) *Char

	Cairo_ps_surface_get_eps func(
		surface *Cairo_surface_t)

	Cairo_ps_surface_set_size func(
		surface *Cairo_surface_t,
		width_in_points Double,
		height_in_points Double) Cairo_bool_t

	Cairo_ps_surface_dsc_comment func(
		surface *Cairo_surface_t,
		comment *Char)

	Cairo_ps_surface_dsc_begin_setup func(
		surface *Cairo_surface_t)

	Cairo_ps_surface_dsc_begin_page_setup func(
		surface *Cairo_surface_t)

	Cairo_svg_surface_create func(
		filename *Char,
		width_in_points Double,
		height_in_points Double) *Cairo_surface_t

	Cairo_svg_surface_create_for_stream func(
		write_func Cairo_write_func_t,
		closure *Void,
		width_in_points Double,
		height_in_points Double) *Cairo_surface_t

	Cairo_svg_surface_restrict_to_version func(
		surface *Cairo_surface_t,
		version Cairo_svg_version_t)

	Cairo_svg_get_versions func(
		versions **Cairo_svg_version_t,
		num_versions *int)

	Cairo_svg_version_to_string func(
		version Cairo_svg_version_t) *Char

	Cairo_win32_surface_create func(
		hdc HDC) *Cairo_surface_t

	Cairo_win32_printing_surface_create func(
		hdc HDC) *Cairo_surface_t

	Cairo_win32_surface_create_with_ddb func(
		hdc HDC,
		format Cairo_format_t,
		width int,
		height int) *Cairo_surface_t

	Cairo_win32_surface_create_with_dib func(
		format Cairo_format_t,
		width int,
		height int) *Cairo_surface_t

	Cairo_win32_surface_get_dc func(
		surface *Cairo_surface_t) HDC

	Cairo_win32_surface_get_image func(
		surface *Cairo_surface_t) *Cairo_surface_t

	Cairo_win32_font_face_create_for_logfontw func(
		logfont *LOGFONTW) *Cairo_font_face_t

	Cairo_win32_font_face_create_for_hfont func(
		font HFONT) *Cairo_font_face_t

	Cairo_win32_font_face_create_for_logfontw_hfont func(
		logfont *LOGFONTW,
		font HFONT) *Cairo_font_face_t

	Cairo_win32_scaled_font_select_font func(
		scaled_font *Cairo_scaled_font_t,
		hdc HDC) Cairo_status_t

	Cairo_win32_scaled_font_done_font func(
		scaled_font *Cairo_scaled_font_t)

	Cairo_win32_scaled_font_get_metrics_factor func(
		scaled_font *Cairo_scaled_font_t) Double

	Cairo_win32_scaled_font_get_logical_to_device func(
		scaled_font *Cairo_scaled_font_t,
		logical_to_device *Cairo_matrix_t)

	Cairo_win32_scaled_font_get_device_to_logical func(
		scaled_font *Cairo_scaled_font_t,
		device_to_logical *Cairo_matrix_t)
)

var dll = "libcairo-2.dll"

var apiList = outside.Apis{
	{"cairo_append_path", &Cairo_append_path},
	{"cairo_arc", &Cairo_arc},
	{"cairo_arc_negative", &Cairo_arc_negative},
	{"cairo_clip", &Cairo_clip},
	{"cairo_clip_extents", &Cairo_clip_extents},
	{"cairo_clip_preserve", &Cairo_clip_preserve},
	{"cairo_close_path", &Cairo_close_path},
	{"cairo_copy_clip_rectangle_list", &Cairo_copy_clip_rectangle_list},
	{"cairo_copy_page", &Cairo_copy_page},
	{"cairo_copy_path", &Cairo_copy_path},
	{"cairo_copy_path_flat", &Cairo_copy_path_flat},
	{"cairo_create", &Cairo_create},
	{"cairo_curve_to", &Cairo_curve_to},
	{"cairo_debug_reset_static_data", &Cairo_debug_reset_static_data},
	{"cairo_destroy", &Cairo_destroy},
	{"cairo_device_acquire", &Cairo_device_acquire},
	{"cairo_device_destroy", &Cairo_device_destroy},
	{"cairo_device_finish", &Cairo_device_finish},
	{"cairo_device_flush", &Cairo_device_flush},
	{"cairo_device_get_reference_count", &Cairo_device_get_reference_count},
	{"cairo_device_get_type", &Cairo_device_get_type},
	{"cairo_device_get_user_data", &Cairo_device_get_user_data},
	{"cairo_device_reference", &Cairo_device_reference},
	{"cairo_device_release", &Cairo_device_release},
	{"cairo_device_set_user_data", &Cairo_device_set_user_data},
	{"cairo_device_status", &Cairo_device_status},
	{"cairo_device_to_user", &Cairo_device_to_user},
	{"cairo_device_to_user_distance", &Cairo_device_to_user_distance},
	{"cairo_fill", &Cairo_fill},
	{"cairo_fill_extents", &Cairo_fill_extents},
	{"cairo_fill_preserve", &Cairo_fill_preserve},
	{"cairo_font_extents", &Cairo_font_extents},
	{"cairo_font_face_destroy", &Cairo_font_face_destroy},
	{"cairo_font_face_get_reference_count", &Cairo_font_face_get_reference_count},
	{"cairo_font_face_get_type", &Cairo_font_face_get_type},
	{"cairo_font_face_get_user_data", &Cairo_font_face_get_user_data},
	{"cairo_font_face_reference", &Cairo_font_face_reference},
	{"cairo_font_face_set_user_data", &Cairo_font_face_set_user_data},
	{"cairo_font_face_status", &Cairo_font_face_status},
	{"cairo_font_options_copy", &Cairo_font_options_copy},
	{"cairo_font_options_create", &Cairo_font_options_create},
	{"cairo_font_options_destroy", &Cairo_font_options_destroy},
	{"cairo_font_options_equal", &Cairo_font_options_equal},
	{"cairo_font_options_get_antialias", &Cairo_font_options_get_antialias},
	{"cairo_font_options_get_hint_metrics", &Cairo_font_options_get_hint_metrics},
	{"cairo_font_options_get_hint_style", &Cairo_font_options_get_hint_style},
	{"cairo_font_options_get_subpixel_order", &Cairo_font_options_get_subpixel_order},
	{"cairo_font_options_hash", &Cairo_font_options_hash},
	{"cairo_font_options_merge", &Cairo_font_options_merge},
	{"cairo_font_options_set_antialias", &Cairo_font_options_set_antialias},
	{"cairo_font_options_set_hint_metrics", &Cairo_font_options_set_hint_metrics},
	{"cairo_font_options_set_hint_style", &Cairo_font_options_set_hint_style},
	{"cairo_font_options_set_subpixel_order", &Cairo_font_options_set_subpixel_order},
	{"cairo_font_options_status", &Cairo_font_options_status},
	{"cairo_format_stride_for_width", &Cairo_format_stride_for_width},
	{"cairo_ft_font_face_create_for_ft_face", &Cairo_ft_font_face_create_for_ft_face},
	{"cairo_ft_font_face_create_for_pattern", &Cairo_ft_font_face_create_for_pattern},
	{"cairo_ft_font_options_substitute", &Cairo_ft_font_options_substitute},
	{"cairo_ft_scaled_font_lock_face", &Cairo_ft_scaled_font_lock_face},
	{"cairo_ft_scaled_font_unlock_face", &Cairo_ft_scaled_font_unlock_face},
	{"cairo_get_antialias", &Cairo_get_antialias},
	{"cairo_get_current_point", &Cairo_get_current_point},
	{"cairo_get_dash", &Cairo_get_dash},
	{"cairo_get_dash_count", &Cairo_get_dash_count},
	{"cairo_get_fill_rule", &Cairo_get_fill_rule},
	{"cairo_get_font_face", &Cairo_get_font_face},
	{"cairo_get_font_matrix", &Cairo_get_font_matrix},
	{"cairo_get_font_options", &Cairo_get_font_options},
	{"cairo_get_group_target", &Cairo_get_group_target},
	{"cairo_get_line_cap", &Cairo_get_line_cap},
	{"cairo_get_line_join", &Cairo_get_line_join},
	{"cairo_get_line_width", &Cairo_get_line_width},
	{"cairo_get_matrix", &Cairo_get_matrix},
	{"cairo_get_miter_limit", &Cairo_get_miter_limit},
	{"cairo_get_operator", &Cairo_get_operator},
	{"cairo_get_reference_count", &Cairo_get_reference_count},
	{"cairo_get_scaled_font", &Cairo_get_scaled_font},
	{"cairo_get_source", &Cairo_get_source},
	{"cairo_get_target", &Cairo_get_target},
	{"cairo_get_tolerance", &Cairo_get_tolerance},
	{"cairo_get_user_data", &Cairo_get_user_data},
	{"cairo_glyph_allocate", &Cairo_glyph_allocate},
	{"cairo_glyph_extents", &Cairo_glyph_extents},
	{"cairo_glyph_free", &Cairo_glyph_free},
	{"cairo_glyph_path", &Cairo_glyph_path},
	{"cairo_has_current_point", &Cairo_has_current_point},
	{"cairo_identity_matrix", &Cairo_identity_matrix},
	{"cairo_image_surface_create", &Cairo_image_surface_create},
	{"cairo_image_surface_create_for_data", &Cairo_image_surface_create_for_data},
	{"cairo_image_surface_create_from_png", &Cairo_image_surface_create_from_png},
	{"cairo_image_surface_create_from_png_stream", &Cairo_image_surface_create_from_png_stream},
	{"cairo_image_surface_get_data", &Cairo_image_surface_get_data},
	{"cairo_image_surface_get_format", &Cairo_image_surface_get_format},
	{"cairo_image_surface_get_height", &Cairo_image_surface_get_height},
	{"cairo_image_surface_get_stride", &Cairo_image_surface_get_stride},
	{"cairo_image_surface_get_width", &Cairo_image_surface_get_width},
	{"cairo_in_clip", &Cairo_in_clip},
	{"cairo_in_fill", &Cairo_in_fill},
	{"cairo_in_stroke", &Cairo_in_stroke},
	{"cairo_line_to", &Cairo_line_to},
	{"cairo_mask", &Cairo_mask},
	{"cairo_mask_surface", &Cairo_mask_surface},
	{"cairo_matrix_init", &Cairo_matrix_init},
	{"cairo_matrix_init_identity", &Cairo_matrix_init_identity},
	{"cairo_matrix_init_rotate", &Cairo_matrix_init_rotate},
	{"cairo_matrix_init_scale", &Cairo_matrix_init_scale},
	{"cairo_matrix_init_translate", &Cairo_matrix_init_translate},
	{"cairo_matrix_invert", &Cairo_matrix_invert},
	{"cairo_matrix_multiply", &Cairo_matrix_multiply},
	{"cairo_matrix_rotate", &Cairo_matrix_rotate},
	{"cairo_matrix_scale", &Cairo_matrix_scale},
	{"cairo_matrix_transform_distance", &Cairo_matrix_transform_distance},
	{"cairo_matrix_transform_point", &Cairo_matrix_transform_point},
	{"cairo_matrix_translate", &Cairo_matrix_translate},
	{"cairo_move_to", &Cairo_move_to},
	{"cairo_new_path", &Cairo_new_path},
	{"cairo_new_sub_path", &Cairo_new_sub_path},
	{"cairo_paint", &Cairo_paint},
	{"cairo_paint_with_alpha", &Cairo_paint_with_alpha},
	{"cairo_path_destroy", &Cairo_path_destroy},
	{"cairo_path_extents", &Cairo_path_extents},
	{"cairo_pattern_add_color_stop_rgb", &Cairo_pattern_add_color_stop_rgb},
	{"cairo_pattern_add_color_stop_rgba", &Cairo_pattern_add_color_stop_rgba},
	{"cairo_pattern_create_for_surface", &Cairo_pattern_create_for_surface},
	{"cairo_pattern_create_linear", &Cairo_pattern_create_linear},
	{"cairo_pattern_create_radial", &Cairo_pattern_create_radial},
	{"cairo_pattern_create_rgb", &Cairo_pattern_create_rgb},
	{"cairo_pattern_create_rgba", &Cairo_pattern_create_rgba},
	{"cairo_pattern_destroy", &Cairo_pattern_destroy},
	{"cairo_pattern_get_color_stop_count", &Cairo_pattern_get_color_stop_count},
	{"cairo_pattern_get_color_stop_rgba", &Cairo_pattern_get_color_stop_rgba},
	{"cairo_pattern_get_extend", &Cairo_pattern_get_extend},
	{"cairo_pattern_get_filter", &Cairo_pattern_get_filter},
	{"cairo_pattern_get_linear_points", &Cairo_pattern_get_linear_points},
	{"cairo_pattern_get_matrix", &Cairo_pattern_get_matrix},
	{"cairo_pattern_get_radial_circles", &Cairo_pattern_get_radial_circles},
	{"cairo_pattern_get_reference_count", &Cairo_pattern_get_reference_count},
	{"cairo_pattern_get_rgba", &Cairo_pattern_get_rgba},
	{"cairo_pattern_get_surface", &Cairo_pattern_get_surface},
	{"cairo_pattern_get_type", &Cairo_pattern_get_type},
	{"cairo_pattern_get_user_data", &Cairo_pattern_get_user_data},
	{"cairo_pattern_reference", &Cairo_pattern_reference},
	{"cairo_pattern_set_extend", &Cairo_pattern_set_extend},
	{"cairo_pattern_set_filter", &Cairo_pattern_set_filter},
	{"cairo_pattern_set_matrix", &Cairo_pattern_set_matrix},
	{"cairo_pattern_set_user_data", &Cairo_pattern_set_user_data},
	{"cairo_pattern_status", &Cairo_pattern_status},
	{"cairo_pdf_get_versions", &Cairo_pdf_get_versions},
	{"cairo_pdf_surface_create", &Cairo_pdf_surface_create},
	{"cairo_pdf_surface_create_for_stream", &Cairo_pdf_surface_create_for_stream},
	{"cairo_pdf_surface_restrict_to_version", &Cairo_pdf_surface_restrict_to_version},
	{"cairo_pdf_surface_set_size", &Cairo_pdf_surface_set_size},
	{"cairo_pdf_version_to_string", &Cairo_pdf_version_to_string},
	{"cairo_pop_group", &Cairo_pop_group},
	{"cairo_pop_group_to_source", &Cairo_pop_group_to_source},
	{"cairo_ps_get_levels", &Cairo_ps_get_levels},
	{"cairo_ps_level_to_string", &Cairo_ps_level_to_string},
	{"cairo_ps_surface_create", &Cairo_ps_surface_create},
	{"cairo_ps_surface_create_for_stream", &Cairo_ps_surface_create_for_stream},
	{"cairo_ps_surface_dsc_begin_page_setup", &Cairo_ps_surface_dsc_begin_page_setup},
	{"cairo_ps_surface_dsc_begin_setup", &Cairo_ps_surface_dsc_begin_setup},
	{"cairo_ps_surface_dsc_comment", &Cairo_ps_surface_dsc_comment},
	{"cairo_ps_surface_get_eps", &Cairo_ps_surface_get_eps},
	{"cairo_ps_surface_restrict_to_level", &Cairo_ps_surface_restrict_to_level},
	{"cairo_ps_surface_set_eps", &Cairo_ps_surface_set_eps},
	{"cairo_ps_surface_set_size", &Cairo_ps_surface_set_size},
	{"cairo_push_group", &Cairo_push_group},
	{"cairo_push_group_with_content", &Cairo_push_group_with_content},
	{"cairo_recording_surface_create", &Cairo_recording_surface_create},
	{"cairo_recording_surface_ink_extents", &Cairo_recording_surface_ink_extents},
	{"cairo_rectangle", &Cairo_rectangle},
	{"cairo_rectangle_list_destroy", &Cairo_rectangle_list_destroy},
	{"cairo_reference", &Cairo_reference},
	{"cairo_region_contains_point", &Cairo_region_contains_point},
	{"cairo_region_contains_rectangle", &Cairo_region_contains_rectangle},
	{"cairo_region_copy", &Cairo_region_copy},
	{"cairo_region_create", &Cairo_region_create},
	{"cairo_region_create_rectangle", &Cairo_region_create_rectangle},
	{"cairo_region_create_rectangles", &Cairo_region_create_rectangles},
	{"cairo_region_destroy", &Cairo_region_destroy},
	{"cairo_region_equal", &Cairo_region_equal},
	{"cairo_region_get_extents", &Cairo_region_get_extents},
	{"cairo_region_get_rectangle", &Cairo_region_get_rectangle},
	{"cairo_region_intersect", &Cairo_region_intersect},
	{"cairo_region_intersect_rectangle", &Cairo_region_intersect_rectangle},
	{"cairo_region_is_empty", &Cairo_region_is_empty},
	{"cairo_region_num_rectangles", &Cairo_region_num_rectangles},
	{"cairo_region_reference", &Cairo_region_reference},
	{"cairo_region_status", &Cairo_region_status},
	{"cairo_region_subtract", &Cairo_region_subtract},
	{"cairo_region_subtract_rectangle", &Cairo_region_subtract_rectangle},
	{"cairo_region_translate", &Cairo_region_translate},
	{"cairo_region_union", &Cairo_region_union},
	{"cairo_region_union_rectangle", &Cairo_region_union_rectangle},
	{"cairo_region_xor", &Cairo_region_xor},
	{"cairo_region_xor_rectangle", &Cairo_region_xor_rectangle},
	{"cairo_rel_curve_to", &Cairo_rel_curve_to},
	{"cairo_rel_line_to", &Cairo_rel_line_to},
	{"cairo_rel_move_to", &Cairo_rel_move_to},
	{"cairo_reset_clip", &Cairo_reset_clip},
	{"cairo_restore", &Cairo_restore},
	{"cairo_rotate", &Cairo_rotate},
	{"cairo_save", &Cairo_save},
	{"cairo_scale", &Cairo_scale},
	{"cairo_scaled_font_create", &Cairo_scaled_font_create},
	{"cairo_scaled_font_destroy", &Cairo_scaled_font_destroy},
	{"cairo_scaled_font_extents", &Cairo_scaled_font_extents},
	{"cairo_scaled_font_get_ctm", &Cairo_scaled_font_get_ctm},
	{"cairo_scaled_font_get_font_face", &Cairo_scaled_font_get_font_face},
	{"cairo_scaled_font_get_font_matrix", &Cairo_scaled_font_get_font_matrix},
	{"cairo_scaled_font_get_font_options", &Cairo_scaled_font_get_font_options},
	{"cairo_scaled_font_get_reference_count", &Cairo_scaled_font_get_reference_count},
	{"cairo_scaled_font_get_scale_matrix", &Cairo_scaled_font_get_scale_matrix},
	{"cairo_scaled_font_get_type", &Cairo_scaled_font_get_type},
	{"cairo_scaled_font_get_user_data", &Cairo_scaled_font_get_user_data},
	{"cairo_scaled_font_glyph_extents", &Cairo_scaled_font_glyph_extents},
	{"cairo_scaled_font_reference", &Cairo_scaled_font_reference},
	{"cairo_scaled_font_set_user_data", &Cairo_scaled_font_set_user_data},
	{"cairo_scaled_font_status", &Cairo_scaled_font_status},
	{"cairo_scaled_font_text_extents", &Cairo_scaled_font_text_extents},
	{"cairo_scaled_font_text_to_glyphs", &Cairo_scaled_font_text_to_glyphs},
	{"cairo_select_font_face", &Cairo_select_font_face},
	{"cairo_set_antialias", &Cairo_set_antialias},
	{"cairo_set_dash", &Cairo_set_dash},
	{"cairo_set_fill_rule", &Cairo_set_fill_rule},
	{"cairo_set_font_face", &Cairo_set_font_face},
	{"cairo_set_font_matrix", &Cairo_set_font_matrix},
	{"cairo_set_font_options", &Cairo_set_font_options},
	{"cairo_set_font_size", &Cairo_set_font_size},
	{"cairo_set_line_cap", &Cairo_set_line_cap},
	{"cairo_set_line_join", &Cairo_set_line_join},
	{"cairo_set_line_width", &Cairo_set_line_width},
	{"cairo_set_matrix", &Cairo_set_matrix},
	{"cairo_set_miter_limit", &Cairo_set_miter_limit},
	{"cairo_set_operator", &Cairo_set_operator},
	{"cairo_set_scaled_font", &Cairo_set_scaled_font},
	{"cairo_set_source", &Cairo_set_source},
	{"cairo_set_source_rgb", &Cairo_set_source_rgb},
	{"cairo_set_source_rgba", &Cairo_set_source_rgba},
	{"cairo_set_source_surface", &Cairo_set_source_surface},
	{"cairo_set_tolerance", &Cairo_set_tolerance},
	{"cairo_set_user_data", &Cairo_set_user_data},
	{"cairo_show_glyphs", &Cairo_show_glyphs},
	{"cairo_show_page", &Cairo_show_page},
	{"cairo_show_text", &Cairo_show_text},
	{"cairo_show_text_glyphs", &Cairo_show_text_glyphs},
	{"cairo_status", &Cairo_status},
	{"cairo_status_to_string", &Cairo_status_to_string},
	{"cairo_stroke", &Cairo_stroke},
	{"cairo_stroke_extents", &Cairo_stroke_extents},
	{"cairo_stroke_preserve", &Cairo_stroke_preserve},
	{"cairo_surface_copy_page", &Cairo_surface_copy_page},
	{"cairo_surface_create_for_rectangle", &Cairo_surface_create_for_rectangle},
	{"cairo_surface_create_similar", &Cairo_surface_create_similar},
	{"cairo_surface_destroy", &Cairo_surface_destroy},
	{"cairo_surface_finish", &Cairo_surface_finish},
	{"cairo_surface_flush", &Cairo_surface_flush},
	{"cairo_surface_get_content", &Cairo_surface_get_content},
	{"cairo_surface_get_device", &Cairo_surface_get_device},
	{"cairo_surface_get_device_offset", &Cairo_surface_get_device_offset},
	{"cairo_surface_get_fallback_resolution", &Cairo_surface_get_fallback_resolution},
	{"cairo_surface_get_font_options", &Cairo_surface_get_font_options},
	{"cairo_surface_get_mime_data", &Cairo_surface_get_mime_data},
	{"cairo_surface_get_reference_count", &Cairo_surface_get_reference_count},
	{"cairo_surface_get_type", &Cairo_surface_get_type},
	{"cairo_surface_get_user_data", &Cairo_surface_get_user_data},
	{"cairo_surface_has_show_text_glyphs", &Cairo_surface_has_show_text_glyphs},
	{"cairo_surface_mark_dirty", &Cairo_surface_mark_dirty},
	{"cairo_surface_mark_dirty_rectangle", &Cairo_surface_mark_dirty_rectangle},
	{"cairo_surface_reference", &Cairo_surface_reference},
	{"cairo_surface_set_device_offset", &Cairo_surface_set_device_offset},
	{"cairo_surface_set_fallback_resolution", &Cairo_surface_set_fallback_resolution},
	{"cairo_surface_set_mime_data", &Cairo_surface_set_mime_data},
	{"cairo_surface_set_user_data", &Cairo_surface_set_user_data},
	{"cairo_surface_show_page", &Cairo_surface_show_page},
	{"cairo_surface_status", &Cairo_surface_status},
	{"cairo_surface_write_to_png", &Cairo_surface_write_to_png},
	{"cairo_surface_write_to_png_stream", &Cairo_surface_write_to_png_stream},
	{"cairo_svg_get_versions", &Cairo_svg_get_versions},
	{"cairo_svg_surface_create", &Cairo_svg_surface_create},
	{"cairo_svg_surface_create_for_stream", &Cairo_svg_surface_create_for_stream},
	{"cairo_svg_surface_restrict_to_version", &Cairo_svg_surface_restrict_to_version},
	{"cairo_svg_version_to_string", &Cairo_svg_version_to_string},
	{"cairo_text_cluster_allocate", &Cairo_text_cluster_allocate},
	{"cairo_text_cluster_free", &Cairo_text_cluster_free},
	{"cairo_text_extents", &Cairo_text_extents},
	{"cairo_text_path", &Cairo_text_path},
	{"cairo_toy_font_face_create", &Cairo_toy_font_face_create},
	{"cairo_toy_font_face_get_family", &Cairo_toy_font_face_get_family},
	{"cairo_toy_font_face_get_slant", &Cairo_toy_font_face_get_slant},
	{"cairo_toy_font_face_get_weight", &Cairo_toy_font_face_get_weight},
	{"cairo_transform", &Cairo_transform},
	{"cairo_translate", &Cairo_translate},
	{"cairo_user_font_face_create", &Cairo_user_font_face_create},
	{"cairo_user_font_face_get_init_func", &Cairo_user_font_face_get_init_func},
	{"cairo_user_font_face_get_render_glyph_func", &Cairo_user_font_face_get_render_glyph_func},
	{"cairo_user_font_face_get_text_to_glyphs_func", &Cairo_user_font_face_get_text_to_glyphs_func},
	{"cairo_user_font_face_get_unicode_to_glyph_func", &Cairo_user_font_face_get_unicode_to_glyph_func},
	{"cairo_user_font_face_set_init_func", &Cairo_user_font_face_set_init_func},
	{"cairo_user_font_face_set_render_glyph_func", &Cairo_user_font_face_set_render_glyph_func},
	{"cairo_user_font_face_set_text_to_glyphs_func", &Cairo_user_font_face_set_text_to_glyphs_func},
	{"cairo_user_font_face_set_unicode_to_glyph_func", &Cairo_user_font_face_set_unicode_to_glyph_func},
	{"cairo_user_to_device", &Cairo_user_to_device},
	{"cairo_user_to_device_distance", &Cairo_user_to_device_distance},
	{"cairo_version", &Cairo_version},
	{"cairo_version_string", &Cairo_version_string},
	{"cairo_win32_font_face_create_for_hfont", &Cairo_win32_font_face_create_for_hfont},
	{"cairo_win32_font_face_create_for_logfontw", &Cairo_win32_font_face_create_for_logfontw},
	{"cairo_win32_font_face_create_for_logfontw_hfont", &Cairo_win32_font_face_create_for_logfontw_hfont},
	{"cairo_win32_printing_surface_create", &Cairo_win32_printing_surface_create},
	{"cairo_win32_scaled_font_done_font", &Cairo_win32_scaled_font_done_font},
	{"cairo_win32_scaled_font_get_device_to_logical", &Cairo_win32_scaled_font_get_device_to_logical},
	{"cairo_win32_scaled_font_get_logical_to_device", &Cairo_win32_scaled_font_get_logical_to_device},
	{"cairo_win32_scaled_font_get_metrics_factor", &Cairo_win32_scaled_font_get_metrics_factor},
	{"cairo_win32_scaled_font_select_font", &Cairo_win32_scaled_font_select_font},
	{"cairo_win32_surface_create", &Cairo_win32_surface_create},
	{"cairo_win32_surface_create_with_ddb", &Cairo_win32_surface_create_with_ddb},
	{"cairo_win32_surface_create_with_dib", &Cairo_win32_surface_create_with_dib},
	{"cairo_win32_surface_get_dc", &Cairo_win32_surface_get_dc},
	{"cairo_win32_surface_get_image", &Cairo_win32_surface_get_image},
}

var dllGobject = "libcairo-gobject-2.dll"

var apiListGobject = outside.Apis{
	{"cairo_gobject_antialias_get_type", &Cairo_gobject_antialias_get_type},
	{"cairo_gobject_content_get_type", &Cairo_gobject_content_get_type},
	{"cairo_gobject_context_get_type", &Cairo_gobject_context_get_type},
	{"cairo_gobject_device_get_type", &Cairo_gobject_device_get_type},
	{"cairo_gobject_device_type_get_type", &Cairo_gobject_device_type_get_type},
	{"cairo_gobject_extend_get_type", &Cairo_gobject_extend_get_type},
	{"cairo_gobject_fill_rule_get_type", &Cairo_gobject_fill_rule_get_type},
	{"cairo_gobject_filter_get_type", &Cairo_gobject_filter_get_type},
	{"cairo_gobject_font_face_get_type", &Cairo_gobject_font_face_get_type},
	{"cairo_gobject_font_options_get_type", &Cairo_gobject_font_options_get_type},
	{"cairo_gobject_font_slant_get_type", &Cairo_gobject_font_slant_get_type},
	{"cairo_gobject_font_type_get_type", &Cairo_gobject_font_type_get_type},
	{"cairo_gobject_font_weight_get_type", &Cairo_gobject_font_weight_get_type},
	{"cairo_gobject_format_get_type", &Cairo_gobject_format_get_type},
	{"cairo_gobject_hint_metrics_get_type", &Cairo_gobject_hint_metrics_get_type},
	{"cairo_gobject_hint_style_get_type", &Cairo_gobject_hint_style_get_type},
	{"cairo_gobject_line_cap_get_type", &Cairo_gobject_line_cap_get_type},
	{"cairo_gobject_line_join_get_type", &Cairo_gobject_line_join_get_type},
	{"cairo_gobject_operator_get_type", &Cairo_gobject_operator_get_type},
	{"cairo_gobject_path_data_type_get_type", &Cairo_gobject_path_data_type_get_type},
	{"cairo_gobject_pattern_get_type", &Cairo_gobject_pattern_get_type},
	{"cairo_gobject_pattern_type_get_type", &Cairo_gobject_pattern_type_get_type},
	{"cairo_gobject_rectangle_get_type", &Cairo_gobject_rectangle_get_type},
	{"cairo_gobject_rectangle_int_get_type", &Cairo_gobject_rectangle_int_get_type},
	{"cairo_gobject_region_get_type", &Cairo_gobject_region_get_type},
	{"cairo_gobject_region_overlap_get_type", &Cairo_gobject_region_overlap_get_type},
	{"cairo_gobject_scaled_font_get_type", &Cairo_gobject_scaled_font_get_type},
	{"cairo_gobject_status_get_type", &Cairo_gobject_status_get_type},
	{"cairo_gobject_subpixel_order_get_type", &Cairo_gobject_subpixel_order_get_type},
	{"cairo_gobject_surface_get_type", &Cairo_gobject_surface_get_type},
	{"cairo_gobject_surface_type_get_type", &Cairo_gobject_surface_type_get_type},
	{"cairo_gobject_text_cluster_flags_get_type", &Cairo_gobject_text_cluster_flags_get_type},
}

var dllScript = "libcairo-script-interpreter-2.dll"

var apiListScript = outside.Apis{
	// Undocumented {"_csi_alloc", &Csi_alloc},
	// Undocumented {"_csi_alloc0", &Csi_alloc0},
	// Undocumented {"_csi_array_execute", &Csi_array_execute},
	// Undocumented {"_csi_error", &Csi_error},
	// Undocumented {"_csi_file_as_string", &Csi_file_as_string},
	// Undocumented {"_csi_file_execute", &Csi_file_execute},
	// Undocumented {"_csi_file_free", &Csi_file_free},
	// Undocumented {"_csi_free", &Csi_free},
	// Undocumented {"_csi_hash_table_fini", &Csi_hash_table_fini},
	// Undocumented {"_csi_hash_table_foreach", &Csi_hash_table_foreach},
	// Undocumented {"_csi_hash_table_init", &Csi_hash_table_init},
	// Undocumented {"_csi_hash_table_insert", &Csi_hash_table_insert},
	// Undocumented {"_csi_hash_table_lookup", &Csi_hash_table_lookup},
	// Undocumented {"_csi_hash_table_remove", &Csi_hash_table_remove},
	// Undocumented {"_csi_integer_constants", &Csi_integer_constants},
	// Undocumented {"_csi_intern_string", &Csi_intern_string},
	// Undocumented {"_csi_name_define", &Csi_name_define},
	// Undocumented {"_csi_name_lookup", &Csi_name_lookup},
	// Undocumented {"_csi_name_undefine", &Csi_name_undefine},
	// Undocumented {"_csi_operators", &Csi_operators},
	// Undocumented {"_csi_parse_number", &Csi_parse_number},
	// Undocumented {"_csi_perm_alloc", &Csi_perm_alloc},
	// Undocumented {"_csi_real_constants", &Csi_real_constants},
	// Undocumented {"_csi_realloc", &Csi_realloc},
	// Undocumented {"_csi_scan_file", &Csi_scan_file},
	// Undocumented {"_csi_scanner_fini", &Csi_scanner_fini},
	// Undocumented {"_csi_scanner_init", &Csi_scanner_init},
	// Undocumented {"_csi_slab_alloc", &Csi_slab_alloc},
	// Undocumented {"_csi_slab_free", &Csi_slab_free},
	// Undocumented {"_csi_stack_exch", &Csi_stack_exch},
	// Undocumented {"_csi_stack_fini", &Csi_stack_fini},
	// Undocumented {"_csi_stack_grow", &Csi_stack_grow},
	// Undocumented {"_csi_stack_init", &Csi_stack_init},
	// Undocumented {"_csi_stack_peek", &Csi_stack_peek},
	// Undocumented {"_csi_stack_pop", &Csi_stack_pop},
	// Undocumented {"_csi_stack_push_internal", &Csi_stack_push_internal},
	// Undocumented {"_csi_stack_roll", &Csi_stack_roll},
	// Undocumented {"_csi_translate_file", &Csi_translate_file},
	// Undocumented {"_get_output_format", &Get_output_format},
	{"cairo_script_interpreter_create", &Cairo_script_interpreter_create},
	{"cairo_script_interpreter_destroy", &Cairo_script_interpreter_destroy},
	{"cairo_script_interpreter_feed_stream", &Cairo_script_interpreter_feed_stream},
	{"cairo_script_interpreter_feed_string", &Cairo_script_interpreter_feed_string},
	{"cairo_script_interpreter_finish", &Cairo_script_interpreter_finish},
	{"cairo_script_interpreter_get_line_number", &Cairo_script_interpreter_get_line_number},
	{"cairo_script_interpreter_install_hooks", &Cairo_script_interpreter_install_hooks},
	{"cairo_script_interpreter_reference", &Cairo_script_interpreter_reference},
	{"cairo_script_interpreter_run", &Cairo_script_interpreter_run},
	{"cairo_script_interpreter_translate_stream", &Cairo_script_interpreter_translate_stream},
	// Undocumented {"csi_array_append", &Csi_array_append},
	// Undocumented {"csi_array_free", &Csi_array_free},
	// Undocumented {"csi_array_get", &Csi_array_get},
	// Undocumented {"csi_array_new", &Csi_array_new},
	// Undocumented {"csi_array_put", &Csi_array_put},
	// Undocumented {"csi_dictionary_free", &Csi_dictionary_free},
	// Undocumented {"csi_dictionary_get", &Csi_dictionary_get},
	// Undocumented {"csi_dictionary_has", &Csi_dictionary_has},
	// Undocumented {"csi_dictionary_new", &Csi_dictionary_new},
	// Undocumented {"csi_dictionary_put", &Csi_dictionary_put},
	// Undocumented {"csi_dictionary_remove", &Csi_dictionary_remove},
	// Undocumented {"csi_file_close", &Csi_file_close},
	// Undocumented {"csi_file_flush", &Csi_file_flush},
	// Undocumented {"csi_file_getc", &Csi_file_getc},
	// Undocumented {"csi_file_new", &Csi_file_new},
	// Undocumented {"csi_file_new_ascii85_decode", &Csi_file_new_ascii85_decode},
	// Undocumented {"csi_file_new_deflate_decode", &Csi_file_new_deflate_decode},
	// Undocumented {"csi_file_new_for_bytes", &Csi_file_new_for_bytes},
	// Undocumented {"csi_file_new_for_stream", &Csi_file_new_for_stream},
	// Undocumented {"csi_file_new_from_string", &Csi_file_new_from_string},
	// Undocumented {"csi_file_putc", &Csi_file_putc},
	// Undocumented {"csi_file_read", &Csi_file_read},
	// Undocumented {"csi_matrix_free", &Csi_matrix_free},
	// Undocumented {"csi_matrix_new", &Csi_matrix_new},
	// Undocumented {"csi_matrix_new_from_array", &Csi_matrix_new_from_array},
	// Undocumented {"csi_matrix_new_from_matrix", &Csi_matrix_new_from_matrix},
	// Undocumented {"csi_matrix_new_from_values", &Csi_matrix_new_from_values},
	// Undocumented {"csi_name_new", &Csi_name_new},
	// Undocumented {"csi_name_new_static", &Csi_name_new_static},
	// Undocumented {"csi_object_as_file", &Csi_object_as_file},
	// Undocumented {"csi_object_compare", &Csi_object_compare},
	// Undocumented {"csi_object_eq", &Csi_object_eq},
	// Undocumented {"csi_object_execute", &Csi_object_execute},
	// Undocumented {"csi_object_free", &Csi_object_free},
	// Undocumented {"csi_object_reference", &Csi_object_reference},
	// Undocumented {"csi_string_deflate_new", &Csi_string_deflate_new},
	// Undocumented {"csi_string_free", &Csi_string_free},
	// Undocumented {"csi_string_new", &Csi_string_new},
	// Undocumented {"csi_string_new_from_bytes", &Csi_string_new_from_bytes},
}
