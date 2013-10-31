// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Label struct {
	Misc  Misc
	Label *T.Gchar
	Bits  uint
	// JType : 2
	// Wrap : 1
	// UseUnderline : 1
	// UseMarkup : 1
	// EllipSize : 3
	// SingleLineMode : 1
	// HaveTransform : 1
	// InClick : 1
	// WrapMode : 3
	// PatternSet : 1
	// TrackLinks : 1
	MnemonicKeyval uint
	Text           *T.Gchar
	Attrs          *P.AttrList
	EffectiveAttrs *P.AttrList
	Layout         *P.Layout
	MnemonicWidget *Widget
	MnemonicWindow *Window
	SelectInfo     *LabelSelectionInfo
}

type LabelSelectionInfo struct{}

var (
	LabelGetType         func() O.Type
	LabelNew             func(str string) *Widget
	LabelNewWithMnemonic func(str string) *Widget

	labelGet                   func(l *Label, str **T.Gchar)
	labelGetAngle              func(l *Label) float64
	labelGetAttributes         func(l *Label) *P.AttrList
	labelGetCurrentUri         func(l *Label) string
	labelGetEllipsize          func(l *Label) P.EllipsizeMode
	labelGetJustify            func(l *Label) Justification
	labelGetLabel              func(l *Label) string
	labelGetLayout             func(l *Label) *P.Layout
	labelGetLayoutOffsets      func(l *Label, x, y *int)
	labelGetLineWrap           func(l *Label) bool
	labelGetLineWrapMode       func(l *Label) P.WrapMode
	labelGetMaxWidthChars      func(l *Label) int
	labelGetMnemonicKeyval     func(l *Label) uint
	labelGetMnemonicWidget     func(l *Label) *Widget
	labelGetSelectable         func(l *Label) bool
	labelGetSelectionBounds    func(l *Label, start, end *int) bool
	labelGetSingleLineMode     func(l *Label) bool
	labelGetText               func(l *Label) string
	labelGetTrackVisitedLinks  func(l *Label) bool
	labelGetUseMarkup          func(l *Label) bool
	labelGetUseUnderline       func(l *Label) bool
	labelGetWidthChars         func(l *Label) int
	labelParseUline            func(l *Label, string string) uint
	labelSelectRegion          func(l *Label, startOffset, endOffset int)
	labelSetAngle              func(l *Label, angle float64)
	labelSetAttributes         func(l *Label, attrs *P.AttrList)
	labelSetEllipsize          func(l *Label, mode P.EllipsizeMode)
	labelSetJustify            func(l *Label, jtype Justification)
	labelSetLabel              func(l *Label, str string)
	labelSetLineWrap           func(l *Label, wrap bool)
	labelSetLineWrapMode       func(l *Label, wrapMode P.WrapMode)
	labelSetMarkup             func(l *Label, str string)
	labelSetMarkupWithMnemonic func(l *Label, str string)
	labelSetMaxWidthChars      func(l *Label, nChars int)
	labelSetMnemonicWidget     func(l *Label, widget *Widget)
	labelSetPattern            func(l *Label, pattern string)
	labelSetSelectable         func(l *Label, setting bool)
	labelSetSingleLineMode     func(l *Label, singleLineMode bool)
	labelSetText               func(l *Label, str string)
	labelSetTextWithMnemonic   func(l *Label, str string)
	labelSetTrackVisitedLinks  func(l *Label, trackLinks bool)
	labelSetUseMarkup          func(l *Label, setting bool)
	labelSetUseUnderline       func(l *Label, setting bool)
	labelSetWidthChars         func(l *Label, nChars int)
)

func (l *Label) Get(str **T.Gchar)             { labelGet(l, str) }
func (l *Label) GetAngle() float64             { return labelGetAngle(l) }
func (l *Label) GetAttributes() *P.AttrList    { return labelGetAttributes(l) }
func (l *Label) GetCurrentUri() string         { return labelGetCurrentUri(l) }
func (l *Label) GetEllipsize() P.EllipsizeMode { return labelGetEllipsize(l) }
func (l *Label) GetJustify() Justification     { return labelGetJustify(l) }
func (l *Label) GetLabel() string              { return labelGetLabel(l) }
func (l *Label) GetLayout() *P.Layout          { return labelGetLayout(l) }
func (l *Label) GetLayoutOffsets(x, y *int)    { labelGetLayoutOffsets(l, x, y) }
func (l *Label) GetLineWrap() bool       { return labelGetLineWrap(l) }
func (l *Label) GetLineWrapMode() P.WrapMode   { return labelGetLineWrapMode(l) }
func (l *Label) GetMaxWidthChars() int         { return labelGetMaxWidthChars(l) }
func (l *Label) GetMnemonicKeyval() uint       { return labelGetMnemonicKeyval(l) }
func (l *Label) GetMnemonicWidget() *Widget    { return labelGetMnemonicWidget(l) }
func (l *Label) GetSelectable() bool           { return labelGetSelectable(l) }
func (l *Label) GetSelectionBounds(start, end *int) bool {
	return labelGetSelectionBounds(l, start, end)
}
func (l *Label) GetSingleLineMode() bool                 { return labelGetSingleLineMode(l) }
func (l *Label) GetText() string                         { return labelGetText(l) }
func (l *Label) GetTrackVisitedLinks() bool              { return labelGetTrackVisitedLinks(l) }
func (l *Label) GetUseMarkup() bool                      { return labelGetUseMarkup(l) }
func (l *Label) GetUseUnderline() bool                   { return labelGetUseUnderline(l) }
func (l *Label) GetWidthChars() int                      { return labelGetWidthChars(l) }
func (l *Label) ParseUline(str string) uint              { return labelParseUline(l, str) }
func (l *Label) SelectRegion(startOffset, endOffset int) { labelSelectRegion(l, startOffset, endOffset) }
func (l *Label) SetAngle(angle float64)                  { labelSetAngle(l, angle) }
func (l *Label) SetAttributes(attrs *P.AttrList)         { labelSetAttributes(l, attrs) }
func (l *Label) SetEllipsize(mode P.EllipsizeMode)       { labelSetEllipsize(l, mode) }
func (l *Label) SetJustify(jtype Justification)          { labelSetJustify(l, jtype) }
func (l *Label) SetLabel(str string)                     { labelSetLabel(l, str) }
func (l *Label) SetLineWrap(wrap bool)                   { labelSetLineWrap(l, wrap) }
func (l *Label) SetLineWrapMode(wrapMode P.WrapMode)     { labelSetLineWrapMode(l, wrapMode) }
func (l *Label) SetMarkup(str string)                    { labelSetMarkup(l, str) }
func (l *Label) SetMarkupWithMnemonic(str string)        { labelSetMarkupWithMnemonic(l, str) }
func (l *Label) SetMaxWidthChars(nChars int)             { labelSetMaxWidthChars(l, nChars) }
func (l *Label) SetMnemonicWidget(widget *Widget)        { labelSetMnemonicWidget(l, widget) }
func (l *Label) SetPattern(pattern string)               { labelSetPattern(l, pattern) }
func (l *Label) SetSelectable(setting bool)              { labelSetSelectable(l, setting) }
func (l *Label) SetSingleLineMode(singleLineMode bool) {
	labelSetSingleLineMode(l, singleLineMode)
}
func (l *Label) SetText(str string)                   { labelSetText(l, str) }
func (l *Label) SetTextWithMnemonic(str string)       { labelSetTextWithMnemonic(l, str) }
func (l *Label) SetTrackVisitedLinks(trackLinks bool) { labelSetTrackVisitedLinks(l, trackLinks) }
func (l *Label) SetUseMarkup(setting bool)            { labelSetUseMarkup(l, setting) }
func (l *Label) SetUseUnderline(setting bool)         { labelSetUseUnderline(l, setting) }
func (l *Label) SetWidthChars(nChars int)             { labelSetWidthChars(l, nChars) }

type LabelClass struct {
	Parent        MiscClass
	MoveCursor    func(label *Label, step MovementStep, count int, extendSelection bool)
	CopyClipboard func(label *Label)
	PopulatePopup func(label *Label, menu *Menu)
	Activate_link func(label *Label, uri *T.Gchar) bool
	_, _, _       func()
}

type Layout struct {
	Container    Container
	Children     *T.GList
	Width        uint
	Height       uint
	Hadjustment  *Adjustment
	Vadjustment  *Adjustment
	Bin_window   *D.Window
	Visibility   D.VisibilityState
	Scroll_x     int
	Scroll_y     int
	Freeze_count uint
}

var (
	LayoutGetType func() O.Type
	LayoutNew     func(hadjustment, vadjustment *Adjustment) *Widget

	layoutFreeze         func(l *Layout)
	layoutGetBinWindow   func(l *Layout) *D.Window
	layoutGetHadjustment func(l *Layout) *Adjustment
	layoutGetSize        func(l *Layout, width, height *uint)
	layoutGetVadjustment func(l *Layout) *Adjustment
	layoutMove           func(l *Layout, childWidget *Widget, x, y int)
	layoutPut            func(l *Layout, childWidget *Widget, x, y int)
	layoutSetHadjustment func(l *Layout, adjustment *Adjustment)
	layoutSetSize        func(l *Layout, width, height uint)
	layoutSetVadjustment func(l *Layout, adjustment *Adjustment)
	layoutThaw           func(l *Layout)
)

func (l *Layout) Freeze()                               { layoutFreeze(l) }
func (l *Layout) GetBinWindow() *D.Window               { return layoutGetBinWindow(l) }
func (l *Layout) GetHadjustment() *Adjustment           { return layoutGetHadjustment(l) }
func (l *Layout) GetSize(width, height *uint)           { layoutGetSize(l, width, height) }
func (l *Layout) GetVadjustment() *Adjustment           { return layoutGetVadjustment(l) }
func (l *Layout) Move(childWidget *Widget, x, y int)    { layoutMove(l, childWidget, x, y) }
func (l *Layout) Put(childWidget *Widget, x, y int)     { layoutPut(l, childWidget, x, y) }
func (l *Layout) SetHadjustment(adjustment *Adjustment) { layoutSetHadjustment(l, adjustment) }
func (l *Layout) SetSize(width, height uint)            { layoutSetSize(l, width, height) }
func (l *Layout) SetVadjustment(adjustment *Adjustment) { layoutSetVadjustment(l, adjustment) }
func (l *Layout) Thaw()                                 { layoutThaw(l) }

type ListStore struct {
	Parent             T.GObject
	Stamp              int
	Seq                T.Gpointer
	_                  T.Gpointer
	SortList           *T.GList
	NColumns           int
	SortColumnId       int
	Order              SortType
	ColumnHeaders      *O.Type
	Length             int
	DefaultSortFunc    TreeIterCompareFunc
	DefaultSortData    T.Gpointer
	DefaultSortDestroy T.GDestroyNotify
	ColumnsDirty       uint // : 1
}

type (
	LinkButton struct {
		Parent Button
		_      *struct{}
	}

	LinkButtonUriFunc func(
		button *LinkButton, link string, userData T.Gpointer)
)

var (
	ListStoreGetType func() O.Type
	ListStoreNew     func(nColumns int, v ...VArg) *ListStore
	ListStoreNewv    func(nColumns int, types *O.Type) *ListStore

	listStoreAppend            func(l *ListStore, iter *TreeIter)
	listStoreClear             func(l *ListStore)
	listStoreInsert            func(l *ListStore, iter *TreeIter, position int)
	listStoreInsertAfter       func(l *ListStore, iter *TreeIter, sibling *TreeIter)
	listStoreInsertBefore      func(l *ListStore, iter *TreeIter, sibling *TreeIter)
	listStoreInsertWithValues  func(l *ListStore, iter *TreeIter, position int, v ...VArg)
	listStoreInsertWithValuesv func(l *ListStore, iter *TreeIter, position int, columns *int, values *T.GValue, nValues int)
	listStoreIterIsValid       func(l *ListStore, iter *TreeIter) bool
	listStoreMoveAfter         func(l *ListStore, iter *TreeIter, position *TreeIter)
	listStoreMoveBefore        func(l *ListStore, iter *TreeIter, position *TreeIter)
	listStorePrepend           func(l *ListStore, iter *TreeIter)
	listStoreRemove            func(l *ListStore, iter *TreeIter) bool
	listStoreReorder           func(l *ListStore, newOrder *int)
	listStoreSet               func(l *ListStore, iter *TreeIter, v ...VArg)
	listStoreSetColumnTypes    func(l *ListStore, nColumns int, types *O.Type)
	listStoreSetValist         func(l *ListStore, iter *TreeIter, varArgs T.VaList)
	listStoreSetValue          func(l *ListStore, iter *TreeIter, column int, value *T.GValue)
	listStoreSetValuesv        func(l *ListStore, iter *TreeIter, columns *int, values *T.GValue, nValues int)
	listStoreSwap              func(l *ListStore, a, b *TreeIter)
)

func (l *ListStore) Append(iter *TreeIter)               { listStoreAppend(l, iter) }
func (l *ListStore) Clear()                              { listStoreClear(l) }
func (l *ListStore) Insert(iter *TreeIter, position int) { listStoreInsert(l, iter, position) }
func (l *ListStore) InsertAfter(iter *TreeIter, sibling *TreeIter) {
	listStoreInsertAfter(l, iter, sibling)
}
func (l *ListStore) InsertBefore(iter *TreeIter, sibling *TreeIter) {
	listStoreInsertBefore(l, iter, sibling)
}
func (l *ListStore) InsertWithValues(iter *TreeIter, position int, v ...VArg) {
	listStoreInsertWithValues(l, iter, position, v)
}
func (l *ListStore) InsertWithValuesv(iter *TreeIter, position int, columns *int, values *T.GValue, nValues int) {
	listStoreInsertWithValuesv(l, iter, position, columns, values, nValues)
}
func (l *ListStore) IterIsValid(iter *TreeIter) bool { return listStoreIterIsValid(l, iter) }
func (l *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	listStoreMoveAfter(l, iter, position)
}
func (l *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	listStoreMoveBefore(l, iter, position)
}
func (l *ListStore) Prepend(iter *TreeIter)        { listStorePrepend(l, iter) }
func (l *ListStore) Remove(iter *TreeIter) bool    { return listStoreRemove(l, iter) }
func (l *ListStore) Reorder(newOrder *int)         { listStoreReorder(l, newOrder) }
func (l *ListStore) Set(iter *TreeIter, v ...VArg) { listStoreSet(l, iter, v) }
func (l *ListStore) SetColumnTypes(nColumns int, types *O.Type) {
	listStoreSetColumnTypes(l, nColumns, types)
}
func (l *ListStore) SetValist(iter *TreeIter, varArgs T.VaList) {
	listStoreSetValist(l, iter, varArgs)
}
func (l *ListStore) SetValue(iter *TreeIter, column int, value *T.GValue) {
	listStoreSetValue(l, iter, column, value)
}
func (l *ListStore) SetValuesv(iter *TreeIter, columns *int, values *T.GValue, nValues int) {
	listStoreSetValuesv(l, iter, columns, values, nValues)
}
func (l *ListStore) Swap(a, b *TreeIter) { listStoreSwap(l, a, b) }

var (
	LinkButtonGetType      func() O.Type
	LinkButtonNew          func(uri string) *Widget
	LinkButtonNewWithLabel func(uri string, label string) *Widget

	LinkButtonSetUriHook func(f LinkButtonUriFunc, data T.Gpointer, destroy T.GDestroyNotify) LinkButtonUriFunc

	linkButtonGetUri     func(l *LinkButton) string
	linkButtonSetUri     func(l *LinkButton, uri string)
	linkButtonGetVisited func(l *LinkButton) bool
	linkButtonSetVisited func(l *LinkButton, visited bool)
)

func (l *LinkButton) GetUri() string          { return linkButtonGetUri(l) }
func (l *LinkButton) GetVisited() bool        { return linkButtonGetVisited(l) }
func (l *LinkButton) SetUri(uri string)       { linkButtonSetUri(l, uri) }
func (l *LinkButton) SetVisited(visited bool) { linkButtonSetVisited(l, visited) }

type List struct {
	Container         Container
	Children          *T.GList
	Selection         *T.GList
	UndoSelectionList *T.GList // Name ambuiguity with method
	UndoUnselection   *T.GList
	LastFocusChild    *Widget
	UndoFocusChild    *Widget
	Htimer            uint
	Vtimer            uint
	Anchor            int
	DragPos           int
	AnchorState       StateType
	Bits              uint
	// SelectionMode : 2
	// DragSelection : 1
	// AddMode : 1
}

var (
	ListGetType          func() O.Type
	ListNew              func() *Widget
	ListItemNewWithLabel func(label string) *Widget

	listInsertItems        func(l *List, items *T.GList, position int)
	listAppendItems        func(l *List, items *T.GList)
	listPrependItems       func(l *List, items *T.GList)
	listRemoveItems        func(l *List, items *T.GList)
	listRemoveItemsNoUnref func(l *List, items *T.GList)
	listClearItems         func(l *List, start int, end int)
	listSelectItem         func(l *List, item int)
	listUnselectItem       func(l *List, item int)
	listSelectChild        func(l *List, child *Widget)
	listUnselectChild      func(l *List, child *Widget)
	listChildPosition      func(l *List, child *Widget) int
	listSetSelectionMode   func(l *List, mode SelectionMode)
	listExtendSelection    func(l *List, scrollType ScrollType, position float32, autoStartSelection bool)
	listStartSelection     func(l *List)
	listEndSelection       func(l *List)
	listSelectAll          func(l *List)
	listUnselectAll        func(l *List)
	listScrollHorizontal   func(l *List, scrollType ScrollType, position float32)
	listScrollVertical     func(l *List, scrollType ScrollType, position float32)
	listToggleAddMode      func(l *List)
	listToggleFocusRow     func(l *List)
	listToggleRow          func(l *List, item *Widget)
	listUndoSelection      func(l *List)
	listEndDragSelection   func(l *List)
)

func (l *List) AppendItems(items *T.GList)      { listAppendItems(l, items) }
func (l *List) ChildPosition(child *Widget) int { return listChildPosition(l, child) }
func (l *List) ClearItems(start int, end int)   { listClearItems(l, start, end) }
func (l *List) EndDragSelection()               { listEndDragSelection(l) }
func (l *List) EndSelection()                   { listEndSelection(l) }
func (l *List) ExtendSelection(scrollType ScrollType, position float32, autoStartSelection bool) {
	listExtendSelection(l, scrollType, position, autoStartSelection)
}
func (l *List) InsertItems(items *T.GList, position int) { listInsertItems(l, items, position) }
func (l *List) PrependItems(items *T.GList)              { listPrependItems(l, items) }
func (l *List) RemoveItems(items *T.GList)               { listRemoveItems(l, items) }
func (l *List) RemoveItemsNoUnref(items *T.GList)        { listRemoveItemsNoUnref(l, items) }
func (l *List) ScrollHorizontal(scrollType ScrollType, position float32) {
	listScrollHorizontal(l, scrollType, position)
}
func (l *List) ScrollVertical(scrollType ScrollType, position float32) {
	listScrollVertical(l, scrollType, position)
}
func (l *List) SelectAll()                          { listSelectAll(l) }
func (l *List) SelectChild(child *Widget)           { listSelectChild(l, child) }
func (l *List) SelectItem(item int)                 { listSelectItem(l, item) }
func (l *List) SetSelectionMode(mode SelectionMode) { listSetSelectionMode(l, mode) }
func (l *List) StartSelection()                     { listStartSelection(l) }
func (l *List) ToggleAddMode()                      { listToggleAddMode(l) }
func (l *List) ToggleFocusRow()                     { listToggleFocusRow(l) }
func (l *List) ToggleRow(item *Widget)              { listToggleRow(l, item) }
func (l *List) UndoSelection()                      { listUndoSelection(l) }
func (l *List) UnselectAll()                        { listUnselectAll(l) }
func (l *List) UnselectChild(child *Widget)         { listUnselectChild(l, child) }
func (l *List) UnselectItem(item int)               { listUnselectItem(l, item) }

type ListItem struct{ Item Item }

var (
	ListItemGetType func() O.Type
	ListItemNew     func() *Widget

	listItemSelect   func(l *ListItem)
	listItemDeselect func(l *ListItem)
)

func (l *ListItem) ItemDeselect() { listItemDeselect(l) }
func (l *ListItem) ItemSelect()   { listItemSelect(l) }
