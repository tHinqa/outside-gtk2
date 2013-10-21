// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type TreeModel struct{}

type TreePath struct{}

type (
	TreeView struct {
		Parent T.GtkContainer
		_      *struct{}
	}

	TreeIter struct {
		Stamp     int
		UserData  T.Gpointer
		UserData2 T.Gpointer
		UserData3 T.Gpointer
	}

	TreeSelection struct {
		Parent   T.GObject
		TreeView *TreeView
		Type     T.GtkSelectionMode
		UserFunc TreeSelectionFunc
		UserData T.Gpointer
		Destroy  T.GDestroyNotify
	}

	TreeViewRowSeparatorFunc func(
		model *TreeModel,
		iter *TreeIter,
		data T.Gpointer) T.Gboolean

	TreeViewSearchEqualFunc func(
		model *TreeModel,
		column int,
		key string,
		iter *TreeIter,
		searchData T.Gpointer) T.Gboolean

	TreeViewSearchPositionFunc func(
		tree_view *TreeView,
		search_dialog *Widget,
		user_data T.Gpointer)

	TreeSelectionFunc func(selection *TreeSelection,
		model *TreeModel, path *TreePath,
		pathCurrentlySelected T.Gboolean, data T.Gpointer) T.Gboolean

	TreeCellDataFunc func(
		tree_column *TreeViewColumn,
		cell *CellRenderer,
		treeModel *TreeModel,
		iter *TreeIter,
		data T.Gpointer)

	TreeViewMappingFunc func(
		treeView *TreeView,
		path *TreePath,
		userData T.Gpointer)

	TreeViewColumnDropFunc func(
		treeView *TreeView,
		column *TreeViewColumn,
		prevColumn *TreeViewColumn,
		nextColumn *TreeViewColumn,
		data T.Gpointer) T.Gboolean

	TreeDestroyCountFunc func(
		treeView *TreeView,
		path *TreePath,
		children int,
		userData T.Gpointer)

	TreeIterCompareFunc func(model *TreeModel,
		a, b *TreeIter, userData T.Gpointer) int
)

type TreeViewDropPosition T.Enum

const (
	TREE_VIEW_DROP_BEFORE TreeViewDropPosition = iota
	TREE_VIEW_DROP_AFTER
	TREE_VIEW_DROP_INTO_OR_BEFORE
	TREE_VIEW_DROP_INTO_OR_AFTER
)

type TreeViewGridLines T.Enum

const (
	TREE_VIEW_GRID_LINES_NONE TreeViewGridLines = iota
	TREE_VIEW_GRID_LINES_HORIZONTAL
	TREE_VIEW_GRID_LINES_VERTICAL
	TREE_VIEW_GRID_LINES_BOTH
)

type TreeViewColumnSizing T.Enum

const (
	TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = iota
	TREE_VIEW_COLUMN_AUTOSIZE
	TREE_VIEW_COLUMN_FIXED
)

var (
	TreeViewGetType      func() T.GType
	TreeViewNew          func() *Widget
	TreeViewNewWithModel func(model *TreeModel) *Widget

	treeViewAppendColumn                   func(t *TreeView, column *TreeViewColumn) int
	treeViewCollapseAll                    func(t *TreeView)
	treeViewCollapseRow                    func(t *TreeView, path *TreePath) T.Gboolean
	treeViewColumnsAutosize                func(t *TreeView)
	treeViewConvertBinWindowToTreeCoords   func(t *TreeView, bx, by int, tx, ty *int)
	treeViewConvertBinWindowToWidgetCoords func(t *TreeView, bx, by int, wx, wy *int)
	treeViewConvertTreeToBinWindowCoords   func(t *TreeView, tx, ty int, bx, by *int)
	treeViewConvertTreeToWidgetCoords      func(t *TreeView, tx, ty int, wx, wy *int)
	treeViewConvertWidgetToBinWindowCoords func(t *TreeView, wx, wy int, bx, by *int)
	treeViewConvertWidgetToTreeCoords      func(t *TreeView, wx, wy int, tx, ty *int)
	treeViewCreateRowDragIcon              func(t *TreeView, path *TreePath) *T.GdkPixmap
	treeViewEnableModelDragDest            func(t *TreeView, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction)
	treeViewEnableModelDragSource          func(t *TreeView, startButtonMask T.GdkModifierType, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction)
	treeViewExpandAll                      func(t *TreeView)
	treeViewExpandRow                      func(t *TreeView, path *TreePath, openAll T.Gboolean) T.Gboolean
	treeViewExpandToPath                   func(t *TreeView, path *TreePath)
	treeViewGetBackgroundArea              func(t *TreeView, path *TreePath, column *TreeViewColumn, rect *T.GdkRectangle)
	treeViewGetBinWindow                   func(t *TreeView) *T.GdkWindow
	treeViewGetCellArea                    func(t *TreeView, path *TreePath, column *TreeViewColumn, rect *T.GdkRectangle)
	treeViewGetColumn                      func(t *TreeView, n int) *TreeViewColumn
	treeViewGetColumns                     func(t *TreeView) *T.GList
	treeViewGetCursor                      func(t *TreeView, path **TreePath, focusColumn **TreeViewColumn)
	treeViewGetDestRowAtPos                func(t *TreeView, dragX, dragY int, path **TreePath, pos *TreeViewDropPosition) T.Gboolean
	treeViewGetDragDestRow                 func(t *TreeView, path **TreePath, pos *TreeViewDropPosition)
	treeViewGetEnableSearch                func(t *TreeView) T.Gboolean
	treeViewGetEnableTreeLines             func(t *TreeView) T.Gboolean
	treeViewGetExpanderColumn              func(t *TreeView) *TreeViewColumn
	treeViewGetFixedHeightMode             func(t *TreeView) T.Gboolean
	treeViewGetGridLines                   func(t *TreeView) TreeViewGridLines
	treeViewGetHadjustment                 func(t *TreeView) *Adjustment
	treeViewGetHeadersClickable            func(t *TreeView) T.Gboolean
	treeViewGetHeadersVisible              func(t *TreeView) T.Gboolean
	treeViewGetHoverExpand                 func(t *TreeView) T.Gboolean
	treeViewGetHoverSelection              func(t *TreeView) T.Gboolean
	treeViewGetLevelIndentation            func(t *TreeView) int
	treeViewGetModel                       func(t *TreeView) *TreeModel
	treeViewGetPathAtPos                   func(t *TreeView, x, y int, path **TreePath, column **TreeViewColumn, cellX, cellY *int) T.Gboolean
	treeViewGetReorderable                 func(t *TreeView) T.Gboolean
	treeViewGetRowSeparatorFunc            func(t *TreeView) TreeViewRowSeparatorFunc
	treeViewGetRubberBanding               func(t *TreeView) T.Gboolean
	treeViewGetRulesHint                   func(t *TreeView) T.Gboolean
	treeViewGetSearchColumn                func(t *TreeView) int
	treeViewGetSearchEntry                 func(t *TreeView) *Entry
	treeViewGetSearchEqualFunc             func(t *TreeView) TreeViewSearchEqualFunc
	treeViewGetSearchPositionFunc          func(t *TreeView) TreeViewSearchPositionFunc
	treeViewGetSelection                   func(t *TreeView) *TreeSelection
	treeViewGetShowExpanders               func(t *TreeView) T.Gboolean
	treeViewGetTooltipColumn               func(t *TreeView) int
	treeViewGetTooltipContext              func(t *TreeView, x, y *int, keyboardTip T.Gboolean, model **TreeModel, path **TreePath, iter *TreeIter) T.Gboolean
	treeViewGetVadjustment                 func(t *TreeView) *Adjustment
	treeViewGetVisibleRange                func(t *TreeView, startPath **TreePath, endPath **TreePath) T.Gboolean
	treeViewGetVisibleRect                 func(t *TreeView, visibleRect *T.GdkRectangle)
	treeViewInsertColumn                   func(t *TreeView, column *TreeViewColumn, position int) int
	treeViewInsertColumnWithAttributes     func(t *TreeView, position int, title string, cell *CellRenderer, v ...VArg) int
	treeViewInsertColumnWithDataFunc       func(t *TreeView, position int, title string, cell *CellRenderer, f TreeCellDataFunc, data T.Gpointer, dnotify T.GDestroyNotify) int
	treeViewIsRubberBandingActive          func(t *TreeView) T.Gboolean
	treeViewMapExpandedRows                func(t *TreeView, f TreeViewMappingFunc, data T.Gpointer)
	treeViewMoveColumnAfter                func(t *TreeView, column *TreeViewColumn, baseColumn *TreeViewColumn)
	treeViewRemoveColumn                   func(t *TreeView, column *TreeViewColumn) int
	treeViewRowActivated                   func(t *TreeView, path *TreePath, column *TreeViewColumn)
	treeViewRowExpanded                    func(t *TreeView, path *TreePath) T.Gboolean
	treeViewScrollToCell                   func(t *TreeView, path *TreePath, column *TreeViewColumn, useAlign T.Gboolean, rowAlign, colAlign float32)
	treeViewScrollToPoint                  func(t *TreeView, treeX, treeY int)
	treeViewSetColumnDragFunction          func(t *TreeView, f TreeViewColumnDropFunc, userData T.Gpointer, destroy T.GDestroyNotify)
	treeViewSetCursor                      func(t *TreeView, path *TreePath, focusColumn *TreeViewColumn, startEditing T.Gboolean)
	treeViewSetCursorOnCell                func(t *TreeView, path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing T.Gboolean)
	treeViewSetDestroyCountFunc            func(t *TreeView, f TreeDestroyCountFunc, data T.Gpointer, destroy T.GDestroyNotify)
	treeViewSetDragDestRow                 func(t *TreeView, path *TreePath, pos TreeViewDropPosition)
	treeViewSetEnableSearch                func(t *TreeView, enableSearch T.Gboolean)
	treeViewSetEnableTreeLines             func(t *TreeView, enabled T.Gboolean)
	treeViewSetExpanderColumn              func(t *TreeView, column *TreeViewColumn)
	treeViewSetFixedHeightMode             func(t *TreeView, enable T.Gboolean)
	treeViewSetGridLines                   func(t *TreeView, gridLines TreeViewGridLines)
	treeViewSetHadjustment                 func(t *TreeView, adjustment *Adjustment)
	treeViewSetHeadersClickable            func(t *TreeView, setting T.Gboolean)
	treeViewSetHeadersVisible              func(t *TreeView, headersVisible T.Gboolean)
	treeViewSetHoverExpand                 func(t *TreeView, expand T.Gboolean)
	treeViewSetHoverSelection              func(t *TreeView, hover T.Gboolean)
	treeViewSetLevelIndentation            func(t *TreeView, indentation int)
	treeViewSetModel                       func(t *TreeView, model *TreeModel)
	treeViewSetReorderable                 func(t *TreeView, reorderable T.Gboolean)
	treeViewSetRowSeparatorFunc            func(t *TreeView, f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify)
	treeViewSetRubberBanding               func(t *TreeView, enable T.Gboolean)
	treeViewSetRulesHint                   func(t *TreeView, setting T.Gboolean)
	treeViewSetSearchColumn                func(t *TreeView, column int)
	treeViewSetSearchEntry                 func(t *TreeView, entry *Entry)
	treeViewSetSearchEqualFunc             func(t *TreeView, searchEqualFunc TreeViewSearchEqualFunc, searchUserData T.Gpointer, searchDestroy T.GDestroyNotify)
	treeViewSetSearchPositionFunc          func(t *TreeView, f TreeViewSearchPositionFunc, data T.Gpointer, destroy T.GDestroyNotify)
	treeViewSetShowExpanders               func(t *TreeView, enabled T.Gboolean)
	treeViewSetTooltipCell                 func(t *TreeView, tooltip *T.GtkTooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer)
	treeViewSetTooltipColumn               func(t *TreeView, column int)
	treeViewSetTooltipRow                  func(t *TreeView, tooltip *T.GtkTooltip, path *TreePath)
	treeViewSetVadjustment                 func(t *TreeView, adjustment *Adjustment)
	treeViewTreeToWidgetCoords             func(t *TreeView, tx, ty int, wx, wy *int)
	treeViewUnsetRowsDragDest              func(t *TreeView)
	treeViewUnsetRowsDragSource            func(t *TreeView)
	treeViewWidgetToTreeCoords             func(t *TreeView, wx, wy int, tx, ty *int)
)

func (t *TreeView) AppendColumn(column *TreeViewColumn) int { return treeViewAppendColumn(t, column) }
func (t *TreeView) CollapseAll()                            { treeViewCollapseAll(t) }
func (t *TreeView) CollapseRow(path *TreePath) T.Gboolean   { return treeViewCollapseRow(t, path) }
func (t *TreeView) ColumnsAutosize()                        { treeViewColumnsAutosize(t) }
func (t *TreeView) ConvertBinWindowToTreeCoords(bx, by int, tx, ty *int) {
	treeViewConvertBinWindowToTreeCoords(t, bx, by, tx, ty)
}
func (t *TreeView) ConvertBinWindowToWidgetCoords(bx, by int, wx, wy *int) {
	treeViewConvertBinWindowToWidgetCoords(t, bx, by, wx, wy)
}
func (t *TreeView) ConvertTreeToBinWindowCoords(tx, ty int, bx, by *int) {
	treeViewConvertTreeToBinWindowCoords(t, tx, ty, bx, by)
}
func (t *TreeView) ConvertTreeToWidgetCoords(tx, ty int, wx, wy *int) {
	treeViewConvertTreeToWidgetCoords(t, tx, ty, wx, wy)
}
func (t *TreeView) ConvertWidgetToBinWindowCoords(wx, wy int, bx, by *int) {
	treeViewConvertWidgetToBinWindowCoords(t, wx, wy, bx, by)
}
func (t *TreeView) ConvertWidgetToTreeCoords(wx, wy int, tx, ty *int) {
	treeViewConvertWidgetToTreeCoords(t, wx, wy, tx, ty)
}
func (t *TreeView) CreateRowDragIcon(path *TreePath) *T.GdkPixmap {
	return treeViewCreateRowDragIcon(t, path)
}
func (t *TreeView) EnableModelDragDest(targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction) {
	treeViewEnableModelDragDest(t, targets, nTargets, actions)
}
func (t *TreeView) EnableModelDragSource(startButtonMask T.GdkModifierType, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction) {
	treeViewEnableModelDragSource(t, startButtonMask, targets, nTargets, actions)
}
func (t *TreeView) ExpandAll() { treeViewExpandAll(t) }
func (t *TreeView) ExpandRow(path *TreePath, openAll T.Gboolean) T.Gboolean {
	return treeViewExpandRow(t, path, openAll)
}
func (t *TreeView) ExpandToPath(path *TreePath) { treeViewExpandToPath(t, path) }
func (t *TreeView) GetBackgroundArea(path *TreePath, column *TreeViewColumn, rect *T.GdkRectangle) {
	treeViewGetBackgroundArea(t, path, column, rect)
}
func (t *TreeView) GetBinWindow() *T.GdkWindow { return treeViewGetBinWindow(t) }
func (t *TreeView) GetCellArea(path *TreePath, column *TreeViewColumn, rect *T.GdkRectangle) {
	treeViewGetCellArea(t, path, column, rect)
}
func (t *TreeView) GetColumn(n int) *TreeViewColumn { return treeViewGetColumn(t, n) }
func (t *TreeView) GetColumns() *T.GList            { return treeViewGetColumns(t) }
func (t *TreeView) GetCursor(path **TreePath, focusColumn **TreeViewColumn) {
	treeViewGetCursor(t, path, focusColumn)
}
func (t *TreeView) GetDestRowAtPos(dragX, dragY int, path **TreePath, pos *TreeViewDropPosition) T.Gboolean {
	return treeViewGetDestRowAtPos(t, dragX, dragY, path, pos)
}
func (t *TreeView) GetDragDestRow(path **TreePath, pos *TreeViewDropPosition) {
	treeViewGetDragDestRow(t, path, pos)
}
func (t *TreeView) GetEnableSearch() T.Gboolean        { return treeViewGetEnableSearch(t) }
func (t *TreeView) GetEnableTreeLines() T.Gboolean     { return treeViewGetEnableTreeLines(t) }
func (t *TreeView) GetExpanderColumn() *TreeViewColumn { return treeViewGetExpanderColumn(t) }
func (t *TreeView) GetFixedHeightMode() T.Gboolean     { return treeViewGetFixedHeightMode(t) }
func (t *TreeView) GetGridLines() TreeViewGridLines    { return treeViewGetGridLines(t) }
func (t *TreeView) GetHadjustment() *Adjustment        { return treeViewGetHadjustment(t) }
func (t *TreeView) GetHeadersClickable() T.Gboolean    { return treeViewGetHeadersClickable(t) }
func (t *TreeView) GetHeadersVisible() T.Gboolean      { return treeViewGetHeadersVisible(t) }
func (t *TreeView) GetHoverExpand() T.Gboolean         { return treeViewGetHoverExpand(t) }
func (t *TreeView) GetHoverSelection() T.Gboolean      { return treeViewGetHoverSelection(t) }
func (t *TreeView) GetLevelIndentation() int           { return treeViewGetLevelIndentation(t) }
func (t *TreeView) GetModel() *TreeModel               { return treeViewGetModel(t) }
func (t *TreeView) GetPathAtPos(x, y int, path **TreePath, column **TreeViewColumn, cellX, cellY *int) T.Gboolean {
	return treeViewGetPathAtPos(t, x, y, path, column, cellX, cellY)
}
func (t *TreeView) GetReorderable() T.Gboolean { return treeViewGetReorderable(t) }
func (t *TreeView) GetRowSeparatorFunc() TreeViewRowSeparatorFunc {
	return treeViewGetRowSeparatorFunc(t)
}
func (t *TreeView) GetRubberBanding() T.Gboolean { return treeViewGetRubberBanding(t) }
func (t *TreeView) GetRulesHint() T.Gboolean     { return treeViewGetRulesHint(t) }
func (t *TreeView) GetSearchColumn() int         { return treeViewGetSearchColumn(t) }
func (t *TreeView) GetSearchEntry() *Entry       { return treeViewGetSearchEntry(t) }
func (t *TreeView) GetSearchEqualFunc() TreeViewSearchEqualFunc {
	return treeViewGetSearchEqualFunc(t)
}
func (t *TreeView) GetSearchPositionFunc() TreeViewSearchPositionFunc {
	return treeViewGetSearchPositionFunc(t)
}
func (t *TreeView) GetSelection() *TreeSelection { return treeViewGetSelection(t) }
func (t *TreeView) GetShowExpanders() T.Gboolean { return treeViewGetShowExpanders(t) }
func (t *TreeView) GetTooltipColumn() int        { return treeViewGetTooltipColumn(t) }
func (t *TreeView) GetTooltipContext(x, y *int, keyboardTip T.Gboolean, model **TreeModel, path **TreePath, iter *TreeIter) T.Gboolean {
	return treeViewGetTooltipContext(t, x, y, keyboardTip, model, path, iter)
}
func (t *TreeView) GetVadjustment() *Adjustment { return treeViewGetVadjustment(t) }
func (t *TreeView) GetVisibleRange(startPath **TreePath, endPath **TreePath) T.Gboolean {
	return treeViewGetVisibleRange(t, startPath, endPath)
}
func (t *TreeView) GetVisibleRect(visibleRect *T.GdkRectangle) {
	treeViewGetVisibleRect(t, visibleRect)
}
func (t *TreeView) InsertColumn(column *TreeViewColumn, position int) int {
	return treeViewInsertColumn(t, column, position)
}
func (t *TreeView) InsertColumnWithAttributes(position int, title string, cell *CellRenderer, v ...VArg) int {
	return treeViewInsertColumnWithAttributes(t, position, title, cell, v)
}
func (t *TreeView) InsertColumnWithDataFunc(position int, title string, cell *CellRenderer, f TreeCellDataFunc, data T.Gpointer, dnotify T.GDestroyNotify) int {
	return treeViewInsertColumnWithDataFunc(t, position, title, cell, f, data, dnotify)
}
func (t *TreeView) IsRubberBandingActive() T.Gboolean { return treeViewIsRubberBandingActive(t) }
func (t *TreeView) MapExpandedRows(f TreeViewMappingFunc, data T.Gpointer) {
	treeViewMapExpandedRows(t, f, data)
}
func (t *TreeView) MoveColumnAfter(column *TreeViewColumn, baseColumn *TreeViewColumn) {
	treeViewMoveColumnAfter(t, column, baseColumn)
}
func (t *TreeView) RemoveColumn(column *TreeViewColumn) int { return treeViewRemoveColumn(t, column) }
func (t *TreeView) RowActivated(path *TreePath, column *TreeViewColumn) {
	treeViewRowActivated(t, path, column)
}
func (t *TreeView) RowExpanded(path *TreePath) T.Gboolean { return treeViewRowExpanded(t, path) }
func (t *TreeView) ScrollToCell(path *TreePath, column *TreeViewColumn, useAlign T.Gboolean, rowAlign, colAlign float32) {
	treeViewScrollToCell(t, path, column, useAlign, rowAlign, colAlign)
}
func (t *TreeView) ScrollToPoint(treeX, treeY int) { treeViewScrollToPoint(t, treeX, treeY) }
func (t *TreeView) SetColumnDragFunction(f TreeViewColumnDropFunc, userData T.Gpointer, destroy T.GDestroyNotify) {
	treeViewSetColumnDragFunction(t, f, userData, destroy)
}
func (t *TreeView) SetCursor(path *TreePath, focusColumn *TreeViewColumn, startEditing T.Gboolean) {
	treeViewSetCursor(t, path, focusColumn, startEditing)
}
func (t *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing T.Gboolean) {
	treeViewSetCursorOnCell(t, path, focusColumn, focusCell, startEditing)
}
func (t *TreeView) SetDestroyCountFunc(f TreeDestroyCountFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeViewSetDestroyCountFunc(t, f, data, destroy)
}
func (t *TreeView) SetDragDestRow(path *TreePath, pos TreeViewDropPosition) {
	treeViewSetDragDestRow(t, path, pos)
}
func (t *TreeView) SetEnableSearch(enableSearch T.Gboolean)  { treeViewSetEnableSearch(t, enableSearch) }
func (t *TreeView) SetEnableTreeLines(enabled T.Gboolean)    { treeViewSetEnableTreeLines(t, enabled) }
func (t *TreeView) SetExpanderColumn(column *TreeViewColumn) { treeViewSetExpanderColumn(t, column) }
func (t *TreeView) SetFixedHeightMode(enable T.Gboolean)     { treeViewSetFixedHeightMode(t, enable) }
func (t *TreeView) SetGridLines(gridLines TreeViewGridLines) { treeViewSetGridLines(t, gridLines) }
func (t *TreeView) SetHadjustment(adjustment *Adjustment)    { treeViewSetHadjustment(t, adjustment) }
func (t *TreeView) SetHeadersClickable(setting T.Gboolean)   { treeViewSetHeadersClickable(t, setting) }
func (t *TreeView) SetHeadersVisible(headersVisible T.Gboolean) {
	treeViewSetHeadersVisible(t, headersVisible)
}
func (t *TreeView) SetHoverExpand(expand T.Gboolean)      { treeViewSetHoverExpand(t, expand) }
func (t *TreeView) SetHoverSelection(hover T.Gboolean)    { treeViewSetHoverSelection(t, hover) }
func (t *TreeView) SetLevelIndentation(indentation int)   { treeViewSetLevelIndentation(t, indentation) }
func (t *TreeView) SetModel(model *TreeModel)             { treeViewSetModel(t, model) }
func (t *TreeView) SetReorderable(reorderable T.Gboolean) { treeViewSetReorderable(t, reorderable) }
func (t *TreeView) SetRowSeparatorFunc(f TreeViewRowSeparatorFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeViewSetRowSeparatorFunc(t, f, data, destroy)
}
func (t *TreeView) SetRubberBanding(enable T.Gboolean) { treeViewSetRubberBanding(t, enable) }
func (t *TreeView) SetRulesHint(setting T.Gboolean)    { treeViewSetRulesHint(t, setting) }
func (t *TreeView) SetSearchColumn(column int)         { treeViewSetSearchColumn(t, column) }
func (t *TreeView) SetSearchEntry(entry *Entry)        { treeViewSetSearchEntry(t, entry) }
func (t *TreeView) SetSearchEqualFunc(searchEqualFunc TreeViewSearchEqualFunc, searchUserData T.Gpointer, searchDestroy T.GDestroyNotify) {
	treeViewSetSearchEqualFunc(t, searchEqualFunc, searchUserData, searchDestroy)
}
func (t *TreeView) SetSearchPositionFunc(f TreeViewSearchPositionFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	treeViewSetSearchPositionFunc(t, f, data, destroy)
}
func (t *TreeView) SetShowExpanders(enabled T.Gboolean) { treeViewSetShowExpanders(t, enabled) }
func (t *TreeView) SetTooltipCell(tooltip *T.GtkTooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer) {
	treeViewSetTooltipCell(t, tooltip, path, column, cell)
}
func (t *TreeView) SetTooltipColumn(column int) { treeViewSetTooltipColumn(t, column) }
func (t *TreeView) SetTooltipRow(tooltip *T.GtkTooltip, path *TreePath) {
	treeViewSetTooltipRow(t, tooltip, path)
}
func (t *TreeView) SetVadjustment(adjustment *Adjustment) { treeViewSetVadjustment(t, adjustment) }
func (t *TreeView) TreeToWidgetCoords(tx, ty int, wx, wy *int) {
	treeViewTreeToWidgetCoords(t, tx, ty, wx, wy)
}
func (t *TreeView) UnsetRowsDragDest()   { treeViewUnsetRowsDragDest(t) }
func (t *TreeView) UnsetRowsDragSource() { treeViewUnsetRowsDragSource(t) }
func (t *TreeView) WidgetToTreeCoords(wx, wy int, tx, ty *int) {
	treeViewWidgetToTreeCoords(t, wx, wy, tx, ty)
}

type TreeViewColumn struct{}

var (
	TreeViewColumnGetType           func() T.GType
	TreeViewColumnNew               func() *TreeViewColumn
	TreeViewColumnNewWithAttributes func(title string, cell *CellRenderer, v ...VArg) *TreeViewColumn

	treeViewColumnAddAttribute     func(t *TreeViewColumn, cellRenderer *CellRenderer, attribute string, column int)
	treeViewColumnCellGetPosition  func(t *TreeViewColumn, cellRenderer *CellRenderer, startPos, width *int) T.Gboolean
	treeViewColumnCellGetSize      func(t *TreeViewColumn, cellArea *T.GdkRectangle, xOffset, yOffset, width, height *int)
	treeViewColumnCellIsVisible    func(t *TreeViewColumn) T.Gboolean
	treeViewColumnCellSetCellData  func(t *TreeViewColumn, treeModel *TreeModel, iter *TreeIter, isExpander, isExpanded T.Gboolean)
	treeViewColumnClear            func(t *TreeViewColumn)
	treeViewColumnClearAttributes  func(t *TreeViewColumn, cellRenderer *CellRenderer)
	treeViewColumnClicked          func(t *TreeViewColumn)
	treeViewColumnFocusCell        func(t *TreeViewColumn, cell *CellRenderer)
	treeViewColumnGetAlignment     func(t *TreeViewColumn) float32
	treeViewColumnGetCellRenderers func(t *TreeViewColumn) *T.GList
	treeViewColumnGetClickable     func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetExpand        func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetFixedWidth    func(t *TreeViewColumn) int
	treeViewColumnGetMaxWidth      func(t *TreeViewColumn) int
	treeViewColumnGetMinWidth      func(t *TreeViewColumn) int
	treeViewColumnGetReorderable   func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetResizable     func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetSizing        func(t *TreeViewColumn) TreeViewColumnSizing
	treeViewColumnGetSortColumnId  func(t *TreeViewColumn) int
	treeViewColumnGetSortIndicator func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetSortOrder     func(t *TreeViewColumn) T.GtkSortType
	treeViewColumnGetSpacing       func(t *TreeViewColumn) int
	treeViewColumnGetTitle         func(t *TreeViewColumn) string
	treeViewColumnGetTreeView      func(t *TreeViewColumn) *Widget
	treeViewColumnGetVisible       func(t *TreeViewColumn) T.Gboolean
	treeViewColumnGetWidget        func(t *TreeViewColumn) *Widget
	treeViewColumnGetWidth         func(t *TreeViewColumn) int
	treeViewColumnPackEnd          func(t *TreeViewColumn, cell *CellRenderer, expand T.Gboolean)
	treeViewColumnPackStart        func(t *TreeViewColumn, cell *CellRenderer, expand T.Gboolean)
	treeViewColumnQueueResize      func(t *TreeViewColumn)
	treeViewColumnSetAlignment     func(t *TreeViewColumn, xalign float32)
	treeViewColumnSetAttributes    func(t *TreeViewColumn, cellRenderer *CellRenderer, v ...VArg)
	treeViewColumnSetCellDataFunc  func(t *TreeViewColumn, cellRenderer *CellRenderer, f TreeCellDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
	treeViewColumnSetClickable     func(t *TreeViewColumn, clickable T.Gboolean)
	treeViewColumnSetExpand        func(t *TreeViewColumn, expand T.Gboolean)
	treeViewColumnSetFixedWidth    func(t *TreeViewColumn, fixedWidth int)
	treeViewColumnSetMaxWidth      func(t *TreeViewColumn, maxWidth int)
	treeViewColumnSetMinWidth      func(t *TreeViewColumn, minWidth int)
	treeViewColumnSetReorderable   func(t *TreeViewColumn, reorderable T.Gboolean)
	treeViewColumnSetResizable     func(t *TreeViewColumn, resizable T.Gboolean)
	treeViewColumnSetSizing        func(t *TreeViewColumn, typ TreeViewColumnSizing)
	treeViewColumnSetSortColumnId  func(t *TreeViewColumn, sortColumnId int)
	treeViewColumnSetSortIndicator func(t *TreeViewColumn, setting T.Gboolean)
	treeViewColumnSetSortOrder     func(t *TreeViewColumn, order T.GtkSortType)
	treeViewColumnSetSpacing       func(t *TreeViewColumn, spacing int)
	treeViewColumnSetTitle         func(t *TreeViewColumn, title string)
	treeViewColumnSetVisible       func(t *TreeViewColumn, visible T.Gboolean)
	treeViewColumnSetWidget        func(t *TreeViewColumn, widget *Widget)
)

func (t *TreeViewColumn) AddAttribute(cellRenderer *CellRenderer, attribute string, column int) {
	treeViewColumnAddAttribute(t, cellRenderer, attribute, column)
}
func (t *TreeViewColumn) CellGetPosition(cellRenderer *CellRenderer, startPos *int, width *int) T.Gboolean {
	return treeViewColumnCellGetPosition(t, cellRenderer, startPos, width)
}
func (t *TreeViewColumn) CellGetSize(cellArea *T.GdkRectangle, xOffset, yOffset, width, height *int) {
	treeViewColumnCellGetSize(t, cellArea, xOffset, yOffset, width, height)
}
func (t *TreeViewColumn) CellIsVisible() T.Gboolean { return treeViewColumnCellIsVisible(t) }
func (t *TreeViewColumn) CellSetCellData(treeModel *TreeModel, iter *TreeIter, isExpander, isExpanded T.Gboolean) {
	treeViewColumnCellSetCellData(t, treeModel, iter, isExpander, isExpanded)
}
func (t *TreeViewColumn) Clear() { treeViewColumnClear(t) }
func (t *TreeViewColumn) ClearAttributes(cellRenderer *CellRenderer) {
	treeViewColumnClearAttributes(t, cellRenderer)
}
func (t *TreeViewColumn) Clicked()                        { treeViewColumnClicked(t) }
func (t *TreeViewColumn) FocusCell(cell *CellRenderer)    { treeViewColumnFocusCell(t, cell) }
func (t *TreeViewColumn) GetAlignment() float32           { return treeViewColumnGetAlignment(t) }
func (t *TreeViewColumn) GetCellRenderers() *T.GList      { return treeViewColumnGetCellRenderers(t) }
func (t *TreeViewColumn) GetClickable() T.Gboolean        { return treeViewColumnGetClickable(t) }
func (t *TreeViewColumn) GetExpand() T.Gboolean           { return treeViewColumnGetExpand(t) }
func (t *TreeViewColumn) GetFixedWidth() int              { return treeViewColumnGetFixedWidth(t) }
func (t *TreeViewColumn) GetMaxWidth() int                { return treeViewColumnGetMaxWidth(t) }
func (t *TreeViewColumn) GetMinWidth() int                { return treeViewColumnGetMinWidth(t) }
func (t *TreeViewColumn) GetReorderable() T.Gboolean      { return treeViewColumnGetReorderable(t) }
func (t *TreeViewColumn) GetResizable() T.Gboolean        { return treeViewColumnGetResizable(t) }
func (t *TreeViewColumn) GetSizing() TreeViewColumnSizing { return treeViewColumnGetSizing(t) }
func (t *TreeViewColumn) GetSortColumnId() int            { return treeViewColumnGetSortColumnId(t) }
func (t *TreeViewColumn) GetSortIndicator() T.Gboolean    { return treeViewColumnGetSortIndicator(t) }
func (t *TreeViewColumn) GetSortOrder() T.GtkSortType     { return treeViewColumnGetSortOrder(t) }
func (t *TreeViewColumn) GetSpacing() int                 { return treeViewColumnGetSpacing(t) }
func (t *TreeViewColumn) GetTitle() string                { return treeViewColumnGetTitle(t) }
func (t *TreeViewColumn) GetTreeView() *Widget            { return treeViewColumnGetTreeView(t) }
func (t *TreeViewColumn) GetVisible() T.Gboolean          { return treeViewColumnGetVisible(t) }
func (t *TreeViewColumn) GetWidget() *Widget              { return treeViewColumnGetWidget(t) }
func (t *TreeViewColumn) GetWidth() int                   { return treeViewColumnGetWidth(t) }
func (t *TreeViewColumn) PackStart(cell *CellRenderer, expand T.Gboolean) {
	treeViewColumnPackEnd(t, cell, expand)
}
func (t *TreeViewColumn) QueueResize()                { treeViewColumnQueueResize(t) }
func (t *TreeViewColumn) SetAlignment(xalign float32) { treeViewColumnSetAlignment(t, xalign) }
func (t *TreeViewColumn) SetAttributes(cellRenderer *CellRenderer, v ...VArg) {
	treeViewColumnSetAttributes(t, cellRenderer, v)
}
func (t *TreeViewColumn) SetCellDataFunc(cellRenderer *CellRenderer, f TreeCellDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	treeViewColumnSetCellDataFunc(t, cellRenderer, f, funcData, destroy)
}
func (t *TreeViewColumn) SetClickable(clickable T.Gboolean) { treeViewColumnSetClickable(t, clickable) }
func (t *TreeViewColumn) SetExpand(expand T.Gboolean)       { treeViewColumnSetExpand(t, expand) }
func (t *TreeViewColumn) SetFixedWidth(fixedWidth int)      { treeViewColumnSetFixedWidth(t, fixedWidth) }
func (t *TreeViewColumn) SetMaxWidth(maxWidth int)          { treeViewColumnSetMaxWidth(t, maxWidth) }
func (t *TreeViewColumn) SetMinWidth(minWidth int)          { treeViewColumnSetMinWidth(t, minWidth) }
func (t *TreeViewColumn) SetReorderable(reorderable T.Gboolean) {
	treeViewColumnSetReorderable(t, reorderable)
}
func (t *TreeViewColumn) SetResizable(resizable T.Gboolean)  { treeViewColumnSetResizable(t, resizable) }
func (t *TreeViewColumn) SetSizing(typ TreeViewColumnSizing) { treeViewColumnSetSizing(t, typ) }
func (t *TreeViewColumn) SetSortColumnId(sortColumnId int) {
	treeViewColumnSetSortColumnId(t, sortColumnId)
}
func (t *TreeViewColumn) SetSortIndicator(setting T.Gboolean) {
	treeViewColumnSetSortIndicator(t, setting)
}
func (t *TreeViewColumn) SetSortOrder(order T.GtkSortType) { treeViewColumnSetSortOrder(t, order) }
func (t *TreeViewColumn) SetSpacing(spacing int)           { treeViewColumnSetSpacing(t, spacing) }
func (t *TreeViewColumn) SetTitle(title string)            { treeViewColumnSetTitle(t, title) }
func (t *TreeViewColumn) SetVisible(visible T.Gboolean)    { treeViewColumnSetVisible(t, visible) }
func (t *TreeViewColumn) SetWidget(widget *Widget)         { treeViewColumnSetWidget(t, widget) }
