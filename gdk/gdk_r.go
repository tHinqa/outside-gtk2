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

	RectangleIntersect func(src1, src2, dest *Rectangle) bool
	RectangleUnion     func(src1, src2, dest *Rectangle)
)

func (r *Rectangle) Intersect(r2, dest *Rectangle) bool { return RectangleIntersect(r, r2, dest) }
func (r *Rectangle) Union(r2, dest *Rectangle)          { RectangleUnion(r, r2, dest) }

type Region struct{}

var (
	RegionNew func() *Region

	RegionPolygon   func(points *Point, nPoints int, fillRule T.GdkFillRule) *Region
	RegionRectangle func(rectangle *Rectangle) *Region

	RegionCopy                  func(r *Region) *Region
	RegionDestroy               func(r *Region)
	RegionEmpty                 func(r *Region) bool
	RegionEqual                 func(r *Region, r2 *Region) bool
	RegionGetClipbox            func(r *Region, rectangle *Rectangle)
	RegionGetRectangles         func(r *Region, rectangles **Rectangle, nRectangles *int)
	RegionIntersect             func(r, r2 *Region)
	RegionOffset                func(r *Region, dx, dy int)
	RegionPointIn               func(r *Region, x, y int) bool
	RegionRectEqual             func(r *Region, rectangle *Rectangle) bool
	RegionRectIn                func(r *Region, rectangle *Rectangle) T.GdkOverlapType
	RegionShrink                func(r *Region, dx, dy int)
	RegionSpansIntersectForeach func(r *Region, spans *T.GdkSpan, nSpans int, sorted bool, function T.GdkSpanFunc, data T.Gpointer)
	RegionSubtract              func(r, r2 *Region)
	RegionUnion                 func(r, r2 *Region)
	RegionUnionWithRect         func(r *Region, rect *Rectangle)
	RegionXor                   func(r, r2 *Region)
)

func (r *Region) Copy() *Region                   { return RegionCopy(r) }
func (r *Region) Destroy()                        { RegionDestroy(r) }
func (r *Region) Empty() bool                     { return RegionEmpty(r) }
func (r *Region) Equal(r2 *Region) bool           { return RegionEqual(r, r2) }
func (r *Region) GetClipbox(rectangle *Rectangle) { RegionGetClipbox(r, rectangle) }
func (r *Region) GetRectangles(rectangles **Rectangle, nRectangles *int) {
	RegionGetRectangles(r, rectangles, nRectangles)
}
func (r *Region) Intersect(r2 *Region)                         { RegionIntersect(r, r2) }
func (r *Region) Offset(dx, dy int)                            { RegionOffset(r, dx, dy) }
func (r *Region) PointIn(x, y int) bool                        { return RegionPointIn(r, x, y) }
func (r *Region) RectEqual(rectangle *Rectangle) bool          { return RegionRectEqual(r, rectangle) }
func (r *Region) RectIn(rectangle *Rectangle) T.GdkOverlapType { return RegionRectIn(r, rectangle) }
func (r *Region) Shrink(dx, dy int)                            { RegionShrink(r, dx, dy) }
func (r *Region) SpansIntersectForeach(spans *T.GdkSpan, nSpans int, sorted bool, function T.GdkSpanFunc, data T.Gpointer) {
	RegionSpansIntersectForeach(r, spans, nSpans, sorted, function, data)
}
func (r *Region) Subtract(r2 *Region)           { RegionSubtract(r, r2) }
func (r *Region) Union(r2 *Region)              { RegionUnion(r, r2) }
func (r *Region) UnionWithRect(rect *Rectangle) { RegionUnionWithRect(r, rect) }
func (r *Region) Xor(r2 *Region)                { RegionXor(r, r2) }
