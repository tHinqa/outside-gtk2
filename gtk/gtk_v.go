// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

var (
	VbuttonBoxGetType func() O.Type
	VbuttonBoxNew     func() *Widget

	VbuttonBoxGetLayoutDefault  func() ButtonBoxStyle
	VbuttonBoxGetSpacingDefault func() int
	VbuttonBoxSetLayoutDefault  func(layout ButtonBoxStyle)
	VbuttonBoxSetSpacingDefault func(spacing int)

	VisibilityGetType func() O.Type

	VolumeButtonGetType func() O.Type
	VolumeButtonNew     func() *Widget

	VpanedGetType func() O.Type
	VpanedNew     func() *Widget

	VrulerGetType func() O.Type
	VrulerNew     func() *Widget

	VscaleGetType      func() O.Type
	VscaleNew          func(adjustment *Adjustment) *Widget
	VscaleNewWithRange func(min, max, step float64) *Widget

	VscrollbarGetType func() O.Type
	VscrollbarNew     func(adjustment *Adjustment) *Widget

	VseparatorGetType func() O.Type
	VseparatorNew     func() *Widget
)

type VBox struct {
	Box Box
}

var (
	VboxGetType func() O.Type
	VboxNew     func(homogeneous bool, spacing int) *Widget
)

type Viewport struct {
	Bin         Bin
	ShadowType  ShadowType
	ViewWindow  *D.Window
	BinWindow   *D.Window
	Hadjustment *Adjustment
	Vadjustment *Adjustment
}

var (
	ViewportGetType func() O.Type
	ViewportNew     func(hadjustment, vadjustment *Adjustment) *Widget

	ViewportGetBinWindow   func(v *Viewport) *D.Window
	ViewportGetHadjustment func(v *Viewport) *Adjustment
	ViewportGetShadowType  func(v *Viewport) ShadowType
	ViewportGetVadjustment func(v *Viewport) *Adjustment
	ViewportGetViewWindow  func(v *Viewport) *D.Window
	ViewportSetHadjustment func(v *Viewport, adjustment *Adjustment)
	ViewportSetShadowType  func(v *Viewport, t ShadowType)
	ViewportSetVadjustment func(v *Viewport, adjustment *Adjustment)
)

func (v *Viewport) GetBinWindow() *D.Window      { return ViewportGetBinWindow(v) }
func (v *Viewport) GetHadjustment() *Adjustment  { return ViewportGetHadjustment(v) }
func (v *Viewport) GetShadowType() ShadowType    { return ViewportGetShadowType(v) }
func (v *Viewport) GetVadjustment() *Adjustment  { return ViewportGetVadjustment(v) }
func (v *Viewport) GetViewWindow() *D.Window     { return ViewportGetViewWindow(v) }
func (v *Viewport) SetHadjustment(a *Adjustment) { ViewportSetHadjustment(v, a) }
func (v *Viewport) SetShadowType(t ShadowType)   { ViewportSetShadowType(v, t) }
func (v *Viewport) SetVadjustment(a *Adjustment) { ViewportSetVadjustment(v, a) }

type Visibility Enum

const (
	VISIBILITY_NONE Visibility = iota
	VISIBILITY_PARTIAL
	VISIBILITY_FULL
)
