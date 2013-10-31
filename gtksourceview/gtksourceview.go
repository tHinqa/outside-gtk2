// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gtksourceview provides API definitions for accessing
//libgtksourceview-2.0-0.dll.
package gtksourceview

import (
	"github.com/tHinqa/outside"
	D "github.com/tHinqa/outside-gtk2/gdk"
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

	bufferBackwardIterToSourceMark         func(b *Buffer, iter *G.TextIter, category string) bool
	bufferBeginNotUndoableAction           func(b *Buffer)
	bufferCanRedo                          func(b *Buffer) bool
	bufferCanUndo                          func(b *Buffer) bool
	bufferCreateSourceMark                 func(b *Buffer, name, category string, where *G.TextIter) *Mark
	bufferEndNotUndoableAction             func(b *Buffer)
	bufferEnsureHighlight                  func(b *Buffer, start, end *G.TextIter)
	bufferForwardIterToSourceMark          func(b *Buffer, iter *G.TextIter, category string) bool
	bufferGetContextClassesAtIter          func(b *Buffer, iter *G.TextIter) []string
	bufferGetHighlightMatchingBrackets     func(b *Buffer) bool
	bufferGetHighlightSyntax               func(b *Buffer) bool
	bufferGetLanguage                      func(b *Buffer) *Language
	bufferGetMaxUndoLevels                 func(b *Buffer) int
	bufferGetSourceMarksAtIter             func(b *Buffer, iter *G.TextIter, category string) *T.GSList
	bufferGetSourceMarksAtLine             func(b *Buffer, line int, category string) *T.GSList
	bufferGetStyleScheme                   func(b *Buffer) *StyleScheme
	bufferGetUndoManager                   func(b *Buffer) *UndoManager
	bufferIterBackwardToContextClassToggle func(b *Buffer, iter *G.TextIter, contextClass string) bool
	bufferIterForwardToContextClassToggle  func(b *Buffer, iter *G.TextIter, contextClass string) bool
	bufferIterHasContextClass              func(b *Buffer, iter *G.TextIter, contextClass string) bool
	bufferRedo                             func(b *Buffer)
	bufferRemoveSourceMarks                func(b *Buffer, start *G.TextIter, end *G.TextIter, category string)
	bufferSetHighlightMatchingBrackets     func(b *Buffer, highlight bool)
	bufferSetHighlightSyntax               func(b *Buffer, highlight bool)
	bufferSetLanguage                      func(b *Buffer, language *Language)
	bufferSetMaxUndoLevels                 func(b *Buffer, maxUndoLevels int)
	bufferSetStyleScheme                   func(b *Buffer, scheme *StyleScheme)
	bufferSetUndoManager                   func(b *Buffer, manager *UndoManager)
	bufferUndo                             func(b *Buffer)
)

func (b *Buffer) BackwardIterToSourceMark(iter *G.TextIter, category string) bool {
	return bufferBackwardIterToSourceMark(b, iter, category)
}
func (b *Buffer) BeginNotUndoableAction() { bufferBeginNotUndoableAction(b) }
func (b *Buffer) CanRedo() bool           { return bufferCanRedo(b) }
func (b *Buffer) CanUndo() bool           { return bufferCanUndo(b) }
func (b *Buffer) ContextClassesAtIter(iter *G.TextIter) []string {
	return bufferGetContextClassesAtIter(b, iter)
}
func (b *Buffer) CreateSourceMark(name, category string, where *G.TextIter) *Mark {
	return bufferCreateSourceMark(b, name, category, where)
}
func (b *Buffer) EndNotUndoableAction()                  { bufferEndNotUndoableAction(b) }
func (b *Buffer) EnsureHighlight(start, end *G.TextIter) { bufferEnsureHighlight(b, start, end) }
func (b *Buffer) ForwardIterToSourceMark(iter *G.TextIter, category string) bool {
	return bufferForwardIterToSourceMark(b, iter, category)
}
func (b *Buffer) HighlightMatchingBrackets() bool {
	return bufferGetHighlightMatchingBrackets(b)
}
func (b *Buffer) HighlightSyntax() bool { return bufferGetHighlightSyntax(b) }
func (b *Buffer) IterBackwardToContextClassToggle(iter *G.TextIter, contextClass string) bool {
	return bufferIterBackwardToContextClassToggle(b, iter, contextClass)
}
func (b *Buffer) IterForwardToContextClassToggle(iter *G.TextIter, contextClass string) bool {
	return bufferIterForwardToContextClassToggle(b, iter, contextClass)
}
func (b *Buffer) IterHasContextClass(iter *G.TextIter, contextClass string) bool {
	return bufferIterHasContextClass(b, iter, contextClass)
}
func (b *Buffer) Language() *Language { return bufferGetLanguage(b) }
func (b *Buffer) MaxUndoLevels() int  { return bufferGetMaxUndoLevels(b) }
func (b *Buffer) Redo()               { bufferRedo(b) }
func (b *Buffer) RemoveSourceMarks(start *G.TextIter, end *G.TextIter, category string) {
	bufferRemoveSourceMarks(b, start, end, category)
}
func (b *Buffer) SetHighlightMatchingBrackets(highlight bool) {
	bufferSetHighlightMatchingBrackets(b, highlight)
}
func (b *Buffer) SetHighlightSyntax(highlight bool)   { bufferSetHighlightSyntax(b, highlight) }
func (b *Buffer) SetLanguage(language *Language)      { bufferSetLanguage(b, language) }
func (b *Buffer) SetMaxUndoLevels(maxUndoLevels int)  { bufferSetMaxUndoLevels(b, maxUndoLevels) }
func (b *Buffer) SetStyleScheme(scheme *StyleScheme)  { bufferSetStyleScheme(b, scheme) }
func (b *Buffer) SetUndoManager(manager *UndoManager) { bufferSetUndoManager(b, manager) }
func (b *Buffer) SourceMarksAtIter(iter *G.TextIter, category string) *T.GSList {
	return bufferGetSourceMarksAtIter(b, iter, category)
}
func (b *Buffer) SourceMarksAtLine(line int, category string) *T.GSList {
	return bufferGetSourceMarksAtLine(b, line, category)
}
func (b *Buffer) StyleScheme() *StyleScheme { return bufferGetStyleScheme(b) }
func (b *Buffer) Undo()                     { bufferUndo(b) }
func (b *Buffer) UndoManager() *UndoManager { return bufferGetUndoManager(b) }

type Completion struct {
	Parent G.Object
	_      *struct{}
}

var (
	CompletionGetType    func() O.Type
	CompletionErrorQuark func() T.GQuark

	completionAddProvider        func(c *Completion, provider *CompletionProvider, err **T.GError) bool
	completionBlockInteractive   func(c *Completion)
	completionCreateContext      func(c *Completion, position *G.TextIter) *CompletionContext
	completionGetInfoWindow      func(c *Completion) *CompletionInfo
	completionGetProviders       func(c *Completion) *T.GList
	completionGetView            func(c *Completion) *View
	completionHide               func(c *Completion)
	completionMoveWindow         func(c *Completion, iter *G.TextIter)
	completionRemoveProvider     func(c *Completion, provider *CompletionProvider, err **T.GError) bool
	completionShow               func(c *Completion, providers *T.GList, context *CompletionContext) bool
	completionUnblockInteractive func(c *Completion)
)

func (c *Completion) AddProvider(provider *CompletionProvider, err **T.GError) bool {
	return completionAddProvider(c, provider, err)
}
func (c *Completion) BlockInteractive() { completionBlockInteractive(c) }
func (c *Completion) CreateContext(position *G.TextIter) *CompletionContext {
	return completionCreateContext(c, position)
}
func (c *Completion) Hide()                       { completionHide(c) }
func (c *Completion) InfoWindow() *CompletionInfo { return completionGetInfoWindow(c) }
func (c *Completion) MoveWindow(iter *G.TextIter) { completionMoveWindow(c, iter) }
func (c *Completion) Providers() *T.GList         { return completionGetProviders(c) }
func (c *Completion) RemoveProvider(provider *CompletionProvider, err **T.GError) bool {
	return completionRemoveProvider(c, provider, err)
}
func (c *Completion) Show(providers *T.GList, context *CompletionContext) bool {
	return completionShow(c, providers, context)
}
func (c *Completion) UnblockInteractive() { completionUnblockInteractive(c) }
func (c *Completion) View() *View         { return completionGetView(c) }

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

	completionContextAddProposals  func(c *CompletionContext, provider *CompletionProvider, proposals *T.GList, finished bool)
	completionContextGetActivation func(c *CompletionContext) CompletionActivation
	completionContextGetIter       func(c *CompletionContext, iter *G.TextIter)
)

func (c *CompletionContext) Activation() CompletionActivation {
	return completionContextGetActivation(c)
}
func (c *CompletionContext) AddProposals(provider *CompletionProvider, proposals *T.GList, finished bool) {
	completionContextAddProposals(c, provider, proposals, finished)
}
func (c *CompletionContext) Iter(iter *G.TextIter) { completionContextGetIter(c, iter) }

var CompletionErrorGetType func() O.Type

type CompletionInfo struct {
	Parent G.Window
	_      *struct{}
}

var (
	CompletionInfoGetType func() O.Type
	CompletionInfoNew     func() *CompletionInfo

	completionInfoGetWidget     func(c *CompletionInfo) *G.Widget
	completionInfoMoveToIter    func(c *CompletionInfo, view *G.TextView, iter *G.TextIter)
	completionInfoProcessResize func(c *CompletionInfo)
	completionInfoSetSizing     func(c *CompletionInfo, width, height int, shrinkWidth bool, shrinkHeight bool)
	completionInfoSetWidget     func(c *CompletionInfo, widget *G.Widget)
)

func (c *CompletionInfo) MoveToIter(view *G.TextView, iter *G.TextIter) {
	completionInfoMoveToIter(c, view, iter)
}
func (c *CompletionInfo) ProcessResize() { completionInfoProcessResize(c) }
func (c *CompletionInfo) SetSizing(width, height int, shrinkWidth bool, shrinkHeight bool) {
	completionInfoSetSizing(c, width, height, shrinkWidth, shrinkHeight)
}
func (c *CompletionInfo) SetWidget(widget *G.Widget) { completionInfoSetWidget(c, widget) }
func (c *CompletionInfo) Widget() *G.Widget          { return completionInfoGetWidget(c) }

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

	completionProposalChanged   func(c *CompletionProposal)
	completionProposalEqual     func(c *CompletionProposal, other *CompletionProposal) bool
	completionProposalGetIcon   func(c *CompletionProposal) *D.Pixbuf
	completionProposalGetInfo   func(c *CompletionProposal) string
	completionProposalGetLabel  func(c *CompletionProposal) string
	completionProposalGetMarkup func(c *CompletionProposal) string
	completionProposalGetText   func(c *CompletionProposal) string
	completionProposalHash      func(c *CompletionProposal) uint
)

func (c *CompletionProposal) Changed() { completionProposalChanged(c) }
func (c *CompletionProposal) Equal(other *CompletionProposal) bool {
	return completionProposalEqual(c, other)
}
func (c *CompletionProposal) Hash() uint      { return completionProposalHash(c) }
func (c *CompletionProposal) Icon() *D.Pixbuf { return completionProposalGetIcon(c) }
func (c *CompletionProposal) Info() string    { return completionProposalGetInfo(c) }
func (c *CompletionProposal) Label() string   { return completionProposalGetLabel(c) }
func (c *CompletionProposal) Markup() string  { return completionProposalGetMarkup(c) }
func (c *CompletionProposal) Text() string    { return completionProposalGetText(c) }

type CompletionProvider struct{}

var (
	CompletionProviderGetType func() O.Type

	completionProviderActivateProposal    func(c *CompletionProvider, proposal *CompletionProposal, iter *G.TextIter) bool
	completionProviderGetActivation       func(c *CompletionProvider) CompletionActivation
	completionProviderGetIcon             func(c *CompletionProvider) *D.Pixbuf
	completionProviderGetInfoWidget       func(c *CompletionProvider, proposal *CompletionProposal) *G.Widget
	completionProviderGetInteractiveDelay func(c *CompletionProvider) int
	completionProviderGetName             func(c *CompletionProvider) string
	completionProviderGetPriority         func(c *CompletionProvider) int
	completionProviderGetStartIter        func(c *CompletionProvider, context *CompletionContext, proposal *CompletionProposal, iter *G.TextIter) bool
	completionProviderMatch               func(c *CompletionProvider, context *CompletionContext) bool
	completionProviderPopulate            func(c *CompletionProvider, context *CompletionContext)
	completionProviderUpdateInfo          func(c *CompletionProvider, proposal *CompletionProposal, info *CompletionInfo)
)

func (c *CompletionProvider) Activation() CompletionActivation {
	return completionProviderGetActivation(c)
}
func (c *CompletionProvider) ActivateProposal(proposal *CompletionProposal, iter *G.TextIter) bool {
	return completionProviderActivateProposal(c, proposal, iter)
}
func (c *CompletionProvider) Icon() *D.Pixbuf { return completionProviderGetIcon(c) }
func (c *CompletionProvider) InfoWidget(proposal *CompletionProposal) *G.Widget {
	return completionProviderGetInfoWidget(c, proposal)
}
func (c *CompletionProvider) InteractiveDelay() int {
	return completionProviderGetInteractiveDelay(c)
}
func (c *CompletionProvider) Match(context *CompletionContext) bool {
	return completionProviderMatch(c, context)
}
func (c *CompletionProvider) Name() string { return completionProviderGetName(c) }
func (c *CompletionProvider) Populate(context *CompletionContext) {
	completionProviderPopulate(c, context)
}
func (c *CompletionProvider) Priority() int { return completionProviderGetPriority(c) }
func (c *CompletionProvider) StartIter(context *CompletionContext, proposal *CompletionProposal, iter *G.TextIter) bool {
	return completionProviderGetStartIter(c, context, proposal, iter)
}
func (c *CompletionProvider) UpdateInfo(proposal *CompletionProposal, info *CompletionInfo) {
	completionProviderUpdateInfo(c, proposal, info)
}

type CompletionWords struct {
	Parent O.Object
	_      *struct{}
}

var (
	CompletionWordsGetType func() O.Type
	CompletionWordsNew     func(name string, icon *D.Pixbuf) *CompletionWords

	completionWordsRegister   func(words *CompletionWords, buffer *G.TextBuffer)
	completionWordsUnregister func(words *CompletionWords, buffer *G.TextBuffer)
)

func (c *CompletionWords) Register(buffer *G.TextBuffer)   { completionWordsRegister(c, buffer) }
func (c *CompletionWords) Unregister(buffer *G.TextBuffer) { completionWordsUnregister(c, buffer) }

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

	gutterGetWindow       func(g *Gutter) *D.Window
	gutterInsert          func(g *Gutter, renderer *G.CellRenderer, position int)
	gutterQueueDraw       func(g *Gutter)
	gutterRemove          func(g *Gutter, renderer *G.CellRenderer)
	gutterReorder         func(g *Gutter, renderer *G.CellRenderer, position int)
	gutterSetCellDataFunc func(g *Gutter, renderer *G.CellRenderer, f GutterDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
	gutterSetCellSizeFunc func(g *Gutter, renderer *G.CellRenderer, f GutterSizeFunc, funcData T.Gpointer, destroy T.GDestroyNotify)
)

func (g *Gutter) Insert(renderer *G.CellRenderer, position int) { gutterInsert(g, renderer, position) }
func (g *Gutter) QueueDraw()                                    { gutterQueueDraw(g) }
func (g *Gutter) Remove(renderer *G.CellRenderer)               { gutterRemove(g, renderer) }
func (g *Gutter) Reorder(renderer *G.CellRenderer, position int) {
	gutterReorder(g, renderer, position)
}
func (g *Gutter) SetCellDataFunc(renderer *G.CellRenderer, f GutterDataFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	gutterSetCellDataFunc(g, renderer, f, funcData, destroy)
}
func (g *Gutter) SetCellSizeFunc(renderer *G.CellRenderer, f GutterSizeFunc, funcData T.Gpointer, destroy T.GDestroyNotify) {
	gutterSetCellSizeFunc(g, renderer, f, funcData, destroy)
}
func (g *Gutter) Window() *D.Window { return gutterGetWindow(g) }

type GutterDataFunc func(gutter *Gutter, cell *G.CellRenderer,
	lineNumber int, currentLine bool, data T.Gpointer)

type GutterSizeFunc func(
	gutter *Gutter, cell *G.CellRenderer, data T.Gpointer)

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

	languageGetGlobs     func(language *Language) []string
	languageGetHidden    func(language *Language) bool
	languageGetId        func(language *Language) string
	languageGetMetadata  func(language *Language, name string) string
	languageGetMimeTypes func(language *Language) []string
	languageGetName      func(language *Language) string
	languageGetSection   func(language *Language) string
	languageGetStyleIds  func(language *Language) []string
	languageGetStyleName func(language *Language, styleId string) string
)

func (l *Language) Globs() []string                 { return languageGetGlobs(l) }
func (l *Language) Hidden() bool                    { return languageGetHidden(l) }
func (l *Language) Id() string                      { return languageGetId(l) }
func (l *Language) Metadata(name string) string     { return languageGetMetadata(l, name) }
func (l *Language) MimeTypes() []string             { return languageGetMimeTypes(l) }
func (l *Language) Name() string                    { return languageGetName(l) }
func (l *Language) Section() string                 { return languageGetSection(l) }
func (l *Language) StyleIds() []string              { return languageGetStyleIds(l) }
func (l *Language) StyleName(styleId string) string { return languageGetStyleName(l, styleId) }

type LanguageManager struct {
	Parent O.Object
	_      *struct{}
}

var (
	LanguageManagerGetType func() O.Type
	LanguageManagerNew     func() *LanguageManager

	LanguageManagerGetDefault func() *LanguageManager

	languageManagerGetLanguage    func(l *LanguageManager, id string) *Language
	languageManagerGetLanguageIds func(l *LanguageManager) []string
	languageManagerGetSearchPath  func(l *LanguageManager) []string
	languageManagerGuessLanguage  func(l *LanguageManager, filename string, contentType string) *Language
	languageManagerSetSearchPath  func(l *LanguageManager, dirs []string)
)

func (l *LanguageManager) GuessLanguage(filename string, contentType string) *Language {
	return languageManagerGuessLanguage(l, filename, contentType)
}
func (l *LanguageManager) Language(id string) *Language { return languageManagerGetLanguage(l, id) }
func (l *LanguageManager) LanguageIds() []string        { return languageManagerGetLanguageIds(l) }
func (l *LanguageManager) SearchPath() []string         { return languageManagerGetSearchPath(l) }
func (l *LanguageManager) SetSearchPath(dirs []string)  { languageManagerSetSearchPath(l, dirs) }

type Mark struct {
	Parent G.TextMark
	_      *struct{}
}

var (
	MarkGetType func() O.Type
	MarkNew     func(name, category string) *Mark

	markGetCategory func(m *Mark) string
	markNext        func(m *Mark, category string) *Mark
	markPrev        func(m *Mark, category string) *Mark
)

func (m *Mark) Category() string           { return markGetCategory(m) }
func (m *Mark) Next(category string) *Mark { return markNext(m, category) }
func (m *Mark) Prev(category string) *Mark { return markPrev(m, category) }

type PrintCompositor struct {
	Parent O.Object
	_      *struct{}
}

var (
	PrintCompositorGetType     func() O.Type
	PrintCompositorNew         func(buffer *Buffer) *PrintCompositor
	PrintCompositorNewFromView func(view *View) *PrintCompositor

	printCompositorDrawPage               func(p *PrintCompositor, context *G.PrintContext, pageNr int)
	printCompositorGetBodyFontName        func(p *PrintCompositor) string
	printCompositorGetBottomMargin        func(p *PrintCompositor, unit G.Unit) float64
	printCompositorGetBuffer              func(p *PrintCompositor) *Buffer
	printCompositorGetFooterFontName      func(p *PrintCompositor) string
	printCompositorGetHeaderFontName      func(p *PrintCompositor) string
	printCompositorGetHighlightSyntax     func(p *PrintCompositor) bool
	printCompositorGetLeftMargin          func(p *PrintCompositor, unit G.Unit) float64
	printCompositorGetLineNumbersFontName func(p *PrintCompositor) string
	printCompositorGetNPages              func(p *PrintCompositor) int
	printCompositorGetPaginationProgress  func(p *PrintCompositor) float64
	printCompositorGetPrintFooter         func(p *PrintCompositor) bool
	printCompositorGetPrintHeader         func(p *PrintCompositor) bool
	printCompositorGetPrintLineNumbers    func(p *PrintCompositor) uint
	printCompositorGetRightMargin         func(p *PrintCompositor, unit G.Unit) float64
	printCompositorGetTabWidth            func(p *PrintCompositor) uint
	printCompositorGetTopMargin           func(p *PrintCompositor, unit G.Unit) float64
	printCompositorGetWrapMode            func(p *PrintCompositor) G.WrapMode
	printCompositorPaginate               func(p *PrintCompositor, context *G.PrintContext) bool
	printCompositorSetBodyFontName        func(p *PrintCompositor, fontName string)
	printCompositorSetBottomMargin        func(p *PrintCompositor, margin float64, unit G.Unit)
	printCompositorSetFooterFontName      func(p *PrintCompositor, fontName string)
	printCompositorSetFooterFormat        func(p *PrintCompositor, separator bool, left, center, right string)
	printCompositorSetHeaderFontName      func(p *PrintCompositor, fontName string)
	printCompositorSetHeaderFormat        func(p *PrintCompositor, separator bool, left, center, right string)
	printCompositorSetHighlightSyntax     func(p *PrintCompositor, highlight bool)
	printCompositorSetLeftMargin          func(p *PrintCompositor, margin float64, unit G.Unit)
	printCompositorSetLineNumbersFontName func(p *PrintCompositor, fontName string)
	printCompositorSetPrintFooter         func(p *PrintCompositor, prt bool)
	printCompositorSetPrintHeader         func(p *PrintCompositor, prt bool)
	printCompositorSetPrintLineNumbers    func(p *PrintCompositor, interval uint)
	printCompositorSetRightMargin         func(p *PrintCompositor, margin float64, unit G.Unit)
	printCompositorSetTabWidth            func(p *PrintCompositor, width uint)
	printCompositorSetTopMargin           func(p *PrintCompositor, margin float64, unit G.Unit)
	printCompositorSetWrapMode            func(p *PrintCompositor, wrapMode G.WrapMode)
)

func (p *PrintCompositor) BodyFontName() string { return printCompositorGetBodyFontName(p) }
func (p *PrintCompositor) BottomMargin(unit G.Unit) float64 {
	return printCompositorGetBottomMargin(p, unit)
}
func (p *PrintCompositor) Buffer() *Buffer { return printCompositorGetBuffer(p) }
func (p *PrintCompositor) DrawPage(context *G.PrintContext, pageNr int) {
	printCompositorDrawPage(p, context, pageNr)
}
func (p *PrintCompositor) FooterFontName() string { return printCompositorGetFooterFontName(p) }
func (p *PrintCompositor) HeaderFontName() string { return printCompositorGetHeaderFontName(p) }
func (p *PrintCompositor) HighlightSyntax() bool  { return printCompositorGetHighlightSyntax(p) }
func (p *PrintCompositor) LeftMargin(unit G.Unit) float64 {
	return printCompositorGetLeftMargin(p, unit)
}
func (p *PrintCompositor) LineNumbersFontName() string {
	return printCompositorGetLineNumbersFontName(p)
}
func (p *PrintCompositor) NPages() int { return printCompositorGetNPages(p) }
func (p *PrintCompositor) Paginate(context *G.PrintContext) bool {
	return printCompositorPaginate(p, context)
}
func (p *PrintCompositor) PaginationProgress() float64 {
	return printCompositorGetPaginationProgress(p)
}
func (p *PrintCompositor) PrintFooter() bool      { return printCompositorGetPrintFooter(p) }
func (p *PrintCompositor) PrintHeader() bool      { return printCompositorGetPrintHeader(p) }
func (p *PrintCompositor) PrintLineNumbers() uint { return printCompositorGetPrintLineNumbers(p) }
func (p *PrintCompositor) RightMargin(unit G.Unit) float64 {
	return printCompositorGetRightMargin(p, unit)
}
func (p *PrintCompositor) SetBodyFontName(fontName string) {
	printCompositorSetBodyFontName(p, fontName)
}
func (p *PrintCompositor) SetBottomMargin(margin float64, unit G.Unit) {
	printCompositorSetBottomMargin(p, margin, unit)
}
func (p *PrintCompositor) SetFooterFontName(fontName string) {
	printCompositorSetFooterFontName(p, fontName)
}
func (p *PrintCompositor) SetFooterFormat(separator bool, left, center, right string) {
	printCompositorSetFooterFormat(p, separator, left, center, right)
}
func (p *PrintCompositor) SetHeaderFontName(fontName string) {
	printCompositorSetHeaderFontName(p, fontName)
}
func (p *PrintCompositor) SetHeaderFormat(separator bool, left, center, right string) {
	printCompositorSetHeaderFormat(p, separator, left, center, right)
}
func (p *PrintCompositor) SetHighlightSyntax(highlight bool) {
	printCompositorSetHighlightSyntax(p, highlight)
}
func (p *PrintCompositor) SetLeftMargin(margin float64, unit G.Unit) {
	printCompositorSetLeftMargin(p, margin, unit)
}
func (p *PrintCompositor) SetLineNumbersFontName(fontName string) {
	printCompositorSetLineNumbersFontName(p, fontName)
}
func (p *PrintCompositor) SetPrintFooter(prt bool) { printCompositorSetPrintFooter(p, prt) }
func (p *PrintCompositor) SetPrintHeader(prt bool) { printCompositorSetPrintHeader(p, prt) }
func (p *PrintCompositor) SetPrintLineNumbers(interval uint) {
	printCompositorSetPrintLineNumbers(p, interval)
}
func (p *PrintCompositor) SetRightMargin(margin float64, unit G.Unit) {
	printCompositorSetRightMargin(p, margin, unit)
}
func (p *PrintCompositor) SetTabWidth(width uint) { printCompositorSetTabWidth(p, width) }
func (p *PrintCompositor) SetTopMargin(margin float64, unit G.Unit) {
	printCompositorSetTopMargin(p, margin, unit)
}
func (p *PrintCompositor) SetWrapMode(wrapMode G.WrapMode) { printCompositorSetWrapMode(p, wrapMode) }
func (p *PrintCompositor) TabWidth() uint                  { return printCompositorGetTabWidth(p) }
func (p *PrintCompositor) TopMargin(unit G.Unit) float64 {
	return printCompositorGetTopMargin(p, unit)
}
func (p *PrintCompositor) WrapMode() G.WrapMode { return printCompositorGetWrapMode(p) }

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

	styleSchemeManagerAppendSearchPath  func(s *StyleSchemeManager, path string)
	styleSchemeManagerForceRescan       func(s *StyleSchemeManager)
	styleSchemeManagerGetScheme         func(s *StyleSchemeManager, schemeId string) *StyleScheme
	styleSchemeManagerGetSchemeIds      func(s *StyleSchemeManager) []string
	styleSchemeManagerGetSearchPath     func(s *StyleSchemeManager) []string
	styleSchemeManagerPrependSearchPath func(s *StyleSchemeManager, path string)
	styleSchemeManagerSetSearchPath     func(s *StyleSchemeManager, path **T.Gchar)
)

func (s *StyleSchemeManager) AppendSearchPath(path string) {
	styleSchemeManagerAppendSearchPath(s, path)
}
func (s *StyleSchemeManager) ForceRescan() { styleSchemeManagerForceRescan(s) }
func (s *StyleSchemeManager) PrependSearchPath(path string) {
	styleSchemeManagerPrependSearchPath(s, path)
}
func (s *StyleSchemeManager) Scheme(schemeId string) *StyleScheme {
	return styleSchemeManagerGetScheme(s, schemeId)
}
func (s *StyleSchemeManager) SchemeIds() []string          { return styleSchemeManagerGetSchemeIds(s) }
func (s *StyleSchemeManager) SearchPath() []string         { return styleSchemeManagerGetSearchPath(s) }
func (s *StyleSchemeManager) SetSearchPath(path **T.Gchar) { styleSchemeManagerSetSearchPath(s, path) }

type UndoManager struct{}

var (
	UndoManagerGetType func() O.Type

	undoManagerBeginNotUndoableAction func(u *UndoManager)
	undoManagerCanRedo                func(u *UndoManager) bool
	undoManagerCanRedoChanged         func(u *UndoManager)
	undoManagerCanUndo                func(u *UndoManager) bool
	undoManagerCanUndoChanged         func(u *UndoManager)
	undoManagerEndNotUndoableAction   func(u *UndoManager)
	undoManagerRedo                   func(u *UndoManager)
	undoManagerUndo                   func(u *UndoManager)
)

func (u *UndoManager) BeginNotUndoableAction() { undoManagerBeginNotUndoableAction(u) }
func (u *UndoManager) CanRedo() bool           { return undoManagerCanRedo(u) }
func (u *UndoManager) CanRedoChanged()         { undoManagerCanRedoChanged(u) }
func (u *UndoManager) CanUndo() bool           { return undoManagerCanUndo(u) }
func (u *UndoManager) CanUndoChanged()         { undoManagerCanUndoChanged(u) }
func (u *UndoManager) EndNotUndoableAction()   { undoManagerEndNotUndoableAction(u) }
func (u *UndoManager) Redo()                   { undoManagerRedo(u) }
func (u *UndoManager) Undo()                   { undoManagerUndo(u) }

type View struct {
	Parent G.TextView
	_      *struct{}
}

var (
	ViewGetType       func() O.Type
	ViewNew           func() *G.Widget
	ViewNewWithBuffer func(buffer *Buffer) *G.Widget

	viewGetAutoIndent                    func(v *View) bool
	viewGetCompletion                    func(v *View) *Completion
	viewGetDrawSpaces                    func(v *View) DrawSpacesFlags
	viewGetGutter                        func(v *View, windowType G.TextWindowType) *Gutter
	viewGetHighlightCurrentLine          func(v *View) bool
	viewGetIndentOnTab                   func(v *View) bool
	viewGetIndentWidth                   func(v *View) int
	viewGetInsertSpacesInsteadOfTabs     func(v *View) bool
	viewGetMarkCategoryBackground        func(v *View, category string, dest *D.Color) bool
	viewGetMarkCategoryPixbuf            func(v *View, category string) *D.Pixbuf
	viewGetMarkCategoryPriority          func(v *View, category string) int
	viewGetRightMarginPosition           func(v *View) uint
	viewGetShowLineMarks                 func(v *View) bool
	viewGetShowLineNumbers               func(v *View) bool
	viewGetShowRightMargin               func(v *View) bool
	viewGetSmartHomeEnd                  func(v *View) SmartHomeEndType
	viewGetTabWidth                      func(v *View) uint
	viewSetAutoIndent                    func(v *View, enable bool)
	viewSetDrawSpaces                    func(v *View, flags DrawSpacesFlags)
	viewSetHighlightCurrentLine          func(v *View, show bool)
	viewSetIndentOnTab                   func(v *View, enable bool)
	viewSetIndentWidth                   func(v *View, width int)
	viewSetInsertSpacesInsteadOfTabs     func(v *View, enable bool)
	viewSetMarkCategoryBackground        func(v *View, category string, color *D.Color)
	viewSetMarkCategoryIconFromIconName  func(v *View, category, name string)
	viewSetMarkCategoryIconFromPixbuf    func(v *View, category string, pixbuf *D.Pixbuf)
	viewSetMarkCategoryIconFromStock     func(v *View, category, stockId string)
	viewSetMarkCategoryPixbuf            func(v *View, category string, pixbuf *D.Pixbuf)
	viewSetMarkCategoryPriority          func(v *View, category string, priority int)
	viewSetMarkCategoryTooltipFunc       func(v *View, category string, f ViewMarkTooltipFunc, userData T.Gpointer, userDataNotify T.GDestroyNotify)
	viewSetMarkCategoryTooltipMarkupFunc func(v *View, category string, markupFunc ViewMarkTooltipFunc, userData T.Gpointer, userDataNotify T.GDestroyNotify)
	viewSetRightMarginPosition           func(v *View, pos uint)
	viewSetShowLineMarks                 func(v *View, show bool)
	viewSetShowLineNumbers               func(v *View, show bool)
	viewSetShowRightMargin               func(v *View, show bool)
	viewSetSmartHomeEnd                  func(v *View, smartHe SmartHomeEndType)
	viewSetTabWidth                      func(v *View, width uint)
)

func (v *View) AutoIndent() bool                           { return viewGetAutoIndent(v) }
func (v *View) Completion() *Completion                    { return viewGetCompletion(v) }
func (v *View) DrawSpaces() DrawSpacesFlags                { return viewGetDrawSpaces(v) }
func (v *View) Gutter(windowType G.TextWindowType) *Gutter { return viewGetGutter(v, windowType) }
func (v *View) HighlightCurrentLine() bool                 { return viewGetHighlightCurrentLine(v) }
func (v *View) IndentOnTab() bool                          { return viewGetIndentOnTab(v) }
func (v *View) IndentWidth() int                           { return viewGetIndentWidth(v) }
func (v *View) InsertSpacesInsteadOfTabs() bool            { return viewGetInsertSpacesInsteadOfTabs(v) }
func (v *View) MarkCategoryBackground(category string, dest *D.Color) bool {
	return viewGetMarkCategoryBackground(v, category, dest)
}
func (v *View) MarkCategoryPixbuf(category string) *D.Pixbuf {
	return viewGetMarkCategoryPixbuf(v, category)
}
func (v *View) MarkCategoryPriority(category string) int {
	return viewGetMarkCategoryPriority(v, category)
}
func (v *View) RightMarginPosition() uint           { return viewGetRightMarginPosition(v) }
func (v *View) SetAutoIndent(enable bool)           { viewSetAutoIndent(v, enable) }
func (v *View) SetDrawSpaces(flags DrawSpacesFlags) { viewSetDrawSpaces(v, flags) }
func (v *View) SetHighlightCurrentLine(show bool)   { viewSetHighlightCurrentLine(v, show) }
func (v *View) SetIndentOnTab(enable bool)          { viewSetIndentOnTab(v, enable) }
func (v *View) SetIndentWidth(width int)            { viewSetIndentWidth(v, width) }
func (v *View) SetInsertSpacesInsteadOfTabs(enable bool) {
	viewSetInsertSpacesInsteadOfTabs(v, enable)
}
func (v *View) SetMarkCategoryBackground(category string, color *D.Color) {
	viewSetMarkCategoryBackground(v, category, color)
}
func (v *View) SetMarkCategoryIconFromIconName(category, name string) {
	viewSetMarkCategoryIconFromIconName(v, category, name)
}
func (v *View) SetMarkCategoryIconFromPixbuf(category string, pixbuf *D.Pixbuf) {
	viewSetMarkCategoryIconFromPixbuf(v, category, pixbuf)
}
func (v *View) SetMarkCategoryIconFromStock(category, stockId string) {
	viewSetMarkCategoryIconFromStock(v, category, stockId)
}
func (v *View) SetMarkCategoryPixbuf(category string, pixbuf *D.Pixbuf) {
	viewSetMarkCategoryPixbuf(v, category, pixbuf)
}
func (v *View) SetMarkCategoryPriority(category string, priority int) {
	viewSetMarkCategoryPriority(v, category, priority)
}
func (v *View) SetMarkCategoryTooltipFunc(category string, f ViewMarkTooltipFunc, userData T.Gpointer, userDataNotify T.GDestroyNotify) {
	viewSetMarkCategoryTooltipFunc(v, category, f, userData, userDataNotify)
}
func (v *View) SetMarkCategoryTooltipMarkupFunc(category string, markupFunc ViewMarkTooltipFunc, userData T.Gpointer, userDataNotify T.GDestroyNotify) {
	viewSetMarkCategoryTooltipMarkupFunc(v, category, markupFunc, userData, userDataNotify)
}
func (v *View) SetRightMarginPosition(pos uint)          { viewSetRightMarginPosition(v, pos) }
func (v *View) SetShowLineMarks(show bool)               { viewSetShowLineMarks(v, show) }
func (v *View) SetShowLineNumbers(show bool)             { viewSetShowLineNumbers(v, show) }
func (v *View) SetShowRightMargin(show bool)             { viewSetShowRightMargin(v, show) }
func (v *View) SetSmartHomeEnd(smartHe SmartHomeEndType) { viewSetSmartHomeEnd(v, smartHe) }
func (v *View) SetTabWidth(width uint)                   { viewSetTabWidth(v, width) }
func (v *View) ShowLineMarks() bool                      { return viewGetShowLineMarks(v) }
func (v *View) ShowLineNumbers() bool                    { return viewGetShowLineNumbers(v) }
func (v *View) ShowRightMargin() bool                    { return viewGetShowRightMargin(v) }
func (v *View) SmartHomeEnd() SmartHomeEndType           { return viewGetSmartHomeEnd(v) }
func (v *View) TabWidth() uint                           { return viewGetTabWidth(v) }

var ViewGutterPositionGetType func() O.Type

type ViewMarkTooltipFunc func(
	mark *Mark, userData T.Gpointer) string

var dll = "libgtksourceview-2.0-0.dll"

var apiList = outside.Apis{
	{"gtk_source_buffer_backward_iter_to_source_mark", &bufferBackwardIterToSourceMark},
	{"gtk_source_buffer_begin_not_undoable_action", &bufferBeginNotUndoableAction},
	{"gtk_source_buffer_can_redo", &bufferCanRedo},
	{"gtk_source_buffer_can_undo", &bufferCanUndo},
	{"gtk_source_buffer_create_source_mark", &bufferCreateSourceMark},
	{"gtk_source_buffer_end_not_undoable_action", &bufferEndNotUndoableAction},
	{"gtk_source_buffer_ensure_highlight", &bufferEnsureHighlight},
	{"gtk_source_buffer_forward_iter_to_source_mark", &bufferForwardIterToSourceMark},
	{"gtk_source_buffer_get_context_classes_at_iter", &bufferGetContextClassesAtIter},
	{"gtk_source_buffer_get_highlight_matching_brackets", &bufferGetHighlightMatchingBrackets},
	{"gtk_source_buffer_get_highlight_syntax", &bufferGetHighlightSyntax},
	{"gtk_source_buffer_get_language", &bufferGetLanguage},
	{"gtk_source_buffer_get_max_undo_levels", &bufferGetMaxUndoLevels},
	{"gtk_source_buffer_get_source_marks_at_iter", &bufferGetSourceMarksAtIter},
	{"gtk_source_buffer_get_source_marks_at_line", &bufferGetSourceMarksAtLine},
	{"gtk_source_buffer_get_style_scheme", &bufferGetStyleScheme},
	{"gtk_source_buffer_get_type", &BufferGetType},
	{"gtk_source_buffer_get_undo_manager", &bufferGetUndoManager},
	{"gtk_source_buffer_iter_backward_to_context_class_toggle", &bufferIterBackwardToContextClassToggle},
	{"gtk_source_buffer_iter_forward_to_context_class_toggle", &bufferIterForwardToContextClassToggle},
	{"gtk_source_buffer_iter_has_context_class", &bufferIterHasContextClass},
	{"gtk_source_buffer_new", &BufferNew},
	{"gtk_source_buffer_new_with_language", &BufferNewWithLanguage},
	{"gtk_source_buffer_redo", &bufferRedo},
	{"gtk_source_buffer_remove_source_marks", &bufferRemoveSourceMarks},
	{"gtk_source_buffer_set_highlight_matching_brackets", &bufferSetHighlightMatchingBrackets},
	{"gtk_source_buffer_set_highlight_syntax", &bufferSetHighlightSyntax},
	{"gtk_source_buffer_set_language", &bufferSetLanguage},
	{"gtk_source_buffer_set_max_undo_levels", &bufferSetMaxUndoLevels},
	{"gtk_source_buffer_set_style_scheme", &bufferSetStyleScheme},
	{"gtk_source_buffer_set_undo_manager", &bufferSetUndoManager},
	{"gtk_source_buffer_undo", &bufferUndo},
	{"gtk_source_completion_activation_get_type", &CompletionActivationGetType},
	{"gtk_source_completion_add_provider", &completionAddProvider},
	{"gtk_source_completion_block_interactive", &completionBlockInteractive},
	{"gtk_source_completion_context_add_proposals", &completionContextAddProposals},
	{"gtk_source_completion_context_get_activation", &completionContextGetActivation},
	{"gtk_source_completion_context_get_iter", &completionContextGetIter},
	{"gtk_source_completion_context_get_type", &CompletionContextGetType},
	{"gtk_source_completion_create_context", &completionCreateContext},
	{"gtk_source_completion_error_get_type", &CompletionErrorGetType},
	{"gtk_source_completion_error_quark", &CompletionErrorQuark},
	{"gtk_source_completion_get_info_window", &completionGetInfoWindow},
	{"gtk_source_completion_get_providers", &completionGetProviders},
	{"gtk_source_completion_get_type", &CompletionGetType},
	{"gtk_source_completion_get_view", &completionGetView},
	{"gtk_source_completion_hide", &completionHide},
	{"gtk_source_completion_info_get_type", &CompletionInfoGetType},
	{"gtk_source_completion_info_get_widget", &completionInfoGetWidget},
	{"gtk_source_completion_info_move_to_iter", &completionInfoMoveToIter},
	{"gtk_source_completion_info_new", &CompletionInfoNew},
	{"gtk_source_completion_info_process_resize", &completionInfoProcessResize},
	{"gtk_source_completion_info_set_sizing", &completionInfoSetSizing},
	{"gtk_source_completion_info_set_widget", &completionInfoSetWidget},
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
	{"gtk_source_completion_move_window", &completionMoveWindow},
	// Undocumented {"gtk_source_completion_new", &CompletionNew},
	{"gtk_source_completion_proposal_changed", &completionProposalChanged},
	{"gtk_source_completion_proposal_equal", &completionProposalEqual},
	{"gtk_source_completion_proposal_get_icon", &completionProposalGetIcon},
	{"gtk_source_completion_proposal_get_info", &completionProposalGetInfo},
	{"gtk_source_completion_proposal_get_label", &completionProposalGetLabel},
	{"gtk_source_completion_proposal_get_markup", &completionProposalGetMarkup},
	{"gtk_source_completion_proposal_get_text", &completionProposalGetText},
	{"gtk_source_completion_proposal_get_type", &CompletionProposalGetType},
	{"gtk_source_completion_proposal_hash", &completionProposalHash},
	{"gtk_source_completion_provider_activate_proposal", &completionProviderActivateProposal},
	{"gtk_source_completion_provider_get_activation", &completionProviderGetActivation},
	{"gtk_source_completion_provider_get_icon", &completionProviderGetIcon},
	{"gtk_source_completion_provider_get_info_widget", &completionProviderGetInfoWidget},
	{"gtk_source_completion_provider_get_interactive_delay", &completionProviderGetInteractiveDelay},
	{"gtk_source_completion_provider_get_name", &completionProviderGetName},
	{"gtk_source_completion_provider_get_priority", &completionProviderGetPriority},
	{"gtk_source_completion_provider_get_start_iter", &completionProviderGetStartIter},
	{"gtk_source_completion_provider_get_type", &CompletionProviderGetType},
	{"gtk_source_completion_provider_match", &completionProviderMatch},
	{"gtk_source_completion_provider_populate", &completionProviderPopulate},
	{"gtk_source_completion_provider_update_info", &completionProviderUpdateInfo},
	{"gtk_source_completion_remove_provider", &completionRemoveProvider},
	{"gtk_source_completion_show", &completionShow},
	{"gtk_source_completion_unblock_interactive", &completionUnblockInteractive},
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
	{"gtk_source_completion_words_register", &completionWordsRegister},
	{"gtk_source_completion_words_unregister", &completionWordsUnregister},
	// Undocumented {"gtk_source_completion_words_utils_backward_word_start", &CompletionWordsUtilsBackwardWordStart},
	// Undocumented {"gtk_source_completion_words_utils_forward_word_end", &CompletionWordsUtilsForwardWordEnd},
	// Undocumented {"gtk_source_context_class_free", &ContextClassFree},
	// Undocumented {"gtk_source_context_class_new", &ContextClassNew},
	{"gtk_source_draw_spaces_flags_get_type", &DrawSpacesFlagsGetType},
	{"gtk_source_gutter_get_type", &GutterGetType},
	{"gtk_source_gutter_get_window", &gutterGetWindow},
	{"gtk_source_gutter_insert", &gutterInsert},
	// Undocumented {"gtk_source_gutter_new", &GutterNew},
	{"gtk_source_gutter_queue_draw", &gutterQueueDraw},
	{"gtk_source_gutter_remove", &gutterRemove},
	{"gtk_source_gutter_reorder", &gutterReorder},
	{"gtk_source_gutter_set_cell_data_func", &gutterSetCellDataFunc},
	{"gtk_source_gutter_set_cell_size_func", &gutterSetCellSizeFunc},
	{"gtk_source_iter_backward_search", &IterBackwardSearch},
	{"gtk_source_iter_forward_search", &IterForwardSearch},
	{"gtk_source_language_get_globs", &languageGetGlobs},
	{"gtk_source_language_get_hidden", &languageGetHidden},
	{"gtk_source_language_get_id", &languageGetId},
	{"gtk_source_language_get_metadata", &languageGetMetadata},
	{"gtk_source_language_get_mime_types", &languageGetMimeTypes},
	{"gtk_source_language_get_name", &languageGetName},
	{"gtk_source_language_get_section", &languageGetSection},
	{"gtk_source_language_get_style_ids", &languageGetStyleIds},
	{"gtk_source_language_get_style_name", &languageGetStyleName},
	{"gtk_source_language_get_type", &LanguageGetType},
	{"gtk_source_language_manager_get_default", &LanguageManagerGetDefault},
	{"gtk_source_language_manager_get_language", &languageManagerGetLanguage},
	{"gtk_source_language_manager_get_language_ids", &languageManagerGetLanguageIds},
	{"gtk_source_language_manager_get_search_path", &languageManagerGetSearchPath},
	{"gtk_source_language_manager_get_type", &LanguageManagerGetType},
	{"gtk_source_language_manager_guess_language", &languageManagerGuessLanguage},
	{"gtk_source_language_manager_new", &LanguageManagerNew},
	{"gtk_source_language_manager_set_search_path", &languageManagerSetSearchPath},
	{"gtk_source_mark_get_category", &markGetCategory},
	{"gtk_source_mark_get_type", &MarkGetType},
	{"gtk_source_mark_new", &MarkNew},
	{"gtk_source_mark_next", &markNext},
	{"gtk_source_mark_prev", &markPrev},
	{"gtk_source_print_compositor_draw_page", &printCompositorDrawPage},
	{"gtk_source_print_compositor_get_body_font_name", &printCompositorGetBodyFontName},
	{"gtk_source_print_compositor_get_bottom_margin", &printCompositorGetBottomMargin},
	{"gtk_source_print_compositor_get_buffer", &printCompositorGetBuffer},
	{"gtk_source_print_compositor_get_footer_font_name", &printCompositorGetFooterFontName},
	{"gtk_source_print_compositor_get_header_font_name", &printCompositorGetHeaderFontName},
	{"gtk_source_print_compositor_get_highlight_syntax", &printCompositorGetHighlightSyntax},
	{"gtk_source_print_compositor_get_left_margin", &printCompositorGetLeftMargin},
	{"gtk_source_print_compositor_get_line_numbers_font_name", &printCompositorGetLineNumbersFontName},
	{"gtk_source_print_compositor_get_n_pages", &printCompositorGetNPages},
	{"gtk_source_print_compositor_get_pagination_progress", &printCompositorGetPaginationProgress},
	{"gtk_source_print_compositor_get_print_footer", &printCompositorGetPrintFooter},
	{"gtk_source_print_compositor_get_print_header", &printCompositorGetPrintHeader},
	{"gtk_source_print_compositor_get_print_line_numbers", &printCompositorGetPrintLineNumbers},
	{"gtk_source_print_compositor_get_right_margin", &printCompositorGetRightMargin},
	{"gtk_source_print_compositor_get_tab_width", &printCompositorGetTabWidth},
	{"gtk_source_print_compositor_get_top_margin", &printCompositorGetTopMargin},
	{"gtk_source_print_compositor_get_type", &PrintCompositorGetType},
	{"gtk_source_print_compositor_get_wrap_mode", &printCompositorGetWrapMode},
	{"gtk_source_print_compositor_new", &PrintCompositorNew},
	{"gtk_source_print_compositor_new_from_view", &PrintCompositorNewFromView},
	{"gtk_source_print_compositor_paginate", &printCompositorPaginate},
	{"gtk_source_print_compositor_set_body_font_name", &printCompositorSetBodyFontName},
	{"gtk_source_print_compositor_set_bottom_margin", &printCompositorSetBottomMargin},
	{"gtk_source_print_compositor_set_footer_font_name", &printCompositorSetFooterFontName},
	{"gtk_source_print_compositor_set_footer_format", &printCompositorSetFooterFormat},
	{"gtk_source_print_compositor_set_header_font_name", &printCompositorSetHeaderFontName},
	{"gtk_source_print_compositor_set_header_format", &printCompositorSetHeaderFormat},
	{"gtk_source_print_compositor_set_highlight_syntax", &printCompositorSetHighlightSyntax},
	{"gtk_source_print_compositor_set_left_margin", &printCompositorSetLeftMargin},
	{"gtk_source_print_compositor_set_line_numbers_font_name", &printCompositorSetLineNumbersFontName},
	{"gtk_source_print_compositor_set_print_footer", &printCompositorSetPrintFooter},
	{"gtk_source_print_compositor_set_print_header", &printCompositorSetPrintHeader},
	{"gtk_source_print_compositor_set_print_line_numbers", &printCompositorSetPrintLineNumbers},
	{"gtk_source_print_compositor_set_right_margin", &printCompositorSetRightMargin},
	{"gtk_source_print_compositor_set_tab_width", &printCompositorSetTabWidth},
	{"gtk_source_print_compositor_set_top_margin", &printCompositorSetTopMargin},
	{"gtk_source_print_compositor_set_wrap_mode", &printCompositorSetWrapMode},
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
	{"gtk_source_style_scheme_manager_append_search_path", &styleSchemeManagerAppendSearchPath},
	{"gtk_source_style_scheme_manager_force_rescan", &styleSchemeManagerForceRescan},
	{"gtk_source_style_scheme_manager_get_default", &StyleSchemeManagerGetDefault},
	{"gtk_source_style_scheme_manager_get_scheme", &styleSchemeManagerGetScheme},
	{"gtk_source_style_scheme_manager_get_scheme_ids", &styleSchemeManagerGetSchemeIds},
	{"gtk_source_style_scheme_manager_get_search_path", &styleSchemeManagerGetSearchPath},
	{"gtk_source_style_scheme_manager_get_type", &StyleSchemeManagerGetType},
	{"gtk_source_style_scheme_manager_new", &StyleSchemeManagerNew},
	{"gtk_source_style_scheme_manager_prepend_search_path", &styleSchemeManagerPrependSearchPath},
	{"gtk_source_style_scheme_manager_set_search_path", &styleSchemeManagerSetSearchPath},
	{"gtk_source_undo_manager_begin_not_undoable_action", &undoManagerBeginNotUndoableAction},
	{"gtk_source_undo_manager_can_redo", &undoManagerCanRedo},
	{"gtk_source_undo_manager_can_redo_changed", &undoManagerCanRedoChanged},
	{"gtk_source_undo_manager_can_undo", &undoManagerCanUndo},
	{"gtk_source_undo_manager_can_undo_changed", &undoManagerCanUndoChanged},
	// Undocumented {"gtk_source_undo_manager_default_get_type", &UndoManagerDefaultGetType},
	// Undocumented {"gtk_source_undo_manager_default_set_max_undo_levels", &UndoManagerDefaultSetMaxUndoLevels},
	{"gtk_source_undo_manager_end_not_undoable_action", &undoManagerEndNotUndoableAction},
	{"gtk_source_undo_manager_get_type", &UndoManagerGetType},
	{"gtk_source_undo_manager_redo", &undoManagerRedo},
	{"gtk_source_undo_manager_undo", &undoManagerUndo},
	{"gtk_source_view_get_auto_indent", &viewGetAutoIndent},
	{"gtk_source_view_get_completion", &viewGetCompletion},
	{"gtk_source_view_get_draw_spaces", &viewGetDrawSpaces},
	{"gtk_source_view_get_gutter", &viewGetGutter},
	{"gtk_source_view_get_highlight_current_line", &viewGetHighlightCurrentLine},
	{"gtk_source_view_get_indent_on_tab", &viewGetIndentOnTab},
	{"gtk_source_view_get_indent_width", &viewGetIndentWidth},
	{"gtk_source_view_get_insert_spaces_instead_of_tabs", &viewGetInsertSpacesInsteadOfTabs},
	{"gtk_source_view_get_mark_category_background", &viewGetMarkCategoryBackground},
	{"gtk_source_view_get_mark_category_pixbuf", &viewGetMarkCategoryPixbuf},
	{"gtk_source_view_get_mark_category_priority", &viewGetMarkCategoryPriority},
	{"gtk_source_view_get_right_margin_position", &viewGetRightMarginPosition},
	{"gtk_source_view_get_show_line_marks", &viewGetShowLineMarks},
	{"gtk_source_view_get_show_line_numbers", &viewGetShowLineNumbers},
	{"gtk_source_view_get_show_right_margin", &viewGetShowRightMargin},
	{"gtk_source_view_get_smart_home_end", &viewGetSmartHomeEnd},
	{"gtk_source_view_get_tab_width", &viewGetTabWidth},
	{"gtk_source_view_get_type", &ViewGetType},
	{"gtk_source_view_gutter_position_get_type", &ViewGutterPositionGetType},
	{"gtk_source_view_new", &ViewNew},
	{"gtk_source_view_new_with_buffer", &ViewNewWithBuffer},
	{"gtk_source_view_set_auto_indent", &viewSetAutoIndent},
	{"gtk_source_view_set_draw_spaces", &viewSetDrawSpaces},
	{"gtk_source_view_set_highlight_current_line", &viewSetHighlightCurrentLine},
	{"gtk_source_view_set_indent_on_tab", &viewSetIndentOnTab},
	{"gtk_source_view_set_indent_width", &viewSetIndentWidth},
	{"gtk_source_view_set_insert_spaces_instead_of_tabs", &viewSetInsertSpacesInsteadOfTabs},
	{"gtk_source_view_set_mark_category_background", &viewSetMarkCategoryBackground},
	{"gtk_source_view_set_mark_category_icon_from_icon_name", &viewSetMarkCategoryIconFromIconName},
	{"gtk_source_view_set_mark_category_icon_from_pixbuf", &viewSetMarkCategoryIconFromPixbuf},
	{"gtk_source_view_set_mark_category_icon_from_stock", &viewSetMarkCategoryIconFromStock},
	{"gtk_source_view_set_mark_category_pixbuf", &viewSetMarkCategoryPixbuf},
	{"gtk_source_view_set_mark_category_priority", &viewSetMarkCategoryPriority},
	{"gtk_source_view_set_mark_category_tooltip_func", &viewSetMarkCategoryTooltipFunc},
	{"gtk_source_view_set_mark_category_tooltip_markup_func", &viewSetMarkCategoryTooltipMarkupFunc},
	{"gtk_source_view_set_right_margin_position", &viewSetRightMarginPosition},
	{"gtk_source_view_set_show_line_marks", &viewSetShowLineMarks},
	{"gtk_source_view_set_show_line_numbers", &viewSetShowLineNumbers},
	{"gtk_source_view_set_show_right_margin", &viewSetShowRightMargin},
	{"gtk_source_view_set_smart_home_end", &viewSetSmartHomeEnd},
	{"gtk_source_view_set_tab_width", &viewSetTabWidth},
}
