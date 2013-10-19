package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Dialog struct {
	Window      T.GtkWindow
	Vbox        *T.GtkWidget
	Action_area *T.GtkWidget
	Separator   *T.GtkWidget
}

var (
	DialogFlagsGetType   func() T.GType
	DialogGetType        func() T.GType
	DialogNew            func() *T.GtkWidget
	DialogNewWithButtons func(title string, parent *T.GtkWindow, flags T.GtkDialogFlags, firstButtonText string, v ...VArg) *T.GtkWidget
)

var (
	DialogAddActionWidget                    func(d *Dialog, child *T.GtkWidget, responseId int)
	DialogAddButton                          func(d *Dialog, buttonText string, responseId int) *T.GtkWidget
	DialogAddButtons                         func(d *Dialog, firstButtonText string, v ...VArg)
	DialogGetActionArea                      func(d *Dialog) *T.GtkWidget
	DialogGetContentArea                     func(d *Dialog) *T.GtkWidget
	DialogGetHasSeparator                    func(d *Dialog) T.Gboolean
	DialogGetResponseForWidget               func(d *Dialog, widget *T.GtkWidget) int
	DialogGetWidgetForResponse               func(d *Dialog, responseId int) *T.GtkWidget
	DialogResponse                           func(d *Dialog, responseId int)
	DialogRun                                func(d *Dialog) int
	DialogSetAlternativeButtonOrder          func(d *Dialog, firstResponseId int, v ...VArg)
	DialogSetAlternativeButtonOrderFromArray func(d *Dialog, nParams int, newOrder *int)
	DialogSetDefaultResponse                 func(d *Dialog, responseId int)
	DialogSetHasSeparator                    func(d *Dialog, setting T.Gboolean)
	DialogSetResponseSensitive               func(d *Dialog, responseId int, setting T.Gboolean)
)

func (d *Dialog) AddActionWidget(child *T.GtkWidget, responseId int) {
	DialogAddActionWidget(d, child, responseId)
}

func (d *Dialog) AddButton(buttonText string, responseId int) *T.GtkWidget {
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

func (d *Dialog) GetWidgetForResponse(responseId int) *T.GtkWidget {
	return DialogGetWidgetForResponse(d, responseId)
}

func (d *Dialog) GetResponseForWidget(widget *T.GtkWidget) int {
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

func (d *Dialog) GetActionArea() *T.GtkWidget {
	return DialogGetActionArea(d)
}

func (d *Dialog) GetContentArea() *T.GtkWidget {
	return DialogGetContentArea(d)
}

type DrawingArea struct {
	Widget    T.GtkWidget
	Draw_data T.Gpointer
}

var (
	DrawingAreaGetType func() T.GType
	DrawingAreaNew     func() *T.GtkWidget

	DrawingAreaSize func(d *DrawingArea, width int, height int)
)

func (d *DrawingArea) Size(width int, height int) {
	DrawingAreaSize(d, width, height)
}
