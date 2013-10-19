package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

//============================================================

type AboutDialog struct {
	Parent T.GtkDialog
	_      *struct{}
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

//============================================================

type AccelGroup struct {
	Parent         T.GObject
	LockCount      uint
	ModifierMask   T.GdkModifierType
	Acceleratables *T.GSList
	NAccels        uint
	_              *AccelGroupEntry
}

type AccelGroupEntry struct {
	Key            T.GtkAccelKey
	Closure        *T.GClosure
	AccelPathQuark T.GQuark
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

	AccelFlagsGetType func() T.GType
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

//============================================================

type Accessible struct {
	Parent T.AtkObject
	Widget *T.GtkWidget
}

var AccessibleGetType func() T.GType

var (
	AccessibleConnectWidgetDestroyed func(a *Accessible)
	AccessibleGetWidget              func(a *Accessible) *T.GtkWidget
	AccessibleSetWidget              func(a *Accessible, widget *T.GtkWidget)
)

func (a *Accessible) SetWidget(widget *T.GtkWidget) {
	AccessibleSetWidget(a, widget)
}

func (a *Accessible) GetWidget() *T.GtkWidget {
	return AccessibleGetWidget(a)
}

func (a *Accessible) ConnectWidgetDestroyed() {
	AccessibleConnectWidgetDestroyed(a)
}

//============================================================

type Action struct {
	Object T.GObject
	_      *struct{}
}

type ActionGroup struct {
	Parent T.GObject
	_      *struct{}
}

type ActionEntry struct {
	Name        *T.Gchar
	Stock_id    *T.Gchar
	Label       *T.Gchar
	Accelerator *T.Gchar
	Tooltip     *T.Gchar
	Callback    T.GCallback
}

var (
	ActionGetType func() T.GType
	ActionNew     func(name, label, tooltip, stockId string) *T.GtkAction

	ActionGroupGetType func() T.GType
	ActionGroupNew     func(name string) *ActionGroup

	WidgetGetAction func(widget *T.GtkWidget) *Action
)

var (
	ActionActivate              func(a *Action)
	ActionBlockActivate         func(a *Action)
	ActionBlockActivateFrom     func(a *Action, proxy *T.GtkWidget)
	ActionConnectAccelerator    func(a *Action)
	ActionConnectProxy          func(a *Action, proxy *T.GtkWidget)
	ActionCreateIcon            func(a *Action, iconSize T.GtkIconSize) *T.GtkWidget
	ActionCreateMenu            func(a *Action) *T.GtkWidget
	ActionCreateMenuItem        func(a *Action) *T.GtkWidget
	ActionCreateToolItem        func(a *Action) *T.GtkWidget
	ActionDisconnectAccelerator func(a *Action)
	ActionDisconnectProxy       func(a *Action, proxy *T.GtkWidget)
	ActionGetAccelClosure       func(a *Action) *T.GClosure
	ActionGetAccelPath          func(a *Action) string
	ActionGetAlwaysShowImage    func(a *Action) T.Gboolean
	ActionGetGicon              func(a *Action) *T.GIcon
	ActionGetIconName           func(a *Action) string
	ActionGetIsImportant        func(a *Action) T.Gboolean
	ActionGetLabel              func(a *Action) string
	ActionGetName               func(a *Action) string
	ActionGetProxies            func(a *Action) *T.GSList
	ActionGetSensitive          func(a *Action) T.Gboolean
	ActionGetShortLabel         func(a *Action) string
	ActionGetStockId            func(a *Action) string
	ActionGetTooltip            func(a *Action) string
	ActionGetVisible            func(a *Action) T.Gboolean
	ActionGetVisibleHorizontal  func(a *Action) T.Gboolean
	ActionGetVisibleVertical    func(a *Action) T.Gboolean
	ActionIsSensitive           func(a *Action) T.Gboolean
	ActionIsVisible             func(a *Action) T.Gboolean
	ActionSetAccelGroup         func(a *Action, accelGroup *AccelGroup)
	ActionSetAccelPath          func(a *Action, accelPath string)
	ActionSetAlwaysShowImage    func(a *Action, alwaysShow T.Gboolean)
	ActionSetGicon              func(a *Action, icon *T.GIcon)
	ActionSetIconName           func(a *Action, iconName string)
	ActionSetIsImportant        func(a *Action, isImportant T.Gboolean)
	ActionSetLabel              func(a *Action, label string)
	ActionSetSensitive          func(a *Action, sensitive T.Gboolean)
	ActionSetShortLabel         func(a *Action, shortLabel string)
	ActionSetStockId            func(a *Action, stockId string)
	ActionSetTooltip            func(a *Action, tooltip string)
	ActionSetVisible            func(a *Action, visible T.Gboolean)
	ActionSetVisibleHorizontal  func(a *Action, visibleHorizontal T.Gboolean)
	ActionSetVisibleVertical    func(a *Action, visibleVertical T.Gboolean)
	ActionUnblockActivate       func(a *Action)
	ActionUnblockActivateFrom   func(a *Action, proxy *T.GtkWidget)
)

func (a *Action) GetName() string {
	return ActionGetName(a)
}

func (a *Action) IsSensitive() T.Gboolean {
	return ActionIsSensitive(a)
}

func (a *Action) GetSensitive() T.Gboolean {
	return ActionGetSensitive(a)
}

func (a *Action) SetSensitive(sensitive T.Gboolean) {
	ActionSetSensitive(a, sensitive)
}

func (a *Action) IsVisible() T.Gboolean {
	return ActionIsVisible(a)
}

func (a *Action) GetVisible() T.Gboolean {
	return ActionGetVisible(a)
}

func (a *Action) SetVisible(visible T.Gboolean) {
	ActionSetVisible(a, visible)
}

func (a *Action) Activate() {
	ActionActivate(a)
}

func (a *Action) CreateIcon(iconSize T.GtkIconSize) *T.GtkWidget {
	return ActionCreateIcon(a, iconSize)
}

func (a *Action) CreateMenuItem() *T.GtkWidget {
	return ActionCreateMenuItem(a)
}

func (a *Action) CreateToolItem() *T.GtkWidget {
	return ActionCreateToolItem(a)
}

func (a *Action) CreateMenu() *T.GtkWidget {
	return ActionCreateMenu(a)
}

func (a *Action) GetProxies() *T.GSList {
	return ActionGetProxies(a)
}

func (a *Action) ConnectAccelerator() {
	ActionConnectAccelerator(a)
}

func (a *Action) DisconnectAccelerator() {
	ActionDisconnectAccelerator(a)
}

func (a *Action) GetAccelPath() string {
	return ActionGetAccelPath(a)
}

func (a *Action) GetAccelClosure() *T.GClosure {
	return ActionGetAccelClosure(a)
}

func (a *Action) ConnectProxy(proxy *T.GtkWidget) {
	ActionConnectProxy(a, proxy)
}

func (a *Action) DisconnectProxy(proxy *T.GtkWidget) {
	ActionDisconnectProxy(a, proxy)
}

func (a *Action) BlockActivateFrom(proxy *T.GtkWidget) {
	ActionBlockActivateFrom(a, proxy)
}

func (a *Action) UnblockActivateFrom(proxy *T.GtkWidget) {
	ActionUnblockActivateFrom(a, proxy)
}

func (a *Action) BlockActivate() { ActionBlockActivate(a) }

func (a *Action) UnblockActivate() { ActionUnblockActivate(a) }

func (a *Action) SetAccelPath(accelPath string) {
	ActionSetAccelPath(a, accelPath)
}

func (a *Action) SetAccelGroup(accelGroup *AccelGroup) {
	ActionSetAccelGroup(a, accelGroup)
}

func (a *Action) SetLabel(label string) {
	ActionSetLabel(a, label)
}

func (a *Action) GetLabel() string { return ActionGetLabel(a) }

func (a *Action) SetShortLabel(shortLabel string) {
	ActionSetShortLabel(a, shortLabel)
}

func (a *Action) GetShortLabel() string {
	return ActionGetShortLabel(a)
}

func (a *Action) SetTooltip(tooltip string) {
	ActionSetTooltip(a, tooltip)
}

func (a *Action) GetTooltip() string {
	return ActionGetTooltip(a)
}

func (a *Action) SetStockId(stockId string) {
	ActionSetStockId(a, stockId)
}

func (a *Action) GetStockId() string {
	return ActionGetStockId(a)
}

func (a *Action) SetGicon(icon *T.GIcon) {
	ActionSetGicon(a, icon)
}

func (a *Action) GetGicon() *T.GIcon { return ActionGetGicon(a) }

func (a *Action) SetIconName(iconName string) {
	ActionSetIconName(a, iconName)
}

func (a *Action) GetIconName() string {
	return ActionGetIconName(a)
}

func (a *Action) SetVisibleHorizontal(visibleHorizontal T.Gboolean) {
	ActionSetVisibleHorizontal(a, visibleHorizontal)
}

func (a *Action) GetVisibleHorizontal() T.Gboolean {
	return ActionGetVisibleHorizontal(a)
}

func (a *Action) SetVisibleVertical(visibleVertical T.Gboolean) {
	ActionSetVisibleVertical(a, visibleVertical)
}

func (a *Action) GetVisibleVertical() T.Gboolean {
	return ActionGetVisibleVertical(a)
}

func (a *Action) SetIsImportant(isImportant T.Gboolean) {
	ActionSetIsImportant(a, isImportant)
}

func (a *Action) GetIsImportant() T.Gboolean {
	return ActionGetIsImportant(a)
}

func (a *Action) SetAlwaysShowImage(alwaysShow T.Gboolean) {
	ActionSetAlwaysShowImage(a, alwaysShow)
}

func (a *Action) GetAlwaysShowImage() T.Gboolean {
	return ActionGetAlwaysShowImage(a)
}

var (
	ActionGroupAddAction            func(a *ActionGroup, action *Action)
	ActionGroupAddActions           func(a *ActionGroup, entries *ActionEntry, nEntries uint, userData T.Gpointer)
	ActionGroupAddActionsFull       func(a *ActionGroup, entries *ActionEntry, nEntries uint, userDataGpointer, destroy T.GDestroyNotify)
	ActionGroupAddActionWithAccel   func(a *ActionGroup, action *Action, accelerator string)
	ActionGroupAddRadioActions      func(a *ActionGroup, entries *T.GtkRadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer)
	ActionGroupAddRadioActionsFull  func(a *ActionGroup, entries *T.GtkRadioActionEntry, nEntries uint, value int, onChange T.GCallback, userDataGpointer, destroy T.GDestroyNotify)
	ActionGroupAddToggleActions     func(a *ActionGroup, entries *T.GtkToggleActionEntry, nEntries uint, userData T.Gpointer)
	ActionGroupAddToggleActionsFull func(a *ActionGroup, entries *T.GtkToggleActionEntry, nEntries uint, userDataGpointer, destroy T.GDestroyNotify)
	ActionGroupGetAction            func(a *ActionGroup, actionName string) *Action
	ActionGroupGetName              func(a *ActionGroup) string
	ActionGroupGetSensitive         func(a *ActionGroup) T.Gboolean
	ActionGroupGetVisible           func(a *ActionGroup) T.Gboolean
	ActionGroupListActions          func(a *ActionGroup) *T.GList
	ActionGroupRemoveAction         func(a *ActionGroup, action *Action)
	ActionGroupSetSensitive         func(a *ActionGroup, sensitive T.Gboolean)
	ActionGroupSetTranslateFunc     func(a *ActionGroup, f T.GtkTranslateFunc, dataGpointer, notify T.GDestroyNotify)
	ActionGroupSetTranslationDomain func(a *ActionGroup, domain string)
	ActionGroupSetVisible           func(a *ActionGroup, visible T.Gboolean)
	ActionGroupTranslateString      func(a *ActionGroup, str string) string
)

func (a *ActionGroup) GetName() string {
	return ActionGroupGetName(a)
}

func (a *ActionGroup) GetSensitive() T.Gboolean {
	return ActionGroupGetSensitive(a)
}

func (a *ActionGroup) SetSensitive(sensitive T.Gboolean) {
	ActionGroupSetSensitive(a, sensitive)
}

func (a *ActionGroup) GetVisible() T.Gboolean {
	return ActionGroupGetVisible(a)
}

func (a *ActionGroup) SetVisible(visible T.Gboolean) {
	ActionGroupSetVisible(a, visible)
}

func (a *ActionGroup) GetAction(actionName string) *Action {
	return ActionGroupGetAction(a, actionName)
}

func (a *ActionGroup) ListActions() *T.GList {
	return ActionGroupListActions(a)
}

func (a *ActionGroup) AddAction(action *Action) {
	ActionGroupAddAction(a, action)
}

func (a *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	ActionGroupAddActionWithAccel(a, action, accelerator)
}

func (a *ActionGroup) RemoveAction(action *Action) {
	ActionGroupRemoveAction(a, action)
}

func (a *ActionGroup) AddActions(entries *ActionEntry, nEntries uint, userData T.Gpointer) {
	ActionGroupAddActions(a, entries, nEntries, userData)
}

func (a *ActionGroup) AddToggleActions(entries *T.GtkToggleActionEntry, nEntries uint, userData T.Gpointer) {
	ActionGroupAddToggleActions(a, entries, nEntries, userData)
}

func (a *ActionGroup) AddRadioActions(entries *T.GtkRadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer) {
	ActionGroupAddRadioActions(a, entries, nEntries, value, onChange, userData)
}

func (a *ActionGroup) AddActionsFull(entries *ActionEntry,
	nEntries uint, userDataGpointer, destroy T.GDestroyNotify) {
	ActionGroupAddActionsFull(a, entries, nEntries, userDataGpointer, destroy)
}

func (a *ActionGroup) AddToggleActionsFull(
	entries *T.GtkToggleActionEntry, nEntries uint,
	userDataGpointer, destroy T.GDestroyNotify) {
	ActionGroupAddToggleActionsFull(a, entries, nEntries, userDataGpointer, destroy)
}

func (a *ActionGroup) AddRadioActionsFull(
	entries *T.GtkRadioActionEntry, nEntries uint,
	value int, onChange T.GCallback,
	userDataGpointer, destroy T.GDestroyNotify) {
	ActionGroupAddRadioActionsFull(a, entries, nEntries, value, onChange, userDataGpointer, destroy)
}

func (a *ActionGroup) SetTranslateFunc(
	f T.GtkTranslateFunc, dataGpointer, notify T.GDestroyNotify) {
	ActionGroupSetTranslateFunc(a, f, dataGpointer, notify)
}

func (a *ActionGroup) SetTranslationDomain(domain string) {
	ActionGroupSetTranslationDomain(a, domain)
}

func (a *ActionGroup) TranslateString(str string) string {
	return ActionGroupTranslateString(a, str)
}

//============================================================

type Activatable struct{}

var ActivatableGetType func() T.GType

var (
	ActivatableDoSetRelatedAction     func(a *Activatable, action *T.GtkAction)
	ActivatableGetRelatedAction       func(a *Activatable) *T.GtkAction
	ActivatableGetUseActionAppearance func(a *Activatable) T.Gboolean
	ActivatableSetRelatedAction       func(a *Activatable, action *T.GtkAction)
	ActivatableSetUseActionAppearance func(a *Activatable, useAppearance T.Gboolean)
	ActivatableSyncActionProperties   func(a *Activatable, action *T.GtkAction)
)

func (a *Activatable) SyncActionProperties(action *T.GtkAction) {
	ActivatableSyncActionProperties(a, action)
}

func (a *Activatable) SetRelatedAction(action *T.GtkAction) {
	ActivatableSetRelatedAction(a, action)
}

func (a *Activatable) GetRelatedAction() *T.GtkAction {
	return ActivatableGetRelatedAction(a)
}

func (a *Activatable) SetUseActionAppearance(useAppearance T.Gboolean) {
	ActivatableSetUseActionAppearance(a, useAppearance)
}

func (a *Activatable) GetUseActionAppearance() T.Gboolean {
	return ActivatableGetUseActionAppearance(a)
}

func (a *Activatable) DoSetRelatedAction(action *T.GtkAction) {
	ActivatableDoSetRelatedAction(a, action)
}

//============================================================

type Adjustment struct {
	Parent        T.GtkObject
	Lower         float64
	Upper         float64
	Value         float64
	StepIncrement float64
	PageIncrement float64
	PageSize      float64
}

var (
	AdjustmentGetType func() T.GType
	AdjustmentNew     func(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *T.GtkObject
)

var (
	AdjustmentChanged          func(a *Adjustment)
	AdjustmentClampPage        func(a *Adjustment, lower, upper float64)
	AdjustmentConfigure        func(a *Adjustment, value, lower, upper, stepIncrement, pageIncrement, pageSize float64)
	AdjustmentGetLower         func(a *Adjustment) float64
	AdjustmentGetPageIncrement func(a *Adjustment) float64
	AdjustmentGetPageSize      func(a *Adjustment) float64
	AdjustmentGetStepIncrement func(a *Adjustment) float64
	AdjustmentGetUpper         func(a *Adjustment) float64
	AdjustmentGetValue         func(a *Adjustment) float64
	AdjustmentSetLower         func(a *Adjustment, lower float64)
	AdjustmentSetPageIncrement func(a *Adjustment, pageIncrement float64)
	AdjustmentSetPageSize      func(a *Adjustment, pageSize float64)
	AdjustmentSetStepIncrement func(a *Adjustment, stepIncrement float64)
	AdjustmentSetUpper         func(a *Adjustment, upper float64)
	AdjustmentSetValue         func(a *Adjustment, value float64)
	AdjustmentValueChanged     func(a *Adjustment)
)

func (a *Adjustment) Changed() { AdjustmentChanged(a) }

func (a *Adjustment) ValueChanged() { AdjustmentValueChanged(a) }

func (a *Adjustment) ClampPage(lower, upper float64) {
	AdjustmentClampPage(a, lower, upper)
}

func (a *Adjustment) GetValue() float64 {
	return AdjustmentGetValue(a)
}

func (a *Adjustment) SetValue(value float64) {
	AdjustmentSetValue(a, value)
}

func (a *Adjustment) GetLower() float64 {
	return AdjustmentGetLower(a)
}

func (a *Adjustment) SetLower(lower float64) {
	AdjustmentSetLower(a, lower)
}

func (a *Adjustment) GetUpper() float64 {
	return AdjustmentGetUpper(a)
}

func (a *Adjustment) SetUpper(upper float64) {
	AdjustmentSetUpper(a, upper)
}

func (a *Adjustment) GetStepIncrement() float64 {
	return AdjustmentGetStepIncrement(a)
}

func (a *Adjustment) SetStepIncrement(stepIncrement float64) {
	AdjustmentSetStepIncrement(a, stepIncrement)
}

func (a *Adjustment) GetPageIncrement() float64 {
	return AdjustmentGetPageIncrement(a)
}

func (a *Adjustment) SetPageIncrement(pageIncrement float64) {
	AdjustmentSetPageIncrement(a, pageIncrement)
}

func (a *Adjustment) GetPageSize() float64 {
	return AdjustmentGetPageSize(a)
}

func (a *Adjustment) SetPageSize(pageSize float64) {
	AdjustmentSetPageSize(a, pageSize)
}

func (a *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	AdjustmentConfigure(a, value, lower, upper, stepIncrement, pageIncrement, pageSize)
}

//============================================================

type Assistant struct {
	Parent  T.GtkWindow
	Cancel  *T.GtkWidget
	Forward *T.GtkWidget
	Back    *T.GtkWidget
	Apply   *T.GtkWidget
	Close   *T.GtkWidget
	Last    *T.GtkWidget
	_       *struct{}
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

//============================================================

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

//============================================================

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

//============================================================

type ComboBox struct {
	ParentInstance T.GtkBin
	_              *T.GtkComboBoxPrivate
}

type ComboBoxText struct {
	ParentInstance *ComboBox
	_              *T.GtkComboBoxTextPrivate
}

type ComboBoxEntry struct {
	ParentInstance ComboBox
	_              *T.GtkComboBoxEntryPrivate
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
