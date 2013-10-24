// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Image struct {
	Parent        T.GObject
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
	ImageGetType func() T.GType
	ImageNew     func(typ ImageType, visual *Visual, width, height int) *Image

	ImageGet func(drawable *Drawable, x, y, width, height int) *Image

	imageGetBitsPerPixel  func(image *Image) uint16
	imageGetByteOrder     func(image *Image) T.GdkByteOrder
	imageGetBytesPerLine  func(image *Image) uint16
	imageGetBytesPerPixel func(image *Image) uint16
	imageGetColormap      func(image *Image) *Colormap
	imageGetDepth         func(image *Image) uint16
	imageGetHeight        func(image *Image) int
	imageGetImageType     func(image *Image) ImageType
	imageGetPixel         func(image *Image, x, y int) T.GUint32
	imageGetPixels        func(image *Image) T.Gpointer
	imageGetVisual        func(image *Image) *Visual
	imageGetWidth         func(image *Image) int
	imagePutPixel         func(image *Image, x, y int, pixel T.GUint32)
	imageRef              func(image *Image) *Image
	imageSetColormap      func(image *Image, colormap *Colormap)
	imageUnref            func(image *Image)
)

func (i *Image) GetBitsPerPixel() uint16            { return imageGetBitsPerPixel(i) }
func (i *Image) GetByteOrder() T.GdkByteOrder       { return imageGetByteOrder(i) }
func (i *Image) GetBytesPerLine() uint16            { return imageGetBytesPerLine(i) }
func (i *Image) GetBytesPerPixel() uint16           { return imageGetBytesPerPixel(i) }
func (i *Image) GetColormap() *Colormap             { return imageGetColormap(i) }
func (i *Image) GetDepth() uint16                   { return imageGetDepth(i) }
func (i *Image) GetHeight() int                     { return imageGetHeight(i) }
func (i *Image) GetImageType() ImageType            { return imageGetImageType(i) }
func (i *Image) GetPixel(x, y int) T.GUint32        { return imageGetPixel(i, x, y) }
func (i *Image) GetPixels() T.Gpointer              { return imageGetPixels(i) }
func (i *Image) GetVisual() *Visual                 { return imageGetVisual(i) }
func (i *Image) GetWidth() int                      { return imageGetWidth(i) }
func (i *Image) PutPixel(x, y int, pixel T.GUint32) { imagePutPixel(i, x, y, pixel) }
func (i *Image) Ref() *Image                        { return imageRef(i) }
func (i *Image) SetColormap(colormap *Colormap)     { imageSetColormap(i, colormap) }
func (i *Image) Unref()                             { imageUnref(i) }

type ImageType Enum

const (
	IMAGE_NORMAL ImageType = iota
	IMAGE_SHARED
	IMAGE_FASTEST
)

var ImageTypeGetType func() T.GType
