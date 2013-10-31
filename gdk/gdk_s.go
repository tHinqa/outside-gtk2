// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	C "github.com/tHinqa/outside-gtk2/cairo"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Screen struct {
	Parent       O.Object
	Closed       uint    // : 1
	NormalGcs    *[32]GC //TODO(t): CHECK
	ExposureGcs  *[32]GC //TODO(t): CHECK
	SubwindowGcs *[32]GC //TODO(t): CHECK
	FontOptions  *C.FontOptions
	Resolution   float64
}

var (
	ScreenGetType func() O.Type

	ScreenGetDefault func() *Screen

	ScreenBroadcastClientMessage func(s *Screen, event *Event)
	ScreenGetActiveWindow        func(s *Screen) *Window
	ScreenGetDefaultColormap     func(s *Screen) *Colormap
	ScreenGetDisplay             func(s *Screen) *Display
	ScreenGetFontOptions         func(s *Screen) *C.FontOptions
	ScreenGetHeight              func(s *Screen) int
	ScreenGetHeightMm            func(s *Screen) int
	ScreenGetMonitorAtPoint      func(s *Screen, x int, y int) int
	ScreenGetMonitorAtWindow     func(s *Screen, window *Window) int
	ScreenGetMonitorGeometry     func(s *Screen, monitorNum int, dest *Rectangle)
	ScreenGetMonitorHeightMm     func(s *Screen, monitorNum int) int
	ScreenGetMonitorPlugName     func(s *Screen, monitorNum int) string
	ScreenGetMonitorWidthMm      func(s *Screen, monitorNum int) int
	ScreenGetNMonitors           func(s *Screen) int
	ScreenGetNumber              func(s *Screen) int
	ScreenGetPrimaryMonitor      func(s *Screen) int
	ScreenGetResolution          func(s *Screen) float64
	ScreenGetRgbaColormap        func(s *Screen) *Colormap
	ScreenGetRgbaVisual          func(s *Screen) *Visual
	ScreenGetRgbColormap         func(s *Screen) *Colormap
	ScreenGetRgbVisual           func(s *Screen) *Visual
	ScreenGetRootWindow          func(s *Screen) *Window
	ScreenGetSetting             func(s *Screen, name string, value *T.GValue) bool
	ScreenGetSystemColormap      func(s *Screen) *Colormap
	ScreenGetSystemVisual        func(s *Screen) *Visual
	ScreenGetToplevelWindows     func(s *Screen) *T.GList
	ScreenGetWidth               func(s *Screen) int
	ScreenGetWidthMm             func(s *Screen) int
	ScreenGetWindowStack         func(s *Screen) *T.GList
	ScreenIsComposited           func(s *Screen) bool
	ScreenListVisuals            func(s *Screen) *T.GList
	ScreenMakeDisplayName        func(s *Screen) string
	ScreenSetDefaultColormap     func(s *Screen, colormap *Colormap)
	ScreenSetFontOptions         func(s *Screen, options *C.FontOptions)
	ScreenSetResolution          func(s *Screen, dpi float64)
)

func (s *Screen) BroadcastClientMessage(event *Event) { ScreenBroadcastClientMessage(s, event) }
func (s *Screen) GetActiveWindow() *Window            { return ScreenGetActiveWindow(s) }
func (s *Screen) GetDefaultColormap() *Colormap       { return ScreenGetDefaultColormap(s) }
func (s *Screen) GetDisplay() *Display                { return ScreenGetDisplay(s) }
func (s *Screen) GetFontOptions() *C.FontOptions      { return ScreenGetFontOptions(s) }
func (s *Screen) GetHeight() int                      { return ScreenGetHeight(s) }
func (s *Screen) GetHeightMm() int                    { return ScreenGetHeightMm(s) }
func (s *Screen) GetMonitorAtPoint(x, y int) int      { return ScreenGetMonitorAtPoint(s, x, y) }
func (s *Screen) GetMonitorAtWindow(window *Window) int {
	return ScreenGetMonitorAtWindow(s, window)
}
func (s *Screen) GetMonitorGeometry(monitorNum int, dest *Rectangle) {
	ScreenGetMonitorGeometry(s, monitorNum, dest)
}
func (s *Screen) GetMonitorHeightMm(monitorNum int) int {
	return ScreenGetMonitorHeightMm(s, monitorNum)
}
func (s *Screen) GetMonitorPlugName(monitorNum int) string {
	return ScreenGetMonitorPlugName(s, monitorNum)
}
func (s *Screen) GetMonitorWidthMm(monitorNum int) int { return ScreenGetMonitorWidthMm(s, monitorNum) }
func (s *Screen) GetNMonitors() int                    { return ScreenGetNMonitors(s) }
func (s *Screen) GetNumber() int                       { return ScreenGetNumber(s) }
func (s *Screen) GetPrimaryMonitor() int               { return ScreenGetPrimaryMonitor(s) }
func (s *Screen) GetResolution() float64               { return ScreenGetResolution(s) }
func (s *Screen) GetRgbaColormap() *Colormap           { return ScreenGetRgbaColormap(s) }
func (s *Screen) GetRgbaVisual() *Visual               { return ScreenGetRgbaVisual(s) }
func (s *Screen) GetRgbColormap() *Colormap            { return ScreenGetRgbColormap(s) }
func (s *Screen) GetRgbVisual() *Visual                { return ScreenGetRgbVisual(s) }
func (s *Screen) GetRootWindow() *Window               { return ScreenGetRootWindow(s) }
func (s *Screen) GetSetting(name string, value *T.GValue) bool {
	return ScreenGetSetting(s, name, value)
}
func (s *Screen) GetSystemColormap() *Colormap          { return ScreenGetSystemColormap(s) }
func (s *Screen) GetSystemVisual() *Visual              { return ScreenGetSystemVisual(s) }
func (s *Screen) GetToplevelWindows() *T.GList          { return ScreenGetToplevelWindows(s) }
func (s *Screen) GetWidth() int                         { return ScreenGetWidth(s) }
func (s *Screen) GetWidthMm() int                       { return ScreenGetWidthMm(s) }
func (s *Screen) GetWindowStack() *T.GList              { return ScreenGetWindowStack(s) }
func (s *Screen) IsComposited() bool                    { return ScreenIsComposited(s) }
func (s *Screen) ListVisuals() *T.GList                 { return ScreenListVisuals(s) }
func (s *Screen) MakeDisplayName() string               { return ScreenMakeDisplayName(s) }
func (s *Screen) SetDefaultColormap(colormap *Colormap) { ScreenSetDefaultColormap(s, colormap) }
func (s *Screen) SetFontOptions(options *C.FontOptions) { ScreenSetFontOptions(s, options) }
func (s *Screen) SetResolution(dpi float64)             { ScreenSetResolution(s, dpi) }
