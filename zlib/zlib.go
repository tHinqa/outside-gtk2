// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package zlib provides API definitions for accessing
//zlib1.dll.
package zlib

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type (
	GzFile        *T.Void
	InternalState struct{}

	AllocFunc func(
		opaque *T.Void,
		items uint,
		size uint) *T.Void

	FreeFunc func(
		opaque *T.Void,
		address *T.Void)

	InFunc func(
		*T.Void,
		**T.UnsignedChar) uint

	OutFunc func(
		*T.Void,
		*T.UnsignedChar,
		uint) int

	ZStream struct {
		NextIn   *byte
		AvailIn  uint
		TotalIn  T.UnsignedLong
		NextOut  *byte
		AvailOut uint
		TotalOut T.UnsignedLong
		Msg      *T.Char
		State    *InternalState
		Zalloc   AllocFunc
		Zfree    FreeFunc
		Opaque   *T.Void
		DataType int
		Adler    T.UnsignedLong
		_        T.UnsignedLong
	}

	GzHeader struct {
		Text     int
		Time     T.UnsignedLong
		Xflags   int
		Os       int
		Extra    *byte
		ExtraLen uint
		ExtraMax uint
		Name     *byte
		NameMax  uint
		Comment  *byte
		CommMax  uint
		Hcrc     int
		Done     int
	}
)

var (
	ZlibVersion func() string

	ZlibCompileFlags func() T.UnsignedLong

	Compress func(
		dst *byte, destLen *T.UnsignedLong,
		src *byte, sourceLen T.UnsignedLong) int

	Compress2 func(
		dst *byte, destLen *T.UnsignedLong,
		src *byte, sourceLen T.UnsignedLong, level int) int

	CompressBound func(
		sourceLen T.UnsignedLong) T.UnsignedLong

	Uncompress func(
		dst *byte, dstLen *T.UnsignedLong,
		src *byte, srcLen T.UnsignedLong) int

	Gzdopen func(fd int, mode string) GzFile

	Gzbuffer func(file GzFile, size uint) int

	Gzsetparams func(file GzFile, level, strategy int) int

	Gzread func(file GzFile, buf *T.Void, len uint) int

	Gzwrite func(file GzFile, buf *T.Void, len uint) int

	Gzprintf func(file GzFile, format string, v ...VArg) int

	Gzputs func(file GzFile, s string) int

	Gzgets func(file GzFile, buf string, len int) string

	Gzputc func(file GzFile, c int) int

	Gzgetc func(file GzFile) int

	Gzungetc func(c int, file GzFile) int

	Gzflush func(file GzFile, flush int) int

	Gzrewind func(file GzFile) int

	Gzeof func(file GzFile) int

	Gzdirect func(file GzFile) int

	Gzclose func(file GzFile) int

	GzcloseR func(file GzFile) int

	GzcloseW func(file GzFile) int

	Gzerror func(file GzFile, errnum *int) string

	Gzclearerr func(file GzFile)

	Adler32 func(
		adler T.UnsignedLong, buf *byte, len uint) T.UnsignedLong

	Crc32 func(
		crc T.UnsignedLong, buf *byte, len uint) T.UnsignedLong

	Gzopen func(string, string) GzFile

	Gzseek func(GzFile, T.Long, int) T.Long

	Gztell func(GzFile) T.Long

	Gzoffset func(GzFile) T.Long

	Adler32Combine func(T.UnsignedLong, T.UnsignedLong, T.Long) T.UnsignedLong

	Crc32Combine func(T.UnsignedLong, T.UnsignedLong, T.Long) T.UnsignedLong

	ZError func(int) string

	GetCrcTable func() *T.UnsignedLong
)

var (
	deflate              func(z *ZStream, flush int) int
	deflateBound         func(z *ZStream, sourceLen T.UnsignedLong) T.UnsignedLong
	deflateCopy          func(dest, source *ZStream) int
	deflateEnd           func(z *ZStream) int
	deflateInit          func(z *ZStream, level int, version string, streamSize int) int
	deflateInit2         func(z *ZStream, level, method, windowBits, memLevel, strategy int, version string, streamSize int) int
	deflateParams        func(z *ZStream, level, strategy int) int
	deflatePrime         func(z *ZStream, bits, value int) int
	deflateReset         func(z *ZStream) int
	deflateSetDictionary func(z *ZStream, dictionary *byte, dictLength uint) int
	deflateSetHeader     func(z *ZStream, head *GzHeader) int
	deflateTune          func(z *ZStream, goodLength, maxLazy, niceLength, maxChain int) int
	inflate              func(z *ZStream, flush int) int
	inflateBack          func(z *ZStream, in InFunc, inDesc *T.Void, out OutFunc, outDesc *T.Void) int
	inflateBackEnd       func(z *ZStream) int
	inflateBackInit      func(z *ZStream, windowBits int, window *T.UnsignedChar, version string, streamSize int) int
	inflateCopy          func(dest, source *ZStream) int
	inflateEnd           func(z *ZStream) int
	inflateGetHeader     func(z *ZStream, head *GzHeader) int
	inflateInit          func(z *ZStream, version string, streamSize int) int
	inflateInit2         func(z *ZStream, windowBits int, version string, streamSize int) int
	inflateMark          func(z *ZStream) T.Long
	inflatePrime         func(z *ZStream, bits, value int) int
	inflateReset         func(z *ZStream) int
	inflateReset2        func(z *ZStream, windowBits int) int
	inflateSetDictionary func(z *ZStream, dictionary *byte, dictLength uint) int
	inflateSync          func(z *ZStream) int
	inflateSyncPoint     func(z *ZStream) int
	inflateUndermine     func(z *ZStream, i int) int
)

func (z *ZStream) Deflate(flush int) int { return deflate(z, flush) }
func (z *ZStream) DeflateBound(sourceLen T.UnsignedLong) T.UnsignedLong {
	return deflateBound(z, sourceLen)
}
func (dst *ZStream) DeflateCopyFrom(src *ZStream) int { return deflateCopy(dst, src) }
func (src *ZStream) DeflateCopyTo(dst *ZStream) int   { return deflateCopy(dst, src) }
func (z *ZStream) DeflateEnd() int                    { return deflateEnd(z) }
func (z *ZStream) DeflateInit(level int, version string, streamSize int) int {
	return deflateInit(z, level, version, streamSize)
}
func (z *ZStream) DeflateInit2(level, method, windowBits, memLevel, strategy int, version string, streamSize int) int {
	return deflateInit2(z, level, method, windowBits, memLevel, strategy, version, streamSize)
}
func (z *ZStream) DeflateParams(level, strategy int) int { return deflateParams(z, level, strategy) }
func (z *ZStream) DeflatePrime(bits, value int) int      { return deflatePrime(z, bits, value) }
func (z *ZStream) DeflateReset() int                     { return deflateReset(z) }
func (z *ZStream) DeflateSetDictionary(dictionary *byte, dictLength uint) int {
	return deflateSetDictionary(z, dictionary, dictLength)
}
func (z *ZStream) DeflateSetHeader(head *GzHeader) int { return deflateSetHeader(z, head) }
func (z *ZStream) DeflateTune(goodLength, maxLazy, niceLength, maxChain int) int {
	return deflateTune(z, goodLength, maxLazy, niceLength, maxChain)
}
func (z *ZStream) Inflate(flush int) int { return inflate(z, flush) }
func (z *ZStream) InflateBack(in InFunc, inDesc *T.Void, out OutFunc, outDesc *T.Void) int {
	return inflateBack(z, in, inDesc, out, outDesc)
}
func (z *ZStream) InflateBackEnd() int { return inflateBackEnd(z) }
func (z *ZStream) InflateBackInit(windowBits int, window *T.UnsignedChar, version string, streamSize int) int {
	return inflateBackInit(z, windowBits, window, version, streamSize)
}
func (dst *ZStream) InflateCopyFrom(src *ZStream) int  { return inflateCopy(dst, src) }
func (src *ZStream) InflateCopyTo(dst *ZStream) int    { return inflateCopy(dst, src) }
func (z *ZStream) InflateEnd() int                     { return inflateEnd(z) }
func (z *ZStream) InflateGetHeader(head *GzHeader) int { return inflateGetHeader(z, head) }
func (z *ZStream) InflateInit(version string, streamSize int) int {
	return inflateInit(z, version, streamSize)
}
func (z *ZStream) InflateInit2(windowBits int, version string, streamSize int) int {
	return inflateInit2(z, windowBits, version, streamSize)
}
func (z *ZStream) InflateMark() T.Long              { return inflateMark(z) }
func (z *ZStream) InflatePrime(bits, value int) int { return inflatePrime(z, bits, value) }
func (z *ZStream) InflateReset() int                { return inflateReset(z) }
func (z *ZStream) InflateReset2(windowBits int) int { return inflateReset2(z, windowBits) }
func (z *ZStream) InflateSetDictionary(dictionary *byte, dictLength uint) int {
	return inflateSetDictionary(z, dictionary, dictLength)
}
func (z *ZStream) InflateSync() int           { return inflateSync(z) }
func (z *ZStream) InflateSyncPoint() int      { return inflateSyncPoint(z) }
func (z *ZStream) InflateUndermine(i int) int { return inflateUndermine(z, i) }

var dll = "zlib1.dll"

var apiList = outside.Apis{
	{"adler32", &Adler32},
	{"adler32_combine", &Adler32Combine},
	{"compress", &Compress},
	{"compress2", &Compress2},
	{"compressBound", &CompressBound},
	{"crc32", &Crc32},
	{"crc32_combine", &Crc32Combine},
	{"deflate", &deflate},
	{"deflateBound", &deflateBound},
	{"deflateCopy", &deflateCopy},
	{"deflateEnd", &deflateEnd},
	{"deflateInit2_", &deflateInit2},
	{"deflateInit_", &deflateInit},
	{"deflateParams", &deflateParams},
	{"deflatePrime", &deflatePrime},
	{"deflateReset", &deflateReset},
	{"deflateSetDictionary", &deflateSetDictionary},
	{"deflateSetHeader", &deflateSetHeader},
	{"deflateTune", &deflateTune},
	{"get_crc_table", &GetCrcTable},
	{"gzbuffer", &Gzbuffer},
	{"gzclearerr", &Gzclearerr},
	{"gzclose", &Gzclose},
	{"gzclose_r", &GzcloseR},
	{"gzclose_w", &GzcloseW},
	{"gzdirect", &Gzdirect},
	{"gzdopen", &Gzdopen},
	{"gzeof", &Gzeof},
	{"gzerror", &Gzerror},
	{"gzflush", &Gzflush},
	{"gzgetc", &Gzgetc},
	{"gzgets", &Gzgets},
	{"gzoffset", &Gzoffset},
	{"gzopen", &Gzopen},
	{"gzprintf", &Gzprintf},
	{"gzputc", &Gzputc},
	{"gzputs", &Gzputs},
	{"gzread", &Gzread},
	{"gzrewind", &Gzrewind},
	{"gzseek", &Gzseek},
	{"gzsetparams", &Gzsetparams},
	{"gztell", &Gztell},
	{"gzungetc", &Gzungetc},
	{"gzwrite", &Gzwrite},
	{"inflate", &inflate},
	{"inflateBack", &inflateBack},
	{"inflateBackEnd", &inflateBackEnd},
	{"inflateBackInit_", &inflateBackInit},
	{"inflateCopy", &inflateCopy},
	{"inflateEnd", &inflateEnd},
	{"inflateGetHeader", &inflateGetHeader},
	{"inflateInit2_", &inflateInit2},
	{"inflateInit_", &inflateInit},
	{"inflateMark", &inflateMark},
	{"inflatePrime", &inflatePrime},
	{"inflateReset", &inflateReset},
	{"inflateReset2", &inflateReset2},
	{"inflateSetDictionary", &inflateSetDictionary},
	{"inflateSync", &inflateSync},
	{"inflateSyncPoint", &inflateSyncPoint},
	{"inflateUndermine", &inflateUndermine},
	{"uncompress", &Uncompress},
	{"zError", &ZError},
	{"zlibCompileFlags", &ZlibCompileFlags},
	{"zlibVersion", &ZlibVersion},
}
