// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
	T "github.com/tHinqa/outside-gtk2/types"
)

var DebugResetStaticData func()

type Device struct{}

var (
	deviceAcquire           func(d *Device) Status
	deviceDestroy           func(d *Device)
	deviceFinish            func(d *Device)
	deviceFlush             func(d *Device)
	deviceGetReferenceCount func(d *Device) uint
	deviceGetType           func(d *Device) DeviceType
	deviceGetUserData       func(d *Device, key *UserDataKey) *T.Void
	deviceReference         func(d *Device) *Device
	deviceRelease           func(d *Device)
	deviceSetUserData       func(d *Device, key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status
	deviceStatus            func(d *Device) Status
)

func (d *Device) Acquire() Status                      { return deviceAcquire(d) }
func (d *Device) Destroy()                             { deviceDestroy(d) }
func (d *Device) Finish()                              { deviceFinish(d) }
func (d *Device) Flush()                               { deviceFlush(d) }
func (d *Device) GetReferenceCount() uint              { return deviceGetReferenceCount(d) }
func (d *Device) GetType() DeviceType                  { return deviceGetType(d) }
func (d *Device) GetUserData(key *UserDataKey) *T.Void { return deviceGetUserData(d, key) }
func (d *Device) Reference() *Device                   { return deviceReference(d) }
func (d *Device) Release()                             { deviceRelease(d) }
func (d *Device) SetUserData(key *UserDataKey, userData *T.Void, destroy DestroyFunc) Status {
	return deviceSetUserData(d, key, userData, destroy)
}
func (d *Device) Status() Status { return deviceStatus(d) }

type DeviceType Enum

const (
	DEVICE_TYPE_DRM DeviceType = iota
	DEVICE_TYPE_GL
	DEVICE_TYPE_SCRIPT
	DEVICE_TYPE_XCB
	DEVICE_TYPE_XLIB
	DEVICE_TYPE_XML
)
