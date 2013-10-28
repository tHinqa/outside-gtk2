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
	GzFile struct{}

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

	ZError func(int) string

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

	Adler32 func(
		adler T.UnsignedLong, buf *byte, len uint) T.UnsignedLong

	Crc32 func(
		crc T.UnsignedLong, buf *byte, len uint) T.UnsignedLong

	Adler32Combine func(T.UnsignedLong, T.UnsignedLong, T.Long) T.UnsignedLong

	Crc32Combine func(T.UnsignedLong, T.UnsignedLong, T.Long) T.UnsignedLong

	GetCrcTable func() *T.UnsignedLong
)
var (
	Gzdopen func(fd int, mode string) *GzFile
	Gzopen  func(string, string) *GzFile

	gzbuffer    func(g *GzFile, size uint) int
	gzclearerr  func(g *GzFile)
	gzclose     func(g *GzFile) int
	gzcloseR    func(g *GzFile) int
	gzcloseW    func(g *GzFile) int
	gzdirect    func(g *GzFile) int
	gzeof       func(g *GzFile) int
	gzerror     func(g *GzFile, errnum *int) string
	gzflush     func(g *GzFile, flush int) int
	gzgetc      func(g *GzFile) int
	gzgets      func(g *GzFile, buf string, leng int) string
	gzoffset    func(g *GzFile) T.Long
	gzprintf    func(g *GzFile, format string, v ...VArg) int
	gzputc      func(g *GzFile, c int) int
	gzputs      func(g *GzFile, s string) int
	gzread      func(g *GzFile, buf *T.Void, leng uint) int
	gzrewind    func(g *GzFile) int
	gzseek      func(*GzFile, T.Long, int) T.Long
	gzsetparams func(g *GzFile, level, strategy int) int
	gztell      func(g *GzFile) T.Long
	gzungetc    func(c int, file *GzFile) int
	gzwrite     func(g *GzFile, buf *T.Void, leng uint) int
)

func (g *GzFile) Buffer(size uint) int                { return gzbuffer(g, size) }
func (g *GzFile) Clearerr()                           { gzclearerr(g) }
func (g *GzFile) Close() int                          { return gzclose(g) }
func (g *GzFile) CloseR() int                         { return gzcloseR(g) }
func (g *GzFile) CloseW() int                         { return gzcloseW(g) }
func (g *GzFile) Direct() int                         { return gzdirect(g) }
func (g *GzFile) Eof() int                            { return gzeof(g) }
func (g *GzFile) Error(errnum *int) string            { return gzerror(g, errnum) }
func (g *GzFile) Flush(flush int) int                 { return gzflush(g, flush) }
func (g *GzFile) Getc() int                           { return gzgetc(g) }
func (g *GzFile) Gets(buf string, leng int) string    { return gzgets(g, buf, leng) }
func (g *GzFile) Offset() T.Long                      { return gzoffset(g) }
func (g *GzFile) Printf(format string, v ...VArg) int { return gzprintf(g, format, v) }
func (g *GzFile) Putc(c int) int                      { return gzputc(g, c) }
func (g *GzFile) Puts(s string) int                   { return gzputs(g, s) }
func (g *GzFile) Read(buf *T.Void, leng uint) int     { return gzread(g, buf, leng) }
func (g *GzFile) Rewind() int                         { return gzrewind(g) }
func (g *GzFile) Seek(a1 T.Long, a2 int) T.Long       { return gzseek(g, a1, a2) }
func (g *GzFile) Setparams(level, strategy int) int   { return gzsetparams(g, level, strategy) }
func (g *GzFile) Tell() T.Long                        { return gztell(g) }
func (g *GzFile) Ungetc(c int) int                    { return gzungetc(c, g) }
func (g *GzFile) Write(buf *T.Void, leng uint) int    { return gzwrite(g, buf, leng) }

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

type ZStream struct {
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
	{"gzbuffer", &gzbuffer},
	{"gzclearerr", &gzclearerr},
	{"gzclose", &gzclose},
	{"gzclose_r", &gzcloseR},
	{"gzclose_w", &gzcloseW},
	{"gzdirect", &gzdirect},
	{"gzdopen", &Gzdopen},
	{"gzeof", &gzeof},
	{"gzerror", &gzerror},
	{"gzflush", &gzflush},
	{"gzgetc", &gzgetc},
	{"gzgets", &gzgets},
	{"gzoffset", &gzoffset},
	{"gzopen", &Gzopen},
	{"gzprintf", &gzprintf},
	{"gzputc", &gzputc},
	{"gzputs", &gzputs},
	{"gzread", &gzread},
	{"gzrewind", &gzrewind},
	{"gzseek", &gzseek},
	{"gzsetparams", &gzsetparams},
	{"gztell", &gztell},
	{"gzungetc", &gzungetc},
	{"gzwrite", &gzwrite},
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
