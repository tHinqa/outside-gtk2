// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type OutputStream struct {
	Parent O.Object
	_      *struct{}
}

var (
	OutputStreamGetType func() O.Type

	OutputStreamClearPending func(o *OutputStream)
	OutputStreamClose        func(o *OutputStream, cancellable *Cancellable, err **L.Error) bool
	OutputStreamCloseAsync   func(o *OutputStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	OutputStreamCloseFinish  func(o *OutputStream, result *AsyncResult, err **L.Error) bool
	OutputStreamFlush        func(o *OutputStream, cancellable *Cancellable, err **L.Error) bool
	OutputStreamFlushAsync   func(o *OutputStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	OutputStreamFlushFinish  func(o *OutputStream, result *AsyncResult, err **L.Error) bool
	OutputStreamHasPending   func(o *OutputStream) bool
	OutputStreamIsClosed     func(o *OutputStream) bool
	OutputStreamIsClosing    func(o *OutputStream) bool
	OutputStreamSetPending   func(o *OutputStream, err **L.Error) bool
	OutputStreamSplice       func(o *OutputStream, source *InputStream, flags OutputStreamSpliceFlags, cancellable *Cancellable, err **L.Error) T.Gssize
	OutputStreamSpliceAsync  func(o *OutputStream, source *InputStream, flags OutputStreamSpliceFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	OutputStreamSpliceFinish func(o *OutputStream, result *AsyncResult, err **L.Error) T.Gssize
	OutputStreamWrite        func(o *OutputStream, buffer *T.Void, count T.Gsize, cancellable *Cancellable, err **L.Error) T.Gssize
	OutputStreamWriteAll     func(o *OutputStream, buffer *T.Void, count T.Gsize, bytesWritten *T.Gsize, cancellable *Cancellable, err **L.Error) bool
	OutputStreamWriteAsync   func(o *OutputStream, buffer *T.Void, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	OutputStreamWriteFinish  func(o *OutputStream, result *AsyncResult, err **L.Error) T.Gssize
)

func (o *OutputStream) ClearPending() { OutputStreamClearPending(o) }
func (o *OutputStream) Close(cancellable *Cancellable, err **L.Error) bool {
	return OutputStreamClose(o, cancellable, err)
}
func (o *OutputStream) CloseAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	OutputStreamCloseAsync(o, ioPriority, cancellable, callback, userData)
}
func (o *OutputStream) CloseFinish(result *AsyncResult, err **L.Error) bool {
	return OutputStreamCloseFinish(o, result, err)
}
func (o *OutputStream) Flush(cancellable *Cancellable, err **L.Error) bool {
	return OutputStreamFlush(o, cancellable, err)
}
func (o *OutputStream) FlushAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	OutputStreamFlushAsync(o, ioPriority, cancellable, callback, userData)
}
func (o *OutputStream) FlushFinish(result *AsyncResult, err **L.Error) bool {
	return OutputStreamFlushFinish(o, result, err)
}
func (o *OutputStream) HasPending() bool              { return OutputStreamHasPending(o) }
func (o *OutputStream) IsClosed() bool                { return OutputStreamIsClosed(o) }
func (o *OutputStream) IsClosing() bool               { return OutputStreamIsClosing(o) }
func (o *OutputStream) SetPending(err **L.Error) bool { return OutputStreamSetPending(o, err) }
func (o *OutputStream) Splice(source *InputStream, flags OutputStreamSpliceFlags, cancellable *Cancellable, err **L.Error) T.Gssize {
	return OutputStreamSplice(o, source, flags, cancellable, err)
}
func (o *OutputStream) SpliceAsync(source *InputStream, flags OutputStreamSpliceFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	OutputStreamSpliceAsync(o, source, flags, ioPriority, cancellable, callback, userData)
}
func (o *OutputStream) SpliceFinish(result *AsyncResult, err **L.Error) T.Gssize {
	return OutputStreamSpliceFinish(o, result, err)
}
func (o *OutputStream) Write(buffer *T.Void, count T.Gsize, cancellable *Cancellable, err **L.Error) T.Gssize {
	return OutputStreamWrite(o, buffer, count, cancellable, err)
}
func (o *OutputStream) WriteAll(buffer *T.Void, count T.Gsize, bytesWritten *T.Gsize, cancellable *Cancellable, err **L.Error) bool {
	return OutputStreamWriteAll(o, buffer, count, bytesWritten, cancellable, err)
}
func (o *OutputStream) WriteAsync(buffer *T.Void, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	OutputStreamWriteAsync(o, buffer, count, ioPriority, cancellable, callback, userData)
}
func (o *OutputStream) WriteFinish(result *AsyncResult, err **L.Error) T.Gssize {
	return OutputStreamWriteFinish(o, result, err)
}

type OutputStreamSpliceFlags Enum

const (
	OUTPUT_STREAM_SPLICE_CLOSE_SOURCE OutputStreamSpliceFlags = 1 << iota
	OUTPUT_STREAM_SPLICE_CLOSE_TARGET
	OUTPUT_STREAM_SPLICE_NONE OutputStreamSpliceFlags = 0
)

var OutputStreamSpliceFlagsGetType func() O.Type
