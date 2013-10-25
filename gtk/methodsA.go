// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	I "github.com/tHinqa/outside-gtk2/gio"
	O "github.com/tHinqa/outside-gtk2/gobject"
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
	AboutDialogGetType func() O.Type
	AboutDialogNew     func() *Widget

	AboutDialogSetEmailHook func(f AboutDialogActivateLinkFunc, data T.Gpointer, destroy T.GDestroyNotify) AboutDialogActivateLinkFunc
	AboutDialogSetUrlHook   func(f AboutDialogActivateLinkFunc, data T.Gpointer, destroy T.GDestroyNotify) AboutDialogActivateLinkFunc

	ShowAboutDialog func(parent *Window, firstPropertyName string, v ...VArg)
)

var (
	aboutDialogGetArtists           func(a *AboutDialog) **T.Gchar
	aboutDialogGetAuthors           func(a *AboutDialog) **T.Gchar
	aboutDialogGetComments          func(a *AboutDialog) string
	aboutDialogGetCopyright         func(a *AboutDialog) string
	aboutDialogGetDocumenters       func(a *AboutDialog) **T.Gchar
	aboutDialogGetLicense           func(a *AboutDialog) string
	aboutDialogGetLogo              func(a *AboutDialog) *D.Pixbuf
	aboutDialogGetLogoIconName      func(a *AboutDialog) string
	aboutDialogGetName              func(a *AboutDialog) string
	aboutDialogGetProgramName       func(a *AboutDialog) string
	aboutDialogGetTranslatorCredits func(a *AboutDialog) string
	aboutDialogGetVersion           func(a *AboutDialog) string
	aboutDialogGetWebsite           func(a *AboutDialog) string
	aboutDialogGetWebsiteLabel      func(a *AboutDialog) string
	aboutDialogGetWrapLicense       func(a *AboutDialog) T.Gboolean
	aboutDialogSetArtists           func(a *AboutDialog, artists **T.Gchar)
	aboutDialogSetAuthors           func(a *AboutDialog, authors **T.Gchar)
	aboutDialogSetComments          func(a *AboutDialog, comments string)
	aboutDialogSetCopyright         func(a *AboutDialog, copyright string)
	aboutDialogSetDocumenters       func(a *AboutDialog, documenters **T.Gchar)
	aboutDialogSetLicense           func(a *AboutDialog, license string)
	aboutDialogSetLogo              func(a *AboutDialog, logo *D.Pixbuf)
	aboutDialogSetLogoIconName      func(a *AboutDialog, iconName string)
	aboutDialogSetName              func(a *AboutDialog, name string)
	aboutDialogSetProgramName       func(a *AboutDialog, name string)
	aboutDialogSetTranslatorCredits func(a *AboutDialog, translatorCredits string)
	aboutDialogSetVersion           func(a *AboutDialog, version string)
	aboutDialogSetWebsite           func(a *AboutDialog, website string)
	aboutDialogSetWebsiteLabel      func(a *AboutDialog, websiteLabel string)
	aboutDialogSetWrapLicense       func(a *AboutDialog, wrapLicense T.Gboolean)
)

func (a *AboutDialog) GetArtists() **T.Gchar                { return aboutDialogGetArtists(a) }
func (a *AboutDialog) GetAuthors() **T.Gchar                { return aboutDialogGetAuthors(a) }
func (a *AboutDialog) GetComments() string                  { return aboutDialogGetComments(a) }
func (a *AboutDialog) GetCopyright() string                 { return aboutDialogGetCopyright(a) }
func (a *AboutDialog) GetDocumenters() **T.Gchar            { return aboutDialogGetDocumenters(a) }
func (a *AboutDialog) GetLicense() string                   { return aboutDialogGetLicense(a) }
func (a *AboutDialog) GetLogo() *D.Pixbuf                   { return aboutDialogGetLogo(a) }
func (a *AboutDialog) GetLogoIconName() string              { return aboutDialogGetLogoIconName(a) }
func (a *AboutDialog) GetName() string                      { return aboutDialogGetName(a) }
func (a *AboutDialog) GetProgramName() string               { return aboutDialogGetProgramName(a) }
func (a *AboutDialog) GetTranslatorCredits() string         { return aboutDialogGetTranslatorCredits(a) }
func (a *AboutDialog) GetVersion() string                   { return aboutDialogGetVersion(a) }
func (a *AboutDialog) GetWebsite() string                   { return aboutDialogGetWebsite(a) }
func (a *AboutDialog) GetWebsiteLabel() string              { return aboutDialogGetWebsiteLabel(a) }
func (a *AboutDialog) GetWrapLicense() T.Gboolean           { return aboutDialogGetWrapLicense(a) }
func (a *AboutDialog) SetArtists(artists **T.Gchar)         { aboutDialogSetArtists(a, artists) }
func (a *AboutDialog) SetAuthors(authors **T.Gchar)         { aboutDialogSetAuthors(a, authors) }
func (a *AboutDialog) SetComments(comments string)          { aboutDialogSetComments(a, comments) }
func (a *AboutDialog) SetCopyright(copyright string)        { aboutDialogSetCopyright(a, copyright) }
func (a *AboutDialog) SetDocumenters(documenters **T.Gchar) { aboutDialogSetDocumenters(a, documenters) }
func (a *AboutDialog) SetLicense(license string)            { aboutDialogSetLicense(a, license) }
func (a *AboutDialog) SetLogo(logo *D.Pixbuf)               { aboutDialogSetLogo(a, logo) }
func (a *AboutDialog) SetLogoIconName(iconName string)      { aboutDialogSetLogoIconName(a, iconName) }
func (a *AboutDialog) SetName(name string)                  { aboutDialogSetName(a, name) }
func (a *AboutDialog) SetProgramName(name string)           { aboutDialogSetProgramName(a, name) }
func (a *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	aboutDialogSetTranslatorCredits(a, translatorCredits)
}
func (a *AboutDialog) SetVersion(version string) { aboutDialogSetVersion(a, version) }
func (a *AboutDialog) SetWebsite(website string) { aboutDialogSetWebsite(a, website) }
func (a *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	aboutDialogSetWebsiteLabel(a, websiteLabel)
}
func (a *AboutDialog) SetWrapLicense(wrapLicense T.Gboolean) {
	aboutDialogSetWrapLicense(a, wrapLicense)
}

var (
	AcceleratorGetDefaultModMask func() uint

	AcceleratorGetLabel func(
		acceleratorKey uint,
		acceleratorMods T.GdkModifierType) string

	AcceleratorName func(
		acceleratorKey uint,
		acceleratorMods T.GdkModifierType) string

	AcceleratorParse func(
		accelerator string,
		acceleratorKey *uint,
		acceleratorMods *T.GdkModifierType)

	AcceleratorSetDefaultModMask func(
		defaultModMask T.GdkModifierType)

	AcceleratorValid func(
		keyval uint,
		modifiers T.GdkModifierType) T.Gboolean
)

var (
	AccelGroupsActivate func(object *T.GObject, accelKey uint,
		accelMods T.GdkModifierType) T.Gboolean

	AccelGroupsFromObject func(object *T.GObject) *T.GSList
)

type AccelLabelClass struct {
	ParentClass    LabelClass
	SignalQuote1   *T.Gchar
	SignalQuote2   *T.Gchar
	ModNameShift   *T.Gchar
	ModNameControl *T.Gchar
	ModNameAlt     *T.Gchar
	ModSeparator   *T.Gchar
	AccelSeperator *T.Gchar
	Latin1ToChar   uint // : 1
	_, _, _, _     func()
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
		Closure        *O.Closure
		AccelPathQuark T.GQuark
	}

	AccelKey struct { //
		AccelKey   uint
		AccelMods  T.GdkModifierType
		AccelFlags uint //: 16
	}

	AccelGroupFindFunc func(key *AccelKey,
		closure *O.Closure, data T.Gpointer) T.Gboolean

	AccelMapForeachFunc func(
		data T.Gpointer,
		accelPath string,
		accelKey uint,
		accelMods T.GdkModifierType,
		changed T.Gboolean)
)

type AccelFlags Enum

const (
	ACCEL_VISIBLE AccelFlags = 1 << iota
	ACCEL_LOCKED
	ACCEL_MASK AccelFlags = 0x07
)

type AccelLabel struct {
	Label       Label
	_           uint
	Padding     uint //TODO(t):_?
	Widget      *Widget
	Closure     *O.Closure
	Group       *AccelGroup
	String      *T.Gchar
	StringWidth uint16
}

var (
	AccelGroupGetType func() O.Type
	AccelGroupNew     func() *AccelGroup

	AccelGroupFromAccelClosure func(closure *O.Closure) *AccelGroup

	AccelLabelGetType func() O.Type
	AccelLabelNew     func(str string) *Widget

	AccelMapGetType func() O.Type

	AccelMapAddEntry          func(accelPath string, accelKey uint, accelMods T.GdkModifierType)
	AccelMapAddFilter         func(filterPattern string)
	AccelMapChangeEntry       func(accelPath string, accelKey uint, accelMods T.GdkModifierType, replace T.Gboolean) T.Gboolean
	AccelMapForeach           func(data T.Gpointer, foreachFunc AccelMapForeachFunc)
	AccelMapForeachUnfiltered func(data T.Gpointer, foreachFunc AccelMapForeachFunc)
	AccelMapGet               func() *AccelMap
	AccelMapLoad              func(fileName string)
	AccelMapLoadFd            func(fd int)
	AccelMapLoadScanner       func(scanner *T.GScanner)
	AccelMapLockPath          func(accelPath string)
	AccelMapLookupEntry       func(accelPath string, key *AccelKey) T.Gboolean
	AccelMapSave              func(fileName string)
	AccelMapSaveFd            func(fd int)
	AccelMapUnlockPath        func(accelPath string)

	AccelFlagsGetType func() O.Type
)

var (
	accelGroupActivate        func(a *AccelGroup, accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) T.Gboolean
	accelGroupConnect         func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags, closure *O.Closure)
	accelGroupConnectByPath   func(a *AccelGroup, accelPath string, closure *O.Closure)
	accelGroupDisconnect      func(a *AccelGroup, closure *O.Closure) T.Gboolean
	accelGroupDisconnectKey   func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType) T.Gboolean
	accelGroupFind            func(a *AccelGroup, findFunc AccelGroupFindFunc, data T.Gpointer) *AccelKey
	accelGroupGetIsLocked     func(a *AccelGroup) T.Gboolean
	accelGroupGetModifierMask func(a *AccelGroup) T.GdkModifierType
	accelGroupLock            func(a *AccelGroup)
	accelGroupQuery           func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *AccelGroupEntry
	accelGroupUnlock          func(a *AccelGroup)

	accelLabelGetAccelWidget  func(a *AccelLabel) *Widget
	accelLabelGetAccelWidth   func(a *AccelLabel) uint
	accelLabelRefetch         func(a *AccelLabel) T.Gboolean
	accelLabelSetAccelClosure func(a *AccelLabel, accelClosure *O.Closure)
	accelLabelSetAccelWidget  func(a *AccelLabel, accelWidget *Widget)
)

func (a *AccelGroup) Activate(accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) T.Gboolean {
	return accelGroupActivate(a, accelQuark, acceleratable, accelKey, accelMods)
}
func (a *AccelGroup) Connect(accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags, closure *O.Closure) {
	accelGroupConnect(a, accelKey, accelMods, accelFlags, closure)
}
func (a *AccelGroup) ConnectByPath(
	accelPath string, closure *O.Closure) {
	accelGroupConnectByPath(a, accelPath, closure)
}
func (a *AccelGroup) Disconnect(closure *O.Closure) T.Gboolean {
	return accelGroupDisconnect(a, closure)
}
func (a *AccelGroup) DisconnectKey(accelKey uint, accelMods T.GdkModifierType) T.Gboolean {
	return accelGroupDisconnectKey(a, accelKey, accelMods)
}
func (a *AccelGroup) Find(findFunc AccelGroupFindFunc, data T.Gpointer) *AccelKey {
	return accelGroupFind(a, findFunc, data)
}
func (a *AccelGroup) GetIsLocked() T.Gboolean            { return accelGroupGetIsLocked(a) }
func (a *AccelGroup) GetModifierMask() T.GdkModifierType { return accelGroupGetModifierMask(a) }
func (a *AccelGroup) Lock()                              { accelGroupLock(a) }
func (a *AccelGroup) Query(accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *AccelGroupEntry {
	return accelGroupQuery(a, accelKey, accelMods, nEntries)
}
func (a *AccelGroup) Unlock() { accelGroupUnlock(a) }

func (a *AccelLabel) GetAccelWidget() *Widget { return accelLabelGetAccelWidget(a) }
func (a *AccelLabel) GetAccelWidth() uint     { return accelLabelGetAccelWidth(a) }
func (a *AccelLabel) Refetch() T.Gboolean     { return accelLabelRefetch(a) }
func (a *AccelLabel) SetAccelClosure(accelClosure *O.Closure) {
	accelLabelSetAccelClosure(a, accelClosure)
}
func (a *AccelLabel) SetAccelWidget(accelWidget *Widget) { accelLabelSetAccelWidget(a, accelWidget) }

type Accessible struct {
	Parent T.AtkObject
	Widget *Widget
}

var AccessibleGetType func() O.Type

var (
	accessibleConnectWidgetDestroyed func(a *Accessible)
	accessibleGetWidget              func(a *Accessible) *Widget
	accessibleSetWidget              func(a *Accessible, widget *Widget)
)

func (a *Accessible) SetWidget(widget *Widget) { accessibleSetWidget(a, widget) }
func (a *Accessible) GetWidget() *Widget       { return accessibleGetWidget(a) }
func (a *Accessible) ConnectWidgetDestroyed()  { accessibleConnectWidgetDestroyed(a) }

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
	ActionGetType func() O.Type
	ActionNew     func(name, label, tooltip, stockId string) *Action

	ActionGroupGetType func() O.Type
	ActionGroupNew     func(name string) *ActionGroup

	WidgetGetAction func(widget *Widget) *Action
)

var (
	actionActivate              func(a *Action)
	actionBlockActivate         func(a *Action)
	actionBlockActivateFrom     func(a *Action, proxy *Widget)
	actionConnectAccelerator    func(a *Action)
	actionConnectProxy          func(a *Action, proxy *Widget)
	actionCreateIcon            func(a *Action, iconSize IconSize) *Widget
	actionCreateMenu            func(a *Action) *Widget
	actionCreateMenuItem        func(a *Action) *Widget
	actionCreateToolItem        func(a *Action) *Widget
	actionDisconnectAccelerator func(a *Action)
	actionDisconnectProxy       func(a *Action, proxy *Widget)
	actionGetAccelClosure       func(a *Action) *O.Closure
	actionGetAccelPath          func(a *Action) string
	actionGetAlwaysShowImage    func(a *Action) T.Gboolean
	actionGetGicon              func(a *Action) *I.Icon
	actionGetIconName           func(a *Action) string
	actionGetIsImportant        func(a *Action) T.Gboolean
	actionGetLabel              func(a *Action) string
	actionGetName               func(a *Action) string
	actionGetProxies            func(a *Action) *T.GSList
	actionGetSensitive          func(a *Action) T.Gboolean
	actionGetShortLabel         func(a *Action) string
	actionGetStockId            func(a *Action) string
	actionGetTooltip            func(a *Action) string
	actionGetVisible            func(a *Action) T.Gboolean
	actionGetVisibleHorizontal  func(a *Action) T.Gboolean
	actionGetVisibleVertical    func(a *Action) T.Gboolean
	actionIsSensitive           func(a *Action) T.Gboolean
	actionIsVisible             func(a *Action) T.Gboolean
	actionSetAccelGroup         func(a *Action, accelGroup *AccelGroup)
	actionSetAccelPath          func(a *Action, accelPath string)
	actionSetAlwaysShowImage    func(a *Action, alwaysShow T.Gboolean)
	actionSetGicon              func(a *Action, icon *I.Icon)
	actionSetIconName           func(a *Action, iconName string)
	actionSetIsImportant        func(a *Action, isImportant T.Gboolean)
	actionSetLabel              func(a *Action, label string)
	actionSetSensitive          func(a *Action, sensitive T.Gboolean)
	actionSetShortLabel         func(a *Action, shortLabel string)
	actionSetStockId            func(a *Action, stockId string)
	actionSetTooltip            func(a *Action, tooltip string)
	actionSetVisible            func(a *Action, visible T.Gboolean)
	actionSetVisibleHorizontal  func(a *Action, visibleHorizontal T.Gboolean)
	actionSetVisibleVertical    func(a *Action, visibleVertical T.Gboolean)
	actionUnblockActivate       func(a *Action)
	actionUnblockActivateFrom   func(a *Action, proxy *Widget)
)

func (a *Action) Activate()                                { actionActivate(a) }
func (a *Action) BlockActivate()                           { actionBlockActivate(a) }
func (a *Action) BlockActivateFrom(proxy *Widget)          { actionBlockActivateFrom(a, proxy) }
func (a *Action) ConnectAccelerator()                      { actionConnectAccelerator(a) }
func (a *Action) ConnectProxy(proxy *Widget)               { actionConnectProxy(a, proxy) }
func (a *Action) CreateIcon(iconSize IconSize) *Widget     { return actionCreateIcon(a, iconSize) }
func (a *Action) CreateMenu() *Widget                      { return actionCreateMenu(a) }
func (a *Action) CreateMenuItem() *Widget                  { return actionCreateMenuItem(a) }
func (a *Action) CreateToolItem() *Widget                  { return actionCreateToolItem(a) }
func (a *Action) DisconnectAccelerator()                   { actionDisconnectAccelerator(a) }
func (a *Action) DisconnectProxy(proxy *Widget)            { actionDisconnectProxy(a, proxy) }
func (a *Action) GetAccelClosure() *O.Closure             { return actionGetAccelClosure(a) }
func (a *Action) GetAccelPath() string                     { return actionGetAccelPath(a) }
func (a *Action) GetAlwaysShowImage() T.Gboolean           { return actionGetAlwaysShowImage(a) }
func (a *Action) GetGicon() *I.Icon                        { return actionGetGicon(a) }
func (a *Action) GetIconName() string                      { return actionGetIconName(a) }
func (a *Action) GetIsImportant() T.Gboolean               { return actionGetIsImportant(a) }
func (a *Action) GetLabel() string                         { return actionGetLabel(a) }
func (a *Action) GetName() string                          { return actionGetName(a) }
func (a *Action) GetProxies() *T.GSList                    { return actionGetProxies(a) }
func (a *Action) GetSensitive() T.Gboolean                 { return actionGetSensitive(a) }
func (a *Action) GetShortLabel() string                    { return actionGetShortLabel(a) }
func (a *Action) GetStockId() string                       { return actionGetStockId(a) }
func (a *Action) GetTooltip() string                       { return actionGetTooltip(a) }
func (a *Action) GetVisible() T.Gboolean                   { return actionGetVisible(a) }
func (a *Action) GetVisibleHorizontal() T.Gboolean         { return actionGetVisibleHorizontal(a) }
func (a *Action) GetVisibleVertical() T.Gboolean           { return actionGetVisibleVertical(a) }
func (a *Action) IsSensitive() T.Gboolean                  { return actionIsSensitive(a) }
func (a *Action) IsVisible() T.Gboolean                    { return actionIsVisible(a) }
func (a *Action) SetAccelGroup(accelGroup *AccelGroup)     { actionSetAccelGroup(a, accelGroup) }
func (a *Action) SetAccelPath(accelPath string)            { actionSetAccelPath(a, accelPath) }
func (a *Action) SetAlwaysShowImage(alwaysShow T.Gboolean) { actionSetAlwaysShowImage(a, alwaysShow) }
func (a *Action) SetGicon(icon *I.Icon)                    { actionSetGicon(a, icon) }
func (a *Action) SetIconName(iconName string)              { actionSetIconName(a, iconName) }
func (a *Action) SetIsImportant(isImportant T.Gboolean)    { actionSetIsImportant(a, isImportant) }
func (a *Action) SetLabel(label string)                    { actionSetLabel(a, label) }
func (a *Action) SetSensitive(sensitive T.Gboolean)        { actionSetSensitive(a, sensitive) }
func (a *Action) SetShortLabel(shortLabel string)          { actionSetShortLabel(a, shortLabel) }
func (a *Action) SetStockId(stockId string)                { actionSetStockId(a, stockId) }
func (a *Action) SetTooltip(tooltip string)                { actionSetTooltip(a, tooltip) }
func (a *Action) SetVisible(visible T.Gboolean)            { actionSetVisible(a, visible) }
func (a *Action) SetVisibleHorizontal(visibleHorizontal T.Gboolean) {
	actionSetVisibleHorizontal(a, visibleHorizontal)
}
func (a *Action) SetVisibleVertical(visibleVertical T.Gboolean) {
	actionSetVisibleVertical(a, visibleVertical)
}
func (a *Action) UnblockActivate()                  { actionUnblockActivate(a) }
func (a *Action) UnblockActivateFrom(proxy *Widget) { actionUnblockActivateFrom(a, proxy) }

var (
	actionGroupAddAction            func(a *ActionGroup, action *Action)
	actionGroupAddActions           func(a *ActionGroup, entries *ActionEntry, nEntries uint, userData T.Gpointer)
	actionGroupAddActionsFull       func(a *ActionGroup, entries *ActionEntry, nEntries uint, userData T.Gpointer, destroy T.GDestroyNotify)
	actionGroupAddActionWithAccel   func(a *ActionGroup, action *Action, accelerator string)
	actionGroupAddRadioActions      func(a *ActionGroup, entries *RadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer)
	actionGroupAddRadioActionsFull  func(a *ActionGroup, entries *RadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer, destroy T.GDestroyNotify)
	actionGroupAddToggleActions     func(a *ActionGroup, entries *ToggleActionEntry, nEntries uint, userData T.Gpointer)
	actionGroupAddToggleActionsFull func(a *ActionGroup, entries *ToggleActionEntry, nEntries uint, userData T.Gpointer, destroy T.GDestroyNotify)
	actionGroupGetAction            func(a *ActionGroup, actionName string) *Action
	actionGroupGetName              func(a *ActionGroup) string
	actionGroupGetSensitive         func(a *ActionGroup) T.Gboolean
	actionGroupGetVisible           func(a *ActionGroup) T.Gboolean
	actionGroupListActions          func(a *ActionGroup) *T.GList
	actionGroupRemoveAction         func(a *ActionGroup, action *Action)
	actionGroupSetSensitive         func(a *ActionGroup, sensitive T.Gboolean)
	actionGroupSetTranslateFunc     func(a *ActionGroup, f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify)
	actionGroupSetTranslationDomain func(a *ActionGroup, domain string)
	actionGroupSetVisible           func(a *ActionGroup, visible T.Gboolean)
	actionGroupTranslateString      func(a *ActionGroup, str string) string
)

func (a *ActionGroup) AddAction(action *Action) { actionGroupAddAction(a, action) }
func (a *ActionGroup) AddActions(entries *ActionEntry, nEntries uint, userData T.Gpointer) {
	actionGroupAddActions(a, entries, nEntries, userData)
}
func (a *ActionGroup) AddActionsFull(entries *ActionEntry, nEntries uint, userData T.Gpointer, destroy T.GDestroyNotify) {
	actionGroupAddActionsFull(a, entries, nEntries, userData, destroy)
}
func (a *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	actionGroupAddActionWithAccel(a, action, accelerator)
}
func (a *ActionGroup) AddRadioActions(entries *RadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer) {
	actionGroupAddRadioActions(a, entries, nEntries, value, onChange, userData)
}
func (a *ActionGroup) AddRadioActionsFull(entries *RadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer, destroy T.GDestroyNotify) {
	actionGroupAddRadioActionsFull(a, entries, nEntries, value, onChange, userData, destroy)
}
func (a *ActionGroup) AddToggleActions(entries *ToggleActionEntry, nEntries uint, userData T.Gpointer) {
	actionGroupAddToggleActions(a, entries, nEntries, userData)
}
func (a *ActionGroup) AddToggleActionsFull(entries *ToggleActionEntry, nEntries uint, userData T.Gpointer, destroy T.GDestroyNotify) {
	actionGroupAddToggleActionsFull(a, entries, nEntries, userData, destroy)
}
func (a *ActionGroup) GetAction(actionName string) *Action { return actionGroupGetAction(a, actionName) }
func (a *ActionGroup) GetName() string                     { return actionGroupGetName(a) }
func (a *ActionGroup) GetSensitive() T.Gboolean            { return actionGroupGetSensitive(a) }
func (a *ActionGroup) GetVisible() T.Gboolean              { return actionGroupGetVisible(a) }
func (a *ActionGroup) ListActions() *T.GList               { return actionGroupListActions(a) }
func (a *ActionGroup) RemoveAction(action *Action)         { actionGroupRemoveAction(a, action) }
func (a *ActionGroup) SetSensitive(sensitive T.Gboolean)   { actionGroupSetSensitive(a, sensitive) }
func (a *ActionGroup) SetTranslateFunc(f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify) {
	actionGroupSetTranslateFunc(a, f, data, notify)
}
func (a *ActionGroup) SetTranslationDomain(domain string) { actionGroupSetTranslationDomain(a, domain) }
func (a *ActionGroup) SetVisible(visible T.Gboolean)      { actionGroupSetVisible(a, visible) }
func (a *ActionGroup) TranslateString(str string) string  { return actionGroupTranslateString(a, str) }

type Activatable struct{}

var ActivatableGetType func() O.Type

var (
	activatableDoSetRelatedAction     func(a *Activatable, action *Action)
	activatableGetRelatedAction       func(a *Activatable) *Action
	activatableGetUseActionAppearance func(a *Activatable) T.Gboolean
	activatableSetRelatedAction       func(a *Activatable, action *Action)
	activatableSetUseActionAppearance func(a *Activatable, useAppearance T.Gboolean)
	activatableSyncActionProperties   func(a *Activatable, action *Action)
)

func (a *Activatable) DoSetRelatedAction(action *Action)  { activatableDoSetRelatedAction(a, action) }
func (a *Activatable) GetRelatedAction() *Action          { return activatableGetRelatedAction(a) }
func (a *Activatable) GetUseActionAppearance() T.Gboolean { return activatableGetUseActionAppearance(a) }
func (a *Activatable) SetRelatedAction(action *Action)    { activatableSetRelatedAction(a, action) }
func (a *Activatable) SetUseActionAppearance(useAppearance T.Gboolean) {
	activatableSetUseActionAppearance(a, useAppearance)
}
func (a *Activatable) SyncActionProperties(action *Action) { activatableSyncActionProperties(a, action) }

type Adjustment struct {
	Parent        Object
	Lower         float64
	Upper         float64
	Value         float64
	StepIncrement float64
	PageIncrement float64
	PageSize      float64
}

var (
	AdjustmentGetType func() O.Type
	AdjustmentNew     func(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Object
)

var (
	adjustmentChanged          func(a *Adjustment)
	adjustmentClampPage        func(a *Adjustment, lower, upper float64)
	adjustmentConfigure        func(a *Adjustment, value, lower, upper, stepIncrement, pageIncrement, pageSize float64)
	adjustmentGetLower         func(a *Adjustment) float64
	adjustmentGetPageIncrement func(a *Adjustment) float64
	adjustmentGetPageSize      func(a *Adjustment) float64
	adjustmentGetStepIncrement func(a *Adjustment) float64
	adjustmentGetUpper         func(a *Adjustment) float64
	adjustmentGetValue         func(a *Adjustment) float64
	adjustmentSetLower         func(a *Adjustment, lower float64)
	adjustmentSetPageIncrement func(a *Adjustment, pageIncrement float64)
	adjustmentSetPageSize      func(a *Adjustment, pageSize float64)
	adjustmentSetStepIncrement func(a *Adjustment, stepIncrement float64)
	adjustmentSetUpper         func(a *Adjustment, upper float64)
	adjustmentSetValue         func(a *Adjustment, value float64)
	adjustmentValueChanged     func(a *Adjustment)
)

func (a *Adjustment) Changed()                       { adjustmentChanged(a) }
func (a *Adjustment) ClampPage(lower, upper float64) { adjustmentClampPage(a, lower, upper) }
func (a *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	adjustmentConfigure(a, value, lower, upper, stepIncrement, pageIncrement, pageSize)
}
func (a *Adjustment) GetLower() float64         { return adjustmentGetLower(a) }
func (a *Adjustment) GetPageIncrement() float64 { return adjustmentGetPageIncrement(a) }
func (a *Adjustment) GetPageSize() float64      { return adjustmentGetPageSize(a) }
func (a *Adjustment) GetStepIncrement() float64 { return adjustmentGetStepIncrement(a) }
func (a *Adjustment) GetUpper() float64         { return adjustmentGetUpper(a) }
func (a *Adjustment) GetValue() float64         { return adjustmentGetValue(a) }
func (a *Adjustment) SetLower(lower float64)    { adjustmentSetLower(a, lower) }
func (a *Adjustment) SetPageIncrement(pageIncrement float64) {
	adjustmentSetPageIncrement(a, pageIncrement)
}
func (a *Adjustment) SetPageSize(pageSize float64) { adjustmentSetPageSize(a, pageSize) }
func (a *Adjustment) SetStepIncrement(stepIncrement float64) {
	adjustmentSetStepIncrement(a, stepIncrement)
}
func (a *Adjustment) SetUpper(upper float64) { adjustmentSetUpper(a, upper) }
func (a *Adjustment) SetValue(value float64) { adjustmentSetValue(a, value) }
func (a *Adjustment) ValueChanged()          { adjustmentValueChanged(a) }

type Alignment struct {
	Bin    Bin
	Xalign float32
	Yalign float32
	Xscale float32
	Yscale float32
}

var (
	AlignmentGetType func() O.Type
	AlignmentNew     func(xalign, yalign, xscale, yscale float32) *Widget

	alignmentSet        func(a *Alignment, xalign, yalign, xscale, yscale float32)
	alignmentSetPadding func(a *Alignment, paddingTop, paddingBottom, paddingLeft, paddingRight uint)
	alignmentGetPadding func(a *Alignment, paddingTop, paddingBottom, paddingLeft, paddingRight *uint)
)

func (a *Alignment) Set(xalign, yalign, xscale, yscale float32) {
	alignmentSet(a, xalign, yalign, xscale, yscale)
}
func (a *Alignment) SetPadding(paddingTop, paddingBottom, paddingLeft, paddingRight uint) {
	alignmentSetPadding(a, paddingTop, paddingBottom, paddingLeft, paddingRight)
}
func (a *Alignment) GetPadding(paddingTop, paddingBottom, paddingLeft, paddingRight *uint) {
	alignmentGetPadding(a, paddingTop, paddingBottom, paddingLeft, paddingRight)
}

type Allocation D.Rectangle

var AlternativeDialogButtonOrder func(
	screen *D.Screen) T.Gboolean

type AnchorType Enum

const (
	ANCHOR_CENTER AnchorType = iota
	ANCHOR_NORTH
	ANCHOR_NORTH_WEST
	ANCHOR_NORTH_EAST
	ANCHOR_SOUTH
	ANCHOR_SOUTH_WEST
	ANCHOR_SOUTH_EAST
	ANCHOR_WEST
	ANCHOR_EAST
	ANCHOR_N  = ANCHOR_NORTH
	ANCHOR_NW = ANCHOR_NORTH_WEST
	ANCHOR_NE = ANCHOR_NORTH_EAST
	ANCHOR_S  = ANCHOR_SOUTH
	ANCHOR_SW = ANCHOR_SOUTH_WEST
	ANCHOR_SE = ANCHOR_SOUTH_EAST
	ANCHOR_W  = ANCHOR_WEST
	ANCHOR_E  = ANCHOR_EAST
)

var AnchorTypeGetType func() O.Type

type Arg struct { //TODO(t):Fix union
	Type O.Type
	Name *T.Gchar
	// Union
	// Gchar char_data;
	// uchar_data  Guchar;
	// bool_data  Gboolean;
	// int_data  int;
	// uint_data  uint;
	// long_data  glong;
	// ulong_data  Gulong;
	// float_data  float32;
	// Double_data  float64;
	// string_data  *Gchar;
	// object_data  *Object;
	// pointer_data  Gpointer;
	// signal_data struct{f GCallback; d Gpointer}
}

type ArgFlags Enum

const (
	ARG_READABLE       ArgFlags = ArgFlags(O.PARAM_READABLE)
	ARG_WRITABLE                = ArgFlags(O.PARAM_WRITABLE)
	ARG_CONSTRUCT               = ArgFlags(O.PARAM_CONSTRUCT)
	ARG_CONSTRUCT_ONLY          = ArgFlags(O.PARAM_CONSTRUCT_ONLY)
	ARG_CHILD_ARG      ArgFlags = 1 << iota
)

var ArgFlagsGetType func() O.Type

type (
	Assistant struct {
		Parent  Window
		Cancel  *Widget
		Forward *Widget
		Back    *Widget
		Apply   *Widget
		Close   *Widget
		Last    *Widget
		_       *struct{}
	}

	AssistantPageFunc func(currentPage int, data T.Gpointer) int
)

type AssistantPageType Enum

const (
	ASSISTANT_PAGE_CONTENT AssistantPageType = iota
	ASSISTANT_PAGE_INTRO
	ASSISTANT_PAGE_CONFIRM
	ASSISTANT_PAGE_SUMMARY
	ASSISTANT_PAGE_PROGRESS
)

var (
	AssistantGetType func() O.Type
	AssistantNew     func() *Widget

	AssistantPageTypeGetType func() O.Type
)

var (
	assistantAddActionWidget    func(a *Assistant, child *Widget)
	assistantAppendPage         func(a *Assistant, page *Widget) int
	assistantCommit             func(a *Assistant)
	assistantGetCurrentPage     func(a *Assistant) int
	assistantGetNPages          func(a *Assistant) int
	assistantGetNthPage         func(a *Assistant, pageNum int) *Widget
	assistantGetPageComplete    func(a *Assistant, page *Widget) T.Gboolean
	assistantGetPageHeaderImage func(a *Assistant, page *Widget) *D.Pixbuf
	assistantGetPageSideImage   func(a *Assistant, page *Widget) *D.Pixbuf
	assistantGetPageTitle       func(a *Assistant, page *Widget) string
	assistantGetPageType        func(a *Assistant, page *Widget) AssistantPageType
	assistantInsertPage         func(a *Assistant, page *Widget, position int) int
	assistantPrependPage        func(a *Assistant, page *Widget) int
	assistantRemoveActionWidget func(a *Assistant, child *Widget)
	assistantSetCurrentPage     func(a *Assistant, pageNum int)
	assistantSetForwardPageFunc func(a *Assistant, pageFunc AssistantPageFunc, data T.Gpointer, destroy T.GDestroyNotify)
	assistantSetPageComplete    func(a *Assistant, page *Widget, complete T.Gboolean)
	assistantSetPageHeaderImage func(a *Assistant, page *Widget, pixbuf *D.Pixbuf)
	assistantSetPageSideImage   func(a *Assistant, page *Widget, pixbuf *D.Pixbuf)
	assistantSetPageTitle       func(a *Assistant, page *Widget, title string)
	assistantSetPageType        func(a *Assistant, page *Widget, t AssistantPageType)
	assistantUpdateButtonsState func(a *Assistant)
)

func (a *Assistant) AddActionWidget(child *Widget)           { assistantAddActionWidget(a, child) }
func (a *Assistant) AppendPage(page *Widget) int             { return assistantAppendPage(a, page) }
func (a *Assistant) Commit()                                 { assistantCommit(a) }
func (a *Assistant) GetCurrentPage() int                     { return assistantGetCurrentPage(a) }
func (a *Assistant) GetNPages() int                          { return assistantGetNPages(a) }
func (a *Assistant) GetNthPage(pageNum int) *Widget          { return assistantGetNthPage(a, pageNum) }
func (a *Assistant) GetPageComplete(page *Widget) T.Gboolean { return assistantGetPageComplete(a, page) }
func (a *Assistant) GetPageHeaderImage(page *Widget) *D.Pixbuf {
	return assistantGetPageHeaderImage(a, page)
}
func (a *Assistant) GetPageSideImage(page *Widget) *D.Pixbuf {
	return assistantGetPageSideImage(a, page)
}
func (a *Assistant) GetPageTitle(page *Widget) string           { return assistantGetPageTitle(a, page) }
func (a *Assistant) GetPageType(page *Widget) AssistantPageType { return assistantGetPageType(a, page) }
func (a *Assistant) InsertPage(page *Widget, position int) int {
	return assistantInsertPage(a, page, position)
}
func (a *Assistant) PrependPage(page *Widget) int     { return assistantPrependPage(a, page) }
func (a *Assistant) RemoveActionWidget(child *Widget) { assistantRemoveActionWidget(a, child) }
func (a *Assistant) SetCurrentPage(pageNum int)       { assistantSetCurrentPage(a, pageNum) }
func (a *Assistant) SetForwardPageFunc(pageFunc AssistantPageFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	assistantSetForwardPageFunc(a, pageFunc, data, destroy)
}
func (a *Assistant) SetPageComplete(page *Widget, complete T.Gboolean) {
	assistantSetPageComplete(a, page, complete)
}
func (a *Assistant) SetPageHeaderImage(page *Widget, pixbuf *D.Pixbuf) {
	assistantSetPageHeaderImage(a, page, pixbuf)
}
func (a *Assistant) SetPageSideImage(page *Widget, pixbuf *D.Pixbuf) {
	assistantSetPageSideImage(a, page, pixbuf)
}
func (a *Assistant) SetPageTitle(page *Widget, title string)       { assistantSetPageTitle(a, page, title) }
func (a *Assistant) SetPageType(page *Widget, t AssistantPageType) { assistantSetPageType(a, page, t) }
func (a *Assistant) UpdateButtonsState()                           { assistantUpdateButtonsState(a) }

type Arrow struct {
	Misc       Misc
	ArrowType  int16
	ShadowType int16
}

type ArrowType Enum

const (
	ARROW_UP ArrowType = iota
	ARROW_DOWN
	ARROW_LEFT
	ARROW_RIGHT
	ARROW_NONE
)

type ArrowPlacement Enum

const (
	ARROWS_BOTH ArrowPlacement = iota
	ARROWS_START
	ARROWS_END
)

var (
	ArrowGetType          func() O.Type
	ArrowTypeGetType      func() O.Type
	ArrowPlacementGetType func() O.Type
)

var (
	arrowSet func(a *Arrow, arrowType ArrowType, shadowType ShadowType)

	arrowNew func(a ArrowType, shadowType ShadowType) *Widget
)

func (a *Arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	arrowSet(a, arrowType, shadowType)
}

func (a ArrowType) New(shadowType ShadowType) *Widget { return arrowNew(a, shadowType) }

type AspectFrame struct {
	Frame            Frame
	Xalign           float32
	Yalign           float32
	Ratio            float32
	ObeyChild        T.Gboolean
	CenterAllocation Allocation
}

var (
	AspectFrameGetType func() O.Type
	AspectFrameNew     func(label string, xalign, yalign, ratio float32, obeyChild T.Gboolean) *Widget
)

var aspectFrameSet func(a *AspectFrame, xalign, yalign, ratio float32, obeyChild T.Gboolean)

func (a *AspectFrame) Set(xalign, yalign, ratio float32, obeyChild T.Gboolean) {
	aspectFrameSet(a, xalign, yalign, ratio, obeyChild)
}

var AttachOptionsGetType func() O.Type

type AttachOptions Enum

const (
	EXPAND AttachOptions = 1 << iota
	SHRINK
	FILL
)
