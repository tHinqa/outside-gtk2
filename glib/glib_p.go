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

	PtrArrayAdd             func(p *PtrArray, data T.Gpointer)
	PtrArrayForeach         func(p *PtrArray, f T.GFunc, userData T.Gpointer)
	PtrArrayFree            func(p *PtrArray, freeSeg bool) *T.Gpointer
	PtrArrayRef             func(p *PtrArray) *PtrArray
	PtrArrayRemove          func(p *PtrArray, data T.Gpointer) bool
	PtrArrayRemoveFast      func(p *PtrArray, data T.Gpointer) bool
	PtrArrayRemoveIndex     func(p *PtrArray, index uint) T.Gpointer
	PtrArrayRemoveIndexFast func(p *PtrArray, index uint) T.Gpointer
	PtrArrayRemoveRange     func(p *PtrArray, index uint, length uint)
	PtrArraySetFreeFunc     func(p *PtrArray, elementFreeFunc T.GDestroyNotify)
	PtrArraySetSize         func(p *PtrArray, length int)
	PtrArraySort            func(p *PtrArray, compareFunc T.GCompareFunc)
	PtrArraySortWithData    func(p *PtrArray, compareFunc T.GCompareDataFunc, userData T.Gpointer)
	PtrArrayUnref           func(p *PtrArray)
)

func (p *PtrArray) Add(data T.Gpointer)                    { PtrArrayAdd(p, data) }
func (p *PtrArray) Foreach(f T.GFunc, userData T.Gpointer) { PtrArrayForeach(p, f, userData) }
func (p *PtrArray) Free(freeSeg bool) *T.Gpointer          { return PtrArrayFree(p, freeSeg) }
func (p *PtrArray) Ref() *PtrArray                         { return PtrArrayRef(p) }
func (p *PtrArray) Remove(data T.Gpointer) bool            { return PtrArrayRemove(p, data) }
func (p *PtrArray) RemoveFast(data T.Gpointer) bool        { return PtrArrayRemoveFast(p, data) }
func (p *PtrArray) RemoveIndex(index uint) T.Gpointer      { return PtrArrayRemoveIndex(p, index) }
func (p *PtrArray) RemoveIndexFast(index uint) T.Gpointer  { return PtrArrayRemoveIndexFast(p, index) }
func (p *PtrArray) RemoveRange(index, length uint)         { PtrArrayRemoveRange(p, index, length) }
func (p *PtrArray) SetFreeFunc(elementFreeFunc T.GDestroyNotify) {
	PtrArraySetFreeFunc(p, elementFreeFunc)
}
func (p *PtrArray) SetSize(length int)              { PtrArraySetSize(p, length) }
func (p *PtrArray) Sort(compareFunc T.GCompareFunc) { PtrArraySort(p, compareFunc) }
func (p *PtrArray) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) {
	PtrArraySortWithData(p, compareFunc, userData)
}
func (p *PtrArray) Unref() { PtrArrayUnref(p) }

type PatternSpec struct{}

var (
	PatternSpecNew func(pattern string) *PatternSpec

	PatternMatchSimple func(pattern, str string) bool

	PatternMatch       func(p *PatternSpec, stringLength uint, str, stringReversed string) bool
	PatternMatchString func(p *PatternSpec, str string) bool
	PatternSpecEqual   func(p *PatternSpec, p2 *PatternSpec) bool
	PatternSpecFree    func(p *PatternSpec)
)

func (p *PatternSpec) Match(stringLength uint, str, stringReversed string) bool {
	return PatternMatch(p, stringLength, str, stringReversed)
}
func (p *PatternSpec) MatchString(str string) bool { return PatternMatchString(p, str) }
func (p *PatternSpec) Equal(p2 *PatternSpec) bool  { return PatternSpecEqual(p, p2) }
func (p *PatternSpec) Free()                       { PatternSpecFree(p) }
