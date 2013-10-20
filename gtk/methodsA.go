package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type (
	AboutDialog struct {
		Parent Dialog
		_      *struct{}
	}

	AboutDialogActivateLinkFunc func(
		about *AboutDialog, link string, data T.Gpointer)
)

var (
	AboutDialogGetType func() T.GType
	AboutDialogNew     func() *T.GtkWidget

	AboutDialogSetEmailHook func(f AboutDialogActivateLinkFunc, dataGpointer, destroy T.GDestroyNotify) AboutDialogActivateLinkFunc
	AboutDialogSetUrlHook   func(f AboutDialogActivateLinkFunc, dataGpointer, destroy T.GDestroyNotify) AboutDialogActivateLinkFunc

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

type (
	AccelMap struct{}

	AccelGroup struct {
		Parent         T.GObject
		LockCount      uint
		ModifierMask   T.GdkModifierType
		Acceleratables *T.GSList
		NAccels        uint
		_              *AccelGroupEntry
	}

	AccelGroupEntry struct {
		Key            AccelKey
		Closure        *T.GClosure
		AccelPathQuark T.GQuark
	}

	AccelKey struct { //
		AccelKey   uint
		AccelMods  T.GdkModifierType
		AccelFlags uint //: 16
	}

	AccelGroupFindFunc func(key *AccelKey,
		closure *T.GClosure, data T.Gpointer) T.Gboolean

	AccelMapForeachFunc func(
		data T.Gpointer,
		accelPath string,
		accelKey uint,
		accelMods T.GdkModifierType,
		changed T.Gboolean)
)

type AccelFlags T.Enum

const (
	ACCEL_VISIBLE AccelFlags = 1 << iota
	ACCEL_LOCKED
	ACCEL_MASK AccelFlags = 0x07
)

type AccelLabel struct {
	Label       Label
	_           uint
	Padding     uint //TODO(t):_?
	Widget      *T.GtkWidget
	Closure     *T.GClosure
	Group       *AccelGroup
	String      *T.Gchar
	StringWidth uint16
}

var (
	AccelGroupGetType func() T.GType
	AccelGroupNew     func() *AccelGroup

	AccelGroupFromAccelClosure func(closure *T.GClosure) *AccelGroup

	AccelLabelGetType func() T.GType
	AccelLabelNew     func(str string) *T.GtkWidget

	AccelMapGetType func() T.GType

	AccelMapAddEntry          func(accelPath string, accelKey uint, accelMods T.GdkModifierType)
	AccelMapAddFilter         func(filterPattern string)
	AccelMapChangeEntry       func(accelPath string, accelKey uint, accelMods T.GdkModifierType, replace T.Gboolean) T.Gboolean
	AccelMapForeach           func(dataGpointer, foreachFunc AccelMapForeachFunc)
	AccelMapForeachUnfiltered func(dataGpointer, foreachFunc AccelMapForeachFunc)
	AccelMapGet               func() *AccelMap
	AccelMapLoad              func(fileName string)
	AccelMapLoadFd            func(fd int)
	AccelMapLoadScanner       func(scanner *T.GScanner)
	AccelMapLockPath          func(accelPath string)
	AccelMapLookupEntry       func(accelPath string, key *AccelKey) T.Gboolean
	AccelMapSave              func(fileName string)
	AccelMapSaveFd            func(fd int)
	AccelMapUnlockPath        func(accelPath string)

	AccelFlagsGetType func() T.GType
)

var (
	AccelGroupActivate        func(a *AccelGroup, accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) T.Gboolean
	AccelGroupConnect         func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags, closure *T.GClosure)
	AccelGroupConnectByPath   func(a *AccelGroup, accelPath string, closure *T.GClosure)
	AccelGroupDisconnect      func(a *AccelGroup, closure *T.GClosure) T.Gboolean
	AccelGroupDisconnectKey   func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType) T.Gboolean
	AccelGroupFind            func(a *AccelGroup, findFunc AccelGroupFindFunc, data T.Gpointer) *AccelKey
	AccelGroupGetIsLocked     func(a *AccelGroup) T.Gboolean
	AccelGroupGetModifierMask func(a *AccelGroup) T.GdkModifierType
	AccelGroupLock            func(a *AccelGroup)
	AccelGroupQuery           func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *AccelGroupEntry
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

func (a *AccelGroup) Connect(accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags, closure *T.GClosure) {
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

func (a *AccelGroup) Find(findFunc AccelGroupFindFunc, data T.Gpointer) *AccelKey {
	return AccelGroupFind(a, findFunc, data)
}

func (a *AccelGroup) Query(accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *AccelGroupEntry {
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
	ActionNew     func(name, label, tooltip, stockId string) *Action

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
	ActionCreateIcon            func(a *Action, iconSize IconSize) *T.GtkWidget
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

func (a *Action) CreateIcon(iconSize IconSize) *T.GtkWidget {
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

type Activatable struct{}

var ActivatableGetType func() T.GType

var (
	ActivatableDoSetRelatedAction     func(a *Activatable, action *Action)
	ActivatableGetRelatedAction       func(a *Activatable) *Action
	ActivatableGetUseActionAppearance func(a *Activatable) T.Gboolean
	ActivatableSetRelatedAction       func(a *Activatable, action *Action)
	ActivatableSetUseActionAppearance func(a *Activatable, useAppearance T.Gboolean)
	ActivatableSyncActionProperties   func(a *Activatable, action *Action)
)

func (a *Activatable) SyncActionProperties(action *Action) {
	ActivatableSyncActionProperties(a, action)
}

func (a *Activatable) SetRelatedAction(action *Action) {
	ActivatableSetRelatedAction(a, action)
}

func (a *Activatable) GetRelatedAction() *Action {
	return ActivatableGetRelatedAction(a)
}

func (a *Activatable) SetUseActionAppearance(useAppearance T.Gboolean) {
	ActivatableSetUseActionAppearance(a, useAppearance)
}

func (a *Activatable) GetUseActionAppearance() T.Gboolean {
	return ActivatableGetUseActionAppearance(a)
}

func (a *Activatable) DoSetRelatedAction(action *Action) {
	ActivatableDoSetRelatedAction(a, action)
}

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

type (
	Assistant struct {
		Parent  T.GtkWindow
		Cancel  *T.GtkWidget
		Forward *T.GtkWidget
		Back    *T.GtkWidget
		Apply   *T.GtkWidget
		Close   *T.GtkWidget
		Last    *T.GtkWidget
		_       *struct{}
	}

	AssistantPageFunc func(currentPage int, data T.Gpointer) int
)

type AssistantPageType T.Enum

const (
	ASSISTANT_PAGE_CONTENT AssistantPageType = iota
	ASSISTANT_PAGE_INTRO
	ASSISTANT_PAGE_CONFIRM
	ASSISTANT_PAGE_SUMMARY
	ASSISTANT_PAGE_PROGRESS
)

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
	AssistantGetPageType        func(a *Assistant, page *T.GtkWidget) AssistantPageType
	AssistantInsertPage         func(a *Assistant, page *T.GtkWidget, position int) int
	AssistantPrependPage        func(a *Assistant, page *T.GtkWidget) int
	AssistantRemoveActionWidget func(a *Assistant, child *T.GtkWidget)
	AssistantSetCurrentPage     func(a *Assistant, pageNum int)
	AssistantSetForwardPageFunc func(a *Assistant, pageFunc AssistantPageFunc, dataGpointer, destroy T.GDestroyNotify)
	AssistantSetPageComplete    func(a *Assistant, page *T.GtkWidget, complete T.Gboolean)
	AssistantSetPageHeaderImage func(a *Assistant, page *T.GtkWidget, pixbuf *T.GdkPixbuf)
	AssistantSetPageSideImage   func(a *Assistant, page *T.GtkWidget, pixbuf *T.GdkPixbuf)
	AssistantSetPageTitle       func(a *Assistant, page *T.GtkWidget, title string)
	AssistantSetPageType        func(a *Assistant, page *T.GtkWidget, t AssistantPageType)
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

func (a *Assistant) SetForwardPageFunc(pageFunc AssistantPageFunc, dataGpointer, destroy T.GDestroyNotify) {
	AssistantSetForwardPageFunc(a, pageFunc, dataGpointer, destroy)
}

func (a *Assistant) SetPageType(page *T.GtkWidget, t AssistantPageType) {
	AssistantSetPageType(a, page, t)
}

func (a *Assistant) GetPageType(page *T.GtkWidget) AssistantPageType {
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

type Arrow struct {
	Misc       T.GtkMisc
	ArrowType  int16
	ShadowType int16
}

type ArrowType T.Enum

const (
	ARROW_UP ArrowType = iota
	ARROW_DOWN
	ARROW_LEFT
	ARROW_RIGHT
	ARROW_NONE
)

type ArrowPlacement T.Enum

const (
	ARROWS_BOTH ArrowPlacement = iota
	ARROWS_START
	ARROWS_END
)

var (
	ArrowGetType          func() T.GType
	ArrowTypeGetType      func() T.GType
	ArrowPlacementGetType func() T.GType
)

var (
	ArrowSet func(a *Arrow, arrowType ArrowType, shadowType T.GtkShadowType)

	ArrowNew func(a ArrowType, shadowType T.GtkShadowType) *T.GtkWidget
)

func (a *Arrow) Set(arrowType ArrowType, shadowType T.GtkShadowType) {
	ArrowSet(a, arrowType, shadowType)
}

func (a ArrowType) New(shadowType T.GtkShadowType) *T.GtkWidget {
	return ArrowNew(a, shadowType)
}

type AspectFrame struct {
	Frame            Frame
	Xalign           float32
	Yalign           float32
	Ratio            float32
	ObeyChild        T.Gboolean
	CenterAllocation T.GtkAllocation
}

var (
	AspectFrameGetType func() T.GType
	AspectFrameNew     func(label string, xalign, yalign, ratio float32, obeyChild T.Gboolean) *T.GtkWidget
)

var AspectFrameSet func(a *AspectFrame,
	xalign, yalign, ratio float32, obeyChild T.Gboolean)

func (a *AspectFrame) Set(
	xalign, yalign, ratio float32, obeyChild T.Gboolean) {
	AspectFrameSet(a, xalign, yalign, ratio, obeyChild)
}
