// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
	T "github.com/tHinqa/outside-gtk2/types"
)

type Path struct {
	Status   Status
	Data     *PathData
	Num_data int
}

type PathData struct {
	// todo(t): Union
	// header struct {
	// 	Type   CairoPathDataType
	// 	length int
	// }
	Point struct {
		X, Y float64
	}
}

type PathDataType Enum

const (
	PATH_MOVE_TO PathDataType = iota
	PATH_LINE_TO
	PATH_CURVE_TO
	PATH_CLOSE_PATH
)

var PathDestroy func(path *Path)

type Pattern struct{}

type Surface struct{}

var (
	PatternCreateForSurface func(surface *Surface) *Pattern
	PatternCreateLinear     func(x0, y0, x1, y1 float64) *Pattern
	PatternCreateRadial     func(cx0, cy0, radius0, cx1, cy1, radius1 float64) *Pattern
	PatternCreateRgb        func(red, green, blue float64) *Pattern
	PatternCreateRgba       func(red, green, blue, alpha float64) *Pattern

	patternAddColorStopRgb   func(p *Pattern, offset, red, green, blue float64)
	patternAddColorStopRgba  func(p *Pattern, offset, red, green, blue, alpha float64)
	patternDestroy           func(p *Pattern)
	patternGetColorStopCount func(p *Pattern, count *int) Status
	patternGetColorStopRgba  func(p *Pattern, index int, offset, red, green, blue, alpha *float64) Status
	patternGetExtend         func(p *Pattern) Extend
	patternGetFilter         func(p *Pattern) Filter
	patternGetLinearPoints   func(p *Pattern, x0, y0, x1, y1 *float64) Status
	patternGetMatrix         func(p *Pattern, matrix *Matrix)
	patternGetRadialCircles  func(p *Pattern, x0, y0, r0, x1, y1, r1 *float64) Status
	patternGetReferenceCount func(p *Pattern) uint
	patternGetRgba           func(p *Pattern, red, green, blue, alpha *float64) Status
	patternGetSurface        func(p *Pattern, surface **Surface) Status
	patternGetType           func(p *Pattern) PatternType
	patternGetUserData       func(p *Pattern, key *UserDataKey) *T.Void
	patternReference         func(p *Pattern) *Pattern
	patternSetExtend         func(p *Pattern, extend Extend)
	patternSetFilter         func(p *Pattern, filter Filter)
	patternSetMatrix         func(p *Pattern, matrix *Matrix)
	patternSetUserData       func(p *Pattern, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	patternStatus            func(p *Pattern) Status
)

func (p *Pattern) AddColorStopRgb(offset, red, green, blue float64) {
	patternAddColorStopRgb(p, offset, red, green, blue)
}
func (p *Pattern) AddColorStopRgba(offset, red, green, blue, alpha float64) {
	patternAddColorStopRgba(p, offset, red, green, blue, alpha)
}
func (p *Pattern) Destroy() { patternDestroy(p) }
func (p *Pattern) GetColorStopCount(count *int) Status {
	return patternGetColorStopCount(p, count)
}
func (p *Pattern) GetColorStopRgba(index int, offset, red, green, blue, alpha *float64) Status {
	return patternGetColorStopRgba(p, index, offset, red, green, blue, alpha)
}
func (p *Pattern) GetExtend() Extend { return patternGetExtend(p) }
func (p *Pattern) GetFilter() Filter { return patternGetFilter(p) }
func (p *Pattern) GetLinearPoints(x0, y0, x1, y1 *float64) Status {
	return patternGetLinearPoints(p, x0, y0, x1, y1)
}
func (p *Pattern) GetMatrix(matrix *Matrix) { patternGetMatrix(p, matrix) }
func (p *Pattern) GetRadialCircles(x0, y0, r0, x1, y1, r1 *float64) Status {
	return patternGetRadialCircles(p, x0, y0, r0, x1, y1, r1)
}
func (p *Pattern) GetReferenceCount() uint { return patternGetReferenceCount(p) }
func (p *Pattern) GetRgba(red, green, blue, alpha *float64) Status {
	return patternGetRgba(p, red, green, blue, alpha)
}
func (p *Pattern) GetSurface(surface **Surface) Status  { return patternGetSurface(p, surface) }
func (p *Pattern) GetType() PatternType                 { return patternGetType(p) }
func (p *Pattern) GetUserData(key *UserDataKey) *T.Void { return patternGetUserData(p, key) }
func (p *Pattern) Reference() *Pattern                  { return patternReference(p) }
func (p *Pattern) SetExtend(extend Extend)              { patternSetExtend(p, extend) }
func (p *Pattern) SetFilter(filter Filter)              { patternSetFilter(p, filter) }
func (p *Pattern) SetMatrix(matrix *Matrix)             { patternSetMatrix(p, matrix) }
func (p *Pattern) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return patternSetUserData(p, key, userData, destroy)
}
func (p *Pattern) Status() Status { return patternStatus(p) }

type PatternType Enum

const (
	PATTERN_TYPE_SOLID PatternType = iota
	PATTERN_TYPE_SURFACE
	PATTERN_TYPE_LINEAR
	PATTERN_TYPE_RADIAL
)

type PdfVersion Enum

const (
	PDF_VERSION_1_4 PdfVersion = iota
	PDF_VERSION_1_5
)

var (
	PdfSurfaceCreate            func(filename string, widthInPoints, heightInPoints float64) *Surface
	PdfSurfaceCreateForStream   func(writeFunc WriteFunc, closure *T.Void, widthInPoints, heightInPoints float64) *Surface
	PdfSurfaceRestrictToVersion func(surface *Surface, version PdfVersion)
	PdfGetVersions              func(versions **PdfVersion, numVersions *int)
	PdfVersionToString          func(version PdfVersion) string
	PdfSurfaceSetSize           func(surface *Surface, widthInPoints, heightInPoints float64)
)

type PsLevel Enum

const (
	PS_LEVEL_2 PsLevel = iota
	PS_LEVEL_3
)

var (
	PsGetLevels     func(levels **PsLevel, numLevels *int)
	PsLevelToString func(level PsLevel)
)

var (
	PsSurfaceCreate          func(filename string, widthInPoints float64, heightInPoints float64) *Surface
	PsSurfaceCreateForStream func(writeFunc WriteFunc, closure *T.Void, widthInPoints float64, heightInPoints float64) *Surface

	psSurfaceDscBeginPageSetup func(s *Surface)
	psSurfaceDscBeginSetup     func(s *Surface)
	psSurfaceDscComment        func(s *Surface, comment string)
	psSurfaceGetEps            func(s *Surface)
	psSurfaceRestrictToLevel   func(s *Surface, level PsLevel)
	psSurfaceSetEps            func(s *Surface, eps Bool) string
	psSurfaceSetSize           func(s *Surface, widthInPoints, heightInPoints float64) Bool
)

func (s *Surface) DscBeginPageSetup()            { psSurfaceDscBeginPageSetup(s) }
func (s *Surface) DscBeginSetup()                { psSurfaceDscBeginSetup(s) }
func (s *Surface) DscComment(comment string)     { psSurfaceDscComment(s, comment) }
func (s *Surface) GetEps()                       { psSurfaceGetEps(s) }
func (s *Surface) RestrictToLevel(level PsLevel) { psSurfaceRestrictToLevel(s, level) }
func (s *Surface) SetEps(eps Bool) string        { return psSurfaceSetEps(s, eps) }
func (s *Surface) SetSize(widthInPoints, heightInPoints float64) Bool {
	return psSurfaceSetSize(s, widthInPoints, heightInPoints)
}
