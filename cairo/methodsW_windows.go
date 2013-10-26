// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
// T "github.com/tHinqa/outside-gtk2/types"
)

var (
	win32FontFaceCreateForHfont         func(font HFONT) *FontFace
	win32FontFaceCreateForLogfontw      func(logfont *LOGFONTW) *FontFace
	win32FontFaceCreateForLogfontwHfont func(logfont *LOGFONTW, font HFONT) *FontFace
)

func (f HFONT) FontFaceCreate() *FontFace {
	return win32FontFaceCreateForHfont(f)
}
func (f *LOGFONTW) FontFaceCreate() *FontFace {
	return win32FontFaceCreateForLogfontw(f)
}
func (f *LOGFONTW) FontFaceCreateHfont(h HFONT) *FontFace {
	return win32FontFaceCreateForLogfontwHfont(f, h)
}

var (
	win32PrintingSurfaceCreate func(h HDC) *Surface
	win32SurfaceCreate         func(h HDC) *Surface
	win32SurfaceCreateWithDdb  func(h HDC, format Format, width, height int) *Surface
	SurfaceCreateWithDib       func(format Format, width, height int) *Surface
)

func (h HDC) PrintingSurfaceCreate() *Surface { return win32PrintingSurfaceCreate(h) }
func (h HDC) SurfaceCreate() *Surface         { return win32SurfaceCreate(h) }
func (h HDC) SurfaceCreateWithDdb(format Format, width, height int) *Surface {
	return win32SurfaceCreateWithDdb(h, format, width, height)
}

var (
	SurfaceGetDc    func(s *Surface) HDC
	SurfaceGetImage func(s *Surface) *Surface
)
var (
	win32ScaledFontDoneFont           func(s *ScaledFont)
	win32ScaledFontGetDeviceToLogical func(s *ScaledFont, deviceToLogical *Matrix)
	win32ScaledFontGetLogicalToDevice func(s *ScaledFont, logicalToDevice *Matrix)
	win32ScaledFontGetMetricsFactor   func(s *ScaledFont) float64
	win32ScaledFontSelectFont         func(s *ScaledFont, hdc HDC) Status
)

func (s *ScaledFont) DoneFont() { win32ScaledFontDoneFont(s) }
func (s *ScaledFont) GetDeviceToLogical(deviceToLogical *Matrix) {
	win32ScaledFontGetDeviceToLogical(s, deviceToLogical)
}
func (s *ScaledFont) GetLogicalToDevice(logicalToDevice *Matrix) {
	win32ScaledFontGetLogicalToDevice(s, logicalToDevice)
}
func (s *ScaledFont) GetMetricsFactor() float64 { return win32ScaledFontGetMetricsFactor(s) }
func (s *ScaledFont) SelectFont(hdc HDC) Status { return win32ScaledFontSelectFont(s, hdc) }
