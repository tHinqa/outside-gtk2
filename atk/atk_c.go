// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
)

type Component struct{}

var (
	ComponentGetType func() O.Type

	componentAddFocusHandler      func(c *Component, handler FocusHandler) uint
	componentContains             func(c *Component, x, y int, coordType CoordType) bool
	componentGetAlpha             func(c *Component) float64
	componentGetExtents           func(c *Component, x, y, width, height *int, coordType CoordType)
	componentGetLayer             func(c *Component) Layer
	componentGetMdiZorder         func(c *Component) int
	componentGetPosition          func(c *Component, x, y *int, coordType CoordType)
	componentGetSize              func(c *Component, width, height *int)
	componentGrabFocus            func(c *Component) bool
	componentRefAccessibleAtPoint func(c *Component, x, y int, coordType CoordType) *Object
	componentRemoveFocusHandler   func(c *Component, handlerId uint)
	componentSetExtents           func(c *Component, x, y, width, height int, coordType CoordType) bool
	componentSetPosition          func(c *Component, x, y int, coordType CoordType) bool
	componentSetSize              func(c *Component, width, height int) bool
)

func (c *Component) AddFocusHandler(handler FocusHandler) uint {
	return componentAddFocusHandler(c, handler)
}
func (c *Component) Contains(x, y int, coordType CoordType) bool {
	return componentContains(c, x, y, coordType)
}
func (c *Component) GetAlpha() float64 { return componentGetAlpha(c) }
func (c *Component) GetExtents(x, y, width, height *int, coordType CoordType) {
	componentGetExtents(c, x, y, width, height, coordType)
}
func (c *Component) GetLayer() Layer   { return componentGetLayer(c) }
func (c *Component) GetMdiZorder() int { return componentGetMdiZorder(c) }
func (c *Component) GetPosition(x, y *int, coordType CoordType) {
	componentGetPosition(c, x, y, coordType)
}
func (c *Component) GetSize(width, height *int) { componentGetSize(c, width, height) }
func (c *Component) GrabFocus() bool            { return componentGrabFocus(c) }
func (c *Component) RefAccessibleAtPoint(x, y int, coordType CoordType) *Object {
	return componentRefAccessibleAtPoint(c, x, y, coordType)
}
func (c *Component) RemoveFocusHandler(handlerId uint) { componentRemoveFocusHandler(c, handlerId) }
func (c *Component) SetExtents(x, y, width, height int, coordType CoordType) bool {
	return componentSetExtents(c, x, y, width, height, coordType)
}
func (c *Component) SetPosition(x, y int, coordType CoordType) bool {
	return componentSetPosition(c, x, y, coordType)
}
func (c *Component) SetSize(width, height int) bool { return componentSetSize(c, width, height) }

type CoordType Enum

const (
	XY_SCREEN CoordType = iota
	XY_WINDOW
)

var CoordTypeGetType func() O.Type
