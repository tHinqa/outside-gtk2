// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type GC struct {
	Parent_instance O.Object
	ClipXOrigin     int
	ClipYOrigin     int
	TsXOrigin       int
	TsYOrigin       int
	Colormap        *Colormap
}

var (
	GcGetType       func() O.Type
	GcNew           func(drawable *Drawable) *GC
	GcNewWithValues func(drawable *Drawable, values *GCValues, valuesMask GCValuesMask) *GC

	GcRef               func(gc *GC) *GC
	GcUnref             func(gc *GC)
	GcGetValues         func(gc *GC, values *GCValues)
	GcSetValues         func(gc *GC, values *GCValues, valuesMask GCValuesMask)
	GcSetForeground     func(gc *GC, color *Color)
	GcSetBackground     func(gc *GC, color *Color)
	GcSetFont           func(gc *GC, font *Font)
	GcSetFunction       func(gc *GC, function T.GdkFunction)
	GcSetFill           func(gc *GC, fill T.GdkFill)
	GcSetTile           func(gc *GC, tile *Pixmap)
	GcSetStipple        func(gc *GC, stipple *Pixmap)
	GcSetTsOrigin       func(gc *GC, x, y int)
	GcSetClipOrigin     func(gc *GC, x, y int)
	GcSetClipMask       func(gc *GC, mask *T.GdkBitmap)
	GcSetClipRectangle  func(gc *GC, rectangle *Rectangle)
	GcSetClipRegion     func(gc *GC, region *Region)
	GcSetSubwindow      func(gc *GC, mode T.GdkSubwindowMode)
	GcSetExposures      func(gc *GC, exposures bool)
	GcSetLineAttributes func(gc *GC, lineWidth int, lineStyle T.GdkLineStyle, capStyle CapStyle, joinStyle T.GdkJoinStyle)
	GcSetDashes         func(gc *GC, dashOffset int, dashList *int8, n int)
	GcOffset            func(gc *GC, xOffset, yOffset int)
	GcCopy              func(dstGc, srcGc *GC)
	GcSetColormap       func(gc *GC, colormap *Colormap)
	GcGetColormap       func(gc *GC) *Colormap
	GcSetRgbFgColor     func(gc *GC, color *Color)
	GcSetRgbBgColor     func(gc *GC, color *Color)
	GcGetScreen         func(gc *GC) *Screen
)

func (g *GC) Ref() *GC                                            { return GcRef(g) }
func (g *GC) Unref()                                              { GcUnref(g) }
func (g *GC) GetValues(values *GCValues)                          { GcGetValues(g, values) }
func (g *GC) SetValues(values *GCValues, valuesMask GCValuesMask) { GcSetValues(g, values, valuesMask) }
func (g *GC) SetForeground(color *Color)                          { GcSetForeground(g, color) }
func (g *GC) SetBackground(color *Color)                          { GcSetBackground(g, color) }
func (g *GC) SetFont(font *Font)                                  { GcSetFont(g, font) }
func (g *GC) SetFunction(function T.GdkFunction)                  { GcSetFunction(g, function) }
func (g *GC) SetFill(fill T.GdkFill)                              { GcSetFill(g, fill) }
func (g *GC) SetTile(tile *Pixmap)                                { GcSetTile(g, tile) }
func (g *GC) SetStipple(stipple *Pixmap)                          { GcSetStipple(g, stipple) }
func (g *GC) SetTsOrigin(x, y int)                                { GcSetTsOrigin(g, x, y) }
func (g *GC) SetClipOrigin(x, y int)                              { GcSetClipOrigin(g, x, y) }
func (g *GC) SetClipMask(mask *T.GdkBitmap)                       { GcSetClipMask(g, mask) }
func (g *GC) SetClipRectangle(rectangle *Rectangle)               { GcSetClipRectangle(g, rectangle) }
func (g *GC) SetClipRegion(region *Region)                        { GcSetClipRegion(g, region) }
func (g *GC) SetSubwindow(mode T.GdkSubwindowMode)                { GcSetSubwindow(g, mode) }
func (g *GC) SetExposures(exposures bool)                         { GcSetExposures(g, exposures) }
func (g *GC) SetLineAttributes(lineWidth int, lineStyle T.GdkLineStyle, capStyle CapStyle, joinStyle T.GdkJoinStyle) {
	GcSetLineAttributes(g, lineWidth, lineStyle, capStyle, joinStyle)
}
func (g *GC) SetDashes(dashOffset int, dashList *int8, n int) {
	GcSetDashes(g, dashOffset, dashList, n)
}
func (g *GC) Offset(xOffset, yOffset int)    { GcOffset(g, xOffset, yOffset) }
func (g *GC) Copy(srcGc *GC)                 { GcCopy(g, srcGc) }
func (g *GC) SetColormap(colormap *Colormap) { GcSetColormap(g, colormap) }
func (g *GC) GetColormap() *Colormap         { return GcGetColormap(g) }
func (g *GC) SetRgbFgColor(color *Color)     { GcSetRgbFgColor(g, color) }
func (g *GC) SetRgbBgColor(color *Color)     { GcSetRgbBgColor(g, color) }
func (g *GC) GetScreen() *Screen             { return GcGetScreen(g) }

type GCValues struct {
	Foreground        Color
	Background        Color
	Font              *Font
	Function          T.GdkFunction
	Fill              T.GdkFill
	Tile              *Pixmap
	Stipple           *Pixmap
	ClipMask          *Pixmap
	SubwindowMode     T.GdkSubwindowMode
	TsXOrigin         int
	TsYOrigin         int
	ClipXOrigin       int
	ClipYOrigin       int
	GraphicsExposures int
	LineWidth         int
	LineStyle         T.GdkLineStyle
	CapStyle          CapStyle
	JoinStyle         T.GdkJoinStyle
}

type GCValuesMask Enum

const (
	GC_FOREGROUND GCValuesMask = 1 << iota
	GC_BACKGROUND
	GC_FONT
	GC_FUNCTION
	GC_FILL
	GC_TILE
	GC_STIPPLE
	GC_CLIP_MASK
	GC_SUBWINDOW
	GC_TS_X_ORIGIN
	GC_TS_Y_ORIGIN
	GC_CLIP_X_ORIGIN
	GC_CLIP_Y_ORIGIN
	GC_EXPOSURES
	GC_LINE_WIDTH
	GC_LINE_STYLE
	GC_CAP_STYLE
	GC_JOIN_STYLE
)

var GcValuesMaskGetType func() O.Type

type Geometry struct {
	MinWidth   int
	MinHeight  int
	MaxWidth   int
	MaxHeight  int
	BaseWidth  int
	BaseHeight int
	WidthInc   int
	HeightInc  int
	MinAspect  float64
	MaxAspect  float64
	WinGravity Gravity
}

type GrabStatus Enum

const (
	GRAB_SUCCESS GrabStatus = iota
	GRAB_ALREADY_GRABBED
	GRAB_INVALID_TIME
	GRAB_NOT_VIEWABLE
	GRAB_FROZEN
)

var GrabStatusGetType func() O.Type

type Gravity Enum

const (
	GRAVITY_NORTH_WEST Gravity = iota + 1
	GRAVITY_NORTH
	GRAVITY_NORTH_EAST
	GRAVITY_WEST
	GRAVITY_CENTER
	GRAVITY_EAST
	GRAVITY_SOUTH_WEST
	GRAVITY_SOUTH
	GRAVITY_SOUTH_EAST
	GRAVITY_STATIC
)

var GravityGetType func() O.Type
