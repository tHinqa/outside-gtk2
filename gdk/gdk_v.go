// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"

)

type VisibilityState Enum

const (
	VISIBILITY_UNOBSCURED VisibilityState = iota
	VISIBILITY_PARTIAL
	VISIBILITY_FULLY_OBSCURED
)

var VisibilityStateGetType func() O.Type

type Visual struct {
	Parent       O.Object
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
	VisualGetType func() O.Type

	VisualGetBestDepth     func() int
	VisualGetBestType      func() VisualType
	VisualGetSystem        func() *Visual
	VisualGetBest          func() *Visual
	VisualGetBestWithDepth func(depth int) *Visual

	VisualGetBitsPerRgb        func(v *Visual) int
	VisualGetBluePixelDetails  func(v *Visual, mask *T.GUint32, shift *int, precision *int)
	VisualGetByteOrder         func(v *Visual) T.GdkByteOrder
	VisualGetColormapSize      func(v *Visual) int
	VisualGetDepth             func(v *Visual) int
	VisualGetGreenPixelDetails func(v *Visual, mask *T.GUint32, shift *int, precision *int)
	VisualGetRedPixelDetails   func(v *Visual, mask *T.GUint32, shift *int, precision *int)
	VisualGetScreen            func(v *Visual) *Screen
	VisualGetVisualType        func(v *Visual) VisualType
)

func (v *Visual) GetBitsPerRgb() int { return VisualGetBitsPerRgb(v) }
func (v *Visual) GetBluePixelDetails(mask *T.GUint32, shift *int, precision *int) {
	VisualGetBluePixelDetails(v, mask, shift, precision)
}
func (v *Visual) GetByteOrder() T.GdkByteOrder { return VisualGetByteOrder(v) }
func (v *Visual) GetColormapSize() int         { return VisualGetColormapSize(v) }
func (v *Visual) GetDepth() int                { return VisualGetDepth(v) }
func (v *Visual) GetGreenPixelDetails(mask *T.GUint32, shift *int, precision *int) {
	VisualGetGreenPixelDetails(v, mask, shift, precision)
}
func (v *Visual) GetRedPixelDetails(mask *T.GUint32, shift *int, precision *int) {
	VisualGetRedPixelDetails(v, mask, shift, precision)
}
func (v *Visual) GetScreen() *Screen        { return VisualGetScreen(v) }
func (v *Visual) GetVisualType() VisualType { return VisualGetVisualType(v) }

type VisualType Enum

const (
	VISUAL_STATIC_GRAY VisualType = iota
	VISUAL_GRAYSCALE
	VISUAL_STATIC_COLOR
	VISUAL_PSEUDO_COLOR
	VISUAL_TRUE_COLOR
	VISUAL_DIRECT_COLOR
)

var (
	VisualTypeGetType func() O.Type

	VisualGetBestWithBoth func(depth int, visualType VisualType) *Visual
	VisualGetBestWithType func(visualType VisualType) *Visual
)
