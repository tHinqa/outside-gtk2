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

	AllocatorFree func(a *Allocator)
)

func (a *Allocator) Free() { AllocatorFree(a) }

type Array struct {
	Data *T.Gchar
	Leng uint
}

var (
	ArrayNew      func(zeroTerminated, clear bool, elementSize uint) *Array
	ArraySizedNew func(zeroTerminated, clear bool, elementSize, reservedSize uint) *Array

	ArrayFree            func(a *Array, freeSegment bool) string
	ArrayRef             func(a *Array) *Array
	ArrayUnref           func(a *Array)
	ArrayGetElementSize  func(a *Array) uint
	ArrayAppendVals      func(a *Array, data T.Gconstpointer, leng uint) *Array
	ArrayPrependVals     func(a *Array, data T.Gconstpointer, leng uint) *Array
	ArrayInsertVals      func(a *Array, index uint, data T.Gconstpointer, leng uint) *Array
	ArraySetSize         func(a *Array, length uint) *Array
	ArrayRemoveIndex     func(a *Array, index uint) *Array
	ArrayRemoveIndexFast func(a *Array, index uint) *Array
	ArrayRemoveRange     func(a *Array, index, length uint) *Array
	ArraySort            func(a *Array, compareFunc T.GCompareFunc)
	ArraySortWithData    func(a *Array, compareFunc T.GCompareDataFunc, userData T.Gpointer)
)

func (a *Array) Free(freeSegment bool) string { return ArrayFree(a, freeSegment) }
func (a *Array) Ref() *Array                  { return ArrayRef(a) }
func (a *Array) Unref()                       { ArrayUnref(a) }
func (a *Array) GetElementSize() uint         { return ArrayGetElementSize(a) }
func (a *Array) AppendVals(data T.Gconstpointer, leng uint) *Array {
	return ArrayAppendVals(a, data, leng)
}
func (a *Array) PrependVals(data T.Gconstpointer, leng uint) *Array {
	return ArrayPrependVals(a, data, leng)
}
func (a *Array) InsertVals(index uint, data T.Gconstpointer, leng uint) *Array {
	return ArrayInsertVals(a, index, data, leng)
}
func (a *Array) SetSize(length uint) *Array            { return ArraySetSize(a, length) }
func (a *Array) RemoveIndex(index uint) *Array         { return ArrayRemoveIndex(a, index) }
func (a *Array) RemoveIndexFast(index uint) *Array     { return ArrayRemoveIndexFast(a, index) }
func (a *Array) RemoveRange(index, length uint) *Array { return ArrayRemoveRange(a, index, length) }
func (a *Array) Sort(compareFunc T.GCompareFunc)       { ArraySort(a, compareFunc) }
func (a *Array) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) {
	ArraySortWithData(a, compareFunc, userData)
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

	AsyncQueueLock               func(a *AsyncQueue)
	AsyncQueueUnlock             func(a *AsyncQueue)
	AsyncQueueRef                func(a *AsyncQueue) *AsyncQueue
	AsyncQueueUnref              func(a *AsyncQueue)
	AsyncQueueRefUnlocked        func(a *AsyncQueue)
	AsyncQueueUnrefAndUnlock     func(a *AsyncQueue)
	AsyncQueuePush               func(a *AsyncQueue, data T.Gpointer)
	AsyncQueuePushUnlocked       func(a *AsyncQueue, data T.Gpointer)
	AsyncQueuePushSorted         func(a *AsyncQueue, data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer)
	AsyncQueuePushSortedUnlocked func(a *AsyncQueue, data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer)
	AsyncQueuePop                func(a *AsyncQueue) T.Gpointer
	AsyncQueuePopUnlocked        func(a *AsyncQueue) T.Gpointer
	AsyncQueueTryPop             func(a *AsyncQueue) T.Gpointer
	AsyncQueueTryPopUnlocked     func(a *AsyncQueue) T.Gpointer
	AsyncQueueTimedPop           func(a *AsyncQueue, endTime *T.GTimeVal) T.Gpointer
	AsyncQueueTimedPopUnlocked   func(a *AsyncQueue, endTime *T.GTimeVal) T.Gpointer
	AsyncQueueLength             func(a *AsyncQueue) int
	AsyncQueueLengthUnlocked     func(a *AsyncQueue) int
	AsyncQueueSort               func(a *AsyncQueue, f T.GCompareDataFunc, userData T.Gpointer)
	AsyncQueueSortUnlocked       func(a *AsyncQueue, f T.GCompareDataFunc, userData T.Gpointer)
)

func (a *AsyncQueue) Lock()                        { AsyncQueueLock(a) }
func (a *AsyncQueue) Unlock()                      { AsyncQueueUnlock(a) }
func (a *AsyncQueue) Ref() *AsyncQueue             { return AsyncQueueRef(a) }
func (a *AsyncQueue) Unref()                       { AsyncQueueUnref(a) }
func (a *AsyncQueue) RefUnlocked()                 { AsyncQueueRefUnlocked(a) }
func (a *AsyncQueue) UnrefAndUnlock()              { AsyncQueueUnrefAndUnlock(a) }
func (a *AsyncQueue) Push(data T.Gpointer)         { AsyncQueuePush(a, data) }
func (a *AsyncQueue) PushUnlocked(data T.Gpointer) { AsyncQueuePushUnlocked(a, data) }
func (a *AsyncQueue) PushSorted(data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) {
	AsyncQueuePushSorted(a, data, f, userData)
}
func (a *AsyncQueue) PushSortedUnlocked(data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) {
	AsyncQueuePushSortedUnlocked(a, data, f, userData)
}
func (a *AsyncQueue) Pop() T.Gpointer                         { return AsyncQueuePop(a) }
func (a *AsyncQueue) PopUnlocked() T.Gpointer                 { return AsyncQueuePopUnlocked(a) }
func (a *AsyncQueue) TryPop() T.Gpointer                      { return AsyncQueueTryPop(a) }
func (a *AsyncQueue) TryPopUnlocked() T.Gpointer              { return AsyncQueueTryPopUnlocked(a) }
func (a *AsyncQueue) TimedPop(endTime *T.GTimeVal) T.Gpointer { return AsyncQueueTimedPop(a, endTime) }
func (a *AsyncQueue) TimedPopUnlocked(endTime *T.GTimeVal) T.Gpointer {
	return AsyncQueueTimedPopUnlocked(a, endTime)
}
func (a *AsyncQueue) Length() int                                    { return AsyncQueueLength(a) }
func (a *AsyncQueue) LengthUnlocked() int                            { return AsyncQueueLengthUnlocked(a) }
func (a *AsyncQueue) Sort(f T.GCompareDataFunc, userData T.Gpointer) { AsyncQueueSort(a, f, userData) }
func (a *AsyncQueue) SortUnlocked(f T.GCompareDataFunc, userData T.Gpointer) {
	AsyncQueueSortUnlocked(a, f, userData)
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
