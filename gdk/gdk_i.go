// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Image struct {
	Parent        O.Object
	Type          ImageType
	Visual        *Visual
	ByteOrder     T.GdkByteOrder
	Width         int
	Height        int
	Depth         uint16
	Bpp           uint16
	Bpl           uint16
	BitsPerPixel  uint16
	Mem           T.Gpointer
	Colormap      *Colormap
	WindowingData T.Gpointer
}

var (
	ImageGetType func() O.Type
	ImageNew     func(typ ImageType, visual *Visual, width, height int) *Image

	ImageGet func(drawable *Drawable, x, y, width, height int) *Image

	ImageGetBitsPerPixel  func(image *Image) uint16
	ImageGetByteOrder     func(image *Image) T.GdkByteOrder
	ImageGetBytesPerLine  func(image *Image) uint16
	ImageGetBytesPerPixel func(image *Image) uint16
	ImageGetColormap      func(image *Image) *Colormap
	ImageGetDepth         func(image *Image) uint16
	ImageGetHeight        func(image *Image) int
	ImageGetImageType     func(image *Image) ImageType
	ImageGetPixel         func(image *Image, x, y int) T.GUint32
	ImageGetPixels        func(image *Image) T.Gpointer
	ImageGetVisual        func(image *Image) *Visual
	ImageGetWidth         func(image *Image) int
	ImagePutPixel         func(image *Image, x, y int, pixel T.GUint32)
	ImageRef              func(image *Image) *Image
	ImageSetColormap      func(image *Image, colormap *Colormap)
	ImageUnref            func(image *Image)
)

func (i *Image) GetBitsPerPixel() uint16            { return ImageGetBitsPerPixel(i) }
func (i *Image) GetByteOrder() T.GdkByteOrder       { return ImageGetByteOrder(i) }
func (i *Image) GetBytesPerLine() uint16            { return ImageGetBytesPerLine(i) }
func (i *Image) GetBytesPerPixel() uint16           { return ImageGetBytesPerPixel(i) }
func (i *Image) GetColormap() *Colormap             { return ImageGetColormap(i) }
func (i *Image) GetDepth() uint16                   { return ImageGetDepth(i) }
func (i *Image) GetHeight() int                     { return ImageGetHeight(i) }
func (i *Image) GetImageType() ImageType            { return ImageGetImageType(i) }
func (i *Image) GetPixel(x, y int) T.GUint32        { return ImageGetPixel(i, x, y) }
func (i *Image) GetPixels() T.Gpointer              { return ImageGetPixels(i) }
func (i *Image) GetVisual() *Visual                 { return ImageGetVisual(i) }
func (i *Image) GetWidth() int                      { return ImageGetWidth(i) }
func (i *Image) PutPixel(x, y int, pixel T.GUint32) { ImagePutPixel(i, x, y, pixel) }
func (i *Image) Ref() *Image                        { return ImageRef(i) }
func (i *Image) SetColormap(colormap *Colormap)     { ImageSetColormap(i, colormap) }
func (i *Image) Unref()                             { ImageUnref(i) }

type ImageType Enum

const (
	IMAGE_NORMAL ImageType = iota
	IMAGE_SHARED
	IMAGE_FASTEST
)

var ImageTypeGetType func() O.Type
