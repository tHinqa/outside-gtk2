// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	I "github.com/tHinqa/outside-gtk2/gio"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

var (
	PangoAttrEmbossColorNew func(color *Color) *P.Attribute
	PangoAttrEmbossedNew    func(embossed T.Gboolean) *P.Attribute

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

	pangoRendererSetDrawable      func(p *PangoRenderer, drawable *Drawable)
	pangoRendererSetGc            func(p *PangoRenderer, gc *GC)
	pangoRendererSetStipple       func(p *PangoRenderer, part P.RenderPart, stipple *T.GdkBitmap)
	pangoRendererSetOverrideColor func(p *PangoRenderer, part P.RenderPart, color *Color)
)

func (p *PangoRenderer) SetDrawable(drawable *Drawable) { pangoRendererSetDrawable(p, drawable) }
func (p *PangoRenderer) SetGc(gc *GC)                   { pangoRendererSetGc(p, gc) }
func (p *PangoRenderer) SetStipple(part P.RenderPart, stipple *T.GdkBitmap) {
	pangoRendererSetStipple(p, part, stipple)
}
func (p *PangoRenderer) SetOverrideColor(part P.RenderPart, color *Color) {
	pangoRendererSetOverrideColor(p, part, color)
}

var PangoAttrStippleNew func(stipple *T.GdkBitmap) *P.Attribute

type Pixbuf struct{}

var (
	PixbufGetType                   func() O.Type
	PixbufNew                       func(colorspace Colorspace, hasAlpha T.Gboolean, bitsPerSample, width, height int) *Pixbuf
	PixbufNewFromData               func(data *T.Guchar, colorspace Colorspace, hasAlpha T.Gboolean, bitsPerSample, width, height, rowstride int, destroyFn PixbufDestroyNotify, destroyFnData T.Gpointer) *Pixbuf
	PixbufNewFromFile               func(filename string, e **T.GError) *Pixbuf
	PixbufNewFromFileAtScale        func(filename string, width, height int, preserveAspectRatio T.Gboolean, e **T.GError) *Pixbuf
	PixbufNewFromFileAtSize         func(filename string, width, height int, e **T.GError) *Pixbuf
	PixbufNewFromInline             func(dataLength int, data *uint8, copyPixels T.Gboolean, e **T.GError) *Pixbuf
	PixbufNewFromStream             func(stream *I.InputStream, cancellable *I.Cancellable, e **T.GError) *Pixbuf
	PixbufNewFromStreamAsync        func(stream *I.InputStream, cancellable *I.Cancellable, callback I.AsyncReadyCallback, userData T.Gpointer)
	PixbufNewFromStreamAtScale      func(stream *I.InputStream, width, height int, preserveAspectRatio T.Gboolean, cancellable *I.Cancellable, e **T.GError) *Pixbuf
	PixbufNewFromStreamAtScaleAsync func(stream *I.InputStream, width, height int, preserveAspectRatio T.Gboolean, cancellable *I.Cancellable, callback I.AsyncReadyCallback, userData T.Gpointer)
	PixbufNewFromStreamFinish       func(asyncResult *I.AsyncResult, e **T.GError) *Pixbuf
	PixbufNewFromXpmData            func(data **T.Char) *Pixbuf

	PixbufErrorQuark         func() T.GQuark
	PixbufGetFileInfo        func(filename string, width, height *int) *PixbufFormat
	PixbufGetFormats         func() *T.GSList
	PixbufSaveToStreamFinish func(asyncResult *I.AsyncResult, e **T.GError) T.Gboolean

	PixbufGetFromDrawable func(dest *Pixbuf, src *Drawable, cmap *Colormap, srcX, srcY, destX, destY, width, height int) *Pixbuf
	PixbufGetFromImage    func(dest *Pixbuf, src *Image, cmap *Colormap, srcX, srcY, destX, destY, width, height int) *Pixbuf

	pixbufAddAlpha                 func(pixbuf *Pixbuf, substituteColor T.Gboolean, r, g, b T.Guchar) *Pixbuf
	pixbufApplyEmbeddedOrientation func(src *Pixbuf) *Pixbuf
	pixbufComposite                func(src, dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType, overallAlpha int)
	pixbufCompositeColor           func(src, dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType, overallAlpha, checkX, checkY, checkSize int, color1, color2 T.GUint32)
	pixbufCompositeColorSimple     func(src *Pixbuf, destWidth, destHeight int, interpType T.GdkInterpType, overallAlpha, checkSize int, color1, color2 T.GUint32) *Pixbuf
	pixbufCopy                     func(pixbuf *Pixbuf) *Pixbuf
	pixbufCopyArea                 func(srcPixbuf *Pixbuf, srcX, srcY, width, height int, destPixbuf *Pixbuf, destX, destY int)
	pixbufFill                     func(pixbuf *Pixbuf, pixel T.GUint32)
	pixbufFlip                     func(src *Pixbuf, horizontal T.Gboolean) *Pixbuf
	pixbufGetBitsPerSample         func(pixbuf *Pixbuf) int
	pixbufGetColorspace            func(pixbuf *Pixbuf) Colorspace
	pixbufGetHasAlpha              func(pixbuf *Pixbuf) T.Gboolean
	pixbufGetHeight                func(pixbuf *Pixbuf) int
	pixbufGetNChannels             func(pixbuf *Pixbuf) int
	pixbufGetOption                func(pixbuf *Pixbuf, key string) string
	pixbufGetPixels                func(pixbuf *Pixbuf) *T.Guchar
	pixbufGetRowstride             func(pixbuf *Pixbuf) int
	pixbufGetWidth                 func(pixbuf *Pixbuf) int
	pixbufNewSubpixbuf             func(pixbuf *Pixbuf, srcX, srcY, width, height int) *Pixbuf
	pixbufRef                      func(pixbuf *Pixbuf) *Pixbuf
	pixbufRotateSimple             func(src *Pixbuf, angle PixbufRotation) *Pixbuf
	pixbufSaturateAndPixelate      func(src, dest *Pixbuf, saturation float32, pixelate T.Gboolean)
	pixbufSaveToBuffer             func(pixbuf *Pixbuf, buffer **T.Gchar, bufferSize *T.Gsize, typ string, e **T.GError, v ...VArg) T.Gboolean
	pixbufSaveToBufferv            func(pixbuf *Pixbuf, buffer **T.Gchar, bufferSize *T.Gsize, typ string, optionKeys, optionValues **T.Char, e **T.GError) T.Gboolean
	pixbufSaveToCallback           func(pixbuf *Pixbuf, saveFunc PixbufSaveFunc, userData T.Gpointer, typ string, err **T.GError, v ...VArg) T.Gboolean
	pixbufSaveToCallbackv          func(pixbuf *Pixbuf, saveFunc PixbufSaveFunc, userData T.Gpointer, typ string, optionKeys, optionValues **T.Char, e **T.GError) T.Gboolean
	pixbufSaveToStream             func(pixbuf *Pixbuf, stream *I.OutputStream, typ string, cancellable *I.Cancellable, e **T.GError, v ...VArg) T.Gboolean
	pixbufSaveToStreamAsync        func(pixbuf *Pixbuf, stream *I.OutputStream, typ string, cancellable *I.Cancellable, callback I.AsyncReadyCallback, userData T.Gpointer, v ...VArg)
	pixbufSaveUtf8                 func(pixbuf *Pixbuf, filename, typ string, e **T.GError, v ...VArg) T.Gboolean
	pixbufSavevUtf8                func(pixbuf *Pixbuf, filename, typ string, optionKeys, optionValues **T.Char, e **T.GError) T.Gboolean
	pixbufScale                    func(src, dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType)
	pixbufScaleSimple              func(src *Pixbuf, destWidth, destHeight int, interpType T.GdkInterpType) *Pixbuf
	pixbufSetOption                func(pixbuf *Pixbuf, key, value string) T.Gboolean
	pixbufUnref                    func(pixbuf *Pixbuf)
)

func (p *Pixbuf) AddAlpha(substituteColor T.Gboolean, r, g, b T.Guchar) *Pixbuf {
	return pixbufAddAlpha(p, substituteColor, r, g, b)
}
func (p *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf { return pixbufApplyEmbeddedOrientation(p) }
func (p *Pixbuf) Composite(dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType, overallAlpha int) {
	pixbufComposite(p, dest, destX, destY, destWidth, destHeight, offsetX, offsetY, scaleX, scaleY, interpType, overallAlpha)
}
func (p *Pixbuf) CompositeColor(dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType, overallAlpha, checkX, checkY, checkSize int, color1, color2 T.GUint32) {
	pixbufCompositeColor(p, dest, destX, destY, destWidth, destHeight, offsetX, offsetY, scaleX, scaleY, interpType, overallAlpha, checkX, checkY, checkSize, color1, color2)
}
func (p *Pixbuf) CompositeColorSimple(destWidth, destHeight int, interpType T.GdkInterpType, overallAlpha, checkSize int, color1, color2 T.GUint32) *Pixbuf {
	return pixbufCompositeColorSimple(p, destWidth, destHeight, interpType, overallAlpha, checkSize, color1, color2)
}
func (p *Pixbuf) Copy() *Pixbuf { return pixbufCopy(p) }
func (p *Pixbuf) CopyArea(srcX, srcY, width, height int, destPixbuf *Pixbuf, destX, destY int) {
	pixbufCopyArea(p, srcX, srcY, width, height, destPixbuf, destX, destY)
}
func (p *Pixbuf) Fill(pixel T.GUint32)               { pixbufFill(p, pixel) }
func (p *Pixbuf) Flip(horizontal T.Gboolean) *Pixbuf { return pixbufFlip(p, horizontal) }
func (p *Pixbuf) GetBitsPerSample() int              { return pixbufGetBitsPerSample(p) }
func (p *Pixbuf) GetColorspace() Colorspace          { return pixbufGetColorspace(p) }
func (p *Pixbuf) GetHasAlpha() T.Gboolean            { return pixbufGetHasAlpha(p) }
func (p *Pixbuf) GetHeight() int                     { return pixbufGetHeight(p) }
func (p *Pixbuf) GetNChannels() int                  { return pixbufGetNChannels(p) }
func (p *Pixbuf) GetOption(key string) string        { return pixbufGetOption(p, key) }
func (p *Pixbuf) GetPixels() *T.Guchar               { return pixbufGetPixels(p) }
func (p *Pixbuf) GetRowstride() int                  { return pixbufGetRowstride(p) }
func (p *Pixbuf) GetWidth() int                      { return pixbufGetWidth(p) }
func (p *Pixbuf) NewSubpixbuf(srcX, srcY, width, height int) *Pixbuf {
	return pixbufNewSubpixbuf(p, srcX, srcY, width, height)
}
func (p *Pixbuf) Ref() *Pixbuf                              { return pixbufRef(p) }
func (p *Pixbuf) RotateSimple(angle PixbufRotation) *Pixbuf { return pixbufRotateSimple(p, angle) }
func (p *Pixbuf) SaturateAndPixelate(dest *Pixbuf, saturation float32, pixelate T.Gboolean) {
	pixbufSaturateAndPixelate(p, dest, saturation, pixelate)
}
func (p *Pixbuf) SaveToBuffer(buffer **T.Gchar, bufferSize *T.Gsize, typ string, e **T.GError, v ...VArg) T.Gboolean {
	return pixbufSaveToBuffer(p, buffer, bufferSize, typ, e, v)
}
func (p *Pixbuf) SaveToBufferv(buffer **T.Gchar, bufferSize *T.Gsize, typ string, optionKeys, optionValues **T.Char, e **T.GError) T.Gboolean {
	return pixbufSaveToBufferv(p, buffer, bufferSize, typ, optionKeys, optionValues, e)
}
func (p *Pixbuf) SaveToCallback(saveFunc PixbufSaveFunc, userData T.Gpointer, typ string, err **T.GError, v ...VArg) T.Gboolean {
	return pixbufSaveToCallback(p, saveFunc, userData, typ, err, v)
}
func (p *Pixbuf) SaveToCallbackv(saveFunc PixbufSaveFunc, userData T.Gpointer, typ string, optionKeys, optionValues **T.Char, e **T.GError) T.Gboolean {
	return pixbufSaveToCallbackv(p, saveFunc, userData, typ, optionKeys, optionValues, e)
}
func (p *Pixbuf) SaveToStream(stream *I.OutputStream, typ string, cancellable *I.Cancellable, e **T.GError, v ...VArg) T.Gboolean {
	return pixbufSaveToStream(p, stream, typ, cancellable, e, v)
}
func (p *Pixbuf) SaveToStreamAsync(stream *I.OutputStream, typ string, cancellable *I.Cancellable, callback I.AsyncReadyCallback, userData T.Gpointer, v ...VArg) {
	pixbufSaveToStreamAsync(p, stream, typ, cancellable, callback, userData, v)
}
func (p *Pixbuf) SaveUtf8(filename, typ string, e **T.GError, v ...VArg) T.Gboolean {
	return pixbufSaveUtf8(p, filename, typ, e, v)
}
func (p *Pixbuf) SavevUtf8(filename, typ string, optionKeys, optionValues **T.Char, e **T.GError) T.Gboolean {
	return pixbufSavevUtf8(p, filename, typ, optionKeys, optionValues, e)
}
func (p *Pixbuf) Scale(dest *Pixbuf, destX, destY, destWidth, destHeight int, offsetX, offsetY, scaleX, scaleY float64, interpType T.GdkInterpType) {
	pixbufScale(p, dest, destX, destY, destWidth, destHeight, offsetX, offsetY, scaleX, scaleY, interpType)
}
func (p *Pixbuf) ScaleSimple(destWidth, destHeight int, interpType T.GdkInterpType) *Pixbuf {
	return pixbufScaleSimple(p, destWidth, destHeight, interpType)
}
func (p *Pixbuf) SetOption(key, value string) T.Gboolean { return pixbufSetOption(p, key, value) }
func (p *Pixbuf) Unref()                                 { pixbufUnref(p) }

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
	PixbufAnimationNewFromFileUtf8 func(filename string, e **T.GError) *PixbufAnimation

	pixbufAnimationGetHeight      func(animation *PixbufAnimation) int
	pixbufAnimationGetIter        func(animation *PixbufAnimation, startTime *T.GTimeVal) *PixbufAnimationIter
	pixbufAnimationGetStaticImage func(animation *PixbufAnimation) *Pixbuf
	pixbufAnimationGetWidth       func(animation *PixbufAnimation) int
	pixbufAnimationIsStaticImage  func(animation *PixbufAnimation) T.Gboolean
	pixbufAnimationRef            func(animation *PixbufAnimation) *PixbufAnimation
	pixbufAnimationUnref          func(animation *PixbufAnimation)
)

func (p *PixbufAnimation) GetHeight() int { return pixbufAnimationGetHeight(p) }
func (p *PixbufAnimation) GetIter(startTime *T.GTimeVal) *PixbufAnimationIter {
	return pixbufAnimationGetIter(p, startTime)
}
func (p *PixbufAnimation) GetStaticImage() *Pixbuf   { return pixbufAnimationGetStaticImage(p) }
func (p *PixbufAnimation) GetWidth() int             { return pixbufAnimationGetWidth(p) }
func (p *PixbufAnimation) IsStaticImage() T.Gboolean { return pixbufAnimationIsStaticImage(p) }
func (p *PixbufAnimation) Ref() *PixbufAnimation     { return pixbufAnimationRef(p) }
func (p *PixbufAnimation) Unref()                    { pixbufAnimationUnref(p) }

type PixbufAnimationIter struct{}

var (
	PixbufAnimationIterGetType func() O.Type
	PixbufAnimationNewFromFile func(filename string, e **T.GError) *PixbufAnimation

	PixbufAnimationIterAdvance                 func(iter *PixbufAnimationIter, currentTime *T.GTimeVal) T.Gboolean
	PixbufAnimationIterGetDelayTime            func(iter *PixbufAnimationIter) int
	PixbufAnimationIterGetPixbuf               func(iter *PixbufAnimationIter) *Pixbuf
	PixbufAnimationIterOnCurrentlyLoadingFrame func(iter *PixbufAnimationIter) T.Gboolean
)

type PixbufDestroyNotify func(
	pixels *T.Guchar,
	data T.Gpointer)

type PixbufFormat struct{}

var (
	PixbufFormatGetType func() O.Type

	PixbufFormatCopy           func(format *PixbufFormat) *PixbufFormat
	PixbufFormatFree           func(format *PixbufFormat)
	PixbufFormatGetDescription func(format *PixbufFormat) string
	PixbufFormatGetExtensions  func(format *PixbufFormat) **T.Gchar
	PixbufFormatGetLicense     func(format *PixbufFormat) string
	PixbufFormatGetMimeTypes   func(format *PixbufFormat) **T.Gchar
	PixbufFormatGetName        func(format *PixbufFormat) string
	PixbufFormatIsDisabled     func(format *PixbufFormat) T.Gboolean
	PixbufFormatIsScalable     func(format *PixbufFormat) T.Gboolean
	PixbufFormatIsWritable     func(format *PixbufFormat) T.Gboolean
	PixbufFormatSetDisabled    func(format *PixbufFormat, disabled T.Gboolean)
)

type PixbufLoader struct {
	Parent O.Object
	_      T.Gpointer
}

var (
	PixbufLoaderGetType         func() O.Type
	PixbufLoaderNew             func() *PixbufLoader
	PixbufLoaderNewWithMimeType func(mimeType string, e **T.GError) *PixbufLoader
	PixbufLoaderNewWithType     func(imageType string, e **T.GError) *PixbufLoader

	PixbufLoaderClose        func(loader *PixbufLoader, e **T.GError) T.Gboolean
	PixbufLoaderGetAnimation func(loader *PixbufLoader) *PixbufAnimation
	PixbufLoaderGetFormat    func(loader *PixbufLoader) *PixbufFormat
	PixbufLoaderGetPixbuf    func(loader *PixbufLoader) *Pixbuf
	PixbufLoaderSetSize      func(loader *PixbufLoader, width, height int)
	PixbufLoaderWrite        func(loader *PixbufLoader, buf *T.Guchar, count T.Gsize, e **T.GError) T.Gboolean
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
	err **T.GError,
	data T.Gpointer) T.Gboolean

type PixbufSimpleAnim struct{}

var (
	PixbufSimpleAnimGetType func() O.Type
	PixbufSimpleAnimNew     func(width, height int, rate float32) *PixbufSimpleAnim

	pixbufSimpleAnimAddFrame func(animation *PixbufSimpleAnim, pixbuf *Pixbuf)
	pixbufSimpleAnimSetLoop  func(animation *PixbufSimpleAnim, loop T.Gboolean)
	pixbufSimpleAnimGetLoop  func(animation *PixbufSimpleAnim) T.Gboolean
)

func (p *PixbufSimpleAnim) AddFrame(pixbuf *Pixbuf) { pixbufSimpleAnimAddFrame(p, pixbuf) }
func (p *PixbufSimpleAnim) SetLoop(loop T.Gboolean) { pixbufSimpleAnimSetLoop(p, loop) }
func (p *PixbufSimpleAnim) GetLoop() T.Gboolean     { return pixbufSimpleAnimGetLoop(p) }

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
	PixbufFromPixdata func(p *Pixdata, copyPixels T.Gboolean, e **T.GError) *Pixbuf

	pixdataDeserialize func(p *Pixdata, streamLength uint, stream *uint8, e **T.GError) T.Gboolean
	pixdataFromPixbuf  func(p *Pixdata, pixbuf *Pixbuf, useRle T.Gboolean) T.Gpointer
	pixdataSerialize   func(p *Pixdata, streamLengthP *uint) *uint8
	pixdataToCsource   func(p *Pixdata, name string, dumpType PixdataDumpType) *T.GString
)

func (p *Pixdata) Deserialize(streamLength uint, stream *uint8, e **T.GError) T.Gboolean {
	return pixdataDeserialize(p, streamLength, stream, e)
}
func (p *Pixdata) FromPixbuf(pixbuf *Pixbuf, useRle T.Gboolean) T.Gpointer {
	return pixdataFromPixbuf(p, pixbuf, useRle)
}
func (p *Pixdata) Serialize(streamLengthP *uint) *uint8 { return pixdataSerialize(p, streamLengthP) }
func (p *Pixdata) ToCsource(name string, dumpType PixdataDumpType) *T.GString {
	return pixdataToCsource(p, name, dumpType)
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

	pixmapGetSize func(p *Pixmap, width, height *int)
)

func (p *Pixmap) GetSize(width, height *int) { pixmapGetSize(p, width, height) }

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
