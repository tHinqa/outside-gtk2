// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type IconFactory struct {
	Parent T.GObject
	Icons  *T.GHashTable
}

var (
	IconFactoryGetType func() T.GType
	IconFactoryNew     func() *IconFactory

	IconFactoryLookupDefault func(stockId string) *IconSet

	iconFactoryAdd           func(i *IconFactory, stockId string, iconSet *IconSet)
	iconFactoryAddDefault    func(i *IconFactory)
	iconFactoryLookup        func(i *IconFactory, stockId string) *IconSet
	iconFactoryRemoveDefault func(i *IconFactory)
)

func (i *IconFactory) Add(stockId string, iconSet *IconSet) { iconFactoryAdd(i, stockId, iconSet) }
func (i *IconFactory) AddDefault()                          { iconFactoryAddDefault(i) }
func (i *IconFactory) Lookup(stockId string) *IconSet       { return iconFactoryLookup(i, stockId) }
func (i *IconFactory) RemoveDefault()                       { iconFactoryRemoveDefault(i) }

type IconInfo struct{}

var (
	IconInfoGetType      func() T.GType
	IconInfoNewForPixbuf func(iconTheme *IconTheme, pixbuf *T.GdkPixbuf) *IconInfo

	iconInfoCopy              func(i *IconInfo) *IconInfo
	iconInfoFree              func(i *IconInfo)
	iconInfoGetAttachPoints   func(i *IconInfo, points **T.GdkPoint, nPoints *int) T.Gboolean
	iconInfoGetBaseSize       func(i *IconInfo) int
	iconInfoGetBuiltinPixbuf  func(i *IconInfo) *T.GdkPixbuf
	iconInfoGetDisplayName    func(i *IconInfo) string
	iconInfoGetEmbeddedRect   func(i *IconInfo, rectangle *T.GdkRectangle) T.Gboolean
	iconInfoGetFilename       func(i *IconInfo) string
	iconInfoLoadIcon          func(i *IconInfo, error **T.GError) *T.GdkPixbuf
	iconInfoSetRawCoordinates func(i *IconInfo, rawCoordinates T.Gboolean)
)

func (i *IconInfo) Copy() *IconInfo { return iconInfoCopy(i) }
func (i *IconInfo) Free()           { iconInfoFree(i) }
func (i *IconInfo) GetAttachPoints(points **T.GdkPoint, nPoints *int) T.Gboolean {
	return iconInfoGetAttachPoints(i, points, nPoints)
}
func (i *IconInfo) GetBaseSize() int               { return iconInfoGetBaseSize(i) }
func (i *IconInfo) GetBuiltinPixbuf() *T.GdkPixbuf { return iconInfoGetBuiltinPixbuf(i) }
func (i *IconInfo) GetDisplayName() string         { return iconInfoGetDisplayName(i) }
func (i *IconInfo) GetEmbeddedRect(rectangle *T.GdkRectangle) T.Gboolean {
	return iconInfoGetEmbeddedRect(i, rectangle)
}
func (i *IconInfo) GetFilename() string                  { return iconInfoGetFilename(i) }
func (i *IconInfo) LoadIcon(err **T.GError) *T.GdkPixbuf { return iconInfoLoadIcon(i, err) }
func (i *IconInfo) SetRawCoordinates(rawCoordinates T.Gboolean) {
	iconInfoSetRawCoordinates(i, rawCoordinates)
}

type IconSet struct{}

var (
	IconSetGetType       func() T.GType
	IconSetNew           func() *IconSet
	IconSetNewFromPixbuf func(pixbuf *T.GdkPixbuf) *IconSet

	iconSetAddSource  func(i *IconSet, source *IconSource)
	iconSetCopy       func(i *IconSet) *IconSet
	iconSetGetSizes   func(i *IconSet, sizes **IconSize, nSizes *int)
	iconSetRef        func(i *IconSet) *IconSet
	iconSetRenderIcon func(i *IconSet, style *T.GtkStyle, direction T.GtkTextDirection, state T.GtkStateType, size IconSize, widget *Widget, detail string) *T.GdkPixbuf
	iconSetUnref      func(i *IconSet)
)

func (i *IconSet) AddSource(source *IconSource)           { iconSetAddSource(i, source) }
func (i *IconSet) Copy() *IconSet                         { return iconSetCopy(i) }
func (i *IconSet) GetSizes(sizes **IconSize, nSizes *int) { iconSetGetSizes(i, sizes, nSizes) }
func (i *IconSet) Ref() *IconSet                          { return iconSetRef(i) }
func (i *IconSet) RenderIcon(style *T.GtkStyle, direction T.GtkTextDirection, state T.GtkStateType, size IconSize, widget *Widget, detail string) *T.GdkPixbuf {
	return iconSetRenderIcon(i, style, direction, state, size, widget, detail)
}
func (i *IconSet) Unref() { iconSetUnref(i) }

type IconSize T.Enum

const (
	ICON_SIZE_INVALID IconSize = iota
	ICON_SIZE_MENU
	ICON_SIZE_SMALL_TOOLBAR
	ICON_SIZE_LARGE_TOOLBAR
	ICON_SIZE_BUTTON
	ICON_SIZE_DND
	ICON_SIZE_DIALOG
)

var (
	IconSizeGetType func() T.GType

	IconSizeFromName      func(name string) IconSize
	IconSizeRegister      func(name string, width int, height int) IconSize
	IconSizeRegisterAlias func(alias string, target IconSize)

	iconSizeGetName func(i IconSize) string
	iconSizeLookup  func(i IconSize, width *int, height *int) T.Gboolean
)

func (i IconSize) GetName() string                           { return iconSizeGetName(i) }
func (i IconSize) Lookup(width *int, height *int) T.Gboolean { return iconSizeLookup(i, width, height) }

type IconSource struct{}

var (
	IconSourceGetType func() T.GType
	IconSourceNew     func() *IconSource

	iconSourceCopy                   func(i *IconSource) *IconSource
	iconSourceFree                   func(i *IconSource)
	iconSourceGetDirection           func(i *IconSource) T.GtkTextDirection
	iconSourceGetDirectionWildcarded func(i *IconSource) T.Gboolean
	iconSourceGetFilename            func(i *IconSource) string
	iconSourceGetIconName            func(i *IconSource) string
	iconSourceGetPixbuf              func(i *IconSource) *T.GdkPixbuf
	iconSourceGetSize                func(i *IconSource) IconSize
	iconSourceGetSizeWildcarded      func(i *IconSource) T.Gboolean
	iconSourceGetState               func(i *IconSource) T.GtkStateType
	iconSourceGetStateWildcarded     func(i *IconSource) T.Gboolean
	iconSourceSetDirection           func(i *IconSource, direction T.GtkTextDirection)
	iconSourceSetDirectionWildcarded func(i *IconSource, setting T.Gboolean)
	iconSourceSetFilename            func(i *IconSource, filename string)
	iconSourceSetIconName            func(i *IconSource, iconName string)
	iconSourceSetPixbuf              func(i *IconSource, pixbuf *T.GdkPixbuf)
	iconSourceSetSize                func(i *IconSource, size IconSize)
	iconSourceSetSizeWildcarded      func(i *IconSource, setting T.Gboolean)
	iconSourceSetState               func(i *IconSource, state T.GtkStateType)
	iconSourceSetStateWildcarded     func(i *IconSource, setting T.Gboolean)
)

func (i *IconSource) Copy() *IconSource                         { return iconSourceCopy(i) }
func (i *IconSource) Free()                                     { iconSourceFree(i) }
func (i *IconSource) GetDirection() T.GtkTextDirection          { return iconSourceGetDirection(i) }
func (i *IconSource) GetDirectionWildcarded() T.Gboolean        { return iconSourceGetDirectionWildcarded(i) }
func (i *IconSource) GetFilename() string                       { return iconSourceGetFilename(i) }
func (i *IconSource) GetIconName() string                       { return iconSourceGetIconName(i) }
func (i *IconSource) GetPixbuf() *T.GdkPixbuf                   { return iconSourceGetPixbuf(i) }
func (i *IconSource) GetSize() IconSize                         { return iconSourceGetSize(i) }
func (i *IconSource) GetSizeWildcarded() T.Gboolean             { return iconSourceGetSizeWildcarded(i) }
func (i *IconSource) GetState() T.GtkStateType                  { return iconSourceGetState(i) }
func (i *IconSource) GetStateWildcarded() T.Gboolean            { return iconSourceGetStateWildcarded(i) }
func (i *IconSource) SetDirection(direction T.GtkTextDirection) { iconSourceSetDirection(i, direction) }
func (i *IconSource) SetDirectionWildcarded(setting T.Gboolean) {
	iconSourceSetDirectionWildcarded(i, setting)
}
func (i *IconSource) SetFilename(filename string)           { iconSourceSetFilename(i, filename) }
func (i *IconSource) SetIconName(iconName string)           { iconSourceSetIconName(i, iconName) }
func (i *IconSource) SetPixbuf(pixbuf *T.GdkPixbuf)         { iconSourceSetPixbuf(i, pixbuf) }
func (i *IconSource) SetSize(size IconSize)                 { iconSourceSetSize(i, size) }
func (i *IconSource) SetSizeWildcarded(setting T.Gboolean)  { iconSourceSetSizeWildcarded(i, setting) }
func (i *IconSource) SetState(state T.GtkStateType)         { iconSourceSetState(i, state) }
func (i *IconSource) SetStateWildcarded(setting T.Gboolean) { iconSourceSetStateWildcarded(i, setting) }

type IconTheme struct {
	Parent T.GObject
	_      *struct{}
}

type IconLookupFlags T.Enum

const (
	ICON_LOOKUP_NO_SVG IconLookupFlags = 1 << iota
	ICON_LOOKUP_FORCE_SVG
	ICON_LOOKUP_USE_BUILTIN
	ICON_LOOKUP_GENERIC_FALLBACK
	ICON_LOOKUP_FORCE_SIZE
)

var (
	IconThemeGetType func() T.GType
	IconThemeNew     func() *IconTheme

	IconThemeErrorGetType func() T.GType
	IconThemeErrorQuark   func() T.GQuark
	IconThemeGetDefault   func() *IconTheme

	IconThemeGetForScreen   func(screen *T.GdkScreen) *IconTheme
	IconThemeAddBuiltinIcon func(iconName string, size int, pixbuf *T.GdkPixbuf)

	iconThemeAppendSearchPath   func(i *IconTheme, path string)
	iconThemeChooseIcon         func(i *IconTheme, iconNames **T.Gchar, size int, flags IconLookupFlags) *IconInfo
	iconThemeGetExampleIconName func(i *IconTheme) string
	iconThemeGetIconSizes       func(i *IconTheme, iconName string) *int
	iconThemeGetSearchPath      func(i *IconTheme, path ***T.Gchar, nElements *int)
	iconThemeHasIcon            func(i *IconTheme, iconName string) T.Gboolean
	iconThemeListContexts       func(i *IconTheme) *T.GList
	iconThemeListIcons          func(i *IconTheme, context string) *T.GList
	iconThemeLoadIcon           func(i *IconTheme, iconName string, size int, flags IconLookupFlags, error **T.GError) *T.GdkPixbuf
	iconThemeLookupByGicon      func(i *IconTheme, icon *T.GIcon, size int, flags IconLookupFlags) *IconInfo
	iconThemeLookupIcon         func(i *IconTheme, iconName string, size int, flags IconLookupFlags) *IconInfo
	iconThemePrependSearchPath  func(i *IconTheme, path string)
	iconThemeRescanIfNeeded     func(i *IconTheme) T.Gboolean
	iconThemeSetCustomTheme     func(i *IconTheme, themeName string)
	iconThemeSetScreen          func(i *IconTheme, screen *T.GdkScreen)
	iconThemeSetSearchPath      func(i *IconTheme, path **T.Gchar, nElements int)
)

func (i *IconTheme) AppendSearchPath(path string) { iconThemeAppendSearchPath(i, path) }
func (i *IconTheme) ChooseIcon(iconNames **T.Gchar, size int, flags IconLookupFlags) *IconInfo {
	return iconThemeChooseIcon(i, iconNames, size, flags)
}
func (i *IconTheme) GetExampleIconName() string        { return iconThemeGetExampleIconName(i) }
func (i *IconTheme) GetIconSizes(iconName string) *int { return iconThemeGetIconSizes(i, iconName) }
func (i *IconTheme) GetSearchPath(path ***T.Gchar, nElements *int) {
	iconThemeGetSearchPath(i, path, nElements)
}
func (i *IconTheme) HasIcon(iconName string) T.Gboolean { return iconThemeHasIcon(i, iconName) }
func (i *IconTheme) ListContexts() *T.GList             { return iconThemeListContexts(i) }
func (i *IconTheme) ListIcons(context string) *T.GList  { return iconThemeListIcons(i, context) }
func (i *IconTheme) LoadIcon(iconName string, size int, flags IconLookupFlags, err **T.GError) *T.GdkPixbuf {
	return iconThemeLoadIcon(i, iconName, size, flags, err)
}
func (i *IconTheme) LookupByGicon(icon *T.GIcon, size int, flags IconLookupFlags) *IconInfo {
	return iconThemeLookupByGicon(i, icon, size, flags)
}
func (i *IconTheme) LookupIcon(iconName string, size int, flags IconLookupFlags) *IconInfo {
	return iconThemeLookupIcon(i, iconName, size, flags)
}
func (i *IconTheme) PrependSearchPath(path string)   { iconThemePrependSearchPath(i, path) }
func (i *IconTheme) RescanIfNeeded() T.Gboolean      { return iconThemeRescanIfNeeded(i) }
func (i *IconTheme) SetCustomTheme(themeName string) { iconThemeSetCustomTheme(i, themeName) }
func (i *IconTheme) SetScreen(screen *T.GdkScreen)   { iconThemeSetScreen(i, screen) }
func (i *IconTheme) SetSearchPath(path **T.Gchar, nElements int) {
	iconThemeSetSearchPath(i, path, nElements)
}

type (
	IconView struct {
		Parent T.GtkContainer
		_      *struct{}
	}

	IconViewForeachFunc func(icon_view *IconView,
		path *TreePath, data T.Gpointer)
)

type IconViewDropPosition T.Enum

const (
	ICON_VIEW_NO_DROP IconViewDropPosition = iota
	ICON_VIEW_DROP_INTO
	ICON_VIEW_DROP_LEFT
	ICON_VIEW_DROP_RIGHT
	ICON_VIEW_DROP_ABOVE
	ICON_VIEW_DROP_BELOW
)

var (
	IconViewGetType      func() T.GType
	IconViewNew          func() *Widget
	IconViewNewWithModel func(model *TreeModel) *Widget

	IconViewDropPositionGetType func() T.GType

	iconViewConvertWidgetToBinWindowCoords func(i *IconView, wx int, wy int, bx *int, by *int)
	iconViewCreateDragIcon                 func(i *IconView, path *TreePath) *T.GdkPixmap
	iconViewEnableModelDragDest            func(i *IconView, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction)
	iconViewEnableModelDragSource          func(i *IconView, startButtonMask T.GdkModifierType, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction)
	iconViewGetColumns                     func(i *IconView) int
	iconViewGetColumnSpacing               func(i *IconView) int
	iconViewGetCursor                      func(i *IconView, path **TreePath, cell **CellRenderer) T.Gboolean
	iconViewGetDestItemAtPos               func(i *IconView, dragX, dragY int, path **TreePath, pos *IconViewDropPosition) T.Gboolean
	iconViewGetDragDestItem                func(i *IconView, path **TreePath, pos *IconViewDropPosition)
	iconViewGetItemAtPos                   func(i *IconView, x, y int, path **TreePath, cell **CellRenderer) T.Gboolean
	iconViewGetItemColumn                  func(i *IconView, path *TreePath) int
	iconViewGetItemOrientation             func(i *IconView) T.GtkOrientation
	iconViewGetItemPadding                 func(i *IconView) int
	iconViewGetItemRow                     func(i *IconView, path *TreePath) int
	iconViewGetItemWidth                   func(i *IconView) int
	iconViewGetMargin                      func(i *IconView) int
	iconViewGetMarkupColumn                func(i *IconView) int
	iconViewGetModel                       func(i *IconView) *TreeModel
	iconViewGetOrientation                 func(i *IconView) T.GtkOrientation
	iconViewGetPathAtPos                   func(i *IconView, x, y int) *TreePath
	iconViewGetPixbufColumn                func(i *IconView) int
	iconViewGetReorderable                 func(i *IconView) T.Gboolean
	iconViewGetRowSpacing                  func(i *IconView) int
	iconViewGetSelectedItems               func(i *IconView) *T.GList
	iconViewGetSelectionMode               func(i *IconView) T.GtkSelectionMode
	iconViewGetSpacing                     func(i *IconView) int
	iconViewGetTextColumn                  func(i *IconView) int
	iconViewGetTooltipColumn               func(i *IconView) int
	iconViewGetTooltipContext              func(i *IconView, x *int, y *int, keyboardTip T.Gboolean, model **TreeModel, path **TreePath, iter *TreeIter) T.Gboolean
	iconViewGetVisibleRange                func(i *IconView, startPath **TreePath, endPath **TreePath) T.Gboolean
	iconViewItemActivated                  func(i *IconView, path *TreePath)
	iconViewPathIsSelected                 func(i *IconView, path *TreePath) T.Gboolean
	iconViewScrollToPath                   func(i *IconView, path *TreePath, useAlign T.Gboolean, rowAlign, colAlign float32)
	iconViewSelectAll                      func(i *IconView)
	iconViewSelectedForeach                func(i *IconView, f IconViewForeachFunc, data T.Gpointer)
	iconViewSelectPath                     func(i *IconView, path *TreePath)
	iconViewSetColumns                     func(i *IconView, columns int)
	iconViewSetColumnSpacing               func(i *IconView, columnSpacing int)
	iconViewSetCursor                      func(i *IconView, path *TreePath, cell *CellRenderer, startEditing T.Gboolean)
	iconViewSetDragDestItem                func(i *IconView, path *TreePath, pos IconViewDropPosition)
	iconViewSetItemOrientation             func(i *IconView, orientation T.GtkOrientation)
	iconViewSetItemPadding                 func(i *IconView, itemPadding int)
	iconViewSetItemWidth                   func(i *IconView, itemWidth int)
	iconViewSetMargin                      func(i *IconView, margin int)
	iconViewSetMarkupColumn                func(i *IconView, column int)
	iconViewSetModel                       func(i *IconView, model *TreeModel)
	iconViewSetOrientation                 func(i *IconView, orientation T.GtkOrientation)
	iconViewSetPixbufColumn                func(i *IconView, column int)
	iconViewSetReorderable                 func(i *IconView, reorderable T.Gboolean)
	iconViewSetRowSpacing                  func(i *IconView, rowSpacing int)
	iconViewSetSelectionMode               func(i *IconView, mode T.GtkSelectionMode)
	iconViewSetSpacing                     func(i *IconView, spacing int)
	iconViewSetTextColumn                  func(i *IconView, column int)
	iconViewSetTooltipCell                 func(i *IconView, tooltip *T.GtkTooltip, path *TreePath, cell *CellRenderer)
	iconViewSetTooltipColumn               func(i *IconView, column int)
	iconViewSetTooltipItem                 func(i *IconView, tooltip *T.GtkTooltip, path *TreePath)
	iconViewUnselectAll                    func(i *IconView)
	iconViewUnselectPath                   func(i *IconView, path *TreePath)
	iconViewUnsetModelDragDest             func(i *IconView)
	iconViewUnsetModelDragSource           func(i *IconView)
)

func (i *IconView) ConvertWidgetToBinWindowCoords(wx, wy int, bx, by *int) {
	iconViewConvertWidgetToBinWindowCoords(i, wx, wy, bx, by)
}
func (i *IconView) CreateDragIcon(path *TreePath) *T.GdkPixmap {
	return iconViewCreateDragIcon(i, path)
}
func (i *IconView) EnableModelDragDest(targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction) {
	iconViewEnableModelDragDest(i, targets, nTargets, actions)
}
func (i *IconView) EnableModelDragSource(startButtonMask T.GdkModifierType, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction) {
	iconViewEnableModelDragSource(i, startButtonMask, targets, nTargets, actions)
}
func (i *IconView) GetColumns() int       { return iconViewGetColumns(i) }
func (i *IconView) GetColumnSpacing() int { return iconViewGetColumnSpacing(i) }
func (i *IconView) GetCursor(path **TreePath, cell **CellRenderer) T.Gboolean {
	return iconViewGetCursor(i, path, cell)
}
func (i *IconView) GetDestItemAtPos(dragX, dragY int, path **TreePath, pos *IconViewDropPosition) T.Gboolean {
	return iconViewGetDestItemAtPos(i, dragX, dragY, path, pos)
}
func (i *IconView) GetDragDestItem(path **TreePath, pos *IconViewDropPosition) {
	iconViewGetDragDestItem(i, path, pos)
}
func (i *IconView) GetItemAtPos(x, y int, path **TreePath, cell **CellRenderer) T.Gboolean {
	return iconViewGetItemAtPos(i, x, y, path, cell)
}
func (i *IconView) GetItemColumn(path *TreePath) int     { return iconViewGetItemColumn(i, path) }
func (i *IconView) GetItemOrientation() T.GtkOrientation { return iconViewGetItemOrientation(i) }
func (i *IconView) GetItemPadding() int                  { return iconViewGetItemPadding(i) }
func (i *IconView) GetItemRow(path *TreePath) int        { return iconViewGetItemRow(i, path) }
func (i *IconView) GetItemWidth() int                    { return iconViewGetItemWidth(i) }
func (i *IconView) GetMargin() int                       { return iconViewGetMargin(i) }
func (i *IconView) GetMarkupColumn() int                 { return iconViewGetMarkupColumn(i) }
func (i *IconView) GetModel() *TreeModel                 { return iconViewGetModel(i) }
func (i *IconView) GetOrientation() T.GtkOrientation     { return iconViewGetOrientation(i) }
func (i *IconView) GetPathAtPos(x, y int) *TreePath      { return iconViewGetPathAtPos(i, x, y) }
func (i *IconView) GetPixbufColumn() int                 { return iconViewGetPixbufColumn(i) }
func (i *IconView) GetReorderable() T.Gboolean           { return iconViewGetReorderable(i) }
func (i *IconView) GetRowSpacing() int                   { return iconViewGetRowSpacing(i) }
func (i *IconView) GetSelectedItems() *T.GList           { return iconViewGetSelectedItems(i) }
func (i *IconView) GetSelectionMode() T.GtkSelectionMode { return iconViewGetSelectionMode(i) }
func (i *IconView) GetSpacing() int                      { return iconViewGetSpacing(i) }
func (i *IconView) GetTextColumn() int                   { return iconViewGetTextColumn(i) }
func (i *IconView) GetTooltipColumn() int                { return iconViewGetTooltipColumn(i) }
func (i *IconView) GetTooltipContext(x, y *int, keyboardTip T.Gboolean, model **TreeModel, path **TreePath, iter *TreeIter) T.Gboolean {
	return iconViewGetTooltipContext(i, x, y, keyboardTip, model, path, iter)
}
func (i *IconView) GetVisibleRange(startPath **TreePath, endPath **TreePath) T.Gboolean {
	return iconViewGetVisibleRange(i, startPath, endPath)
}
func (i *IconView) ItemActivated(path *TreePath) { iconViewItemActivated(i, path) }
func (i *IconView) PathIsSelected(path *TreePath) T.Gboolean {
	return iconViewPathIsSelected(i, path)
}
func (i *IconView) ScrollToPath(path *TreePath, useAlign T.Gboolean, rowAlign, colAlign float32) {
	iconViewScrollToPath(i, path, useAlign, rowAlign, colAlign)
}
func (i *IconView) SelectAll() { iconViewSelectAll(i) }
func (i *IconView) SelectedForeach(f IconViewForeachFunc, data T.Gpointer) {
	iconViewSelectedForeach(i, f, data)
}
func (i *IconView) SelectPath(path *TreePath)          { iconViewSelectPath(i, path) }
func (i *IconView) SetColumns(columns int)             { iconViewSetColumns(i, columns) }
func (i *IconView) SetColumnSpacing(columnSpacing int) { iconViewSetColumnSpacing(i, columnSpacing) }
func (i *IconView) SetCursor(path *TreePath, cell *CellRenderer, startEditing T.Gboolean) {
	iconViewSetCursor(i, path, cell, startEditing)
}
func (i *IconView) SetDragDestItem(path *TreePath, pos IconViewDropPosition) {
	iconViewSetDragDestItem(i, path, pos)
}
func (i *IconView) SetItemOrientation(orientation T.GtkOrientation) {
	iconViewSetItemOrientation(i, orientation)
}
func (i *IconView) SetItemPadding(itemPadding int) { iconViewSetItemPadding(i, itemPadding) }
func (i *IconView) SetItemWidth(itemWidth int)     { iconViewSetItemWidth(i, itemWidth) }
func (i *IconView) SetMargin(margin int)           { iconViewSetMargin(i, margin) }
func (i *IconView) SetMarkupColumn(column int)     { iconViewSetMarkupColumn(i, column) }
func (i *IconView) SetModel(model *TreeModel)      { iconViewSetModel(i, model) }
func (i *IconView) SetOrientation(orientation T.GtkOrientation) {
	iconViewSetOrientation(i, orientation)
}
func (i *IconView) SetPixbufColumn(column int)               { iconViewSetPixbufColumn(i, column) }
func (i *IconView) SetReorderable(reorderable T.Gboolean)    { iconViewSetReorderable(i, reorderable) }
func (i *IconView) SetRowSpacing(rowSpacing int)             { iconViewSetRowSpacing(i, rowSpacing) }
func (i *IconView) SetSelectionMode(mode T.GtkSelectionMode) { iconViewSetSelectionMode(i, mode) }
func (i *IconView) SetSpacing(spacing int)                   { iconViewSetSpacing(i, spacing) }
func (i *IconView) SetTextColumn(column int)                 { iconViewSetTextColumn(i, column) }
func (i *IconView) SetTooltipCell(tooltip *T.GtkTooltip, path *TreePath, cell *CellRenderer) {
	iconViewSetTooltipCell(i, tooltip, path, cell)
}
func (i *IconView) SetTooltipColumn(column int) { iconViewSetTooltipColumn(i, column) }
func (i *IconView) SetTooltipItem(tooltip *T.GtkTooltip, path *TreePath) {
	iconViewSetTooltipItem(i, tooltip, path)
}
func (i *IconView) UnselectAll()                { iconViewUnselectAll(i) }
func (i *IconView) UnselectPath(path *TreePath) { iconViewUnselectPath(i, path) }
func (i *IconView) UnsetModelDragDest()         { iconViewUnsetModelDragDest(i) }
func (i *IconView) UnsetModelDragSource()       { iconViewUnsetModelDragSource(i) }

type Image struct{}

/*
type Image  struct  {
	misc  GtkMisc
	storageType  GtkImageType
	union {
		pixmap  GtkImagePixmapData
		image  GtkImageImageData
		pixbuf  GtkImagePixbufData
		stock  GtkImageStockData
		iconSet  GtkImageIconSetData
		anim  GtkImageAnimationData
		name  GtkImageIconNameData
		gicon  GtkImageGIconData
	}
	mask  *GdkBitmap
	iconSize  GtkIconSize
}
*/

type ImageType T.Enum

const (
	IMAGE_EMPTY ImageType = iota
	IMAGE_PIXMAP
	IMAGE_IMAGE
	IMAGE_PIXBUF
	IMAGE_STOCK
	IMAGE_ICON_SET
	IMAGE_ANIMATION
	IMAGE_ICON_NAME
	IMAGE_GICON
)

var (
	ImageGetType          func() T.GType
	ImageNew              func() *Widget
	ImageNewFromPixmap    func(pixmap *T.GdkPixmap, mask *T.GdkBitmap) *Widget
	ImageNewFromImage     func(image *T.GdkImage, mask *T.GdkBitmap) *Widget
	ImageNewFromFile      func(filename string) *Widget
	ImageNewFromPixbuf    func(pixbuf *T.GdkPixbuf) *Widget
	ImageNewFromStock     func(stockId string, size IconSize) *Widget
	ImageNewFromIconSet   func(iconSet *IconSet, size IconSize) *Widget
	ImageNewFromAnimation func(animation *T.GdkPixbufAnimation) *Widget
	ImageNewFromIconName  func(iconName string, size IconSize) *Widget
	ImageNewFromGicon     func(icon *T.GIcon, size IconSize) *Widget

	imageClear            func(i *Image)
	imageGet              func(i *Image, val **T.GdkImage, mask **T.GdkBitmap)
	imageGetAnimation     func(i *Image) *T.GdkPixbufAnimation
	imageGetGicon         func(i *Image, gicon **T.GIcon, size *IconSize)
	imageGetIconName      func(i *Image, iconName **T.Gchar, size *IconSize)
	imageGetIconSet       func(i *Image, iconSet **IconSet, size *IconSize)
	imageGetImage         func(i *Image, gdkImage **T.GdkImage, mask **T.GdkBitmap)
	imageGetPixbuf        func(i *Image) *T.GdkPixbuf
	imageGetPixelSize     func(i *Image) int
	imageGetPixmap        func(i *Image, pixmap **T.GdkPixmap, mask **T.GdkBitmap)
	imageGetStock         func(i *Image, stockId **T.Gchar, size *IconSize)
	imageGetStorageType   func(i *Image) ImageType
	imageSet              func(i *Image, val *T.GdkImage, mask *T.GdkBitmap)
	imageSetFromAnimation func(i *Image, animation *T.GdkPixbufAnimation)
	imageSetFromFile      func(i *Image, filename string)
	imageSetFromGicon     func(i *Image, icon *T.GIcon, size IconSize)
	imageSetFromIconName  func(i *Image, iconName string, size IconSize)
	imageSetFromIconSet   func(i *Image, iconSet *IconSet, size IconSize)
	imageSetFromImage     func(i *Image, gdkImage *T.GdkImage, mask *T.GdkBitmap)
	imageSetFromPixbuf    func(i *Image, pixbuf *T.GdkPixbuf)
	imageSetFromPixmap    func(i *Image, pixmap *T.GdkPixmap, mask *T.GdkBitmap)
	imageSetFromStock     func(i *Image, stockId string, size IconSize)
	imageSetPixelSize     func(i *Image, pixelSize int)
)

func (i *Image) Clear()                                             { imageClear(i) }
func (i *Image) Get(val **T.GdkImage, mask **T.GdkBitmap)           { imageGet(i, val, mask) }
func (i *Image) GetAnimation() *T.GdkPixbufAnimation                { return imageGetAnimation(i) }
func (i *Image) GetGicon(gicon **T.GIcon, size *IconSize)           { imageGetGicon(i, gicon, size) }
func (i *Image) GetIconName(iconName **T.Gchar, size *IconSize)     { imageGetIconName(i, iconName, size) }
func (i *Image) GetIconSet(iconSet **IconSet, size *IconSize)       { imageGetIconSet(i, iconSet, size) }
func (i *Image) GetImage(gdkImage **T.GdkImage, mask **T.GdkBitmap) { imageGetImage(i, gdkImage, mask) }
func (i *Image) GetPixbuf() *T.GdkPixbuf                            { return imageGetPixbuf(i) }
func (i *Image) GetPixelSize() int                                  { return imageGetPixelSize(i) }
func (i *Image) GetPixmap(pixmap **T.GdkPixmap, mask **T.GdkBitmap) { imageGetPixmap(i, pixmap, mask) }
func (i *Image) GetStock(stockId **T.Gchar, size *IconSize)         { imageGetStock(i, stockId, size) }
func (i *Image) GetStorageType() ImageType                          { return imageGetStorageType(i) }
func (i *Image) Set(val *T.GdkImage, mask *T.GdkBitmap)             { imageSet(i, val, mask) }
func (i *Image) SetFromAnimation(animation *T.GdkPixbufAnimation)   { imageSetFromAnimation(i, animation) }
func (i *Image) SetFromFile(filename string)                        { imageSetFromFile(i, filename) }
func (i *Image) SetFromGicon(icon *T.GIcon, size IconSize)          { imageSetFromGicon(i, icon, size) }
func (i *Image) SetFromIconName(iconName string, size IconSize) {
	imageSetFromIconName(i, iconName, size)
}
func (i *Image) SetFromIconSet(iconSet *IconSet, size IconSize) { imageSetFromIconSet(i, iconSet, size) }
func (i *Image) SetFromImage(gdkImage *T.GdkImage, mask *T.GdkBitmap) {
	imageSetFromImage(i, gdkImage, mask)
}
func (i *Image) SetFromPixbuf(pixbuf *T.GdkPixbuf) { imageSetFromPixbuf(i, pixbuf) }
func (i *Image) SetFromPixmap(pixmap *T.GdkPixmap, mask *T.GdkBitmap) {
	imageSetFromPixmap(i, pixmap, mask)
}
func (i *Image) SetFromStock(stockId string, size IconSize) { imageSetFromStock(i, stockId, size) }
func (i *Image) SetPixelSize(pixelSize int)                 { imageSetPixelSize(i, pixelSize) }

type ImageMenuItem struct {
	MenuItem MenuItem
	Image    *Widget
}

var (
	ImageMenuItemGetType         func() T.GType
	ImageMenuItemNew             func() *Widget
	ImageMenuItemNewWithLabel    func(label string) *Widget
	ImageMenuItemNewWithMnemonic func(label string) *Widget
	ImageMenuItemNewFromStock    func(stockId string, accelGroup *AccelGroup) *Widget

	imageMenuItemGetAlwaysShowImage func(i *ImageMenuItem) T.Gboolean
	imageMenuItemGetImage           func(i *ImageMenuItem) *Widget
	imageMenuItemGetUseStock        func(i *ImageMenuItem) T.Gboolean
	imageMenuItemSetAccelGroup      func(i *ImageMenuItem, accelGroup *AccelGroup)
	imageMenuItemSetAlwaysShowImage func(i *ImageMenuItem, alwaysShow T.Gboolean)
	imageMenuItemSetImage           func(i *ImageMenuItem, image *Widget)
	imageMenuItemSetUseStock        func(i *ImageMenuItem, useStock T.Gboolean)
)

func (i *ImageMenuItem) GetAlwaysShowImage() T.Gboolean { return imageMenuItemGetAlwaysShowImage(i) }
func (i *ImageMenuItem) GetImage() *Widget              { return imageMenuItemGetImage(i) }
func (i *ImageMenuItem) GetUseStock() T.Gboolean        { return imageMenuItemGetUseStock(i) }
func (i *ImageMenuItem) SetAccelGroup(accelGroup *AccelGroup) {
	imageMenuItemSetAccelGroup(i, accelGroup)
}
func (i *ImageMenuItem) SetAlwaysShowImage(alwaysShow T.Gboolean) {
	imageMenuItemSetAlwaysShowImage(i, alwaysShow)
}
func (i *ImageMenuItem) SetImage(image *Widget)          { imageMenuItemSetImage(i, image) }
func (i *ImageMenuItem) SetUseStock(useStock T.Gboolean) { imageMenuItemSetUseStock(i, useStock) }

type IMContext simpleObject

var (
	ImContextGetType func() T.GType

	imContextDeleteSurrounding func(i *IMContext, offset int, nChars int) T.Gboolean
	imContextFilterKeypress    func(i *IMContext, event *T.GdkEventKey) T.Gboolean
	imContextFocusIn           func(i *IMContext)
	imContextFocusOut          func(i *IMContext)
	imContextGetPreeditString  func(i *IMContext, str **T.Gchar, attrs **T.PangoAttrList, cursorPos *int)
	imContextGetSurrounding    func(i *IMContext, text **T.Gchar, cursorIndex *int) T.Gboolean
	imContextReset             func(i *IMContext)
	imContextSetClientWindow   func(i *IMContext, window *T.GdkWindow)
	imContextSetCursorLocation func(i *IMContext, area *T.GdkRectangle)
	imContextSetSurrounding    func(i *IMContext, text string, len int, cursorIndex int)
	imContextSetUsePreedit     func(i *IMContext, usePreedit T.Gboolean)
)

func (i *IMContext) DeleteSurrounding(offset, nChars int) T.Gboolean {
	return imContextDeleteSurrounding(i, offset, nChars)
}
func (i *IMContext) FilterKeypress(event *T.GdkEventKey) T.Gboolean {
	return imContextFilterKeypress(i, event)
}
func (i *IMContext) FocusIn()  { imContextFocusIn(i) }
func (i *IMContext) FocusOut() { imContextFocusOut(i) }
func (i *IMContext) GetPreeditString(str **T.Gchar, attrs **T.PangoAttrList, cursorPos *int) {
	imContextGetPreeditString(i, str, attrs, cursorPos)
}
func (i *IMContext) GetSurrounding(text **T.Gchar, cursorIndex *int) T.Gboolean {
	return imContextGetSurrounding(i, text, cursorIndex)
}
func (i *IMContext) Reset()                                 { imContextReset(i) }
func (i *IMContext) SetClientWindow(window *T.GdkWindow)    { imContextSetClientWindow(i, window) }
func (i *IMContext) SetCursorLocation(area *T.GdkRectangle) { imContextSetCursorLocation(i, area) }
func (i *IMContext) SetSurrounding(text string, leng int, cursorIndex int) {
	imContextSetSurrounding(i, text, leng, cursorIndex)
}
func (i *IMContext) SetUsePreedit(usePreedit T.Gboolean) { imContextSetUsePreedit(i, usePreedit) }

type IMContextSimple struct {
	Object            IMContext
	Tables            *T.GSList
	ComposeBuffer     [7 + 1]uint //TODO(t):Symbolic
	TentativeMatch    T.Gunichar
	TentativeMatchLen int
	Bits              uint
	// InHexSequence : 1
	// ModifiersDropped : 1
}

var (
	ImContextSimpleGetType func() T.GType
	ImContextSimpleNew     func() *IMContext

	imContextSimpleAddTable func(i *IMContextSimple, data *uint16, maxSeqLen, nSeqs int)
)

func (i *IMContextSimple) AddTable(data *uint16, maxSeqLen, nSeqs int) {
	imContextSimpleAddTable(i, data, maxSeqLen, nSeqs)
}

type IMMulticontext struct {
	Object    IMContext
	Slave     *IMContext
	_         *struct{}
	ContextId *T.Gchar
}

var (
	ImMulticontextGetType func() T.GType
	ImMulticontextNew     func() *IMContext

	imMulticontextAppendMenuitems func(i *IMMulticontext, menushell *MenuShell)
	imMulticontextGetContextId    func(i *IMMulticontext) string
	imMulticontextSetContextId    func(i *IMMulticontext, contextId string)
)

func (i *IMMulticontext) AppendMenuitems(menushell *MenuShell) {
	imMulticontextAppendMenuitems(i, menushell)
}
func (i *IMMulticontext) GetContextId() string          { return imMulticontextGetContextId(i) }
func (i *IMMulticontext) SetContextId(contextId string) { imMulticontextSetContextId(i, contextId) }

type InfoBar struct {
	Parent HBox
	_      *struct{}
}

var (
	InfoBarGetType        func() T.GType
	InfoBarNew            func() *Widget
	InfoBarNewWithButtons func(firstButtonText string, v ...VArg) *Widget

	infoBarAddActionWidget      func(i *InfoBar, child *Widget, responseId int)
	infoBarAddButton            func(i *InfoBar, buttonText string, responseId int) *Widget
	infoBarAddButtons           func(i *InfoBar, firstButtonText string, v ...VArg)
	infoBarGetActionArea        func(i *InfoBar) *Widget
	infoBarGetContentArea       func(i *InfoBar) *Widget
	infoBarGetMessageType       func(i *InfoBar) T.GtkMessageType
	infoBarResponse             func(i *InfoBar, responseId int)
	infoBarSetDefaultResponse   func(i *InfoBar, responseId int)
	infoBarSetMessageType       func(i *InfoBar, messageType T.GtkMessageType)
	infoBarSetResponseSensitive func(i *InfoBar, responseId int, setting T.Gboolean)
)

func (i *InfoBar) AddActionWidget(child *Widget, responseId int) {
	infoBarAddActionWidget(i, child, responseId)
}
func (i *InfoBar) AddButton(buttonText string, responseId int) *Widget {
	return infoBarAddButton(i, buttonText, responseId)
}
func (i *InfoBar) AddButtons(firstButtonText string, v ...VArg) {
	infoBarAddButtons(i, firstButtonText, v)
}
func (i *InfoBar) GetActionArea() *Widget                      { return infoBarGetActionArea(i) }
func (i *InfoBar) GetContentArea() *Widget                     { return infoBarGetContentArea(i) }
func (i *InfoBar) GetMessageType() T.GtkMessageType            { return infoBarGetMessageType(i) }
func (i *InfoBar) Response(responseId int)                     { infoBarResponse(i, responseId) }
func (i *InfoBar) SetDefaultResponse(responseId int)           { infoBarSetDefaultResponse(i, responseId) }
func (i *InfoBar) SetMessageType(messageType T.GtkMessageType) { infoBarSetMessageType(i, messageType) }
func (i *InfoBar) SetResponseSensitive(responseId int, setting T.Gboolean) {
	infoBarSetResponseSensitive(i, responseId, setting)
}

type Invisible struct {
	Widget          Widget
	HasUserRefCount T.Gboolean
	Screen          *T.GdkScreen
}

var (
	InvisibleGetType      func() T.GType
	InvisibleNew          func() *Widget
	InvisibleNewForScreen func(screen *T.GdkScreen) *Widget

	invisibleGetScreen func(i *Invisible) *T.GdkScreen
	invisibleSetScreen func(i *Invisible, screen *T.GdkScreen)
)

func (i *Invisible) GetScreen() *T.GdkScreen       { return invisibleGetScreen(i) }
func (i *Invisible) SetScreen(screen *T.GdkScreen) { invisibleSetScreen(i, screen) }

type Item struct {
	Bin Bin
}

var (
	ItemGetType func() T.GType

	itemDeselect func(i *Item)
	itemSelect   func(i *Item)
	itemToggle   func(i *Item)
)

func (i *Item) Deselect() { itemDeselect(i) }
func (i *Item) Select()   { itemSelect(i) }
func (i *Item) Toggle()   { itemToggle(i) }

type (
	ItemFactory struct {
		Object           T.GtkObject
		Path             *T.Gchar
		Accel_group      *AccelGroup
		Widget           *Widget
		Items            *T.GSList
		Translate_func   T.GtkTranslateFunc
		Translate_data   T.Gpointer
		Translate_notify T.GDestroyNotify
	}

	ItemFactoryEntry struct {
		Path            *T.Gchar
		Accelerator     *T.Gchar
		Callback        ItemFactoryCallback
		Callback_action uint
		Item_type       *T.Gchar
		Extra_data      T.Gconstpointer
	}

	ItemFactoryCallback func()
)

var (
	ItemFactoryGetType func() T.GType
	ItemFactoryNew     func(containerType T.GType, path string, accelGroup *AccelGroup) *ItemFactory

	ItemFactoriesPathDelete        func(ifactoryPath string, path string)
	ItemFactoryAddForeign          func(accelWidget *Widget, fullPath string, accelGroup *AccelGroup, keyval uint, modifiers T.GdkModifierType)
	ItemFactoryCreateMenuEntries   func(nEntries uint, entries *MenuEntry)
	ItemFactoryFromPath            func(path string) *ItemFactory
	ItemFactoryFromWidget          func(widget *Widget) *ItemFactory
	ItemFactoryPathFromWidget      func(widget *Widget) string
	ItemFactoryPopupDataFromWidget func(widget *Widget) T.Gpointer

	itemFactoryConstruct         func(i *ItemFactory, containerType T.GType, path string, accelGroup *AccelGroup)
	itemFactoryCreateItem        func(i *ItemFactory, entry *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint)
	itemFactoryCreateItems       func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer)
	itemFactoryCreateItemsAc     func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint)
	itemFactoryDeleteEntries     func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry)
	itemFactoryDeleteEntry       func(i *ItemFactory, entry *ItemFactoryEntry)
	itemFactoryDeleteItem        func(i *ItemFactory, path string)
	itemFactoryGetItem           func(i *ItemFactory, path string) *Widget
	itemFactoryGetItemByAction   func(i *ItemFactory, action uint) *Widget
	itemFactoryGetWidget         func(i *ItemFactory, path string) *Widget
	itemFactoryGetWidgetByAction func(i *ItemFactory, action uint) *Widget
	itemFactoryPopup             func(i *ItemFactory, x uint, y uint, mouseButton uint, time T.GUint32)
	itemFactoryPopupData         func(i *ItemFactory) T.Gpointer
	itemFactoryPopupWithData     func(i *ItemFactory, popupData T.Gpointer, destroy T.GDestroyNotify, x uint, y uint, mouseButton uint, time T.GUint32)
	itemFactorySetTranslateFunc  func(i *ItemFactory, f T.GtkTranslateFunc, data T.Gpointer, notify T.GDestroyNotify)
)

func (i *ItemFactory) Construct(containerType T.GType, path string, accelGroup *AccelGroup) {
	itemFactoryConstruct(i, containerType, path, accelGroup)
}
func (i *ItemFactory) CreateItem(entry *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint) {
	itemFactoryCreateItem(i, entry, callbackData, callbackType)
}
func (i *ItemFactory) CreateItems(nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer) {
	itemFactoryCreateItems(i, nEntries, entries, callbackData)
}
func (i *ItemFactory) CreateItemsAc(nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint) {
	itemFactoryCreateItemsAc(i, nEntries, entries, callbackData, callbackType)
}
func (i *ItemFactory) DeleteEntries(nEntries uint, entries *ItemFactoryEntry) {
	itemFactoryDeleteEntries(i, nEntries, entries)
}
func (i *ItemFactory) DeleteEntry(entry *ItemFactoryEntry) { itemFactoryDeleteEntry(i, entry) }
func (i *ItemFactory) DeleteItem(path string)              { itemFactoryDeleteItem(i, path) }
func (i *ItemFactory) GetItem(path string) *Widget         { return itemFactoryGetItem(i, path) }
func (i *ItemFactory) GetItemByAction(action uint) *Widget {
	return itemFactoryGetItemByAction(i, action)
}
func (i *ItemFactory) GetWidget(path string) *Widget { return itemFactoryGetWidget(i, path) }
func (i *ItemFactory) GetWidgetByAction(action uint) *Widget {
	return itemFactoryGetWidgetByAction(i, action)
}
func (i *ItemFactory) Popup(x, y, mouseButton uint, time T.GUint32) {
	itemFactoryPopup(i, x, y, mouseButton, time)
}
func (i *ItemFactory) PopupData() T.Gpointer { return itemFactoryPopupData(i) }
func (i *ItemFactory) PopupWithData(popupData T.Gpointer, destroy T.GDestroyNotify, x, y, mouseButton uint, time T.GUint32) {
	itemFactoryPopupWithData(i, popupData, destroy, x, y, mouseButton, time)
}
func (i *ItemFactory) SetTranslateFunc(f T.GtkTranslateFunc, data T.Gpointer, notify T.GDestroyNotify) {
	itemFactorySetTranslateFunc(i, f, data, notify)
}
