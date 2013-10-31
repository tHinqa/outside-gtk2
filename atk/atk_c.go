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

	ComponentAddFocusHandler      func(c *Component, handler FocusHandler) uint
	ComponentContains             func(c *Component, x, y int, coordType CoordType) bool
	ComponentGetAlpha             func(c *Component) float64
	ComponentGetExtents           func(c *Component, x, y, width, height *int, coordType CoordType)
	ComponentGetLayer             func(c *Component) Layer
	ComponentGetMdiZorder         func(c *Component) int
	ComponentGetPosition          func(c *Component, x, y *int, coordType CoordType)
	ComponentGetSize              func(c *Component, width, height *int)
	ComponentGrabFocus            func(c *Component) bool
	ComponentRefAccessibleAtPoint func(c *Component, x, y int, coordType CoordType) *Object
	ComponentRemoveFocusHandler   func(c *Component, handlerId uint)
	ComponentSetExtents           func(c *Component, x, y, width, height int, coordType CoordType) bool
	ComponentSetPosition          func(c *Component, x, y int, coordType CoordType) bool
	ComponentSetSize              func(c *Component, width, height int) bool
)

func (c *Component) AddFocusHandler(handler FocusHandler) uint {
	return ComponentAddFocusHandler(c, handler)
}
func (c *Component) Contains(x, y int, coordType CoordType) bool {
	return ComponentContains(c, x, y, coordType)
}
func (c *Component) GetAlpha() float64 { return ComponentGetAlpha(c) }
func (c *Component) GetExtents(x, y, width, height *int, coordType CoordType) {
	ComponentGetExtents(c, x, y, width, height, coordType)
}
func (c *Component) GetLayer() Layer   { return ComponentGetLayer(c) }
func (c *Component) GetMdiZorder() int { return ComponentGetMdiZorder(c) }
func (c *Component) GetPosition(x, y *int, coordType CoordType) {
	ComponentGetPosition(c, x, y, coordType)
}
func (c *Component) GetSize(width, height *int) { ComponentGetSize(c, width, height) }
func (c *Component) GrabFocus() bool            { return ComponentGrabFocus(c) }
func (c *Component) RefAccessibleAtPoint(x, y int, coordType CoordType) *Object {
	return ComponentRefAccessibleAtPoint(c, x, y, coordType)
}
func (c *Component) RemoveFocusHandler(handlerId uint) { ComponentRemoveFocusHandler(c, handlerId) }
func (c *Component) SetExtents(x, y, width, height int, coordType CoordType) bool {
	return ComponentSetExtents(c, x, y, width, height, coordType)
}
func (c *Component) SetPosition(x, y int, coordType CoordType) bool {
	return ComponentSetPosition(c, x, y, coordType)
}
func (c *Component) SetSize(width, height int) bool { return ComponentSetSize(c, width, height) }

type CoordType Enum

const (
	XY_SCREEN CoordType = iota
	XY_WINDOW
)

var CoordTypeGetType func() O.Type
