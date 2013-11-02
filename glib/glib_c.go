// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	T "github.com/tHinqa/outside-gtk2/types"
)

type (
	Cache            struct{}
	CacheDestroyFunc func(value T.Gpointer)
	CacheDupFunc     func(value T.Gpointer) T.Gpointer
	CacheNewFunc     func(key T.Gpointer) T.Gpointer
)

var (
	CacheNew func(valueNewFunc CacheNewFunc,
		valueDestroyFunc CacheDestroyFunc,
		keyDupFunc CacheDupFunc,
		keyDestroyFunc CacheDestroyFunc,
		hashKeyFunc HashFunc,
		hashValueFunc HashFunc,
		keyEqualFunc T.GEqualFunc) *Cache

	CacheDestroy      func(c *Cache)
	CacheInsert       func(c *Cache, key T.Gpointer) T.Gpointer
	CacheKeyForeach   func(c *Cache, f HFunc, userData T.Gpointer)
	CacheRemove       func(c *Cache, value T.Gconstpointer)
	CacheValueForeach func(c *Cache, f HFunc, userData T.Gpointer)
)

func (c *Cache) Destroy()                                  { CacheDestroy(c) }
func (c *Cache) Insert(key T.Gpointer) T.Gpointer          { return CacheInsert(c, key) }
func (c *Cache) KeyForeach(f HFunc, userData T.Gpointer)   { CacheKeyForeach(c, f, userData) }
func (c *Cache) Remove(value T.Gconstpointer)              { CacheRemove(c, value) }
func (c *Cache) ValueForeach(f HFunc, userData T.Gpointer) { CacheValueForeach(c, f, userData) }

type Checksum struct{}

type ChecksumType Enum

const (
	CHECKSUM_MD5 ChecksumType = iota
	CHECKSUM_SHA1
	CHECKSUM_SHA256
)

var (
	ChecksumNew func(checksumType ChecksumType) *Checksum

	ChecksumCopy      func(c *Checksum) *Checksum
	ChecksumFree      func(c *Checksum)
	ChecksumGetDigest func(c *Checksum, buffer *uint8, digestLen *T.Gsize)
	ChecksumGetString func(c *Checksum) string
	ChecksumReset     func(c *Checksum)
	ChecksumUpdate    func(c *Checksum, data *T.Guchar, length T.Gssize)

	ChecksumTypeGetLength func(c ChecksumType) T.Gssize
)

func (c *Checksum) Copy() *Checksum { return ChecksumCopy(c) }
func (c *Checksum) Free()           { ChecksumFree(c) }
func (c *Checksum) GetDigest(buffer *uint8, digestLen *T.Gsize) {
	ChecksumGetDigest(c, buffer, digestLen)
}
func (c *Checksum) GetString() string                      { return ChecksumGetString(c) }
func (c *Checksum) Reset()                                 { ChecksumReset(c) }
func (c *Checksum) Update(data *T.Guchar, length T.Gssize) { ChecksumUpdate(c, data, length) }

func (c ChecksumType) GetLength() T.Gssize { return ChecksumTypeGetLength(c) }

type (
	Completion struct {
		Items       *List
		Fnc         CompletionFunc
		Prefix      *T.Gchar
		Cache       *List
		StrncmpFunc CompletionStrncmpFunc
	}

	CompletionFunc func(T.Gpointer) string

	CompletionStrncmpFunc func(s1, s2 string, n T.Gsize) int
)

var (
	CompletionNew func(f CompletionFunc) *Completion

	CompletionAddItems     func(c *Completion, items *List)
	CompletionClearItems   func(c *Completion)
	CompletionComplete     func(c *Completion, prefix string, newPrefix **T.Gchar) *List
	CompletionCompleteUtf8 func(c *Completion, prefix string, newPrefix **T.Gchar) *List
	CompletionFree         func(c *Completion)
	CompletionRemoveItems  func(c *Completion, items *List)
	CompletionSetCompare   func(c *Completion, strncmpFunc CompletionStrncmpFunc)
)

func (c *Completion) AddItems(items *List) { CompletionAddItems(c, items) }
func (c *Completion) ClearItems()          { CompletionClearItems(c) }
func (c *Completion) Complete(prefix string, newPrefix **T.Gchar) *List {
	return CompletionComplete(c, prefix, newPrefix)
}
func (c *Completion) CompleteUtf8(prefix string, newPrefix **T.Gchar) *List {
	return CompletionCompleteUtf8(c, prefix, newPrefix)
}
func (c *Completion) Free()                   { CompletionFree(c) }
func (c *Completion) RemoveItems(items *List) { CompletionRemoveItems(c, items) }
func (c *Completion) SetCompare(strncmpFunc CompletionStrncmpFunc) {
	CompletionSetCompare(c, strncmpFunc)
}

var (
	ComputeChecksumForData func(c ChecksumType,
		data *T.Guchar, length T.Gsize) string

	ComputeChecksumForString func(c ChecksumType,
		str string, length T.Gssize) string
)

type Cond struct{}
