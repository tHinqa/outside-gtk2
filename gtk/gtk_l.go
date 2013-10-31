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

	LabelGet                   func(l *Label, str **T.Gchar)
	LabelGetAngle              func(l *Label) float64
	LabelGetAttributes         func(l *Label) *P.AttrList
	LabelGetCurrentUri         func(l *Label) string
	LabelGetEllipsize          func(l *Label) P.EllipsizeMode
	LabelGetJustify            func(l *Label) Justification
	LabelGetLabel              func(l *Label) string
	LabelGetLayout             func(l *Label) *P.Layout
	LabelGetLayoutOffsets      func(l *Label, x, y *int)
	LabelGetLineWrap           func(l *Label) bool
	LabelGetLineWrapMode       func(l *Label) P.WrapMode
	LabelGetMaxWidthChars      func(l *Label) int
	LabelGetMnemonicKeyval     func(l *Label) uint
	LabelGetMnemonicWidget     func(l *Label) *Widget
	LabelGetSelectable         func(l *Label) bool
	LabelGetSelectionBounds    func(l *Label, start, end *int) bool
	LabelGetSingleLineMode     func(l *Label) bool
	LabelGetText               func(l *Label) string
	LabelGetTrackVisitedLinks  func(l *Label) bool
	LabelGetUseMarkup          func(l *Label) bool
	LabelGetUseUnderline       func(l *Label) bool
	LabelGetWidthChars         func(l *Label) int
	LabelParseUline            func(l *Label, string string) uint
	LabelSelectRegion          func(l *Label, startOffset, endOffset int)
	LabelSetAngle              func(l *Label, angle float64)
	LabelSetAttributes         func(l *Label, attrs *P.AttrList)
	LabelSetEllipsize          func(l *Label, mode P.EllipsizeMode)
	LabelSetJustify            func(l *Label, jtype Justification)
	LabelSetLabel              func(l *Label, str string)
	LabelSetLineWrap           func(l *Label, wrap bool)
	LabelSetLineWrapMode       func(l *Label, wrapMode P.WrapMode)
	LabelSetMarkup             func(l *Label, str string)
	LabelSetMarkupWithMnemonic func(l *Label, str string)
	LabelSetMaxWidthChars      func(l *Label, nChars int)
	LabelSetMnemonicWidget     func(l *Label, widget *Widget)
	LabelSetPattern            func(l *Label, pattern string)
	LabelSetSelectable         func(l *Label, setting bool)
	LabelSetSingleLineMode     func(l *Label, singleLineMode bool)
	LabelSetText               func(l *Label, str string)
	LabelSetTextWithMnemonic   func(l *Label, str string)
	LabelSetTrackVisitedLinks  func(l *Label, trackLinks bool)
	LabelSetUseMarkup          func(l *Label, setting bool)
	LabelSetUseUnderline       func(l *Label, setting bool)
	LabelSetWidthChars         func(l *Label, nChars int)
)

func (l *Label) Get(str **T.Gchar)             { LabelGet(l, str) }
func (l *Label) GetAngle() float64             { return LabelGetAngle(l) }
func (l *Label) GetAttributes() *P.AttrList    { return LabelGetAttributes(l) }
func (l *Label) GetCurrentUri() string         { return LabelGetCurrentUri(l) }
func (l *Label) GetEllipsize() P.EllipsizeMode { return LabelGetEllipsize(l) }
func (l *Label) GetJustify() Justification     { return LabelGetJustify(l) }
func (l *Label) GetLabel() string              { return LabelGetLabel(l) }
func (l *Label) GetLayout() *P.Layout          { return LabelGetLayout(l) }
func (l *Label) GetLayoutOffsets(x, y *int)    { LabelGetLayoutOffsets(l, x, y) }
func (l *Label) GetLineWrap() bool       { return LabelGetLineWrap(l) }
func (l *Label) GetLineWrapMode() P.WrapMode   { return LabelGetLineWrapMode(l) }
func (l *Label) GetMaxWidthChars() int         { return LabelGetMaxWidthChars(l) }
func (l *Label) GetMnemonicKeyval() uint       { return LabelGetMnemonicKeyval(l) }
func (l *Label) GetMnemonicWidget() *Widget    { return LabelGetMnemonicWidget(l) }
func (l *Label) GetSelectable() bool           { return LabelGetSelectable(l) }
func (l *Label) GetSelectionBounds(start, end *int) bool {
	return LabelGetSelectionBounds(l, start, end)
}
func (l *Label) GetSingleLineMode() bool                 { return LabelGetSingleLineMode(l) }
func (l *Label) GetText() string                         { return LabelGetText(l) }
func (l *Label) GetTrackVisitedLinks() bool              { return LabelGetTrackVisitedLinks(l) }
func (l *Label) GetUseMarkup() bool                      { return LabelGetUseMarkup(l) }
func (l *Label) GetUseUnderline() bool                   { return LabelGetUseUnderline(l) }
func (l *Label) GetWidthChars() int                      { return LabelGetWidthChars(l) }
func (l *Label) ParseUline(str string) uint              { return LabelParseUline(l, str) }
func (l *Label) SelectRegion(startOffset, endOffset int) { LabelSelectRegion(l, startOffset, endOffset) }
func (l *Label) SetAngle(angle float64)                  { LabelSetAngle(l, angle) }
func (l *Label) SetAttributes(attrs *P.AttrList)         { LabelSetAttributes(l, attrs) }
func (l *Label) SetEllipsize(mode P.EllipsizeMode)       { LabelSetEllipsize(l, mode) }
func (l *Label) SetJustify(jtype Justification)          { LabelSetJustify(l, jtype) }
func (l *Label) SetLabel(str string)                     { LabelSetLabel(l, str) }
func (l *Label) SetLineWrap(wrap bool)                   { LabelSetLineWrap(l, wrap) }
func (l *Label) SetLineWrapMode(wrapMode P.WrapMode)     { LabelSetLineWrapMode(l, wrapMode) }
func (l *Label) SetMarkup(str string)                    { LabelSetMarkup(l, str) }
func (l *Label) SetMarkupWithMnemonic(str string)        { LabelSetMarkupWithMnemonic(l, str) }
func (l *Label) SetMaxWidthChars(nChars int)             { LabelSetMaxWidthChars(l, nChars) }
func (l *Label) SetMnemonicWidget(widget *Widget)        { LabelSetMnemonicWidget(l, widget) }
func (l *Label) SetPattern(pattern string)               { LabelSetPattern(l, pattern) }
func (l *Label) SetSelectable(setting bool)              { LabelSetSelectable(l, setting) }
func (l *Label) SetSingleLineMode(singleLineMode bool) {
	LabelSetSingleLineMode(l, singleLineMode)
}
func (l *Label) SetText(str string)                   { LabelSetText(l, str) }
func (l *Label) SetTextWithMnemonic(str string)       { LabelSetTextWithMnemonic(l, str) }
func (l *Label) SetTrackVisitedLinks(trackLinks bool) { LabelSetTrackVisitedLinks(l, trackLinks) }
func (l *Label) SetUseMarkup(setting bool)            { LabelSetUseMarkup(l, setting) }
func (l *Label) SetUseUnderline(setting bool)         { LabelSetUseUnderline(l, setting) }
func (l *Label) SetWidthChars(nChars int)             { LabelSetWidthChars(l, nChars) }

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

	LayoutFreeze         func(l *Layout)
	LayoutGetBinWindow   func(l *Layout) *D.Window
	LayoutGetHadjustment func(l *Layout) *Adjustment
	LayoutGetSize        func(l *Layout, width, height *uint)
	LayoutGetVadjustment func(l *Layout) *Adjustment
	LayoutMove           func(l *Layout, childWidget *Widget, x, y int)
	LayoutPut            func(l *Layout, childWidget *Widget, x, y int)
	LayoutSetHadjustment func(l *Layout, adjustment *Adjustment)
	LayoutSetSize        func(l *Layout, width, height uint)
	LayoutSetVadjustment func(l *Layout, adjustment *Adjustment)
	LayoutThaw           func(l *Layout)
)

func (l *Layout) Freeze()                               { LayoutFreeze(l) }
func (l *Layout) GetBinWindow() *D.Window               { return LayoutGetBinWindow(l) }
func (l *Layout) GetHadjustment() *Adjustment           { return LayoutGetHadjustment(l) }
func (l *Layout) GetSize(width, height *uint)           { LayoutGetSize(l, width, height) }
func (l *Layout) GetVadjustment() *Adjustment           { return LayoutGetVadjustment(l) }
func (l *Layout) Move(childWidget *Widget, x, y int)    { LayoutMove(l, childWidget, x, y) }
func (l *Layout) Put(childWidget *Widget, x, y int)     { LayoutPut(l, childWidget, x, y) }
func (l *Layout) SetHadjustment(adjustment *Adjustment) { LayoutSetHadjustment(l, adjustment) }
func (l *Layout) SetSize(width, height uint)            { LayoutSetSize(l, width, height) }
func (l *Layout) SetVadjustment(adjustment *Adjustment) { LayoutSetVadjustment(l, adjustment) }
func (l *Layout) Thaw()                                 { LayoutThaw(l) }

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

	ListStoreAppend            func(l *ListStore, iter *TreeIter)
	ListStoreClear             func(l *ListStore)
	ListStoreInsert            func(l *ListStore, iter *TreeIter, position int)
	ListStoreInsertAfter       func(l *ListStore, iter *TreeIter, sibling *TreeIter)
	ListStoreInsertBefore      func(l *ListStore, iter *TreeIter, sibling *TreeIter)
	ListStoreInsertWithValues  func(l *ListStore, iter *TreeIter, position int, v ...VArg)
	ListStoreInsertWithValuesv func(l *ListStore, iter *TreeIter, position int, columns *int, values *T.GValue, nValues int)
	ListStoreIterIsValid       func(l *ListStore, iter *TreeIter) bool
	ListStoreMoveAfter         func(l *ListStore, iter *TreeIter, position *TreeIter)
	ListStoreMoveBefore        func(l *ListStore, iter *TreeIter, position *TreeIter)
	ListStorePrepend           func(l *ListStore, iter *TreeIter)
	ListStoreRemove            func(l *ListStore, iter *TreeIter) bool
	ListStoreReorder           func(l *ListStore, newOrder *int)
	ListStoreSet               func(l *ListStore, iter *TreeIter, v ...VArg)
	ListStoreSetColumnTypes    func(l *ListStore, nColumns int, types *O.Type)
	ListStoreSetValist         func(l *ListStore, iter *TreeIter, varArgs T.VaList)
	ListStoreSetValue          func(l *ListStore, iter *TreeIter, column int, value *T.GValue)
	ListStoreSetValuesv        func(l *ListStore, iter *TreeIter, columns *int, values *T.GValue, nValues int)
	ListStoreSwap              func(l *ListStore, a, b *TreeIter)
)

func (l *ListStore) Append(iter *TreeIter)               { ListStoreAppend(l, iter) }
func (l *ListStore) Clear()                              { ListStoreClear(l) }
func (l *ListStore) Insert(iter *TreeIter, position int) { ListStoreInsert(l, iter, position) }
func (l *ListStore) InsertAfter(iter *TreeIter, sibling *TreeIter) {
	ListStoreInsertAfter(l, iter, sibling)
}
func (l *ListStore) InsertBefore(iter *TreeIter, sibling *TreeIter) {
	ListStoreInsertBefore(l, iter, sibling)
}
func (l *ListStore) InsertWithValues(iter *TreeIter, position int, v ...VArg) {
	ListStoreInsertWithValues(l, iter, position, v)
}
func (l *ListStore) InsertWithValuesv(iter *TreeIter, position int, columns *int, values *T.GValue, nValues int) {
	ListStoreInsertWithValuesv(l, iter, position, columns, values, nValues)
}
func (l *ListStore) IterIsValid(iter *TreeIter) bool { return ListStoreIterIsValid(l, iter) }
func (l *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	ListStoreMoveAfter(l, iter, position)
}
func (l *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	ListStoreMoveBefore(l, iter, position)
}
func (l *ListStore) Prepend(iter *TreeIter)        { ListStorePrepend(l, iter) }
func (l *ListStore) Remove(iter *TreeIter) bool    { return ListStoreRemove(l, iter) }
func (l *ListStore) Reorder(newOrder *int)         { ListStoreReorder(l, newOrder) }
func (l *ListStore) Set(iter *TreeIter, v ...VArg) { ListStoreSet(l, iter, v) }
func (l *ListStore) SetColumnTypes(nColumns int, types *O.Type) {
	ListStoreSetColumnTypes(l, nColumns, types)
}
func (l *ListStore) SetValist(iter *TreeIter, varArgs T.VaList) {
	ListStoreSetValist(l, iter, varArgs)
}
func (l *ListStore) SetValue(iter *TreeIter, column int, value *T.GValue) {
	ListStoreSetValue(l, iter, column, value)
}
func (l *ListStore) SetValuesv(iter *TreeIter, columns *int, values *T.GValue, nValues int) {
	ListStoreSetValuesv(l, iter, columns, values, nValues)
}
func (l *ListStore) Swap(a, b *TreeIter) { ListStoreSwap(l, a, b) }

var (
	LinkButtonGetType      func() O.Type
	LinkButtonNew          func(uri string) *Widget
	LinkButtonNewWithLabel func(uri string, label string) *Widget

	LinkButtonSetUriHook func(f LinkButtonUriFunc, data T.Gpointer, destroy T.GDestroyNotify) LinkButtonUriFunc

	LinkButtonGetUri     func(l *LinkButton) string
	LinkButtonSetUri     func(l *LinkButton, uri string)
	LinkButtonGetVisited func(l *LinkButton) bool
	LinkButtonSetVisited func(l *LinkButton, visited bool)
)

func (l *LinkButton) GetUri() string          { return LinkButtonGetUri(l) }
func (l *LinkButton) GetVisited() bool        { return LinkButtonGetVisited(l) }
func (l *LinkButton) SetUri(uri string)       { LinkButtonSetUri(l, uri) }
func (l *LinkButton) SetVisited(visited bool) { LinkButtonSetVisited(l, visited) }

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

	ListInsertItems        func(l *List, items *T.GList, position int)
	ListAppendItems        func(l *List, items *T.GList)
	ListPrependItems       func(l *List, items *T.GList)
	ListRemoveItems        func(l *List, items *T.GList)
	ListRemoveItemsNoUnref func(l *List, items *T.GList)
	ListClearItems         func(l *List, start int, end int)
	ListSelectItem         func(l *List, item int)
	ListUnselectItem       func(l *List, item int)
	ListSelectChild        func(l *List, child *Widget)
	ListUnselectChild      func(l *List, child *Widget)
	ListChildPosition      func(l *List, child *Widget) int
	ListSetSelectionMode   func(l *List, mode SelectionMode)
	ListExtendSelection    func(l *List, scrollType ScrollType, position float32, autoStartSelection bool)
	ListStartSelection     func(l *List)
	ListEndSelection       func(l *List)
	ListSelectAll          func(l *List)
	ListUnselectAll        func(l *List)
	ListScrollHorizontal   func(l *List, scrollType ScrollType, position float32)
	ListScrollVertical     func(l *List, scrollType ScrollType, position float32)
	ListToggleAddMode      func(l *List)
	ListToggleFocusRow     func(l *List)
	ListToggleRow          func(l *List, item *Widget)
	ListUndoSelection      func(l *List)
	ListEndDragSelection   func(l *List)
)

func (l *List) AppendItems(items *T.GList)      { ListAppendItems(l, items) }
func (l *List) ChildPosition(child *Widget) int { return ListChildPosition(l, child) }
func (l *List) ClearItems(start int, end int)   { ListClearItems(l, start, end) }
func (l *List) EndDragSelection()               { ListEndDragSelection(l) }
func (l *List) EndSelection()                   { ListEndSelection(l) }
func (l *List) ExtendSelection(scrollType ScrollType, position float32, autoStartSelection bool) {
	ListExtendSelection(l, scrollType, position, autoStartSelection)
}
func (l *List) InsertItems(items *T.GList, position int) { ListInsertItems(l, items, position) }
func (l *List) PrependItems(items *T.GList)              { ListPrependItems(l, items) }
func (l *List) RemoveItems(items *T.GList)               { ListRemoveItems(l, items) }
func (l *List) RemoveItemsNoUnref(items *T.GList)        { ListRemoveItemsNoUnref(l, items) }
func (l *List) ScrollHorizontal(scrollType ScrollType, position float32) {
	ListScrollHorizontal(l, scrollType, position)
}
func (l *List) ScrollVertical(scrollType ScrollType, position float32) {
	ListScrollVertical(l, scrollType, position)
}
func (l *List) SelectAll()                          { ListSelectAll(l) }
func (l *List) SelectChild(child *Widget)           { ListSelectChild(l, child) }
func (l *List) SelectItem(item int)                 { ListSelectItem(l, item) }
func (l *List) SetSelectionMode(mode SelectionMode) { ListSetSelectionMode(l, mode) }
func (l *List) StartSelection()                     { ListStartSelection(l) }
func (l *List) ToggleAddMode()                      { ListToggleAddMode(l) }
func (l *List) ToggleFocusRow()                     { ListToggleFocusRow(l) }
func (l *List) ToggleRow(item *Widget)              { ListToggleRow(l, item) }
func (l *List) UndoSelection()                      { ListUndoSelection(l) }
func (l *List) UnselectAll()                        { ListUnselectAll(l) }
func (l *List) UnselectChild(child *Widget)         { ListUnselectChild(l, child) }
func (l *List) UnselectItem(item int)               { ListUnselectItem(l, item) }

type ListItem struct{ Item Item }

var (
	ListItemGetType func() O.Type
	ListItemNew     func() *Widget

	ListItemSelect   func(l *ListItem)
	ListItemDeselect func(l *ListItem)
)

func (l *ListItem) ItemDeselect() { ListItemDeselect(l) }
func (l *ListItem) ItemSelect()   { ListItemSelect(l) }
