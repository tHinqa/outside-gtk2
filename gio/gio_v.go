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

	VfsGetFileForPath         func(v *Vfs, path string) *File
	VfsGetFileForUri          func(v *Vfs, uri string) *File
	VfsGetSupportedUriSchemes func(v *Vfs) []string
	VfsIsActive               func(v *Vfs) bool
	VfsParseName              func(v *Vfs, parseName string) *File
)

func (v *Vfs) GetFileForPath(path string) *File { return VfsGetFileForPath(v, path) }
func (v *Vfs) GetFileForUri(uri string) *File   { return VfsGetFileForUri(v, uri) }
func (v *Vfs) GetSupportedUriSchemes() []string { return VfsGetSupportedUriSchemes(v) }
func (v *Vfs) IsActive() bool                   { return VfsIsActive(v) }
func (v *Vfs) ParseName(parseName string) *File { return VfsParseName(v, parseName) }

type Volume struct{}

var (
	VolumeGetType func() O.Type

	VolumeCanEject                 func(v *Volume) bool
	VolumeCanMount                 func(v *Volume) bool
	VolumeEject                    func(v *Volume, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	VolumeEjectFinish              func(v *Volume, result *AsyncResult, err **T.GError) bool
	VolumeEjectWithOperation       func(v *Volume, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	VolumeEjectWithOperationFinish func(v *Volume, result *AsyncResult, err **T.GError) bool
	VolumeEnumerateIdentifiers     func(v *Volume) []string
	VolumeGetActivationRoot        func(v *Volume) *File
	VolumeGetDrive                 func(v *Volume) *Drive
	VolumeGetIcon                  func(v *Volume) *Icon
	VolumeGetIdentifier            func(v *Volume, kind string) string
	VolumeGetMount                 func(v *Volume) *Mount
	VolumeGetName                  func(v *Volume) string
	VolumeGetUuid                  func(v *Volume) string
	VolumeMount                    func(v *Volume, flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	VolumeMountFinish              func(v *Volume, result *AsyncResult, err **T.GError) bool
	VolumeShouldAutomount          func(v *Volume) bool
)

func (v *Volume) CanEject() bool { return VolumeCanEject(v) }
func (v *Volume) CanMount() bool { return VolumeCanMount(v) }
func (v *Volume) Eject(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	VolumeEject(v, flags, cancellable, callback, userData)
}
func (v *Volume) EjectFinish(result *AsyncResult, err **T.GError) bool {
	return VolumeEjectFinish(v, result, err)
}
func (v *Volume) EjectWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	VolumeEjectWithOperation(v, flags, mountOperation, cancellable, callback, userData)
}
func (v *Volume) EjectWithOperationFinish(result *AsyncResult, err **T.GError) bool {
	return VolumeEjectWithOperationFinish(v, result, err)
}
func (v *Volume) EnumerateIdentifiers() []string   { return VolumeEnumerateIdentifiers(v) }
func (v *Volume) GetActivationRoot() *File         { return VolumeGetActivationRoot(v) }
func (v *Volume) GetDrive() *Drive                 { return VolumeGetDrive(v) }
func (v *Volume) GetIcon() *Icon                   { return VolumeGetIcon(v) }
func (v *Volume) GetIdentifier(kind string) string { return VolumeGetIdentifier(v, kind) }
func (v *Volume) GetMount() *Mount                 { return VolumeGetMount(v) }
func (v *Volume) GetName() string                  { return VolumeGetName(v) }
func (v *Volume) GetUuid() string                  { return VolumeGetUuid(v) }
func (v *Volume) Mount(flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	VolumeMount(v, flags, mountOperation, cancellable, callback, userData)
}
func (v *Volume) MountFinish(result *AsyncResult, err **T.GError) bool {
	return VolumeMountFinish(v, result, err)
}
func (v *Volume) ShouldAutomount() bool { return VolumeShouldAutomount(v) }

type VolumeMonitor struct {
	Parent O.Object
	Priv   T.Gpointer
}

var (
	VolumeMonitorGetType func() O.Type

	VolumeMonitorAdoptOrphanMount func(mount *Mount) *Volume
	VolumeMonitorGet              func() *VolumeMonitor

	VolumeMonitorGetConnectedDrives func(v *VolumeMonitor) *T.GList
	VolumeMonitorGetMountForUuid    func(v *VolumeMonitor, uuid string) *Mount
	VolumeMonitorGetMounts          func(v *VolumeMonitor) *T.GList
	VolumeMonitorGetVolumeForUuid   func(v *VolumeMonitor, uuid string) *Volume
	VolumeMonitorGetVolumes         func(v *VolumeMonitor) *T.GList
)

func (v *VolumeMonitor) GetConnectedDrives() *T.GList { return VolumeMonitorGetConnectedDrives(v) }
func (v *VolumeMonitor) GetMountForUuid(uuid string) *Mount {
	return VolumeMonitorGetMountForUuid(v, uuid)
}
func (v *VolumeMonitor) GetMounts() *T.GList { return VolumeMonitorGetMounts(v) }
func (v *VolumeMonitor) GetVolumeForUuid(uuid string) *Volume {
	return VolumeMonitorGetVolumeForUuid(v, uuid)
}
func (v *VolumeMonitor) GetVolumes() *T.GList { return VolumeMonitorGetVolumes(v) }
