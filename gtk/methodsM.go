// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type MatchType T.Enum

const (
	MATCH_ALL MatchType = iota
	MATCH_ALL_TAIL
	MATCH_HEAD
	MATCH_TAIL
	MATCH_EXACT
	MATCH_LAST
)

var MatchTypeGetType func() T.GType

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
		NavigationRegion  *T.GdkRegion
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
		x, y *int, pushIn *T.Gboolean, userData T.Gpointer)

	MenuDetachFunc func(ttachWidget *Widget, menu *Menu)
)

type MenuDirectionType T.Enum

const (
	MENU_DIR_PARENT MenuDirectionType = iota
	ENU_DIR_CHILD
	MENU_DIR_NEXT
	MENU_DIR_PREV
)

var (
	MenuGetType func() T.GType
	MenuNew     func() *Widget

	MenuDirectionTypeGetType func() T.GType
	MenuGetForAttachWidget   func(widget *Widget) *T.GList

	menuAttach               func(m *Menu, child *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint)
	menuAttachToWidget       func(m *Menu, attachWidget *Widget, detacher MenuDetachFunc)
	menuDetach               func(m *Menu)
	menuGetAccelGroup        func(m *Menu) *AccelGroup
	menuGetAccelPath         func(m *Menu) string
	menuGetActive            func(m *Menu) *Widget
	menuGetAttachWidget      func(m *Menu) *Widget
	menuGetMonitor           func(m *Menu) int
	menuGetReserveToggleSize func(m *Menu) T.Gboolean
	menuGetTearoffState      func(m *Menu) T.Gboolean
	menuGetTitle             func(m *Menu) string
	menuPopdown              func(m *Menu)
	menuPopup                func(m *Menu, parentMenuShell, parentMenuItem *Widget, f MenuPositionFunc, data T.Gpointer, button uint, activateTime T.GUint32)
	menuReorderChild         func(m *Menu, child *Widget, position int)
	menuReposition           func(m *Menu)
	menuSetAccelGroup        func(m *Menu, accelGroup *AccelGroup)
	menuSetAccelPath         func(m *Menu, accelPath string)
	menuSetActive            func(m *Menu, index uint)
	menuSetMonitor           func(m *Menu, monitorNum int)
	menuSetReserveToggleSize func(m *Menu, reserveToggleSize T.Gboolean)
	menuSetScreen            func(m *Menu, screen *D.Screen)
	menuSetTearoffState      func(m *Menu, tornOff T.Gboolean)
	menuSetTitle             func(m *Menu, title string)
)

func (m *Menu) Attach(child *Widget, leftAttach, rightAttach, topAttach, bottomAttach uint) {
	menuAttach(m, child, leftAttach, rightAttach, topAttach, bottomAttach)
}
func (m *Menu) AttachToWidget(attachWidget *Widget, detacher MenuDetachFunc) {
	menuAttachToWidget(m, attachWidget, detacher)
}
func (m *Menu) Detach()                          { menuDetach(m) }
func (m *Menu) GetAccelGroup() *AccelGroup       { return menuGetAccelGroup(m) }
func (m *Menu) GetAccelPath() string             { return menuGetAccelPath(m) }
func (m *Menu) GetActive() *Widget               { return menuGetActive(m) }
func (m *Menu) GetAttachWidget() *Widget         { return menuGetAttachWidget(m) }
func (m *Menu) GetMonitor() int                  { return menuGetMonitor(m) }
func (m *Menu) GetReserveToggleSize() T.Gboolean { return menuGetReserveToggleSize(m) }
func (m *Menu) GetTearoffState() T.Gboolean      { return menuGetTearoffState(m) }
func (m *Menu) GetTitle() string                 { return menuGetTitle(m) }
func (m *Menu) Popdown()                         { menuPopdown(m) }
func (m *Menu) Popup(parentMenuShell, parentMenuItem *Widget, f MenuPositionFunc, data T.Gpointer, button uint, activateTime T.GUint32) {
	menuPopup(m, parentMenuShell, parentMenuItem, f, data, button, activateTime)
}
func (m *Menu) ReorderChild(child *Widget, position int) { menuReorderChild(m, child, position) }
func (m *Menu) Reposition()                              { menuReposition(m) }
func (m *Menu) SetAccelGroup(accelGroup *AccelGroup)     { menuSetAccelGroup(m, accelGroup) }
func (m *Menu) SetAccelPath(accelPath string)            { menuSetAccelPath(m, accelPath) }
func (m *Menu) SetActive(index uint)                     { menuSetActive(m, index) }
func (m *Menu) SetMonitor(monitorNum int)                { menuSetMonitor(m, monitorNum) }
func (m *Menu) SetReserveToggleSize(reserveToggleSize T.Gboolean) {
	menuSetReserveToggleSize(m, reserveToggleSize)
}
func (m *Menu) SetScreen(screen *D.Screen)         { menuSetScreen(m, screen) }
func (m *Menu) SetTearoffState(tornOff T.Gboolean) { menuSetTearoffState(m, tornOff) }
func (m *Menu) SetTitle(title string)              { menuSetTitle(m, title) }

type MenuBar struct {
	MenuShell MenuShell
}

var (
	MenuBarGetType func() T.GType
	MenuBarNew     func() *Widget

	menuBarGetChildPackDirection func(m *MenuBar) PackDirection
	menuBarGetPackDirection      func(m *MenuBar) PackDirection
	menuBarSetChildPackDirection func(m *MenuBar, childPackDir PackDirection)
	menuBarSetPackDirection      func(m *MenuBar, packDir PackDirection)
)

func (m *MenuBar) GetChildPackDirection() PackDirection { return menuBarGetChildPackDirection(m) }
func (m *MenuBar) GetPackDirection() PackDirection      { return menuBarGetPackDirection(m) }
func (m *MenuBar) SetChildPackDirection(childPackDir PackDirection) {
	menuBarSetChildPackDirection(m, childPackDir)
}
func (m *MenuBar) SetPackDirection(packDir PackDirection) { menuBarSetPackDirection(m, packDir) }

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
	MenuItemGetType         func() T.GType
	MenuItemNew             func() *Widget
	MenuItemNewWithLabel    func(label string) *Widget
	MenuItemNewWithMnemonic func(label string) *Widget

	menuItemActivate           func(m *MenuItem)
	menuItemDeselect           func(m *MenuItem)
	menuItemGetAccelPath       func(m *MenuItem) string
	menuItemGetLabel           func(m *MenuItem) string
	menuItemGetRightJustified  func(m *MenuItem) T.Gboolean
	menuItemGetSubmenu         func(m *MenuItem) *Widget
	menuItemGetUseUnderline    func(m *MenuItem) T.Gboolean
	menuItemRemoveSubmenu      func(m *MenuItem)
	menuItemSelect             func(m *MenuItem)
	menuItemSetAccelPath       func(m *MenuItem, accelPath string)
	menuItemSetLabel           func(m *MenuItem, label string)
	menuItemSetRightJustified  func(m *MenuItem, rightJustified T.Gboolean)
	menuItemSetSubmenu         func(m *MenuItem, submenu *Widget)
	menuItemSetUseUnderline    func(m *MenuItem, setting T.Gboolean)
	menuItemToggleSizeAllocate func(m *MenuItem, allocation int)
	menuItemToggleSizeRequest  func(m *MenuItem, requisition *int)
)

func (m *MenuItem) Activate()                     { menuItemActivate(m) }
func (m *MenuItem) Deselect()                     { menuItemDeselect(m) }
func (m *MenuItem) GetAccelPath() string          { return menuItemGetAccelPath(m) }
func (m *MenuItem) GetLabel() string              { return menuItemGetLabel(m) }
func (m *MenuItem) GetRightJustified() T.Gboolean { return menuItemGetRightJustified(m) }
func (m *MenuItem) GetSubmenu() *Widget           { return menuItemGetSubmenu(m) }
func (m *MenuItem) GetUseUnderline() T.Gboolean   { return menuItemGetUseUnderline(m) }
func (m *MenuItem) RemoveSubmenu()                { menuItemRemoveSubmenu(m) }
func (m *MenuItem) Select()                       { menuItemSelect(m) }
func (m *MenuItem) SetAccelPath(accelPath string) { menuItemSetAccelPath(m, accelPath) }
func (m *MenuItem) SetLabel(label string)         { menuItemSetLabel(m, label) }
func (m *MenuItem) SetRightJustified(rightJustified T.Gboolean) {
	menuItemSetRightJustified(m, rightJustified)
}
func (m *MenuItem) SetSubmenu(submenu *Widget)         { menuItemSetSubmenu(m, submenu) }
func (m *MenuItem) SetUseUnderline(setting T.Gboolean) { menuItemSetUseUnderline(m, setting) }
func (m *MenuItem) ToggleSizeAllocate(allocation int)  { menuItemToggleSizeAllocate(m, allocation) }
func (m *MenuItem) ToggleSizeRequest(requisition *int) { menuItemToggleSizeRequest(m, requisition) }

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
	MenuShellGetType func() T.GType

	menuShellActivateItem func(m *MenuShell, menuItem *Widget, forceDeactivate T.Gboolean)
	menuShellAppend       func(m *MenuShell, child *Widget)
	menuShellCancel       func(m *MenuShell)
	menuShellDeactivate   func(m *MenuShell)
	menuShellDeselect     func(m *MenuShell)
	menuShellGetTakeFocus func(m *MenuShell) T.Gboolean
	menuShellInsert       func(m *MenuShell, child *Widget, position int)
	menuShellPrepend      func(m *MenuShell, child *Widget)
	menuShellSelectFirst  func(m *MenuShell, searchSensitive T.Gboolean)
	menuShellSelectItem   func(m *MenuShell, menuItem *Widget)
	menuShellSetTakeFocus func(m *MenuShell, takeFocus T.Gboolean)
)

func (m *MenuShell) ActivateItem(menuItem *Widget, forceDeactivate T.Gboolean) {
	menuShellActivateItem(m, menuItem, forceDeactivate)
}
func (m *MenuShell) Append(child *Widget)                   { menuShellAppend(m, child) }
func (m *MenuShell) Cancel()                                { menuShellCancel(m) }
func (m *MenuShell) Deactivate()                            { menuShellDeactivate(m) }
func (m *MenuShell) Deselect()                              { menuShellDeselect(m) }
func (m *MenuShell) GetTakeFocus() T.Gboolean               { return menuShellGetTakeFocus(m) }
func (m *MenuShell) Insert(child *Widget, position int)     { menuShellInsert(m, child, position) }
func (m *MenuShell) Prepend(child *Widget)                  { menuShellPrepend(m, child) }
func (m *MenuShell) SelectFirst(searchSensitive T.Gboolean) { menuShellSelectFirst(m, searchSensitive) }
func (m *MenuShell) SelectItem(menuItem *Widget)            { menuShellSelectItem(m, menuItem) }
func (m *MenuShell) SetTakeFocus(takeFocus T.Gboolean)      { menuShellSetTakeFocus(m, takeFocus) }

type MenuToolButton struct {
	Parent ToolButton
	_      *struct{}
}

var (
	MenuToolButtonGetType      func() T.GType
	MenuToolButtonNew          func(iconWidget *Widget, label string) *ToolItem
	MenuToolButtonNewFromStock func(stockId string) *ToolItem

	menuToolButtonGetMenu               func(t *MenuToolButton) *Widget
	menuToolButtonSetArrowTooltip       func(t *MenuToolButton, tooltips *Tooltips, tipText, tipPrivate string)
	menuToolButtonSetArrowTooltipMarkup func(t *MenuToolButton, markup string)
	menuToolButtonSetArrowTooltipText   func(t *MenuToolButton, text string)
	menuToolButtonSetMenu               func(t *MenuToolButton, menu *Widget)
)

func (t *MenuToolButton) GetMenu() *Widget { return menuToolButtonGetMenu(t) }
func (t *MenuToolButton) SetArrowTooltip(tooltips *Tooltips, tipText, tipPrivate string) {
	menuToolButtonSetArrowTooltip(t, tooltips, tipText, tipPrivate)
}
func (t *MenuToolButton) SetArrowTooltipMarkup(markup string) {
	menuToolButtonSetArrowTooltipMarkup(t, markup)
}
func (t *MenuToolButton) SetArrowTooltipText(text string) { menuToolButtonSetArrowTooltipText(t, text) }
func (t *MenuToolButton) SetMenu(menu *Widget)            { menuToolButtonSetMenu(t, menu) }

type MessageType T.Enum

const (
	MESSAGE_INFO MessageType = iota
	MESSAGE_WARNING
	MESSAGE_QUESTION
	MESSAGE_ERROR
	MESSAGE_OTHER
)

var MessageTypeGetType func() T.GType

type Misc struct {
	Widget Widget
	Xalign float32
	Yalign float32
	Xpad   uint16
	Ypad   uint16
}

var (
	MiscGetType func() T.GType

	miscGetAlignment func(m *Misc, xalign, yalign *float32)
	miscGetPadding   func(m *Misc, xpad, ypad *int)
	miscSetAlignment func(m *Misc, xalign, yalign float32)
	miscSetPadding   func(m *Misc, xpad, ypad int)
)

func (m *Misc) GetAlignment(xalign, yalign *float32) { miscGetAlignment(m, xalign, yalign) }
func (m *Misc) GetPadding(xpad, ypad *int)           { miscGetPadding(m, xpad, ypad) }
func (m *Misc) SetAlignment(xalign, yalign float32)  { miscSetAlignment(m, xalign, yalign) }
func (m *Misc) SetPadding(xpad, ypad int)            { miscSetPadding(m, xpad, ypad) }

type MiscClass struct {
	ParentClass WidgetClass
}

type MessageDialog struct {
	Parent Dialog
	Image  *Widget
	Label  *Widget
}

var (
	MessageDialogGetType       func() T.GType
	MessageDialogNew           func(parent *Window, flags DialogFlags, t MessageType, buttons ButtonsType, messageFormat string, v ...VArg) *Widget
	MessageDialogNewWithMarkup func(parent *Window, flags DialogFlags, t MessageType, buttons ButtonsType, messageFormat string, v ...VArg) *Widget

	messageDialogFormatSecondaryMarkup func(m *MessageDialog, messageFormat string, v ...VArg)
	messageDialogFormatSecondaryText   func(m *MessageDialog, messageFormat string, v ...VArg)
	messageDialogGetImage              func(m *MessageDialog) *Widget
	messageDialogGetMessageArea        func(m *MessageDialog) *Widget
	messageDialogSetImage              func(m *MessageDialog, image *Widget)
	messageDialogSetMarkup             func(m *MessageDialog, str string)
)

func (m *MessageDialog) FormatSecondaryMarkup(messageFormat string, v ...VArg) {
	messageDialogFormatSecondaryMarkup(m, messageFormat, v)
}
func (m *MessageDialog) FormatSecondaryText(messageFormat string, v ...VArg) {
	messageDialogFormatSecondaryText(m, messageFormat, v)
}
func (m *MessageDialog) GetImage() *Widget       { return messageDialogGetImage(m) }
func (m *MessageDialog) GetMessageArea() *Widget { return messageDialogGetMessageArea(m) }
func (m *MessageDialog) SetImage(image *Widget)  { messageDialogSetImage(m, image) }
func (m *MessageDialog) SetMarkup(str string)    { messageDialogSetMarkup(m, str) }

type MetricType T.Enum

const (
	PIXELS MetricType = iota
	INCHES
	CENTIMETERS
)

var MetricTypeGetType func() T.GType

type MountOperation struct {
	Parent T.GMountOperation
	_      *struct{}
}

var (
	MountOperationGetType func() T.GType
	MountOperationNew     func(parent *Window) *MountOperation

	mountOperationGetParent func(m *MountOperation) *Window
	mountOperationGetScreen func(m *MountOperation) *D.Screen
	mountOperationIsShowing func(m *MountOperation) T.Gboolean
	mountOperationSetParent func(m *MountOperation, parent *Window)
	mountOperationSetScreen func(m *MountOperation, screen *D.Screen)
)

func (m *MountOperation) GetParent() *Window         { return mountOperationGetParent(m) }
func (m *MountOperation) GetScreen() *D.Screen       { return mountOperationGetScreen(m) }
func (m *MountOperation) IsShowing() T.Gboolean      { return mountOperationIsShowing(m) }
func (m *MountOperation) SetParent(parent *Window)   { mountOperationSetParent(m, parent) }
func (m *MountOperation) SetScreen(screen *D.Screen) { mountOperationSetScreen(m, screen) }

type MovementStep T.Enum

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

var MovementStepGetType func() T.GType
