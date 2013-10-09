// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package pangowin provides API definitions for accessing
//libpangowin32-1.0-0.dll.
package pangowin

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
	w "github.com/tHinqa/outside-windows/types"
)

//TODO(t): sort out A/W once structure better handled in outside

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	Pango_win32_get_context func() *PangoContext

	Pango_win32_render func(hdc w.HDC,
		font *PangoFont, glyphs *PangoGlyphString, x, y Gint)

	Pango_win32_render_layout_line func(
		hdc w.HDC, line *PangoLayoutLine, x, y int)

	Pango_win32_render_layout func(
		hdc w.HDC, layout *PangoLayout, x, y int)

	Pango_win32_render_transformed func(
		hdc w.HDC,
		matrix *PangoMatrix,
		font *PangoFont,
		glyphs *PangoGlyphString,
		x, y int)

	Pango_win32_get_unknown_glyph func(
		font *PangoFont, wc Gunichar) PangoGlyph

	Pango_win32_font_get_glyph_index func(
		font *PangoFont, wc Gunichar) Gint

	Pango_win32_get_dc func() w.HDC

	Pango_win32_get_debug_flag func() Gboolean

	Pango_win32_font_select_font func(
		font *PangoFont, hdc w.HDC) Gboolean

	Pango_win32_font_done_font func(font *PangoFont)

	Pango_win32_font_get_metrics_factor func(font *PangoFont) Double

	Pango_win32_font_cache_new func() *PangoWin32FontCache

	Pango_win32_font_cache_free func(cache *PangoWin32FontCache)

	Pango_win32_font_cache_load func(
		cache *PangoWin32FontCache, logfont *w.LOGFONTA) w.HFONT

	Pango_win32_font_cache_loadw func(
		cache *PangoWin32FontCache, logfont *w.LOGFONTW) w.HFONT

	Pango_win32_font_cache_unload func(
		cache *PangoWin32FontCache, hfont w.HFONT)

	Pango_win32_font_map_for_display func() *PangoFontMap

	Pango_win32_shutdown_display func()

	Pango_win32_font_map_get_font_cache func(
		font_map *PangoFontMap) *PangoWin32FontCache

	Pango_win32_font_logfont func(font *PangoFont) *w.LOGFONTA

	Pango_win32_font_logfontw func(font *PangoFont) *w.LOGFONTW

	Pango_win32_font_description_from_logfont func(
		lfp *w.LOGFONTA) *PangoFontDescription

	Pango_win32_font_description_from_logfontw func(
		lfp *w.LOGFONTW) *PangoFontDescription
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
