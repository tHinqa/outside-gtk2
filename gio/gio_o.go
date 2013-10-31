// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
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

	outputStreamClearPending func(o *OutputStream)
	outputStreamClose        func(o *OutputStream, cancellable *Cancellable, err **T.GError) bool
	outputStreamCloseAsync   func(o *OutputStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	outputStreamCloseFinish  func(o *OutputStream, result *AsyncResult, err **T.GError) bool
	outputStreamFlush        func(o *OutputStream, cancellable *Cancellable, err **T.GError) bool
	outputStreamFlushAsync   func(o *OutputStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	outputStreamFlushFinish  func(o *OutputStream, result *AsyncResult, err **T.GError) bool
	outputStreamHasPending   func(o *OutputStream) bool
	outputStreamIsClosed     func(o *OutputStream) bool
	outputStreamIsClosing    func(o *OutputStream) bool
	outputStreamSetPending   func(o *OutputStream, err **T.GError) bool
	outputStreamSplice       func(o *OutputStream, source *InputStream, flags OutputStreamSpliceFlags, cancellable *Cancellable, err **T.GError) T.Gssize
	outputStreamSpliceAsync  func(o *OutputStream, source *InputStream, flags OutputStreamSpliceFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	outputStreamSpliceFinish func(o *OutputStream, result *AsyncResult, err **T.GError) T.Gssize
	outputStreamWrite        func(o *OutputStream, buffer *T.Void, count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	outputStreamWriteAll     func(o *OutputStream, buffer *T.Void, count T.Gsize, bytesWritten *T.Gsize, cancellable *Cancellable, err **T.GError) bool
	outputStreamWriteAsync   func(o *OutputStream, buffer *T.Void, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	outputStreamWriteFinish  func(o *OutputStream, result *AsyncResult, err **T.GError) T.Gssize
)

func (o *OutputStream) ClearPending() { outputStreamClearPending(o) }
func (o *OutputStream) Close(cancellable *Cancellable, err **T.GError) bool {
	return outputStreamClose(o, cancellable, err)
}
func (o *OutputStream) CloseAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	outputStreamCloseAsync(o, ioPriority, cancellable, callback, userData)
}
func (o *OutputStream) CloseFinish(result *AsyncResult, err **T.GError) bool {
	return outputStreamCloseFinish(o, result, err)
}
func (o *OutputStream) Flush(cancellable *Cancellable, err **T.GError) bool {
	return outputStreamFlush(o, cancellable, err)
}
func (o *OutputStream) FlushAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	outputStreamFlushAsync(o, ioPriority, cancellable, callback, userData)
}
func (o *OutputStream) FlushFinish(result *AsyncResult, err **T.GError) bool {
	return outputStreamFlushFinish(o, result, err)
}
func (o *OutputStream) HasPending() bool               { return outputStreamHasPending(o) }
func (o *OutputStream) IsClosed() bool                 { return outputStreamIsClosed(o) }
func (o *OutputStream) IsClosing() bool                { return outputStreamIsClosing(o) }
func (o *OutputStream) SetPending(err **T.GError) bool { return outputStreamSetPending(o, err) }
func (o *OutputStream) Splice(source *InputStream, flags OutputStreamSpliceFlags, cancellable *Cancellable, err **T.GError) T.Gssize {
	return outputStreamSplice(o, source, flags, cancellable, err)
}
func (o *OutputStream) SpliceAsync(source *InputStream, flags OutputStreamSpliceFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	outputStreamSpliceAsync(o, source, flags, ioPriority, cancellable, callback, userData)
}
func (o *OutputStream) SpliceFinish(result *AsyncResult, err **T.GError) T.Gssize {
	return outputStreamSpliceFinish(o, result, err)
}
func (o *OutputStream) Write(buffer *T.Void, count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return outputStreamWrite(o, buffer, count, cancellable, err)
}
func (o *OutputStream) WriteAll(buffer *T.Void, count T.Gsize, bytesWritten *T.Gsize, cancellable *Cancellable, err **T.GError) bool {
	return outputStreamWriteAll(o, buffer, count, bytesWritten, cancellable, err)
}
func (o *OutputStream) WriteAsync(buffer *T.Void, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	outputStreamWriteAsync(o, buffer, count, ioPriority, cancellable, callback, userData)
}
func (o *OutputStream) WriteFinish(result *AsyncResult, err **T.GError) T.Gssize {
	return outputStreamWriteFinish(o, result, err)
}

type OutputStreamSpliceFlags Enum

const (
	OUTPUT_STREAM_SPLICE_CLOSE_SOURCE OutputStreamSpliceFlags = 1 << iota
	OUTPUT_STREAM_SPLICE_CLOSE_TARGET
	OUTPUT_STREAM_SPLICE_NONE OutputStreamSpliceFlags = 0
)

var OutputStreamSpliceFlagsGetType func() O.Type
