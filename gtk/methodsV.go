// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

var (
	VbuttonBoxGetType func() T.GType
	VbuttonBoxNew     func() *Widget

	VbuttonBoxGetLayoutDefault  func() ButtonBoxStyle
	VbuttonBoxGetSpacingDefault func() int
	VbuttonBoxSetLayoutDefault  func(layout ButtonBoxStyle)
	VbuttonBoxSetSpacingDefault func(spacing int)

	VisibilityGetType func() T.GType

	VolumeButtonGetType func() T.GType
	VolumeButtonNew     func() *Widget

	VpanedGetType func() T.GType
	VpanedNew     func() *Widget

	VrulerGetType func() T.GType
	VrulerNew     func() *Widget

	VscaleGetType      func() T.GType
	VscaleNew          func(adjustment *Adjustment) *Widget
	VscaleNewWithRange func(min, max, step float64) *Widget

	VscrollbarGetType func() T.GType
	VscrollbarNew     func(adjustment *Adjustment) *Widget

	VseparatorGetType func() T.GType
	VseparatorNew     func() *Widget
)

type VBox struct {
	Box Box
}

var (
	VboxGetType func() T.GType
	VboxNew     func(homogeneous T.Gboolean, spacing int) *Widget
)

type Viewport struct {
	Bin         Bin
	ShadowType  ShadowType
	ViewWindow  *T.GdkWindow
	BinWindow   *T.GdkWindow
	Hadjustment *Adjustment
	Vadjustment *Adjustment
}

var (
	ViewportGetType func() T.GType
	ViewportNew     func(hadjustment, vadjustment *Adjustment) *Widget

	viewportGetBinWindow   func(v *Viewport) *T.GdkWindow
	viewportGetHadjustment func(v *Viewport) *Adjustment
	viewportGetShadowType  func(v *Viewport) ShadowType
	viewportGetVadjustment func(v *Viewport) *Adjustment
	viewportGetViewWindow  func(v *Viewport) *T.GdkWindow
	viewportSetHadjustment func(v *Viewport, adjustment *Adjustment)
	viewportSetShadowType  func(v *Viewport, t ShadowType)
	viewportSetVadjustment func(v *Viewport, adjustment *Adjustment)
)

func (v *Viewport) GetBinWindow() *T.GdkWindow   { return viewportGetBinWindow(v) }
func (v *Viewport) GetHadjustment() *Adjustment  { return viewportGetHadjustment(v) }
func (v *Viewport) GetShadowType() ShadowType    { return viewportGetShadowType(v) }
func (v *Viewport) GetVadjustment() *Adjustment  { return viewportGetVadjustment(v) }
func (v *Viewport) GetViewWindow() *T.GdkWindow  { return viewportGetViewWindow(v) }
func (v *Viewport) SetHadjustment(a *Adjustment) { viewportSetHadjustment(v, a) }
func (v *Viewport) SetShadowType(t ShadowType)   { viewportSetShadowType(v, t) }
func (v *Viewport) SetVadjustment(a *Adjustment) { viewportSetVadjustment(v, a) }

type Visibility T.Enum

const (
	VISIBILITY_NONE Visibility = iota
	VISIBILITY_PARTIAL
	VISIBILITY_FULL
)
