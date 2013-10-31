// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Action struct{}

var (
	ActionGetType func() O.Type

	actionDoAction         func(a *Action, i int) bool
	actionGetDescription   func(a *Action, i int) string
	actionGetKeybinding    func(a *Action, i int) string
	actionGetLocalizedName func(a *Action, i int) string
	actionGetNActions      func(a *Action) int
	actionGetName          func(a *Action, i int) string
	actionSetDescription   func(a *Action, i int, desc string) bool
)

func (a *Action) DoAction(i int) bool           { return actionDoAction(a, i) }
func (a *Action) GetDescription(i int) string   { return actionGetDescription(a, i) }
func (a *Action) GetKeybinding(i int) string    { return actionGetKeybinding(a, i) }
func (a *Action) GetLocalizedName(i int) string { return actionGetLocalizedName(a, i) }
func (a *Action) GetNActions() int              { return actionGetNActions(a) }
func (a *Action) GetName(i int) string          { return actionGetName(a, i) }
func (a *Action) SetDescription(i int, desc string) bool {
	return actionSetDescription(a, i, desc)
}

type (
	EventListener func(obj *Object)

	KeyEventStruct struct {
		Type      int
		State     uint
		Keyval    uint
		Length    int
		String    *T.Gchar
		Keycode   uint16
		Timestamp T.GUint32
	}

	KeySnoopFunc func(
		event *KeyEventStruct, funcData T.Gpointer) int
)

var (
	AddFocusTracker func(focusTracker EventListener) uint

	AddGlobalEventListener func(listener T.GSignalEmissionHook, eventType string) uint

	AddKeyEventListener func(listener KeySnoopFunc, data T.Gpointer) uint

	AttributeSetFree func(attribSet *AttributeSet)
)
