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
	HandleBoxNew     func() *T.GtkWidget

	HandleBoxGetChildDetached  func(h *HandleBox) T.Gboolean
	HandleBoxGetHandlePosition func(h *HandleBox) T.GtkPositionType
	HandleBoxGetShadowType     func(h *HandleBox) T.GtkShadowType
	HandleBoxGetSnapEdge       func(h *HandleBox) T.GtkPositionType
	HandleBoxSetHandlePosition func(h *HandleBox, position T.GtkPositionType)
	HandleBoxSetShadowType     func(h *HandleBox, t T.GtkShadowType)
	HandleBoxSetSnapEdge       func(h *HandleBox, edge T.GtkPositionType)
)

func (h *HandleBox) SetShadowType(t T.GtkShadowType) {
	HandleBoxSetShadowType(h, t)
}

func (h *HandleBox) GetShadowType() T.GtkShadowType {
	return HandleBoxGetShadowType(h)
}

func (h *HandleBox) SetHandlePosition(position T.GtkPositionType) {
	HandleBoxSetHandlePosition(h, position)
}

func (h *HandleBox) GetHandlePosition() T.GtkPositionType {
	return HandleBoxGetHandlePosition(h)
}

func (h *HandleBox) SetSnapEdge(edge T.GtkPositionType) {
	HandleBoxSetSnapEdge(h, edge)
}

func (h *HandleBox) GetSnapEdge() T.GtkPositionType {
	return HandleBoxGetSnapEdge(h)
}

func (h *HandleBox) GetChildDetached() T.Gboolean {
	return HandleBoxGetChildDetached(h)
}

type HSV struct {
	Parent T.GtkWidget
	_      T.Gpointer
}

var (
	HsvGetType func() T.GType
	HsvNew     func() *T.GtkWidget

	HsvGetColor    func(hsv *HSV, h, s, v *float64)
	HsvGetMetrics  func(hsv *HSV, size *int, ringWidth *int)
	HsvIsAdjusting func(hsv *HSV) T.Gboolean
	HsvSetColor    func(hsv *HSV, h, s, v float64)
	HsvSetMetrics  func(hsv *HSV, size int, ringWidth int)
)

func (hsv *HSV) SetColor(h, s, v float64) {
	HsvSetColor(hsv, h, s, v)
}

func (hsv *HSV) GetColor(h, s, v *float64) {
	HsvGetColor(hsv, h, s, v)
}

func (hsv *HSV) SetMetrics(size, ringWidth int) {
	HsvSetMetrics(hsv, size, ringWidth)
}

func (hsv *HSV) GetMetrics(size, ringWidth *int) {
	HsvGetMetrics(hsv, size, ringWidth)
}

func (hsv *HSV) IsAdjusting() T.Gboolean {
	return HsvIsAdjusting(hsv)
}
