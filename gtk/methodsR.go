// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
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
		EventWindow                *D.Window
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
	rangeGetLowerStepperSensitivity func(r *Range) SensitivityType
	rangeGetMinSliderSize           func(r *Range) int
	rangeGetRangeRect               func(r *Range, rangeRect *T.GdkRectangle)
	rangeGetRestrictToFillLevel     func(r *Range) T.Gboolean
	rangeGetRoundDigits             func(r *Range) int
	rangeGetShowFillLevel           func(r *Range) T.Gboolean
	rangeGetSliderRange             func(r *Range, sliderStart, sliderEnd *int)
	rangeGetSliderSizeFixed         func(r *Range) T.Gboolean
	rangeGetUpdatePolicy            func(r *Range) UpdateType
	rangeGetUpperStepperSensitivity func(r *Range) SensitivityType
	rangeGetValue                   func(r *Range) float64
	rangeSetAdjustment              func(r *Range, adjustment *Adjustment)
	rangeSetFillLevel               func(r *Range, fillLevel float64)
	rangeSetFlippable               func(r *Range, flippable T.Gboolean)
	rangeSetIncrements              func(r *Range, step, page float64)
	rangeSetInverted                func(r *Range, setting T.Gboolean)
	rangeSetLowerStepperSensitivity func(r *Range, sensitivity SensitivityType)
	rangeSetMinSliderSize           func(r *Range, minSize T.Gboolean)
	rangeSetRange                   func(r *Range, min, max float64)
	rangeSetRestrictToFillLevel     func(r *Range, restrictToFillLevel T.Gboolean)
	rangeSetRoundDigits             func(r *Range, roundDigits int)
	rangeSetShowFillLevel           func(r *Range, showFillLevel T.Gboolean)
	rangeSetSliderSizeFixed         func(r *Range, sizeFixed T.Gboolean)
	rangeSetUpdatePolicy            func(r *Range, policy UpdateType)
	rangeSetUpperStepperSensitivity func(r *Range, sensitivity SensitivityType)
	rangeSetValue                   func(r *Range, value float64)
)

func (r *Range) GetAdjustment() *Adjustment { return rangeGetAdjustment(r) }
func (r *Range) GetFillLevel() float64      { return rangeGetFillLevel(r) }
func (r *Range) GetFlippable() T.Gboolean   { return rangeGetFlippable(r) }
func (r *Range) GetInverted() T.Gboolean    { return rangeGetInverted(r) }
func (r *Range) GetLowerStepperSensitivity() SensitivityType {
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
func (r *Range) GetSliderSizeFixed() T.Gboolean { return rangeGetSliderSizeFixed(r) }
func (r *Range) GetUpdatePolicy() UpdateType    { return rangeGetUpdatePolicy(r) }
func (r *Range) GetUpperStepperSensitivity() SensitivityType {
	return rangeGetUpperStepperSensitivity(r)
}
func (r *Range) GetValue() float64                    { return rangeGetValue(r) }
func (r *Range) SetAdjustment(adjustment *Adjustment) { rangeSetAdjustment(r, adjustment) }
func (r *Range) SetFillLevel(fillLevel float64)       { rangeSetFillLevel(r, fillLevel) }
func (r *Range) SetFlippable(flippable T.Gboolean)    { rangeSetFlippable(r, flippable) }
func (r *Range) SetIncrements(step, page float64)     { rangeSetIncrements(r, step, page) }
func (r *Range) SetInverted(setting T.Gboolean)       { rangeSetInverted(r, setting) }
func (r *Range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
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
func (r *Range) SetUpdatePolicy(policy UpdateType)         { rangeSetUpdatePolicy(r, policy) }
func (r *Range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	rangeSetUpperStepperSensitivity(r, sensitivity)
}
func (r *Range) SetValue(value float64) { rangeSetValue(r, value) }

type RcContext struct{}

var (
	RcFlagsGetType     func() T.GType
	RcTokenTypeGetType func() T.GType

	RcAddDefaultFile           func(filename string)
	RcFindModuleInPath         func(moduleFile string) string
	RcFindPixmapInPath         func(settings *Settings, scanner *T.GScanner, pixmapFile string) string
	RcGetDefaultFiles          func() **T.Gchar
	RcGetImModuleFile          func() string
	RcGetImModulePath          func() string
	RcGetModuleDir             func() string
	RcGetStyle                 func(widget *Widget) *Style
	RcGetStyleByPaths          func(settings *Settings, widgetPath, classPath string, t T.GType) *Style
	RcGetThemeDir              func() string
	RcParse                    func(filename string)
	RcParseString              func(rcString string)
	RcPropertyParseBorder      func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) T.Gboolean
	RcPropertyParseColor       func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) T.Gboolean
	RcPropertyParseEnum        func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) T.Gboolean
	RcPropertyParseFlags       func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) T.Gboolean
	RcPropertyParseRequisition func(pspec *T.GParamSpec, gstring *T.GString, propertyValue *T.GValue) T.Gboolean
	RcReparseAll               func() T.Gboolean
	RcReparseAllForSettings    func(settings *Settings, forceLoad T.Gboolean) T.Gboolean
	RcResetStyles              func(settings *Settings)
	RcSetDefaultFiles          func(filenames **T.Gchar)

	RcScannerNew func() *T.GScanner

	RcParseColor     func(scanner *T.GScanner, color *D.Color) uint
	RcParseColorFull func(scanner *T.GScanner, style *RcStyle, color *D.Color) uint
	RcParseState     func(scanner *T.GScanner, state *StateType) uint
	RcParsePriority  func(scanner *T.GScanner, priority *PathPriorityType) uint
)

type RcPropertyParser func(pspec *T.GParamSpec,
	rcString *T.GString, property_value *T.GValue) T.Gboolean

type RcStyle struct {
	Parent          T.GObject
	Name            *T.Gchar
	BgPixmapName    *[5]T.Gchar //TODO(t): CHECK
	FontDesc        *T.PangoFontDescription
	ColorFlags      [5]RcFlags
	Fg              [5]D.Color
	Bg              [5]D.Color
	Text            [5]D.Color
	Base            [5]D.Color
	Xthickness      int
	Ythickness      int
	RcProperties    *T.GArray
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
	RcStyleGetType func() T.GType
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

type RecentSortType Enum

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
	recentInfoGetIcon            func(r *RecentInfo, size int) *D.Pixbuf
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
func (r *RecentInfo) GetIcon(size int) *D.Pixbuf          { return recentInfoGetIcon(r, size) }
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
	RecentManagerGetForScreen func(screen *D.Screen) *RecentManager

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
	recentManagerSetScreen  func(r *RecentManager, screen *D.Screen)
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
func (r *RecentManager) SetLimit(limit int)         { recentManagerSetLimit(r, limit) }
func (r *RecentManager) SetScreen(screen *D.Screen) { recentManagerSetScreen(r, screen) }

var RecentSortTypeGetType func() T.GType

type ReliefStyle Enum

const (
	RELIEF_NORMAL ReliefStyle = iota
	RELIEF_HALF
	RELIEF_NONE
)

var ReliefStyleGetType func() T.GType

type Requisition struct {
	Width  int
	Height int
}

var (
	RequisitionGetType func() T.GType

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

var ResizeModeGetType func() T.GType

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

var ResponseTypeGetType func() T.GType

type (
	Ruler struct {
		Widget       Widget
		BackingStore *D.Pixmap
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
