// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type BufferedInputStream struct {
	Parent FilterInputStream
	_      *struct{}
}

var (
	BufferedInputStreamGetType  func() O.Type
	BufferedInputStreamNew      func(baseStream *InputStream) *InputStream
	BufferedInputStreamNewSized func(baseStream *InputStream, size T.Gsize) *InputStream

	bufferedInputStreamFill          func(b *BufferedInputStream, count T.Gssize, cancellable *Cancellable, err **T.GError) T.Gssize
	bufferedInputStreamFillAsync     func(b *BufferedInputStream, count T.Gssize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	bufferedInputStreamFillFinish    func(b *BufferedInputStream, result *AsyncResult, err **T.GError) T.Gssize
	bufferedInputStreamGetAvailable  func(b *BufferedInputStream) T.Gsize
	bufferedInputStreamGetBufferSize func(b *BufferedInputStream) T.Gsize
	bufferedInputStreamPeek          func(b *BufferedInputStream, buffer *T.Void, offset, count T.Gsize) T.Gsize
	bufferedInputStreamPeekBuffer    func(b *BufferedInputStream, count *T.Gsize) *T.Void
	bufferedInputStreamReadByte      func(b *BufferedInputStream, cancellable *Cancellable, err **T.GError) int
	bufferedInputStreamSetBufferSize func(b *BufferedInputStream, size T.Gsize)
)

func (b *BufferedInputStream) Fill(count T.Gssize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return bufferedInputStreamFill(b, count, cancellable, err)
}
func (b *BufferedInputStream) FillAsync(count T.Gssize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	bufferedInputStreamFillAsync(b, count, ioPriority, cancellable, callback, userData)
}
func (b *BufferedInputStream) FillFinish(result *AsyncResult, err **T.GError) T.Gssize {
	return bufferedInputStreamFillFinish(b, result, err)
}
func (b *BufferedInputStream) GetAvailable() T.Gsize  { return bufferedInputStreamGetAvailable(b) }
func (b *BufferedInputStream) GetBufferSize() T.Gsize { return bufferedInputStreamGetBufferSize(b) }
func (b *BufferedInputStream) Peek(buffer *T.Void, offset, count T.Gsize) T.Gsize {
	return bufferedInputStreamPeek(b, buffer, offset, count)
}
func (b *BufferedInputStream) PeekBuffer(count *T.Gsize) *T.Void {
	return bufferedInputStreamPeekBuffer(b, count)
}
func (b *BufferedInputStream) ReadByte(cancellable *Cancellable, err **T.GError) int {
	return bufferedInputStreamReadByte(b, cancellable, err)
}
func (b *BufferedInputStream) SetBufferSize(size T.Gsize) { bufferedInputStreamSetBufferSize(b, size) }

type BufferedOutputStream struct {
	Parent FilterOutputStream
	_      *struct{}
}

var (
	BufferedOutputStreamGetType  func() O.Type
	BufferedOutputStreamNew      func(baseStream *OutputStream) *OutputStream
	BufferedOutputStreamNewSized func(baseStream *OutputStream, size T.Gsize) *OutputStream

	bufferedOutputStreamGetAutoGrow   func(b *BufferedOutputStream) bool
	bufferedOutputStreamGetBufferSize func(b *BufferedOutputStream) T.Gsize
	bufferedOutputStreamSetAutoGrow   func(b *BufferedOutputStream, autoGrow bool)
	bufferedOutputStreamSetBufferSize func(b *BufferedOutputStream, size T.Gsize)
)

func (b *BufferedOutputStream) GetAutoGrow() bool      { return bufferedOutputStreamGetAutoGrow(b) }
func (b *BufferedOutputStream) GetBufferSize() T.Gsize { return bufferedOutputStreamGetBufferSize(b) }
func (b *BufferedOutputStream) SetAutoGrow(autoGrow bool) {
	bufferedOutputStreamSetAutoGrow(b, autoGrow)
}
func (b *BufferedOutputStream) SetBufferSize(size T.Gsize) {
	bufferedOutputStreamSetBufferSize(b, size)
}

type BusAcquiredCallback func(
	connection *DBusConnection,
	name string,
	userData T.Gpointer)

type BusNameAcquiredCallback func(
	connection *DBusConnection,
	name string,
	userData T.Gpointer)

type BusNameAppearedCallback func(
	connection *DBusConnection,
	name string,
	nameOwner string,
	userData T.Gpointer)

type BusNameLostCallback func(
	connection *DBusConnection,
	name string,
	userData T.Gpointer)

type BusNameOwnerFlags Enum

const (
	BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT BusNameOwnerFlags = 1 << iota
	BUS_NAME_OWNER_FLAGS_REPLACE
	BUS_NAME_OWNER_FLAGS_NONE BusNameOwnerFlags = 0
)

var BusNameOwnerFlagsGetType func() O.Type

type BusNameVanishedCallback func(
	connection *DBusConnection,
	name string,
	userData T.Gpointer)

type BusNameWatcherFlags Enum

const (
	BUS_NAME_WATCHER_FLAGS_AUTO_START BusNameWatcherFlags = 1 << iota
	BUS_NAME_WATCHER_FLAGS_NONE       BusNameWatcherFlags = 0
)

var BusNameWatcherFlagsGetType func() O.Type

var (
	BusOwnNameOnConnection               func(b *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredHandler BusNameAcquiredCallback, nameLostHandler BusNameLostCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint
	BusOwnNameOnConnectionWithClosures   func(b *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredClosure, nameLostClosure *O.Closure) uint
	BusWatchNameOnConnection             func(b *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedHandler BusNameAppearedCallback, nameVanishedHandler BusNameVanishedCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint
	BusWatchNameOnConnectionWithClosures func(b *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedClosure, nameVanishedClosure *O.Closure) uint
)

type BusType Enum

const (
	BUS_TYPE_STARTER BusType = iota - 1
	BUS_TYPE_NONE
	BUS_TYPE_SYSTEM
	BUS_TYPE_SESSION
)

var (
	BusTypeGetType func() O.Type
	BusGetFinish   func(res *AsyncResult, err **T.GError) *DBusConnection
	BusUnownName   func(ownerId uint)
	BusUnwatchName func(watcherId uint)

	busGet                   func(b BusType, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	busGetSync               func(b BusType, cancellable *Cancellable, err **T.GError) *DBusConnection
	busOwnName               func(b BusType, name string, flags BusNameOwnerFlags, busAcquiredHandler BusAcquiredCallback, nameAcquiredHandler BusNameAcquiredCallback, nameLostHandler BusNameLostCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint
	busOwnNameWithClosures   func(b BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure, nameAcquiredClosure, nameLostClosure *O.Closure) uint
	busWatchName             func(b BusType, name string, flags BusNameWatcherFlags, nameAppearedHandler BusNameAppearedCallback, nameVanishedHandler BusNameVanishedCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint
	busWatchNameWithClosures func(b BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure, nameVanishedClosure *O.Closure) uint
)

func (b BusType) Get(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	busGet(b, cancellable, callback, userData)
}
func (b BusType) GetSync(cancellable *Cancellable, err **T.GError) *DBusConnection {
	return busGetSync(b, cancellable, err)
}
func (b BusType) OwnName(name string, flags BusNameOwnerFlags, busAcquiredHandler BusAcquiredCallback, nameAcquiredHandler BusNameAcquiredCallback, nameLostHandler BusNameLostCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint {
	return busOwnName(b, name, flags, busAcquiredHandler, nameAcquiredHandler, nameLostHandler, userData, userDataFreeFunc)
}
func (b BusType) OwnNameWithClosures(name string, flags BusNameOwnerFlags, busAcquiredClosure, nameAcquiredClosure, nameLostClosure *O.Closure) uint {
	return busOwnNameWithClosures(b, name, flags, busAcquiredClosure, nameAcquiredClosure, nameLostClosure)
}
func (b BusType) WatchName(name string, flags BusNameWatcherFlags, nameAppearedHandler BusNameAppearedCallback, nameVanishedHandler BusNameVanishedCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint {
	return busWatchName(b, name, flags, nameAppearedHandler, nameVanishedHandler, userData, userDataFreeFunc)
}
func (b BusType) WatchNameWithClosures(name string, flags BusNameWatcherFlags, nameAppearedClosure, nameVanishedClosure *O.Closure) uint {
	return busWatchNameWithClosures(b, name, flags, nameAppearedClosure, nameVanishedClosure)
}
