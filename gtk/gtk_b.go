// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Bin struct {
	Container Container
	Child     *Widget
}

var BinGetType func() O.Type

var BinGetChild func(b *Bin) *Widget

func (b *Bin) GetChild() *Widget { return BinGetChild(b) }

type (
	BindingSet struct {
		SetName           *T.Gchar
		Priority          int
		WidgetPathPspecs  *L.SList
		WidgetClassPspecs *L.SList
		ClassBranchPspecs *L.SList
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
		ArgType O.Type
		// Union
		// LongData  Glong
		DoubleData float64 // largest for size
		// StringData  *Gchar
	}
)

var (
	BindingSetNew func(setName string) *BindingSet

	BindingParseBinding   func(scanner *L.Scanner) uint
	BindingsActivate      func(object *Object, keyval uint, modifiers T.GdkModifierType) bool
	BindingsActivateEvent func(object *Object, event *D.EventKey) bool

	BindingSetByClass func(objectClass T.Gpointer) *BindingSet
	BindingSetFind    func(setName string) *BindingSet

	BindingEntryAddSignal  func(b *BindingSet, keyval uint, modifiers T.GdkModifierType, signalName string, nArgs uint, varg ...VArg)
	BindingEntryAddSignall func(b *BindingSet, keyval uint, modifiers T.GdkModifierType, signalName string, bindingArgs *L.SList)
	BindingEntryClear      func(b *BindingSet, keyval uint, modifiers T.GdkModifierType)
	BindingEntryRemove     func(b *BindingSet, keyval uint, modifiers T.GdkModifierType)
	BindingEntrySkip       func(b *BindingSet, keyval uint, modifiers T.GdkModifierType)
	BindingSetActivate     func(b *BindingSet, keyval uint, modifiers T.GdkModifierType, object *Object) bool
	BindingSetAddPath      func(b *BindingSet, pathType PathType, pathPattern string, priority PathPriorityType)
)

func (b *BindingSet) EntryAddSignal(keyval uint, modifiers T.GdkModifierType, signalName string, nArgs uint, varg ...VArg) {
	BindingEntryAddSignal(b, keyval, modifiers, signalName, nArgs, varg)
}
func (b *BindingSet) EntryAddSignall(keyval uint, modifiers T.GdkModifierType, signalName string, bindingArgs *L.SList) {
	BindingEntryAddSignall(b, keyval, modifiers, signalName, bindingArgs)
}
func (b *BindingSet) EntryClear(keyval uint, modifiers T.GdkModifierType) {
	BindingEntryClear(b, keyval, modifiers)
}
func (b *BindingSet) EntryRemove(keyval uint, modifiers T.GdkModifierType) {
	BindingEntryRemove(b, keyval, modifiers)
}
func (b *BindingSet) EntrySkip(keyval uint, modifiers T.GdkModifierType) {
	BindingEntrySkip(b, keyval, modifiers)
}
func (b *BindingSet) Activate(keyval uint, modifiers T.GdkModifierType, object *Object) bool {
	return BindingSetActivate(b, keyval, modifiers, object)
}
func (b *BindingSet) AddPath(pathType PathType, pathPattern string, priority PathPriorityType) {
	BindingSetAddPath(b, pathType, pathPattern, priority)
}

type Border struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

var (
	BorderGetType func() O.Type
	BorderNew     func() *Border
)

var (
	BorderCopy func(b *Border) *Border
	BorderFree func(b *Border)
)

func (b *Border) Copy() *Border { return BorderCopy(b) }
func (b *Border) Free()         { BorderFree(b) }

type Box struct {
	Container   Container
	Children    *L.List
	Spacing     int16
	Homogeneous uint // : 1
}

var BoxGetType func() O.Type

var (
	BoxGetHomogeneous    func(b *Box) bool
	BoxGetSpacing        func(b *Box) int
	BoxPackEnd           func(b *Box, child *Widget, expand bool, fill bool, padding uint)
	BoxPackEndDefaults   func(b *Box, widget *Widget)
	BoxPackStart         func(b *Box, child *Widget, expand bool, fill bool, padding uint)
	BoxPackStartDefaults func(b *Box, widget *Widget)
	BoxQueryChildPacking func(b *Box, child *Widget, expand *bool, fill *bool, padding *uint, packType *PackType)
	BoxReorderChild      func(b *Box, child *Widget, position int)
	BoxSetChildPacking   func(b *Box, child *Widget, expand bool, fill bool, padding uint, packType PackType)
	BoxSetHomogeneous    func(b *Box, homogeneous bool)
	BoxSetSpacing        func(b *Box, spacing int)
)

func (b *Box) GetHomogeneous() bool { return BoxGetHomogeneous(b) }
func (b *Box) GetSpacing() int      { return BoxGetSpacing(b) }
func (b *Box) PackEnd(child *Widget, expand bool, fill bool, padding uint) {
	BoxPackEnd(b, child, expand, fill, padding)
}
func (b *Box) PackEndDefaults(widget *Widget) { BoxPackEndDefaults(b, widget) }
func (b *Box) PackStart(child *Widget, expand bool, fill bool, padding uint) {
	BoxPackStart(b, child, expand, fill, padding)
}
func (b *Box) PackStartDefaults(widget *Widget) { BoxPackStartDefaults(b, widget) }
func (b *Box) QueryChildPacking(child *Widget, expand *bool, fill *bool, padding *uint, packType *PackType) {
	BoxQueryChildPacking(b, child, expand, fill, padding, packType)
}
func (b *Box) ReorderChild(child *Widget, position int) { BoxReorderChild(b, child, position) }
func (b *Box) SetChildPacking(child *Widget, expand bool, fill bool, padding uint, packType PackType) {
	BoxSetChildPacking(b, child, expand, fill, padding, packType)
}
func (b *Box) SetHomogeneous(homogeneous bool) { BoxSetHomogeneous(b, homogeneous) }
func (b *Box) SetSpacing(spacing int)          { BoxSetSpacing(b, spacing) }

type Buildable struct{}

var BuildableGetType func() O.Type

var (
	BuildableAddChild             func(b *Buildable, builder *Builder, child *O.Object, typ string)
	BuildableConstructChild       func(b *Buildable, builder *Builder, name string) *O.Object
	BuildableCustomFinished       func(b *Buildable, builder *Builder, child *O.Object, tagname string, data T.Gpointer)
	BuildableCustomTagEnd         func(b *Buildable, builder *Builder, child *O.Object, tagname string, data *T.Gpointer)
	BuildableCustomTagStart       func(b *Buildable, builder *Builder, child *O.Object, tagname string, parser *L.MarkupParser, data *T.Gpointer) bool
	BuildableGetInternalChild     func(b *Buildable, builder *Builder, childname string) *O.Object
	BuildableGetName              func(b *Buildable) string
	BuildableParserFinished       func(b *Buildable, builder *Builder)
	BuildableSetBuildableProperty func(b *Buildable, builder *Builder, name string, value *O.Value)
	BuildableSetName              func(b *Buildable, name string)
)

func (b *Buildable) AddChild(builder *Builder, child *O.Object, typ string) {
	BuildableAddChild(b, builder, child, typ)
}
func (b *Buildable) ConstructChild(builder *Builder, name string) *O.Object {
	return BuildableConstructChild(b, builder, name)
}
func (b *Buildable) CustomFinished(builder *Builder, child *O.Object, tagname string, data T.Gpointer) {
	BuildableCustomFinished(b, builder, child, tagname, data)
}
func (b *Buildable) CustomTagEnd(builder *Builder, child *O.Object, tagname string, data *T.Gpointer) {
	BuildableCustomTagEnd(b, builder, child, tagname, data)
}
func (b *Buildable) CustomTagStart(builder *Builder, child *O.Object, tagname string, parser *L.MarkupParser, data *T.Gpointer) bool {
	return BuildableCustomTagStart(b, builder, child, tagname, parser, data)
}
func (b *Buildable) GetInternalChild(builder *Builder, childname string) *O.Object {
	return BuildableGetInternalChild(b, builder, childname)
}
func (b *Buildable) GetName() string                 { return BuildableGetName(b) }
func (b *Buildable) ParserFinished(builder *Builder) { BuildableParserFinished(b, builder) }
func (b *Buildable) SetBuildableProperty(builder *Builder, name string, value *O.Value) {
	BuildableSetBuildableProperty(b, builder, name, value)
}
func (b *Buildable) SetName(name string) { BuildableSetName(b, name) }

type (
	Builder struct {
		Parent O.Object
		_      *struct{}
	}

	BuilderConnectFunc func(builder *Builder, object *O.Object,
		signalName, handlerName string, connectObject *O.Object,
		flags T.GConnectFlags, userData T.Gpointer)
)

var (
	BuilderGetType func() O.Type
	BuilderNew     func() *Builder

	BuilderErrorGetType func() O.Type
	BuilderErrorQuark   func() L.Quark
)

var (
	BuilderAddFromFile          func(b *Builder, filename string, err **L.Error) uint
	BuilderAddFromString        func(b *Builder, buffer string, length T.Gsize, err **L.Error) uint
	BuilderAddObjectsFromFile   func(b *Builder, filename string, objectIds []string, err **L.Error) uint
	BuilderAddObjectsFromString func(b *Builder, buffer string, length T.Gsize, objectIds []string, err **L.Error) uint
	BuilderConnectSignals       func(b *Builder, userData T.Gpointer)
	BuilderConnectSignalsFull   func(b *Builder, f BuilderConnectFunc, userData T.Gpointer)
	BuilderGetObject            func(b *Builder, name string) *O.Object
	BuilderGetObjects           func(b *Builder) *L.SList
	BuilderGetTranslationDomain func(b *Builder) string
	BuilderGetTypeFromName      func(b *Builder, typeName string) O.Type
	BuilderSetTranslationDomain func(b *Builder, domain string)
	BuilderValueFromString      func(b *Builder, pspec *T.GParamSpec, str string, value *O.Value, err **L.Error) bool
	BuilderValueFromStringType  func(b *Builder, t O.Type, str string, value *O.Value, err **L.Error) bool
)

func (b *Builder) AddFromFile(filename string, err **L.Error) uint {
	return BuilderAddFromFile(b, filename, err)
}
func (b *Builder) AddFromString(buffer string, length T.Gsize, err **L.Error) uint {
	return BuilderAddFromString(b, buffer, length, err)
}
func (b *Builder) AddObjectsFromFile(filename string, objectIds []string, err **L.Error) uint {
	return BuilderAddObjectsFromFile(b, filename, objectIds, err)
}
func (b *Builder) AddObjectsFromString(buffer string, length T.Gsize, objectIds []string, err **L.Error) uint {
	return BuilderAddObjectsFromString(b, buffer, length, objectIds, err)
}
func (b *Builder) ConnectSignals(userData T.Gpointer) { BuilderConnectSignals(b, userData) }
func (b *Builder) ConnectSignalsFull(f BuilderConnectFunc, userData T.Gpointer) {
	BuilderConnectSignalsFull(b, f, userData)
}
func (b *Builder) GetObject(name string) *O.Object        { return BuilderGetObject(b, name) }
func (b *Builder) GetObjects() *L.SList                   { return BuilderGetObjects(b) }
func (b *Builder) GetTranslationDomain() string           { return BuilderGetTranslationDomain(b) }
func (b *Builder) GetTypeFromName(typeName string) O.Type { return BuilderGetTypeFromName(b, typeName) }
func (b *Builder) SetTranslationDomain(domain string)     { BuilderSetTranslationDomain(b, domain) }
func (b *Builder) ValueFromString(pspec *T.GParamSpec, str string, value *O.Value, err **L.Error) bool {
	return BuilderValueFromString(b, pspec, str, value, err)
}
func (b *Builder) ValueFromStringType(t O.Type, str string, value *O.Value, err **L.Error) bool {
	return BuilderValueFromStringType(b, t, str, value, err)
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
	ButtonGetType         func() O.Type
	ButtonNew             func() *Widget
	ButtonNewFromStock    func(stockId string) *Widget
	ButtonNewWithLabel    func(label string) *Widget
	ButtonNewWithMnemonic func(label string) *Widget

	ButtonActionGetType   func() O.Type
	ButtonBoxGetType      func() O.Type
	ButtonBoxStyleGetType func() O.Type

	ButtonsTypeGetType func() O.Type
)

var (
	ButtonClicked          func(b *Button)
	ButtonEnter            func(b *Button)
	ButtonGetAlignment     func(b *Button, xalign, yalign *float32)
	ButtonGetEventWindow   func(b *Button) *D.Window
	ButtonGetFocusOnClick  func(b *Button) bool
	ButtonGetImage         func(b *Button) *Widget
	ButtonGetImagePosition func(b *Button) PositionType
	ButtonGetLabel         func(b *Button) string
	ButtonGetRelief        func(b *Button) ReliefStyle
	ButtonGetUseStock      func(b *Button) bool
	ButtonGetUseUnderline  func(b *Button) bool
	ButtonLeave            func(b *Button)
	ButtonPressed          func(b *Button)
	ButtonReleased         func(b *Button)
	ButtonSetAlignment     func(b *Button, xalign, yalign float32)
	ButtonSetFocusOnClick  func(b *Button, focusOnClick bool)
	ButtonSetImage         func(b *Button, image *Widget)
	ButtonSetImagePosition func(b *Button, position PositionType)
	ButtonSetLabel         func(b *Button, label string)
	ButtonSetRelief        func(b *Button, newstyle ReliefStyle)
	ButtonSetUseStock      func(b *Button, useStock bool)
	ButtonSetUseUnderline  func(b *Button, useUnderline bool)
)

func (b *Button) Clicked()                               { ButtonClicked(b) }
func (b *Button) Enter()                                 { ButtonEnter(b) }
func (b *Button) GetAlignment(xalign, yalign *float32)   { ButtonGetAlignment(b, xalign, yalign) }
func (b *Button) GetEventWindow() *D.Window              { return ButtonGetEventWindow(b) }
func (b *Button) GetFocusOnClick() bool                  { return ButtonGetFocusOnClick(b) }
func (b *Button) GetImage() *Widget                      { return ButtonGetImage(b) }
func (b *Button) GetImagePosition() PositionType         { return ButtonGetImagePosition(b) }
func (b *Button) GetLabel() string                       { return ButtonGetLabel(b) }
func (b *Button) GetRelief() ReliefStyle                 { return ButtonGetRelief(b) }
func (b *Button) GetUseStock() bool                      { return ButtonGetUseStock(b) }
func (b *Button) GetUseUnderline() bool                  { return ButtonGetUseUnderline(b) }
func (b *Button) Leave()                                 { ButtonLeave(b) }
func (b *Button) Pressed()                               { ButtonPressed(b) }
func (b *Button) Released()                              { ButtonReleased(b) }
func (b *Button) SetAlignment(xalign, yalign float32)    { ButtonSetAlignment(b, xalign, yalign) }
func (b *Button) SetFocusOnClick(focusOnClick bool)      { ButtonSetFocusOnClick(b, focusOnClick) }
func (b *Button) SetImage(image *Widget)                 { ButtonSetImage(b, image) }
func (b *Button) SetImagePosition(position PositionType) { ButtonSetImagePosition(b, position) }
func (b *Button) SetLabel(label string)                  { ButtonSetLabel(b, label) }
func (b *Button) SetRelief(newstyle ReliefStyle)         { ButtonSetRelief(b, newstyle) }
func (b *Button) SetUseStock(useStock bool)              { ButtonSetUseStock(b, useStock) }
func (b *Button) SetUseUnderline(useUnderline bool)      { ButtonSetUseUnderline(b, useUnderline) }

type ButtonBox struct {
	Box            Box
	ChildMinWidth  int
	ChildMinHeight int
	ChildIpadX     int
	ChildIpadY     int
	LayoutStyle    ButtonBoxStyle
}

type ButtonBoxStyle Enum

const (
	BUTTONBOX_DEFAULT_STYLE ButtonBoxStyle = iota
	BUTTONBOX_SPREAD
	BUTTONBOX_EDGE
	BUTTONBOX_START
	BUTTONBOX_END
	BUTTONBOX_CENTER
)

var (
	ButtonBoxGetChildIpadding  func(b *ButtonBox, ipadX, ipadY *int)
	ButtonBoxGetChildSecondary func(b *ButtonBox, child *Widget) bool
	ButtonBoxGetChildSize      func(b *ButtonBox, minWidth, minHeight *int)
	ButtonBoxGetLayout         func(b *ButtonBox) ButtonBoxStyle
	ButtonBoxSetChildIpadding  func(b *ButtonBox, ipadX, ipadY int)
	ButtonBoxSetChildSecondary func(b *ButtonBox, child *Widget, isSecondary bool)
	ButtonBoxSetChildSize      func(b *ButtonBox, minWidth, minHeight int)
	ButtonBoxSetLayout         func(b *ButtonBox, layoutStyle ButtonBoxStyle)
)

func (b *ButtonBox) GetChildIpadding(ipadX, ipadY *int) { ButtonBoxGetChildIpadding(b, ipadX, ipadY) }
func (b *ButtonBox) GetChildSecondary(child *Widget) bool {
	return ButtonBoxGetChildSecondary(b, child)
}
func (b *ButtonBox) GetChildSize(minWidth, minHeight *int) {
	ButtonBoxGetChildSize(b, minWidth, minHeight)
}
func (b *ButtonBox) GetLayout() ButtonBoxStyle         { return ButtonBoxGetLayout(b) }
func (b *ButtonBox) SetChildIpadding(ipadX, ipadY int) { ButtonBoxSetChildIpadding(b, ipadX, ipadY) }
func (b *ButtonBox) SetChildSecondary(child *Widget, isSecondary bool) {
	ButtonBoxSetChildSecondary(b, child, isSecondary)
}
func (b *ButtonBox) SetChildSize(minWidth, minHeight int) {
	ButtonBoxSetChildSize(b, minWidth, minHeight)
}
func (b *ButtonBox) SetLayout(layoutStyle ButtonBoxStyle) { ButtonBoxSetLayout(b, layoutStyle) }

type ButtonsType Enum

const (
	BUTTONS_NONE ButtonsType = iota
	BUTTONS_OK
	BUTTONS_CLOSE
	BUTTONS_CANCEL
	BUTTONS_YES_NO
	BUTTONS_OK_CANCEL
)
