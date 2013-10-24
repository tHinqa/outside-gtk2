// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package glib provides API definitions for accessing
//libglib-2.0-0.dll and libgthread-2.0-0.dll.
package glib

import (
	"github.com/tHinqa/outside"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllApis(dllThread, false, apiListThread)
	outside.AddDllData(dll, false, dataList)
}

type (
	//TODO(t):Fix (stat/stat32)
	GStatBuf struct{}
)

var (
	ArrayNew func(zeroTerminated T.Gboolean,
		clear T.Gboolean,
		elementSize uint) *T.GArray

	ArraySizedNew func(zeroTerminated T.Gboolean,
		clear T.Gboolean,
		elementSize uint,
		reservedSize uint) *T.GArray

	ArrayFree func(array *T.GArray,
		freeSegment T.Gboolean) string

	ArrayRef func(array *T.GArray) *T.GArray

	ArrayUnref func(array *T.GArray)

	ArrayGetElementSize func(array *T.GArray) uint

	ArrayAppendVals func(array *T.GArray,
		data T.Gconstpointer,
		leng uint) *T.GArray

	ArrayPrependVals func(array *T.GArray,
		data T.Gconstpointer,
		leng uint) *T.GArray

	ArrayInsertVals func(array *T.GArray,
		index uint,
		data T.Gconstpointer,
		leng uint) *T.GArray

	ArraySetSize func(array *T.GArray,
		length uint) *T.GArray

	ArrayRemoveIndex func(array *T.GArray,
		index uint) *T.GArray

	ArrayRemoveIndexFast func(array *T.GArray,
		index uint) *T.GArray

	ArrayRemoveRange func(array *T.GArray,
		index uint,
		length uint) *T.GArray

	ArraySort func(array *T.GArray,
		compareFunc T.GCompareFunc)

	ArraySortWithData func(array *T.GArray,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer)

	PtrArrayNew func() *T.GPtrArray

	PtrArrayNewWithFreeFunc func(elementFreeFunc T.GDestroyNotify) *T.GPtrArray

	PtrArraySizedNew func(reservedSize uint) *T.GPtrArray

	PtrArrayFree func(array *T.GPtrArray,
		freeSeg T.Gboolean) *T.Gpointer

	PtrArrayRef func(array *T.GPtrArray) *T.GPtrArray

	PtrArrayUnref func(array *T.GPtrArray)

	PtrArraySetFreeFunc func(array *T.GPtrArray,
		elementFreeFunc T.GDestroyNotify)

	PtrArraySetSize func(array *T.GPtrArray,
		length int)

	PtrArrayRemoveIndex func(array *T.GPtrArray,
		index uint) T.Gpointer

	PtrArrayRemoveIndexFast func(array *T.GPtrArray,
		index uint) T.Gpointer

	PtrArrayRemove func(array *T.GPtrArray,
		data T.Gpointer) T.Gboolean

	PtrArrayRemoveFast func(array *T.GPtrArray,
		data T.Gpointer) T.Gboolean

	PtrArrayRemoveRange func(array *T.GPtrArray,
		index uint,
		length uint)

	PtrArrayAdd func(array *T.GPtrArray, data T.Gpointer)

	PtrArraySort func(array *T.GPtrArray,
		compareFunc T.GCompareFunc)

	PtrArraySortWithData func(array *T.GPtrArray,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer)

	PtrArrayForeach func(array *T.GPtrArray,
		f T.GFunc, userData T.Gpointer)

	ByteArrayNew func() *T.GByteArray

	ByteArraySizedNew func(reservedSize uint) *T.GByteArray

	ByteArrayFree func(array *T.GByteArray,
		freeSegment T.Gboolean) *uint8

	ByteArrayRef func(array *T.GByteArray) *T.GByteArray

	ByteArrayUnref func(array *T.GByteArray)

	ByteArrayAppend func(array *T.GByteArray,
		data *uint8, leng uint) *T.GByteArray

	ByteArrayPrepend func(array *T.GByteArray,
		data *uint8, leng uint) *T.GByteArray

	ByteArraySetSize func(array *T.GByteArray,
		length uint) *T.GByteArray

	ByteArrayRemoveIndex func(array *T.GByteArray,
		index uint) *T.GByteArray

	ByteArrayRemoveIndexFast func(array *T.GByteArray,
		index uint) *T.GByteArray

	ByteArrayRemoveRange func(array *T.GByteArray,
		index uint,
		length uint) *T.GByteArray

	ByteArraySort func(array *T.GByteArray,
		compareFunc T.GCompareFunc)

	ByteArraySortWithData func(array *T.GByteArray,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer)

	QuarkTryString func(str string) T.GQuark

	QuarkFromStaticString func(str string) T.GQuark

	QuarkFromString func(str string) T.GQuark

	QuarkToString func(quark T.GQuark) string

	InternString func(str string) string

	InternStaticString func(str string) string

	ErrorNew func(domain T.GQuark, code int, format string,
		v ...VArg) *T.GError

	ErrorNewLiteral func(domain T.GQuark,
		code int,
		message string) *T.GError

	ErrorNewValist func(domain T.GQuark,
		code int,
		format string,
		args T.VaList) *T.GError

	ErrorFree func(e *T.GError)

	ErrorCopy func(e *T.GError) *T.GError

	ErrorMatches func(e *T.GError,
		domain T.GQuark,
		code int) T.Gboolean

	SetError func(err **T.GError, domain T.GQuark, code int,
		format string, v ...VArg)

	SetErrorLiteral func(err **T.GError,
		domain T.GQuark,
		code int,
		message string)

	PropagateError func(dest **T.GError,
		src *T.GError)

	ClearError func(err **T.GError)

	PrefixError func(err **T.GError, format string, v ...VArg)

	PropagatePrefixedError func(dest **T.GError, src *T.GError,
		format string, v ...VArg)

	GetUserName func() string

	GetRealName func() string

	GetHomeDir func() string

	GetTmpDir func() string

	GetHostName func() string

	GetPrgname func() string

	SetPrgname func(prgname string)

	GetApplicationName func() string

	SetApplicationName func(applicationName string)

	ReloadUserSpecialDirsCache func()

	GetUserDataDir func() string

	GetUserConfigDir func() string

	GetUserCacheDir func() string

	GetSystemDataDirs func() **T.Gchar

	Win32GetSystemDataDirsForModule func(
		f func()) **T.Gchar

	ParseDebugString func(
		str string, keys *T.GDebugKey, nkeys uint) uint

	Snprintf func(str string, n T.Gulong, format string,
		v ...VArg) int

	Vsnprintf func(str string,
		n T.Gulong,
		format string,
		args T.VaList) int

	PathIsAbsolute func(fileName string) T.Gboolean

	PathSkipRoot func(fileName string) string

	Basename func(fileName string) string

	GetCurrentDir func() string

	PathGetBasename func(fileName string) string

	PathGetDirname func(fileName string) string

	NullifyPointer func(nullifyLocation *T.Gpointer)

	Getenv func(variable string) string

	Setenv func(variable string,
		value string,
		overwrite T.Gboolean) T.Gboolean

	Unsetenv func(variable string)

	Listenv func() **T.Gchar

	GetEnviron func() **T.Gchar

	ThreadInit func(vtable *T.GThreadFunctions)

	ThreadInitWithErrorcheckMutexes func(vtable *T.GThreadFunctions)

	ThreadGetInitialized func() T.Gboolean

	StaticMutexGetMutexImpl func(mutex **T.GMutex) *T.GMutex

	ThreadCreateFull func(f T.GThreadFunc,
		data T.Gpointer,
		stackSize T.Gulong,
		joinable T.Gboolean,
		bound T.Gboolean,
		priority T.GThreadPriority,
		e **T.GError) *T.GThread

	ThreadSelf func() *T.GThread

	ThreadExit func(retval T.Gpointer)

	ThreadJoin func(thread *T.GThread) T.Gpointer

	ThreadSetPriority func(thread *T.GThread,
		priority T.GThreadPriority)

	StaticMutexInit func(mutex *T.GStaticMutex)

	StaticMutexFree func(mutex *T.GStaticMutex)

	StaticRecMutexInit func(mutex *T.GStaticRecMutex)

	StaticRecMutexLock func(mutex *T.GStaticRecMutex)

	StaticRecMutexTrylock func(mutex *T.GStaticRecMutex) T.Gboolean

	StaticRecMutexUnlock func(mutex *T.GStaticRecMutex)

	StaticRecMutexLockFull func(mutex *T.GStaticRecMutex,
		depth uint)

	StaticRecMutexUnlockFull func(mutex *T.GStaticRecMutex) uint

	StaticRecMutexFree func(mutex *T.GStaticRecMutex)

	StaticRwLockInit func(
		lock *T.GStaticRWLock)

	StaticRwLockReaderLock func(
		lock *T.GStaticRWLock)

	StaticRwLockReaderTrylock func(
		lock *T.GStaticRWLock) T.Gboolean

	StaticRwLockReaderUnlock func(
		lock *T.GStaticRWLock)

	StaticRwLockWriterLock func(
		lock *T.GStaticRWLock)

	StaticRwLockWriterTrylock func(
		lock *T.GStaticRWLock) T.Gboolean

	StaticRwLockWriterUnlock func(
		lock *T.GStaticRWLock)

	StaticRwLockFree func(
		lock *T.GStaticRWLock)

	ThreadForeach func(threadFunc T.GFunc,
		userData T.Gpointer)

	OnceImpl func(
		once *T.GOnce,
		f T.GThreadFunc,
		arg T.Gpointer) T.Gpointer

	OnceInitEnter func(
		valueLocation *T.Gsize) T.Gboolean

	OnceInitEnterImpl func(
		valueLocation *T.Gsize) T.Gboolean

	OnceInitLeave func(
		valueLocation *T.Gsize,
		initializationValue T.Gsize)

	AsyncQueueNew func() *T.GAsyncQueue

	AsyncQueueNewFull func(itemFreeFunc T.GDestroyNotify) *T.GAsyncQueue

	AsyncQueueLock func(queue *T.GAsyncQueue)

	AsyncQueueUnlock func(queue *T.GAsyncQueue)

	AsyncQueueRef func(queue *T.GAsyncQueue) *T.GAsyncQueue

	AsyncQueueUnref func(queue *T.GAsyncQueue)

	AsyncQueueRefUnlocked func(queue *T.GAsyncQueue)

	AsyncQueueUnrefAndUnlock func(queue *T.GAsyncQueue)

	AsyncQueuePush func(queue *T.GAsyncQueue,
		data T.Gpointer)

	AsyncQueuePushUnlocked func(queue *T.GAsyncQueue,
		data T.Gpointer)

	AsyncQueuePushSorted func(queue *T.GAsyncQueue,
		data T.Gpointer,
		f T.GCompareDataFunc,
		userData T.Gpointer)

	AsyncQueuePushSortedUnlocked func(queue *T.GAsyncQueue,
		data T.Gpointer,
		f T.GCompareDataFunc,
		userData T.Gpointer)

	AsyncQueuePop func(queue *T.GAsyncQueue) T.Gpointer

	AsyncQueuePopUnlocked func(queue *T.GAsyncQueue) T.Gpointer

	AsyncQueueTryPop func(queue *T.GAsyncQueue) T.Gpointer

	AsyncQueueTryPopUnlocked func(queue *T.GAsyncQueue) T.Gpointer

	AsyncQueueTimedPop func(queue *T.GAsyncQueue,
		endTime *T.GTimeVal) T.Gpointer

	AsyncQueueTimedPopUnlocked func(queue *T.GAsyncQueue,
		endTime *T.GTimeVal) T.Gpointer

	AsyncQueueLength func(queue *T.GAsyncQueue) int

	AsyncQueueLengthUnlocked func(queue *T.GAsyncQueue) int

	AsyncQueueSort func(queue *T.GAsyncQueue,
		f T.GCompareDataFunc,
		userData T.Gpointer)

	AsyncQueueSortUnlocked func(queue *T.GAsyncQueue,
		f T.GCompareDataFunc,
		userData T.Gpointer)

	OnErrorQuery func(prgName string)

	OnErrorStackTrace func(prgName string)

	Base64EncodeStep func(in *T.Guchar,
		leng T.Gsize,
		breakLines T.Gboolean,
		out string,
		state *int,
		save *int) T.Gsize

	Base64EncodeClose func(breakLines T.Gboolean,
		out string,
		state *int,
		save *int) T.Gsize

	Base64Encode func(data *T.Guchar,
		leng T.Gsize) string

	Base64DecodeStep func(in string,
		leng T.Gsize,
		out *T.Guchar,
		state *int,
		save *uint) T.Gsize

	Base64Decode func(text string,
		outLen *T.Gsize) *T.Guchar

	Base64DecodeInplace func(text string,
		outLen *T.Gsize) *T.Guchar

	BitLock func(
		address *int,
		lockBit int)

	BitTrylock func(
		address *int,
		lockBit int) T.Gboolean

	BitUnlock func(
		address *int,
		lockBit int)

	GetSystemConfigDirs func() **T.Gchar

	GetUserRuntimeDir func() string

	GetLanguageNames func() **T.Gchar

	GetLocaleVariants func(locale string) **T.Gchar

	Atexit func(f T.GVoidFunc)

	GlibCheckVersion func(requiredMajor uint,
		requiredMinor uint,
		requiredMicro uint) string

	AtomicIntExchangeAndAdd func(atomic *int,
		val int) int

	AtomicIntAdd func(atomic *int,
		val int)

	AtomicIntCompareAndExchange func(atomic *int,
		oldval int,
		newval int) T.Gboolean

	AtomicPointerCompareAndExchange func(atomic *T.Gpointer,
		oldval T.Gpointer,
		newval T.Gpointer) T.Gboolean

	AtomicIntGet func(atomic *int) int

	AtomicIntSet func(atomic *int,
		newval int)

	AtomicPointerGet func(atomic *T.Gpointer) T.Gpointer

	AtomicPointerSet func(atomic *T.Gpointer,
		newval T.Gpointer)

	ThreadErrorQuark func() T.GQuark

	BookmarkFileErrorQuark func() T.GQuark

	BookmarkFileNew func() *T.GBookmarkFile

	BookmarkFileFree func(bookmark *T.GBookmarkFile)

	BookmarkFileLoadFromFile func(bookmark *T.GBookmarkFile,
		filename string,
		e **T.GError) T.Gboolean

	BookmarkFileLoadFromData func(bookmark *T.GBookmarkFile,
		data string,
		length T.Gsize,
		e **T.GError) T.Gboolean

	BookmarkFileLoadFromDataDirs func(bookmark *T.GBookmarkFile,
		file string,
		fullPath **T.Gchar,
		e **T.GError) T.Gboolean

	BookmarkFileToData func(bookmark *T.GBookmarkFile,
		length *T.Gsize,
		e **T.GError) string

	BookmarkFileToFile func(bookmark *T.GBookmarkFile,
		filename string,
		e **T.GError) T.Gboolean

	BookmarkFileSetTitle func(bookmark *T.GBookmarkFile,
		uri string,
		title string)

	BookmarkFileGetTitle func(bookmark *T.GBookmarkFile,
		uri string,
		e **T.GError) string

	BookmarkFileSetDescription func(bookmark *T.GBookmarkFile,
		uri string,
		description string)

	BookmarkFileGetDescription func(bookmark *T.GBookmarkFile,
		uri string,
		e **T.GError) string

	BookmarkFileSetMimeType func(bookmark *T.GBookmarkFile,
		uri string,
		mimeType string)

	BookmarkFileGetMimeType func(bookmark *T.GBookmarkFile,
		uri string,
		e **T.GError) string

	BookmarkFileSetGroups func(bookmark *T.GBookmarkFile,
		uri string,
		groups **T.Gchar,
		length T.Gsize)

	BookmarkFileAddGroup func(bookmark *T.GBookmarkFile,
		uri string,
		group string)

	BookmarkFileHasGroup func(bookmark *T.GBookmarkFile,
		uri string,
		group string,
		e **T.GError) T.Gboolean

	BookmarkFileGetGroups func(bookmark *T.GBookmarkFile,
		uri string,
		length *T.Gsize,
		e **T.GError) **T.Gchar

	BookmarkFileAddApplication func(bookmark *T.GBookmarkFile,
		uri string,
		name string,
		exec string)

	BookmarkFileHasApplication func(bookmark *T.GBookmarkFile,
		uri string,
		name string,
		e **T.GError) T.Gboolean

	BookmarkFileGetApplications func(bookmark *T.GBookmarkFile,
		uri string,
		length *T.Gsize,
		e **T.GError) **T.Gchar

	BookmarkFileSetAppInfo func(bookmark *T.GBookmarkFile,
		uri string,
		name string,
		exec string,
		count int,
		stamp T.TimeT,
		e **T.GError) T.Gboolean

	BookmarkFileGetAppInfo func(bookmark *T.GBookmarkFile,
		uri string,
		name string,
		exec **T.Gchar,
		count *uint,
		stamp *T.TimeT,
		e **T.GError) T.Gboolean

	BookmarkFileSetIsPrivate func(bookmark *T.GBookmarkFile,
		uri string,
		isPrivate T.Gboolean)

	BookmarkFileGetIsPrivate func(bookmark *T.GBookmarkFile,
		uri string,
		e **T.GError) T.Gboolean

	BookmarkFileSetIcon func(bookmark *T.GBookmarkFile,
		uri string,
		href string,
		mimeType string)

	BookmarkFileGetIcon func(bookmark *T.GBookmarkFile,
		uri string,
		href **T.Gchar,
		mimeType **T.Gchar,
		e **T.GError) T.Gboolean

	BookmarkFileSetAdded func(bookmark *T.GBookmarkFile,
		uri string,
		added T.TimeT)

	BookmarkFileGetAdded func(bookmark *T.GBookmarkFile,
		uri string,
		e **T.GError) T.TimeT

	BookmarkFileSetModified func(bookmark *T.GBookmarkFile,
		uri string,
		modified T.TimeT)

	BookmarkFileGetModified func(bookmark *T.GBookmarkFile,
		uri string,
		e **T.GError) T.TimeT

	BookmarkFileSetVisited func(bookmark *T.GBookmarkFile,
		uri string,
		visited T.TimeT)

	BookmarkFileGetVisited func(bookmark *T.GBookmarkFile,
		uri string,
		e **T.GError) T.TimeT

	BookmarkFileHasItem func(bookmark *T.GBookmarkFile,
		uri string) T.Gboolean

	BookmarkFileGetSize func(bookmark *T.GBookmarkFile) int

	BookmarkFileGetUris func(bookmark *T.GBookmarkFile,
		length *T.Gsize) **T.Gchar

	BookmarkFileRemoveGroup func(bookmark *T.GBookmarkFile,
		uri string,
		group string,
		e **T.GError) T.Gboolean

	BookmarkFileRemoveApplication func(bookmark *T.GBookmarkFile,
		uri string,
		name string,
		e **T.GError) T.Gboolean

	BookmarkFileRemoveItem func(bookmark *T.GBookmarkFile,
		uri string,
		e **T.GError) T.Gboolean

	BookmarkFileMoveItem func(bookmark *T.GBookmarkFile,
		oldUri string,
		newUri string,
		e **T.GError) T.Gboolean

	SliceAlloc func(blockSize T.Gsize) T.Gpointer

	SliceAlloc0 func(blockSize T.Gsize) T.Gpointer

	SliceCopy func(blockSize T.Gsize,
		memBlock T.Gconstpointer) T.Gpointer

	SliceFree1 func(blockSize T.Gsize,
		memBlock T.Gpointer)

	SliceFreeChainWithOffset func(blockSize T.Gsize,
		memChain T.Gpointer,
		nextOffset T.Gsize)

	SliceSetConfig func(ckey T.GSliceConfig, value int64)

	SliceGetConfig func(ckey T.GSliceConfig) int64

	SliceGetConfigState func(ckey T.GSliceConfig, address int64, nValues *uint) *int64

	Free func(mem T.Gpointer)

	Malloc func(nBytes T.Gsize) T.Gpointer

	Malloc0 func(nBytes T.Gsize) T.Gpointer

	Realloc func(mem T.Gpointer,
		nBytes T.Gsize) T.Gpointer

	TryMalloc func(nBytes T.Gsize) T.Gpointer

	TryMalloc0 func(nBytes T.Gsize) T.Gpointer

	TryRealloc func(mem T.Gpointer,
		nBytes T.Gsize) T.Gpointer

	MallocN func(nBlocks T.Gsize,
		nBlockBytes T.Gsize) T.Gpointer

	Malloc0N func(nBlocks T.Gsize,
		nBlockBytes T.Gsize) T.Gpointer

	ReallocN func(mem T.Gpointer,
		nBlocks T.Gsize,
		nBlockBytes T.Gsize) T.Gpointer

	TryMallocN func(nBlocks T.Gsize,
		nBlockBytes T.Gsize) T.Gpointer

	TryMalloc0N func(nBlocks T.Gsize,
		nBlockBytes T.Gsize) T.Gpointer

	TryReallocN func(mem T.Gpointer,
		nBlocks T.Gsize,
		nBlockBytes T.Gsize) T.Gpointer

	MemIsSystemMalloc func() T.Gboolean

	MemProfile func()

	MemChunkNew func(name string, atomSize int,
		areaSize T.Gsize, typ int) *T.GMemChunk

	MemChunkDestroy func(memChunk *T.GMemChunk)

	MemChunkAlloc func(memChunk *T.GMemChunk) T.Gpointer

	MemChunkAlloc0 func(memChunk *T.GMemChunk) T.Gpointer

	MemChunkFree func(memChunk *T.GMemChunk,
		mem T.Gpointer)

	MemChunkClean func(memChunk *T.GMemChunk)

	MemChunkReset func(memChunk *T.GMemChunk)

	MemChunkPrint func(memChunk *T.GMemChunk)

	MemChunkInfo func()

	BlowChunks func()

	AllocatorNew func(name string,
		nPreallocs uint) *T.GAllocator

	AllocatorFree func(allocator *T.GAllocator)

	ListAlloc func() *T.GList

	ListFree func(list *T.GList)

	ListFree1 func(list *T.GList)

	ListFreeFull func(list *T.GList,
		freeFunc T.GDestroyNotify)

	ListAppend func(list *T.GList,
		data T.Gpointer) *T.GList

	ListPrepend func(list *T.GList,
		data T.Gpointer) *T.GList

	ListInsert func(list *T.GList,
		data T.Gpointer,
		position int) *T.GList

	ListInsertSorted func(list *T.GList,
		data T.Gpointer,
		f T.GCompareFunc) *T.GList

	ListInsertSortedWithData func(list *T.GList,
		data T.Gpointer,
		fGCompareDataFunc,
		userData T.Gpointer) *T.GList

	ListInsertBefore func(list *T.GList,
		sibling *T.GList,
		data T.Gpointer) *T.GList

	ListConcat func(list1 *T.GList,
		list2 *T.GList) *T.GList

	ListRemove func(list *T.GList,
		data T.Gconstpointer) *T.GList

	ListRemoveAll func(list *T.GList,
		data T.Gconstpointer) *T.GList

	ListRemoveLink func(list *T.GList,
		llink *T.GList) *T.GList

	ListDeleteLink func(list *T.GList,
		link *T.GList) *T.GList

	ListReverse func(list *T.GList) *T.GList

	ListCopy func(list *T.GList) *T.GList

	ListNth func(list *T.GList,
		n uint) *T.GList

	ListNthPrev func(list *T.GList,
		n uint) *T.GList

	ListFind func(list *T.GList,
		data T.Gconstpointer) *T.GList

	ListFindCustom func(list *T.GList,
		data T.Gconstpointer,
		f T.GCompareFunc) *T.GList

	ListPosition func(list *T.GList,
		llink *T.GList) int

	ListIndex func(list *T.GList,
		data T.Gconstpointer) int

	ListLast func(list *T.GList) *T.GList

	ListFirst func(list *T.GList) *T.GList

	ListLength func(list *T.GList) uint

	ListForeach func(list *T.GList,
		f T.GFunc,
		userData T.Gpointer)

	ListSort func(list *T.GList,
		compareFunc T.GCompareFunc) *T.GList

	ListSortWithData func(list *T.GList,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer) *T.GList

	ListNthData func(list *T.GList,
		n uint) T.Gpointer

	ListPushAllocator func(allocator T.Gpointer)

	ListPopAllocator func()

	CacheNew func(valueNewFunc T.GCacheNewFunc,
		valueDestroyFunc T.GCacheDestroyFunc,
		keyDupFunc T.GCacheDupFunc,
		keyDestroyFunc T.GCacheDestroyFunc,
		hashKeyFunc T.GHashFunc,
		hashValueFunc T.GHashFunc,
		keyEqualFunc T.GEqualFunc) *T.GCache

	CacheDestroy func(cache *T.GCache)

	CacheInsert func(cache *T.GCache,
		key T.Gpointer) T.Gpointer

	CacheRemove func(cache *T.GCache,
		value T.Gconstpointer)

	CacheKeyForeach func(cache *T.GCache,
		f T.GHFunc, userData T.Gpointer)

	CacheValueForeach func(cache *T.GCache,
		f T.GHFunc, userData T.Gpointer)

	ChecksumTypeGetLength func(checksumType T.GChecksumType) T.Gssize

	ChecksumNew func(checksumType T.GChecksumType) *T.GChecksum

	ChecksumReset func(checksum *T.GChecksum)

	ChecksumCopy func(checksum *T.GChecksum) *T.GChecksum

	ChecksumFree func(checksum *T.GChecksum)

	ChecksumUpdate func(checksum *T.GChecksum,
		data *T.Guchar, length T.Gssize)

	ChecksumGetString func(checksum *T.GChecksum) string

	ChecksumGetDigest func(checksum *T.GChecksum,
		buffer *uint8, digestLen *T.Gsize)

	ComputeChecksumForData func(checksumType T.GChecksumType,
		data *T.Guchar, length T.Gsize) string

	ComputeChecksumForString func(checksumType T.GChecksumType,
		str string, length T.Gssize) string

	CompletionNew func(f T.GCompletionFunc) *T.GCompletion

	CompletionAddItems func(cmp *T.GCompletion, items *T.GList)

	CompletionRemoveItems func(cmp *T.GCompletion,
		items *T.GList)

	CompletionClearItems func(cmp *T.GCompletion)

	CompletionComplete func(cmp *T.GCompletion,
		prefix string,
		newPrefix **T.Gchar) *T.GList

	CompletionCompleteUtf8 func(cmp *T.GCompletion,
		prefix string,
		newPrefix **T.Gchar) *T.GList

	CompletionSetCompare func(cmp *T.GCompletion,
		strncmpFunc T.GCompletionStrncmpFunc)

	CompletionFree func(cmp *T.GCompletion)

	ConvertErrorQuark func() T.GQuark

	IconvOpen func(toCodeset string,
		fromCodeset string) T.GIConv

	Iconv func(converter T.GIConv,
		inbuf **T.Gchar,
		inbytesLeft *T.Gsize,
		outbuf **T.Gchar,
		outbytesLeft *T.Gsize) T.Gsize

	IconvClose func(converter T.GIConv) int

	Convert func(str string,
		leng T.Gssize,
		toCodeset string,
		fromCodeset string,
		bytesRead *T.Gsize,
		bytesWritten *T.Gsize,
		e **T.GError) string

	ConvertWithIconv func(str string,
		leng T.Gssize,
		converter T.GIConv,
		bytesRead *T.Gsize,
		bytesWritten *T.Gsize,
		e **T.GError) string

	ConvertWithFallback func(str string,
		leng T.Gssize,
		toCodeset string,
		fromCodeset string,
		fallback string,
		bytesRead *T.Gsize,
		bytesWritten *T.Gsize,
		e **T.GError) string

	LocaleToUtf8 func(opsysstr string,
		leng T.Gssize,
		bytesRead *T.Gsize,
		bytesWritten *T.Gsize,
		e **T.GError) string

	LocaleFromUtf8 func(utf8str string,
		leng T.Gssize,
		bytesRead *T.Gsize,
		bytesWritten *T.Gsize,
		e **T.GError) string

	FilenameToUtf8 func(opsysstr string,
		leng T.Gssize,
		bytesRead *T.Gsize,
		bytesWritten *T.Gsize,
		e **T.GError) string

	FilenameFromUtf8 func(utf8str string,
		leng T.Gssize,
		bytesRead *T.Gsize,
		bytesWritten *T.Gsize,
		e **T.GError) string

	FilenameFromUri func(uri string,
		hostname **T.Gchar,
		e **T.GError) string

	FilenameToUri func(filename string,
		hostname string,
		e **T.GError) string

	FilenameDisplayName func(filename string) string

	GetFilenameCharsets func(charsets ***T.Gchar) T.Gboolean

	FilenameDisplayBasename func(filename string) string

	UriListExtractUris func(uriList string) **T.Gchar

	DatalistInit func(datalist **T.GData)

	DatalistClear func(datalist **T.GData)

	DatalistIdGetData func(datalist **T.GData,
		keyId T.GQuark) T.Gpointer

	DatalistIdSetDataFull func(datalist **T.GData,
		keyId T.GQuark,
		data T.Gpointer,
		destroyFunc T.GDestroyNotify)

	DatalistIdRemoveNoNotify func(datalist **T.GData,
		keyId T.GQuark) T.Gpointer

	DatalistForeach func(datalist **T.GData,
		f T.GDataForeachFunc,
		userData T.Gpointer)

	DatalistSetFlags func(datalist **T.GData,
		flags uint)

	DatalistUnsetFlags func(datalist **T.GData,
		flags uint)

	DatalistGetFlags func(datalist **T.GData) uint

	DatasetDestroy func(datasetLocation T.Gconstpointer)

	DatasetIdGetData func(datasetLocation T.Gconstpointer,
		keyId T.GQuark) T.Gpointer

	DatasetIdSetDataFull func(datasetLocation T.Gconstpointer,
		keyId T.GQuark,
		data T.Gpointer,
		destroyFunc T.GDestroyNotify)

	DatasetIdRemoveNoNotify func(datasetLocation T.Gconstpointer,
		keyId T.GQuark) T.Gpointer

	DatasetForeach func(datasetLocation T.Gconstpointer,
		f T.GDataForeachFunc,
		userData T.Gpointer)

	DateNew func() *T.GDate

	DateNewDmy func(day T.GDateDay,
		month T.GDateMonth,
		year T.GDateYear) *T.GDate

	DateNewJulian func(julianDay T.GUint32) *T.GDate

	DateFree func(date *T.GDate)

	DateValid func(date *T.GDate) T.Gboolean

	DateValidDay func(day T.GDateDay) T.Gboolean

	DateValidMonth func(month T.GDateMonth) T.Gboolean

	DateValidYear func(year T.GDateYear) T.Gboolean

	DateValidWeekday func(weekday T.GDateWeekday) T.Gboolean

	DateValidJulian func(julianDate T.GUint32) T.Gboolean

	DateValidDmy func(day T.GDateDay,
		month T.GDateMonth,
		year T.GDateYear) T.Gboolean

	DateGetWeekday func(date *T.GDate) T.GDateWeekday

	DateGetMonth func(date *T.GDate) T.GDateMonth

	DateGetYear func(date *T.GDate) T.GDateYear

	DateGetDay func(date *T.GDate) T.GDateDay

	DateGetJulian func(date *T.GDate) T.GUint32

	DateGetDayOfYear func(date *T.GDate) uint

	DateGetMondayWeekOfYear func(date *T.GDate) uint

	DateGetSundayWeekOfYear func(date *T.GDate) uint

	DateGetIso8601WeekOfYear func(date *T.GDate) uint

	DateClear func(date *T.GDate, nDates uint)

	DateSetParse func(date *T.GDate, str string)

	DateSetTimeT func(date *T.GDate, timet T.TimeT)

	DateSetTimeVal func(date *T.GDate, timeval *T.GTimeVal)

	DateSetTime func(date *T.GDate, time T.GTime)

	DateSetMonth func(date *T.GDate, month T.GDateMonth)

	DateSetDay func(date *T.GDate, day T.GDateDay)

	DateSetYear func(date *T.GDate, year T.GDateYear)

	DateSetDmy func(date *T.GDate,
		day T.GDateDay, month T.GDateMonth, y T.GDateYear)

	DateSetJulian func(date *T.GDate, julianDate T.GUint32)

	DateIsFirstOfMonth func(date *T.GDate) T.Gboolean

	DateIsLastOfMonth func(date *T.GDate) T.Gboolean

	DateAddDays func(date *T.GDate, nDays uint)

	DateSubtractDays func(date *T.GDate, nDays uint)

	DateAddMonths func(date *T.GDate, nMonths uint)

	DateSubtractMonths func(date *T.GDate, nMonths uint)

	DateAddYears func(date *T.GDate, nYears uint)

	DateSubtractYears func(date *T.GDate, nYears uint)

	DateIsLeapYear func(year T.GDateYear) T.Gboolean

	DateGetDaysInMonth func(month T.GDateMonth,
		year T.GDateYear) uint8

	DateGetMondayWeeksInYear func(year T.GDateYear) uint8

	DateGetSundayWeeksInYear func(year T.GDateYear) uint8

	DateDaysBetween func(date1 *T.GDate,
		date2 *T.GDate) int

	DateCompare func(lhs *T.GDate,
		rhs *T.GDate) int

	DateToStructTm func(date *T.GDate, tm *T.Tm)

	DateClamp func(date, minDate, maxDate *T.GDate)

	DateOrder func(date1 *T.GDate, date2 *T.GDate)

	DateStrftime func(s string,
		slen T.Gsize,
		format string,
		date *T.GDate) T.Gsize

	TimeZoneNew func(identifier string) *T.GTimeZone

	TimeZoneNewUtc func() *T.GTimeZone

	TimeZoneNewLocal func() *T.GTimeZone

	TimeZoneRef func(tz *T.GTimeZone) *T.GTimeZone

	TimeZoneUnref func(tz *T.GTimeZone)

	TimeZoneFindInterval func(tz *T.GTimeZone,
		typ T.GTimeType,
		time int64) int

	TimeZoneAdjustTime func(tz *T.GTimeZone,
		typ T.GTimeType,
		time *int64) int

	TimeZoneGetAbbreviation func(tz *T.GTimeZone,
		interval int) string

	TimeZoneGetOffset func(tz *T.GTimeZone,
		interval int) T.GInt32

	TimeZoneIsDst func(tz *T.GTimeZone,
		interval int) T.Gboolean

	DateTimeUnref func(datetime *T.GDateTime)

	DateTimeRef func(datetime *T.GDateTime) *T.GDateTime

	DateTimeNewNow func(tz *T.GTimeZone) *T.GDateTime

	DateTimeNewNowLocal func() *T.GDateTime

	DateTimeNewNowUtc func() *T.GDateTime

	DateTimeNewFromUnixLocal func(t int64) *T.GDateTime

	DateTimeNewFromUnixUtc func(t int64) *T.GDateTime

	DateTimeNewFromTimevalLocal func(tv *T.GTimeVal) *T.GDateTime

	DateTimeNewFromTimevalUtc func(tv *T.GTimeVal) *T.GDateTime

	DateTimeNew func(tz *T.GTimeZone,
		year int,
		month int,
		day int,
		hour int,
		minute int,
		seconds float64) *T.GDateTime

	DateTimeNewLocal func(year int,
		month int,
		day int,
		hour int,
		minute int,
		seconds float64) *T.GDateTime

	DateTimeNewUtc func(year int,
		month int,
		day int,
		hour int,
		minute int,
		seconds float64) *T.GDateTime

	DateTimeAdd func(datetime *T.GDateTime,
		timespan T.GTimeSpan) *T.GDateTime

	DateTimeAddYears func(datetime *T.GDateTime,
		years int) *T.GDateTime

	DateTimeAddMonths func(datetime *T.GDateTime,
		months int) *T.GDateTime

	DateTimeAddWeeks func(datetime *T.GDateTime,
		weeks int) *T.GDateTime

	DateTimeAddDays func(datetime *T.GDateTime,
		days int) *T.GDateTime

	DateTimeAddHours func(datetime *T.GDateTime,
		hours int) *T.GDateTime

	DateTimeAddMinutes func(datetime *T.GDateTime,
		minutes int) *T.GDateTime

	DateTimeAddSeconds func(datetime *T.GDateTime,
		seconds float64) *T.GDateTime

	DateTimeAddFull func(datetime *T.GDateTime,
		years int,
		months int,
		days int,
		hours int,
		minutes int,
		seconds float64) *T.GDateTime

	DateTimeCompare func(dt1 T.Gconstpointer,
		dt2 T.Gconstpointer) int

	DateTimeDifference func(end *T.GDateTime,
		begin *T.GDateTime) T.GTimeSpan

	DateTimeHash func(datetime T.Gconstpointer) uint

	DateTimeEqual func(dt1 T.Gconstpointer,
		dt2 T.Gconstpointer) T.Gboolean

	DateTimeGetYmd func(datetime *T.GDateTime,
		year *int,
		month *int,
		day *int)

	DateTimeGetYear func(datetime *T.GDateTime) int

	DateTimeGetMonth func(datetime *T.GDateTime) int

	DateTimeGetDayOfMonth func(datetime *T.GDateTime) int

	DateTimeGetWeekNumberingYear func(datetime *T.GDateTime) int

	DateTimeGetWeekOfYear func(datetime *T.GDateTime) int

	DateTimeGetDayOfWeek func(datetime *T.GDateTime) int

	DateTimeGetDayOfYear func(datetime *T.GDateTime) int

	DateTimeGetHour func(datetime *T.GDateTime) int

	DateTimeGetMinute func(datetime *T.GDateTime) int

	DateTimeGetSecond func(datetime *T.GDateTime) int

	DateTimeGetMicrosecond func(datetime *T.GDateTime) int

	DateTimeGetSeconds func(datetime *T.GDateTime) float64

	DateTimeToUnix func(datetime *T.GDateTime) int64

	DateTimeToTimeval func(datetime *T.GDateTime,
		tv *T.GTimeVal) T.Gboolean

	DateTimeGetUtcOffset func(datetime *T.GDateTime) T.GTimeSpan

	DateTimeGetTimezoneAbbreviation func(datetime *T.GDateTime) string

	DateTimeIsDaylightSavings func(datetime *T.GDateTime) T.Gboolean

	DateTimeToTimezone func(datetime *T.GDateTime,
		tz *T.GTimeZone) *T.GDateTime

	DateTimeToLocal func(datetime *T.GDateTime) *T.GDateTime

	DateTimeToUtc func(datetime *T.GDateTime) *T.GDateTime

	DateTimeFormat func(datetime *T.GDateTime,
		format string) string

	DirOpenUtf8 func(path string,
		flags uint,
		e **T.GError) *T.GDir

	DirReadNameUtf8 func(dir *T.GDir) string

	DirRewind func(dir *T.GDir)

	DirClose func(dir *T.GDir)

	FileErrorQuark func() T.GQuark

	FileErrorFromErrno func(errNo int) T.GFileError

	FileTest func(filename string,
		test T.GFileTest) T.Gboolean

	FileGetContents func(filename string,
		contents **T.Gchar,
		length *T.Gsize,
		e **T.GError) T.Gboolean

	FileSetContents func(filename string,
		contents string,
		length T.Gssize,
		e **T.GError) T.Gboolean

	FileReadLink func(filename string,
		e **T.GError) string

	Mkstemp func(tmpl string) int

	MkstempFull func(tmpl string,
		flags int,
		mode int) int

	FileOpenTmp func(tmpl string,
		nameUsed **T.Gchar,
		e **T.GError) int

	FormatSizeForDisplay func(size T.Goffset) string

	BuildPath func(separator string, firstElement string,
		v ...VArg) string

	BuildPathv func(separator string,
		args **T.Gchar) string

	BuildFilename func(firstElement string, v ...VArg) string

	BuildFilenamev func(args **T.Gchar) string

	MkdirWithParents func(pathname string,
		mode int) int

	HashTableNew func(hashFunc T.GHashFunc,
		keyEqualFunc T.GEqualFunc) *T.GHashTable

	HashTableNewFull func(hashFunc T.GHashFunc,
		keyEqualFunc T.GEqualFunc,
		keyDestroyFunc T.GDestroyNotify,
		valueDestroyFunc T.GDestroyNotify) *T.GHashTable

	HashTableDestroy func(hashTable *T.GHashTable)

	HashTableInsert func(hashTable *T.GHashTable,
		key T.Gpointer,
		value T.Gpointer)

	HashTableReplace func(hashTable *T.GHashTable,
		key T.Gpointer,
		value T.Gpointer)

	HashTableRemove func(hashTable *T.GHashTable,
		key T.Gconstpointer) T.Gboolean

	HashTableRemoveAll func(hashTable *T.GHashTable)

	HashTableSteal func(hashTable *T.GHashTable,
		key T.Gconstpointer) T.Gboolean

	HashTableStealAll func(hashTable *T.GHashTable)

	HashTableLookup func(hashTable *T.GHashTable,
		key T.Gconstpointer) T.Gpointer

	HashTableLookupExtended func(hashTable *T.GHashTable,
		lookupKey T.Gconstpointer,
		origKey, value *T.Gpointer) T.Gboolean

	HashTableForeach func(hashTable *T.GHashTable,
		f T.GHFunc, userData T.Gpointer)

	HashTableFind func(hashTable *T.GHashTable,
		predicate T.GHRFunc, userData T.Gpointer) T.Gpointer

	HashTableForeachRemove func(hashTable *T.GHashTable,
		f T.GHRFunc, userData T.Gpointer) uint

	HashTableForeachSteal func(hashTable *T.GHashTable,
		f T.GHRFunc, userData T.Gpointer) uint

	HashTableSize func(hashTable *T.GHashTable) uint

	HashTableGetKeys func(hashTable *T.GHashTable) *T.GList

	HashTableGetValues func(hashTable *T.GHashTable) *T.GList

	HashTableIterInit func(iter *T.GHashTableIter,
		hashTable *T.GHashTable)

	HashTableIterNext func(iter *T.GHashTableIter,
		key *T.Gpointer,
		value *T.Gpointer) T.Gboolean

	HashTableIterGetHashTable func(iter *T.GHashTableIter) *T.GHashTable

	HashTableIterRemove func(iter *T.GHashTableIter)

	HashTableIterSteal func(iter *T.GHashTableIter)

	HashTableRef func(hashTable *T.GHashTable) *T.GHashTable

	HashTableUnref func(hashTable *T.GHashTable)

	StrEqual func(v1 T.Gconstpointer,
		v2 T.Gconstpointer) T.Gboolean

	StrHash func(v T.Gconstpointer) uint

	IntEqual func(v1 T.Gconstpointer,
		v2 T.Gconstpointer) T.Gboolean

	IntHash func(v T.Gconstpointer) uint

	Int64Equal func(v1 T.Gconstpointer,
		v2 T.Gconstpointer) T.Gboolean

	Int64Hash func(v T.Gconstpointer) uint

	DoubleEqual func(v1 T.Gconstpointer,
		v2 T.Gconstpointer) T.Gboolean

	DoubleHash func(v T.Gconstpointer) uint

	DirectHash func(v T.Gconstpointer) uint

	DirectEqual func(v1 T.Gconstpointer,
		v2 T.Gconstpointer) T.Gboolean

	HookListInit func(hookList *T.GHookList,
		hookSize uint)

	HookListClear func(hookList *T.GHookList)

	HookAlloc func(hookList *T.GHookList) *T.GHook

	HookFree func(hookList *T.GHookList,
		hook *T.GHook)

	HookRef func(hookList *T.GHookList,
		hook *T.GHook) *T.GHook

	HookUnref func(hookList *T.GHookList,
		hook *T.GHook)

	HookDestroy func(hookList *T.GHookList,
		hookId T.Gulong) T.Gboolean

	HookDestroyLink func(hookList *T.GHookList,
		hook *T.GHook)

	HookPrepend func(hookList *T.GHookList,
		hook *T.GHook)

	HookInsertBefore func(hookList *T.GHookList,
		sibling *T.GHook,
		hook *T.GHook)

	HookInsertSorted func(hookList *T.GHookList,
		hook *T.GHook,
		f T.GHookCompareFunc)

	HookGet func(hookList *T.GHookList,
		hookId T.Gulong) *T.GHook

	HookFind func(hookList *T.GHookList,
		needValids T.Gboolean,
		f T.GHookFindFunc,
		data T.Gpointer) *T.GHook

	HookFindData func(hookList *T.GHookList,
		needValids T.Gboolean,
		data T.Gpointer) *T.GHook

	HookFindFunc func(hookList *T.GHookList,
		needValids T.Gboolean,
		f T.Gpointer) *T.GHook

	HookFindFuncData func(hookList *T.GHookList,
		needValids T.Gboolean,
		f T.Gpointer,
		data T.Gpointer) *T.GHook

	HookFirstValid func(hookList *T.GHookList,
		mayBeInCall T.Gboolean) *T.GHook

	HookNextValid func(hookList *T.GHookList,
		hook *T.GHook,
		mayBeInCall T.Gboolean) *T.GHook

	HookCompareIds func(newHook *T.GHook,
		sibling *T.GHook) int

	HookListInvoke func(hookList *T.GHookList,
		mayRecurse T.Gboolean)

	HookListInvokeCheck func(hookList *T.GHookList,
		mayRecurse T.Gboolean)

	HookListMarshal func(hookList *T.GHookList,
		mayRecurse T.Gboolean,
		marshaller T.GHookMarshaller,
		marshalData T.Gpointer)

	HookListMarshalCheck func(hookList *T.GHookList,
		mayRecurse T.Gboolean,
		marshaller T.GHookCheckMarshaller,
		marshalData T.Gpointer)

	HostnameIsNonAscii func(hostname string) T.Gboolean

	HostnameIsAsciiEncoded func(hostname string) T.Gboolean

	HostnameIsIpAddress func(hostname string) T.Gboolean

	HostnameToAscii func(hostname string) string

	HostnameToUnicode func(hostname string) string

	Poll func(fds *T.GPollFD,
		nfds uint,
		timeout int) int

	SlistAlloc func() *T.GSList

	SlistFree func(list *T.GSList)

	SlistFree1 func(list *T.GSList)

	SlistFreeFull func(list *T.GSList,
		freeFunc T.GDestroyNotify)

	SlistAppend func(list *T.GSList,
		data T.Gpointer) *T.GSList

	SlistPrepend func(list *T.GSList,
		data T.Gpointer) *T.GSList

	SlistInsert func(list *T.GSList,
		data T.Gpointer,
		position int) *T.GSList

	SlistInsertSorted func(list *T.GSList,
		data T.Gpointer,
		f T.GCompareFunc) *T.GSList

	SlistInsertSortedWithData func(list *T.GSList,
		data T.Gpointer,
		f T.GCompareDataFunc,
		userData T.Gpointer) *T.GSList

	SlistInsertBefore func(slist *T.GSList,
		sibling *T.GSList,
		data T.Gpointer) *T.GSList

	SlistConcat func(list1 *T.GSList,
		list2 *T.GSList) *T.GSList

	SlistRemove func(list *T.GSList,
		data T.Gconstpointer) *T.GSList

	SlistRemoveAll func(list *T.GSList,
		data T.Gconstpointer) *T.GSList

	SlistRemoveLink func(list *T.GSList,
		link *T.GSList) *T.GSList

	SlistDeleteLink func(list *T.GSList,
		link *T.GSList) *T.GSList

	SlistReverse func(list *T.GSList) *T.GSList

	SlistCopy func(list *T.GSList) *T.GSList

	SlistNth func(list *T.GSList,
		n uint) *T.GSList

	SlistFind func(list *T.GSList,
		data T.Gconstpointer) *T.GSList

	SlistFindCustom func(list *T.GSList,
		data T.Gconstpointer,
		f T.GCompareFunc) *T.GSList

	SlistPosition func(list *T.GSList,
		llink *T.GSList) int

	SlistIndex func(list *T.GSList,
		data T.Gconstpointer) int

	SlistLast func(list *T.GSList) *T.GSList

	SlistLength func(list *T.GSList) uint

	SlistForeach func(list *T.GSList,
		f T.GFunc,
		userData T.Gpointer)

	SlistSort func(list *T.GSList,
		compareFunc T.GCompareFunc) *T.GSList

	SlistSortWithData func(list *T.GSList,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer) *T.GSList

	SlistNthData func(list *T.GSList,
		n uint) T.Gpointer

	SlistPushAllocator func(dummy T.Gpointer)

	SlistPopAllocator func()

	MainContextNew func() *T.GMainContext

	MainContextRef func(context *T.GMainContext) *T.GMainContext

	MainContextUnref func(context *T.GMainContext)

	MainContextDefault func() *T.GMainContext

	MainContextIteration func(context *T.GMainContext,
		mayBlock T.Gboolean) T.Gboolean

	MainContextPending func(context *T.GMainContext) T.Gboolean

	MainContextFindSourceById func(context *T.GMainContext,
		sourceId uint) *O.Source

	MainContextFindSourceByUserData func(context *T.GMainContext,
		userData T.Gpointer) *O.Source

	MainContextFindSourceByFuncsUserData func(context *T.GMainContext,
		funcs *O.SourceFuncs,
		userData T.Gpointer) *O.Source

	MainContextWakeup func(context *T.GMainContext)

	MainContextAcquire func(context *T.GMainContext) T.Gboolean

	MainContextRelease func(context *T.GMainContext)

	MainContextIsOwner func(context *T.GMainContext) T.Gboolean

	MainContextWait func(context *T.GMainContext,
		cond *T.GCond, mutex *T.GMutex) T.Gboolean

	MainContextPrepare func(context *T.GMainContext,
		priority *int) T.Gboolean

	MainContextQuery func(context *T.GMainContext,
		maxPriority int, timeout *int,
		fds *T.GPollFD, nFds int) int

	MainContextCheck func(context *T.GMainContext,
		maxPriority int, fds *T.GPollFD,
		nFds int) int

	MainContextDispatch func(context *T.GMainContext)

	MainContextSetPollFunc func(context *T.GMainContext,
		f T.GPollFunc)

	MainContextGetPollFunc func(context *T.GMainContext) T.GPollFunc

	MainContextAddPoll func(context *T.GMainContext,
		fd *T.GPollFD,
		priority int)

	MainContextRemovePoll func(context *T.GMainContext,
		fd *T.GPollFD)

	MainDepth func() int

	MainCurrentSource func() *O.Source

	MainContextPushThreadDefault func(context *T.GMainContext)

	MainContextPopThreadDefault func(context *T.GMainContext)

	MainContextGetThreadDefault func() *T.GMainContext

	MainLoopNew func(context *T.GMainContext,
		isRunning T.Gboolean) *T.GMainLoop

	MainLoopRun func(loop *T.GMainLoop)

	MainLoopQuit func(loop *T.GMainLoop)

	MainLoopRef func(loop *T.GMainLoop) *T.GMainLoop

	MainLoopUnref func(loop *T.GMainLoop)

	MainLoopIsRunning func(loop *T.GMainLoop) T.Gboolean

	MainLoopGetContext func(loop *T.GMainLoop) *T.GMainContext

	SourceNew func(sourceFuncs *O.SourceFuncs,

		structSize uint) *O.Source

	SourceRef func(source *O.Source) *O.Source

	SourceUnref func(source *O.Source)

	SourceAttach func(source *O.Source,
		context *T.GMainContext) uint

	SourceDestroy func(source *O.Source)

	SourceSetPriority func(source *O.Source,
		priority int)

	SourceGetPriority func(source *O.Source) int

	SourceSetCanRecurse func(source *O.Source,
		canRecurse T.Gboolean)

	SourceGetCanRecurse func(source *O.Source) T.Gboolean

	SourceGetId func(source *O.Source) uint

	SourceGetContext func(source *O.Source) *T.GMainContext

	SourceSetCallback func(source *O.Source,
		f O.SourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify)

	SourceSetFuncs func(source *O.Source,
		funcs *O.SourceFuncs)

	SourceIsDestroyed func(source *O.Source) T.Gboolean

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

	IdleSourceNew func() *O.Source

	ChildWatchSourceNew func(pid T.GPid) *O.Source

	TimeoutSourceNew func(interval uint) *O.Source

	TimeoutSourceNewSeconds func(interval uint) *O.Source

	GetCurrentTime func(result *T.GTimeVal)

	GetMonotonicTime func() int64

	GetRealTime func() int64

	SourceRemove func(tag uint) T.Gboolean

	SourceRemoveByUserData func(userData T.Gpointer) T.Gboolean

	SourceRemoveByFuncsUserData func(funcs *O.SourceFuncs,
		userData T.Gpointer) T.Gboolean

	TimeoutAddFull func(priority int,
		interval uint,
		function O.SourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify) uint

	TimeoutAdd func(interval uint,
		function O.SourceFunc,
		data T.Gpointer) uint

	TimeoutAddSecondsFull func(priority int,
		interval uint,
		function O.SourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify) uint

	TimeoutAddSeconds func(interval uint,
		function O.SourceFunc,
		data T.Gpointer) uint

	ChildWatchAddFull func(priority int,
		pid T.GPid,
		function T.GChildWatchFunc,
		data T.Gpointer,
		notify T.GDestroyNotify) uint

	ChildWatchAdd func(pid T.GPid,
		function T.GChildWatchFunc,
		data T.Gpointer) uint

	IdleAdd func(function O.SourceFunc,
		data T.Gpointer) uint

	IdleAddFull func(priority int,
		function O.SourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify) uint

	IdleRemoveByData func(data T.Gpointer) T.Gboolean

	MainContextInvokeFull func(context *T.GMainContext,
		priority int,
		function O.SourceFunc,
		data T.Gpointer,
		notify T.GDestroyNotify)

	MainContextInvoke func(context *T.GMainContext,
		function O.SourceFunc,
		data T.Gpointer)

	GetCharset func(charset **T.Char) T.Gboolean

	UnicharIsalnum func(c T.Gunichar) T.Gboolean

	UnicharIsalpha func(c T.Gunichar) T.Gboolean

	UnicharIscntrl func(c T.Gunichar) T.Gboolean

	UnicharIsdigit func(c T.Gunichar) T.Gboolean

	UnicharIsgraph func(c T.Gunichar) T.Gboolean

	UnicharIslower func(c T.Gunichar) T.Gboolean

	UnicharIsprint func(c T.Gunichar) T.Gboolean

	UnicharIspunct func(c T.Gunichar) T.Gboolean

	UnicharIsspace func(c T.Gunichar) T.Gboolean

	UnicharIsupper func(c T.Gunichar) T.Gboolean

	UnicharIsxdigit func(c T.Gunichar) T.Gboolean

	UnicharIstitle func(c T.Gunichar) T.Gboolean

	UnicharIsdefined func(c T.Gunichar) T.Gboolean

	UnicharIswide func(c T.Gunichar) T.Gboolean

	UnicharIswideCjk func(c T.Gunichar) T.Gboolean

	UnicharIszerowidth func(c T.Gunichar) T.Gboolean

	UnicharIsmark func(c T.Gunichar) T.Gboolean

	UnicharToupper func(c T.Gunichar) T.Gunichar

	UnicharTolower func(c T.Gunichar) T.Gunichar

	UnicharTotitle func(c T.Gunichar) T.Gunichar

	UnicharDigitValue func(c T.Gunichar) int

	UnicharXdigitValue func(c T.Gunichar) int

	UnicharType func(c T.Gunichar) T.GUnicodeType

	UnicharBreakType func(c T.Gunichar) T.GUnicodeBreakType

	UnicharCombiningClass func(uc T.Gunichar) int

	UnicodeCanonicalOrdering func(str *T.Gunichar,
		leng T.Gsize)

	UnicodeCanonicalDecomposition func(ch T.Gunichar,
		resultLen *T.Gsize) *T.Gunichar

	Utf8GetChar func(p string) T.Gunichar

	Utf8GetCharValidated func(p string,
		maxLen T.Gssize) T.Gunichar

	Utf8OffsetToPointer func(str string,
		offset T.Glong) string

	Utf8PointerToOffset func(str string,
		pos string) T.Glong

	Utf8PrevChar func(p string) string

	Utf8FindNextChar func(p string,
		end string) string

	Utf8FindPrevChar func(str string,
		p string) string

	Utf8Strlen func(p string,
		max T.Gssize) T.Glong

	Utf8Strncpy func(dest string,
		src string,
		n T.Gsize) string

	Utf8Strchr func(p string,
		leng T.Gssize,
		c T.Gunichar) string

	Utf8Strrchr func(p string,
		leng T.Gssize,
		c T.Gunichar) string

	Utf8Strreverse func(str string,
		leng T.Gssize) string

	Utf8ToUtf16 func(str string,
		leng T.Glong,
		itemsRead *T.Glong,
		itemsWritten *T.Glong,
		e **T.GError) *T.Gunichar2

	Utf8ToUcs4 func(str string,
		leng T.Glong,
		itemsRead *T.Glong,
		itemsWritten *T.Glong,
		e **T.GError) *T.Gunichar

	Utf8ToUcs4Fast func(str string,
		leng T.Glong,
		itemsWritten *T.Glong) *T.Gunichar

	Utf16ToUcs4 func(str *T.Gunichar2,
		leng T.Glong,
		itemsRead *T.Glong,
		itemsWritten *T.Glong,
		e **T.GError) *T.Gunichar

	Utf16ToUtf8 func(str *T.Gunichar2,
		leng T.Glong,
		itemsRead *T.Glong,
		itemsWritten *T.Glong,
		e **T.GError) string

	Ucs4ToUtf16 func(str *T.Gunichar,
		leng T.Glong,
		itemsRead *T.Glong,
		itemsWritten *T.Glong,
		e **T.GError) *T.Gunichar2

	Ucs4ToUtf8 func(str *T.Gunichar,
		leng T.Glong,
		itemsRead *T.Glong,
		itemsWritten *T.Glong,
		e **T.GError) string

	UnicharToUtf8 func(c T.Gunichar,
		outbuf string) int

	Utf8Validate func(str string,
		maxLen T.Gssize,
		end **T.Gchar) T.Gboolean

	UnicharValidate func(ch T.Gunichar) T.Gboolean
	Utf8Strup       func(str string,
		leng T.Gssize) string

	Utf8Strdown func(str string,
		leng T.Gssize) string

	Utf8Casefold func(str string,
		leng T.Gssize) string

	Utf8Normalize func(str string,
		leng T.Gssize,
		mode T.GNormalizeMode) string

	Utf8Collate func(str1 string,
		str2 string) int

	Utf8CollateKey func(str string,
		leng T.Gssize) string

	Utf8CollateKeyForFilename func(str string,
		leng T.Gssize) string

	UnicharGetMirrorChar func(ch T.Gunichar,
		mirroredCh *T.Gunichar) T.Gboolean

	UnicharGetScript func(ch T.Gunichar) T.GUnicodeScript

	StringChunkNew func(size T.Gsize) *T.GStringChunk

	StringChunkFree func(chunk *T.GStringChunk)

	StringChunkClear func(chunk *T.GStringChunk)

	StringChunkInsert func(chunk *T.GStringChunk,
		str string) string

	StringChunkInsertLen func(chunk *T.GStringChunk,
		str string,
		leng T.Gssize) string

	StringNew func(init string) *T.GString

	StringNewLen func(init string,
		leng T.Gssize) *T.GString

	StringSizedNew func(dflSize T.Gsize) *T.GString

	StringFree func(str *T.GString,
		freeSegment T.Gboolean) string

	StringEqual func(v *T.GString,
		v2 *T.GString) T.Gboolean

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
		allowUtf8 T.Gboolean) *T.GString

	IoChannelInit func(channel *T.GIOChannel)

	IoChannelRef func(channel *T.GIOChannel) *T.GIOChannel

	IoChannelUnref func(channel *T.GIOChannel)

	IoChannelRead func(channel *T.GIOChannel,
		buf string, count T.Gsize,
		bytesRead *T.Gsize) T.GIOError

	IoChannelWrite func(channel *T.GIOChannel,
		buf string, count T.Gsize,
		bytesWritten *T.Gsize) T.GIOError

	IoChannelSeek func(channel *T.GIOChannel,
		offset int64, typ T.GSeekType) T.GIOError

	IoChannelClose func(channel *T.GIOChannel)

	IoChannelShutdown func(channel *T.GIOChannel,
		flush T.Gboolean, err **T.GError) T.GIOStatus

	IoAddWatchFull func(channel *T.GIOChannel,
		priority int,
		condition T.GIOCondition,
		f T.GIOFunc,
		userData T.Gpointer,
		notify T.GDestroyNotify) uint

	IoCreateWatch func(channel *T.GIOChannel,
		condition T.GIOCondition) *O.Source

	IoAddWatch func(channel *T.GIOChannel,
		condition T.GIOCondition,
		f T.GIOFunc,
		userData T.Gpointer) uint

	IoChannelSetBufferSize func(channel *T.GIOChannel,
		size T.Gsize)

	IoChannelGetBufferSize func(channel *T.GIOChannel) T.Gsize

	IoChannelGetBufferCondition func(channel *T.GIOChannel) T.GIOCondition

	IoChannelSetFlags func(channel *T.GIOChannel,
		flags T.GIOFlags,
		e **T.GError) T.GIOStatus

	IoChannelGetFlags func(channel *T.GIOChannel) T.GIOFlags

	IoChannelSetLineTerm func(channel *T.GIOChannel,
		lineTerm string,
		length int)

	IoChannelGetLineTerm func(channel *T.GIOChannel,
		length *int) string

	IoChannelSetBuffered func(channel *T.GIOChannel,
		buffered T.Gboolean)

	IoChannelGetBuffered func(channel *T.GIOChannel) T.Gboolean

	IoChannelSetEncoding func(channel *T.GIOChannel,
		encoding string,
		e **T.GError) T.GIOStatus

	IoChannelGetEncoding func(channel *T.GIOChannel) string

	IoChannelSetCloseOnUnref func(channel *T.GIOChannel,
		doClose T.Gboolean)

	IoChannelGetCloseOnUnref func(channel *T.GIOChannel) T.Gboolean

	IoChannelFlush func(channel *T.GIOChannel,
		e **T.GError) T.GIOStatus

	IoChannelReadLine func(channel *T.GIOChannel,
		strReturn **T.Gchar,
		length *T.Gsize,
		terminatorPos *T.Gsize,
		e **T.GError) T.GIOStatus

	IoChannelReadLineString func(channel *T.GIOChannel,
		buffer *T.GString,
		terminatorPos *T.Gsize,
		e **T.GError) T.GIOStatus

	IoChannelReadToEnd func(channel *T.GIOChannel,
		strReturn **T.Gchar,
		length *T.Gsize,
		e **T.GError) T.GIOStatus

	IoChannelReadChars func(channel *T.GIOChannel,
		buf string,
		count T.Gsize,
		bytesRead *T.Gsize,
		e **T.GError) T.GIOStatus

	IoChannelReadUnichar func(channel *T.GIOChannel,
		thechar *T.Gunichar,
		e **T.GError) T.GIOStatus

	IoChannelWriteChars func(channel *T.GIOChannel,
		buf string,
		count T.Gssize,
		bytesWritten *T.Gsize,
		e **T.GError) T.GIOStatus

	IoChannelWriteUnichar func(channel *T.GIOChannel,
		thechar T.Gunichar,
		e **T.GError) T.GIOStatus

	IoChannelSeekPosition func(channel *T.GIOChannel,
		offset int64,
		typ T.GSeekType,
		e **T.GError) T.GIOStatus

	IoChannelNewFile func(filename string,
		mode string,
		e **T.GError) *T.GIOChannel

	IoChannelErrorQuark func() T.GQuark

	IoChannelErrorFromErrno func(en int) T.GIOChannelError

	IoChannelUnixNew func(fd int) *T.GIOChannel

	IoChannelUnixGetFd func(channel *T.GIOChannel) int

	IoChannelWin32MakePollfd func(channel *T.GIOChannel,
		condition T.GIOCondition,
		fd *T.GPollFD)

	IoChannelWin32Poll func(fds *T.GPollFD,
		nFds int,
		timeout int) int

	IoChannelWin32NewMessages func(hwnd uint) *T.GIOChannel

	IoChannelWin32NewFd func(fd int) *T.GIOChannel

	IoChannelWin32GetFd func(channel *T.GIOChannel) int

	IoChannelWin32NewSocket func(socket int) *T.GIOChannel

	KeyFileNew func() *T.GKeyFile

	KeyFileFree func(keyFile *T.GKeyFile)

	KeyFileSetListSeparator func(keyFile *T.GKeyFile,
		separator T.Gchar)

	KeyFileLoadFromFile func(keyFile *T.GKeyFile,
		file string,
		flags T.GKeyFileFlags,
		e **T.GError) T.Gboolean

	KeyFileLoadFromData func(keyFile *T.GKeyFile,
		data string,
		length T.Gsize,
		flags T.GKeyFileFlags,
		e **T.GError) T.Gboolean

	KeyFileLoadFromDirs func(keyFile *T.GKeyFile,
		file string,
		searchDirs **T.Gchar,
		fullPath **T.Gchar,
		flags T.GKeyFileFlags,
		e **T.GError) T.Gboolean

	KeyFileLoadFromDataDirs func(keyFile *T.GKeyFile,
		file string,
		fullPath **T.Gchar,
		flags T.GKeyFileFlags,
		e **T.GError) T.Gboolean

	KeyFileToData func(keyFile *T.GKeyFile,
		length *T.Gsize,
		e **T.GError) string

	KeyFileGetStartGroup func(keyFile *T.GKeyFile) string

	KeyFileGetGroups func(keyFile *T.GKeyFile,
		length *T.Gsize) **T.Gchar

	KeyFileGetKeys func(keyFile *T.GKeyFile,
		groupName string,
		length *T.Gsize,
		e **T.GError) **T.Gchar

	KeyFileHasGroup func(keyFile *T.GKeyFile,
		groupName string) T.Gboolean

	KeyFileHasKey func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) T.Gboolean

	KeyFileGetValue func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) string

	KeyFileSetValue func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		value string)

	KeyFileGetString func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) string

	KeyFileSetString func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		str string)

	KeyFileGetLocaleString func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		locale string,
		e **T.GError) string

	KeyFileSetLocaleString func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		locale string,
		str string)

	KeyFileGetBoolean func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) T.Gboolean

	KeyFileSetBoolean func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		value T.Gboolean)

	KeyFileGetInteger func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) int

	KeyFileSetInteger func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		value int)

	KeyFileGetInt64 func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) int64

	KeyFileSetInt64 func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		value int64)

	KeyFileGetUint64 func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) uint64

	KeyFileSetUint64 func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		value uint64)

	KeyFileGetDouble func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) float64

	KeyFileSetDouble func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		value float64)

	KeyFileGetStringList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		length *T.Gsize,
		e **T.GError) **T.Gchar

	KeyFileSetStringList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		list **T.Gchar,
		length T.Gsize)

	KeyFileGetLocaleStringList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		locale string,
		length *T.Gsize,
		e **T.GError) **T.Gchar

	KeyFileSetLocaleStringList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		locale string,
		list **T.Gchar,
		length T.Gsize)

	KeyFileGetBooleanList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		length *T.Gsize,
		e **T.GError) *T.Gboolean

	KeyFileSetBooleanList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		list *T.Gboolean,
		length T.Gsize)

	KeyFileGetIntegerList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		length *T.Gsize,
		e **T.GError) *int

	KeyFileSetDoubleList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		list *float64,
		length T.Gsize)

	KeyFileGetDoubleList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		length *T.Gsize,
		e **T.GError) *float64

	KeyFileSetIntegerList func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		list *int,
		length T.Gsize)

	KeyFileSetComment func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		comment string,
		e **T.GError) T.Gboolean

	KeyFileGetComment func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) string

	KeyFileRemoveComment func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) T.Gboolean

	KeyFileRemoveKey func(keyFile *T.GKeyFile,
		groupName string,
		key string,
		e **T.GError) T.Gboolean

	KeyFileRemoveGroup func(keyFile *T.GKeyFile,
		groupName string,
		e **T.GError) T.Gboolean

	MappedFileNew func(filename string,
		writable T.Gboolean,
		e **T.GError) *T.GMappedFile

	MappedFileGetLength func(file *T.GMappedFile) T.Gsize

	MappedFileGetContents func(file *T.GMappedFile) string

	MappedFileRef func(file *T.GMappedFile) *T.GMappedFile

	MappedFileUnref func(file *T.GMappedFile)

	MappedFileFree func(file *T.GMappedFile)

	MarkupParseContextNew func(parser *T.GMarkupParser,
		flags T.GMarkupParseFlags,
		userData T.Gpointer,
		userDataDnotify T.GDestroyNotify) *T.GMarkupParseContext

	MarkupParseContextFree func(context *T.GMarkupParseContext)

	MarkupParseContextParse func(context *T.GMarkupParseContext,
		text string,
		textLen T.Gssize,
		e **T.GError) T.Gboolean

	MarkupParseContextPush func(context *T.GMarkupParseContext,
		parser *T.GMarkupParser,
		userData T.Gpointer)

	MarkupParseContextPop func(context *T.GMarkupParseContext) T.Gpointer

	MarkupParseContextEndParse func(context *T.GMarkupParseContext,
		e **T.GError) T.Gboolean

	MarkupParseContextGetElement func(context *T.GMarkupParseContext) string

	MarkupParseContextGetElementStack func(context *T.GMarkupParseContext) *T.GSList

	MarkupParseContextGetPosition func(context *T.GMarkupParseContext,
		lineNumber *int,
		charNumber *int)

	MarkupParseContextGetUserData func(context *T.GMarkupParseContext) T.Gpointer

	MarkupEscapeText func(text string,
		length T.Gssize) string

	MarkupPrintfEscaped func(format string, v ...VArg) string

	MarkupVprintfEscaped func(format string,
		args T.VaList) string

	LogSetHandler func(logDomain string,
		logLevels T.GLogLevelFlags,
		logFunc T.GLogFunc,
		userData T.Gpointer) uint

	LogRemoveHandler func(logDomain string,
		handlerId uint)

	LogDefaultHandler func(logDomain string,
		logLevel T.GLogLevelFlags,
		message string,
		unusedData T.Gpointer)

	LogSetDefaultHandler func(logFunc T.GLogFunc,
		userData T.Gpointer) T.GLogFunc

	Log func(logDomain string, logLevel T.GLogLevelFlags,
		format string, v ...VArg)

	Logv func(logDomain string,
		logLevel T.GLogLevelFlags,
		format string,
		args T.VaList)

	LogSetFatalMask func(logDomain string,
		fatalMask T.GLogLevelFlags) T.GLogLevelFlags

	LogSetAlwaysFatal func(fatalMask T.GLogLevelFlags) T.GLogLevelFlags

	ReturnIfFailWarning func(logDomain string,
		prettyFunction string,
		expression string)

	WarnMessage func(domain string,
		file string,
		line int,
		f string,
		warnexpr string)

	AssertWarning func(logDomain string,
		file string,
		line int,
		prettyFunction string,
		expression string)

	Print func(format string, v ...VArg)

	SetPrintHandler func(f T.GPrintFunc) T.GPrintFunc

	Printerr func(format string, v ...VArg)

	SetPrinterrHandler func(f T.GPrintFunc) T.GPrintFunc

	NodeNew func(data T.Gpointer) *T.GNode

	NodeDestroy func(root *T.GNode)

	NodeUnlink func(node *T.GNode)

	NodeCopyDeep func(node *T.GNode,
		copyFunc T.GCopyFunc, data T.Gpointer) *T.GNode

	NodeCopy func(node *T.GNode) *T.GNode

	NodeInsert func(parent *T.GNode,
		position int, node *T.GNode) *T.GNode

	NodeInsertBefore func(parent *T.GNode,
		sibling *T.GNode, node *T.GNode) *T.GNode

	NodeInsertAfter func(parent *T.GNode,
		sibling *T.GNode, node *T.GNode) *T.GNode

	NodePrepend func(parent *T.GNode, node *T.GNode) *T.GNode

	NodeNNodes func(root *T.GNode,
		flags T.GTraverseFlags) uint

	NodeGetRoot func(node *T.GNode) *T.GNode

	NodeIsAncestor func(node *T.GNode,
		descendant *T.GNode) T.Gboolean

	NodeDepth func(node *T.GNode) uint

	NodeFind func(root *T.GNode,
		order T.GTraverseType,
		flags T.GTraverseFlags,
		data T.Gpointer) *T.GNode

	NodeTraverse func(root *T.GNode,
		order T.GTraverseType,
		flags T.GTraverseFlags,
		maxDepth int,
		f T.GNodeTraverseFunc,
		data T.Gpointer)

	NodeMaxHeight func(root *T.GNode) uint

	NodeChildrenForeach func(node *T.GNode,
		flags T.GTraverseFlags,
		f T.GNodeForeachFunc,
		data T.Gpointer)

	NodeReverseChildren func(node *T.GNode)

	NodeNChildren func(node *T.GNode) uint

	NodeNthChild func(node *T.GNode,
		n uint) *T.GNode

	NodeLastChild func(node *T.GNode) *T.GNode

	NodeFindChild func(node *T.GNode,
		flags T.GTraverseFlags,
		data T.Gpointer) *T.GNode

	NodeChildPosition func(node *T.GNode,
		child *T.GNode) int

	NodeChildIndex func(node *T.GNode,
		data T.Gpointer) int

	NodeFirstSibling func(node *T.GNode) *T.GNode

	NodeLastSibling func(node *T.GNode) *T.GNode

	NodePushAllocator func(dummy T.Gpointer)

	NodePopAllocator func()

	OptionContextNew func(parameterString string) *T.GOptionContext

	OptionContextSetSummary func(context *T.GOptionContext,
		summary string)

	OptionContextGetSummary func(context *T.GOptionContext) string

	OptionContextSetDescription func(context *T.GOptionContext,
		description string)

	OptionContextGetDescription func(context *T.GOptionContext) string

	OptionContextFree func(context *T.GOptionContext)

	OptionContextSetHelpEnabled func(context *T.GOptionContext,
		helpEnabled T.Gboolean)

	OptionContextGetHelpEnabled func(context *T.GOptionContext) T.Gboolean

	OptionContextSetIgnoreUnknownOptions func(context *T.GOptionContext,
		ignoreUnknown T.Gboolean)

	OptionContextGetIgnoreUnknownOptions func(context *T.GOptionContext) T.Gboolean

	OptionContextAddMainEntries func(context *T.GOptionContext,
		entries *T.GOptionEntry,
		translationDomain string)

	OptionContextParse func(context *T.GOptionContext,
		argc *int,
		argv ***T.Gchar,
		e **T.GError) T.Gboolean

	OptionContextSetTranslateFunc func(context *T.GOptionContext,
		f T.GTranslateFunc,
		data T.Gpointer,
		destroyNotify T.GDestroyNotify)

	OptionContextSetTranslationDomain func(context *T.GOptionContext,
		domain string)

	OptionContextAddGroup func(context *T.GOptionContext,
		group *T.GOptionGroup)

	OptionContextSetMainGroup func(context *T.GOptionContext,
		group *T.GOptionGroup)
	OptionContextGetMainGroup func(context *T.GOptionContext) *T.GOptionGroup

	OptionContextGetHelp func(context *T.GOptionContext,
		mainHelp T.Gboolean,
		group *T.GOptionGroup) string

	OptionGroupNew func(name string,
		description string,
		helpDescription string,
		userData T.Gpointer,
		destroy T.GDestroyNotify) *T.GOptionGroup

	OptionGroupSetParseHooks func(group *T.GOptionGroup,
		preParseFunc T.GOptionParseFunc,
		postParseFunc T.GOptionParseFunc)

	OptionGroupSetErrorHook func(group *T.GOptionGroup,
		errorFunc T.GOptionErrorFunc)

	OptionGroupFree func(group *T.GOptionGroup)

	OptionGroupAddEntries func(group *T.GOptionGroup,
		entries *T.GOptionEntry)

	OptionGroupSetTranslateFunc func(group *T.GOptionGroup,
		f T.GTranslateFunc,
		data T.Gpointer,
		destroyNotify T.GDestroyNotify)

	OptionGroupSetTranslationDomain func(group *T.GOptionGroup,
		domain string)

	PatternSpecNew func(pattern string) *T.GPatternSpec

	PatternSpecFree func(pspec *T.GPatternSpec)

	PatternSpecEqual func(pspec1 *T.GPatternSpec,
		pspec2 *T.GPatternSpec) T.Gboolean

	PatternMatch func(pspec *T.GPatternSpec,
		stringLength uint,
		str string,
		stringReversed string) T.Gboolean

	PatternMatchString func(pspec *T.GPatternSpec,
		str string) T.Gboolean

	PatternMatchSimple func(pattern string,
		str string) T.Gboolean

	SpacedPrimesClosest func(num uint) uint

	QsortWithData func(pbase T.Gconstpointer,
		totalElems int,
		size T.Gsize,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer)

	QueueNew func() *T.GQueue

	QueueFree func(queue *T.GQueue)

	QueueInit func(queue *T.GQueue)

	QueueClear func(queue *T.GQueue)

	QueueIsEmpty func(queue *T.GQueue) T.Gboolean

	QueueGetLength func(queue *T.GQueue) uint

	QueueReverse func(queue *T.GQueue)

	QueueCopy func(queue *T.GQueue) *T.GQueue

	QueueForeach func(queue *T.GQueue,
		f T.GFunc,
		userData T.Gpointer)

	QueueFind func(queue *T.GQueue,
		data T.Gconstpointer) *T.GList

	QueueFindCustom func(queue *T.GQueue,
		data T.Gconstpointer,
		f T.GCompareFunc) *T.GList

	QueueSort func(queue *T.GQueue,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer)

	QueuePushHead func(queue *T.GQueue,
		data T.Gpointer)

	QueuePushTail func(queue *T.GQueue,
		data T.Gpointer)

	QueuePushNth func(queue *T.GQueue,
		data T.Gpointer,
		n int)

	QueuePopHead func(queue *T.GQueue) T.Gpointer

	QueuePopTail func(queue *T.GQueue) T.Gpointer

	QueuePopNth func(queue *T.GQueue,
		n uint) T.Gpointer

	QueuePeekHead func(queue *T.GQueue) T.Gpointer

	QueuePeekTail func(queue *T.GQueue) T.Gpointer

	QueuePeekNth func(queue *T.GQueue,
		n uint) T.Gpointer

	QueueIndex func(queue *T.GQueue,
		data T.Gconstpointer) int

	QueueRemove func(queue *T.GQueue,
		data T.Gconstpointer)

	QueueRemoveAll func(queue *T.GQueue,
		data T.Gconstpointer)

	QueueInsertBefore func(queue *T.GQueue,
		sibling *T.GList,
		data T.Gpointer)

	QueueInsertAfter func(queue *T.GQueue,
		sibling *T.GList,
		data T.Gpointer)

	QueueInsertSorted func(queue *T.GQueue,
		data T.Gpointer,
		f T.GCompareDataFunc,
		userData T.Gpointer)

	QueuePushHeadLink func(queue *T.GQueue,
		link *T.GList)

	QueuePushTailLink func(queue *T.GQueue,
		link *T.GList)

	QueuePushNthLink func(queue *T.GQueue,
		n int,
		link *T.GList)

	QueuePopHeadLink func(queue *T.GQueue) *T.GList

	QueuePopTailLink func(queue *T.GQueue) *T.GList

	QueuePopNthLink func(queue *T.GQueue,
		n uint) *T.GList

	QueuePeekHeadLink func(queue *T.GQueue) *T.GList

	QueuePeekTailLink func(queue *T.GQueue) *T.GList

	QueuePeekNthLink func(queue *T.GQueue,
		n uint) *T.GList

	QueueLinkIndex func(queue *T.GQueue,
		link *T.GList) int

	QueueUnlink func(queue *T.GQueue, link *T.GList)

	QueueDeleteLink func(queue *T.GQueue, link *T.GList)

	RandNewWithSeed func(seed T.GUint32) *T.GRand

	RandNewWithSeedArray func(seed *T.GUint32,
		seedLength uint) *T.GRand

	RandNew func() *T.GRand

	RandFree func(rand *T.GRand)

	RandCopy func(rand *T.GRand) *T.GRand

	RandSetSeed func(rand *T.GRand,
		seed T.GUint32)

	RandSetSeedArray func(rand *T.GRand,
		seed *T.GUint32,
		seedLength uint)

	RandInt func(rand *T.GRand) T.GUint32

	RandIntRange func(rand *T.GRand,
		begin T.GInt32,
		end T.GInt32) T.GInt32

	RandDouble func(rand *T.GRand) float64

	RandDoubleRange func(rand *T.GRand,
		begin float64,
		end float64) float64

	RandomSetSeed func(seed T.GUint32)

	RandomInt func() T.GUint32

	RandomIntRange func(begin T.GInt32,
		end T.GInt32) T.GInt32

	RandomDouble func() float64

	RandomDoubleRange func(begin float64,
		end float64) float64

	RelationNew func(fields int) *T.GRelation

	RelationDestroy func(relation *T.GRelation)

	RelationIndex func(relation *T.GRelation,
		field int,
		hashFunc T.GHashFunc,
		keyEqualFunc T.GEqualFunc)

	RelationInsert func(relation *T.GRelation, v ...VArg)

	RelationDelete func(relation *T.GRelation,
		key T.Gconstpointer,
		field int) int

	RelationSelect func(relation *T.GRelation,
		key T.Gconstpointer,
		field int) *T.GTuples

	RelationCount func(relation *T.GRelation,
		key T.Gconstpointer,
		field int) int

	RelationExists func(
		relation *T.GRelation, v ...VArg) T.Gboolean

	RelationPrint func(relation *T.GRelation)

	TuplesDestroy func(tuples *T.GTuples)

	TuplesIndex func(tuples *T.GTuples,
		index int,
		field int) T.Gpointer

	RegexNew func(pattern string,
		compileOptions T.GRegexCompileFlags,
		matchOptions T.GRegexMatchFlags,
		e **T.GError) *T.GRegex

	RegexRef func(regex *T.GRegex) *T.GRegex

	RegexUnref func(regex *T.GRegex)

	RegexGetPattern func(regex *T.GRegex) string

	RegexGetMaxBackref func(regex *T.GRegex) int

	RegexGetCaptureCount func(regex *T.GRegex) int

	RegexGetStringNumber func(regex *T.GRegex,
		name string) int

	RegexEscapeString func(str string,
		length int) string

	RegexGetCompileFlags func(regex *T.GRegex) T.GRegexCompileFlags

	RegexGetMatchFlags func(regex *T.GRegex) T.GRegexMatchFlags

	RegexMatchSimple func(pattern string,
		str string,
		compileOptions T.GRegexCompileFlags,
		matchOptions T.GRegexMatchFlags) T.Gboolean

	RegexMatch func(regex *T.GRegex,
		str string,
		matchOptions T.GRegexMatchFlags,
		matchInfo **T.GMatchInfo) T.Gboolean

	RegexMatchFull func(regex *T.GRegex,
		str string,
		stringLen T.Gssize,
		startPosition int,
		matchOptions T.GRegexMatchFlags,
		matchInfo **T.GMatchInfo,
		e **T.GError) T.Gboolean

	RegexMatchAll func(regex *T.GRegex,
		str string,
		matchOptions T.GRegexMatchFlags,
		matchInfo **T.GMatchInfo) T.Gboolean

	RegexMatchAllFull func(regex *T.GRegex,
		str string,
		stringLen T.Gssize,
		startPosition int,
		matchOptions T.GRegexMatchFlags,
		matchInfo **T.GMatchInfo,
		e **T.GError) T.Gboolean

	RegexSplitSimple func(pattern string,
		str string,
		compileOptions T.GRegexCompileFlags,
		matchOptions T.GRegexMatchFlags) **T.Gchar

	RegexSplit func(regex *T.GRegex,
		str string,
		matchOptions T.GRegexMatchFlags) **T.Gchar

	RegexSplitFull func(regex *T.GRegex,
		str string,
		stringLen T.Gssize,
		startPosition int,
		matchOptions T.GRegexMatchFlags,
		maxTokens int,
		e **T.GError) **T.Gchar

	RegexReplace func(regex *T.GRegex,
		str string,
		stringLen T.Gssize,
		startPosition int,
		replacement string,
		matchOptions T.GRegexMatchFlags,
		e **T.GError) string

	RegexReplaceLiteral func(regex *T.GRegex,
		str string,
		stringLen T.Gssize,
		startPosition int,
		replacement string,
		matchOptions T.GRegexMatchFlags,
		e **T.GError) string

	RegexReplaceEval func(regex *T.GRegex,
		str string,
		stringLen T.Gssize,
		startPosition int,
		matchOptions T.GRegexMatchFlags,
		eval T.GRegexEvalCallback,
		userData T.Gpointer,
		e **T.GError) string

	RegexCheckReplacement func(replacement string,
		hasReferences *T.Gboolean,
		e **T.GError) T.Gboolean

	MatchInfoGetRegex func(matchInfo *T.GMatchInfo) *T.GRegex

	MatchInfoGetString func(matchInfo *T.GMatchInfo) string

	MatchInfoFree func(matchInfo *T.GMatchInfo)

	MatchInfoNext func(matchInfo *T.GMatchInfo,
		e **T.GError) T.Gboolean

	MatchInfoMatches func(matchInfo *T.GMatchInfo) T.Gboolean

	MatchInfoGetMatchCount func(matchInfo *T.GMatchInfo) int

	MatchInfoIsPartialMatch func(matchInfo *T.GMatchInfo) T.Gboolean

	MatchInfoExpandReferences func(matchInfo *T.GMatchInfo,
		stringToExpand string,
		e **T.GError) string

	MatchInfoFetch func(matchInfo *T.GMatchInfo,
		matchNum int) string

	MatchInfoFetchPos func(matchInfo *T.GMatchInfo,
		matchNum int,
		startPos *int,
		endPos *int) T.Gboolean

	MatchInfoFetchNamed func(matchInfo *T.GMatchInfo,
		name string) string

	MatchInfoFetchNamedPos func(matchInfo *T.GMatchInfo,
		name string, startPos *int, endPos *int) T.Gboolean

	MatchInfoFetchAll func(
		matchInfo *T.GMatchInfo) **T.Gchar

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

	ScannerEof func(scanner *T.GScanner) T.Gboolean

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

	SequenceIterIsBegin func(iter *T.GSequenceIter) T.Gboolean

	SequenceIterIsEnd func(iter *T.GSequenceIter) T.Gboolean

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

	ShellErrorQuark func() T.GQuark

	ShellQuote func(unquotedString string) string

	ShellUnquote func(quotedString string,
		e **T.GError) string

	ShellParseArgv func(commandLine string,
		argcp *int,
		argvp ***T.Gchar,
		e **T.GError) T.Gboolean

	SpawnErrorQuark func() T.GQuark

	SpawnAsync func(workingDirectory string,
		argv **T.Gchar,
		envp **T.Gchar,
		flags T.GSpawnFlags,
		childSetup T.GSpawnChildSetupFunc,
		userData T.Gpointer,
		childPid *T.GPid,
		e **T.GError) T.Gboolean

	SpawnAsyncWithPipes func(workingDirectory string,
		argv **T.Gchar,
		envp **T.Gchar,
		flags T.GSpawnFlags,
		childSetup T.GSpawnChildSetupFunc,
		userData T.Gpointer,
		childPid *T.GPid,
		standardInput *int,
		standardOutput *int,
		standardError *int,
		e **T.GError) T.Gboolean

	SpawnSync func(workingDirectory string,
		argv **T.Gchar,
		envp **T.Gchar,
		flags T.GSpawnFlags,
		childSetup T.GSpawnChildSetupFunc,
		userData T.Gpointer,
		standardOutput **T.Gchar,
		standardError **T.Gchar,
		exitStatus *int,
		e **T.GError) T.Gboolean

	SpawnCommandLineSync func(commandLine string,
		standardOutput **T.Gchar,
		standardError **T.Gchar,
		exitStatus *int,
		e **T.GError) T.Gboolean

	SpawnCommandLineAsync func(commandLine string,
		e **T.GError) T.Gboolean

	SpawnClosePid func(pid T.GPid)

	AsciiTolower func(c T.Gchar) T.Gchar

	AsciiToupper func(c T.Gchar) T.Gchar

	AsciiDigitValue func(c T.Gchar) int

	AsciiXdigitValue func(c T.Gchar) int

	Strdelimit func(str string,
		delimiters string,
		newDelimiter T.Gchar) string

	Strcanon func(str string,
		validChars string,
		substitutor T.Gchar) string

	Strerror func(errnum int) string

	Strsignal func(signum int) string

	Strreverse func(str string) string

	Strlcpy func(dest string,
		src string,
		destSize T.Gsize) T.Gsize

	Strlcat func(dest string,
		src string,
		destSize T.Gsize) T.Gsize

	StrstrLen func(haystack string,
		haystackLen T.Gssize,
		needle string) string

	Strrstr func(haystack string,
		needle string) string

	StrrstrLen func(haystack string,
		haystackLen T.Gssize,
		needle string) string

	StrHasSuffix func(str string,
		suffix string) T.Gboolean

	StrHasPrefix func(str string,
		prefix string) T.Gboolean

	Strtod func(nptr string,
		endptr **T.Gchar) float64

	AsciiStrtod func(nptr string,
		endptr **T.Gchar) float64

	AsciiStrtoull func(nptr string,
		endptr **T.Gchar,
		base uint) uint64

	AsciiStrtoll func(nptr string,
		endptr **T.Gchar,
		base uint) int64

	AsciiDtostr func(buffer string,
		bufLen int,
		d float64) string

	AsciiFormatd func(buffer string,
		bufLen int,
		format string,
		d float64) string

	Strchug func(str string) string

	Strchomp func(str string) string

	AsciiStrcasecmp func(s1 string,
		s2 string) int

	AsciiStrncasecmp func(s1 string,
		s2 string,
		n T.Gsize) int

	AsciiStrdown func(str string,
		leng T.Gssize) string

	AsciiStrup func(str string,
		leng T.Gssize) string

	Strcasecmp func(s1 string,
		s2 string) int

	Strncasecmp func(s1 string,
		s2 string,
		n uint) int

	Strdown func(str string) string

	Strup func(str string) string

	Strdup func(str string) string

	StrdupPrintf func(format string, v ...VArg) string

	StrdupVprintf func(format string,
		args T.VaList) string

	Strndup func(str string,
		n T.Gsize) string

	Strnfill func(length T.Gsize,
		fillChar T.Gchar) string

	Strconcat func(string1 string, v ...VArg) string

	Strjoin func(separator string, v ...VArg) string

	Strcompress func(source string) string

	Strescape func(source string,
		exceptions string) string

	Memdup func(mem T.Gconstpointer,
		byteSize uint) T.Gpointer

	Strsplit func(str string,
		delimiter string,
		maxTokens int) **T.Gchar

	StrsplitSet func(str string,
		delimiters string,
		maxTokens int) **T.Gchar

	Strjoinv func(separator string,
		str_Array **T.Gchar) string

	Strfreev func(strArray **T.Gchar)

	Strdupv func(strArray **T.Gchar) **T.Gchar

	StrvLength func(strArray **T.Gchar) uint

	Stpcpy func(dest, src string) string

	StripContext func(msgid, msgval string) string

	Dgettext func(domain, msgid string) string

	Dcgettext func(domain, msgid string,
		category int) string

	Dngettext func(domain, msgid, msgidPlural string,
		n T.Gulong) string

	Dpgettext func(domain, msgctxtid string,
		msgidoffset T.Gsize) string

	Dpgettext2 func(domain, context, msgid string) string

	Strcmp0 func(str1, str2 string) int

	TestMinimizedResult func(minimizedQuantity float64,
		format string, v ...VArg)

	TestMaximizedResult func(maximizedQuantity float64,
		format string, v ...VArg)

	TestInit func(argc *int, argv ***T.Char, v ...VArg)

	TestRun func() int

	TestAddFunc func(testpath string,
		testFunc T.GTestFunc)

	TestAddDataFunc func(testpath string,
		testData T.Gconstpointer,
		testFunc T.GTestDataFunc)

	TestMessage func(format string, v ...VArg)

	TestBugBase func(uriPattern string)

	TestBug func(bugUriSnippet string)

	TestTimerStart func()

	TestTimerElapsed func() float64

	TestTimerLast func() float64

	TestQueueFree func(gfreePointer T.Gpointer)

	TestQueueDestroy func(destroyFunc T.GDestroyNotify,
		destroyData T.Gpointer)

	TestTrapFork func(usecTimeout uint64,
		testTrapFlags T.GTestTrapFlags) T.Gboolean

	TestTrapHasPassed func() T.Gboolean

	TestTrapReachedTimeout func() T.Gboolean

	TestRandInt func() T.GInt32

	TestRandIntRange func(begin T.GInt32,
		end T.GInt32) T.GInt32

	TestRandDouble func() float64

	TestRandDoubleRange func(rangeStart float64,
		rangeEnd float64) float64

	TestCreateCase func(testName string,
		dataSize T.Gsize,
		testData T.Gconstpointer,
		dataSetup T.GTestFixtureFunc,
		dataTest T.GTestFixtureFunc,
		dataTeardown T.GTestFixtureFunc) *T.GTestCase

	TestCreateSuite func(suiteName string) *T.GTestSuite

	TestGetRoot func() *T.GTestSuite

	TestSuiteAdd func(suite *T.GTestSuite,
		testCase *T.GTestCase)

	TestSuiteAddSuite func(suite *T.GTestSuite,
		nestedsuite *T.GTestSuite)

	TestRunSuite func(suite *T.GTestSuite) int

	TestTrapAssertions func(domain string,
		file string,
		line int,
		f string,
		assertionFlags uint64,
		pattern string)

	AssertionMessage func(domain string,
		file string,
		line int,
		f string,
		message string)

	AssertionMessageExpr func(domain string,
		file string,
		line int,
		f string,
		expr string)

	AssertionMessageCmpstr func(domain string,
		file string,
		line int,
		f string,
		expr string,
		arg1 string,
		cmp string,
		arg2 string)

	AssertionMessageCmpnum func(domain string,
		file string,
		line int,
		f string,
		expr string,
		arg1 T.LongDouble,
		cmp string,
		arg2 T.LongDouble,
		numtype T.Char)

	AssertionMessageError func(domain string,
		file string,
		line int,
		f string,
		expr string,
		e *T.GError,
		errorDomain T.GQuark,
		errorCode int)

	TestAddVtable func(testpath string,
		dataSize T.Gsize,
		testData T.Gconstpointer,
		dataSetup T.GTestFixtureFunc,
		dataTest T.GTestFixtureFunc,
		dataTeardown T.GTestFixtureFunc)

	TestLogTypeName func(logType T.GTestLogType) string

	TestLogBufferNew func() *T.GTestLogBuffer

	TestLogBufferFree func(tbuffer *T.GTestLogBuffer)

	TestLogBufferPush func(tbuffer *T.GTestLogBuffer,
		nBytes uint,
		bytes *uint8)

	TestLogBufferPop func(tbuffer *T.GTestLogBuffer) *T.GTestLogMsg

	TestLogMsgFree func(tmsg *T.GTestLogMsg)

	TestLogSetFatalHandler func(logFunc T.GTestLogFatalFunc,
		userData T.Gpointer)

	ThreadPoolNew func(f T.GFunc,
		userData T.Gpointer,
		maxThreads int,
		exclusive T.Gboolean,
		e **T.GError) *T.GThreadPool

	ThreadPoolPush func(pool *T.GThreadPool,
		data T.Gpointer,
		e **T.GError)

	ThreadPoolSetMaxThreads func(pool *T.GThreadPool,
		maxThreads int,
		e **T.GError)

	ThreadPoolGetMaxThreads func(pool *T.GThreadPool) int

	ThreadPoolGetNumThreads func(pool *T.GThreadPool) uint

	ThreadPoolUnprocessed func(pool *T.GThreadPool) uint

	ThreadPoolFree func(pool *T.GThreadPool,
		immediate T.Gboolean,
		wait T.Gboolean)

	ThreadPoolSetMaxUnusedThreads func(maxThreads int)

	ThreadPoolGetMaxUnusedThreads func() int

	ThreadPoolGetNumUnusedThreads func() uint

	ThreadPoolStopUnusedThreads func()

	ThreadPoolSetSortFunction func(pool *T.GThreadPool,
		f T.GCompareDataFunc,
		userData T.Gpointer)

	ThreadPoolSetMaxIdleTime func(interval uint)

	ThreadPoolGetMaxIdleTime func() uint

	TimerNew func() *T.GTimer

	TimerDestroy func(timer *T.GTimer)

	TimerStart func(timer *T.GTimer)

	TimerStop func(timer *T.GTimer)

	TimerReset func(timer *T.GTimer)

	TimerContinue func(timer *T.GTimer)

	TimerElapsed func(timer *T.GTimer,
		microseconds *T.Gulong) float64

	Usleep func(microseconds T.Gulong)

	TimeValAdd func(time *T.GTimeVal,
		microseconds T.Glong)

	TimeValFromIso8601 func(isoDate string,
		time *T.GTimeVal) T.Gboolean

	TimeValToIso8601 func(time *T.GTimeVal) string

	TreeNew func(keyCompareFunc T.GCompareFunc) *T.GTree

	TreeNewWithData func(keyCompareFunc T.GCompareDataFunc,
		keyCompareData T.Gpointer) *T.GTree

	TreeNewFull func(keyCompareFunc T.GCompareDataFunc,
		keyCompareData T.Gpointer,
		keyDestroyFunc T.GDestroyNotify,
		valueDestroyFunc T.GDestroyNotify) *T.GTree

	TreeRef func(tree *T.GTree) *T.GTree

	TreeUnref func(tree *T.GTree)

	TreeDestroy func(tree *T.GTree)

	TreeInsert func(tree *T.GTree,
		key T.Gpointer, value T.Gpointer)

	TreeReplace func(tree *T.GTree,
		key T.Gpointer, value T.Gpointer)

	TreeRemove func(tree *T.GTree,
		key T.Gconstpointer) T.Gboolean

	TreeSteal func(tree *T.GTree,
		key T.Gconstpointer) T.Gboolean

	TreeLookup func(tree *T.GTree,
		key T.Gconstpointer) T.Gpointer

	TreeLookupExtended func(tree *T.GTree,
		lookupKey T.Gconstpointer,
		origKey *T.Gpointer,
		value *T.Gpointer) T.Gboolean

	TreeForeach func(tree *T.GTree,
		f T.GTraverseFunc,
		userData T.Gpointer)

	TreeTraverse func(tree *T.GTree,
		traverseFunc T.GTraverseFunc,
		traverseType T.GTraverseType,
		userData T.Gpointer)

	TreeSearch func(tree *T.GTree,
		searchFunc T.GCompareFunc,
		userData T.Gconstpointer) T.Gpointer

	TreeHeight func(tree *T.GTree) int

	TreeNnodes func(tree *T.GTree) int

	UriUnescapeString func(escapedString string,
		illegalCharacters string) string

	UriUnescapeSegment func(escapedString string,
		escapedStringEnd string,
		illegalCharacters string) string

	UriParseScheme func(uri string) string

	UriEscapeString func(unescaped string,
		reservedCharsAllowed string,
		allowUtf8 T.Gboolean) string

	VariantTypeStringIsValid func(
		typeString string) T.Gboolean

	VariantTypeStringScan func(str string,
		limit string,
		endptr **T.Gchar) T.Gboolean

	VariantTypeFree func(typ *T.GVariantType)

	VariantTypeCopy func(typ *T.GVariantType) *T.GVariantType

	VariantTypeNew func(typeString string) *T.GVariantType

	VariantTypeGetStringLength func(typ *T.GVariantType) T.Gsize

	VariantTypePeekString func(typ *T.GVariantType) string

	VariantTypeDupString func(typ *T.GVariantType) string

	VariantTypeIsDefinite func(typ *T.GVariantType) T.Gboolean

	VariantTypeIsContainer func(typ *T.GVariantType) T.Gboolean

	VariantTypeIsBasic func(typ *T.GVariantType) T.Gboolean

	VariantTypeIsMaybe func(typ *T.GVariantType) T.Gboolean

	VariantTypeIsArray func(typ *T.GVariantType) T.Gboolean

	VariantTypeIsTuple func(typ *T.GVariantType) T.Gboolean

	VariantTypeIsDictEntry func(typ *T.GVariantType) T.Gboolean

	VariantTypeIsVariant func(typ *T.GVariantType) T.Gboolean

	VariantTypeHash func(typ T.Gconstpointer) uint

	VariantTypeEqual func(type1 T.Gconstpointer,
		type2 T.Gconstpointer) T.Gboolean

	VariantTypeIsSubtypeOf func(typ *T.GVariantType,
		supertyp *T.GVariantType) T.Gboolean

	VariantTypeElement func(typ *T.GVariantType) *T.GVariantType

	VariantTypeFirst func(typ *T.GVariantType) *T.GVariantType

	VariantTypeNext func(typ *T.GVariantType) *T.GVariantType

	VariantTypeNItems func(typ *T.GVariantType) T.Gsize

	VariantTypeKey func(typ *T.GVariantType) *T.GVariantType

	VariantTypeValue func(typ *T.GVariantType) *T.GVariantType

	VariantTypeNewArray func(element *T.GVariantType) *T.GVariantType

	VariantTypeNewMaybe func(element *T.GVariantType) *T.GVariantType

	VariantTypeNewTuple func(items **T.GVariantType,
		length int) *T.GVariantType

	VariantTypeNewDictEntry func(key *T.GVariantType,
		value *T.GVariantType) *T.GVariantType

	VariantTypeChecked func(string) *T.GVariantType

	VariantUnref func(value *T.GVariant)

	VariantRef func(value *T.GVariant) *T.GVariant

	VariantRefSink func(value *T.GVariant) *T.GVariant

	VariantIsFloating func(value *T.GVariant) T.Gboolean

	VariantGetType func(value *T.GVariant) *T.GVariantType

	VariantGetTypeString func(value *T.GVariant) string

	VariantIsOfType func(value *T.GVariant,
		typ *T.GVariantType) T.Gboolean

	VariantIsContainer func(value *T.GVariant) T.Gboolean

	VariantClassify func(value *T.GVariant) T.GVariantClass

	VariantNewBoolean func(value T.Gboolean) *T.GVariant

	VariantNewByte func(value T.Guchar) *T.GVariant

	VariantNewInt16 func(value int16) *T.GVariant

	VariantNewUint16 func(value uint16) *T.GVariant

	VariantNewInt32 func(value T.GInt32) *T.GVariant

	VariantNewUint32 func(value T.GUint32) *T.GVariant

	VariantNewInt64 func(value int64) *T.GVariant

	VariantNewUint64 func(value uint64) *T.GVariant

	VariantNewHandle func(value T.GInt32) *T.GVariant

	VariantNewDouble func(value float64) *T.GVariant

	VariantNewString func(str string) *T.GVariant

	VariantNewObjectPath func(objectPath string) *T.GVariant

	VariantIsObjectPath func(str string) T.Gboolean

	VariantNewSignature func(signature string) *T.GVariant

	VariantIsSignature func(str string) T.Gboolean

	VariantNewVariant func(value *T.GVariant) *T.GVariant

	VariantNewStrv func(strv **T.Gchar,
		length T.Gssize) *T.GVariant

	VariantNewBytestring func(str string) *T.GVariant

	VariantNewBytestringArray func(strv **T.Gchar,
		length T.Gssize) *T.GVariant

	VariantGetBoolean func(value *T.GVariant) T.Gboolean

	VariantGetByte func(value *T.GVariant) T.Guchar

	VariantGetInt16 func(value *T.GVariant) int16

	VariantGetUint16 func(value *T.GVariant) uint16

	VariantGetInt32 func(value *T.GVariant) T.GInt32

	VariantGetUint32 func(value *T.GVariant) T.GUint32

	VariantGetInt64 func(value *T.GVariant) int64

	VariantGetUint64 func(value *T.GVariant) uint64

	VariantGetHandle func(value *T.GVariant) T.GInt32

	VariantGetDouble func(value *T.GVariant) float64

	VariantGetVariant func(value *T.GVariant) *T.GVariant

	VariantGetString func(value *T.GVariant,
		length *T.Gsize) string

	VariantDupString func(value *T.GVariant,
		length *T.Gsize) string

	VariantGetStrv func(value *T.GVariant,
		length *T.Gsize) **T.Gchar

	VariantDupStrv func(value *T.GVariant,
		length *T.Gsize) **T.Gchar

	VariantGetBytestring func(value *T.GVariant) string

	VariantDupBytestring func(value *T.GVariant,
		length *T.Gsize) string

	VariantGetBytestringArray func(value *T.GVariant,
		length *T.Gsize) **T.Gchar

	VariantDupBytestringArray func(value *T.GVariant,
		length *T.Gsize) **T.Gchar

	VariantNewMaybe func(childType *T.GVariantType,
		child *T.GVariant) *T.GVariant

	VariantNewArray func(childType *T.GVariantType,
		children **T.GVariant,
		nChildren T.Gsize) *T.GVariant

	VariantNewTuple func(children **T.GVariant,
		nChildren T.Gsize) *T.GVariant

	VariantNewDictEntry func(key *T.GVariant,
		value *T.GVariant) *T.GVariant

	VariantGetMaybe func(value *T.GVariant) *T.GVariant

	VariantNChildren func(value *T.GVariant) T.Gsize

	VariantGetChild func(value *T.GVariant, index T.Gsize,
		formatString string, v ...VArg)

	VariantGetChildValue func(value *T.GVariant,
		index T.Gsize) *T.GVariant

	VariantLookup func(dictionary *T.GVariant, key string,
		formatString string, v ...VArg) T.Gboolean

	VariantLookupValue func(dictionary *T.GVariant,
		key string,
		expectedType *T.GVariantType) *T.GVariant

	VariantGetFixedArray func(value *T.GVariant,
		nElements *T.Gsize,
		elementSize T.Gsize) T.Gconstpointer

	VariantGetSize func(value *T.GVariant) T.Gsize

	VariantGetData func(value *T.GVariant) T.Gconstpointer

	VariantStore func(value *T.GVariant,
		data T.Gpointer)

	VariantPrint func(value *T.GVariant,
		typeAnnotate T.Gboolean) string

	VariantPrintString func(value *T.GVariant,
		str *T.GString,
		typeAnnotate T.Gboolean) *T.GString

	VariantHash func(value T.Gconstpointer) uint

	VariantEqual func(one T.Gconstpointer,
		two T.Gconstpointer) T.Gboolean

	VariantGetNormalForm func(value *T.GVariant) *T.GVariant

	VariantIsNormalForm func(value *T.GVariant) T.Gboolean

	VariantByteswap func(value *T.GVariant) *T.GVariant

	VariantNewFromData func(typ *T.GVariantType,
		data T.Gconstpointer,
		size T.Gsize,
		trusted T.Gboolean,
		notify T.GDestroyNotify,
		userData T.Gpointer) *T.GVariant

	VariantIterNew func(value *T.GVariant) *T.GVariantIter

	VariantIterInit func(iter *T.GVariantIter,
		value *T.GVariant) T.Gsize

	VariantIterCopy func(iter *T.GVariantIter) *T.GVariantIter

	VariantIterNChildren func(iter *T.GVariantIter) T.Gsize

	VariantIterFree func(iter *T.GVariantIter)

	VariantIterNextValue func(iter *T.GVariantIter) *T.GVariant

	VariantIterNext func(iter *T.GVariantIter,
		formatString string, v ...VArg) T.Gboolean

	VariantIterLoop func(iter *T.GVariantIter,
		formatString string, v ...VArg) T.Gboolean

	VariantParserGetErrorQuark func() T.GQuark

	VariantBuilderNew func(typ *T.GVariantType) *T.GVariantBuilder

	VariantBuilderUnref func(builder *T.GVariantBuilder)

	VariantBuilderRef func(builder *T.GVariantBuilder) *T.GVariantBuilder

	VariantBuilderInit func(builder *T.GVariantBuilder,
		typ *T.GVariantType)

	VariantBuilderEnd func(builder *T.GVariantBuilder) *T.GVariant

	VariantBuilderClear func(builder *T.GVariantBuilder)

	VariantBuilderOpen func(
		builder *T.GVariantBuilder, typ *T.GVariantType)

	VariantBuilderClose func(builder *T.GVariantBuilder)

	VariantBuilderAddValue func(
		builder *T.GVariantBuilder, value *T.GVariant)

	VariantBuilderAdd func(builder *T.GVariantBuilder,
		formatString string, v ...VArg)

	VariantBuilderAddParsed func(
		builder *T.GVariantBuilder, format string, v ...VArg)

	VariantNew func(
		formatString string, v ...VArg) *T.GVariant

	VariantGet func(
		value *T.GVariant, formatString string, v ...VArg)

	VariantNewVa func(formatString string,
		endptr **T.Gchar, app *T.VaList) *T.GVariant

	VariantGetVa func(value *T.GVariant, formatString string,
		endptr **T.Gchar, app *T.VaList)

	VariantParse func(typ *T.GVariantType, text, limit string,
		endptr **T.Gchar, err **T.GError) *T.GVariant

	VariantNewParsed func(
		format string, v ...VArg) *T.GVariant

	VariantNewParsedVa func(format string,
		app *T.VaList) *T.GVariant

	VariantCompare func(one T.Gconstpointer,
		two T.Gconstpointer) int

	Win32Ftruncate func(f int,
		size uint) int

	Win32Getlocale func() string

	Win32ErrorMessage func(err int) string

	Win32GetPackageInstallationDirectory func(
		pkg, dllName string) string

	Win32GetPackageInstallationSubdirectory func(
		pkg, dllName, subdir string) string

	Win32GetPackageInstallationDirectoryOfModule func(
		hmodule T.Gpointer) string

	Win32GetWindowsVersion func() uint

	Win32LocaleFilenameFromUtf8 func(
		utf8filename string) string

	Access func(filename string, mode int) int

	Chmod func(filename string, mode int) int

	Open func(filename string, flags int, mode int) int

	Creat func(filename string, mode int) int

	Rename func(oldfilename string, newfilename string) int

	Mkdir func(filename string, mode int) int

	Chdir func(path string) int

	Stat func(filename string, buf *GStatBuf) int

	Lstat func(filename string, buf *GStatBuf) int

	Unlink func(filename string) int

	Remove func(filename string) int

	Rmdir func(filename string) int

	Fopen func(filename string, mode string) *T.FILE

	Freopen func(
		filename string, mode string, stream *T.FILE) *T.FILE

	Utime func(filename string, utb *T.Utimbuf) int

	FindProgramInPath func(program string) string

	Printf func(format string, v ...VArg) int

	Fprintf func(file *T.FILE, format string, v ...VArg) int

	Sprintf func(str string, format string, v ...VArg) int

	Vprintf func(format string, args T.VaList) int

	Vfprintf func(
		file *T.FILE, format string, args T.VaList) int

	Vsprintf func(
		str string, format string, args T.VaList) int

	Vasprintf func(
		string **T.Gchar, format string, args T.VaList) int

	GetUserSpecialDir func(directory T.GUserDirectory) string

	KeyFileErrorQuark func() T.GQuark

	MarkupCollectAttributes func(elementName string,
		attributeNames **T.Gchar, attributeValues **T.Gchar,
		error **T.GError, firstType T.GMarkupCollectType,
		firstAttr string, v ...VArg) T.Gboolean

	MarkupErrorQuark func() T.GQuark

	MemSetVtable func(vtable *T.GMemVTable)

	OptionErrorQuark func() T.GQuark

	PrintfStringUpperBound func(
		format string, args T.VaList) T.Gsize

	RegexErrorQuark func() T.GQuark

	StaticPrivateFree func(
		privateKey *T.GStaticPrivate)

	StaticPrivateGet func(
		privateKey *T.GStaticPrivate) T.Gpointer

	StaticPrivateInit func(privateKey *T.GStaticPrivate)

	StaticPrivateSet func(privateKey *T.GStaticPrivate,
		data T.Gpointer, notify T.GDestroyNotify)

	StringChunkInsertConst func(
		chunk *T.GStringChunk, str string) string

	StringDown func(str *T.GString) *T.GString

	StringUp func(str *T.GString) *T.GString

	TrashStackPush func(
		stackP **T.GTrashStack, dataP T.Gpointer)

	TrashStackPop func(stackP **T.GTrashStack) T.Gpointer

	TrashStackPeek func(stackP **T.GTrashStack) T.Gpointer

	TrashStackHeight func(stackP **T.GTrashStack) uint
)

var dll = "libglib-2.0-0.dll"

var apiList = outside.Apis{
	// Undocumented {"_g_debug_flags", &_gDebugFlags},
	// Undocumented {"_g_debug_initialized", &_gDebugInitialized},
	{"g_access", &Access},
	{"g_allocator_free", &AllocatorFree},
	{"g_allocator_new", &AllocatorNew},
	{"g_array_append_vals", &ArrayAppendVals},
	{"g_array_free", &ArrayFree},
	{"g_array_get_element_size", &ArrayGetElementSize},
	{"g_array_insert_vals", &ArrayInsertVals},
	{"g_array_new", &ArrayNew},
	{"g_array_prepend_vals", &ArrayPrependVals},
	{"g_array_ref", &ArrayRef},
	{"g_array_remove_index", &ArrayRemoveIndex},
	{"g_array_remove_index_fast", &ArrayRemoveIndexFast},
	{"g_array_remove_range", &ArrayRemoveRange},
	{"g_array_set_size", &ArraySetSize},
	{"g_array_sized_new", &ArraySizedNew},
	{"g_array_sort", &ArraySort},
	{"g_array_sort_with_data", &ArraySortWithData},
	{"g_array_unref", &ArrayUnref},
	{"g_ascii_digit_value", &AsciiDigitValue},
	{"g_ascii_dtostr", &AsciiDtostr},
	{"g_ascii_formatd", &AsciiFormatd},
	{"g_ascii_strcasecmp", &AsciiStrcasecmp},
	{"g_ascii_strdown", &AsciiStrdown},
	{"g_ascii_strncasecmp", &AsciiStrncasecmp},
	{"g_ascii_strtod", &AsciiStrtod},
	{"g_ascii_strtoll", &AsciiStrtoll},
	{"g_ascii_strtoull", &AsciiStrtoull},
	{"g_ascii_strup", &AsciiStrup},
	{"g_ascii_tolower", &AsciiTolower},
	{"g_ascii_toupper", &AsciiToupper},
	{"g_ascii_xdigit_value", &AsciiXdigitValue},
	{"g_assert_warning", &AssertWarning},
	{"g_assertion_message", &AssertionMessage},
	{"g_assertion_message_cmpnum", &AssertionMessageCmpnum},
	{"g_assertion_message_cmpstr", &AssertionMessageCmpstr},
	{"g_assertion_message_error", &AssertionMessageError},
	{"g_assertion_message_expr", &AssertionMessageExpr},
	{"g_async_queue_length", &AsyncQueueLength},
	{"g_async_queue_length_unlocked", &AsyncQueueLengthUnlocked},
	{"g_async_queue_lock", &AsyncQueueLock},
	{"g_async_queue_new", &AsyncQueueNew},
	{"g_async_queue_new_full", &AsyncQueueNewFull},
	{"g_async_queue_pop", &AsyncQueuePop},
	{"g_async_queue_pop_unlocked", &AsyncQueuePopUnlocked},
	{"g_async_queue_push", &AsyncQueuePush},
	{"g_async_queue_push_sorted", &AsyncQueuePushSorted},
	{"g_async_queue_push_sorted_unlocked", &AsyncQueuePushSortedUnlocked},
	{"g_async_queue_push_unlocked", &AsyncQueuePushUnlocked},
	{"g_async_queue_ref", &AsyncQueueRef},
	{"g_async_queue_ref_unlocked", &AsyncQueueRefUnlocked},
	{"g_async_queue_sort", &AsyncQueueSort},
	{"g_async_queue_sort_unlocked", &AsyncQueueSortUnlocked},
	{"g_async_queue_timed_pop", &AsyncQueueTimedPop},
	{"g_async_queue_timed_pop_unlocked", &AsyncQueueTimedPopUnlocked},
	{"g_async_queue_try_pop", &AsyncQueueTryPop},
	{"g_async_queue_try_pop_unlocked", &AsyncQueueTryPopUnlocked},
	{"g_async_queue_unlock", &AsyncQueueUnlock},
	{"g_async_queue_unref", &AsyncQueueUnref},
	{"g_async_queue_unref_and_unlock", &AsyncQueueUnrefAndUnlock},
	{"g_atexit", &Atexit},
	{"g_atomic_int_add", &AtomicIntAdd},
	{"g_atomic_int_compare_and_exchange", &AtomicIntCompareAndExchange},
	{"g_atomic_int_exchange_and_add", &AtomicIntExchangeAndAdd},
	{"g_atomic_int_get", &AtomicIntGet},
	{"g_atomic_int_set", &AtomicIntSet},
	{"g_atomic_pointer_compare_and_exchange", &AtomicPointerCompareAndExchange},
	{"g_atomic_pointer_get", &AtomicPointerGet},
	{"g_atomic_pointer_set", &AtomicPointerSet},
	{"g_base64_decode", &Base64Decode},
	{"g_base64_decode_inplace", &Base64DecodeInplace},
	{"g_base64_decode_step", &Base64DecodeStep},
	{"g_base64_encode", &Base64Encode},
	{"g_base64_encode_close", &Base64EncodeClose},
	{"g_base64_encode_step", &Base64EncodeStep},
	{"g_basename", &Basename},
	{"g_bit_lock", &BitLock},
	// Inline {"g_bit_nth_lsf", &BitNthLsf},
	// Inline {"g_bit_nth_msf", &BitNthMsf},
	// Inline {"g_bit_storage", &BitStorage},
	{"g_bit_trylock", &BitTrylock},
	{"g_bit_unlock", &BitUnlock},
	{"g_blow_chunks", &BlowChunks},
	{"g_bookmark_file_add_application", &BookmarkFileAddApplication},
	{"g_bookmark_file_add_group", &BookmarkFileAddGroup},
	{"g_bookmark_file_error_quark", &BookmarkFileErrorQuark},
	{"g_bookmark_file_free", &BookmarkFileFree},
	{"g_bookmark_file_get_added", &BookmarkFileGetAdded},
	{"g_bookmark_file_get_app_info", &BookmarkFileGetAppInfo},
	{"g_bookmark_file_get_applications", &BookmarkFileGetApplications},
	{"g_bookmark_file_get_description", &BookmarkFileGetDescription},
	{"g_bookmark_file_get_groups", &BookmarkFileGetGroups},
	{"g_bookmark_file_get_icon", &BookmarkFileGetIcon},
	{"g_bookmark_file_get_is_private", &BookmarkFileGetIsPrivate},
	{"g_bookmark_file_get_mime_type", &BookmarkFileGetMimeType},
	{"g_bookmark_file_get_modified", &BookmarkFileGetModified},
	{"g_bookmark_file_get_size", &BookmarkFileGetSize},
	{"g_bookmark_file_get_title", &BookmarkFileGetTitle},
	{"g_bookmark_file_get_uris", &BookmarkFileGetUris},
	{"g_bookmark_file_get_visited", &BookmarkFileGetVisited},
	{"g_bookmark_file_has_application", &BookmarkFileHasApplication},
	{"g_bookmark_file_has_group", &BookmarkFileHasGroup},
	{"g_bookmark_file_has_item", &BookmarkFileHasItem},
	{"g_bookmark_file_load_from_data", &BookmarkFileLoadFromData},
	{"g_bookmark_file_load_from_data_dirs", &BookmarkFileLoadFromDataDirs},
	{"g_bookmark_file_load_from_file", &BookmarkFileLoadFromFile},
	{"g_bookmark_file_move_item", &BookmarkFileMoveItem},
	{"g_bookmark_file_new", &BookmarkFileNew},
	{"g_bookmark_file_remove_application", &BookmarkFileRemoveApplication},
	{"g_bookmark_file_remove_group", &BookmarkFileRemoveGroup},
	{"g_bookmark_file_remove_item", &BookmarkFileRemoveItem},
	{"g_bookmark_file_set_added", &BookmarkFileSetAdded},
	{"g_bookmark_file_set_app_info", &BookmarkFileSetAppInfo},
	{"g_bookmark_file_set_description", &BookmarkFileSetDescription},
	{"g_bookmark_file_set_groups", &BookmarkFileSetGroups},
	{"g_bookmark_file_set_icon", &BookmarkFileSetIcon},
	{"g_bookmark_file_set_is_private", &BookmarkFileSetIsPrivate},
	{"g_bookmark_file_set_mime_type", &BookmarkFileSetMimeType},
	{"g_bookmark_file_set_modified", &BookmarkFileSetModified},
	{"g_bookmark_file_set_title", &BookmarkFileSetTitle},
	{"g_bookmark_file_set_visited", &BookmarkFileSetVisited},
	{"g_bookmark_file_to_data", &BookmarkFileToData},
	{"g_bookmark_file_to_file", &BookmarkFileToFile},
	{"g_build_filename", &BuildFilename},
	{"g_build_filenamev", &BuildFilenamev},
	{"g_build_path", &BuildPath},
	{"g_build_pathv", &BuildPathv},
	{"g_byte_array_append", &ByteArrayAppend},
	{"g_byte_array_free", &ByteArrayFree},
	{"g_byte_array_new", &ByteArrayNew},
	{"g_byte_array_prepend", &ByteArrayPrepend},
	{"g_byte_array_ref", &ByteArrayRef},
	{"g_byte_array_remove_index", &ByteArrayRemoveIndex},
	{"g_byte_array_remove_index_fast", &ByteArrayRemoveIndexFast},
	{"g_byte_array_remove_range", &ByteArrayRemoveRange},
	{"g_byte_array_set_size", &ByteArraySetSize},
	{"g_byte_array_sized_new", &ByteArraySizedNew},
	{"g_byte_array_sort", &ByteArraySort},
	{"g_byte_array_sort_with_data", &ByteArraySortWithData},
	{"g_byte_array_unref", &ByteArrayUnref},
	{"g_cache_destroy", &CacheDestroy},
	{"g_cache_insert", &CacheInsert},
	{"g_cache_key_foreach", &CacheKeyForeach},
	{"g_cache_new", &CacheNew},
	{"g_cache_remove", &CacheRemove},
	{"g_cache_value_foreach", &CacheValueForeach},
	{"g_chdir", &Chdir},
	{"g_checksum_copy", &ChecksumCopy},
	{"g_checksum_free", &ChecksumFree},
	{"g_checksum_get_digest", &ChecksumGetDigest},
	{"g_checksum_get_string", &ChecksumGetString},
	{"g_checksum_new", &ChecksumNew},
	{"g_checksum_reset", &ChecksumReset},
	{"g_checksum_type_get_length", &ChecksumTypeGetLength},
	{"g_checksum_update", &ChecksumUpdate},
	{"g_child_watch_add", &ChildWatchAdd},
	{"g_child_watch_add_full", &ChildWatchAddFull},
	{"g_child_watch_source_new", &ChildWatchSourceNew},
	{"g_chmod", &Chmod},
	{"g_clear_error", &ClearError},
	{"g_completion_add_items", &CompletionAddItems},
	{"g_completion_clear_items", &CompletionClearItems},
	{"g_completion_complete", &CompletionComplete},
	{"g_completion_complete_utf8", &CompletionCompleteUtf8},
	{"g_completion_free", &CompletionFree},
	{"g_completion_new", &CompletionNew},
	{"g_completion_remove_items", &CompletionRemoveItems},
	{"g_completion_set_compare", &CompletionSetCompare},
	{"g_compute_checksum_for_data", &ComputeChecksumForData},
	{"g_compute_checksum_for_string", &ComputeChecksumForString},
	{"g_convert", &Convert},
	{"g_convert_error_quark", &ConvertErrorQuark},
	{"g_convert_with_fallback", &ConvertWithFallback},
	{"g_convert_with_iconv", &ConvertWithIconv},
	{"g_creat", &Creat},
	{"g_datalist_clear", &DatalistClear},
	{"g_datalist_foreach", &DatalistForeach},
	{"g_datalist_get_flags", &DatalistGetFlags},
	{"g_datalist_id_get_data", &DatalistIdGetData},
	{"g_datalist_id_remove_no_notify", &DatalistIdRemoveNoNotify},
	{"g_datalist_id_set_data_full", &DatalistIdSetDataFull},
	{"g_datalist_init", &DatalistInit},
	{"g_datalist_set_flags", &DatalistSetFlags},
	{"g_datalist_unset_flags", &DatalistUnsetFlags},
	{"g_dataset_destroy", &DatasetDestroy},
	{"g_dataset_foreach", &DatasetForeach},
	{"g_dataset_id_get_data", &DatasetIdGetData},
	{"g_dataset_id_remove_no_notify", &DatasetIdRemoveNoNotify},
	{"g_dataset_id_set_data_full", &DatasetIdSetDataFull},
	{"g_date_add_days", &DateAddDays},
	{"g_date_add_months", &DateAddMonths},
	{"g_date_add_years", &DateAddYears},
	{"g_date_clamp", &DateClamp},
	{"g_date_clear", &DateClear},
	{"g_date_compare", &DateCompare},
	{"g_date_days_between", &DateDaysBetween},
	{"g_date_free", &DateFree},
	{"g_date_get_day", &DateGetDay},
	{"g_date_get_day_of_year", &DateGetDayOfYear},
	{"g_date_get_days_in_month", &DateGetDaysInMonth},
	{"g_date_get_iso8601_week_of_year", &DateGetIso8601WeekOfYear},
	{"g_date_get_julian", &DateGetJulian},
	{"g_date_get_monday_week_of_year", &DateGetMondayWeekOfYear},
	{"g_date_get_monday_weeks_in_year", &DateGetMondayWeeksInYear},
	{"g_date_get_month", &DateGetMonth},
	{"g_date_get_sunday_week_of_year", &DateGetSundayWeekOfYear},
	{"g_date_get_sunday_weeks_in_year", &DateGetSundayWeeksInYear},
	{"g_date_get_weekday", &DateGetWeekday},
	{"g_date_get_year", &DateGetYear},
	{"g_date_is_first_of_month", &DateIsFirstOfMonth},
	{"g_date_is_last_of_month", &DateIsLastOfMonth},
	{"g_date_is_leap_year", &DateIsLeapYear},
	{"g_date_new", &DateNew},
	{"g_date_new_dmy", &DateNewDmy},
	{"g_date_new_julian", &DateNewJulian},
	{"g_date_order", &DateOrder},
	{"g_date_set_day", &DateSetDay},
	{"g_date_set_dmy", &DateSetDmy},
	{"g_date_set_julian", &DateSetJulian},
	{"g_date_set_month", &DateSetMonth},
	{"g_date_set_parse", &DateSetParse},
	{"g_date_set_time", &DateSetTime},
	{"g_date_set_time_t", &DateSetTimeT},
	{"g_date_set_time_val", &DateSetTimeVal},
	{"g_date_set_year", &DateSetYear},
	{"g_date_strftime", &DateStrftime},
	{"g_date_subtract_days", &DateSubtractDays},
	{"g_date_subtract_months", &DateSubtractMonths},
	{"g_date_subtract_years", &DateSubtractYears},
	{"g_date_time_add", &DateTimeAdd},
	{"g_date_time_add_days", &DateTimeAddDays},
	{"g_date_time_add_full", &DateTimeAddFull},
	{"g_date_time_add_hours", &DateTimeAddHours},
	{"g_date_time_add_minutes", &DateTimeAddMinutes},
	{"g_date_time_add_months", &DateTimeAddMonths},
	{"g_date_time_add_seconds", &DateTimeAddSeconds},
	{"g_date_time_add_weeks", &DateTimeAddWeeks},
	{"g_date_time_add_years", &DateTimeAddYears},
	{"g_date_time_compare", &DateTimeCompare},
	{"g_date_time_difference", &DateTimeDifference},
	{"g_date_time_equal", &DateTimeEqual},
	{"g_date_time_format", &DateTimeFormat},
	{"g_date_time_get_day_of_month", &DateTimeGetDayOfMonth},
	{"g_date_time_get_day_of_week", &DateTimeGetDayOfWeek},
	{"g_date_time_get_day_of_year", &DateTimeGetDayOfYear},
	{"g_date_time_get_hour", &DateTimeGetHour},
	{"g_date_time_get_microsecond", &DateTimeGetMicrosecond},
	{"g_date_time_get_minute", &DateTimeGetMinute},
	{"g_date_time_get_month", &DateTimeGetMonth},
	{"g_date_time_get_second", &DateTimeGetSecond},
	{"g_date_time_get_seconds", &DateTimeGetSeconds},
	{"g_date_time_get_timezone_abbreviation", &DateTimeGetTimezoneAbbreviation},
	{"g_date_time_get_utc_offset", &DateTimeGetUtcOffset},
	{"g_date_time_get_week_numbering_year", &DateTimeGetWeekNumberingYear},
	{"g_date_time_get_week_of_year", &DateTimeGetWeekOfYear},
	{"g_date_time_get_year", &DateTimeGetYear},
	{"g_date_time_get_ymd", &DateTimeGetYmd},
	{"g_date_time_hash", &DateTimeHash},
	{"g_date_time_is_daylight_savings", &DateTimeIsDaylightSavings},
	{"g_date_time_new", &DateTimeNew},
	{"g_date_time_new_from_timeval_local", &DateTimeNewFromTimevalLocal},
	{"g_date_time_new_from_timeval_utc", &DateTimeNewFromTimevalUtc},
	{"g_date_time_new_from_unix_local", &DateTimeNewFromUnixLocal},
	{"g_date_time_new_from_unix_utc", &DateTimeNewFromUnixUtc},
	{"g_date_time_new_local", &DateTimeNewLocal},
	{"g_date_time_new_now", &DateTimeNewNow},
	{"g_date_time_new_now_local", &DateTimeNewNowLocal},
	{"g_date_time_new_now_utc", &DateTimeNewNowUtc},
	{"g_date_time_new_utc", &DateTimeNewUtc},
	{"g_date_time_ref", &DateTimeRef},
	{"g_date_time_to_local", &DateTimeToLocal},
	{"g_date_time_to_timeval", &DateTimeToTimeval},
	{"g_date_time_to_timezone", &DateTimeToTimezone},
	{"g_date_time_to_unix", &DateTimeToUnix},
	{"g_date_time_to_utc", &DateTimeToUtc},
	{"g_date_time_unref", &DateTimeUnref},
	{"g_date_to_struct_tm", &DateToStructTm},
	{"g_date_valid", &DateValid},
	{"g_date_valid_day", &DateValidDay},
	{"g_date_valid_dmy", &DateValidDmy},
	{"g_date_valid_julian", &DateValidJulian},
	{"g_date_valid_month", &DateValidMonth},
	{"g_date_valid_weekday", &DateValidWeekday},
	{"g_date_valid_year", &DateValidYear},
	{"g_dcgettext", &Dcgettext},
	{"g_dgettext", &Dgettext},
	{"g_dir_close", &DirClose},
	// Windows: _utf8 {"g_dir_open", &DirOpen},
	{"g_dir_open_utf8", &DirOpenUtf8},
	// Windows: _utf8 {"g_dir_read_name", &DirReadName},
	{"g_dir_read_name_utf8", &DirReadNameUtf8},
	{"g_dir_rewind", &DirRewind},
	{"g_direct_equal", &DirectEqual},
	{"g_direct_hash", &DirectHash},
	{"g_dngettext", &Dngettext},
	{"g_double_equal", &DoubleEqual},
	{"g_double_hash", &DoubleHash},
	{"g_dpgettext", &Dpgettext},
	{"g_dpgettext2", &Dpgettext2},
	{"g_error_copy", &ErrorCopy},
	{"g_error_free", &ErrorFree},
	{"g_error_matches", &ErrorMatches},
	{"g_error_new", &ErrorNew},
	{"g_error_new_literal", &ErrorNewLiteral},
	{"g_error_new_valist", &ErrorNewValist},
	{"g_file_error_from_errno", &FileErrorFromErrno},
	{"g_file_error_quark", &FileErrorQuark},
	// Windows: _utf8 {"g_file_get_contents", &FileGetContents},
	{"g_file_get_contents_utf8", &FileGetContents},
	// Windows: _utf8 {"g_file_open_tmp", &FileOpenTmp},
	{"g_file_open_tmp_utf8", &FileOpenTmp},
	{"g_file_read_link", &FileReadLink},
	{"g_file_set_contents", &FileSetContents},
	// Windows: _utf8 {"g_file_test", &FileTest},
	{"g_file_test_utf8", &FileTest},
	{"g_filename_display_basename", &FilenameDisplayBasename},
	{"g_filename_display_name", &FilenameDisplayName},
	// Windows: _utf8 {"g_filename_from_uri", &FilenameFromUri},
	{"g_filename_from_uri_utf8", &FilenameFromUri},
	// Windows: _utf8 {"g_filename_from_utf8", &FilenameFromUtf8},
	{"g_filename_from_utf8_utf8", &FilenameFromUtf8},
	// Windows: _utf8 {"g_filename_to_uri", &FilenameToUri},
	{"g_filename_to_uri_utf8", &FilenameToUri},
	// Windows: _utf8 {"g_filename_to_utf8", &FilenameToUtf8},
	{"g_filename_to_utf8_utf8", &FilenameToUtf8},
	// Windows: _utf8 {"g_find_program_in_path", &FindProgramInPath},
	{"g_find_program_in_path_utf8", &FindProgramInPath},
	{"g_fopen", &Fopen},
	{"g_format_size_for_display", &FormatSizeForDisplay},
	{"g_fprintf", &Fprintf},
	{"g_free", &Free},
	{"g_freopen", &Freopen},
	{"g_get_application_name", &GetApplicationName},
	{"g_get_charset", &GetCharset},
	// Undocumented {"g_get_codeset", &GetCodeset},
	// Windows: _utf8 {"g_get_current_dir", &GetCurrentDir},
	{"g_get_current_dir_utf8", &GetCurrentDir},
	{"g_get_current_time", &GetCurrentTime},
	{"g_get_environ", &GetEnviron},
	{"g_get_filename_charsets", &GetFilenameCharsets},
	// Windows: _utf8 {"g_get_home_dir", &GetHomeDir},
	{"g_get_home_dir_utf8", &GetHomeDir},
	{"g_get_host_name", &GetHostName},
	{"g_get_language_names", &GetLanguageNames},
	{"g_get_locale_variants", &GetLocaleVariants},
	{"g_get_monotonic_time", &GetMonotonicTime},
	{"g_get_prgname", &GetPrgname},
	// Windows: _utf8 {"g_get_real_name", &GetRealName},
	{"g_get_real_name_utf8", &GetRealName},
	{"g_get_real_time", &GetRealTime},
	{"g_get_system_config_dirs", &GetSystemConfigDirs},
	{"g_get_system_data_dirs", &GetSystemDataDirs},
	// Windows: _utf8 {"g_get_tmp_dir", &GetTmpDir},
	{"g_get_tmp_dir_utf8", &GetTmpDir},
	{"g_get_user_cache_dir", &GetUserCacheDir},
	{"g_get_user_config_dir", &GetUserConfigDir},
	{"g_get_user_data_dir", &GetUserDataDir},
	// Windows: _utf8 {"g_get_user_name", &GetUserName},
	{"g_get_user_name_utf8", &GetUserName},
	{"g_get_user_runtime_dir", &GetUserRuntimeDir},
	{"g_get_user_special_dir", &GetUserSpecialDir},
	// Windows: _utf8 {"g_getenv", &Getenv},
	{"g_getenv_utf8", &Getenv},
	{"g_hash_table_destroy", &HashTableDestroy},
	{"g_hash_table_find", &HashTableFind},
	{"g_hash_table_foreach", &HashTableForeach},
	{"g_hash_table_foreach_remove", &HashTableForeachRemove},
	{"g_hash_table_foreach_steal", &HashTableForeachSteal},
	{"g_hash_table_get_keys", &HashTableGetKeys},
	{"g_hash_table_get_values", &HashTableGetValues},
	{"g_hash_table_insert", &HashTableInsert},
	{"g_hash_table_iter_get_hash_table", &HashTableIterGetHashTable},
	{"g_hash_table_iter_init", &HashTableIterInit},
	{"g_hash_table_iter_next", &HashTableIterNext},
	{"g_hash_table_iter_remove", &HashTableIterRemove},
	{"g_hash_table_iter_steal", &HashTableIterSteal},
	{"g_hash_table_lookup", &HashTableLookup},
	{"g_hash_table_lookup_extended", &HashTableLookupExtended},
	{"g_hash_table_new", &HashTableNew},
	{"g_hash_table_new_full", &HashTableNewFull},
	{"g_hash_table_ref", &HashTableRef},
	{"g_hash_table_remove", &HashTableRemove},
	{"g_hash_table_remove_all", &HashTableRemoveAll},
	{"g_hash_table_replace", &HashTableReplace},
	{"g_hash_table_size", &HashTableSize},
	{"g_hash_table_steal", &HashTableSteal},
	{"g_hash_table_steal_all", &HashTableStealAll},
	{"g_hash_table_unref", &HashTableUnref},
	{"g_hook_alloc", &HookAlloc},
	{"g_hook_compare_ids", &HookCompareIds},
	{"g_hook_destroy", &HookDestroy},
	{"g_hook_destroy_link", &HookDestroyLink},
	{"g_hook_find", &HookFind},
	{"g_hook_find_data", &HookFindData},
	{"g_hook_find_func", &HookFindFunc},
	{"g_hook_find_func_data", &HookFindFuncData},
	{"g_hook_first_valid", &HookFirstValid},
	{"g_hook_free", &HookFree},
	{"g_hook_get", &HookGet},
	{"g_hook_insert_before", &HookInsertBefore},
	{"g_hook_insert_sorted", &HookInsertSorted},
	{"g_hook_list_clear", &HookListClear},
	{"g_hook_list_init", &HookListInit},
	{"g_hook_list_invoke", &HookListInvoke},
	{"g_hook_list_invoke_check", &HookListInvokeCheck},
	{"g_hook_list_marshal", &HookListMarshal},
	{"g_hook_list_marshal_check", &HookListMarshalCheck},
	{"g_hook_next_valid", &HookNextValid},
	{"g_hook_prepend", &HookPrepend},
	{"g_hook_ref", &HookRef},
	{"g_hook_unref", &HookUnref},
	{"g_hostname_is_ascii_encoded", &HostnameIsAsciiEncoded},
	{"g_hostname_is_ip_address", &HostnameIsIpAddress},
	{"g_hostname_is_non_ascii", &HostnameIsNonAscii},
	{"g_hostname_to_ascii", &HostnameToAscii},
	{"g_hostname_to_unicode", &HostnameToUnicode},
	{"g_iconv", &Iconv},
	{"g_iconv_close", &IconvClose},
	{"g_iconv_open", &IconvOpen},
	{"g_idle_add", &IdleAdd},
	{"g_idle_add_full", &IdleAddFull},
	{"g_idle_remove_by_data", &IdleRemoveByData},
	{"g_idle_source_new", &IdleSourceNew},
	{"g_int64_equal", &Int64Equal},
	{"g_int64_hash", &Int64Hash},
	{"g_int_equal", &IntEqual},
	{"g_int_hash", &IntHash},
	{"g_intern_static_string", &InternStaticString},
	{"g_intern_string", &InternString},
	{"g_io_add_watch", &IoAddWatch},
	{"g_io_add_watch_full", &IoAddWatchFull},
	{"g_io_channel_close", &IoChannelClose},
	{"g_io_channel_error_from_errno", &IoChannelErrorFromErrno},
	{"g_io_channel_error_quark", &IoChannelErrorQuark},
	{"g_io_channel_flush", &IoChannelFlush},
	{"g_io_channel_get_buffer_condition", &IoChannelGetBufferCondition},
	{"g_io_channel_get_buffer_size", &IoChannelGetBufferSize},
	{"g_io_channel_get_buffered", &IoChannelGetBuffered},
	{"g_io_channel_get_close_on_unref", &IoChannelGetCloseOnUnref},
	{"g_io_channel_get_encoding", &IoChannelGetEncoding},
	{"g_io_channel_get_flags", &IoChannelGetFlags},
	{"g_io_channel_get_line_term", &IoChannelGetLineTerm},
	{"g_io_channel_init", &IoChannelInit},
	// Windows: _utf8 {"g_io_channel_new_file", &IoChannelNewFile},
	{"g_io_channel_new_file_utf8", &IoChannelNewFile},
	{"g_io_channel_read", &IoChannelRead},
	{"g_io_channel_read_chars", &IoChannelReadChars},
	{"g_io_channel_read_line", &IoChannelReadLine},
	{"g_io_channel_read_line_string", &IoChannelReadLineString},
	{"g_io_channel_read_to_end", &IoChannelReadToEnd},
	{"g_io_channel_read_unichar", &IoChannelReadUnichar},
	{"g_io_channel_ref", &IoChannelRef},
	{"g_io_channel_seek", &IoChannelSeek},
	{"g_io_channel_seek_position", &IoChannelSeekPosition},
	{"g_io_channel_set_buffer_size", &IoChannelSetBufferSize},
	{"g_io_channel_set_buffered", &IoChannelSetBuffered},
	{"g_io_channel_set_close_on_unref", &IoChannelSetCloseOnUnref},
	{"g_io_channel_set_encoding", &IoChannelSetEncoding},
	{"g_io_channel_set_flags", &IoChannelSetFlags},
	{"g_io_channel_set_line_term", &IoChannelSetLineTerm},
	{"g_io_channel_shutdown", &IoChannelShutdown},
	{"g_io_channel_unix_get_fd", &IoChannelUnixGetFd},
	{"g_io_channel_unix_new", &IoChannelUnixNew},
	{"g_io_channel_unref", &IoChannelUnref},
	{"g_io_channel_win32_get_fd", &IoChannelWin32GetFd},
	{"g_io_channel_win32_make_pollfd", &IoChannelWin32MakePollfd},
	{"g_io_channel_win32_new_fd", &IoChannelWin32NewFd},
	{"g_io_channel_win32_new_messages", &IoChannelWin32NewMessages},
	{"g_io_channel_win32_new_socket", &IoChannelWin32NewSocket},
	// Undocumented {"g_io_channel_win32_new_stream_socket", &IoChannelWin32NewStreamSocket},
	{"g_io_channel_win32_poll", &IoChannelWin32Poll},
	// Undocumented {"g_io_channel_win32_set_debug", &IoChannelWin32SetDebug},
	{"g_io_channel_write", &IoChannelWrite},
	{"g_io_channel_write_chars", &IoChannelWriteChars},
	{"g_io_channel_write_unichar", &IoChannelWriteUnichar},
	{"g_io_create_watch", &IoCreateWatch},
	{"g_key_file_error_quark", &KeyFileErrorQuark},
	{"g_key_file_free", &KeyFileFree},
	{"g_key_file_get_boolean", &KeyFileGetBoolean},
	{"g_key_file_get_boolean_list", &KeyFileGetBooleanList},
	{"g_key_file_get_comment", &KeyFileGetComment},
	{"g_key_file_get_double", &KeyFileGetDouble},
	{"g_key_file_get_double_list", &KeyFileGetDoubleList},
	{"g_key_file_get_groups", &KeyFileGetGroups},
	{"g_key_file_get_int64", &KeyFileGetInt64},
	{"g_key_file_get_integer", &KeyFileGetInteger},
	{"g_key_file_get_integer_list", &KeyFileGetIntegerList},
	{"g_key_file_get_keys", &KeyFileGetKeys},
	{"g_key_file_get_locale_string", &KeyFileGetLocaleString},
	{"g_key_file_get_locale_string_list", &KeyFileGetLocaleStringList},
	{"g_key_file_get_start_group", &KeyFileGetStartGroup},
	{"g_key_file_get_string", &KeyFileGetString},
	{"g_key_file_get_string_list", &KeyFileGetStringList},
	{"g_key_file_get_uint64", &KeyFileGetUint64},
	{"g_key_file_get_value", &KeyFileGetValue},
	{"g_key_file_has_group", &KeyFileHasGroup},
	{"g_key_file_has_key", &KeyFileHasKey},
	{"g_key_file_load_from_data", &KeyFileLoadFromData},
	{"g_key_file_load_from_data_dirs", &KeyFileLoadFromDataDirs},
	{"g_key_file_load_from_dirs", &KeyFileLoadFromDirs},
	{"g_key_file_load_from_file", &KeyFileLoadFromFile},
	{"g_key_file_new", &KeyFileNew},
	{"g_key_file_remove_comment", &KeyFileRemoveComment},
	{"g_key_file_remove_group", &KeyFileRemoveGroup},
	{"g_key_file_remove_key", &KeyFileRemoveKey},
	{"g_key_file_set_boolean", &KeyFileSetBoolean},
	{"g_key_file_set_boolean_list", &KeyFileSetBooleanList},
	{"g_key_file_set_comment", &KeyFileSetComment},
	{"g_key_file_set_double", &KeyFileSetDouble},
	{"g_key_file_set_double_list", &KeyFileSetDoubleList},
	{"g_key_file_set_int64", &KeyFileSetInt64},
	{"g_key_file_set_integer", &KeyFileSetInteger},
	{"g_key_file_set_integer_list", &KeyFileSetIntegerList},
	{"g_key_file_set_list_separator", &KeyFileSetListSeparator},
	{"g_key_file_set_locale_string", &KeyFileSetLocaleString},
	{"g_key_file_set_locale_string_list", &KeyFileSetLocaleStringList},
	{"g_key_file_set_string", &KeyFileSetString},
	{"g_key_file_set_string_list", &KeyFileSetStringList},
	{"g_key_file_set_uint64", &KeyFileSetUint64},
	{"g_key_file_set_value", &KeyFileSetValue},
	{"g_key_file_to_data", &KeyFileToData},
	{"g_list_alloc", &ListAlloc},
	{"g_list_append", &ListAppend},
	{"g_list_concat", &ListConcat},
	{"g_list_copy", &ListCopy},
	{"g_list_delete_link", &ListDeleteLink},
	{"g_list_find", &ListFind},
	{"g_list_find_custom", &ListFindCustom},
	{"g_list_first", &ListFirst},
	{"g_list_foreach", &ListForeach},
	{"g_list_free", &ListFree},
	{"g_list_free_1", &ListFree1},
	{"g_list_free_full", &ListFreeFull},
	{"g_list_index", &ListIndex},
	{"g_list_insert", &ListInsert},
	{"g_list_insert_before", &ListInsertBefore},
	{"g_list_insert_sorted", &ListInsertSorted},
	{"g_list_insert_sorted_with_data", &ListInsertSortedWithData},
	{"g_list_last", &ListLast},
	{"g_list_length", &ListLength},
	{"g_list_nth", &ListNth},
	{"g_list_nth_data", &ListNthData},
	{"g_list_nth_prev", &ListNthPrev},
	{"g_list_pop_allocator", &ListPopAllocator},
	{"g_list_position", &ListPosition},
	{"g_list_prepend", &ListPrepend},
	{"g_list_push_allocator", &ListPushAllocator},
	{"g_list_remove", &ListRemove},
	{"g_list_remove_all", &ListRemoveAll},
	{"g_list_remove_link", &ListRemoveLink},
	{"g_list_reverse", &ListReverse},
	{"g_list_sort", &ListSort},
	{"g_list_sort_with_data", &ListSortWithData},
	{"g_listenv", &Listenv},
	{"g_locale_from_utf8", &LocaleFromUtf8},
	{"g_locale_to_utf8", &LocaleToUtf8},
	{"g_log", &Log},
	{"g_log_default_handler", &LogDefaultHandler},
	{"g_log_remove_handler", &LogRemoveHandler},
	{"g_log_set_always_fatal", &LogSetAlwaysFatal},
	{"g_log_set_default_handler", &LogSetDefaultHandler},
	{"g_log_set_fatal_mask", &LogSetFatalMask},
	{"g_log_set_handler", &LogSetHandler},
	{"g_logv", &Logv},
	{"g_lstat", &Lstat},
	{"g_main_context_acquire", &MainContextAcquire},
	{"g_main_context_add_poll", &MainContextAddPoll},
	{"g_main_context_check", &MainContextCheck},
	{"g_main_context_default", &MainContextDefault},
	{"g_main_context_dispatch", &MainContextDispatch},
	{"g_main_context_find_source_by_funcs_user_data", &MainContextFindSourceByFuncsUserData},
	{"g_main_context_find_source_by_id", &MainContextFindSourceById},
	{"g_main_context_find_source_by_user_data", &MainContextFindSourceByUserData},
	{"g_main_context_get_poll_func", &MainContextGetPollFunc},
	{"g_main_context_get_thread_default", &MainContextGetThreadDefault},
	{"g_main_context_invoke", &MainContextInvoke},
	{"g_main_context_invoke_full", &MainContextInvokeFull},
	{"g_main_context_is_owner", &MainContextIsOwner},
	{"g_main_context_iteration", &MainContextIteration},
	{"g_main_context_new", &MainContextNew},
	{"g_main_context_pending", &MainContextPending},
	{"g_main_context_pop_thread_default", &MainContextPopThreadDefault},
	{"g_main_context_prepare", &MainContextPrepare},
	{"g_main_context_push_thread_default", &MainContextPushThreadDefault},
	{"g_main_context_query", &MainContextQuery},
	{"g_main_context_ref", &MainContextRef},
	{"g_main_context_release", &MainContextRelease},
	{"g_main_context_remove_poll", &MainContextRemovePoll},
	{"g_main_context_set_poll_func", &MainContextSetPollFunc},
	{"g_main_context_unref", &MainContextUnref},
	{"g_main_context_wait", &MainContextWait},
	{"g_main_context_wakeup", &MainContextWakeup},
	{"g_main_current_source", &MainCurrentSource},
	{"g_main_depth", &MainDepth},
	{"g_main_loop_get_context", &MainLoopGetContext},
	{"g_main_loop_is_running", &MainLoopIsRunning},
	{"g_main_loop_new", &MainLoopNew},
	{"g_main_loop_quit", &MainLoopQuit},
	{"g_main_loop_ref", &MainLoopRef},
	{"g_main_loop_run", &MainLoopRun},
	{"g_main_loop_unref", &MainLoopUnref},
	{"g_malloc", &Malloc},
	{"g_malloc0", &Malloc0},
	{"g_malloc0_n", &Malloc0N},
	{"g_malloc_n", &MallocN},
	{"g_mapped_file_free", &MappedFileFree},
	{"g_mapped_file_get_contents", &MappedFileGetContents},
	{"g_mapped_file_get_length", &MappedFileGetLength},
	{"g_mapped_file_new", &MappedFileNew},
	{"g_mapped_file_ref", &MappedFileRef},
	{"g_mapped_file_unref", &MappedFileUnref},
	{"g_markup_collect_attributes", &MarkupCollectAttributes},
	{"g_markup_error_quark", &MarkupErrorQuark},
	{"g_markup_escape_text", &MarkupEscapeText},
	{"g_markup_parse_context_end_parse", &MarkupParseContextEndParse},
	{"g_markup_parse_context_free", &MarkupParseContextFree},
	{"g_markup_parse_context_get_element", &MarkupParseContextGetElement},
	{"g_markup_parse_context_get_element_stack", &MarkupParseContextGetElementStack},
	{"g_markup_parse_context_get_position", &MarkupParseContextGetPosition},
	{"g_markup_parse_context_get_user_data", &MarkupParseContextGetUserData},
	{"g_markup_parse_context_new", &MarkupParseContextNew},
	{"g_markup_parse_context_parse", &MarkupParseContextParse},
	{"g_markup_parse_context_pop", &MarkupParseContextPop},
	{"g_markup_parse_context_push", &MarkupParseContextPush},
	{"g_markup_printf_escaped", &MarkupPrintfEscaped},
	{"g_markup_vprintf_escaped", &MarkupVprintfEscaped},
	{"g_match_info_expand_references", &MatchInfoExpandReferences},
	{"g_match_info_fetch", &MatchInfoFetch},
	{"g_match_info_fetch_all", &MatchInfoFetchAll},
	{"g_match_info_fetch_named", &MatchInfoFetchNamed},
	{"g_match_info_fetch_named_pos", &MatchInfoFetchNamedPos},
	{"g_match_info_fetch_pos", &MatchInfoFetchPos},
	{"g_match_info_free", &MatchInfoFree},
	{"g_match_info_get_match_count", &MatchInfoGetMatchCount},
	{"g_match_info_get_regex", &MatchInfoGetRegex},
	{"g_match_info_get_string", &MatchInfoGetString},
	{"g_match_info_is_partial_match", &MatchInfoIsPartialMatch},
	{"g_match_info_matches", &MatchInfoMatches},
	{"g_match_info_next", &MatchInfoNext},
	{"g_mem_chunk_alloc", &MemChunkAlloc},
	{"g_mem_chunk_alloc0", &MemChunkAlloc0},
	{"g_mem_chunk_clean", &MemChunkClean},
	{"g_mem_chunk_destroy", &MemChunkDestroy},
	{"g_mem_chunk_free", &MemChunkFree},
	{"g_mem_chunk_info", &MemChunkInfo},
	{"g_mem_chunk_new", &MemChunkNew},
	{"g_mem_chunk_print", &MemChunkPrint},
	{"g_mem_chunk_reset", &MemChunkReset},
	{"g_mem_is_system_malloc", &MemIsSystemMalloc},
	{"g_mem_profile", &MemProfile},
	{"g_mem_set_vtable", &MemSetVtable},
	{"g_memdup", &Memdup},
	{"g_mkdir", &Mkdir},
	{"g_mkdir_with_parents", &MkdirWithParents},
	// Windows: _utf8 {"g_mkstemp", &Mkstemp},
	{"g_mkstemp_full", &MkstempFull},
	{"g_mkstemp_utf8", &Mkstemp},
	{"g_node_child_index", &NodeChildIndex},
	{"g_node_child_position", &NodeChildPosition},
	{"g_node_children_foreach", &NodeChildrenForeach},
	{"g_node_copy", &NodeCopy},
	{"g_node_copy_deep", &NodeCopyDeep},
	{"g_node_depth", &NodeDepth},
	{"g_node_destroy", &NodeDestroy},
	{"g_node_find", &NodeFind},
	{"g_node_find_child", &NodeFindChild},
	{"g_node_first_sibling", &NodeFirstSibling},
	{"g_node_get_root", &NodeGetRoot},
	{"g_node_insert", &NodeInsert},
	{"g_node_insert_after", &NodeInsertAfter},
	{"g_node_insert_before", &NodeInsertBefore},
	{"g_node_is_ancestor", &NodeIsAncestor},
	{"g_node_last_child", &NodeLastChild},
	{"g_node_last_sibling", &NodeLastSibling},
	{"g_node_max_height", &NodeMaxHeight},
	{"g_node_n_children", &NodeNChildren},
	{"g_node_n_nodes", &NodeNNodes},
	{"g_node_new", &NodeNew},
	{"g_node_nth_child", &NodeNthChild},
	{"g_node_pop_allocator", &NodePopAllocator},
	{"g_node_prepend", &NodePrepend},
	{"g_node_push_allocator", &NodePushAllocator},
	{"g_node_reverse_children", &NodeReverseChildren},
	{"g_node_traverse", &NodeTraverse},
	{"g_node_unlink", &NodeUnlink},
	{"g_nullify_pointer", &NullifyPointer},
	{"g_on_error_query", &OnErrorQuery},
	{"g_on_error_stack_trace", &OnErrorStackTrace},
	{"g_once_impl", &OnceImpl},
	{"g_once_init_enter", &OnceInitEnter},
	{"g_once_init_enter_impl", &OnceInitEnterImpl},
	{"g_once_init_leave", &OnceInitLeave},
	{"g_open", &Open},
	{"g_option_context_add_group", &OptionContextAddGroup},
	{"g_option_context_add_main_entries", &OptionContextAddMainEntries},
	{"g_option_context_free", &OptionContextFree},
	{"g_option_context_get_description", &OptionContextGetDescription},
	{"g_option_context_get_help", &OptionContextGetHelp},
	{"g_option_context_get_help_enabled", &OptionContextGetHelpEnabled},
	{"g_option_context_get_ignore_unknown_options", &OptionContextGetIgnoreUnknownOptions},
	{"g_option_context_get_main_group", &OptionContextGetMainGroup},
	{"g_option_context_get_summary", &OptionContextGetSummary},
	{"g_option_context_new", &OptionContextNew},
	{"g_option_context_parse", &OptionContextParse},
	{"g_option_context_set_description", &OptionContextSetDescription},
	{"g_option_context_set_help_enabled", &OptionContextSetHelpEnabled},
	{"g_option_context_set_ignore_unknown_options", &OptionContextSetIgnoreUnknownOptions},
	{"g_option_context_set_main_group", &OptionContextSetMainGroup},
	{"g_option_context_set_summary", &OptionContextSetSummary},
	{"g_option_context_set_translate_func", &OptionContextSetTranslateFunc},
	{"g_option_context_set_translation_domain", &OptionContextSetTranslationDomain},
	{"g_option_error_quark", &OptionErrorQuark},
	{"g_option_group_add_entries", &OptionGroupAddEntries},
	{"g_option_group_free", &OptionGroupFree},
	{"g_option_group_new", &OptionGroupNew},
	{"g_option_group_set_error_hook", &OptionGroupSetErrorHook},
	{"g_option_group_set_parse_hooks", &OptionGroupSetParseHooks},
	{"g_option_group_set_translate_func", &OptionGroupSetTranslateFunc},
	{"g_option_group_set_translation_domain", &OptionGroupSetTranslationDomain},
	{"g_parse_debug_string", &ParseDebugString},
	{"g_path_get_basename", &PathGetBasename},
	{"g_path_get_dirname", &PathGetDirname},
	{"g_path_is_absolute", &PathIsAbsolute},
	{"g_path_skip_root", &PathSkipRoot},
	{"g_pattern_match", &PatternMatch},
	{"g_pattern_match_simple", &PatternMatchSimple},
	{"g_pattern_match_string", &PatternMatchString},
	{"g_pattern_spec_equal", &PatternSpecEqual},
	{"g_pattern_spec_free", &PatternSpecFree},
	{"g_pattern_spec_new", &PatternSpecNew},
	{"g_poll", &Poll},
	{"g_prefix_error", &PrefixError},
	{"g_print", &Print},
	{"g_printerr", &Printerr},
	{"g_printf", &Printf},
	{"g_printf_string_upper_bound", &PrintfStringUpperBound},
	{"g_propagate_error", &PropagateError},
	{"g_propagate_prefixed_error", &PropagatePrefixedError},
	{"g_ptr_array_add", &PtrArrayAdd},
	{"g_ptr_array_foreach", &PtrArrayForeach},
	{"g_ptr_array_free", &PtrArrayFree},
	{"g_ptr_array_new", &PtrArrayNew},
	{"g_ptr_array_new_with_free_func", &PtrArrayNewWithFreeFunc},
	{"g_ptr_array_ref", &PtrArrayRef},
	{"g_ptr_array_remove", &PtrArrayRemove},
	{"g_ptr_array_remove_fast", &PtrArrayRemoveFast},
	{"g_ptr_array_remove_index", &PtrArrayRemoveIndex},
	{"g_ptr_array_remove_index_fast", &PtrArrayRemoveIndexFast},
	{"g_ptr_array_remove_range", &PtrArrayRemoveRange},
	{"g_ptr_array_set_free_func", &PtrArraySetFreeFunc},
	{"g_ptr_array_set_size", &PtrArraySetSize},
	{"g_ptr_array_sized_new", &PtrArraySizedNew},
	{"g_ptr_array_sort", &PtrArraySort},
	{"g_ptr_array_sort_with_data", &PtrArraySortWithData},
	{"g_ptr_array_unref", &PtrArrayUnref},
	{"g_qsort_with_data", &QsortWithData},
	{"g_quark_from_static_string", &QuarkFromStaticString},
	{"g_quark_from_string", &QuarkFromString},
	{"g_quark_to_string", &QuarkToString},
	{"g_quark_try_string", &QuarkTryString},
	{"g_queue_clear", &QueueClear},
	{"g_queue_copy", &QueueCopy},
	{"g_queue_delete_link", &QueueDeleteLink},
	{"g_queue_find", &QueueFind},
	{"g_queue_find_custom", &QueueFindCustom},
	{"g_queue_foreach", &QueueForeach},
	{"g_queue_free", &QueueFree},
	{"g_queue_get_length", &QueueGetLength},
	{"g_queue_index", &QueueIndex},
	{"g_queue_init", &QueueInit},
	{"g_queue_insert_after", &QueueInsertAfter},
	{"g_queue_insert_before", &QueueInsertBefore},
	{"g_queue_insert_sorted", &QueueInsertSorted},
	{"g_queue_is_empty", &QueueIsEmpty},
	{"g_queue_link_index", &QueueLinkIndex},
	{"g_queue_new", &QueueNew},
	{"g_queue_peek_head", &QueuePeekHead},
	{"g_queue_peek_head_link", &QueuePeekHeadLink},
	{"g_queue_peek_nth", &QueuePeekNth},
	{"g_queue_peek_nth_link", &QueuePeekNthLink},
	{"g_queue_peek_tail", &QueuePeekTail},
	{"g_queue_peek_tail_link", &QueuePeekTailLink},
	{"g_queue_pop_head", &QueuePopHead},
	{"g_queue_pop_head_link", &QueuePopHeadLink},
	{"g_queue_pop_nth", &QueuePopNth},
	{"g_queue_pop_nth_link", &QueuePopNthLink},
	{"g_queue_pop_tail", &QueuePopTail},
	{"g_queue_pop_tail_link", &QueuePopTailLink},
	{"g_queue_push_head", &QueuePushHead},
	{"g_queue_push_head_link", &QueuePushHeadLink},
	{"g_queue_push_nth", &QueuePushNth},
	{"g_queue_push_nth_link", &QueuePushNthLink},
	{"g_queue_push_tail", &QueuePushTail},
	{"g_queue_push_tail_link", &QueuePushTailLink},
	{"g_queue_remove", &QueueRemove},
	{"g_queue_remove_all", &QueueRemoveAll},
	{"g_queue_reverse", &QueueReverse},
	{"g_queue_sort", &QueueSort},
	{"g_queue_unlink", &QueueUnlink},
	{"g_rand_copy", &RandCopy},
	{"g_rand_double", &RandDouble},
	{"g_rand_double_range", &RandDoubleRange},
	{"g_rand_free", &RandFree},
	{"g_rand_int", &RandInt},
	{"g_rand_int_range", &RandIntRange},
	{"g_rand_new", &RandNew},
	{"g_rand_new_with_seed", &RandNewWithSeed},
	{"g_rand_new_with_seed_array", &RandNewWithSeedArray},
	{"g_rand_set_seed", &RandSetSeed},
	{"g_rand_set_seed_array", &RandSetSeedArray},
	{"g_random_double", &RandomDouble},
	{"g_random_double_range", &RandomDoubleRange},
	{"g_random_int", &RandomInt},
	{"g_random_int_range", &RandomIntRange},
	{"g_random_set_seed", &RandomSetSeed},
	{"g_realloc", &Realloc},
	{"g_realloc_n", &ReallocN},
	{"g_regex_check_replacement", &RegexCheckReplacement},
	{"g_regex_error_quark", &RegexErrorQuark},
	{"g_regex_escape_string", &RegexEscapeString},
	{"g_regex_get_capture_count", &RegexGetCaptureCount},
	{"g_regex_get_compile_flags", &RegexGetCompileFlags},
	{"g_regex_get_match_flags", &RegexGetMatchFlags},
	{"g_regex_get_max_backref", &RegexGetMaxBackref},
	{"g_regex_get_pattern", &RegexGetPattern},
	{"g_regex_get_string_number", &RegexGetStringNumber},
	{"g_regex_match", &RegexMatch},
	{"g_regex_match_all", &RegexMatchAll},
	{"g_regex_match_all_full", &RegexMatchAllFull},
	{"g_regex_match_full", &RegexMatchFull},
	{"g_regex_match_simple", &RegexMatchSimple},
	{"g_regex_new", &RegexNew},
	{"g_regex_ref", &RegexRef},
	{"g_regex_replace", &RegexReplace},
	{"g_regex_replace_eval", &RegexReplaceEval},
	{"g_regex_replace_literal", &RegexReplaceLiteral},
	{"g_regex_split", &RegexSplit},
	{"g_regex_split_full", &RegexSplitFull},
	{"g_regex_split_simple", &RegexSplitSimple},
	{"g_regex_unref", &RegexUnref},
	{"g_relation_count", &RelationCount},
	{"g_relation_delete", &RelationDelete},
	{"g_relation_destroy", &RelationDestroy},
	{"g_relation_exists", &RelationExists},
	{"g_relation_index", &RelationIndex},
	{"g_relation_insert", &RelationInsert},
	{"g_relation_new", &RelationNew},
	{"g_relation_print", &RelationPrint},
	{"g_relation_select", &RelationSelect},
	{"g_reload_user_special_dirs_cache", &ReloadUserSpecialDirsCache},
	{"g_remove", &Remove},
	{"g_rename", &Rename},
	{"g_return_if_fail_warning", &ReturnIfFailWarning},
	{"g_rmdir", &Rmdir},
	{"g_scanner_cur_line", &ScannerCurLine},
	{"g_scanner_cur_position", &ScannerCurPosition},
	{"g_scanner_cur_token", &ScannerCurToken},
	{"g_scanner_cur_value", &ScannerCurValue},
	{"g_scanner_destroy", &ScannerDestroy},
	{"g_scanner_eof", &ScannerEof},
	{"g_scanner_error", &ScannerError},
	{"g_scanner_get_next_token", &ScannerGetNextToken},
	{"g_scanner_input_file", &ScannerInputFile},
	{"g_scanner_input_text", &ScannerInputText},
	{"g_scanner_lookup_symbol", &ScannerLookupSymbol},
	{"g_scanner_new", &ScannerNew},
	{"g_scanner_peek_next_token", &ScannerPeekNextToken},
	{"g_scanner_scope_add_symbol", &ScannerScopeAddSymbol},
	{"g_scanner_scope_foreach_symbol", &ScannerScopeForeachSymbol},
	{"g_scanner_scope_lookup_symbol", &ScannerScopeLookupSymbol},
	{"g_scanner_scope_remove_symbol", &ScannerScopeRemoveSymbol},
	{"g_scanner_set_scope", &ScannerSetScope},
	{"g_scanner_sync_file_offset", &ScannerSyncFileOffset},
	{"g_scanner_unexp_token", &ScannerUnexpToken},
	{"g_scanner_warn", &ScannerWarn},
	{"g_sequence_append", &SequenceAppend},
	{"g_sequence_foreach", &SequenceForeach},
	{"g_sequence_foreach_range", &SequenceForeachRange},
	{"g_sequence_free", &SequenceFree},
	{"g_sequence_get", &SequenceGet},
	{"g_sequence_get_begin_iter", &SequenceGetBeginIter},
	{"g_sequence_get_end_iter", &SequenceGetEndIter},
	{"g_sequence_get_iter_at_pos", &SequenceGetIterAtPos},
	{"g_sequence_get_length", &SequenceGetLength},
	{"g_sequence_insert_before", &SequenceInsertBefore},
	{"g_sequence_insert_sorted", &SequenceInsertSorted},
	{"g_sequence_insert_sorted_iter", &SequenceInsertSortedIter},
	{"g_sequence_iter_compare", &SequenceIterCompare},
	{"g_sequence_iter_get_position", &SequenceIterGetPosition},
	{"g_sequence_iter_get_sequence", &SequenceIterGetSequence},
	{"g_sequence_iter_is_begin", &SequenceIterIsBegin},
	{"g_sequence_iter_is_end", &SequenceIterIsEnd},
	{"g_sequence_iter_move", &SequenceIterMove},
	{"g_sequence_iter_next", &SequenceIterNext},
	{"g_sequence_iter_prev", &SequenceIterPrev},
	{"g_sequence_lookup", &SequenceLookup},
	{"g_sequence_lookup_iter", &SequenceLookupIter},
	{"g_sequence_move", &SequenceMove},
	{"g_sequence_move_range", &SequenceMoveRange},
	{"g_sequence_new", &SequenceNew},
	{"g_sequence_prepend", &SequencePrepend},
	{"g_sequence_range_get_midpoint", &SequenceRangeGetMidpoint},
	{"g_sequence_remove", &SequenceRemove},
	{"g_sequence_remove_range", &SequenceRemoveRange},
	{"g_sequence_search", &SequenceSearch},
	{"g_sequence_search_iter", &SequenceSearchIter},
	{"g_sequence_set", &SequenceSet},
	{"g_sequence_sort", &SequenceSort},
	{"g_sequence_sort_changed", &SequenceSortChanged},
	{"g_sequence_sort_changed_iter", &SequenceSortChangedIter},
	{"g_sequence_sort_iter", &SequenceSortIter},
	{"g_sequence_swap", &SequenceSwap},
	{"g_set_application_name", &SetApplicationName},
	{"g_set_error", &SetError},
	{"g_set_error_literal", &SetErrorLiteral},
	{"g_set_prgname", &SetPrgname},
	{"g_set_print_handler", &SetPrintHandler},
	{"g_set_printerr_handler", &SetPrinterrHandler},
	// Windows: _utf8 {"g_setenv", &Setenv},
	{"g_setenv_utf8", &Setenv},
	{"g_shell_error_quark", &ShellErrorQuark},
	{"g_shell_parse_argv", &ShellParseArgv},
	{"g_shell_quote", &ShellQuote},
	{"g_shell_unquote", &ShellUnquote},
	{"g_slice_alloc", &SliceAlloc},
	{"g_slice_alloc0", &SliceAlloc0},
	{"g_slice_copy", &SliceCopy},
	{"g_slice_free1", &SliceFree1},
	{"g_slice_free_chain_with_offset", &SliceFreeChainWithOffset},
	{"g_slice_get_config", &SliceGetConfig},
	{"g_slice_get_config_state", &SliceGetConfigState},
	{"g_slice_set_config", &SliceSetConfig},
	{"g_slist_alloc", &SlistAlloc},
	{"g_slist_append", &SlistAppend},
	{"g_slist_concat", &SlistConcat},
	{"g_slist_copy", &SlistCopy},
	{"g_slist_delete_link", &SlistDeleteLink},
	{"g_slist_find", &SlistFind},
	{"g_slist_find_custom", &SlistFindCustom},
	{"g_slist_foreach", &SlistForeach},
	{"g_slist_free", &SlistFree},
	{"g_slist_free_1", &SlistFree1},
	{"g_slist_free_full", &SlistFreeFull},
	{"g_slist_index", &SlistIndex},
	{"g_slist_insert", &SlistInsert},
	{"g_slist_insert_before", &SlistInsertBefore},
	{"g_slist_insert_sorted", &SlistInsertSorted},
	{"g_slist_insert_sorted_with_data", &SlistInsertSortedWithData},
	{"g_slist_last", &SlistLast},
	{"g_slist_length", &SlistLength},
	{"g_slist_nth", &SlistNth},
	{"g_slist_nth_data", &SlistNthData},
	{"g_slist_pop_allocator", &SlistPopAllocator},
	{"g_slist_position", &SlistPosition},
	{"g_slist_prepend", &SlistPrepend},
	{"g_slist_push_allocator", &SlistPushAllocator},
	{"g_slist_remove", &SlistRemove},
	{"g_slist_remove_all", &SlistRemoveAll},
	{"g_slist_remove_link", &SlistRemoveLink},
	{"g_slist_reverse", &SlistReverse},
	{"g_slist_sort", &SlistSort},
	{"g_slist_sort_with_data", &SlistSortWithData},
	{"g_snprintf", &Snprintf},
	{"g_source_add_child_source", &SourceAddChildSource},
	{"g_source_add_poll", &SourceAddPoll},
	{"g_source_attach", &SourceAttach},
	{"g_source_destroy", &SourceDestroy},
	{"g_source_get_can_recurse", &SourceGetCanRecurse},
	{"g_source_get_context", &SourceGetContext},
	{"g_source_get_current_time", &SourceGetCurrentTime},
	{"g_source_get_id", &SourceGetId},
	{"g_source_get_name", &SourceGetName},
	{"g_source_get_priority", &SourceGetPriority},
	{"g_source_get_time", &SourceGetTime},
	{"g_source_is_destroyed", &SourceIsDestroyed},
	{"g_source_new", &SourceNew},
	{"g_source_ref", &SourceRef},
	{"g_source_remove", &SourceRemove},
	{"g_source_remove_by_funcs_user_data", &SourceRemoveByFuncsUserData},
	{"g_source_remove_by_user_data", &SourceRemoveByUserData},
	{"g_source_remove_child_source", &SourceRemoveChildSource},
	{"g_source_remove_poll", &SourceRemovePoll},
	{"g_source_set_callback", &SourceSetCallback},
	{"g_source_set_callback_indirect", &SourceSetCallbackIndirect},
	{"g_source_set_can_recurse", &SourceSetCanRecurse},
	{"g_source_set_funcs", &SourceSetFuncs},
	{"g_source_set_name", &SourceSetName},
	{"g_source_set_name_by_id", &SourceSetNameById},
	{"g_source_set_priority", &SourceSetPriority},
	{"g_source_unref", &SourceUnref},
	{"g_spaced_primes_closest", &SpacedPrimesClosest},
	// Windows: _utf8 {"g_spawn_async", &SpawnAsync},
	{"g_spawn_async_utf8", &SpawnAsync},
	// Windows: _utf8 {"g_spawn_async_with_pipes", &SpawnAsyncWithPipes},
	{"g_spawn_async_with_pipes_utf8", &SpawnAsyncWithPipes},
	{"g_spawn_close_pid", &SpawnClosePid},
	// Windows: _utf8 {"g_spawn_command_line_async", &SpawnCommandLineAsync},
	{"g_spawn_command_line_async_utf8", &SpawnCommandLineAsync},
	// Windows: _utf8 {"g_spawn_command_line_sync", &SpawnCommandLineSync},
	{"g_spawn_command_line_sync_utf8", &SpawnCommandLineSync},
	{"g_spawn_error_quark", &SpawnErrorQuark},
	// Windows: _utf8 {"g_spawn_sync", &SpawnSync},
	{"g_spawn_sync_utf8", &SpawnSync},
	{"g_sprintf", &Sprintf},
	{"g_stat", &Stat},
	{"g_static_mutex_free", &StaticMutexFree},
	{"g_static_mutex_get_mutex_impl", &StaticMutexGetMutexImpl},
	{"g_static_mutex_init", &StaticMutexInit},
	{"g_static_private_free", &StaticPrivateFree},
	{"g_static_private_get", &StaticPrivateGet},
	{"g_static_private_init", &StaticPrivateInit},
	{"g_static_private_set", &StaticPrivateSet},
	{"g_static_rec_mutex_free", &StaticRecMutexFree},
	{"g_static_rec_mutex_init", &StaticRecMutexInit},
	{"g_static_rec_mutex_lock", &StaticRecMutexLock},
	{"g_static_rec_mutex_lock_full", &StaticRecMutexLockFull},
	{"g_static_rec_mutex_trylock", &StaticRecMutexTrylock},
	{"g_static_rec_mutex_unlock", &StaticRecMutexUnlock},
	{"g_static_rec_mutex_unlock_full", &StaticRecMutexUnlockFull},
	{"g_static_rw_lock_free", &StaticRwLockFree},
	{"g_static_rw_lock_init", &StaticRwLockInit},
	{"g_static_rw_lock_reader_lock", &StaticRwLockReaderLock},
	{"g_static_rw_lock_reader_trylock", &StaticRwLockReaderTrylock},
	{"g_static_rw_lock_reader_unlock", &StaticRwLockReaderUnlock},
	{"g_static_rw_lock_writer_lock", &StaticRwLockWriterLock},
	{"g_static_rw_lock_writer_trylock", &StaticRwLockWriterTrylock},
	{"g_static_rw_lock_writer_unlock", &StaticRwLockWriterUnlock},
	{"g_stpcpy", &Stpcpy},
	{"g_str_equal", &StrEqual},
	{"g_str_has_prefix", &StrHasPrefix},
	{"g_str_has_suffix", &StrHasSuffix},
	{"g_str_hash", &StrHash},
	{"g_strcanon", &Strcanon},
	{"g_strcasecmp", &Strcasecmp},
	{"g_strchomp", &Strchomp},
	{"g_strchug", &Strchug},
	{"g_strcmp0", &Strcmp0},
	{"g_strcompress", &Strcompress},
	{"g_strconcat", &Strconcat},
	{"g_strdelimit", &Strdelimit},
	{"g_strdown", &Strdown},
	{"g_strdup", &Strdup},
	{"g_strdup_printf", &StrdupPrintf},
	{"g_strdup_vprintf", &StrdupVprintf},
	{"g_strdupv", &Strdupv},
	{"g_strerror", &Strerror},
	{"g_strescape", &Strescape},
	{"g_strfreev", &Strfreev},
	{"g_string_append", &StringAppend},
	{"g_string_append_c", &StringAppendC},
	{"g_string_append_len", &StringAppendLen},
	{"g_string_append_printf", &StringAppendPrintf},
	{"g_string_append_unichar", &StringAppendUnichar},
	{"g_string_append_uri_escaped", &StringAppendUriEscaped},
	{"g_string_append_vprintf", &StringAppendVprintf},
	{"g_string_ascii_down", &StringAsciiDown},
	{"g_string_ascii_up", &StringAsciiUp},
	{"g_string_assign", &StringAssign},
	{"g_string_chunk_clear", &StringChunkClear},
	{"g_string_chunk_free", &StringChunkFree},
	{"g_string_chunk_insert", &StringChunkInsert},
	{"g_string_chunk_insert_const", &StringChunkInsertConst},
	{"g_string_chunk_insert_len", &StringChunkInsertLen},
	{"g_string_chunk_new", &StringChunkNew},
	{"g_string_down", &StringDown},
	{"g_string_equal", &StringEqual},
	{"g_string_erase", &StringErase},
	{"g_string_free", &StringFree},
	{"g_string_hash", &StringHash},
	{"g_string_insert", &StringInsert},
	{"g_string_insert_c", &StringInsertC},
	{"g_string_insert_len", &StringInsertLen},
	{"g_string_insert_unichar", &StringInsertUnichar},
	{"g_string_new", &StringNew},
	{"g_string_new_len", &StringNewLen},
	{"g_string_overwrite", &StringOverwrite},
	{"g_string_overwrite_len", &StringOverwriteLen},
	{"g_string_prepend", &StringPrepend},
	{"g_string_prepend_c", &StringPrependC},
	{"g_string_prepend_len", &StringPrependLen},
	{"g_string_prepend_unichar", &StringPrependUnichar},
	{"g_string_printf", &StringPrintf},
	{"g_string_set_size", &StringSetSize},
	{"g_string_sized_new", &StringSizedNew},
	{"g_string_truncate", &StringTruncate},
	{"g_string_up", &StringUp},
	{"g_string_vprintf", &StringVprintf},
	{"g_strip_context", &StripContext},
	{"g_strjoin", &Strjoin},
	{"g_strjoinv", &Strjoinv},
	{"g_strlcat", &Strlcat},
	{"g_strlcpy", &Strlcpy},
	{"g_strncasecmp", &Strncasecmp},
	{"g_strndup", &Strndup},
	{"g_strnfill", &Strnfill},
	{"g_strreverse", &Strreverse},
	{"g_strrstr", &Strrstr},
	{"g_strrstr_len", &StrrstrLen},
	{"g_strsignal", &Strsignal},
	{"g_strsplit", &Strsplit},
	{"g_strsplit_set", &StrsplitSet},
	{"g_strstr_len", &StrstrLen},
	{"g_strtod", &Strtod},
	{"g_strup", &Strup},
	{"g_strv_length", &StrvLength},
	{"g_test_add_data_func", &TestAddDataFunc},
	{"g_test_add_func", &TestAddFunc},
	{"g_test_add_vtable", &TestAddVtable},
	{"g_test_bug", &TestBug},
	{"g_test_bug_base", &TestBugBase},
	{"g_test_create_case", &TestCreateCase},
	{"g_test_create_suite", &TestCreateSuite},
	{"g_test_get_root", &TestGetRoot},
	{"g_test_init", &TestInit},
	{"g_test_log_buffer_free", &TestLogBufferFree},
	{"g_test_log_buffer_new", &TestLogBufferNew},
	{"g_test_log_buffer_pop", &TestLogBufferPop},
	{"g_test_log_buffer_push", &TestLogBufferPush},
	{"g_test_log_msg_free", &TestLogMsgFree},
	{"g_test_log_set_fatal_handler", &TestLogSetFatalHandler},
	{"g_test_log_type_name", &TestLogTypeName},
	{"g_test_maximized_result", &TestMaximizedResult},
	{"g_test_message", &TestMessage},
	{"g_test_minimized_result", &TestMinimizedResult},
	{"g_test_queue_destroy", &TestQueueDestroy},
	{"g_test_queue_free", &TestQueueFree},
	{"g_test_rand_double", &TestRandDouble},
	{"g_test_rand_double_range", &TestRandDoubleRange},
	{"g_test_rand_int", &TestRandInt},
	{"g_test_rand_int_range", &TestRandIntRange},
	{"g_test_run", &TestRun},
	{"g_test_run_suite", &TestRunSuite},
	{"g_test_suite_add", &TestSuiteAdd},
	{"g_test_suite_add_suite", &TestSuiteAddSuite},
	{"g_test_timer_elapsed", &TestTimerElapsed},
	{"g_test_timer_last", &TestTimerLast},
	{"g_test_timer_start", &TestTimerStart},
	{"g_test_trap_assertions", &TestTrapAssertions},
	{"g_test_trap_fork", &TestTrapFork},
	{"g_test_trap_has_passed", &TestTrapHasPassed},
	{"g_test_trap_reached_timeout", &TestTrapReachedTimeout},
	{"g_thread_create_full", &ThreadCreateFull},
	{"g_thread_error_quark", &ThreadErrorQuark},
	{"g_thread_exit", &ThreadExit},
	{"g_thread_foreach", &ThreadForeach},
	{"g_thread_get_initialized", &ThreadGetInitialized},
	// Undocumented {"g_thread_init_glib", &ThreadInitGlib},
	{"g_thread_join", &ThreadJoin},
	{"g_thread_pool_free", &ThreadPoolFree},
	{"g_thread_pool_get_max_idle_time", &ThreadPoolGetMaxIdleTime},
	{"g_thread_pool_get_max_threads", &ThreadPoolGetMaxThreads},
	{"g_thread_pool_get_max_unused_threads", &ThreadPoolGetMaxUnusedThreads},
	{"g_thread_pool_get_num_threads", &ThreadPoolGetNumThreads},
	{"g_thread_pool_get_num_unused_threads", &ThreadPoolGetNumUnusedThreads},
	{"g_thread_pool_new", &ThreadPoolNew},
	{"g_thread_pool_push", &ThreadPoolPush},
	{"g_thread_pool_set_max_idle_time", &ThreadPoolSetMaxIdleTime},
	{"g_thread_pool_set_max_threads", &ThreadPoolSetMaxThreads},
	{"g_thread_pool_set_max_unused_threads", &ThreadPoolSetMaxUnusedThreads},
	{"g_thread_pool_set_sort_function", &ThreadPoolSetSortFunction},
	{"g_thread_pool_stop_unused_threads", &ThreadPoolStopUnusedThreads},
	{"g_thread_pool_unprocessed", &ThreadPoolUnprocessed},
	{"g_thread_self", &ThreadSelf},
	{"g_thread_set_priority", &ThreadSetPriority},
	{"g_time_val_add", &TimeValAdd},
	{"g_time_val_from_iso8601", &TimeValFromIso8601},
	{"g_time_val_to_iso8601", &TimeValToIso8601},
	{"g_time_zone_adjust_time", &TimeZoneAdjustTime},
	{"g_time_zone_find_interval", &TimeZoneFindInterval},
	{"g_time_zone_get_abbreviation", &TimeZoneGetAbbreviation},
	{"g_time_zone_get_offset", &TimeZoneGetOffset},
	{"g_time_zone_is_dst", &TimeZoneIsDst},
	{"g_time_zone_new", &TimeZoneNew},
	{"g_time_zone_new_local", &TimeZoneNewLocal},
	{"g_time_zone_new_utc", &TimeZoneNewUtc},
	{"g_time_zone_ref", &TimeZoneRef},
	{"g_time_zone_unref", &TimeZoneUnref},
	{"g_timeout_add", &TimeoutAdd},
	{"g_timeout_add_full", &TimeoutAddFull},
	{"g_timeout_add_seconds", &TimeoutAddSeconds},
	{"g_timeout_add_seconds_full", &TimeoutAddSecondsFull},
	{"g_timeout_source_new", &TimeoutSourceNew},
	{"g_timeout_source_new_seconds", &TimeoutSourceNewSeconds},
	{"g_timer_continue", &TimerContinue},
	{"g_timer_destroy", &TimerDestroy},
	{"g_timer_elapsed", &TimerElapsed},
	{"g_timer_new", &TimerNew},
	{"g_timer_reset", &TimerReset},
	{"g_timer_start", &TimerStart},
	{"g_timer_stop", &TimerStop},
	{"g_trash_stack_height", &TrashStackHeight},
	{"g_trash_stack_peek", &TrashStackPeek},
	{"g_trash_stack_pop", &TrashStackPop},
	{"g_trash_stack_push", &TrashStackPush},
	{"g_tree_destroy", &TreeDestroy},
	{"g_tree_foreach", &TreeForeach},
	{"g_tree_height", &TreeHeight},
	{"g_tree_insert", &TreeInsert},
	{"g_tree_lookup", &TreeLookup},
	{"g_tree_lookup_extended", &TreeLookupExtended},
	{"g_tree_new", &TreeNew},
	{"g_tree_new_full", &TreeNewFull},
	{"g_tree_new_with_data", &TreeNewWithData},
	{"g_tree_nnodes", &TreeNnodes},
	{"g_tree_ref", &TreeRef},
	{"g_tree_remove", &TreeRemove},
	{"g_tree_replace", &TreeReplace},
	{"g_tree_search", &TreeSearch},
	{"g_tree_steal", &TreeSteal},
	{"g_tree_traverse", &TreeTraverse},
	{"g_tree_unref", &TreeUnref},
	{"g_try_malloc", &TryMalloc},
	{"g_try_malloc0", &TryMalloc0},
	{"g_try_malloc0_n", &TryMalloc0N},
	{"g_try_malloc_n", &TryMallocN},
	{"g_try_realloc", &TryRealloc},
	{"g_try_realloc_n", &TryReallocN},
	{"g_tuples_destroy", &TuplesDestroy},
	{"g_tuples_index", &TuplesIndex},
	{"g_ucs4_to_utf16", &Ucs4ToUtf16},
	{"g_ucs4_to_utf8", &Ucs4ToUtf8},
	{"g_unichar_break_type", &UnicharBreakType},
	{"g_unichar_combining_class", &UnicharCombiningClass},
	{"g_unichar_digit_value", &UnicharDigitValue},
	{"g_unichar_get_mirror_char", &UnicharGetMirrorChar},
	{"g_unichar_get_script", &UnicharGetScript},
	{"g_unichar_isalnum", &UnicharIsalnum},
	{"g_unichar_isalpha", &UnicharIsalpha},
	{"g_unichar_iscntrl", &UnicharIscntrl},
	{"g_unichar_isdefined", &UnicharIsdefined},
	{"g_unichar_isdigit", &UnicharIsdigit},
	{"g_unichar_isgraph", &UnicharIsgraph},
	{"g_unichar_islower", &UnicharIslower},
	{"g_unichar_ismark", &UnicharIsmark},
	{"g_unichar_isprint", &UnicharIsprint},
	{"g_unichar_ispunct", &UnicharIspunct},
	{"g_unichar_isspace", &UnicharIsspace},
	{"g_unichar_istitle", &UnicharIstitle},
	{"g_unichar_isupper", &UnicharIsupper},
	{"g_unichar_iswide", &UnicharIswide},
	{"g_unichar_iswide_cjk", &UnicharIswideCjk},
	{"g_unichar_isxdigit", &UnicharIsxdigit},
	{"g_unichar_iszerowidth", &UnicharIszerowidth},
	{"g_unichar_to_utf8", &UnicharToUtf8},
	{"g_unichar_tolower", &UnicharTolower},
	{"g_unichar_totitle", &UnicharTotitle},
	{"g_unichar_toupper", &UnicharToupper},
	{"g_unichar_type", &UnicharType},
	{"g_unichar_validate", &UnicharValidate},
	{"g_unichar_xdigit_value", &UnicharXdigitValue},
	{"g_unicode_canonical_decomposition", &UnicodeCanonicalDecomposition},
	{"g_unicode_canonical_ordering", &UnicodeCanonicalOrdering},
	{"g_unlink", &Unlink},
	// Windows: _utf8 {"g_unsetenv", &Unsetenv},
	{"g_unsetenv_utf8", &Unsetenv},
	{"g_uri_escape_string", &UriEscapeString},
	{"g_uri_list_extract_uris", &UriListExtractUris},
	{"g_uri_parse_scheme", &UriParseScheme},
	{"g_uri_unescape_segment", &UriUnescapeSegment},
	{"g_uri_unescape_string", &UriUnescapeString},
	{"g_usleep", &Usleep},
	{"g_utf16_to_ucs4", &Utf16ToUcs4},
	{"g_utf16_to_utf8", &Utf16ToUtf8},
	{"g_utf8_casefold", &Utf8Casefold},
	{"g_utf8_collate", &Utf8Collate},
	{"g_utf8_collate_key", &Utf8CollateKey},
	{"g_utf8_collate_key_for_filename", &Utf8CollateKeyForFilename},
	{"g_utf8_find_next_char", &Utf8FindNextChar},
	{"g_utf8_find_prev_char", &Utf8FindPrevChar},
	{"g_utf8_get_char", &Utf8GetChar},
	{"g_utf8_get_char_validated", &Utf8GetCharValidated},
	{"g_utf8_normalize", &Utf8Normalize},
	{"g_utf8_offset_to_pointer", &Utf8OffsetToPointer},
	{"g_utf8_pointer_to_offset", &Utf8PointerToOffset},
	{"g_utf8_prev_char", &Utf8PrevChar},
	{"g_utf8_strchr", &Utf8Strchr},
	{"g_utf8_strdown", &Utf8Strdown},
	{"g_utf8_strlen", &Utf8Strlen},
	{"g_utf8_strncpy", &Utf8Strncpy},
	{"g_utf8_strrchr", &Utf8Strrchr},
	{"g_utf8_strreverse", &Utf8Strreverse},
	{"g_utf8_strup", &Utf8Strup},
	{"g_utf8_to_ucs4", &Utf8ToUcs4},
	{"g_utf8_to_ucs4_fast", &Utf8ToUcs4Fast},
	{"g_utf8_to_utf16", &Utf8ToUtf16},
	{"g_utf8_validate", &Utf8Validate},
	{"g_utime", &Utime},
	{"g_variant_builder_add", &VariantBuilderAdd},
	{"g_variant_builder_add_parsed", &VariantBuilderAddParsed},
	{"g_variant_builder_add_value", &VariantBuilderAddValue},
	{"g_variant_builder_clear", &VariantBuilderClear},
	{"g_variant_builder_close", &VariantBuilderClose},
	{"g_variant_builder_end", &VariantBuilderEnd},
	{"g_variant_builder_init", &VariantBuilderInit},
	{"g_variant_builder_new", &VariantBuilderNew},
	{"g_variant_builder_open", &VariantBuilderOpen},
	{"g_variant_builder_ref", &VariantBuilderRef},
	{"g_variant_builder_unref", &VariantBuilderUnref},
	{"g_variant_byteswap", &VariantByteswap},
	{"g_variant_classify", &VariantClassify},
	{"g_variant_compare", &VariantCompare},
	{"g_variant_dup_bytestring", &VariantDupBytestring},
	{"g_variant_dup_bytestring_array", &VariantDupBytestringArray},
	{"g_variant_dup_string", &VariantDupString},
	{"g_variant_dup_strv", &VariantDupStrv},
	{"g_variant_equal", &VariantEqual},
	// Undocumented {"g_variant_format_string_scan", &VariantFormatStringScan},
	// Undocumented {"g_variant_format_string_scan_type", &VariantFormatStringScanType},
	{"g_variant_get", &VariantGet},
	{"g_variant_get_boolean", &VariantGetBoolean},
	{"g_variant_get_byte", &VariantGetByte},
	{"g_variant_get_bytestring", &VariantGetBytestring},
	{"g_variant_get_bytestring_array", &VariantGetBytestringArray},
	{"g_variant_get_child", &VariantGetChild},
	{"g_variant_get_child_value", &VariantGetChildValue},
	{"g_variant_get_data", &VariantGetData},
	{"g_variant_get_double", &VariantGetDouble},
	{"g_variant_get_fixed_array", &VariantGetFixedArray},
	{"g_variant_get_handle", &VariantGetHandle},
	{"g_variant_get_int16", &VariantGetInt16},
	{"g_variant_get_int32", &VariantGetInt32},
	{"g_variant_get_int64", &VariantGetInt64},
	{"g_variant_get_maybe", &VariantGetMaybe},
	{"g_variant_get_normal_form", &VariantGetNormalForm},
	{"g_variant_get_size", &VariantGetSize},
	{"g_variant_get_string", &VariantGetString},
	{"g_variant_get_strv", &VariantGetStrv},
	{"g_variant_get_type", &VariantGetType},
	{"g_variant_get_type_string", &VariantGetTypeString},
	{"g_variant_get_uint16", &VariantGetUint16},
	{"g_variant_get_uint32", &VariantGetUint32},
	{"g_variant_get_uint64", &VariantGetUint64},
	{"g_variant_get_va", &VariantGetVa},
	{"g_variant_get_variant", &VariantGetVariant},
	{"g_variant_hash", &VariantHash},
	{"g_variant_is_container", &VariantIsContainer},
	{"g_variant_is_floating", &VariantIsFloating},
	{"g_variant_is_normal_form", &VariantIsNormalForm},
	{"g_variant_is_object_path", &VariantIsObjectPath},
	{"g_variant_is_of_type", &VariantIsOfType},
	{"g_variant_is_signature", &VariantIsSignature},
	{"g_variant_iter_copy", &VariantIterCopy},
	{"g_variant_iter_free", &VariantIterFree},
	{"g_variant_iter_init", &VariantIterInit},
	{"g_variant_iter_loop", &VariantIterLoop},
	{"g_variant_iter_n_children", &VariantIterNChildren},
	{"g_variant_iter_new", &VariantIterNew},
	{"g_variant_iter_next", &VariantIterNext},
	{"g_variant_iter_next_value", &VariantIterNextValue},
	{"g_variant_lookup", &VariantLookup},
	{"g_variant_lookup_value", &VariantLookupValue},
	{"g_variant_n_children", &VariantNChildren},
	{"g_variant_new", &VariantNew},
	{"g_variant_new_array", &VariantNewArray},
	{"g_variant_new_boolean", &VariantNewBoolean},
	{"g_variant_new_byte", &VariantNewByte},
	{"g_variant_new_bytestring", &VariantNewBytestring},
	{"g_variant_new_bytestring_array", &VariantNewBytestringArray},
	{"g_variant_new_dict_entry", &VariantNewDictEntry},
	{"g_variant_new_double", &VariantNewDouble},
	{"g_variant_new_from_data", &VariantNewFromData},
	{"g_variant_new_handle", &VariantNewHandle},
	{"g_variant_new_int16", &VariantNewInt16},
	{"g_variant_new_int32", &VariantNewInt32},
	{"g_variant_new_int64", &VariantNewInt64},
	{"g_variant_new_maybe", &VariantNewMaybe},
	{"g_variant_new_object_path", &VariantNewObjectPath},
	{"g_variant_new_parsed", &VariantNewParsed},
	{"g_variant_new_parsed_va", &VariantNewParsedVa},
	{"g_variant_new_signature", &VariantNewSignature},
	{"g_variant_new_string", &VariantNewString},
	{"g_variant_new_strv", &VariantNewStrv},
	{"g_variant_new_tuple", &VariantNewTuple},
	{"g_variant_new_uint16", &VariantNewUint16},
	{"g_variant_new_uint32", &VariantNewUint32},
	{"g_variant_new_uint64", &VariantNewUint64},
	{"g_variant_new_va", &VariantNewVa},
	{"g_variant_new_variant", &VariantNewVariant},
	{"g_variant_parse", &VariantParse},
	{"g_variant_parser_get_error_quark", &VariantParserGetErrorQuark},
	{"g_variant_print", &VariantPrint},
	{"g_variant_print_string", &VariantPrintString},
	{"g_variant_ref", &VariantRef},
	{"g_variant_ref_sink", &VariantRefSink},
	// Undocumented {"g_variant_serialised_byteswap", &VariantSerialisedByteswap},
	// Undocumented {"g_variant_serialised_get_child", &VariantSerialisedGetChild},
	// Undocumented {"g_variant_serialised_is_normal", &VariantSerialisedIsNormal},
	// Undocumented {"g_variant_serialised_n_children", &VariantSerialisedNChildren},
	// Undocumented {"g_variant_serialiser_is_object_path", &VariantSerialiserIsObjectPath},
	// Undocumented {"g_variant_serialiser_is_signature", &VariantSerialiserIsSignature},
	// Undocumented {"g_variant_serialiser_is_string", &VariantSerialiserIsString},
	// Undocumented {"g_variant_serialiser_needed_size", &VariantSerialiserNeededSize},
	// Undocumented {"g_variant_serialiser_serialise", &VariantSerialiserSerialise},
	{"g_variant_store", &VariantStore},
	{"g_variant_type_checked_", &VariantTypeChecked},
	{"g_variant_type_copy", &VariantTypeCopy},
	{"g_variant_type_dup_string", &VariantTypeDupString},
	{"g_variant_type_element", &VariantTypeElement},
	{"g_variant_type_equal", &VariantTypeEqual},
	{"g_variant_type_first", &VariantTypeFirst},
	{"g_variant_type_free", &VariantTypeFree},
	{"g_variant_type_get_string_length", &VariantTypeGetStringLength},
	{"g_variant_type_hash", &VariantTypeHash},
	// Undocumented {"g_variant_type_info_assert_no_infos", &VariantTypeInfoAssertNoInfos},
	// Undocumented {"g_variant_type_info_element", &VariantTypeInfoElement},
	// Undocumented {"g_variant_type_info_get", &VariantTypeInfoGet},
	// Undocumented {"g_variant_type_info_get_type_string", &VariantTypeInfoGetTypeString},
	// Undocumented {"g_variant_type_info_member_info", &VariantTypeInfoMemberInfo},
	// Undocumented {"g_variant_type_info_n_members", &VariantTypeInfoNMembers},
	// Undocumented {"g_variant_type_info_query", &VariantTypeInfoQuery},
	// Undocumented {"g_variant_type_info_query_element", &VariantTypeInfoQueryElement},
	// Undocumented {"g_variant_type_info_ref", &VariantTypeInfoRef},
	// Undocumented {"g_variant_type_info_unref", &VariantTypeInfoUnref},
	{"g_variant_type_is_array", &VariantTypeIsArray},
	{"g_variant_type_is_basic", &VariantTypeIsBasic},
	{"g_variant_type_is_container", &VariantTypeIsContainer},
	{"g_variant_type_is_definite", &VariantTypeIsDefinite},
	{"g_variant_type_is_dict_entry", &VariantTypeIsDictEntry},
	{"g_variant_type_is_maybe", &VariantTypeIsMaybe},
	{"g_variant_type_is_subtype_of", &VariantTypeIsSubtypeOf},
	{"g_variant_type_is_tuple", &VariantTypeIsTuple},
	{"g_variant_type_is_variant", &VariantTypeIsVariant},
	{"g_variant_type_key", &VariantTypeKey},
	{"g_variant_type_n_items", &VariantTypeNItems},
	{"g_variant_type_new", &VariantTypeNew},
	{"g_variant_type_new_array", &VariantTypeNewArray},
	{"g_variant_type_new_dict_entry", &VariantTypeNewDictEntry},
	{"g_variant_type_new_maybe", &VariantTypeNewMaybe},
	{"g_variant_type_new_tuple", &VariantTypeNewTuple},
	{"g_variant_type_next", &VariantTypeNext},
	{"g_variant_type_peek_string", &VariantTypePeekString},
	{"g_variant_type_string_is_valid", &VariantTypeStringIsValid},
	{"g_variant_type_string_scan", &VariantTypeStringScan},
	{"g_variant_type_value", &VariantTypeValue},
	{"g_variant_unref", &VariantUnref},
	{"g_vasprintf", &Vasprintf},
	{"g_vfprintf", &Vfprintf},
	{"g_vprintf", &Vprintf},
	{"g_vsnprintf", &Vsnprintf},
	{"g_vsprintf", &Vsprintf},
	{"g_warn_message", &WarnMessage},
	{"g_win32_error_message", &Win32ErrorMessage},
	{"g_win32_ftruncate", &Win32Ftruncate},
	// Windows: _utf8 {"g_win32_get_package_installation_directory", &Win32GetPackageInstallationDirectory},
	{"g_win32_get_package_installation_directory_of_module", &Win32GetPackageInstallationDirectoryOfModule},
	{"g_win32_get_package_installation_directory_utf8", &Win32GetPackageInstallationDirectory},
	// Windows: _utf8 {"g_win32_get_package_installation_subdirectory", &Win32GetPackageInstallationSubdirectory},
	{"g_win32_get_package_installation_subdirectory_utf8", &Win32GetPackageInstallationSubdirectory},
	{"g_win32_get_system_data_dirs_for_module", &Win32GetSystemDataDirsForModule},
	{"g_win32_get_windows_version", &Win32GetWindowsVersion},
	{"g_win32_getlocale", &Win32Getlocale},
	{"g_win32_locale_filename_from_utf8", &Win32LocaleFilenameFromUtf8},
	{"glib_check_version", &GlibCheckVersion},
	// Undocumented {"glib_gettext", &GlibGettext},
	// Undocumented {"glib_on_error_halt", &GlibOnErrorHalt},
	// Undocumented {"glib_pgettext", &GlibPgettext},
}

var dllThread = "libgthread-2.0-0.dll"

var apiListThread = outside.Apis{
	{"g_thread_init", &ThreadInit},
	{"g_thread_init_with_errorcheck_mutexes", &ThreadInitWithErrorcheckMutexes},
}

var dataList = outside.Data{
	{"g_ascii_table", (*uint16)(nil)},
	{"g_child_watch_funcs", (*O.SourceFuncs)(nil)},
	{"g_idle_funcs", (*O.SourceFuncs)(nil)},
	{"g_io_watch_funcs", (*O.SourceFuncs)(nil)},
	{"g_mem_gc_friendly", (*T.Gboolean)(nil)},
	{"g_test_config_vars", (*T.GTestConfig)(nil)},
	{"g_thread_functions_for_glib_use", (*T.GThreadFunctions)(nil)},
	{"g_thread_gettime", (*T.G_thread_gettime)(nil)},
	{"g_thread_use_default_impl", (*T.Gboolean)(nil)},
	{"g_threads_got_initialized", (*T.Gboolean)(nil)},
	{"g_timeout_funcs", (*O.SourceFuncs)(nil)},
	{"g_utf8_skip", (*T.Gchar)(nil)},
	{"glib_binary_age", (*uint)(nil)},
	{"glib_interface_age", (*uint)(nil)},
	{"glib_major_version", (*uint)(nil)},
	{"glib_mem_profiler_table", (*T.GMemVTable)(nil)},
	{"glib_micro_version", (*uint)(nil)},
	{"glib_minor_version", (*uint)(nil)},
}
