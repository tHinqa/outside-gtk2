// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Display struct {
	Parent          T.GObject
	QueuedEvents    *T.GList
	QueuedTail      *T.GList
	ButtonClickTime [2]T.GUint32
	ButtonWindow    [2]*Window
	ButtonNumber    [2]int
	DoubleClickTime uint
	CorePointer     *T.GdkDevice
	PointerHooks    *DisplayPointerHooks
	Bits            uint
	// Closed : 1
	// IgnoreCoreEvents : 1
	DoubleClickDistance uint
	ButtonX             [2]int
	ButtonY             [2]int
	PointerGrabs        *T.GList
	KeyboardGrab        T.GdkKeyboardGrabInfo
	PointerInfo         T.GdkPointerWindowInfo
	LastEventTime       T.GUint32
}

var (
	DisplayGetType func() T.GType

	DisplayGetDefault            func() *Display
	DisplayOpen                  func(displayName string) *Display
	DisplayOpenDefaultLibgtkOnly func() *Display

	displayAddClientMessageFilter        func(d *Display, messageType T.GdkAtom, f T.GdkFilterFunc, data T.Gpointer)
	displayBeep                          func(d *Display)
	displayClose                         func(d *Display)
	displayFlush                         func(d *Display)
	displayGetCorePointer                func(d *Display) *T.GdkDevice
	displayGetDefaultCursorSize          func(d *Display) uint
	displayGetDefaultGroup               func(d *Display) *Window
	displayGetDefaultScreen              func(d *Display) *Screen
	displayGetEvent                      func(d *Display) *Event
	displayGetMaximalCursorSize          func(d *Display, width, height *uint)
	displayGetName                       func(d *Display) string
	displayGetNScreens                   func(d *Display) int
	displayGetPointer                    func(d *Display, screen **Screen, x, y *int, mask *T.GdkModifierType)
	displayGetScreen                     func(d *Display, screenNum int) *Screen
	displayGetWindowAtPointer            func(d *Display, winX, winY *int) *Window
	displayIsClosed                      func(d *Display) T.Gboolean
	displayKeyboardUngrab                func(d *Display, time T.GUint32)
	displayListDevices                   func(d *Display) *T.GList
	displayPeekEvent                     func(d *Display) *Event
	displayPointerIsGrabbed              func(d *Display) T.Gboolean
	displayPointerUngrab                 func(d *Display, time T.GUint32)
	displayPutEvent                      func(d *Display, event *Event)
	displayRequestSelectionNotification  func(d *Display, selection T.GdkAtom) T.Gboolean
	displaySetDoubleClickDistance        func(d *Display, distance uint)
	displaySetDoubleClickTime            func(d *Display, msec uint)
	displaySetPointerHooks               func(d *Display, newHooks *DisplayPointerHooks) *DisplayPointerHooks
	displayStoreClipboard                func(d *Display, clipboardWindow *Window, time T.GUint32, targets *T.GdkAtom, nTargets int)
	displaySupportsClipboardPersistence  func(d *Display) T.Gboolean
	displaySupportsComposite             func(d *Display) T.Gboolean
	displaySupportsCursorAlpha           func(d *Display) T.Gboolean
	displaySupportsCursorColor           func(d *Display) T.Gboolean
	displaySupportsInputShapes           func(d *Display) T.Gboolean
	displaySupportsSelectionNotification func(d *Display) T.Gboolean
	displaySupportsShapes                func(d *Display) T.Gboolean
	displaySync                          func(d *Display)
	displayWarpPointer                   func(d *Display, screen *Screen, x, y int)
)

func (d *Display) AddClientMessageFilter(messageType T.GdkAtom, f T.GdkFilterFunc, data T.Gpointer) {
	displayAddClientMessageFilter(d, messageType, f, data)
}
func (d *Display) Beep()                        { displayBeep(d) }
func (d *Display) Close()                       { displayClose(d) }
func (d *Display) Flush()                       { displayFlush(d) }
func (d *Display) GetCorePointer() *T.GdkDevice { return displayGetCorePointer(d) }
func (d *Display) GetDefaultCursorSize() uint   { return displayGetDefaultCursorSize(d) }
func (d *Display) GetDefaultGroup() *Window     { return displayGetDefaultGroup(d) }
func (d *Display) GetDefaultScreen() *Screen    { return displayGetDefaultScreen(d) }
func (d *Display) GetEvent() *Event             { return displayGetEvent(d) }
func (d *Display) GetMaximalCursorSize(width, height *uint) {
	displayGetMaximalCursorSize(d, width, height)
}
func (d *Display) GetName() string  { return displayGetName(d) }
func (d *Display) GetNScreens() int { return displayGetNScreens(d) }
func (d *Display) GetPointer(screen **Screen, x, y *int, mask *T.GdkModifierType) {
	displayGetPointer(d, screen, x, y, mask)
}
func (d *Display) GetScreen(screenNum int) *Screen { return displayGetScreen(d, screenNum) }
func (d *Display) GetWindowAtPointer(winX, winY *int) *Window {
	return displayGetWindowAtPointer(d, winX, winY)
}
func (d *Display) IsClosed() T.Gboolean          { return displayIsClosed(d) }
func (d *Display) KeyboardUngrab(time T.GUint32) { displayKeyboardUngrab(d, time) }
func (d *Display) ListDevices() *T.GList         { return displayListDevices(d) }
func (d *Display) PeekEvent() *Event             { return displayPeekEvent(d) }
func (d *Display) PointerIsGrabbed() T.Gboolean  { return displayPointerIsGrabbed(d) }
func (d *Display) PointerUngrab(time T.GUint32)  { displayPointerUngrab(d, time) }
func (d *Display) PutEvent(event *Event)         { displayPutEvent(d, event) }
func (d *Display) RequestSelectionNotification(selection T.GdkAtom) T.Gboolean {
	return displayRequestSelectionNotification(d, selection)
}
func (d *Display) SetDoubleClickDistance(distance uint) { displaySetDoubleClickDistance(d, distance) }
func (d *Display) SetDoubleClickTime(msec uint)         { displaySetDoubleClickTime(d, msec) }
func (d *Display) SetPointerHooks(newHooks *DisplayPointerHooks) *DisplayPointerHooks {
	return displaySetPointerHooks(d, newHooks)
}
func (d *Display) StoreClipboard(clipboardWindow *Window, time T.GUint32, targets *T.GdkAtom, nTargets int) {
	displayStoreClipboard(d, clipboardWindow, time, targets, nTargets)
}
func (d *Display) SupportsClipboardPersistence() T.Gboolean {
	return displaySupportsClipboardPersistence(d)
}
func (d *Display) SupportsComposite() T.Gboolean   { return displaySupportsComposite(d) }
func (d *Display) SupportsCursorAlpha() T.Gboolean { return displaySupportsCursorAlpha(d) }
func (d *Display) SupportsCursorColor() T.Gboolean { return displaySupportsCursorColor(d) }
func (d *Display) SupportsInputShapes() T.Gboolean { return displaySupportsInputShapes(d) }
func (d *Display) SupportsSelectionNotification() T.Gboolean {
	return displaySupportsSelectionNotification(d)
}
func (d *Display) SupportsShapes() T.Gboolean           { return displaySupportsShapes(d) }
func (d *Display) Sync()                                { displaySync(d) }
func (d *Display) WarpPointer(screen *Screen, x, y int) { displayWarpPointer(d, screen, x, y) }

type DisplayPointerHooks struct {
	GetPointer func(
		display *Display, screen **Screen,
		x, y *int, mask *T.GdkModifierType)
	WindowGetPointer func(
		display *Display, window *Window,
		x, y *int, mask *T.GdkModifierType) *Window
	WindowAtPointer func(
		display *Display, winX, winY *int) *Window
}

type DisplayManager struct{}

var (
	DisplayManagerGetType func() T.GType

	DisplayManagerGet func() *DisplayManager

	displayManagerGetDefaultDisplay func(d *DisplayManager) *Display
	displayManagerListDisplays      func(d *DisplayManager) *T.GSList
	displayManagerSetDefaultDisplay func(d *DisplayManager, display *Display)
)

func (d *DisplayManager) GetDefaultDisplay() *Display { return displayManagerGetDefaultDisplay(d) }
func (d *DisplayManager) ListDisplays() *T.GSList     { return displayManagerListDisplays(d) }
func (d *DisplayManager) SetDefaultDisplay(display *Display) {
	displayManagerSetDefaultDisplay(d, display)
}
