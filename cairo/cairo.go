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

	Bool int
	Enum int
)

type (
	Cairo struct{}

	UserDataKey struct {
		Unused int
	}

	Glyph struct {
		Index T.UnsignedLong
		X     float64
		Y     float64
	}

	DestroyFunc func(data *T.Void)

	ReadFunc func(
		closure *T.Void, data *T.UnsignedChar, length uint) Status

	WriteFunc func(
		closure *T.Void, data *T.UnsignedChar, length uint) Status
)

type Antialias Enum

const (
	ANTIALIAS_DEFAULT Antialias = iota
	ANTIALIAS_NONE
	ANTIALIAS_GRAY
	ANTIALIAS_SUBPIXEL
)

type Extend Enum

const (
	EXTEND_NONE Extend = iota
	EXTEND_REPEAT
	EXTEND_REFLECT
	EXTEND_PAD
)

type HintMetrics Enum

const (
	HINT_METRICS_DEFAULT HintMetrics = iota
	HINT_METRICS_OFF
	HINT_METRICS_ON
)

type HintStyle Enum

const (
	HINT_STYLE_DEFAULT HintStyle = iota
	HINT_STYLE_NONE
	HINT_STYLE_SLIGHT
	HINT_STYLE_MEDIUM
	HINT_STYLE_FULL
)

type LineCap Enum

const (
	LINE_CAP_BUTT LineCap = iota
	LINE_CAP_ROUND
	LINE_CAP_SQUARE
)

type LineJoin Enum

const (
	LINE_JOIN_MITER LineJoin = iota
	LINE_JOIN_ROUND
	LINE_JOIN_BEVEL
)

type Operator Enum

const (
	OPERATOR_CLEAR Operator = iota
	OPERATOR_SOURCE
	OPERATOR_OVER
	OPERATOR_IN
	OPERATOR_OUT
	OPERATOR_ATOP
	OPERATOR_DEST
	OPERATOR_DEST_OVER
	OPERATOR_DEST_IN
	OPERATOR_DEST_OUT
	OPERATOR_DEST_ATOP
	OPERATOR_XOR
	OPERATOR_ADD
	OPERATOR_SATURATE
	OPERATOR_MULTIPLY
	OPERATOR_SCREEN
	OPERATOR_OVERLAY
	OPERATOR_DARKEN
	OPERATOR_LIGHTEN
	OPERATOR_COLOR_DODGE
	OPERATOR_COLOR_BURN
	OPERATOR_HARD_LIGHT
	OPERATOR_SOFT_LIGHT
	OPERATOR_DIFFERENCE
	OPERATOR_EXCLUSION
	OPERATOR_HSL_HUE
	OPERATOR_HSL_SATURATION
	OPERATOR_HSL_COLOR
	OPERATOR_HSL_LUMINOSITY
)

type SubpixelOrder Enum

const (
	SUBPIXEL_ORDER_DEFAULT SubpixelOrder = iota
	SUBPIXEL_ORDER_RGB
	SUBPIXEL_ORDER_BGR
	SUBPIXEL_ORDER_VRGB
	SUBPIXEL_ORDER_VBGR
)

var (
	Create func(target *Surface) *Cairo

	AppendPath            func(cr *Cairo, path *Path)
	Arc                   func(cr *Cairo, xc, yc, radius, angle1, angle2 float64)
	ArcNegative           func(cr *Cairo, xc, yc, radius, angle1, angle2 float64)
	Clip                  func(cr *Cairo)
	ClipExtents           func(cr *Cairo, x1, y1, x2, y2 *float64)
	ClipPreserve          func(cr *Cairo)
	ClosePath             func(cr *Cairo)
	CopyClipRectangleList func(cr *Cairo) *RectangleList
	CopyPage              func(cr *Cairo)
	CopyPath              func(cr *Cairo) *Path
	CopyPathFlat          func(cr *Cairo) *Path
	CurveTo               func(cr *Cairo, x1, y1, x2, y2, x3, y3 float64)
	Destroy               func(cr *Cairo)
	DeviceToUser          func(cr *Cairo, x, y *float64)
	DeviceToUserDistance  func(cr *Cairo, dx, dy *float64)
	Fill                  func(cr *Cairo)
	FillExtents           func(cr *Cairo, x1, y1, x2, y2 *float64)
	FillPreserve          func(cr *Cairo)
	FontExtents           func(cr *Cairo, extents *FontExtentsS)
	GetAntialias          func(cr *Cairo) Antialias
	GetCurrentPoint       func(cr *Cairo, x, y *float64)
	GetDash               func(cr *Cairo, dashes, offset *float64)
	GetDashCount          func(cr *Cairo) int
	GetFillRule           func(cr *Cairo) FillRule
	GetFontFace           func(cr *Cairo) *FontFace
	GetFontMatrix         func(cr *Cairo, matrix *Matrix)
	GetFontOptions        func(cr *Cairo, options *FontOptions)
	GetGroupTarget        func(cr *Cairo) *Surface
	GetLineCap            func(cr *Cairo) LineCap
	GetLineJoin           func(cr *Cairo) LineJoin
	GetLineWidth          func(cr *Cairo) float64
	GetMatrix             func(cr *Cairo, matrix *Matrix)
	GetMiterLimit         func(cr *Cairo) float64
	GetOperator           func(cr *Cairo) Operator
	GetReferenceCount     func(cr *Cairo) uint
	GetScaledFont         func(cr *Cairo) *ScaledFont
	GetSource             func(cr *Cairo) *Pattern
	GetTarget             func(cr *Cairo) *Surface
	GetTolerance          func(cr *Cairo) float64
	GetUserData           func(cr *Cairo, key *UserDataKey) *T.Void
	GlyphExtents          func(cr *Cairo, glyphs *Glyph, numGlyphs int, extents *TextExtents)
	GlyphPath             func(cr *Cairo, glyphs *Glyph, numGlyphs int)
	HasCurrentPoint       func(cr *Cairo) Bool
	IdentityMatrix        func(cr *Cairo)
	InClip                func(cr *Cairo, x, y float64) Bool
	InFill                func(cr *Cairo, x, y float64) Bool
	InStroke              func(cr *Cairo, x, y float64) Bool
	LineTo                func(cr *Cairo, x, y float64)
	Mask                  func(cr *Cairo, pattern *Pattern)
	MaskSurface           func(cr *Cairo, surface *Surface, surfaceX, surfaceY float64)
	MoveTo                func(cr *Cairo, x, y float64)
	NewPath               func(cr *Cairo)
	NewSubPath            func(cr *Cairo)
	Paint                 func(c *Cairo)
	PaintWithAlpha        func(c *Cairo, alpha float64)
	PathExtents           func(c *Cairo, x1, y1, x2, y2 *float64)
	PopGroup              func(c *Cairo) *Pattern
	PopGroupToSource      func(c *Cairo)
	PushGroup             func(c *Cairo)
	PushGroupWithContent  func(c *Cairo, content Content)
	Rectangle_            func(cr *Cairo, x, y, width, height float64)
	Reference             func(cr *Cairo) *Cairo
	RelCurveTo            func(cr *Cairo, dx1, dy1, dx2, dy2, dx3, dy3 float64)
	RelLineTo             func(cr *Cairo, dx, dy float64)
	RelMoveTo             func(cr *Cairo, dx, dy float64)
	ResetClip             func(cr *Cairo)
	Restore               func(cr *Cairo)
	Rotate                func(cr *Cairo, angle float64)
	Save                  func(cr *Cairo)
	Scale                 func(cr *Cairo, sx, sy float64)
	SelectFontFace        func(cr *Cairo, family string, slant FontSlant, weight FontWeight)
	SetAntialias          func(cr *Cairo, antialias Antialias)
	SetDash               func(cr *Cairo, dashes *float64, numDashes int, offset float64)
	SetFillRule           func(cr *Cairo, fillRule FillRule)
	SetFontFace           func(cr *Cairo, fontFace *FontFace)
	SetFontMatrix         func(cr *Cairo, matrix *Matrix)
	SetFontOptions        func(cr *Cairo, options *FontOptions)
	SetFontSize           func(cr *Cairo, size float64)
	SetLineCap            func(cr *Cairo, lineCap LineCap)
	SetLineJoin           func(cr *Cairo, lineJoin LineJoin)
	SetLineWidth          func(cr *Cairo, width float64)
	SetMatrix             func(cr *Cairo, matrix *Matrix)
	SetMiterLimit         func(cr *Cairo, limit float64)
	SetOperator           func(cr *Cairo, op Operator)
	SetScaledFont         func(cr *Cairo, scaledFont *ScaledFont)
	SetSource             func(cr *Cairo, source *Pattern)
	SetSourceRgb          func(cr *Cairo, red, green, blue float64)
	SetSourceRgba         func(cr *Cairo, red, green, blue, alpha float64)
	SetSourceSurface      func(cr *Cairo, surface *Surface, x, y float64)
	SetTolerance          func(cr *Cairo, tolerance float64)
	SetUserData           func(cr *Cairo, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	ShowGlyphs            func(cr *Cairo, glyphs *Glyph, numGlyphs int)
	ShowPage              func(cr *Cairo)
	ShowText              func(cr *Cairo, utf8 string)
	ShowTextGlyphs        func(cr *Cairo, utf8 string, utf8Len int, glyphs *Glyph, numGlyphs int, clusters *TextCluster, numClusters int, clusterFlags TextClusterFlags)
	Status_               func(cr *Cairo) Status
	Stroke                func(cr *Cairo)
	StrokeExtents         func(cr *Cairo, x1, y1, x2, y2 *float64)
	StrokePreserve        func(cr *Cairo)
	TextExtents_          func(cr *Cairo, utf8 string, extents *TextExtents)
	TextPath              func(cr *Cairo, utf8 string)
	Transform             func(cr *Cairo, matrix *Matrix)
	Translate             func(cr *Cairo, tx, ty float64)
	UserToDevice          func(cr *Cairo, x, y *float64)
	UserToDeviceDistance  func(cr *Cairo, dx, dy *float64)
)

func (c *Cairo) AppendPath(path *Path)                      { AppendPath(c, path) }
func (c *Cairo) Arc(xc, yc, radius, angle1, angle2 float64) { Arc(c, xc, yc, radius, angle1, angle2) }
func (c *Cairo) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	ArcNegative(c, xc, yc, radius, angle1, angle2)
}
func (c *Cairo) Clip()                                  { Clip(c) }
func (c *Cairo) ClipExtents(x1, y1, x2, y2 *float64)    { ClipExtents(c, x1, y1, x2, y2) }
func (c *Cairo) ClipPreserve()                          { ClipPreserve(c) }
func (c *Cairo) ClosePath()                             { ClosePath(c) }
func (c *Cairo) CopyClipRectangleList() *RectangleList  { return CopyClipRectangleList(c) }
func (c *Cairo) CopyPage()                              { CopyPage(c) }
func (c *Cairo) CopyPath() *Path                        { return CopyPath(c) }
func (c *Cairo) CopyPathFlat() *Path                    { return CopyPathFlat(c) }
func (c *Cairo) CurveTo(x1, y1, x2, y2, x3, y3 float64) { CurveTo(c, x1, y1, x2, y2, x3, y3) }
func (c *Cairo) Destroy()                               { Destroy(c) }
func (c *Cairo) DeviceToUser(x, y *float64)             { DeviceToUser(c, x, y) }
func (c *Cairo) DeviceToUserDistance(dx, dy *float64)   { DeviceToUserDistance(c, dx, dy) }
func (c *Cairo) Fill()                                  { Fill(c) }
func (c *Cairo) FillExtents(x1, y1, x2, y2 *float64)    { FillExtents(c, x1, y1, x2, y2) }
func (c *Cairo) FillPreserve()                          { FillPreserve(c) }
func (c *Cairo) FontExtents(extents *FontExtentsS)      { FontExtents(c, extents) }
func (c *Cairo) GetAntialias() Antialias                { return GetAntialias(c) }
func (c *Cairo) GetCurrentPoint(x, y *float64)          { GetCurrentPoint(c, x, y) }
func (c *Cairo) GetDash(dashes, offset *float64)        { GetDash(c, dashes, offset) }
func (c *Cairo) GetDashCount() int                      { return GetDashCount(c) }
func (c *Cairo) GetFillRule() FillRule                  { return GetFillRule(c) }
func (c *Cairo) GetFontFace() *FontFace                 { return GetFontFace(c) }
func (c *Cairo) GetFontMatrix(matrix *Matrix)           { GetFontMatrix(c, matrix) }
func (c *Cairo) GetFontOptions(options *FontOptions)    { GetFontOptions(c, options) }
func (c *Cairo) GetGroupTarget() *Surface               { return GetGroupTarget(c) }
func (c *Cairo) GetLineCap() LineCap                    { return GetLineCap(c) }
func (c *Cairo) GetLineJoin() LineJoin                  { return GetLineJoin(c) }
func (c *Cairo) GetLineWidth() float64                  { return GetLineWidth(c) }
func (c *Cairo) GetMatrix(matrix *Matrix)               { GetMatrix(c, matrix) }
func (c *Cairo) GetMiterLimit() float64                 { return GetMiterLimit(c) }
func (c *Cairo) GetOperator() Operator                  { return GetOperator(c) }
func (c *Cairo) GetReferenceCount() uint                { return GetReferenceCount(c) }
func (c *Cairo) GetScaledFont() *ScaledFont             { return GetScaledFont(c) }
func (c *Cairo) GetSource() *Pattern                    { return GetSource(c) }
func (c *Cairo) GetTarget() *Surface                    { return GetTarget(c) }
func (c *Cairo) GetTolerance() float64                  { return GetTolerance(c) }
func (c *Cairo) GetUserData(key *UserDataKey) *T.Void   { return GetUserData(c, key) }
func (c *Cairo) GlyphExtents(glyphs *Glyph, numGlyphs int, extents *TextExtents) {
	GlyphExtents(c, glyphs, numGlyphs, extents)
}
func (c *Cairo) GlyphPath(glyphs *Glyph, numGlyphs int) { GlyphPath(c, glyphs, numGlyphs) }
func (c *Cairo) HasCurrentPoint() Bool                  { return HasCurrentPoint(c) }
func (c *Cairo) IdentityMatrix()                        { IdentityMatrix(c) }
func (c *Cairo) InClip(x, y float64) Bool               { return InClip(c, x, y) }
func (c *Cairo) InFill(x, y float64) Bool               { return InFill(c, x, y) }
func (c *Cairo) InStroke(x, y float64) Bool             { return InStroke(c, x, y) }
func (c *Cairo) LineTo(x, y float64)                    { LineTo(c, x, y) }
func (c *Cairo) Mask(pattern *Pattern)                  { Mask(c, pattern) }
func (c *Cairo) MaskSurface(surface *Surface, surfaceX, surfaceY float64) {
	MaskSurface(c, surface, surfaceX, surfaceY)
}
func (c *Cairo) MoveTo(x, y float64)                   { MoveTo(c, x, y) }
func (c *Cairo) NewPath()                              { NewPath(c) }
func (c *Cairo) NewSubPath()                           { NewSubPath(c) }
func (c *Cairo) Paint()                                { Paint(c) }
func (c *Cairo) PaintWithAlpha(alpha float64)          { PaintWithAlpha(c, alpha) }
func (c *Cairo) PathExtents(x1, y1, x2, y2 *float64)   { PathExtents(c, x1, y1, x2, y2) }
func (c *Cairo) PopGroup() *Pattern                    { return PopGroup(c) }
func (c *Cairo) PopGroupToSource()                     { PopGroupToSource(c) }
func (c *Cairo) PushGroup()                            { PushGroup(c) }
func (c *Cairo) PushGroupWithContent(content Content)  { PushGroupWithContent(c, content) }
func (c *Cairo) Rectangle(x, y, width, height float64) { Rectangle_(c, x, y, width, height) }
func (c *Cairo) Reference() *Cairo                     { return Reference(c) }
func (c *Cairo) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	RelCurveTo(c, dx1, dy1, dx2, dy2, dx3, dy3)
}
func (c *Cairo) RelLineTo(dx, dy float64) { RelLineTo(c, dx, dy) }
func (c *Cairo) RelMoveTo(dx, dy float64) { RelMoveTo(c, dx, dy) }
func (c *Cairo) ResetClip()               { ResetClip(c) }
func (c *Cairo) Restore()                 { Restore(c) }
func (c *Cairo) Rotate(angle float64)     { Rotate(c, angle) }
func (c *Cairo) Save()                    { Save(c) }
func (c *Cairo) Scale(sx, sy float64)     { Scale(c, sx, sy) }
func (c *Cairo) SelectFontFace(family string, slant FontSlant, weight FontWeight) {
	SelectFontFace(c, family, slant, weight)
}
func (c *Cairo) SetAntialias(antialias Antialias) { SetAntialias(c, antialias) }
func (c *Cairo) SetDash(dashes *float64, numDashes int, offset float64) {
	SetDash(c, dashes, numDashes, offset)
}
func (c *Cairo) SetFillRule(fillRule FillRule)         { SetFillRule(c, fillRule) }
func (c *Cairo) SetFontFace(fontFace *FontFace)        { SetFontFace(c, fontFace) }
func (c *Cairo) SetFontMatrix(matrix *Matrix)          { SetFontMatrix(c, matrix) }
func (c *Cairo) SetFontOptions(options *FontOptions)   { SetFontOptions(c, options) }
func (c *Cairo) SetFontSize(size float64)              { SetFontSize(c, size) }
func (c *Cairo) SetLineCap(lineCap LineCap)            { SetLineCap(c, lineCap) }
func (c *Cairo) SetLineJoin(lineJoin LineJoin)         { SetLineJoin(c, lineJoin) }
func (c *Cairo) SetLineWidth(width float64)            { SetLineWidth(c, width) }
func (c *Cairo) SetMatrix(matrix *Matrix)              { SetMatrix(c, matrix) }
func (c *Cairo) SetMiterLimit(limit float64)           { SetMiterLimit(c, limit) }
func (c *Cairo) SetOperator(op Operator)               { SetOperator(c, op) }
func (c *Cairo) SetScaledFont(scaledFont *ScaledFont)  { SetScaledFont(c, scaledFont) }
func (c *Cairo) SetSource(source *Pattern)             { SetSource(c, source) }
func (c *Cairo) SetSourceRgb(red, green, blue float64) { SetSourceRgb(c, red, green, blue) }
func (c *Cairo) SetSourceRgba(red, green, blue, alpha float64) {
	SetSourceRgba(c, red, green, blue, alpha)
}
func (c *Cairo) SetSourceSurface(surface *Surface, x, y float64) {
	SetSourceSurface(c, surface, x, y)
}
func (c *Cairo) SetTolerance(tolerance float64) { SetTolerance(c, tolerance) }
func (c *Cairo) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return SetUserData(c, key, userData, destroy)
}
func (c *Cairo) ShowGlyphs(glyphs *Glyph, numGlyphs int) { ShowGlyphs(c, glyphs, numGlyphs) }
func (c *Cairo) ShowPage()                               { ShowPage(c) }
func (c *Cairo) ShowText(utf8 string)                    { ShowText(c, utf8) }
func (c *Cairo) ShowTextGlyphs(utf8 string, utf8Len int, glyphs *Glyph, numGlyphs int, clusters *TextCluster, numClusters int, clusterFlags TextClusterFlags) {
	ShowTextGlyphs(c, utf8, utf8Len, glyphs, numGlyphs, clusters, numClusters, clusterFlags)
}
func (c *Cairo) Status() Status                                { return Status_(c) }
func (c *Cairo) Stroke()                                       { Stroke(c) }
func (c *Cairo) StrokeExtents(x1, y1, x2, y2 *float64)         { StrokeExtents(c, x1, y1, x2, y2) }
func (c *Cairo) StrokePreserve()                               { StrokePreserve(c) }
func (c *Cairo) TextExtents(utf8 string, extents *TextExtents) { TextExtents_(c, utf8, extents) }
func (c *Cairo) TextPath(utf8 string)                          { TextPath(c, utf8) }
func (c *Cairo) Transform(matrix *Matrix)                      { Transform(c, matrix) }
func (c *Cairo) Translate(tx, ty float64)                      { Translate(c, tx, ty) }
func (c *Cairo) UserToDevice(x, y *float64)                    { UserToDevice(c, x, y) }
func (c *Cairo) UserToDeviceDistance(dx, dy *float64)          { UserToDeviceDistance(c, dx, dy) }

var (
	Version func() int

	VersionString func() string

	GlyphAllocate func(numGlyphs int) *Glyph
	GlyphFree     func(glyphs *Glyph)

	GobjectContextGetType          func()
	GobjectDeviceGetType           func()
	GobjectPatternGetType          func()
	GobjectSurfaceGetType          func()
	GobjectRectangleGetType        func()
	GobjectScaledFontGetType       func()
	GobjectFontFaceGetType         func()
	GobjectFontOptionsGetType      func()
	GobjectRectangleIntGetType     func()
	GobjectRegionGetType           func()
	GobjectStatusGetType           func()
	GobjectContentGetType          func()
	GobjectOperatorGetType         func()
	GobjectAntialiasGetType        func()
	GobjectFillRuleGetType         func()
	GobjectLineCapGetType          func()
	GobjectLineJoinGetType         func()
	GobjectTextClusterFlagsGetType func()
	GobjectFontSlantGetType        func()
	GobjectFontWeightGetType       func()
	GobjectSubpixelOrderGetType    func()
	GobjectHintStyleGetType        func()
	GobjectHintMetricsGetType      func()
	GobjectFontTypeGetType         func()
	GobjectPathDataTypeGetType     func()
	GobjectDeviceTypeGetType       func()
	GobjectSurfaceTypeGetType      func()
	GobjectFormatGetType           func()
	GobjectPatternTypeGetType      func()
	GobjectExtendGetType           func()
	GobjectFilterGetType           func()
	GobjectRegionOverlapGetType    func()
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
	{"cairo_rectangle", &Rectangle_},
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
	{"cairo_status", &Status_},
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
	{"cairo_text_extents", &TextExtents_},
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
	{"cairo_win32_surface_create_with_dib", &SurfaceCreateWithDib},
	{"cairo_win32_surface_get_dc", &SurfaceGetDc},
	{"cairo_win32_surface_get_image", &SurfaceGetImage},
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
