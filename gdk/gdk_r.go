// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

var (
	RectangleGetType func() O.Type

	rectangleIntersect func(src1, src2, dest *Rectangle) T.Gboolean
	rectangleUnion     func(src1, src2, dest *Rectangle)
)

func (r *Rectangle) Intersect(r2, dest *Rectangle) T.Gboolean { return rectangleIntersect(r, r2, dest) }
func (r *Rectangle) Union(r2, dest *Rectangle)                { rectangleUnion(r, r2, dest) }

type Region struct{}

var (
	RegionNew func() *Region

	RegionPolygon   func(points *Point, nPoints int, fillRule T.GdkFillRule) *Region
	RegionRectangle func(rectangle *Rectangle) *Region

	regionCopy                  func(r *Region) *Region
	regionDestroy               func(r *Region)
	regionEmpty                 func(r *Region) T.Gboolean
	regionEqual                 func(r *Region, r2 *Region) T.Gboolean
	regionGetClipbox            func(r *Region, rectangle *Rectangle)
	regionGetRectangles         func(r *Region, rectangles **Rectangle, nRectangles *int)
	regionIntersect             func(r, r2 *Region)
	regionOffset                func(r *Region, dx, dy int)
	regionPointIn               func(r *Region, x, y int) T.Gboolean
	regionRectEqual             func(r *Region, rectangle *Rectangle) T.Gboolean
	regionRectIn                func(r *Region, rectangle *Rectangle) T.GdkOverlapType
	regionShrink                func(r *Region, dx, dy int)
	regionSpansIntersectForeach func(r *Region, spans *T.GdkSpan, nSpans int, sorted T.Gboolean, function T.GdkSpanFunc, data T.Gpointer)
	regionSubtract              func(r, r2 *Region)
	regionUnion                 func(r, r2 *Region)
	regionUnionWithRect         func(r *Region, rect *Rectangle)
	regionXor                   func(r, r2 *Region)
)

func (r *Region) Copy() *Region                   { return regionCopy(r) }
func (r *Region) Destroy()                        { regionDestroy(r) }
func (r *Region) Empty() T.Gboolean               { return regionEmpty(r) }
func (r *Region) Equal(r2 *Region) T.Gboolean     { return regionEqual(r, r2) }
func (r *Region) GetClipbox(rectangle *Rectangle) { regionGetClipbox(r, rectangle) }
func (r *Region) GetRectangles(rectangles **Rectangle, nRectangles *int) {
	regionGetRectangles(r, rectangles, nRectangles)
}
func (r *Region) Intersect(r2 *Region)                         { regionIntersect(r, r2) }
func (r *Region) Offset(dx, dy int)                            { regionOffset(r, dx, dy) }
func (r *Region) PointIn(x, y int) T.Gboolean                  { return regionPointIn(r, x, y) }
func (r *Region) RectEqual(rectangle *Rectangle) T.Gboolean    { return regionRectEqual(r, rectangle) }
func (r *Region) RectIn(rectangle *Rectangle) T.GdkOverlapType { return regionRectIn(r, rectangle) }
func (r *Region) Shrink(dx, dy int)                            { regionShrink(r, dx, dy) }
func (r *Region) SpansIntersectForeach(spans *T.GdkSpan, nSpans int, sorted T.Gboolean, function T.GdkSpanFunc, data T.Gpointer) {
	regionSpansIntersectForeach(r, spans, nSpans, sorted, function, data)
}
func (r *Region) Subtract(r2 *Region)           { regionSubtract(r, r2) }
func (r *Region) Union(r2 *Region)              { regionUnion(r, r2) }
func (r *Region) UnionWithRect(rect *Rectangle) { regionUnionWithRect(r, rect) }
func (r *Region) Xor(r2 *Region)                { regionXor(r, r2) }
