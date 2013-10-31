// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
	T "github.com/tHinqa/outside-gtk2/types"
)

var (
	ScaledFontCreate func(fontFace *FontFace, fontMatrix *Matrix, ctm *Matrix, options *FontOptions) *ScaledFont

	ScaledFontDestroy           func(s *ScaledFont)
	ScaledFontExtents           func(s *ScaledFont, extents *FontExtentsS)
	ScaledFontGetCtm            func(s *ScaledFont, ctm *Matrix)
	ScaledFontGetFontFace       func(s *ScaledFont) *FontFace
	ScaledFontGetFontMatrix     func(s *ScaledFont, fontMatrix *Matrix)
	ScaledFontGetFontOptions    func(s *ScaledFont, options *FontOptions)
	ScaledFontGetReferenceCount func(s *ScaledFont) uint
	ScaledFontGetScaleMatrix    func(s *ScaledFont, scaleMatrix *Matrix)
	ScaledFontGetType           func(s *ScaledFont) FontType
	ScaledFontGetUserData       func(s *ScaledFont, key *UserDataKey) *T.Void
	ScaledFontGlyphExtents      func(s *ScaledFont, glyphs *Glyph, numGlyphs int, extents *TextExtents)
	ScaledFontReference         func(s *ScaledFont) *ScaledFont
	ScaledFontSetUserData       func(s *ScaledFont, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	ScaledFontStatus            func(s *ScaledFont) Status
	ScaledFontTextExtents       func(s *ScaledFont, utf8 string, extents *TextExtents)
	ScaledFontTextToGlyphs      func(s *ScaledFont, x, y float64, utf8 string, utf8Len int, glyphs **Glyph, numGlyphs *int, clusters **TextCluster, numClusters *int, clusterFlags *TextClusterFlags) Status
)

func (s *ScaledFont) Destroy()                             { ScaledFontDestroy(s) }
func (s *ScaledFont) Extents(extents *FontExtentsS)        { ScaledFontExtents(s, extents) }
func (s *ScaledFont) GetCtm(ctm *Matrix)                   { ScaledFontGetCtm(s, ctm) }
func (s *ScaledFont) GetFontFace() *FontFace               { return ScaledFontGetFontFace(s) }
func (s *ScaledFont) GetFontMatrix(fontMatrix *Matrix)     { ScaledFontGetFontMatrix(s, fontMatrix) }
func (s *ScaledFont) GetFontOptions(options *FontOptions)  { ScaledFontGetFontOptions(s, options) }
func (s *ScaledFont) GetReferenceCount() uint              { return ScaledFontGetReferenceCount(s) }
func (s *ScaledFont) GetScaleMatrix(scaleMatrix *Matrix)   { ScaledFontGetScaleMatrix(s, scaleMatrix) }
func (s *ScaledFont) GetType() FontType                    { return ScaledFontGetType(s) }
func (s *ScaledFont) GetUserData(key *UserDataKey) *T.Void { return ScaledFontGetUserData(s, key) }
func (s *ScaledFont) GlyphExtents(glyphs *Glyph, numGlyphs int, extents *TextExtents) {
	ScaledFontGlyphExtents(s, glyphs, numGlyphs, extents)
}
func (s *ScaledFont) Reference() *ScaledFont { return ScaledFontReference(s) }
func (s *ScaledFont) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return ScaledFontSetUserData(s, key, userData, destroy)
}
func (s *ScaledFont) Status() Status { return ScaledFontStatus(s) }
func (s *ScaledFont) TextExtents(utf8 string, extents *TextExtents) {
	ScaledFontTextExtents(s, utf8, extents)
}
func (s *ScaledFont) TextToGlyphs(x, y float64, utf8 string, utf8Len int, glyphs **Glyph, numGlyphs *int, clusters **TextCluster, numClusters *int, clusterFlags *TextClusterFlags) Status {
	return ScaledFontTextToGlyphs(s, x, y, utf8, utf8Len, glyphs, numGlyphs, clusters, numClusters, clusterFlags)
}

type ScriptInterpreter struct{}

var (
	ScriptInterpreterCreate func() *ScriptInterpreter

	ScriptInterpreterTranslateStream func(stream *T.FILE, writeFunc WriteFunc, closure *T.Void) Status

	ScriptInterpreterInstallHooks  func(ctx *ScriptInterpreter, hooks *ScriptInterpreterHooks)
	ScriptInterpreterRun           func(ctx *ScriptInterpreter, filename string) Status
	ScriptInterpreterFeedStream    func(ctx *ScriptInterpreter, stream *T.FILE) Status
	ScriptInterpreterFeedString    func(ctx *ScriptInterpreter, line string, leng int) Status
	ScriptInterpreterGetLineNumber func(ctx *ScriptInterpreter) uint
	ScriptInterpreterReference     func(ctx *ScriptInterpreter) *ScriptInterpreter
	ScriptInterpreterFinish        func(ctx *ScriptInterpreter) Status
	ScriptInterpreterDestroy       func(ctx *ScriptInterpreter) Status
)

func (s *ScriptInterpreter) InstallHooks(hooks *ScriptInterpreterHooks) {
	ScriptInterpreterInstallHooks(s, hooks)
}
func (s *ScriptInterpreter) Run(filename string) Status { return ScriptInterpreterRun(s, filename) }
func (s *ScriptInterpreter) FeedStream(stream *T.FILE) Status {
	return ScriptInterpreterFeedStream(s, stream)
}
func (s *ScriptInterpreter) FeedString(line string, leng int) Status {
	return ScriptInterpreterFeedString(s, line, leng)
}
func (s *ScriptInterpreter) GetLineNumber() uint           { return ScriptInterpreterGetLineNumber(s) }
func (s *ScriptInterpreter) Reference() *ScriptInterpreter { return ScriptInterpreterReference(s) }
func (s *ScriptInterpreter) Finish() Status                { return ScriptInterpreterFinish(s) }
func (s *ScriptInterpreter) Destroy() Status               { return ScriptInterpreterDestroy(s) }

type (
	ScriptInterpreterHooks struct {
		Closure        *T.Void
		SurfaceCreate  CsiSurfaceCreateFunc
		SurfaceDestroy CsiDestroyFunc
		ContextCreate  CsiContextCreateFunc
		ContextDestroy CsiDestroyFunc
		ShowPage       CsiShowPageFunc
		CopyPage       CsiCopyPageFunc
	}

	CsiSurfaceCreateFunc func(
		closure *T.Void, content Content, width, height float64,
		uid T.Long) *Surface

	CsiDestroyFunc func(closure, ptr *T.Void)

	CsiContextCreateFunc func(
		closure *T.Void, surface *Surface) *Cairo

	CsiShowPageFunc func(closure *T.Void, cr *Cairo)

	CsiCopyPageFunc func(closure *T.Void, cr *Cairo)
)

type Content Enum

const (
	CONTENT_COLOR       Content = 0x1000
	CONTENT_ALPHA       Content = 0x2000
	CONTENT_COLOR_ALPHA Content = 0x3000
)

var (
	SurfaceCopyPage              func(surface *Surface)
	SurfaceCreateForRectangle    func(target *Surface, x, y, width, height float64) *Surface
	SurfaceCreateSimilar         func(other *Surface, content Content, width, height int) *Surface
	SurfaceDestroy               func(surface *Surface)
	SurfaceFinish                func(surface *Surface)
	SurfaceFlush                 func(surface *Surface)
	SurfaceGetContent            func(surface *Surface) Content
	SurfaceGetDevice             func(surface *Surface) *Device
	SurfaceGetDeviceOffset       func(surface *Surface, xOffset, yOffset *float64)
	SurfaceGetFallbackResolution func(surface *Surface, xPixelsPerInch, yPixelsPerInch *float64)
	SurfaceGetFontOptions        func(surface *Surface, options *FontOptions)
	SurfaceGetMimeData           func(surface *Surface, mimeType string, data **T.UnsignedChar, length *T.UnsignedLong)
	SurfaceGetReferenceCount     func(surface *Surface) uint
	SurfaceGetType               func(surface *Surface) SurfaceType
	SurfaceGetUserData           func(surface *Surface, key *UserDataKey) *T.Void
	SurfaceHasShowTextGlyphs     func(surface *Surface) Bool
	SurfaceMarkDirty             func(surface *Surface)
	SurfaceMarkDirtyRectangle    func(surface *Surface, x, y, width, height int)
	SurfaceReference             func(surface *Surface) *Surface
	SurfaceSetDeviceOffset       func(surface *Surface, xOffset, yOffset float64)
	SurfaceSetFallbackResolution func(surface *Surface, xPixelsPerInch, yPixelsPerInch float64)
	SurfaceSetMimeData           func(surface *Surface, mimeType string, data *T.UnsignedChar, length T.UnsignedLong, destroy DestroyFunc, closure *T.Void) Status
	SurfaceSetUserData           func(surface *Surface, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	SurfaceShowPage              func(surface *Surface)
	SurfaceStatus                func(surface *Surface) Status
	SurfaceWriteToPng            func(surface *Surface, filename string) Status
	SurfaceWriteToPngStream      func(surface *Surface, writeFunc WriteFunc, closure *T.Void) Status

	ImageSurfaceCreate              func(format Format, width, height int) *Surface
	ImageSurfaceCreateForData       func(data *T.UnsignedChar, format Format, width, height, stride int) *Surface
	ImageSurfaceCreateFromPng       func(filename string) *Surface
	ImageSurfaceCreateFromPngStream func(readFunc ReadFunc, closure *T.Void) *Surface
	RecordingSurfaceCreate          func(content Content, extents *Rectangle) *Surface

	ImageSurfaceGetData        func(surface *Surface) *T.UnsignedChar
	ImageSurfaceGetFormat      func(surface *Surface) Format
	ImageSurfaceGetHeight      func(surface *Surface) int
	ImageSurfaceGetStride      func(surface *Surface) int
	ImageSurfaceGetWidth       func(surface *Surface) int
	RecordingSurfaceInkExtents func(surface *Surface, x0, y0, width, height *float64)
)

func (s *Surface) CopyPage() { SurfaceCopyPage(s) }
func (s *Surface) CreateForRectangle(x, y, width, height float64) *Surface {
	return SurfaceCreateForRectangle(s, x, y, width, height)
}
func (s *Surface) CreateSimilar(content Content, width, height int) *Surface {
	return SurfaceCreateSimilar(s, content, width, height)
}
func (s *Surface) Destroy()            { SurfaceDestroy(s) }
func (s *Surface) Finish()             { SurfaceFinish(s) }
func (s *Surface) Flush()              { SurfaceFlush(s) }
func (s *Surface) GetContent() Content { return SurfaceGetContent(s) }
func (s *Surface) GetDevice() *Device  { return SurfaceGetDevice(s) }
func (s *Surface) GetDeviceOffset(xOffset, yOffset *float64) {
	SurfaceGetDeviceOffset(s, xOffset, yOffset)
}
func (s *Surface) GetFallbackResolution(xPixelsPerInch, yPixelsPerInch *float64) {
	SurfaceGetFallbackResolution(s, xPixelsPerInch, yPixelsPerInch)
}
func (s *Surface) GetFontOptions(options *FontOptions) { SurfaceGetFontOptions(s, options) }
func (s *Surface) GetMimeData(mimeType string, data **T.UnsignedChar, length *T.UnsignedLong) {
	SurfaceGetMimeData(s, mimeType, data, length)
}
func (s *Surface) GetReferenceCount() uint              { return SurfaceGetReferenceCount(s) }
func (s *Surface) GetType() SurfaceType                 { return SurfaceGetType(s) }
func (s *Surface) GetUserData(key *UserDataKey) *T.Void { return SurfaceGetUserData(s, key) }
func (s *Surface) HasShowTextGlyphs() Bool              { return SurfaceHasShowTextGlyphs(s) }
func (s *Surface) MarkDirty()                           { SurfaceMarkDirty(s) }
func (s *Surface) MarkDirtyRectangle(x, y, width, height int) {
	SurfaceMarkDirtyRectangle(s, x, y, width, height)
}
func (s *Surface) Reference() *Surface { return SurfaceReference(s) }
func (s *Surface) SetDeviceOffset(xOffset, yOffset float64) {
	SurfaceSetDeviceOffset(s, xOffset, yOffset)
}
func (s *Surface) SetFallbackResolution(xPixelsPerInch, yPixelsPerInch float64) {
	SurfaceSetFallbackResolution(s, xPixelsPerInch, yPixelsPerInch)
}
func (s *Surface) SetMimeData(mimeType string, data *T.UnsignedChar, length T.UnsignedLong, destroy DestroyFunc, closure *T.Void) Status {
	return SurfaceSetMimeData(s, mimeType, data, length, destroy, closure)
}
func (s *Surface) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return SurfaceSetUserData(s, key, userData, destroy)
}
func (s *Surface) ShowPage()                         { SurfaceShowPage(s) }
func (s *Surface) Status() Status                    { return SurfaceStatus(s) }
func (s *Surface) WriteToPng(filename string) Status { return SurfaceWriteToPng(s, filename) }
func (s *Surface) WriteToPngStream(writeFunc WriteFunc, closure *T.Void) Status {
	return SurfaceWriteToPngStream(s, writeFunc, closure)
}

func (s *Surface) GetData() *T.UnsignedChar { return ImageSurfaceGetData(s) }
func (s *Surface) GetFormat() Format        { return ImageSurfaceGetFormat(s) }
func (s *Surface) GetHeight() int           { return ImageSurfaceGetHeight(s) }
func (s *Surface) GetStride() int           { return ImageSurfaceGetStride(s) }
func (s *Surface) GetWidth() int            { return ImageSurfaceGetWidth(s) }
func (s *Surface) InkExtents(x0, y0, width, height *float64) {
	RecordingSurfaceInkExtents(s, x0, y0, width, height)
}

type SurfaceType Enum

const (
	SURFACE_TYPE_IMAGE SurfaceType = iota
	SURFACE_TYPE_PDF
	SURFACE_TYPE_PS
	SURFACE_TYPE_XLIB
	SURFACE_TYPE_XCB
	SURFACE_TYPE_GLITZ
	SURFACE_TYPE_QUARTZ
	SURFACE_TYPE_WIN32
	SURFACE_TYPE_BEOS
	SURFACE_TYPE_DIRECTFB
	SURFACE_TYPE_SVG
	SURFACE_TYPE_OS2
	SURFACE_TYPE_WIN32_PRINTING
	SURFACE_TYPE_QUARTZ_IMAGE
	SURFACE_TYPE_SCRIPT
	SURFACE_TYPE_QT
	SURFACE_TYPE_RECORDING
	SURFACE_TYPE_VG
	SURFACE_TYPE_GL
	SURFACE_TYPE_DRM
	SURFACE_TYPE_TEE
	SURFACE_TYPE_XML
	SURFACE_TYPE_SKIA
	SURFACE_TYPE_SUBSURFACE
)

var (
	SvgGetVersions              func(versions **SvgVersion, numVersions *int)
	SvgSurfaceCreate            func(filename string, widthInPoints float64, heightInPoints float64) *Surface
	SvgSurfaceCreateForStream   func(writeFunc WriteFunc, closure *T.Void, widthInPoints float64, heightInPoints float64) *Surface
	SvgSurfaceRestrictToVersion func(surface *Surface, version SvgVersion)
	SvgVersionToString          func(version SvgVersion) string
)

type Status Enum

const (
	STATUS_SUCCESS Status = iota
	STATUS_NO_MEMORY
	STATUS_INVALID_RESTORE
	STATUS_INVALID_POP_GROUP
	STATUS_NO_CURRENT_POINT
	STATUS_INVALID_MATRIX
	STATUS_INVALID_STATUS
	STATUS_NULL_POINTER
	STATUS_INVALID_STRING
	STATUS_INVALID_PATH_DATA
	STATUS_READ_ERROR
	STATUS_WRITE_ERROR
	STATUS_SURFACE_FINISHED
	STATUS_SURFACE_TYPE_MISMATCH
	STATUS_PATTERN_TYPE_MISMATCH
	STATUS_INVALID_CONTENT
	STATUS_INVALID_FORMAT
	STATUS_INVALID_VISUAL
	STATUS_FILE_NOT_FOUND
	STATUS_INVALID_DASH
	STATUS_INVALID_DSC_COMMENT
	STATUS_INVALID_INDEX
	STATUS_CLIP_NOT_REPRESENTABLE
	STATUS_TEMP_FILE_ERROR
	STATUS_INVALID_STRIDE
	STATUS_FONT_TYPE_MISMATCH
	STATUS_USER_FONT_IMMUTABLE
	STATUS_USER_FONT_ERROR
	STATUS_NEGATIVE_COUNT
	STATUS_INVALID_CLUSTERS
	STATUS_INVALID_SLANT
	STATUS_INVALID_WEIGHT
	STATUS_INVALID_SIZE
	STATUS_USER_FONT_NOT_IMPLEMENTED
	STATUS_DEVICE_TYPE_MISMATCH
	STATUS_DEVICE_ERROR
	STATUS_LAST_STATUS
)

var StatusToString func(status Status) string

type SvgVersion Enum

const (
	SVG_VERSION_1_1 SvgVersion = iota
	SVG_VERSION_1_2
)
