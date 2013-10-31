// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Object struct {
	Parent T.GInitiallyUnowned
	Flags  T.GUint32 //TODO(t): ObjectFlags?
}

var (
	ObjectGetType func() O.Type
	ObjectNew     func(t O.Type, firstPropertyName string, v ...VArg) *Object

	ObjectAddArgType func(argName string, argType O.Type, argFlags, argId uint)

	ObjectDestroy            func(o *Object)
	ObjectGet                func(o *Object, firstPropertyName string, v ...VArg)
	ObjectGetData            func(o *Object, key string) T.Gpointer
	ObjectGetDataById        func(o *Object, dataId T.GQuark) T.Gpointer
	ObjectGetUserData        func(o *Object) T.Gpointer
	ObjectRef                func(o *Object) *Object
	ObjectRemoveData         func(o *Object, key string)
	ObjectRemoveDataById     func(o *Object, dataId T.GQuark)
	ObjectRemoveNoNotify     func(o *Object, key string)
	ObjectRemoveNoNotifyById func(o *Object, keyId T.GQuark)
	ObjectSet                func(o *Object, firstPropertyName string, v ...VArg)
	ObjectSetData            func(o *Object, key string, data T.Gpointer)
	ObjectSetDataById        func(o *Object, dataId T.GQuark, data T.Gpointer)
	ObjectSetDataByIdFull    func(o *Object, dataId T.GQuark, data T.Gpointer, destroy T.GDestroyNotify)
	ObjectSetDataFull        func(o *Object, key string, data T.Gpointer, destroy T.GDestroyNotify)
	ObjectSetUserData        func(o *Object, data T.Gpointer)
	ObjectSink               func(o *Object)
	ObjectUnref              func(o *Object)
	ObjectWeakref            func(o *Object, notify T.GDestroyNotify, data T.Gpointer)
	ObjectWeakunref          func(o *Object, notify T.GDestroyNotify, data T.Gpointer)
)

func (o *Object) Destroy()                                     { ObjectDestroy(o) }
func (o *Object) Get(firstPropertyName string, v ...VArg)      { ObjectGet(o, firstPropertyName, v) }
func (o *Object) GetData(key string) T.Gpointer                { return ObjectGetData(o, key) }
func (o *Object) GetDataById(dataId T.GQuark) T.Gpointer       { return ObjectGetDataById(o, dataId) }
func (o *Object) GetUserData() T.Gpointer                      { return ObjectGetUserData(o) }
func (o *Object) Ref() *Object                                 { return ObjectRef(o) }
func (o *Object) RemoveData(key string)                        { ObjectRemoveData(o, key) }
func (o *Object) RemoveDataById(dataId T.GQuark)               { ObjectRemoveDataById(o, dataId) }
func (o *Object) RemoveNoNotify(key string)                    { ObjectRemoveNoNotify(o, key) }
func (o *Object) RemoveNoNotifyById(keyId T.GQuark)            { ObjectRemoveNoNotifyById(o, keyId) }
func (o *Object) Set(firstPropertyName string, v ...VArg)      { ObjectSet(o, firstPropertyName, v) }
func (o *Object) SetData(key string, data T.Gpointer)          { ObjectSetData(o, key, data) }
func (o *Object) SetDataById(dataId T.GQuark, data T.Gpointer) { ObjectSetDataById(o, dataId, data) }
func (o *Object) SetDataByIdFull(dataId T.GQuark, data T.Gpointer, destroy T.GDestroyNotify) {
	ObjectSetDataByIdFull(o, dataId, data, destroy)
}
func (o *Object) SetDataFull(key string, data T.Gpointer, destroy T.GDestroyNotify) {
	ObjectSetDataFull(o, key, data, destroy)
}
func (o *Object) SetUserData(data T.Gpointer)                      { ObjectSetUserData(o, data) }
func (o *Object) Sink()                                            { ObjectSink(o) }
func (o *Object) Unref()                                           { ObjectUnref(o) }
func (o *Object) Weakref(notify T.GDestroyNotify, data T.Gpointer) { ObjectWeakref(o, notify, data) }
func (o *Object) Weakunref(notify T.GDestroyNotify, data T.Gpointer) {
	ObjectWeakunref(o, notify, data)
}

type ObjectClass struct {
	ParentClass T.GInitiallyUnownedClass
	SetArg      func(object *Object, arg *Arg, argId uint)
	GetArg      func(object *Object, arg *Arg, argId uint)
	Destroy     func(object *Object)
}

type ObjectFlags Enum

const (
	IN_DESTRUCTION ObjectFlags = 1 << iota
	FLOATING
	RESERVED_1
	RESERVED_2
)

var ObjectFlagsGetType func() O.Type

type ObjectInitFunc T.GInstanceInitFunc

type OffscreenWindow struct {
	Parent Window
}

var (
	OffscreenWindowGetType func() O.Type
	OffscreenWindowNew     func() *Widget

	OffscreenWindowGetPixbuf func(o *OffscreenWindow) *D.Pixbuf
	OffscreenWindowGetPixmap func(o *OffscreenWindow) *D.Pixmap
)

func (o *OffscreenWindow) GetPixbuf() *D.Pixbuf { return OffscreenWindowGetPixbuf(o) }
func (o *OffscreenWindow) GetPixmap() *D.Pixmap { return OffscreenWindowGetPixmap(o) }

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
	OldEditableGetType func() O.Type

	OldEditableClaimSelection func(o *OldEditable, claim bool, time T.GUint32)
	OldEditableChanged        func(o *OldEditable)
)

func (o *OldEditable) ClaimSelection(claim bool, time T.GUint32) {
	OldEditableClaimSelection(o, claim, time)
}
func (o *OldEditable) Changed() { OldEditableChanged(o) }

type OptionMenu struct {
	Button   Button
	Menu     *Widget
	MenuItem *Widget
	Width    uint16
	Height   uint16
}

var (
	OptionMenuGetType func() O.Type
	OptionMenuNew     func() *Widget

	OptionMenuGetHistory func(o *OptionMenu) int
	OptionMenuGetMenu    func(o *OptionMenu) *Widget
	OptionMenuRemoveMenu func(o *OptionMenu)
	OptionMenuSetHistory func(o *OptionMenu, index uint)
	OptionMenuSetMenu    func(o *OptionMenu, menu *Widget)
)

func (o *OptionMenu) GetHistory() int       { return OptionMenuGetHistory(o) }
func (o *OptionMenu) GetMenu() *Widget      { return OptionMenuGetMenu(o) }
func (o *OptionMenu) RemoveMenu()           { OptionMenuRemoveMenu(o) }
func (o *OptionMenu) SetHistory(index uint) { OptionMenuSetHistory(o, index) }
func (o *OptionMenu) SetMenu(menu *Widget)  { OptionMenuSetMenu(o, menu) }

type Orientable struct{}

type Orientation Enum

const (
	ORIENTATION_HORIZONTAL Orientation = iota
	ORIENTATION_VERTICAL
)

var (
	OrientableGetType func() O.Type

	OrientationGetType func() O.Type

	OrientableGetOrientation func(o *Orientable) Orientation
	OrientableSetOrientation func(o *Orientable, orientation Orientation)
)

func (o *Orientable) GetOrientation() Orientation { return OrientableGetOrientation(o) }
func (o *Orientable) SetOrientation(orientation Orientation) {
	OrientableSetOrientation(o, orientation)
}
