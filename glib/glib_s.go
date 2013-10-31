// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

var (
	ScannerNew func(
		configTempl *T.GScannerConfig) *T.GScanner

	ScannerDestroy func(scanner *T.GScanner)

	ScannerInputFile func(scanner *T.GScanner,
		inputFd int)

	ScannerSyncFileOffset func(scanner *T.GScanner)

	ScannerInputText func(scanner *T.GScanner,
		text string, textLen uint)

	ScannerGetNextToken func(scanner *T.GScanner) T.GTokenType

	ScannerPeekNextToken func(scanner *T.GScanner) T.GTokenType

	ScannerCurToken func(scanner *T.GScanner) T.GTokenType

	ScannerCurValue func(scanner *T.GScanner) T.GTokenValue

	ScannerCurLine func(scanner *T.GScanner) uint

	ScannerCurPosition func(scanner *T.GScanner) uint

	ScannerEof func(scanner *T.GScanner) bool

	ScannerSetScope func(scanner *T.GScanner,
		scopeId uint) uint

	ScannerScopeAddSymbol func(scanner *T.GScanner,
		scopeId uint, symbol string, value T.Gpointer)

	ScannerScopeRemoveSymbol func(scanner *T.GScanner,
		scopeId uint, symbol string)

	ScannerScopeLookupSymbol func(scanner *T.GScanner,
		scopeId uint, symbol string) T.Gpointer

	ScannerScopeForeachSymbol func(scanner *T.GScanner,
		scopeId uint, f T.GHFunc, userData T.Gpointer)

	ScannerLookupSymbol func(scanner *T.GScanner,
		symbol string) T.Gpointer

	ScannerUnexpToken func(scanner *T.GScanner,
		expectedToken T.GTokenType,
		identifierSpec, symbolSpec, symbolName, message string,
		isError int)

	ScannerError func(scanner *T.GScanner,
		format string, v ...VArg)

	ScannerWarn func(scanner *T.GScanner,
		format string, v ...VArg)
)

var (
	SequenceNew func(dataDestroy T.GDestroyNotify) *T.GSequence

	SequenceFree func(seq *T.GSequence)

	SequenceGetLength func(seq *T.GSequence) int

	SequenceForeach func(seq *T.GSequence,
		f T.GFunc, userData T.Gpointer)

	SequenceForeachRange func(begin *T.GSequenceIter,
		end *T.GSequenceIter, f T.GFunc, userData T.Gpointer)

	SequenceSort func(seq *T.GSequence,
		cmpFunc T.GCompareDataFunc, cmpData T.Gpointer)

	SequenceSortIter func(seq *T.GSequence,
		cmpFunc T.GSequenceIterCompareFunc, cmpData T.Gpointer)

	SequenceGetBeginIter func(seq *T.GSequence) *T.GSequenceIter

	SequenceGetEndIter func(seq *T.GSequence) *T.GSequenceIter

	SequenceGetIterAtPos func(seq *T.GSequence,
		pos int) *T.GSequenceIter

	SequenceAppend func(seq *T.GSequence,
		data T.Gpointer) *T.GSequenceIter

	SequencePrepend func(seq *T.GSequence,
		data T.Gpointer) *T.GSequenceIter

	SequenceInsertBefore func(iter *T.GSequenceIter,
		data T.Gpointer) *T.GSequenceIter

	SequenceMove func(src *T.GSequenceIter,
		dest *T.GSequenceIter)

	SequenceSwap func(a *T.GSequenceIter,
		b *T.GSequenceIter)

	SequenceInsertSorted func(seq *T.GSequence,
		data T.Gpointer,
		cmpFunc T.GCompareDataFunc,
		cmpData T.Gpointer) *T.GSequenceIter

	SequenceInsertSortedIter func(seq *T.GSequence,
		data T.Gpointer,
		iterCmp T.GSequenceIterCompareFunc,
		cmpData T.Gpointer) *T.GSequenceIter

	SequenceSortChanged func(iter *T.GSequenceIter,
		cmpFunc T.GCompareDataFunc,
		cmpData T.Gpointer)

	SequenceSortChangedIter func(iter *T.GSequenceIter,
		iterCmp T.GSequenceIterCompareFunc,
		cmpData T.Gpointer)

	SequenceRemove func(iter *T.GSequenceIter)

	SequenceRemoveRange func(begin *T.GSequenceIter,
		end *T.GSequenceIter)

	SequenceMoveRange func(dest *T.GSequenceIter,
		begin *T.GSequenceIter,
		end *T.GSequenceIter)

	SequenceSearch func(seq *T.GSequence,
		data T.Gpointer,
		cmpFunc T.GCompareDataFunc,
		cmpData T.Gpointer) *T.GSequenceIter

	SequenceSearchIter func(seq *T.GSequence,
		data T.Gpointer,
		iterCmp T.GSequenceIterCompareFunc,
		cmpData T.Gpointer) *T.GSequenceIter

	SequenceLookup func(seq *T.GSequence,
		data T.Gpointer,
		cmpFunc T.GCompareDataFunc,
		cmpData T.Gpointer) *T.GSequenceIter

	SequenceLookupIter func(seq *T.GSequence,
		data T.Gpointer,
		iterCmp T.GSequenceIterCompareFunc,
		cmpData T.Gpointer) *T.GSequenceIter

	SequenceGet func(iter *T.GSequenceIter) T.Gpointer

	SequenceSet func(iter *T.GSequenceIter,
		data T.Gpointer)

	SequenceIterIsBegin func(iter *T.GSequenceIter) bool

	SequenceIterIsEnd func(iter *T.GSequenceIter) bool

	SequenceIterNext func(iter *T.GSequenceIter) *T.GSequenceIter

	SequenceIterPrev func(iter *T.GSequenceIter) *T.GSequenceIter

	SequenceIterGetPosition func(iter *T.GSequenceIter) int

	SequenceIterMove func(iter *T.GSequenceIter,
		delta int) *T.GSequenceIter

	SequenceIterGetSequence func(iter *T.GSequenceIter) *T.GSequence

	SequenceIterCompare func(a *T.GSequenceIter,
		b *T.GSequenceIter) int

	SequenceRangeGetMidpoint func(begin *T.GSequenceIter,
		end *T.GSequenceIter) *T.GSequenceIter
)

var (
	ShellErrorQuark func() T.GQuark

	ShellQuote func(unquotedString string) string

	ShellUnquote func(quotedString string,
		e **T.GError) string

	ShellParseArgv func(commandLine string,
		argcp *int,
		argvp ***T.Gchar,
		e **T.GError) bool
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

	SliceSetConfig      func(s SliceConfig, value int64)
	SliceGetConfig      func(s SliceConfig) int64
	SliceGetConfigState func(s SliceConfig, address int64, nValues *uint) *int64
)

type SList struct {
	Data T.Gpointer
	Next *SList
}

var (
	SlistAlloc         func() *SList
	SlistPopAllocator  func()
	SlistPushAllocator func(dummy T.Gpointer)

	slistAppend               func(s *SList, data T.Gpointer) *SList
	slistConcat               func(s *SList, s2 *SList) *SList
	slistCopy                 func(s *SList) *SList
	slistDeleteLink           func(s *SList, link *SList) *SList
	slistFind                 func(s *SList, data T.Gconstpointer) *SList
	slistFindCustom           func(s *SList, data T.Gconstpointer, f T.GCompareFunc) *SList
	slistForeach              func(s *SList, f T.GFunc, userData T.Gpointer)
	slistFree                 func(s *SList)
	slistFree1                func(s *SList)
	slistFreeFull             func(s *SList, freeFunc T.GDestroyNotify)
	slistIndex                func(s *SList, data T.Gconstpointer) int
	slistInsert               func(s *SList, data T.Gpointer, position int) *SList
	slistInsertBefore         func(s *SList, sibling *SList, data T.Gpointer) *SList
	slistInsertSorted         func(s *SList, data T.Gpointer, f T.GCompareFunc) *SList
	slistInsertSortedWithData func(s *SList, data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) *SList
	slistLast                 func(s *SList) *SList
	slistLength               func(s *SList) uint
	slistNth                  func(s *SList, n uint) *SList
	slistNthData              func(s *SList, n uint) T.Gpointer
	slistPosition             func(s *SList, llink *SList) int
	slistPrepend              func(s *SList, data T.Gpointer) *SList
	slistRemove               func(s *SList, data T.Gconstpointer) *SList
	slistRemoveAll            func(s *SList, data T.Gconstpointer) *SList
	slistRemoveLink           func(s *SList, link *SList) *SList
	slistReverse              func(s *SList) *SList
	slistSort                 func(s *SList, compareFunc T.GCompareFunc) *SList
	slistSortWithData         func(s *SList, compareFunc T.GCompareDataFunc, userData T.Gpointer) *SList
)

func (s *SList) Append(data T.Gpointer) *SList    { return slistAppend(s, data) }
func (s *SList) Concat(s2 *SList) *SList          { return slistConcat(s, s2) }
func (s *SList) Copy() *SList                     { return slistCopy(s) }
func (s *SList) DeleteLink(link *SList) *SList    { return slistDeleteLink(s, link) }
func (s *SList) Find(data T.Gconstpointer) *SList { return slistFind(s, data) }
func (s *SList) FindCustom(data T.Gconstpointer, f T.GCompareFunc) *SList {
	return slistFindCustom(s, data, f)
}
func (s *SList) Foreach(f T.GFunc, userData T.Gpointer)      { slistForeach(s, f, userData) }
func (s *SList) Free()                                       { slistFree(s) }
func (s *SList) Free1()                                      { slistFree1(s) }
func (s *SList) FreeFull(freeFunc T.GDestroyNotify)          { slistFreeFull(s, freeFunc) }
func (s *SList) Index(data T.Gconstpointer) int              { return slistIndex(s, data) }
func (s *SList) Insert(data T.Gpointer, position int) *SList { return slistInsert(s, data, position) }
func (s *SList) InsertBefore(sibling *SList, data T.Gpointer) *SList {
	return slistInsertBefore(s, sibling, data)
}
func (s *SList) InsertSorted(data T.Gpointer, f T.GCompareFunc) *SList {
	return slistInsertSorted(s, data, f)
}
func (s *SList) InsertSortedWithData(data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) *SList {
	return slistInsertSortedWithData(s, data, f, userData)
}
func (s *SList) Last() *SList                           { return slistLast(s) }
func (s *SList) Length() uint                           { return slistLength(s) }
func (s *SList) Nth(n uint) *SList                      { return slistNth(s, n) }
func (s *SList) NthData(n uint) T.Gpointer              { return slistNthData(s, n) }
func (s *SList) Position(llink *SList) int              { return slistPosition(s, llink) }
func (s *SList) Prepend(data T.Gpointer) *SList         { return slistPrepend(s, data) }
func (s *SList) Remove(data T.Gconstpointer) *SList     { return slistRemove(s, data) }
func (s *SList) RemoveAll(data T.Gconstpointer) *SList  { return slistRemoveAll(s, data) }
func (s *SList) RemoveLink(link *SList) *SList          { return slistRemoveLink(s, link) }
func (s *SList) Reverse() *SList                        { return slistReverse(s) }
func (s *SList) Sort(compareFunc T.GCompareFunc) *SList { return slistSort(s, compareFunc) }
func (s *SList) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) *SList {
	return slistSortWithData(s, compareFunc, userData)
}

type Source O.Source

var (
	SourceNew func(sourceFuncs *O.SourceFuncs,
		structSize uint) *O.Source

	SourceRemove func(tag uint) bool

	SourceRemoveByUserData func(userData T.Gpointer) bool

	SourceRemoveByFuncsUserData func(funcs *O.SourceFuncs,
		userData T.Gpointer) bool

	SourceRef func(source *O.Source) *O.Source

	SourceUnref func(source *O.Source)

	SourceAttach func(source *O.Source,
		context *T.GMainContext) uint

	SourceDestroy func(source *O.Source)

	SourceSetPriority func(source *O.Source,
		priority int)

	SourceGetPriority func(source *O.Source) int

	SourceSetCanRecurse func(source *O.Source,
		canRecurse bool)

	SourceGetCanRecurse func(source *O.Source) bool

	SourceGetId func(source *O.Source) uint

	SourceGetContext func(source *O.Source) *T.GMainContext

	SourceSetCallback func(source *O.Source,
		f O.SourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify)

	SourceSetFuncs func(source *O.Source,
		funcs *O.SourceFuncs)

	SourceIsDestroyed func(source *O.Source) bool

	SourceSetName func(source *O.Source,
		name string)

	SourceGetName func(source *O.Source) string

	SourceSetNameById func(tag uint,
		name string)

	SourceSetCallbackIndirect func(source *O.Source,
		callbackData T.Gpointer,
		callbackFuncs *O.SourceCallbackFuncs)

	SourceAddPoll func(source *O.Source,
		fd *T.GPollFD)

	SourceRemovePoll func(source *O.Source,
		fd *T.GPollFD)

	SourceAddChildSource func(source *O.Source,
		childSource *O.Source)

	SourceRemoveChildSource func(source *O.Source,
		childSource *O.Source)

	SourceGetCurrentTime func(source *O.Source,
		timeval *T.GTimeVal)

	SourceGetTime func(source *O.Source) int64
)

var (
	SpawnErrorQuark func() T.GQuark

	SpawnAsync func(workingDirectory string,
		argv []string,
		envp []string,
		flags T.GSpawnFlags,
		childSetup T.GSpawnChildSetupFunc,
		userData T.Gpointer,
		childPid *T.GPid,
		e **T.GError) bool

	SpawnAsyncWithPipes func(workingDirectory string,
		argv []string,
		envp []string,
		flags T.GSpawnFlags,
		childSetup T.GSpawnChildSetupFunc,
		userData T.Gpointer,
		childPid *T.GPid,
		standardInput *int,
		standardOutput *int,
		standardError *int,
		e **T.GError) bool

	SpawnSync func(workingDirectory string,
		argv []string,
		envp []string,
		flags T.GSpawnFlags,
		childSetup T.GSpawnChildSetupFunc,
		userData T.Gpointer,
		standardOutput **T.Gchar,
		standardError **T.Gchar,
		exitStatus *int,
		e **T.GError) bool

	SpawnCommandLineSync func(commandLine string,
		standardOutput **T.Gchar,
		standardError **T.Gchar,
		exitStatus *int,
		e **T.GError) bool

	SpawnCommandLineAsync func(commandLine string,
		e **T.GError) bool

	SpawnClosePid func(pid T.GPid)
)

type StaticMutex *T.GMutex

var (
	StaticMutexInit func(s *StaticMutex)
	StaticMutexFree func(s *StaticMutex)
)

var (
	StaticPrivateFree func(
		privateKey *T.GStaticPrivate)

	StaticPrivateGet func(
		privateKey *T.GStaticPrivate) T.Gpointer

	StaticPrivateInit func(privateKey *T.GStaticPrivate)

	StaticPrivateSet func(privateKey *T.GStaticPrivate,
		data T.Gpointer, notify T.GDestroyNotify)
)

type StaticRecMutex struct {
	Mutex StaticMutex
	Depth uint
	Owner SystemThread
}

var (
	staticRecMutexFree       func(s *StaticRecMutex)
	staticRecMutexInit       func(s *StaticRecMutex)
	staticRecMutexLock       func(s *StaticRecMutex)
	staticRecMutexLockFull   func(s *StaticRecMutex, depth uint)
	staticRecMutexTrylock    func(s *StaticRecMutex) bool
	staticRecMutexUnlock     func(s *StaticRecMutex)
	staticRecMutexUnlockFull func(s *StaticRecMutex) uint
)

func (s *StaticRecMutex) Free()               { staticRecMutexFree(s) }
func (s *StaticRecMutex) Init()               { staticRecMutexInit(s) }
func (s *StaticRecMutex) Lock()               { staticRecMutexLock(s) }
func (s *StaticRecMutex) LockFull(depth uint) { staticRecMutexLockFull(s, depth) }
func (s *StaticRecMutex) Trylock() bool       { return staticRecMutexTrylock(s) }
func (s *StaticRecMutex) Unlock()             { staticRecMutexUnlock(s) }
func (s *StaticRecMutex) UnlockFull() uint    { return staticRecMutexUnlockFull(s) }

type StaticRWLock struct {
	Mutex       StaticMutex
	ReadCond    *T.GCond
	WriteCond   *T.GCond
	ReadCounter uint
	HaveWriter  T.Gboolean
	WantToRead  uint
	WantToWrite uint
}

var (
	staticRwLockFree          func(s *StaticRWLock)
	staticRwLockInit          func(s *StaticRWLock)
	staticRwLockReaderLock    func(s *StaticRWLock)
	staticRwLockReaderTrylock func(s *StaticRWLock) bool
	staticRwLockReaderUnlock  func(s *StaticRWLock)
	staticRwLockWriterLock    func(s *StaticRWLock)
	staticRwLockWriterTrylock func(s *StaticRWLock) bool
	staticRwLockWriterUnlock  func(s *StaticRWLock)
)

func (s *StaticRWLock) Free()               { staticRwLockFree(s) }
func (s *StaticRWLock) Init()               { staticRwLockInit(s) }
func (s *StaticRWLock) ReaderLock()         { staticRwLockReaderLock(s) }
func (s *StaticRWLock) ReaderTrylock() bool { return staticRwLockReaderTrylock(s) }
func (s *StaticRWLock) ReaderUnlock()       { staticRwLockReaderUnlock(s) }
func (s *StaticRWLock) WriterLock()         { staticRwLockWriterLock(s) }
func (s *StaticRWLock) WriterTrylock() bool { return staticRwLockWriterTrylock(s) }
func (s *StaticRWLock) WriterUnlock()       { staticRwLockWriterUnlock(s) }

type SystemThread struct {
	// Union
	// Data[4] Char
	DummyDouble float64
	// DummyPointer *Void
	// DummyLong Long
}

var (
	StringNew func(init string) *T.GString

	StringNewLen func(init string,
		leng T.Gssize) *T.GString

	StringSizedNew func(dflSize T.Gsize) *T.GString

	StringFree func(str *T.GString,
		freeSegment bool) string

	StringEqual func(v *T.GString,
		v2 *T.GString) bool

	StringHash func(str *T.GString) uint

	StringAssign func(str *T.GString,
		rval string) *T.GString

	StringTruncate func(str *T.GString,
		leng T.Gsize) *T.GString

	StringSetSize func(str *T.GString,
		leng T.Gsize) *T.GString

	StringInsertLen func(str *T.GString,
		pos T.Gssize,
		val string,
		leng T.Gssize) *T.GString

	StringAppend func(str *T.GString,
		val string) *T.GString

	StringAppendLen func(str *T.GString,
		val string,
		leng T.Gssize) *T.GString

	StringAppendC func(str *T.GString,
		c T.Gchar) *T.GString

	StringAppendUnichar func(str *T.GString,
		wc T.Gunichar) *T.GString

	StringPrepend func(str *T.GString,
		val string) *T.GString

	StringPrependC func(str *T.GString,
		c T.Gchar) *T.GString

	StringPrependUnichar func(str *T.GString,
		wc T.Gunichar) *T.GString

	StringPrependLen func(str *T.GString,
		val string,
		leng T.Gssize) *T.GString

	StringInsert func(str *T.GString,
		pos T.Gssize,
		val string) *T.GString

	StringInsertC func(str *T.GString,
		pos T.Gssize,
		c T.Gchar) *T.GString

	StringInsertUnichar func(str *T.GString,
		pos T.Gssize,
		wc T.Gunichar) *T.GString

	StringOverwrite func(str *T.GString,
		pos T.Gsize,
		val string) *T.GString

	StringOverwriteLen func(str *T.GString,
		pos T.Gsize,
		val string,
		leng T.Gssize) *T.GString

	StringErase func(str *T.GString,
		pos T.Gssize,
		leng T.Gssize) *T.GString

	StringAsciiDown func(str *T.GString) *T.GString

	StringAsciiUp func(str *T.GString) *T.GString

	StringVprintf func(str *T.GString,
		format string,
		args T.VaList)

	StringPrintf func(str *T.GString, format string, v ...VArg)

	StringAppendVprintf func(str *T.GString,
		format string,
		args T.VaList)

	StringAppendPrintf func(str *T.GString,
		format string, v ...VArg)

	StringAppendUriEscaped func(str *T.GString,
		unescaped string,
		reservedChars_Allowed string,
		allowUtf8 bool) *T.GString

	StringDown func(str *T.GString) *T.GString

	StringUp func(str *T.GString) *T.GString
)

var (
	StringChunkNew func(size T.Gsize) *T.GStringChunk

	StringChunkFree func(chunk *T.GStringChunk)

	StringChunkClear func(chunk *T.GStringChunk)

	StringChunkInsert func(chunk *T.GStringChunk,
		str string) string

	StringChunkInsertLen func(chunk *T.GStringChunk,
		str string,
		leng T.Gssize) string

	StringChunkInsertConst func(
		chunk *T.GStringChunk, str string) string
)
