// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type IOChannel struct {
	RefCount        int
	Funcs           *IOFuncs
	Encoding        *T.Gchar
	ReadCd          IConv
	WriteCd         IConv
	LineTerm        *T.Gchar
	LineTermLen     uint
	BufSize         T.Gsize
	ReadBuf         *String
	EncodedReadBuf  *String
	WriteBuf        *String
	PartialWriteBuf [6]T.Gchar
	Bits            uint
	// UseBuffer : 1
	// DoEncode : 1
	// CloseOnUnref : 1
	// IsReadable : 1
	// IsWriteable : 1
	// IsSeekable : 1
	_, _ T.Gpointer
}

var (
	IoChannelErrorFromErrno   func(en int) IOChannelError
	IoChannelErrorQuark       func() Quark
	IoChannelNewFile          func(filename string, mode string, e **Error) *IOChannel
	IoChannelUnixNew          func(fd int) *IOChannel
	IoChannelWin32NewFd       func(fd int) *IOChannel
	IoChannelWin32NewMessages func(hwnd uint) *IOChannel
	IoChannelWin32NewSocket   func(socket int) *IOChannel
	IoChannelWin32Poll        func(fds *T.GPollFD, nFds int, timeout int) int

	IoAddWatch                  func(i *IOChannel, condition IOCondition, f IOFunc, userData T.Gpointer) uint
	IoAddWatchFull              func(i *IOChannel, priority int, condition IOCondition, f IOFunc, userData T.Gpointer, notify O.DestroyNotify) uint
	IoChannelClose              func(i *IOChannel)
	IoChannelFlush              func(i *IOChannel, e **Error) IOStatus
	IoChannelGetBufferCondition func(i *IOChannel) IOCondition
	IoChannelGetBuffered        func(i *IOChannel) bool
	IoChannelGetBufferSize      func(i *IOChannel) T.Gsize
	IoChannelGetCloseOnUnref    func(i *IOChannel) bool
	IoChannelGetEncoding        func(i *IOChannel) string
	IoChannelGetFlags           func(i *IOChannel) IOFlags
	IoChannelGetLineTerm        func(i *IOChannel, length *int) string
	IoChannelInit               func(i *IOChannel)
	IoChannelRead               func(i *IOChannel, buf string, count T.Gsize, bytesRead *T.Gsize) IOError
	IoChannelReadChars          func(i *IOChannel, buf string, count T.Gsize, bytesRead *T.Gsize, e **Error) IOStatus
	IoChannelReadLine           func(i *IOChannel, strReturn **T.Gchar, length, terminatorPos *T.Gsize, e **Error) IOStatus
	IoChannelReadLineString     func(i *IOChannel, buffer *String, terminatorPos *T.Gsize, e **Error) IOStatus
	IoChannelReadToEnd          func(i *IOChannel, strReturn **T.Gchar, length *T.Gsize, e **Error) IOStatus
	IoChannelReadUnichar        func(i *IOChannel, thechar *Unichar, e **Error) IOStatus
	IoChannelRef                func(i *IOChannel) *IOChannel
	IoChannelSeek               func(i *IOChannel, offset int64, typ SeekType) IOError
	IoChannelSeekPosition       func(i *IOChannel, offset int64, typ SeekType, e **Error) IOStatus
	IoChannelSetBuffered        func(i *IOChannel, buffered bool)
	IoChannelSetBufferSize      func(i *IOChannel, size T.Gsize)
	IoChannelSetCloseOnUnref    func(i *IOChannel, doClose bool)
	IoChannelSetEncoding        func(i *IOChannel, encoding string, e **Error) IOStatus
	IoChannelSetFlags           func(i *IOChannel, flags IOFlags, e **Error) IOStatus
	IoChannelSetLineTerm        func(i *IOChannel, lineTerm string, length int)
	IoChannelShutdown           func(i *IOChannel, flush bool, err **Error) IOStatus
	IoChannelUnixGetFd          func(i *IOChannel) int
	IoChannelUnref              func(i *IOChannel)
	IoChannelWin32GetFd         func(i *IOChannel) int
	IoChannelWin32MakePollfd    func(i *IOChannel, condition IOCondition, fd *T.GPollFD)
	IoChannelWrite              func(i *IOChannel, buf string, count T.Gsize, bytesWritten *T.Gsize) IOError
	IoChannelWriteChars         func(i *IOChannel, buf string, count T.Gssize, bytesWritten *T.Gsize, e **Error) IOStatus
	IoChannelWriteUnichar       func(i *IOChannel, thechar Unichar, e **Error) IOStatus
	IoCreateWatch               func(i *IOChannel, condition IOCondition) *O.Source
)

func (i *IOChannel) AddWatch(condition IOCondition, f IOFunc, userData T.Gpointer) uint {
	return IoAddWatch(i, condition, f, userData)
}
func (i *IOChannel) AddWatchFull(priority int, condition IOCondition, f IOFunc, userData T.Gpointer, notify O.DestroyNotify) uint {
	return IoAddWatchFull(i, priority, condition, f, userData, notify)
}
func (i *IOChannel) Close()                                      { IoChannelClose(i) }
func (i *IOChannel) CreateWatch(condition IOCondition) *O.Source { return IoCreateWatch(i, condition) }
func (i *IOChannel) Flush(e **Error) IOStatus                    { return IoChannelFlush(i, e) }
func (i *IOChannel) GetBufferCondition() IOCondition             { return IoChannelGetBufferCondition(i) }
func (i *IOChannel) GetBuffered() bool                           { return IoChannelGetBuffered(i) }
func (i *IOChannel) GetBufferSize() T.Gsize                      { return IoChannelGetBufferSize(i) }
func (i *IOChannel) GetCloseOnUnref() bool                       { return IoChannelGetCloseOnUnref(i) }
func (i *IOChannel) GetEncoding() string                         { return IoChannelGetEncoding(i) }
func (i *IOChannel) GetFlags() IOFlags                           { return IoChannelGetFlags(i) }
func (i *IOChannel) GetLineTerm(length *int) string              { return IoChannelGetLineTerm(i, length) }
func (i *IOChannel) Init()                                       { IoChannelInit(i) }
func (i *IOChannel) Read(buf string, count T.Gsize, bytesRead *T.Gsize) IOError {
	return IoChannelRead(i, buf, count, bytesRead)
}
func (i *IOChannel) ReadChars(buf string, count T.Gsize, bytesRead *T.Gsize, e **Error) IOStatus {
	return IoChannelReadChars(i, buf, count, bytesRead, e)
}
func (i *IOChannel) ReadLine(strReturn **T.Gchar, length, terminatorPos *T.Gsize, e **Error) IOStatus {
	return IoChannelReadLine(i, strReturn, length, terminatorPos, e)
}
func (i *IOChannel) ReadLineString(buffer *String, terminatorPos *T.Gsize, e **Error) IOStatus {
	return IoChannelReadLineString(i, buffer, terminatorPos, e)
}
func (i *IOChannel) ReadToEnd(strReturn **T.Gchar, length *T.Gsize, e **Error) IOStatus {
	return IoChannelReadToEnd(i, strReturn, length, e)
}
func (i *IOChannel) ReadUnichar(thechar *Unichar, e **Error) IOStatus {
	return IoChannelReadUnichar(i, thechar, e)
}
func (i *IOChannel) Ref() *IOChannel                         { return IoChannelRef(i) }
func (i *IOChannel) Seek(offset int64, typ SeekType) IOError { return IoChannelSeek(i, offset, typ) }
func (i *IOChannel) SeekPosition(offset int64, typ SeekType, e **Error) IOStatus {
	return IoChannelSeekPosition(i, offset, typ, e)
}
func (i *IOChannel) SetBuffered(buffered bool)    { IoChannelSetBuffered(i, buffered) }
func (i *IOChannel) SetBufferSize(size T.Gsize)   { IoChannelSetBufferSize(i, size) }
func (i *IOChannel) SetCloseOnUnref(doClose bool) { IoChannelSetCloseOnUnref(i, doClose) }
func (i *IOChannel) SetEncoding(encoding string, e **Error) IOStatus {
	return IoChannelSetEncoding(i, encoding, e)
}
func (i *IOChannel) SetFlags(flags IOFlags, e **Error) IOStatus {
	return IoChannelSetFlags(i, flags, e)
}
func (i *IOChannel) SetLineTerm(lineTerm string, length int) {
	IoChannelSetLineTerm(i, lineTerm, length)
}
func (i *IOChannel) Shutdown(flush bool, err **Error) IOStatus {
	return IoChannelShutdown(i, flush, err)
}
func (i *IOChannel) UnixGetFd() int  { return IoChannelUnixGetFd(i) }
func (i *IOChannel) Unref()          { IoChannelUnref(i) }
func (i *IOChannel) Win32GetFd() int { return IoChannelWin32GetFd(i) }
func (i *IOChannel) Win32MakePollfd(condition IOCondition, fd *T.GPollFD) {
	IoChannelWin32MakePollfd(i, condition, fd)
}
func (i *IOChannel) Write(buf string, count T.Gsize, bytesWritten *T.Gsize) IOError {
	return IoChannelWrite(i, buf, count, bytesWritten)
}
func (i *IOChannel) WriteChars(buf string, count T.Gssize, bytesWritten *T.Gsize, e **Error) IOStatus {
	return IoChannelWriteChars(i, buf, count, bytesWritten, e)
}
func (i *IOChannel) WriteUnichar(thechar Unichar, e **Error) IOStatus {
	return IoChannelWriteUnichar(i, thechar, e)
}

type IOChannelError Enum

const (
	IO_CHANNEL_ERROR_FBIG IOChannelError = iota
	IO_CHANNEL_ERROR_INVAL
	IO_CHANNEL_ERROR_IO
	IO_CHANNEL_ERROR_ISDIR
	IO_CHANNEL_ERROR_NOSPC
	IO_CHANNEL_ERROR_NXIO
	IO_CHANNEL_ERROR_OVERFLOW
	IO_CHANNEL_ERROR_PIPE
	IO_CHANNEL_ERROR_FAILED
)

type IOCondition Enum

const (
	IO_IN IOCondition = 1 << iota
	IO_PRI
	IO_OUT
	IO_ERR
	IO_HUP
	IO_NVAL
)

type IConv struct{}

var (
	IconvOpen func(toCodeset string, fromCodeset string) IConv

	Iconv      func(i IConv, inbuf **T.Gchar, inbytesLeft *T.Gsize, outbuf **T.Gchar, outbytesLeft *T.Gsize) T.Gsize
	IconvClose func(i IConv) int
)

type IOError Enum

const (
	IO_ERROR_NONE IOError = iota
	IO_ERROR_AGAIN
	IO_ERROR_INVAL
	IO_ERROR_UNKNOWN
)

type IOFlags Enum

const (
	IO_FLAG_APPEND IOFlags = 1 << iota
	IO_FLAG_NONBLOCK
	IO_FLAG_IS_READABLE
	IO_FLAG_IS_WRITEABLE
	IO_FLAG_IS_SEEKABLE
	IO_FLAG_MASK     IOFlags = (1 << iota) - 1
	IO_FLAG_GET_MASK         = IO_FLAG_MASK
	IO_FLAG_SET_MASK         = IO_FLAG_APPEND | IO_FLAG_NONBLOCK
)

type IOFunc func(source *IOChannel,
	condition IOCondition, data T.Gpointer) T.Gboolean

type IOFuncs struct {
	IoRead func(channel *IOChannel, buf *T.Gchar,
		count T.Gsize, bytesRead *T.Gsize,
		err **Error) IOStatus
	IoWrite func(channel *IOChannel, buf *T.Gchar,
		count T.Gsize, bytesWritten *T.Gsize,
		err **Error) IOStatus
	IoSeek func(channel *IOChannel, offset int64,
		typ SeekType, err **Error) IOStatus
	IoClose func(channel *IOChannel,
		err **Error) IOStatus
	IoCreateWatch func(channel *IOChannel,
		condition IOCondition) *Source
	IoFree     func(channel *IOChannel)
	IoSetFlags func(channel *IOChannel,
		flags IOFlags, err **Error) IOStatus
	IoGetFlags func(channel *IOChannel) IOFlags
}

type IOStatus Enum

const (
	IO_STATUS_ERROR IOStatus = iota
	IO_STATUS_NORMAL
	IO_STATUS_EOF
	IO_STATUS_AGAIN
)
