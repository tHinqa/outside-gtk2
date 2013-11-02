// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type (
	Scanner struct {
		UserData       T.Gpointer
		MaxParseErrors uint
		ParseErrors    uint
		InputName      *T.Gchar
		Qdata          *T.GData
		Config         *ScannerConfig
		Token          T.GTokenType
		Value          T.GTokenValue
		Line           uint
		Position       uint
		NextToken      T.GTokenType
		NextValue      T.GTokenValue
		NextLine       uint
		NextPosition   uint
		SymbolTable    *HashTable
		InputFd        int
		Text           *T.Gchar
		TextEnd        *T.Gchar
		Buffer         *T.Gchar
		ScopeId        uint
		MsgHandler     ScannerMsgFunc
	}

	ScannerConfig struct {
		CsetSkipCharacters  *T.Gchar
		CsetIdentifierFirst *T.Gchar
		CsetIdentifierNth   *T.Gchar
		CpairCommentSingle  *T.Gchar
		Bits                uint
		// CaseSensitive : 1
		// SkipCommentMulti : 1
		// SkipCommentSingle : 1
		// ScanCommentMulti : 1
		// ScanIdentifier : 1
		// ScanIdentifier1char : 1
		// ScanIdentifier_NULL : 1
		// ScanSymbols : 1
		// ScanBinary : 1
		// ScanOctal : 1
		// ScanFloat : 1
		// ScanHex : 1
		// ScanHexDollar : 1
		// ScanStringSq : 1
		// ScanStringDq : 1
		// Numbers2Int : 1
		// Int2Float : 1
		// Identifier2String : 1
		// Char2Token : 1
		// Symbol2Token : 1
		// Scope0Fallback : 1
		// StoreInt64 : 1
		PaddingDummy uint
	}

	ScannerMsgFunc func(
		scanner *Scanner, message string, err T.Gboolean)
)

var (
	ScannerNew func(configTempl *ScannerConfig) *Scanner

	ScannerDestroy            func(s *Scanner)
	ScannerInputFile          func(s *Scanner, inputFd int)
	ScannerSyncFileOffset     func(s *Scanner)
	ScannerInputText          func(s *Scanner, text string, textLen uint)
	ScannerGetNextToken       func(s *Scanner) T.GTokenType
	ScannerPeekNextToken      func(s *Scanner) T.GTokenType
	ScannerCurToken           func(s *Scanner) T.GTokenType
	ScannerCurValue           func(s *Scanner) T.GTokenValue
	ScannerCurLine            func(s *Scanner) uint
	ScannerCurPosition        func(s *Scanner) uint
	ScannerEof                func(s *Scanner) bool
	ScannerSetScope           func(s *Scanner, scopeId uint) uint
	ScannerScopeAddSymbol     func(s *Scanner, scopeId uint, symbol string, value T.Gpointer)
	ScannerScopeRemoveSymbol  func(s *Scanner, scopeId uint, symbol string)
	ScannerScopeLookupSymbol  func(s *Scanner, scopeId uint, symbol string) T.Gpointer
	ScannerScopeForeachSymbol func(s *Scanner, scopeId uint, f HFunc, userData T.Gpointer)
	ScannerLookupSymbol       func(s *Scanner, symbol string) T.Gpointer
	ScannerUnexpToken         func(s *Scanner, expectedToken T.GTokenType, identifierSpec, symbolSpec, symbolName, message string, isError int)
	ScannerError              func(s *Scanner, format string, v ...VArg)
	ScannerWarn               func(s *Scanner, format string, v ...VArg)
)

func (s *Scanner) Destroy()                            { ScannerDestroy(s) }
func (s *Scanner) InputFile(inputFd int)               { ScannerInputFile(s, inputFd) }
func (s *Scanner) SyncFileOffset()                     { ScannerSyncFileOffset(s) }
func (s *Scanner) InputText(text string, textLen uint) { ScannerInputText(s, text, textLen) }
func (s *Scanner) GetNextToken() T.GTokenType          { return ScannerGetNextToken(s) }
func (s *Scanner) PeekNextToken() T.GTokenType         { return ScannerPeekNextToken(s) }
func (s *Scanner) CurToken() T.GTokenType              { return ScannerCurToken(s) }
func (s *Scanner) CurValue() T.GTokenValue             { return ScannerCurValue(s) }
func (s *Scanner) CurLine() uint                       { return ScannerCurLine(s) }
func (s *Scanner) CurPosition() uint                   { return ScannerCurPosition(s) }
func (s *Scanner) Eof() bool                           { return ScannerEof(s) }
func (s *Scanner) SetScope(scopeId uint) uint          { return ScannerSetScope(s, scopeId) }
func (s *Scanner) ScopeAddSymbol(scopeId uint, symbol string, value T.Gpointer) {
	ScannerScopeAddSymbol(s, scopeId, symbol, value)
}
func (s *Scanner) ScopeRemoveSymbol(scopeId uint, symbol string) {
	ScannerScopeRemoveSymbol(s, scopeId, symbol)
}
func (s *Scanner) ScopeLookupSymbol(scopeId uint, symbol string) T.Gpointer {
	return ScannerScopeLookupSymbol(s, scopeId, symbol)
}
func (s *Scanner) ScopeForeachSymbol(scopeId uint, f HFunc, userData T.Gpointer) {
	ScannerScopeForeachSymbol(s, scopeId, f, userData)
}
func (s *Scanner) LookupSymbol(symbol string) T.Gpointer { return ScannerLookupSymbol(s, symbol) }
func (s *Scanner) UnexpToken(expectedToken T.GTokenType, identifierSpec, symbolSpec, symbolName, message string, isError int) {
	ScannerUnexpToken(s, expectedToken, identifierSpec, symbolSpec, symbolName, message, isError)
}
func (s *Scanner) Error(format string, v ...VArg) { ScannerError(s, format, v) }
func (s *Scanner) Warn(format string, v ...VArg)  { ScannerWarn(s, format, v) }

type SeekType Enum

const (
	SEEK_CUR SeekType = iota
	SEEK_SET
	SEEK_END
)

type Sequence struct{}

var (
	SequenceNew func(dataDestroy O.DestroyNotify) *Sequence

	SequenceFree             func(s *Sequence)
	SequenceGetLength        func(s *Sequence) int
	SequenceForeach          func(s *Sequence, f T.GFunc, userData T.Gpointer)
	SequenceSort             func(s *Sequence, cmpFunc T.GCompareDataFunc, cmpData T.Gpointer)
	SequenceSortIter         func(s *Sequence, cmpFunc SequenceIterCompareFunc, cmpData T.Gpointer)
	SequenceGetBeginIter     func(s *Sequence) *SequenceIter
	SequenceGetEndIter       func(s *Sequence) *SequenceIter
	SequenceGetIterAtPos     func(s *Sequence, pos int) *SequenceIter
	SequenceAppend           func(s *Sequence, data T.Gpointer) *SequenceIter
	SequencePrepend          func(s *Sequence, data T.Gpointer) *SequenceIter
	SequenceInsertSorted     func(s *Sequence, data T.Gpointer, cmpFunc T.GCompareDataFunc, cmpData T.Gpointer) *SequenceIter
	SequenceInsertSortedIter func(s *Sequence, data T.Gpointer, iterCmp SequenceIterCompareFunc, cmpData T.Gpointer) *SequenceIter
	SequenceSearch           func(s *Sequence, data T.Gpointer, cmpFunc T.GCompareDataFunc, cmpData T.Gpointer) *SequenceIter
	SequenceSearchIter       func(s *Sequence, data T.Gpointer, iterCmp SequenceIterCompareFunc, cmpData T.Gpointer) *SequenceIter
	SequenceLookup           func(s *Sequence, data T.Gpointer, cmpFunc T.GCompareDataFunc, cmpData T.Gpointer) *SequenceIter
	SequenceLookupIter       func(s *Sequence, data T.Gpointer, iterCmp SequenceIterCompareFunc, cmpData T.Gpointer) *SequenceIter
)

func (s *Sequence) Free()                                  { SequenceFree(s) }
func (s *Sequence) Length() int                            { return SequenceGetLength(s) }
func (s *Sequence) Foreach(f T.GFunc, userData T.Gpointer) { SequenceForeach(s, f, userData) }
func (s *Sequence) Sort(cmpFunc T.GCompareDataFunc, cmpData T.Gpointer) {
	SequenceSort(s, cmpFunc, cmpData)
}
func (s *Sequence) SortIter(cmpFunc SequenceIterCompareFunc, cmpData T.Gpointer) {
	SequenceSortIter(s, cmpFunc, cmpData)
}
func (s *Sequence) BeginIter() *SequenceIter              { return SequenceGetBeginIter(s) }
func (s *Sequence) EndIter() *SequenceIter                { return SequenceGetEndIter(s) }
func (s *Sequence) IterAtPos(pos int) *SequenceIter       { return SequenceGetIterAtPos(s, pos) }
func (s *Sequence) Append(data T.Gpointer) *SequenceIter  { return SequenceAppend(s, data) }
func (s *Sequence) Prepend(data T.Gpointer) *SequenceIter { return SequencePrepend(s, data) }
func (s *Sequence) InsertSorted(data T.Gpointer, cmpFunc T.GCompareDataFunc, cmpData T.Gpointer) *SequenceIter {
	return SequenceInsertSorted(s, data, cmpFunc, cmpData)
}
func (s *Sequence) InsertSortedIter(data T.Gpointer, iterCmp SequenceIterCompareFunc, cmpData T.Gpointer) *SequenceIter {
	return SequenceInsertSortedIter(s, data, iterCmp, cmpData)
}
func (s *Sequence) Search(data T.Gpointer, cmpFunc T.GCompareDataFunc, cmpData T.Gpointer) *SequenceIter {
	return SequenceSearch(s, data, cmpFunc, cmpData)
}
func (s *Sequence) SearchIter(data T.Gpointer, iterCmp SequenceIterCompareFunc, cmpData T.Gpointer) *SequenceIter {
	return SequenceSearchIter(s, data, iterCmp, cmpData)
}
func (s *Sequence) Lookup(data T.Gpointer, cmpFunc T.GCompareDataFunc, cmpData T.Gpointer) *SequenceIter {
	return SequenceLookup(s, data, cmpFunc, cmpData)
}
func (s *Sequence) LookupIter(data T.Gpointer, iterCmp SequenceIterCompareFunc, cmpData T.Gpointer) *SequenceIter {
	return SequenceLookupIter(s, data, iterCmp, cmpData)
}

type (
	SequenceIter SequenceNode

	SequenceNode struct{}

	SequenceIterCompareFunc func(
		a, b *SequenceIter, data T.Gpointer) int
)

var (
	SequenceForeachRange     func(s *SequenceIter, end *SequenceIter, f T.GFunc, userData T.Gpointer)
	SequenceInsertBefore     func(s *SequenceIter, data T.Gpointer) *SequenceIter
	SequenceMove             func(s *SequenceIter, dest *SequenceIter)
	SequenceSwap             func(a *SequenceIter, b *SequenceIter)
	SequenceSortChanged      func(s *SequenceIter, cmpFunc T.GCompareDataFunc, cmpData T.Gpointer)
	SequenceSortChangedIter  func(s *SequenceIter, iterCmp SequenceIterCompareFunc, cmpData T.Gpointer)
	SequenceRemove           func(s *SequenceIter)
	SequenceRemoveRange      func(s *SequenceIter, end *SequenceIter)
	SequenceMoveRange        func(s *SequenceIter, begin *SequenceIter, end *SequenceIter)
	SequenceGet              func(s *SequenceIter) T.Gpointer
	SequenceSet              func(s *SequenceIter, data T.Gpointer)
	SequenceIterIsBegin      func(s *SequenceIter) bool
	SequenceIterIsEnd        func(s *SequenceIter) bool
	SequenceIterNext         func(s *SequenceIter) *SequenceIter
	SequenceIterPrev         func(s *SequenceIter) *SequenceIter
	SequenceIterGetPosition  func(s *SequenceIter) int
	SequenceIterMove         func(s *SequenceIter, delta int) *SequenceIter
	SequenceIterGetSequence  func(s *SequenceIter) *Sequence
	SequenceIterCompare      func(s *SequenceIter, s2 *SequenceIter) int
	SequenceRangeGetMidpoint func(s *SequenceIter, end *SequenceIter) *SequenceIter
)

func (s *SequenceIter) ForeachRange(end *SequenceIter, f T.GFunc, userData T.Gpointer) {
	SequenceForeachRange(s, end, f, userData)
}
func (s *SequenceIter) InsertBefore(data T.Gpointer) *SequenceIter {
	return SequenceInsertBefore(s, data)
}
func (s *SequenceIter) Move(dest *SequenceIter) { SequenceMove(s, dest) }
func (s *SequenceIter) Swap(s2 *SequenceIter)   { SequenceSwap(s, s2) }
func (s *SequenceIter) SortChanged(cmpFunc T.GCompareDataFunc, cmpData T.Gpointer) {
	SequenceSortChanged(s, cmpFunc, cmpData)
}
func (s *SequenceIter) SortChangedIter(iterCmp SequenceIterCompareFunc, cmpData T.Gpointer) {
	SequenceSortChangedIter(s, iterCmp, cmpData)
}
func (s *SequenceIter) Remove()                       { SequenceRemove(s) }
func (s *SequenceIter) RemoveRange(end *SequenceIter) { SequenceRemoveRange(s, end) }
func (s *SequenceIter) MoveRange(begin *SequenceIter, end *SequenceIter) {
	SequenceMoveRange(s, begin, end)
}
func (s *SequenceIter) Get() T.Gpointer                { return SequenceGet(s) }
func (s *SequenceIter) Set(data T.Gpointer)            { SequenceSet(s, data) }
func (s *SequenceIter) IsBegin() bool                  { return SequenceIterIsBegin(s) }
func (s *SequenceIter) IsEnd() bool                    { return SequenceIterIsEnd(s) }
func (s *SequenceIter) Next() *SequenceIter            { return SequenceIterNext(s) }
func (s *SequenceIter) Prev() *SequenceIter            { return SequenceIterPrev(s) }
func (s *SequenceIter) GetPosition() int               { return SequenceIterGetPosition(s) }
func (s *SequenceIter) MoveBy(delta int) *SequenceIter { return SequenceIterMove(s, delta) } //TODO(t):renamed name==action?
func (s *SequenceIter) GetSequence() *Sequence         { return SequenceIterGetSequence(s) }
func (s *SequenceIter) Compare(s2 *SequenceIter) int   { return SequenceIterCompare(s, s2) }
func (s *SequenceIter) RangeGetMidpoint(end *SequenceIter) *SequenceIter {
	return SequenceRangeGetMidpoint(s, end)
}

var (
	ShellErrorQuark func() Quark
	ShellQuote      func(unquotedString string) string
	ShellUnquote    func(quotedString string, e **Error) string
	ShellParseArgv  func(commandLine string, argcp *int, argvp ***T.Gchar, e **Error) bool
)

type SliceConfig Enum

const (
	SLICE_CONFIG_ALWAYS_MALLOC SliceConfig = iota + 1
	SLICE_CONFIG_BYPASS_MAGAZINES
	SLICE_CONFIG_WORKING_SET_MSECS
	SLICE_CONFIG_COLOR_INCREMENT
	SLICE_CONFIG_CHUNK_SIZES
	SLICE_CONFIG_CONTENTION_COUNTER
)

var (
	SliceAlloc               func(blockSize T.Gsize) T.Gpointer
	SliceAlloc0              func(blockSize T.Gsize) T.Gpointer
	SliceCopy                func(blockSize T.Gsize, memBlock T.Gconstpointer) T.Gpointer
	SliceFree1               func(blockSize T.Gsize, memBlock T.Gpointer)
	SliceFreeChainWithOffset func(blockSize T.Gsize, memChain T.Gpointer, nextOffset T.Gsize)

	SliceGetConfig      func(s SliceConfig) int64
	SliceGetConfigState func(s SliceConfig, address int64, nValues *uint) *int64
	SliceSetConfig      func(s SliceConfig, value int64)
)

func (s SliceConfig) Get() int64 { return SliceGetConfig(s) }
func (s SliceConfig) GetState(address int64, nValues *uint) *int64 {
	return SliceGetConfigState(s, address, nValues)
}
func (s SliceConfig) Set(value int64) { SliceSetConfig(s, value) }

type SList T.SList

var (
	SlistAlloc         func() *SList
	SlistPopAllocator  func()
	SlistPushAllocator func(dummy T.Gpointer)

	SlistAppend               func(s *SList, data T.Gpointer) *SList
	SlistConcat               func(s *SList, s2 *SList) *SList
	SlistCopy                 func(s *SList) *SList
	SlistDeleteLink           func(s *SList, link *SList) *SList
	SlistFind                 func(s *SList, data T.Gconstpointer) *SList
	SlistFindCustom           func(s *SList, data T.Gconstpointer, f T.GCompareFunc) *SList
	SlistForeach              func(s *SList, f T.GFunc, userData T.Gpointer)
	SlistFree                 func(s *SList)
	SlistFree1                func(s *SList)
	SlistFreeFull             func(s *SList, freeFunc O.DestroyNotify)
	SlistIndex                func(s *SList, data T.Gconstpointer) int
	SlistInsert               func(s *SList, data T.Gpointer, position int) *SList
	SlistInsertBefore         func(s *SList, sibling *SList, data T.Gpointer) *SList
	SlistInsertSorted         func(s *SList, data T.Gpointer, f T.GCompareFunc) *SList
	SlistInsertSortedWithData func(s *SList, data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) *SList
	SlistLast                 func(s *SList) *SList
	SlistLength               func(s *SList) uint
	SlistNth                  func(s *SList, n uint) *SList
	SlistNthData              func(s *SList, n uint) T.Gpointer
	SlistPosition             func(s *SList, llink *SList) int
	SlistPrepend              func(s *SList, data T.Gpointer) *SList
	SlistRemove               func(s *SList, data T.Gconstpointer) *SList
	SlistRemoveAll            func(s *SList, data T.Gconstpointer) *SList
	SlistRemoveLink           func(s *SList, link *SList) *SList
	SlistReverse              func(s *SList) *SList
	SlistSort                 func(s *SList, compareFunc T.GCompareFunc) *SList
	SlistSortWithData         func(s *SList, compareFunc T.GCompareDataFunc, userData T.Gpointer) *SList
)

func (s *SList) Append(data T.Gpointer) *SList    { return SlistAppend(s, data) }
func (s *SList) Concat(s2 *SList) *SList          { return SlistConcat(s, s2) }
func (s *SList) Copy() *SList                     { return SlistCopy(s) }
func (s *SList) DeleteLink(link *SList) *SList    { return SlistDeleteLink(s, link) }
func (s *SList) Find(data T.Gconstpointer) *SList { return SlistFind(s, data) }
func (s *SList) FindCustom(data T.Gconstpointer, f T.GCompareFunc) *SList {
	return SlistFindCustom(s, data, f)
}
func (s *SList) Foreach(f T.GFunc, userData T.Gpointer)      { SlistForeach(s, f, userData) }
func (s *SList) Free()                                       { SlistFree(s) }
func (s *SList) Free1()                                      { SlistFree1(s) }
func (s *SList) FreeFull(freeFunc O.DestroyNotify)           { SlistFreeFull(s, freeFunc) }
func (s *SList) Index(data T.Gconstpointer) int              { return SlistIndex(s, data) }
func (s *SList) Insert(data T.Gpointer, position int) *SList { return SlistInsert(s, data, position) }
func (s *SList) InsertBefore(sibling *SList, data T.Gpointer) *SList {
	return SlistInsertBefore(s, sibling, data)
}
func (s *SList) InsertSorted(data T.Gpointer, f T.GCompareFunc) *SList {
	return SlistInsertSorted(s, data, f)
}
func (s *SList) InsertSortedWithData(data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) *SList {
	return SlistInsertSortedWithData(s, data, f, userData)
}
func (s *SList) Last() *SList                           { return SlistLast(s) }
func (s *SList) Length() uint                           { return SlistLength(s) }
func (s *SList) Nth(n uint) *SList                      { return SlistNth(s, n) }
func (s *SList) NthData(n uint) T.Gpointer              { return SlistNthData(s, n) }
func (s *SList) Position(llink *SList) int              { return SlistPosition(s, llink) }
func (s *SList) Prepend(data T.Gpointer) *SList         { return SlistPrepend(s, data) }
func (s *SList) Remove(data T.Gconstpointer) *SList     { return SlistRemove(s, data) }
func (s *SList) RemoveAll(data T.Gconstpointer) *SList  { return SlistRemoveAll(s, data) }
func (s *SList) RemoveLink(link *SList) *SList          { return SlistRemoveLink(s, link) }
func (s *SList) Reverse() *SList                        { return SlistReverse(s) }
func (s *SList) Sort(compareFunc T.GCompareFunc) *SList { return SlistSort(s, compareFunc) }
func (s *SList) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) *SList {
	return SlistSortWithData(s, compareFunc, userData)
}

type Source O.Source

var (
	SourceNew                   func(sourceFuncs *O.SourceFuncs, structSize uint) *Source
	SourceRemove                func(tag uint) bool
	SourceRemoveByFuncsUserData func(funcs *O.SourceFuncs, userData T.Gpointer) bool
	SourceRemoveByUserData      func(userData T.Gpointer) bool
	SourceSetNameById           func(tag uint, name string)

	SourceAddChildSource      func(s *Source, childSource *Source)
	SourceAddPoll             func(s *Source, fd *T.GPollFD)
	SourceAttach              func(s *Source, context *MainContext) uint
	SourceDestroy             func(s *Source)
	SourceGetCanRecurse       func(s *Source) bool
	SourceGetContext          func(s *Source) *MainContext
	SourceGetCurrentTime      func(s *Source, timeval *TimeVal)
	SourceGetId               func(s *Source) uint
	SourceGetName             func(s *Source) string
	SourceGetPriority         func(s *Source) int
	SourceGetTime             func(s *Source) int64
	SourceIsDestroyed         func(s *Source) bool
	SourceRef                 func(s *Source) *Source
	SourceRemoveChildSource   func(s *Source, childSource *Source)
	SourceRemovePoll          func(s *Source, fd *T.GPollFD)
	SourceSetCallback         func(s *Source, f O.SourceFunc, data T.Gpointer, notify O.DestroyNotify)
	SourceSetCallbackIndirect func(s *Source, callbackData T.Gpointer, callbackFuncs *O.SourceCallbackFuncs)
	SourceSetCanRecurse       func(s *Source, canRecurse bool)
	SourceSetFuncs            func(s *Source, funcs *O.SourceFuncs)
	SourceSetName             func(s *Source, name string)
	SourceSetPriority         func(s *Source, priority int)
	SourceUnref               func(s *Source)
)

func (s *Source) AddChildSource(childSource *Source)    { SourceAddChildSource(s, childSource) }
func (s *Source) AddPoll(fd *T.GPollFD)                 { SourceAddPoll(s, fd) }
func (s *Source) Attach(context *MainContext) uint      { return SourceAttach(s, context) }
func (s *Source) Destroy()                              { SourceDestroy(s) }
func (s *Source) GetCanRecurse() bool                   { return SourceGetCanRecurse(s) }
func (s *Source) GetContext() *MainContext              { return SourceGetContext(s) }
func (s *Source) GetCurrentTime(timeval *TimeVal)       { SourceGetCurrentTime(s, timeval) }
func (s *Source) GetId() uint                           { return SourceGetId(s) }
func (s *Source) GetName() string                       { return SourceGetName(s) }
func (s *Source) GetPriority() int                      { return SourceGetPriority(s) }
func (s *Source) GetTime() int64                        { return SourceGetTime(s) }
func (s *Source) IsDestroyed() bool                     { return SourceIsDestroyed(s) }
func (s *Source) Ref() *Source                          { return SourceRef(s) }
func (s *Source) RemoveChildSource(childSource *Source) { SourceRemoveChildSource(s, childSource) }
func (s *Source) RemovePoll(fd *T.GPollFD)              { SourceRemovePoll(s, fd) }
func (s *Source) SetCallback(f O.SourceFunc, data T.Gpointer, notify O.DestroyNotify) {
	SourceSetCallback(s, f, data, notify)
}
func (s *Source) SetCallbackIndirect(callbackData T.Gpointer, callbackFuncs *O.SourceCallbackFuncs) {
	SourceSetCallbackIndirect(s, callbackData, callbackFuncs)
}
func (s *Source) SetCanRecurse(canRecurse bool) { SourceSetCanRecurse(s, canRecurse) }
func (s *Source) SetFuncs(funcs *O.SourceFuncs) { SourceSetFuncs(s, funcs) }
func (s *Source) SetName(name string)           { SourceSetName(s, name) }
func (s *Source) SetPriority(priority int)      { SourceSetPriority(s, priority) }
func (s *Source) Unref()                        { SourceUnref(s) }

type SpawnFlags Enum

const (
	SPAWN_LEAVE_DESCRIPTORS_OPEN SpawnFlags = 1 << iota
	SPAWN_DO_NOT_REAP_CHILD
	SPAWN_SEARCH_PATH
	SPAWN_STDOUT_TO_DEV_NULL
	SPAWN_STDERR_TO_DEV_NULL
	SPAWN_CHILD_INHERITS_STDIN
	SPAWN_FILE_AND_ARGV_ZERO
)

type SpawnChildSetupFunc func(userData T.Gpointer)

var (
	SpawnAsync            func(workingDirectory string, argv, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc, userData T.Gpointer, childPid *T.GPid, e **Error) bool
	SpawnAsyncWithPipes   func(workingDirectory string, argv, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc, userData T.Gpointer, childPid *T.GPid, standardInput *int, standardOutput *int, standardError *int, e **Error) bool
	SpawnClosePid         func(pid T.GPid)
	SpawnCommandLineAsync func(commandLine string, e **Error) bool
	SpawnCommandLineSync  func(commandLine string, standardOutput **T.Gchar, standardError **T.Gchar, exitStatus *int, e **Error) bool
	SpawnErrorQuark       func() Quark
	SpawnSync             func(workingDirectory string, argv, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc, userData T.Gpointer, standardOutput **T.Gchar, standardError **T.Gchar, exitStatus *int, e **Error) bool
)

type StaticMutex *Mutex

var (
	StaticMutexInit func(s *StaticMutex) // Can't methodize **!
	StaticMutexFree func(s *StaticMutex) // Can't methodize **!
)

type StaticPrivate struct {
	Index uint
}

var (
	StaticPrivateFree func(s *StaticPrivate)
	StaticPrivateGet  func(s *StaticPrivate) T.Gpointer
	StaticPrivateInit func(s *StaticPrivate)
	StaticPrivateSet  func(s *StaticPrivate, data T.Gpointer, notify O.DestroyNotify)
)

func (s *StaticPrivate) Free()           { StaticPrivateFree(s) }
func (s *StaticPrivate) Get() T.Gpointer { return StaticPrivateGet(s) }
func (s *StaticPrivate) Init()           { StaticPrivateInit(s) }
func (s *StaticPrivate) Set(data T.Gpointer, notify O.DestroyNotify) {
	StaticPrivateSet(s, data, notify)
}

type StaticRecMutex struct {
	Mutex StaticMutex
	Depth uint
	Owner SystemThread
}

var (
	StaticRecMutexFree       func(s *StaticRecMutex)
	StaticRecMutexInit       func(s *StaticRecMutex)
	StaticRecMutexLock       func(s *StaticRecMutex)
	StaticRecMutexLockFull   func(s *StaticRecMutex, depth uint)
	StaticRecMutexTrylock    func(s *StaticRecMutex) bool
	StaticRecMutexUnlock     func(s *StaticRecMutex)
	StaticRecMutexUnlockFull func(s *StaticRecMutex) uint
)

func (s *StaticRecMutex) Free()               { StaticRecMutexFree(s) }
func (s *StaticRecMutex) Init()               { StaticRecMutexInit(s) }
func (s *StaticRecMutex) Lock()               { StaticRecMutexLock(s) }
func (s *StaticRecMutex) LockFull(depth uint) { StaticRecMutexLockFull(s, depth) }
func (s *StaticRecMutex) Trylock() bool       { return StaticRecMutexTrylock(s) }
func (s *StaticRecMutex) Unlock()             { StaticRecMutexUnlock(s) }
func (s *StaticRecMutex) UnlockFull() uint    { return StaticRecMutexUnlockFull(s) }

type StaticRWLock struct {
	Mutex       StaticMutex
	ReadCond    *Cond
	WriteCond   *Cond
	ReadCounter uint
	HaveWriter  T.Gboolean
	WantToRead  uint
	WantToWrite uint
}

var (
	StaticRwLockFree          func(s *StaticRWLock)
	StaticRwLockInit          func(s *StaticRWLock)
	StaticRwLockReaderLock    func(s *StaticRWLock)
	StaticRwLockReaderTrylock func(s *StaticRWLock) bool
	StaticRwLockReaderUnlock  func(s *StaticRWLock)
	StaticRwLockWriterLock    func(s *StaticRWLock)
	StaticRwLockWriterTrylock func(s *StaticRWLock) bool
	StaticRwLockWriterUnlock  func(s *StaticRWLock)
)

func (s *StaticRWLock) Free()               { StaticRwLockFree(s) }
func (s *StaticRWLock) Init()               { StaticRwLockInit(s) }
func (s *StaticRWLock) ReaderLock()         { StaticRwLockReaderLock(s) }
func (s *StaticRWLock) ReaderTrylock() bool { return StaticRwLockReaderTrylock(s) }
func (s *StaticRWLock) ReaderUnlock()       { StaticRwLockReaderUnlock(s) }
func (s *StaticRWLock) WriterLock()         { StaticRwLockWriterLock(s) }
func (s *StaticRWLock) WriterTrylock() bool { return StaticRwLockWriterTrylock(s) }
func (s *StaticRWLock) WriterUnlock()       { StaticRwLockWriterUnlock(s) }

type SystemThread struct {
	// Union
	// Data[4] Char
	DummyDouble float64
	// DummyPointer *Void
	// DummyLong Long
}

type String struct {
	Str          *T.Gchar
	Len          T.Gsize
	AllocatedLen T.Gsize
}

var (
	StringNew      func(init string) *String
	StringNewLen   func(init string, leng T.Gssize) *String
	StringSizedNew func(dflSize T.Gsize) *String

	StringAppend           func(s *String, val string) *String
	StringAppendC          func(s *String, c T.Gchar) *String
	StringAppendLen        func(s *String, val string, leng T.Gssize) *String
	StringAppendPrintf     func(s *String, format string, v ...VArg)
	StringAppendUnichar    func(s *String, wc Unichar) *String
	StringAppendUriEscaped func(s *String, unescaped, reservedCharsAllowed string, allowUtf8 bool) *String
	StringAppendVprintf    func(s *String, format string, args T.VaList)
	StringAsciiDown        func(s *String) *String
	StringAsciiUp          func(s *String) *String
	StringAssign           func(s *String, rval string) *String
	StringDown             func(s *String) *String
	StringEqual            func(s *String, s2 *String) bool
	StringErase            func(s *String, pos T.Gssize, leng T.Gssize) *String
	StringFree             func(s *String, freeSegment bool) string
	StringHash             func(s *String) uint
	StringInsert           func(s *String, pos T.Gssize, val string) *String
	StringInsertC          func(s *String, pos T.Gssize, c T.Gchar) *String
	StringInsertLen        func(s *String, pos T.Gssize, val string, leng T.Gssize) *String
	StringInsertUnichar    func(s *String, pos T.Gssize, wc Unichar) *String
	StringOverwrite        func(s *String, pos T.Gsize, val string) *String
	StringOverwriteLen     func(s *String, pos T.Gsize, val string, leng T.Gssize) *String
	StringPrepend          func(s *String, val string) *String
	StringPrependC         func(s *String, c T.Gchar) *String
	StringPrependLen       func(s *String, val string, leng T.Gssize) *String
	StringPrependUnichar   func(s *String, wc Unichar) *String
	StringPrintf           func(s *String, format string, v ...VArg)
	StringSetSize          func(s *String, leng T.Gsize) *String
	StringTruncate         func(s *String, leng T.Gsize) *String
	StringUp               func(s *String) *String
	StringVprintf          func(s *String, format string, args T.VaList)
)

func (s *String) Append(val string) *String                   { return StringAppend(s, val) }
func (s *String) AppendC(c T.Gchar) *String                   { return StringAppendC(s, c) }
func (s *String) AppendLen(val string, leng T.Gssize) *String { return StringAppendLen(s, val, leng) }
func (s *String) AppendPrintf(format string, v ...VArg)       { StringAppendPrintf(s, format, v) }
func (s *String) AppendUnichar(wc Unichar) *String            { return StringAppendUnichar(s, wc) }
func (s *String) AppendUriEscaped(unescaped, reservedCharsAllowed string, allowUtf8 bool) *String {
	return StringAppendUriEscaped(s, unescaped, reservedCharsAllowed, allowUtf8)
}
func (s *String) AppendVprintf(format string, args T.VaList) { StringAppendVprintf(s, format, args) }
func (s *String) AsciiDown() *String                         { return StringAsciiDown(s) }
func (s *String) AsciiUp() *String                           { return StringAsciiUp(s) }
func (s *String) Assign(rval string) *String                 { return StringAssign(s, rval) }
func (s *String) Down() *String                              { return StringDown(s) }
func (s *String) Equal(s2 *String) bool                      { return StringEqual(s, s2) }
func (s *String) Erase(pos T.Gssize, leng T.Gssize) *String  { return StringErase(s, pos, leng) }
func (s *String) Free(freeSegment bool) string               { return StringFree(s, freeSegment) }
func (s *String) Hash() uint                                 { return StringHash(s) }
func (s *String) Insert(pos T.Gssize, val string) *String    { return StringInsert(s, pos, val) }
func (s *String) InsertC(pos T.Gssize, c T.Gchar) *String    { return StringInsertC(s, pos, c) }
func (s *String) InsertLen(pos T.Gssize, val string, leng T.Gssize) *String {
	return StringInsertLen(s, pos, val, leng)
}
func (s *String) InsertUnichar(pos T.Gssize, wc Unichar) *String {
	return StringInsertUnichar(s, pos, wc)
}
func (s *String) Overwrite(pos T.Gsize, val string) *String { return StringOverwrite(s, pos, val) }
func (s *String) OverwriteLen(pos T.Gsize, val string, leng T.Gssize) *String {
	return StringOverwriteLen(s, pos, val, leng)
}
func (s *String) Prepend(val string) *String                   { return StringPrepend(s, val) }
func (s *String) PrependC(c T.Gchar) *String                   { return StringPrependC(s, c) }
func (s *String) PrependLen(val string, leng T.Gssize) *String { return StringPrependLen(s, val, leng) }
func (s *String) PrependUnichar(wc Unichar) *String            { return StringPrependUnichar(s, wc) }
func (s *String) Printf(format string, v ...VArg)              { StringPrintf(s, format, v) }
func (s *String) SetSize(leng T.Gsize) *String                 { return StringSetSize(s, leng) }
func (s *String) Truncate(leng T.Gsize) *String                { return StringTruncate(s, leng) }
func (s *String) Up() *String                                  { return StringUp(s) }
func (s *String) Vprintf(format string, args T.VaList)         { StringVprintf(s, format, args) }

type StringChunk struct{}

var (
	StringChunkNew func(size T.Gsize) *StringChunk

	StringChunkFree        func(s *StringChunk)
	StringChunkClear       func(s *StringChunk)
	StringChunkInsert      func(s *StringChunk, str string) string
	StringChunkInsertLen   func(s *StringChunk, str string, leng T.Gssize) string
	StringChunkInsertConst func(s *StringChunk, str string) string
)

func (s *StringChunk) Free()                    { StringChunkFree(s) }
func (s *StringChunk) Clear()                   { StringChunkClear(s) }
func (s *StringChunk) Insert(str string) string { return StringChunkInsert(s, str) }
func (s *StringChunk) InsertLen(str string, leng T.Gssize) string {
	return StringChunkInsertLen(s, str, leng)
}
func (s *StringChunk) InsertConst(str string) string { return StringChunkInsertConst(s, str) }
