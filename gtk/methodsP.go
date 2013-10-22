// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type PageSetup struct{}

var (
	PageSetupGetType        func() T.GType
	PageSetupNew            func() *T.GtkPageSetup
	PageSetupNewFromFile    func(fileName string, error **T.GError) *T.GtkPageSetup
	PageSetupNewFromKeyFile func(keyFile *T.GKeyFile, groupName string, error **T.GError) *T.GtkPageSetup

	pageSetupCopy                          func(p *PageSetup) *T.GtkPageSetup
	pageSetupGetBottomMargin               func(p *PageSetup, unit T.GtkUnit) float64
	pageSetupGetLeftMargin                 func(p *PageSetup, unit T.GtkUnit) float64
	pageSetupGetOrientation                func(p *PageSetup) T.GtkPageOrientation
	pageSetupGetPageHeight                 func(p *PageSetup, unit T.GtkUnit) float64
	pageSetupGetPageWidth                  func(p *PageSetup, unit T.GtkUnit) float64
	pageSetupGetPaperHeight                func(p *PageSetup, unit T.GtkUnit) float64
	pageSetupGetPaperSize                  func(p *PageSetup) *PaperSize
	pageSetupGetPaperWidth                 func(p *PageSetup, unit T.GtkUnit) float64
	pageSetupGetRightMargin                func(p *PageSetup, unit T.GtkUnit) float64
	pageSetupGetTopMargin                  func(p *PageSetup, unit T.GtkUnit) float64
	pageSetupLoadFile                      func(p *PageSetup, fileName string, error **T.GError) T.Gboolean
	pageSetupLoadKeyFile                   func(p *PageSetup, keyFile *T.GKeyFile, groupName string, error **T.GError) T.Gboolean
	pageSetupSetBottomMargin               func(p *PageSetup, margin float64, unit T.GtkUnit)
	pageSetupSetLeftMargin                 func(p *PageSetup, margin float64, unit T.GtkUnit)
	pageSetupSetOrientation                func(p *PageSetup, orientation T.GtkPageOrientation)
	pageSetupSetPaperSize                  func(p *PageSetup, size *PaperSize)
	pageSetupSetPaperSizeAndDefaultMargins func(p *PageSetup, size *PaperSize)
	pageSetupSetRightMargin                func(p *PageSetup, margin float64, unit T.GtkUnit)
	pageSetupSetTopMargin                  func(p *PageSetup, margin float64, unit T.GtkUnit)
	pageSetupToFile                        func(p *PageSetup, fileName string, error **T.GError) T.Gboolean
	pageSetupToKeyFile                     func(p *PageSetup, keyFile *T.GKeyFile, groupName string)
)

func (p *PageSetup) Copy() *T.GtkPageSetup                  { return pageSetupCopy(p) }
func (p *PageSetup) GetBottomMargin(unit T.GtkUnit) float64 { return pageSetupGetBottomMargin(p, unit) }
func (p *PageSetup) GetLeftMargin(unit T.GtkUnit) float64   { return pageSetupGetLeftMargin(p, unit) }
func (p *PageSetup) GetOrientation() T.GtkPageOrientation   { return pageSetupGetOrientation(p) }
func (p *PageSetup) GetPageHeight(unit T.GtkUnit) float64   { return pageSetupGetPageHeight(p, unit) }
func (p *PageSetup) GetPageWidth(unit T.GtkUnit) float64    { return pageSetupGetPageWidth(p, unit) }
func (p *PageSetup) GetPaperHeight(unit T.GtkUnit) float64  { return pageSetupGetPaperHeight(p, unit) }
func (p *PageSetup) GetPaperSize() *PaperSize               { return pageSetupGetPaperSize(p) }
func (p *PageSetup) GetPaperWidth(unit T.GtkUnit) float64   { return pageSetupGetPaperWidth(p, unit) }
func (p *PageSetup) GetRightMargin(unit T.GtkUnit) float64  { return pageSetupGetRightMargin(p, unit) }
func (p *PageSetup) GetTopMargin(unit T.GtkUnit) float64    { return pageSetupGetTopMargin(p, unit) }
func (p *PageSetup) LoadFile(fileName string, err **T.GError) T.Gboolean {
	return pageSetupLoadFile(p, fileName, err)
}
func (p *PageSetup) LoadKeyFile(keyFile *T.GKeyFile, groupName string, err **T.GError) T.Gboolean {
	return pageSetupLoadKeyFile(p, keyFile, groupName, err)
}
func (p *PageSetup) SetBottomMargin(margin float64, unit T.GtkUnit) {
	pageSetupSetBottomMargin(p, margin, unit)
}
func (p *PageSetup) SetLeftMargin(margin float64, unit T.GtkUnit) {
	pageSetupSetLeftMargin(p, margin, unit)
}
func (p *PageSetup) SetOrientation(orientation T.GtkPageOrientation) {
	pageSetupSetOrientation(p, orientation)
}
func (p *PageSetup) SetPaperSize(size *PaperSize) { pageSetupSetPaperSize(p, size) }
func (p *PageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	pageSetupSetPaperSizeAndDefaultMargins(p, size)
}
func (p *PageSetup) SetRightMargin(margin float64, unit T.GtkUnit) {
	pageSetupSetRightMargin(p, margin, unit)
}
func (p *PageSetup) SetTopMargin(margin float64, unit T.GtkUnit) {
	pageSetupSetTopMargin(p, margin, unit)
}
func (p *PageSetup) ToFile(fileName string, err **T.GError) T.Gboolean {
	return pageSetupToFile(p, fileName, err)
}
func (p *PageSetup) ToKeyFile(keyFile *T.GKeyFile, groupName string) {
	pageSetupToKeyFile(p, keyFile, groupName)
}

type Paned struct {
	Container      Container
	Child1         *Widget
	Child2         *Widget
	Handle         *T.GdkWindow
	XorGc          *T.GdkGC
	CursorType     T.GdkCursorType
	HandlePos      T.GdkRectangle
	Child1Size     int
	LastAllocation int
	MinPosition    int
	MaxPosition    int
	Bits           uint
	// PositionSet : 1
	// InDrag : 1
	// Child1Shrink : 1
	// Child1Resize : 1
	// Child2Shrink : 1
	// Child2Resize : 1
	// Orientation : 1
	// InRecursion : 1
	// HandlePrelit : 1
	LastChild1Focus  *Widget
	LastChild2Focus  *Widget
	_                *struct{}
	DragPos          int
	OriginalPosition int
}

var (
	PanedGetType func() T.GType

	panedAdd1            func(p *Paned, child *Widget)
	panedAdd2            func(p *Paned, child *Widget)
	panedComputePosition func(p *Paned, allocation, child1Req, child2Req int)
	panedGetChild1       func(p *Paned) *Widget
	panedGetChild2       func(p *Paned) *Widget
	panedGetHandleWindow func(p *Paned) *T.GdkWindow
	panedGetPosition     func(p *Paned) int
	panedPack1           func(p *Paned, child *Widget, resize, shrink T.Gboolean)
	panedPack2           func(p *Paned, child *Widget, resize, shrink T.Gboolean)
	panedSetPosition     func(p *Paned, position int)
)

func (p *Paned) Add1(child *Widget) { panedAdd1(p, child) }
func (p *Paned) Add2(child *Widget) { panedAdd2(p, child) }
func (p *Paned) ComputePosition(allocation, child1Req, child2Req int) {
	panedComputePosition(p, allocation, child1Req, child2Req)
}
func (p *Paned) GetChild1() *Widget            { return panedGetChild1(p) }
func (p *Paned) GetChild2() *Widget            { return panedGetChild2(p) }
func (p *Paned) GetHandleWindow() *T.GdkWindow { return panedGetHandleWindow(p) }
func (p *Paned) GetPosition() int              { return panedGetPosition(p) }
func (p *Paned) Pack1(child *Widget, resize T.Gboolean, shrink T.Gboolean) {
	panedPack1(p, child, resize, shrink)
}
func (p *Paned) Pack2(child *Widget, resize T.Gboolean, shrink T.Gboolean) {
	panedPack2(p, child, resize, shrink)
}
func (p *Paned) SetPosition(position int) { panedSetPosition(p, position) }

type PaperSize struct{}

var (
	PaperSizeGetType        func() T.GType
	PaperSizeNew            func(name string) *PaperSize
	PaperSizeNewCustom      func(name, displayName string, width, height float64, unit T.GtkUnit) *PaperSize
	PaperSizeNewFromKeyFile func(keyFile *T.GKeyFile, groupName string, err **T.GError) *PaperSize
	PaperSizeNewFromPpd     func(ppdName, ppdDisplayName string, width, height float64) *PaperSize

	PaperSizeGetDefault    func() string
	PaperSizeGetPaperSizes func(includeCustom T.Gboolean) *T.GList

	paperSizeCopy                   func(s *PaperSize) *PaperSize
	paperSizeFree                   func(s *PaperSize)
	paperSizeGetDefaultBottomMargin func(s *PaperSize, unit T.GtkUnit) float64
	paperSizeGetDefaultLeftMargin   func(s *PaperSize, unit T.GtkUnit) float64
	paperSizeGetDefaultRightMargin  func(s *PaperSize, unit T.GtkUnit) float64
	paperSizeGetDefaultTopMargin    func(s *PaperSize, unit T.GtkUnit) float64
	paperSizeGetDisplayName         func(s *PaperSize) string
	paperSizeGetHeight              func(s *PaperSize, unit T.GtkUnit) float64
	paperSizeGetName                func(s *PaperSize) string
	paperSizeGetPpdName             func(s *PaperSize) string
	paperSizeGetWidth               func(s *PaperSize, unit T.GtkUnit) float64
	paperSizeIsCustom               func(s *PaperSize) T.Gboolean
	paperSizeIsEqual                func(s *PaperSize, size2 *PaperSize) T.Gboolean
	paperSizeSetSize                func(s *PaperSize, width, height float64, unit T.GtkUnit)
	paperSizeToKeyFile              func(s *PaperSize, keyFile *T.GKeyFile, groupName string)
)

func (s *PaperSize) Copy() *PaperSize { return paperSizeCopy(s) }
func (s *PaperSize) Free()            { paperSizeFree(s) }
func (s *PaperSize) GetDefaultBottomMargin(unit T.GtkUnit) float64 {
	return paperSizeGetDefaultBottomMargin(s, unit)
}
func (s *PaperSize) GetDefaultLeftMargin(unit T.GtkUnit) float64 {
	return paperSizeGetDefaultLeftMargin(s, unit)
}
func (s *PaperSize) GetDefaultRightMargin(unit T.GtkUnit) float64 {
	return paperSizeGetDefaultRightMargin(s, unit)
}
func (s *PaperSize) GetDefaultTopMargin(unit T.GtkUnit) float64 {
	return paperSizeGetDefaultTopMargin(s, unit)
}
func (s *PaperSize) GetDisplayName() string              { return paperSizeGetDisplayName(s) }
func (s *PaperSize) GetHeight(unit T.GtkUnit) float64    { return paperSizeGetHeight(s, unit) }
func (s *PaperSize) GetName() string                     { return paperSizeGetName(s) }
func (s *PaperSize) GetPpdName() string                  { return paperSizeGetPpdName(s) }
func (s *PaperSize) GetWidth(unit T.GtkUnit) float64     { return paperSizeGetWidth(s, unit) }
func (s *PaperSize) IsCustom() T.Gboolean                { return paperSizeIsCustom(s) }
func (s *PaperSize) IsEqual(size2 *PaperSize) T.Gboolean { return paperSizeIsEqual(s, size2) }
func (s *PaperSize) SetSize(width, height float64, unit T.GtkUnit) {
	paperSizeSetSize(s, width, height, unit)
}
func (s *PaperSize) ToKeyFile(keyFile *T.GKeyFile, groupName string) {
	paperSizeToKeyFile(s, keyFile, groupName)
}

type Pixmap struct {
	Misc              Misc
	Pixmap            *T.GdkPixmap
	Mask              *T.GdkBitmap
	PixmapInsensitive *T.GdkPixmap
	BuildInsensitive  uint // : 1
}

var (
	PixmapGetType func() T.GType
	PixmapNew     func(pixmap *T.GdkPixmap, mask *T.GdkBitmap) *Widget

	pixmapGet                 func(p *Pixmap, val **T.GdkPixmap, mask **T.GdkBitmap)
	pixmapSet                 func(p *Pixmap, val *T.GdkPixmap, mask *T.GdkBitmap)
	pixmapSetBuildInsensitive func(p *Pixmap, build T.Gboolean)
)

func (p *Pixmap) Get(val **T.GdkPixmap, mask **T.GdkBitmap) { pixmapGet(p, val, mask) }
func (p *Pixmap) Set(val *T.GdkPixmap, mask *T.GdkBitmap)   { pixmapSet(p, val, mask) }
func (p *Pixmap) SetBuildInsensitive(build T.Gboolean)      { pixmapSetBuildInsensitive(p, build) }

type Plug struct {
	Window         Window
	SocketWindow   *T.GdkWindow
	ModalityWindow *Widget
	ModalityGroup  *WindowGroup
	GrabbedKeys    *T.GHashTable
	SameApp        uint // : 1
}

var (
	PlugGetType       func() T.GType
	PlugNew           func(socketId T.GdkNativeWindow) *Widget
	PlugNewForDisplay func(display *T.GdkDisplay, socketId T.GdkNativeWindow) *Widget

	plugConstruct           func(p *Plug, socketId T.GdkNativeWindow)
	plugConstructForDisplay func(p *Plug, display *T.GdkDisplay, socketId T.GdkNativeWindow)
	plugGetEmbedded         func(p *Plug) T.Gboolean
	plugGetId               func(p *Plug) T.GdkNativeWindow
	plugGetSocketWindow     func(p *Plug) *T.GdkWindow
)

func (p *Plug) Construct(socketId T.GdkNativeWindow) { plugConstruct(p, socketId) }
func (p *Plug) ConstructForDisplay(display *T.GdkDisplay, socketId T.GdkNativeWindow) {
	plugConstructForDisplay(p, display, socketId)
}
func (p *Plug) GetEmbedded() T.Gboolean       { return plugGetEmbedded(p) }
func (p *Plug) GetId() T.GdkNativeWindow      { return plugGetId(p) }
func (p *Plug) GetSocketWindow() *T.GdkWindow { return plugGetSocketWindow(p) }

type Preview struct {
	Widget       Widget
	Buffer       *T.Guchar
	BufferWidth  uint16
	BufferHeight uint16
	Bpp          uint16
	Rowstride    uint16
	Dither       T.GdkRgbDither
	Bits         uint
	// Type : 1
	// Expand : 1
}

type PreviewType T.Enum

const (
	PREVIEW_COLOR PreviewType = iota
	PREVIEW_GRAYSCALE
)

var (
	PreviewGetType     func() T.GType
	PreviewNew         func(t PreviewType) *Widget
	PreviewTypeGetType func() T.GType

	PreviewGetCmap        func() *T.GdkColormap
	PreviewGetInfo        func() *T.GtkPreviewInfo
	PreviewGetVisual      func() *T.GdkVisual
	PreviewReset          func()
	PreviewSetColorCube   func(nredShades uint, ngreenShades, nblueShades, ngrayShades uint)
	PreviewSetGamma       func(gamma float64)
	PreviewSetInstallCmap func(installCmap int)
	PreviewSetReserved    func(nreserved int)
	PreviewUninit         func()

	previewDrawRow   func(p *Preview, data *T.Guchar, x, y, w int)
	previewPut       func(p *Preview, window *T.GdkWindow, gc *T.GdkGC, srcx, srcy, destx, desty, width, height int)
	previewSetDither func(p *Preview, dither T.GdkRgbDither)
	previewSetExpand func(p *Preview, expand T.Gboolean)
	previewSize      func(p *Preview, width, height int)
)

func (p *Preview) DrawRow(data *T.Guchar, x, y, w int) { previewDrawRow(p, data, x, y, w) }
func (p *Preview) Put(window *T.GdkWindow, gc *T.GdkGC, srcx, srcy, destx, desty, width, height int) {
	previewPut(p, window, gc, srcx, srcy, destx, desty, width, height)
}
func (p *Preview) SetDither(dither T.GdkRgbDither) { previewSetDither(p, dither) }
func (p *Preview) SetExpand(expand T.Gboolean)     { previewSetExpand(p, expand) }
func (p *Preview) Size(width, height int)          { previewSize(p, width, height) }

type PrintContext struct{}

var (
	PrintContextGetType func() T.GType

	printContextCreatePangoContext func(p *PrintContext) *T.PangoContext
	printContextCreatePangoLayout  func(p *PrintContext) *T.PangoLayout
	printContextGetCairoContext    func(p *PrintContext) *T.Cairo
	printContextGetDpiX            func(p *PrintContext) float64
	printContextGetDpiY            func(p *PrintContext) float64
	printContextGetHardMargins     func(p *PrintContext, top, bottom, left, right *float64) T.Gboolean
	printContextGetHeight          func(p *PrintContext) float64
	printContextGetPageSetup       func(p *PrintContext) *T.GtkPageSetup
	printContextGetPangoFontmap    func(p *PrintContext) *T.PangoFontMap
	printContextGetWidth           func(p *PrintContext) float64
	printContextSetCairoContext    func(p *PrintContext, cr *T.Cairo, dpiX, dpiY float64)
)

func (p *PrintContext) CreatePangoContext() *T.PangoContext { return printContextCreatePangoContext(p) }
func (p *PrintContext) CreatePangoLayout() *T.PangoLayout   { return printContextCreatePangoLayout(p) }
func (p *PrintContext) GetCairoContext() *T.Cairo           { return printContextGetCairoContext(p) }
func (p *PrintContext) GetDpiX() float64                    { return printContextGetDpiX(p) }
func (p *PrintContext) GetDpiY() float64                    { return printContextGetDpiY(p) }
func (p *PrintContext) GetHardMargins(top, bottom, left, right *float64) T.Gboolean {
	return printContextGetHardMargins(p, top, bottom, left, right)
}
func (p *PrintContext) GetHeight() float64               { return printContextGetHeight(p) }
func (p *PrintContext) GetPageSetup() *T.GtkPageSetup    { return printContextGetPageSetup(p) }
func (p *PrintContext) GetPangoFontmap() *T.PangoFontMap { return printContextGetPangoFontmap(p) }
func (p *PrintContext) GetWidth() float64                { return printContextGetWidth(p) }
func (p *PrintContext) SetCairoContext(cr *T.Cairo, dpiX, dpiY float64) {
	printContextSetCairoContext(p, cr, dpiX, dpiY)
}

type PrintOperation struct {
	Parent_instance T.GObject
	_               *struct{}
}

var (
	PrintOperationGetType       func() T.GType
	PrintOperationNew           func() *PrintOperation
	PrintOperationActionGetType func() T.GType
	PrintOperationResultGetType func() T.GType

	printOperationCancel              func(p *PrintOperation)
	printOperationDrawPageFinish      func(p *PrintOperation)
	printOperationGetDefaultPageSetup func(p *PrintOperation) *T.GtkPageSetup
	printOperationGetEmbedPageSetup   func(p *PrintOperation) T.Gboolean
	printOperationGetError            func(p *PrintOperation, err **T.GError)
	printOperationGetHasSelection     func(p *PrintOperation) T.Gboolean
	printOperationGetNPagesToPrint    func(p *PrintOperation) int
	printOperationGetPrintSettings    func(p *PrintOperation) *PrintSettings
	printOperationGetStatus           func(p *PrintOperation) T.GtkPrintStatus
	printOperationGetStatusString     func(p *PrintOperation) string
	printOperationGetSupportSelection func(p *PrintOperation) T.Gboolean
	printOperationIsFinished          func(p *PrintOperation) T.Gboolean
	printOperationRun                 func(p *PrintOperation, action T.GtkPrintOperationAction, parent *Window, err **T.GError) T.GtkPrintOperationResult
	printOperationSetAllowAsync       func(p *PrintOperation, allowAsync T.Gboolean)
	printOperationSetCurrentPage      func(p *PrintOperation, currentPage int)
	printOperationSetCustomTabLabel   func(p *PrintOperation, label string)
	printOperationSetDefaultPageSetup func(p *PrintOperation, defaultPageSetup *T.GtkPageSetup)
	printOperationSetDeferDrawing     func(p *PrintOperation)
	printOperationSetEmbedPageSetup   func(p *PrintOperation, embed T.Gboolean)
	printOperationSetExportFilename   func(p *PrintOperation, filename string)
	printOperationSetHasSelection     func(p *PrintOperation, hasSelection T.Gboolean)
	printOperationSetJobName          func(p *PrintOperation, jobName string)
	printOperationSetNPages           func(p *PrintOperation, nPages int)
	printOperationSetPrintSettings    func(p *PrintOperation, printSettings *PrintSettings)
	printOperationSetShowProgress     func(p *PrintOperation, showProgress T.Gboolean)
	printOperationSetSupportSelection func(p *PrintOperation, supportSelection T.Gboolean)
	printOperationSetTrackPrintStatus func(p *PrintOperation, trackStatus T.Gboolean)
	printOperationSetUnit             func(p *PrintOperation, unit T.GtkUnit)
	printOperationSetUseFullPage      func(p *PrintOperation, fullPage T.Gboolean)
)

func (p *PrintOperation) Cancel()         { printOperationCancel(p) }
func (p *PrintOperation) DrawPageFinish() { printOperationDrawPageFinish(p) }
func (p *PrintOperation) GetDefaultPageSetup() *T.GtkPageSetup {
	return printOperationGetDefaultPageSetup(p)
}
func (p *PrintOperation) GetEmbedPageSetup() T.Gboolean    { return printOperationGetEmbedPageSetup(p) }
func (p *PrintOperation) GetError(err **T.GError)          { printOperationGetError(p, err) }
func (p *PrintOperation) GetHasSelection() T.Gboolean      { return printOperationGetHasSelection(p) }
func (p *PrintOperation) GetNPagesToPrint() int            { return printOperationGetNPagesToPrint(p) }
func (p *PrintOperation) GetPrintSettings() *PrintSettings { return printOperationGetPrintSettings(p) }
func (p *PrintOperation) GetStatus() T.GtkPrintStatus      { return printOperationGetStatus(p) }
func (p *PrintOperation) GetStatusString() string          { return printOperationGetStatusString(p) }
func (p *PrintOperation) GetSupportSelection() T.Gboolean  { return printOperationGetSupportSelection(p) }
func (p *PrintOperation) IsFinished() T.Gboolean           { return printOperationIsFinished(p) }
func (p *PrintOperation) Run(action T.GtkPrintOperationAction, parent *Window, err **T.GError) T.GtkPrintOperationResult {
	return printOperationRun(p, action, parent, err)
}
func (p *PrintOperation) SetAllowAsync(allowAsync T.Gboolean) {
	printOperationSetAllowAsync(p, allowAsync)
}
func (p *PrintOperation) SetCurrentPage(currentPage int) {
	printOperationSetCurrentPage(p, currentPage)
}
func (p *PrintOperation) SetCustomTabLabel(label string) { printOperationSetCustomTabLabel(p, label) }
func (p *PrintOperation) SetDefaultPageSetup(defaultPageSetup *T.GtkPageSetup) {
	printOperationSetDefaultPageSetup(p, defaultPageSetup)
}
func (p *PrintOperation) SetDeferDrawing() { printOperationSetDeferDrawing(p) }
func (p *PrintOperation) SetEmbedPageSetup(embed T.Gboolean) {
	printOperationSetEmbedPageSetup(p, embed)
}
func (p *PrintOperation) SetExportFilename(filename string) {
	printOperationSetExportFilename(p, filename)
}
func (p *PrintOperation) SetHasSelection(hasSelection T.Gboolean) {
	printOperationSetHasSelection(p, hasSelection)
}
func (p *PrintOperation) SetJobName(jobName string) { printOperationSetJobName(p, jobName) }
func (p *PrintOperation) SetNPages(nPages int)      { printOperationSetNPages(p, nPages) }
func (p *PrintOperation) SetPrintSettings(printSettings *PrintSettings) {
	printOperationSetPrintSettings(p, printSettings)
}
func (p *PrintOperation) SetShowProgress(showProgress T.Gboolean) {
	printOperationSetShowProgress(p, showProgress)
}
func (p *PrintOperation) SetSupportSelection(supportSelection T.Gboolean) {
	printOperationSetSupportSelection(p, supportSelection)
}
func (p *PrintOperation) SetTrackPrintStatus(trackStatus T.Gboolean) {
	printOperationSetTrackPrintStatus(p, trackStatus)
}
func (p *PrintOperation) SetUnit(unit T.GtkUnit) { printOperationSetUnit(p, unit) }
func (p *PrintOperation) SetUseFullPage(fullPage T.Gboolean) {
	printOperationSetUseFullPage(p, fullPage)
}

type PrintOperationPreview struct{}

var (
	PrintOperationPreviewGetType func() T.GType

	printOperationPreviewEndPreview func(p *PrintOperationPreview)
	printOperationPreviewIsSelected func(p *PrintOperationPreview, pageNr int) T.Gboolean
	printOperationPreviewRenderPage func(p *PrintOperationPreview, pageNr int)
)

func (p *PrintOperationPreview) EndPreview() { printOperationPreviewEndPreview(p) }
func (p *PrintOperationPreview) IsSelected(pageNr int) T.Gboolean {
	return printOperationPreviewIsSelected(p, pageNr)
}
func (p *PrintOperationPreview) RenderPage(pageNr int) { printOperationPreviewRenderPage(p, pageNr) }

type (
	PrintSettings struct{}

	PrintSettingsFunc func(key, value string, userData T.Gpointer)
)

var (
	PrintSettingsGetType        func() T.GType
	PrintSettingsNew            func() *PrintSettings
	PrintSettingsNewFromFile    func(fileName string, err **T.GError) *PrintSettings
	PrintSettingsNewFromKeyFile func(keyFile *T.GKeyFile, groupName string, err **T.GError) *PrintSettings

	printSettingsCopy                 func(p *PrintSettings) *PrintSettings
	printSettingsForeach              func(p *PrintSettings, f PrintSettingsFunc, userData T.Gpointer)
	printSettingsGet                  func(p *PrintSettings, key string) string
	printSettingsGetBool              func(p *PrintSettings, key string) T.Gboolean
	printSettingsGetCollate           func(p *PrintSettings) T.Gboolean
	printSettingsGetDefaultSource     func(p *PrintSettings) string
	printSettingsGetDither            func(p *PrintSettings) string
	printSettingsGetDouble            func(p *PrintSettings, key string) float64
	printSettingsGetDoubleWithDefault func(p *PrintSettings, key string, def float64) float64
	printSettingsGetDuplex            func(p *PrintSettings) T.GtkPrintDuplex
	printSettingsGetFinishings        func(p *PrintSettings) string
	printSettingsGetInt               func(p *PrintSettings, key string) int
	printSettingsGetIntWithDefault    func(p *PrintSettings, key string, def int) int
	printSettingsGetLength            func(p *PrintSettings, key string, unit T.GtkUnit) float64
	printSettingsGetMediaType         func(p *PrintSettings) string
	printSettingsGetNCopies           func(p *PrintSettings) int
	printSettingsGetNumberUp          func(p *PrintSettings) int
	printSettingsGetNumberUpLayout    func(p *PrintSettings) T.GtkNumberUpLayout
	printSettingsGetOrientation       func(p *PrintSettings) T.GtkPageOrientation
	printSettingsGetOutputBin         func(p *PrintSettings) string
	printSettingsGetPageRanges        func(p *PrintSettings, numRanges *int) *T.GtkPageRange
	printSettingsGetPageSet           func(p *PrintSettings) T.GtkPageSet
	printSettingsGetPaperHeight       func(p *PrintSettings, unit T.GtkUnit) float64
	printSettingsGetPaperSize         func(p *PrintSettings) *PaperSize
	printSettingsGetPaperWidth        func(p *PrintSettings, unit T.GtkUnit) float64
	printSettingsGetPrinter           func(p *PrintSettings) string
	printSettingsGetPrinterLpi        func(p *PrintSettings) float64
	printSettingsGetPrintPages        func(p *PrintSettings) T.GtkPrintPages
	printSettingsGetQuality           func(p *PrintSettings) T.GtkPrintQuality
	printSettingsGetResolution        func(p *PrintSettings) int
	printSettingsGetResolutionX       func(p *PrintSettings) int
	printSettingsGetResolutionY       func(p *PrintSettings) int
	printSettingsGetReverse           func(p *PrintSettings) T.Gboolean
	printSettingsGetScale             func(p *PrintSettings) float64
	printSettingsGetUseColor          func(p *PrintSettings) T.Gboolean
	printSettingsHasKey               func(p *PrintSettings, key string) T.Gboolean
	printSettingsLoadFile             func(p *PrintSettings, fileName string, err **T.GError) T.Gboolean
	printSettingsLoadKeyFile          func(p *PrintSettings, keyFile *T.GKeyFile, groupName string, err **T.GError) T.Gboolean
	printSettingsSet                  func(p *PrintSettings, key, value string)
	printSettingsSetBool              func(p *PrintSettings, key string, value T.Gboolean)
	printSettingsSetCollate           func(p *PrintSettings, collate T.Gboolean)
	printSettingsSetDefaultSource     func(p *PrintSettings, defaultSource string)
	printSettingsSetDither            func(p *PrintSettings, dither string)
	printSettingsSetDouble            func(p *PrintSettings, key string, value float64)
	printSettingsSetDuplex            func(p *PrintSettings, duplex T.GtkPrintDuplex)
	printSettingsSetFinishings        func(p *PrintSettings, finishings string)
	printSettingsSetInt               func(p *PrintSettings, key string, value int)
	printSettingsSetLength            func(p *PrintSettings, key string, value float64, unit T.GtkUnit)
	printSettingsSetMediaType         func(p *PrintSettings, mediaType string)
	printSettingsSetNCopies           func(p *PrintSettings, numCopies int)
	printSettingsSetNumberUp          func(p *PrintSettings, numberUp int)
	printSettingsSetNumberUpLayout    func(p *PrintSettings, numberUpLayout T.GtkNumberUpLayout)
	printSettingsSetOrientation       func(p *PrintSettings, orientation T.GtkPageOrientation)
	printSettingsSetOutputBin         func(p *PrintSettings, outputBin string)
	printSettingsSetPageRanges        func(p *PrintSettings, pageRanges *T.GtkPageRange, numRanges int)
	printSettingsSetPageSet           func(p *PrintSettings, pageSet T.GtkPageSet)
	printSettingsSetPaperHeight       func(p *PrintSettings, height float64, unit T.GtkUnit)
	printSettingsSetPaperSize         func(p *PrintSettings, paperSize *PaperSize)
	printSettingsSetPaperWidth        func(p *PrintSettings, width float64, unit T.GtkUnit)
	printSettingsSetPrinter           func(p *PrintSettings, printer string)
	printSettingsSetPrinterLpi        func(p *PrintSettings, lpi float64)
	printSettingsSetPrintPages        func(p *PrintSettings, pages T.GtkPrintPages)
	printSettingsSetQuality           func(p *PrintSettings, quality T.GtkPrintQuality)
	printSettingsSetResolution        func(p *PrintSettings, resolution int)
	printSettingsSetResolutionXy      func(p *PrintSettings, resolutionX, resolutionY int)
	printSettingsSetReverse           func(p *PrintSettings, reverse T.Gboolean)
	printSettingsSetScale             func(p *PrintSettings, scale float64)
	printSettingsSetUseColor          func(p *PrintSettings, useColor T.Gboolean)
	printSettingsToFile               func(p *PrintSettings, fileName string, err **T.GError) T.Gboolean
	printSettingsToKeyFile            func(p *PrintSettings, keyFile *T.GKeyFile, groupName string)
	printSettingsUnset                func(p *PrintSettings, key string)
)

func (p *PrintSettings) Copy() *PrintSettings { return printSettingsCopy(p) }
func (p *PrintSettings) Foreach(f PrintSettingsFunc, userData T.Gpointer) {
	printSettingsForeach(p, f, userData)
}
func (p *PrintSettings) Get(key string) string         { return printSettingsGet(p, key) }
func (p *PrintSettings) GetBool(key string) T.Gboolean { return printSettingsGetBool(p, key) }
func (p *PrintSettings) GetCollate() T.Gboolean        { return printSettingsGetCollate(p) }
func (p *PrintSettings) GetDefaultSource() string      { return printSettingsGetDefaultSource(p) }
func (p *PrintSettings) GetDither() string             { return printSettingsGetDither(p) }
func (p *PrintSettings) GetDouble(key string) float64  { return printSettingsGetDouble(p, key) }
func (p *PrintSettings) GetDoubleWithDefault(key string, def float64) float64 {
	return printSettingsGetDoubleWithDefault(p, key, def)
}
func (p *PrintSettings) GetDuplex() T.GtkPrintDuplex { return printSettingsGetDuplex(p) }
func (p *PrintSettings) GetFinishings() string       { return printSettingsGetFinishings(p) }
func (p *PrintSettings) GetInt(key string) int       { return printSettingsGetInt(p, key) }
func (p *PrintSettings) GetIntWithDefault(key string, def int) int {
	return printSettingsGetIntWithDefault(p, key, def)
}
func (p *PrintSettings) GetLength(key string, unit T.GtkUnit) float64 {
	return printSettingsGetLength(p, key, unit)
}
func (p *PrintSettings) GetMediaType() string { return printSettingsGetMediaType(p) }
func (p *PrintSettings) GetNCopies() int      { return printSettingsGetNCopies(p) }
func (p *PrintSettings) GetNumberUp() int     { return printSettingsGetNumberUp(p) }
func (p *PrintSettings) GetNumberUpLayout() T.GtkNumberUpLayout {
	return printSettingsGetNumberUpLayout(p)
}
func (p *PrintSettings) GetOrientation() T.GtkPageOrientation { return printSettingsGetOrientation(p) }
func (p *PrintSettings) GetOutputBin() string                 { return printSettingsGetOutputBin(p) }
func (p *PrintSettings) GetPageRanges(numRanges *int) *T.GtkPageRange {
	return printSettingsGetPageRanges(p, numRanges)
}
func (p *PrintSettings) GetPageSet() T.GtkPageSet { return printSettingsGetPageSet(p) }
func (p *PrintSettings) GetPaperHeight(unit T.GtkUnit) float64 {
	return printSettingsGetPaperHeight(p, unit)
}
func (p *PrintSettings) GetPaperSize() *PaperSize { return printSettingsGetPaperSize(p) }
func (p *PrintSettings) GetPaperWidth(unit T.GtkUnit) float64 {
	return printSettingsGetPaperWidth(p, unit)
}
func (p *PrintSettings) GetPrinter() string             { return printSettingsGetPrinter(p) }
func (p *PrintSettings) GetPrinterLpi() float64         { return printSettingsGetPrinterLpi(p) }
func (p *PrintSettings) GetPrintPages() T.GtkPrintPages { return printSettingsGetPrintPages(p) }
func (p *PrintSettings) GetQuality() T.GtkPrintQuality  { return printSettingsGetQuality(p) }
func (p *PrintSettings) GetResolution() int             { return printSettingsGetResolution(p) }
func (p *PrintSettings) GetResolutionX() int            { return printSettingsGetResolutionX(p) }
func (p *PrintSettings) GetResolutionY() int            { return printSettingsGetResolutionY(p) }
func (p *PrintSettings) GetReverse() T.Gboolean         { return printSettingsGetReverse(p) }
func (p *PrintSettings) GetScale() float64              { return printSettingsGetScale(p) }
func (p *PrintSettings) GetUseColor() T.Gboolean        { return printSettingsGetUseColor(p) }
func (p *PrintSettings) HasKey(key string) T.Gboolean   { return printSettingsHasKey(p, key) }
func (p *PrintSettings) LoadFile(fileName string, err **T.GError) T.Gboolean {
	return printSettingsLoadFile(p, fileName, err)
}
func (p *PrintSettings) LoadKeyFile(keyFile *T.GKeyFile, groupName string, err **T.GError) T.Gboolean {
	return printSettingsLoadKeyFile(p, keyFile, groupName, err)
}
func (p *PrintSettings) Set(key, value string)                { printSettingsSet(p, key, value) }
func (p *PrintSettings) SetBool(key string, value T.Gboolean) { printSettingsSetBool(p, key, value) }
func (p *PrintSettings) SetCollate(collate T.Gboolean)        { printSettingsSetCollate(p, collate) }
func (p *PrintSettings) SetDefaultSource(defaultSource string) {
	printSettingsSetDefaultSource(p, defaultSource)
}
func (p *PrintSettings) SetDither(dither string)             { printSettingsSetDither(p, dither) }
func (p *PrintSettings) SetDouble(key string, value float64) { printSettingsSetDouble(p, key, value) }
func (p *PrintSettings) SetDuplex(duplex T.GtkPrintDuplex)   { printSettingsSetDuplex(p, duplex) }
func (p *PrintSettings) SetFinishings(finishings string)     { printSettingsSetFinishings(p, finishings) }
func (p *PrintSettings) SetInt(key string, value int)        { printSettingsSetInt(p, key, value) }
func (p *PrintSettings) SetLength(key string, value float64, unit T.GtkUnit) {
	printSettingsSetLength(p, key, value, unit)
}
func (p *PrintSettings) SetMediaType(mediaType string) { printSettingsSetMediaType(p, mediaType) }
func (p *PrintSettings) SetNCopies(numCopies int)      { printSettingsSetNCopies(p, numCopies) }
func (p *PrintSettings) SetNumberUp(numberUp int)      { printSettingsSetNumberUp(p, numberUp) }
func (p *PrintSettings) SetNumberUpLayout(numberUpLayout T.GtkNumberUpLayout) {
	printSettingsSetNumberUpLayout(p, numberUpLayout)
}
func (p *PrintSettings) SetOrientation(orientation T.GtkPageOrientation) {
	printSettingsSetOrientation(p, orientation)
}
func (p *PrintSettings) SetOutputBin(outputBin string) { printSettingsSetOutputBin(p, outputBin) }
func (p *PrintSettings) SetPageRanges(pageRanges *T.GtkPageRange, numRanges int) {
	printSettingsSetPageRanges(p, pageRanges, numRanges)
}
func (p *PrintSettings) SetPageSet(pageSet T.GtkPageSet) { printSettingsSetPageSet(p, pageSet) }
func (p *PrintSettings) SetPaperHeight(height float64, unit T.GtkUnit) {
	printSettingsSetPaperHeight(p, height, unit)
}
func (p *PrintSettings) SetPaperSize(paperSize *PaperSize) { printSettingsSetPaperSize(p, paperSize) }
func (p *PrintSettings) SetPaperWidth(width float64, unit T.GtkUnit) {
	printSettingsSetPaperWidth(p, width, unit)
}
func (p *PrintSettings) SetPrinter(printer string)            { printSettingsSetPrinter(p, printer) }
func (p *PrintSettings) SetPrinterLpi(lpi float64)            { printSettingsSetPrinterLpi(p, lpi) }
func (p *PrintSettings) SetPrintPages(pages T.GtkPrintPages)  { printSettingsSetPrintPages(p, pages) }
func (p *PrintSettings) SetQuality(quality T.GtkPrintQuality) { printSettingsSetQuality(p, quality) }
func (p *PrintSettings) SetResolution(resolution int)         { printSettingsSetResolution(p, resolution) }
func (p *PrintSettings) SetResolutionXy(resolutionX, resolutionY int) {
	printSettingsSetResolutionXy(p, resolutionX, resolutionY)
}
func (p *PrintSettings) SetReverse(reverse T.Gboolean)   { printSettingsSetReverse(p, reverse) }
func (p *PrintSettings) SetScale(scale float64)          { printSettingsSetScale(p, scale) }
func (p *PrintSettings) SetUseColor(useColor T.Gboolean) { printSettingsSetUseColor(p, useColor) }
func (p *PrintSettings) ToFile(fileName string, err **T.GError) T.Gboolean {
	return printSettingsToFile(p, fileName, err)
}
func (p *PrintSettings) ToKeyFile(keyFile *T.GKeyFile, groupName string) {
	printSettingsToKeyFile(p, keyFile, groupName)
}
func (p *PrintSettings) Unset(key string) { printSettingsUnset(p, key) }

type Progress struct {
	Widget          Widget
	Adjustment      *Adjustment
	OffscreenPixmap *T.GdkPixmap
	Format          *T.Gchar
	XAlign          float32
	YAlign          float32
	Bits            uint
	// ShowText : 1
	// ActivityMode : 1
	// UseTextFormat : 1
}

var (
	ProgressGetType func() T.GType

	progressSetShowText            func(p *Progress, showText T.Gboolean)
	progressSetTextAlignment       func(p *Progress, xAlign, yAlign float32)
	progressSetFormatString        func(p *Progress, format string)
	progressSetAdjustment          func(p *Progress, adjustment *Adjustment)
	progressConfigure              func(p *Progress, value, min, max float64)
	progressSetPercentage          func(p *Progress, percentage float64)
	progressSetValue               func(p *Progress, value float64)
	progressGetValue               func(p *Progress) float64
	progressSetActivityMode        func(p *Progress, activityMode T.Gboolean)
	progressGetCurrentText         func(p *Progress) string
	progressGetTextFromValue       func(p *Progress, value float64) string
	progressGetCurrentPercentage   func(p *Progress) float64
	progressGetPercentageFromValue func(p *Progress, value float64) float64
)

func (p *Progress) SetShowText(showText T.Gboolean) { progressSetShowText(p, showText) }
func (p *Progress) SetTextAlignment(xAlign, yAlign float32) {
	progressSetTextAlignment(p, xAlign, yAlign)
}
func (p *Progress) SetFormatString(format string)           { progressSetFormatString(p, format) }
func (p *Progress) SetAdjustment(adjustment *Adjustment)    { progressSetAdjustment(p, adjustment) }
func (p *Progress) Configure(value, min, max float64)       { progressConfigure(p, value, min, max) }
func (p *Progress) SetPercentage(percentage float64)        { progressSetPercentage(p, percentage) }
func (p *Progress) SetValue(value float64)                  { progressSetValue(p, value) }
func (p *Progress) GetValue() float64                       { return progressGetValue(p) }
func (p *Progress) SetActivityMode(activityMode T.Gboolean) { progressSetActivityMode(p, activityMode) }
func (p *Progress) GetCurrentText() string                  { return progressGetCurrentText(p) }
func (p *Progress) GetTextFromValue(value float64) string   { return progressGetTextFromValue(p, value) }
func (p *Progress) GetCurrentPercentage() float64           { return progressGetCurrentPercentage(p) }
func (p *Progress) GetPercentageFromValue(value float64) float64 {
	return progressGetPercentageFromValue(p, value)
}

type ProgressBar struct {
	Progress       Progress
	BarStyle       ProgressBarStyle
	Orientation    ProgressBarOrientation
	Blocks         uint
	InBlock        int
	ActivityPos    int
	ActivityStep   uint
	ActivityBlocks uint
	PulseFraction  float64
	Bits           uint
	// ActivityDir : 1
	// Ellipsize : 3
	// Dirty : 1
}
type ProgressBarStyle T.Enum

const (
	PROGRESS_CONTINUOUS ProgressBarStyle = iota
	PROGRESS_DISCRETE
)

type ProgressBarOrientation T.Enum

const (
	PROGRESS_LEFT_TO_RIGHT ProgressBarOrientation = iota
	PROGRESS_RIGHT_TO_LEFT
	PROGRESS_BOTTOM_TO_TOP
	PROGRESS_TOP_TO_BOTTOM
)

var (
	ProgressBarGetType           func() T.GType
	ProgressBarNew               func() *Widget
	ProgressBarNewWithAdjustment func(adjustment *Adjustment) *Widget

	ProgressBarOrientationGetType func() T.GType
	ProgressBarStyleGetType       func() T.GType

	progressBarGetEllipsize      func(p *ProgressBar) T.PangoEllipsizeMode
	progressBarGetFraction       func(p *ProgressBar) float64
	progressBarGetOrientation    func(p *ProgressBar) ProgressBarOrientation
	progressBarGetPulseStep      func(p *ProgressBar) float64
	progressBarGetText           func(p *ProgressBar) string
	progressBarPulse             func(p *ProgressBar)
	progressBarSetActivityBlocks func(p *ProgressBar, blocks uint)
	progressBarSetActivityStep   func(p *ProgressBar, step uint)
	progressBarSetBarStyle       func(p *ProgressBar, style ProgressBarStyle)
	progressBarSetDiscreteBlocks func(p *ProgressBar, blocks uint)
	progressBarSetEllipsize      func(p *ProgressBar, mode T.PangoEllipsizeMode)
	progressBarSetFraction       func(p *ProgressBar, fraction float64)
	progressBarSetOrientation    func(p *ProgressBar, orientation ProgressBarOrientation)
	progressBarSetPulseStep      func(p *ProgressBar, fraction float64)
	progressBarSetText           func(p *ProgressBar, text string)
	progressBarUpdate            func(p *ProgressBar, percentage float64)
)

func (p *ProgressBar) GetEllipsize() T.PangoEllipsizeMode     { return progressBarGetEllipsize(p) }
func (p *ProgressBar) GetFraction() float64                   { return progressBarGetFraction(p) }
func (p *ProgressBar) GetOrientation() ProgressBarOrientation { return progressBarGetOrientation(p) }
func (p *ProgressBar) GetPulseStep() float64                  { return progressBarGetPulseStep(p) }
func (p *ProgressBar) GetText() string                        { return progressBarGetText(p) }
func (p *ProgressBar) Pulse()                                 { progressBarPulse(p) }
func (p *ProgressBar) SetActivityBlocks(blocks uint)          { progressBarSetActivityBlocks(p, blocks) }
func (p *ProgressBar) SetActivityStep(step uint)              { progressBarSetActivityStep(p, step) }
func (p *ProgressBar) SetBarStyle(style ProgressBarStyle)     { progressBarSetBarStyle(p, style) }
func (p *ProgressBar) SetDiscreteBlocks(blocks uint)          { progressBarSetDiscreteBlocks(p, blocks) }
func (p *ProgressBar) SetEllipsize(mode T.PangoEllipsizeMode) { progressBarSetEllipsize(p, mode) }
func (p *ProgressBar) SetFraction(fraction float64)           { progressBarSetFraction(p, fraction) }
func (p *ProgressBar) SetOrientation(orientation ProgressBarOrientation) {
	progressBarSetOrientation(p, orientation)
}
func (p *ProgressBar) SetPulseStep(fraction float64) { progressBarSetPulseStep(p, fraction) }
func (p *ProgressBar) SetText(text string)           { progressBarSetText(p, text) }
func (p *ProgressBar) Update(percentage float64)     { progressBarUpdate(p, percentage) }
