// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

type Rectangle struct {
	X, Y, Width, Height float64
}

type RectangleInt struct {
	X, Y          int
	Width, Height int
}

type RectangleList struct {
	Status        Status
	Rectangles    *Rectangle
	NumRectangles int
}

var RectangleListDestroy func(r *RectangleList)

func (r *RectangleList) Destroy() { RectangleListDestroy(r) }

type Region struct{}

var (
	RegionCreate func() *Region

	RegionCreateRectangle  func(rectangle *RectangleInt) *Region
	RegionCreateRectangles func(rects *RectangleInt, count int) *Region

	RegionContainsPoint      func(region *Region, x, y int) Bool
	RegionContainsRectangle  func(region *Region, rectangle *RectangleInt) RegionOverlap
	RegionCopy               func(original *Region) *Region
	RegionDestroy            func(region *Region)
	RegionEqual              func(a, b *Region) Bool
	RegionGetExtents         func(region *Region, extents *RectangleInt)
	RegionGetRectangle       func(region *Region, nth int, rectangle *RectangleInt)
	RegionIntersect          func(dst, other *Region) Status
	RegionIntersectRectangle func(dst *Region, rectangle *RectangleInt) Status
	RegionIsEmpty            func(region *Region) Bool
	RegionNumRectangles      func(region *Region) int
	RegionReference          func(region *Region) *Region
	RegionStatus             func(region *Region) Status
	RegionSubtract           func(dst, other *Region) Status
	RegionSubtractRectangle  func(dst *Region, rectangle *RectangleInt) Status
	RegionTranslate          func(region *Region, dx, dy int)
	RegionUnion              func(dst, other *Region) Status
	RegionUnionRectangle     func(dst *Region, rectangle *RectangleInt) Status
	RegionXor                func(dst, other *Region) Status
	RegionXorRectangle       func(dst *Region, rectangle *RectangleInt) Status
)

func (r *Region) ContainsPoint(x, y int) Bool { return RegionContainsPoint(r, x, y) }
func (r *Region) ContainsRectangle(rectangle *RectangleInt) RegionOverlap {
	return RegionContainsRectangle(r, rectangle)
}
func (r *Region) Copy() *Region                                 { return RegionCopy(r) }
func (r *Region) Destroy()                                      { RegionDestroy(r) }
func (r *Region) Equal(r2 *Region) Bool                         { return RegionEqual(r, r2) }
func (r *Region) GetExtents(extents *RectangleInt)              { RegionGetExtents(r, extents) }
func (r *Region) GetRectangle(nth int, rectangle *RectangleInt) { RegionGetRectangle(r, nth, rectangle) }
func (r *Region) Intersect(other *Region) Status                { return RegionIntersect(r, other) }
func (r *Region) IntersectRectangle(rectangle *RectangleInt) Status {
	return RegionIntersectRectangle(r, rectangle)
}
func (r *Region) IsEmpty() Bool                 { return RegionIsEmpty(r) }
func (r *Region) NumRectangles() int            { return RegionNumRectangles(r) }
func (r *Region) Reference() *Region            { return RegionReference(r) }
func (r *Region) Status() Status                { return RegionStatus(r) }
func (r *Region) Subtract(other *Region) Status { return RegionSubtract(r, other) }
func (r *Region) SubtractRectangle(rectangle *RectangleInt) Status {
	return RegionSubtractRectangle(r, rectangle)
}
func (r *Region) Translate(dx, dy int)       { RegionTranslate(r, dx, dy) }
func (r *Region) Union(other *Region) Status { return RegionUnion(r, other) }
func (r *Region) UnionRectangle(rectangle *RectangleInt) Status {
	return RegionUnionRectangle(r, rectangle)
}
func (r *Region) Xor(other *Region) Status                    { return RegionXor(r, other) }
func (r *Region) XorRectangle(rectangle *RectangleInt) Status { return RegionXorRectangle(r, rectangle) }

type RegionOverlap Enum

const (
	REGION_OVERLAP_IN RegionOverlap = iota
	REGION_OVERLAP_OUT
	REGION_OVERLAP_PART
)
