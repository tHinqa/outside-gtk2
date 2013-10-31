// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	A "github.com/tHinqa/outside-gtk2/atk"
	D "github.com/tHinqa/outside-gtk2/gdk"
	I "github.com/tHinqa/outside-gtk2/gio"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
	"unsafe"
)

type (
	AboutDialog struct {
		Parent Dialog
		_      *struct{}
	}

	AboutDialogActivateLink func(
		about *AboutDialog, link string, data T.Gpointer)
)

var (
	AboutDialogGetType func() O.Type
	AboutDialogNew     func() *Widget // 2.6

	AboutDialogSetEmailHook func(f AboutDialogActivateLink, data T.Gpointer, destroy T.GDestroyNotify) AboutDialogActivateLink // 2.6
	AboutDialogSetUrlHook   func(f AboutDialogActivateLink, data T.Gpointer, destroy T.GDestroyNotify) AboutDialogActivateLink // 2.6

	ShowAboutDialog func(parent *Window, firstPropertyName string, v ...VArg) // 2.6

	aboutDialogGetArtists      func(a *AboutDialog) []string  // 2.6
	aboutDialogGetAuthors      func(a *AboutDialog) []string  // 2.6
	aboutDialogGetComments     func(a *AboutDialog) string    // 2.6
	aboutDialogGetCopyright    func(a *AboutDialog) string    // 2.6
	aboutDialogGetDocumenters  func(a *AboutDialog) []string  // 2.6
	aboutDialogGetLicense      func(a *AboutDialog) string    // 2.6
	aboutDialogGetLogo         func(a *AboutDialog) *D.Pixbuf // 2.6
	aboutDialogGetLogoIconName func(a *AboutDialog) string    // 2.6
	// 2.6 Deprecated 2.12 aboutDialogGetName func(a *AboutDialog) string
	aboutDialogGetProgramName       func(a *AboutDialog) string                // 2.12
	aboutDialogGetTranslatorCredits func(a *AboutDialog) string                // 2.6
	aboutDialogGetVersion           func(a *AboutDialog) string                // 2.6
	aboutDialogGetWebsite           func(a *AboutDialog) string                // 2.6
	aboutDialogGetWebsiteLabel      func(a *AboutDialog) string                // 2.6
	aboutDialogGetWrapLicense       func(a *AboutDialog) bool                  // 2.8
	aboutDialogSetArtists           func(a *AboutDialog, artists []string)     // 2.6
	aboutDialogSetAuthors           func(a *AboutDialog, authors []string)     // 2.6
	aboutDialogSetComments          func(a *AboutDialog, comments string)      // 2.6
	aboutDialogSetCopyright         func(a *AboutDialog, copyright string)     // 2.6
	aboutDialogSetDocumenters       func(a *AboutDialog, documenters []string) // 2.6
	aboutDialogSetLicense           func(a *AboutDialog, license string)       // 2.6
	aboutDialogSetLogo              func(a *AboutDialog, logo *D.Pixbuf)       // 2.6
	aboutDialogSetLogoIconName      func(a *AboutDialog, iconName string)      // 2.6
	// 2.6 Deprecated 2.12 aboutDialogSetName func(a *AboutDialog, name string)
	aboutDialogSetProgramName       func(a *AboutDialog, name string)              // 2.12
	aboutDialogSetTranslatorCredits func(a *AboutDialog, translatorCredits string) // 2.6
	aboutDialogSetVersion           func(a *AboutDialog, version string)           // 2.6
	aboutDialogSetWebsite           func(a *AboutDialog, website string)           // 2.6
	aboutDialogSetWebsiteLabel      func(a *AboutDialog, websiteLabel string)      // 2.6
	aboutDialogSetWrapLicense       func(a *AboutDialog, wrapLicense bool)         // 2.8
)

func NewAboutDialog() (w *Widget, a *AboutDialog) {
	w = AboutDialogNew()
	a = (*AboutDialog)(unsafe.Pointer(w))
	return
}

func (a *AboutDialog) Artists() []string                   { return aboutDialogGetArtists(a) }
func (a *AboutDialog) Authors() []string                   { return aboutDialogGetAuthors(a) }
func (a *AboutDialog) Comments() string                    { return aboutDialogGetComments(a) }
func (a *AboutDialog) Copyright() string                   { return aboutDialogGetCopyright(a) }
func (a *AboutDialog) Documenters() []string               { return aboutDialogGetDocumenters(a) }
func (a *AboutDialog) License() string                     { return aboutDialogGetLicense(a) }
func (a *AboutDialog) Logo() *D.Pixbuf                     { return aboutDialogGetLogo(a) }
func (a *AboutDialog) LogoIconName() string                { return aboutDialogGetLogoIconName(a) }
func (a *AboutDialog) ProgramName() string                 { return aboutDialogGetProgramName(a) }
func (a *AboutDialog) TranslatorCredits() string           { return aboutDialogGetTranslatorCredits(a) }
func (a *AboutDialog) Version() string                     { return aboutDialogGetVersion(a) }
func (a *AboutDialog) Website() string                     { return aboutDialogGetWebsite(a) }
func (a *AboutDialog) WebsiteLabel() string                { return aboutDialogGetWebsiteLabel(a) }
func (a *AboutDialog) WrapLicense() bool                   { return aboutDialogGetWrapLicense(a) }
func (a *AboutDialog) SetArtists(artists []string)         { aboutDialogSetArtists(a, artists) }
func (a *AboutDialog) SetAuthors(authors []string)         { aboutDialogSetAuthors(a, authors) }
func (a *AboutDialog) SetComments(comments string)         { aboutDialogSetComments(a, comments) }
func (a *AboutDialog) SetCopyright(copyright string)       { aboutDialogSetCopyright(a, copyright) }
func (a *AboutDialog) SetDocumenters(documenters []string) { aboutDialogSetDocumenters(a, documenters) }
func (a *AboutDialog) SetLicense(license string)           { aboutDialogSetLicense(a, license) }
func (a *AboutDialog) SetLogo(logo *D.Pixbuf)              { aboutDialogSetLogo(a, logo) }
func (a *AboutDialog) SetLogoIconName(iconName string)     { aboutDialogSetLogoIconName(a, iconName) }

// Deprecated 2.12 func (a *AboutDialog) SetName(name string)                  { aboutDialogSetName(a, name) }
func (a *AboutDialog) SetProgramName(name string) { aboutDialogSetProgramName(a, name) }
func (a *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	aboutDialogSetTranslatorCredits(a, translatorCredits)
}
func (a *AboutDialog) SetVersion(version string) { aboutDialogSetVersion(a, version) }
func (a *AboutDialog) SetWebsite(website string) { aboutDialogSetWebsite(a, website) }
func (a *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	aboutDialogSetWebsiteLabel(a, websiteLabel)
}
func (a *AboutDialog) SetWrapLicense(wrapLicense bool) {
	aboutDialogSetWrapLicense(a, wrapLicense)
}

// Deprecated 2.12 func (a *AboutDialog) GetName() string { return aboutDialogGetName(a) }

/*
Properties
"artists"                  GStrv*                : Read / Write
"authors"                  GStrv*                : Read / Write
"comments"                 gchar*                : Read / Write
"copyright"                gchar*                : Read / Write
"documenters"              GStrv*                : Read / Write
"license"                  gchar*                : Read / Write
"logo"                     GdkPixbuf*            : Read / Write
"logo-icon-name"           gchar*                : Read / Write
"program-name"             gchar*                : Read / Write
"translator-credits"       gchar*                : Read / Write
"version"                  gchar*                : Read / Write
"website"                  gchar*                : Read / Write
"website-label"            gchar*                : Read / Write
"wrap-license"             gboolean              : Read / Write
*/

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
		modifiers T.GdkModifierType) bool
)

var (
	AccelGroupsActivate func(object *T.GObject, accelKey uint,
		accelMods T.GdkModifierType) bool

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
		closure *O.Closure, data T.Gpointer) bool

	AccelMapForeachFunc func(
		data T.Gpointer,
		accelPath string,
		accelKey uint,
		accelMods T.GdkModifierType,
		changed bool)
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
	AccelMapChangeEntry       func(accelPath string, accelKey uint, accelMods T.GdkModifierType, replace bool) bool
	AccelMapForeach           func(data T.Gpointer, foreachFunc AccelMapForeachFunc)
	AccelMapForeachUnfiltered func(data T.Gpointer, foreachFunc AccelMapForeachFunc)
	AccelMapGet               func() *AccelMap
	AccelMapLoad              func(fileName string)
	AccelMapLoadFd            func(fd int)
	AccelMapLoadScanner       func(scanner *T.GScanner)
	AccelMapLockPath          func(accelPath string)
	AccelMapLookupEntry       func(accelPath string, key *AccelKey) bool
	AccelMapSave              func(fileName string)
	AccelMapSaveFd            func(fd int)
	AccelMapUnlockPath        func(accelPath string)

	AccelFlagsGetType func() O.Type
)

var (
	accelGroupActivate        func(a *AccelGroup, accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) bool
	accelGroupConnect         func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags, closure *O.Closure)
	accelGroupConnectByPath   func(a *AccelGroup, accelPath string, closure *O.Closure)
	accelGroupDisconnect      func(a *AccelGroup, closure *O.Closure) bool
	accelGroupDisconnectKey   func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType) bool
	accelGroupFind            func(a *AccelGroup, findFunc AccelGroupFindFunc, data T.Gpointer) *AccelKey
	accelGroupGetIsLocked     func(a *AccelGroup) bool
	accelGroupGetModifierMask func(a *AccelGroup) T.GdkModifierType
	accelGroupLock            func(a *AccelGroup)
	accelGroupQuery           func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *AccelGroupEntry
	accelGroupUnlock          func(a *AccelGroup)

	accelLabelGetAccelWidget  func(a *AccelLabel) *Widget
	accelLabelGetAccelWidth   func(a *AccelLabel) uint
	accelLabelRefetch         func(a *AccelLabel) bool
	accelLabelSetAccelClosure func(a *AccelLabel, accelClosure *O.Closure)
	accelLabelSetAccelWidget  func(a *AccelLabel, accelWidget *Widget)
)

func (a *AccelGroup) Activate(accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) bool {
	return accelGroupActivate(a, accelQuark, acceleratable, accelKey, accelMods)
}
func (a *AccelGroup) Connect(accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags, closure *O.Closure) {
	accelGroupConnect(a, accelKey, accelMods, accelFlags, closure)
}
func (a *AccelGroup) ConnectByPath(
	accelPath string, closure *O.Closure) {
	accelGroupConnectByPath(a, accelPath, closure)
}
func (a *AccelGroup) Disconnect(closure *O.Closure) bool {
	return accelGroupDisconnect(a, closure)
}
func (a *AccelGroup) DisconnectKey(accelKey uint, accelMods T.GdkModifierType) bool {
	return accelGroupDisconnectKey(a, accelKey, accelMods)
}
func (a *AccelGroup) Find(findFunc AccelGroupFindFunc, data T.Gpointer) *AccelKey {
	return accelGroupFind(a, findFunc, data)
}
func (a *AccelGroup) GetIsLocked() bool                  { return accelGroupGetIsLocked(a) }
func (a *AccelGroup) GetModifierMask() T.GdkModifierType { return accelGroupGetModifierMask(a) }
func (a *AccelGroup) Lock()                              { accelGroupLock(a) }
func (a *AccelGroup) Query(accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *AccelGroupEntry {
	return accelGroupQuery(a, accelKey, accelMods, nEntries)
}
func (a *AccelGroup) Unlock() { accelGroupUnlock(a) }

func (a *AccelLabel) GetAccelWidget() *Widget { return accelLabelGetAccelWidget(a) }
func (a *AccelLabel) GetAccelWidth() uint     { return accelLabelGetAccelWidth(a) }
func (a *AccelLabel) Refetch() bool           { return accelLabelRefetch(a) }
func (a *AccelLabel) SetAccelClosure(accelClosure *O.Closure) {
	accelLabelSetAccelClosure(a, accelClosure)
}
func (a *AccelLabel) SetAccelWidget(accelWidget *Widget) { accelLabelSetAccelWidget(a, accelWidget) }

type Accessible struct {
	Parent A.Object
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
	actionGetAlwaysShowImage    func(a *Action) bool
	actionGetGicon              func(a *Action) *I.Icon
	actionGetIconName           func(a *Action) string
	actionGetIsImportant        func(a *Action) bool
	actionGetLabel              func(a *Action) string
	actionGetName               func(a *Action) string
	actionGetProxies            func(a *Action) *T.GSList
	actionGetSensitive          func(a *Action) bool
	actionGetShortLabel         func(a *Action) string
	actionGetStockId            func(a *Action) string
	actionGetTooltip            func(a *Action) string
	actionGetVisible            func(a *Action) bool
	actionGetVisibleHorizontal  func(a *Action) bool
	actionGetVisibleVertical    func(a *Action) bool
	actionIsSensitive           func(a *Action) bool
	actionIsVisible             func(a *Action) bool
	actionSetAccelGroup         func(a *Action, accelGroup *AccelGroup)
	actionSetAccelPath          func(a *Action, accelPath string)
	actionSetAlwaysShowImage    func(a *Action, alwaysShow bool)
	actionSetGicon              func(a *Action, icon *I.Icon)
	actionSetIconName           func(a *Action, iconName string)
	actionSetIsImportant        func(a *Action, isImportant bool)
	actionSetLabel              func(a *Action, label string)
	actionSetSensitive          func(a *Action, sensitive bool)
	actionSetShortLabel         func(a *Action, shortLabel string)
	actionSetStockId            func(a *Action, stockId string)
	actionSetTooltip            func(a *Action, tooltip string)
	actionSetVisible            func(a *Action, visible bool)
	actionSetVisibleHorizontal  func(a *Action, visibleHorizontal bool)
	actionSetVisibleVertical    func(a *Action, visibleVertical bool)
	actionUnblockActivate       func(a *Action)
	actionUnblockActivateFrom   func(a *Action, proxy *Widget)
)

func (a *Action) Activate()                            { actionActivate(a) }
func (a *Action) BlockActivate()                       { actionBlockActivate(a) }
func (a *Action) BlockActivateFrom(proxy *Widget)      { actionBlockActivateFrom(a, proxy) }
func (a *Action) ConnectAccelerator()                  { actionConnectAccelerator(a) }
func (a *Action) ConnectProxy(proxy *Widget)           { actionConnectProxy(a, proxy) }
func (a *Action) CreateIcon(iconSize IconSize) *Widget { return actionCreateIcon(a, iconSize) }
func (a *Action) CreateMenu() *Widget                  { return actionCreateMenu(a) }
func (a *Action) CreateMenuItem() *Widget              { return actionCreateMenuItem(a) }
func (a *Action) CreateToolItem() *Widget              { return actionCreateToolItem(a) }
func (a *Action) DisconnectAccelerator()               { actionDisconnectAccelerator(a) }
func (a *Action) DisconnectProxy(proxy *Widget)        { actionDisconnectProxy(a, proxy) }
func (a *Action) GetAccelClosure() *O.Closure          { return actionGetAccelClosure(a) }
func (a *Action) GetAccelPath() string                 { return actionGetAccelPath(a) }
func (a *Action) GetAlwaysShowImage() bool             { return actionGetAlwaysShowImage(a) }
func (a *Action) GetGicon() *I.Icon                    { return actionGetGicon(a) }
func (a *Action) GetIconName() string                  { return actionGetIconName(a) }
func (a *Action) GetIsImportant() bool                 { return actionGetIsImportant(a) }
func (a *Action) GetLabel() string                     { return actionGetLabel(a) }
func (a *Action) GetName() string                      { return actionGetName(a) }
func (a *Action) GetProxies() *T.GSList                { return actionGetProxies(a) }
func (a *Action) GetSensitive() bool                   { return actionGetSensitive(a) }
func (a *Action) GetShortLabel() string                { return actionGetShortLabel(a) }
func (a *Action) GetStockId() string                   { return actionGetStockId(a) }
func (a *Action) GetTooltip() string                   { return actionGetTooltip(a) }
func (a *Action) GetVisible() bool                     { return actionGetVisible(a) }
func (a *Action) GetVisibleHorizontal() bool           { return actionGetVisibleHorizontal(a) }
func (a *Action) GetVisibleVertical() bool             { return actionGetVisibleVertical(a) }
func (a *Action) IsSensitive() bool                    { return actionIsSensitive(a) }
func (a *Action) IsVisible() bool                      { return actionIsVisible(a) }
func (a *Action) SetAccelGroup(accelGroup *AccelGroup) { actionSetAccelGroup(a, accelGroup) }
func (a *Action) SetAccelPath(accelPath string)        { actionSetAccelPath(a, accelPath) }
func (a *Action) SetAlwaysShowImage(alwaysShow bool)   { actionSetAlwaysShowImage(a, alwaysShow) }
func (a *Action) SetGicon(icon *I.Icon)                { actionSetGicon(a, icon) }
func (a *Action) SetIconName(iconName string)          { actionSetIconName(a, iconName) }
func (a *Action) SetIsImportant(isImportant bool)      { actionSetIsImportant(a, isImportant) }
func (a *Action) SetLabel(label string)                { actionSetLabel(a, label) }
func (a *Action) SetSensitive(sensitive bool)          { actionSetSensitive(a, sensitive) }
func (a *Action) SetShortLabel(shortLabel string)      { actionSetShortLabel(a, shortLabel) }
func (a *Action) SetStockId(stockId string)            { actionSetStockId(a, stockId) }
func (a *Action) SetTooltip(tooltip string)            { actionSetTooltip(a, tooltip) }
func (a *Action) SetVisible(visible bool)              { actionSetVisible(a, visible) }
func (a *Action) SetVisibleHorizontal(visibleHorizontal bool) {
	actionSetVisibleHorizontal(a, visibleHorizontal)
}
func (a *Action) SetVisibleVertical(visibleVertical bool) {
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
	actionGroupGetSensitive         func(a *ActionGroup) bool
	actionGroupGetVisible           func(a *ActionGroup) bool
	actionGroupListActions          func(a *ActionGroup) *T.GList
	actionGroupRemoveAction         func(a *ActionGroup, action *Action)
	actionGroupSetSensitive         func(a *ActionGroup, sensitive bool)
	actionGroupSetTranslateFunc     func(a *ActionGroup, f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify)
	actionGroupSetTranslationDomain func(a *ActionGroup, domain string)
	actionGroupSetVisible           func(a *ActionGroup, visible bool)
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
func (a *ActionGroup) GetSensitive() bool                  { return actionGroupGetSensitive(a) }
func (a *ActionGroup) GetVisible() bool                    { return actionGroupGetVisible(a) }
func (a *ActionGroup) ListActions() *T.GList               { return actionGroupListActions(a) }
func (a *ActionGroup) RemoveAction(action *Action)         { actionGroupRemoveAction(a, action) }
func (a *ActionGroup) SetSensitive(sensitive bool)         { actionGroupSetSensitive(a, sensitive) }
func (a *ActionGroup) SetTranslateFunc(f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify) {
	actionGroupSetTranslateFunc(a, f, data, notify)
}
func (a *ActionGroup) SetTranslationDomain(domain string) { actionGroupSetTranslationDomain(a, domain) }
func (a *ActionGroup) SetVisible(visible bool)            { actionGroupSetVisible(a, visible) }
func (a *ActionGroup) TranslateString(str string) string  { return actionGroupTranslateString(a, str) }

type Activatable struct{}

var ActivatableGetType func() O.Type

var (
	activatableDoSetRelatedAction     func(a *Activatable, action *Action)
	activatableGetRelatedAction       func(a *Activatable) *Action
	activatableGetUseActionAppearance func(a *Activatable) bool
	activatableSetRelatedAction       func(a *Activatable, action *Action)
	activatableSetUseActionAppearance func(a *Activatable, useAppearance bool)
	activatableSyncActionProperties   func(a *Activatable, action *Action)
)

func (a *Activatable) DoSetRelatedAction(action *Action) { activatableDoSetRelatedAction(a, action) }
func (a *Activatable) GetRelatedAction() *Action         { return activatableGetRelatedAction(a) }
func (a *Activatable) GetUseActionAppearance() bool      { return activatableGetUseActionAppearance(a) }
func (a *Activatable) SetRelatedAction(action *Action)   { activatableSetRelatedAction(a, action) }
func (a *Activatable) SetUseActionAppearance(useAppearance bool) {
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
	screen *D.Screen) bool

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
	assistantGetPageComplete    func(a *Assistant, page *Widget) bool
	assistantGetPageHeaderImage func(a *Assistant, page *Widget) *D.Pixbuf
	assistantGetPageSideImage   func(a *Assistant, page *Widget) *D.Pixbuf
	assistantGetPageTitle       func(a *Assistant, page *Widget) string
	assistantGetPageType        func(a *Assistant, page *Widget) AssistantPageType
	assistantInsertPage         func(a *Assistant, page *Widget, position int) int
	assistantPrependPage        func(a *Assistant, page *Widget) int
	assistantRemoveActionWidget func(a *Assistant, child *Widget)
	assistantSetCurrentPage     func(a *Assistant, pageNum int)
	assistantSetForwardPageFunc func(a *Assistant, pageFunc AssistantPageFunc, data T.Gpointer, destroy T.GDestroyNotify)
	assistantSetPageComplete    func(a *Assistant, page *Widget, complete bool)
	assistantSetPageHeaderImage func(a *Assistant, page *Widget, pixbuf *D.Pixbuf)
	assistantSetPageSideImage   func(a *Assistant, page *Widget, pixbuf *D.Pixbuf)
	assistantSetPageTitle       func(a *Assistant, page *Widget, title string)
	assistantSetPageType        func(a *Assistant, page *Widget, t AssistantPageType)
	assistantUpdateButtonsState func(a *Assistant)
)

func (a *Assistant) AddActionWidget(child *Widget)     { assistantAddActionWidget(a, child) }
func (a *Assistant) AppendPage(page *Widget) int       { return assistantAppendPage(a, page) }
func (a *Assistant) Commit()                           { assistantCommit(a) }
func (a *Assistant) GetCurrentPage() int               { return assistantGetCurrentPage(a) }
func (a *Assistant) GetNPages() int                    { return assistantGetNPages(a) }
func (a *Assistant) GetNthPage(pageNum int) *Widget    { return assistantGetNthPage(a, pageNum) }
func (a *Assistant) GetPageComplete(page *Widget) bool { return assistantGetPageComplete(a, page) }
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
func (a *Assistant) SetPageComplete(page *Widget, complete bool) {
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
	ObeyChild        bool
	CenterAllocation Allocation
}

var (
	AspectFrameGetType func() O.Type
	AspectFrameNew     func(label string, xalign, yalign, ratio float32, obeyChild bool) *Widget
)

var aspectFrameSet func(a *AspectFrame, xalign, yalign, ratio float32, obeyChild bool)

func (a *AspectFrame) Set(xalign, yalign, ratio float32, obeyChild bool) {
	aspectFrameSet(a, xalign, yalign, ratio, obeyChild)
}

var AttachOptionsGetType func() O.Type

type AttachOptions Enum

const (
	EXPAND AttachOptions = 1 << iota
	SHRINK
	FILL
)
