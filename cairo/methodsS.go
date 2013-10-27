// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
	T "github.com/tHinqa/outside-gtk2/types"
)

var (
	ScaledFontCreate func(fontFace *FontFace, fontMatrix *Matrix, ctm *Matrix, options *FontOptions) *ScaledFont

	scaledFontDestroy           func(s *ScaledFont)
	scaledFontExtents           func(s *ScaledFont, extents *FontExtentsS)
	scaledFontGetCtm            func(s *ScaledFont, ctm *Matrix)
	scaledFontGetFontFace       func(s *ScaledFont) *FontFace
	scaledFontGetFontMatrix     func(s *ScaledFont, fontMatrix *Matrix)
	scaledFontGetFontOptions    func(s *ScaledFont, options *FontOptions)
	scaledFontGetReferenceCount func(s *ScaledFont) uint
	scaledFontGetScaleMatrix    func(s *ScaledFont, scaleMatrix *Matrix)
	scaledFontGetType           func(s *ScaledFont) FontType
	scaledFontGetUserData       func(s *ScaledFont, key *UserDataKey) *T.Void
	scaledFontGlyphExtents      func(s *ScaledFont, glyphs *Glyph, numGlyphs int, extents *TextExtents)
	scaledFontReference         func(s *ScaledFont) *ScaledFont
	scaledFontSetUserData       func(s *ScaledFont, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	scaledFontStatus            func(s *ScaledFont) Status
	scaledFontTextExtents       func(s *ScaledFont, utf8 string, extents *TextExtents)
	scaledFontTextToGlyphs      func(s *ScaledFont, x, y float64, utf8 string, utf8Len int, glyphs **Glyph, numGlyphs *int, clusters **TextCluster, numClusters *int, clusterFlags *TextClusterFlags) Status
)

func (s *ScaledFont) Destroy()                             { scaledFontDestroy(s) }
func (s *ScaledFont) Extents(extents *FontExtentsS)        { scaledFontExtents(s, extents) }
func (s *ScaledFont) GetCtm(ctm *Matrix)                   { scaledFontGetCtm(s, ctm) }
func (s *ScaledFont) GetFontFace() *FontFace               { return scaledFontGetFontFace(s) }
func (s *ScaledFont) GetFontMatrix(fontMatrix *Matrix)     { scaledFontGetFontMatrix(s, fontMatrix) }
func (s *ScaledFont) GetFontOptions(options *FontOptions)  { scaledFontGetFontOptions(s, options) }
func (s *ScaledFont) GetReferenceCount() uint              { return scaledFontGetReferenceCount(s) }
func (s *ScaledFont) GetScaleMatrix(scaleMatrix *Matrix)   { scaledFontGetScaleMatrix(s, scaleMatrix) }
func (s *ScaledFont) GetType() FontType                    { return scaledFontGetType(s) }
func (s *ScaledFont) GetUserData(key *UserDataKey) *T.Void { return scaledFontGetUserData(s, key) }
func (s *ScaledFont) GlyphExtents(glyphs *Glyph, numGlyphs int, extents *TextExtents) {
	scaledFontGlyphExtents(s, glyphs, numGlyphs, extents)
}
func (s *ScaledFont) Reference() *ScaledFont { return scaledFontReference(s) }
func (s *ScaledFont) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return scaledFontSetUserData(s, key, userData, destroy)
}
func (s *ScaledFont) Status() Status { return scaledFontStatus(s) }
func (s *ScaledFont) TextExtents(utf8 string, extents *TextExtents) {
	scaledFontTextExtents(s, utf8, extents)
}
func (s *ScaledFont) TextToGlyphs(x, y float64, utf8 string, utf8Len int, glyphs **Glyph, numGlyphs *int, clusters **TextCluster, numClusters *int, clusterFlags *TextClusterFlags) Status {
	return scaledFontTextToGlyphs(s, x, y, utf8, utf8Len, glyphs, numGlyphs, clusters, numClusters, clusterFlags)
}

type ScriptInterpreter struct{}

var (
	ScriptInterpreterCreate func() *ScriptInterpreter

	ScriptInterpreterTranslateStream func(stream *T.FILE, writeFunc WriteFunc, closure *T.Void) Status

	scriptInterpreterInstallHooks  func(ctx *ScriptInterpreter, hooks *ScriptInterpreterHooks)
	scriptInterpreterRun           func(ctx *ScriptInterpreter, filename string) Status
	scriptInterpreterFeedStream    func(ctx *ScriptInterpreter, stream *T.FILE) Status
	scriptInterpreterFeedString    func(ctx *ScriptInterpreter, line string, leng int) Status
	scriptInterpreterGetLineNumber func(ctx *ScriptInterpreter) uint
	scriptInterpreterReference     func(ctx *ScriptInterpreter) *ScriptInterpreter
	scriptInterpreterFinish        func(ctx *ScriptInterpreter) Status
	scriptInterpreterDestroy       func(ctx *ScriptInterpreter) Status
)

func (s *ScriptInterpreter) InstallHooks(hooks *ScriptInterpreterHooks) {
	scriptInterpreterInstallHooks(s, hooks)
}
func (s *ScriptInterpreter) Run(filename string) Status { return scriptInterpreterRun(s, filename) }
func (s *ScriptInterpreter) FeedStream(stream *T.FILE) Status {
	return scriptInterpreterFeedStream(s, stream)
}
func (s *ScriptInterpreter) FeedString(line string, leng int) Status {
	return scriptInterpreterFeedString(s, line, leng)
}
func (s *ScriptInterpreter) GetLineNumber() uint           { return scriptInterpreterGetLineNumber(s) }
func (s *ScriptInterpreter) Reference() *ScriptInterpreter { return scriptInterpreterReference(s) }
func (s *ScriptInterpreter) Finish() Status                { return scriptInterpreterFinish(s) }
func (s *ScriptInterpreter) Destroy() Status               { return scriptInterpreterDestroy(s) }

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
	surfaceCopyPage              func(surface *Surface)
	surfaceCreateForRectangle    func(target *Surface, x, y, width, height float64) *Surface
	surfaceCreateSimilar         func(other *Surface, content Content, width, height int) *Surface
	surfaceDestroy               func(surface *Surface)
	surfaceFinish                func(surface *Surface)
	surfaceFlush                 func(surface *Surface)
	surfaceGetContent            func(surface *Surface) Content
	surfaceGetDevice             func(surface *Surface) *Device
	surfaceGetDeviceOffset       func(surface *Surface, xOffset, yOffset *float64)
	surfaceGetFallbackResolution func(surface *Surface, xPixelsPerInch, yPixelsPerInch *float64)
	surfaceGetFontOptions        func(surface *Surface, options *FontOptions)
	surfaceGetMimeData           func(surface *Surface, mimeType string, data **T.UnsignedChar, length *T.UnsignedLong)
	surfaceGetReferenceCount     func(surface *Surface) uint
	surfaceGetType               func(surface *Surface) SurfaceType
	surfaceGetUserData           func(surface *Surface, key *UserDataKey) *T.Void
	surfaceHasShowTextGlyphs     func(surface *Surface) Bool
	surfaceMarkDirty             func(surface *Surface)
	surfaceMarkDirtyRectangle    func(surface *Surface, x, y, width, height int)
	surfaceReference             func(surface *Surface) *Surface
	surfaceSetDeviceOffset       func(surface *Surface, xOffset, yOffset float64)
	surfaceSetFallbackResolution func(surface *Surface, xPixelsPerInch, yPixelsPerInch float64)
	surfaceSetMimeData           func(surface *Surface, mimeType string, data *T.UnsignedChar, length T.UnsignedLong, destroy DestroyFunc, closure *T.Void) Status
	surfaceSetUserData           func(surface *Surface, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	surfaceShowPage              func(surface *Surface)
	surfaceStatus                func(surface *Surface) Status
	surfaceWriteToPng            func(surface *Surface, filename string) Status
	surfaceWriteToPngStream      func(surface *Surface, writeFunc WriteFunc, closure *T.Void) Status

	ImageSurfaceCreate              func(format Format, width, height int) *Surface
	ImageSurfaceCreateForData       func(data *T.UnsignedChar, format Format, width, height, stride int) *Surface
	ImageSurfaceCreateFromPng       func(filename string) *Surface
	ImageSurfaceCreateFromPngStream func(readFunc ReadFunc, closure *T.Void) *Surface
	RecordingSurfaceCreate          func(content Content, extents *Rectangle) *Surface

	imageSurfaceGetData        func(surface *Surface) *T.UnsignedChar
	imageSurfaceGetFormat      func(surface *Surface) Format
	imageSurfaceGetHeight      func(surface *Surface) int
	imageSurfaceGetStride      func(surface *Surface) int
	imageSurfaceGetWidth       func(surface *Surface) int
	recordingSurfaceInkExtents func(surface *Surface, x0, y0, width, height *float64)
)

func (s *Surface) CopyPage() { surfaceCopyPage(s) }
func (s *Surface) CreateForRectangle(x, y, width, height float64) *Surface {
	return surfaceCreateForRectangle(s, x, y, width, height)
}
func (s *Surface) CreateSimilar(content Content, width, height int) *Surface {
	return surfaceCreateSimilar(s, content, width, height)
}
func (s *Surface) Destroy()            { surfaceDestroy(s) }
func (s *Surface) Finish()             { surfaceFinish(s) }
func (s *Surface) Flush()              { surfaceFlush(s) }
func (s *Surface) GetContent() Content { return surfaceGetContent(s) }
func (s *Surface) GetDevice() *Device  { return surfaceGetDevice(s) }
func (s *Surface) GetDeviceOffset(xOffset, yOffset *float64) {
	surfaceGetDeviceOffset(s, xOffset, yOffset)
}
func (s *Surface) GetFallbackResolution(xPixelsPerInch, yPixelsPerInch *float64) {
	surfaceGetFallbackResolution(s, xPixelsPerInch, yPixelsPerInch)
}
func (s *Surface) GetFontOptions(options *FontOptions) { surfaceGetFontOptions(s, options) }
func (s *Surface) GetMimeData(mimeType string, data **T.UnsignedChar, length *T.UnsignedLong) {
	surfaceGetMimeData(s, mimeType, data, length)
}
func (s *Surface) GetReferenceCount() uint              { return surfaceGetReferenceCount(s) }
func (s *Surface) GetType() SurfaceType                 { return surfaceGetType(s) }
func (s *Surface) GetUserData(key *UserDataKey) *T.Void { return surfaceGetUserData(s, key) }
func (s *Surface) HasShowTextGlyphs() Bool              { return surfaceHasShowTextGlyphs(s) }
func (s *Surface) MarkDirty()                           { surfaceMarkDirty(s) }
func (s *Surface) MarkDirtyRectangle(x, y, width, height int) {
	surfaceMarkDirtyRectangle(s, x, y, width, height)
}
func (s *Surface) Reference() *Surface { return surfaceReference(s) }
func (s *Surface) SetDeviceOffset(xOffset, yOffset float64) {
	surfaceSetDeviceOffset(s, xOffset, yOffset)
}
func (s *Surface) SetFallbackResolution(xPixelsPerInch, yPixelsPerInch float64) {
	surfaceSetFallbackResolution(s, xPixelsPerInch, yPixelsPerInch)
}
func (s *Surface) SetMimeData(mimeType string, data *T.UnsignedChar, length T.UnsignedLong, destroy DestroyFunc, closure *T.Void) Status {
	return surfaceSetMimeData(s, mimeType, data, length, destroy, closure)
}
func (s *Surface) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return surfaceSetUserData(s, key, userData, destroy)
}
func (s *Surface) ShowPage()                         { surfaceShowPage(s) }
func (s *Surface) Status() Status                    { return surfaceStatus(s) }
func (s *Surface) WriteToPng(filename string) Status { return surfaceWriteToPng(s, filename) }
func (s *Surface) WriteToPngStream(writeFunc WriteFunc, closure *T.Void) Status {
	return surfaceWriteToPngStream(s, writeFunc, closure)
}

func (s *Surface) GetData() *T.UnsignedChar { return imageSurfaceGetData(s) }
func (s *Surface) GetFormat() Format        { return imageSurfaceGetFormat(s) }
func (s *Surface) GetHeight() int           { return imageSurfaceGetHeight(s) }
func (s *Surface) GetStride() int           { return imageSurfaceGetStride(s) }
func (s *Surface) GetWidth() int            { return imageSurfaceGetWidth(s) }
func (s *Surface) InkExtents(x0, y0, width, height *float64) {
	recordingSurfaceInkExtents(s, x0, y0, width, height)
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
