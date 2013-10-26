// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
	T "github.com/tHinqa/outside-gtk2/types"
)

var (
	ScaledFontCreate func(fontFace *FontFace, fontMatrix *Matrix, ctm *Matrix, options *FontOptions) *ScaledFont

	ScaledFontReference         func(s *ScaledFont) *ScaledFont
	ScaledFontDestroy           func(s *ScaledFont)
	ScaledFontGetReferenceCount func(s *ScaledFont) uint
	ScaledFontStatus            func(s *ScaledFont) Status
	ScaledFontGetType           func(s *ScaledFont) FontType
	ScaledFontGetUserData       func(s *ScaledFont, key *UserDataKey) *T.Void
	ScaledFontSetUserData       func(s *ScaledFont, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	ScaledFontExtents           func(s *ScaledFont, extents *FontExtentsS)
	ScaledFontTextExtents       func(s *ScaledFont, utf8 string, extents *TextExtents)
	ScaledFontGlyphExtents      func(s *ScaledFont, glyphs *Glyph, numGlyphs int, extents *TextExtents)
	ScaledFontTextToGlyphs      func(s *ScaledFont, x, y float64, utf8 string, utf8Len int, glyphs **Glyph, numGlyphs *int, clusters **TextCluster, numClusters *int, clusterFlags *TextClusterFlags) Status
	ScaledFontGetFontFace       func(s *ScaledFont) *FontFace
	ScaledFontGetFontMatrix     func(s *ScaledFont, fontMatrix *Matrix)
	ScaledFontGetCtm            func(s *ScaledFont, ctm *Matrix)
	ScaledFontGetScaleMatrix    func(s *ScaledFont, scaleMatrix *Matrix)
	ScaledFontGetFontOptions    func(s *ScaledFont, options *FontOptions)
)

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
