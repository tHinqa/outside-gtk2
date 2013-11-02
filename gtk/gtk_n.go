// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type (
	Notebook struct {
		Container   Container
		CurPage     *NotebookPage
		Children    *L.List
		FirstTab    *L.List
		FocusTab    *L.List
		Menu        *Widget
		EventWindow *D.Window
		Timer       T.GUint32
		TabHborder  uint16
		TabVborder  uint16
		Bits        uint
		// ShowTabs : 1
		// Homogeneous : 1
		// ShowBorder : 1
		// TabPos : 2
		// Scrollable : 1
		// InChild : 3
		// ClickChild : 3
		// Button : 2
		// NeedTimer : 1
		// ChildHasFocus : 1
		// HaveVisibleChild : 1
		// FocusOut : 1
		// HasBeforePrevious : 1
		// HasBeforeNext : 1
		// HasAfterPrevious : 1
		// HasAfterNext : 1
	}

	NotebookPage struct{}

	NotebookWindowCreationFunc func(source *Notebook,
		page *Widget, x, y int, data T.Gpointer) *Notebook
)

var (
	NotebookGetType    func() O.Type
	NotebookNew        func() *Widget
	NotebookTabGetType func() O.Type

	NotebookSetWindowCreationHook func(f NotebookWindowCreationFunc, data T.Gpointer, destroy O.DestroyNotify)

	NotebookAppendPage           func(n *Notebook, child, tabLabel *Widget) int
	NotebookAppendPageMenu       func(n *Notebook, child, tabLabel, menuLabel *Widget) int
	NotebookGetActionWidget      func(n *Notebook, packType PackType) *Widget
	NotebookGetCurrentPage       func(n *Notebook) int
	NotebookGetGroup             func(n *Notebook) T.Gpointer
	NotebookGetGroupId           func(n *Notebook) int
	NotebookGetGroupName         func(n *Notebook) string
	NotebookGetMenuLabel         func(n *Notebook, child *Widget) *Widget
	NotebookGetMenuLabelText     func(n *Notebook, child *Widget) string
	NotebookGetNPages            func(n *Notebook) int
	NotebookGetNthPage           func(n *Notebook, pageNum int) *Widget
	NotebookGetScrollable        func(n *Notebook) bool
	NotebookGetShowBorder        func(n *Notebook) bool
	NotebookGetShowTabs          func(n *Notebook) bool
	NotebookGetTabDetachable     func(n *Notebook, child *Widget) bool
	NotebookGetTabHborder        func(n *Notebook) uint16
	NotebookGetTabLabel          func(n *Notebook, child *Widget) *Widget
	NotebookGetTabLabelText      func(n *Notebook, child *Widget) string
	NotebookGetTabPos            func(n *Notebook) PositionType
	NotebookGetTabReorderable    func(n *Notebook, child *Widget) bool
	NotebookGetTabVborder        func(n *Notebook) uint16
	NotebookInsertPage           func(n *Notebook, child, tabLabel *Widget, position int) int
	NotebookInsertPageMenu       func(n *Notebook, child, tabLabel, menuLabel *Widget, position int) int
	NotebookNextPage             func(n *Notebook)
	NotebookPageNum              func(n *Notebook, child *Widget) int
	NotebookPopupDisable         func(n *Notebook)
	NotebookPopupEnable          func(n *Notebook)
	NotebookPrependPage          func(n *Notebook, child, tabLabel *Widget) int
	NotebookPrependPageMenu      func(n *Notebook, child, tabLabel, menuLabel *Widget) int
	NotebookPrevPage             func(n *Notebook)
	NotebookQueryTabLabelPacking func(n *Notebook, child *Widget, expand, fill *bool, packType *PackType)
	NotebookRemovePage           func(n *Notebook, pageNum int)
	NotebookReorderChild         func(n *Notebook, child *Widget, position int)
	NotebookSetActionWidget      func(n *Notebook, widget *Widget, packType PackType)
	NotebookSetCurrentPage       func(n *Notebook, pageNum int)
	NotebookSetGroup             func(n *Notebook, group T.Gpointer)
	NotebookSetGroupId           func(n *Notebook, groupId int)
	NotebookSetGroupName         func(n *Notebook, groupName string)
	NotebookSetHomogeneousTabs   func(n *Notebook, homogeneous bool)
	NotebookSetMenuLabel         func(n *Notebook, child, menuLabel *Widget)
	NotebookSetMenuLabelText     func(n *Notebook, child *Widget, menuText string)
	NotebookSetScrollable        func(n *Notebook, scrollable bool)
	NotebookSetShowBorder        func(n *Notebook, showBorder bool)
	NotebookSetShowTabs          func(n *Notebook, showTabs bool)
	NotebookSetTabBorder         func(n *Notebook, borderWidth uint)
	NotebookSetTabDetachable     func(n *Notebook, child *Widget, detachable bool)
	NotebookSetTabHborder        func(n *Notebook, tabHborder uint)
	NotebookSetTabLabel          func(n *Notebook, child, tabLabel *Widget)
	NotebookSetTabLabelPacking   func(n *Notebook, child *Widget, expand, fill bool, packType PackType)
	NotebookSetTabLabelText      func(n *Notebook, child *Widget, tabText string)
	NotebookSetTabPos            func(n *Notebook, pos PositionType)
	NotebookSetTabReorderable    func(n *Notebook, child *Widget, reorderable bool)
	NotebookSetTabVborder        func(n *Notebook, tabVborder uint)
)

func (n *Notebook) AppendPage(child, tabLabel *Widget) int {
	return NotebookAppendPage(n, child, tabLabel)
}
func (n *Notebook) AppendPageMenu(child, tabLabel, menuLabel *Widget) int {
	return NotebookAppendPageMenu(n, child, tabLabel, menuLabel)
}
func (n *Notebook) GetActionWidget(packType PackType) *Widget {
	return NotebookGetActionWidget(n, packType)
}
func (n *Notebook) GetCurrentPage() int                   { return NotebookGetCurrentPage(n) }
func (n *Notebook) GetGroup() T.Gpointer                  { return NotebookGetGroup(n) }
func (n *Notebook) GetGroupId() int                       { return NotebookGetGroupId(n) }
func (n *Notebook) GetGroupName() string                  { return NotebookGetGroupName(n) }
func (n *Notebook) GetMenuLabel(child *Widget) *Widget    { return NotebookGetMenuLabel(n, child) }
func (n *Notebook) GetMenuLabelText(child *Widget) string { return NotebookGetMenuLabelText(n, child) }
func (n *Notebook) GetNPages() int                        { return NotebookGetNPages(n) }
func (n *Notebook) GetNthPage(pageNum int) *Widget        { return NotebookGetNthPage(n, pageNum) }
func (n *Notebook) GetScrollable() bool                   { return NotebookGetScrollable(n) }
func (n *Notebook) GetShowBorder() bool                   { return NotebookGetShowBorder(n) }
func (n *Notebook) GetShowTabs() bool                     { return NotebookGetShowTabs(n) }
func (n *Notebook) GetTabDetachable(child *Widget) bool {
	return NotebookGetTabDetachable(n, child)
}
func (n *Notebook) GetTabHborder() uint16                { return NotebookGetTabHborder(n) }
func (n *Notebook) GetTabLabel(child *Widget) *Widget    { return NotebookGetTabLabel(n, child) }
func (n *Notebook) GetTabLabelText(child *Widget) string { return NotebookGetTabLabelText(n, child) }
func (n *Notebook) GetTabPos() PositionType              { return NotebookGetTabPos(n) }
func (n *Notebook) GetTabReorderable(child *Widget) bool {
	return NotebookGetTabReorderable(n, child)
}
func (n *Notebook) GetTabVborder() uint16 { return NotebookGetTabVborder(n) }
func (n *Notebook) InsertPage(child, tabLabel *Widget, position int) int {
	return NotebookInsertPage(n, child, tabLabel, position)
}
func (n *Notebook) InsertPageMenu(child, tabLabel, menuLabel *Widget, position int) int {
	return NotebookInsertPageMenu(n, child, tabLabel, menuLabel, position)
}
func (n *Notebook) NextPage()                 { NotebookNextPage(n) }
func (n *Notebook) PageNum(child *Widget) int { return NotebookPageNum(n, child) }
func (n *Notebook) PopupDisable()             { NotebookPopupDisable(n) }
func (n *Notebook) PopupEnable()              { NotebookPopupEnable(n) }
func (n *Notebook) PrependPage(child, tabLabel *Widget) int {
	return NotebookPrependPage(n, child, tabLabel)
}
func (n *Notebook) PrependPageMenu(child, tabLabel, menuLabel *Widget) int {
	return NotebookPrependPageMenu(n, child, tabLabel, menuLabel)
}
func (n *Notebook) PrevPage() { NotebookPrevPage(n) }
func (n *Notebook) QueryTabLabelPacking(child *Widget, expand, fill *bool, packType *PackType) {
	NotebookQueryTabLabelPacking(n, child, expand, fill, packType)
}
func (n *Notebook) RemovePage(pageNum int) { NotebookRemovePage(n, pageNum) }
func (n *Notebook) ReorderChild(child *Widget, position int) {
	NotebookReorderChild(n, child, position)
}
func (n *Notebook) SetActionWidget(widget *Widget, packType PackType) {
	NotebookSetActionWidget(n, widget, packType)
}
func (n *Notebook) SetCurrentPage(pageNum int)    { NotebookSetCurrentPage(n, pageNum) }
func (n *Notebook) SetGroup(group T.Gpointer)     { NotebookSetGroup(n, group) }
func (n *Notebook) SetGroupId(groupId int)        { NotebookSetGroupId(n, groupId) }
func (n *Notebook) SetGroupName(groupName string) { NotebookSetGroupName(n, groupName) }
func (n *Notebook) SetHomogeneousTabs(homogeneous bool) {
	NotebookSetHomogeneousTabs(n, homogeneous)
}
func (n *Notebook) SetMenuLabel(child, menuLabel *Widget) { NotebookSetMenuLabel(n, child, menuLabel) }
func (n *Notebook) SetMenuLabelText(child *Widget, menuText string) {
	NotebookSetMenuLabelText(n, child, menuText)
}
func (n *Notebook) SetScrollable(scrollable bool) { NotebookSetScrollable(n, scrollable) }
func (n *Notebook) SetShowBorder(showBorder bool) { NotebookSetShowBorder(n, showBorder) }
func (n *Notebook) SetShowTabs(showTabs bool)     { NotebookSetShowTabs(n, showTabs) }
func (n *Notebook) SetTabBorder(borderWidth uint) { NotebookSetTabBorder(n, borderWidth) }
func (n *Notebook) SetTabDetachable(child *Widget, detachable bool) {
	NotebookSetTabDetachable(n, child, detachable)
}
func (n *Notebook) SetTabHborder(tabHborder uint)       { NotebookSetTabHborder(n, tabHborder) }
func (n *Notebook) SetTabLabel(child, tabLabel *Widget) { NotebookSetTabLabel(n, child, tabLabel) }
func (n *Notebook) SetTabLabelPacking(child *Widget, expand, fill bool, packType PackType) {
	NotebookSetTabLabelPacking(n, child, expand, fill, packType)
}
func (n *Notebook) SetTabLabelText(child *Widget, tabText string) {
	NotebookSetTabLabelText(n, child, tabText)
}
func (n *Notebook) SetTabPos(pos PositionType) { NotebookSetTabPos(n, pos) }
func (n *Notebook) SetTabReorderable(child *Widget, reorderable bool) {
	NotebookSetTabReorderable(n, child, reorderable)
}
func (n *Notebook) SetTabVborder(tabVborder uint) { NotebookSetTabVborder(n, tabVborder) }

type NumberUpLayout Enum

const (
	NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM NumberUpLayout = iota
	NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP
	NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM
	NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP
	NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT
	NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT
	NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT
	NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT
)

var NumberUpLayoutGetType func() O.Type
