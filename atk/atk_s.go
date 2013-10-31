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

	StateTypeGetName func(s StateType) string
)

func (s StateType) GetName() string { return StateTypeGetName(s) }

type Selection struct{}

var (
	SelectionGetType func() O.Type

	SelectionAddSelection       func(s *Selection, i int) bool
	SelectionClearSelection     func(s *Selection) bool
	SelectionGetSelectionCount  func(s *Selection) int
	SelectionIsChildSelected    func(s *Selection, i int) bool
	SelectionRefSelection       func(s *Selection, i int) *Object
	SelectionRemoveSelection    func(s *Selection, i int) bool
	SelectionSelectAllSelection func(s *Selection) bool
)

func (s *Selection) AddSelection(i int) bool    { return SelectionAddSelection(s, i) }
func (s *Selection) ClearSelection() bool       { return SelectionClearSelection(s) }
func (s *Selection) GetSelectionCount() int     { return SelectionGetSelectionCount(s) }
func (s *Selection) IsChildSelected(i int) bool { return SelectionIsChildSelected(s, i) }
func (s *Selection) RefSelection(i int) *Object { return SelectionRefSelection(s, i) }
func (s *Selection) RemoveSelection(i int) bool { return SelectionRemoveSelection(s, i) }
func (s *Selection) SelectAllSelection() bool   { return SelectionSelectAllSelection(s) }

type Socket struct{}

var (
	SocketNew func() *Object

	SocketEmbed      func(obj *Socket, plugId string)
	SocketIsOccupied func(obj *Socket) bool
)

func (s *Socket) Embed(plugId string) { SocketEmbed(s, plugId) }
func (s *Socket) IsOccupied() bool    { return SocketIsOccupied(s) }

type State struct{}

type StateSet struct{ parent O.Object }

var (
	StateSetGetType func() O.Type
	StateSetNew     func() *StateSet

	StateSetAddState       func(s *StateSet, t StateType) bool
	StateSetAddStates      func(s *StateSet, types *StateType, nTypes int)
	StateSetAndSets        func(s *StateSet, compareSet *StateSet) *StateSet
	StateSetClearStates    func(s *StateSet)
	StateSetContainsState  func(s *StateSet, t StateType) bool
	StateSetContainsStates func(s *StateSet, types *StateType, nTypes int) bool
	StateSetIsEmpty        func(s *StateSet) bool
	StateSetOrSets         func(s *StateSet, compareSet *StateSet) *StateSet
	StateSetRemoveState    func(s *StateSet, t StateType) bool
	StateSetXorSets        func(s *StateSet, compareSet *StateSet) *StateSet
)

func (s *StateSet) AddState(t StateType) bool              { return StateSetAddState(s, t) }
func (s *StateSet) AddStates(types *StateType, nTypes int) { StateSetAddStates(s, types, nTypes) }
func (s *StateSet) AndSets(compareSet *StateSet) *StateSet { return StateSetAndSets(s, compareSet) }
func (s *StateSet) ClearStates()                           { StateSetClearStates(s) }
func (s *StateSet) ContainsState(t StateType) bool         { return StateSetContainsState(s, t) }
func (s *StateSet) ContainsStates(types *StateType, nTypes int) bool {
	return StateSetContainsStates(s, types, nTypes)
}
func (s *StateSet) IsEmpty() bool                          { return StateSetIsEmpty(s) }
func (s *StateSet) OrSets(compareSet *StateSet) *StateSet  { return StateSetOrSets(s, compareSet) }
func (s *StateSet) RemoveState(t StateType) bool           { return StateSetRemoveState(s, t) }
func (s *StateSet) XorSets(compareSet *StateSet) *StateSet { return StateSetXorSets(s, compareSet) }

type StreamableContent struct{}

var (
	StreamableContentGetType func() O.Type

	StreamableContentGetMimeType   func(s *StreamableContent, i int) string
	StreamableContentGetNMimeTypes func(s *StreamableContent) int
	StreamableContentGetStream     func(s *StreamableContent, mimeType string) *T.GIOChannel
	StreamableContentGetUri        func(s *StreamableContent, mimeType string) string
)

func (s *StreamableContent) GetMimeType(i int) string { return StreamableContentGetMimeType(s, i) }
func (s *StreamableContent) GetNMimeTypes() int       { return StreamableContentGetNMimeTypes(s) }
func (s *StreamableContent) GetStream(mimeType string) *T.GIOChannel {
	return StreamableContentGetStream(s, mimeType)
}
func (s *StreamableContent) GetUri(mimeType string) string {
	return StreamableContentGetUri(s, mimeType)
}

var StateTypeGetType func() O.Type
