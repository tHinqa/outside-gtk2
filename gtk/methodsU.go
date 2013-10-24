// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
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

var UnitGetType func() T.GType

type UIManager struct {
	Parent T.GObject
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
	UiManagerGetType func() T.GType
	UiManagerNew     func() *UIManager

	UiManagerItemTypeGetType func() T.GType

	uiManagerSetAddTearoffs    func(u *UIManager, addTearoffs T.Gboolean)
	uiManagerGetAddTearoffs    func(u *UIManager) T.Gboolean
	uiManagerInsertActionGroup func(u *UIManager, actionGroup *ActionGroup, pos int)
	uiManagerRemoveActionGroup func(u *UIManager, actionGroup *ActionGroup)
	uiManagerGetActionGroups   func(u *UIManager) *T.GList
	uiManagerGetAccelGroup     func(u *UIManager) *AccelGroup
	uiManagerGetWidget         func(u *UIManager, path string) *Widget
	uiManagerGetToplevels      func(u *UIManager, types UIManagerItemType) *T.GSList
	uiManagerGetAction         func(u *UIManager, path string) *Action
	uiManagerAddUiFromString   func(u *UIManager, buffer string, length T.Gssize, err **T.GError) uint
	uiManagerAddUiFromFile     func(u *UIManager, filename string, err **T.GError) uint
	uiManagerAddUi             func(u *UIManager, mergeId uint, path, name, action string, t UIManagerItemType, top T.Gboolean)
	uiManagerRemoveUi          func(u *UIManager, mergeId uint)
	uiManagerGetUi             func(u *UIManager) string
	uiManagerEnsureUpdate      func(u *UIManager)
	uiManagerNewMergeId        func(u *UIManager) uint
)

func (u *UIManager) SetAddTearoffs(addTearoffs T.Gboolean) { uiManagerSetAddTearoffs(u, addTearoffs) }
func (u *UIManager) GetAddTearoffs() T.Gboolean            { return uiManagerGetAddTearoffs(u) }
func (u *UIManager) InsertActionGroup(actionGroup *ActionGroup, pos int) {
	uiManagerInsertActionGroup(u, actionGroup, pos)
}
func (u *UIManager) RemoveActionGroup(actionGroup *ActionGroup) {
	uiManagerRemoveActionGroup(u, actionGroup)
}
func (u *UIManager) GetActionGroups() *T.GList     { return uiManagerGetActionGroups(u) }
func (u *UIManager) GetAccelGroup() *AccelGroup    { return uiManagerGetAccelGroup(u) }
func (u *UIManager) GetWidget(path string) *Widget { return uiManagerGetWidget(u, path) }
func (u *UIManager) GetToplevels(types UIManagerItemType) *T.GSList {
	return uiManagerGetToplevels(u, types)
}
func (u *UIManager) GetAction(path string) *Action { return uiManagerGetAction(u, path) }
func (u *UIManager) AddUiFromString(buffer string, length T.Gssize, err **T.GError) uint {
	return uiManagerAddUiFromString(u, buffer, length, err)
}
func (u *UIManager) AddUiFromFile(filename string, err **T.GError) uint {
	return uiManagerAddUiFromFile(u, filename, err)
}
func (u *UIManager) AddUi(mergeId uint, path, name, action string, t UIManagerItemType, top T.Gboolean) {
	uiManagerAddUi(u, mergeId, path, name, action, t, top)
}
func (u *UIManager) RemoveUi(mergeId uint) { uiManagerRemoveUi(u, mergeId) }
func (u *UIManager) GetUi() string         { return uiManagerGetUi(u) }
func (u *UIManager) EnsureUpdate()         { uiManagerEnsureUpdate(u) }
func (u *UIManager) NewMergeId() uint      { return uiManagerNewMergeId(u) }

type UpdateType Enum

const (
	UPDATE_CONTINUOUS UpdateType = iota
	UPDATE_DISCONTINUOUS
	UPDATE_DELAYED
)

var UpdateTypeGetType func() T.GType
