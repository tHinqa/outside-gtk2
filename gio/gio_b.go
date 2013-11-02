// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type BufferedInputStream struct {
	Parent FilterInputStream
	_      *struct{}
}

var (
	BufferedInputStreamGetType  func() O.Type
	BufferedInputStreamNew      func(baseStream *InputStream) *InputStream
	BufferedInputStreamNewSized func(baseStream *InputStream, size T.Gsize) *InputStream

	BufferedInputStreamFill          func(b *BufferedInputStream, count T.Gssize, cancellable *Cancellable, err **L.Error) T.Gssize
	BufferedInputStreamFillAsync     func(b *BufferedInputStream, count T.Gssize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	BufferedInputStreamFillFinish    func(b *BufferedInputStream, result *AsyncResult, err **L.Error) T.Gssize
	BufferedInputStreamGetAvailable  func(b *BufferedInputStream) T.Gsize
	BufferedInputStreamGetBufferSize func(b *BufferedInputStream) T.Gsize
	BufferedInputStreamPeek          func(b *BufferedInputStream, buffer *T.Void, offset, count T.Gsize) T.Gsize
	BufferedInputStreamPeekBuffer    func(b *BufferedInputStream, count *T.Gsize) *T.Void
	BufferedInputStreamReadByte      func(b *BufferedInputStream, cancellable *Cancellable, err **L.Error) int
	BufferedInputStreamSetBufferSize func(b *BufferedInputStream, size T.Gsize)
)

func (b *BufferedInputStream) Fill(count T.Gssize, cancellable *Cancellable, err **L.Error) T.Gssize {
	return BufferedInputStreamFill(b, count, cancellable, err)
}
func (b *BufferedInputStream) FillAsync(count T.Gssize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	BufferedInputStreamFillAsync(b, count, ioPriority, cancellable, callback, userData)
}
func (b *BufferedInputStream) FillFinish(result *AsyncResult, err **L.Error) T.Gssize {
	return BufferedInputStreamFillFinish(b, result, err)
}
func (b *BufferedInputStream) GetAvailable() T.Gsize  { return BufferedInputStreamGetAvailable(b) }
func (b *BufferedInputStream) GetBufferSize() T.Gsize { return BufferedInputStreamGetBufferSize(b) }
func (b *BufferedInputStream) Peek(buffer *T.Void, offset, count T.Gsize) T.Gsize {
	return BufferedInputStreamPeek(b, buffer, offset, count)
}
func (b *BufferedInputStream) PeekBuffer(count *T.Gsize) *T.Void {
	return BufferedInputStreamPeekBuffer(b, count)
}
func (b *BufferedInputStream) ReadByte(cancellable *Cancellable, err **L.Error) int {
	return BufferedInputStreamReadByte(b, cancellable, err)
}
func (b *BufferedInputStream) SetBufferSize(size T.Gsize) { BufferedInputStreamSetBufferSize(b, size) }

type BufferedOutputStream struct {
	Parent FilterOutputStream
	_      *struct{}
}

var (
	BufferedOutputStreamGetType  func() O.Type
	BufferedOutputStreamNew      func(baseStream *OutputStream) *OutputStream
	BufferedOutputStreamNewSized func(baseStream *OutputStream, size T.Gsize) *OutputStream

	BufferedOutputStreamGetAutoGrow   func(b *BufferedOutputStream) bool
	BufferedOutputStreamGetBufferSize func(b *BufferedOutputStream) T.Gsize
	BufferedOutputStreamSetAutoGrow   func(b *BufferedOutputStream, autoGrow bool)
	BufferedOutputStreamSetBufferSize func(b *BufferedOutputStream, size T.Gsize)
)

func (b *BufferedOutputStream) GetAutoGrow() bool      { return BufferedOutputStreamGetAutoGrow(b) }
func (b *BufferedOutputStream) GetBufferSize() T.Gsize { return BufferedOutputStreamGetBufferSize(b) }
func (b *BufferedOutputStream) SetAutoGrow(autoGrow bool) {
	BufferedOutputStreamSetAutoGrow(b, autoGrow)
}
func (b *BufferedOutputStream) SetBufferSize(size T.Gsize) {
	BufferedOutputStreamSetBufferSize(b, size)
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
	BusOwnNameOnConnection               func(b *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredHandler BusNameAcquiredCallback, nameLostHandler BusNameLostCallback, userData T.Gpointer, userDataFreeFunc O.DestroyNotify) uint
	BusOwnNameOnConnectionWithClosures   func(b *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredClosure, nameLostClosure *O.Closure) uint
	BusWatchNameOnConnection             func(b *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedHandler BusNameAppearedCallback, nameVanishedHandler BusNameVanishedCallback, userData T.Gpointer, userDataFreeFunc O.DestroyNotify) uint
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
	BusGetFinish   func(res *AsyncResult, err **L.Error) *DBusConnection
	BusUnownName   func(ownerId uint)
	BusUnwatchName func(watcherId uint)

	BusGet                   func(b BusType, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	BusGetSync               func(b BusType, cancellable *Cancellable, err **L.Error) *DBusConnection
	BusOwnName               func(b BusType, name string, flags BusNameOwnerFlags, busAcquiredHandler BusAcquiredCallback, nameAcquiredHandler BusNameAcquiredCallback, nameLostHandler BusNameLostCallback, userData T.Gpointer, userDataFreeFunc O.DestroyNotify) uint
	BusOwnNameWithClosures   func(b BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure, nameAcquiredClosure, nameLostClosure *O.Closure) uint
	BusWatchName             func(b BusType, name string, flags BusNameWatcherFlags, nameAppearedHandler BusNameAppearedCallback, nameVanishedHandler BusNameVanishedCallback, userData T.Gpointer, userDataFreeFunc O.DestroyNotify) uint
	BusWatchNameWithClosures func(b BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure, nameVanishedClosure *O.Closure) uint
)

func (b BusType) Get(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	BusGet(b, cancellable, callback, userData)
}
func (b BusType) GetSync(cancellable *Cancellable, err **L.Error) *DBusConnection {
	return BusGetSync(b, cancellable, err)
}
func (b BusType) OwnName(name string, flags BusNameOwnerFlags, busAcquiredHandler BusAcquiredCallback, nameAcquiredHandler BusNameAcquiredCallback, nameLostHandler BusNameLostCallback, userData T.Gpointer, userDataFreeFunc O.DestroyNotify) uint {
	return BusOwnName(b, name, flags, busAcquiredHandler, nameAcquiredHandler, nameLostHandler, userData, userDataFreeFunc)
}
func (b BusType) OwnNameWithClosures(name string, flags BusNameOwnerFlags, busAcquiredClosure, nameAcquiredClosure, nameLostClosure *O.Closure) uint {
	return BusOwnNameWithClosures(b, name, flags, busAcquiredClosure, nameAcquiredClosure, nameLostClosure)
}
func (b BusType) WatchName(name string, flags BusNameWatcherFlags, nameAppearedHandler BusNameAppearedCallback, nameVanishedHandler BusNameVanishedCallback, userData T.Gpointer, userDataFreeFunc O.DestroyNotify) uint {
	return BusWatchName(b, name, flags, nameAppearedHandler, nameVanishedHandler, userData, userDataFreeFunc)
}
func (b BusType) WatchNameWithClosures(name string, flags BusNameWatcherFlags, nameAppearedClosure, nameVanishedClosure *O.Closure) uint {
	return BusWatchNameWithClosures(b, name, flags, nameAppearedClosure, nameVanishedClosure)
}
