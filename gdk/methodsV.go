// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"

)

type VisibilityState Enum

const (
	VISIBILITY_UNOBSCURED VisibilityState = iota
	VISIBILITY_PARTIAL
	VISIBILITY_FULLY_OBSCURED
)

var VisibilityStateGetType func() T.GType

type Visual struct {
	Parent       T.GObject
	Type         VisualType
	Depth        int
	ByteOrder    T.GdkByteOrder
	ColormapSize int
	BitsPerRgb   int
	RedMask      T.GUint32
	RedShift     int
	RedPrec      int
	GreenMask    T.GUint32
	GreenShift   int
	GreenPrec    int
	BlueMask     T.GUint32
	BlueShift    int
	BluePrec     int
}

var (
	VisualGetType func() T.GType

	VisualGetBestDepth     func() int
	VisualGetBestType      func() VisualType
	VisualGetSystem        func() *Visual
	VisualGetBest          func() *Visual
	VisualGetBestWithDepth func(depth int) *Visual

	visualGetBitsPerRgb        func(v *Visual) int
	visualGetBluePixelDetails  func(v *Visual, mask *T.GUint32, shift *int, precision *int)
	visualGetByteOrder         func(v *Visual) T.GdkByteOrder
	visualGetColormapSize      func(v *Visual) int
	visualGetDepth             func(v *Visual) int
	visualGetGreenPixelDetails func(v *Visual, mask *T.GUint32, shift *int, precision *int)
	visualGetRedPixelDetails   func(v *Visual, mask *T.GUint32, shift *int, precision *int)
	visualGetScreen            func(v *Visual) *Screen
	visualGetVisualType        func(v *Visual) VisualType
)

func (v *Visual) GetBitsPerRgb() int { return visualGetBitsPerRgb(v) }
func (v *Visual) GetBluePixelDetails(mask *T.GUint32, shift *int, precision *int) {
	visualGetBluePixelDetails(v, mask, shift, precision)
}
func (v *Visual) GetByteOrder() T.GdkByteOrder { return visualGetByteOrder(v) }
func (v *Visual) GetColormapSize() int         { return visualGetColormapSize(v) }
func (v *Visual) GetDepth() int                { return visualGetDepth(v) }
func (v *Visual) GetGreenPixelDetails(mask *T.GUint32, shift *int, precision *int) {
	visualGetGreenPixelDetails(v, mask, shift, precision)
}
func (v *Visual) GetRedPixelDetails(mask *T.GUint32, shift *int, precision *int) {
	visualGetRedPixelDetails(v, mask, shift, precision)
}
func (v *Visual) GetScreen() *Screen        { return visualGetScreen(v) }
func (v *Visual) GetVisualType() VisualType { return visualGetVisualType(v) }

type VisualType T.Enum

const (
	VISUAL_STATIC_GRAY VisualType = iota
	VISUAL_GRAYSCALE
	VISUAL_STATIC_COLOR
	VISUAL_PSEUDO_COLOR
	VISUAL_TRUE_COLOR
	VISUAL_DIRECT_COLOR
)

var (
	VisualTypeGetType func() T.GType

	VisualGetBestWithBoth func(depth int, visualType VisualType) *Visual
	VisualGetBestWithType func(visualType VisualType) *Visual
)
