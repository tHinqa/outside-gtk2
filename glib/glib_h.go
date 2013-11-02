// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type (
	HashFunc func(key T.Gconstpointer) uint
	HFunc    func(key, value, userData T.Gpointer)
	HRFunc   func(key, value, userData T.Gpointer) bool
)

type HashTable struct{}

var (
	HashTableNew     func(hashFunc HashFunc, keyEqualFunc T.GEqualFunc) *HashTable
	HashTableNewFull func(hashFunc HashFunc, keyEqualFunc T.GEqualFunc, keyDestroyFunc O.DestroyNotify, valueDestroyFunc O.DestroyNotify) *HashTable

	HashTableDestroy        func(h *HashTable)
	HashTableFind           func(h *HashTable, predicate HRFunc, userData T.Gpointer) T.Gpointer
	HashTableForeach        func(h *HashTable, f HFunc, userData T.Gpointer)
	HashTableForeachRemove  func(h *HashTable, f HRFunc, userData T.Gpointer) uint
	HashTableForeachSteal   func(h *HashTable, f HRFunc, userData T.Gpointer) uint
	HashTableGetKeys        func(h *HashTable) *List
	HashTableGetValues      func(h *HashTable) *List
	HashTableInsert         func(h *HashTable, key T.Gpointer, value T.Gpointer)
	HashTableLookup         func(h *HashTable, key T.Gconstpointer) T.Gpointer
	HashTableLookupExtended func(h *HashTable, lookupKey T.Gconstpointer, origKey, value *T.Gpointer) bool
	HashTableRef            func(h *HashTable) *HashTable
	HashTableRemove         func(h *HashTable, key T.Gconstpointer) bool
	HashTableRemoveAll      func(h *HashTable)
	HashTableReplace        func(h *HashTable, key T.Gpointer, value T.Gpointer)
	HashTableSize           func(h *HashTable) uint
	HashTableSteal          func(h *HashTable, key T.Gconstpointer) bool
	HashTableStealAll       func(h *HashTable)
	HashTableUnref          func(h *HashTable)
)

func (h *HashTable) Destroy() { HashTableDestroy(h) }
func (h *HashTable) Find(predicate HRFunc, userData T.Gpointer) T.Gpointer {
	return HashTableFind(h, predicate, userData)
}
func (h *HashTable) Foreach(f HFunc, userData T.Gpointer) { HashTableForeach(h, f, userData) }
func (h *HashTable) ForeachRemove(f HRFunc, userData T.Gpointer) uint {
	return HashTableForeachRemove(h, f, userData)
}
func (h *HashTable) ForeachSteal(f HRFunc, userData T.Gpointer) uint {
	return HashTableForeachSteal(h, f, userData)
}
func (h *HashTable) GetKeys() *List                          { return HashTableGetKeys(h) }
func (h *HashTable) GetValues() *List                        { return HashTableGetValues(h) }
func (h *HashTable) Insert(key T.Gpointer, value T.Gpointer) { HashTableInsert(h, key, value) }
func (h *HashTable) Lookup(key T.Gconstpointer) T.Gpointer   { return HashTableLookup(h, key) }
func (h *HashTable) LookupExtended(lookupKey T.Gconstpointer, origKey, value *T.Gpointer) bool {
	return HashTableLookupExtended(h, lookupKey, origKey, value)
}
func (h *HashTable) Ref() *HashTable                          { return HashTableRef(h) }
func (h *HashTable) Remove(key T.Gconstpointer) bool          { return HashTableRemove(h, key) }
func (h *HashTable) RemoveAll()                               { HashTableRemoveAll(h) }
func (h *HashTable) Replace(key T.Gpointer, value T.Gpointer) { HashTableReplace(h, key, value) }
func (h *HashTable) Size() uint                               { return HashTableSize(h) }
func (h *HashTable) Steal(key T.Gconstpointer) bool           { return HashTableSteal(h, key) }
func (h *HashTable) StealAll()                                { HashTableStealAll(h) }
func (h *HashTable) Unref()                                   { HashTableUnref(h) }

type HashTableIter struct {
	Dummy1 T.Gpointer
	Dummy2 T.Gpointer
	Dummy3 T.Gpointer
	Dummy4 int
	Dummy5 T.Gboolean
	Dummy6 T.Gpointer
}

var (
	HashTableIterGetHashTable func(h *HashTableIter) *HashTable
	HashTableIterInit         func(h *HashTableIter, hashTable *HashTable)
	HashTableIterNext         func(h *HashTableIter, key *T.Gpointer, value *T.Gpointer) bool
	HashTableIterRemove       func(h *HashTableIter)
	HashTableIterSteal        func(h *HashTableIter)
)

func (h *HashTableIter) GetHashTable() *HashTable  { return HashTableIterGetHashTable(h) }
func (h *HashTableIter) Init(hashTable *HashTable) { HashTableIterInit(h, hashTable) }
func (h *HashTableIter) Next(key *T.Gpointer, value *T.Gpointer) bool {
	return HashTableIterNext(h, key, value)
}
func (h *HashTableIter) Remove() { HashTableIterRemove(h) }
func (h *HashTableIter) Steal()  { HashTableIterSteal(h) }

type (
	Hook struct {
		Data     T.Gpointer
		Next     *Hook
		Prev     *Hook
		RefCount uint
		HookId   T.Gulong
		Flags    uint
		Fnc      T.Gpointer
		Destroy  O.DestroyNotify
	}

	HookList struct {
		SeqId T.Gulong
		Bits  uint
		// HookSize : 16
		// IsSetup : 1
		Hooks        *Hook
		Dummy3       T.Gpointer
		FinalizeHook HookFinalizeFunc
		Dummy        [2]T.Gpointer
	}

	HookCheckMarshaller func(
		hook *Hook, marshalData T.Gpointer) T.Gboolean

	HookCompareFunc func(newHook, sibling *Hook) int

	HookFinalizeFunc func(hookList *HookList, hook *Hook)

	HookFindFunc func(hook *Hook, data T.Gpointer) T.Gboolean

	HookMarshaller func(hook *Hook, marshalData T.Gpointer)
)

var (
	HookCompareIds func(newHook *Hook, sibling *Hook) int

	HookAlloc            func(h *HookList) *Hook
	HookDestroy          func(h *HookList, hookId T.Gulong) bool
	HookDestroyLink      func(h *HookList, hook *Hook)
	HookFind             func(h *HookList, needValids bool, f HookFindFunc, data T.Gpointer) *Hook
	HookFindData         func(h *HookList, needValids bool, data T.Gpointer) *Hook
	HookFindFunc_        func(h *HookList, needValids bool, f T.Gpointer) *Hook
	HookFindFuncData     func(h *HookList, needValids bool, f T.Gpointer, data T.Gpointer) *Hook
	HookFirstValid       func(h *HookList, mayBeInCall bool) *Hook
	HookFree             func(h *HookList, hook *Hook)
	HookGet              func(h *HookList, hookId T.Gulong) *Hook
	HookInsertBefore     func(h *HookList, sibling, hook *Hook)
	HookInsertSorted     func(h *HookList, hook *Hook, f HookCompareFunc)
	HookListClear        func(h *HookList)
	HookListInit         func(h *HookList, hookSize uint)
	HookListInvoke       func(h *HookList, mayRecurse bool)
	HookListInvokeCheck  func(h *HookList, mayRecurse bool)
	HookListMarshal      func(h *HookList, mayRecurse bool, marshaller HookMarshaller, marshalData T.Gpointer)
	HookListMarshalCheck func(h *HookList, mayRecurse bool, marshaller HookCheckMarshaller, marshalData T.Gpointer)
	HookNextValid        func(h *HookList, hook *Hook, mayBeInCall bool) *Hook
	HookPrepend          func(h *HookList, hook *Hook)
	HookRef              func(h *HookList, hook *Hook) *Hook
	HookUnref            func(h *HookList, hook *Hook)
)

func (h *HookList) Alloc() *Hook                 { return HookAlloc(h) }
func (h *HookList) Destroy(hookId T.Gulong) bool { return HookDestroy(h, hookId) }
func (h *HookList) DestroyLink(hook *Hook)       { HookDestroyLink(h, hook) }
func (h *HookList) Find(needValids bool, f HookFindFunc, data T.Gpointer) *Hook {
	return HookFind(h, needValids, f, data)
}
func (h *HookList) FindData(needValids bool, data T.Gpointer) *Hook {
	return HookFindData(h, needValids, data)
}
func (h *HookList) FindFunc_(needValids bool, f T.Gpointer) *Hook {
	return HookFindFunc_(h, needValids, f)
}
func (h *HookList) FindFuncData(needValids bool, f T.Gpointer, data T.Gpointer) *Hook {
	return HookFindFuncData(h, needValids, f, data)
}
func (h *HookList) FirstValid(mayBeInCall bool) *Hook          { return HookFirstValid(h, mayBeInCall) }
func (h *HookList) Free(hook *Hook)                            { HookFree(h, hook) }
func (h *HookList) Get(hookId T.Gulong) *Hook                  { return HookGet(h, hookId) }
func (h *HookList) InsertBefore(sibling, hook *Hook)           { HookInsertBefore(h, sibling, hook) }
func (h *HookList) InsertSorted(hook *Hook, f HookCompareFunc) { HookInsertSorted(h, hook, f) }
func (h *HookList) Clear()                                     { HookListClear(h) }
func (h *HookList) Init(hookSize uint)                         { HookListInit(h, hookSize) }
func (h *HookList) Invoke(mayRecurse bool)                     { HookListInvoke(h, mayRecurse) }
func (h *HookList) InvokeCheck(mayRecurse bool)                { HookListInvokeCheck(h, mayRecurse) }
func (h *HookList) Marshal(mayRecurse bool, marshaller HookMarshaller, marshalData T.Gpointer) {
	HookListMarshal(h, mayRecurse, marshaller, marshalData)
}
func (h *HookList) MarshalCheck(mayRecurse bool, marshaller HookCheckMarshaller, marshalData T.Gpointer) {
	HookListMarshalCheck(h, mayRecurse, marshaller, marshalData)
}
func (h *HookList) NextValid(hook *Hook, mayBeInCall bool) *Hook {
	return HookNextValid(h, hook, mayBeInCall)
}
func (h *HookList) Prepend(hook *Hook)   { HookPrepend(h, hook) }
func (h *HookList) Ref(hook *Hook) *Hook { return HookRef(h, hook) }
func (h *HookList) Unref(hook *Hook)     { HookUnref(h, hook) }
