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

	KeymapAddVirtualModifiers    func(k *Keymap, state *T.GdkModifierType)
	KeymapGetCapsLockState       func(k *Keymap) bool
	KeymapGetDirection           func(k *Keymap) P.Direction
	KeymapGetEntriesForKeycode   func(k *Keymap, hardwareKeycode uint, keys **KeymapKey, keyvals **uint, nEntries *int) bool
	KeymapGetEntriesForKeyval    func(k *Keymap, keyval uint, keys **KeymapKey, nKeys *int) bool
	KeymapHaveBidiLayouts        func(k *Keymap) bool
	KeymapLookupKey              func(k *Keymap, key *KeymapKey) uint
	KeymapMapVirtualModifiers    func(k *Keymap, state *T.GdkModifierType) bool
	KeymapTranslateKeyboardState func(k *Keymap, hardwareKeycode uint, state T.GdkModifierType, group int, keyval *uint, effectiveGroup, level *int, consumedModifiers *T.GdkModifierType) bool
)

func (k *Keymap) AddVirtualModifiers(state *T.GdkModifierType) { KeymapAddVirtualModifiers(k, state) }
func (k *Keymap) GetCapsLockState() bool                       { return KeymapGetCapsLockState(k) }
func (k *Keymap) GetDirection() P.Direction                    { return KeymapGetDirection(k) }
func (k *Keymap) GetEntriesForKeycode(hardwareKeycode uint, keys **KeymapKey, keyvals **uint, nEntries *int) bool {
	return KeymapGetEntriesForKeycode(k, hardwareKeycode, keys, keyvals, nEntries)
}
func (k *Keymap) GetEntriesForKeyval(keyval uint, keys **KeymapKey, nKeys *int) bool {
	return KeymapGetEntriesForKeyval(k, keyval, keys, nKeys)
}
func (k *Keymap) HaveBidiLayouts() bool         { return KeymapHaveBidiLayouts(k) }
func (k *Keymap) LookupKey(key *KeymapKey) uint { return KeymapLookupKey(k, key) }
func (k *Keymap) MapVirtualModifiers(state *T.GdkModifierType) bool {
	return KeymapMapVirtualModifiers(k, state)
}
func (k *Keymap) TranslateKeyboardState(hardwareKeycode uint, state T.GdkModifierType, group int, keyval *uint, effectiveGroup, level *int, consumedModifiers *T.GdkModifierType) bool {
	return KeymapTranslateKeyboardState(k, hardwareKeycode, state, group, keyval, effectiveGroup, level, consumedModifiers)
}

type KeymapKey struct {
	Keycode uint
	Group   int
	Level   int
}
