// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Viewport struct {
	Bin         Bin
	ShadowType  T.GtkShadowType
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
	viewportGetShadowType  func(v *Viewport) T.GtkShadowType
	viewportGetVadjustment func(v *Viewport) *Adjustment
	viewportGetViewWindow  func(v *Viewport) *T.GdkWindow
	viewportSetHadjustment func(v *Viewport, adjustment *Adjustment)
	viewportSetShadowType  func(v *Viewport, t T.GtkShadowType)
	viewportSetVadjustment func(v *Viewport, adjustment *Adjustment)
)

func (v *Viewport) GetBinWindow() *T.GdkWindow            { return viewportGetBinWindow(v) }
func (v *Viewport) GetHadjustment() *Adjustment           { return viewportGetHadjustment(v) }
func (v *Viewport) GetShadowType() T.GtkShadowType        { return viewportGetShadowType(v) }
func (v *Viewport) GetVadjustment() *Adjustment           { return viewportGetVadjustment(v) }
func (v *Viewport) GetViewWindow() *T.GdkWindow           { return viewportGetViewWindow(v) }
func (v *Viewport) SetHadjustment(adjustment *Adjustment) { viewportSetHadjustment(v, adjustment) }
func (v *Viewport) SetShadowType(t T.GtkShadowType)       { viewportSetShadowType(v, t) }
func (v *Viewport) SetVadjustment(adjustment *Adjustment) { viewportSetVadjustment(v, adjustment) }
