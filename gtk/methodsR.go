// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type RecentAction struct {
	Parent Action
	_      *struct{}
}

var (
	RecentActionGetType       func() T.GType
	RecentActionNew           func(name, label, tooltip, stockId string) *Action
	RecentActionNewForManager func(name, label, tooltip, stockId string, manager *RecentManager) *Action

	recentActionGetShowNumbers func(r *RecentAction) T.Gboolean
	recentActionSetShowNumbers func(r *RecentAction, showNumbers T.Gboolean)
)

func (r *RecentAction) GetShowNumbers() T.Gboolean { return recentActionGetShowNumbers(r) }
func (r *RecentAction) SetShowNumbers(showNumbers T.Gboolean) {
	recentActionSetShowNumbers(r, showNumbers)
}

type RecentChooser struct{}

type RecentSortType T.Enum

const (
	RECENT_SORT_NONE RecentSortType = iota
	RECENT_SORT_MRU
	RECENT_SORT_LRU
	RECENT_SORT_CUSTOM
)

var (
	RecentChooserGetType func() T.GType

	RecentChooserErrorGetType        func() T.GType
	RecentChooserMenuNewForManager   func(manager *RecentManager) *Widget
	RecentChooserWidgetGetType       func() T.GType
	RecentChooserWidgetNew           func() *Widget
	RecentChooserWidgetNewForManager func(manager *RecentManager) *Widget

	RecentChooserDialogGetType       func() T.GType
	RecentChooserDialogNew           func(title string, parent *Window, firstButtonText string, v ...VArg) *Widget
	RecentChooserDialogNewForManager func(title string, parent *Window, manager *RecentManager, firstButtonText string, v ...VArg) *Widget

	RecentChooserErrorQuark func() T.GQuark

	recentChooserAddFilter         func(r *RecentChooser, filter *RecentFilter)
	recentChooserGetCurrentItem    func(r *RecentChooser) *RecentInfo
	recentChooserGetCurrentUri     func(r *RecentChooser) string
	recentChooserGetFilter         func(r *RecentChooser) *RecentFilter
	recentChooserGetItems          func(r *RecentChooser) *T.GList
	recentChooserGetLimit          func(r *RecentChooser) int
	recentChooserGetLocalOnly      func(r *RecentChooser) T.Gboolean
	recentChooserGetSelectMultiple func(r *RecentChooser) T.Gboolean
	recentChooserGetShowIcons      func(r *RecentChooser) T.Gboolean
	recentChooserGetShowNotFound   func(r *RecentChooser) T.Gboolean
	recentChooserGetShowNumbers    func(r *RecentChooser) T.Gboolean
	recentChooserGetShowPrivate    func(r *RecentChooser) T.Gboolean
	recentChooserGetShowTips       func(r *RecentChooser) T.Gboolean
	recentChooserGetSortType       func(r *RecentChooser) RecentSortType
	recentChooserGetUris           func(r *RecentChooser, length *T.Gsize) **T.Gchar
	recentChooserListFilters       func(r *RecentChooser) *T.GSList
	recentChooserRemoveFilter      func(r *RecentChooser, filter *RecentFilter)
	recentChooserSelectAll         func(r *RecentChooser)
	recentChooserSelectUri         func(r *RecentChooser, uri string, err **T.GError) T.Gboolean
	recentChooserSetCurrentUri     func(r *RecentChooser, uri string, err **T.GError) T.Gboolean
	recentChooserSetFilter         func(r *RecentChooser, filter *RecentFilter)
	recentChooserSetLimit          func(r *RecentChooser, limit int)
	recentChooserSetLocalOnly      func(r *RecentChooser, localOnly T.Gboolean)
	recentChooserSetSelectMultiple func(r *RecentChooser, selectMultiple T.Gboolean)
	recentChooserSetShowIcons      func(r *RecentChooser, showIcons T.Gboolean)
	recentChooserSetShowNotFound   func(r *RecentChooser, showNotFound T.Gboolean)
	recentChooserSetShowNumbers    func(r *RecentChooser, showNumbers T.Gboolean)
	recentChooserSetShowPrivate    func(r *RecentChooser, showPrivate T.Gboolean)
	recentChooserSetShowTips       func(r *RecentChooser, showTips T.Gboolean)
	recentChooserSetSortFunc       func(r *RecentChooser, sortFunc RecentSortFunc, sortData T.Gpointer, dataDestroy T.GDestroyNotify)
	recentChooserSetSortType       func(r *RecentChooser, sortType RecentSortType)
	recentChooserUnselectAll       func(r *RecentChooser)
	recentChooserUnselectUri       func(r *RecentChooser, uri string)
)

func (r *RecentChooser) AddFilter(filter *RecentFilter)    { recentChooserAddFilter(r, filter) }
func (r *RecentChooser) GetCurrentItem() *RecentInfo       { return recentChooserGetCurrentItem(r) }
func (r *RecentChooser) GetCurrentUri() string             { return recentChooserGetCurrentUri(r) }
func (r *RecentChooser) GetFilter() *RecentFilter          { return recentChooserGetFilter(r) }
func (r *RecentChooser) GetItems() *T.GList                { return recentChooserGetItems(r) }
func (r *RecentChooser) GetLimit() int                     { return recentChooserGetLimit(r) }
func (r *RecentChooser) GetLocalOnly() T.Gboolean          { return recentChooserGetLocalOnly(r) }
func (r *RecentChooser) GetSelectMultiple() T.Gboolean     { return recentChooserGetSelectMultiple(r) }
func (r *RecentChooser) GetShowIcons() T.Gboolean          { return recentChooserGetShowIcons(r) }
func (r *RecentChooser) GetShowNotFound() T.Gboolean       { return recentChooserGetShowNotFound(r) }
func (r *RecentChooser) GetShowNumbers() T.Gboolean        { return recentChooserGetShowNumbers(r) }
func (r *RecentChooser) GetShowPrivate() T.Gboolean        { return recentChooserGetShowPrivate(r) }
func (r *RecentChooser) GetShowTips() T.Gboolean           { return recentChooserGetShowTips(r) }
func (r *RecentChooser) GetSortType() RecentSortType       { return recentChooserGetSortType(r) }
func (r *RecentChooser) GetUris(length *T.Gsize) **T.Gchar { return recentChooserGetUris(r, length) }
func (r *RecentChooser) ListFilters() *T.GSList            { return recentChooserListFilters(r) }
func (r *RecentChooser) RemoveFilter(filter *RecentFilter) { recentChooserRemoveFilter(r, filter) }
func (r *RecentChooser) SelectAll()                        { recentChooserSelectAll(r) }
func (r *RecentChooser) SelectUri(uri string, err **T.GError) T.Gboolean {
	return recentChooserSelectUri(r, uri, err)
}
func (r *RecentChooser) SetCurrentUri(uri string, err **T.GError) T.Gboolean {
	return recentChooserSetCurrentUri(r, uri, err)
}
func (r *RecentChooser) SetFilter(filter *RecentFilter)    { recentChooserSetFilter(r, filter) }
func (r *RecentChooser) SetLimit(limit int)                { recentChooserSetLimit(r, limit) }
func (r *RecentChooser) SetLocalOnly(localOnly T.Gboolean) { recentChooserSetLocalOnly(r, localOnly) }
func (r *RecentChooser) SetSelectMultiple(selectMultiple T.Gboolean) {
	recentChooserSetSelectMultiple(r, selectMultiple)
}
func (r *RecentChooser) SetShowIcons(showIcons T.Gboolean) { recentChooserSetShowIcons(r, showIcons) }
func (r *RecentChooser) SetShowNotFound(showNotFound T.Gboolean) {
	recentChooserSetShowNotFound(r, showNotFound)
}
func (r *RecentChooser) SetShowNumbers(showNumbers T.Gboolean) {
	recentChooserSetShowNumbers(r, showNumbers)
}
func (r *RecentChooser) SetShowPrivate(showPrivate T.Gboolean) {
	recentChooserSetShowPrivate(r, showPrivate)
}
func (r *RecentChooser) SetShowTips(showTips T.Gboolean) { recentChooserSetShowTips(r, showTips) }
func (r *RecentChooser) SetSortFunc(sortFunc RecentSortFunc, sortData T.Gpointer, dataDestroy T.GDestroyNotify) {
	recentChooserSetSortFunc(r, sortFunc, sortData, dataDestroy)
}
func (r *RecentChooser) SetSortType(sortType RecentSortType) { recentChooserSetSortType(r, sortType) }
func (r *RecentChooser) UnselectAll()                        { recentChooserUnselectAll(r) }
func (r *RecentChooser) UnselectUri(uri string)              { recentChooserUnselectUri(r, uri) }

type RecentChooserMenu struct {
	Parent Menu
	_      *struct{}
}

var (
	RecentChooserMenuGetType func() T.GType
	RecentChooserMenuNew     func() *Widget

	recentChooserMenuGetShowNumbers func(r *RecentChooserMenu) T.Gboolean
	recentChooserMenuSetShowNumbers func(r *RecentChooserMenu, showNumbers T.Gboolean)
)

func (r *RecentChooserMenu) GetShowNumbers() T.Gboolean { return recentChooserMenuGetShowNumbers(r) }
func (r *RecentChooserMenu) SetShowNumbers(showNumbers T.Gboolean) {
	recentChooserMenuSetShowNumbers(r, showNumbers)
}

type (
	RecentInfo struct{}

	RecentSortFunc func(
		a, b *RecentInfo, userData T.Gpointer) int
)

var (
	RecentInfoGetType func() T.GType

	recentInfoExists             func(r *RecentInfo) T.Gboolean
	recentInfoGetAdded           func(r *RecentInfo) T.TimeT
	recentInfoGetAge             func(r *RecentInfo) int
	recentInfoGetApplicationInfo func(r *RecentInfo, appName string, appExec **T.Gchar, count *uint, time *T.TimeT) T.Gboolean
	recentInfoGetApplications    func(r *RecentInfo, length *T.Gsize) **T.Gchar
	recentInfoGetDescription     func(r *RecentInfo) string
	recentInfoGetDisplayName     func(r *RecentInfo) string
	recentInfoGetGroups          func(r *RecentInfo, length *T.Gsize) **T.Gchar
	recentInfoGetIcon            func(r *RecentInfo, size int) *T.GdkPixbuf
	recentInfoGetMimeType        func(r *RecentInfo) string
	recentInfoGetModified        func(r *RecentInfo) T.TimeT
	recentInfoGetPrivateHint     func(r *RecentInfo) T.Gboolean
	recentInfoGetShortName       func(r *RecentInfo) string
	recentInfoGetUri             func(r *RecentInfo) string
	recentInfoGetUriDisplay      func(r *RecentInfo) string
	recentInfoGetVisited         func(r *RecentInfo) T.TimeT
	recentInfoHasApplication     func(r *RecentInfo, appName string) T.Gboolean
	recentInfoHasGroup           func(r *RecentInfo, groupName string) T.Gboolean
	recentInfoIsLocal            func(r *RecentInfo) T.Gboolean
	recentInfoLastApplication    func(r *RecentInfo) string
	recentInfoMatch              func(r *RecentInfo, infoB *RecentInfo) T.Gboolean
	recentInfoRef                func(r *RecentInfo) *RecentInfo
	recentInfoUnref              func(r *RecentInfo)
)

func (r *RecentInfo) Exists() T.Gboolean { return recentInfoExists(r) }
func (r *RecentInfo) GetAdded() T.TimeT  { return recentInfoGetAdded(r) }
func (r *RecentInfo) GetAge() int        { return recentInfoGetAge(r) }
func (r *RecentInfo) GetApplicationInfo(appName string, appExec **T.Gchar, count *uint, time *T.TimeT) T.Gboolean {
	return recentInfoGetApplicationInfo(r, appName, appExec, count, time)
}
func (r *RecentInfo) GetApplications(length *T.Gsize) **T.Gchar {
	return recentInfoGetApplications(r, length)
}
func (r *RecentInfo) GetDescription() string              { return recentInfoGetDescription(r) }
func (r *RecentInfo) GetDisplayName() string              { return recentInfoGetDisplayName(r) }
func (r *RecentInfo) GetGroups(length *T.Gsize) **T.Gchar { return recentInfoGetGroups(r, length) }
func (r *RecentInfo) GetIcon(size int) *T.GdkPixbuf       { return recentInfoGetIcon(r, size) }
func (r *RecentInfo) GetMimeType() string                 { return recentInfoGetMimeType(r) }
func (r *RecentInfo) GetModified() T.TimeT                { return recentInfoGetModified(r) }
func (r *RecentInfo) GetPrivateHint() T.Gboolean          { return recentInfoGetPrivateHint(r) }
func (r *RecentInfo) GetShortName() string                { return recentInfoGetShortName(r) }
func (r *RecentInfo) GetUri() string                      { return recentInfoGetUri(r) }
func (r *RecentInfo) GetUriDisplay() string               { return recentInfoGetUriDisplay(r) }
func (r *RecentInfo) GetVisited() T.TimeT                 { return recentInfoGetVisited(r) }
func (r *RecentInfo) HasApplication(appName string) T.Gboolean {
	return recentInfoHasApplication(r, appName)
}
func (r *RecentInfo) HasGroup(groupName string) T.Gboolean { return recentInfoHasGroup(r, groupName) }
func (r *RecentInfo) IsLocal() T.Gboolean                  { return recentInfoIsLocal(r) }
func (r *RecentInfo) LastApplication() string              { return recentInfoLastApplication(r) }
func (r *RecentInfo) Match(infoB *RecentInfo) T.Gboolean   { return recentInfoMatch(r, infoB) }
func (r *RecentInfo) Ref() *RecentInfo                     { return recentInfoRef(r) }
func (r *RecentInfo) Unref()                               { recentInfoUnref(r) }

type (
	RecentFilter struct{}

	RecentFilterFunc func(
		filterInfo *RecentFilterInfo,
		userData T.Gpointer) T.Gboolean

	RecentFilterInfo struct {
		Contains     RecentFilterFlags
		Uri          *T.Gchar
		Display_name *T.Gchar
		Mime_type    *T.Gchar
		Applications **T.Gchar
		Groups       **T.Gchar
		Age          int
	}
)

type RecentFilterFlags T.Enum

const (
	RECENT_FILTER_URI RecentFilterFlags = 1 << iota
	RECENT_FILTER_DISPLAY_NAME
	RECENT_FILTER_MIME_TYPE
	RECENT_FILTER_APPLICATION
	RECENT_FILTER_GROUP
	RECENT_FILTER_AGE
)

var (
	RecentFilterGetType func() T.GType
	RecentFilterNew     func() *RecentFilter

	RecentFilterFlagsGetType func() T.GType

	recentFilterAddAge           func(r *RecentFilter, days int)
	recentFilterAddApplication   func(r *RecentFilter, application string)
	recentFilterAddCustom        func(r *RecentFilter, needed RecentFilterFlags, f RecentFilterFunc, data T.Gpointer, dataDestroy T.GDestroyNotify)
	recentFilterAddGroup         func(r *RecentFilter, group string)
	recentFilterAddMimeType      func(r *RecentFilter, mimeType string)
	recentFilterAddPattern       func(r *RecentFilter, pattern string)
	recentFilterAddPixbufFormats func(r *RecentFilter)
	recentFilterFilter           func(r *RecentFilter, filterInfo *RecentFilterInfo) T.Gboolean
	recentFilterGetName          func(r *RecentFilter) string
	recentFilterGetNeeded        func(r *RecentFilter) RecentFilterFlags
	recentFilterSetName          func(r *RecentFilter, name string)
)

func (r *RecentFilter) AddAge(days int)                   { recentFilterAddAge(r, days) }
func (r *RecentFilter) AddApplication(application string) { recentFilterAddApplication(r, application) }
func (r *RecentFilter) AddCustom(needed RecentFilterFlags, f RecentFilterFunc, data T.Gpointer, dataDestroy T.GDestroyNotify) {
	recentFilterAddCustom(r, needed, f, data, dataDestroy)
}
func (r *RecentFilter) AddGroup(group string)       { recentFilterAddGroup(r, group) }
func (r *RecentFilter) AddMimeType(mimeType string) { recentFilterAddMimeType(r, mimeType) }
func (r *RecentFilter) AddPattern(pattern string)   { recentFilterAddPattern(r, pattern) }
func (r *RecentFilter) AddPixbufFormats()           { recentFilterAddPixbufFormats(r) }
func (r *RecentFilter) Filter(filterInfo *RecentFilterInfo) T.Gboolean {
	return recentFilterFilter(r, filterInfo)
}
func (r *RecentFilter) GetName() string              { return recentFilterGetName(r) }
func (r *RecentFilter) GetNeeded() RecentFilterFlags { return recentFilterGetNeeded(r) }
func (r *RecentFilter) SetName(name string)          { recentFilterSetName(r, name) }

type RecentManager struct {
	Parent T.GObject
	_      *struct{}
}

type RecentData struct {
	DisplayName *T.Gchar
	Description *T.Gchar
	MimeType    *T.Gchar
	AppName     *T.Gchar
	AppExec     *T.Gchar
	Groups      **T.Gchar
	IsPrivate   T.Gboolean
}

var (
	RecentManagerGetType func() T.GType
	RecentManagerNew     func() *RecentManager

	RecentManagerErrorGetType func() T.GType
	RecentManagerErrorQuark   func() T.GQuark

	RecentManagerGetDefault   func() *RecentManager
	RecentManagerGetForScreen func(screen *T.GdkScreen) *RecentManager

	recentManagerAddFull    func(r *RecentManager, uri string, recentData *RecentData) T.Gboolean
	recentManagerAddItem    func(r *RecentManager, uri string) T.Gboolean
	recentManagerGetItems   func(r *RecentManager) *T.GList
	recentManagerGetLimit   func(r *RecentManager) int
	recentManagerHasItem    func(r *RecentManager, uri string) T.Gboolean
	recentManagerLookupItem func(r *RecentManager, uri string, err **T.GError) *RecentInfo
	recentManagerMoveItem   func(r *RecentManager, uri string, newUri string, err **T.GError) T.Gboolean
	recentManagerPurgeItems func(r *RecentManager, err **T.GError) int
	recentManagerRemoveItem func(r *RecentManager, uri string, err **T.GError) T.Gboolean
	recentManagerSetLimit   func(r *RecentManager, limit int)
	recentManagerSetScreen  func(r *RecentManager, screen *T.GdkScreen)
)

func (r *RecentManager) AddFull(uri string, recentData *RecentData) T.Gboolean {
	return recentManagerAddFull(r, uri, recentData)
}
func (r *RecentManager) AddItem(uri string) T.Gboolean { return recentManagerAddItem(r, uri) }
func (r *RecentManager) GetItems() *T.GList            { return recentManagerGetItems(r) }
func (r *RecentManager) GetLimit() int                 { return recentManagerGetLimit(r) }
func (r *RecentManager) HasItem(uri string) T.Gboolean { return recentManagerHasItem(r, uri) }
func (r *RecentManager) LookupItem(uri string, err **T.GError) *RecentInfo {
	return recentManagerLookupItem(r, uri, err)
}
func (r *RecentManager) MoveItem(uri string, newUri string, err **T.GError) T.Gboolean {
	return recentManagerMoveItem(r, uri, newUri, err)
}
func (r *RecentManager) PurgeItems(err **T.GError) int { return recentManagerPurgeItems(r, err) }
func (r *RecentManager) RemoveItem(uri string, err **T.GError) T.Gboolean {
	return recentManagerRemoveItem(r, uri, err)
}
func (r *RecentManager) SetLimit(limit int)            { recentManagerSetLimit(r, limit) }
func (r *RecentManager) SetScreen(screen *T.GdkScreen) { recentManagerSetScreen(r, screen) }
