// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type (
	TestCase struct{}

	TestSuite struct{}

	TestFunc func()

	TestFixtureFunc func(
		fixture T.Gpointer, userData T.Gconstpointer)

	TestDataFunc func(userData T.Gconstpointer)

	TestLogFatalFunc func(
		logDomain string, logLevel T.GLogLevelFlags,
		message string, userData T.Gpointer) T.Gboolean
)

var (
	TestMinimizedResult func(minimizedQuantity float64,
		format string, v ...VArg)

	TestMaximizedResult func(maximizedQuantity float64,
		format string, v ...VArg)

	TestInit func(argc *int, argv ***T.Char, v ...VArg)

	TestRun func() int

	TestAddFunc func(testpath string,
		testFunc TestFunc)

	TestAddDataFunc func(testpath string,
		testData T.Gconstpointer,
		testFunc TestDataFunc)

	TestMessage func(format string, v ...VArg)

	TestBugBase func(uriPattern string)

	TestBug func(bugUriSnippet string)

	TestTimerStart func()

	TestTimerElapsed func() float64

	TestTimerLast func() float64

	TestQueueFree func(gfreePointer T.Gpointer)

	TestQueueDestroy func(destroyFunc O.DestroyNotify,
		destroyData T.Gpointer)

	TestTrapFork func(usecTimeout uint64,
		testTrapFlags TestTrapFlags) bool

	TestTrapHasPassed func() bool

	TestTrapReachedTimeout func() bool

	TestRandInt func() T.GInt32

	TestRandIntRange func(begin T.GInt32,
		end T.GInt32) T.GInt32

	TestRandDouble func() float64

	TestRandDoubleRange func(rangeStart float64,
		rangeEnd float64) float64

	TestCreateCase func(testName string,
		dataSize T.Gsize,
		testData T.Gconstpointer,
		dataSetup TestFixtureFunc,
		dataTest TestFixtureFunc,
		dataTeardown TestFixtureFunc) *TestCase

	TestCreateSuite func(suiteName string) *TestSuite

	TestGetRoot func() *TestSuite

	TestSuiteAdd func(suite *TestSuite,
		testCase *TestCase)

	TestSuiteAddSuite func(suite *TestSuite,
		nestedsuite *TestSuite)

	TestRunSuite func(suite *TestSuite) int

	TestTrapAssertions func(domain string,
		file string,
		line int,
		f string,
		assertionFlags uint64,
		pattern string)

	TestAddVtable func(testpath string,
		dataSize T.Gsize,
		testData T.Gconstpointer,
		dataSetup TestFixtureFunc,
		dataTest TestFixtureFunc,
		dataTeardown TestFixtureFunc)
)

type (
	TestLogBuffer struct {
		Data *T.GString
		Msgs *SList
	}

	TestLogMsg struct {
		LogType  TestLogType
		NStrings uint
		Strings  **T.Gchar
		NNums    uint
		Nums     *T.LongDouble
	}
)

type TestLogType Enum

const (
	TEST_LOG_NONE TestLogType = iota
	TEST_LOG_ERROR
	TEST_LOG_START_BINARY
	TEST_LOG_LIST_CASE
	TEST_LOG_SKIP_CASE
	TEST_LOG_START_CASE
	TEST_LOG_STOP_CASE
	TEST_LOG_MIN_RESULT
	TEST_LOG_MAX_RESULT
	TEST_LOG_MESSAGE
)

var (
	TestLogTypeName func(logType TestLogType) string

	TestLogBufferNew func() *TestLogBuffer

	TestLogBufferFree func(tbuffer *TestLogBuffer)

	TestLogBufferPush func(tbuffer *TestLogBuffer,
		nBytes uint,
		bytes *uint8)

	TestLogBufferPop func(tbuffer *TestLogBuffer) *TestLogMsg

	TestLogMsgFree func(tmsg *TestLogMsg)

	TestLogSetFatalHandler func(logFunc TestLogFatalFunc,
		userData T.Gpointer)
)

type TestTrapFlags Enum

const (
	TEST_TRAP_SILENCE_STDOUT TestTrapFlags = 1 << (7 + iota)
	TEST_TRAP_SILENCE_STDERR
	TEST_TRAP_INHERIT_STDIN
)

type Thread struct {
	Func     ThreadFunc
	Data     T.Gpointer
	Joinable T.Gboolean
	Priority ThreadPriority
}

var (
	ThreadCreateFull                func(f ThreadFunc, data T.Gpointer, stackSize T.Gulong, joinable bool, bound bool, priority ThreadPriority, e **Error) *Thread
	ThreadErrorQuark                func() Quark
	ThreadExit                      func(retval T.Gpointer)
	ThreadForeach                   func(threadFunc T.GFunc, userData T.Gpointer)
	ThreadGetInitialized            func() bool
	ThreadInit                      func(vtable *ThreadFunctions)
	ThreadInitWithErrorcheckMutexes func(vtable *ThreadFunctions)
	ThreadSelf                      func() *Thread

	ThreadJoin        func(t *Thread) T.Gpointer
	ThreadSetPriority func(t *Thread, priority ThreadPriority)
)

func (t *Thread) Join() T.Gpointer                    { return ThreadJoin(t) }
func (t *Thread) SetPriority(priority ThreadPriority) { ThreadSetPriority(t, priority) }

type (
	ThreadFunc func(data T.Gpointer) T.Gpointer

	ThreadFunctions struct {
		MutexNew      func()
		MutexLock     func(mutex *Mutex) *Mutex
		MutexTrylock  func(mutex *Mutex) bool
		MutexUnlock   func(mutex *Mutex)
		MutexFree     func(mutex *Mutex)
		CondNew       func() *Cond
		CondSignal    func(cond *Cond)
		CondBroadcast func(cond *Cond)
		CondWait      func(cond *Cond, mutex *Mutex)
		CondTimedWait func(
			cond *Cond,
			mutex *Mutex,
			endTime *TimeVal) bool
		CondFree     func(cond *Cond)
		PrivateNew   func(destructor O.DestroyNotify) *Private
		PrivateGet   func(privateKey *Private) T.Gpointer
		PrivateSet   func(privateKey *Private, data T.Gpointer)
		ThreadCreate func(
			fnc ThreadFunc,
			data T.Gpointer,
			stackSize T.Gulong,
			joinable, bound bool,
			priority ThreadPriority,
			thread T.Gpointer,
			error **Error)
		Threadield        func()
		ThreadJoin        func(thread T.Gpointer)
		ThreadExit        func()
		ThreadSetPriority func(
			thread T.Gpointer,
			priority ThreadPriority)
		ThreadSelf  func(thread T.Gpointer)
		ThreadEqual func(
			thread1 T.Gpointer,
			thread2 T.Gpointer) T.Gboolean
	}

	Private struct{}
)

type ThreadPool struct {
	Func      T.GFunc
	UserData  T.Gpointer
	Exclusive T.Gboolean
}

var (
	ThreadPoolNew func(f T.GFunc, userData T.Gpointer, maxThreads int, exclusive bool, e **Error) *ThreadPool

	ThreadPoolGetMaxIdleTime      func() uint
	ThreadPoolGetMaxUnusedThreads func() int
	ThreadPoolGetNumUnusedThreads func() uint
	ThreadPoolSetMaxIdleTime      func(interval uint)
	ThreadPoolSetMaxUnusedThreads func(maxThreads int)
	ThreadPoolStopUnusedThreads   func()

	ThreadPoolFree            func(t *ThreadPool, immediate, wait bool)
	ThreadPoolGetMaxThreads   func(t *ThreadPool) int
	ThreadPoolGetNumThreads   func(t *ThreadPool) uint
	ThreadPoolPush            func(t *ThreadPool, data T.Gpointer, e **Error)
	ThreadPoolSetMaxThreads   func(t *ThreadPool, maxThreads int, e **Error)
	ThreadPoolSetSortFunction func(t *ThreadPool, f T.GCompareDataFunc, userData T.Gpointer)
	ThreadPoolUnprocessed     func(t *ThreadPool) uint
)

func (t *ThreadPool) Free(immediate, wait bool)       { ThreadPoolFree(t, immediate, wait) }
func (t *ThreadPool) GetMaxThreads() int              { return ThreadPoolGetMaxThreads(t) }
func (t *ThreadPool) GetNumThreads() uint             { return ThreadPoolGetNumThreads(t) }
func (t *ThreadPool) Push(data T.Gpointer, e **Error) { ThreadPoolPush(t, data, e) }
func (t *ThreadPool) SetMaxThreads(maxThreads int, e **Error) {
	ThreadPoolSetMaxThreads(t, maxThreads, e)
}
func (t *ThreadPool) SetSortFunction(f T.GCompareDataFunc, userData T.Gpointer) {
	ThreadPoolSetSortFunction(t, f, userData)
}
func (t *ThreadPool) Unprocessed() uint { return ThreadPoolUnprocessed(t) }

type ThreadPriority Enum

const (
	THREAD_PRIORITY_LOW ThreadPriority = iota
	THREAD_PRIORITY_NORMAL
	THREAD_PRIORITY_HIGH
	THREAD_PRIORITY_URGENT
)

type Timer struct{}

var (
	TimerNew func() *Timer

	TimerContinue func(t *Timer)
	TimerDestroy  func(t *Timer)
	TimerElapsed  func(t *Timer, microseconds *T.Gulong) float64
	TimerReset    func(t *Timer)
	TimerStart    func(t *Timer)
	TimerStop     func(t *Timer)
)

func (t *Timer) Continue()                              { TimerContinue(t) }
func (t *Timer) Destroy()                               { TimerDestroy(t) }
func (t *Timer) Elapsed(microseconds *T.Gulong) float64 { return TimerElapsed(t, microseconds) }
func (t *Timer) Reset()                                 { TimerReset(t) }
func (t *Timer) Start()                                 { TimerStart(t) }
func (t *Timer) Stop()                                  { TimerStop(t) }

type TimeType Enum

const (
	TIME_TYPE_STANDARD TimeType = iota
	TIME_TYPE_DAYLIGHT
	TIME_TYPE_UNIVERSAL
)

type TimeVal struct {
	TvSec  T.Glong
	TvUsec T.Glong
}

var (
	TimeValFromIso8601 func(isoDate string, time *TimeVal) bool

	TimeValAdd       func(t *TimeVal, microseconds T.Glong)
	TimeValToIso8601 func(t *TimeVal) string
)

func (t *TimeVal) Add(microseconds T.Glong) { TimeValAdd(t, microseconds) }
func (t *TimeVal) ToIso8601() string        { return TimeValToIso8601(t) }

type TimeZone struct{}

var (
	TimeZoneNew      func(identifier string) *TimeZone
	TimeZoneNewLocal func() *TimeZone
	TimeZoneNewUtc   func() *TimeZone

	TimeZoneAdjustTime      func(t *TimeZone, typ TimeType, time *int64) int
	TimeZoneFindInterval    func(t *TimeZone, typ TimeType, time int64) int
	TimeZoneGetAbbreviation func(t *TimeZone, interval int) string
	TimeZoneGetOffset       func(t *TimeZone, interval int) T.GInt32
	TimeZoneIsDst           func(t *TimeZone, interval int) bool
	TimeZoneRef             func(t *TimeZone) *TimeZone
	TimeZoneUnref           func(t *TimeZone)
)

func (t *TimeZone) AdjustTime(typ TimeType, time *int64) int { return TimeZoneAdjustTime(t, typ, time) }
func (t *TimeZone) FindInterval(typ TimeType, time int64) int {
	return TimeZoneFindInterval(t, typ, time)
}
func (t *TimeZone) GetAbbreviation(interval int) string { return TimeZoneGetAbbreviation(t, interval) }
func (t *TimeZone) GetOffset(interval int) T.GInt32     { return TimeZoneGetOffset(t, interval) }
func (t *TimeZone) IsDst(interval int) bool             { return TimeZoneIsDst(t, interval) }
func (t *TimeZone) Ref() *TimeZone                      { return TimeZoneRef(t) }
func (t *TimeZone) Unref()                              { TimeZoneUnref(t) }

type TrashStack struct {
	Next *TrashStack
}

var (
	TrashStackHeight func(t **TrashStack) uint
	TrashStackPeek   func(t **TrashStack) T.Gpointer
	TrashStackPop    func(t **TrashStack) T.Gpointer
	TrashStackPush   func(t **TrashStack, dataP T.Gpointer)
)

type TraverseFlags Enum

const (
	TRAVERSE_LEAVES TraverseFlags = 1 << iota
	TRAVERSE_NON_LEAVES
	TRAVERSE_ALL       = TRAVERSE_LEAVES | TRAVERSE_NON_LEAVES
	TRAVERSE_MASK      = TRAVERSE_ALL
	TRAVERSE_LEAFS     = TRAVERSE_LEAVES
	TRAVERSE_NON_LEAFS = TRAVERSE_NON_LEAVES
)

type TraverseFunc func(key, value, data T.Gpointer) bool

type TraverseType Enum

const (
	IN_ORDER TraverseType = iota
	PRE_ORDER
	POST_ORDER
	LEVEL_ORDER
)

type Tree struct{}

var (
	TreeNew         func(keyCompareFunc T.GCompareFunc) *Tree
	TreeNewFull     func(keyCompareFunc T.GCompareDataFunc, keyCompareData T.Gpointer, keyDestroyFunc O.DestroyNotify, valueDestroyFunc O.DestroyNotify) *Tree
	TreeNewWithData func(keyCompareFunc T.GCompareDataFunc, keyCompareData T.Gpointer) *Tree

	TreeDestroy        func(t *Tree)
	TreeForeach        func(t *Tree, f TraverseFunc, userData T.Gpointer)
	TreeHeight         func(t *Tree) int
	TreeInsert         func(t *Tree, key, value T.Gpointer)
	TreeLookup         func(t *Tree, key T.Gconstpointer) T.Gpointer
	TreeLookupExtended func(t *Tree, lookupKey T.Gconstpointer, origKey, value *T.Gpointer) bool
	TreeNnodes         func(t *Tree) int
	TreeRef            func(t *Tree) *Tree
	TreeRemove         func(t *Tree, key T.Gconstpointer) bool
	TreeReplace        func(t *Tree, key, value T.Gpointer)
	TreeSearch         func(t *Tree, searchFunc T.GCompareFunc, userData T.Gconstpointer) T.Gpointer
	TreeSteal          func(t *Tree, key T.Gconstpointer) bool
	TreeTraverse       func(t *Tree, traverseFunc TraverseFunc, traverseType TraverseType, userData T.Gpointer)
	TreeUnref          func(t *Tree)
)

func (t *Tree) Destroy()                                    { TreeDestroy(t) }
func (t *Tree) Foreach(f TraverseFunc, userData T.Gpointer) { TreeForeach(t, f, userData) }
func (t *Tree) Height() int                                 { return TreeHeight(t) }
func (t *Tree) Insert(key, value T.Gpointer)                { TreeInsert(t, key, value) }
func (t *Tree) Lookup(key T.Gconstpointer) T.Gpointer       { return TreeLookup(t, key) }
func (t *Tree) LookupExtended(lookupKey T.Gconstpointer, origKey, value *T.Gpointer) bool {
	return TreeLookupExtended(t, lookupKey, origKey, value)
}
func (t *Tree) Nnodes() int                     { return TreeNnodes(t) }
func (t *Tree) Ref() *Tree                      { return TreeRef(t) }
func (t *Tree) Remove(key T.Gconstpointer) bool { return TreeRemove(t, key) }
func (t *Tree) Replace(key, value T.Gpointer)   { TreeReplace(t, key, value) }
func (t *Tree) Search(searchFunc T.GCompareFunc, userData T.Gconstpointer) T.Gpointer {
	return TreeSearch(t, searchFunc, userData)
}
func (t *Tree) Steal(key T.Gconstpointer) bool { return TreeSteal(t, key) }
func (t *Tree) Traverse(traverseFunc TraverseFunc, traverseType TraverseType, userData T.Gpointer) {
	TreeTraverse(t, traverseFunc, traverseType, userData)
}
func (t *Tree) Unref() { TreeUnref(t) }

type Tuples struct {
	Len uint
}

var (
	TuplesDestroy func(t *Tuples)
	TuplesIndex   func(t *Tuples, index, field int) T.Gpointer
)

func (t *Tuples) Destroy()                          { TuplesDestroy(t) }
func (t *Tuples) Index(index, field int) T.Gpointer { return TuplesIndex(t, index, field) }
