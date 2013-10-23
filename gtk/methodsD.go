// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type DebugFlag T.Enum

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

var DebugFlagGetType func() T.GType

type DeleteType T.Enum

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

var DeleteTypeGetType func() T.GType

type DestDefaults T.Enum

const (
	DEST_DEFAULT_MOTION DestDefaults = 1 << iota
	DEST_DEFAULT_HIGHLIGHT
	DEST_DEFAULT_DROP
	DEST_DEFAULT_ALL DestDefaults = 0x07
)

var DestDefaultsGetType func() T.GType

type Dialog struct {
	Window      Window
	Vbox        *Widget
	Action_area *Widget
	Separator   *Widget
}

type DialogFlags T.Enum

const (
	DIALOG_MODAL DialogFlags = 1 << iota
	DIALOG_DESTROY_WITH_PARENT
	DIALOG_NO_SEPARATOR
)

var (
	DialogFlagsGetType   func() T.GType
	DialogGetType        func() T.GType
	DialogNew            func() *Widget
	DialogNewWithButtons func(title string, parent *Window, flags DialogFlags, firstButtonText string, v ...VArg) *Widget
)

var (
	dialogAddActionWidget                    func(d *Dialog, child *Widget, responseId int)
	dialogAddButton                          func(d *Dialog, buttonText string, responseId int) *Widget
	dialogAddButtons                         func(d *Dialog, firstButtonText string, v ...VArg)
	dialogGetActionArea                      func(d *Dialog) *Widget
	dialogGetContentArea                     func(d *Dialog) *Widget
	dialogGetHasSeparator                    func(d *Dialog) T.Gboolean
	dialogGetResponseForWidget               func(d *Dialog, widget *Widget) int
	dialogGetWidgetForResponse               func(d *Dialog, responseId int) *Widget
	dialogResponse                           func(d *Dialog, responseId int)
	dialogRun                                func(d *Dialog) int
	dialogSetAlternativeButtonOrder          func(d *Dialog, firstResponseId int, v ...VArg)
	dialogSetAlternativeButtonOrderFromArray func(d *Dialog, nParams int, newOrder *int)
	dialogSetDefaultResponse                 func(d *Dialog, responseId int)
	dialogSetHasSeparator                    func(d *Dialog, setting T.Gboolean)
	dialogSetResponseSensitive               func(d *Dialog, responseId int, setting T.Gboolean)
)

func (d *Dialog) AddActionWidget(child *Widget, responseId int) {
	dialogAddActionWidget(d, child, responseId)
}
func (d *Dialog) AddButton(buttonText string, responseId int) *Widget {
	return dialogAddButton(d, buttonText, responseId)
}
func (d *Dialog) AddButtons(firstButtonText string, v ...VArg) {
	dialogAddButtons(d, firstButtonText, v)
}
func (d *Dialog) GetActionArea() *Widget      { return dialogGetActionArea(d) }
func (d *Dialog) GetContentArea() *Widget     { return dialogGetContentArea(d) }
func (d *Dialog) GetHasSeparator() T.Gboolean { return dialogGetHasSeparator(d) }
func (d *Dialog) GetResponseForWidget(widget *Widget) int {
	return dialogGetResponseForWidget(d, widget)
}
func (d *Dialog) GetWidgetForResponse(responseId int) *Widget {
	return dialogGetWidgetForResponse(d, responseId)
}
func (d *Dialog) Response(responseId int) { dialogResponse(d, responseId) }
func (d *Dialog) Run() int                { return dialogRun(d) }
func (d *Dialog) SetAlternativeButtonOrder(firstResponseId int, v ...VArg) {
	dialogSetAlternativeButtonOrder(d, firstResponseId, v)
}
func (d *Dialog) SetAlternativeButtonOrderFromArray(nParams int, newOrder *int) {
	dialogSetAlternativeButtonOrderFromArray(d, nParams, newOrder)
}
func (d *Dialog) SetDefaultResponse(responseId int)  { dialogSetDefaultResponse(d, responseId) }
func (d *Dialog) SetHasSeparator(setting T.Gboolean) { dialogSetHasSeparator(d, setting) }
func (d *Dialog) SetResponseSensitive(responseId int, setting T.Gboolean) {
	dialogSetResponseSensitive(d, responseId, setting)
}

type DirectionType T.Enum

const (
	DIR_TAB_FORWARD DirectionType = iota
	DIR_TAB_BACKWARD
	DIR_UP
	DIR_DOWN
	DIR_LEFT
	DIR_RIGHT
)

var DirectionTypeGetType func() T.GType

var DisableSetlocale func()

var (
	DragBegin                 func(widget *Widget, targets *TargetList, actions T.GdkDragAction, button int, event *T.GdkEvent) *T.GdkDragContext
	DragCheckThreshold        func(widget *Widget, startX, startY, currentX, currentY int) T.Gboolean
	DragDestAddImageTargets   func(widget *Widget)
	DragDestAddTextTargets    func(widget *Widget)
	DragDestAddUriTargets     func(widget *Widget)
	DragDestFindTarget        func(widget *Widget, context *T.GdkDragContext, targetList *TargetList) T.GdkAtom
	DragDestGetTargetList     func(widget *Widget) *TargetList
	DragDestGetTrackMotion    func(widget *Widget) T.Gboolean
	DragDestSet               func(widget *Widget, flags DestDefaults, targets *TargetEntry, nTargets int, actions T.GdkDragAction)
	DragDestSetProxy          func(widget *Widget, proxyWindow *T.GdkWindow, protocol T.GdkDragProtocol, useCoordinates T.Gboolean)
	DragDestSetTargetList     func(widget *Widget, targetList *TargetList)
	DragDestSetTrackMotion    func(widget *Widget, trackMotion T.Gboolean)
	DragDestUnset             func(widget *Widget)
	DragFinish                func(context *T.GdkDragContext, success, del T.Gboolean, time T.GUint32)
	DragGetData               func(widget *Widget, context *T.GdkDragContext, target T.GdkAtom, time T.GUint32)
	DragGetSourceWidget       func(context *T.GdkDragContext) *Widget
	DragHighlight             func(widget *Widget)
	DragSetDefaultIcon        func(colormap *T.GdkColormap, pixmap *T.GdkPixmap, mask *T.GdkBitmap, hotX, hotY int)
	DragSetIconDefault        func(context *T.GdkDragContext)
	DragSetIconName           func(context *T.GdkDragContext, iconName string, hotX, hotY int)
	DragSetIconPixbuf         func(context *T.GdkDragContext, pixbuf *T.GdkPixbuf, hotX, hotY int)
	DragSetIconPixmap         func(context *T.GdkDragContext, colormap *T.GdkColormap, pixmap *T.GdkPixmap, mask *T.GdkBitmap, hotX, hotY int)
	DragSetIconStock          func(context *T.GdkDragContext, stockId string, hotX, hotY int)
	DragSetIconWidget         func(context *T.GdkDragContext, widget *Widget, hotX, hotY int)
	DragSourceAddImageTargets func(widget *Widget)
	DragSourceAddTextTargets  func(widget *Widget)
	DragSourceAddUriTargets   func(widget *Widget)
	DragSourceGetTargetList   func(widget *Widget) *TargetList
	DragSourceSet             func(widget *Widget, startButtonMask T.GdkModifierType, targets *TargetEntry, nTargets int, actions T.GdkDragAction)
	DragSourceSetIcon         func(widget *Widget, colormap *T.GdkColormap, pixmap *T.GdkPixmap, mask *T.GdkBitmap)
	DragSourceSetIconName     func(widget *Widget, iconName string)
	DragSourceSetIconPixbuf   func(widget *Widget, pixbuf *T.GdkPixbuf)
	DragSourceSetIconStock    func(widget *Widget, stockId string)
	DragSourceSetTargetList   func(widget *Widget, targetList *TargetList)
	DragSourceUnset           func(widget *Widget)
	DragUnhighlight           func(widget *Widget)
)

type DragResult T.Enum

const (
	DRAG_RESULT_SUCCESS DragResult = iota
	DRAG_RESULT_NO_TARGET
	DRAG_RESULT_USER_CANCELLED
	DRAG_RESULT_TIMEOUT_EXPIRED
	DRAG_RESULT_GRAB_BROKEN
	DRAG_RESULT_ERROR
)

var DragResultGetType func() T.GType

type DrawingArea struct {
	Widget    Widget
	Draw_data T.Gpointer
}

var (
	DrawingAreaGetType func() T.GType
	DrawingAreaNew     func() *Widget

	drawingAreaSize func(d *DrawingArea, width int, height int)
)

func (d *DrawingArea) Size(width int, height int) { drawingAreaSize(d, width, height) }
