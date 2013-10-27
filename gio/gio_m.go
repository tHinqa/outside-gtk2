// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type MemoryInputStream struct {
	Parent InputStream
	_      *struct{}
}

var (
	MemoryInputStreamGetType     func() O.Type
	MemoryInputStreamNew         func() *InputStream
	MemoryInputStreamNewFromData func(data *T.Void, len T.Gssize, destroy T.GDestroyNotify) *InputStream

	memoryInputStreamAddData func(m *MemoryInputStream, data *T.Void, len T.Gssize, destroy T.GDestroyNotify)
)

func (m *MemoryInputStream) AddData(data *T.Void, len T.Gssize, destroy T.GDestroyNotify) {
	memoryInputStreamAddData(m, data, len, destroy)
}

type MemoryOutputStream struct {
	Parent OutputStream
	Priv   *struct{}
}

var (
	MemoryOutputStreamGetType func() O.Type
	MemoryOutputStreamNew     func(data T.Gpointer, size T.Gsize, reallocFunction T.GReallocFunc, destroyFunction T.GDestroyNotify) *OutputStream

	memoryOutputStreamGetData     func(m *MemoryOutputStream) T.Gpointer
	memoryOutputStreamGetDataSize func(m *MemoryOutputStream) T.Gsize
	memoryOutputStreamGetSize     func(m *MemoryOutputStream) T.Gsize
	memoryOutputStreamStealData   func(m *MemoryOutputStream) T.Gpointer
)

func (m *MemoryOutputStream) GetData() T.Gpointer   { return memoryOutputStreamGetData(m) }
func (m *MemoryOutputStream) GetDataSize() T.Gsize  { return memoryOutputStreamGetDataSize(m) }
func (m *MemoryOutputStream) GetSize() T.Gsize      { return memoryOutputStreamGetSize(m) }
func (m *MemoryOutputStream) StealData() T.Gpointer { return memoryOutputStreamStealData(m) }

var MemorySettingsBackendNew func() *SettingsBackend

var (
	MountGetType func() O.Type

	mountCanEject                   func(m *Mount) T.Gboolean
	mountCanUnmount                 func(m *Mount) T.Gboolean
	mountEject                      func(m *Mount, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	mountEjectFinish                func(m *Mount, result *AsyncResult, err **T.GError) T.Gboolean
	mountEjectWithOperation         func(m *Mount, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	mountEjectWithOperationFinish   func(m *Mount, result *AsyncResult, err **T.GError) T.Gboolean
	mountGetDefaultLocation         func(m *Mount) *File
	mountGetDrive                   func(m *Mount) *Drive
	mountGetIcon                    func(m *Mount) *Icon
	mountGetName                    func(m *Mount) string
	mountGetRoot                    func(m *Mount) *File
	mountGetUuid                    func(m *Mount) string
	mountGetVolume                  func(m *Mount) *Volume
	mountGuessContentType           func(m *Mount, forceRescan T.Gboolean, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	mountGuessContentTypeFinish     func(m *Mount, result *AsyncResult, err **T.GError) **T.Gchar
	mountGuessContentTypeSync       func(m *Mount, forceRescan T.Gboolean, cancellable *Cancellable, err **T.GError) **T.Gchar
	mountIsShadowed                 func(m *Mount) T.Gboolean
	mountRemount                    func(m *Mount, flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	mountRemountFinish              func(m *Mount, result *AsyncResult, err **T.GError) T.Gboolean
	mountShadow                     func(m *Mount)
	mountUnmount                    func(m *Mount, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	mountUnmountFinish              func(m *Mount, result *AsyncResult, err **T.GError) T.Gboolean
	mountUnmountWithOperation       func(m *Mount, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	mountUnmountWithOperationFinish func(m *Mount, result *AsyncResult, err **T.GError) T.Gboolean
	mountUnshadow                   func(m *Mount)
)

type Mount struct{}

func (m *Mount) CanEject() T.Gboolean   { return mountCanEject(m) }
func (m *Mount) CanUnmount() T.Gboolean { return mountCanUnmount(m) }
func (m *Mount) Eject(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	mountEject(m, flags, cancellable, callback, userData)
}
func (m *Mount) EjectFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return mountEjectFinish(m, result, err)
}
func (m *Mount) EjectWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	mountEjectWithOperation(m, flags, mountOperation, cancellable, callback, userData)
}
func (m *Mount) EjectWithOperationFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return mountEjectWithOperationFinish(m, result, err)
}
func (m *Mount) GetDefaultLocation() *File { return mountGetDefaultLocation(m) }
func (m *Mount) GetDrive() *Drive          { return mountGetDrive(m) }
func (m *Mount) GetIcon() *Icon            { return mountGetIcon(m) }
func (m *Mount) GetName() string           { return mountGetName(m) }
func (m *Mount) GetRoot() *File            { return mountGetRoot(m) }
func (m *Mount) GetUuid() string           { return mountGetUuid(m) }
func (m *Mount) GetVolume() *Volume        { return mountGetVolume(m) }
func (m *Mount) GuessContentType(forceRescan T.Gboolean, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	mountGuessContentType(m, forceRescan, cancellable, callback, userData)
}
func (m *Mount) GuessContentTypeFinish(result *AsyncResult, err **T.GError) **T.Gchar {
	return mountGuessContentTypeFinish(m, result, err)
}
func (m *Mount) GuessContentTypeSync(forceRescan T.Gboolean, cancellable *Cancellable, err **T.GError) **T.Gchar {
	return mountGuessContentTypeSync(m, forceRescan, cancellable, err)
}
func (m *Mount) IsShadowed() T.Gboolean { return mountIsShadowed(m) }
func (m *Mount) Remount(flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	mountRemount(m, flags, mountOperation, cancellable, callback, userData)
}
func (m *Mount) RemountFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return mountRemountFinish(m, result, err)
}
func (m *Mount) Shadow() { mountShadow(m) }
func (m *Mount) Unmount(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	mountUnmount(m, flags, cancellable, callback, userData)
}
func (m *Mount) UnmountFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return mountUnmountFinish(m, result, err)
}
func (m *Mount) UnmountWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	mountUnmountWithOperation(m, flags, mountOperation, cancellable, callback, userData)
}
func (m *Mount) UnmountWithOperationFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return mountUnmountWithOperationFinish(m, result, err)
}
func (m *Mount) Unshadow() { mountUnshadow(m) }

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

	mountOperationGetAnonymous    func(m *MountOperation) T.Gboolean
	mountOperationGetChoice       func(m *MountOperation) int
	mountOperationGetDomain       func(m *MountOperation) string
	mountOperationGetPassword     func(m *MountOperation) string
	mountOperationGetPasswordSave func(m *MountOperation) T.GPasswordSave
	mountOperationGetUsername     func(m *MountOperation) string
	mountOperationReply           func(m *MountOperation, result MountOperationResult)
	mountOperationSetAnonymous    func(m *MountOperation, anonymous T.Gboolean)
	mountOperationSetChoice       func(m *MountOperation, choice int)
	mountOperationSetDomain       func(m *MountOperation, domain string)
	mountOperationSetPassword     func(m *MountOperation, password string)
	mountOperationSetPasswordSave func(m *MountOperation, save T.GPasswordSave)
	mountOperationSetUsername     func(m *MountOperation, username string)
)

func (m *MountOperation) GetAnonymous() T.Gboolean          { return mountOperationGetAnonymous(m) }
func (m *MountOperation) GetChoice() int                    { return mountOperationGetChoice(m) }
func (m *MountOperation) GetDomain() string                 { return mountOperationGetDomain(m) }
func (m *MountOperation) GetPassword() string               { return mountOperationGetPassword(m) }
func (m *MountOperation) GetPasswordSave() T.GPasswordSave  { return mountOperationGetPasswordSave(m) }
func (m *MountOperation) GetUsername() string               { return mountOperationGetUsername(m) }
func (m *MountOperation) Reply(result MountOperationResult) { mountOperationReply(m, result) }
func (m *MountOperation) SetAnonymous(anonymous T.Gboolean) { mountOperationSetAnonymous(m, anonymous) }
func (m *MountOperation) SetChoice(choice int)              { mountOperationSetChoice(m, choice) }
func (m *MountOperation) SetDomain(domain string)           { mountOperationSetDomain(m, domain) }
func (m *MountOperation) SetPassword(password string)       { mountOperationSetPassword(m, password) }
func (m *MountOperation) SetPasswordSave(save T.GPasswordSave) {
	mountOperationSetPasswordSave(m, save)
}
func (m *MountOperation) SetUsername(username string) { mountOperationSetUsername(m, username) }

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
