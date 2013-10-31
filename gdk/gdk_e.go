// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Event struct {
	Type EventType
	// Union
	// Any  EventAny
	// Expose  EventExpose
	// NoExpose  EventNoExpose
	// Visibility  EventVisibility
	// Motion  EventMotion
	// Button  EventButton
	// Scroll  EventScroll
	// Key  EventKey
	// Crossing  EventCrossing
	// FocusChange  EventFocus
	// Configure  EventConfigure
	// Property  EventProperty
	// Selection  EventSelection
	// OwnerChange  EventOwnerChange
	// Proximity  EventProximity
	// Client  EventClient
	// Dnd  EventDND
	// WindowState  EventWindowState
	// Setting  EventSetting
	// GrabBroken  EventGrabBroken
}

var (
	EventGetType func() O.Type
	EventNew     func(t EventType) *Event

	EventGet                         func() *Event
	EventGetGraphicsExpose           func(window *Window) *Event
	EventHandlerSet                  func(f EventFunc, data T.Gpointer, notify T.GDestroyNotify)
	EventPeek                        func() *Event
	EventRequestMotions              func(event *EventMotion)
	EventSendClientMessageForDisplay func(display *Display, event *Event, winid T.GdkNativeWindow) bool

	EventCopy                   func(event *Event) *Event
	EventFree                   func(event *Event)
	EventGetAxis                func(event *Event, axisUse T.GdkAxisUse, value *float64) bool
	EventGetCoords              func(event *Event, xWin, yWin *float64) bool
	EventGetRootCoords          func(event *Event, xRoot, yRoot *float64) bool
	EventGetScreen              func(event *Event) *Screen
	EventGetState               func(event *Event, state *T.GdkModifierType) bool
	EventGetTime                func(event *Event) T.GUint32
	EventPut                    func(event *Event)
	EventSendClientMessage      func(event *Event, winid T.GdkNativeWindow) bool
	EventSendClientmessageToall func(event *Event)
	EventSetScreen              func(event *Event, screen *Screen)
)

func (e *Event) Copy() *Event { return EventCopy(e) }
func (e *Event) Free()        { EventFree(e) }
func (e *Event) GetAxis(axisUse T.GdkAxisUse, value *float64) bool {
	return EventGetAxis(e, axisUse, value)
}
func (e *Event) GetCoords(xWin, yWin *float64) bool { return EventGetCoords(e, xWin, yWin) }
func (e *Event) GetRootCoords(xRoot, yRoot *float64) bool {
	return EventGetRootCoords(e, xRoot, yRoot)
}
func (e *Event) GetScreen() *Screen                     { return EventGetScreen(e) }
func (e *Event) GetState(state *T.GdkModifierType) bool { return EventGetState(e, state) }
func (e *Event) GetTime() T.GUint32                     { return EventGetTime(e) }
func (e *Event) Put()                                   { EventPut(e) }
func (e *Event) SendClientMessage(winid T.GdkNativeWindow) bool {
	return EventSendClientMessage(e, winid)
}
func (e *Event) SendClientmessageToall()  { EventSendClientmessageToall(e) }
func (e *Event) SetScreen(screen *Screen) { EventSetScreen(e, screen) }

type EventFunc func(event *Event, data T.Gpointer)

type EventAny struct {
	Type      EventType
	Window    *Window
	SendEvent int8
}

type EventButton struct {
	Type         EventType
	Window       *Window
	SendEvent    int8
	Time         T.GUint32
	X            float64
	Y            float64
	Axes         *float64
	State        uint
	Button       uint
	Device       *Device
	XRoot, YRoot float64
}

type EventClient struct {
	Type        EventType
	Window      *Window
	SendEvent   int8
	MessageType Atom
	DataFormat  uint16
	// Union
	B [20]T.Char
	// s [10]T.Short
	// l [5]T.Long
}

type EventConfigure struct {
	Type      EventType
	Window    *Window
	SendEvent int8
	X, Y      int
	Width     int
	Height    int
}

type EventCrossing struct {
	Type       EventType
	Window     *Window
	Send_event int8
	Subwindow  *Window
	Time       T.GUint32
	X          float64
	Y          float64
	XRoot      float64
	YRoot      float64
	Mode       CrossingMode
	Detail     T.GdkNotifyType
	Focus      T.Gboolean
	State      uint
}

type EventExpose struct {
	Type      EventType
	Window    *Window
	SendEvent int8
	Area      Rectangle
	Region    *Region
	Count     int
}

type EventFocus struct {
	Type      EventType
	Window    *Window
	SendEvent int8
	In        int16
}

type EventGrabBroken struct {
	Type       EventType
	Window     *Window
	SendEvent  int8
	Keyboard   T.Gboolean
	Implicit   T.Gboolean
	GrabWindow *Window
}

type EventKey struct {
	Type            EventType
	Window          *Window
	SendEvent       int8
	Time            T.GUint32
	State           uint
	Keyval          uint
	Length          int
	String          *T.Gchar
	HardwareKeycode uint16
	Group           uint8
	IsModifier      uint // : 1
}

type EventMask Enum

const (
	EXPOSURE_MASK EventMask = 1 << (iota + 1)
	POINTER_MOTION_MASK
	POINTER_MOTION_HINT_MASK
	BUTTON_MOTION_MASK
	BUTTON1_MOTION_MASK
	BUTTON2_MOTION_MASK
	BUTTON3_MOTION_MASK
	BUTTON_PRESS_MASK
	BUTTON_RELEASE_MASK
	KEY_PRESS_MASK
	KEY_RELEASE_MASK
	ENTER_NOTIFY_MASK
	LEAVE_NOTIFY_MASK
	FOCUS_CHANGE_MASK
	STRUCTURE_MASK
	PROPERTY_CHANGE_MASK
	VISIBILITY_NOTIFY_MASK
	PROXIMITY_IN_MASK
	PROXIMITY_OUT_MASK
	SUBSTRUCTURE_MASK
	SCROLL_MASK
	ALL_EVENTS_MASK EventMask = 0x3FFFFE
)

var EventMaskGetType func() O.Type

type EventMotion struct {
	Type         EventType
	Window       *Window
	SendEvent    int8
	Time         T.GUint32
	X            float64
	Y            float64
	Axes         *float64
	State        uint
	IsHint       int16
	Device       *Device
	XRoot, YRoot float64
}

type EventProperty struct {
	Type      EventType
	Window    *Window
	SendEvent int8
	Atom      Atom
	Time      T.GUint32
	State     uint
}

type EventProximity struct {
	Type      EventType
	Window    *Window
	SendEvent int8
	Time      T.GUint32
	Device    *Device
}

type EventScroll struct {
	Type         EventType
	Window       *Window
	SendEvent    int8
	Time         T.GUint32
	X            float64
	Y            float64
	State        uint
	Direction    T.GdkScrollDirection
	Device       *Device
	XRoot, YRoot float64
}

var EventsPending func() bool

type EventSelection struct {
	Type      EventType
	Window    *Window
	SendEvent int8
	Selection Atom
	Target    Atom
	Property  Atom
	Time      T.GUint32
	Requestor T.GdkNativeWindow
}

type EventType Enum

const (
	NOTHING EventType = iota - 1
	DELETE
	DESTROY
	EXPOSE
	MOTION_NOTIFY
	BUTTON_PRESS
	X2BUTTON_PRESS
	X3BUTTON_PRESS
	BUTTON_RELEASE
	KEY_PRESS
	KEY_RELEASE
	ENTER_NOTIFY
	LEAVE_NOTIFY
	FOCUS_CHANGE
	CONFIGURE
	MAP
	UNMAP
	PROPERTY_NOTIFY
	SELECTION_CLEAR
	SELECTION_REQUEST
	SELECTION_NOTIFY
	PROXIMITY_IN
	PROXIMITY_OUT
	DRAG_ENTER
	DRAG_LEAVE
	DRAG_MOTION
	DRAG_STATUS
	DROP_START
	DROP_FINISHED
	CLIENT_EVENT
	VISIBILITY_NOTIFY
	NO_EXPOSE
	SCROLL
	WINDOW_STATE
	SETTING
	OWNER_CHANGE
	GRAB_BROKEN
	DAMAGE
	EVENT_LAST
)

var EventTypeGetType func() O.Type

type EventVisibility struct {
	Type      EventType
	Window    *Window
	SendEvent int8
	State     VisibilityState
}

type EventWindowState struct {
	Type           EventType
	Window         *Window
	SendEvent      int8
	ChangedMask    WindowState
	NewWindowState WindowState
}

type ExtensionMode Enum

const (
	EXTENSION_EVENTS_NONE ExtensionMode = iota
	EXTENSION_EVENTS_ALL
	EXTENSION_EVENTS_CURSOR
)

var ExtensionModeGetType func() O.Type
