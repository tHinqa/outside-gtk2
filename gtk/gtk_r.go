// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type RadioAction struct {
	Parent ToggleAction
	_      *struct{}
}

var (
	RadioActionGetType func() O.Type
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

type RadioActionEntry struct {
	Name        *T.Gchar
	StockId     *T.Gchar
	Label       *T.Gchar
	Accelerator *T.Gchar
	Tooltip     *T.Gchar
	Value       int
}

type RadioButton struct {
	CheckButton CheckButton
	Group       *T.GSList
}

var (
	RadioButtonGetType                   func() O.Type
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
	RadioMenuItemGetType                   func() O.Type
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
	RadioToolButtonGetType                func() O.Type
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
		UpdatePolicy UpdateType
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
		RangeRect     D.Rectangle
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
		EventWindow                *D.Window
	}

	RangeLayout struct{}

	RangeStepTimer struct{}
)

var (
	RangeGetType func() O.Type

	rangeGetAdjustment              func(r *Range) *Adjustment
	rangeGetFillLevel               func(r *Range) float64
	rangeGetFlippable               func(r *Range) bool
	rangeGetInverted                func(r *Range) bool
	rangeGetLowerStepperSensitivity func(r *Range) SensitivityType
	rangeGetMinSliderSize           func(r *Range) int
	rangeGetRangeRect               func(r *Range, rangeRect *D.Rectangle)
	rangeGetRestrictToFillLevel     func(r *Range) bool
	rangeGetRoundDigits             func(r *Range) int
	rangeGetShowFillLevel           func(r *Range) bool
	rangeGetSliderRange             func(r *Range, sliderStart, sliderEnd *int)
	rangeGetSliderSizeFixed         func(r *Range) bool
	rangeGetUpdatePolicy            func(r *Range) UpdateType
	rangeGetUpperStepperSensitivity func(r *Range) SensitivityType
	rangeGetValue                   func(r *Range) float64
	rangeSetAdjustment              func(r *Range, adjustment *Adjustment)
	rangeSetFillLevel               func(r *Range, fillLevel float64)
	rangeSetFlippable               func(r *Range, flippable bool)
	rangeSetIncrements              func(r *Range, step, page float64)
	rangeSetInverted                func(r *Range, setting bool)
	rangeSetLowerStepperSensitivity func(r *Range, sensitivity SensitivityType)
	rangeSetMinSliderSize           func(r *Range, minSize bool)
	rangeSetRange                   func(r *Range, min, max float64)
	rangeSetRestrictToFillLevel     func(r *Range, restrictToFillLevel bool)
	rangeSetRoundDigits             func(r *Range, roundDigits int)
	rangeSetShowFillLevel           func(r *Range, showFillLevel bool)
	rangeSetSliderSizeFixed         func(r *Range, sizeFixed bool)
	rangeSetUpdatePolicy            func(r *Range, policy UpdateType)
	rangeSetUpperStepperSensitivity func(r *Range, sensitivity SensitivityType)
	rangeSetValue                   func(r *Range, value float64)
)

func (r *Range) GetAdjustment() *Adjustment { return rangeGetAdjustment(r) }
func (r *Range) GetFillLevel() float64      { return rangeGetFillLevel(r) }
func (r *Range) GetFlippable() bool         { return rangeGetFlippable(r) }
func (r *Range) GetInverted() bool          { return rangeGetInverted(r) }
func (r *Range) GetLowerStepperSensitivity() SensitivityType {
	return rangeGetLowerStepperSensitivity(r)
}
func (r *Range) GetMinSliderSize() int               { return rangeGetMinSliderSize(r) }
func (r *Range) GetRangeRect(rangeRect *D.Rectangle) { rangeGetRangeRect(r, rangeRect) }
func (r *Range) GetRestrictToFillLevel() bool        { return rangeGetRestrictToFillLevel(r) }
func (r *Range) GetRoundDigits() int                 { return rangeGetRoundDigits(r) }
func (r *Range) GetShowFillLevel() bool              { return rangeGetShowFillLevel(r) }
func (r *Range) GetSliderRange(sliderStart, sliderEnd *int) {
	rangeGetSliderRange(r, sliderStart, sliderEnd)
}
func (r *Range) GetSliderSizeFixed() bool    { return rangeGetSliderSizeFixed(r) }
func (r *Range) GetUpdatePolicy() UpdateType { return rangeGetUpdatePolicy(r) }
func (r *Range) GetUpperStepperSensitivity() SensitivityType {
	return rangeGetUpperStepperSensitivity(r)
}
func (r *Range) GetValue() float64                    { return rangeGetValue(r) }
func (r *Range) SetAdjustment(adjustment *Adjustment) { rangeSetAdjustment(r, adjustment) }
func (r *Range) SetFillLevel(fillLevel float64)       { rangeSetFillLevel(r, fillLevel) }
func (r *Range) SetFlippable(flippable bool)          { rangeSetFlippable(r, flippable) }
func (r *Range) SetIncrements(step, page float64)     { rangeSetIncrements(r, step, page) }
func (r *Range) SetInverted(setting bool)             { rangeSetInverted(r, setting) }
func (r *Range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
	rangeSetLowerStepperSensitivity(r, sensitivity)
}
func (r *Range) SetMinSliderSize(minSize bool) { rangeSetMinSliderSize(r, minSize) }
func (r *Range) SetRange(min, max float64)     { rangeSetRange(r, min, max) }
func (r *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	rangeSetRestrictToFillLevel(r, restrictToFillLevel)
}
func (r *Range) SetRoundDigits(roundDigits int)      { rangeSetRoundDigits(r, roundDigits) }
func (r *Range) SetShowFillLevel(showFillLevel bool) { rangeSetShowFillLevel(r, showFillLevel) }
func (r *Range) SetSliderSizeFixed(sizeFixed bool)   { rangeSetSliderSizeFixed(r, sizeFixed) }
func (r *Range) SetUpdatePolicy(policy UpdateType)   { rangeSetUpdatePolicy(r, policy) }
func (r *Range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	rangeSetUpperStepperSensitivity(r, sensitivity)
}
func (r *Range) SetValue(value float64) { rangeSetValue(r, value) }

type RcContext struct{}

var (
	RcFlagsGetType     func() O.Type
	RcTokenTypeGetType func() O.Type

	RcAddDefaultFile           func(filename string)
	RcFindModuleInPath         func(moduleFile string) string
	RcFindPixmapInPath         func(settings *Settings, scanner *T.GScanner, pixmapFile string) string
	RcGetDefaultFiles          func() []string
	RcGetImModuleFile          func() string
	RcGetImModulePath          func() string
	RcGetModuleDir             func() string
	RcGetStyle                 func(widget *Widget) *Style
	RcGetStyleByPaths          func(settings *Settings, widgetPath, classPath string, t O.Type) *Style
	RcGetThemeDir              func() string
	RcParse                    func(filename string)
	RcParseString              func(rcString string)
	RcPropertyParseBorder      func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) bool
	RcPropertyParseColor       func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) bool
	RcPropertyParseEnum        func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) bool
	RcPropertyParseFlags       func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) bool
	RcPropertyParseRequisition func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) bool
	RcReparseAll               func() bool
	RcReparseAllForSettings    func(settings *Settings, forceLoad bool) bool
	RcResetStyles              func(settings *Settings)
	RcSetDefaultFiles          func(filenames []string)

	RcScannerNew func() *T.GScanner

	RcParseColor     func(scanner *T.GScanner, color *D.Color) uint
	RcParseColorFull func(scanner *T.GScanner, style *RcStyle, color *D.Color) uint
	RcParseState     func(scanner *T.GScanner, state *StateType) uint
	RcParsePriority  func(scanner *T.GScanner, priority *PathPriorityType) uint
)

type RcPropertyParser func(pspec *T.GParamSpec,
	rcString *T.GString, property_value *T.GValue) bool

type RcStyle struct {
	Parent          T.GObject
	Name            *T.Gchar
	BgPixmapName    *[5]T.Gchar //TODO(t): CHECK
	FontDesc        *P.FontDescription
	ColorFlags      [5]RcFlags
	Fg              [5]D.Color
	Bg              [5]D.Color
	Text            [5]D.Color
	Base            [5]D.Color
	Xthickness      int
	Ythickness      int
	RcProperties    *L.Array
	RcStyleLists    *T.GSList
	IconFactories   *T.GSList
	EngineSpecified uint //: 1
}

type RcFlags Enum

const (
	RC_FG RcFlags = 1 << iota
	RC_BG
	RC_TEXT
	RC_BASE
)

var (
	RcStyleGetType func() O.Type
	RcStyleNew     func() *RcStyle

	RcAddWidgetNameStyle  func(r *RcStyle, pattern string)
	RcAddWidgetClassStyle func(r *RcStyle, pattern string)
	RcAddClassStyle       func(r *RcStyle, pattern string)

	rcStyleCopy  func(r *RcStyle) *RcStyle
	rcStyleRef   func(r *RcStyle)
	rcStyleUnref func(r *RcStyle)
)

func (r *RcStyle) Copy() *RcStyle { return rcStyleCopy(r) }
func (r *RcStyle) Ref()           { rcStyleRef(r) }
func (r *RcStyle) Unref()         { rcStyleUnref(r) }

type RecentAction struct {
	Parent Action
	_      *struct{}
}

var (
	RecentActionGetType       func() O.Type
	RecentActionNew           func(name, label, tooltip, stockId string) *Action
	RecentActionNewForManager func(name, label, tooltip, stockId string, manager *RecentManager) *Action

	recentActionGetShowNumbers func(r *RecentAction) bool
	recentActionSetShowNumbers func(r *RecentAction, showNumbers bool)
)

func (r *RecentAction) GetShowNumbers() bool { return recentActionGetShowNumbers(r) }
func (r *RecentAction) SetShowNumbers(showNumbers bool) {
	recentActionSetShowNumbers(r, showNumbers)
}

type RecentChooser struct{}

type RecentSortType Enum

const (
	RECENT_SORT_NONE RecentSortType = iota
	RECENT_SORT_MRU
	RECENT_SORT_LRU
	RECENT_SORT_CUSTOM
)

var (
	RecentChooserGetType func() O.Type

	RecentChooserErrorGetType        func() O.Type
	RecentChooserMenuNewForManager   func(manager *RecentManager) *Widget
	RecentChooserWidgetGetType       func() O.Type
	RecentChooserWidgetNew           func() *Widget
	RecentChooserWidgetNewForManager func(manager *RecentManager) *Widget

	RecentChooserDialogGetType       func() O.Type
	RecentChooserDialogNew           func(title string, parent *Window, firstButtonText string, v ...VArg) *Widget
	RecentChooserDialogNewForManager func(title string, parent *Window, manager *RecentManager, firstButtonText string, v ...VArg) *Widget

	RecentChooserErrorQuark func() T.GQuark

	recentChooserAddFilter         func(r *RecentChooser, filter *RecentFilter)
	recentChooserGetCurrentItem    func(r *RecentChooser) *RecentInfo
	recentChooserGetCurrentUri     func(r *RecentChooser) string
	recentChooserGetFilter         func(r *RecentChooser) *RecentFilter
	recentChooserGetItems          func(r *RecentChooser) *T.GList
	recentChooserGetLimit          func(r *RecentChooser) int
	recentChooserGetLocalOnly      func(r *RecentChooser) bool
	recentChooserGetSelectMultiple func(r *RecentChooser) bool
	recentChooserGetShowIcons      func(r *RecentChooser) bool
	recentChooserGetShowNotFound   func(r *RecentChooser) bool
	recentChooserGetShowNumbers    func(r *RecentChooser) bool
	recentChooserGetShowPrivate    func(r *RecentChooser) bool
	recentChooserGetShowTips       func(r *RecentChooser) bool
	recentChooserGetSortType       func(r *RecentChooser) RecentSortType
	recentChooserGetUris           func(r *RecentChooser, length *T.Gsize) []string
	recentChooserListFilters       func(r *RecentChooser) *T.GSList
	recentChooserRemoveFilter      func(r *RecentChooser, filter *RecentFilter)
	recentChooserSelectAll         func(r *RecentChooser)
	recentChooserSelectUri         func(r *RecentChooser, uri string, err **T.GError) bool
	recentChooserSetCurrentUri     func(r *RecentChooser, uri string, err **T.GError) bool
	recentChooserSetFilter         func(r *RecentChooser, filter *RecentFilter)
	recentChooserSetLimit          func(r *RecentChooser, limit int)
	recentChooserSetLocalOnly      func(r *RecentChooser, localOnly bool)
	recentChooserSetSelectMultiple func(r *RecentChooser, selectMultiple bool)
	recentChooserSetShowIcons      func(r *RecentChooser, showIcons bool)
	recentChooserSetShowNotFound   func(r *RecentChooser, showNotFound bool)
	recentChooserSetShowNumbers    func(r *RecentChooser, showNumbers bool)
	recentChooserSetShowPrivate    func(r *RecentChooser, showPrivate bool)
	recentChooserSetShowTips       func(r *RecentChooser, showTips bool)
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
func (r *RecentChooser) GetLocalOnly() bool                { return recentChooserGetLocalOnly(r) }
func (r *RecentChooser) GetSelectMultiple() bool           { return recentChooserGetSelectMultiple(r) }
func (r *RecentChooser) GetShowIcons() bool                { return recentChooserGetShowIcons(r) }
func (r *RecentChooser) GetShowNotFound() bool             { return recentChooserGetShowNotFound(r) }
func (r *RecentChooser) GetShowNumbers() bool              { return recentChooserGetShowNumbers(r) }
func (r *RecentChooser) GetShowPrivate() bool              { return recentChooserGetShowPrivate(r) }
func (r *RecentChooser) GetShowTips() bool                 { return recentChooserGetShowTips(r) }
func (r *RecentChooser) GetSortType() RecentSortType       { return recentChooserGetSortType(r) }
func (r *RecentChooser) GetUris(length *T.Gsize) []string  { return recentChooserGetUris(r, length) }
func (r *RecentChooser) ListFilters() *T.GSList            { return recentChooserListFilters(r) }
func (r *RecentChooser) RemoveFilter(filter *RecentFilter) { recentChooserRemoveFilter(r, filter) }
func (r *RecentChooser) SelectAll()                        { recentChooserSelectAll(r) }
func (r *RecentChooser) SelectUri(uri string, err **T.GError) bool {
	return recentChooserSelectUri(r, uri, err)
}
func (r *RecentChooser) SetCurrentUri(uri string, err **T.GError) bool {
	return recentChooserSetCurrentUri(r, uri, err)
}
func (r *RecentChooser) SetFilter(filter *RecentFilter) { recentChooserSetFilter(r, filter) }
func (r *RecentChooser) SetLimit(limit int)             { recentChooserSetLimit(r, limit) }
func (r *RecentChooser) SetLocalOnly(localOnly bool)    { recentChooserSetLocalOnly(r, localOnly) }
func (r *RecentChooser) SetSelectMultiple(selectMultiple bool) {
	recentChooserSetSelectMultiple(r, selectMultiple)
}
func (r *RecentChooser) SetShowIcons(showIcons bool) { recentChooserSetShowIcons(r, showIcons) }
func (r *RecentChooser) SetShowNotFound(showNotFound bool) {
	recentChooserSetShowNotFound(r, showNotFound)
}
func (r *RecentChooser) SetShowNumbers(showNumbers bool) {
	recentChooserSetShowNumbers(r, showNumbers)
}
func (r *RecentChooser) SetShowPrivate(showPrivate bool) {
	recentChooserSetShowPrivate(r, showPrivate)
}
func (r *RecentChooser) SetShowTips(showTips bool) { recentChooserSetShowTips(r, showTips) }
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
	RecentChooserMenuGetType func() O.Type
	RecentChooserMenuNew     func() *Widget

	recentChooserMenuGetShowNumbers func(r *RecentChooserMenu) bool
	recentChooserMenuSetShowNumbers func(r *RecentChooserMenu, showNumbers bool)
)

func (r *RecentChooserMenu) GetShowNumbers() bool { return recentChooserMenuGetShowNumbers(r) }
func (r *RecentChooserMenu) SetShowNumbers(showNumbers bool) {
	recentChooserMenuSetShowNumbers(r, showNumbers)
}

type (
	RecentInfo struct{}

	RecentSortFunc func(
		a, b *RecentInfo, userData T.Gpointer) int
)

var (
	RecentInfoGetType func() O.Type

	recentInfoExists             func(r *RecentInfo) bool
	recentInfoGetAdded           func(r *RecentInfo) T.TimeT
	recentInfoGetAge             func(r *RecentInfo) int
	recentInfoGetApplicationInfo func(r *RecentInfo, appName string, appExec **T.Gchar, count *uint, time *T.TimeT) bool
	recentInfoGetApplications    func(r *RecentInfo, length *T.Gsize) []string
	recentInfoGetDescription     func(r *RecentInfo) string
	recentInfoGetDisplayName     func(r *RecentInfo) string
	recentInfoGetGroups          func(r *RecentInfo, length *T.Gsize) []string
	recentInfoGetIcon            func(r *RecentInfo, size int) *D.Pixbuf
	recentInfoGetMimeType        func(r *RecentInfo) string
	recentInfoGetModified        func(r *RecentInfo) T.TimeT
	recentInfoGetPrivateHint     func(r *RecentInfo) bool
	recentInfoGetShortName       func(r *RecentInfo) string
	recentInfoGetUri             func(r *RecentInfo) string
	recentInfoGetUriDisplay      func(r *RecentInfo) string
	recentInfoGetVisited         func(r *RecentInfo) T.TimeT
	recentInfoHasApplication     func(r *RecentInfo, appName string) bool
	recentInfoHasGroup           func(r *RecentInfo, groupName string) bool
	recentInfoIsLocal            func(r *RecentInfo) bool
	recentInfoLastApplication    func(r *RecentInfo) string
	recentInfoMatch              func(r *RecentInfo, infoB *RecentInfo) bool
	recentInfoRef                func(r *RecentInfo) *RecentInfo
	recentInfoUnref              func(r *RecentInfo)
)

func (r *RecentInfo) Exists() bool      { return recentInfoExists(r) }
func (r *RecentInfo) GetAdded() T.TimeT { return recentInfoGetAdded(r) }
func (r *RecentInfo) GetAge() int       { return recentInfoGetAge(r) }
func (r *RecentInfo) GetApplicationInfo(appName string, appExec **T.Gchar, count *uint, time *T.TimeT) bool {
	return recentInfoGetApplicationInfo(r, appName, appExec, count, time)
}
func (r *RecentInfo) GetApplications(length *T.Gsize) []string {
	return recentInfoGetApplications(r, length)
}
func (r *RecentInfo) GetDescription() string             { return recentInfoGetDescription(r) }
func (r *RecentInfo) GetDisplayName() string             { return recentInfoGetDisplayName(r) }
func (r *RecentInfo) GetGroups(length *T.Gsize) []string { return recentInfoGetGroups(r, length) }
func (r *RecentInfo) GetIcon(size int) *D.Pixbuf         { return recentInfoGetIcon(r, size) }
func (r *RecentInfo) GetMimeType() string                { return recentInfoGetMimeType(r) }
func (r *RecentInfo) GetModified() T.TimeT               { return recentInfoGetModified(r) }
func (r *RecentInfo) GetPrivateHint() bool               { return recentInfoGetPrivateHint(r) }
func (r *RecentInfo) GetShortName() string               { return recentInfoGetShortName(r) }
func (r *RecentInfo) GetUri() string                     { return recentInfoGetUri(r) }
func (r *RecentInfo) GetUriDisplay() string              { return recentInfoGetUriDisplay(r) }
func (r *RecentInfo) GetVisited() T.TimeT                { return recentInfoGetVisited(r) }
func (r *RecentInfo) HasApplication(appName string) bool {
	return recentInfoHasApplication(r, appName)
}
func (r *RecentInfo) HasGroup(groupName string) bool { return recentInfoHasGroup(r, groupName) }
func (r *RecentInfo) IsLocal() bool                  { return recentInfoIsLocal(r) }
func (r *RecentInfo) LastApplication() string        { return recentInfoLastApplication(r) }
func (r *RecentInfo) Match(infoB *RecentInfo) bool   { return recentInfoMatch(r, infoB) }
func (r *RecentInfo) Ref() *RecentInfo               { return recentInfoRef(r) }
func (r *RecentInfo) Unref()                         { recentInfoUnref(r) }

type (
	RecentFilter struct{}

	RecentFilterFunc func(
		filterInfo *RecentFilterInfo,
		userData T.Gpointer) bool

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

type RecentFilterFlags Enum

const (
	RECENT_FILTER_URI RecentFilterFlags = 1 << iota
	RECENT_FILTER_DISPLAY_NAME
	RECENT_FILTER_MIME_TYPE
	RECENT_FILTER_APPLICATION
	RECENT_FILTER_GROUP
	RECENT_FILTER_AGE
)

var (
	RecentFilterGetType func() O.Type
	RecentFilterNew     func() *RecentFilter

	RecentFilterFlagsGetType func() O.Type

	recentFilterAddAge           func(r *RecentFilter, days int)
	recentFilterAddApplication   func(r *RecentFilter, application string)
	recentFilterAddCustom        func(r *RecentFilter, needed RecentFilterFlags, f RecentFilterFunc, data T.Gpointer, dataDestroy T.GDestroyNotify)
	recentFilterAddGroup         func(r *RecentFilter, group string)
	recentFilterAddMimeType      func(r *RecentFilter, mimeType string)
	recentFilterAddPattern       func(r *RecentFilter, pattern string)
	recentFilterAddPixbufFormats func(r *RecentFilter)
	recentFilterFilter           func(r *RecentFilter, filterInfo *RecentFilterInfo) bool
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
func (r *RecentFilter) Filter(filterInfo *RecentFilterInfo) bool {
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
	IsPrivate   bool
}

var (
	RecentManagerGetType func() O.Type
	RecentManagerNew     func() *RecentManager

	RecentManagerErrorGetType func() O.Type
	RecentManagerErrorQuark   func() T.GQuark

	RecentManagerGetDefault   func() *RecentManager
	RecentManagerGetForScreen func(screen *D.Screen) *RecentManager

	recentManagerAddFull    func(r *RecentManager, uri string, recentData *RecentData) bool
	recentManagerAddItem    func(r *RecentManager, uri string) bool
	recentManagerGetItems   func(r *RecentManager) *T.GList
	recentManagerGetLimit   func(r *RecentManager) int
	recentManagerHasItem    func(r *RecentManager, uri string) bool
	recentManagerLookupItem func(r *RecentManager, uri string, err **T.GError) *RecentInfo
	recentManagerMoveItem   func(r *RecentManager, uri string, newUri string, err **T.GError) bool
	recentManagerPurgeItems func(r *RecentManager, err **T.GError) int
	recentManagerRemoveItem func(r *RecentManager, uri string, err **T.GError) bool
	recentManagerSetLimit   func(r *RecentManager, limit int)
	recentManagerSetScreen  func(r *RecentManager, screen *D.Screen)
)

func (r *RecentManager) AddFull(uri string, recentData *RecentData) bool {
	return recentManagerAddFull(r, uri, recentData)
}
func (r *RecentManager) AddItem(uri string) bool { return recentManagerAddItem(r, uri) }
func (r *RecentManager) GetItems() *T.GList      { return recentManagerGetItems(r) }
func (r *RecentManager) GetLimit() int           { return recentManagerGetLimit(r) }
func (r *RecentManager) HasItem(uri string) bool { return recentManagerHasItem(r, uri) }
func (r *RecentManager) LookupItem(uri string, err **T.GError) *RecentInfo {
	return recentManagerLookupItem(r, uri, err)
}
func (r *RecentManager) MoveItem(uri string, newUri string, err **T.GError) bool {
	return recentManagerMoveItem(r, uri, newUri, err)
}
func (r *RecentManager) PurgeItems(err **T.GError) int { return recentManagerPurgeItems(r, err) }
func (r *RecentManager) RemoveItem(uri string, err **T.GError) bool {
	return recentManagerRemoveItem(r, uri, err)
}
func (r *RecentManager) SetLimit(limit int)         { recentManagerSetLimit(r, limit) }
func (r *RecentManager) SetScreen(screen *D.Screen) { recentManagerSetScreen(r, screen) }

var RecentSortTypeGetType func() O.Type

type ReliefStyle Enum

const (
	RELIEF_NORMAL ReliefStyle = iota
	RELIEF_HALF
	RELIEF_NONE
)

var ReliefStyleGetType func() O.Type

type Requisition struct {
	Width  int
	Height int
}

var (
	RequisitionGetType func() O.Type

	requisitionCopy func(r *Requisition) *Requisition
	requisitionFree func(r *Requisition)
)

func (r *Requisition) Copy() *Requisition { return requisitionCopy(r) }
func (r *Requisition) Free()              { requisitionFree(r) }

type ResizeMode Enum

const (
	RESIZE_PARENT ResizeMode = iota
	RESIZE_QUEUE
	RESIZE_IMMEDIATE
)

var ResizeModeGetType func() O.Type

type ResponseType Enum

const (
	RESPONSE_NONE ResponseType = -(iota + 1)
	RESPONSE_REJECT
	RESPONSE_ACCEPT
	RESPONSE_DELETE_EVENT
	RESPONSE_OK
	RESPONSE_CANCEL
	RESPONSE_CLOSE
	RESPONSE_YES
	RESPONSE_NO
	RESPONSE_APPLY
	RESPONSE_HELP
)

var ResponseTypeGetType func() O.Type

type (
	Ruler struct {
		Widget       Widget
		BackingStore *D.Pixmap
		NonGrExpGc   *D.GC
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
	RulerGetType func() O.Type

	rulerDrawPos   func(r *Ruler)
	rulerDrawTicks func(r *Ruler)
	rulerGetMetric func(r *Ruler) MetricType
	rulerGetRange  func(r *Ruler, lower, upper, position, maxSize *float64)
	rulerSetMetric func(r *Ruler, metric MetricType)
	rulerSetRange  func(r *Ruler, lower, upper, position, maxSize float64)
)

func (r *Ruler) DrawPos()              { rulerDrawPos(r) }
func (r *Ruler) DrawTicks()            { rulerDrawTicks(r) }
func (r *Ruler) GetMetric() MetricType { return rulerGetMetric(r) }
func (r *Ruler) GetRange(lower, upper, position, maxSize *float64) {
	rulerGetRange(r, lower, upper, position, maxSize)
}
func (r *Ruler) SetMetric(metric MetricType) { rulerSetMetric(r, metric) }
func (r *Ruler) SetRange(lower, upper, position, maxSize float64) {
	rulerSetRange(r, lower, upper, position, maxSize)
}
