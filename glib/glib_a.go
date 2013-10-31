// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Allocator struct{}

var (
	AllocatorNew func(name string, nPreallocs uint) *Allocator

	allocatorFree func(a *Allocator)
)

func (a *Allocator) Free() { allocatorFree(a) }

type Array struct {
	Data *T.Gchar
	Leng uint
}

var (
	ArrayNew      func(zeroTerminated, clear bool, elementSize uint) *Array
	ArraySizedNew func(zeroTerminated, clear bool, elementSize, reservedSize uint) *Array

	arrayFree            func(a *Array, freeSegment bool) string
	arrayRef             func(a *Array) *Array
	arrayUnref           func(a *Array)
	arrayGetElementSize  func(a *Array) uint
	arrayAppendVals      func(a *Array, data T.Gconstpointer, leng uint) *Array
	arrayPrependVals     func(a *Array, data T.Gconstpointer, leng uint) *Array
	arrayInsertVals      func(a *Array, index uint, data T.Gconstpointer, leng uint) *Array
	arraySetSize         func(a *Array, length uint) *Array
	arrayRemoveIndex     func(a *Array, index uint) *Array
	arrayRemoveIndexFast func(a *Array, index uint) *Array
	arrayRemoveRange     func(a *Array, index, length uint) *Array
	arraySort            func(a *Array, compareFunc T.GCompareFunc)
	arraySortWithData    func(a *Array, compareFunc T.GCompareDataFunc, userData T.Gpointer)
)

func (a *Array) Free(freeSegment bool) string { return arrayFree(a, freeSegment) }
func (a *Array) Ref() *Array                  { return arrayRef(a) }
func (a *Array) Unref()                       { arrayUnref(a) }
func (a *Array) GetElementSize() uint         { return arrayGetElementSize(a) }
func (a *Array) AppendVals(data T.Gconstpointer, leng uint) *Array {
	return arrayAppendVals(a, data, leng)
}
func (a *Array) PrependVals(data T.Gconstpointer, leng uint) *Array {
	return arrayPrependVals(a, data, leng)
}
func (a *Array) InsertVals(index uint, data T.Gconstpointer, leng uint) *Array {
	return arrayInsertVals(a, index, data, leng)
}
func (a *Array) SetSize(length uint) *Array            { return arraySetSize(a, length) }
func (a *Array) RemoveIndex(index uint) *Array         { return arrayRemoveIndex(a, index) }
func (a *Array) RemoveIndexFast(index uint) *Array     { return arrayRemoveIndexFast(a, index) }
func (a *Array) RemoveRange(index, length uint) *Array { return arrayRemoveRange(a, index, length) }
func (a *Array) Sort(compareFunc T.GCompareFunc)       { arraySort(a, compareFunc) }
func (a *Array) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) {
	arraySortWithData(a, compareFunc, userData)
}

var (
	AssertionMessage       func(domain, file string, line int, f, message string)
	AssertionMessageExpr   func(domain, file string, line int, f, expr string)
	AssertionMessageCmpstr func(domain, file string, line int, f, expr, arg1, cmp, arg2 string)
	AssertionMessageCmpnum func(domain, file string, line int, f, expr string, arg1 T.LongDouble, cmp string, arg2 T.LongDouble, numtype T.Char)
	AssertionMessageError  func(domain, file string, line int, f, expr string, e *T.GError, errorDomain T.GQuark, errorCode int)
)

type AsyncQueue struct{}

var (
	AsyncQueueNew     func() *AsyncQueue
	AsyncQueueNewFull func(itemFreeFunc T.GDestroyNotify) *AsyncQueue

	asyncQueueLock               func(a *AsyncQueue)
	asyncQueueUnlock             func(a *AsyncQueue)
	asyncQueueRef                func(a *AsyncQueue) *AsyncQueue
	asyncQueueUnref              func(a *AsyncQueue)
	asyncQueueRefUnlocked        func(a *AsyncQueue)
	asyncQueueUnrefAndUnlock     func(a *AsyncQueue)
	asyncQueuePush               func(a *AsyncQueue, data T.Gpointer)
	asyncQueuePushUnlocked       func(a *AsyncQueue, data T.Gpointer)
	asyncQueuePushSorted         func(a *AsyncQueue, data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer)
	asyncQueuePushSortedUnlocked func(a *AsyncQueue, data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer)
	asyncQueuePop                func(a *AsyncQueue) T.Gpointer
	asyncQueuePopUnlocked        func(a *AsyncQueue) T.Gpointer
	asyncQueueTryPop             func(a *AsyncQueue) T.Gpointer
	asyncQueueTryPopUnlocked     func(a *AsyncQueue) T.Gpointer
	asyncQueueTimedPop           func(a *AsyncQueue, endTime *T.GTimeVal) T.Gpointer
	asyncQueueTimedPopUnlocked   func(a *AsyncQueue, endTime *T.GTimeVal) T.Gpointer
	asyncQueueLength             func(a *AsyncQueue) int
	asyncQueueLengthUnlocked     func(a *AsyncQueue) int
	asyncQueueSort               func(a *AsyncQueue, f T.GCompareDataFunc, userData T.Gpointer)
	asyncQueueSortUnlocked       func(a *AsyncQueue, f T.GCompareDataFunc, userData T.Gpointer)
)

func (a *AsyncQueue) Lock()                        { asyncQueueLock(a) }
func (a *AsyncQueue) Unlock()                      { asyncQueueUnlock(a) }
func (a *AsyncQueue) Ref() *AsyncQueue             { return asyncQueueRef(a) }
func (a *AsyncQueue) Unref()                       { asyncQueueUnref(a) }
func (a *AsyncQueue) RefUnlocked()                 { asyncQueueRefUnlocked(a) }
func (a *AsyncQueue) UnrefAndUnlock()              { asyncQueueUnrefAndUnlock(a) }
func (a *AsyncQueue) Push(data T.Gpointer)         { asyncQueuePush(a, data) }
func (a *AsyncQueue) PushUnlocked(data T.Gpointer) { asyncQueuePushUnlocked(a, data) }
func (a *AsyncQueue) PushSorted(data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) {
	asyncQueuePushSorted(a, data, f, userData)
}
func (a *AsyncQueue) PushSortedUnlocked(data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) {
	asyncQueuePushSortedUnlocked(a, data, f, userData)
}
func (a *AsyncQueue) Pop() T.Gpointer                         { return asyncQueuePop(a) }
func (a *AsyncQueue) PopUnlocked() T.Gpointer                 { return asyncQueuePopUnlocked(a) }
func (a *AsyncQueue) TryPop() T.Gpointer                      { return asyncQueueTryPop(a) }
func (a *AsyncQueue) TryPopUnlocked() T.Gpointer              { return asyncQueueTryPopUnlocked(a) }
func (a *AsyncQueue) TimedPop(endTime *T.GTimeVal) T.Gpointer { return asyncQueueTimedPop(a, endTime) }
func (a *AsyncQueue) TimedPopUnlocked(endTime *T.GTimeVal) T.Gpointer {
	return asyncQueueTimedPopUnlocked(a, endTime)
}
func (a *AsyncQueue) Length() int                                    { return asyncQueueLength(a) }
func (a *AsyncQueue) LengthUnlocked() int                            { return asyncQueueLengthUnlocked(a) }
func (a *AsyncQueue) Sort(f T.GCompareDataFunc, userData T.Gpointer) { asyncQueueSort(a, f, userData) }
func (a *AsyncQueue) SortUnlocked(f T.GCompareDataFunc, userData T.Gpointer) {
	asyncQueueSortUnlocked(a, f, userData)
}

var (
	AtomicIntExchangeAndAdd         func(atomic *int, val int) int
	AtomicIntAdd                    func(atomic *int, val int)
	AtomicIntCompareAndExchange     func(atomic *int, oldval, newval int) bool
	AtomicPointerCompareAndExchange func(atomic *T.Gpointer, oldval, newval T.Gpointer) bool
	AtomicIntGet                    func(atomic *int) int
	AtomicIntSet                    func(atomic *int, newval int)
	AtomicPointerGet                func(atomic *T.Gpointer) T.Gpointer
	AtomicPointerSet                func(atomic *T.Gpointer, newval T.Gpointer)
)
