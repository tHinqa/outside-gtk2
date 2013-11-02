// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	I "github.com/tHinqa/outside-gtk2/gio"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type IconFactory struct {
	Parent O.Object
	Icons  *L.HashTable
}

var (
	IconFactoryGetType func() O.Type
	IconFactoryNew     func() *IconFactory

	IconFactoryLookupDefault func(stockId string) *IconSet

	IconFactoryAdd           func(i *IconFactory, stockId string, iconSet *IconSet)
	IconFactoryAddDefault    func(i *IconFactory)
	IconFactoryLookup        func(i *IconFactory, stockId string) *IconSet
	IconFactoryRemoveDefault func(i *IconFactory)
)

func (i *IconFactory) Add(stockId string, iconSet *IconSet) { IconFactoryAdd(i, stockId, iconSet) }
func (i *IconFactory) AddDefault()                          { IconFactoryAddDefault(i) }
func (i *IconFactory) Lookup(stockId string) *IconSet       { return IconFactoryLookup(i, stockId) }
func (i *IconFactory) RemoveDefault()                       { IconFactoryRemoveDefault(i) }

type IconInfo struct{}

var (
	IconInfoGetType      func() O.Type
	IconInfoNewForPixbuf func(iconTheme *IconTheme, pixbuf *D.Pixbuf) *IconInfo

	IconInfoCopy              func(i *IconInfo) *IconInfo
	IconInfoFree              func(i *IconInfo)
	IconInfoGetAttachPoints   func(i *IconInfo, points **D.Point, nPoints *int) bool
	IconInfoGetBaseSize       func(i *IconInfo) int
	IconInfoGetBuiltinPixbuf  func(i *IconInfo) *D.Pixbuf
	IconInfoGetDisplayName    func(i *IconInfo) string
	IconInfoGetEmbeddedRect   func(i *IconInfo, rectangle *D.Rectangle) bool
	IconInfoGetFilename       func(i *IconInfo) string
	IconInfoLoadIcon          func(i *IconInfo, error **L.Error) *D.Pixbuf
	IconInfoSetRawCoordinates func(i *IconInfo, rawCoordinates bool)
)

func (i *IconInfo) Copy() *IconInfo { return IconInfoCopy(i) }
func (i *IconInfo) Free()           { IconInfoFree(i) }
func (i *IconInfo) GetAttachPoints(points **D.Point, nPoints *int) bool {
	return IconInfoGetAttachPoints(i, points, nPoints)
}
func (i *IconInfo) GetBaseSize() int            { return IconInfoGetBaseSize(i) }
func (i *IconInfo) GetBuiltinPixbuf() *D.Pixbuf { return IconInfoGetBuiltinPixbuf(i) }
func (i *IconInfo) GetDisplayName() string      { return IconInfoGetDisplayName(i) }
func (i *IconInfo) GetEmbeddedRect(rectangle *D.Rectangle) bool {
	return IconInfoGetEmbeddedRect(i, rectangle)
}
func (i *IconInfo) GetFilename() string               { return IconInfoGetFilename(i) }
func (i *IconInfo) LoadIcon(err **L.Error) *D.Pixbuf { return IconInfoLoadIcon(i, err) }
func (i *IconInfo) SetRawCoordinates(rawCoordinates bool) {
	IconInfoSetRawCoordinates(i, rawCoordinates)
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

	IconSetAddSource  func(i *IconSet, source *IconSource)
	IconSetCopy       func(i *IconSet) *IconSet
	IconSetGetSizes   func(i *IconSet, sizes **IconSize, nSizes *int)
	IconSetRef        func(i *IconSet) *IconSet
	IconSetRenderIcon func(i *IconSet, style *Style, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *D.Pixbuf
	IconSetUnref      func(i *IconSet)
)

func (i *IconSet) AddSource(source *IconSource)           { IconSetAddSource(i, source) }
func (i *IconSet) Copy() *IconSet                         { return IconSetCopy(i) }
func (i *IconSet) GetSizes(sizes **IconSize, nSizes *int) { IconSetGetSizes(i, sizes, nSizes) }
func (i *IconSet) Ref() *IconSet                          { return IconSetRef(i) }
func (i *IconSet) RenderIcon(style *Style, direction TextDirection, state StateType, size IconSize, widget *Widget, detail string) *D.Pixbuf {
	return IconSetRenderIcon(i, style, direction, state, size, widget, detail)
}
func (i *IconSet) Unref() { IconSetUnref(i) }

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

	IconSizeGetName func(i IconSize) string
	IconSizeLookup  func(i IconSize, width *int, height *int) bool
)

func (i IconSize) GetName() string                     { return IconSizeGetName(i) }
func (i IconSize) Lookup(width *int, height *int) bool { return IconSizeLookup(i, width, height) }

type IconSource struct{}

var (
	IconSourceGetType func() O.Type
	IconSourceNew     func() *IconSource

	IconSourceCopy                   func(i *IconSource) *IconSource
	IconSourceFree                   func(i *IconSource)
	IconSourceGetDirection           func(i *IconSource) TextDirection
	IconSourceGetDirectionWildcarded func(i *IconSource) bool
	IconSourceGetFilename            func(i *IconSource) string
	IconSourceGetIconName            func(i *IconSource) string
	IconSourceGetPixbuf              func(i *IconSource) *D.Pixbuf
	IconSourceGetSize                func(i *IconSource) IconSize
	IconSourceGetSizeWildcarded      func(i *IconSource) bool
	IconSourceGetState               func(i *IconSource) StateType
	IconSourceGetStateWildcarded     func(i *IconSource) bool
	IconSourceSetDirection           func(i *IconSource, direction TextDirection)
	IconSourceSetDirectionWildcarded func(i *IconSource, setting bool)
	IconSourceSetFilename            func(i *IconSource, filename string)
	IconSourceSetIconName            func(i *IconSource, iconName string)
	IconSourceSetPixbuf              func(i *IconSource, pixbuf *D.Pixbuf)
	IconSourceSetSize                func(i *IconSource, size IconSize)
	IconSourceSetSizeWildcarded      func(i *IconSource, setting bool)
	IconSourceSetState               func(i *IconSource, state StateType)
	IconSourceSetStateWildcarded     func(i *IconSource, setting bool)
)

func (i *IconSource) Copy() *IconSource                    { return IconSourceCopy(i) }
func (i *IconSource) Free()                                { IconSourceFree(i) }
func (i *IconSource) GetDirection() TextDirection          { return IconSourceGetDirection(i) }
func (i *IconSource) GetDirectionWildcarded() bool         { return IconSourceGetDirectionWildcarded(i) }
func (i *IconSource) GetFilename() string                  { return IconSourceGetFilename(i) }
func (i *IconSource) GetIconName() string                  { return IconSourceGetIconName(i) }
func (i *IconSource) GetPixbuf() *D.Pixbuf                 { return IconSourceGetPixbuf(i) }
func (i *IconSource) GetSize() IconSize                    { return IconSourceGetSize(i) }
func (i *IconSource) GetSizeWildcarded() bool              { return IconSourceGetSizeWildcarded(i) }
func (i *IconSource) GetState() StateType                  { return IconSourceGetState(i) }
func (i *IconSource) GetStateWildcarded() bool             { return IconSourceGetStateWildcarded(i) }
func (i *IconSource) SetDirection(direction TextDirection) { IconSourceSetDirection(i, direction) }
func (i *IconSource) SetDirectionWildcarded(setting bool) {
	IconSourceSetDirectionWildcarded(i, setting)
}
func (i *IconSource) SetFilename(filename string)     { IconSourceSetFilename(i, filename) }
func (i *IconSource) SetIconName(iconName string)     { IconSourceSetIconName(i, iconName) }
func (i *IconSource) SetPixbuf(pixbuf *D.Pixbuf)      { IconSourceSetPixbuf(i, pixbuf) }
func (i *IconSource) SetSize(size IconSize)           { IconSourceSetSize(i, size) }
func (i *IconSource) SetSizeWildcarded(setting bool)  { IconSourceSetSizeWildcarded(i, setting) }
func (i *IconSource) SetState(state StateType)        { IconSourceSetState(i, state) }
func (i *IconSource) SetStateWildcarded(setting bool) { IconSourceSetStateWildcarded(i, setting) }

type IconTheme struct {
	Parent O.Object
	_      *struct{}
}

var (
	IconThemeGetType func() O.Type
	IconThemeNew     func() *IconTheme

	IconThemeErrorGetType func() O.Type
	IconThemeErrorQuark   func() L.Quark
	IconThemeGetDefault   func() *IconTheme

	IconThemeGetForScreen   func(screen *D.Screen) *IconTheme
	IconThemeAddBuiltinIcon func(iconName string, size int, pixbuf *D.Pixbuf)

	IconThemeAppendSearchPath   func(i *IconTheme, path string)
	IconThemeChooseIcon         func(i *IconTheme, iconNames []string, size int, flags IconLookupFlags) *IconInfo
	IconThemeGetExampleIconName func(i *IconTheme) string
	IconThemeGetIconSizes       func(i *IconTheme, iconName string) *int
	IconThemeGetSearchPath      func(i *IconTheme, path ***T.Gchar, nElements *int)
	IconThemeHasIcon            func(i *IconTheme, iconName string) bool
	IconThemeListContexts       func(i *IconTheme) *L.List
	IconThemeListIcons          func(i *IconTheme, context string) *L.List
	IconThemeLoadIcon           func(i *IconTheme, iconName string, size int, flags IconLookupFlags, error **L.Error) *D.Pixbuf
	IconThemeLookupByGicon      func(i *IconTheme, icon *I.Icon, size int, flags IconLookupFlags) *IconInfo
	IconThemeLookupIcon         func(i *IconTheme, iconName string, size int, flags IconLookupFlags) *IconInfo
	IconThemePrependSearchPath  func(i *IconTheme, path string)
	IconThemeRescanIfNeeded     func(i *IconTheme) bool
	IconThemeSetCustomTheme     func(i *IconTheme, themeName string)
	IconThemeSetScreen          func(i *IconTheme, screen *D.Screen)
	IconThemeSetSearchPath      func(i *IconTheme, path **T.Gchar, nElements int)
)

func (i *IconTheme) AppendSearchPath(path string) { IconThemeAppendSearchPath(i, path) }
func (i *IconTheme) ChooseIcon(iconNames []string, size int, flags IconLookupFlags) *IconInfo {
	return IconThemeChooseIcon(i, iconNames, size, flags)
}
func (i *IconTheme) GetExampleIconName() string        { return IconThemeGetExampleIconName(i) }
func (i *IconTheme) GetIconSizes(iconName string) *int { return IconThemeGetIconSizes(i, iconName) }
func (i *IconTheme) GetSearchPath(path ***T.Gchar, nElements *int) {
	IconThemeGetSearchPath(i, path, nElements)
}
func (i *IconTheme) HasIcon(iconName string) bool     { return IconThemeHasIcon(i, iconName) }
func (i *IconTheme) ListContexts() *L.List            { return IconThemeListContexts(i) }
func (i *IconTheme) ListIcons(context string) *L.List { return IconThemeListIcons(i, context) }
func (i *IconTheme) LoadIcon(iconName string, size int, flags IconLookupFlags, err **L.Error) *D.Pixbuf {
	return IconThemeLoadIcon(i, iconName, size, flags, err)
}
func (i *IconTheme) LookupByGicon(icon *I.Icon, size int, flags IconLookupFlags) *IconInfo {
	return IconThemeLookupByGicon(i, icon, size, flags)
}
func (i *IconTheme) LookupIcon(iconName string, size int, flags IconLookupFlags) *IconInfo {
	return IconThemeLookupIcon(i, iconName, size, flags)
}
func (i *IconTheme) PrependSearchPath(path string)   { IconThemePrependSearchPath(i, path) }
func (i *IconTheme) RescanIfNeeded() bool            { return IconThemeRescanIfNeeded(i) }
func (i *IconTheme) SetCustomTheme(themeName string) { IconThemeSetCustomTheme(i, themeName) }
func (i *IconTheme) SetScreen(screen *D.Screen)      { IconThemeSetScreen(i, screen) }
func (i *IconTheme) SetSearchPath(path **T.Gchar, nElements int) {
	IconThemeSetSearchPath(i, path, nElements)
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

	IconViewConvertWidgetToBinWindowCoords func(i *IconView, wx int, wy int, bx *int, by *int)
	IconViewCreateDragIcon                 func(i *IconView, path *TreePath) *D.Pixmap
	IconViewEnableModelDragDest            func(i *IconView, targets *TargetEntry, nTargets int, actions D.DragAction)
	IconViewEnableModelDragSource          func(i *IconView, startButtonMask T.GdkModifierType, targets *TargetEntry, nTargets int, actions D.DragAction)
	IconViewGetColumns                     func(i *IconView) int
	IconViewGetColumnSpacing               func(i *IconView) int
	IconViewGetCursor                      func(i *IconView, path **TreePath, cell **CellRenderer) bool
	IconViewGetDestItemAtPos               func(i *IconView, dragX, dragY int, path **TreePath, pos *IconViewDropPosition) bool
	IconViewGetDragDestItem                func(i *IconView, path **TreePath, pos *IconViewDropPosition)
	IconViewGetItemAtPos                   func(i *IconView, x, y int, path **TreePath, cell **CellRenderer) bool
	IconViewGetItemColumn                  func(i *IconView, path *TreePath) int
	IconViewGetItemOrientation             func(i *IconView) Orientation
	IconViewGetItemPadding                 func(i *IconView) int
	IconViewGetItemRow                     func(i *IconView, path *TreePath) int
	IconViewGetItemWidth                   func(i *IconView) int
	IconViewGetMargin                      func(i *IconView) int
	IconViewGetMarkupColumn                func(i *IconView) int
	IconViewGetModel                       func(i *IconView) *TreeModel
	IconViewGetOrientation                 func(i *IconView) Orientation
	IconViewGetPathAtPos                   func(i *IconView, x, y int) *TreePath
	IconViewGetPixbufColumn                func(i *IconView) int
	IconViewGetReorderable                 func(i *IconView) bool
	IconViewGetRowSpacing                  func(i *IconView) int
	IconViewGetSelectedItems               func(i *IconView) *L.List
	IconViewGetSelectionMode               func(i *IconView) SelectionMode
	IconViewGetSpacing                     func(i *IconView) int
	IconViewGetTextColumn                  func(i *IconView) int
	IconViewGetTooltipColumn               func(i *IconView) int
	IconViewGetTooltipContext              func(i *IconView, x *int, y *int, keyboardTip bool, model **TreeModel, path **TreePath, iter *TreeIter) bool
	IconViewGetVisibleRange                func(i *IconView, startPath **TreePath, endPath **TreePath) bool
	IconViewItemActivated                  func(i *IconView, path *TreePath)
	IconViewPathIsSelected                 func(i *IconView, path *TreePath) bool
	IconViewScrollToPath                   func(i *IconView, path *TreePath, useAlign bool, rowAlign, colAlign float32)
	IconViewSelectAll                      func(i *IconView)
	IconViewSelectedForeach                func(i *IconView, f IconViewForeachFunc, data T.Gpointer)
	IconViewSelectPath                     func(i *IconView, path *TreePath)
	IconViewSetColumns                     func(i *IconView, columns int)
	IconViewSetColumnSpacing               func(i *IconView, columnSpacing int)
	IconViewSetCursor                      func(i *IconView, path *TreePath, cell *CellRenderer, startEditing bool)
	IconViewSetDragDestItem                func(i *IconView, path *TreePath, pos IconViewDropPosition)
	IconViewSetItemOrientation             func(i *IconView, orientation Orientation)
	IconViewSetItemPadding                 func(i *IconView, itemPadding int)
	IconViewSetItemWidth                   func(i *IconView, itemWidth int)
	IconViewSetMargin                      func(i *IconView, margin int)
	IconViewSetMarkupColumn                func(i *IconView, column int)
	IconViewSetModel                       func(i *IconView, model *TreeModel)
	IconViewSetOrientation                 func(i *IconView, orientation Orientation)
	IconViewSetPixbufColumn                func(i *IconView, column int)
	IconViewSetReorderable                 func(i *IconView, reorderable bool)
	IconViewSetRowSpacing                  func(i *IconView, rowSpacing int)
	IconViewSetSelectionMode               func(i *IconView, mode SelectionMode)
	IconViewSetSpacing                     func(i *IconView, spacing int)
	IconViewSetTextColumn                  func(i *IconView, column int)
	IconViewSetTooltipCell                 func(i *IconView, tooltip *Tooltip, path *TreePath, cell *CellRenderer)
	IconViewSetTooltipColumn               func(i *IconView, column int)
	IconViewSetTooltipItem                 func(i *IconView, tooltip *Tooltip, path *TreePath)
	IconViewUnselectAll                    func(i *IconView)
	IconViewUnselectPath                   func(i *IconView, path *TreePath)
	IconViewUnsetModelDragDest             func(i *IconView)
	IconViewUnsetModelDragSource           func(i *IconView)
)

func (i *IconView) ConvertWidgetToBinWindowCoords(wx, wy int, bx, by *int) {
	IconViewConvertWidgetToBinWindowCoords(i, wx, wy, bx, by)
}
func (i *IconView) CreateDragIcon(path *TreePath) *D.Pixmap {
	return IconViewCreateDragIcon(i, path)
}
func (i *IconView) EnableModelDragDest(targets *TargetEntry, nTargets int, actions D.DragAction) {
	IconViewEnableModelDragDest(i, targets, nTargets, actions)
}
func (i *IconView) EnableModelDragSource(startButtonMask T.GdkModifierType, targets *TargetEntry, nTargets int, actions D.DragAction) {
	IconViewEnableModelDragSource(i, startButtonMask, targets, nTargets, actions)
}
func (i *IconView) GetColumns() int       { return IconViewGetColumns(i) }
func (i *IconView) GetColumnSpacing() int { return IconViewGetColumnSpacing(i) }
func (i *IconView) GetCursor(path **TreePath, cell **CellRenderer) bool {
	return IconViewGetCursor(i, path, cell)
}
func (i *IconView) GetDestItemAtPos(dragX, dragY int, path **TreePath, pos *IconViewDropPosition) bool {
	return IconViewGetDestItemAtPos(i, dragX, dragY, path, pos)
}
func (i *IconView) GetDragDestItem(path **TreePath, pos *IconViewDropPosition) {
	IconViewGetDragDestItem(i, path, pos)
}
func (i *IconView) GetItemAtPos(x, y int, path **TreePath, cell **CellRenderer) bool {
	return IconViewGetItemAtPos(i, x, y, path, cell)
}
func (i *IconView) GetItemColumn(path *TreePath) int { return IconViewGetItemColumn(i, path) }
func (i *IconView) GetItemOrientation() Orientation  { return IconViewGetItemOrientation(i) }
func (i *IconView) GetItemPadding() int              { return IconViewGetItemPadding(i) }
func (i *IconView) GetItemRow(path *TreePath) int    { return IconViewGetItemRow(i, path) }
func (i *IconView) GetItemWidth() int                { return IconViewGetItemWidth(i) }
func (i *IconView) GetMargin() int                   { return IconViewGetMargin(i) }
func (i *IconView) GetMarkupColumn() int             { return IconViewGetMarkupColumn(i) }
func (i *IconView) GetModel() *TreeModel             { return IconViewGetModel(i) }
func (i *IconView) GetOrientation() Orientation      { return IconViewGetOrientation(i) }
func (i *IconView) GetPathAtPos(x, y int) *TreePath  { return IconViewGetPathAtPos(i, x, y) }
func (i *IconView) GetPixbufColumn() int             { return IconViewGetPixbufColumn(i) }
func (i *IconView) GetReorderable() bool             { return IconViewGetReorderable(i) }
func (i *IconView) GetRowSpacing() int               { return IconViewGetRowSpacing(i) }
func (i *IconView) GetSelectedItems() *L.List        { return IconViewGetSelectedItems(i) }
func (i *IconView) GetSelectionMode() SelectionMode  { return IconViewGetSelectionMode(i) }
func (i *IconView) GetSpacing() int                  { return IconViewGetSpacing(i) }
func (i *IconView) GetTextColumn() int               { return IconViewGetTextColumn(i) }
func (i *IconView) GetTooltipColumn() int            { return IconViewGetTooltipColumn(i) }
func (i *IconView) GetTooltipContext(x, y *int, keyboardTip bool, model **TreeModel, path **TreePath, iter *TreeIter) bool {
	return IconViewGetTooltipContext(i, x, y, keyboardTip, model, path, iter)
}
func (i *IconView) GetVisibleRange(startPath **TreePath, endPath **TreePath) bool {
	return IconViewGetVisibleRange(i, startPath, endPath)
}
func (i *IconView) ItemActivated(path *TreePath) { IconViewItemActivated(i, path) }
func (i *IconView) PathIsSelected(path *TreePath) bool {
	return IconViewPathIsSelected(i, path)
}
func (i *IconView) ScrollToPath(path *TreePath, useAlign bool, rowAlign, colAlign float32) {
	IconViewScrollToPath(i, path, useAlign, rowAlign, colAlign)
}
func (i *IconView) SelectAll() { IconViewSelectAll(i) }
func (i *IconView) SelectedForeach(f IconViewForeachFunc, data T.Gpointer) {
	IconViewSelectedForeach(i, f, data)
}
func (i *IconView) SelectPath(path *TreePath)          { IconViewSelectPath(i, path) }
func (i *IconView) SetColumns(columns int)             { IconViewSetColumns(i, columns) }
func (i *IconView) SetColumnSpacing(columnSpacing int) { IconViewSetColumnSpacing(i, columnSpacing) }
func (i *IconView) SetCursor(path *TreePath, cell *CellRenderer, startEditing bool) {
	IconViewSetCursor(i, path, cell, startEditing)
}
func (i *IconView) SetDragDestItem(path *TreePath, pos IconViewDropPosition) {
	IconViewSetDragDestItem(i, path, pos)
}
func (i *IconView) SetItemOrientation(orientation Orientation) {
	IconViewSetItemOrientation(i, orientation)
}
func (i *IconView) SetItemPadding(itemPadding int) { IconViewSetItemPadding(i, itemPadding) }
func (i *IconView) SetItemWidth(itemWidth int)     { IconViewSetItemWidth(i, itemWidth) }
func (i *IconView) SetMargin(margin int)           { IconViewSetMargin(i, margin) }
func (i *IconView) SetMarkupColumn(column int)     { IconViewSetMarkupColumn(i, column) }
func (i *IconView) SetModel(model *TreeModel)      { IconViewSetModel(i, model) }
func (i *IconView) SetOrientation(orientation Orientation) {
	IconViewSetOrientation(i, orientation)
}
func (i *IconView) SetPixbufColumn(column int)          { IconViewSetPixbufColumn(i, column) }
func (i *IconView) SetReorderable(reorderable bool)     { IconViewSetReorderable(i, reorderable) }
func (i *IconView) SetRowSpacing(rowSpacing int)        { IconViewSetRowSpacing(i, rowSpacing) }
func (i *IconView) SetSelectionMode(mode SelectionMode) { IconViewSetSelectionMode(i, mode) }
func (i *IconView) SetSpacing(spacing int)              { IconViewSetSpacing(i, spacing) }
func (i *IconView) SetTextColumn(column int)            { IconViewSetTextColumn(i, column) }
func (i *IconView) SetTooltipCell(tooltip *Tooltip, path *TreePath, cell *CellRenderer) {
	IconViewSetTooltipCell(i, tooltip, path, cell)
}
func (i *IconView) SetTooltipColumn(column int) { IconViewSetTooltipColumn(i, column) }
func (i *IconView) SetTooltipItem(tooltip *Tooltip, path *TreePath) {
	IconViewSetTooltipItem(i, tooltip, path)
}
func (i *IconView) UnselectAll()                { IconViewUnselectAll(i) }
func (i *IconView) UnselectPath(path *TreePath) { IconViewUnselectPath(i, path) }
func (i *IconView) UnsetModelDragDest()         { IconViewUnsetModelDragDest(i) }
func (i *IconView) UnsetModelDragSource()       { IconViewUnsetModelDragSource(i) }

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

	ImageClear            func(i *Image)
	ImageGet              func(i *Image, val **D.Image, mask **T.GdkBitmap)
	ImageGetAnimation     func(i *Image) *D.PixbufAnimation
	ImageGetGicon         func(i *Image, gicon **I.Icon, size *IconSize)
	ImageGetIconName      func(i *Image, iconName **T.Gchar, size *IconSize)
	ImageGetIconSet       func(i *Image, iconSet **IconSet, size *IconSize)
	ImageGetImage         func(i *Image, gdkImage **D.Image, mask **T.GdkBitmap)
	ImageGetPixbuf        func(i *Image) *D.Pixbuf
	ImageGetPixelSize     func(i *Image) int
	ImageGetPixmap        func(i *Image, pixmap **D.Pixmap, mask **T.GdkBitmap)
	ImageGetStock         func(i *Image, stockId **T.Gchar, size *IconSize)
	ImageGetStorageType   func(i *Image) ImageType
	ImageSet              func(i *Image, val *D.Image, mask *T.GdkBitmap)
	ImageSetFromAnimation func(i *Image, animation *D.PixbufAnimation)
	ImageSetFromFile      func(i *Image, filename string)
	ImageSetFromGicon     func(i *Image, icon *I.Icon, size IconSize)
	ImageSetFromIconName  func(i *Image, iconName string, size IconSize)
	ImageSetFromIconSet   func(i *Image, iconSet *IconSet, size IconSize)
	ImageSetFromImage     func(i *Image, gdkImage *D.Image, mask *T.GdkBitmap)
	ImageSetFromPixbuf    func(i *Image, pixbuf *D.Pixbuf)
	ImageSetFromPixmap    func(i *Image, pixmap *D.Pixmap, mask *T.GdkBitmap)
	ImageSetFromStock     func(i *Image, stockId string, size IconSize)
	ImageSetPixelSize     func(i *Image, pixelSize int)
)

func (i *Image) Clear()                                          { ImageClear(i) }
func (i *Image) Get(val **D.Image, mask **T.GdkBitmap)           { ImageGet(i, val, mask) }
func (i *Image) GetAnimation() *D.PixbufAnimation                { return ImageGetAnimation(i) }
func (i *Image) GetGicon(gicon **I.Icon, size *IconSize)         { ImageGetGicon(i, gicon, size) }
func (i *Image) GetIconName(iconName **T.Gchar, size *IconSize)  { ImageGetIconName(i, iconName, size) }
func (i *Image) GetIconSet(iconSet **IconSet, size *IconSize)    { ImageGetIconSet(i, iconSet, size) }
func (i *Image) GetImage(gdkImage **D.Image, mask **T.GdkBitmap) { ImageGetImage(i, gdkImage, mask) }
func (i *Image) GetPixbuf() *D.Pixbuf                            { return ImageGetPixbuf(i) }
func (i *Image) GetPixelSize() int                               { return ImageGetPixelSize(i) }
func (i *Image) GetPixmap(pixmap **D.Pixmap, mask **T.GdkBitmap) { ImageGetPixmap(i, pixmap, mask) }
func (i *Image) GetStock(stockId **T.Gchar, size *IconSize)      { ImageGetStock(i, stockId, size) }
func (i *Image) GetStorageType() ImageType                       { return ImageGetStorageType(i) }
func (i *Image) Set(val *D.Image, mask *T.GdkBitmap)             { ImageSet(i, val, mask) }
func (i *Image) SetFromAnimation(animation *D.PixbufAnimation)   { ImageSetFromAnimation(i, animation) }
func (i *Image) SetFromFile(filename string)                     { ImageSetFromFile(i, filename) }
func (i *Image) SetFromGicon(icon *I.Icon, size IconSize)        { ImageSetFromGicon(i, icon, size) }
func (i *Image) SetFromIconName(iconName string, size IconSize) {
	ImageSetFromIconName(i, iconName, size)
}
func (i *Image) SetFromIconSet(iconSet *IconSet, size IconSize) { ImageSetFromIconSet(i, iconSet, size) }
func (i *Image) SetFromImage(gdkImage *D.Image, mask *T.GdkBitmap) {
	ImageSetFromImage(i, gdkImage, mask)
}
func (i *Image) SetFromPixbuf(pixbuf *D.Pixbuf) { ImageSetFromPixbuf(i, pixbuf) }
func (i *Image) SetFromPixmap(pixmap *D.Pixmap, mask *T.GdkBitmap) {
	ImageSetFromPixmap(i, pixmap, mask)
}
func (i *Image) SetFromStock(stockId string, size IconSize) { ImageSetFromStock(i, stockId, size) }
func (i *Image) SetPixelSize(pixelSize int)                 { ImageSetPixelSize(i, pixelSize) }

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

	ImageMenuItemGetAlwaysShowImage func(i *ImageMenuItem) bool
	ImageMenuItemGetImage           func(i *ImageMenuItem) *Widget
	ImageMenuItemGetUseStock        func(i *ImageMenuItem) bool
	ImageMenuItemSetAccelGroup      func(i *ImageMenuItem, accelGroup *AccelGroup)
	ImageMenuItemSetAlwaysShowImage func(i *ImageMenuItem, alwaysShow bool)
	ImageMenuItemSetImage           func(i *ImageMenuItem, image *Widget)
	ImageMenuItemSetUseStock        func(i *ImageMenuItem, useStock bool)
)

func (i *ImageMenuItem) GetAlwaysShowImage() bool { return ImageMenuItemGetAlwaysShowImage(i) }
func (i *ImageMenuItem) GetImage() *Widget        { return ImageMenuItemGetImage(i) }
func (i *ImageMenuItem) GetUseStock() bool        { return ImageMenuItemGetUseStock(i) }
func (i *ImageMenuItem) SetAccelGroup(accelGroup *AccelGroup) {
	ImageMenuItemSetAccelGroup(i, accelGroup)
}
func (i *ImageMenuItem) SetAlwaysShowImage(alwaysShow bool) {
	ImageMenuItemSetAlwaysShowImage(i, alwaysShow)
}
func (i *ImageMenuItem) SetImage(image *Widget)    { ImageMenuItemSetImage(i, image) }
func (i *ImageMenuItem) SetUseStock(useStock bool) { ImageMenuItemSetUseStock(i, useStock) }

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

	ImContextDeleteSurrounding func(i *IMContext, offset int, nChars int) bool
	ImContextFilterKeypress    func(i *IMContext, event *D.EventKey) bool
	ImContextFocusIn           func(i *IMContext)
	ImContextFocusOut          func(i *IMContext)
	ImContextGetPreeditString  func(i *IMContext, str **T.Gchar, attrs **P.AttrList, cursorPos *int)
	ImContextGetSurrounding    func(i *IMContext, text **T.Gchar, cursorIndex *int) bool
	ImContextReset             func(i *IMContext)
	ImContextSetClientWindow   func(i *IMContext, window *D.Window)
	ImContextSetCursorLocation func(i *IMContext, area *D.Rectangle)
	ImContextSetSurrounding    func(i *IMContext, text string, len int, cursorIndex int)
	ImContextSetUsePreedit     func(i *IMContext, usePreedit bool)
)

func (i *IMContext) DeleteSurrounding(offset, nChars int) bool {
	return ImContextDeleteSurrounding(i, offset, nChars)
}
func (i *IMContext) FilterKeypress(event *D.EventKey) bool {
	return ImContextFilterKeypress(i, event)
}
func (i *IMContext) FocusIn()  { ImContextFocusIn(i) }
func (i *IMContext) FocusOut() { ImContextFocusOut(i) }
func (i *IMContext) GetPreeditString(str **T.Gchar, attrs **P.AttrList, cursorPos *int) {
	ImContextGetPreeditString(i, str, attrs, cursorPos)
}
func (i *IMContext) GetSurrounding(text **T.Gchar, cursorIndex *int) bool {
	return ImContextGetSurrounding(i, text, cursorIndex)
}
func (i *IMContext) Reset()                              { ImContextReset(i) }
func (i *IMContext) SetClientWindow(window *D.Window)    { ImContextSetClientWindow(i, window) }
func (i *IMContext) SetCursorLocation(area *D.Rectangle) { ImContextSetCursorLocation(i, area) }
func (i *IMContext) SetSurrounding(text string, leng int, cursorIndex int) {
	ImContextSetSurrounding(i, text, leng, cursorIndex)
}
func (i *IMContext) SetUsePreedit(usePreedit bool) { ImContextSetUsePreedit(i, usePreedit) }

type IMContextSimple struct {
	Object            IMContext
	Tables            *L.SList
	ComposeBuffer     [7 + 1]uint //TODO(t):Symbolic
	TentativeMatch    L.Unichar
	TentativeMatchLen int
	Bits              uint
	// InHexSequence : 1
	// ModifiersDropped : 1
}

var (
	ImContextSimpleGetType func() O.Type
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
	ImMulticontextGetType func() O.Type
	ImMulticontextNew     func() *IMContext

	ImMulticontextAppendMenuitems func(i *IMMulticontext, menushell *MenuShell)
	ImMulticontextGetContextId    func(i *IMMulticontext) string
	ImMulticontextSetContextId    func(i *IMMulticontext, contextId string)
)

func (i *IMMulticontext) AppendMenuitems(menushell *MenuShell) {
	ImMulticontextAppendMenuitems(i, menushell)
}
func (i *IMMulticontext) GetContextId() string          { return ImMulticontextGetContextId(i) }
func (i *IMMulticontext) SetContextId(contextId string) { ImMulticontextSetContextId(i, contextId) }

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

	InfoBarAddActionWidget      func(i *InfoBar, child *Widget, responseId int)
	InfoBarAddButton            func(i *InfoBar, buttonText string, responseId int) *Widget
	InfoBarAddButtons           func(i *InfoBar, firstButtonText string, v ...VArg)
	InfoBarGetActionArea        func(i *InfoBar) *Widget
	InfoBarGetContentArea       func(i *InfoBar) *Widget
	InfoBarGetMessageType       func(i *InfoBar) MessageType
	InfoBarResponse             func(i *InfoBar, responseId int)
	InfoBarSetDefaultResponse   func(i *InfoBar, responseId int)
	InfoBarSetMessageType       func(i *InfoBar, messageType MessageType)
	InfoBarSetResponseSensitive func(i *InfoBar, responseId int, setting bool)
)

func (i *InfoBar) AddActionWidget(child *Widget, responseId int) {
	InfoBarAddActionWidget(i, child, responseId)
}
func (i *InfoBar) AddButton(buttonText string, responseId int) *Widget {
	return InfoBarAddButton(i, buttonText, responseId)
}
func (i *InfoBar) AddButtons(firstButtonText string, v ...VArg) {
	InfoBarAddButtons(i, firstButtonText, v)
}
func (i *InfoBar) GetActionArea() *Widget                 { return InfoBarGetActionArea(i) }
func (i *InfoBar) GetContentArea() *Widget                { return InfoBarGetContentArea(i) }
func (i *InfoBar) GetMessageType() MessageType            { return InfoBarGetMessageType(i) }
func (i *InfoBar) Response(responseId int)                { InfoBarResponse(i, responseId) }
func (i *InfoBar) SetDefaultResponse(responseId int)      { InfoBarSetDefaultResponse(i, responseId) }
func (i *InfoBar) SetMessageType(messageType MessageType) { InfoBarSetMessageType(i, messageType) }
func (i *InfoBar) SetResponseSensitive(responseId int, setting bool) {
	InfoBarSetResponseSensitive(i, responseId, setting)
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

	InvisibleGetScreen func(i *Invisible) *D.Screen
	InvisibleSetScreen func(i *Invisible, screen *D.Screen)
)

func (i *Invisible) GetScreen() *D.Screen       { return InvisibleGetScreen(i) }
func (i *Invisible) SetScreen(screen *D.Screen) { InvisibleSetScreen(i, screen) }

type Item struct {
	Bin Bin
}

var (
	ItemGetType func() O.Type

	ItemDeselect func(i *Item)
	ItemSelect   func(i *Item)
	ItemToggle   func(i *Item)
)

func (i *Item) Deselect() { ItemDeselect(i) }
func (i *Item) Select()   { ItemSelect(i) }
func (i *Item) Toggle()   { ItemToggle(i) }

type (
	ItemFactory struct {
		Object           Object
		Path             *T.Gchar
		Accel_group      *AccelGroup
		Widget           *Widget
		Items            *L.SList
		Translate_func   TranslateFunc
		Translate_data   T.Gpointer
		Translate_notify O.DestroyNotify
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

	ItemFactoryConstruct         func(i *ItemFactory, containerType O.Type, path string, accelGroup *AccelGroup)
	ItemFactoryCreateItem        func(i *ItemFactory, entry *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint)
	ItemFactoryCreateItems       func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer)
	ItemFactoryCreateItemsAc     func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint)
	ItemFactoryDeleteEntries     func(i *ItemFactory, nEntries uint, entries *ItemFactoryEntry)
	ItemFactoryDeleteEntry       func(i *ItemFactory, entry *ItemFactoryEntry)
	ItemFactoryDeleteItem        func(i *ItemFactory, path string)
	ItemFactoryGetItem           func(i *ItemFactory, path string) *Widget
	ItemFactoryGetItemByAction   func(i *ItemFactory, action uint) *Widget
	ItemFactoryGetWidget         func(i *ItemFactory, path string) *Widget
	ItemFactoryGetWidgetByAction func(i *ItemFactory, action uint) *Widget
	ItemFactoryPopup             func(i *ItemFactory, x uint, y uint, mouseButton uint, time T.GUint32)
	ItemFactoryPopupData         func(i *ItemFactory) T.Gpointer
	ItemFactoryPopupWithData     func(i *ItemFactory, popupData T.Gpointer, destroy O.DestroyNotify, x uint, y uint, mouseButton uint, time T.GUint32)
	ItemFactorySetTranslateFunc  func(i *ItemFactory, f TranslateFunc, data T.Gpointer, notify O.DestroyNotify)
)

func (i *ItemFactory) Construct(containerType O.Type, path string, accelGroup *AccelGroup) {
	ItemFactoryConstruct(i, containerType, path, accelGroup)
}
func (i *ItemFactory) CreateItem(entry *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint) {
	ItemFactoryCreateItem(i, entry, callbackData, callbackType)
}
func (i *ItemFactory) CreateItems(nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer) {
	ItemFactoryCreateItems(i, nEntries, entries, callbackData)
}
func (i *ItemFactory) CreateItemsAc(nEntries uint, entries *ItemFactoryEntry, callbackData T.Gpointer, callbackType uint) {
	ItemFactoryCreateItemsAc(i, nEntries, entries, callbackData, callbackType)
}
func (i *ItemFactory) DeleteEntries(nEntries uint, entries *ItemFactoryEntry) {
	ItemFactoryDeleteEntries(i, nEntries, entries)
}
func (i *ItemFactory) DeleteEntry(entry *ItemFactoryEntry) { ItemFactoryDeleteEntry(i, entry) }
func (i *ItemFactory) DeleteItem(path string)              { ItemFactoryDeleteItem(i, path) }
func (i *ItemFactory) GetItem(path string) *Widget         { return ItemFactoryGetItem(i, path) }
func (i *ItemFactory) GetItemByAction(action uint) *Widget {
	return ItemFactoryGetItemByAction(i, action)
}
func (i *ItemFactory) GetWidget(path string) *Widget { return ItemFactoryGetWidget(i, path) }
func (i *ItemFactory) GetWidgetByAction(action uint) *Widget {
	return ItemFactoryGetWidgetByAction(i, action)
}
func (i *ItemFactory) Popup(x, y, mouseButton uint, time T.GUint32) {
	ItemFactoryPopup(i, x, y, mouseButton, time)
}
func (i *ItemFactory) PopupData() T.Gpointer { return ItemFactoryPopupData(i) }
func (i *ItemFactory) PopupWithData(popupData T.Gpointer, destroy O.DestroyNotify, x, y, mouseButton uint, time T.GUint32) {
	ItemFactoryPopupWithData(i, popupData, destroy, x, y, mouseButton, time)
}
func (i *ItemFactory) SetTranslateFunc(f TranslateFunc, data T.Gpointer, notify O.DestroyNotify) {
	ItemFactorySetTranslateFunc(i, f, data, notify)
}
