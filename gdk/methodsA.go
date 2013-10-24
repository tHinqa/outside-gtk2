// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type AppLaunchContext struct {
	Parent T.GAppLaunchContext
	_      *struct{}
}

var (
	AppLaunchContextGetType func() O.Type
	AppLaunchContextNew     func() *AppLaunchContext

	appLaunchContextSetDesktop   func(a *AppLaunchContext, desktop int)
	appLaunchContextSetDisplay   func(a *AppLaunchContext, display *Display)
	appLaunchContextSetIcon      func(a *AppLaunchContext, icon *T.GIcon)
	appLaunchContextSetIconName  func(a *AppLaunchContext, iconName string)
	appLaunchContextSetScreen    func(a *AppLaunchContext, screen *Screen)
	appLaunchContextSetTimestamp func(a *AppLaunchContext, timestamp T.GUint32)
)

func (a *AppLaunchContext) SetDesktop(desktop int)      { appLaunchContextSetDesktop(a, desktop) }
func (a *AppLaunchContext) SetDisplay(display *Display) { appLaunchContextSetDisplay(a, display) }
func (a *AppLaunchContext) SetIcon(icon *T.GIcon)       { appLaunchContextSetIcon(a, icon) }
func (a *AppLaunchContext) SetIconName(iconName string) { appLaunchContextSetIconName(a, iconName) }
func (a *AppLaunchContext) SetScreen(screen *Screen)    { appLaunchContextSetScreen(a, screen) }
func (a *AppLaunchContext) SetTimestamp(timestamp T.GUint32) {
	appLaunchContextSetTimestamp(a, timestamp)
}

type Atom *struct{}
