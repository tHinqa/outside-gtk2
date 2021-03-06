// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gtk provides API definitions for accessing
//libgtk-win32-2.0-0.dll.
package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Enum int

var (
	True  func() bool
	False func() bool
)

var (
	DrawHline func(
		style *Style,
		window *D.Window,
		stateType StateType,
		x1, x2, y int)

	DrawVline func(
		style *Style,
		window *D.Window,
		stateType StateType,
		y1, y2, x int)

	DrawShadow func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawPolygon func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		points *D.Point,
		npoints int,
		fill bool)

	DrawArrow func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		arrowType ArrowType,
		fill bool,
		x, y, width, height int)

	DrawDiamond func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawBox func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawFlatBox func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawCheck func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawOption func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawTab func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int)

	DrawShadowGap func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		gapSide PositionType,
		gapX, gapWidth int)

	DrawBoxGap func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		gapSide PositionType,
		gapX, gapWidth int)

	DrawExtension func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		gapSide PositionType)

	DrawFocus func(
		style *Style,
		window *D.Window,
		x, y, width, height int)

	DrawSlider func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		orientation Orientation)

	DrawHandle func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		x, y, width, height int,
		orientation Orientation)

	DrawExpander func(
		style *Style,
		window *D.Window,
		stateType StateType,
		x, y int,
		expanderStyle ExpanderStyle)

	DrawLayout func(
		style *Style,
		window *D.Window,
		stateType StateType,
		useText bool,
		x, y int,
		layout *P.Layout)

	DrawResizeGrip func(
		style *Style,
		window *D.Window,
		stateType StateType,
		edge D.WindowEdge,
		x, y, width, height int)

	PaintHline func(
		style *Style,
		window *D.Window,
		stateType StateType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x1, x2, y int)

	PaintVline func(
		style *Style,
		window *D.Window,
		stateType StateType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		y1, y2, x int)

	PaintShadow func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintPolygon func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		points *D.Point,
		nPoints int,
		fill bool)

	PaintArrow func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		arrowType ArrowType,
		fill bool,
		x, y, width, height int)

	PaintDiamond func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintBox func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintFlatBox func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintCheck func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintOption func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintTab func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintShadowGap func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		gapSide PositionType,
		gapX, gapWidth int)

	PaintBoxGap func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		gapSide PositionType,
		gapX, gapWidth int)

	PaintExtension func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		gapSide PositionType)

	PaintFocus func(
		style *Style,
		window *D.Window,
		stateType StateType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int)

	PaintSlider func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		orientation Orientation)

	PaintHandle func(
		style *Style,
		window *D.Window,
		stateType StateType,
		shadowType ShadowType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y, width, height int,
		orientation Orientation)

	PaintExpander func(
		style *Style,
		window *D.Window,
		stateType StateType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y int,
		expanderStyle ExpanderStyle)

	PaintLayout func(
		style *Style,
		window *D.Window,
		stateType StateType,
		useText bool,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y int,
		layout *P.Layout)

	PaintResizeGrip func(
		style *Style,
		window *D.Window,
		stateType StateType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		edge D.WindowEdge,
		x, y, width, height int)

	PaintSpinner func(
		style *Style,
		window *D.Window,
		stateType StateType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		step uint,
		x, y, width, height int)

	DrawString func(
		style *Style,
		window *D.Window,
		stateType StateType,
		x, y int,
		str string)

	PaintString func(
		style *Style,
		window *D.Window,
		stateType StateType,
		area *D.Rectangle,
		widget *Widget,
		detail string,
		x, y int,
		str string)

	DrawInsertionCursor func(
		widget *Widget,
		drawable *D.Drawable,
		area, location *D.Rectangle,
		isPrimary bool,
		direction TextDirection,
		drawArrow bool)

	SignalNewv func(
		name string,
		signalFlags SignalRunType,
		objectType O.Type,
		functionOffset uint,
		marshaller O.SignalCMarshaller,
		returnVal O.Type,
		nArgs uint,
		args *O.Type) uint

	SignalNew func(name string,
		signalFlags SignalRunType,
		objectType O.Type, functionOffset uint,
		marshaller O.SignalCMarshaller, returnVal O.Type,
		nArgs uint, v ...VArg) uint

	SignalEmitStopByName func(object *Object, name string)

	SignalConnectObjectWhileAlive func(
		object *Object,
		name string,
		f O.Callback,
		aliveObject *Object)

	SignalConnectWhileAlive func(
		object *Object,
		name string,
		f O.Callback,
		funcData T.Gpointer,
		aliveObject *Object)

	SignalConnectFull func(
		object *Object,
		name string,
		f O.Callback,
		unsupported CallbackMarshal,
		data T.Gpointer,
		destroyFunc O.DestroyNotify,
		objectSignal, after int) T.Gulong

	SignalEmitv func(object *Object, signalId uint, args *Arg)

	SignalEmit func(object *Object, signalId uint, v ...VArg)

	SignalEmitByName func(object *Object, name string, v ...VArg)

	SignalEmitvByName func(object *Object, name string, args *Arg)

	SignalCompatMatched func(
		object *Object,
		f O.Callback,
		data T.Gpointer,
		match O.SignalMatchType,
		action uint)

	GcGet func(
		depth int,
		colormap *D.Colormap,
		values *D.GCValues,
		valuesMask D.GCValuesMask) *D.GC

	GcRelease func(gc *D.GC)

	RgbToHsv func(r, g, b float64, h, s, v *float64)

	IconSizeLookupForSettings func(settings *Settings, size IconSize, width *int, height *int) bool

	ParseArgs func(argc *int, argv ***T.Char) bool

	Init func(argc *int, argv ***T.Char)

	InitCheck func(argc *int, argv ***T.Char) bool

	InitWithArgs func(
		argc *int,
		argv ***T.Char,
		parameterString string,
		entries *L.OptionEntry,
		translationDomain string,
		error **L.Error) bool

	GetOptionGroup func(
		openDefaultDisplay bool) *L.OptionGroup

	InitAbiCheck func(
		argc *int,
		argv ***T.Char,
		numChecks int,
		sizeofGtkWindow, sizeofGtkBox T.SizeT)

	InitCheckAbiCheck func(
		argc *int,
		argv ***T.Char,
		numChecks int,
		sizeof_GtkWindow, sizeof_GtkBox T.SizeT) bool

	Exit func(errorCode int)

	SetLocale func() string

	GetDefaultLanguage func() *P.Language

	EventsPending func() bool

	MainDoEvent func(event *D.Event)

	Main func()

	MainLevel func() uint

	MainQuit func()

	MainIteration func() bool

	MainIterationDo func(
		blocking bool) bool

	GrabAdd func(widget *Widget)

	GrabGetCurrent func() *Widget

	GrabRemove func(widget *Widget)

	InitAdd func(function Function, data T.Gpointer)

	QuitAddDestroy func(mainLevel uint, object *Object)

	QuitAdd func(
		mainLevel uint, function Function, data T.Gpointer) uint

	QuitAddFull func(
		mainLevel uint,
		function Function,
		marshal CallbackMarshal,
		data T.Gpointer,
		destroy O.DestroyNotify) uint

	QuitRemove func(quitHandlerId uint)

	QuitRemoveByData func(data T.Gpointer)

	IdleAdd func(function Function, data T.Gpointer) uint

	IdleAddPriority func(
		priority int, function Function, data T.Gpointer) uint

	IdleAddFull func(
		priority int,
		function Function,
		marshal CallbackMarshal,
		data T.Gpointer,
		destroy O.DestroyNotify) uint

	IdleRemove func(idleHandlerId uint)

	IdleRemoveByData func(data T.Gpointer)

	InputAddFull func(
		source int,
		condition T.GdkInputCondition,
		function T.GdkInputFunction,
		marshal CallbackMarshal,
		data T.Gpointer,
		destroy O.DestroyNotify) uint

	InputRemove func(inputHandlerId uint)

	KeySnooperInstall func(
		snooper KeySnoopFunc, funcData T.Gpointer) uint

	KeySnooperRemove func(snooperHandlerId uint)

	GetCurrentEvent func() *D.Event

	GetCurrentEventTime func() T.GUint32

	GetCurrentEventState func(state *T.GdkModifierType) bool

	GetEventWidget func(event *D.Event) *Widget

	PropagateEvent func(widget *Widget, event *D.Event)

	ShowUri func(
		screen *D.Screen,
		uri string,
		timestamp T.GUint32,
		error **L.Error) bool

	TestInit func(argcp *int, argvp ***T.Char, v ...VArg)

	TestRegisterAllTypes func()

	TestListAllTypes func(nTypes *uint) *O.Type

	TestFindWidget func(widget *Widget,
		labelPattern string,
		widgetType O.Type) *Widget

	TestCreateWidget func(widgetType O.Type,
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
		upwards bool) bool

	TestWidgetClick func(widget *Widget,
		button uint,
		modifiers T.GdkModifierType) bool

	TestWidgetSendKey func(widget *Widget,
		keyval uint,
		modifiers T.GdkModifierType) bool

	TestTextSet func(widget *Widget,
		str string)

	TestTextGet func(widget *Widget) string

	TestFindSibling func(baseWidget *Widget,
		widgetType O.Type) *Widget

	TestFindLabel func(widget *Widget,
		labelPattern string) *Widget

	Marshal_BOOLEAN__VOID func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_POINTER_INT_INT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_INT_INT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_INT_INT_UINT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_BOOLEAN__POINTER_STRING_STRING_POINTER func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_ENUM__ENUM func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_INT__POINTER func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_INT__POINTER_CHAR_CHAR func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__ENUM_FLOAT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__ENUM_FLOAT_BOOLEAN func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__INT_INT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__INT_INT_POINTER func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_INT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER_POINTER func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_STRING_STRING func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT_ENUM func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_POINTER_UINT_UINT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_INT_INT_POINTER_UINT_UINT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__POINTER_UINT_UINT func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__STRING_INT_POINTER func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__UINT_POINTER_UINT_ENUM_ENUM_POINTER func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__UINT_POINTER_UINT_UINT_ENUM func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)

	Marshal_VOID__UINT_STRING func(
		closure *O.Closure,
		returnValue *O.Value,
		nParamValues uint,
		paramValues *O.Value,
		invocationHint, marshalData T.Gpointer)
)
