// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
	T "github.com/tHinqa/outside-gtk2/types"
)

var DebugResetStaticData func()

type Device struct{}

var (
	DeviceAcquire           func(d *Device) Status
	DeviceDestroy           func(d *Device)
	DeviceFinish            func(d *Device)
	DeviceFlush             func(d *Device)
	DeviceGetReferenceCount func(d *Device) uint
	DeviceGetType           func(d *Device) DeviceType
	DeviceGetUserData       func(d *Device, key *UserDataKey) *T.Void
	DeviceReference         func(d *Device) *Device
	DeviceRelease           func(d *Device)
	DeviceSetUserData       func(d *Device, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	DeviceStatus            func(d *Device) Status
)

func (d *Device) Acquire() Status                      { return DeviceAcquire(d) }
func (d *Device) Destroy()                             { DeviceDestroy(d) }
func (d *Device) Finish()                              { DeviceFinish(d) }
func (d *Device) Flush()                               { DeviceFlush(d) }
func (d *Device) GetReferenceCount() uint              { return DeviceGetReferenceCount(d) }
func (d *Device) GetType() DeviceType                  { return DeviceGetType(d) }
func (d *Device) GetUserData(key *UserDataKey) *T.Void { return DeviceGetUserData(d, key) }
func (d *Device) Reference() *Device                   { return DeviceReference(d) }
func (d *Device) Release()                             { DeviceRelease(d) }
func (d *Device) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return DeviceSetUserData(d, key, userData, destroy)
}
func (d *Device) Status() Status { return DeviceStatus(d) }

type DeviceType Enum

const (
	DEVICE_TYPE_DRM DeviceType = iota
	DEVICE_TYPE_GL
	DEVICE_TYPE_SCRIPT
	DEVICE_TYPE_XCB
	DEVICE_TYPE_XLIB
	DEVICE_TYPE_XML
)
