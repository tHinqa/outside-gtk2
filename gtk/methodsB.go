// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Bin struct {
	Container T.GtkContainer
	Child     *Widget
}

var BinGetType func() T.GType

var BinGetChild func(b *Bin) *Widget

func (b *Bin) GetChild() *Widget {
	return BinGetChild(b)
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
	Container   T.GtkContainer
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
	boxQueryChildPacking func(b *Box, child *Widget, expand *T.Gboolean, fill *T.Gboolean, padding *uint, packType *T.GtkPackType)
	boxReorderChild      func(b *Box, child *Widget, position int)
	boxSetChildPacking   func(b *Box, child *Widget, expand T.Gboolean, fill T.Gboolean, padding uint, packType T.GtkPackType)
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
func (b *Box) QueryChildPacking(child *Widget, expand *T.Gboolean, fill *T.Gboolean, padding *uint, packType *T.GtkPackType) {
	boxQueryChildPacking(b, child, expand, fill, padding, packType)
}
func (b *Box) ReorderChild(child *Widget, position int) { boxReorderChild(b, child, position) }
func (b *Box) SetChildPacking(child *Widget, expand T.Gboolean, fill T.Gboolean, padding uint, packType T.GtkPackType) {
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
	EventWindow     *T.GdkWindow
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
	buttonGetEventWindow   func(b *Button) *T.GdkWindow
	buttonGetFocusOnClick  func(b *Button) T.Gboolean
	buttonGetImage         func(b *Button) *Widget
	buttonGetImagePosition func(b *Button) T.GtkPositionType
	buttonGetLabel         func(b *Button) string
	buttonGetRelief        func(b *Button) T.GtkReliefStyle
	buttonGetUseStock      func(b *Button) T.Gboolean
	buttonGetUseUnderline  func(b *Button) T.Gboolean
	buttonLeave            func(b *Button)
	buttonPressed          func(b *Button)
	buttonReleased         func(b *Button)
	buttonSetAlignment     func(b *Button, xalign, yalign float32)
	buttonSetFocusOnClick  func(b *Button, focusOnClick T.Gboolean)
	buttonSetImage         func(b *Button, image *Widget)
	buttonSetImagePosition func(b *Button, position T.GtkPositionType)
	buttonSetLabel         func(b *Button, label string)
	buttonSetRelief        func(b *Button, newstyle T.GtkReliefStyle)
	buttonSetUseStock      func(b *Button, useStock T.Gboolean)
	buttonSetUseUnderline  func(b *Button, useUnderline T.Gboolean)
)

func (b *Button) Clicked()                                    { buttonClicked(b) }
func (b *Button) Enter()                                      { buttonEnter(b) }
func (b *Button) GetAlignment(xalign, yalign *float32)        { buttonGetAlignment(b, xalign, yalign) }
func (b *Button) GetEventWindow() *T.GdkWindow                { return buttonGetEventWindow(b) }
func (b *Button) GetFocusOnClick() T.Gboolean                 { return buttonGetFocusOnClick(b) }
func (b *Button) GetImage() *Widget                           { return buttonGetImage(b) }
func (b *Button) GetImagePosition() T.GtkPositionType         { return buttonGetImagePosition(b) }
func (b *Button) GetLabel() string                            { return buttonGetLabel(b) }
func (b *Button) GetRelief() T.GtkReliefStyle                 { return buttonGetRelief(b) }
func (b *Button) GetUseStock() T.Gboolean                     { return buttonGetUseStock(b) }
func (b *Button) GetUseUnderline() T.Gboolean                 { return buttonGetUseUnderline(b) }
func (b *Button) Leave()                                      { buttonLeave(b) }
func (b *Button) Pressed()                                    { buttonPressed(b) }
func (b *Button) Released()                                   { buttonReleased(b) }
func (b *Button) SetAlignment(xalign, yalign float32)         { buttonSetAlignment(b, xalign, yalign) }
func (b *Button) SetFocusOnClick(focusOnClick T.Gboolean)     { buttonSetFocusOnClick(b, focusOnClick) }
func (b *Button) SetImage(image *Widget)                      { buttonSetImage(b, image) }
func (b *Button) SetImagePosition(position T.GtkPositionType) { buttonSetImagePosition(b, position) }
func (b *Button) SetLabel(label string)                       { buttonSetLabel(b, label) }
func (b *Button) SetRelief(newstyle T.GtkReliefStyle)         { buttonSetRelief(b, newstyle) }
func (b *Button) SetUseStock(useStock T.Gboolean)             { buttonSetUseStock(b, useStock) }
func (b *Button) SetUseUnderline(useUnderline T.Gboolean)     { buttonSetUseUnderline(b, useUnderline) }

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
