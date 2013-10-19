package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Bin struct {
	Container T.GtkContainer
	Child     *T.GtkWidget
}

var BinGetType func() T.GType

var BinGetChild func(b *Bin) *T.GtkWidget

func (b *Bin) GetChild() *T.GtkWidget {
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
	BorderCopy func(b *Border) *Border
	BorderFree func(b *Border)
)

func (b *Border) Copy() *Border { return BorderCopy(b) }

func (b *Border) Free() { BorderFree(b) }

type Box struct {
	Container   T.GtkContainer
	Children    *T.GList
	Spacing     int16
	Homogeneous uint // : 1
}

var BoxGetType func() T.GType

var (
	BoxGetHomogeneous    func(b *Box) T.Gboolean
	BoxGetSpacing        func(b *Box) int
	BoxPackEnd           func(b *Box, child *T.GtkWidget, expand T.Gboolean, fill T.Gboolean, padding uint)
	BoxPackEndDefaults   func(b *Box, widget *T.GtkWidget)
	BoxPackStart         func(b *Box, child *T.GtkWidget, expand T.Gboolean, fill T.Gboolean, padding uint)
	BoxPackStartDefaults func(b *Box, widget *T.GtkWidget)
	BoxQueryChildPacking func(b *Box, child *T.GtkWidget, expand *T.Gboolean, fill *T.Gboolean, padding *uint, packType *T.GtkPackType)
	BoxReorderChild      func(b *Box, child *T.GtkWidget, position int)
	BoxSetChildPacking   func(b *Box, child *T.GtkWidget, expand T.Gboolean, fill T.Gboolean, padding uint, packType T.GtkPackType)
	BoxSetHomogeneous    func(b *Box, homogeneous T.Gboolean)
	BoxSetSpacing        func(b *Box, spacing int)
)

func (b *Box) PackStart(child *T.GtkWidget,
	expand T.Gboolean, fill T.Gboolean, padding uint) {
	BoxPackStart(b, child, expand, fill, padding)
}

func (b *Box) PackEnd(child *T.GtkWidget,
	expand T.Gboolean, fill T.Gboolean, padding uint) {
	BoxPackEnd(b, child, expand, fill, padding)
}

func (b *Box) PackStartDefaults(widget *T.GtkWidget) {
	BoxPackStartDefaults(b, widget)
}

func (b *Box) PackEndDefaults(widget *T.GtkWidget) {
	BoxPackEndDefaults(b, widget)
}

func (b *Box) SetHomogeneous(homogeneous T.Gboolean) {
	BoxSetHomogeneous(b, homogeneous)
}

func (b *Box) GetHomogeneous() T.Gboolean {
	return BoxGetHomogeneous(b)
}

func (b *Box) SetSpacing(spacing int) { BoxSetSpacing(b, spacing) }

func (b *Box) GetSpacing() int { return BoxGetSpacing(b) }

func (b *Box) ReorderChild(child *T.GtkWidget, position int) {
	BoxReorderChild(b, child, position)
}

func (b *Box) QueryChildPacking(
	child *T.GtkWidget, expand *T.Gboolean, fill *T.Gboolean,
	padding *uint, packType *T.GtkPackType) {
	BoxQueryChildPacking(
		b, child, expand, fill, padding, packType)
}

func (b *Box) SetChildPacking(
	child *T.GtkWidget, expand T.Gboolean, fill T.Gboolean,
	padding uint, packType T.GtkPackType) {
	BoxSetChildPacking(b, child, expand, fill, padding, packType)
}

type Buildable struct{}

var BuildableGetType func() T.GType

var (
	BuildableAddChild             func(b *Buildable, builder *T.GtkBuilder, child *T.GObject, typ string)
	BuildableConstructChild       func(b *Buildable, builder *T.GtkBuilder, name string) *T.GObject
	BuildableCustomFinished       func(b *Buildable, builder *T.GtkBuilder, child *T.GObject, tagname string, data T.Gpointer)
	BuildableCustomTagEnd         func(b *Buildable, builder *T.GtkBuilder, child *T.GObject, tagname string, data *T.Gpointer)
	BuildableCustomTagStart       func(b *Buildable, builder *T.GtkBuilder, child *T.GObject, tagname string, parser *T.GMarkupParser, data *T.Gpointer) T.Gboolean
	BuildableGetInternalChild     func(b *Buildable, builder *T.GtkBuilder, childname string) *T.GObject
	BuildableGetName              func(b *Buildable) string
	BuildableParserFinished       func(b *Buildable, builder *T.GtkBuilder)
	BuildableSetBuildableProperty func(b *Buildable, builder *T.GtkBuilder, name string, value *T.GValue)
	BuildableSetName              func(b *Buildable, name string)
)

func (b *Buildable) SetName(name string) {
	BuildableSetName(b, name)
}

func (b *Buildable) GetName() string {
	return BuildableGetName(b)
}

func (b *Buildable) AddChild(
	builder *T.GtkBuilder, child *T.GObject, typ string) {
	BuildableAddChild(b, builder, child, typ)
}

func (b *Buildable) SetBuildableProperty(
	builder *T.GtkBuilder, name string, value *T.GValue) {
	BuildableSetBuildableProperty(b, builder, name, value)
}

func (b *Buildable) ConstructChild(
	builder *T.GtkBuilder, name string) *T.GObject {
	return BuildableConstructChild(b, builder, name)
}

func (b *Buildable) CustomTagStart(
	builder *T.GtkBuilder, child *T.GObject, tagname string,
	parser *T.GMarkupParser, data *T.Gpointer) T.Gboolean {
	return BuildableCustomTagStart(
		b, builder, child, tagname, parser, data)
}

func (b *Buildable) CustomTagEnd(builder *T.GtkBuilder,
	child *T.GObject, tagname string, data *T.Gpointer) {
	BuildableCustomTagEnd(b, builder, child, tagname, data)
}

func (b *Buildable) CustomFinished(builder *T.GtkBuilder,
	child *T.GObject, tagname string, data T.Gpointer) {
	BuildableCustomFinished(b, builder, child, tagname, data)
}

func (b *Buildable) ParserFinished(builder *T.GtkBuilder) {
	BuildableParserFinished(b, builder)
}

func (b *Buildable) GetInternalChild(
	builder *T.GtkBuilder, childname string) *T.GObject {
	return BuildableGetInternalChild(b, builder, childname)
}

type Builder struct {
	Parent T.GObject
	_      *struct{}
}

var (
	BuilderGetType func() T.GType
	BuilderNew     func() *T.GtkBuilder

	BuilderErrorGetType func() T.GType
	BuilderErrorQuark   func() T.GQuark
)

var (
	BuilderAddFromFile          func(b *Builder, filename string, err **T.GError) uint
	BuilderAddFromString        func(b *Builder, buffer string, length T.Gsize, err **T.GError) uint
	BuilderAddObjectsFromFile   func(b *Builder, filename string, objectIds **T.Gchar, err **T.GError) uint
	BuilderAddObjectsFromString func(b *Builder, buffer string, length T.Gsize, objectIds **T.Gchar, err **T.GError) uint
	BuilderConnectSignals       func(b *Builder, userData T.Gpointer)
	BuilderConnectSignalsFull   func(b *Builder, f T.GtkBuilderConnectFunc, userData T.Gpointer)
	BuilderGetObject            func(b *Builder, name string) *T.GObject
	BuilderGetObjects           func(b *Builder) *T.GSList
	BuilderGetTranslationDomain func(b *Builder) string
	BuilderGetTypeFromName      func(b *Builder, typeName string) T.GType
	BuilderSetTranslationDomain func(b *Builder, domain string)
	BuilderValueFromString      func(b *Builder, pspec *T.GParamSpec, str string, value *T.GValue, err **T.GError) T.Gboolean
	BuilderValueFromStringType  func(b *Builder, t T.GType, str string, value *T.GValue, err **T.GError) T.Gboolean
)

func (b *Builder) AddFromFile(
	filename string, err **T.GError) uint {
	return BuilderAddFromFile(b, filename, err)
}

func (b *Builder) AddFromString(
	buffer string, length T.Gsize, err **T.GError) uint {
	return BuilderAddFromString(b, buffer, length, err)
}

func (b *Builder) AddObjectsFromFile(
	filename string, objectIds **T.Gchar, err **T.GError) uint {
	return BuilderAddObjectsFromFile(b, filename, objectIds, err)
}

func (b *Builder) AddObjectsFromString(buffer string,
	length T.Gsize, objectIds **T.Gchar, err **T.GError) uint {
	return BuilderAddObjectsFromString(
		b, buffer, length, objectIds, err)
}

func (b *Builder) GetObject(name string) *T.GObject {
	return BuilderGetObject(b, name)
}

func (b *Builder) GetObjects() *T.GSList {
	return BuilderGetObjects(b)
}

func (b *Builder) ConnectSignals(userData T.Gpointer) {
	BuilderConnectSignals(b, userData)
}

func (b *Builder) ConnectSignalsFull(
	f T.GtkBuilderConnectFunc, userData T.Gpointer) {
	BuilderConnectSignalsFull(b, f, userData)
}

func (b *Builder) SetTranslationDomain(domain string) {
	BuilderSetTranslationDomain(b, domain)
}

func (b *Builder) GetTranslationDomain() string {
	return BuilderGetTranslationDomain(b)
}

func (b *Builder) GetTypeFromName(typeName string) T.GType {
	return BuilderGetTypeFromName(b, typeName)
}

func (b *Builder) ValueFromString(pspec *T.GParamSpec,
	str string, value *T.GValue, err **T.GError) T.Gboolean {
	return BuilderValueFromString(b, pspec, str, value, err)
}

func (b *Builder) ValueFromStringType(t T.GType, str string,
	value *T.GValue, err **T.GError) T.Gboolean {
	return BuilderValueFromStringType(b, t, str, value, err)
}

type Button struct {
	Bin             T.GtkBin
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
	ButtonNew             func() *T.GtkWidget
	ButtonNewFromStock    func(stockId string) *T.GtkWidget
	ButtonNewWithLabel    func(label string) *T.GtkWidget
	ButtonNewWithMnemonic func(label string) *T.GtkWidget

	ButtonActionGetType   func() T.GType
	ButtonBoxGetType      func() T.GType
	ButtonBoxStyleGetType func() T.GType

	ButtonsTypeGetType func() T.GType
)

var (
	ButtonClicked          func(b *Button)
	ButtonEnter            func(b *Button)
	ButtonGetAlignment     func(b *Button, xalign, yalign *float32)
	ButtonGetEventWindow   func(b *Button) *T.GdkWindow
	ButtonGetFocusOnClick  func(b *Button) T.Gboolean
	ButtonGetImage         func(b *Button) *T.GtkWidget
	ButtonGetImagePosition func(b *Button) T.GtkPositionType
	ButtonGetLabel         func(b *Button) string
	ButtonGetRelief        func(b *Button) T.GtkReliefStyle
	ButtonGetUseStock      func(b *Button) T.Gboolean
	ButtonGetUseUnderline  func(b *Button) T.Gboolean
	ButtonLeave            func(b *Button)
	ButtonPressed          func(b *Button)
	ButtonReleased         func(b *Button)
	ButtonSetAlignment     func(b *Button, xalign, yalign float32)
	ButtonSetFocusOnClick  func(b *Button, focusOnClick T.Gboolean)
	ButtonSetImage         func(b *Button, image *T.GtkWidget)
	ButtonSetImagePosition func(b *Button, position T.GtkPositionType)
	ButtonSetLabel         func(b *Button, label string)
	ButtonSetRelief        func(b *Button, newstyle T.GtkReliefStyle)
	ButtonSetUseStock      func(b *Button, useStock T.Gboolean)
	ButtonSetUseUnderline  func(b *Button, useUnderline T.Gboolean)
)

func (b *Button) Pressed() { ButtonPressed(b) }

func (b *Button) Released() { ButtonReleased(b) }

func (b *Button) Clicked() { ButtonClicked(b) }

func (b *Button) Enter() { ButtonEnter(b) }

func (b *Button) Leave() { ButtonLeave(b) }

func (b *Button) SetRelief(newstyle T.GtkReliefStyle) {
	ButtonSetRelief(b, newstyle)
}

func (b *Button) GetRelief() T.GtkReliefStyle {
	return ButtonGetRelief(b)
}

func (b *Button) SetLabel(label string) {
	ButtonSetLabel(b, label)
}

func (b *Button) GetLabel() string { return ButtonGetLabel(b) }

func (b *Button) SetUseUnderline(useUnderline T.Gboolean) {
	ButtonSetUseUnderline(b, useUnderline)
}

func (b *Button) GetUseUnderline() T.Gboolean {
	return ButtonGetUseUnderline(b)
}

func (b *Button) SetUseStock(useStock T.Gboolean) {
	ButtonSetUseStock(b, useStock)
}

func (b *Button) GetUseStock() T.Gboolean {
	return ButtonGetUseStock(b)
}

func (b *Button) SetFocusOnClick(focusOnClick T.Gboolean) {
	ButtonSetFocusOnClick(b, focusOnClick)
}

func (b *Button) GetFocusOnClick() T.Gboolean {
	return ButtonGetFocusOnClick(b)
}

func (b *Button) SetAlignment(xalign, yalign float32) {
	ButtonSetAlignment(b, xalign, yalign)
}

func (b *Button) GetAlignment(xalign, yalign *float32) {
	ButtonGetAlignment(b, xalign, yalign)
}

func (b *Button) SetImage(image *T.GtkWidget) {
	ButtonSetImage(b, image)
}

func (b *Button) GetImage() *T.GtkWidget {
	return ButtonGetImage(b)
}

func (b *Button) SetImagePosition(position T.GtkPositionType) {
	ButtonSetImagePosition(b, position)
}

func (b *Button) GetImagePosition() T.GtkPositionType {
	return ButtonGetImagePosition(b)
}

func (b *Button) GetEventWindow() *T.GdkWindow {
	return ButtonGetEventWindow(b)
}

type ButtonBox struct {
	Box            Box
	ChildMinWidth  int
	ChildMinHeight int
	ChildIpadX     int
	ChildIpadY     int
	LayoutStyle    T.GtkButtonBoxStyle
}

var (
	ButtonBoxGetChildIpadding  func(b *ButtonBox, ipadX, ipadY *int)
	ButtonBoxGetChildSecondary func(b *ButtonBox, child *T.GtkWidget) T.Gboolean
	ButtonBoxGetChildSize      func(b *ButtonBox, minWidth, minHeight *int)
	ButtonBoxGetLayout         func(b *ButtonBox) T.GtkButtonBoxStyle
	ButtonBoxSetChildIpadding  func(b *ButtonBox, ipadX, ipadY int)
	ButtonBoxSetChildSecondary func(b *ButtonBox, child *T.GtkWidget, isSecondary T.Gboolean)
	ButtonBoxSetChildSize      func(b *ButtonBox, minWidth, minHeight int)
	ButtonBoxSetLayout         func(b *ButtonBox, layoutStyle T.GtkButtonBoxStyle)
)

func (b *ButtonBox) GetLayout() T.GtkButtonBoxStyle {
	return ButtonBoxGetLayout(b)
}

func (b *ButtonBox) SetLayout(layoutStyle T.GtkButtonBoxStyle) {
	ButtonBoxSetLayout(b, layoutStyle)
}

func (b *ButtonBox) GetChildSecondary(child *T.GtkWidget) T.Gboolean {
	return ButtonBoxGetChildSecondary(b, child)
}

func (b *ButtonBox) SetChildSecondary(child *T.GtkWidget, isSecondary T.Gboolean) {
	ButtonBoxSetChildSecondary(b, child, isSecondary)
}

func (b *ButtonBox) SetChildSize(minWidth, minHeight int) {
	ButtonBoxSetChildSize(b, minWidth, minHeight)
}

func (b *ButtonBox) SetChildIpadding(ipadX, ipadY int) {
	ButtonBoxSetChildIpadding(b, ipadX, ipadY)
}

func (b *ButtonBox) GetChildSize(minWidth, minHeight *int) {
	ButtonBoxGetChildSize(b, minWidth, minHeight)
}

func (b *ButtonBox) GetChildIpadding(ipadX, ipadY *int) {
	ButtonBoxGetChildIpadding(b, ipadX, ipadY)
}
