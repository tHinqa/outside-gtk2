// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
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
	EventGetType func() T.GType

	EventsPending func() T.Gboolean

	EventGet func() *Event

	EventPeek func() *Event

	EventGetGraphicsExpose func(window *Window) *Event

	EventPut func(event *Event)

	EventNew func(t EventType) *Event

	EventCopy func(event *Event) *Event

	EventFree func(event *Event)

	EventGetTime func(event *Event) T.GUint32

	EventGetState func(
		event *Event,
		state *T.GdkModifierType) T.Gboolean

	EventGetCoords func(
		event *Event,
		xWin,
		yWin *float64) T.Gboolean

	EventGetRootCoords func(
		event *Event,
		xRoot,
		yRoot *float64) T.Gboolean

	EventGetAxis func(
		event *Event,
		axisUse T.GdkAxisUse,
		value *float64) T.Gboolean

	EventRequestMotions func(
		event *T.GdkEventMotion)

	EventHandlerSet func(
		f T.GdkEventFunc,
		data T.Gpointer,
		notify T.GDestroyNotify)

	EventSetScreen func(
		event *Event,
		screen *Screen)

	EventGetScreen func(
		event *Event) *Screen
)

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
