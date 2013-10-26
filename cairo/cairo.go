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

	appendPath            func(cr *Cairo, path *Path)
	arc                   func(cr *Cairo, xc, yc, radius, angle1, angle2 float64)
	arcNegative           func(cr *Cairo, xc, yc, radius, angle1, angle2 float64)
	clip                  func(cr *Cairo)
	clipExtents           func(cr *Cairo, x1, y1, x2, y2 *float64)
	clipPreserve          func(cr *Cairo)
	closePath             func(cr *Cairo)
	copyClipRectangleList func(cr *Cairo) *RectangleList
	copyPage              func(cr *Cairo)
	copyPath              func(cr *Cairo) *Path
	copyPathFlat          func(cr *Cairo) *Path
	curveTo               func(cr *Cairo, x1, y1, x2, y2, x3, y3 float64)
	destroy               func(cr *Cairo)
	deviceToUser          func(cr *Cairo, x, y *float64)
	deviceToUserDistance  func(cr *Cairo, dx, dy *float64)
	fill                  func(cr *Cairo)
	fillExtents           func(cr *Cairo, x1, y1, x2, y2 *float64)
	fillPreserve          func(cr *Cairo)
	fontExtents           func(cr *Cairo, extents *FontExtentsS)
	getAntialias          func(cr *Cairo) Antialias
	getCurrentPoint       func(cr *Cairo, x, y *float64)
	getDash               func(cr *Cairo, dashes, offset *float64)
	getDashCount          func(cr *Cairo) int
	getFillRule           func(cr *Cairo) FillRule
	getFontFace           func(cr *Cairo) *FontFace
	getFontMatrix         func(cr *Cairo, matrix *Matrix)
	getFontOptions        func(cr *Cairo, options *FontOptions)
	getGroupTarget        func(cr *Cairo) *Surface
	getLineCap            func(cr *Cairo) LineCap
	getLineJoin           func(cr *Cairo) LineJoin
	getLineWidth          func(cr *Cairo) float64
	getMatrix             func(cr *Cairo, matrix *Matrix)
	getMiterLimit         func(cr *Cairo) float64
	getOperator           func(cr *Cairo) Operator
	getReferenceCount     func(cr *Cairo) uint
	getScaledFont         func(cr *Cairo) *ScaledFont
	getSource             func(cr *Cairo) *Pattern
	getTarget             func(cr *Cairo) *Surface
	getTolerance          func(cr *Cairo) float64
	getUserData           func(cr *Cairo, key *UserDataKey) *T.Void
	glyphExtents          func(cr *Cairo, glyphs *Glyph, numGlyphs int, extents *TextExtents)
	glyphPath             func(cr *Cairo, glyphs *Glyph, numGlyphs int)
	hasCurrentPoint       func(cr *Cairo) Bool
	identityMatrix        func(cr *Cairo)
	inClip                func(cr *Cairo, x, y float64) Bool
	inFill                func(cr *Cairo, x, y float64) Bool
	inStroke              func(cr *Cairo, x, y float64) Bool
	lineTo                func(cr *Cairo, x, y float64)
	mask                  func(cr *Cairo, pattern *Pattern)
	maskSurface           func(cr *Cairo, surface *Surface, surfaceX, surfaceY float64)
	moveTo                func(cr *Cairo, x, y float64)
	newPath               func(cr *Cairo)
	newSubPath            func(cr *Cairo)
	paint                 func(c *Cairo)
	paintWithAlpha        func(c *Cairo, alpha float64)
	pathExtents           func(c *Cairo, x1, y1, x2, y2 *float64)
	popGroup              func(c *Cairo) *Pattern
	popGroupToSource      func(c *Cairo)
	pushGroup             func(c *Cairo)
	pushGroupWithContent  func(c *Cairo, content Content)
	rectangle             func(cr *Cairo, x, y, width, height float64)
	reference             func(cr *Cairo) *Cairo
	relCurveTo            func(cr *Cairo, dx1, dy1, dx2, dy2, dx3, dy3 float64)
	relLineTo             func(cr *Cairo, dx, dy float64)
	relMoveTo             func(cr *Cairo, dx, dy float64)
	resetClip             func(cr *Cairo)
	restore               func(cr *Cairo)
	rotate                func(cr *Cairo, angle float64)
	save                  func(cr *Cairo)
	scale                 func(cr *Cairo, sx, sy float64)
	selectFontFace        func(cr *Cairo, family string, slant FontSlant, weight FontWeight)
	setAntialias          func(cr *Cairo, antialias Antialias)
	setDash               func(cr *Cairo, dashes *float64, numDashes int, offset float64)
	setFillRule           func(cr *Cairo, fillRule FillRule)
	setFontFace           func(cr *Cairo, fontFace *FontFace)
	setFontMatrix         func(cr *Cairo, matrix *Matrix)
	setFontOptions        func(cr *Cairo, options *FontOptions)
	setFontSize           func(cr *Cairo, size float64)
	setLineCap            func(cr *Cairo, lineCap LineCap)
	setLineJoin           func(cr *Cairo, lineJoin LineJoin)
	setLineWidth          func(cr *Cairo, width float64)
	setMatrix             func(cr *Cairo, matrix *Matrix)
	setMiterLimit         func(cr *Cairo, limit float64)
	setOperator           func(cr *Cairo, op Operator)
	setScaledFont         func(cr *Cairo, scaledFont *ScaledFont)
	setSource             func(cr *Cairo, source *Pattern)
	setSourceRgb          func(cr *Cairo, red, green, blue float64)
	setSourceRgba         func(cr *Cairo, red, green, blue, alpha float64)
	setSourceSurface      func(cr *Cairo, surface *Surface, x, y float64)
	setTolerance          func(cr *Cairo, tolerance float64)
	setUserData           func(cr *Cairo, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	showGlyphs            func(cr *Cairo, glyphs *Glyph, numGlyphs int)
	showPage              func(cr *Cairo)
	showText              func(cr *Cairo, utf8 string)
	showTextGlyphs        func(cr *Cairo, utf8 string, utf8Len int, glyphs *Glyph, numGlyphs int, clusters *TextCluster, numClusters int, clusterFlags TextClusterFlags)
	status                func(cr *Cairo) Status
	stroke                func(cr *Cairo)
	strokeExtents         func(cr *Cairo, x1, y1, x2, y2 *float64)
	strokePreserve        func(cr *Cairo)
	textExtents           func(cr *Cairo, utf8 string, extents *TextExtents)
	textPath              func(cr *Cairo, utf8 string)
	transform             func(cr *Cairo, matrix *Matrix)
	translate             func(cr *Cairo, tx, ty float64)
	userToDevice          func(cr *Cairo, x, y *float64)
	userToDeviceDistance  func(cr *Cairo, dx, dy *float64)
)

func (c *Cairo) AppendPath(path *Path)                      { appendPath(c, path) }
func (c *Cairo) Arc(xc, yc, radius, angle1, angle2 float64) { arc(c, xc, yc, radius, angle1, angle2) }
func (c *Cairo) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	arcNegative(c, xc, yc, radius, angle1, angle2)
}
func (c *Cairo) Clip()                                  { clip(c) }
func (c *Cairo) ClipExtents(x1, y1, x2, y2 *float64)    { clipExtents(c, x1, y1, x2, y2) }
func (c *Cairo) ClipPreserve()                          { clipPreserve(c) }
func (c *Cairo) ClosePath()                             { closePath(c) }
func (c *Cairo) CopyClipRectangleList() *RectangleList  { return copyClipRectangleList(c) }
func (c *Cairo) CopyPage()                              { copyPage(c) }
func (c *Cairo) CopyPath() *Path                        { return copyPath(c) }
func (c *Cairo) CopyPathFlat() *Path                    { return copyPathFlat(c) }
func (c *Cairo) CurveTo(x1, y1, x2, y2, x3, y3 float64) { curveTo(c, x1, y1, x2, y2, x3, y3) }
func (c *Cairo) Destroy()                               { destroy(c) }
func (c *Cairo) DeviceToUser(x, y *float64)             { deviceToUser(c, x, y) }
func (c *Cairo) DeviceToUserDistance(dx, dy *float64)   { deviceToUserDistance(c, dx, dy) }
func (c *Cairo) Fill()                                  { fill(c) }
func (c *Cairo) FillExtents(x1, y1, x2, y2 *float64)    { fillExtents(c, x1, y1, x2, y2) }
func (c *Cairo) FillPreserve()                          { fillPreserve(c) }
func (c *Cairo) FontExtents(extents *FontExtentsS)      { fontExtents(c, extents) }
func (c *Cairo) GetAntialias() Antialias                { return getAntialias(c) }
func (c *Cairo) GetCurrentPoint(x, y *float64)          { getCurrentPoint(c, x, y) }
func (c *Cairo) GetDash(dashes, offset *float64)        { getDash(c, dashes, offset) }
func (c *Cairo) GetDashCount() int                      { return getDashCount(c) }
func (c *Cairo) GetFillRule() FillRule                  { return getFillRule(c) }
func (c *Cairo) GetFontFace() *FontFace                 { return getFontFace(c) }
func (c *Cairo) GetFontMatrix(matrix *Matrix)           { getFontMatrix(c, matrix) }
func (c *Cairo) GetFontOptions(options *FontOptions)    { getFontOptions(c, options) }
func (c *Cairo) GetGroupTarget() *Surface               { return getGroupTarget(c) }
func (c *Cairo) GetLineCap() LineCap                    { return getLineCap(c) }
func (c *Cairo) GetLineJoin() LineJoin                  { return getLineJoin(c) }
func (c *Cairo) GetLineWidth() float64                  { return getLineWidth(c) }
func (c *Cairo) GetMatrix(matrix *Matrix)               { getMatrix(c, matrix) }
func (c *Cairo) GetMiterLimit() float64                 { return getMiterLimit(c) }
func (c *Cairo) GetOperator() Operator                  { return getOperator(c) }
func (c *Cairo) GetReferenceCount() uint                { return getReferenceCount(c) }
func (c *Cairo) GetScaledFont() *ScaledFont             { return getScaledFont(c) }
func (c *Cairo) GetSource() *Pattern                    { return getSource(c) }
func (c *Cairo) GetTarget() *Surface                    { return getTarget(c) }
func (c *Cairo) GetTolerance() float64                  { return getTolerance(c) }
func (c *Cairo) GetUserData(key *UserDataKey) *T.Void   { return getUserData(c, key) }
func (c *Cairo) GlyphExtents(glyphs *Glyph, numGlyphs int, extents *TextExtents) {
	glyphExtents(c, glyphs, numGlyphs, extents)
}
func (c *Cairo) GlyphPath(glyphs *Glyph, numGlyphs int) { glyphPath(c, glyphs, numGlyphs) }
func (c *Cairo) HasCurrentPoint() Bool                  { return hasCurrentPoint(c) }
func (c *Cairo) IdentityMatrix()                        { identityMatrix(c) }
func (c *Cairo) InClip(x, y float64) Bool               { return inClip(c, x, y) }
func (c *Cairo) InFill(x, y float64) Bool               { return inFill(c, x, y) }
func (c *Cairo) InStroke(x, y float64) Bool             { return inStroke(c, x, y) }
func (c *Cairo) LineTo(x, y float64)                    { lineTo(c, x, y) }
func (c *Cairo) Mask(pattern *Pattern)                  { mask(c, pattern) }
func (c *Cairo) MaskSurface(surface *Surface, surfaceX, surfaceY float64) {
	maskSurface(c, surface, surfaceX, surfaceY)
}
func (c *Cairo) MoveTo(x, y float64)                   { moveTo(c, x, y) }
func (c *Cairo) NewPath()                              { newPath(c) }
func (c *Cairo) NewSubPath()                           { newSubPath(c) }
func (c *Cairo) Paint()                                { paint(c) }
func (c *Cairo) PaintWithAlpha(alpha float64)          { paintWithAlpha(c, alpha) }
func (c *Cairo) PathExtents(x1, y1, x2, y2 *float64)   { pathExtents(c, x1, y1, x2, y2) }
func (c *Cairo) PopGroup() *Pattern                    { return popGroup(c) }
func (c *Cairo) PopGroupToSource()                     { popGroupToSource(c) }
func (c *Cairo) PushGroup()                            { pushGroup(c) }
func (c *Cairo) PushGroupWithContent(content Content)  { pushGroupWithContent(c, content) }
func (c *Cairo) Rectangle(x, y, width, height float64) { rectangle(c, x, y, width, height) }
func (c *Cairo) Reference() *Cairo                     { return reference(c) }
func (c *Cairo) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	relCurveTo(c, dx1, dy1, dx2, dy2, dx3, dy3)
}
func (c *Cairo) RelLineTo(dx, dy float64) { relLineTo(c, dx, dy) }
func (c *Cairo) RelMoveTo(dx, dy float64) { relMoveTo(c, dx, dy) }
func (c *Cairo) ResetClip()               { resetClip(c) }
func (c *Cairo) Restore()                 { restore(c) }
func (c *Cairo) Rotate(angle float64)     { rotate(c, angle) }
func (c *Cairo) Save()                    { save(c) }
func (c *Cairo) Scale(sx, sy float64)     { scale(c, sx, sy) }
func (c *Cairo) SelectFontFace(family string, slant FontSlant, weight FontWeight) {
	selectFontFace(c, family, slant, weight)
}
func (c *Cairo) SetAntialias(antialias Antialias) { setAntialias(c, antialias) }
func (c *Cairo) SetDash(dashes *float64, numDashes int, offset float64) {
	setDash(c, dashes, numDashes, offset)
}
func (c *Cairo) SetFillRule(fillRule FillRule)         { setFillRule(c, fillRule) }
func (c *Cairo) SetFontFace(fontFace *FontFace)        { setFontFace(c, fontFace) }
func (c *Cairo) SetFontMatrix(matrix *Matrix)          { setFontMatrix(c, matrix) }
func (c *Cairo) SetFontOptions(options *FontOptions)   { setFontOptions(c, options) }
func (c *Cairo) SetFontSize(size float64)              { setFontSize(c, size) }
func (c *Cairo) SetLineCap(lineCap LineCap)            { setLineCap(c, lineCap) }
func (c *Cairo) SetLineJoin(lineJoin LineJoin)         { setLineJoin(c, lineJoin) }
func (c *Cairo) SetLineWidth(width float64)            { setLineWidth(c, width) }
func (c *Cairo) SetMatrix(matrix *Matrix)              { setMatrix(c, matrix) }
func (c *Cairo) SetMiterLimit(limit float64)           { setMiterLimit(c, limit) }
func (c *Cairo) SetOperator(op Operator)               { setOperator(c, op) }
func (c *Cairo) SetScaledFont(scaledFont *ScaledFont)  { setScaledFont(c, scaledFont) }
func (c *Cairo) SetSource(source *Pattern)             { setSource(c, source) }
func (c *Cairo) SetSourceRgb(red, green, blue float64) { setSourceRgb(c, red, green, blue) }
func (c *Cairo) SetSourceRgba(red, green, blue, alpha float64) {
	setSourceRgba(c, red, green, blue, alpha)
}
func (c *Cairo) SetSourceSurface(surface *Surface, x, y float64) {
	setSourceSurface(c, surface, x, y)
}
func (c *Cairo) SetTolerance(tolerance float64) { setTolerance(c, tolerance) }
func (c *Cairo) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return setUserData(c, key, userData, destroy)
}
func (c *Cairo) ShowGlyphs(glyphs *Glyph, numGlyphs int) { showGlyphs(c, glyphs, numGlyphs) }
func (c *Cairo) ShowPage()                               { showPage(c) }
func (c *Cairo) ShowText(utf8 string)                    { showText(c, utf8) }
func (c *Cairo) ShowTextGlyphs(utf8 string, utf8Len int, glyphs *Glyph, numGlyphs int, clusters *TextCluster, numClusters int, clusterFlags TextClusterFlags) {
	showTextGlyphs(c, utf8, utf8Len, glyphs, numGlyphs, clusters, numClusters, clusterFlags)
}
func (c *Cairo) Status() Status                                { return status(c) }
func (c *Cairo) Stroke()                                       { stroke(c) }
func (c *Cairo) StrokeExtents(x1, y1, x2, y2 *float64)         { strokeExtents(c, x1, y1, x2, y2) }
func (c *Cairo) StrokePreserve()                               { strokePreserve(c) }
func (c *Cairo) TextExtents(utf8 string, extents *TextExtents) { textExtents(c, utf8, extents) }
func (c *Cairo) TextPath(utf8 string)                          { textPath(c, utf8) }
func (c *Cairo) Transform(matrix *Matrix)                      { transform(c, matrix) }
func (c *Cairo) Translate(tx, ty float64)                      { translate(c, tx, ty) }
func (c *Cairo) UserToDevice(x, y *float64)                    { userToDevice(c, x, y) }
func (c *Cairo) UserToDeviceDistance(dx, dy *float64)          { userToDeviceDistance(c, dx, dy) }

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
	{"cairo_append_path", &appendPath},
	{"cairo_arc", &arc},
	{"cairo_arc_negative", &arcNegative},
	{"cairo_clip", &clip},
	{"cairo_clip_extents", &clipExtents},
	{"cairo_clip_preserve", &clipPreserve},
	{"cairo_close_path", &closePath},
	{"cairo_copy_clip_rectangle_list", &copyClipRectangleList},
	{"cairo_copy_page", &copyPage},
	{"cairo_copy_path", &copyPath},
	{"cairo_copy_path_flat", &copyPathFlat},
	{"cairo_create", &Create},
	{"cairo_curve_to", &curveTo},
	{"cairo_debug_reset_static_data", &DebugResetStaticData},
	{"cairo_destroy", &destroy},
	{"cairo_device_acquire", &deviceAcquire},
	{"cairo_device_destroy", &deviceDestroy},
	{"cairo_device_finish", &deviceFinish},
	{"cairo_device_flush", &deviceFlush},
	{"cairo_device_get_reference_count", &deviceGetReferenceCount},
	{"cairo_device_get_type", &deviceGetType},
	{"cairo_device_get_user_data", &deviceGetUserData},
	{"cairo_device_reference", &deviceReference},
	{"cairo_device_release", &deviceRelease},
	{"cairo_device_set_user_data", &deviceSetUserData},
	{"cairo_device_status", &deviceStatus},
	{"cairo_device_to_user", &deviceToUser},
	{"cairo_device_to_user_distance", &deviceToUserDistance},
	{"cairo_fill", &fill},
	{"cairo_fill_extents", &fillExtents},
	{"cairo_fill_preserve", &fillPreserve},
	{"cairo_font_extents", &fontExtents},
	{"cairo_font_face_destroy", &fontFaceDestroy},
	{"cairo_font_face_get_reference_count", &fontFaceGetReferenceCount},
	{"cairo_font_face_get_type", &fontFaceGetType},
	{"cairo_font_face_get_user_data", &fontFaceGetUserData},
	{"cairo_font_face_reference", &fontFaceReference},
	{"cairo_font_face_set_user_data", &fontFaceSetUserData},
	{"cairo_font_face_status", &fontFaceStatus},
	{"cairo_font_options_copy", &fontOptionsCopy},
	{"cairo_font_options_create", &FontOptionsCreate},
	{"cairo_font_options_destroy", &fontOptionsDestroy},
	{"cairo_font_options_equal", &fontOptionsEqual},
	{"cairo_font_options_get_antialias", &fontOptionsGetAntialias},
	{"cairo_font_options_get_hint_metrics", &fontOptionsGetHintMetrics},
	{"cairo_font_options_get_hint_style", &fontOptionsGetHintStyle},
	{"cairo_font_options_get_subpixel_order", &fontOptionsGetSubpixelOrder},
	{"cairo_font_options_hash", &fontOptionsHash},
	{"cairo_font_options_merge", &fontOptionsMerge},
	{"cairo_font_options_set_antialias", &fontOptionsSetAntialias},
	{"cairo_font_options_set_hint_metrics", &fontOptionsSetHintMetrics},
	{"cairo_font_options_set_hint_style", &fontOptionsSetHintStyle},
	{"cairo_font_options_set_subpixel_order", &fontOptionsSetSubpixelOrder},
	{"cairo_font_options_status", &fontOptionsStatus},
	{"cairo_format_stride_for_width", &FormatStrideForWidth},
	{"cairo_ft_font_face_create_for_ft_face", &FtFontFaceCreateForFtFace},
	{"cairo_ft_font_face_create_for_pattern", &FtFontFaceCreateForPattern},
	{"cairo_ft_font_options_substitute", &FtFontOptionsSubstitute},
	{"cairo_ft_scaled_font_lock_face", &ftScaledFontLockFace},
	{"cairo_ft_scaled_font_unlock_face", &ftScaledFontUnlockFace},
	{"cairo_get_antialias", &getAntialias},
	{"cairo_get_current_point", &getCurrentPoint},
	{"cairo_get_dash", &getDash},
	{"cairo_get_dash_count", &getDashCount},
	{"cairo_get_fill_rule", &getFillRule},
	{"cairo_get_font_face", &getFontFace},
	{"cairo_get_font_matrix", &getFontMatrix},
	{"cairo_get_font_options", &getFontOptions},
	{"cairo_get_group_target", &getGroupTarget},
	{"cairo_get_line_cap", &getLineCap},
	{"cairo_get_line_join", &getLineJoin},
	{"cairo_get_line_width", &getLineWidth},
	{"cairo_get_matrix", &getMatrix},
	{"cairo_get_miter_limit", &getMiterLimit},
	{"cairo_get_operator", &getOperator},
	{"cairo_get_reference_count", &getReferenceCount},
	{"cairo_get_scaled_font", &getScaledFont},
	{"cairo_get_source", &getSource},
	{"cairo_get_target", &getTarget},
	{"cairo_get_tolerance", &getTolerance},
	{"cairo_get_user_data", &getUserData},
	{"cairo_glyph_allocate", &GlyphAllocate},
	{"cairo_glyph_extents", &glyphExtents},
	{"cairo_glyph_free", &GlyphFree},
	{"cairo_glyph_path", &glyphPath},
	{"cairo_has_current_point", &hasCurrentPoint},
	{"cairo_identity_matrix", &identityMatrix},
	{"cairo_image_surface_create", &ImageSurfaceCreate},
	{"cairo_image_surface_create_for_data", &ImageSurfaceCreateForData},
	{"cairo_image_surface_create_from_png", &ImageSurfaceCreateFromPng},
	{"cairo_image_surface_create_from_png_stream", &ImageSurfaceCreateFromPngStream},
	{"cairo_image_surface_get_data", &ImageSurfaceGetData},
	{"cairo_image_surface_get_format", &ImageSurfaceGetFormat},
	{"cairo_image_surface_get_height", &ImageSurfaceGetHeight},
	{"cairo_image_surface_get_stride", &ImageSurfaceGetStride},
	{"cairo_image_surface_get_width", &ImageSurfaceGetWidth},
	{"cairo_in_clip", &inClip},
	{"cairo_in_fill", &inFill},
	{"cairo_in_stroke", &inStroke},
	{"cairo_line_to", &lineTo},
	{"cairo_mask", &mask},
	{"cairo_mask_surface", &maskSurface},
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
	{"cairo_move_to", &moveTo},
	{"cairo_new_path", &newPath},
	{"cairo_new_sub_path", &newSubPath},
	{"cairo_paint", &paint},
	{"cairo_paint_with_alpha", &paintWithAlpha},
	{"cairo_path_destroy", &PathDestroy},
	{"cairo_path_extents", &pathExtents},
	{"cairo_pattern_add_color_stop_rgb", &patternAddColorStopRgb},
	{"cairo_pattern_add_color_stop_rgba", &patternAddColorStopRgba},
	{"cairo_pattern_create_for_surface", &PatternCreateForSurface},
	{"cairo_pattern_create_linear", &PatternCreateLinear},
	{"cairo_pattern_create_radial", &PatternCreateRadial},
	{"cairo_pattern_create_rgb", &PatternCreateRgb},
	{"cairo_pattern_create_rgba", &PatternCreateRgba},
	{"cairo_pattern_destroy", &patternDestroy},
	{"cairo_pattern_get_color_stop_count", &patternGetColorStopCount},
	{"cairo_pattern_get_color_stop_rgba", &patternGetColorStopRgba},
	{"cairo_pattern_get_extend", &patternGetExtend},
	{"cairo_pattern_get_filter", &patternGetFilter},
	{"cairo_pattern_get_linear_points", &patternGetLinearPoints},
	{"cairo_pattern_get_matrix", &patternGetMatrix},
	{"cairo_pattern_get_radial_circles", &patternGetRadialCircles},
	{"cairo_pattern_get_reference_count", &patternGetReferenceCount},
	{"cairo_pattern_get_rgba", &patternGetRgba},
	{"cairo_pattern_get_surface", &patternGetSurface},
	{"cairo_pattern_get_type", &patternGetType},
	{"cairo_pattern_get_user_data", &patternGetUserData},
	{"cairo_pattern_reference", &patternReference},
	{"cairo_pattern_set_extend", &patternSetExtend},
	{"cairo_pattern_set_filter", &patternSetFilter},
	{"cairo_pattern_set_matrix", &patternSetMatrix},
	{"cairo_pattern_set_user_data", &patternSetUserData},
	{"cairo_pattern_status", &patternStatus},
	{"cairo_pdf_get_versions", &PdfGetVersions},
	{"cairo_pdf_surface_create", &PdfSurfaceCreate},
	{"cairo_pdf_surface_create_for_stream", &PdfSurfaceCreateForStream},
	{"cairo_pdf_surface_restrict_to_version", &PdfSurfaceRestrictToVersion},
	{"cairo_pdf_surface_set_size", &PdfSurfaceSetSize},
	{"cairo_pdf_version_to_string", &PdfVersionToString},
	{"cairo_pop_group", &popGroup},
	{"cairo_pop_group_to_source", &popGroupToSource},
	{"cairo_ps_get_levels", &PsGetLevels},
	{"cairo_ps_level_to_string", &PsLevelToString},
	{"cairo_ps_surface_create", &PsSurfaceCreate},
	{"cairo_ps_surface_create_for_stream", &PsSurfaceCreateForStream},
	{"cairo_ps_surface_dsc_begin_page_setup", &psSurfaceDscBeginPageSetup},
	{"cairo_ps_surface_dsc_begin_setup", &psSurfaceDscBeginSetup},
	{"cairo_ps_surface_dsc_comment", &psSurfaceDscComment},
	{"cairo_ps_surface_get_eps", &psSurfaceGetEps},
	{"cairo_ps_surface_restrict_to_level", &psSurfaceRestrictToLevel},
	{"cairo_ps_surface_set_eps", &psSurfaceSetEps},
	{"cairo_ps_surface_set_size", &psSurfaceSetSize},
	{"cairo_push_group", &pushGroup},
	{"cairo_push_group_with_content", &pushGroupWithContent},
	{"cairo_recording_surface_create", &RecordingSurfaceCreate},
	{"cairo_recording_surface_ink_extents", &RecordingSurfaceInkExtents},
	{"cairo_rectangle", &rectangle},
	{"cairo_rectangle_list_destroy", &rectangleListDestroy},
	{"cairo_reference", &reference},
	{"cairo_region_contains_point", &regionContainsPoint},
	{"cairo_region_contains_rectangle", &regionContainsRectangle},
	{"cairo_region_copy", &regionCopy},
	{"cairo_region_create", &RegionCreate},
	{"cairo_region_create_rectangle", &RegionCreateRectangle},
	{"cairo_region_create_rectangles", &RegionCreateRectangles},
	{"cairo_region_destroy", &regionDestroy},
	{"cairo_region_equal", &regionEqual},
	{"cairo_region_get_extents", &regionGetExtents},
	{"cairo_region_get_rectangle", &regionGetRectangle},
	{"cairo_region_intersect", &regionIntersect},
	{"cairo_region_intersect_rectangle", &regionIntersectRectangle},
	{"cairo_region_is_empty", &regionIsEmpty},
	{"cairo_region_num_rectangles", &regionNumRectangles},
	{"cairo_region_reference", &regionReference},
	{"cairo_region_status", &regionStatus},
	{"cairo_region_subtract", &regionSubtract},
	{"cairo_region_subtract_rectangle", &regionSubtractRectangle},
	{"cairo_region_translate", &regionTranslate},
	{"cairo_region_union", &regionUnion},
	{"cairo_region_union_rectangle", &regionUnionRectangle},
	{"cairo_region_xor", &regionXor},
	{"cairo_region_xor_rectangle", &regionXorRectangle},
	{"cairo_rel_curve_to", &relCurveTo},
	{"cairo_rel_line_to", &relLineTo},
	{"cairo_rel_move_to", &relMoveTo},
	{"cairo_reset_clip", &resetClip},
	{"cairo_restore", &restore},
	{"cairo_rotate", &rotate},
	{"cairo_save", &save},
	{"cairo_scale", &scale},
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
	{"cairo_select_font_face", &selectFontFace},
	{"cairo_set_antialias", &setAntialias},
	{"cairo_set_dash", &setDash},
	{"cairo_set_fill_rule", &setFillRule},
	{"cairo_set_font_face", &setFontFace},
	{"cairo_set_font_matrix", &setFontMatrix},
	{"cairo_set_font_options", &setFontOptions},
	{"cairo_set_font_size", &setFontSize},
	{"cairo_set_line_cap", &setLineCap},
	{"cairo_set_line_join", &setLineJoin},
	{"cairo_set_line_width", &setLineWidth},
	{"cairo_set_matrix", &setMatrix},
	{"cairo_set_miter_limit", &setMiterLimit},
	{"cairo_set_operator", &setOperator},
	{"cairo_set_scaled_font", &setScaledFont},
	{"cairo_set_source", &setSource},
	{"cairo_set_source_rgb", &setSourceRgb},
	{"cairo_set_source_rgba", &setSourceRgba},
	{"cairo_set_source_surface", &setSourceSurface},
	{"cairo_set_tolerance", &setTolerance},
	{"cairo_set_user_data", &setUserData},
	{"cairo_show_glyphs", &showGlyphs},
	{"cairo_show_page", &showPage},
	{"cairo_show_text", &showText},
	{"cairo_show_text_glyphs", &showTextGlyphs},
	{"cairo_status", &status},
	{"cairo_status_to_string", &StatusToString},
	{"cairo_stroke", &stroke},
	{"cairo_stroke_extents", &strokeExtents},
	{"cairo_stroke_preserve", &strokePreserve},
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
	{"cairo_text_cluster_free", &textClusterFree},
	{"cairo_text_extents", &textExtents},
	{"cairo_text_path", &textPath},
	{"cairo_toy_font_face_create", &ToyFontFaceCreate},
	{"cairo_toy_font_face_get_family", &toyFontFaceGetFamily},
	{"cairo_toy_font_face_get_slant", &toyFontFaceGetSlant},
	{"cairo_toy_font_face_get_weight", &toyFontFaceGetWeight},
	{"cairo_transform", &transform},
	{"cairo_translate", &translate},
	{"cairo_user_font_face_create", &UserFontFaceCreate},
	{"cairo_user_font_face_get_init_func", &UserFontFaceGetInitFunc},
	{"cairo_user_font_face_get_render_glyph_func", &UserFontFaceGetRenderGlyphFunc},
	{"cairo_user_font_face_get_text_to_glyphs_func", &UserFontFaceGetTextToGlyphsFunc},
	{"cairo_user_font_face_get_unicode_to_glyph_func", &UserFontFaceGetUnicodeToGlyphFunc},
	{"cairo_user_font_face_set_init_func", &UserFontFaceSetInitFunc},
	{"cairo_user_font_face_set_render_glyph_func", &UserFontFaceSetRenderGlyphFunc},
	{"cairo_user_font_face_set_text_to_glyphs_func", &UserFontFaceSetTextToGlyphsFunc},
	{"cairo_user_font_face_set_unicode_to_glyph_func", &UserFontFaceSetUnicodeToGlyphFunc},
	{"cairo_user_to_device", &userToDevice},
	{"cairo_user_to_device_distance", &userToDeviceDistance},
	{"cairo_version", &Version},
	{"cairo_version_string", &VersionString},
	{"cairo_win32_font_face_create_for_hfont", &win32FontFaceCreateForHfont},
	{"cairo_win32_font_face_create_for_logfontw", &win32FontFaceCreateForLogfontw},
	{"cairo_win32_font_face_create_for_logfontw_hfont", &win32FontFaceCreateForLogfontwHfont},
	{"cairo_win32_printing_surface_create", &win32PrintingSurfaceCreate},
	{"cairo_win32_scaled_font_done_font", &win32ScaledFontDoneFont},
	{"cairo_win32_scaled_font_get_device_to_logical", &win32ScaledFontGetDeviceToLogical},
	{"cairo_win32_scaled_font_get_logical_to_device", &win32ScaledFontGetLogicalToDevice},
	{"cairo_win32_scaled_font_get_metrics_factor", &win32ScaledFontGetMetricsFactor},
	{"cairo_win32_scaled_font_select_font", &win32ScaledFontSelectFont},
	{"cairo_win32_surface_create", &win32SurfaceCreate},
	{"cairo_win32_surface_create_with_ddb", &win32SurfaceCreateWithDdb},
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
