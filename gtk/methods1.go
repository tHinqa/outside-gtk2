package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type AboutDialog struct {
	ParentInstance T.GtkDialog
	PrivateData    T.Gpointer
}

var (
	AboutDialogGetType func() T.GType
	AboutDialogNew     func() *T.GtkWidget

	AboutDialogSetEmailHook func(f T.GtkAboutDialogActivateLinkFunc, dataGpointer, destroy T.GDestroyNotify) T.GtkAboutDialogActivateLinkFunc
	AboutDialogSetUrlHook   func(f T.GtkAboutDialogActivateLinkFunc, dataGpointer, destroy T.GDestroyNotify) T.GtkAboutDialogActivateLinkFunc

	ShowAboutDialog func(parent *T.GtkWindow, firstPropertyName string, v ...VArg)
)

var (
	AboutDialogGetArtists           func(a *AboutDialog) **T.Gchar
	AboutDialogGetAuthors           func(a *AboutDialog) **T.Gchar
	AboutDialogGetComments          func(a *AboutDialog) string
	AboutDialogGetCopyright         func(a *AboutDialog) string
	AboutDialogGetDocumenters       func(a *AboutDialog) **T.Gchar
	AboutDialogGetLicense           func(a *AboutDialog) string
	AboutDialogGetLogo              func(a *AboutDialog) *T.GdkPixbuf
	AboutDialogGetLogoIconName      func(a *AboutDialog) string
	AboutDialogGetName              func(a *AboutDialog) string
	AboutDialogGetProgramName       func(a *AboutDialog) string
	AboutDialogGetTranslatorCredits func(a *AboutDialog) string
	AboutDialogGetVersion           func(a *AboutDialog) string
	AboutDialogGetWebsite           func(a *AboutDialog) string
	AboutDialogGetWebsiteLabel      func(a *AboutDialog) string
	AboutDialogGetWrapLicense       func(a *AboutDialog) T.Gboolean
	AboutDialogSetArtists           func(a *AboutDialog, artists **T.Gchar)
	AboutDialogSetAuthors           func(a *AboutDialog, authors **T.Gchar)
	AboutDialogSetComments          func(a *AboutDialog, comments string)
	AboutDialogSetCopyright         func(a *AboutDialog, copyright string)
	AboutDialogSetDocumenters       func(a *AboutDialog, documenters **T.Gchar)
	AboutDialogSetLicense           func(a *AboutDialog, license string)
	AboutDialogSetLogo              func(a *AboutDialog, logo *T.GdkPixbuf)
	AboutDialogSetLogoIconName      func(a *AboutDialog, iconName string)
	AboutDialogSetName              func(a *AboutDialog, name string)
	AboutDialogSetProgramName       func(a *AboutDialog, name string)
	AboutDialogSetTranslatorCredits func(a *AboutDialog, translatorCredits string)
	AboutDialogSetVersion           func(a *AboutDialog, version string)
	AboutDialogSetWebsite           func(a *AboutDialog, website string)
	AboutDialogSetWebsiteLabel      func(a *AboutDialog, websiteLabel string)
	AboutDialogSetWrapLicense       func(a *AboutDialog, wrapLicense T.Gboolean)
)

func (a *AboutDialog) GetName() string {
	return AboutDialogGetName(a)
}

func (a *AboutDialog) SetName(name string) {
	AboutDialogSetName(a, name)
}

func (a *AboutDialog) GetProgramName() string {
	return AboutDialogGetProgramName(a)
}

func (a *AboutDialog) SetProgramName(name string) {
	AboutDialogSetProgramName(a, name)
}

func (a *AboutDialog) GetVersion() string {
	return AboutDialogGetVersion(a)
}

func (a *AboutDialog) SetVersion(version string) {
	AboutDialogSetVersion(a, version)
}

func (a *AboutDialog) GetCopyright() string {
	return AboutDialogGetCopyright(a)
}

func (a *AboutDialog) SetCopyright(copyright string) {
	AboutDialogSetCopyright(a, copyright)
}

func (a *AboutDialog) GetComments() string {
	return AboutDialogGetComments(a)
}

func (a *AboutDialog) SetComments(comments string) {
	AboutDialogSetComments(a, comments)
}

func (a *AboutDialog) GetLicense() string {
	return AboutDialogGetLicense(a)
}

func (a *AboutDialog) SetLicense(license string) {
	AboutDialogSetLicense(a, license)
}

func (a *AboutDialog) GetWrapLicense() T.Gboolean {
	return AboutDialogGetWrapLicense(a)
}

func (a *AboutDialog) SetWrapLicense(wrapLicense T.Gboolean) {
	AboutDialogSetWrapLicense(a, wrapLicense)
}

func (a *AboutDialog) GetWebsite() string {
	return AboutDialogGetWebsite(a)
}

func (a *AboutDialog) SetWebsite(website string) {
	AboutDialogSetWebsite(a, website)
}

func (a *AboutDialog) GetWebsiteLabel() string {
	return AboutDialogGetWebsiteLabel(a)
}

func (a *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	AboutDialogSetWebsiteLabel(a, websiteLabel)
}

func (a *AboutDialog) GetAuthors() **T.Gchar {
	return AboutDialogGetAuthors(a)
}

func (a *AboutDialog) SetAuthors(authors **T.Gchar) {
	AboutDialogSetAuthors(a, authors)
}

func (a *AboutDialog) GetDocumenters() **T.Gchar {
	return AboutDialogGetDocumenters(a)
}

func (a *AboutDialog) SetDocumenters(documenters **T.Gchar) {
	AboutDialogSetDocumenters(a, documenters)
}

func (a *AboutDialog) GetArtists() **T.Gchar {
	return AboutDialogGetArtists(a)
}

func (a *AboutDialog) SetArtists(artists **T.Gchar) {
	AboutDialogSetArtists(a, artists)
}

func (a *AboutDialog) GetTranslatorCredits() string {
	return AboutDialogGetTranslatorCredits(a)
}

func (a *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	AboutDialogSetTranslatorCredits(a, translatorCredits)
}

func (a *AboutDialog) GetLogo() *T.GdkPixbuf {
	return AboutDialogGetLogo(a)
}

func (a *AboutDialog) SetLogo(logo *T.GdkPixbuf) {
	AboutDialogSetLogo(a, logo)
}

func (a *AboutDialog) GetLogoIconName() string {
	return AboutDialogGetLogoIconName(a)
}

func (a *AboutDialog) SetLogoIconName(iconName string) {
	AboutDialogSetLogoIconName(a, iconName)
}

type AccelGroup struct {
	Parent         T.GObject
	LockCount      uint
	ModifierMask   T.GdkModifierType
	Acceleratables *T.GSList
	NAccels        uint
	_              *T.GtkAccelGroupEntry
}

type AccelLabel struct {
	Label       T.GtkLabel
	_           uint
	Padding     uint //TODO(t):_?
	Widget      *T.GtkWidget
	Closure     *T.GClosure
	Group       *T.GtkAccelGroup
	String      *T.Gchar
	StringWidth uint16
}

var (
	AccelGroupGetType func() T.GType
	AccelGroupNew     func() *T.GtkAccelGroup

	AccelGroupFromAccelClosure func(closure *T.GClosure) *T.GtkAccelGroup

	AccelLabelGetType func() T.GType
	AccelLabelNew     func(str string) *T.GtkWidget

	AccelMapGetType func() T.GType

	AccelMapAddEntry          func(accelPath string, accelKey uint, accelMods T.GdkModifierType)
	AccelMapAddFilter         func(filterPattern string)
	AccelMapChangeEntry       func(accelPath string, accelKey uint, accelMods T.GdkModifierType, replace T.Gboolean) T.Gboolean
	AccelMapForeach           func(dataGpointer, foreachFunc T.GtkAccelMapForeach)
	AccelMapForeachUnfiltered func(dataGpointer, foreachFunc T.GtkAccelMapForeach)
	AccelMapGet               func() *T.GtkAccelMap
	AccelMapLoad              func(fileName string)
	AccelMapLoadFd            func(fd int)
	AccelMapLoadScanner       func(scanner *T.GScanner)
	AccelMapLockPath          func(accelPath string)
	AccelMapLookupEntry       func(accelPath string, key *T.GtkAccelKey) T.Gboolean
	AccelMapSave              func(fileName string)
	AccelMapSaveFd            func(fd int)
	AccelMapUnlockPath        func(accelPath string)
)

var (
	AccelGroupActivate        func(a *AccelGroup, accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) T.Gboolean
	AccelGroupConnect         func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags T.GtkAccelFlags, closure *T.GClosure)
	AccelGroupConnectByPath   func(a *AccelGroup, accelPath string, closure *T.GClosure)
	AccelGroupDisconnect      func(a *AccelGroup, closure *T.GClosure) T.Gboolean
	AccelGroupDisconnectKey   func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType) T.Gboolean
	AccelGroupFind            func(a *AccelGroup, findFunc T.GtkAccelGroupFindFunc, data T.Gpointer) *T.GtkAccelKey
	AccelGroupGetIsLocked     func(a *AccelGroup) T.Gboolean
	AccelGroupGetModifierMask func(a *AccelGroup) T.GdkModifierType
	AccelGroupLock            func(a *AccelGroup)
	AccelGroupQuery           func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *T.GtkAccelGroupEntry
	AccelGroupUnlock          func(a *AccelGroup)

	AccelLabelGetAccelWidget  func(a *AccelLabel) *T.GtkWidget
	AccelLabelGetAccelWidth   func(a *AccelLabel) uint
	AccelLabelRefetch         func(a *AccelLabel) T.Gboolean
	AccelLabelSetAccelClosure func(a *AccelLabel, accelClosure *T.GClosure)
	AccelLabelSetAccelWidget  func(a *AccelLabel, accelWidget *T.GtkWidget)
)

func (a *AccelGroup) GetIsLocked() T.Gboolean {
	return AccelGroupGetIsLocked(a)
}

func (a *AccelGroup) GetModifierMask() T.GdkModifierType {
	return AccelGroupGetModifierMask(a)
}

func (a *AccelGroup) Lock() { AccelGroupLock(a) }

func (a *AccelGroup) Unlock() { AccelGroupUnlock(a) }

func (a *AccelGroup) Connect(accelKey uint, accelMods T.GdkModifierType, accelFlags T.GtkAccelFlags, closure *T.GClosure) {
	AccelGroupConnect(a, accelKey, accelMods, accelFlags, closure)
}

func (a *AccelGroup) ConnectByPath(accelPath string, closure *T.GClosure) {
	AccelGroupConnectByPath(a, accelPath, closure)
}

func (a *AccelGroup) Disconnect(closure *T.GClosure) T.Gboolean {
	return AccelGroupDisconnect(a, closure)
}

func (a *AccelGroup) DisconnectKey(accelKey uint, accelMods T.GdkModifierType) T.Gboolean {
	return AccelGroupDisconnectKey(a, accelKey, accelMods)
}

func (a *AccelGroup) Activate(accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) T.Gboolean {
	return AccelGroupActivate(a, accelQuark, acceleratable, accelKey, accelMods)
}

func (a *AccelGroup) Find(findFunc T.GtkAccelGroupFindFunc, data T.Gpointer) *T.GtkAccelKey {
	return AccelGroupFind(a, findFunc, data)
}

func (a *AccelGroup) Query(accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *T.GtkAccelGroupEntry {
	return AccelGroupQuery(a, accelKey, accelMods, nEntries)
}

func (a *AccelLabel) GetAccelWidget() *T.GtkWidget {
	return AccelLabelGetAccelWidget(a)
}

func (a *AccelLabel) GetAccelWidth() uint {
	return AccelLabelGetAccelWidth(a)
}

func (a *AccelLabel) SetAccelWidget(accelWidget *T.GtkWidget) {
	AccelLabelSetAccelWidget(a, accelWidget)
}

func (a *AccelLabel) SetAccelClosure(accelClosure *T.GClosure) {
	AccelLabelSetAccelClosure(a, accelClosure)
}

func (a *AccelLabel) Refetch() T.Gboolean {
	return AccelLabelRefetch(a)
}

type Assistant struct {
	Parent  T.GtkWindow
	Cancel  *T.GtkWidget
	Forward *T.GtkWidget
	Back    *T.GtkWidget
	Apply   *T.GtkWidget
	Close   *T.GtkWidget
	Last    *T.GtkWidget
	Priv    *T.GtkAssistantPrivate
}

var (
	AssistantGetType func() T.GType
	AssistantNew     func() *T.GtkWidget

	AssistantPageTypeGetType func() T.GType
)

var (
	AssistantAddActionWidget    func(a *Assistant, child *T.GtkWidget)
	AssistantAppendPage         func(a *Assistant, page *T.GtkWidget) int
	AssistantCommit             func(a *Assistant)
	AssistantGetCurrentPage     func(a *Assistant) int
	AssistantGetNPages          func(a *Assistant) int
	AssistantGetNthPage         func(a *Assistant, pageNum int) *T.GtkWidget
	AssistantGetPageComplete    func(a *Assistant, page *T.GtkWidget) T.Gboolean
	AssistantGetPageHeaderImage func(a *Assistant, page *T.GtkWidget) *T.GdkPixbuf
	AssistantGetPageSideImage   func(a *Assistant, page *T.GtkWidget) *T.GdkPixbuf
	AssistantGetPageTitle       func(a *Assistant, page *T.GtkWidget) string
	AssistantGetPageType        func(a *Assistant, page *T.GtkWidget) T.GtkAssistantPageType
	AssistantInsertPage         func(a *Assistant, page *T.GtkWidget, position int) int
	AssistantPrependPage        func(a *Assistant, page *T.GtkWidget) int
	AssistantRemoveActionWidget func(a *Assistant, child *T.GtkWidget)
	AssistantSetCurrentPage     func(a *Assistant, pageNum int)
	AssistantSetForwardPageFunc func(a *Assistant, pageFunc T.GtkAssistantPageFunc, dataGpointer, destroy T.GDestroyNotify)
	AssistantSetPageComplete    func(a *Assistant, page *T.GtkWidget, complete T.Gboolean)
	AssistantSetPageHeaderImage func(a *Assistant, page *T.GtkWidget, pixbuf *T.GdkPixbuf)
	AssistantSetPageSideImage   func(a *Assistant, page *T.GtkWidget, pixbuf *T.GdkPixbuf)
	AssistantSetPageTitle       func(a *Assistant, page *T.GtkWidget, title string)
	AssistantSetPageType        func(a *Assistant, page *T.GtkWidget, t T.GtkAssistantPageType)
	AssistantUpdateButtonsState func(a *Assistant)
)

func (a *Assistant) GetCurrentPage() int {
	return AssistantGetCurrentPage(a)
}

func (a *Assistant) SetCurrentPage(pageNum int) {
	AssistantSetCurrentPage(a, pageNum)
}

func (a *Assistant) GetNPages() int {
	return AssistantGetNPages(a)
}

func (a *Assistant) GetNthPage(pageNum int) *T.GtkWidget {
	return AssistantGetNthPage(a, pageNum)
}

func (a *Assistant) PrependPage(page *T.GtkWidget) int {
	return AssistantPrependPage(a, page)
}

func (a *Assistant) AppendPage(page *T.GtkWidget) int {
	return AssistantAppendPage(a, page)
}

func (a *Assistant) InsertPage(page *T.GtkWidget, position int) int {
	return AssistantInsertPage(a, page, position)
}

func (a *Assistant) SetForwardPageFunc(pageFunc T.GtkAssistantPageFunc, dataGpointer, destroy T.GDestroyNotify) {
	AssistantSetForwardPageFunc(a, pageFunc, dataGpointer, destroy)
}

func (a *Assistant) SetPageType(page *T.GtkWidget, t T.GtkAssistantPageType) {
	AssistantSetPageType(a, page, t)
}

func (a *Assistant) GetPageType(page *T.GtkWidget) T.GtkAssistantPageType {
	return AssistantGetPageType(a, page)
}

func (a *Assistant) SetPageTitle(page *T.GtkWidget, title string) {
	AssistantSetPageTitle(a, page, title)
}

func (a *Assistant) GetPageTitle(page *T.GtkWidget) string {
	return AssistantGetPageTitle(a, page)
}

func (a *Assistant) SetPageHeaderImage(page *T.GtkWidget, pixbuf *T.GdkPixbuf) {
	AssistantSetPageHeaderImage(a, page, pixbuf)
}

func (a *Assistant) GetPageHeaderImage(page *T.GtkWidget) *T.GdkPixbuf {
	return AssistantGetPageHeaderImage(a, page)
}

func (a *Assistant) SetPageSideImage(page *T.GtkWidget, pixbuf *T.GdkPixbuf) {
	AssistantSetPageSideImage(a, page, pixbuf)
}

func (a *Assistant) GetPageSideImage(page *T.GtkWidget) *T.GdkPixbuf {
	return AssistantGetPageSideImage(a, page)
}

func (a *Assistant) SetPageComplete(page *T.GtkWidget, complete T.Gboolean) {
	AssistantSetPageComplete(a, page, complete)
}

func (a *Assistant) GetPageComplete(page *T.GtkWidget) T.Gboolean {
	return AssistantGetPageComplete(a, page)
}

func (a *Assistant) AddActionWidget(child *T.GtkWidget) {
	AssistantAddActionWidget(a, child)
}

func (a *Assistant) RemoveActionWidget(child *T.GtkWidget) {
	AssistantRemoveActionWidget(a, child)
}

func (a *Assistant) UpdateButtonsState() {
	AssistantUpdateButtonsState(a)
}

func (a *Assistant) Commit() { AssistantCommit(a) }

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
	Box            T.GtkBox
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

var (
	ComboBoxGetType              func() T.GType
	ComboBoxNew                  func() *T.GtkWidget
	ComboBoxNewWithEntry         func() *T.GtkWidget
	ComboBoxNewWithModel         func(model *T.GtkTreeModel) *T.GtkWidget
	ComboBoxNewWithModelAndEntry func(model *T.GtkTreeModel) *T.GtkWidget

	ComboBoxTextGetType      func() T.GType
	ComboBoxTextNew          func() *T.GtkWidget
	ComboBoxTextNewWithEntry func() *T.GtkWidget

	ComboBoxNewText func() *T.GtkWidget

	ComboBoxEntryGetType      func() T.GType
	ComboBoxEntryNew          func() *T.GtkWidget
	ComboBoxEntryNewWithModel func(model *T.GtkTreeModel, textColumn int) *T.GtkWidget
	ComboBoxEntryNewText      func() *T.GtkWidget
)

type ComboBox struct {
	ParentInstance T.GtkBin
	_              *T.GtkComboBoxPrivate
}

var (
	ComboBoxAppendText           func(c *ComboBox, text string)
	ComboBoxGetActive            func(c *ComboBox) int
	ComboBoxGetActiveIter        func(c *ComboBox, iter *T.GtkTreeIter) T.Gboolean
	ComboBoxGetActiveText        func(c *ComboBox) string
	ComboBoxGetAddTearoffs       func(c *ComboBox) T.Gboolean
	ComboBoxGetButtonSensitivity func(c *ComboBox) T.GtkSensitivityType
	ComboBoxGetColumnSpanColumn  func(c *ComboBox) int
	ComboBoxGetEntryTextColumn   func(c *ComboBox) int
	ComboBoxGetFocusOnClick      func(c *ComboBox) T.Gboolean
	ComboBoxGetHasEntry          func(c *ComboBox) T.Gboolean
	ComboBoxGetModel             func(c *ComboBox) *T.GtkTreeModel
	ComboBoxGetPopupAccessible   func(c *ComboBox) *T.AtkObject
	ComboBoxGetRowSeparatorFunc  func(c *ComboBox) T.GtkTreeViewRowSeparatorFunc
	ComboBoxGetRowSpanColumn     func(c *ComboBox) int
	ComboBoxGetTitle             func(c *ComboBox) string
	ComboBoxGetWrapWidth         func(c *ComboBox) int
	ComboBoxInsertText           func(c *ComboBox, position int, text string)
	ComboBoxPopdown              func(c *ComboBox)
	ComboBoxPopup                func(c *ComboBox)
	ComboBoxPrependText          func(c *ComboBox, text string)
	ComboBoxRemoveText           func(c *ComboBox, position int)
	ComboBoxSetActive            func(c *ComboBox, index int)
	ComboBoxSetActiveIter        func(c *ComboBox, iter *T.GtkTreeIter)
	ComboBoxSetAddTearoffs       func(c *ComboBox, addTearoffs T.Gboolean)
	ComboBoxSetButtonSensitivity func(c *ComboBox, sensitivity T.GtkSensitivityType)
	ComboBoxSetColumnSpanColumn  func(c *ComboBox, columnSpan int)
	ComboBoxSetEntryTextColumn   func(c *ComboBox, textColumn int)
	ComboBoxSetFocusOnClick      func(c *ComboBox, focusOnClick T.Gboolean)
	ComboBoxSetModel             func(c *ComboBox, model *T.GtkTreeModel)
	ComboBoxSetRowSeparatorFunc  func(c *ComboBox, f T.GtkTreeViewRowSeparatorFunc, dataGpointer, destroy T.GDestroyNotify)
	ComboBoxSetRowSpanColumn     func(c *ComboBox, rowSpan int)
	ComboBoxSetTitle             func(c *ComboBox, title string)
	ComboBoxSetWrapWidth         func(c *ComboBox, width int)
)

func (c *ComboBox) GetWrapWidth() int {
	return ComboBoxGetWrapWidth(c)
}

func (c *ComboBox) SetWrapWidth(width int) {
	ComboBoxSetWrapWidth(c, width)
}

func (c *ComboBox) GetRowSpanColumn() int {
	return ComboBoxGetRowSpanColumn(c)
}

func (c *ComboBox) SetRowSpanColumn(rowSpan int) {
	ComboBoxSetRowSpanColumn(c, rowSpan)
}

func (c *ComboBox) GetColumnSpanColumn() int {
	return ComboBoxGetColumnSpanColumn(c)
}

func (c *ComboBox) SetColumnSpanColumn(columnSpan int) {
	ComboBoxSetColumnSpanColumn(c, columnSpan)
}

func (c *ComboBox) GetAddTearoffs() T.Gboolean {
	return ComboBoxGetAddTearoffs(c)
}

func (c *ComboBox) SetAddTearoffs(addTearoffs T.Gboolean) {
	ComboBoxSetAddTearoffs(c, addTearoffs)
}

func (c *ComboBox) GetTitle() string {
	return ComboBoxGetTitle(c)
}

func (c *ComboBox) SetTitle(title string) {
	ComboBoxSetTitle(c, title)
}

func (c *ComboBox) GetFocusOnClick() T.Gboolean {
	return ComboBoxGetFocusOnClick(c)
}

func (c *ComboBox) SetFocusOnClick(focusOnClick T.Gboolean) {
	ComboBoxSetFocusOnClick(c, focusOnClick)
}

func (c *ComboBox) GetActive() int { return ComboBoxGetActive(c) }

func (c *ComboBox) SetActive(index int) {
	ComboBoxSetActive(c, index)
}

func (c *ComboBox) GetActiveIter(iter *T.GtkTreeIter) T.Gboolean {
	return ComboBoxGetActiveIter(c, iter)
}

func (c *ComboBox) SetActiveIter(iter *T.GtkTreeIter) {
	ComboBoxSetActiveIter(c, iter)
}

func (c *ComboBox) SetModel(model *T.GtkTreeModel) {
	ComboBoxSetModel(c, model)
}

func (c *ComboBox) GetModel() *T.GtkTreeModel {
	return ComboBoxGetModel(c)
}

func (c *ComboBox) GetRowSeparatorFunc() T.GtkTreeViewRowSeparatorFunc {
	return ComboBoxGetRowSeparatorFunc(c)
}

func (c *ComboBox) SetRowSeparatorFunc(f T.GtkTreeViewRowSeparatorFunc, dataGpointer, destroy T.GDestroyNotify) {
	ComboBoxSetRowSeparatorFunc(c, f, dataGpointer, destroy)
}

func (c *ComboBox) SetButtonSensitivity(sensitivity T.GtkSensitivityType) {
	ComboBoxSetButtonSensitivity(c, sensitivity)
}

func (c *ComboBox) GetButtonSensitivity() T.GtkSensitivityType {
	return ComboBoxGetButtonSensitivity(c)
}

func (c *ComboBox) GetHasEntry() T.Gboolean {
	return ComboBoxGetHasEntry(c)
}

func (c *ComboBox) SetEntryTextColumn(textColumn int) {
	ComboBoxSetEntryTextColumn(c, textColumn)
}

func (c *ComboBox) GetEntryTextColumn() int {
	return ComboBoxGetEntryTextColumn(c)
}

func (c *ComboBox) AppendText(text string) {
	ComboBoxAppendText(c, text)
}

func (c *ComboBox) InsertText(position int, text string) {
	ComboBoxInsertText(c, position, text)
}

func (c *ComboBox) PrependText(text string) {
	ComboBoxPrependText(c, text)
}

func (c *ComboBox) RemoveText(position int) {
	ComboBoxRemoveText(c, position)
}

func (c *ComboBox) GetActiveText() string {
	return ComboBoxGetActiveText(c)
}

func (c *ComboBox) Popup() { ComboBoxPopup(c) }

func (c *ComboBox) Popdown() { ComboBoxPopdown(c) }

func (c *ComboBox) GetPopupAccessible() *T.AtkObject {
	return ComboBoxGetPopupAccessible(c)
}

type ComboBoxText struct {
	ParentInstance *ComboBox
	_              *T.GtkComboBoxTextPrivate
}

var (
	ComboBoxTextAppendText    func(c *ComboBoxText, text string)
	ComboBoxTextGetActiveText func(c *ComboBoxText) string
	ComboBoxTextInsertText    func(c *ComboBoxText, position int, text string)
	ComboBoxTextPrependText   func(c *ComboBoxText, text string)
	ComboBoxTextRemove        func(c *ComboBoxText, position int)
)

func (c *ComboBoxText) AppendText(text string) {
	ComboBoxTextAppendText(c, text)
}

func (c *ComboBoxText) InsertText(position int, text string) {
	ComboBoxTextInsertText(c, position, text)
}

func (c *ComboBoxText) PrependText(text string) {
	ComboBoxTextPrependText(c, text)
}

func (c *ComboBoxText) Remove(position int) {
	ComboBoxTextRemove(c, position)
}

func (c *ComboBoxText) GetActiveText() string {
	return ComboBoxTextGetActiveText(c)
}

type ComboBoxEntry struct {
	ParentInstance ComboBox
	_              *T.GtkComboBoxEntryPrivate
}

var (
	ComboBoxEntrySetTextColumn func(e *ComboBoxEntry, textColumn int)
	ComboBoxEntryGetTextColumn func(e *ComboBoxEntry) int
)

func (e *ComboBoxEntry) SetTextColumn(textColumn int) {
	ComboBoxEntrySetTextColumn(e, textColumn)
}

func (e *ComboBoxEntry) GetTextColumn() int {
	return ComboBoxEntryGetTextColumn(e)
}
