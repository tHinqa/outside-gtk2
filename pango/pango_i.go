// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

var IsZeroWidth func(ch T.Gunichar) bool

type Item struct {
	Offset   int
	Length   int
	NumChars int
	Analysis Analysis
}

var (
	ItemGetType func() O.Type
	ItemNew     func() *Item

	ItemCopy  func(i *Item) *Item
	ItemFree  func(i *Item)
	ItemSplit func(i *Item, splitIndex, splitOffset int) *Item
)

func (i *Item) Copy() *Item                   { return ItemCopy(i) }
func (i *Item) Free()                         { ItemFree(i) }
func (i *Item) Split(index, offset int) *Item { return ItemSplit(i, index, offset) }

var (
	Itemize            func(context *Context, text string, startIndex, length int, Attrs *AttrList, cachedIter *AttrIterator) *T.GList
	ItemizeWithBaseDir func(context *Context, baseDir Direction, text string, startIndex, length int, Attrs *AttrList, cachedIter *AttrIterator) *T.GList
)
