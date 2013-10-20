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
	DialogAddActionWidget                    func(d *Dialog, child *Widget, responseId int)
	DialogAddButton                          func(d *Dialog, buttonText string, responseId int) *Widget
	DialogAddButtons                         func(d *Dialog, firstButtonText string, v ...VArg)
	DialogGetActionArea                      func(d *Dialog) *Widget
	DialogGetContentArea                     func(d *Dialog) *Widget
	DialogGetHasSeparator                    func(d *Dialog) T.Gboolean
	DialogGetResponseForWidget               func(d *Dialog, widget *Widget) int
	DialogGetWidgetForResponse               func(d *Dialog, responseId int) *Widget
	DialogResponse                           func(d *Dialog, responseId int)
	DialogRun                                func(d *Dialog) int
	DialogSetAlternativeButtonOrder          func(d *Dialog, firstResponseId int, v ...VArg)
	DialogSetAlternativeButtonOrderFromArray func(d *Dialog, nParams int, newOrder *int)
	DialogSetDefaultResponse                 func(d *Dialog, responseId int)
	DialogSetHasSeparator                    func(d *Dialog, setting T.Gboolean)
	DialogSetResponseSensitive               func(d *Dialog, responseId int, setting T.Gboolean)
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

func (d *Dialog) SetResponseSensitive(responseId int, setting T.Gboolean) {
	DialogSetResponseSensitive(d, responseId, setting)
}

func (d *Dialog) SetDefaultResponse(responseId int) {
	DialogSetDefaultResponse(d, responseId)
}

func (d *Dialog) GetWidgetForResponse(responseId int) *Widget {
	return DialogGetWidgetForResponse(d, responseId)
}

func (d *Dialog) GetResponseForWidget(widget *Widget) int {
	return DialogGetResponseForWidget(d, widget)
}

func (d *Dialog) SetHasSeparator(setting T.Gboolean) {
	DialogSetHasSeparator(d, setting)
}

func (d *Dialog) GetHasSeparator() T.Gboolean {
	return DialogGetHasSeparator(d)
}

func (d *Dialog) SetAlternativeButtonOrder(firstResponseId int, v ...VArg) {
	DialogSetAlternativeButtonOrder(d, firstResponseId, v)
}

func (d *Dialog) SetAlternativeButtonOrderFromArray(nParams int, newOrder *int) {
	DialogSetAlternativeButtonOrderFromArray(d, nParams, newOrder)
}

func (d *Dialog) Response(responseId int) {
	DialogResponse(d, responseId)
}

func (d *Dialog) Run() int { return DialogRun(d) }

func (d *Dialog) GetActionArea() *Widget {
	return DialogGetActionArea(d)
}

func (d *Dialog) GetContentArea() *Widget {
	return DialogGetContentArea(d)
}

type DrawingArea struct {
	Widget    Widget
	Draw_data T.Gpointer
}

var (
	DrawingAreaGetType func() T.GType
	DrawingAreaNew     func() *Widget

	DrawingAreaSize func(d *DrawingArea, width int, height int)
)

func (d *DrawingArea) Size(width int, height int) {
	DrawingAreaSize(d, width, height)
}
