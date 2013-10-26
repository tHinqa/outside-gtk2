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

var rectangleListDestroy func(r *RectangleList)

func (r *RectangleList) Destroy() { rectangleListDestroy(r) }

type Region struct{}

var (
	RegionCreate func() *Region

	RegionCreateRectangle  func(rectangle *RectangleInt) *Region
	RegionCreateRectangles func(rects *RectangleInt, count int) *Region

	regionContainsPoint      func(region *Region, x, y int) Bool
	regionContainsRectangle  func(region *Region, rectangle *RectangleInt) RegionOverlap
	regionCopy               func(original *Region) *Region
	regionDestroy            func(region *Region)
	regionEqual              func(a, b *Region) Bool
	regionGetExtents         func(region *Region, extents *RectangleInt)
	regionGetRectangle       func(region *Region, nth int, rectangle *RectangleInt)
	regionIntersect          func(dst, other *Region) Status
	regionIntersectRectangle func(dst *Region, rectangle *RectangleInt) Status
	regionIsEmpty            func(region *Region) Bool
	regionNumRectangles      func(region *Region) int
	regionReference          func(region *Region) *Region
	regionStatus             func(region *Region) Status
	regionSubtract           func(dst, other *Region) Status
	regionSubtractRectangle  func(dst *Region, rectangle *RectangleInt) Status
	regionTranslate          func(region *Region, dx, dy int)
	regionUnion              func(dst, other *Region) Status
	regionUnionRectangle     func(dst *Region, rectangle *RectangleInt) Status
	regionXor                func(dst, other *Region) Status
	regionXorRectangle       func(dst *Region, rectangle *RectangleInt) Status
)

func (r *Region) ContainsPoint(x, y int) Bool { return regionContainsPoint(r, x, y) }
func (r *Region) ContainsRectangle(rectangle *RectangleInt) RegionOverlap {
	return regionContainsRectangle(r, rectangle)
}
func (r *Region) Copy() *Region                                 { return regionCopy(r) }
func (r *Region) Destroy()                                      { regionDestroy(r) }
func (r *Region) Equal(r2 *Region) Bool                         { return regionEqual(r, r2) }
func (r *Region) GetExtents(extents *RectangleInt)              { regionGetExtents(r, extents) }
func (r *Region) GetRectangle(nth int, rectangle *RectangleInt) { regionGetRectangle(r, nth, rectangle) }
func (r *Region) Intersect(other *Region) Status                { return regionIntersect(r, other) }
func (r *Region) IntersectRectangle(rectangle *RectangleInt) Status {
	return regionIntersectRectangle(r, rectangle)
}
func (r *Region) IsEmpty() Bool                 { return regionIsEmpty(r) }
func (r *Region) NumRectangles() int            { return regionNumRectangles(r) }
func (r *Region) Reference() *Region            { return regionReference(r) }
func (r *Region) Status() Status                { return regionStatus(r) }
func (r *Region) Subtract(other *Region) Status { return regionSubtract(r, other) }
func (r *Region) SubtractRectangle(rectangle *RectangleInt) Status {
	return regionSubtractRectangle(r, rectangle)
}
func (r *Region) Translate(dx, dy int)       { regionTranslate(r, dx, dy) }
func (r *Region) Union(other *Region) Status { return regionUnion(r, other) }
func (r *Region) UnionRectangle(rectangle *RectangleInt) Status {
	return regionUnionRectangle(r, rectangle)
}
func (r *Region) Xor(other *Region) Status                    { return regionXor(r, other) }
func (r *Region) XorRectangle(rectangle *RectangleInt) Status { return regionXorRectangle(r, rectangle) }

type RegionOverlap Enum

const (
	REGION_OVERLAP_IN RegionOverlap = iota
	REGION_OVERLAP_OUT
	REGION_OVERLAP_PART
)
