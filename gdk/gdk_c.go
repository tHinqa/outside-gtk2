// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Cairo struct{} //TODO(t):Def and push methods to /cairo? pragmatic pros/cons

var (
	CairoCreate func(drawable *Drawable) *Cairo

	CairoResetClip       func(c *Cairo, drawable *Drawable)
	CairoSetSourceColor  func(c *Cairo, color *Color)
	CairoSetSourcePixbuf func(c *Cairo, pixbuf *Pixbuf, pixbufX, pixbufY float64)
	CairoSetSourcePixmap func(c *Cairo, pixmap *Pixmap, pixmapX, pixmapY float64)
	CairoSetSourceWindow func(c *Cairo, window *Window, x, y float64)
	CairoRectangle       func(c *Cairo, rectangle *Rectangle)
	CairoRegion          func(c *Cairo, region *Region)
)

func (c *Cairo) ResetClip(drawable *Drawable) { CairoResetClip(c, drawable) }
func (c *Cairo) SetSourceColor(color *Color)  { CairoSetSourceColor(c, color) }
func (c *Cairo) SetSourcePixbuf(pixbuf *Pixbuf, pixbufX, pixbufY float64) {
	CairoSetSourcePixbuf(c, pixbuf, pixbufX, pixbufY)
}
func (c *Cairo) SetSourcePixmap(pixmap *Pixmap, pixmapX, pixmapY float64) {
	CairoSetSourcePixmap(c, pixmap, pixmapX, pixmapY)
}
func (c *Cairo) SetSourceWindow(window *Window, x, y float64) { CairoSetSourceWindow(c, window, x, y) }
func (c *Cairo) Rectangle(rectangle *Rectangle)               { CairoRectangle(c, rectangle) }
func (c *Cairo) Region(region *Region)                        { CairoRegion(c, region) }

type CapStyle Enum

const (
	CAP_NOT_LAST CapStyle = iota
	CAP_BUTT
	CAP_ROUND
	CAP_PROJECTING
)

var CapStyleGetType func() O.Type

type Color struct {
	Pixel T.GUint32
	Red   uint16
	Green uint16
	Blue  uint16
}

var (
	ColorGetType func() O.Type

	ColorParse func(spec string, color *Color) bool

	ColorAlloc  func(c *Colormap, color *Color) int
	ColorBlack  func(c *Colormap, color *Color) int
	ColorChange func(c *Colormap, color *Color) int
	ColorsAlloc func(c *Colormap, contiguous bool, planes *T.Gulong, nplanes int, pixels *T.Gulong, npixels int) int
	ColorsFree  func(c *Colormap, pixels *T.Gulong, npixels int, planes T.Gulong)
	ColorWhite  func(c *Colormap, color *Color) int

	ColorCopy     func(c *Color) *Color
	ColorEqual    func(c *Color, colorb *Color) bool
	ColorFree     func(c *Color)
	ColorHash     func(c *Color) uint
	ColorToString func(c *Color) string
)

func (c *Color) Copy() *Color             { return ColorCopy(c) }
func (c *Color) Equal(colorb *Color) bool { return ColorEqual(c, colorb) }
func (c *Color) Free()                    { ColorFree(c) }
func (c *Color) Hash() uint               { return ColorHash(c) }
func (c *Color) ToString() string         { return ColorToString(c) }

type Colormap struct {
	Parent        O.Object
	Size          int
	Colors        *Color
	Visual        *Visual
	WindowingData T.Gpointer
}

var (
	ColormapGetType func() O.Type
	ColormapNew     func(visual *Visual, allocate bool) *Colormap

	ColormapGetSystem     func() *Colormap
	ColormapGetSystemSize func() int

	ColorsStore func(c *Colormap, colors *Color, ncolors int)

	ColormapAllocColor  func(c *Colormap, color *Color, writeable bool, bestMatch bool) bool
	ColormapAllocColors func(c *Colormap, colors *Color, nColors int, writeable bool, bestMatch bool, success *bool) int
	ColormapChange      func(c *Colormap, ncolors int)
	ColormapFreeColors  func(c *Colormap, colors *Color, nColors int)
	ColormapGetScreen   func(c *Colormap) *Screen
	ColormapGetVisual   func(c *Colormap) *Visual
	ColormapQueryColor  func(c *Colormap, pixel T.Gulong, result *Color)
	ColormapRef         func(c *Colormap) *Colormap
	ColormapUnref       func(c *Colormap)
)

func (c *Colormap) AllocColor(color *Color, writeable bool, bestMatch bool) bool {
	return ColormapAllocColor(c, color, writeable, bestMatch)
}
func (c *Colormap) AllocColors(colors *Color, nColors int, writeable bool, bestMatch bool, success *bool) int {
	return ColormapAllocColors(c, colors, nColors, writeable, bestMatch, success)
}
func (c *Colormap) Change(ncolors int)                       { ColormapChange(c, ncolors) }
func (c *Colormap) FreeColors(colors *Color, nColors int)    { ColormapFreeColors(c, colors, nColors) }
func (c *Colormap) GetScreen() *Screen                       { return ColormapGetScreen(c) }
func (c *Colormap) GetVisual() *Visual                       { return ColormapGetVisual(c) }
func (c *Colormap) QueryColor(pixel T.Gulong, result *Color) { ColormapQueryColor(c, pixel, result) }
func (c *Colormap) Ref() *Colormap                           { return ColormapRef(c) }
func (c *Colormap) Unref()                                   { ColormapUnref(c) }

type Colorspace Enum

const COLORSPACE_RGB Colorspace = 0

var ColorspaceGetType func() O.Type

type CrossingMode Enum

const (
	CROSSING_NORMAL CrossingMode = iota
	CROSSING_GRAB
	CROSSING_UNGRAB
	CROSSING_GTK_GRAB
	CROSSING_GTK_UNGRAB
	CROSSING_STATE_CHANGED
)

var CrossingModeGetType func() O.Type

type Cursor struct {
	Type     CursorType
	RefCount uint
}

var (
	CursorGetType       func() O.Type
	CursorNew           func(cursorType CursorType) *Cursor
	CursorNewForDisplay func(display *Display, cursorType CursorType) *Cursor
	CursorNewFromName   func(display *Display, name string) *Cursor
	CursorNewFromPixbuf func(display *Display, pixbuf *Pixbuf, x int, y int) *Cursor
	CursorNewFromPixmap func(source *Pixmap, mask *Pixmap, fg *Color, bg *Color, x int, y int) *Cursor

	CursorGetCursorType func(cursor *Cursor) CursorType
	CursorGetDisplay    func(cursor *Cursor) *Display
	CursorGetImage      func(cursor *Cursor) *Pixbuf
	CursorRef           func(cursor *Cursor) *Cursor
	CursorUnref         func(cursor *Cursor)
)

func (c *Cursor) GetCursorType() CursorType { return CursorGetCursorType(c) }
func (c *Cursor) GetDisplay() *Display      { return CursorGetDisplay(c) }
func (c *Cursor) GetImage() *Pixbuf         { return CursorGetImage(c) }
func (c *Cursor) Ref() *Cursor              { return CursorRef(c) }
func (c *Cursor) Unref()                    { CursorUnref(c) }

type CursorType Enum

const (
	X_CURSOR CursorType = iota * 2
	ARROW
	BASED_ARROW_DOWN
	BASED_ARROW_UP
	BOAT
	BOGOSITY
	BOTTOM_LEFT_CORNER
	BOTTOM_RIGHT_CORNER
	BOTTOM_SIDE
	BOTTOM_TEE
	BOX_SPIRAL
	CENTER_PTR
	CIRCLE
	CLOCK
	COFFEE_MUG
	CROSS
	CROSS_REVERSE
	CROSSHAIR
	DIAMOND_CROSS
	DOT
	DOTBOX
	DOUBLE_ARROW
	DRAFT_LARGE
	DRAFT_SMALL
	DRAPED_BOX
	EXCHANGE
	FLEUR
	GOBBLER
	GUMBY
	HAND1
	HAND2
	HEART
	ICON
	IRON_CROSS
	LEFT_PTR
	LEFT_SIDE
	LEFT_TEE
	LEFTBUTTON
	LL_ANGLE
	LR_ANGLE
	MAN
	MIDDLEBUTTON
	MOUSE
	PENCIL
	PIRATE
	PLUS
	QUESTION_ARROW
	RIGHT_PTR
	RIGHT_SIDE
	RIGHT_TEE
	RIGHTBUTTON
	RTL_LOGO
	SAILBOAT
	SB_DOWN_ARROW
	SB_H_DOUBLE_ARROW
	SB_LEFT_ARROW
	SB_RIGHT_ARROW
	SB_UP_ARROW
	SB_V_DOUBLE_ARROW
	SHUTTLE
	SIZING
	SPIDER
	SPRAYCAN
	STAR
	TARGET
	TCROSS
	TOP_LEFT_ARROW
	TOP_LEFT_CORNER
	TOP_RIGHT_CORNER
	TOP_SIDE
	TOP_TEE
	TREK
	UL_ANGLE
	UMBRELLA
	UR_ANGLE
	WATCH
	XTERM
	LAST_CURSOR                 = XTERM + 1
	BLANK_CURSOR     CursorType = -2
	CURSOR_IS_PIXMAP CursorType = -1
)

var CursorTypeGetType func() O.Type
