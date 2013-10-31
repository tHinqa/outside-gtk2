// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	A "github.com/tHinqa/outside-gtk2/atk"
	D "github.com/tHinqa/outside-gtk2/gdk"
	I "github.com/tHinqa/outside-gtk2/gio"
	L "github.com/tHinqa/outside-gtk2/glib"
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

	AboutDialogGetArtists      func(a *AboutDialog) []string  // 2.6
	AboutDialogGetAuthors      func(a *AboutDialog) []string  // 2.6
	AboutDialogGetComments     func(a *AboutDialog) string    // 2.6
	AboutDialogGetCopyright    func(a *AboutDialog) string    // 2.6
	AboutDialogGetDocumenters  func(a *AboutDialog) []string  // 2.6
	AboutDialogGetLicense      func(a *AboutDialog) string    // 2.6
	AboutDialogGetLogo         func(a *AboutDialog) *D.Pixbuf // 2.6
	AboutDialogGetLogoIconName func(a *AboutDialog) string    // 2.6
	// 2.6 Deprecated 2.12 aboutDialogGetName func(a *AboutDialog) string
	AboutDialogGetProgramName       func(a *AboutDialog) string                // 2.12
	AboutDialogGetTranslatorCredits func(a *AboutDialog) string                // 2.6
	AboutDialogGetVersion           func(a *AboutDialog) string                // 2.6
	AboutDialogGetWebsite           func(a *AboutDialog) string                // 2.6
	AboutDialogGetWebsiteLabel      func(a *AboutDialog) string                // 2.6
	AboutDialogGetWrapLicense       func(a *AboutDialog) bool                  // 2.8
	AboutDialogSetArtists           func(a *AboutDialog, artists []string)     // 2.6
	AboutDialogSetAuthors           func(a *AboutDialog, authors []string)     // 2.6
	AboutDialogSetComments          func(a *AboutDialog, comments string)      // 2.6
	AboutDialogSetCopyright         func(a *AboutDialog, copyright string)     // 2.6
	AboutDialogSetDocumenters       func(a *AboutDialog, documenters []string) // 2.6
	AboutDialogSetLicense           func(a *AboutDialog, license string)       // 2.6
	AboutDialogSetLogo              func(a *AboutDialog, logo *D.Pixbuf)       // 2.6
	AboutDialogSetLogoIconName      func(a *AboutDialog, iconName string)      // 2.6
	// 2.6 Deprecated 2.12 aboutDialogSetName func(a *AboutDialog, name string)
	AboutDialogSetProgramName       func(a *AboutDialog, name string)              // 2.12
	AboutDialogSetTranslatorCredits func(a *AboutDialog, translatorCredits string) // 2.6
	AboutDialogSetVersion           func(a *AboutDialog, version string)           // 2.6
	AboutDialogSetWebsite           func(a *AboutDialog, website string)           // 2.6
	AboutDialogSetWebsiteLabel      func(a *AboutDialog, websiteLabel string)      // 2.6
	AboutDialogSetWrapLicense       func(a *AboutDialog, wrapLicense bool)         // 2.8
)

func NewAboutDialog() (w *Widget, a *AboutDialog) {
	w = AboutDialogNew()
	a = (*AboutDialog)(unsafe.Pointer(w))
	return
}

func (a *AboutDialog) Artists() []string                   { return AboutDialogGetArtists(a) }
func (a *AboutDialog) Authors() []string                   { return AboutDialogGetAuthors(a) }
func (a *AboutDialog) Comments() string                    { return AboutDialogGetComments(a) }
func (a *AboutDialog) Copyright() string                   { return AboutDialogGetCopyright(a) }
func (a *AboutDialog) Documenters() []string               { return AboutDialogGetDocumenters(a) }
func (a *AboutDialog) License() string                     { return AboutDialogGetLicense(a) }
func (a *AboutDialog) Logo() *D.Pixbuf                     { return AboutDialogGetLogo(a) }
func (a *AboutDialog) LogoIconName() string                { return AboutDialogGetLogoIconName(a) }
func (a *AboutDialog) ProgramName() string                 { return AboutDialogGetProgramName(a) }
func (a *AboutDialog) TranslatorCredits() string           { return AboutDialogGetTranslatorCredits(a) }
func (a *AboutDialog) Version() string                     { return AboutDialogGetVersion(a) }
func (a *AboutDialog) Website() string                     { return AboutDialogGetWebsite(a) }
func (a *AboutDialog) WebsiteLabel() string                { return AboutDialogGetWebsiteLabel(a) }
func (a *AboutDialog) WrapLicense() bool                   { return AboutDialogGetWrapLicense(a) }
func (a *AboutDialog) SetArtists(artists []string)         { AboutDialogSetArtists(a, artists) }
func (a *AboutDialog) SetAuthors(authors []string)         { AboutDialogSetAuthors(a, authors) }
func (a *AboutDialog) SetComments(comments string)         { AboutDialogSetComments(a, comments) }
func (a *AboutDialog) SetCopyright(copyright string)       { AboutDialogSetCopyright(a, copyright) }
func (a *AboutDialog) SetDocumenters(documenters []string) { AboutDialogSetDocumenters(a, documenters) }
func (a *AboutDialog) SetLicense(license string)           { AboutDialogSetLicense(a, license) }
func (a *AboutDialog) SetLogo(logo *D.Pixbuf)              { AboutDialogSetLogo(a, logo) }
func (a *AboutDialog) SetLogoIconName(iconName string)     { AboutDialogSetLogoIconName(a, iconName) }

// Deprecated 2.12 func (a *AboutDialog) SetName(name string)                  { AboutDialogSetName(a, name) }
func (a *AboutDialog) SetProgramName(name string) { AboutDialogSetProgramName(a, name) }
func (a *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	AboutDialogSetTranslatorCredits(a, translatorCredits)
}
func (a *AboutDialog) SetVersion(version string) { AboutDialogSetVersion(a, version) }
func (a *AboutDialog) SetWebsite(website string) { AboutDialogSetWebsite(a, website) }
func (a *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	AboutDialogSetWebsiteLabel(a, websiteLabel)
}
func (a *AboutDialog) SetWrapLicense(wrapLicense bool) {
	AboutDialogSetWrapLicense(a, wrapLicense)
}

// Deprecated 2.12 func (a *AboutDialog) GetName() string { return AboutDialogGetName(a) }

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

	AccelGroupsFromObject func(object *T.GObject) *L.SList
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
		Acceleratables *L.SList
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
	AccelMapLoadScanner       func(scanner *L.Scanner)
	AccelMapLockPath          func(accelPath string)
	AccelMapLookupEntry       func(accelPath string, key *AccelKey) bool
	AccelMapSave              func(fileName string)
	AccelMapSaveFd            func(fd int)
	AccelMapUnlockPath        func(accelPath string)

	AccelFlagsGetType func() O.Type
)

var (
	AccelGroupActivate        func(a *AccelGroup, accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) bool
	AccelGroupConnect         func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags, closure *O.Closure)
	AccelGroupConnectByPath   func(a *AccelGroup, accelPath string, closure *O.Closure)
	AccelGroupDisconnect      func(a *AccelGroup, closure *O.Closure) bool
	AccelGroupDisconnectKey   func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType) bool
	AccelGroupFind            func(a *AccelGroup, findFunc AccelGroupFindFunc, data T.Gpointer) *AccelKey
	AccelGroupGetIsLocked     func(a *AccelGroup) bool
	AccelGroupGetModifierMask func(a *AccelGroup) T.GdkModifierType
	AccelGroupLock            func(a *AccelGroup)
	AccelGroupQuery           func(a *AccelGroup, accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *AccelGroupEntry
	AccelGroupUnlock          func(a *AccelGroup)

	AccelLabelGetAccelWidget  func(a *AccelLabel) *Widget
	AccelLabelGetAccelWidth   func(a *AccelLabel) uint
	AccelLabelRefetch         func(a *AccelLabel) bool
	AccelLabelSetAccelClosure func(a *AccelLabel, accelClosure *O.Closure)
	AccelLabelSetAccelWidget  func(a *AccelLabel, accelWidget *Widget)
)

func (a *AccelGroup) Activate(accelQuark T.GQuark, acceleratable *T.GObject, accelKey uint, accelMods T.GdkModifierType) bool {
	return AccelGroupActivate(a, accelQuark, acceleratable, accelKey, accelMods)
}
func (a *AccelGroup) Connect(accelKey uint, accelMods T.GdkModifierType, accelFlags AccelFlags, closure *O.Closure) {
	AccelGroupConnect(a, accelKey, accelMods, accelFlags, closure)
}
func (a *AccelGroup) ConnectByPath(accelPath string, closure *O.Closure) {
	AccelGroupConnectByPath(a, accelPath, closure)
}
func (a *AccelGroup) Disconnect(closure *O.Closure) bool {
	return AccelGroupDisconnect(a, closure)
}
func (a *AccelGroup) DisconnectKey(accelKey uint, accelMods T.GdkModifierType) bool {
	return AccelGroupDisconnectKey(a, accelKey, accelMods)
}
func (a *AccelGroup) Find(findFunc AccelGroupFindFunc, data T.Gpointer) *AccelKey {
	return AccelGroupFind(a, findFunc, data)
}
func (a *AccelGroup) GetIsLocked() bool                  { return AccelGroupGetIsLocked(a) }
func (a *AccelGroup) GetModifierMask() T.GdkModifierType { return AccelGroupGetModifierMask(a) }
func (a *AccelGroup) Lock()                              { AccelGroupLock(a) }
func (a *AccelGroup) Query(accelKey uint, accelMods T.GdkModifierType, nEntries *uint) *AccelGroupEntry {
	return AccelGroupQuery(a, accelKey, accelMods, nEntries)
}
func (a *AccelGroup) Unlock() { AccelGroupUnlock(a) }

func (a *AccelLabel) GetAccelWidget() *Widget { return AccelLabelGetAccelWidget(a) }
func (a *AccelLabel) GetAccelWidth() uint     { return AccelLabelGetAccelWidth(a) }
func (a *AccelLabel) Refetch() bool           { return AccelLabelRefetch(a) }
func (a *AccelLabel) SetAccelClosure(accelClosure *O.Closure) {
	AccelLabelSetAccelClosure(a, accelClosure)
}
func (a *AccelLabel) SetAccelWidget(accelWidget *Widget) { AccelLabelSetAccelWidget(a, accelWidget) }

type Accessible struct {
	Parent A.Object
	Widget *Widget
}

var AccessibleGetType func() O.Type

var (
	AccessibleConnectWidgetDestroyed func(a *Accessible)
	AccessibleGetWidget              func(a *Accessible) *Widget
	AccessibleSetWidget              func(a *Accessible, widget *Widget)
)

func (a *Accessible) SetWidget(widget *Widget) { AccessibleSetWidget(a, widget) }
func (a *Accessible) GetWidget() *Widget       { return AccessibleGetWidget(a) }
func (a *Accessible) ConnectWidgetDestroyed()  { AccessibleConnectWidgetDestroyed(a) }

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
	ActionActivate              func(a *Action)
	ActionBlockActivate         func(a *Action)
	ActionBlockActivateFrom     func(a *Action, proxy *Widget)
	ActionConnectAccelerator    func(a *Action)
	ActionConnectProxy          func(a *Action, proxy *Widget)
	ActionCreateIcon            func(a *Action, iconSize IconSize) *Widget
	ActionCreateMenu            func(a *Action) *Widget
	ActionCreateMenuItem        func(a *Action) *Widget
	ActionCreateToolItem        func(a *Action) *Widget
	ActionDisconnectAccelerator func(a *Action)
	ActionDisconnectProxy       func(a *Action, proxy *Widget)
	ActionGetAccelClosure       func(a *Action) *O.Closure
	ActionGetAccelPath          func(a *Action) string
	ActionGetAlwaysShowImage    func(a *Action) bool
	ActionGetGicon              func(a *Action) *I.Icon
	ActionGetIconName           func(a *Action) string
	ActionGetIsImportant        func(a *Action) bool
	ActionGetLabel              func(a *Action) string
	ActionGetName               func(a *Action) string
	ActionGetProxies            func(a *Action) *L.SList
	ActionGetSensitive          func(a *Action) bool
	ActionGetShortLabel         func(a *Action) string
	ActionGetStockId            func(a *Action) string
	ActionGetTooltip            func(a *Action) string
	ActionGetVisible            func(a *Action) bool
	ActionGetVisibleHorizontal  func(a *Action) bool
	ActionGetVisibleVertical    func(a *Action) bool
	ActionIsSensitive           func(a *Action) bool
	ActionIsVisible             func(a *Action) bool
	ActionSetAccelGroup         func(a *Action, accelGroup *AccelGroup)
	ActionSetAccelPath          func(a *Action, accelPath string)
	ActionSetAlwaysShowImage    func(a *Action, alwaysShow bool)
	ActionSetGicon              func(a *Action, icon *I.Icon)
	ActionSetIconName           func(a *Action, iconName string)
	ActionSetIsImportant        func(a *Action, isImportant bool)
	ActionSetLabel              func(a *Action, label string)
	ActionSetSensitive          func(a *Action, sensitive bool)
	ActionSetShortLabel         func(a *Action, shortLabel string)
	ActionSetStockId            func(a *Action, stockId string)
	ActionSetTooltip            func(a *Action, tooltip string)
	ActionSetVisible            func(a *Action, visible bool)
	ActionSetVisibleHorizontal  func(a *Action, visibleHorizontal bool)
	ActionSetVisibleVertical    func(a *Action, visibleVertical bool)
	ActionUnblockActivate       func(a *Action)
	ActionUnblockActivateFrom   func(a *Action, proxy *Widget)
)

func (a *Action) Activate()                            { ActionActivate(a) }
func (a *Action) BlockActivate()                       { ActionBlockActivate(a) }
func (a *Action) BlockActivateFrom(proxy *Widget)      { ActionBlockActivateFrom(a, proxy) }
func (a *Action) ConnectAccelerator()                  { ActionConnectAccelerator(a) }
func (a *Action) ConnectProxy(proxy *Widget)           { ActionConnectProxy(a, proxy) }
func (a *Action) CreateIcon(iconSize IconSize) *Widget { return ActionCreateIcon(a, iconSize) }
func (a *Action) CreateMenu() *Widget                  { return ActionCreateMenu(a) }
func (a *Action) CreateMenuItem() *Widget              { return ActionCreateMenuItem(a) }
func (a *Action) CreateToolItem() *Widget              { return ActionCreateToolItem(a) }
func (a *Action) DisconnectAccelerator()               { ActionDisconnectAccelerator(a) }
func (a *Action) DisconnectProxy(proxy *Widget)        { ActionDisconnectProxy(a, proxy) }
func (a *Action) GetAccelClosure() *O.Closure          { return ActionGetAccelClosure(a) }
func (a *Action) GetAccelPath() string                 { return ActionGetAccelPath(a) }
func (a *Action) GetAlwaysShowImage() bool             { return ActionGetAlwaysShowImage(a) }
func (a *Action) GetGicon() *I.Icon                    { return ActionGetGicon(a) }
func (a *Action) GetIconName() string                  { return ActionGetIconName(a) }
func (a *Action) GetIsImportant() bool                 { return ActionGetIsImportant(a) }
func (a *Action) GetLabel() string                     { return ActionGetLabel(a) }
func (a *Action) GetName() string                      { return ActionGetName(a) }
func (a *Action) GetProxies() *L.SList                 { return ActionGetProxies(a) }
func (a *Action) GetSensitive() bool                   { return ActionGetSensitive(a) }
func (a *Action) GetShortLabel() string                { return ActionGetShortLabel(a) }
func (a *Action) GetStockId() string                   { return ActionGetStockId(a) }
func (a *Action) GetTooltip() string                   { return ActionGetTooltip(a) }
func (a *Action) GetVisible() bool                     { return ActionGetVisible(a) }
func (a *Action) GetVisibleHorizontal() bool           { return ActionGetVisibleHorizontal(a) }
func (a *Action) GetVisibleVertical() bool             { return ActionGetVisibleVertical(a) }
func (a *Action) IsSensitive() bool                    { return ActionIsSensitive(a) }
func (a *Action) IsVisible() bool                      { return ActionIsVisible(a) }
func (a *Action) SetAccelGroup(accelGroup *AccelGroup) { ActionSetAccelGroup(a, accelGroup) }
func (a *Action) SetAccelPath(accelPath string)        { ActionSetAccelPath(a, accelPath) }
func (a *Action) SetAlwaysShowImage(alwaysShow bool)   { ActionSetAlwaysShowImage(a, alwaysShow) }
func (a *Action) SetGicon(icon *I.Icon)                { ActionSetGicon(a, icon) }
func (a *Action) SetIconName(iconName string)          { ActionSetIconName(a, iconName) }
func (a *Action) SetIsImportant(isImportant bool)      { ActionSetIsImportant(a, isImportant) }
func (a *Action) SetLabel(label string)                { ActionSetLabel(a, label) }
func (a *Action) SetSensitive(sensitive bool)          { ActionSetSensitive(a, sensitive) }
func (a *Action) SetShortLabel(shortLabel string)      { ActionSetShortLabel(a, shortLabel) }
func (a *Action) SetStockId(stockId string)            { ActionSetStockId(a, stockId) }
func (a *Action) SetTooltip(tooltip string)            { ActionSetTooltip(a, tooltip) }
func (a *Action) SetVisible(visible bool)              { ActionSetVisible(a, visible) }
func (a *Action) SetVisibleHorizontal(visibleHorizontal bool) {
	ActionSetVisibleHorizontal(a, visibleHorizontal)
}
func (a *Action) SetVisibleVertical(visibleVertical bool) {
	ActionSetVisibleVertical(a, visibleVertical)
}
func (a *Action) UnblockActivate()                  { ActionUnblockActivate(a) }
func (a *Action) UnblockActivateFrom(proxy *Widget) { ActionUnblockActivateFrom(a, proxy) }

var (
	ActionGroupAddAction            func(a *ActionGroup, action *Action)
	ActionGroupAddActions           func(a *ActionGroup, entries *ActionEntry, nEntries uint, userData T.Gpointer)
	ActionGroupAddActionsFull       func(a *ActionGroup, entries *ActionEntry, nEntries uint, userData T.Gpointer, destroy T.GDestroyNotify)
	ActionGroupAddActionWithAccel   func(a *ActionGroup, action *Action, accelerator string)
	ActionGroupAddRadioActions      func(a *ActionGroup, entries *RadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer)
	ActionGroupAddRadioActionsFull  func(a *ActionGroup, entries *RadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer, destroy T.GDestroyNotify)
	ActionGroupAddToggleActions     func(a *ActionGroup, entries *ToggleActionEntry, nEntries uint, userData T.Gpointer)
	ActionGroupAddToggleActionsFull func(a *ActionGroup, entries *ToggleActionEntry, nEntries uint, userData T.Gpointer, destroy T.GDestroyNotify)
	ActionGroupGetAction            func(a *ActionGroup, actionName string) *Action
	ActionGroupGetName              func(a *ActionGroup) string
	ActionGroupGetSensitive         func(a *ActionGroup) bool
	ActionGroupGetVisible           func(a *ActionGroup) bool
	ActionGroupListActions          func(a *ActionGroup) *T.GList
	ActionGroupRemoveAction         func(a *ActionGroup, action *Action)
	ActionGroupSetSensitive         func(a *ActionGroup, sensitive bool)
	ActionGroupSetTranslateFunc     func(a *ActionGroup, f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify)
	ActionGroupSetTranslationDomain func(a *ActionGroup, domain string)
	ActionGroupSetVisible           func(a *ActionGroup, visible bool)
	ActionGroupTranslateString      func(a *ActionGroup, str string) string
)

func (a *ActionGroup) AddAction(action *Action) { ActionGroupAddAction(a, action) }
func (a *ActionGroup) AddActions(entries *ActionEntry, nEntries uint, userData T.Gpointer) {
	ActionGroupAddActions(a, entries, nEntries, userData)
}
func (a *ActionGroup) AddActionsFull(entries *ActionEntry, nEntries uint, userData T.Gpointer, destroy T.GDestroyNotify) {
	ActionGroupAddActionsFull(a, entries, nEntries, userData, destroy)
}
func (a *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	ActionGroupAddActionWithAccel(a, action, accelerator)
}
func (a *ActionGroup) AddRadioActions(entries *RadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer) {
	ActionGroupAddRadioActions(a, entries, nEntries, value, onChange, userData)
}
func (a *ActionGroup) AddRadioActionsFull(entries *RadioActionEntry, nEntries uint, value int, onChange T.GCallback, userData T.Gpointer, destroy T.GDestroyNotify) {
	ActionGroupAddRadioActionsFull(a, entries, nEntries, value, onChange, userData, destroy)
}
func (a *ActionGroup) AddToggleActions(entries *ToggleActionEntry, nEntries uint, userData T.Gpointer) {
	ActionGroupAddToggleActions(a, entries, nEntries, userData)
}
func (a *ActionGroup) AddToggleActionsFull(entries *ToggleActionEntry, nEntries uint, userData T.Gpointer, destroy T.GDestroyNotify) {
	ActionGroupAddToggleActionsFull(a, entries, nEntries, userData, destroy)
}
func (a *ActionGroup) GetAction(actionName string) *Action { return ActionGroupGetAction(a, actionName) }
func (a *ActionGroup) GetName() string                     { return ActionGroupGetName(a) }
func (a *ActionGroup) GetSensitive() bool                  { return ActionGroupGetSensitive(a) }
func (a *ActionGroup) GetVisible() bool                    { return ActionGroupGetVisible(a) }
func (a *ActionGroup) ListActions() *T.GList               { return ActionGroupListActions(a) }
func (a *ActionGroup) RemoveAction(action *Action)         { ActionGroupRemoveAction(a, action) }
func (a *ActionGroup) SetSensitive(sensitive bool)         { ActionGroupSetSensitive(a, sensitive) }
func (a *ActionGroup) SetTranslateFunc(f TranslateFunc, data T.Gpointer, notify T.GDestroyNotify) {
	ActionGroupSetTranslateFunc(a, f, data, notify)
}
func (a *ActionGroup) SetTranslationDomain(domain string) { ActionGroupSetTranslationDomain(a, domain) }
func (a *ActionGroup) SetVisible(visible bool)            { ActionGroupSetVisible(a, visible) }
func (a *ActionGroup) TranslateString(str string) string  { return ActionGroupTranslateString(a, str) }

type Activatable struct{}

var ActivatableGetType func() O.Type

var (
	ActivatableDoSetRelatedAction     func(a *Activatable, action *Action)
	ActivatableGetRelatedAction       func(a *Activatable) *Action
	ActivatableGetUseActionAppearance func(a *Activatable) bool
	ActivatableSetRelatedAction       func(a *Activatable, action *Action)
	ActivatableSetUseActionAppearance func(a *Activatable, useAppearance bool)
	ActivatableSyncActionProperties   func(a *Activatable, action *Action)
)

func (a *Activatable) DoSetRelatedAction(action *Action) { ActivatableDoSetRelatedAction(a, action) }
func (a *Activatable) GetRelatedAction() *Action         { return ActivatableGetRelatedAction(a) }
func (a *Activatable) GetUseActionAppearance() bool      { return ActivatableGetUseActionAppearance(a) }
func (a *Activatable) SetRelatedAction(action *Action)   { ActivatableSetRelatedAction(a, action) }
func (a *Activatable) SetUseActionAppearance(useAppearance bool) {
	ActivatableSetUseActionAppearance(a, useAppearance)
}
func (a *Activatable) SyncActionProperties(action *Action) { ActivatableSyncActionProperties(a, action) }

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

func (a *Adjustment) Changed()                       { AdjustmentChanged(a) }
func (a *Adjustment) ClampPage(lower, upper float64) { AdjustmentClampPage(a, lower, upper) }
func (a *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	AdjustmentConfigure(a, value, lower, upper, stepIncrement, pageIncrement, pageSize)
}
func (a *Adjustment) GetLower() float64         { return AdjustmentGetLower(a) }
func (a *Adjustment) GetPageIncrement() float64 { return AdjustmentGetPageIncrement(a) }
func (a *Adjustment) GetPageSize() float64      { return AdjustmentGetPageSize(a) }
func (a *Adjustment) GetStepIncrement() float64 { return AdjustmentGetStepIncrement(a) }
func (a *Adjustment) GetUpper() float64         { return AdjustmentGetUpper(a) }
func (a *Adjustment) GetValue() float64         { return AdjustmentGetValue(a) }
func (a *Adjustment) SetLower(lower float64)    { AdjustmentSetLower(a, lower) }
func (a *Adjustment) SetPageIncrement(pageIncrement float64) {
	AdjustmentSetPageIncrement(a, pageIncrement)
}
func (a *Adjustment) SetPageSize(pageSize float64) { AdjustmentSetPageSize(a, pageSize) }
func (a *Adjustment) SetStepIncrement(stepIncrement float64) {
	AdjustmentSetStepIncrement(a, stepIncrement)
}
func (a *Adjustment) SetUpper(upper float64) { AdjustmentSetUpper(a, upper) }
func (a *Adjustment) SetValue(value float64) { AdjustmentSetValue(a, value) }
func (a *Adjustment) ValueChanged()          { AdjustmentValueChanged(a) }

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

	AlignmentSet        func(a *Alignment, xalign, yalign, xscale, yscale float32)
	AlignmentSetPadding func(a *Alignment, paddingTop, paddingBottom, paddingLeft, paddingRight uint)
	AlignmentGetPadding func(a *Alignment, paddingTop, paddingBottom, paddingLeft, paddingRight *uint)
)

func (a *Alignment) Set(xalign, yalign, xscale, yscale float32) {
	AlignmentSet(a, xalign, yalign, xscale, yscale)
}
func (a *Alignment) SetPadding(paddingTop, paddingBottom, paddingLeft, paddingRight uint) {
	AlignmentSetPadding(a, paddingTop, paddingBottom, paddingLeft, paddingRight)
}
func (a *Alignment) GetPadding(paddingTop, paddingBottom, paddingLeft, paddingRight *uint) {
	AlignmentGetPadding(a, paddingTop, paddingBottom, paddingLeft, paddingRight)
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
	AssistantAddActionWidget    func(a *Assistant, child *Widget)
	AssistantAppendPage         func(a *Assistant, page *Widget) int
	AssistantCommit             func(a *Assistant)
	AssistantGetCurrentPage     func(a *Assistant) int
	AssistantGetNPages          func(a *Assistant) int
	AssistantGetNthPage         func(a *Assistant, pageNum int) *Widget
	AssistantGetPageComplete    func(a *Assistant, page *Widget) bool
	AssistantGetPageHeaderImage func(a *Assistant, page *Widget) *D.Pixbuf
	AssistantGetPageSideImage   func(a *Assistant, page *Widget) *D.Pixbuf
	AssistantGetPageTitle       func(a *Assistant, page *Widget) string
	AssistantGetPageType        func(a *Assistant, page *Widget) AssistantPageType
	AssistantInsertPage         func(a *Assistant, page *Widget, position int) int
	AssistantPrependPage        func(a *Assistant, page *Widget) int
	AssistantRemoveActionWidget func(a *Assistant, child *Widget)
	AssistantSetCurrentPage     func(a *Assistant, pageNum int)
	AssistantSetForwardPageFunc func(a *Assistant, pageFunc AssistantPageFunc, data T.Gpointer, destroy T.GDestroyNotify)
	AssistantSetPageComplete    func(a *Assistant, page *Widget, complete bool)
	AssistantSetPageHeaderImage func(a *Assistant, page *Widget, pixbuf *D.Pixbuf)
	AssistantSetPageSideImage   func(a *Assistant, page *Widget, pixbuf *D.Pixbuf)
	AssistantSetPageTitle       func(a *Assistant, page *Widget, title string)
	AssistantSetPageType        func(a *Assistant, page *Widget, t AssistantPageType)
	AssistantUpdateButtonsState func(a *Assistant)
)

func (a *Assistant) AddActionWidget(child *Widget)     { AssistantAddActionWidget(a, child) }
func (a *Assistant) AppendPage(page *Widget) int       { return AssistantAppendPage(a, page) }
func (a *Assistant) Commit()                           { AssistantCommit(a) }
func (a *Assistant) GetCurrentPage() int               { return AssistantGetCurrentPage(a) }
func (a *Assistant) GetNPages() int                    { return AssistantGetNPages(a) }
func (a *Assistant) GetNthPage(pageNum int) *Widget    { return AssistantGetNthPage(a, pageNum) }
func (a *Assistant) GetPageComplete(page *Widget) bool { return AssistantGetPageComplete(a, page) }
func (a *Assistant) GetPageHeaderImage(page *Widget) *D.Pixbuf {
	return AssistantGetPageHeaderImage(a, page)
}
func (a *Assistant) GetPageSideImage(page *Widget) *D.Pixbuf {
	return AssistantGetPageSideImage(a, page)
}
func (a *Assistant) GetPageTitle(page *Widget) string           { return AssistantGetPageTitle(a, page) }
func (a *Assistant) GetPageType(page *Widget) AssistantPageType { return AssistantGetPageType(a, page) }
func (a *Assistant) InsertPage(page *Widget, position int) int {
	return AssistantInsertPage(a, page, position)
}
func (a *Assistant) PrependPage(page *Widget) int     { return AssistantPrependPage(a, page) }
func (a *Assistant) RemoveActionWidget(child *Widget) { AssistantRemoveActionWidget(a, child) }
func (a *Assistant) SetCurrentPage(pageNum int)       { AssistantSetCurrentPage(a, pageNum) }
func (a *Assistant) SetForwardPageFunc(pageFunc AssistantPageFunc, data T.Gpointer, destroy T.GDestroyNotify) {
	AssistantSetForwardPageFunc(a, pageFunc, data, destroy)
}
func (a *Assistant) SetPageComplete(page *Widget, complete bool) {
	AssistantSetPageComplete(a, page, complete)
}
func (a *Assistant) SetPageHeaderImage(page *Widget, pixbuf *D.Pixbuf) {
	AssistantSetPageHeaderImage(a, page, pixbuf)
}
func (a *Assistant) SetPageSideImage(page *Widget, pixbuf *D.Pixbuf) {
	AssistantSetPageSideImage(a, page, pixbuf)
}
func (a *Assistant) SetPageTitle(page *Widget, title string)       { AssistantSetPageTitle(a, page, title) }
func (a *Assistant) SetPageType(page *Widget, t AssistantPageType) { AssistantSetPageType(a, page, t) }
func (a *Assistant) UpdateButtonsState()                           { AssistantUpdateButtonsState(a) }

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
	ArrowSet func(a *Arrow, arrowType ArrowType, shadowType ShadowType)

	ArrowNew func(a ArrowType, shadowType ShadowType) *Widget
)

func (a *Arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	ArrowSet(a, arrowType, shadowType)
}

func (a ArrowType) New(shadowType ShadowType) *Widget { return ArrowNew(a, shadowType) }

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

var AspectFrameSet func(a *AspectFrame, xalign, yalign, ratio float32, obeyChild bool)

func (a *AspectFrame) Set(xalign, yalign, ratio float32, obeyChild bool) {
	AspectFrameSet(a, xalign, yalign, ratio, obeyChild)
}

var AttachOptionsGetType func() O.Type

type AttachOptions Enum

const (
	EXPAND AttachOptions = 1 << iota
	SHRINK
	FILL
)
