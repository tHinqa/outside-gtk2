// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package cairo provides API definitions for accessing
//libcairo-2.dll, libcairo-gobject-2.dll and
//libcairo-script-interpreter-2.dll.
package cairo

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllApis(dllGobject, false, apiListGobject)
	outside.AddDllApis(dllScript, false, apiListScript)
}

type (
	//TODO(t):Fix
	HDC      uint32
	HFONT    uint32
	LOGFONTW struct{}
)

type PdfVersion T.Enum

const (
	PDF_VERSION_1_4 PdfVersion = iota
	PDF_VERSION_1_5
)

type PsLevel T.Enum

const (
	PS_LEVEL_2 PsLevel = iota
	PS_LEVEL_3
)

type SvgVersion T.Enum

const (
	SVG_VERSION_1_1 SvgVersion = iota
	SVG_VERSION_1_2
)

var (
	Version func() int

	VersionString func() string

	Create func(target *T.CairoSurface) *T.Cairo

	Reference func(cr *T.Cairo) *T.Cairo

	Destroy func(cr *T.Cairo)

	GetReferenceCount func(cr *T.Cairo) uint

	GetUserData func(
		cr *T.Cairo, key *T.CairoUserDataKey) *T.Void

	SetUserData func(
		cr *T.Cairo,
		key *T.CairoUserDataKey,
		userData *T.Void,
		destroy T.CairoDestroyFunc) T.CairoStatus

	Save func(cr *T.Cairo)

	Restore func(cr *T.Cairo)

	PushGroup func(cr *T.Cairo)

	PushGroupWithContent func(
		cr *T.Cairo, content T.CairoContent)

	PopGroup func(cr *T.Cairo) *T.CairoPattern

	PopGroupToSource func(cr *T.Cairo)

	SetOperator func(cr *T.Cairo, op T.CairoOperator)

	SetSource func(cr *T.Cairo, source *T.CairoPattern)

	SetSourceRgb func(cr *T.Cairo, red, green, blue float64)

	SetSourceRgba func(
		cr *T.Cairo, red, green, blue, alpha float64)

	SetSourceSurface func(
		cr *T.Cairo, surface *T.CairoSurface, x, y float64)

	SetTolerance func(cr *T.Cairo, tolerance float64)

	SetAntialias func(cr *T.Cairo, antialias T.CairoAntialias)

	SetFillRule func(cr *T.Cairo, fillRule T.CairoFillRule)

	SetLineWidth func(cr *T.Cairo, width float64)

	SetLineCap func(cr *T.Cairo, lineCap T.CairoLineCap)

	SetLineJoin func(cr *T.Cairo, lineJoin T.CairoLineJoin)

	SetDash func(cr *T.Cairo,
		dashes *float64, numDashes int, offset float64)

	SetMiterLimit func(cr *T.Cairo, limit float64)

	Translate func(cr *T.Cairo, tx, ty float64)

	Scale func(cr *T.Cairo, sx, sy float64)

	Rotate func(cr *T.Cairo, angle float64)

	Transform func(cr *T.Cairo, matrix *T.CairoMatrix)

	SetMatrix func(cr *T.Cairo, matrix *T.CairoMatrix)

	IdentityMatrix func(cr *T.Cairo)

	UserToDevice func(cr *T.Cairo, x, y *float64)

	UserToDeviceDistance func(cr *T.Cairo, dx, dy *float64)

	DeviceToUser func(cr *T.Cairo, x, y *float64)

	DeviceToUserDistance func(cr *T.Cairo, dx, dy *float64)

	NewPath func(cr *T.Cairo)

	MoveTo func(cr *T.Cairo, x, y float64)

	NewSubPath func(cr *T.Cairo)

	LineTo func(cr *T.Cairo, x, y float64)

	CurveTo func(cr *T.Cairo, x1, y1, x2, y2, x3, y3 float64)

	Arc func(
		cr *T.Cairo, xc, yc, radius, angle1, angle2 float64)

	ArcNegative func(
		cr *T.Cairo, xc, yc, radius, angle1, angle2 float64)

	RelMoveTo func(cr *T.Cairo, dx, dy float64)

	RelLineTo func(cr *T.Cairo, dx, dy float64)

	RelCurveTo func(
		cr *T.Cairo, dx1, dy1, dx2, dy2, dx3, dy3 float64)

	Rectangle func(cr *T.Cairo, x, y, width, height float64)

	ClosePath func(cr *T.Cairo)

	PathExtents func(cr *T.Cairo, x1, y1, x2, y2 *float64)

	Paint func(cr *T.Cairo)

	PaintWithAlpha func(cr *T.Cairo, alpha float64)

	Mask func(cr *T.Cairo, pattern *T.CairoPattern)

	MaskSurface func(cr *T.Cairo,
		surface *T.CairoSurface, surfaceX, surfaceY float64)

	Stroke func(cr *T.Cairo)

	StrokePreserve func(cr *T.Cairo)

	Fill func(cr *T.Cairo)

	FillPreserve func(cr *T.Cairo)

	CopyPage func(cr *T.Cairo)

	ShowPage func(cr *T.Cairo)

	InStroke func(cr *T.Cairo, x, y float64) T.CairoBool

	InFill func(cr *T.Cairo, x, y float64) T.CairoBool

	InClip func(cr *T.Cairo, x, y float64) T.CairoBool

	StrokeExtents func(cr *T.Cairo, x1, y1, x2, y2 *float64)

	FillExtents func(cr *T.Cairo, x1, y1, x2, y2 *float64)

	ResetClip func(cr *T.Cairo)

	Clip func(cr *T.Cairo)

	ClipPreserve func(cr *T.Cairo)

	ClipExtents func(cr *T.Cairo, x1, y1, x2, y2 *float64)

	CopyClipRectangleList func(
		cr *T.Cairo) *T.CairoRectangleList

	RectangleListDestroy func(
		rectangleList *T.CairoRectangleList)

	GlyphAllocate func(numGlyphs int) *T.CairoGlyph

	GlyphFree func(
		glyphs *T.CairoGlyph)

	TextClusterAllocate func(
		numClusters int) *T.CairoTextCluster

	TextClusterFree func(
		clusters *T.CairoTextCluster)

	FontOptionsCreate func() *T.CairoFontOptions

	FontOptionsCopy func(
		original *T.CairoFontOptions) *T.CairoFontOptions

	FontOptionsDestroy func(options *T.CairoFontOptions)

	FontOptionsStatus func(
		options *T.CairoFontOptions) T.CairoStatus

	FontOptionsMerge func(
		options *T.CairoFontOptions, other *T.CairoFontOptions)

	FontOptionsEqual func(options *T.CairoFontOptions,
		other *T.CairoFontOptions) T.CairoBool

	FontOptionsHash func(
		options *T.CairoFontOptions) T.UnsignedLong

	FontOptionsSetAntialias func(
		options *T.CairoFontOptions, antialias T.CairoAntialias)

	FontOptionsGetAntialias func(
		options *T.CairoFontOptions) T.CairoAntialias

	FontOptionsSetSubpixelOrder func(options *T.CairoFontOptions,
		subpixelOrder T.CairoSubpixelOrder)

	FontOptionsGetSubpixelOrder func(
		options *T.CairoFontOptions) T.CairoSubpixelOrder

	FontOptionsSetHintStyle func(
		options *T.CairoFontOptions, hintStyle T.CairoHintStyle)

	FontOptionsGetHintStyle func(
		options *T.CairoFontOptions) T.CairoHintStyle

	FontOptionsSetHintMetrics func(options *T.CairoFontOptions,
		hintMetrics T.CairoHintMetrics)

	FontOptionsGetHintMetrics func(
		options *T.CairoFontOptions) T.CairoHintMetrics

	SelectFontFace func(cr *T.Cairo, family string,
		slant T.CairoFontSlant, weight T.CairoFontWeight)

	SetFontSize func(cr *T.Cairo, size float64)

	SetFontMatrix func(cr *T.Cairo, matrix *T.CairoMatrix)

	GetFontMatrix func(cr *T.Cairo, matrix *T.CairoMatrix)

	SetFontOptions func(cr *T.Cairo, options *T.CairoFontOptions)

	GetFontOptions func(cr *T.Cairo, options *T.CairoFontOptions)

	SetFontFace func(cr *T.Cairo, fontFace *T.CairoFontFace)

	GetFontFace func(cr *T.Cairo) *T.CairoFontFace

	SetScaledFont func(
		cr *T.Cairo, scaledFont *T.CairoScaledFont)

	GetScaledFont func(cr *T.Cairo) *T.CairoScaledFont

	ShowText func(cr *T.Cairo, utf8 string)

	ShowGlyphs func(
		cr *T.Cairo, glyphs *T.CairoGlyph, numGlyphs int)

	ShowTextGlyphs func(cr *T.Cairo, utf8 string, utf8Len int,
		glyphs *T.CairoGlyph, numGlyphs int,
		clusters *T.CairoTextCluster, numClusters int,
		clusterFlags T.CairoTextClusterFlags)

	TextPath func(cr *T.Cairo, utf8 string)

	GlyphPath func(
		cr *T.Cairo, glyphs *T.CairoGlyph, numGlyphs int)

	TextExtents func(
		cr *T.Cairo, utf8 string, extents *T.CairoTextExtents)

	GlyphExtents func(cr *T.Cairo, glyphs *T.CairoGlyph,
		numGlyphs int, extents *T.CairoTextExtents)

	FontExtents func(cr *T.Cairo, extents *T.CairoFontExtents)

	FontFaceReference func(
		fontFace *T.CairoFontFace) *T.CairoFontFace

	FontFaceDestroy func(fontFace *T.CairoFontFace)

	FontFaceGetReferenceCount func(
		fontFace *T.CairoFontFace) uint

	FontFaceStatus func(fontFace *T.CairoFontFace) T.CairoStatus

	FontFaceGetType func(
		fontFace *T.CairoFontFace) T.CairoFontType

	FontFaceGetUserData func(fontFace *T.CairoFontFace,
		key *T.CairoUserDataKey) *T.Void

	FontFaceSetUserData func(fontFace *T.CairoFontFace,
		key *T.CairoUserDataKey, userData *T.Void,
		destroy T.CairoDestroyFunc) T.CairoStatus

	ScaledFontCreate func(fontFace *T.CairoFontFace,
		fontMatrix *T.CairoMatrix, ctm *T.CairoMatrix,
		options *T.CairoFontOptions) *T.CairoScaledFont

	ScaledFontReference func(
		scaledFont *T.CairoScaledFont) *T.CairoScaledFont

	ScaledFontDestroy func(scaledFont *T.CairoScaledFont)

	ScaledFontGetReferenceCount func(
		scaledFont *T.CairoScaledFont) uint

	ScaledFontStatus func(
		scaledFont *T.CairoScaledFont) T.CairoStatus

	ScaledFontGetType func(
		scaledFont *T.CairoScaledFont) T.CairoFontType

	ScaledFontGetUserData func(scaledFont *T.CairoScaledFont,
		key *T.CairoUserDataKey) *T.Void

	ScaledFontSetUserData func(scaledFont *T.CairoScaledFont,
		key *T.CairoUserDataKey, userData *T.Void,
		destroy T.CairoDestroyFunc) T.CairoStatus

	ScaledFontExtents func(scaledFont *T.CairoScaledFont,
		extents *T.CairoFontExtents)

	ScaledFontTextExtents func(scaledFont *T.CairoScaledFont,
		utf8 string, extents *T.CairoTextExtents)

	ScaledFontGlyphExtents func(scaledFont *T.CairoScaledFont,
		glyphs *T.CairoGlyph, numGlyphs int,
		extents *T.CairoTextExtents)

	ScaledFontTextToGlyphs func(scaledFont *T.CairoScaledFont,
		x, y float64, utf8 string, utf8Len int,
		glyphs **T.CairoGlyph, numGlyphs *int,
		clusters **T.CairoTextCluster, numClusters *int,
		clusterFlags *T.CairoTextClusterFlags) T.CairoStatus

	ScaledFontGetFontFace func(
		scaledFont *T.CairoScaledFont) *T.CairoFontFace

	ScaledFontGetFontMatrix func(
		scaledFont *T.CairoScaledFont, fontMatrix *T.CairoMatrix)

	ScaledFontGetCtm func(
		scaledFont *T.CairoScaledFont, ctm *T.CairoMatrix)

	ScaledFontGetScaleMatrix func(scaledFont *T.CairoScaledFont,
		scaleMatrix *T.CairoMatrix)

	ScaledFontGetFontOptions func(scaledFont *T.CairoScaledFont,
		options *T.CairoFontOptions)

	ToyFontFaceCreate func(family string,
		slant T.CairoFontSlant,
		weight T.CairoFontWeight) *T.CairoFontFace

	ToyFontFaceGetFamily func(
		fontFace *T.CairoFontFace) string

	ToyFontFaceGetSlant func(
		fontFace *T.CairoFontFace) T.CairoFontSlant

	ToyFontFaceGetWeight func(
		fontFace *T.CairoFontFace) T.CairoFontWeight

	UserFontFaceCreate func() *T.CairoFontFace

	UserFontFaceSetInitFunc func(fontFace *T.CairoFontFace,
		initFunc T.CairoUserScaledFontInitFunc)

	UserFontFaceSetRenderGlyphFunc func(
		fontFace *T.CairoFontFace,
		renderGlyphFunc T.CairoUserScaledFontRenderGlyphFunc)

	UserFontFaceSetTextToGlyphsFunc func(
		fontFace *T.CairoFontFace,
		textToGlyphsFunc T.CairoUserScaledFontTextToGlyphsFunc)

	UserFontFaceSetUnicodeToGlyphFunc func(
		fontFace *T.CairoFontFace,
		unicodeToGlyphFunc T.CairoUserScaledFontUnicodeToGlyphFunc)

	UserFontFaceGetInitFunc func(
		fontFace *T.CairoFontFace) T.CairoUserScaledFontInitFunc

	UserFontFaceGetRenderGlyphFunc func(
		fontFace *T.CairoFontFace) T.CairoUserScaledFontRenderGlyphFunc

	UserFontFaceGetTextToGlyphsFunc func(
		fontFace *T.CairoFontFace) T.CairoUserScaledFontTextToGlyphsFunc

	UserFontFaceGetUnicodeToGlyphFunc func(
		fontFace *T.CairoFontFace) T.CairoUserScaledFontUnicodeToGlyphFunc

	GetOperator func(cr *T.Cairo) T.CairoOperator

	GetSource func(cr *T.Cairo) *T.CairoPattern

	GetTolerance func(cr *T.Cairo) float64

	GetAntialias func(cr *T.Cairo) T.CairoAntialias

	HasCurrentPoint func(cr *T.Cairo) T.CairoBool

	GetCurrentPoint func(cr *T.Cairo, x, y *float64)

	GetFillRule func(cr *T.Cairo) T.CairoFillRule

	GetLineWidth func(cr *T.Cairo) float64

	GetLineCap func(cr *T.Cairo) T.CairoLineCap

	GetLineJoin func(cr *T.Cairo) T.CairoLineJoin

	GetMiterLimit func(cr *T.Cairo) float64

	GetDashCount func(cr *T.Cairo) int

	GetDash func(cr *T.Cairo, dashes, offset *float64)

	GetMatrix func(cr *T.Cairo, matrix *T.CairoMatrix)

	GetTarget func(cr *T.Cairo) *T.CairoSurface

	GetGroupTarget func(cr *T.Cairo) *T.CairoSurface

	CopyPath func(cr *T.Cairo) *T.CairoPath

	CopyPathFlat func(cr *T.Cairo) *T.CairoPath

	AppendPath func(cr *T.Cairo, path *T.CairoPath)

	PathDestroy func(path *T.CairoPath)

	Status func(cr *T.Cairo) T.CairoStatus

	StatusToString func(status T.CairoStatus) string

	DeviceReference func(device *T.CairoDevice) *T.CairoDevice

	DeviceGetType func(device *T.CairoDevice) T.CairoDeviceType

	DeviceStatus func(device *T.CairoDevice) T.CairoStatus

	DeviceAcquire func(device *T.CairoDevice) T.CairoStatus

	DeviceRelease func(device *T.CairoDevice)

	DeviceFlush func(device *T.CairoDevice)

	DeviceFinish func(device *T.CairoDevice)

	DeviceDestroy func(device *T.CairoDevice)

	DeviceGetReferenceCount func(device *T.CairoDevice) uint

	DeviceGetUserData func(device *T.CairoDevice,
		key *T.CairoUserDataKey) *T.Void

	DeviceSetUserData func(device *T.CairoDevice,
		key *T.CairoUserDataKey, userData *T.Void,
		destroy T.CairoDestroyFunc) T.CairoStatus

	SurfaceCreateSimilar func(
		other *T.CairoSurface, content T.CairoContent,
		width, height int) *T.CairoSurface

	SurfaceCreateForRectangle func(target *T.CairoSurface,
		x, y, width, height float64) *T.CairoSurface

	SurfaceReference func(
		surface *T.CairoSurface) *T.CairoSurface

	SurfaceFinish func(surface *T.CairoSurface)

	SurfaceDestroy func(surface *T.CairoSurface)

	SurfaceGetDevice func(surface *T.CairoSurface) *T.CairoDevice

	SurfaceGetReferenceCount func(
		surface *T.CairoSurface) uint

	SurfaceStatus func(surface *T.CairoSurface) T.CairoStatus

	SurfaceGetType func(
		surface *T.CairoSurface) T.CairoSurfaceType

	SurfaceGetContent func(
		surface *T.CairoSurface) T.CairoContent

	SurfaceWriteToPng func(
		surface *T.CairoSurface, filename string) T.CairoStatus

	SurfaceWriteToPngStream func(surface *T.CairoSurface,
		writeFunc T.CairoWriteFunc,
		closure *T.Void) T.CairoStatus

	SurfaceGetUserData func(
		surface *T.CairoSurface, key *T.CairoUserDataKey) *T.Void

	SurfaceSetUserData func(surface *T.CairoSurface,
		key *T.CairoUserDataKey, userData *T.Void,
		destroy T.CairoDestroyFunc) T.CairoStatus

	SurfaceGetMimeData func(surface *T.CairoSurface,
		mimeType string, data **T.UnsignedChar,
		length *T.UnsignedLong)

	SurfaceSetMimeData func(surface *T.CairoSurface,
		mimeType string, data *T.UnsignedChar,
		length T.UnsignedLong, destroy T.CairoDestroyFunc,
		closure *T.Void) T.CairoStatus

	SurfaceGetFontOptions func(
		surface *T.CairoSurface, options *T.CairoFontOptions)

	SurfaceFlush func(surface *T.CairoSurface)

	SurfaceMarkDirty func(surface *T.CairoSurface)

	SurfaceMarkDirtyRectangle func(
		surface *T.CairoSurface, x, y, width, height int)

	SurfaceSetDeviceOffset func(
		surface *T.CairoSurface, xOffset, yOffset float64)

	SurfaceGetDeviceOffset func(
		surface *T.CairoSurface, xOffset, yOffset *float64)

	SurfaceSetFallbackResolution func(surface *T.CairoSurface,
		xPixelsPerInch, yPixelsPerInch float64)

	SurfaceGetFallbackResolution func(surface *T.CairoSurface,
		xPixelsPerInch, yPixelsPerInch *float64)

	SurfaceCopyPage func(surface *T.CairoSurface)

	SurfaceShowPage func(surface *T.CairoSurface)

	SurfaceHasShowTextGlyphs func(
		surface *T.CairoSurface) T.CairoBool

	ImageSurfaceCreate func(format T.CairoFormat,
		width, height int) *T.CairoSurface

	FormatStrideForWidth func(
		format T.CairoFormat, width int) int

	ImageSurfaceCreateForData func(
		data *T.UnsignedChar, format T.CairoFormat,
		width, height, stride int) *T.CairoSurface

	ImageSurfaceGetData func(
		surface *T.CairoSurface) *T.UnsignedChar

	ImageSurfaceGetFormat func(
		surface *T.CairoSurface) T.CairoFormat

	ImageSurfaceGetWidth func(surface *T.CairoSurface) int

	ImageSurfaceGetHeight func(surface *T.CairoSurface) int

	ImageSurfaceGetStride func(surface *T.CairoSurface) int

	ImageSurfaceCreateFromPng func(
		filename string) *T.CairoSurface

	ImageSurfaceCreateFromPngStream func(
		readFunc T.CairoReadFunc,
		closure *T.Void) *T.CairoSurface

	RecordingSurfaceCreate func(
		content T.CairoContent,
		extents *T.CairoRectangle) *T.CairoSurface

	RecordingSurfaceInkExtents func(
		surface *T.CairoSurface, x0, y0, width, height *float64)

	PatternCreateRgb func(
		red, green, blue float64) *T.CairoPattern

	PatternCreateRgba func(
		red, green, blue, alpha float64) *T.CairoPattern

	PatternCreateForSurface func(
		surface *T.CairoSurface) *T.CairoPattern

	PatternCreateLinear func(
		x0, y0, x1, y1 float64) *T.CairoPattern

	PatternCreateRadial func(
		cx0, cy0, radius0,
		cx1, cy1, radius1 float64) *T.CairoPattern

	PatternReference func(
		pattern *T.CairoPattern) *T.CairoPattern

	PatternDestroy func(pattern *T.CairoPattern)

	PatternGetReferenceCount func(pattern *T.CairoPattern) uint

	PatternStatus func(pattern *T.CairoPattern) T.CairoStatus

	PatternGetUserData func(
		pattern *T.CairoPattern, key *T.CairoUserDataKey) *T.Void

	PatternSetUserData func(pattern *T.CairoPattern,
		key *T.CairoUserDataKey, userData *T.Void,
		destroy T.CairoDestroyFunc) T.CairoStatus

	PatternGetType func(
		pattern *T.CairoPattern) T.CairoPatternType

	PatternAddColorStopRgb func(pattern *T.CairoPattern,
		offset, red, green, blue float64)

	PatternAddColorStopRgba func(pattern *T.CairoPattern,
		offset, red, green, blue, alpha float64)

	PatternSetMatrix func(
		pattern *T.CairoPattern, matrix *T.CairoMatrix)

	PatternGetMatrix func(
		pattern *T.CairoPattern, matrix *T.CairoMatrix)

	PatternSetExtend func(
		pattern *T.CairoPattern, extend T.CairoExtend)

	PatternGetExtend func(
		pattern *T.CairoPattern) T.CairoExtend

	PatternSetFilter func(
		pattern *T.CairoPattern, filter T.CairoFilter)

	PatternGetFilter func(
		pattern *T.CairoPattern) T.CairoFilter

	PatternGetRgba func(pattern *T.CairoPattern,
		red, green, blue, alpha *float64) T.CairoStatus

	PatternGetSurface func(pattern *T.CairoPattern,
		surface **T.CairoSurface) T.CairoStatus

	PatternGetColorStopRgba func(
		pattern *T.CairoPattern, index int,
		offset, red, green, blue, alpha *float64) T.CairoStatus

	PatternGetColorStopCount func(
		pattern *T.CairoPattern, count *int) T.CairoStatus

	PatternGetLinearPoints func(pattern *T.CairoPattern,
		x0, y0, x1, y1 *float64) T.CairoStatus

	PatternGetRadialCircles func(pattern *T.CairoPattern,
		x0, y0, r0, x1, y1, r1 *float64) T.CairoStatus

	MatrixInit func(
		matrix *T.CairoMatrix, xx, yx, xy, yy, x0, y0 float64)

	MatrixInitIdentity func(matrix *T.CairoMatrix)

	MatrixInitTranslate func(
		matrix *T.CairoMatrix, tx, ty float64)

	MatrixInitScale func(
		matrix *T.CairoMatrix, sx, sy float64)

	MatrixInitRotate func(matrix *T.CairoMatrix, radians float64)

	MatrixTranslate func(matrix *T.CairoMatrix, tx, ty float64)

	MatrixScale func(matrix *T.CairoMatrix, sx, sy float64)

	MatrixRotate func(matrix *T.CairoMatrix, radians float64)

	MatrixInvert func(matrix *T.CairoMatrix) T.CairoStatus

	MatrixMultiply func(result, a, b *T.CairoMatrix)

	MatrixTransformDistance func(
		matrix *T.CairoMatrix, dx, dy *float64)

	MatrixTransformPoint func(
		matrix *T.CairoMatrix, x, y *float64)

	RegionCreate func() *T.CairoRegion

	RegionCreateRectangle func(
		rectangle *T.CairoRectangleInt) *T.CairoRegion

	RegionCreateRectangles func(
		rects *T.CairoRectangleInt,
		count int) *T.CairoRegion

	RegionCopy func(original *T.CairoRegion) *T.CairoRegion

	RegionReference func(region *T.CairoRegion) *T.CairoRegion

	RegionDestroy func(region *T.CairoRegion)

	RegionEqual func(a, b *T.CairoRegion) T.CairoBool

	RegionStatus func(region *T.CairoRegion) T.CairoStatus

	RegionGetExtents func(region *T.CairoRegion,
		extents *T.CairoRectangleInt)

	RegionNumRectangles func(region *T.CairoRegion) int

	RegionGetRectangle func(region *T.CairoRegion,
		nth int, rectangle *T.CairoRectangleInt)

	RegionIsEmpty func(region *T.CairoRegion) T.CairoBool

	RegionContainsRectangle func(region *T.CairoRegion,
		rectangle *T.CairoRectangleInt) T.CairoRegionOverlap

	RegionContainsPoint func(
		region *T.CairoRegion, x, y int) T.CairoBool

	RegionTranslate func(region *T.CairoRegion, dx, dy int)

	RegionSubtract func(dst, other *T.CairoRegion) T.CairoStatus

	RegionSubtractRectangle func(dst *T.CairoRegion,
		rectangle *T.CairoRectangleInt) T.CairoStatus

	RegionIntersect func(dst, other *T.CairoRegion) T.CairoStatus

	RegionIntersectRectangle func(dst *T.CairoRegion,
		rectangle *T.CairoRectangleInt) T.CairoStatus

	RegionUnion func(dst, other *T.CairoRegion) T.CairoStatus

	RegionUnionRectangle func(dst *T.CairoRegion,
		rectangle *T.CairoRectangleInt) T.CairoStatus

	RegionXor func(dst, other *T.CairoRegion) T.CairoStatus

	RegionXorRectangle func(dst *T.CairoRegion,
		rectangle *T.CairoRectangleInt) T.CairoStatus

	DebugResetStaticData func()

	GobjectContextGetType func()

	GobjectDeviceGetType func()

	GobjectPatternGetType func()

	GobjectSurfaceGetType func()

	GobjectRectangleGetType func()

	GobjectScaledFontGetType func()

	GobjectFontFaceGetType func()

	GobjectFontOptionsGetType func()

	GobjectRectangleIntGetType func()

	GobjectRegionGetType func()

	GobjectStatusGetType func()

	GobjectContentGetType func()

	GobjectOperatorGetType func()

	GobjectAntialiasGetType func()

	GobjectFillRuleGetType func()

	GobjectLineCapGetType func()

	GobjectLineJoinGetType func()

	GobjectTextClusterFlagsGetType func()

	GobjectFontSlantGetType func()

	GobjectFontWeightGetType func()

	GobjectSubpixelOrderGetType func()

	GobjectHintStyleGetType func()

	GobjectHintMetricsGetType func()

	GobjectFontTypeGetType func()

	GobjectPathDataTypeGetType func()

	GobjectDeviceTypeGetType func()

	GobjectSurfaceTypeGetType func()

	GobjectFormatGetType func()

	GobjectPatternTypeGetType func()

	GobjectExtendGetType func()

	GobjectFilterGetType func()

	GobjectRegionOverlapGetType func()

	ScriptInterpreterCreate func() *T.CairoScriptInterpreter

	ScriptInterpreterInstallHooks func(
		ctx *T.CairoScriptInterpreter,
		hooks *T.CairoScriptInterpreterHooks)

	ScriptInterpreterRun func(
		ctx *T.CairoScriptInterpreter,
		filename string) T.CairoStatus

	ScriptInterpreterFeedStream func(
		ctx *T.CairoScriptInterpreter,
		stream *T.FILE) T.CairoStatus

	ScriptInterpreterFeedString func(
		ctx *T.CairoScriptInterpreter,
		line string,
		leng int) T.CairoStatus

	ScriptInterpreterGetLineNumber func(
		ctx *T.CairoScriptInterpreter) uint

	ScriptInterpreterReference func(
		ctx *T.CairoScriptInterpreter) *T.CairoScriptInterpreter

	ScriptInterpreterFinish func(
		ctx *T.CairoScriptInterpreter) T.CairoStatus

	ScriptInterpreterDestroy func(
		ctx *T.CairoScriptInterpreter) T.CairoStatus

	ScriptInterpreterTranslateStream func(
		stream *T.FILE,
		writeFunc T.CairoWriteFunc,
		closure *T.Void) T.CairoStatus

	FtFontFaceCreateForFtFace func(
		face T.FTFace,
		loadFlags int) *T.CairoFontFace

	FtScaledFontLockFace func(
		scaledFont *T.CairoScaledFont) T.FTFace

	FtScaledFontUnlockFace func(
		scaledFont *T.CairoScaledFont)

	FtFontFaceCreateForPattern func(
		pattern *T.FcPattern) *T.CairoFontFace

	FtFontOptionsSubstitute func(
		options *T.CairoFontOptions,
		pattern *T.FcPattern)

	PdfSurfaceCreate func(
		filename string,
		widthInPoints float64,
		heightInPoints float64) *T.CairoSurface

	PdfSurfaceCreateForStream func(
		writeFunc T.CairoWriteFunc,
		closure *T.Void,
		widthInPoints float64,
		heightInPoints float64) *T.CairoSurface

	PdfSurfaceRestrictToVersion func(
		surface *T.CairoSurface,
		version PdfVersion)

	PdfGetVersions func(
		versions **PdfVersion,
		numVersions *int)

	PdfVersionToString func(
		version PdfVersion) string

	PdfSurfaceSetSize func(
		surface *T.CairoSurface,
		widthInPoints float64,
		heightInPoints float64)

	PsSurfaceCreate func(
		filename string,
		widthInPoints float64,
		heightInPoints float64) *T.CairoSurface

	PsSurfaceCreateForStream func(
		writeFunc T.CairoWriteFunc,
		closure *T.Void,
		widthInPoints float64,
		heightInPoints float64) *T.CairoSurface

	PsSurfaceRestrictToLevel func(
		surface *T.CairoSurface,
		level PsLevel)

	PsGetLevels func(
		levels **PsLevel,
		numLevels *int)

	PsLevelToString func(
		level PsLevel)

	PsSurfaceSetEps func(
		surface *T.CairoSurface,
		eps T.CairoBool) string

	PsSurfaceGetEps func(
		surface *T.CairoSurface)

	PsSurfaceSetSize func(
		surface *T.CairoSurface,
		widthInPoints float64,
		heightInPoints float64) T.CairoBool

	PsSurfaceDscComment func(
		surface *T.CairoSurface,
		comment string)

	PsSurfaceDscBeginSetup func(
		surface *T.CairoSurface)

	PsSurfaceDscBeginPageSetup func(
		surface *T.CairoSurface)

	SvgSurfaceCreate func(
		filename string,
		widthInPoints float64,
		heightInPoints float64) *T.CairoSurface

	SvgSurfaceCreateForStream func(
		writeFunc T.CairoWriteFunc,
		closure *T.Void,
		widthInPoints float64,
		heightInPoints float64) *T.CairoSurface

	SvgSurfaceRestrictToVersion func(
		surface *T.CairoSurface,
		version SvgVersion)

	SvgGetVersions func(
		versions **SvgVersion,
		numVersions *int)

	SvgVersionToString func(
		version SvgVersion) string

	Win32SurfaceCreate func(hdc HDC) *T.CairoSurface

	Win32PrintingSurfaceCreate func(hdc HDC) *T.CairoSurface

	Win32SurfaceCreateWithDdb func(hdc HDC, format T.CairoFormat,
		width, height int) *T.CairoSurface

	Win32SurfaceCreateWithDib func(
		format T.CairoFormat,
		width int,
		height int) *T.CairoSurface

	Win32SurfaceGetDc func(
		surface *T.CairoSurface) HDC

	Win32SurfaceGetImage func(
		surface *T.CairoSurface) *T.CairoSurface

	Win32FontFaceCreateForLogfontw func(
		logfont *LOGFONTW) *T.CairoFontFace

	Win32FontFaceCreateForHfont func(
		font HFONT) *T.CairoFontFace

	Win32FontFaceCreateForLogfontwHfont func(
		logfont *LOGFONTW,
		font HFONT) *T.CairoFontFace

	Win32ScaledFontSelectFont func(
		scaledFont *T.CairoScaledFont,
		hdc HDC) T.CairoStatus

	Win32ScaledFontDoneFont func(
		scaledFont *T.CairoScaledFont)

	Win32ScaledFontGetMetricsFactor func(
		scaledFont *T.CairoScaledFont) float64

	Win32ScaledFontGetLogicalToDevice func(
		scaledFont *T.CairoScaledFont,
		logicalToDevice *T.CairoMatrix)

	Win32ScaledFontGetDeviceToLogical func(
		scaledFont *T.CairoScaledFont,
		deviceToLogical *T.CairoMatrix)
)

var dll = "libcairo-2.dll"

var apiList = outside.Apis{
	{"cairo_append_path", &AppendPath},
	{"cairo_arc", &Arc},
	{"cairo_arc_negative", &ArcNegative},
	{"cairo_clip", &Clip},
	{"cairo_clip_extents", &ClipExtents},
	{"cairo_clip_preserve", &ClipPreserve},
	{"cairo_close_path", &ClosePath},
	{"cairo_copy_clip_rectangle_list", &CopyClipRectangleList},
	{"cairo_copy_page", &CopyPage},
	{"cairo_copy_path", &CopyPath},
	{"cairo_copy_path_flat", &CopyPathFlat},
	{"cairo_create", &Create},
	{"cairo_curve_to", &CurveTo},
	{"cairo_debug_reset_static_data", &DebugResetStaticData},
	{"cairo_destroy", &Destroy},
	{"cairo_device_acquire", &DeviceAcquire},
	{"cairo_device_destroy", &DeviceDestroy},
	{"cairo_device_finish", &DeviceFinish},
	{"cairo_device_flush", &DeviceFlush},
	{"cairo_device_get_reference_count", &DeviceGetReferenceCount},
	{"cairo_device_get_type", &DeviceGetType},
	{"cairo_device_get_user_data", &DeviceGetUserData},
	{"cairo_device_reference", &DeviceReference},
	{"cairo_device_release", &DeviceRelease},
	{"cairo_device_set_user_data", &DeviceSetUserData},
	{"cairo_device_status", &DeviceStatus},
	{"cairo_device_to_user", &DeviceToUser},
	{"cairo_device_to_user_distance", &DeviceToUserDistance},
	{"cairo_fill", &Fill},
	{"cairo_fill_extents", &FillExtents},
	{"cairo_fill_preserve", &FillPreserve},
	{"cairo_font_extents", &FontExtents},
	{"cairo_font_face_destroy", &FontFaceDestroy},
	{"cairo_font_face_get_reference_count", &FontFaceGetReferenceCount},
	{"cairo_font_face_get_type", &FontFaceGetType},
	{"cairo_font_face_get_user_data", &FontFaceGetUserData},
	{"cairo_font_face_reference", &FontFaceReference},
	{"cairo_font_face_set_user_data", &FontFaceSetUserData},
	{"cairo_font_face_status", &FontFaceStatus},
	{"cairo_font_options_copy", &FontOptionsCopy},
	{"cairo_font_options_create", &FontOptionsCreate},
	{"cairo_font_options_destroy", &FontOptionsDestroy},
	{"cairo_font_options_equal", &FontOptionsEqual},
	{"cairo_font_options_get_antialias", &FontOptionsGetAntialias},
	{"cairo_font_options_get_hint_metrics", &FontOptionsGetHintMetrics},
	{"cairo_font_options_get_hint_style", &FontOptionsGetHintStyle},
	{"cairo_font_options_get_subpixel_order", &FontOptionsGetSubpixelOrder},
	{"cairo_font_options_hash", &FontOptionsHash},
	{"cairo_font_options_merge", &FontOptionsMerge},
	{"cairo_font_options_set_antialias", &FontOptionsSetAntialias},
	{"cairo_font_options_set_hint_metrics", &FontOptionsSetHintMetrics},
	{"cairo_font_options_set_hint_style", &FontOptionsSetHintStyle},
	{"cairo_font_options_set_subpixel_order", &FontOptionsSetSubpixelOrder},
	{"cairo_font_options_status", &FontOptionsStatus},
	{"cairo_format_stride_for_width", &FormatStrideForWidth},
	{"cairo_ft_font_face_create_for_ft_face", &FtFontFaceCreateForFtFace},
	{"cairo_ft_font_face_create_for_pattern", &FtFontFaceCreateForPattern},
	{"cairo_ft_font_options_substitute", &FtFontOptionsSubstitute},
	{"cairo_ft_scaled_font_lock_face", &FtScaledFontLockFace},
	{"cairo_ft_scaled_font_unlock_face", &FtScaledFontUnlockFace},
	{"cairo_get_antialias", &GetAntialias},
	{"cairo_get_current_point", &GetCurrentPoint},
	{"cairo_get_dash", &GetDash},
	{"cairo_get_dash_count", &GetDashCount},
	{"cairo_get_fill_rule", &GetFillRule},
	{"cairo_get_font_face", &GetFontFace},
	{"cairo_get_font_matrix", &GetFontMatrix},
	{"cairo_get_font_options", &GetFontOptions},
	{"cairo_get_group_target", &GetGroupTarget},
	{"cairo_get_line_cap", &GetLineCap},
	{"cairo_get_line_join", &GetLineJoin},
	{"cairo_get_line_width", &GetLineWidth},
	{"cairo_get_matrix", &GetMatrix},
	{"cairo_get_miter_limit", &GetMiterLimit},
	{"cairo_get_operator", &GetOperator},
	{"cairo_get_reference_count", &GetReferenceCount},
	{"cairo_get_scaled_font", &GetScaledFont},
	{"cairo_get_source", &GetSource},
	{"cairo_get_target", &GetTarget},
	{"cairo_get_tolerance", &GetTolerance},
	{"cairo_get_user_data", &GetUserData},
	{"cairo_glyph_allocate", &GlyphAllocate},
	{"cairo_glyph_extents", &GlyphExtents},
	{"cairo_glyph_free", &GlyphFree},
	{"cairo_glyph_path", &GlyphPath},
	{"cairo_has_current_point", &HasCurrentPoint},
	{"cairo_identity_matrix", &IdentityMatrix},
	{"cairo_image_surface_create", &ImageSurfaceCreate},
	{"cairo_image_surface_create_for_data", &ImageSurfaceCreateForData},
	{"cairo_image_surface_create_from_png", &ImageSurfaceCreateFromPng},
	{"cairo_image_surface_create_from_png_stream", &ImageSurfaceCreateFromPngStream},
	{"cairo_image_surface_get_data", &ImageSurfaceGetData},
	{"cairo_image_surface_get_format", &ImageSurfaceGetFormat},
	{"cairo_image_surface_get_height", &ImageSurfaceGetHeight},
	{"cairo_image_surface_get_stride", &ImageSurfaceGetStride},
	{"cairo_image_surface_get_width", &ImageSurfaceGetWidth},
	{"cairo_in_clip", &InClip},
	{"cairo_in_fill", &InFill},
	{"cairo_in_stroke", &InStroke},
	{"cairo_line_to", &LineTo},
	{"cairo_mask", &Mask},
	{"cairo_mask_surface", &MaskSurface},
	{"cairo_matrix_init", &MatrixInit},
	{"cairo_matrix_init_identity", &MatrixInitIdentity},
	{"cairo_matrix_init_rotate", &MatrixInitRotate},
	{"cairo_matrix_init_scale", &MatrixInitScale},
	{"cairo_matrix_init_translate", &MatrixInitTranslate},
	{"cairo_matrix_invert", &MatrixInvert},
	{"cairo_matrix_multiply", &MatrixMultiply},
	{"cairo_matrix_rotate", &MatrixRotate},
	{"cairo_matrix_scale", &MatrixScale},
	{"cairo_matrix_transform_distance", &MatrixTransformDistance},
	{"cairo_matrix_transform_point", &MatrixTransformPoint},
	{"cairo_matrix_translate", &MatrixTranslate},
	{"cairo_move_to", &MoveTo},
	{"cairo_new_path", &NewPath},
	{"cairo_new_sub_path", &NewSubPath},
	{"cairo_paint", &Paint},
	{"cairo_paint_with_alpha", &PaintWithAlpha},
	{"cairo_path_destroy", &PathDestroy},
	{"cairo_path_extents", &PathExtents},
	{"cairo_pattern_add_color_stop_rgb", &PatternAddColorStopRgb},
	{"cairo_pattern_add_color_stop_rgba", &PatternAddColorStopRgba},
	{"cairo_pattern_create_for_surface", &PatternCreateForSurface},
	{"cairo_pattern_create_linear", &PatternCreateLinear},
	{"cairo_pattern_create_radial", &PatternCreateRadial},
	{"cairo_pattern_create_rgb", &PatternCreateRgb},
	{"cairo_pattern_create_rgba", &PatternCreateRgba},
	{"cairo_pattern_destroy", &PatternDestroy},
	{"cairo_pattern_get_color_stop_count", &PatternGetColorStopCount},
	{"cairo_pattern_get_color_stop_rgba", &PatternGetColorStopRgba},
	{"cairo_pattern_get_extend", &PatternGetExtend},
	{"cairo_pattern_get_filter", &PatternGetFilter},
	{"cairo_pattern_get_linear_points", &PatternGetLinearPoints},
	{"cairo_pattern_get_matrix", &PatternGetMatrix},
	{"cairo_pattern_get_radial_circles", &PatternGetRadialCircles},
	{"cairo_pattern_get_reference_count", &PatternGetReferenceCount},
	{"cairo_pattern_get_rgba", &PatternGetRgba},
	{"cairo_pattern_get_surface", &PatternGetSurface},
	{"cairo_pattern_get_type", &PatternGetType},
	{"cairo_pattern_get_user_data", &PatternGetUserData},
	{"cairo_pattern_reference", &PatternReference},
	{"cairo_pattern_set_extend", &PatternSetExtend},
	{"cairo_pattern_set_filter", &PatternSetFilter},
	{"cairo_pattern_set_matrix", &PatternSetMatrix},
	{"cairo_pattern_set_user_data", &PatternSetUserData},
	{"cairo_pattern_status", &PatternStatus},
	{"cairo_pdf_get_versions", &PdfGetVersions},
	{"cairo_pdf_surface_create", &PdfSurfaceCreate},
	{"cairo_pdf_surface_create_for_stream", &PdfSurfaceCreateForStream},
	{"cairo_pdf_surface_restrict_to_version", &PdfSurfaceRestrictToVersion},
	{"cairo_pdf_surface_set_size", &PdfSurfaceSetSize},
	{"cairo_pdf_version_to_string", &PdfVersionToString},
	{"cairo_pop_group", &PopGroup},
	{"cairo_pop_group_to_source", &PopGroupToSource},
	{"cairo_ps_get_levels", &PsGetLevels},
	{"cairo_ps_level_to_string", &PsLevelToString},
	{"cairo_ps_surface_create", &PsSurfaceCreate},
	{"cairo_ps_surface_create_for_stream", &PsSurfaceCreateForStream},
	{"cairo_ps_surface_dsc_begin_page_setup", &PsSurfaceDscBeginPageSetup},
	{"cairo_ps_surface_dsc_begin_setup", &PsSurfaceDscBeginSetup},
	{"cairo_ps_surface_dsc_comment", &PsSurfaceDscComment},
	{"cairo_ps_surface_get_eps", &PsSurfaceGetEps},
	{"cairo_ps_surface_restrict_to_level", &PsSurfaceRestrictToLevel},
	{"cairo_ps_surface_set_eps", &PsSurfaceSetEps},
	{"cairo_ps_surface_set_size", &PsSurfaceSetSize},
	{"cairo_push_group", &PushGroup},
	{"cairo_push_group_with_content", &PushGroupWithContent},
	{"cairo_recording_surface_create", &RecordingSurfaceCreate},
	{"cairo_recording_surface_ink_extents", &RecordingSurfaceInkExtents},
	{"cairo_rectangle", &Rectangle},
	{"cairo_rectangle_list_destroy", &RectangleListDestroy},
	{"cairo_reference", &Reference},
	{"cairo_region_contains_point", &RegionContainsPoint},
	{"cairo_region_contains_rectangle", &RegionContainsRectangle},
	{"cairo_region_copy", &RegionCopy},
	{"cairo_region_create", &RegionCreate},
	{"cairo_region_create_rectangle", &RegionCreateRectangle},
	{"cairo_region_create_rectangles", &RegionCreateRectangles},
	{"cairo_region_destroy", &RegionDestroy},
	{"cairo_region_equal", &RegionEqual},
	{"cairo_region_get_extents", &RegionGetExtents},
	{"cairo_region_get_rectangle", &RegionGetRectangle},
	{"cairo_region_intersect", &RegionIntersect},
	{"cairo_region_intersect_rectangle", &RegionIntersectRectangle},
	{"cairo_region_is_empty", &RegionIsEmpty},
	{"cairo_region_num_rectangles", &RegionNumRectangles},
	{"cairo_region_reference", &RegionReference},
	{"cairo_region_status", &RegionStatus},
	{"cairo_region_subtract", &RegionSubtract},
	{"cairo_region_subtract_rectangle", &RegionSubtractRectangle},
	{"cairo_region_translate", &RegionTranslate},
	{"cairo_region_union", &RegionUnion},
	{"cairo_region_union_rectangle", &RegionUnionRectangle},
	{"cairo_region_xor", &RegionXor},
	{"cairo_region_xor_rectangle", &RegionXorRectangle},
	{"cairo_rel_curve_to", &RelCurveTo},
	{"cairo_rel_line_to", &RelLineTo},
	{"cairo_rel_move_to", &RelMoveTo},
	{"cairo_reset_clip", &ResetClip},
	{"cairo_restore", &Restore},
	{"cairo_rotate", &Rotate},
	{"cairo_save", &Save},
	{"cairo_scale", &Scale},
	{"cairo_scaled_font_create", &ScaledFontCreate},
	{"cairo_scaled_font_destroy", &ScaledFontDestroy},
	{"cairo_scaled_font_extents", &ScaledFontExtents},
	{"cairo_scaled_font_get_ctm", &ScaledFontGetCtm},
	{"cairo_scaled_font_get_font_face", &ScaledFontGetFontFace},
	{"cairo_scaled_font_get_font_matrix", &ScaledFontGetFontMatrix},
	{"cairo_scaled_font_get_font_options", &ScaledFontGetFontOptions},
	{"cairo_scaled_font_get_reference_count", &ScaledFontGetReferenceCount},
	{"cairo_scaled_font_get_scale_matrix", &ScaledFontGetScaleMatrix},
	{"cairo_scaled_font_get_type", &ScaledFontGetType},
	{"cairo_scaled_font_get_user_data", &ScaledFontGetUserData},
	{"cairo_scaled_font_glyph_extents", &ScaledFontGlyphExtents},
	{"cairo_scaled_font_reference", &ScaledFontReference},
	{"cairo_scaled_font_set_user_data", &ScaledFontSetUserData},
	{"cairo_scaled_font_status", &ScaledFontStatus},
	{"cairo_scaled_font_text_extents", &ScaledFontTextExtents},
	{"cairo_scaled_font_text_to_glyphs", &ScaledFontTextToGlyphs},
	{"cairo_select_font_face", &SelectFontFace},
	{"cairo_set_antialias", &SetAntialias},
	{"cairo_set_dash", &SetDash},
	{"cairo_set_fill_rule", &SetFillRule},
	{"cairo_set_font_face", &SetFontFace},
	{"cairo_set_font_matrix", &SetFontMatrix},
	{"cairo_set_font_options", &SetFontOptions},
	{"cairo_set_font_size", &SetFontSize},
	{"cairo_set_line_cap", &SetLineCap},
	{"cairo_set_line_join", &SetLineJoin},
	{"cairo_set_line_width", &SetLineWidth},
	{"cairo_set_matrix", &SetMatrix},
	{"cairo_set_miter_limit", &SetMiterLimit},
	{"cairo_set_operator", &SetOperator},
	{"cairo_set_scaled_font", &SetScaledFont},
	{"cairo_set_source", &SetSource},
	{"cairo_set_source_rgb", &SetSourceRgb},
	{"cairo_set_source_rgba", &SetSourceRgba},
	{"cairo_set_source_surface", &SetSourceSurface},
	{"cairo_set_tolerance", &SetTolerance},
	{"cairo_set_user_data", &SetUserData},
	{"cairo_show_glyphs", &ShowGlyphs},
	{"cairo_show_page", &ShowPage},
	{"cairo_show_text", &ShowText},
	{"cairo_show_text_glyphs", &ShowTextGlyphs},
	{"cairo_status", &Status},
	{"cairo_status_to_string", &StatusToString},
	{"cairo_stroke", &Stroke},
	{"cairo_stroke_extents", &StrokeExtents},
	{"cairo_stroke_preserve", &StrokePreserve},
	{"cairo_surface_copy_page", &SurfaceCopyPage},
	{"cairo_surface_create_for_rectangle", &SurfaceCreateForRectangle},
	{"cairo_surface_create_similar", &SurfaceCreateSimilar},
	{"cairo_surface_destroy", &SurfaceDestroy},
	{"cairo_surface_finish", &SurfaceFinish},
	{"cairo_surface_flush", &SurfaceFlush},
	{"cairo_surface_get_content", &SurfaceGetContent},
	{"cairo_surface_get_device", &SurfaceGetDevice},
	{"cairo_surface_get_device_offset", &SurfaceGetDeviceOffset},
	{"cairo_surface_get_fallback_resolution", &SurfaceGetFallbackResolution},
	{"cairo_surface_get_font_options", &SurfaceGetFontOptions},
	{"cairo_surface_get_mime_data", &SurfaceGetMimeData},
	{"cairo_surface_get_reference_count", &SurfaceGetReferenceCount},
	{"cairo_surface_get_type", &SurfaceGetType},
	{"cairo_surface_get_user_data", &SurfaceGetUserData},
	{"cairo_surface_has_show_text_glyphs", &SurfaceHasShowTextGlyphs},
	{"cairo_surface_mark_dirty", &SurfaceMarkDirty},
	{"cairo_surface_mark_dirty_rectangle", &SurfaceMarkDirtyRectangle},
	{"cairo_surface_reference", &SurfaceReference},
	{"cairo_surface_set_device_offset", &SurfaceSetDeviceOffset},
	{"cairo_surface_set_fallback_resolution", &SurfaceSetFallbackResolution},
	{"cairo_surface_set_mime_data", &SurfaceSetMimeData},
	{"cairo_surface_set_user_data", &SurfaceSetUserData},
	{"cairo_surface_show_page", &SurfaceShowPage},
	{"cairo_surface_status", &SurfaceStatus},
	{"cairo_surface_write_to_png", &SurfaceWriteToPng},
	{"cairo_surface_write_to_png_stream", &SurfaceWriteToPngStream},
	{"cairo_svg_get_versions", &SvgGetVersions},
	{"cairo_svg_surface_create", &SvgSurfaceCreate},
	{"cairo_svg_surface_create_for_stream", &SvgSurfaceCreateForStream},
	{"cairo_svg_surface_restrict_to_version", &SvgSurfaceRestrictToVersion},
	{"cairo_svg_version_to_string", &SvgVersionToString},
	{"cairo_text_cluster_allocate", &TextClusterAllocate},
	{"cairo_text_cluster_free", &TextClusterFree},
	{"cairo_text_extents", &TextExtents},
	{"cairo_text_path", &TextPath},
	{"cairo_toy_font_face_create", &ToyFontFaceCreate},
	{"cairo_toy_font_face_get_family", &ToyFontFaceGetFamily},
	{"cairo_toy_font_face_get_slant", &ToyFontFaceGetSlant},
	{"cairo_toy_font_face_get_weight", &ToyFontFaceGetWeight},
	{"cairo_transform", &Transform},
	{"cairo_translate", &Translate},
	{"cairo_user_font_face_create", &UserFontFaceCreate},
	{"cairo_user_font_face_get_init_func", &UserFontFaceGetInitFunc},
	{"cairo_user_font_face_get_render_glyph_func", &UserFontFaceGetRenderGlyphFunc},
	{"cairo_user_font_face_get_text_to_glyphs_func", &UserFontFaceGetTextToGlyphsFunc},
	{"cairo_user_font_face_get_unicode_to_glyph_func", &UserFontFaceGetUnicodeToGlyphFunc},
	{"cairo_user_font_face_set_init_func", &UserFontFaceSetInitFunc},
	{"cairo_user_font_face_set_render_glyph_func", &UserFontFaceSetRenderGlyphFunc},
	{"cairo_user_font_face_set_text_to_glyphs_func", &UserFontFaceSetTextToGlyphsFunc},
	{"cairo_user_font_face_set_unicode_to_glyph_func", &UserFontFaceSetUnicodeToGlyphFunc},
	{"cairo_user_to_device", &UserToDevice},
	{"cairo_user_to_device_distance", &UserToDeviceDistance},
	{"cairo_version", &Version},
	{"cairo_version_string", &VersionString},
	{"cairo_win32_font_face_create_for_hfont", &Win32FontFaceCreateForHfont},
	{"cairo_win32_font_face_create_for_logfontw", &Win32FontFaceCreateForLogfontw},
	{"cairo_win32_font_face_create_for_logfontw_hfont", &Win32FontFaceCreateForLogfontwHfont},
	{"cairo_win32_printing_surface_create", &Win32PrintingSurfaceCreate},
	{"cairo_win32_scaled_font_done_font", &Win32ScaledFontDoneFont},
	{"cairo_win32_scaled_font_get_device_to_logical", &Win32ScaledFontGetDeviceToLogical},
	{"cairo_win32_scaled_font_get_logical_to_device", &Win32ScaledFontGetLogicalToDevice},
	{"cairo_win32_scaled_font_get_metrics_factor", &Win32ScaledFontGetMetricsFactor},
	{"cairo_win32_scaled_font_select_font", &Win32ScaledFontSelectFont},
	{"cairo_win32_surface_create", &Win32SurfaceCreate},
	{"cairo_win32_surface_create_with_ddb", &Win32SurfaceCreateWithDdb},
	{"cairo_win32_surface_create_with_dib", &Win32SurfaceCreateWithDib},
	{"cairo_win32_surface_get_dc", &Win32SurfaceGetDc},
	{"cairo_win32_surface_get_image", &Win32SurfaceGetImage},
}

var dllGobject = "libcairo-gobject-2.dll"

var apiListGobject = outside.Apis{
	{"cairo_gobject_antialias_get_type", &GobjectAntialiasGetType},
	{"cairo_gobject_content_get_type", &GobjectContentGetType},
	{"cairo_gobject_context_get_type", &GobjectContextGetType},
	{"cairo_gobject_device_get_type", &GobjectDeviceGetType},
	{"cairo_gobject_device_type_get_type", &GobjectDeviceTypeGetType},
	{"cairo_gobject_extend_get_type", &GobjectExtendGetType},
	{"cairo_gobject_fill_rule_get_type", &GobjectFillRuleGetType},
	{"cairo_gobject_filter_get_type", &GobjectFilterGetType},
	{"cairo_gobject_font_face_get_type", &GobjectFontFaceGetType},
	{"cairo_gobject_font_options_get_type", &GobjectFontOptionsGetType},
	{"cairo_gobject_font_slant_get_type", &GobjectFontSlantGetType},
	{"cairo_gobject_font_type_get_type", &GobjectFontTypeGetType},
	{"cairo_gobject_font_weight_get_type", &GobjectFontWeightGetType},
	{"cairo_gobject_format_get_type", &GobjectFormatGetType},
	{"cairo_gobject_hint_metrics_get_type", &GobjectHintMetricsGetType},
	{"cairo_gobject_hint_style_get_type", &GobjectHintStyleGetType},
	{"cairo_gobject_line_cap_get_type", &GobjectLineCapGetType},
	{"cairo_gobject_line_join_get_type", &GobjectLineJoinGetType},
	{"cairo_gobject_operator_get_type", &GobjectOperatorGetType},
	{"cairo_gobject_path_data_type_get_type", &GobjectPathDataTypeGetType},
	{"cairo_gobject_pattern_get_type", &GobjectPatternGetType},
	{"cairo_gobject_pattern_type_get_type", &GobjectPatternTypeGetType},
	{"cairo_gobject_rectangle_get_type", &GobjectRectangleGetType},
	{"cairo_gobject_rectangle_int_get_type", &GobjectRectangleIntGetType},
	{"cairo_gobject_region_get_type", &GobjectRegionGetType},
	{"cairo_gobject_region_overlap_get_type", &GobjectRegionOverlapGetType},
	{"cairo_gobject_scaled_font_get_type", &GobjectScaledFontGetType},
	{"cairo_gobject_status_get_type", &GobjectStatusGetType},
	{"cairo_gobject_subpixel_order_get_type", &GobjectSubpixelOrderGetType},
	{"cairo_gobject_surface_get_type", &GobjectSurfaceGetType},
	{"cairo_gobject_surface_type_get_type", &GobjectSurfaceTypeGetType},
	{"cairo_gobject_text_cluster_flags_get_type", &GobjectTextClusterFlagsGetType},
}

var dllScript = "libcairo-script-interpreter-2.dll"

var apiListScript = outside.Apis{
	// Undocumented {"_csi_alloc", &CsiAlloc},
	// Undocumented {"_csi_alloc0", &CsiAlloc0},
	// Undocumented {"_csi_array_execute", &CsiArrayExecute},
	// Undocumented {"_csi_error", &CsiError},
	// Undocumented {"_csi_file_as_string", &CsiFileAsString},
	// Undocumented {"_csi_file_execute", &CsiFileExecute},
	// Undocumented {"_csi_file_free", &CsiFileFree},
	// Undocumented {"_csi_free", &CsiFree},
	// Undocumented {"_csi_hash_table_fini", &CsiHashTableFini},
	// Undocumented {"_csi_hash_table_foreach", &CsiHashTableForeach},
	// Undocumented {"_csi_hash_table_init", &CsiHashTableInit},
	// Undocumented {"_csi_hash_table_insert", &CsiHashTableInsert},
	// Undocumented {"_csi_hash_table_lookup", &CsiHashTableLookup},
	// Undocumented {"_csi_hash_table_remove", &CsiHashTableRemove},
	// Undocumented {"_csi_integer_constants", &CsiIntegerConstants},
	// Undocumented {"_csi_intern_string", &CsiInternString},
	// Undocumented {"_csi_name_define", &CsiNameDefine},
	// Undocumented {"_csi_name_lookup", &CsiNameLookup},
	// Undocumented {"_csi_name_undefine", &CsiNameUndefine},
	// Undocumented {"_csi_operators", &CsiOperators},
	// Undocumented {"_csi_parse_number", &CsiParseNumber},
	// Undocumented {"_csi_perm_alloc", &CsiPermAlloc},
	// Undocumented {"_csi_real_constants", &CsiRealConstants},
	// Undocumented {"_csi_realloc", &CsiRealloc},
	// Undocumented {"_csi_scan_file", &CsiScanFile},
	// Undocumented {"_csi_scanner_fini", &CsiScannerFini},
	// Undocumented {"_csi_scanner_init", &CsiScannerInit},
	// Undocumented {"_csi_slab_alloc", &CsiSlabAlloc},
	// Undocumented {"_csi_slab_free", &CsiSlabFree},
	// Undocumented {"_csi_stack_exch", &CsiStackExch},
	// Undocumented {"_csi_stack_fini", &CsiStackFini},
	// Undocumented {"_csi_stack_grow", &CsiStackGrow},
	// Undocumented {"_csi_stack_init", &CsiStackInit},
	// Undocumented {"_csi_stack_peek", &CsiStackPeek},
	// Undocumented {"_csi_stack_pop", &CsiStackPop},
	// Undocumented {"_csi_stack_push_internal", &CsiStackPushInternal},
	// Undocumented {"_csi_stack_roll", &CsiStackRoll},
	// Undocumented {"_csi_translate_file", &CsiTranslateFile},
	// Undocumented {"_get_output_format", &GetOutputFormat},
	{"cairo_script_interpreter_create", &ScriptInterpreterCreate},
	{"cairo_script_interpreter_destroy", &ScriptInterpreterDestroy},
	{"cairo_script_interpreter_feed_stream", &ScriptInterpreterFeedStream},
	{"cairo_script_interpreter_feed_string", &ScriptInterpreterFeedString},
	{"cairo_script_interpreter_finish", &ScriptInterpreterFinish},
	{"cairo_script_interpreter_get_line_number", &ScriptInterpreterGetLineNumber},
	{"cairo_script_interpreter_install_hooks", &ScriptInterpreterInstallHooks},
	{"cairo_script_interpreter_reference", &ScriptInterpreterReference},
	{"cairo_script_interpreter_run", &ScriptInterpreterRun},
	{"cairo_script_interpreter_translate_stream", &ScriptInterpreterTranslateStream},
	// Undocumented {"csi_array_append", &CsiArrayAppend},
	// Undocumented {"csi_array_free", &CsiArrayFree},
	// Undocumented {"csi_array_get", &CsiArrayGet},
	// Undocumented {"csi_array_new", &CsiArrayNew},
	// Undocumented {"csi_array_put", &CsiArrayPut},
	// Undocumented {"csi_dictionary_free", &CsiDictionaryFree},
	// Undocumented {"csi_dictionary_get", &CsiDictionaryGet},
	// Undocumented {"csi_dictionary_has", &CsiDictionaryHas},
	// Undocumented {"csi_dictionary_new", &CsiDictionaryNew},
	// Undocumented {"csi_dictionary_put", &CsiDictionaryPut},
	// Undocumented {"csi_dictionary_remove", &CsiDictionaryRemove},
	// Undocumented {"csi_file_close", &CsiFileClose},
	// Undocumented {"csi_file_flush", &CsiFileFlush},
	// Undocumented {"csi_file_getc", &CsiFileGetc},
	// Undocumented {"csi_file_new", &CsiFileNew},
	// Undocumented {"csi_file_new_ascii85_decode", &CsiFileNewAscii85Decode},
	// Undocumented {"csi_file_new_deflate_decode", &CsiFileNewDeflateDecode},
	// Undocumented {"csi_file_new_for_bytes", &CsiFileNewForBytes},
	// Undocumented {"csi_file_new_for_stream", &CsiFileNewForStream},
	// Undocumented {"csi_file_new_from_string", &CsiFileNewFromString},
	// Undocumented {"csi_file_putc", &CsiFilePutc},
	// Undocumented {"csi_file_read", &CsiFileRead},
	// Undocumented {"csi_matrix_free", &CsiMatrixFree},
	// Undocumented {"csi_matrix_new", &CsiMatrixNew},
	// Undocumented {"csi_matrix_new_from_array", &CsiMatrixNewFromArray},
	// Undocumented {"csi_matrix_new_from_matrix", &CsiMatrixNewFromMatrix},
	// Undocumented {"csi_matrix_new_from_values", &CsiMatrixNewFromValues},
	// Undocumented {"csi_name_new", &CsiNameNew},
	// Undocumented {"csi_name_new_static", &CsiNameNewStatic},
	// Undocumented {"csi_object_as_file", &CsiObjectAsFile},
	// Undocumented {"csi_object_compare", &CsiObjectCompare},
	// Undocumented {"csi_object_eq", &CsiObjectEq},
	// Undocumented {"csi_object_execute", &CsiObjectExecute},
	// Undocumented {"csi_object_free", &CsiObjectFree},
	// Undocumented {"csi_object_reference", &CsiObjectReference},
	// Undocumented {"csi_string_deflate_new", &CsiStringDeflateNew},
	// Undocumented {"csi_string_free", &CsiStringFree},
	// Undocumented {"csi_string_new", &CsiStringNew},
	// Undocumented {"csi_string_new_from_bytes", &CsiStringNewFromBytes},
}
