// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type DebugFlag Enum

const (
	DEBUG_MISC DebugFlag = 1 << iota
	DEBUG_PLUGSOCKET
	DEBUG_TEXT
	DEBUG_TREE
	DEBUG_UPDATES
	DEBUG_KEYBINDINGS
	DEBUG_MULTIHEAD
	DEBUG_MODULES
	DEBUG_GEOMETRY
	DEBUG_ICONTHEME
	DEBUG_PRINTING
	DEBUG_BUILDER
)

var DebugFlagGetType func() O.Type

type DeleteType Enum

const (
	DELETE_CHARS DeleteType = iota
	DELETE_WORD_ENDS
	DELETE_WORDS
	DELETE_DISPLAY_LINES
	DELETE_DISPLAY_LINE_ENDS
	DELETE_PARAGRAPH_ENDS
	DELETE_PARAGRAPHS
	DELETE_WHITESPACE
)

var DeleteTypeGetType func() O.Type

type DestDefaults Enum

const (
	DEST_DEFAULT_MOTION DestDefaults = 1 << iota
	DEST_DEFAULT_HIGHLIGHT
	DEST_DEFAULT_DROP
	DEST_DEFAULT_ALL DestDefaults = 0x07
)

var DestDefaultsGetType func() O.Type

type Dialog struct {
	Window      Window
	Vbox        *Widget
	Action_area *Widget
	Separator   *Widget
}

type DialogFlags Enum

const (
	DIALOG_MODAL DialogFlags = 1 << iota
	DIALOG_DESTROY_WITH_PARENT
	DIALOG_NO_SEPARATOR
)

var (
	DialogFlagsGetType   func() O.Type
	DialogGetType        func() O.Type
	DialogNew            func() *Widget
	DialogNewWithButtons func(title string, parent *Window, flags DialogFlags, firstButtonText string, v ...VArg) *Widget
)

var (
	DialogAddActionWidget                    func(d *Dialog, child *Widget, responseId int)
	DialogAddButton                          func(d *Dialog, buttonText string, responseId int) *Widget
	DialogAddButtons                         func(d *Dialog, firstButtonText string, v ...VArg)
	DialogGetActionArea                      func(d *Dialog) *Widget
	DialogGetContentArea                     func(d *Dialog) *Widget
	DialogGetHasSeparator                    func(d *Dialog) bool
	DialogGetResponseForWidget               func(d *Dialog, widget *Widget) int
	DialogGetWidgetForResponse               func(d *Dialog, responseId int) *Widget
	DialogResponse                           func(d *Dialog, responseId int)
	DialogRun                                func(d *Dialog) int
	DialogSetAlternativeButtonOrder          func(d *Dialog, firstResponseId int, v ...VArg)
	DialogSetAlternativeButtonOrderFromArray func(d *Dialog, nParams int, newOrder *int)
	DialogSetDefaultResponse                 func(d *Dialog, responseId int)
	DialogSetHasSeparator                    func(d *Dialog, setting bool)
	DialogSetResponseSensitive               func(d *Dialog, responseId int, setting bool)
)

func (d *Dialog) AddActionWidget(child *Widget, responseId int) {
	DialogAddActionWidget(d, child, responseId)
}
func (d *Dialog) AddButton(buttonText string, responseId int) *Widget {
	return DialogAddButton(d, buttonText, responseId)
}
func (d *Dialog) AddButtons(firstButtonText string, v ...VArg) {
	DialogAddButtons(d, firstButtonText, v)
}
func (d *Dialog) GetActionArea() *Widget  { return DialogGetActionArea(d) }
func (d *Dialog) GetContentArea() *Widget { return DialogGetContentArea(d) }
func (d *Dialog) GetHasSeparator() bool   { return DialogGetHasSeparator(d) }
func (d *Dialog) GetResponseForWidget(widget *Widget) int {
	return DialogGetResponseForWidget(d, widget)
}
func (d *Dialog) GetWidgetForResponse(responseId int) *Widget {
	return DialogGetWidgetForResponse(d, responseId)
}
func (d *Dialog) Response(responseId int) { DialogResponse(d, responseId) }
func (d *Dialog) Run() int                { return DialogRun(d) }
func (d *Dialog) SetAlternativeButtonOrder(firstResponseId int, v ...VArg) {
	DialogSetAlternativeButtonOrder(d, firstResponseId, v)
}
func (d *Dialog) SetAlternativeButtonOrderFromArray(nParams int, newOrder *int) {
	DialogSetAlternativeButtonOrderFromArray(d, nParams, newOrder)
}
func (d *Dialog) SetDefaultResponse(responseId int) { DialogSetDefaultResponse(d, responseId) }
func (d *Dialog) SetHasSeparator(setting bool)      { DialogSetHasSeparator(d, setting) }
func (d *Dialog) SetResponseSensitive(responseId int, setting bool) {
	DialogSetResponseSensitive(d, responseId, setting)
}

type DirectionType Enum

const (
	DIR_TAB_FORWARD DirectionType = iota
	DIR_TAB_BACKWARD
	DIR_UP
	DIR_DOWN
	DIR_LEFT
	DIR_RIGHT
)

var DirectionTypeGetType func() O.Type

var DisableSetlocale func()

var (
	DragBegin                 func(widget *Widget, targets *TargetList, actions D.DragAction, button int, event *D.Event) *D.DragContext
	DragCheckThreshold        func(widget *Widget, startX, startY, currentX, currentY int) bool
	DragDestAddImageTargets   func(widget *Widget)
	DragDestAddTextTargets    func(widget *Widget)
	DragDestAddUriTargets     func(widget *Widget)
	DragDestFindTarget        func(widget *Widget, context *D.DragContext, targetList *TargetList) D.Atom
	DragDestGetTargetList     func(widget *Widget) *TargetList
	DragDestGetTrackMotion    func(widget *Widget) bool
	DragDestSet               func(widget *Widget, flags DestDefaults, targets *TargetEntry, nTargets int, actions D.DragAction)
	DragDestSetProxy          func(widget *Widget, proxyWindow *D.Window, protocol D.DragProtocol, useCoordinates bool)
	DragDestSetTargetList     func(widget *Widget, targetList *TargetList)
	DragDestSetTrackMotion    func(widget *Widget, trackMotion bool)
	DragDestUnset             func(widget *Widget)
	DragFinish                func(context *D.DragContext, success, del bool, time T.GUint32)
	DragGetData               func(widget *Widget, context *D.DragContext, target D.Atom, time T.GUint32)
	DragGetSourceWidget       func(context *D.DragContext) *Widget
	DragHighlight             func(widget *Widget)
	DragSetDefaultIcon        func(colormap *D.Colormap, pixmap *D.Pixmap, mask *T.GdkBitmap, hotX, hotY int)
	DragSetIconDefault        func(context *D.DragContext)
	DragSetIconName           func(context *D.DragContext, iconName string, hotX, hotY int)
	DragSetIconPixbuf         func(context *D.DragContext, pixbuf *D.Pixbuf, hotX, hotY int)
	DragSetIconPixmap         func(context *D.DragContext, colormap *D.Colormap, pixmap *D.Pixmap, mask *T.GdkBitmap, hotX, hotY int)
	DragSetIconStock          func(context *D.DragContext, stockId string, hotX, hotY int)
	DragSetIconWidget         func(context *D.DragContext, widget *Widget, hotX, hotY int)
	DragSourceAddImageTargets func(widget *Widget)
	DragSourceAddTextTargets  func(widget *Widget)
	DragSourceAddUriTargets   func(widget *Widget)
	DragSourceGetTargetList   func(widget *Widget) *TargetList
	DragSourceSet             func(widget *Widget, startButtonMask T.GdkModifierType, targets *TargetEntry, nTargets int, actions D.DragAction)
	DragSourceSetIcon         func(widget *Widget, colormap *D.Colormap, pixmap *D.Pixmap, mask *T.GdkBitmap)
	DragSourceSetIconName     func(widget *Widget, iconName string)
	DragSourceSetIconPixbuf   func(widget *Widget, pixbuf *D.Pixbuf)
	DragSourceSetIconStock    func(widget *Widget, stockId string)
	DragSourceSetTargetList   func(widget *Widget, targetList *TargetList)
	DragSourceUnset           func(widget *Widget)
	DragUnhighlight           func(widget *Widget)
)

type DragResult Enum

const (
	DRAG_RESULT_SUCCESS DragResult = iota
	DRAG_RESULT_NO_TARGET
	DRAG_RESULT_USER_CANCELLED
	DRAG_RESULT_TIMEOUT_EXPIRED
	DRAG_RESULT_GRAB_BROKEN
	DRAG_RESULT_ERROR
)

var DragResultGetType func() O.Type

type DrawingArea struct {
	Widget    Widget
	Draw_data T.Gpointer
}

var (
	DrawingAreaGetType func() O.Type
	DrawingAreaNew     func() *Widget

	DrawingAreaSize func(d *DrawingArea, width int, height int)
)

func (d *DrawingArea) Size(width int, height int) { DrawingAreaSize(d, width, height) }
