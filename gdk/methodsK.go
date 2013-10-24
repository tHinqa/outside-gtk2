// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Keymap struct {
	Parent  T.GObject
	Display *Display
}

var (
	KeymapGetType func() T.GType

	KeymapGetDefault    func() *Keymap
	KeymapGetForDisplay func(display *Display) *Keymap

	keymapAddVirtualModifiers    func(k *Keymap, state *T.GdkModifierType)
	keymapGetCapsLockState       func(k *Keymap) T.Gboolean
	keymapGetDirection           func(k *Keymap) T.PangoDirection
	keymapGetEntriesForKeycode   func(k *Keymap, hardwareKeycode uint, keys **KeymapKey, keyvals **uint, nEntries *int) T.Gboolean
	keymapGetEntriesForKeyval    func(k *Keymap, keyval uint, keys **KeymapKey, nKeys *int) T.Gboolean
	keymapHaveBidiLayouts        func(k *Keymap) T.Gboolean
	keymapLookupKey              func(k *Keymap, key *KeymapKey) uint
	keymapMapVirtualModifiers    func(k *Keymap, state *T.GdkModifierType) T.Gboolean
	keymapTranslateKeyboardState func(k *Keymap, hardwareKeycode uint, state T.GdkModifierType, group int, keyval *uint, effectiveGroup, level *int, consumedModifiers *T.GdkModifierType) T.Gboolean
)

func (k *Keymap) AddVirtualModifiers(state *T.GdkModifierType) { keymapAddVirtualModifiers(k, state) }
func (k *Keymap) GetCapsLockState() T.Gboolean                 { return keymapGetCapsLockState(k) }
func (k *Keymap) GetDirection() T.PangoDirection               { return keymapGetDirection(k) }
func (k *Keymap) GetEntriesForKeycode(hardwareKeycode uint, keys **KeymapKey, keyvals **uint, nEntries *int) T.Gboolean {
	return keymapGetEntriesForKeycode(k, hardwareKeycode, keys, keyvals, nEntries)
}
func (k *Keymap) GetEntriesForKeyval(keyval uint, keys **KeymapKey, nKeys *int) T.Gboolean {
	return keymapGetEntriesForKeyval(k, keyval, keys, nKeys)
}
func (k *Keymap) HaveBidiLayouts() T.Gboolean   { return keymapHaveBidiLayouts(k) }
func (k *Keymap) LookupKey(key *KeymapKey) uint { return keymapLookupKey(k, key) }
func (k *Keymap) MapVirtualModifiers(state *T.GdkModifierType) T.Gboolean {
	return keymapMapVirtualModifiers(k, state)
}
func (k *Keymap) TranslateKeyboardState(hardwareKeycode uint, state T.GdkModifierType, group int, keyval *uint, effectiveGroup, level *int, consumedModifiers *T.GdkModifierType) T.Gboolean {
	return keymapTranslateKeyboardState(k, hardwareKeycode, state, group, keyval, effectiveGroup, level, consumedModifiers)
}

type KeymapKey struct {
	Keycode uint
	Group   int
	Level   int
}
