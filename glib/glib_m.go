// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type MainContext T.MainContext // chicken/egg

var (
	MainContextNew func() *MainContext

	MainContextDefault          func() *MainContext
	MainContextGetThreadDefault func() *MainContext

	MainCurrentSource func() *O.Source
	MainDepth         func() int

	MainContextAcquire                   func(m *MainContext) bool
	MainContextAddPoll                   func(m *MainContext, fd *T.GPollFD, priority int)
	MainContextCheck                     func(m *MainContext, maxPriority int, fds *T.GPollFD, nFds int) int
	MainContextDispatch                  func(m *MainContext)
	MainContextFindSourceByFuncsUserData func(m *MainContext, funcs *O.SourceFuncs, userData T.Gpointer) *O.Source
	MainContextFindSourceById            func(m *MainContext, sourceId uint) *O.Source
	MainContextFindSourceByUserData      func(m *MainContext, userData T.Gpointer) *O.Source
	MainContextGetPollFunc               func(m *MainContext) T.GPollFunc
	MainContextInvoke                    func(m *MainContext, function O.SourceFunc, data T.Gpointer)
	MainContextInvokeFull                func(m *MainContext, priority int, function O.SourceFunc, data T.Gpointer, notify O.DestroyNotify)
	MainContextIsOwner                   func(m *MainContext) bool
	MainContextIteration                 func(m *MainContext, mayBlock bool) bool
	MainContextPending                   func(m *MainContext) bool
	MainContextPopThreadDefault          func(m *MainContext)
	MainContextPrepare                   func(m *MainContext, priority *int) bool
	MainContextPushThreadDefault         func(m *MainContext)
	MainContextQuery                     func(m *MainContext, maxPriority int, timeout *int, fds *T.GPollFD, nFds int) int
	MainContextRef                       func(m *MainContext) *MainContext
	MainContextRelease                   func(m *MainContext)
	MainContextRemovePoll                func(m *MainContext, fd *T.GPollFD)
	MainContextSetPollFunc               func(m *MainContext, f T.GPollFunc)
	MainContextUnref                     func(m *MainContext)
	MainContextWait                      func(m *MainContext, cond *Cond, mutex *Mutex) bool
	MainContextWakeup                    func(m *MainContext)
)

func (m *MainContext) Acquire() bool                       { return MainContextAcquire(m) }
func (m *MainContext) AddPoll(fd *T.GPollFD, priority int) { MainContextAddPoll(m, fd, priority) }
func (m *MainContext) Check(maxPriority int, fds *T.GPollFD, nFds int) int {
	return MainContextCheck(m, maxPriority, fds, nFds)
}
func (m *MainContext) Dispatch() { MainContextDispatch(m) }
func (m *MainContext) FindSourceByFuncsUserData(funcs *O.SourceFuncs, userData T.Gpointer) *O.Source {
	return MainContextFindSourceByFuncsUserData(m, funcs, userData)
}
func (m *MainContext) FindSourceById(sourceId uint) *O.Source {
	return MainContextFindSourceById(m, sourceId)
}
func (m *MainContext) FindSourceByUserData(userData T.Gpointer) *O.Source {
	return MainContextFindSourceByUserData(m, userData)
}
func (m *MainContext) GetPollFunc() T.GPollFunc { return MainContextGetPollFunc(m) }
func (m *MainContext) Invoke(function O.SourceFunc, data T.Gpointer) {
	MainContextInvoke(m, function, data)
}
func (m *MainContext) InvokeFull(priority int, function O.SourceFunc, data T.Gpointer, notify O.DestroyNotify) {
	MainContextInvokeFull(m, priority, function, data, notify)
}
func (m *MainContext) IsOwner() bool                { return MainContextIsOwner(m) }
func (m *MainContext) Iteration(mayBlock bool) bool { return MainContextIteration(m, mayBlock) }
func (m *MainContext) Pending() bool                { return MainContextPending(m) }
func (m *MainContext) PopThreadDefault()            { MainContextPopThreadDefault(m) }
func (m *MainContext) Prepare(priority *int) bool   { return MainContextPrepare(m, priority) }
func (m *MainContext) PushThreadDefault()           { MainContextPushThreadDefault(m) }
func (m *MainContext) Query(maxPriority int, timeout *int, fds *T.GPollFD, nFds int) int {
	return MainContextQuery(m, maxPriority, timeout, fds, nFds)
}
func (m *MainContext) Ref() *MainContext         { return MainContextRef(m) }
func (m *MainContext) Release()                  { MainContextRelease(m) }
func (m *MainContext) RemovePoll(fd *T.GPollFD)  { MainContextRemovePoll(m, fd) }
func (m *MainContext) SetPollFunc(f T.GPollFunc) { MainContextSetPollFunc(m, f) }
func (m *MainContext) Unref()                    { MainContextUnref(m) }
func (m *MainContext) Wait(cond *Cond, mutex *Mutex) bool {
	return MainContextWait(m, cond, mutex)
}
func (m *MainContext) Wakeup() { MainContextWakeup(m) }

type MainLoop struct{}

var (
	MainLoopNew func(context *MainContext, isRunning bool) *MainLoop

	MainLoopGetContext func(m *MainLoop) *MainContext
	MainLoopIsRunning  func(m *MainLoop) bool
	MainLoopQuit       func(m *MainLoop)
	MainLoopRef        func(m *MainLoop) *MainLoop
	MainLoopRun        func(m *MainLoop)
	MainLoopUnref      func(m *MainLoop)
)

func (m *MainLoop) GetContext() *MainContext { return MainLoopGetContext(m) }
func (m *MainLoop) IsRunning() bool          { return MainLoopIsRunning(m) }
func (m *MainLoop) Quit()                    { MainLoopQuit(m) }
func (m *MainLoop) Ref() *MainLoop           { return MainLoopRef(m) }
func (m *MainLoop) Run()                     { MainLoopRun(m) }
func (m *MainLoop) Unref()                   { MainLoopUnref(m) }

type MappedFile struct{}

var (
	MappedFileNew func(filename string, writable bool, e **Error) *MappedFile

	MappedFileFree        func(m *MappedFile)
	MappedFileGetContents func(m *MappedFile) string
	MappedFileGetLength   func(m *MappedFile) T.Gsize
	MappedFileRef         func(m *MappedFile) *MappedFile
	MappedFileUnref       func(m *MappedFile)
)

func (m *MappedFile) Free()               { MappedFileFree(m) }
func (m *MappedFile) GetContents() string { return MappedFileGetContents(m) }
func (m *MappedFile) GetLength() T.Gsize  { return MappedFileGetLength(m) }
func (m *MappedFile) Ref() *MappedFile    { return MappedFileRef(m) }
func (m *MappedFile) Unref()              { MappedFileUnref(m) }

type MarkupParseContext struct{}

var (
	MarkupParseContextNew func(parser *MarkupParser, flags MarkupParseFlags, userData T.Gpointer, userDataDnotify O.DestroyNotify) *MarkupParseContext

	MarkupEscapeText     func(text string, length T.Gssize) string
	MarkupPrintfEscaped  func(format string, v ...VArg) string
	MarkupVprintfEscaped func(format string, args VAList) string

	MarkupParseContextEndParse        func(m *MarkupParseContext, e **Error) bool
	MarkupParseContextFree            func(m *MarkupParseContext)
	MarkupParseContextGetElement      func(m *MarkupParseContext) string
	MarkupParseContextGetElementStack func(m *MarkupParseContext) *SList
	MarkupParseContextGetPosition     func(m *MarkupParseContext, lineNumber, charNumber *int)
	MarkupParseContextGetUserData     func(m *MarkupParseContext) T.Gpointer
	MarkupParseContextParse           func(m *MarkupParseContext, text string, textLen T.Gssize, e **Error) bool
	MarkupParseContextPop             func(m *MarkupParseContext) T.Gpointer
	MarkupParseContextPush            func(m *MarkupParseContext, parser *MarkupParser, userData T.Gpointer)
)

func (m *MarkupParseContext) EndParse(e **Error) bool { return MarkupParseContextEndParse(m, e) }
func (m *MarkupParseContext) Free()                   { MarkupParseContextFree(m) }
func (m *MarkupParseContext) GetElement() string      { return MarkupParseContextGetElement(m) }
func (m *MarkupParseContext) GetElementStack() *SList { return MarkupParseContextGetElementStack(m) }
func (m *MarkupParseContext) GetPosition(lineNumber, charNumber *int) {
	MarkupParseContextGetPosition(m, lineNumber, charNumber)
}
func (m *MarkupParseContext) GetUserData() T.Gpointer { return MarkupParseContextGetUserData(m) }
func (m *MarkupParseContext) Parse(text string, textLen T.Gssize, e **Error) bool {
	return MarkupParseContextParse(m, text, textLen, e)
}
func (m *MarkupParseContext) Pop() T.Gpointer { return MarkupParseContextPop(m) }
func (m *MarkupParseContext) Push(parser *MarkupParser, userData T.Gpointer) {
	MarkupParseContextPush(m, parser, userData)
}

type MarkupParseFlags Enum

const (
	MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG MarkupParseFlags = 1 << iota
	MARKUP_TREAT_CDATA_AS_TEXT
	MARKUP_PREFIX_ERROR_POSITION
)

type MarkupParser struct {
	StartElement func(
		context *MarkupParseContext,
		elementName *T.Gchar,
		attributeNames **T.Gchar,
		attributeValues **T.Gchar,
		userData T.Gpointer,
		error **Error)
	EndElement func(
		context *MarkupParseContext,
		elementName *T.Gchar,
		userData T.Gpointer,
		error **Error)
	Text func(
		context *MarkupParseContext,
		text *T.Gchar,
		textLen T.Gsize,
		userData T.Gpointer,
		error **Error)
	Passthrough func(
		context *MarkupParseContext,
		passthroughText *T.Gchar,
		textLen T.Gsize,
		userData T.Gpointer,
		error **Error)
	Error func(
		context *MarkupParseContext,
		error *Error,
		userData T.Gpointer)
}

type MemChunk struct{}

var (
	MemChunkNew  func(name string, atomSize int, areaSize T.Gsize, typ int) *MemChunk
	MemChunkInfo func()

	MemChunkAlloc   func(m *MemChunk) T.Gpointer
	MemChunkAlloc0  func(m *MemChunk) T.Gpointer
	MemChunkClean   func(m *MemChunk)
	MemChunkDestroy func(m *MemChunk)
	MemChunkFree    func(m *MemChunk, mem T.Gpointer)
	MemChunkPrint   func(m *MemChunk)
	MemChunkReset   func(m *MemChunk)
)

func (m *MemChunk) Alloc() T.Gpointer   { return MemChunkAlloc(m) }
func (m *MemChunk) Alloc0() T.Gpointer  { return MemChunkAlloc0(m) }
func (m *MemChunk) Clean()              { MemChunkClean(m) }
func (m *MemChunk) Destroy()            { MemChunkDestroy(m) }
func (m *MemChunk) Free(mem T.Gpointer) { MemChunkFree(m, mem) }
func (m *MemChunk) Print()              { MemChunkPrint(m) }
func (m *MemChunk) Reset()              { MemChunkReset(m) }

type Mutex struct{}
