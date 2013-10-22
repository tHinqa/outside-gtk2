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
	AnchorTypeGetType           func() T.GType
	ArgFlagsGetType             func() T.GType
	AttachOptionsGetType        func() T.GType
	CornerTypeGetType           func() T.GType
	DebugFlagGetType            func() T.GType
	DeleteTypeGetType           func() T.GType
	DestDefaultsGetType         func() T.GType
	DirectionTypeGetType        func() T.GType
	DragResultGetType           func() T.GType
	GammaCurveGetType           func() T.GType
	GammaCurveNew               func() *Widget
	HbuttonBoxGetType           func() T.GType
	HbuttonBoxNew               func() *Widget
	HbuttonBoxGetSpacingDefault func() int
	HbuttonBoxGetLayoutDefault  func() ButtonBoxStyle
	HbuttonBoxSetSpacingDefault func(spacing int)
	HbuttonBoxSetLayoutDefault  func(layout ButtonBoxStyle)
	HpanedGetType               func() T.GType
	HpanedNew                   func() *Widget
	HrulerGetType               func() T.GType
	HrulerNew                   func() *Widget
	HscaleGetType               func() T.GType
	HscaleNew                   func(adjustment *Adjustment) *Widget
	HscaleNewWithRange          func(min, max, step float64) *Widget
	HscrollbarGetType           func() T.GType
	HscrollbarNew               func(adjustment *Adjustment) *Widget
	HseparatorGetType           func() T.GType
	HseparatorNew               func() *Widget
	IconLookupFlagsGetType      func() T.GType
	IdentifierGetType           func() T.GType
	ImageTypeGetType            func() T.GType
	ImPreeditStyleGetType       func() T.GType
	ImStatusStyleGetType        func() T.GType
	InputDialogGetType          func() T.GType
	InputDialogNew              func() *Widget
	JustificationGetType        func() T.GType
	MatchTypeGetType            func() T.GType
	MessageDialogGetType        func() T.GType
	MessageTypeGetType          func() T.GType
	MetricTypeGetType           func() T.GType
	MountOperationGetType       func() T.GType
	MovementStepGetType         func() T.GType
	NotebookGetType             func() T.GType
	NotebookNew                 func() *Widget
	NotebookTabGetType          func() T.GType
	NumberUpLayoutGetType       func() T.GType
	ObjectFlagsGetType          func() T.GType
	ObjectGetType               func() T.GType
	OffscreenWindowGetType      func() T.GType
	OffscreenWindowNew          func() *Widget
	OldEditableGetType          func() T.GType
	OptionMenuGetType           func() T.GType
	OptionMenuNew               func() *Widget
	OrientableGetType           func() T.GType
	OrientationGetType          func() T.GType
	PackDirectionGetType        func() T.GType
	PackTypeGetType             func() T.GType
	PageOrientationGetType      func() T.GType
	PageSetGetType              func() T.GType
	PathPriorityTypeGetType     func() T.GType
	PathTypeGetType             func() T.GType
	PolicyTypeGetType           func() T.GType
	PositionTypeGetType         func() T.GType
	PrintDuplexGetType          func() T.GType
	PrintErrorGetType           func() T.GType
	PrintPagesGetType           func() T.GType
	PrintQualityGetType         func() T.GType
	PrintStatusGetType          func() T.GType
	PrivateFlagsGetType         func() T.GType
	RadioActionGetType          func() T.GType
	RadioButtonGetType          func() T.GType
	RadioMenuItemGetType        func() T.GType
	RadioToolButtonGetType      func() T.GType
	RangeGetType                func() T.GType
	RcFlagsGetType              func() T.GType
	RcStyleGetType              func() T.GType
	RcTokenTypeGetType          func() T.GType
	RecentSortTypeGetType       func() T.GType
	ReliefStyleGetType          func() T.GType
	RequisitionGetType          func() T.GType
	ResizeModeGetType           func() T.GType
	ResponseTypeGetType         func() T.GType
	RulerGetType                func() T.GType
	ScrollbarGetType            func() T.GType
	ScrollStepGetType           func() T.GType
	ScrollTypeGetType           func() T.GType
	TreeGetRowDragData          func(s *SelectionData, treeModel **TreeModel, path **TreePath) T.Gboolean
	TreeSetRowDragData          func(s *SelectionData, treeModel *TreeModel, path *TreePath) T.Gboolean
	SelectionModeGetType        func() T.GType
	SensitivityTypeGetType      func() T.GType
	SeparatorGetType            func() T.GType
	SeparatorMenuItemGetType    func() T.GType
	SeparatorMenuItemNew        func() *Widget
	ShadowTypeGetType           func() T.GType
	SideTypeGetType             func() T.GType
	SignalRunTypeGetType        func() T.GType
	SortTypeGetType             func() T.GType
	SpinTypeGetType             func() T.GType
	StateTypeGetType            func() T.GType
	SubmenuDirectionGetType     func() T.GType
	SubmenuPlacementGetType     func() T.GType
	TargetFlagsGetType          func() T.GType
	TearoffMenuItemGetType      func() T.GType
	TearoffMenuItemNew          func() *Widget
	TextDirectionGetType        func() T.GType
	TextSearchFlagsGetType      func() T.GType
	TextWindowTypeGetType       func() T.GType
	UiManagerGetType            func() T.GType
	UiManagerItemTypeGetType    func() T.GType
	UnitGetType                 func() T.GType
	UpdateTypeGetType           func() T.GType
	VboxGetType                 func() T.GType
	VbuttonBoxGetType           func() T.GType
	VbuttonBoxNew               func() *Widget
	ViewportGetType             func() T.GType
	VisibilityGetType           func() T.GType
	VolumeButtonGetType         func() T.GType
	VolumeButtonNew             func() *Widget
	VpanedGetType               func() T.GType
	VpanedNew                   func() *Widget
	VrulerGetType               func() T.GType
	VrulerNew                   func() *Widget
	VscaleGetType               func() T.GType
	VscrollbarGetType           func() T.GType
	VseparatorGetType           func() T.GType
	VseparatorNew               func() *Widget
	WindowGroupGetType          func() T.GType
	WindowPositionGetType       func() T.GType
	WindowTypeGetType           func() T.GType
	WrapModeGetType             func() T.GType

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
		data T.Gpointer,
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
		settings *Settings,
		widgetPath string,
		classPath string,
		t T.GType) *T.GtkStyle

	RcReparseAllForSettings func(
		settings *Settings,
		forceLoad T.Gboolean) T.Gboolean

	RcResetStyles func(
		settings *Settings)

	RcFindPixmapInPath func(
		settings *Settings,
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

	RequisitionCopy func(
		requisition *T.GtkRequisition) *T.GtkRequisition

	RequisitionFree func(
		requisition *T.GtkRequisition)

	WindowGroupGetCurrentGrab func(
		windowGroup *T.GtkWindowGroup) *Widget

	AccelGroupsActivate func(object *T.GObject, accelKey uint,
		accelMods T.GdkModifierType) T.Gboolean

	AccelGroupsFromObject func(object *T.GObject) *T.GSList

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
		funcData T.Gpointer,
		aliveObject *T.GtkObject)

	SignalConnectFull func(
		object *T.GtkObject,
		name string,
		f T.GCallback,
		unsupported T.GtkCallbackMarshal,
		data T.Gpointer,
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
		data T.Gpointer,
		match T.GSignalMatchType,
		action uint)

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

	TargetsIncludeText func(
		targets *T.GdkAtom,
		nTargets int) T.Gboolean

	TargetsIncludeRichText func(
		targets *T.GdkAtom,
		nTargets int,
		buffer *TextBuffer) T.Gboolean

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
		targetList *TargetList) T.GdkAtom

	DragDestGetTargetList func(
		widget *Widget) *TargetList

	DragDestSetTargetList func(
		widget *Widget,
		targetList *TargetList)

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
		widget *Widget) *TargetList

	DragSourceSetTargetList func(
		widget *Widget,
		targetList *TargetList)

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
		targets *TargetList,
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

	GcGet func(
		depth int,
		colormap *T.GdkColormap,
		values *T.GdkGCValues,
		valuesMask T.GdkGCValuesMask) *T.GdkGC

	GcRelease func(
		gc *T.GdkGC)

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

	HsvToRgb func(h, s, v float64, r, g, b *float64)

	RgbToHsv func(r, g, b float64, h, s, v *float64)

	IconSizeLookupForSettings func(settings *Settings, size IconSize, width *int, height *int) T.Gboolean

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
		data T.Gpointer,
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
		data T.Gpointer,
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
		data T.Gpointer,
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
		data T.Gpointer,
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

	TooltipsDataGet func(
		widget *Widget) *TooltipsData

	TooltipsGetInfoFromTipWindow func(
		tipWindow *Window,
		tooltips **Tooltips,
		currentWidget **Widget) T.Gboolean

	TooltipTriggerTooltipQuery func(
		display *T.GdkDisplay)

	MessageDialogNew func(
		parent *Window,
		flags DialogFlags,
		t T.GtkMessageType,
		buttons T.GtkButtonsType,
		messageFormat string,
		v ...VArg) *Widget

	MessageDialogNewWithMarkup func(
		parent *Window,
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
		parent *Window) *T.GMountOperation

	MountOperationIsShowing func(
		op *T.GtkMountOperation) T.Gboolean

	MountOperationSetParent func(
		op *T.GtkMountOperation,
		parent *Window)

	MountOperationGetParent func(
		op *T.GtkMountOperation) *Window

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
		data T.Gpointer,
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

	PrintErrorQuark func() T.GQuark

	PrintRunPageSetupDialog func(
		parent *Window,
		pageSetup *T.GtkPageSetup,
		settings *PrintSettings) *T.GtkPageSetup

	PrintRunPageSetupDialogAsync func(
		parent *Window,
		pageSetup *T.GtkPageSetup,
		settings *PrintSettings,
		doneCb T.GtkPageSetupDoneFunc,
		data T.Gpointer)

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

	RadioToolButtonNew func(
		group *T.GSList) *ToolItem

	RadioToolButtonNewFromStock func(
		group *T.GSList,
		stockId string) *ToolItem

	RadioToolButtonNewFromWidget func(
		group *T.GtkRadioToolButton) *ToolItem

	RadioToolButtonNewWithStockFromWidget func(
		group *T.GtkRadioToolButton,
		stockId string) *ToolItem

	RadioToolButtonGetGroup func(
		button *T.GtkRadioToolButton) *T.GSList

	RadioToolButtonSetGroup func(
		button *T.GtkRadioToolButton,
		group *T.GSList)

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

	ShowUri func(
		screen *T.GdkScreen,
		uri string,
		timestamp T.GUint32,
		error **T.GError) T.Gboolean

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
		data T.Gpointer,
		notify T.GDestroyNotify)

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

	TestSpinButtonClick func(spinner *SpinButton,
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
)
