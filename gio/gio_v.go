// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Vfs struct {
	Parent O.Object
}

var (
	VfsGetType func() O.Type

	VfsGetDefault func() *Vfs
	VfsGetLocal   func() *Vfs

	vfsGetFileForPath         func(v *Vfs, path string) *File
	vfsGetFileForUri          func(v *Vfs, uri string) *File
	vfsGetSupportedUriSchemes func(v *Vfs) []string
	vfsIsActive               func(v *Vfs) T.Gboolean
	vfsParseName              func(v *Vfs, parseName string) *File
)

func (v *Vfs) GetFileForPath(path string) *File { return vfsGetFileForPath(v, path) }
func (v *Vfs) GetFileForUri(uri string) *File   { return vfsGetFileForUri(v, uri) }
func (v *Vfs) GetSupportedUriSchemes() []string { return vfsGetSupportedUriSchemes(v) }
func (v *Vfs) IsActive() T.Gboolean             { return vfsIsActive(v) }
func (v *Vfs) ParseName(parseName string) *File { return vfsParseName(v, parseName) }

type Volume struct{}

var (
	VolumeGetType func() O.Type

	volumeCanEject                 func(v *Volume) T.Gboolean
	volumeCanMount                 func(v *Volume) T.Gboolean
	volumeEject                    func(v *Volume, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	volumeEjectFinish              func(v *Volume, result *AsyncResult, err **T.GError) T.Gboolean
	volumeEjectWithOperation       func(v *Volume, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	volumeEjectWithOperationFinish func(v *Volume, result *AsyncResult, err **T.GError) T.Gboolean
	volumeEnumerateIdentifiers     func(v *Volume) []string
	volumeGetActivationRoot        func(v *Volume) *File
	volumeGetDrive                 func(v *Volume) *Drive
	volumeGetIcon                  func(v *Volume) *Icon
	volumeGetIdentifier            func(v *Volume, kind string) string
	volumeGetMount                 func(v *Volume) *Mount
	volumeGetName                  func(v *Volume) string
	volumeGetUuid                  func(v *Volume) string
	volumeMount                    func(v *Volume, flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	volumeMountFinish              func(v *Volume, result *AsyncResult, err **T.GError) T.Gboolean
	volumeShouldAutomount          func(v *Volume) T.Gboolean
)

func (v *Volume) CanEject() T.Gboolean { return volumeCanEject(v) }
func (v *Volume) CanMount() T.Gboolean { return volumeCanMount(v) }
func (v *Volume) Eject(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	volumeEject(v, flags, cancellable, callback, userData)
}
func (v *Volume) EjectFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return volumeEjectFinish(v, result, err)
}
func (v *Volume) EjectWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	volumeEjectWithOperation(v, flags, mountOperation, cancellable, callback, userData)
}
func (v *Volume) EjectWithOperationFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return volumeEjectWithOperationFinish(v, result, err)
}
func (v *Volume) EnumerateIdentifiers() []string   { return volumeEnumerateIdentifiers(v) }
func (v *Volume) GetActivationRoot() *File         { return volumeGetActivationRoot(v) }
func (v *Volume) GetDrive() *Drive                 { return volumeGetDrive(v) }
func (v *Volume) GetIcon() *Icon                   { return volumeGetIcon(v) }
func (v *Volume) GetIdentifier(kind string) string { return volumeGetIdentifier(v, kind) }
func (v *Volume) GetMount() *Mount                 { return volumeGetMount(v) }
func (v *Volume) GetName() string                  { return volumeGetName(v) }
func (v *Volume) GetUuid() string                  { return volumeGetUuid(v) }
func (v *Volume) Mount(flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	volumeMount(v, flags, mountOperation, cancellable, callback, userData)
}
func (v *Volume) MountFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return volumeMountFinish(v, result, err)
}
func (v *Volume) ShouldAutomount() T.Gboolean { return volumeShouldAutomount(v) }

type VolumeMonitor struct {
	Parent O.Object
	Priv   T.Gpointer
}

var (
	VolumeMonitorGetType func() O.Type

	VolumeMonitorAdoptOrphanMount func(mount *Mount) *Volume
	VolumeMonitorGet              func() *VolumeMonitor

	volumeMonitorGetConnectedDrives func(v *VolumeMonitor) *T.GList
	volumeMonitorGetMountForUuid    func(v *VolumeMonitor, uuid string) *Mount
	volumeMonitorGetMounts          func(v *VolumeMonitor) *T.GList
	volumeMonitorGetVolumeForUuid   func(v *VolumeMonitor, uuid string) *Volume
	volumeMonitorGetVolumes         func(v *VolumeMonitor) *T.GList
)

func (v *VolumeMonitor) GetConnectedDrives() *T.GList { return volumeMonitorGetConnectedDrives(v) }
func (v *VolumeMonitor) GetMountForUuid(uuid string) *Mount {
	return volumeMonitorGetMountForUuid(v, uuid)
}
func (v *VolumeMonitor) GetMounts() *T.GList { return volumeMonitorGetMounts(v) }
func (v *VolumeMonitor) GetVolumeForUuid(uuid string) *Volume {
	return volumeMonitorGetVolumeForUuid(v, uuid)
}
func (v *VolumeMonitor) GetVolumes() *T.GList { return volumeMonitorGetVolumes(v) }
