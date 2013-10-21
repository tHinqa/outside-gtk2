// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Dialog struct {
	Window      T.GtkWindow
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
	DialogNewWithButtons func(title string, parent *T.GtkWindow, flags DialogFlags, firstButtonText string, v ...VArg) *Widget
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
