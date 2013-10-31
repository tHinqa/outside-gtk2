// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	I "github.com/tHinqa/outside-gtk2/gio"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type AppLaunchContext struct {
	Parent I.AppLaunchContext
	_      *struct{}
}

var (
	AppLaunchContextGetType func() O.Type
	AppLaunchContextNew     func() *AppLaunchContext

	AppLaunchContextSetDesktop   func(a *AppLaunchContext, desktop int)
	AppLaunchContextSetDisplay   func(a *AppLaunchContext, display *Display)
	AppLaunchContextSetIcon      func(a *AppLaunchContext, icon *I.Icon)
	AppLaunchContextSetIconName  func(a *AppLaunchContext, iconName string)
	AppLaunchContextSetScreen    func(a *AppLaunchContext, screen *Screen)
	AppLaunchContextSetTimestamp func(a *AppLaunchContext, timestamp T.GUint32)
)

func (a *AppLaunchContext) SetDesktop(desktop int)      { AppLaunchContextSetDesktop(a, desktop) }
func (a *AppLaunchContext) SetDisplay(display *Display) { AppLaunchContextSetDisplay(a, display) }
func (a *AppLaunchContext) SetIcon(icon *I.Icon)        { AppLaunchContextSetIcon(a, icon) }
func (a *AppLaunchContext) SetIconName(iconName string) { AppLaunchContextSetIconName(a, iconName) }
func (a *AppLaunchContext) SetScreen(screen *Screen)    { AppLaunchContextSetScreen(a, screen) }
func (a *AppLaunchContext) SetTimestamp(timestamp T.GUint32) {
	AppLaunchContextSetTimestamp(a, timestamp)
}

type Atom *struct{}
