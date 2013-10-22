// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type ObjectInitFunc T.GInstanceInitFunc

type OffscreenWindow struct {
	Parent Window
}

var (
	OffscreenWindowGetType func() T.GType
	OffscreenWindowNew     func() *Widget

	offscreenWindowGetPixbuf func(o *OffscreenWindow) *T.GdkPixbuf
	offscreenWindowGetPixmap func(o *OffscreenWindow) *T.GdkPixmap
)

func (o *OffscreenWindow) GetPixbuf() *T.GdkPixbuf { return offscreenWindowGetPixbuf(o) }
func (o *OffscreenWindow) GetPixmap() *T.GdkPixmap { return offscreenWindowGetPixmap(o) }

type OldEditable struct {
	Widget            Widget
	CurrentPos        uint
	SelectionStartPos uint
	SelectionEndPos   uint
	Bits              uint
	// HasSelection : 1
	// Editable : 1
	// Visible : 1
	ClipboardText *T.Gchar
}

var (
	OldEditableGetType func() T.GType

	oldEditableClaimSelection func(o *OldEditable, claim T.Gboolean, time T.GUint32)
	oldEditableChanged        func(o *OldEditable)
)

func (o *OldEditable) ClaimSelection(claim T.Gboolean, time T.GUint32) {
	oldEditableClaimSelection(o, claim, time)
}
func (o *OldEditable) Changed() { oldEditableChanged(o) }

type OptionMenu struct {
	Button   Button
	Menu     *Widget
	MenuItem *Widget
	Width    uint16
	Height   uint16
}

var (
	OptionMenuGetType func() T.GType
	OptionMenuNew     func() *Widget

	optionMenuGetHistory func(o *OptionMenu) int
	optionMenuGetMenu    func(o *OptionMenu) *Widget
	optionMenuRemoveMenu func(o *OptionMenu)
	optionMenuSetHistory func(o *OptionMenu, index uint)
	optionMenuSetMenu    func(o *OptionMenu, menu *Widget)
)

func (o *OptionMenu) GetHistory() int       { return optionMenuGetHistory(o) }
func (o *OptionMenu) GetMenu() *Widget      { return optionMenuGetMenu(o) }
func (o *OptionMenu) RemoveMenu()           { optionMenuRemoveMenu(o) }
func (o *OptionMenu) SetHistory(index uint) { optionMenuSetHistory(o, index) }
func (o *OptionMenu) SetMenu(menu *Widget)  { optionMenuSetMenu(o, menu) }

type Orientable struct{}

type Orientation T.Enum

const (
	ORIENTATION_HORIZONTAL Orientation = iota
	ORIENTATION_VERTICAL
)

var (
	OrientableGetType func() T.GType

	OrientationGetType func() T.GType

	orientableGetOrientation func(o *Orientable) Orientation
	orientableSetOrientation func(o *Orientable, orientation Orientation)
)

func (o *Orientable) GetOrientation() Orientation { return orientableGetOrientation(o) }
func (o *Orientable) SetOrientation(orientation Orientation) {
	orientableSetOrientation(o, orientation)
}
