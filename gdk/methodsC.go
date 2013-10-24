// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Cairo struct{} //TODO(t):Def and push methods to /cairo? pragmatic pros/cons

var (
	CairoCreate func(drawable *Drawable) *Cairo

	cairoResetClip       func(c *Cairo, drawable *Drawable)
	cairoSetSourceColor  func(c *Cairo, color *Color)
	cairoSetSourcePixbuf func(c *Cairo, pixbuf *Pixbuf, pixbufX, pixbufY float64)
	cairoSetSourcePixmap func(c *Cairo, pixmap *Pixmap, pixmapX, pixmapY float64)
	cairoSetSourceWindow func(c *Cairo, window *Window, x, y float64)
	cairoRectangle       func(c *Cairo, rectangle *Rectangle)
	cairoRegion          func(c *Cairo, region *Region)
)

func (c *Cairo) ResetClip(drawable *Drawable) { cairoResetClip(c, drawable) }
func (c *Cairo) SetSourceColor(color *Color)  { cairoSetSourceColor(c, color) }
func (c *Cairo) SetSourcePixbuf(pixbuf *Pixbuf, pixbufX, pixbufY float64) {
	cairoSetSourcePixbuf(c, pixbuf, pixbufX, pixbufY)
}
func (c *Cairo) SetSourcePixmap(pixmap *Pixmap, pixmapX, pixmapY float64) {
	cairoSetSourcePixmap(c, pixmap, pixmapX, pixmapY)
}
func (c *Cairo) SetSourceWindow(window *Window, x, y float64) { cairoSetSourceWindow(c, window, x, y) }
func (c *Cairo) Rectangle(rectangle *Rectangle)               { cairoRectangle(c, rectangle) }
func (c *Cairo) Region(region *Region)                        { cairoRegion(c, region) }

type CapStyle Enum

const (
	CAP_NOT_LAST CapStyle = iota
	CAP_BUTT
	CAP_ROUND
	CAP_PROJECTING
)

var CapStyleGetType func() T.GType

type Color struct {
	Pixel T.GUint32
	Red   uint16
	Green uint16
	Blue  uint16
}

var (
	ColorGetType func() T.GType

	ColorParse func(spec string, color *Color) T.Gboolean

	ColorAlloc  func(c *Colormap, color *Color) int
	ColorBlack  func(c *Colormap, color *Color) int
	ColorChange func(c *Colormap, color *Color) int
	ColorsAlloc func(c *Colormap, contiguous T.Gboolean, planes *T.Gulong, nplanes int, pixels *T.Gulong, npixels int) int
	ColorsFree  func(c *Colormap, pixels *T.Gulong, npixels int, planes T.Gulong)
	ColorWhite  func(c *Colormap, color *Color) int

	colorCopy     func(c *Color) *Color
	colorEqual    func(c *Color, colorb *Color) T.Gboolean
	colorFree     func(c *Color)
	colorHash     func(c *Color) uint
	colorToString func(c *Color) string
)

func (c *Color) Copy() *Color                   { return colorCopy(c) }
func (c *Color) Equal(colorb *Color) T.Gboolean { return colorEqual(c, colorb) }
func (c *Color) Free()                          { colorFree(c) }
func (c *Color) Hash() uint                     { return colorHash(c) }
func (c *Color) ToString() string               { return colorToString(c) }

type Colormap struct {
	Parent        T.GObject
	Size          int
	Colors        *Color
	Visual        *Visual
	WindowingData T.Gpointer
}

var (
	ColormapGetType func() T.GType
	ColormapNew     func(visual *Visual, allocate T.Gboolean) *Colormap

	ColormapGetSystem     func() *Colormap
	ColormapGetSystemSize func() int

	ColorsStore func(c *Colormap, colors *Color, ncolors int)

	colormapAllocColor  func(c *Colormap, color *Color, writeable T.Gboolean, bestMatch T.Gboolean) T.Gboolean
	colormapAllocColors func(c *Colormap, colors *Color, nColors int, writeable T.Gboolean, bestMatch T.Gboolean, success *T.Gboolean) int
	colormapChange      func(c *Colormap, ncolors int)
	colormapFreeColors  func(c *Colormap, colors *Color, nColors int)
	colormapGetScreen   func(c *Colormap) *Screen
	colormapGetVisual   func(c *Colormap) *Visual
	colormapQueryColor  func(c *Colormap, pixel T.Gulong, result *Color)
	colormapRef         func(c *Colormap) *Colormap
	colormapUnref       func(c *Colormap)
)

func (c *Colormap) AllocColor(color *Color, writeable T.Gboolean, bestMatch T.Gboolean) T.Gboolean {
	return colormapAllocColor(c, color, writeable, bestMatch)
}
func (c *Colormap) AllocColors(colors *Color, nColors int, writeable T.Gboolean, bestMatch T.Gboolean, success *T.Gboolean) int {
	return colormapAllocColors(c, colors, nColors, writeable, bestMatch, success)
}
func (c *Colormap) Change(ncolors int)                       { colormapChange(c, ncolors) }
func (c *Colormap) FreeColors(colors *Color, nColors int)    { colormapFreeColors(c, colors, nColors) }
func (c *Colormap) GetScreen() *Screen                       { return colormapGetScreen(c) }
func (c *Colormap) GetVisual() *Visual                       { return colormapGetVisual(c) }
func (c *Colormap) QueryColor(pixel T.Gulong, result *Color) { colormapQueryColor(c, pixel, result) }
func (c *Colormap) Ref() *Colormap                           { return colormapRef(c) }
func (c *Colormap) Unref()                                   { colormapUnref(c) }

type Colorspace Enum

const COLORSPACE_RGB Colorspace = 0

var ColorspaceGetType func() T.GType

type CrossingMode Enum

const (
	CROSSING_NORMAL CrossingMode = iota
	CROSSING_GRAB
	CROSSING_UNGRAB
	CROSSING_GTK_GRAB
	CROSSING_GTK_UNGRAB
	CROSSING_STATE_CHANGED
)

var CrossingModeGetType func() T.GType

type Cursor struct {
	Type     CursorType
	RefCount uint
}

var (
	CursorGetType       func() T.GType
	CursorNew           func(cursorType CursorType) *Cursor
	CursorNewForDisplay func(display *Display, cursorType CursorType) *Cursor
	CursorNewFromName   func(display *Display, name string) *Cursor
	CursorNewFromPixbuf func(display *Display, pixbuf *Pixbuf, x int, y int) *Cursor
	CursorNewFromPixmap func(source *Pixmap, mask *Pixmap, fg *Color, bg *Color, x int, y int) *Cursor

	cursorGetCursorType func(cursor *Cursor) CursorType
	cursorGetDisplay    func(cursor *Cursor) *Display
	cursorGetImage      func(cursor *Cursor) *Pixbuf
	cursorRef           func(cursor *Cursor) *Cursor
	cursorUnref         func(cursor *Cursor)
)

func (c *Cursor) GetCursorType() CursorType { return cursorGetCursorType(c) }
func (c *Cursor) GetDisplay() *Display      { return cursorGetDisplay(c) }
func (c *Cursor) GetImage() *Pixbuf         { return cursorGetImage(c) }
func (c *Cursor) Ref() *Cursor              { return cursorRef(c) }
func (c *Cursor) Unref()                    { cursorUnref(c) }

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

var CursorTypeGetType func() T.GType
