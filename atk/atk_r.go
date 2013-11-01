// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
)

var RectangleGetType func() O.Type //TODO(t):Use?

type Registry struct {
	Parent                O.Object
	FactoryTypeRegistry   *L.HashTable
	FactorySingletonCache *L.HashTable
}

var (
	RegistryGetType func() O.Type

	RegistryGetFactory     func(r *Registry, t O.Type) *ObjectFactory
	RegistrySetFactoryType func(r *Registry, t O.Type, factoryType O.Type)
	RegistryGetFactoryType func(r *Registry, t O.Type) O.Type
)

func (r *Registry) GetFactory(t O.Type) *ObjectFactory { return RegistryGetFactory(r, t) }
func (r *Registry) SetFactoryType(t O.Type, factoryType O.Type) {
	RegistrySetFactoryType(r, t, factoryType)
}
func (r *Registry) GetFactoryType(t O.Type) O.Type { return RegistryGetFactoryType(r, t) }

type Relation struct {
	Parent       O.Object
	Target       *L.PtrArray
	Relationship RelationType
}

var (
	RelationGetType func() O.Type
	RelationNew     func(targets **Object, nTargets int, relationship RelationType) *Relation

	RelationAddTarget       func(r *Relation, target *Object)
	RelationGetRelationType func(r *Relation) RelationType
	RelationGetTarget       func(r *Relation) *L.PtrArray
	RelationRemoveTarget    func(r *Relation, target *Object) bool
)

func (r *Relation) AddTarget(target *Object)         { RelationAddTarget(r, target) }
func (r *Relation) GetRelationType() RelationType    { return RelationGetRelationType(r) }
func (r *Relation) GetTarget() *L.PtrArray           { return RelationGetTarget(r) }
func (r *Relation) RemoveTarget(target *Object) bool { return RelationRemoveTarget(r, target) }

type RelationSet struct {
	Parent    O.Object
	Relations *L.PtrArray
}

var (
	RelationSetGetType func() O.Type
	RelationSetNew     func() *RelationSet

	RelationSetAdd               func(r *RelationSet, relation *Relation)
	RelationSetAddRelationByType func(r *RelationSet, relationship RelationType, target *Object)
	RelationSetContains          func(r *RelationSet, relationship RelationType) bool
	RelationSetGetNRelations     func(r *RelationSet) int
	RelationSetGetRelation       func(r *RelationSet, i int) *Relation
	RelationSetGetRelationByType func(r *RelationSet, relationship RelationType) *Relation
	RelationSetRemove            func(r *RelationSet, relation *Relation)
)

func (r *RelationSet) Add(relation *Relation) { RelationSetAdd(r, relation) }
func (r *RelationSet) AddRelationByType(relationship RelationType, target *Object) {
	RelationSetAddRelationByType(r, relationship, target)
}
func (r *RelationSet) Contains(relationship RelationType) bool {
	return RelationSetContains(r, relationship)
}
func (r *RelationSet) GetNRelations() int          { return RelationSetGetNRelations(r) }
func (r *RelationSet) GetRelation(i int) *Relation { return RelationSetGetRelation(r, i) }
func (r *RelationSet) GetRelationByType(relationship RelationType) *Relation {
	return RelationSetGetRelationByType(r, relationship)
}
func (r *RelationSet) Remove(relation *Relation) { RelationSetRemove(r, relation) }

type RelationType Enum

const (
	RELATION_NULL RelationType = iota
	RELATION_CONTROLLED_BY
	RELATION_CONTROLLER_FOR
	RELATION_LABEL_FOR
	RELATION_LABELLED_BY
	RELATION_MEMBER_OF
	RELATION_NODE_CHILD_OF
	RELATION_FLOWS_TO
	RELATION_FLOWS_FROM
	RELATION_SUBWINDOW_OF
	RELATION_EMBEDS
	RELATION_EMBEDDED_BY
	RELATION_POPUP_FOR
	RELATION_PARENT_WINDOW_OF
	RELATION_DESCRIBED_BY
	RELATION_DESCRIPTION_FOR
	RELATION_NODE_PARENT_OF
	RELATION_LAST_DEFINED
)

var (
	RelationTypeGetType func() O.Type

	RelationTypeRegister func(name string) RelationType
	RelationTypeForName  func(name string) RelationType

	RelationTypeGetName func(r RelationType) string
)

func (r RelationType) GetName() string { return RelationTypeGetName(r) }

var (
	RemoveFocusTracker        func(trackerId uint)
	RemoveGlobalEventListener func(listenerId uint)
	RemoveKeyEventListener    func(listenerId uint)
)

type Role Enum

const (
	ROLE_INVALID Role = iota
	ROLE_ACCEL_LABEL
	ROLE_ALERT
	ROLE_ANIMATION
	ROLE_ARROW
	ROLE_CALENDAR
	ROLE_CANVAS
	ROLE_CHECK_BOX
	ROLE_CHECK_MENU_ITEM
	ROLE_COLOR_CHOOSER
	ROLE_COLUMN_HEADER
	ROLE_COMBO_BOX
	ROLE_DATE_EDITOR
	ROLE_DESKTOP_ICON
	ROLE_DESKTOP_FRAME
	ROLE_DIAL
	ROLE_DIALOG
	ROLE_DIRECTORY_PANE
	ROLE_DRAWING_AREA
	ROLE_FILE_CHOOSER
	ROLE_FILLER
	ROLE_FONT_CHOOSER
	ROLE_FRAME
	ROLE_GLASS_PANE
	ROLE_HTML_CONTAINER
	ROLE_ICON
	ROLE_IMAGE
	ROLE_INTERNAL_FRAME
	ROLE_LABEL
	ROLE_LAYERED_PANE
	ROLE_LIST
	ROLE_LIST_ITEM
	ROLE_MENU
	ROLE_MENU_BAR
	ROLE_MENU_ITEM
	ROLE_OPTION_PANE
	ROLE_PAGE_TAB
	ROLE_PAGE_TAB_LIST
	ROLE_PANEL
	ROLE_PASSWORD_TEXT
	ROLE_POPUP_MENU
	ROLE_PROGRESS_BAR
	ROLE_PUSH_BUTTON
	ROLE_RADIO_BUTTON
	ROLE_RADIO_MENU_ITEM
	ROLE_ROOT_PANE
	ROLE_ROW_HEADER
	ROLE_SCROLL_BAR
	ROLE_SCROLL_PANE
	ROLE_SEPARATOR
	ROLE_SLIDER
	ROLE_SPLIT_PANE
	ROLE_SPIN_BUTTON
	ROLE_STATUSBAR
	ROLE_TABLE
	ROLE_TABLE_CELL
	ROLE_TABLE_COLUMN_HEADER
	ROLE_TABLE_ROW_HEADER
	ROLE_TEAR_OFF_MENU_ITEM
	ROLE_TERMINAL
	ROLE_TEXT
	ROLE_TOGGLE_BUTTON
	ROLE_TOOL_BAR
	ROLE_TOOL_TIP
	ROLE_TREE
	ROLE_TREE_TABLE
	ROLE_UNKNOWN
	ROLE_VIEWPORT
	ROLE_WINDOW
	ROLE_HEADER
	ROLE_FOOTER
	ROLE_PARAGRAPH
	ROLE_RULER
	ROLE_APPLICATION
	ROLE_AUTOCOMPLETE
	ROLE_EDITBAR
	ROLE_EMBEDDED
	ROLE_ENTRY
	ROLE_CHART
	ROLE_CAPTION
	ROLE_DOCUMENT_FRAME
	ROLE_HEADING
	ROLE_PAGE
	ROLE_SECTION
	ROLE_REDUNDANT_OBJECT
	ROLE_FORM
	ROLE_LINK
	ROLE_INPUT_METHOD_WINDOW
	ROLE_LAST_DEFINED
)

var (
	RoleGetType func() O.Type

	RoleForName  func(name string) Role
	RoleRegister func(name string) Role

	RoleGetLocalizedName func(r Role) string
	RoleGetName          func(r Role) string
)

func (r Role) GetLocalizedName() string { return RoleGetLocalizedName(r) }
func (r Role) GetName() string          { return RoleGetName(r) }
