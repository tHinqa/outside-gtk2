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

type MatchType Enum

const (
	MATCH_ALL MatchType = iota
	MATCH_ALL_TAIL
	MATCH_HEAD
	MATCH_TAIL
	MATCH_EXACT
	MATCH_LAST
)

var MatchTypeGetType func() O.Type

type (
	Menu struct {
		MenuShell         MenuShell
		ParentMenuItem    *Widget
		OldActiveMenuItem *Widget
		AccelGroup        *AccelGroup
		AccelPath         *T.Gchar
		PositionFunc      MenuPositionFunc
		PositionFuncData  T.Gpointer
		ToggleSize        uint
		Toplevel          *Widget
		TearoffWindow     *Widget
		TearoffHbox       *Widget
		TearoffScrollbar  *Widget
		TearoffAdjustment *Adjustment
		ViewWindow        *D.Window
		BinWindow         *D.Window
		ScrollOffset      int
		SavedScrollOffset int
		ScrollStep        int
		TimeoutId         uint
		NavigationRegion  *D.Region
		NavigationTimeout uint
		Bits              uint
		// NeedsDestructionRefCount : 1
		// TornOff : 1
		// TearoffActive : 1
		// ScrollFast : 1
		// upperArrowVisible : 1
		// lowerArrowVisible : 1
		// upperArrowPrelight : 1
		// lowerArrowPrelight : 1
	}

	MenuEntry struct {
		Path         *T.Gchar
		Accelerator  *T.Gchar
		Callback     MenuCallback
		CallbackData T.Gpointer
		Widget       *Widget
	}

	MenuCallback func(widget *Widget, userData T.Gpointer)

	MenuPositionFunc func(menu *Menu,
		x, y *int, pushIn *bool, userData T.Gpointer)

	MenuDetachFunc func(ttachWidget *Widget, menu *Menu)
)

type MenuDirectionType Enum

const (
	MENU_DIR_PARENT MenuDirectionType = iota
	ENU_DIR_CHILD
	MENU_DIR_NEXT
	MENU_DIR_PREV
)

var (
	MenuGetType func() O.Type
	MenuNew     func() *Widget

	MenuDirectionTypeGetType func() O.Type
	MenuGetForAttachWidget   func(widget *Widget) *T.GList

	MenuAttach               func(m *Menu, child *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint)
	MenuAttachToWidget       func(m *Menu, attachWidget *Widget, detacher MenuDetachFunc)
	MenuDetach               func(m *Menu)
	MenuGetAccelGroup        func(m *Menu) *AccelGroup
	MenuGetAccelPath         func(m *Menu) string
	MenuGetActive            func(m *Menu) *Widget
	MenuGetAttachWidget      func(m *Menu) *Widget
	MenuGetMonitor           func(m *Menu) int
	MenuGetReserveToggleSize func(m *Menu) bool
	MenuGetTearoffState      func(m *Menu) bool
	MenuGetTitle             func(m *Menu) string
	MenuPopdown              func(m *Menu)
	MenuPopup                func(m *Menu, parentMenuShell, parentMenuItem *Widget, f MenuPositionFunc, data T.Gpointer, button uint, activateTime T.GUint32)
	MenuReorderChild         func(m *Menu, child *Widget, position int)
	MenuReposition           func(m *Menu)
	MenuSetAccelGroup        func(m *Menu, accelGroup *AccelGroup)
	MenuSetAccelPath         func(m *Menu, accelPath string)
	MenuSetActive            func(m *Menu, index uint)
	MenuSetMonitor           func(m *Menu, monitorNum int)
	MenuSetReserveToggleSize func(m *Menu, reserveToggleSize bool)
	MenuSetScreen            func(m *Menu, screen *D.Screen)
	MenuSetTearoffState      func(m *Menu, tornOff bool)
	MenuSetTitle             func(m *Menu, title string)
)

func (m *Menu) Attach(child *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint) {
	MenuAttach(m, child, leftAttach, rightAttach, topAttach, bottomAttach)
}
func (m *Menu) AttachToWidget(attachWidget *Widget, detacher MenuDetachFunc) {
	MenuAttachToWidget(m, attachWidget, detacher)
}
func (m *Menu) Detach()                    { MenuDetach(m) }
func (m *Menu) GetAccelGroup() *AccelGroup { return MenuGetAccelGroup(m) }
func (m *Menu) GetAccelPath() string       { return MenuGetAccelPath(m) }
func (m *Menu) GetActive() *Widget         { return MenuGetActive(m) }
func (m *Menu) GetAttachWidget() *Widget   { return MenuGetAttachWidget(m) }
func (m *Menu) GetMonitor() int            { return MenuGetMonitor(m) }
func (m *Menu) GetReserveToggleSize() bool { return MenuGetReserveToggleSize(m) }
func (m *Menu) GetTearoffState() bool      { return MenuGetTearoffState(m) }
func (m *Menu) GetTitle() string           { return MenuGetTitle(m) }
func (m *Menu) Popdown()                   { MenuPopdown(m) }
func (m *Menu) Popup(parentMenuShell, parentMenuItem *Widget, f MenuPositionFunc, data T.Gpointer, button uint, activateTime T.GUint32) {
	MenuPopup(m, parentMenuShell, parentMenuItem, f, data, button, activateTime)
}
func (m *Menu) ReorderChild(child *Widget, position int) { MenuReorderChild(m, child, position) }
func (m *Menu) Reposition()                              { MenuReposition(m) }
func (m *Menu) SetAccelGroup(accelGroup *AccelGroup)     { MenuSetAccelGroup(m, accelGroup) }
func (m *Menu) SetAccelPath(accelPath string)            { MenuSetAccelPath(m, accelPath) }
func (m *Menu) SetActive(index uint)                     { MenuSetActive(m, index) }
func (m *Menu) SetMonitor(monitorNum int)                { MenuSetMonitor(m, monitorNum) }
func (m *Menu) SetReserveToggleSize(reserveToggleSize bool) {
	MenuSetReserveToggleSize(m, reserveToggleSize)
}
func (m *Menu) SetScreen(screen *D.Screen)   { MenuSetScreen(m, screen) }
func (m *Menu) SetTearoffState(tornOff bool) { MenuSetTearoffState(m, tornOff) }
func (m *Menu) SetTitle(title string)        { MenuSetTitle(m, title) }

type MenuBar struct {
	MenuShell MenuShell
}

var (
	MenuBarGetType func() O.Type
	MenuBarNew     func() *Widget

	MenuBarGetChildPackDirection func(m *MenuBar) PackDirection
	MenuBarGetPackDirection      func(m *MenuBar) PackDirection
	MenuBarSetChildPackDirection func(m *MenuBar, childPackDir PackDirection)
	MenuBarSetPackDirection      func(m *MenuBar, packDir PackDirection)
)

func (m *MenuBar) GetChildPackDirection() PackDirection { return MenuBarGetChildPackDirection(m) }
func (m *MenuBar) GetPackDirection() PackDirection      { return MenuBarGetPackDirection(m) }
func (m *MenuBar) SetChildPackDirection(childPackDir PackDirection) {
	MenuBarSetChildPackDirection(m, childPackDir)
}
func (m *MenuBar) SetPackDirection(packDir PackDirection) { MenuBarSetPackDirection(m, packDir) }

type (
	MenuItem struct {
		Item             Item
		Submenu          *Widget
		EventWindow      *D.Window
		ToggleSize       uint16
		AcceleratorWidth uint16
		AccelPath        *T.Gchar
		Bits             uint
		// ShowSubmenuIndicator : 1
		// SubmenuPlacement : 1
		// SubmenuDirection : 1
		// RightJustify: 1
		// TimerFromKeypress : 1
		// FromMenubar : 1
		Timer uint
	}
)

var (
	MenuItemGetType         func() O.Type
	MenuItemNew             func() *Widget
	MenuItemNewWithLabel    func(label string) *Widget
	MenuItemNewWithMnemonic func(label string) *Widget

	MenuItemActivate           func(m *MenuItem)
	MenuItemDeselect           func(m *MenuItem)
	MenuItemGetAccelPath       func(m *MenuItem) string
	MenuItemGetLabel           func(m *MenuItem) string
	MenuItemGetRightJustified  func(m *MenuItem) bool
	MenuItemGetSubmenu         func(m *MenuItem) *Widget
	MenuItemGetUseUnderline    func(m *MenuItem) bool
	MenuItemRemoveSubmenu      func(m *MenuItem)
	MenuItemSelect             func(m *MenuItem)
	MenuItemSetAccelPath       func(m *MenuItem, accelPath string)
	MenuItemSetLabel           func(m *MenuItem, label string)
	MenuItemSetRightJustified  func(m *MenuItem, rightJustified bool)
	MenuItemSetSubmenu         func(m *MenuItem, submenu *Widget)
	MenuItemSetUseUnderline    func(m *MenuItem, setting bool)
	MenuItemToggleSizeAllocate func(m *MenuItem, allocation int)
	MenuItemToggleSizeRequest  func(m *MenuItem, requisition *int)
)

func (m *MenuItem) Activate()                     { MenuItemActivate(m) }
func (m *MenuItem) Deselect()                     { MenuItemDeselect(m) }
func (m *MenuItem) GetAccelPath() string          { return MenuItemGetAccelPath(m) }
func (m *MenuItem) GetLabel() string              { return MenuItemGetLabel(m) }
func (m *MenuItem) GetRightJustified() bool       { return MenuItemGetRightJustified(m) }
func (m *MenuItem) GetSubmenu() *Widget           { return MenuItemGetSubmenu(m) }
func (m *MenuItem) GetUseUnderline() bool         { return MenuItemGetUseUnderline(m) }
func (m *MenuItem) RemoveSubmenu()                { MenuItemRemoveSubmenu(m) }
func (m *MenuItem) Select()                       { MenuItemSelect(m) }
func (m *MenuItem) SetAccelPath(accelPath string) { MenuItemSetAccelPath(m, accelPath) }
func (m *MenuItem) SetLabel(label string)         { MenuItemSetLabel(m, label) }
func (m *MenuItem) SetRightJustified(rightJustified bool) {
	MenuItemSetRightJustified(m, rightJustified)
}
func (m *MenuItem) SetSubmenu(submenu *Widget)         { MenuItemSetSubmenu(m, submenu) }
func (m *MenuItem) SetUseUnderline(setting bool)       { MenuItemSetUseUnderline(m, setting) }
func (m *MenuItem) ToggleSizeAllocate(allocation int)  { MenuItemToggleSizeAllocate(m, allocation) }
func (m *MenuItem) ToggleSizeRequest(requisition *int) { MenuItemToggleSizeRequest(m, requisition) }

type MenuShell struct {
	Container       Container
	Children        *T.GList
	ActiveMenuItem  *Widget
	ParentMenuShell *Widget
	Button          uint
	ActivateTime    T.GUint32
	Bits            uint
	// Active : 1
	// HaveGrab : 1
	// HaveXgrab : 1
	// IgnoreLeave : 1
	// MenuFlag : 1
	// IgnoreEnter : 1
	// KeyboardMode : 1
}

var (
	MenuShellGetType func() O.Type

	MenuShellActivateItem func(m *MenuShell, menuItem *Widget, forceDeactivate bool)
	MenuShellAppend       func(m *MenuShell, child *Widget)
	MenuShellCancel       func(m *MenuShell)
	MenuShellDeactivate   func(m *MenuShell)
	MenuShellDeselect     func(m *MenuShell)
	MenuShellGetTakeFocus func(m *MenuShell) bool
	MenuShellInsert       func(m *MenuShell, child *Widget, position int)
	MenuShellPrepend      func(m *MenuShell, child *Widget)
	MenuShellSelectFirst  func(m *MenuShell, searchSensitive bool)
	MenuShellSelectItem   func(m *MenuShell, menuItem *Widget)
	MenuShellSetTakeFocus func(m *MenuShell, takeFocus bool)
)

func (m *MenuShell) ActivateItem(menuItem *Widget, forceDeactivate bool) {
	MenuShellActivateItem(m, menuItem, forceDeactivate)
}
func (m *MenuShell) Append(child *Widget)               { MenuShellAppend(m, child) }
func (m *MenuShell) Cancel()                            { MenuShellCancel(m) }
func (m *MenuShell) Deactivate()                        { MenuShellDeactivate(m) }
func (m *MenuShell) Deselect()                          { MenuShellDeselect(m) }
func (m *MenuShell) GetTakeFocus() bool                 { return MenuShellGetTakeFocus(m) }
func (m *MenuShell) Insert(child *Widget, position int) { MenuShellInsert(m, child, position) }
func (m *MenuShell) Prepend(child *Widget)              { MenuShellPrepend(m, child) }
func (m *MenuShell) SelectFirst(searchSensitive bool)   { MenuShellSelectFirst(m, searchSensitive) }
func (m *MenuShell) SelectItem(menuItem *Widget)        { MenuShellSelectItem(m, menuItem) }
func (m *MenuShell) SetTakeFocus(takeFocus bool)        { MenuShellSetTakeFocus(m, takeFocus) }

type MenuToolButton struct {
	Parent ToolButton
	_      *struct{}
}

var (
	MenuToolButtonGetType      func() O.Type
	MenuToolButtonNew          func(iconWidget *Widget, label string) *ToolItem
	MenuToolButtonNewFromStock func(stockId string) *ToolItem

	MenuToolButtonGetMenu               func(t *MenuToolButton) *Widget
	MenuToolButtonSetArrowTooltip       func(t *MenuToolButton, tooltips *Tooltips, tipText, tipPrivate string)
	MenuToolButtonSetArrowTooltipMarkup func(t *MenuToolButton, markup string)
	MenuToolButtonSetArrowTooltipText   func(t *MenuToolButton, text string)
	MenuToolButtonSetMenu               func(t *MenuToolButton, menu *Widget)
)

func (t *MenuToolButton) GetMenu() *Widget { return MenuToolButtonGetMenu(t) }
func (t *MenuToolButton) SetArrowTooltip(tooltips *Tooltips, tipText, tipPrivate string) {
	MenuToolButtonSetArrowTooltip(t, tooltips, tipText, tipPrivate)
}
func (t *MenuToolButton) SetArrowTooltipMarkup(markup string) {
	MenuToolButtonSetArrowTooltipMarkup(t, markup)
}
func (t *MenuToolButton) SetArrowTooltipText(text string) { MenuToolButtonSetArrowTooltipText(t, text) }
func (t *MenuToolButton) SetMenu(menu *Widget)            { MenuToolButtonSetMenu(t, menu) }

type MessageType Enum

const (
	MESSAGE_INFO MessageType = iota
	MESSAGE_WARNING
	MESSAGE_QUESTION
	MESSAGE_ERROR
	MESSAGE_OTHER
)

var MessageTypeGetType func() O.Type

type Misc struct {
	Widget Widget
	Xalign float32
	Yalign float32
	Xpad   uint16
	Ypad   uint16
}

var (
	MiscGetType func() O.Type

	MiscGetAlignment func(m *Misc, xalign, yalign *float32)
	MiscGetPadding   func(m *Misc, xpad, ypad *int)
	MiscSetAlignment func(m *Misc, xalign, yalign float32)
	MiscSetPadding   func(m *Misc, xpad, ypad int)
)

func (m *Misc) GetAlignment(xalign, yalign *float32) { MiscGetAlignment(m, xalign, yalign) }
func (m *Misc) GetPadding(xpad, ypad *int)           { MiscGetPadding(m, xpad, ypad) }
func (m *Misc) SetAlignment(xalign, yalign float32)  { MiscSetAlignment(m, xalign, yalign) }
func (m *Misc) SetPadding(xpad, ypad int)            { MiscSetPadding(m, xpad, ypad) }

type MiscClass struct {
	ParentClass WidgetClass
}

type MessageDialog struct {
	Parent Dialog
	Image  *Widget
	Label  *Widget
}

var (
	MessageDialogGetType       func() O.Type
	MessageDialogNew           func(parent *Window, flags DialogFlags, t MessageType, buttons ButtonsType, messageFormat string, v ...VArg) *Widget
	MessageDialogNewWithMarkup func(parent *Window, flags DialogFlags, t MessageType, buttons ButtonsType, messageFormat string, v ...VArg) *Widget

	MessageDialogFormatSecondaryMarkup func(m *MessageDialog, messageFormat string, v ...VArg)
	MessageDialogFormatSecondaryText   func(m *MessageDialog, messageFormat string, v ...VArg)
	MessageDialogGetImage              func(m *MessageDialog) *Widget
	MessageDialogGetMessageArea        func(m *MessageDialog) *Widget
	MessageDialogSetImage              func(m *MessageDialog, image *Widget)
	MessageDialogSetMarkup             func(m *MessageDialog, str string)
)

func (m *MessageDialog) FormatSecondaryMarkup(messageFormat string, v ...VArg) {
	MessageDialogFormatSecondaryMarkup(m, messageFormat, v)
}
func (m *MessageDialog) FormatSecondaryText(messageFormat string, v ...VArg) {
	MessageDialogFormatSecondaryText(m, messageFormat, v)
}
func (m *MessageDialog) GetImage() *Widget       { return MessageDialogGetImage(m) }
func (m *MessageDialog) GetMessageArea() *Widget { return MessageDialogGetMessageArea(m) }
func (m *MessageDialog) SetImage(image *Widget)  { MessageDialogSetImage(m, image) }
func (m *MessageDialog) SetMarkup(str string)    { MessageDialogSetMarkup(m, str) }

type MetricType Enum

const (
	PIXELS MetricType = iota
	INCHES
	CENTIMETERS
)

var MetricTypeGetType func() O.Type

type MountOperation struct {
	Parent I.MountOperation
	_      *struct{}
}

var (
	MountOperationGetType func() O.Type
	MountOperationNew     func(parent *Window) *MountOperation

	MountOperationGetParent func(m *MountOperation) *Window
	MountOperationGetScreen func(m *MountOperation) *D.Screen
	MountOperationIsShowing func(m *MountOperation) bool
	MountOperationSetParent func(m *MountOperation, parent *Window)
	MountOperationSetScreen func(m *MountOperation, screen *D.Screen)
)

func (m *MountOperation) GetParent() *Window         { return MountOperationGetParent(m) }
func (m *MountOperation) GetScreen() *D.Screen       { return MountOperationGetScreen(m) }
func (m *MountOperation) IsShowing() bool            { return MountOperationIsShowing(m) }
func (m *MountOperation) SetParent(parent *Window)   { MountOperationSetParent(m, parent) }
func (m *MountOperation) SetScreen(screen *D.Screen) { MountOperationSetScreen(m, screen) }

type MovementStep Enum

const (
	MOVEMENT_LOGICAL_POSITIONS MovementStep = iota
	MOVEMENT_VISUAL_POSITIONS
	MOVEMENT_WORDS
	MOVEMENT_DISPLAY_LINES
	MOVEMENT_DISPLAY_LINE_ENDS
	MOVEMENT_PARAGRAPHS
	MOVEMENT_PARAGRAPH_ENDS
	MOVEMENT_PAGES
	MOVEMENT_BUFFER_ENDS
	MOVEMENT_HORIZONTAL_PAGES
)

var MovementStepGetType func() O.Type
