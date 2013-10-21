// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type HandleBox struct {
	Bin         Bin
	BinWindow   *T.GdkWindow
	FloatWindow *T.GdkWindow
	ShadowType  T.GtkShadowType
	Bits        uint
	// HandlePosition : 2;
	// FloatWindowMapped : 1;
	// ChildDetached : 1;
	// InDrag : 1;
	// ShrinkOnDetach : 1;
	// SnapEdge : 3 // signed?!
	DeskoffX         int
	DeskoffY         int
	AttachAllocation T.GtkAllocation
	FloatAllocation  T.GtkAllocation
}

var (
	HandleBoxGetType func() T.GType
	HandleBoxNew     func() *Widget

	handleBoxGetChildDetached  func(h *HandleBox) T.Gboolean
	handleBoxGetHandlePosition func(h *HandleBox) T.GtkPositionType
	handleBoxGetShadowType     func(h *HandleBox) T.GtkShadowType
	handleBoxGetSnapEdge       func(h *HandleBox) T.GtkPositionType
	handleBoxSetHandlePosition func(h *HandleBox, position T.GtkPositionType)
	handleBoxSetShadowType     func(h *HandleBox, t T.GtkShadowType)
	handleBoxSetSnapEdge       func(h *HandleBox, edge T.GtkPositionType)
)

func (h *HandleBox) GetChildDetached() T.Gboolean         { return handleBoxGetChildDetached(h) }
func (h *HandleBox) GetHandlePosition() T.GtkPositionType { return handleBoxGetHandlePosition(h) }
func (h *HandleBox) GetShadowType() T.GtkShadowType       { return handleBoxGetShadowType(h) }
func (h *HandleBox) GetSnapEdge() T.GtkPositionType       { return handleBoxGetSnapEdge(h) }
func (h *HandleBox) SetHandlePosition(position T.GtkPositionType) {
	handleBoxSetHandlePosition(h, position)
}
func (h *HandleBox) SetShadowType(t T.GtkShadowType)    { handleBoxSetShadowType(h, t) }
func (h *HandleBox) SetSnapEdge(edge T.GtkPositionType) { handleBoxSetSnapEdge(h, edge) }

type HSV struct {
	Parent Widget
	_      T.Gpointer
}

type HBox struct {
	Box Box
}

var (
	HboxGetType func() T.GType

	HboxNew func(homogeneous T.Gboolean, spacing int) *Widget
)

var (
	HsvGetType func() T.GType
	HsvNew     func() *Widget

	hsvGetColor    func(hsv *HSV, h, s, v *float64)
	hsvGetMetrics  func(hsv *HSV, size *int, ringWidth *int)
	hsvIsAdjusting func(hsv *HSV) T.Gboolean
	hsvSetColor    func(hsv *HSV, h, s, v float64)
	hsvSetMetrics  func(hsv *HSV, size int, ringWidth int)
)

func (hsv *HSV) GetColor(h, s, v *float64)       { hsvGetColor(hsv, h, s, v) }
func (hsv *HSV) GetMetrics(size, ringWidth *int) { hsvGetMetrics(hsv, size, ringWidth) }
func (hsv *HSV) IsAdjusting() T.Gboolean         { return hsvIsAdjusting(hsv) }
func (hsv *HSV) SetColor(h, s, v float64)        { hsvSetColor(hsv, h, s, v) }
func (hsv *HSV) SetMetrics(size, ringWidth int)  { hsvSetMetrics(hsv, size, ringWidth) }
