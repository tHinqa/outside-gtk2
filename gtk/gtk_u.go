// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Unit Enum

const (
	UNIT_PIXEL Unit = iota
	UNIT_POINTS
	UNIT_INCH
	UNIT_MM
)

var UnitGetType func() O.Type

type UIManager struct {
	Parent O.Object
	_      *struct{}
}

type UIManagerItemType Enum

const (
	UI_MANAGER_MENUBAR UIManagerItemType = 1 << iota
	UI_MANAGER_MENU
	UI_MANAGER_TOOLBAR
	UI_MANAGER_PLACEHOLDER
	UI_MANAGER_POPUP
	UI_MANAGER_MENUITEM
	UI_MANAGER_TOOLITEM
	UI_MANAGER_SEPARATOR
	UI_MANAGER_ACCELERATOR
	UI_MANAGER_POPUP_WITH_ACCELS
	UI_MANAGER_AUTO UIManagerItemType = 0
)

var (
	UiManagerGetType func() O.Type
	UiManagerNew     func() *UIManager

	UiManagerItemTypeGetType func() O.Type

	UiManagerSetAddTearoffs    func(u *UIManager, addTearoffs bool)
	UiManagerGetAddTearoffs    func(u *UIManager) bool
	UiManagerInsertActionGroup func(u *UIManager, actionGroup *ActionGroup, pos int)
	UiManagerRemoveActionGroup func(u *UIManager, actionGroup *ActionGroup)
	UiManagerGetActionGroups   func(u *UIManager) *L.List
	UiManagerGetAccelGroup     func(u *UIManager) *AccelGroup
	UiManagerGetWidget         func(u *UIManager, path string) *Widget
	UiManagerGetToplevels      func(u *UIManager, types UIManagerItemType) *L.SList
	UiManagerGetAction         func(u *UIManager, path string) *Action
	UiManagerAddUiFromString   func(u *UIManager, buffer string, length T.Gssize, err **L.Error) uint
	UiManagerAddUiFromFile     func(u *UIManager, filename string, err **L.Error) uint
	UiManagerAddUi             func(u *UIManager, mergeId uint, path, name, action string, t UIManagerItemType, top bool)
	UiManagerRemoveUi          func(u *UIManager, mergeId uint)
	UiManagerGetUi             func(u *UIManager) string
	UiManagerEnsureUpdate      func(u *UIManager)
	UiManagerNewMergeId        func(u *UIManager) uint
)

func (u *UIManager) SetAddTearoffs(addTearoffs bool) { UiManagerSetAddTearoffs(u, addTearoffs) }
func (u *UIManager) GetAddTearoffs() bool            { return UiManagerGetAddTearoffs(u) }
func (u *UIManager) InsertActionGroup(actionGroup *ActionGroup, pos int) {
	UiManagerInsertActionGroup(u, actionGroup, pos)
}
func (u *UIManager) RemoveActionGroup(actionGroup *ActionGroup) {
	UiManagerRemoveActionGroup(u, actionGroup)
}
func (u *UIManager) GetActionGroups() *L.List      { return UiManagerGetActionGroups(u) }
func (u *UIManager) GetAccelGroup() *AccelGroup    { return UiManagerGetAccelGroup(u) }
func (u *UIManager) GetWidget(path string) *Widget { return UiManagerGetWidget(u, path) }
func (u *UIManager) GetToplevels(types UIManagerItemType) *L.SList {
	return UiManagerGetToplevels(u, types)
}
func (u *UIManager) GetAction(path string) *Action { return UiManagerGetAction(u, path) }
func (u *UIManager) AddUiFromString(buffer string, length T.Gssize, err **L.Error) uint {
	return UiManagerAddUiFromString(u, buffer, length, err)
}
func (u *UIManager) AddUiFromFile(filename string, err **L.Error) uint {
	return UiManagerAddUiFromFile(u, filename, err)
}
func (u *UIManager) AddUi(mergeId uint, path, name, action string, t UIManagerItemType, top bool) {
	UiManagerAddUi(u, mergeId, path, name, action, t, top)
}
func (u *UIManager) RemoveUi(mergeId uint) { UiManagerRemoveUi(u, mergeId) }
func (u *UIManager) GetUi() string         { return UiManagerGetUi(u) }
func (u *UIManager) EnsureUpdate()         { UiManagerEnsureUpdate(u) }
func (u *UIManager) NewMergeId() uint      { return UiManagerNewMergeId(u) }

type UpdateType Enum

const (
	UPDATE_CONTINUOUS UpdateType = iota
	UPDATE_DISCONTINUOUS
	UPDATE_DELAYED
)

var UpdateTypeGetType func() O.Type
