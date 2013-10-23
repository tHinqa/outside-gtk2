// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Bin struct {
	Container Container
	Child     *Widget
}

var BinGetType func() T.GType

var binGetChild func(b *Bin) *Widget

func (b *Bin) GetChild() *Widget { return binGetChild(b) }

type (
	BindingSet struct {
		SetName           *T.Gchar
		Priority          int
		WidgetPathPspecs  *T.GSList
		WidgetClassPspecs *T.GSList
		ClassBranchPspecs *T.GSList
		Entries           *BindingEntry
		Current           *BindingEntry
		Parsed            uint // : 1
	}

	BindingEntry struct {
		Keyval     uint
		Modifiers  T.GdkModifierType
		BindingSet *BindingSet
		Bits       uint
		// Destroyed : 1
		// InEmission : 1
		// MarksUnbound : 1
		SetNext  *BindingEntry
		HashNext *BindingEntry
		Signals  *BindingSignal
	}

	BindingSignal struct {
		Next        *BindingSignal
		Signal_name *T.Gchar
		N_args      uint
		Args        *BindingArg
	}

	BindingArg struct {
		ArgType T.GType
		// Union
		// LongData  Glong
		DoubleData float64 // largest for size
		// StringData  *Gchar
	}
)

var (
	BindingSetNew func(setName string) *BindingSet

	BindingParseBinding   func(scanner *T.GScanner) uint
	BindingsActivate      func(object *Object, keyval uint, modifiers T.GdkModifierType) T.Gboolean
	BindingsActivateEvent func(object *Object, event *T.GdkEventKey) T.Gboolean

	BindingSetByClass func(objectClass T.Gpointer) *BindingSet
	BindingSetFind    func(setName string) *BindingSet

	bindingEntryAddSignal  func(b *BindingSet, keyval uint, modifiers T.GdkModifierType, signalName string, nArgs uint, varg ...VArg)
	bindingEntryAddSignall func(b *BindingSet, keyval uint, modifiers T.GdkModifierType, signalName string, bindingArgs *T.GSList)
	bindingEntryClear      func(b *BindingSet, keyval uint, modifiers T.GdkModifierType)
	bindingEntryRemove     func(b *BindingSet, keyval uint, modifiers T.GdkModifierType)
	bindingEntrySkip       func(b *BindingSet, keyval uint, modifiers T.GdkModifierType)
	bindingSetActivate     func(b *BindingSet, keyval uint, modifiers T.GdkModifierType, object *Object) T.Gboolean
	bindingSetAddPath      func(b *BindingSet, pathType PathType, pathPattern string, priority PathPriorityType)
)

func (b *BindingSet) EntryAddSignal(keyval uint, modifiers T.GdkModifierType, signalName string, nArgs uint, varg ...VArg) {
	bindingEntryAddSignal(b, keyval, modifiers, signalName, nArgs, varg)
}
func (b *BindingSet) EntryAddSignall(keyval uint, modifiers T.GdkModifierType, signalName string, bindingArgs *T.GSList) {
	bindingEntryAddSignall(b, keyval, modifiers, signalName, bindingArgs)
}
func (b *BindingSet) EntryClear(keyval uint, modifiers T.GdkModifierType) {
	bindingEntryClear(b, keyval, modifiers)
}
func (b *BindingSet) EntryRemove(keyval uint, modifiers T.GdkModifierType) {
	bindingEntryRemove(b, keyval, modifiers)
}
func (b *BindingSet) EntrySkip(keyval uint, modifiers T.GdkModifierType) {
	bindingEntrySkip(b, keyval, modifiers)
}
func (b *BindingSet) Activate(keyval uint, modifiers T.GdkModifierType, object *Object) T.Gboolean {
	return bindingSetActivate(b, keyval, modifiers, object)
}
func (b *BindingSet) AddPath(pathType PathType, pathPattern string, priority PathPriorityType) {
	bindingSetAddPath(b, pathType, pathPattern, priority)
}

type Border struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

var (
	BorderGetType func() T.GType
	BorderNew     func() *Border
)

var (
	borderCopy func(b *Border) *Border
	borderFree func(b *Border)
)

func (b *Border) Copy() *Border { return borderCopy(b) }
func (b *Border) Free()         { borderFree(b) }

type Box struct {
	Container   Container
	Children    *T.GList
	Spacing     int16
	Homogeneous uint // : 1
}

var BoxGetType func() T.GType

var (
	boxGetHomogeneous    func(b *Box) T.Gboolean
	boxGetSpacing        func(b *Box) int
	boxPackEnd           func(b *Box, child *Widget, expand T.Gboolean, fill T.Gboolean, padding uint)
	boxPackEndDefaults   func(b *Box, widget *Widget)
	boxPackStart         func(b *Box, child *Widget, expand T.Gboolean, fill T.Gboolean, padding uint)
	boxPackStartDefaults func(b *Box, widget *Widget)
	boxQueryChildPacking func(b *Box, child *Widget, expand *T.Gboolean, fill *T.Gboolean, padding *uint, packType *PackType)
	boxReorderChild      func(b *Box, child *Widget, position int)
	boxSetChildPacking   func(b *Box, child *Widget, expand T.Gboolean, fill T.Gboolean, padding uint, packType PackType)
	boxSetHomogeneous    func(b *Box, homogeneous T.Gboolean)
	boxSetSpacing        func(b *Box, spacing int)
)

func (b *Box) GetHomogeneous() T.Gboolean { return boxGetHomogeneous(b) }
func (b *Box) GetSpacing() int            { return boxGetSpacing(b) }
func (b *Box) PackEnd(child *Widget, expand T.Gboolean, fill T.Gboolean, padding uint) {
	boxPackEnd(b, child, expand, fill, padding)
}
func (b *Box) PackEndDefaults(widget *Widget) { boxPackEndDefaults(b, widget) }
func (b *Box) PackStart(child *Widget, expand T.Gboolean, fill T.Gboolean, padding uint) {
	boxPackStart(b, child, expand, fill, padding)
}
func (b *Box) PackStartDefaults(widget *Widget) { boxPackStartDefaults(b, widget) }
func (b *Box) QueryChildPacking(child *Widget, expand *T.Gboolean, fill *T.Gboolean, padding *uint, packType *PackType) {
	boxQueryChildPacking(b, child, expand, fill, padding, packType)
}
func (b *Box) ReorderChild(child *Widget, position int) { boxReorderChild(b, child, position) }
func (b *Box) SetChildPacking(child *Widget, expand T.Gboolean, fill T.Gboolean, padding uint, packType PackType) {
	boxSetChildPacking(b, child, expand, fill, padding, packType)
}
func (b *Box) SetHomogeneous(homogeneous T.Gboolean) { boxSetHomogeneous(b, homogeneous) }
func (b *Box) SetSpacing(spacing int)                { boxSetSpacing(b, spacing) }

type Buildable struct{}

var BuildableGetType func() T.GType

var (
	buildableAddChild             func(b *Buildable, builder *Builder, child *T.GObject, typ string)
	buildableConstructChild       func(b *Buildable, builder *Builder, name string) *T.GObject
	buildableCustomFinished       func(b *Buildable, builder *Builder, child *T.GObject, tagname string, data T.Gpointer)
	buildableCustomTagEnd         func(b *Buildable, builder *Builder, child *T.GObject, tagname string, data *T.Gpointer)
	buildableCustomTagStart       func(b *Buildable, builder *Builder, child *T.GObject, tagname string, parser *T.GMarkupParser, data *T.Gpointer) T.Gboolean
	buildableGetInternalChild     func(b *Buildable, builder *Builder, childname string) *T.GObject
	buildableGetName              func(b *Buildable) string
	buildableParserFinished       func(b *Buildable, builder *Builder)
	buildableSetBuildableProperty func(b *Buildable, builder *Builder, name string, value *T.GValue)
	buildableSetName              func(b *Buildable, name string)
)

func (b *Buildable) AddChild(builder *Builder, child *T.GObject, typ string) {
	buildableAddChild(b, builder, child, typ)
}
func (b *Buildable) ConstructChild(builder *Builder, name string) *T.GObject {
	return buildableConstructChild(b, builder, name)
}
func (b *Buildable) CustomFinished(builder *Builder, child *T.GObject, tagname string, data T.Gpointer) {
	buildableCustomFinished(b, builder, child, tagname, data)
}
func (b *Buildable) CustomTagEnd(builder *Builder, child *T.GObject, tagname string, data *T.Gpointer) {
	buildableCustomTagEnd(b, builder, child, tagname, data)
}
func (b *Buildable) CustomTagStart(builder *Builder, child *T.GObject, tagname string, parser *T.GMarkupParser, data *T.Gpointer) T.Gboolean {
	return buildableCustomTagStart(b, builder, child, tagname, parser, data)
}
func (b *Buildable) GetInternalChild(builder *Builder, childname string) *T.GObject {
	return buildableGetInternalChild(b, builder, childname)
}
func (b *Buildable) GetName() string                 { return buildableGetName(b) }
func (b *Buildable) ParserFinished(builder *Builder) { buildableParserFinished(b, builder) }
func (b *Buildable) SetBuildableProperty(builder *Builder, name string, value *T.GValue) {
	buildableSetBuildableProperty(b, builder, name, value)
}
func (b *Buildable) SetName(name string) { buildableSetName(b, name) }

type (
	Builder struct {
		Parent T.GObject
		_      *struct{}
	}

	BuilderConnectFunc func(builder *Builder, object *T.GObject,
		signalName, handlerName string, connectObject *T.GObject,
		flags T.GConnectFlags, userData T.Gpointer)
)

var (
	BuilderGetType func() T.GType
	BuilderNew     func() *Builder

	BuilderErrorGetType func() T.GType
	BuilderErrorQuark   func() T.GQuark
)

var (
	builderAddFromFile          func(b *Builder, filename string, err **T.GError) uint
	builderAddFromString        func(b *Builder, buffer string, length T.Gsize, err **T.GError) uint
	builderAddObjectsFromFile   func(b *Builder, filename string, objectIds **T.Gchar, err **T.GError) uint
	builderAddObjectsFromString func(b *Builder, buffer string, length T.Gsize, objectIds **T.Gchar, err **T.GError) uint
	builderConnectSignals       func(b *Builder, userData T.Gpointer)
	builderConnectSignalsFull   func(b *Builder, f BuilderConnectFunc, userData T.Gpointer)
	builderGetObject            func(b *Builder, name string) *T.GObject
	builderGetObjects           func(b *Builder) *T.GSList
	builderGetTranslationDomain func(b *Builder) string
	builderGetTypeFromName      func(b *Builder, typeName string) T.GType
	builderSetTranslationDomain func(b *Builder, domain string)
	builderValueFromString      func(b *Builder, pspec *T.GParamSpec, str string, value *T.GValue, err **T.GError) T.Gboolean
	builderValueFromStringType  func(b *Builder, t T.GType, str string, value *T.GValue, err **T.GError) T.Gboolean
)

func (b *Builder) AddFromFile(filename string, err **T.GError) uint {
	return builderAddFromFile(b, filename, err)
}
func (b *Builder) AddFromString(buffer string, length T.Gsize, err **T.GError) uint {
	return builderAddFromString(b, buffer, length, err)
}
func (b *Builder) AddObjectsFromFile(filename string, objectIds **T.Gchar, err **T.GError) uint {
	return builderAddObjectsFromFile(b, filename, objectIds, err)
}
func (b *Builder) AddObjectsFromString(buffer string, length T.Gsize, objectIds **T.Gchar, err **T.GError) uint {
	return builderAddObjectsFromString(b, buffer, length, objectIds, err)
}
func (b *Builder) ConnectSignals(userData T.Gpointer) { builderConnectSignals(b, userData) }
func (b *Builder) ConnectSignalsFull(f BuilderConnectFunc, userData T.Gpointer) {
	builderConnectSignalsFull(b, f, userData)
}
func (b *Builder) GetObject(name string) *T.GObject        { return builderGetObject(b, name) }
func (b *Builder) GetObjects() *T.GSList                   { return builderGetObjects(b) }
func (b *Builder) GetTranslationDomain() string            { return builderGetTranslationDomain(b) }
func (b *Builder) GetTypeFromName(typeName string) T.GType { return builderGetTypeFromName(b, typeName) }
func (b *Builder) SetTranslationDomain(domain string)      { builderSetTranslationDomain(b, domain) }
func (b *Builder) ValueFromString(pspec *T.GParamSpec, str string, value *T.GValue, err **T.GError) T.Gboolean {
	return builderValueFromString(b, pspec, str, value, err)
}
func (b *Builder) ValueFromStringType(t T.GType, str string, value *T.GValue, err **T.GError) T.Gboolean {
	return builderValueFromStringType(b, t, str, value, err)
}

type Button struct {
	Bin             Bin
	EventWindow     *D.Window
	LabelText       *T.Gchar
	ActivateTimeout uint
	Bits            uint
	// Constructed : 1
	// InButton : 1
	// ButtonDown : 1
	// Relief : 2
	// UseUnderline : 1
	// UseStock : 1
	// Depressed : 1
	// DepressOnActivate : 1
	// FocusOnClick : 1
}

var (
	ButtonGetType         func() T.GType
	ButtonNew             func() *Widget
	ButtonNewFromStock    func(stockId string) *Widget
	ButtonNewWithLabel    func(label string) *Widget
	ButtonNewWithMnemonic func(label string) *Widget

	ButtonActionGetType   func() T.GType
	ButtonBoxGetType      func() T.GType
	ButtonBoxStyleGetType func() T.GType

	ButtonsTypeGetType func() T.GType
)

var (
	buttonClicked          func(b *Button)
	buttonEnter            func(b *Button)
	buttonGetAlignment     func(b *Button, xalign, yalign *float32)
	buttonGetEventWindow   func(b *Button) *D.Window
	buttonGetFocusOnClick  func(b *Button) T.Gboolean
	buttonGetImage         func(b *Button) *Widget
	buttonGetImagePosition func(b *Button) PositionType
	buttonGetLabel         func(b *Button) string
	buttonGetRelief        func(b *Button) ReliefStyle
	buttonGetUseStock      func(b *Button) T.Gboolean
	buttonGetUseUnderline  func(b *Button) T.Gboolean
	buttonLeave            func(b *Button)
	buttonPressed          func(b *Button)
	buttonReleased         func(b *Button)
	buttonSetAlignment     func(b *Button, xalign, yalign float32)
	buttonSetFocusOnClick  func(b *Button, focusOnClick T.Gboolean)
	buttonSetImage         func(b *Button, image *Widget)
	buttonSetImagePosition func(b *Button, position PositionType)
	buttonSetLabel         func(b *Button, label string)
	buttonSetRelief        func(b *Button, newstyle ReliefStyle)
	buttonSetUseStock      func(b *Button, useStock T.Gboolean)
	buttonSetUseUnderline  func(b *Button, useUnderline T.Gboolean)
)

func (b *Button) Clicked()                                { buttonClicked(b) }
func (b *Button) Enter()                                  { buttonEnter(b) }
func (b *Button) GetAlignment(xalign, yalign *float32)    { buttonGetAlignment(b, xalign, yalign) }
func (b *Button) GetEventWindow() *D.Window               { return buttonGetEventWindow(b) }
func (b *Button) GetFocusOnClick() T.Gboolean             { return buttonGetFocusOnClick(b) }
func (b *Button) GetImage() *Widget                       { return buttonGetImage(b) }
func (b *Button) GetImagePosition() PositionType          { return buttonGetImagePosition(b) }
func (b *Button) GetLabel() string                        { return buttonGetLabel(b) }
func (b *Button) GetRelief() ReliefStyle                  { return buttonGetRelief(b) }
func (b *Button) GetUseStock() T.Gboolean                 { return buttonGetUseStock(b) }
func (b *Button) GetUseUnderline() T.Gboolean             { return buttonGetUseUnderline(b) }
func (b *Button) Leave()                                  { buttonLeave(b) }
func (b *Button) Pressed()                                { buttonPressed(b) }
func (b *Button) Released()                               { buttonReleased(b) }
func (b *Button) SetAlignment(xalign, yalign float32)     { buttonSetAlignment(b, xalign, yalign) }
func (b *Button) SetFocusOnClick(focusOnClick T.Gboolean) { buttonSetFocusOnClick(b, focusOnClick) }
func (b *Button) SetImage(image *Widget)                  { buttonSetImage(b, image) }
func (b *Button) SetImagePosition(position PositionType)  { buttonSetImagePosition(b, position) }
func (b *Button) SetLabel(label string)                   { buttonSetLabel(b, label) }
func (b *Button) SetRelief(newstyle ReliefStyle)          { buttonSetRelief(b, newstyle) }
func (b *Button) SetUseStock(useStock T.Gboolean)         { buttonSetUseStock(b, useStock) }
func (b *Button) SetUseUnderline(useUnderline T.Gboolean) { buttonSetUseUnderline(b, useUnderline) }

type ButtonBox struct {
	Box            Box
	ChildMinWidth  int
	ChildMinHeight int
	ChildIpadX     int
	ChildIpadY     int
	LayoutStyle    ButtonBoxStyle
}

type ButtonBoxStyle T.Enum

const (
	BUTTONBOX_DEFAULT_STYLE ButtonBoxStyle = iota
	BUTTONBOX_SPREAD
	BUTTONBOX_EDGE
	BUTTONBOX_START
	BUTTONBOX_END
	BUTTONBOX_CENTER
)

var (
	buttonBoxGetChildIpadding  func(b *ButtonBox, ipadX, ipadY *int)
	buttonBoxGetChildSecondary func(b *ButtonBox, child *Widget) T.Gboolean
	buttonBoxGetChildSize      func(b *ButtonBox, minWidth, minHeight *int)
	buttonBoxGetLayout         func(b *ButtonBox) ButtonBoxStyle
	buttonBoxSetChildIpadding  func(b *ButtonBox, ipadX, ipadY int)
	buttonBoxSetChildSecondary func(b *ButtonBox, child *Widget, isSecondary T.Gboolean)
	buttonBoxSetChildSize      func(b *ButtonBox, minWidth, minHeight int)
	buttonBoxSetLayout         func(b *ButtonBox, layoutStyle ButtonBoxStyle)
)

func (b *ButtonBox) GetChildIpadding(ipadX, ipadY *int) { buttonBoxGetChildIpadding(b, ipadX, ipadY) }
func (b *ButtonBox) GetChildSecondary(child *Widget) T.Gboolean {
	return buttonBoxGetChildSecondary(b, child)
}
func (b *ButtonBox) GetChildSize(minWidth, minHeight *int) {
	buttonBoxGetChildSize(b, minWidth, minHeight)
}
func (b *ButtonBox) GetLayout() ButtonBoxStyle         { return buttonBoxGetLayout(b) }
func (b *ButtonBox) SetChildIpadding(ipadX, ipadY int) { buttonBoxSetChildIpadding(b, ipadX, ipadY) }
func (b *ButtonBox) SetChildSecondary(child *Widget, isSecondary T.Gboolean) {
	buttonBoxSetChildSecondary(b, child, isSecondary)
}
func (b *ButtonBox) SetChildSize(minWidth, minHeight int) {
	buttonBoxSetChildSize(b, minWidth, minHeight)
}
func (b *ButtonBox) SetLayout(layoutStyle ButtonBoxStyle) { buttonBoxSetLayout(b, layoutStyle) }

type ButtonsType T.Enum

const (
	BUTTONS_NONE ButtonsType = iota
	BUTTONS_OK
	BUTTONS_CLOSE
	BUTTONS_CANCEL
	BUTTONS_YES_NO
	BUTTONS_OK_CANCEL
)
