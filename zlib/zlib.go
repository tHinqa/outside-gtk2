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

	Gzbuffer    func(g *GzFile, size uint) int
	Gzclearerr  func(g *GzFile)
	Gzclose     func(g *GzFile) int
	GzcloseR    func(g *GzFile) int
	GzcloseW    func(g *GzFile) int
	Gzdirect    func(g *GzFile) int
	Gzeof       func(g *GzFile) int
	Gzerror     func(g *GzFile, errnum *int) string
	Gzflush     func(g *GzFile, flush int) int
	Gzgetc      func(g *GzFile) int
	Gzgets      func(g *GzFile, buf string, leng int) string
	Gzoffset    func(g *GzFile) T.Long
	Gzprintf    func(g *GzFile, format string, v ...VArg) int
	Gzputc      func(g *GzFile, c int) int
	Gzputs      func(g *GzFile, s string) int
	Gzread      func(g *GzFile, buf *T.Void, leng uint) int
	Gzrewind    func(g *GzFile) int
	Gzseek      func(*GzFile, T.Long, int) T.Long
	Gzsetparams func(g *GzFile, level, strategy int) int
	Gztell      func(g *GzFile) T.Long
	Gzungetc    func(c int, file *GzFile) int
	Gzwrite     func(g *GzFile, buf *T.Void, leng uint) int
)

func (g *GzFile) Buffer(size uint) int                { return Gzbuffer(g, size) }
func (g *GzFile) Clearerr()                           { Gzclearerr(g) }
func (g *GzFile) Close() int                          { return Gzclose(g) }
func (g *GzFile) CloseR() int                         { return GzcloseR(g) }
func (g *GzFile) CloseW() int                         { return GzcloseW(g) }
func (g *GzFile) Direct() int                         { return Gzdirect(g) }
func (g *GzFile) Eof() int                            { return Gzeof(g) }
func (g *GzFile) Error(errnum *int) string            { return Gzerror(g, errnum) }
func (g *GzFile) Flush(flush int) int                 { return Gzflush(g, flush) }
func (g *GzFile) Getc() int                           { return Gzgetc(g) }
func (g *GzFile) Gets(buf string, leng int) string    { return Gzgets(g, buf, leng) }
func (g *GzFile) Offset() T.Long                      { return Gzoffset(g) }
func (g *GzFile) Printf(format string, v ...VArg) int { return Gzprintf(g, format, v) }
func (g *GzFile) Putc(c int) int                      { return Gzputc(g, c) }
func (g *GzFile) Puts(s string) int                   { return Gzputs(g, s) }
func (g *GzFile) Read(buf *T.Void, leng uint) int     { return Gzread(g, buf, leng) }
func (g *GzFile) Rewind() int                         { return Gzrewind(g) }
func (g *GzFile) Seek(a1 T.Long, a2 int) T.Long       { return Gzseek(g, a1, a2) }
func (g *GzFile) Setparams(level, strategy int) int   { return Gzsetparams(g, level, strategy) }
func (g *GzFile) Tell() T.Long                        { return Gztell(g) }
func (g *GzFile) Ungetc(c int) int                    { return Gzungetc(c, g) }
func (g *GzFile) Write(buf *T.Void, leng uint) int    { return Gzwrite(g, buf, leng) }

var (
	Deflate              func(z *ZStream, flush int) int
	DeflateBound         func(z *ZStream, sourceLen T.UnsignedLong) T.UnsignedLong
	DeflateCopy          func(dest, source *ZStream) int
	DeflateEnd           func(z *ZStream) int
	DeflateInit          func(z *ZStream, level int, version string, streamSize int) int
	DeflateInit2         func(z *ZStream, level, method, windowBits, memLevel, strategy int, version string, streamSize int) int
	DeflateParams        func(z *ZStream, level, strategy int) int
	DeflatePrime         func(z *ZStream, bits, value int) int
	DeflateReset         func(z *ZStream) int
	DeflateSetDictionary func(z *ZStream, dictionary *byte, dictLength uint) int
	DeflateSetHeader     func(z *ZStream, head *GzHeader) int
	DeflateTune          func(z *ZStream, goodLength, maxLazy, niceLength, maxChain int) int
	Inflate              func(z *ZStream, flush int) int
	InflateBack          func(z *ZStream, in InFunc, inDesc *T.Void, out OutFunc, outDesc *T.Void) int
	InflateBackEnd       func(z *ZStream) int
	InflateBackInit      func(z *ZStream, windowBits int, window *T.UnsignedChar, version string, streamSize int) int
	InflateCopy          func(dest, source *ZStream) int
	InflateEnd           func(z *ZStream) int
	InflateGetHeader     func(z *ZStream, head *GzHeader) int
	InflateInit          func(z *ZStream, version string, streamSize int) int
	InflateInit2         func(z *ZStream, windowBits int, version string, streamSize int) int
	InflateMark          func(z *ZStream) T.Long
	InflatePrime         func(z *ZStream, bits, value int) int
	InflateReset         func(z *ZStream) int
	InflateReset2        func(z *ZStream, windowBits int) int
	InflateSetDictionary func(z *ZStream, dictionary *byte, dictLength uint) int
	InflateSync          func(z *ZStream) int
	InflateSyncPoint     func(z *ZStream) int
	InflateUndermine     func(z *ZStream, i int) int
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

func (z *ZStream) Deflate(flush int) int { return Deflate(z, flush) }
func (z *ZStream) DeflateBound(sourceLen T.UnsignedLong) T.UnsignedLong {
	return DeflateBound(z, sourceLen)
}
func (dst *ZStream) DeflateCopyFrom(src *ZStream) int { return DeflateCopy(dst, src) }
func (src *ZStream) DeflateCopyTo(dst *ZStream) int   { return DeflateCopy(dst, src) }
func (z *ZStream) DeflateEnd() int                    { return DeflateEnd(z) }
func (z *ZStream) DeflateInit(level int, version string, streamSize int) int {
	return DeflateInit(z, level, version, streamSize)
}
func (z *ZStream) DeflateInit2(level, method, windowBits, memLevel, strategy int, version string, streamSize int) int {
	return DeflateInit2(z, level, method, windowBits, memLevel, strategy, version, streamSize)
}
func (z *ZStream) DeflateParams(level, strategy int) int { return DeflateParams(z, level, strategy) }
func (z *ZStream) DeflatePrime(bits, value int) int      { return DeflatePrime(z, bits, value) }
func (z *ZStream) DeflateReset() int                     { return DeflateReset(z) }
func (z *ZStream) DeflateSetDictionary(dictionary *byte, dictLength uint) int {
	return DeflateSetDictionary(z, dictionary, dictLength)
}
func (z *ZStream) DeflateSetHeader(head *GzHeader) int { return DeflateSetHeader(z, head) }
func (z *ZStream) DeflateTune(goodLength, maxLazy, niceLength, maxChain int) int {
	return DeflateTune(z, goodLength, maxLazy, niceLength, maxChain)
}
func (z *ZStream) Inflate(flush int) int { return Inflate(z, flush) }
func (z *ZStream) InflateBack(in InFunc, inDesc *T.Void, out OutFunc, outDesc *T.Void) int {
	return InflateBack(z, in, inDesc, out, outDesc)
}
func (z *ZStream) InflateBackEnd() int { return InflateBackEnd(z) }
func (z *ZStream) InflateBackInit(windowBits int, window *T.UnsignedChar, version string, streamSize int) int {
	return InflateBackInit(z, windowBits, window, version, streamSize)
}
func (dst *ZStream) InflateCopyFrom(src *ZStream) int  { return InflateCopy(dst, src) }
func (src *ZStream) InflateCopyTo(dst *ZStream) int    { return InflateCopy(dst, src) }
func (z *ZStream) InflateEnd() int                     { return InflateEnd(z) }
func (z *ZStream) InflateGetHeader(head *GzHeader) int { return InflateGetHeader(z, head) }
func (z *ZStream) InflateInit(version string, streamSize int) int {
	return InflateInit(z, version, streamSize)
}
func (z *ZStream) InflateInit2(windowBits int, version string, streamSize int) int {
	return InflateInit2(z, windowBits, version, streamSize)
}
func (z *ZStream) InflateMark() T.Long              { return InflateMark(z) }
func (z *ZStream) InflatePrime(bits, value int) int { return InflatePrime(z, bits, value) }
func (z *ZStream) InflateReset() int                { return InflateReset(z) }
func (z *ZStream) InflateReset2(windowBits int) int { return InflateReset2(z, windowBits) }
func (z *ZStream) InflateSetDictionary(dictionary *byte, dictLength uint) int {
	return InflateSetDictionary(z, dictionary, dictLength)
}
func (z *ZStream) InflateSync() int           { return InflateSync(z) }
func (z *ZStream) InflateSyncPoint() int      { return InflateSyncPoint(z) }
func (z *ZStream) InflateUndermine(i int) int { return InflateUndermine(z, i) }

var dll = "zlib1.dll"

var apiList = outside.Apis{
	{"adler32", &Adler32},
	{"adler32_combine", &Adler32Combine},
	{"compress", &Compress},
	{"compress2", &Compress2},
	{"compressBound", &CompressBound},
	{"crc32", &Crc32},
	{"crc32_combine", &Crc32Combine},
	{"deflate", &Deflate},
	{"deflateBound", &DeflateBound},
	{"deflateCopy", &DeflateCopy},
	{"deflateEnd", &DeflateEnd},
	{"deflateInit2_", &DeflateInit2},
	{"deflateInit_", &DeflateInit},
	{"deflateParams", &DeflateParams},
	{"deflatePrime", &DeflatePrime},
	{"deflateReset", &DeflateReset},
	{"deflateSetDictionary", &DeflateSetDictionary},
	{"deflateSetHeader", &DeflateSetHeader},
	{"deflateTune", &DeflateTune},
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
	{"inflate", &Inflate},
	{"inflateBack", &InflateBack},
	{"inflateBackEnd", &InflateBackEnd},
	{"inflateBackInit_", &InflateBackInit},
	{"inflateCopy", &InflateCopy},
	{"inflateEnd", &InflateEnd},
	{"inflateGetHeader", &InflateGetHeader},
	{"inflateInit2_", &InflateInit2},
	{"inflateInit_", &InflateInit},
	{"inflateMark", &InflateMark},
	{"inflatePrime", &InflatePrime},
	{"inflateReset", &InflateReset},
	{"inflateReset2", &InflateReset2},
	{"inflateSetDictionary", &InflateSetDictionary},
	{"inflateSync", &InflateSync},
	{"inflateSyncPoint", &InflateSyncPoint},
	{"inflateUndermine", &InflateUndermine},
	{"uncompress", &Uncompress},
	{"zError", &ZError},
	{"zlibCompileFlags", &ZlibCompileFlags},
	{"zlibVersion", &ZlibVersion},
}
