// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type HandleBox struct {
	Bin         Bin
	BinWindow   *D.Window
	FloatWindow *D.Window
	ShadowType  ShadowType
	Bits        uint
	// HandlePosition : 2;
	// FloatWindowMapped : 1;
	// ChildDetached : 1;
	// InDrag : 1;
	// ShrinkOnDetach : 1;
	// SnapEdge : 3 // signed?!
	DeskoffX         int
	DeskoffY         int
	AttachAllocation Allocation
	FloatAllocation  Allocation
}

var (
	HandleBoxGetType func() O.Type
	HandleBoxNew     func() *Widget

	handleBoxGetChildDetached  func(h *HandleBox) bool
	handleBoxGetHandlePosition func(h *HandleBox) PositionType
	handleBoxGetShadowType     func(h *HandleBox) ShadowType
	handleBoxGetSnapEdge       func(h *HandleBox) PositionType
	handleBoxSetHandlePosition func(h *HandleBox, position PositionType)
	handleBoxSetShadowType     func(h *HandleBox, t ShadowType)
	handleBoxSetSnapEdge       func(h *HandleBox, edge PositionType)
)

func (h *HandleBox) GetChildDetached() bool          { return handleBoxGetChildDetached(h) }
func (h *HandleBox) GetHandlePosition() PositionType { return handleBoxGetHandlePosition(h) }
func (h *HandleBox) GetShadowType() ShadowType       { return handleBoxGetShadowType(h) }
func (h *HandleBox) GetSnapEdge() PositionType       { return handleBoxGetSnapEdge(h) }
func (h *HandleBox) SetHandlePosition(position PositionType) {
	handleBoxSetHandlePosition(h, position)
}
func (h *HandleBox) SetShadowType(t ShadowType)    { handleBoxSetShadowType(h, t) }
func (h *HandleBox) SetSnapEdge(edge PositionType) { handleBoxSetSnapEdge(h, edge) }

type HBox struct {
	Box Box
}

var (
	HboxGetType func() O.Type

	HboxNew func(homogeneous bool, spacing int) *Widget
)
var (
	HbuttonBoxGetType func() O.Type
	HbuttonBoxNew     func() *Widget

	HbuttonBoxGetSpacingDefault func() int
	HbuttonBoxGetLayoutDefault  func() ButtonBoxStyle
	HbuttonBoxSetSpacingDefault func(spacing int)
	HbuttonBoxSetLayoutDefault  func(layout ButtonBoxStyle)
)
var (
	HpanedGetType func() O.Type
	HpanedNew     func() *Widget
)
var (
	HrulerGetType func() O.Type
	HrulerNew     func() *Widget
)
var (
	HscaleGetType      func() O.Type
	HscaleNew          func(adjustment *Adjustment) *Widget
	HscaleNewWithRange func(min, max, step float64) *Widget
)
var (
	HscrollbarGetType func() O.Type
	HscrollbarNew     func(adjustment *Adjustment) *Widget
)
var (
	HseparatorGetType func() O.Type
	HseparatorNew     func() *Widget
)

type HSV struct {
	Parent Widget
	_      T.Gpointer
}

var (
	HsvGetType func() O.Type
	HsvNew     func() *Widget

	hsvGetColor    func(hsv *HSV, h, s, v *float64)
	hsvGetMetrics  func(hsv *HSV, size *int, ringWidth *int)
	hsvIsAdjusting func(hsv *HSV) bool
	hsvSetColor    func(hsv *HSV, h, s, v float64)
	hsvSetMetrics  func(hsv *HSV, size int, ringWidth int)
)

func (hsv *HSV) GetColor(h, s, v *float64)       { hsvGetColor(hsv, h, s, v) }
func (hsv *HSV) GetMetrics(size, ringWidth *int) { hsvGetMetrics(hsv, size, ringWidth) }
func (hsv *HSV) IsAdjusting() bool               { return hsvIsAdjusting(hsv) }
func (hsv *HSV) SetColor(h, s, v float64)        { hsvSetColor(hsv, h, s, v) }
func (hsv *HSV) SetMetrics(size, ringWidth int)  { hsvSetMetrics(hsv, size, ringWidth) }

var HsvToRgb func(h, s, v float64, r, g, b *float64)
