// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Object struct {
	Parent           O.Object
	Description      *T.Gchar
	Name             *T.Gchar
	AccessibleParent *Object
	Role             Role
	RelationSet      *RelationSet
	Layer            Layer
}

var (
	ObjectGetType func() O.Type

	ObjectAddRelationship              func(o *Object, relationship RelationType, target *Object) bool
	ObjectConnectPropertyChangeHandler func(o *Object, handler *PropertyChangeHandler) uint
	ObjectGetAttributes                func(o *Object) *AttributeSet
	ObjectGetDescription               func(o *Object) string
	ObjectGetIndexInParent             func(o *Object) int
	ObjectGetLayer                     func(o *Object) Layer
	ObjectGetMdiZorder                 func(o *Object) int
	ObjectGetNAccessibleChildren       func(o *Object) int
	ObjectGetName                      func(o *Object) string
	ObjectGetParent                    func(o *Object) *Object
	ObjectGetRole                      func(o *Object) Role
	ObjectInitialize                   func(o *Object, data T.Gpointer)
	ObjectNotifyStateChange            func(o *Object, state State, value bool)
	ObjectRefAccessibleChild           func(o *Object, i int) *Object
	ObjectRefRelationSet               func(o *Object) *RelationSet
	ObjectRefStateSet                  func(o *Object) *StateSet
	ObjectRemovePropertyChangeHandler  func(o *Object, handlerId uint)
	ObjectRemoveRelationship           func(o *Object, relationship RelationType, target *Object) bool
	ObjectSetDescription               func(o *Object, description string)
	ObjectSetName                      func(o *Object, name string)
	ObjectSetParent                    func(o *Object, parent *Object)
	ObjectSetRole                      func(o *Object, role Role)
)

func (o *Object) AddRelationship(relationship RelationType, target *Object) bool {
	return ObjectAddRelationship(o, relationship, target)
}
func (o *Object) ConnectPropertyChangeHandler(handler *PropertyChangeHandler) uint {
	return ObjectConnectPropertyChangeHandler(o, handler)
}
func (o *Object) GetAttributes() *AttributeSet { return ObjectGetAttributes(o) }
func (o *Object) GetDescription() string       { return ObjectGetDescription(o) }
func (o *Object) GetIndexInParent() int        { return ObjectGetIndexInParent(o) }
func (o *Object) GetLayer() Layer              { return ObjectGetLayer(o) }
func (o *Object) GetMdiZorder() int            { return ObjectGetMdiZorder(o) }
func (o *Object) GetNAccessibleChildren() int  { return ObjectGetNAccessibleChildren(o) }
func (o *Object) GetName() string              { return ObjectGetName(o) }
func (o *Object) GetParent() *Object           { return ObjectGetParent(o) }
func (o *Object) GetRole() Role                { return ObjectGetRole(o) }
func (o *Object) Initialize(data T.Gpointer)   { ObjectInitialize(o, data) }
func (o *Object) NotifyStateChange(state State, value bool) {
	ObjectNotifyStateChange(o, state, value)
}
func (o *Object) RefAccessibleChild(i int) *Object { return ObjectRefAccessibleChild(o, i) }
func (o *Object) RefRelationSet() *RelationSet     { return ObjectRefRelationSet(o) }
func (o *Object) RefStateSet() *StateSet           { return ObjectRefStateSet(o) }
func (o *Object) RemovePropertyChangeHandler(handlerId uint) {
	ObjectRemovePropertyChangeHandler(o, handlerId)
}
func (o *Object) RemoveRelationship(relationship RelationType, target *Object) bool {
	return ObjectRemoveRelationship(o, relationship, target)
}
func (o *Object) SetDescription(description string) { ObjectSetDescription(o, description) }
func (o *Object) SetName(name string)               { ObjectSetName(o, name) }
func (o *Object) SetParent(parent *Object)          { ObjectSetParent(o, parent) }
func (o *Object) SetRole(role Role)                 { ObjectSetRole(o, role) }

type ObjectFactory struct{ parent O.Object }

var (
	ObjectFactoryGetType func() O.Type

	ObjectFactoryCreateAccessible  func(o *ObjectFactory, obj *O.Object) *Object
	ObjectFactoryGetAccessibleType func(o *ObjectFactory) O.Type
	ObjectFactoryInvalidate        func(o *ObjectFactory)
)

func (o *ObjectFactory) CreateAccessible(obj *O.Object) *Object {
	return ObjectFactoryCreateAccessible(o, obj)
}
func (o *ObjectFactory) GetAccessibleType() O.Type { return ObjectFactoryGetAccessibleType(o) }
func (o *ObjectFactory) Invalidate()               { ObjectFactoryInvalidate(o) }
