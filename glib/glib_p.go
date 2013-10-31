// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type PtrArray struct {
	Data *T.Gpointer
	Len  uint
}

var (
	PtrArrayNew             func() *PtrArray
	PtrArrayNewWithFreeFunc func(elementFreeFunc T.GDestroyNotify) *PtrArray
	PtrArraySizedNew        func(reservedSize uint) *PtrArray

	ptrArrayAdd             func(p *PtrArray, data T.Gpointer)
	ptrArrayForeach         func(p *PtrArray, f T.GFunc, userData T.Gpointer)
	ptrArrayFree            func(p *PtrArray, freeSeg bool) *T.Gpointer
	ptrArrayRef             func(p *PtrArray) *PtrArray
	ptrArrayRemove          func(p *PtrArray, data T.Gpointer) bool
	ptrArrayRemoveFast      func(p *PtrArray, data T.Gpointer) bool
	ptrArrayRemoveIndex     func(p *PtrArray, index uint) T.Gpointer
	ptrArrayRemoveIndexFast func(p *PtrArray, index uint) T.Gpointer
	ptrArrayRemoveRange     func(p *PtrArray, index uint, length uint)
	ptrArraySetFreeFunc     func(p *PtrArray, elementFreeFunc T.GDestroyNotify)
	ptrArraySetSize         func(p *PtrArray, length int)
	ptrArraySort            func(p *PtrArray, compareFunc T.GCompareFunc)
	ptrArraySortWithData    func(p *PtrArray, compareFunc T.GCompareDataFunc, userData T.Gpointer)
	ptrArrayUnref           func(p *PtrArray)
)

func (p *PtrArray) Add(data T.Gpointer)                    { ptrArrayAdd(p, data) }
func (p *PtrArray) Foreach(f T.GFunc, userData T.Gpointer) { ptrArrayForeach(p, f, userData) }
func (p *PtrArray) Free(freeSeg bool) *T.Gpointer          { return ptrArrayFree(p, freeSeg) }
func (p *PtrArray) Ref() *PtrArray                         { return ptrArrayRef(p) }
func (p *PtrArray) Remove(data T.Gpointer) bool            { return ptrArrayRemove(p, data) }
func (p *PtrArray) RemoveFast(data T.Gpointer) bool        { return ptrArrayRemoveFast(p, data) }
func (p *PtrArray) RemoveIndex(index uint) T.Gpointer      { return ptrArrayRemoveIndex(p, index) }
func (p *PtrArray) RemoveIndexFast(index uint) T.Gpointer  { return ptrArrayRemoveIndexFast(p, index) }
func (p *PtrArray) RemoveRange(index, length uint)         { ptrArrayRemoveRange(p, index, length) }
func (p *PtrArray) SetFreeFunc(elementFreeFunc T.GDestroyNotify) {
	ptrArraySetFreeFunc(p, elementFreeFunc)
}
func (p *PtrArray) SetSize(length int)              { ptrArraySetSize(p, length) }
func (p *PtrArray) Sort(compareFunc T.GCompareFunc) { ptrArraySort(p, compareFunc) }
func (p *PtrArray) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) {
	ptrArraySortWithData(p, compareFunc, userData)
}
func (p *PtrArray) Unref() { ptrArrayUnref(p) }

type PatternSpec struct{}

var (
	PatternSpecNew func(pattern string) *PatternSpec

	PatternMatchSimple func(pattern, str string) bool

	patternMatch       func(p *PatternSpec, stringLength uint, str, stringReversed string) bool
	patternMatchString func(p *PatternSpec, str string) bool
	patternSpecEqual   func(p *PatternSpec, p2 *PatternSpec) bool
	patternSpecFree    func(p *PatternSpec)
)

func (p *PatternSpec) Match(stringLength uint, str, stringReversed string) bool {
	return patternMatch(p, stringLength, str, stringReversed)
}
func (p *PatternSpec) MatchString(str string) bool { return patternMatchString(p, str) }
func (p *PatternSpec) Equal(p2 *PatternSpec) bool  { return patternSpecEqual(p, p2) }
func (p *PatternSpec) Free()                       { patternSpecFree(p) }
