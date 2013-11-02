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

	ActionDoAction         func(a *Action, i int) bool
	ActionGetDescription   func(a *Action, i int) string
	ActionGetKeybinding    func(a *Action, i int) string
	ActionGetLocalizedName func(a *Action, i int) string
	ActionGetNActions      func(a *Action) int
	ActionGetName          func(a *Action, i int) string
	ActionSetDescription   func(a *Action, i int, desc string) bool
)

func (a *Action) DoAction(i int) bool           { return ActionDoAction(a, i) }
func (a *Action) GetDescription(i int) string   { return ActionGetDescription(a, i) }
func (a *Action) GetKeybinding(i int) string    { return ActionGetKeybinding(a, i) }
func (a *Action) GetLocalizedName(i int) string { return ActionGetLocalizedName(a, i) }
func (a *Action) GetNActions() int              { return ActionGetNActions(a) }
func (a *Action) GetName(i int) string          { return ActionGetName(a, i) }
func (a *Action) SetDescription(i int, desc string) bool {
	return ActionSetDescription(a, i, desc)
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

	AddGlobalEventListener func(listener O.SignalEmissionHook, eventType string) uint

	AddKeyEventListener func(listener KeySnoopFunc, data T.Gpointer) uint

	AttributeSetFree func(attribSet *AttributeSet)
)
