// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Quark T.Quark // chicken/egg glib/gobject

var (
	QuarkFromStaticString func(str string) Quark
	QuarkFromString       func(str string) Quark
	QuarkTryString        func(str string) Quark

	QuarkToString func(q Quark) string
)

func (q Quark) ToString() string { return QuarkToString(q) }

type Queue struct {
	Head   *List
	Tail   *List
	Length uint
}

var (
	QueueNew func() *Queue

	QueueClear        func(q *Queue)
	QueueCopy         func(q *Queue) *Queue
	QueueDeleteLink   func(q *Queue, link *List)
	QueueFind         func(q *Queue, data T.Gconstpointer) *List
	QueueFindCustom   func(q *Queue, data T.Gconstpointer, f T.GCompareFunc) *List
	QueueForeach      func(q *Queue, f T.GFunc, userData T.Gpointer)
	QueueFree         func(q *Queue)
	QueueGetLength    func(q *Queue) uint
	QueueIndex        func(q *Queue, data T.Gconstpointer) int
	QueueInit         func(q *Queue)
	QueueInsertAfter  func(q *Queue, sibling *List, data T.Gpointer)
	QueueInsertBefore func(q *Queue, sibling *List, data T.Gpointer)
	QueueInsertSorted func(q *Queue, data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer)
	QueueIsEmpty      func(q *Queue) bool
	QueueLinkIndex    func(q *Queue, link *List) int
	QueuePeekHead     func(q *Queue) T.Gpointer
	QueuePeekHeadLink func(q *Queue) *List
	QueuePeekNth      func(q *Queue, n uint) T.Gpointer
	QueuePeekNthLink  func(q *Queue, n uint) *List
	QueuePeekTail     func(q *Queue) T.Gpointer
	QueuePeekTailLink func(q *Queue) *List
	QueuePopHead      func(q *Queue) T.Gpointer
	QueuePopHeadLink  func(q *Queue) *List
	QueuePopNth       func(q *Queue, n uint) T.Gpointer
	QueuePopNthLink   func(q *Queue, n uint) *List
	QueuePopTail      func(q *Queue) T.Gpointer
	QueuePopTailLink  func(q *Queue) *List
	QueuePushHead     func(q *Queue, data T.Gpointer)
	QueuePushHeadLink func(q *Queue, link *List)
	QueuePushNth      func(q *Queue, data T.Gpointer, n int)
	QueuePushNthLink  func(q *Queue, n int, link *List)
	QueuePushTail     func(q *Queue, data T.Gpointer)
	QueuePushTailLink func(q *Queue, link *List)
	QueueRemove       func(q *Queue, data T.Gconstpointer)
	QueueRemoveAll    func(q *Queue, data T.Gconstpointer)
	QueueReverse      func(q *Queue)
	QueueSort         func(q *Queue, compareFunc T.GCompareDataFunc, userData T.Gpointer)
	QueueUnlink       func(q *Queue, link *List)
)

func (q *Queue) Clear()                          { QueueClear(q) }
func (q *Queue) Copy() *Queue                    { return QueueCopy(q) }
func (q *Queue) DeleteLink(link *List)           { QueueDeleteLink(q, link) }
func (q *Queue) Find(data T.Gconstpointer) *List { return QueueFind(q, data) }
func (q *Queue) FindCustom(data T.Gconstpointer, f T.GCompareFunc) *List {
	return QueueFindCustom(q, data, f)
}
func (q *Queue) Foreach(f T.GFunc, userData T.Gpointer)      { QueueForeach(q, f, userData) }
func (q *Queue) Free()                                       { QueueFree(q) }
func (q *Queue) GetLength() uint                             { return QueueGetLength(q) }
func (q *Queue) Index(data T.Gconstpointer) int              { return QueueIndex(q, data) }
func (q *Queue) Init()                                       { QueueInit(q) }
func (q *Queue) InsertAfter(sibling *List, data T.Gpointer)  { QueueInsertAfter(q, sibling, data) }
func (q *Queue) InsertBefore(sibling *List, data T.Gpointer) { QueueInsertBefore(q, sibling, data) }
func (q *Queue) InsertSorted(data T.Gpointer, f T.GCompareDataFunc, userData T.Gpointer) {
	QueueInsertSorted(q, data, f, userData)
}
func (q *Queue) IsEmpty() bool                  { return QueueIsEmpty(q) }
func (q *Queue) LinkIndex(link *List) int       { return QueueLinkIndex(q, link) }
func (q *Queue) PeekHead() T.Gpointer           { return QueuePeekHead(q) }
func (q *Queue) PeekHeadLink() *List            { return QueuePeekHeadLink(q) }
func (q *Queue) PeekNth(n uint) T.Gpointer      { return QueuePeekNth(q, n) }
func (q *Queue) PeekNthLink(n uint) *List       { return QueuePeekNthLink(q, n) }
func (q *Queue) PeekTail() T.Gpointer           { return QueuePeekTail(q) }
func (q *Queue) PeekTailLink() *List            { return QueuePeekTailLink(q) }
func (q *Queue) PopHead() T.Gpointer            { return QueuePopHead(q) }
func (q *Queue) PopHeadLink() *List             { return QueuePopHeadLink(q) }
func (q *Queue) PopNth(n uint) T.Gpointer       { return QueuePopNth(q, n) }
func (q *Queue) PopNthLink(n uint) *List        { return QueuePopNthLink(q, n) }
func (q *Queue) PopTail() T.Gpointer            { return QueuePopTail(q) }
func (q *Queue) PopTailLink() *List             { return QueuePopTailLink(q) }
func (q *Queue) PushHead(data T.Gpointer)       { QueuePushHead(q, data) }
func (q *Queue) PushHeadLink(link *List)        { QueuePushHeadLink(q, link) }
func (q *Queue) PushNth(data T.Gpointer, n int) { QueuePushNth(q, data, n) }
func (q *Queue) PushNthLink(n int, link *List)  { QueuePushNthLink(q, n, link) }
func (q *Queue) PushTail(data T.Gpointer)       { QueuePushTail(q, data) }
func (q *Queue) PushTailLink(link *List)        { QueuePushTailLink(q, link) }
func (q *Queue) Remove(data T.Gconstpointer)    { QueueRemove(q, data) }
func (q *Queue) RemoveAll(data T.Gconstpointer) { QueueRemoveAll(q, data) }
func (q *Queue) Reverse()                       { QueueReverse(q) }
func (q *Queue) Sort(compareFunc T.GCompareDataFunc, userData T.Gpointer) {
	QueueSort(q, compareFunc, userData)
}
func (q *Queue) Unlink(link *List) { QueueUnlink(q, link) }
