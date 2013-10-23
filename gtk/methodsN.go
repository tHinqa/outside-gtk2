// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type (
	Notebook struct {
		Container   Container
		CurPage     *NotebookPage
		Children    *T.GList
		FirstTab    *T.GList
		FocusTab    *T.GList
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
	NotebookGetType    func() T.GType
	NotebookNew        func() *Widget
	NotebookTabGetType func() T.GType

	NotebookSetWindowCreationHook func(f NotebookWindowCreationFunc, data T.Gpointer, destroy T.GDestroyNotify)

	notebookAppendPage           func(n *Notebook, child, tabLabel *Widget) int
	notebookAppendPageMenu       func(n *Notebook, child, tabLabel, menuLabel *Widget) int
	notebookGetActionWidget      func(n *Notebook, packType PackType) *Widget
	notebookGetCurrentPage       func(n *Notebook) int
	notebookGetGroup             func(n *Notebook) T.Gpointer
	notebookGetGroupId           func(n *Notebook) int
	notebookGetGroupName         func(n *Notebook) string
	notebookGetMenuLabel         func(n *Notebook, child *Widget) *Widget
	notebookGetMenuLabelText     func(n *Notebook, child *Widget) string
	notebookGetNPages            func(n *Notebook) int
	notebookGetNthPage           func(n *Notebook, pageNum int) *Widget
	notebookGetScrollable        func(n *Notebook) T.Gboolean
	notebookGetShowBorder        func(n *Notebook) T.Gboolean
	notebookGetShowTabs          func(n *Notebook) T.Gboolean
	notebookGetTabDetachable     func(n *Notebook, child *Widget) T.Gboolean
	notebookGetTabHborder        func(n *Notebook) uint16
	notebookGetTabLabel          func(n *Notebook, child *Widget) *Widget
	notebookGetTabLabelText      func(n *Notebook, child *Widget) string
	notebookGetTabPos            func(n *Notebook) PositionType
	notebookGetTabReorderable    func(n *Notebook, child *Widget) T.Gboolean
	notebookGetTabVborder        func(n *Notebook) uint16
	notebookInsertPage           func(n *Notebook, child, tabLabel *Widget, position int) int
	notebookInsertPageMenu       func(n *Notebook, child, tabLabel, menuLabel *Widget, position int) int
	notebookNextPage             func(n *Notebook)
	notebookPageNum              func(n *Notebook, child *Widget) int
	notebookPopupDisable         func(n *Notebook)
	notebookPopupEnable          func(n *Notebook)
	notebookPrependPage          func(n *Notebook, child, tabLabel *Widget) int
	notebookPrependPageMenu      func(n *Notebook, child, tabLabel, menuLabel *Widget) int
	notebookPrevPage             func(n *Notebook)
	notebookQueryTabLabelPacking func(n *Notebook, child *Widget, expand, fill *T.Gboolean, packType *PackType)
	notebookRemovePage           func(n *Notebook, pageNum int)
	notebookReorderChild         func(n *Notebook, child *Widget, position int)
	notebookSetActionWidget      func(n *Notebook, widget *Widget, packType PackType)
	notebookSetCurrentPage       func(n *Notebook, pageNum int)
	notebookSetGroup             func(n *Notebook, group T.Gpointer)
	notebookSetGroupId           func(n *Notebook, groupId int)
	notebookSetGroupName         func(n *Notebook, groupName string)
	notebookSetHomogeneousTabs   func(n *Notebook, homogeneous T.Gboolean)
	notebookSetMenuLabel         func(n *Notebook, child, menuLabel *Widget)
	notebookSetMenuLabelText     func(n *Notebook, child *Widget, menuText string)
	notebookSetScrollable        func(n *Notebook, scrollable T.Gboolean)
	notebookSetShowBorder        func(n *Notebook, showBorder T.Gboolean)
	notebookSetShowTabs          func(n *Notebook, showTabs T.Gboolean)
	notebookSetTabBorder         func(n *Notebook, borderWidth uint)
	notebookSetTabDetachable     func(n *Notebook, child *Widget, detachable T.Gboolean)
	notebookSetTabHborder        func(n *Notebook, tabHborder uint)
	notebookSetTabLabel          func(n *Notebook, child, tabLabel *Widget)
	notebookSetTabLabelPacking   func(n *Notebook, child *Widget, expand, fill T.Gboolean, packType PackType)
	notebookSetTabLabelText      func(n *Notebook, child *Widget, tabText string)
	notebookSetTabPos            func(n *Notebook, pos PositionType)
	notebookSetTabReorderable    func(n *Notebook, child *Widget, reorderable T.Gboolean)
	notebookSetTabVborder        func(n *Notebook, tabVborder uint)
)

func (n *Notebook) AppendPage(child, tabLabel *Widget) int {
	return notebookAppendPage(n, child, tabLabel)
}
func (n *Notebook) AppendPageMenu(child, tabLabel, menuLabel *Widget) int {
	return notebookAppendPageMenu(n, child, tabLabel, menuLabel)
}
func (n *Notebook) GetActionWidget(packType PackType) *Widget {
	return notebookGetActionWidget(n, packType)
}
func (n *Notebook) GetCurrentPage() int                   { return notebookGetCurrentPage(n) }
func (n *Notebook) GetGroup() T.Gpointer                  { return notebookGetGroup(n) }
func (n *Notebook) GetGroupId() int                       { return notebookGetGroupId(n) }
func (n *Notebook) GetGroupName() string                  { return notebookGetGroupName(n) }
func (n *Notebook) GetMenuLabel(child *Widget) *Widget    { return notebookGetMenuLabel(n, child) }
func (n *Notebook) GetMenuLabelText(child *Widget) string { return notebookGetMenuLabelText(n, child) }
func (n *Notebook) GetNPages() int                        { return notebookGetNPages(n) }
func (n *Notebook) GetNthPage(pageNum int) *Widget        { return notebookGetNthPage(n, pageNum) }
func (n *Notebook) GetScrollable() T.Gboolean             { return notebookGetScrollable(n) }
func (n *Notebook) GetShowBorder() T.Gboolean             { return notebookGetShowBorder(n) }
func (n *Notebook) GetShowTabs() T.Gboolean               { return notebookGetShowTabs(n) }
func (n *Notebook) GetTabDetachable(child *Widget) T.Gboolean {
	return notebookGetTabDetachable(n, child)
}
func (n *Notebook) GetTabHborder() uint16                { return notebookGetTabHborder(n) }
func (n *Notebook) GetTabLabel(child *Widget) *Widget    { return notebookGetTabLabel(n, child) }
func (n *Notebook) GetTabLabelText(child *Widget) string { return notebookGetTabLabelText(n, child) }
func (n *Notebook) GetTabPos() PositionType              { return notebookGetTabPos(n) }
func (n *Notebook) GetTabReorderable(child *Widget) T.Gboolean {
	return notebookGetTabReorderable(n, child)
}
func (n *Notebook) GetTabVborder() uint16 { return notebookGetTabVborder(n) }
func (n *Notebook) InsertPage(child, tabLabel *Widget, position int) int {
	return notebookInsertPage(n, child, tabLabel, position)
}
func (n *Notebook) InsertPageMenu(child, tabLabel, menuLabel *Widget, position int) int {
	return notebookInsertPageMenu(n, child, tabLabel, menuLabel, position)
}
func (n *Notebook) NextPage()                 { notebookNextPage(n) }
func (n *Notebook) PageNum(child *Widget) int { return notebookPageNum(n, child) }
func (n *Notebook) PopupDisable()             { notebookPopupDisable(n) }
func (n *Notebook) PopupEnable()              { notebookPopupEnable(n) }
func (n *Notebook) PrependPage(child, tabLabel *Widget) int {
	return notebookPrependPage(n, child, tabLabel)
}
func (n *Notebook) PrependPageMenu(child, tabLabel, menuLabel *Widget) int {
	return notebookPrependPageMenu(n, child, tabLabel, menuLabel)
}
func (n *Notebook) PrevPage() { notebookPrevPage(n) }
func (n *Notebook) QueryTabLabelPacking(child *Widget, expand, fill *T.Gboolean, packType *PackType) {
	notebookQueryTabLabelPacking(n, child, expand, fill, packType)
}
func (n *Notebook) RemovePage(pageNum int) { notebookRemovePage(n, pageNum) }
func (n *Notebook) ReorderChild(child *Widget, position int) {
	notebookReorderChild(n, child, position)
}
func (n *Notebook) SetActionWidget(widget *Widget, packType PackType) {
	notebookSetActionWidget(n, widget, packType)
}
func (n *Notebook) SetCurrentPage(pageNum int)    { notebookSetCurrentPage(n, pageNum) }
func (n *Notebook) SetGroup(group T.Gpointer)     { notebookSetGroup(n, group) }
func (n *Notebook) SetGroupId(groupId int)        { notebookSetGroupId(n, groupId) }
func (n *Notebook) SetGroupName(groupName string) { notebookSetGroupName(n, groupName) }
func (n *Notebook) SetHomogeneousTabs(homogeneous T.Gboolean) {
	notebookSetHomogeneousTabs(n, homogeneous)
}
func (n *Notebook) SetMenuLabel(child, menuLabel *Widget) { notebookSetMenuLabel(n, child, menuLabel) }
func (n *Notebook) SetMenuLabelText(child *Widget, menuText string) {
	notebookSetMenuLabelText(n, child, menuText)
}
func (n *Notebook) SetScrollable(scrollable T.Gboolean) { notebookSetScrollable(n, scrollable) }
func (n *Notebook) SetShowBorder(showBorder T.Gboolean) { notebookSetShowBorder(n, showBorder) }
func (n *Notebook) SetShowTabs(showTabs T.Gboolean)     { notebookSetShowTabs(n, showTabs) }
func (n *Notebook) SetTabBorder(borderWidth uint)       { notebookSetTabBorder(n, borderWidth) }
func (n *Notebook) SetTabDetachable(child *Widget, detachable T.Gboolean) {
	notebookSetTabDetachable(n, child, detachable)
}
func (n *Notebook) SetTabHborder(tabHborder uint)       { notebookSetTabHborder(n, tabHborder) }
func (n *Notebook) SetTabLabel(child, tabLabel *Widget) { notebookSetTabLabel(n, child, tabLabel) }
func (n *Notebook) SetTabLabelPacking(child *Widget, expand, fill T.Gboolean, packType PackType) {
	notebookSetTabLabelPacking(n, child, expand, fill, packType)
}
func (n *Notebook) SetTabLabelText(child *Widget, tabText string) {
	notebookSetTabLabelText(n, child, tabText)
}
func (n *Notebook) SetTabPos(pos PositionType) { notebookSetTabPos(n, pos) }
func (n *Notebook) SetTabReorderable(child *Widget, reorderable T.Gboolean) {
	notebookSetTabReorderable(n, child, reorderable)
}
func (n *Notebook) SetTabVborder(tabVborder uint) { notebookSetTabVborder(n, tabVborder) }

type NumberUpLayout T.Enum

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

var NumberUpLayoutGetType func() T.GType
