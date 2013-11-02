// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	C "github.com/tHinqa/outside-gtk2/cairo"
	D "github.com/tHinqa/outside-gtk2/gdk"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type PackDirection Enum

const (
	PACK_DIRECTION_LTR PackDirection = iota
	PACK_DIRECTION_RTL
	PACK_DIRECTION_TTB
	PACK_DIRECTION_BTT
)

var PackDirectionGetType func() O.Type

type PackType Enum

const (
	PACK_START PackType = iota
	PACK_END
)

var PackTypeGetType func() O.Type

type PageOrientation Enum

const (
	PAGE_ORIENTATION_PORTRAIT PageOrientation = iota
	PAGE_ORIENTATION_LANDSCAPE
	PAGE_ORIENTATION_REVERSE_PORTRAIT
	PAGE_ORIENTATION_REVERSE_LANDSCAPE
)

var PageOrientationGetType func() O.Type

type PageRange struct {
	Start int
	End   int
}

type PageSet Enum

const (
	PAGE_SET_ALL PageSet = iota
	PAGE_SET_EVEN
	PAGE_SET_ODD
)

var PageSetGetType func() O.Type

type PageSetup struct{}

var (
	PageSetupGetType        func() O.Type
	PageSetupNew            func() *PageSetup
	PageSetupNewFromFile    func(fileName string, error **L.Error) *PageSetup
	PageSetupNewFromKeyFile func(keyFile *L.KeyFile, groupName string, error **L.Error) *PageSetup

	PageSetupCopy                          func(p *PageSetup) *PageSetup
	PageSetupGetBottomMargin               func(p *PageSetup, unit Unit) float64
	PageSetupGetLeftMargin                 func(p *PageSetup, unit Unit) float64
	PageSetupGetOrientation                func(p *PageSetup) Orientation
	PageSetupGetPageHeight                 func(p *PageSetup, unit Unit) float64
	PageSetupGetPageWidth                  func(p *PageSetup, unit Unit) float64
	PageSetupGetPaperHeight                func(p *PageSetup, unit Unit) float64
	PageSetupGetPaperSize                  func(p *PageSetup) *PaperSize
	PageSetupGetPaperWidth                 func(p *PageSetup, unit Unit) float64
	PageSetupGetRightMargin                func(p *PageSetup, unit Unit) float64
	PageSetupGetTopMargin                  func(p *PageSetup, unit Unit) float64
	PageSetupLoadFile                      func(p *PageSetup, fileName string, error **L.Error) bool
	PageSetupLoadKeyFile                   func(p *PageSetup, keyFile *L.KeyFile, groupName string, error **L.Error) bool
	PageSetupSetBottomMargin               func(p *PageSetup, margin float64, unit Unit)
	PageSetupSetLeftMargin                 func(p *PageSetup, margin float64, unit Unit)
	PageSetupSetOrientation                func(p *PageSetup, orientation Orientation)
	PageSetupSetPaperSize                  func(p *PageSetup, size *PaperSize)
	PageSetupSetPaperSizeAndDefaultMargins func(p *PageSetup, size *PaperSize)
	PageSetupSetRightMargin                func(p *PageSetup, margin float64, unit Unit)
	PageSetupSetTopMargin                  func(p *PageSetup, margin float64, unit Unit)
	PageSetupToFile                        func(p *PageSetup, fileName string, error **L.Error) bool
	PageSetupToKeyFile                     func(p *PageSetup, keyFile *L.KeyFile, groupName string)
)

func (p *PageSetup) Copy() *PageSetup                  { return PageSetupCopy(p) }
func (p *PageSetup) GetBottomMargin(unit Unit) float64 { return PageSetupGetBottomMargin(p, unit) }
func (p *PageSetup) GetLeftMargin(unit Unit) float64   { return PageSetupGetLeftMargin(p, unit) }
func (p *PageSetup) GetOrientation() Orientation       { return PageSetupGetOrientation(p) }
func (p *PageSetup) GetPageHeight(unit Unit) float64   { return PageSetupGetPageHeight(p, unit) }
func (p *PageSetup) GetPageWidth(unit Unit) float64    { return PageSetupGetPageWidth(p, unit) }
func (p *PageSetup) GetPaperHeight(unit Unit) float64  { return PageSetupGetPaperHeight(p, unit) }
func (p *PageSetup) GetPaperSize() *PaperSize          { return PageSetupGetPaperSize(p) }
func (p *PageSetup) GetPaperWidth(unit Unit) float64   { return PageSetupGetPaperWidth(p, unit) }
func (p *PageSetup) GetRightMargin(unit Unit) float64  { return PageSetupGetRightMargin(p, unit) }
func (p *PageSetup) GetTopMargin(unit Unit) float64    { return PageSetupGetTopMargin(p, unit) }
func (p *PageSetup) LoadFile(fileName string, err **L.Error) bool {
	return PageSetupLoadFile(p, fileName, err)
}
func (p *PageSetup) LoadKeyFile(keyFile *L.KeyFile, groupName string, err **L.Error) bool {
	return PageSetupLoadKeyFile(p, keyFile, groupName, err)
}
func (p *PageSetup) SetBottomMargin(margin float64, unit Unit) {
	PageSetupSetBottomMargin(p, margin, unit)
}
func (p *PageSetup) SetLeftMargin(margin float64, unit Unit) {
	PageSetupSetLeftMargin(p, margin, unit)
}
func (p *PageSetup) SetOrientation(orientation Orientation) {
	PageSetupSetOrientation(p, orientation)
}
func (p *PageSetup) SetPaperSize(size *PaperSize) { PageSetupSetPaperSize(p, size) }
func (p *PageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	PageSetupSetPaperSizeAndDefaultMargins(p, size)
}
func (p *PageSetup) SetRightMargin(margin float64, unit Unit) {
	PageSetupSetRightMargin(p, margin, unit)
}
func (p *PageSetup) SetTopMargin(margin float64, unit Unit) {
	PageSetupSetTopMargin(p, margin, unit)
}
func (p *PageSetup) ToFile(fileName string, err **L.Error) bool {
	return PageSetupToFile(p, fileName, err)
}
func (p *PageSetup) ToKeyFile(keyFile *L.KeyFile, groupName string) {
	PageSetupToKeyFile(p, keyFile, groupName)
}

type PageSetupDoneFunc func(pageSetup *PageSetup, data T.Gpointer)

type Paned struct {
	Container      Container
	Child1         *Widget
	Child2         *Widget
	Handle         *D.Window
	XorGc          *D.GC
	CursorType     D.CursorType
	HandlePos      D.Rectangle
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
	PanedGetType func() O.Type

	PanedAdd1            func(p *Paned, child *Widget)
	PanedAdd2            func(p *Paned, child *Widget)
	PanedComputePosition func(p *Paned, allocation, child1Req, child2Req int)
	PanedGetChild1       func(p *Paned) *Widget
	PanedGetChild2       func(p *Paned) *Widget
	PanedGetHandleWindow func(p *Paned) *D.Window
	PanedGetPosition     func(p *Paned) int
	PanedPack1           func(p *Paned, child *Widget, resize, shrink bool)
	PanedPack2           func(p *Paned, child *Widget, resize, shrink bool)
	PanedSetPosition     func(p *Paned, position int)
)

func (p *Paned) Add1(child *Widget) { PanedAdd1(p, child) }
func (p *Paned) Add2(child *Widget) { PanedAdd2(p, child) }
func (p *Paned) ComputePosition(allocation, child1Req, child2Req int) {
	PanedComputePosition(p, allocation, child1Req, child2Req)
}
func (p *Paned) GetChild1() *Widget         { return PanedGetChild1(p) }
func (p *Paned) GetChild2() *Widget         { return PanedGetChild2(p) }
func (p *Paned) GetHandleWindow() *D.Window { return PanedGetHandleWindow(p) }
func (p *Paned) GetPosition() int           { return PanedGetPosition(p) }
func (p *Paned) Pack1(child *Widget, resize bool, shrink bool) {
	PanedPack1(p, child, resize, shrink)
}
func (p *Paned) Pack2(child *Widget, resize bool, shrink bool) {
	PanedPack2(p, child, resize, shrink)
}
func (p *Paned) SetPosition(position int) { PanedSetPosition(p, position) }

type PaperSize struct{}

var (
	PaperSizeGetType        func() O.Type
	PaperSizeNew            func(name string) *PaperSize
	PaperSizeNewCustom      func(name, displayName string, width, height float64, unit Unit) *PaperSize
	PaperSizeNewFromKeyFile func(keyFile *L.KeyFile, groupName string, err **L.Error) *PaperSize
	PaperSizeNewFromPpd     func(ppdName, ppdDisplayName string, width, height float64) *PaperSize

	PaperSizeGetDefault    func() string
	PaperSizeGetPaperSizes func(includeCustom bool) *L.List

	PaperSizeCopy                   func(s *PaperSize) *PaperSize
	PaperSizeFree                   func(s *PaperSize)
	PaperSizeGetDefaultBottomMargin func(s *PaperSize, unit Unit) float64
	PaperSizeGetDefaultLeftMargin   func(s *PaperSize, unit Unit) float64
	PaperSizeGetDefaultRightMargin  func(s *PaperSize, unit Unit) float64
	PaperSizeGetDefaultTopMargin    func(s *PaperSize, unit Unit) float64
	PaperSizeGetDisplayName         func(s *PaperSize) string
	PaperSizeGetHeight              func(s *PaperSize, unit Unit) float64
	PaperSizeGetName                func(s *PaperSize) string
	PaperSizeGetPpdName             func(s *PaperSize) string
	PaperSizeGetWidth               func(s *PaperSize, unit Unit) float64
	PaperSizeIsCustom               func(s *PaperSize) bool
	PaperSizeIsEqual                func(s *PaperSize, size2 *PaperSize) bool
	PaperSizeSetSize                func(s *PaperSize, width, height float64, unit Unit)
	PaperSizeToKeyFile              func(s *PaperSize, keyFile *L.KeyFile, groupName string)
)

func (s *PaperSize) Copy() *PaperSize { return PaperSizeCopy(s) }
func (s *PaperSize) Free()            { PaperSizeFree(s) }
func (s *PaperSize) GetDefaultBottomMargin(unit Unit) float64 {
	return PaperSizeGetDefaultBottomMargin(s, unit)
}
func (s *PaperSize) GetDefaultLeftMargin(unit Unit) float64 {
	return PaperSizeGetDefaultLeftMargin(s, unit)
}
func (s *PaperSize) GetDefaultRightMargin(unit Unit) float64 {
	return PaperSizeGetDefaultRightMargin(s, unit)
}
func (s *PaperSize) GetDefaultTopMargin(unit Unit) float64 {
	return PaperSizeGetDefaultTopMargin(s, unit)
}
func (s *PaperSize) GetDisplayName() string        { return PaperSizeGetDisplayName(s) }
func (s *PaperSize) GetHeight(unit Unit) float64   { return PaperSizeGetHeight(s, unit) }
func (s *PaperSize) GetName() string               { return PaperSizeGetName(s) }
func (s *PaperSize) GetPpdName() string            { return PaperSizeGetPpdName(s) }
func (s *PaperSize) GetWidth(unit Unit) float64    { return PaperSizeGetWidth(s, unit) }
func (s *PaperSize) IsCustom() bool                { return PaperSizeIsCustom(s) }
func (s *PaperSize) IsEqual(size2 *PaperSize) bool { return PaperSizeIsEqual(s, size2) }
func (s *PaperSize) SetSize(width, height float64, unit Unit) {
	PaperSizeSetSize(s, width, height, unit)
}
func (s *PaperSize) ToKeyFile(keyFile *L.KeyFile, groupName string) {
	PaperSizeToKeyFile(s, keyFile, groupName)
}

type PathPriorityType Enum

const (
	PATH_PRIO_LOWEST PathPriorityType = iota * 2
	_
	PATH_PRIO_GTK
	PATH_PRIO_APPLICATION
	PATH_PRIO_THEME
	PATH_PRIO_RC
	PATH_PRIO_HIGHEST PathPriorityType = iota*2 - 1
)

var PathPriorityTypeGetType func() O.Type

type PathType Enum

const (
	PATH_WIDGET PathType = iota
	PATH_WIDGET_CLASS
	PATH_CLASS
)

var PathTypeGetType func() O.Type

type Pixmap struct {
	Misc              Misc
	Pixmap            *D.Pixmap
	Mask              *T.GdkBitmap
	PixmapInsensitive *D.Pixmap
	BuildInsensitive  uint // : 1
}

var (
	PixmapGetType func() O.Type
	PixmapNew     func(pixmap *D.Pixmap, mask *T.GdkBitmap) *Widget

	PixmapGet                 func(p *Pixmap, val **D.Pixmap, mask **T.GdkBitmap)
	PixmapSet                 func(p *Pixmap, val *D.Pixmap, mask *T.GdkBitmap)
	PixmapSetBuildInsensitive func(p *Pixmap, build bool)
)

func (p *Pixmap) Get(val **D.Pixmap, mask **T.GdkBitmap) { PixmapGet(p, val, mask) }
func (p *Pixmap) Set(val *D.Pixmap, mask *T.GdkBitmap)   { PixmapSet(p, val, mask) }
func (p *Pixmap) SetBuildInsensitive(build bool)         { PixmapSetBuildInsensitive(p, build) }

type Plug struct {
	Window         Window
	SocketWindow   *D.Window
	ModalityWindow *Widget
	ModalityGroup  *WindowGroup
	GrabbedKeys    *L.HashTable
	SameApp        uint // : 1
}

var (
	PlugGetType       func() O.Type
	PlugNew           func(socketId T.GdkNativeWindow) *Widget
	PlugNewForDisplay func(display *D.Display, socketId T.GdkNativeWindow) *Widget

	PlugConstruct           func(p *Plug, socketId T.GdkNativeWindow)
	PlugConstructForDisplay func(p *Plug, display *D.Display, socketId T.GdkNativeWindow)
	PlugGetEmbedded         func(p *Plug) bool
	PlugGetId               func(p *Plug) T.GdkNativeWindow
	PlugGetSocketWindow     func(p *Plug) *D.Window
)

func (p *Plug) Construct(socketId T.GdkNativeWindow) { PlugConstruct(p, socketId) }
func (p *Plug) ConstructForDisplay(display *D.Display, socketId T.GdkNativeWindow) {
	PlugConstructForDisplay(p, display, socketId)
}
func (p *Plug) GetEmbedded() bool          { return PlugGetEmbedded(p) }
func (p *Plug) GetId() T.GdkNativeWindow   { return PlugGetId(p) }
func (p *Plug) GetSocketWindow() *D.Window { return PlugGetSocketWindow(p) }

type PolicyType Enum

const (
	POLICY_ALWAYS PolicyType = iota
	POLICY_AUTOMATIC
	POLICY_NEVER
)

var PolicyTypeGetType func() O.Type

type PositionType Enum

const (
	POS_LEFT PositionType = iota
	POS_RIGHT
	POS_TOP
	POS_BOTTOM
)

var PositionTypeGetType func() O.Type

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

var (
	PreviewGetType func() O.Type
	PreviewNew     func(t PreviewType) *Widget

	PreviewGetCmap        func() *D.Colormap
	PreviewGetInfo        func() *PreviewInfo
	PreviewGetVisual      func() *D.Visual
	PreviewReset          func()
	PreviewSetColorCube   func(nredShades uint, ngreenShades, nblueShades, ngrayShades uint)
	PreviewSetGamma       func(gamma float64)
	PreviewSetInstallCmap func(installCmap int)
	PreviewSetReserved    func(nreserved int)
	PreviewUninit         func()

	PreviewDrawRow   func(p *Preview, data *T.Guchar, x, y, w int)
	PreviewPut       func(p *Preview, window *D.Window, gc *D.GC, srcx, srcy, destx, desty, width, height int)
	PreviewSetDither func(p *Preview, dither T.GdkRgbDither)
	PreviewSetExpand func(p *Preview, expand bool)
	PreviewSize      func(p *Preview, width, height int)
)

func (p *Preview) DrawRow(data *T.Guchar, x, y, w int) { PreviewDrawRow(p, data, x, y, w) }
func (p *Preview) Put(window *D.Window, gc *D.GC, srcx, srcy, destx, desty, width, height int) {
	PreviewPut(p, window, gc, srcx, srcy, destx, desty, width, height)
}
func (p *Preview) SetDither(dither T.GdkRgbDither) { PreviewSetDither(p, dither) }
func (p *Preview) SetExpand(expand bool)           { PreviewSetExpand(p, expand) }
func (p *Preview) Size(width, height int)          { PreviewSize(p, width, height) }

type PreviewInfo struct {
	Lookup *T.Guchar
	Gamma  float64
}

type PreviewType Enum

const (
	PREVIEW_COLOR PreviewType = iota
	PREVIEW_GRAYSCALE
)

var PreviewTypeGetType func() O.Type

type PrintContext struct{}

var (
	PrintContextGetType func() O.Type

	PrintContextCreatePangoContext func(p *PrintContext) *P.Context
	PrintContextCreatePangoLayout  func(p *PrintContext) *P.Layout
	PrintContextGetCairoContext    func(p *PrintContext) *C.Cairo
	PrintContextGetDpiX            func(p *PrintContext) float64
	PrintContextGetDpiY            func(p *PrintContext) float64
	PrintContextGetHardMargins     func(p *PrintContext, top, bottom, left, right *float64) bool
	PrintContextGetHeight          func(p *PrintContext) float64
	PrintContextGetPageSetup       func(p *PrintContext) *PageSetup
	PrintContextGetPangoFontmap    func(p *PrintContext) *P.FontMap
	PrintContextGetWidth           func(p *PrintContext) float64
	PrintContextSetCairoContext    func(p *PrintContext, cr *C.Cairo, dpiX, dpiY float64)
)

func (p *PrintContext) CreatePangoContext() *P.Context { return PrintContextCreatePangoContext(p) }
func (p *PrintContext) CreatePangoLayout() *P.Layout   { return PrintContextCreatePangoLayout(p) }
func (p *PrintContext) GetCairoContext() *C.Cairo      { return PrintContextGetCairoContext(p) }
func (p *PrintContext) GetDpiX() float64               { return PrintContextGetDpiX(p) }
func (p *PrintContext) GetDpiY() float64               { return PrintContextGetDpiY(p) }
func (p *PrintContext) GetHardMargins(top, bottom, left, right *float64) bool {
	return PrintContextGetHardMargins(p, top, bottom, left, right)
}
func (p *PrintContext) GetHeight() float64          { return PrintContextGetHeight(p) }
func (p *PrintContext) GetPageSetup() *PageSetup    { return PrintContextGetPageSetup(p) }
func (p *PrintContext) GetPangoFontmap() *P.FontMap { return PrintContextGetPangoFontmap(p) }
func (p *PrintContext) GetWidth() float64           { return PrintContextGetWidth(p) }
func (p *PrintContext) SetCairoContext(cr *C.Cairo, dpiX, dpiY float64) {
	PrintContextSetCairoContext(p, cr, dpiX, dpiY)
}

type PrintDuplex Enum

const (
	PRINT_DUPLEX_SIMPLEX PrintDuplex = iota
	PRINT_DUPLEX_HORIZONTAL
	PRINT_DUPLEX_VERTICAL
)

var PrintDuplexGetType func() O.Type

type PrintError Enum

const (
	PRINT_ERROR_GENERAL PrintError = iota
	PRINT_ERROR_INTERNAL_ERROR
	PRINT_ERROR_NOMEM
	PRINT_ERROR_INVALID_FILE
)

var (
	PrintErrorGetType func() O.Type

	PrintErrorQuark func() L.Quark
)

type PrintOperation struct {
	Parent O.Object
	_      *struct{}
}

var (
	PrintOperationGetType       func() O.Type
	PrintOperationNew           func() *PrintOperation
	PrintOperationActionGetType func() O.Type
	PrintOperationResultGetType func() O.Type

	PrintOperationCancel              func(p *PrintOperation)
	PrintOperationDrawPageFinish      func(p *PrintOperation)
	PrintOperationGetDefaultPageSetup func(p *PrintOperation) *PageSetup
	PrintOperationGetEmbedPageSetup   func(p *PrintOperation) bool
	PrintOperationGetError            func(p *PrintOperation, err **L.Error)
	PrintOperationGetHasSelection     func(p *PrintOperation) bool
	PrintOperationGetNPagesToPrint    func(p *PrintOperation) int
	PrintOperationGetPrintSettings    func(p *PrintOperation) *PrintSettings
	PrintOperationGetStatus           func(p *PrintOperation) PrintStatus
	PrintOperationGetStatusString     func(p *PrintOperation) string
	PrintOperationGetSupportSelection func(p *PrintOperation) bool
	PrintOperationIsFinished          func(p *PrintOperation) bool
	PrintOperationRun                 func(p *PrintOperation, action PrintOperationAction, parent *Window, err **L.Error) PrintOperationResult
	PrintOperationSetAllowAsync       func(p *PrintOperation, allowAsync bool)
	PrintOperationSetCurrentPage      func(p *PrintOperation, currentPage int)
	PrintOperationSetCustomTabLabel   func(p *PrintOperation, label string)
	PrintOperationSetDefaultPageSetup func(p *PrintOperation, defaultPageSetup *PageSetup)
	PrintOperationSetDeferDrawing     func(p *PrintOperation)
	PrintOperationSetEmbedPageSetup   func(p *PrintOperation, embed bool)
	PrintOperationSetExportFilename   func(p *PrintOperation, filename string)
	PrintOperationSetHasSelection     func(p *PrintOperation, hasSelection bool)
	PrintOperationSetJobName          func(p *PrintOperation, jobName string)
	PrintOperationSetNPages           func(p *PrintOperation, nPages int)
	PrintOperationSetPrintSettings    func(p *PrintOperation, printSettings *PrintSettings)
	PrintOperationSetShowProgress     func(p *PrintOperation, showProgress bool)
	PrintOperationSetSupportSelection func(p *PrintOperation, supportSelection bool)
	PrintOperationSetTrackPrintStatus func(p *PrintOperation, trackStatus bool)
	PrintOperationSetUnit             func(p *PrintOperation, unit Unit)
	PrintOperationSetUseFullPage      func(p *PrintOperation, fullPage bool)
)

func (p *PrintOperation) Cancel()         { PrintOperationCancel(p) }
func (p *PrintOperation) DrawPageFinish() { PrintOperationDrawPageFinish(p) }
func (p *PrintOperation) GetDefaultPageSetup() *PageSetup {
	return PrintOperationGetDefaultPageSetup(p)
}
func (p *PrintOperation) GetEmbedPageSetup() bool          { return PrintOperationGetEmbedPageSetup(p) }
func (p *PrintOperation) GetError(err **L.Error)          { PrintOperationGetError(p, err) }
func (p *PrintOperation) GetHasSelection() bool            { return PrintOperationGetHasSelection(p) }
func (p *PrintOperation) GetNPagesToPrint() int            { return PrintOperationGetNPagesToPrint(p) }
func (p *PrintOperation) GetPrintSettings() *PrintSettings { return PrintOperationGetPrintSettings(p) }
func (p *PrintOperation) GetStatus() PrintStatus           { return PrintOperationGetStatus(p) }
func (p *PrintOperation) GetStatusString() string          { return PrintOperationGetStatusString(p) }
func (p *PrintOperation) GetSupportSelection() bool        { return PrintOperationGetSupportSelection(p) }
func (p *PrintOperation) IsFinished() bool                 { return PrintOperationIsFinished(p) }
func (p *PrintOperation) Run(action PrintOperationAction, parent *Window, err **L.Error) PrintOperationResult {
	return PrintOperationRun(p, action, parent, err)
}
func (p *PrintOperation) SetAllowAsync(allowAsync bool) {
	PrintOperationSetAllowAsync(p, allowAsync)
}
func (p *PrintOperation) SetCurrentPage(currentPage int) {
	PrintOperationSetCurrentPage(p, currentPage)
}
func (p *PrintOperation) SetCustomTabLabel(label string) { PrintOperationSetCustomTabLabel(p, label) }
func (p *PrintOperation) SetDefaultPageSetup(defaultPageSetup *PageSetup) {
	PrintOperationSetDefaultPageSetup(p, defaultPageSetup)
}
func (p *PrintOperation) SetDeferDrawing() { PrintOperationSetDeferDrawing(p) }
func (p *PrintOperation) SetEmbedPageSetup(embed bool) {
	PrintOperationSetEmbedPageSetup(p, embed)
}
func (p *PrintOperation) SetExportFilename(filename string) {
	PrintOperationSetExportFilename(p, filename)
}
func (p *PrintOperation) SetHasSelection(hasSelection bool) {
	PrintOperationSetHasSelection(p, hasSelection)
}
func (p *PrintOperation) SetJobName(jobName string) { PrintOperationSetJobName(p, jobName) }
func (p *PrintOperation) SetNPages(nPages int)      { PrintOperationSetNPages(p, nPages) }
func (p *PrintOperation) SetPrintSettings(printSettings *PrintSettings) {
	PrintOperationSetPrintSettings(p, printSettings)
}
func (p *PrintOperation) SetShowProgress(showProgress bool) {
	PrintOperationSetShowProgress(p, showProgress)
}
func (p *PrintOperation) SetSupportSelection(supportSelection bool) {
	PrintOperationSetSupportSelection(p, supportSelection)
}
func (p *PrintOperation) SetTrackPrintStatus(trackStatus bool) {
	PrintOperationSetTrackPrintStatus(p, trackStatus)
}
func (p *PrintOperation) SetUnit(unit Unit) { PrintOperationSetUnit(p, unit) }
func (p *PrintOperation) SetUseFullPage(fullPage bool) {
	PrintOperationSetUseFullPage(p, fullPage)
}

type PrintOperationAction Enum

const (
	PRINT_OPERATION_ACTION_PRINT_DIALOG PrintOperationAction = iota
	PRINT_OPERATION_ACTION_PRINT
	PRINT_OPERATION_ACTION_PREVIEW
	PRINT_OPERATION_ACTION_EXPORT
)

type PrintOperationPreview struct{}

var (
	PrintOperationPreviewGetType func() O.Type

	PrintOperationPreviewEndPreview func(p *PrintOperationPreview)
	PrintOperationPreviewIsSelected func(p *PrintOperationPreview, pageNr int) bool
	PrintOperationPreviewRenderPage func(p *PrintOperationPreview, pageNr int)
)

func (p *PrintOperationPreview) EndPreview() { PrintOperationPreviewEndPreview(p) }
func (p *PrintOperationPreview) IsSelected(pageNr int) bool {
	return PrintOperationPreviewIsSelected(p, pageNr)
}
func (p *PrintOperationPreview) RenderPage(pageNr int) { PrintOperationPreviewRenderPage(p, pageNr) }

type PrintOperationResult Enum

const (
	PRINT_OPERATION_RESULT_ERROR PrintOperationResult = iota
	PRINT_OPERATION_RESULT_APPLY
	PRINT_OPERATION_RESULT_CANCEL
	PRINT_OPERATION_RESULT_IN_PROGRESS
)

type PrintPages Enum

const (
	PRINT_PAGES_ALL PrintPages = iota
	PRINT_PAGES_CURRENT
	PRINT_PAGES_RANGES
	PRINT_PAGES_SELECTION
)

var PrintPagesGetType func() O.Type

type (
	PrintSettings struct{}

	PrintSettingsFunc func(key, value string, userData T.Gpointer)
)

type PrintQuality Enum

const (
	PRINT_QUALITY_LOW PrintQuality = iota
	PRINT_QUALITY_NORMAL
	PRINT_QUALITY_HIGH
	PRINT_QUALITY_DRAFT
)

var PrintQualityGetType func() O.Type
var (
	PrintRunPageSetupDialog func(
		parent *Window,
		pageSetup *PageSetup,
		settings *PrintSettings) *PageSetup

	PrintRunPageSetupDialogAsync func(
		parent *Window,
		pageSetup *PageSetup,
		settings *PrintSettings,
		doneCb PageSetupDoneFunc,
		data T.Gpointer)
)
var (
	PrintSettingsGetType        func() O.Type
	PrintSettingsNew            func() *PrintSettings
	PrintSettingsNewFromFile    func(fileName string, err **L.Error) *PrintSettings
	PrintSettingsNewFromKeyFile func(keyFile *L.KeyFile, groupName string, err **L.Error) *PrintSettings

	PrintSettingsCopy                 func(p *PrintSettings) *PrintSettings
	PrintSettingsForeach              func(p *PrintSettings, f PrintSettingsFunc, userData T.Gpointer)
	PrintSettingsGet                  func(p *PrintSettings, key string) string
	PrintSettingsGetBool              func(p *PrintSettings, key string) bool
	PrintSettingsGetCollate           func(p *PrintSettings) bool
	PrintSettingsGetDefaultSource     func(p *PrintSettings) string
	PrintSettingsGetDither            func(p *PrintSettings) string
	PrintSettingsGetDouble            func(p *PrintSettings, key string) float64
	PrintSettingsGetDoubleWithDefault func(p *PrintSettings, key string, def float64) float64
	PrintSettingsGetDuplex            func(p *PrintSettings) PrintDuplex
	PrintSettingsGetFinishings        func(p *PrintSettings) string
	PrintSettingsGetInt               func(p *PrintSettings, key string) int
	PrintSettingsGetIntWithDefault    func(p *PrintSettings, key string, def int) int
	PrintSettingsGetLength            func(p *PrintSettings, key string, unit Unit) float64
	PrintSettingsGetMediaType         func(p *PrintSettings) string
	PrintSettingsGetNCopies           func(p *PrintSettings) int
	PrintSettingsGetNumberUp          func(p *PrintSettings) int
	PrintSettingsGetNumberUpLayout    func(p *PrintSettings) NumberUpLayout
	PrintSettingsGetOrientation       func(p *PrintSettings) Orientation
	PrintSettingsGetOutputBin         func(p *PrintSettings) string
	PrintSettingsGetPageRanges        func(p *PrintSettings, numRanges *int) *PageRange
	PrintSettingsGetPageSet           func(p *PrintSettings) PageSet
	PrintSettingsGetPaperHeight       func(p *PrintSettings, unit Unit) float64
	PrintSettingsGetPaperSize         func(p *PrintSettings) *PaperSize
	PrintSettingsGetPaperWidth        func(p *PrintSettings, unit Unit) float64
	PrintSettingsGetPrinter           func(p *PrintSettings) string
	PrintSettingsGetPrinterLpi        func(p *PrintSettings) float64
	PrintSettingsGetPrintPages        func(p *PrintSettings) PrintPages
	PrintSettingsGetQuality           func(p *PrintSettings) PrintQuality
	PrintSettingsGetResolution        func(p *PrintSettings) int
	PrintSettingsGetResolutionX       func(p *PrintSettings) int
	PrintSettingsGetResolutionY       func(p *PrintSettings) int
	PrintSettingsGetReverse           func(p *PrintSettings) bool
	PrintSettingsGetScale             func(p *PrintSettings) float64
	PrintSettingsGetUseColor          func(p *PrintSettings) bool
	PrintSettingsHasKey               func(p *PrintSettings, key string) bool
	PrintSettingsLoadFile             func(p *PrintSettings, fileName string, err **L.Error) bool
	PrintSettingsLoadKeyFile          func(p *PrintSettings, keyFile *L.KeyFile, groupName string, err **L.Error) bool
	PrintSettingsSet                  func(p *PrintSettings, key, value string)
	PrintSettingsSetBool              func(p *PrintSettings, key string, value bool)
	PrintSettingsSetCollate           func(p *PrintSettings, collate bool)
	PrintSettingsSetDefaultSource     func(p *PrintSettings, defaultSource string)
	PrintSettingsSetDither            func(p *PrintSettings, dither string)
	PrintSettingsSetDouble            func(p *PrintSettings, key string, value float64)
	PrintSettingsSetDuplex            func(p *PrintSettings, duplex PrintDuplex)
	PrintSettingsSetFinishings        func(p *PrintSettings, finishings string)
	PrintSettingsSetInt               func(p *PrintSettings, key string, value int)
	PrintSettingsSetLength            func(p *PrintSettings, key string, value float64, unit Unit)
	PrintSettingsSetMediaType         func(p *PrintSettings, mediaType string)
	PrintSettingsSetNCopies           func(p *PrintSettings, numCopies int)
	PrintSettingsSetNumberUp          func(p *PrintSettings, numberUp int)
	PrintSettingsSetNumberUpLayout    func(p *PrintSettings, numberUpLayout NumberUpLayout)
	PrintSettingsSetOrientation       func(p *PrintSettings, orientation Orientation)
	PrintSettingsSetOutputBin         func(p *PrintSettings, outputBin string)
	PrintSettingsSetPageRanges        func(p *PrintSettings, pageRanges *PageRange, numRanges int)
	PrintSettingsSetPageSet           func(p *PrintSettings, pageSet PageSet)
	PrintSettingsSetPaperHeight       func(p *PrintSettings, height float64, unit Unit)
	PrintSettingsSetPaperSize         func(p *PrintSettings, paperSize *PaperSize)
	PrintSettingsSetPaperWidth        func(p *PrintSettings, width float64, unit Unit)
	PrintSettingsSetPrinter           func(p *PrintSettings, printer string)
	PrintSettingsSetPrinterLpi        func(p *PrintSettings, lpi float64)
	PrintSettingsSetPrintPages        func(p *PrintSettings, pages PrintPages)
	PrintSettingsSetQuality           func(p *PrintSettings, quality PrintQuality)
	PrintSettingsSetResolution        func(p *PrintSettings, resolution int)
	PrintSettingsSetResolutionXy      func(p *PrintSettings, resolutionX, resolutionY int)
	PrintSettingsSetReverse           func(p *PrintSettings, reverse bool)
	PrintSettingsSetScale             func(p *PrintSettings, scale float64)
	PrintSettingsSetUseColor          func(p *PrintSettings, useColor bool)
	PrintSettingsToFile               func(p *PrintSettings, fileName string, err **L.Error) bool
	PrintSettingsToKeyFile            func(p *PrintSettings, keyFile *L.KeyFile, groupName string)
	PrintSettingsUnset                func(p *PrintSettings, key string)
)

func (p *PrintSettings) Copy() *PrintSettings { return PrintSettingsCopy(p) }
func (p *PrintSettings) Foreach(f PrintSettingsFunc, userData T.Gpointer) {
	PrintSettingsForeach(p, f, userData)
}
func (p *PrintSettings) Get(key string) string        { return PrintSettingsGet(p, key) }
func (p *PrintSettings) GetBool(key string) bool      { return PrintSettingsGetBool(p, key) }
func (p *PrintSettings) GetCollate() bool             { return PrintSettingsGetCollate(p) }
func (p *PrintSettings) GetDefaultSource() string     { return PrintSettingsGetDefaultSource(p) }
func (p *PrintSettings) GetDither() string            { return PrintSettingsGetDither(p) }
func (p *PrintSettings) GetDouble(key string) float64 { return PrintSettingsGetDouble(p, key) }
func (p *PrintSettings) GetDoubleWithDefault(key string, def float64) float64 {
	return PrintSettingsGetDoubleWithDefault(p, key, def)
}
func (p *PrintSettings) GetDuplex() PrintDuplex { return PrintSettingsGetDuplex(p) }
func (p *PrintSettings) GetFinishings() string  { return PrintSettingsGetFinishings(p) }
func (p *PrintSettings) GetInt(key string) int  { return PrintSettingsGetInt(p, key) }
func (p *PrintSettings) GetIntWithDefault(key string, def int) int {
	return PrintSettingsGetIntWithDefault(p, key, def)
}
func (p *PrintSettings) GetLength(key string, unit Unit) float64 {
	return PrintSettingsGetLength(p, key, unit)
}
func (p *PrintSettings) GetMediaType() string { return PrintSettingsGetMediaType(p) }
func (p *PrintSettings) GetNCopies() int      { return PrintSettingsGetNCopies(p) }
func (p *PrintSettings) GetNumberUp() int     { return PrintSettingsGetNumberUp(p) }
func (p *PrintSettings) GetNumberUpLayout() NumberUpLayout {
	return PrintSettingsGetNumberUpLayout(p)
}
func (p *PrintSettings) GetOrientation() Orientation { return PrintSettingsGetOrientation(p) }
func (p *PrintSettings) GetOutputBin() string        { return PrintSettingsGetOutputBin(p) }
func (p *PrintSettings) GetPageRanges(numRanges *int) *PageRange {
	return PrintSettingsGetPageRanges(p, numRanges)
}
func (p *PrintSettings) GetPageSet() PageSet { return PrintSettingsGetPageSet(p) }
func (p *PrintSettings) GetPaperHeight(unit Unit) float64 {
	return PrintSettingsGetPaperHeight(p, unit)
}
func (p *PrintSettings) GetPaperSize() *PaperSize { return PrintSettingsGetPaperSize(p) }
func (p *PrintSettings) GetPaperWidth(unit Unit) float64 {
	return PrintSettingsGetPaperWidth(p, unit)
}
func (p *PrintSettings) GetPrinter() string        { return PrintSettingsGetPrinter(p) }
func (p *PrintSettings) GetPrinterLpi() float64    { return PrintSettingsGetPrinterLpi(p) }
func (p *PrintSettings) GetPrintPages() PrintPages { return PrintSettingsGetPrintPages(p) }
func (p *PrintSettings) GetQuality() PrintQuality  { return PrintSettingsGetQuality(p) }
func (p *PrintSettings) GetResolution() int        { return PrintSettingsGetResolution(p) }
func (p *PrintSettings) GetResolutionX() int       { return PrintSettingsGetResolutionX(p) }
func (p *PrintSettings) GetResolutionY() int       { return PrintSettingsGetResolutionY(p) }
func (p *PrintSettings) GetReverse() bool          { return PrintSettingsGetReverse(p) }
func (p *PrintSettings) GetScale() float64         { return PrintSettingsGetScale(p) }
func (p *PrintSettings) GetUseColor() bool         { return PrintSettingsGetUseColor(p) }
func (p *PrintSettings) HasKey(key string) bool    { return PrintSettingsHasKey(p, key) }
func (p *PrintSettings) LoadFile(fileName string, err **L.Error) bool {
	return PrintSettingsLoadFile(p, fileName, err)
}
func (p *PrintSettings) LoadKeyFile(keyFile *L.KeyFile, groupName string, err **L.Error) bool {
	return PrintSettingsLoadKeyFile(p, keyFile, groupName, err)
}
func (p *PrintSettings) Set(key, value string)          { PrintSettingsSet(p, key, value) }
func (p *PrintSettings) SetBool(key string, value bool) { PrintSettingsSetBool(p, key, value) }
func (p *PrintSettings) SetCollate(collate bool)        { PrintSettingsSetCollate(p, collate) }
func (p *PrintSettings) SetDefaultSource(defaultSource string) {
	PrintSettingsSetDefaultSource(p, defaultSource)
}
func (p *PrintSettings) SetDither(dither string)             { PrintSettingsSetDither(p, dither) }
func (p *PrintSettings) SetDouble(key string, value float64) { PrintSettingsSetDouble(p, key, value) }
func (p *PrintSettings) SetDuplex(duplex PrintDuplex)        { PrintSettingsSetDuplex(p, duplex) }
func (p *PrintSettings) SetFinishings(finishings string)     { PrintSettingsSetFinishings(p, finishings) }
func (p *PrintSettings) SetInt(key string, value int)        { PrintSettingsSetInt(p, key, value) }
func (p *PrintSettings) SetLength(key string, value float64, unit Unit) {
	PrintSettingsSetLength(p, key, value, unit)
}
func (p *PrintSettings) SetMediaType(mediaType string) { PrintSettingsSetMediaType(p, mediaType) }
func (p *PrintSettings) SetNCopies(numCopies int)      { PrintSettingsSetNCopies(p, numCopies) }
func (p *PrintSettings) SetNumberUp(numberUp int)      { PrintSettingsSetNumberUp(p, numberUp) }
func (p *PrintSettings) SetNumberUpLayout(numberUpLayout NumberUpLayout) {
	PrintSettingsSetNumberUpLayout(p, numberUpLayout)
}
func (p *PrintSettings) SetOrientation(orientation Orientation) {
	PrintSettingsSetOrientation(p, orientation)
}
func (p *PrintSettings) SetOutputBin(outputBin string) { PrintSettingsSetOutputBin(p, outputBin) }
func (p *PrintSettings) SetPageRanges(pageRanges *PageRange, numRanges int) {
	PrintSettingsSetPageRanges(p, pageRanges, numRanges)
}
func (p *PrintSettings) SetPageSet(pageSet PageSet) { PrintSettingsSetPageSet(p, pageSet) }
func (p *PrintSettings) SetPaperHeight(height float64, unit Unit) {
	PrintSettingsSetPaperHeight(p, height, unit)
}
func (p *PrintSettings) SetPaperSize(paperSize *PaperSize) { PrintSettingsSetPaperSize(p, paperSize) }
func (p *PrintSettings) SetPaperWidth(width float64, unit Unit) {
	PrintSettingsSetPaperWidth(p, width, unit)
}
func (p *PrintSettings) SetPrinter(printer string)       { PrintSettingsSetPrinter(p, printer) }
func (p *PrintSettings) SetPrinterLpi(lpi float64)       { PrintSettingsSetPrinterLpi(p, lpi) }
func (p *PrintSettings) SetPrintPages(pages PrintPages)  { PrintSettingsSetPrintPages(p, pages) }
func (p *PrintSettings) SetQuality(quality PrintQuality) { PrintSettingsSetQuality(p, quality) }
func (p *PrintSettings) SetResolution(resolution int)    { PrintSettingsSetResolution(p, resolution) }
func (p *PrintSettings) SetResolutionXy(resolutionX, resolutionY int) {
	PrintSettingsSetResolutionXy(p, resolutionX, resolutionY)
}
func (p *PrintSettings) SetReverse(reverse bool)   { PrintSettingsSetReverse(p, reverse) }
func (p *PrintSettings) SetScale(scale float64)    { PrintSettingsSetScale(p, scale) }
func (p *PrintSettings) SetUseColor(useColor bool) { PrintSettingsSetUseColor(p, useColor) }
func (p *PrintSettings) ToFile(fileName string, err **L.Error) bool {
	return PrintSettingsToFile(p, fileName, err)
}
func (p *PrintSettings) ToKeyFile(keyFile *L.KeyFile, groupName string) {
	PrintSettingsToKeyFile(p, keyFile, groupName)
}
func (p *PrintSettings) Unset(key string) { PrintSettingsUnset(p, key) }

type PrintStatus Enum

const (
	PRINT_STATUS_INITIAL PrintStatus = iota
	PRINT_STATUS_PREPARING
	PRINT_STATUS_GENERATING_DATA
	PRINT_STATUS_SENDING_DATA
	PRINT_STATUS_PENDING
	PRINT_STATUS_PENDING_ISSUE
	PRINT_STATUS_PRINTING
	PRINT_STATUS_FINISHED
	PRINT_STATUS_FINISHED_ABORTED
)

var PrintStatusGetType func() O.Type

type Progress struct {
	Widget          Widget
	Adjustment      *Adjustment
	OffscreenPixmap *D.Pixmap
	Format          *T.Gchar
	XAlign          float32
	YAlign          float32
	Bits            uint
	// ShowText : 1
	// ActivityMode : 1
	// UseTextFormat : 1
}

var PrivateFlagsGetType func() O.Type //TODO(t):Use?

var (
	ProgressGetType func() O.Type

	ProgressSetShowText            func(p *Progress, showText bool)
	ProgressSetTextAlignment       func(p *Progress, xAlign, yAlign float32)
	ProgressSetFormatString        func(p *Progress, format string)
	ProgressSetAdjustment          func(p *Progress, adjustment *Adjustment)
	ProgressConfigure              func(p *Progress, value, min, max float64)
	ProgressSetPercentage          func(p *Progress, percentage float64)
	ProgressSetValue               func(p *Progress, value float64)
	ProgressGetValue               func(p *Progress) float64
	ProgressSetActivityMode        func(p *Progress, activityMode bool)
	ProgressGetCurrentText         func(p *Progress) string
	ProgressGetTextFromValue       func(p *Progress, value float64) string
	ProgressGetCurrentPercentage   func(p *Progress) float64
	ProgressGetPercentageFromValue func(p *Progress, value float64) float64
)

func (p *Progress) SetShowText(showText bool) { ProgressSetShowText(p, showText) }
func (p *Progress) SetTextAlignment(xAlign, yAlign float32) {
	ProgressSetTextAlignment(p, xAlign, yAlign)
}
func (p *Progress) SetFormatString(format string)         { ProgressSetFormatString(p, format) }
func (p *Progress) SetAdjustment(adjustment *Adjustment)  { ProgressSetAdjustment(p, adjustment) }
func (p *Progress) Configure(value, min, max float64)     { ProgressConfigure(p, value, min, max) }
func (p *Progress) SetPercentage(percentage float64)      { ProgressSetPercentage(p, percentage) }
func (p *Progress) SetValue(value float64)                { ProgressSetValue(p, value) }
func (p *Progress) GetValue() float64                     { return ProgressGetValue(p) }
func (p *Progress) SetActivityMode(activityMode bool)     { ProgressSetActivityMode(p, activityMode) }
func (p *Progress) GetCurrentText() string                { return ProgressGetCurrentText(p) }
func (p *Progress) GetTextFromValue(value float64) string { return ProgressGetTextFromValue(p, value) }
func (p *Progress) GetCurrentPercentage() float64         { return ProgressGetCurrentPercentage(p) }
func (p *Progress) GetPercentageFromValue(value float64) float64 {
	return ProgressGetPercentageFromValue(p, value)
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
type ProgressBarStyle Enum

const (
	PROGRESS_CONTINUOUS ProgressBarStyle = iota
	PROGRESS_DISCRETE
)

type ProgressBarOrientation Enum

const (
	PROGRESS_LEFT_TO_RIGHT ProgressBarOrientation = iota
	PROGRESS_RIGHT_TO_LEFT
	PROGRESS_BOTTOM_TO_TOP
	PROGRESS_TOP_TO_BOTTOM
)

var (
	ProgressBarGetType           func() O.Type
	ProgressBarNew               func() *Widget
	ProgressBarNewWithAdjustment func(adjustment *Adjustment) *Widget

	ProgressBarOrientationGetType func() O.Type
	ProgressBarStyleGetType       func() O.Type

	ProgressBarGetEllipsize      func(p *ProgressBar) P.EllipsizeMode
	ProgressBarGetFraction       func(p *ProgressBar) float64
	ProgressBarGetOrientation    func(p *ProgressBar) ProgressBarOrientation
	ProgressBarGetPulseStep      func(p *ProgressBar) float64
	ProgressBarGetText           func(p *ProgressBar) string
	ProgressBarPulse             func(p *ProgressBar)
	ProgressBarSetActivityBlocks func(p *ProgressBar, blocks uint)
	ProgressBarSetActivityStep   func(p *ProgressBar, step uint)
	ProgressBarSetBarStyle       func(p *ProgressBar, style ProgressBarStyle)
	ProgressBarSetDiscreteBlocks func(p *ProgressBar, blocks uint)
	ProgressBarSetEllipsize      func(p *ProgressBar, mode P.EllipsizeMode)
	ProgressBarSetFraction       func(p *ProgressBar, fraction float64)
	ProgressBarSetOrientation    func(p *ProgressBar, orientation ProgressBarOrientation)
	ProgressBarSetPulseStep      func(p *ProgressBar, fraction float64)
	ProgressBarSetText           func(p *ProgressBar, text string)
	ProgressBarUpdate            func(p *ProgressBar, percentage float64)
)

func (p *ProgressBar) GetEllipsize() P.EllipsizeMode          { return ProgressBarGetEllipsize(p) }
func (p *ProgressBar) GetFraction() float64                   { return ProgressBarGetFraction(p) }
func (p *ProgressBar) GetOrientation() ProgressBarOrientation { return ProgressBarGetOrientation(p) }
func (p *ProgressBar) GetPulseStep() float64                  { return ProgressBarGetPulseStep(p) }
func (p *ProgressBar) GetText() string                        { return ProgressBarGetText(p) }
func (p *ProgressBar) Pulse()                                 { ProgressBarPulse(p) }
func (p *ProgressBar) SetActivityBlocks(blocks uint)          { ProgressBarSetActivityBlocks(p, blocks) }
func (p *ProgressBar) SetActivityStep(step uint)              { ProgressBarSetActivityStep(p, step) }
func (p *ProgressBar) SetBarStyle(style ProgressBarStyle)     { ProgressBarSetBarStyle(p, style) }
func (p *ProgressBar) SetDiscreteBlocks(blocks uint)          { ProgressBarSetDiscreteBlocks(p, blocks) }
func (p *ProgressBar) SetEllipsize(mode P.EllipsizeMode)      { ProgressBarSetEllipsize(p, mode) }
func (p *ProgressBar) SetFraction(fraction float64)           { ProgressBarSetFraction(p, fraction) }
func (p *ProgressBar) SetOrientation(orientation ProgressBarOrientation) {
	ProgressBarSetOrientation(p, orientation)
}
func (p *ProgressBar) SetPulseStep(fraction float64) { ProgressBarSetPulseStep(p, fraction) }
func (p *ProgressBar) SetText(text string)           { ProgressBarSetText(p, text) }
func (p *ProgressBar) Update(percentage float64)     { ProgressBarUpdate(p, percentage) }
