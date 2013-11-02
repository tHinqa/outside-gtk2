// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	I "github.com/tHinqa/outside-gtk2/gio"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

var (
	PangoAttrEmbossColorNew func(color *Color) *P.Attribute
	PangoAttrEmbossedNew    func(embossed bool) *P.Attribute

	PangoContextGet          func() *P.Context
	PangoContextGetForScreen func(screen *Screen) *P.Context
	PangoContextSetColormap  func(context *P.Context, colormap *Colormap)

	PangoLayoutGetClipRegion     func(layout *P.Layout, xOrigin, yOrigin int, indexRanges *int, nRanges int) *Region
	PangoLayoutLineGetClipRegion func(line *P.LayoutLine, xOrigin, yOrigin int, indexRanges *int, nRanges int) *Region
)

type PangoRenderer struct {
	Parent P.Renderer
	_      *struct{}
}

var (
	PangoRendererGetType func() O.Type
	PangoRendererNew     func(screen *Screen) *P.Renderer

	PangoRendererGetDefault func(screen *Screen) *P.Renderer

	PangoRendererSetDrawable      func(p *PangoRenderer, drawable *Drawable)
	PangoRendererSetGc            func(p *PangoRenderer, gc *GC)
	PangoRendererSetStipple       func(p *PangoRenderer, part P.RenderPart, stipple *T.GdkBitmap)
	PangoRendererSetOverrideColor func(p *PangoRenderer, part P.RenderPart, color *Color)
)

func (p *PangoRenderer) SetDrawable(drawable *Drawable) { PangoRendererSetDrawable(p, drawable) }
func (p *PangoRenderer) SetGc(gc *GC)                   { PangoRendererSetGc(p, gc) }
func (p *PangoRenderer) SetStipple(part P.RenderPart, stipple *T.GdkBitmap) {
	PangoRendererSetStipple(p, part, stipple)
}
func (p *PangoRenderer) SetOverrideColor(part P.RenderPart, color *Color) {
	PangoRendererSetOverrideColor(p, part, color)
}

var PangoAttrStippleNew func(stipple *T.GdkBitmap) *P.Attribute

type Pixbuf struct{}

var (
	PixbufGetType                   func() O.Type
	PixbufNew                       func(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) *Pixbuf
	PixbufNewFromData               func(data *T.Guchar, colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height, rowstride int, destroyFn PixbufDestroyNotify, destroyFnData T.Gpointer) *Pixbuf
	PixbufNewFromFile               func(filename string, e **L.Error) *Pixbuf
	PixbufNewFromFileAtScale        func(filename string, width, height int, preserveAspectRatio bool, e **L.Error) *Pixbuf
	PixbufNewFromFileAtSize         func(filename string, width, height int, e **L.Error) *Pixbuf
	PixbufNewFromInline             func(dataLength int, data *uint8, copyPixels bool, e **L.Error) *Pixbuf
	PixbufNewFromStream             func(stream *I.InputStream, cancellable *I.Cancellable, e **L.Error) *Pixbuf
	PixbufNewFromStreamAsync        func(stream *I.InputStream, cancellable *I.Cancellable, callback I.AsyncReadyCallback, userData T.Gpointer)
	PixbufNewFromStreamAtScale      func(stream *I.InputStream, width, height int, preserveAspectRatio bool, cancellable *I.Cancellable, e **L.Error) *Pixbuf
	PixbufNewFromStreamAtScaleAsync func(stream *I.InputStream, width, height int, preserveAspectRatio bool, cancellable *I.Cancellable, callback I.AsyncReadyCallback, userData T.Gpointer)
	PixbufNewFromStreamFinish       func(asyncResult *I.AsyncResult, e **L.Error) *Pixbuf
	PixbufNewFromXpmData            func(data **T.Char) *Pixbuf

	PixbufErrorQuark         func() L.Quark
	PixbufGetFileInfo        func(filename string, width, height *int) *PixbufFormat
	PixbufGetFormats         func() *L.SList
	PixbufSaveToStreamFinish func(asyncResult *I.AsyncResult, e **L.Error) bool

	PixbufGetFromDrawable func(dest *Pixbuf, src *Drawable, cmap *Colormap, srcX, srcY, destX, destY, width, height int) *Pixbuf
	PixbufGetFromImage    func(dest *Pixbuf, src *Image, cmap *Colormap, srcX, srcY, destX, destY, width, height int) *Pixbuf

	PixbufAddAlpha                 func(pixbuf *Pixbuf, substituteColor bool, r, g, b T.Guchar) *Pixbuf
	PixbufApplyEmbeddedOrientation func(src *Pixbuf) *Pixbuf
	PixbufComposite                func(src, dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType, overallAlpha int)
	PixbufCompositeColor           func(src, dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType, overallAlpha, checkX, checkY, checkSize int, color1, color2 T.GUint32)
	PixbufCompositeColorSimple     func(src *Pixbuf, destWidth, destHeight int, interpType T.GdkInterpType, overallAlpha, checkSize int, color1, color2 T.GUint32) *Pixbuf
	PixbufCopy                     func(pixbuf *Pixbuf) *Pixbuf
	PixbufCopyArea                 func(srcPixbuf *Pixbuf, srcX, srcY, width, height int, destPixbuf *Pixbuf, destX, destY int)
	PixbufFill                     func(pixbuf *Pixbuf, pixel T.GUint32)
	PixbufFlip                     func(src *Pixbuf, horizontal bool) *Pixbuf
	PixbufGetBitsPerSample         func(pixbuf *Pixbuf) int
	PixbufGetColorspace            func(pixbuf *Pixbuf) Colorspace
	PixbufGetHasAlpha              func(pixbuf *Pixbuf) bool
	PixbufGetHeight                func(pixbuf *Pixbuf) int
	PixbufGetNChannels             func(pixbuf *Pixbuf) int
	PixbufGetOption                func(pixbuf *Pixbuf, key string) string
	PixbufGetPixels                func(pixbuf *Pixbuf) *T.Guchar
	PixbufGetRowstride             func(pixbuf *Pixbuf) int
	PixbufGetWidth                 func(pixbuf *Pixbuf) int
	PixbufNewSubpixbuf             func(pixbuf *Pixbuf, srcX, srcY, width, height int) *Pixbuf
	PixbufRef                      func(pixbuf *Pixbuf) *Pixbuf
	PixbufRotateSimple             func(src *Pixbuf, angle PixbufRotation) *Pixbuf
	PixbufSaturateAndPixelate      func(src, dest *Pixbuf, saturation float32, pixelate bool)
	PixbufSaveToBuffer             func(pixbuf *Pixbuf, buffer **T.Gchar, bufferSize *T.Gsize, typ string, e **L.Error, v ...VArg) bool
	PixbufSaveToBufferv            func(pixbuf *Pixbuf, buffer **T.Gchar, bufferSize *T.Gsize, typ string, optionKeys, optionValues **T.Char, e **L.Error) bool
	PixbufSaveToCallback           func(pixbuf *Pixbuf, saveFunc PixbufSaveFunc, userData T.Gpointer, typ string, err **L.Error, v ...VArg) bool
	PixbufSaveToCallbackv          func(pixbuf *Pixbuf, saveFunc PixbufSaveFunc, userData T.Gpointer, typ string, optionKeys, optionValues **T.Char, e **L.Error) bool
	PixbufSaveToStream             func(pixbuf *Pixbuf, stream *I.OutputStream, typ string, cancellable *I.Cancellable, e **L.Error, v ...VArg) bool
	PixbufSaveToStreamAsync        func(pixbuf *Pixbuf, stream *I.OutputStream, typ string, cancellable *I.Cancellable, callback I.AsyncReadyCallback, userData T.Gpointer, v ...VArg)
	PixbufSaveUtf8                 func(pixbuf *Pixbuf, filename, typ string, e **L.Error, v ...VArg) bool
	PixbufSavevUtf8                func(pixbuf *Pixbuf, filename, typ string, optionKeys, optionValues **T.Char, e **L.Error) bool
	PixbufScale                    func(src, dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType)
	PixbufScaleSimple              func(src *Pixbuf, destWidth, destHeight int, interpType T.GdkInterpType) *Pixbuf
	PixbufSetOption                func(pixbuf *Pixbuf, key, value string) bool
	PixbufUnref                    func(pixbuf *Pixbuf)
)

func (p *Pixbuf) AddAlpha(substituteColor bool, r, g, b T.Guchar) *Pixbuf {
	return PixbufAddAlpha(p, substituteColor, r, g, b)
}
func (p *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf { return PixbufApplyEmbeddedOrientation(p) }
func (p *Pixbuf) Composite(dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType, overallAlpha int) {
	PixbufComposite(p, dest, destX, destY, destWidth, destHeight, offsetX, offsetY, scaleX, scaleY, interpType, overallAlpha)
}
func (p *Pixbuf) CompositeColor(dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType, overallAlpha, checkX, checkY, checkSize int, color1, color2 T.GUint32) {
	PixbufCompositeColor(p, dest, destX, destY, destWidth, destHeight, offsetX, offsetY, scaleX, scaleY, interpType, overallAlpha, checkX, checkY, checkSize, color1, color2)
}
func (p *Pixbuf) CompositeColorSimple(destWidth, destHeight int, interpType T.GdkInterpType, overallAlpha, checkSize int, color1, color2 T.GUint32) *Pixbuf {
	return PixbufCompositeColorSimple(p, destWidth, destHeight, interpType, overallAlpha, checkSize, color1, color2)
}
func (p *Pixbuf) Copy() *Pixbuf { return PixbufCopy(p) }
func (p *Pixbuf) CopyArea(srcX, srcY, width, height int, destPixbuf *Pixbuf, destX, destY int) {
	PixbufCopyArea(p, srcX, srcY, width, height, destPixbuf, destX, destY)
}
func (p *Pixbuf) Fill(pixel T.GUint32)         { PixbufFill(p, pixel) }
func (p *Pixbuf) Flip(horizontal bool) *Pixbuf { return PixbufFlip(p, horizontal) }
func (p *Pixbuf) GetBitsPerSample() int        { return PixbufGetBitsPerSample(p) }
func (p *Pixbuf) GetColorspace() Colorspace    { return PixbufGetColorspace(p) }
func (p *Pixbuf) GetHasAlpha() bool            { return PixbufGetHasAlpha(p) }
func (p *Pixbuf) GetHeight() int               { return PixbufGetHeight(p) }
func (p *Pixbuf) GetNChannels() int            { return PixbufGetNChannels(p) }
func (p *Pixbuf) GetOption(key string) string  { return PixbufGetOption(p, key) }
func (p *Pixbuf) GetPixels() *T.Guchar         { return PixbufGetPixels(p) }
func (p *Pixbuf) GetRowstride() int            { return PixbufGetRowstride(p) }
func (p *Pixbuf) GetWidth() int                { return PixbufGetWidth(p) }
func (p *Pixbuf) NewSubpixbuf(srcX, srcY, width, height int) *Pixbuf {
	return PixbufNewSubpixbuf(p, srcX, srcY, width, height)
}
func (p *Pixbuf) Ref() *Pixbuf                              { return PixbufRef(p) }
func (p *Pixbuf) RotateSimple(angle PixbufRotation) *Pixbuf { return PixbufRotateSimple(p, angle) }
func (p *Pixbuf) SaturateAndPixelate(dest *Pixbuf, saturation float32, pixelate bool) {
	PixbufSaturateAndPixelate(p, dest, saturation, pixelate)
}
func (p *Pixbuf) SaveToBuffer(buffer **T.Gchar, bufferSize *T.Gsize, typ string, e **L.Error, v ...VArg) bool {
	return PixbufSaveToBuffer(p, buffer, bufferSize, typ, e, v)
}
func (p *Pixbuf) SaveToBufferv(buffer **T.Gchar, bufferSize *T.Gsize, typ string, optionKeys, optionValues **T.Char, e **L.Error) bool {
	return PixbufSaveToBufferv(p, buffer, bufferSize, typ, optionKeys, optionValues, e)
}
func (p *Pixbuf) SaveToCallback(saveFunc PixbufSaveFunc, userData T.Gpointer, typ string, err **L.Error, v ...VArg) bool {
	return PixbufSaveToCallback(p, saveFunc, userData, typ, err, v)
}
func (p *Pixbuf) SaveToCallbackv(saveFunc PixbufSaveFunc, userData T.Gpointer, typ string, optionKeys, optionValues **T.Char, e **L.Error) bool {
	return PixbufSaveToCallbackv(p, saveFunc, userData, typ, optionKeys, optionValues, e)
}
func (p *Pixbuf) SaveToStream(stream *I.OutputStream, typ string, cancellable *I.Cancellable, e **L.Error, v ...VArg) bool {
	return PixbufSaveToStream(p, stream, typ, cancellable, e, v)
}
func (p *Pixbuf) SaveToStreamAsync(stream *I.OutputStream, typ string, cancellable *I.Cancellable, callback I.AsyncReadyCallback, userData T.Gpointer, v ...VArg) {
	PixbufSaveToStreamAsync(p, stream, typ, cancellable, callback, userData, v)
}
func (p *Pixbuf) SaveUtf8(filename, typ string, e **L.Error, v ...VArg) bool {
	return PixbufSaveUtf8(p, filename, typ, e, v)
}
func (p *Pixbuf) SavevUtf8(filename, typ string, optionKeys, optionValues **T.Char, e **L.Error) bool {
	return PixbufSavevUtf8(p, filename, typ, optionKeys, optionValues, e)
}
func (p *Pixbuf) Scale(dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType) {
	PixbufScale(p, dest, destX, destY, destWidth, destHeight, offsetX, offsetY, scaleX, scaleY, interpType)
}
func (p *Pixbuf) ScaleSimple(destWidth, destHeight int, interpType T.GdkInterpType) *Pixbuf {
	return PixbufScaleSimple(p, destWidth, destHeight, interpType)
}
func (p *Pixbuf) SetOption(key, value string) bool { return PixbufSetOption(p, key, value) }
func (p *Pixbuf) Unref()                           { PixbufUnref(p) }

var (
	PixbufErrorGetType func() O.Type
)

type PixbufAlphaMode Enum

const (
	PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = iota
	PIXBUF_ALPHA_FULL
)

var PixbufAlphaModeGetType func() O.Type

type PixbufAnimation struct{}

var (
	PixbufAnimationGetType         func() O.Type
	PixbufAnimationNewFromFileUtf8 func(filename string, e **L.Error) *PixbufAnimation

	PixbufAnimationGetHeight      func(animation *PixbufAnimation) int
	PixbufAnimationGetIter        func(animation *PixbufAnimation, startTime *L.TimeVal) *PixbufAnimationIter
	PixbufAnimationGetStaticImage func(animation *PixbufAnimation) *Pixbuf
	PixbufAnimationGetWidth       func(animation *PixbufAnimation) int
	PixbufAnimationIsStaticImage  func(animation *PixbufAnimation) bool
	PixbufAnimationRef            func(animation *PixbufAnimation) *PixbufAnimation
	PixbufAnimationUnref          func(animation *PixbufAnimation)
)

func (p *PixbufAnimation) GetHeight() int { return PixbufAnimationGetHeight(p) }
func (p *PixbufAnimation) GetIter(startTime *L.TimeVal) *PixbufAnimationIter {
	return PixbufAnimationGetIter(p, startTime)
}
func (p *PixbufAnimation) GetStaticImage() *Pixbuf { return PixbufAnimationGetStaticImage(p) }
func (p *PixbufAnimation) GetWidth() int           { return PixbufAnimationGetWidth(p) }
func (p *PixbufAnimation) IsStaticImage() bool     { return PixbufAnimationIsStaticImage(p) }
func (p *PixbufAnimation) Ref() *PixbufAnimation   { return PixbufAnimationRef(p) }
func (p *PixbufAnimation) Unref()                  { PixbufAnimationUnref(p) }

type PixbufAnimationIter struct{}

var (
	PixbufAnimationIterGetType func() O.Type
	PixbufAnimationNewFromFile func(filename string, e **L.Error) *PixbufAnimation

	PixbufAnimationIterAdvance                 func(iter *PixbufAnimationIter, currentTime *L.TimeVal) bool
	PixbufAnimationIterGetDelayTime            func(iter *PixbufAnimationIter) int
	PixbufAnimationIterGetPixbuf               func(iter *PixbufAnimationIter) *Pixbuf
	PixbufAnimationIterOnCurrentlyLoadingFrame func(iter *PixbufAnimationIter) bool
)

type PixbufDestroyNotify func(
	Pixels *T.Guchar,
	data T.Gpointer)

type PixbufFormat struct{}

var (
	PixbufFormatGetType func() O.Type

	PixbufFormatCopy           func(format *PixbufFormat) *PixbufFormat
	PixbufFormatFree           func(format *PixbufFormat)
	PixbufFormatGetDescription func(format *PixbufFormat) string
	PixbufFormatGetExtensions  func(format *PixbufFormat) []string
	PixbufFormatGetLicense     func(format *PixbufFormat) string
	PixbufFormatGetMimeTypes   func(format *PixbufFormat) []string
	PixbufFormatGetName        func(format *PixbufFormat) string
	PixbufFormatIsDisabled     func(format *PixbufFormat) bool
	PixbufFormatIsScalable     func(format *PixbufFormat) bool
	PixbufFormatIsWritable     func(format *PixbufFormat) bool
	PixbufFormatSetDisabled    func(format *PixbufFormat, disabled bool)
)

type PixbufLoader struct {
	Parent O.Object
	_      T.Gpointer
}

var (
	PixbufLoaderGetType         func() O.Type
	PixbufLoaderNew             func() *PixbufLoader
	PixbufLoaderNewWithMimeType func(mimeType string, e **L.Error) *PixbufLoader
	PixbufLoaderNewWithType     func(imageType string, e **L.Error) *PixbufLoader

	PixbufLoaderClose        func(loader *PixbufLoader, e **L.Error) bool
	PixbufLoaderGetAnimation func(loader *PixbufLoader) *PixbufAnimation
	PixbufLoaderGetFormat    func(loader *PixbufLoader) *PixbufFormat
	PixbufLoaderGetPixbuf    func(loader *PixbufLoader) *Pixbuf
	PixbufLoaderSetSize      func(loader *PixbufLoader, width, height int)
	PixbufLoaderWrite        func(loader *PixbufLoader, buf *T.Guchar, count T.Gsize, e **L.Error) bool
)

var (
	PixbufNonAnimGetType func() O.Type
	PixbufNonAnimNew     func(pixbuf *Pixbuf) *PixbufAnimation
)

var (
	PixbufRenderPixmapAndMask            func(pixbuf *Pixbuf, pixmapReturn **Pixmap, maskReturn **T.GdkBitmap, alphaThreshold int)
	PixbufRenderPixmapAndMaskForColormap func(pixbuf *Pixbuf, colormap *Colormap, pixmapReturn **Pixmap, maskReturn **T.GdkBitmap, alphaThreshold int)
	PixbufRenderThresholdAlpha           func(pixbuf *Pixbuf, bitmap *T.GdkBitmap, srcX, srcY, destX, destY, width, height, alphaThreshold int)
	PixbufRenderToDrawable               func(pixbuf *Pixbuf, drawable *Drawable, gc *GC, srcX, srcY, destX, destY, width, height int, dither T.GdkRgbDither, xDither, yDither int)
	PixbufRenderToDrawableAlpha          func(pixbuf *Pixbuf, drawable *Drawable, srcX, srcY, destX, destY, width, height int, alphaMode PixbufAlphaMode, alphaThreshold int, dither T.GdkRgbDither, xDither, yDither int)
)

type PixbufRotation Enum

const (
	PIXBUF_ROTATE_NONE             PixbufRotation = 0
	PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = 180
	PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = 270
)

var PixbufRotationGetType func() O.Type

type PixbufSaveFunc func(
	buf string,
	count T.Gsize,
	err **L.Error,
	data T.Gpointer) bool

type PixbufSimpleAnim struct{}

var (
	PixbufSimpleAnimGetType func() O.Type
	PixbufSimpleAnimNew     func(width, height int, rate float32) *PixbufSimpleAnim

	PixbufSimpleAnimAddFrame func(animation *PixbufSimpleAnim, pixbuf *Pixbuf)
	PixbufSimpleAnimSetLoop  func(animation *PixbufSimpleAnim, loop bool)
	PixbufSimpleAnimGetLoop  func(animation *PixbufSimpleAnim) bool
)

func (p *PixbufSimpleAnim) AddFrame(pixbuf *Pixbuf) { PixbufSimpleAnimAddFrame(p, pixbuf) }
func (p *PixbufSimpleAnim) SetLoop(loop bool)       { PixbufSimpleAnimSetLoop(p, loop) }
func (p *PixbufSimpleAnim) GetLoop() bool           { return PixbufSimpleAnimGetLoop(p) }

var PixbufSimpleAnimIterGetType func() O.Type

type Pixdata struct {
	Magic       T.GUint32
	Length      T.GInt32
	PixdataType PixdataType
	Rowstride   T.GUint32
	Width       T.GUint32
	Height      T.GUint32
	PixelData   *uint8
}

var (
	PixbufFromPixdata func(p *Pixdata, copyPixels bool, e **L.Error) *Pixbuf

	PixdataDeserialize func(p *Pixdata, streamLength uint, stream *uint8, e **L.Error) bool
	PixdataFromPixbuf  func(p *Pixdata, pixbuf *Pixbuf, useRle bool) T.Gpointer
	PixdataSerialize   func(p *Pixdata, streamLengthP *uint) *uint8
	PixdataToCsource   func(p *Pixdata, name string, dumpType PixdataDumpType) *L.String
)

func (p *Pixdata) Deserialize(streamLength uint, stream *uint8, e **L.Error) bool {
	return PixdataDeserialize(p, streamLength, stream, e)
}
func (p *Pixdata) FromPixbuf(pixbuf *Pixbuf, useRle bool) T.Gpointer {
	return PixdataFromPixbuf(p, pixbuf, useRle)
}
func (p *Pixdata) Serialize(streamLengthP *uint) *uint8 { return PixdataSerialize(p, streamLengthP) }
func (p *Pixdata) ToCsource(name string, dumpType PixdataDumpType) *L.String {
	return PixdataToCsource(p, name, dumpType)
}

type PixdataType Enum

const (
	PIXDATA_COLOR_TYPE_RGB PixdataType = 0x01 << iota
	PIXDATA_COLOR_TYPE_RGBA
	PIXDATA_COLOR_TYPE_MASK PixdataType = 0xff
)
const (
	PIXDATA_SAMPLE_WIDTH_8    PixdataType = 0x01 << 16
	PIXDATA_SAMPLE_WIDTH_MASK PixdataType = 0x0f << 16
)
const (
	PIXDATA_ENCODING_RAW PixdataType = 0x01 << (24 + iota)
	PIXDATA_ENCODING_RLE
	PIXDATA_ENCODING_MASK PixdataType = 0x0f << 24
)

type PixdataDumpType Enum

const (
	PIXDATA_DUMP_PIXDATA_STRUCT PixdataDumpType = 1 << iota
	PIXDATA_DUMP_MACROS
	_
	_
	_
	_
	_
	_
	PIXDATA_DUMP_CTYPES
	PIXDATA_DUMP_STATIC
	PIXDATA_DUMP_CONST
	_
	_
	_
	_
	_
	PIXDATA_DUMP_RLE_DECODER
	PIXDATA_DUMP_PIXDATA_STREAM PixdataDumpType = 0
	PIXDATA_DUMP_GTYPES         PixdataDumpType = 0
)

type Pixmap Drawable

var (
	PixmapGetType func() O.Type
	PixmapNew     func(drawable *Drawable, width, height, depth int) *Pixmap

	PixmapColormapCreateFromXpm  func(drawable *Drawable, colormap *Colormap, mask **T.GdkBitmap, transparentColor *Color, filename string) *Pixmap
	PixmapColormapCreateFromXpmD func(drawable *Drawable, colormap *Colormap, mask **T.GdkBitmap, transparentColor *Color, data **T.Gchar) *Pixmap
	PixmapCreateFromData         func(drawable *Drawable, data string, width, height, depth int, fg, bg *Color) *Pixmap
	PixmapCreateFromXpm          func(drawable *Drawable, mask **T.GdkBitmap, transparentColor *Color, filename string) *Pixmap
	PixmapCreateFromXpmD         func(drawable *Drawable, mask **T.GdkBitmap, transparentColor *Color, data **T.Gchar) *Pixmap
	PixmapForeignNew             func(anid T.GdkNativeWindow) *Pixmap
	PixmapForeignNewForDisplay   func(display *Display, anid T.GdkNativeWindow) *Pixmap
	PixmapForeignNewForScreen    func(screen *Screen, anid T.GdkNativeWindow, width, height, depth int) *Pixmap
	PixmapLookup                 func(anid T.GdkNativeWindow) *Pixmap
	PixmapLookupForDisplay       func(display *Display, anid T.GdkNativeWindow) *Pixmap

	PixmapGetSize func(p *Pixmap, width, height *int)
)

func (p *Pixmap) GetSize(width, height *int) { PixmapGetSize(p, width, height) }

type Point struct {
	X int
	Y int
}

type PropMode Enum

const (
	PROP_MODE_REPLACE PropMode = iota
	PROP_MODE_PREPEND
	PROP_MODE_APPEND
)

var PropModeGetType func() O.Type

type PropertyState Enum

const (
	PROPERTY_NEW_VALUE PropertyState = iota
	PROPERTY_DELETE
)

var PropertyStateGetType func() O.Type
