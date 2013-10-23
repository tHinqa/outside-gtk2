// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Object struct {
	Parent T.GInitiallyUnowned
	Flags  T.GUint32 //TODO(t): ObjectFlags?
}

var (
	ObjectGetType func() T.GType
	ObjectNew     func(t T.GType, firstPropertyName string, v ...VArg) *Object

	ObjectAddArgType func(argName string, argType T.GType, argFlags, argId uint)

	objectDestroy            func(o *Object)
	objectGet                func(o *Object, firstPropertyName string, v ...VArg)
	objectGetData            func(o *Object, key string) T.Gpointer
	objectGetDataById        func(o *Object, dataId T.GQuark) T.Gpointer
	objectGetUserData        func(o *Object) T.Gpointer
	objectRef                func(o *Object) *Object
	objectRemoveData         func(o *Object, key string)
	objectRemoveDataById     func(o *Object, dataId T.GQuark)
	objectRemoveNoNotify     func(o *Object, key string)
	objectRemoveNoNotifyById func(o *Object, keyId T.GQuark)
	objectSet                func(o *Object, firstPropertyName string, v ...VArg)
	objectSetData            func(o *Object, key string, data T.Gpointer)
	objectSetDataById        func(o *Object, dataId T.GQuark, data T.Gpointer)
	objectSetDataByIdFull    func(o *Object, dataId T.GQuark, data T.Gpointer, destroy T.GDestroyNotify)
	objectSetDataFull        func(o *Object, key string, data T.Gpointer, destroy T.GDestroyNotify)
	objectSetUserData        func(o *Object, data T.Gpointer)
	objectSink               func(o *Object)
	objectUnref              func(o *Object)
	objectWeakref            func(o *Object, notify T.GDestroyNotify, data T.Gpointer)
	objectWeakunref          func(o *Object, notify T.GDestroyNotify, data T.Gpointer)
)

func (o *Object) Destroy()                                     { objectDestroy(o) }
func (o *Object) Get(firstPropertyName string, v ...VArg)      { objectGet(o, firstPropertyName, v) }
func (o *Object) GetData(key string) T.Gpointer                { return objectGetData(o, key) }
func (o *Object) GetDataById(dataId T.GQuark) T.Gpointer       { return objectGetDataById(o, dataId) }
func (o *Object) GetUserData() T.Gpointer                      { return objectGetUserData(o) }
func (o *Object) Ref() *Object                                 { return objectRef(o) }
func (o *Object) RemoveData(key string)                        { objectRemoveData(o, key) }
func (o *Object) RemoveDataById(dataId T.GQuark)               { objectRemoveDataById(o, dataId) }
func (o *Object) RemoveNoNotify(key string)                    { objectRemoveNoNotify(o, key) }
func (o *Object) RemoveNoNotifyById(keyId T.GQuark)            { objectRemoveNoNotifyById(o, keyId) }
func (o *Object) Set(firstPropertyName string, v ...VArg)      { objectSet(o, firstPropertyName, v) }
func (o *Object) SetData(key string, data T.Gpointer)          { objectSetData(o, key, data) }
func (o *Object) SetDataById(dataId T.GQuark, data T.Gpointer) { objectSetDataById(o, dataId, data) }
func (o *Object) SetDataByIdFull(dataId T.GQuark, data T.Gpointer, destroy T.GDestroyNotify) {
	objectSetDataByIdFull(o, dataId, data, destroy)
}
func (o *Object) SetDataFull(key string, data T.Gpointer, destroy T.GDestroyNotify) {
	objectSetDataFull(o, key, data, destroy)
}
func (o *Object) SetUserData(data T.Gpointer)                      { objectSetUserData(o, data) }
func (o *Object) Sink()                                            { objectSink(o) }
func (o *Object) Unref()                                           { objectUnref(o) }
func (o *Object) Weakref(notify T.GDestroyNotify, data T.Gpointer) { objectWeakref(o, notify, data) }
func (o *Object) Weakunref(notify T.GDestroyNotify, data T.Gpointer) {
	objectWeakunref(o, notify, data)
}

type ObjectClass struct {
	ParentClass T.GInitiallyUnownedClass
	SetArg      func(object *Object, arg *Arg, argId uint)
	GetArg      func(object *Object, arg *Arg, argId uint)
	Destroy     func(object *Object)
}

type ObjectFlags T.Enum

const (
	IN_DESTRUCTION ObjectFlags = 1 << iota
	FLOATING
	RESERVED_1
	RESERVED_2
)

var ObjectFlagsGetType func() T.GType

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
