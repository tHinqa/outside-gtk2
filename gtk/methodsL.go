// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Label struct {
	Misc  T.GtkMisc
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
	Attrs          *T.PangoAttrList
	EffectiveAttrs *T.PangoAttrList
	Layout         *T.PangoLayout
	MnemonicWidget *Widget
	MnemonicWindow *T.GtkWindow
	SelectInfo     *LabelSelectionInfo
}

type LabelSelectionInfo struct{}

var (
	LabelGetType         func() T.GType
	LabelNew             func(str string) *Widget
	LabelNewWithMnemonic func(str string) *Widget

	LabelGet                   func(l *Label, str **T.Gchar)
	LabelGetAngle              func(l *Label) float64
	LabelGetAttributes         func(l *Label) *T.PangoAttrList
	LabelGetCurrentUri         func(l *Label) string
	LabelGetEllipsize          func(l *Label) T.PangoEllipsizeMode
	LabelGetJustify            func(l *Label) T.GtkJustification
	LabelGetLabel              func(l *Label) string
	LabelGetLayout             func(l *Label) *T.PangoLayout
	LabelGetLayoutOffsets      func(l *Label, x, y *int)
	LabelGetLineWrap           func(l *Label) T.Gboolean
	LabelGetLineWrapMode       func(l *Label) T.PangoWrapMode
	LabelGetMaxWidthChars      func(l *Label) int
	LabelGetMnemonicKeyval     func(l *Label) uint
	LabelGetMnemonicWidget     func(l *Label) *Widget
	LabelGetSelectable         func(l *Label) T.Gboolean
	LabelGetSelectionBounds    func(l *Label, start, end *int) T.Gboolean
	LabelGetSingleLineMode     func(l *Label) T.Gboolean
	LabelGetText               func(l *Label) string
	LabelGetTrackVisitedLinks  func(l *Label) T.Gboolean
	LabelGetUseMarkup          func(l *Label) T.Gboolean
	LabelGetUseUnderline       func(l *Label) T.Gboolean
	LabelGetWidthChars         func(l *Label) int
	LabelParseUline            func(l *Label, string string) uint
	LabelSelectRegion          func(l *Label, startOffset, endOffset int)
	LabelSetAngle              func(l *Label, angle float64)
	LabelSetAttributes         func(l *Label, attrs *T.PangoAttrList)
	LabelSetEllipsize          func(l *Label, mode T.PangoEllipsizeMode)
	LabelSetJustify            func(l *Label, jtype T.GtkJustification)
	LabelSetLabel              func(l *Label, str string)
	LabelSetLineWrap           func(l *Label, wrap T.Gboolean)
	LabelSetLineWrapMode       func(l *Label, wrapMode T.PangoWrapMode)
	LabelSetMarkup             func(l *Label, str string)
	LabelSetMarkupWithMnemonic func(l *Label, str string)
	LabelSetMaxWidthChars      func(l *Label, nChars int)
	LabelSetMnemonicWidget     func(l *Label, widget *Widget)
	LabelSetPattern            func(l *Label, pattern string)
	LabelSetSelectable         func(l *Label, setting T.Gboolean)
	LabelSetSingleLineMode     func(l *Label, singleLineMode T.Gboolean)
	LabelSetText               func(l *Label, str string)
	LabelSetTextWithMnemonic   func(l *Label, str string)
	LabelSetTrackVisitedLinks  func(l *Label, trackLinks T.Gboolean)
	LabelSetUseMarkup          func(l *Label, setting T.Gboolean)
	LabelSetUseUnderline       func(l *Label, setting T.Gboolean)
	LabelSetWidthChars         func(l *Label, nChars int)
)

func (l *Label) SetText(str string) {
	LabelSetText(l, str)
}

func (l *Label) GetText() string {
	return LabelGetText(l)
}

func (l *Label) SetAttributes(attrs *T.PangoAttrList) {
	LabelSetAttributes(l, attrs)
}

func (l *Label) GetAttributes() *T.PangoAttrList {
	return LabelGetAttributes(l)
}

func (l *Label) SetLabel(str string) {
	LabelSetLabel(l, str)
}

func (l *Label) GetLabel() string {
	return LabelGetLabel(l)
}

func (l *Label) SetMarkup(str string) {
	LabelSetMarkup(l, str)
}

func (l *Label) SetUseMarkup(setting T.Gboolean) {
	LabelSetUseMarkup(l, setting)
}

func (l *Label) GetUseMarkup() T.Gboolean {
	return LabelGetUseMarkup(l)
}

func (l *Label) SetUseUnderline(setting T.Gboolean) {
	LabelSetUseUnderline(l, setting)
}

func (l *Label) GetUseUnderline() T.Gboolean {
	return LabelGetUseUnderline(l)
}

func (l *Label) SetMarkupWithMnemonic(str string) {
	LabelSetMarkupWithMnemonic(l, str)
}

func (l *Label) GetMnemonicKeyval() uint {
	return LabelGetMnemonicKeyval(l)
}

func (l *Label) SetMnemonicWidget(widget *Widget) {
	LabelSetMnemonicWidget(l, widget)
}

func (l *Label) GetMnemonicWidget() *Widget {
	return LabelGetMnemonicWidget(l)
}

func (l *Label) SetTextWithMnemonic(str string) {
	LabelSetTextWithMnemonic(l, str)
}

func (l *Label) SetJustify(jtype T.GtkJustification) {
	LabelSetJustify(l, jtype)
}

func (l *Label) GetJustify() T.GtkJustification {
	return LabelGetJustify(l)
}

func (l *Label) SetEllipsize(mode T.PangoEllipsizeMode) {
	LabelSetEllipsize(l, mode)
}

func (l *Label) GetEllipsize() T.PangoEllipsizeMode {
	return LabelGetEllipsize(l)
}

func (l *Label) SetWidthChars(nChars int) {
	LabelSetWidthChars(l, nChars)
}

func (l *Label) GetWidthChars() int {
	return LabelGetWidthChars(l)
}

func (l *Label) SetMaxWidthChars(nChars int) {
	LabelSetMaxWidthChars(l, nChars)
}

func (l *Label) GetMaxWidthChars() int {
	return LabelGetMaxWidthChars(l)
}

func (l *Label) SetPattern(pattern string) {
	LabelSetPattern(l, pattern)
}

func (l *Label) SetLineWrap(wrap T.Gboolean) {
	LabelSetLineWrap(l, wrap)
}

func (l *Label) GetLineWrap() T.Gboolean {
	return LabelGetLineWrap(l)
}

func (l *Label) SetLineWrapMode(wrapMode T.PangoWrapMode) {
	LabelSetLineWrapMode(l, wrapMode)
}

func (l *Label) GetLineWrapMode() T.PangoWrapMode {
	return LabelGetLineWrapMode(l)
}

func (l *Label) SetSelectable(setting T.Gboolean) {
	LabelSetSelectable(l, setting)
}

func (l *Label) GetSelectable() T.Gboolean {
	return LabelGetSelectable(l)
}

func (l *Label) SetAngle(angle float64) {
	LabelSetAngle(l, angle)
}

func (l *Label) GetAngle() float64 {
	return LabelGetAngle(l)
}

func (l *Label) SelectRegion(startOffset, endOffset int) {
	LabelSelectRegion(l, startOffset, endOffset)
}

func (l *Label) GetSelectionBounds(start, end *int) T.Gboolean {
	return LabelGetSelectionBounds(l, start, end)
}

func (l *Label) GetLayout() *T.PangoLayout {
	return LabelGetLayout(l)
}

func (l *Label) GetLayoutOffsets(x, y *int) {
	LabelGetLayoutOffsets(l, x, y)
}

func (l *Label) SetSingleLineMode(singleLineMode T.Gboolean) {
	LabelSetSingleLineMode(l, singleLineMode)
}

func (l *Label) GetSingleLineMode() T.Gboolean {
	return LabelGetSingleLineMode(l)
}

func (l *Label) GetCurrentUri() string {
	return LabelGetCurrentUri(l)
}

func (l *Label) SetTrackVisitedLinks(trackLinks T.Gboolean) {
	LabelSetTrackVisitedLinks(l, trackLinks)
}

func (l *Label) GetTrackVisitedLinks() T.Gboolean {
	return LabelGetTrackVisitedLinks(l)
}

func (l *Label) Get(str **T.Gchar) {
	LabelGet(l, str)
}

func (l *Label) ParseUline(str string) uint {
	return LabelParseUline(l, str)
}

type Layout struct {
	Container    T.GtkContainer
	Children     *T.GList
	Width        uint
	Height       uint
	Hadjustment  *Adjustment
	Vadjustment  *Adjustment
	Bin_window   *T.GdkWindow
	Visibility   T.GdkVisibilityState
	Scroll_x     int
	Scroll_y     int
	Freeze_count uint
}

var (
	LayoutGetType func() T.GType
	LayoutNew     func(hadjustment, vadjustment *Adjustment) *Widget

	LayoutFreeze         func(l *Layout)
	LayoutGetBinWindow   func(l *Layout) *T.GdkWindow
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

func (l *Layout) GetBinWindow() *T.GdkWindow {
	return LayoutGetBinWindow(l)
}

func (l *Layout) Put(childWidget *Widget, x, y int) {
	LayoutPut(l, childWidget, x, y)
}

func (l *Layout) Move(childWidget *Widget, x, y int) {
	LayoutMove(l, childWidget, x, y)
}

func (l *Layout) SetSize(width, height uint) {
	LayoutSetSize(l, width, height)
}

func (l *Layout) GetSize(width, height *uint) {
	LayoutGetSize(l, width, height)
}

func (l *Layout) GetHadjustment() *Adjustment {
	return LayoutGetHadjustment(l)
}

func (l *Layout) GetVadjustment() *Adjustment {
	return LayoutGetVadjustment(l)
}

func (l *Layout) SetHadjustment(adjustment *Adjustment) {
	LayoutSetHadjustment(l, adjustment)
}

func (l *Layout) SetVadjustment(adjustment *Adjustment) {
	LayoutSetVadjustment(l, adjustment)
}

func (l *Layout) Freeze() { LayoutFreeze(l) }

func (l *Layout) Thaw() { LayoutThaw(l) }

type ListStore struct {
	Parent             T.GObject
	Stamp              int
	Seq                T.Gpointer
	_                  T.Gpointer
	SortList           *T.GList
	NColumns           int
	SortColumnId       int
	Order              T.GtkSortType
	ColumnHeaders      *T.GType
	Length             int
	DefaultSortFunc    T.GtkTreeIterCompareFunc
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
	ListStoreGetType func() T.GType
	ListStoreNew     func(nColumns int, v ...VArg) *ListStore
	ListStoreNewv    func(nColumns int, types *T.GType) *ListStore

	ListStoreAppend            func(l *ListStore, iter *T.GtkTreeIter)
	ListStoreClear             func(l *ListStore)
	ListStoreInsert            func(l *ListStore, iter *T.GtkTreeIter, position int)
	ListStoreInsertAfter       func(l *ListStore, iter *T.GtkTreeIter, sibling *T.GtkTreeIter)
	ListStoreInsertBefore      func(l *ListStore, iter *T.GtkTreeIter, sibling *T.GtkTreeIter)
	ListStoreInsertWithValues  func(l *ListStore, iter *T.GtkTreeIter, position int, v ...VArg)
	ListStoreInsertWithValuesv func(l *ListStore, iter *T.GtkTreeIter, position int, columns *int, values *T.GValue, nValues int)
	ListStoreIterIsValid       func(l *ListStore, iter *T.GtkTreeIter) T.Gboolean
	ListStoreMoveAfter         func(l *ListStore, iter *T.GtkTreeIter, position *T.GtkTreeIter)
	ListStoreMoveBefore        func(l *ListStore, iter *T.GtkTreeIter, position *T.GtkTreeIter)
	ListStorePrepend           func(l *ListStore, iter *T.GtkTreeIter)
	ListStoreRemove            func(l *ListStore, iter *T.GtkTreeIter) T.Gboolean
	ListStoreReorder           func(l *ListStore, newOrder *int)
	ListStoreSet               func(l *ListStore, iter *T.GtkTreeIter, v ...VArg)
	ListStoreSetColumnTypes    func(l *ListStore, nColumns int, types *T.GType)
	ListStoreSetValist         func(l *ListStore, iter *T.GtkTreeIter, varArgs T.VaList)
	ListStoreSetValue          func(l *ListStore, iter *T.GtkTreeIter, column int, value *T.GValue)
	ListStoreSetValuesv        func(l *ListStore, iter *T.GtkTreeIter, columns *int, values *T.GValue, nValues int)
	ListStoreSwap              func(l *ListStore, a, b *T.GtkTreeIter)
)

func (l *ListStore) SetColumnTypes(nColumns int, types *T.GType) {
	ListStoreSetColumnTypes(l, nColumns, types)
}

func (l *ListStore) SetValue(iter *T.GtkTreeIter, column int, value *T.GValue) {
	ListStoreSetValue(l, iter, column, value)
}

func (l *ListStore) Set(iter *T.GtkTreeIter, v ...VArg) {
	ListStoreSet(l, iter, v)
}

func (l *ListStore) SetValuesv(iter *T.GtkTreeIter, columns *int, values *T.GValue, nValues int) {
	ListStoreSetValuesv(l, iter, columns, values, nValues)
}

func (l *ListStore) SetValist(iter *T.GtkTreeIter, varArgs T.VaList) {
	ListStoreSetValist(l, iter, varArgs)
}

func (l *ListStore) Remove(iter *T.GtkTreeIter) T.Gboolean {
	return ListStoreRemove(l, iter)
}

func (l *ListStore) Insert(iter *T.GtkTreeIter, position int) {
	ListStoreInsert(l, iter, position)
}

func (l *ListStore) InsertBefore(iter *T.GtkTreeIter, sibling *T.GtkTreeIter) {
	ListStoreInsertBefore(l, iter, sibling)
}

func (l *ListStore) InsertAfter(iter *T.GtkTreeIter, sibling *T.GtkTreeIter) {
	ListStoreInsertAfter(l, iter, sibling)
}

func (l *ListStore) InsertWithValues(iter *T.GtkTreeIter, position int, v ...VArg) {
	ListStoreInsertWithValues(l, iter, position, v)
}

func (l *ListStore) InsertWithValuesv(iter *T.GtkTreeIter, position int, columns *int, values *T.GValue, nValues int) {
	ListStoreInsertWithValuesv(l, iter, position, columns, values, nValues)
}

func (l *ListStore) Prepend(iter *T.GtkTreeIter) {
	ListStorePrepend(l, iter)
}

func (l *ListStore) Append(iter *T.GtkTreeIter) {
	ListStoreAppend(l, iter)
}

func (l *ListStore) Clear() { ListStoreClear(l) }

func (l *ListStore) IterIsValid(iter *T.GtkTreeIter) T.Gboolean {
	return ListStoreIterIsValid(l, iter)
}

func (l *ListStore) Reorder(newOrder *int) {
	ListStoreReorder(l, newOrder)
}

func (l *ListStore) Swap(a, b *T.GtkTreeIter) {
	ListStoreSwap(l, a, b)
}

func (l *ListStore) MoveAfter(iter *T.GtkTreeIter, position *T.GtkTreeIter) {
	ListStoreMoveAfter(l, iter, position)
}

func (l *ListStore) MoveBefore(iter *T.GtkTreeIter, position *T.GtkTreeIter) {
	ListStoreMoveBefore(l, iter, position)
}

var (
	LinkButtonGetType      func() T.GType
	LinkButtonNew          func(uri string) *Widget
	LinkButtonNewWithLabel func(uri string, label string) *Widget

	LinkButtonSetUriHook func(f LinkButtonUriFunc, dataGpointer, destroy T.GDestroyNotify) LinkButtonUriFunc

	LinkButtonGetUri     func(l *LinkButton) string
	LinkButtonSetUri     func(l *LinkButton, uri string)
	LinkButtonGetVisited func(l *LinkButton) T.Gboolean
	LinkButtonSetVisited func(l *LinkButton, visited T.Gboolean)
)

type List struct {
	Container         T.GtkContainer
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
	AnchorState       T.GtkStateType
	Bits              uint
	// SelectionMode : 2
	// DragSelection : 1
	// AddMode : 1
}

var (
	ListGetType          func() T.GType
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
	ListSetSelectionMode   func(l *List, mode T.GtkSelectionMode)
	ListExtendSelection    func(l *List, scrollType T.GtkScrollType, position float32, autoStartSelection T.Gboolean)
	ListStartSelection     func(l *List)
	ListEndSelection       func(l *List)
	ListSelectAll          func(l *List)
	ListUnselectAll        func(l *List)
	ListScrollHorizontal   func(l *List, scrollType T.GtkScrollType, position float32)
	ListScrollVertical     func(l *List, scrollType T.GtkScrollType, position float32)
	ListToggleAddMode      func(l *List)
	ListToggleFocusRow     func(l *List)
	ListToggleRow          func(l *List, item *Widget)
	ListUndoSelection      func(l *List)
	ListEndDragSelection   func(l *List)
)

func (l *List) InsertItems(items *T.GList, position int) {
	ListInsertItems(l, items, position)
}

func (l *List) AppendItems(items *T.GList) {
	ListAppendItems(l, items)
}

func (l *List) PrependItems(items *T.GList) {
	ListPrependItems(l, items)
}

func (l *List) RemoveItems(items *T.GList) {
	ListRemoveItems(l, items)
}

func (l *List) RemoveItemsNoUnref(items *T.GList) {
	ListRemoveItemsNoUnref(l, items)
}

func (l *List) ClearItems(start int, end int) {
	ListClearItems(l, start, end)
}

func (l *List) SelectItem(item int) { ListSelectItem(l, item) }

func (l *List) UnselectItem(item int) {
	ListUnselectItem(l, item)
}

func (l *List) SelectChild(child *Widget) {
	ListSelectChild(l, child)
}

func (l *List) UnselectChild(child *Widget) {
	ListUnselectChild(l, child)
}

func (l *List) ChildPosition(child *Widget) int {
	return ListChildPosition(l, child)
}

func (l *List) SetSelectionMode(mode T.GtkSelectionMode) {
	ListSetSelectionMode(l, mode)
}

func (l *List) ExtendSelection(scrollType T.GtkScrollType,
	position float32, autoStartSelection T.Gboolean) {
	ListExtendSelection(
		l, scrollType, position, autoStartSelection)
}

func (l *List) StartSelection() { ListStartSelection(l) }

func (l *List) EndSelection() { ListEndSelection(l) }

func (l *List) SelectAll() { ListSelectAll(l) }

func (l *List) UnselectAll() { ListUnselectAll(l) }

func (l *List) ScrollHorizontal(
	scrollType T.GtkScrollType, position float32) {
	ListScrollHorizontal(l, scrollType, position)
}

func (l *List) ScrollVertical(
	scrollType T.GtkScrollType, position float32) {
	ListScrollVertical(l, scrollType, position)
}

func (l *List) ToggleAddMode() { ListToggleAddMode(l) }

func (l *List) ToggleFocusRow() { ListToggleFocusRow(l) }

func (l *List) ToggleRow(item *Widget) {
	ListToggleRow(l, item)
}

func (l *List) UndoSelection() { ListUndoSelection(l) }

func (l *List) EndDragSelection() { ListEndDragSelection(l) }

type ListItem struct{ Item Item }

var (
	ListItemGetType func() T.GType
	ListItemNew     func() *Widget

	ListItemSelect   func(l *ListItem)
	ListItemDeselect func(l *ListItem)
)

func (l *ListItem) ItemSelect() { ListItemSelect(l) }

func (l *ListItem) ItemDeselect() { ListItemDeselect(l) }
