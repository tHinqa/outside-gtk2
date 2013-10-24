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
	ObjectNewv      func(objectType Type, nParameters uint, parameters *T.GParameter) T.Gpointer
	ObjectNewValist func(objectType Type, firstPropertyName string, varArgs T.VaList) *Object

	ObjectBindProperty             func(source T.Gpointer, sourceProperty string, target T.Gpointer, targetProperty string, flags BindingFlags) *Binding
	ObjectBindPropertyFull         func(source T.Gpointer, sourceProperty string, target T.Gpointer, targetProperty string, flags BindingFlags, transformTo BindingTransformFunc, transformFrom BindingTransformFunc, userData T.Gpointer, notify T.GDestroyNotify) *Binding
	ObjectBindPropertyWithClosures func(source T.Gpointer, sourceProperty string, target T.Gpointer, targetProperty string, flags BindingFlags, transformTo *T.GClosure, transformFrom *T.GClosure) *Binding
	ObjectConnect                  func(object T.Gpointer, signalSpec string, v ...VArg) T.Gpointer
	ObjectDisconnect               func(object T.Gpointer, signalSpec string, v ...VArg)
	ObjectGet                      func(object T.Gpointer, firstPropertyName string, v ...VArg)
	ObjectIsFloating               func(object T.Gpointer) T.Gboolean
	ObjectRef                      func(object T.Gpointer) T.Gpointer
	ObjectRefSink                  func(object T.Gpointer) T.Gpointer
	ObjectSet                      func(object T.Gpointer, firstPropertyName string, v ...VArg)
	ObjectUnref                    func(object T.Gpointer)

	objectAddToggleRef      func(o *Object, notify T.GToggleNotify, data T.Gpointer)
	objectAddWeakPointer    func(o *Object, weakPointerLocation *T.Gpointer)
	objectForceFloating     func(o *Object)
	objectFreezeNotify      func(o *Object)
	objectGetData           func(o *Object, key string) T.Gpointer
	objectGetProperty       func(o *Object, propertyName string, value *Value)
	objectGetQdata          func(o *Object, quark T.GQuark) T.Gpointer
	objectGetValist         func(o *Object, firstPropertyName string, varArgs T.VaList)
	objectNotify            func(o *Object, propertyName string)
	objectNotifyByPspec     func(o *Object, pspec *ParamSpec)
	objectRemoveToggleRef   func(o *Object, notify T.GToggleNotify, data T.Gpointer)
	objectRemoveWeakPointer func(o *Object, weakPointerLocation *T.Gpointer)
	objectRunDispose        func(o *Object)
	objectSetData           func(o *Object, key string, data T.Gpointer)
	objectSetDataFull       func(o *Object, key string, data T.Gpointer, destroy T.GDestroyNotify)
	objectSetProperty       func(o *Object, propertyName string, value *Value)
	objectSetQdata          func(o *Object, quark T.GQuark, data T.Gpointer)
	objectSetQdataFull      func(o *Object, quark T.GQuark, data T.Gpointer, destroy T.GDestroyNotify)
	objectSetValist         func(o *Object, firstPropertyName string, varArgs T.VaList)
	objectStealData         func(o *Object, key string) T.Gpointer
	objectStealQdata        func(o *Object, quark T.GQuark) T.Gpointer
	objectThawNotify        func(o *Object)
	objectWatchClosure      func(o *Object, closure *Closure)
	objectWeakRef           func(o *Object, notify T.GWeakNotify, data T.Gpointer)
	objectWeakUnref         func(o *Object, notify T.GWeakNotify, data T.Gpointer)
)

func (o *Object) AddToggleRef(notify T.GToggleNotify, data T.Gpointer) {
	objectAddToggleRef(o, notify, data)
}
func (o *Object) AddWeakPointer(weakPointerLocation *T.Gpointer) {
	objectAddWeakPointer(o, weakPointerLocation)
}
func (o *Object) ForceFloating()                { objectForceFloating(o) }
func (o *Object) FreezeNotify()                 { objectFreezeNotify(o) }
func (o *Object) GetData(key string) T.Gpointer { return objectGetData(o, key) }
func (o *Object) GetProperty(propertyName string, value *Value) {
	objectGetProperty(o, propertyName, value)
}
func (o *Object) GetQdata(quark T.GQuark) T.Gpointer { return objectGetQdata(o, quark) }
func (o *Object) GetValist(firstPropertyName string, varArgs T.VaList) {
	objectGetValist(o, firstPropertyName, varArgs)
}
func (o *Object) Notify(propertyName string)     { objectNotify(o, propertyName) }
func (o *Object) NotifyByPspec(pspec *ParamSpec) { objectNotifyByPspec(o, pspec) }
func (o *Object) RemoveToggleRef(notify T.GToggleNotify, data T.Gpointer) {
	objectRemoveToggleRef(o, notify, data)
}
func (o *Object) RemoveWeakPointer(weakPointerLocation *T.Gpointer) {
	objectRemoveWeakPointer(o, weakPointerLocation)
}
func (o *Object) RunDispose()                         { objectRunDispose(o) }
func (o *Object) SetData(key string, data T.Gpointer) { objectSetData(o, key, data) }
func (o *Object) SetDataFull(key string, data T.Gpointer, destroy T.GDestroyNotify) {
	objectSetDataFull(o, key, data, destroy)
}
func (o *Object) SetProperty(propertyName string, value *Value) {
	objectSetProperty(o, propertyName, value)
}
func (o *Object) SetQdata(quark T.GQuark, data T.Gpointer) { objectSetQdata(o, quark, data) }
func (o *Object) SetQdataFull(quark T.GQuark, data T.Gpointer, destroy T.GDestroyNotify) {
	objectSetQdataFull(o, quark, data, destroy)
}
func (o *Object) SetValist(firstPropertyName string, varArgs T.VaList) {
	objectSetValist(o, firstPropertyName, varArgs)
}
func (o *Object) StealData(key string) T.Gpointer                 { return objectStealData(o, key) }
func (o *Object) StealQdata(quark T.GQuark) T.Gpointer            { return objectStealQdata(o, quark) }
func (o *Object) ThawNotify()                                     { objectThawNotify(o) }
func (o *Object) WatchClosure(closure *Closure)                   { objectWatchClosure(o, closure) }
func (o *Object) WeakRef(notify T.GWeakNotify, data T.Gpointer)   { objectWeakRef(o, notify, data) }
func (o *Object) WeakUnref(notify T.GWeakNotify, data T.Gpointer) { objectWeakUnref(o, notify, data) }

type ObjectClass struct {
	TypeClass           TypeClass
	ConstructProperties *T.GSList

	Constructor               func(Type Type, nConstructProperties uint, constructProperties *T.GObjectConstructParam) *Object
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
	ObjectClassInstallProperty   func(o *ObjectClass, propertyId uint, pspec *ParamSpec)
	ObjectClassFindProperty      func(o *ObjectClass, propertyName string) *ParamSpec
	ObjectClassListProperties    func(o *ObjectClass, nProperties *uint) **ParamSpec
	ObjectClassOverrideProperty  func(o *ObjectClass, propertyId uint, name string)
	ObjectClassInstallProperties func(o *ObjectClass, nPspecs uint, pspecs **ParamSpec)
)

var ObjectTypeInit func()
