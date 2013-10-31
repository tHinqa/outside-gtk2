// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type TabAlign Enum

const TAB_LEFT TabAlign = 0

var TabAlignGetType func() O.Type

type TabArray struct{}

var (
	TabArrayGetType          func() O.Type
	TabArrayNew              func(initialSize int, positionsInPixels bool) *TabArray
	TabArrayNewWithPositions func(size int, positionsInPixels bool, firstAlignment TabAlign, firstPosition int, v ...VArg) *TabArray

	tabArrayCopy                 func(t *TabArray) *TabArray
	tabArrayFree                 func(t *TabArray)
	tabArrayGetPositionsInPixels func(t *TabArray) bool
	tabArrayGetSize              func(t *TabArray) int
	tabArrayGetTab               func(t *TabArray, tabIndex int, Alignment *TabAlign, location *int)
	tabArrayGetTabs              func(t *TabArray, Alignments **TabAlign, locations **int)
	tabArrayResize               func(t *TabArray, newSize int)
	tabArraySetTab               func(t *TabArray, tabIndex int, Alignment TabAlign, location int)
)

func (t *TabArray) Copy() *TabArray            { return tabArrayCopy(t) }
func (t *TabArray) Free()                      { tabArrayFree(t) }
func (t *TabArray) GetPositionsInPixels() bool { return tabArrayGetPositionsInPixels(t) }
func (t *TabArray) GetSize() int               { return tabArrayGetSize(t) }
func (t *TabArray) GetTab(tabIndex int, Alignment *TabAlign, location *int) {
	tabArrayGetTab(t, tabIndex, Alignment, location)
}
func (t *TabArray) GetTabs(Alignments **TabAlign, locations **int) {
	tabArrayGetTabs(t, Alignments, locations)
}
func (t *TabArray) Resize(newSize int) { tabArrayResize(t, newSize) }
func (t *TabArray) SetTab(tabIndex int, Alignment TabAlign, location int) {
	tabArraySetTab(t, tabIndex, Alignment, location)
}

var TrimString func(str string) string
