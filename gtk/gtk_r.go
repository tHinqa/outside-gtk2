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

	RadioActionGetCurrentValue func(r *RadioAction) int
	RadioActionGetGroup        func(r *RadioAction) *L.SList
	RadioActionSetCurrentValue func(r *RadioAction, currentValue int)
	RadioActionSetGroup        func(r *RadioAction, group *L.SList)
)

func (r *RadioAction) GetCurrentValue() int             { return RadioActionGetCurrentValue(r) }
func (r *RadioAction) GetGroup() *L.SList               { return RadioActionGetGroup(r) }
func (r *RadioAction) SetCurrentValue(currentValue int) { RadioActionSetCurrentValue(r, currentValue) }
func (r *RadioAction) SetGroup(group *L.SList)          { RadioActionSetGroup(r, group) }

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
	Group       *L.SList
}

var (
	RadioButtonGetType                   func() O.Type
	RadioButtonNew                       func(group *L.SList) *Widget
	RadioButtonNewFromWidget             func(radioGroupMember *RadioButton) *Widget
	RadioButtonNewWithLabel              func(group *L.SList, label string) *Widget
	RadioButtonNewWithLabelFromWidget    func(radioGroupMember *RadioButton, label string) *Widget
	RadioButtonNewWithMnemonic           func(group *L.SList, label string) *Widget
	RadioButtonNewWithMnemonicFromWidget func(radioGroupMember *RadioButton, label string) *Widget

	RadioButtonGetGroup func(r *RadioButton) *L.SList
	RadioButtonSetGroup func(r *RadioButton, group *L.SList)
)

func (r *RadioButton) GetGroup() *L.SList      { return RadioButtonGetGroup(r) }
func (r *RadioButton) SetGroup(group *L.SList) { RadioButtonSetGroup(r, group) }

type RadioMenuItem struct {
	CheckMenuItem CheckMenuItem
	Group         *L.SList
}

var (
	RadioMenuItemGetType                   func() O.Type
	RadioMenuItemNew                       func(group *L.SList) *Widget
	RadioMenuItemNewFromWidget             func(group *RadioMenuItem) *Widget
	RadioMenuItemNewWithLabel              func(group *L.SList, label string) *Widget
	RadioMenuItemNewWithLabelFromWidget    func(group *RadioMenuItem, label string) *Widget
	RadioMenuItemNewWithMnemonic           func(group *L.SList, label string) *Widget
	RadioMenuItemNewWithMnemonicFromWidget func(group *RadioMenuItem, label string) *Widget

	RadioMenuItemGetGroup func(r *RadioMenuItem) *L.SList
	RadioMenuItemSetGroup func(r *RadioMenuItem, group *L.SList)
)

func (r *RadioMenuItem) GetGroup() *L.SList      { return RadioMenuItemGetGroup(r) }
func (r *RadioMenuItem) SetGroup(group *L.SList) { RadioMenuItemSetGroup(r, group) }

type RadioToolButton struct {
	Parent ToggleToolButton
}

var (
	RadioToolButtonGetType                func() O.Type
	RadioToolButtonNew                    func(group *L.SList) *ToolItem
	RadioToolButtonNewFromStock           func(group *L.SList, stockId string) *ToolItem
	RadioToolButtonNewFromWidget          func(group *RadioToolButton) *ToolItem
	RadioToolButtonNewWithStockFromWidget func(group *RadioToolButton, stockId string) *ToolItem

	RadioToolButtonGetGroup func(r *RadioToolButton) *L.SList
	RadioToolButtonSetGroup func(r *RadioToolButton, group *L.SList)
)

func (r *RadioToolButton) GetGroup() *L.SList      { return RadioToolButtonGetGroup(r) }
func (r *RadioToolButton) SetGroup(group *L.SList) { RadioToolButtonSetGroup(r, group) }

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

	RangeGetAdjustment              func(r *Range) *Adjustment
	RangeGetFillLevel               func(r *Range) float64
	RangeGetFlippable               func(r *Range) bool
	RangeGetInverted                func(r *Range) bool
	RangeGetLowerStepperSensitivity func(r *Range) SensitivityType
	RangeGetMinSliderSize           func(r *Range) int
	RangeGetRangeRect               func(r *Range, rangeRect *D.Rectangle)
	RangeGetRestrictToFillLevel     func(r *Range) bool
	RangeGetRoundDigits             func(r *Range) int
	RangeGetShowFillLevel           func(r *Range) bool
	RangeGetSliderRange             func(r *Range, sliderStart, sliderEnd *int)
	RangeGetSliderSizeFixed         func(r *Range) bool
	RangeGetUpdatePolicy            func(r *Range) UpdateType
	RangeGetUpperStepperSensitivity func(r *Range) SensitivityType
	RangeGetValue                   func(r *Range) float64
	RangeSetAdjustment              func(r *Range, adjustment *Adjustment)
	RangeSetFillLevel               func(r *Range, fillLevel float64)
	RangeSetFlippable               func(r *Range, flippable bool)
	RangeSetIncrements              func(r *Range, step, page float64)
	RangeSetInverted                func(r *Range, setting bool)
	RangeSetLowerStepperSensitivity func(r *Range, sensitivity SensitivityType)
	RangeSetMinSliderSize           func(r *Range, minSize bool)
	RangeSetRange                   func(r *Range, min, max float64)
	RangeSetRestrictToFillLevel     func(r *Range, restrictToFillLevel bool)
	RangeSetRoundDigits             func(r *Range, roundDigits int)
	RangeSetShowFillLevel           func(r *Range, showFillLevel bool)
	RangeSetSliderSizeFixed         func(r *Range, sizeFixed bool)
	RangeSetUpdatePolicy            func(r *Range, policy UpdateType)
	RangeSetUpperStepperSensitivity func(r *Range, sensitivity SensitivityType)
	RangeSetValue                   func(r *Range, value float64)
)

func (r *Range) GetAdjustment() *Adjustment { return RangeGetAdjustment(r) }
func (r *Range) GetFillLevel() float64      { return RangeGetFillLevel(r) }
func (r *Range) GetFlippable() bool         { return RangeGetFlippable(r) }
func (r *Range) GetInverted() bool          { return RangeGetInverted(r) }
func (r *Range) GetLowerStepperSensitivity() SensitivityType {
	return RangeGetLowerStepperSensitivity(r)
}
func (r *Range) GetMinSliderSize() int               { return RangeGetMinSliderSize(r) }
func (r *Range) GetRangeRect(rangeRect *D.Rectangle) { RangeGetRangeRect(r, rangeRect) }
func (r *Range) GetRestrictToFillLevel() bool        { return RangeGetRestrictToFillLevel(r) }
func (r *Range) GetRoundDigits() int                 { return RangeGetRoundDigits(r) }
func (r *Range) GetShowFillLevel() bool              { return RangeGetShowFillLevel(r) }
func (r *Range) GetSliderRange(sliderStart, sliderEnd *int) {
	RangeGetSliderRange(r, sliderStart, sliderEnd)
}
func (r *Range) GetSliderSizeFixed() bool    { return RangeGetSliderSizeFixed(r) }
func (r *Range) GetUpdatePolicy() UpdateType { return RangeGetUpdatePolicy(r) }
func (r *Range) GetUpperStepperSensitivity() SensitivityType {
	return RangeGetUpperStepperSensitivity(r)
}
func (r *Range) GetValue() float64                    { return RangeGetValue(r) }
func (r *Range) SetAdjustment(adjustment *Adjustment) { RangeSetAdjustment(r, adjustment) }
func (r *Range) SetFillLevel(fillLevel float64)       { RangeSetFillLevel(r, fillLevel) }
func (r *Range) SetFlippable(flippable bool)          { RangeSetFlippable(r, flippable) }
func (r *Range) SetIncrements(step, page float64)     { RangeSetIncrements(r, step, page) }
func (r *Range) SetInverted(setting bool)             { RangeSetInverted(r, setting) }
func (r *Range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
	RangeSetLowerStepperSensitivity(r, sensitivity)
}
func (r *Range) SetMinSliderSize(minSize bool) { RangeSetMinSliderSize(r, minSize) }
func (r *Range) SetRange(min, max float64)     { RangeSetRange(r, min, max) }
func (r *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	RangeSetRestrictToFillLevel(r, restrictToFillLevel)
}
func (r *Range) SetRoundDigits(roundDigits int)      { RangeSetRoundDigits(r, roundDigits) }
func (r *Range) SetShowFillLevel(showFillLevel bool) { RangeSetShowFillLevel(r, showFillLevel) }
func (r *Range) SetSliderSizeFixed(sizeFixed bool)   { RangeSetSliderSizeFixed(r, sizeFixed) }
func (r *Range) SetUpdatePolicy(policy UpdateType)   { RangeSetUpdatePolicy(r, policy) }
func (r *Range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	RangeSetUpperStepperSensitivity(r, sensitivity)
}
func (r *Range) SetValue(value float64) { RangeSetValue(r, value) }

type RcContext struct{}

var (
	RcFlagsGetType     func() O.Type
	RcTokenTypeGetType func() O.Type

	RcAddDefaultFile           func(filename string)
	RcFindModuleInPath         func(moduleFile string) string
	RcFindPixmapInPath         func(settings *Settings, scanner *L.Scanner, pixmapFile string) string
	RcGetDefaultFiles          func() []string
	RcGetImModuleFile          func() string
	RcGetImModulePath          func() string
	RcGetModuleDir             func() string
	RcGetStyle                 func(widget *Widget) *Style
	RcGetStyleByPaths          func(settings *Settings, widgetPath, classPath string, t O.Type) *Style
	RcGetThemeDir              func() string
	RcParse                    func(filename string)
	RcParseString              func(rcString string)
	RcPropertyParseBorder      func(pspec *O.ParamSpec, gstring *L.String, propertyValue *O.Value) bool
	RcPropertyParseColor       func(pspec *O.ParamSpec, gstring *L.String, propertyValue *O.Value) bool
	RcPropertyParseEnum        func(pspec *O.ParamSpec, gstring *L.String, propertyValue *O.Value) bool
	RcPropertyParseFlags       func(pspec *O.ParamSpec, gstring *L.String, propertyValue *O.Value) bool
	RcPropertyParseRequisition func(pspec *O.ParamSpec, gstring *L.String, propertyValue *O.Value) bool
	RcReparseAll               func() bool
	RcReparseAllForSettings    func(settings *Settings, forceLoad bool) bool
	RcResetStyles              func(settings *Settings)
	RcSetDefaultFiles          func(filenames []string)

	RcScannerNew func() *L.Scanner

	RcParseColor     func(scanner *L.Scanner, color *D.Color) uint
	RcParseColorFull func(scanner *L.Scanner, style *RcStyle, color *D.Color) uint
	RcParseState     func(scanner *L.Scanner, state *StateType) uint
	RcParsePriority  func(scanner *L.Scanner, priority *PathPriorityType) uint
)

type RcPropertyParser func(pspec *O.ParamSpec,
	rcString *L.String, property_value *O.Value) bool

type RcStyle struct {
	Parent          O.Object
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
	RcStyleLists    *L.SList
	IconFactories   *L.SList
	EngineSpecified uint //: 1
}

type RcFlags Enum

const (
	RC_FG RcFlags = 1 << iota
	RC_BG
	RC_TEXT
	RC_BASE
)

type RcProperty struct {
	TypeName     L.Quark
	PropertyName L.Quark
	Origin       *T.Gchar
	Value        O.Value
}

var (
	RcStyleGetType func() O.Type
	RcStyleNew     func() *RcStyle

	RcAddWidgetNameStyle  func(r *RcStyle, pattern string)
	RcAddWidgetClassStyle func(r *RcStyle, pattern string)
	RcAddClassStyle       func(r *RcStyle, pattern string)

	RcStyleCopy  func(r *RcStyle) *RcStyle
	RcStyleRef   func(r *RcStyle)
	RcStyleUnref func(r *RcStyle)
)

func (r *RcStyle) Copy() *RcStyle { return RcStyleCopy(r) }
func (r *RcStyle) Ref()           { RcStyleRef(r) }
func (r *RcStyle) Unref()         { RcStyleUnref(r) }

type RecentAction struct {
	Parent Action
	_      *struct{}
}

var (
	RecentActionGetType       func() O.Type
	RecentActionNew           func(name, label, tooltip, stockId string) *Action
	RecentActionNewForManager func(name, label, tooltip, stockId string, manager *RecentManager) *Action

	RecentActionGetShowNumbers func(r *RecentAction) bool
	RecentActionSetShowNumbers func(r *RecentAction, showNumbers bool)
)

func (r *RecentAction) GetShowNumbers() bool { return RecentActionGetShowNumbers(r) }
func (r *RecentAction) SetShowNumbers(showNumbers bool) {
	RecentActionSetShowNumbers(r, showNumbers)
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

	RecentChooserErrorQuark func() L.Quark

	RecentChooserAddFilter         func(r *RecentChooser, filter *RecentFilter)
	RecentChooserGetCurrentItem    func(r *RecentChooser) *RecentInfo
	RecentChooserGetCurrentUri     func(r *RecentChooser) string
	RecentChooserGetFilter         func(r *RecentChooser) *RecentFilter
	RecentChooserGetItems          func(r *RecentChooser) *L.List
	RecentChooserGetLimit          func(r *RecentChooser) int
	RecentChooserGetLocalOnly      func(r *RecentChooser) bool
	RecentChooserGetSelectMultiple func(r *RecentChooser) bool
	RecentChooserGetShowIcons      func(r *RecentChooser) bool
	RecentChooserGetShowNotFound   func(r *RecentChooser) bool
	RecentChooserGetShowNumbers    func(r *RecentChooser) bool
	RecentChooserGetShowPrivate    func(r *RecentChooser) bool
	RecentChooserGetShowTips       func(r *RecentChooser) bool
	RecentChooserGetSortType       func(r *RecentChooser) RecentSortType
	RecentChooserGetUris           func(r *RecentChooser, length *T.Gsize) []string
	RecentChooserListFilters       func(r *RecentChooser) *L.SList
	RecentChooserRemoveFilter      func(r *RecentChooser, filter *RecentFilter)
	RecentChooserSelectAll         func(r *RecentChooser)
	RecentChooserSelectUri         func(r *RecentChooser, uri string, err **L.Error) bool
	RecentChooserSetCurrentUri     func(r *RecentChooser, uri string, err **L.Error) bool
	RecentChooserSetFilter         func(r *RecentChooser, filter *RecentFilter)
	RecentChooserSetLimit          func(r *RecentChooser, limit int)
	RecentChooserSetLocalOnly      func(r *RecentChooser, localOnly bool)
	RecentChooserSetSelectMultiple func(r *RecentChooser, selectMultiple bool)
	RecentChooserSetShowIcons      func(r *RecentChooser, showIcons bool)
	RecentChooserSetShowNotFound   func(r *RecentChooser, showNotFound bool)
	RecentChooserSetShowNumbers    func(r *RecentChooser, showNumbers bool)
	RecentChooserSetShowPrivate    func(r *RecentChooser, showPrivate bool)
	RecentChooserSetShowTips       func(r *RecentChooser, showTips bool)
	RecentChooserSetSortFunc       func(r *RecentChooser, sortFunc RecentSortFunc, sortData T.Gpointer, dataDestroy O.DestroyNotify)
	RecentChooserSetSortType       func(r *RecentChooser, sortType RecentSortType)
	RecentChooserUnselectAll       func(r *RecentChooser)
	RecentChooserUnselectUri       func(r *RecentChooser, uri string)
)

func (r *RecentChooser) AddFilter(filter *RecentFilter)    { RecentChooserAddFilter(r, filter) }
func (r *RecentChooser) GetCurrentItem() *RecentInfo       { return RecentChooserGetCurrentItem(r) }
func (r *RecentChooser) GetCurrentUri() string             { return RecentChooserGetCurrentUri(r) }
func (r *RecentChooser) GetFilter() *RecentFilter          { return RecentChooserGetFilter(r) }
func (r *RecentChooser) GetItems() *L.List                 { return RecentChooserGetItems(r) }
func (r *RecentChooser) GetLimit() int                     { return RecentChooserGetLimit(r) }
func (r *RecentChooser) GetLocalOnly() bool                { return RecentChooserGetLocalOnly(r) }
func (r *RecentChooser) GetSelectMultiple() bool           { return RecentChooserGetSelectMultiple(r) }
func (r *RecentChooser) GetShowIcons() bool                { return RecentChooserGetShowIcons(r) }
func (r *RecentChooser) GetShowNotFound() bool             { return RecentChooserGetShowNotFound(r) }
func (r *RecentChooser) GetShowNumbers() bool              { return RecentChooserGetShowNumbers(r) }
func (r *RecentChooser) GetShowPrivate() bool              { return RecentChooserGetShowPrivate(r) }
func (r *RecentChooser) GetShowTips() bool                 { return RecentChooserGetShowTips(r) }
func (r *RecentChooser) GetSortType() RecentSortType       { return RecentChooserGetSortType(r) }
func (r *RecentChooser) GetUris(length *T.Gsize) []string  { return RecentChooserGetUris(r, length) }
func (r *RecentChooser) ListFilters() *L.SList             { return RecentChooserListFilters(r) }
func (r *RecentChooser) RemoveFilter(filter *RecentFilter) { RecentChooserRemoveFilter(r, filter) }
func (r *RecentChooser) SelectAll()                        { RecentChooserSelectAll(r) }
func (r *RecentChooser) SelectUri(uri string, err **L.Error) bool {
	return RecentChooserSelectUri(r, uri, err)
}
func (r *RecentChooser) SetCurrentUri(uri string, err **L.Error) bool {
	return RecentChooserSetCurrentUri(r, uri, err)
}
func (r *RecentChooser) SetFilter(filter *RecentFilter) { RecentChooserSetFilter(r, filter) }
func (r *RecentChooser) SetLimit(limit int)             { RecentChooserSetLimit(r, limit) }
func (r *RecentChooser) SetLocalOnly(localOnly bool)    { RecentChooserSetLocalOnly(r, localOnly) }
func (r *RecentChooser) SetSelectMultiple(selectMultiple bool) {
	RecentChooserSetSelectMultiple(r, selectMultiple)
}
func (r *RecentChooser) SetShowIcons(showIcons bool) { RecentChooserSetShowIcons(r, showIcons) }
func (r *RecentChooser) SetShowNotFound(showNotFound bool) {
	RecentChooserSetShowNotFound(r, showNotFound)
}
func (r *RecentChooser) SetShowNumbers(showNumbers bool) {
	RecentChooserSetShowNumbers(r, showNumbers)
}
func (r *RecentChooser) SetShowPrivate(showPrivate bool) {
	RecentChooserSetShowPrivate(r, showPrivate)
}
func (r *RecentChooser) SetShowTips(showTips bool) { RecentChooserSetShowTips(r, showTips) }
func (r *RecentChooser) SetSortFunc(sortFunc RecentSortFunc, sortData T.Gpointer, dataDestroy O.DestroyNotify) {
	RecentChooserSetSortFunc(r, sortFunc, sortData, dataDestroy)
}
func (r *RecentChooser) SetSortType(sortType RecentSortType) { RecentChooserSetSortType(r, sortType) }
func (r *RecentChooser) UnselectAll()                        { RecentChooserUnselectAll(r) }
func (r *RecentChooser) UnselectUri(uri string)              { RecentChooserUnselectUri(r, uri) }

type RecentChooserMenu struct {
	Parent Menu
	_      *struct{}
}

var (
	RecentChooserMenuGetType func() O.Type
	RecentChooserMenuNew     func() *Widget

	RecentChooserMenuGetShowNumbers func(r *RecentChooserMenu) bool
	RecentChooserMenuSetShowNumbers func(r *RecentChooserMenu, showNumbers bool)
)

func (r *RecentChooserMenu) GetShowNumbers() bool { return RecentChooserMenuGetShowNumbers(r) }
func (r *RecentChooserMenu) SetShowNumbers(showNumbers bool) {
	RecentChooserMenuSetShowNumbers(r, showNumbers)
}

type (
	RecentInfo struct{}

	RecentSortFunc func(
		a, b *RecentInfo, userData T.Gpointer) int
)

var (
	RecentInfoGetType func() O.Type

	RecentInfoExists             func(r *RecentInfo) bool
	RecentInfoGetAdded           func(r *RecentInfo) T.TimeT
	RecentInfoGetAge             func(r *RecentInfo) int
	RecentInfoGetApplicationInfo func(r *RecentInfo, appName string, appExec **T.Gchar, count *uint, time *T.TimeT) bool
	RecentInfoGetApplications    func(r *RecentInfo, length *T.Gsize) []string
	RecentInfoGetDescription     func(r *RecentInfo) string
	RecentInfoGetDisplayName     func(r *RecentInfo) string
	RecentInfoGetGroups          func(r *RecentInfo, length *T.Gsize) []string
	RecentInfoGetIcon            func(r *RecentInfo, size int) *D.Pixbuf
	RecentInfoGetMimeType        func(r *RecentInfo) string
	RecentInfoGetModified        func(r *RecentInfo) T.TimeT
	RecentInfoGetPrivateHint     func(r *RecentInfo) bool
	RecentInfoGetShortName       func(r *RecentInfo) string
	RecentInfoGetUri             func(r *RecentInfo) string
	RecentInfoGetUriDisplay      func(r *RecentInfo) string
	RecentInfoGetVisited         func(r *RecentInfo) T.TimeT
	RecentInfoHasApplication     func(r *RecentInfo, appName string) bool
	RecentInfoHasGroup           func(r *RecentInfo, groupName string) bool
	RecentInfoIsLocal            func(r *RecentInfo) bool
	RecentInfoLastApplication    func(r *RecentInfo) string
	RecentInfoMatch              func(r *RecentInfo, infoB *RecentInfo) bool
	RecentInfoRef                func(r *RecentInfo) *RecentInfo
	RecentInfoUnref              func(r *RecentInfo)
)

func (r *RecentInfo) Exists() bool      { return RecentInfoExists(r) }
func (r *RecentInfo) GetAdded() T.TimeT { return RecentInfoGetAdded(r) }
func (r *RecentInfo) GetAge() int       { return RecentInfoGetAge(r) }
func (r *RecentInfo) GetApplicationInfo(appName string, appExec **T.Gchar, count *uint, time *T.TimeT) bool {
	return RecentInfoGetApplicationInfo(r, appName, appExec, count, time)
}
func (r *RecentInfo) GetApplications(length *T.Gsize) []string {
	return RecentInfoGetApplications(r, length)
}
func (r *RecentInfo) GetDescription() string             { return RecentInfoGetDescription(r) }
func (r *RecentInfo) GetDisplayName() string             { return RecentInfoGetDisplayName(r) }
func (r *RecentInfo) GetGroups(length *T.Gsize) []string { return RecentInfoGetGroups(r, length) }
func (r *RecentInfo) GetIcon(size int) *D.Pixbuf         { return RecentInfoGetIcon(r, size) }
func (r *RecentInfo) GetMimeType() string                { return RecentInfoGetMimeType(r) }
func (r *RecentInfo) GetModified() T.TimeT               { return RecentInfoGetModified(r) }
func (r *RecentInfo) GetPrivateHint() bool               { return RecentInfoGetPrivateHint(r) }
func (r *RecentInfo) GetShortName() string               { return RecentInfoGetShortName(r) }
func (r *RecentInfo) GetUri() string                     { return RecentInfoGetUri(r) }
func (r *RecentInfo) GetUriDisplay() string              { return RecentInfoGetUriDisplay(r) }
func (r *RecentInfo) GetVisited() T.TimeT                { return RecentInfoGetVisited(r) }
func (r *RecentInfo) HasApplication(appName string) bool {
	return RecentInfoHasApplication(r, appName)
}
func (r *RecentInfo) HasGroup(groupName string) bool { return RecentInfoHasGroup(r, groupName) }
func (r *RecentInfo) IsLocal() bool                  { return RecentInfoIsLocal(r) }
func (r *RecentInfo) LastApplication() string        { return RecentInfoLastApplication(r) }
func (r *RecentInfo) Match(infoB *RecentInfo) bool   { return RecentInfoMatch(r, infoB) }
func (r *RecentInfo) Ref() *RecentInfo               { return RecentInfoRef(r) }
func (r *RecentInfo) Unref()                         { RecentInfoUnref(r) }

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

	RecentFilterAddAge           func(r *RecentFilter, days int)
	RecentFilterAddApplication   func(r *RecentFilter, application string)
	RecentFilterAddCustom        func(r *RecentFilter, needed RecentFilterFlags, f RecentFilterFunc, data T.Gpointer, dataDestroy O.DestroyNotify)
	RecentFilterAddGroup         func(r *RecentFilter, group string)
	RecentFilterAddMimeType      func(r *RecentFilter, mimeType string)
	RecentFilterAddPattern       func(r *RecentFilter, pattern string)
	RecentFilterAddPixbufFormats func(r *RecentFilter)
	RecentFilterFilter           func(r *RecentFilter, filterInfo *RecentFilterInfo) bool
	RecentFilterGetName          func(r *RecentFilter) string
	RecentFilterGetNeeded        func(r *RecentFilter) RecentFilterFlags
	RecentFilterSetName          func(r *RecentFilter, name string)
)

func (r *RecentFilter) AddAge(days int)                   { RecentFilterAddAge(r, days) }
func (r *RecentFilter) AddApplication(application string) { RecentFilterAddApplication(r, application) }
func (r *RecentFilter) AddCustom(needed RecentFilterFlags, f RecentFilterFunc, data T.Gpointer, dataDestroy O.DestroyNotify) {
	RecentFilterAddCustom(r, needed, f, data, dataDestroy)
}
func (r *RecentFilter) AddGroup(group string)       { RecentFilterAddGroup(r, group) }
func (r *RecentFilter) AddMimeType(mimeType string) { RecentFilterAddMimeType(r, mimeType) }
func (r *RecentFilter) AddPattern(pattern string)   { RecentFilterAddPattern(r, pattern) }
func (r *RecentFilter) AddPixbufFormats()           { RecentFilterAddPixbufFormats(r) }
func (r *RecentFilter) Filter(filterInfo *RecentFilterInfo) bool {
	return RecentFilterFilter(r, filterInfo)
}
func (r *RecentFilter) GetName() string              { return RecentFilterGetName(r) }
func (r *RecentFilter) GetNeeded() RecentFilterFlags { return RecentFilterGetNeeded(r) }
func (r *RecentFilter) SetName(name string)          { RecentFilterSetName(r, name) }

type RecentManager struct {
	Parent O.Object
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
	RecentManagerErrorQuark   func() L.Quark

	RecentManagerGetDefault   func() *RecentManager
	RecentManagerGetForScreen func(screen *D.Screen) *RecentManager

	RecentManagerAddFull    func(r *RecentManager, uri string, recentData *RecentData) bool
	RecentManagerAddItem    func(r *RecentManager, uri string) bool
	RecentManagerGetItems   func(r *RecentManager) *L.List
	RecentManagerGetLimit   func(r *RecentManager) int
	RecentManagerHasItem    func(r *RecentManager, uri string) bool
	RecentManagerLookupItem func(r *RecentManager, uri string, err **L.Error) *RecentInfo
	RecentManagerMoveItem   func(r *RecentManager, uri string, newUri string, err **L.Error) bool
	RecentManagerPurgeItems func(r *RecentManager, err **L.Error) int
	RecentManagerRemoveItem func(r *RecentManager, uri string, err **L.Error) bool
	RecentManagerSetLimit   func(r *RecentManager, limit int)
	RecentManagerSetScreen  func(r *RecentManager, screen *D.Screen)
)

func (r *RecentManager) AddFull(uri string, recentData *RecentData) bool {
	return RecentManagerAddFull(r, uri, recentData)
}
func (r *RecentManager) AddItem(uri string) bool { return RecentManagerAddItem(r, uri) }
func (r *RecentManager) GetItems() *L.List       { return RecentManagerGetItems(r) }
func (r *RecentManager) GetLimit() int           { return RecentManagerGetLimit(r) }
func (r *RecentManager) HasItem(uri string) bool { return RecentManagerHasItem(r, uri) }
func (r *RecentManager) LookupItem(uri string, err **L.Error) *RecentInfo {
	return RecentManagerLookupItem(r, uri, err)
}
func (r *RecentManager) MoveItem(uri string, newUri string, err **L.Error) bool {
	return RecentManagerMoveItem(r, uri, newUri, err)
}
func (r *RecentManager) PurgeItems(err **L.Error) int { return RecentManagerPurgeItems(r, err) }
func (r *RecentManager) RemoveItem(uri string, err **L.Error) bool {
	return RecentManagerRemoveItem(r, uri, err)
}
func (r *RecentManager) SetLimit(limit int)         { RecentManagerSetLimit(r, limit) }
func (r *RecentManager) SetScreen(screen *D.Screen) { RecentManagerSetScreen(r, screen) }

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

	RequisitionCopy func(r *Requisition) *Requisition
	RequisitionFree func(r *Requisition)
)

func (r *Requisition) Copy() *Requisition { return RequisitionCopy(r) }
func (r *Requisition) Free()              { RequisitionFree(r) }

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

	RulerDrawPos   func(r *Ruler)
	RulerDrawTicks func(r *Ruler)
	RulerGetMetric func(r *Ruler) MetricType
	RulerGetRange  func(r *Ruler, lower, upper, position, maxSize *float64)
	RulerSetMetric func(r *Ruler, metric MetricType)
	RulerSetRange  func(r *Ruler, lower, upper, position, maxSize float64)
)

func (r *Ruler) DrawPos()              { RulerDrawPos(r) }
func (r *Ruler) DrawTicks()            { RulerDrawTicks(r) }
func (r *Ruler) GetMetric() MetricType { return RulerGetMetric(r) }
func (r *Ruler) GetRange(lower, upper, position, maxSize *float64) {
	RulerGetRange(r, lower, upper, position, maxSize)
}
func (r *Ruler) SetMetric(metric MetricType) { RulerSetMetric(r, metric) }
func (r *Ruler) SetRange(lower, upper, position, maxSize float64) {
	RulerSetRange(r, lower, upper, position, maxSize)
}
