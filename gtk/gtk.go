// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gtk provides API definitions for accessing
//libgtk-win32-2.0-0.dll.
package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type simpleObject struct{ parent T.GObject }

var (
	AlignmentGetType              func() T.GType
	AnchorTypeGetType             func() T.GType
	ArgFlagsGetType               func() T.GType
	AttachOptionsGetType          func() T.GType
	ContainerGetType              func() T.GType
	CornerTypeGetType             func() T.GType
	DebugFlagGetType              func() T.GType
	DeleteTypeGetType             func() T.GType
	DestDefaultsGetType           func() T.GType
	DirectionTypeGetType          func() T.GType
	DragResultGetType             func() T.GType
	GammaCurveGetType             func() T.GType
	GammaCurveNew                 func() *Widget
	HbuttonBoxGetType             func() T.GType
	HbuttonBoxNew                 func() *Widget
	HbuttonBoxGetSpacingDefault   func() int
	HbuttonBoxGetLayoutDefault    func() ButtonBoxStyle
	HbuttonBoxSetSpacingDefault   func(spacing int)
	HbuttonBoxSetLayoutDefault    func(layout ButtonBoxStyle)
	HpanedGetType                 func() T.GType
	HpanedNew                     func() *Widget
	HrulerGetType                 func() T.GType
	HrulerNew                     func() *Widget
	HscaleGetType                 func() T.GType
	HscaleNew                     func(adjustment *Adjustment) *Widget
	HscaleNewWithRange            func(min, max, step float64) *Widget
	HscrollbarGetType             func() T.GType
	HscrollbarNew                 func(adjustment *Adjustment) *Widget
	HseparatorGetType             func() T.GType
	HseparatorNew                 func() *Widget
	IconLookupFlagsGetType        func() T.GType
	IdentifierGetType             func() T.GType
	ImageTypeGetType              func() T.GType
	ImPreeditStyleGetType         func() T.GType
	ImStatusStyleGetType          func() T.GType
	InputDialogGetType            func() T.GType
	InputDialogNew                func() *Widget
	JustificationGetType          func() T.GType
	MatchTypeGetType              func() T.GType
	MenuBarGetType                func() T.GType
	MenuBarNew                    func() *Widget
	MenuDirectionTypeGetType      func() T.GType
	MenuGetType                   func() T.GType
	MenuItemGetType               func() T.GType
	MenuItemNew                   func() *Widget
	MenuNew                       func() *Widget
	MenuShellGetType              func() T.GType
	MenuToolButtonGetType         func() T.GType
	MessageDialogGetType          func() T.GType
	MessageTypeGetType            func() T.GType
	MetricTypeGetType             func() T.GType
	MiscGetType                   func() T.GType
	MountOperationGetType         func() T.GType
	MovementStepGetType           func() T.GType
	NotebookGetType               func() T.GType
	NotebookNew                   func() *Widget
	NotebookTabGetType            func() T.GType
	NumberUpLayoutGetType         func() T.GType
	ObjectFlagsGetType            func() T.GType
	ObjectGetType                 func() T.GType
	OffscreenWindowGetType        func() T.GType
	OffscreenWindowNew            func() *Widget
	OldEditableGetType            func() T.GType
	OptionMenuGetType             func() T.GType
	OptionMenuNew                 func() *Widget
	OrientableGetType             func() T.GType
	OrientationGetType            func() T.GType
	PackDirectionGetType          func() T.GType
	PackTypeGetType               func() T.GType
	PageOrientationGetType        func() T.GType
	PageSetGetType                func() T.GType
	PageSetupGetType              func() T.GType
	PanedGetType                  func() T.GType
	PaperSizeGetType              func() T.GType
	PathPriorityTypeGetType       func() T.GType
	PathTypeGetType               func() T.GType
	PixmapGetType                 func() T.GType
	PlugGetType                   func() T.GType
	PolicyTypeGetType             func() T.GType
	PositionTypeGetType           func() T.GType
	PreviewGetType                func() T.GType
	PreviewTypeGetType            func() T.GType
	PrintContextGetType           func() T.GType
	PrintDuplexGetType            func() T.GType
	PrintErrorGetType             func() T.GType
	PrintOperationActionGetType   func() T.GType
	PrintOperationGetType         func() T.GType
	PrintOperationPreviewGetType  func() T.GType
	PrintOperationResultGetType   func() T.GType
	PrintPagesGetType             func() T.GType
	PrintQualityGetType           func() T.GType
	PrintSettingsGetType          func() T.GType
	PrintStatusGetType            func() T.GType
	PrivateFlagsGetType           func() T.GType
	ProgressBarGetType            func() T.GType
	ProgressBarNew                func() *Widget
	ProgressBarOrientationGetType func() T.GType
	ProgressBarStyleGetType       func() T.GType
	ProgressGetType               func() T.GType
	RadioActionGetType            func() T.GType
	RadioButtonGetType            func() T.GType
	RadioMenuItemGetType          func() T.GType
	RadioToolButtonGetType        func() T.GType
	RangeGetType                  func() T.GType
	RcFlagsGetType                func() T.GType
	RcStyleGetType                func() T.GType
	RcTokenTypeGetType            func() T.GType
	RecentActionGetType           func() T.GType
	RecentChooserDialogGetType    func() T.GType
	RecentChooserErrorGetType     func() T.GType
	RecentChooserGetType          func() T.GType
	RecentChooserMenuGetType      func() T.GType
	RecentChooserMenuNew          func() *Widget
	RecentChooserWidgetGetType    func() T.GType
	RecentChooserWidgetNew        func() *Widget
	RecentFilterFlagsGetType      func() T.GType
	RecentFilterGetType           func() T.GType
	RecentInfoGetType             func() T.GType
	RecentManagerErrorGetType     func() T.GType
	RecentManagerGetType          func() T.GType
	RecentSortTypeGetType         func() T.GType
	ReliefStyleGetType            func() T.GType
	RequisitionGetType            func() T.GType
	ResizeModeGetType             func() T.GType
	ResponseTypeGetType           func() T.GType
	RulerGetType                  func() T.GType
	ScaleButtonGetType            func() T.GType
	ScaleGetType                  func() T.GType
	ScrollbarGetType              func() T.GType
	ScrolledWindowGetType         func() T.GType
	ScrollStepGetType             func() T.GType
	ScrollTypeGetType             func() T.GType
	SelectionDataGetType          func() T.GType
	SelectionModeGetType          func() T.GType
	SensitivityTypeGetType        func() T.GType
	SeparatorGetType              func() T.GType
	SeparatorMenuItemGetType      func() T.GType
	SeparatorMenuItemNew          func() *Widget
	SeparatorToolItemGetType      func() T.GType
	SettingsGetType               func() T.GType
	ShadowTypeGetType             func() T.GType
	SideTypeGetType               func() T.GType
	SignalRunTypeGetType          func() T.GType
	SizeGroupGetType              func() T.GType
	SizeGroupModeGetType          func() T.GType
	SocketGetType                 func() T.GType
	SocketNew                     func() *Widget
	SortTypeGetType               func() T.GType
	SpinButtonGetType             func() T.GType
	SpinButtonUpdatePolicyGetType func() T.GType
	SpinnerGetType                func() T.GType
	SpinnerNew                    func() *Widget
	SpinTypeGetType               func() T.GType
	StateTypeGetType              func() T.GType
	StatusbarGetType              func() T.GType
	StatusbarNew                  func() *Widget
	StatusIconGetType             func() T.GType
	StyleGetType                  func() T.GType
	SubmenuDirectionGetType       func() T.GType
	SubmenuPlacementGetType       func() T.GType
	TableGetType                  func() T.GType
	TargetFlagsGetType            func() T.GType
	TargetListGetType             func() T.GType
	TearoffMenuItemGetType        func() T.GType
	TearoffMenuItemNew            func() *Widget
	TextAttributesGetType         func() T.GType
	TextBufferGetType             func() T.GType
	TextBufferTargetInfoGetType   func() T.GType
	TextChildAnchorGetType        func() T.GType
	TextDirectionGetType          func() T.GType
	TextIterGetType               func() T.GType
	TextMarkGetType               func() T.GType
	TextSearchFlagsGetType        func() T.GType
	TextTagGetType                func() T.GType
	TextTagTableGetType           func() T.GType
	TextViewGetType               func() T.GType
	TextViewNew                   func() *Widget
	TextWindowTypeGetType         func() T.GType
	TipsQueryGetType              func() T.GType
	TipsQueryNew                  func() *Widget
	ToggleActionGetType           func() T.GType
	ToggleButtonGetType           func() T.GType
	ToggleButtonNew               func() *Widget
	ToggleToolButtonGetType       func() T.GType
	ToolbarChildTypeGetType       func() T.GType
	ToolbarGetType                func() T.GType
	ToolbarNew                    func() *Widget
	ToolbarSpaceStyleGetType      func() T.GType
	ToolbarStyleGetType           func() T.GType
	ToolButtonGetType             func() T.GType
	ToolItemGetType               func() T.GType
	ToolItemGroupGetType          func() T.GType
	ToolPaletteDragTargetsGetType func() T.GType
	ToolPaletteGetType            func() T.GType
	ToolPaletteNew                func() *Widget
	ToolShellGetType              func() T.GType
	TooltipGetType                func() T.GType
	TooltipsGetType               func() T.GType
	TreeDragDestGetType           func() T.GType
	TreeDragSourceGetType         func() T.GType
	TreeGetType                   func() T.GType
	TreeItemGetType               func() T.GType
	TreeItemNew                   func() *Widget
	TreeIterGetType               func() T.GType
	TreeModelFilterGetType        func() T.GType
	TreeModelFlagsGetType         func() T.GType
	TreeModelGetType              func() T.GType
	TreeModelSortGetType          func() T.GType
	TreeNew                       func() *Widget
	TreePathGetType               func() T.GType
	TreeRowReferenceGetType       func() T.GType
	TreeSelectionGetType          func() T.GType
	TreeSortableGetType           func() T.GType
	TreeStoreGetType              func() T.GType
	TreeViewColumnGetType         func() T.GType
	TreeViewColumnSizingGetType   func() T.GType
	TreeViewDropPositionGetType   func() T.GType
	TreeViewGetType               func() T.GType
	TreeViewGridLinesGetType      func() T.GType
	TreeViewModeGetType           func() T.GType
	TreeViewNew                   func() *Widget
	UiManagerGetType              func() T.GType
	UiManagerItemTypeGetType      func() T.GType
	UnitGetType                   func() T.GType
	UpdateTypeGetType             func() T.GType
	VboxGetType                   func() T.GType
	VbuttonBoxGetType             func() T.GType
	VbuttonBoxNew                 func() *Widget
	ViewportGetType               func() T.GType
	VisibilityGetType             func() T.GType
	VolumeButtonGetType           func() T.GType
	VolumeButtonNew               func() *Widget
	VpanedGetType                 func() T.GType
	VpanedNew                     func() *Widget
	VrulerGetType                 func() T.GType
	VrulerNew                     func() *Widget
	VscaleGetType                 func() T.GType
	VscrollbarGetType             func() T.GType
	VseparatorGetType             func() T.GType
	VseparatorNew                 func() *Widget
	WidgetFlagsGetType            func() T.GType
	WidgetHelpTypeGetType         func() T.GType
	WindowGroupGetType            func() T.GType
	WindowPositionGetType         func() T.GType
	WindowTypeGetType             func() T.GType
	WrapModeGetType               func() T.GType

	AlternativeDialogButtonOrder func(
		screen *T.GdkScreen) T.Gboolean

	AcceleratorValid func(
		keyval uint,
		modifiers T.GdkModifierType) T.Gboolean

	AcceleratorParse func(
		accelerator string,
		acceleratorKey *uint,
		acceleratorMods *T.GdkModifierType)

	AcceleratorName func(
		acceleratorKey uint,
		acceleratorMods T.GdkModifierType) string

	AcceleratorGetLabel func(
		acceleratorKey uint,
		acceleratorMods T.GdkModifierType) string

	AcceleratorSetDefaultModMask func(
		defaultModMask T.GdkModifierType)

	AcceleratorGetDefaultModMask func() uint

	TypeInit func(
		debugFlags T.GTypeDebugFlags)

	TypeUnique func(
		parentType T.GtkType,
		gtkinfo *T.GtkTypeInfo) T.GtkType

	TypeClass func(
		t T.GtkType) T.Gpointer

	TypeNew func(
		t T.GtkType) T.Gpointer

	ObjectSink func(
		object *T.GtkObject)

	ObjectDestroy func(
		object *T.GtkObject)

	ObjectNew func(t T.GType, firstPropertyName string,
		v ...VArg) *T.GtkObject

	ObjectRef func(
		object *T.GtkObject) *T.GtkObject

	ObjectUnref func(
		object *T.GtkObject)

	ObjectWeakref func(
		object *T.GtkObject,
		notify T.GDestroyNotify,
		data T.Gpointer)

	ObjectWeakunref func(
		object *T.GtkObject,
		notify T.GDestroyNotify,
		data T.Gpointer)

	ObjectSetData func(
		object *T.GtkObject,
		key string,
		data T.Gpointer)

	ObjectSetDataFull func(
		object *T.GtkObject,
		key string,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	ObjectRemoveData func(
		object *T.GtkObject,
		key string)

	ObjectGetData func(
		object *T.GtkObject,
		key string) T.Gpointer

	ObjectRemoveNoNotify func(
		object *T.GtkObject,
		key string)

	ObjectSetUserData func(
		object *T.GtkObject,
		data T.Gpointer)

	ObjectGetUserData func(
		object *T.GtkObject) T.Gpointer

	ObjectSetDataById func(
		object *T.GtkObject,
		dataId T.GQuark,
		data T.Gpointer)

	ObjectSetDataByIdFull func(
		object *T.GtkObject,
		dataId T.GQuark,
		dataGpointer,
		destroy T.GDestroyNotify)

	ObjectGetDataById func(
		object *T.GtkObject,
		dataId T.GQuark) T.Gpointer

	ObjectRemoveDataById func(
		object *T.GtkObject,
		dataId T.GQuark)

	ObjectRemoveNoNotifyById func(
		object *T.GtkObject,
		keyId T.GQuark)

	ObjectGet func(object *T.GtkObject,
		firstPropertyName string, v ...VArg)

	ObjectSet func(object *T.GtkObject,
		firstPropertyName string, v ...VArg)

	ObjectAddArgType func(
		argName string,
		argType T.GType,
		argFlags uint,
		argId uint)

	StyleNew func() *T.GtkStyle

	StyleCopy func(
		style *T.GtkStyle) *T.GtkStyle

	StyleAttach func(
		style *T.GtkStyle,
		window *T.GdkWindow) *T.GtkStyle

	StyleDetach func(
		style *T.GtkStyle)

	StyleRef func(
		style *T.GtkStyle) *T.GtkStyle

	StyleUnref func(
		style *T.GtkStyle)

	StyleGetFont func(
		style *T.GtkStyle) *T.GdkFont

	StyleSetFont func(
		style *T.GtkStyle,
		font *T.GdkFont)

	StyleSetBackground func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType)

	StyleApplyDefaultBackground func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		setBg T.Gboolean,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		x int,
		y int,
		width int,
		height int)

	StyleLookupIconSet func(
		style *T.GtkStyle,
		stockId string) *IconSet

	StyleLookupColor func(
		style *T.GtkStyle,
		colorName string,
		color *T.GdkColor) T.Gboolean

	StyleRenderIcon func(
		style *T.GtkStyle,
		source *IconSource,
		direction T.GtkTextDirection,
		state T.GtkStateType,
		size IconSize,
		widget *Widget,
		detail string) *T.GdkPixbuf

	DrawHline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		x1 int,
		x2 int,
		y int)

	DrawVline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		y1 int,
		y2 int,
		x int)

	DrawShadow func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int)

	DrawPolygon func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		points *T.GdkPoint,
		npoints int,
		fill T.Gboolean)

	DrawArrow func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		arrowType T.GtkArrowType,
		fill T.Gboolean,
		x int,
		y int,
		width int,
		height int)

	DrawDiamond func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int)

	DrawBox func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int)

	DrawFlatBox func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int)

	DrawCheck func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int)

	DrawOption func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int)

	DrawTab func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int)

	DrawShadowGap func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int,
		gapSide T.GtkPositionType,
		gapX int,
		gapWidth int)

	DrawBoxGap func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int,
		gapSide T.GtkPositionType,
		gapX int,
		gapWidth int)

	DrawExtension func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int,
		gapSide T.GtkPositionType)

	DrawFocus func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		x int,
		y int,
		width int,
		height int)

	DrawSlider func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int,
		orientation T.GtkOrientation)

	DrawHandle func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		x int,
		y int,
		width int,
		height int,
		orientation T.GtkOrientation)

	DrawExpander func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		x int,
		y int,
		expanderStyle ExpanderStyle)

	DrawLayout func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		useText T.Gboolean,
		x int,
		y int,
		layout *T.PangoLayout)

	DrawResizeGrip func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		edge T.GdkWindowEdge,
		x int,
		y int,
		width int,
		height int)

	PaintHline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x1 int,
		x2 int,
		y int)

	PaintVline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		y1 int,
		y2 int,
		x int)

	PaintShadow func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int)

	PaintPolygon func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		points *T.GdkPoint,
		nPoints int,
		fill T.Gboolean)

	PaintArrow func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		arrowType T.GtkArrowType,
		fill T.Gboolean,
		x int,
		y int,
		width int,
		height int)

	PaintDiamond func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int)

	PaintBox func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int)

	PaintFlatBox func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int)

	PaintCheck func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int)

	PaintOption func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int)

	PaintTab func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int)

	PaintShadowGap func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int,
		gapSide T.GtkPositionType,
		gapX int,
		gapWidth int)

	PaintBoxGap func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int,
		gapSide T.GtkPositionType,
		gapX int,
		gapWidth int)

	PaintExtension func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int,
		gapSide T.GtkPositionType)

	PaintFocus func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int)

	PaintSlider func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int,
		orientation T.GtkOrientation)

	PaintHandle func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		shadowType T.GtkShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		width int,
		height int,
		orientation T.GtkOrientation)

	PaintExpander func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		expanderStyle ExpanderStyle)

	PaintLayout func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		useText T.Gboolean,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		layout *T.PangoLayout)

	PaintResizeGrip func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		edge T.GdkWindowEdge,
		x int,
		y int,
		width int,
		height int)

	PaintSpinner func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		step uint,
		x int,
		y int,
		width int,
		height int)

	StyleGetStyleProperty func(
		style *T.GtkStyle,
		widgetType T.GType,
		propertyName string,
		value *T.GValue)

	StyleGetValist func(
		style *T.GtkStyle,
		widgetType T.GType,
		firstPropertyName string,
		varArgs T.VaList)

	StyleGet func(style *T.GtkStyle, widgetType T.GType,
		firstPropertyName string, v ...VArg)

	DrawString func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		x int,
		y int,
		string string)

	PaintString func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x int,
		y int,
		string string)

	DrawInsertionCursor func(
		widget *Widget,
		drawable *T.GdkDrawable,
		area *T.GdkRectangle,
		location *T.GdkRectangle,
		isPrimary T.Gboolean,
		direction T.GtkTextDirection,
		drawArrow T.Gboolean)

	RcAddDefaultFile func(
		filename string)

	RcSetDefaultFiles func(
		filenames **T.Gchar)

	RcGetDefaultFiles func() **T.Gchar

	RcGetStyle func(
		widget *Widget) *T.GtkStyle

	RcGetStyleByPaths func(
		settings *T.GtkSettings,
		widgetPath string,
		classPath string,
		t T.GType) *T.GtkStyle

	RcReparseAllForSettings func(
		settings *T.GtkSettings,
		forceLoad T.Gboolean) T.Gboolean

	RcResetStyles func(
		settings *T.GtkSettings)

	RcFindPixmapInPath func(
		settings *T.GtkSettings,
		scanner *T.GScanner,
		pixmapFile string) string

	RcParse func(
		filename string)

	RcParseString func(
		rcString string)

	RcReparseAll func() T.Gboolean

	RcAddWidgetNameStyle func(
		rcStyle *T.GtkRcStyle,
		pattern string)

	RcAddWidgetClassStyle func(
		rcStyle *T.GtkRcStyle,
		pattern string)

	RcAddClassStyle func(
		rcStyle *T.GtkRcStyle,
		pattern string)

	RcStyleNew func() *T.GtkRcStyle

	RcStyleCopy func(
		orig *T.GtkRcStyle) *T.GtkRcStyle

	RcStyleRef func(
		rcStyle *T.GtkRcStyle)

	RcStyleUnref func(
		rcStyle *T.GtkRcStyle)

	RcFindModuleInPath func(
		moduleFile string) string

	RcGetThemeDir func() string

	RcGetModuleDir func() string

	RcGetImModulePath func() string

	RcGetImModuleFile func() string

	RcScannerNew func() *T.GScanner

	RcParseColor func(
		scanner *T.GScanner,
		color *T.GdkColor) uint

	RcParseColorFull func(
		scanner *T.GScanner,
		style *T.GtkRcStyle,
		color *T.GdkColor) uint

	RcParseState func(
		scanner *T.GScanner,
		state *T.GtkStateType) uint

	RcParsePriority func(
		scanner *T.GScanner,
		priority *T.GtkPathPriorityType) uint

	SettingsGetDefault func() *T.GtkSettings

	SettingsGetForScreen func(
		screen *T.GdkScreen) *T.GtkSettings

	SettingsInstallProperty func(
		pspec *T.GParamSpec)

	SettingsInstallPropertyParser func(
		pspec *T.GParamSpec,
		parser T.GtkRcPropertyParser)

	RcPropertyParseColor func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		propertyValue *T.GValue) T.Gboolean

	RcPropertyParseEnum func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		propertyValue *T.GValue) T.Gboolean

	RcPropertyParseFlags func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		propertyValue *T.GValue) T.Gboolean

	RcPropertyParseRequisition func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		propertyValue *T.GValue) T.Gboolean

	RcPropertyParseBorder func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		propertyValue *T.GValue) T.Gboolean

	SettingsSetPropertyValue func(
		settings *T.GtkSettings,
		name string,
		svalue *T.GtkSettingsValue)

	SettingsSetStringProperty func(
		settings *T.GtkSettings,
		name string,
		vString string,
		origin string)

	SettingsSetLongProperty func(
		settings *T.GtkSettings,
		name string,
		vLong T.Glong,
		origin string)

	SettingsSetDoubleProperty func(
		settings *T.GtkSettings,
		name string,
		vDouble float64,
		origin string)

	WidgetPushColormap func(
		cmap *T.GdkColormap)

	WidgetPushCompositeChild func()

	WidgetPopCompositeChild func()

	WidgetPopColormap func()

	WidgetClassInstallStyleProperty func(
		klass *T.GtkWidgetClass,
		pspec *T.GParamSpec)

	WidgetClassInstallStylePropertyParser func(
		klass *T.GtkWidgetClass,
		pspec *T.GParamSpec,
		parser T.GtkRcPropertyParser)

	WidgetClassFindStyleProperty func(
		klass *T.GtkWidgetClass,
		propertyName string) *T.GParamSpec

	WidgetClassListStyleProperties func(
		klass *T.GtkWidgetClass,
		nProperties *uint) **T.GParamSpec

	WidgetStyleGetProperty func(
		widget *Widget,
		propertyName string,
		value *T.GValue)

	WidgetStyleGetValist func(
		widget *Widget,
		firstPropertyName string,
		varArgs T.VaList)

	WidgetStyleGet func(widget *Widget,
		firstPropertyName string, v ...VArg)

	WidgetSetDefaultColormap func(
		colormap *T.GdkColormap)

	WidgetGetDefaultStyle func() *T.GtkStyle

	WidgetGetDefaultColormap func() *T.GdkColormap

	WidgetGetDefaultVisual func() *T.GdkVisual

	WidgetSetDirection func(
		widget *Widget,
		dir T.GtkTextDirection)

	WidgetGetDirection func(
		widget *Widget) T.GtkTextDirection

	WidgetSetDefaultDirection func(
		dir T.GtkTextDirection)

	WidgetGetDefaultDirection func() T.GtkTextDirection

	WidgetIsComposited func(
		widget *Widget) T.Gboolean

	WidgetShapeCombineMask func(
		widget *Widget,
		shapeMask *T.GdkBitmap,
		offsetX int,
		offsetY int)

	WidgetInputShapeCombineMask func(
		widget *Widget,
		shapeMask *T.GdkBitmap,
		offsetX int,
		offsetY int)

	WidgetResetShapes func(
		widget *Widget)

	WidgetPath func(
		widget *Widget,
		pathLength *uint,
		path **T.Gchar,
		pathReversed **T.Gchar)

	WidgetClassPath func(
		widget *Widget,
		pathLength *uint,
		path **T.Gchar,
		pathReversed **T.Gchar)

	WidgetListMnemonicLabels func(
		widget *Widget) *T.GList

	WidgetAddMnemonicLabel func(
		widget *Widget,
		label *Widget)

	WidgetRemoveMnemonicLabel func(
		widget *Widget,
		label *Widget)

	WidgetSetTooltipWindow func(
		widget *Widget,
		customWindow *T.GtkWindow)

	WidgetGetTooltipWindow func(
		widget *Widget) *T.GtkWindow

	WidgetTriggerTooltipQuery func(
		widget *Widget)

	WidgetSetTooltipText func(
		widget *Widget,
		text string)

	WidgetGetTooltipText func(
		widget *Widget) string

	WidgetSetTooltipMarkup func(
		widget *Widget,
		markup string)

	WidgetGetTooltipMarkup func(
		widget *Widget) string

	WidgetSetHasTooltip func(
		widget *Widget,
		hasTooltip T.Gboolean)

	WidgetGetHasTooltip func(
		widget *Widget) T.Gboolean

	RequisitionCopy func(
		requisition *T.GtkRequisition) *T.GtkRequisition

	RequisitionFree func(
		requisition *T.GtkRequisition)

	ContainerSetBorderWidth func(
		container *T.GtkContainer,
		borderWidth uint)

	ContainerGetBorderWidth func(
		container *T.GtkContainer) uint

	ContainerAdd func(
		container *T.GtkContainer,
		widget *Widget)

	ContainerRemove func(
		container *T.GtkContainer,
		widget *Widget)

	ContainerSetResizeMode func(
		container *T.GtkContainer,
		resizeMode T.GtkResizeMode)

	ContainerGetResizeMode func(
		container *T.GtkContainer) T.GtkResizeMode

	ContainerCheckResize func(
		container *T.GtkContainer)

	ContainerForeach func(
		container *T.GtkContainer,
		callback T.GtkCallback,
		callbackData T.Gpointer)

	ContainerForeachFull func(
		container *T.GtkContainer,
		callback T.GtkCallback,
		marshal T.GtkCallbackMarshal,
		callbackDataGpointer,
		notify T.GDestroyNotify)

	ContainerGetChildren func(
		container *T.GtkContainer) *T.GList

	ContainerPropagateExpose func(
		container *T.GtkContainer,
		child *Widget,
		event *T.GdkEventExpose)

	ContainerSetFocusChain func(
		container *T.GtkContainer,
		focusableWidgets *T.GList)

	ContainerGetFocusChain func(
		container *T.GtkContainer,
		focusableWidgets **T.GList) T.Gboolean

	ContainerUnsetFocusChain func(
		container *T.GtkContainer)

	ContainerSetReallocateRedraws func(
		container *T.GtkContainer,
		needsRedraws T.Gboolean)

	ContainerSetFocusChild func(
		container *T.GtkContainer,
		child *Widget)

	ContainerGetFocusChild func(
		container *T.GtkContainer) *Widget

	ContainerSetFocusVadjustment func(
		container *T.GtkContainer,
		adjustment *Adjustment)

	ContainerGetFocusVadjustment func(
		container *T.GtkContainer) *Adjustment

	ContainerSetFocusHadjustment func(
		container *T.GtkContainer,
		adjustment *Adjustment)

	ContainerGetFocusHadjustment func(
		container *T.GtkContainer) *Adjustment

	ContainerResizeChildren func(
		container *T.GtkContainer)

	ContainerChildType func(
		container *T.GtkContainer) T.GType

	ContainerClassInstallChildProperty func(
		cclass *T.GtkContainerClass,
		propertyId uint,
		pspec *T.GParamSpec)

	ContainerClassFindChildProperty func(
		cclass *T.GObjectClass,
		propertyName string) *T.GParamSpec

	ContainerClassListChildProperties func(
		cclass *T.GObjectClass,
		nProperties *uint) **T.GParamSpec

	ContainerAddWithProperties func(
		container *T.GtkContainer, widget *Widget,
		firstPropName string, v ...VArg)

	ContainerChildSet func(
		container *T.GtkContainer, child *Widget,
		firstPropName string, v ...VArg)

	ContainerChildGet func(container *T.GtkContainer,
		child *Widget, firstPropName string, v ...VArg)

	ContainerChildSetValist func(
		container *T.GtkContainer,
		child *Widget,
		firstPropertyName string,
		varArgs T.VaList)

	ContainerChildGetValist func(
		container *T.GtkContainer,
		child *Widget,
		firstPropertyName string,
		varArgs T.VaList)

	ContainerChildSetProperty func(
		container *T.GtkContainer,
		child *Widget,
		propertyName string,
		value *T.GValue)

	ContainerChildGetProperty func(
		container *T.GtkContainer,
		child *Widget,
		propertyName string,
		value *T.GValue)

	ContainerForall func(
		container *T.GtkContainer,
		callback T.GtkCallback,
		callbackData T.Gpointer)

	WindowGroupGetCurrentGrab func(
		windowGroup *T.GtkWindowGroup) *Widget

	MiscSetAlignment func(
		misc *T.GtkMisc,
		xalign float32,
		yalign float32)

	MiscGetAlignment func(
		misc *T.GtkMisc,
		xalign *float32,
		yalign *float32)

	MiscSetPadding func(
		misc *T.GtkMisc,
		xpad int,
		ypad int)

	MiscGetPadding func(
		misc *T.GtkMisc,
		xpad *int,
		ypad *int)

	MenuShellAppend func(
		menuShell *T.GtkMenuShell,
		child *Widget)

	MenuShellPrepend func(
		menuShell *T.GtkMenuShell,
		child *Widget)

	MenuShellInsert func(
		menuShell *T.GtkMenuShell,
		child *Widget,
		position int)

	MenuShellDeactivate func(
		menuShell *T.GtkMenuShell)

	MenuShellSelectItem func(
		menuShell *T.GtkMenuShell,
		menuItem *Widget)

	MenuShellDeselect func(
		menuShell *T.GtkMenuShell)

	MenuShellActivateItem func(
		menuShell *T.GtkMenuShell,
		menuItem *Widget,
		forceDeactivate T.Gboolean)

	MenuShellSelectFirst func(
		menuShell *T.GtkMenuShell,
		searchSensitive T.Gboolean)

	MenuShellCancel func(
		menuShell *T.GtkMenuShell)

	MenuShellGetTakeFocus func(
		menuShell *T.GtkMenuShell) T.Gboolean

	MenuShellSetTakeFocus func(
		menuShell *T.GtkMenuShell,
		takeFocus T.Gboolean)

	MenuPopup func(
		menu *T.GtkMenu,
		parentMenuShell *Widget,
		parentMenuItem *Widget,
		f T.GtkMenuPositionFunc,
		dataGpointer,
		button uint,
		activateTime T.GUint32)

	MenuReposition func(
		menu *T.GtkMenu)

	MenuPopdown func(
		menu *T.GtkMenu)

	MenuGetActive func(
		menu *T.GtkMenu) *Widget

	MenuSetActive func(
		menu *T.GtkMenu,
		index uint)

	MenuSetAccelGroup func(
		menu *T.GtkMenu,
		accelGroup *AccelGroup)

	MenuGetAccelGroup func(
		menu *T.GtkMenu) *AccelGroup

	MenuSetAccelPath func(
		menu *T.GtkMenu,
		accelPath string)

	MenuGetAccelPath func(
		menu *T.GtkMenu) string

	MenuAttachToWidget func(
		menu *T.GtkMenu,
		attachWidget *Widget,
		detacher T.GtkMenuDetachFunc)

	MenuDetach func(
		menu *T.GtkMenu)

	MenuGetAttachWidget func(
		menu *T.GtkMenu) *Widget

	MenuSetTearoffState func(
		menu *T.GtkMenu,
		tornOff T.Gboolean)

	MenuGetTearoffState func(
		menu *T.GtkMenu) T.Gboolean

	MenuSetTitle func(
		menu *T.GtkMenu,
		title string)

	MenuGetTitle func(
		menu *T.GtkMenu) string

	MenuReorderChild func(
		menu *T.GtkMenu,
		child *Widget,
		position int)

	MenuSetScreen func(
		menu *T.GtkMenu,
		screen *T.GdkScreen)

	MenuAttach func(
		menu *T.GtkMenu,
		child *Widget,
		leftAttach uint,
		rightAttach uint,
		topAttach uint,
		bottomAttach uint)

	MenuSetMonitor func(
		menu *T.GtkMenu,
		monitorNum int)

	MenuGetMonitor func(
		menu *T.GtkMenu) int

	MenuGetForAttachWidget func(
		widget *Widget) *T.GList

	MenuSetReserveToggleSize func(
		menu *T.GtkMenu,
		reserveToggleSize T.Gboolean)

	MenuGetReserveToggleSize func(
		menu *T.GtkMenu) T.Gboolean

	AccelGroupsActivate func(object *T.GObject, accelKey uint,
		accelMods T.GdkModifierType) T.Gboolean

	AccelGroupsFromObject func(object *T.GObject) *T.GSList

	AlignmentNew func(
		xalign float32,
		yalign float32,
		xscale float32,
		yscale float32) *Widget

	AlignmentSet func(
		alignment *T.GtkAlignment,
		xalign float32,
		yalign float32,
		xscale float32,
		yscale float32)

	AlignmentSetPadding func(
		alignment *T.GtkAlignment,
		paddingTop uint,
		paddingBottom uint,
		paddingLeft uint,
		paddingRight uint)

	AlignmentGetPadding func(
		alignment *T.GtkAlignment,
		paddingTop *uint,
		paddingBottom *uint,
		paddingLeft *uint,
		paddingRight *uint)

	BindingSetNew func(
		setName string) *T.GtkBindingSet

	BindingSetByClass func(
		objectClass T.Gpointer) *T.GtkBindingSet

	BindingSetFind func(
		setName string) *T.GtkBindingSet

	BindingsActivate func(
		object *T.GtkObject,
		keyval uint,
		modifiers T.GdkModifierType) T.Gboolean

	BindingsActivateEvent func(
		object *T.GtkObject,
		event *T.GdkEventKey) T.Gboolean

	BindingSetActivate func(
		bindingSet *T.GtkBindingSet,
		keyval uint,
		modifiers T.GdkModifierType,
		object *T.GtkObject) T.Gboolean

	BindingEntryClear func(
		bindingSet *T.GtkBindingSet,
		keyval uint,
		modifiers T.GdkModifierType)

	BindingParseBinding func(
		scanner *T.GScanner) uint

	BindingEntrySkip func(
		bindingSet *T.GtkBindingSet,
		keyval uint,
		modifiers T.GdkModifierType)

	BindingEntryAddSignal func(bindingSet *T.GtkBindingSet,
		keyval uint, modifiers T.GdkModifierType,
		signalName string, nArgs uint, varg ...VArg)

	BindingEntryAddSignall func(
		bindingSet *T.GtkBindingSet,
		keyval uint,
		modifiers T.GdkModifierType,
		signalName string,
		bindingArgs *T.GSList)

	BindingEntryRemove func(
		bindingSet *T.GtkBindingSet,
		keyval uint,
		modifiers T.GdkModifierType)

	BindingSetAddPath func(
		bindingSet *T.GtkBindingSet,
		pathType T.GtkPathType,
		pathPattern string,
		priority T.GtkPathPriorityType)

	SignalNewv func(
		name string,
		signalFlags T.GtkSignalRunType,
		objectType T.GType,
		functionOffset uint,
		marshaller T.GSignalCMarshaller,
		returnVal T.GType,
		nArgs uint,
		args *T.GType) uint

	SignalNew func(name string,
		signalFlags T.GtkSignalRunType,
		objectType T.GType, functionOffset uint,
		marshaller T.GSignalCMarshaller, returnVal T.GType,
		nArgs uint, v ...VArg) uint

	SignalEmitStopByName func(
		object *T.GtkObject, name string)

	SignalConnectObjectWhileAlive func(
		object *T.GtkObject,
		name string,
		f T.GCallback,
		aliveObject *T.GtkObject)

	SignalConnectWhileAlive func(
		object *T.GtkObject,
		name string,
		f T.GCallback,
		funcDataGpointer,
		aliveObject *T.GtkObject)

	SignalConnectFull func(
		object *T.GtkObject,
		name string,
		f T.GCallback,
		unsupported T.GtkCallbackMarshal,
		dataGpointer,
		destroyFunc T.GDestroyNotify,
		objectSignal int,
		after int) T.Gulong

	SignalEmitv func(
		object *T.GtkObject,
		signalId uint,
		args *T.GtkArg)

	SignalEmit func(object *T.GtkObject, signalId uint,
		v ...VArg)

	SignalEmitByName func(object *T.GtkObject,
		name string, v ...VArg)

	SignalEmitvByName func(
		object *T.GtkObject,
		name string,
		args *T.GtkArg)

	SignalCompatMatched func(
		object *T.GtkObject,
		f T.GCallback,
		dataGpointer,
		match T.GSignalMatchType,
		action uint)

	TreePathNew func() *T.GtkTreePath

	TreePathNewFromString func(
		path string) *T.GtkTreePath

	TreePathNewFromIndices func(firstIndex int,
		v ...VArg) *T.GtkTreePath

	TreePathToString func(
		path *T.GtkTreePath) string

	TreePathNewFirst func() *T.GtkTreePath

	TreePathAppendIndex func(
		path *T.GtkTreePath,
		index int)

	TreePathPrependIndex func(
		path *T.GtkTreePath,
		index int)

	TreePathGetDepth func(
		path *T.GtkTreePath) int

	TreePathGetIndices func(
		path *T.GtkTreePath) *int

	TreePathGetIndicesWithDepth func(
		path *T.GtkTreePath,
		depth *int) *int

	TreePathFree func(
		path *T.GtkTreePath)

	TreePathCopy func(
		path *T.GtkTreePath) *T.GtkTreePath

	TreePathCompare func(
		a *T.GtkTreePath,
		b *T.GtkTreePath) int

	TreePathNext func(
		path *T.GtkTreePath)

	TreePathPrev func(
		path *T.GtkTreePath) T.Gboolean

	TreePathUp func(
		path *T.GtkTreePath) T.Gboolean

	TreePathDown func(
		path *T.GtkTreePath)

	TreePathIsAncestor func(
		path *T.GtkTreePath,
		descendant *T.GtkTreePath) T.Gboolean

	TreePathIsDescendant func(
		path *T.GtkTreePath,
		ancestor *T.GtkTreePath) T.Gboolean

	TreeRowReferenceNew func(
		model *T.GtkTreeModel,
		path *T.GtkTreePath) *T.GtkTreeRowReference

	TreeRowReferenceNewProxy func(
		proxy *T.GObject,
		model *T.GtkTreeModel,
		path *T.GtkTreePath) *T.GtkTreeRowReference

	TreeRowReferenceGetPath func(
		reference *T.GtkTreeRowReference) *T.GtkTreePath

	TreeRowReferenceGetModel func(
		reference *T.GtkTreeRowReference) *T.GtkTreeModel

	TreeRowReferenceValid func(
		reference *T.GtkTreeRowReference) T.Gboolean

	TreeRowReferenceCopy func(
		reference *T.GtkTreeRowReference) *T.GtkTreeRowReference

	TreeRowReferenceFree func(
		reference *T.GtkTreeRowReference)

	TreeRowReferenceInserted func(
		proxy *T.GObject,
		path *T.GtkTreePath)

	TreeRowReferenceDeleted func(
		proxy *T.GObject,
		path *T.GtkTreePath)

	TreeRowReferenceReordered func(
		proxy *T.GObject,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter,
		newOrder *int)

	TreeIterCopy func(
		iter *T.GtkTreeIter) *T.GtkTreeIter

	TreeIterFree func(
		iter *T.GtkTreeIter)

	TreeModelGetFlags func(
		treeModel *T.GtkTreeModel) T.GtkTreeModelFlags

	TreeModelGetNColumns func(
		treeModel *T.GtkTreeModel) int

	TreeModelGetColumnType func(
		treeModel *T.GtkTreeModel,
		index int) T.GType

	TreeModelGetIter func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		path *T.GtkTreePath) T.Gboolean

	TreeModelGetIterFromString func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		pathString string) T.Gboolean

	TreeModelGetStringFromIter func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter) string

	TreeModelGetIterFirst func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gboolean

	TreeModelGetPath func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter) *T.GtkTreePath

	TreeModelGetValue func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		column int,
		value *T.GValue)

	TreeModelIterNext func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gboolean

	TreeModelIterChildren func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter) T.Gboolean

	TreeModelIterHasChild func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gboolean

	TreeModelIterNChildren func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter) int

	TreeModelIterNthChild func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		n int) T.Gboolean

	TreeModelIterParent func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		child *T.GtkTreeIter) T.Gboolean

	TreeModelRefNode func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter)

	TreeModelUnrefNode func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter)

	TreeModelGet func(treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter, v ...VArg)

	TreeModelGetValist func(
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		varArgs T.VaList)

	TreeModelForeach func(
		model *T.GtkTreeModel,
		f T.GtkTreeModelForeachFunc,
		userData T.Gpointer)

	TreeModelRowChanged func(
		treeModel *T.GtkTreeModel,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter)

	TreeModelRowInserted func(
		treeModel *T.GtkTreeModel,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter)

	TreeModelRowHasChildToggled func(
		treeModel *T.GtkTreeModel,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter)

	TreeModelRowDeleted func(
		treeModel *T.GtkTreeModel,
		path *T.GtkTreePath)

	TreeModelRowsReordered func(
		treeModel *T.GtkTreeModel,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter,
		newOrder *int)

	TreeSortableSortColumnChanged func(
		sortable *T.GtkTreeSortable)

	TreeSortableGetSortColumnId func(
		sortable *T.GtkTreeSortable,
		sortColumnId *int,
		order *T.GtkSortType) T.Gboolean

	TreeSortableSetSortColumnId func(
		sortable *T.GtkTreeSortable,
		sortColumnId int,
		order T.GtkSortType)

	TreeSortableSetSortFunc func(
		sortable *T.GtkTreeSortable,
		sortColumnId int,
		sortFunc T.GtkTreeIterCompareFunc,
		userDataGpointer,
		destroy T.GDestroyNotify)

	TreeSortableSetDefaultSortFunc func(
		sortable *T.GtkTreeSortable,
		sortFunc T.GtkTreeIterCompareFunc,
		userDataGpointer,
		destroy T.GDestroyNotify)

	TreeSortableHasDefaultSortFunc func(
		sortable *T.GtkTreeSortable) T.Gboolean

	ToggleButtonNewWithLabel func(
		label string) *Widget

	ToggleButtonNewWithMnemonic func(
		label string) *Widget

	ToggleButtonSetMode func(
		toggleButton *T.GtkToggleButton,
		drawIndicator T.Gboolean)

	ToggleButtonGetMode func(
		toggleButton *T.GtkToggleButton) T.Gboolean

	ToggleButtonSetActive func(
		toggleButton *T.GtkToggleButton,
		isActive T.Gboolean)

	ToggleButtonGetActive func(
		toggleButton *T.GtkToggleButton) T.Gboolean

	ToggleButtonToggled func(
		toggleButton *T.GtkToggleButton)

	ToggleButtonSetInconsistent func(
		toggleButton *T.GtkToggleButton,
		setting T.Gboolean)

	ToggleButtonGetInconsistent func(
		toggleButton *T.GtkToggleButton) T.Gboolean

	MenuItemNewWithLabel func(
		label string) *Widget

	MenuItemNewWithMnemonic func(
		label string) *Widget

	MenuItemSetSubmenu func(
		menuItem *T.GtkMenuItem,
		submenu *Widget)

	MenuItemGetSubmenu func(
		menuItem *T.GtkMenuItem) *Widget

	MenuItemSelect func(
		menuItem *T.GtkMenuItem)

	MenuItemDeselect func(
		menuItem *T.GtkMenuItem)

	MenuItemActivate func(
		menuItem *T.GtkMenuItem)

	MenuItemToggleSizeRequest func(
		menuItem *T.GtkMenuItem,
		requisition *int)

	MenuItemToggleSizeAllocate func(
		menuItem *T.GtkMenuItem,
		allocation int)

	MenuItemSetRightJustified func(
		menuItem *T.GtkMenuItem,
		rightJustified T.Gboolean)

	MenuItemGetRightJustified func(
		menuItem *T.GtkMenuItem) T.Gboolean

	MenuItemSetAccelPath func(
		menuItem *T.GtkMenuItem,
		accelPath string)

	MenuItemGetAccelPath func(
		menuItem *T.GtkMenuItem) string

	MenuItemSetLabel func(
		menuItem *T.GtkMenuItem,
		label string)

	MenuItemGetLabel func(
		menuItem *T.GtkMenuItem) string

	MenuItemSetUseUnderline func(
		menuItem *T.GtkMenuItem,
		setting T.Gboolean)

	MenuItemGetUseUnderline func(
		menuItem *T.GtkMenuItem) T.Gboolean

	MenuItemRemoveSubmenu func(
		menuItem *T.GtkMenuItem)

	TextTagNew func(
		name string) *T.GtkTextTag

	TextTagGetPriority func(
		tag *T.GtkTextTag) int

	TextTagSetPriority func(
		tag *T.GtkTextTag,
		priority int)

	TextTagEvent func(
		tag *T.GtkTextTag,
		eventObject *T.GObject,
		event *T.GdkEvent,
		iter *T.GtkTextIter) T.Gboolean

	TextAttributesNew func() *T.GtkTextAttributes

	TextAttributesCopy func(
		src *T.GtkTextAttributes) *T.GtkTextAttributes

	TextAttributesCopyValues func(
		src *T.GtkTextAttributes,
		dest *T.GtkTextAttributes)

	TextAttributesUnref func(
		values *T.GtkTextAttributes)

	TextAttributesRef func(
		values *T.GtkTextAttributes) *T.GtkTextAttributes

	TextChildAnchorNew func() *T.GtkTextChildAnchor

	TextChildAnchorGetWidgets func(
		anchor *T.GtkTextChildAnchor) *T.GList

	TextChildAnchorGetDeleted func(
		anchor *T.GtkTextChildAnchor) T.Gboolean

	TextIterGetBuffer func(
		iter *T.GtkTextIter) *T.GtkTextBuffer

	TextIterCopy func(
		iter *T.GtkTextIter) *T.GtkTextIter

	TextIterFree func(
		iter *T.GtkTextIter)

	TextIterGetOffset func(
		iter *T.GtkTextIter) int

	TextIterGetLine func(
		iter *T.GtkTextIter) int

	TextIterGetLineOffset func(
		iter *T.GtkTextIter) int

	TextIterGetLineIndex func(
		iter *T.GtkTextIter) int

	TextIterGetVisibleLineOffset func(
		iter *T.GtkTextIter) int

	TextIterGetVisibleLineIndex func(
		iter *T.GtkTextIter) int

	TextIterGetChar func(
		iter *T.GtkTextIter) T.Gunichar

	TextIterGetSlice func(
		start *T.GtkTextIter,
		end *T.GtkTextIter) string

	TextIterGetText func(
		start *T.GtkTextIter,
		end *T.GtkTextIter) string

	TextIterGetVisibleSlice func(
		start *T.GtkTextIter,
		end *T.GtkTextIter) string

	TextIterGetVisibleText func(
		start *T.GtkTextIter,
		end *T.GtkTextIter) string

	TextIterGetPixbuf func(
		iter *T.GtkTextIter) *T.GdkPixbuf

	TextIterGetMarks func(
		iter *T.GtkTextIter) *T.GSList

	TextIterGetChildAnchor func(
		iter *T.GtkTextIter) *T.GtkTextChildAnchor

	TextIterGetToggledTags func(
		iter *T.GtkTextIter,
		toggledOn T.Gboolean) *T.GSList

	TextIterBeginsTag func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	TextIterEndsTag func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	TextIterTogglesTag func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	TextIterHasTag func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	TextIterGetTags func(
		iter *T.GtkTextIter) *T.GSList

	TextIterEditable func(
		iter *T.GtkTextIter,
		defaultSetting T.Gboolean) T.Gboolean

	TextIterCanInsert func(
		iter *T.GtkTextIter,
		defaultEditability T.Gboolean) T.Gboolean

	TextIterStartsWord func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterEndsWord func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterInsideWord func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterStartsSentence func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterEndsSentence func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterInsideSentence func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterStartsLine func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterEndsLine func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterIsCursorPosition func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterGetCharsInLine func(
		iter *T.GtkTextIter) int

	TextIterGetBytesInLine func(
		iter *T.GtkTextIter) int

	TextIterGetAttributes func(
		iter *T.GtkTextIter,
		values *T.GtkTextAttributes) T.Gboolean

	TextIterGetLanguage func(
		iter *T.GtkTextIter) *T.PangoLanguage

	TextIterIsEnd func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterIsStart func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardChar func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterBackwardChar func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardChars func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterBackwardChars func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterForwardLine func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterBackwardLine func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardLines func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterBackwardLines func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterForwardWordEnd func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterBackwardWordStart func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardWordEnds func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterBackwardWordStarts func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterForwardVisibleLine func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterBackwardVisibleLine func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardVisibleLines func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterBackwardVisibleLines func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterForwardVisibleWordEnd func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterBackwardVisibleWordStart func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardVisibleWordEnds func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterBackwardVisibleWordStarts func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterForwardSentenceEnd func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterBackwardSentenceStart func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardSentenceEnds func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterBackwardSentenceStarts func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterForwardCursorPosition func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterBackwardCursorPosition func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardCursorPositions func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterBackwardCursorPositions func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterForwardVisibleCursorPosition func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterBackwardVisibleCursorPosition func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterForwardVisibleCursorPositions func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterBackwardVisibleCursorPositions func(
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextIterSetOffset func(
		iter *T.GtkTextIter,
		charOffset int)

	TextIterSetLine func(
		iter *T.GtkTextIter,
		lineNumber int)

	TextIterSetLineOffset func(
		iter *T.GtkTextIter,
		charOnLine int)

	TextIterSetLineIndex func(
		iter *T.GtkTextIter,
		byteOnLine int)

	TextIterForwardToEnd func(
		iter *T.GtkTextIter)

	TextIterForwardToLineEnd func(
		iter *T.GtkTextIter) T.Gboolean

	TextIterSetVisibleLineOffset func(
		iter *T.GtkTextIter,
		charOnLine int)

	TextIterSetVisibleLineIndex func(
		iter *T.GtkTextIter,
		byteOnLine int)

	TextIterForwardToTagToggle func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	TextIterBackwardToTagToggle func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	TextIterForwardFindChar func(
		iter *T.GtkTextIter,
		pred T.GtkTextCharPredicate,
		userDataGpointer,
		limit *T.GtkTextIter) T.Gboolean

	TextIterBackwardFindChar func(
		iter *T.GtkTextIter,
		pred T.GtkTextCharPredicate,
		userDataGpointer,
		limit *T.GtkTextIter) T.Gboolean

	TextIterForwardSearch func(
		iter *T.GtkTextIter,
		str string,
		flags T.GtkTextSearchFlags,
		matchStart *T.GtkTextIter,
		matchEnd *T.GtkTextIter,
		limit *T.GtkTextIter) T.Gboolean

	TextIterBackwardSearch func(
		iter *T.GtkTextIter,
		str string,
		flags T.GtkTextSearchFlags,
		matchStart *T.GtkTextIter,
		matchEnd *T.GtkTextIter,
		limit *T.GtkTextIter) T.Gboolean

	TextIterEqual func(
		lhs *T.GtkTextIter,
		rhs *T.GtkTextIter) T.Gboolean

	TextIterCompare func(
		lhs *T.GtkTextIter,
		rhs *T.GtkTextIter) int

	TextIterInRange func(
		iter *T.GtkTextIter,
		start *T.GtkTextIter,
		end *T.GtkTextIter) T.Gboolean

	TextIterOrder func(
		first *T.GtkTextIter,
		second *T.GtkTextIter)

	TargetListNew func(
		targets *T.GtkTargetEntry,
		ntargets uint) *T.GtkTargetList

	TargetListRef func(
		list *T.GtkTargetList) *T.GtkTargetList

	TargetListUnref func(
		list *T.GtkTargetList)

	TargetListAdd func(
		list *T.GtkTargetList,
		target T.GdkAtom,
		flags uint,
		info uint)

	TargetListAddTextTargets func(
		list *T.GtkTargetList,
		info uint)

	TargetListAddRichTextTargets func(
		list *T.GtkTargetList,
		info uint,
		deserializable T.Gboolean,
		buffer *T.GtkTextBuffer)

	TargetListAddImageTargets func(
		list *T.GtkTargetList,
		info uint,
		writable T.Gboolean)

	TargetListAddUriTargets func(
		list *T.GtkTargetList,
		info uint)

	TargetListAddTable func(
		list *T.GtkTargetList,
		targets *T.GtkTargetEntry,
		ntargets uint)

	TargetListRemove func(
		list *T.GtkTargetList,
		target T.GdkAtom)

	TargetListFind func(
		list *T.GtkTargetList,
		target T.GdkAtom,
		info *uint) T.Gboolean

	TargetTableNewFromList func(
		list *T.GtkTargetList,
		nTargets *int) *T.GtkTargetEntry

	TargetTableFree func(
		targets *T.GtkTargetEntry,
		nTargets int)

	SelectionOwnerSet func(
		widget *Widget,
		selection T.GdkAtom,
		time T.GUint32) T.Gboolean

	SelectionOwnerSetForDisplay func(
		display *T.GdkDisplay,
		widget *Widget,
		selection T.GdkAtom,
		time T.GUint32) T.Gboolean

	SelectionAddTarget func(
		widget *Widget,
		selection T.GdkAtom,
		target T.GdkAtom,
		info uint)

	SelectionAddTargets func(
		widget *Widget,
		selection T.GdkAtom,
		targets *T.GtkTargetEntry,
		ntargets uint)

	SelectionClearTargets func(
		widget *Widget,
		selection T.GdkAtom)

	SelectionConvert func(
		widget *Widget,
		selection T.GdkAtom,
		target T.GdkAtom,
		time T.GUint32) T.Gboolean

	SelectionDataGetSelection func(
		selectionData *T.GtkSelectionData) T.GdkAtom

	SelectionDataGetTarget func(
		selectionData *T.GtkSelectionData) T.GdkAtom

	SelectionDataGetDataType func(
		selectionData *T.GtkSelectionData) T.GdkAtom

	SelectionDataGetFormat func(
		selectionData *T.GtkSelectionData) int

	SelectionDataGetData func(
		selectionData *T.GtkSelectionData) *T.Guchar

	SelectionDataGetLength func(
		selectionData *T.GtkSelectionData) int

	SelectionDataGetDisplay func(
		selectionData *T.GtkSelectionData) *T.GdkDisplay

	SelectionDataSet func(
		selectionData *T.GtkSelectionData,
		typ T.GdkAtom,
		format int,
		data *T.Guchar,
		length int)

	SelectionDataSetText func(
		selectionData *T.GtkSelectionData,
		str string,
		len int) T.Gboolean

	SelectionDataGetText func(
		selectionData *T.GtkSelectionData) *T.Guchar

	SelectionDataSetPixbuf func(
		selectionData *T.GtkSelectionData,
		pixbuf *T.GdkPixbuf) T.Gboolean

	SelectionDataGetPixbuf func(
		selectionData *T.GtkSelectionData) *T.GdkPixbuf

	SelectionDataSetUris func(
		selectionData *T.GtkSelectionData,
		uris **T.Gchar) T.Gboolean

	SelectionDataGetUris func(
		selectionData *T.GtkSelectionData) **T.Gchar

	SelectionDataGetTargets func(
		selectionData *T.GtkSelectionData,
		targets **T.GdkAtom,
		nAtoms *int) T.Gboolean

	SelectionDataTargetsIncludeText func(
		selectionData *T.GtkSelectionData) T.Gboolean

	SelectionDataTargetsIncludeRichText func(
		selectionData *T.GtkSelectionData,
		buffer *T.GtkTextBuffer) T.Gboolean

	SelectionDataTargetsIncludeImage func(
		selectionData *T.GtkSelectionData,
		writable T.Gboolean) T.Gboolean

	SelectionDataTargetsIncludeUri func(
		selectionData *T.GtkSelectionData) T.Gboolean

	TargetsIncludeText func(
		targets *T.GdkAtom,
		nTargets int) T.Gboolean

	TargetsIncludeRichText func(
		targets *T.GdkAtom,
		nTargets int,
		buffer *T.GtkTextBuffer) T.Gboolean

	TargetsIncludeImage func(
		targets *T.GdkAtom,
		nTargets int,
		writable T.Gboolean) T.Gboolean

	TargetsIncludeUri func(
		targets *T.GdkAtom,
		nTargets int) T.Gboolean

	SelectionRemoveAll func(
		widget *Widget)

	SelectionClear func(
		widget *Widget,
		event *T.GdkEventSelection) T.Gboolean

	SelectionDataCopy func(
		data *T.GtkSelectionData) *T.GtkSelectionData

	SelectionDataFree func(
		data *T.GtkSelectionData)

	VboxNew func(
		homogeneous T.Gboolean,
		spacing int) *Widget

	DragGetData func(
		widget *Widget,
		context *T.GdkDragContext,
		target T.GdkAtom,
		time T.GUint32)

	DragFinish func(
		context *T.GdkDragContext,
		success T.Gboolean,
		del T.Gboolean,
		time T.GUint32)

	DragGetSourceWidget func(
		context *T.GdkDragContext) *Widget

	DragHighlight func(
		widget *Widget)

	DragUnhighlight func(
		widget *Widget)

	DragDestSet func(
		widget *Widget,
		flags T.GtkDestDefaults,
		targets *T.GtkTargetEntry,
		nTargets int,
		actions T.GdkDragAction)

	DragDestSetProxy func(
		widget *Widget,
		proxyWindow *T.GdkWindow,
		protocol T.GdkDragProtocol,
		useCoordinates T.Gboolean)

	DragDestUnset func(
		widget *Widget)

	DragDestFindTarget func(
		widget *Widget,
		context *T.GdkDragContext,
		targetList *T.GtkTargetList) T.GdkAtom

	DragDestGetTargetList func(
		widget *Widget) *T.GtkTargetList

	DragDestSetTargetList func(
		widget *Widget,
		targetList *T.GtkTargetList)

	DragDestAddTextTargets func(
		widget *Widget)

	DragDestAddImageTargets func(
		widget *Widget)

	DragDestAddUriTargets func(
		widget *Widget)

	DragDestSetTrackMotion func(
		widget *Widget,
		trackMotion T.Gboolean)

	DragDestGetTrackMotion func(
		widget *Widget) T.Gboolean

	DragSourceSet func(
		widget *Widget,
		startButtonMask T.GdkModifierType,
		targets *T.GtkTargetEntry,
		nTargets int,
		actions T.GdkDragAction)

	DragSourceUnset func(
		widget *Widget)

	DragSourceGetTargetList func(
		widget *Widget) *T.GtkTargetList

	DragSourceSetTargetList func(
		widget *Widget,
		targetList *T.GtkTargetList)

	DragSourceAddTextTargets func(
		widget *Widget)

	DragSourceAddImageTargets func(
		widget *Widget)

	DragSourceAddUriTargets func(
		widget *Widget)

	DragSourceSetIcon func(
		widget *Widget,
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	DragSourceSetIconPixbuf func(
		widget *Widget,
		pixbuf *T.GdkPixbuf)

	DragSourceSetIconStock func(
		widget *Widget,
		stockId string)

	DragSourceSetIconName func(
		widget *Widget,
		iconName string)

	DragBegin func(
		widget *Widget,
		targets *T.GtkTargetList,
		actions T.GdkDragAction,
		button int,
		event *T.GdkEvent) *T.GdkDragContext

	DragSetIconWidget func(
		context *T.GdkDragContext,
		widget *Widget,
		hotX int,
		hotY int)

	DragSetIconPixmap func(
		context *T.GdkDragContext,
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap,
		hotX int,
		hotY int)

	DragSetIconPixbuf func(
		context *T.GdkDragContext,
		pixbuf *T.GdkPixbuf,
		hotX int,
		hotY int)

	DragSetIconStock func(
		context *T.GdkDragContext,
		stockId string,
		hotX int,
		hotY int)

	DragSetIconName func(
		context *T.GdkDragContext,
		iconName string,
		hotX int,
		hotY int)

	DragSetIconDefault func(
		context *T.GdkDragContext)

	DragCheckThreshold func(
		widget *Widget,
		startX int,
		startY int,
		currentX int,
		currentY int) T.Gboolean

	DragSetDefaultIcon func(
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap,
		hotX int,
		hotY int)

	TreeModelFilterNew func(
		childModel *T.GtkTreeModel,
		root *T.GtkTreePath) *T.GtkTreeModel

	TreeModelFilterSetVisibleFunc func(
		filter *T.GtkTreeModelFilter,
		f T.GtkTreeModelFilterVisibleFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	TreeModelFilterSetModifyFunc func(
		filter *T.GtkTreeModelFilter,
		nColumns int,
		types *T.GType,
		f T.GtkTreeModelFilterModifyFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	TreeModelFilterSetVisibleColumn func(
		filter *T.GtkTreeModelFilter,
		column int)

	TreeModelFilterGetModel func(
		filter *T.GtkTreeModelFilter) *T.GtkTreeModel

	TreeModelFilterConvertChildIterToIter func(
		filter *T.GtkTreeModelFilter,
		filterIter *T.GtkTreeIter,
		childIter *T.GtkTreeIter) T.Gboolean

	TreeModelFilterConvertIterToChildIter func(
		filter *T.GtkTreeModelFilter,
		childIter *T.GtkTreeIter,
		filterIter *T.GtkTreeIter)

	TreeModelFilterConvertChildPathToPath func(
		filter *T.GtkTreeModelFilter,
		childPath *T.GtkTreePath) *T.GtkTreePath

	TreeModelFilterConvertPathToChildPath func(
		filter *T.GtkTreeModelFilter,
		filterPath *T.GtkTreePath) *T.GtkTreePath

	TreeModelFilterRefilter func(
		filter *T.GtkTreeModelFilter)

	TreeModelFilterClearCache func(
		filter *T.GtkTreeModelFilter)

	TreeViewNewWithModel func(
		model *T.GtkTreeModel) *Widget

	TreeViewGetModel func(
		treeView *T.GtkTreeView) *T.GtkTreeModel

	TreeViewSetModel func(
		treeView *T.GtkTreeView,
		model *T.GtkTreeModel)

	TreeViewGetSelection func(
		treeView *T.GtkTreeView) *T.GtkTreeSelection

	TreeViewGetHadjustment func(
		treeView *T.GtkTreeView) *Adjustment

	TreeViewSetHadjustment func(
		treeView *T.GtkTreeView, adjustment *Adjustment)

	TreeViewGetVadjustment func(
		treeView *T.GtkTreeView) *Adjustment

	TreeViewSetVadjustment func(
		treeView *T.GtkTreeView, adjustment *Adjustment)

	TreeViewGetHeadersVisible func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewSetHeadersVisible func(
		treeView *T.GtkTreeView,
		headersVisible T.Gboolean)

	TreeViewColumnsAutosize func(
		treeView *T.GtkTreeView)

	TreeViewGetHeadersClickable func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewSetHeadersClickable func(
		treeView *T.GtkTreeView,
		setting T.Gboolean)

	TreeViewSetRulesHint func(
		treeView *T.GtkTreeView,
		setting T.Gboolean)

	TreeViewGetRulesHint func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewAppendColumn func(
		treeView *T.GtkTreeView,
		column *TreeViewColumn) int

	TreeViewRemoveColumn func(
		treeView *T.GtkTreeView,
		column *TreeViewColumn) int

	TreeViewInsertColumn func(
		treeView *T.GtkTreeView,
		column *TreeViewColumn,
		position int) int

	TreeViewInsertColumnWithAttributes func(
		treeView *T.GtkTreeView, position int, title string,
		cell *CellRenderer, v ...VArg) int

	TreeViewInsertColumnWithDataFunc func(
		treeView *T.GtkTreeView,
		position int,
		title string,
		cell *CellRenderer,
		f T.GtkTreeCellDataFunc,
		dataGpointer,
		dnotify T.GDestroyNotify) int

	TreeViewGetColumn func(
		treeView *T.GtkTreeView,
		n int) *TreeViewColumn

	TreeViewGetColumns func(
		treeView *T.GtkTreeView) *T.GList

	TreeViewMoveColumnAfter func(
		treeView *T.GtkTreeView,
		column *TreeViewColumn,
		baseColumn *TreeViewColumn)

	TreeViewSetExpanderColumn func(
		treeView *T.GtkTreeView,
		column *TreeViewColumn)

	TreeViewGetExpanderColumn func(
		treeView *T.GtkTreeView) *TreeViewColumn

	TreeViewSetColumnDragFunction func(
		treeView *T.GtkTreeView,
		f T.GtkTreeViewColumnDropFunc,
		userDataGpointer,
		destroy T.GDestroyNotify)

	TreeViewScrollToPoint func(
		treeView *T.GtkTreeView,
		treeX int,
		treeY int)

	TreeViewScrollToCell func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath,
		column *TreeViewColumn,
		useAlign T.Gboolean,
		rowAlign float32,
		colAlign float32)

	TreeViewRowActivated func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath,
		column *TreeViewColumn)

	TreeViewExpandAll func(
		treeView *T.GtkTreeView)

	TreeViewCollapseAll func(
		treeView *T.GtkTreeView)

	TreeViewExpandToPath func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath)

	TreeViewExpandRow func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath,
		openAll T.Gboolean) T.Gboolean

	TreeViewCollapseRow func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath) T.Gboolean

	TreeViewMapExpandedRows func(
		treeView *T.GtkTreeView,
		f T.GtkTreeViewMappingFunc,
		data T.Gpointer)

	TreeViewRowExpanded func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath) T.Gboolean

	TreeViewSetReorderable func(
		treeView *T.GtkTreeView,
		reorderable T.Gboolean)

	TreeViewGetReorderable func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewSetCursor func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath,
		focusColumn *TreeViewColumn,
		startEditing T.Gboolean)

	TreeViewSetCursorOnCell func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath,
		focusColumn *TreeViewColumn,
		focusCell *CellRenderer,
		startEditing T.Gboolean)

	TreeViewGetCursor func(
		treeView *T.GtkTreeView,
		path **T.GtkTreePath,
		focusColumn **TreeViewColumn)

	TreeViewGetBinWindow func(
		treeView *T.GtkTreeView) *T.GdkWindow

	TreeViewGetPathAtPos func(
		treeView *T.GtkTreeView,
		x int,
		y int,
		path **T.GtkTreePath,
		column **TreeViewColumn,
		cellX *int,
		cellY *int) T.Gboolean

	TreeViewGetCellArea func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath,
		column *TreeViewColumn,
		rect *T.GdkRectangle)

	TreeViewGetBackgroundArea func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath,
		column *TreeViewColumn,
		rect *T.GdkRectangle)

	TreeViewGetVisibleRect func(
		treeView *T.GtkTreeView,
		visibleRect *T.GdkRectangle)

	TreeViewWidgetToTreeCoords func(
		treeView *T.GtkTreeView,
		wx int,
		wy int,
		tx *int,
		ty *int)

	TreeViewTreeToWidgetCoords func(
		treeView *T.GtkTreeView,
		tx int,
		ty int,
		wx *int,
		wy *int)

	TreeViewGetVisibleRange func(
		treeView *T.GtkTreeView,
		startPath **T.GtkTreePath,
		endPath **T.GtkTreePath) T.Gboolean

	TreeViewEnableModelDragSource func(
		treeView *T.GtkTreeView,
		startButtonMask T.GdkModifierType,
		targets *T.GtkTargetEntry,
		nTargets int,
		actions T.GdkDragAction)

	TreeViewEnableModelDragDest func(
		treeView *T.GtkTreeView,
		targets *T.GtkTargetEntry,
		nTargets int,
		actions T.GdkDragAction)

	TreeViewUnsetRowsDragSource func(
		treeView *T.GtkTreeView)

	TreeViewUnsetRowsDragDest func(
		treeView *T.GtkTreeView)

	TreeViewSetDragDestRow func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath,
		pos T.GtkTreeViewDropPosition)

	TreeViewGetDragDestRow func(
		treeView *T.GtkTreeView,
		path **T.GtkTreePath,
		pos *T.GtkTreeViewDropPosition)

	TreeViewGetDestRowAtPos func(
		treeView *T.GtkTreeView,
		dragX int,
		dragY int,
		path **T.GtkTreePath,
		pos *T.GtkTreeViewDropPosition) T.Gboolean

	TreeViewCreateRowDragIcon func(
		treeView *T.GtkTreeView,
		path *T.GtkTreePath) *T.GdkPixmap

	TreeViewSetEnableSearch func(
		treeView *T.GtkTreeView,
		enableSearch T.Gboolean)

	TreeViewGetEnableSearch func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewGetSearchColumn func(
		treeView *T.GtkTreeView) int

	TreeViewSetSearchColumn func(
		treeView *T.GtkTreeView,
		column int)

	TreeViewGetSearchEqualFunc func(
		treeView *T.GtkTreeView) T.GtkTreeViewSearchEqualFunc

	TreeViewSetSearchEqualFunc func(
		treeView *T.GtkTreeView,
		searchEqualFunc T.GtkTreeViewSearchEqualFunc,
		searchUserDataGpointer,
		searchDestroy T.GDestroyNotify)

	TreeViewGetSearchEntry func(
		treeView *T.GtkTreeView) *Entry

	TreeViewSetSearchEntry func(
		treeView *T.GtkTreeView,
		entry *Entry)

	TreeViewGetSearchPositionFunc func(
		treeView *T.GtkTreeView) T.GtkTreeViewSearchPositionFunc

	TreeViewSetSearchPositionFunc func(
		treeView *T.GtkTreeView,
		f T.GtkTreeViewSearchPositionFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	TreeViewConvertWidgetToTreeCoords func(
		treeView *T.GtkTreeView,
		wx int,
		wy int,
		tx *int,
		ty *int)

	TreeViewConvertTreeToWidgetCoords func(
		treeView *T.GtkTreeView,
		tx int,
		ty int,
		wx *int,
		wy *int)

	TreeViewConvertWidgetToBinWindowCoords func(
		treeView *T.GtkTreeView,
		wx int,
		wy int,
		bx *int,
		by *int)

	TreeViewConvertBinWindowToWidgetCoords func(
		treeView *T.GtkTreeView,
		bx int,
		by int,
		wx *int,
		wy *int)

	TreeViewConvertTreeToBinWindowCoords func(
		treeView *T.GtkTreeView,
		tx int,
		ty int,
		bx *int,
		by *int)

	TreeViewConvertBinWindowToTreeCoords func(
		treeView *T.GtkTreeView,
		bx int,
		by int,
		tx *int,
		ty *int)

	TreeViewSetDestroyCountFunc func(
		treeView *T.GtkTreeView,
		f T.GtkTreeDestroyCountFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	TreeViewSetFixedHeightMode func(
		treeView *T.GtkTreeView,
		enable T.Gboolean)

	TreeViewGetFixedHeightMode func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewSetHoverSelection func(
		treeView *T.GtkTreeView,
		hover T.Gboolean)

	TreeViewGetHoverSelection func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewSetHoverExpand func(
		treeView *T.GtkTreeView,
		expand T.Gboolean)

	TreeViewGetHoverExpand func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewSetRubberBanding func(
		treeView *T.GtkTreeView,
		enable T.Gboolean)

	TreeViewGetRubberBanding func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewIsRubberBandingActive func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewGetRowSeparatorFunc func(
		treeView *T.GtkTreeView) T.GtkTreeViewRowSeparatorFunc

	TreeViewSetRowSeparatorFunc func(
		treeView *T.GtkTreeView,
		f T.GtkTreeViewRowSeparatorFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	TreeViewGetGridLines func(
		treeView *T.GtkTreeView) T.GtkTreeViewGridLines

	TreeViewSetGridLines func(
		treeView *T.GtkTreeView,
		gridLines T.GtkTreeViewGridLines)

	TreeViewGetEnableTreeLines func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewSetEnableTreeLines func(
		treeView *T.GtkTreeView,
		enabled T.Gboolean)

	TreeViewSetShowExpanders func(
		treeView *T.GtkTreeView,
		enabled T.Gboolean)

	TreeViewGetShowExpanders func(
		treeView *T.GtkTreeView) T.Gboolean

	TreeViewSetLevelIndentation func(
		treeView *T.GtkTreeView,
		indentation int)

	TreeViewGetLevelIndentation func(
		treeView *T.GtkTreeView) int

	TreeViewSetTooltipRow func(
		treeView *T.GtkTreeView,
		tooltip *T.GtkTooltip,
		path *T.GtkTreePath)

	TreeViewSetTooltipCell func(
		treeView *T.GtkTreeView,
		tooltip *T.GtkTooltip,
		path *T.GtkTreePath,
		column *TreeViewColumn,
		cell *CellRenderer)

	TreeViewGetTooltipContext func(
		treeView *T.GtkTreeView,
		x *int,
		y *int,
		keyboardTip T.Gboolean,
		model **T.GtkTreeModel,
		path **T.GtkTreePath,
		iter *T.GtkTreeIter) T.Gboolean

	TreeViewSetTooltipColumn func(
		treeView *T.GtkTreeView,
		column int)

	TreeViewGetTooltipColumn func(
		treeView *T.GtkTreeView) int

	GcGet func(
		depth int,
		colormap *T.GdkColormap,
		values *T.GdkGCValues,
		valuesMask T.GdkGCValuesMask) *T.GdkGC

	GcRelease func(
		gc *T.GdkGC)

	PanedAdd1 func(
		paned *T.GtkPaned,
		child *Widget)

	PanedAdd2 func(
		paned *T.GtkPaned,
		child *Widget)

	PanedPack1 func(
		paned *T.GtkPaned,
		child *Widget,
		resize T.Gboolean,
		shrink T.Gboolean)

	PanedPack2 func(
		paned *T.GtkPaned,
		child *Widget,
		resize T.Gboolean,
		shrink T.Gboolean)

	PanedGetPosition func(
		paned *T.GtkPaned) int

	PanedSetPosition func(
		paned *T.GtkPaned,
		position int)

	PanedGetChild1 func(
		paned *T.GtkPaned) *Widget

	PanedGetChild2 func(
		paned *T.GtkPaned) *Widget

	PanedGetHandleWindow func(
		paned *T.GtkPaned) *T.GdkWindow

	PanedComputePosition func(
		paned *T.GtkPaned,
		allocation int,
		child1Req int,
		child2Req int)

	RulerSetMetric func(
		ruler *T.GtkRuler,
		metric T.GtkMetricType)

	RulerGetMetric func(
		ruler *T.GtkRuler) T.GtkMetricType

	RulerSetRange func(
		ruler *T.GtkRuler,
		lower float64,
		upper float64,
		position float64,
		maxSize float64)

	RulerGetRange func(
		ruler *T.GtkRuler,
		lower *float64,
		upper *float64,
		position *float64,
		maxSize *float64)

	RulerDrawTicks func(
		ruler *T.GtkRuler)

	RulerDrawPos func(
		ruler *T.GtkRuler)

	RangeSetUpdatePolicy func(
		r *T.GtkRange,
		policy T.GtkUpdateType)

	RangeGetUpdatePolicy func(
		r *T.GtkRange) T.GtkUpdateType

	RangeSetAdjustment func(
		r *T.GtkRange,
		adjustment *Adjustment)

	RangeGetAdjustment func(
		r *T.GtkRange) *Adjustment

	RangeSetInverted func(
		r *T.GtkRange,
		setting T.Gboolean)

	RangeGetInverted func(
		r *T.GtkRange) T.Gboolean

	RangeSetFlippable func(
		r *T.GtkRange,
		flippable T.Gboolean)

	RangeGetFlippable func(
		r *T.GtkRange) T.Gboolean

	RangeSetSliderSizeFixed func(
		r *T.GtkRange,
		sizeFixed T.Gboolean)

	RangeGetSliderSizeFixed func(
		r *T.GtkRange) T.Gboolean

	RangeSetMinSliderSize func(
		r *T.GtkRange,
		minSize T.Gboolean)

	RangeGetMinSliderSize func(
		r *T.GtkRange) int

	RangeGetRangeRect func(
		r *T.GtkRange,
		rangeRect *T.GdkRectangle)

	RangeGetSliderRange func(
		r *T.GtkRange,
		sliderStart *int,
		sliderEnd *int)

	RangeSetLowerStepperSensitivity func(
		r *T.GtkRange,
		sensitivity T.GtkSensitivityType)

	RangeGetLowerStepperSensitivity func(
		r *T.GtkRange) T.GtkSensitivityType

	RangeSetUpperStepperSensitivity func(
		r *T.GtkRange,
		sensitivity T.GtkSensitivityType)

	RangeGetUpperStepperSensitivity func(
		r *T.GtkRange) T.GtkSensitivityType

	RangeSetIncrements func(
		r *T.GtkRange,
		step float64,
		page float64)

	RangeSetRange func(
		r *T.GtkRange,
		min float64,
		max float64)

	RangeSetValue func(
		r *T.GtkRange,
		value float64)

	RangeGetValue func(
		r *T.GtkRange) float64

	RangeSetShowFillLevel func(
		r *T.GtkRange,
		showFillLevel T.Gboolean)

	RangeGetShowFillLevel func(
		r *T.GtkRange) T.Gboolean

	RangeSetRestrictToFillLevel func(
		r *T.GtkRange,
		restrictToFillLevel T.Gboolean)

	RangeGetRestrictToFillLevel func(
		r *T.GtkRange) T.Gboolean

	RangeSetFillLevel func(
		r *T.GtkRange,
		fillLevel float64)

	RangeGetFillLevel func(
		r *T.GtkRange) float64

	RangeSetRoundDigits func(
		r *T.GtkRange,
		roundDigits int)

	RangeGetRoundDigits func(
		r *T.GtkRange) int

	ScaleSetDigits func(
		scale *T.GtkScale,
		digits int)

	ScaleGetDigits func(
		scale *T.GtkScale) int

	ScaleSetDrawValue func(
		scale *T.GtkScale,
		drawValue T.Gboolean)

	ScaleGetDrawValue func(
		scale *T.GtkScale) T.Gboolean

	ScaleSetValuePos func(
		scale *T.GtkScale,
		pos T.GtkPositionType)

	ScaleGetValuePos func(
		scale *T.GtkScale) T.GtkPositionType

	ScaleGetLayout func(
		scale *T.GtkScale) *T.PangoLayout

	ScaleGetLayoutOffsets func(
		scale *T.GtkScale,
		x *int,
		y *int)

	ScaleAddMark func(
		scale *T.GtkScale,
		value float64,
		position T.GtkPositionType,
		markup string)

	ScaleClearMarks func(
		scale *T.GtkScale)

	HsvToRgb func(h, s, v float64, r, g, b *float64)

	RgbToHsv func(r, g, b float64, h, s, v *float64)

	IconSizeLookupForSettings func(settings *T.GtkSettings, size IconSize, width *int, height *int) T.Gboolean

	TooltipSetMarkup func(
		tooltip *T.GtkTooltip,
		markup string)

	TooltipSetText func(
		tooltip *T.GtkTooltip,
		text string)

	TooltipSetIcon func(
		tooltip *T.GtkTooltip,
		pixbuf *T.GdkPixbuf)

	TooltipSetIconFromStock func(
		tooltip *T.GtkTooltip,
		stockId string,
		size IconSize)

	TooltipSetIconFromIconName func(
		tooltip *T.GtkTooltip,
		iconName string,
		size IconSize)

	TooltipSetIconFromGicon func(
		tooltip *T.GtkTooltip,
		gicon *T.GIcon,
		size IconSize)

	TooltipSetCustom func(
		tooltip *T.GtkTooltip,
		customWidget *Widget)

	TooltipSetTipArea func(
		tooltip *T.GtkTooltip,
		rect *T.GdkRectangle)

	TooltipTriggerTooltipQuery func(
		display *T.GdkDisplay)

	CheckVersion func(
		requiredMajor uint,
		requiredMinor uint,
		requiredMicro uint) string

	ParseArgs func(
		argc *int,
		argv ***T.Char) T.Gboolean

	Init func(
		argc *int,
		argv ***T.Char)

	InitCheck func(
		argc *int,
		argv ***T.Char) T.Gboolean

	InitWithArgs func(
		argc *int,
		argv ***T.Char,
		parameterString string,
		entries *T.GOptionEntry,
		translationDomain string,
		error **T.GError) T.Gboolean

	GetOptionGroup func(
		openDefaultDisplay T.Gboolean) *T.GOptionGroup

	InitAbiCheck func(
		argc *int,
		argv ***T.Char,
		numChecks int,
		sizeof_GtkWindow T.SizeT,
		sizeof_GtkBox T.SizeT)

	InitCheckAbiCheck func(
		argc *int,
		argv ***T.Char,
		numChecks int,
		sizeof_GtkWindow T.SizeT,
		sizeof_GtkBox T.SizeT) T.Gboolean

	Exit func(
		errorCode int)

	SetLocale func() string

	DisableSetlocale func()

	GetDefaultLanguage func() *T.PangoLanguage

	EventsPending func() T.Gboolean

	MainDoEvent func(
		event *T.GdkEvent)

	Main func()

	MainLevel func() uint

	MainQuit func()

	MainIteration func() T.Gboolean

	MainIterationDo func(
		blocking T.Gboolean) T.Gboolean

	True func() T.Gboolean

	False func() T.Gboolean

	GrabAdd func(
		widget *Widget)

	GrabGetCurrent func() *Widget

	GrabRemove func(
		widget *Widget)

	InitAdd func(
		function T.GtkFunction,
		data T.Gpointer)

	QuitAddDestroy func(
		mainLevel uint,
		object *T.GtkObject)

	QuitAdd func(
		mainLevel uint,
		function T.GtkFunction,
		data T.Gpointer) uint

	QuitAddFull func(
		mainLevel uint,
		function T.GtkFunction,
		marshal T.GtkCallbackMarshal,
		dataGpointer,
		destroy T.GDestroyNotify) uint

	QuitRemove func(
		quitHandlerId uint)

	QuitRemoveByData func(
		data T.Gpointer)

	TimeoutAdd func(
		interval T.GUint32,
		function T.GtkFunction,
		data T.Gpointer) uint

	TimeoutAddFull func(
		interval T.GUint32,
		function T.GtkFunction,
		marshal T.GtkCallbackMarshal,
		dataGpointer,
		destroy T.GDestroyNotify) uint

	TimeoutRemove func(
		timeoutHandlerId uint)

	IdleAdd func(
		function T.GtkFunction,
		data T.Gpointer) uint

	IdleAddPriority func(
		priority int,
		function T.GtkFunction,
		data T.Gpointer) uint

	IdleAddFull func(
		priority int,
		function T.GtkFunction,
		marshal T.GtkCallbackMarshal,
		dataGpointer,
		destroy T.GDestroyNotify) uint

	IdleRemove func(
		idleHandlerId uint)

	IdleRemoveByData func(
		data T.Gpointer)

	InputAddFull func(
		source int,
		condition T.GdkInputCondition,
		function T.GdkInputFunction,
		marshal T.GtkCallbackMarshal,
		dataGpointer,
		destroy T.GDestroyNotify) uint

	InputRemove func(
		inputHandlerId uint)

	KeySnooperInstall func(
		snooper T.GtkKeySnoopFunc,
		funcData T.Gpointer) uint

	KeySnooperRemove func(
		snooperHandlerId uint)

	GetCurrentEvent func() *T.GdkEvent

	GetCurrentEventTime func() T.GUint32

	GetCurrentEventState func(
		state *T.GdkModifierType) T.Gboolean

	GetEventWidget func(
		event *T.GdkEvent) *Widget

	PropagateEvent func(
		widget *Widget,
		event *T.GdkEvent)

	MenuBarGetPackDirection func(
		menubar *T.GtkMenuBar) T.GtkPackDirection

	MenuBarSetPackDirection func(
		menubar *T.GtkMenuBar,
		packDir T.GtkPackDirection)

	MenuBarGetChildPackDirection func(
		menubar *T.GtkMenuBar) T.GtkPackDirection

	MenuBarSetChildPackDirection func(
		menubar *T.GtkMenuBar,
		childPackDir T.GtkPackDirection)

	TooltipsNew func() *T.GtkTooltips

	TooltipsEnable func(
		tooltips *T.GtkTooltips)

	TooltipsDisable func(
		tooltips *T.GtkTooltips)

	TooltipsSetDelay func(
		tooltips *T.GtkTooltips,
		delay uint)

	TooltipsSetTip func(
		tooltips *T.GtkTooltips,
		widget *Widget,
		tipText string,
		tipPrivate string)

	TooltipsDataGet func(
		widget *Widget) *T.GtkTooltipsData

	TooltipsForceWindow func(
		tooltips *T.GtkTooltips)

	TooltipsGetInfoFromTipWindow func(
		tipWindow *T.GtkWindow,
		tooltips **T.GtkTooltips,
		currentWidget **Widget) T.Gboolean

	SizeGroupNew func(
		mode T.GtkSizeGroupMode) *T.GtkSizeGroup

	SizeGroupSetMode func(
		sizeGroup *T.GtkSizeGroup,
		mode T.GtkSizeGroupMode)

	SizeGroupGetMode func(
		sizeGroup *T.GtkSizeGroup) T.GtkSizeGroupMode

	SizeGroupSetIgnoreHidden func(
		sizeGroup *T.GtkSizeGroup,
		ignoreHidden T.Gboolean)

	SizeGroupGetIgnoreHidden func(
		sizeGroup *T.GtkSizeGroup) T.Gboolean

	SizeGroupAddWidget func(
		sizeGroup *T.GtkSizeGroup,
		widget *Widget)

	SizeGroupRemoveWidget func(
		sizeGroup *T.GtkSizeGroup,
		widget *Widget)

	SizeGroupGetWidgets func(
		sizeGroup *T.GtkSizeGroup) *T.GSList

	ToolItemNew func() *T.GtkToolItem

	ToolItemSetHomogeneous func(
		toolItem *T.GtkToolItem,
		homogeneous T.Gboolean)

	ToolItemGetHomogeneous func(
		toolItem *T.GtkToolItem) T.Gboolean

	ToolItemSetExpand func(
		toolItem *T.GtkToolItem,
		expand T.Gboolean)

	ToolItemGetExpand func(
		toolItem *T.GtkToolItem) T.Gboolean

	ToolItemSetTooltip func(
		toolItem *T.GtkToolItem,
		tooltips *T.GtkTooltips,
		tipText string,
		tipPrivate string)

	ToolItemSetTooltipText func(
		toolItem *T.GtkToolItem,
		text string)

	ToolItemSetTooltipMarkup func(
		toolItem *T.GtkToolItem,
		markup string)

	ToolItemSetUseDragWindow func(
		toolItem *T.GtkToolItem,
		useDragWindow T.Gboolean)

	ToolItemGetUseDragWindow func(
		toolItem *T.GtkToolItem) T.Gboolean

	ToolItemSetVisibleHorizontal func(
		toolItem *T.GtkToolItem,
		visibleHorizontal T.Gboolean)

	ToolItemGetVisibleHorizontal func(
		toolItem *T.GtkToolItem) T.Gboolean

	ToolItemSetVisibleVertical func(
		toolItem *T.GtkToolItem,
		visibleVertical T.Gboolean)

	ToolItemGetVisibleVertical func(
		toolItem *T.GtkToolItem) T.Gboolean

	ToolItemGetIsImportant func(
		toolItem *T.GtkToolItem) T.Gboolean

	ToolItemSetIsImportant func(
		toolItem *T.GtkToolItem,
		isImportant T.Gboolean)

	ToolItemGetEllipsizeMode func(
		toolItem *T.GtkToolItem) T.PangoEllipsizeMode

	ToolItemGetIconSize func(
		toolItem *T.GtkToolItem) IconSize

	ToolItemGetOrientation func(
		toolItem *T.GtkToolItem) T.GtkOrientation

	ToolItemGetToolbarStyle func(
		toolItem *T.GtkToolItem) T.GtkToolbarStyle

	ToolItemGetReliefStyle func(
		toolItem *T.GtkToolItem) T.GtkReliefStyle

	ToolItemGetTextAlignment func(
		toolItem *T.GtkToolItem) float32

	ToolItemGetTextOrientation func(
		toolItem *T.GtkToolItem) T.GtkOrientation

	ToolItemGetTextSizeGroup func(
		toolItem *T.GtkToolItem) *T.GtkSizeGroup

	ToolItemRetrieveProxyMenuItem func(
		toolItem *T.GtkToolItem) *Widget

	ToolItemGetProxyMenuItem func(
		toolItem *T.GtkToolItem,
		menuItemId string) *Widget

	ToolItemSetProxyMenuItem func(
		toolItem *T.GtkToolItem,
		menuItemId string,
		menuItem *Widget)

	ToolItemRebuildMenu func(
		toolItem *T.GtkToolItem)

	ToolItemToolbarReconfigured func(
		toolItem *T.GtkToolItem)

	ToolButtonNew func(
		iconWidget *Widget,
		label string) *T.GtkToolItem

	ToolButtonNewFromStock func(
		stockId string) *T.GtkToolItem

	ToolButtonSetLabel func(
		button *T.GtkToolButton,
		label string)

	ToolButtonGetLabel func(
		button *T.GtkToolButton) string

	ToolButtonSetUseUnderline func(
		button *T.GtkToolButton,
		useUnderline T.Gboolean)

	ToolButtonGetUseUnderline func(
		button *T.GtkToolButton) T.Gboolean

	ToolButtonSetStockId func(
		button *T.GtkToolButton,
		stockId string)

	ToolButtonGetStockId func(
		button *T.GtkToolButton) string

	ToolButtonSetIconName func(
		button *T.GtkToolButton,
		iconName string)

	ToolButtonGetIconName func(
		button *T.GtkToolButton) string

	ToolButtonSetIconWidget func(
		button *T.GtkToolButton,
		iconWidget *Widget)

	ToolButtonGetIconWidget func(
		button *T.GtkToolButton) *Widget

	ToolButtonSetLabelWidget func(
		button *T.GtkToolButton,
		labelWidget *Widget)

	ToolButtonGetLabelWidget func(
		button *T.GtkToolButton) *Widget

	MenuToolButtonNew func(
		iconWidget *Widget,
		label string) *T.GtkToolItem

	MenuToolButtonNewFromStock func(
		stockId string) *T.GtkToolItem

	MenuToolButtonSetMenu func(
		button *T.GtkMenuToolButton,
		menu *Widget)

	MenuToolButtonGetMenu func(
		button *T.GtkMenuToolButton) *Widget

	MenuToolButtonSetArrowTooltip func(
		button *T.GtkMenuToolButton,
		tooltips *T.GtkTooltips,
		tipText string,
		tipPrivate string)

	MenuToolButtonSetArrowTooltipText func(
		button *T.GtkMenuToolButton,
		text string)

	MenuToolButtonSetArrowTooltipMarkup func(
		button *T.GtkMenuToolButton,
		markup string)

	MessageDialogNew func(
		parent *T.GtkWindow,
		flags DialogFlags,
		t T.GtkMessageType,
		buttons T.GtkButtonsType,
		messageFormat string,
		v ...VArg) *Widget

	MessageDialogNewWithMarkup func(
		parent *T.GtkWindow,
		flags DialogFlags,
		t T.GtkMessageType,
		buttons T.GtkButtonsType,
		messageFormat string,
		v ...VArg) *Widget

	MessageDialogSetImage func(
		dialog *T.GtkMessageDialog,
		image *Widget)

	MessageDialogGetImage func(
		dialog *T.GtkMessageDialog) *Widget

	MessageDialogSetMarkup func(
		messageDialog *T.GtkMessageDialog,
		str string)

	MessageDialogFormatSecondaryText func(
		messageDialog *T.GtkMessageDialog,
		messageFormat string,
		v ...VArg)

	MessageDialogFormatSecondaryMarkup func(
		messageDialog *T.GtkMessageDialog,
		messageFormat string,
		v ...VArg)

	MessageDialogGetMessageArea func(
		messageDialog *T.GtkMessageDialog) *Widget

	MountOperationNew func(
		parent *T.GtkWindow) *T.GMountOperation

	MountOperationIsShowing func(
		op *T.GtkMountOperation) T.Gboolean

	MountOperationSetParent func(
		op *T.GtkMountOperation,
		parent *T.GtkWindow)

	MountOperationGetParent func(
		op *T.GtkMountOperation) *T.GtkWindow

	MountOperationSetScreen func(
		op *T.GtkMountOperation,
		screen *T.GdkScreen)

	MountOperationGetScreen func(
		op *T.GtkMountOperation) *T.GdkScreen

	NotebookAppendPage func(
		notebook *T.GtkNotebook,
		child *Widget,
		tabLabel *Widget) int

	NotebookAppendPageMenu func(
		notebook *T.GtkNotebook,
		child *Widget,
		tabLabel *Widget,
		menuLabel *Widget) int

	NotebookPrependPage func(
		notebook *T.GtkNotebook,
		child *Widget,
		tabLabel *Widget) int

	NotebookPrependPageMenu func(
		notebook *T.GtkNotebook,
		child *Widget,
		tabLabel *Widget,
		menuLabel *Widget) int

	NotebookInsertPage func(
		notebook *T.GtkNotebook,
		child *Widget,
		tabLabel *Widget,
		position int) int

	NotebookInsertPageMenu func(
		notebook *T.GtkNotebook,
		child *Widget,
		tabLabel *Widget,
		menuLabel *Widget,
		position int) int

	NotebookRemovePage func(
		notebook *T.GtkNotebook,
		pageNum int)

	NotebookSetWindowCreationHook func(
		f T.GtkNotebookWindowCreationFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	NotebookSetGroupId func(
		notebook *T.GtkNotebook,
		groupId int)

	NotebookGetGroupId func(
		notebook *T.GtkNotebook) int

	NotebookSetGroup func(
		notebook *T.GtkNotebook,
		group T.Gpointer)

	NotebookGetGroup func(
		notebook *T.GtkNotebook) T.Gpointer

	NotebookSetGroupName func(
		notebook *T.GtkNotebook,
		groupName string)

	NotebookGetGroupName func(
		notebook *T.GtkNotebook) string

	NotebookGetCurrentPage func(
		notebook *T.GtkNotebook) int

	NotebookGetNthPage func(
		notebook *T.GtkNotebook,
		pageNum int) *Widget

	NotebookGetNPages func(
		notebook *T.GtkNotebook) int

	NotebookPageNum func(
		notebook *T.GtkNotebook,
		child *Widget) int

	NotebookSetCurrentPage func(
		notebook *T.GtkNotebook,
		pageNum int)

	NotebookNextPage func(
		notebook *T.GtkNotebook)

	NotebookPrevPage func(
		notebook *T.GtkNotebook)

	NotebookSetShowBorder func(
		notebook *T.GtkNotebook,
		showBorder T.Gboolean)

	NotebookGetShowBorder func(
		notebook *T.GtkNotebook) T.Gboolean

	NotebookSetShowTabs func(
		notebook *T.GtkNotebook,
		showTabs T.Gboolean)

	NotebookGetShowTabs func(
		notebook *T.GtkNotebook) T.Gboolean

	NotebookSetTabPos func(
		notebook *T.GtkNotebook,
		pos T.GtkPositionType)

	NotebookGetTabPos func(
		notebook *T.GtkNotebook) T.GtkPositionType

	NotebookSetHomogeneousTabs func(
		notebook *T.GtkNotebook,
		homogeneous T.Gboolean)

	NotebookSetTabBorder func(
		notebook *T.GtkNotebook,
		borderWidth uint)

	NotebookSetTabHborder func(
		notebook *T.GtkNotebook,
		tabHborder uint)

	NotebookSetTabVborder func(
		notebook *T.GtkNotebook,
		tabVborder uint)

	NotebookSetScrollable func(
		notebook *T.GtkNotebook,
		scrollable T.Gboolean)

	NotebookGetScrollable func(
		notebook *T.GtkNotebook) T.Gboolean

	NotebookGetTabHborder func(
		notebook *T.GtkNotebook) uint16

	NotebookGetTabVborder func(
		notebook *T.GtkNotebook) uint16

	NotebookPopupEnable func(
		notebook *T.GtkNotebook)

	NotebookPopupDisable func(
		notebook *T.GtkNotebook)

	NotebookGetTabLabel func(
		notebook *T.GtkNotebook,
		child *Widget) *Widget

	NotebookSetTabLabel func(
		notebook *T.GtkNotebook,
		child *Widget,
		tabLabel *Widget)

	NotebookSetTabLabelText func(
		notebook *T.GtkNotebook,
		child *Widget,
		tabText string)

	NotebookGetTabLabelText func(
		notebook *T.GtkNotebook,
		child *Widget) string

	NotebookGetMenuLabel func(
		notebook *T.GtkNotebook,
		child *Widget) *Widget

	NotebookSetMenuLabel func(
		notebook *T.GtkNotebook,
		child *Widget,
		menuLabel *Widget)

	NotebookSetMenuLabelText func(
		notebook *T.GtkNotebook,
		child *Widget,
		menuText string)

	NotebookGetMenuLabelText func(
		notebook *T.GtkNotebook,
		child *Widget) string

	NotebookQueryTabLabelPacking func(
		notebook *T.GtkNotebook,
		child *Widget,
		expand *T.Gboolean,
		fill *T.Gboolean,
		packType *T.GtkPackType)

	NotebookSetTabLabelPacking func(
		notebook *T.GtkNotebook,
		child *Widget,
		expand T.Gboolean,
		fill T.Gboolean,
		packType T.GtkPackType)

	NotebookReorderChild func(
		notebook *T.GtkNotebook,
		child *Widget,
		position int)

	NotebookGetTabReorderable func(
		notebook *T.GtkNotebook,
		child *Widget) T.Gboolean

	NotebookSetTabReorderable func(
		notebook *T.GtkNotebook,
		child *Widget,
		reorderable T.Gboolean)

	NotebookGetTabDetachable func(
		notebook *T.GtkNotebook,
		child *Widget) T.Gboolean

	NotebookSetTabDetachable func(
		notebook *T.GtkNotebook,
		child *Widget,
		detachable T.Gboolean)

	NotebookGetActionWidget func(
		notebook *T.GtkNotebook,
		packType T.GtkPackType) *Widget

	NotebookSetActionWidget func(
		notebook *T.GtkNotebook,
		widget *Widget,
		packType T.GtkPackType)

	OffscreenWindowGetPixmap func(
		offscreen *T.GtkOffscreenWindow) *T.GdkPixmap

	OffscreenWindowGetPixbuf func(
		offscreen *T.GtkOffscreenWindow) *T.GdkPixbuf

	OrientableSetOrientation func(
		orientable *T.GtkOrientable,
		orientation T.GtkOrientation)

	OrientableGetOrientation func(
		orientable *T.GtkOrientable) T.GtkOrientation

	PaperSizeNew func(
		name string) *T.GtkPaperSize

	PaperSizeNewFromPpd func(
		ppdName string,
		ppdDisplayName string,
		width float64,
		height float64) *T.GtkPaperSize

	PaperSizeNewCustom func(
		name string,
		displayName string,
		width float64,
		height float64,
		unit T.GtkUnit) *T.GtkPaperSize

	PaperSizeCopy func(
		other *T.GtkPaperSize) *T.GtkPaperSize

	PaperSizeFree func(
		size *T.GtkPaperSize)

	PaperSizeIsEqual func(
		size1 *T.GtkPaperSize,
		size2 *T.GtkPaperSize) T.Gboolean

	PaperSizeGetPaperSizes func(
		includeCustom T.Gboolean) *T.GList

	PaperSizeGetName func(
		size *T.GtkPaperSize) string

	PaperSizeGetDisplayName func(
		size *T.GtkPaperSize) string

	PaperSizeGetPpdName func(
		size *T.GtkPaperSize) string

	PaperSizeGetWidth func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) float64

	PaperSizeGetHeight func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) float64

	PaperSizeIsCustom func(
		size *T.GtkPaperSize) T.Gboolean

	PaperSizeSetSize func(
		size *T.GtkPaperSize,
		width float64,
		height float64,
		unit T.GtkUnit)

	PaperSizeGetDefaultTopMargin func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) float64

	PaperSizeGetDefaultBottomMargin func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) float64

	PaperSizeGetDefaultLeftMargin func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) float64

	PaperSizeGetDefaultRightMargin func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) float64

	PaperSizeGetDefault func() string

	PaperSizeNewFromKeyFile func(
		keyFile *T.GKeyFile,
		groupName string,
		error **T.GError) *T.GtkPaperSize

	PaperSizeToKeyFile func(
		size *T.GtkPaperSize,
		keyFile *T.GKeyFile,
		groupName string)

	PageSetupNew func() *T.GtkPageSetup

	PageSetupCopy func(
		other *T.GtkPageSetup) *T.GtkPageSetup

	PageSetupGetOrientation func(
		setup *T.GtkPageSetup) T.GtkPageOrientation

	PageSetupSetOrientation func(
		setup *T.GtkPageSetup,
		orientation T.GtkPageOrientation)

	PageSetupGetPaperSize func(
		setup *T.GtkPageSetup) *T.GtkPaperSize

	PageSetupSetPaperSize func(
		setup *T.GtkPageSetup,
		size *T.GtkPaperSize)

	PageSetupGetTopMargin func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) float64

	PageSetupSetTopMargin func(
		setup *T.GtkPageSetup,
		margin float64,
		unit T.GtkUnit)

	PageSetupGetBottomMargin func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) float64

	PageSetupSetBottomMargin func(
		setup *T.GtkPageSetup,
		margin float64,
		unit T.GtkUnit)

	PageSetupGetLeftMargin func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) float64

	PageSetupSetLeftMargin func(
		setup *T.GtkPageSetup,
		margin float64,
		unit T.GtkUnit)

	PageSetupGetRightMargin func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) float64

	PageSetupSetRightMargin func(
		setup *T.GtkPageSetup,
		margin float64,
		unit T.GtkUnit)

	PageSetupSetPaperSizeAndDefaultMargins func(
		setup *T.GtkPageSetup,
		size *T.GtkPaperSize)

	PageSetupGetPaperWidth func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) float64

	PageSetupGetPaperHeight func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) float64

	PageSetupGetPageWidth func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) float64

	PageSetupGetPageHeight func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) float64

	PageSetupNewFromFile func(
		fileName string,
		error **T.GError) *T.GtkPageSetup

	PageSetupLoadFile func(
		setup *T.GtkPageSetup,
		fileName string,
		error **T.GError) T.Gboolean

	PageSetupToFile func(
		setup *T.GtkPageSetup,
		fileName string,
		error **T.GError) T.Gboolean

	PageSetupNewFromKeyFile func(
		keyFile *T.GKeyFile,
		groupName string,
		error **T.GError) *T.GtkPageSetup

	PageSetupLoadKeyFile func(
		setup *T.GtkPageSetup,
		keyFile *T.GKeyFile,
		groupName string,
		error **T.GError) T.Gboolean

	PageSetupToKeyFile func(
		setup *T.GtkPageSetup,
		keyFile *T.GKeyFile,
		groupName string)

	SocketAddId func(
		socket *T.GtkSocket,
		windowId T.GdkNativeWindow)

	SocketGetId func(
		socket *T.GtkSocket) T.GdkNativeWindow

	SocketGetPlugWindow func(
		socket *T.GtkSocket) *T.GdkWindow

	SocketSteal func(
		socket *T.GtkSocket,
		wid T.GdkNativeWindow)

	PlugConstruct func(
		plug *T.GtkPlug,
		socketId T.GdkNativeWindow)

	PlugNew func(
		socketId T.GdkNativeWindow) *Widget

	PlugConstructForDisplay func(
		plug *T.GtkPlug,
		display *T.GdkDisplay,
		socketId T.GdkNativeWindow)

	PlugNewForDisplay func(
		display *T.GdkDisplay,
		socketId T.GdkNativeWindow) *Widget

	PlugGetId func(
		plug *T.GtkPlug) T.GdkNativeWindow

	PlugGetEmbedded func(
		plug *T.GtkPlug) T.Gboolean

	PlugGetSocketWindow func(
		plug *T.GtkPlug) *T.GdkWindow

	PrintContextGetCairoContext func(
		context *T.GtkPrintContext) *T.Cairo

	PrintContextGetPageSetup func(
		context *T.GtkPrintContext) *T.GtkPageSetup

	PrintContextGetWidth func(
		context *T.GtkPrintContext) float64

	PrintContextGetHeight func(
		context *T.GtkPrintContext) float64

	PrintContextGetDpiX func(
		context *T.GtkPrintContext) float64

	PrintContextGetDpiY func(
		context *T.GtkPrintContext) float64

	PrintContextGetHardMargins func(
		context *T.GtkPrintContext,
		top *float64,
		bottom *float64,
		left *float64,
		right *float64) T.Gboolean

	PrintContextGetPangoFontmap func(
		context *T.GtkPrintContext) *T.PangoFontMap

	PrintContextCreatePangoContext func(
		context *T.GtkPrintContext) *T.PangoContext

	PrintContextCreatePangoLayout func(
		context *T.GtkPrintContext) *T.PangoLayout

	PrintSettingsNew func() *T.GtkPrintSettings

	PrintSettingsCopy func(
		other *T.GtkPrintSettings) *T.GtkPrintSettings

	PrintSettingsNewFromFile func(
		fileName string,
		error **T.GError) *T.GtkPrintSettings

	PrintSettingsLoadFile func(
		settings *T.GtkPrintSettings,
		fileName string,
		error **T.GError) T.Gboolean

	PrintSettingsToFile func(
		settings *T.GtkPrintSettings,
		fileName string,
		error **T.GError) T.Gboolean

	PrintSettingsNewFromKeyFile func(
		keyFile *T.GKeyFile,
		groupName string,
		error **T.GError) *T.GtkPrintSettings

	PrintSettingsLoadKeyFile func(
		settings *T.GtkPrintSettings,
		keyFile *T.GKeyFile,
		groupName string,
		error **T.GError) T.Gboolean

	PrintSettingsToKeyFile func(
		settings *T.GtkPrintSettings,
		keyFile *T.GKeyFile,
		groupName string)

	PrintSettingsHasKey func(
		settings *T.GtkPrintSettings,
		key string) T.Gboolean

	PrintSettingsGet func(
		settings *T.GtkPrintSettings,
		key string) string

	PrintSettingsSet func(
		settings *T.GtkPrintSettings,
		key string,
		value string)

	PrintSettingsUnset func(
		settings *T.GtkPrintSettings,
		key string)

	PrintSettingsForeach func(
		settings *T.GtkPrintSettings,
		f T.GtkPrintSettingsFunc,
		userData T.Gpointer)

	PrintSettingsGetBool func(
		settings *T.GtkPrintSettings,
		key string) T.Gboolean

	PrintSettingsSetBool func(
		settings *T.GtkPrintSettings,
		key string,
		value T.Gboolean)

	PrintSettingsGetDouble func(
		settings *T.GtkPrintSettings,
		key string) float64

	PrintSettingsGetDoubleWithDefault func(
		settings *T.GtkPrintSettings,
		key string,
		def float64) float64

	PrintSettingsSetDouble func(
		settings *T.GtkPrintSettings,
		key string,
		value float64)

	PrintSettingsGetLength func(
		settings *T.GtkPrintSettings,
		key string,
		unit T.GtkUnit) float64

	PrintSettingsSetLength func(
		settings *T.GtkPrintSettings,
		key string,
		value float64,
		unit T.GtkUnit)

	PrintSettingsGetInt func(
		settings *T.GtkPrintSettings,
		key string) int

	PrintSettingsGetIntWithDefault func(
		settings *T.GtkPrintSettings,
		key string,
		def int) int

	PrintSettingsSetInt func(
		settings *T.GtkPrintSettings,
		key string,
		value int)

	PrintSettingsGetPrinter func(
		settings *T.GtkPrintSettings) string

	PrintSettingsSetPrinter func(
		settings *T.GtkPrintSettings,
		printer string)

	PrintSettingsGetOrientation func(
		settings *T.GtkPrintSettings) T.GtkPageOrientation

	PrintSettingsSetOrientation func(
		settings *T.GtkPrintSettings,
		orientation T.GtkPageOrientation)

	PrintSettingsGetPaperSize func(
		settings *T.GtkPrintSettings) *T.GtkPaperSize

	PrintSettingsSetPaperSize func(
		settings *T.GtkPrintSettings,
		paperSize *T.GtkPaperSize)

	PrintSettingsGetPaperWidth func(
		settings *T.GtkPrintSettings,
		unit T.GtkUnit) float64

	PrintSettingsSetPaperWidth func(
		settings *T.GtkPrintSettings,
		width float64,
		unit T.GtkUnit)

	PrintSettingsGetPaperHeight func(
		settings *T.GtkPrintSettings,
		unit T.GtkUnit) float64

	PrintSettingsSetPaperHeight func(
		settings *T.GtkPrintSettings,
		height float64,
		unit T.GtkUnit)

	PrintSettingsGetUseColor func(
		settings *T.GtkPrintSettings) T.Gboolean

	PrintSettingsSetUseColor func(
		settings *T.GtkPrintSettings,
		useColor T.Gboolean)

	PrintSettingsGetCollate func(
		settings *T.GtkPrintSettings) T.Gboolean

	PrintSettingsSetCollate func(
		settings *T.GtkPrintSettings,
		collate T.Gboolean)

	PrintSettingsGetReverse func(
		settings *T.GtkPrintSettings) T.Gboolean

	PrintSettingsSetReverse func(
		settings *T.GtkPrintSettings,
		reverse T.Gboolean)

	PrintSettingsGetDuplex func(
		settings *T.GtkPrintSettings) T.GtkPrintDuplex

	PrintSettingsSetDuplex func(
		settings *T.GtkPrintSettings,
		duplex T.GtkPrintDuplex)

	PrintSettingsGetQuality func(
		settings *T.GtkPrintSettings) T.GtkPrintQuality

	PrintSettingsSetQuality func(
		settings *T.GtkPrintSettings,
		quality T.GtkPrintQuality)

	PrintSettingsGetNCopies func(
		settings *T.GtkPrintSettings) int

	PrintSettingsSetNCopies func(
		settings *T.GtkPrintSettings,
		numCopies int)

	PrintSettingsGetNumberUp func(
		settings *T.GtkPrintSettings) int

	PrintSettingsSetNumberUp func(
		settings *T.GtkPrintSettings,
		numberUp int)

	PrintSettingsGetNumberUpLayout func(
		settings *T.GtkPrintSettings) T.GtkNumberUpLayout

	PrintSettingsSetNumberUpLayout func(
		settings *T.GtkPrintSettings,
		numberUpLayout T.GtkNumberUpLayout)

	PrintSettingsGetResolution func(
		settings *T.GtkPrintSettings) int

	PrintSettingsSetResolution func(
		settings *T.GtkPrintSettings,
		resolution int)

	PrintSettingsGetResolutionX func(
		settings *T.GtkPrintSettings) int

	PrintSettingsGetResolutionY func(
		settings *T.GtkPrintSettings) int

	PrintSettingsSetResolutionXy func(
		settings *T.GtkPrintSettings,
		resolutionX int,
		resolutionY int)

	PrintSettingsGetPrinterLpi func(
		settings *T.GtkPrintSettings) float64

	PrintSettingsSetPrinterLpi func(
		settings *T.GtkPrintSettings,
		lpi float64)

	PrintSettingsGetScale func(
		settings *T.GtkPrintSettings) float64

	PrintSettingsSetScale func(
		settings *T.GtkPrintSettings,
		scale float64)

	PrintSettingsGetPrintPages func(
		settings *T.GtkPrintSettings) T.GtkPrintPages

	PrintSettingsSetPrintPages func(
		settings *T.GtkPrintSettings,
		pages T.GtkPrintPages)

	PrintSettingsGetPageRanges func(
		settings *T.GtkPrintSettings,
		numRanges *int) *T.GtkPageRange

	PrintSettingsSetPageRanges func(
		settings *T.GtkPrintSettings,
		pageRanges *T.GtkPageRange,
		numRanges int)

	PrintSettingsGetPageSet func(
		settings *T.GtkPrintSettings) T.GtkPageSet

	PrintSettingsSetPageSet func(
		settings *T.GtkPrintSettings,
		pageSet T.GtkPageSet)

	PrintSettingsGetDefaultSource func(
		settings *T.GtkPrintSettings) string

	PrintSettingsSetDefaultSource func(
		settings *T.GtkPrintSettings,
		defaultSource string)

	PrintSettingsGetMediaType func(
		settings *T.GtkPrintSettings) string

	PrintSettingsSetMediaType func(
		settings *T.GtkPrintSettings,
		mediaType string)

	PrintSettingsGetDither func(
		settings *T.GtkPrintSettings) string

	PrintSettingsSetDither func(
		settings *T.GtkPrintSettings,
		dither string)

	PrintSettingsGetFinishings func(
		settings *T.GtkPrintSettings) string

	PrintSettingsSetFinishings func(
		settings *T.GtkPrintSettings,
		finishings string)

	PrintSettingsGetOutputBin func(
		settings *T.GtkPrintSettings) string

	PrintSettingsSetOutputBin func(
		settings *T.GtkPrintSettings,
		outputBin string)

	PrintOperationPreviewRenderPage func(
		preview *T.GtkPrintOperationPreview,
		pageNr int)

	PrintOperationPreviewEndPreview func(
		preview *T.GtkPrintOperationPreview)

	PrintOperationPreviewIsSelected func(
		preview *T.GtkPrintOperationPreview,
		pageNr int) T.Gboolean

	PrintErrorQuark func() T.GQuark

	PrintOperationNew func() *T.GtkPrintOperation

	PrintOperationSetDefaultPageSetup func(
		op *T.GtkPrintOperation,
		defaultPageSetup *T.GtkPageSetup)

	PrintOperationGetDefaultPageSetup func(
		op *T.GtkPrintOperation) *T.GtkPageSetup

	PrintOperationSetPrintSettings func(
		op *T.GtkPrintOperation,
		printSettings *T.GtkPrintSettings)

	PrintOperationGetPrintSettings func(
		op *T.GtkPrintOperation) *T.GtkPrintSettings

	PrintOperationSetJobName func(
		op *T.GtkPrintOperation,
		jobName string)

	PrintOperationSetNPages func(
		op *T.GtkPrintOperation,
		nPages int)

	PrintOperationSetCurrentPage func(
		op *T.GtkPrintOperation,
		currentPage int)

	PrintOperationSetUseFullPage func(
		op *T.GtkPrintOperation,
		fullPage T.Gboolean)

	PrintOperationSetUnit func(
		op *T.GtkPrintOperation,
		unit T.GtkUnit)

	PrintOperationSetExportFilename func(
		op *T.GtkPrintOperation,
		filename string)

	PrintOperationSetTrackPrintStatus func(
		op *T.GtkPrintOperation,
		trackStatus T.Gboolean)

	PrintOperationSetShowProgress func(
		op *T.GtkPrintOperation,
		showProgress T.Gboolean)

	PrintOperationSetAllowAsync func(
		op *T.GtkPrintOperation,
		allowAsync T.Gboolean)

	PrintOperationSetCustomTabLabel func(
		op *T.GtkPrintOperation,
		label string)

	PrintOperationRun func(
		op *T.GtkPrintOperation,
		action T.GtkPrintOperationAction,
		parent *T.GtkWindow,
		error **T.GError) T.GtkPrintOperationResult

	PrintOperationGetError func(
		op *T.GtkPrintOperation,
		error **T.GError)

	PrintOperationGetStatus func(
		op *T.GtkPrintOperation) T.GtkPrintStatus

	PrintOperationGetStatusString func(
		op *T.GtkPrintOperation) string

	PrintOperationIsFinished func(
		op *T.GtkPrintOperation) T.Gboolean

	PrintOperationCancel func(
		op *T.GtkPrintOperation)

	PrintOperationDrawPageFinish func(
		op *T.GtkPrintOperation)

	PrintOperationSetDeferDrawing func(
		op *T.GtkPrintOperation)

	PrintOperationSetSupportSelection func(
		op *T.GtkPrintOperation,
		supportSelection T.Gboolean)

	PrintOperationGetSupportSelection func(
		op *T.GtkPrintOperation) T.Gboolean

	PrintOperationSetHasSelection func(
		op *T.GtkPrintOperation,
		hasSelection T.Gboolean)

	PrintOperationGetHasSelection func(
		op *T.GtkPrintOperation) T.Gboolean

	PrintOperationSetEmbedPageSetup func(
		op *T.GtkPrintOperation,
		embed T.Gboolean)

	PrintOperationGetEmbedPageSetup func(
		op *T.GtkPrintOperation) T.Gboolean

	PrintOperationGetNPagesToPrint func(
		op *T.GtkPrintOperation) int

	PrintRunPageSetupDialog func(
		parent *T.GtkWindow,
		pageSetup *T.GtkPageSetup,
		settings *T.GtkPrintSettings) *T.GtkPageSetup

	PrintRunPageSetupDialogAsync func(
		parent *T.GtkWindow,
		pageSetup *T.GtkPageSetup,
		settings *T.GtkPrintSettings,
		doneCb T.GtkPageSetupDoneFunc,
		data T.Gpointer)

	ProgressSetShowText func(
		progress *T.GtkProgress,
		showText T.Gboolean)

	ProgressSetTextAlignment func(
		progress *T.GtkProgress,
		xAlign float32,
		yAlign float32)

	ProgressSetFormatString func(
		progress *T.GtkProgress,
		format string)

	ProgressSetAdjustment func(
		progress *T.GtkProgress,
		adjustment *Adjustment)

	ProgressConfigure func(
		progress *T.GtkProgress,
		value float64,
		min float64,
		max float64)

	ProgressSetPercentage func(
		progress *T.GtkProgress,
		percentage float64)

	ProgressSetValue func(
		progress *T.GtkProgress,
		value float64)

	ProgressGetValue func(
		progress *T.GtkProgress) float64

	ProgressSetActivityMode func(
		progress *T.GtkProgress,
		activityMode T.Gboolean)

	ProgressGetCurrentText func(
		progress *T.GtkProgress) string

	ProgressGetTextFromValue func(
		progress *T.GtkProgress,
		value float64) string

	ProgressGetCurrentPercentage func(
		progress *T.GtkProgress) float64

	ProgressGetPercentageFromValue func(
		progress *T.GtkProgress,
		value float64) float64

	ProgressBarPulse func(
		pbar *T.GtkProgressBar)

	ProgressBarSetText func(
		pbar *T.GtkProgressBar,
		text string)

	ProgressBarSetFraction func(
		pbar *T.GtkProgressBar,
		fraction float64)

	ProgressBarSetPulseStep func(
		pbar *T.GtkProgressBar,
		fraction float64)

	ProgressBarSetOrientation func(
		pbar *T.GtkProgressBar,
		orientation T.GtkProgressBarOrientation)

	ProgressBarGetText func(
		pbar *T.GtkProgressBar) string

	ProgressBarGetFraction func(
		pbar *T.GtkProgressBar) float64

	ProgressBarGetPulseStep func(
		pbar *T.GtkProgressBar) float64

	ProgressBarGetOrientation func(
		pbar *T.GtkProgressBar) T.GtkProgressBarOrientation

	ProgressBarSetEllipsize func(
		pbar *T.GtkProgressBar,
		mode T.PangoEllipsizeMode)

	ProgressBarGetEllipsize func(
		pbar *T.GtkProgressBar) T.PangoEllipsizeMode

	ProgressBarNewWithAdjustment func(
		adjustment *Adjustment) *Widget

	ProgressBarSetBarStyle func(
		pbar *T.GtkProgressBar,
		style T.GtkProgressBarStyle)

	ProgressBarSetDiscreteBlocks func(
		pbar *T.GtkProgressBar,
		blocks uint)

	ProgressBarSetActivityStep func(
		pbar *T.GtkProgressBar,
		step uint)

	ProgressBarSetActivityBlocks func(
		pbar *T.GtkProgressBar,
		blocks uint)

	ProgressBarUpdate func(
		pbar *T.GtkProgressBar,
		percentage float64)

	ToggleActionNew func(
		name string,
		label string,
		tooltip string,
		stockId string) *T.GtkToggleAction

	ToggleActionToggled func(
		action *T.GtkToggleAction)

	ToggleActionSetActive func(
		action *T.GtkToggleAction,
		isActive T.Gboolean)

	ToggleActionGetActive func(
		action *T.GtkToggleAction) T.Gboolean

	ToggleActionSetDrawAsRadio func(
		action *T.GtkToggleAction,
		drawAsRadio T.Gboolean)

	ToggleActionGetDrawAsRadio func(
		action *T.GtkToggleAction) T.Gboolean

	RadioActionNew func(
		name string,
		label string,
		tooltip string,
		stockId string,
		value int) *T.GtkRadioAction

	RadioActionGetGroup func(
		action *T.GtkRadioAction) *T.GSList

	RadioActionSetGroup func(
		action *T.GtkRadioAction,
		group *T.GSList)

	RadioActionGetCurrentValue func(
		action *T.GtkRadioAction) int

	RadioActionSetCurrentValue func(
		action *T.GtkRadioAction,
		currentValue int)

	RadioButtonNew func(
		group *T.GSList) *Widget

	RadioButtonNewFromWidget func(
		radioGroupMember *T.GtkRadioButton) *Widget

	RadioButtonNewWithLabel func(
		group *T.GSList,
		label string) *Widget

	RadioButtonNewWithLabelFromWidget func(
		radioGroupMember *T.GtkRadioButton,
		label string) *Widget

	RadioButtonNewWithMnemonic func(
		group *T.GSList,
		label string) *Widget

	RadioButtonNewWithMnemonicFromWidget func(
		radioGroupMember *T.GtkRadioButton,
		label string) *Widget

	RadioButtonGetGroup func(
		radioButton *T.GtkRadioButton) *T.GSList

	RadioButtonSetGroup func(
		radioButton *T.GtkRadioButton,
		group *T.GSList)

	RadioMenuItemNew func(
		group *T.GSList) *Widget

	RadioMenuItemNewWithLabel func(
		group *T.GSList,
		label string) *Widget

	RadioMenuItemNewWithMnemonic func(
		group *T.GSList,
		label string) *Widget

	RadioMenuItemNewFromWidget func(
		group *T.GtkRadioMenuItem) *Widget

	RadioMenuItemNewWithMnemonicFromWidget func(
		group *T.GtkRadioMenuItem,
		label string) *Widget

	RadioMenuItemNewWithLabelFromWidget func(
		group *T.GtkRadioMenuItem,
		label string) *Widget

	RadioMenuItemGetGroup func(
		radioMenuItem *T.GtkRadioMenuItem) *T.GSList

	RadioMenuItemSetGroup func(
		radioMenuItem *T.GtkRadioMenuItem,
		group *T.GSList)

	ToggleToolButtonNew func() *T.GtkToolItem

	ToggleToolButtonNewFromStock func(
		stockId string) *T.GtkToolItem

	ToggleToolButtonSetActive func(
		button *T.GtkToggleToolButton,
		isActive T.Gboolean)

	ToggleToolButtonGetActive func(
		button *T.GtkToggleToolButton) T.Gboolean

	RadioToolButtonNew func(
		group *T.GSList) *T.GtkToolItem

	RadioToolButtonNewFromStock func(
		group *T.GSList,
		stockId string) *T.GtkToolItem

	RadioToolButtonNewFromWidget func(
		group *T.GtkRadioToolButton) *T.GtkToolItem

	RadioToolButtonNewWithStockFromWidget func(
		group *T.GtkRadioToolButton,
		stockId string) *T.GtkToolItem

	RadioToolButtonGetGroup func(
		button *T.GtkRadioToolButton) *T.GSList

	RadioToolButtonSetGroup func(
		button *T.GtkRadioToolButton,
		group *T.GSList)

	RecentManagerErrorQuark func() T.GQuark

	RecentManagerNew func() *T.GtkRecentManager

	RecentManagerGetDefault func() *T.GtkRecentManager

	RecentManagerGetForScreen func(
		screen *T.GdkScreen) *T.GtkRecentManager

	RecentManagerSetScreen func(
		manager *T.GtkRecentManager,
		screen *T.GdkScreen)

	RecentManagerAddItem func(
		manager *T.GtkRecentManager,
		uri string) T.Gboolean

	RecentManagerAddFull func(
		manager *T.GtkRecentManager,
		uri string,
		recentData *T.GtkRecentData) T.Gboolean

	RecentManagerRemoveItem func(
		manager *T.GtkRecentManager,
		uri string,
		error **T.GError) T.Gboolean

	RecentManagerLookupItem func(
		manager *T.GtkRecentManager,
		uri string,
		error **T.GError) *T.GtkRecentInfo

	RecentManagerHasItem func(
		manager *T.GtkRecentManager,
		uri string) T.Gboolean

	RecentManagerMoveItem func(
		manager *T.GtkRecentManager,
		uri string,
		newUri string,
		error **T.GError) T.Gboolean

	RecentManagerSetLimit func(
		manager *T.GtkRecentManager,
		limit int)

	RecentManagerGetLimit func(
		manager *T.GtkRecentManager) int

	RecentManagerGetItems func(
		manager *T.GtkRecentManager) *T.GList

	RecentManagerPurgeItems func(
		manager *T.GtkRecentManager,
		error **T.GError) int

	RecentInfoRef func(
		info *T.GtkRecentInfo) *T.GtkRecentInfo

	RecentInfoUnref func(
		info *T.GtkRecentInfo)

	RecentInfoGetUri func(
		info *T.GtkRecentInfo) string

	RecentInfoGetDisplayName func(
		info *T.GtkRecentInfo) string

	RecentInfoGetDescription func(
		info *T.GtkRecentInfo) string

	RecentInfoGetMimeType func(
		info *T.GtkRecentInfo) string

	RecentInfoGetAdded func(
		info *T.GtkRecentInfo) T.TimeT

	RecentInfoGetModified func(
		info *T.GtkRecentInfo) T.TimeT

	RecentInfoGetVisited func(
		info *T.GtkRecentInfo) T.TimeT

	RecentInfoGetPrivateHint func(
		info *T.GtkRecentInfo) T.Gboolean

	RecentInfoGetApplicationInfo func(
		info *T.GtkRecentInfo,
		appName string,
		appExec **T.Gchar,
		count *uint,
		time *T.TimeT) T.Gboolean

	RecentInfoGetApplications func(
		info *T.GtkRecentInfo,
		length *T.Gsize) **T.Gchar

	RecentInfoLastApplication func(
		info *T.GtkRecentInfo) string

	RecentInfoHasApplication func(
		info *T.GtkRecentInfo,
		appName string) T.Gboolean

	RecentInfoGetGroups func(
		info *T.GtkRecentInfo,
		length *T.Gsize) **T.Gchar

	RecentInfoHasGroup func(
		info *T.GtkRecentInfo,
		groupName string) T.Gboolean

	RecentInfoGetIcon func(
		info *T.GtkRecentInfo,
		size int) *T.GdkPixbuf

	RecentInfoGetShortName func(
		info *T.GtkRecentInfo) string

	RecentInfoGetUriDisplay func(
		info *T.GtkRecentInfo) string

	RecentInfoGetAge func(
		info *T.GtkRecentInfo) int

	RecentInfoIsLocal func(
		info *T.GtkRecentInfo) T.Gboolean

	RecentInfoExists func(
		info *T.GtkRecentInfo) T.Gboolean

	RecentInfoMatch func(
		infoA *T.GtkRecentInfo,
		infoB *T.GtkRecentInfo) T.Gboolean

	RecentActionNew func(
		name string,
		label string,
		tooltip string,
		stockId string) *Action

	RecentActionNewForManager func(
		name string,
		label string,
		tooltip string,
		stockId string,
		manager *T.GtkRecentManager) *Action

	RecentActionGetShowNumbers func(
		action *T.GtkRecentAction) T.Gboolean

	RecentActionSetShowNumbers func(
		action *T.GtkRecentAction,
		showNumbers T.Gboolean)

	RecentFilterNew func() *T.GtkRecentFilter

	RecentFilterSetName func(
		filter *T.GtkRecentFilter,
		name string)

	RecentFilterGetName func(
		filter *T.GtkRecentFilter) string

	RecentFilterAddMimeType func(
		filter *T.GtkRecentFilter,
		mimeType string)

	RecentFilterAddPattern func(
		filter *T.GtkRecentFilter,
		pattern string)

	RecentFilterAddPixbufFormats func(
		filter *T.GtkRecentFilter)

	RecentFilterAddApplication func(
		filter *T.GtkRecentFilter,
		application string)

	RecentFilterAddGroup func(
		filter *T.GtkRecentFilter,
		group string)

	RecentFilterAddAge func(
		filter *T.GtkRecentFilter,
		days int)

	RecentFilterAddCustom func(
		filter *T.GtkRecentFilter,
		needed T.GtkRecentFilterFlags,
		f T.GtkRecentFilterFunc,
		dataGpointer,
		dataDestroy T.GDestroyNotify)

	RecentFilterGetNeeded func(
		filter *T.GtkRecentFilter) T.GtkRecentFilterFlags

	RecentFilterFilter func(
		filter *T.GtkRecentFilter,
		filterInfo *T.GtkRecentFilterInfo) T.Gboolean

	RecentChooserErrorQuark func() T.GQuark

	RecentChooserSetShowPrivate func(
		chooser *T.GtkRecentChooser,
		showPrivate T.Gboolean)

	RecentChooserGetShowPrivate func(
		chooser *T.GtkRecentChooser) T.Gboolean

	RecentChooserSetShowNotFound func(
		chooser *T.GtkRecentChooser,
		showNotFound T.Gboolean)

	RecentChooserGetShowNotFound func(
		chooser *T.GtkRecentChooser) T.Gboolean

	RecentChooserSetSelectMultiple func(
		chooser *T.GtkRecentChooser,
		selectMultiple T.Gboolean)

	RecentChooserGetSelectMultiple func(
		chooser *T.GtkRecentChooser) T.Gboolean

	RecentChooserSetLimit func(
		chooser *T.GtkRecentChooser,
		limit int)

	RecentChooserGetLimit func(
		chooser *T.GtkRecentChooser) int

	RecentChooserSetLocalOnly func(
		chooser *T.GtkRecentChooser,
		localOnly T.Gboolean)

	RecentChooserGetLocalOnly func(
		chooser *T.GtkRecentChooser) T.Gboolean

	RecentChooserSetShowTips func(
		chooser *T.GtkRecentChooser,
		showTips T.Gboolean)

	RecentChooserGetShowTips func(
		chooser *T.GtkRecentChooser) T.Gboolean

	RecentChooserSetShowNumbers func(
		chooser *T.GtkRecentChooser,
		showNumbers T.Gboolean)

	RecentChooserGetShowNumbers func(
		chooser *T.GtkRecentChooser) T.Gboolean

	RecentChooserSetShowIcons func(
		chooser *T.GtkRecentChooser,
		showIcons T.Gboolean)

	RecentChooserGetShowIcons func(
		chooser *T.GtkRecentChooser) T.Gboolean

	RecentChooserSetSortType func(
		chooser *T.GtkRecentChooser,
		sortType T.GtkRecentSortType)

	RecentChooserGetSortType func(
		chooser *T.GtkRecentChooser) T.GtkRecentSortType

	RecentChooserSetSortFunc func(
		chooser *T.GtkRecentChooser,
		sortFunc T.GtkRecentSortFunc,
		sortDataGpointer,
		dataDestroy T.GDestroyNotify)

	RecentChooserSetCurrentUri func(
		chooser *T.GtkRecentChooser,
		uri string,
		error **T.GError) T.Gboolean

	RecentChooserGetCurrentUri func(
		chooser *T.GtkRecentChooser) string

	RecentChooserGetCurrentItem func(
		chooser *T.GtkRecentChooser) *T.GtkRecentInfo

	RecentChooserSelectUri func(
		chooser *T.GtkRecentChooser,
		uri string,
		error **T.GError) T.Gboolean

	RecentChooserUnselectUri func(
		chooser *T.GtkRecentChooser,
		uri string)

	RecentChooserSelectAll func(
		chooser *T.GtkRecentChooser)

	RecentChooserUnselectAll func(
		chooser *T.GtkRecentChooser)

	RecentChooserGetItems func(
		chooser *T.GtkRecentChooser) *T.GList

	RecentChooserGetUris func(
		chooser *T.GtkRecentChooser,
		length *T.Gsize) **T.Gchar

	RecentChooserAddFilter func(
		chooser *T.GtkRecentChooser,
		filter *T.GtkRecentFilter)

	RecentChooserRemoveFilter func(
		chooser *T.GtkRecentChooser,
		filter *T.GtkRecentFilter)

	RecentChooserListFilters func(
		chooser *T.GtkRecentChooser) *T.GSList

	RecentChooserSetFilter func(
		chooser *T.GtkRecentChooser,
		filter *T.GtkRecentFilter)

	RecentChooserGetFilter func(
		chooser *T.GtkRecentChooser) *T.GtkRecentFilter

	RecentChooserDialogNew func(
		title string, parent *T.GtkWindow,
		firstButtonText string, v ...VArg) *Widget

	RecentChooserDialogNewForManager func(
		title string, parent *T.GtkWindow,
		manager *T.GtkRecentManager,
		firstButtonText string, v ...VArg) *Widget

	RecentChooserMenuNewForManager func(
		manager *T.GtkRecentManager) *Widget

	RecentChooserMenuGetShowNumbers func(
		menu *T.GtkRecentChooserMenu) T.Gboolean

	RecentChooserMenuSetShowNumbers func(
		menu *T.GtkRecentChooserMenu,
		showNumbers T.Gboolean)

	RecentChooserWidgetNewForManager func(
		manager *T.GtkRecentManager) *Widget

	ScaleButtonNew func(
		size IconSize,
		min float64,
		max float64,
		step float64,
		icons **T.Gchar) *Widget

	ScaleButtonSetIcons func(
		button *T.GtkScaleButton,
		icons **T.Gchar)

	ScaleButtonGetValue func(
		button *T.GtkScaleButton) float64

	ScaleButtonSetValue func(
		button *T.GtkScaleButton,
		value float64)

	ScaleButtonGetAdjustment func(
		button *T.GtkScaleButton) *Adjustment

	ScaleButtonSetAdjustment func(
		button *T.GtkScaleButton,
		adjustment *Adjustment)

	ScaleButtonGetPlusButton func(
		button *T.GtkScaleButton) *Widget

	ScaleButtonGetMinusButton func(
		button *T.GtkScaleButton) *Widget

	ScaleButtonGetPopup func(
		button *T.GtkScaleButton) *Widget

	ScaleButtonGetOrientation func(
		button *T.GtkScaleButton) T.GtkOrientation

	ScaleButtonSetOrientation func(
		button *T.GtkScaleButton,
		orientation T.GtkOrientation)

	VscrollbarNew func(
		adjustment *Adjustment) *Widget

	ViewportNew func(
		hadjustment,
		vadjustment *Adjustment) *Widget

	ViewportGetHadjustment func(
		viewport *T.GtkViewport) *Adjustment

	ViewportGetVadjustment func(
		viewport *T.GtkViewport) *Adjustment

	ViewportSetHadjustment func(
		viewport *T.GtkViewport,
		adjustment *Adjustment)

	ViewportSetVadjustment func(
		viewport *T.GtkViewport,
		adjustment *Adjustment)

	ViewportSetShadowType func(
		viewport *T.GtkViewport,
		t T.GtkShadowType)

	ViewportGetShadowType func(
		viewport *T.GtkViewport) T.GtkShadowType

	ViewportGetBinWindow func(
		viewport *T.GtkViewport) *T.GdkWindow

	ViewportGetViewWindow func(
		viewport *T.GtkViewport) *T.GdkWindow

	ScrolledWindowNew func(
		hadjustment,
		vadjustment *Adjustment) *Widget

	ScrolledWindowSetHadjustment func(
		scrolledWindow *T.GtkScrolledWindow,
		hadjustment *Adjustment)

	ScrolledWindowSetVadjustment func(
		scrolledWindow *T.GtkScrolledWindow,
		vadjustment *Adjustment)

	ScrolledWindowGetHadjustment func(
		scrolledWindow *T.GtkScrolledWindow) *Adjustment

	ScrolledWindowGetVadjustment func(
		scrolledWindow *T.GtkScrolledWindow) *Adjustment

	ScrolledWindowGetHscrollbar func(
		scrolledWindow *T.GtkScrolledWindow) *Widget

	ScrolledWindowGetVscrollbar func(
		scrolledWindow *T.GtkScrolledWindow) *Widget

	ScrolledWindowSetPolicy func(
		scrolledWindow *T.GtkScrolledWindow,
		hscrollbarPolicy T.GtkPolicyType,
		vscrollbarPolicy T.GtkPolicyType)

	ScrolledWindowGetPolicy func(
		scrolledWindow *T.GtkScrolledWindow,
		hscrollbarPolicy *T.GtkPolicyType,
		vscrollbarPolicy *T.GtkPolicyType)

	ScrolledWindowSetPlacement func(
		scrolledWindow *T.GtkScrolledWindow,
		windowPlacement T.GtkCornerType)

	ScrolledWindowUnsetPlacement func(
		scrolledWindow *T.GtkScrolledWindow)

	ScrolledWindowGetPlacement func(
		scrolledWindow *T.GtkScrolledWindow) T.GtkCornerType

	ScrolledWindowSetShadowType func(
		scrolledWindow *T.GtkScrolledWindow,
		t T.GtkShadowType)

	ScrolledWindowGetShadowType func(
		scrolledWindow *T.GtkScrolledWindow) T.GtkShadowType

	ScrolledWindowAddWithViewport func(
		scrolledWindow *T.GtkScrolledWindow,
		child *Widget)

	SeparatorToolItemNew func() *T.GtkToolItem

	SeparatorToolItemGetDraw func(
		item *T.GtkSeparatorToolItem) T.Gboolean

	SeparatorToolItemSetDraw func(
		item *T.GtkSeparatorToolItem,
		draw T.Gboolean)

	ShowUri func(
		screen *T.GdkScreen,
		uri string,
		timestamp T.GUint32,
		error **T.GError) T.Gboolean

	SpinButtonConfigure func(
		spinButton *T.GtkSpinButton,
		adjustment *Adjustment,
		climbRate float64,
		digits uint)

	SpinButtonNew func(
		adjustment *Adjustment,
		climbRate float64,
		digits uint) *Widget

	SpinButtonNewWithRange func(
		min,
		max,
		step float64) *Widget

	SpinButtonSetAdjustment func(
		spinButton *T.GtkSpinButton,
		adjustment *Adjustment)

	SpinButtonGetAdjustment func(
		spinButton *T.GtkSpinButton) *Adjustment

	SpinButtonSetDigits func(
		spinButton *T.GtkSpinButton,
		digits uint)

	SpinButtonGetDigits func(
		spinButton *T.GtkSpinButton) uint

	SpinButtonSetIncrements func(
		spinButton *T.GtkSpinButton,
		step float64,
		page float64)

	SpinButtonGetIncrements func(
		spinButton *T.GtkSpinButton,
		step *float64,
		page *float64)

	SpinButtonSetRange func(
		spinButton *T.GtkSpinButton,
		min float64,
		max float64)

	SpinButtonGetRange func(
		spinButton *T.GtkSpinButton,
		min *float64,
		max *float64)

	SpinButtonGetValue func(
		spinButton *T.GtkSpinButton) float64

	SpinButtonGetValueAsInt func(
		spinButton *T.GtkSpinButton) int

	SpinButtonSetValue func(
		spinButton *T.GtkSpinButton,
		value float64)

	SpinButtonSetUpdatePolicy func(
		spinButton *T.GtkSpinButton,
		policy T.GtkSpinButtonUpdatePolicy)

	SpinButtonGetUpdatePolicy func(
		spinButton *T.GtkSpinButton) T.GtkSpinButtonUpdatePolicy

	SpinButtonSetNumeric func(
		spinButton *T.GtkSpinButton,
		numeric T.Gboolean)

	SpinButtonGetNumeric func(
		spinButton *T.GtkSpinButton) T.Gboolean

	SpinButtonSpin func(
		spinButton *T.GtkSpinButton,
		direction T.GtkSpinType,
		increment float64)

	SpinButtonSetWrap func(
		spinButton *T.GtkSpinButton,
		wrap T.Gboolean)

	SpinButtonGetWrap func(
		spinButton *T.GtkSpinButton) T.Gboolean

	SpinButtonSetSnapToTicks func(
		spinButton *T.GtkSpinButton,
		snapToTicks T.Gboolean)

	SpinButtonGetSnapToTicks func(
		spinButton *T.GtkSpinButton) T.Gboolean

	SpinButtonUpdate func(
		spinButton *T.GtkSpinButton)

	SpinnerStart func(
		spinner *T.GtkSpinner)

	SpinnerStop func(
		spinner *T.GtkSpinner)

	StatusbarGetContextId func(
		statusbar *T.GtkStatusbar,
		contextDescription string) uint

	StatusbarPush func(
		statusbar *T.GtkStatusbar,
		contextId uint,
		text string) uint

	StatusbarPop func(
		statusbar *T.GtkStatusbar,
		contextId uint)

	StatusbarRemove func(
		statusbar *T.GtkStatusbar,
		contextId uint,
		messageId uint)

	StatusbarRemoveAll func(
		statusbar *T.GtkStatusbar,
		contextId uint)

	StatusbarSetHasResizeGrip func(
		statusbar *T.GtkStatusbar,
		setting T.Gboolean)

	StatusbarGetHasResizeGrip func(
		statusbar *T.GtkStatusbar) T.Gboolean

	StatusbarGetMessageArea func(
		statusbar *T.GtkStatusbar) *Widget

	StatusIconNew func() *T.GtkStatusIcon

	StatusIconNewFromPixbuf func(
		pixbuf *T.GdkPixbuf) *T.GtkStatusIcon

	StatusIconNewFromFile func(
		filename string) *T.GtkStatusIcon

	StatusIconNewFromStock func(
		stockId string) *T.GtkStatusIcon

	StatusIconNewFromIconName func(
		iconName string) *T.GtkStatusIcon

	StatusIconNewFromGicon func(
		icon *T.GIcon) *T.GtkStatusIcon

	StatusIconSetFromPixbuf func(
		statusIcon *T.GtkStatusIcon,
		pixbuf *T.GdkPixbuf)

	StatusIconSetFromFile func(
		statusIcon *T.GtkStatusIcon,
		filename string)

	StatusIconSetFromStock func(
		statusIcon *T.GtkStatusIcon,
		stockId string)

	StatusIconSetFromIconName func(
		statusIcon *T.GtkStatusIcon,
		iconName string)

	StatusIconSetFromGicon func(
		statusIcon *T.GtkStatusIcon,
		icon *T.GIcon)

	StatusIconGetStorageType func(
		statusIcon *T.GtkStatusIcon) ImageType

	StatusIconGetPixbuf func(
		statusIcon *T.GtkStatusIcon) *T.GdkPixbuf

	StatusIconGetStock func(
		statusIcon *T.GtkStatusIcon) string

	StatusIconGetIconName func(
		statusIcon *T.GtkStatusIcon) string

	StatusIconGetGicon func(
		statusIcon *T.GtkStatusIcon) *T.GIcon

	StatusIconGetSize func(
		statusIcon *T.GtkStatusIcon) int

	StatusIconSetScreen func(
		statusIcon *T.GtkStatusIcon,
		screen *T.GdkScreen)

	StatusIconGetScreen func(
		statusIcon *T.GtkStatusIcon) *T.GdkScreen

	StatusIconSetTooltip func(
		statusIcon *T.GtkStatusIcon,
		tooltipText string)

	StatusIconSetHasTooltip func(
		statusIcon *T.GtkStatusIcon,
		hasTooltip T.Gboolean)

	StatusIconSetTooltipText func(
		statusIcon *T.GtkStatusIcon,
		text string)

	StatusIconSetTooltipMarkup func(
		statusIcon *T.GtkStatusIcon,
		markup string)

	StatusIconSetTitle func(
		statusIcon *T.GtkStatusIcon,
		title string)

	StatusIconGetTitle func(
		statusIcon *T.GtkStatusIcon) string

	StatusIconSetName func(
		statusIcon *T.GtkStatusIcon,
		name string)

	StatusIconSetVisible func(
		statusIcon *T.GtkStatusIcon,
		visible T.Gboolean)

	StatusIconGetVisible func(
		statusIcon *T.GtkStatusIcon) T.Gboolean

	StatusIconSetBlinking func(
		statusIcon *T.GtkStatusIcon,
		blinking T.Gboolean)

	StatusIconGetBlinking func(
		statusIcon *T.GtkStatusIcon) T.Gboolean

	StatusIconIsEmbedded func(
		statusIcon *T.GtkStatusIcon) T.Gboolean

	StatusIconPositionMenu func(
		menu *T.GtkMenu,
		x *int,
		y *int,
		pushIn *T.Gboolean,
		userData T.Gpointer)

	StatusIconGetGeometry func(
		statusIcon *T.GtkStatusIcon,
		screen **T.GdkScreen,
		area *T.GdkRectangle,
		orientation *T.GtkOrientation) T.Gboolean

	StatusIconGetHasTooltip func(
		statusIcon *T.GtkStatusIcon) T.Gboolean

	StatusIconGetTooltipText func(
		statusIcon *T.GtkStatusIcon) string

	StatusIconGetTooltipMarkup func(
		statusIcon *T.GtkStatusIcon) string

	StatusIconGetX11WindowId func(
		statusIcon *T.GtkStatusIcon) T.GUint32

	StockAdd func(
		items *T.GtkStockItem,
		nItems uint)

	StockAddStatic func(
		items *T.GtkStockItem,
		nItems uint)

	StockLookup func(
		stockId string,
		item *T.GtkStockItem) T.Gboolean

	StockListIds func() *T.GSList

	StockItemCopy func(
		item *T.GtkStockItem) *T.GtkStockItem

	StockItemFree func(
		item *T.GtkStockItem)

	StockSetTranslateFunc func(
		domain string,
		f T.GtkTranslateFunc,
		dataGpointer,
		notify T.GDestroyNotify)

	TableNew func(
		rows,
		columns uint,
		homogeneous T.Gboolean) *Widget

	TableResize func(
		table *T.GtkTable,
		rows,
		columns uint)

	TableAttach func(
		table *T.GtkTable,
		child *Widget,
		leftAttach,
		rightAttach,
		topAttach,
		bottomAttach uint,
		xoptions,
		yoptions T.GtkAttachOptions,
		xpadding,
		ypadding uint)

	TableAttachDefaults func(
		table *T.GtkTable,
		widget *Widget,
		leftAttach,
		rightAttach,
		topAttach,
		bottomAttach uint)

	TableSetRowSpacing func(
		table *T.GtkTable,
		row uint,
		spacing uint)

	TableGetRowSpacing func(
		table *T.GtkTable,
		row uint) uint

	TableSetColSpacing func(
		table *T.GtkTable,
		column uint,
		spacing uint)

	TableGetColSpacing func(
		table *T.GtkTable,
		column uint) uint

	TableSetRowSpacings func(
		table *T.GtkTable,
		spacing uint)

	TableGetDefaultRowSpacing func(
		table *T.GtkTable) uint

	TableSetColSpacings func(
		table *T.GtkTable,
		spacing uint)

	TableGetDefaultColSpacing func(
		table *T.GtkTable) uint

	TableSetHomogeneous func(
		table *T.GtkTable,
		homogeneous T.Gboolean)

	TableGetHomogeneous func(
		table *T.GtkTable) T.Gboolean

	TableGetSize func(
		table *T.GtkTable,
		rows *uint,
		columns *uint)

	TextTagTableNew func() *T.GtkTextTagTable

	TextTagTableAdd func(
		table *T.GtkTextTagTable,
		tag *T.GtkTextTag)

	TextTagTableRemove func(
		table *T.GtkTextTagTable,
		tag *T.GtkTextTag)

	TextTagTableLookup func(
		table *T.GtkTextTagTable,
		name string) *T.GtkTextTag

	TextTagTableForeach func(
		table *T.GtkTextTagTable,
		f T.GtkTextTagTableForeach,
		data T.Gpointer)

	TextTagTableGetSize func(
		table *T.GtkTextTagTable) int

	TextMarkSetVisible func(
		mark *T.GtkTextMark,
		setting T.Gboolean)

	TextMarkGetVisible func(
		mark *T.GtkTextMark) T.Gboolean

	TextMarkNew func(
		name string,
		leftGravity T.Gboolean) *T.GtkTextMark

	TextMarkGetName func(
		mark *T.GtkTextMark) string

	TextMarkGetDeleted func(
		mark *T.GtkTextMark) T.Gboolean

	TextMarkGetBuffer func(
		mark *T.GtkTextMark) *T.GtkTextBuffer

	TextMarkGetLeftGravity func(
		mark *T.GtkTextMark) T.Gboolean

	TextBufferNew func(
		table *T.GtkTextTagTable) *T.GtkTextBuffer

	TextBufferGetLineCount func(
		buffer *T.GtkTextBuffer) int

	TextBufferGetCharCount func(
		buffer *T.GtkTextBuffer) int

	TextBufferGetTagTable func(
		buffer *T.GtkTextBuffer) *T.GtkTextTagTable

	TextBufferSetText func(
		buffer *T.GtkTextBuffer,
		text string,
		len int)

	TextBufferInsert func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		text string,
		len int)

	TextBufferInsertAtCursor func(
		buffer *T.GtkTextBuffer,
		text string,
		len int)

	TextBufferInsertInteractive func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		text string,
		len int,
		defaultEditable T.Gboolean) T.Gboolean

	TextBufferInsertInteractiveAtCursor func(
		buffer *T.GtkTextBuffer,
		text string,
		len int,
		defaultEditable T.Gboolean) T.Gboolean

	TextBufferInsertRange func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	TextBufferInsertRangeInteractive func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		start *T.GtkTextIter,
		end *T.GtkTextIter,
		defaultEditable T.Gboolean) T.Gboolean

	TextBufferInsertWithTags func(buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter, text string, len int,
		firstTag *T.GtkTextTag, v ...VArg)

	TextBufferInsertWithTagsByName func(
		buffer *T.GtkTextBuffer, iter *T.GtkTextIter, text string,
		len int, firstTagName string, v ...VArg)

	TextBufferDelete func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	TextBufferDeleteInteractive func(
		buffer *T.GtkTextBuffer,
		startIter *T.GtkTextIter,
		endIter *T.GtkTextIter,
		defaultEditable T.Gboolean) T.Gboolean

	TextBufferBackspace func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		interactive T.Gboolean,
		defaultEditable T.Gboolean) T.Gboolean

	TextBufferGetText func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter,
		includeHiddenChars T.Gboolean) string

	TextBufferGetSlice func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter,
		includeHiddenChars T.Gboolean) string

	TextBufferInsertPixbuf func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		pixbuf *T.GdkPixbuf)

	TextBufferInsertChildAnchor func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		anchor *T.GtkTextChildAnchor)

	TextBufferCreateChildAnchor func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter) *T.GtkTextChildAnchor

	TextBufferAddMark func(
		buffer *T.GtkTextBuffer,
		mark *T.GtkTextMark,
		where *T.GtkTextIter)

	TextBufferCreateMark func(
		buffer *T.GtkTextBuffer,
		markName string,
		where *T.GtkTextIter,
		leftGravity T.Gboolean) *T.GtkTextMark

	TextBufferMoveMark func(
		buffer *T.GtkTextBuffer,
		mark *T.GtkTextMark,
		where *T.GtkTextIter)

	TextBufferDeleteMark func(
		buffer *T.GtkTextBuffer,
		mark *T.GtkTextMark)

	TextBufferGetMark func(
		buffer *T.GtkTextBuffer,
		name string) *T.GtkTextMark

	TextBufferMoveMarkByName func(
		buffer *T.GtkTextBuffer,
		name string,
		where *T.GtkTextIter)

	TextBufferDeleteMarkByName func(
		buffer *T.GtkTextBuffer,
		name string)

	TextBufferGetInsert func(
		buffer *T.GtkTextBuffer) *T.GtkTextMark

	TextBufferGetSelectionBound func(
		buffer *T.GtkTextBuffer) *T.GtkTextMark

	TextBufferPlaceCursor func(
		buffer *T.GtkTextBuffer,
		where *T.GtkTextIter)

	TextBufferSelectRange func(
		buffer *T.GtkTextBuffer,
		ins *T.GtkTextIter,
		bound *T.GtkTextIter)

	TextBufferApplyTag func(
		buffer *T.GtkTextBuffer,
		tag *T.GtkTextTag,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	TextBufferRemoveTag func(
		buffer *T.GtkTextBuffer,
		tag *T.GtkTextTag,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	TextBufferApplyTagByName func(
		buffer *T.GtkTextBuffer,
		name string,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	TextBufferRemoveTagByName func(
		buffer *T.GtkTextBuffer,
		name string,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	TextBufferRemoveAllTags func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	TextBufferCreateTag func(
		buffer *T.GtkTextBuffer, tagName string,
		firstPropertyName string, v ...VArg) *T.GtkTextTag

	TextBufferGetIterAtLineOffset func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		lineNumber int,
		charOffset int)

	TextBufferGetIterAtLineIndex func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		lineNumber int,
		byteIndex int)

	TextBufferGetIterAtOffset func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		charOffset int)

	TextBufferGetIterAtLine func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		lineNumber int)

	TextBufferGetStartIter func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter)

	TextBufferGetEndIter func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter)

	TextBufferGetBounds func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	TextBufferGetIterAtMark func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		mark *T.GtkTextMark)

	TextBufferGetIterAtChildAnchor func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		anchor *T.GtkTextChildAnchor)

	TextBufferGetModified func(
		buffer *T.GtkTextBuffer) T.Gboolean

	TextBufferSetModified func(
		buffer *T.GtkTextBuffer,
		setting T.Gboolean)

	TextBufferGetHasSelection func(
		buffer *T.GtkTextBuffer) T.Gboolean

	TextBufferAddSelectionClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *Clipboard)

	TextBufferRemoveSelectionClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *Clipboard)

	TextBufferCutClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *Clipboard,
		defaultEditable T.Gboolean)

	TextBufferCopyClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *Clipboard)

	TextBufferPasteClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *Clipboard,
		overrideLocation *T.GtkTextIter,
		defaultEditable T.Gboolean)

	TextBufferGetSelectionBounds func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter) T.Gboolean

	TextBufferDeleteSelection func(
		buffer *T.GtkTextBuffer,
		interactive T.Gboolean,
		defaultEditable T.Gboolean) T.Gboolean

	TextBufferBeginUserAction func(
		buffer *T.GtkTextBuffer)

	TextBufferEndUserAction func(
		buffer *T.GtkTextBuffer)

	TextBufferGetCopyTargetList func(
		buffer *T.GtkTextBuffer) *T.GtkTargetList

	TextBufferGetPasteTargetList func(
		buffer *T.GtkTextBuffer) *T.GtkTargetList

	TextBufferRegisterSerializeFormat func(
		buffer *T.GtkTextBuffer,
		mimeType string,
		function T.GtkTextBufferSerializeFunc,
		userDataGpointer,
		userDataDestroy T.GDestroyNotify) T.GdkAtom

	TextBufferRegisterSerializeTagset func(
		buffer *T.GtkTextBuffer,
		tagsetName string) T.GdkAtom

	TextBufferRegisterDeserializeFormat func(
		buffer *T.GtkTextBuffer,
		mimeType string,
		function T.GtkTextBufferDeserializeFunc,
		userDataGpointer,
		userDataDestroy T.GDestroyNotify) T.GdkAtom

	TextBufferRegisterDeserializeTagset func(
		buffer *T.GtkTextBuffer,
		tagsetName string) T.GdkAtom

	TextBufferUnregisterSerializeFormat func(
		buffer *T.GtkTextBuffer,
		format T.GdkAtom)

	TextBufferUnregisterDeserializeFormat func(
		buffer *T.GtkTextBuffer,
		format T.GdkAtom)

	TextBufferDeserializeSetCanCreateTags func(
		buffer *T.GtkTextBuffer,
		format T.GdkAtom,
		canCreateTags T.Gboolean)

	TextBufferDeserializeGetCanCreateTags func(
		buffer *T.GtkTextBuffer,
		format T.GdkAtom) T.Gboolean

	TextBufferGetSerializeFormats func(
		buffer *T.GtkTextBuffer,
		nFormats *int) *T.GdkAtom

	TextBufferGetDeserializeFormats func(
		buffer *T.GtkTextBuffer,
		nFormats *int) *T.GdkAtom

	TextBufferSerialize func(
		registerBuffer *T.GtkTextBuffer,
		contentBuffer *T.GtkTextBuffer,
		format T.GdkAtom,
		start *T.GtkTextIter,
		end *T.GtkTextIter,
		length *T.Gsize) *uint8

	TextBufferDeserialize func(
		registerBuffer *T.GtkTextBuffer,
		contentBuffer *T.GtkTextBuffer,
		format T.GdkAtom,
		iter *T.GtkTextIter,
		data *uint8,
		length T.Gsize,
		error **T.GError) T.Gboolean

	TextViewNewWithBuffer func(buffer *T.GtkTextBuffer) *Widget

	TextViewSetBuffer func(textView *T.GtkTextView,
		buffer *T.GtkTextBuffer)

	TextViewGetBuffer func(textView *T.GtkTextView) *T.GtkTextBuffer

	TextViewScrollToIter func(textView *T.GtkTextView,
		iter *T.GtkTextIter,
		withinMargin float64,
		useAlign T.Gboolean,
		xalign float64,
		yalign float64) T.Gboolean

	TextViewScrollToMark func(textView *T.GtkTextView,
		mark *T.GtkTextMark,
		withinMargin float64,
		useAlign T.Gboolean,
		xalign float64,
		yalign float64)

	TextViewScrollMarkOnscreen func(textView *T.GtkTextView,
		mark *T.GtkTextMark)

	TextViewMoveMarkOnscreen func(textView *T.GtkTextView,
		mark *T.GtkTextMark) T.Gboolean

	TextViewPlaceCursorOnscreen func(textView *T.GtkTextView) T.Gboolean

	TextViewGetVisibleRect func(textView *T.GtkTextView,
		visibleRect *T.GdkRectangle)

	TextViewSetCursorVisible func(textView *T.GtkTextView,
		setting T.Gboolean)

	TextViewGetCursorVisible func(textView *T.GtkTextView) T.Gboolean

	TextViewGetIterLocation func(textView *T.GtkTextView,
		iter *T.GtkTextIter,
		location *T.GdkRectangle)

	TextViewGetIterAtLocation func(textView *T.GtkTextView,
		iter *T.GtkTextIter,
		x int,
		y int)

	TextViewGetIterAtPosition func(textView *T.GtkTextView,
		iter *T.GtkTextIter,
		trailing *int,
		x int,
		y int)

	TextViewGetLineYrange func(textView *T.GtkTextView,
		iter *T.GtkTextIter,
		y *int,
		height *int)

	TextViewGetLineAtY func(textView *T.GtkTextView,
		targetIter *T.GtkTextIter,
		y int,
		lineTop *int)

	TextViewBufferToWindowCoords func(textView *T.GtkTextView,
		win T.GtkTextWindowType,
		bufferX int,
		bufferY int,
		windowX *int,
		windowY *int)

	TextViewWindowToBufferCoords func(textView *T.GtkTextView,
		win T.GtkTextWindowType,
		windowX int,
		windowY int,
		bufferX *int,
		bufferY *int)

	TextViewGetHadjustment func(textView *T.GtkTextView) *Adjustment

	TextViewGetVadjustment func(textView *T.GtkTextView) *Adjustment

	TextViewGetWindow func(textView *T.GtkTextView,
		win T.GtkTextWindowType) *T.GdkWindow

	TextViewGetWindowType func(textView *T.GtkTextView,
		window *T.GdkWindow) T.GtkTextWindowType

	TextViewSetBorderWindowSize func(textView *T.GtkTextView,
		t T.GtkTextWindowType,
		size int)

	TextViewGetBorderWindowSize func(textView *T.GtkTextView,
		t T.GtkTextWindowType) int

	TextViewForwardDisplayLine func(textView *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	TextViewBackwardDisplayLine func(textView *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	TextViewForwardDisplayLineEnd func(textView *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	TextViewBackwardDisplayLineStart func(textView *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	TextViewStartsDisplayLine func(textView *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	TextViewMoveVisually func(textView *T.GtkTextView,
		iter *T.GtkTextIter,
		count int) T.Gboolean

	TextViewImContextFilterKeypress func(textView *T.GtkTextView,
		event *T.GdkEventKey) T.Gboolean

	TextViewResetImContext func(textView *T.GtkTextView)

	TextViewAddChildAtAnchor func(textView *T.GtkTextView,
		child *Widget,
		anchor *T.GtkTextChildAnchor)

	TextViewAddChildInWindow func(textView *T.GtkTextView,
		child *Widget,
		whichWindow T.GtkTextWindowType,
		xpos int,
		ypos int)

	TextViewMoveChild func(textView *T.GtkTextView,
		child *Widget,
		xpos int,
		ypos int)

	TextViewSetWrapMode func(textView *T.GtkTextView,
		wrapMode T.GtkWrapMode)

	TextViewGetWrapMode func(textView *T.GtkTextView) T.GtkWrapMode

	TextViewSetEditable func(textView *T.GtkTextView,
		setting T.Gboolean)

	TextViewGetEditable func(textView *T.GtkTextView) T.Gboolean

	TextViewSetOverwrite func(textView *T.GtkTextView,
		overwrite T.Gboolean)

	TextViewGetOverwrite func(textView *T.GtkTextView) T.Gboolean

	TextViewSetAcceptsTab func(textView *T.GtkTextView,
		acceptsTab T.Gboolean)

	TextViewGetAcceptsTab func(textView *T.GtkTextView) T.Gboolean

	TextViewSetPixelsAboveLines func(textView *T.GtkTextView,
		pixelsAboveLines int)

	TextViewGetPixelsAboveLines func(textView *T.GtkTextView) int

	TextViewSetPixelsBelowLines func(textView *T.GtkTextView,
		pixelsBelowLines int)

	TextViewGetPixelsBelowLines func(textView *T.GtkTextView) int

	TextViewSetPixelsInsideWrap func(textView *T.GtkTextView,
		pixelsInsideWrap int)

	TextViewGetPixelsInsideWrap func(textView *T.GtkTextView) int

	TextViewSetJustification func(textView *T.GtkTextView,
		justification T.GtkJustification)

	TextViewGetJustification func(textView *T.GtkTextView) T.GtkJustification

	TextViewSetLeftMargin func(textView *T.GtkTextView,
		leftMargin int)

	TextViewGetLeftMargin func(textView *T.GtkTextView) int

	TextViewSetRightMargin func(textView *T.GtkTextView,
		rightMargin int)

	TextViewGetRightMargin func(textView *T.GtkTextView) int

	TextViewSetIndent func(textView *T.GtkTextView,
		indent int)

	TextViewGetIndent func(textView *T.GtkTextView) int

	TextViewSetTabs func(textView *T.GtkTextView,
		tabs *T.PangoTabArray)

	TextViewGetTabs func(textView *T.GtkTextView) *T.PangoTabArray

	TextViewGetDefaultAttributes func(textView *T.GtkTextView) *T.GtkTextAttributes

	PixmapNew func(pixmap *T.GdkPixmap,
		mask *T.GdkBitmap) *Widget

	PixmapSet func(pixmap *T.GtkPixmap,
		val *T.GdkPixmap,
		mask *T.GdkBitmap)

	PixmapGet func(pixmap *T.GtkPixmap,
		val **T.GdkPixmap,
		mask **T.GdkBitmap)

	PixmapSetBuildInsensitive func(pixmap *T.GtkPixmap,
		build T.Gboolean)

	ToolbarInsert func(toolbar *T.GtkToolbar,
		item *T.GtkToolItem,
		pos int)

	ToolbarGetItemIndex func(toolbar *T.GtkToolbar,
		item *T.GtkToolItem) int

	ToolbarGetNItems func(toolbar *T.GtkToolbar) int

	ToolbarGetNthItem func(toolbar *T.GtkToolbar,
		n int) *T.GtkToolItem

	ToolbarGetShowArrow func(toolbar *T.GtkToolbar) T.Gboolean

	ToolbarSetShowArrow func(toolbar *T.GtkToolbar,
		showArrow T.Gboolean)

	ToolbarGetStyle func(toolbar *T.GtkToolbar) T.GtkToolbarStyle

	ToolbarSetStyle func(toolbar *T.GtkToolbar,
		style T.GtkToolbarStyle)

	ToolbarUnsetStyle func(toolbar *T.GtkToolbar)

	ToolbarGetIconSize func(toolbar *T.GtkToolbar) IconSize

	ToolbarSetIconSize func(toolbar *T.GtkToolbar,
		iconSize IconSize)

	ToolbarUnsetIconSize func(toolbar *T.GtkToolbar)

	ToolbarGetReliefStyle func(toolbar *T.GtkToolbar) T.GtkReliefStyle

	ToolbarGetDropIndex func(toolbar *T.GtkToolbar,
		x int,
		y int) int

	ToolbarSetDropHighlightItem func(toolbar *T.GtkToolbar,
		toolItem *T.GtkToolItem,
		index int)

	ToolbarGetOrientation func(toolbar *T.GtkToolbar) T.GtkOrientation

	ToolbarSetOrientation func(toolbar *T.GtkToolbar,
		orientation T.GtkOrientation)

	ToolbarGetTooltips func(toolbar *T.GtkToolbar) T.Gboolean

	ToolbarSetTooltips func(toolbar *T.GtkToolbar,
		enable T.Gboolean)

	ToolbarAppendItem func(toolbar *T.GtkToolbar,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *Widget,
		callback T.GCallback,
		userData T.Gpointer) *Widget

	ToolbarPrependItem func(toolbar *T.GtkToolbar,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *Widget,
		callback T.GCallback,
		userData T.Gpointer) *Widget

	ToolbarInsertItem func(toolbar *T.GtkToolbar,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *Widget,
		callback T.GCallback,
		userDataGpointer,
		position int) *Widget

	ToolbarInsertStock func(toolbar *T.GtkToolbar,
		stockId string,
		tooltipText string,
		tooltipPrivateText string,
		callback T.GCallback,
		userDataGpointer,
		position int) *Widget

	ToolbarAppendSpace func(toolbar *T.GtkToolbar)

	ToolbarPrependSpace func(toolbar *T.GtkToolbar)

	ToolbarInsertSpace func(toolbar *T.GtkToolbar,
		position int)

	ToolbarRemoveSpace func(toolbar *T.GtkToolbar,
		position int)

	ToolbarAppendElement func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *Widget,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *Widget,
		callback T.GCallback,
		userData T.Gpointer) *Widget

	ToolbarPrependElement func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *Widget,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *Widget,
		callback T.GCallback,
		userData T.Gpointer) *Widget

	ToolbarInsertElement func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *Widget,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *Widget,
		callback T.GCallback,
		userDataGpointer,
		position int) *Widget

	ToolbarAppendWidget func(toolbar *T.GtkToolbar,
		widget *Widget,
		tooltipText string,
		tooltipPrivateText string)

	ToolbarPrependWidget func(toolbar *T.GtkToolbar,
		widget *Widget,
		tooltipText string,
		tooltipPrivateText string)

	ToolbarInsertWidget func(toolbar *T.GtkToolbar,
		widget *Widget,
		tooltipText string,
		tooltipPrivateText string,
		position int)

	ToolItemGroupNew func(label string) *Widget

	ToolItemGroupSetLabel func(group *T.GtkToolItemGroup,
		label string)

	ToolItemGroupSetLabelWidget func(group *T.GtkToolItemGroup,
		labelWidget *Widget)

	ToolItemGroupSetCollapsed func(group *T.GtkToolItemGroup,
		collapsed T.Gboolean)

	ToolItemGroupSetEllipsize func(group *T.GtkToolItemGroup,
		ellipsize T.PangoEllipsizeMode)

	ToolItemGroupSetHeaderRelief func(group *T.GtkToolItemGroup,
		style T.GtkReliefStyle)

	ToolItemGroupGetLabel func(group *T.GtkToolItemGroup) string

	ToolItemGroupGetLabelWidget func(group *T.GtkToolItemGroup) *Widget

	ToolItemGroupGetCollapsed func(group *T.GtkToolItemGroup) T.Gboolean

	ToolItemGroupGetEllipsize func(group *T.GtkToolItemGroup) T.PangoEllipsizeMode

	ToolItemGroupGetHeaderRelief func(group *T.GtkToolItemGroup) T.GtkReliefStyle

	ToolItemGroupInsert func(group *T.GtkToolItemGroup,
		item *T.GtkToolItem,
		position int)

	ToolItemGroupSetItemPosition func(group *T.GtkToolItemGroup,
		item *T.GtkToolItem,
		position int)

	ToolItemGroupGetItemPosition func(group *T.GtkToolItemGroup,
		item *T.GtkToolItem) int

	ToolItemGroupGetNItems func(group *T.GtkToolItemGroup) uint

	ToolItemGroupGetNthItem func(group *T.GtkToolItemGroup,
		index uint) *T.GtkToolItem

	ToolItemGroupGetDropItem func(group *T.GtkToolItemGroup,
		x int,
		y int) *T.GtkToolItem

	ToolPaletteSetGroupPosition func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup,
		position int)

	ToolPaletteSetExclusive func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup,
		exclusive T.Gboolean)

	ToolPaletteSetExpand func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup,
		expand T.Gboolean)

	ToolPaletteGetGroupPosition func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup) int

	ToolPaletteGetExclusive func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup) T.Gboolean

	ToolPaletteGetExpand func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup) T.Gboolean

	ToolPaletteSetIconSize func(palette *T.GtkToolPalette,
		iconSize IconSize)

	ToolPaletteUnsetIconSize func(palette *T.GtkToolPalette)

	ToolPaletteSetStyle func(palette *T.GtkToolPalette,
		style T.GtkToolbarStyle)

	ToolPaletteUnsetStyle func(palette *T.GtkToolPalette)

	ToolPaletteGetIconSize func(palette *T.GtkToolPalette) IconSize

	ToolPaletteGetStyle func(palette *T.GtkToolPalette) T.GtkToolbarStyle

	ToolPaletteGetDropItem func(palette *T.GtkToolPalette,
		x int,
		y int) *T.GtkToolItem

	ToolPaletteGetDropGroup func(palette *T.GtkToolPalette,
		x int,
		y int) *T.GtkToolItemGroup

	ToolPaletteGetDragItem func(palette *T.GtkToolPalette,
		selection *T.GtkSelectionData) *Widget

	ToolPaletteSetDragSource func(palette *T.GtkToolPalette,
		targets T.GtkToolPaletteDragTargets)

	ToolPaletteAddDragDest func(palette *T.GtkToolPalette,
		widget *Widget,
		flags T.GtkDestDefaults,
		targets T.GtkToolPaletteDragTargets,
		actions T.GdkDragAction)

	ToolPaletteGetHadjustment func(palette *T.GtkToolPalette) *Adjustment

	ToolPaletteGetVadjustment func(palette *T.GtkToolPalette) *Adjustment

	ToolPaletteGetDragTargetItem func() *T.GtkTargetEntry

	ToolPaletteGetDragTargetGroup func() *T.GtkTargetEntry

	ToolShellGetIconSize func(shell *T.GtkToolShell) IconSize

	ToolShellGetOrientation func(shell *T.GtkToolShell) T.GtkOrientation

	ToolShellGetStyle func(shell *T.GtkToolShell) T.GtkToolbarStyle

	ToolShellGetReliefStyle func(shell *T.GtkToolShell) T.GtkReliefStyle

	ToolShellRebuildMenu func(shell *T.GtkToolShell)

	ToolShellGetTextOrientation func(shell *T.GtkToolShell) T.GtkOrientation

	ToolShellGetTextAlignment func(shell *T.GtkToolShell) float32

	ToolShellGetEllipsizeMode func(shell *T.GtkToolShell) T.PangoEllipsizeMode

	ToolShellGetTextSizeGroup func(shell *T.GtkToolShell) *T.GtkSizeGroup

	TestInit func(argcp *int, argvp ***T.Char, v ...VArg)

	TestRegisterAllTypes func()

	TestListAllTypes func(nTypes *uint) *T.GType

	TestFindWidget func(widget *Widget,
		labelPattern string,
		widgetType T.GType) *Widget

	TestCreateWidget func(widgetType T.GType,
		firstPropertyName string, v ...VArg) *Widget

	TestCreateSimpleWindow func(windowTitle string,
		dialogText string) *Widget

	TestDisplayButtonWindow func(windowTitle string,
		dialogText string, v ...VArg) *Widget

	TestSliderSetPerc func(widget *Widget,
		percentage float64)

	TestSliderGetValue func(widget *Widget) float64

	TestSpinButtonClick func(spinner *T.GtkSpinButton,
		button uint,
		upwards T.Gboolean) T.Gboolean

	TestWidgetClick func(widget *Widget,
		button uint,
		modifiers T.GdkModifierType) T.Gboolean

	TestWidgetSendKey func(widget *Widget,
		keyval uint,
		modifiers T.GdkModifierType) T.Gboolean

	TestTextSet func(widget *Widget,
		string string)

	TestTextGet func(widget *Widget) string

	TestFindSibling func(baseWidget *Widget,
		widgetType T.GType) *Widget

	TestFindLabel func(widget *Widget,
		labelPattern string) *Widget

	TreeDragSourceRowDraggable func(dragSource *T.GtkTreeDragSource,
		path *T.GtkTreePath) T.Gboolean

	TreeDragSourceDragDataDelete func(dragSource *T.GtkTreeDragSource,
		path *T.GtkTreePath) T.Gboolean

	TreeDragSourceDragDataGet func(dragSource *T.GtkTreeDragSource,
		path *T.GtkTreePath,
		selectionData *T.GtkSelectionData) T.Gboolean

	TreeDragDestDragDataReceived func(dragDest *T.GtkTreeDragDest,
		dest *T.GtkTreePath,
		selectionData *T.GtkSelectionData) T.Gboolean

	TreeDragDestRowDropPossible func(dragDest *T.GtkTreeDragDest,
		destPath *T.GtkTreePath,
		selectionData *T.GtkSelectionData) T.Gboolean

	TreeSetRowDragData func(selectionData *T.GtkSelectionData,
		treeModel *T.GtkTreeModel,
		path *T.GtkTreePath) T.Gboolean

	TreeGetRowDragData func(selectionData *T.GtkSelectionData,
		treeModel **T.GtkTreeModel,
		path **T.GtkTreePath) T.Gboolean

	TreeModelSortNewWithModel func(childModel *T.GtkTreeModel) *T.GtkTreeModel

	TreeModelSortGetModel func(treeModel *T.GtkTreeModelSort) *T.GtkTreeModel

	TreeModelSortConvertChildPathToPath func(treeModelSort *T.GtkTreeModelSort,
		childPath *T.GtkTreePath) *T.GtkTreePath

	TreeModelSortConvertChildIterToIter func(treeModelSort *T.GtkTreeModelSort,
		sortIter *T.GtkTreeIter,
		childIter *T.GtkTreeIter) T.Gboolean

	TreeModelSortConvertPathToChildPath func(treeModelSort *T.GtkTreeModelSort,
		sortedPath *T.GtkTreePath) *T.GtkTreePath

	TreeModelSortConvertIterToChildIter func(treeModelSort *T.GtkTreeModelSort,
		childIter *T.GtkTreeIter,
		sortedIter *T.GtkTreeIter)

	TreeModelSortResetDefaultSortFunc func(treeModelSort *T.GtkTreeModelSort)

	TreeModelSortClearCache func(treeModelSort *T.GtkTreeModelSort)

	TreeModelSortIterIsValid func(treeModelSort *T.GtkTreeModelSort,
		iter *T.GtkTreeIter) T.Gboolean

	TreeSelectionSetMode func(selection *T.GtkTreeSelection,
		typ T.GtkSelectionMode)

	TreeSelectionGetMode func(selection *T.GtkTreeSelection) T.GtkSelectionMode

	TreeSelectionSetSelectFunction func(selection *T.GtkTreeSelection,
		f T.GtkTreeSelectionFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	TreeSelectionGetUserData func(selection *T.GtkTreeSelection) T.Gpointer

	TreeSelectionGetTreeView func(selection *T.GtkTreeSelection) *T.GtkTreeView

	TreeSelectionGetSelectFunction func(selection *T.GtkTreeSelection) T.GtkTreeSelectionFunc

	TreeSelectionGetSelected func(selection *T.GtkTreeSelection,
		model **T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gboolean

	TreeSelectionGetSelectedRows func(selection *T.GtkTreeSelection,
		model **T.GtkTreeModel) *T.GList

	TreeSelectionCountSelectedRows func(selection *T.GtkTreeSelection) int

	TreeSelectionSelectedForeach func(selection *T.GtkTreeSelection,
		f T.GtkTreeSelectionForeachFunc,
		data T.Gpointer)

	TreeSelectionSelectPath func(selection *T.GtkTreeSelection,
		path *T.GtkTreePath)

	TreeSelectionUnselectPath func(selection *T.GtkTreeSelection,
		path *T.GtkTreePath)

	TreeSelectionSelectIter func(selection *T.GtkTreeSelection,
		iter *T.GtkTreeIter)

	TreeSelectionUnselectIter func(selection *T.GtkTreeSelection,
		iter *T.GtkTreeIter)

	TreeSelectionPathIsSelected func(selection *T.GtkTreeSelection,
		path *T.GtkTreePath) T.Gboolean

	TreeSelectionIterIsSelected func(selection *T.GtkTreeSelection,
		iter *T.GtkTreeIter) T.Gboolean

	TreeSelectionSelectAll func(selection *T.GtkTreeSelection)

	TreeSelectionUnselectAll func(selection *T.GtkTreeSelection)

	TreeSelectionSelectRange func(selection *T.GtkTreeSelection,
		startPath *T.GtkTreePath,
		endPath *T.GtkTreePath)

	TreeSelectionUnselectRange func(selection *T.GtkTreeSelection,
		startPath *T.GtkTreePath,
		endPath *T.GtkTreePath)

	TreeStoreNew func(
		nColumns int, v ...VArg) *T.GtkTreeStore

	TreeStoreNewv func(nColumns int,
		types *T.GType) *T.GtkTreeStore

	TreeStoreSetColumnTypes func(treeStore *T.GtkTreeStore,
		nColumns int,
		types *T.GType)

	TreeStoreSetValue func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		column int,
		value *T.GValue)

	TreeStoreSet func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter, v ...VArg)

	TreeStoreSetValuesv func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		columns *int,
		values *T.GValue,
		nValues int)

	TreeStoreSetValist func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		varArgs T.VaList)

	TreeStoreRemove func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter) T.Gboolean

	TreeStoreInsert func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		position int)

	TreeStoreInsertBefore func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		sibling *T.GtkTreeIter)

	TreeStoreInsertAfter func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		sibling *T.GtkTreeIter)

	TreeStoreInsertWithValues func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		position int, v ...VArg)

	TreeStoreInsertWithValuesv func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		position int,
		columns *int,
		values *T.GValue,
		nValues int)

	TreeStorePrepend func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter)

	TreeStoreAppend func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter)

	TreeStoreIsAncestor func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		descendant *T.GtkTreeIter) T.Gboolean

	TreeStoreIterDepth func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter) int

	TreeStoreClear func(treeStore *T.GtkTreeStore)

	TreeStoreIterIsValid func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter) T.Gboolean

	TreeStoreReorder func(treeStore *T.GtkTreeStore,
		parent *T.GtkTreeIter,
		newOrder *int)

	TreeStoreSwap func(treeStore *T.GtkTreeStore,
		a *T.GtkTreeIter,
		b *T.GtkTreeIter)

	TreeStoreMoveBefore func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		position *T.GtkTreeIter)

	TreeStoreMoveAfter func(treeStore *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		position *T.GtkTreeIter)

	UiManagerNew func() *T.GtkUIManager

	UiManagerSetAddTearoffs func(self *T.GtkUIManager,
		addTearoffs T.Gboolean)

	UiManagerGetAddTearoffs func(self *T.GtkUIManager) T.Gboolean

	UiManagerInsertActionGroup func(self *T.GtkUIManager,
		actionGroup *ActionGroup,
		pos int)

	UiManagerRemoveActionGroup func(self *T.GtkUIManager,
		actionGroup *ActionGroup)

	UiManagerGetActionGroups func(self *T.GtkUIManager) *T.GList

	UiManagerGetAccelGroup func(self *T.GtkUIManager) *AccelGroup

	UiManagerGetWidget func(self *T.GtkUIManager,
		path string) *Widget

	UiManagerGetToplevels func(self *T.GtkUIManager,
		types T.GtkUIManagerItemType) *T.GSList

	UiManagerGetAction func(self *T.GtkUIManager,
		path string) *Action

	UiManagerAddUiFromString func(self *T.GtkUIManager,
		buffer string,
		length T.Gssize,
		err **T.GError) uint

	UiManagerAddUiFromFile func(self *T.GtkUIManager,
		filename string,
		err **T.GError) uint

	UiManagerAddUi func(self *T.GtkUIManager,
		mergeId uint,
		path string,
		name string,
		action string,
		t T.GtkUIManagerItemType,
		top T.Gboolean)

	UiManagerRemoveUi func(self *T.GtkUIManager,
		mergeId uint)

	UiManagerGetUi func(self *T.GtkUIManager) string

	UiManagerEnsureUpdate func(self *T.GtkUIManager)

	UiManagerNewMergeId func(self *T.GtkUIManager) uint

	VbuttonBoxGetSpacingDefault func() int

	VbuttonBoxSetSpacingDefault func(spacing int)

	VbuttonBoxGetLayoutDefault func() ButtonBoxStyle

	VbuttonBoxSetLayoutDefault func(layout ButtonBoxStyle)

	VscaleNew func(adjustment *Adjustment) *Widget

	VscaleNewWithRange func(min float64,
		max float64,
		step float64) *Widget

	OldEditableClaimSelection func(oldEditable *T.GtkOldEditable,
		claim T.Gboolean,
		time T.GUint32)

	OldEditableChanged func(oldEditable *T.GtkOldEditable)

	OptionMenuGetMenu func(optionMenu *T.GtkOptionMenu) *Widget

	OptionMenuSetMenu func(optionMenu *T.GtkOptionMenu,
		menu *Widget)

	OptionMenuRemoveMenu func(optionMenu *T.GtkOptionMenu)

	OptionMenuGetHistory func(optionMenu *T.GtkOptionMenu) int

	OptionMenuSetHistory func(optionMenu *T.GtkOptionMenu,
		index uint)

	PreviewUninit func()

	PreviewNew func(t T.GtkPreviewType) *Widget

	PreviewSize func(preview *T.GtkPreview,
		width int,
		height int)

	PreviewPut func(preview *T.GtkPreview,
		window *T.GdkWindow,
		gc *T.GdkGC,
		srcx int,
		srcy int,
		destx int,
		desty int,
		width int,
		height int)

	PreviewDrawRow func(preview *T.GtkPreview,
		data *T.Guchar,
		x int,
		y int,
		w int)

	PreviewSetExpand func(preview *T.GtkPreview,
		expand T.Gboolean)

	PreviewSetGamma func(gamma float64)

	PreviewSetColorCube func(nredShades uint,
		ngreenShades uint,
		nblueShades uint,
		ngrayShades uint)

	PreviewSetInstallCmap func(installCmap int)

	PreviewSetReserved func(nreserved int)

	PreviewSetDither func(preview *T.GtkPreview,
		dither T.GdkRgbDither)

	PreviewGetVisual func() *T.GdkVisual

	PreviewGetCmap func() *T.GdkColormap

	PreviewGetInfo func() *T.GtkPreviewInfo

	PreviewReset func()

	TipsQueryStartQuery func(tipsQuery *T.GtkTipsQuery)

	TipsQueryStopQuery func(tipsQuery *T.GtkTipsQuery)

	TipsQuerySetCaller func(tipsQuery *T.GtkTipsQuery,
		caller *Widget)

	TipsQuerySetLabels func(tipsQuery *T.GtkTipsQuery,
		labelInactive string,
		labelNoTip string)

	Marshal_BOOLEAN__VOID func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_POINTER_INT_INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_INT_INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_INT_INT_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_STRING_STRING_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_ENUM__ENUM func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_INT__POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_INT__POINTER_CHAR_CHAR func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__ENUM_FLOAT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__ENUM_FLOAT_BOOLEAN func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__INT_INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__INT_INT_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_STRING_STRING func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT_ENUM func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER_UINT_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_INT_INT_POINTER_UINT_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__STRING_INT_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__UINT_POINTER_UINT_ENUM_ENUM_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__UINT_POINTER_UINT_UINT_ENUM func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	Marshal_VOID__UINT_STRING func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint T.Gpointer,
		marshalData T.Gpointer)

	PrintContextSetCairoContext func(
		context *T.GtkPrintContext,
		cr *T.Cairo,
		dpiX, dpiY float64)

	TreeAppend func(tree *T.GtkTree, treeItem *Widget)

	TreePrepend func(tree *T.GtkTree, treeItem *Widget)

	TreeInsert func(
		tree *T.GtkTree, treeItem *Widget, position int)

	TreeRemoveItems func(tree *T.GtkTree, items *T.GList)

	TreeClearItems func(
		tree *T.GtkTree, start int, end int)

	TreeSelectItem func(tree *T.GtkTree, item int)

	TreeUnselectItem func(tree *T.GtkTree, item int)

	TreeSelectChild func(
		tree *T.GtkTree, treeItem *Widget)

	TreeUnselectChild func(
		tree *T.GtkTree, treeItem *Widget)

	TreeChildPosition func(
		tree *T.GtkTree, child *Widget) int

	TreeSetSelectionMode func(
		tree *T.GtkTree, mode T.GtkSelectionMode)

	TreeSetViewMode func(
		tree *T.GtkTree, mode T.GtkTreeViewMode)

	TreeSetViewLines func(tree *T.GtkTree, flag T.Gboolean)

	TreeRemoveItem func(tree *T.GtkTree, child *Widget)

	TreeItemNewWithLabel func(label string) *Widget

	TreeItemSetSubtree func(
		treeItem *T.GtkTreeItem, subtree *Widget)

	TreeItemRemoveSubtree func(treeItem *T.GtkTreeItem)

	TreeItemSelect func(treeItem *T.GtkTreeItem)

	TreeItemDeselect func(treeItem *T.GtkTreeItem)

	TreeItemExpand func(treeItem *T.GtkTreeItem)

	TreeItemCollapse func(treeItem *T.GtkTreeItem)

	TypeEnumGetValues func(
		enumType T.GtkType) *T.GtkEnumValue

	TypeFlagsGetValues func(
		flagsType T.GtkType) *T.GtkFlagValue

	TypeEnumFindValue func(
		enumType T.GtkType,
		valueName string) *T.GtkEnumValue

	TypeFlagsFindValue func(
		flagsType T.GtkType,

		valueName string) *T.GtkFlagValue

	// misc

	TreeViewColumnNew func() *TreeViewColumn

	// funcs with methods

	TreeViewColumnNewWithAttributes func(
		title string, cell *CellRenderer,
		v ...VArg) *TreeViewColumn

	TreeViewColumnPackStart func(
		treeColumn *TreeViewColumn,
		cell *CellRenderer,
		expand T.Gboolean)

	TreeViewColumnPackEnd func(
		treeColumn *TreeViewColumn,
		cell *CellRenderer,
		expand T.Gboolean)

	TreeViewColumnClear func(
		treeColumn *TreeViewColumn)

	TreeViewColumnGetCellRenderers func(
		treeColumn *TreeViewColumn) *T.GList

	TreeViewColumnAddAttribute func(
		treeColumn *TreeViewColumn,
		cellRenderer *CellRenderer,
		attribute string,
		column int)

	TreeViewColumnSetAttributes func(
		treeColumn *TreeViewColumn,
		cellRenderer *CellRenderer, v ...VArg)

	TreeViewColumnSetCellDataFunc func(
		treeColumn *TreeViewColumn,
		cellRenderer *CellRenderer,
		f T.GtkTreeCellDataFunc,
		funcDataGpointer,
		destroy T.GDestroyNotify)

	TreeViewColumnClearAttributes func(
		treeColumn *TreeViewColumn,
		cellRenderer *CellRenderer)

	TreeViewColumnSetSpacing func(
		treeColumn *TreeViewColumn,
		spacing int)

	TreeViewColumnGetSpacing func(
		treeColumn *TreeViewColumn) int

	TreeViewColumnSetVisible func(
		treeColumn *TreeViewColumn,
		visible T.Gboolean)

	TreeViewColumnGetVisible func(
		treeColumn *TreeViewColumn) T.Gboolean

	TreeViewColumnSetResizable func(
		treeColumn *TreeViewColumn,
		resizable T.Gboolean)

	TreeViewColumnGetResizable func(
		treeColumn *TreeViewColumn) T.Gboolean

	TreeViewColumnSetSizing func(
		treeColumn *TreeViewColumn,
		typ T.GtkTreeViewColumnSizing)

	TreeViewColumnGetSizing func(
		treeColumn *TreeViewColumn) T.GtkTreeViewColumnSizing

	TreeViewColumnGetWidth func(
		treeColumn *TreeViewColumn) int

	TreeViewColumnGetFixedWidth func(
		treeColumn *TreeViewColumn) int

	TreeViewColumnSetFixedWidth func(
		treeColumn *TreeViewColumn,
		fixedWidth int)

	TreeViewColumnSetMinWidth func(
		treeColumn *TreeViewColumn,
		minWidth int)

	TreeViewColumnGetMinWidth func(
		treeColumn *TreeViewColumn) int

	TreeViewColumnSetMaxWidth func(
		treeColumn *TreeViewColumn,
		maxWidth int)

	TreeViewColumnGetMaxWidth func(
		treeColumn *TreeViewColumn) int

	TreeViewColumnClicked func(
		treeColumn *TreeViewColumn)

	TreeViewColumnSetTitle func(
		treeColumn *TreeViewColumn,
		title string)

	TreeViewColumnGetTitle func(
		treeColumn *TreeViewColumn) string

	TreeViewColumnSetExpand func(
		treeColumn *TreeViewColumn,
		expand T.Gboolean)

	TreeViewColumnGetExpand func(
		treeColumn *TreeViewColumn) T.Gboolean

	TreeViewColumnSetClickable func(
		treeColumn *TreeViewColumn,
		clickable T.Gboolean)

	TreeViewColumnGetClickable func(
		treeColumn *TreeViewColumn) T.Gboolean

	TreeViewColumnSetWidget func(
		treeColumn *TreeViewColumn,
		widget *Widget)

	TreeViewColumnGetWidget func(
		treeColumn *TreeViewColumn) *Widget

	TreeViewColumnSetAlignment func(
		treeColumn *TreeViewColumn,
		xalign float32)

	TreeViewColumnGetAlignment func(
		treeColumn *TreeViewColumn) float32

	TreeViewColumnSetReorderable func(
		treeColumn *TreeViewColumn,
		reorderable T.Gboolean)

	TreeViewColumnGetReorderable func(
		treeColumn *TreeViewColumn) T.Gboolean

	TreeViewColumnSetSortColumnId func(
		treeColumn *TreeViewColumn,
		sortColumnId int)

	TreeViewColumnGetSortColumnId func(
		treeColumn *TreeViewColumn) int

	TreeViewColumnSetSortIndicator func(
		treeColumn *TreeViewColumn,
		setting T.Gboolean)

	TreeViewColumnGetSortIndicator func(
		treeColumn *TreeViewColumn) T.Gboolean

	TreeViewColumnSetSortOrder func(
		treeColumn *TreeViewColumn,
		order T.GtkSortType)

	TreeViewColumnGetSortOrder func(
		treeColumn *TreeViewColumn) T.GtkSortType

	TreeViewColumnCellSetCellData func(
		treeColumn *TreeViewColumn,
		treeModel *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		isExpander T.Gboolean,
		isExpanded T.Gboolean)

	TreeViewColumnCellGetSize func(
		treeColumn *TreeViewColumn,
		cellArea *T.GdkRectangle,
		xOffset *int,
		yOffset *int,
		width *int,
		height *int)

	TreeViewColumnCellIsVisible func(
		treeColumn *TreeViewColumn) T.Gboolean

	TreeViewColumnFocusCell func(
		treeColumn *TreeViewColumn,
		cell *CellRenderer)

	TreeViewColumnCellGetPosition func(
		treeColumn *TreeViewColumn,
		cellRenderer *CellRenderer,
		startPos *int,
		width *int) T.Gboolean

	TreeViewColumnQueueResize func(
		treeColumn *TreeViewColumn)

	TreeViewColumnGetTreeView func(
		treeColumn *TreeViewColumn) *Widget
)

type TreeViewColumn struct{}

func (tc *TreeViewColumn) PackStart(
	cell *CellRenderer, expand T.Gboolean) {
	TreeViewColumnPackEnd(tc, cell, expand)
}

func (tc *TreeViewColumn) Clear() { TreeViewColumnClear(tc) }

func (tc *TreeViewColumn) GetCellRenderers() *T.GList {
	return TreeViewColumnGetCellRenderers(tc)
}

func (tc *TreeViewColumn) AddAttribute(
	cellRenderer *CellRenderer,
	attribute string, column int) {
	TreeViewColumnAddAttribute(
		tc, cellRenderer, attribute, column)
}

func (tc *TreeViewColumn) SetAttributes(
	cellRenderer *CellRenderer, v ...VArg) {
	TreeViewColumnSetAttributes(tc, cellRenderer, v)
}

func (tc *TreeViewColumn) SetCellDataFunc(
	cellRenderer *CellRenderer, f T.GtkTreeCellDataFunc,
	funcDataGpointer, destroy T.GDestroyNotify) {
	TreeViewColumnSetCellDataFunc(
		tc, cellRenderer, f, funcDataGpointer, destroy)
}

func (tc *TreeViewColumn) ClearAttributes(
	cellRenderer *CellRenderer) {
	TreeViewColumnClearAttributes(tc, cellRenderer)
}

func (tc *TreeViewColumn) SetSpacing(spacing int) {
	TreeViewColumnSetSpacing(tc, spacing)
}

func (tc *TreeViewColumn) GetSpacing() int {
	return TreeViewColumnGetSpacing(tc)
}

func (tc *TreeViewColumn) SetVisible(visible T.Gboolean) {
	TreeViewColumnSetVisible(tc, visible)
}

func (tc *TreeViewColumn) GetVisible() T.Gboolean {
	return TreeViewColumnGetVisible(tc)
}

func (tc *TreeViewColumn) SetResizable(resizable T.Gboolean) {
	TreeViewColumnSetResizable(tc, resizable)
}

func (tc *TreeViewColumn) GetResizable() T.Gboolean {
	return TreeViewColumnGetResizable(tc)
}

func (tc *TreeViewColumn) SetSizing(typ T.GtkTreeViewColumnSizing) {
	TreeViewColumnSetSizing(tc, typ)
}

func (tc *TreeViewColumn) GetSizing() T.GtkTreeViewColumnSizing {
	return TreeViewColumnGetSizing(tc)
}

func (tc *TreeViewColumn) GetWidth() int {
	return TreeViewColumnGetWidth(tc)
}

func (tc *TreeViewColumn) GetFixedWidth() int {
	return TreeViewColumnGetFixedWidth(tc)
}

func (tc *TreeViewColumn) SetFixedWidth(fixedWidth int) {
	TreeViewColumnSetFixedWidth(tc, fixedWidth)
}

func (tc *TreeViewColumn) SetMinWidth(minWidth int) {
	TreeViewColumnSetMinWidth(tc, minWidth)
}

func (tc *TreeViewColumn) GetMinWidth() int {
	return TreeViewColumnGetMinWidth(tc)
}

func (tc *TreeViewColumn) SetMaxWidth(maxWidth int) {
	TreeViewColumnSetMaxWidth(tc, maxWidth)
}

func (tc *TreeViewColumn) GetMaxWidth() int {
	return TreeViewColumnGetMaxWidth(tc)
}

func (tc *TreeViewColumn) Clicked() {
	TreeViewColumnClicked(tc)
}

func (tc *TreeViewColumn) SetTitle(title string) {
	TreeViewColumnSetTitle(tc, title)
}

func (tc *TreeViewColumn) GetTitle() string {
	return TreeViewColumnGetTitle(tc)
}

func (tc *TreeViewColumn) SetExpand(expand T.Gboolean) {
	TreeViewColumnSetExpand(tc, expand)
}

func (tc *TreeViewColumn) GetExpand() T.Gboolean {
	return TreeViewColumnGetExpand(tc)
}

func (tc *TreeViewColumn) SetClickable(clickable T.Gboolean) {
	TreeViewColumnSetClickable(tc, clickable)
}

func (tc *TreeViewColumn) GetClickable() T.Gboolean {
	return TreeViewColumnGetClickable(tc)
}

func (tc *TreeViewColumn) SetWidget(widget *Widget) {
	TreeViewColumnSetWidget(tc, widget)
}

func (tc *TreeViewColumn) GetWidget() *Widget {
	return TreeViewColumnGetWidget(tc)
}

func (tc *TreeViewColumn) SetAlignment(xalign float32) {
	TreeViewColumnSetAlignment(tc, xalign)
}

func (tc *TreeViewColumn) GetAlignment() float32 {
	return TreeViewColumnGetAlignment(tc)
}

func (tc *TreeViewColumn) SetReorderable(reorderable T.Gboolean) {
	TreeViewColumnSetReorderable(tc, reorderable)
}

func (tc *TreeViewColumn) GetReorderable() T.Gboolean {
	return TreeViewColumnGetReorderable(tc)
}

func (tc *TreeViewColumn) SetSortColumnId(sortColumnId int) {
	TreeViewColumnSetSortColumnId(tc, sortColumnId)
}

func (tc *TreeViewColumn) GetSortColumnId() int {
	return TreeViewColumnGetSortColumnId(tc)
}

func (tc *TreeViewColumn) SetSortIndicator(setting T.Gboolean) {
	TreeViewColumnSetSortIndicator(tc, setting)
}

func (tc *TreeViewColumn) GetSortIndicator() T.Gboolean {
	return TreeViewColumnGetSortIndicator(tc)
}

func (tc *TreeViewColumn) SetSortOrder(order T.GtkSortType) {
	TreeViewColumnSetSortOrder(tc, order)
}

func (tc *TreeViewColumn) GetSortOrder() T.GtkSortType {
	return TreeViewColumnGetSortOrder(tc)
}

func (tc *TreeViewColumn) CellSetCellData(
	treeModel *T.GtkTreeModel, iter *T.GtkTreeIter,
	isExpander, isExpanded T.Gboolean) {
	TreeViewColumnCellSetCellData(
		tc, treeModel, iter, isExpander, isExpanded)
}

func (tc *TreeViewColumn) CellGetSize(cellArea *T.GdkRectangle,
	xOffset, yOffset, width, height *int) {
	TreeViewColumnCellGetSize(
		tc, cellArea, xOffset, yOffset, width, height)
}

func (tc *TreeViewColumn) CellIsVisible() T.Gboolean {
	return TreeViewColumnCellIsVisible(tc)
}

func (tc *TreeViewColumn) FocusCell(cell *CellRenderer) {
	TreeViewColumnFocusCell(tc, cell)
}

func (tc *TreeViewColumn) CellGetPosition(
	cellRenderer *CellRenderer, startPos *int,
	width *int) T.Gboolean {
	return TreeViewColumnCellGetPosition(
		tc, cellRenderer, startPos, width)
}

func (tc *TreeViewColumn) QueueResize() {
	TreeViewColumnQueueResize(tc)
}

func (tc *TreeViewColumn) GetTreeView() *Widget {
	return TreeViewColumnGetTreeView(tc)
}
