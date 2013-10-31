// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gtksourceview provides API definitions for accessing
//libgtksourceview-2.0-0.dll.
package gtksourceview

import (
	"github.com/tHinqa/outside"
	D "github.com/tHinqa/outside-gtk2/gdk"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	G "github.com/tHinqa/outside-gtk2/gtk"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type Enum int

type Buffer struct {
	Parent G.TextBuffer
	_      *struct{}
}

var (
	BufferGetType         func() O.Type
	BufferNew             func(table *G.TextTagTable) *Buffer
	BufferNewWithLanguage func(language *Language) *Buffer

	BufferBackwardIterToSourceMark         func(b *Buffer, iter *G.TextIter, category string) bool
	BufferBeginNotUndoableAction           func(b *Buffer)
	BufferCanRedo                          func(b *Buffer) bool
	BufferCanUndo                          func(b *Buffer) bool
	BufferCreateSourceMark                 func(b *Buffer, name, category string, where *G.TextIter) *Mark
	BufferEndNotUndoableAction             func(b *Buffer)
	BufferEnsureHighlight                  func(b *Buffer, start, end *G.TextIter)
	BufferForwardIterToSourceMark          func(b *Buffer, iter *G.TextIter, category string) bool
	BufferGetContextClassesAtIter          func(b *Buffer, iter *G.TextIter) []string
	BufferGetHighlightMatchingBrackets     func(b *Buffer) bool
	BufferGetHighlightSyntax               func(b *Buffer) bool
	BufferGetLanguage                      func(b *Buffer) *Language
	BufferGetMaxUndoLevels                 func(b *Buffer) int
	BufferGetSourceMarksAtIter             func(b *Buffer, iter *G.TextIter, category string) *L.SList
	BufferGetSourceMarksAtLine             func(b *Buffer, line int, category string) *L.SList
	BufferGetStyleScheme                   func(b *Buffer) *StyleScheme
	BufferGetUndoManager                   func(b *Buffer) *UndoManager
	BufferIterBackwardToContextClassToggle func(b *Buffer, iter *G.TextIter, contextClass string) bool
	BufferIterForwardToContextClassToggle  func(b *Buffer, iter *G.TextIter, contextClass string) bool
	BufferIterHasContextClass              func(b *Buffer, iter *G.TextIter, contextClass string) bool
	BufferRedo                             func(b *Buffer)
	BufferRemoveSourceMarks                func(b *Buffer, start *G.TextIter, end *G.TextIter, category string)
	BufferSetHighlightMatchingBrackets     func(b *Buffer, highlight bool)
	BufferSetHighlightSyntax               func(b *Buffer, highlight bool)
	BufferSetLanguage                      func(b *Buffer, language *Language)
	BufferSetMaxUndoLevels                 func(b *Buffer, maxUndoLevels int)
	BufferSetStyleScheme                   func(b *Buffer, scheme *StyleScheme)
	BufferSetUndoManager                   func(b *Buffer, manager *UndoManager)
	BufferUndo                             func(b *Buffer)
)

func (b *Buffer) BackwardIterToSourceMark(iter *G.TextIter, category string) bool {
	return BufferBackwardIterToSourceMark(b, iter, category)
}
func (b *Buffer) BeginNotUndoableAction() { BufferBeginNotUndoableAction(b) }
func (b *Buffer) CanRedo() bool           { return BufferCanRedo(b) }
func (b *Buffer) CanUndo() bool           { return BufferCanUndo(b) }
func (b *Buffer) ContextClassesAtIter(iter *G.TextIter) []string {
	return BufferGetContextClassesAtIter(b, iter)
}
func (b *Buffer) CreateSourceMark(name, category string, where *G.TextIter) *Mark {
	return BufferCreateSourceMark(b, name, category, where)
}
func (b *Buffer) EndNotUndoableAction()                  { BufferEndNotUndoableAction(b) }
func (b *Buffer) EnsureHighlight(start, end *G.TextIter) { BufferEnsureHighlight(b, start, end) }
func (b *Buffer) ForwardIterToSourceMark(iter *G.TextIter, category string) bool {
	return BufferForwardIterToSourceMark(b, iter, category)
}
func (b *Buffer) HighlightMatchingBrackets() bool {
	return BufferGetHighlightMatchingBrackets(b)
}
func (b *Buffer) HighlightSyntax() bool { return BufferGetHighlightSyntax(b) }
func (b *Buffer) IterBackwardToContextClassToggle(iter *G.TextIter, contextClass string) bool {
	return BufferIterBackwardToContextClassToggle(b, iter, contextClass)
}
func (b *Buffer) IterForwardToContextClassToggle(iter *G.TextIter, contextClass string) bool {
	return BufferIterForwardToContextClassToggle(b, iter, contextClass)
}
func (b *Buffer) IterHasContextClass(iter *G.TextIter, contextClass string) bool {
	return BufferIterHasContextClass(b, iter, contextClass)
}
func (b *Buffer) Language() *Language { return BufferGetLanguage(b) }
func (b *Buffer) MaxUndoLevels() int  { return BufferGetMaxUndoLevels(b) }
func (b *Buffer) Redo()               { BufferRedo(b) }
func (b *Buffer) RemoveSourceMarks(start *G.TextIter, end *G.TextIter, category string) {
	BufferRemoveSourceMarks(b, start, end, category)
}
func (b *Buffer) SetHighlightMatchingBrackets(highlight bool) {
	BufferSetHighlightMatchingBrackets(b, highlight)
}
func (b *Buffer) SetHighlightSyntax(highlight bool)   { BufferSetHighlightSyntax(b, highlight) }
func (b *Buffer) SetLanguage(language *Language)      { BufferSetLanguage(b, language) }
func (b *Buffer) SetMaxUndoLevels(maxUndoLevels int)  { BufferSetMaxUndoLevels(b, maxUndoLevels) }
func (b *Buffer) SetStyleScheme(scheme *StyleScheme)  { BufferSetStyleScheme(b, scheme) }
func (b *Buffer) SetUndoManager(manager *UndoManager) { BufferSetUndoManager(b, manager) }
func (b *Buffer) SourceMarksAtIter(iter *G.TextIter, category string) *L.SList {
	return BufferGetSourceMarksAtIter(b, iter, category)
}
func (b *Buffer) SourceMarksAtLine(line int, category string) *L.SList {
	return BufferGetSourceMarksAtLine(b, line, category)
}
func (b *Buffer) StyleScheme() *StyleScheme { return BufferGetStyleScheme(b) }
func (b *Buffer) Undo()                     { BufferUndo(b) }
func (b *Buffer) UndoManager() *UndoManager { return BufferGetUndoManager(b) }

type Completion struct {
	Parent G.Object
	_      *struct{}
}

var (
	CompletionGetType    func() O.Type
	CompletionErrorQuark func() T.GQuark

	CompletionAddProvider        func(c *Completion, provider *CompletionProvider, err **T.GError) bool
	CompletionBlockInteractive   func(c *Completion)
	CompletionCreateContext      func(c *Completion, position *G.TextIter) *CompletionContext
	CompletionGetInfoWindow      func(c *Completion) *CompletionInfo
	CompletionGetProviders       func(c *Completion) *T.GList
	CompletionGetView            func(c *Completion) *View
	CompletionHide               func(c *Completion)
	CompletionMoveWindow         func(c *Completion, iter *G.TextIter)
	CompletionRemoveProvider     func(c *Completion, provider *CompletionProvider, err **T.GError) bool
	CompletionShow               func(c *Completion, providers *T.GList, context *CompletionContext) bool
	CompletionUnblockInteractive func(c *Completion)
)

func (c *Completion) AddProvider(provider *CompletionProvider, err **T.GError) bool {
	return CompletionAddProvider(c, provider, err)
}
func (c *Completion) BlockInteractive() { CompletionBlockInteractive(c) }
func (c *Completion) CreateContext(position *G.TextIter) *CompletionContext {
	return CompletionCreateContext(c, position)
}
func (c *Completion) Hide()                       { CompletionHide(c) }
func (c *Completion) InfoWindow() *CompletionInfo { return CompletionGetInfoWindow(c) }
func (c *Completion) MoveWindow(iter *G.TextIter) { CompletionMoveWindow(c, iter) }
func (c *Completion) Providers() *T.GList         { return CompletionGetProviders(c) }
func (c *Completion) RemoveProvider(provider *CompletionProvider, err **T.GError) bool {
	return CompletionRemoveProvider(c, provider, err)
}
func (c *Completion) Show(providers *T.GList, context *CompletionContext) bool {
	return CompletionShow(c, providers, context)
}
func (c *Completion) UnblockInteractive() { CompletionUnblockInteractive(c) }
func (c *Completion) View() *View         { return CompletionGetView(c) }

type CompletionActivation Enum

const (
	COMPLETION_ACTIVATION_INTERACTIVE CompletionActivation = 1 << iota
	COMPLETION_ACTIVATION_USER_REQUESTED
	COMPLETION_ACTIVATION_NONE CompletionActivation = 0
)

var CompletionActivationGetType func() O.Type

type CompletionContext struct {
	Parent T.GInitiallyUnowned
	_      *struct{}
}

var (
	CompletionContextGetType func() O.Type

	CompletionContextAddProposals  func(c *CompletionContext, provider *CompletionProvider, proposals *T.GList, finished bool)
	CompletionContextGetActivation func(c *CompletionContext) CompletionActivation
	CompletionContextGetIter       func(c *CompletionContext, iter *G.TextIter)
)

func (c *CompletionContext) Activation() CompletionActivation {
	return CompletionContextGetActivation(c)
}
func (c *CompletionContext) AddProposals(provider *CompletionProvider, proposals *T.GList, finished bool) {
	CompletionContextAddProposals(c, provider, proposals, finished)
}
func (c *CompletionContext) Iter(iter *G.TextIter) { CompletionContextGetIter(c, iter) }

var CompletionErrorGetType func() O.Type

type CompletionInfo struct {
	Parent G.Window
	_      *struct{}
}

var (
	CompletionInfoGetType func() O.Type
	CompletionInfoNew     func() *CompletionInfo

	CompletionInfoGetWidget     func(c *CompletionInfo) *G.Widget
	CompletionInfoMoveToIter    func(c *CompletionInfo, view *G.TextView, iter *G.TextIter)
	CompletionInfoProcessResize func(c *CompletionInfo)
	CompletionInfoSetSizing     func(c *CompletionInfo, width, height int, shrinkWidth bool, shrinkHeight bool)
	CompletionInfoSetWidget     func(c *CompletionInfo, widget *G.Widget)
)

func (c *CompletionInfo) MoveToIter(view *G.TextView, iter *G.TextIter) {
	CompletionInfoMoveToIter(c, view, iter)
}
func (c *CompletionInfo) ProcessResize() { CompletionInfoProcessResize(c) }
func (c *CompletionInfo) SetSizing(width, height int, shrinkWidth bool, shrinkHeight bool) {
	CompletionInfoSetSizing(c, width, height, shrinkWidth, shrinkHeight)
}
func (c *CompletionInfo) SetWidget(widget *G.Widget) { CompletionInfoSetWidget(c, widget) }
func (c *CompletionInfo) Widget() *G.Widget          { return CompletionInfoGetWidget(c) }

type CompletionItem struct {
	Parent O.Object
	_      *struct{}
}

var (
	CompletionItemGetType       func() O.Type
	CompletionItemNew           func(label, text string, icon *D.Pixbuf, info string) *CompletionItem
	CompletionItemNewWithMarkup func(markup, text string, icon *D.Pixbuf, info string) *CompletionItem
	CompletionItemNewFromStock  func(label, text, stock, info string) *CompletionItem
)

type CompletionProposal struct{}

var (
	CompletionProposalGetType func() O.Type

	CompletionProposalChanged   func(c *CompletionProposal)
	CompletionProposalEqual     func(c *CompletionProposal, other *CompletionProposal) bool
	CompletionProposalGetIcon   func(c *CompletionProposal) *D.Pixbuf
	CompletionProposalGetInfo   func(c *CompletionProposal) string
	CompletionProposalGetLabel  func(c *CompletionProposal) string
	CompletionProposalGetMarkup func(c *CompletionProposal) string
	CompletionProposalGetText   func(c *CompletionProposal) string
	CompletionProposalHash      func(c *CompletionProposal) uint
)

func (c *CompletionProposal) Changed() { CompletionProposalChanged(c) }
func (c *CompletionProposal) Equal(other *CompletionProposal) bool {
	return CompletionProposalEqual(c, other)
}
func (c *CompletionProposal) Hash() uint      { return CompletionProposalHash(c) }
func (c *CompletionProposal) Icon() *D.Pixbuf { return CompletionProposalGetIcon(c) }
func (c *CompletionProposal) Info() string    { return CompletionProposalGetInfo(c) }
func (c *CompletionProposal) Label() string   { return CompletionProposalGetLabel(c) }
func (c *CompletionProposal) Markup() string  { return CompletionProposalGetMarkup(c) }
func (c *CompletionProposal) Text() string    { return CompletionProposalGetText(c) }

type CompletionProvider struct{}

var (
	CompletionProviderGetType func() O.Type

	CompletionProviderActivateProposal    func(c *CompletionProvider, proposal *CompletionProposal, iter *G.TextIter) bool
	CompletionProviderGetActivation       func(c *CompletionProvider) CompletionActivation
	CompletionProviderGetIcon             func(c *CompletionProvider) *D.Pixbuf
	CompletionProviderGetInfoWidget       func(c *CompletionProvider, proposal *CompletionProposal) *G.Widget
	CompletionProviderGetInteractiveDelay func(c *CompletionProvider) int
	CompletionProviderGetName             func(c *CompletionProvider) string
	CompletionProviderGetPriority         func(c *CompletionProvider) int
	CompletionProviderGetStartIter        func(c *CompletionProvider, context *CompletionContext, proposal *CompletionProposal, iter *G.TextIter) bool
	CompletionProviderMatch               func(c *CompletionProvider, context *CompletionContext) bool
	CompletionProviderPopulate            func(c *CompletionProvider, context *CompletionContext)
	CompletionProviderUpdateInfo          func(c *CompletionProvider, proposal *CompletionProposal, info *CompletionInfo)
)

func (c *CompletionProvider) Activation() CompletionActivation {
	return CompletionProviderGetActivation(c)
}
func (c *CompletionProvider) ActivateProposal(proposal *CompletionProposal, iter *G.TextIter) bool {
	return CompletionProviderActivateProposal(c, proposal, iter)
}
func (c *CompletionProvider) Icon() *D.Pixbuf { return CompletionProviderGetIcon(c) }
func (c *CompletionProvider) InfoWidget(proposal *CompletionProposal) *G.Widget {
	return CompletionProviderGetInfoWidget(c, proposal)
}
func (c *CompletionProvider) InteractiveDelay() int {
	return CompletionProviderGetInteractiveDelay(c)
}
func (c *CompletionProvider) Match(context *CompletionContext) bool {
	return CompletionProviderMatch(c, context)
}
func (c *CompletionProvider) Name() string { return CompletionProviderGetName(c) }
func (c *CompletionProvider) Populate(context *CompletionContext) {
	CompletionProviderPopulate(c, context)
}
func (c *CompletionProvider) Priority() int { return CompletionProviderGetPriority(c) }
func (c *CompletionProvider) StartIter(context *CompletionContext, proposal *CompletionProposal, iter *G.TextIter) bool {
	return CompletionProviderGetStartIter(c, context, proposal, iter)
}
func (c *CompletionProvider) UpdateInfo(proposal *CompletionProposal, info *CompletionInfo) {
	CompletionProviderUpdateInfo(c, proposal, info)
}

type CompletionWords struct {
	Parent O.Object
	_      *struct{}
}

var (
	CompletionWordsGetType func() O.Type
	CompletionWordsNew     func(name string, icon *D.Pixbuf) *CompletionWords

	CompletionWordsRegister   func(words *CompletionWords, buffer *G.TextBuffer)
	CompletionWordsUnregister func(words *CompletionWords, buffer *G.TextBuffer)
)

func (c *CompletionWords) Register(buffer *G.TextBuffer)   { CompletionWordsRegister(c, buffer) }
func (c *CompletionWords) Unregister(buffer *G.TextBuffer) { CompletionWordsUnregister(c, buffer) }

type DrawSpacesFlags Enum

const (
	DRAW_SPACES_SPACE DrawSpacesFlags = 1 << iota
	DRAW_SPACES_TAB
	DRAW_SPACES_NEWLINE
	DRAW_SPACES_NBSP
	DRAW_SPACES_LEADING
	DRAW_SPACES_TEXT
	DRAW_SPACES_TRAILING
	DRAW_SPACES_ALL = DRAW_SPACES_SPACE |
		DRAW_SPACES_TAB |
		DRAW_SPACES_NEWLINE |
		DRAW_SPACES_NBSP |
		DRAW_SPACES_LEADING |
		DRAW_SPACES_TEXT |
		DRAW_SPACES_TRAILING
)

var DrawSpacesFlagsGetType func() O.Type

type Gutter struct {
	Parent O.Object
	_      *struct{}
}

var (
	GutterGetType func() O.Type

	GutterGetWindow       func(g *Gutter) *D.Window
	GutterInsert          func(g *Gutter, renderer *G.CellRenderer, position int)
	GutterQueueDraw       func(g *Gutter)
	GutterRemove          func(g *Gutter, renderer *G.CellRenderer)
	GutterReorder         func(g *Gutter, renderer *G.CellRenderer, position int)
	GutterSetCellDataFunc func(g *Gutter, renderer *G.CellRenderer, f GutterDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
	GutterSetCellSizeFunc func(g *Gutter, renderer *G.CellRenderer, f GutterSizeFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
)

func (g *Gutter) Insert(renderer *G.CellRenderer, position int) { GutterInsert(g, renderer, position) }
func (g *Gutter) QueueDraw()                                    { GutterQueueDraw(g) }
func (g *Gutter) Remove(renderer *G.CellRenderer)               { GutterRemove(g, renderer) }
func (g *Gutter) Reorder(renderer *G.CellRenderer, position int) {
	GutterReorder(g, renderer, position)
}
func (g *Gutter) SetCellDataFunc(renderer *G.CellRenderer, f GutterDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	GutterSetCellDataFunc(g, renderer, f, funcData, destroy)
}
func (g *Gutter) SetCellSizeFunc(renderer *G.CellRenderer, f GutterSizeFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	GutterSetCellSizeFunc(g, renderer, f, funcData, destroy)
}
func (g *Gutter) Window() *D.Window { return GutterGetWindow(g) }

type GutterDataFunc func(gutter *Gutter, cell *G.CellRenderer,
	LineNumber int, currentLine bool, data T.Gpointer)

type GutterSizeFunc func(
	Gutter *Gutter, cell *G.CellRenderer, data T.Gpointer)

var (
	IterForwardSearch  func(iter *G.TextIter, str string, flags SearchFlags, matchStart *G.TextIter, matchEnd *G.TextIter, limit *G.TextIter) bool
	IterBackwardSearch func(iter *G.TextIter, str string, flags SearchFlags, matchStart *G.TextIter, matchEnd *G.TextIter, limit *G.TextIter) bool
)

type Language struct {
	Parent O.Object
	_      *struct{}
}

var (
	LanguageGetType func() O.Type

	LanguageGetGlobs     func(language *Language) []string
	LanguageGetHidden    func(language *Language) bool
	LanguageGetId        func(language *Language) string
	LanguageGetMetadata  func(language *Language, name string) string
	LanguageGetMimeTypes func(language *Language) []string
	LanguageGetName      func(language *Language) string
	LanguageGetSection   func(language *Language) string
	LanguageGetStyleIds  func(language *Language) []string
	LanguageGetStyleName func(language *Language, styleId string) string
)

func (l *Language) Globs() []string                 { return LanguageGetGlobs(l) }
func (l *Language) Hidden() bool                    { return LanguageGetHidden(l) }
func (l *Language) Id() string                      { return LanguageGetId(l) }
func (l *Language) Metadata(name string) string     { return LanguageGetMetadata(l, name) }
func (l *Language) MimeTypes() []string             { return LanguageGetMimeTypes(l) }
func (l *Language) Name() string                    { return LanguageGetName(l) }
func (l *Language) Section() string                 { return LanguageGetSection(l) }
func (l *Language) StyleIds() []string              { return LanguageGetStyleIds(l) }
func (l *Language) StyleName(styleId string) string { return LanguageGetStyleName(l, styleId) }

type LanguageManager struct {
	Parent O.Object
	_      *struct{}
}

var (
	LanguageManagerGetType func() O.Type
	LanguageManagerNew     func() *LanguageManager

	LanguageManagerGetDefault func() *LanguageManager

	LanguageManagerGetLanguage    func(l *LanguageManager, id string) *Language
	LanguageManagerGetLanguageIds func(l *LanguageManager) []string
	LanguageManagerGetSearchPath  func(l *LanguageManager) []string
	LanguageManagerGuessLanguage  func(l *LanguageManager, filename string, contentType string) *Language
	LanguageManagerSetSearchPath  func(l *LanguageManager, dirs []string)
)

func (l *LanguageManager) GuessLanguage(filename string, contentType string) *Language {
	return LanguageManagerGuessLanguage(l, filename, contentType)
}
func (l *LanguageManager) Language(id string) *Language { return LanguageManagerGetLanguage(l, id) }
func (l *LanguageManager) LanguageIds() []string        { return LanguageManagerGetLanguageIds(l) }
func (l *LanguageManager) SearchPath() []string         { return LanguageManagerGetSearchPath(l) }
func (l *LanguageManager) SetSearchPath(dirs []string)  { LanguageManagerSetSearchPath(l, dirs) }

type Mark struct {
	Parent G.TextMark
	_      *struct{}
}

var (
	MarkGetType func() O.Type
	MarkNew     func(name, category string) *Mark

	MarkGetCategory func(m *Mark) string
	MarkNext        func(m *Mark, category string) *Mark
	MarkPrev        func(m *Mark, category string) *Mark
)

func (m *Mark) Category() string           { return MarkGetCategory(m) }
func (m *Mark) Next(category string) *Mark { return MarkNext(m, category) }
func (m *Mark) Prev(category string) *Mark { return MarkPrev(m, category) }

type PrintCompositor struct {
	Parent O.Object
	_      *struct{}
}

var (
	PrintCompositorGetType     func() O.Type
	PrintCompositorNew         func(buffer *Buffer) *PrintCompositor
	PrintCompositorNewFromView func(view *View) *PrintCompositor

	PrintCompositorDrawPage               func(p *PrintCompositor, context *G.PrintContext, pageNr int)
	PrintCompositorGetBodyFontName        func(p *PrintCompositor) string
	PrintCompositorGetBottomMargin        func(p *PrintCompositor, unit G.Unit) float64
	PrintCompositorGetBuffer              func(p *PrintCompositor) *Buffer
	PrintCompositorGetFooterFontName      func(p *PrintCompositor) string
	PrintCompositorGetHeaderFontName      func(p *PrintCompositor) string
	PrintCompositorGetHighlightSyntax     func(p *PrintCompositor) bool
	PrintCompositorGetLeftMargin          func(p *PrintCompositor, unit G.Unit) float64
	PrintCompositorGetLineNumbersFontName func(p *PrintCompositor) string
	PrintCompositorGetNPages              func(p *PrintCompositor) int
	PrintCompositorGetPaginationProgress  func(p *PrintCompositor) float64
	PrintCompositorGetPrintFooter         func(p *PrintCompositor) bool
	PrintCompositorGetPrintHeader         func(p *PrintCompositor) bool
	PrintCompositorGetPrintLineNumbers    func(p *PrintCompositor) uint
	PrintCompositorGetRightMargin         func(p *PrintCompositor, unit G.Unit) float64
	PrintCompositorGetTabWidth            func(p *PrintCompositor) uint
	PrintCompositorGetTopMargin           func(p *PrintCompositor, unit G.Unit) float64
	PrintCompositorGetWrapMode            func(p *PrintCompositor) G.WrapMode
	PrintCompositorPaginate               func(p *PrintCompositor, context *G.PrintContext) bool
	PrintCompositorSetBodyFontName        func(p *PrintCompositor, fontName string)
	PrintCompositorSetBottomMargin        func(p *PrintCompositor, margin float64, unit G.Unit)
	PrintCompositorSetFooterFontName      func(p *PrintCompositor, fontName string)
	PrintCompositorSetFooterFormat        func(p *PrintCompositor, separator bool, left, center, right string)
	PrintCompositorSetHeaderFontName      func(p *PrintCompositor, fontName string)
	PrintCompositorSetHeaderFormat        func(p *PrintCompositor, separator bool, left, center, right string)
	PrintCompositorSetHighlightSyntax     func(p *PrintCompositor, highlight bool)
	PrintCompositorSetLeftMargin          func(p *PrintCompositor, margin float64, unit G.Unit)
	PrintCompositorSetLineNumbersFontName func(p *PrintCompositor, fontName string)
	PrintCompositorSetPrintFooter         func(p *PrintCompositor, prt bool)
	PrintCompositorSetPrintHeader         func(p *PrintCompositor, prt bool)
	PrintCompositorSetPrintLineNumbers    func(p *PrintCompositor, interval uint)
	PrintCompositorSetRightMargin         func(p *PrintCompositor, margin float64, unit G.Unit)
	PrintCompositorSetTabWidth            func(p *PrintCompositor, width uint)
	PrintCompositorSetTopMargin           func(p *PrintCompositor, margin float64, unit G.Unit)
	PrintCompositorSetWrapMode            func(p *PrintCompositor, wrapMode G.WrapMode)
)

func (p *PrintCompositor) BodyFontName() string { return PrintCompositorGetBodyFontName(p) }
func (p *PrintCompositor) BottomMargin(unit G.Unit) float64 {
	return PrintCompositorGetBottomMargin(p, unit)
}
func (p *PrintCompositor) Buffer() *Buffer { return PrintCompositorGetBuffer(p) }
func (p *PrintCompositor) DrawPage(context *G.PrintContext, pageNr int) {
	PrintCompositorDrawPage(p, context, pageNr)
}
func (p *PrintCompositor) FooterFontName() string { return PrintCompositorGetFooterFontName(p) }
func (p *PrintCompositor) HeaderFontName() string { return PrintCompositorGetHeaderFontName(p) }
func (p *PrintCompositor) HighlightSyntax() bool  { return PrintCompositorGetHighlightSyntax(p) }
func (p *PrintCompositor) LeftMargin(unit G.Unit) float64 {
	return PrintCompositorGetLeftMargin(p, unit)
}
func (p *PrintCompositor) LineNumbersFontName() string {
	return PrintCompositorGetLineNumbersFontName(p)
}
func (p *PrintCompositor) NPages() int { return PrintCompositorGetNPages(p) }
func (p *PrintCompositor) Paginate(context *G.PrintContext) bool {
	return PrintCompositorPaginate(p, context)
}
func (p *PrintCompositor) PaginationProgress() float64 {
	return PrintCompositorGetPaginationProgress(p)
}
func (p *PrintCompositor) PrintFooter() bool      { return PrintCompositorGetPrintFooter(p) }
func (p *PrintCompositor) PrintHeader() bool      { return PrintCompositorGetPrintHeader(p) }
func (p *PrintCompositor) PrintLineNumbers() uint { return PrintCompositorGetPrintLineNumbers(p) }
func (p *PrintCompositor) RightMargin(unit G.Unit) float64 {
	return PrintCompositorGetRightMargin(p, unit)
}
func (p *PrintCompositor) SetBodyFontName(fontName string) {
	PrintCompositorSetBodyFontName(p, fontName)
}
func (p *PrintCompositor) SetBottomMargin(margin float64, unit G.Unit) {
	PrintCompositorSetBottomMargin(p, margin, unit)
}
func (p *PrintCompositor) SetFooterFontName(fontName string) {
	PrintCompositorSetFooterFontName(p, fontName)
}
func (p *PrintCompositor) SetFooterFormat(separator bool, left, center, right string) {
	PrintCompositorSetFooterFormat(p, separator, left, center, right)
}
func (p *PrintCompositor) SetHeaderFontName(fontName string) {
	PrintCompositorSetHeaderFontName(p, fontName)
}
func (p *PrintCompositor) SetHeaderFormat(separator bool, left, center, right string) {
	PrintCompositorSetHeaderFormat(p, separator, left, center, right)
}
func (p *PrintCompositor) SetHighlightSyntax(highlight bool) {
	PrintCompositorSetHighlightSyntax(p, highlight)
}
func (p *PrintCompositor) SetLeftMargin(margin float64, unit G.Unit) {
	PrintCompositorSetLeftMargin(p, margin, unit)
}
func (p *PrintCompositor) SetLineNumbersFontName(fontName string) {
	PrintCompositorSetLineNumbersFontName(p, fontName)
}
func (p *PrintCompositor) SetPrintFooter(prt bool) { PrintCompositorSetPrintFooter(p, prt) }
func (p *PrintCompositor) SetPrintHeader(prt bool) { PrintCompositorSetPrintHeader(p, prt) }
func (p *PrintCompositor) SetPrintLineNumbers(interval uint) {
	PrintCompositorSetPrintLineNumbers(p, interval)
}
func (p *PrintCompositor) SetRightMargin(margin float64, unit G.Unit) {
	PrintCompositorSetRightMargin(p, margin, unit)
}
func (p *PrintCompositor) SetTabWidth(width uint) { PrintCompositorSetTabWidth(p, width) }
func (p *PrintCompositor) SetTopMargin(margin float64, unit G.Unit) {
	PrintCompositorSetTopMargin(p, margin, unit)
}
func (p *PrintCompositor) SetWrapMode(wrapMode G.WrapMode) { PrintCompositorSetWrapMode(p, wrapMode) }
func (p *PrintCompositor) TabWidth() uint                  { return PrintCompositorGetTabWidth(p) }
func (p *PrintCompositor) TopMargin(unit G.Unit) float64 {
	return PrintCompositorGetTopMargin(p, unit)
}
func (p *PrintCompositor) WrapMode() G.WrapMode { return PrintCompositorGetWrapMode(p) }

type SearchFlags Enum

const (
	SEARCH_VISIBLE_ONLY SearchFlags = 1 << iota
	SEARCH_TEXT_ONLY
	SEARCH_CASE_INSENSITIVE
)

var SearchFlagsGetType func() O.Type

type SmartHomeEndType Enum

const (
	SMART_HOME_END_DISABLED SmartHomeEndType = iota
	SMART_HOME_END_BEFORE
	SMART_HOME_END_AFTER
	SMART_HOME_END_ALWAYS
)

var SmartHomeEndTypeGetType func() O.Type

type Style struct{}

var (
	StyleGetType func() O.Type

	StyleCopy func(s *Style) *Style
)

func (s *Style) Copy() *Style { return StyleCopy(s) }

type StyleScheme struct {
	Base O.Object
	_    *struct{}
}

var (
	StyleSchemeGetType func() O.Type

	StyleSchemeGetAuthors     func(s *StyleScheme) []string
	StyleSchemeGetDescription func(s *StyleScheme) string
	StyleSchemeGetFilename    func(s *StyleScheme) string
	StyleSchemeGetId          func(s *StyleScheme) string
	StyleSchemeGetName        func(s *StyleScheme) string
	StyleSchemeGetStyle       func(s *StyleScheme, styleId string) *Style
)

func (s *StyleScheme) Authors() []string           { return StyleSchemeGetAuthors(s) }
func (s *StyleScheme) Description() string         { return StyleSchemeGetDescription(s) }
func (s *StyleScheme) Filename() string            { return StyleSchemeGetFilename(s) }
func (s *StyleScheme) Id() string                  { return StyleSchemeGetId(s) }
func (s *StyleScheme) Name() string                { return StyleSchemeGetName(s) }
func (s *StyleScheme) Style(styleId string) *Style { return StyleSchemeGetStyle(s, styleId) }

type StyleSchemeManager struct {
	Parent O.Object
	_      *struct{}
}

var (
	StyleSchemeManagerGetType func() O.Type
	StyleSchemeManagerNew     func() *StyleSchemeManager

	StyleSchemeManagerGetDefault func() *StyleSchemeManager

	StyleSchemeManagerAppendSearchPath  func(s *StyleSchemeManager, path string)
	StyleSchemeManagerForceRescan       func(s *StyleSchemeManager)
	StyleSchemeManagerGetScheme         func(s *StyleSchemeManager, schemeId string) *StyleScheme
	StyleSchemeManagerGetSchemeIds      func(s *StyleSchemeManager) []string
	StyleSchemeManagerGetSearchPath     func(s *StyleSchemeManager) []string
	StyleSchemeManagerPrependSearchPath func(s *StyleSchemeManager, path string)
	StyleSchemeManagerSetSearchPath     func(s *StyleSchemeManager, path **T.Gchar)
)

func (s *StyleSchemeManager) AppendSearchPath(path string) {
	StyleSchemeManagerAppendSearchPath(s, path)
}
func (s *StyleSchemeManager) ForceRescan() { StyleSchemeManagerForceRescan(s) }
func (s *StyleSchemeManager) PrependSearchPath(path string) {
	StyleSchemeManagerPrependSearchPath(s, path)
}
func (s *StyleSchemeManager) Scheme(schemeId string) *StyleScheme {
	return StyleSchemeManagerGetScheme(s, schemeId)
}
func (s *StyleSchemeManager) SchemeIds() []string          { return StyleSchemeManagerGetSchemeIds(s) }
func (s *StyleSchemeManager) SearchPath() []string         { return StyleSchemeManagerGetSearchPath(s) }
func (s *StyleSchemeManager) SetSearchPath(path **T.Gchar) { StyleSchemeManagerSetSearchPath(s, path) }

type UndoManager struct{}

var (
	UndoManagerGetType func() O.Type

	UndoManagerBeginNotUndoableAction func(u *UndoManager)
	UndoManagerCanRedo                func(u *UndoManager) bool
	UndoManagerCanRedoChanged         func(u *UndoManager)
	UndoManagerCanUndo                func(u *UndoManager) bool
	UndoManagerCanUndoChanged         func(u *UndoManager)
	UndoManagerEndNotUndoableAction   func(u *UndoManager)
	UndoManagerRedo                   func(u *UndoManager)
	UndoManagerUndo                   func(u *UndoManager)
)

func (u *UndoManager) BeginNotUndoableAction() { UndoManagerBeginNotUndoableAction(u) }
func (u *UndoManager) CanRedo() bool           { return UndoManagerCanRedo(u) }
func (u *UndoManager) CanRedoChanged()         { UndoManagerCanRedoChanged(u) }
func (u *UndoManager) CanUndo() bool           { return UndoManagerCanUndo(u) }
func (u *UndoManager) CanUndoChanged()         { UndoManagerCanUndoChanged(u) }
func (u *UndoManager) EndNotUndoableAction()   { UndoManagerEndNotUndoableAction(u) }
func (u *UndoManager) Redo()                   { UndoManagerRedo(u) }
func (u *UndoManager) Undo()                   { UndoManagerUndo(u) }

type View struct {
	Parent G.TextView
	_      *struct{}
}

var (
	ViewGetType       func() O.Type
	ViewNew           func() *G.Widget
	ViewNewWithBuffer func(buffer *Buffer) *G.Widget

	ViewGetAutoIndent                    func(v *View) bool
	ViewGetCompletion                    func(v *View) *Completion
	ViewGetDrawSpaces                    func(v *View) DrawSpacesFlags
	ViewGetGutter                        func(v *View, windowType G.TextWindowType) *Gutter
	ViewGetHighlightCurrentLine          func(v *View) bool
	ViewGetIndentOnTab                   func(v *View) bool
	ViewGetIndentWidth                   func(v *View) int
	ViewGetInsertSpacesInsteadOfTabs     func(v *View) bool
	ViewGetMarkCategoryBackground        func(v *View, category string, dest *D.Color) bool
	ViewGetMarkCategoryPixbuf            func(v *View, category string) *D.Pixbuf
	ViewGetMarkCategoryPriority          func(v *View, category string) int
	ViewGetRightMarginPosition           func(v *View) uint
	ViewGetShowLineMarks                 func(v *View) bool
	ViewGetShowLineNumbers               func(v *View) bool
	ViewGetShowRightMargin               func(v *View) bool
	ViewGetSmartHomeEnd                  func(v *View) SmartHomeEndType
	ViewGetTabWidth                      func(v *View) uint
	ViewSetAutoIndent                    func(v *View, enable bool)
	ViewSetDrawSpaces                    func(v *View, flags DrawSpacesFlags)
	ViewSetHighlightCurrentLine          func(v *View, show bool)
	ViewSetIndentOnTab                   func(v *View, enable bool)
	ViewSetIndentWidth                   func(v *View, width int)
	ViewSetInsertSpacesInsteadOfTabs     func(v *View, enable bool)
	ViewSetMarkCategoryBackground        func(v *View, category string, color *D.Color)
	ViewSetMarkCategoryIconFromIconName  func(v *View, category, name string)
	ViewSetMarkCategoryIconFromPixbuf    func(v *View, category string, pixbuf *D.Pixbuf)
	ViewSetMarkCategoryIconFromStock     func(v *View, category, stockId string)
	ViewSetMarkCategoryPixbuf            func(v *View, category string, pixbuf *D.Pixbuf)
	ViewSetMarkCategoryPriority          func(v *View, category string, priority int)
	ViewSetMarkCategoryTooltipFunc       func(v *View, category string, f ViewMarkTooltipFunc, userData T.Gpointer, userDataNotify T.GDestroyNotify)
	ViewSetMarkCategoryTooltipMarkupFunc func(v *View, category string, markupFunc ViewMarkTooltipFunc, userData T.Gpointer, userDataNotify T.GDestroyNotify)
	ViewSetRightMarginPosition           func(v *View, pos uint)
	ViewSetShowLineMarks                 func(v *View, show bool)
	ViewSetShowLineNumbers               func(v *View, show bool)
	ViewSetShowRightMargin               func(v *View, show bool)
	ViewSetSmartHomeEnd                  func(v *View, smartHe SmartHomeEndType)
	ViewSetTabWidth                      func(v *View, width uint)
)

func (v *View) AutoIndent() bool                           { return ViewGetAutoIndent(v) }
func (v *View) Completion() *Completion                    { return ViewGetCompletion(v) }
func (v *View) DrawSpaces() DrawSpacesFlags                { return ViewGetDrawSpaces(v) }
func (v *View) Gutter(windowType G.TextWindowType) *Gutter { return ViewGetGutter(v, windowType) }
func (v *View) HighlightCurrentLine() bool                 { return ViewGetHighlightCurrentLine(v) }
func (v *View) IndentOnTab() bool                          { return ViewGetIndentOnTab(v) }
func (v *View) IndentWidth() int                           { return ViewGetIndentWidth(v) }
func (v *View) InsertSpacesInsteadOfTabs() bool            { return ViewGetInsertSpacesInsteadOfTabs(v) }
func (v *View) MarkCategoryBackground(category string, dest *D.Color) bool {
	return ViewGetMarkCategoryBackground(v, category, dest)
}
func (v *View) MarkCategoryPixbuf(category string) *D.Pixbuf {
	return ViewGetMarkCategoryPixbuf(v, category)
}
func (v *View) MarkCategoryPriority(category string) int {
	return ViewGetMarkCategoryPriority(v, category)
}
func (v *View) RightMarginPosition() uint           { return ViewGetRightMarginPosition(v) }
func (v *View) SetAutoIndent(enable bool)           { ViewSetAutoIndent(v, enable) }
func (v *View) SetDrawSpaces(flags DrawSpacesFlags) { ViewSetDrawSpaces(v, flags) }
func (v *View) SetHighlightCurrentLine(show bool)   { ViewSetHighlightCurrentLine(v, show) }
func (v *View) SetIndentOnTab(enable bool)          { ViewSetIndentOnTab(v, enable) }
func (v *View) SetIndentWidth(width int)            { ViewSetIndentWidth(v, width) }
func (v *View) SetInsertSpacesInsteadOfTabs(enable bool) {
	ViewSetInsertSpacesInsteadOfTabs(v, enable)
}
func (v *View) SetMarkCategoryBackground(category string, color *D.Color) {
	ViewSetMarkCategoryBackground(v, category, color)
}
func (v *View) SetMarkCategoryIconFromIconName(category, name string) {
	ViewSetMarkCategoryIconFromIconName(v, category, name)
}
func (v *View) SetMarkCategoryIconFromPixbuf(category string, pixbuf *D.Pixbuf) {
	ViewSetMarkCategoryIconFromPixbuf(v, category, pixbuf)
}
func (v *View) SetMarkCategoryIconFromStock(category, stockId string) {
	ViewSetMarkCategoryIconFromStock(v, category, stockId)
}
func (v *View) SetMarkCategoryPixbuf(category string, pixbuf *D.Pixbuf) {
	ViewSetMarkCategoryPixbuf(v, category, pixbuf)
}
func (v *View) SetMarkCategoryPriority(category string, priority int) {
	ViewSetMarkCategoryPriority(v, category, priority)
}
func (v *View) SetMarkCategoryTooltipFunc(category string, f ViewMarkTooltipFunc, userData T.Gpointer, userDataNotify T.GDestroyNotify) {
	ViewSetMarkCategoryTooltipFunc(v, category, f, userData, userDataNotify)
}
func (v *View) SetMarkCategoryTooltipMarkupFunc(category string, markupFunc ViewMarkTooltipFunc, userData T.Gpointer, userDataNotify T.GDestroyNotify) {
	ViewSetMarkCategoryTooltipMarkupFunc(v, category, markupFunc, userData, userDataNotify)
}
func (v *View) SetRightMarginPosition(pos uint)          { ViewSetRightMarginPosition(v, pos) }
func (v *View) SetShowLineMarks(show bool)               { ViewSetShowLineMarks(v, show) }
func (v *View) SetShowLineNumbers(show bool)             { ViewSetShowLineNumbers(v, show) }
func (v *View) SetShowRightMargin(show bool)             { ViewSetShowRightMargin(v, show) }
func (v *View) SetSmartHomeEnd(smartHe SmartHomeEndType) { ViewSetSmartHomeEnd(v, smartHe) }
func (v *View) SetTabWidth(width uint)                   { ViewSetTabWidth(v, width) }
func (v *View) ShowLineMarks() bool                      { return ViewGetShowLineMarks(v) }
func (v *View) ShowLineNumbers() bool                    { return ViewGetShowLineNumbers(v) }
func (v *View) ShowRightMargin() bool                    { return ViewGetShowRightMargin(v) }
func (v *View) SmartHomeEnd() SmartHomeEndType           { return ViewGetSmartHomeEnd(v) }
func (v *View) TabWidth() uint                           { return ViewGetTabWidth(v) }

var ViewGutterPositionGetType func() O.Type

type ViewMarkTooltipFunc func(
	Mark *Mark, userData T.Gpointer) string

var dll = "libgtksourceview-2.0-0.dll"

var apiList = outside.Apis{
	{"gtk_source_buffer_backward_iter_to_source_mark", &BufferBackwardIterToSourceMark},
	{"gtk_source_buffer_begin_not_undoable_action", &BufferBeginNotUndoableAction},
	{"gtk_source_buffer_can_redo", &BufferCanRedo},
	{"gtk_source_buffer_can_undo", &BufferCanUndo},
	{"gtk_source_buffer_create_source_mark", &BufferCreateSourceMark},
	{"gtk_source_buffer_end_not_undoable_action", &BufferEndNotUndoableAction},
	{"gtk_source_buffer_ensure_highlight", &BufferEnsureHighlight},
	{"gtk_source_buffer_forward_iter_to_source_mark", &BufferForwardIterToSourceMark},
	{"gtk_source_buffer_get_context_classes_at_iter", &BufferGetContextClassesAtIter},
	{"gtk_source_buffer_get_highlight_matching_brackets", &BufferGetHighlightMatchingBrackets},
	{"gtk_source_buffer_get_highlight_syntax", &BufferGetHighlightSyntax},
	{"gtk_source_buffer_get_language", &BufferGetLanguage},
	{"gtk_source_buffer_get_max_undo_levels", &BufferGetMaxUndoLevels},
	{"gtk_source_buffer_get_source_marks_at_iter", &BufferGetSourceMarksAtIter},
	{"gtk_source_buffer_get_source_marks_at_line", &BufferGetSourceMarksAtLine},
	{"gtk_source_buffer_get_style_scheme", &BufferGetStyleScheme},
	{"gtk_source_buffer_get_type", &BufferGetType},
	{"gtk_source_buffer_get_undo_manager", &BufferGetUndoManager},
	{"gtk_source_buffer_iter_backward_to_context_class_toggle", &BufferIterBackwardToContextClassToggle},
	{"gtk_source_buffer_iter_forward_to_context_class_toggle", &BufferIterForwardToContextClassToggle},
	{"gtk_source_buffer_iter_has_context_class", &BufferIterHasContextClass},
	{"gtk_source_buffer_new", &BufferNew},
	{"gtk_source_buffer_new_with_language", &BufferNewWithLanguage},
	{"gtk_source_buffer_redo", &BufferRedo},
	{"gtk_source_buffer_remove_source_marks", &BufferRemoveSourceMarks},
	{"gtk_source_buffer_set_highlight_matching_brackets", &BufferSetHighlightMatchingBrackets},
	{"gtk_source_buffer_set_highlight_syntax", &BufferSetHighlightSyntax},
	{"gtk_source_buffer_set_language", &BufferSetLanguage},
	{"gtk_source_buffer_set_max_undo_levels", &BufferSetMaxUndoLevels},
	{"gtk_source_buffer_set_style_scheme", &BufferSetStyleScheme},
	{"gtk_source_buffer_set_undo_manager", &BufferSetUndoManager},
	{"gtk_source_buffer_undo", &BufferUndo},
	{"gtk_source_completion_activation_get_type", &CompletionActivationGetType},
	{"gtk_source_completion_add_provider", &CompletionAddProvider},
	{"gtk_source_completion_block_interactive", &CompletionBlockInteractive},
	{"gtk_source_completion_context_add_proposals", &CompletionContextAddProposals},
	{"gtk_source_completion_context_get_activation", &CompletionContextGetActivation},
	{"gtk_source_completion_context_get_iter", &CompletionContextGetIter},
	{"gtk_source_completion_context_get_type", &CompletionContextGetType},
	{"gtk_source_completion_create_context", &CompletionCreateContext},
	{"gtk_source_completion_error_get_type", &CompletionErrorGetType},
	{"gtk_source_completion_error_quark", &CompletionErrorQuark},
	{"gtk_source_completion_get_info_window", &CompletionGetInfoWindow},
	{"gtk_source_completion_get_providers", &CompletionGetProviders},
	{"gtk_source_completion_get_type", &CompletionGetType},
	{"gtk_source_completion_get_view", &CompletionGetView},
	{"gtk_source_completion_hide", &CompletionHide},
	{"gtk_source_completion_info_get_type", &CompletionInfoGetType},
	{"gtk_source_completion_info_get_widget", &CompletionInfoGetWidget},
	{"gtk_source_completion_info_move_to_iter", &CompletionInfoMoveToIter},
	{"gtk_source_completion_info_new", &CompletionInfoNew},
	{"gtk_source_completion_info_process_resize", &CompletionInfoProcessResize},
	{"gtk_source_completion_info_set_sizing", &CompletionInfoSetSizing},
	{"gtk_source_completion_info_set_widget", &CompletionInfoSetWidget},
	{"gtk_source_completion_item_get_type", &CompletionItemGetType},
	{"gtk_source_completion_item_new", &CompletionItemNew},
	{"gtk_source_completion_item_new_from_stock", &CompletionItemNewFromStock},
	{"gtk_source_completion_item_new_with_markup", &CompletionItemNewWithMarkup},
	// Undocumented {"gtk_source_completion_model_append", &CompletionModelAppend},
	// Undocumented {"gtk_source_completion_model_begin", &CompletionModelBegin},
	// Undocumented {"gtk_source_completion_model_cancel", &CompletionModelCancel},
	// Undocumented {"gtk_source_completion_model_clear", &CompletionModelClear},
	// Undocumented {"gtk_source_completion_model_end", &CompletionModelEnd},
	// Undocumented {"gtk_source_completion_model_get_providers", &CompletionModelGetProviders},
	// Undocumented {"gtk_source_completion_model_get_type", &CompletionModelGetType},
	// Undocumented {"gtk_source_completion_model_get_visible_providers", &CompletionModelGetVisibleProviders},
	// Undocumented {"gtk_source_completion_model_is_empty", &CompletionModelIsEmpty},
	// Undocumented {"gtk_source_completion_model_iter_equal", &CompletionModelIterEqual},
	// Undocumented {"gtk_source_completion_model_iter_is_header", &CompletionModelIterIsHeader},
	// Undocumented {"gtk_source_completion_model_iter_last", &CompletionModelIterLast},
	// Undocumented {"gtk_source_completion_model_iter_previous", &CompletionModelIterPrevious},
	// Undocumented {"gtk_source_completion_model_n_proposals", &CompletionModelNProposals},
	// Undocumented {"gtk_source_completion_model_new", &CompletionModelNew},
	// Undocumented {"gtk_source_completion_model_set_show_headers", &CompletionModelSetShowHeaders},
	// Undocumented {"gtk_source_completion_model_set_visible_providers", &CompletionModelSetVisibleProviders},
	{"gtk_source_completion_move_window", &CompletionMoveWindow},
	// Undocumented {"gtk_source_completion_new", &CompletionNew},
	{"gtk_source_completion_proposal_changed", &CompletionProposalChanged},
	{"gtk_source_completion_proposal_equal", &CompletionProposalEqual},
	{"gtk_source_completion_proposal_get_icon", &CompletionProposalGetIcon},
	{"gtk_source_completion_proposal_get_info", &CompletionProposalGetInfo},
	{"gtk_source_completion_proposal_get_label", &CompletionProposalGetLabel},
	{"gtk_source_completion_proposal_get_markup", &CompletionProposalGetMarkup},
	{"gtk_source_completion_proposal_get_text", &CompletionProposalGetText},
	{"gtk_source_completion_proposal_get_type", &CompletionProposalGetType},
	{"gtk_source_completion_proposal_hash", &CompletionProposalHash},
	{"gtk_source_completion_provider_activate_proposal", &CompletionProviderActivateProposal},
	{"gtk_source_completion_provider_get_activation", &CompletionProviderGetActivation},
	{"gtk_source_completion_provider_get_icon", &CompletionProviderGetIcon},
	{"gtk_source_completion_provider_get_info_widget", &CompletionProviderGetInfoWidget},
	{"gtk_source_completion_provider_get_interactive_delay", &CompletionProviderGetInteractiveDelay},
	{"gtk_source_completion_provider_get_name", &CompletionProviderGetName},
	{"gtk_source_completion_provider_get_priority", &CompletionProviderGetPriority},
	{"gtk_source_completion_provider_get_start_iter", &CompletionProviderGetStartIter},
	{"gtk_source_completion_provider_get_type", &CompletionProviderGetType},
	{"gtk_source_completion_provider_match", &CompletionProviderMatch},
	{"gtk_source_completion_provider_populate", &CompletionProviderPopulate},
	{"gtk_source_completion_provider_update_info", &CompletionProviderUpdateInfo},
	{"gtk_source_completion_remove_provider", &CompletionRemoveProvider},
	{"gtk_source_completion_show", &CompletionShow},
	{"gtk_source_completion_unblock_interactive", &CompletionUnblockInteractive},
	// Undocumented {"gtk_source_completion_utils_get_word", &CompletionUtilsGetWord},
	// Undocumented {"gtk_source_completion_utils_get_word_iter", &CompletionUtilsGetWordIter},
	// Undocumented {"gtk_source_completion_utils_is_separator", &CompletionUtilsIsSeparator},
	// Undocumented {"gtk_source_completion_utils_move_to_cursor", &CompletionUtilsMoveToCursor},
	// Undocumented {"gtk_source_completion_utils_move_to_iter", &CompletionUtilsMoveToIter},
	// Undocumented {"gtk_source_completion_utils_replace_current_word", &CompletionUtilsReplaceCurrentWord},
	// Undocumented {"gtk_source_completion_utils_replace_word", &CompletionUtilsReplaceWord},
	// Undocumented {"gtk_source_completion_words_buffer_get_buffer", &CompletionWordsBufferGetBuffer},
	// Undocumented {"gtk_source_completion_words_buffer_get_mark", &CompletionWordsBufferGetMark},
	// Undocumented {"gtk_source_completion_words_buffer_get_type", &CompletionWordsBufferGetType},
	// Undocumented{"gtk_source_completion_words_buffer_new", &CompletionWordsBufferNew},
	// Undocumented{"gtk_source_completion_words_buffer_set_minimum_word_size", &CompletionWordsBufferSetMinimumWordSize},
	// Undocumented{"gtk_source_completion_words_buffer_set_scan_batch_size", &CompletionWordsBufferSetScanBatchSize},
	{"gtk_source_completion_words_get_type", &CompletionWordsGetType},
	// Undocumented {"gtk_source_completion_words_library_add_word", &CompletionWordsLibraryAddWord},
	// Undocumented {"gtk_source_completion_words_library_find", &CompletionWordsLibraryFind},
	// Undocumented {"gtk_source_completion_words_library_find_first", &CompletionWordsLibrary_findFirst},
	// Undocumented {"gtk_source_completion_words_library_find_next", &CompletionWordsLibraryFindNext},
	// Undocumented {"gtk_source_completion_words_library_get_proposal", &CompletionWordsLibraryGetProposal},
	// Undocumented {"gtk_source_completion_words_library_get_type", &CompletionWordsLibraryGetType},
	// Undocumented {"gtk_source_completion_words_library_is_locked", &CompletionWordsLibraryIsLocked},
	// Undocumented {"gtk_source_completion_words_library_lock", &CompletionWordsLibraryLock},
	// Undocumented {"gtk_source_completion_words_library_new", &CompletionWordsLibraryNew},
	// Undocumented {"gtk_source_completion_words_library_remove_word", &CompletionWordsLibraryRemoveWord},
	// Undocumented {"gtk_source_completion_words_library_unlock", &CompletionWordsLibraryUnlock},
	{"gtk_source_completion_words_new", &CompletionWordsNew},
	// Undocumented {"gtk_source_completion_words_proposal_get_type", &CompletionWordsProposalGetType},
	// Undocumented {"gtk_source_completion_words_proposal_get_word", &CompletionWordsProposalGetWord},
	// Undocumented {"gtk_source_completion_words_proposal_new", &CompletionWordsProposalNew},
	// Undocumented {"gtk_source_completion_words_proposal_unuse", &CompletionWordsProposalUnuse},
	// Undocumented {"gtk_source_completion_words_proposal_use", &CompletionWordsProposalUse},
	{"gtk_source_completion_words_register", &CompletionWordsRegister},
	{"gtk_source_completion_words_unregister", &CompletionWordsUnregister},
	// Undocumented {"gtk_source_completion_words_utils_backward_word_start", &CompletionWordsUtilsBackwardWordStart},
	// Undocumented {"gtk_source_completion_words_utils_forward_word_end", &CompletionWordsUtilsForwardWordEnd},
	// Undocumented {"gtk_source_context_class_free", &ContextClassFree},
	// Undocumented {"gtk_source_context_class_new", &ContextClassNew},
	{"gtk_source_draw_spaces_flags_get_type", &DrawSpacesFlagsGetType},
	{"gtk_source_gutter_get_type", &GutterGetType},
	{"gtk_source_gutter_get_window", &GutterGetWindow},
	{"gtk_source_gutter_insert", &GutterInsert},
	// Undocumented {"gtk_source_gutter_new", &GutterNew},
	{"gtk_source_gutter_queue_draw", &GutterQueueDraw},
	{"gtk_source_gutter_remove", &GutterRemove},
	{"gtk_source_gutter_reorder", &GutterReorder},
	{"gtk_source_gutter_set_cell_data_func", &GutterSetCellDataFunc},
	{"gtk_source_gutter_set_cell_size_func", &GutterSetCellSizeFunc},
	{"gtk_source_iter_backward_search", &IterBackwardSearch},
	{"gtk_source_iter_forward_search", &IterForwardSearch},
	{"gtk_source_language_get_globs", &LanguageGetGlobs},
	{"gtk_source_language_get_hidden", &LanguageGetHidden},
	{"gtk_source_language_get_id", &LanguageGetId},
	{"gtk_source_language_get_metadata", &LanguageGetMetadata},
	{"gtk_source_language_get_mime_types", &LanguageGetMimeTypes},
	{"gtk_source_language_get_name", &LanguageGetName},
	{"gtk_source_language_get_section", &LanguageGetSection},
	{"gtk_source_language_get_style_ids", &LanguageGetStyleIds},
	{"gtk_source_language_get_style_name", &LanguageGetStyleName},
	{"gtk_source_language_get_type", &LanguageGetType},
	{"gtk_source_language_manager_get_default", &LanguageManagerGetDefault},
	{"gtk_source_language_manager_get_language", &LanguageManagerGetLanguage},
	{"gtk_source_language_manager_get_language_ids", &LanguageManagerGetLanguageIds},
	{"gtk_source_language_manager_get_search_path", &LanguageManagerGetSearchPath},
	{"gtk_source_language_manager_get_type", &LanguageManagerGetType},
	{"gtk_source_language_manager_guess_language", &LanguageManagerGuessLanguage},
	{"gtk_source_language_manager_new", &LanguageManagerNew},
	{"gtk_source_language_manager_set_search_path", &LanguageManagerSetSearchPath},
	{"gtk_source_mark_get_category", &MarkGetCategory},
	{"gtk_source_mark_get_type", &MarkGetType},
	{"gtk_source_mark_new", &MarkNew},
	{"gtk_source_mark_next", &MarkNext},
	{"gtk_source_mark_prev", &MarkPrev},
	{"gtk_source_print_compositor_draw_page", &PrintCompositorDrawPage},
	{"gtk_source_print_compositor_get_body_font_name", &PrintCompositorGetBodyFontName},
	{"gtk_source_print_compositor_get_bottom_margin", &PrintCompositorGetBottomMargin},
	{"gtk_source_print_compositor_get_buffer", &PrintCompositorGetBuffer},
	{"gtk_source_print_compositor_get_footer_font_name", &PrintCompositorGetFooterFontName},
	{"gtk_source_print_compositor_get_header_font_name", &PrintCompositorGetHeaderFontName},
	{"gtk_source_print_compositor_get_highlight_syntax", &PrintCompositorGetHighlightSyntax},
	{"gtk_source_print_compositor_get_left_margin", &PrintCompositorGetLeftMargin},
	{"gtk_source_print_compositor_get_line_numbers_font_name", &PrintCompositorGetLineNumbersFontName},
	{"gtk_source_print_compositor_get_n_pages", &PrintCompositorGetNPages},
	{"gtk_source_print_compositor_get_pagination_progress", &PrintCompositorGetPaginationProgress},
	{"gtk_source_print_compositor_get_print_footer", &PrintCompositorGetPrintFooter},
	{"gtk_source_print_compositor_get_print_header", &PrintCompositorGetPrintHeader},
	{"gtk_source_print_compositor_get_print_line_numbers", &PrintCompositorGetPrintLineNumbers},
	{"gtk_source_print_compositor_get_right_margin", &PrintCompositorGetRightMargin},
	{"gtk_source_print_compositor_get_tab_width", &PrintCompositorGetTabWidth},
	{"gtk_source_print_compositor_get_top_margin", &PrintCompositorGetTopMargin},
	{"gtk_source_print_compositor_get_type", &PrintCompositorGetType},
	{"gtk_source_print_compositor_get_wrap_mode", &PrintCompositorGetWrapMode},
	{"gtk_source_print_compositor_new", &PrintCompositorNew},
	{"gtk_source_print_compositor_new_from_view", &PrintCompositorNewFromView},
	{"gtk_source_print_compositor_paginate", &PrintCompositorPaginate},
	{"gtk_source_print_compositor_set_body_font_name", &PrintCompositorSetBodyFontName},
	{"gtk_source_print_compositor_set_bottom_margin", &PrintCompositorSetBottomMargin},
	{"gtk_source_print_compositor_set_footer_font_name", &PrintCompositorSetFooterFontName},
	{"gtk_source_print_compositor_set_footer_format", &PrintCompositorSetFooterFormat},
	{"gtk_source_print_compositor_set_header_font_name", &PrintCompositorSetHeaderFontName},
	{"gtk_source_print_compositor_set_header_format", &PrintCompositorSetHeaderFormat},
	{"gtk_source_print_compositor_set_highlight_syntax", &PrintCompositorSetHighlightSyntax},
	{"gtk_source_print_compositor_set_left_margin", &PrintCompositorSetLeftMargin},
	{"gtk_source_print_compositor_set_line_numbers_font_name", &PrintCompositorSetLineNumbersFontName},
	{"gtk_source_print_compositor_set_print_footer", &PrintCompositorSetPrintFooter},
	{"gtk_source_print_compositor_set_print_header", &PrintCompositorSetPrintHeader},
	{"gtk_source_print_compositor_set_print_line_numbers", &PrintCompositorSetPrintLineNumbers},
	{"gtk_source_print_compositor_set_right_margin", &PrintCompositorSetRightMargin},
	{"gtk_source_print_compositor_set_tab_width", &PrintCompositorSetTabWidth},
	{"gtk_source_print_compositor_set_top_margin", &PrintCompositorSetTopMargin},
	{"gtk_source_print_compositor_set_wrap_mode", &PrintCompositorSetWrapMode},
	{"gtk_source_search_flags_get_type", &SearchFlagsGetType},
	{"gtk_source_smart_home_end_type_get_type", &SmartHomeEndTypeGetType},
	{"gtk_source_style_copy", &StyleCopy},
	{"gtk_source_style_get_type", &StyleGetType},
	{"gtk_source_style_scheme_get_authors", &StyleSchemeGetAuthors},
	{"gtk_source_style_scheme_get_description", &StyleSchemeGetDescription},
	{"gtk_source_style_scheme_get_filename", &StyleSchemeGetFilename},
	{"gtk_source_style_scheme_get_id", &StyleSchemeGetId},
	{"gtk_source_style_scheme_get_name", &StyleSchemeGetName},
	{"gtk_source_style_scheme_get_style", &StyleSchemeGetStyle},
	{"gtk_source_style_scheme_get_type", &StyleSchemeGetType},
	{"gtk_source_style_scheme_manager_append_search_path", &StyleSchemeManagerAppendSearchPath},
	{"gtk_source_style_scheme_manager_force_rescan", &StyleSchemeManagerForceRescan},
	{"gtk_source_style_scheme_manager_get_default", &StyleSchemeManagerGetDefault},
	{"gtk_source_style_scheme_manager_get_scheme", &StyleSchemeManagerGetScheme},
	{"gtk_source_style_scheme_manager_get_scheme_ids", &StyleSchemeManagerGetSchemeIds},
	{"gtk_source_style_scheme_manager_get_search_path", &StyleSchemeManagerGetSearchPath},
	{"gtk_source_style_scheme_manager_get_type", &StyleSchemeManagerGetType},
	{"gtk_source_style_scheme_manager_new", &StyleSchemeManagerNew},
	{"gtk_source_style_scheme_manager_prepend_search_path", &StyleSchemeManagerPrependSearchPath},
	{"gtk_source_style_scheme_manager_set_search_path", &StyleSchemeManagerSetSearchPath},
	{"gtk_source_undo_manager_begin_not_undoable_action", &UndoManagerBeginNotUndoableAction},
	{"gtk_source_undo_manager_can_redo", &UndoManagerCanRedo},
	{"gtk_source_undo_manager_can_redo_changed", &UndoManagerCanRedoChanged},
	{"gtk_source_undo_manager_can_undo", &UndoManagerCanUndo},
	{"gtk_source_undo_manager_can_undo_changed", &UndoManagerCanUndoChanged},
	// Undocumented {"gtk_source_undo_manager_default_get_type", &UndoManagerDefaultGetType},
	// Undocumented {"gtk_source_undo_manager_default_set_max_undo_levels", &UndoManagerDefaultSetMaxUndoLevels},
	{"gtk_source_undo_manager_end_not_undoable_action", &UndoManagerEndNotUndoableAction},
	{"gtk_source_undo_manager_get_type", &UndoManagerGetType},
	{"gtk_source_undo_manager_redo", &UndoManagerRedo},
	{"gtk_source_undo_manager_undo", &UndoManagerUndo},
	{"gtk_source_view_get_auto_indent", &ViewGetAutoIndent},
	{"gtk_source_view_get_completion", &ViewGetCompletion},
	{"gtk_source_view_get_draw_spaces", &ViewGetDrawSpaces},
	{"gtk_source_view_get_gutter", &ViewGetGutter},
	{"gtk_source_view_get_highlight_current_line", &ViewGetHighlightCurrentLine},
	{"gtk_source_view_get_indent_on_tab", &ViewGetIndentOnTab},
	{"gtk_source_view_get_indent_width", &ViewGetIndentWidth},
	{"gtk_source_view_get_insert_spaces_instead_of_tabs", &ViewGetInsertSpacesInsteadOfTabs},
	{"gtk_source_view_get_mark_category_background", &ViewGetMarkCategoryBackground},
	{"gtk_source_view_get_mark_category_pixbuf", &ViewGetMarkCategoryPixbuf},
	{"gtk_source_view_get_mark_category_priority", &ViewGetMarkCategoryPriority},
	{"gtk_source_view_get_right_margin_position", &ViewGetRightMarginPosition},
	{"gtk_source_view_get_show_line_marks", &ViewGetShowLineMarks},
	{"gtk_source_view_get_show_line_numbers", &ViewGetShowLineNumbers},
	{"gtk_source_view_get_show_right_margin", &ViewGetShowRightMargin},
	{"gtk_source_view_get_smart_home_end", &ViewGetSmartHomeEnd},
	{"gtk_source_view_get_tab_width", &ViewGetTabWidth},
	{"gtk_source_view_get_type", &ViewGetType},
	{"gtk_source_view_gutter_position_get_type", &ViewGutterPositionGetType},
	{"gtk_source_view_new", &ViewNew},
	{"gtk_source_view_new_with_buffer", &ViewNewWithBuffer},
	{"gtk_source_view_set_auto_indent", &ViewSetAutoIndent},
	{"gtk_source_view_set_draw_spaces", &ViewSetDrawSpaces},
	{"gtk_source_view_set_highlight_current_line", &ViewSetHighlightCurrentLine},
	{"gtk_source_view_set_indent_on_tab", &ViewSetIndentOnTab},
	{"gtk_source_view_set_indent_width", &ViewSetIndentWidth},
	{"gtk_source_view_set_insert_spaces_instead_of_tabs", &ViewSetInsertSpacesInsteadOfTabs},
	{"gtk_source_view_set_mark_category_background", &ViewSetMarkCategoryBackground},
	{"gtk_source_view_set_mark_category_icon_from_icon_name", &ViewSetMarkCategoryIconFromIconName},
	{"gtk_source_view_set_mark_category_icon_from_pixbuf", &ViewSetMarkCategoryIconFromPixbuf},
	{"gtk_source_view_set_mark_category_icon_from_stock", &ViewSetMarkCategoryIconFromStock},
	{"gtk_source_view_set_mark_category_pixbuf", &ViewSetMarkCategoryPixbuf},
	{"gtk_source_view_set_mark_category_priority", &ViewSetMarkCategoryPriority},
	{"gtk_source_view_set_mark_category_tooltip_func", &ViewSetMarkCategoryTooltipFunc},
	{"gtk_source_view_set_mark_category_tooltip_markup_func", &ViewSetMarkCategoryTooltipMarkupFunc},
	{"gtk_source_view_set_right_margin_position", &ViewSetRightMarginPosition},
	{"gtk_source_view_set_show_line_marks", &ViewSetShowLineMarks},
	{"gtk_source_view_set_show_line_numbers", &ViewSetShowLineNumbers},
	{"gtk_source_view_set_show_right_margin", &ViewSetShowRightMargin},
	{"gtk_source_view_set_smart_home_end", &ViewSetSmartHomeEnd},
	{"gtk_source_view_set_tab_width", &ViewSetTabWidth},
}
