// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Keymap struct {
	Parent  O.Object
	Display *Display
}

var (
	KeymapGetType func() O.Type

	KeymapGetDefault    func() *Keymap
	KeymapGetForDisplay func(display *Display) *Keymap

	keymapAddVirtualModifiers    func(k *Keymap, state *T.GdkModifierType)
	keymapGetCapsLockState       func(k *Keymap) bool
	keymapGetDirection           func(k *Keymap) P.Direction
	keymapGetEntriesForKeycode   func(k *Keymap, hardwareKeycode uint, keys **KeymapKey, keyvals **uint, nEntries *int) bool
	keymapGetEntriesForKeyval    func(k *Keymap, keyval uint, keys **KeymapKey, nKeys *int) bool
	keymapHaveBidiLayouts        func(k *Keymap) bool
	keymapLookupKey              func(k *Keymap, key *KeymapKey) uint
	keymapMapVirtualModifiers    func(k *Keymap, state *T.GdkModifierType) bool
	keymapTranslateKeyboardState func(k *Keymap, hardwareKeycode uint, state T.GdkModifierType, group int, keyval *uint, effectiveGroup, level *int, consumedModifiers *T.GdkModifierType) bool
)

func (k *Keymap) AddVirtualModifiers(state *T.GdkModifierType) { keymapAddVirtualModifiers(k, state) }
func (k *Keymap) GetCapsLockState() bool                       { return keymapGetCapsLockState(k) }
func (k *Keymap) GetDirection() P.Direction                    { return keymapGetDirection(k) }
func (k *Keymap) GetEntriesForKeycode(hardwareKeycode uint, keys **KeymapKey, keyvals **uint, nEntries *int) bool {
	return keymapGetEntriesForKeycode(k, hardwareKeycode, keys, keyvals, nEntries)
}
func (k *Keymap) GetEntriesForKeyval(keyval uint, keys **KeymapKey, nKeys *int) bool {
	return keymapGetEntriesForKeyval(k, keyval, keys, nKeys)
}
func (k *Keymap) HaveBidiLayouts() bool         { return keymapHaveBidiLayouts(k) }
func (k *Keymap) LookupKey(key *KeymapKey) uint { return keymapLookupKey(k, key) }
func (k *Keymap) MapVirtualModifiers(state *T.GdkModifierType) bool {
	return keymapMapVirtualModifiers(k, state)
}
func (k *Keymap) TranslateKeyboardState(hardwareKeycode uint, state T.GdkModifierType, group int, keyval *uint, effectiveGroup, level *int, consumedModifiers *T.GdkModifierType) bool {
	return keymapTranslateKeyboardState(k, hardwareKeycode, state, group, keyval, effectiveGroup, level, consumedModifiers)
}

type KeymapKey struct {
	Keycode uint
	Group   int
	Level   int
}
