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

	PatternAddColorStopRgb   func(p *Pattern, offset, red, green, blue float64)
	PatternAddColorStopRgba  func(p *Pattern, offset, red, green, blue, alpha float64)
	PatternDestroy           func(p *Pattern)
	PatternGetColorStopCount func(p *Pattern, count *int) Status
	PatternGetColorStopRgba  func(p *Pattern, index int, offset, red, green, blue, alpha *float64) Status
	PatternGetExtend         func(p *Pattern) Extend
	PatternGetFilter         func(p *Pattern) Filter
	PatternGetLinearPoints   func(p *Pattern, x0, y0, x1, y1 *float64) Status
	PatternGetMatrix         func(p *Pattern, matrix *Matrix)
	PatternGetRadialCircles  func(p *Pattern, x0, y0, r0, x1, y1, r1 *float64) Status
	PatternGetReferenceCount func(p *Pattern) uint
	PatternGetRgba           func(p *Pattern, red, green, blue, alpha *float64) Status
	PatternGetSurface        func(p *Pattern, surface **Surface) Status
	PatternGetType           func(p *Pattern) PatternType
	PatternGetUserData       func(p *Pattern, key *UserDataKey) *T.Void
	PatternReference         func(p *Pattern) *Pattern
	PatternSetExtend         func(p *Pattern, extend Extend)
	PatternSetFilter         func(p *Pattern, filter Filter)
	PatternSetMatrix         func(p *Pattern, matrix *Matrix)
	PatternSetUserData       func(p *Pattern, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	PatternStatus            func(p *Pattern) Status
)

func (p *Pattern) AddColorStopRgb(offset, red, green, blue float64) {
	PatternAddColorStopRgb(p, offset, red, green, blue)
}
func (p *Pattern) AddColorStopRgba(offset, red, green, blue, alpha float64) {
	PatternAddColorStopRgba(p, offset, red, green, blue, alpha)
}
func (p *Pattern) Destroy() { PatternDestroy(p) }
func (p *Pattern) GetColorStopCount(count *int) Status {
	return PatternGetColorStopCount(p, count)
}
func (p *Pattern) GetColorStopRgba(index int, offset, red, green, blue, alpha *float64) Status {
	return PatternGetColorStopRgba(p, index, offset, red, green, blue, alpha)
}
func (p *Pattern) GetExtend() Extend { return PatternGetExtend(p) }
func (p *Pattern) GetFilter() Filter { return PatternGetFilter(p) }
func (p *Pattern) GetLinearPoints(x0, y0, x1, y1 *float64) Status {
	return PatternGetLinearPoints(p, x0, y0, x1, y1)
}
func (p *Pattern) GetMatrix(matrix *Matrix) { PatternGetMatrix(p, matrix) }
func (p *Pattern) GetRadialCircles(x0, y0, r0, x1, y1, r1 *float64) Status {
	return PatternGetRadialCircles(p, x0, y0, r0, x1, y1, r1)
}
func (p *Pattern) GetReferenceCount() uint { return PatternGetReferenceCount(p) }
func (p *Pattern) GetRgba(red, green, blue, alpha *float64) Status {
	return PatternGetRgba(p, red, green, blue, alpha)
}
func (p *Pattern) GetSurface(surface **Surface) Status  { return PatternGetSurface(p, surface) }
func (p *Pattern) GetType() PatternType                 { return PatternGetType(p) }
func (p *Pattern) GetUserData(key *UserDataKey) *T.Void { return PatternGetUserData(p, key) }
func (p *Pattern) Reference() *Pattern                  { return PatternReference(p) }
func (p *Pattern) SetExtend(extend Extend)              { PatternSetExtend(p, extend) }
func (p *Pattern) SetFilter(filter Filter)              { PatternSetFilter(p, filter) }
func (p *Pattern) SetMatrix(matrix *Matrix)             { PatternSetMatrix(p, matrix) }
func (p *Pattern) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return PatternSetUserData(p, key, userData, destroy)
}
func (p *Pattern) Status() Status { return PatternStatus(p) }

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

	PsSurfaceDscBeginPageSetup func(s *Surface)
	PsSurfaceDscBeginSetup     func(s *Surface)
	PsSurfaceDscComment        func(s *Surface, comment string)
	PsSurfaceGetEps            func(s *Surface)
	PsSurfaceRestrictToLevel   func(s *Surface, level PsLevel)
	PsSurfaceSetEps            func(s *Surface, eps Bool) string
	PsSurfaceSetSize           func(s *Surface, widthInPoints, heightInPoints float64) Bool
)

func (s *Surface) DscBeginPageSetup()            { PsSurfaceDscBeginPageSetup(s) }
func (s *Surface) DscBeginSetup()                { PsSurfaceDscBeginSetup(s) }
func (s *Surface) DscComment(comment string)     { PsSurfaceDscComment(s, comment) }
func (s *Surface) GetEps()                       { PsSurfaceGetEps(s) }
func (s *Surface) RestrictToLevel(level PsLevel) { PsSurfaceRestrictToLevel(s, level) }
func (s *Surface) SetEps(eps Bool) string        { return PsSurfaceSetEps(s, eps) }
func (s *Surface) SetSize(widthInPoints, heightInPoints float64) Bool {
	return PsSurfaceSetSize(s, widthInPoints, heightInPoints)
}
