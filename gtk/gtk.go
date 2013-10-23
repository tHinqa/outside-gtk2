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
	PrintErrorGetType        func() T.GType
	PrivateFlagsGetType      func() T.GType
	RecentSortTypeGetType    func() T.GType
	ResponseTypeGetType      func() T.GType
	ScrollbarGetType         func() T.GType
	ScrollStepGetType        func() T.GType
	TreeGetRowDragData       func(s *SelectionData, treeModel **TreeModel, path **TreePath) T.Gboolean
	TreeSetRowDragData       func(s *SelectionData, treeModel *TreeModel, path *TreePath) T.Gboolean
	SeparatorGetType         func() T.GType
	SeparatorMenuItemGetType func() T.GType
	SeparatorMenuItemNew     func() *Widget
	SideTypeGetType          func() T.GType
	SignalRunTypeGetType     func() T.GType
	SortTypeGetType          func() T.GType
	SubmenuDirectionGetType  func() T.GType
	SubmenuPlacementGetType  func() T.GType
	TargetFlagsGetType       func() T.GType
	TearoffMenuItemGetType   func() T.GType
	TearoffMenuItemNew       func() *Widget

	TypeInit func(
		debugFlags T.GTypeDebugFlags)

	TypeUnique func(
		parentType Type,
		gtkinfo *TypeInfo) Type

	TypeClass func(
		t Type) T.Gpointer

	TypeNew func(
		t Type) T.Gpointer

	ObjectSink func(
		object *Object)

	ObjectDestroy func(
		object *Object)

	ObjectNew func(t T.GType, firstPropertyName string,
		v ...VArg) *Object

	ObjectRef func(
		object *Object) *Object

	ObjectUnref func(
		object *Object)

	ObjectWeakref func(
		object *Object,
		notify T.GDestroyNotify,
		data T.Gpointer)

	ObjectWeakunref func(
		object *Object,
		notify T.GDestroyNotify,
		data T.Gpointer)

	ObjectSetData func(
		object *Object,
		key string,
		data T.Gpointer)

	ObjectSetDataFull func(
		object *Object,
		key string,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	ObjectRemoveData func(
		object *Object,
		key string)

	ObjectGetData func(
		object *Object,
		key string) T.Gpointer

	ObjectRemoveNoNotify func(
		object *Object,
		key string)

	ObjectSetUserData func(
		object *Object,
		data T.Gpointer)

	ObjectGetUserData func(
		object *Object) T.Gpointer

	ObjectSetDataById func(
		object *Object,
		dataId T.GQuark,
		data T.Gpointer)

	ObjectSetDataByIdFull func(
		object *Object,
		dataId T.GQuark,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	ObjectGetDataById func(
		object *Object,
		dataId T.GQuark) T.Gpointer

	ObjectRemoveDataById func(
		object *Object,
		dataId T.GQuark)

	ObjectRemoveNoNotifyById func(
		object *Object,
		keyId T.GQuark)

	ObjectGet func(object *Object,
		firstPropertyName string, v ...VArg)

	ObjectSet func(object *Object,
		firstPropertyName string, v ...VArg)

	ObjectAddArgType func(
		argName string,
		argType T.GType,
		argFlags, argId uint)

	DrawHline func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		x1, x2, y int)

	DrawVline func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		y1, y2, x int)

	DrawShadow func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawPolygon func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		points *T.GdkPoint,
		npoints int,
		fill T.Gboolean)

	DrawArrow func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		arrowType ArrowType,
		fill T.Gboolean,
		x, y, width, height int)

	DrawDiamond func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawBox func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawFlatBox func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawCheck func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawOption func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawTab func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawShadowGap func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		gapSide PositionType,
		gapX, gapWidth int)

	DrawBoxGap func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		gapSide PositionType,
		gapX, gapWidth int)

	DrawExtension func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		gapSide PositionType)

	DrawFocus func(
		style *Style,
		window *T.GdkWindow,
		x, y, width, height int)

	DrawSlider func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		orientation Orientation)

	DrawHandle func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		orientation Orientation)

	DrawExpander func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		x, y int,
		expanderStyle ExpanderStyle)

	DrawLayout func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		useText T.Gboolean,
		x, y int,
		layout *T.PangoLayout)

	DrawResizeGrip func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		edge T.GdkWindowEdge,
		x, y, width, height int)

	PaintHline func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x1, x2, y int)

	PaintVline func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		y1, y2, x int)

	PaintShadow func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintPolygon func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		points *T.GdkPoint,
		nPoints int,
		fill T.Gboolean)

	PaintArrow func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		arrowType ArrowType,
		fill T.Gboolean,
		x, y, width, height int)

	PaintDiamond func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintBox func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintFlatBox func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintCheck func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintOption func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintTab func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintShadowGap func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		gapSide PositionType,
		gapX, gapWidth int)

	PaintBoxGap func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		gapSide PositionType,
		gapX, gapWidth int)

	PaintExtension func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		gapSide PositionType)

	PaintFocus func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintSlider func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		orientation Orientation)

	PaintHandle func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		shadowType ShadowType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		orientation Orientation)

	PaintExpander func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y int,
		expanderStyle ExpanderStyle)

	PaintLayout func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		useText T.Gboolean,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y int,
		layout *T.PangoLayout)

	PaintResizeGrip func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		edge T.GdkWindowEdge,
		x, y, width, height int)

	PaintSpinner func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		step uint,
		x, y, width, height int)

	StyleGetStyleProperty func(
		style *Style,
		widgetType T.GType,
		propertyName string,
		value *T.GValue)

	StyleGetValist func(
		style *Style,
		widgetType T.GType,
		firstPropertyName string,
		varArgs T.VaList)

	StyleGet func(style *Style, widgetType T.GType,
		firstPropertyName string, v ...VArg)

	DrawString func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		x, y int,
		str string)

	PaintString func(
		style *Style,
		window *T.GdkWindow,
		stateType StateType,
		area *T.GdkRectangle,
		widget *Widget,
		detail string,
		x, y int,
		str string)

	DrawInsertionCursor func(
		widget *Widget,
		drawable *T.GdkDrawable,
		area, location *T.GdkRectangle,
		isPrimary T.Gboolean,
		direction TextDirection,
		drawArrow T.Gboolean)

	RequisitionCopy func(
		requisition *Requisition) *Requisition

	RequisitionFree func(
		requisition *Requisition)

	AccelGroupsActivate func(object *T.GObject, accelKey uint,
		accelMods T.GdkModifierType) T.Gboolean

	AccelGroupsFromObject func(object *T.GObject) *T.GSList

	SignalNewv func(
		name string,
		signalFlags SignalRunType,
		objectType T.GType,
		functionOffset uint,
		marshaller T.GSignalCMarshaller,
		returnVal T.GType,
		nArgs uint,
		args *T.GType) uint

	SignalNew func(name string,
		signalFlags SignalRunType,
		objectType T.GType, functionOffset uint,
		marshaller T.GSignalCMarshaller, returnVal T.GType,
		nArgs uint, v ...VArg) uint

	SignalEmitStopByName func(
		object *Object, name string)

	SignalConnectObjectWhileAlive func(
		object *Object,
		name string,
		f T.GCallback,
		aliveObject *Object)

	SignalConnectWhileAlive func(
		object *Object,
		name string,
		f T.GCallback,
		funcData T.Gpointer,
		aliveObject *Object)

	SignalConnectFull func(
		object *Object,
		name string,
		f T.GCallback,
		unsupported CallbackMarshal,
		data T.Gpointer,
		destroyFunc T.GDestroyNotify,
		objectSignal, after int) T.Gulong

	SignalEmitv func(object *Object, signalId uint, args *Arg)

	SignalEmit func(object *Object, signalId uint, v ...VArg)

	SignalEmitByName func(object *Object, name string, v ...VArg)

	SignalEmitvByName func(object *Object, name string, args *Arg)

	SignalCompatMatched func(
		object *Object,
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
		selection, target T.GdkAtom,
		info uint)

	SelectionAddTargets func(
		widget *Widget,
		selection T.GdkAtom,
		targets *TargetEntry,
		ntargets uint)

	SelectionClearTargets func(
		widget *Widget,
		selection T.GdkAtom)

	SelectionConvert func(
		widget *Widget,
		selection, target T.GdkAtom,
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

	DragGetData func(
		widget *Widget,
		context *T.GdkDragContext,
		target T.GdkAtom,
		time T.GUint32)

	DragFinish func(
		context *T.GdkDragContext,
		success, del T.Gboolean,
		time T.GUint32)

	DragGetSourceWidget func(
		context *T.GdkDragContext) *Widget

	DragHighlight func(
		widget *Widget)

	DragUnhighlight func(
		widget *Widget)

	DragDestSet func(
		widget *Widget,
		flags DestDefaults,
		targets *TargetEntry,
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
		targets *TargetEntry,
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
		hotX, hotY int)

	DragSetIconPixmap func(
		context *T.GdkDragContext,
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap,
		hotX, hotY int)

	DragSetIconPixbuf func(
		context *T.GdkDragContext,
		pixbuf *T.GdkPixbuf,
		hotX, hotY int)

	DragSetIconStock func(
		context *T.GdkDragContext,
		stockId string,
		hotX, hotY int)

	DragSetIconName func(
		context *T.GdkDragContext,
		iconName string,
		hotX, hotY int)

	DragSetIconDefault func(
		context *T.GdkDragContext)

	DragCheckThreshold func(
		widget *Widget,
		startX, startY, currentX, currentY int) T.Gboolean

	DragSetDefaultIcon func(
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap,
		hotX, hotY int)

	GcGet func(
		depth int,
		colormap *T.GdkColormap,
		values *T.GdkGCValues,
		valuesMask T.GdkGCValuesMask) *T.GdkGC

	GcRelease func(
		gc *T.GdkGC)

	RgbToHsv func(r, g, b float64, h, s, v *float64)

	IconSizeLookupForSettings func(settings *Settings, size IconSize, width *int, height *int) T.Gboolean

	ParseArgs func(argc *int, argv ***T.Char) T.Gboolean

	Init func(argc *int, argv ***T.Char)

	InitCheck func(argc *int, argv ***T.Char) T.Gboolean

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
		sizeof_GtkWindow, sizeof_GtkBox T.SizeT)

	InitCheckAbiCheck func(
		argc *int,
		argv ***T.Char,
		numChecks int,
		sizeof_GtkWindow, sizeof_GtkBox T.SizeT) T.Gboolean

	Exit func(
		errorCode int)

	SetLocale func() string

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
		function Function,
		data T.Gpointer)

	QuitAddDestroy func(
		mainLevel uint,
		object *Object)

	QuitAdd func(
		mainLevel uint,
		function Function,
		data T.Gpointer) uint

	QuitAddFull func(
		mainLevel uint,
		function Function,
		marshal CallbackMarshal,
		data T.Gpointer,
		destroy T.GDestroyNotify) uint

	QuitRemove func(
		quitHandlerId uint)

	QuitRemoveByData func(
		data T.Gpointer)

	TimeoutAdd func(
		interval T.GUint32,
		function Function,
		data T.Gpointer) uint

	TimeoutAddFull func(
		interval T.GUint32,
		function Function,
		marshal CallbackMarshal,
		data T.Gpointer,
		destroy T.GDestroyNotify) uint

	TimeoutRemove func(
		timeoutHandlerId uint)

	IdleAdd func(
		function Function,
		data T.Gpointer) uint

	IdleAddPriority func(
		priority int,
		function Function,
		data T.Gpointer) uint

	IdleAddFull func(
		priority int,
		function Function,
		marshal CallbackMarshal,
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
		marshal CallbackMarshal,
		data T.Gpointer,
		destroy T.GDestroyNotify) uint

	InputRemove func(
		inputHandlerId uint)

	KeySnooperInstall func(
		snooper KeySnoopFunc,
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

	PrintErrorQuark func() T.GQuark

	PrintRunPageSetupDialog func(
		parent *Window,
		pageSetup *PageSetup,
		settings *PrintSettings) *PageSetup

	PrintRunPageSetupDialogAsync func(
		parent *Window,
		pageSetup *PageSetup,
		settings *PrintSettings,
		doneCb PageSetupDoneFunc,
		data T.Gpointer)

	ShowUri func(
		screen *T.GdkScreen,
		uri string,
		timestamp T.GUint32,
		error **T.GError) T.Gboolean

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
		str string)

	TestTextGet func(widget *Widget) string

	TestFindSibling func(baseWidget *Widget,
		widgetType T.GType) *Widget

	TestFindLabel func(widget *Widget,
		labelPattern string) *Widget

	Marshal_BOOLEAN__VOID func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_POINTER_INT_INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_INT_INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_INT_INT_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_STRING_STRING_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_ENUM__ENUM func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_INT__POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_INT__POINTER_CHAR_CHAR func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__ENUM_FLOAT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__ENUM_FLOAT_BOOLEAN func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__INT_INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__INT_INT_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_INT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_STRING_STRING func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT_ENUM func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER_UINT_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_INT_INT_POINTER_UINT_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT_UINT func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__STRING_INT_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__UINT_POINTER_UINT_ENUM_ENUM_POINTER func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__UINT_POINTER_UINT_UINT_ENUM func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__UINT_STRING func(
		closure *T.GClosure,
		returnValue *T.GValue,
		nParamValues uint,
		paramValues *T.GValue,
		invocationHint, marshalData T.Gpointer)

	TypeEnumGetValues func(enumType Type) *EnumValue

	TypeFlagsGetValues func(flagsType Type) *FlagValue

	TypeEnumFindValue func(
		enumType Type,
		valueName string) *EnumValue

	TypeFlagsFindValue func(
		flagsType Type,
		valueName string) *FlagValue
)
