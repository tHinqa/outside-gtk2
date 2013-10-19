// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gtk provides API definitions for accessing
//libgtk-win32-2.0-0.dll.
package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

var (
	AccelFlagsGetType              func() T.GType
	AccessibleGetType              func() T.GType
	ActionGetType                  func() T.GType
	ActionGroupGetType             func() T.GType
	ActivatableGetType             func() T.GType
	AdjustmentGetType              func() T.GType
	AlignmentGetType               func() T.GType
	AnchorTypeGetType              func() T.GType
	ArgFlagsGetType                func() T.GType
	ArrowGetType                   func() T.GType
	ArrowPlacementGetType          func() T.GType
	ArrowTypeGetType               func() T.GType
	AspectFrameGetType             func() T.GType
	AttachOptionsGetType           func() T.GType
	BinGetType                     func() T.GType
	BorderGetType                  func() T.GType
	BoxGetType                     func() T.GType
	BuildableGetType               func() T.GType
	BuilderErrorGetType            func() T.GType
	BuilderGetType                 func() T.GType
	CalendarDisplayOptionsGetType  func() T.GType
	CalendarGetType                func() T.GType
	CalendarNew                    func() *T.GtkWidget
	CellEditableGetType            func() T.GType
	CellLayoutGetType              func() T.GType
	CellRendererAccelGetType       func() T.GType
	CellRendererAccelModeGetType   func() T.GType
	CellRendererComboGetType       func() T.GType
	CellRendererGetType            func() T.GType
	CellRendererModeGetType        func() T.GType
	CellRendererPixbufGetType      func() T.GType
	CellRendererProgressGetType    func() T.GType
	CellRendererSpinGetType        func() T.GType
	CellRendererSpinnerGetType     func() T.GType
	CellRendererStateGetType       func() T.GType
	CellRendererTextGetType        func() T.GType
	CellRendererToggleGetType      func() T.GType
	CellTypeGetType                func() T.GType
	CellViewGetType                func() T.GType
	CellViewNew                    func() *T.GtkWidget
	CheckButtonGetType             func() T.GType
	CheckButtonNew                 func() *T.GtkWidget
	CheckMenuItemGetType           func() T.GType
	CheckMenuItemNew               func() *T.GtkWidget
	ClipboardGetType               func() T.GType
	ClistDragPosGetType            func() T.GType
	ClistGetType                   func() T.GType
	ColorButtonGetType             func() T.GType
	ColorButtonNew                 func() *T.GtkWidget
	ColorSelectionDialogGetType    func() T.GType
	ColorSelectionGetType          func() T.GType
	ColorSelectionNew              func() *T.GtkWidget
	ComboGetType                   func() T.GType
	ComboNew                       func() *T.GtkWidget
	ContainerGetType               func() T.GType
	CornerTypeGetType              func() T.GType
	CtreeExpanderStyleGetType      func() T.GType
	CtreeExpansionTypeGetType      func() T.GType
	CtreeGetType                   func() T.GType
	CtreeLineStyleGetType          func() T.GType
	CtreeNodeGetType               func() T.GType
	CtreePosGetType                func() T.GType
	CurveGetType                   func() T.GType
	CurveNew                       func() *T.GtkWidget
	CurveTypeGetType               func() T.GType
	DebugFlagGetType               func() T.GType
	DeleteTypeGetType              func() T.GType
	DestDefaultsGetType            func() T.GType
	DialogFlagsGetType             func() T.GType
	DialogGetType                  func() T.GType
	DialogNew                      func() *T.GtkWidget
	DirectionTypeGetType           func() T.GType
	DragResultGetType              func() T.GType
	DrawingAreaGetType             func() T.GType
	DrawingAreaNew                 func() *T.GtkWidget
	EditableGetType                func() T.GType
	EntryBufferGetType             func() T.GType
	EntryCompletionGetType         func() T.GType
	EntryGetType                   func() T.GType
	EntryIconPositionGetType       func() T.GType
	EntryNew                       func() *T.GtkWidget
	EventBoxGetType                func() T.GType
	EventBoxNew                    func() *T.GtkWidget
	ExpanderGetType                func() T.GType
	ExpanderStyleGetType           func() T.GType
	FileChooserActionGetType       func() T.GType
	FileChooserButtonGetType       func() T.GType
	FileChooserConfirmationGetType func() T.GType
	FileChooserDialogGetType       func() T.GType
	FileChooserErrorGetType        func() T.GType
	FileChooserGetType             func() T.GType
	FileChooserWidgetGetType       func() T.GType
	FileFilterFlagsGetType         func() T.GType
	FileFilterGetType              func() T.GType
	FileSelectionGetType           func() T.GType
	FixedGetType                   func() T.GType
	FixedNew                       func() *T.GtkWidget
	FontButtonGetType              func() T.GType
	FontButtonNew                  func() *T.GtkWidget
	FontSelectionDialogGetType     func() T.GType
	FontSelectionGetType           func() T.GType
	FontSelectionNew               func() *T.GtkWidget
	FrameGetType                   func() T.GType
	GammaCurveGetType              func() T.GType
	GammaCurveNew                  func() *T.GtkWidget
	HandleBoxGetType               func() T.GType
	HandleBoxNew                   func() *T.GtkWidget
	HboxGetType                    func() T.GType
	HbuttonBoxGetType              func() T.GType
	HbuttonBoxNew                  func() *T.GtkWidget
	HpanedGetType                  func() T.GType
	HpanedNew                      func() *T.GtkWidget
	HrulerGetType                  func() T.GType
	HrulerNew                      func() *T.GtkWidget
	HscaleGetType                  func() T.GType
	HscrollbarGetType              func() T.GType
	HseparatorGetType              func() T.GType
	HseparatorNew                  func() *T.GtkWidget
	HsvGetType                     func() T.GType
	HsvNew                         func() *T.GtkWidget
	IconFactoryGetType             func() T.GType
	IconInfoGetType                func() T.GType
	IconLookupFlagsGetType         func() T.GType
	IconSetGetType                 func() T.GType
	IconSizeGetType                func() T.GType
	IconSourceGetType              func() T.GType
	IconThemeErrorGetType          func() T.GType
	IconThemeGetType               func() T.GType
	IconViewDropPositionGetType    func() T.GType
	IconViewGetType                func() T.GType
	IconViewNew                    func() *T.GtkWidget
	IdentifierGetType              func() T.GType
	ImageGetType                   func() T.GType
	ImageMenuItemGetType           func() T.GType
	ImageMenuItemNew               func() *T.GtkWidget
	ImageNew                       func() *T.GtkWidget
	ImageTypeGetType               func() T.GType
	ImContextGetType               func() T.GType
	ImContextSimpleGetType         func() T.GType
	ImMulticontextGetType          func() T.GType
	ImPreeditStyleGetType          func() T.GType
	ImStatusStyleGetType           func() T.GType
	InfoBarGetType                 func() T.GType
	InfoBarNew                     func() *T.GtkWidget
	InputDialogGetType             func() T.GType
	InputDialogNew                 func() *T.GtkWidget
	InvisibleGetType               func() T.GType
	InvisibleNew                   func() *T.GtkWidget
	ItemFactoryGetType             func() T.GType
	ItemGetType                    func() T.GType
	JustificationGetType           func() T.GType
	LabelGetType                   func() T.GType
	LayoutGetType                  func() T.GType
	LinkButtonGetType              func() T.GType
	ListGetType                    func() T.GType
	ListItemGetType                func() T.GType
	ListItemNew                    func() *T.GtkWidget
	ListNew                        func() *T.GtkWidget
	ListStoreGetType               func() T.GType
	MatchTypeGetType               func() T.GType
	MenuBarGetType                 func() T.GType
	MenuBarNew                     func() *T.GtkWidget
	MenuDirectionTypeGetType       func() T.GType
	MenuGetType                    func() T.GType
	MenuItemGetType                func() T.GType
	MenuItemNew                    func() *T.GtkWidget
	MenuNew                        func() *T.GtkWidget
	MenuShellGetType               func() T.GType
	MenuToolButtonGetType          func() T.GType
	MessageDialogGetType           func() T.GType
	MessageTypeGetType             func() T.GType
	MetricTypeGetType              func() T.GType
	MiscGetType                    func() T.GType
	MountOperationGetType          func() T.GType
	MovementStepGetType            func() T.GType
	NotebookGetType                func() T.GType
	NotebookNew                    func() *T.GtkWidget
	NotebookTabGetType             func() T.GType
	NumberUpLayoutGetType          func() T.GType
	ObjectFlagsGetType             func() T.GType
	ObjectGetType                  func() T.GType
	OffscreenWindowGetType         func() T.GType
	OffscreenWindowNew             func() *T.GtkWidget
	OldEditableGetType             func() T.GType
	OptionMenuGetType              func() T.GType
	OptionMenuNew                  func() *T.GtkWidget
	OrientableGetType              func() T.GType
	OrientationGetType             func() T.GType
	PackDirectionGetType           func() T.GType
	PackTypeGetType                func() T.GType
	PageOrientationGetType         func() T.GType
	PageSetGetType                 func() T.GType
	PageSetupGetType               func() T.GType
	PanedGetType                   func() T.GType
	PaperSizeGetType               func() T.GType
	PathPriorityTypeGetType        func() T.GType
	PathTypeGetType                func() T.GType
	PixmapGetType                  func() T.GType
	PlugGetType                    func() T.GType
	PolicyTypeGetType              func() T.GType
	PositionTypeGetType            func() T.GType
	PreviewGetType                 func() T.GType
	PreviewTypeGetType             func() T.GType
	PrintContextGetType            func() T.GType
	PrintDuplexGetType             func() T.GType
	PrintErrorGetType              func() T.GType
	PrintOperationActionGetType    func() T.GType
	PrintOperationGetType          func() T.GType
	PrintOperationPreviewGetType   func() T.GType
	PrintOperationResultGetType    func() T.GType
	PrintPagesGetType              func() T.GType
	PrintQualityGetType            func() T.GType
	PrintSettingsGetType           func() T.GType
	PrintStatusGetType             func() T.GType
	PrivateFlagsGetType            func() T.GType
	ProgressBarGetType             func() T.GType
	ProgressBarNew                 func() *T.GtkWidget
	ProgressBarOrientationGetType  func() T.GType
	ProgressBarStyleGetType        func() T.GType
	ProgressGetType                func() T.GType
	RadioActionGetType             func() T.GType
	RadioButtonGetType             func() T.GType
	RadioMenuItemGetType           func() T.GType
	RadioToolButtonGetType         func() T.GType
	RangeGetType                   func() T.GType
	RcFlagsGetType                 func() T.GType
	RcStyleGetType                 func() T.GType
	RcTokenTypeGetType             func() T.GType
	RecentActionGetType            func() T.GType
	RecentChooserDialogGetType     func() T.GType
	RecentChooserErrorGetType      func() T.GType
	RecentChooserGetType           func() T.GType
	RecentChooserMenuGetType       func() T.GType
	RecentChooserMenuNew           func() *T.GtkWidget
	RecentChooserWidgetGetType     func() T.GType
	RecentChooserWidgetNew         func() *T.GtkWidget
	RecentFilterFlagsGetType       func() T.GType
	RecentFilterGetType            func() T.GType
	RecentInfoGetType              func() T.GType
	RecentManagerErrorGetType      func() T.GType
	RecentManagerGetType           func() T.GType
	RecentSortTypeGetType          func() T.GType
	ReliefStyleGetType             func() T.GType
	RequisitionGetType             func() T.GType
	ResizeModeGetType              func() T.GType
	ResponseTypeGetType            func() T.GType
	RulerGetType                   func() T.GType
	ScaleButtonGetType             func() T.GType
	ScaleGetType                   func() T.GType
	ScrollbarGetType               func() T.GType
	ScrolledWindowGetType          func() T.GType
	ScrollStepGetType              func() T.GType
	ScrollTypeGetType              func() T.GType
	SelectionDataGetType           func() T.GType
	SelectionModeGetType           func() T.GType
	SensitivityTypeGetType         func() T.GType
	SeparatorGetType               func() T.GType
	SeparatorMenuItemGetType       func() T.GType
	SeparatorMenuItemNew           func() *T.GtkWidget
	SeparatorToolItemGetType       func() T.GType
	SettingsGetType                func() T.GType
	ShadowTypeGetType              func() T.GType
	SideTypeGetType                func() T.GType
	SignalRunTypeGetType           func() T.GType
	SizeGroupGetType               func() T.GType
	SizeGroupModeGetType           func() T.GType
	SocketGetType                  func() T.GType
	SocketNew                      func() *T.GtkWidget
	SortTypeGetType                func() T.GType
	SpinButtonGetType              func() T.GType
	SpinButtonUpdatePolicyGetType  func() T.GType
	SpinnerGetType                 func() T.GType
	SpinnerNew                     func() *T.GtkWidget
	SpinTypeGetType                func() T.GType
	StateTypeGetType               func() T.GType
	StatusbarGetType               func() T.GType
	StatusbarNew                   func() *T.GtkWidget
	StatusIconGetType              func() T.GType
	StyleGetType                   func() T.GType
	SubmenuDirectionGetType        func() T.GType
	SubmenuPlacementGetType        func() T.GType
	TableGetType                   func() T.GType
	TargetFlagsGetType             func() T.GType
	TargetListGetType              func() T.GType
	TearoffMenuItemGetType         func() T.GType
	TearoffMenuItemNew             func() *T.GtkWidget
	TextAttributesGetType          func() T.GType
	TextBufferGetType              func() T.GType
	TextBufferTargetInfoGetType    func() T.GType
	TextChildAnchorGetType         func() T.GType
	TextDirectionGetType           func() T.GType
	TextIterGetType                func() T.GType
	TextMarkGetType                func() T.GType
	TextSearchFlagsGetType         func() T.GType
	TextTagGetType                 func() T.GType
	TextTagTableGetType            func() T.GType
	TextViewGetType                func() T.GType
	TextViewNew                    func() *T.GtkWidget
	TextWindowTypeGetType          func() T.GType
	TipsQueryGetType               func() T.GType
	TipsQueryNew                   func() *T.GtkWidget
	ToggleActionGetType            func() T.GType
	ToggleButtonGetType            func() T.GType
	ToggleButtonNew                func() *T.GtkWidget
	ToggleToolButtonGetType        func() T.GType
	ToolbarChildTypeGetType        func() T.GType
	ToolbarGetType                 func() T.GType
	ToolbarNew                     func() *T.GtkWidget
	ToolbarSpaceStyleGetType       func() T.GType
	ToolbarStyleGetType            func() T.GType
	ToolButtonGetType              func() T.GType
	ToolItemGetType                func() T.GType
	ToolItemGroupGetType           func() T.GType
	ToolPaletteDragTargetsGetType  func() T.GType
	ToolPaletteGetType             func() T.GType
	ToolPaletteNew                 func() *T.GtkWidget
	ToolShellGetType               func() T.GType
	TooltipGetType                 func() T.GType
	TooltipsGetType                func() T.GType
	TreeDragDestGetType            func() T.GType
	TreeDragSourceGetType          func() T.GType
	TreeGetType                    func() T.GType
	TreeItemGetType                func() T.GType
	TreeItemNew                    func() *T.GtkWidget
	TreeIterGetType                func() T.GType
	TreeModelFilterGetType         func() T.GType
	TreeModelFlagsGetType          func() T.GType
	TreeModelGetType               func() T.GType
	TreeModelSortGetType           func() T.GType
	TreeNew                        func() *T.GtkWidget
	TreePathGetType                func() T.GType
	TreeRowReferenceGetType        func() T.GType
	TreeSelectionGetType           func() T.GType
	TreeSortableGetType            func() T.GType
	TreeStoreGetType               func() T.GType
	TreeViewColumnGetType          func() T.GType
	TreeViewColumnSizingGetType    func() T.GType
	TreeViewDropPositionGetType    func() T.GType
	TreeViewGetType                func() T.GType
	TreeViewGridLinesGetType       func() T.GType
	TreeViewModeGetType            func() T.GType
	TreeViewNew                    func() *T.GtkWidget
	UiManagerGetType               func() T.GType
	UiManagerItemTypeGetType       func() T.GType
	UnitGetType                    func() T.GType
	UpdateTypeGetType              func() T.GType
	VboxGetType                    func() T.GType
	VbuttonBoxGetType              func() T.GType
	VbuttonBoxNew                  func() *T.GtkWidget
	ViewportGetType                func() T.GType
	VisibilityGetType              func() T.GType
	VolumeButtonGetType            func() T.GType
	VolumeButtonNew                func() *T.GtkWidget
	VpanedGetType                  func() T.GType
	VpanedNew                      func() *T.GtkWidget
	VrulerGetType                  func() T.GType
	VrulerNew                      func() *T.GtkWidget
	VscaleGetType                  func() T.GType
	VscrollbarGetType              func() T.GType
	VseparatorGetType              func() T.GType
	VseparatorNew                  func() *T.GtkWidget
	WidgetFlagsGetType             func() T.GType
	WidgetGetType                  func() T.GType
	WidgetHelpTypeGetType          func() T.GType
	WindowGetType                  func() T.GType
	WindowGroupGetType             func() T.GType
	WindowPositionGetType          func() T.GType
	WindowTypeGetType              func() T.GType
	WrapModeGetType                func() T.GType

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

	AdjustmentNew func(
		value float64,
		lower float64,
		upper float64,
		stepIncrement float64,
		pageIncrement float64,
		pageSize float64) *T.GtkObject

	AdjustmentChanged func(
		adjustment *T.GtkAdjustment)

	AdjustmentValueChanged func(
		adjustment *T.GtkAdjustment)

	AdjustmentClampPage func(
		adjustment *T.GtkAdjustment, lower, upper float64)

	AdjustmentGetValue func(
		adjustment *T.GtkAdjustment) float64

	AdjustmentSetValue func(
		adjustment *T.GtkAdjustment, value float64)

	AdjustmentGetLower func(
		adjustment *T.GtkAdjustment) float64

	AdjustmentSetLower func(
		adjustment *T.GtkAdjustment, lower float64)

	AdjustmentGetUpper func(
		adjustment *T.GtkAdjustment) float64

	AdjustmentSetUpper func(
		adjustment *T.GtkAdjustment, upper float64)

	AdjustmentGetStepIncrement func(
		adjustment *T.GtkAdjustment) float64

	AdjustmentSetStepIncrement func(
		adjustment *T.GtkAdjustment,
		stepIncrement float64)

	AdjustmentGetPageIncrement func(
		adjustment *T.GtkAdjustment) float64

	AdjustmentSetPageIncrement func(
		adjustment *T.GtkAdjustment,
		pageIncrement float64)

	AdjustmentGetPageSize func(
		adjustment *T.GtkAdjustment) float64

	AdjustmentSetPageSize func(
		adjustment *T.GtkAdjustment,
		pageSize float64)

	AdjustmentConfigure func(
		adjustment *T.GtkAdjustment,
		value float64,
		lower float64,
		upper float64,
		stepIncrement float64,
		pageIncrement float64,
		pageSize float64)

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
		stockId string) *T.GtkIconSet

	StyleLookupColor func(
		style *T.GtkStyle,
		colorName string,
		color *T.GdkColor) T.Gboolean

	StyleRenderIcon func(
		style *T.GtkStyle,
		source *T.GtkIconSource,
		direction T.GtkTextDirection,
		state T.GtkStateType,
		size T.GtkIconSize,
		widget *T.GtkWidget,
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
		expanderStyle T.GtkExpanderStyle)

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
		widget *T.GtkWidget,
		detail string,
		x1 int,
		x2 int,
		y int)

	PaintVline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
		detail string,
		x int,
		y int,
		expanderStyle T.GtkExpanderStyle)

	PaintLayout func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		useText T.Gboolean,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x int,
		y int,
		layout *T.PangoLayout)

	PaintResizeGrip func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		stateType T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
		detail string,
		step uint,
		x int,
		y int,
		width int,
		height int)

	BorderNew func() *T.GtkBorder

	BorderCopy func(
		border *T.GtkBorder) *T.GtkBorder

	BorderFree func(
		border *T.GtkBorder)

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
		widget *T.GtkWidget,
		detail string,
		x int,
		y int,
		string string)

	DrawInsertionCursor func(
		widget *T.GtkWidget,
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
		widget *T.GtkWidget) *T.GtkStyle

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

	WidgetNew func(t T.GType, firstPropertyName string,
		v ...VArg) *T.GtkWidget

	WidgetDestroy func(
		widget *T.GtkWidget)

	WidgetDestroyed func(
		widget *T.GtkWidget,
		widgetPointer **T.GtkWidget)

	WidgetRef func(
		widget *T.GtkWidget) *T.GtkWidget

	WidgetUnref func(
		widget *T.GtkWidget)

	WidgetSet func(widget *T.GtkWidget,
		firstPropertyName string, v ...VArg)

	WidgetHideAll func(
		widget *T.GtkWidget)

	WidgetUnparent func(
		widget *T.GtkWidget)

	WidgetShow func(
		widget *T.GtkWidget)

	WidgetShowNow func(
		widget *T.GtkWidget)

	WidgetHide func(
		widget *T.GtkWidget)

	WidgetShowAll func(
		widget *T.GtkWidget)

	WidgetSetNoShowAll func(
		widget *T.GtkWidget,
		noShowAll T.Gboolean)

	WidgetGetNoShowAll func(
		widget *T.GtkWidget) T.Gboolean

	WidgetMap func(
		widget *T.GtkWidget)

	WidgetUnmap func(
		widget *T.GtkWidget)

	WidgetRealize func(
		widget *T.GtkWidget)

	WidgetUnrealize func(
		widget *T.GtkWidget)

	WidgetQueueDraw func(
		widget *T.GtkWidget)

	WidgetQueueDrawArea func(
		widget *T.GtkWidget, x, y, width, height int)

	WidgetQueueClear func(
		widget *T.GtkWidget)

	WidgetQueueClearArea func(
		widget *T.GtkWidget, x, y, width, height int)

	WidgetQueueResize func(widget *T.GtkWidget)

	WidgetQueueResizeNoRedraw func(widget *T.GtkWidget)

	WidgetDraw func(widget *T.GtkWidget, area *T.GdkRectangle)

	WidgetSizeRequest func(
		widget *T.GtkWidget,
		requisition *T.GtkRequisition)

	WidgetSizeAllocate func(
		widget *T.GtkWidget,
		allocation *T.GtkAllocation)

	WidgetGetChildRequisition func(
		widget *T.GtkWidget,
		requisition *T.GtkRequisition)

	WidgetAddAccelerator func(
		widget *T.GtkWidget,
		accelSignal string,
		accelGroup *T.GtkAccelGroup,
		accelKey uint,
		accelMods T.GdkModifierType,
		accelFlags T.GtkAccelFlags)

	WidgetRemoveAccelerator func(
		widget *T.GtkWidget,
		accelGroup *T.GtkAccelGroup,
		accelKey uint,
		accelMods T.GdkModifierType) T.Gboolean

	WidgetSetAccelPath func(
		widget *T.GtkWidget,
		accelPath string,
		accelGroup *T.GtkAccelGroup)

	WidgetListAccelClosures func(
		widget *T.GtkWidget) *T.GList

	WidgetCanActivateAccel func(
		widget *T.GtkWidget,
		signalId uint) T.Gboolean

	WidgetMnemonicActivate func(
		widget *T.GtkWidget,
		groupCycling T.Gboolean) T.Gboolean

	WidgetEvent func(
		widget *T.GtkWidget,
		event *T.GdkEvent) T.Gboolean

	WidgetSendExpose func(
		widget *T.GtkWidget,
		event *T.GdkEvent) int

	WidgetSendFocusChange func(
		widget *T.GtkWidget,
		event *T.GdkEvent) T.Gboolean

	WidgetActivate func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetScrollAdjustments func(
		widget *T.GtkWidget,
		hadjustment *T.GtkAdjustment,
		vadjustment *T.GtkAdjustment) T.Gboolean

	WidgetReparent func(
		widget *T.GtkWidget,
		newParent *T.GtkWidget)

	WidgetIntersect func(
		widget *T.GtkWidget,
		area *T.GdkRectangle,
		intersection *T.GdkRectangle) T.Gboolean

	WidgetRegionIntersect func(
		widget *T.GtkWidget,
		region *T.GdkRegion) *T.GdkRegion

	WidgetFreezeChildNotify func(
		widget *T.GtkWidget)

	WidgetChildNotify func(
		widget *T.GtkWidget,
		childProperty string)

	WidgetThawChildNotify func(
		widget *T.GtkWidget)

	WidgetSetCanFocus func(
		widget *T.GtkWidget,
		canFocus T.Gboolean)

	WidgetGetCanFocus func(
		widget *T.GtkWidget) T.Gboolean

	WidgetHasFocus func(
		widget *T.GtkWidget) T.Gboolean

	WidgetIsFocus func(
		widget *T.GtkWidget) T.Gboolean

	WidgetGrabFocus func(
		widget *T.GtkWidget)

	WidgetSetCanDefault func(
		widget *T.GtkWidget,
		canDefault T.Gboolean)

	WidgetGetCanDefault func(
		widget *T.GtkWidget) T.Gboolean

	WidgetHasDefault func(
		widget *T.GtkWidget) T.Gboolean

	WidgetGrabDefault func(
		widget *T.GtkWidget)

	WidgetSetReceivesDefault func(
		widget *T.GtkWidget,
		receivesDefault T.Gboolean)

	WidgetGetReceivesDefault func(
		widget *T.GtkWidget) T.Gboolean

	WidgetHasGrab func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetName func(
		widget *T.GtkWidget,
		name string)

	WidgetGetName func(
		widget *T.GtkWidget) string

	WidgetSetState func(
		widget *T.GtkWidget,
		state T.GtkStateType)

	WidgetGetState func(
		widget *T.GtkWidget) T.GtkStateType

	WidgetSetSensitive func(
		widget *T.GtkWidget,
		sensitive T.Gboolean)

	WidgetGetSensitive func(
		widget *T.GtkWidget) T.Gboolean

	WidgetIsSensitive func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetVisible func(
		widget *T.GtkWidget,
		visible T.Gboolean)

	WidgetGetVisible func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetHasWindow func(
		widget *T.GtkWidget,
		hasWindow T.Gboolean)

	WidgetGetHasWindow func(
		widget *T.GtkWidget) T.Gboolean

	WidgetIsToplevel func(
		widget *T.GtkWidget) T.Gboolean

	WidgetIsDrawable func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetRealized func(
		widget *T.GtkWidget,
		realized T.Gboolean)

	WidgetGetRealized func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetMapped func(
		widget *T.GtkWidget,
		mapped T.Gboolean)

	WidgetGetMapped func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetAppPaintable func(
		widget *T.GtkWidget,
		appPaintable T.Gboolean)

	WidgetGetAppPaintable func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetDoubleBuffered func(
		widget *T.GtkWidget,
		doubleBuffered T.Gboolean)

	WidgetGetDoubleBuffered func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetRedrawOnAllocate func(
		widget *T.GtkWidget,
		redrawOnAllocate T.Gboolean)

	WidgetSetParent func(
		widget *T.GtkWidget,
		parent *T.GtkWidget)

	WidgetGetParent func(
		widget *T.GtkWidget) *T.GtkWidget

	WidgetSetParentWindow func(
		widget *T.GtkWidget,
		parentWindow *T.GdkWindow)

	WidgetGetParentWindow func(
		widget *T.GtkWidget) *T.GdkWindow

	WidgetSetChildVisible func(
		widget *T.GtkWidget,
		isVisible T.Gboolean)

	WidgetGetChildVisible func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetWindow func(
		widget *T.GtkWidget,
		window *T.GdkWindow)

	WidgetGetWindow func(
		widget *T.GtkWidget) *T.GdkWindow

	WidgetGetAllocation func(
		widget *T.GtkWidget,
		allocation *T.GtkAllocation)

	WidgetSetAllocation func(
		widget *T.GtkWidget,
		allocation *T.GtkAllocation)

	WidgetGetRequisition func(
		widget *T.GtkWidget,
		requisition *T.GtkRequisition)

	WidgetChildFocus func(
		widget *T.GtkWidget,
		direction T.GtkDirectionType) T.Gboolean

	WidgetKeynavFailed func(
		widget *T.GtkWidget,
		direction T.GtkDirectionType) T.Gboolean

	WidgetErrorBell func(
		widget *T.GtkWidget)

	WidgetSetSizeRequest func(
		widget *T.GtkWidget,
		width int,
		height int)

	WidgetGetSizeRequest func(
		widget *T.GtkWidget,
		width *int,
		height *int)

	WidgetSetUposition func(
		widget *T.GtkWidget,
		x int,
		y int)

	WidgetSetUsize func(
		widget *T.GtkWidget,
		width int,
		height int)

	WidgetSetEvents func(
		widget *T.GtkWidget,
		events int)

	WidgetAddEvents func(
		widget *T.GtkWidget,
		events int)

	WidgetSetExtensionEvents func(
		widget *T.GtkWidget,
		mode T.GdkExtensionMode)

	WidgetGetExtensionEvents func(
		widget *T.GtkWidget) T.GdkExtensionMode

	WidgetGetToplevel func(
		widget *T.GtkWidget) *T.GtkWidget

	WidgetGetAncestor func(
		widget *T.GtkWidget,
		widgetType T.GType) *T.GtkWidget

	WidgetGetColormap func(
		widget *T.GtkWidget) *T.GdkColormap

	WidgetGetVisual func(
		widget *T.GtkWidget) *T.GdkVisual

	WidgetGetScreen func(
		widget *T.GtkWidget) *T.GdkScreen

	WidgetHasScreen func(
		widget *T.GtkWidget) T.Gboolean

	WidgetGetDisplay func(
		widget *T.GtkWidget) *T.GdkDisplay

	WidgetGetRootWindow func(
		widget *T.GtkWidget) *T.GdkWindow

	WidgetGetSettings func(
		widget *T.GtkWidget) *T.GtkSettings

	WidgetGetClipboard func(
		widget *T.GtkWidget,
		selection T.GdkAtom) *T.GtkClipboard

	WidgetGetSnapshot func(
		widget *T.GtkWidget,
		clipRect *T.GdkRectangle) *T.GdkPixmap

	WidgetGetAccessible func(
		widget *T.GtkWidget) *T.AtkObject

	WidgetSetColormap func(
		widget *T.GtkWidget,
		colormap *T.GdkColormap)

	WidgetGetEvents func(
		widget *T.GtkWidget) int

	WidgetGetPointer func(
		widget *T.GtkWidget,
		x *int,
		y *int)

	WidgetIsAncestor func(
		widget *T.GtkWidget,
		ancestor *T.GtkWidget) T.Gboolean

	WidgetTranslateCoordinates func(
		srcWidget *T.GtkWidget,
		destWidget *T.GtkWidget,
		srcX int,
		srcY int,
		destX *int,
		destY *int) T.Gboolean

	WidgetHideOnDelete func(
		widget *T.GtkWidget) T.Gboolean

	WidgetStyleAttach func(
		style *T.GtkWidget)

	WidgetHasRcStyle func(
		widget *T.GtkWidget) T.Gboolean

	WidgetSetStyle func(
		widget *T.GtkWidget,
		style *T.GtkStyle)

	WidgetEnsureStyle func(
		widget *T.GtkWidget)

	WidgetGetStyle func(
		widget *T.GtkWidget) *T.GtkStyle

	WidgetModifyStyle func(
		widget *T.GtkWidget,
		style *T.GtkRcStyle)

	WidgetGetModifierStyle func(
		widget *T.GtkWidget) *T.GtkRcStyle

	WidgetModifyFg func(
		widget *T.GtkWidget,
		state T.GtkStateType,
		color *T.GdkColor)

	WidgetModifyBg func(
		widget *T.GtkWidget,
		state T.GtkStateType,
		color *T.GdkColor)

	WidgetModifyText func(
		widget *T.GtkWidget,
		state T.GtkStateType,
		color *T.GdkColor)

	WidgetModifyBase func(
		widget *T.GtkWidget,
		state T.GtkStateType,
		color *T.GdkColor)

	WidgetModifyCursor func(
		widget *T.GtkWidget,
		primary *T.GdkColor,
		secondary *T.GdkColor)

	WidgetModifyFont func(
		widget *T.GtkWidget,
		fontDesc *T.PangoFontDescription)

	WidgetCreatePangoContext func(
		widget *T.GtkWidget) *T.PangoContext

	WidgetGetPangoContext func(
		widget *T.GtkWidget) *T.PangoContext

	WidgetCreatePangoLayout func(
		widget *T.GtkWidget,
		text string) *T.PangoLayout

	WidgetRenderIcon func(
		widget *T.GtkWidget,
		stockId string,
		size T.GtkIconSize,
		detail string) *T.GdkPixbuf

	WidgetSetCompositeName func(
		widget *T.GtkWidget,
		name string)

	WidgetGetCompositeName func(
		widget *T.GtkWidget) string

	WidgetResetRcStyles func(
		widget *T.GtkWidget)

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
		widget *T.GtkWidget,
		propertyName string,
		value *T.GValue)

	WidgetStyleGetValist func(
		widget *T.GtkWidget,
		firstPropertyName string,
		varArgs T.VaList)

	WidgetStyleGet func(widget *T.GtkWidget,
		firstPropertyName string, v ...VArg)

	WidgetSetDefaultColormap func(
		colormap *T.GdkColormap)

	WidgetGetDefaultStyle func() *T.GtkStyle

	WidgetGetDefaultColormap func() *T.GdkColormap

	WidgetGetDefaultVisual func() *T.GdkVisual

	WidgetSetDirection func(
		widget *T.GtkWidget,
		dir T.GtkTextDirection)

	WidgetGetDirection func(
		widget *T.GtkWidget) T.GtkTextDirection

	WidgetSetDefaultDirection func(
		dir T.GtkTextDirection)

	WidgetGetDefaultDirection func() T.GtkTextDirection

	WidgetIsComposited func(
		widget *T.GtkWidget) T.Gboolean

	WidgetShapeCombineMask func(
		widget *T.GtkWidget,
		shapeMask *T.GdkBitmap,
		offsetX int,
		offsetY int)

	WidgetInputShapeCombineMask func(
		widget *T.GtkWidget,
		shapeMask *T.GdkBitmap,
		offsetX int,
		offsetY int)

	WidgetResetShapes func(
		widget *T.GtkWidget)

	WidgetPath func(
		widget *T.GtkWidget,
		pathLength *uint,
		path **T.Gchar,
		pathReversed **T.Gchar)

	WidgetClassPath func(
		widget *T.GtkWidget,
		pathLength *uint,
		path **T.Gchar,
		pathReversed **T.Gchar)

	WidgetListMnemonicLabels func(
		widget *T.GtkWidget) *T.GList

	WidgetAddMnemonicLabel func(
		widget *T.GtkWidget,
		label *T.GtkWidget)

	WidgetRemoveMnemonicLabel func(
		widget *T.GtkWidget,
		label *T.GtkWidget)

	WidgetSetTooltipWindow func(
		widget *T.GtkWidget,
		customWindow *T.GtkWindow)

	WidgetGetTooltipWindow func(
		widget *T.GtkWidget) *T.GtkWindow

	WidgetTriggerTooltipQuery func(
		widget *T.GtkWidget)

	WidgetSetTooltipText func(
		widget *T.GtkWidget,
		text string)

	WidgetGetTooltipText func(
		widget *T.GtkWidget) string

	WidgetSetTooltipMarkup func(
		widget *T.GtkWidget,
		markup string)

	WidgetGetTooltipMarkup func(
		widget *T.GtkWidget) string

	WidgetSetHasTooltip func(
		widget *T.GtkWidget,
		hasTooltip T.Gboolean)

	WidgetGetHasTooltip func(
		widget *T.GtkWidget) T.Gboolean

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
		widget *T.GtkWidget)

	ContainerRemove func(
		container *T.GtkContainer,
		widget *T.GtkWidget)

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
		child *T.GtkWidget,
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
		child *T.GtkWidget)

	ContainerGetFocusChild func(
		container *T.GtkContainer) *T.GtkWidget

	ContainerSetFocusVadjustment func(
		container *T.GtkContainer,
		adjustment *T.GtkAdjustment)

	ContainerGetFocusVadjustment func(
		container *T.GtkContainer) *T.GtkAdjustment

	ContainerSetFocusHadjustment func(
		container *T.GtkContainer,
		adjustment *T.GtkAdjustment)

	ContainerGetFocusHadjustment func(
		container *T.GtkContainer) *T.GtkAdjustment

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
		container *T.GtkContainer, widget *T.GtkWidget,
		firstPropName string, v ...VArg)

	ContainerChildSet func(
		container *T.GtkContainer, child *T.GtkWidget,
		firstPropName string, v ...VArg)

	ContainerChildGet func(container *T.GtkContainer,
		child *T.GtkWidget, firstPropName string, v ...VArg)

	ContainerChildSetValist func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		firstPropertyName string,
		varArgs T.VaList)

	ContainerChildGetValist func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		firstPropertyName string,
		varArgs T.VaList)

	ContainerChildSetProperty func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		propertyName string,
		value *T.GValue)

	ContainerChildGetProperty func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		propertyName string,
		value *T.GValue)

	ContainerForall func(
		container *T.GtkContainer,
		callback T.GtkCallback,
		callbackData T.Gpointer)

	BinGetChild func(
		bin *T.GtkBin) *T.GtkWidget

	WindowNew func(
		t T.GtkWindowType) *T.GtkWidget

	WindowSetTitle func(
		window *T.GtkWindow,
		title string)

	WindowGetTitle func(
		window *T.GtkWindow) string

	WindowSetWmclass func(
		window *T.GtkWindow,
		wmclassName string,
		wmclassClass string)

	WindowSetRole func(
		window *T.GtkWindow,
		role string)

	WindowSetStartupId func(
		window *T.GtkWindow,
		startupId string)

	WindowGetRole func(
		window *T.GtkWindow) string

	WindowAddAccelGroup func(
		window *T.GtkWindow,
		accelGroup *T.GtkAccelGroup)

	WindowRemoveAccelGroup func(
		window *T.GtkWindow,
		accelGroup *T.GtkAccelGroup)

	WindowSetPosition func(
		window *T.GtkWindow,
		position T.GtkWindowPosition)

	WindowActivateFocus func(
		window *T.GtkWindow) T.Gboolean

	WindowSetFocus func(
		window *T.GtkWindow,
		focus *T.GtkWidget)

	WindowGetFocus func(
		window *T.GtkWindow) *T.GtkWidget

	WindowSetDefault func(
		window *T.GtkWindow,
		defaultWidget *T.GtkWidget)

	WindowGetDefaultWidget func(
		window *T.GtkWindow) *T.GtkWidget

	WindowActivateDefault func(
		window *T.GtkWindow) T.Gboolean

	WindowSetTransientFor func(
		window *T.GtkWindow,
		parent *T.GtkWindow)

	WindowGetTransientFor func(
		window *T.GtkWindow) *T.GtkWindow

	WindowSetOpacity func(
		window *T.GtkWindow,
		opacity float64)

	WindowGetOpacity func(
		window *T.GtkWindow) float64

	WindowSetTypeHint func(
		window *T.GtkWindow,
		hint T.GdkWindowTypeHint)

	WindowGetTypeHint func(
		window *T.GtkWindow) T.GdkWindowTypeHint

	WindowSetSkipTaskbarHint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetSkipTaskbarHint func(
		window *T.GtkWindow) T.Gboolean

	WindowSetSkipPagerHint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetSkipPagerHint func(
		window *T.GtkWindow) T.Gboolean

	WindowSetUrgencyHint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetUrgencyHint func(
		window *T.GtkWindow) T.Gboolean

	WindowSetAcceptFocus func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetAcceptFocus func(
		window *T.GtkWindow) T.Gboolean

	WindowSetFocusOnMap func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetFocusOnMap func(
		window *T.GtkWindow) T.Gboolean

	WindowSetDestroyWithParent func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetDestroyWithParent func(
		window *T.GtkWindow) T.Gboolean

	WindowSetMnemonicsVisible func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetMnemonicsVisible func(
		window *T.GtkWindow) T.Gboolean

	WindowSetResizable func(
		window *T.GtkWindow,
		resizable T.Gboolean)

	WindowGetResizable func(
		window *T.GtkWindow) T.Gboolean

	WindowSetGravity func(
		window *T.GtkWindow,
		gravity T.GdkGravity)

	WindowGetGravity func(
		window *T.GtkWindow) T.GdkGravity

	WindowSetGeometryHints func(
		window *T.GtkWindow,
		geometryWidget *T.GtkWidget,
		geometry *T.GdkGeometry,
		geomMask T.GdkWindowHints)

	WindowSetScreen func(
		window *T.GtkWindow,
		screen *T.GdkScreen)

	WindowGetScreen func(
		window *T.GtkWindow) *T.GdkScreen

	WindowIsActive func(
		window *T.GtkWindow) T.Gboolean

	WindowHasToplevelFocus func(
		window *T.GtkWindow) T.Gboolean

	WindowSetHasFrame func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetHasFrame func(
		window *T.GtkWindow) T.Gboolean

	WindowSetFrameDimensions func(
		window *T.GtkWindow,
		left int,
		top int,
		right int,
		bottom int)

	WindowGetFrameDimensions func(
		window *T.GtkWindow,
		left *int,
		top *int,
		right *int,
		bottom *int)

	WindowSetDecorated func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetDecorated func(
		window *T.GtkWindow) T.Gboolean

	WindowSetDeletable func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowGetDeletable func(
		window *T.GtkWindow) T.Gboolean

	WindowSetIconList func(
		window *T.GtkWindow,
		list *T.GList)

	WindowGetIconList func(
		window *T.GtkWindow) *T.GList

	WindowSetIcon func(
		window *T.GtkWindow,
		icon *T.GdkPixbuf)

	WindowSetIconName func(
		window *T.GtkWindow,
		name string)

	WindowSetIconFromFile func(
		window *T.GtkWindow,
		filename string,
		err **T.GError) T.Gboolean

	WindowGetIcon func(
		window *T.GtkWindow) *T.GdkPixbuf

	WindowGetIconName func(
		window *T.GtkWindow) string

	WindowSetDefaultIconList func(
		list *T.GList)

	WindowGetDefaultIconList func() *T.GList

	WindowSetDefaultIcon func(
		icon *T.GdkPixbuf)

	WindowSetDefaultIconName func(
		name string)

	WindowGetDefaultIconName func() string

	WindowSetDefaultIconFromFile func(
		filename string,
		err **T.GError) T.Gboolean

	WindowSetAutoStartupNotification func(
		setting T.Gboolean)

	WindowSetModal func(
		window *T.GtkWindow,
		modal T.Gboolean)

	WindowGetModal func(
		window *T.GtkWindow) T.Gboolean

	WindowListToplevels func() *T.GList

	WindowAddMnemonic func(
		window *T.GtkWindow,
		keyval uint,
		target *T.GtkWidget)

	WindowRemoveMnemonic func(
		window *T.GtkWindow,
		keyval uint,
		target *T.GtkWidget)

	WindowMnemonicActivate func(
		window *T.GtkWindow,
		keyval uint,
		modifier T.GdkModifierType) T.Gboolean

	WindowSetMnemonicModifier func(
		window *T.GtkWindow,
		modifier T.GdkModifierType)

	WindowGetMnemonicModifier func(
		window *T.GtkWindow) T.GdkModifierType

	WindowActivateKey func(
		window *T.GtkWindow,
		event *T.GdkEventKey) T.Gboolean

	WindowPropagateKeyEvent func(
		window *T.GtkWindow,
		event *T.GdkEventKey) T.Gboolean

	WindowPresent func(
		window *T.GtkWindow)

	WindowPresentWithTime func(
		window *T.GtkWindow,
		timestamp T.GUint32)

	WindowIconify func(
		window *T.GtkWindow)

	WindowDeiconify func(
		window *T.GtkWindow)

	WindowStick func(
		window *T.GtkWindow)

	WindowUnstick func(
		window *T.GtkWindow)

	WindowMaximize func(
		window *T.GtkWindow)

	WindowUnmaximize func(
		window *T.GtkWindow)

	WindowFullscreen func(
		window *T.GtkWindow)

	WindowUnfullscreen func(
		window *T.GtkWindow)

	WindowSetKeepAbove func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowSetKeepBelow func(
		window *T.GtkWindow,
		setting T.Gboolean)

	WindowBeginResizeDrag func(
		window *T.GtkWindow,
		edge T.GdkWindowEdge,
		button int,
		rootX int,
		rootY int,
		timestamp T.GUint32)

	WindowBeginMoveDrag func(
		window *T.GtkWindow,
		button int,
		rootX int,
		rootY int,
		timestamp T.GUint32)

	WindowSetPolicy func(
		window *T.GtkWindow,
		allowShrink int,
		allowGrow int,
		autoShrink int)

	WindowSetDefaultSize func(
		window *T.GtkWindow,
		width int,
		height int)

	WindowGetDefaultSize func(
		window *T.GtkWindow,
		width *int,
		height *int)

	WindowResize func(
		window *T.GtkWindow,
		width int,
		height int)

	WindowGetSize func(
		window *T.GtkWindow,
		width *int,
		height *int)

	WindowMove func(
		window *T.GtkWindow,
		x int,
		y int)

	WindowGetPosition func(
		window *T.GtkWindow,
		rootX *int,
		rootY *int)

	WindowParseGeometry func(
		window *T.GtkWindow,
		geometry string) T.Gboolean

	WindowGetGroup func(
		window *T.GtkWindow) *T.GtkWindowGroup

	WindowHasGroup func(
		window *T.GtkWindow) T.Gboolean

	WindowReshowWithInitialSize func(
		window *T.GtkWindow)

	WindowGetWindowType func(
		window *T.GtkWindow) T.GtkWindowType

	WindowGroupNew func() *T.GtkWindowGroup

	WindowGroupAddWindow func(
		windowGroup *T.GtkWindowGroup,
		window *T.GtkWindow)

	WindowGroupRemoveWindow func(
		windowGroup *T.GtkWindowGroup,
		window *T.GtkWindow)

	WindowGroupListWindows func(
		windowGroup *T.GtkWindowGroup) *T.GList

	WindowRemoveEmbeddedXid func(
		window *T.GtkWindow,
		xid T.GdkNativeWindow)

	WindowAddEmbeddedXid func(
		window *T.GtkWindow,
		xid T.GdkNativeWindow)

	WindowGroupGetCurrentGrab func(
		windowGroup *T.GtkWindowGroup) *T.GtkWidget

	DialogNewWithButtons func(
		title string, parent *T.GtkWindow,
		flags T.GtkDialogFlags, firstButtonText string,
		v ...VArg) *T.GtkWidget

	DialogAddActionWidget func(
		dialog *T.GtkDialog,
		child *T.GtkWidget,
		responseId int)

	DialogAddButton func(
		dialog *T.GtkDialog,
		buttonText string,
		responseId int) *T.GtkWidget

	DialogAddButtons func(dialog *T.GtkDialog,
		firstButtonText string, v ...VArg)

	DialogSetResponseSensitive func(
		dialog *T.GtkDialog,
		responseId int,
		setting T.Gboolean)

	DialogSetDefaultResponse func(
		dialog *T.GtkDialog,
		responseId int)

	DialogGetWidgetForResponse func(
		dialog *T.GtkDialog,
		responseId int) *T.GtkWidget

	DialogGetResponseForWidget func(
		dialog *T.GtkDialog,
		widget *T.GtkWidget) int

	DialogSetHasSeparator func(
		dialog *T.GtkDialog,
		setting T.Gboolean)

	DialogGetHasSeparator func(
		dialog *T.GtkDialog) T.Gboolean

	AlternativeDialogButtonOrder func(
		screen *T.GdkScreen) T.Gboolean

	DialogSetAlternativeButtonOrder func(
		dialog *T.GtkDialog, firstResponseId int, v ...VArg)

	DialogSetAlternativeButtonOrderFromArray func(
		dialog *T.GtkDialog,
		nParams int,
		newOrder *int)

	DialogResponse func(
		dialog *T.GtkDialog,
		responseId int)

	DialogRun func(
		dialog *T.GtkDialog) int

	DialogGetActionArea func(
		dialog *T.GtkDialog) *T.GtkWidget

	DialogGetContentArea func(
		dialog *T.GtkDialog) *T.GtkWidget

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
		child *T.GtkWidget)

	MenuShellPrepend func(
		menuShell *T.GtkMenuShell,
		child *T.GtkWidget)

	MenuShellInsert func(
		menuShell *T.GtkMenuShell,
		child *T.GtkWidget,
		position int)

	MenuShellDeactivate func(
		menuShell *T.GtkMenuShell)

	MenuShellSelectItem func(
		menuShell *T.GtkMenuShell,
		menuItem *T.GtkWidget)

	MenuShellDeselect func(
		menuShell *T.GtkMenuShell)

	MenuShellActivateItem func(
		menuShell *T.GtkMenuShell,
		menuItem *T.GtkWidget,
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
		parentMenuShell *T.GtkWidget,
		parentMenuItem *T.GtkWidget,
		f T.GtkMenuPositionFunc,
		dataGpointer,
		button uint,
		activateTime T.GUint32)

	MenuReposition func(
		menu *T.GtkMenu)

	MenuPopdown func(
		menu *T.GtkMenu)

	MenuGetActive func(
		menu *T.GtkMenu) *T.GtkWidget

	MenuSetActive func(
		menu *T.GtkMenu,
		index uint)

	MenuSetAccelGroup func(
		menu *T.GtkMenu,
		accelGroup *T.GtkAccelGroup)

	MenuGetAccelGroup func(
		menu *T.GtkMenu) *T.GtkAccelGroup

	MenuSetAccelPath func(
		menu *T.GtkMenu,
		accelPath string)

	MenuGetAccelPath func(
		menu *T.GtkMenu) string

	MenuAttachToWidget func(
		menu *T.GtkMenu,
		attachWidget *T.GtkWidget,
		detacher T.GtkMenuDetachFunc)

	MenuDetach func(
		menu *T.GtkMenu)

	MenuGetAttachWidget func(
		menu *T.GtkMenu) *T.GtkWidget

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
		child *T.GtkWidget,
		position int)

	MenuSetScreen func(
		menu *T.GtkMenu,
		screen *T.GdkScreen)

	MenuAttach func(
		menu *T.GtkMenu,
		child *T.GtkWidget,
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
		widget *T.GtkWidget) *T.GList

	MenuSetReserveToggleSize func(
		menu *T.GtkMenu,
		reserveToggleSize T.Gboolean)

	MenuGetReserveToggleSize func(
		menu *T.GtkMenu) T.Gboolean

	LabelNew func(
		str string) *T.GtkWidget

	LabelNewWithMnemonic func(
		str string) *T.GtkWidget

	LabelSetText func(
		label *T.GtkLabel,
		str string)

	LabelGetText func(
		label *T.GtkLabel) string

	LabelSetAttributes func(
		label *T.GtkLabel,
		attrs *T.PangoAttrList)

	LabelGetAttributes func(
		label *T.GtkLabel) *T.PangoAttrList

	LabelSetLabel func(
		label *T.GtkLabel,
		str string)

	LabelGetLabel func(
		label *T.GtkLabel) string

	LabelSetMarkup func(
		label *T.GtkLabel,
		str string)

	LabelSetUseMarkup func(
		label *T.GtkLabel,
		setting T.Gboolean)

	LabelGetUseMarkup func(
		label *T.GtkLabel) T.Gboolean

	LabelSetUseUnderline func(
		label *T.GtkLabel,
		setting T.Gboolean)

	LabelGetUseUnderline func(
		label *T.GtkLabel) T.Gboolean

	LabelSetMarkupWithMnemonic func(
		label *T.GtkLabel,
		str string)

	LabelGetMnemonicKeyval func(
		label *T.GtkLabel) uint

	LabelSetMnemonicWidget func(
		label *T.GtkLabel,
		widget *T.GtkWidget)

	LabelGetMnemonicWidget func(
		label *T.GtkLabel) *T.GtkWidget

	LabelSetTextWithMnemonic func(
		label *T.GtkLabel,
		str string)

	LabelSetJustify func(
		label *T.GtkLabel,
		jtype T.GtkJustification)

	LabelGetJustify func(
		label *T.GtkLabel) T.GtkJustification

	LabelSetEllipsize func(
		label *T.GtkLabel,
		mode T.PangoEllipsizeMode)

	LabelGetEllipsize func(
		label *T.GtkLabel) T.PangoEllipsizeMode

	LabelSetWidthChars func(
		label *T.GtkLabel,
		nChars int)

	LabelGetWidthChars func(
		label *T.GtkLabel) int

	LabelSetMaxWidthChars func(
		label *T.GtkLabel,
		nChars int)

	LabelGetMaxWidthChars func(
		label *T.GtkLabel) int

	LabelSetPattern func(
		label *T.GtkLabel,
		pattern string)

	LabelSetLineWrap func(
		label *T.GtkLabel,
		wrap T.Gboolean)

	LabelGetLineWrap func(
		label *T.GtkLabel) T.Gboolean

	LabelSetLineWrapMode func(
		label *T.GtkLabel,
		wrapMode T.PangoWrapMode)

	LabelGetLineWrapMode func(
		label *T.GtkLabel) T.PangoWrapMode

	LabelSetSelectable func(
		label *T.GtkLabel,
		setting T.Gboolean)

	LabelGetSelectable func(
		label *T.GtkLabel) T.Gboolean

	LabelSetAngle func(
		label *T.GtkLabel,
		angle float64)

	LabelGetAngle func(
		label *T.GtkLabel) float64

	LabelSelectRegion func(
		label *T.GtkLabel,
		startOffset int,
		endOffset int)

	LabelGetSelectionBounds func(
		label *T.GtkLabel,
		start *int,
		end *int) T.Gboolean

	LabelGetLayout func(
		label *T.GtkLabel) *T.PangoLayout

	LabelGetLayoutOffsets func(
		label *T.GtkLabel,
		x *int,
		y *int)

	LabelSetSingleLineMode func(
		label *T.GtkLabel,
		singleLineMode T.Gboolean)

	LabelGetSingleLineMode func(
		label *T.GtkLabel) T.Gboolean

	LabelGetCurrentUri func(
		label *T.GtkLabel) string

	LabelSetTrackVisitedLinks func(
		label *T.GtkLabel,
		trackLinks T.Gboolean)

	LabelGetTrackVisitedLinks func(
		label *T.GtkLabel) T.Gboolean

	LabelGet func(
		label *T.GtkLabel,
		str **T.Gchar)

	LabelParseUline func(
		label *T.GtkLabel,
		string string) uint

	AccelGroupsActivate func(object *T.GObject, accelKey uint,
		accelMods T.GdkModifierType) T.Gboolean

	AccelGroupsFromObject func(object *T.GObject) *T.GSList

	AccessibleSetWidget func(
		accessible *T.GtkAccessible,
		widget *T.GtkWidget)

	AccessibleGetWidget func(
		accessible *T.GtkAccessible) *T.GtkWidget

	AccessibleConnectWidgetDestroyed func(
		accessible *T.GtkAccessible)

	ActionNew func(
		name string,
		label string,
		tooltip string,
		stockId string) *T.GtkAction

	ActionGetName func(
		action *T.GtkAction) string

	ActionIsSensitive func(
		action *T.GtkAction) T.Gboolean

	ActionGetSensitive func(
		action *T.GtkAction) T.Gboolean

	ActionSetSensitive func(
		action *T.GtkAction,
		sensitive T.Gboolean)

	ActionIsVisible func(
		action *T.GtkAction) T.Gboolean

	ActionGetVisible func(
		action *T.GtkAction) T.Gboolean

	ActionSetVisible func(
		action *T.GtkAction,
		visible T.Gboolean)

	ActionActivate func(
		action *T.GtkAction)

	ActionCreateIcon func(
		action *T.GtkAction,
		iconSize T.GtkIconSize) *T.GtkWidget

	ActionCreateMenuItem func(
		action *T.GtkAction) *T.GtkWidget

	ActionCreateToolItem func(
		action *T.GtkAction) *T.GtkWidget

	ActionCreateMenu func(
		action *T.GtkAction) *T.GtkWidget

	ActionGetProxies func(
		action *T.GtkAction) *T.GSList

	ActionConnectAccelerator func(
		action *T.GtkAction)

	ActionDisconnectAccelerator func(
		action *T.GtkAction)

	ActionGetAccelPath func(
		action *T.GtkAction) string

	ActionGetAccelClosure func(
		action *T.GtkAction) *T.GClosure

	WidgetGetAction func(
		widget *T.GtkWidget) *T.GtkAction

	ActionConnectProxy func(
		action *T.GtkAction,
		proxy *T.GtkWidget)

	ActionDisconnectProxy func(
		action *T.GtkAction,
		proxy *T.GtkWidget)

	ActionBlockActivateFrom func(
		action *T.GtkAction,
		proxy *T.GtkWidget)

	ActionUnblockActivateFrom func(
		action *T.GtkAction,
		proxy *T.GtkWidget)

	ActionBlockActivate func(
		action *T.GtkAction)

	ActionUnblockActivate func(
		action *T.GtkAction)

	ActionSetAccelPath func(
		action *T.GtkAction,
		accelPath string)

	ActionSetAccelGroup func(
		action *T.GtkAction,
		accelGroup *T.GtkAccelGroup)

	ActionSetLabel func(
		action *T.GtkAction,
		label string)

	ActionGetLabel func(
		action *T.GtkAction) string

	ActionSetShortLabel func(
		action *T.GtkAction,
		shortLabel string)

	ActionGetShortLabel func(
		action *T.GtkAction) string

	ActionSetTooltip func(
		action *T.GtkAction,
		tooltip string)

	ActionGetTooltip func(
		action *T.GtkAction) string

	ActionSetStockId func(
		action *T.GtkAction,
		stockId string)

	ActionGetStockId func(
		action *T.GtkAction) string

	ActionSetGicon func(
		action *T.GtkAction,
		icon *T.GIcon)

	ActionGetGicon func(
		action *T.GtkAction) *T.GIcon

	ActionSetIconName func(
		action *T.GtkAction,
		iconName string)

	ActionGetIconName func(
		action *T.GtkAction) string

	ActionSetVisibleHorizontal func(
		action *T.GtkAction,
		visibleHorizontal T.Gboolean)

	ActionGetVisibleHorizontal func(
		action *T.GtkAction) T.Gboolean

	ActionSetVisibleVertical func(
		action *T.GtkAction,
		visibleVertical T.Gboolean)

	ActionGetVisibleVertical func(
		action *T.GtkAction) T.Gboolean

	ActionSetIsImportant func(
		action *T.GtkAction,
		isImportant T.Gboolean)

	ActionGetIsImportant func(
		action *T.GtkAction) T.Gboolean

	ActionSetAlwaysShowImage func(
		action *T.GtkAction,
		alwaysShow T.Gboolean)

	ActionGetAlwaysShowImage func(
		action *T.GtkAction) T.Gboolean

	ActionGroupNew func(
		name string) *T.GtkActionGroup

	ActionGroupGetName func(
		actionGroup *T.GtkActionGroup) string

	ActionGroupGetSensitive func(
		actionGroup *T.GtkActionGroup) T.Gboolean

	ActionGroupSetSensitive func(
		actionGroup *T.GtkActionGroup,
		sensitive T.Gboolean)

	ActionGroupGetVisible func(
		actionGroup *T.GtkActionGroup) T.Gboolean

	ActionGroupSetVisible func(
		actionGroup *T.GtkActionGroup,
		visible T.Gboolean)

	ActionGroupGetAction func(
		actionGroup *T.GtkActionGroup,
		actionName string) *T.GtkAction

	ActionGroupListActions func(
		actionGroup *T.GtkActionGroup) *T.GList

	ActionGroupAddAction func(
		actionGroup *T.GtkActionGroup,
		action *T.GtkAction)

	ActionGroupAddActionWithAccel func(
		actionGroup *T.GtkActionGroup,
		action *T.GtkAction,
		accelerator string)

	ActionGroupRemoveAction func(
		actionGroup *T.GtkActionGroup,
		action *T.GtkAction)

	ActionGroupAddActions func(
		actionGroup *T.GtkActionGroup,
		entries *T.GtkActionEntry,
		nEntries uint,
		userData T.Gpointer)

	ActionGroupAddToggleActions func(
		actionGroup *T.GtkActionGroup,
		entries *T.GtkToggleActionEntry,
		nEntries uint,
		userData T.Gpointer)

	ActionGroupAddRadioActions func(
		actionGroup *T.GtkActionGroup,
		entries *T.GtkRadioActionEntry,
		nEntries uint,
		value int,
		onChange T.GCallback,
		userData T.Gpointer)

	ActionGroupAddActionsFull func(
		actionGroup *T.GtkActionGroup,
		entries *T.GtkActionEntry,
		nEntries uint,
		userDataGpointer,
		destroy T.GDestroyNotify)

	ActionGroupAddToggleActionsFull func(
		actionGroup *T.GtkActionGroup,
		entries *T.GtkToggleActionEntry,
		nEntries uint,
		userDataGpointer,
		destroy T.GDestroyNotify)

	ActionGroupAddRadioActionsFull func(
		actionGroup *T.GtkActionGroup,
		entries *T.GtkRadioActionEntry,
		nEntries uint,
		value int,
		onChange T.GCallback,
		userDataGpointer,
		destroy T.GDestroyNotify)

	ActionGroupSetTranslateFunc func(
		actionGroup *T.GtkActionGroup,
		f T.GtkTranslateFunc,
		dataGpointer,
		notify T.GDestroyNotify)

	ActionGroupSetTranslationDomain func(
		actionGroup *T.GtkActionGroup,
		domain string)

	ActionGroupTranslateString func(
		actionGroup *T.GtkActionGroup,
		str string) string

	ActivatableSyncActionProperties func(
		activatable *T.GtkActivatable,
		action *T.GtkAction)

	ActivatableSetRelatedAction func(
		activatable *T.GtkActivatable,
		action *T.GtkAction)

	ActivatableGetRelatedAction func(
		activatable *T.GtkActivatable) *T.GtkAction

	ActivatableSetUseActionAppearance func(
		activatable *T.GtkActivatable,
		useAppearance T.Gboolean)

	ActivatableGetUseActionAppearance func(
		activatable *T.GtkActivatable) T.Gboolean

	ActivatableDoSetRelatedAction func(
		activatable *T.GtkActivatable,
		action *T.GtkAction)

	AlignmentNew func(
		xalign float32,
		yalign float32,
		xscale float32,
		yscale float32) *T.GtkWidget

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

	ArrowNew func(
		arrowType T.GtkArrowType,
		shadowType T.GtkShadowType) *T.GtkWidget

	ArrowSet func(
		arrow *T.GtkArrow,
		arrowType T.GtkArrowType,
		shadowType T.GtkShadowType)

	FrameNew func(
		label string) *T.GtkWidget

	FrameSetLabel func(
		frame *T.GtkFrame,
		label string)

	FrameGetLabel func(
		frame *T.GtkFrame) string

	FrameSetLabelWidget func(
		frame *T.GtkFrame,
		labelWidget *T.GtkWidget)

	FrameGetLabelWidget func(
		frame *T.GtkFrame) *T.GtkWidget

	FrameSetLabelAlign func(
		frame *T.GtkFrame,
		xalign float32,
		yalign float32)

	FrameGetLabelAlign func(
		frame *T.GtkFrame,
		xalign *float32,
		yalign *float32)

	FrameSetShadowType func(
		frame *T.GtkFrame,
		t T.GtkShadowType)

	FrameGetShadowType func(
		frame *T.GtkFrame) T.GtkShadowType

	AspectFrameNew func(
		label string,
		xalign float32,
		yalign float32,
		ratio float32,
		obeyChild T.Gboolean) *T.GtkWidget

	AspectFrameSet func(
		aspectFrame *T.GtkAspectFrame,
		xalign float32,
		yalign float32,
		ratio float32,
		obeyChild T.Gboolean)

	BoxPackStart func(
		box *T.GtkBox,
		child *T.GtkWidget,
		expand T.Gboolean,
		fill T.Gboolean,
		padding uint)

	BoxPackEnd func(
		box *T.GtkBox,
		child *T.GtkWidget,
		expand T.Gboolean,
		fill T.Gboolean,
		padding uint)

	BoxPackStartDefaults func(
		box *T.GtkBox,
		widget *T.GtkWidget)

	BoxPackEndDefaults func(
		box *T.GtkBox,
		widget *T.GtkWidget)

	BoxSetHomogeneous func(
		box *T.GtkBox,
		homogeneous T.Gboolean)

	BoxGetHomogeneous func(
		box *T.GtkBox) T.Gboolean

	BoxSetSpacing func(
		box *T.GtkBox,
		spacing int)

	BoxGetSpacing func(
		box *T.GtkBox) int

	BoxReorderChild func(
		box *T.GtkBox,
		child *T.GtkWidget,
		position int)

	BoxQueryChildPacking func(
		box *T.GtkBox,
		child *T.GtkWidget,
		expand *T.Gboolean,
		fill *T.Gboolean,
		padding *uint,
		packType *T.GtkPackType)

	BoxSetChildPacking func(
		box *T.GtkBox,
		child *T.GtkWidget,
		expand T.Gboolean,
		fill T.Gboolean,
		padding uint,
		packType T.GtkPackType)

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

	BuilderNew func() *T.GtkBuilder

	BuilderAddFromFile func(
		builder *T.GtkBuilder,
		filename string,
		error **T.GError) uint

	BuilderAddFromString func(
		builder *T.GtkBuilder,
		buffer string,
		length T.Gsize,
		error **T.GError) uint

	BuilderAddObjectsFromFile func(
		builder *T.GtkBuilder,
		filename string,
		objectIds **T.Gchar,
		error **T.GError) uint

	BuilderAddObjectsFromString func(
		builder *T.GtkBuilder,
		buffer string,
		length T.Gsize,
		objectIds **T.Gchar,
		error **T.GError) uint

	BuilderGetObject func(
		builder *T.GtkBuilder,
		name string) *T.GObject

	BuilderGetObjects func(
		builder *T.GtkBuilder) *T.GSList

	BuilderConnectSignals func(
		builder *T.GtkBuilder,
		userData T.Gpointer)

	BuilderConnectSignalsFull func(
		builder *T.GtkBuilder,
		f T.GtkBuilderConnectFunc,
		userData T.Gpointer)

	BuilderSetTranslationDomain func(
		builder *T.GtkBuilder,
		domain string)

	BuilderGetTranslationDomain func(
		builder *T.GtkBuilder) string

	BuilderGetTypeFromName func(
		builder *T.GtkBuilder,
		typeName string) T.GType

	BuilderValueFromString func(
		builder *T.GtkBuilder,
		pspec *T.GParamSpec,
		string string,
		value *T.GValue,
		error **T.GError) T.Gboolean

	BuilderValueFromStringType func(
		builder *T.GtkBuilder,
		t T.GType,
		string string,
		value *T.GValue,
		error **T.GError) T.Gboolean

	BuildableSetName func(
		buildable *T.GtkBuildable,
		name string)

	BuildableGetName func(
		buildable *T.GtkBuildable) string

	BuildableAddChild func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		child *T.GObject,
		typ string)

	BuildableSetBuildableProperty func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		name string,
		value *T.GValue)

	BuildableConstructChild func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		name string) *T.GObject

	BuildableCustomTagStart func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		child *T.GObject,
		tagname string,
		parser *T.GMarkupParser,
		data *T.Gpointer) T.Gboolean

	BuildableCustomTagEnd func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		child *T.GObject,
		tagname string,
		data *T.Gpointer)

	BuildableCustomFinished func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		child *T.GObject,
		tagname string,
		data T.Gpointer)

	BuildableParserFinished func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder)

	BuildableGetInternalChild func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		childname string) *T.GObject

	ImageNewFromPixmap func(
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap) *T.GtkWidget

	ImageNewFromImage func(
		image *T.GdkImage,
		mask *T.GdkBitmap) *T.GtkWidget

	ImageNewFromFile func(
		filename string) *T.GtkWidget

	ImageNewFromPixbuf func(
		pixbuf *T.GdkPixbuf) *T.GtkWidget

	ImageNewFromStock func(
		stockId string,
		size T.GtkIconSize) *T.GtkWidget

	ImageNewFromIconSet func(
		iconSet *T.GtkIconSet,
		size T.GtkIconSize) *T.GtkWidget

	ImageNewFromAnimation func(
		animation *T.GdkPixbufAnimation) *T.GtkWidget

	ImageNewFromIconName func(
		iconName string,
		size T.GtkIconSize) *T.GtkWidget

	ImageNewFromGicon func(
		icon *T.GIcon,
		size T.GtkIconSize) *T.GtkWidget

	ImageClear func(
		image *T.GtkImage)

	ImageSetFromPixmap func(
		image *T.GtkImage,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	ImageSetFromImage func(
		image *T.GtkImage,
		gdkImage *T.GdkImage,
		mask *T.GdkBitmap)

	ImageSetFromFile func(
		image *T.GtkImage,
		filename string)

	ImageSetFromPixbuf func(
		image *T.GtkImage,
		pixbuf *T.GdkPixbuf)

	ImageSetFromStock func(
		image *T.GtkImage,
		stockId string,
		size T.GtkIconSize)

	ImageSetFromIconSet func(
		image *T.GtkImage,
		iconSet *T.GtkIconSet,
		size T.GtkIconSize)

	ImageSetFromAnimation func(
		image *T.GtkImage,
		animation *T.GdkPixbufAnimation)

	ImageSetFromIconName func(
		image *T.GtkImage,
		iconName string,
		size T.GtkIconSize)

	ImageSetFromGicon func(
		image *T.GtkImage,
		icon *T.GIcon,
		size T.GtkIconSize)

	ImageSetPixelSize func(
		image *T.GtkImage,
		pixelSize int)

	ImageGetStorageType func(
		image *T.GtkImage) T.GtkImageType

	ImageGetPixmap func(
		image *T.GtkImage,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap)

	ImageGetImage func(
		image *T.GtkImage,
		gdkImage **T.GdkImage,
		mask **T.GdkBitmap)

	ImageGetPixbuf func(
		image *T.GtkImage) *T.GdkPixbuf

	ImageGetStock func(
		image *T.GtkImage,
		stockId **T.Gchar,
		size *T.GtkIconSize)

	ImageGetIconSet func(
		image *T.GtkImage,
		iconSet **T.GtkIconSet,
		size *T.GtkIconSize)

	ImageGetAnimation func(
		image *T.GtkImage) *T.GdkPixbufAnimation

	ImageGetIconName func(
		image *T.GtkImage,
		iconName **T.Gchar,
		size *T.GtkIconSize)

	ImageGetGicon func(
		image *T.GtkImage,
		gicon **T.GIcon,
		size *T.GtkIconSize)

	ImageGetPixelSize func(
		image *T.GtkImage) int

	ImageSet func(
		image *T.GtkImage,
		val *T.GdkImage,
		mask *T.GdkBitmap)

	ImageGet func(
		image *T.GtkImage,
		val **T.GdkImage,
		mask **T.GdkBitmap)

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

	CalendarSelectMonth func(
		calendar *T.GtkCalendar,
		month uint,
		year uint) T.Gboolean

	CalendarSelectDay func(
		calendar *T.GtkCalendar,
		day uint)

	CalendarMarkDay func(
		calendar *T.GtkCalendar,
		day uint) T.Gboolean

	CalendarUnmarkDay func(
		calendar *T.GtkCalendar,
		day uint) T.Gboolean

	CalendarClearMarks func(
		calendar *T.GtkCalendar)

	CalendarSetDisplayOptions func(
		calendar *T.GtkCalendar,
		flags T.GtkCalendarDisplayOptions)

	CalendarGetDisplayOptions func(
		calendar *T.GtkCalendar) T.GtkCalendarDisplayOptions

	CalendarDisplayOptions func(
		calendar *T.GtkCalendar,
		flags T.GtkCalendarDisplayOptions)

	CalendarGetDate func(
		calendar *T.GtkCalendar,
		year *uint,
		month *uint,
		day *uint)

	CalendarSetDetailFunc func(
		calendar *T.GtkCalendar,
		f T.GtkCalendarDetailFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	CalendarSetDetailWidthChars func(
		calendar *T.GtkCalendar,
		chars int)

	CalendarSetDetailHeightRows func(
		calendar *T.GtkCalendar,
		rows int)

	CalendarGetDetailWidthChars func(
		calendar *T.GtkCalendar) int

	CalendarGetDetailHeightRows func(
		calendar *T.GtkCalendar) int

	CalendarFreeze func(
		calendar *T.GtkCalendar)

	CalendarThaw func(
		calendar *T.GtkCalendar)

	CellEditableStartEditing func(
		cellEditable *T.GtkCellEditable,
		event *T.GdkEvent)

	CellEditableEditingDone func(
		cellEditable *T.GtkCellEditable)

	CellEditableRemoveWidget func(
		cellEditable *T.GtkCellEditable)

	CellRendererGetSize func(
		cell *T.GtkCellRenderer,
		widget *T.GtkWidget,
		cellArea *T.GdkRectangle,
		xOffset *int,
		yOffset *int,
		width *int,
		height *int)

	CellRendererRender func(
		cell *T.GtkCellRenderer,
		window *T.GdkWindow,
		widget *T.GtkWidget,
		backgroundArea *T.GdkRectangle,
		cellArea *T.GdkRectangle,
		exposeArea *T.GdkRectangle,
		flags T.GtkCellRendererState)

	CellRendererActivate func(
		cell *T.GtkCellRenderer,
		event *T.GdkEvent,
		widget *T.GtkWidget,
		path string,
		backgroundArea *T.GdkRectangle,
		cellArea *T.GdkRectangle,
		flags T.GtkCellRendererState) T.Gboolean

	CellRendererStartEditing func(
		cell *T.GtkCellRenderer,
		event *T.GdkEvent,
		widget *T.GtkWidget,
		path string,
		backgroundArea *T.GdkRectangle,
		cellArea *T.GdkRectangle,
		flags T.GtkCellRendererState) *T.GtkCellEditable

	CellRendererSetFixedSize func(
		cell *T.GtkCellRenderer,
		width int,
		height int)

	CellRendererGetFixedSize func(
		cell *T.GtkCellRenderer,
		width *int,
		height *int)

	CellRendererSetAlignment func(
		cell *T.GtkCellRenderer,
		xalign float32,
		yalign float32)

	CellRendererGetAlignment func(
		cell *T.GtkCellRenderer,
		xalign *float32,
		yalign *float32)

	CellRendererSetPadding func(
		cell *T.GtkCellRenderer,
		xpad int,
		ypad int)

	CellRendererGetPadding func(
		cell *T.GtkCellRenderer,
		xpad *int,
		ypad *int)

	CellRendererSetVisible func(
		cell *T.GtkCellRenderer,
		visible T.Gboolean)

	CellRendererGetVisible func(
		cell *T.GtkCellRenderer) T.Gboolean

	CellRendererSetSensitive func(
		cell *T.GtkCellRenderer,
		sensitive T.Gboolean)

	CellRendererGetSensitive func(
		cell *T.GtkCellRenderer) T.Gboolean

	CellRendererEditingCanceled func(
		cell *T.GtkCellRenderer)

	CellRendererStopEditing func(
		cell *T.GtkCellRenderer,
		canceled T.Gboolean)

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

	CellLayoutPackStart func(
		cellLayout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		expand T.Gboolean)

	CellLayoutPackEnd func(
		cellLayout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		expand T.Gboolean)

	CellLayoutGetCells func(
		cellLayout *T.GtkCellLayout) *T.GList

	CellLayoutClear func(
		cellLayout *T.GtkCellLayout)

	CellLayoutSetAttributes func(cellLayout *T.GtkCellLayout,
		cell *T.GtkCellRenderer, v ...VArg)

	CellLayoutAddAttribute func(
		cellLayout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		attribute string,
		column int)

	CellLayoutSetCellDataFunc func(
		cellLayout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		f T.GtkCellLayoutDataFunc,
		funcDataGpointer,
		destroy T.GDestroyNotify)

	CellLayoutClearAttributes func(
		cellLayout *T.GtkCellLayout,
		cell *T.GtkCellRenderer)

	CellLayoutReorder func(
		cellLayout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		position int)

	CellRendererTextNew func() *T.GtkCellRenderer

	CellRendererTextSetFixedHeightFromFont func(
		renderer *T.GtkCellRendererText,
		numberOfRows int)

	CellRendererAccelNew func() *T.GtkCellRenderer

	CellRendererComboNew func() *T.GtkCellRenderer

	CellRendererPixbufNew func() *T.GtkCellRenderer

	CellRendererProgressNew func() *T.GtkCellRenderer

	CellRendererSpinNew func() *T.GtkCellRenderer

	CellRendererSpinnerNew func() *T.GtkCellRenderer

	CellRendererToggleNew func() *T.GtkCellRenderer

	CellRendererToggleGetRadio func(
		toggle *T.GtkCellRendererToggle) T.Gboolean

	CellRendererToggleSetRadio func(
		toggle *T.GtkCellRendererToggle,
		radio T.Gboolean)

	CellRendererToggleGetActive func(
		toggle *T.GtkCellRendererToggle) T.Gboolean

	CellRendererToggleSetActive func(
		toggle *T.GtkCellRendererToggle,
		setting T.Gboolean)

	CellRendererToggleGetActivatable func(
		toggle *T.GtkCellRendererToggle) T.Gboolean

	CellRendererToggleSetActivatable func(
		toggle *T.GtkCellRendererToggle,
		setting T.Gboolean)

	CellViewNewWithText func(
		text string) *T.GtkWidget

	CellViewNewWithMarkup func(
		markup string) *T.GtkWidget

	CellViewNewWithPixbuf func(
		pixbuf *T.GdkPixbuf) *T.GtkWidget

	CellViewSetModel func(
		cellView *T.GtkCellView,
		model *T.GtkTreeModel)

	CellViewGetModel func(
		cellView *T.GtkCellView) *T.GtkTreeModel

	CellViewSetDisplayedRow func(
		cellView *T.GtkCellView,
		path *T.GtkTreePath)

	CellViewGetDisplayedRow func(
		cellView *T.GtkCellView) *T.GtkTreePath

	CellViewGetSizeOfRow func(
		cellView *T.GtkCellView,
		path *T.GtkTreePath,
		requisition *T.GtkRequisition) T.Gboolean

	CellViewSetBackgroundColor func(
		cellView *T.GtkCellView,
		color *T.GdkColor)

	CellViewGetCellRenderers func(
		cellView *T.GtkCellView) *T.GList

	ToggleButtonNewWithLabel func(
		label string) *T.GtkWidget

	ToggleButtonNewWithMnemonic func(
		label string) *T.GtkWidget

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

	CheckButtonNewWithLabel func(
		label string) *T.GtkWidget

	CheckButtonNewWithMnemonic func(
		label string) *T.GtkWidget

	ItemSelect func(item *T.GtkItem)

	ItemDeselect func(item *T.GtkItem)

	ItemToggle func(item *T.GtkItem)

	MenuItemNewWithLabel func(
		label string) *T.GtkWidget

	MenuItemNewWithMnemonic func(
		label string) *T.GtkWidget

	MenuItemSetSubmenu func(
		menuItem *T.GtkMenuItem,
		submenu *T.GtkWidget)

	MenuItemGetSubmenu func(
		menuItem *T.GtkMenuItem) *T.GtkWidget

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

	CheckMenuItemNewWithLabel func(
		label string) *T.GtkWidget

	CheckMenuItemNewWithMnemonic func(
		label string) *T.GtkWidget

	CheckMenuItemSetActive func(
		checkMenuItem *T.GtkCheckMenuItem,
		isActive T.Gboolean)

	CheckMenuItemGetActive func(
		checkMenuItem *T.GtkCheckMenuItem) T.Gboolean

	CheckMenuItemToggled func(
		checkMenuItem *T.GtkCheckMenuItem)

	CheckMenuItemSetInconsistent func(
		checkMenuItem *T.GtkCheckMenuItem,
		setting T.Gboolean)

	CheckMenuItemGetInconsistent func(
		checkMenuItem *T.GtkCheckMenuItem) T.Gboolean

	CheckMenuItemSetDrawAsRadio func(
		checkMenuItem *T.GtkCheckMenuItem,
		drawAsRadio T.Gboolean)

	CheckMenuItemGetDrawAsRadio func(
		checkMenuItem *T.GtkCheckMenuItem) T.Gboolean

	CheckMenuItemSetShowToggle func(
		menuItem *T.GtkCheckMenuItem,
		always T.Gboolean)

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
		widget *T.GtkWidget,
		selection T.GdkAtom,
		time T.GUint32) T.Gboolean

	SelectionOwnerSetForDisplay func(
		display *T.GdkDisplay,
		widget *T.GtkWidget,
		selection T.GdkAtom,
		time T.GUint32) T.Gboolean

	SelectionAddTarget func(
		widget *T.GtkWidget,
		selection T.GdkAtom,
		target T.GdkAtom,
		info uint)

	SelectionAddTargets func(
		widget *T.GtkWidget,
		selection T.GdkAtom,
		targets *T.GtkTargetEntry,
		ntargets uint)

	SelectionClearTargets func(
		widget *T.GtkWidget,
		selection T.GdkAtom)

	SelectionConvert func(
		widget *T.GtkWidget,
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
		widget *T.GtkWidget)

	SelectionClear func(
		widget *T.GtkWidget,
		event *T.GdkEventSelection) T.Gboolean

	SelectionDataCopy func(
		data *T.GtkSelectionData) *T.GtkSelectionData

	SelectionDataFree func(
		data *T.GtkSelectionData)

	ClipboardGetForDisplay func(
		display *T.GdkDisplay,
		selection T.GdkAtom) *T.GtkClipboard

	ClipboardGet func(
		selection T.GdkAtom) *T.GtkClipboard

	ClipboardGetDisplay func(
		clipboard *T.GtkClipboard) *T.GdkDisplay

	ClipboardSetWithData func(
		clipboard *T.GtkClipboard,
		targets *T.GtkTargetEntry,
		nTargets uint,
		getFunc T.GtkClipboardGetFunc,
		clearFunc T.GtkClipboardClearFunc,
		userData T.Gpointer) T.Gboolean

	ClipboardSetWithOwner func(
		clipboard *T.GtkClipboard,
		targets *T.GtkTargetEntry,
		nTargets uint,
		getFunc T.GtkClipboardGetFunc,
		clearFunc T.GtkClipboardClearFunc,
		owner *T.GObject) T.Gboolean

	ClipboardGetOwner func(
		clipboard *T.GtkClipboard) *T.GObject

	ClipboardClear func(
		clipboard *T.GtkClipboard)

	ClipboardSetText func(
		clipboard *T.GtkClipboard,
		text string,
		len int)

	ClipboardSetImage func(
		clipboard *T.GtkClipboard,
		pixbuf *T.GdkPixbuf)

	ClipboardRequestContents func(
		clipboard *T.GtkClipboard,
		target T.GdkAtom,
		callback T.GtkClipboardReceivedFunc,
		userData T.Gpointer)

	ClipboardRequestText func(
		clipboard *T.GtkClipboard,
		callback T.GtkClipboardTextReceivedFunc,
		userData T.Gpointer)

	ClipboardRequestRichText func(
		clipboard *T.GtkClipboard,
		buffer *T.GtkTextBuffer,
		callback T.GtkClipboardRichTextReceivedFunc,
		userData T.Gpointer)

	ClipboardRequestImage func(
		clipboard *T.GtkClipboard,
		callback T.GtkClipboardImageReceivedFunc,
		userData T.Gpointer)

	ClipboardRequestUris func(
		clipboard *T.GtkClipboard,
		callback T.GtkClipboardURIReceivedFunc,
		userData T.Gpointer)

	ClipboardRequestTargets func(
		clipboard *T.GtkClipboard,
		callback T.GtkClipboardTargetsReceivedFunc,
		userData T.Gpointer)

	ClipboardWaitForContents func(
		clipboard *T.GtkClipboard,
		target T.GdkAtom) *T.GtkSelectionData

	ClipboardWaitForText func(
		clipboard *T.GtkClipboard) string

	ClipboardWaitForRichText func(
		clipboard *T.GtkClipboard,
		buffer *T.GtkTextBuffer,
		format *T.GdkAtom,
		length *T.Gsize) *uint8

	ClipboardWaitForImage func(
		clipboard *T.GtkClipboard) *T.GdkPixbuf

	ClipboardWaitForUris func(
		clipboard *T.GtkClipboard) **T.Gchar

	ClipboardWaitForTargets func(
		clipboard *T.GtkClipboard,
		targets **T.GdkAtom,
		nTargets *int) T.Gboolean

	ClipboardWaitIsTextAvailable func(
		clipboard *T.GtkClipboard) T.Gboolean

	ClipboardWaitIsRichTextAvailable func(
		clipboard *T.GtkClipboard,
		buffer *T.GtkTextBuffer) T.Gboolean

	ClipboardWaitIsImageAvailable func(
		clipboard *T.GtkClipboard) T.Gboolean

	ClipboardWaitIsUrisAvailable func(
		clipboard *T.GtkClipboard) T.Gboolean

	ClipboardWaitIsTargetAvailable func(
		clipboard *T.GtkClipboard,
		target T.GdkAtom) T.Gboolean

	ClipboardSetCanStore func(
		clipboard *T.GtkClipboard,
		targets *T.GtkTargetEntry,
		nTargets int)

	ClipboardStore func(
		clipboard *T.GtkClipboard)

	ColorButtonNewWithColor func(
		color *T.GdkColor) *T.GtkWidget

	ColorButtonSetColor func(
		colorButton *T.GtkColorButton,
		color *T.GdkColor)

	ColorButtonSetAlpha func(
		colorButton *T.GtkColorButton,
		alpha uint16)

	ColorButtonGetColor func(
		colorButton *T.GtkColorButton,
		color *T.GdkColor)

	ColorButtonGetAlpha func(
		colorButton *T.GtkColorButton) uint16

	ColorButtonSetUseAlpha func(
		colorButton *T.GtkColorButton,
		useAlpha T.Gboolean)

	ColorButtonGetUseAlpha func(
		colorButton *T.GtkColorButton) T.Gboolean

	ColorButtonSetTitle func(
		colorButton *T.GtkColorButton,
		title string)

	ColorButtonGetTitle func(
		colorButton *T.GtkColorButton) string

	VboxNew func(
		homogeneous T.Gboolean,
		spacing int) *T.GtkWidget

	ColorSelectionGetHasOpacityControl func(
		colorsel *T.GtkColorSelection) T.Gboolean

	ColorSelectionSetHasOpacityControl func(
		colorsel *T.GtkColorSelection,
		hasOpacity T.Gboolean)

	ColorSelectionGetHasPalette func(
		colorsel *T.GtkColorSelection) T.Gboolean

	ColorSelectionSetHasPalette func(
		colorsel *T.GtkColorSelection,
		hasPalette T.Gboolean)

	ColorSelectionSetCurrentColor func(
		colorsel *T.GtkColorSelection,
		color *T.GdkColor)

	ColorSelectionSetCurrentAlpha func(
		colorsel *T.GtkColorSelection,
		alpha uint16)

	ColorSelectionGetCurrentColor func(
		colorsel *T.GtkColorSelection,
		color *T.GdkColor)

	ColorSelectionGetCurrentAlpha func(
		colorsel *T.GtkColorSelection) uint16

	ColorSelectionSetPreviousColor func(
		colorsel *T.GtkColorSelection,
		color *T.GdkColor)

	ColorSelectionSetPreviousAlpha func(
		colorsel *T.GtkColorSelection,
		alpha uint16)

	ColorSelectionGetPreviousColor func(
		colorsel *T.GtkColorSelection,
		color *T.GdkColor)

	ColorSelectionGetPreviousAlpha func(
		colorsel *T.GtkColorSelection) uint16

	ColorSelectionIsAdjusting func(
		colorsel *T.GtkColorSelection) T.Gboolean

	ColorSelectionPaletteFromString func(
		str string,
		colors **T.GdkColor,
		nColors *int) T.Gboolean

	ColorSelectionPaletteToString func(
		colors *T.GdkColor,
		nColors int) string

	ColorSelectionSetChangePaletteHook func(
		f T.GtkColorSelectionChangePaletteFunc) T.GtkColorSelectionChangePaletteFunc

	ColorSelectionSetChangePaletteWithScreenHook func(
		f T.GtkColorSelectionChangePaletteWithScreenFunc) T.GtkColorSelectionChangePaletteWithScreenFunc

	ColorSelectionSetColor func(
		colorsel *T.GtkColorSelection,
		color *float64)

	ColorSelectionGetColor func(
		colorsel *T.GtkColorSelection,
		color *float64)

	ColorSelectionSetUpdatePolicy func(
		colorsel *T.GtkColorSelection,
		policy T.GtkUpdateType)

	ColorSelectionDialogNew func(
		title string) *T.GtkWidget

	ColorSelectionDialogGetColorSelection func(
		colorsel *T.GtkColorSelectionDialog) *T.GtkWidget

	DragGetData func(
		widget *T.GtkWidget,
		context *T.GdkDragContext,
		target T.GdkAtom,
		time T.GUint32)

	DragFinish func(
		context *T.GdkDragContext,
		success T.Gboolean,
		del T.Gboolean,
		time T.GUint32)

	DragGetSourceWidget func(
		context *T.GdkDragContext) *T.GtkWidget

	DragHighlight func(
		widget *T.GtkWidget)

	DragUnhighlight func(
		widget *T.GtkWidget)

	DragDestSet func(
		widget *T.GtkWidget,
		flags T.GtkDestDefaults,
		targets *T.GtkTargetEntry,
		nTargets int,
		actions T.GdkDragAction)

	DragDestSetProxy func(
		widget *T.GtkWidget,
		proxyWindow *T.GdkWindow,
		protocol T.GdkDragProtocol,
		useCoordinates T.Gboolean)

	DragDestUnset func(
		widget *T.GtkWidget)

	DragDestFindTarget func(
		widget *T.GtkWidget,
		context *T.GdkDragContext,
		targetList *T.GtkTargetList) T.GdkAtom

	DragDestGetTargetList func(
		widget *T.GtkWidget) *T.GtkTargetList

	DragDestSetTargetList func(
		widget *T.GtkWidget,
		targetList *T.GtkTargetList)

	DragDestAddTextTargets func(
		widget *T.GtkWidget)

	DragDestAddImageTargets func(
		widget *T.GtkWidget)

	DragDestAddUriTargets func(
		widget *T.GtkWidget)

	DragDestSetTrackMotion func(
		widget *T.GtkWidget,
		trackMotion T.Gboolean)

	DragDestGetTrackMotion func(
		widget *T.GtkWidget) T.Gboolean

	DragSourceSet func(
		widget *T.GtkWidget,
		startButtonMask T.GdkModifierType,
		targets *T.GtkTargetEntry,
		nTargets int,
		actions T.GdkDragAction)

	DragSourceUnset func(
		widget *T.GtkWidget)

	DragSourceGetTargetList func(
		widget *T.GtkWidget) *T.GtkTargetList

	DragSourceSetTargetList func(
		widget *T.GtkWidget,
		targetList *T.GtkTargetList)

	DragSourceAddTextTargets func(
		widget *T.GtkWidget)

	DragSourceAddImageTargets func(
		widget *T.GtkWidget)

	DragSourceAddUriTargets func(
		widget *T.GtkWidget)

	DragSourceSetIcon func(
		widget *T.GtkWidget,
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	DragSourceSetIconPixbuf func(
		widget *T.GtkWidget,
		pixbuf *T.GdkPixbuf)

	DragSourceSetIconStock func(
		widget *T.GtkWidget,
		stockId string)

	DragSourceSetIconName func(
		widget *T.GtkWidget,
		iconName string)

	DragBegin func(
		widget *T.GtkWidget,
		targets *T.GtkTargetList,
		actions T.GdkDragAction,
		button int,
		event *T.GdkEvent) *T.GdkDragContext

	DragSetIconWidget func(
		context *T.GdkDragContext,
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
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

	EditableSelectRegion func(
		editable *T.GtkEditable,
		startPos int,
		endPos int)

	EditableGetSelectionBounds func(
		editable *T.GtkEditable,
		startPos *int,
		endPos *int) T.Gboolean

	EditableInsertText func(
		editable *T.GtkEditable,
		newText string,
		newTextLength int,
		position *int)

	EditableDeleteText func(
		editable *T.GtkEditable,
		startPos int,
		endPos int)

	EditableGetChars func(
		editable *T.GtkEditable,
		startPos int,
		endPos int) string

	EditableCutClipboard func(
		editable *T.GtkEditable)

	EditableCopyClipboard func(
		editable *T.GtkEditable)

	EditablePasteClipboard func(
		editable *T.GtkEditable)

	EditableDeleteSelection func(
		editable *T.GtkEditable)

	EditableSetPosition func(
		editable *T.GtkEditable,
		position int)

	EditableGetPosition func(
		editable *T.GtkEditable) int

	EditableSetEditable func(
		editable *T.GtkEditable,
		isEditable T.Gboolean)

	EditableGetEditable func(
		editable *T.GtkEditable) T.Gboolean

	ImContextSetClientWindow func(
		context *T.GtkIMContext,
		window *T.GdkWindow)

	ImContextGetPreeditString func(
		context *T.GtkIMContext,
		str **T.Gchar,
		attrs **T.PangoAttrList,
		cursorPos *int)

	ImContextFilterKeypress func(
		context *T.GtkIMContext,
		event *T.GdkEventKey) T.Gboolean

	ImContextFocusIn func(
		context *T.GtkIMContext)

	ImContextFocusOut func(
		context *T.GtkIMContext)

	ImContextReset func(
		context *T.GtkIMContext)

	ImContextSetCursorLocation func(
		context *T.GtkIMContext,
		area *T.GdkRectangle)

	ImContextSetUsePreedit func(
		context *T.GtkIMContext,
		usePreedit T.Gboolean)

	ImContextSetSurrounding func(
		context *T.GtkIMContext,
		text string,
		len int,
		cursorIndex int)

	ImContextGetSurrounding func(
		context *T.GtkIMContext,
		text **T.Gchar,
		cursorIndex *int) T.Gboolean

	ImContextDeleteSurrounding func(
		context *T.GtkIMContext,
		offset int,
		nChars int) T.Gboolean

	EntryBufferNew func(
		initialChars string,
		nInitialChars int) *T.GtkEntryBuffer

	EntryBufferGetBytes func(
		buffer *T.GtkEntryBuffer) T.Gsize

	EntryBufferGetLength func(
		buffer *T.GtkEntryBuffer) uint

	EntryBufferGetText func(
		buffer *T.GtkEntryBuffer) string

	EntryBufferSetText func(
		buffer *T.GtkEntryBuffer,
		chars string,
		nChars int)

	EntryBufferSetMaxLength func(
		buffer *T.GtkEntryBuffer,
		maxLength int)

	EntryBufferGetMaxLength func(
		buffer *T.GtkEntryBuffer) int

	EntryBufferInsertText func(
		buffer *T.GtkEntryBuffer,
		position uint,
		chars string,
		nChars int) uint

	EntryBufferDeleteText func(
		buffer *T.GtkEntryBuffer,
		position uint,
		nChars int) uint

	EntryBufferEmitInsertedText func(
		buffer *T.GtkEntryBuffer,
		position uint,
		chars string,
		nChars uint)

	EntryBufferEmitDeletedText func(
		buffer *T.GtkEntryBuffer,
		position uint,
		nChars uint)

	ListStoreNew func(nColumns int,
		v ...VArg) *T.GtkListStore

	ListStoreNewv func(
		nColumns int,
		types *T.GType) *T.GtkListStore

	ListStoreSetColumnTypes func(
		listStore *T.GtkListStore,
		nColumns int,
		types *T.GType)

	ListStoreSetValue func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter,
		column int,
		value *T.GValue)

	ListStoreSet func(listStore *T.GtkListStore,
		iter *T.GtkTreeIter, v ...VArg)

	ListStoreSetValuesv func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter,
		columns *int,
		values *T.GValue,
		nValues int)

	ListStoreSetValist func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter,
		varArgs T.VaList)

	ListStoreRemove func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter) T.Gboolean

	ListStoreInsert func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter,
		position int)

	ListStoreInsertBefore func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter,
		sibling *T.GtkTreeIter)

	ListStoreInsertAfter func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter,
		sibling *T.GtkTreeIter)

	ListStoreInsertWithValues func(listStore *T.GtkListStore,
		iter *T.GtkTreeIter, position int, v ...VArg)

	ListStoreInsertWithValuesv func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter,
		position int,
		columns *int,
		values *T.GValue,
		nValues int)

	ListStorePrepend func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter)

	ListStoreAppend func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter)

	ListStoreClear func(
		listStore *T.GtkListStore)

	ListStoreIterIsValid func(
		listStore *T.GtkListStore,
		iter *T.GtkTreeIter) T.Gboolean

	ListStoreReorder func(
		store *T.GtkListStore,
		newOrder *int)

	ListStoreSwap func(
		store *T.GtkListStore,
		a *T.GtkTreeIter,
		b *T.GtkTreeIter)

	ListStoreMoveAfter func(
		store *T.GtkListStore,
		iter *T.GtkTreeIter,
		position *T.GtkTreeIter)

	ListStoreMoveBefore func(
		store *T.GtkListStore,
		iter *T.GtkTreeIter,
		position *T.GtkTreeIter)

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

	EntryCompletionNew func() *T.GtkEntryCompletion

	EntryCompletionGetEntry func(
		completion *T.GtkEntryCompletion) *T.GtkWidget

	EntryCompletionSetModel func(
		completion *T.GtkEntryCompletion,
		model *T.GtkTreeModel)

	EntryCompletionGetModel func(
		completion *T.GtkEntryCompletion) *T.GtkTreeModel

	EntryCompletionSetMatchFunc func(
		completion *T.GtkEntryCompletion,
		f T.GtkEntryCompletionMatchFunc,
		funcDataGpointer,
		funcNotify T.GDestroyNotify)

	EntryCompletionSetMinimumKeyLength func(
		completion *T.GtkEntryCompletion,
		length int)

	EntryCompletionGetMinimumKeyLength func(
		completion *T.GtkEntryCompletion) int

	EntryCompletionComplete func(
		completion *T.GtkEntryCompletion)

	EntryCompletionInsertPrefix func(
		completion *T.GtkEntryCompletion)

	EntryCompletionInsertActionText func(
		completion *T.GtkEntryCompletion,
		index int,
		text string)

	EntryCompletionInsertActionMarkup func(
		completion *T.GtkEntryCompletion,
		index int,
		markup string)

	EntryCompletionDeleteAction func(
		completion *T.GtkEntryCompletion,
		index int)

	EntryCompletionSetInlineCompletion func(
		completion *T.GtkEntryCompletion,
		inlineCompletion T.Gboolean)

	EntryCompletionGetInlineCompletion func(
		completion *T.GtkEntryCompletion) T.Gboolean

	EntryCompletionSetInlineSelection func(
		completion *T.GtkEntryCompletion,
		inlineSelection T.Gboolean)

	EntryCompletionGetInlineSelection func(
		completion *T.GtkEntryCompletion) T.Gboolean

	EntryCompletionSetPopupCompletion func(
		completion *T.GtkEntryCompletion,
		popupCompletion T.Gboolean)

	EntryCompletionGetPopupCompletion func(
		completion *T.GtkEntryCompletion) T.Gboolean

	EntryCompletionSetPopupSetWidth func(
		completion *T.GtkEntryCompletion,
		popupSetWidth T.Gboolean)

	EntryCompletionGetPopupSetWidth func(
		completion *T.GtkEntryCompletion) T.Gboolean

	EntryCompletionSetPopupSingleMatch func(
		completion *T.GtkEntryCompletion,
		popupSingleMatch T.Gboolean)

	EntryCompletionGetPopupSingleMatch func(
		completion *T.GtkEntryCompletion) T.Gboolean

	EntryCompletionGetCompletionPrefix func(
		completion *T.GtkEntryCompletion) string

	EntryCompletionSetTextColumn func(
		completion *T.GtkEntryCompletion,
		column int)

	EntryCompletionGetTextColumn func(
		completion *T.GtkEntryCompletion) int

	EntryNewWithBuffer func(
		buffer *T.GtkEntryBuffer) *T.GtkWidget

	EntryGetBuffer func(
		entry *T.GtkEntry) *T.GtkEntryBuffer

	EntrySetBuffer func(
		entry *T.GtkEntry,
		buffer *T.GtkEntryBuffer)

	EntryGetTextWindow func(
		entry *T.GtkEntry) *T.GdkWindow

	EntrySetVisibility func(
		entry *T.GtkEntry,
		visible T.Gboolean)

	EntryGetVisibility func(
		entry *T.GtkEntry) T.Gboolean

	EntrySetInvisibleChar func(
		entry *T.GtkEntry,
		ch T.Gunichar)

	EntryGetInvisibleChar func(
		entry *T.GtkEntry) T.Gunichar

	EntryUnsetInvisibleChar func(
		entry *T.GtkEntry)

	EntrySetHasFrame func(
		entry *T.GtkEntry,
		setting T.Gboolean)

	EntryGetHasFrame func(
		entry *T.GtkEntry) T.Gboolean

	EntrySetInnerBorder func(
		entry *T.GtkEntry,
		border *T.GtkBorder)

	EntryGetInnerBorder func(
		entry *T.GtkEntry) *T.GtkBorder

	EntrySetOverwriteMode func(
		entry *T.GtkEntry,
		overwrite T.Gboolean)

	EntryGetOverwriteMode func(
		entry *T.GtkEntry) T.Gboolean

	EntrySetMaxLength func(
		entry *T.GtkEntry,
		max int)

	EntryGetMaxLength func(
		entry *T.GtkEntry) int

	EntryGetTextLength func(
		entry *T.GtkEntry) uint16

	EntrySetActivatesDefault func(
		entry *T.GtkEntry,
		setting T.Gboolean)

	EntryGetActivatesDefault func(
		entry *T.GtkEntry) T.Gboolean

	EntrySetWidthChars func(
		entry *T.GtkEntry,
		nChars int)

	EntryGetWidthChars func(
		entry *T.GtkEntry) int

	EntrySetText func(
		entry *T.GtkEntry,
		text string)

	EntryGetText func(
		entry *T.GtkEntry) string

	EntryGetLayout func(
		entry *T.GtkEntry) *T.PangoLayout

	EntryGetLayoutOffsets func(
		entry *T.GtkEntry,
		x *int,
		y *int)

	EntrySetAlignment func(
		entry *T.GtkEntry,
		xalign float32)

	EntryGetAlignment func(
		entry *T.GtkEntry) float32

	EntrySetCompletion func(
		entry *T.GtkEntry,
		completion *T.GtkEntryCompletion)

	EntryGetCompletion func(
		entry *T.GtkEntry) *T.GtkEntryCompletion

	EntryLayoutIndexToTextIndex func(
		entry *T.GtkEntry,
		layoutIndex int) int

	EntryTextIndexToLayoutIndex func(
		entry *T.GtkEntry,
		textIndex int) int

	EntrySetCursorHadjustment func(
		entry *T.GtkEntry,
		adjustment *T.GtkAdjustment)

	EntryGetCursorHadjustment func(
		entry *T.GtkEntry) *T.GtkAdjustment

	EntrySetProgressFraction func(
		entry *T.GtkEntry,
		fraction float64)

	EntryGetProgressFraction func(
		entry *T.GtkEntry) float64

	EntrySetProgressPulseStep func(
		entry *T.GtkEntry,
		fraction float64)

	EntryGetProgressPulseStep func(
		entry *T.GtkEntry) float64

	EntryProgressPulse func(
		entry *T.GtkEntry)

	EntrySetIconFromPixbuf func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		pixbuf *T.GdkPixbuf)

	EntrySetIconFromStock func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		stockId string)

	EntrySetIconFromIconName func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		iconName string)

	EntrySetIconFromGicon func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		icon *T.GIcon)

	EntryGetIconStorageType func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) T.GtkImageType

	EntryGetIconPixbuf func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) *T.GdkPixbuf

	EntryGetIconStock func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) string

	EntryGetIconName func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) string

	EntryGetIconGicon func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) *T.GIcon

	EntrySetIconActivatable func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		activatable T.Gboolean)

	EntryGetIconActivatable func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) T.Gboolean

	EntrySetIconSensitive func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		sensitive T.Gboolean)

	EntryGetIconSensitive func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) T.Gboolean

	EntryGetIconAtPos func(
		entry *T.GtkEntry,
		x int,
		y int) int

	EntrySetIconTooltipText func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		tooltip string)

	EntryGetIconTooltipText func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) string

	EntrySetIconTooltipMarkup func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		tooltip string)

	EntryGetIconTooltipMarkup func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) string

	EntrySetIconDragSource func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition,
		targetList *T.GtkTargetList,
		actions T.GdkDragAction)

	EntryGetCurrentIconDragSource func(
		entry *T.GtkEntry) int

	EntryGetIconWindow func(
		entry *T.GtkEntry,
		iconPos T.GtkEntryIconPosition) *T.GdkWindow

	EntryImContextFilterKeypress func(
		entry *T.GtkEntry,
		event *T.GdkEventKey) T.Gboolean

	EntryResetImContext func(
		entry *T.GtkEntry)

	EntryNewWithMaxLength func(
		max int) *T.GtkWidget

	EntryAppendText func(
		entry *T.GtkEntry,
		text string)

	EntryPrependText func(
		entry *T.GtkEntry,
		text string)

	EntrySetPosition func(
		entry *T.GtkEntry,
		position int)

	EntrySelectRegion func(
		entry *T.GtkEntry,
		start int,
		end int)

	EntrySetEditable func(
		entry *T.GtkEntry,
		editable T.Gboolean)

	TreeViewNewWithModel func(
		model *T.GtkTreeModel) *T.GtkWidget

	TreeViewGetModel func(
		treeView *T.GtkTreeView) *T.GtkTreeModel

	TreeViewSetModel func(
		treeView *T.GtkTreeView,
		model *T.GtkTreeModel)

	TreeViewGetSelection func(
		treeView *T.GtkTreeView) *T.GtkTreeSelection

	TreeViewGetHadjustment func(
		treeView *T.GtkTreeView) *T.GtkAdjustment

	TreeViewSetHadjustment func(
		treeView *T.GtkTreeView,
		adjustment *T.GtkAdjustment)

	TreeViewGetVadjustment func(
		treeView *T.GtkTreeView) *T.GtkAdjustment

	TreeViewSetVadjustment func(
		treeView *T.GtkTreeView,
		adjustment *T.GtkAdjustment)

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
		cell *T.GtkCellRenderer, v ...VArg) int

	TreeViewInsertColumnWithDataFunc func(
		treeView *T.GtkTreeView,
		position int,
		title string,
		cell *T.GtkCellRenderer,
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
		focusCell *T.GtkCellRenderer,
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
		treeView *T.GtkTreeView) *T.GtkEntry

	TreeViewSetSearchEntry func(
		treeView *T.GtkTreeView,
		entry *T.GtkEntry)

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
		cell *T.GtkCellRenderer)

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

	DrawingAreaSize func(
		darea *T.GtkDrawingArea,
		width int,
		height int)

	EventBoxGetVisibleWindow func(
		eventBox *T.GtkEventBox) T.Gboolean

	EventBoxSetVisibleWindow func(
		eventBox *T.GtkEventBox,
		visibleWindow T.Gboolean)

	EventBoxGetAboveChild func(
		eventBox *T.GtkEventBox) T.Gboolean

	EventBoxSetAboveChild func(
		eventBox *T.GtkEventBox,
		aboveChild T.Gboolean)

	ExpanderNew func(
		label string) *T.GtkWidget

	ExpanderNewWithMnemonic func(
		label string) *T.GtkWidget

	ExpanderSetExpanded func(
		expander *T.GtkExpander,
		expanded T.Gboolean)

	ExpanderGetExpanded func(
		expander *T.GtkExpander) T.Gboolean

	ExpanderSetSpacing func(
		expander *T.GtkExpander,
		spacing int)

	ExpanderGetSpacing func(
		expander *T.GtkExpander) int

	ExpanderSetLabel func(
		expander *T.GtkExpander,
		label string)

	ExpanderGetLabel func(
		expander *T.GtkExpander) string

	ExpanderSetUseUnderline func(
		expander *T.GtkExpander,
		useUnderline T.Gboolean)

	ExpanderGetUseUnderline func(
		expander *T.GtkExpander) T.Gboolean

	ExpanderSetUseMarkup func(
		expander *T.GtkExpander,
		useMarkup T.Gboolean)

	ExpanderGetUseMarkup func(
		expander *T.GtkExpander) T.Gboolean

	ExpanderSetLabelWidget func(
		expander *T.GtkExpander,
		labelWidget *T.GtkWidget)

	ExpanderGetLabelWidget func(
		expander *T.GtkExpander) *T.GtkWidget

	ExpanderSetLabelFill func(
		expander *T.GtkExpander,
		labelFill T.Gboolean)

	ExpanderGetLabelFill func(
		expander *T.GtkExpander) T.Gboolean

	FixedPut func(
		fixed *T.GtkFixed,
		widget *T.GtkWidget,
		x int,
		y int)

	FixedMove func(
		fixed *T.GtkFixed,
		widget *T.GtkWidget,
		x int,
		y int)

	FixedSetHasWindow func(
		fixed *T.GtkFixed,
		hasWindow T.Gboolean)

	FixedGetHasWindow func(
		fixed *T.GtkFixed) T.Gboolean

	FileFilterNew func() *T.GtkFileFilter

	FileFilterSetName func(
		filter *T.GtkFileFilter,
		name string)

	FileFilterGetName func(
		filter *T.GtkFileFilter) string

	FileFilterAddMimeType func(
		filter *T.GtkFileFilter,
		mimeType string)

	FileFilterAddPattern func(
		filter *T.GtkFileFilter,
		pattern string)

	FileFilterAddPixbufFormats func(
		filter *T.GtkFileFilter)

	FileFilterAddCustom func(
		filter *T.GtkFileFilter,
		needed T.GtkFileFilterFlags,
		f T.GtkFileFilterFunc,
		dataGpointer,
		notify T.GDestroyNotify)

	FileFilterGetNeeded func(
		filter *T.GtkFileFilter) T.GtkFileFilterFlags

	FileFilterFilter func(
		filter *T.GtkFileFilter,
		filterInfo *T.GtkFileFilterInfo) T.Gboolean

	FileChooserErrorQuark func() T.GQuark

	FileChooserSetAction func(
		chooser *T.GtkFileChooser,
		action T.GtkFileChooserAction)

	FileChooserGetAction func(
		chooser *T.GtkFileChooser) T.GtkFileChooserAction

	FileChooserSetLocalOnly func(
		chooser *T.GtkFileChooser,
		localOnly T.Gboolean)

	FileChooserGetLocalOnly func(
		chooser *T.GtkFileChooser) T.Gboolean

	FileChooserSetSelectMultiple func(
		chooser *T.GtkFileChooser,
		selectMultiple T.Gboolean)

	FileChooserGetSelectMultiple func(
		chooser *T.GtkFileChooser) T.Gboolean

	FileChooserSetShowHidden func(
		chooser *T.GtkFileChooser,
		showHidden T.Gboolean)

	FileChooserGetShowHidden func(
		chooser *T.GtkFileChooser) T.Gboolean

	FileChooserSetDoOverwriteConfirmation func(
		chooser *T.GtkFileChooser,
		doOverwriteConfirmation T.Gboolean)

	FileChooserGetDoOverwriteConfirmation func(
		chooser *T.GtkFileChooser) T.Gboolean

	FileChooserSetCreateFolders func(
		chooser *T.GtkFileChooser,
		createFolders T.Gboolean)

	FileChooserGetCreateFolders func(
		chooser *T.GtkFileChooser) T.Gboolean

	FileChooserSetCurrentName func(
		chooser *T.GtkFileChooser,
		name string)

	FileChooserGetFilename func(
		chooser *T.GtkFileChooser) string

	FileChooserSetFilename func(
		chooser *T.GtkFileChooser,
		filename string) T.Gboolean

	FileChooserSelectFilename func(
		chooser *T.GtkFileChooser,
		filename string) T.Gboolean

	FileChooserUnselectFilename func(
		chooser *T.GtkFileChooser,
		filename string)

	FileChooserSelectAll func(
		chooser *T.GtkFileChooser)

	FileChooserUnselectAll func(
		chooser *T.GtkFileChooser)

	FileChooserGetFilenames func(
		chooser *T.GtkFileChooser) *T.GSList

	FileChooserSetCurrentFolder func(
		chooser *T.GtkFileChooser,
		filename string) T.Gboolean

	FileChooserGetCurrentFolder func(
		chooser *T.GtkFileChooser) string

	FileChooserGetUri func(
		chooser *T.GtkFileChooser) string

	FileChooserSetUri func(
		chooser *T.GtkFileChooser,
		uri string) T.Gboolean

	FileChooserSelectUri func(
		chooser *T.GtkFileChooser,
		uri string) T.Gboolean

	FileChooserUnselectUri func(
		chooser *T.GtkFileChooser,
		uri string)

	FileChooserGetUris func(
		chooser *T.GtkFileChooser) *T.GSList

	FileChooserSetCurrentFolderUri func(
		chooser *T.GtkFileChooser,
		uri string) T.Gboolean

	FileChooserGetCurrentFolderUri func(
		chooser *T.GtkFileChooser) string

	FileChooserGetFile func(
		chooser *T.GtkFileChooser) *T.GFile

	FileChooserSetFile func(
		chooser *T.GtkFileChooser,
		file *T.GFile,
		error **T.GError) T.Gboolean

	FileChooserSelectFile func(
		chooser *T.GtkFileChooser,
		file *T.GFile,
		error **T.GError) T.Gboolean

	FileChooserUnselectFile func(
		chooser *T.GtkFileChooser,
		file *T.GFile)

	FileChooserGetFiles func(
		chooser *T.GtkFileChooser) *T.GSList

	FileChooserSetCurrentFolderFile func(
		chooser *T.GtkFileChooser,
		file *T.GFile,
		error **T.GError) T.Gboolean

	FileChooserGetCurrentFolderFile func(
		chooser *T.GtkFileChooser) *T.GFile

	FileChooserSetPreviewWidget func(
		chooser *T.GtkFileChooser,
		previewWidget *T.GtkWidget)

	FileChooserGetPreviewWidget func(
		chooser *T.GtkFileChooser) *T.GtkWidget

	FileChooserSetPreviewWidgetActive func(
		chooser *T.GtkFileChooser,
		active T.Gboolean)

	FileChooserGetPreviewWidgetActive func(
		chooser *T.GtkFileChooser) T.Gboolean

	FileChooserSetUsePreviewLabel func(
		chooser *T.GtkFileChooser,
		useLabel T.Gboolean)

	FileChooserGetUsePreviewLabel func(
		chooser *T.GtkFileChooser) T.Gboolean

	FileChooserGetPreviewFilename func(
		chooser *T.GtkFileChooser) string

	FileChooserGetPreviewUri func(
		chooser *T.GtkFileChooser) string

	FileChooserGetPreviewFile func(
		chooser *T.GtkFileChooser) *T.GFile

	FileChooserSetExtraWidget func(
		chooser *T.GtkFileChooser,
		extraWidget *T.GtkWidget)

	FileChooserGetExtraWidget func(
		chooser *T.GtkFileChooser) *T.GtkWidget

	FileChooserAddFilter func(
		chooser *T.GtkFileChooser,
		filter *T.GtkFileFilter)

	FileChooserRemoveFilter func(
		chooser *T.GtkFileChooser,
		filter *T.GtkFileFilter)

	FileChooserListFilters func(
		chooser *T.GtkFileChooser) *T.GSList

	FileChooserSetFilter func(
		chooser *T.GtkFileChooser,
		filter *T.GtkFileFilter)

	FileChooserGetFilter func(
		chooser *T.GtkFileChooser) *T.GtkFileFilter

	FileChooserAddShortcutFolder func(
		chooser *T.GtkFileChooser,
		folder string,
		error **T.GError) T.Gboolean

	FileChooserRemoveShortcutFolder func(
		chooser *T.GtkFileChooser,
		folder string,
		error **T.GError) T.Gboolean

	FileChooserListShortcutFolders func(
		chooser *T.GtkFileChooser) *T.GSList

	FileChooserAddShortcutFolderUri func(
		chooser *T.GtkFileChooser,
		uri string,
		error **T.GError) T.Gboolean

	FileChooserRemoveShortcutFolderUri func(
		chooser *T.GtkFileChooser,
		uri string,
		error **T.GError) T.Gboolean

	FileChooserListShortcutFolderUris func(
		chooser *T.GtkFileChooser) *T.GSList

	HboxNew func(
		homogeneous T.Gboolean,
		spacing int) *T.GtkWidget

	FileChooserButtonNew func(
		title string,
		action T.GtkFileChooserAction) *T.GtkWidget

	FileChooserButtonNewWithBackend func(
		title string,
		action T.GtkFileChooserAction,
		backend string) *T.GtkWidget

	FileChooserButtonNewWithDialog func(
		dialog *T.GtkWidget) *T.GtkWidget

	FileChooserButtonGetTitle func(
		button *T.GtkFileChooserButton) string

	FileChooserButtonSetTitle func(
		button *T.GtkFileChooserButton,
		title string)

	FileChooserButtonGetWidthChars func(
		button *T.GtkFileChooserButton) int

	FileChooserButtonSetWidthChars func(
		button *T.GtkFileChooserButton,
		nChars int)

	FileChooserButtonGetFocusOnClick func(
		button *T.GtkFileChooserButton) T.Gboolean

	FileChooserButtonSetFocusOnClick func(
		button *T.GtkFileChooserButton,
		focusOnClick T.Gboolean)

	FileChooserDialogNew func(title string, parent *T.GtkWindow,
		action T.GtkFileChooserAction, firstButtonText string,
		v ...VArg) *T.GtkWidget

	FileChooserDialogNewWithBackend func(title string,
		parent *T.GtkWindow, action T.GtkFileChooserAction,
		backend string, firstButtonText string,
		v ...VArg) *T.GtkWidget

	FileChooserWidgetNew func(
		action T.GtkFileChooserAction) *T.GtkWidget

	FileChooserWidgetNewWithBackend func(
		action T.GtkFileChooserAction,
		backend string) *T.GtkWidget

	FontButtonNewWithFont func(
		fontname string) *T.GtkWidget

	FontButtonGetTitle func(
		fontButton *T.GtkFontButton) string

	FontButtonSetTitle func(
		fontButton *T.GtkFontButton,
		title string)

	FontButtonGetUseFont func(
		fontButton *T.GtkFontButton) T.Gboolean

	FontButtonSetUseFont func(
		fontButton *T.GtkFontButton,
		useFont T.Gboolean)

	FontButtonGetUseSize func(
		fontButton *T.GtkFontButton) T.Gboolean

	FontButtonSetUseSize func(
		fontButton *T.GtkFontButton,
		useSize T.Gboolean)

	FontButtonGetFontName func(
		fontButton *T.GtkFontButton) string

	FontButtonSetFontName func(
		fontButton *T.GtkFontButton,
		fontname string) T.Gboolean

	FontButtonGetShowStyle func(
		fontButton *T.GtkFontButton) T.Gboolean

	FontButtonSetShowStyle func(
		fontButton *T.GtkFontButton,
		showStyle T.Gboolean)

	FontButtonGetShowSize func(
		fontButton *T.GtkFontButton) T.Gboolean

	FontButtonSetShowSize func(
		fontButton *T.GtkFontButton,
		showSize T.Gboolean)

	FontSelectionGetFamilyList func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	FontSelectionGetFaceList func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	FontSelectionGetSizeEntry func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	FontSelectionGetSizeList func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	FontSelectionGetPreviewEntry func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	FontSelectionGetFamily func(
		fontsel *T.GtkFontSelection) *T.PangoFontFamily

	FontSelectionGetFace func(
		fontsel *T.GtkFontSelection) *T.PangoFontFace

	FontSelectionGetSize func(
		fontsel *T.GtkFontSelection) int

	FontSelectionGetFontName func(
		fontsel *T.GtkFontSelection) string

	FontSelectionGetFont func(
		fontsel *T.GtkFontSelection) *T.GdkFont

	FontSelectionSetFontName func(
		fontsel *T.GtkFontSelection,
		fontname string) T.Gboolean

	FontSelectionGetPreviewText func(
		fontsel *T.GtkFontSelection) string

	FontSelectionSetPreviewText func(
		fontsel *T.GtkFontSelection,
		text string)

	FontSelectionDialogNew func(
		title string) *T.GtkWidget

	FontSelectionDialogGetOkButton func(
		fsd *T.GtkFontSelectionDialog) *T.GtkWidget

	FontSelectionDialogGetApplyButton func(
		fsd *T.GtkFontSelectionDialog) *T.GtkWidget

	FontSelectionDialogGetCancelButton func(
		fsd *T.GtkFontSelectionDialog) *T.GtkWidget

	FontSelectionDialogGetFontSelection func(
		fsd *T.GtkFontSelectionDialog) *T.GtkWidget

	FontSelectionDialogGetFontName func(
		fsd *T.GtkFontSelectionDialog) string

	FontSelectionDialogGetFont func(
		fsd *T.GtkFontSelectionDialog) *T.GdkFont

	FontSelectionDialogSetFontName func(
		fsd *T.GtkFontSelectionDialog,
		fontname string) T.Gboolean

	FontSelectionDialogGetPreviewText func(
		fsd *T.GtkFontSelectionDialog) string

	FontSelectionDialogSetPreviewText func(
		fsd *T.GtkFontSelectionDialog,
		text string)

	GcGet func(
		depth int,
		colormap *T.GdkColormap,
		values *T.GdkGCValues,
		valuesMask T.GdkGCValuesMask) *T.GdkGC

	GcRelease func(
		gc *T.GdkGC)

	HandleBoxSetShadowType func(
		handleBox *T.GtkHandleBox,
		t T.GtkShadowType)

	HandleBoxGetShadowType func(
		handleBox *T.GtkHandleBox) T.GtkShadowType

	HandleBoxSetHandlePosition func(
		handleBox *T.GtkHandleBox,
		position T.GtkPositionType)

	HandleBoxGetHandlePosition func(
		handleBox *T.GtkHandleBox) T.GtkPositionType

	HandleBoxSetSnapEdge func(
		handleBox *T.GtkHandleBox,
		edge T.GtkPositionType)

	HandleBoxGetSnapEdge func(
		handleBox *T.GtkHandleBox) T.GtkPositionType

	HandleBoxGetChildDetached func(
		handleBox *T.GtkHandleBox) T.Gboolean

	HbuttonBoxGetSpacingDefault func() int

	HbuttonBoxGetLayoutDefault func() T.GtkButtonBoxStyle

	HbuttonBoxSetSpacingDefault func(
		spacing int)

	HbuttonBoxSetLayoutDefault func(
		layout T.GtkButtonBoxStyle)

	PanedAdd1 func(
		paned *T.GtkPaned,
		child *T.GtkWidget)

	PanedAdd2 func(
		paned *T.GtkPaned,
		child *T.GtkWidget)

	PanedPack1 func(
		paned *T.GtkPaned,
		child *T.GtkWidget,
		resize T.Gboolean,
		shrink T.Gboolean)

	PanedPack2 func(
		paned *T.GtkPaned,
		child *T.GtkWidget,
		resize T.Gboolean,
		shrink T.Gboolean)

	PanedGetPosition func(
		paned *T.GtkPaned) int

	PanedSetPosition func(
		paned *T.GtkPaned,
		position int)

	PanedGetChild1 func(
		paned *T.GtkPaned) *T.GtkWidget

	PanedGetChild2 func(
		paned *T.GtkPaned) *T.GtkWidget

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
		adjustment *T.GtkAdjustment)

	RangeGetAdjustment func(
		r *T.GtkRange) *T.GtkAdjustment

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

	HscaleNew func(
		adjustment *T.GtkAdjustment) *T.GtkWidget

	HscaleNewWithRange func(
		min float64,
		max float64,
		step float64) *T.GtkWidget

	HscrollbarNew func(
		adjustment *T.GtkAdjustment) *T.GtkWidget

	HsvSetColor func(
		hsv *T.GtkHSV, h, s, v float64)

	HsvGetColor func(
		hsv *T.GtkHSV, h, s, v *float64)

	HsvSetMetrics func(
		hsv *T.GtkHSV,
		size int,
		ringWidth int)

	HsvGetMetrics func(
		hsv *T.GtkHSV,
		size *int,
		ringWidth *int)

	HsvIsAdjusting func(
		hsv *T.GtkHSV) T.Gboolean

	HsvToRgb func(h, s, v float64, r, g, b *float64)

	RgbToHsv func(r, g, b float64, h, s, v *float64)

	IconFactoryNew func() *T.GtkIconFactory

	IconFactoryAdd func(
		factory *T.GtkIconFactory,
		stockId string,
		iconSet *T.GtkIconSet)

	IconFactoryLookup func(
		factory *T.GtkIconFactory,
		stockId string) *T.GtkIconSet

	IconFactoryAddDefault func(
		factory *T.GtkIconFactory)

	IconFactoryRemoveDefault func(
		factory *T.GtkIconFactory)

	IconFactoryLookupDefault func(
		stockId string) *T.GtkIconSet

	IconSizeLookup func(
		size T.GtkIconSize,
		width *int,
		height *int) T.Gboolean

	IconSizeLookupForSettings func(
		settings *T.GtkSettings,
		size T.GtkIconSize,
		width *int,
		height *int) T.Gboolean

	IconSizeRegister func(
		name string,
		width int,
		height int) T.GtkIconSize

	IconSizeRegisterAlias func(
		alias string,
		target T.GtkIconSize)

	IconSizeFromName func(
		name string) T.GtkIconSize

	IconSizeGetName func(
		size T.GtkIconSize) string

	IconSetNew func() *T.GtkIconSet

	IconSetNewFromPixbuf func(
		pixbuf *T.GdkPixbuf) *T.GtkIconSet

	IconSetRef func(
		iconSet *T.GtkIconSet) *T.GtkIconSet

	IconSetUnref func(
		iconSet *T.GtkIconSet)

	IconSetCopy func(
		iconSet *T.GtkIconSet) *T.GtkIconSet

	IconSetRenderIcon func(
		iconSet *T.GtkIconSet,
		style *T.GtkStyle,
		direction T.GtkTextDirection,
		state T.GtkStateType,
		size T.GtkIconSize,
		widget *T.GtkWidget,
		detail string) *T.GdkPixbuf

	IconSetAddSource func(
		iconSet *T.GtkIconSet,
		source *T.GtkIconSource)

	IconSetGetSizes func(
		iconSet *T.GtkIconSet,
		sizes **T.GtkIconSize,
		nSizes *int)

	IconSourceNew func() *T.GtkIconSource

	IconSourceCopy func(
		source *T.GtkIconSource) *T.GtkIconSource

	IconSourceFree func(
		source *T.GtkIconSource)

	IconSourceSetFilename func(
		source *T.GtkIconSource,
		filename string)

	IconSourceSetIconName func(
		source *T.GtkIconSource,
		iconName string)

	IconSourceSetPixbuf func(
		source *T.GtkIconSource,
		pixbuf *T.GdkPixbuf)

	IconSourceGetFilename func(
		source *T.GtkIconSource) string

	IconSourceGetIconName func(
		source *T.GtkIconSource) string

	IconSourceGetPixbuf func(
		source *T.GtkIconSource) *T.GdkPixbuf

	IconSourceSetDirectionWildcarded func(
		source *T.GtkIconSource,
		setting T.Gboolean)

	IconSourceSetStateWildcarded func(
		source *T.GtkIconSource,
		setting T.Gboolean)

	IconSourceSetSizeWildcarded func(
		source *T.GtkIconSource,
		setting T.Gboolean)

	IconSourceGetSizeWildcarded func(
		source *T.GtkIconSource) T.Gboolean

	IconSourceGetStateWildcarded func(
		source *T.GtkIconSource) T.Gboolean

	IconSourceGetDirectionWildcarded func(
		source *T.GtkIconSource) T.Gboolean

	IconSourceSetDirection func(
		source *T.GtkIconSource,
		direction T.GtkTextDirection)

	IconSourceSetState func(
		source *T.GtkIconSource,
		state T.GtkStateType)

	IconSourceSetSize func(
		source *T.GtkIconSource,
		size T.GtkIconSize)

	IconSourceGetDirection func(
		source *T.GtkIconSource) T.GtkTextDirection

	IconSourceGetState func(
		source *T.GtkIconSource) T.GtkStateType

	IconSourceGetSize func(
		source *T.GtkIconSource) T.GtkIconSize

	IconThemeErrorQuark func() T.GQuark

	IconThemeNew func() *T.GtkIconTheme

	IconThemeGetDefault func() *T.GtkIconTheme

	IconThemeGetForScreen func(
		screen *T.GdkScreen) *T.GtkIconTheme

	IconThemeSetScreen func(
		iconTheme *T.GtkIconTheme,
		screen *T.GdkScreen)

	IconThemeSetSearchPath func(
		iconTheme *T.GtkIconTheme,
		path **T.Gchar,
		nElements int)

	IconThemeGetSearchPath func(
		iconTheme *T.GtkIconTheme,
		path ***T.Gchar,
		nElements *int)

	IconThemeAppendSearchPath func(
		iconTheme *T.GtkIconTheme,
		path string)

	IconThemePrependSearchPath func(
		iconTheme *T.GtkIconTheme,
		path string)

	IconThemeSetCustomTheme func(
		iconTheme *T.GtkIconTheme,
		themeName string)

	IconThemeHasIcon func(
		iconTheme *T.GtkIconTheme,
		iconName string) T.Gboolean

	IconThemeGetIconSizes func(
		iconTheme *T.GtkIconTheme,
		iconName string) *int

	IconThemeLookupIcon func(
		iconTheme *T.GtkIconTheme,
		iconName string,
		size int,
		flags T.GtkIconLookupFlags) *T.GtkIconInfo

	IconThemeChooseIcon func(
		iconTheme *T.GtkIconTheme,
		iconNames **T.Gchar,
		size int,
		flags T.GtkIconLookupFlags) *T.GtkIconInfo

	IconThemeLoadIcon func(
		iconTheme *T.GtkIconTheme,
		iconName string,
		size int,
		flags T.GtkIconLookupFlags,
		error **T.GError) *T.GdkPixbuf

	IconThemeLookupByGicon func(
		iconTheme *T.GtkIconTheme,
		icon *T.GIcon,
		size int,
		flags T.GtkIconLookupFlags) *T.GtkIconInfo

	IconThemeListIcons func(
		iconTheme *T.GtkIconTheme,
		context string) *T.GList

	IconThemeListContexts func(
		iconTheme *T.GtkIconTheme) *T.GList

	IconThemeGetExampleIconName func(
		iconTheme *T.GtkIconTheme) string

	IconThemeRescanIfNeeded func(
		iconTheme *T.GtkIconTheme) T.Gboolean

	IconThemeAddBuiltinIcon func(
		iconName string,
		size int,
		pixbuf *T.GdkPixbuf)

	IconInfoCopy func(
		iconInfo *T.GtkIconInfo) *T.GtkIconInfo

	IconInfoFree func(
		iconInfo *T.GtkIconInfo)

	IconInfoNewForPixbuf func(
		iconTheme *T.GtkIconTheme,
		pixbuf *T.GdkPixbuf) *T.GtkIconInfo

	IconInfoGetBaseSize func(
		iconInfo *T.GtkIconInfo) int

	IconInfoGetFilename func(
		iconInfo *T.GtkIconInfo) string

	IconInfoGetBuiltinPixbuf func(
		iconInfo *T.GtkIconInfo) *T.GdkPixbuf

	IconInfoLoadIcon func(
		iconInfo *T.GtkIconInfo,
		error **T.GError) *T.GdkPixbuf

	IconInfoSetRawCoordinates func(
		iconInfo *T.GtkIconInfo,
		rawCoordinates T.Gboolean)

	IconInfoGetEmbeddedRect func(
		iconInfo *T.GtkIconInfo,
		rectangle *T.GdkRectangle) T.Gboolean

	IconInfoGetAttachPoints func(
		iconInfo *T.GtkIconInfo,
		points **T.GdkPoint,
		nPoints *int) T.Gboolean

	IconInfoGetDisplayName func(
		iconInfo *T.GtkIconInfo) string

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
		size T.GtkIconSize)

	TooltipSetIconFromIconName func(
		tooltip *T.GtkTooltip,
		iconName string,
		size T.GtkIconSize)

	TooltipSetIconFromGicon func(
		tooltip *T.GtkTooltip,
		gicon *T.GIcon,
		size T.GtkIconSize)

	TooltipSetCustom func(
		tooltip *T.GtkTooltip,
		customWidget *T.GtkWidget)

	TooltipSetTipArea func(
		tooltip *T.GtkTooltip,
		rect *T.GdkRectangle)

	TooltipTriggerTooltipQuery func(
		display *T.GdkDisplay)

	IconViewNewWithModel func(
		model *T.GtkTreeModel) *T.GtkWidget

	IconViewSetModel func(
		iconView *T.GtkIconView,
		model *T.GtkTreeModel)

	IconViewGetModel func(
		iconView *T.GtkIconView) *T.GtkTreeModel

	IconViewSetTextColumn func(
		iconView *T.GtkIconView,
		column int)

	IconViewGetTextColumn func(
		iconView *T.GtkIconView) int

	IconViewSetMarkupColumn func(
		iconView *T.GtkIconView,
		column int)

	IconViewGetMarkupColumn func(
		iconView *T.GtkIconView) int

	IconViewSetPixbufColumn func(
		iconView *T.GtkIconView,
		column int)

	IconViewGetPixbufColumn func(
		iconView *T.GtkIconView) int

	IconViewSetOrientation func(
		iconView *T.GtkIconView,
		orientation T.GtkOrientation)

	IconViewGetOrientation func(
		iconView *T.GtkIconView) T.GtkOrientation

	IconViewSetItemOrientation func(
		iconView *T.GtkIconView,
		orientation T.GtkOrientation)

	IconViewGetItemOrientation func(
		iconView *T.GtkIconView) T.GtkOrientation

	IconViewSetColumns func(
		iconView *T.GtkIconView,
		columns int)

	IconViewGetColumns func(
		iconView *T.GtkIconView) int

	IconViewSetItemWidth func(
		iconView *T.GtkIconView,
		itemWidth int)

	IconViewGetItemWidth func(
		iconView *T.GtkIconView) int

	IconViewSetSpacing func(
		iconView *T.GtkIconView,
		spacing int)

	IconViewGetSpacing func(
		iconView *T.GtkIconView) int

	IconViewSetRowSpacing func(
		iconView *T.GtkIconView,
		rowSpacing int)

	IconViewGetRowSpacing func(
		iconView *T.GtkIconView) int

	IconViewSetColumnSpacing func(
		iconView *T.GtkIconView,
		columnSpacing int)

	IconViewGetColumnSpacing func(
		iconView *T.GtkIconView) int

	IconViewSetMargin func(
		iconView *T.GtkIconView,
		margin int)

	IconViewGetMargin func(
		iconView *T.GtkIconView) int

	IconViewSetItemPadding func(
		iconView *T.GtkIconView,
		itemPadding int)

	IconViewGetItemPadding func(
		iconView *T.GtkIconView) int

	IconViewGetPathAtPos func(
		iconView *T.GtkIconView,
		x int,
		y int) *T.GtkTreePath

	IconViewGetItemAtPos func(
		iconView *T.GtkIconView,
		x int,
		y int,
		path **T.GtkTreePath,
		cell **T.GtkCellRenderer) T.Gboolean

	IconViewGetVisibleRange func(
		iconView *T.GtkIconView,
		startPath **T.GtkTreePath,
		endPath **T.GtkTreePath) T.Gboolean

	IconViewSelectedForeach func(
		iconView *T.GtkIconView,
		f T.GtkIconViewForeachFunc,
		data T.Gpointer)

	IconViewSetSelectionMode func(
		iconView *T.GtkIconView,
		mode T.GtkSelectionMode)

	IconViewGetSelectionMode func(
		iconView *T.GtkIconView) T.GtkSelectionMode

	IconViewSelectPath func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath)

	IconViewUnselectPath func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath)

	IconViewPathIsSelected func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath) T.Gboolean

	IconViewGetItemRow func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath) int

	IconViewGetItemColumn func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath) int

	IconViewGetSelectedItems func(
		iconView *T.GtkIconView) *T.GList

	IconViewSelectAll func(
		iconView *T.GtkIconView)

	IconViewUnselectAll func(
		iconView *T.GtkIconView)

	IconViewItemActivated func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath)

	IconViewSetCursor func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath,
		cell *T.GtkCellRenderer,
		startEditing T.Gboolean)

	IconViewGetCursor func(
		iconView *T.GtkIconView,
		path **T.GtkTreePath,
		cell **T.GtkCellRenderer) T.Gboolean

	IconViewScrollToPath func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath,
		useAlign T.Gboolean,
		rowAlign float32,
		colAlign float32)

	IconViewEnableModelDragSource func(
		iconView *T.GtkIconView,
		startButtonMask T.GdkModifierType,
		targets *T.GtkTargetEntry,
		nTargets int,
		actions T.GdkDragAction)

	IconViewEnableModelDragDest func(
		iconView *T.GtkIconView,
		targets *T.GtkTargetEntry,
		nTargets int,
		actions T.GdkDragAction)

	IconViewUnsetModelDragSource func(
		iconView *T.GtkIconView)

	IconViewUnsetModelDragDest func(
		iconView *T.GtkIconView)

	IconViewSetReorderable func(
		iconView *T.GtkIconView,
		reorderable T.Gboolean)

	IconViewGetReorderable func(
		iconView *T.GtkIconView) T.Gboolean

	IconViewSetDragDestItem func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath,
		pos T.GtkIconViewDropPosition)

	IconViewGetDragDestItem func(
		iconView *T.GtkIconView,
		path **T.GtkTreePath,
		pos *T.GtkIconViewDropPosition)

	IconViewGetDestItemAtPos func(
		iconView *T.GtkIconView,
		dragX int,
		dragY int,
		path **T.GtkTreePath,
		pos *T.GtkIconViewDropPosition) T.Gboolean

	IconViewCreateDragIcon func(
		iconView *T.GtkIconView,
		path *T.GtkTreePath) *T.GdkPixmap

	IconViewConvertWidgetToBinWindowCoords func(
		iconView *T.GtkIconView,
		wx int,
		wy int,
		bx *int,
		by *int)

	IconViewSetTooltipItem func(
		iconView *T.GtkIconView,
		tooltip *T.GtkTooltip,
		path *T.GtkTreePath)

	IconViewSetTooltipCell func(
		iconView *T.GtkIconView,
		tooltip *T.GtkTooltip,
		path *T.GtkTreePath,
		cell *T.GtkCellRenderer)

	IconViewGetTooltipContext func(
		iconView *T.GtkIconView,
		x *int,
		y *int,
		keyboardTip T.Gboolean,
		model **T.GtkTreeModel,
		path **T.GtkTreePath,
		iter *T.GtkTreeIter) T.Gboolean

	IconViewSetTooltipColumn func(
		iconView *T.GtkIconView,
		column int)

	IconViewGetTooltipColumn func(
		iconView *T.GtkIconView) int

	ImageMenuItemNewWithLabel func(
		label string) *T.GtkWidget

	ImageMenuItemNewWithMnemonic func(
		label string) *T.GtkWidget

	ImageMenuItemNewFromStock func(
		stockId string,
		accelGroup *T.GtkAccelGroup) *T.GtkWidget

	ImageMenuItemSetAlwaysShowImage func(
		imageMenuItem *T.GtkImageMenuItem,
		alwaysShow T.Gboolean)

	ImageMenuItemGetAlwaysShowImage func(
		imageMenuItem *T.GtkImageMenuItem) T.Gboolean

	ImageMenuItemSetImage func(
		imageMenuItem *T.GtkImageMenuItem,
		image *T.GtkWidget)

	ImageMenuItemGetImage func(
		imageMenuItem *T.GtkImageMenuItem) *T.GtkWidget

	ImageMenuItemSetUseStock func(
		imageMenuItem *T.GtkImageMenuItem,
		useStock T.Gboolean)

	ImageMenuItemGetUseStock func(
		imageMenuItem *T.GtkImageMenuItem) T.Gboolean

	ImageMenuItemSetAccelGroup func(
		imageMenuItem *T.GtkImageMenuItem,
		accelGroup *T.GtkAccelGroup)

	ImContextSimpleNew func() *T.GtkIMContext

	ImContextSimpleAddTable func(
		contextSimple *T.GtkIMContextSimple,
		data *uint16,
		maxSeqLen int,
		nSeqs int)

	ImMulticontextNew func() *T.GtkIMContext

	ImMulticontextAppendMenuitems func(
		context *T.GtkIMMulticontext,
		menushell *T.GtkMenuShell)

	ImMulticontextGetContextId func(
		context *T.GtkIMMulticontext) string

	ImMulticontextSetContextId func(
		context *T.GtkIMMulticontext,
		contextId string)

	InfoBarNewWithButtons func(
		firstButtonText string, v ...VArg) *T.GtkWidget

	InfoBarGetActionArea func(
		infoBar *T.GtkInfoBar) *T.GtkWidget

	InfoBarGetContentArea func(
		infoBar *T.GtkInfoBar) *T.GtkWidget

	InfoBarAddActionWidget func(
		infoBar *T.GtkInfoBar,
		child *T.GtkWidget,
		responseId int)

	InfoBarAddButton func(
		infoBar *T.GtkInfoBar,
		buttonText string,
		responseId int) *T.GtkWidget

	InfoBarAddButtons func(infoBar *T.GtkInfoBar,
		firstButtonText string, v ...VArg)

	InfoBarSetResponseSensitive func(
		infoBar *T.GtkInfoBar,
		responseId int,
		setting T.Gboolean)

	InfoBarSetDefaultResponse func(
		infoBar *T.GtkInfoBar,
		responseId int)

	InfoBarResponse func(
		infoBar *T.GtkInfoBar,
		responseId int)

	InfoBarSetMessageType func(
		infoBar *T.GtkInfoBar,
		messageType T.GtkMessageType)

	InfoBarGetMessageType func(
		infoBar *T.GtkInfoBar) T.GtkMessageType

	InvisibleNewForScreen func(
		screen *T.GdkScreen) *T.GtkWidget

	InvisibleSetScreen func(
		invisible *T.GtkInvisible,
		screen *T.GdkScreen)

	InvisibleGetScreen func(
		invisible *T.GtkInvisible) *T.GdkScreen

	LayoutNew func(
		hadjustment *T.GtkAdjustment,
		vadjustment *T.GtkAdjustment) *T.GtkWidget

	LayoutGetBinWindow func(
		layout *T.GtkLayout) *T.GdkWindow

	LayoutPut func(
		layout *T.GtkLayout,
		childWidget *T.GtkWidget,
		x int,
		y int)

	LayoutMove func(
		layout *T.GtkLayout,
		childWidget *T.GtkWidget,
		x int,
		y int)

	LayoutSetSize func(
		layout *T.GtkLayout,
		width uint,
		height uint)

	LayoutGetSize func(
		layout *T.GtkLayout,
		width *uint,
		height *uint)

	LayoutGetHadjustment func(
		layout *T.GtkLayout) *T.GtkAdjustment

	LayoutGetVadjustment func(
		layout *T.GtkLayout) *T.GtkAdjustment

	LayoutSetHadjustment func(
		layout *T.GtkLayout,
		adjustment *T.GtkAdjustment)

	LayoutSetVadjustment func(
		layout *T.GtkLayout,
		adjustment *T.GtkAdjustment)

	LayoutFreeze func(
		layout *T.GtkLayout)

	LayoutThaw func(
		layout *T.GtkLayout)

	LinkButtonNew func(
		uri string) *T.GtkWidget

	LinkButtonNewWithLabel func(
		uri string,
		label string) *T.GtkWidget

	LinkButtonGetUri func(
		linkButton *T.GtkLinkButton) string

	LinkButtonSetUri func(
		linkButton *T.GtkLinkButton,
		uri string)

	LinkButtonSetUriHook func(
		f T.GtkLinkButtonUriFunc,
		dataGpointer,
		destroy T.GDestroyNotify) T.GtkLinkButtonUriFunc

	LinkButtonGetVisited func(
		linkButton *T.GtkLinkButton) T.Gboolean

	LinkButtonSetVisited func(
		linkButton *T.GtkLinkButton,
		visited T.Gboolean)

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
		widget *T.GtkWidget)

	GrabGetCurrent func() *T.GtkWidget

	GrabRemove func(
		widget *T.GtkWidget)

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
		event *T.GdkEvent) *T.GtkWidget

	PropagateEvent func(
		widget *T.GtkWidget,
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
		widget *T.GtkWidget,
		tipText string,
		tipPrivate string)

	TooltipsDataGet func(
		widget *T.GtkWidget) *T.GtkTooltipsData

	TooltipsForceWindow func(
		tooltips *T.GtkTooltips)

	TooltipsGetInfoFromTipWindow func(
		tipWindow *T.GtkWindow,
		tooltips **T.GtkTooltips,
		currentWidget **T.GtkWidget) T.Gboolean

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
		widget *T.GtkWidget)

	SizeGroupRemoveWidget func(
		sizeGroup *T.GtkSizeGroup,
		widget *T.GtkWidget)

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
		toolItem *T.GtkToolItem) T.GtkIconSize

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
		toolItem *T.GtkToolItem) *T.GtkWidget

	ToolItemGetProxyMenuItem func(
		toolItem *T.GtkToolItem,
		menuItemId string) *T.GtkWidget

	ToolItemSetProxyMenuItem func(
		toolItem *T.GtkToolItem,
		menuItemId string,
		menuItem *T.GtkWidget)

	ToolItemRebuildMenu func(
		toolItem *T.GtkToolItem)

	ToolItemToolbarReconfigured func(
		toolItem *T.GtkToolItem)

	ToolButtonNew func(
		iconWidget *T.GtkWidget,
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
		iconWidget *T.GtkWidget)

	ToolButtonGetIconWidget func(
		button *T.GtkToolButton) *T.GtkWidget

	ToolButtonSetLabelWidget func(
		button *T.GtkToolButton,
		labelWidget *T.GtkWidget)

	ToolButtonGetLabelWidget func(
		button *T.GtkToolButton) *T.GtkWidget

	MenuToolButtonNew func(
		iconWidget *T.GtkWidget,
		label string) *T.GtkToolItem

	MenuToolButtonNewFromStock func(
		stockId string) *T.GtkToolItem

	MenuToolButtonSetMenu func(
		button *T.GtkMenuToolButton,
		menu *T.GtkWidget)

	MenuToolButtonGetMenu func(
		button *T.GtkMenuToolButton) *T.GtkWidget

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
		flags T.GtkDialogFlags,
		t T.GtkMessageType,
		buttons T.GtkButtonsType,
		messageFormat string,
		v ...VArg) *T.GtkWidget

	MessageDialogNewWithMarkup func(
		parent *T.GtkWindow,
		flags T.GtkDialogFlags,
		t T.GtkMessageType,
		buttons T.GtkButtonsType,
		messageFormat string,
		v ...VArg) *T.GtkWidget

	MessageDialogSetImage func(
		dialog *T.GtkMessageDialog,
		image *T.GtkWidget)

	MessageDialogGetImage func(
		dialog *T.GtkMessageDialog) *T.GtkWidget

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
		messageDialog *T.GtkMessageDialog) *T.GtkWidget

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
		child *T.GtkWidget,
		tabLabel *T.GtkWidget) int

	NotebookAppendPageMenu func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tabLabel *T.GtkWidget,
		menuLabel *T.GtkWidget) int

	NotebookPrependPage func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tabLabel *T.GtkWidget) int

	NotebookPrependPageMenu func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tabLabel *T.GtkWidget,
		menuLabel *T.GtkWidget) int

	NotebookInsertPage func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tabLabel *T.GtkWidget,
		position int) int

	NotebookInsertPageMenu func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tabLabel *T.GtkWidget,
		menuLabel *T.GtkWidget,
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
		pageNum int) *T.GtkWidget

	NotebookGetNPages func(
		notebook *T.GtkNotebook) int

	NotebookPageNum func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) int

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
		child *T.GtkWidget) *T.GtkWidget

	NotebookSetTabLabel func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tabLabel *T.GtkWidget)

	NotebookSetTabLabelText func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tabText string)

	NotebookGetTabLabelText func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) string

	NotebookGetMenuLabel func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) *T.GtkWidget

	NotebookSetMenuLabel func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		menuLabel *T.GtkWidget)

	NotebookSetMenuLabelText func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		menuText string)

	NotebookGetMenuLabelText func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) string

	NotebookQueryTabLabelPacking func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		expand *T.Gboolean,
		fill *T.Gboolean,
		packType *T.GtkPackType)

	NotebookSetTabLabelPacking func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		expand T.Gboolean,
		fill T.Gboolean,
		packType T.GtkPackType)

	NotebookReorderChild func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		position int)

	NotebookGetTabReorderable func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) T.Gboolean

	NotebookSetTabReorderable func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		reorderable T.Gboolean)

	NotebookGetTabDetachable func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) T.Gboolean

	NotebookSetTabDetachable func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		detachable T.Gboolean)

	NotebookGetActionWidget func(
		notebook *T.GtkNotebook,
		packType T.GtkPackType) *T.GtkWidget

	NotebookSetActionWidget func(
		notebook *T.GtkNotebook,
		widget *T.GtkWidget,
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
		socketId T.GdkNativeWindow) *T.GtkWidget

	PlugConstructForDisplay func(
		plug *T.GtkPlug,
		display *T.GdkDisplay,
		socketId T.GdkNativeWindow)

	PlugNewForDisplay func(
		display *T.GdkDisplay,
		socketId T.GdkNativeWindow) *T.GtkWidget

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
		adjustment *T.GtkAdjustment)

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
		adjustment *T.GtkAdjustment) *T.GtkWidget

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
		group *T.GSList) *T.GtkWidget

	RadioButtonNewFromWidget func(
		radioGroupMember *T.GtkRadioButton) *T.GtkWidget

	RadioButtonNewWithLabel func(
		group *T.GSList,
		label string) *T.GtkWidget

	RadioButtonNewWithLabelFromWidget func(
		radioGroupMember *T.GtkRadioButton,
		label string) *T.GtkWidget

	RadioButtonNewWithMnemonic func(
		group *T.GSList,
		label string) *T.GtkWidget

	RadioButtonNewWithMnemonicFromWidget func(
		radioGroupMember *T.GtkRadioButton,
		label string) *T.GtkWidget

	RadioButtonGetGroup func(
		radioButton *T.GtkRadioButton) *T.GSList

	RadioButtonSetGroup func(
		radioButton *T.GtkRadioButton,
		group *T.GSList)

	RadioMenuItemNew func(
		group *T.GSList) *T.GtkWidget

	RadioMenuItemNewWithLabel func(
		group *T.GSList,
		label string) *T.GtkWidget

	RadioMenuItemNewWithMnemonic func(
		group *T.GSList,
		label string) *T.GtkWidget

	RadioMenuItemNewFromWidget func(
		group *T.GtkRadioMenuItem) *T.GtkWidget

	RadioMenuItemNewWithMnemonicFromWidget func(
		group *T.GtkRadioMenuItem,
		label string) *T.GtkWidget

	RadioMenuItemNewWithLabelFromWidget func(
		group *T.GtkRadioMenuItem,
		label string) *T.GtkWidget

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
		stockId string) *T.GtkAction

	RecentActionNewForManager func(
		name string,
		label string,
		tooltip string,
		stockId string,
		manager *T.GtkRecentManager) *T.GtkAction

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
		firstButtonText string, v ...VArg) *T.GtkWidget

	RecentChooserDialogNewForManager func(
		title string, parent *T.GtkWindow,
		manager *T.GtkRecentManager,
		firstButtonText string, v ...VArg) *T.GtkWidget

	RecentChooserMenuNewForManager func(
		manager *T.GtkRecentManager) *T.GtkWidget

	RecentChooserMenuGetShowNumbers func(
		menu *T.GtkRecentChooserMenu) T.Gboolean

	RecentChooserMenuSetShowNumbers func(
		menu *T.GtkRecentChooserMenu,
		showNumbers T.Gboolean)

	RecentChooserWidgetNewForManager func(
		manager *T.GtkRecentManager) *T.GtkWidget

	ScaleButtonNew func(
		size T.GtkIconSize,
		min float64,
		max float64,
		step float64,
		icons **T.Gchar) *T.GtkWidget

	ScaleButtonSetIcons func(
		button *T.GtkScaleButton,
		icons **T.Gchar)

	ScaleButtonGetValue func(
		button *T.GtkScaleButton) float64

	ScaleButtonSetValue func(
		button *T.GtkScaleButton,
		value float64)

	ScaleButtonGetAdjustment func(
		button *T.GtkScaleButton) *T.GtkAdjustment

	ScaleButtonSetAdjustment func(
		button *T.GtkScaleButton,
		adjustment *T.GtkAdjustment)

	ScaleButtonGetPlusButton func(
		button *T.GtkScaleButton) *T.GtkWidget

	ScaleButtonGetMinusButton func(
		button *T.GtkScaleButton) *T.GtkWidget

	ScaleButtonGetPopup func(
		button *T.GtkScaleButton) *T.GtkWidget

	ScaleButtonGetOrientation func(
		button *T.GtkScaleButton) T.GtkOrientation

	ScaleButtonSetOrientation func(
		button *T.GtkScaleButton,
		orientation T.GtkOrientation)

	VscrollbarNew func(
		adjustment *T.GtkAdjustment) *T.GtkWidget

	ViewportNew func(
		hadjustment *T.GtkAdjustment,
		vadjustment *T.GtkAdjustment) *T.GtkWidget

	ViewportGetHadjustment func(
		viewport *T.GtkViewport) *T.GtkAdjustment

	ViewportGetVadjustment func(
		viewport *T.GtkViewport) *T.GtkAdjustment

	ViewportSetHadjustment func(
		viewport *T.GtkViewport,
		adjustment *T.GtkAdjustment)

	ViewportSetVadjustment func(
		viewport *T.GtkViewport,
		adjustment *T.GtkAdjustment)

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
		hadjustment *T.GtkAdjustment,
		vadjustment *T.GtkAdjustment) *T.GtkWidget

	ScrolledWindowSetHadjustment func(
		scrolledWindow *T.GtkScrolledWindow,
		hadjustment *T.GtkAdjustment)

	ScrolledWindowSetVadjustment func(
		scrolledWindow *T.GtkScrolledWindow,
		vadjustment *T.GtkAdjustment)

	ScrolledWindowGetHadjustment func(
		scrolledWindow *T.GtkScrolledWindow) *T.GtkAdjustment

	ScrolledWindowGetVadjustment func(
		scrolledWindow *T.GtkScrolledWindow) *T.GtkAdjustment

	ScrolledWindowGetHscrollbar func(
		scrolledWindow *T.GtkScrolledWindow) *T.GtkWidget

	ScrolledWindowGetVscrollbar func(
		scrolledWindow *T.GtkScrolledWindow) *T.GtkWidget

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
		child *T.GtkWidget)

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
		adjustment *T.GtkAdjustment,
		climbRate float64,
		digits uint)

	SpinButtonNew func(
		adjustment *T.GtkAdjustment,
		climbRate float64,
		digits uint) *T.GtkWidget

	SpinButtonNewWithRange func(
		min float64,
		max float64,
		step float64) *T.GtkWidget

	SpinButtonSetAdjustment func(
		spinButton *T.GtkSpinButton,
		adjustment *T.GtkAdjustment)

	SpinButtonGetAdjustment func(
		spinButton *T.GtkSpinButton) *T.GtkAdjustment

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
		statusbar *T.GtkStatusbar) *T.GtkWidget

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
		statusIcon *T.GtkStatusIcon) T.GtkImageType

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
		rows uint,
		columns uint,
		homogeneous T.Gboolean) *T.GtkWidget

	TableResize func(
		table *T.GtkTable,
		rows uint,
		columns uint)

	TableAttach func(
		table *T.GtkTable,
		child *T.GtkWidget,
		leftAttach uint,
		rightAttach uint,
		topAttach uint,
		bottomAttach uint,
		xoptions T.GtkAttachOptions,
		yoptions T.GtkAttachOptions,
		xpadding uint,
		ypadding uint)

	TableAttachDefaults func(
		table *T.GtkTable,
		widget *T.GtkWidget,
		leftAttach uint,
		rightAttach uint,
		topAttach uint,
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
		clipboard *T.GtkClipboard)

	TextBufferRemoveSelectionClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard)

	TextBufferCutClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard,
		defaultEditable T.Gboolean)

	TextBufferCopyClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard)

	TextBufferPasteClipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard,
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

	TextViewNewWithBuffer func(buffer *T.GtkTextBuffer) *T.GtkWidget

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

	TextViewGetHadjustment func(textView *T.GtkTextView) *T.GtkAdjustment

	TextViewGetVadjustment func(textView *T.GtkTextView) *T.GtkAdjustment

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
		child *T.GtkWidget,
		anchor *T.GtkTextChildAnchor)

	TextViewAddChildInWindow func(textView *T.GtkTextView,
		child *T.GtkWidget,
		whichWindow T.GtkTextWindowType,
		xpos int,
		ypos int)

	TextViewMoveChild func(textView *T.GtkTextView,
		child *T.GtkWidget,
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
		mask *T.GdkBitmap) *T.GtkWidget

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

	ToolbarGetIconSize func(toolbar *T.GtkToolbar) T.GtkIconSize

	ToolbarSetIconSize func(toolbar *T.GtkToolbar,
		iconSize T.GtkIconSize)

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
		icon *T.GtkWidget,
		callback T.GCallback,
		userData T.Gpointer) *T.GtkWidget

	ToolbarPrependItem func(toolbar *T.GtkToolbar,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *T.GtkWidget,
		callback T.GCallback,
		userData T.Gpointer) *T.GtkWidget

	ToolbarInsertItem func(toolbar *T.GtkToolbar,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *T.GtkWidget,
		callback T.GCallback,
		userDataGpointer,
		position int) *T.GtkWidget

	ToolbarInsertStock func(toolbar *T.GtkToolbar,
		stockId string,
		tooltipText string,
		tooltipPrivateText string,
		callback T.GCallback,
		userDataGpointer,
		position int) *T.GtkWidget

	ToolbarAppendSpace func(toolbar *T.GtkToolbar)

	ToolbarPrependSpace func(toolbar *T.GtkToolbar)

	ToolbarInsertSpace func(toolbar *T.GtkToolbar,
		position int)

	ToolbarRemoveSpace func(toolbar *T.GtkToolbar,
		position int)

	ToolbarAppendElement func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *T.GtkWidget,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *T.GtkWidget,
		callback T.GCallback,
		userData T.Gpointer) *T.GtkWidget

	ToolbarPrependElement func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *T.GtkWidget,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *T.GtkWidget,
		callback T.GCallback,
		userData T.Gpointer) *T.GtkWidget

	ToolbarInsertElement func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *T.GtkWidget,
		text string,
		tooltipText string,
		tooltipPrivateText string,
		icon *T.GtkWidget,
		callback T.GCallback,
		userDataGpointer,
		position int) *T.GtkWidget

	ToolbarAppendWidget func(toolbar *T.GtkToolbar,
		widget *T.GtkWidget,
		tooltipText string,
		tooltipPrivateText string)

	ToolbarPrependWidget func(toolbar *T.GtkToolbar,
		widget *T.GtkWidget,
		tooltipText string,
		tooltipPrivateText string)

	ToolbarInsertWidget func(toolbar *T.GtkToolbar,
		widget *T.GtkWidget,
		tooltipText string,
		tooltipPrivateText string,
		position int)

	ToolItemGroupNew func(label string) *T.GtkWidget

	ToolItemGroupSetLabel func(group *T.GtkToolItemGroup,
		label string)

	ToolItemGroupSetLabelWidget func(group *T.GtkToolItemGroup,
		labelWidget *T.GtkWidget)

	ToolItemGroupSetCollapsed func(group *T.GtkToolItemGroup,
		collapsed T.Gboolean)

	ToolItemGroupSetEllipsize func(group *T.GtkToolItemGroup,
		ellipsize T.PangoEllipsizeMode)

	ToolItemGroupSetHeaderRelief func(group *T.GtkToolItemGroup,
		style T.GtkReliefStyle)

	ToolItemGroupGetLabel func(group *T.GtkToolItemGroup) string

	ToolItemGroupGetLabelWidget func(group *T.GtkToolItemGroup) *T.GtkWidget

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
		iconSize T.GtkIconSize)

	ToolPaletteUnsetIconSize func(palette *T.GtkToolPalette)

	ToolPaletteSetStyle func(palette *T.GtkToolPalette,
		style T.GtkToolbarStyle)

	ToolPaletteUnsetStyle func(palette *T.GtkToolPalette)

	ToolPaletteGetIconSize func(palette *T.GtkToolPalette) T.GtkIconSize

	ToolPaletteGetStyle func(palette *T.GtkToolPalette) T.GtkToolbarStyle

	ToolPaletteGetDropItem func(palette *T.GtkToolPalette,
		x int,
		y int) *T.GtkToolItem

	ToolPaletteGetDropGroup func(palette *T.GtkToolPalette,
		x int,
		y int) *T.GtkToolItemGroup

	ToolPaletteGetDragItem func(palette *T.GtkToolPalette,
		selection *T.GtkSelectionData) *T.GtkWidget

	ToolPaletteSetDragSource func(palette *T.GtkToolPalette,
		targets T.GtkToolPaletteDragTargets)

	ToolPaletteAddDragDest func(palette *T.GtkToolPalette,
		widget *T.GtkWidget,
		flags T.GtkDestDefaults,
		targets T.GtkToolPaletteDragTargets,
		actions T.GdkDragAction)

	ToolPaletteGetHadjustment func(palette *T.GtkToolPalette) *T.GtkAdjustment

	ToolPaletteGetVadjustment func(palette *T.GtkToolPalette) *T.GtkAdjustment

	ToolPaletteGetDragTargetItem func() *T.GtkTargetEntry

	ToolPaletteGetDragTargetGroup func() *T.GtkTargetEntry

	ToolShellGetIconSize func(shell *T.GtkToolShell) T.GtkIconSize

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

	TestFindWidget func(widget *T.GtkWidget,
		labelPattern string,
		widgetType T.GType) *T.GtkWidget

	TestCreateWidget func(widgetType T.GType,
		firstPropertyName string, v ...VArg) *T.GtkWidget

	TestCreateSimpleWindow func(windowTitle string,
		dialogText string) *T.GtkWidget

	TestDisplayButtonWindow func(windowTitle string,
		dialogText string, v ...VArg) *T.GtkWidget

	TestSliderSetPerc func(widget *T.GtkWidget,
		percentage float64)

	TestSliderGetValue func(widget *T.GtkWidget) float64

	TestSpinButtonClick func(spinner *T.GtkSpinButton,
		button uint,
		upwards T.Gboolean) T.Gboolean

	TestWidgetClick func(widget *T.GtkWidget,
		button uint,
		modifiers T.GdkModifierType) T.Gboolean

	TestWidgetSendKey func(widget *T.GtkWidget,
		keyval uint,
		modifiers T.GdkModifierType) T.Gboolean

	TestTextSet func(widget *T.GtkWidget,
		string string)

	TestTextGet func(widget *T.GtkWidget) string

	TestFindSibling func(baseWidget *T.GtkWidget,
		widgetType T.GType) *T.GtkWidget

	TestFindLabel func(widget *T.GtkWidget,
		labelPattern string) *T.GtkWidget

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
		actionGroup *T.GtkActionGroup,
		pos int)

	UiManagerRemoveActionGroup func(self *T.GtkUIManager,
		actionGroup *T.GtkActionGroup)

	UiManagerGetActionGroups func(self *T.GtkUIManager) *T.GList

	UiManagerGetAccelGroup func(self *T.GtkUIManager) *T.GtkAccelGroup

	UiManagerGetWidget func(self *T.GtkUIManager,
		path string) *T.GtkWidget

	UiManagerGetToplevels func(self *T.GtkUIManager,
		types T.GtkUIManagerItemType) *T.GSList

	UiManagerGetAction func(self *T.GtkUIManager,
		path string) *T.GtkAction

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

	VbuttonBoxGetLayoutDefault func() T.GtkButtonBoxStyle

	VbuttonBoxSetLayoutDefault func(layout T.GtkButtonBoxStyle)

	VscaleNew func(adjustment *T.GtkAdjustment) *T.GtkWidget

	VscaleNewWithRange func(min float64,
		max float64,
		step float64) *T.GtkWidget

	ClistNew func(columns int) *T.GtkWidget

	ClistNewWithTitles func(columns int,
		titles **T.Gchar) *T.GtkWidget

	ClistSetHadjustment func(clist *T.GtkCList,
		adjustment *T.GtkAdjustment)

	ClistSetVadjustment func(clist *T.GtkCList,
		adjustment *T.GtkAdjustment)

	ClistGetHadjustment func(clist *T.GtkCList) *T.GtkAdjustment

	ClistGetVadjustment func(clist *T.GtkCList) *T.GtkAdjustment

	ClistSetShadowType func(clist *T.GtkCList,
		t T.GtkShadowType)

	ClistSetSelectionMode func(clist *T.GtkCList,
		mode T.GtkSelectionMode)

	ClistSetReorderable func(clist *T.GtkCList,
		reorderable T.Gboolean)

	ClistSetUseDragIcons func(clist *T.GtkCList,
		useIcons T.Gboolean)

	ClistSetButtonActions func(clist *T.GtkCList,
		button uint,
		buttonActions uint8)

	ClistFreeze func(clist *T.GtkCList)

	ClistThaw func(clist *T.GtkCList)

	ClistColumnTitlesShow func(clist *T.GtkCList)

	ClistColumnTitlesHide func(clist *T.GtkCList)

	ClistColumnTitleActive func(clist *T.GtkCList,
		column int)

	ClistColumnTitlePassive func(clist *T.GtkCList,
		column int)

	ClistColumnTitlesActive func(clist *T.GtkCList)

	ClistColumnTitlesPassive func(clist *T.GtkCList)

	ClistSetColumnTitle func(clist *T.GtkCList,
		column int,
		title string)

	ClistGetColumnTitle func(clist *T.GtkCList,
		column int) string

	ClistSetColumnWidget func(clist *T.GtkCList,
		column int,
		widget *T.GtkWidget)

	ClistGetColumnWidget func(clist *T.GtkCList,
		column int) *T.GtkWidget

	ClistSetColumnJustification func(clist *T.GtkCList,
		column int,
		justification T.GtkJustification)

	ClistSetColumnVisibility func(clist *T.GtkCList,
		column int,
		visible T.Gboolean)

	ClistSetColumnResizeable func(clist *T.GtkCList,
		column int,
		resizeable T.Gboolean)

	ClistSetColumnAutoResize func(clist *T.GtkCList,
		column int,
		autoResize T.Gboolean)

	ClistColumnsAutosize func(clist *T.GtkCList) int

	ClistOptimalColumnWidth func(clist *T.GtkCList,
		column int) int

	ClistSetColumnWidth func(clist *T.GtkCList,
		column int,
		width int)

	ClistSetColumnMinWidth func(clist *T.GtkCList,
		column int,
		minWidth int)

	ClistSetColumnMaxWidth func(clist *T.GtkCList,
		column int,
		maxWidth int)

	ClistSetRowHeight func(clist *T.GtkCList,
		height uint)

	ClistMoveto func(clist *T.GtkCList,
		row int,
		column int,
		rowAlign float32,
		colAlign float32)

	ClistRowIsVisible func(clist *T.GtkCList,
		row int) T.GtkVisibility

	ClistGetCellType func(clist *T.GtkCList,
		row int,
		column int) T.GtkCellType

	ClistSetText func(clist *T.GtkCList,
		row int,
		column int,
		text string)

	ClistGetText func(clist *T.GtkCList,
		row int,
		column int,
		text **T.Gchar) int

	ClistSetPixmap func(clist *T.GtkCList,
		row int,
		column int,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	ClistGetPixmap func(clist *T.GtkCList,
		row int,
		column int,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap) int

	ClistSetPixtext func(clist *T.GtkCList,
		row int,
		column int,
		text string,
		spacing uint8,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	ClistGetPixtext func(clist *T.GtkCList,
		row int,
		column int,
		text **T.Gchar,
		spacing *uint8,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap) int

	ClistSetForeground func(clist *T.GtkCList,
		row int,
		color *T.GdkColor)

	ClistSetBackground func(clist *T.GtkCList,
		row int,
		color *T.GdkColor)

	ClistSetCellStyle func(clist *T.GtkCList,
		row int,
		column int,
		style *T.GtkStyle)

	ClistGetCellStyle func(clist *T.GtkCList,
		row int,
		column int) *T.GtkStyle

	ClistSetRowStyle func(clist *T.GtkCList,
		row int,
		style *T.GtkStyle)

	ClistGetRowStyle func(clist *T.GtkCList,
		row int) *T.GtkStyle

	ClistSetShift func(clist *T.GtkCList,
		row int,
		column int,
		vertical int,
		horizontal int)

	ClistSetSelectable func(clist *T.GtkCList,
		row int,
		selectable T.Gboolean)

	ClistGetSelectable func(clist *T.GtkCList,
		row int) T.Gboolean

	ClistPrepend func(clist *T.GtkCList,
		text **T.Gchar) int

	ClistAppend func(clist *T.GtkCList,
		text **T.Gchar) int

	ClistInsert func(clist *T.GtkCList,
		row int,
		text **T.Gchar) int

	ClistRemove func(clist *T.GtkCList,
		row int)

	ClistSetRowData func(clist *T.GtkCList,
		row int,
		data T.Gpointer)

	ClistSetRowDataFull func(clist *T.GtkCList,
		row int,
		dataGpointer,
		destroy T.GDestroyNotify)

	ClistGetRowData func(clist *T.GtkCList,
		row int) T.Gpointer

	ClistFindRowFromData func(clist *T.GtkCList,
		data T.Gpointer) int

	ClistSelectRow func(clist *T.GtkCList,
		row int,
		column int)

	ClistUnselectRow func(clist *T.GtkCList,
		row int,
		column int)

	ClistUndoSelection func(clist *T.GtkCList)

	ClistClear func(clist *T.GtkCList)

	ClistGetSelectionInfo func(clist *T.GtkCList,
		x int,
		y int,
		row *int,
		column *int) int

	ClistSelectAll func(clist *T.GtkCList)

	ClistUnselectAll func(clist *T.GtkCList)

	ClistSwapRows func(clist *T.GtkCList,
		row1 int,
		row2 int)

	ClistRowMove func(clist *T.GtkCList,
		sourceRow int,
		destRow int)

	ClistSetCompareFunc func(clist *T.GtkCList,
		cmpFunc T.GtkCListCompareFunc)

	ClistSetSortColumn func(clist *T.GtkCList,
		column int)

	ClistSetSortType func(clist *T.GtkCList,
		sortType T.GtkSortType)

	ClistSort func(clist *T.GtkCList)

	ClistSetAutoSort func(clist *T.GtkCList,
		autoSort T.Gboolean)

	ComboSetValueInList func(combo *T.GtkCombo,
		val T.Gboolean,
		okIfEmpty T.Gboolean)

	ComboSetUseArrows func(combo *T.GtkCombo,
		val T.Gboolean)

	ComboSetUseArrowsAlways func(combo *T.GtkCombo,
		val T.Gboolean)

	ComboSetCaseSensitive func(combo *T.GtkCombo,
		val T.Gboolean)

	ComboSetItemString func(combo *T.GtkCombo,
		item *T.GtkItem,
		itemValue string)

	ComboSetPopdownStrings func(combo *T.GtkCombo,
		strings *T.GList)

	ComboDisableActivate func(combo *T.GtkCombo)

	CtreeNewWithTitles func(columns int,
		treeColumn int,
		titles **T.Gchar) *T.GtkWidget

	CtreeNew func(columns int,
		treeColumn int) *T.GtkWidget

	CtreeInsertNode func(ctree *T.GtkCTree,
		parent *T.GtkCTreeNode,
		sibling *T.GtkCTreeNode,
		text **T.Gchar,
		spacing uint8,
		pixmapClosed *T.GdkPixmap,
		maskClosed *T.GdkBitmap,
		pixmapOpened *T.GdkPixmap,
		maskOpened *T.GdkBitmap,
		isLeaf T.Gboolean,
		expanded T.Gboolean) *T.GtkCTreeNode

	CtreeRemoveNode func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeInsertGnode func(ctree *T.GtkCTree,
		parent *T.GtkCTreeNode,
		sibling *T.GtkCTreeNode,
		gnode *T.GNode,
		f T.GtkCTreeGNodeFunc,
		data T.Gpointer) *T.GtkCTreeNode

	CtreeExportToGnode func(ctree *T.GtkCTree,
		parent *T.GNode,
		sibling *T.GNode,
		node *T.GtkCTreeNode,
		f T.GtkCTreeGNodeFunc,
		data T.Gpointer) *T.GNode

	CtreePostRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		f T.GtkCTreeFunc,
		data T.Gpointer)

	CtreePostRecursiveToDepth func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		depth int,
		f T.GtkCTreeFunc,
		data T.Gpointer)

	CtreePreRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		f T.GtkCTreeFunc,
		data T.Gpointer)

	CtreePreRecursiveToDepth func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		depth int,
		f T.GtkCTreeFunc,
		data T.Gpointer)

	CtreeIsViewable func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) T.Gboolean

	CtreeLast func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) *T.GtkCTreeNode

	CtreeFindNodePtr func(ctree *T.GtkCTree,
		ctreeRow *T.GtkCTreeRow) *T.GtkCTreeNode

	CtreeNodeNth func(ctree *T.GtkCTree,
		row uint) *T.GtkCTreeNode

	CtreeFind func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		child *T.GtkCTreeNode) T.Gboolean

	CtreeIsAncestor func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		child *T.GtkCTreeNode) T.Gboolean

	CtreeFindByRowData func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		data T.Gpointer) *T.GtkCTreeNode

	CtreeFindAllByRowData func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		data T.Gpointer) *T.GList

	CtreeFindByRowDataCustom func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		dataGpointer,
		f T.GCompareFunc) *T.GtkCTreeNode

	CtreeFindAllByRowDataCustom func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		dataGpointer,
		f T.GCompareFunc) *T.GList

	CtreeIsHotSpot func(ctree *T.GtkCTree,
		x int,
		y int) T.Gboolean

	CtreeMove func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		newParent *T.GtkCTreeNode,
		newSibling *T.GtkCTreeNode)

	CtreeExpand func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeExpandRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeExpandToDepth func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		depth int)

	CtreeCollapse func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeCollapseRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeCollapseToDepth func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		depth int)

	CtreeToggleExpansion func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeToggleExpansionRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeSelect func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeSelectRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeUnselect func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeUnselectRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeRealSelectRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		state int)

	CtreeNodeSetText func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		text string)

	CtreeNodeSetPixmap func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	CtreeNodeSetPixtext func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		text string,
		spacing uint8,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	CtreeSetNodeInfo func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		text string,
		spacing uint8,
		pixmapClosed *T.GdkPixmap,
		maskClosed *T.GdkBitmap,
		pixmapOpened *T.GdkPixmap,
		maskOpened *T.GdkBitmap,
		isLeaf T.Gboolean,
		expanded T.Gboolean)

	CtreeNodeSetShift func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		vertical int,
		horizontal int)

	CtreeNodeSetSelectable func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		selectable T.Gboolean)

	CtreeNodeGetSelectable func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) T.Gboolean

	CtreeNodeGetCellType func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int) T.GtkCellType

	CtreeNodeGetText func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		text **T.Gchar) T.Gboolean

	CtreeNodeGetPixmap func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap) T.Gboolean

	CtreeNodeGetPixtext func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		text **T.Gchar,
		spacing *uint8,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap) T.Gboolean

	CtreeGetNodeInfo func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		text **T.Gchar,
		spacing *uint8,
		pixmapClosed **T.GdkPixmap,
		maskClosed **T.GdkBitmap,
		pixmapOpened **T.GdkPixmap,
		maskOpened **T.GdkBitmap,
		isLeaf *T.Gboolean,
		expanded *T.Gboolean) T.Gboolean

	CtreeNodeSetRowStyle func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		style *T.GtkStyle)

	CtreeNodeGetRowStyle func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) *T.GtkStyle

	CtreeNodeSetCellStyle func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		style *T.GtkStyle)

	CtreeNodeGetCellStyle func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int) *T.GtkStyle

	CtreeNodeSetForeground func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		color *T.GdkColor)

	CtreeNodeSetBackground func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		color *T.GdkColor)

	CtreeNodeSetRowData func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		data T.Gpointer)

	CtreeNodeSetRowDataFull func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		dataGpointer,
		destroy T.GDestroyNotify)

	CtreeNodeGetRowData func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) T.Gpointer

	CtreeNodeMoveto func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column int,
		rowAlign float32,
		colAlign float32)

	CtreeNodeIsVisible func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) T.GtkVisibility

	CtreeSetIndent func(ctree *T.GtkCTree,
		indent int)

	CtreeSetSpacing func(ctree *T.GtkCTree,
		spacing int)

	CtreeSetShowStub func(ctree *T.GtkCTree,
		showStub T.Gboolean)

	CtreeSetLineStyle func(ctree *T.GtkCTree,
		lineStyle T.GtkCTreeLineStyle)

	CtreeSetExpanderStyle func(ctree *T.GtkCTree,
		expanderStyle T.GtkCTreeExpanderStyle)

	CtreeSetDragCompareFunc func(ctree *T.GtkCTree,
		cmpFunc T.GtkCTreeCompareDragFunc)

	CtreeSortNode func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CtreeSortRecursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	CurveReset func(curve *T.GtkCurve)

	CurveSetGamma func(curve *T.GtkCurve,
		gamma float32)

	CurveSetRange func(curve *T.GtkCurve,
		minX float32,
		maxX float32,
		minY float32,
		maxY float32)

	CurveGetVector func(curve *T.GtkCurve,
		veclen int,
		vector *float32)

	CurveSetVector func(curve *T.GtkCurve,
		veclen int,
		vector *float32)

	CurveSetCurveType func(curve *T.GtkCurve,
		t T.GtkCurveType)

	FileSelectionNew func(title string) *T.GtkWidget

	FileSelectionSetFilename func(filesel *T.GtkFileSelection,
		filename string)

	FileSelectionGetFilename func(filesel *T.GtkFileSelection) string

	FileSelectionComplete func(filesel *T.GtkFileSelection,
		pattern string)

	FileSelectionShowFileopButtons func(filesel *T.GtkFileSelection)

	FileSelectionHideFileopButtons func(filesel *T.GtkFileSelection)

	FileSelectionGetSelections func(filesel *T.GtkFileSelection) **T.Gchar

	FileSelectionSetSelectMultiple func(filesel *T.GtkFileSelection,
		selectMultiple T.Gboolean)

	FileSelectionGetSelectMultiple func(filesel *T.GtkFileSelection) T.Gboolean

	ItemFactoryNew func(containerType T.GType,
		path string,
		accelGroup *T.GtkAccelGroup) *T.GtkItemFactory

	ItemFactoryConstruct func(ifactory *T.GtkItemFactory,
		containerType T.GType,
		path string,
		accelGroup *T.GtkAccelGroup)

	ItemFactoryAddForeign func(accelWidget *T.GtkWidget,
		fullPath string,
		accelGroup *T.GtkAccelGroup,
		keyval uint,
		modifiers T.GdkModifierType)

	ItemFactoryFromWidget func(widget *T.GtkWidget) *T.GtkItemFactory

	ItemFactoryPathFromWidget func(widget *T.GtkWidget) string

	ItemFactoryGetItem func(ifactory *T.GtkItemFactory,
		path string) *T.GtkWidget

	ItemFactoryGetWidget func(ifactory *T.GtkItemFactory,
		path string) *T.GtkWidget

	ItemFactoryGetWidgetByAction func(ifactory *T.GtkItemFactory,
		action uint) *T.GtkWidget

	ItemFactoryGetItemByAction func(ifactory *T.GtkItemFactory,
		action uint) *T.GtkWidget

	ItemFactoryCreateItem func(ifactory *T.GtkItemFactory,
		entry *T.GtkItemFactoryEntry,
		callbackDataGpointer,
		callbackType uint)

	ItemFactoryCreateItems func(ifactory *T.GtkItemFactory,
		nEntries uint,
		entries *T.GtkItemFactoryEntry,
		callbackData T.Gpointer)

	ItemFactoryDeleteItem func(ifactory *T.GtkItemFactory,
		path string)

	ItemFactoryDeleteEntry func(ifactory *T.GtkItemFactory,
		entry *T.GtkItemFactoryEntry)

	ItemFactoryDeleteEntries func(ifactory *T.GtkItemFactory,
		nEntries uint,
		entries *T.GtkItemFactoryEntry)

	ItemFactoryPopup func(ifactory *T.GtkItemFactory,
		x uint,
		y uint,
		mouseButton uint,
		time T.GUint32)

	ItemFactoryPopupWithData func(ifactory *T.GtkItemFactory,
		popupDataGpointer,
		destroy T.GDestroyNotify,
		x uint,
		y uint,
		mouseButton uint,
		time T.GUint32)

	ItemFactoryPopupData func(ifactory *T.GtkItemFactory) T.Gpointer

	ItemFactoryPopupDataFromWidget func(widget *T.GtkWidget) T.Gpointer

	ItemFactorySetTranslateFunc func(ifactory *T.GtkItemFactory,
		f T.GtkTranslateFunc,
		dataGpointer,
		notify T.GDestroyNotify)

	ItemFactoryCreateItemsAc func(ifactory *T.GtkItemFactory,
		nEntries uint,
		entries *T.GtkItemFactoryEntry,
		callbackDataGpointer,
		callbackType uint)

	ItemFactoryFromPath func(path string) *T.GtkItemFactory

	ItemFactoryCreateMenuEntries func(nEntries uint,
		entries *T.GtkMenuEntry)

	ItemFactoriesPathDelete func(ifactoryPath string,
		path string)

	OldEditableClaimSelection func(oldEditable *T.GtkOldEditable,
		claim T.Gboolean,
		time T.GUint32)

	OldEditableChanged func(oldEditable *T.GtkOldEditable)

	OptionMenuGetMenu func(optionMenu *T.GtkOptionMenu) *T.GtkWidget

	OptionMenuSetMenu func(optionMenu *T.GtkOptionMenu,
		menu *T.GtkWidget)

	OptionMenuRemoveMenu func(optionMenu *T.GtkOptionMenu)

	OptionMenuGetHistory func(optionMenu *T.GtkOptionMenu) int

	OptionMenuSetHistory func(optionMenu *T.GtkOptionMenu,
		index uint)

	PreviewUninit func()

	PreviewNew func(t T.GtkPreviewType) *T.GtkWidget

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
		caller *T.GtkWidget)

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

	BuilderErrorQuark func() T.GQuark

	PrintContextSetCairoContext func(
		context *T.GtkPrintContext,
		cr *T.Cairo,
		dpiX, dpiY float64)

	TreeAppend func(tree *T.GtkTree, treeItem *T.GtkWidget)

	TreePrepend func(tree *T.GtkTree, treeItem *T.GtkWidget)

	TreeInsert func(
		tree *T.GtkTree, treeItem *T.GtkWidget, position int)

	TreeRemoveItems func(tree *T.GtkTree, items *T.GList)

	TreeClearItems func(
		tree *T.GtkTree, start int, end int)

	TreeSelectItem func(tree *T.GtkTree, item int)

	TreeUnselectItem func(tree *T.GtkTree, item int)

	TreeSelectChild func(
		tree *T.GtkTree, treeItem *T.GtkWidget)

	TreeUnselectChild func(
		tree *T.GtkTree, treeItem *T.GtkWidget)

	TreeChildPosition func(
		tree *T.GtkTree, child *T.GtkWidget) int

	TreeSetSelectionMode func(
		tree *T.GtkTree, mode T.GtkSelectionMode)

	TreeSetViewMode func(
		tree *T.GtkTree, mode T.GtkTreeViewMode)

	TreeSetViewLines func(tree *T.GtkTree, flag T.Gboolean)

	TreeRemoveItem func(tree *T.GtkTree, child *T.GtkWidget)

	TreeItemNewWithLabel func(label string) *T.GtkWidget

	TreeItemSetSubtree func(
		treeItem *T.GtkTreeItem, subtree *T.GtkWidget)

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
		title string, cell *T.GtkCellRenderer,
		v ...VArg) *TreeViewColumn

	TreeViewColumnPackStart func(
		treeColumn *TreeViewColumn,
		cell *T.GtkCellRenderer,
		expand T.Gboolean)

	TreeViewColumnPackEnd func(
		treeColumn *TreeViewColumn,
		cell *T.GtkCellRenderer,
		expand T.Gboolean)

	TreeViewColumnClear func(
		treeColumn *TreeViewColumn)

	TreeViewColumnGetCellRenderers func(
		treeColumn *TreeViewColumn) *T.GList

	TreeViewColumnAddAttribute func(
		treeColumn *TreeViewColumn,
		cellRenderer *T.GtkCellRenderer,
		attribute string,
		column int)

	TreeViewColumnSetAttributes func(
		treeColumn *TreeViewColumn,
		cellRenderer *T.GtkCellRenderer, v ...VArg)

	TreeViewColumnSetCellDataFunc func(
		treeColumn *TreeViewColumn,
		cellRenderer *T.GtkCellRenderer,
		f T.GtkTreeCellDataFunc,
		funcDataGpointer,
		destroy T.GDestroyNotify)

	TreeViewColumnClearAttributes func(
		treeColumn *TreeViewColumn,
		cellRenderer *T.GtkCellRenderer)

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
		widget *T.GtkWidget)

	TreeViewColumnGetWidget func(
		treeColumn *TreeViewColumn) *T.GtkWidget

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
		cell *T.GtkCellRenderer)

	TreeViewColumnCellGetPosition func(
		treeColumn *TreeViewColumn,
		cellRenderer *T.GtkCellRenderer,
		startPos *int,
		width *int) T.Gboolean

	TreeViewColumnQueueResize func(
		treeColumn *TreeViewColumn)

	TreeViewColumnGetTreeView func(
		treeColumn *TreeViewColumn) *T.GtkWidget

	// funcs with methods

	ListInsertItems func(list *List,
		items *T.GList,
		position int)

	ListAppendItems func(list *List,
		items *T.GList)

	ListPrependItems func(list *List,
		items *T.GList)

	ListRemoveItems func(list *List,
		items *T.GList)

	ListRemoveItemsNoUnref func(list *List,
		items *T.GList)

	ListClearItems func(list *List,
		start int,
		end int)

	ListSelectItem func(list *List,
		item int)

	ListUnselectItem func(list *List,
		item int)

	ListSelectChild func(list *List,
		child *T.GtkWidget)

	ListUnselectChild func(list *List,
		child *T.GtkWidget)

	ListChildPosition func(list *List,
		child *T.GtkWidget) int

	ListSetSelectionMode func(list *List,
		mode T.GtkSelectionMode)

	ListExtendSelection func(list *List,
		scrollType T.GtkScrollType,
		position float32,
		autoStartSelection T.Gboolean)

	ListStartSelection func(list *List)

	ListEndSelection func(list *List)

	ListSelectAll func(list *List)

	ListUnselectAll func(list *List)

	ListScrollHorizontal func(list *List,
		scrollType T.GtkScrollType,
		position float32)

	ListScrollVertical func(list *List,
		scrollType T.GtkScrollType,
		position float32)

	ListToggleAddMode func(list *List)

	ListToggleFocusRow func(list *List)

	ListToggleRow func(list *List,
		item *T.GtkWidget)

	ListUndoSelection func(list *List)

	ListEndDragSelection func(list *List)

	ListItemNewWithLabel func(label string) *T.GtkWidget

	ListItemSelect func(listItem *ListItem)

	ListItemDeselect func(listItem *ListItem)
)

type (
	List struct {
		Container        T.GtkContainer
		Children         *T.GList
		Selection        *T.GList
		Undo_selection   *T.GList
		Undo_unselection *T.GList
		Last_focus_child *T.GtkWidget
		Undo_focus_child *T.GtkWidget
		Htimer           uint
		Vtimer           uint
		Anchor           int
		Drag_pos         int
		Anchor_state     T.GtkStateType
		Bits             uint
		// selection_mode : 2
		// drag_selection : 1
		// add_mode : 1
	}

	ListItem struct {
		Item T.GtkItem
	}
)

type TreeViewColumn struct{}

func (tc *TreeViewColumn) PackStart(
	cell *T.GtkCellRenderer, expand T.Gboolean) {
	TreeViewColumnPackEnd(tc, cell, expand)
}

func (tc *TreeViewColumn) Clear() { TreeViewColumnClear(tc) }

func (tc *TreeViewColumn) GetCellRenderers() *T.GList {
	return TreeViewColumnGetCellRenderers(tc)
}

func (tc *TreeViewColumn) AddAttribute(
	cellRenderer *T.GtkCellRenderer,
	attribute string, column int) {
	TreeViewColumnAddAttribute(
		tc, cellRenderer, attribute, column)
}

func (tc *TreeViewColumn) SetAttributes(
	cellRenderer *T.GtkCellRenderer, v ...VArg) {
	TreeViewColumnSetAttributes(tc, cellRenderer, v)
}

func (tc *TreeViewColumn) SetCellDataFunc(
	cellRenderer *T.GtkCellRenderer, f T.GtkTreeCellDataFunc,
	funcDataGpointer, destroy T.GDestroyNotify) {
	TreeViewColumnSetCellDataFunc(
		tc, cellRenderer, f, funcDataGpointer, destroy)
}

func (tc *TreeViewColumn) ClearAttributes(
	cellRenderer *T.GtkCellRenderer) {
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

func (tc *TreeViewColumn) SetWidget(widget *T.GtkWidget) {
	TreeViewColumnSetWidget(tc, widget)
}

func (tc *TreeViewColumn) GetWidget() *T.GtkWidget {
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

func (tc *TreeViewColumn) FocusCell(cell *T.GtkCellRenderer) {
	TreeViewColumnFocusCell(tc, cell)
}

func (tc *TreeViewColumn) CellGetPosition(
	cellRenderer *T.GtkCellRenderer, startPos *int,
	width *int) T.Gboolean {
	return TreeViewColumnCellGetPosition(
		tc, cellRenderer, startPos, width)
}

func (tc *TreeViewColumn) QueueResize() {
	TreeViewColumnQueueResize(tc)
}

func (tc *TreeViewColumn) GetTreeView() *T.GtkWidget {
	return TreeViewColumnGetTreeView(tc)
}

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

func (l *List) SelectChild(child *T.GtkWidget) {
	ListSelectChild(l, child)
}

func (l *List) UnselectChild(child *T.GtkWidget) {
	ListUnselectChild(l, child)
}

func (l *List) ChildPosition(child *T.GtkWidget) int {
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

func (l *List) ToggleRow(item *T.GtkWidget) {
	ListToggleRow(l, item)
}

func (l *List) UndoSelection() { ListUndoSelection(l) }

func (l *List) EndDragSelection() { ListEndDragSelection(l) }

func (i *ListItem) ItemSelect() { ListItemSelect(i) }

func (i *ListItem) ItemDeselect() { ListItemDeselect(i) }
