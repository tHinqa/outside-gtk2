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

	objectAddRelationship              func(o *Object, relationship RelationType, target *Object) T.Gboolean
	objectConnectPropertyChangeHandler func(o *Object, handler *PropertyChangeHandler) uint
	objectGetAttributes                func(o *Object) *AttributeSet
	objectGetDescription               func(o *Object) string
	objectGetIndexInParent             func(o *Object) int
	objectGetLayer                     func(o *Object) Layer
	objectGetMdiZorder                 func(o *Object) int
	objectGetNAccessibleChildren       func(o *Object) int
	objectGetName                      func(o *Object) string
	objectGetParent                    func(o *Object) *Object
	objectGetRole                      func(o *Object) Role
	objectInitialize                   func(o *Object, data T.Gpointer)
	objectNotifyStateChange            func(o *Object, state State, value T.Gboolean)
	objectRefAccessibleChild           func(o *Object, i int) *Object
	objectRefRelationSet               func(o *Object) *RelationSet
	objectRefStateSet                  func(o *Object) *StateSet
	objectRemovePropertyChangeHandler  func(o *Object, handlerId uint)
	objectRemoveRelationship           func(o *Object, relationship RelationType, target *Object) T.Gboolean
	objectSetDescription               func(o *Object, description string)
	objectSetName                      func(o *Object, name string)
	objectSetParent                    func(o *Object, parent *Object)
	objectSetRole                      func(o *Object, role Role)
)

func (o *Object) AddRelationship(relationship RelationType, target *Object) T.Gboolean {
	return objectAddRelationship(o, relationship, target)
}
func (o *Object) ConnectPropertyChangeHandler(handler *PropertyChangeHandler) uint {
	return objectConnectPropertyChangeHandler(o, handler)
}
func (o *Object) GetAttributes() *AttributeSet { return objectGetAttributes(o) }
func (o *Object) GetDescription() string       { return objectGetDescription(o) }
func (o *Object) GetIndexInParent() int        { return objectGetIndexInParent(o) }
func (o *Object) GetLayer() Layer              { return objectGetLayer(o) }
func (o *Object) GetMdiZorder() int            { return objectGetMdiZorder(o) }
func (o *Object) GetNAccessibleChildren() int  { return objectGetNAccessibleChildren(o) }
func (o *Object) GetName() string              { return objectGetName(o) }
func (o *Object) GetParent() *Object           { return objectGetParent(o) }
func (o *Object) GetRole() Role                { return objectGetRole(o) }
func (o *Object) Initialize(data T.Gpointer)   { objectInitialize(o, data) }
func (o *Object) NotifyStateChange(state State, value T.Gboolean) {
	objectNotifyStateChange(o, state, value)
}
func (o *Object) RefAccessibleChild(i int) *Object { return objectRefAccessibleChild(o, i) }
func (o *Object) RefRelationSet() *RelationSet     { return objectRefRelationSet(o) }
func (o *Object) RefStateSet() *StateSet           { return objectRefStateSet(o) }
func (o *Object) RemovePropertyChangeHandler(handlerId uint) {
	objectRemovePropertyChangeHandler(o, handlerId)
}
func (o *Object) RemoveRelationship(relationship RelationType, target *Object) T.Gboolean {
	return objectRemoveRelationship(o, relationship, target)
}
func (o *Object) SetDescription(description string) { objectSetDescription(o, description) }
func (o *Object) SetName(name string)               { objectSetName(o, name) }
func (o *Object) SetParent(parent *Object)          { objectSetParent(o, parent) }
func (o *Object) SetRole(role Role)                 { objectSetRole(o, role) }

type ObjectFactory struct{ parent O.Object }

var (
	ObjectFactoryGetType func() O.Type

	objectFactoryCreateAccessible  func(o *ObjectFactory, obj *O.Object) *Object
	objectFactoryGetAccessibleType func(o *ObjectFactory) O.Type
	objectFactoryInvalidate        func(o *ObjectFactory)
)

func (o *ObjectFactory) CreateAccessible(obj *O.Object) *Object {
	return objectFactoryCreateAccessible(o, obj)
}
func (o *ObjectFactory) GetAccessibleType() O.Type { return objectFactoryGetAccessibleType(o) }
func (o *ObjectFactory) Invalidate()               { objectFactoryInvalidate(o) }
