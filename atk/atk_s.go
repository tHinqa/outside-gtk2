// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type StateType Enum

const (
	STATE_INVALID StateType = iota
	STATE_ACTIVE
	STATE_ARMED
	STATE_BUSY
	STATE_CHECKED
	STATE_DEFUNCT
	STATE_EDITABLE
	STATE_ENABLED
	STATE_EXPANDABLE
	STATE_EXPANDED
	STATE_FOCUSABLE
	STATE_FOCUSED
	STATE_HORIZONTAL
	STATE_ICONIFIED
	STATE_MODAL
	STATE_MULTI_LINE
	STATE_MULTISELECTABLE
	STATE_OPAQUE
	STATE_PRESSED
	STATE_RESIZABLE
	STATE_SELECTABLE
	STATE_SELECTED
	STATE_SENSITIVE
	STATE_SHOWING
	STATE_SINGLE_LINE
	STATE_STALE
	STATE_TRANSIENT
	STATE_VERTICAL
	STATE_VISIBLE
	STATE_MANAGES_DESCENDANTS
	STATE_INDETERMINATE
	STATE_TRUNCATED
	STATE_REQUIRED
	STATE_INVALID_ENTRY
	STATE_SUPPORTS_AUTOCOMPLETION
	STATE_SELECTABLE_TEXT
	STATE_DEFAULT
	STATE_ANIMATED
	STATE_VISITED
	STATE_LAST_DEFINED
)

var (
	StateTypeRegister func(name string) StateType
	StateTypeForName  func(name string) StateType

	stateTypeGetName func(s StateType) string
)

func (s StateType) GetName() string { return stateTypeGetName(s) }

type Selection struct{}

var (
	SelectionGetType func() O.Type

	selectionAddSelection       func(s *Selection, i int) bool
	selectionClearSelection     func(s *Selection) bool
	selectionGetSelectionCount  func(s *Selection) int
	selectionIsChildSelected    func(s *Selection, i int) bool
	selectionRefSelection       func(s *Selection, i int) *Object
	selectionRemoveSelection    func(s *Selection, i int) bool
	selectionSelectAllSelection func(s *Selection) bool
)

func (s *Selection) AddSelection(i int) bool    { return selectionAddSelection(s, i) }
func (s *Selection) ClearSelection() bool       { return selectionClearSelection(s) }
func (s *Selection) GetSelectionCount() int     { return selectionGetSelectionCount(s) }
func (s *Selection) IsChildSelected(i int) bool { return selectionIsChildSelected(s, i) }
func (s *Selection) RefSelection(i int) *Object { return selectionRefSelection(s, i) }
func (s *Selection) RemoveSelection(i int) bool { return selectionRemoveSelection(s, i) }
func (s *Selection) SelectAllSelection() bool   { return selectionSelectAllSelection(s) }

type Socket struct{}

var (
	SocketNew func() *Object

	socketEmbed      func(obj *Socket, plugId string)
	socketIsOccupied func(obj *Socket) bool
)

func (s *Socket) Embed(plugId string) { socketEmbed(s, plugId) }
func (s *Socket) IsOccupied() bool    { return socketIsOccupied(s) }

type State struct{}

type StateSet struct{ parent O.Object }

var (
	StateSetGetType func() O.Type
	StateSetNew     func() *StateSet

	stateSetAddState       func(s *StateSet, t StateType) bool
	stateSetAddStates      func(s *StateSet, types *StateType, nTypes int)
	stateSetAndSets        func(s *StateSet, compareSet *StateSet) *StateSet
	stateSetClearStates    func(s *StateSet)
	stateSetContainsState  func(s *StateSet, t StateType) bool
	stateSetContainsStates func(s *StateSet, types *StateType, nTypes int) bool
	stateSetIsEmpty        func(s *StateSet) bool
	stateSetOrSets         func(s *StateSet, compareSet *StateSet) *StateSet
	stateSetRemoveState    func(s *StateSet, t StateType) bool
	stateSetXorSets        func(s *StateSet, compareSet *StateSet) *StateSet
)

func (s *StateSet) AddState(t StateType) bool              { return stateSetAddState(s, t) }
func (s *StateSet) AddStates(types *StateType, nTypes int) { stateSetAddStates(s, types, nTypes) }
func (s *StateSet) AndSets(compareSet *StateSet) *StateSet { return stateSetAndSets(s, compareSet) }
func (s *StateSet) ClearStates()                           { stateSetClearStates(s) }
func (s *StateSet) ContainsState(t StateType) bool         { return stateSetContainsState(s, t) }
func (s *StateSet) ContainsStates(types *StateType, nTypes int) bool {
	return stateSetContainsStates(s, types, nTypes)
}
func (s *StateSet) IsEmpty() bool                          { return stateSetIsEmpty(s) }
func (s *StateSet) OrSets(compareSet *StateSet) *StateSet  { return stateSetOrSets(s, compareSet) }
func (s *StateSet) RemoveState(t StateType) bool           { return stateSetRemoveState(s, t) }
func (s *StateSet) XorSets(compareSet *StateSet) *StateSet { return stateSetXorSets(s, compareSet) }

type StreamableContent struct{}

var (
	StreamableContentGetType func() O.Type

	streamableContentGetMimeType   func(s *StreamableContent, i int) string
	streamableContentGetNMimeTypes func(s *StreamableContent) int
	streamableContentGetStream     func(s *StreamableContent, mimeType string) *T.GIOChannel
	streamableContentGetUri        func(s *StreamableContent, mimeType string) string
)

func (s *StreamableContent) GetMimeType(i int) string { return streamableContentGetMimeType(s, i) }
func (s *StreamableContent) GetNMimeTypes() int       { return streamableContentGetNMimeTypes(s) }
func (s *StreamableContent) GetStream(mimeType string) *T.GIOChannel {
	return streamableContentGetStream(s, mimeType)
}
func (s *StreamableContent) GetUri(mimeType string) string {
	return streamableContentGetUri(s, mimeType)
}

var StateTypeGetType func() O.Type
