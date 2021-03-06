// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Object struct {
	TypeInstance TypeInstance
	RefCount     uint
	Qdata        *T.GData
}

type Type T.Gsize

var (
	ObjectGetType   func() Type
	ObjectNew       func(objectType Type, firstPropertyName string, v ...VArg) T.Gpointer
	ObjectNewv      func(objectType Type, nParameters uint, parameters *Parameter) T.Gpointer
	ObjectNewValist func(objectType Type, firstPropertyName string, varArgs VAList) *Object

	ObjectBindProperty             func(source T.Gpointer, sourceProperty string, target T.Gpointer, targetProperty string, flags BindingFlags) *Binding
	ObjectBindPropertyFull         func(source T.Gpointer, sourceProperty string, target T.Gpointer, targetProperty string, flags BindingFlags, transformTo BindingTransformFunc, transformFrom BindingTransformFunc, userData T.Gpointer, notify DestroyNotify) *Binding
	ObjectBindPropertyWithClosures func(source T.Gpointer, sourceProperty string, target T.Gpointer, targetProperty string, flags BindingFlags, transformTo *Closure, transformFrom *Closure) *Binding
	ObjectConnect                  func(object T.Gpointer, signalSpec string, v ...VArg) T.Gpointer
	ObjectDisconnect               func(object T.Gpointer, signalSpec string, v ...VArg)
	ObjectGet                      func(object T.Gpointer, firstPropertyName string, v ...VArg)
	ObjectIsFloating               func(object T.Gpointer) bool
	ObjectRef                      func(object T.Gpointer) T.Gpointer
	ObjectRefSink                  func(object T.Gpointer) T.Gpointer
	ObjectSet                      func(object T.Gpointer, firstPropertyName string, v ...VArg)
	ObjectUnref                    func(object T.Gpointer)

	ObjectAddToggleRef      func(o *Object, notify T.GToggleNotify, data T.Gpointer)
	ObjectAddWeakPointer    func(o *Object, weakPointerLocation *T.Gpointer)
	ObjectForceFloating     func(o *Object)
	ObjectFreezeNotify      func(o *Object)
	ObjectGetData           func(o *Object, key string) T.Gpointer
	ObjectGetProperty       func(o *Object, propertyName string, value *Value)
	ObjectGetQdata          func(o *Object, quark T.Quark) T.Gpointer
	ObjectGetValist         func(o *Object, firstPropertyName string, varArgs VAList)
	ObjectNotify            func(o *Object, propertyName string)
	ObjectNotifyByPspec     func(o *Object, pspec *ParamSpec)
	ObjectRemoveToggleRef   func(o *Object, notify T.GToggleNotify, data T.Gpointer)
	ObjectRemoveWeakPointer func(o *Object, weakPointerLocation *T.Gpointer)
	ObjectRunDispose        func(o *Object)
	ObjectSetData           func(o *Object, key string, data T.Gpointer)
	ObjectSetDataFull       func(o *Object, key string, data T.Gpointer, destroy DestroyNotify)
	ObjectSetProperty       func(o *Object, propertyName string, value *Value)
	ObjectSetQdata          func(o *Object, quark T.Quark, data T.Gpointer)
	ObjectSetQdataFull      func(o *Object, quark T.Quark, data T.Gpointer, destroy DestroyNotify)
	ObjectSetValist         func(o *Object, firstPropertyName string, varArgs VAList)
	ObjectStealData         func(o *Object, key string) T.Gpointer
	ObjectStealQdata        func(o *Object, quark T.Quark) T.Gpointer
	ObjectThawNotify        func(o *Object)
	ObjectWatchClosure      func(o *Object, closure *Closure)
	ObjectWeakRef           func(o *Object, notify T.GWeakNotify, data T.Gpointer)
	ObjectWeakUnref         func(o *Object, notify T.GWeakNotify, data T.Gpointer)
)

func (o *Object) AddToggleRef(notify T.GToggleNotify, data T.Gpointer) {
	ObjectAddToggleRef(o, notify, data)
}
func (o *Object) AddWeakPointer(weakPointerLocation *T.Gpointer) {
	ObjectAddWeakPointer(o, weakPointerLocation)
}
func (o *Object) ForceFloating()                { ObjectForceFloating(o) }
func (o *Object) FreezeNotify()                 { ObjectFreezeNotify(o) }
func (o *Object) GetData(key string) T.Gpointer { return ObjectGetData(o, key) }
func (o *Object) GetProperty(propertyName string, value *Value) {
	ObjectGetProperty(o, propertyName, value)
}
func (o *Object) GetQdata(quark T.Quark) T.Gpointer { return ObjectGetQdata(o, quark) }
func (o *Object) GetValist(firstPropertyName string, varArgs VAList) {
	ObjectGetValist(o, firstPropertyName, varArgs)
}
func (o *Object) Notify(propertyName string)     { ObjectNotify(o, propertyName) }
func (o *Object) NotifyByPspec(pspec *ParamSpec) { ObjectNotifyByPspec(o, pspec) }
func (o *Object) RemoveToggleRef(notify T.GToggleNotify, data T.Gpointer) {
	ObjectRemoveToggleRef(o, notify, data)
}
func (o *Object) RemoveWeakPointer(weakPointerLocation *T.Gpointer) {
	ObjectRemoveWeakPointer(o, weakPointerLocation)
}
func (o *Object) RunDispose()                         { ObjectRunDispose(o) }
func (o *Object) SetData(key string, data T.Gpointer) { ObjectSetData(o, key, data) }
func (o *Object) SetDataFull(key string, data T.Gpointer, destroy DestroyNotify) {
	ObjectSetDataFull(o, key, data, destroy)
}
func (o *Object) SetProperty(propertyName string, value *Value) {
	ObjectSetProperty(o, propertyName, value)
}
func (o *Object) SetQdata(quark T.Quark, data T.Gpointer) { ObjectSetQdata(o, quark, data) }
func (o *Object) SetQdataFull(quark T.Quark, data T.Gpointer, destroy DestroyNotify) {
	ObjectSetQdataFull(o, quark, data, destroy)
}
func (o *Object) SetValist(firstPropertyName string, varArgs VAList) {
	ObjectSetValist(o, firstPropertyName, varArgs)
}
func (o *Object) StealData(key string) T.Gpointer                 { return ObjectStealData(o, key) }
func (o *Object) StealQdata(quark T.Quark) T.Gpointer             { return ObjectStealQdata(o, quark) }
func (o *Object) ThawNotify()                                     { ObjectThawNotify(o) }
func (o *Object) WatchClosure(closure *Closure)                   { ObjectWatchClosure(o, closure) }
func (o *Object) WeakRef(notify T.GWeakNotify, data T.Gpointer)   { ObjectWeakRef(o, notify, data) }
func (o *Object) WeakUnref(notify T.GWeakNotify, data T.Gpointer) { ObjectWeakUnref(o, notify, data) }

type ObjectClass struct {
	TypeClass           TypeClass
	ConstructProperties *SList

	Constructor               func(Type Type, nConstructProperties uint, constructProperties *ObjectConstructParam) *Object
	SetProperty               func(object *Object, propertyId uint, value Value, pspec *ParamSpec)
	GetProperty               func(object *Object, propertyId uint, value *Value, pspec *ParamSpec)
	Dispose                   func(object *Object)
	Finalize                  func(object *Object)
	DispatchPropertiesChanged func(object *Object, nPspecs uint, pspecs **ParamSpec)
	Notify                    func(object *Object, pspec *ParamSpec)
	Constructed               func(object *Object)
	Flags                     T.Gsize
	Pdummy                    [6]T.Gpointer
}

var (
	ObjectClassFindProperty      func(o *ObjectClass, propertyName string) *ParamSpec
	ObjectClassInstallProperties func(o *ObjectClass, nPspecs uint, pspecs **ParamSpec)
	ObjectClassInstallProperty   func(o *ObjectClass, propertyId uint, pspec *ParamSpec)
	ObjectClassListProperties    func(o *ObjectClass, nProperties *uint) **ParamSpec
	ObjectClassOverrideProperty  func(o *ObjectClass, propertyId uint, name string)
)

func (o *ObjectClass) FindProperty(propertyName string) *ParamSpec {
	return ObjectClassFindProperty(o, propertyName)
}
func (o *ObjectClass) InstallProperties(nPspecs uint, pspecs **ParamSpec) {
	ObjectClassInstallProperties(o, nPspecs, pspecs)
}
func (o *ObjectClass) InstallProperty(propertyId uint, pspec *ParamSpec) {
	ObjectClassInstallProperty(o, propertyId, pspec)
}
func (o *ObjectClass) ListProperties(nProperties *uint) **ParamSpec {
	return ObjectClassListProperties(o, nProperties)
}
func (o *ObjectClass) OverrideProperty(propertyId uint, name string) {
	ObjectClassOverrideProperty(o, propertyId, name)
}

type ObjectConstructParam struct {
	Pspec *ParamSpec
	Value *Value
}

var ObjectTypeInit func()
