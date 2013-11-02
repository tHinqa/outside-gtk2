// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type MemoryInputStream struct {
	Parent InputStream
	_      *struct{}
}

var (
	MemoryInputStreamGetType     func() O.Type
	MemoryInputStreamNew         func() *InputStream
	MemoryInputStreamNewFromData func(data *T.Void, len T.Gssize, destroy O.DestroyNotify) *InputStream

	MemoryInputStreamAddData func(m *MemoryInputStream, data *T.Void, len T.Gssize, destroy O.DestroyNotify)
)

func (m *MemoryInputStream) AddData(data *T.Void, len T.Gssize, destroy O.DestroyNotify) {
	MemoryInputStreamAddData(m, data, len, destroy)
}

type MemoryOutputStream struct {
	Parent OutputStream
	Priv   *struct{}
}

var (
	MemoryOutputStreamGetType func() O.Type
	MemoryOutputStreamNew     func(data T.Gpointer, size T.Gsize, reallocFunction T.GReallocFunc, destroyFunction O.DestroyNotify) *OutputStream

	MemoryOutputStreamGetData     func(m *MemoryOutputStream) T.Gpointer
	MemoryOutputStreamGetDataSize func(m *MemoryOutputStream) T.Gsize
	MemoryOutputStreamGetSize     func(m *MemoryOutputStream) T.Gsize
	MemoryOutputStreamStealData   func(m *MemoryOutputStream) T.Gpointer
)

func (m *MemoryOutputStream) GetData() T.Gpointer   { return MemoryOutputStreamGetData(m) }
func (m *MemoryOutputStream) GetDataSize() T.Gsize  { return MemoryOutputStreamGetDataSize(m) }
func (m *MemoryOutputStream) GetSize() T.Gsize      { return MemoryOutputStreamGetSize(m) }
func (m *MemoryOutputStream) StealData() T.Gpointer { return MemoryOutputStreamStealData(m) }

var MemorySettingsBackendNew func() *SettingsBackend

var (
	MountGetType func() O.Type

	MountCanEject                   func(m *Mount) bool
	MountCanUnmount                 func(m *Mount) bool
	MountEject                      func(m *Mount, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	MountEjectFinish                func(m *Mount, result *AsyncResult, err **L.Error) bool
	MountEjectWithOperation         func(m *Mount, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	MountEjectWithOperationFinish   func(m *Mount, result *AsyncResult, err **L.Error) bool
	MountGetDefaultLocation         func(m *Mount) *File
	MountGetDrive                   func(m *Mount) *Drive
	MountGetIcon                    func(m *Mount) *Icon
	MountGetName                    func(m *Mount) string
	MountGetRoot                    func(m *Mount) *File
	MountGetUuid                    func(m *Mount) string
	MountGetVolume                  func(m *Mount) *Volume
	MountGuessContentType           func(m *Mount, forceRescan bool, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	MountGuessContentTypeFinish     func(m *Mount, result *AsyncResult, err **L.Error) []string
	MountGuessContentTypeSync       func(m *Mount, forceRescan bool, cancellable *Cancellable, err **L.Error) []string
	MountIsShadowed                 func(m *Mount) bool
	MountRemount                    func(m *Mount, flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	MountRemountFinish              func(m *Mount, result *AsyncResult, err **L.Error) bool
	MountShadow                     func(m *Mount)
	MountUnmount                    func(m *Mount, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	MountUnmountFinish              func(m *Mount, result *AsyncResult, err **L.Error) bool
	MountUnmountWithOperation       func(m *Mount, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	MountUnmountWithOperationFinish func(m *Mount, result *AsyncResult, err **L.Error) bool
	MountUnshadow                   func(m *Mount)
)

type Mount struct{}

func (m *Mount) CanEject() bool   { return MountCanEject(m) }
func (m *Mount) CanUnmount() bool { return MountCanUnmount(m) }
func (m *Mount) Eject(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	MountEject(m, flags, cancellable, callback, userData)
}
func (m *Mount) EjectFinish(result *AsyncResult, err **L.Error) bool {
	return MountEjectFinish(m, result, err)
}
func (m *Mount) EjectWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	MountEjectWithOperation(m, flags, mountOperation, cancellable, callback, userData)
}
func (m *Mount) EjectWithOperationFinish(result *AsyncResult, err **L.Error) bool {
	return MountEjectWithOperationFinish(m, result, err)
}
func (m *Mount) GetDefaultLocation() *File { return MountGetDefaultLocation(m) }
func (m *Mount) GetDrive() *Drive          { return MountGetDrive(m) }
func (m *Mount) GetIcon() *Icon            { return MountGetIcon(m) }
func (m *Mount) GetName() string           { return MountGetName(m) }
func (m *Mount) GetRoot() *File            { return MountGetRoot(m) }
func (m *Mount) GetUuid() string           { return MountGetUuid(m) }
func (m *Mount) GetVolume() *Volume        { return MountGetVolume(m) }
func (m *Mount) GuessContentType(forceRescan bool, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	MountGuessContentType(m, forceRescan, cancellable, callback, userData)
}
func (m *Mount) GuessContentTypeFinish(result *AsyncResult, err **L.Error) []string {
	return MountGuessContentTypeFinish(m, result, err)
}
func (m *Mount) GuessContentTypeSync(forceRescan bool, cancellable *Cancellable, err **L.Error) []string {
	return MountGuessContentTypeSync(m, forceRescan, cancellable, err)
}
func (m *Mount) IsShadowed() bool { return MountIsShadowed(m) }
func (m *Mount) Remount(flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	MountRemount(m, flags, mountOperation, cancellable, callback, userData)
}
func (m *Mount) RemountFinish(result *AsyncResult, err **L.Error) bool {
	return MountRemountFinish(m, result, err)
}
func (m *Mount) Shadow() { MountShadow(m) }
func (m *Mount) Unmount(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	MountUnmount(m, flags, cancellable, callback, userData)
}
func (m *Mount) UnmountFinish(result *AsyncResult, err **L.Error) bool {
	return MountUnmountFinish(m, result, err)
}
func (m *Mount) UnmountWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	MountUnmountWithOperation(m, flags, mountOperation, cancellable, callback, userData)
}
func (m *Mount) UnmountWithOperationFinish(result *AsyncResult, err **L.Error) bool {
	return MountUnmountWithOperationFinish(m, result, err)
}
func (m *Mount) Unshadow() { MountUnshadow(m) }

type MountMountFlags Enum

const MOUNT_MOUNT_NONE MountMountFlags = 0

var MountMountFlagsGetType func() O.Type

type MountOperation struct {
	Parent O.Object
	_      *struct{}
}

var (
	MountOperationGetType func() O.Type
	MountOperationNew     func() *MountOperation

	MountOperationGetAnonymous    func(m *MountOperation) bool
	MountOperationGetChoice       func(m *MountOperation) int
	MountOperationGetDomain       func(m *MountOperation) string
	MountOperationGetPassword     func(m *MountOperation) string
	MountOperationGetPasswordSave func(m *MountOperation) T.GPasswordSave
	MountOperationGetUsername     func(m *MountOperation) string
	MountOperationReply           func(m *MountOperation, result MountOperationResult)
	MountOperationSetAnonymous    func(m *MountOperation, anonymous bool)
	MountOperationSetChoice       func(m *MountOperation, choice int)
	MountOperationSetDomain       func(m *MountOperation, domain string)
	MountOperationSetPassword     func(m *MountOperation, password string)
	MountOperationSetPasswordSave func(m *MountOperation, save T.GPasswordSave)
	MountOperationSetUsername     func(m *MountOperation, username string)
)

func (m *MountOperation) GetAnonymous() bool                { return MountOperationGetAnonymous(m) }
func (m *MountOperation) GetChoice() int                    { return MountOperationGetChoice(m) }
func (m *MountOperation) GetDomain() string                 { return MountOperationGetDomain(m) }
func (m *MountOperation) GetPassword() string               { return MountOperationGetPassword(m) }
func (m *MountOperation) GetPasswordSave() T.GPasswordSave  { return MountOperationGetPasswordSave(m) }
func (m *MountOperation) GetUsername() string               { return MountOperationGetUsername(m) }
func (m *MountOperation) Reply(result MountOperationResult) { MountOperationReply(m, result) }
func (m *MountOperation) SetAnonymous(anonymous bool)       { MountOperationSetAnonymous(m, anonymous) }
func (m *MountOperation) SetChoice(choice int)              { MountOperationSetChoice(m, choice) }
func (m *MountOperation) SetDomain(domain string)           { MountOperationSetDomain(m, domain) }
func (m *MountOperation) SetPassword(password string)       { MountOperationSetPassword(m, password) }
func (m *MountOperation) SetPasswordSave(save T.GPasswordSave) {
	MountOperationSetPasswordSave(m, save)
}
func (m *MountOperation) SetUsername(username string) { MountOperationSetUsername(m, username) }

type MountOperationResult Enum

const (
	MOUNT_OPERATION_HANDLED MountOperationResult = iota
	MOUNT_OPERATION_ABORTED
	MOUNT_OPERATION_UNHANDLED
)

var MountOperationResultGetType func() O.Type

type MountUnmountFlags Enum

const (
	MOUNT_UNMOUNT_FORCE MountUnmountFlags = 1 << iota
	MOUNT_UNMOUNT_NONE  MountUnmountFlags = 0
)

var MountUnmountFlagsGetType func() O.Type
