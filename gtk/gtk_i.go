// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	I "github.com/tHinqa/outside-gtk2/gio"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type IconFactory struct {
	Parent T.GObject
	Icons  *T.GHashTable
}

var (
	IconFactoryGetType func() O.Type
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
	IconInfoGetType      func() O.Type
	IconInfoNewForPixbuf func(iconTheme *IconTheme, pixbuf *D.Pixbuf) *IconInfo

	iconInfoCopy              func(i *IconInfo) *IconInfo
	iconInfoFree              func(i *IconInfo)
	iconInfoGetAttachPoints   func(i *IconInfo, points **D.Point, nPoints *int) bool
	iconInfoGetBaseSize       func(i *IconInfo) int
	iconInfoGetBuiltinPixbuf  func(i *IconInfo) *D.Pixbuf
	iconInfoGetDisplayName    func(i *IconInfo) string
	iconInfoGetEmbeddedRect   func(i *IconInfo, rectangle *D.Rectangle) bool
	iconInfoGetFilename       func(i *IconInfo) string
	iconInfoLoadIcon          func(i *IconInfo, error **T.GError) *D.Pixbuf
	iconInfoSetRawCoordinates func(i *IconInfo, rawCoordinates bool)
)

func (i *IconInfo) Copy() *IconInfo { return iconInfoCopy(i) }
func (i *IconInfo) Free()           { iconInfoFree(i) }
func (i *IconInfo) GetAttachPoints(points **D.Point, nPoints *int) bool {
	return iconInfoGetAttachPoints(i, points, nPoints)
}
func (i *IconInfo) GetBaseSize() int            { return iconInfoGetBaseSize(i) }
func (i *IconInfo) GetBuiltinPixbuf() *D.Pixbuf { return iconInfoGetBuiltinPixbuf(i) }
func (i *IconInfo) GetDisplayName() string      { return iconInfoGetDisplayName(i) }
func (i *IconInfo) GetEmbeddedRect(rectangle *D.Rectangle) bool {
	return iconInfoGetEmbeddedRect(i, rectangle)
}
func (i *IconInfo) GetFilename() string               { return iconInfoGetFilename(i) }
func (i *IconInfo) LoadIcon(err **T.GError) *D.Pixbuf { return iconInfoLoadIcon(i, err) }
func (i *IconInfo) SetRawCoordinates(rawCoordinates bool) {
	iconInfoSetRawCoordinates(i, rawCoordinates)
}

type IconLookupFlags Enum

const (
	ICON_LOOKUP_NO_SVG IconLookupFlags = 1 << iota
	ICON_LOOKUP_FORCE_SVG
	ICON_LOOKUP_USE_BUILTIN
	ICON_LOOKUP_GENERIC_FALLBACK
	ICON_LOOKUP_FORCE_SIZE
)

var IconLookupFlagsGetType func() O.Type

type IconSet struct{}

var (
	IconSetGetType       func() O.Type
	IconSetNew           func() *IconSet
	IconSetNewFromPixbuf func(pixbuf *D.Pixbuf) *IconSet

	iconSetAddSource  func(i *IconSet, source *IconSource)
	iconSetCopy       func(i *IconSet) *IconSet
	iconSetGetSizes   func(i *IconSet, sizes **IconSize, nSizes *int)
	iconSetRef        func(i *IconSet) *IconSet
	iconSetRenderIcon func(i *IconSet, style *Style, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *D.Pixbuf
	iconSetUnref      func(i *IconSet)
)

func (i *IconSet) AddSource(source *IconSource)           { iconSetAddSource(i, source) }
func (i *IconSet) Copy() *IconSet                         { return iconSetCopy(i) }
func (i *IconSet) GetSizes(sizes **IconSize, nSizes *int) { iconSetGetSizes(i, sizes, nSizes) }
func (i *IconSet) Ref() *IconSet                          { return iconSetRef(i) }
func (i *IconSet) RenderIcon(style *Style, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *D.Pixbuf {
	return iconSetRenderIcon(i, style, direction, state, size, widget, detail)
}
func (i *IconSet) Unref() { iconSetUnref(i) }

type IconSize Enum

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
	IconSizeGetType func() O.Type

	IconSizeFromName      func(name string) IconSize
	IconSizeRegister      func(name string, width int, height int) IconSize
	IconSizeRegisterAlias func(alias string, target IconSize)

	iconSizeGetName func(i IconSize) string
	iconSizeLookup  func(i IconSize, width *int, height *int) bool
)

func (i IconSize) GetName() string                     { return iconSizeGetName(i) }
func (i IconSize) Lookup(width *int, height *int) bool { return iconSizeLookup(i, width, height) }

type IconSource struct{}

var (
	IconSourceGetType func() O.Type
	IconSourceNew     func() *IconSource

	iconSourceCopy                   func(i *IconSource) *IconSource
	iconSourceFree                   func(i *IconSource)
	iconSourceGetDirection           func(i *IconSource) TextDirection
	iconSourceGetDirectionWildcarded func(i *IconSource) bool
	iconSourceGetFilename            func(i *IconSource) string
	iconSourceGetIconName            func(i *IconSource) string
	iconSourceGetPixbuf              func(i *IconSource) *D.Pixbuf
	iconSourceGetSize                func(i *IconSource) IconSize
	iconSourceGetSizeWildcarded      func(i *IconSource) bool
	iconSourceGetState               func(i *IconSource) StateType
	iconSourceGetStateWildcarded     func(i *IconSource) bool
	iconSourceSetDirection           func(i *IconSource, direction TextDirection)
	iconSourceSetDirectionWildcarded func(i *IconSource, setting bool)
	iconSourceSetFilename            func(i *IconSource, filename string)
	iconSourceSetIconName            func(i *IconSource, iconName string)
	iconSourceSetPixbuf              func(i *IconSource, pixbuf *D.Pixbuf)
	iconSourceSetSize                func(i *IconSource, size IconSize)
	iconSourceSetSizeWildcarded      func(i *IconSource, setting bool)
	iconSourceSetState               func(i *IconSource, state StateType)
	iconSourceSetStateWildcarded     func(i *IconSource, setting bool)
)

func (i *IconSource) Copy() *IconSource                    { return iconSourceCopy(i) }
func (i *IconSource) Free()                                { iconSourceFree(i) }
func (i *IconSource) GetDirection() TextDirection          { return iconSourceGetDirection(i) }
func (i *IconSource) GetDirectionWildcarded() bool         { return iconSourceGetDirectionWildcarded(i) }
func (i *IconSource) GetFilename() string                  { return iconSourceGetFilename(i) }
func (i *IconSource) GetIconName() string                  { return iconSourceGetIconName(i) }
func (i *IconSource) GetPixbuf() *D.Pixbuf                 { return iconSourceGetPixbuf(i) }
func (i *IconSource) GetSize() IconSize                    { return iconSourceGetSize(i) }
func (i *IconSource) GetSizeWildcarded() bool              { return iconSourceGetSizeWildcarded(i) }
func (i *IconSource) GetState() StateType                  { return iconSourceGetState(i) }
func (i *IconSource) GetStateWildcarded() bool             { return iconSourceGetStateWildcarded(i) }
func (i *IconSource) SetDirection(direction TextDirection) { iconSourceSetDirection(i, direction) }
func (i *IconSource) SetDirectionWildcarded(setting bool) {
	iconSourceSetDirectionWildcarded(i, setting)
}
func (i *IconSource) SetFilename(filename string)     { iconSourceSetFilename(i, filename) }
func (i *IconSource) SetIconName(iconName string)     { iconSourceSetIconName(i, iconName) }
func (i *IconSource) SetPixbuf(pixbuf *D.Pixbuf)      { iconSourceSetPixbuf(i, pixbuf) }
func (i *IconSource) SetSize(size IconSize)           { iconSourceSetSize(i, size) }
func (i *IconSource) SetSizeWildcarded(setting bool)  { iconSourceSetSizeWildcarded(i, setting) }
func (i *IconSource) SetState(state StateType)        { iconSourceSetState(i, state) }
func (i *IconSource) SetStateWildcarded(setting bool) { iconSourceSetStateWildcarded(i, setting) }

type IconTheme struct {
	Parent T.GObject
	_      *struct{}
}

var (
	IconThemeGetType func() O.Type
	IconThemeNew     func() *IconTheme

	IconThemeErrorGetType func() O.Type
	IconThemeErrorQuark   func() T.GQuark
	IconThemeGetDefault   func() *IconTheme

	IconThemeGetForScreen   func(screen *D.Screen) *IconTheme
	IconThemeAddBuiltinIcon func(iconName string, size int, pixbuf *D.Pixbuf)

	iconThemeAppendSearchPath   func(i *IconTheme, path string)
	iconThemeChooseIcon         func(i *IconTheme, iconNames []string, size int, flags IconLookupFlags) *IconInfo
	iconThemeGetExampleIconName func(i *IconTheme) string
	iconThemeGetIconSizes       func(i *IconTheme, iconName string) *int
	iconThemeGetSearchPath      func(i *IconTheme, path ***T.Gchar, nElements *int)
	iconThemeHasIcon            func(i *IconTheme, iconName string) bool
	iconThemeListContexts       func(i *IconTheme) *T.GList
	iconThemeListIcons          func(i *IconTheme, context string) *T.GList
	iconThemeLoadIcon           func(i *IconTheme, iconName string, size int, flags IconLookupFlags, error **T.GError) *D.Pixbuf
	iconThemeLookupByGicon      func(i *IconTheme, icon *I.Icon, size int, flags IconLookupFlags) *IconInfo
	iconThemeLookupIcon         func(i *IconTheme, iconName string, size int, flags IconLookupFlags) *IconInfo
	iconThemePrependSearchPath  func(i *IconTheme, path string)
	iconThemeRescanIfNeeded     func(i *IconTheme) bool
	iconThemeSetCustomTheme     func(i *IconTheme, themeName string)
	iconThemeSetScreen          func(i *IconTheme, screen *D.Screen)
	iconThemeSetSearchPath      func(i *IconTheme, path **T.Gchar, nElements int)
)

func (i *IconTheme) AppendSearchPath(path string) { iconThemeAppendSearchPath(i, path) }
func (i *IconTheme) ChooseIcon(iconNames []string, size int, flags IconLookupFlags) *IconInfo {
	return iconThemeChooseIcon(i, iconNames, size, flags)
}
func (i *IconTheme) GetExampleIconName() string        { return iconThemeGetExampleIconName(i) }
func (i *IconTheme) GetIconSizes(iconName string) *int { return iconThemeGetIconSizes(i, iconName) }
func (i *IconTheme) GetSearchPath(path ***T.Gchar, nElements *int) {
	iconThemeGetSearchPath(i, path, nElements)
}
func (i *IconTheme) HasIcon(iconName string) bool      { return iconThemeHasIcon(i, iconName) }
func (i *IconTheme) ListContexts() *T.GList            { return iconThemeListContexts(i) }
func (i *IconTheme) ListIcons(context string) *T.GList { return iconThemeListIcons(i, context) }
func (i *IconTheme) LoadIcon(iconName string, size int, flags IconLookupFlags, err **T.GError) *D.Pixbuf {
	return iconThemeLoadIcon(i, iconName, size, flags, err)
}
func (i *IconTheme) LookupByGicon(icon *I.Icon, size int, flags IconLookupFlags) *IconInfo {
	return iconThemeLookupByGicon(i, icon, size, flags)
}
func (i *IconTheme) LookupIcon(iconName string, size int, flags IconLookupFlags) *IconInfo {
	return iconThemeLookupIcon(i, iconName, size, flags)
}
func (i *IconTheme) PrependSearchPath(path string)   { iconThemePrependSearchPath(i, path) }
func (i *IconTheme) RescanIfNeeded() bool            { return iconThemeRescanIfNeeded(i) }
func (i *IconTheme) SetCustomTheme(themeName string) { iconThemeSetCustomTheme(i, themeName) }
func (i *IconTheme) SetScreen(screen *D.Screen)      { iconThemeSetScreen(i, screen) }
func (i *IconTheme) SetSearchPath(path **T.Gchar, nElements int) {
	iconThemeSetSearchPath(i, path, nElements)
}

type (
	IconView struct {
		Parent Container
		_      *struct{}
	}

	IconViewForeachFunc func(icon_view *IconView,
		path *TreePath, data T.Gpointer)
)

type IconViewDropPosition Enum

const (
	ICON_VIEW_NO_DROP IconViewDropPosition = iota
	ICON_VIEW_DROP_INTO
	ICON_VIEW_DROP_LEFT
	ICON_VIEW_DROP_RIGHT
	ICON_VIEW_DROP_ABOVE
	ICON_VIEW_DROP_BELOW
)

var (
	IconViewGetType      func() O.Type
	IconViewNew          func() *Widget
	IconViewNewWithModel func(model *TreeModel) *Widget

	IconViewDropPositionGetType func() O.Type

	iconViewConvertWidgetToBinWindowCoords func(i *IconView, wx int, wy int, bx *int, by *int)
	iconViewCreateDragIcon                 func(i *IconView, path *TreePath) *D.Pixmap
	iconViewEnableModelDragDest            func(i *IconView, targets *TargetEntry, nTargets int, actions D.DragAction)
	iconViewEnableModelDragSource          func(i *IconView, startButtonMask T.GdkModifierType, targets *TargetEntry, nTargets int, actions D.DragAction)
	iconViewGetColumns                     func(i *IconView) int
	iconViewGetColumnSpacing               func(i *IconView) int
	iconViewGetCursor                      func(i *IconView, path **TreePath, cell **CellRenderer) bool
	iconViewGetDestItemAtPos               func(i *IconView, dragX, dragY int, path **TreePath, pos *IconViewDropPosition) bool
	iconViewGetDragDestItem                func(i *IconView, path **TreePath, pos *IconViewDropPosition)
	iconViewGetItemAtPos                   func(i *IconView, x, y int, path **TreePath, cell **CellRenderer) bool
	iconViewGetItemColumn                  func(i *IconView, path *TreePath) int
	iconViewGetItemOrientation             func(i *IconView) Orientation
	iconViewGetItemPadding                 func(i *IconView) int
	iconViewGetItemRow                     func(i *IconView, path *TreePath) int
	iconViewGetItemWidth                   func(i *IconView) int
	iconViewGetMargin                      func(i *IconView) int
	iconViewGetMarkupColumn                func(i *IconView) int
	iconViewGetModel                       func(i *IconView) *TreeModel
	iconViewGetOrientation                 func(i *IconView) Orientation
	iconViewGetPathAtPos                   func(i *IconView, x, y int) *TreePath
	iconViewGetPixbufColumn                func(i *IconView) int
	iconViewGetReorderable                 func(i *IconView) bool
	iconViewGetRowSpacing                  func(i *IconView) int
	iconViewGetSelectedItems               func(i *IconView) *T.GList
	iconViewGetSelectionMode               func(i *IconView) SelectionMode
	iconViewGetSpacing                     func(i *IconView) int
	iconViewGetTextColumn                  func(i *IconView) int
	iconViewGetTooltipColumn               func(i *IconView) int
	iconViewGetTooltipContext              func(i *IconView, x *int, y *int, keyboardTip bool, model **TreeModel, path **TreePath, iter *TreeIter) bool
	iconViewGetVisibleRange                func(i *IconView, startPath **TreePath, endPath **TreePath) bool
	iconViewItemActivated                  func(i *IconView, path *TreePath)
	iconViewPathIsSelected                 func(i *IconView, path *TreePath) bool
	iconViewScrollToPath                   func(i *IconView, path *TreePath, useAlign bool, rowAlign, colAlign float32)
	iconViewSelectAll                      func(i *IconView)
	iconViewSelectedForeach                func(i *IconView, f IconViewForeachFunc, data T.Gpointer)
	iconViewSelectPath                     func(i *IconView, path *TreePath)
	iconViewSetColumns                     func(i *IconView, columns int)
	iconViewSetColumnSpacing               func(i *IconView, columnSpacing int)
	iconViewSetCursor                      func(i *IconView, path *TreePath, cell *CellRenderer, startEditing bool)
	iconViewSetDragDestItem                func(i *IconView, path *TreePath, pos IconViewDropPosition)
	iconViewSetItemOrientation             func(i *IconView, orientation Orientation)
	iconViewSetItemPadding                 func(i *IconView, itemPadding int)
	iconViewSetItemWidth                   func(i *IconView, itemWidth int)
	iconViewSetMargin                      func(i *IconView, margin int)
	iconViewSetMarkupColumn                func(i *IconView, column int)
	iconViewSetModel                       func(i *IconView, model *TreeModel)
	iconViewSetOrientation                 func(i *IconView, orientation Orientation)
	iconViewSetPixbufColumn                func(i *IconView, column int)
	iconViewSetReorderable                 func(i *IconView, reorderable bool)
	iconViewSetRowSpacing                  func(i *IconView, rowSpacing int)
	iconViewSetSelectionMode               func(i *IconView, mode SelectionMode)
	iconViewSetSpacing                     func(i *IconView, spacing int)
	iconViewSetTextColumn                  func(i *IconView, column int)
	iconViewSetTooltipCell                 func(i *IconView, tooltip *Tooltip, path *TreePath, cell *CellRenderer)
	iconViewSetTooltipColumn               func(i *IconView, column int)
	iconViewSetTooltipItem                 func(i *IconView, tooltip *Tooltip, path *TreePath)
	iconViewUnselectAll                    func(i *IconView)
	iconViewUnselectPath                   func(i *IconView, path *TreePath)
	iconViewUnsetModelDragDest             func(i *IconView)
	iconViewUnsetModelDragSource           func(i *IconView)
)

func (i *IconView) ConvertWidgetToBinWindowCoords(wx, wy int, bx, by *int) {
	iconViewConvertWidgetToBinWindowCoords(i, wx, wy, bx, by)
}
func (i *IconView) CreateDragIcon(path *TreePath) *D.Pixmap {
	return iconViewCreateDragIcon(i, path)
}
func (i *IconView) EnableModelDragDest(targets *TargetEntry, nTargets int, actions D.DragAction) {
	iconViewEnableModelDragDest(i, targets, nTargets, actions)
}
func (i *IconView) EnableModelDragSource(startButtonMask T.GdkModifierType, targets *TargetEntry, nTargets int, actions D.DragAction) {
	iconViewEnableModelDragSource(i, startButtonMask, targets, nTargets, actions)
}
func (i *IconView) GetColumns() int       { return iconViewGetColumns(i) }
func (i *IconView) GetColumnSpacing() int { return iconViewGetColumnSpacing(i) }
func (i *IconView) GetCursor(path **TreePath, cell **CellRenderer) bool {
	return iconViewGetCursor(i, path, cell)
}
func (i *IconView) GetDestItemAtPos(dragX, dragY int, path **TreePath, pos *IconViewDropPosition) bool {
	return iconViewGetDestItemAtPos(i, dragX, dragY, path, pos)
}
func (i *IconView) GetDragDestItem(path **TreePath, pos *IconViewDropPosition) {
	iconViewGetDragDestItem(i, path, pos)
}
func (i *IconView) GetItemAtPos(x, y int, path **TreePath, cell **CellRenderer) bool {
	return iconViewGetItemAtPos(i, x, y, path, cell)
}
func (i *IconView) GetItemColumn(path *TreePath) int { return iconViewGetItemColumn(i, path) }
func (i *IconView) GetItemOrientation() Orientation  { return iconViewGetItemOrientation(i) }
func (i *IconView) GetItemPadding() int              { return iconViewGetItemPadding(i) }
func (i *IconView) GetItemRow(path *TreePath) int    { return iconViewGetItemRow(i, path) }
func (i *IconView) GetItemWidth() int                { return iconViewGetItemWidth(i) }
func (i *IconView) GetMargin() int                   { return iconViewGetMargin(i) }
func (i *IconView) GetMarkupColumn() int             { return iconViewGetMarkupColumn(i) }
func (i *IconView) GetModel() *TreeModel             { return iconViewGetModel(i) }
func (i *IconView) GetOrientation() Orientation      { return iconViewGetOrientation(i) }
func (i *IconView) GetPathAtPos(x, y int) *TreePath  { return iconViewGetPathAtPos(i, x, y) }
func (i *IconView) GetPixbufColumn() int             { return iconViewGetPixbufColumn(i) }
func (i *IconView) GetReorderable() bool             { return iconViewGetReorderable(i) }
func (i *IconView) GetRowSpacing() int               { return iconViewGetRowSpacing(i) }
func (i *IconView) GetSelectedItems() *T.GList       { return iconViewGetSelectedItems(i) }
func (i *IconView) GetSelectionMode() SelectionMode  { return iconViewGetSelectionMode(i) }
func (i *IconView) GetSpacing() int                  { return iconViewGetSpacing(i) }
func (i *IconView) GetTextColumn() int               { return iconViewGetTextColumn(i) }
func (i *IconView) GetTooltipColumn() int            { return iconViewGetTooltipColumn(i) }
func (i *IconView) GetTooltipContext(x, y *int, keyboardTip bool, model **TreeModel, path **TreePath, iter *TreeIter) bool {
	return iconViewGetTooltipContext(i, x, y, keyboardTip, model, path, iter)
}
func (i *IconView) GetVisibleRange(startPath **TreePath, endPath **TreePath) bool {
	return iconViewGetVisibleRange(i, startPath, endPath)
}
func (i *IconView) ItemActivated(path *TreePath) { iconViewItemActivated(i, path) }
func (i *IconView) PathIsSelected(path *TreePath) bool {
	return iconViewPathIsSelected(i, path)
}
func (i *IconView) ScrollToPath(path *TreePath, useAlign bool, rowAlign, colAlign float32) {
	iconViewScrollToPath(i, path, useAlign, rowAlign, colAlign)
}
func (i *IconView) SelectAll() { iconViewSelectAll(i) }
func (i *IconView) SelectedForeach(f IconViewForeachFunc, data T.Gpointer) {
	iconViewSelectedForeach(i, f, data)
}
func (i *IconView) SelectPath(path *TreePath)          { iconViewSelectPath(i, path) }
func (i *IconView) SetColumns(columns int)             { iconViewSetColumns(i, columns) }
func (i *IconView) SetColumnSpacing(columnSpacing int) { iconViewSetColumnSpacing(i, columnSpacing) }
func (i *IconView) SetCursor(path *TreePath, cell *CellRenderer, startEditing bool) {
	iconViewSetCursor(i, path, cell, startEditing)
}
func (i *IconView) SetDragDestItem(path *TreePath, pos IconViewDropPosition) {
	iconViewSetDragDestItem(i, path, pos)
}
func (i *IconView) SetItemOrientation(orientation Orientation) {
	iconViewSetItemOrientation(i, orientation)
}
func (i *IconView) SetItemPadding(itemPadding int) { iconViewSetItemPadding(i, itemPadding) }
func (i *IconView) SetItemWidth(itemWidth int)     { iconViewSetItemWidth(i, itemWidth) }
func (i *IconView) SetMargin(margin int)           { iconViewSetMargin(i, margin) }
func (i *IconView) SetMarkupColumn(column int)     { iconViewSetMarkupColumn(i, column) }
func (i *IconView) SetModel(model *TreeModel)      { iconViewSetModel(i, model) }
func (i *IconView) SetOrientation(orientation Orientation) {
	iconViewSetOrientation(i, orientation)
}
func (i *IconView) SetPixbufColumn(column int)          { iconViewSetPixbufColumn(i, column) }
func (i *IconView) SetReorderable(reorderable bool)     { iconViewSetReorderable(i, reorderable) }
func (i *IconView) SetRowSpacing(rowSpacing int)        { iconViewSetRowSpacing(i, rowSpacing) }
func (i *IconView) SetSelectionMode(mode SelectionMode) { iconViewSetSelectionMode(i, mode) }
func (i *IconView) SetSpacing(spacing int)              { iconViewSetSpacing(i, spacing) }
func (i *IconView) SetTextColumn(column int)            { iconViewSetTextColumn(i, column) }
func (i *IconView) SetTooltipCell(tooltip *Tooltip, path *TreePath, cell *CellRenderer) {
	iconViewSetTooltipCell(i, tooltip, path, cell)
}
func (i *IconView) SetTooltipColumn(column int) { iconViewSetTooltipColumn(i, column) }
func (i *IconView) SetTooltipItem(tooltip *Tooltip, path *TreePath) {
	iconViewSetTooltipItem(i, tooltip, path)
}
func (i *IconView) UnselectAll()                { iconViewUnselectAll(i) }
func (i *IconView) UnselectPath(path *TreePath) { iconViewUnselectPath(i, path) }
func (i *IconView) UnsetModelDragDest()         { iconViewUnsetModelDragDest(i) }
func (i *IconView) UnsetModelDragSource()       { iconViewUnsetModelDragSource(i) }

var IdentifierGetType func() O.Type

//TODO(t):Fix
type Image struct {
	Misc        Misc
	StorageType ImageType
	// union
	Pixmap ImagePixmapData
	// 	Image  ImageImageData
	// 	Pixbuf  ImagePixbufData
	// 	Stock  ImageStockData
	// 	IconSet  ImageIconSetData
	// 	Anim  ImageAnimationData
	// 	Name  ImageIconNameData
	// 	Gicon  ImageGIconData
	// }
	Mask     *T.GdkBitmap
	IconSize IconSize
}

type (
	ImagePixmapData    struct{ Pixmap *D.Pixmap }
	ImageImageData     struct{ Image *D.Image }
	ImagePixbufData    struct{ Pixbuf *D.Pixbuf }
	ImageStockData     struct{ StockId *T.Gchar }
	ImageIconSetData   struct{ IconSet *IconSet }
	ImageAnimationData struct {
		Anim         *D.PixbufAnimation
		Iter         *D.PixbufAnimationIter
		FrameTimeout uint
	}
	ImageIconNameData struct {
		IconName      *T.Gchar
		Pixbuf        *D.Pixbuf
		ThemeChangeId uint
	}
	ImageGIconData struct {
		icon          *I.Icon
		Pixbuf        *D.Pixbuf
		ThemeChangeId uint
	}
	ImageClass struct {
		Parent     MiscClass
		_, _, _, _ func()
	}
)

var (
	ImageGetType          func() O.Type
	ImageNew              func() *Widget
	ImageNewFromPixmap    func(pixmap *D.Pixmap, mask *T.GdkBitmap) *Widget
	ImageNewFromImage     func(image *D.Image, mask *T.GdkBitmap) *Widget
	ImageNewFromFile      func(filename string) *Widget
	ImageNewFromPixbuf    func(pixbuf *D.Pixbuf) *Widget
	ImageNewFromStock     func(stockId string, size IconSize) *Widget
	ImageNewFromIconSet   func(iconSet *IconSet, size IconSize) *Widget
	ImageNewFromAnimation func(animation *D.PixbufAnimation) *Widget
	ImageNewFromIconName  func(iconName string, size IconSize) *Widget
	ImageNewFromGicon     func(icon *I.Icon, size IconSize) *Widget

	imageClear            func(i *Image)
	imageGet              func(i *Image, val **D.Image, mask **T.GdkBitmap)
	imageGetAnimation     func(i *Image) *D.PixbufAnimation
	imageGetGicon         func(i *Image, gicon **I.Icon, size *IconSize)
	imageGetIconName      func(i *Image, iconName **T.Gchar, size *IconSize)
	imageGetIconSet       func(i *Image, iconSet **IconSet, size *IconSize)
	imageGetImage         func(i *Image, gdkImage **D.Image, mask **T.GdkBitmap)
	imageGetPixbuf        func(i *Image) *D.Pixbuf
	imageGetPixelSize     func(i *Image) int
	imageGetPixmap        func(i *Image, pixmap **D.Pixmap, mask **T.GdkBitmap)
	imageGetStock         func(i *Image, stockId **T.Gchar, size *IconSize)
	imageGetStorageType   func(i *Image) ImageType
	imageSet              func(i *Image, val *D.Image, mask *T.GdkBitmap)
	imageSetFromAnimation func(i *Image, animation *D.PixbufAnimation)
	imageSetFromFile      func(i *Image, filename string)
	imageSetFromGicon     func(i *Image, icon *I.Icon, size IconSize)
	imageSetFromIconName  func(i *Image, iconName string, size IconSize)
	imageSetFromIconSet   func(i *Image, iconSet *IconSet, size IconSize)
	imageSetFromImage     func(i *Image, gdkImage *D.Image, mask *T.GdkBitmap)
	imageSetFromPixbuf    func(i *Image, pixbuf *D.Pixbuf)
	imageSetFromPixmap    func(i *Image, pixmap *D.Pixmap, mask *T.GdkBitmap)
	imageSetFromStock     func(i *Image, stockId string, size IconSize)
	imageSetPixelSize     func(i *Image, pixelSize int)
)

func (i *Image) Clear()                                          { imageClear(i) }
func (i *Image) Get(val **D.Image, mask **T.GdkBitmap)           { imageGet(i, val, mask) }
func (i *Image) GetAnimation() *D.PixbufAnimation                { return imageGetAnimation(i) }
func (i *Image) GetGicon(gicon **I.Icon, size *IconSize)         { imageGetGicon(i, gicon, size) }
func (i *Image) GetIconName(iconName **T.Gchar, size *IconSize)  { imageGetIconName(i, iconName, size) }
func (i *Image) GetIconSet(iconSet **IconSet, size *IconSize)    { imageGetIconSet(i, iconSet, size) }
func (i *Image) GetImage(gdkImage **D.Image, mask **T.GdkBitmap) { imageGetImage(i, gdkImage, mask) }
func (i *Image) GetPixbuf() *D.Pixbuf                            { return imageGetPixbuf(i) }
func (i *Image) GetPixelSize() int                               { return imageGetPixelSize(i) }
func (i *Image) GetPixmap(pixmap **D.Pixmap, mask **T.GdkBitmap) { imageGetPixmap(i, pixmap, mask) }
func (i *Image) GetStock(stockId **T.Gchar, size *IconSize)      { imageGetStock(i, stockId, size) }
func (i *Image) GetStorageType() ImageType                       { return imageGetStorageType(i) }
func (i *Image) Set(val *D.Image, mask *T.GdkBitmap)             { imageSet(i, val, mask) }
func (i *Image) SetFromAnimation(animation *D.PixbufAnimation)   { imageSetFromAnimation(i, animation) }
func (i *Image) SetFromFile(filename string)                     { imageSetFromFile(i, filename) }
func (i *Image) SetFromGicon(icon *I.Icon, size IconSize)        { imageSetFromGicon(i, icon, size) }
func (i *Image) SetFromIconName(iconName string, size IconSize) {
	imageSetFromIconName(i, iconName, size)
}
func (i *Image) SetFromIconSet(iconSet *IconSet, size IconSize) { imageSetFromIconSet(i, iconSet, size) }
func (i *Image) SetFromImage(gdkImage *D.Image, mask *T.GdkBitmap) {
	imageSetFromImage(i, gdkImage, mask)
}
func (i *Image) SetFromPixbuf(pixbuf *D.Pixbuf) { imageSetFromPixbuf(i, pixbuf) }
func (i *Image) SetFromPixmap(pixmap *D.Pixmap, mask *T.GdkBitmap) {
	imageSetFromPixmap(i, pixmap, mask)
}
func (i *Image) SetFromStock(stockId string, size IconSize) { imageSetFromStock(i, stockId, size) }
func (i *Image) SetPixelSize(pixelSize int)                 { imageSetPixelSize(i, pixelSize) }

type ImageMenuItem struct {
	MenuItem MenuItem
	Image    *Widget
}

var (
	ImageMenuItemGetType         func() O.Type
	ImageMenuItemNew             func() *Widget
	ImageMenuItemNewWithLabel    func(label string) *Widget
	ImageMenuItemNewWithMnemonic func(label string) *Widget
	ImageMenuItemNewFromStock    func(stockId string, accelGroup *AccelGroup) *Widget

	imageMenuItemGetAlwaysShowImage func(i *ImageMenuItem) bool
	imageMenuItemGetImage           func(i *ImageMenuItem) *Widget
	imageMenuItemGetUseStock        func(i *ImageMenuItem) bool
	imageMenuItemSetAccelGroup      func(i *ImageMenuItem, accelGroup *AccelGroup)
	imageMenuItemSetAlwaysShowImage func(i *ImageMenuItem, alwaysShow bool)
	imageMenuItemSetImage           func(i *ImageMenuItem, image *Widget)
	imageMenuItemSetUseStock        func(i *ImageMenuItem, useStock bool)
)

func (i *ImageMenuItem) GetAlwaysShowImage() bool { return imageMenuItemGetAlwaysShowImage(i) }
func (i *ImageMenuItem) GetImage() *Widget        { return imageMenuItemGetImage(i) }
func (i *ImageMenuItem) GetUseStock() bool        { return imageMenuItemGetUseStock(i) }
func (i *ImageMenuItem) SetAccelGroup(accelGroup *AccelGroup) {
	imageMenuItemSetAccelGroup(i, accelGroup)
}
func (i *ImageMenuItem) SetAlwaysShowImage(alwaysShow bool) {
	imageMenuItemSetAlwaysShowImage(i, alwaysShow)
}
func (i *ImageMenuItem) SetImage(image *Widget)    { imageMenuItemSetImage(i, image) }
func (i *ImageMenuItem) SetUseStock(useStock bool) { imageMenuItemSetUseStock(i, useStock) }

type ImageType Enum

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

var ImageTypeGetType func() O.Type

type IMContext struct{ parent O.Object }

var (
	ImContextGetType func() O.Type

	imContextDeleteSurrounding func(i *IMContext, offset int, nChars int) bool
	imContextFilterKeypress    func(i *IMContext, event *D.EventKey) bool
	imContextFocusIn           func(i *IMContext)
	imContextFocusOut          func(i *IMContext)
	imContextGetPreeditString  func(i *IMContext, str **T.Gchar, attrs **P.AttrList, cursorPos *int)
	imContextGetSurrounding    func(i *IMContext, text **T.Gchar, cursorIndex *int) bool
	imContextReset             func(i *IMContext)
	imContextSetClientWindow   func(i *IMContext, window *D.Window)
	imContextSetCursorLocation func(i *IMContext, area *D.Rectangle)
	imContextSetSurrounding    func(i *IMContext, text string, len int, cursorIndex int)
	imContextSetUsePreedit     func(i *IMContext, usePreedit bool)
)

func (i *IMContext) DeleteSurrounding(offset, nChars int) bool {
	return imContextDeleteSurrounding(i, offset, nChars)
}
func (i *IMContext) FilterKeypress(event *D.EventKey) bool {
	return imContextFilterKeypress(i, event)
}
func (i *IMContext) FocusIn()  { imContextFocusIn(i) }
func (i *IMContext) FocusOut() { imContextFocusOut(i) }
func (i *IMContext) GetPreeditString(str **T.Gchar, attrs **P.AttrList, cursorPos *int) {
	imContextGetPreeditString(i, str, attrs, cursorPos)
}
func (i *IMContext) GetSurrounding(text **T.Gchar, cursorIndex *int) bool {
	return imContextGetSurrounding(i, text, cursorIndex)
}
func (i *IMContext) Reset()                              { imContextReset(i) }
func (i *IMContext) SetClientWindow(window *D.Window)    { imContextSetClientWindow(i, window) }
func (i *IMContext) SetCursorLocation(area *D.Rectangle) { imContextSetCursorLocation(i, area) }
func (i *IMContext) SetSurrounding(text string, leng int, cursorIndex int) {
	imContextSetSurrounding(i, text, leng, cursorIndex)
}
func (i *IMContext) SetUsePreedit(usePreedit bool) { imContextSetUsePreedit(i, usePreedit) }

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
	ImContextSimpleGetType func() O.Type
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
	ImMulticontextGetType func() O.Type
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

var ImPreeditStyleGetType func() O.Type

var ImStatusStyleGetType func() O.Type

type InfoBar struct {
	Parent HBox
	_      *struct{}
}

var (
	InfoBarGetType        func() O.Type
	InfoBarNew            func() *Widget
	InfoBarNewWithButtons func(firstButtonText string, v ...VArg) *Widget

	infoBarAddActionWidget      func(i *InfoBar, child *Widget, responseId int)
	infoBarAddButton            func(i *InfoBar, buttonText string, responseId int) *Widget
	infoBarAddButtons           func(i *InfoBar, firstButtonText string, v ...VArg)
	infoBarGetActionArea        func(i *InfoBar) *Widget
	infoBarGetContentArea       func(i *InfoBar) *Widget
	infoBarGetMessageType       func(i *InfoBar) MessageType
	infoBarResponse             func(i *InfoBar, responseId int)
	infoBarSetDefaultResponse   func(i *InfoBar, responseId int)
	infoBarSetMessageType       func(i *InfoBar, messageType MessageType)
	infoBarSetResponseSensitive func(i *InfoBar, responseId int, setting bool)
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
func (i *InfoBar) GetActionArea() *Widget                 { return infoBarGetActionArea(i) }
func (i *InfoBar) GetContentArea() *Widget                { return infoBarGetContentArea(i) }
func (i *InfoBar) GetMessageType() MessageType            { return infoBarGetMessageType(i) }
func (i *InfoBar) Response(responseId int)                { infoBarResponse(i, responseId) }
func (i *InfoBar) SetDefaultResponse(responseId int)      { infoBarSetDefaultResponse(i, responseId) }
func (i *InfoBar) SetMessageType(messageType MessageType) { infoBarSetMessageType(i, messageType) }
func (i *InfoBar) SetResponseSensitive(responseId int, setting bool) {
	infoBarSetResponseSensitive(i, responseId, setting)
}

var (
	InputDialogGetType func() O.Type
	InputDialogNew     func() *Widget
)

type Invisible struct {
	Widget          Widget
	HasUserRefCount bool
	Screen          *D.Screen
}

var (
	InvisibleGetType      func() O.Type
	InvisibleNew          func() *Widget
	InvisibleNewForScreen func(screen *D.Screen) *Widget

	invisibleGetScreen func(i *Invisible) *D.Screen
	invisibleSetScreen func(i *Invisible, screen *D.Screen)
)

func (i *Invisible) GetScreen() *D.Screen       { return invisibleGetScreen(i) }
func (i *Invisible) SetScreen(screen *D.Screen) { invisibleSetScreen(i, screen) }

type Item struct {
	Bin Bin
}

var (
	ItemGetType func() O.Type

	itemDeselect func(i *Item)
	itemSelect   func(i *Item)
	itemToggle   func(i *Item)
)

func (i *Item) Deselect() { itemDeselect(i) }
func (i *Item) Select()   { itemSelect(i) }
func (i *Item) Toggle()   { itemToggle(i) }

type (
	ItemFactory struct {
		Object           Object
		Path             *T.Gchar
		Accel_group      *AccelGroup
		Widget           *Widget
		Items            *T.GSList
		Translate_func   TranslateFunc
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
	ItemFactoryGetType func() O.Type
	ItemFactoryNew     func(containerType O.Type, path string, accelGroup *AccelGroup) *ItemFactory

	ItemFactoriesPathDelete        func(ifactoryPath string, path string)
	ItemFactoryAddForeign          func(accelWidget *Widget, fullPath string, accelGroup *AccelGroup, keyval uint, modifiers T.GdkModifierType)
	ItemFactoryCreateMenuEntries   func(nEntries uint, entries *MenuEntry)
	ItemFactoryFromPath            func(path string) *ItemFactory
	ItemFactoryFromWidget          func(widget *Widget) *ItemFactory
	ItemFactoryPathFromWidget      func(widget *Widget) string
	ItemFactoryPopupDataFromWidget func(widget *Widget) T.Gpointer

	itemFactoryConstruct         func(i *ItemFactory, containerType O.Type, path string, accelGroup *AccelGroup)
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
	itemFactorySetTranslateFunc  func(i *ItemFactory, f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify)
)

func (i *ItemFactory) Construct(containerType O.Type, path string, accelGroup *AccelGroup) {
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
func (i *ItemFactory) SetTranslateFunc(f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify) {
	itemFactorySetTranslateFunc(i, f, data, notify)
}
