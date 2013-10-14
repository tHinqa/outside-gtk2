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
	Byte           T.Unsigned_char
	Bytef          Byte
	GzFile         Voidp
	Internal_state struct{}
	UInt           T.Unsigned_int
	ULong          T.Unsigned_long
	ULongf         ULong
	Voidp          *T.Void
	Voidpf         *T.Void
	Z_streamp      *Z_stream
	Gz_headerp     *Gz_header

	Alloc_func func(
		opaque Voidpf,
		items UInt,
		size UInt) Voidpf

	Free_func func(
		opaque Voidpf,
		address Voidpf)

	In_func func(
		*T.Void,
		**T.Unsigned_char) uint

	Out_func func(
		*T.Void,
		*T.Unsigned_char,
		uint) int

	Z_stream struct {
		Next_in   *Bytef
		Avail_in  UInt
		Total_in  ULong
		Next_out  *Bytef
		Avail_out UInt
		Total_out ULong
		Msg       *T.Char
		State     *Internal_state
		Zalloc    Alloc_func
		Zfree     Free_func
		Opaque    Voidpf
		Data_type int
		Adler     ULong
		_         ULong
	}

	Gz_header struct {
		Text      int
		Time      ULong
		Xflags    int
		Os        int
		Extra     *Bytef
		Extra_len UInt
		Extra_max UInt
		Name      *Bytef
		Name_max  UInt
		Comment   *Bytef
		Comm_max  UInt
		Hcrc      int
		Done      int
	}
)

var (
	ZlibVersion func() string

	Deflate func(strm Z_streamp, flush int) int

	DeflateEnd func(strm Z_streamp) int

	Inflate func(strm Z_streamp, flush int) int

	InflateEnd func(strm Z_streamp) int

	DeflateSetDictionary func(
		strm Z_streamp,
		dictionary *Bytef,
		dictLength UInt) int

	DeflateCopy func(dest, source Z_streamp) int

	DeflateReset func(strm Z_streamp) int

	DeflateParams func(strm Z_streamp, level, strategy int) int

	DeflateTune func(strm Z_streamp,
		good_length, max_lazy, nice_length, max_chain int) int

	DeflateBound func(strm Z_streamp, sourceLen ULong) ULong

	DeflatePrime func(strm Z_streamp, bits, value int) int

	DeflateSetHeader func(strm Z_streamp, head Gz_headerp) int

	InflateSetDictionary func(
		strm Z_streamp, dictionary *Bytef, dictLength UInt) int

	InflateSync func(strm Z_streamp) int

	InflateCopy func(dest, source Z_streamp) int

	InflateReset func(strm Z_streamp) int

	InflateReset2 func(strm Z_streamp, windowBits int) int

	InflatePrime func(strm Z_streamp, bits, value int) int

	InflateMark func(strm Z_streamp) T.Long

	InflateGetHeader func(strm Z_streamp, head Gz_headerp) int

	InflateBack func(
		strm Z_streamp,
		in In_func,
		in_desc *T.Void,
		out Out_func,
		out_desc *T.Void) int

	InflateBackEnd func(strm Z_streamp) int

	ZlibCompileFlags func() ULong

	Compress func(
		dest *Bytef,
		destLen *ULongf,
		source *Bytef,
		sourceLen ULong) int

	Compress2 func(
		dest *Bytef,
		destLen *ULongf,
		source *Bytef,
		sourceLen ULong,
		level int) int

	CompressBound func(
		sourceLen ULong) ULong

	Uncompress func(
		dest *Bytef,
		destLen *ULongf,
		source *Bytef,
		sourceLen ULong) int

	Gzdopen func(
		fd int,
		mode string) GzFile

	Gzbuffer func(
		file GzFile,
		size uint) int

	Gzsetparams func(
		file GzFile,
		level int,
		strategy int) int

	Gzread func(
		file GzFile,
		buf Voidp,
		len uint) int

	Gzwrite func(
		file GzFile,
		buf Voidp,
		len uint) int

	Gzprintf func(
		file GzFile,
		format string,
		v ...VArg) int

	Gzputs func(
		file GzFile,
		s string) int

	Gzgets func(
		file GzFile,
		buf string,
		len int) string

	Gzputc func(
		file GzFile,
		c int) int

	Gzgetc func(
		file GzFile) int

	Gzungetc func(
		c int,
		file GzFile) int

	Gzflush func(
		file GzFile,
		flush int) int

	Gzrewind func(
		file GzFile) int

	Gzeof func(
		file GzFile) int

	Gzdirect func(
		file GzFile) int

	Gzclose func(
		file GzFile) int

	Gzclose_r func(
		file GzFile) int

	Gzclose_w func(
		file GzFile) int

	Gzerror func(
		file GzFile,
		errnum *int) string

	Gzclearerr func(
		file GzFile)

	Adler32 func(
		adler ULong,
		buf *Bytef,
		len UInt) ULong

	Crc32 func(
		crc ULong,
		buf *Bytef,
		len UInt) ULong

	DeflateInit_ func(
		strm Z_streamp,
		level int,
		version string,
		stream_size int) int

	InflateInit_ func(
		strm Z_streamp,
		version string,
		stream_size int) int

	DeflateInit2_ func(
		strm Z_streamp,
		level int,
		method int,
		windowBits int,
		memLevel int,
		strategy int,
		version string,
		stream_size int) int

	InflateInit2_ func(
		strm Z_streamp,
		windowBits int,
		version string,
		stream_size int) int

	InflateBackInit_ func(
		strm Z_streamp,
		windowBits int,
		window *T.Unsigned_char,
		version string,
		stream_size int) int

	Gzopen func(string, string) GzFile

	Gzseek func(GzFile, T.Long, int) T.Long

	Gztell func(GzFile) T.Long

	Gzoffset func(GzFile) T.Long

	Adler32_combine func(ULong, ULong, T.Long) ULong

	Crc32_combine func(ULong, ULong, T.Long) ULong

	ZError func(int) string

	InflateSyncPoint func(Z_streamp) int

	Get_crc_table func() *ULongf

	InflateUndermine func(Z_streamp, int) int
)

var dll = "zlib1.dll"

var apiList = outside.Apis{
	{"adler32", &Adler32},
	{"adler32_combine", &Adler32_combine},
	{"compress", &Compress},
	{"compress2", &Compress2},
	{"compressBound", &CompressBound},
	{"crc32", &Crc32},
	{"crc32_combine", &Crc32_combine},
	{"deflate", &Deflate},
	{"deflateBound", &DeflateBound},
	{"deflateCopy", &DeflateCopy},
	{"deflateEnd", &DeflateEnd},
	{"deflateInit2_", &DeflateInit2_},
	{"deflateInit_", &DeflateInit_},
	{"deflateParams", &DeflateParams},
	{"deflatePrime", &DeflatePrime},
	{"deflateReset", &DeflateReset},
	{"deflateSetDictionary", &DeflateSetDictionary},
	{"deflateSetHeader", &DeflateSetHeader},
	{"deflateTune", &DeflateTune},
	{"get_crc_table", &Get_crc_table},
	{"gzbuffer", &Gzbuffer},
	{"gzclearerr", &Gzclearerr},
	{"gzclose", &Gzclose},
	{"gzclose_r", &Gzclose_r},
	{"gzclose_w", &Gzclose_w},
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
	{"inflateBackInit_", &InflateBackInit_},
	{"inflateCopy", &InflateCopy},
	{"inflateEnd", &InflateEnd},
	{"inflateGetHeader", &InflateGetHeader},
	{"inflateInit2_", &InflateInit2_},
	{"inflateInit_", &InflateInit_},
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
