// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package pangowin provides API definitions for accessing
//libpangowin32-1.0-0.dll.
package pangowin

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
	w "github.com/tHinqa/outside-windows/types"
)

//TODO(t): sort out A/W once structure better handled in outside

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	Pango_win32_get_context func() *T.PangoContext

	Pango_win32_render func(hdc w.HDC,
		font *T.PangoFont, glyphs *T.PangoGlyphString, x, y T.Gint)

	Pango_win32_render_layout_line func(
		hdc w.HDC, line *T.PangoLayoutLine, x, y int)

	Pango_win32_render_layout func(
		hdc w.HDC, layout *T.PangoLayout, x, y int)

	Pango_win32_render_transformed func(
		hdc w.HDC,
		matrix *T.PangoMatrix,
		font *T.PangoFont,
		glyphs *T.PangoGlyphString,
		x, y int)

	Pango_win32_get_unknown_glyph func(
		font *T.PangoFont, wc T.Gunichar) T.PangoGlyph

	Pango_win32_font_get_glyph_index func(
		font *T.PangoFont, wc T.Gunichar) T.Gint

	Pango_win32_get_dc func() w.HDC

	Pango_win32_get_debug_flag func() T.Gboolean

	Pango_win32_font_select_font func(
		font *T.PangoFont, hdc w.HDC) T.Gboolean

	Pango_win32_font_done_font func(font *T.PangoFont)

	Pango_win32_font_get_metrics_factor func(font *T.PangoFont) T.Double

	Pango_win32_font_cache_new func() *T.PangoWin32FontCache

	Pango_win32_font_cache_free func(cache *T.PangoWin32FontCache)

	Pango_win32_font_cache_load func(
		cache *T.PangoWin32FontCache, logfont *w.LOGFONTA) w.HFONT

	Pango_win32_font_cache_loadw func(
		cache *T.PangoWin32FontCache, logfont *w.LOGFONTW) w.HFONT

	Pango_win32_font_cache_unload func(
		cache *T.PangoWin32FontCache, hfont w.HFONT)

	Pango_win32_font_map_for_display func() *T.PangoFontMap

	Pango_win32_shutdown_display func()

	Pango_win32_font_map_get_font_cache func(
		font_map *T.PangoFontMap) *T.PangoWin32FontCache

	Pango_win32_font_logfont func(font *T.PangoFont) *w.LOGFONTA

	Pango_win32_font_logfontw func(font *T.PangoFont) *w.LOGFONTW

	Pango_win32_font_description_from_logfont func(
		lfp *w.LOGFONTA) *T.PangoFontDescription

	Pango_win32_font_description_from_logfontw func(
		lfp *w.LOGFONTW) *T.PangoFontDescription
)

var dll = "libpangowin32-1.0-0.dll"

var apiList = outside.Apis{
	// Undocumented {"_pango_win32_font_get_hfont",&_pango_win32_font_get_hfont},
	// Undocumented {"_pango_win32_font_get_type",&_pango_win32_font_get_type},
	// Undocumented {"_pango_win32_font_map_get_type",&_pango_win32_font_map_get_type},
	// Undocumented {"_pango_win32_fontmap_cache_remove",&_pango_win32_fontmap_cache_remove},
	// Undocumented {"_pango_win32_make_matching_logfontw",&_pango_win32_make_matching_logfontw},
	{"pango_win32_font_cache_free", &Pango_win32_font_cache_free},
	{"pango_win32_font_cache_load", &Pango_win32_font_cache_load},
	{"pango_win32_font_cache_loadw", &Pango_win32_font_cache_loadw},
	{"pango_win32_font_cache_new", &Pango_win32_font_cache_new},
	{"pango_win32_font_cache_unload", &Pango_win32_font_cache_unload},
	{"pango_win32_font_description_from_logfont", &Pango_win32_font_description_from_logfont},
	{"pango_win32_font_description_from_logfontw", &Pango_win32_font_description_from_logfontw},
	{"pango_win32_font_done_font", &Pango_win32_font_done_font},
	{"pango_win32_font_get_glyph_index", &Pango_win32_font_get_glyph_index},
	{"pango_win32_font_get_metrics_factor", &Pango_win32_font_get_metrics_factor},
	{"pango_win32_font_logfont", &Pango_win32_font_logfont},
	{"pango_win32_font_logfontw", &Pango_win32_font_logfontw},
	{"pango_win32_font_map_for_display", &Pango_win32_font_map_for_display},
	{"pango_win32_font_map_get_font_cache", &Pango_win32_font_map_get_font_cache},
	{"pango_win32_font_select_font", &Pango_win32_font_select_font},
	{"pango_win32_get_context", &Pango_win32_get_context},
	{"pango_win32_get_dc", &Pango_win32_get_dc},
	{"pango_win32_get_debug_flag", &Pango_win32_get_debug_flag},
	{"pango_win32_get_unknown_glyph", &Pango_win32_get_unknown_glyph},
	{"pango_win32_render", &Pango_win32_render},
	{"pango_win32_render_layout", &Pango_win32_render_layout},
	{"pango_win32_render_layout_line", &Pango_win32_render_layout_line},
	{"pango_win32_render_transformed", &Pango_win32_render_transformed},
	{"pango_win32_shutdown_display", &Pango_win32_shutdown_display},
}
