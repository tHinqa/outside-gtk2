// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
)

type ZlibCompressor struct{}

var (
	ZlibCompressorGetType func() O.Type
	ZlibCompressorNew     func(format ZlibCompressorFormat, level int) *ZlibCompressor

	ZlibCompressorGetFileInfo func(compressor *ZlibCompressor) *FileInfo
	ZlibCompressorSetFileInfo func(compressor *ZlibCompressor, fileInfo *FileInfo)
)

func (z *ZlibCompressor) GetFileInfo() *FileInfo         { return ZlibCompressorGetFileInfo(z) }
func (z *ZlibCompressor) SetFileInfo(fileInfo *FileInfo) { ZlibCompressorSetFileInfo(z, fileInfo) }

type ZlibDecompressor struct{}

var (
	ZlibDecompressorGetType func() O.Type
	ZlibDecompressorNew     func(format ZlibCompressorFormat) *ZlibDecompressor

	ZlibDecompressorGetFileInfo func(decompressor *ZlibDecompressor) *FileInfo
)

func (z *ZlibDecompressor) GetFileInfo() *FileInfo { return ZlibDecompressorGetFileInfo(z) }

type ZlibCompressorFormat Enum

const (
	ZLIB_COMPRESSOR_FORMAT_ZLIB ZlibCompressorFormat = iota
	ZLIB_COMPRESSOR_FORMAT_GZIP
	ZLIB_COMPRESSOR_FORMAT_RAW
)

var ZlibCompressorFormatGetType func() O.Type
