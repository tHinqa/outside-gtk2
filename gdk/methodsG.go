// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type GC struct {
	Parent_instance T.GObject
	ClipXOrigin     int
	ClipYOrigin     int
	TsXOrigin       int
	TsYOrigin       int
	Colormap        *Colormap
}

var (
	GcGetType       func() T.GType
	GcNew           func(drawable *Drawable) *GC
	GcNewWithValues func(drawable *Drawable, values *GCValues, valuesMask GCValuesMask) *GC

	gcRef               func(gc *GC) *GC
	gcUnref             func(gc *GC)
	gcGetValues         func(gc *GC, values *GCValues)
	gcSetValues         func(gc *GC, values *GCValues, valuesMask GCValuesMask)
	gcSetForeground     func(gc *GC, color *Color)
	gcSetBackground     func(gc *GC, color *Color)
	gcSetFont           func(gc *GC, font *Font)
	gcSetFunction       func(gc *GC, function T.GdkFunction)
	gcSetFill           func(gc *GC, fill T.GdkFill)
	gcSetTile           func(gc *GC, tile *Pixmap)
	gcSetStipple        func(gc *GC, stipple *Pixmap)
	gcSetTsOrigin       func(gc *GC, x, y int)
	gcSetClipOrigin     func(gc *GC, x, y int)
	gcSetClipMask       func(gc *GC, mask *T.GdkBitmap)
	gcSetClipRectangle  func(gc *GC, rectangle *Rectangle)
	gcSetClipRegion     func(gc *GC, region *Region)
	gcSetSubwindow      func(gc *GC, mode T.GdkSubwindowMode)
	gcSetExposures      func(gc *GC, exposures T.Gboolean)
	gcSetLineAttributes func(gc *GC, lineWidth int, lineStyle T.GdkLineStyle, capStyle CapStyle, joinStyle T.GdkJoinStyle)
	gcSetDashes         func(gc *GC, dashOffset int, dashList *int8, n int)
	gcOffset            func(gc *GC, xOffset, yOffset int)
	gcCopy              func(dstGc, srcGc *GC)
	gcSetColormap       func(gc *GC, colormap *Colormap)
	gcGetColormap       func(gc *GC) *Colormap
	gcSetRgbFgColor     func(gc *GC, color *Color)
	gcSetRgbBgColor     func(gc *GC, color *Color)
	gcGetScreen         func(gc *GC) *Screen
)

func (g *GC) Ref() *GC                                            { return gcRef(g) }
func (g *GC) Unref()                                              { gcUnref(g) }
func (g *GC) GetValues(values *GCValues)                          { gcGetValues(g, values) }
func (g *GC) SetValues(values *GCValues, valuesMask GCValuesMask) { gcSetValues(g, values, valuesMask) }
func (g *GC) SetForeground(color *Color)                          { gcSetForeground(g, color) }
func (g *GC) SetBackground(color *Color)                          { gcSetBackground(g, color) }
func (g *GC) SetFont(font *Font)                                  { gcSetFont(g, font) }
func (g *GC) SetFunction(function T.GdkFunction)                  { gcSetFunction(g, function) }
func (g *GC) SetFill(fill T.GdkFill)                              { gcSetFill(g, fill) }
func (g *GC) SetTile(tile *Pixmap)                                { gcSetTile(g, tile) }
func (g *GC) SetStipple(stipple *Pixmap)                          { gcSetStipple(g, stipple) }
func (g *GC) SetTsOrigin(x, y int)                                { gcSetTsOrigin(g, x, y) }
func (g *GC) SetClipOrigin(x, y int)                              { gcSetClipOrigin(g, x, y) }
func (g *GC) SetClipMask(mask *T.GdkBitmap)                       { gcSetClipMask(g, mask) }
func (g *GC) SetClipRectangle(rectangle *Rectangle)               { gcSetClipRectangle(g, rectangle) }
func (g *GC) SetClipRegion(region *Region)                        { gcSetClipRegion(g, region) }
func (g *GC) SetSubwindow(mode T.GdkSubwindowMode)                { gcSetSubwindow(g, mode) }
func (g *GC) SetExposures(exposures T.Gboolean)                   { gcSetExposures(g, exposures) }
func (g *GC) SetLineAttributes(lineWidth int, lineStyle T.GdkLineStyle, capStyle CapStyle, joinStyle T.GdkJoinStyle) {
	gcSetLineAttributes(g, lineWidth, lineStyle, capStyle, joinStyle)
}
func (g *GC) SetDashes(dashOffset int, dashList *int8, n int) {
	gcSetDashes(g, dashOffset, dashList, n)
}
func (g *GC) Offset(xOffset, yOffset int)    { gcOffset(g, xOffset, yOffset) }
func (g *GC) Copy(srcGc *GC)                 { gcCopy(g, srcGc) }
func (g *GC) SetColormap(colormap *Colormap) { gcSetColormap(g, colormap) }
func (g *GC) GetColormap() *Colormap         { return gcGetColormap(g) }
func (g *GC) SetRgbFgColor(color *Color)     { gcSetRgbFgColor(g, color) }
func (g *GC) SetRgbBgColor(color *Color)     { gcSetRgbBgColor(g, color) }
func (g *GC) GetScreen() *Screen             { return gcGetScreen(g) }

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

var GcValuesMaskGetType func() T.GType

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

var GrabStatusGetType func() T.GType

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

var GravityGetType func() T.GType
