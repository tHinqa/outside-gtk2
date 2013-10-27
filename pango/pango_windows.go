// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//On Windows systems there are also definitions for accessing
//libpangowin32-1.0-0.dll.
package pango

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
	W "github.com/tHinqa/outside-windows/types"
)

//TODO(t): sort out A/W once structure better handled in outside

func init() {
	outside.AddDllApis(dllwin32, false, apiListwin32)
}

type Win32FontCache struct{}

var (
	Win32GetContext func() *Context

	Win32Render func(hdc W.HDC,
		font *Font, glyphs *GlyphString, x, y int)

	Win32RenderLayoutLine func(
		hdc W.HDC, line *LayoutLine, x, y int)

	Win32RenderLayout func(
		hdc W.HDC, layout *Layout, x, y int)

	Win32RenderTransformed func(
		hdc W.HDC,
		matrix *Matrix,
		font *Font,
		glyphs *GlyphString,
		x, y int)

	Win32GetDc func() W.HDC

	Win32GetDebugFlag func() T.Gboolean

	Win32FontMapForDisplay func() *FontMap

	Win32ShutdownDisplay func()

	Win32FontMapGetFontCache func(
		font_map *FontMap) *Win32FontCache

	Win32FontDescriptionFromLogfont func(
		lfp *W.LOGFONTA) *FontDescription

	Win32FontDescriptionFromLogfontw func(
		lfp *W.LOGFONTW) *FontDescription
)

var (
	win32FontDoneFont         func(font *Font)
	win32FontGetGlyphIndex    func(font *Font, wc T.Gunichar) int
	win32FontGetMetricsFactor func(font *Font) float64
	win32FontLogfont          func(font *Font) *W.LOGFONTA
	win32FontLogfontw         func(font *Font) *W.LOGFONTW
	win32FontSelectFont       func(font *Font, hdc W.HDC) T.Gboolean
	win32GetUnknownGlyph      func(font *Font, wc T.Gunichar) Glyph
)

func (f *Font) DoneFont()                           { win32FontDoneFont(f) }
func (f *Font) GetGlyphIndex(wc T.Gunichar) int     { return win32FontGetGlyphIndex(f, wc) }
func (f *Font) GetMetricsFactor() float64           { return win32FontGetMetricsFactor(f) }
func (f *Font) Logfont() *W.LOGFONTA                { return win32FontLogfont(f) }
func (f *Font) Logfontw() *W.LOGFONTW               { return win32FontLogfontw(f) }
func (f *Font) SelectFont(hdc W.HDC) T.Gboolean     { return win32FontSelectFont(f, hdc) }
func (f *Font) GetUnknownGlyph(wc T.Gunichar) Glyph { return win32GetUnknownGlyph(f, wc) }

var (
	Win32FontCacheNew func() *Win32FontCache

	win32FontCacheFree   func(cache *Win32FontCache)
	win32FontCacheLoad   func(cache *Win32FontCache, logfont *W.LOGFONTA) W.HFONT
	win32FontCacheLoadw  func(cache *Win32FontCache, logfont *W.LOGFONTW) W.HFONT
	win32FontCacheUnload func(cache *Win32FontCache, hfont W.HFONT)
)

func (f *Win32FontCache) Free()                             { win32FontCacheFree(f) }
func (f *Win32FontCache) Load(logfont *W.LOGFONTA) W.HFONT  { return win32FontCacheLoad(f, logfont) }
func (f *Win32FontCache) Loadw(logfont *W.LOGFONTW) W.HFONT { return win32FontCacheLoadw(f, logfont) }
func (f *Win32FontCache) Unload(hfont W.HFONT)              { win32FontCacheUnload(f, hfont) }

var dllwin32 = "libpangowin32-1.0-0.dll"

var apiListwin32 = outside.Apis{
	// Undocumented {"_pango_win32_font_get_hfont",&_pango_win32_font_get_hfont},
	// Undocumented {"_pango_win32_font_get_type",&_pango_win32_font_get_type},
	// Undocumented {"_pango_win32_font_map_get_type",&_pango_win32_font_map_get_type},
	// Undocumented {"_pango_win32_fontmap_cache_remove",&_pango_win32_fontmap_cache_remove},
	// Undocumented {"_pango_win32_make_matching_logfontw",&_pango_win32_make_matching_logfontw},
	{"pango_win32_font_cache_free", &win32FontCacheFree},
	{"pango_win32_font_cache_load", &win32FontCacheLoad},
	{"pango_win32_font_cache_loadw", &win32FontCacheLoadw},
	{"pango_win32_font_cache_new", &Win32FontCacheNew},
	{"pango_win32_font_cache_unload", &win32FontCacheUnload},
	{"pango_win32_font_description_from_logfont", &Win32FontDescriptionFromLogfont},
	{"pango_win32_font_description_from_logfontw", &Win32FontDescriptionFromLogfontw},
	{"pango_win32_font_done_font", &win32FontDoneFont},
	{"pango_win32_font_get_glyph_index", &win32FontGetGlyphIndex},
	{"pango_win32_font_get_metrics_factor", &win32FontGetMetricsFactor},
	{"pango_win32_font_logfont", &win32FontLogfont},
	{"pango_win32_font_logfontw", &win32FontLogfontw},
	{"pango_win32_font_map_for_display", &Win32FontMapForDisplay},
	{"pango_win32_font_map_get_font_cache", &Win32FontMapGetFontCache},
	{"pango_win32_font_select_font", &win32FontSelectFont},
	{"pango_win32_get_context", &Win32GetContext},
	{"pango_win32_get_dc", &Win32GetDc},
	{"pango_win32_get_debug_flag", &Win32GetDebugFlag},
	{"pango_win32_get_unknown_glyph", &win32GetUnknownGlyph},
	{"pango_win32_render", &Win32Render},
	{"pango_win32_render_layout", &Win32RenderLayout},
	{"pango_win32_render_layout_line", &Win32RenderLayoutLine},
	{"pango_win32_render_transformed", &Win32RenderTransformed},
	{"pango_win32_shutdown_display", &Win32ShutdownDisplay},
}
