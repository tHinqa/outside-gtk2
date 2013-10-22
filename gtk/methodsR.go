// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type RadioAction struct {
	Parent ToggleAction
	_      *struct{}
}

var (
	RadioActionGetType func() T.GType
	RadioActionNew     func(name, label, tooltip, stockId string, value int) *RadioAction

	radioActionGetCurrentValue func(r *RadioAction) int
	radioActionGetGroup        func(r *RadioAction) *T.GSList
	radioActionSetCurrentValue func(r *RadioAction, currentValue int)
	radioActionSetGroup        func(r *RadioAction, group *T.GSList)
)

func (r *RadioAction) GetCurrentValue() int             { return radioActionGetCurrentValue(r) }
func (r *RadioAction) GetGroup() *T.GSList              { return radioActionGetGroup(r) }
func (r *RadioAction) SetCurrentValue(currentValue int) { radioActionSetCurrentValue(r, currentValue) }
func (r *RadioAction) SetGroup(group *T.GSList)         { radioActionSetGroup(r, group) }

type RadioButton struct {
	CheckButton T.GtkCheckButton
	Group       *T.GSList
}

var (
	RadioButtonGetType                   func() T.GType
	RadioButtonNew                       func(group *T.GSList) *Widget
	RadioButtonNewFromWidget             func(radioGroupMember *RadioButton) *Widget
	RadioButtonNewWithLabel              func(group *T.GSList, label string) *Widget
	RadioButtonNewWithLabelFromWidget    func(radioGroupMember *RadioButton, label string) *Widget
	RadioButtonNewWithMnemonic           func(group *T.GSList, label string) *Widget
	RadioButtonNewWithMnemonicFromWidget func(radioGroupMember *RadioButton, label string) *Widget

	radioButtonGetGroup func(r *RadioButton) *T.GSList
	radioButtonSetGroup func(r *RadioButton, group *T.GSList)
)

func (r *RadioButton) GetGroup() *T.GSList      { return radioButtonGetGroup(r) }
func (r *RadioButton) SetGroup(group *T.GSList) { radioButtonSetGroup(r, group) }

type RadioMenuItem struct {
	CheckMenuItem CheckMenuItem
	Group         *T.GSList
}

var (
	RadioMenuItemGetType                   func() T.GType
	RadioMenuItemNew                       func(group *T.GSList) *Widget
	RadioMenuItemNewFromWidget             func(group *RadioMenuItem) *Widget
	RadioMenuItemNewWithLabel              func(group *T.GSList, label string) *Widget
	RadioMenuItemNewWithLabelFromWidget    func(group *RadioMenuItem, label string) *Widget
	RadioMenuItemNewWithMnemonic           func(group *T.GSList, label string) *Widget
	RadioMenuItemNewWithMnemonicFromWidget func(group *RadioMenuItem, label string) *Widget

	radioMenuItemGetGroup func(r *RadioMenuItem) *T.GSList
	radioMenuItemSetGroup func(r *RadioMenuItem, group *T.GSList)
)

func (r *RadioMenuItem) GetGroup() *T.GSList      { return radioMenuItemGetGroup(r) }
func (r *RadioMenuItem) SetGroup(group *T.GSList) { radioMenuItemSetGroup(r, group) }

type RadioToolButton struct {
	Parent ToggleToolButton
}

var (
	RadioToolButtonGetType                func() T.GType
	RadioToolButtonNew                    func(group *T.GSList) *ToolItem
	RadioToolButtonNewFromStock           func(group *T.GSList, stockId string) *ToolItem
	RadioToolButtonNewFromWidget          func(group *RadioToolButton) *ToolItem
	RadioToolButtonNewWithStockFromWidget func(group *RadioToolButton, stockId string) *ToolItem

	radioToolButtonGetGroup func(r *RadioToolButton) *T.GSList
	radioToolButtonSetGroup func(r *RadioToolButton, group *T.GSList)
)

func (r *RadioToolButton) GetGroup() *T.GSList      { return radioToolButtonGetGroup(r) }
func (r *RadioToolButton) SetGroup(group *T.GSList) { radioToolButtonSetGroup(r, group) }

type (
	Range struct {
		Widget       Widget
		Adjustment   *Adjustment
		UpdatePolicy T.GtkUpdateType
		Bits         uint
		// Inverted : 1
		// Flippable : 1
		// HasStepperA : 1
		// HasStepperB : 1
		// HasStepperC : 1
		// HasStepperD : 1
		// NeedRecalc : 1
		// SliderSizeFixed : 1
		MinSliderSize int
		Orientation   Orientation
		RangeRect     T.GdkRectangle
		SliderStart   int
		SliderEnd     int
		RoundDigits   int
		Bits2         uint
		// TroughClickForward : 1
		// UpdatePending : 1
		Layout                     *RangeLayout
		Timer                      *RangeStepTimer
		SlideInitialSliderPosition int
		SlideInitialCoordinate     int
		UpdateTimeoutId            uint
		EventWindow                *T.GdkWindow
	}

	RangeLayout struct{}

	RangeStepTimer struct{}
)

var (
	RangeGetType func() T.GType

	rangeGetAdjustment              func(r *Range) *Adjustment
	rangeGetFillLevel               func(r *Range) float64
	rangeGetFlippable               func(r *Range) T.Gboolean
	rangeGetInverted                func(r *Range) T.Gboolean
	rangeGetLowerStepperSensitivity func(r *Range) T.GtkSensitivityType
	rangeGetMinSliderSize           func(r *Range) int
	rangeGetRangeRect               func(r *Range, rangeRect *T.GdkRectangle)
	rangeGetRestrictToFillLevel     func(r *Range) T.Gboolean
	rangeGetRoundDigits             func(r *Range) int
	rangeGetShowFillLevel           func(r *Range) T.Gboolean
	rangeGetSliderRange             func(r *Range, sliderStart, sliderEnd *int)
	rangeGetSliderSizeFixed         func(r *Range) T.Gboolean
	rangeGetUpdatePolicy            func(r *Range) T.GtkUpdateType
	rangeGetUpperStepperSensitivity func(r *Range) T.GtkSensitivityType
	rangeGetValue                   func(r *Range) float64
	rangeSetAdjustment              func(r *Range, adjustment *Adjustment)
	rangeSetFillLevel               func(r *Range, fillLevel float64)
	rangeSetFlippable               func(r *Range, flippable T.Gboolean)
	rangeSetIncrements              func(r *Range, step, page float64)
	rangeSetInverted                func(r *Range, setting T.Gboolean)
	rangeSetLowerStepperSensitivity func(r *Range, sensitivity T.GtkSensitivityType)
	rangeSetMinSliderSize           func(r *Range, minSize T.Gboolean)
	rangeSetRange                   func(r *Range, min, max float64)
	rangeSetRestrictToFillLevel     func(r *Range, restrictToFillLevel T.Gboolean)
	rangeSetRoundDigits             func(r *Range, roundDigits int)
	rangeSetShowFillLevel           func(r *Range, showFillLevel T.Gboolean)
	rangeSetSliderSizeFixed         func(r *Range, sizeFixed T.Gboolean)
	rangeSetUpdatePolicy            func(r *Range, policy T.GtkUpdateType)
	rangeSetUpperStepperSensitivity func(r *Range, sensitivity T.GtkSensitivityType)
	rangeSetValue                   func(r *Range, value float64)
)

func (r *Range) GetAdjustment() *Adjustment { return rangeGetAdjustment(r) }
func (r *Range) GetFillLevel() float64      { return rangeGetFillLevel(r) }
func (r *Range) GetFlippable() T.Gboolean   { return rangeGetFlippable(r) }
func (r *Range) GetInverted() T.Gboolean    { return rangeGetInverted(r) }
func (r *Range) GetLowerStepperSensitivity() T.GtkSensitivityType {
	return rangeGetLowerStepperSensitivity(r)
}
func (r *Range) GetMinSliderSize() int                  { return rangeGetMinSliderSize(r) }
func (r *Range) GetRangeRect(rangeRect *T.GdkRectangle) { rangeGetRangeRect(r, rangeRect) }
func (r *Range) GetRestrictToFillLevel() T.Gboolean     { return rangeGetRestrictToFillLevel(r) }
func (r *Range) GetRoundDigits() int                    { return rangeGetRoundDigits(r) }
func (r *Range) GetShowFillLevel() T.Gboolean           { return rangeGetShowFillLevel(r) }
func (r *Range) GetSliderRange(sliderStart, sliderEnd *int) {
	rangeGetSliderRange(r, sliderStart, sliderEnd)
}
func (r *Range) GetSliderSizeFixed() T.Gboolean   { return rangeGetSliderSizeFixed(r) }
func (r *Range) GetUpdatePolicy() T.GtkUpdateType { return rangeGetUpdatePolicy(r) }
func (r *Range) GetUpperStepperSensitivity() T.GtkSensitivityType {
	return rangeGetUpperStepperSensitivity(r)
}
func (r *Range) GetValue() float64                    { return rangeGetValue(r) }
func (r *Range) SetAdjustment(adjustment *Adjustment) { rangeSetAdjustment(r, adjustment) }
func (r *Range) SetFillLevel(fillLevel float64)       { rangeSetFillLevel(r, fillLevel) }
func (r *Range) SetFlippable(flippable T.Gboolean)    { rangeSetFlippable(r, flippable) }
func (r *Range) SetIncrements(step, page float64)     { rangeSetIncrements(r, step, page) }
func (r *Range) SetInverted(setting T.Gboolean)       { rangeSetInverted(r, setting) }
func (r *Range) SetLowerStepperSensitivity(sensitivity T.GtkSensitivityType) {
	rangeSetLowerStepperSensitivity(r, sensitivity)
}
func (r *Range) SetMinSliderSize(minSize T.Gboolean) { rangeSetMinSliderSize(r, minSize) }
func (r *Range) SetRange(min, max float64)           { rangeSetRange(r, min, max) }
func (r *Range) SetRestrictToFillLevel(restrictToFillLevel T.Gboolean) {
	rangeSetRestrictToFillLevel(r, restrictToFillLevel)
}
func (r *Range) SetRoundDigits(roundDigits int)            { rangeSetRoundDigits(r, roundDigits) }
func (r *Range) SetShowFillLevel(showFillLevel T.Gboolean) { rangeSetShowFillLevel(r, showFillLevel) }
func (r *Range) SetSliderSizeFixed(sizeFixed T.Gboolean)   { rangeSetSliderSizeFixed(r, sizeFixed) }
func (r *Range) SetUpdatePolicy(policy T.GtkUpdateType)    { rangeSetUpdatePolicy(r, policy) }
func (r *Range) SetUpperStepperSensitivity(sensitivity T.GtkSensitivityType) {
	rangeSetUpperStepperSensitivity(r, sensitivity)
}
func (r *Range) SetValue(value float64) { rangeSetValue(r, value) }

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

type (
	Ruler struct {
		Widget       Widget
		BackingStore *T.GdkPixmap
		NonGrExpGc   *T.GdkGC
		Metric       *RulerMetric
		Xsrc         int
		Ysrc         int
		SliderSize   int
		Lower        float64
		Upper        float64
		Position     float64
		MaxSize      float64
	}

	RulerMetric struct {
		MetricName    *T.Gchar
		Abbrev        *T.Gchar
		PixelsPerUnit float64
		RulerScale    [10]float64
		Subdivide     [5]int
	}
)

var (
	RulerGetType func() T.GType

	rulerDrawPos   func(r *Ruler)
	rulerDrawTicks func(r *Ruler)
	rulerGetMetric func(r *Ruler) T.GtkMetricType
	rulerGetRange  func(r *Ruler, lower, upper, position, maxSize *float64)
	rulerSetMetric func(r *Ruler, metric T.GtkMetricType)
	rulerSetRange  func(r *Ruler, lower, upper, position, maxSize float64)
)

func (r *Ruler) DrawPos()                   { rulerDrawPos(r) }
func (r *Ruler) DrawTicks()                 { rulerDrawTicks(r) }
func (r *Ruler) GetMetric() T.GtkMetricType { return rulerGetMetric(r) }
func (r *Ruler) GetRange(lower, upper, position, maxSize *float64) {
	rulerGetRange(r, lower, upper, position, maxSize)
}
func (r *Ruler) SetMetric(metric T.GtkMetricType) { rulerSetMetric(r, metric) }
func (r *Ruler) SetRange(lower, upper, position, maxSize float64) {
	rulerSetRange(r, lower, upper, position, maxSize)
}
