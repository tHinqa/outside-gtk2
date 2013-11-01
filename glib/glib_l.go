// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type List struct {
	Data T.Gpointer
	Next *List
	Prev *List
}

var (
	ListAlloc         func() *List
	ListPopAllocator  func()
	ListPushAllocator func(allocator T.Gpointer)

	ListAppend               func(l *List, data T.Gpointer) *List
	ListConcat               func(l *List, l2 *List) *List
	ListCopy                 func(l *List) *List
	ListDeleteLink           func(l *List, link *List) *List
	ListFind                 func(l *List, data T.Gconstpointer) *List
	ListFindCustom           func(l *List, data T.Gconstpointer, f T.GCompareFunc) *List
	ListFirst                func(l *List) *List
	ListForeach              func(l *List, f T.GFunc, userData T.Gpointer)
	ListFree                 func(l *List)
	ListFree1                func(l *List)
	ListFreeFull             func(l *List, freeFunc T.GDestroyNotify)
	ListIndex                func(l *List, data T.Gconstpointer) int
	ListInsert               func(l *List, data T.Gpointer, position int) *List
	ListInsertBefore         func(l *List, sibling *List, data T.Gpointer) *List
	ListInsertSorted         func(l *List, data T.Gpointer, f T.GCompareFunc) *List
	ListInsertSortedWithData func(l *List, data T.Gpointer, fGCompareDataFunc, userData T.Gpointer) *List
	ListLast                 func(l *List) *List
	ListLength               func(l *List) uint
	ListNth                  func(l *List, n uint) *List
	ListNthData              func(l *List, n uint) T.Gpointer
	ListNthPrev              func(l *List, n uint) *List
	ListPosition             func(l *List, llink *List) int
	ListPrepend              func(l *List, data T.Gpointer) *List
	ListRemove               func(l *List, data T.Gconstpointer) *List
	ListRemoveAll            func(l *List, data T.Gconstpointer) *List
	ListRemoveLink           func(l *List, llink *List) *List
	ListReverse              func(l *List) *List
	ListSort                 func(l *List, compareFunc T.GCompareFunc) *List
	ListSortWithData         func(l *List, compareFunc T.GCompareDataFunc, userData T.Gpointer) *List
)

func (l *List) Append(data T.Gpointer) *List    { return ListAppend(l, data) }
func (l *List) Concat(l2 *List) *List           { return ListConcat(l, l2) }
func (l *List) Copy() *List                     { return ListCopy(l) }
func (l *List) DeleteLink(link *List) *List     { return ListDeleteLink(l, link) }
func (l *List) Find(data T.Gconstpointer) *List { return ListFind(l, data) }
func (l *List) FindCustom(data T.Gconstpointer, f T.GCompareFunc) *List {
	return ListFindCustom(l, data, f)
}
func (l *List) First() *List                               { return ListFirst(l) }
func (l *List) Foreach(f T.GFunc, userData T.Gpointer)     { ListForeach(l, f, userData) }
func (l *List) Free()                                      { ListFree(l) }
func (l *List) Free1()                                     { ListFree1(l) }
func (l *List) FreeFull(freeFunc T.GDestroyNotify)         { ListFreeFull(l, freeFunc) }
func (l *List) Index(data T.Gconstpointer) int             { return ListIndex(l, data) }
func (l *List) Insert(data T.Gpointer, position int) *List { return ListInsert(l, data, position) }
func (l *List) InsertBefore(sibling *List, data T.Gpointer) *List {
	return ListInsertBefore(l, sibling, data)
}
func (l *List) InsertSorted(data T.Gpointer, f T.GCompareFunc) *List {
	return ListInsertSorted(l, data, f)
}
func (l *List) InsertSortedWithData(data T.Gpointer, fGCompareDataFunc, userData T.Gpointer) *List {
	return ListInsertSortedWithData(l, data, fGCompareDataFunc, userData)
}
func (l *List) Last() *List                           { return ListLast(l) }
func (l *List) Length() uint                          { return ListLength(l) }
func (l *List) Nth(n uint) *List                      { return ListNth(l, n) }
func (l *List) NthData(n uint) T.Gpointer             { return ListNthData(l, n) }
func (l *List) NthPrev(n uint) *List                  { return ListNthPrev(l, n) }
func (l *List) Position(llink *List) int              { return ListPosition(l, llink) }
func (l *List) Prepend(data T.Gpointer) *List         { return ListPrepend(l, data) }
func (l *List) Remove(data T.Gconstpointer) *List     { return ListRemove(l, data) }
func (l *List) RemoveAll(data T.Gconstpointer) *List  { return ListRemoveAll(l, data) }
func (l *List) RemoveLink(llink *List) *List          { return ListRemoveLink(l, llink) }
func (l *List) Reverse() *List                        { return ListReverse(l) }
func (l *List) Sort(compareFunc T.GCompareFunc) *List { return ListSort(l, compareFunc) }
func (l *List) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) *List {
	return ListSortWithData(l, compareFunc, userData)
}
