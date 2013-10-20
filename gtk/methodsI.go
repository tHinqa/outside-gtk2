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

	IconFactoryAdd           func(i *IconFactory, stockId string, iconSet *IconSet)
	IconFactoryAddDefault    func(i *IconFactory)
	IconFactoryLookup        func(i *IconFactory, stockId string) *IconSet
	IconFactoryRemoveDefault func(i *IconFactory)
)

func (i *IconFactory) Add(stockId string, iconSet *IconSet) {
	IconFactoryAdd(i, stockId, iconSet)
}

func (i *IconFactory) Lookup(stockId string) *IconSet {
	return IconFactoryLookup(i, stockId)
}

func (i *IconFactory) AddDefault() { IconFactoryAddDefault(i) }

func (i *IconFactory) RemoveDefault() {
	IconFactoryRemoveDefault(i)
}

type IconInfo struct{}

var (
	IconInfoGetType      func() T.GType
	IconInfoNewForPixbuf func(iconTheme *IconTheme, pixbuf *T.GdkPixbuf) *IconInfo

	IconInfoCopy              func(i *IconInfo) *IconInfo
	IconInfoFree              func(i *IconInfo)
	IconInfoGetAttachPoints   func(i *IconInfo, points **T.GdkPoint, nPoints *int) T.Gboolean
	IconInfoGetBaseSize       func(i *IconInfo) int
	IconInfoGetBuiltinPixbuf  func(i *IconInfo) *T.GdkPixbuf
	IconInfoGetDisplayName    func(i *IconInfo) string
	IconInfoGetEmbeddedRect   func(i *IconInfo, rectangle *T.GdkRectangle) T.Gboolean
	IconInfoGetFilename       func(i *IconInfo) string
	IconInfoLoadIcon          func(i *IconInfo, error **T.GError) *T.GdkPixbuf
	IconInfoSetRawCoordinates func(i *IconInfo, rawCoordinates T.Gboolean)
)

func (i *IconInfo) Copy() *IconInfo { return IconInfoCopy(i) }

func (i *IconInfo) Free() { IconInfoFree(i) }

func (i *IconInfo) GetBaseSize() int {
	return IconInfoGetBaseSize(i)
}

func (i *IconInfo) GetFilename() string {
	return IconInfoGetFilename(i)
}

func (i *IconInfo) GetBuiltinPixbuf() *T.GdkPixbuf {
	return IconInfoGetBuiltinPixbuf(i)
}

func (i *IconInfo) LoadIcon(err **T.GError) *T.GdkPixbuf {
	return IconInfoLoadIcon(i, err)
}

func (i *IconInfo) SetRawCoordinates(rawCoordinates T.Gboolean) {
	IconInfoSetRawCoordinates(i, rawCoordinates)
}

func (i *IconInfo) GetEmbeddedRect(
	rectangle *T.GdkRectangle) T.Gboolean {
	return IconInfoGetEmbeddedRect(i, rectangle)
}

func (i *IconInfo) GetAttachPoints(
	points **T.GdkPoint, nPoints *int) T.Gboolean {
	return IconInfoGetAttachPoints(i, points, nPoints)
}

func (i *IconInfo) GetDisplayName() string {
	return IconInfoGetDisplayName(i)
}

type IconSet struct{}

var (
	IconSetGetType       func() T.GType
	IconSetNew           func() *IconSet
	IconSetNewFromPixbuf func(pixbuf *T.GdkPixbuf) *IconSet

	IconSetAddSource  func(i *IconSet, source *IconSource)
	IconSetCopy       func(i *IconSet) *IconSet
	IconSetGetSizes   func(i *IconSet, sizes **IconSize, nSizes *int)
	IconSetRef        func(i *IconSet) *IconSet
	IconSetRenderIcon func(i *IconSet, style *T.GtkStyle, direction T.GtkTextDirection, state T.GtkStateType, size IconSize, widget *T.GtkWidget, detail string) *T.GdkPixbuf
	IconSetUnref      func(i *IconSet)
)

func (i *IconSet) Ref() *IconSet { return IconSetRef(i) }

func (i *IconSet) Unref() { IconSetUnref(i) }

func (i *IconSet) Copy() *IconSet { return IconSetCopy(i) }

func (i *IconSet) RenderIcon(style *T.GtkStyle,
	direction T.GtkTextDirection, state T.GtkStateType,
	size IconSize, widget *T.GtkWidget, detail string) *T.GdkPixbuf {
	return IconSetRenderIcon(
		i, style, direction, state, size, widget, detail)
}

func (i *IconSet) AddSource(source *IconSource) {
	IconSetAddSource(i, source)
}

func (i *IconSet) GetSizes(sizes **IconSize, nSizes *int) {
	IconSetGetSizes(i, sizes, nSizes)
}

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

	IconSizeGetName func(i IconSize) string
	IconSizeLookup  func(i IconSize, width *int, height *int) T.Gboolean
)

func (i IconSize) Lookup(width *int, height *int) T.Gboolean {
	return IconSizeLookup(i, width, height)
}

func (i IconSize) GetName() string {
	return IconSizeGetName(i)
}

type IconSource struct{}

var (
	IconSourceGetType func() T.GType
	IconSourceNew     func() *IconSource

	IconSourceCopy                   func(i *IconSource) *IconSource
	IconSourceFree                   func(i *IconSource)
	IconSourceGetDirection           func(i *IconSource) T.GtkTextDirection
	IconSourceGetDirectionWildcarded func(i *IconSource) T.Gboolean
	IconSourceGetFilename            func(i *IconSource) string
	IconSourceGetIconName            func(i *IconSource) string
	IconSourceGetPixbuf              func(i *IconSource) *T.GdkPixbuf
	IconSourceGetSize                func(i *IconSource) IconSize
	IconSourceGetSizeWildcarded      func(i *IconSource) T.Gboolean
	IconSourceGetState               func(i *IconSource) T.GtkStateType
	IconSourceGetStateWildcarded     func(i *IconSource) T.Gboolean
	IconSourceSetDirection           func(i *IconSource, direction T.GtkTextDirection)
	IconSourceSetDirectionWildcarded func(i *IconSource, setting T.Gboolean)
	IconSourceSetFilename            func(i *IconSource, filename string)
	IconSourceSetIconName            func(i *IconSource, iconName string)
	IconSourceSetPixbuf              func(i *IconSource, pixbuf *T.GdkPixbuf)
	IconSourceSetSize                func(i *IconSource, size IconSize)
	IconSourceSetSizeWildcarded      func(i *IconSource, setting T.Gboolean)
	IconSourceSetState               func(i *IconSource, state T.GtkStateType)
	IconSourceSetStateWildcarded     func(i *IconSource, setting T.Gboolean)
)

func (i *IconSource) Copy() *IconSource {
	return IconSourceCopy(i)
}

func (i *IconSource) Free() { IconSourceFree(i) }

func (i *IconSource) SetFilename(filename string) {
	IconSourceSetFilename(i, filename)
}

func (i *IconSource) SetIconName(iconName string) {
	IconSourceSetIconName(i, iconName)
}

func (i *IconSource) SetPixbuf(pixbuf *T.GdkPixbuf) {
	IconSourceSetPixbuf(i, pixbuf)
}

func (i *IconSource) GetFilename() string {
	return IconSourceGetFilename(i)
}

func (i *IconSource) GetIconName() string {
	return IconSourceGetIconName(i)
}

func (i *IconSource) GetPixbuf() *T.GdkPixbuf {
	return IconSourceGetPixbuf(i)
}

func (i *IconSource) SetDirectionWildcarded(setting T.Gboolean) {
	IconSourceSetDirectionWildcarded(i, setting)
}

func (i *IconSource) SetStateWildcarded(setting T.Gboolean) {
	IconSourceSetStateWildcarded(i, setting)
}

func (i *IconSource) SetSizeWildcarded(setting T.Gboolean) {
	IconSourceSetSizeWildcarded(i, setting)
}

func (i *IconSource) GetSizeWildcarded() T.Gboolean {
	return IconSourceGetSizeWildcarded(i)
}

func (i *IconSource) GetStateWildcarded() T.Gboolean {
	return IconSourceGetStateWildcarded(i)
}

func (i *IconSource) GetDirectionWildcarded() T.Gboolean {
	return IconSourceGetDirectionWildcarded(i)
}

func (i *IconSource) SetDirection(direction T.GtkTextDirection) {
	IconSourceSetDirection(i, direction)
}

func (i *IconSource) SetState(state T.GtkStateType) {
	IconSourceSetState(i, state)
}

func (i *IconSource) SetSize(size IconSize) {
	IconSourceSetSize(i, size)
}

func (i *IconSource) GetDirection() T.GtkTextDirection {
	return IconSourceGetDirection(i)
}

func (i *IconSource) GetState() T.GtkStateType {
	return IconSourceGetState(i)
}

func (i *IconSource) GetSize() IconSize {
	return IconSourceGetSize(i)
}

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

	IconThemeAppendSearchPath   func(i *IconTheme, path string)
	IconThemeChooseIcon         func(i *IconTheme, iconNames **T.Gchar, size int, flags IconLookupFlags) *IconInfo
	IconThemeGetExampleIconName func(i *IconTheme) string
	IconThemeGetIconSizes       func(i *IconTheme, iconName string) *int
	IconThemeGetSearchPath      func(i *IconTheme, path ***T.Gchar, nElements *int)
	IconThemeHasIcon            func(i *IconTheme, iconName string) T.Gboolean
	IconThemeListContexts       func(i *IconTheme) *T.GList
	IconThemeListIcons          func(i *IconTheme, context string) *T.GList
	IconThemeLoadIcon           func(i *IconTheme, iconName string, size int, flags IconLookupFlags, error **T.GError) *T.GdkPixbuf
	IconThemeLookupByGicon      func(i *IconTheme, icon *T.GIcon, size int, flags IconLookupFlags) *IconInfo
	IconThemeLookupIcon         func(i *IconTheme, iconName string, size int, flags IconLookupFlags) *IconInfo
	IconThemePrependSearchPath  func(i *IconTheme, path string)
	IconThemeRescanIfNeeded     func(i *IconTheme) T.Gboolean
	IconThemeSetCustomTheme     func(i *IconTheme, themeName string)
	IconThemeSetScreen          func(i *IconTheme, screen *T.GdkScreen)
	IconThemeSetSearchPath      func(i *IconTheme, path **T.Gchar, nElements int)
)

func (i *IconTheme) SetScreen(screen *T.GdkScreen) {
	IconThemeSetScreen(i, screen)
}

func (i *IconTheme) SetSearchPath(path **T.Gchar, nElements int) {
	IconThemeSetSearchPath(i, path, nElements)
}

func (i *IconTheme) GetSearchPath(path ***T.Gchar, nElements *int) {
	IconThemeGetSearchPath(i, path, nElements)
}

func (i *IconTheme) AppendSearchPath(path string) {
	IconThemeAppendSearchPath(i, path)
}

func (i *IconTheme) PrependSearchPath(path string) {
	IconThemePrependSearchPath(i, path)
}

func (i *IconTheme) SetCustomTheme(themeName string) {
	IconThemeSetCustomTheme(i, themeName)
}

func (i *IconTheme) HasIcon(iconName string) T.Gboolean {
	return IconThemeHasIcon(i, iconName)
}

func (i *IconTheme) GetIconSizes(iconName string) *int {
	return IconThemeGetIconSizes(i, iconName)
}

func (i *IconTheme) LookupIcon(iconName string, size int,
	flags IconLookupFlags) *IconInfo {
	return IconThemeLookupIcon(i, iconName, size, flags)
}

func (i *IconTheme) ChooseIcon(iconNames **T.Gchar, size int,
	flags IconLookupFlags) *IconInfo {
	return IconThemeChooseIcon(i, iconNames, size, flags)
}

func (i *IconTheme) LoadIcon(iconName string, size int,
	flags IconLookupFlags, err **T.GError) *T.GdkPixbuf {
	return IconThemeLoadIcon(i, iconName, size, flags, err)
}

func (i *IconTheme) LookupByGicon(icon *T.GIcon, size int,
	flags IconLookupFlags) *IconInfo {
	return IconThemeLookupByGicon(i, icon, size, flags)
}

func (i *IconTheme) ListIcons(context string) *T.GList {
	return IconThemeListIcons(i, context)
}

func (i *IconTheme) ListContexts() *T.GList {
	return IconThemeListContexts(i)
}

func (i *IconTheme) GetExampleIconName() string {
	return IconThemeGetExampleIconName(i)
}

func (i *IconTheme) RescanIfNeeded() T.Gboolean {
	return IconThemeRescanIfNeeded(i)
}

type (
	IconView struct {
		Parent T.GtkContainer
		_      *struct{}
	}

	IconViewForeachFunc func(icon_view *IconView,
		path *T.GtkTreePath, data T.Gpointer)
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
	IconViewNew          func() *T.GtkWidget
	IconViewNewWithModel func(model *T.GtkTreeModel) *T.GtkWidget

	IconViewDropPositionGetType func() T.GType

	IconViewConvertWidgetToBinWindowCoords func(i *IconView, wx int, wy int, bx *int, by *int)
	IconViewCreateDragIcon                 func(i *IconView, path *T.GtkTreePath) *T.GdkPixmap
	IconViewEnableModelDragDest            func(i *IconView, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction)
	IconViewEnableModelDragSource          func(i *IconView, startButtonMask T.GdkModifierType, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction)
	IconViewGetColumns                     func(i *IconView) int
	IconViewGetColumnSpacing               func(i *IconView) int
	IconViewGetCursor                      func(i *IconView, path **T.GtkTreePath, cell **CellRenderer) T.Gboolean
	IconViewGetDestItemAtPos               func(i *IconView, dragX, dragY int, path **T.GtkTreePath, pos *IconViewDropPosition) T.Gboolean
	IconViewGetDragDestItem                func(i *IconView, path **T.GtkTreePath, pos *IconViewDropPosition)
	IconViewGetItemAtPos                   func(i *IconView, x, y int, path **T.GtkTreePath, cell **CellRenderer) T.Gboolean
	IconViewGetItemColumn                  func(i *IconView, path *T.GtkTreePath) int
	IconViewGetItemOrientation             func(i *IconView) T.GtkOrientation
	IconViewGetItemPadding                 func(i *IconView) int
	IconViewGetItemRow                     func(i *IconView, path *T.GtkTreePath) int
	IconViewGetItemWidth                   func(i *IconView) int
	IconViewGetMargin                      func(i *IconView) int
	IconViewGetMarkupColumn                func(i *IconView) int
	IconViewGetModel                       func(i *IconView) *T.GtkTreeModel
	IconViewGetOrientation                 func(i *IconView) T.GtkOrientation
	IconViewGetPathAtPos                   func(i *IconView, x, y int) *T.GtkTreePath
	IconViewGetPixbufColumn                func(i *IconView) int
	IconViewGetReorderable                 func(i *IconView) T.Gboolean
	IconViewGetRowSpacing                  func(i *IconView) int
	IconViewGetSelectedItems               func(i *IconView) *T.GList
	IconViewGetSelectionMode               func(i *IconView) T.GtkSelectionMode
	IconViewGetSpacing                     func(i *IconView) int
	IconViewGetTextColumn                  func(i *IconView) int
	IconViewGetTooltipColumn               func(i *IconView) int
	IconViewGetTooltipContext              func(i *IconView, x *int, y *int, keyboardTip T.Gboolean, model **T.GtkTreeModel, path **T.GtkTreePath, iter *T.GtkTreeIter) T.Gboolean
	IconViewGetVisibleRange                func(i *IconView, startPath **T.GtkTreePath, endPath **T.GtkTreePath) T.Gboolean
	IconViewItemActivated                  func(i *IconView, path *T.GtkTreePath)
	IconViewPathIsSelected                 func(i *IconView, path *T.GtkTreePath) T.Gboolean
	IconViewScrollToPath                   func(i *IconView, path *T.GtkTreePath, useAlign T.Gboolean, rowAlign, colAlign float32)
	IconViewSelectAll                      func(i *IconView)
	IconViewSelectedForeach                func(i *IconView, f IconViewForeachFunc, data T.Gpointer)
	IconViewSelectPath                     func(i *IconView, path *T.GtkTreePath)
	IconViewSetColumns                     func(i *IconView, columns int)
	IconViewSetColumnSpacing               func(i *IconView, columnSpacing int)
	IconViewSetCursor                      func(i *IconView, path *T.GtkTreePath, cell *CellRenderer, startEditing T.Gboolean)
	IconViewSetDragDestItem                func(i *IconView, path *T.GtkTreePath, pos IconViewDropPosition)
	IconViewSetItemOrientation             func(i *IconView, orientation T.GtkOrientation)
	IconViewSetItemPadding                 func(i *IconView, itemPadding int)
	IconViewSetItemWidth                   func(i *IconView, itemWidth int)
	IconViewSetMargin                      func(i *IconView, margin int)
	IconViewSetMarkupColumn                func(i *IconView, column int)
	IconViewSetModel                       func(i *IconView, model *T.GtkTreeModel)
	IconViewSetOrientation                 func(i *IconView, orientation T.GtkOrientation)
	IconViewSetPixbufColumn                func(i *IconView, column int)
	IconViewSetReorderable                 func(i *IconView, reorderable T.Gboolean)
	IconViewSetRowSpacing                  func(i *IconView, rowSpacing int)
	IconViewSetSelectionMode               func(i *IconView, mode T.GtkSelectionMode)
	IconViewSetSpacing                     func(i *IconView, spacing int)
	IconViewSetTextColumn                  func(i *IconView, column int)
	IconViewSetTooltipCell                 func(i *IconView, tooltip *T.GtkTooltip, path *T.GtkTreePath, cell *CellRenderer)
	IconViewSetTooltipColumn               func(i *IconView, column int)
	IconViewSetTooltipItem                 func(i *IconView, tooltip *T.GtkTooltip, path *T.GtkTreePath)
	IconViewUnselectAll                    func(i *IconView)
	IconViewUnselectPath                   func(i *IconView, path *T.GtkTreePath)
	IconViewUnsetModelDragDest             func(i *IconView)
	IconViewUnsetModelDragSource           func(i *IconView)
)

func (i *IconView) SetModel(model *T.GtkTreeModel) {
	IconViewSetModel(i, model)
}

func (i *IconView) GetModel() *T.GtkTreeModel {
	return IconViewGetModel(i)
}

func (i *IconView) SetTextColumn(column int) {
	IconViewSetTextColumn(i, column)
}

func (i *IconView) GetTextColumn() int {
	return IconViewGetTextColumn(i)
}

func (i *IconView) SetMarkupColumn(column int) {
	IconViewSetMarkupColumn(i, column)
}

func (i *IconView) GetMarkupColumn() int {
	return IconViewGetMarkupColumn(i)
}

func (i *IconView) SetPixbufColumn(column int) {
	IconViewSetPixbufColumn(i, column)
}

func (i *IconView) GetPixbufColumn() int {
	return IconViewGetPixbufColumn(i)
}

func (i *IconView) SetOrientation(orientation T.GtkOrientation) {
	IconViewSetOrientation(i, orientation)
}

func (i *IconView) GetOrientation() T.GtkOrientation {
	return IconViewGetOrientation(i)
}

func (i *IconView) SetItemOrientation(orientation T.GtkOrientation) {
	IconViewSetItemOrientation(i, orientation)
}

func (i *IconView) GetItemOrientation() T.GtkOrientation {
	return IconViewGetItemOrientation(i)
}

func (i *IconView) SetColumns(columns int) {
	IconViewSetColumns(i, columns)
}

func (i *IconView) GetColumns() int {
	return IconViewGetColumns(i)
}

func (i *IconView) SetItemWidth(itemWidth int) {
	IconViewSetItemWidth(i, itemWidth)
}

func (i *IconView) GetItemWidth() int {
	return IconViewGetItemWidth(i)
}

func (i *IconView) SetSpacing(spacing int) {
	IconViewSetSpacing(i, spacing)
}

func (i *IconView) GetSpacing() int {
	return IconViewGetSpacing(i)
}

func (i *IconView) SetRowSpacing(rowSpacing int) {
	IconViewSetRowSpacing(i, rowSpacing)
}

func (i *IconView) GetRowSpacing() int {
	return IconViewGetRowSpacing(i)
}

func (i *IconView) SetColumnSpacing(columnSpacing int) {
	IconViewSetColumnSpacing(i, columnSpacing)
}

func (i *IconView) GetColumnSpacing() int {
	return IconViewGetColumnSpacing(i)
}

func (i *IconView) SetMargin(margin int) {
	IconViewSetMargin(i, margin)
}

func (i *IconView) GetMargin() int {
	return IconViewGetMargin(i)
}

func (i *IconView) SetItemPadding(itemPadding int) {
	IconViewSetItemPadding(i, itemPadding)
}

func (i *IconView) GetItemPadding() int {
	return IconViewGetItemPadding(i)
}

func (i *IconView) GetPathAtPos(x, y int) *T.GtkTreePath {
	return IconViewGetPathAtPos(i, x, y)
}

func (i *IconView) GetItemAtPos(x, y int, path **T.GtkTreePath, cell **CellRenderer) T.Gboolean {
	return IconViewGetItemAtPos(i, x, y, path, cell)
}

func (i *IconView) GetVisibleRange(startPath **T.GtkTreePath, endPath **T.GtkTreePath) T.Gboolean {
	return IconViewGetVisibleRange(i, startPath, endPath)
}

func (i *IconView) SelectedForeach(f IconViewForeachFunc, data T.Gpointer) {
	IconViewSelectedForeach(i, f, data)
}

func (i *IconView) SetSelectionMode(mode T.GtkSelectionMode) {
	IconViewSetSelectionMode(i, mode)
}

func (i *IconView) GetSelectionMode() T.GtkSelectionMode {
	return IconViewGetSelectionMode(i)
}

func (i *IconView) SelectPath(path *T.GtkTreePath) {
	IconViewSelectPath(i, path)
}

func (i *IconView) UnselectPath(path *T.GtkTreePath) {
	IconViewUnselectPath(i, path)
}

func (i *IconView) PathIsSelected(path *T.GtkTreePath) T.Gboolean {
	return IconViewPathIsSelected(i, path)
}

func (i *IconView) GetItemRow(path *T.GtkTreePath) int {
	return IconViewGetItemRow(i, path)
}

func (i *IconView) GetItemColumn(path *T.GtkTreePath) int {
	return IconViewGetItemColumn(i, path)
}

func (i *IconView) GetSelectedItems() *T.GList {
	return IconViewGetSelectedItems(i)
}

func (i *IconView) SelectAll() { IconViewSelectAll(i) }

func (i *IconView) UnselectAll() { IconViewUnselectAll(i) }

func (i *IconView) ItemActivated(path *T.GtkTreePath) {
	IconViewItemActivated(i, path)
}

func (i *IconView) SetCursor(path *T.GtkTreePath, cell *CellRenderer, startEditing T.Gboolean) {
	IconViewSetCursor(i, path, cell, startEditing)
}

func (i *IconView) GetCursor(path **T.GtkTreePath, cell **CellRenderer) T.Gboolean {
	return IconViewGetCursor(i, path, cell)
}

func (i *IconView) ScrollToPath(path *T.GtkTreePath, useAlign T.Gboolean, rowAlign, colAlign float32) {
	IconViewScrollToPath(i, path, useAlign, rowAlign, colAlign)
}

func (i *IconView) EnableModelDragSource(startButtonMask T.GdkModifierType, targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction) {
	IconViewEnableModelDragSource(i, startButtonMask, targets, nTargets, actions)
}

func (i *IconView) EnableModelDragDest(targets *T.GtkTargetEntry, nTargets int, actions T.GdkDragAction) {
	IconViewEnableModelDragDest(i, targets, nTargets, actions)
}

func (i *IconView) UnsetModelDragSource() {
	IconViewUnsetModelDragSource(i)
}

func (i *IconView) UnsetModelDragDest() {
	IconViewUnsetModelDragDest(i)
}

func (i *IconView) SetReorderable(reorderable T.Gboolean) {
	IconViewSetReorderable(i, reorderable)
}

func (i *IconView) GetReorderable() T.Gboolean {
	return IconViewGetReorderable(i)
}

func (i *IconView) SetDragDestItem(path *T.GtkTreePath, pos IconViewDropPosition) {
	IconViewSetDragDestItem(i, path, pos)
}

func (i *IconView) GetDragDestItem(path **T.GtkTreePath, pos *IconViewDropPosition) {
	IconViewGetDragDestItem(i, path, pos)
}

func (i *IconView) GetDestItemAtPos(dragX, dragY int, path **T.GtkTreePath, pos *IconViewDropPosition) T.Gboolean {
	return IconViewGetDestItemAtPos(i, dragX, dragY, path, pos)
}

func (i *IconView) CreateDragIcon(path *T.GtkTreePath) *T.GdkPixmap {
	return IconViewCreateDragIcon(i, path)
}

func (i *IconView) ConvertWidgetToBinWindowCoords(wx, wy int, bx, by *int) {
	IconViewConvertWidgetToBinWindowCoords(i, wx, wy, bx, by)
}

func (i *IconView) SetTooltipItem(tooltip *T.GtkTooltip, path *T.GtkTreePath) {
	IconViewSetTooltipItem(i, tooltip, path)
}

func (i *IconView) SetTooltipCell(tooltip *T.GtkTooltip, path *T.GtkTreePath, cell *CellRenderer) {
	IconViewSetTooltipCell(i, tooltip, path, cell)
}

func (i *IconView) GetTooltipContext(x, y *int, keyboardTip T.Gboolean, model **T.GtkTreeModel, path **T.GtkTreePath, iter *T.GtkTreeIter) T.Gboolean {
	return IconViewGetTooltipContext(i, x, y, keyboardTip, model, path, iter)
}

func (i *IconView) SetTooltipColumn(column int) {
	IconViewSetTooltipColumn(i, column)
}

func (i *IconView) GetTooltipColumn() int {
	return IconViewGetTooltipColumn(i)
}

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
	ImageNew              func() *T.GtkWidget
	ImageNewFromPixmap    func(pixmap *T.GdkPixmap, mask *T.GdkBitmap) *T.GtkWidget
	ImageNewFromImage     func(image *T.GdkImage, mask *T.GdkBitmap) *T.GtkWidget
	ImageNewFromFile      func(filename string) *T.GtkWidget
	ImageNewFromPixbuf    func(pixbuf *T.GdkPixbuf) *T.GtkWidget
	ImageNewFromStock     func(stockId string, size IconSize) *T.GtkWidget
	ImageNewFromIconSet   func(iconSet *IconSet, size IconSize) *T.GtkWidget
	ImageNewFromAnimation func(animation *T.GdkPixbufAnimation) *T.GtkWidget
	ImageNewFromIconName  func(iconName string, size IconSize) *T.GtkWidget
	ImageNewFromGicon     func(icon *T.GIcon, size IconSize) *T.GtkWidget

	ImageClear            func(i *Image)
	ImageGet              func(i *Image, val **T.GdkImage, mask **T.GdkBitmap)
	ImageGetAnimation     func(i *Image) *T.GdkPixbufAnimation
	ImageGetGicon         func(i *Image, gicon **T.GIcon, size *IconSize)
	ImageGetIconName      func(i *Image, iconName **T.Gchar, size *IconSize)
	ImageGetIconSet       func(i *Image, iconSet **IconSet, size *IconSize)
	ImageGetImage         func(i *Image, gdkImage **T.GdkImage, mask **T.GdkBitmap)
	ImageGetPixbuf        func(i *Image) *T.GdkPixbuf
	ImageGetPixelSize     func(i *Image) int
	ImageGetPixmap        func(i *Image, pixmap **T.GdkPixmap, mask **T.GdkBitmap)
	ImageGetStock         func(i *Image, stockId **T.Gchar, size *IconSize)
	ImageGetStorageType   func(i *Image) ImageType
	ImageSet              func(i *Image, val *T.GdkImage, mask *T.GdkBitmap)
	ImageSetFromAnimation func(i *Image, animation *T.GdkPixbufAnimation)
	ImageSetFromFile      func(i *Image, filename string)
	ImageSetFromGicon     func(i *Image, icon *T.GIcon, size IconSize)
	ImageSetFromIconName  func(i *Image, iconName string, size IconSize)
	ImageSetFromIconSet   func(i *Image, iconSet *IconSet, size IconSize)
	ImageSetFromImage     func(i *Image, gdkImage *T.GdkImage, mask *T.GdkBitmap)
	ImageSetFromPixbuf    func(i *Image, pixbuf *T.GdkPixbuf)
	ImageSetFromPixmap    func(i *Image, pixmap *T.GdkPixmap, mask *T.GdkBitmap)
	ImageSetFromStock     func(i *Image, stockId string, size IconSize)
	ImageSetPixelSize     func(i *Image, pixelSize int)
)

func (i *Image) Clear() { ImageClear(i) }

func (i *Image) SetFromPixmap(
	pixmap *T.GdkPixmap, mask *T.GdkBitmap) {
	ImageSetFromPixmap(i, pixmap, mask)
}

func (i *Image) SetFromImage(
	gdkImage *T.GdkImage, mask *T.GdkBitmap) {
	ImageSetFromImage(i, gdkImage, mask)
}

func (i *Image) SetFromFile(filename string) {
	ImageSetFromFile(i, filename)
}

func (i *Image) SetFromPixbuf(pixbuf *T.GdkPixbuf) {
	ImageSetFromPixbuf(i, pixbuf)
}

func (i *Image) SetFromStock(stockId string, size IconSize) {
	ImageSetFromStock(i, stockId, size)
}

func (i *Image) SetFromIconSet(iconSet *IconSet, size IconSize) {
	ImageSetFromIconSet(i, iconSet, size)
}

func (i *Image) SetFromAnimation(animation *T.GdkPixbufAnimation) {
	ImageSetFromAnimation(i, animation)
}

func (i *Image) SetFromIconName(iconName string, size IconSize) {
	ImageSetFromIconName(i, iconName, size)
}

func (i *Image) SetFromGicon(icon *T.GIcon, size IconSize) {
	ImageSetFromGicon(i, icon, size)
}

func (i *Image) SetPixelSize(pixelSize int) {
	ImageSetPixelSize(i, pixelSize)
}

func (i *Image) GetStorageType() ImageType {
	return ImageGetStorageType(i)
}

func (i *Image) GetPixmap(
	pixmap **T.GdkPixmap, mask **T.GdkBitmap) {
	ImageGetPixmap(i, pixmap, mask)
}

func (i *Image) GetImage(
	gdkImage **T.GdkImage, mask **T.GdkBitmap) {
	ImageGetImage(i, gdkImage, mask)
}

func (i *Image) GetPixbuf() *T.GdkPixbuf {
	return ImageGetPixbuf(i)
}

func (i *Image) GetStock(stockId **T.Gchar, size *IconSize) {
	ImageGetStock(i, stockId, size)
}

func (i *Image) GetIconSet(iconSet **IconSet, size *IconSize) {
	ImageGetIconSet(i, iconSet, size)
}

func (i *Image) GetAnimation() *T.GdkPixbufAnimation {
	return ImageGetAnimation(i)
}

func (i *Image) GetIconName(iconName **T.Gchar, size *IconSize) {
	ImageGetIconName(i, iconName, size)
}

func (i *Image) GetGicon(gicon **T.GIcon, size *IconSize) {
	ImageGetGicon(i, gicon, size)
}

func (i *Image) GetPixelSize() int {
	return ImageGetPixelSize(i)
}

func (i *Image) Set(val *T.GdkImage, mask *T.GdkBitmap) {
	ImageSet(i, val, mask)
}

func (i *Image) Get(val **T.GdkImage, mask **T.GdkBitmap) {
	ImageGet(i, val, mask)
}

type ImageMenuItem struct {
	MenuItem T.GtkMenuItem
	Image    *T.GtkWidget
}

var (
	ImageMenuItemGetType         func() T.GType
	ImageMenuItemNew             func() *T.GtkWidget
	ImageMenuItemNewWithLabel    func(label string) *T.GtkWidget
	ImageMenuItemNewWithMnemonic func(label string) *T.GtkWidget
	ImageMenuItemNewFromStock    func(stockId string, accelGroup *AccelGroup) *T.GtkWidget

	ImageMenuItemGetAlwaysShowImage func(i *ImageMenuItem) T.Gboolean
	ImageMenuItemGetImage           func(i *ImageMenuItem) *T.GtkWidget
	ImageMenuItemGetUseStock        func(i *ImageMenuItem) T.Gboolean
	ImageMenuItemSetAccelGroup      func(i *ImageMenuItem, accelGroup *AccelGroup)
	ImageMenuItemSetAlwaysShowImage func(i *ImageMenuItem, alwaysShow T.Gboolean)
	ImageMenuItemSetImage           func(i *ImageMenuItem, image *T.GtkWidget)
	ImageMenuItemSetUseStock        func(i *ImageMenuItem, useStock T.Gboolean)
)

func (i *ImageMenuItem) SetAlwaysShowImage(alwaysShow T.Gboolean) {
	ImageMenuItemSetAlwaysShowImage(i, alwaysShow)
}

func (i *ImageMenuItem) GetAlwaysShowImage() T.Gboolean {
	return ImageMenuItemGetAlwaysShowImage(i)
}

func (i *ImageMenuItem) SetImage(image *T.GtkWidget) {
	ImageMenuItemSetImage(i, image)
}

func (i *ImageMenuItem) GetImage() *T.GtkWidget {
	return ImageMenuItemGetImage(i)
}

func (i *ImageMenuItem) SetUseStock(useStock T.Gboolean) {
	ImageMenuItemSetUseStock(i, useStock)
}

func (i *ImageMenuItem) GetUseStock() T.Gboolean {
	return ImageMenuItemGetUseStock(i)
}

func (i *ImageMenuItem) SetAccelGroup(accelGroup *AccelGroup) {
	ImageMenuItemSetAccelGroup(i, accelGroup)
}

type IMContext simpleObject

var (
	ImContextGetType func() T.GType

	ImContextDeleteSurrounding func(i *IMContext, offset int, nChars int) T.Gboolean
	ImContextFilterKeypress    func(i *IMContext, event *T.GdkEventKey) T.Gboolean
	ImContextFocusIn           func(i *IMContext)
	ImContextFocusOut          func(i *IMContext)
	ImContextGetPreeditString  func(i *IMContext, str **T.Gchar, attrs **T.PangoAttrList, cursorPos *int)
	ImContextGetSurrounding    func(i *IMContext, text **T.Gchar, cursorIndex *int) T.Gboolean
	ImContextReset             func(i *IMContext)
	ImContextSetClientWindow   func(i *IMContext, window *T.GdkWindow)
	ImContextSetCursorLocation func(i *IMContext, area *T.GdkRectangle)
	ImContextSetSurrounding    func(i *IMContext, text string, len int, cursorIndex int)
	ImContextSetUsePreedit     func(i *IMContext, usePreedit T.Gboolean)
)

func (i *IMContext) SetClientWindow(window *T.GdkWindow) {
	ImContextSetClientWindow(i, window)
}

func (i *IMContext) GetPreeditString(
	str **T.Gchar, attrs **T.PangoAttrList, cursorPos *int) {
	ImContextGetPreeditString(i, str, attrs, cursorPos)
}

func (i *IMContext) FilterKeypress(
	event *T.GdkEventKey) T.Gboolean {
	return ImContextFilterKeypress(i, event)
}

func (i *IMContext) FocusIn() {
	ImContextFocusIn(i)
}

func (i *IMContext) FocusOut() {
	ImContextFocusOut(i)
}

func (i *IMContext) Reset() {
	ImContextReset(i)
}

func (i *IMContext) SetCursorLocation(area *T.GdkRectangle) {
	ImContextSetCursorLocation(i, area)
}

func (i *IMContext) SetUsePreedit(usePreedit T.Gboolean) {
	ImContextSetUsePreedit(i, usePreedit)
}

func (i *IMContext) SetSurrounding(
	text string, leng int, cursorIndex int) {
	ImContextSetSurrounding(i, text, leng, cursorIndex)
}

func (i *IMContext) GetSurrounding(
	text **T.Gchar, cursorIndex *int) T.Gboolean {
	return ImContextGetSurrounding(i, text, cursorIndex)
}

func (i *IMContext) DeleteSurrounding(
	offset, nChars int) T.Gboolean {
	return ImContextDeleteSurrounding(i, offset, nChars)
}

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

	ImContextSimpleAddTable func(i *IMContextSimple, data *uint16, maxSeqLen, nSeqs int)
)

func (i *IMContextSimple) AddTable(data *uint16, maxSeqLen, nSeqs int) {
	ImContextSimpleAddTable(i, data, maxSeqLen, nSeqs)
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

	ImMulticontextAppendMenuitems func(i *IMMulticontext, menushell *T.GtkMenuShell)
	ImMulticontextGetContextId    func(i *IMMulticontext) string
	ImMulticontextSetContextId    func(i *IMMulticontext, contextId string)
)

func (i *IMMulticontext) AppendMenuitems(menushell *T.GtkMenuShell) {
	ImMulticontextAppendMenuitems(i, menushell)
}

func (i *IMMulticontext) GetContextId() string {
	return ImMulticontextGetContextId(i)
}

func (i *IMMulticontext) SetContextId(contextId string) {
	ImMulticontextSetContextId(i, contextId)
}

type InfoBar struct {
	Parent T.GtkHBox
	_      *struct{}
}

var (
	InfoBarGetType        func() T.GType
	InfoBarNew            func() *T.GtkWidget
	InfoBarNewWithButtons func(firstButtonText string, v ...VArg) *T.GtkWidget

	InfoBarAddActionWidget      func(i *InfoBar, child *T.GtkWidget, responseId int)
	InfoBarAddButton            func(i *InfoBar, buttonText string, responseId int) *T.GtkWidget
	InfoBarAddButtons           func(i *InfoBar, firstButtonText string, v ...VArg)
	InfoBarGetActionArea        func(i *InfoBar) *T.GtkWidget
	InfoBarGetContentArea       func(i *InfoBar) *T.GtkWidget
	InfoBarGetMessageType       func(i *InfoBar) T.GtkMessageType
	InfoBarResponse             func(i *InfoBar, responseId int)
	InfoBarSetDefaultResponse   func(i *InfoBar, responseId int)
	InfoBarSetMessageType       func(i *InfoBar, messageType T.GtkMessageType)
	InfoBarSetResponseSensitive func(i *InfoBar, responseId int, setting T.Gboolean)
)

func (i *InfoBar) GetActionArea() *T.GtkWidget {
	return InfoBarGetActionArea(i)
}

func (i *InfoBar) GetContentArea() *T.GtkWidget {
	return InfoBarGetContentArea(i)
}

func (i *InfoBar) AddActionWidget(child *T.GtkWidget, responseId int) {
	InfoBarAddActionWidget(i, child, responseId)
}

func (i *InfoBar) AddButton(buttonText string, responseId int) *T.GtkWidget {
	return InfoBarAddButton(i, buttonText, responseId)
}

func (i *InfoBar) AddButtons(firstButtonText string, v ...VArg) {
	InfoBarAddButtons(i, firstButtonText, v)
}

func (i *InfoBar) SetResponseSensitive(responseId int, setting T.Gboolean) {
	InfoBarSetResponseSensitive(i, responseId, setting)
}

func (i *InfoBar) SetDefaultResponse(responseId int) {
	InfoBarSetDefaultResponse(i, responseId)
}

func (i *InfoBar) Response(responseId int) {
	InfoBarResponse(i, responseId)
}

func (i *InfoBar) SetMessageType(messageType T.GtkMessageType) {
	InfoBarSetMessageType(i, messageType)
}

func (i *InfoBar) GetMessageType() T.GtkMessageType {
	return InfoBarGetMessageType(i)
}

type Invisible struct {
	Widget          T.GtkWidget
	HasUserRefCount T.Gboolean
	Screen          *T.GdkScreen
}

var (
	InvisibleGetType      func() T.GType
	InvisibleNew          func() *T.GtkWidget
	InvisibleNewForScreen func(screen *T.GdkScreen) *T.GtkWidget

	InvisibleGetScreen func(i *Invisible) *T.GdkScreen
	InvisibleSetScreen func(i *Invisible, screen *T.GdkScreen)
)

func (i *Invisible) SetScreen(screen *T.GdkScreen) {
	InvisibleSetScreen(i, screen)
}

func (i *Invisible) GetScreen() *T.GdkScreen {
	return InvisibleGetScreen(i)
}

type Item struct {
	Bin Bin
}

var (
	ItemGetType func() T.GType

	ItemDeselect func(i *Item)
	ItemSelect   func(i *Item)
	ItemToggle   func(i *Item)
)

func (i *Item) Select() { ItemSelect(i) }

func (i *Item) Deselect() { ItemDeselect(i) }

func (i *Item) Toggle() { ItemToggle(i) }

type (
	ItemFactory struct {
		Object           T.GtkObject
		Path             *T.Gchar
		Accel_group      *AccelGroup
		Widget           *T.GtkWidget
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
	ItemFactoryAddForeign          func(accelWidget *T.GtkWidget, fullPath string, accelGroup *AccelGroup, keyval uint, modifiers T.GdkModifierType)
	ItemFactoryCreateMenuEntries   func(nEntries uint, entries *T.GtkMenuEntry)
	ItemFactoryFromPath            func(path string) *ItemFactory
	ItemFactoryFromWidget          func(widget *T.GtkWidget) *ItemFactory
	ItemFactoryPathFromWidget      func(widget *T.GtkWidget) string
	ItemFactoryPopupDataFromWidget func(widget *T.GtkWidget) T.Gpointer

	ItemFactoryConstruct         func(i *ItemFactory, containerType T.GType, path string, accelGroup *AccelGroup)
	ItemFactoryCreateItem        func(i *ItemFactory, entry *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint)
	ItemFactoryCreateItems       func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer)
	ItemFactoryCreateItemsAc     func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint)
	ItemFactoryDeleteEntries     func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry)
	ItemFactoryDeleteEntry       func(i *ItemFactory, entry *ItemFactoryEntry)
	ItemFactoryDeleteItem        func(i *ItemFactory, path string)
	ItemFactoryGetItem           func(i *ItemFactory, path string) *T.GtkWidget
	ItemFactoryGetItemByAction   func(i *ItemFactory, action uint) *T.GtkWidget
	ItemFactoryGetWidget         func(i *ItemFactory, path string) *T.GtkWidget
	ItemFactoryGetWidgetByAction func(i *ItemFactory, action uint) *T.GtkWidget
	ItemFactoryPopup             func(i *ItemFactory, x uint, y uint, mouseButton uint, time T.GUint32)
	ItemFactoryPopupData         func(i *ItemFactory) T.Gpointer
	ItemFactoryPopupWithData     func(i *ItemFactory, popupData T.Gpointer, destroy T.GDestroyNotify, x uint, y uint, mouseButton uint, time T.GUint32)
	ItemFactorySetTranslateFunc  func(i *ItemFactory, f T.GtkTranslateFunc, data T.Gpointer, notify T.GDestroyNotify)
)

func (i *ItemFactory) Construct(containerType T.GType,
	path string, accelGroup *AccelGroup) {
	ItemFactoryConstruct(i, containerType, path, accelGroup)
}

func (i *ItemFactory) GetItem(path string) *T.GtkWidget {
	return ItemFactoryGetItem(i, path)
}

func (i *ItemFactory) GetWidget(path string) *T.GtkWidget {
	return ItemFactoryGetWidget(i, path)
}

func (i *ItemFactory) GetWidgetByAction(action uint) *T.GtkWidget {
	return ItemFactoryGetWidgetByAction(i, action)
}

func (i *ItemFactory) GetItemByAction(action uint) *T.GtkWidget {
	return ItemFactoryGetItemByAction(i, action)
}

func (i *ItemFactory) CreateItem(entry *ItemFactoryEntry,
	callbackData T.Gpointer, callbackType uint) {
	ItemFactoryCreateItem(i, entry, callbackData, callbackType)
}

func (i *ItemFactory) CreateItems(nEntries uint,
	entries *ItemFactoryEntry, callbackData T.Gpointer) {
	ItemFactoryCreateItems(i, nEntries, entries, callbackData)
}

func (i *ItemFactory) DeleteItem(path string) {
	ItemFactoryDeleteItem(i, path)
}

func (i *ItemFactory) DeleteEntry(entry *ItemFactoryEntry) {
	ItemFactoryDeleteEntry(i, entry)
}

func (i *ItemFactory) DeleteEntries(
	nEntries uint, entries *ItemFactoryEntry) {
	ItemFactoryDeleteEntries(i, nEntries, entries)
}

func (i *ItemFactory) Popup(
	x, y, mouseButton uint, time T.GUint32) {
	ItemFactoryPopup(i, x, y, mouseButton, time)
}

func (i *ItemFactory) PopupWithData(popupData T.Gpointer,
	destroy T.GDestroyNotify, x, y, mouseButton uint,
	time T.GUint32) {
	ItemFactoryPopupWithData(
		i, popupData, destroy, x, y, mouseButton, time)
}

func (i *ItemFactory) PopupData() T.Gpointer {
	return ItemFactoryPopupData(i)
}

func (i *ItemFactory) SetTranslateFunc(
	f T.GtkTranslateFunc, data T.Gpointer, notify T.GDestroyNotify) {
	ItemFactorySetTranslateFunc(i, f, data, notify)
}

func (i *ItemFactory) CreateItemsAc(
	nEntries uint, entries *ItemFactoryEntry,
	callbackData T.Gpointer, callbackType uint) {
	ItemFactoryCreateItemsAc(
		i, nEntries, entries, callbackData, callbackType)
}
