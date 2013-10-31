// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
// T "github.com/tHinqa/outside-gtk2/types"
)

var (
	Win32FontFaceCreateForHfont         func(font HFONT) *FontFace
	Win32FontFaceCreateForLogfontw      func(logfont *LOGFONTW) *FontFace
	Win32FontFaceCreateForLogfontwHfont func(logfont *LOGFONTW, font HFONT) *FontFace
)

func (f HFONT) FontFaceCreate() *FontFace {
	return Win32FontFaceCreateForHfont(f)
}
func (f *LOGFONTW) FontFaceCreate() *FontFace {
	return Win32FontFaceCreateForLogfontw(f)
}
func (f *LOGFONTW) FontFaceCreateHfont(h HFONT) *FontFace {
	return Win32FontFaceCreateForLogfontwHfont(f, h)
}

var (
	Win32PrintingSurfaceCreate func(h HDC) *Surface
	Win32SurfaceCreate         func(h HDC) *Surface
	Win32SurfaceCreateWithDdb  func(h HDC, format Format, width, height int) *Surface
	SurfaceCreateWithDib       func(format Format, width, height int) *Surface
)

func (h HDC) PrintingSurfaceCreate() *Surface { return Win32PrintingSurfaceCreate(h) }
func (h HDC) SurfaceCreate() *Surface         { return Win32SurfaceCreate(h) }
func (h HDC) SurfaceCreateWithDdb(format Format, width, height int) *Surface {
	return Win32SurfaceCreateWithDdb(h, format, width, height)
}

var (
	SurfaceGetDc    func(s *Surface) HDC
	SurfaceGetImage func(s *Surface) *Surface
)
var (
	Win32ScaledFontDoneFont           func(s *ScaledFont)
	Win32ScaledFontGetDeviceToLogical func(s *ScaledFont, deviceToLogical *Matrix)
	Win32ScaledFontGetLogicalToDevice func(s *ScaledFont, logicalToDevice *Matrix)
	Win32ScaledFontGetMetricsFactor   func(s *ScaledFont) float64
	Win32ScaledFontSelectFont         func(s *ScaledFont, hdc HDC) Status
)

func (s *ScaledFont) DoneFont() { Win32ScaledFontDoneFont(s) }
func (s *ScaledFont) GetDeviceToLogical(deviceToLogical *Matrix) {
	Win32ScaledFontGetDeviceToLogical(s, deviceToLogical)
}
func (s *ScaledFont) GetLogicalToDevice(logicalToDevice *Matrix) {
	Win32ScaledFontGetLogicalToDevice(s, logicalToDevice)
}
func (s *ScaledFont) GetMetricsFactor() float64 { return Win32ScaledFontGetMetricsFactor(s) }
func (s *ScaledFont) SelectFont(hdc HDC) Status { return Win32ScaledFontSelectFont(s, hdc) }
